# Flex GPU Device Plugin

Kubernetes device plugin for multiple container nvidia gpu share.

⚠This project is for code example purpose, if you want to use it in production, I am happy to provide support, feel
free to contact me.

Related project: [WLBF/flex-gpu-scheduler](https://github.com/WLBF/flex-gpu-scheduler)

## Overview

Flex GPU device plugin will detect nvidia gpu and register two type of resource for each gpu.

* `nvidia.flex.com/gpu` is for exclusively gpu usage
  like [NVIDIA/k8s-device-plugin](https://github.com/NVIDIA/k8s-device-plugin).

* `nvidia.flex.com/memory` is for gpu share usage. For now gpu memory resource unit is GiB.

### Example

The kubectl describe command show the node `v124-worker-0` has 3 gpu and 8 GiB memory each gpu, 24 GiB in total.

```
# kubectl describe no v124-worker-0

...
Capacity:
  cpu:                     2
  ephemeral-storage:       4893836Ki
  hugepages-1Gi:           0
  hugepages-2Mi:           0
  memory:                  4026052Ki
  nvidia.flex.com/gpu:     3
  nvidia.flex.com/memory:  24
  pods:                    110
Allocatable:
  cpu:                     2
  ephemeral-storage:       4510159251
  hugepages-1Gi:           0
  hugepages-2Mi:           0
  memory:                  3923652Ki
  nvidia.flex.com/gpu:     3
  nvidia.flex.com/memory:  24
  pods:                    110
...
```

## Install

Device plugin can be installed by helm chart. For development use `values.dev.yaml` instead of `values.pord.yaml`.

```
helm install flex-gpu-device-plugin -f  ./manifests/flexgpu/values.prod.yaml ./manifests/flexgpu
```
