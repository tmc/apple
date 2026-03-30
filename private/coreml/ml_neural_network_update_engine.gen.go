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

// The class instance for the [MLNeuralNetworkUpdateEngine] class.
var (
	_MLNeuralNetworkUpdateEngineClass     MLNeuralNetworkUpdateEngineClass
	_MLNeuralNetworkUpdateEngineClassOnce sync.Once
)

func getMLNeuralNetworkUpdateEngineClass() MLNeuralNetworkUpdateEngineClass {
	_MLNeuralNetworkUpdateEngineClassOnce.Do(func() {
		_MLNeuralNetworkUpdateEngineClass = MLNeuralNetworkUpdateEngineClass{class: objc.GetClass("MLNeuralNetworkUpdateEngine")}
	})
	return _MLNeuralNetworkUpdateEngineClass
}

// GetMLNeuralNetworkUpdateEngineClass returns the class object for MLNeuralNetworkUpdateEngine.
func GetMLNeuralNetworkUpdateEngineClass() MLNeuralNetworkUpdateEngineClass {
	return getMLNeuralNetworkUpdateEngineClass()
}

type MLNeuralNetworkUpdateEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkUpdateEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkUpdateEngineClass) Alloc() MLNeuralNetworkUpdateEngine {
	rv := objc.Send[MLNeuralNetworkUpdateEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralNetworkUpdateEngine.BiasForLayerError]
//   - [MLNeuralNetworkUpdateEngine.CancelUpdate]
//   - [MLNeuralNetworkUpdateEngine.ClassLabelToIndexMap]
//   - [MLNeuralNetworkUpdateEngine.CollectMetricsFromTaskContextIsInCallBack]
//   - [MLNeuralNetworkUpdateEngine.ContinueWithUpdate]
//   - [MLNeuralNetworkUpdateEngine.SetContinueWithUpdate]
//   - [MLNeuralNetworkUpdateEngine.CoreMLToEspressoParamsMap]
//   - [MLNeuralNetworkUpdateEngine.SetCoreMLToEspressoParamsMap]
//   - [MLNeuralNetworkUpdateEngine.CreateEspressoTaskFromUpdateParametersLossInputNameLossTargetNameLossOutputNameUpdatableLayerNamesConfigurationError]
//   - [MLNeuralNetworkUpdateEngine.LoadLossInputNameUpdatableLayerNamesFromCompiledArchive]
//   - [MLNeuralNetworkUpdateEngine.LoadLossTargetNameLossOutputNameFromUpdateParameters]
//   - [MLNeuralNetworkUpdateEngine.LossOutputName]
//   - [MLNeuralNetworkUpdateEngine.SetLossOutputName]
//   - [MLNeuralNetworkUpdateEngine.LossTargetName]
//   - [MLNeuralNetworkUpdateEngine.SetLossTargetName]
//   - [MLNeuralNetworkUpdateEngine.ParameterContainer]
//   - [MLNeuralNetworkUpdateEngine.SetParameterContainer]
//   - [MLNeuralNetworkUpdateEngine.ParameterValueForKey]
//   - [MLNeuralNetworkUpdateEngine.ParamsForLayerParameterTypeError]
//   - [MLNeuralNetworkUpdateEngine.ProgressHandlers]
//   - [MLNeuralNetworkUpdateEngine.SetProgressHandlers]
//   - [MLNeuralNetworkUpdateEngine.ProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkUpdateEngine.SetProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkUpdateEngine.ResumeUpdate]
//   - [MLNeuralNetworkUpdateEngine.ResumeUpdateWithParameters]
//   - [MLNeuralNetworkUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLNeuralNetworkUpdateEngine.SetWeightsOrBiasesForLayerLayerTypeValueError]
//   - [MLNeuralNetworkUpdateEngine.ShuffableTrainingData]
//   - [MLNeuralNetworkUpdateEngine.SetShuffableTrainingData]
//   - [MLNeuralNetworkUpdateEngine.Snapshot]
//   - [MLNeuralNetworkUpdateEngine.SetSnapshot]
//   - [MLNeuralNetworkUpdateEngine.StringForDataType]
//   - [MLNeuralNetworkUpdateEngine.Task]
//   - [MLNeuralNetworkUpdateEngine.SetTask]
//   - [MLNeuralNetworkUpdateEngine.UpdateLearningRateWithTaskContextIsInCallBackError]
//   - [MLNeuralNetworkUpdateEngine.UpdateModelWithData]
//   - [MLNeuralNetworkUpdateEngine.UpdateParameters]
//   - [MLNeuralNetworkUpdateEngine.UpdateWeightsAndBiasesFromConfigParamsError]
//   - [MLNeuralNetworkUpdateEngine.WeightsForLayerError]
//   - [MLNeuralNetworkUpdateEngine.WriteToURLError]
//   - [MLNeuralNetworkUpdateEngine.InitWithCompiledArchiveNnContainerConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine
type MLNeuralNetworkUpdateEngine struct {
	MLNeuralNetworkV1Engine
}

// MLNeuralNetworkUpdateEngineFromID constructs a [MLNeuralNetworkUpdateEngine] from an objc.ID.
func MLNeuralNetworkUpdateEngineFromID(id objc.ID) MLNeuralNetworkUpdateEngine {
	return MLNeuralNetworkUpdateEngine{MLNeuralNetworkV1Engine: MLNeuralNetworkV1EngineFromID(id)}
}

// Ensure MLNeuralNetworkUpdateEngine implements IMLNeuralNetworkUpdateEngine.
var _ IMLNeuralNetworkUpdateEngine = MLNeuralNetworkUpdateEngine{}

// An interface definition for the [MLNeuralNetworkUpdateEngine] class.
//
// # Methods
//
//   - [IMLNeuralNetworkUpdateEngine.BiasForLayerError]
//   - [IMLNeuralNetworkUpdateEngine.CancelUpdate]
//   - [IMLNeuralNetworkUpdateEngine.ClassLabelToIndexMap]
//   - [IMLNeuralNetworkUpdateEngine.CollectMetricsFromTaskContextIsInCallBack]
//   - [IMLNeuralNetworkUpdateEngine.ContinueWithUpdate]
//   - [IMLNeuralNetworkUpdateEngine.SetContinueWithUpdate]
//   - [IMLNeuralNetworkUpdateEngine.CoreMLToEspressoParamsMap]
//   - [IMLNeuralNetworkUpdateEngine.SetCoreMLToEspressoParamsMap]
//   - [IMLNeuralNetworkUpdateEngine.CreateEspressoTaskFromUpdateParametersLossInputNameLossTargetNameLossOutputNameUpdatableLayerNamesConfigurationError]
//   - [IMLNeuralNetworkUpdateEngine.LoadLossInputNameUpdatableLayerNamesFromCompiledArchive]
//   - [IMLNeuralNetworkUpdateEngine.LoadLossTargetNameLossOutputNameFromUpdateParameters]
//   - [IMLNeuralNetworkUpdateEngine.LossOutputName]
//   - [IMLNeuralNetworkUpdateEngine.SetLossOutputName]
//   - [IMLNeuralNetworkUpdateEngine.LossTargetName]
//   - [IMLNeuralNetworkUpdateEngine.SetLossTargetName]
//   - [IMLNeuralNetworkUpdateEngine.ParameterContainer]
//   - [IMLNeuralNetworkUpdateEngine.SetParameterContainer]
//   - [IMLNeuralNetworkUpdateEngine.ParameterValueForKey]
//   - [IMLNeuralNetworkUpdateEngine.ParamsForLayerParameterTypeError]
//   - [IMLNeuralNetworkUpdateEngine.ProgressHandlers]
//   - [IMLNeuralNetworkUpdateEngine.SetProgressHandlers]
//   - [IMLNeuralNetworkUpdateEngine.ProgressHandlersDispatchQueue]
//   - [IMLNeuralNetworkUpdateEngine.SetProgressHandlersDispatchQueue]
//   - [IMLNeuralNetworkUpdateEngine.ResumeUpdate]
//   - [IMLNeuralNetworkUpdateEngine.ResumeUpdateWithParameters]
//   - [IMLNeuralNetworkUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [IMLNeuralNetworkUpdateEngine.SetWeightsOrBiasesForLayerLayerTypeValueError]
//   - [IMLNeuralNetworkUpdateEngine.ShuffableTrainingData]
//   - [IMLNeuralNetworkUpdateEngine.SetShuffableTrainingData]
//   - [IMLNeuralNetworkUpdateEngine.Snapshot]
//   - [IMLNeuralNetworkUpdateEngine.SetSnapshot]
//   - [IMLNeuralNetworkUpdateEngine.StringForDataType]
//   - [IMLNeuralNetworkUpdateEngine.Task]
//   - [IMLNeuralNetworkUpdateEngine.SetTask]
//   - [IMLNeuralNetworkUpdateEngine.UpdateLearningRateWithTaskContextIsInCallBackError]
//   - [IMLNeuralNetworkUpdateEngine.UpdateModelWithData]
//   - [IMLNeuralNetworkUpdateEngine.UpdateParameters]
//   - [IMLNeuralNetworkUpdateEngine.UpdateWeightsAndBiasesFromConfigParamsError]
//   - [IMLNeuralNetworkUpdateEngine.WeightsForLayerError]
//   - [IMLNeuralNetworkUpdateEngine.WriteToURLError]
//   - [IMLNeuralNetworkUpdateEngine.InitWithCompiledArchiveNnContainerConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine
type IMLNeuralNetworkUpdateEngine interface {
	IMLNeuralNetworkV1Engine

	// Topic: Methods

	BiasForLayerError(layer objectivec.IObject) (objectivec.IObject, error)
	CancelUpdate()
	ClassLabelToIndexMap() foundation.INSDictionary
	CollectMetricsFromTaskContextIsInCallBack(context objectivec.IObject, back bool) objectivec.IObject
	ContinueWithUpdate() bool
	SetContinueWithUpdate(value bool)
	CoreMLToEspressoParamsMap() foundation.INSDictionary
	SetCoreMLToEspressoParamsMap(value foundation.INSDictionary)
	CreateEspressoTaskFromUpdateParametersLossInputNameLossTargetNameLossOutputNameUpdatableLayerNamesConfigurationError(from objectivec.IObject, parameters unsafe.Pointer, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, names objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error)
	LoadLossInputNameUpdatableLayerNamesFromCompiledArchive(name []objectivec.IObject, names []objectivec.IObject, archive unsafe.Pointer)
	LoadLossTargetNameLossOutputNameFromUpdateParameters(name []objectivec.IObject, name2 []objectivec.IObject, parameters unsafe.Pointer)
	LossOutputName() string
	SetLossOutputName(value string)
	LossTargetName() string
	SetLossTargetName(value string)
	ParameterContainer() IMLParameterContainer
	SetParameterContainer(value IMLParameterContainer)
	ParameterValueForKey(key objectivec.IObject) objectivec.IObject
	ParamsForLayerParameterTypeError(layer objectivec.IObject, type_ uint64) (objectivec.IObject, error)
	ProgressHandlers() IMLUpdateProgressHandlers
	SetProgressHandlers(value IMLUpdateProgressHandlers)
	ProgressHandlersDispatchQueue() objectivec.Object
	SetProgressHandlersDispatchQueue(value objectivec.Object)
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject)
	SetWeightsOrBiasesForLayerLayerTypeValueError(layer objectivec.IObject, type_ uint64, value objectivec.IObject) (bool, error)
	ShuffableTrainingData() IMLShufflingBatchProvider
	SetShuffableTrainingData(value IMLShufflingBatchProvider)
	Snapshot() objectivec.IObject
	SetSnapshot(value objectivec.IObject)
	StringForDataType(type_ uint64) objectivec.IObject
	Task() objectivec.IObject
	SetTask(value objectivec.IObject)
	UpdateLearningRateWithTaskContextIsInCallBackError(context objectivec.IObject, back bool) (bool, error)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() objectivec.IObject
	UpdateWeightsAndBiasesFromConfigParamsError(params objectivec.IObject) (bool, error)
	WeightsForLayerError(layer objectivec.IObject) (objectivec.IObject, error)
	WriteToURLError(url foundation.INSURL) (bool, error)
	InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkUpdateEngine, error)
}

