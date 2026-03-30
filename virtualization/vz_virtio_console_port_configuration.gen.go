// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioConsolePortConfiguration] class.
var (
	_VZVirtioConsolePortConfigurationClass     VZVirtioConsolePortConfigurationClass
	_VZVirtioConsolePortConfigurationClassOnce sync.Once
)

func getVZVirtioConsolePortConfigurationClass() VZVirtioConsolePortConfigurationClass {
	_VZVirtioConsolePortConfigurationClassOnce.Do(func() {
		_VZVirtioConsolePortConfigurationClass = VZVirtioConsolePortConfigurationClass{class: objc.GetClass("VZVirtioConsolePortConfiguration")}
	})
	return _VZVirtioConsolePortConfigurationClass
}

// GetVZVirtioConsolePortConfigurationClass returns the class object for VZVirtioConsolePortConfiguration.
func GetVZVirtioConsolePortConfigurationClass() VZVirtioConsolePortConfigurationClass {
	return getVZVirtioConsolePortConfigurationClass()
}

type VZVirtioConsolePortConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioConsolePortConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioConsolePortConfigurationClass) Alloc() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents the configuration options you can set on a Virtio
// console port.
//
// # Overview
//
// A console port is a two-way communication channel between a host
// [VZSerialPortAttachment] and a VM console port. A Virtio device can have
// one or more attached console devices. Optionally, you can set a name for a
// console port and also configure a console port that the guest can use as
// the system console.
//
// # Configuring the port
//
//   - [VZVirtioConsolePortConfiguration.IsConsole]: A Boolean value that indicates whether this port is a console.
//   - [VZVirtioConsolePortConfiguration.SetIsConsole]
//   - [VZVirtioConsolePortConfiguration.Name]: The name of the port.
//   - [VZVirtioConsolePortConfiguration.SetName]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration
type VZVirtioConsolePortConfiguration struct {
	VZConsolePortConfiguration
}

// VZVirtioConsolePortConfigurationFromID constructs a [VZVirtioConsolePortConfiguration] from an objc.ID.
//
// A class that represents the configuration options you can set on a Virtio
// console port.
func VZVirtioConsolePortConfigurationFromID(id objc.ID) VZVirtioConsolePortConfiguration {
	return VZVirtioConsolePortConfiguration{VZConsolePortConfiguration: VZConsolePortConfigurationFromID(id)}
}

// NOTE: VZVirtioConsolePortConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioConsolePortConfiguration] class.
//
// # Configuring the port
//
//   - [IVZVirtioConsolePortConfiguration.IsConsole]: A Boolean value that indicates whether this port is a console.
//   - [IVZVirtioConsolePortConfiguration.SetIsConsole]
//   - [IVZVirtioConsolePortConfiguration.Name]: The name of the port.
//   - [IVZVirtioConsolePortConfiguration.SetName]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration
type IVZVirtioConsolePortConfiguration interface {
	IVZConsolePortConfiguration

	// Topic: Configuring the port

	// A Boolean value that indicates whether this port is a console.
	IsConsole() bool
	SetIsConsole(value bool)
	// The name of the port.
	Name() string
	SetName(value string)

	// The array of console devices that you expose to the guest operating system.
	ConsoleDevices() IVZConsoleDeviceConfiguration
	SetConsoleDevices(value IVZConsoleDeviceConfiguration)
}

// Init initializes the instance.
func (v VZVirtioConsolePortConfiguration) Init() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioConsolePortConfiguration) Autorelease() VZVirtioConsolePortConfiguration {
	rv := objc.Send[VZVirtioConsolePortConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioConsolePortConfiguration creates a new VZVirtioConsolePortConfiguration instance.
func NewVZVirtioConsolePortConfiguration() VZVirtioConsolePortConfiguration {
	class := getVZVirtioConsolePortConfigurationClass()
	rv := objc.Send[VZVirtioConsolePortConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether this port is a console.
//
// # Discussion
//
// The framework may mark the console port for use as the system console. The
// default is `false`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/isConsole
func (v VZVirtioConsolePortConfiguration) IsConsole() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isConsole"))
	return rv
}
func (v VZVirtioConsolePortConfiguration) SetIsConsole(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIsConsole:"), value)
}

// The name of the port.
//
// # Discussion
//
// The default behavior is to not use a name unless set.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsolePortConfiguration/name
func (v VZVirtioConsolePortConfiguration) Name() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZVirtioConsolePortConfiguration) SetName(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setName:"), objc.String(value))
}

// The array of console devices that you expose to the guest operating system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/consoledevices
func (v VZVirtioConsolePortConfiguration) ConsoleDevices() IVZConsoleDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("consoleDevices"))
	return VZConsoleDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioConsolePortConfiguration) SetConsoleDevices(value IVZConsoleDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setConsoleDevices:"), value)
}
