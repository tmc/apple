// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleGazetteer] class.
var (
	_MLAppleGazetteerClass     MLAppleGazetteerClass
	_MLAppleGazetteerClassOnce sync.Once
)

func getMLAppleGazetteerClass() MLAppleGazetteerClass {
	_MLAppleGazetteerClassOnce.Do(func() {
		_MLAppleGazetteerClass = MLAppleGazetteerClass{class: objc.GetClass("MLAppleGazetteer")}
	})
	return _MLAppleGazetteerClass
}

// GetMLAppleGazetteerClass returns the class object for MLAppleGazetteer.
func GetMLAppleGazetteerClass() MLAppleGazetteerClass {
	return getMLAppleGazetteerClass()
}

type MLAppleGazetteerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleGazetteerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleGazetteerClass) Alloc() MLAppleGazetteer {
	rv := objc.Send[MLAppleGazetteer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleGazetteer.Parameters]
//   - [MLAppleGazetteer.PredictionFromFeaturesOptionsError]
//   - [MLAppleGazetteer.InitWithParametersModelDescriptionNlpHandleConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer
type MLAppleGazetteer struct {
	MLModel
}

// MLAppleGazetteerFromID constructs a [MLAppleGazetteer] from an objc.ID.
func MLAppleGazetteerFromID(id objc.ID) MLAppleGazetteer {
	return MLAppleGazetteer{MLModel: MLModelFromID(id)}
}
// Ensure MLAppleGazetteer implements IMLAppleGazetteer.
var _ IMLAppleGazetteer = MLAppleGazetteer{}

// An interface definition for the [MLAppleGazetteer] class.
//
// # Methods
//
//   - [IMLAppleGazetteer.Parameters]
//   - [IMLAppleGazetteer.PredictionFromFeaturesOptionsError]
//   - [IMLAppleGazetteer.InitWithParametersModelDescriptionNlpHandleConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer
type IMLAppleGazetteer interface {
	IMLModel

	// Topic: Methods

	Parameters() IMLAppleGazetteerParameters
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleGazetteer, error)
}

// Init initializes the instance.
func (a MLAppleGazetteer) Init() MLAppleGazetteer {
	rv := objc.Send[MLAppleGazetteer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleGazetteer) Autorelease() MLAppleGazetteer {
	rv := objc.Send[MLAppleGazetteer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleGazetteer creates a new MLAppleGazetteer instance.
func NewMLAppleGazetteer() MLAppleGazetteer {
	class := getMLAppleGazetteerClass()
	rv := objc.Send[MLAppleGazetteer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewAppleGazetteerDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleGazetteer, error) {
	var errorPtr objc.ID
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleGazetteerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewAppleGazetteerInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleGazetteer, error) {
	var errorPtr objc.ID
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleGazetteerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewAppleGazetteerWithConfiguration(configuration objectivec.IObject) MLAppleGazetteer {
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleGazetteerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewAppleGazetteerWithDescription(description objectivec.IObject) MLAppleGazetteer {
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleGazetteerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewAppleGazetteerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleGazetteer {
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleGazetteerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewAppleGazetteerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleGazetteer {
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleGazetteerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer/initWithParameters:modelDescription:nlpHandle:configuration:error:
func NewAppleGazetteerWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleGazetteer, error) {
	var errorPtr objc.ID
	instance := getMLAppleGazetteerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleGazetteerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer/predictionFromFeatures:options:error:
func (a MLAppleGazetteer) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer/initWithParameters:modelDescription:nlpHandle:configuration:error:
func (a MLAppleGazetteer) InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleGazetteer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleGazetteer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleGazetteerFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer/loadModelFromSpecification:configuration:error:
func (_MLAppleGazetteerClass MLAppleGazetteerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleGazetteerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer/saveAppleGazetteerModelToURL:gazetteerParameters:error:
func (_MLAppleGazetteerClass MLAppleGazetteerClass) SaveAppleGazetteerModelToURLGazetteerParametersError(url foundation.INSURL, parameters objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLAppleGazetteerClass.class), objc.Sel("saveAppleGazetteerModelToURL:gazetteerParameters:error:"), url, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveAppleGazetteerModelToURL:gazetteerParameters:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleGazetteer/parameters
func (a MLAppleGazetteer) Parameters() IMLAppleGazetteerParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameters"))
	return MLAppleGazetteerParametersFromID(objc.ID(rv))
}

