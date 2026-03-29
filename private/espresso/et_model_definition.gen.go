// Code generated from Apple documentation for espresso. DO NOT EDIT.

package espresso

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ETModelDefinition] class.
var (
	_ETModelDefinitionClass     ETModelDefinitionClass
	_ETModelDefinitionClassOnce sync.Once
)

func getETModelDefinitionClass() ETModelDefinitionClass {
	_ETModelDefinitionClassOnce.Do(func() {
		_ETModelDefinitionClass = ETModelDefinitionClass{class: objc.GetClass("ETModelDefinition")}
	})
	return _ETModelDefinitionClass
}

// GetETModelDefinitionClass returns the class object for ETModelDefinition.
func GetETModelDefinitionClass() ETModelDefinitionClass {
	return getETModelDefinitionClass()
}

type ETModelDefinitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ec ETModelDefinitionClass) Class() objc.Class {
	return ec.class
}

// Alloc allocates memory for a new instance of the class.
func (ec ETModelDefinitionClass) Alloc() ETModelDefinition {
	rv := objc.Send[ETModelDefinition](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ETModelDefinition.Initializer]
//   - [ETModelDefinition.Inputs]
//   - [ETModelDefinition.IsTrainingGlobalName]
//   - [ETModelDefinition.LayerNames]
//   - [ETModelDefinition.ModelURL]
//   - [ETModelDefinition.Outputs]
//   - [ETModelDefinition.TrainingInputs]
//   - [ETModelDefinition.TrainingOutputs]
//   - [ETModelDefinition.InitWithInferenceNetworkPathError]
//   - [ETModelDefinition.InitWithInferenceNetworkPathInferenceInputsInferenceOutputsError]
//   - [ETModelDefinition.InitWithTrainingNetworkPathInferenceInputsInferenceOutputsTrainingInputsTrainingOutputsTrainingControlVariableNameWithInitializerError]
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition
type ETModelDefinition struct {
	objectivec.Object
}

// ETModelDefinitionFromID constructs a [ETModelDefinition] from an objc.ID.
func ETModelDefinitionFromID(id objc.ID) ETModelDefinition {
	return ETModelDefinition{objectivec.Object{ID: id}}
}
// Ensure ETModelDefinition implements IETModelDefinition.
var _ IETModelDefinition = ETModelDefinition{}

// An interface definition for the [ETModelDefinition] class.
//
// # Methods
//
//   - [IETModelDefinition.Initializer]
//   - [IETModelDefinition.Inputs]
//   - [IETModelDefinition.IsTrainingGlobalName]
//   - [IETModelDefinition.LayerNames]
//   - [IETModelDefinition.ModelURL]
//   - [IETModelDefinition.Outputs]
//   - [IETModelDefinition.TrainingInputs]
//   - [IETModelDefinition.TrainingOutputs]
//   - [IETModelDefinition.InitWithInferenceNetworkPathError]
//   - [IETModelDefinition.InitWithInferenceNetworkPathInferenceInputsInferenceOutputsError]
//   - [IETModelDefinition.InitWithTrainingNetworkPathInferenceInputsInferenceOutputsTrainingInputsTrainingOutputsTrainingControlVariableNameWithInitializerError]
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition
type IETModelDefinition interface {
	objectivec.IObject

	// Topic: Methods

	Initializer() string
	Inputs() foundation.INSArray
	IsTrainingGlobalName() string
	LayerNames() foundation.INSArray
	ModelURL() foundation.INSURL
	Outputs() foundation.INSArray
	TrainingInputs() foundation.INSArray
	TrainingOutputs() foundation.INSArray
	InitWithInferenceNetworkPathError(path objectivec.IObject) (ETModelDefinition, error)
	InitWithInferenceNetworkPathInferenceInputsInferenceOutputsError(path objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject) (ETModelDefinition, error)
	InitWithTrainingNetworkPathInferenceInputsInferenceOutputsTrainingInputsTrainingOutputsTrainingControlVariableNameWithInitializerError(path objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject, inputs2 objectivec.IObject, outputs2 objectivec.IObject, name objectivec.IObject, initializer objectivec.IObject) (ETModelDefinition, error)
}

// Init initializes the instance.
func (e ETModelDefinition) Init() ETModelDefinition {
	rv := objc.Send[ETModelDefinition](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e ETModelDefinition) Autorelease() ETModelDefinition {
	rv := objc.Send[ETModelDefinition](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewETModelDefinition creates a new ETModelDefinition instance.
func NewETModelDefinition() ETModelDefinition {
	class := getETModelDefinitionClass()
	rv := objc.Send[ETModelDefinition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initWithInferenceNetworkPath:error:
func NewETModelDefinitionWithInferenceNetworkPathError(path objectivec.IObject) (ETModelDefinition, error) {
	var errorPtr objc.ID
	instance := getETModelDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInferenceNetworkPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETModelDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETModelDefinitionFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initWithInferenceNetworkPath:inferenceInputs:inferenceOutputs:error:
func NewETModelDefinitionWithInferenceNetworkPathInferenceInputsInferenceOutputsError(path objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject) (ETModelDefinition, error) {
	var errorPtr objc.ID
	instance := getETModelDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInferenceNetworkPath:inferenceInputs:inferenceOutputs:error:"), path, inputs, outputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETModelDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETModelDefinitionFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initWithTrainingNetworkPath:inferenceInputs:inferenceOutputs:trainingInputs:trainingOutputs:trainingControlVariableName:withInitializer:error:
func NewETModelDefinitionWithTrainingNetworkPathInferenceInputsInferenceOutputsTrainingInputsTrainingOutputsTrainingControlVariableNameWithInitializerError(path objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject, inputs2 objectivec.IObject, outputs2 objectivec.IObject, name objectivec.IObject, initializer objectivec.IObject) (ETModelDefinition, error) {
	var errorPtr objc.ID
	instance := getETModelDefinitionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTrainingNetworkPath:inferenceInputs:inferenceOutputs:trainingInputs:trainingOutputs:trainingControlVariableName:withInitializer:error:"), path, inputs, outputs, inputs2, outputs2, name, initializer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETModelDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETModelDefinitionFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initWithInferenceNetworkPath:error:
func (e ETModelDefinition) InitWithInferenceNetworkPathError(path objectivec.IObject) (ETModelDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithInferenceNetworkPath:error:"), path, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETModelDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETModelDefinitionFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initWithInferenceNetworkPath:inferenceInputs:inferenceOutputs:error:
func (e ETModelDefinition) InitWithInferenceNetworkPathInferenceInputsInferenceOutputsError(path objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject) (ETModelDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithInferenceNetworkPath:inferenceInputs:inferenceOutputs:error:"), path, inputs, outputs, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETModelDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETModelDefinitionFromID(rv), nil

}
//
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initWithTrainingNetworkPath:inferenceInputs:inferenceOutputs:trainingInputs:trainingOutputs:trainingControlVariableName:withInitializer:error:
func (e ETModelDefinition) InitWithTrainingNetworkPathInferenceInputsInferenceOutputsTrainingInputsTrainingOutputsTrainingControlVariableNameWithInitializerError(path objectivec.IObject, inputs objectivec.IObject, outputs objectivec.IObject, inputs2 objectivec.IObject, outputs2 objectivec.IObject, name objectivec.IObject, initializer objectivec.IObject) (ETModelDefinition, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initWithTrainingNetworkPath:inferenceInputs:inferenceOutputs:trainingInputs:trainingOutputs:trainingControlVariableName:withInitializer:error:"), path, inputs, outputs, inputs2, outputs2, name, initializer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ETModelDefinition{}, foundation.NSErrorFrom(errorPtr)
	}
	return ETModelDefinitionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/initializer
func (e ETModelDefinition) Initializer() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("initializer"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/inputs
func (e ETModelDefinition) Inputs() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/isTrainingGlobalName
func (e ETModelDefinition) IsTrainingGlobalName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("isTrainingGlobalName"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/layerNames
func (e ETModelDefinition) LayerNames() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("layerNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/modelURL
func (e ETModelDefinition) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/outputs
func (e ETModelDefinition) Outputs() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/trainingInputs
func (e ETModelDefinition) TrainingInputs() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("trainingInputs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Espresso/ETModelDefinition/trainingOutputs
func (e ETModelDefinition) TrainingOutputs() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("trainingOutputs"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

