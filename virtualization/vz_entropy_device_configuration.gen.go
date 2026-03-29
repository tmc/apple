// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZEntropyDeviceConfiguration] class.
var (
	_VZEntropyDeviceConfigurationClass     VZEntropyDeviceConfigurationClass
	_VZEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZEntropyDeviceConfigurationClass() VZEntropyDeviceConfigurationClass {
	_VZEntropyDeviceConfigurationClassOnce.Do(func() {
		_VZEntropyDeviceConfigurationClass = VZEntropyDeviceConfigurationClass{class: objc.GetClass("VZEntropyDeviceConfiguration")}
	})
	return _VZEntropyDeviceConfigurationClass
}

// GetVZEntropyDeviceConfigurationClass returns the class object for VZEntropyDeviceConfiguration.
func GetVZEntropyDeviceConfigurationClass() VZEntropyDeviceConfigurationClass {
	return getVZEntropyDeviceConfigurationClass()
}

type VZEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZEntropyDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZEntropyDeviceConfigurationClass) Alloc() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The common configuration traits for entropy devices.
//
// # Overview
// 
// Don’t create a VZEntropyDeviceConfiguration object directly. Instead,
// instantiate a subclass such as [VZVirtioEntropyDeviceConfiguration] to
// configure a source of entropy for your virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration
type VZEntropyDeviceConfiguration struct {
	objectivec.Object
}

// VZEntropyDeviceConfigurationFromID constructs a [VZEntropyDeviceConfiguration] from an objc.ID.
//
// The common configuration traits for entropy devices.
func VZEntropyDeviceConfigurationFromID(id objc.ID) VZEntropyDeviceConfiguration {
	return VZEntropyDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZEntropyDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZEntropyDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZEntropyDeviceConfiguration
type IVZEntropyDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (e VZEntropyDeviceConfiguration) Init() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e VZEntropyDeviceConfiguration) Autorelease() VZEntropyDeviceConfiguration {
	rv := objc.Send[VZEntropyDeviceConfiguration](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZEntropyDeviceConfiguration creates a new VZEntropyDeviceConfiguration instance.
func NewVZEntropyDeviceConfiguration() VZEntropyDeviceConfiguration {
	class := getVZEntropyDeviceConfigurationClass()
	rv := objc.Send[VZEntropyDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

