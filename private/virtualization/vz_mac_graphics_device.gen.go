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

//
// # Methods
//
//   - [VZMacGraphicsDevice._deviceFeatureLevel]
//   - [VZMacGraphicsDevice._prefersLowPower]
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice
type VZMacGraphicsDevice struct {
	VZGraphicsDevice
}

// VZMacGraphicsDeviceFromID constructs a [VZMacGraphicsDevice] from an objc.ID.
func VZMacGraphicsDeviceFromID(id objc.ID) VZMacGraphicsDevice {
	return VZMacGraphicsDevice{VZGraphicsDevice: VZGraphicsDeviceFromID(id)}
}
// Ensure VZMacGraphicsDevice implements IVZMacGraphicsDevice.
var _ IVZMacGraphicsDevice = VZMacGraphicsDevice{}

// An interface definition for the [VZMacGraphicsDevice] class.
//
// # Methods
//
//   - [IVZMacGraphicsDevice._deviceFeatureLevel]
//   - [IVZMacGraphicsDevice._prefersLowPower]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice
type IVZMacGraphicsDevice interface {
	IVZGraphicsDevice

	// Topic: Methods

	_deviceFeatureLevel() int64
	_prefersLowPower() bool
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

// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice/_deviceFeatureLevel
func (m VZMacGraphicsDevice) _deviceFeatureLevel() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("_deviceFeatureLevel"))
	return rv
}

// DeviceFeatureLevel is an exported wrapper for the private method _deviceFeatureLevel.
func (m VZMacGraphicsDevice) DeviceFeatureLevel() int64 {
	return m._deviceFeatureLevel()
}
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDevice/_prefersLowPower
func (m VZMacGraphicsDevice) _prefersLowPower() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_prefersLowPower"))
	return rv
}

// PrefersLowPower is an exported wrapper for the private method _prefersLowPower.
func (m VZMacGraphicsDevice) PrefersLowPower() bool {
	return m._prefersLowPower()
}

