// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMacGraphicsDeviceConfiguration] class.
var (
	_VZMacGraphicsDeviceConfigurationClass     VZMacGraphicsDeviceConfigurationClass
	_VZMacGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZMacGraphicsDeviceConfigurationClass() VZMacGraphicsDeviceConfigurationClass {
	_VZMacGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZMacGraphicsDeviceConfigurationClass = VZMacGraphicsDeviceConfigurationClass{class: objc.GetClass("VZMacGraphicsDeviceConfiguration")}
	})
	return _VZMacGraphicsDeviceConfigurationClass
}

// GetVZMacGraphicsDeviceConfigurationClass returns the class object for VZMacGraphicsDeviceConfiguration.
func GetVZMacGraphicsDeviceConfigurationClass() VZMacGraphicsDeviceConfigurationClass {
	return getVZMacGraphicsDeviceConfigurationClass()
}

type VZMacGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacGraphicsDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacGraphicsDeviceConfigurationClass) Alloc() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Configuration for a display attached to a Mac graphics device.
//
// # Overview
//
// Use this device to attach a display that’s shown in a
// [VZVirtualMachineView].
//
// # Configuring displays
//
//   - [VZMacGraphicsDeviceConfiguration.Displays]: The displays associated with this graphics device.
//   - [VZMacGraphicsDeviceConfiguration.SetDisplays]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration
type VZMacGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZMacGraphicsDeviceConfigurationFromID constructs a [VZMacGraphicsDeviceConfiguration] from an objc.ID.
//
// Configuration for a display attached to a Mac graphics device.
func VZMacGraphicsDeviceConfigurationFromID(id objc.ID) VZMacGraphicsDeviceConfiguration {
	return VZMacGraphicsDeviceConfiguration{VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFromID(id)}
}

// NOTE: VZMacGraphicsDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacGraphicsDeviceConfiguration] class.
//
// # Configuring displays
//
//   - [IVZMacGraphicsDeviceConfiguration.Displays]: The displays associated with this graphics device.
//   - [IVZMacGraphicsDeviceConfiguration.SetDisplays]
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration
type IVZMacGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration

	// Topic: Configuring displays

	// The displays associated with this graphics device.
	Displays() []VZMacGraphicsDisplayConfiguration
	SetDisplays(value []VZMacGraphicsDisplayConfiguration)
}

// Init initializes the instance.
func (m VZMacGraphicsDeviceConfiguration) Init() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacGraphicsDeviceConfiguration) Autorelease() VZMacGraphicsDeviceConfiguration {
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacGraphicsDeviceConfiguration creates a new VZMacGraphicsDeviceConfiguration instance.
func NewVZMacGraphicsDeviceConfiguration() VZMacGraphicsDeviceConfiguration {
	class := getVZMacGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZMacGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The displays associated with this graphics device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacGraphicsDeviceConfiguration/displays
func (m VZMacGraphicsDeviceConfiguration) Displays() []VZMacGraphicsDisplayConfiguration {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("displays"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZMacGraphicsDisplayConfiguration {
		return VZMacGraphicsDisplayConfigurationFromID(id)
	})
}
func (m VZMacGraphicsDeviceConfiguration) SetDisplays(value []VZMacGraphicsDisplayConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setDisplays:"), objectivec.IObjectSliceToNSArray(value))
}
