// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioGenericDeviceSpecificConfiguration] class.
var (
	_VZVirtioGenericDeviceSpecificConfigurationClass     VZVirtioGenericDeviceSpecificConfigurationClass
	_VZVirtioGenericDeviceSpecificConfigurationClassOnce sync.Once
)

func getVZVirtioGenericDeviceSpecificConfigurationClass() VZVirtioGenericDeviceSpecificConfigurationClass {
	_VZVirtioGenericDeviceSpecificConfigurationClassOnce.Do(func() {
		_VZVirtioGenericDeviceSpecificConfigurationClass = VZVirtioGenericDeviceSpecificConfigurationClass{class: objc.GetClass("_VZVirtioGenericDeviceSpecificConfiguration")}
	})
	return _VZVirtioGenericDeviceSpecificConfigurationClass
}

// GetVZVirtioGenericDeviceSpecificConfigurationClass returns the class object for _VZVirtioGenericDeviceSpecificConfiguration.
func GetVZVirtioGenericDeviceSpecificConfigurationClass() VZVirtioGenericDeviceSpecificConfigurationClass {
	return getVZVirtioGenericDeviceSpecificConfigurationClass()
}

type VZVirtioGenericDeviceSpecificConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioGenericDeviceSpecificConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGenericDeviceSpecificConfigurationClass) Alloc() VZVirtioGenericDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioGenericDeviceSpecificConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtioGenericDeviceSpecificConfiguration.ConfigurationData]
//   - [VZVirtioGenericDeviceSpecificConfiguration.InitWithConfigurationData]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioGenericDeviceSpecificConfiguration
type VZVirtioGenericDeviceSpecificConfiguration struct {
	VZVirtioDeviceSpecificConfiguration
}

// VZVirtioGenericDeviceSpecificConfigurationFromID constructs a [VZVirtioGenericDeviceSpecificConfiguration] from an objc.ID.
func VZVirtioGenericDeviceSpecificConfigurationFromID(id objc.ID) VZVirtioGenericDeviceSpecificConfiguration {
	return VZVirtioGenericDeviceSpecificConfiguration{VZVirtioDeviceSpecificConfiguration: VZVirtioDeviceSpecificConfigurationFromID(id)}
}
// Ensure VZVirtioGenericDeviceSpecificConfiguration implements IVZVirtioGenericDeviceSpecificConfiguration.
var _ IVZVirtioGenericDeviceSpecificConfiguration = VZVirtioGenericDeviceSpecificConfiguration{}

// An interface definition for the [VZVirtioGenericDeviceSpecificConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioGenericDeviceSpecificConfiguration.ConfigurationData]
//   - [IVZVirtioGenericDeviceSpecificConfiguration.InitWithConfigurationData]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioGenericDeviceSpecificConfiguration
type IVZVirtioGenericDeviceSpecificConfiguration interface {
	IVZVirtioDeviceSpecificConfiguration

	// Topic: Methods

	ConfigurationData() foundation.INSData
	InitWithConfigurationData(data objectivec.IObject) VZVirtioGenericDeviceSpecificConfiguration
}

// Init initializes the instance.
func (v VZVirtioGenericDeviceSpecificConfiguration) Init() VZVirtioGenericDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioGenericDeviceSpecificConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGenericDeviceSpecificConfiguration) Autorelease() VZVirtioGenericDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioGenericDeviceSpecificConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGenericDeviceSpecificConfiguration creates a new VZVirtioGenericDeviceSpecificConfiguration instance.
func NewVZVirtioGenericDeviceSpecificConfiguration() VZVirtioGenericDeviceSpecificConfiguration {
	class := getVZVirtioGenericDeviceSpecificConfigurationClass()
	rv := objc.Send[VZVirtioGenericDeviceSpecificConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioGenericDeviceSpecificConfiguration/initWithConfigurationData:
func NewVZVirtioGenericDeviceSpecificConfigurationWithConfigurationData(data objectivec.IObject) VZVirtioGenericDeviceSpecificConfiguration {
	instance := getVZVirtioGenericDeviceSpecificConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfigurationData:"), data)
	return VZVirtioGenericDeviceSpecificConfigurationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioGenericDeviceSpecificConfiguration/initWithConfigurationData:
func (v VZVirtioGenericDeviceSpecificConfiguration) InitWithConfigurationData(data objectivec.IObject) VZVirtioGenericDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioGenericDeviceSpecificConfiguration](v.ID, objc.Sel("initWithConfigurationData:"), data)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioGenericDeviceSpecificConfiguration/configurationData
func (v VZVirtioGenericDeviceSpecificConfiguration) ConfigurationData() foundation.INSData {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("configurationData"))
	return foundation.NSDataFromID(objc.ID(rv))
}

