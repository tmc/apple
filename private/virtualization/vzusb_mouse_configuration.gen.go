// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZUSBMouseConfiguration] class.
var (
	_VZUSBMouseConfigurationClass     VZUSBMouseConfigurationClass
	_VZUSBMouseConfigurationClassOnce sync.Once
)

func getVZUSBMouseConfigurationClass() VZUSBMouseConfigurationClass {
	_VZUSBMouseConfigurationClassOnce.Do(func() {
		_VZUSBMouseConfigurationClass = VZUSBMouseConfigurationClass{class: objc.GetClass("_VZUSBMouseConfiguration")}
	})
	return _VZUSBMouseConfigurationClass
}

// GetVZUSBMouseConfigurationClass returns the class object for _VZUSBMouseConfiguration.
func GetVZUSBMouseConfigurationClass() VZUSBMouseConfigurationClass {
	return getVZUSBMouseConfigurationClass()
}

type VZUSBMouseConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBMouseConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBMouseConfigurationClass) Alloc() VZUSBMouseConfiguration {
	rv := objc.SendIfResponds[VZUSBMouseConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZUSBMouseConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZUSBMouseConfigurationFromID constructs a [VZUSBMouseConfiguration] from an objc.ID.
func VZUSBMouseConfigurationFromID(id objc.ID) VZUSBMouseConfiguration {
	return VZUSBMouseConfiguration{VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFromID(id)}
}

// Ensure VZUSBMouseConfiguration implements IVZUSBMouseConfiguration.
var _ IVZUSBMouseConfiguration = VZUSBMouseConfiguration{}

// An interface definition for the [VZUSBMouseConfiguration] class.
type IVZUSBMouseConfiguration interface {
	IVZPointingDeviceConfiguration
}

// Init initializes the instance.
func (v VZUSBMouseConfiguration) Init() VZUSBMouseConfiguration {
	rv := objc.SendIfResponds[VZUSBMouseConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBMouseConfiguration) Autorelease() VZUSBMouseConfiguration {
	rv := objc.SendIfResponds[VZUSBMouseConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMouseConfiguration creates a new VZUSBMouseConfiguration instance.
func NewVZUSBMouseConfiguration() VZUSBMouseConfiguration {
	class := getVZUSBMouseConfigurationClass()
	rv := objc.SendIfResponds[VZUSBMouseConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
