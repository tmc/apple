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
	rv := objc.Send[ETTaskDefinition](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ETTaskDefinition.CheckShapesWithSampleWithError]
//   - [ETTaskDefinition.Context_for_runtime_platform]
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
//   - [ETTaskDefinition.NamesVectorToFoundationArray]
//   - [ETTaskDefinition.Optimizer]
//   - [ETTaskDefinition.SetOptimizer]
//   - [ETTaskDefinition.Platform]
//   - [ETTaskDefinition.SetPlatform]
//   - [ETTaskDefinition.PlatformForLayerNamedError]
//   - [ETTaskDefinition.PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError]
//   - [ETTaskDefinition.ReloadOnRuntimePlatform]
//   - [ETTaskDefinition.SaveNetworkInplaceError]
//   - [ETTaskDefinition.SaveTrainingNetworkCheckpointError]
//   - [ETTaskDefinition.SetInferenceNetworkWeightsError]
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
//
// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition
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
//   - [IETTaskDefinition.Context_for_runtime_platform]
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
//   - [IETTaskDefinition.NamesVectorToFoundationArray]
//   - [IETTaskDefinition.Optimizer]
//   - [IETTaskDefinition.SetOptimizer]
//   - [IETTaskDefinition.Platform]
//   - [IETTaskDefinition.SetPlatform]
//   - [IETTaskDefinition.PlatformForLayerNamedError]
//   - [IETTaskDefinition.PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError]
//   - [IETTaskDefinition.ReloadOnRuntimePlatform]
//   - [IETTaskDefinition.SaveNetworkInplaceError]
//   - [IETTaskDefinition.SaveTrainingNetworkCheckpointError]
//   - [IETTaskDefinition.SetInferenceNetworkWeightsError]
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
//
// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition
type IETTaskDefinition interface {
	objectivec.IObject

	// Topic: Methods

	CheckShapesWithSampleWithError(shapes unsafe.Pointer, sample unsafe.Pointer) error
	Context_for_runtime_platform(context_for_runtime_platform []objectivec.IObject) objectivec.IObject
	DataTypeForParameterOfTypeFromLayerNamed(type_ uint64, named objectivec.IObject) uint64
	DoInferenceOnDataError(data objectivec.IObject) (objectivec.IObject, error)
	DoTrainingOnDataForNumberOfEpochsWithCallbackError(data objectivec.IObject, epochs uint64, callback objectivec.IObject) (bool, error)
	GetParameterOfTypeForLayerNamedError(type_ uint64, named objectivec.IObject) (objectivec.IObject, error)
	GetTensorNamed(named objectivec.IObject) objectivec.IObject
	GetTensorNamedDirectBind(named objectivec.IObject, bind bool) objectivec.IObject
	InferenceGraphNetPtr() objectivec.IObject
	SetInferenceGraphNetPtr(value objectivec.IObject)
	InferenceModel() IETModelDefinition
	SetInferenceModel(value IETModelDefinition)
	NamesVectorToFoundationArray(array objectivec.IObject) objectivec.IObject
	Optimizer() IETOptimizerDefinition
	SetOptimizer(value IETOptimizerDefinition)
	Platform() uint64
	SetPlatform(value uint64)
	PlatformForLayerNamedError(named objectivec.IObject) (uint64, error)
	PrivateDoTrainingOnDataForNumberOfEpochsWithCallbackError(data objectivec.IObject, epochs uint64, callback objectivec.IObject) (bool, error)
	ReloadOnRuntimePlatform(platform []objectivec.IObject)
	SaveNetworkInplaceError(network objectivec.IObject, inplace bool) (bool, error)
	SaveTrainingNetworkCheckpointError(network objectivec.IObject, checkpoint objectivec.IObject) (bool, error)
	SetInferenceNetworkWeightsError(weights objectivec.IObject) (bool, error)
	SetParameterOfTypeForLayerNamedWithValueError(type_ uint64, named objectivec.IObject, value objectivec.IObject) (bool, error)
	SetTensorNamedWithValueError(named objectivec.IObject, value objectivec.IObject) (bool, error)
	SetWeightsOfInferenceNetworkLoadedFromAndSaveToError(from objectivec.IObject, to objectivec.IObject) (bool, error)
	SetupInputOutputShapes(shapes []objectivec.IObject)
	SetupShapesForBlobsWithError(shapes unsafe.Pointer, blobs objectivec.IObject) error
	ShareWeights()
	TaskState() IETTaskState
	SetTaskState(value IETTaskState)
	TrainingGraphNetPtr() objectivec.IObject
	SetTrainingGraphNetPtr(value objectivec.IObject)
	InitWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError(definition objectivec.IObject, definition2 objectivec.IObject, definition3 objectivec.IObject, definition4 objectivec.IObject, platform uint64) (ETTaskDefinition, error)
	InitWithTrainingModelDefinitionForPlatformError(definition objectivec.IObject, platform uint64) (ETTaskDefinition, error)
}

