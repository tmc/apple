// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNLPSentenceClassifierModel] class.
var (
	_MLNLPSentenceClassifierModelClass     MLNLPSentenceClassifierModelClass
	_MLNLPSentenceClassifierModelClassOnce sync.Once
)

func getMLNLPSentenceClassifierModelClass() MLNLPSentenceClassifierModelClass {
	_MLNLPSentenceClassifierModelClassOnce.Do(func() {
		_MLNLPSentenceClassifierModelClass = MLNLPSentenceClassifierModelClass{class: objc.GetClass("_MLNLPSentenceClassifierModel")}
	})
	return _MLNLPSentenceClassifierModelClass
}

// GetMLNLPSentenceClassifierModelClass returns the class object for _MLNLPSentenceClassifierModel.
func GetMLNLPSentenceClassifierModelClass() MLNLPSentenceClassifierModelClass {
	return getMLNLPSentenceClassifierModelClass()
}

type MLNLPSentenceClassifierModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNLPSentenceClassifierModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNLPSentenceClassifierModelClass) Alloc() MLNLPSentenceClassifierModel {
	rv := objc.Send[MLNLPSentenceClassifierModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNLPSentenceClassifierModel.ModelDescription]
//   - [MLNLPSentenceClassifierModel.SetModelDescription]
//   - [MLNLPSentenceClassifierModel.PredictionFromFeaturesOptionsError]
//   - [MLNLPSentenceClassifierModel.InitWithModelDescriptionParameterDictionaryError]
// See: https://developer.apple.com/documentation/CoreML/_MLNLPSentenceClassifierModel
type MLNLPSentenceClassifierModel struct {
	objectivec.Object
}

// MLNLPSentenceClassifierModelFromID constructs a [MLNLPSentenceClassifierModel] from an objc.ID.
func MLNLPSentenceClassifierModelFromID(id objc.ID) MLNLPSentenceClassifierModel {
	return MLNLPSentenceClassifierModel{objectivec.Object{ID: id}}
}
// Ensure MLNLPSentenceClassifierModel implements IMLNLPSentenceClassifierModel.
var _ IMLNLPSentenceClassifierModel = MLNLPSentenceClassifierModel{}

// An interface definition for the [MLNLPSentenceClassifierModel] class.
//
// # Methods
//
//   - [IMLNLPSentenceClassifierModel.ModelDescription]
//   - [IMLNLPSentenceClassifierModel.SetModelDescription]
//   - [IMLNLPSentenceClassifierModel.PredictionFromFeaturesOptionsError]
//   - [IMLNLPSentenceClassifierModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLNLPSentenceClassifierModel
type IMLNLPSentenceClassifierModel interface {
	objectivec.IObject

	// Topic: Methods

	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLNLPSentenceClassifierModel, error)
}

// Init initializes the instance.
func (m MLNLPSentenceClassifierModel) Init() MLNLPSentenceClassifierModel {
	rv := objc.Send[MLNLPSentenceClassifierModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNLPSentenceClassifierModel) Autorelease() MLNLPSentenceClassifierModel {
	rv := objc.Send[MLNLPSentenceClassifierModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNLPSentenceClassifierModel creates a new MLNLPSentenceClassifierModel instance.
func NewMLNLPSentenceClassifierModel() MLNLPSentenceClassifierModel {
	class := getMLNLPSentenceClassifierModelClass()
	rv := objc.Send[MLNLPSentenceClassifierModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLNLPSentenceClassifierModel/initWithModelDescription:parameterDictionary:error:
func NewMLNLPSentenceClassifierModelWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLNLPSentenceClassifierModel, error) {
	var errorPtr objc.ID
	instance := getMLNLPSentenceClassifierModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNLPSentenceClassifierModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPSentenceClassifierModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/_MLNLPSentenceClassifierModel/predictionFromFeatures:options:error:
func (m MLNLPSentenceClassifierModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/_MLNLPSentenceClassifierModel/initWithModelDescription:parameterDictionary:error:
func (m MLNLPSentenceClassifierModel) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLNLPSentenceClassifierModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNLPSentenceClassifierModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNLPSentenceClassifierModelFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLNLPSentenceClassifierModel/modelDescription
func (m MLNLPSentenceClassifierModel) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLNLPSentenceClassifierModel) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelDescription:"), value)
}

