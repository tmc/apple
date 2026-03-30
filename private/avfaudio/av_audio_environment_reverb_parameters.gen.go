// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioEnvironmentReverbParameters] class.
var (
	_AVAudioEnvironmentReverbParametersClass     AVAudioEnvironmentReverbParametersClass
	_AVAudioEnvironmentReverbParametersClassOnce sync.Once
)

func getAVAudioEnvironmentReverbParametersClass() AVAudioEnvironmentReverbParametersClass {
	_AVAudioEnvironmentReverbParametersClassOnce.Do(func() {
		_AVAudioEnvironmentReverbParametersClass = AVAudioEnvironmentReverbParametersClass{class: objc.GetClass("AVAudioEnvironmentReverbParameters")}
	})
	return _AVAudioEnvironmentReverbParametersClass
}

// GetAVAudioEnvironmentReverbParametersClass returns the class object for AVAudioEnvironmentReverbParameters.
func GetAVAudioEnvironmentReverbParametersClass() AVAudioEnvironmentReverbParametersClass {
	return getAVAudioEnvironmentReverbParametersClass()
}

type AVAudioEnvironmentReverbParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioEnvironmentReverbParametersClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioEnvironmentReverbParametersClass) Alloc() AVAudioEnvironmentReverbParameters {
	rv := objc.Send[AVAudioEnvironmentReverbParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioEnvironmentReverbParameters.InitWithEnvironment]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters
type AVAudioEnvironmentReverbParameters struct {
	objectivec.Object
}

// AVAudioEnvironmentReverbParametersFromID constructs a [AVAudioEnvironmentReverbParameters] from an objc.ID.
func AVAudioEnvironmentReverbParametersFromID(id objc.ID) AVAudioEnvironmentReverbParameters {
	return AVAudioEnvironmentReverbParameters{objectivec.Object{ID: id}}
}

// Ensure AVAudioEnvironmentReverbParameters implements IAVAudioEnvironmentReverbParameters.
var _ IAVAudioEnvironmentReverbParameters = AVAudioEnvironmentReverbParameters{}

// An interface definition for the [AVAudioEnvironmentReverbParameters] class.
//
// # Methods
//
//   - [IAVAudioEnvironmentReverbParameters.InitWithEnvironment]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters
type IAVAudioEnvironmentReverbParameters interface {
	objectivec.IObject

	// Topic: Methods

	InitWithEnvironment(environment unsafe.Pointer) AVAudioEnvironmentReverbParameters
}

// Init initializes the instance.
func (a AVAudioEnvironmentReverbParameters) Init() AVAudioEnvironmentReverbParameters {
	rv := objc.Send[AVAudioEnvironmentReverbParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioEnvironmentReverbParameters) Autorelease() AVAudioEnvironmentReverbParameters {
	rv := objc.Send[AVAudioEnvironmentReverbParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioEnvironmentReverbParameters creates a new AVAudioEnvironmentReverbParameters instance.
func NewAVAudioEnvironmentReverbParameters() AVAudioEnvironmentReverbParameters {
	class := getAVAudioEnvironmentReverbParametersClass()
	rv := objc.Send[AVAudioEnvironmentReverbParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters/initWithEnvironment:
func NewAudioEnvironmentReverbParametersWithEnvironment(environment unsafe.Pointer) AVAudioEnvironmentReverbParameters {
	instance := getAVAudioEnvironmentReverbParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEnvironment:"), environment)
	return AVAudioEnvironmentReverbParametersFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters/initWithEnvironment:
func (a AVAudioEnvironmentReverbParameters) InitWithEnvironment(environment unsafe.Pointer) AVAudioEnvironmentReverbParameters {
	rv := objc.Send[AVAudioEnvironmentReverbParameters](a.ID, objc.Sel("initWithEnvironment:"), environment)
	return rv
}
