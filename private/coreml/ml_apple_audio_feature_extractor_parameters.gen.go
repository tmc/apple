// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleAudioFeatureExtractorParameters] class.
var (
	_MLAppleAudioFeatureExtractorParametersClass     MLAppleAudioFeatureExtractorParametersClass
	_MLAppleAudioFeatureExtractorParametersClassOnce sync.Once
)

func getMLAppleAudioFeatureExtractorParametersClass() MLAppleAudioFeatureExtractorParametersClass {
	_MLAppleAudioFeatureExtractorParametersClassOnce.Do(func() {
		_MLAppleAudioFeatureExtractorParametersClass = MLAppleAudioFeatureExtractorParametersClass{class: objc.GetClass("MLAppleAudioFeatureExtractorParameters")}
	})
	return _MLAppleAudioFeatureExtractorParametersClass
}

// GetMLAppleAudioFeatureExtractorParametersClass returns the class object for MLAppleAudioFeatureExtractorParameters.
func GetMLAppleAudioFeatureExtractorParametersClass() MLAppleAudioFeatureExtractorParametersClass {
	return getMLAppleAudioFeatureExtractorParametersClass()
}

type MLAppleAudioFeatureExtractorParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleAudioFeatureExtractorParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleAudioFeatureExtractorParametersClass) Alloc() MLAppleAudioFeatureExtractorParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleAudioFeatureExtractorParameters.FeatureExtractorParameters]
//   - [MLAppleAudioFeatureExtractorParameters.InitWithSoundPrintParameters]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorParameters
type MLAppleAudioFeatureExtractorParameters struct {
	objectivec.Object
}

// MLAppleAudioFeatureExtractorParametersFromID constructs a [MLAppleAudioFeatureExtractorParameters] from an objc.ID.
func MLAppleAudioFeatureExtractorParametersFromID(id objc.ID) MLAppleAudioFeatureExtractorParameters {
	return MLAppleAudioFeatureExtractorParameters{objectivec.Object{ID: id}}
}

// Ensure MLAppleAudioFeatureExtractorParameters implements IMLAppleAudioFeatureExtractorParameters.
var _ IMLAppleAudioFeatureExtractorParameters = MLAppleAudioFeatureExtractorParameters{}

// An interface definition for the [MLAppleAudioFeatureExtractorParameters] class.
//
// # Methods
//
//   - [IMLAppleAudioFeatureExtractorParameters.FeatureExtractorParameters]
//   - [IMLAppleAudioFeatureExtractorParameters.InitWithSoundPrintParameters]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorParameters
type IMLAppleAudioFeatureExtractorParameters interface {
	objectivec.IObject

	// Topic: Methods

	FeatureExtractorParameters() objectivec.IObject
	InitWithSoundPrintParameters(parameters objectivec.IObject) MLAppleAudioFeatureExtractorParameters
}

// Init initializes the instance.
func (a MLAppleAudioFeatureExtractorParameters) Init() MLAppleAudioFeatureExtractorParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorParameters](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleAudioFeatureExtractorParameters) Autorelease() MLAppleAudioFeatureExtractorParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorParameters](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleAudioFeatureExtractorParameters creates a new MLAppleAudioFeatureExtractorParameters instance.
func NewMLAppleAudioFeatureExtractorParameters() MLAppleAudioFeatureExtractorParameters {
	class := getMLAppleAudioFeatureExtractorParametersClass()
	rv := objc.Send[MLAppleAudioFeatureExtractorParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorParameters/initWithSoundPrintParameters:
func NewAppleAudioFeatureExtractorParametersWithSoundPrintParameters(parameters objectivec.IObject) MLAppleAudioFeatureExtractorParameters {
	instance := getMLAppleAudioFeatureExtractorParametersClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSoundPrintParameters:"), parameters)
	return MLAppleAudioFeatureExtractorParametersFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorParameters/initWithSoundPrintParameters:
func (a MLAppleAudioFeatureExtractorParameters) InitWithSoundPrintParameters(parameters objectivec.IObject) MLAppleAudioFeatureExtractorParameters {
	rv := objc.Send[MLAppleAudioFeatureExtractorParameters](a.ID, objc.Sel("initWithSoundPrintParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractorParameters/featureExtractorParameters
func (a MLAppleAudioFeatureExtractorParameters) FeatureExtractorParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("featureExtractorParameters"))
	return objectivec.Object{ID: rv}
}
