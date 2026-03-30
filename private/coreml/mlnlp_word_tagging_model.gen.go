// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNLPWordTaggingModel] class.
var (
	_MLNLPWordTaggingModelClass     MLNLPWordTaggingModelClass
	_MLNLPWordTaggingModelClassOnce sync.Once
)

func getMLNLPWordTaggingModelClass() MLNLPWordTaggingModelClass {
	_MLNLPWordTaggingModelClassOnce.Do(func() {
		_MLNLPWordTaggingModelClass = MLNLPWordTaggingModelClass{class: objc.GetClass("_MLNLPWordTaggingModel")}
	})
	return _MLNLPWordTaggingModelClass
}

// GetMLNLPWordTaggingModelClass returns the class object for _MLNLPWordTaggingModel.
func GetMLNLPWordTaggingModelClass() MLNLPWordTaggingModelClass {
	return getMLNLPWordTaggingModelClass()
}

type MLNLPWordTaggingModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNLPWordTaggingModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNLPWordTaggingModelClass) Alloc() MLNLPWordTaggingModel {
	rv := objc.Send[MLNLPWordTaggingModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNLPWordTaggingModel.ModelDescription]
//   - [MLNLPWordTaggingModel.SetModelDescription]
//   - [MLNLPWordTaggingModel.PredictionFromFeaturesOptionsError]
//   - [MLNLPWordTaggingModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLNLPWordTaggingModel
type MLNLPWordTaggingModel struct {
	objectivec.Object
}

// MLNLPWordTaggingModelFromID constructs a [MLNLPWordTaggingModel] from an objc.ID.
func MLNLPWordTaggingModelFromID(id objc.ID) MLNLPWordTaggingModel {
	return MLNLPWordTaggingModel{objectivec.Object{ID: id}}
}

// Ensure MLNLPWordTaggingModel implements IMLNLPWordTaggingModel.
var _ IMLNLPWordTaggingModel = MLNLPWordTaggingModel{}

// An interface definition for the [MLNLPWordTaggingModel] class.
//
// # Methods
//
//   - [IMLNLPWordTaggingModel.ModelDescription]
//   - [IMLNLPWordTaggingModel.SetModelDescription]
//   - [IMLNLPWordTaggingModel.PredictionFromFeaturesOptionsError]
//   - [IMLNLPWordTaggingModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLNLPWordTaggingModel
type IMLNLPWordTaggingModel interface {
	objectivec.IObject

	// Topic: Methods

	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLNLPWordTaggingModel, error)
}

// Init initializes the instance.
func (m MLNLPWordTaggingModel) Init() MLNLPWordTaggingModel {
	rv := objc.Send[MLNLPWordTaggingModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNLPWordTaggingModel) Autorelease() MLNLPWordTaggingModel {
	rv := objc.Send[MLNLPWordTaggingModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNLPWordTaggingModel creates a new MLNLPWordTaggingModel instance.
func NewMLNLPWordTaggingModel() MLNLPWordTaggingModel {
	class := getMLNLPWordTaggingModelClass()
	rv := objc.Send[MLNLPWordTaggingModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLNLPWordTaggingModel/initWithModelDescription:parameterDictionary:error:
func NewMLNLPWordTaggingModelWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLNLPWordTaggingModel, error) {
	var errorPtr objc.ID
	instance := getMLNLPWordTaggingModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNLPWordTaggingModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPWordTaggingModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/_MLNLPWordTaggingModel/predictionFromFeatures:options:error:
func (m MLNLPWordTaggingModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLNLPWordTaggingModel/initWithModelDescription:parameterDictionary:error:
func (m MLNLPWordTaggingModel) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLNLPWordTaggingModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNLPWordTaggingModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPWordTaggingModelFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLNLPWordTaggingModel/modelDescription
func (m MLNLPWordTaggingModel) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLNLPWordTaggingModel) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelDescription:"), value)
}
