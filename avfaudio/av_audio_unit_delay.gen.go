// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitDelay] class.
var (
	_AVAudioUnitDelayClass     AVAudioUnitDelayClass
	_AVAudioUnitDelayClassOnce sync.Once
)

func getAVAudioUnitDelayClass() AVAudioUnitDelayClass {
	_AVAudioUnitDelayClassOnce.Do(func() {
		_AVAudioUnitDelayClass = AVAudioUnitDelayClass{class: objc.GetClass("AVAudioUnitDelay")}
	})
	return _AVAudioUnitDelayClass
}

// GetAVAudioUnitDelayClass returns the class object for AVAudioUnitDelay.
func GetAVAudioUnitDelayClass() AVAudioUnitDelayClass {
	return getAVAudioUnitDelayClass()
}

type AVAudioUnitDelayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitDelayClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitDelayClass) Alloc() AVAudioUnitDelay {
	rv := objc.Send[AVAudioUnitDelay](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that implements a delay effect.
//
// # Overview
//
// A delay unit delays the input signal by the specified time interval and
// then blends it with the input signal. You can also control the amount of
// high-frequency roll-off to simulate the effect of a tape delay.
//
// # Getting and setting the delay values
//
//   - [AVAudioUnitDelay.DelayTime]: The time for the input signal to reach the output.
//   - [AVAudioUnitDelay.SetDelayTime]
//   - [AVAudioUnitDelay.Feedback]: The amount of the output signal that feeds back into the delay line.
//   - [AVAudioUnitDelay.SetFeedback]
//   - [AVAudioUnitDelay.LowPassCutoff]: The cutoff frequency above which high frequency content rolls off, in hertz.
//   - [AVAudioUnitDelay.SetLowPassCutoff]
//   - [AVAudioUnitDelay.WetDryMix]: The blend of the wet and dry signals.
//   - [AVAudioUnitDelay.SetWetDryMix]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay
type AVAudioUnitDelay struct {
	AVAudioUnitEffect
}

// AVAudioUnitDelayFromID constructs a [AVAudioUnitDelay] from an objc.ID.
//
// An object that implements a delay effect.
func AVAudioUnitDelayFromID(id objc.ID) AVAudioUnitDelay {
	return AVAudioUnitDelay{AVAudioUnitEffect: AVAudioUnitEffectFromID(id)}
}

// NOTE: AVAudioUnitDelay adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitDelay] class.
//
// # Getting and setting the delay values
//
//   - [IAVAudioUnitDelay.DelayTime]: The time for the input signal to reach the output.
//   - [IAVAudioUnitDelay.SetDelayTime]
//   - [IAVAudioUnitDelay.Feedback]: The amount of the output signal that feeds back into the delay line.
//   - [IAVAudioUnitDelay.SetFeedback]
//   - [IAVAudioUnitDelay.LowPassCutoff]: The cutoff frequency above which high frequency content rolls off, in hertz.
//   - [IAVAudioUnitDelay.SetLowPassCutoff]
//   - [IAVAudioUnitDelay.WetDryMix]: The blend of the wet and dry signals.
//   - [IAVAudioUnitDelay.SetWetDryMix]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay
type IAVAudioUnitDelay interface {
	IAVAudioUnitEffect

	// Topic: Getting and setting the delay values

	// The time for the input signal to reach the output.
	DelayTime() float64
	SetDelayTime(value float64)
	// The amount of the output signal that feeds back into the delay line.
	Feedback() float32
	SetFeedback(value float32)
	// The cutoff frequency above which high frequency content rolls off, in hertz.
	LowPassCutoff() float32
	SetLowPassCutoff(value float32)
	// The blend of the wet and dry signals.
	WetDryMix() float32
	SetWetDryMix(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitDelay) Init() AVAudioUnitDelay {
	rv := objc.Send[AVAudioUnitDelay](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitDelay) Autorelease() AVAudioUnitDelay {
	rv := objc.Send[AVAudioUnitDelay](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitDelay creates a new AVAudioUnitDelay instance.
func NewAVAudioUnitDelay() AVAudioUnitDelay {
	class := getAVAudioUnitDelayClass()
	rv := objc.Send[AVAudioUnitDelay](objc.ID(class.class), objc.Sel("new"))
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
func NewAudioUnitDelayWithAudioComponentDescription(audioComponentDescription unsafe.Pointer) AVAudioUnitDelay {
	instance := getAVAudioUnitDelayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitDelayFromID(rv)
}

// The time for the input signal to reach the output.
//
// # Discussion
//
// You specify the delay in seconds. The default value is `1`. The valid range
// of values is `0` to `2` seconds.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/delayTime
func (a AVAudioUnitDelay) DelayTime() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("delayTime"))
	return rv
}
func (a AVAudioUnitDelay) SetDelayTime(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelayTime:"), value)
}

// The amount of the output signal that feeds back into the delay line.
//
// # Discussion
//
// You specify the feedback as a percentage. The default value is `50%`. The
// valid range of values is `-100%` to `100%`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/feedback
func (a AVAudioUnitDelay) Feedback() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("feedback"))
	return rv
}
func (a AVAudioUnitDelay) SetFeedback(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setFeedback:"), value)
}

// The cutoff frequency above which high frequency content rolls off, in
// hertz.
//
// # Discussion
//
// The default value is `15000 Hz`. The valid range of values is `10 Hz`
// through `(sampleRate/2)`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/lowPassCutoff
func (a AVAudioUnitDelay) LowPassCutoff() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("lowPassCutoff"))
	return rv
}
func (a AVAudioUnitDelay) SetLowPassCutoff(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setLowPassCutoff:"), value)
}

// The blend of the wet and dry signals.
//
// # Discussion
//
// You specify the blend as a percentage. The default value is `100%`. The
// valid range of values is `0%` through `100%`, where `0%` represents all
// dry.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitDelay/wetDryMix
func (a AVAudioUnitDelay) WetDryMix() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("wetDryMix"))
	return rv
}
func (a AVAudioUnitDelay) SetWetDryMix(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setWetDryMix:"), value)
}
