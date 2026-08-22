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

// The class instance for the [MLAppleTextClassifier] class.
var (
	_MLAppleTextClassifierClass     MLAppleTextClassifierClass
	_MLAppleTextClassifierClassOnce sync.Once
)

func getMLAppleTextClassifierClass() MLAppleTextClassifierClass {
	_MLAppleTextClassifierClassOnce.Do(func() {
		_MLAppleTextClassifierClass = MLAppleTextClassifierClass{class: objc.GetClass("MLAppleTextClassifier")}
	})
	return _MLAppleTextClassifierClass
}

// GetMLAppleTextClassifierClass returns the class object for MLAppleTextClassifier.
func GetMLAppleTextClassifierClass() MLAppleTextClassifierClass {
	return getMLAppleTextClassifierClass()
}

type MLAppleTextClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleTextClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleTextClassifierClass) Alloc() MLAppleTextClassifier {
	rv := objc.SendIfResponds[MLAppleTextClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleTextClassifier.Parameters]
//   - [MLAppleTextClassifier.PredictionFromFeaturesOptionsError]
//   - [MLAppleTextClassifier.InitWithParametersModelDescriptionNlpHandleConfigurationError]
type MLAppleTextClassifier struct {
	MLModel
}

// MLAppleTextClassifierFromID constructs a [MLAppleTextClassifier] from an objc.ID.
func MLAppleTextClassifierFromID(id objc.ID) MLAppleTextClassifier {
	return MLAppleTextClassifier{MLModel: MLModelFromID(id)}
}

// Ensure MLAppleTextClassifier implements IMLAppleTextClassifier.
var _ IMLAppleTextClassifier = MLAppleTextClassifier{}

// An interface definition for the [MLAppleTextClassifier] class.
//
// # Methods
//
//   - [IMLAppleTextClassifier.Parameters]
//   - [IMLAppleTextClassifier.PredictionFromFeaturesOptionsError]
//   - [IMLAppleTextClassifier.InitWithParametersModelDescriptionNlpHandleConfigurationError]
type IMLAppleTextClassifier interface {
	IMLModel

	// Topic: Methods

	Parameters() IMLAppleTextClassifierParameters
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleTextClassifier, error)
}

// Init initializes the instance.
func (m MLAppleTextClassifier) Init() MLAppleTextClassifier {
	rv := objc.SendIfResponds[MLAppleTextClassifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLAppleTextClassifier) Autorelease() MLAppleTextClassifier {
	rv := objc.SendIfResponds[MLAppleTextClassifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleTextClassifier creates a new MLAppleTextClassifier instance.
func NewMLAppleTextClassifier() MLAppleTextClassifier {
	class := getMLAppleTextClassifierClass()
	rv := objc.SendIfResponds[MLAppleTextClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAppleTextClassifierDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleTextClassifier, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleTextClassifier{}, objc.ErrInitFailed
	}
	return MLAppleTextClassifierFromID(rv), nil
}

func NewAppleTextClassifierInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleTextClassifier, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleTextClassifier{}, objc.ErrInitFailed
	}
	return MLAppleTextClassifierFromID(rv), nil
}

func NewAppleTextClassifierWithConfiguration(configuration objectivec.IObject) MLAppleTextClassifier {
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleTextClassifierFromID(rv)
}

func NewAppleTextClassifierWithDescription(description objectivec.IObject) MLAppleTextClassifier {
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleTextClassifierFromID(rv)
}

func NewAppleTextClassifierWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleTextClassifier {
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleTextClassifierFromID(rv)
}

func NewAppleTextClassifierWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleTextClassifier {
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleTextClassifierFromID(rv)
}

func NewAppleTextClassifierWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleTextClassifier, error) {
	var errorPtr objc.ID
	instance := getMLAppleTextClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleTextClassifier{}, objc.ErrInitFailed
	}
	return MLAppleTextClassifierFromID(rv), nil
}

func (m MLAppleTextClassifier) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLAppleTextClassifier) InitWithParametersModelDescriptionNlpHandleConfigurationError(parameters objectivec.IObject, description objectivec.IObject, handle objectivec.IObject, configuration objectivec.IObject) (MLAppleTextClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithParameters:modelDescription:nlpHandle:configuration:error:"), parameters, description, handle, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleTextClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleTextClassifierFromID(rv), nil

}

func (_MLAppleTextClassifierClass MLAppleTextClassifierClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleTextClassifierClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLAppleTextClassifierClass MLAppleTextClassifierClass) SaveAppleTextClassifierModelToURLTextClassifierParametersError(url foundation.NSURL, parameters objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLAppleTextClassifierClass.class), objc.Sel("saveAppleTextClassifierModelToURL:textClassifierParameters:error:"), url, parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveAppleTextClassifierModelToURL:textClassifierParameters:error: returned NO with nil NSError")
	}
	return rv, nil

}

func (m MLAppleTextClassifier) Parameters() IMLAppleTextClassifierParameters {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameters"))
	return MLAppleTextClassifierParametersFromID(objc.ID(rv))
}
