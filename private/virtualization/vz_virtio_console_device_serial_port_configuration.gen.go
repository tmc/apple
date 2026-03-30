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

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioConsoleDeviceSerialPortConfiguration
type VZVirtioConsoleDeviceSerialPortConfiguration struct {
	VZSerialPortConfiguration
}

// VZVirtioConsoleDeviceSerialPortConfigurationFromID constructs a [VZVirtioConsoleDeviceSerialPortConfiguration] from an objc.ID.
func VZVirtioConsoleDeviceSerialPortConfigurationFromID(id objc.ID) VZVirtioConsoleDeviceSerialPortConfiguration {
	return VZVirtioConsoleDeviceSerialPortConfiguration{VZSerialPortConfiguration: VZSerialPortConfigurationFromID(id)}
}

// Ensure VZVirtioConsoleDeviceSerialPortConfiguration implements IVZVirtioConsoleDeviceSerialPortConfiguration.
var _ IVZVirtioConsoleDeviceSerialPortConfiguration = VZVirtioConsoleDeviceSerialPortConfiguration{}

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
