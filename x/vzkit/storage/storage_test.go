package storage

import (
	"os"
	"path/filepath"
	"testing"

	vz "github.com/tmc/apple/virtualization"
)

func TestNSDataRoundTrip(t *testing.T) {
	in := []byte("hello")
	out := NSDataToBytes(NSDataFromBytes(in))
	if string(out) != string(in) {
		t.Fatalf("round trip = %q, want %q", out, in)
	}
}

func TestCreateDiskImage(t *testing.T) {
	path := filepath.Join(t.TempDir(), "disk.img")
	if err := CreateDiskImage(path, 1); err != nil {
		t.Fatalf("CreateDiskImage: %v", err)
	}
	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if info.Size() == 0 {
		t.Fatal("CreateDiskImage created zero-sized file")
	}
}

func TestStorageDeviceConstructors(t *testing.T) {
	path := filepath.Join(t.TempDir(), "disk.img")
	if err := CreateDiskImage(path, 1); err != nil {
		t.Fatalf("CreateDiskImage: %v", err)
	}
	attachment, err := CreateDiskAttachment(path, false)
	if err != nil {
		t.Fatalf("CreateDiskAttachment: %v", err)
	}
	base := attachment.VZStorageDeviceAttachment

	tests := []struct {
		name string
		run  func() (vz.VZStorageDeviceConfiguration, error)
	}{
		{
			name: "virtio",
			run: func() (vz.VZStorageDeviceConfiguration, error) {
				device, err := CreateBlockDeviceWithAttachment(base)
				return vz.VZStorageDeviceConfigurationFromID(device.ID), err
			},
		},
		{
			name: "nvme",
			run: func() (vz.VZStorageDeviceConfiguration, error) {
				device, err := CreateNVMeDeviceWithAttachment(base)
				return vz.VZStorageDeviceConfigurationFromID(device.ID), err
			},
		},
		{
			name: "usb",
			run: func() (vz.VZStorageDeviceConfiguration, error) {
				device, err := CreateUSBMassStorageDeviceWithAttachment(base)
				return vz.VZStorageDeviceConfigurationFromID(device.ID), err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			device, err := tt.run()
			if err != nil {
				t.Fatalf("%s: %v", tt.name, err)
			}
			if device.ID == 0 {
				t.Fatalf("%s returned nil device", tt.name)
			}
		})
	}
}

func TestAppendStorageDevices(t *testing.T) {
	path := filepath.Join(t.TempDir(), "disk.img")
	if err := CreateDiskImage(path, 1); err != nil {
		t.Fatalf("CreateDiskImage: %v", err)
	}
	attachment, err := CreateDiskAttachment(path, false)
	if err != nil {
		t.Fatalf("CreateDiskAttachment: %v", err)
	}
	device, err := CreateBlockDeviceWithAttachment(attachment.VZStorageDeviceAttachment)
	if err != nil {
		t.Fatalf("CreateBlockDeviceWithAttachment: %v", err)
	}

	config := vz.NewVZVirtualMachineConfiguration()
	AppendStorageDevices(config, vz.VZStorageDeviceConfigurationFromID(device.ID))
	if got := len(config.StorageDevices()); got != 1 {
		t.Fatalf("storage devices = %d, want 1", got)
	}
}

func TestEnsureUSBController(t *testing.T) {
	config := vz.NewVZVirtualMachineConfiguration()
	EnsureUSBController(config)
	if got := len(config.UsbControllers()); got != 1 {
		t.Fatalf("usb controllers = %d, want 1", got)
	}
	EnsureUSBController(config)
	if got := len(config.UsbControllers()); got != 1 {
		t.Fatalf("usb controllers after second call = %d, want 1", got)
	}
}
