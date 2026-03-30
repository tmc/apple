// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioConsoleDeviceConfiguration] class.
var (
	_VZVirtioConsoleDeviceConfigurationClass     VZVirtioConsoleDeviceConfigurationClass
	_VZVirtioConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceConfigurationClass() VZVirtioConsoleDeviceConfigurationClass {
	_VZVirtioConsoleDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioConsoleDeviceConfigurationClass = VZVirtioConsoleDeviceConfigurationClass{class: objc.GetClass("VZVirtioConsoleDeviceConfiguration")}
	})
	return _VZVirtioConsoleDeviceConfigurationClass
}

// GetVZVirtioConsoleDeviceConfigurationClass returns the class object for VZVirtioConsoleDeviceConfiguration.
func GetVZVirtioConsoleDeviceConfigurationClass() VZVirtioConsoleDeviceConfigurationClass {
	return getVZVirtioConsoleDeviceConfigurationClass()
}

type VZVirtioConsoleDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioConsoleDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsoleDeviceConfigurationClass) Alloc() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration
type VZVirtioConsoleDeviceConfiguration struct {
	VZConsoleDeviceConfiguration
}

// VZVirtioConsoleDeviceConfigurationFromID constructs a [VZVirtioConsoleDeviceConfiguration] from an objc.ID.
func VZVirtioConsoleDeviceConfigurationFromID(id objc.ID) VZVirtioConsoleDeviceConfiguration {
	return VZVirtioConsoleDeviceConfiguration{VZConsoleDeviceConfiguration: VZConsoleDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioConsoleDeviceConfiguration implements IVZVirtioConsoleDeviceConfiguration.
var _ IVZVirtioConsoleDeviceConfiguration = VZVirtioConsoleDeviceConfiguration{}

// An interface definition for the [VZVirtioConsoleDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration
type IVZVirtioConsoleDeviceConfiguration interface {
	IVZConsoleDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioConsoleDeviceConfiguration) Init() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsoleDeviceConfiguration) Autorelease() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceConfiguration creates a new VZVirtioConsoleDeviceConfiguration instance.
func NewVZVirtioConsoleDeviceConfiguration() VZVirtioConsoleDeviceConfiguration {
	class := getVZVirtioConsoleDeviceConfigurationClass()
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
