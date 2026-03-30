// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleWordEmbedding] class.
var (
	_MLAppleWordEmbeddingClass     MLAppleWordEmbeddingClass
	_MLAppleWordEmbeddingClassOnce sync.Once
)

func getMLAppleWordEmbeddingClass() MLAppleWordEmbeddingClass {
	_MLAppleWordEmbeddingClassOnce.Do(func() {
		_MLAppleWordEmbeddingClass = MLAppleWordEmbeddingClass{class: objc.GetClass("MLAppleWordEmbedding")}
	})
	return _MLAppleWordEmbeddingClass
}

// GetMLAppleWordEmbeddingClass returns the class object for MLAppleWordEmbedding.
func GetMLAppleWordEmbeddingClass() MLAppleWordEmbeddingClass {
	return getMLAppleWordEmbeddingClass()
}

type MLAppleWordEmbeddingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleWordEmbeddingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleWordEmbeddingClass) Alloc() MLAppleWordEmbedding {
	rv := objc.Send[MLAppleWordEmbedding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleWordEmbedding.Parameters]
//   - [MLAppleWordEmbedding.PredictionFromFeaturesOptionsError]
//   - [MLAppleWordEmbedding.InitWithParametersModelDescriptionNlpHandleConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding
type MLAppleWordEmbedding struct {
	MLModel
}

// MLAppleWordEmbeddingFromID constructs a [MLAppleWordEmbedding] from an objc.ID.
func MLAppleWordEmbeddingFromID(id objc.ID) MLAppleWordEmbedding {
	return MLAppleWordEmbedding{MLModel: MLModelFromID(id)}
}

// Ensure MLAppleWordEmbedding implements IMLAppleWordEmbedding.
var _ IMLAppleWordEmbedding = MLAppleWordEmbedding{}

// An interface definition for the [MLAppleWordEmbedding] class.
//
// # Methods
//
//   - [IMLAppleWordEmbedding.Parameters]
//   - [IMLAppleWordEmbedding.PredictionFromFeaturesOptionsError]
//   - [IMLAppleWordEmbedding.InitWithParametersModelDescriptionNlpHandleConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding
type IMLAppleWordEmbedding interface {
	IMLModel

	// Topic: Methods

	Parameters() IMLAppleWordEmbeddingParameters
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleWordEmbedding, error)
}

// Init initializes the instance.
func (a MLAppleWordEmbedding) Init() MLAppleWordEmbedding {
	rv := objc.Send[MLAppleWordEmbedding](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleWordEmbedding) Autorelease() MLAppleWordEmbedding {
	rv := objc.Send[MLAppleWordEmbedding](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleWordEmbedding creates a new MLAppleWordEmbedding instance.
func NewMLAppleWordEmbedding() MLAppleWordEmbedding {
	class := getMLAppleWordEmbeddingClass()
	rv := objc.Send[MLAppleWordEmbedding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewAppleWordEmbeddingDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleWordEmbedding, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewAppleWordEmbeddingInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleWordEmbedding, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewAppleWordEmbeddingWithConfiguration(configuration objectivec.IObject) MLAppleWordEmbedding {
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleWordEmbeddingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewAppleWordEmbeddingWithDescription(description objectivec.IObject) MLAppleWordEmbedding {
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleWordEmbeddingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewAppleWordEmbeddingWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleWordEmbedding {
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleWordEmbeddingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewAppleWordEmbeddingWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleWordEmbedding {
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleWordEmbeddingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding/initWithParameters:modelDescription:nlpHandle:configuration:error:
func NewAppleWordEmbeddingWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleWordEmbedding, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding/predictionFromFeatures:options:error:
func (a MLAppleWordEmbedding) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding/initWithParameters:modelDescription:nlpHandle:configuration:error:
func (a MLAppleWordEmbedding) InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleWordEmbedding, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordEmbeddingFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding/loadModelFromSpecification:configuration:error:
func (_MLAppleWordEmbeddingClass MLAppleWordEmbeddingClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleWordEmbeddingClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding/saveAppleWordEmbeddingModelToURL:wordEmbeddingParameters:error:
func (_MLAppleWordEmbeddingClass MLAppleWordEmbeddingClass) SaveAppleWordEmbeddingModelToURLWordEmbeddingParametersError(url foundation.INSURL, parameters objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLAppleWordEmbeddingClass.class), objc.Sel("saveAppleWordEmbeddingModelToURL:wordEmbeddingParameters:error:"), url, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveAppleWordEmbeddingModelToURL:wordEmbeddingParameters:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordEmbedding/parameters
func (a MLAppleWordEmbedding) Parameters() IMLAppleWordEmbeddingParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameters"))
	return MLAppleWordEmbeddingParametersFromID(objc.ID(rv))
}
