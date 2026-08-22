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

// The class instance for the [MLE5Engine] class.
var (
	_MLE5EngineClass     MLE5EngineClass
	_MLE5EngineClassOnce sync.Once
)

func getMLE5EngineClass() MLE5EngineClass {
	_MLE5EngineClassOnce.Do(func() {
		_MLE5EngineClass = MLE5EngineClass{class: objc.GetClass("MLE5Engine")}
	})
	return _MLE5EngineClass
}

// GetMLE5EngineClass returns the class object for MLE5Engine.
func GetMLE5EngineClass() MLE5EngineClass {
	return getMLE5EngineClass()
}

type MLE5EngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5EngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5EngineClass) Alloc() MLE5Engine {
	rv := objc.SendIfResponds[MLE5Engine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5Engine._classProbabilitiesInOutputFeaturesError]
//   - [MLE5Engine._classifierResultFromOutputFeaturesClassifyTopKError]
//   - [MLE5Engine._cleanUpAndReconfigureStreamForInputFeaturesError]
//   - [MLE5Engine._cleanUpStream]
//   - [MLE5Engine._conformInputFeaturesError]
//   - [MLE5Engine._conformStateError]
//   - [MLE5Engine._extractSupportFromBackendDict]
//   - [MLE5Engine._extractSupportedComputeUnitFromString]
//   - [MLE5Engine._newRequestForModelInputFeaturesUsingStateOptionsError]
//   - [MLE5Engine._outputFeaturesByAddingClassifierResultToClassifyTopKError]
//   - [MLE5Engine._postProcessingForOutputsOptionsError]
//   - [MLE5Engine._predictionFromFeaturesOptionsCompletionHandler]
//   - [MLE5Engine._predictionFromFeaturesOptionsError]
//   - [MLE5Engine._predictionFromFeaturesStreamOptionsError]
//   - [MLE5Engine._predictionFromFeaturesUsingStateOptionsError]
//   - [MLE5Engine._probabilityDictionaryWithMultiArrayClassifyTopK]
//   - [MLE5Engine._totalRuntimeInMilliSecondsFromE5AnalyticsDictionary]
//   - [MLE5Engine._trimQuotesFromBackendName]
//   - [MLE5Engine._validateStreamReuseExpectationError]
//   - [MLE5Engine.BatchMaxInFlightSem]
//   - [MLE5Engine.ClassLabels]
//   - [MLE5Engine.ClassLabelsSharedKey]
//   - [MLE5Engine.ClassProbabilitiesFeatureName]
//   - [MLE5Engine.ClassifyOptionsError]
//   - [MLE5Engine.CompilerVersionInfo]
//   - [MLE5Engine.EvaluateFunctionArgumentsError]
//   - [MLE5Engine.FunctionName]
//   - [MLE5Engine.InputFeatureConformer]
//   - [MLE5Engine.NewContextAndReturnError]
//   - [MLE5Engine.NewRequestForModelInputFeaturesUsingStateOptionsError]
//   - [MLE5Engine.NewStateWithClientBuffers]
//   - [MLE5Engine.OperationPool]
//   - [MLE5Engine.PredictionFromFeaturesUsingStateOptionsError]
//   - [MLE5Engine.PrepareWithConcurrencyHintError]
//   - [MLE5Engine.ProgramLibrary]
//   - [MLE5Engine.SerializedMILText]
//   - [MLE5Engine.StateFeatureConformer]
//   - [MLE5Engine.StreamPool]
//   - [MLE5Engine.InitWithContainerConfigurationError]
//   - [MLE5Engine.InitWithProgramLibraryModelDescriptionConfigurationFunctionNameClassProbabilitiesFeatureNameOptionalInputDefaultValuesCompilerVersionInfo]
type MLE5Engine struct {
	MLModelEngine
}

// MLE5EngineFromID constructs a [MLE5Engine] from an objc.ID.
func MLE5EngineFromID(id objc.ID) MLE5Engine {
	return MLE5Engine{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLE5Engine implements IMLE5Engine.
var _ IMLE5Engine = MLE5Engine{}

// An interface definition for the [MLE5Engine] class.
//
// # Methods
//
//   - [IMLE5Engine._classProbabilitiesInOutputFeaturesError]
//   - [IMLE5Engine._classifierResultFromOutputFeaturesClassifyTopKError]
//   - [IMLE5Engine._cleanUpAndReconfigureStreamForInputFeaturesError]
//   - [IMLE5Engine._cleanUpStream]
//   - [IMLE5Engine._conformInputFeaturesError]
//   - [IMLE5Engine._conformStateError]
//   - [IMLE5Engine._extractSupportFromBackendDict]
//   - [IMLE5Engine._extractSupportedComputeUnitFromString]
//   - [IMLE5Engine._newRequestForModelInputFeaturesUsingStateOptionsError]
//   - [IMLE5Engine._outputFeaturesByAddingClassifierResultToClassifyTopKError]
//   - [IMLE5Engine._postProcessingForOutputsOptionsError]
//   - [IMLE5Engine._predictionFromFeaturesOptionsCompletionHandler]
//   - [IMLE5Engine._predictionFromFeaturesOptionsError]
//   - [IMLE5Engine._predictionFromFeaturesStreamOptionsError]
//   - [IMLE5Engine._predictionFromFeaturesUsingStateOptionsError]
//   - [IMLE5Engine._probabilityDictionaryWithMultiArrayClassifyTopK]
//   - [IMLE5Engine._totalRuntimeInMilliSecondsFromE5AnalyticsDictionary]
//   - [IMLE5Engine._trimQuotesFromBackendName]
//   - [IMLE5Engine._validateStreamReuseExpectationError]
//   - [IMLE5Engine.BatchMaxInFlightSem]
//   - [IMLE5Engine.ClassLabels]
//   - [IMLE5Engine.ClassLabelsSharedKey]
//   - [IMLE5Engine.ClassProbabilitiesFeatureName]
//   - [IMLE5Engine.ClassifyOptionsError]
//   - [IMLE5Engine.CompilerVersionInfo]
//   - [IMLE5Engine.EvaluateFunctionArgumentsError]
//   - [IMLE5Engine.FunctionName]
//   - [IMLE5Engine.InputFeatureConformer]
//   - [IMLE5Engine.NewContextAndReturnError]
//   - [IMLE5Engine.NewRequestForModelInputFeaturesUsingStateOptionsError]
//   - [IMLE5Engine.NewStateWithClientBuffers]
//   - [IMLE5Engine.OperationPool]
//   - [IMLE5Engine.PredictionFromFeaturesUsingStateOptionsError]
//   - [IMLE5Engine.PrepareWithConcurrencyHintError]
//   - [IMLE5Engine.ProgramLibrary]
//   - [IMLE5Engine.SerializedMILText]
//   - [IMLE5Engine.StateFeatureConformer]
//   - [IMLE5Engine.StreamPool]
//   - [IMLE5Engine.InitWithContainerConfigurationError]
//   - [IMLE5Engine.InitWithProgramLibraryModelDescriptionConfigurationFunctionNameClassProbabilitiesFeatureNameOptionalInputDefaultValuesCompilerVersionInfo]
type IMLE5Engine interface {
	IMLModelEngine

	// Topic: Methods

	_classProbabilitiesInOutputFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	_classifierResultFromOutputFeaturesClassifyTopKError(features objectivec.IObject, k uint64) (objectivec.IObject, error)
	_cleanUpAndReconfigureStreamForInputFeaturesError(stream objectivec.IObject, features objectivec.IObject) (bool, error)
	_cleanUpStream(stream objectivec.IObject)
	_conformInputFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	_conformStateError(state objectivec.IObject) (objectivec.IObject, error)
	_extractSupportFromBackendDict(dict objectivec.IObject) uint64
	_extractSupportedComputeUnitFromString(string_ objectivec.IObject) uint64
	_newRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_outputFeaturesByAddingClassifierResultToClassifyTopKError(to objectivec.IObject, k uint64) (objectivec.IObject, error)
	_postProcessingForOutputsOptionsError(outputs objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_predictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler)
	_predictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_predictionFromFeaturesStreamOptionsError(features objectivec.IObject, stream objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_predictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	_probabilityDictionaryWithMultiArrayClassifyTopK(array objectivec.IObject, k int64) objectivec.IObject
	_totalRuntimeInMilliSecondsFromE5AnalyticsDictionary(dictionary objectivec.IObject) float64
	_trimQuotesFromBackendName(name objectivec.IObject) objectivec.IObject
	_validateStreamReuseExpectationError(reuse bool, expectation objectivec.IObject) (bool, error)
	BatchMaxInFlightSem() objectivec.Object
	ClassLabels() objectivec.IObject
	ClassLabelsSharedKey() objectivec.IObject
	ClassProbabilitiesFeatureName() string
	ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	CompilerVersionInfo() IMLVersionInfo
	EvaluateFunctionArgumentsError(function objectivec.IObject, arguments objectivec.IObject) (objectivec.IObject, error)
	FunctionName() string
	InputFeatureConformer() IMLFeatureProviderConformer
	NewContextAndReturnError() (objectivec.IObject, error)
	NewRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject
	OperationPool() unsafe.Pointer
	PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PrepareWithConcurrencyHintError(hint int64) (bool, error)
	ProgramLibrary() IMLE5ProgramLibrary
	SerializedMILText() string
	StateFeatureConformer() IMLFeatureProviderConformer
	StreamPool() IMLE5ExecutionStreamPool
	InitWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLE5Engine, error)
	InitWithProgramLibraryModelDescriptionConfigurationFunctionNameClassProbabilitiesFeatureNameOptionalInputDefaultValuesCompilerVersionInfo(library objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, values objectivec.IObject, info objectivec.IObject) MLE5Engine
}

// Init initializes the instance.
func (m MLE5Engine) Init() MLE5Engine {
	rv := objc.SendIfResponds[MLE5Engine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLE5Engine) Autorelease() MLE5Engine {
	rv := objc.SendIfResponds[MLE5Engine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5Engine creates a new MLE5Engine instance.
func NewMLE5Engine() MLE5Engine {
	class := getMLE5EngineClass()
	rv := objc.SendIfResponds[MLE5Engine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewE5EngineWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLE5Engine, error) {
	var errorPtr objc.ID
	instance := getMLE5EngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLE5Engine{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLE5Engine{}, objc.ErrInitFailed
	}
	return MLE5EngineFromID(rv), nil
}

func NewE5EngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLE5Engine {
	instance := getMLE5EngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLE5EngineFromID(rv)
}

func NewE5EngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLE5Engine {
	instance := getMLE5EngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLE5EngineFromID(rv)
}

func NewE5EngineWithProgramLibraryModelDescriptionConfigurationFunctionNameClassProbabilitiesFeatureNameOptionalInputDefaultValuesCompilerVersionInfo(library objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, values objectivec.IObject, info objectivec.IObject) MLE5Engine {
	instance := getMLE5EngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:modelDescription:configuration:functionName:classProbabilitiesFeatureName:optionalInputDefaultValues:compilerVersionInfo:"), library, description, configuration, name, name2, values, info)
	return MLE5EngineFromID(rv)
}

func (m MLE5Engine) _classProbabilitiesInOutputFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_classProbabilitiesInOutputFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ClassProbabilitiesInOutputFeaturesError is an exported wrapper for the private method _classProbabilitiesInOutputFeaturesError.
func (m MLE5Engine) ClassProbabilitiesInOutputFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_classProbabilitiesInOutputFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_classProbabilitiesInOutputFeatures:error:"}
		return nil, err
	}
	return m._classProbabilitiesInOutputFeaturesError(features)
}

// CanClassProbabilitiesInOutputFeaturesError reports whether the receiver responds to the private selector _classProbabilitiesInOutputFeatures:error:.
func (m MLE5Engine) CanClassProbabilitiesInOutputFeaturesError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_classProbabilitiesInOutputFeatures:error:"))
}
func (m MLE5Engine) _classifierResultFromOutputFeaturesClassifyTopKError(features objectivec.IObject, k uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_classifierResultFromOutputFeatures:classifyTopK:error:"), features, k, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ClassifierResultFromOutputFeaturesClassifyTopKError is an exported wrapper for the private method _classifierResultFromOutputFeaturesClassifyTopKError.
func (m MLE5Engine) ClassifierResultFromOutputFeaturesClassifyTopKError(features objectivec.IObject, k uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_classifierResultFromOutputFeatures:classifyTopK:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_classifierResultFromOutputFeatures:classifyTopK:error:"}
		return nil, err
	}
	return m._classifierResultFromOutputFeaturesClassifyTopKError(features, k)
}

// CanClassifierResultFromOutputFeaturesClassifyTopKError reports whether the receiver responds to the private selector _classifierResultFromOutputFeatures:classifyTopK:error:.
func (m MLE5Engine) CanClassifierResultFromOutputFeaturesClassifyTopKError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_classifierResultFromOutputFeatures:classifyTopK:error:"))
}
func (m MLE5Engine) _cleanUpAndReconfigureStreamForInputFeaturesError(stream objectivec.IObject, features objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_cleanUpAndReconfigureStream:forInputFeatures:error:"), stream, features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_cleanUpAndReconfigureStream:forInputFeatures:error: returned NO with nil NSError")
	}
	return rv, nil

}

// CleanUpAndReconfigureStreamForInputFeaturesError is an exported wrapper for the private method _cleanUpAndReconfigureStreamForInputFeaturesError.
func (m MLE5Engine) CleanUpAndReconfigureStreamForInputFeaturesError(stream objectivec.IObject, features objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_cleanUpAndReconfigureStream:forInputFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_cleanUpAndReconfigureStream:forInputFeatures:error:"}
		return false, err
	}
	return m._cleanUpAndReconfigureStreamForInputFeaturesError(stream, features)
}

