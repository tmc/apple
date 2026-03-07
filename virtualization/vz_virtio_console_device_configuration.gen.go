// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioConsoleDeviceConfiguration] class.
var (
	_VZVirtioConsoleDeviceConfigurationClass     VZVirtioConsoleDeviceConfigurationClass
	_VZVirtioConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceConfigurationClass() VZVirtioConsoleDeviceConfigurationClass {
	_VZVirtioConsoleDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioConsoleDeviceConfigurationClass = VZVirtioConsoleDeviceConfigurationClass{class: objc.GetClass("VZVirtioConsoleDeviceConfiguration")}
	})
	return _VZVirtioConsoleDeviceConfigurationClass
}

// GetVZVirtioConsoleDeviceConfigurationClass returns the class object for VZVirtioConsoleDeviceConfiguration.
func GetVZVirtioConsoleDeviceConfigurationClass() VZVirtioConsoleDeviceConfigurationClass {
	return getVZVirtioConsoleDeviceConfigurationClass()
}

type VZVirtioConsoleDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsoleDeviceConfigurationClass) Alloc() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// A console device that enables communication between the host and the guest
// using console ports through a Virtio interface.
//
// # Overview
// 
// A [VZVirtioConsoleDeviceConfiguration] object enables serial communication
// between the guest-operating system and host computer through the Virtio
// interface. The device sets up one or more ports through
// [VZVirtioConsolePortConfiguration] on the Virtio console device.
//
// # Configuring the console ports
//
//   - [VZVirtioConsoleDeviceConfiguration.Ports]: The list of Virtio port configurations.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration
type VZVirtioConsoleDeviceConfiguration struct {
	VZConsoleDeviceConfiguration
}

// VZVirtioConsoleDeviceConfigurationFromID constructs a [VZVirtioConsoleDeviceConfiguration] from an objc.ID.
//
// A console device that enables communication between the host and the guest
// using console ports through a Virtio interface.
func VZVirtioConsoleDeviceConfigurationFromID(id objc.ID) VZVirtioConsoleDeviceConfiguration {
	return VZVirtioConsoleDeviceConfiguration{VZConsoleDeviceConfiguration: VZConsoleDeviceConfigurationFromID(id)}
}
// NOTE: VZVirtioConsoleDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioConsoleDeviceConfiguration] class.
//
// # Configuring the console ports
//
//   - [IVZVirtioConsoleDeviceConfiguration.Ports]: The list of Virtio port configurations.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration
type IVZVirtioConsoleDeviceConfiguration interface {
	IVZConsoleDeviceConfiguration

	// Topic: Configuring the console ports

	// The list of Virtio port configurations.
	Ports() IVZVirtioConsolePortConfigurationArray

	// The array of console devices that you expose to the guest operating system.
	ConsoleDevices() IVZConsoleDeviceConfiguration
	SetConsoleDevices(value IVZConsoleDeviceConfiguration)
}





// Init initializes the instance.
func (v VZVirtioConsoleDeviceConfiguration) Init() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsoleDeviceConfiguration) Autorelease() VZVirtioConsoleDeviceConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceConfiguration creates a new VZVirtioConsoleDeviceConfiguration instance.
func NewVZVirtioConsoleDeviceConfiguration() VZVirtioConsoleDeviceConfiguration {
	class := getVZVirtioConsoleDeviceConfigurationClass()
	rv := objc.Send[VZVirtioConsoleDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}






















// The list of Virtio port configurations.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceConfiguration/ports
func (v VZVirtioConsoleDeviceConfiguration) Ports() IVZVirtioConsolePortConfigurationArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("ports"))
	return VZVirtioConsolePortConfigurationArrayFromID(objc.ID(rv))
}



// The array of console devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v VZVirtioConsoleDeviceConfiguration) ConsoleDevices() IVZConsoleDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("consoleDevices"))
	return VZConsoleDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioConsoleDeviceConfiguration) SetConsoleDevices(value IVZConsoleDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setConsoleDevices:"), value)
}
























