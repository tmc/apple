// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioConsolePortConfiguration] class.
var (
	_VZVirtioConsolePortConfigurationClass     VZVirtioConsolePortConfigurationClass
	_VZVirtioConsolePortConfigurationClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationClass() VZVirtioConsolePortConfigurationClass {
	_VZVirtioConsolePortConfigurationClassOnce.Do(func() {
		_VZVirtioConsolePortConfigurationClass = VZVirtioConsolePortConfigurationClass{class: objc.GetClass("VZVirtioConsolePortConfiguration")}
	})
	return _VZVirtioConsolePortConfigurationClass
}

// GetVZVirtioConsolePortConfigurationClass returns the class object for VZVirtioConsolePortConfiguration.
func GetVZVirtioConsolePortConfigurationClass() VZVirtioConsolePortConfigurationClass {
	return getVZVirtioConsolePortConfigurationClass()
}

type VZVirtioConsolePortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioConsolePortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsolePortConfigurationClass) Alloc() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration
type VZVirtioConsolePortConfiguration struct {
	VZConsolePortConfiguration
}

// VZVirtioConsolePortConfigurationFromID constructs a [VZVirtioConsolePortConfiguration] from an objc.ID.
func VZVirtioConsolePortConfigurationFromID(id objc.ID) VZVirtioConsolePortConfiguration {
	return VZVirtioConsolePortConfiguration{VZConsolePortConfiguration: VZConsolePortConfigurationFromID(id)}
}
// Ensure VZVirtioConsolePortConfiguration implements IVZVirtioConsolePortConfiguration.
var _ IVZVirtioConsolePortConfiguration = VZVirtioConsolePortConfiguration{}

// An interface definition for the [VZVirtioConsolePortConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration
type IVZVirtioConsolePortConfiguration interface {
	IVZConsolePortConfiguration
}

// Init initializes the instance.
func (v VZVirtioConsolePortConfiguration) Init() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsolePortConfiguration) Autorelease() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfiguration creates a new VZVirtioConsolePortConfiguration instance.
func NewVZVirtioConsolePortConfiguration() VZVirtioConsolePortConfiguration {
	class := getVZVirtioConsolePortConfigurationClass()
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