// CanCleanUpAndReconfigureStreamForInputFeaturesError reports whether the receiver responds to the private selector _cleanUpAndReconfigureStream:forInputFeatures:error:.
func (m MLE5Engine) CanCleanUpAndReconfigureStreamForInputFeaturesError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_cleanUpAndReconfigureStream:forInputFeatures:error:"))
}
func (m MLE5Engine) _cleanUpStream(stream objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_cleanUpStream:"), stream)
}

// CleanUpStream is an exported wrapper for the private method _cleanUpStream.
func (m MLE5Engine) CleanUpStream(stream objectivec.IObject) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_cleanUpStream:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_cleanUpStream:"}
		return err
	}
	m._cleanUpStream(stream)
	return nil
}

// CanCleanUpStream reports whether the receiver responds to the private selector _cleanUpStream:.
func (m MLE5Engine) CanCleanUpStream() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_cleanUpStream:"))
}
func (m MLE5Engine) _conformInputFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_conformInputFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ConformInputFeaturesError is an exported wrapper for the private method _conformInputFeaturesError.
func (m MLE5Engine) ConformInputFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_conformInputFeatures:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_conformInputFeatures:error:"}
		return nil, err
	}
	return m._conformInputFeaturesError(features)
}

// CanConformInputFeaturesError reports whether the receiver responds to the private selector _conformInputFeatures:error:.
func (m MLE5Engine) CanConformInputFeaturesError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_conformInputFeatures:error:"))
}
func (m MLE5Engine) _conformStateError(state objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_conformState:error:"), state, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ConformStateError is an exported wrapper for the private method _conformStateError.
func (m MLE5Engine) ConformStateError(state objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_conformState:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_conformState:error:"}
		return nil, err
	}
	return m._conformStateError(state)
}

