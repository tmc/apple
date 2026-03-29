// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZUSBTouchScreenConfiguration] class.
var (
	_VZUSBTouchScreenConfigurationClass     VZUSBTouchScreenConfigurationClass
	_VZUSBTouchScreenConfigurationClassOnce sync.Once
)

func getVZUSBTouchScreenConfigurationClass() VZUSBTouchScreenConfigurationClass {
	_VZUSBTouchScreenConfigurationClassOnce.Do(func() {
		_VZUSBTouchScreenConfigurationClass = VZUSBTouchScreenConfigurationClass{class: objc.GetClass("_VZUSBTouchScreenConfiguration")}
	})
	return _VZUSBTouchScreenConfigurationClass
}

// GetVZUSBTouchScreenConfigurationClass returns the class object for _VZUSBTouchScreenConfiguration.
func GetVZUSBTouchScreenConfigurationClass() VZUSBTouchScreenConfigurationClass {
	return getVZUSBTouchScreenConfigurationClass()
}

type VZUSBTouchScreenConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBTouchScreenConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBTouchScreenConfigurationClass) Alloc() VZUSBTouchScreenConfiguration {
	rv := objc.Send[VZUSBTouchScreenConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBTouchScreenConfiguration
type VZUSBTouchScreenConfiguration struct {
	VZMultiTouchDeviceConfiguration
}

// VZUSBTouchScreenConfigurationFromID constructs a [VZUSBTouchScreenConfiguration] from an objc.ID.
func VZUSBTouchScreenConfigurationFromID(id objc.ID) VZUSBTouchScreenConfiguration {
	return VZUSBTouchScreenConfiguration{VZMultiTouchDeviceConfiguration: VZMultiTouchDeviceConfigurationFromID(id)}
}
// Ensure VZUSBTouchScreenConfiguration implements IVZUSBTouchScreenConfiguration.
var _ IVZUSBTouchScreenConfiguration = VZUSBTouchScreenConfiguration{}

// An interface definition for the [VZUSBTouchScreenConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBTouchScreenConfiguration
type IVZUSBTouchScreenConfiguration interface {
	IVZMultiTouchDeviceConfiguration
}

// Init initializes the instance.
func (v VZUSBTouchScreenConfiguration) Init() VZUSBTouchScreenConfiguration {
	rv := objc.Send[VZUSBTouchScreenConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBTouchScreenConfiguration) Autorelease() VZUSBTouchScreenConfiguration {
	rv := objc.Send[VZUSBTouchScreenConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBTouchScreenConfiguration creates a new VZUSBTouchScreenConfiguration instance.
func NewVZUSBTouchScreenConfiguration() VZUSBTouchScreenConfiguration {
	class := getVZUSBTouchScreenConfigurationClass()
	rv := objc.Send[VZUSBTouchScreenConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

