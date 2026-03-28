// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioEntropyDeviceConfiguration] class.
var (
	_VZVirtioEntropyDeviceConfigurationClass     VZVirtioEntropyDeviceConfigurationClass
	_VZVirtioEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioEntropyDeviceConfigurationClass() VZVirtioEntropyDeviceConfigurationClass {
	_VZVirtioEntropyDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioEntropyDeviceConfigurationClass = VZVirtioEntropyDeviceConfigurationClass{class: objc.GetClass("VZVirtioEntropyDeviceConfiguration")}
	})
	return _VZVirtioEntropyDeviceConfigurationClass
}

// GetVZVirtioEntropyDeviceConfigurationClass returns the class object for VZVirtioEntropyDeviceConfiguration.
func GetVZVirtioEntropyDeviceConfigurationClass() VZVirtioEntropyDeviceConfigurationClass {
	return getVZVirtioEntropyDeviceConfigurationClass()
}

type VZVirtioEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioEntropyDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioEntropyDeviceConfigurationClass) Alloc() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioEntropyDeviceConfiguration
type VZVirtioEntropyDeviceConfiguration struct {
	VZEntropyDeviceConfiguration
}

// VZVirtioEntropyDeviceConfigurationFromID constructs a [VZVirtioEntropyDeviceConfiguration] from an objc.ID.
func VZVirtioEntropyDeviceConfigurationFromID(id objc.ID) VZVirtioEntropyDeviceConfiguration {
	return VZVirtioEntropyDeviceConfiguration{VZEntropyDeviceConfiguration: VZEntropyDeviceConfigurationFromID(id)}
}
// Ensure VZVirtioEntropyDeviceConfiguration implements IVZVirtioEntropyDeviceConfiguration.
var _ IVZVirtioEntropyDeviceConfiguration = VZVirtioEntropyDeviceConfiguration{}

// An interface definition for the [VZVirtioEntropyDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioEntropyDeviceConfiguration
type IVZVirtioEntropyDeviceConfiguration interface {
	IVZEntropyDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioEntropyDeviceConfiguration) Init() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioEntropyDeviceConfiguration) Autorelease() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioEntropyDeviceConfiguration creates a new VZVirtioEntropyDeviceConfiguration instance.
func NewVZVirtioEntropyDeviceConfiguration() VZVirtioEntropyDeviceConfiguration {
	class := getVZVirtioEntropyDeviceConfigurationClass()
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

