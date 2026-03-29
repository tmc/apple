// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBPassthroughDevice] class.
var (
	_VZUSBPassthroughDeviceClass     VZUSBPassthroughDeviceClass
	_VZUSBPassthroughDeviceClassOnce sync.Once
)

func getVZUSBPassthroughDeviceClass() VZUSBPassthroughDeviceClass {
	_VZUSBPassthroughDeviceClassOnce.Do(func() {
		_VZUSBPassthroughDeviceClass = VZUSBPassthroughDeviceClass{class: objc.GetClass("_VZUSBPassthroughDevice")}
	})
	return _VZUSBPassthroughDeviceClass
}

// GetVZUSBPassthroughDeviceClass returns the class object for _VZUSBPassthroughDevice.
func GetVZUSBPassthroughDeviceClass() VZUSBPassthroughDeviceClass {
	return getVZUSBPassthroughDeviceClass()
}

type VZUSBPassthroughDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBPassthroughDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBPassthroughDeviceClass) Alloc() VZUSBPassthroughDevice {
	rv := objc.Send[VZUSBPassthroughDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZUSBPassthroughDevice.Configuration]
//   - [VZUSBPassthroughDevice.SetConfiguration]
//   - [VZUSBPassthroughDevice.IsPointingDevice]
//   - [VZUSBPassthroughDevice.Signature]
//   - [VZUSBPassthroughDevice.UsbController]
//   - [VZUSBPassthroughDevice.SetUsbController]
//   - [VZUSBPassthroughDevice.Uuid]
//   - [VZUSBPassthroughDevice.VirtualMachine]
//   - [VZUSBPassthroughDevice.SetVirtualMachine]
//   - [VZUSBPassthroughDevice.InitWithConfigurationError]
//   - [VZUSBPassthroughDevice.DebugDescription]
//   - [VZUSBPassthroughDevice.Description]
//   - [VZUSBPassthroughDevice.Hash]
//   - [VZUSBPassthroughDevice.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice
type VZUSBPassthroughDevice struct {
	objectivec.Object
}

// VZUSBPassthroughDeviceFromID constructs a [VZUSBPassthroughDevice] from an objc.ID.
func VZUSBPassthroughDeviceFromID(id objc.ID) VZUSBPassthroughDevice {
	return VZUSBPassthroughDevice{objectivec.Object{ID: id}}
}
// Ensure VZUSBPassthroughDevice implements IVZUSBPassthroughDevice.
var _ IVZUSBPassthroughDevice = VZUSBPassthroughDevice{}

// An interface definition for the [VZUSBPassthroughDevice] class.
//
// # Methods
//
//   - [IVZUSBPassthroughDevice.Configuration]
//   - [IVZUSBPassthroughDevice.SetConfiguration]
//   - [IVZUSBPassthroughDevice.IsPointingDevice]
//   - [IVZUSBPassthroughDevice.Signature]
//   - [IVZUSBPassthroughDevice.UsbController]
//   - [IVZUSBPassthroughDevice.SetUsbController]
//   - [IVZUSBPassthroughDevice.Uuid]
//   - [IVZUSBPassthroughDevice.VirtualMachine]
//   - [IVZUSBPassthroughDevice.SetVirtualMachine]
//   - [IVZUSBPassthroughDevice.InitWithConfigurationError]
//   - [IVZUSBPassthroughDevice.DebugDescription]
//   - [IVZUSBPassthroughDevice.Description]
//   - [IVZUSBPassthroughDevice.Hash]
//   - [IVZUSBPassthroughDevice.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice
type IVZUSBPassthroughDevice interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() *VZUSBPassthroughDeviceConfiguration
	SetConfiguration(value *VZUSBPassthroughDeviceConfiguration)
	IsPointingDevice() bool
	Signature() objectivec.IObject
	UsbController() IVZUSBController
	SetUsbController(value IVZUSBController)
	Uuid() foundation.NSUUID
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
	InitWithConfigurationError(configuration objectivec.IObject) (VZUSBPassthroughDevice, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZUSBPassthroughDevice) Init() VZUSBPassthroughDevice {
	rv := objc.Send[VZUSBPassthroughDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBPassthroughDevice) Autorelease() VZUSBPassthroughDevice {
	rv := objc.Send[VZUSBPassthroughDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBPassthroughDevice creates a new VZUSBPassthroughDevice instance.
func NewVZUSBPassthroughDevice() VZUSBPassthroughDevice {
	class := getVZUSBPassthroughDeviceClass()
	rv := objc.Send[VZUSBPassthroughDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/initWithConfiguration:error:
func NewVZUSBPassthroughDeviceWithConfigurationError(configuration objectivec.IObject) (VZUSBPassthroughDevice, error) {
	var errorPtr objc.ID
	instance := getVZUSBPassthroughDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZUSBPassthroughDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZUSBPassthroughDeviceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/signature
func (v VZUSBPassthroughDevice) Signature() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("signature"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/initWithConfiguration:error:
func (v VZUSBPassthroughDevice) InitWithConfigurationError(configuration objectivec.IObject) (VZUSBPassthroughDevice, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZUSBPassthroughDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZUSBPassthroughDeviceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/configuration
func (v VZUSBPassthroughDevice) Configuration() *VZUSBPassthroughDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configuration"))
	if rv == 0 {
		return nil
	}
	val := VZUSBPassthroughDeviceConfigurationFromID(objc.ID(rv))
	return &val
}
func (v VZUSBPassthroughDevice) SetConfiguration(value *VZUSBPassthroughDeviceConfiguration) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setConfiguration:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setConfiguration:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/debugDescription
func (v VZUSBPassthroughDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/description
func (v VZUSBPassthroughDevice) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/hash
func (v VZUSBPassthroughDevice) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/isPointingDevice
func (v VZUSBPassthroughDevice) IsPointingDevice() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPointingDevice"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/superclass
func (v VZUSBPassthroughDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/usbController
func (v VZUSBPassthroughDevice) UsbController() IVZUSBController {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (v VZUSBPassthroughDevice) SetUsbController(value IVZUSBController) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsbController:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/uuid
func (v VZUSBPassthroughDevice) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBPassthroughDevice/virtualMachine
func (v VZUSBPassthroughDevice) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZUSBPassthroughDevice) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}

