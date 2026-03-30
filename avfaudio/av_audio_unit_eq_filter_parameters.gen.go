// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioUnitEQFilterParameters] class.
var (
	_AVAudioUnitEQFilterParametersClass     AVAudioUnitEQFilterParametersClass
	_AVAudioUnitEQFilterParametersClassOnce sync.Once
)

func getAVAudioUnitEQFilterParametersClass() AVAudioUnitEQFilterParametersClass {
	_AVAudioUnitEQFilterParametersClassOnce.Do(func() {
		_AVAudioUnitEQFilterParametersClass = AVAudioUnitEQFilterParametersClass{class: objc.GetClass("AVAudioUnitEQFilterParameters")}
	})
	return _AVAudioUnitEQFilterParametersClass
}

// GetAVAudioUnitEQFilterParametersClass returns the class object for AVAudioUnitEQFilterParameters.
func GetAVAudioUnitEQFilterParametersClass() AVAudioUnitEQFilterParametersClass {
	return getAVAudioUnitEQFilterParametersClass()
}

type AVAudioUnitEQFilterParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitEQFilterParametersClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitEQFilterParametersClass) Alloc() AVAudioUnitEQFilterParameters {
	rv := objc.Send[AVAudioUnitEQFilterParameters](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates the parameters that the equalizer uses.
//
// # Overview
//
// # Getting and Setting Equalizer Filter Parameters
//
//   - [AVAudioUnitEQFilterParameters.Bandwidth]: The bandwidth of the equalizer filter, in octaves.
//   - [AVAudioUnitEQFilterParameters.SetBandwidth]
//   - [AVAudioUnitEQFilterParameters.Bypass]: The bypass state of the equalizer filter band.
//   - [AVAudioUnitEQFilterParameters.SetBypass]
//   - [AVAudioUnitEQFilterParameters.FilterType]: The equalizer filter type.
//   - [AVAudioUnitEQFilterParameters.SetFilterType]
//   - [AVAudioUnitEQFilterParameters.Frequency]: The frequency of the equalizer filter, in hertz.
//   - [AVAudioUnitEQFilterParameters.SetFrequency]
//   - [AVAudioUnitEQFilterParameters.Gain]: The gain of the equalizer filter, in decibels.
//   - [AVAudioUnitEQFilterParameters.SetGain]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters
type AVAudioUnitEQFilterParameters struct {
	objectivec.Object
}

// AVAudioUnitEQFilterParametersFromID constructs a [AVAudioUnitEQFilterParameters] from an objc.ID.
//
// An object that encapsulates the parameters that the equalizer uses.
func AVAudioUnitEQFilterParametersFromID(id objc.ID) AVAudioUnitEQFilterParameters {
	return AVAudioUnitEQFilterParameters{objectivec.Object{ID: id}}
}

// NOTE: AVAudioUnitEQFilterParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioUnitEQFilterParameters] class.
//
// # Getting and Setting Equalizer Filter Parameters
//
//   - [IAVAudioUnitEQFilterParameters.Bandwidth]: The bandwidth of the equalizer filter, in octaves.
//   - [IAVAudioUnitEQFilterParameters.SetBandwidth]
//   - [IAVAudioUnitEQFilterParameters.Bypass]: The bypass state of the equalizer filter band.
//   - [IAVAudioUnitEQFilterParameters.SetBypass]
//   - [IAVAudioUnitEQFilterParameters.FilterType]: The equalizer filter type.
//   - [IAVAudioUnitEQFilterParameters.SetFilterType]
//   - [IAVAudioUnitEQFilterParameters.Frequency]: The frequency of the equalizer filter, in hertz.
//   - [IAVAudioUnitEQFilterParameters.SetFrequency]
//   - [IAVAudioUnitEQFilterParameters.Gain]: The gain of the equalizer filter, in decibels.
//   - [IAVAudioUnitEQFilterParameters.SetGain]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters
type IAVAudioUnitEQFilterParameters interface {
	objectivec.IObject

	// Topic: Getting and Setting Equalizer Filter Parameters

	// The bandwidth of the equalizer filter, in octaves.
	Bandwidth() float32
	SetBandwidth(value float32)
	// The bypass state of the equalizer filter band.
	Bypass() bool
	SetBypass(value bool)
	// The equalizer filter type.
	FilterType() AVAudioUnitEQFilterType
	SetFilterType(value AVAudioUnitEQFilterType)
	// The frequency of the equalizer filter, in hertz.
	Frequency() float32
	SetFrequency(value float32)
	// The gain of the equalizer filter, in decibels.
	Gain() float32
	SetGain(value float32)

	// An array of equalizer filter parameters.
	Bands() IAVAudioUnitEQFilterParameters
	SetBands(value IAVAudioUnitEQFilterParameters)
	// The overall gain adjustment that the audio unit applies to the signal, in decibels.
	GlobalGain() float32
	SetGlobalGain(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitEQFilterParameters) Init() AVAudioUnitEQFilterParameters {
	rv := objc.Send[AVAudioUnitEQFilterParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitEQFilterParameters) Autorelease() AVAudioUnitEQFilterParameters {
	rv := objc.Send[AVAudioUnitEQFilterParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitEQFilterParameters creates a new AVAudioUnitEQFilterParameters instance.
func NewAVAudioUnitEQFilterParameters() AVAudioUnitEQFilterParameters {
	class := getAVAudioUnitEQFilterParametersClass()
	rv := objc.Send[AVAudioUnitEQFilterParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The bandwidth of the equalizer filter, in octaves.
//
// # Discussion
//
// The value range of values is `0.05` to `5.0` octaves.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/bandwidth
func (a AVAudioUnitEQFilterParameters) Bandwidth() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("bandwidth"))
	return rv
}
func (a AVAudioUnitEQFilterParameters) SetBandwidth(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setBandwidth:"), value)
}

// The bypass state of the equalizer filter band.
//
// # Discussion
//
// true if the bypass is active; otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/bypass
func (a AVAudioUnitEQFilterParameters) Bypass() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("bypass"))
	return rv
}
func (a AVAudioUnitEQFilterParameters) SetBypass(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setBypass:"), value)
}

// The equalizer filter type.
//
// # Discussion
//
// The default value is [AVAudioUnitEQFilterTypeParametric].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/filterType
func (a AVAudioUnitEQFilterParameters) FilterType() AVAudioUnitEQFilterType {
	rv := objc.Send[AVAudioUnitEQFilterType](a.ID, objc.Sel("filterType"))
	return AVAudioUnitEQFilterType(rv)
}
func (a AVAudioUnitEQFilterParameters) SetFilterType(value AVAudioUnitEQFilterType) {
	objc.Send[struct{}](a.ID, objc.Sel("setFilterType:"), value)
}

// The frequency of the equalizer filter, in hertz.
//
// # Discussion
//
// The valid range of values is `20 Hz` through `(SampleRate/2)`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/frequency
func (a AVAudioUnitEQFilterParameters) Frequency() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("frequency"))
	return rv
}
func (a AVAudioUnitEQFilterParameters) SetFrequency(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setFrequency:"), value)
}

// The gain of the equalizer filter, in decibels.
//
// # Discussion
//
// The default value is `0 dB`. The valid range of values is `-96 dB` through
// `24 dB`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitEQFilterParameters/gain
func (a AVAudioUnitEQFilterParameters) Gain() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("gain"))
	return rv
}
func (a AVAudioUnitEQFilterParameters) SetGain(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setGain:"), value)
}

// An array of equalizer filter parameters.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiouniteq/bands
func (a AVAudioUnitEQFilterParameters) Bands() IAVAudioUnitEQFilterParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("bands"))
	return AVAudioUnitEQFilterParametersFromID(objc.ID(rv))
}
func (a AVAudioUnitEQFilterParameters) SetBands(value IAVAudioUnitEQFilterParameters) {
	objc.Send[struct{}](a.ID, objc.Sel("setBands:"), value)
}

// The overall gain adjustment that the audio unit applies to the signal, in
// decibels.
//
// See: https://developer.apple.com/documentation/avfaudio/avaudiouniteq/globalgain
func (a AVAudioUnitEQFilterParameters) GlobalGain() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("globalGain"))
	return rv
}
func (a AVAudioUnitEQFilterParameters) SetGlobalGain(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setGlobalGain:"), value)
}