// Init initializes the instance.
func (e ETTaskDefinition) Init() ETTaskDefinition {
	rv := objc.Send[ETTaskDefinition](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETTaskDefinition) Autorelease() ETTaskDefinition {
	rv := objc.Send[ETTaskDefinition](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETTaskDefinition creates a new ETTaskDefinition instance.
func NewETTaskDefinition() ETTaskDefinition {
	class := getETTaskDefinitionClass()
	rv := objc.Send[ETTaskDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/initWithModelDefinition:lossDefinition:variablesDefinition:optimizerDefinition:forPlatform:error:
func NewETTaskDefinitionWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError(definition objectivec.IObject, definition2 objectivec.IObject, definition3 objectivec.IObject, definition4 objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	instance := getETTaskDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDefinition:lossDefinition:variablesDefinition:optimizerDefinition:forPlatform:error:"), definition, definition2, definition3, definition4, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETTaskDefinitionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/initWithTrainingModelDefinition:forPlatform:error:
func NewETTaskDefinitionWithTrainingModelDefinitionForPlatformError(definition objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	instance := getETTaskDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTrainingModelDefinition:forPlatform:error:"), definition, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETTaskDefinitionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/checkShapes:withSample:withError:
func (e ETTaskDefinition) CheckShapesWithSampleWithError(shapes unsafe.Pointer, sample unsafe.Pointer) error {
	var errorPtr objc.ID
	objc.Send[struct{}](e.ID, objc.Sel("checkShapes:withSample:withError:"), shapes, sample, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/context_for_runtime_platform:
func (e ETTaskDefinition) Context_for_runtime_platform(context_for_runtime_platform []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("context_for_runtime_platform:"), objectivec.IObjectSliceToNSArray(context_for_runtime_platform))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/dataTypeForParameterOfType:fromLayerNamed:
func (e ETTaskDefinition) DataTypeForParameterOfTypeFromLayerNamed(type_ uint64, named objectivec.IObject) uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("dataTypeForParameterOfType:fromLayerNamed:"), type_, named)
	return rv
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/doInferenceOnData:error:
func (e ETTaskDefinition) DoInferenceOnDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("doInferenceOnData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/doTrainingOnData:forNumberOfEpochs:withCallback:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/getParameterOfType:forLayerNamed:error:
func (e ETTaskDefinition) GetParameterOfTypeForLayerNamedError(type_ uint64, named objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getParameterOfType:forLayerNamed:error:"), type_, named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/getTensorNamed:
func (e ETTaskDefinition) GetTensorNamed(named objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getTensorNamed:"), named)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/getTensorNamed:directBind:
func (e ETTaskDefinition) GetTensorNamedDirectBind(named objectivec.IObject, bind bool) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("getTensorNamed:directBind:"), named, bind)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/namesVectorToFoundationArray:
func (e ETTaskDefinition) NamesVectorToFoundationArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("namesVectorToFoundationArray:"), array)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/platformForLayerNamed:error:
func (e ETTaskDefinition) PlatformForLayerNamedError(named objectivec.IObject) (uint64, error) {
	var errorPtr objc.ID
	rv := objc.Send[uint64](e.ID, objc.Sel("platformForLayerNamed:error:"), named, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/privateDoTrainingOnData:forNumberOfEpochs:withCallback:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/reloadOnRuntimePlatform:
func (e ETTaskDefinition) ReloadOnRuntimePlatform(platform []objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("reloadOnRuntimePlatform:"), objectivec.IObjectSliceToNSArray(platform))
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/saveNetwork:inplace:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/saveTrainingNetwork:checkpoint:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/setInferenceNetworkWeights:error:
func (e ETTaskDefinition) SetInferenceNetworkWeightsError(weights objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("setInferenceNetworkWeights:error:"), weights, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setInferenceNetworkWeights:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/setParameterOfType:forLayerNamed:withValue:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/setTensorNamed:withValue:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/setWeightsOfInferenceNetworkLoadedFrom:AndSaveTo:error:
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

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/setupInputOutputShapes:
func (e ETTaskDefinition) SetupInputOutputShapes(shapes []objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setupInputOutputShapes:"), objectivec.IObjectSliceToNSArray(shapes))
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/setupShapes:forBlobs:withError:
func (e ETTaskDefinition) SetupShapesForBlobsWithError(shapes unsafe.Pointer, blobs objectivec.IObject) error {
	var errorPtr objc.ID
	objc.Send[struct{}](e.ID, objc.Sel("setupShapes:forBlobs:withError:"), shapes, blobs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/shareWeights
func (e ETTaskDefinition) ShareWeights() {
	objc.Send[objc.ID](e.ID, objc.Sel("shareWeights"))
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/initWithModelDefinition:lossDefinition:variablesDefinition:optimizerDefinition:forPlatform:error:
func (e ETTaskDefinition) InitWithModelDefinitionLossDefinitionVariablesDefinitionOptimizerDefinitionForPlatformError(definition objectivec.IObject, definition2 objectivec.IObject, definition3 objectivec.IObject, definition4 objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithModelDefinition:lossDefinition:variablesDefinition:optimizerDefinition:forPlatform:error:"), definition, definition2, definition3, definition4, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETTaskDefinitionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/initWithTrainingModelDefinition:forPlatform:error:
func (e ETTaskDefinition) InitWithTrainingModelDefinitionForPlatformError(definition objectivec.IObject, platform uint64) (ETTaskDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithTrainingModelDefinition:forPlatform:error:"), definition, platform, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETTaskDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETTaskDefinitionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/inferenceGraphNetPtr
func (e ETTaskDefinition) InferenceGraphNetPtr() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inferenceGraphNetPtr"))
	return objectivec.Object{ID: rv}
}
func (e ETTaskDefinition) SetInferenceGraphNetPtr(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setInferenceGraphNetPtr:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/inferenceModel
func (e ETTaskDefinition) InferenceModel() IETModelDefinition {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inferenceModel"))
	return ETModelDefinitionFromID(objc.ID(rv))
}
func (e ETTaskDefinition) SetInferenceModel(value IETModelDefinition) {
	objc.Send[struct{}](e.ID, objc.Sel("setInferenceModel:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/optimizer
func (e ETTaskDefinition) Optimizer() IETOptimizerDefinition {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("optimizer"))
	return ETOptimizerDefinitionFromID(objc.ID(rv))
}
func (e ETTaskDefinition) SetOptimizer(value IETOptimizerDefinition) {
	objc.Send[struct{}](e.ID, objc.Sel("setOptimizer:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/platform
func (e ETTaskDefinition) Platform() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("platform"))
	return rv
}
func (e ETTaskDefinition) SetPlatform(value uint64) {
	objc.Send[struct{}](e.ID, objc.Sel("setPlatform:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/taskState
func (e ETTaskDefinition) TaskState() IETTaskState {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("taskState"))
	return ETTaskStateFromID(objc.ID(rv))
}
func (e ETTaskDefinition) SetTaskState(value IETTaskState) {
	objc.Send[struct{}](e.ID, objc.Sel("setTaskState:"), value)
}

// See: https://developer.apple.com/documentation/Espresso/ETTaskDefinition/trainingGraphNetPtr
func (e ETTaskDefinition) TrainingGraphNetPtr() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("trainingGraphNetPtr"))
	return objectivec.Object{ID: rv}
}
func (e ETTaskDefinition) SetTrainingGraphNetPtr(value objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setTrainingGraphNetPtr:"), value)
}
