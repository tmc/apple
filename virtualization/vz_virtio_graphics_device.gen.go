// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioGraphicsDevice] class.
var (
	_VZVirtioGraphicsDeviceClass     VZVirtioGraphicsDeviceClass
	_VZVirtioGraphicsDeviceClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceClass() VZVirtioGraphicsDeviceClass {
	_VZVirtioGraphicsDeviceClassOnce.Do(func() {
		_VZVirtioGraphicsDeviceClass = VZVirtioGraphicsDeviceClass{class: objc.GetClass("VZVirtioGraphicsDevice")}
	})
	return _VZVirtioGraphicsDeviceClass
}

// GetVZVirtioGraphicsDeviceClass returns the class object for VZVirtioGraphicsDevice.
func GetVZVirtioGraphicsDeviceClass() VZVirtioGraphicsDeviceClass {
	return getVZVirtioGraphicsDeviceClass()
}

type VZVirtioGraphicsDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsDeviceClass) Alloc() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A Virtio graphics device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDevice
type VZVirtioGraphicsDevice struct {
	VZGraphicsDevice
}

// VZVirtioGraphicsDeviceFromID constructs a [VZVirtioGraphicsDevice] from an objc.ID.
//
// A Virtio graphics device.
func VZVirtioGraphicsDeviceFromID(id objc.ID) VZVirtioGraphicsDevice {
	return VZVirtioGraphicsDevice{VZGraphicsDevice: VZGraphicsDeviceFromID(id)}
}
// NOTE: VZVirtioGraphicsDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioGraphicsDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDevice
type IVZVirtioGraphicsDevice interface {
	IVZGraphicsDevice
}

// Init initializes the instance.
func (v VZVirtioGraphicsDevice) Init() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsDevice) Autorelease() VZVirtioGraphicsDevice {
	rv := objc.Send[VZVirtioGraphicsDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDevice creates a new VZVirtioGraphicsDevice instance.
func NewVZVirtioGraphicsDevice() VZVirtioGraphicsDevice {
	class := getVZVirtioGraphicsDeviceClass()
	rv := objc.Send[VZVirtioGraphicsDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

