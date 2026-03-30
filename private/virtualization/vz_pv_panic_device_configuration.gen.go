// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZPvPanicDeviceConfiguration] class.
var (
	_VZPvPanicDeviceConfigurationClass     VZPvPanicDeviceConfigurationClass
	_VZPvPanicDeviceConfigurationClassOnce sync.Once
)

func getVZPvPanicDeviceConfigurationClass() VZPvPanicDeviceConfigurationClass {
	_VZPvPanicDeviceConfigurationClassOnce.Do(func() {
		_VZPvPanicDeviceConfigurationClass = VZPvPanicDeviceConfigurationClass{class: objc.GetClass("_VZPvPanicDeviceConfiguration")}
	})
	return _VZPvPanicDeviceConfigurationClass
}

// GetVZPvPanicDeviceConfigurationClass returns the class object for _VZPvPanicDeviceConfiguration.
func GetVZPvPanicDeviceConfigurationClass() VZPvPanicDeviceConfigurationClass {
	return getVZPvPanicDeviceConfigurationClass()
}

type VZPvPanicDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPvPanicDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPvPanicDeviceConfigurationClass) Alloc() VZPvPanicDeviceConfiguration {
	rv := objc.Send[VZPvPanicDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPvPanicDeviceConfiguration
type VZPvPanicDeviceConfiguration struct {
	VZPanicDeviceConfiguration
}

// VZPvPanicDeviceConfigurationFromID constructs a [VZPvPanicDeviceConfiguration] from an objc.ID.
func VZPvPanicDeviceConfigurationFromID(id objc.ID) VZPvPanicDeviceConfiguration {
	return VZPvPanicDeviceConfiguration{VZPanicDeviceConfiguration: VZPanicDeviceConfigurationFromID(id)}
}

// Ensure VZPvPanicDeviceConfiguration implements IVZPvPanicDeviceConfiguration.
var _ IVZPvPanicDeviceConfiguration = VZPvPanicDeviceConfiguration{}

// An interface definition for the [VZPvPanicDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPvPanicDeviceConfiguration
type IVZPvPanicDeviceConfiguration interface {
	IVZPanicDeviceConfiguration
}

// Init initializes the instance.
func (v VZPvPanicDeviceConfiguration) Init() VZPvPanicDeviceConfiguration {
	rv := objc.Send[VZPvPanicDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPvPanicDeviceConfiguration) Autorelease() VZPvPanicDeviceConfiguration {
	rv := objc.Send[VZPvPanicDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPvPanicDeviceConfiguration creates a new VZPvPanicDeviceConfiguration instance.
func NewVZPvPanicDeviceConfiguration() VZPvPanicDeviceConfiguration {
	class := getVZPvPanicDeviceConfigurationClass()
	rv := objc.Send[VZPvPanicDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
