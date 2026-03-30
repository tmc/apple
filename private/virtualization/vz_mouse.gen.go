// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMouse] class.
var (
	_VZMouseClass     VZMouseClass
	_VZMouseClassOnce sync.Once
)

func getVZMouseClass() VZMouseClass {
	_VZMouseClassOnce.Do(func() {
		_VZMouseClass = VZMouseClass{class: objc.GetClass("_VZMouse")}
	})
	return _VZMouseClass
}

// GetVZMouseClass returns the class object for _VZMouse.
func GetVZMouseClass() VZMouseClass {
	return getVZMouseClass()
}

type VZMouseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMouseClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMouseClass) Alloc() VZMouse {
	rv := objc.Send[VZMouse](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMouse.SendMouseEvents]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMouse
type VZMouse struct {
	VZPointingDevice
}

// VZMouseFromID constructs a [VZMouse] from an objc.ID.
func VZMouseFromID(id objc.ID) VZMouse {
	return VZMouse{VZPointingDevice: VZPointingDeviceFromID(id)}
}

// Ensure VZMouse implements IVZMouse.
var _ IVZMouse = VZMouse{}

// An interface definition for the [VZMouse] class.
//
// # Methods
//
//   - [IVZMouse.SendMouseEvents]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMouse
type IVZMouse interface {
	IVZPointingDevice

	// Topic: Methods

	SendMouseEvents(events objectivec.IObject)
}

// Init initializes the instance.
func (v VZMouse) Init() VZMouse {
	rv := objc.Send[VZMouse](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMouse) Autorelease() VZMouse {
	rv := objc.Send[VZMouse](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMouse creates a new VZMouse instance.
func NewVZMouse() VZMouse {
	class := getVZMouseClass()
	rv := objc.Send[VZMouse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/initWithType:virtualMachine:pointingDeviceIndex:
func NewVZMouseWithTypeVirtualMachinePointingDeviceIndex(type_ int64, machine objectivec.IObject, index uint64) VZMouse {
	instance := getVZMouseClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:virtualMachine:pointingDeviceIndex:"), type_, machine, index)
	return VZMouseFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMouse/sendMouseEvents:
func (v VZMouse) SendMouseEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMouseEvents:"), events)
}
