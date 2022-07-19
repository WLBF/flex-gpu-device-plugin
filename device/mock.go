package device

import (
	"fmt"
	"k8s.io/klog/v2"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
	"strconv"
	"strings"
)

type MockManager struct {
	gpus []*GPU
}

var _ Manager = &MockManager{}

func NewMockManager(devs string) *MockManager {
	strs := strings.Split(devs, ",")
	var gpus []*GPU
	for i, str := range strs {
		mem, err := strconv.ParseUint(str, 10, 64)
		if err != nil {
			klog.Fatal("invalid number str", err)
		}
		klog.V(6).InfoS("mock devices", "index", i, "memory", mem)
		gpu := GPU{
			index:  i,
			memory: mem,
		}
		gpus = append(gpus, &gpu)
	}
	return &MockManager{
		gpus: gpus,
	}
}

func (m *MockManager) GetMemoryDevs() []*pluginapi.Device {
	var devs []*pluginapi.Device
	for _, gpu := range m.gpus {
		// minimum unit GiB

		sz := gpu.memory / GiB

		klog.V(6).InfoS("device memory size", "index", gpu.index, "size", sz)
		for j := uint64(0); j < sz; j++ {
			dev := pluginapi.Device{
				ID:     fmt.Sprintf("MEM-%d-%d", gpu.index, j),
				Health: pluginapi.Healthy,
			}
			devs = append(devs, &dev)
		}
	}

	klog.V(6).InfoS("total device memory size", "size", len(devs))
	return devs
}

func (m *MockManager) GetGPUDevs() []*pluginapi.Device {
	var devs []*pluginapi.Device
	for _, gpu := range m.gpus {
		dev := pluginapi.Device{
			ID:     fmt.Sprintf("GPU-%d", gpu.index),
			Health: pluginapi.Healthy,
		}
		devs = append(devs, &dev)
	}
	return devs
}
