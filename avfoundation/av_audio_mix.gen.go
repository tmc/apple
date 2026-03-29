// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioMix] class.
var (
	_AVAudioMixClass     AVAudioMixClass
	_AVAudioMixClassOnce sync.Once
)

func getAVAudioMixClass() AVAudioMixClass {
	_AVAudioMixClassOnce.Do(func() {
		_AVAudioMixClass = AVAudioMixClass{class: objc.GetClass("AVAudioMix")}
	})
	return _AVAudioMixClass
}

// GetAVAudioMixClass returns the class object for AVAudioMix.
func GetAVAudioMixClass() AVAudioMixClass {
	return getAVAudioMixClass()
}

type AVAudioMixClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioMixClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioMixClass) Alloc() AVAudioMix {
	rv := objc.Send[AVAudioMix](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the input parameters for mixing audio tracks.
//
// # Retrieving input parameters
//
//   - [AVAudioMix.InputParameters]: An array of input parameters for the mix.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMix
type AVAudioMix struct {
	objectivec.Object
}

// AVAudioMixFromID constructs a [AVAudioMix] from an objc.ID.
//
// An object that manages the input parameters for mixing audio tracks.
func AVAudioMixFromID(id objc.ID) AVAudioMix {
	return AVAudioMix{objectivec.Object{ID: id}}
}
// NOTE: AVAudioMix adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioMix] class.
//
// # Retrieving input parameters
//
//   - [IAVAudioMix.InputParameters]: An array of input parameters for the mix.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMix
type IAVAudioMix interface {
	objectivec.IObject

	// Topic: Retrieving input parameters

	// An array of input parameters for the mix.
	InputParameters() []AVAudioMixInputParameters
}

// Init initializes the instance.
func (a AVAudioMix) Init() AVAudioMix {
	rv := objc.Send[AVAudioMix](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioMix) Autorelease() AVAudioMix {
	rv := objc.Send[AVAudioMix](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMix creates a new AVAudioMix instance.
func NewAVAudioMix() AVAudioMix {
	class := getAVAudioMixClass()
	rv := objc.Send[AVAudioMix](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An array of input parameters for the mix.
//
// # Discussion
// 
// The array contains instances of [AVAudioMixInputParameters].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioMix/inputParameters
func (a AVAudioMix) InputParameters() []AVAudioMixInputParameters {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("inputParameters"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAudioMixInputParameters {
		return AVAudioMixInputParametersFromID(id)
	})
}

