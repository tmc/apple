// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitEQ] class.
var (
	_AVAudioUnitEQClass     AVAudioUnitEQClass
	_AVAudioUnitEQClassOnce sync.Once
)

func getAVAudioUnitEQClass() AVAudioUnitEQClass {
	_AVAudioUnitEQClassOnce.Do(func() {
		_AVAudioUnitEQClass = AVAudioUnitEQClass{class: objc.GetClass("AVAudioUnitEQ")}
	})
	return _AVAudioUnitEQClass
}

// GetAVAudioUnitEQClass returns the class object for AVAudioUnitEQ.
func GetAVAudioUnitEQClass() AVAudioUnitEQClass {
	return getAVAudioUnitEQClass()
}

type AVAudioUnitEQClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitEQClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitEQClass) Alloc() AVAudioUnitEQ {
	rv := objc.Send[AVAudioUnitEQ](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that implements a multiband equalizer.
//
// # Overview
// 
// The [AVAudioUnitEQFilterParameters] class encapsulates the filter
// parameters that the [AVAudioUnitEQ.Bands] property array returns.
//
// # Creating an equalizer
//
//   - [AVAudioUnitEQ.InitWithNumberOfBands]: Creates an audio unit equalizer object with the specified number of bands.
//
// # Getting and setting the equalizer values
//
//   - [AVAudioUnitEQ.Bands]: An array of equalizer filter parameters.
//   - [AVAudioUnitEQ.GlobalGain]: The overall gain adjustment that the audio unit applies to the signal, in decibels.
//   - [AVAudioUnitEQ.SetGlobalGain]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ
type AVAudioUnitEQ struct {
	AVAudioUnitEffect
}

// AVAudioUnitEQFromID constructs a [AVAudioUnitEQ] from an objc.ID.
//
// An object that implements a multiband equalizer.
func AVAudioUnitEQFromID(id objc.ID) AVAudioUnitEQ {
	return AVAudioUnitEQ{AVAudioUnitEffect: AVAudioUnitEffectFromID(id)}
}
// NOTE: AVAudioUnitEQ adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitEQ] class.
//
// # Creating an equalizer
//
//   - [IAVAudioUnitEQ.InitWithNumberOfBands]: Creates an audio unit equalizer object with the specified number of bands.
//
// # Getting and setting the equalizer values
//
//   - [IAVAudioUnitEQ.Bands]: An array of equalizer filter parameters.
//   - [IAVAudioUnitEQ.GlobalGain]: The overall gain adjustment that the audio unit applies to the signal, in decibels.
//   - [IAVAudioUnitEQ.SetGlobalGain]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ
type IAVAudioUnitEQ interface {
	IAVAudioUnitEffect

	// Topic: Creating an equalizer

	// Creates an audio unit equalizer object with the specified number of bands.
	InitWithNumberOfBands(numberOfBands uint) AVAudioUnitEQ

	// Topic: Getting and setting the equalizer values

	// An array of equalizer filter parameters.
	Bands() []AVAudioUnitEQFilterParameters
	// The overall gain adjustment that the audio unit applies to the signal, in decibels.
	GlobalGain() float32
	SetGlobalGain(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitEQ) Init() AVAudioUnitEQ {
	rv := objc.Send[AVAudioUnitEQ](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitEQ) Autorelease() AVAudioUnitEQ {
	rv := objc.Send[AVAudioUnitEQ](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitEQ creates a new AVAudioUnitEQ instance.
func NewAVAudioUnitEQ() AVAudioUnitEQ {
	class := getAVAudioUnitEQClass()
	rv := objc.Send[AVAudioUnitEQ](objc.ID(class.class), objc.Sel("new"))
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
func NewAudioUnitEQWithAudioComponentDescription(audioComponentDescription objectivec.IObject) AVAudioUnitEQ {
	instance := getAVAudioUnitEQClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioComponentDescription:"), audioComponentDescription)
	return AVAudioUnitEQFromID(rv)
}

// Creates an audio unit equalizer object with the specified number of bands.
//
// numberOfBands: The number of bands that the equalizer creates.
//
// # Return Value
// 
// A new [AVAudioUnitEQ] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/init(numberOfBands:)
func NewAudioUnitEQWithNumberOfBands(numberOfBands uint) AVAudioUnitEQ {
	instance := getAVAudioUnitEQClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNumberOfBands:"), numberOfBands)
	return AVAudioUnitEQFromID(rv)
}

// Creates an audio unit equalizer object with the specified number of bands.
//
// numberOfBands: The number of bands that the equalizer creates.
//
// # Return Value
// 
// A new [AVAudioUnitEQ] instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/init(numberOfBands:)
func (a AVAudioUnitEQ) InitWithNumberOfBands(numberOfBands uint) AVAudioUnitEQ {
	rv := objc.Send[AVAudioUnitEQ](a.ID, objc.Sel("initWithNumberOfBands:"), numberOfBands)
	return rv
}

// An array of equalizer filter parameters.
//
// # Discussion
// 
// The number of elements in the array is equal to the number of bands.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/bands
func (a AVAudioUnitEQ) Bands() []AVAudioUnitEQFilterParameters {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("bands"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioUnitEQFilterParameters {
		return AVAudioUnitEQFilterParametersFromID(id)
	})
}
// The overall gain adjustment that the audio unit applies to the signal, in
// decibels.
//
// # Discussion
// 
// The default value is `0 db`. The valid range of values is `-96 db` to `24
// db`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQ/globalGain
func (a AVAudioUnitEQ) GlobalGain() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("globalGain"))
	return rv
}
func (a AVAudioUnitEQ) SetGlobalGain(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setGlobalGain:"), value)
}

