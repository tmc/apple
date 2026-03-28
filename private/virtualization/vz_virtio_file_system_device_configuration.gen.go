// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioFileSystemDeviceConfiguration] class.
var (
	_VZVirtioFileSystemDeviceConfigurationClass     VZVirtioFileSystemDeviceConfigurationClass
	_VZVirtioFileSystemDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioFileSystemDeviceConfigurationClass() VZVirtioFileSystemDeviceConfigurationClass {
	_VZVirtioFileSystemDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioFileSystemDeviceConfigurationClass = VZVirtioFileSystemDeviceConfigurationClass{class: objc.GetClass("VZVirtioFileSystemDeviceConfiguration")}
	})
	return _VZVirtioFileSystemDeviceConfigurationClass
}

// GetVZVirtioFileSystemDeviceConfigurationClass returns the class object for VZVirtioFileSystemDeviceConfiguration.
func GetVZVirtioFileSystemDeviceConfigurationClass() VZVirtioFileSystemDeviceConfigurationClass {
	return getVZVirtioFileSystemDeviceConfigurationClass()
}

type VZVirtioFileSystemDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioFileSystemDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioFileSystemDeviceConfigurationClass) Alloc() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration
type VZVirtioFileSystemDeviceConfiguration struct {
	VZDirectorySharingDeviceConfiguration
}

// VZVirtioFileSystemDeviceConfigurationFromID constructs a [VZVirtioFileSystemDeviceConfiguration] from an objc.ID.
func VZVirtioFileSystemDeviceConfigurationFromID(id objc.ID) VZVirtioFileSystemDeviceConfiguration {
	return VZVirtioFileSystemDeviceConfiguration{VZDirectorySharingDeviceConfiguration: VZDirectorySharingDeviceConfigurationFromID(id)}
}
// Ensure VZVirtioFileSystemDeviceConfiguration implements IVZVirtioFileSystemDeviceConfiguration.
var _ IVZVirtioFileSystemDeviceConfiguration = VZVirtioFileSystemDeviceConfiguration{}

// An interface definition for the [VZVirtioFileSystemDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioFileSystemDeviceConfiguration
type IVZVirtioFileSystemDeviceConfiguration interface {
	IVZDirectorySharingDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioFileSystemDeviceConfiguration) Init() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioFileSystemDeviceConfiguration) Autorelease() VZVirtioFileSystemDeviceConfiguration {
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioFileSystemDeviceConfiguration creates a new VZVirtioFileSystemDeviceConfiguration instance.
func NewVZVirtioFileSystemDeviceConfiguration() VZVirtioFileSystemDeviceConfiguration {
	class := getVZVirtioFileSystemDeviceConfigurationClass()
	rv := objc.Send[VZVirtioFileSystemDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

