// Code generated from Apple documentation for virtualization. DO NOT EDIT.

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
	rv := objc.SendIfResponds[VZIOUSBHostPassthroughDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZIOUSBHostPassthroughDevice.Configuration]
//   - [VZIOUSBHostPassthroughDevice.SetConfiguration]
//   - [VZIOUSBHostPassthroughDevice.IsPointingDevice]
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
//   - [IVZIOUSBHostPassthroughDevice.Configuration]
//   - [IVZIOUSBHostPassthroughDevice.SetConfiguration]
//   - [IVZIOUSBHostPassthroughDevice.IsPointingDevice]
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
type IVZIOUSBHostPassthroughDevice interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IVZIOUSBHostPassthroughDeviceConfiguration
	SetConfiguration(value IVZIOUSBHostPassthroughDeviceConfiguration)
	IsPointingDevice() bool
	UsbController() IVZUSBController
	SetUsbController(value IVZUSBController)
	Uuid() foundation.NSUUID
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
	InitWithConfigurationError(configuration objectivec.IObject) (VZIOUSBHostPassthroughDevice, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZIOUSBHostPassthroughDevice) Init() VZIOUSBHostPassthroughDevice {
	rv := objc.SendIfResponds[VZIOUSBHostPassthroughDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZIOUSBHostPassthroughDevice) Autorelease() VZIOUSBHostPassthroughDevice {
	rv := objc.SendIfResponds[VZIOUSBHostPassthroughDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZIOUSBHostPassthroughDevice creates a new VZIOUSBHostPassthroughDevice instance.
func NewVZIOUSBHostPassthroughDevice() VZIOUSBHostPassthroughDevice {
	class := getVZIOUSBHostPassthroughDeviceClass()
	rv := objc.SendIfResponds[VZIOUSBHostPassthroughDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZIOUSBHostPassthroughDeviceWithConfigurationError(configuration objectivec.IObject) (VZIOUSBHostPassthroughDevice, error) {
	var errorPtr objc.ID
	instance := getVZIOUSBHostPassthroughDeviceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VZIOUSBHostPassthroughDevice{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return VZIOUSBHostPassthroughDevice{}, objc.ErrInitFailed
	}
	return VZIOUSBHostPassthroughDeviceFromID(rv), nil
}

func (v VZIOUSBHostPassthroughDevice) InitWithConfigurationError(configuration objectivec.IObject) (VZIOUSBHostPassthroughDevice, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](v.ID, objc.Sel("initWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(VZIOUSBHostPassthroughDevice), foundation.NSErrorFrom(errorPtr)
	}
	return VZIOUSBHostPassthroughDeviceFromID(rv), nil

}

func (v VZIOUSBHostPassthroughDevice) Configuration() IVZIOUSBHostPassthroughDeviceConfiguration {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("configuration"))
	return VZIOUSBHostPassthroughDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDevice) SetConfiguration(value IVZIOUSBHostPassthroughDeviceConfiguration) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setConfiguration:"), value)
}
func (v VZIOUSBHostPassthroughDevice) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZIOUSBHostPassthroughDevice) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZIOUSBHostPassthroughDevice) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZIOUSBHostPassthroughDevice) IsPointingDevice() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isPointingDevice"))
	return rv
}
func (v VZIOUSBHostPassthroughDevice) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZIOUSBHostPassthroughDevice) UsbController() IVZUSBController {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDevice) SetUsbController(value IVZUSBController) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setUsbController:"), value)
}
func (v VZIOUSBHostPassthroughDevice) Uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[foundation.NSUUID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUID(rv)
}
func (v VZIOUSBHostPassthroughDevice) VirtualMachine() IVZVirtualMachine {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZIOUSBHostPassthroughDevice) SetVirtualMachine(value IVZVirtualMachine) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}
