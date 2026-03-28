// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioGraphicsScanoutConfiguration] class.
var (
	_VZVirtioGraphicsScanoutConfigurationClass     VZVirtioGraphicsScanoutConfigurationClass
	_VZVirtioGraphicsScanoutConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsScanoutConfigurationClass() VZVirtioGraphicsScanoutConfigurationClass {
	_VZVirtioGraphicsScanoutConfigurationClassOnce.Do(func() {
		_VZVirtioGraphicsScanoutConfigurationClass = VZVirtioGraphicsScanoutConfigurationClass{class: objc.GetClass("VZVirtioGraphicsScanoutConfiguration")}
	})
	return _VZVirtioGraphicsScanoutConfigurationClass
}

// GetVZVirtioGraphicsScanoutConfigurationClass returns the class object for VZVirtioGraphicsScanoutConfiguration.
func GetVZVirtioGraphicsScanoutConfigurationClass() VZVirtioGraphicsScanoutConfigurationClass {
	return getVZVirtioGraphicsScanoutConfigurationClass()
}

type VZVirtioGraphicsScanoutConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioGraphicsScanoutConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsScanoutConfigurationClass) Alloc() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration
type VZVirtioGraphicsScanoutConfiguration struct {
	VZGraphicsDisplayConfiguration
}

// VZVirtioGraphicsScanoutConfigurationFromID constructs a [VZVirtioGraphicsScanoutConfiguration] from an objc.ID.
func VZVirtioGraphicsScanoutConfigurationFromID(id objc.ID) VZVirtioGraphicsScanoutConfiguration {
	return VZVirtioGraphicsScanoutConfiguration{VZGraphicsDisplayConfiguration: VZGraphicsDisplayConfigurationFromID(id)}
}
// Ensure VZVirtioGraphicsScanoutConfiguration implements IVZVirtioGraphicsScanoutConfiguration.
var _ IVZVirtioGraphicsScanoutConfiguration = VZVirtioGraphicsScanoutConfiguration{}

// An interface definition for the [VZVirtioGraphicsScanoutConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsScanoutConfiguration
type IVZVirtioGraphicsScanoutConfiguration interface {
	IVZGraphicsDisplayConfiguration
}

// Init initializes the instance.
func (v VZVirtioGraphicsScanoutConfiguration) Init() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsScanoutConfiguration) Autorelease() VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsScanoutConfiguration creates a new VZVirtioGraphicsScanoutConfiguration instance.
func NewVZVirtioGraphicsScanoutConfiguration() VZVirtioGraphicsScanoutConfiguration {
	class := getVZVirtioGraphicsScanoutConfigurationClass()
	rv := objc.Send[VZVirtioGraphicsScanoutConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

