// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitReverb] class.
var (
	_AVAudioUnitReverbClass     AVAudioUnitReverbClass
	_AVAudioUnitReverbClassOnce sync.Once
)

func getAVAudioUnitReverbClass() AVAudioUnitReverbClass {
	_AVAudioUnitReverbClassOnce.Do(func() {
		_AVAudioUnitReverbClass = AVAudioUnitReverbClass{class: objc.GetClass("AVAudioUnitReverb")}
	})
	return _AVAudioUnitReverbClass
}

// GetAVAudioUnitReverbClass returns the class object for AVAudioUnitReverb.
func GetAVAudioUnitReverbClass() AVAudioUnitReverbClass {
	return getAVAudioUnitReverbClass()
}

type AVAudioUnitReverbClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitReverbClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitReverbClass) Alloc() AVAudioUnitReverb {
	rv := objc.Send[AVAudioUnitReverb](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that implements a reverb effect.
//
// # Overview
// 
// A reverb simulates the acoustic characteristics of a particular
// environment. Use the different presets to simulate a particular space and
// blend it in with the original signal using the [AVAudioUnitReverb.WetDryMix] property.
//
// # Configure the reverb
//
//   - [AVAudioUnitReverb.LoadFactoryPreset]: Configures the audio unit as a reverb preset.
//
// # Getting and setting the reverb values
//
//   - [AVAudioUnitReverb.WetDryMix]: The blend of the wet and dry signals.
//   - [AVAudioUnitReverb.SetWetDryMix]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb
type AVAudioUnitReverb struct {
	AVAudioUnitEffect
}

// AVAudioUnitReverbFromID constructs a [AVAudioUnitReverb] from an objc.ID.
//
// An object that implements a reverb effect.
func AVAudioUnitReverbFromID(id objc.ID) AVAudioUnitReverb {
	return AVAudioUnitReverb{AVAudioUnitEffect: AVAudioUnitEffectFromID(id)}
}
// NOTE: AVAudioUnitReverb adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitReverb] class.
//
// # Configure the reverb
//
//   - [IAVAudioUnitReverb.LoadFactoryPreset]: Configures the audio unit as a reverb preset.
//
// # Getting and setting the reverb values
//
//   - [IAVAudioUnitReverb.WetDryMix]: The blend of the wet and dry signals.
//   - [IAVAudioUnitReverb.SetWetDryMix]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb
type IAVAudioUnitReverb interface {
	IAVAudioUnitEffect

	// Topic: Configure the reverb

	// Configures the audio unit as a reverb preset.
	LoadFactoryPreset(preset AVAudioUnitReverbPreset)

	// Topic: Getting and setting the reverb values

	// The blend of the wet and dry signals.
	WetDryMix() float32
	SetWetDryMix(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitReverb) Init() AVAudioUnitReverb {
	rv := objc.Send[AVAudioUnitReverb](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitReverb) Autorelease() AVAudioUnitReverb {
	rv := objc.Send[AVAudioUnitReverb](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitReverb creates a new AVAudioUnitReverb instance.
func NewAVAudioUnitReverb() AVAudioUnitReverb {
	class := getAVAudioUnitReverbClass()
	rv := objc.Send[AVAudioUnitReverb](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio unit effect object with the specified description.
//
// audioComponentDescription: The description of the audio unit to create.
// 
// The `audioComponentDescription` must be one of these types
// `kAudioUnitType_Effect`, `kAudioUnitType_MusicEffect`,
// `kAudioUnitType_Panner`, `kAudioUnitType_RemoteEffect`, or
// `kAudioUnitType_RemoteMusicEffect`.
//
// # Return Value
// 
// A new [AVAudioUnitEffect] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect/init(audioComponentDescription:)
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
func NewAudioUnitReverbWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitReverb {
	instance := getAVAudioUnitReverbClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitReverbFromID(rv)
}

// Configures the audio unit as a reverb preset.
//
// preset: The reverb preset.
//
// # Discussion
// 
// For more information about possible values, see [AVAudioUnitReverbPreset].
// The default value is [AudioUnitReverbPresetMediumHall].
//
// [AVAudioUnitReverbPreset]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverbPreset
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb/loadFactoryPreset(_:)
func (a AVAudioUnitReverb) LoadFactoryPreset(preset AVAudioUnitReverbPreset) {
	objc.Send[objc.ID](a.ID, objc.Sel("loadFactoryPreset:"), preset)
}

// The blend of the wet and dry signals.
//
// # Discussion
// 
// You specify the blend as a percentage. The range is `0%` through `100%`,
// where `0%` represents all dry.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitReverb/wetDryMix
func (a AVAudioUnitReverb) WetDryMix() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("wetDryMix"))
	return rv
}
func (a AVAudioUnitReverb) SetWetDryMix(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setWetDryMix:"), value)
}

