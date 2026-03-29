// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZConsoleDeviceConfiguration] class.
var (
	_VZConsoleDeviceConfigurationClass     VZConsoleDeviceConfigurationClass
	_VZConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZConsoleDeviceConfigurationClass() VZConsoleDeviceConfigurationClass {
	_VZConsoleDeviceConfigurationClassOnce.Do(func() {
		_VZConsoleDeviceConfigurationClass = VZConsoleDeviceConfigurationClass{class: objc.GetClass("VZConsoleDeviceConfiguration")}
	})
	return _VZConsoleDeviceConfigurationClass
}

// GetVZConsoleDeviceConfigurationClass returns the class object for VZConsoleDeviceConfiguration.
func GetVZConsoleDeviceConfigurationClass() VZConsoleDeviceConfigurationClass {
	return getVZConsoleDeviceConfigurationClass()
}

type VZConsoleDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZConsoleDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZConsoleDeviceConfigurationClass) Alloc() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for a console device configuration.
//
// # Overview
// 
// Don’t instantiate VZConsoleDeviceConfiguration directly, instead use one
// of its subclasses like [VZVirtioConsoleDeviceConfiguration] instead.
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration
type VZConsoleDeviceConfiguration struct {
	objectivec.Object
}

// VZConsoleDeviceConfigurationFromID constructs a [VZConsoleDeviceConfiguration] from an objc.ID.
//
// The base class for a console device configuration.
func VZConsoleDeviceConfigurationFromID(id objc.ID) VZConsoleDeviceConfiguration {
	return VZConsoleDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZConsoleDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZConsoleDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration
type IVZConsoleDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (c VZConsoleDeviceConfiguration) Init() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VZConsoleDeviceConfiguration) Autorelease() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDeviceConfiguration creates a new VZConsoleDeviceConfiguration instance.
func NewVZConsoleDeviceConfiguration() VZConsoleDeviceConfiguration {
	class := getVZConsoleDeviceConfigurationClass()
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

