// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSocketDevice] class.
var (
	_VZSocketDeviceClass     VZSocketDeviceClass
	_VZSocketDeviceClassOnce sync.Once
)

func getVZSocketDeviceClass() VZSocketDeviceClass {
	_VZSocketDeviceClassOnce.Do(func() {
		_VZSocketDeviceClass = VZSocketDeviceClass{class: objc.GetClass("VZSocketDevice")}
	})
	return _VZSocketDeviceClass
}

// GetVZSocketDeviceClass returns the class object for VZSocketDevice.
func GetVZSocketDeviceClass() VZSocketDeviceClass {
	return getVZSocketDeviceClass()
}

type VZSocketDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSocketDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSocketDeviceClass) Alloc() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZSocketDevice._init]
// See: https://developer.apple.com/documentation/Virtualization/VZSocketDevice
type VZSocketDevice struct {
	objectivec.Object
}

// VZSocketDeviceFromID constructs a [VZSocketDevice] from an objc.ID.
func VZSocketDeviceFromID(id objc.ID) VZSocketDevice {
	return VZSocketDevice{objectivec.Object{ID: id}}
}
// Ensure VZSocketDevice implements IVZSocketDevice.
var _ IVZSocketDevice = VZSocketDevice{}

// An interface definition for the [VZSocketDevice] class.
//
// # Methods
//
//   - [IVZSocketDevice._init]
//
// See: https://developer.apple.com/documentation/Virtualization/VZSocketDevice
type IVZSocketDevice interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (s VZSocketDevice) Init() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSocketDevice) Autorelease() VZSocketDevice {
	rv := objc.Send[VZSocketDevice](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDevice creates a new VZSocketDevice instance.
func NewVZSocketDevice() VZSocketDevice {
	class := getVZSocketDeviceClass()
	rv := objc.Send[VZSocketDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZSocketDevice/_init
func (s VZSocketDevice) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

