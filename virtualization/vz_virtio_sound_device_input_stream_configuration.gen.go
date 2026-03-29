// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
var (
	_VZVirtioSoundDeviceInputStreamConfigurationClass     VZVirtioSoundDeviceInputStreamConfigurationClass
	_VZVirtioSoundDeviceInputStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceInputStreamConfigurationClass() VZVirtioSoundDeviceInputStreamConfigurationClass {
	_VZVirtioSoundDeviceInputStreamConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceInputStreamConfigurationClass = VZVirtioSoundDeviceInputStreamConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceInputStreamConfiguration")}
	})
	return _VZVirtioSoundDeviceInputStreamConfigurationClass
}

// GetVZVirtioSoundDeviceInputStreamConfigurationClass returns the class object for VZVirtioSoundDeviceInputStreamConfiguration.
func GetVZVirtioSoundDeviceInputStreamConfigurationClass() VZVirtioSoundDeviceInputStreamConfigurationClass {
	return getVZVirtioSoundDeviceInputStreamConfigurationClass()
}

type VZVirtioSoundDeviceInputStreamConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceInputStreamConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceInputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A PCM stream of input audio data, such as from a microphone.
//
// # Overview
// 
// This device represents a PCM stream of audio data. Don’t instantiate
// [VZVirtioSoundDeviceStreamConfiguration] directly. Instead, use one of its
// subclasses such as [VZVirtioSoundDeviceInputStreamConfiguration] or
// [VZVirtioSoundDeviceOutputStreamConfiguration].
//
// # Accessing the sound source
//
//   - [VZVirtioSoundDeviceInputStreamConfiguration.Source]: An audio stream source that defines how the host supplies audio data for the guest.
//   - [VZVirtioSoundDeviceInputStreamConfiguration.SetSource]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceInputStreamConfiguration
type VZVirtioSoundDeviceInputStreamConfiguration struct {
	VZVirtioSoundDeviceStreamConfiguration
}

// VZVirtioSoundDeviceInputStreamConfigurationFromID constructs a [VZVirtioSoundDeviceInputStreamConfiguration] from an objc.ID.
//
// A PCM stream of input audio data, such as from a microphone.
func VZVirtioSoundDeviceInputStreamConfigurationFromID(id objc.ID) VZVirtioSoundDeviceInputStreamConfiguration {
	return VZVirtioSoundDeviceInputStreamConfiguration{VZVirtioSoundDeviceStreamConfiguration: VZVirtioSoundDeviceStreamConfigurationFromID(id)}
}
// NOTE: VZVirtioSoundDeviceInputStreamConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
//
// # Accessing the sound source
//
//   - [IVZVirtioSoundDeviceInputStreamConfiguration.Source]: An audio stream source that defines how the host supplies audio data for the guest.
//   - [IVZVirtioSoundDeviceInputStreamConfiguration.SetSource]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceInputStreamConfiguration
type IVZVirtioSoundDeviceInputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration

	// Topic: Accessing the sound source

	// An audio stream source that defines how the host supplies audio data for the guest.
	Source() IVZAudioInputStreamSource
	SetSource(value IVZAudioInputStreamSource)
}

// Init initializes the instance.
func (v VZVirtioSoundDeviceInputStreamConfiguration) Init() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceInputStreamConfiguration) Autorelease() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceInputStreamConfiguration creates a new VZVirtioSoundDeviceInputStreamConfiguration instance.
func NewVZVirtioSoundDeviceInputStreamConfiguration() VZVirtioSoundDeviceInputStreamConfiguration {
	class := getVZVirtioSoundDeviceInputStreamConfigurationClass()
	rv := objc.Send[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An audio stream source that defines how the host supplies audio data for
// the guest.
//
// # Discussion
// 
// Not specifying a source results in a default handler that produces audio
// silence. The default is `nil`.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceInputStreamConfiguration/source
func (v VZVirtioSoundDeviceInputStreamConfiguration) Source() IVZAudioInputStreamSource {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("source"))
	return VZAudioInputStreamSourceFromID(objc.ID(rv))
}
func (v VZVirtioSoundDeviceInputStreamConfiguration) SetSource(value IVZAudioInputStreamSource) {
	objc.Send[struct{}](v.ID, objc.Sel("setSource:"), value)
}

