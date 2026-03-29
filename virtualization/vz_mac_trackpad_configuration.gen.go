// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacTrackpadConfiguration] class.
var (
	_VZMacTrackpadConfigurationClass     VZMacTrackpadConfigurationClass
	_VZMacTrackpadConfigurationClassOnce sync.Once
)

func getVZMacTrackpadConfigurationClass() VZMacTrackpadConfigurationClass {
	_VZMacTrackpadConfigurationClassOnce.Do(func() {
		_VZMacTrackpadConfigurationClass = VZMacTrackpadConfigurationClass{class: objc.GetClass("VZMacTrackpadConfiguration")}
	})
	return _VZMacTrackpadConfigurationClass
}

// GetVZMacTrackpadConfigurationClass returns the class object for VZMacTrackpadConfiguration.
func GetVZMacTrackpadConfigurationClass() VZMacTrackpadConfigurationClass {
	return getVZMacTrackpadConfigurationClass()
}

type VZMacTrackpadConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacTrackpadConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacTrackpadConfigurationClass) Alloc() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The class that represents the configuration for a Mac trackpad.
//
// # Overview
// 
// The [VZVirtualMachineView] uses this device to send pointer events and
// multi-touch trackpad gestures to the virtual machine. In macOS 13 and
// later, guests use the multi-touch trackpad device, while earlier versions
// of macOS uses the USB pointing device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacTrackpadConfiguration
type VZMacTrackpadConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZMacTrackpadConfigurationFromID constructs a [VZMacTrackpadConfiguration] from an objc.ID.
//
// The class that represents the configuration for a Mac trackpad.
func VZMacTrackpadConfigurationFromID(id objc.ID) VZMacTrackpadConfiguration {
	return VZMacTrackpadConfiguration{VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFromID(id)}
}
// NOTE: VZMacTrackpadConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZMacTrackpadConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacTrackpadConfiguration
type IVZMacTrackpadConfiguration interface {
	IVZPointingDeviceConfiguration

	// The list of pointing devices.
	PointingDevices() IVZPointingDeviceConfiguration
	SetPointingDevices(value IVZPointingDeviceConfiguration)
}

// Init initializes the instance.
func (m VZMacTrackpadConfiguration) Init() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m VZMacTrackpadConfiguration) Autorelease() VZMacTrackpadConfiguration {
	rv := objc.Send[VZMacTrackpadConfiguration](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacTrackpadConfiguration creates a new VZMacTrackpadConfiguration instance.
func NewVZMacTrackpadConfiguration() VZMacTrackpadConfiguration {
	class := getVZMacTrackpadConfigurationClass()
	rv := objc.Send[VZMacTrackpadConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The list of pointing devices.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/pointingdevices
func (m VZMacTrackpadConfiguration) PointingDevices() IVZPointingDeviceConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pointingDevices"))
	return VZPointingDeviceConfigurationFromID(objc.ID(rv))
}
func (m VZMacTrackpadConfiguration) SetPointingDevices(value IVZPointingDeviceConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setPointingDevices:"), value)
}

