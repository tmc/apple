// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGraphicsDeviceConfiguration] class.
var (
	_VZGraphicsDeviceConfigurationClass     VZGraphicsDeviceConfigurationClass
	_VZGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZGraphicsDeviceConfigurationClass() VZGraphicsDeviceConfigurationClass {
	_VZGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZGraphicsDeviceConfigurationClass = VZGraphicsDeviceConfigurationClass{class: objc.GetClass("VZGraphicsDeviceConfiguration")}
	})
	return _VZGraphicsDeviceConfigurationClass
}

// GetVZGraphicsDeviceConfigurationClass returns the class object for VZGraphicsDeviceConfiguration.
func GetVZGraphicsDeviceConfigurationClass() VZGraphicsDeviceConfigurationClass {
	return getVZGraphicsDeviceConfigurationClass()
}

type VZGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGraphicsDeviceConfigurationClass) Alloc() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a graphics device configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration
type VZGraphicsDeviceConfiguration struct {
	objectivec.Object
}

// VZGraphicsDeviceConfigurationFromID constructs a [VZGraphicsDeviceConfiguration] from an objc.ID.
//
// The base class for a graphics device configuration.
func VZGraphicsDeviceConfigurationFromID(id objc.ID) VZGraphicsDeviceConfiguration {
	return VZGraphicsDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZGraphicsDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZGraphicsDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZGraphicsDeviceConfiguration
type IVZGraphicsDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (g VZGraphicsDeviceConfiguration) Init() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g VZGraphicsDeviceConfiguration) Autorelease() VZGraphicsDeviceConfiguration {
	rv := objc.Send[VZGraphicsDeviceConfiguration](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGraphicsDeviceConfiguration creates a new VZGraphicsDeviceConfiguration instance.
func NewVZGraphicsDeviceConfiguration() VZGraphicsDeviceConfiguration {
	class := getVZGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

