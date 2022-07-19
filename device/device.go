package device

import (
	"fmt"
	"github.com/NVIDIA/go-nvml/pkg/nvml"
	"k8s.io/klog/v2"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	MiB = 1
	GiB = 1024 * MiB
)

type Manager interface {
	GetMemoryDevs() []*pluginapi.Device
	GetGPUDevs() []*pluginapi.Device
}

type GPU struct {
	index  int
	memory uint64
}

type GPUManager struct {
	gpus []*GPU
}

var _ Manager = &GPUManager{}

func NewGPUManager() *GPUManager {
	initNVML()
	var gpus []*GPU
	cnt := getDeviceCount()
	for i := 0; i < cnt; i++ {
		mem := getDeviceMemory(i)
		gpu := GPU{
			index:  i,
			memory: mem,
		}
		gpus = append(gpus, &gpu)
	}

	return &GPUManager{
		gpus: gpus,
	}
}

func (m *GPUManager) GetMemoryDevs() []*pluginapi.Device {
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

func (m *GPUManager) GetGPUDevs() []*pluginapi.Device {
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

func initNVML() {
	ret := nvml.Init()
	if ret != nvml.SUCCESS {
		klog.Fatalf("Unable to initialize NVML: %v", nvml.ErrorString(ret))
	}
	defer func() {
		ret := nvml.Shutdown()
		if ret != nvml.SUCCESS {
			klog.Fatalf("Unable to shutdown NVML: %v", nvml.ErrorString(ret))
		}
	}()
}

func getDeviceCount() int {
	count, ret := nvml.DeviceGetCount()
	if ret != nvml.SUCCESS {
		klog.Fatalf("Unable to get device count: %v", nvml.ErrorString(ret))
	}
	return count
}

func getDeviceMemory(idx int) uint64 {
	dev, ret := nvml.DeviceGetHandleByIndex(idx)
	if ret != nvml.SUCCESS {
		klog.Fatalf("Unable to get device by index %v: %v", idx, nvml.ErrorString(ret))
	}

	mem, ret := dev.GetMemoryInfo()
	if ret != nvml.SUCCESS {
		klog.Fatalf("Unable to get device memory: %v", nvml.ErrorString(ret))
	}

	return mem.Total
}
