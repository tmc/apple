// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPointingDeviceConfiguration] class.
var (
	_VZPointingDeviceConfigurationClass     VZPointingDeviceConfigurationClass
	_VZPointingDeviceConfigurationClassOnce sync.Once
)

func getVZPointingDeviceConfigurationClass() VZPointingDeviceConfigurationClass {
	_VZPointingDeviceConfigurationClassOnce.Do(func() {
		_VZPointingDeviceConfigurationClass = VZPointingDeviceConfigurationClass{class: objc.GetClass("VZPointingDeviceConfiguration")}
	})
	return _VZPointingDeviceConfigurationClass
}

// GetVZPointingDeviceConfigurationClass returns the class object for VZPointingDeviceConfiguration.
func GetVZPointingDeviceConfigurationClass() VZPointingDeviceConfigurationClass {
	return getVZPointingDeviceConfigurationClass()
}

type VZPointingDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPointingDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPointingDeviceConfigurationClass) Alloc() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a pointing device configuration.
//
// # Overview
// 
// Don’t instantiate a [VZPointingDeviceConfiguration] directly, use one of
// its subclasses like [VZUSBScreenCoordinatePointingDeviceConfiguration]
// instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration
type VZPointingDeviceConfiguration struct {
	objectivec.Object
}

// VZPointingDeviceConfigurationFromID constructs a [VZPointingDeviceConfiguration] from an objc.ID.
//
// The base class for a pointing device configuration.
func VZPointingDeviceConfigurationFromID(id objc.ID) VZPointingDeviceConfiguration {
	return VZPointingDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZPointingDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZPointingDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration
type IVZPointingDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p VZPointingDeviceConfiguration) Init() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VZPointingDeviceConfiguration) Autorelease() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPointingDeviceConfiguration creates a new VZPointingDeviceConfiguration instance.
func NewVZPointingDeviceConfiguration() VZPointingDeviceConfiguration {
	class := getVZPointingDeviceConfigurationClass()
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

