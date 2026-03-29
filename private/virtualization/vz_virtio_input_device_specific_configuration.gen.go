// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioInputDeviceSpecificConfiguration] class.
var (
	_VZVirtioInputDeviceSpecificConfigurationClass     VZVirtioInputDeviceSpecificConfigurationClass
	_VZVirtioInputDeviceSpecificConfigurationClassOnce sync.Once
)

func getVZVirtioInputDeviceSpecificConfigurationClass() VZVirtioInputDeviceSpecificConfigurationClass {
	_VZVirtioInputDeviceSpecificConfigurationClassOnce.Do(func() {
		_VZVirtioInputDeviceSpecificConfigurationClass = VZVirtioInputDeviceSpecificConfigurationClass{class: objc.GetClass("_VZVirtioInputDeviceSpecificConfiguration")}
	})
	return _VZVirtioInputDeviceSpecificConfigurationClass
}

// GetVZVirtioInputDeviceSpecificConfigurationClass returns the class object for _VZVirtioInputDeviceSpecificConfiguration.
func GetVZVirtioInputDeviceSpecificConfigurationClass() VZVirtioInputDeviceSpecificConfigurationClass {
	return getVZVirtioInputDeviceSpecificConfigurationClass()
}

type VZVirtioInputDeviceSpecificConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioInputDeviceSpecificConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioInputDeviceSpecificConfigurationClass) Alloc() VZVirtioInputDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioInputDeviceSpecificConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtioInputDeviceSpecificConfiguration.Configurations]
//   - [VZVirtioInputDeviceSpecificConfiguration.InitWithConfigurations]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioInputDeviceSpecificConfiguration
type VZVirtioInputDeviceSpecificConfiguration struct {
	VZVirtioDeviceSpecificConfiguration
}

// VZVirtioInputDeviceSpecificConfigurationFromID constructs a [VZVirtioInputDeviceSpecificConfiguration] from an objc.ID.
func VZVirtioInputDeviceSpecificConfigurationFromID(id objc.ID) VZVirtioInputDeviceSpecificConfiguration {
	return VZVirtioInputDeviceSpecificConfiguration{VZVirtioDeviceSpecificConfiguration: VZVirtioDeviceSpecificConfigurationFromID(id)}
}
// Ensure VZVirtioInputDeviceSpecificConfiguration implements IVZVirtioInputDeviceSpecificConfiguration.
var _ IVZVirtioInputDeviceSpecificConfiguration = VZVirtioInputDeviceSpecificConfiguration{}

// An interface definition for the [VZVirtioInputDeviceSpecificConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioInputDeviceSpecificConfiguration.Configurations]
//   - [IVZVirtioInputDeviceSpecificConfiguration.InitWithConfigurations]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioInputDeviceSpecificConfiguration
type IVZVirtioInputDeviceSpecificConfiguration interface {
	IVZVirtioDeviceSpecificConfiguration

	// Topic: Methods

	Configurations() foundation.INSArray
	InitWithConfigurations(configurations objectivec.IObject) VZVirtioInputDeviceSpecificConfiguration
}

// Init initializes the instance.
func (v VZVirtioInputDeviceSpecificConfiguration) Init() VZVirtioInputDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioInputDeviceSpecificConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioInputDeviceSpecificConfiguration) Autorelease() VZVirtioInputDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioInputDeviceSpecificConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioInputDeviceSpecificConfiguration creates a new VZVirtioInputDeviceSpecificConfiguration instance.
func NewVZVirtioInputDeviceSpecificConfiguration() VZVirtioInputDeviceSpecificConfiguration {
	class := getVZVirtioInputDeviceSpecificConfigurationClass()
	rv := objc.Send[VZVirtioInputDeviceSpecificConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioInputDeviceSpecificConfiguration/initWithConfigurations:
func NewVZVirtioInputDeviceSpecificConfigurationWithConfigurations(configurations objectivec.IObject) VZVirtioInputDeviceSpecificConfiguration {
	instance := getVZVirtioInputDeviceSpecificConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfigurations:"), configurations)
	return VZVirtioInputDeviceSpecificConfigurationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioInputDeviceSpecificConfiguration/initWithConfigurations:
func (v VZVirtioInputDeviceSpecificConfiguration) InitWithConfigurations(configurations objectivec.IObject) VZVirtioInputDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioInputDeviceSpecificConfiguration](v.ID, objc.Sel("initWithConfigurations:"), configurations)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioInputDeviceSpecificConfiguration/configurations
func (v VZVirtioInputDeviceSpecificConfiguration) Configurations() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configurations"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

