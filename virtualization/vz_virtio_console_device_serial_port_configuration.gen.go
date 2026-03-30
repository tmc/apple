// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioConsoleDeviceSerialPortConfiguration] class.
var (
	_VZVirtioConsoleDeviceSerialPortConfigurationClass     VZVirtioConsoleDeviceSerialPortConfigurationClass
	_VZVirtioConsoleDeviceSerialPortConfigurationClassOnce sync.Once
)

func getVZVirtioConsoleDeviceSerialPortConfigurationClass() VZVirtioConsoleDeviceSerialPortConfigurationClass {
	_VZVirtioConsoleDeviceSerialPortConfigurationClassOnce.Do(func() {
		_VZVirtioConsoleDeviceSerialPortConfigurationClass = VZVirtioConsoleDeviceSerialPortConfigurationClass{class: objc.GetClass("VZVirtioConsoleDeviceSerialPortConfiguration")}
	})
	return _VZVirtioConsoleDeviceSerialPortConfigurationClass
}

// GetVZVirtioConsoleDeviceSerialPortConfigurationClass returns the class object for VZVirtioConsoleDeviceSerialPortConfiguration.
func GetVZVirtioConsoleDeviceSerialPortConfigurationClass() VZVirtioConsoleDeviceSerialPortConfigurationClass {
	return getVZVirtioConsoleDeviceSerialPortConfigurationClass()
}

type VZVirtioConsoleDeviceSerialPortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioConsoleDeviceSerialPortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsoleDeviceSerialPortConfigurationClass) Alloc() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that requests the creation of a console device to
// communicate with the guest system.
//
// # Overview
//
// A [VZVirtioConsoleDeviceSerialPortConfiguration] object enables serial
// communication between the guest operating system and host computer through
// the Virtio interface. After you create this configuration object, configure
// its inherited [VZVirtioConsoleDeviceSerialPortConfiguration.Attachment] property with an object that defines the type of
// serial communication you want to enable. Use a
// [VZFileHandleSerialPortAttachment] object to enable two-way communication
// between the guest and host, and use a [VZFileSerialPortAttachment] object
// to enable one-way communication from the guest to the file you designate.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceSerialPortConfiguration
type VZVirtioConsoleDeviceSerialPortConfiguration struct {
	VZSerialPortConfiguration
}

// VZVirtioConsoleDeviceSerialPortConfigurationFromID constructs a [VZVirtioConsoleDeviceSerialPortConfiguration] from an objc.ID.
//
// A configuration object that requests the creation of a console device to
// communicate with the guest system.
func VZVirtioConsoleDeviceSerialPortConfigurationFromID(id objc.ID) VZVirtioConsoleDeviceSerialPortConfiguration {
	return VZVirtioConsoleDeviceSerialPortConfiguration{VZSerialPortConfiguration: VZSerialPortConfigurationFromID(id)}
}

// NOTE: VZVirtioConsoleDeviceSerialPortConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioConsoleDeviceSerialPortConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceSerialPortConfiguration
type IVZVirtioConsoleDeviceSerialPortConfiguration interface {
	IVZSerialPortConfiguration
}

// Init initializes the instance.
func (v VZVirtioConsoleDeviceSerialPortConfiguration) Init() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsoleDeviceSerialPortConfiguration) Autorelease() VZVirtioConsoleDeviceSerialPortConfiguration {
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsoleDeviceSerialPortConfiguration creates a new VZVirtioConsoleDeviceSerialPortConfiguration instance.
func NewVZVirtioConsoleDeviceSerialPortConfiguration() VZVirtioConsoleDeviceSerialPortConfiguration {
	class := getVZVirtioConsoleDeviceSerialPortConfigurationClass()
	rv := objc.Send[VZVirtioConsoleDeviceSerialPortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
