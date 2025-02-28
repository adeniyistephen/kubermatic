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

package certificates

import (
	"errors"
	"fmt"
	"time"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/resources"
	"k8c.io/kubermatic/v2/pkg/resources/certificates/triple"
	"k8c.io/kubermatic/v2/pkg/resources/reconciling"

	corev1 "k8s.io/api/core/v1"
	certutil "k8s.io/client-go/util/cert"
)

// GetCACreator returns a function to create a secret containing a CA with the specified name.
func GetCACreator(commonName string) reconciling.SecretCreator {
	return func(se *corev1.Secret) (*corev1.Secret, error) {
		if se.Data == nil {
			se.Data = map[string][]byte{}
		}

		// if the CA exists, only check if it's expired but never attempt to replace an existing CA
		if certPEM, exists := se.Data[resources.CACertSecretKey]; exists {
			certs, err := certutil.ParseCertsPEM(certPEM)
			if err != nil {
				return se, fmt.Errorf("certificate is not valid PEM-encoded: %w", err)
			}

			if time.Now().After(certs[0].NotAfter) {
				return se, errors.New("certificate has expired")
			}

			return se, nil
		}

		caKp, err := triple.NewCA(commonName)
		if err != nil {
			return nil, fmt.Errorf("unable to create a new CA: %w", err)
		}

		se.Data[resources.CAKeySecretKey] = triple.EncodePrivateKeyPEM(caKp.Key)
		se.Data[resources.CACertSecretKey] = triple.EncodeCertPEM(caKp.Cert)

		return se, nil
	}
}

type caCreatorData interface {
	Cluster() *kubermaticv1.Cluster
}

// RootCACreator returns a function to create a secret with the root ca.
func RootCACreator(data caCreatorData) reconciling.NamedSecretCreatorGetter {
	return func() (string, reconciling.SecretCreator) {
		return resources.CASecretName, GetCACreator(fmt.Sprintf("root-ca.%s", data.Cluster().GetAddress().ExternalName))
	}
}

// FrontProxyCACreator returns a function to create a secret with front proxy ca.
func FrontProxyCACreator() reconciling.NamedSecretCreatorGetter {
	return func() (string, reconciling.SecretCreator) {
		return resources.FrontProxyCASecretName, GetCACreator("front-proxy-ca")
	}
}
