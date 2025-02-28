//go:build !ee

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

package test

import (
	"k8c.io/kubermatic/v2/pkg/provider"
	"k8c.io/kubermatic/v2/pkg/provider/kubernetes"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

func resourceQuotaProviderFactory(_ kubernetes.ImpersonationClient, _ ctrlruntimeclient.Client) provider.ResourceQuotaProvider {
	return nil
}

func groupProjectBindingProviderFactory(createMasterImpersonatedClient kubernetes.ImpersonationClient, privilegedClient ctrlruntimeclient.Client) provider.GroupProjectBindingProvider {
	return nil
}
