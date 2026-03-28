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
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacKeyboardConfiguration._setSoftwareKeyboard]
//   - [VZMacKeyboardConfiguration._setSupportsGlobeKey]
//   - [VZMacKeyboardConfiguration._softwareKeyboard]
//   - [VZMacKeyboardConfiguration.Set_softwareKeyboard]
//   - [VZMacKeyboardConfiguration._supportsGlobeKey]
//   - [VZMacKeyboardConfiguration.Set_supportsGlobeKey]
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration
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
func (m VZMacKeyboardConfiguration) Init() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacKeyboardConfiguration) Autorelease() VZMacKeyboardConfiguration {
	rv := objc.Send[VZMacKeyboardConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacKeyboardConfiguration creates a new VZMacKeyboardConfiguration instance.
func NewVZMacKeyboardConfiguration() VZMacKeyboardConfiguration {
	class := getVZMacKeyboardConfigurationClass()
	rv := objc.Send[VZMacKeyboardConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration/_setSoftwareKeyboard:
func (m VZMacKeyboardConfiguration) _setSoftwareKeyboard(keyboard bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setSoftwareKeyboard:"), keyboard)
}

// SetSoftwareKeyboard is an exported wrapper for the private method _setSoftwareKeyboard.
func (m VZMacKeyboardConfiguration) SetSoftwareKeyboard(keyboard bool) {
	m._setSoftwareKeyboard(keyboard)
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration/_setSupportsGlobeKey:
func (m VZMacKeyboardConfiguration) _setSupportsGlobeKey(key bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("_setSupportsGlobeKey:"), key)
}

// SetSupportsGlobeKey is an exported wrapper for the private method _setSupportsGlobeKey.
func (m VZMacKeyboardConfiguration) SetSupportsGlobeKey(key bool) {
	m._setSupportsGlobeKey(key)
}

// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration/_softwareKeyboard
func (m VZMacKeyboardConfiguration) _softwareKeyboard() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_softwareKeyboard"))
	return rv
}
func (m VZMacKeyboardConfiguration) Set_softwareKeyboard(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_softwareKeyboard:"), value)
}
// See: https://developer.apple.com/documentation/Virtualization/VZMacKeyboardConfiguration/_supportsGlobeKey
func (m VZMacKeyboardConfiguration) _supportsGlobeKey() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("_supportsGlobeKey"))
	return rv
}
func (m VZMacKeyboardConfiguration) Set_supportsGlobeKey(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("set_supportsGlobeKey:"), value)
}

