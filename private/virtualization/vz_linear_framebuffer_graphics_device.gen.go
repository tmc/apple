// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZLinearFramebufferGraphicsDevice] class.
var (
	_VZLinearFramebufferGraphicsDeviceClass     VZLinearFramebufferGraphicsDeviceClass
	_VZLinearFramebufferGraphicsDeviceClassOnce sync.Once
)

func getVZLinearFramebufferGraphicsDeviceClass() VZLinearFramebufferGraphicsDeviceClass {
	_VZLinearFramebufferGraphicsDeviceClassOnce.Do(func() {
		_VZLinearFramebufferGraphicsDeviceClass = VZLinearFramebufferGraphicsDeviceClass{class: objc.GetClass("_VZLinearFramebufferGraphicsDevice")}
	})
	return _VZLinearFramebufferGraphicsDeviceClass
}

// GetVZLinearFramebufferGraphicsDeviceClass returns the class object for _VZLinearFramebufferGraphicsDevice.
func GetVZLinearFramebufferGraphicsDeviceClass() VZLinearFramebufferGraphicsDeviceClass {
	return getVZLinearFramebufferGraphicsDeviceClass()
}

type VZLinearFramebufferGraphicsDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZLinearFramebufferGraphicsDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZLinearFramebufferGraphicsDeviceClass) Alloc() VZLinearFramebufferGraphicsDevice {
	rv := objc.Send[VZLinearFramebufferGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDevice
type VZLinearFramebufferGraphicsDevice struct {
	VZGraphicsDevice
}

// VZLinearFramebufferGraphicsDeviceFromID constructs a [VZLinearFramebufferGraphicsDevice] from an objc.ID.
func VZLinearFramebufferGraphicsDeviceFromID(id objc.ID) VZLinearFramebufferGraphicsDevice {
	return VZLinearFramebufferGraphicsDevice{VZGraphicsDevice: VZGraphicsDeviceFromID(id)}
}

// Ensure VZLinearFramebufferGraphicsDevice implements IVZLinearFramebufferGraphicsDevice.
var _ IVZLinearFramebufferGraphicsDevice = VZLinearFramebufferGraphicsDevice{}

// An interface definition for the [VZLinearFramebufferGraphicsDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZLinearFramebufferGraphicsDevice
type IVZLinearFramebufferGraphicsDevice interface {
	IVZGraphicsDevice
}

// Init initializes the instance.
func (v VZLinearFramebufferGraphicsDevice) Init() VZLinearFramebufferGraphicsDevice {
	rv := objc.Send[VZLinearFramebufferGraphicsDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZLinearFramebufferGraphicsDevice) Autorelease() VZLinearFramebufferGraphicsDevice {
	rv := objc.Send[VZLinearFramebufferGraphicsDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZLinearFramebufferGraphicsDevice creates a new VZLinearFramebufferGraphicsDevice instance.
func NewVZLinearFramebufferGraphicsDevice() VZLinearFramebufferGraphicsDevice {
	class := getVZLinearFramebufferGraphicsDeviceClass()
	rv := objc.Send[VZLinearFramebufferGraphicsDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}
