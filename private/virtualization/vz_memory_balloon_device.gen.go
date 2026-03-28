// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMemoryBalloonDevice] class.
var (
	_VZMemoryBalloonDeviceClass     VZMemoryBalloonDeviceClass
	_VZMemoryBalloonDeviceClassOnce sync.Once
)

func getVZMemoryBalloonDeviceClass() VZMemoryBalloonDeviceClass {
	_VZMemoryBalloonDeviceClassOnce.Do(func() {
		_VZMemoryBalloonDeviceClass = VZMemoryBalloonDeviceClass{class: objc.GetClass("VZMemoryBalloonDevice")}
	})
	return _VZMemoryBalloonDeviceClass
}

// GetVZMemoryBalloonDeviceClass returns the class object for VZMemoryBalloonDevice.
func GetVZMemoryBalloonDeviceClass() VZMemoryBalloonDeviceClass {
	return getVZMemoryBalloonDeviceClass()
}

type VZMemoryBalloonDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMemoryBalloonDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMemoryBalloonDeviceClass) Alloc() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMemoryBalloonDevice._init]
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice
type VZMemoryBalloonDevice struct {
	objectivec.Object
}

// VZMemoryBalloonDeviceFromID constructs a [VZMemoryBalloonDevice] from an objc.ID.
func VZMemoryBalloonDeviceFromID(id objc.ID) VZMemoryBalloonDevice {
	return VZMemoryBalloonDevice{objectivec.Object{ID: id}}
}
// Ensure VZMemoryBalloonDevice implements IVZMemoryBalloonDevice.
var _ IVZMemoryBalloonDevice = VZMemoryBalloonDevice{}

// An interface definition for the [VZMemoryBalloonDevice] class.
//
// # Methods
//
//   - [IVZMemoryBalloonDevice._init]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice
type IVZMemoryBalloonDevice interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (m VZMemoryBalloonDevice) Init() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMemoryBalloonDevice) Autorelease() VZMemoryBalloonDevice {
	rv := objc.Send[VZMemoryBalloonDevice](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMemoryBalloonDevice creates a new VZMemoryBalloonDevice instance.
func NewVZMemoryBalloonDevice() VZMemoryBalloonDevice {
	class := getVZMemoryBalloonDeviceClass()
	rv := objc.Send[VZMemoryBalloonDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZMemoryBalloonDevice/_init
func (m VZMemoryBalloonDevice) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

