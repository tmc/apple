// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModel] class.
var (
	_MLModelClass     MLModelClass
	_MLModelClassOnce sync.Once
)

func getMLModelClass() MLModelClass {
	_MLModelClassOnce.Do(func() {
		_MLModelClass = MLModelClass{class: objc.GetClass("MLModel")}
	})
	return _MLModelClass
}

// GetMLModelClass returns the class object for MLModel.
func GetMLModelClass() MLModelClass {
	return getMLModelClass()
}

type MLModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelClass) Alloc() MLModel {
	rv := objc.Send[MLModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModel.CancelPredictionRequest]
//   - [MLModel.Classifier]
//   - [MLModel.DebugQuickLookObject]
//   - [MLModel.DecryptSession]
//   - [MLModel.SetDecryptSession]
//   - [MLModel.EnableInstrumentsTracing]
//   - [MLModel.EnableInstrumentsTracingIfNeeded]
//   - [MLModel.ExecutionSchedule]
//   - [MLModel.InternalEngine]
//   - [MLModel.ModelPath]
//   - [MLModel.NeuralNetwork]
//   - [MLModel.NewRequestForModelInputFeaturesOptionsError]
//   - [MLModel.NewRequestWithInputFeaturesOptionsError]
//   - [MLModel.NewRequestWithInputFeaturesUsingStateOptionsError]
//   - [MLModel.NewState]
//   - [MLModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLModel.NewStateWithClientBuffers]
//   - [MLModel.NextPredictionRequestID]
//   - [MLModel.ObjectBoundingBoxOutputDescription]
//   - [MLModel.Pipeline]
//   - [MLModel.PipelineOfPostVisionFeaturePrintModelsFromPipeline]
//   - [MLModel.PredictionEvent]
//   - [MLModel.SetPredictionEvent]
//   - [MLModel.PredictionFromFeaturesCompletionHandler]
//   - [MLModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [MLModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [MLModel.PredictionTypeForKTrace]
//   - [MLModel.PrepareWithCompletionHandler]
//   - [MLModel.PrepareWithConcurrencyHint]
//   - [MLModel.Program]
//   - [MLModel.RecordsPredictionEvent]
//   - [MLModel.Regressor]
//   - [MLModel.SetModelPathModelName]
//   - [MLModel.SignpostID]
//   - [MLModel.SetSignpostID]
//   - [MLModel.SubmitPredictionRequestCompletionHandler]
//   - [MLModel.SupportsConcurrentSubmissions]
//   - [MLModel.Updatable]
//   - [MLModel.VectorizeInputError]
//   - [MLModel.VisionFeaturePrintInfo]
//   - [MLModel.Writable]
//   - [MLModel.InitDescriptionOnlyWithSpecificationConfigurationError]
//   - [MLModel.InitInterfaceAndMetadataWithCompiledArchiveError]
//   - [MLModel.InitWithConfiguration]
//   - [MLModel.InitWithDescription]
//   - [MLModel.InitWithDescriptionConfiguration]
//   - [MLModel.InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//   - [MLModel.Configuration]
//   - [MLModel.SetConfiguration]
//   - [MLModel.DebugDescription]
//   - [MLModel.Description]
//   - [MLModel.Hash]
//   - [MLModel.ModelDescription]
//   - [MLModel.SetModelDescription]
//   - [MLModel.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLModel
type MLModel struct {
	objectivec.Object
}

// MLModelFromID constructs a [MLModel] from an objc.ID.
func MLModelFromID(id objc.ID) MLModel {
	return MLModel{objectivec.Object{ID: id}}
}
// Ensure MLModel implements IMLModel.
var _ IMLModel = MLModel{}

// An interface definition for the [MLModel] class.
//
// # Methods
//
//   - [IMLModel.CancelPredictionRequest]
//   - [IMLModel.Classifier]
//   - [IMLModel.DebugQuickLookObject]
//   - [IMLModel.DecryptSession]
//   - [IMLModel.SetDecryptSession]
//   - [IMLModel.EnableInstrumentsTracing]
//   - [IMLModel.EnableInstrumentsTracingIfNeeded]
//   - [IMLModel.ExecutionSchedule]
//   - [IMLModel.InternalEngine]
//   - [IMLModel.ModelPath]
//   - [IMLModel.NeuralNetwork]
//   - [IMLModel.NewRequestForModelInputFeaturesOptionsError]
//   - [IMLModel.NewRequestWithInputFeaturesOptionsError]
//   - [IMLModel.NewRequestWithInputFeaturesUsingStateOptionsError]
//   - [IMLModel.NewState]
//   - [IMLModel.NewStateForFeatureNamedInitializerBlock]
//   - [IMLModel.NewStateWithClientBuffers]
//   - [IMLModel.NextPredictionRequestID]
//   - [IMLModel.ObjectBoundingBoxOutputDescription]
//   - [IMLModel.Pipeline]
//   - [IMLModel.PipelineOfPostVisionFeaturePrintModelsFromPipeline]
//   - [IMLModel.PredictionEvent]
//   - [IMLModel.SetPredictionEvent]
//   - [IMLModel.PredictionFromFeaturesCompletionHandler]
//   - [IMLModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [IMLModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
//   - [IMLModel.PredictionTypeForKTrace]
//   - [IMLModel.PrepareWithCompletionHandler]
//   - [IMLModel.PrepareWithConcurrencyHint]
//   - [IMLModel.Program]
//   - [IMLModel.RecordsPredictionEvent]
//   - [IMLModel.Regressor]
//   - [IMLModel.SetModelPathModelName]
//   - [IMLModel.SignpostID]
//   - [IMLModel.SetSignpostID]
//   - [IMLModel.SubmitPredictionRequestCompletionHandler]
//   - [IMLModel.SupportsConcurrentSubmissions]
//   - [IMLModel.Updatable]
//   - [IMLModel.VectorizeInputError]
//   - [IMLModel.VisionFeaturePrintInfo]
//   - [IMLModel.Writable]
//   - [IMLModel.InitDescriptionOnlyWithSpecificationConfigurationError]
//   - [IMLModel.InitInterfaceAndMetadataWithCompiledArchiveError]
//   - [IMLModel.InitWithConfiguration]
//   - [IMLModel.InitWithDescription]
//   - [IMLModel.InitWithDescriptionConfiguration]
//   - [IMLModel.InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//   - [IMLModel.Configuration]
//   - [IMLModel.SetConfiguration]
//   - [IMLModel.DebugDescription]
//   - [IMLModel.Description]
//   - [IMLModel.Hash]
//   - [IMLModel.ModelDescription]
//   - [IMLModel.SetModelDescription]
//   - [IMLModel.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLModel
type IMLModel interface {
	objectivec.IObject

	// Topic: Methods

	CancelPredictionRequest(request objectivec.IObject)
	Classifier() objectivec.IObject
	DebugQuickLookObject() objectivec.IObject
	DecryptSession() IMLFairPlayDecryptSession
	SetDecryptSession(value IMLFairPlayDecryptSession)
	EnableInstrumentsTracing()
	EnableInstrumentsTracingIfNeeded()
	ExecutionSchedule() objectivec.IObject
	InternalEngine() objectivec.IObject
	ModelPath() objectivec.IObject
	NeuralNetwork() objectivec.IObject
	NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewRequestWithInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewRequestWithInputFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewState() objectivec.IObject
	NewStateForFeatureNamedInitializerBlock(named objectivec.IObject, block VoidHandler) objectivec.IObject
	NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject
	NextPredictionRequestID() uint64
	ObjectBoundingBoxOutputDescription() objectivec.IObject
	Pipeline() objectivec.IObject
	PipelineOfPostVisionFeaturePrintModelsFromPipeline(pipeline objectivec.IObject) objectivec.IObject
	PredictionEvent() IMLPredictionEvent
	SetPredictionEvent(value IMLPredictionEvent)
	PredictionFromFeaturesCompletionHandler(features objectivec.IObject, handler ErrorHandler)
	PredictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	PredictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	PredictionTypeForKTrace() uint64
	PrepareWithCompletionHandler(handler ErrorHandler)
	PrepareWithConcurrencyHint(hint int64)
	Program() objectivec.IObject
	RecordsPredictionEvent() bool
	Regressor() objectivec.IObject
	SetModelPathModelName(path objectivec.IObject, name objectivec.IObject)
	SignpostID() uint64
	SetSignpostID(value uint64)
	SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler)
	SupportsConcurrentSubmissions() bool
	Updatable() objectivec.IObject
	VectorizeInputError(input objectivec.IObject) (objectivec.IObject, error)
	VisionFeaturePrintInfo() objectivec.IObject
	Writable() objectivec.IObject
	InitDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLModel, error)
	InitInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLModel, error)
	InitWithConfiguration(configuration objectivec.IObject) MLModel
	InitWithDescription(description objectivec.IObject) MLModel
	InitWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModel
	InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModel
	Configuration() IMLModelConfiguration
	SetConfiguration(value IMLModelConfiguration)
	DebugDescription() string
	Description() string
	Hash() uint64
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLModel) Init() MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModel) Autorelease() MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModel creates a new MLModel instance.
func NewMLModel() MLModel {
	class := getMLModelClass()
	rv := objc.Send[MLModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLModel, error) {
	var errorPtr objc.ID
	instance := getMLModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLModel, error) {
	var errorPtr objc.ID
	instance := getMLModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewModelWithConfiguration(configuration objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewModelWithDescription(description objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/cancelPredictionRequest:
func (m MLModel) CancelPredictionRequest(request objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelPredictionRequest:"), request)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/debugQuickLookObject
func (m MLModel) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/enableInstrumentsTracing
func (m MLModel) EnableInstrumentsTracing() {
	objc.Send[objc.ID](m.ID, objc.Sel("enableInstrumentsTracing"))
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/enableInstrumentsTracingIfNeeded
func (m MLModel) EnableInstrumentsTracingIfNeeded() {
	objc.Send[objc.ID](m.ID, objc.Sel("enableInstrumentsTracingIfNeeded"))
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/executionSchedule
func (m MLModel) ExecutionSchedule() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/internalEngine
func (m MLModel) InternalEngine() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("internalEngine"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelPath
func (m MLModel) ModelPath() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelPath"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/newRequestForModel:inputFeatures:options:error:
func (m MLModel) NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestForModel:inputFeatures:options:error:"), model, features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/newRequestWithInputFeatures:options:error:
func (m MLModel) NewRequestWithInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestWithInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/newRequestWithInputFeatures:usingState:options:error:
func (m MLModel) NewRequestWithInputFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestWithInputFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLModel/newState
func (m MLModel) NewState() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newState"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/newStateForFeatureNamed:initializerBlock:
func (m MLModel) NewStateForFeatureNamedInitializerBlock(named objectivec.IObject, block VoidHandler) objectivec.IObject {
_block1, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newStateForFeatureNamed:initializerBlock:"), named, _block1)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/newStateWithClientBuffers:
func (m MLModel) NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newStateWithClientBuffers:"), buffers)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/nextPredictionRequestID
func (m MLModel) NextPredictionRequestID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("nextPredictionRequestID"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/objectBoundingBoxOutputDescription
func (m MLModel) ObjectBoundingBoxOutputDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectBoundingBoxOutputDescription"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/pipelineOfPostVisionFeaturePrintModelsFromPipeline:
func (m MLModel) PipelineOfPostVisionFeaturePrintModelsFromPipeline(pipeline objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pipelineOfPostVisionFeaturePrintModelsFromPipeline:"), pipeline)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:completionHandler:
func (m MLModel) PredictionFromFeaturesCompletionHandler(features objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:completionHandler:"), features, _block1)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:options:completionHandler:
func (m MLModel) PredictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:completionHandler:"), features, options, _block2)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionFromFeatures:usingState:options:completionHandler:
func (m MLModel) PredictionFromFeaturesUsingStateOptionsCompletionHandler(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
_block3, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:completionHandler:"), features, state, options, _block3)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prepareWithCompletionHandler:
func (m MLModel) PrepareWithCompletionHandler(handler ErrorHandler) {
_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("prepareWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/prepareWithConcurrencyHint:
func (m MLModel) PrepareWithConcurrencyHint(hint int64) {
	objc.Send[objc.ID](m.ID, objc.Sel("prepareWithConcurrencyHint:"), hint)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/setModelPath:modelName:
func (m MLModel) SetModelPathModelName(path objectivec.IObject, name objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setModelPath:modelName:"), path, name)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/submitPredictionRequest:completionHandler:
func (m MLModel) SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](m.ID, objc.Sel("submitPredictionRequest:completionHandler:"), request, _block1)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/updatable
func (m MLModel) Updatable() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("updatable"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/vectorizeInput:error:
func (m MLModel) VectorizeInputError(input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vectorizeInput:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLModel/visionFeaturePrintInfo
func (m MLModel) VisionFeaturePrintInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("visionFeaturePrintInfo"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func (m MLModel) InitDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func (m MLModel) InitInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func (m MLModel) InitWithConfiguration(configuration objectivec.IObject) MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func (m MLModel) InitWithDescription(description objectivec.IObject) MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("initWithDescription:"), description)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func (m MLModel) InitWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func (m MLModel) InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModel {
	rv := objc.Send[MLModel](m.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/_compileModelAtURL:options:error:
func (_MLModelClass MLModelClass) _compileModelAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("_compileModelAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// CompileModelAtURLOptionsError is an exported wrapper for the private method _compileModelAtURLOptionsError.
func (_MLModelClass MLModelClass) CompileModelAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	return _MLModelClass._compileModelAtURLOptionsError(url, options)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/availableComputeDevices
func (_MLModelClass MLModelClass) AvailableComputeDevices() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("availableComputeDevices"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/compileModelAtURL:completionHandler:
func (_MLModelClass MLModelClass) CompileModelAtURLCompletionHandler(url foundation.INSURL, handler ErrorHandler) {
_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("compileModelAtURL:completionHandler:"), url, _block1)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/compileModelWithoutAutoreleaseAtURL:options:error:
func (_MLModelClass MLModelClass) CompileModelWithoutAutoreleaseAtURLOptionsError(url foundation.INSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("compileModelWithoutAutoreleaseAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLModel/generateSignpostId
func (_MLModelClass MLModelClass) GenerateSignpostId() uint64 {
	rv := objc.Send[uint64](objc.ID(_MLModelClass.class), objc.Sel("generateSignpostId"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/loadContentsOfURL:configuration:completionHandler:
func (_MLModelClass MLModelClass) LoadContentsOfURLConfigurationCompletionHandler(url foundation.INSURL, configuration objectivec.IObject, handler ErrorHandler) {
_block2, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, _block2)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/maxPredictionsInFlight
func (_MLModelClass MLModelClass) MaxPredictionsInFlight() int64 {
	rv := objc.Send[int64](objc.ID(_MLModelClass.class), objc.Sel("maxPredictionsInFlight"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelWithContentsOfURL:configuration:error:
func (_MLModelClass MLModelClass) ModelWithContentsOfURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelWithContentsOfURL:error:
func (_MLModelClass MLModelClass) ModelWithContentsOfURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionsFromLoopingOverBatch:model:options:error:
func (_MLModelClass MLModelClass) PredictionsFromLoopingOverBatchModelOptionsError(batch objectivec.IObject, model objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("predictionsFromLoopingOverBatch:model:options:error:"), batch, model, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionsFromSubbatchingBatch:maxSubbatchLength:predictionBlock:options:error:
func (_MLModelClass MLModelClass) PredictionsFromSubbatchingBatchMaxSubbatchLengthPredictionBlockOptionsError(batch objectivec.IObject, length int64, block func(), options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("predictionsFromSubbatchingBatch:maxSubbatchLength:predictionBlock:options:error:"), batch, length, block, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/serializeInterfaceAndMetadata:toArchive:error:
func (_MLModelClass MLModelClass) SerializeInterfaceAndMetadataToArchiveError(metadata unsafe.Pointer, archive unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLModelClass.class), objc.Sel("serializeInterfaceAndMetadata:toArchive:error:"), metadata, archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("serializeInterfaceAndMetadata:toArchive:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModel/classifier
func (m MLModel) Classifier() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classifier"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/configuration
func (m MLModel) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModel) SetConfiguration(value IMLModelConfiguration) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfiguration:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/debugDescription
func (m MLModel) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/decryptSession
func (m MLModel) DecryptSession() IMLFairPlayDecryptSession {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("decryptSession"))
	return MLFairPlayDecryptSessionFromID(objc.ID(rv))
}
func (m MLModel) SetDecryptSession(value IMLFairPlayDecryptSession) {
	objc.Send[struct{}](m.ID, objc.Sel("setDecryptSession:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/description
func (m MLModel) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/hash
func (m MLModel) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (m MLModel) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLModel) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelDescription:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/neuralNetwork
func (m MLModel) NeuralNetwork() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("neuralNetwork"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/pipeline
func (m MLModel) Pipeline() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pipeline"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionEvent
func (m MLModel) PredictionEvent() IMLPredictionEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionEvent"))
	return MLPredictionEventFromID(objc.ID(rv))
}
func (m MLModel) SetPredictionEvent(value IMLPredictionEvent) {
	objc.Send[struct{}](m.ID, objc.Sel("setPredictionEvent:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/predictionTypeForKTrace
func (m MLModel) PredictionTypeForKTrace() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/program
func (m MLModel) Program() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("program"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/recordsPredictionEvent
func (m MLModel) RecordsPredictionEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/regressor
func (m MLModel) Regressor() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regressor"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/signpostID
func (m MLModel) SignpostID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}
func (m MLModel) SetSignpostID(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setSignpostID:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/superclass
func (m MLModel) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/supportsConcurrentSubmissions
func (m MLModel) SupportsConcurrentSubmissions() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLModel/writable
func (m MLModel) Writable() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("writable"))
	return objectivec.Object{ID: rv}
}

// NewStateForFeatureNamedInitializerBlockSync is a synchronous wrapper around [MLModel.NewStateForFeatureNamedInitializerBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) NewStateForFeatureNamedInitializerBlockSync(ctx context.Context, named objectivec.IObject) error {
	done := make(chan struct{}, 1)
	m.NewStateForFeatureNamedInitializerBlock(named, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PredictionFromFeatures is a synchronous wrapper around [MLModel.PredictionFromFeaturesCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) PredictionFromFeatures(ctx context.Context, features objectivec.IObject) error {
	done := make(chan error, 1)
	m.PredictionFromFeaturesCompletionHandler(features, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PredictionFromFeaturesOptions is a synchronous wrapper around [MLModel.PredictionFromFeaturesOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) PredictionFromFeaturesOptions(ctx context.Context, features objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m.PredictionFromFeaturesOptionsCompletionHandler(features, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PredictionFromFeaturesUsingStateOptions is a synchronous wrapper around [MLModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) PredictionFromFeaturesUsingStateOptions(ctx context.Context, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m.PredictionFromFeaturesUsingStateOptionsCompletionHandler(features, state, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Prepare is a synchronous wrapper around [MLModel.PrepareWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) Prepare(ctx context.Context) error {
	done := make(chan error, 1)
	m.PrepareWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SubmitPredictionRequest is a synchronous wrapper around [MLModel.SubmitPredictionRequestCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModel) SubmitPredictionRequest(ctx context.Context, request objectivec.IObject) error {
	done := make(chan error, 1)
	m.SubmitPredictionRequestCompletionHandler(request, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// CompileModelAtURL is a synchronous wrapper around [MLModel.CompileModelAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelClass) CompileModelAtURL(ctx context.Context, url foundation.INSURL) error {
	done := make(chan error, 1)
	mc.CompileModelAtURLCompletionHandler(url, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LoadContentsOfURLConfiguration is a synchronous wrapper around [MLModel.LoadContentsOfURLConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelClass) LoadContentsOfURLConfiguration(ctx context.Context, url foundation.INSURL, configuration objectivec.IObject) error {
	done := make(chan error, 1)
	mc.LoadContentsOfURLConfigurationCompletionHandler(url, configuration, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

