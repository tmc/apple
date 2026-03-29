// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZHIDDevice] class.
var (
	_VZHIDDeviceClass     VZHIDDeviceClass
	_VZHIDDeviceClassOnce sync.Once
)

func getVZHIDDeviceClass() VZHIDDeviceClass {
	_VZHIDDeviceClassOnce.Do(func() {
		_VZHIDDeviceClass = VZHIDDeviceClass{class: objc.GetClass("_VZHIDDevice")}
	})
	return _VZHIDDeviceClass
}

// GetVZHIDDeviceClass returns the class object for _VZHIDDevice.
func GetVZHIDDeviceClass() VZHIDDeviceClass {
	return getVZHIDDeviceClass()
}

type VZHIDDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZHIDDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHIDDeviceClass) Alloc() VZHIDDevice {
	rv := objc.Send[VZHIDDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZHIDDevice.SendIOHIDEvents]
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDevice
type VZHIDDevice struct {
	objectivec.Object
}

// VZHIDDeviceFromID constructs a [VZHIDDevice] from an objc.ID.
func VZHIDDeviceFromID(id objc.ID) VZHIDDevice {
	return VZHIDDevice{objectivec.Object{ID: id}}
}
// Ensure VZHIDDevice implements IVZHIDDevice.
var _ IVZHIDDevice = VZHIDDevice{}

// An interface definition for the [VZHIDDevice] class.
//
// # Methods
//
//   - [IVZHIDDevice.SendIOHIDEvents]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDevice
type IVZHIDDevice interface {
	objectivec.IObject

	// Topic: Methods

	SendIOHIDEvents(iOHIDEvents objectivec.IObject)
}

// Init initializes the instance.
func (v VZHIDDevice) Init() VZHIDDevice {
	rv := objc.Send[VZHIDDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHIDDevice) Autorelease() VZHIDDevice {
	rv := objc.Send[VZHIDDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHIDDevice creates a new VZHIDDevice instance.
func NewVZHIDDevice() VZHIDDevice {
	class := getVZHIDDeviceClass()
	rv := objc.Send[VZHIDDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDevice/sendIOHIDEvents:
func (v VZHIDDevice) SendIOHIDEvents(iOHIDEvents objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendIOHIDEvents:"), iOHIDEvents)
}

