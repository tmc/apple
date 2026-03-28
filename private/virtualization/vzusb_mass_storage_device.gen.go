// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

//
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
func (u VZUSBMassStorageDevice) Init() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBMassStorageDevice) Autorelease() VZUSBMassStorageDevice {
	rv := objc.Send[VZUSBMassStorageDevice](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDevice creates a new VZUSBMassStorageDevice instance.
func NewVZUSBMassStorageDevice() VZUSBMassStorageDevice {
	class := getVZUSBMassStorageDeviceClass()
	rv := objc.Send[VZUSBMassStorageDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/configuration
func (u VZUSBMassStorageDevice) Configuration() IVZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configuration"))
	return VZUSBMassStorageDeviceConfigurationFromID(objc.ID(rv))
}
func (u VZUSBMassStorageDevice) SetConfiguration(value IVZUSBMassStorageDeviceConfiguration) {
	objc.Send[struct{}](u.ID, objc.Sel("setConfiguration:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/debugDescription
func (u VZUSBMassStorageDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/description
func (u VZUSBMassStorageDevice) Description() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/hash
func (u VZUSBMassStorageDevice) Hash() uint64 {
	rv := objc.Send[uint64](u.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/isPointingDevice
func (u VZUSBMassStorageDevice) IsPointingDevice() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("isPointingDevice"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/superclass
func (u VZUSBMassStorageDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](u.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/usbController
func (u VZUSBMassStorageDevice) UsbController() IVZUSBController {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (u VZUSBMassStorageDevice) SetUsbController(value IVZUSBController) {
	objc.Send[struct{}](u.ID, objc.Sel("setUsbController:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/usbMassStorageConfiguration
func (u VZUSBMassStorageDevice) UsbMassStorageConfiguration() IVZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("usbMassStorageConfiguration"))
	return VZUSBMassStorageDeviceConfigurationFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/VZUSBMassStorageDevice/virtualMachine
func (u VZUSBMassStorageDevice) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (u VZUSBMassStorageDevice) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[struct{}](u.ID, objc.Sel("setVirtualMachine:"), value)
}

