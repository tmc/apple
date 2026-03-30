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

// The class instance for the [MLAppleWordTagger] class.
var (
	_MLAppleWordTaggerClass     MLAppleWordTaggerClass
	_MLAppleWordTaggerClassOnce sync.Once
)

func getMLAppleWordTaggerClass() MLAppleWordTaggerClass {
	_MLAppleWordTaggerClassOnce.Do(func() {
		_MLAppleWordTaggerClass = MLAppleWordTaggerClass{class: objc.GetClass("MLAppleWordTagger")}
	})
	return _MLAppleWordTaggerClass
}

// GetMLAppleWordTaggerClass returns the class object for MLAppleWordTagger.
func GetMLAppleWordTaggerClass() MLAppleWordTaggerClass {
	return getMLAppleWordTaggerClass()
}

type MLAppleWordTaggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleWordTaggerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleWordTaggerClass) Alloc() MLAppleWordTagger {
	rv := objc.Send[MLAppleWordTagger](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleWordTagger.Parameters]
//   - [MLAppleWordTagger.PredictionFromFeaturesOptionsError]
//   - [MLAppleWordTagger.InitWithParametersModelDescriptionNlpHandleConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger
type MLAppleWordTagger struct {
	MLModel
}

// MLAppleWordTaggerFromID constructs a [MLAppleWordTagger] from an objc.ID.
func MLAppleWordTaggerFromID(id objc.ID) MLAppleWordTagger {
	return MLAppleWordTagger{MLModel: MLModelFromID(id)}
}

// Ensure MLAppleWordTagger implements IMLAppleWordTagger.
var _ IMLAppleWordTagger = MLAppleWordTagger{}

// An interface definition for the [MLAppleWordTagger] class.
//
// # Methods
//
//   - [IMLAppleWordTagger.Parameters]
//   - [IMLAppleWordTagger.PredictionFromFeaturesOptionsError]
//   - [IMLAppleWordTagger.InitWithParametersModelDescriptionNlpHandleConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger
type IMLAppleWordTagger interface {
	IMLModel

	// Topic: Methods

	Parameters() IMLAppleWordTaggerParameters
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleWordTagger, error)
}

// Init initializes the instance.
func (a MLAppleWordTagger) Init() MLAppleWordTagger {
	rv := objc.Send[MLAppleWordTagger](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleWordTagger) Autorelease() MLAppleWordTagger {
	rv := objc.Send[MLAppleWordTagger](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleWordTagger creates a new MLAppleWordTagger instance.
func NewMLAppleWordTagger() MLAppleWordTagger {
	class := getMLAppleWordTaggerClass()
	rv := objc.Send[MLAppleWordTagger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewAppleWordTaggerDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleWordTagger, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTagger{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewAppleWordTaggerInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleWordTagger, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTagger{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewAppleWordTaggerWithConfiguration(configuration objectivec.IObject) MLAppleWordTagger {
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleWordTaggerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewAppleWordTaggerWithDescription(description objectivec.IObject) MLAppleWordTagger {
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleWordTaggerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewAppleWordTaggerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleWordTagger {
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleWordTaggerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewAppleWordTaggerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleWordTagger {
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleWordTaggerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger/initWithParameters:modelDescription:nlpHandle:configuration:error:
func NewAppleWordTaggerWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleWordTagger, error) {
	var errorPtr objc.ID
	instance := getMLAppleWordTaggerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTagger{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger/predictionFromFeatures:options:error:
func (a MLAppleWordTagger) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger/initWithParameters:modelDescription:nlpHandle:configuration:error:
func (a MLAppleWordTagger) InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleWordTagger, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleWordTagger{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleWordTaggerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger/loadModelFromSpecification:configuration:error:
func (_MLAppleWordTaggerClass MLAppleWordTaggerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleWordTaggerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger/saveAppleWordTaggingModelToURL:wordTaggerParameters:error:
func (_MLAppleWordTaggerClass MLAppleWordTaggerClass) SaveAppleWordTaggingModelToURLWordTaggerParametersError(url foundation.INSURL, parameters objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLAppleWordTaggerClass.class), objc.Sel("saveAppleWordTaggingModelToURL:wordTaggerParameters:error:"), url, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveAppleWordTaggingModelToURL:wordTaggerParameters:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleWordTagger/parameters
func (a MLAppleWordTagger) Parameters() IMLAppleWordTaggerParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameters"))
	return MLAppleWordTaggerParametersFromID(objc.ID(rv))
}
