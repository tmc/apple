// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPointingDevice] class.
var (
	_VZPointingDeviceClass     VZPointingDeviceClass
	_VZPointingDeviceClassOnce sync.Once
)

func getVZPointingDeviceClass() VZPointingDeviceClass {
	_VZPointingDeviceClassOnce.Do(func() {
		_VZPointingDeviceClass = VZPointingDeviceClass{class: objc.GetClass("_VZPointingDevice")}
	})
	return _VZPointingDeviceClass
}

// GetVZPointingDeviceClass returns the class object for _VZPointingDevice.
func GetVZPointingDeviceClass() VZPointingDeviceClass {
	return getVZPointingDeviceClass()
}

type VZPointingDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPointingDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPointingDeviceClass) Alloc() VZPointingDevice {
	rv := objc.Send[VZPointingDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZPointingDevice.SendMagnifyEvents]
//   - [VZPointingDevice.SendQuickLookEvents]
//   - [VZPointingDevice.SendRotationEvents]
//   - [VZPointingDevice.SendScrollWheelEvents]
//   - [VZPointingDevice.SendSmartMagnifyEvents]
//   - [VZPointingDevice.Type]
//   - [VZPointingDevice.InitWithTypeVirtualMachinePointingDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice
type VZPointingDevice struct {
	objectivec.Object
}

// VZPointingDeviceFromID constructs a [VZPointingDevice] from an objc.ID.
func VZPointingDeviceFromID(id objc.ID) VZPointingDevice {
	return VZPointingDevice{objectivec.Object{ID: id}}
}

// Ensure VZPointingDevice implements IVZPointingDevice.
var _ IVZPointingDevice = VZPointingDevice{}

// An interface definition for the [VZPointingDevice] class.
//
// # Methods
//
//   - [IVZPointingDevice.SendMagnifyEvents]
//   - [IVZPointingDevice.SendQuickLookEvents]
//   - [IVZPointingDevice.SendRotationEvents]
//   - [IVZPointingDevice.SendScrollWheelEvents]
//   - [IVZPointingDevice.SendSmartMagnifyEvents]
//   - [IVZPointingDevice.Type]
//   - [IVZPointingDevice.InitWithTypeVirtualMachinePointingDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice
type IVZPointingDevice interface {
	objectivec.IObject

	// Topic: Methods

	SendMagnifyEvents(events objectivec.IObject)
	SendQuickLookEvents(events objectivec.IObject)
	SendRotationEvents(events objectivec.IObject)
	SendScrollWheelEvents(events objectivec.IObject)
	SendSmartMagnifyEvents(events objectivec.IObject)
	Type() int64
	InitWithTypeVirtualMachinePointingDeviceIndex(type_ int64, machine objectivec.IObject, index uint64) VZPointingDevice
}

// Init initializes the instance.
func (v VZPointingDevice) Init() VZPointingDevice {
	rv := objc.Send[VZPointingDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPointingDevice) Autorelease() VZPointingDevice {
	rv := objc.Send[VZPointingDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPointingDevice creates a new VZPointingDevice instance.
func NewVZPointingDevice() VZPointingDevice {
	class := getVZPointingDeviceClass()
	rv := objc.Send[VZPointingDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/initWithType:virtualMachine:pointingDeviceIndex:
func NewVZPointingDeviceWithTypeVirtualMachinePointingDeviceIndex(type_ int64, machine objectivec.IObject, index uint64) VZPointingDevice {
	instance := getVZPointingDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:virtualMachine:pointingDeviceIndex:"), type_, machine, index)
	return VZPointingDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/sendMagnifyEvents:
func (v VZPointingDevice) SendMagnifyEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMagnifyEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/sendQuickLookEvents:
func (v VZPointingDevice) SendQuickLookEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendQuickLookEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/sendRotationEvents:
func (v VZPointingDevice) SendRotationEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendRotationEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/sendScrollWheelEvents:
func (v VZPointingDevice) SendScrollWheelEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendScrollWheelEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/sendSmartMagnifyEvents:
func (v VZPointingDevice) SendSmartMagnifyEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendSmartMagnifyEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/initWithType:virtualMachine:pointingDeviceIndex:
func (v VZPointingDevice) InitWithTypeVirtualMachinePointingDeviceIndex(type_ int64, machine objectivec.IObject, index uint64) VZPointingDevice {
	rv := objc.Send[VZPointingDevice](v.ID, objc.Sel("initWithType:virtualMachine:pointingDeviceIndex:"), type_, machine, index)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/requiresGrabbingMouseInput
func (_VZPointingDeviceClass VZPointingDeviceClass) RequiresGrabbingMouseInput() bool {
	rv := objc.Send[bool](objc.ID(_VZPointingDeviceClass.class), objc.Sel("requiresGrabbingMouseInput"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/type
func (v VZPointingDevice) Type() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("type"))
	return rv
}
