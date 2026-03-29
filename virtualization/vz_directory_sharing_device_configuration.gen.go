// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDirectorySharingDeviceConfiguration] class.
var (
	_VZDirectorySharingDeviceConfigurationClass     VZDirectorySharingDeviceConfigurationClass
	_VZDirectorySharingDeviceConfigurationClassOnce sync.Once
)

func getVZDirectorySharingDeviceConfigurationClass() VZDirectorySharingDeviceConfigurationClass {
	_VZDirectorySharingDeviceConfigurationClassOnce.Do(func() {
		_VZDirectorySharingDeviceConfigurationClass = VZDirectorySharingDeviceConfigurationClass{class: objc.GetClass("VZDirectorySharingDeviceConfiguration")}
	})
	return _VZDirectorySharingDeviceConfigurationClass
}

// GetVZDirectorySharingDeviceConfigurationClass returns the class object for VZDirectorySharingDeviceConfiguration.
func GetVZDirectorySharingDeviceConfigurationClass() VZDirectorySharingDeviceConfigurationClass {
	return getVZDirectorySharingDeviceConfigurationClass()
}

type VZDirectorySharingDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDirectorySharingDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDirectorySharingDeviceConfigurationClass) Alloc() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a directory sharing device configuration.
//
// # Overview
// 
// Don’t instantiate [VZDirectorySharingDeviceConfiguration] directly.
// Instead use one of its subclasses, like
// [VZVirtioFileSystemDeviceConfiguration].
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration
type VZDirectorySharingDeviceConfiguration struct {
	objectivec.Object
}

// VZDirectorySharingDeviceConfigurationFromID constructs a [VZDirectorySharingDeviceConfiguration] from an objc.ID.
//
// The base class for a directory sharing device configuration.
func VZDirectorySharingDeviceConfigurationFromID(id objc.ID) VZDirectorySharingDeviceConfiguration {
	return VZDirectorySharingDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZDirectorySharingDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZDirectorySharingDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDeviceConfiguration
type IVZDirectorySharingDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d VZDirectorySharingDeviceConfiguration) Init() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDirectorySharingDeviceConfiguration) Autorelease() VZDirectorySharingDeviceConfiguration {
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDeviceConfiguration creates a new VZDirectorySharingDeviceConfiguration instance.
func NewVZDirectorySharingDeviceConfiguration() VZDirectorySharingDeviceConfiguration {
	class := getVZDirectorySharingDeviceConfigurationClass()
	rv := objc.Send[VZDirectorySharingDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

