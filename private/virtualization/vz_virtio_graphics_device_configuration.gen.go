// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioGraphicsDeviceConfiguration] class.
var (
	_VZVirtioGraphicsDeviceConfigurationClass     VZVirtioGraphicsDeviceConfigurationClass
	_VZVirtioGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceConfigurationClass() VZVirtioGraphicsDeviceConfigurationClass {
	_VZVirtioGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioGraphicsDeviceConfigurationClass = VZVirtioGraphicsDeviceConfigurationClass{class: objc.GetClass("VZVirtioGraphicsDeviceConfiguration")}
	})
	return _VZVirtioGraphicsDeviceConfigurationClass
}

// GetVZVirtioGraphicsDeviceConfigurationClass returns the class object for VZVirtioGraphicsDeviceConfiguration.
func GetVZVirtioGraphicsDeviceConfigurationClass() VZVirtioGraphicsDeviceConfigurationClass {
	return getVZVirtioGraphicsDeviceConfigurationClass()
}

type VZVirtioGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioGraphicsDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsDeviceConfigurationClass) Alloc() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type VZVirtioGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZVirtioGraphicsDeviceConfigurationFromID constructs a [VZVirtioGraphicsDeviceConfiguration] from an objc.ID.
func VZVirtioGraphicsDeviceConfigurationFromID(id objc.ID) VZVirtioGraphicsDeviceConfiguration {
	return VZVirtioGraphicsDeviceConfiguration{VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioGraphicsDeviceConfiguration implements IVZVirtioGraphicsDeviceConfiguration.
var _ IVZVirtioGraphicsDeviceConfiguration = VZVirtioGraphicsDeviceConfiguration{}

// An interface definition for the [VZVirtioGraphicsDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type IVZVirtioGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioGraphicsDeviceConfiguration) Init() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsDeviceConfiguration) Autorelease() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDeviceConfiguration creates a new VZVirtioGraphicsDeviceConfiguration instance.
func NewVZVirtioGraphicsDeviceConfiguration() VZVirtioGraphicsDeviceConfiguration {
	class := getVZVirtioGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/_maximumAllowedDisplayCount
func (_VZVirtioGraphicsDeviceConfigurationClass VZVirtioGraphicsDeviceConfigurationClass) _maximumAllowedDisplayCount() uint64 {
	rv := objc.Send[uint64](objc.ID(_VZVirtioGraphicsDeviceConfigurationClass.class), objc.Sel("_maximumAllowedDisplayCount"))
	return rv
}

// MaximumAllowedDisplayCount is an exported wrapper for the private method _maximumAllowedDisplayCount.
func (_VZVirtioGraphicsDeviceConfigurationClass VZVirtioGraphicsDeviceConfigurationClass) MaximumAllowedDisplayCount() uint64 {
	return _VZVirtioGraphicsDeviceConfigurationClass._maximumAllowedDisplayCount()
}