// CanConformStateError reports whether the receiver responds to the private selector _conformState:error:.
func (m MLE5Engine) CanConformStateError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_conformState:error:"))
}
func (m MLE5Engine) _extractSupportFromBackendDict(dict objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("_extractSupportFromBackendDict:"), dict)
	return rv
}

// ExtractSupportFromBackendDict is an exported wrapper for the private method _extractSupportFromBackendDict.
func (m MLE5Engine) ExtractSupportFromBackendDict(dict objectivec.IObject) (uint64, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_extractSupportFromBackendDict:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_extractSupportFromBackendDict:"}
		return 0, err
	}
	return m._extractSupportFromBackendDict(dict), nil
}

// CanExtractSupportFromBackendDict reports whether the receiver responds to the private selector _extractSupportFromBackendDict:.
func (m MLE5Engine) CanExtractSupportFromBackendDict() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_extractSupportFromBackendDict:"))
}
func (m MLE5Engine) _extractSupportedComputeUnitFromString(string_ objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("_extractSupportedComputeUnitFromString:"), string_)
	return rv
}

// ExtractSupportedComputeUnitFromString is an exported wrapper for the private method _extractSupportedComputeUnitFromString.
func (m MLE5Engine) ExtractSupportedComputeUnitFromString(string_ objectivec.IObject) (uint64, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_extractSupportedComputeUnitFromString:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_extractSupportedComputeUnitFromString:"}
		return 0, err
	}
	return m._extractSupportedComputeUnitFromString(string_), nil
}

