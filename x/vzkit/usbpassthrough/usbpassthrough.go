package usbpassthrough

import (
	"context"
	"fmt"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	pvz "github.com/tmc/apple/private/virtualization"
	"github.com/tmc/apple/x/vzkit"
)

// HostDevice wraps a live host USB passthrough device.
type HostDevice struct {
	raw pvz.VZIOUSBHostPassthroughDevice
}

// Controller wraps a live XHCI controller.
type Controller struct {
	raw pvz.VZXHCIController
}

// NewHostPassthroughConfigurationFromService creates a host passthrough configuration from a service ID.
func NewHostPassthroughConfigurationFromService(service uint32) (pvz.VZIOUSBHostPassthroughDeviceConfiguration, error) {
	return pvz.NewVZIOUSBHostPassthroughDeviceConfigurationWithServiceError(service)
}

// NewHostPassthroughConfigurationFromLocationID creates a host passthrough configuration from a location ID.
func NewHostPassthroughConfigurationFromLocationID(locationID uint32) (pvz.VZIOUSBHostPassthroughDeviceConfiguration, error) {
	cfg, err := pvz.GetVZIOUSBHostPassthroughDeviceConfigurationClass().FromLocationIDError(locationID)
	if err != nil {
		return pvz.VZIOUSBHostPassthroughDeviceConfiguration{}, fmt.Errorf("create host passthrough configuration: %w", err)
	}
	if cfg == nil || cfg.GetID() == 0 {
		return pvz.VZIOUSBHostPassthroughDeviceConfiguration{}, fmt.Errorf("create host passthrough configuration")
	}
	return pvz.VZIOUSBHostPassthroughDeviceConfigurationFromID(cfg.GetID()), nil
}

// NewHostDevice creates a passthrough device from configuration.
func NewHostDevice(cfg pvz.VZIOUSBHostPassthroughDeviceConfiguration) (HostDevice, error) {
	if cfg.ID == 0 {
		return HostDevice{}, fmt.Errorf("configuration is nil")
	}
	raw, err := pvz.NewVZIOUSBHostPassthroughDeviceWithConfigurationError(cfg)
	if err != nil {
		return HostDevice{}, fmt.Errorf("create host passthrough device: %w", err)
	}
	return HostDevice{raw: raw}, nil
}

// NewHostDeviceWithLocationID is a convenience wrapper for location-based lookup.
func NewHostDeviceWithLocationID(locationID uint32) (HostDevice, error) {
	cfg, err := NewHostPassthroughConfigurationFromLocationID(locationID)
	if err != nil {
		return HostDevice{}, err
	}
	return NewHostDevice(cfg)
}

// NewHostDeviceWithService is a convenience wrapper for service-based lookup.
func NewHostDeviceWithService(service uint32) (HostDevice, error) {
	cfg, err := NewHostPassthroughConfigurationFromService(service)
	if err != nil {
		return HostDevice{}, err
	}
	return NewHostDevice(cfg)
}

// Raw returns the underlying host passthrough device.
func (d HostDevice) Raw() pvz.VZIOUSBHostPassthroughDevice {
	return d.raw
}

// Configuration returns the attached configuration, if any.
func (d HostDevice) Configuration() *pvz.VZIOUSBHostPassthroughDeviceConfiguration {
	return d.raw.Configuration()
}

// Release asks the device to release the host USB device.
func (d HostDevice) Release() {
	if d.raw.ID == 0 {
		return
	}
	d.raw.ReleaseDevice()
}

// DeviceSignature returns the opaque signature bytes for a host passthrough configuration.
func DeviceSignature(cfg pvz.VZIOUSBHostPassthroughDeviceConfiguration) []byte {
	if cfg.ID == 0 {
		return nil
	}
	sig := cfg.Signature()
	if sig == nil || sig.GetID() == 0 {
		return nil
	}
	return vzkit.NSDataToBytes(foundation.NSDataFromID(sig.GetID()))
}

// NewUSBPassthroughConfigurationFromDevice creates a passthrough configuration for an accessory device.
func NewUSBPassthroughConfigurationFromDevice(device objectivec.IObject) (pvz.VZUSBPassthroughDeviceConfiguration, error) {
	if device == nil || device.GetID() == 0 {
		return pvz.VZUSBPassthroughDeviceConfiguration{}, fmt.Errorf("device is nil")
	}
	cfg := pvz.NewVZUSBPassthroughDeviceConfigurationWithDevice(device)
	if cfg.ID == 0 {
		return pvz.VZUSBPassthroughDeviceConfiguration{}, fmt.Errorf("create passthrough configuration")
	}
	return cfg, nil
}

// NewControllerConfiguration creates an XHCI controller configuration.
func NewControllerConfiguration() pvz.VZXHCIControllerConfiguration {
	return pvz.NewVZXHCIControllerConfiguration()
}

// NewController creates a live XHCI controller for a VM.
func NewController(config pvz.VZXHCIControllerConfiguration, vm objectivec.IObject, index uint64, devices objectivec.IObject) (Controller, error) {
	if config.ID == 0 {
		return Controller{}, fmt.Errorf("controller configuration is nil")
	}
	raw := config.MakeUSBControllerForVirtualMachineUsbControllerIndexUsbDevices(vm, index, devices)
	if raw == nil || raw.GetID() == 0 {
		return Controller{}, fmt.Errorf("create xhci controller")
	}
	return Controller{raw: pvz.VZXHCIControllerFromID(raw.GetID())}, nil
}

// Raw returns the underlying controller.
func (c Controller) Raw() pvz.VZXHCIController {
	return c.raw
}

// Capture asks the controller to capture host passthrough devices.
func (c Controller) Capture(ctx context.Context) error {
	if c.raw.ID == 0 {
		return fmt.Errorf("controller is nil")
	}
	done := make(chan error, 1)
	c.raw.CapturePassthroughDevicesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("capture passthrough devices: %w", err)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Release releases captured passthrough devices.
func (c Controller) Release() {
	if c.raw.ID == 0 {
		return
	}
	c.raw.ReleasePassthroughDevices()
}

// Attach attaches a device to the controller.
func (c Controller) Attach(ctx context.Context, device objectivec.IObject) error {
	if c.raw.ID == 0 {
		return fmt.Errorf("controller is nil")
	}
	if device == nil || device.GetID() == 0 {
		return fmt.Errorf("device is nil")
	}
	done := make(chan error, 1)
	c.raw.AttachDeviceCompletionHandler(device, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("attach device: %w", err)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Detach detaches a device from the controller.
func (c Controller) Detach(ctx context.Context, device objectivec.IObject) error {
	if c.raw.ID == 0 {
		return fmt.Errorf("controller is nil")
	}
	if device == nil || device.GetID() == 0 {
		return fmt.Errorf("device is nil")
	}
	done := make(chan error, 1)
	c.raw.DetachDeviceCompletionHandler(device, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		if err != nil {
			return fmt.Errorf("detach device: %w", err)
		}
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
