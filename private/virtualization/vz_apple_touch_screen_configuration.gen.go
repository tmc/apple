// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZAppleTouchScreenConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZAppleTouchScreenConfiguration._registryProperties]
//   - [VZAppleTouchScreenConfiguration.Set_registryProperties]
//   - [VZAppleTouchScreenConfiguration._setRegistryProperties]
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
type IVZAppleTouchScreenConfiguration interface {
	IVZMultiTouchDeviceConfiguration

	// Topic: Methods

	_registryProperties() foundation.INSDictionary
	Set_registryProperties(value foundation.INSDictionary)
	_setRegistryProperties(properties objectivec.IObject)
}

// Init initializes the instance.
func (v VZAppleTouchScreenConfiguration) Init() VZAppleTouchScreenConfiguration {
	rv := objc.SendIfResponds[VZAppleTouchScreenConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZAppleTouchScreenConfiguration) Autorelease() VZAppleTouchScreenConfiguration {
	rv := objc.SendIfResponds[VZAppleTouchScreenConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAppleTouchScreenConfiguration creates a new VZAppleTouchScreenConfiguration instance.
func NewVZAppleTouchScreenConfiguration() VZAppleTouchScreenConfiguration {
	class := getVZAppleTouchScreenConfigurationClass()
	rv := objc.SendIfResponds[VZAppleTouchScreenConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZAppleTouchScreenConfiguration) _setRegistryProperties(properties objectivec.IObject) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setRegistryProperties:"), properties)
}

// SetRegistryProperties is an exported wrapper for the private method _setRegistryProperties.
func (v VZAppleTouchScreenConfiguration) SetRegistryProperties(properties objectivec.IObject) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setRegistryProperties:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setRegistryProperties:"}
		return err
	}
	v._setRegistryProperties(properties)
	return nil
}

// CanSetRegistryProperties reports whether the receiver responds to the private selector _setRegistryProperties:.
func (v VZAppleTouchScreenConfiguration) CanSetRegistryProperties() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setRegistryProperties:"))
}

func (v VZAppleTouchScreenConfiguration) _registryProperties() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_registryProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// CanRegistryProperties reports whether the receiver responds to the private selector _registryProperties.
func (v VZAppleTouchScreenConfiguration) CanRegistryProperties() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_registryProperties"))
}

// RegistryProperties is an exported wrapper for the private property _registryProperties.
func (v VZAppleTouchScreenConfiguration) RegistryProperties() (foundation.INSDictionary, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_registryProperties")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_registryProperties"}
	}
	return v._registryProperties(), nil
}
func (v VZAppleTouchScreenConfiguration) Set_registryProperties(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_registryProperties:"), value)
}
