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

// # Methods
//
//   - [VZMacGraphicsDevice._deviceFeatureLevel]
//   - [VZMacGraphicsDevice._prefersLowPower]
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
type IVZMacGraphicsDevice interface {
	IVZGraphicsDevice

	// Topic: Methods

	_deviceFeatureLevel() int64
	_prefersLowPower() bool
}

// Init initializes the instance.
func (v VZMacGraphicsDevice) Init() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacGraphicsDevice) Autorelease() VZMacGraphicsDevice {
	rv := objc.Send[VZMacGraphicsDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDevice creates a new VZMacGraphicsDevice instance.
func NewVZMacGraphicsDevice() VZMacGraphicsDevice {
	class := getVZMacGraphicsDeviceClass()
	rv := objc.Send[VZMacGraphicsDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacGraphicsDevice) _deviceFeatureLevel() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("_deviceFeatureLevel"))
	return rv
}

// DeviceFeatureLevel is an exported wrapper for the private method _deviceFeatureLevel.
func (v VZMacGraphicsDevice) DeviceFeatureLevel() (int64, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_deviceFeatureLevel")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_deviceFeatureLevel"}
		return 0, err
	}
	return v._deviceFeatureLevel(), nil
}

// CanDeviceFeatureLevel reports whether the receiver responds to the private selector _deviceFeatureLevel.
func (v VZMacGraphicsDevice) CanDeviceFeatureLevel() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_deviceFeatureLevel"))
}
func (v VZMacGraphicsDevice) _prefersLowPower() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_prefersLowPower"))
	return rv
}

// PrefersLowPower is an exported wrapper for the private method _prefersLowPower.
func (v VZMacGraphicsDevice) PrefersLowPower() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_prefersLowPower")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_prefersLowPower"}
		return false, err
	}
	return v._prefersLowPower(), nil
}

// CanPrefersLowPower reports whether the receiver responds to the private selector _prefersLowPower.
func (v VZMacGraphicsDevice) CanPrefersLowPower() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_prefersLowPower"))
}
