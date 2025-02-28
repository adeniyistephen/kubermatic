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

package validation

import (
	"context"
	"errors"
	"fmt"
	"math"
	"net"
	"strings"

	kubermaticv1 "k8c.io/kubermatic/v2/pkg/apis/kubermatic/v1"
	"k8c.io/kubermatic/v2/pkg/provider"

	"k8s.io/apimachinery/pkg/runtime"
	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// validator for validating Resource Quota CRD.
type validator struct {
	seedGetter       provider.SeedGetter
	seedClientGetter provider.SeedClientGetter
}

// NewValidator returns a new Resource Quota validator.
func NewValidator(seedGetter provider.SeedGetter, seedClientGetter provider.SeedClientGetter) *validator {
	return &validator{
		seedGetter:       seedGetter,
		seedClientGetter: seedClientGetter,
	}
}

var _ admission.CustomValidator = &validator{}

func (v *validator) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return v.validate(ctx, obj)
}

func (v *validator) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	if err := v.validate(ctx, newObj); err != nil {
		return err
	}

	newIPAMPool := newObj.(*kubermaticv1.IPAMPool)
	oldIPAMPool := oldObj.(*kubermaticv1.IPAMPool)

	// loop old IPAMPool datacenters
	for dc, dcOldConfig := range oldIPAMPool.Spec.Datacenters {
		dcNewConfig, dcExistsInNewPool := newIPAMPool.Spec.Datacenters[dc]
		if !dcExistsInNewPool {
			err := v.validateDCRemoval(ctx, oldIPAMPool, dc)
			if err != nil {
				return err
			}
			continue
		}

		if dcOldConfig.PoolCIDR != dcNewConfig.PoolCIDR {
			return errors.New("it's not allowed to update the pool CIDR for a datacenter")
		}

		if dcOldConfig.Type != dcNewConfig.Type {
			return errors.New("it's not allowed to update the allocation type for a datacenter")
		}

		switch dcOldConfig.Type {
		case kubermaticv1.IPAMPoolAllocationTypeRange:
			if dcOldConfig.AllocationRange != dcNewConfig.AllocationRange {
				return errors.New("it's not allowed to update the allocation range for a datacenter")
			}
		case kubermaticv1.IPAMPoolAllocationTypePrefix:
			if dcOldConfig.AllocationPrefix != dcNewConfig.AllocationPrefix {
				return errors.New("it's not allowed to update the allocation prefix for a datacenter")
			}
		}
	}

	return nil
}

func (v *validator) ValidateDelete(_ context.Context, _ runtime.Object) error {
	// NOP we allow delete operation
	return nil
}

func (v *validator) validate(ctx context.Context, obj runtime.Object) error {
	ipamPool, ok := obj.(*kubermaticv1.IPAMPool)
	if !ok {
		return errors.New("object is not a IPAMPool")
	}

	for _, dcConfig := range ipamPool.Spec.Datacenters {
		_, poolSubnet, err := net.ParseCIDR(string(dcConfig.PoolCIDR))
		if err != nil {
			return err
		}
		poolPrefix, bits := poolSubnet.Mask.Size()

		switch dcConfig.Type {
		case kubermaticv1.IPAMPoolAllocationTypeRange:
			if dcConfig.AllocationRange <= 0 {
				return errors.New("allocation range should be greater than zero")
			}

			numberOfPoolSubnetIPsFloat64 := math.Pow(2, float64(bits-poolPrefix))
			numberOfPoolSubnetIPs := int(numberOfPoolSubnetIPsFloat64)
			if float64(numberOfPoolSubnetIPs) != numberOfPoolSubnetIPsFloat64 {
				return errors.New("the pool is too big to be processed")
			}

			if bits-poolPrefix > 12 {
				return errors.New("pool prefix is too low for range allocation type")
			}

			if dcConfig.AllocationRange > numberOfPoolSubnetIPs {
				return errors.New("allocation range cannot be greater than the pool subnet possible number of IP addresses")
			}
		case kubermaticv1.IPAMPoolAllocationTypePrefix:
			if dcConfig.AllocationPrefix < poolPrefix {
				return errors.New("allocation prefix cannot be smaller than the pool subnet mask size")
			}
			if dcConfig.AllocationPrefix > bits {
				return errors.New("invalid allocation prefix for IP version")
			}
		}
	}

	return nil
}

func (v *validator) validateDCRemoval(ctx context.Context, ipamPool *kubermaticv1.IPAMPool, dc string) error {
	seedClient, err := v.getSeedClient(ctx)
	if err != nil {
		return err
	}

	// List all IPAM allocations
	ipamAllocationList := &kubermaticv1.IPAMAllocationList{}
	err = seedClient.List(ctx, ipamAllocationList)
	if err != nil {
		return fmt.Errorf("failed to list IPAM allocations: %w", err)
	}

	// Iterate current IPAM allocations to check if there is an allocation for the pool to be deleted
	var ipamAllocationsNamespaces []string
	for _, ipamAllocation := range ipamAllocationList.Items {
		if ipamAllocation.Name == ipamPool.Name && ipamAllocation.Spec.DC == dc {
			ipamAllocationsNamespaces = append(ipamAllocationsNamespaces, ipamAllocation.Namespace)
		}
	}

	if len(ipamAllocationsNamespaces) > 0 {
		return fmt.Errorf("cannot delete some datacenter IPAMPool because there is existing IPAMAllocation in namespaces (%s)", strings.Join(ipamAllocationsNamespaces, ", "))
	}

	return nil
}

func (v *validator) getSeedClient(ctx context.Context) (ctrlruntimeclient.Client, error) {
	seed, err := v.seedGetter()
	if err != nil {
		return nil, fmt.Errorf("failed to get current seed: %w", err)
	}
	if seed == nil {
		return nil, errors.New("webhook not configured for a seed cluster")
	}

	client, err := v.seedClientGetter(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to get seed client: %w", err)
	}

	return client, nil
}
