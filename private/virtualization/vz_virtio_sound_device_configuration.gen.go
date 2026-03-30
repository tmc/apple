// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSoundDeviceConfiguration] class.
var (
	_VZVirtioSoundDeviceConfigurationClass     VZVirtioSoundDeviceConfigurationClass
	_VZVirtioSoundDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceConfigurationClass() VZVirtioSoundDeviceConfigurationClass {
	_VZVirtioSoundDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceConfigurationClass = VZVirtioSoundDeviceConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceConfiguration")}
	})
	return _VZVirtioSoundDeviceConfigurationClass
}

// GetVZVirtioSoundDeviceConfigurationClass returns the class object for VZVirtioSoundDeviceConfiguration.
func GetVZVirtioSoundDeviceConfigurationClass() VZVirtioSoundDeviceConfigurationClass {
	return getVZVirtioSoundDeviceConfigurationClass()
}

type VZVirtioSoundDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceConfigurationClass) Alloc() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration
type VZVirtioSoundDeviceConfiguration struct {
	VZAudioDeviceConfiguration
}

// VZVirtioSoundDeviceConfigurationFromID constructs a [VZVirtioSoundDeviceConfiguration] from an objc.ID.
func VZVirtioSoundDeviceConfigurationFromID(id objc.ID) VZVirtioSoundDeviceConfiguration {
	return VZVirtioSoundDeviceConfiguration{VZAudioDeviceConfiguration: VZAudioDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioSoundDeviceConfiguration implements IVZVirtioSoundDeviceConfiguration.
var _ IVZVirtioSoundDeviceConfiguration = VZVirtioSoundDeviceConfiguration{}

// An interface definition for the [VZVirtioSoundDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration
type IVZVirtioSoundDeviceConfiguration interface {
	IVZAudioDeviceConfiguration
}

// Init initializes the instance.
func (v VZVirtioSoundDeviceConfiguration) Init() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceConfiguration) Autorelease() VZVirtioSoundDeviceConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceConfiguration creates a new VZVirtioSoundDeviceConfiguration instance.
func NewVZVirtioSoundDeviceConfiguration() VZVirtioSoundDeviceConfiguration {
	class := getVZVirtioSoundDeviceConfigurationClass()
	rv := objc.Send[VZVirtioSoundDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
