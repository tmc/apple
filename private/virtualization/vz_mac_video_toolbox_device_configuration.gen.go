// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacVideoToolboxDeviceConfiguration] class.
var (
	_VZMacVideoToolboxDeviceConfigurationClass     VZMacVideoToolboxDeviceConfigurationClass
	_VZMacVideoToolboxDeviceConfigurationClassOnce sync.Once
)

func getVZMacVideoToolboxDeviceConfigurationClass() VZMacVideoToolboxDeviceConfigurationClass {
	_VZMacVideoToolboxDeviceConfigurationClassOnce.Do(func() {
		_VZMacVideoToolboxDeviceConfigurationClass = VZMacVideoToolboxDeviceConfigurationClass{class: objc.GetClass("_VZMacVideoToolboxDeviceConfiguration")}
	})
	return _VZMacVideoToolboxDeviceConfigurationClass
}

// GetVZMacVideoToolboxDeviceConfigurationClass returns the class object for _VZMacVideoToolboxDeviceConfiguration.
func GetVZMacVideoToolboxDeviceConfigurationClass() VZMacVideoToolboxDeviceConfigurationClass {
	return getVZMacVideoToolboxDeviceConfigurationClass()
}

type VZMacVideoToolboxDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacVideoToolboxDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacVideoToolboxDeviceConfigurationClass) Alloc() VZMacVideoToolboxDeviceConfiguration {
	rv := objc.Send[VZMacVideoToolboxDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacVideoToolboxDeviceConfiguration
type VZMacVideoToolboxDeviceConfiguration struct {
	VZAcceleratorDeviceConfiguration
}

// VZMacVideoToolboxDeviceConfigurationFromID constructs a [VZMacVideoToolboxDeviceConfiguration] from an objc.ID.
func VZMacVideoToolboxDeviceConfigurationFromID(id objc.ID) VZMacVideoToolboxDeviceConfiguration {
	return VZMacVideoToolboxDeviceConfiguration{VZAcceleratorDeviceConfiguration: VZAcceleratorDeviceConfigurationFromID(id)}
}
// Ensure VZMacVideoToolboxDeviceConfiguration implements IVZMacVideoToolboxDeviceConfiguration.
var _ IVZMacVideoToolboxDeviceConfiguration = VZMacVideoToolboxDeviceConfiguration{}

// An interface definition for the [VZMacVideoToolboxDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacVideoToolboxDeviceConfiguration
type IVZMacVideoToolboxDeviceConfiguration interface {
	IVZAcceleratorDeviceConfiguration
}

// Init initializes the instance.
func (v VZMacVideoToolboxDeviceConfiguration) Init() VZMacVideoToolboxDeviceConfiguration {
	rv := objc.Send[VZMacVideoToolboxDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacVideoToolboxDeviceConfiguration) Autorelease() VZMacVideoToolboxDeviceConfiguration {
	rv := objc.Send[VZMacVideoToolboxDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacVideoToolboxDeviceConfiguration creates a new VZMacVideoToolboxDeviceConfiguration instance.
func NewVZMacVideoToolboxDeviceConfiguration() VZMacVideoToolboxDeviceConfiguration {
	class := getVZMacVideoToolboxDeviceConfigurationClass()
	rv := objc.Send[VZMacVideoToolboxDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacVideoToolboxDeviceConfiguration/_isSupported
func (_VZMacVideoToolboxDeviceConfigurationClass VZMacVideoToolboxDeviceConfigurationClass) _isSupported() bool {
	rv := objc.Send[bool](objc.ID(_VZMacVideoToolboxDeviceConfigurationClass.class), objc.Sel("_isSupported"))
	return rv
}

// IsSupported is an exported wrapper for the private method _isSupported.
func (_VZMacVideoToolboxDeviceConfigurationClass VZMacVideoToolboxDeviceConfigurationClass) IsSupported() bool {
	return _VZMacVideoToolboxDeviceConfigurationClass._isSupported()
}

