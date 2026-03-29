package vzkit

import (
	vz "github.com/tmc/apple/virtualization"
	balloonx "github.com/tmc/apple/x/vzkit/balloon"
)

// BalloonInfo contains memory balloon device information.
type BalloonInfo = balloonx.Info

// GetBalloonInfo retrieves memory balloon information from a running VM.
func GetBalloonInfo(queue *Queue, vm vz.VZVirtualMachine) (BalloonInfo, error) {
	return balloonx.GetInfo(queue, vm)
}

// SetBalloonTarget sets the memory balloon target size in gigabytes.
func SetBalloonTarget(queue *Queue, vm vz.VZVirtualMachine, sizeGB float64) error {
	return balloonx.SetTarget(queue, vm, sizeGB)
}

// AddMemoryBalloonDevice adds a VirtIO traditional memory balloon device.
func AddMemoryBalloonDevice(config vz.VZVirtualMachineConfiguration) { balloonx.AddDevice(config) }
