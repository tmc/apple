// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLVNDetectionPrintCustomModel] class.
var (
	_MLVNDetectionPrintCustomModelClass     MLVNDetectionPrintCustomModelClass
	_MLVNDetectionPrintCustomModelClassOnce sync.Once
)

func getMLVNDetectionPrintCustomModelClass() MLVNDetectionPrintCustomModelClass {
	_MLVNDetectionPrintCustomModelClassOnce.Do(func() {
		_MLVNDetectionPrintCustomModelClass = MLVNDetectionPrintCustomModelClass{class: objc.GetClass("_MLVNDetectionPrintCustomModel")}
	})
	return _MLVNDetectionPrintCustomModelClass
}

// GetMLVNDetectionPrintCustomModelClass returns the class object for _MLVNDetectionPrintCustomModel.
func GetMLVNDetectionPrintCustomModelClass() MLVNDetectionPrintCustomModelClass {
	return getMLVNDetectionPrintCustomModelClass()
}

type MLVNDetectionPrintCustomModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLVNDetectionPrintCustomModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLVNDetectionPrintCustomModelClass) Alloc() MLVNDetectionPrintCustomModel {
	rv := objc.Send[MLVNDetectionPrintCustomModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLVNDetectionPrintCustomModel.Configuration]
//   - [MLVNDetectionPrintCustomModel.DetectionPrintRequestRevision]
//   - [MLVNDetectionPrintCustomModel.ExpectedOutputShapeV1]
//   - [MLVNDetectionPrintCustomModel.FeatureValueFromDetectionPrintFeatureName]
//   - [MLVNDetectionPrintCustomModel.ModelDescription]
//   - [MLVNDetectionPrintCustomModel.PredictionFromFeaturesOptionsError]
//   - [MLVNDetectionPrintCustomModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel
type MLVNDetectionPrintCustomModel struct {
	objectivec.Object
}

// MLVNDetectionPrintCustomModelFromID constructs a [MLVNDetectionPrintCustomModel] from an objc.ID.
func MLVNDetectionPrintCustomModelFromID(id objc.ID) MLVNDetectionPrintCustomModel {
	return MLVNDetectionPrintCustomModel{objectivec.Object{ID: id}}
}

// Ensure MLVNDetectionPrintCustomModel implements IMLVNDetectionPrintCustomModel.
var _ IMLVNDetectionPrintCustomModel = MLVNDetectionPrintCustomModel{}

// An interface definition for the [MLVNDetectionPrintCustomModel] class.
//
// # Methods
//
//   - [IMLVNDetectionPrintCustomModel.Configuration]
//   - [IMLVNDetectionPrintCustomModel.DetectionPrintRequestRevision]
//   - [IMLVNDetectionPrintCustomModel.ExpectedOutputShapeV1]
//   - [IMLVNDetectionPrintCustomModel.FeatureValueFromDetectionPrintFeatureName]
//   - [IMLVNDetectionPrintCustomModel.ModelDescription]
//   - [IMLVNDetectionPrintCustomModel.PredictionFromFeaturesOptionsError]
//   - [IMLVNDetectionPrintCustomModel.InitWithModelDescriptionParameterDictionaryError]
//
// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel
type IMLVNDetectionPrintCustomModel interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IMLModelConfiguration
	DetectionPrintRequestRevision() uint64
	ExpectedOutputShapeV1() foundation.INSDictionary
	FeatureValueFromDetectionPrintFeatureName(print_ objectivec.IObject, name objectivec.IObject) objectivec.IObject
	ModelDescription() IMLModelDescription
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLVNDetectionPrintCustomModel, error)
}

// Init initializes the instance.
func (m MLVNDetectionPrintCustomModel) Init() MLVNDetectionPrintCustomModel {
	rv := objc.Send[MLVNDetectionPrintCustomModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLVNDetectionPrintCustomModel) Autorelease() MLVNDetectionPrintCustomModel {
	rv := objc.Send[MLVNDetectionPrintCustomModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLVNDetectionPrintCustomModel creates a new MLVNDetectionPrintCustomModel instance.
func NewMLVNDetectionPrintCustomModel() MLVNDetectionPrintCustomModel {
	class := getMLVNDetectionPrintCustomModelClass()
	rv := objc.Send[MLVNDetectionPrintCustomModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/initWithModelDescription:parameterDictionary:error:
func NewMLVNDetectionPrintCustomModelWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLVNDetectionPrintCustomModel, error) {
	var errorPtr objc.ID
	instance := getMLVNDetectionPrintCustomModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLVNDetectionPrintCustomModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLVNDetectionPrintCustomModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/featureValueFromDetectionPrint:featureName:
func (m MLVNDetectionPrintCustomModel) FeatureValueFromDetectionPrintFeatureName(print_ objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureValueFromDetectionPrint:featureName:"), print_, name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/predictionFromFeatures:options:error:
func (m MLVNDetectionPrintCustomModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/initWithModelDescription:parameterDictionary:error:
func (m MLVNDetectionPrintCustomModel) InitWithModelDescriptionParameterDictionaryError(description objectivec.IObject, dictionary objectivec.IObject) (MLVNDetectionPrintCustomModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescription:parameterDictionary:error:"), description, dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLVNDetectionPrintCustomModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLVNDetectionPrintCustomModelFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/configuration
func (m MLVNDetectionPrintCustomModel) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/detectionPrintRequestRevision
func (m MLVNDetectionPrintCustomModel) DetectionPrintRequestRevision() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("detectionPrintRequestRevision"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/expectedOutputShapeV1
func (m MLVNDetectionPrintCustomModel) ExpectedOutputShapeV1() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("expectedOutputShapeV1"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/_MLVNDetectionPrintCustomModel/modelDescription
func (m MLVNDetectionPrintCustomModel) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
