// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
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

// A class that encapsulates the parameters that you use to control the reverb
// of the environment node class.
//
// # Overview
// 
// Use reverberation to simulate the acoustic characteristics of an
// environment. The [AVAudioEnvironmentNode] class has a built-in reverb that
// describe the space that the listener is in.
// 
// The reverb has a single filter that sits at the end of the chain. You use
// this filter to shape the overall sound of the reverb. For instance, select
// one of the reverb presets to simulate the general space, and then use the
// filter to brighten or darken the overall sound.
// 
// You can’t create a standalone instance of
// [AVAudioEnvironmentReverbParameters]. Only an instance vended by a source
// object is valid, such as an [AVAudioEnvironmentNode] instance.
//
// # Enabling and Disabling Reverb
//
//   - [AVAudioEnvironmentReverbParameters.Enable]: A Boolean value that indicates whether reverberation is in an enabled state.
//   - [AVAudioEnvironmentReverbParameters.SetEnable]
//
// # Getting and Setting Reverb Values
//
//   - [AVAudioEnvironmentReverbParameters.Level]: Controls the amount of reverb, in decibels.
//   - [AVAudioEnvironmentReverbParameters.SetLevel]
//   - [AVAudioEnvironmentReverbParameters.FilterParameters]: A filter that the system applies to the output.
//   - [AVAudioEnvironmentReverbParameters.LoadFactoryReverbPreset]: Loads one of the reverbs factory presets.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters
type AVAudioEnvironmentReverbParameters struct {
	objectivec.Object
}

// AVAudioEnvironmentReverbParametersFromID constructs a [AVAudioEnvironmentReverbParameters] from an objc.ID.
//
// A class that encapsulates the parameters that you use to control the reverb
// of the environment node class.
func AVAudioEnvironmentReverbParametersFromID(id objc.ID) AVAudioEnvironmentReverbParameters {
	return AVAudioEnvironmentReverbParameters{objectivec.Object{ID: id}}
}
// NOTE: AVAudioEnvironmentReverbParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioEnvironmentReverbParameters] class.
//
// # Enabling and Disabling Reverb
//
//   - [IAVAudioEnvironmentReverbParameters.Enable]: A Boolean value that indicates whether reverberation is in an enabled state.
//   - [IAVAudioEnvironmentReverbParameters.SetEnable]
//
// # Getting and Setting Reverb Values
//
//   - [IAVAudioEnvironmentReverbParameters.Level]: Controls the amount of reverb, in decibels.
//   - [IAVAudioEnvironmentReverbParameters.SetLevel]
//   - [IAVAudioEnvironmentReverbParameters.FilterParameters]: A filter that the system applies to the output.
//   - [IAVAudioEnvironmentReverbParameters.LoadFactoryReverbPreset]: Loads one of the reverbs factory presets.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters
type IAVAudioEnvironmentReverbParameters interface {
	objectivec.IObject

	// Topic: Enabling and Disabling Reverb

	// A Boolean value that indicates whether reverberation is in an enabled state.
	Enable() bool
	SetEnable(value bool)

	// Topic: Getting and Setting Reverb Values

	// Controls the amount of reverb, in decibels.
	Level() float32
	SetLevel(value float32)
	// A filter that the system applies to the output.
	FilterParameters() IAVAudioUnitEQFilterParameters
	// Loads one of the reverbs factory presets.
	LoadFactoryReverbPreset(preset AVAudioUnitReverbPreset)
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

// Loads one of the reverbs factory presets.
//
// preset: A reverb preset to load.
//
// # Discussion
// 
// Loading a factory reverb preset changes the sound of the reverb. This is
// independent of the filter which follows the reverb in the signal chain.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters/loadFactoryReverbPreset(_:)
func (a AVAudioEnvironmentReverbParameters) LoadFactoryReverbPreset(preset AVAudioUnitReverbPreset) {
	objc.Send[objc.ID](a.ID, objc.Sel("loadFactoryReverbPreset:"), preset)
}

// A Boolean value that indicates whether reverberation is in an enabled
// state.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters/enable
func (a AVAudioEnvironmentReverbParameters) Enable() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("enable"))
	return rv
}
func (a AVAudioEnvironmentReverbParameters) SetEnable(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setEnable:"), value)
}
// Controls the amount of reverb, in decibels.
//
// # Discussion
// 
// The default value is `0.0`. The values must be within the range of `-40` to
// `40` dB.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters/level
func (a AVAudioEnvironmentReverbParameters) Level() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("level"))
	return rv
}
func (a AVAudioEnvironmentReverbParameters) SetLevel(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setLevel:"), value)
}
// A filter that the system applies to the output.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentReverbParameters/filterParameters
func (a AVAudioEnvironmentReverbParameters) FilterParameters() IAVAudioUnitEQFilterParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("filterParameters"))
	return AVAudioUnitEQFilterParametersFromID(objc.ID(rv))
}

