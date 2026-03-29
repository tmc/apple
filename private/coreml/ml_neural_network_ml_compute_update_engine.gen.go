// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkMLComputeUpdateEngine] class.
var (
	_MLNeuralNetworkMLComputeUpdateEngineClass     MLNeuralNetworkMLComputeUpdateEngineClass
	_MLNeuralNetworkMLComputeUpdateEngineClassOnce sync.Once
)

func getMLNeuralNetworkMLComputeUpdateEngineClass() MLNeuralNetworkMLComputeUpdateEngineClass {
	_MLNeuralNetworkMLComputeUpdateEngineClassOnce.Do(func() {
		_MLNeuralNetworkMLComputeUpdateEngineClass = MLNeuralNetworkMLComputeUpdateEngineClass{class: objc.GetClass("MLNeuralNetworkMLComputeUpdateEngine")}
	})
	return _MLNeuralNetworkMLComputeUpdateEngineClass
}

// GetMLNeuralNetworkMLComputeUpdateEngineClass returns the class object for MLNeuralNetworkMLComputeUpdateEngine.
func GetMLNeuralNetworkMLComputeUpdateEngineClass() MLNeuralNetworkMLComputeUpdateEngineClass {
	return getMLNeuralNetworkMLComputeUpdateEngineClass()
}

type MLNeuralNetworkMLComputeUpdateEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkMLComputeUpdateEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkMLComputeUpdateEngineClass) Alloc() MLNeuralNetworkMLComputeUpdateEngine {
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNeuralNetworkMLComputeUpdateEngine.BatchSize]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetBatchSize]
//   - [MLNeuralNetworkMLComputeUpdateEngine.CancelUpdate]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ClassLabelToIndexMap]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ClassifierOutputIsSigmoidOutput]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetClassifierOutputIsSigmoidOutput]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ContinueWithUpdate]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetContinueWithUpdate]
//   - [MLNeuralNetworkMLComputeUpdateEngine.FinalLossValue]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetFinalLossValue]
//   - [MLNeuralNetworkMLComputeUpdateEngine.LoadLossTargetName]
//   - [MLNeuralNetworkMLComputeUpdateEngine.LossTargetName]
//   - [MLNeuralNetworkMLComputeUpdateEngine.MlcGraph]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetMlcGraph]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ParameterContainer]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetParameterContainer]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ParameterValueForKey]
//   - [MLNeuralNetworkMLComputeUpdateEngine.PerformInferenceWithOutputNameToLayerMapError]
//   - [MLNeuralNetworkMLComputeUpdateEngine.PerformTrainingWithCallBacksNumberOfEpochsError]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ProgressHandlers]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetProgressHandlers]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ResumeUpdate]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ResumeUpdateWithParameters]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkMLComputeUpdateEngine.ShuffableTrainingData]
//   - [MLNeuralNetworkMLComputeUpdateEngine.SetShuffableTrainingData]
//   - [MLNeuralNetworkMLComputeUpdateEngine.UpdateLearningRateWithValue]
//   - [MLNeuralNetworkMLComputeUpdateEngine.UpdateModelWithData]
//   - [MLNeuralNetworkMLComputeUpdateEngine.UpdateParameters]
//   - [MLNeuralNetworkMLComputeUpdateEngine.WriteToURLError]
//   - [MLNeuralNetworkMLComputeUpdateEngine.InitWithCompiledArchiveNnContainerConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine
type MLNeuralNetworkMLComputeUpdateEngine struct {
	MLNeuralNetworkV1Engine
}

// MLNeuralNetworkMLComputeUpdateEngineFromID constructs a [MLNeuralNetworkMLComputeUpdateEngine] from an objc.ID.
func MLNeuralNetworkMLComputeUpdateEngineFromID(id objc.ID) MLNeuralNetworkMLComputeUpdateEngine {
	return MLNeuralNetworkMLComputeUpdateEngine{MLNeuralNetworkV1Engine: MLNeuralNetworkV1EngineFromID(id)}
}
// Ensure MLNeuralNetworkMLComputeUpdateEngine implements IMLNeuralNetworkMLComputeUpdateEngine.
var _ IMLNeuralNetworkMLComputeUpdateEngine = MLNeuralNetworkMLComputeUpdateEngine{}

