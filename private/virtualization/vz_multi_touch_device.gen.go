// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMultiTouchDevice] class.
var (
	_VZMultiTouchDeviceClass     VZMultiTouchDeviceClass
	_VZMultiTouchDeviceClassOnce sync.Once
)

func getVZMultiTouchDeviceClass() VZMultiTouchDeviceClass {
	_VZMultiTouchDeviceClassOnce.Do(func() {
		_VZMultiTouchDeviceClass = VZMultiTouchDeviceClass{class: objc.GetClass("_VZMultiTouchDevice")}
	})
	return _VZMultiTouchDeviceClass
}

// GetVZMultiTouchDeviceClass returns the class object for _VZMultiTouchDevice.
func GetVZMultiTouchDeviceClass() VZMultiTouchDeviceClass {
	return getVZMultiTouchDeviceClass()
}

type VZMultiTouchDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMultiTouchDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMultiTouchDeviceClass) Alloc() VZMultiTouchDevice {
	rv := objc.Send[VZMultiTouchDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMultiTouchDevice.AssociationIdentifier]
//   - [VZMultiTouchDevice.SendMultiTouchEvents]
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchDevice
type VZMultiTouchDevice struct {
	objectivec.Object
}

// VZMultiTouchDeviceFromID constructs a [VZMultiTouchDevice] from an objc.ID.
func VZMultiTouchDeviceFromID(id objc.ID) VZMultiTouchDevice {
	return VZMultiTouchDevice{objectivec.Object{ID: id}}
}
// Ensure VZMultiTouchDevice implements IVZMultiTouchDevice.
var _ IVZMultiTouchDevice = VZMultiTouchDevice{}

// An interface definition for the [VZMultiTouchDevice] class.
//
// # Methods
//
//   - [IVZMultiTouchDevice.AssociationIdentifier]
//   - [IVZMultiTouchDevice.SendMultiTouchEvents]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchDevice
type IVZMultiTouchDevice interface {
	objectivec.IObject

	// Topic: Methods

	AssociationIdentifier() foundation.NSUUID
	SendMultiTouchEvents(events objectivec.IObject)
}

// Init initializes the instance.
func (v VZMultiTouchDevice) Init() VZMultiTouchDevice {
	rv := objc.Send[VZMultiTouchDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMultiTouchDevice) Autorelease() VZMultiTouchDevice {
	rv := objc.Send[VZMultiTouchDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMultiTouchDevice creates a new VZMultiTouchDevice instance.
func NewVZMultiTouchDevice() VZMultiTouchDevice {
	class := getVZMultiTouchDeviceClass()
	rv := objc.Send[VZMultiTouchDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchDevice/sendMultiTouchEvents:
func (v VZMultiTouchDevice) SendMultiTouchEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMultiTouchEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchDevice/associationIdentifier
func (v VZMultiTouchDevice) AssociationIdentifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("associationIdentifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

