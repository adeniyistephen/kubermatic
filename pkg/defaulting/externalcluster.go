/*
Copyright 2022 The Kubermatic Kubernetes Platform contributors.

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

package defaulting

import (
	"context"
	"fmt"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"
)

// DefaultExternalClusterSpec defaults the cluster spec when creating a new external cluster.
// This function assumes that the KubermaticConfiguration has already been defaulted
// (as the KubermaticConfigurationGetter does that automatically).
func DefaultExternalClusterSpec(ctx context.Context, spec *kubermaticv1.ExternalClusterSpec) error {
	// The import cluster feature doesn't expect any cloud spec as this imported cluster could be any type of
	// Kubernetes cluster, thus there are no specs in regard of cloud nor infrastructure. Setting the cluster spec
	// to nil would lead rejecting the cluster so we need to initialize it here.
	if spec == nil {
		spec = &kubermaticv1.ExternalClusterSpec{
			CloudSpec: &kubermaticv1.ExternalClusterCloudSpec{},
		}
	}

	// Ensure provider name matches the given spec
	providerName, err := provider.ExternalClusterCloudProviderName(spec.CloudSpec)
	if err != nil {
		return fmt.Errorf("failed to determine cloud provider: %w", err)
	}

	spec.CloudSpec.ProviderName = kubermaticv1.ExternalClusterProvider(providerName)

	return nil
}
