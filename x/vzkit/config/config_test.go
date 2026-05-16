package config

import (
	"testing"

	vz "github.com/tmc/apple/virtualization"
)

func TestSetBasicDevices(t *testing.T) {
	vm := vz.NewVZVirtualMachineConfiguration()

	entropy := vz.NewVZVirtioEntropyDeviceConfiguration()
	SetEntropyDevices(vm, entropy)
	if got := len(vm.EntropyDevices()); got != 1 {
		t.Fatalf("entropy devices = %d, want 1", got)
	}

	balloon := vz.NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration()
	SetMemoryBalloonDevices(vm, balloon)
	if got := len(vm.MemoryBalloonDevices()); got != 1 {
		t.Fatalf("memory balloon devices = %d, want 1", got)
	}

	socket := vz.NewVZVirtioSocketDeviceConfiguration()
	SetSocketDevices(vm, socket)
	if got := len(vm.SocketDevices()); got != 1 {
		t.Fatalf("socket devices = %d, want 1", got)
	}
}

func TestSetInputDevices(t *testing.T) {
	vm := vz.NewVZVirtualMachineConfiguration()

	keyboard := vz.NewVZUSBKeyboardConfiguration()
	SetKeyboards(vm, keyboard)
	if got := len(vm.Keyboards()); got != 1 {
		t.Fatalf("keyboards = %d, want 1", got)
	}

	pointing := vz.NewVZUSBScreenCoordinatePointingDeviceConfiguration()
	SetPointingDevices(vm, []vz.IVZPointingDeviceConfiguration{pointing})
	if got := len(vm.PointingDevices()); got != 1 {
		t.Fatalf("pointing devices = %d, want 1", got)
	}
}