// CanExtractSupportedComputeUnitFromString reports whether the receiver responds to the private selector _extractSupportedComputeUnitFromString:.
func (m MLE5Engine) CanExtractSupportedComputeUnitFromString() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_extractSupportedComputeUnitFromString:"))
}
func (m MLE5Engine) _newRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_newRequestForModel:inputFeatures:usingState:options:error:"), model, features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) _outputFeaturesByAddingClassifierResultToClassifyTopKError(to objectivec.IObject, k uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_outputFeaturesByAddingClassifierResultTo:classifyTopK:error:"), to, k, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// OutputFeaturesByAddingClassifierResultToClassifyTopKError is an exported wrapper for the private method _outputFeaturesByAddingClassifierResultToClassifyTopKError.
func (m MLE5Engine) OutputFeaturesByAddingClassifierResultToClassifyTopKError(to objectivec.IObject, k uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_outputFeaturesByAddingClassifierResultTo:classifyTopK:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_outputFeaturesByAddingClassifierResultTo:classifyTopK:error:"}
		return nil, err
	}
	return m._outputFeaturesByAddingClassifierResultToClassifyTopKError(to, k)
}

// CanOutputFeaturesByAddingClassifierResultToClassifyTopKError reports whether the receiver responds to the private selector _outputFeaturesByAddingClassifierResultTo:classifyTopK:error:.
func (m MLE5Engine) CanOutputFeaturesByAddingClassifierResultToClassifyTopKError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_outputFeaturesByAddingClassifierResultTo:classifyTopK:error:"))
}
func (m MLE5Engine) _postProcessingForOutputsOptionsError(outputs objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_postProcessingForOutputs:options:error:"), outputs, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// PostProcessingForOutputsOptionsError is an exported wrapper for the private method _postProcessingForOutputsOptionsError.
func (m MLE5Engine) PostProcessingForOutputsOptionsError(outputs objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_postProcessingForOutputs:options:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_postProcessingForOutputs:options:error:"}
		return nil, err
	}
	return m._postProcessingForOutputsOptionsError(outputs, options)
}

