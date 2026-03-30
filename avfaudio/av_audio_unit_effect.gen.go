// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitEffect] class.
var (
	_AVAudioUnitEffectClass     AVAudioUnitEffectClass
	_AVAudioUnitEffectClassOnce sync.Once
)

func getAVAudioUnitEffectClass() AVAudioUnitEffectClass {
	_AVAudioUnitEffectClassOnce.Do(func() {
		_AVAudioUnitEffectClass = AVAudioUnitEffectClass{class: objc.GetClass("AVAudioUnitEffect")}
	})
	return _AVAudioUnitEffectClass
}

// GetAVAudioUnitEffectClass returns the class object for AVAudioUnitEffect.
func GetAVAudioUnitEffectClass() AVAudioUnitEffectClass {
	return getAVAudioUnitEffectClass()
}

type AVAudioUnitEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitEffectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitEffectClass) Alloc() AVAudioUnitEffect {
	rv := objc.Send[AVAudioUnitEffect](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that processes audio in real time.
//
// # Overview
//
// This processing uses [AudioUnit] of type effect, music effect, panner,
// remote effect, or remote music effect. These effects run in real time and
// process some number of audio input samples to produce several audio output
// samples. A delay unit is an example of an effect unit.
//
// # Creating an audio effect
//
//   - [AVAudioUnitEffect.InitWithAudioComponentDescription]: Creates an audio unit effect object with the specified description.
//
// # Getting the bypass state
//
//   - [AVAudioUnitEffect.Bypass]: The bypass state of the audio unit.
//   - [AVAudioUnitEffect.SetBypass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect
//
// [AudioUnit]: https://developer.apple.com/documentation/AudioToolbox/AudioUnit
type AVAudioUnitEffect struct {
	AVAudioUnit
}

// AVAudioUnitEffectFromID constructs a [AVAudioUnitEffect] from an objc.ID.
//
// An object that processes audio in real time.
func AVAudioUnitEffectFromID(id objc.ID) AVAudioUnitEffect {
	return AVAudioUnitEffect{AVAudioUnit: AVAudioUnitFromID(id)}
}

// NOTE: AVAudioUnitEffect adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitEffect] class.
//
// # Creating an audio effect
//
//   - [IAVAudioUnitEffect.InitWithAudioComponentDescription]: Creates an audio unit effect object with the specified description.
//
// # Getting the bypass state
//
//   - [IAVAudioUnitEffect.Bypass]: The bypass state of the audio unit.
//   - [IAVAudioUnitEffect.SetBypass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect
type IAVAudioUnitEffect interface {
	IAVAudioUnit

	// Topic: Creating an audio effect

	// Creates an audio unit effect object with the specified description.
	InitWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitEffect

	// Topic: Getting the bypass state

	// The bypass state of the audio unit.
	Bypass() bool
	SetBypass(value bool)
}

// Init initializes the instance.
func (a AVAudioUnitEffect) Init() AVAudioUnitEffect {
	rv := objc.Send[AVAudioUnitEffect](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitEffect) Autorelease() AVAudioUnitEffect {
	rv := objc.Send[AVAudioUnitEffect](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitEffect creates a new AVAudioUnitEffect instance.
func NewAVAudioUnitEffect() AVAudioUnitEffect {
	class := getAVAudioUnitEffectClass()
	rv := objc.Send[AVAudioUnitEffect](objc.ID(class.class), objc.Sel("new"))
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
func NewAudioUnitEffectWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitEffect {
	instance := getAVAudioUnitEffectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitEffectFromID(rv)
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
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
//
// # Return Value
//
// A new [AVAudioUnitEffect] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect/init(audioComponentDescription:)
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
func (a AVAudioUnitEffect) InitWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitEffect {
	rv := objc.Send[AVAudioUnitEffect](a.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return rv
}

// The bypass state of the audio unit.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEffect/bypass
func (a AVAudioUnitEffect) Bypass() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("bypass"))
	return rv
}
func (a AVAudioUnitEffect) SetBypass(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setBypass:"), value)
}
