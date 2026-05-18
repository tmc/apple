// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(vc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice
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
	Superclass() objc.Class
	UsbController() IVZUSBController
	SetUsbController(value IVZUSBController)
}

// Init initializes the instance.
func (v VZUSBMassStorageDevice) Init() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBMassStorageDevice) Autorelease() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDevice creates a new VZUSBMassStorageDevice instance.
func NewVZUSBMassStorageDevice() VZUSBMassStorageDevice {
	class := getVZUSBMassStorageDeviceClass()
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/configuration
func (v VZUSBMassStorageDevice) Configuration() IVZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configuration"))
	return VZUSBMassStorageDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) SetConfiguration(value IVZUSBMassStorageDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setConfiguration:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/debugDescription
func (v VZUSBMassStorageDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/description
func (v VZUSBMassStorageDevice) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/hash
func (v VZUSBMassStorageDevice) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/isPointingDevice
func (v VZUSBMassStorageDevice) IsPointingDevice() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPointingDevice"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/superclass
func (v VZUSBMassStorageDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/usbController
func (v VZUSBMassStorageDevice) UsbController() IVZUSBController {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) SetUsbController(value IVZUSBController) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsbController:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/usbMassStorageConfiguration
func (v VZUSBMassStorageDevice) UsbMassStorageConfiguration() IVZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("usbMassStorageConfiguration"))
	return VZUSBMassStorageDeviceConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/virtualMachine
func (v VZUSBMassStorageDevice) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZUSBMassStorageDevice) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}
