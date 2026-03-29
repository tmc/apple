// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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

// An object that defines a Virtio sound device configuration.
//
// # Overview
// 
// Use a [VZVirtioSoundDeviceConfiguration] object to configure an audio
// device for your VM. After creating this object, assign appropriate values
// to the [VZVirtioSoundDeviceConfiguration.Streams] array property which defines the behaviors of the
// underlying audio streams for this audio device.
// 
// After creating and configuring a [VZVirtioSoundDeviceConfiguration] object,
// assign it to the [VZVirtioSoundDeviceConfiguration.AudioDevices] property of your VM’s configuration.
//
// # Accessing the sound streams
//
//   - [VZVirtioSoundDeviceConfiguration.Streams]: List of audio streams exposed by this device.
//   - [VZVirtioSoundDeviceConfiguration.SetStreams]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration
type VZVirtioSoundDeviceConfiguration struct {
	VZAudioDeviceConfiguration
}

// VZVirtioSoundDeviceConfigurationFromID constructs a [VZVirtioSoundDeviceConfiguration] from an objc.ID.
//
// An object that defines a Virtio sound device configuration.
func VZVirtioSoundDeviceConfigurationFromID(id objc.ID) VZVirtioSoundDeviceConfiguration {
	return VZVirtioSoundDeviceConfiguration{VZAudioDeviceConfiguration: VZAudioDeviceConfigurationFromID(id)}
}
// NOTE: VZVirtioSoundDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioSoundDeviceConfiguration] class.
//
// # Accessing the sound streams
//
//   - [IVZVirtioSoundDeviceConfiguration.Streams]: List of audio streams exposed by this device.
//   - [IVZVirtioSoundDeviceConfiguration.SetStreams]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration
type IVZVirtioSoundDeviceConfiguration interface {
	IVZAudioDeviceConfiguration

	// Topic: Accessing the sound streams

	// List of audio streams exposed by this device.
	Streams() []VZVirtioSoundDeviceStreamConfiguration
	SetStreams(value []VZVirtioSoundDeviceStreamConfiguration)

	// The list of audio devices.
	AudioDevices() IVZAudioDeviceConfiguration
	SetAudioDevices(value IVZAudioDeviceConfiguration)
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

// List of audio streams exposed by this device.
//
// # Discussion
// 
// Empty by default.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceConfiguration/streams
func (v VZVirtioSoundDeviceConfiguration) Streams() []VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("streams"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZVirtioSoundDeviceStreamConfiguration {
		return VZVirtioSoundDeviceStreamConfigurationFromID(id)
	})
}
func (v VZVirtioSoundDeviceConfiguration) SetStreams(value []VZVirtioSoundDeviceStreamConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setStreams:"), objectivec.IObjectSliceToNSArray(value))
}
// The list of audio devices.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachineconfiguration/audiodevices
func (v VZVirtioSoundDeviceConfiguration) AudioDevices() IVZAudioDeviceConfiguration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("audioDevices"))
	return VZAudioDeviceConfigurationFromID(objc.ID(rv))
}
func (v VZVirtioSoundDeviceConfiguration) SetAudioDevices(value IVZAudioDeviceConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setAudioDevices:"), value)
}

