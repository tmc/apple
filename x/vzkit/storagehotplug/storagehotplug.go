package storagehotplug

import (
	"context"
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	vz "github.com/tmc/apple/virtualization"
	"github.com/tmc/apple/x/vzkit/vm"
)

// Device wraps a live VZStorageDevice.
type Device struct {
	raw pvz.VZStorageDevice
}

// NewDiskImageAttachment opens path as a disk image attachment.
func NewDiskImageAttachment(path string, readOnly bool) (pvz.VZDiskImageStorageDeviceAttachment, error) {
	url := foundation.NewURLFileURLWithPath(path)
	if url.ID == 0 {
		return pvz.VZDiskImageStorageDeviceAttachment{}, fmt.Errorf("create file url")
	}
	attachment, err := vz.NewDiskImageStorageDeviceAttachmentWithURLReadOnlyError(url, readOnly)
	if err != nil {
		return pvz.VZDiskImageStorageDeviceAttachment{}, fmt.Errorf("create disk image attachment: %w", err)
	}
	return pvz.VZDiskImageStorageDeviceAttachmentFromID(attachment.GetID()), nil
}

// NewDevice creates a live storage device for attachment.
func NewDevice(attachment pvz.VZStorageDeviceAttachment) (Device, error) {
	raw := pvz.NewVZStorageDevice().InitWithAttachment(attachment)
	if raw == nil || raw.GetID() == 0 {
		return Device{}, fmt.Errorf("create storage device")
	}
	return Device{raw: pvz.VZStorageDeviceFromID(raw.GetID())}, nil
}

// NewDiskImageDevice opens path and creates a live storage device.
func NewDiskImageDevice(path string, readOnly bool) (Device, pvz.VZDiskImageStorageDeviceAttachment, error) {
	attachment, err := NewDiskImageAttachment(path, readOnly)
	if err != nil {
		return Device{}, pvz.VZDiskImageStorageDeviceAttachment{}, err
	}
	device, err := NewDevice(attachment.VZStorageDeviceAttachment)
	if err != nil {
		return Device{}, pvz.VZDiskImageStorageDeviceAttachment{}, err
	}
	return device, attachment, nil
}

// FromID wraps an existing storage device ID.
func FromID(id objectivec.IObject) Device {
	if id == nil || id.GetID() == 0 {
		return Device{}
	}
	return Device{raw: pvz.VZStorageDeviceFromID(id.GetID())}
}

// Raw returns the underlying Objective-C wrapper.
func (d Device) Raw() pvz.VZStorageDevice {
	return d.raw
}

// Attachment returns the current attachment.
func (d Device) Attachment() objectivec.IObject {
	if d.raw.ID == 0 {
		return nil
	}
	return d.raw.Attachment()
}

// SetVirtualMachine associates the device with a VM.
func (d Device) SetVirtualMachine(machine objectivec.IObject) {
	if d.raw.ID == 0 {
		return
	}
	d.raw.SetVirtualMachine(machine)
}

// SetAttachment swaps the device attachment and waits for completion.
func (d Device) SetAttachment(ctx context.Context, queue *vm.Queue, attachment pvz.VZStorageDeviceAttachment) error {
	return SetAttachment(ctx, queue, d.raw, attachment)
}

// SetAttachment swaps the device attachment and waits for completion.
func SetAttachment(ctx context.Context, queue *vm.Queue, device pvz.VZStorageDevice, attachment pvz.VZStorageDeviceAttachment) error {
	if queue == nil {
		return fmt.Errorf("queue is nil")
	}
	if device.ID == 0 {
		return fmt.Errorf("storage device is nil")
	}
	if attachment.ID == 0 {
		return fmt.Errorf("attachment is nil")
	}

	done := make(chan error, 1)
	queue.Sync(func() {
		device.SetAttachmentCompletionHandler(attachment, func(err error) {
			done <- err
		})
	})

	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("set attachment: %w", err)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UpdateDiskSize updates the size of a disk image attachment.
func UpdateDiskSize(attachment pvz.VZDiskImageStorageDeviceAttachment, size uint64) error {
	if attachment.ID == 0 {
		return fmt.Errorf("disk attachment is nil")
	}
	attachment.UpdateDiskSize(size)
	return nil
}
