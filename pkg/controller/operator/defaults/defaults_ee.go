//go:build ee

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

package defaults

const (
	DefaultNodeportProxyDockerRepository = "quay.io/kubermatic/nodeport-proxy"

	// DefaultKubermaticImage defines the default Docker repository containing the Kubermatic API image.
	DefaultKubermaticImage = "quay.io/kubermatic/kubermatic-ee"

	// DefaultEtcdLauncherImage defines the default Docker repository containing the etcd launcher image.
	DefaultEtcdLauncherImage = "quay.io/kubermatic/etcd-launcher"

	// DefaultDNATControllerImage defines the default Docker repository containing the DNAT controller image.
	DefaultDNATControllerImage = "quay.io/kubermatic/kubeletdnat-controller"

	// DefaultDashboardAddonImage defines the default Docker repository containing the dashboard image.
	DefaultDashboardImage = "quay.io/kubermatic/dashboard-ee"

	// DefaultKubernetesAddonImage defines the default Docker repository containing the Kubernetes addons.
	DefaultKubernetesAddonImage = "quay.io/kubermatic/addons"
)
