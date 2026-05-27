package config

import (
	"path/filepath"
	"testing"

	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/audio"
	"github.com/tmc/apple/x/vzkit/display"
	"github.com/tmc/apple/x/vzkit/network"
	"github.com/tmc/apple/x/vzkit/storage"
	"github.com/tmc/apple/x/vzkit/virtiofs"
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

func TestSetStorageDevices(t *testing.T) {
	path := filepath.Join(t.TempDir(), "disk.img")
	if err := storage.CreateDiskImage(path, 1); err != nil {
		t.Fatalf("CreateDiskImage: %v", err)
	}
	attachment, err := storage.CreateDiskAttachment(path, false)
	if err != nil {
		t.Fatalf("CreateDiskAttachment: %v", err)
	}
	device := storage.CreateBlockDevice(attachment)

	vm := vz.NewVZVirtualMachineConfiguration()
	SetStorageDevices(vm, device)
	if got := len(vm.StorageDevices()); got != 1 {
		t.Fatalf("storage devices = %d, want 1", got)
	}
}

func TestSetMacGraphicsDevices(t *testing.T) {
	device, err := display.CreateMacGraphicsConfig([]display.Config{{Width: 1024, Height: 768, PPI: 144}})
	if err != nil {
		t.Fatalf("CreateMacGraphicsConfig: %v", err)
	}
	displayConfig := vz.NewMacGraphicsDisplayConfigurationWithWidthInPixelsHeightInPixelsPixelsPerInch(1280, 800, 144)

	SetMacGraphicsDisplays(device, displayConfig)
	if got := len(device.Displays()); got != 1 {
		t.Fatalf("mac graphics displays = %d, want 1", got)
	}

	vm := vz.NewVZVirtualMachineConfiguration()
	SetMacGraphicsDevices(vm, device)
	if got := len(vm.GraphicsDevices()); got != 1 {
		t.Fatalf("graphics devices = %d, want 1", got)
	}
}

func TestSetVirtioGraphicsDevices(t *testing.T) {
	device, err := display.CreateVirtioGraphicsConfig([]display.Config{{Width: 1024, Height: 768, PPI: 144}})
	if err != nil {
		t.Fatalf("CreateVirtioGraphicsConfig: %v", err)
	}
	scanout := vz.NewVirtioGraphicsScanoutConfigurationWithWidthInPixelsHeightInPixels(1280, 800)

	SetVirtioScanouts(device, scanout)
	if got := len(device.Scanouts()); got != 1 {
		t.Fatalf("virtio scanouts = %d, want 1", got)
	}

	vm := vz.NewVZVirtualMachineConfiguration()
	SetVirtioGraphicsDevices(vm, device)
	if got := len(vm.GraphicsDevices()); got != 1 {
		t.Fatalf("graphics devices = %d, want 1", got)
	}
}

func TestSetNetworkDevices(t *testing.T) {
	attachment, err := network.CreateNATAttachment(false)
	if err != nil {
		t.Fatalf("CreateNATAttachment: %v", err)
	}
	device := vz.NewVZVirtioNetworkDeviceConfiguration()
	device.SetAttachment(&attachment)

	vm := vz.NewVZVirtualMachineConfiguration()
	SetNetworkDevices(vm, device)
	if got := len(vm.NetworkDevices()); got != 1 {
		t.Fatalf("network devices = %d, want 1", got)
	}
}

func TestSetAudioDevices(t *testing.T) {
	device, err := audio.CreateDevice(audio.Config{OutputEnabled: true})
	if err != nil {
		t.Fatalf("CreateDevice: %v", err)
	}

	vm := vz.NewVZVirtualMachineConfiguration()
	SetAudioDevices(vm, device)
	if got := len(vm.AudioDevices()); got != 1 {
		t.Fatalf("audio devices = %d, want 1", got)
	}
}

func TestSetSerialPorts(t *testing.T) {
	serial := vz.NewVZVirtioConsoleDeviceSerialPortConfiguration()

	vm := vz.NewVZVirtualMachineConfiguration()
	SetSerialPorts(vm, serial.VZSerialPortConfiguration)
	if got := len(vm.SerialPorts()); got != 1 {
		t.Fatalf("serial ports = %d, want 1", got)
	}
}

func TestSetDirectorySharingDevices(t *testing.T) {
	devices, err := virtiofs.CreateDevices([]virtiofs.Mount{{
		HostPath: t.TempDir(),
		Tag:      "share",
	}})
	if err != nil {
		t.Fatalf("CreateDevices: %v", err)
	}

	vm := vz.NewVZVirtualMachineConfiguration()
	SetDirectorySharingDevices(vm, devices)
	if got := len(vm.DirectorySharingDevices()); got != 1 {
		t.Fatalf("directory sharing devices = %d, want 1", got)
	}
}
