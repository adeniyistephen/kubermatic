# Copyright 2020 The Kubermatic Kubernetes Platform contributors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

DOCKER_REPO ?= "quay.io/kubermatic"

.PHONY: docker
docker:
	docker build -t $(DOCKER_REPO)/addons:$(TAG) .

.PHONY: hubble
hubble: OUTPUT_FILE=hubble/hubble_v1.12.yaml
hubble:
	helm repo add cilium https://helm.cilium.io/ || true
	helm repo update
	cat hubble/_header.txt > $(OUTPUT_FILE)
	helm template cilium cilium/cilium \
	  --version 1.12.0 \
	  --namespace kube-system \
	  --values values-hubble.yaml \
	  >> $(OUTPUT_FILE)
	cat hubble/_extras.txt >> $(OUTPUT_FILE)
	cat hubble/_footer.txt >> $(OUTPUT_FILE)
	sed -i 's/quay.io/{{ Registry "quay.io" }}/g' $(OUTPUT_FILE)

.PHONY: aws-node-termination-handler
aws-node-termination-handler: OUTPUT_FILE=aws-node-termination-handler/aws-node-termination-handler.yaml
aws-node-termination-handler:
	helm repo add eks https://aws.github.io/eks-charts/ || true
	helm repo update
	cat aws-node-termination-handler/_header.txt > $(OUTPUT_FILE)
	helm template aws-node-termination-handler eks/aws-node-termination-handler \
	  --version 0.18.2 \
	  --namespace kube-system \
	  --values values-aws-node-termination-handler.yaml \
	  >> $(OUTPUT_FILE)
	cat aws-node-termination-handler/_footer.txt >> $(OUTPUT_FILE)
	sed -i 's/public.ecr.aws/{{ Registry "public.ecr.aws" }}/g' $(OUTPUT_FILE)

.PHONY: metallb
metallb: OUTPUT_FILE=metallb/00_metallb.yaml
metallb:
	cat metallb/_header.txt > $(OUTPUT_FILE)
	wget -q -O- https://raw.githubusercontent.com/metallb/metallb/v0.13.3/config/manifests/metallb-native.yaml >> $(OUTPUT_FILE)
	sed -i 's/quay.io/{{ Registry "quay.io" }}/g' $(OUTPUT_FILE)	
