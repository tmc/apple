// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAppleTouchScreenConfiguration] class.
var (
	_VZAppleTouchScreenConfigurationClass     VZAppleTouchScreenConfigurationClass
	_VZAppleTouchScreenConfigurationClassOnce sync.Once
)

func getVZAppleTouchScreenConfigurationClass() VZAppleTouchScreenConfigurationClass {
	_VZAppleTouchScreenConfigurationClassOnce.Do(func() {
		_VZAppleTouchScreenConfigurationClass = VZAppleTouchScreenConfigurationClass{class: objc.GetClass("_VZAppleTouchScreenConfiguration")}
	})
	return _VZAppleTouchScreenConfigurationClass
}

// GetVZAppleTouchScreenConfigurationClass returns the class object for _VZAppleTouchScreenConfiguration.
func GetVZAppleTouchScreenConfigurationClass() VZAppleTouchScreenConfigurationClass {
	return getVZAppleTouchScreenConfigurationClass()
}

type VZAppleTouchScreenConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZAppleTouchScreenConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAppleTouchScreenConfigurationClass) Alloc() VZAppleTouchScreenConfiguration {
	rv := objc.Send[VZAppleTouchScreenConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZAppleTouchScreenConfiguration._registryProperties]
//   - [VZAppleTouchScreenConfiguration.Set_registryProperties]
//   - [VZAppleTouchScreenConfiguration._setRegistryProperties]
// See: https://developer.apple.com/documentation/Virtualization/_VZAppleTouchScreenConfiguration
type VZAppleTouchScreenConfiguration struct {
	VZMultiTouchDeviceConfiguration
}

// VZAppleTouchScreenConfigurationFromID constructs a [VZAppleTouchScreenConfiguration] from an objc.ID.
func VZAppleTouchScreenConfigurationFromID(id objc.ID) VZAppleTouchScreenConfiguration {
	return VZAppleTouchScreenConfiguration{VZMultiTouchDeviceConfiguration: VZMultiTouchDeviceConfigurationFromID(id)}
}
// Ensure VZAppleTouchScreenConfiguration implements IVZAppleTouchScreenConfiguration.
var _ IVZAppleTouchScreenConfiguration = VZAppleTouchScreenConfiguration{}

// An interface definition for the [VZAppleTouchScreenConfiguration] class.
//
// # Methods
//
//   - [IVZAppleTouchScreenConfiguration._registryProperties]
//   - [IVZAppleTouchScreenConfiguration.Set_registryProperties]
//   - [IVZAppleTouchScreenConfiguration._setRegistryProperties]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZAppleTouchScreenConfiguration
type IVZAppleTouchScreenConfiguration interface {
	IVZMultiTouchDeviceConfiguration

	// Topic: Methods

	_registryProperties() foundation.INSDictionary
	Set_registryProperties(value foundation.INSDictionary)
	_setRegistryProperties(properties objectivec.IObject)
}

// Init initializes the instance.
func (v VZAppleTouchScreenConfiguration) Init() VZAppleTouchScreenConfiguration {
	rv := objc.Send[VZAppleTouchScreenConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAppleTouchScreenConfiguration) Autorelease() VZAppleTouchScreenConfiguration {
	rv := objc.Send[VZAppleTouchScreenConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAppleTouchScreenConfiguration creates a new VZAppleTouchScreenConfiguration instance.
func NewVZAppleTouchScreenConfiguration() VZAppleTouchScreenConfiguration {
	class := getVZAppleTouchScreenConfigurationClass()
	rv := objc.Send[VZAppleTouchScreenConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZAppleTouchScreenConfiguration/_setRegistryProperties:
func (v VZAppleTouchScreenConfiguration) _setRegistryProperties(properties objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("_setRegistryProperties:"), properties)
}

// SetRegistryProperties is an exported wrapper for the private method _setRegistryProperties.
func (v VZAppleTouchScreenConfiguration) SetRegistryProperties(properties objectivec.IObject) {
	v._setRegistryProperties(properties)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZAppleTouchScreenConfiguration/_registryProperties
func (v VZAppleTouchScreenConfiguration) _registryProperties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_registryProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v VZAppleTouchScreenConfiguration) Set_registryProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("set_registryProperties:"), value)
}