// Init initializes the instance.
func (n MLNeuralNetworkUpdateEngine) Init() MLNeuralNetworkUpdateEngine {
	rv := objc.Send[MLNeuralNetworkUpdateEngine](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkUpdateEngine) Autorelease() MLNeuralNetworkUpdateEngine {
	rv := objc.Send[MLNeuralNetworkUpdateEngine](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkUpdateEngine creates a new MLNeuralNetworkUpdateEngine instance.
func NewMLNeuralNetworkUpdateEngine() MLNeuralNetworkUpdateEngine {
	class := getMLNeuralNetworkUpdateEngineClass()
	rv := objc.Send[MLNeuralNetworkUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/initWithCoder:
func NewNeuralNetworkUpdateEngineWithCoder(coder objectivec.IObject) MLNeuralNetworkUpdateEngine {
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNeuralNetworkUpdateEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/initWithCompiledArchive:nnContainer:configuration:error:
func NewNeuralNetworkUpdateEngineWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompiledArchive:nnContainer:configuration:error:"), archive, container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkUpdateEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:configuration:error:
func NewNeuralNetworkUpdateEngineWithContainerConfigurationError(container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:configuration:error:"), container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkUpdateEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkEngine/initWithContainer:error:
func NewNeuralNetworkUpdateEngineWithContainerError(container objectivec.IObject) (MLNeuralNetworkUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContainer:error:"), container, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkUpdateEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewNeuralNetworkUpdateEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkUpdateEngine {
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkUpdateEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewNeuralNetworkUpdateEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkUpdateEngine {
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkUpdateEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/biasForLayer:error:
func (n MLNeuralNetworkUpdateEngine) BiasForLayerError(layer objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("biasForLayer:error:"), layer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/cancelUpdate
func (n MLNeuralNetworkUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](n.ID, objc.Sel("cancelUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/collectMetricsFromTaskContext:isInCallBack:
func (n MLNeuralNetworkUpdateEngine) CollectMetricsFromTaskContextIsInCallBack(context objectivec.IObject, back bool) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("collectMetricsFromTaskContext:isInCallBack:"), context, back)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/createEspressoTaskFrom:updateParameters:lossInputName:lossTargetName:lossOutputName:updatableLayerNames:configuration:error:
func (n MLNeuralNetworkUpdateEngine) CreateEspressoTaskFromUpdateParametersLossInputNameLossTargetNameLossOutputNameUpdatableLayerNamesConfigurationError(from objectivec.IObject, parameters unsafe.Pointer, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, names objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createEspressoTaskFrom:updateParameters:lossInputName:lossTargetName:lossOutputName:updatableLayerNames:configuration:error:"), from, parameters, name, name2, name3, names, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/loadLossInputName:updatableLayerNames:fromCompiledArchive:
func (n MLNeuralNetworkUpdateEngine) LoadLossInputNameUpdatableLayerNamesFromCompiledArchive(name []objectivec.IObject, names []objectivec.IObject, archive unsafe.Pointer) {
	objc.Send[objc.ID](n.ID, objc.Sel("loadLossInputName:updatableLayerNames:fromCompiledArchive:"), objectivec.IObjectSliceToNSArray(name), objectivec.IObjectSliceToNSArray(names), archive)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/loadLossTargetName:lossOutputName:fromUpdateParameters:
func (n MLNeuralNetworkUpdateEngine) LoadLossTargetNameLossOutputNameFromUpdateParameters(name []objectivec.IObject, name2 []objectivec.IObject, parameters unsafe.Pointer) {
	objc.Send[objc.ID](n.ID, objc.Sel("loadLossTargetName:lossOutputName:fromUpdateParameters:"), objectivec.IObjectSliceToNSArray(name), objectivec.IObjectSliceToNSArray(name2), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/parameterValueForKey:
func (n MLNeuralNetworkUpdateEngine) ParameterValueForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameterValueForKey:"), key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/paramsForLayer:parameterType:error:
func (n MLNeuralNetworkUpdateEngine) ParamsForLayerParameterTypeError(layer objectivec.IObject, type_ uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("paramsForLayer:parameterType:error:"), layer, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/resumeUpdate
func (n MLNeuralNetworkUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](n.ID, objc.Sel("resumeUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/resumeUpdateWithParameters:
func (n MLNeuralNetworkUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/setUpdateProgressHandlers:dispatchQueue:
func (n MLNeuralNetworkUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
	_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](n.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/setWeightsOrBiasesForLayer:layerType:value:error:
func (n MLNeuralNetworkUpdateEngine) SetWeightsOrBiasesForLayerLayerTypeValueError(layer objectivec.IObject, type_ uint64, value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("setWeightsOrBiasesForLayer:layerType:value:error:"), layer, type_, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWeightsOrBiasesForLayer:layerType:value:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/stringForDataType:
func (n MLNeuralNetworkUpdateEngine) StringForDataType(type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("stringForDataType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/updateLearningRateWithTaskContext:isInCallBack:error:
func (n MLNeuralNetworkUpdateEngine) UpdateLearningRateWithTaskContextIsInCallBackError(context objectivec.IObject, back bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("updateLearningRateWithTaskContext:isInCallBack:error:"), context, back, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateLearningRateWithTaskContext:isInCallBack:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/updateModelWithData:
func (n MLNeuralNetworkUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](n.ID, objc.Sel("updateModelWithData:"), data)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/updateParameters
func (n MLNeuralNetworkUpdateEngine) UpdateParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("updateParameters"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/updateWeightsAndBiasesFromConfigParams:error:
func (n MLNeuralNetworkUpdateEngine) UpdateWeightsAndBiasesFromConfigParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("updateWeightsAndBiasesFromConfigParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateWeightsAndBiasesFromConfigParams:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/weightsForLayer:error:
func (n MLNeuralNetworkUpdateEngine) WeightsForLayerError(layer objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("weightsForLayer:error:"), layer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/writeToURL:error:
func (n MLNeuralNetworkUpdateEngine) WriteToURLError(url foundation.INSURL) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/initWithCompiledArchive:nnContainer:configuration:error:
func (n MLNeuralNetworkUpdateEngine) InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("initWithCompiledArchive:nnContainer:configuration:error:"), archive, container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkUpdateEngineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/createCoreMLToEspressoParamsMap
func (_MLNeuralNetworkUpdateEngineClass MLNeuralNetworkUpdateEngineClass) CreateCoreMLToEspressoParamsMap() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkUpdateEngineClass.class), objc.Sel("createCoreMLToEspressoParamsMap"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/supportsSecureCoding
func (_MLNeuralNetworkUpdateEngineClass MLNeuralNetworkUpdateEngineClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNeuralNetworkUpdateEngineClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/classLabelToIndexMap
func (n MLNeuralNetworkUpdateEngine) ClassLabelToIndexMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("classLabelToIndexMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/continueWithUpdate
func (n MLNeuralNetworkUpdateEngine) ContinueWithUpdate() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (n MLNeuralNetworkUpdateEngine) SetContinueWithUpdate(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setContinueWithUpdate:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/coreMLToEspressoParamsMap
func (n MLNeuralNetworkUpdateEngine) CoreMLToEspressoParamsMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("coreMLToEspressoParamsMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkUpdateEngine) SetCoreMLToEspressoParamsMap(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setCoreMLToEspressoParamsMap:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/lossOutputName
func (n MLNeuralNetworkUpdateEngine) LossOutputName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("lossOutputName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNeuralNetworkUpdateEngine) SetLossOutputName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setLossOutputName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/lossTargetName
func (n MLNeuralNetworkUpdateEngine) LossTargetName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("lossTargetName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNeuralNetworkUpdateEngine) SetLossTargetName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setLossTargetName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/parameterContainer
func (n MLNeuralNetworkUpdateEngine) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (n MLNeuralNetworkUpdateEngine) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](n.ID, objc.Sel("setParameterContainer:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/progressHandlers
func (n MLNeuralNetworkUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (n MLNeuralNetworkUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](n.ID, objc.Sel("setProgressHandlers:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/progressHandlersDispatchQueue
func (n MLNeuralNetworkUpdateEngine) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (n MLNeuralNetworkUpdateEngine) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](n.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/shuffableTrainingData
func (n MLNeuralNetworkUpdateEngine) ShuffableTrainingData() IMLShufflingBatchProvider {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("shuffableTrainingData"))
	return MLShufflingBatchProviderFromID(objc.ID(rv))
}
func (n MLNeuralNetworkUpdateEngine) SetShuffableTrainingData(value IMLShufflingBatchProvider) {
	objc.Send[struct{}](n.ID, objc.Sel("setShuffableTrainingData:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/snapshot
func (n MLNeuralNetworkUpdateEngine) Snapshot() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("snapshot"))
	return objectivec.Object{ID: rv}
}
func (n MLNeuralNetworkUpdateEngine) SetSnapshot(value objectivec.IObject) {
	objc.Send[struct{}](n.ID, objc.Sel("setSnapshot:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkUpdateEngine/task
func (n MLNeuralNetworkUpdateEngine) Task() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("task"))
	return objectivec.Object{ID: rv}
}
func (n MLNeuralNetworkUpdateEngine) SetTask(value objectivec.IObject) {
	objc.Send[struct{}](n.ID, objc.Sel("setTask:"), value)
}
