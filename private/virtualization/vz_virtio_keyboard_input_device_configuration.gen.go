// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZVirtioKeyboardInputDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

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
type IVZVirtioKeyboardInputDeviceConfiguration interface {
	IVZKeyboardConfiguration
}

// Init initializes the instance.
func (v VZVirtioKeyboardInputDeviceConfiguration) Init() VZVirtioKeyboardInputDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioKeyboardInputDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioKeyboardInputDeviceConfiguration) Autorelease() VZVirtioKeyboardInputDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioKeyboardInputDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioKeyboardInputDeviceConfiguration creates a new VZVirtioKeyboardInputDeviceConfiguration instance.
func NewVZVirtioKeyboardInputDeviceConfiguration() VZVirtioKeyboardInputDeviceConfiguration {
	class := getVZVirtioKeyboardInputDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZVirtioKeyboardInputDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
