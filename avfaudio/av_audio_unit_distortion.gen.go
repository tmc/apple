// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitDistortion] class.
var (
	_AVAudioUnitDistortionClass     AVAudioUnitDistortionClass
	_AVAudioUnitDistortionClassOnce sync.Once
)

func getAVAudioUnitDistortionClass() AVAudioUnitDistortionClass {
	_AVAudioUnitDistortionClassOnce.Do(func() {
		_AVAudioUnitDistortionClass = AVAudioUnitDistortionClass{class: objc.GetClass("AVAudioUnitDistortion")}
	})
	return _AVAudioUnitDistortionClass
}

// GetAVAudioUnitDistortionClass returns the class object for AVAudioUnitDistortion.
func GetAVAudioUnitDistortionClass() AVAudioUnitDistortionClass {
	return getAVAudioUnitDistortionClass()
}

type AVAudioUnitDistortionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitDistortionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitDistortionClass) Alloc() AVAudioUnitDistortion {
	rv := objc.Send[AVAudioUnitDistortion](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that implements a multistage distortion effect.
//
// # Configuring the distortion
//
//   - [AVAudioUnitDistortion.LoadFactoryPreset]: Configures the audio distortion unit by loading a distortion preset.
//
// # Getting and setting the distortion values
//
//   - [AVAudioUnitDistortion.PreGain]: The gain that the audio unit applies to the signal before distortion, in decibels.
//   - [AVAudioUnitDistortion.SetPreGain]
//   - [AVAudioUnitDistortion.WetDryMix]: The blend of the distorted and dry signals.
//   - [AVAudioUnitDistortion.SetWetDryMix]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion
type AVAudioUnitDistortion struct {
	AVAudioUnitEffect
}

// AVAudioUnitDistortionFromID constructs a [AVAudioUnitDistortion] from an objc.ID.
//
// An object that implements a multistage distortion effect.
func AVAudioUnitDistortionFromID(id objc.ID) AVAudioUnitDistortion {
	return AVAudioUnitDistortion{AVAudioUnitEffect: AVAudioUnitEffectFromID(id)}
}

// NOTE: AVAudioUnitDistortion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitDistortion] class.
//
// # Configuring the distortion
//
//   - [IAVAudioUnitDistortion.LoadFactoryPreset]: Configures the audio distortion unit by loading a distortion preset.
//
// # Getting and setting the distortion values
//
//   - [IAVAudioUnitDistortion.PreGain]: The gain that the audio unit applies to the signal before distortion, in decibels.
//   - [IAVAudioUnitDistortion.SetPreGain]
//   - [IAVAudioUnitDistortion.WetDryMix]: The blend of the distorted and dry signals.
//   - [IAVAudioUnitDistortion.SetWetDryMix]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion
type IAVAudioUnitDistortion interface {
	IAVAudioUnitEffect

	// Topic: Configuring the distortion

	// Configures the audio distortion unit by loading a distortion preset.
	LoadFactoryPreset(preset AVAudioUnitDistortionPreset)

	// Topic: Getting and setting the distortion values

	// The gain that the audio unit applies to the signal before distortion, in decibels.
	PreGain() float32
	SetPreGain(value float32)
	// The blend of the distorted and dry signals.
	WetDryMix() float32
	SetWetDryMix(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitDistortion) Init() AVAudioUnitDistortion {
	rv := objc.Send[AVAudioUnitDistortion](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitDistortion) Autorelease() AVAudioUnitDistortion {
	rv := objc.Send[AVAudioUnitDistortion](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitDistortion creates a new AVAudioUnitDistortion instance.
func NewAVAudioUnitDistortion() AVAudioUnitDistortion {
	class := getAVAudioUnitDistortionClass()
	rv := objc.Send[AVAudioUnitDistortion](objc.ID(class.class), objc.Sel("new"))
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
func NewAudioUnitDistortionWithAudioComponentDescription(audioComponentDescription unsafe.Pointer) AVAudioUnitDistortion {
	instance := getAVAudioUnitDistortionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitDistortionFromID(rv)
}

// Configures the audio distortion unit by loading a distortion preset.
//
// preset: The distortion preset.
//
// # Discussion
//
// For more information about possible values for `preset`, see
// [AVAudioUnitDistortionPreset]. The default value is
// [AVAudioUnitDistortionPresetDrumsBitBrush].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/loadFactoryPreset(_:)
//
// [AVAudioUnitDistortionPreset]: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortionPreset
func (a AVAudioUnitDistortion) LoadFactoryPreset(preset AVAudioUnitDistortionPreset) {
	objc.Send[objc.ID](a.ID, objc.Sel("loadFactoryPreset:"), preset)
}

// The gain that the audio unit applies to the signal before distortion, in
// decibels.
//
// # Discussion
//
// The default value is `-6 dB`. The valid range of values is `-80 dB` to `20
// dB`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/preGain
func (a AVAudioUnitDistortion) PreGain() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("preGain"))
	return rv
}
func (a AVAudioUnitDistortion) SetPreGain(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreGain:"), value)
}

// The blend of the distorted and dry signals.
//
// # Discussion
//
// You specify the blend as a percentage. The default value is `50%`. The
// valid range is `0%` through `100%`, where `0` represents all dry.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDistortion/wetDryMix
func (a AVAudioUnitDistortion) WetDryMix() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("wetDryMix"))
	return rv
}
func (a AVAudioUnitDistortion) SetWetDryMix(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setWetDryMix:"), value)
}
