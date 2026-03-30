// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioKeyboardInputDeviceConfiguration] class.
var (
	_VZVirtioKeyboardInputDeviceConfigurationClass     VZVirtioKeyboardInputDeviceConfigurationClass
	_VZVirtioKeyboardInputDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioKeyboardInputDeviceConfigurationClass() VZVirtioKeyboardInputDeviceConfigurationClass {
	_VZVirtioKeyboardInputDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioKeyboardInputDeviceConfigurationClass = VZVirtioKeyboardInputDeviceConfigurationClass{class: objc.GetClass("_VZVirtioKeyboardInputDeviceConfiguration")}
	})
	return _VZVirtioKeyboardInputDeviceConfigurationClass
}

// GetVZVirtioKeyboardInputDeviceConfigurationClass returns the class object for _VZVirtioKeyboardInputDeviceConfiguration.
func GetVZVirtioKeyboardInputDeviceConfigurationClass() VZVirtioKeyboardInputDeviceConfigurationClass {
	return getVZVirtioKeyboardInputDeviceConfigurationClass()
}

type VZVirtioKeyboardInputDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioKeyboardInputDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioKeyboardInputDeviceConfigurationClass) Alloc() VZVirtioKeyboardInputDeviceConfiguration {
	rv := objc.Send[VZVirtioKeyboardInputDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtioKeyboardInputDeviceConfiguration._keyboardWithDeviceIdentifier]
//   - [VZVirtioKeyboardInputDeviceConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioKeyboardInputDeviceConfiguration
type VZVirtioKeyboardInputDeviceConfiguration struct {
	VZKeyboardConfiguration
}

// VZVirtioKeyboardInputDeviceConfigurationFromID constructs a [VZVirtioKeyboardInputDeviceConfiguration] from an objc.ID.
func VZVirtioKeyboardInputDeviceConfigurationFromID(id objc.ID) VZVirtioKeyboardInputDeviceConfiguration {
	return VZVirtioKeyboardInputDeviceConfiguration{VZKeyboardConfiguration: VZKeyboardConfigurationFromID(id)}
}

// Ensure VZVirtioKeyboardInputDeviceConfiguration implements IVZVirtioKeyboardInputDeviceConfiguration.
var _ IVZVirtioKeyboardInputDeviceConfiguration = VZVirtioKeyboardInputDeviceConfiguration{}

// An interface definition for the [VZVirtioKeyboardInputDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZVirtioKeyboardInputDeviceConfiguration._keyboardWithDeviceIdentifier]
//   - [IVZVirtioKeyboardInputDeviceConfiguration.EncodeWithEncoder]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioKeyboardInputDeviceConfiguration
type IVZVirtioKeyboardInputDeviceConfiguration interface {
	IVZKeyboardConfiguration

	// Topic: Methods

	_keyboardWithDeviceIdentifier(identifier uint32) objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v VZVirtioKeyboardInputDeviceConfiguration) Init() VZVirtioKeyboardInputDeviceConfiguration {
	rv := objc.Send[VZVirtioKeyboardInputDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioKeyboardInputDeviceConfiguration) Autorelease() VZVirtioKeyboardInputDeviceConfiguration {
	rv := objc.Send[VZVirtioKeyboardInputDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioKeyboardInputDeviceConfiguration creates a new VZVirtioKeyboardInputDeviceConfiguration instance.
func NewVZVirtioKeyboardInputDeviceConfiguration() VZVirtioKeyboardInputDeviceConfiguration {
	class := getVZVirtioKeyboardInputDeviceConfigurationClass()
	rv := objc.Send[VZVirtioKeyboardInputDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioKeyboardInputDeviceConfiguration/_keyboardWithDeviceIdentifier:
func (v VZVirtioKeyboardInputDeviceConfiguration) _keyboardWithDeviceIdentifier(identifier uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_keyboardWithDeviceIdentifier:"), identifier)
	return objectivec.Object{ID: rv}
}

// KeyboardWithDeviceIdentifier is an exported wrapper for the private method _keyboardWithDeviceIdentifier.
func (v VZVirtioKeyboardInputDeviceConfiguration) KeyboardWithDeviceIdentifier(identifier uint32) objectivec.IObject {
	return v._keyboardWithDeviceIdentifier(identifier)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioKeyboardInputDeviceConfiguration/encodeWithEncoder:
func (v VZVirtioKeyboardInputDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}
