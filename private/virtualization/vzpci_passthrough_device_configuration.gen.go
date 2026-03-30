// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZPCIPassthroughDeviceConfiguration] class.
var (
	_VZPCIPassthroughDeviceConfigurationClass     VZPCIPassthroughDeviceConfigurationClass
	_VZPCIPassthroughDeviceConfigurationClassOnce sync.Once
)

func getVZPCIPassthroughDeviceConfigurationClass() VZPCIPassthroughDeviceConfigurationClass {
	_VZPCIPassthroughDeviceConfigurationClassOnce.Do(func() {
		_VZPCIPassthroughDeviceConfigurationClass = VZPCIPassthroughDeviceConfigurationClass{class: objc.GetClass("_VZPCIPassthroughDeviceConfiguration")}
	})
	return _VZPCIPassthroughDeviceConfigurationClass
}

// GetVZPCIPassthroughDeviceConfigurationClass returns the class object for _VZPCIPassthroughDeviceConfiguration.
func GetVZPCIPassthroughDeviceConfigurationClass() VZPCIPassthroughDeviceConfigurationClass {
	return getVZPCIPassthroughDeviceConfigurationClass()
}

type VZPCIPassthroughDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPCIPassthroughDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPCIPassthroughDeviceConfigurationClass) Alloc() VZPCIPassthroughDeviceConfiguration {
	rv := objc.Send[VZPCIPassthroughDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZPCIPassthroughDeviceConfiguration.InitWithDomainBusDeviceFunction]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIPassthroughDeviceConfiguration
type VZPCIPassthroughDeviceConfiguration struct {
	VZPCIDeviceConfiguration
}

// VZPCIPassthroughDeviceConfigurationFromID constructs a [VZPCIPassthroughDeviceConfiguration] from an objc.ID.
func VZPCIPassthroughDeviceConfigurationFromID(id objc.ID) VZPCIPassthroughDeviceConfiguration {
	return VZPCIPassthroughDeviceConfiguration{VZPCIDeviceConfiguration: VZPCIDeviceConfigurationFromID(id)}
}

// Ensure VZPCIPassthroughDeviceConfiguration implements IVZPCIPassthroughDeviceConfiguration.
var _ IVZPCIPassthroughDeviceConfiguration = VZPCIPassthroughDeviceConfiguration{}

// An interface definition for the [VZPCIPassthroughDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZPCIPassthroughDeviceConfiguration.InitWithDomainBusDeviceFunction]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPCIPassthroughDeviceConfiguration
type IVZPCIPassthroughDeviceConfiguration interface {
	IVZPCIDeviceConfiguration

	// Topic: Methods

	InitWithDomainBusDeviceFunction(domain uint32, bus byte, device byte, function byte) VZPCIPassthroughDeviceConfiguration
}

// Init initializes the instance.
func (v VZPCIPassthroughDeviceConfiguration) Init() VZPCIPassthroughDeviceConfiguration {
	rv := objc.Send[VZPCIPassthroughDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPCIPassthroughDeviceConfiguration) Autorelease() VZPCIPassthroughDeviceConfiguration {
	rv := objc.Send[VZPCIPassthroughDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPCIPassthroughDeviceConfiguration creates a new VZPCIPassthroughDeviceConfiguration instance.
func NewVZPCIPassthroughDeviceConfiguration() VZPCIPassthroughDeviceConfiguration {
	class := getVZPCIPassthroughDeviceConfigurationClass()
	rv := objc.Send[VZPCIPassthroughDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPCIPassthroughDeviceConfiguration/initWithDomain:bus:device:function:
func NewVZPCIPassthroughDeviceConfigurationWithDomainBusDeviceFunction(domain uint32, bus byte, device byte, function byte) VZPCIPassthroughDeviceConfiguration {
	instance := getVZPCIPassthroughDeviceConfigurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDomain:bus:device:function:"), domain, bus, device, function)
	return VZPCIPassthroughDeviceConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPCIPassthroughDeviceConfiguration/initWithDomain:bus:device:function:
func (v VZPCIPassthroughDeviceConfiguration) InitWithDomainBusDeviceFunction(domain uint32, bus byte, device byte, function byte) VZPCIPassthroughDeviceConfiguration {
	rv := objc.Send[VZPCIPassthroughDeviceConfiguration](v.ID, objc.Sel("initWithDomain:bus:device:function:"), domain, bus, device, function)
	return rv
}