// CanPostProcessingForOutputsOptionsError reports whether the receiver responds to the private selector _postProcessingForOutputs:options:error:.
func (m MLE5Engine) CanPostProcessingForOutputsOptionsError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_postProcessingForOutputs:options:error:"))
}
func (m MLE5Engine) _predictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler) {
	_block2, _ := NewErrorBlock(handler)
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_predictionFromFeatures:options:completionHandler:"), features, options, _block2)
}

// PredictionFromFeaturesOptionsCompletionHandler is an exported wrapper for the private method _predictionFromFeaturesOptionsCompletionHandler.
func (m MLE5Engine) PredictionFromFeaturesOptionsCompletionHandler(features objectivec.IObject, options objectivec.IObject, handler ErrorHandler) error {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_predictionFromFeatures:options:completionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_predictionFromFeatures:options:completionHandler:"}
		return err
	}
	m._predictionFromFeaturesOptionsCompletionHandler(features, options, handler)
	return nil
}

// CanPredictionFromFeaturesOptionsCompletionHandler reports whether the receiver responds to the private selector _predictionFromFeatures:options:completionHandler:.
func (m MLE5Engine) CanPredictionFromFeaturesOptionsCompletionHandler() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_predictionFromFeatures:options:completionHandler:"))
}
func (m MLE5Engine) _predictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) _predictionFromFeaturesStreamOptionsError(features objectivec.IObject, stream objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_predictionFromFeatures:stream:options:error:"), features, stream, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// PredictionFromFeaturesStreamOptionsError is an exported wrapper for the private method _predictionFromFeaturesStreamOptionsError.
func (m MLE5Engine) PredictionFromFeaturesStreamOptionsError(features objectivec.IObject, stream objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_predictionFromFeatures:stream:options:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_predictionFromFeatures:stream:options:error:"}
		return nil, err
	}
	return m._predictionFromFeaturesStreamOptionsError(features, stream, options)
}

