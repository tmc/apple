// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacTouchIDDeviceConfiguration] class.
var (
	_VZMacTouchIDDeviceConfigurationClass     VZMacTouchIDDeviceConfigurationClass
	_VZMacTouchIDDeviceConfigurationClassOnce sync.Once
)

func getVZMacTouchIDDeviceConfigurationClass() VZMacTouchIDDeviceConfigurationClass {
	_VZMacTouchIDDeviceConfigurationClassOnce.Do(func() {
		_VZMacTouchIDDeviceConfigurationClass = VZMacTouchIDDeviceConfigurationClass{class: objc.GetClass("_VZMacTouchIDDeviceConfiguration")}
	})
	return _VZMacTouchIDDeviceConfigurationClass
}

// GetVZMacTouchIDDeviceConfigurationClass returns the class object for _VZMacTouchIDDeviceConfiguration.
func GetVZMacTouchIDDeviceConfigurationClass() VZMacTouchIDDeviceConfigurationClass {
	return getVZMacTouchIDDeviceConfigurationClass()
}

type VZMacTouchIDDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacTouchIDDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacTouchIDDeviceConfigurationClass) Alloc() VZMacTouchIDDeviceConfiguration {
	rv := objc.Send[VZMacTouchIDDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacTouchIDDeviceConfiguration
type VZMacTouchIDDeviceConfiguration struct {
	VZBiometricDeviceConfiguration
}

// VZMacTouchIDDeviceConfigurationFromID constructs a [VZMacTouchIDDeviceConfiguration] from an objc.ID.
func VZMacTouchIDDeviceConfigurationFromID(id objc.ID) VZMacTouchIDDeviceConfiguration {
	return VZMacTouchIDDeviceConfiguration{VZBiometricDeviceConfiguration: VZBiometricDeviceConfigurationFromID(id)}
}

// Ensure VZMacTouchIDDeviceConfiguration implements IVZMacTouchIDDeviceConfiguration.
var _ IVZMacTouchIDDeviceConfiguration = VZMacTouchIDDeviceConfiguration{}

// An interface definition for the [VZMacTouchIDDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacTouchIDDeviceConfiguration
type IVZMacTouchIDDeviceConfiguration interface {
	IVZBiometricDeviceConfiguration
}

// Init initializes the instance.
func (v VZMacTouchIDDeviceConfiguration) Init() VZMacTouchIDDeviceConfiguration {
	rv := objc.Send[VZMacTouchIDDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacTouchIDDeviceConfiguration) Autorelease() VZMacTouchIDDeviceConfiguration {
	rv := objc.Send[VZMacTouchIDDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacTouchIDDeviceConfiguration creates a new VZMacTouchIDDeviceConfiguration instance.
func NewVZMacTouchIDDeviceConfiguration() VZMacTouchIDDeviceConfiguration {
	class := getVZMacTouchIDDeviceConfigurationClass()
	rv := objc.Send[VZMacTouchIDDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
