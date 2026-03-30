// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZUSBKeyboardConfiguration] class.
var (
	_VZUSBKeyboardConfigurationClass     VZUSBKeyboardConfigurationClass
	_VZUSBKeyboardConfigurationClassOnce sync.Once
)

func getVZUSBKeyboardConfigurationClass() VZUSBKeyboardConfigurationClass {
	_VZUSBKeyboardConfigurationClassOnce.Do(func() {
		_VZUSBKeyboardConfigurationClass = VZUSBKeyboardConfigurationClass{class: objc.GetClass("VZUSBKeyboardConfiguration")}
	})
	return _VZUSBKeyboardConfigurationClass
}

// GetVZUSBKeyboardConfigurationClass returns the class object for VZUSBKeyboardConfiguration.
func GetVZUSBKeyboardConfigurationClass() VZUSBKeyboardConfigurationClass {
	return getVZUSBKeyboardConfigurationClass()
}

type VZUSBKeyboardConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBKeyboardConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBKeyboardConfigurationClass) Alloc() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBKeyboardConfiguration
type VZUSBKeyboardConfiguration struct {
	VZKeyboardConfiguration
}

// VZUSBKeyboardConfigurationFromID constructs a [VZUSBKeyboardConfiguration] from an objc.ID.
func VZUSBKeyboardConfigurationFromID(id objc.ID) VZUSBKeyboardConfiguration {
	return VZUSBKeyboardConfiguration{VZKeyboardConfiguration: VZKeyboardConfigurationFromID(id)}
}

// Ensure VZUSBKeyboardConfiguration implements IVZUSBKeyboardConfiguration.
var _ IVZUSBKeyboardConfiguration = VZUSBKeyboardConfiguration{}

// An interface definition for the [VZUSBKeyboardConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBKeyboardConfiguration
type IVZUSBKeyboardConfiguration interface {
	IVZKeyboardConfiguration
}

// Init initializes the instance.
func (u VZUSBKeyboardConfiguration) Init() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBKeyboardConfiguration) Autorelease() VZUSBKeyboardConfiguration {
	rv := objc.Send[VZUSBKeyboardConfiguration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBKeyboardConfiguration creates a new VZUSBKeyboardConfiguration instance.
func NewVZUSBKeyboardConfiguration() VZUSBKeyboardConfiguration {
	class := getVZUSBKeyboardConfigurationClass()
	rv := objc.Send[VZUSBKeyboardConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
