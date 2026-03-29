// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSNVGGishFeatureEmbedding] class.
var (
	_MLSNVGGishFeatureEmbeddingClass     MLSNVGGishFeatureEmbeddingClass
	_MLSNVGGishFeatureEmbeddingClassOnce sync.Once
)

func getMLSNVGGishFeatureEmbeddingClass() MLSNVGGishFeatureEmbeddingClass {
	_MLSNVGGishFeatureEmbeddingClassOnce.Do(func() {
		_MLSNVGGishFeatureEmbeddingClass = MLSNVGGishFeatureEmbeddingClass{class: objc.GetClass("_MLSNVGGishFeatureEmbedding")}
	})
	return _MLSNVGGishFeatureEmbeddingClass
}

// GetMLSNVGGishFeatureEmbeddingClass returns the class object for _MLSNVGGishFeatureEmbedding.
func GetMLSNVGGishFeatureEmbeddingClass() MLSNVGGishFeatureEmbeddingClass {
	return getMLSNVGGishFeatureEmbeddingClass()
}

type MLSNVGGishFeatureEmbeddingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSNVGGishFeatureEmbeddingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSNVGGishFeatureEmbeddingClass) Alloc() MLSNVGGishFeatureEmbedding {
	rv := objc.Send[MLSNVGGishFeatureEmbedding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSNVGGishFeatureEmbedding.ModelDescription]
//   - [MLSNVGGishFeatureEmbedding.PredictionFromFeaturesOptionsError]
//   - [MLSNVGGishFeatureEmbedding.InitWithModelDescriptionParameterDictionaryError]
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFeatureEmbedding
type MLSNVGGishFeatureEmbedding struct {
	objectivec.Object
}

// MLSNVGGishFeatureEmbeddingFromID constructs a [MLSNVGGishFeatureEmbedding] from an objc.ID.
func MLSNVGGishFeatureEmbeddingFromID(id objc.ID) MLSNVGGishFeatureEmbedding {
	return MLSNVGGishFeatureEmbedding{objectivec.Object{ID: id}}
}
// Ensure MLSNVGGishFeatureEmbedding implements IMLSNVGGishFeatureEmbedding.
var _ IMLSNVGGishFeatureEmbedding = MLSNVGGishFeatureEmbedding{}

// An interface definition for the [MLSNVGGishFeatureEmbedding] class.
//
// # Methods
//
//   - [IMLSNVGGishFeatureEmbedding.ModelDescription]
//   - [IMLSNVGGishFeatureEmbedding.PredictionFromFeaturesOptionsError]
//   - [IMLSNVGGishFeatureEmbedding.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFeatureEmbedding
type IMLSNVGGishFeatureEmbedding interface {
	objectivec.IObject

	// Topic: Methods

	ModelDescription() IMLModelDescription
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNVGGishFeatureEmbedding, error)
}

// Init initializes the instance.
func (m MLSNVGGishFeatureEmbedding) Init() MLSNVGGishFeatureEmbedding {
	rv := objc.Send[MLSNVGGishFeatureEmbedding](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSNVGGishFeatureEmbedding) Autorelease() MLSNVGGishFeatureEmbedding {
	rv := objc.Send[MLSNVGGishFeatureEmbedding](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSNVGGishFeatureEmbedding creates a new MLSNVGGishFeatureEmbedding instance.
func NewMLSNVGGishFeatureEmbedding() MLSNVGGishFeatureEmbedding {
	class := getMLSNVGGishFeatureEmbeddingClass()
	rv := objc.Send[MLSNVGGishFeatureEmbedding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFeatureEmbedding/initWithModelDescription:parameterDictionary:error:
func NewMLSNVGGishFeatureEmbeddingWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNVGGishFeatureEmbedding, error) {
	var errorPtr objc.ID
	instance := getMLSNVGGishFeatureEmbeddingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSNVGGishFeatureEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSNVGGishFeatureEmbeddingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFeatureEmbedding/predictionFromFeatures:options:error:
func (m MLSNVGGishFeatureEmbedding) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFeatureEmbedding/initWithModelDescription:parameterDictionary:error:
func (m MLSNVGGishFeatureEmbedding) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLSNVGGishFeatureEmbedding, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSNVGGishFeatureEmbedding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSNVGGishFeatureEmbeddingFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLSNVGGishFeatureEmbedding/modelDescription
func (m MLSNVGGishFeatureEmbedding) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

