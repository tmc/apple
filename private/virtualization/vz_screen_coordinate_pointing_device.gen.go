// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZScreenCoordinatePointingDevice] class.
var (
	_VZScreenCoordinatePointingDeviceClass     VZScreenCoordinatePointingDeviceClass
	_VZScreenCoordinatePointingDeviceClassOnce sync.Once
)

func getVZScreenCoordinatePointingDeviceClass() VZScreenCoordinatePointingDeviceClass {
	_VZScreenCoordinatePointingDeviceClassOnce.Do(func() {
		_VZScreenCoordinatePointingDeviceClass = VZScreenCoordinatePointingDeviceClass{class: objc.GetClass("_VZScreenCoordinatePointingDevice")}
	})
	return _VZScreenCoordinatePointingDeviceClass
}

// GetVZScreenCoordinatePointingDeviceClass returns the class object for _VZScreenCoordinatePointingDevice.
func GetVZScreenCoordinatePointingDeviceClass() VZScreenCoordinatePointingDeviceClass {
	return getVZScreenCoordinatePointingDeviceClass()
}

type VZScreenCoordinatePointingDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZScreenCoordinatePointingDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZScreenCoordinatePointingDeviceClass) Alloc() VZScreenCoordinatePointingDevice {
	rv := objc.Send[VZScreenCoordinatePointingDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZScreenCoordinatePointingDevice.SendPointerEvents]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointingDevice
type VZScreenCoordinatePointingDevice struct {
	VZPointingDevice
}

// VZScreenCoordinatePointingDeviceFromID constructs a [VZScreenCoordinatePointingDevice] from an objc.ID.
func VZScreenCoordinatePointingDeviceFromID(id objc.ID) VZScreenCoordinatePointingDevice {
	return VZScreenCoordinatePointingDevice{VZPointingDevice: VZPointingDeviceFromID(id)}
}

// Ensure VZScreenCoordinatePointingDevice implements IVZScreenCoordinatePointingDevice.
var _ IVZScreenCoordinatePointingDevice = VZScreenCoordinatePointingDevice{}

// An interface definition for the [VZScreenCoordinatePointingDevice] class.
//
// # Methods
//
//   - [IVZScreenCoordinatePointingDevice.SendPointerEvents]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointingDevice
type IVZScreenCoordinatePointingDevice interface {
	IVZPointingDevice

	// Topic: Methods

	SendPointerEvents(events objectivec.IObject)
}

// Init initializes the instance.
func (v VZScreenCoordinatePointingDevice) Init() VZScreenCoordinatePointingDevice {
	rv := objc.Send[VZScreenCoordinatePointingDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZScreenCoordinatePointingDevice) Autorelease() VZScreenCoordinatePointingDevice {
	rv := objc.Send[VZScreenCoordinatePointingDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZScreenCoordinatePointingDevice creates a new VZScreenCoordinatePointingDevice instance.
func NewVZScreenCoordinatePointingDevice() VZScreenCoordinatePointingDevice {
	class := getVZScreenCoordinatePointingDeviceClass()
	rv := objc.Send[VZScreenCoordinatePointingDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingDevice/initWithType:virtualMachine:pointingDeviceIndex:
func NewVZScreenCoordinatePointingDeviceWithTypeVirtualMachinePointingDeviceIndex(type_ int64, machine objectivec.IObject, index uint64) VZScreenCoordinatePointingDevice {
	instance := getVZScreenCoordinatePointingDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:virtualMachine:pointingDeviceIndex:"), type_, machine, index)
	return VZScreenCoordinatePointingDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointingDevice/sendPointerEvents:
func (v VZScreenCoordinatePointingDevice) SendPointerEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendPointerEvents:"), events)
}
