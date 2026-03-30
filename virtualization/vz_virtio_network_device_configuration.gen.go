// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioNetworkDeviceConfiguration] class.
var (
	_VZVirtioNetworkDeviceConfigurationClass     VZVirtioNetworkDeviceConfigurationClass
	_VZVirtioNetworkDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioNetworkDeviceConfigurationClass() VZVirtioNetworkDeviceConfigurationClass {
	_VZVirtioNetworkDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioNetworkDeviceConfigurationClass = VZVirtioNetworkDeviceConfigurationClass{class: objc.GetClass("VZVirtioNetworkDeviceConfiguration")}
	})
	return _VZVirtioNetworkDeviceConfigurationClass
}

// GetVZVirtioNetworkDeviceConfigurationClass returns the class object for VZVirtioNetworkDeviceConfiguration.
func GetVZVirtioNetworkDeviceConfigurationClass() VZVirtioNetworkDeviceConfigurationClass {
	return getVZVirtioNetworkDeviceConfigurationClass()
}

type VZVirtioNetworkDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioNetworkDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioNetworkDeviceConfigurationClass) Alloc() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that requests the creation of a network device for
// the guest system.
//
// # Overview
//
// Use a [VZVirtioNetworkDeviceConfiguration] object to configure one network
// interface of your virtual machine. After creating this object, assign an
// appropriate value to its inherited [VZVirtioNetworkDeviceConfiguration.Attachment] property to define the type
// of network interface you want. You can also assign a specific MAC address,
// or let the system generate a random address for you.
//
// After creating and configuring a [VZVirtioNetworkDeviceConfiguration]
// object, assign it to the [VZVirtioNetworkDeviceConfiguration.NetworkDevices] property of your virtual
// machine’s configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioNetworkDeviceConfiguration
type VZVirtioNetworkDeviceConfiguration struct {
	VZNetworkDeviceConfiguration
}

// VZVirtioNetworkDeviceConfigurationFromID constructs a [VZVirtioNetworkDeviceConfiguration] from an objc.ID.
//
// A configuration object that requests the creation of a network device for
// the guest system.
func VZVirtioNetworkDeviceConfigurationFromID(id objc.ID) VZVirtioNetworkDeviceConfiguration {
	return VZVirtioNetworkDeviceConfiguration{VZNetworkDeviceConfiguration: VZNetworkDeviceConfigurationFromID(id)}
}

// NOTE: VZVirtioNetworkDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioNetworkDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioNetworkDeviceConfiguration
type IVZVirtioNetworkDeviceConfiguration interface {
	IVZNetworkDeviceConfiguration

	// The array of network devices that you expose to the guest operating system.
	NetworkDevices() IVZNetworkDeviceConfiguration
	SetNetworkDevices(value IVZNetworkDeviceConfiguration)
}

// Init initializes the instance.
func (v VZVirtioNetworkDeviceConfiguration) Init() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioNetworkDeviceConfiguration) Autorelease() VZVirtioNetworkDeviceConfiguration {
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioNetworkDeviceConfiguration creates a new VZVirtioNetworkDeviceConfiguration instance.
func NewVZVirtioNetworkDeviceConfiguration() VZVirtioNetworkDeviceConfiguration {
	class := getVZVirtioNetworkDeviceConfigurationClass()
	rv := objc.Send[VZVirtioNetworkDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of network devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/networkdevices
func (v VZVirtioNetworkDeviceConfiguration) NetworkDevices() IVZNetworkDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("networkDevices"))
	return VZNetworkDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioNetworkDeviceConfiguration) SetNetworkDevices(value IVZNetworkDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setNetworkDevices:"), value)
}
