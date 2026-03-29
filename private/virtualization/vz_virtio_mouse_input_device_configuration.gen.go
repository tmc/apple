// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioMouseInputDeviceConfiguration] class.
var (
	_VZVirtioMouseInputDeviceConfigurationClass     VZVirtioMouseInputDeviceConfigurationClass
	_VZVirtioMouseInputDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioMouseInputDeviceConfigurationClass() VZVirtioMouseInputDeviceConfigurationClass {
	_VZVirtioMouseInputDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioMouseInputDeviceConfigurationClass = VZVirtioMouseInputDeviceConfigurationClass{class: objc.GetClass("_VZVirtioMouseInputDeviceConfiguration")}
	})
	return _VZVirtioMouseInputDeviceConfigurationClass
}

// GetVZVirtioMouseInputDeviceConfigurationClass returns the class object for _VZVirtioMouseInputDeviceConfiguration.
func GetVZVirtioMouseInputDeviceConfigurationClass() VZVirtioMouseInputDeviceConfigurationClass {
	return getVZVirtioMouseInputDeviceConfigurationClass()
}

type VZVirtioMouseInputDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioMouseInputDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioMouseInputDeviceConfigurationClass) Alloc() VZVirtioMouseInputDeviceConfiguration {
	rv := objc.Send[VZVirtioMouseInputDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZVirtioMouseInputDeviceConfiguration.EncodeWithEncoder]
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioMouseInputDeviceConfiguration
type VZVirtioMouseInputDeviceConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZVirtioMouseInputDeviceConfigurationFromID constructs a [VZVirtioMouseInputDeviceConfiguration] from an objc.ID.
func VZVirtioMouseInputDeviceConfigurationFromID(id objc.ID) VZVirtioMouseInputDeviceConfiguration {
	return VZVirtioMouseInputDeviceConfiguration{VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFromID(id)}
}
// Ensure VZVirtioMouseInputDeviceConfiguration implements IVZVirtioMouseInputDeviceConfiguration.
var _ IVZVirtioMouseInputDeviceConfiguration = VZVirtioMouseInputDeviceConfiguration{}

// An interface definition for the [VZVirtioMouseInputDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioMouseInputDeviceConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioMouseInputDeviceConfiguration
type IVZVirtioMouseInputDeviceConfiguration interface {
	IVZPointingDeviceConfiguration

	// Topic: Methods

	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v VZVirtioMouseInputDeviceConfiguration) Init() VZVirtioMouseInputDeviceConfiguration {
	rv := objc.Send[VZVirtioMouseInputDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioMouseInputDeviceConfiguration) Autorelease() VZVirtioMouseInputDeviceConfiguration {
	rv := objc.Send[VZVirtioMouseInputDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioMouseInputDeviceConfiguration creates a new VZVirtioMouseInputDeviceConfiguration instance.
func NewVZVirtioMouseInputDeviceConfiguration() VZVirtioMouseInputDeviceConfiguration {
	class := getVZVirtioMouseInputDeviceConfigurationClass()
	rv := objc.Send[VZVirtioMouseInputDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioMouseInputDeviceConfiguration/encodeWithEncoder:
func (v VZVirtioMouseInputDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

