// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacGraphicsDevice] class.
var (
	_VZMacGraphicsDeviceClass     VZMacGraphicsDeviceClass
	_VZMacGraphicsDeviceClassOnce sync.Once
)

func getVZMacGraphicsDeviceClass() VZMacGraphicsDeviceClass {
	_VZMacGraphicsDeviceClassOnce.Do(func() {
		_VZMacGraphicsDeviceClass = VZMacGraphicsDeviceClass{class: objc.GetClass("VZMacGraphicsDevice")}
	})
	return _VZMacGraphicsDeviceClass
}

// GetVZMacGraphicsDeviceClass returns the class object for VZMacGraphicsDevice.
func GetVZMacGraphicsDeviceClass() VZMacGraphicsDeviceClass {
	return getVZMacGraphicsDeviceClass()
}

type VZMacGraphicsDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacGraphicsDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDeviceClass) Alloc() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a Mac graphics device.
//
// # Overview
//
// You don’t instantiate a [VZMacGraphicsDevice] directly. Graphics devices
// are first configured on the [VZVirtualMachineConfiguration] through a
// subclass of [VZMacGraphicsDeviceConfiguration]. When the framework creates
// a VZVirtualMachine from the configuration, the graphics devices are
// available through the `graphicsDevices` property.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice
type VZMacGraphicsDevice struct {
	VZGraphicsDevice
}

// VZMacGraphicsDeviceFromID constructs a [VZMacGraphicsDevice] from an objc.ID.
//
// An object that represents a Mac graphics device.
func VZMacGraphicsDeviceFromID(id objc.ID) VZMacGraphicsDevice {
	return VZMacGraphicsDevice{VZGraphicsDevice: VZGraphicsDeviceFromID(id)}
}

// NOTE: VZMacGraphicsDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacGraphicsDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice
type IVZMacGraphicsDevice interface {
	IVZGraphicsDevice
}

// Init initializes the instance.
func (m VZMacGraphicsDevice) Init() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDevice) Autorelease() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDevice creates a new VZMacGraphicsDevice instance.
func NewVZMacGraphicsDevice() VZMacGraphicsDevice {
	class := getVZMacGraphicsDeviceClass()
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}
