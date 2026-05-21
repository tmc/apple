// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject)
	ShuffableTrainingData() IMLShufflingBatchProvider
	SetShuffableTrainingData(value IMLShufflingBatchProvider)
	UpdateLearningRateWithValue(value float32)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() objectivec.IObject
	WriteToURLError(url foundation.NSURL) (bool, error)
	InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error)
}

// Init initializes the instance.
func (m MLNeuralNetworkMLComputeUpdateEngine) Init() MLNeuralNetworkMLComputeUpdateEngine {
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworkMLComputeUpdateEngine) Autorelease() MLNeuralNetworkMLComputeUpdateEngine {
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkMLComputeUpdateEngine creates a new MLNeuralNetworkMLComputeUpdateEngine instance.
func NewMLNeuralNetworkMLComputeUpdateEngine() MLNeuralNetworkMLComputeUpdateEngine {
	class := getMLNeuralNetworkMLComputeUpdateEngineClass()
	rv := objc.Send[MLNeuralNetworkMLComputeUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

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

func NewNeuralNetworkMLComputeUpdateEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkMLComputeUpdateEngine {
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv)
}

func NewNeuralNetworkMLComputeUpdateEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkMLComputeUpdateEngine {
	instance := getMLNeuralNetworkMLComputeUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv)
}

func (m MLNeuralNetworkMLComputeUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) LoadLossTargetName(name unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("loadLossTargetName:"), name)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ParameterValueForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:"), key)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkMLComputeUpdateEngine) PerformInferenceWithOutputNameToLayerMapError(with objectivec.IObject, map_ objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("performInferenceWith:outputNameToLayerMap:error:"), with, map_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkMLComputeUpdateEngine) PerformTrainingWithCallBacksNumberOfEpochsError(with objectivec.IObject, backs objectivec.IObject, epochs uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("performTrainingWith:callBacks:numberOfEpochs:error:"), with, backs, epochs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("performTrainingWith:callBacks:numberOfEpochs:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkMLComputeUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) UpdateLearningRateWithValue(value float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateLearningRateWithValue:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) UpdateParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("updateParameters"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkMLComputeUpdateEngine) WriteToURLError(url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkMLComputeUpdateEngine) InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkMLComputeUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithCompiledArchive:nnContainer:configuration:error:"), archive, container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeUpdateEngineFromID(rv), nil

}

func (m MLNeuralNetworkMLComputeUpdateEngine) BatchSize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetBatchSize(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ClassLabelToIndexMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classLabelToIndexMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ClassifierOutputIsSigmoidOutput() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("classifierOutputIsSigmoidOutput"))
	return rv
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetClassifierOutputIsSigmoidOutput(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassifierOutputIsSigmoidOutput:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ContinueWithUpdate() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetContinueWithUpdate(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setContinueWithUpdate:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) FinalLossValue() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("finalLossValue"))
	return rv
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetFinalLossValue(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setFinalLossValue:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) LossTargetName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lossTargetName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkMLComputeUpdateEngine) MlcGraph() IMLNeuralNetworkMLComputeGraph {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mlcGraph"))
	return MLNeuralNetworkMLComputeGraphFromID(objc.ID(rv))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetMlcGraph(value IMLNeuralNetworkMLComputeGraph) {
	objc.Send[struct{}](m.ID, objc.Sel("setMlcGraph:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameterContainer:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlers:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}
func (m MLNeuralNetworkMLComputeUpdateEngine) ShuffableTrainingData() IMLShufflingBatchProvider {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shuffableTrainingData"))
	return MLShufflingBatchProviderFromID(objc.ID(rv))
}
func (m MLNeuralNetworkMLComputeUpdateEngine) SetShuffableTrainingData(value IMLShufflingBatchProvider) {
	objc.Send[struct{}](m.ID, objc.Sel("setShuffableTrainingData:"), value)
}
