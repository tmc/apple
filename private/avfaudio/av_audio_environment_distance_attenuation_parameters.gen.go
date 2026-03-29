// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioEnvironmentDistanceAttenuationParameters] class.
var (
	_AVAudioEnvironmentDistanceAttenuationParametersClass     AVAudioEnvironmentDistanceAttenuationParametersClass
	_AVAudioEnvironmentDistanceAttenuationParametersClassOnce sync.Once
)

func getAVAudioEnvironmentDistanceAttenuationParametersClass() AVAudioEnvironmentDistanceAttenuationParametersClass {
	_AVAudioEnvironmentDistanceAttenuationParametersClassOnce.Do(func() {
		_AVAudioEnvironmentDistanceAttenuationParametersClass = AVAudioEnvironmentDistanceAttenuationParametersClass{class: objc.GetClass("AVAudioEnvironmentDistanceAttenuationParameters")}
	})
	return _AVAudioEnvironmentDistanceAttenuationParametersClass
}

// GetAVAudioEnvironmentDistanceAttenuationParametersClass returns the class object for AVAudioEnvironmentDistanceAttenuationParameters.
func GetAVAudioEnvironmentDistanceAttenuationParametersClass() AVAudioEnvironmentDistanceAttenuationParametersClass {
	return getAVAudioEnvironmentDistanceAttenuationParametersClass()
}

type AVAudioEnvironmentDistanceAttenuationParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioEnvironmentDistanceAttenuationParametersClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioEnvironmentDistanceAttenuationParametersClass) Alloc() AVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioEnvironmentDistanceAttenuationParameters.InitWithEnvironment]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters
type AVAudioEnvironmentDistanceAttenuationParameters struct {
	objectivec.Object
}

// AVAudioEnvironmentDistanceAttenuationParametersFromID constructs a [AVAudioEnvironmentDistanceAttenuationParameters] from an objc.ID.
func AVAudioEnvironmentDistanceAttenuationParametersFromID(id objc.ID) AVAudioEnvironmentDistanceAttenuationParameters {
	return AVAudioEnvironmentDistanceAttenuationParameters{objectivec.Object{ID: id}}
}
// Ensure AVAudioEnvironmentDistanceAttenuationParameters implements IAVAudioEnvironmentDistanceAttenuationParameters.
var _ IAVAudioEnvironmentDistanceAttenuationParameters = AVAudioEnvironmentDistanceAttenuationParameters{}

// An interface definition for the [AVAudioEnvironmentDistanceAttenuationParameters] class.
//
// # Methods
//
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.InitWithEnvironment]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters
type IAVAudioEnvironmentDistanceAttenuationParameters interface {
	objectivec.IObject

	// Topic: Methods

	InitWithEnvironment(environment unsafe.Pointer) AVAudioEnvironmentDistanceAttenuationParameters
}

// Init initializes the instance.
func (a AVAudioEnvironmentDistanceAttenuationParameters) Init() AVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioEnvironmentDistanceAttenuationParameters) Autorelease() AVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioEnvironmentDistanceAttenuationParameters creates a new AVAudioEnvironmentDistanceAttenuationParameters instance.
func NewAVAudioEnvironmentDistanceAttenuationParameters() AVAudioEnvironmentDistanceAttenuationParameters {
	class := getAVAudioEnvironmentDistanceAttenuationParametersClass()
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/initWithEnvironment:
func NewAudioEnvironmentDistanceAttenuationParametersWithEnvironment(environment unsafe.Pointer) AVAudioEnvironmentDistanceAttenuationParameters {
	instance := getAVAudioEnvironmentDistanceAttenuationParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEnvironment:"), environment)
	return AVAudioEnvironmentDistanceAttenuationParametersFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/initWithEnvironment:
func (a AVAudioEnvironmentDistanceAttenuationParameters) InitWithEnvironment(environment unsafe.Pointer) AVAudioEnvironmentDistanceAttenuationParameters {
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationParameters](a.ID, objc.Sel("initWithEnvironment:"), environment)
	return rv
}

