// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLWrappedModel] class.
var (
	_MLWrappedModelClass     MLWrappedModelClass
	_MLWrappedModelClassOnce sync.Once
)

func getMLWrappedModelClass() MLWrappedModelClass {
	_MLWrappedModelClassOnce.Do(func() {
		_MLWrappedModelClass = MLWrappedModelClass{class: objc.GetClass("MLWrappedModel")}
	})
	return _MLWrappedModelClass
}

// GetMLWrappedModelClass returns the class object for MLWrappedModel.
func GetMLWrappedModelClass() MLWrappedModelClass {
	return getMLWrappedModelClass()
}

type MLWrappedModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLWrappedModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLWrappedModelClass) Alloc() MLWrappedModel {
	rv := objc.Send[MLWrappedModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLWrappedModel.ClearInnerModelWithReason]
//   - [MLWrappedModel.InnerModel]
//   - [MLWrappedModel.SetInnerModel]
//   - [MLWrappedModel.ParameterValueForKeyError]
//   - [MLWrappedModel.PredictionFromFeaturesError]
//   - [MLWrappedModel.PredictionFromFeaturesOptionsError]
//   - [MLWrappedModel.PredictionsFromBatchError]
//   - [MLWrappedModel.PredictionsFromBatchOptionsError]
//   - [MLWrappedModel.Reason]
//   - [MLWrappedModel.SetReason]
//   - [MLWrappedModel.InitWithInnerModel]
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel
type MLWrappedModel struct {
	MLModel
}

// MLWrappedModelFromID constructs a [MLWrappedModel] from an objc.ID.
func MLWrappedModelFromID(id objc.ID) MLWrappedModel {
	return MLWrappedModel{MLModel: MLModelFromID(id)}
}
// Ensure MLWrappedModel implements IMLWrappedModel.
var _ IMLWrappedModel = MLWrappedModel{}

// An interface definition for the [MLWrappedModel] class.
//
// # Methods
//
//   - [IMLWrappedModel.ClearInnerModelWithReason]
//   - [IMLWrappedModel.InnerModel]
//   - [IMLWrappedModel.SetInnerModel]
//   - [IMLWrappedModel.ParameterValueForKeyError]
//   - [IMLWrappedModel.PredictionFromFeaturesError]
//   - [IMLWrappedModel.PredictionFromFeaturesOptionsError]
//   - [IMLWrappedModel.PredictionsFromBatchError]
//   - [IMLWrappedModel.PredictionsFromBatchOptionsError]
//   - [IMLWrappedModel.Reason]
//   - [IMLWrappedModel.SetReason]
//   - [IMLWrappedModel.InitWithInnerModel]
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel
type IMLWrappedModel interface {
	IMLModel

	// Topic: Methods

	ClearInnerModelWithReason(reason objectivec.IObject)
	InnerModel() IMLModel
	SetInnerModel(value IMLModel)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	Reason() string
	SetReason(value string)
	InitWithInnerModel(model objectivec.IObject) MLWrappedModel
}

// Init initializes the instance.
func (w MLWrappedModel) Init() MLWrappedModel {
	rv := objc.Send[MLWrappedModel](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w MLWrappedModel) Autorelease() MLWrappedModel {
	rv := objc.Send[MLWrappedModel](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLWrappedModel creates a new MLWrappedModel instance.
func NewMLWrappedModel() MLWrappedModel {
	class := getMLWrappedModelClass()
	rv := objc.Send[MLWrappedModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewWrappedModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLWrappedModel, error) {
	var errorPtr objc.ID
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWrappedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWrappedModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewWrappedModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLWrappedModel, error) {
	var errorPtr objc.ID
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWrappedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWrappedModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewWrappedModelWithConfiguration(configuration objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLWrappedModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewWrappedModelWithDescription(description objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLWrappedModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewWrappedModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLWrappedModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/initWithInnerModel:
func NewWrappedModelWithInnerModel(model objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInnerModel:"), model)
	return MLWrappedModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewWrappedModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLWrappedModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/clearInnerModelWithReason:
func (w MLWrappedModel) ClearInnerModelWithReason(reason objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("clearInnerModelWithReason:"), reason)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/parameterValueForKey:error:
func (w MLWrappedModel) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/predictionFromFeatures:error:
func (w MLWrappedModel) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("predictionFromFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/predictionFromFeatures:options:error:
func (w MLWrappedModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/predictionsFromBatch:error:
func (w MLWrappedModel) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("predictionsFromBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/predictionsFromBatch:options:error:
func (w MLWrappedModel) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/initWithInnerModel:
func (w MLWrappedModel) InitWithInnerModel(model objectivec.IObject) MLWrappedModel {
	rv := objc.Send[MLWrappedModel](w.ID, objc.Sel("initWithInnerModel:"), model)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/innerModel
func (w MLWrappedModel) InnerModel() IMLModel {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("innerModel"))
	return MLModelFromID(objc.ID(rv))
}
func (w MLWrappedModel) SetInnerModel(value IMLModel) {
	objc.Send[struct{}](w.ID, objc.Sel("setInnerModel:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLWrappedModel/reason
func (w MLWrappedModel) Reason() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("reason"))
	return foundation.NSStringFromID(rv).String()
}
func (w MLWrappedModel) SetReason(value string) {
	objc.Send[struct{}](w.ID, objc.Sel("setReason:"), objc.String(value))
}

