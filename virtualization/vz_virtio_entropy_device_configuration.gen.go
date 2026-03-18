// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioEntropyDeviceConfiguration] class.
var (
	_VZVirtioEntropyDeviceConfigurationClass     VZVirtioEntropyDeviceConfigurationClass
	_VZVirtioEntropyDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioEntropyDeviceConfigurationClass() VZVirtioEntropyDeviceConfigurationClass {
	_VZVirtioEntropyDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioEntropyDeviceConfigurationClass = VZVirtioEntropyDeviceConfigurationClass{class: objc.GetClass("VZVirtioEntropyDeviceConfiguration")}
	})
	return _VZVirtioEntropyDeviceConfigurationClass
}

// GetVZVirtioEntropyDeviceConfigurationClass returns the class object for VZVirtioEntropyDeviceConfiguration.
func GetVZVirtioEntropyDeviceConfigurationClass() VZVirtioEntropyDeviceConfigurationClass {
	return getVZVirtioEntropyDeviceConfigurationClass()
}

type VZVirtioEntropyDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioEntropyDeviceConfigurationClass) Alloc() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A source of entropy for the guest’s random number generator.
//
// # Overview
// 
// Use a [VZVirtioEntropyDeviceConfiguration] object to expose a source of
// entropy for the guest operating system’s random-number generator. When
// you create this object and add it to your virtual machine’s
// configuration, the virtual machine configures a Virtio-compliant entropy
// device. The guest operating system uses this device as a seed to generate
// random numbers.
// 
// Create a [VZVirtioEntropyDeviceConfiguration] object and add it to the
// [VZVirtioEntropyDeviceConfiguration.EntropyDevices] property of your virtual machine’s configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioEntropyDeviceConfiguration
type VZVirtioEntropyDeviceConfiguration struct {
	VZEntropyDeviceConfiguration
}

// VZVirtioEntropyDeviceConfigurationFromID constructs a [VZVirtioEntropyDeviceConfiguration] from an objc.ID.
//
// A source of entropy for the guest’s random number generator.
func VZVirtioEntropyDeviceConfigurationFromID(id objc.ID) VZVirtioEntropyDeviceConfiguration {
	return VZVirtioEntropyDeviceConfiguration{VZEntropyDeviceConfiguration: VZEntropyDeviceConfigurationFromID(id)}
}
// NOTE: VZVirtioEntropyDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioEntropyDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioEntropyDeviceConfiguration
type IVZVirtioEntropyDeviceConfiguration interface {
	IVZEntropyDeviceConfiguration

	// The array of randomization devices that you expose to the guest operating system.
	EntropyDevices() IVZEntropyDeviceConfiguration
	SetEntropyDevices(value IVZEntropyDeviceConfiguration)
}

// Init initializes the instance.
func (v VZVirtioEntropyDeviceConfiguration) Init() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioEntropyDeviceConfiguration) Autorelease() VZVirtioEntropyDeviceConfiguration {
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioEntropyDeviceConfiguration creates a new VZVirtioEntropyDeviceConfiguration instance.
func NewVZVirtioEntropyDeviceConfiguration() VZVirtioEntropyDeviceConfiguration {
	class := getVZVirtioEntropyDeviceConfigurationClass()
	rv := objc.Send[VZVirtioEntropyDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of randomization devices that you expose to the guest operating
// system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/entropydevices
func (v VZVirtioEntropyDeviceConfiguration) EntropyDevices() IVZEntropyDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("entropyDevices"))
	return VZEntropyDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioEntropyDeviceConfiguration) SetEntropyDevices(value IVZEntropyDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setEntropyDevices:"), value)
}

