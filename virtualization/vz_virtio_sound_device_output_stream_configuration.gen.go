// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
var (
	_VZVirtioSoundDeviceOutputStreamConfigurationClass     VZVirtioSoundDeviceOutputStreamConfigurationClass
	_VZVirtioSoundDeviceOutputStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceOutputStreamConfigurationClass() VZVirtioSoundDeviceOutputStreamConfigurationClass {
	_VZVirtioSoundDeviceOutputStreamConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceOutputStreamConfigurationClass = VZVirtioSoundDeviceOutputStreamConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceOutputStreamConfiguration")}
	})
	return _VZVirtioSoundDeviceOutputStreamConfigurationClass
}

// GetVZVirtioSoundDeviceOutputStreamConfigurationClass returns the class object for VZVirtioSoundDeviceOutputStreamConfiguration.
func GetVZVirtioSoundDeviceOutputStreamConfigurationClass() VZVirtioSoundDeviceOutputStreamConfigurationClass {
	return getVZVirtioSoundDeviceOutputStreamConfigurationClass()
}

type VZVirtioSoundDeviceOutputStreamConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceOutputStreamConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceOutputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines a Virtio sound device output stream configuration.
//
// # Overview
// 
// A PCM stream of output audio data, such as to a speaker.
//
// # Accessing the sound sink
//
//   - [VZVirtioSoundDeviceOutputStreamConfiguration.Sink]: An audio stream sink that defines how the host handles audio data produced by the guest.
//   - [VZVirtioSoundDeviceOutputStreamConfiguration.SetSink]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration
type VZVirtioSoundDeviceOutputStreamConfiguration struct {
	VZVirtioSoundDeviceStreamConfiguration
}

// VZVirtioSoundDeviceOutputStreamConfigurationFromID constructs a [VZVirtioSoundDeviceOutputStreamConfiguration] from an objc.ID.
//
// An object that defines a Virtio sound device output stream configuration.
func VZVirtioSoundDeviceOutputStreamConfigurationFromID(id objc.ID) VZVirtioSoundDeviceOutputStreamConfiguration {
	return VZVirtioSoundDeviceOutputStreamConfiguration{VZVirtioSoundDeviceStreamConfiguration: VZVirtioSoundDeviceStreamConfigurationFromID(id)}
}
// NOTE: VZVirtioSoundDeviceOutputStreamConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
//
// # Accessing the sound sink
//
//   - [IVZVirtioSoundDeviceOutputStreamConfiguration.Sink]: An audio stream sink that defines how the host handles audio data produced by the guest.
//   - [IVZVirtioSoundDeviceOutputStreamConfiguration.SetSink]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration
type IVZVirtioSoundDeviceOutputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration

	// Topic: Accessing the sound sink

	// An audio stream sink that defines how the host handles audio data produced by the guest.
	Sink() IVZAudioOutputStreamSink
	SetSink(value IVZAudioOutputStreamSink)
}

// Init initializes the instance.
func (v VZVirtioSoundDeviceOutputStreamConfiguration) Init() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceOutputStreamConfiguration) Autorelease() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceOutputStreamConfiguration creates a new VZVirtioSoundDeviceOutputStreamConfiguration instance.
func NewVZVirtioSoundDeviceOutputStreamConfiguration() VZVirtioSoundDeviceOutputStreamConfiguration {
	class := getVZVirtioSoundDeviceOutputStreamConfigurationClass()
	rv := objc.Send[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An audio stream sink that defines how the host handles audio data produced
// by the guest.
//
// # Discussion
// 
// Not specifying a sink results in a default handler that suppresses the
// audio. The default is `nil`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceOutputStreamConfiguration/sink
func (v VZVirtioSoundDeviceOutputStreamConfiguration) Sink() IVZAudioOutputStreamSink {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("sink"))
	return VZAudioOutputStreamSinkFromID(objc.ID(rv))
}
func (v VZVirtioSoundDeviceOutputStreamConfiguration) SetSink(value IVZAudioOutputStreamSink) {
	objc.Send[struct{}](v.ID, objc.Sel("setSink:"), value)
}