// CanPredictionFromFeaturesStreamOptionsError reports whether the receiver responds to the private selector _predictionFromFeatures:stream:options:error:.
func (m MLE5Engine) CanPredictionFromFeaturesStreamOptionsError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_predictionFromFeatures:stream:options:error:"))
}
func (m MLE5Engine) _predictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) _probabilityDictionaryWithMultiArrayClassifyTopK(array objectivec.IObject, k int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_probabilityDictionaryWithMultiArray:classifyTopK:"), array, k)
	return objectivec.Object{ID: rv}
}

// ProbabilityDictionaryWithMultiArrayClassifyTopK is an exported wrapper for the private method _probabilityDictionaryWithMultiArrayClassifyTopK.
func (m MLE5Engine) ProbabilityDictionaryWithMultiArrayClassifyTopK(array objectivec.IObject, k int64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_probabilityDictionaryWithMultiArray:classifyTopK:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_probabilityDictionaryWithMultiArray:classifyTopK:"}
		return nil, err
	}
	return m._probabilityDictionaryWithMultiArrayClassifyTopK(array, k), nil
}

// CanProbabilityDictionaryWithMultiArrayClassifyTopK reports whether the receiver responds to the private selector _probabilityDictionaryWithMultiArray:classifyTopK:.
func (m MLE5Engine) CanProbabilityDictionaryWithMultiArrayClassifyTopK() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_probabilityDictionaryWithMultiArray:classifyTopK:"))
}
func (m MLE5Engine) _totalRuntimeInMilliSecondsFromE5AnalyticsDictionary(dictionary objectivec.IObject) float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("_totalRuntimeInMilliSecondsFromE5AnalyticsDictionary:"), dictionary)
	return rv
}

// TotalRuntimeInMilliSecondsFromE5AnalyticsDictionary is an exported wrapper for the private method _totalRuntimeInMilliSecondsFromE5AnalyticsDictionary.
func (m MLE5Engine) TotalRuntimeInMilliSecondsFromE5AnalyticsDictionary(dictionary objectivec.IObject) (float64, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_totalRuntimeInMilliSecondsFromE5AnalyticsDictionary:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_totalRuntimeInMilliSecondsFromE5AnalyticsDictionary:"}
		return 0.0, err
	}
	return m._totalRuntimeInMilliSecondsFromE5AnalyticsDictionary(dictionary), nil
}

// CanTotalRuntimeInMilliSecondsFromE5AnalyticsDictionary reports whether the receiver responds to the private selector _totalRuntimeInMilliSecondsFromE5AnalyticsDictionary:.
func (m MLE5Engine) CanTotalRuntimeInMilliSecondsFromE5AnalyticsDictionary() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_totalRuntimeInMilliSecondsFromE5AnalyticsDictionary:"))
}
func (m MLE5Engine) _trimQuotesFromBackendName(name objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("_trimQuotesFromBackendName:"), name)
	return objectivec.Object{ID: rv}
}

// TrimQuotesFromBackendName is an exported wrapper for the private method _trimQuotesFromBackendName.
func (m MLE5Engine) TrimQuotesFromBackendName(name objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_trimQuotesFromBackendName:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_trimQuotesFromBackendName:"}
		return nil, err
	}
	return m._trimQuotesFromBackendName(name), nil
}

// CanTrimQuotesFromBackendName reports whether the receiver responds to the private selector _trimQuotesFromBackendName:.
func (m MLE5Engine) CanTrimQuotesFromBackendName() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_trimQuotesFromBackendName:"))
}
func (m MLE5Engine) _validateStreamReuseExpectationError(reuse bool, expectation objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_validateStreamReuse:expectation:error:"), reuse, expectation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_validateStreamReuse:expectation:error: returned NO with nil NSError")
	}
	return rv, nil

}

