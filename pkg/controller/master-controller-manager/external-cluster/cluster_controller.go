/*
Copyright 2020 The Kubermatic Kubernetes Platform contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package externalcluster

import (
	"context"
	"fmt"
	"time"

	"go.uber.org/zap"

	providerconfig "github.com/kubermatic/machine-controller/pkg/providerconfig/types"
	apiv2 "k8c.io/kubermatic/v2/pkg/api/v2"
	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	kuberneteshelper "k8c.io/kubermatic/v2/pkg/kubernetes"
	"k8c.io/kubermatic/v2/pkg/provider"
	"k8c.io/kubermatic/v2/pkg/provider/cloud/aks"
	"k8c.io/kubermatic/v2/pkg/provider/cloud/eks"
	"k8c.io/kubermatic/v2/pkg/provider/cloud/gke"
	"k8c.io/kubermatic/v2/pkg/resources"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/clientcmd/api"
	"k8s.io/client-go/tools/record"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

const (
	ControllerName = "kkp-external-cluster-controller"
)

// Reconciler is a controller which is responsible for managing clusters.
type Reconciler struct {
	ctrlruntimeclient.Client
	log      *zap.SugaredLogger
	recorder record.EventRecorder
}

// Add creates a cluster controller.
func Add(
	ctx context.Context,
	mgr manager.Manager,
	log *zap.SugaredLogger) error {
	reconciler := &Reconciler{
		log:      log.Named(ControllerName),
		Client:   mgr.GetClient(),
		recorder: mgr.GetEventRecorderFor(ControllerName),
	}
	c, err := controller.New(ControllerName, mgr, controller.Options{Reconciler: reconciler})
	if err != nil {
		return err
	}

	// Watch for changes to ExternalCluster except KubeOne clusters.
	return c.Watch(&source.Kind{Type: &kubermaticv1.ExternalCluster{}}, &handler.EnqueueRequestForObject{}, withEventFilter())
}

func withKubeOnefilter(obj ctrlruntimeclient.Object) bool {
	externalCluster, ok := obj.(*kubermaticv1.ExternalCluster)
	if !ok {
		return false
	}
	return externalCluster.Spec.CloudSpec.KubeOne == nil
}

// ExternalCluster controller doesn't process KubeOne clusters.
func withEventFilter() predicate.Predicate {
	return predicate.Funcs{
		CreateFunc: func(e event.CreateEvent) bool {
			return withKubeOnefilter(e.Object)
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return withKubeOnefilter(e.ObjectNew) && e.ObjectOld.GetGeneration() != e.ObjectNew.GetGeneration()
		},
		DeleteFunc: func(e event.DeleteEvent) bool {
			return withKubeOnefilter(e.Object)
		},
		GenericFunc: func(e event.GenericEvent) bool {
			return withKubeOnefilter(e.Object)
		},
	}
}

func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	resourceName := request.Name
	log := r.log.With("externalcluster", request)
	log.Debug("Processing...")

	cluster := &kubermaticv1.ExternalCluster{}
	if err := r.Get(ctx, ctrlruntimeclient.ObjectKey{Namespace: metav1.NamespaceAll, Name: resourceName}, cluster); err != nil {
		if apierrors.IsNotFound(err) {
			log.Debug("Could not find imported cluster")
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err
	}

	return r.reconcile(ctx, cluster)
}

func (r *Reconciler) handleDeletion(ctx context.Context, cluster *kubermaticv1.ExternalCluster) error {
	if kuberneteshelper.HasFinalizer(cluster, kubermaticv1.ExternalClusterKubeconfigCleanupFinalizer) {
		if err := r.cleanUpKubeconfigSecret(ctx, cluster); err != nil {
			return err
		}
	}
	if kuberneteshelper.HasFinalizer(cluster, kubermaticv1.CredentialsSecretsCleanupFinalizer) {
		if err := r.cleanUpCredentialsSecret(ctx, cluster); err != nil {
			return err
		}
	}
	return nil
}

func (r *Reconciler) reconcile(ctx context.Context, cluster *kubermaticv1.ExternalCluster) (reconcile.Result, error) {
	// handling deletion
	if !cluster.DeletionTimestamp.IsZero() {
		if err := r.handleDeletion(ctx, cluster); err != nil {
			return reconcile.Result{}, fmt.Errorf("handling deletion of externalcluster: %w", err)
		}
		return reconcile.Result{}, nil
	}

	cloud := cluster.Spec.CloudSpec
	secretKeySelector := provider.SecretKeySelectorValueFuncFactory(ctx, r.Client)
	if cloud.ProviderName == "" {
		if cluster.Spec.KubeconfigReference != nil {
			if err := kuberneteshelper.TryAddFinalizer(ctx, r.Client, cluster, kubermaticv1.ExternalClusterKubeconfigCleanupFinalizer); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to add kubeconfig secret finalizer: %w", err)
			}
		}
		return reconcile.Result{}, nil
	}

	if cloud.GKE != nil {
		r.log.Debugf("reconcile GKE cluster")
		if cloud.GKE.CredentialsReference != nil {
			if err := kuberneteshelper.TryAddFinalizer(ctx, r.Client, cluster, kubermaticv1.CredentialsSecretsCleanupFinalizer); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to add credential secret finalizer: %w", err)
			}
		}
		status, err := gke.GetClusterStatus(ctx, secretKeySelector, cloud)
		if err != nil {
			r.log.Debugf("failed to get GKE cluster status %v", err)
			r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
			return reconcile.Result{}, err
		}
		if status.State == apiv2.ProvisioningExternalClusterState {
			// repeat after some time to get/store kubeconfig
			return reconcile.Result{RequeueAfter: time.Second * 10}, err
		}
		if status.State == apiv2.RunningExternalClusterState || status.State == apiv2.ReconcilingExternalClusterState {
			err = r.createOrUpdateGKEKubeconfig(ctx, cluster)
			if err != nil {
				r.log.Errorf("failed to create or update kubeconfig secret %v", err)
				r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
				return reconcile.Result{}, err
			}
			// the kubeconfig token is valid 1h, it will update token every 30min
			return reconcile.Result{RequeueAfter: time.Minute * 30}, nil
		}
	}
	if cloud.EKS != nil {
		r.log.Debugf("reconcile EKS cluster %v", cluster.Name)
		if cloud.EKS.CredentialsReference != nil {
			if err := kuberneteshelper.TryAddFinalizer(ctx, r.Client, cluster, kubermaticv1.CredentialsSecretsCleanupFinalizer); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to add credential secret finalizer: %w", err)
			}
		}
		status, err := eks.GetClusterStatus(secretKeySelector, cloud)
		if err != nil {
			r.log.Debugf("failed to get EKS cluster status %v", err)
			r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
			return reconcile.Result{}, err
		}
		if status.State == apiv2.ProvisioningExternalClusterState {
			// repeat after some time to get/store kubeconfig
			return reconcile.Result{RequeueAfter: time.Second * 10}, err
		}
		if status.State == apiv2.RunningExternalClusterState || status.State == apiv2.ReconcilingExternalClusterState {
			err = r.createOrUpdateEKSKubeconfig(ctx, cluster)
			if err != nil {
				r.log.Errorf("failed to create or update kubeconfig secret %v", err)
				r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
				return reconcile.Result{}, err
			}
			// the kubeconfig token is valid 14min, it will update token every 10min
			return reconcile.Result{RequeueAfter: time.Minute * 10}, nil
		}
	}
	if cloud.AKS != nil {
		r.log.Debugf("reconcile AKS cluster %v", cluster.Name)
		if cloud.AKS.CredentialsReference != nil {
			if err := kuberneteshelper.TryAddFinalizer(ctx, r.Client, cluster, kubermaticv1.CredentialsSecretsCleanupFinalizer); err != nil {
				return reconcile.Result{}, fmt.Errorf("failed to add credential secret finalizer: %w", err)
			}
		}
		status, err := aks.GetClusterStatus(ctx, secretKeySelector, cloud)
		if err != nil {
			r.log.Debugf("failed to get AKS cluster status %v", err)
			r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
			return reconcile.Result{}, err
		}
		if status.State == apiv2.ProvisioningExternalClusterState {
			// repeat after some time to get/store kubeconfig
			return reconcile.Result{RequeueAfter: time.Second * 10}, err
		}
		if status.State == apiv2.RunningExternalClusterState || status.State == apiv2.ReconcilingExternalClusterState {
			err = r.createOrUpdateAKSKubeconfig(ctx, cluster)
			if err != nil {
				r.log.Errorf("failed to create or update kubeconfig secret %v", err)
				r.recorder.Event(cluster, corev1.EventTypeWarning, "ReconcilingError", err.Error())
				return reconcile.Result{}, err
			}
		}
		// reconcile to update kubeconfig for cases like starting a stopped cluster
		return reconcile.Result{RequeueAfter: time.Minute * 2}, nil
	}
	return reconcile.Result{}, nil
}

func (r *Reconciler) cleanUpKubeconfigSecret(ctx context.Context, cluster *kubermaticv1.ExternalCluster) error {
	if err := r.deleteSecret(ctx, cluster.GetKubeconfigSecretName()); err != nil {
		return err
	}

	return kuberneteshelper.TryRemoveFinalizer(ctx, r, cluster, kubermaticv1.ExternalClusterKubeconfigCleanupFinalizer)
}

func (r *Reconciler) cleanUpCredentialsSecret(ctx context.Context, cluster *kubermaticv1.ExternalCluster) error {
	if err := r.deleteSecret(ctx, cluster.GetCredentialsSecretName()); err != nil {
		return err
	}

	return kuberneteshelper.TryRemoveFinalizer(ctx, r, cluster, kubermaticv1.CredentialsSecretsCleanupFinalizer)
}

func (r *Reconciler) deleteSecret(ctx context.Context, secretName string) error {
	if secretName == "" {
		return nil
	}

	secret := &corev1.Secret{}
	name := types.NamespacedName{Name: secretName, Namespace: resources.KubermaticNamespace}
	err := r.Get(ctx, name, secret)
	// Its already gone
	if apierrors.IsNotFound(err) {
		return nil
	}

	// Something failed while loading the secret
	if err != nil {
		return fmt.Errorf("failed to get Secret %q: %w", name.String(), err)
	}

	if err := r.Delete(ctx, secret); err != nil {
		return fmt.Errorf("failed to delete Secret %q: %w", name.String(), err)
	}

	// We successfully deleted the secret
	return nil
}

func createKubeconfigSecret(ctx context.Context, client ctrlruntimeclient.Client, name, projectID string, secretData map[string][]byte) (*providerconfig.GlobalSecretKeySelector, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: resources.KubermaticNamespace,
			Labels:    map[string]string{kubermaticv1.ProjectIDLabelKey: projectID},
		},
		Type: corev1.SecretTypeOpaque,
		Data: secretData,
	}
	if err := client.Create(ctx, secret); err != nil {
		return nil, fmt.Errorf("failed to create kubeconfig secret: %w", err)
	}
	return &providerconfig.GlobalSecretKeySelector{
		ObjectReference: corev1.ObjectReference{
			Name:      name,
			Namespace: resources.KubermaticNamespace,
		},
	}, nil
}

func (r *Reconciler) createOrUpdateGKEKubeconfig(ctx context.Context, cluster *kubermaticv1.ExternalCluster) error {
	cloud := cluster.Spec.CloudSpec
	cred, err := resources.GetGKECredentials(ctx, r.Client, cluster)
	if err != nil {
		return err
	}
	config, err := gke.GetClusterConfig(ctx, cred.ServiceAccount, cloud.GKE.Name, cloud.GKE.Zone)
	if err != nil {
		return err
	}

	return r.updateKubeconfigSecret(ctx, config, cluster)
}

func (r *Reconciler) createOrUpdateEKSKubeconfig(ctx context.Context, cluster *kubermaticv1.ExternalCluster) error {
	cloud := cluster.Spec.CloudSpec
	cred, err := resources.GetEKSCredentials(ctx, r.Client, cluster)
	if err != nil {
		return err
	}
	config, err := eks.GetClusterConfig(ctx, cred.AccessKeyID, cred.SecretAccessKey, cloud.EKS.Name, cloud.EKS.Region)
	if err != nil {
		return err
	}

	return r.updateKubeconfigSecret(ctx, config, cluster)
}

func (r *Reconciler) createOrUpdateAKSKubeconfig(ctx context.Context, cluster *kubermaticv1.ExternalCluster) error {
	cloud := cluster.Spec.CloudSpec
	cred, err := resources.GetAKSCredentials(ctx, r.Client, cluster)
	if err != nil {
		return err
	}
	config, err := aks.GetClusterConfig(ctx, cred, cloud.AKS.Name, cloud.AKS.ResourceGroup)
	if err != nil {
		return err
	}

	return r.updateKubeconfigSecret(ctx, config, cluster)
}

func (r *Reconciler) updateKubeconfigSecret(ctx context.Context, config *api.Config, cluster *kubermaticv1.ExternalCluster) error {
	kubeconfigSecretName := cluster.GetKubeconfigSecretName()
	kubeconfig, err := clientcmd.Write(*config)
	if err != nil {
		return err
	}

	projectID := ""
	if cluster.Labels != nil {
		projectID = cluster.Labels[kubermaticv1.ProjectIDLabelKey]
	}

	namespacedName := types.NamespacedName{Namespace: resources.KubermaticNamespace, Name: kubeconfigSecretName}

	existingSecret := &corev1.Secret{}
	if err := r.Get(ctx, namespacedName, existingSecret); err != nil && !apierrors.IsNotFound(err) {
		return fmt.Errorf("failed to probe for secret %v: %w", namespacedName, err)
	}

	secretData := map[string][]byte{
		resources.ExternalClusterKubeconfig: kubeconfig,
	}

	// update if already exists
	if existingSecret.Name != "" {
		existingSecret.Data = secretData
		r.log.Debugf("update kubeconfig for cluster %s", cluster.Name)

		return r.Update(ctx, existingSecret)
	}

	keyRef, err := createKubeconfigSecret(ctx, r.Client, kubeconfigSecretName, projectID, secretData)
	if err != nil {
		return err
	}
	cluster.Spec.KubeconfigReference = keyRef
	kuberneteshelper.AddFinalizer(cluster, kubermaticv1.ExternalClusterKubeconfigCleanupFinalizer)

	return r.Update(ctx, cluster)
}
