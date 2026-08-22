// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

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
//   - [MLModel.Metadata]
//   - [MLModel.ModelPath]
//   - [MLModel.NeuralNetwork]
//   - [MLModel.NewRequestForModelInputFeaturesOptionsError]
//   - [MLModel.NewRequestWithInputFeaturesOptionsError]
//   - [MLModel.NewRequestWithInputFeaturesUsingStateOptionsError]
//   - [MLModel.NewStateForFeatureNamedInitializerBlock]
//   - [MLModel.NewStateWithClientBuffers]
//   - [MLModel.NextPredictionRequestID]
//   - [MLModel.ObjectBoundingBoxOutputDescription]
//   - [MLModel.Pipeline]
//   - [MLModel.PipelineOfPostVisionFeaturePrintModelsFromPipeline]
//   - [MLModel.PredictionEvent]
//   - [MLModel.SetPredictionEvent]
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
//   - [IMLModel.Metadata]
//   - [IMLModel.ModelPath]
//   - [IMLModel.NeuralNetwork]
//   - [IMLModel.NewRequestForModelInputFeaturesOptionsError]
//   - [IMLModel.NewRequestWithInputFeaturesOptionsError]
//   - [IMLModel.NewRequestWithInputFeaturesUsingStateOptionsError]
//   - [IMLModel.NewStateForFeatureNamedInitializerBlock]
//   - [IMLModel.NewStateWithClientBuffers]
//   - [IMLModel.NextPredictionRequestID]
//   - [IMLModel.ObjectBoundingBoxOutputDescription]
//   - [IMLModel.Pipeline]
//   - [IMLModel.PipelineOfPostVisionFeaturePrintModelsFromPipeline]
//   - [IMLModel.PredictionEvent]
//   - [IMLModel.SetPredictionEvent]
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
type IMLModel interface {
	objectivec.IObject

	// Topic: Methods

	CancelPredictionRequest(request objectivec.IObject)
	Classifier() unsafe.Pointer
	DebugQuickLookObject() objectivec.IObject
	DecryptSession() IMLFairPlayDecryptSession
	SetDecryptSession(value IMLFairPlayDecryptSession)
	EnableInstrumentsTracing()
	EnableInstrumentsTracingIfNeeded()
	ExecutionSchedule() objectivec.IObject
	InternalEngine() objectivec.IObject
	Metadata() IMLModelMetadata
	ModelPath() objectivec.IObject
	NeuralNetwork() unsafe.Pointer
	NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewRequestWithInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewRequestWithInputFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewStateForFeatureNamedInitializerBlock(named objectivec.IObject, block VoidHandler) objectivec.IObject
	NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject
	NextPredictionRequestID() uint64
	ObjectBoundingBoxOutputDescription() objectivec.IObject
	Pipeline() unsafe.Pointer
	PipelineOfPostVisionFeaturePrintModelsFromPipeline(pipeline objectivec.IObject) objectivec.IObject
	PredictionEvent() IMLPredictionEvent
	SetPredictionEvent(value IMLPredictionEvent)
	PredictionTypeForKTrace() uint64
	PrepareWithCompletionHandler(handler ErrorHandler)
	PrepareWithConcurrencyHint(hint int64)
	Program() unsafe.Pointer
	RecordsPredictionEvent() bool
	Regressor() unsafe.Pointer
	SetModelPathModelName(path objectivec.IObject, name objectivec.IObject)
	SignpostID() uint64
	SetSignpostID(value uint64)
	SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler)
	SupportsConcurrentSubmissions() bool
	Updatable() objectivec.IObject
	VectorizeInputError(input objectivec.IObject) (objectivec.IObject, error)
	VisionFeaturePrintInfo() objectivec.IObject
	Writable() unsafe.Pointer
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLModel) Init() MLModel {
	rv := objc.SendIfResponds[MLModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModel) Autorelease() MLModel {
	rv := objc.SendIfResponds[MLModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModel creates a new MLModel instance.
func NewMLModel() MLModel {
	class := getMLModelClass()
	rv := objc.SendIfResponds[MLModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLModel, error) {
	var errorPtr objc.ID
	instance := getMLModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModel{}, objc.ErrInitFailed
	}
	return MLModelFromID(rv), nil
}

func NewModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLModel, error) {
	var errorPtr objc.ID
	instance := getMLModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModel{}, objc.ErrInitFailed
	}
	return MLModelFromID(rv), nil
}

func NewModelWithConfiguration(configuration objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLModelFromID(rv)
}

func NewModelWithDescription(description objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLModelFromID(rv)
}

func NewModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLModelFromID(rv)
}

func NewModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModel {
	instance := getMLModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLModelFromID(rv)
}

func (m MLModel) CancelPredictionRequest(request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("cancelPredictionRequest:"), request)
}
func (m MLModel) DebugQuickLookObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) EnableInstrumentsTracing() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("enableInstrumentsTracing"))
}
func (m MLModel) EnableInstrumentsTracingIfNeeded() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("enableInstrumentsTracingIfNeeded"))
}
func (m MLModel) ExecutionSchedule() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) InternalEngine() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("internalEngine"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) ModelPath() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelPath"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestForModel:inputFeatures:options:error:"), model, features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModel) NewRequestWithInputFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestWithInputFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModel) NewRequestWithInputFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestWithInputFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModel) NewStateForFeatureNamedInitializerBlock(named objectivec.IObject, block VoidHandler) objectivec.IObject {
	_block1, _ := NewVoidBlock(block)
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("newStateForFeatureNamed:initializerBlock:"), named, _block1)
	return objectivec.Object{ID: rv}
}
func (m MLModel) NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("newStateWithClientBuffers:"), buffers)
	return objectivec.Object{ID: rv}
}
func (m MLModel) NextPredictionRequestID() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("nextPredictionRequestID"))
	return rv
}
func (m MLModel) ObjectBoundingBoxOutputDescription() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("objectBoundingBoxOutputDescription"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) PipelineOfPostVisionFeaturePrintModelsFromPipeline(pipeline objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("pipelineOfPostVisionFeaturePrintModelsFromPipeline:"), pipeline)
	return objectivec.Object{ID: rv}
}
func (m MLModel) PrepareWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("prepareWithCompletionHandler:"), _block0)
}
func (m MLModel) PrepareWithConcurrencyHint(hint int64) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("prepareWithConcurrencyHint:"), hint)
}
func (m MLModel) SetModelPathModelName(path objectivec.IObject, name objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setModelPath:modelName:"), path, name)
}
func (m MLModel) SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("submitPredictionRequest:completionHandler:"), request, _block1)
}
func (m MLModel) Updatable() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updatable"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) VectorizeInputError(input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vectorizeInput:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModel) VisionFeaturePrintInfo() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("visionFeaturePrintInfo"))
	return objectivec.Object{ID: rv}
}
func (m MLModel) InitDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil

}
func (m MLModel) InitInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelFromID(rv), nil

}
func (m MLModel) InitWithConfiguration(configuration objectivec.IObject) MLModel {
	rv := objc.SendIfResponds[MLModel](m.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}
func (m MLModel) InitWithDescription(description objectivec.IObject) MLModel {
	rv := objc.SendIfResponds[MLModel](m.ID, objc.Sel("initWithDescription:"), description)
	return rv
}
func (m MLModel) InitWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModel {
	rv := objc.SendIfResponds[MLModel](m.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return rv
}
func (m MLModel) InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModel {
	rv := objc.SendIfResponds[MLModel](m.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return rv
}

func (_MLModelClass MLModelClass) _compileModelAtURLOptionsError(url foundation.NSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("_compileModelAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// CompileModelAtURLOptionsError is an exported wrapper for the private method _compileModelAtURLOptionsError.
func (_MLModelClass MLModelClass) CompileModelAtURLOptionsError(url foundation.NSURL, options objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(objc.ID(_MLModelClass.class), objc.Sel("_compileModelAtURL:options:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_compileModelAtURL:options:error:"}
		return nil, err
	}
	return _MLModelClass._compileModelAtURLOptionsError(url, options)
}

// CanCompileModelAtURLOptionsError reports whether the receiver responds to the private selector _compileModelAtURL:options:error:.
func (_MLModelClass MLModelClass) CanCompileModelAtURLOptionsError() bool {
	return objc.RespondsToSelector(objc.ID(_MLModelClass.class), objc.Sel("_compileModelAtURL:options:error:"))
}
func (_MLModelClass MLModelClass) CompileModelAtURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("compileModelAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelClass MLModelClass) CompileModelWithoutAutoreleaseAtURLOptionsError(url foundation.NSURL, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("compileModelWithoutAutoreleaseAtURL:options:error:"), url, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelClass MLModelClass) GenerateSignpostId() uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_MLModelClass.class), objc.Sel("generateSignpostId"))
	return rv
}
func (_MLModelClass MLModelClass) MaxPredictionsInFlight() int64 {
	rv := objc.SendIfResponds[int64](objc.ID(_MLModelClass.class), objc.Sel("maxPredictionsInFlight"))
	return rv
}
func (_MLModelClass MLModelClass) ModelWithContentsOfURLConfigurationError(url foundation.NSURL, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("modelWithContentsOfURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelClass MLModelClass) ModelWithContentsOfURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelClass MLModelClass) PredictionsFromLoopingOverBatchModelOptionsError(batch objectivec.IObject, model objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("predictionsFromLoopingOverBatch:model:options:error:"), batch, model, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLModelClass MLModelClass) PredictionsFromSubbatchingBatchMaxSubbatchLengthPredictionBlockOptionsError(batch objectivec.IObject, length int64, block func(), options objectivec.IObject) (objectivec.IObject, error) {
	_block2, _cleanup2 := NewVoidBlock(block)
	defer _cleanup2()
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLModelClass.class), objc.Sel("predictionsFromSubbatchingBatch:maxSubbatchLength:predictionBlock:options:error:"), batch, length, objc.ID(_block2), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
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

func (m MLModel) Classifier() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("classifier"))
	return rv
}
func (m MLModel) Configuration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModel) SetConfiguration(value IMLModelConfiguration) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setConfiguration:"), value)
}
func (m MLModel) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModel) DecryptSession() IMLFairPlayDecryptSession {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("decryptSession"))
	return MLFairPlayDecryptSessionFromID(objc.ID(rv))
}
func (m MLModel) SetDecryptSession(value IMLFairPlayDecryptSession) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDecryptSession:"), value)
}
func (m MLModel) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModel) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLModel) Metadata() IMLModelMetadata {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}
func (m MLModel) ModelDescription() IMLModelDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLModel) SetModelDescription(value IMLModelDescription) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelDescription:"), value)
}
func (m MLModel) NeuralNetwork() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("neuralNetwork"))
	return rv
}
func (m MLModel) Pipeline() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("pipeline"))
	return rv
}
func (m MLModel) PredictionEvent() IMLPredictionEvent {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("predictionEvent"))
	return MLPredictionEventFromID(objc.ID(rv))
}
func (m MLModel) SetPredictionEvent(value IMLPredictionEvent) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPredictionEvent:"), value)
}
func (m MLModel) PredictionTypeForKTrace() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}
func (m MLModel) Program() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("program"))
	return rv
}
func (m MLModel) RecordsPredictionEvent() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}
func (m MLModel) Regressor() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("regressor"))
	return rv
}
func (m MLModel) SignpostID() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}
func (m MLModel) SetSignpostID(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setSignpostID:"), value)
}
func (m MLModel) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLModel) SupportsConcurrentSubmissions() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
func (m MLModel) Writable() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("writable"))
	return rv
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
