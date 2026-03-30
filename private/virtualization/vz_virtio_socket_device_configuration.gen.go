// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSocketDeviceConfiguration] class.
var (
	_VZVirtioSocketDeviceConfigurationClass     VZVirtioSocketDeviceConfigurationClass
	_VZVirtioSocketDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioSocketDeviceConfigurationClass() VZVirtioSocketDeviceConfigurationClass {
	_VZVirtioSocketDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioSocketDeviceConfigurationClass = VZVirtioSocketDeviceConfigurationClass{class: objc.GetClass("VZVirtioSocketDeviceConfiguration")}
	})
	return _VZVirtioSocketDeviceConfigurationClass
}

// GetVZVirtioSocketDeviceConfigurationClass returns the class object for VZVirtioSocketDeviceConfiguration.
func GetVZVirtioSocketDeviceConfigurationClass() VZVirtioSocketDeviceConfigurationClass {
	return getVZVirtioSocketDeviceConfigurationClass()
}

type VZVirtioSocketDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSocketDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketDeviceConfigurationClass) Alloc() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDeviceConfiguration
type VZVirtioSocketDeviceConfiguration struct {
	VZSocketDeviceConfiguration
}

// VZVirtioSocketDeviceConfigurationFromID constructs a [VZVirtioSocketDeviceConfiguration] from an objc.ID.
func VZVirtioSocketDeviceConfigurationFromID(id objc.ID) VZVirtioSocketDeviceConfiguration {
	return VZVirtioSocketDeviceConfiguration{VZSocketDeviceConfiguration: VZSocketDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioSocketDeviceConfiguration implements IVZVirtioSocketDeviceConfiguration.
var _ IVZVirtioSocketDeviceConfiguration = VZVirtioSocketDeviceConfiguration{}

// An interface definition for the [VZVirtioSocketDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDeviceConfiguration
type IVZVirtioSocketDeviceConfiguration interface {
	IVZSocketDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioSocketDeviceConfiguration) Init() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketDeviceConfiguration) Autorelease() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDeviceConfiguration creates a new VZVirtioSocketDeviceConfiguration instance.
func NewVZVirtioSocketDeviceConfiguration() VZVirtioSocketDeviceConfiguration {
	class := getVZVirtioSocketDeviceConfigurationClass()
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
