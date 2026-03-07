package vzkit

import (
	"fmt"

	vz "github.com/tmc/apple/virtualization"
)

// BalloonInfo contains memory balloon device information.
type BalloonInfo struct {
	HasBalloon     bool
	TargetGB       float64
	MinimumAllowMB uint64
}

// BalloonInfo retrieves memory balloon information from a running VM.
// The call dispatches to the VM's queue.
func GetBalloonInfo(queue *Queue, vm vz.VZVirtualMachine) (BalloonInfo, error) {
	var info BalloonInfo
	var err error

	queue.Sync(func() {
		devices := vm.MemoryBalloonDevices()
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

// SetBalloonTarget sets the memory balloon target size in gigabytes.
// The size is rounded down to a 1 MB boundary.
func SetBalloonTarget(queue *Queue, vm vz.VZVirtualMachine, sizeGB float64) error {
	sizeBytes := uint64(sizeGB * 1024 * 1024 * 1024)
	sizeBytes = (sizeBytes / (1024 * 1024)) * (1024 * 1024)

	var err error
	queue.Sync(func() {
		devices := vm.MemoryBalloonDevices()
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
			err = fmt.Errorf("target size %.2f GB is below minimum allowed %.2f GB",
				sizeGB, float64(minSize)/(1024*1024*1024))
			return
		}

		traditionalDevice.SetTargetVirtualMachineMemorySize(sizeBytes)
	})

	return err
}

// AddMemoryBalloonDevice adds a VirtIO traditional memory balloon device
// to a VM configuration. This enables runtime memory control.
func AddMemoryBalloonDevice(config vz.VZVirtualMachineConfiguration) {
	balloonConfig := vz.NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration()
	if balloonConfig.ID != 0 {
		config.SetMemoryBalloonDevices([]vz.VZMemoryBalloonDeviceConfiguration{
			balloonConfig.VZMemoryBalloonDeviceConfiguration,
		})
	}
}
