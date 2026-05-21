// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

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
type IMLProgramContext interface {
	objectivec.IObject

	// Topic: Methods

	ExecutionState() unsafe.Pointer
	SetExecutionState(value unsafe.Pointer)
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
func (m MLProgramContext) Init() MLProgramContext {
	rv := objc.Send[MLProgramContext](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProgramContext) Autorelease() MLProgramContext {
	rv := objc.Send[MLProgramContext](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramContext creates a new MLProgramContext instance.
func NewMLProgramContext() MLProgramContext {
	class := getMLProgramContextClass()
	rv := objc.Send[MLProgramContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewProgramContextWithExecutionStateFunctionNameToStateMap(state objectivec.IObject, map_ objectivec.IObject) MLProgramContext {
	instance := getMLProgramContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExecutionState:functionNameToStateMap:"), state, map_)
	return MLProgramContextFromID(rv)
}

func (m MLProgramContext) InitWithExecutionStateFunctionNameToStateMap(state objectivec.IObject, map_ objectivec.IObject) MLProgramContext {
	rv := objc.Send[MLProgramContext](m.ID, objc.Sel("initWithExecutionState:functionNameToStateMap:"), state, map_)
	return rv
}

func (m MLProgramContext) ExecutionState() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("executionState"))
	return rv
}
func (m MLProgramContext) SetExecutionState(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setExecutionState:"), value)
}
func (m MLProgramContext) ForwardFunctionLossName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("forwardFunctionLossName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLProgramContext) SetForwardFunctionLossName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setForwardFunctionLossName:"), objc.String(value))
}
func (m MLProgramContext) FunctionNameToInputLayersNames() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionNameToInputLayersNames"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLProgramContext) SetFunctionNameToInputLayersNames(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionNameToInputLayersNames:"), value)
}
func (m MLProgramContext) FunctionNameToOutputLayersNames() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionNameToOutputLayersNames"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLProgramContext) SetFunctionNameToOutputLayersNames(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionNameToOutputLayersNames:"), value)
}
func (m MLProgramContext) FunctionNameToStateMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionNameToStateMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLProgramContext) SetFunctionNameToStateMap(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionNameToStateMap:"), value)
}
func (m MLProgramContext) TrainFunctionLossName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trainFunctionLossName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLProgramContext) SetTrainFunctionLossName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setTrainFunctionLossName:"), objc.String(value))
}
