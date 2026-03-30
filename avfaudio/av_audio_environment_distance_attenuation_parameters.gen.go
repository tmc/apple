// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
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

// An object that specifies the amount of attenuation distance, the gradual
// loss in audio intensity, and other characteristics.
//
// # Overview
//
// # Getting and Setting the Attenuation Model
//
//   - [AVAudioEnvironmentDistanceAttenuationParameters.DistanceAttenuationModel]: The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//   - [AVAudioEnvironmentDistanceAttenuationParameters.SetDistanceAttenuationModel]
//
// # Getting and Setting the Attenuation Values
//
//   - [AVAudioEnvironmentDistanceAttenuationParameters.MaximumDistance]: The distance beyond which the node applies no further attenuation, in meters.
//   - [AVAudioEnvironmentDistanceAttenuationParameters.SetMaximumDistance]
//   - [AVAudioEnvironmentDistanceAttenuationParameters.ReferenceDistance]: The minimum distance at which the node applies attenuation, in meters.
//   - [AVAudioEnvironmentDistanceAttenuationParameters.SetReferenceDistance]
//   - [AVAudioEnvironmentDistanceAttenuationParameters.RolloffFactor]: A factor that determines the attenuation curve.
//   - [AVAudioEnvironmentDistanceAttenuationParameters.SetRolloffFactor]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters
type AVAudioEnvironmentDistanceAttenuationParameters struct {
	objectivec.Object
}

// AVAudioEnvironmentDistanceAttenuationParametersFromID constructs a [AVAudioEnvironmentDistanceAttenuationParameters] from an objc.ID.
//
// An object that specifies the amount of attenuation distance, the gradual
// loss in audio intensity, and other characteristics.
func AVAudioEnvironmentDistanceAttenuationParametersFromID(id objc.ID) AVAudioEnvironmentDistanceAttenuationParameters {
	return AVAudioEnvironmentDistanceAttenuationParameters{objectivec.Object{ID: id}}
}

// NOTE: AVAudioEnvironmentDistanceAttenuationParameters adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioEnvironmentDistanceAttenuationParameters] class.
//
// # Getting and Setting the Attenuation Model
//
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.DistanceAttenuationModel]: The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.SetDistanceAttenuationModel]
//
// # Getting and Setting the Attenuation Values
//
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.MaximumDistance]: The distance beyond which the node applies no further attenuation, in meters.
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.SetMaximumDistance]
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.ReferenceDistance]: The minimum distance at which the node applies attenuation, in meters.
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.SetReferenceDistance]
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.RolloffFactor]: A factor that determines the attenuation curve.
//   - [IAVAudioEnvironmentDistanceAttenuationParameters.SetRolloffFactor]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters
type IAVAudioEnvironmentDistanceAttenuationParameters interface {
	objectivec.IObject

	// Topic: Getting and Setting the Attenuation Model

	// The distance attenuation model that describes the drop-off in gain as the source moves away from the listener.
	DistanceAttenuationModel() AVAudioEnvironmentDistanceAttenuationModel
	SetDistanceAttenuationModel(value AVAudioEnvironmentDistanceAttenuationModel)

	// Topic: Getting and Setting the Attenuation Values

	// The distance beyond which the node applies no further attenuation, in meters.
	MaximumDistance() float32
	SetMaximumDistance(value float32)
	// The minimum distance at which the node applies attenuation, in meters.
	ReferenceDistance() float32
	SetReferenceDistance(value float32)
	// A factor that determines the attenuation curve.
	RolloffFactor() float32
	SetRolloffFactor(value float32)
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

// The distance attenuation model that describes the drop-off in gain as the
// source moves away from the listener.
//
// # Discussion
//
// The default value is the
// [AVAudioEnvironmentDistanceAttenuationModelInverse] attenuation model.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/distanceAttenuationModel
func (a AVAudioEnvironmentDistanceAttenuationParameters) DistanceAttenuationModel() AVAudioEnvironmentDistanceAttenuationModel {
	rv := objc.Send[AVAudioEnvironmentDistanceAttenuationModel](a.ID, objc.Sel("distanceAttenuationModel"))
	return AVAudioEnvironmentDistanceAttenuationModel(rv)
}
func (a AVAudioEnvironmentDistanceAttenuationParameters) SetDistanceAttenuationModel(value AVAudioEnvironmentDistanceAttenuationModel) {
	objc.Send[struct{}](a.ID, objc.Sel("setDistanceAttenuationModel:"), value)
}

// The distance beyond which the node applies no further attenuation, in
// meters.
//
// # Discussion
//
// The default value is `100000.0` meters.
//
// This property is relevant for
// [AVAudioEnvironmentDistanceAttenuationModelInverse].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/maximumDistance
func (a AVAudioEnvironmentDistanceAttenuationParameters) MaximumDistance() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("maximumDistance"))
	return rv
}
func (a AVAudioEnvironmentDistanceAttenuationParameters) SetMaximumDistance(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMaximumDistance:"), value)
}

// The minimum distance at which the node applies attenuation, in meters.
//
// # Discussion
//
// The default value is `1.0` meter.
//
// This property is relevant for
// [AVAudioEnvironmentDistanceAttenuationModelInverse] and
// [AVAudioEnvironmentDistanceAttenuationModelLinear].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/referenceDistance
func (a AVAudioEnvironmentDistanceAttenuationParameters) ReferenceDistance() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("referenceDistance"))
	return rv
}
func (a AVAudioEnvironmentDistanceAttenuationParameters) SetReferenceDistance(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setReferenceDistance:"), value)
}

// A factor that determines the attenuation curve.
//
// # Discussion
//
// A higher value results in a steeper attenuation curve. The default value is
// `1.0`, and the value must be greater than `0.0`.
//
// This property is relevant for
// [AVAudioEnvironmentDistanceAttenuationModelExponential],
// [AVAudioEnvironmentDistanceAttenuationModelInverse], and
// [AVAudioEnvironmentDistanceAttenuationModelLinear].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioEnvironmentDistanceAttenuationParameters/rolloffFactor
func (a AVAudioEnvironmentDistanceAttenuationParameters) RolloffFactor() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("rolloffFactor"))
	return rv
}
func (a AVAudioEnvironmentDistanceAttenuationParameters) SetRolloffFactor(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRolloffFactor:"), value)
}
