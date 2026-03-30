// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProgramContext] class.
var (
	_MLProgramContextClass     MLProgramContextClass
	_MLProgramContextClassOnce sync.Once
)

func getMLProgramContextClass() MLProgramContextClass {
	_MLProgramContextClassOnce.Do(func() {
		_MLProgramContextClass = MLProgramContextClass{class: objc.GetClass("MLProgramContext")}
	})
	return _MLProgramContextClass
}

// GetMLProgramContextClass returns the class object for MLProgramContext.
func GetMLProgramContextClass() MLProgramContextClass {
	return getMLProgramContextClass()
}

type MLProgramContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramContextClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramContextClass) Alloc() MLProgramContext {
	rv := objc.Send[MLProgramContext](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProgramContext.ExecutionState]
//   - [MLProgramContext.SetExecutionState]
//   - [MLProgramContext.ForwardFunctionLossName]
//   - [MLProgramContext.SetForwardFunctionLossName]
//   - [MLProgramContext.FunctionNameToInputLayersNames]
//   - [MLProgramContext.SetFunctionNameToInputLayersNames]
//   - [MLProgramContext.FunctionNameToOutputLayersNames]
//   - [MLProgramContext.SetFunctionNameToOutputLayersNames]
//   - [MLProgramContext.FunctionNameToStateMap]
//   - [MLProgramContext.SetFunctionNameToStateMap]
//   - [MLProgramContext.TrainFunctionLossName]
//   - [MLProgramContext.SetTrainFunctionLossName]
//   - [MLProgramContext.InitWithExecutionStateFunctionNameToStateMap]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramContext
type MLProgramContext struct {
	objectivec.Object
}

// MLProgramContextFromID constructs a [MLProgramContext] from an objc.ID.
func MLProgramContextFromID(id objc.ID) MLProgramContext {
	return MLProgramContext{objectivec.Object{ID: id}}
}

// Ensure MLProgramContext implements IMLProgramContext.
var _ IMLProgramContext = MLProgramContext{}

// An interface definition for the [MLProgramContext] class.
//
// # Methods
//
//   - [IMLProgramContext.ExecutionState]
//   - [IMLProgramContext.SetExecutionState]
//   - [IMLProgramContext.ForwardFunctionLossName]
//   - [IMLProgramContext.SetForwardFunctionLossName]
//   - [IMLProgramContext.FunctionNameToInputLayersNames]
//   - [IMLProgramContext.SetFunctionNameToInputLayersNames]
//   - [IMLProgramContext.FunctionNameToOutputLayersNames]
//   - [IMLProgramContext.SetFunctionNameToOutputLayersNames]
//   - [IMLProgramContext.FunctionNameToStateMap]
//   - [IMLProgramContext.SetFunctionNameToStateMap]
//   - [IMLProgramContext.TrainFunctionLossName]
//   - [IMLProgramContext.SetTrainFunctionLossName]
//   - [IMLProgramContext.InitWithExecutionStateFunctionNameToStateMap]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramContext
type IMLProgramContext interface {
	objectivec.IObject

	// Topic: Methods

	ExecutionState() objectivec.IObject
	SetExecutionState(value objectivec.IObject)
	ForwardFunctionLossName() string
	SetForwardFunctionLossName(value string)
	FunctionNameToInputLayersNames() foundation.INSDictionary
	SetFunctionNameToInputLayersNames(value foundation.INSDictionary)
	FunctionNameToOutputLayersNames() foundation.INSDictionary
	SetFunctionNameToOutputLayersNames(value foundation.INSDictionary)
	FunctionNameToStateMap() foundation.INSDictionary
	SetFunctionNameToStateMap(value foundation.INSDictionary)
	TrainFunctionLossName() string
	SetTrainFunctionLossName(value string)
	InitWithExecutionStateFunctionNameToStateMap(state objectivec.IObject, map_ objectivec.IObject) MLProgramContext
}

// Init initializes the instance.
func (p MLProgramContext) Init() MLProgramContext {
	rv := objc.Send[MLProgramContext](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProgramContext) Autorelease() MLProgramContext {
	rv := objc.Send[MLProgramContext](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramContext creates a new MLProgramContext instance.
func NewMLProgramContext() MLProgramContext {
	class := getMLProgramContextClass()
	rv := objc.Send[MLProgramContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/initWithExecutionState:functionNameToStateMap:
func NewProgramContextWithExecutionStateFunctionNameToStateMap(state objectivec.IObject, map_ objectivec.IObject) MLProgramContext {
	instance := getMLProgramContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExecutionState:functionNameToStateMap:"), state, map_)
	return MLProgramContextFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/initWithExecutionState:functionNameToStateMap:
func (p MLProgramContext) InitWithExecutionStateFunctionNameToStateMap(state objectivec.IObject, map_ objectivec.IObject) MLProgramContext {
	rv := objc.Send[MLProgramContext](p.ID, objc.Sel("initWithExecutionState:functionNameToStateMap:"), state, map_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/executionState
func (p MLProgramContext) ExecutionState() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("executionState"))
	return objectivec.Object{ID: rv}
}
func (p MLProgramContext) SetExecutionState(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setExecutionState:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/forwardFunctionLossName
func (p MLProgramContext) ForwardFunctionLossName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("forwardFunctionLossName"))
	return foundation.NSStringFromID(rv).String()
}
func (p MLProgramContext) SetForwardFunctionLossName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setForwardFunctionLossName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/functionNameToInputLayersNames
func (p MLProgramContext) FunctionNameToInputLayersNames() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("functionNameToInputLayersNames"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLProgramContext) SetFunctionNameToInputLayersNames(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setFunctionNameToInputLayersNames:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/functionNameToOutputLayersNames
func (p MLProgramContext) FunctionNameToOutputLayersNames() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("functionNameToOutputLayersNames"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLProgramContext) SetFunctionNameToOutputLayersNames(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setFunctionNameToOutputLayersNames:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/functionNameToStateMap
func (p MLProgramContext) FunctionNameToStateMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("functionNameToStateMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p MLProgramContext) SetFunctionNameToStateMap(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setFunctionNameToStateMap:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContext/trainFunctionLossName
func (p MLProgramContext) TrainFunctionLossName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("trainFunctionLossName"))
	return foundation.NSStringFromID(rv).String()
}
func (p MLProgramContext) SetTrainFunctionLossName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setTrainFunctionLossName:"), objc.String(value))
}
