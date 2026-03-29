// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitTimePitch] class.
var (
	_AVAudioUnitTimePitchClass     AVAudioUnitTimePitchClass
	_AVAudioUnitTimePitchClassOnce sync.Once
)

func getAVAudioUnitTimePitchClass() AVAudioUnitTimePitchClass {
	_AVAudioUnitTimePitchClassOnce.Do(func() {
		_AVAudioUnitTimePitchClass = AVAudioUnitTimePitchClass{class: objc.GetClass("AVAudioUnitTimePitch")}
	})
	return _AVAudioUnitTimePitchClass
}

// GetAVAudioUnitTimePitchClass returns the class object for AVAudioUnitTimePitch.
func GetAVAudioUnitTimePitchClass() AVAudioUnitTimePitchClass {
	return getAVAudioUnitTimePitchClass()
}

type AVAudioUnitTimePitchClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitTimePitchClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitTimePitchClass) Alloc() AVAudioUnitTimePitch {
	rv := objc.Send[AVAudioUnitTimePitch](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides a good-quality playback rate and pitch shifting
// independently of each other.
//
// # Getting and setting time pitch values
//
//   - [AVAudioUnitTimePitch.Overlap]: The amount of overlap between segments of the input audio signal.
//   - [AVAudioUnitTimePitch.SetOverlap]
//   - [AVAudioUnitTimePitch.Pitch]: The amount to use to pitch shift the input signal.
//   - [AVAudioUnitTimePitch.SetPitch]
//   - [AVAudioUnitTimePitch.Rate]: The playback rate of the input signal.
//   - [AVAudioUnitTimePitch.SetRate]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch
type AVAudioUnitTimePitch struct {
	AVAudioUnitTimeEffect
}

// AVAudioUnitTimePitchFromID constructs a [AVAudioUnitTimePitch] from an objc.ID.
//
// An object that provides a good-quality playback rate and pitch shifting
// independently of each other.
func AVAudioUnitTimePitchFromID(id objc.ID) AVAudioUnitTimePitch {
	return AVAudioUnitTimePitch{AVAudioUnitTimeEffect: AVAudioUnitTimeEffectFromID(id)}
}
// NOTE: AVAudioUnitTimePitch adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitTimePitch] class.
//
// # Getting and setting time pitch values
//
//   - [IAVAudioUnitTimePitch.Overlap]: The amount of overlap between segments of the input audio signal.
//   - [IAVAudioUnitTimePitch.SetOverlap]
//   - [IAVAudioUnitTimePitch.Pitch]: The amount to use to pitch shift the input signal.
//   - [IAVAudioUnitTimePitch.SetPitch]
//   - [IAVAudioUnitTimePitch.Rate]: The playback rate of the input signal.
//   - [IAVAudioUnitTimePitch.SetRate]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch
type IAVAudioUnitTimePitch interface {
	IAVAudioUnitTimeEffect

	// Topic: Getting and setting time pitch values

	// The amount of overlap between segments of the input audio signal.
	Overlap() float32
	SetOverlap(value float32)
	// The amount to use to pitch shift the input signal.
	Pitch() float32
	SetPitch(value float32)
	// The playback rate of the input signal.
	Rate() float32
	SetRate(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitTimePitch) Init() AVAudioUnitTimePitch {
	rv := objc.Send[AVAudioUnitTimePitch](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitTimePitch) Autorelease() AVAudioUnitTimePitch {
	rv := objc.Send[AVAudioUnitTimePitch](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitTimePitch creates a new AVAudioUnitTimePitch instance.
func NewAVAudioUnitTimePitch() AVAudioUnitTimePitch {
	class := getAVAudioUnitTimePitchClass()
	rv := objc.Send[AVAudioUnitTimePitch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a time effect audio unit with the specified description.
//
// audioComponentDescription: The description of the audio unit to create.
//
// # Return Value
// 
// A new [AVAudioUnitTimeEffect] instance.
//
// # Discussion
// 
// The `componentType` field of the description structure must be
// `kAudioUnitType_FormatConverter` (”`aufc`”); otherwise, the method
// raises an exception.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect/init(audioComponentDescription:)
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
func NewAudioUnitTimePitchWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitTimePitch {
	instance := getAVAudioUnitTimePitchClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitTimePitchFromID(rv)
}

// The amount of overlap between segments of the input audio signal.
//
// # Discussion
// 
// A higher value results in fewer artifacts in the output signal. The default
// value is `8.0`. The range of values is `3.0` to `32.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/overlap
func (a AVAudioUnitTimePitch) Overlap() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("overlap"))
	return rv
}
func (a AVAudioUnitTimePitch) SetOverlap(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setOverlap:"), value)
}
// The amount to use to pitch shift the input signal.
//
// # Discussion
// 
// The audio unit measures the pitch in , a logarithmic value you use for
// measuring musical intervals. One octave is equal to 1200 cents. One musical
// semitone is equal to 100 cents.
// 
// The default value is `0``.0`. The range of values is `-2400` to `2400`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/pitch
func (a AVAudioUnitTimePitch) Pitch() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("pitch"))
	return rv
}
func (a AVAudioUnitTimePitch) SetPitch(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPitch:"), value)
}
// The playback rate of the input signal.
//
// # Discussion
// 
// The default value is 1.0. The range of supported values is `1/32` to
// `32.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimePitch/rate
func (a AVAudioUnitTimePitch) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioUnitTimePitch) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}

