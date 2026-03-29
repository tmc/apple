package balloon

import (
	"fmt"

	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/vm"
)

// Info contains memory balloon device information.
type Info struct {
	HasBalloon     bool
	TargetGB       float64
	MinimumAllowMB uint64
}

// GetInfo retrieves memory balloon information from a running VM.
func GetInfo(queue *vm.Queue, machine vz.VZVirtualMachine) (Info, error) {
	var info Info
	var err error
	queue.Sync(func() {
		devices := machine.MemoryBalloonDevices()
		if len(devices) == 0 {
			return
		}
		info.HasBalloon = true
		device := devices[0]
		traditionalDevice := vz.VZVirtioTraditionalMemoryBalloonDeviceFromID(device.GetID())
		if traditionalDevice.ID == 0 {
			err = fmt.Errorf("balloon device is not a VirtioTraditional type")
			return
		}
		targetBytes := traditionalDevice.TargetVirtualMachineMemorySize()
		info.TargetGB = float64(targetBytes) / (1024 * 1024 * 1024)
		info.MinimumAllowMB = vz.GetVZVirtualMachineConfigurationClass().MinimumAllowedMemorySize() / (1024 * 1024)
	})
	return info, err
}

// SetTarget sets the balloon target size in gigabytes.
func SetTarget(queue *vm.Queue, machine vz.VZVirtualMachine, sizeGB float64) error {
	sizeBytes := uint64(sizeGB * 1024 * 1024 * 1024)
	sizeBytes = (sizeBytes / (1024 * 1024)) * (1024 * 1024)

	var err error
	queue.Sync(func() {
		devices := machine.MemoryBalloonDevices()
		if len(devices) == 0 {
			err = fmt.Errorf("no memory balloon device configured")
			return
		}
		device := devices[0]
		traditionalDevice := vz.VZVirtioTraditionalMemoryBalloonDeviceFromID(device.GetID())
		if traditionalDevice.ID == 0 {
			err = fmt.Errorf("balloon device is not a VirtioTraditional type")
			return
		}
		minSize := vz.GetVZVirtualMachineConfigurationClass().MinimumAllowedMemorySize()
		if sizeBytes < minSize {
			err = fmt.Errorf("target size %.2f GB is below minimum allowed %.2f GB", sizeGB, float64(minSize)/(1024*1024*1024))
			return
		}
		traditionalDevice.SetTargetVirtualMachineMemorySize(sizeBytes)
	})
	return err
}

// AddDevice adds a Virtio traditional memory balloon device to config.
func AddDevice(config vz.VZVirtualMachineConfiguration) {
	balloonConfig := vz.NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration()
	if balloonConfig.ID != 0 {
		config.SetMemoryBalloonDevices([]vz.VZMemoryBalloonDeviceConfiguration{
			balloonConfig.VZMemoryBalloonDeviceConfiguration,
		})
	}
}