// An interface definition for the [MLNeuralNetworkMLComputeUpdateEngine] class.
//
// # Methods
//
//   - [IMLNeuralNetworkMLComputeUpdateEngine.BatchSize]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetBatchSize]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.CancelUpdate]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ClassLabelToIndexMap]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ClassifierOutputIsSigmoidOutput]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetClassifierOutputIsSigmoidOutput]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ContinueWithUpdate]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetContinueWithUpdate]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.FinalLossValue]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetFinalLossValue]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.LoadLossTargetName]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.LossTargetName]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.MlcGraph]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetMlcGraph]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ParameterContainer]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetParameterContainer]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ParameterValueForKey]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.PerformInferenceWithOutputNameToLayerMapError]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.PerformTrainingWithCallBacksNumberOfEpochsError]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ProgressHandlers]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetProgressHandlers]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ProgressHandlersDispatchQueue]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetProgressHandlersDispatchQueue]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ResumeUpdate]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ResumeUpdateWithParameters]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.ShuffableTrainingData]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.SetShuffableTrainingData]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.UpdateLearningRateWithValue]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.UpdateModelWithData]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.UpdateParameters]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.WriteToURLError]
//   - [IMLNeuralNetworkMLComputeUpdateEngine.InitWithCompiledArchiveNnContainerConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine
type IMLNeuralNetworkMLComputeUpdateEngine interface {
	IMLNeuralNetworkV1Engine

	// Topic: Methods

	BatchSize() uint64
	SetBatchSize(value uint64)
	CancelUpdate()
	ClassLabelToIndexMap() foundation.INSDictionary
	ClassifierOutputIsSigmoidOutput() bool
	SetClassifierOutputIsSigmoidOutput(value bool)
	ContinueWithUpdate() bool
	SetContinueWithUpdate(value bool)
	FinalLossValue() float32
	SetFinalLossValue(value float32)
	LoadLossTargetName(name unsafe.Pointer) objectivec.IObject
	LossTargetName() string
	MlcGraph() IMLNeuralNetworkMLComputeGraph
	SetMlcGraph(value IMLNeuralNetworkMLComputeGraph)
	ParameterContainer() IMLParameterContainer
	SetParameterContainer(value IMLParameterContainer)
	ParameterValueForKey(key objectivec.IObject) objectivec.IObject
	PerformInferenceWithOutputNameToLayerMapError(with objectivec.IObject, map_ objectivec.IObject) (objectivec.IObject, error)
	PerformTrainingWithCallBacksNumberOfEpochsError(with objectivec.IObject, backs objectivec.IObject, epochs uint64) (bool, error)
	ProgressHandlers() IMLUpdateProgressHandlers
	SetProgressHandlers(value IMLUpdateProgressHandlers)
	ProgressHandlersDispatchQueue() objectivec.Object
	SetProgressHandlersDispatchQueue(value objectivec.Object)
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject)
	ShuffableTrainingData() IMLShufflingBatchProvider
	SetShuffableTrainingData(value IMLShufflingBatchProvider)
	UpdateLearningRateWithValue(value float32)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() objectivec.IObject
	WriteToURLError(url foundation.INSURL) (bool, error)
	InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error)
}

