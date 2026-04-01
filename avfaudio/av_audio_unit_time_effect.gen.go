// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitTimeEffect] class.
var (
	_AVAudioUnitTimeEffectClass     AVAudioUnitTimeEffectClass
	_AVAudioUnitTimeEffectClassOnce sync.Once
)

func getAVAudioUnitTimeEffectClass() AVAudioUnitTimeEffectClass {
	_AVAudioUnitTimeEffectClassOnce.Do(func() {
		_AVAudioUnitTimeEffectClass = AVAudioUnitTimeEffectClass{class: objc.GetClass("AVAudioUnitTimeEffect")}
	})
	return _AVAudioUnitTimeEffectClass
}

// GetAVAudioUnitTimeEffectClass returns the class object for AVAudioUnitTimeEffect.
func GetAVAudioUnitTimeEffectClass() AVAudioUnitTimeEffectClass {
	return getAVAudioUnitTimeEffectClass()
}

type AVAudioUnitTimeEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitTimeEffectClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitTimeEffectClass) Alloc() AVAudioUnitTimeEffect {
	rv := objc.Send[AVAudioUnitTimeEffect](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that processes audio in nonreal time.
//
// # Overview
//
// A time effect audio unit represents an [AVAudioUnit] with a type
// `kAudioUnitType_FormatConverter` (`aufc)`. These effects don’t process
// audio in real time. The [AVAudioUnitVarispeed] class is an example of a
// time effect unit.
//
// # Creating a time effect
//
//   - [AVAudioUnitTimeEffect.InitWithAudioComponentDescription]: Creates a time effect audio unit with the specified description.
//
// # Getting and setting the time effect
//
//   - [AVAudioUnitTimeEffect.Bypass]: The bypass state of the audio unit.
//   - [AVAudioUnitTimeEffect.SetBypass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect
type AVAudioUnitTimeEffect struct {
	AVAudioUnit
}

// AVAudioUnitTimeEffectFromID constructs a [AVAudioUnitTimeEffect] from an objc.ID.
//
// An object that processes audio in nonreal time.
func AVAudioUnitTimeEffectFromID(id objc.ID) AVAudioUnitTimeEffect {
	return AVAudioUnitTimeEffect{AVAudioUnit: AVAudioUnitFromID(id)}
}

// NOTE: AVAudioUnitTimeEffect adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitTimeEffect] class.
//
// # Creating a time effect
//
//   - [IAVAudioUnitTimeEffect.InitWithAudioComponentDescription]: Creates a time effect audio unit with the specified description.
//
// # Getting and setting the time effect
//
//   - [IAVAudioUnitTimeEffect.Bypass]: The bypass state of the audio unit.
//   - [IAVAudioUnitTimeEffect.SetBypass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect
type IAVAudioUnitTimeEffect interface {
	IAVAudioUnit

	// Topic: Creating a time effect

	// Creates a time effect audio unit with the specified description.
	InitWithAudioComponentDescription(audioComponentDescription unsafe.Pointer) AVAudioUnitTimeEffect

	// Topic: Getting and setting the time effect

	// The bypass state of the audio unit.
	Bypass() bool
	SetBypass(value bool)
}

// Init initializes the instance.
func (a AVAudioUnitTimeEffect) Init() AVAudioUnitTimeEffect {
	rv := objc.Send[AVAudioUnitTimeEffect](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitTimeEffect) Autorelease() AVAudioUnitTimeEffect {
	rv := objc.Send[AVAudioUnitTimeEffect](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitTimeEffect creates a new AVAudioUnitTimeEffect instance.
func NewAVAudioUnitTimeEffect() AVAudioUnitTimeEffect {
	class := getAVAudioUnitTimeEffectClass()
	rv := objc.Send[AVAudioUnitTimeEffect](objc.ID(class.class), objc.Sel("new"))
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
func NewAudioUnitTimeEffectWithAudioComponentDescription(audioComponentDescription unsafe.Pointer) AVAudioUnitTimeEffect {
	instance := getAVAudioUnitTimeEffectClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitTimeEffectFromID(rv)
}

// Creates a time effect audio unit with the specified description.
//
// audioComponentDescription: The description of the audio unit to create.
//
// audioComponentDescription is a [audiotoolbox.AudioComponentDescription].
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
func (a AVAudioUnitTimeEffect) InitWithAudioComponentDescription(audioComponentDescription unsafe.Pointer) AVAudioUnitTimeEffect {
	rv := objc.Send[AVAudioUnitTimeEffect](a.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return rv
}

// The bypass state of the audio unit.
//
// # Discussion
//
// If true, the audio unit bypasses audio processing.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitTimeEffect/bypass
func (a AVAudioUnitTimeEffect) Bypass() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("bypass"))
	return rv
}
func (a AVAudioUnitTimeEffect) SetBypass(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setBypass:"), value)
}
