// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSocketDeviceConfiguration] class.
var (
	_VZVirtioSocketDeviceConfigurationClass     VZVirtioSocketDeviceConfigurationClass
	_VZVirtioSocketDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioSocketDeviceConfigurationClass() VZVirtioSocketDeviceConfigurationClass {
	_VZVirtioSocketDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioSocketDeviceConfigurationClass = VZVirtioSocketDeviceConfigurationClass{class: objc.GetClass("VZVirtioSocketDeviceConfiguration")}
	})
	return _VZVirtioSocketDeviceConfigurationClass
}

// GetVZVirtioSocketDeviceConfigurationClass returns the class object for VZVirtioSocketDeviceConfiguration.
func GetVZVirtioSocketDeviceConfigurationClass() VZVirtioSocketDeviceConfigurationClass {
	return getVZVirtioSocketDeviceConfigurationClass()
}

type VZVirtioSocketDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSocketDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSocketDeviceConfigurationClass) Alloc() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that requests the creation of a socket device to
// communicate with the guest system.
//
// # Overview
//
// Use a [VZVirtioSocketDeviceConfiguration] object to implement port-based
// communication between the guest operating system and the host computer.
// When you add this object to the [VZVirtioSocketDeviceConfiguration.SocketDevices] property of your
// [VZVirtualMachineConfiguration], the virtual machine provides a
// corresponding [VZVirtioSocketDevice] object to use to configure the ports.
// Add only one [VZVirtioSocketDeviceConfiguration] to your virtual
// machine’s configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDeviceConfiguration
type VZVirtioSocketDeviceConfiguration struct {
	VZSocketDeviceConfiguration
}

// VZVirtioSocketDeviceConfigurationFromID constructs a [VZVirtioSocketDeviceConfiguration] from an objc.ID.
//
// A configuration object that requests the creation of a socket device to
// communicate with the guest system.
func VZVirtioSocketDeviceConfigurationFromID(id objc.ID) VZVirtioSocketDeviceConfiguration {
	return VZVirtioSocketDeviceConfiguration{VZSocketDeviceConfiguration: VZSocketDeviceConfigurationFromID(id)}
}

// NOTE: VZVirtioSocketDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioSocketDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSocketDeviceConfiguration
type IVZVirtioSocketDeviceConfiguration interface {
	IVZSocketDeviceConfiguration

	// The socket device that you use to implement port-based communication with the guest operating system.
	SocketDevices() IVZSocketDeviceConfiguration
	SetSocketDevices(value IVZSocketDeviceConfiguration)
}

// Init initializes the instance.
func (v VZVirtioSocketDeviceConfiguration) Init() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSocketDeviceConfiguration) Autorelease() VZVirtioSocketDeviceConfiguration {
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSocketDeviceConfiguration creates a new VZVirtioSocketDeviceConfiguration instance.
func NewVZVirtioSocketDeviceConfiguration() VZVirtioSocketDeviceConfiguration {
	class := getVZVirtioSocketDeviceConfigurationClass()
	rv := objc.Send[VZVirtioSocketDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The socket device that you use to implement port-based communication with
// the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/socketdevices
func (v VZVirtioSocketDeviceConfiguration) SocketDevices() IVZSocketDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("socketDevices"))
	return VZSocketDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioSocketDeviceConfiguration) SetSocketDevices(value IVZSocketDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setSocketDevices:"), value)
}
