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

// The class instance for the [MLNeuralNetworkMLComputeGraph] class.
var (
	_MLNeuralNetworkMLComputeGraphClass     MLNeuralNetworkMLComputeGraphClass
	_MLNeuralNetworkMLComputeGraphClassOnce sync.Once
)

func getMLNeuralNetworkMLComputeGraphClass() MLNeuralNetworkMLComputeGraphClass {
	_MLNeuralNetworkMLComputeGraphClassOnce.Do(func() {
		_MLNeuralNetworkMLComputeGraphClass = MLNeuralNetworkMLComputeGraphClass{class: objc.GetClass("MLNeuralNetworkMLComputeGraph")}
	})
	return _MLNeuralNetworkMLComputeGraphClass
}

// GetMLNeuralNetworkMLComputeGraphClass returns the class object for MLNeuralNetworkMLComputeGraph.
func GetMLNeuralNetworkMLComputeGraphClass() MLNeuralNetworkMLComputeGraphClass {
	return getMLNeuralNetworkMLComputeGraphClass()
}

type MLNeuralNetworkMLComputeGraphClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkMLComputeGraphClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkMLComputeGraphClass) Alloc() MLNeuralNetworkMLComputeGraph {
	rv := objc.Send[MLNeuralNetworkMLComputeGraph](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNeuralNetworkMLComputeGraph.BuildGraphsForBatchSizeNumberOfClassesError]
//   - [MLNeuralNetworkMLComputeGraph.BuildInferenceGraphUpdateParamsLayersMlcTensorByNameOutputNameToLayerMapError]
//   - [MLNeuralNetworkMLComputeGraph.BuildMLComputeTensorDescriptorWithFeatureNameBatchSizeNumberOfClassesError]
//   - [MLNeuralNetworkMLComputeGraph.BuildTrainingGraphFromUpdateParamsNumberOfClassesMlcTensorByNameError]
//   - [MLNeuralNetworkMLComputeGraph.ClassifierOutputIsSigmoidOutput]
//   - [MLNeuralNetworkMLComputeGraph.SetClassifierOutputIsSigmoidOutput]
//   - [MLNeuralNetworkMLComputeGraph.CopyWeightsFromToError]
//   - [MLNeuralNetworkMLComputeGraph.CreateMultiArrayFromTensorError]
//   - [MLNeuralNetworkMLComputeGraph.Device]
//   - [MLNeuralNetworkMLComputeGraph.SetDevice]
//   - [MLNeuralNetworkMLComputeGraph.ExecutionOptions]
//   - [MLNeuralNetworkMLComputeGraph.SetExecutionOptions]
//   - [MLNeuralNetworkMLComputeGraph.FusedLayerInputName]
//   - [MLNeuralNetworkMLComputeGraph.SetFusedLayerInputName]
//   - [MLNeuralNetworkMLComputeGraph.GetBiasesForLayerNamedError]
//   - [MLNeuralNetworkMLComputeGraph.GetWeightsForLayerNamedError]
//   - [MLNeuralNetworkMLComputeGraph.Graph]
//   - [MLNeuralNetworkMLComputeGraph.SetGraph]
//   - [MLNeuralNetworkMLComputeGraph.InferenceGraph]
//   - [MLNeuralNetworkMLComputeGraph.SetInferenceGraph]
//   - [MLNeuralNetworkMLComputeGraph.InputTensorMapWithBatchSizeNumberOfClassesError]
//   - [MLNeuralNetworkMLComputeGraph.LabelTensorMapWithBatchSizeNumberOfClassesError]
//   - [MLNeuralNetworkMLComputeGraph.LayerFusedToLoss]
//   - [MLNeuralNetworkMLComputeGraph.SetLayerFusedToLoss]
//   - [MLNeuralNetworkMLComputeGraph.LayersMap]
//   - [MLNeuralNetworkMLComputeGraph.SetLayersMap]
//   - [MLNeuralNetworkMLComputeGraph.LossInputsFromUpdateParams]
//   - [MLNeuralNetworkMLComputeGraph.MlcDeviceTypeForComputeUnit]
//   - [MLNeuralNetworkMLComputeGraph.MlcInputs]
//   - [MLNeuralNetworkMLComputeGraph.SetMlcInputs]
//   - [MLNeuralNetworkMLComputeGraph.MlcLabels]
//   - [MLNeuralNetworkMLComputeGraph.SetMlcLabels]
//   - [MLNeuralNetworkMLComputeGraph.ModelDescription]
//   - [MLNeuralNetworkMLComputeGraph.SetModelDescription]
//   - [MLNeuralNetworkMLComputeGraph.OutputNameToLayerMap]
//   - [MLNeuralNetworkMLComputeGraph.SaveUpdatedWeightsToError]
//   - [MLNeuralNetworkMLComputeGraph.TrainingGraph]
//   - [MLNeuralNetworkMLComputeGraph.SetTrainingGraph]
//   - [MLNeuralNetworkMLComputeGraph.InitWithCompiledArchiveModelDescriptionBatchSizeNumberOfClassesComputeUnitsError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph
type MLNeuralNetworkMLComputeGraph struct {
	objectivec.Object
}

// MLNeuralNetworkMLComputeGraphFromID constructs a [MLNeuralNetworkMLComputeGraph] from an objc.ID.
func MLNeuralNetworkMLComputeGraphFromID(id objc.ID) MLNeuralNetworkMLComputeGraph {
	return MLNeuralNetworkMLComputeGraph{objectivec.Object{ID: id}}
}

// Ensure MLNeuralNetworkMLComputeGraph implements IMLNeuralNetworkMLComputeGraph.
var _ IMLNeuralNetworkMLComputeGraph = MLNeuralNetworkMLComputeGraph{}

// An interface definition for the [MLNeuralNetworkMLComputeGraph] class.
//
// # Methods
//
//   - [IMLNeuralNetworkMLComputeGraph.BuildGraphsForBatchSizeNumberOfClassesError]
//   - [IMLNeuralNetworkMLComputeGraph.BuildInferenceGraphUpdateParamsLayersMlcTensorByNameOutputNameToLayerMapError]
//   - [IMLNeuralNetworkMLComputeGraph.BuildMLComputeTensorDescriptorWithFeatureNameBatchSizeNumberOfClassesError]
//   - [IMLNeuralNetworkMLComputeGraph.BuildTrainingGraphFromUpdateParamsNumberOfClassesMlcTensorByNameError]
//   - [IMLNeuralNetworkMLComputeGraph.ClassifierOutputIsSigmoidOutput]
//   - [IMLNeuralNetworkMLComputeGraph.SetClassifierOutputIsSigmoidOutput]
//   - [IMLNeuralNetworkMLComputeGraph.CopyWeightsFromToError]
//   - [IMLNeuralNetworkMLComputeGraph.CreateMultiArrayFromTensorError]
//   - [IMLNeuralNetworkMLComputeGraph.Device]
//   - [IMLNeuralNetworkMLComputeGraph.SetDevice]
//   - [IMLNeuralNetworkMLComputeGraph.ExecutionOptions]
//   - [IMLNeuralNetworkMLComputeGraph.SetExecutionOptions]
//   - [IMLNeuralNetworkMLComputeGraph.FusedLayerInputName]
//   - [IMLNeuralNetworkMLComputeGraph.SetFusedLayerInputName]
//   - [IMLNeuralNetworkMLComputeGraph.GetBiasesForLayerNamedError]
//   - [IMLNeuralNetworkMLComputeGraph.GetWeightsForLayerNamedError]
//   - [IMLNeuralNetworkMLComputeGraph.Graph]
//   - [IMLNeuralNetworkMLComputeGraph.SetGraph]
//   - [IMLNeuralNetworkMLComputeGraph.InferenceGraph]
//   - [IMLNeuralNetworkMLComputeGraph.SetInferenceGraph]
//   - [IMLNeuralNetworkMLComputeGraph.InputTensorMapWithBatchSizeNumberOfClassesError]
//   - [IMLNeuralNetworkMLComputeGraph.LabelTensorMapWithBatchSizeNumberOfClassesError]
//   - [IMLNeuralNetworkMLComputeGraph.LayerFusedToLoss]
//   - [IMLNeuralNetworkMLComputeGraph.SetLayerFusedToLoss]
//   - [IMLNeuralNetworkMLComputeGraph.LayersMap]
//   - [IMLNeuralNetworkMLComputeGraph.SetLayersMap]
//   - [IMLNeuralNetworkMLComputeGraph.LossInputsFromUpdateParams]
//   - [IMLNeuralNetworkMLComputeGraph.MlcDeviceTypeForComputeUnit]
//   - [IMLNeuralNetworkMLComputeGraph.MlcInputs]
//   - [IMLNeuralNetworkMLComputeGraph.SetMlcInputs]
//   - [IMLNeuralNetworkMLComputeGraph.MlcLabels]
//   - [IMLNeuralNetworkMLComputeGraph.SetMlcLabels]
//   - [IMLNeuralNetworkMLComputeGraph.ModelDescription]
//   - [IMLNeuralNetworkMLComputeGraph.SetModelDescription]
//   - [IMLNeuralNetworkMLComputeGraph.OutputNameToLayerMap]
//   - [IMLNeuralNetworkMLComputeGraph.SaveUpdatedWeightsToError]
//   - [IMLNeuralNetworkMLComputeGraph.TrainingGraph]
//   - [IMLNeuralNetworkMLComputeGraph.SetTrainingGraph]
//   - [IMLNeuralNetworkMLComputeGraph.InitWithCompiledArchiveModelDescriptionBatchSizeNumberOfClassesComputeUnitsError]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph
type IMLNeuralNetworkMLComputeGraph interface {
	objectivec.IObject

	// Topic: Methods

	BuildGraphsForBatchSizeNumberOfClassesError(for_ unsafe.Pointer, size uint64, classes uint64) (bool, error)
	BuildInferenceGraphUpdateParamsLayersMlcTensorByNameOutputNameToLayerMapError(graph objectivec.IObject, params unsafe.Pointer, layers unsafe.Pointer, name objectivec.IObject, map_ objectivec.IObject) (objectivec.IObject, error)
	BuildMLComputeTensorDescriptorWithFeatureNameBatchSizeNumberOfClassesError(with objectivec.IObject, name objectivec.IObject, size uint64, classes uint64) (objectivec.IObject, error)
	BuildTrainingGraphFromUpdateParamsNumberOfClassesMlcTensorByNameError(from objectivec.IObject, params unsafe.Pointer, classes uint64, name objectivec.IObject) (objectivec.IObject, error)
	ClassifierOutputIsSigmoidOutput() bool
	SetClassifierOutputIsSigmoidOutput(value bool)
	CopyWeightsFromToError(from objectivec.IObject, to unsafe.Pointer) (bool, error)
	CreateMultiArrayFromTensorError(tensor objectivec.IObject) (objectivec.IObject, error)
	Device() unsafe.Pointer
	SetDevice(value unsafe.Pointer)
	ExecutionOptions() uint64
	SetExecutionOptions(value uint64)
	FusedLayerInputName() string
	SetFusedLayerInputName(value string)
	GetBiasesForLayerNamedError(named objectivec.IObject) (objectivec.IObject, error)
	GetWeightsForLayerNamedError(named objectivec.IObject) (objectivec.IObject, error)
	Graph() unsafe.Pointer
	SetGraph(value unsafe.Pointer)
	InferenceGraph() unsafe.Pointer
	SetInferenceGraph(value unsafe.Pointer)
	InputTensorMapWithBatchSizeNumberOfClassesError(size uint64, classes uint64) (objectivec.IObject, error)
	LabelTensorMapWithBatchSizeNumberOfClassesError(size uint64, classes uint64) (objectivec.IObject, error)
	LayerFusedToLoss() bool
	SetLayerFusedToLoss(value bool)
	LayersMap() foundation.INSDictionary
	SetLayersMap(value foundation.INSDictionary)
	LossInputsFromUpdateParams(params unsafe.Pointer) objectivec.IObject
	MlcDeviceTypeForComputeUnit(unit int64) int
	MlcInputs() foundation.INSDictionary
	SetMlcInputs(value foundation.INSDictionary)
	MlcLabels() foundation.INSDictionary
	SetMlcLabels(value foundation.INSDictionary)
	ModelDescription() IMLModelDescription
	SetModelDescription(value IMLModelDescription)
	OutputNameToLayerMap() foundation.INSDictionary
	SaveUpdatedWeightsToError(to unsafe.Pointer) (bool, error)
	TrainingGraph() unsafe.Pointer
	SetTrainingGraph(value unsafe.Pointer)
	InitWithCompiledArchiveModelDescriptionBatchSizeNumberOfClassesComputeUnitsError(archive unsafe.Pointer, description objectivec.IObject, size uint64, classes uint64, units int64) (MLNeuralNetworkMLComputeGraph, error)
}

// Init initializes the instance.
func (n MLNeuralNetworkMLComputeGraph) Init() MLNeuralNetworkMLComputeGraph {
	rv := objc.Send[MLNeuralNetworkMLComputeGraph](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkMLComputeGraph) Autorelease() MLNeuralNetworkMLComputeGraph {
	rv := objc.Send[MLNeuralNetworkMLComputeGraph](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkMLComputeGraph creates a new MLNeuralNetworkMLComputeGraph instance.
func NewMLNeuralNetworkMLComputeGraph() MLNeuralNetworkMLComputeGraph {
	class := getMLNeuralNetworkMLComputeGraphClass()
	rv := objc.Send[MLNeuralNetworkMLComputeGraph](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/initWithCompiledArchive:modelDescription:batchSize:numberOfClasses:computeUnits:error:
func NewNeuralNetworkMLComputeGraphWithCompiledArchiveModelDescriptionBatchSizeNumberOfClassesComputeUnitsError(archive unsafe.Pointer, description objectivec.IObject, size uint64, classes uint64, units int64) (MLNeuralNetworkMLComputeGraph, error) {
	var errorPtr objc.ID
	instance := getMLNeuralNetworkMLComputeGraphClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompiledArchive:modelDescription:batchSize:numberOfClasses:computeUnits:error:"), archive, description, size, classes, units, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeGraphFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/buildGraphsFor:batchSize:numberOfClasses:error:
func (n MLNeuralNetworkMLComputeGraph) BuildGraphsForBatchSizeNumberOfClassesError(for_ unsafe.Pointer, size uint64, classes uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("buildGraphsFor:batchSize:numberOfClasses:error:"), for_, size, classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("buildGraphsFor:batchSize:numberOfClasses:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/buildInferenceGraph:updateParams:layers:mlcTensorByName:outputNameToLayerMap:error:
func (n MLNeuralNetworkMLComputeGraph) BuildInferenceGraphUpdateParamsLayersMlcTensorByNameOutputNameToLayerMapError(graph objectivec.IObject, params unsafe.Pointer, layers unsafe.Pointer, name objectivec.IObject, map_ objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("buildInferenceGraph:updateParams:layers:mlcTensorByName:outputNameToLayerMap:error:"), graph, params, layers, name, map_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/buildMLComputeTensorDescriptorWith:featureName:batchSize:numberOfClasses:error:
func (n MLNeuralNetworkMLComputeGraph) BuildMLComputeTensorDescriptorWithFeatureNameBatchSizeNumberOfClassesError(with objectivec.IObject, name objectivec.IObject, size uint64, classes uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("buildMLComputeTensorDescriptorWith:featureName:batchSize:numberOfClasses:error:"), with, name, size, classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/buildTrainingGraphFrom:updateParams:numberOfClasses:mlcTensorByName:error:
func (n MLNeuralNetworkMLComputeGraph) BuildTrainingGraphFromUpdateParamsNumberOfClassesMlcTensorByNameError(from objectivec.IObject, params unsafe.Pointer, classes uint64, name objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("buildTrainingGraphFrom:updateParams:numberOfClasses:mlcTensorByName:error:"), from, params, classes, name, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/copyWeightsFrom:to:error:
func (n MLNeuralNetworkMLComputeGraph) CopyWeightsFromToError(from objectivec.IObject, to unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("copyWeightsFrom:to:error:"), from, to, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("copyWeightsFrom:to:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/createMultiArrayFromTensor:error:
func (n MLNeuralNetworkMLComputeGraph) CreateMultiArrayFromTensorError(tensor objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("createMultiArrayFromTensor:error:"), tensor, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/getBiasesForLayerNamed:error:
func (n MLNeuralNetworkMLComputeGraph) GetBiasesForLayerNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("getBiasesForLayerNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/getWeightsForLayerNamed:error:
func (n MLNeuralNetworkMLComputeGraph) GetWeightsForLayerNamedError(named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("getWeightsForLayerNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/inputTensorMapWithBatchSize:numberOfClasses:error:
func (n MLNeuralNetworkMLComputeGraph) InputTensorMapWithBatchSizeNumberOfClassesError(size uint64, classes uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("inputTensorMapWithBatchSize:numberOfClasses:error:"), size, classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/labelTensorMapWithBatchSize:numberOfClasses:error:
func (n MLNeuralNetworkMLComputeGraph) LabelTensorMapWithBatchSizeNumberOfClassesError(size uint64, classes uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("labelTensorMapWithBatchSize:numberOfClasses:error:"), size, classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/lossInputsFromUpdateParams:
func (n MLNeuralNetworkMLComputeGraph) LossInputsFromUpdateParams(params unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("lossInputsFromUpdateParams:"), params)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/mlcDeviceTypeForComputeUnit:
func (n MLNeuralNetworkMLComputeGraph) MlcDeviceTypeForComputeUnit(unit int64) int {
	rv := objc.Send[int](n.ID, objc.Sel("mlcDeviceTypeForComputeUnit:"), unit)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/saveUpdatedWeightsTo:error:
func (n MLNeuralNetworkMLComputeGraph) SaveUpdatedWeightsToError(to unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](n.ID, objc.Sel("saveUpdatedWeightsTo:error:"), to, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveUpdatedWeightsTo:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/initWithCompiledArchive:modelDescription:batchSize:numberOfClasses:computeUnits:error:
func (n MLNeuralNetworkMLComputeGraph) InitWithCompiledArchiveModelDescriptionBatchSizeNumberOfClassesComputeUnitsError(archive unsafe.Pointer, description objectivec.IObject, size uint64, classes uint64, units int64) (MLNeuralNetworkMLComputeGraph, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](n.ID, objc.Sel("initWithCompiledArchive:modelDescription:batchSize:numberOfClasses:computeUnits:error:"), archive, description, size, classes, units, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLNeuralNetworkMLComputeGraph{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLNeuralNetworkMLComputeGraphFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/graphFromCompiledArchive:modelDescription:batchSize:numberOfClasses:computeUnits:error:
func (_MLNeuralNetworkMLComputeGraphClass MLNeuralNetworkMLComputeGraphClass) GraphFromCompiledArchiveModelDescriptionBatchSizeNumberOfClassesComputeUnitsError(archive unsafe.Pointer, description objectivec.IObject, size uint64, classes uint64, units int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkMLComputeGraphClass.class), objc.Sel("graphFromCompiledArchive:modelDescription:batchSize:numberOfClasses:computeUnits:error:"), archive, description, size, classes, units, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/classifierOutputIsSigmoidOutput
func (n MLNeuralNetworkMLComputeGraph) ClassifierOutputIsSigmoidOutput() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("classifierOutputIsSigmoidOutput"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetClassifierOutputIsSigmoidOutput(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setClassifierOutputIsSigmoidOutput:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/device
func (n MLNeuralNetworkMLComputeGraph) Device() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n.ID, objc.Sel("device"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetDevice(value unsafe.Pointer) {
	objc.Send[struct{}](n.ID, objc.Sel("setDevice:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/executionOptions
func (n MLNeuralNetworkMLComputeGraph) ExecutionOptions() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("executionOptions"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetExecutionOptions(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setExecutionOptions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/fusedLayerInputName
func (n MLNeuralNetworkMLComputeGraph) FusedLayerInputName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("fusedLayerInputName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNeuralNetworkMLComputeGraph) SetFusedLayerInputName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setFusedLayerInputName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/graph
func (n MLNeuralNetworkMLComputeGraph) Graph() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n.ID, objc.Sel("graph"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetGraph(value unsafe.Pointer) {
	objc.Send[struct{}](n.ID, objc.Sel("setGraph:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/inferenceGraph
func (n MLNeuralNetworkMLComputeGraph) InferenceGraph() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n.ID, objc.Sel("inferenceGraph"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetInferenceGraph(value unsafe.Pointer) {
	objc.Send[struct{}](n.ID, objc.Sel("setInferenceGraph:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/layerFusedToLoss
func (n MLNeuralNetworkMLComputeGraph) LayerFusedToLoss() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("layerFusedToLoss"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetLayerFusedToLoss(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setLayerFusedToLoss:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/layersMap
func (n MLNeuralNetworkMLComputeGraph) LayersMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("layersMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeGraph) SetLayersMap(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setLayersMap:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/mlcInputs
func (n MLNeuralNetworkMLComputeGraph) MlcInputs() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("mlcInputs"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeGraph) SetMlcInputs(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setMlcInputs:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/mlcLabels
func (n MLNeuralNetworkMLComputeGraph) MlcLabels() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("mlcLabels"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeGraph) SetMlcLabels(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setMlcLabels:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/modelDescription
func (n MLNeuralNetworkMLComputeGraph) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (n MLNeuralNetworkMLComputeGraph) SetModelDescription(value IMLModelDescription) {
	objc.Send[struct{}](n.ID, objc.Sel("setModelDescription:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/outputNameToLayerMap
func (n MLNeuralNetworkMLComputeGraph) OutputNameToLayerMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("outputNameToLayerMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkMLComputeGraph/trainingGraph
func (n MLNeuralNetworkMLComputeGraph) TrainingGraph() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](n.ID, objc.Sel("trainingGraph"))
	return rv
}
func (n MLNeuralNetworkMLComputeGraph) SetTrainingGraph(value unsafe.Pointer) {
	objc.Send[struct{}](n.ID, objc.Sel("setTrainingGraph:"), value)
}
