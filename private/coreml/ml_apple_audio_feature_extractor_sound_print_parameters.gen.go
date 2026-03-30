// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleAudioFeatureExtractorSoundPrintParameters] class.
var (
	_MLAppleAudioFeatureExtractorSoundPrintParametersClass     MLAppleAudioFeatureExtractorSoundPrintParametersClass
	_MLAppleAudioFeatureExtractorSoundPrintParametersClassOnce sync.Once
)

func getMLAppleAudioFeatureExtractorSoundPrintParametersClass() MLAppleAudioFeatureExtractorSoundPrintParametersClass {
	_MLAppleAudioFeatureExtractorSoundPrintParametersClassOnce.Do(func() {
		_MLAppleAudioFeatureExtractorSoundPrintParametersClass = MLAppleAudioFeatureExtractorSoundPrintParametersClass{class: objc.GetClass("MLAppleAudioFeatureExtractorSoundPrintParameters")}
	})
	return _MLAppleAudioFeatureExtractorSoundPrintParametersClass
}

// GetMLAppleAudioFeatureExtractorSoundPrintParametersClass returns the class object for MLAppleAudioFeatureExtractorSoundPrintParameters.
func GetMLAppleAudioFeatureExtractorSoundPrintParametersClass() MLAppleAudioFeatureExtractorSoundPrintParametersClass {
	return getMLAppleAudioFeatureExtractorSoundPrintParametersClass()
}

type MLAppleAudioFeatureExtractorSoundPrintParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleAudioFeatureExtractorSoundPrintParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleAudioFeatureExtractorSoundPrintParametersClass) Alloc() MLAppleAudioFeatureExtractorSoundPrintParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorSoundPrintParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleAudioFeatureExtractorSoundPrintParameters.SoundPrintVersion]
//   - [MLAppleAudioFeatureExtractorSoundPrintParameters.InitSoundPrintParameters]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorSoundPrintParameters
type MLAppleAudioFeatureExtractorSoundPrintParameters struct {
	objectivec.Object
}

// MLAppleAudioFeatureExtractorSoundPrintParametersFromID constructs a [MLAppleAudioFeatureExtractorSoundPrintParameters] from an objc.ID.
func MLAppleAudioFeatureExtractorSoundPrintParametersFromID(id objc.ID) MLAppleAudioFeatureExtractorSoundPrintParameters {
	return MLAppleAudioFeatureExtractorSoundPrintParameters{objectivec.Object{ID: id}}
}

// Ensure MLAppleAudioFeatureExtractorSoundPrintParameters implements IMLAppleAudioFeatureExtractorSoundPrintParameters.
var _ IMLAppleAudioFeatureExtractorSoundPrintParameters = MLAppleAudioFeatureExtractorSoundPrintParameters{}

// An interface definition for the [MLAppleAudioFeatureExtractorSoundPrintParameters] class.
//
// # Methods
//
//   - [IMLAppleAudioFeatureExtractorSoundPrintParameters.SoundPrintVersion]
//   - [IMLAppleAudioFeatureExtractorSoundPrintParameters.InitSoundPrintParameters]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorSoundPrintParameters
type IMLAppleAudioFeatureExtractorSoundPrintParameters interface {
	objectivec.IObject

	// Topic: Methods

	SoundPrintVersion() int64
	InitSoundPrintParameters(parameters int64) MLAppleAudioFeatureExtractorSoundPrintParameters
}

// Init initializes the instance.
func (a MLAppleAudioFeatureExtractorSoundPrintParameters) Init() MLAppleAudioFeatureExtractorSoundPrintParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorSoundPrintParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleAudioFeatureExtractorSoundPrintParameters) Autorelease() MLAppleAudioFeatureExtractorSoundPrintParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorSoundPrintParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleAudioFeatureExtractorSoundPrintParameters creates a new MLAppleAudioFeatureExtractorSoundPrintParameters instance.
func NewMLAppleAudioFeatureExtractorSoundPrintParameters() MLAppleAudioFeatureExtractorSoundPrintParameters {
	class := getMLAppleAudioFeatureExtractorSoundPrintParametersClass()
	rv := objc.Send[MLAppleAudioFeatureExtractorSoundPrintParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorSoundPrintParameters/initSoundPrintParameters:
func NewAppleAudioFeatureExtractorSoundPrintParametersSoundPrintParameters(parameters int64) MLAppleAudioFeatureExtractorSoundPrintParameters {
	instance := getMLAppleAudioFeatureExtractorSoundPrintParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initSoundPrintParameters:"), parameters)
	return MLAppleAudioFeatureExtractorSoundPrintParametersFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorSoundPrintParameters/initSoundPrintParameters:
func (a MLAppleAudioFeatureExtractorSoundPrintParameters) InitSoundPrintParameters(parameters int64) MLAppleAudioFeatureExtractorSoundPrintParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorSoundPrintParameters](a.ID, objc.Sel("initSoundPrintParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorSoundPrintParameters/soundPrintVersion
func (a MLAppleAudioFeatureExtractorSoundPrintParameters) SoundPrintVersion() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("soundPrintVersion"))
	return rv
}
