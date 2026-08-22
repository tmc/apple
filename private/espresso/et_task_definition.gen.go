// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETTaskDefinition] class.
var (
	_ETTaskDefinitionClass     ETTaskDefinitionClass
	_ETTaskDefinitionClassOnce sync.Once
)

func getETTaskDefinitionClass() ETTaskDefinitionClass {
	_ETTaskDefinitionClassOnce.Do(func() {
		_ETTaskDefinitionClass = ETTaskDefinitionClass{class: objc.GetClass("ETTaskDefinition")}
	})
	return _ETTaskDefinitionClass
}

// GetETTaskDefinitionClass returns the class object for ETTaskDefinition.
func GetETTaskDefinitionClass() ETTaskDefinitionClass {
	return getETTaskDefinitionClass()
}

type ETTaskDefinitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETTaskDefinitionClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETTaskDefinitionClass) Alloc() ETTaskDefinition {
	rv := objc.SendIfResponds[ETTaskDefinition](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETTaskDefinition.CheckShapesWithSampleWithError]
//   - [ETTaskDefinition.DataTypeForParameterOfTypeFromLayerNamed]
//   - [ETTaskDefinition.DoInferenceOnDataError]
//   - [ETTaskDefinition.DoTrainingOnDataForNumberOfEpochsWithCallbackError]
//   - [ETTaskDefinition.GetParameterOfTypeForLayerNamedError]
//   - [ETTaskDefinition.GetTensorNamed]
//   - [ETTaskDefinition.GetTensorNamedDirectBind]
//   - [ETTaskDefinition.InferenceGraphNetPtr]
//   - [ETTaskDefinition.SetInferenceGraphNetPtr]
//   - [ETTaskDefinition.InferenceModel]
//   - [ETTaskDefinition.SetInferenceModel]
//   - [ETTaskDefinition.Optimizer]
//   - [ETTaskDefinition.SetOptimizer]
//   - [ETTaskDefinition.Platform]
//   - [ETTaskDefinition.SetPlatform]
//   - [ETTaskDefinition.PlatformForLayerNamedError]
//   - [ETTaskDefinition.PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError]
//   - [ETTaskDefinition.ReloadOnRuntimePlatform]
//   - [ETTaskDefinition.SaveNetworkInplaceError]
//   - [ETTaskDefinition.SaveTrainingNetworkCheckpointError]
//   - [ETTaskDefinition.SetParameterOfTypeForLayerNamedWithValueError]
//   - [ETTaskDefinition.SetTensorNamedWithValueError]
//   - [ETTaskDefinition.SetWeightsOfInferenceNetworkLoadedFromAndSaveToError]
//   - [ETTaskDefinition.SetupInputOutputShapes]
//   - [ETTaskDefinition.SetupShapesForBlobsWithError]
//   - [ETTaskDefinition.ShareWeights]
//   - [ETTaskDefinition.TaskState]
//   - [ETTaskDefinition.SetTaskState]
//   - [ETTaskDefinition.TrainingGraphNetPtr]
//   - [ETTaskDefinition.SetTrainingGraphNetPtr]
//   - [ETTaskDefinition.InitWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError]
//   - [ETTaskDefinition.InitWithTrainingModelDefinitionForPlatformError]
type ETTaskDefinition struct {
	objectivec.Object
}

// ETTaskDefinitionFromID constructs a [ETTaskDefinition] from an objc.ID.
func ETTaskDefinitionFromID(id objc.ID) ETTaskDefinition {
	return ETTaskDefinition{objectivec.Object{ID: id}}
}

// Ensure ETTaskDefinition implements IETTaskDefinition.
var _ IETTaskDefinition = ETTaskDefinition{}

// An interface definition for the [ETTaskDefinition] class.
//
// # Methods
//
//   - [IETTaskDefinition.CheckShapesWithSampleWithError]
//   - [IETTaskDefinition.DataTypeForParameterOfTypeFromLayerNamed]
//   - [IETTaskDefinition.DoInferenceOnDataError]
//   - [IETTaskDefinition.DoTrainingOnDataForNumberOfEpochsWithCallbackError]
//   - [IETTaskDefinition.GetParameterOfTypeForLayerNamedError]
//   - [IETTaskDefinition.GetTensorNamed]
//   - [IETTaskDefinition.GetTensorNamedDirectBind]
//   - [IETTaskDefinition.InferenceGraphNetPtr]
//   - [IETTaskDefinition.SetInferenceGraphNetPtr]
//   - [IETTaskDefinition.InferenceModel]
//   - [IETTaskDefinition.SetInferenceModel]
//   - [IETTaskDefinition.Optimizer]
//   - [IETTaskDefinition.SetOptimizer]
//   - [IETTaskDefinition.Platform]
//   - [IETTaskDefinition.SetPlatform]
//   - [IETTaskDefinition.PlatformForLayerNamedError]
//   - [IETTaskDefinition.PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError]
//   - [IETTaskDefinition.ReloadOnRuntimePlatform]
//   - [IETTaskDefinition.SaveNetworkInplaceError]
//   - [IETTaskDefinition.SaveTrainingNetworkCheckpointError]
//   - [IETTaskDefinition.SetParameterOfTypeForLayerNamedWithValueError]
//   - [IETTaskDefinition.SetTensorNamedWithValueError]
//   - [IETTaskDefinition.SetWeightsOfInferenceNetworkLoadedFromAndSaveToError]
//   - [IETTaskDefinition.SetupInputOutputShapes]
//   - [IETTaskDefinition.SetupShapesForBlobsWithError]
//   - [IETTaskDefinition.ShareWeights]
//   - [IETTaskDefinition.TaskState]
//   - [IETTaskDefinition.SetTaskState]
//   - [IETTaskDefinition.TrainingGraphNetPtr]
//   - [IETTaskDefinition.SetTrainingGraphNetPtr]
//   - [IETTaskDefinition.InitWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError]
//   - [IETTaskDefinition.InitWithTrainingModelDefinitionForPlatformError]
type IETTaskDefinition interface {
	objectivec.IObject

	// Topic: Methods

	CheckShapesWithSampleWithError(shapes unsafe.Pointer, sample unsafe.Pointer) error
	DataTypeForParameterOfTypeFromLayerNamed(type_ uint64, named objectivec.IObject) uint64
	DoInferenceOnDataError(data objectivec.IObject) (objectivec.IObject, error)
	DoTrainingOnDataForNumberOfEpochsWithCallbackError(data objectivec.IObject, epochs uint64, callback objectivec.IObject) (bool, error)
	GetParameterOfTypeForLayerNamedError(type_ uint64, named objectivec.IObject) (objectivec.IObject, error)
	GetTensorNamed(named objectivec.IObject) objectivec.IObject
	GetTensorNamedDirectBind(named objectivec.IObject, bind bool) objectivec.IObject
	InferenceGraphNetPtr() unsafe.Pointer
	SetInferenceGraphNetPtr(value unsafe.Pointer)
	InferenceModel() IETModelDefinition
	SetInferenceModel(value IETModelDefinition)
	Optimizer() IETOptimizerDefinition
	SetOptimizer(value IETOptimizerDefinition)
	Platform() uint64
	SetPlatform(value uint64)
	PlatformForLayerNamedError(named objectivec.IObject) (uint64, error)
	PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError(data objectivec.IObject, epochs uint64, callback objectivec.IObject) (bool, error)
	ReloadOnRuntimePlatform(platform []objectivec.IObject)
	SaveNetworkInplaceError(network objectivec.IObject, inplace bool) (bool, error)
	SaveTrainingNetworkCheckpointError(network objectivec.IObject, checkpoint objectivec.IObject) (bool, error)
	SetParameterOfTypeForLayerNamedWithValueError(type_ uint64, named objectivec.IObject, value objectivec.IObject) (bool, error)
	SetTensorNamedWithValueError(named objectivec.IObject, value objectivec.IObject) (bool, error)
	SetWeightsOfInferenceNetworkLoadedFromAndSaveToError(from objectivec.IObject, to objectivec.IObject) (bool, error)
	SetupInputOutputShapes(shapes []objectivec.IObject)
	SetupShapesForBlobsWithError(shapes unsafe.Pointer, blobs objectivec.IObject) error
	ShareWeights()
	TaskState() IETTaskState
	SetTaskState(value IETTaskState)
	TrainingGraphNetPtr() unsafe.Pointer
	SetTrainingGraphNetPtr(value unsafe.Pointer)
	InitWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError(definition objectivec.IObject, definition2 objectivec.IObject, definition3 objectivec.IObject, definition4 objectivec.IObject, platform uint64) (ETTaskDefinition, error)
	InitWithTrainingModelDefinitionForPlatformError(definition objectivec.IObject, platform uint64) (ETTaskDefinition, error)
}

// Init initializes the instance.
func (e ETTaskDefinition) Init() ETTaskDefinition {
	rv := objc.SendIfResponds[ETTaskDefinition](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETTaskDefinition) Autorelease() ETTaskDefinition {
	rv := objc.SendIfResponds[ETTaskDefinition](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETTaskDefinition creates a new ETTaskDefinition instance.
func NewETTaskDefinition() ETTaskDefinition {
	class := getETTaskDefinitionClass()
	rv := objc.SendIfResponds[ETTaskDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewETTaskDefinitionWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError(definition objectivec.IObject, definition2 objectivec.IObject, definition3 objectivec.IObject, definition4 objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	instance := getETTaskDefinitionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelDefinition:lossDefinition:variablesDefinition:optimizerDefinition:forPlatform:error:"), definition, definition2, definition3, definition4, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ETTaskDefinition{}, objc.ErrInitFailed
	}
	return ETTaskDefinitionFromID(rv), nil
}

func NewETTaskDefinitionWithTrainingModelDefinitionForPlatformError(definition objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	instance := getETTaskDefinitionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTrainingModelDefinition:forPlatform:error:"), definition, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ETTaskDefinition{}, objc.ErrInitFailed
	}
	return ETTaskDefinitionFromID(rv), nil
}

func (e ETTaskDefinition) CheckShapesWithSampleWithError(shapes unsafe.Pointer, sample unsafe.Pointer) error {
	var errorPtr objc.ID
	objc.Send[struct{}](e.ID, objc.Sel("checkShapes:withSample:withError:"), shapes, sample, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}
func (e ETTaskDefinition) DataTypeForParameterOfTypeFromLayerNamed(type_ uint64, named objectivec.IObject) uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("dataTypeForParameterOfType:fromLayerNamed:"), type_, named)
	return rv
}
func (e ETTaskDefinition) DoInferenceOnDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("doInferenceOnData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (e ETTaskDefinition) DoTrainingOnDataForNumberOfEpochsWithCallbackError(data objectivec.IObject, epochs uint64, callback objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("doTrainingOnData:forNumberOfEpochs:withCallback:error:"), data, epochs, callback, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doTrainingOnData:forNumberOfEpochs:withCallback:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) GetParameterOfTypeForLayerNamedError(type_ uint64, named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getParameterOfType:forLayerNamed:error:"), type_, named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (e ETTaskDefinition) GetTensorNamed(named objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("getTensorNamed:"), named)
	return objectivec.Object{ID: rv}
}
func (e ETTaskDefinition) GetTensorNamedDirectBind(named objectivec.IObject, bind bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("getTensorNamed:directBind:"), named, bind)
	return objectivec.Object{ID: rv}
}
func (e ETTaskDefinition) PlatformForLayerNamedError(named objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](e.ID, objc.Sel("platformForLayerNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (e ETTaskDefinition) PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError(data objectivec.IObject, epochs uint64, callback objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("privateDoTrainingOnData:forNumberOfEpochs:withCallback:error:"), data, epochs, callback, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("privateDoTrainingOnData:forNumberOfEpochs:withCallback:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) ReloadOnRuntimePlatform(platform []objectivec.IObject) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("reloadOnRuntimePlatform:"), objectivec.IObjectSliceToNSArray(platform))
}
func (e ETTaskDefinition) SaveNetworkInplaceError(network objectivec.IObject, inplace bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("saveNetwork:inplace:error:"), network, inplace, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveNetwork:inplace:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) SaveTrainingNetworkCheckpointError(network objectivec.IObject, checkpoint objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("saveTrainingNetwork:checkpoint:error:"), network, checkpoint, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("saveTrainingNetwork:checkpoint:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) SetParameterOfTypeForLayerNamedWithValueError(type_ uint64, named objectivec.IObject, value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("setParameterOfType:forLayerNamed:withValue:error:"), type_, named, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setParameterOfType:forLayerNamed:withValue:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) SetTensorNamedWithValueError(named objectivec.IObject, value objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("setTensorNamed:withValue:error:"), named, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setTensorNamed:withValue:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) SetWeightsOfInferenceNetworkLoadedFromAndSaveToError(from objectivec.IObject, to objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("setWeightsOfInferenceNetworkLoadedFrom:AndSaveTo:error:"), from, to, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setWeightsOfInferenceNetworkLoadedFrom:AndSaveTo:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (e ETTaskDefinition) SetupInputOutputShapes(shapes []objectivec.IObject) {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("setupInputOutputShapes:"), objectivec.IObjectSliceToNSArray(shapes))
}
func (e ETTaskDefinition) SetupShapesForBlobsWithError(shapes unsafe.Pointer, blobs objectivec.IObject) error {
	var errorPtr objc.ID
	objc.Send[struct{}](e.ID, objc.Sel("setupShapes:forBlobs:withError:"), shapes, blobs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}
func (e ETTaskDefinition) ShareWeights() {
	objc.SendIfResponds[objc.ID](e.ID, objc.Sel("shareWeights"))
}
func (e ETTaskDefinition) InitWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError(definition objectivec.IObject, definition2 objectivec.IObject, definition3 objectivec.IObject, definition4 objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithModelDefinition:lossDefinition:variablesDefinition:optimizerDefinition:forPlatform:error:"), definition, definition2, definition3, definition4, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETTaskDefinitionFromID(rv), nil

}
func (e ETTaskDefinition) InitWithTrainingModelDefinitionForPlatformError(definition objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithTrainingModelDefinition:forPlatform:error:"), definition, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETTaskDefinitionFromID(rv), nil

}

func (e ETTaskDefinition) InferenceGraphNetPtr() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("inferenceGraphNetPtr"))
	return rv
}
func (e ETTaskDefinition) SetInferenceGraphNetPtr(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setInferenceGraphNetPtr:"), value)
}
func (e ETTaskDefinition) InferenceModel() IETModelDefinition {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("inferenceModel"))
	return ETModelDefinitionFromID(objc.ID(rv))
}
func (e ETTaskDefinition) SetInferenceModel(value IETModelDefinition) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setInferenceModel:"), value)
}
func (e ETTaskDefinition) Optimizer() IETOptimizerDefinition {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("optimizer"))
	return ETOptimizerDefinitionFromID(objc.ID(rv))
}
func (e ETTaskDefinition) SetOptimizer(value IETOptimizerDefinition) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setOptimizer:"), value)
}
func (e ETTaskDefinition) Platform() uint64 {
	rv := objc.SendIfResponds[uint64](e.ID, objc.Sel("platform"))
	return rv
}
func (e ETTaskDefinition) SetPlatform(value uint64) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setPlatform:"), value)
}
func (e ETTaskDefinition) TaskState() IETTaskState {
	rv := objc.SendIfResponds[objc.ID](e.ID, objc.Sel("taskState"))
	return ETTaskStateFromID(objc.ID(rv))
}
func (e ETTaskDefinition) SetTaskState(value IETTaskState) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setTaskState:"), value)
}
func (e ETTaskDefinition) TrainingGraphNetPtr() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](e.ID, objc.Sel("trainingGraphNetPtr"))
	return rv
}
func (e ETTaskDefinition) SetTrainingGraphNetPtr(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](e.ID, objc.Sel("setTrainingGraphNetPtr:"), value)
}
