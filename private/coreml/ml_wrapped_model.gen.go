// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLWrappedModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
func (m MLWrappedModel) Init() MLWrappedModel {
	rv := objc.SendIfResponds[MLWrappedModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLWrappedModel) Autorelease() MLWrappedModel {
	rv := objc.SendIfResponds[MLWrappedModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLWrappedModel creates a new MLWrappedModel instance.
func NewMLWrappedModel() MLWrappedModel {
	class := getMLWrappedModelClass()
	rv := objc.SendIfResponds[MLWrappedModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWrappedModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLWrappedModel, error) {
	var errorPtr objc.ID
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWrappedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLWrappedModel{}, objc.ErrInitFailed
	}
	return MLWrappedModelFromID(rv), nil
}

func NewWrappedModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLWrappedModel, error) {
	var errorPtr objc.ID
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWrappedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLWrappedModel{}, objc.ErrInitFailed
	}
	return MLWrappedModelFromID(rv), nil
}

func NewWrappedModelWithConfiguration(configuration objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLWrappedModelFromID(rv)
}

func NewWrappedModelWithDescription(description objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLWrappedModelFromID(rv)
}

func NewWrappedModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLWrappedModelFromID(rv)
}

func NewWrappedModelWithInnerModel(model objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInnerModel:"), model)
	return MLWrappedModelFromID(rv)
}

func NewWrappedModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLWrappedModel {
	instance := getMLWrappedModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLWrappedModelFromID(rv)
}

func (m MLWrappedModel) ClearInnerModelWithReason(reason objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("clearInnerModelWithReason:"), reason)
}
func (m MLWrappedModel) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLWrappedModel) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLWrappedModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLWrappedModel) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLWrappedModel) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLWrappedModel) InitWithInnerModel(model objectivec.IObject) MLWrappedModel {
	rv := objc.SendIfResponds[MLWrappedModel](m.ID, objc.Sel("initWithInnerModel:"), model)
	return rv
}

func (m MLWrappedModel) InnerModel() IMLModel {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("innerModel"))
	return MLModelFromID(objc.ID(rv))
}
func (m MLWrappedModel) SetInnerModel(value IMLModel) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setInnerModel:"), value)
}
func (m MLWrappedModel) Reason() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("reason"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLWrappedModel) SetReason(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setReason:"), objc.String(value))
}
