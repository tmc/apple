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
	SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject)
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
	WriteToURLError(url foundation.NSURL) (bool, error)
	InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkUpdateEngine, error)
}

// Init initializes the instance.
func (m MLNeuralNetworkUpdateEngine) Init() MLNeuralNetworkUpdateEngine {
	rv := objc.Send[MLNeuralNetworkUpdateEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNeuralNetworkUpdateEngine) Autorelease() MLNeuralNetworkUpdateEngine {
	rv := objc.Send[MLNeuralNetworkUpdateEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkUpdateEngine creates a new MLNeuralNetworkUpdateEngine instance.
func NewMLNeuralNetworkUpdateEngine() MLNeuralNetworkUpdateEngine {
	class := getMLNeuralNetworkUpdateEngineClass()
	rv := objc.Send[MLNeuralNetworkUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewNeuralNetworkUpdateEngineWithCoder(coder objectivec.IObject) MLNeuralNetworkUpdateEngine {
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLNeuralNetworkUpdateEngineFromID(rv)
}

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

func NewNeuralNetworkUpdateEngineWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkUpdateEngine {
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNeuralNetworkUpdateEngineFromID(rv)
}

func NewNeuralNetworkUpdateEngineWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNeuralNetworkUpdateEngine {
	instance := getMLNeuralNetworkUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNeuralNetworkUpdateEngineFromID(rv)
}

func (m MLNeuralNetworkUpdateEngine) BiasForLayerError(layer objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("biasForLayer:error:"), layer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}
func (m MLNeuralNetworkUpdateEngine) CollectMetricsFromTaskContextIsInCallBack(context objectivec.IObject, back bool) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("collectMetricsFromTaskContext:isInCallBack:"), context, back)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkUpdateEngine) CreateEspressoTaskFromUpdateParametersLossInputNameLossTargetNameLossOutputNameUpdatableLayerNamesConfigurationError(from objectivec.IObject, parameters unsafe.Pointer, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, names objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("createEspressoTaskFrom:updateParameters:lossInputName:lossTargetName:lossOutputName:updatableLayerNames:configuration:error:"), from, parameters, name, name2, name3, names, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkUpdateEngine) LoadLossInputNameUpdatableLayerNamesFromCompiledArchive(name []objectivec.IObject, names []objectivec.IObject, archive unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("loadLossInputName:updatableLayerNames:fromCompiledArchive:"), objectivec.IObjectSliceToNSArray(name), objectivec.IObjectSliceToNSArray(names), archive)
}
func (m MLNeuralNetworkUpdateEngine) LoadLossTargetNameLossOutputNameFromUpdateParameters(name []objectivec.IObject, name2 []objectivec.IObject, parameters unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("loadLossTargetName:lossOutputName:fromUpdateParameters:"), objectivec.IObjectSliceToNSArray(name), objectivec.IObjectSliceToNSArray(name2), parameters)
}
func (m MLNeuralNetworkUpdateEngine) ParameterValueForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:"), key)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkUpdateEngine) ParamsForLayerParameterTypeError(layer objectivec.IObject, type_ uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("paramsForLayer:parameterType:error:"), layer, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}
func (m MLNeuralNetworkUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
func (m MLNeuralNetworkUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}
func (m MLNeuralNetworkUpdateEngine) SetWeightsOrBiasesForLayerLayerTypeValueError(layer objectivec.IObject, type_ uint64, value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setWeightsOrBiasesForLayer:layerType:value:error:"), layer, type_, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWeightsOrBiasesForLayer:layerType:value:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkUpdateEngine) StringForDataType(type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringForDataType:"), type_)
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkUpdateEngine) UpdateLearningRateWithTaskContextIsInCallBackError(context objectivec.IObject, back bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updateLearningRateWithTaskContext:isInCallBack:error:"), context, back, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateLearningRateWithTaskContext:isInCallBack:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}
func (m MLNeuralNetworkUpdateEngine) UpdateParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("updateParameters"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkUpdateEngine) UpdateWeightsAndBiasesFromConfigParamsError(params objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updateWeightsAndBiasesFromConfigParams:error:"), params, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateWeightsAndBiasesFromConfigParams:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLNeuralNetworkUpdateEngine) WeightsForLayerError(layer objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("weightsForLayer:error:"), layer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLNeuralNetworkUpdateEngine) WriteToURLError(url foundation.NSURL) (bool, error) {
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
func (m MLNeuralNetworkUpdateEngine) InitWithCompiledArchiveNnContainerConfigurationError(archive unsafe.Pointer, container objectivec.IObject, configuration objectivec.IObject) (MLNeuralNetworkUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithCompiledArchive:nnContainer:configuration:error:"), archive, container, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkUpdateEngineFromID(rv), nil

}

func (_MLNeuralNetworkUpdateEngineClass MLNeuralNetworkUpdateEngineClass) CreateCoreMLToEspressoParamsMap() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkUpdateEngineClass.class), objc.Sel("createCoreMLToEspressoParamsMap"))
	return objectivec.Object{ID: rv}
}
func (_MLNeuralNetworkUpdateEngineClass MLNeuralNetworkUpdateEngineClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLNeuralNetworkUpdateEngineClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLNeuralNetworkUpdateEngine) ClassLabelToIndexMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classLabelToIndexMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkUpdateEngine) ContinueWithUpdate() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (m MLNeuralNetworkUpdateEngine) SetContinueWithUpdate(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setContinueWithUpdate:"), value)
}
func (m MLNeuralNetworkUpdateEngine) CoreMLToEspressoParamsMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("coreMLToEspressoParamsMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLNeuralNetworkUpdateEngine) SetCoreMLToEspressoParamsMap(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setCoreMLToEspressoParamsMap:"), value)
}
func (m MLNeuralNetworkUpdateEngine) LossOutputName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lossOutputName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkUpdateEngine) SetLossOutputName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLossOutputName:"), objc.String(value))
}
func (m MLNeuralNetworkUpdateEngine) LossTargetName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lossTargetName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNeuralNetworkUpdateEngine) SetLossTargetName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLossTargetName:"), objc.String(value))
}
func (m MLNeuralNetworkUpdateEngine) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (m MLNeuralNetworkUpdateEngine) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameterContainer:"), value)
}
func (m MLNeuralNetworkUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLNeuralNetworkUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlers:"), value)
}
func (m MLNeuralNetworkUpdateEngine) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLNeuralNetworkUpdateEngine) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}
func (m MLNeuralNetworkUpdateEngine) ShuffableTrainingData() IMLShufflingBatchProvider {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("shuffableTrainingData"))
	return MLShufflingBatchProviderFromID(objc.ID(rv))
}
func (m MLNeuralNetworkUpdateEngine) SetShuffableTrainingData(value IMLShufflingBatchProvider) {
	objc.Send[struct{}](m.ID, objc.Sel("setShuffableTrainingData:"), value)
}
func (m MLNeuralNetworkUpdateEngine) Snapshot() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("snapshot"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkUpdateEngine) SetSnapshot(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setSnapshot:"), value)
}
func (m MLNeuralNetworkUpdateEngine) Task() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("task"))
	return objectivec.Object{ID: rv}
}
func (m MLNeuralNetworkUpdateEngine) SetTask(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setTask:"), value)
}
