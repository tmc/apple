// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleAudioFeatureExtractor] class.
var (
	_MLAppleAudioFeatureExtractorClass     MLAppleAudioFeatureExtractorClass
	_MLAppleAudioFeatureExtractorClassOnce sync.Once
)

func getMLAppleAudioFeatureExtractorClass() MLAppleAudioFeatureExtractorClass {
	_MLAppleAudioFeatureExtractorClassOnce.Do(func() {
		_MLAppleAudioFeatureExtractorClass = MLAppleAudioFeatureExtractorClass{class: objc.GetClass("MLAppleAudioFeatureExtractor")}
	})
	return _MLAppleAudioFeatureExtractorClass
}

// GetMLAppleAudioFeatureExtractorClass returns the class object for MLAppleAudioFeatureExtractor.
func GetMLAppleAudioFeatureExtractorClass() MLAppleAudioFeatureExtractorClass {
	return getMLAppleAudioFeatureExtractorClass()
}

type MLAppleAudioFeatureExtractorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleAudioFeatureExtractorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleAudioFeatureExtractorClass) Alloc() MLAppleAudioFeatureExtractor {
	rv := objc.Send[MLAppleAudioFeatureExtractor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleAudioFeatureExtractor.Parameters]
//   - [MLAppleAudioFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [MLAppleAudioFeatureExtractor.InitWithParametersModelDescriptionConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor
type MLAppleAudioFeatureExtractor struct {
	MLModel
}

// MLAppleAudioFeatureExtractorFromID constructs a [MLAppleAudioFeatureExtractor] from an objc.ID.
func MLAppleAudioFeatureExtractorFromID(id objc.ID) MLAppleAudioFeatureExtractor {
	return MLAppleAudioFeatureExtractor{MLModel: MLModelFromID(id)}
}
// Ensure MLAppleAudioFeatureExtractor implements IMLAppleAudioFeatureExtractor.
var _ IMLAppleAudioFeatureExtractor = MLAppleAudioFeatureExtractor{}

// An interface definition for the [MLAppleAudioFeatureExtractor] class.
//
// # Methods
//
//   - [IMLAppleAudioFeatureExtractor.Parameters]
//   - [IMLAppleAudioFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [IMLAppleAudioFeatureExtractor.InitWithParametersModelDescriptionConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor
type IMLAppleAudioFeatureExtractor interface {
	IMLModel

	// Topic: Methods

	Parameters() IMLAppleAudioFeatureExtractorParameters
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithParametersModelDescriptionConfigurationError(parameters objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLAppleAudioFeatureExtractor, error)
}

// Init initializes the instance.
func (a MLAppleAudioFeatureExtractor) Init() MLAppleAudioFeatureExtractor {
	rv := objc.Send[MLAppleAudioFeatureExtractor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleAudioFeatureExtractor) Autorelease() MLAppleAudioFeatureExtractor {
	rv := objc.Send[MLAppleAudioFeatureExtractor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleAudioFeatureExtractor creates a new MLAppleAudioFeatureExtractor instance.
func NewMLAppleAudioFeatureExtractor() MLAppleAudioFeatureExtractor {
	class := getMLAppleAudioFeatureExtractorClass()
	rv := objc.Send[MLAppleAudioFeatureExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewAppleAudioFeatureExtractorDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleAudioFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleAudioFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleAudioFeatureExtractorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewAppleAudioFeatureExtractorInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleAudioFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleAudioFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleAudioFeatureExtractorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewAppleAudioFeatureExtractorWithConfiguration(configuration objectivec.IObject) MLAppleAudioFeatureExtractor {
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleAudioFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewAppleAudioFeatureExtractorWithDescription(description objectivec.IObject) MLAppleAudioFeatureExtractor {
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleAudioFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewAppleAudioFeatureExtractorWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleAudioFeatureExtractor {
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleAudioFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewAppleAudioFeatureExtractorWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleAudioFeatureExtractor {
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleAudioFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor/initWithParameters:modelDescription:configuration:error:
func NewAppleAudioFeatureExtractorWithParametersModelDescriptionConfigurationError(parameters objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLAppleAudioFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleAudioFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:configuration:error:"), parameters, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleAudioFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleAudioFeatureExtractorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor/predictionFromFeatures:options:error:
func (a MLAppleAudioFeatureExtractor) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor/initWithParameters:modelDescription:configuration:error:
func (a MLAppleAudioFeatureExtractor) InitWithParametersModelDescriptionConfigurationError(parameters objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject) (MLAppleAudioFeatureExtractor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithParameters:modelDescription:configuration:error:"), parameters, description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleAudioFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleAudioFeatureExtractorFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor/loadModelFromSpecification:configuration:error:
func (_MLAppleAudioFeatureExtractorClass MLAppleAudioFeatureExtractorClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleAudioFeatureExtractorClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleAudioFeatureExtractor/parameters
func (a MLAppleAudioFeatureExtractor) Parameters() IMLAppleAudioFeatureExtractorParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameters"))
	return MLAppleAudioFeatureExtractorParametersFromID(objc.ID(rv))
}

