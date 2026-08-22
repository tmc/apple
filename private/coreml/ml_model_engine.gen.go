// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelEngine] class.
var (
	_MLModelEngineClass     MLModelEngineClass
	_MLModelEngineClassOnce sync.Once
)

func getMLModelEngineClass() MLModelEngineClass {
	_MLModelEngineClassOnce.Do(func() {
		_MLModelEngineClass = MLModelEngineClass{class: objc.GetClass("MLModelEngine")}
	})
	return _MLModelEngineClass
}

// GetMLModelEngineClass returns the class object for MLModelEngine.
func GetMLModelEngineClass() MLModelEngineClass {
	return getMLModelEngineClass()
}

type MLModelEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelEngineClass) Alloc() MLModelEngine {
	rv := objc.SendIfResponds[MLModelEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelEngine.Configuration]
//   - [MLModelEngine.EnableInstrumentsTracing]
//   - [MLModelEngine.ExecutionSchedule]
//   - [MLModelEngine.Metadata]
//   - [MLModelEngine.ModelDescription]
//   - [MLModelEngine.ModelPath]
//   - [MLModelEngine.NewRequestForModelInputFeaturesOptionsError]
//   - [MLModelEngine.ParameterValueForKeyError]
//   - [MLModelEngine.PredictionFromFeaturesError]
//   - [MLModelEngine.PredictionFromFeaturesOptionsError]
//   - [MLModelEngine.PredictionTypeForKTrace]
//   - [MLModelEngine.PredictionsFromBatchError]
//   - [MLModelEngine.PredictionsFromBatchOptionsError]
//   - [MLModelEngine.RecordsPredictionEvent]
//   - [MLModelEngine.SetModelPathModelName]
//   - [MLModelEngine.SignpostID]
//   - [MLModelEngine.SubmitPredictionRequestCompletionHandler]
//   - [MLModelEngine.SupportsConcurrentSubmissions]
//   - [MLModelEngine.VectorizeInputError]
//   - [MLModelEngine.InitWithDescriptionConfiguration]
//   - [MLModelEngine.InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//   - [MLModelEngine.DebugDescription]
//   - [MLModelEngine.Description]
//   - [MLModelEngine.Hash]
//   - [MLModelEngine.Superclass]
type MLModelEngine struct {
	objectivec.Object
}

// MLModelEngineFromID constructs a [MLModelEngine] from an objc.ID.
func MLModelEngineFromID(id objc.ID) MLModelEngine {
	return MLModelEngine{objectivec.Object{ID: id}}
}

// Ensure MLModelEngine implements IMLModelEngine.
var _ IMLModelEngine = MLModelEngine{}

// An interface definition for the [MLModelEngine] class.
//
// # Methods
//
//   - [IMLModelEngine.Configuration]
//   - [IMLModelEngine.EnableInstrumentsTracing]
//   - [IMLModelEngine.ExecutionSchedule]
//   - [IMLModelEngine.Metadata]
//   - [IMLModelEngine.ModelDescription]
//   - [IMLModelEngine.ModelPath]
//   - [IMLModelEngine.NewRequestForModelInputFeaturesOptionsError]
//   - [IMLModelEngine.ParameterValueForKeyError]
//   - [IMLModelEngine.PredictionFromFeaturesError]
//   - [IMLModelEngine.PredictionFromFeaturesOptionsError]
//   - [IMLModelEngine.PredictionTypeForKTrace]
//   - [IMLModelEngine.PredictionsFromBatchError]
//   - [IMLModelEngine.PredictionsFromBatchOptionsError]
//   - [IMLModelEngine.RecordsPredictionEvent]
//   - [IMLModelEngine.SetModelPathModelName]
//   - [IMLModelEngine.SignpostID]
//   - [IMLModelEngine.SubmitPredictionRequestCompletionHandler]
//   - [IMLModelEngine.SupportsConcurrentSubmissions]
//   - [IMLModelEngine.VectorizeInputError]
//   - [IMLModelEngine.InitWithDescriptionConfiguration]
//   - [IMLModelEngine.InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//   - [IMLModelEngine.DebugDescription]
//   - [IMLModelEngine.Description]
//   - [IMLModelEngine.Hash]
//   - [IMLModelEngine.Superclass]
type IMLModelEngine interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() IMLModelConfiguration
	EnableInstrumentsTracing()
	ExecutionSchedule() objectivec.IObject
	Metadata() IMLModelMetadata
	ModelDescription() IMLModelDescription
	ModelPath() objectivec.IObject
	NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionTypeForKTrace() uint64
	PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	RecordsPredictionEvent() bool
	SetModelPathModelName(path objectivec.IObject, name objectivec.IObject)
	SignpostID() uint64
	SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler)
	SupportsConcurrentSubmissions() bool
	VectorizeInputError(input objectivec.IObject) (objectivec.IObject, error)
	InitWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModelEngine
	InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModelEngine
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLModelEngine) Init() MLModelEngine {
	rv := objc.SendIfResponds[MLModelEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelEngine) Autorelease() MLModelEngine {
	rv := objc.SendIfResponds[MLModelEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelEngine creates a new MLModelEngine instance.
func NewMLModelEngine() MLModelEngine {
	class := getMLModelEngineClass()
	rv := objc.SendIfResponds[MLModelEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModelEngine {
	instance := getMLModelEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLModelEngineFromID(rv)
}

func NewModelEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModelEngine {
	instance := getMLModelEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLModelEngineFromID(rv)
}

func (m MLModelEngine) EnableInstrumentsTracing() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("enableInstrumentsTracing"))
}
func (m MLModelEngine) ExecutionSchedule() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("executionSchedule"))
	return objectivec.Object{ID: rv}
}
func (m MLModelEngine) ModelPath() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelPath"))
	return objectivec.Object{ID: rv}
}
func (m MLModelEngine) NewRequestForModelInputFeaturesOptionsError(model objectivec.IObject, features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestForModel:inputFeatures:options:error:"), model, features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) SetModelPathModelName(path objectivec.IObject, name objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setModelPath:modelName:"), path, name)
}
func (m MLModelEngine) SubmitPredictionRequestCompletionHandler(request objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("submitPredictionRequest:completionHandler:"), request, _block1)
}
func (m MLModelEngine) VectorizeInputError(input objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vectorizeInput:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLModelEngine) InitWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLModelEngine {
	rv := objc.SendIfResponds[MLModelEngine](m.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return rv
}
func (m MLModelEngine) InitWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLModelEngine {
	rv := objc.SendIfResponds[MLModelEngine](m.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return rv
}

func (m MLModelEngine) Configuration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLModelEngine) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelEngine) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelEngine) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLModelEngine) Metadata() IMLModelMetadata {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}
func (m MLModelEngine) ModelDescription() IMLModelDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLModelEngine) PredictionTypeForKTrace() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}
func (m MLModelEngine) RecordsPredictionEvent() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}
func (m MLModelEngine) SignpostID() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}
func (m MLModelEngine) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLModelEngine) SupportsConcurrentSubmissions() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}

// SubmitPredictionRequest is a synchronous wrapper around [MLModelEngine.SubmitPredictionRequestCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelEngine) SubmitPredictionRequest(ctx context.Context, request objectivec.IObject) error {
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
