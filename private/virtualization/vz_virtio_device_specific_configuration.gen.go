// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioDeviceSpecificConfiguration] class.
var (
	_VZVirtioDeviceSpecificConfigurationClass     VZVirtioDeviceSpecificConfigurationClass
	_VZVirtioDeviceSpecificConfigurationClassOnce sync.Once
)

func getVZVirtioDeviceSpecificConfigurationClass() VZVirtioDeviceSpecificConfigurationClass {
	_VZVirtioDeviceSpecificConfigurationClassOnce.Do(func() {
		_VZVirtioDeviceSpecificConfigurationClass = VZVirtioDeviceSpecificConfigurationClass{class: objc.GetClass("_VZVirtioDeviceSpecificConfiguration")}
	})
	return _VZVirtioDeviceSpecificConfigurationClass
}

// GetVZVirtioDeviceSpecificConfigurationClass returns the class object for _VZVirtioDeviceSpecificConfiguration.
func GetVZVirtioDeviceSpecificConfigurationClass() VZVirtioDeviceSpecificConfigurationClass {
	return getVZVirtioDeviceSpecificConfigurationClass()
}

type VZVirtioDeviceSpecificConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioDeviceSpecificConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioDeviceSpecificConfigurationClass) Alloc() VZVirtioDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioDeviceSpecificConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtioDeviceSpecificConfiguration._configuration]
//   - [VZVirtioDeviceSpecificConfiguration._init]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioDeviceSpecificConfiguration
type VZVirtioDeviceSpecificConfiguration struct {
	objectivec.Object
}

// VZVirtioDeviceSpecificConfigurationFromID constructs a [VZVirtioDeviceSpecificConfiguration] from an objc.ID.
func VZVirtioDeviceSpecificConfigurationFromID(id objc.ID) VZVirtioDeviceSpecificConfiguration {
	return VZVirtioDeviceSpecificConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZVirtioDeviceSpecificConfiguration implements IVZVirtioDeviceSpecificConfiguration.
var _ IVZVirtioDeviceSpecificConfiguration = VZVirtioDeviceSpecificConfiguration{}

// An interface definition for the [VZVirtioDeviceSpecificConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioDeviceSpecificConfiguration._configuration]
//   - [IVZVirtioDeviceSpecificConfiguration._init]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioDeviceSpecificConfiguration
type IVZVirtioDeviceSpecificConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_configuration() objectivec.IObject
	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZVirtioDeviceSpecificConfiguration) Init() VZVirtioDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioDeviceSpecificConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioDeviceSpecificConfiguration) Autorelease() VZVirtioDeviceSpecificConfiguration {
	rv := objc.Send[VZVirtioDeviceSpecificConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioDeviceSpecificConfiguration creates a new VZVirtioDeviceSpecificConfiguration instance.
func NewVZVirtioDeviceSpecificConfiguration() VZVirtioDeviceSpecificConfiguration {
	class := getVZVirtioDeviceSpecificConfigurationClass()
	rv := objc.Send[VZVirtioDeviceSpecificConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioDeviceSpecificConfiguration/_init
func (v VZVirtioDeviceSpecificConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioDeviceSpecificConfiguration/_configuration
func (v VZVirtioDeviceSpecificConfiguration) _configuration() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_configuration"))
	return objectivec.Object{ID: rv}
}

