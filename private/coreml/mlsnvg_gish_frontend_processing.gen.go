// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSNVGGishFrontendProcessing] class.
var (
	_MLSNVGGishFrontendProcessingClass     MLSNVGGishFrontendProcessingClass
	_MLSNVGGishFrontendProcessingClassOnce sync.Once
)

func getMLSNVGGishFrontendProcessingClass() MLSNVGGishFrontendProcessingClass {
	_MLSNVGGishFrontendProcessingClassOnce.Do(func() {
		_MLSNVGGishFrontendProcessingClass = MLSNVGGishFrontendProcessingClass{class: objc.GetClass("_MLSNVGGishFrontendProcessing")}
	})
	return _MLSNVGGishFrontendProcessingClass
}

// GetMLSNVGGishFrontendProcessingClass returns the class object for _MLSNVGGishFrontendProcessing.
func GetMLSNVGGishFrontendProcessingClass() MLSNVGGishFrontendProcessingClass {
	return getMLSNVGGishFrontendProcessingClass()
}

type MLSNVGGishFrontendProcessingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSNVGGishFrontendProcessingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSNVGGishFrontendProcessingClass) Alloc() MLSNVGGishFrontendProcessing {
	rv := objc.Send[MLSNVGGishFrontendProcessing](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSNVGGishFrontendProcessing.ModelDescription]
//   - [MLSNVGGishFrontendProcessing.PredictionFromFeaturesOptionsError]
//   - [MLSNVGGishFrontendProcessing.InitWithModelDescriptionParameterDictionaryError]
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFrontendProcessing
type MLSNVGGishFrontendProcessing struct {
	objectivec.Object
}

// MLSNVGGishFrontendProcessingFromID constructs a [MLSNVGGishFrontendProcessing] from an objc.ID.
func MLSNVGGishFrontendProcessingFromID(id objc.ID) MLSNVGGishFrontendProcessing {
	return MLSNVGGishFrontendProcessing{objectivec.Object{ID: id}}
}
// Ensure MLSNVGGishFrontendProcessing implements IMLSNVGGishFrontendProcessing.
var _ IMLSNVGGishFrontendProcessing = MLSNVGGishFrontendProcessing{}

// An interface definition for the [MLSNVGGishFrontendProcessing] class.
//
// # Methods
//
//   - [IMLSNVGGishFrontendProcessing.ModelDescription]
//   - [IMLSNVGGishFrontendProcessing.PredictionFromFeaturesOptionsError]
//   - [IMLSNVGGishFrontendProcessing.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFrontendProcessing
type IMLSNVGGishFrontendProcessing interface {
	objectivec.IObject

	// Topic: Methods

	ModelDescription() IMLModelDescription
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNVGGishFrontendProcessing, error)
}

// Init initializes the instance.
func (m MLSNVGGishFrontendProcessing) Init() MLSNVGGishFrontendProcessing {
	rv := objc.Send[MLSNVGGishFrontendProcessing](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSNVGGishFrontendProcessing) Autorelease() MLSNVGGishFrontendProcessing {
	rv := objc.Send[MLSNVGGishFrontendProcessing](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSNVGGishFrontendProcessing creates a new MLSNVGGishFrontendProcessing instance.
func NewMLSNVGGishFrontendProcessing() MLSNVGGishFrontendProcessing {
	class := getMLSNVGGishFrontendProcessingClass()
	rv := objc.Send[MLSNVGGishFrontendProcessing](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFrontendProcessing/initWithModelDescription:parameterDictionary:error:
func NewMLSNVGGishFrontendProcessingWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNVGGishFrontendProcessing, error) {
	var errorPtr objc.ID
	instance := getMLSNVGGishFrontendProcessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSNVGGishFrontendProcessing{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSNVGGishFrontendProcessingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFrontendProcessing/predictionFromFeatures:options:error:
func (m MLSNVGGishFrontendProcessing) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFrontendProcessing/initWithModelDescription:parameterDictionary:error:
func (m MLSNVGGishFrontendProcessing) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNVGGishFrontendProcessing, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSNVGGishFrontendProcessing{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSNVGGishFrontendProcessingFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFrontendProcessing/modelDescription
func (m MLSNVGGishFrontendProcessing) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