// Init initializes the instance.
func (n MLNeuralNetworkMLComputeUpdateEngine) Init() MLNeuralNetworkMLComputeUpdateEngine {
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkMLComputeUpdateEngine) Autorelease() MLNeuralNetworkMLComputeUpdateEngine {
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkMLComputeUpdateEngine creates a new MLNeuralNetworkMLComputeUpdateEngine instance.
func NewMLNeuralNetworkMLComputeUpdateEngine() MLNeuralNetworkMLComputeUpdateEngine {
	class := getMLNeuralNetworkMLComputeUpdateEngineClass()
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/initWithCompiledArchive:nnContainer:configuration:error:
func NewNeuralNetworkMLComputeUpdateEngineWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompiledArchive:nnContainer:configuration:error:"), archive, container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:configuration:error:
func NewNeuralNetworkMLComputeUpdateEngineWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:error:
func NewNeuralNetworkMLComputeUpdateEngineWithContainerError(container objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewNeuralNetworkMLComputeUpdateEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkMLComputeUpdateEngine {
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewNeuralNetworkMLComputeUpdateEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkMLComputeUpdateEngine {
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/cancelUpdate
func (n MLNeuralNetworkMLComputeUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](n.ID, objc.Sel("cancelUpdate"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/loadLossTargetName:
func (n MLNeuralNetworkMLComputeUpdateEngine) LoadLossTargetName(name unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("loadLossTargetName:"), name)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/parameterValueForKey:
func (n MLNeuralNetworkMLComputeUpdateEngine) ParameterValueForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameterValueForKey:"), key)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/performInferenceWith:outputNameToLayerMap:error:
func (n MLNeuralNetworkMLComputeUpdateEngine) PerformInferenceWithOutputNameToLayerMapError(with objectivec.IObject, map_ objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("performInferenceWith:outputNameToLayerMap:error:"), with, map_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/performTrainingWith:callBacks:numberOfEpochs:error:
func (n MLNeuralNetworkMLComputeUpdateEngine) PerformTrainingWithCallBacksNumberOfEpochsError(with objectivec.IObject, backs objectivec.IObject, epochs uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("performTrainingWith:callBacks:numberOfEpochs:error:"), with, backs, epochs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performTrainingWith:callBacks:numberOfEpochs:error: returned NO with nil NSError")
	}
	return rv, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/resumeUpdate
func (n MLNeuralNetworkMLComputeUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](n.ID, objc.Sel("resumeUpdate"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/resumeUpdateWithParameters:
func (n MLNeuralNetworkMLComputeUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/setUpdateProgressHandlers:dispatchQueue:
func (n MLNeuralNetworkMLComputeUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](n.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/updateLearningRateWithValue:
func (n MLNeuralNetworkMLComputeUpdateEngine) UpdateLearningRateWithValue(value float32) {
	objc.Send[objc.ID](n.ID, objc.Sel("updateLearningRateWithValue:"), value)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/updateModelWithData:
func (n MLNeuralNetworkMLComputeUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("updateModelWithData:"), data)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/updateParameters
func (n MLNeuralNetworkMLComputeUpdateEngine) UpdateParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("updateParameters"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/writeToURL:error:
func (n MLNeuralNetworkMLComputeUpdateEngine) WriteToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/initWithCompiledArchive:nnContainer:configuration:error:
func (n MLNeuralNetworkMLComputeUpdateEngine) InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("initWithCompiledArchive:nnContainer:configuration:error:"), archive, container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/batchSize
func (n MLNeuralNetworkMLComputeUpdateEngine) BatchSize() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("batchSize"))
	return rv
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetBatchSize(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setBatchSize:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/classLabelToIndexMap
func (n MLNeuralNetworkMLComputeUpdateEngine) ClassLabelToIndexMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("classLabelToIndexMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/classifierOutputIsSigmoidOutput
func (n MLNeuralNetworkMLComputeUpdateEngine) ClassifierOutputIsSigmoidOutput() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("classifierOutputIsSigmoidOutput"))
	return rv
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetClassifierOutputIsSigmoidOutput(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setClassifierOutputIsSigmoidOutput:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/continueWithUpdate
func (n MLNeuralNetworkMLComputeUpdateEngine) ContinueWithUpdate() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetContinueWithUpdate(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setContinueWithUpdate:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/finalLossValue
func (n MLNeuralNetworkMLComputeUpdateEngine) FinalLossValue() float32 {
	rv := objc.Send[float32](n.ID, objc.Sel("finalLossValue"))
	return rv
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetFinalLossValue(value float32) {
	objc.Send[struct{}](n.ID, objc.Sel("setFinalLossValue:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/lossTargetName
func (n MLNeuralNetworkMLComputeUpdateEngine) LossTargetName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("lossTargetName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/mlcGraph
func (n MLNeuralNetworkMLComputeUpdateEngine) MlcGraph() IMLNeuralNetworkMLComputeGraph {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("mlcGraph"))
	return MLNeuralNetworkMLComputeGraphFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetMlcGraph(value IMLNeuralNetworkMLComputeGraph) {
	objc.Send[struct{}](n.ID, objc.Sel("setMlcGraph:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/parameterContainer
func (n MLNeuralNetworkMLComputeUpdateEngine) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](n.ID, objc.Sel("setParameterContainer:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/progressHandlers
func (n MLNeuralNetworkMLComputeUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](n.ID, objc.Sel("setProgressHandlers:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/progressHandlersDispatchQueue
func (n MLNeuralNetworkMLComputeUpdateEngine) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](n.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeUpdateEngine/shuffableTrainingData
func (n MLNeuralNetworkMLComputeUpdateEngine) ShuffableTrainingData() IMLShufflingBatchProvider {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("shuffableTrainingData"))
	return MLShufflingBatchProviderFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeUpdateEngine) SetShuffableTrainingData(value IMLShufflingBatchProvider) {
	objc.Send[struct{}](n.ID, objc.Sel("setShuffableTrainingData:"), value)
}

