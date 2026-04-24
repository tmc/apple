// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZIOUSBHostPassthroughDevice] class.
var (
	_VZIOUSBHostPassthroughDeviceClass     VZIOUSBHostPassthroughDeviceClass
	_VZIOUSBHostPassthroughDeviceClassOnce sync.Once
)

func getVZIOUSBHostPassthroughDeviceClass() VZIOUSBHostPassthroughDeviceClass {
	_VZIOUSBHostPassthroughDeviceClassOnce.Do(func() {
		_VZIOUSBHostPassthroughDeviceClass = VZIOUSBHostPassthroughDeviceClass{class: objc.GetClass("_VZIOUSBHostPassthroughDevice")}
	})
	return _VZIOUSBHostPassthroughDeviceClass
}

// GetVZIOUSBHostPassthroughDeviceClass returns the class object for _VZIOUSBHostPassthroughDevice.
func GetVZIOUSBHostPassthroughDeviceClass() VZIOUSBHostPassthroughDeviceClass {
	return getVZIOUSBHostPassthroughDeviceClass()
}

type VZIOUSBHostPassthroughDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZIOUSBHostPassthroughDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZIOUSBHostPassthroughDeviceClass) Alloc() VZIOUSBHostPassthroughDevice {
	rv := objc.Send[VZIOUSBHostPassthroughDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZIOUSBHostPassthroughDevice._processIOUSBHostDeviceMessageMessageArgumentVirtualMachine]
//   - [VZIOUSBHostPassthroughDevice._releaseDevice]
//   - [VZIOUSBHostPassthroughDevice.Configuration]
//   - [VZIOUSBHostPassthroughDevice.SetConfiguration]
//   - [VZIOUSBHostPassthroughDevice.IoUSBHostDeviceConfiguration]
//   - [VZIOUSBHostPassthroughDevice.IsPointingDevice]
//   - [VZIOUSBHostPassthroughDevice.Signature]
//   - [VZIOUSBHostPassthroughDevice.UsbController]
//   - [VZIOUSBHostPassthroughDevice.SetUsbController]
//   - [VZIOUSBHostPassthroughDevice.Uuid]
//   - [VZIOUSBHostPassthroughDevice.VirtualMachine]
//   - [VZIOUSBHostPassthroughDevice.SetVirtualMachine]
//   - [VZIOUSBHostPassthroughDevice.InitWithConfigurationError]
//   - [VZIOUSBHostPassthroughDevice.DebugDescription]
//   - [VZIOUSBHostPassthroughDevice.Description]
//   - [VZIOUSBHostPassthroughDevice.Hash]
//   - [VZIOUSBHostPassthroughDevice.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice
type VZIOUSBHostPassthroughDevice struct {
	objectivec.Object
}

// VZIOUSBHostPassthroughDeviceFromID constructs a [VZIOUSBHostPassthroughDevice] from an objc.ID.
func VZIOUSBHostPassthroughDeviceFromID(id objc.ID) VZIOUSBHostPassthroughDevice {
	return VZIOUSBHostPassthroughDevice{objectivec.Object{ID: id}}
}

// Ensure VZIOUSBHostPassthroughDevice implements IVZIOUSBHostPassthroughDevice.
var _ IVZIOUSBHostPassthroughDevice = VZIOUSBHostPassthroughDevice{}

// An interface definition for the [VZIOUSBHostPassthroughDevice] class.
//
// # Methods
//
//   - [IVZIOUSBHostPassthroughDevice._processIOUSBHostDeviceMessageMessageArgumentVirtualMachine]
//   - [IVZIOUSBHostPassthroughDevice._releaseDevice]
//   - [IVZIOUSBHostPassthroughDevice.Configuration]
//   - [IVZIOUSBHostPassthroughDevice.SetConfiguration]
//   - [IVZIOUSBHostPassthroughDevice.IoUSBHostDeviceConfiguration]
//   - [IVZIOUSBHostPassthroughDevice.IsPointingDevice]
//   - [IVZIOUSBHostPassthroughDevice.Signature]
//   - [IVZIOUSBHostPassthroughDevice.UsbController]
//   - [IVZIOUSBHostPassthroughDevice.SetUsbController]
//   - [IVZIOUSBHostPassthroughDevice.Uuid]
//   - [IVZIOUSBHostPassthroughDevice.VirtualMachine]
//   - [IVZIOUSBHostPassthroughDevice.SetVirtualMachine]
//   - [IVZIOUSBHostPassthroughDevice.InitWithConfigurationError]
//   - [IVZIOUSBHostPassthroughDevice.DebugDescription]
//   - [IVZIOUSBHostPassthroughDevice.Description]
//   - [IVZIOUSBHostPassthroughDevice.Hash]
//   - [IVZIOUSBHostPassthroughDevice.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice
type IVZIOUSBHostPassthroughDevice interface {
	objectivec.IObject

	// Topic: Methods

	_processIOUSBHostDeviceMessageMessageArgumentVirtualMachine(message uint32, argument unsafe.Pointer, machine objectivec.IObject)
	_releaseDevice()
	Configuration() IVZIOUSBHostPassthroughDeviceConfiguration
	SetConfiguration(value IVZIOUSBHostPassthroughDeviceConfiguration)
	IoUSBHostDeviceConfiguration() IVZIOUSBHostPassthroughDeviceConfiguration
	IsPointingDevice() bool
	Signature() objectivec.IObject
	UsbController() IVZUSBController
	SetUsbController(value IVZUSBController)
	Uuid() foundation.NSUUID
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
	InitWithConfigurationError(configuration objectivec.IObject) (VZIOUSBHostPassthroughDevice, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZIOUSBHostPassthroughDevice) Init() VZIOUSBHostPassthroughDevice {
	rv := objc.Send[VZIOUSBHostPassthroughDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZIOUSBHostPassthroughDevice) Autorelease() VZIOUSBHostPassthroughDevice {
	rv := objc.Send[VZIOUSBHostPassthroughDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZIOUSBHostPassthroughDevice creates a new VZIOUSBHostPassthroughDevice instance.
func NewVZIOUSBHostPassthroughDevice() VZIOUSBHostPassthroughDevice {
	class := getVZIOUSBHostPassthroughDeviceClass()
	rv := objc.Send[VZIOUSBHostPassthroughDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/initWithConfiguration:error:
func NewVZIOUSBHostPassthroughDeviceWithConfigurationError(configuration objectivec.IObject) (VZIOUSBHostPassthroughDevice, error) {
	var errorPtr objc.ID
	instance := getVZIOUSBHostPassthroughDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZIOUSBHostPassthroughDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	return VZIOUSBHostPassthroughDeviceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/_processIOUSBHostDeviceMessage:messageArgument:virtualMachine:
func (v VZIOUSBHostPassthroughDevice) _processIOUSBHostDeviceMessageMessageArgumentVirtualMachine(message uint32, argument unsafe.Pointer, machine objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_processIOUSBHostDeviceMessage:messageArgument:virtualMachine:"), message, argument, machine)
}

// ProcessIOUSBHostDeviceMessageMessageArgumentVirtualMachine is an exported wrapper for the private method _processIOUSBHostDeviceMessageMessageArgumentVirtualMachine.
func (v VZIOUSBHostPassthroughDevice) ProcessIOUSBHostDeviceMessageMessageArgumentVirtualMachine(message uint32, argument unsafe.Pointer, machine objectivec.IObject) {
	v._processIOUSBHostDeviceMessageMessageArgumentVirtualMachine(message, argument, machine)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/_releaseDevice
func (v VZIOUSBHostPassthroughDevice) _releaseDevice() {
	objc.Send[objc.ID](v.ID, objc.Sel("_releaseDevice"))
}

// ReleaseDevice is an exported wrapper for the private method _releaseDevice.
func (v VZIOUSBHostPassthroughDevice) ReleaseDevice() {
	v._releaseDevice()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/signature
func (v VZIOUSBHostPassthroughDevice) Signature() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("signature"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/initWithConfiguration:error:
func (v VZIOUSBHostPassthroughDevice) InitWithConfigurationError(configuration objectivec.IObject) (VZIOUSBHostPassthroughDevice, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZIOUSBHostPassthroughDevice), foundation.NSErrorFrom(errorPtr)
	}
	return VZIOUSBHostPassthroughDeviceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/configuration
func (v VZIOUSBHostPassthroughDevice) Configuration() IVZIOUSBHostPassthroughDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configuration"))
	return VZIOUSBHostPassthroughDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDevice) SetConfiguration(value IVZIOUSBHostPassthroughDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setConfiguration:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/debugDescription
func (v VZIOUSBHostPassthroughDevice) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/description
func (v VZIOUSBHostPassthroughDevice) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/hash
func (v VZIOUSBHostPassthroughDevice) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/ioUSBHostDeviceConfiguration
func (v VZIOUSBHostPassthroughDevice) IoUSBHostDeviceConfiguration() IVZIOUSBHostPassthroughDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ioUSBHostDeviceConfiguration"))
	return VZIOUSBHostPassthroughDeviceConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/isPointingDevice
func (v VZIOUSBHostPassthroughDevice) IsPointingDevice() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPointingDevice"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/superclass
func (v VZIOUSBHostPassthroughDevice) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/usbController
func (v VZIOUSBHostPassthroughDevice) UsbController() IVZUSBController {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDevice) SetUsbController(value IVZUSBController) {
	objc.Send[struct{}](v.ID, objc.Sel("setUsbController:"), value)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/uuid
func (v VZIOUSBHostPassthroughDevice) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOUSBHostPassthroughDevice/virtualMachine
func (v VZIOUSBHostPassthroughDevice) VirtualMachine() IVZVirtualMachine {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDevice) SetVirtualMachine(value IVZVirtualMachine) {
	objc.Send[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}
