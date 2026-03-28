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

// See: https://developer.apple.com/documentation/Virtualization/VZMacTrackpadConfiguration
type VZMacTrackpadConfiguration struct {
	VZPointingDeviceConfiguration
}

// VZMacTrackpadConfigurationFromID constructs a [VZMacTrackpadConfiguration] from an objc.ID.
func VZMacTrackpadConfigurationFromID(id objc.ID) VZMacTrackpadConfiguration {
	return VZMacTrackpadConfiguration{VZPointingDeviceConfiguration: VZPointingDeviceConfigurationFromID(id)}
}
// Ensure VZMacTrackpadConfiguration implements IVZMacTrackpadConfiguration.
var _ IVZMacTrackpadConfiguration = VZMacTrackpadConfiguration{}

// An interface definition for the [VZMacTrackpadConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZMacTrackpadConfiguration
type IVZMacTrackpadConfiguration interface {
	IVZPointingDeviceConfiguration
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

