// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacWallPowerSourceDeviceConfiguration] class.
var (
	_VZMacWallPowerSourceDeviceConfigurationClass     VZMacWallPowerSourceDeviceConfigurationClass
	_VZMacWallPowerSourceDeviceConfigurationClassOnce sync.Once
)

func getVZMacWallPowerSourceDeviceConfigurationClass() VZMacWallPowerSourceDeviceConfigurationClass {
	_VZMacWallPowerSourceDeviceConfigurationClassOnce.Do(func() {
		_VZMacWallPowerSourceDeviceConfigurationClass = VZMacWallPowerSourceDeviceConfigurationClass{class: objc.GetClass("_VZMacWallPowerSourceDeviceConfiguration")}
	})
	return _VZMacWallPowerSourceDeviceConfigurationClass
}

// GetVZMacWallPowerSourceDeviceConfigurationClass returns the class object for _VZMacWallPowerSourceDeviceConfiguration.
func GetVZMacWallPowerSourceDeviceConfigurationClass() VZMacWallPowerSourceDeviceConfigurationClass {
	return getVZMacWallPowerSourceDeviceConfigurationClass()
}

type VZMacWallPowerSourceDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacWallPowerSourceDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacWallPowerSourceDeviceConfigurationClass) Alloc() VZMacWallPowerSourceDeviceConfiguration {
	rv := objc.Send[VZMacWallPowerSourceDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacWallPowerSourceDeviceConfiguration
type VZMacWallPowerSourceDeviceConfiguration struct {
	VZPowerSourceDeviceConfiguration
}

// VZMacWallPowerSourceDeviceConfigurationFromID constructs a [VZMacWallPowerSourceDeviceConfiguration] from an objc.ID.
func VZMacWallPowerSourceDeviceConfigurationFromID(id objc.ID) VZMacWallPowerSourceDeviceConfiguration {
	return VZMacWallPowerSourceDeviceConfiguration{VZPowerSourceDeviceConfiguration: VZPowerSourceDeviceConfigurationFromID(id)}
}

// Ensure VZMacWallPowerSourceDeviceConfiguration implements IVZMacWallPowerSourceDeviceConfiguration.
var _ IVZMacWallPowerSourceDeviceConfiguration = VZMacWallPowerSourceDeviceConfiguration{}

// An interface definition for the [VZMacWallPowerSourceDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacWallPowerSourceDeviceConfiguration
type IVZMacWallPowerSourceDeviceConfiguration interface {
	IVZPowerSourceDeviceConfiguration
}

// Init initializes the instance.
func (v VZMacWallPowerSourceDeviceConfiguration) Init() VZMacWallPowerSourceDeviceConfiguration {
	rv := objc.Send[VZMacWallPowerSourceDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacWallPowerSourceDeviceConfiguration) Autorelease() VZMacWallPowerSourceDeviceConfiguration {
	rv := objc.Send[VZMacWallPowerSourceDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacWallPowerSourceDeviceConfiguration creates a new VZMacWallPowerSourceDeviceConfiguration instance.
func NewVZMacWallPowerSourceDeviceConfiguration() VZMacWallPowerSourceDeviceConfiguration {
	class := getVZMacWallPowerSourceDeviceConfigurationClass()
	rv := objc.Send[VZMacWallPowerSourceDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
