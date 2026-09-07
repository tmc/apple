// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioNetworkDeviceConfiguration] class.
var (
	_VZVirtioNetworkDeviceConfigurationClass     VZVirtioNetworkDeviceConfigurationClass
	_VZVirtioNetworkDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioNetworkDeviceConfigurationClass() VZVirtioNetworkDeviceConfigurationClass {
	_VZVirtioNetworkDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioNetworkDeviceConfigurationClass = VZVirtioNetworkDeviceConfigurationClass{class: objc.GetClass("VZVirtioNetworkDeviceConfiguration")}
	})
	return _VZVirtioNetworkDeviceConfigurationClass
}

// GetVZVirtioNetworkDeviceConfigurationClass returns the class object for VZVirtioNetworkDeviceConfiguration.
func GetVZVirtioNetworkDeviceConfigurationClass() VZVirtioNetworkDeviceConfigurationClass {
	return getVZVirtioNetworkDeviceConfigurationClass()
}

type VZVirtioNetworkDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioNetworkDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioNetworkDeviceConfigurationClass) Alloc() VZVirtioNetworkDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZVirtioNetworkDeviceConfiguration struct {
	VZNetworkDeviceConfiguration
}

// VZVirtioNetworkDeviceConfigurationFromID constructs a [VZVirtioNetworkDeviceConfiguration] from an objc.ID.
func VZVirtioNetworkDeviceConfigurationFromID(id objc.ID) VZVirtioNetworkDeviceConfiguration {
	return VZVirtioNetworkDeviceConfiguration{VZNetworkDeviceConfiguration: VZNetworkDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioNetworkDeviceConfiguration implements IVZVirtioNetworkDeviceConfiguration.
var _ IVZVirtioNetworkDeviceConfiguration = VZVirtioNetworkDeviceConfiguration{}

// An interface definition for the [VZVirtioNetworkDeviceConfiguration] class.
type IVZVirtioNetworkDeviceConfiguration interface {
	IVZNetworkDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioNetworkDeviceConfiguration) Init() VZVirtioNetworkDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioNetworkDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioNetworkDeviceConfiguration) Autorelease() VZVirtioNetworkDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioNetworkDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioNetworkDeviceConfiguration creates a new VZVirtioNetworkDeviceConfiguration instance.
func NewVZVirtioNetworkDeviceConfiguration() VZVirtioNetworkDeviceConfiguration {
	class := getVZVirtioNetworkDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZVirtioNetworkDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
