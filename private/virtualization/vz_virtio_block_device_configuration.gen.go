// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioBlockDeviceConfiguration] class.
var (
	_VZVirtioBlockDeviceConfigurationClass     VZVirtioBlockDeviceConfigurationClass
	_VZVirtioBlockDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioBlockDeviceConfigurationClass() VZVirtioBlockDeviceConfigurationClass {
	_VZVirtioBlockDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioBlockDeviceConfigurationClass = VZVirtioBlockDeviceConfigurationClass{class: objc.GetClass("VZVirtioBlockDeviceConfiguration")}
	})
	return _VZVirtioBlockDeviceConfigurationClass
}

// GetVZVirtioBlockDeviceConfigurationClass returns the class object for VZVirtioBlockDeviceConfiguration.
func GetVZVirtioBlockDeviceConfigurationClass() VZVirtioBlockDeviceConfigurationClass {
	return getVZVirtioBlockDeviceConfigurationClass()
}

type VZVirtioBlockDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioBlockDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioBlockDeviceConfigurationClass) Alloc() VZVirtioBlockDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioBlockDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZVirtioBlockDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZVirtioBlockDeviceConfigurationFromID constructs a [VZVirtioBlockDeviceConfiguration] from an objc.ID.
func VZVirtioBlockDeviceConfigurationFromID(id objc.ID) VZVirtioBlockDeviceConfiguration {
	return VZVirtioBlockDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioBlockDeviceConfiguration implements IVZVirtioBlockDeviceConfiguration.
var _ IVZVirtioBlockDeviceConfiguration = VZVirtioBlockDeviceConfiguration{}

// An interface definition for the [VZVirtioBlockDeviceConfiguration] class.
type IVZVirtioBlockDeviceConfiguration interface {
	IVZStorageDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioBlockDeviceConfiguration) Init() VZVirtioBlockDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioBlockDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioBlockDeviceConfiguration) Autorelease() VZVirtioBlockDeviceConfiguration {
	rv := objc.SendIfResponds[VZVirtioBlockDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioBlockDeviceConfiguration creates a new VZVirtioBlockDeviceConfiguration instance.
func NewVZVirtioBlockDeviceConfiguration() VZVirtioBlockDeviceConfiguration {
	class := getVZVirtioBlockDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZVirtioBlockDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