// ValidateStreamReuseExpectationError is an exported wrapper for the private method _validateStreamReuseExpectationError.
func (m MLE5Engine) ValidateStreamReuseExpectationError(reuse bool, expectation objectivec.IObject) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_validateStreamReuse:expectation:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_validateStreamReuse:expectation:error:"}
		return false, err
	}
	return m._validateStreamReuseExpectationError(reuse, expectation)
}

// CanValidateStreamReuseExpectationError reports whether the receiver responds to the private selector _validateStreamReuse:expectation:error:.
func (m MLE5Engine) CanValidateStreamReuseExpectationError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_validateStreamReuse:expectation:error:"))
}
func (m MLE5Engine) ClassLabels() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classLabels"))
	return objectivec.Object{ID: rv}
}
func (m MLE5Engine) ClassifyOptionsError(classify objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classify:options:error:"), classify, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) EvaluateFunctionArgumentsError(function objectivec.IObject, arguments objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateFunction:arguments:error:"), function, arguments, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) NewContextAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newContextAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) NewRequestForModelInputFeaturesUsingStateOptionsError(model objectivec.IObject, features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newRequestForModel:inputFeatures:usingState:options:error:"), model, features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) NewStateWithClientBuffers(buffers objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("newStateWithClientBuffers:"), buffers)
	return objectivec.Object{ID: rv}
}
func (m MLE5Engine) PredictionFromFeaturesUsingStateOptionsError(features objectivec.IObject, state objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:usingState:options:error:"), features, state, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLE5Engine) PrepareWithConcurrencyHintError(hint int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("prepareWithConcurrencyHint:error:"), hint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareWithConcurrencyHint:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLE5Engine) InitWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLE5Engine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLE5Engine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLE5EngineFromID(rv), nil

}
func (m MLE5Engine) InitWithProgramLibraryModelDescriptionConfigurationFunctionNameClassProbabilitiesFeatureNameOptionalInputDefaultValuesCompilerVersionInfo(library objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, values objectivec.IObject, info objectivec.IObject) MLE5Engine {
	rv := objc.SendIfResponds[MLE5Engine](m.ID, objc.Sel("initWithProgramLibrary:modelDescription:configuration:functionName:classProbabilitiesFeatureName:optionalInputDefaultValues:compilerVersionInfo:"), library, description, configuration, name, name2, values, info)
	return rv
}

func (_MLE5EngineClass MLE5EngineClass) ContainerClass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](objc.ID(_MLE5EngineClass.class), objc.Sel("containerClass"))
	return objectivec.Class(rv)
}
func (_MLE5EngineClass MLE5EngineClass) LoadModelAssetDescriptionFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive MLModelInputArchiverRef, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLE5EngineClass.class), objc.Sel("loadModelAssetDescriptionFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLE5EngineClass MLE5EngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive MLModelInputArchiverRef, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLE5EngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLE5Engine) BatchMaxInFlightSem() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("batchMaxInFlightSem"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLE5Engine) ClassLabelsSharedKey() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classLabelsSharedKey"))
	return objectivec.Object{ID: rv}
}
func (m MLE5Engine) ClassProbabilitiesFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classProbabilitiesFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5Engine) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
func (m MLE5Engine) FunctionName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5Engine) InputFeatureConformer() IMLFeatureProviderConformer {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputFeatureConformer"))
	return MLFeatureProviderConformerFromID(objc.ID(rv))
}
func (m MLE5Engine) OperationPool() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("operationPool"))
	return rv
}
func (m MLE5Engine) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}
func (m MLE5Engine) SerializedMILText() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("serializedMILText"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLE5Engine) StateFeatureConformer() IMLFeatureProviderConformer {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("stateFeatureConformer"))
	return MLFeatureProviderConformerFromID(objc.ID(rv))
}
func (m MLE5Engine) StreamPool() IMLE5ExecutionStreamPool {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("streamPool"))
	return MLE5ExecutionStreamPoolFromID(objc.ID(rv))
}

// _predictionFromFeaturesOptions is a synchronous wrapper around [MLE5Engine._predictionFromFeaturesOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLE5Engine) _predictionFromFeaturesOptions(ctx context.Context, features objectivec.IObject, options objectivec.IObject) error {
	done := make(chan error, 1)
	m._predictionFromFeaturesOptionsCompletionHandler(features, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
