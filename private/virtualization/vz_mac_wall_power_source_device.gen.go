// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacWallPowerSourceDevice] class.
var (
	_VZMacWallPowerSourceDeviceClass     VZMacWallPowerSourceDeviceClass
	_VZMacWallPowerSourceDeviceClassOnce sync.Once
)

func getVZMacWallPowerSourceDeviceClass() VZMacWallPowerSourceDeviceClass {
	_VZMacWallPowerSourceDeviceClassOnce.Do(func() {
		_VZMacWallPowerSourceDeviceClass = VZMacWallPowerSourceDeviceClass{class: objc.GetClass("_VZMacWallPowerSourceDevice")}
	})
	return _VZMacWallPowerSourceDeviceClass
}

// GetVZMacWallPowerSourceDeviceClass returns the class object for _VZMacWallPowerSourceDevice.
func GetVZMacWallPowerSourceDeviceClass() VZMacWallPowerSourceDeviceClass {
	return getVZMacWallPowerSourceDeviceClass()
}

type VZMacWallPowerSourceDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacWallPowerSourceDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacWallPowerSourceDeviceClass) Alloc() VZMacWallPowerSourceDevice {
	rv := objc.Send[VZMacWallPowerSourceDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacWallPowerSourceDevice
type VZMacWallPowerSourceDevice struct {
	VZPowerSourceDevice
}

// VZMacWallPowerSourceDeviceFromID constructs a [VZMacWallPowerSourceDevice] from an objc.ID.
func VZMacWallPowerSourceDeviceFromID(id objc.ID) VZMacWallPowerSourceDevice {
	return VZMacWallPowerSourceDevice{VZPowerSourceDevice: VZPowerSourceDeviceFromID(id)}
}

// Ensure VZMacWallPowerSourceDevice implements IVZMacWallPowerSourceDevice.
var _ IVZMacWallPowerSourceDevice = VZMacWallPowerSourceDevice{}

// An interface definition for the [VZMacWallPowerSourceDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacWallPowerSourceDevice
type IVZMacWallPowerSourceDevice interface {
	IVZPowerSourceDevice
}

// Init initializes the instance.
func (v VZMacWallPowerSourceDevice) Init() VZMacWallPowerSourceDevice {
	rv := objc.Send[VZMacWallPowerSourceDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacWallPowerSourceDevice) Autorelease() VZMacWallPowerSourceDevice {
	rv := objc.Send[VZMacWallPowerSourceDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacWallPowerSourceDevice creates a new VZMacWallPowerSourceDevice instance.
func NewVZMacWallPowerSourceDevice() VZMacWallPowerSourceDevice {
	class := getVZMacWallPowerSourceDeviceClass()
	rv := objc.Send[VZMacWallPowerSourceDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}
