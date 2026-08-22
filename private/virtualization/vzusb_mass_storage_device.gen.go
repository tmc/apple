// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBMassStorageDevice] class.
var (
	_VZUSBMassStorageDeviceClass     VZUSBMassStorageDeviceClass
	_VZUSBMassStorageDeviceClassOnce sync.Once
)

func getVZUSBMassStorageDeviceClass() VZUSBMassStorageDeviceClass {
	_VZUSBMassStorageDeviceClassOnce.Do(func() {
		_VZUSBMassStorageDeviceClass = VZUSBMassStorageDeviceClass{class: objc.GetClass("VZUSBMassStorageDevice")}
	})
	return _VZUSBMassStorageDeviceClass
}

// GetVZUSBMassStorageDeviceClass returns the class object for VZUSBMassStorageDevice.
func GetVZUSBMassStorageDeviceClass() VZUSBMassStorageDeviceClass {
	return getVZUSBMassStorageDeviceClass()
}

type VZUSBMassStorageDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBMassStorageDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBMassStorageDeviceClass) Alloc() VZUSBMassStorageDevice {
	rv := objc.SendIfResponds[VZUSBMassStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBMassStorageDevice.Configuration]
//   - [VZUSBMassStorageDevice.SetConfiguration]
//   - [VZUSBMassStorageDevice.IsPointingDevice]
//   - [VZUSBMassStorageDevice.UsbMassStorageConfiguration]
//   - [VZUSBMassStorageDevice.VirtualMachine]
//   - [VZUSBMassStorageDevice.SetVirtualMachine]
//   - [VZUSBMassStorageDevice.DebugDescription]
//   - [VZUSBMassStorageDevice.Description]
//   - [VZUSBMassStorageDevice.Hash]
//   - [VZUSBMassStorageDevice.Superclass]
//   - [VZUSBMassStorageDevice.UsbController]
//   - [VZUSBMassStorageDevice.SetUsbController]
type VZUSBMassStorageDevice struct {
	VZStorageDevice
}

// VZUSBMassStorageDeviceFromID constructs a [VZUSBMassStorageDevice] from an objc.ID.
func VZUSBMassStorageDeviceFromID(id objc.ID) VZUSBMassStorageDevice {
	return VZUSBMassStorageDevice{VZStorageDevice: VZStorageDeviceFromID(id)}
}

// Ensure VZUSBMassStorageDevice implements IVZUSBMassStorageDevice.
var _ IVZUSBMassStorageDevice = VZUSBMassStorageDevice{}

// An interface definition for the [VZUSBMassStorageDevice] class.
//
// # Methods
//
//   - [IVZUSBMassStorageDevice.Configuration]
//   - [IVZUSBMassStorageDevice.SetConfiguration]
//   - [IVZUSBMassStorageDevice.IsPointingDevice]
//   - [IVZUSBMassStorageDevice.UsbMassStorageConfiguration]
//   - [IVZUSBMassStorageDevice.VirtualMachine]
//   - [IVZUSBMassStorageDevice.SetVirtualMachine]
//   - [IVZUSBMassStorageDevice.DebugDescription]
//   - [IVZUSBMassStorageDevice.Description]
//   - [IVZUSBMassStorageDevice.Hash]
//   - [IVZUSBMassStorageDevice.Superclass]
//   - [IVZUSBMassStorageDevice.UsbController]
//   - [IVZUSBMassStorageDevice.SetUsbController]
type IVZUSBMassStorageDevice interface {
	IVZStorageDevice

	// Topic: Methods

	Configuration() IVZUSBMassStorageDeviceConfiguration
	SetConfiguration(value IVZUSBMassStorageDeviceConfiguration)
	IsPointingDevice() bool
	UsbMassStorageConfiguration() IVZUSBMassStorageDeviceConfiguration
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
	UsbController() IVZUSBController
	SetUsbController(value IVZUSBController)
}

// Init initializes the instance.
func (v VZUSBMassStorageDevice) Init() VZUSBMassStorageDevice {
	rv := objc.SendIfResponds[VZUSBMassStorageDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBMassStorageDevice) Autorelease() VZUSBMassStorageDevice {
	rv := objc.SendIfResponds[VZUSBMassStorageDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDevice creates a new VZUSBMassStorageDevice instance.
func NewVZUSBMassStorageDevice() VZUSBMassStorageDevice {
	class := getVZUSBMassStorageDeviceClass()
	rv := objc.SendIfResponds[VZUSBMassStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZUSBMassStorageDevice) Configuration() IVZUSBMassStorageDeviceConfiguration {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("configuration"))
	return VZUSBMassStorageDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) SetConfiguration(value IVZUSBMassStorageDeviceConfiguration) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setConfiguration:"), value)
}
func (v VZUSBMassStorageDevice) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZUSBMassStorageDevice) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZUSBMassStorageDevice) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZUSBMassStorageDevice) IsPointingDevice() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isPointingDevice"))
	return rv
}
func (v VZUSBMassStorageDevice) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZUSBMassStorageDevice) UsbController() IVZUSBController {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) SetUsbController(value IVZUSBController) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setUsbController:"), value)
}
func (v VZUSBMassStorageDevice) UsbMassStorageConfiguration() IVZUSBMassStorageDeviceConfiguration {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("usbMassStorageConfiguration"))
	return VZUSBMassStorageDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) VirtualMachine() IVZVirtualMachine {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) SetVirtualMachine(value IVZVirtualMachine) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}
