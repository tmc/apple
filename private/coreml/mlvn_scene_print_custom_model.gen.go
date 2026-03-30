// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLVNScenePrintCustomModel] class.
var (
	_MLVNScenePrintCustomModelClass     MLVNScenePrintCustomModelClass
	_MLVNScenePrintCustomModelClassOnce sync.Once
)

func getMLVNScenePrintCustomModelClass() MLVNScenePrintCustomModelClass {
	_MLVNScenePrintCustomModelClassOnce.Do(func() {
		_MLVNScenePrintCustomModelClass = MLVNScenePrintCustomModelClass{class: objc.GetClass("_MLVNScenePrintCustomModel")}
	})
	return _MLVNScenePrintCustomModelClass
}

// GetMLVNScenePrintCustomModelClass returns the class object for _MLVNScenePrintCustomModel.
func GetMLVNScenePrintCustomModelClass() MLVNScenePrintCustomModelClass {
	return getMLVNScenePrintCustomModelClass()
}

type MLVNScenePrintCustomModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLVNScenePrintCustomModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLVNScenePrintCustomModelClass) Alloc() MLVNScenePrintCustomModel {
	rv := objc.Send[MLVNScenePrintCustomModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLVNScenePrintCustomModel.Configuration]
//   - [MLVNScenePrintCustomModel.FeatureValueFromScenePrintElementSize]
//   - [MLVNScenePrintCustomModel.ModelDescription]
//   - [MLVNScenePrintCustomModel.PredictionFromFeaturesOptionsError]
//   - [MLVNScenePrintCustomModel.ScenePrintRequestRevision]
//   - [MLVNScenePrintCustomModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel
type MLVNScenePrintCustomModel struct {
	objectivec.Object
}

// MLVNScenePrintCustomModelFromID constructs a [MLVNScenePrintCustomModel] from an objc.ID.
func MLVNScenePrintCustomModelFromID(id objc.ID) MLVNScenePrintCustomModel {
	return MLVNScenePrintCustomModel{objectivec.Object{ID: id}}
}

// Ensure MLVNScenePrintCustomModel implements IMLVNScenePrintCustomModel.
var _ IMLVNScenePrintCustomModel = MLVNScenePrintCustomModel{}

// An interface definition for the [MLVNScenePrintCustomModel] class.
//
// # Methods
//
//   - [IMLVNScenePrintCustomModel.Configuration]
//   - [IMLVNScenePrintCustomModel.FeatureValueFromScenePrintElementSize]
//   - [IMLVNScenePrintCustomModel.ModelDescription]
//   - [IMLVNScenePrintCustomModel.PredictionFromFeaturesOptionsError]
//   - [IMLVNScenePrintCustomModel.ScenePrintRequestRevision]
//   - [IMLVNScenePrintCustomModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel
type IMLVNScenePrintCustomModel interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IMLModelConfiguration
	FeatureValueFromScenePrintElementSize(print_ objectivec.IObject, size uint64) objectivec.IObject
	ModelDescription() IMLModelDescription
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ScenePrintRequestRevision() uint64
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLVNScenePrintCustomModel, error)
}

// Init initializes the instance.
func (m MLVNScenePrintCustomModel) Init() MLVNScenePrintCustomModel {
	rv := objc.Send[MLVNScenePrintCustomModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLVNScenePrintCustomModel) Autorelease() MLVNScenePrintCustomModel {
	rv := objc.Send[MLVNScenePrintCustomModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLVNScenePrintCustomModel creates a new MLVNScenePrintCustomModel instance.
func NewMLVNScenePrintCustomModel() MLVNScenePrintCustomModel {
	class := getMLVNScenePrintCustomModelClass()
	rv := objc.Send[MLVNScenePrintCustomModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/initWithModelDescription:parameterDictionary:error:
func NewMLVNScenePrintCustomModelWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLVNScenePrintCustomModel, error) {
	var errorPtr objc.ID
	instance := getMLVNScenePrintCustomModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLVNScenePrintCustomModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLVNScenePrintCustomModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/featureValueFromScenePrint:elementSize:
func (m MLVNScenePrintCustomModel) FeatureValueFromScenePrintElementSize(print_ objectivec.IObject, size uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureValueFromScenePrint:elementSize:"), print_, size)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/predictionFromFeatures:options:error:
func (m MLVNScenePrintCustomModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/initWithModelDescription:parameterDictionary:error:
func (m MLVNScenePrintCustomModel) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLVNScenePrintCustomModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLVNScenePrintCustomModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLVNScenePrintCustomModelFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/configuration
func (m MLVNScenePrintCustomModel) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/modelDescription
func (m MLVNScenePrintCustomModel) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNScenePrintCustomModel/scenePrintRequestRevision
func (m MLVNScenePrintCustomModel) ScenePrintRequestRevision() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("scenePrintRequestRevision"))
	return rv
}
