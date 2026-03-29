// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitVarispeed] class.
var (
	_AVAudioUnitVarispeedClass     AVAudioUnitVarispeedClass
	_AVAudioUnitVarispeedClassOnce sync.Once
)

func getAVAudioUnitVarispeedClass() AVAudioUnitVarispeedClass {
	_AVAudioUnitVarispeedClassOnce.Do(func() {
		_AVAudioUnitVarispeedClass = AVAudioUnitVarispeedClass{class: objc.GetClass("AVAudioUnitVarispeed")}
	})
	return _AVAudioUnitVarispeedClass
}

// GetAVAudioUnitVarispeedClass returns the class object for AVAudioUnitVarispeed.
func GetAVAudioUnitVarispeedClass() AVAudioUnitVarispeedClass {
	return getAVAudioUnitVarispeedClass()
}

type AVAudioUnitVarispeedClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitVarispeedClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitVarispeedClass) Alloc() AVAudioUnitVarispeed {
	rv := objc.Send[AVAudioUnitVarispeed](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that allows control of the playback rate.
//
// # Getting and setting the playback rate
//
//   - [AVAudioUnitVarispeed.Rate]: The audio playback rate.
//   - [AVAudioUnitVarispeed.SetRate]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed
type AVAudioUnitVarispeed struct {
	AVAudioUnitTimeEffect
}

// AVAudioUnitVarispeedFromID constructs a [AVAudioUnitVarispeed] from an objc.ID.
//
// An object that allows control of the playback rate.
func AVAudioUnitVarispeedFromID(id objc.ID) AVAudioUnitVarispeed {
	return AVAudioUnitVarispeed{AVAudioUnitTimeEffect: AVAudioUnitTimeEffectFromID(id)}
}
// NOTE: AVAudioUnitVarispeed adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitVarispeed] class.
//
// # Getting and setting the playback rate
//
//   - [IAVAudioUnitVarispeed.Rate]: The audio playback rate.
//   - [IAVAudioUnitVarispeed.SetRate]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed
type IAVAudioUnitVarispeed interface {
	IAVAudioUnitTimeEffect

	// Topic: Getting and setting the playback rate

	// The audio playback rate.
	Rate() float32
	SetRate(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitVarispeed) Init() AVAudioUnitVarispeed {
	rv := objc.Send[AVAudioUnitVarispeed](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitVarispeed) Autorelease() AVAudioUnitVarispeed {
	rv := objc.Send[AVAudioUnitVarispeed](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitVarispeed creates a new AVAudioUnitVarispeed instance.
func NewAVAudioUnitVarispeed() AVAudioUnitVarispeed {
	class := getAVAudioUnitVarispeedClass()
	rv := objc.Send[AVAudioUnitVarispeed](objc.ID(class.class), objc.Sel("new"))
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
func NewAudioUnitVarispeedWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitVarispeed {
	instance := getAVAudioUnitVarispeedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitVarispeedFromID(rv)
}

// The audio playback rate.
//
// # Discussion
// 
// The varispeed audio unit resamples the input signal, and as a result,
// changing the playback rate also changes the pitch. For example, changing
// the rate to `2.0` results in the output audio playing one octave higher.
// Similarly changing the rate to `0.5`, results in the output audio playing
// one octave lower.
// 
// The audio unit measures the pitch in , a logarithmic value you use for
// measuring musical intervals. One octave is equal to 1200 cents. One musical
// semitone is equal to 100 cents.
// 
// Using the `rate` value you calculate the pitch (in cents) using the formula
// `pitch = 1200.0 * log2(rate)`. Conversely, you calculate the appropriate
// `rate` for a desired pitch with the formula `rate = pow(2, cents/1200.0)`.
// 
// The default value is `1.0`. The range of values is `0.25` to `4.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitVarispeed/rate
func (a AVAudioUnitVarispeed) Rate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rate"))
	return rv
}
func (a AVAudioUnitVarispeed) SetRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRate:"), value)
}

