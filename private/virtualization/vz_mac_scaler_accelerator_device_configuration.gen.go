// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacScalerAcceleratorDeviceConfiguration] class.
var (
	_VZMacScalerAcceleratorDeviceConfigurationClass     VZMacScalerAcceleratorDeviceConfigurationClass
	_VZMacScalerAcceleratorDeviceConfigurationClassOnce sync.Once
)

func getVZMacScalerAcceleratorDeviceConfigurationClass() VZMacScalerAcceleratorDeviceConfigurationClass {
	_VZMacScalerAcceleratorDeviceConfigurationClassOnce.Do(func() {
		_VZMacScalerAcceleratorDeviceConfigurationClass = VZMacScalerAcceleratorDeviceConfigurationClass{class: objc.GetClass("_VZMacScalerAcceleratorDeviceConfiguration")}
	})
	return _VZMacScalerAcceleratorDeviceConfigurationClass
}

// GetVZMacScalerAcceleratorDeviceConfigurationClass returns the class object for _VZMacScalerAcceleratorDeviceConfiguration.
func GetVZMacScalerAcceleratorDeviceConfigurationClass() VZMacScalerAcceleratorDeviceConfigurationClass {
	return getVZMacScalerAcceleratorDeviceConfigurationClass()
}

type VZMacScalerAcceleratorDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacScalerAcceleratorDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacScalerAcceleratorDeviceConfigurationClass) Alloc() VZMacScalerAcceleratorDeviceConfiguration {
	rv := objc.Send[VZMacScalerAcceleratorDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacScalerAcceleratorDeviceConfiguration
type VZMacScalerAcceleratorDeviceConfiguration struct {
	VZAcceleratorDeviceConfiguration
}

// VZMacScalerAcceleratorDeviceConfigurationFromID constructs a [VZMacScalerAcceleratorDeviceConfiguration] from an objc.ID.
func VZMacScalerAcceleratorDeviceConfigurationFromID(id objc.ID) VZMacScalerAcceleratorDeviceConfiguration {
	return VZMacScalerAcceleratorDeviceConfiguration{VZAcceleratorDeviceConfiguration: VZAcceleratorDeviceConfigurationFromID(id)}
}

// Ensure VZMacScalerAcceleratorDeviceConfiguration implements IVZMacScalerAcceleratorDeviceConfiguration.
var _ IVZMacScalerAcceleratorDeviceConfiguration = VZMacScalerAcceleratorDeviceConfiguration{}

// An interface definition for the [VZMacScalerAcceleratorDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacScalerAcceleratorDeviceConfiguration
type IVZMacScalerAcceleratorDeviceConfiguration interface {
	IVZAcceleratorDeviceConfiguration
}

// Init initializes the instance.
func (v VZMacScalerAcceleratorDeviceConfiguration) Init() VZMacScalerAcceleratorDeviceConfiguration {
	rv := objc.Send[VZMacScalerAcceleratorDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacScalerAcceleratorDeviceConfiguration) Autorelease() VZMacScalerAcceleratorDeviceConfiguration {
	rv := objc.Send[VZMacScalerAcceleratorDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacScalerAcceleratorDeviceConfiguration creates a new VZMacScalerAcceleratorDeviceConfiguration instance.
func NewVZMacScalerAcceleratorDeviceConfiguration() VZMacScalerAcceleratorDeviceConfiguration {
	class := getVZMacScalerAcceleratorDeviceConfigurationClass()
	rv := objc.Send[VZMacScalerAcceleratorDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
