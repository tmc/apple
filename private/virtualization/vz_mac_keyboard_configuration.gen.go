// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacKeyboardConfiguration] class.
var (
	_VZMacKeyboardConfigurationClass     VZMacKeyboardConfigurationClass
	_VZMacKeyboardConfigurationClassOnce sync.Once
)

func getVZMacKeyboardConfigurationClass() VZMacKeyboardConfigurationClass {
	_VZMacKeyboardConfigurationClassOnce.Do(func() {
		_VZMacKeyboardConfigurationClass = VZMacKeyboardConfigurationClass{class: objc.GetClass("VZMacKeyboardConfiguration")}
	})
	return _VZMacKeyboardConfigurationClass
}

// GetVZMacKeyboardConfigurationClass returns the class object for VZMacKeyboardConfiguration.
func GetVZMacKeyboardConfigurationClass() VZMacKeyboardConfigurationClass {
	return getVZMacKeyboardConfigurationClass()
}

type VZMacKeyboardConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacKeyboardConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacKeyboardConfigurationClass) Alloc() VZMacKeyboardConfiguration {
	rv := objc.SendIfResponds[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacKeyboardConfiguration._setSoftwareKeyboard]
//   - [VZMacKeyboardConfiguration._setSupportsGlobeKey]
//   - [VZMacKeyboardConfiguration._softwareKeyboard]
//   - [VZMacKeyboardConfiguration.Set_softwareKeyboard]
//   - [VZMacKeyboardConfiguration._supportsGlobeKey]
//   - [VZMacKeyboardConfiguration.Set_supportsGlobeKey]
type VZMacKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZMacKeyboardConfigurationFromID constructs a [VZMacKeyboardConfiguration] from an objc.ID.
func VZMacKeyboardConfigurationFromID(id objc.ID) VZMacKeyboardConfiguration {
	return VZMacKeyboardConfiguration{VZKeyboardConfiguration: VZKeyboardConfigurationFromID(id)}
}

// Ensure VZMacKeyboardConfiguration implements IVZMacKeyboardConfiguration.
var _ IVZMacKeyboardConfiguration = VZMacKeyboardConfiguration{}

// An interface definition for the [VZMacKeyboardConfiguration] class.
//
// # Methods
//
//   - [IVZMacKeyboardConfiguration._setSoftwareKeyboard]
//   - [IVZMacKeyboardConfiguration._setSupportsGlobeKey]
//   - [IVZMacKeyboardConfiguration._softwareKeyboard]
//   - [IVZMacKeyboardConfiguration.Set_softwareKeyboard]
//   - [IVZMacKeyboardConfiguration._supportsGlobeKey]
//   - [IVZMacKeyboardConfiguration.Set_supportsGlobeKey]
type IVZMacKeyboardConfiguration interface {
	IVZKeyboardConfiguration

	// Topic: Methods

	_setSoftwareKeyboard(keyboard bool)
	_setSupportsGlobeKey(key bool)
	_softwareKeyboard() bool
	Set_softwareKeyboard(value bool)
	_supportsGlobeKey() bool
	Set_supportsGlobeKey(value bool)
}

// Init initializes the instance.
func (v VZMacKeyboardConfiguration) Init() VZMacKeyboardConfiguration {
	rv := objc.SendIfResponds[VZMacKeyboardConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacKeyboardConfiguration) Autorelease() VZMacKeyboardConfiguration {
	rv := objc.SendIfResponds[VZMacKeyboardConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacKeyboardConfiguration creates a new VZMacKeyboardConfiguration instance.
func NewVZMacKeyboardConfiguration() VZMacKeyboardConfiguration {
	class := getVZMacKeyboardConfigurationClass()
	rv := objc.SendIfResponds[VZMacKeyboardConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZMacKeyboardConfiguration) _setSoftwareKeyboard(keyboard bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setSoftwareKeyboard:"), keyboard)
}

// SetSoftwareKeyboard is an exported wrapper for the private method _setSoftwareKeyboard.
func (v VZMacKeyboardConfiguration) SetSoftwareKeyboard(keyboard bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setSoftwareKeyboard:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setSoftwareKeyboard:"}
		return err
	}
	v._setSoftwareKeyboard(keyboard)
	return nil
}

// CanSetSoftwareKeyboard reports whether the receiver responds to the private selector _setSoftwareKeyboard:.
func (v VZMacKeyboardConfiguration) CanSetSoftwareKeyboard() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setSoftwareKeyboard:"))
}
func (v VZMacKeyboardConfiguration) _setSupportsGlobeKey(key bool) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_setSupportsGlobeKey:"), key)
}

// SetSupportsGlobeKey is an exported wrapper for the private method _setSupportsGlobeKey.
func (v VZMacKeyboardConfiguration) SetSupportsGlobeKey(key bool) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_setSupportsGlobeKey:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_setSupportsGlobeKey:"}
		return err
	}
	v._setSupportsGlobeKey(key)
	return nil
}

// CanSetSupportsGlobeKey reports whether the receiver responds to the private selector _setSupportsGlobeKey:.
func (v VZMacKeyboardConfiguration) CanSetSupportsGlobeKey() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_setSupportsGlobeKey:"))
}

func (v VZMacKeyboardConfiguration) _softwareKeyboard() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_softwareKeyboard"))
	return rv
}

// CanSoftwareKeyboard reports whether the receiver responds to the private selector _softwareKeyboard.
func (v VZMacKeyboardConfiguration) CanSoftwareKeyboard() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_softwareKeyboard"))
}

// SoftwareKeyboard is an exported wrapper for the private property _softwareKeyboard.
func (v VZMacKeyboardConfiguration) SoftwareKeyboard() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_softwareKeyboard")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_softwareKeyboard"}
	}
	return v._softwareKeyboard(), nil
}
func (v VZMacKeyboardConfiguration) Set_softwareKeyboard(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_softwareKeyboard:"), value)
}
func (v VZMacKeyboardConfiguration) _supportsGlobeKey() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_supportsGlobeKey"))
	return rv
}

// CanSupportsGlobeKey reports whether the receiver responds to the private selector _supportsGlobeKey.
func (v VZMacKeyboardConfiguration) CanSupportsGlobeKey() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_supportsGlobeKey"))
}

// SupportsGlobeKey is an exported wrapper for the private property _supportsGlobeKey.
func (v VZMacKeyboardConfiguration) SupportsGlobeKey() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_supportsGlobeKey")) {
		return false, &objc.UnrecognizedSelectorError{Selector: "_supportsGlobeKey"}
	}
	return v._supportsGlobeKey(), nil
}
func (v VZMacKeyboardConfiguration) Set_supportsGlobeKey(value bool) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("set_supportsGlobeKey:"), value)
}
