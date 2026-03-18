// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPlatformConfiguration] class.
var (
	_VZPlatformConfigurationClass     VZPlatformConfigurationClass
	_VZPlatformConfigurationClassOnce sync.Once
)

func getVZPlatformConfigurationClass() VZPlatformConfigurationClass {
	_VZPlatformConfigurationClassOnce.Do(func() {
		_VZPlatformConfigurationClass = VZPlatformConfigurationClass{class: objc.GetClass("VZPlatformConfiguration")}
	})
	return _VZPlatformConfigurationClass
}

// GetVZPlatformConfigurationClass returns the class object for VZPlatformConfiguration.
func GetVZPlatformConfigurationClass() VZPlatformConfigurationClass {
	return getVZPlatformConfigurationClass()
}

type VZPlatformConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPlatformConfigurationClass) Alloc() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a platform configuration.
//
// # Overview
// 
// Don’t instantiate directly [VZPlatformConfiguration], use one of its
// subclasses, such as [VZGenericPlatformConfiguration] or
// [VZMacPlatformConfiguration] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration
type VZPlatformConfiguration struct {
	objectivec.Object
}

// VZPlatformConfigurationFromID constructs a [VZPlatformConfiguration] from an objc.ID.
//
// The base class for a platform configuration.
func VZPlatformConfigurationFromID(id objc.ID) VZPlatformConfiguration {
	return VZPlatformConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZPlatformConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZPlatformConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZPlatformConfiguration
type IVZPlatformConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p VZPlatformConfiguration) Init() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VZPlatformConfiguration) Autorelease() VZPlatformConfiguration {
	rv := objc.Send[VZPlatformConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPlatformConfiguration creates a new VZPlatformConfiguration instance.
func NewVZPlatformConfiguration() VZPlatformConfiguration {
	class := getVZPlatformConfigurationClass()
	rv := objc.Send[VZPlatformConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

