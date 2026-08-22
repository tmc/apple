// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProgramEvaluator] class.
var (
	_MLProgramEvaluatorClass     MLProgramEvaluatorClass
	_MLProgramEvaluatorClassOnce sync.Once
)

func getMLProgramEvaluatorClass() MLProgramEvaluatorClass {
	_MLProgramEvaluatorClassOnce.Do(func() {
		_MLProgramEvaluatorClass = MLProgramEvaluatorClass{class: objc.GetClass("MLProgramEvaluator")}
	})
	return _MLProgramEvaluatorClass
}

// GetMLProgramEvaluatorClass returns the class object for MLProgramEvaluator.
func GetMLProgramEvaluatorClass() MLProgramEvaluatorClass {
	return getMLProgramEvaluatorClass()
}

type MLProgramEvaluatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramEvaluatorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramEvaluatorClass) Alloc() MLProgramEvaluator {
	rv := objc.SendIfResponds[MLProgramEvaluator](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProgramEvaluator.EvaluateFunctionArgumentsContextError]
//   - [MLProgramEvaluator.EvaluateFunctionArgumentsContextUpdateContextError]
//   - [MLProgramEvaluator.Model]
//   - [MLProgramEvaluator.NewContextAndReturnError]
//   - [MLProgramEvaluator.PrepareArgumentsFromFeaturesContextForFunctionName]
//   - [MLProgramEvaluator.Program]
//   - [MLProgramEvaluator.SetProgram]
//   - [MLProgramEvaluator.UpdateContextFunctionNameResult]
//   - [MLProgramEvaluator.InitWithProgramError]
type MLProgramEvaluator struct {
	objectivec.Object
}

// MLProgramEvaluatorFromID constructs a [MLProgramEvaluator] from an objc.ID.
func MLProgramEvaluatorFromID(id objc.ID) MLProgramEvaluator {
	return MLProgramEvaluator{objectivec.Object{ID: id}}
}

// Ensure MLProgramEvaluator implements IMLProgramEvaluator.
var _ IMLProgramEvaluator = MLProgramEvaluator{}

// An interface definition for the [MLProgramEvaluator] class.
//
// # Methods
//
//   - [IMLProgramEvaluator.EvaluateFunctionArgumentsContextError]
//   - [IMLProgramEvaluator.EvaluateFunctionArgumentsContextUpdateContextError]
//   - [IMLProgramEvaluator.Model]
//   - [IMLProgramEvaluator.NewContextAndReturnError]
//   - [IMLProgramEvaluator.PrepareArgumentsFromFeaturesContextForFunctionName]
//   - [IMLProgramEvaluator.Program]
//   - [IMLProgramEvaluator.SetProgram]
//   - [IMLProgramEvaluator.UpdateContextFunctionNameResult]
//   - [IMLProgramEvaluator.InitWithProgramError]
type IMLProgramEvaluator interface {
	objectivec.IObject

	// Topic: Methods

	EvaluateFunctionArgumentsContextError(function objectivec.IObject, arguments objectivec.IObject, context objectivec.IObject) (objectivec.IObject, error)
	EvaluateFunctionArgumentsContextUpdateContextError(function objectivec.IObject, arguments objectivec.IObject, context objectivec.IObject, context2 bool) (objectivec.IObject, error)
	Model() unsafe.Pointer
	NewContextAndReturnError() (objectivec.IObject, error)
	PrepareArgumentsFromFeaturesContextForFunctionName(features objectivec.IObject, context objectivec.IObject, name objectivec.IObject) objectivec.IObject
	Program() unsafe.Pointer
	SetProgram(value unsafe.Pointer)
	UpdateContextFunctionNameResult(context objectivec.IObject, name objectivec.IObject, result objectivec.IObject)
	InitWithProgramError(program objectivec.IObject) (MLProgramEvaluator, error)
}

// Init initializes the instance.
func (m MLProgramEvaluator) Init() MLProgramEvaluator {
	rv := objc.SendIfResponds[MLProgramEvaluator](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProgramEvaluator) Autorelease() MLProgramEvaluator {
	rv := objc.SendIfResponds[MLProgramEvaluator](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramEvaluator creates a new MLProgramEvaluator instance.
func NewMLProgramEvaluator() MLProgramEvaluator {
	class := getMLProgramEvaluatorClass()
	rv := objc.SendIfResponds[MLProgramEvaluator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewProgramEvaluatorWithProgramError(program objectivec.IObject) (MLProgramEvaluator, error) {
	var errorPtr objc.ID
	instance := getMLProgramEvaluatorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithProgram:error:"), program, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramEvaluator{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLProgramEvaluator{}, objc.ErrInitFailed
	}
	return MLProgramEvaluatorFromID(rv), nil
}

func (m MLProgramEvaluator) EvaluateFunctionArgumentsContextError(function objectivec.IObject, arguments objectivec.IObject, context objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateFunction:arguments:context:error:"), function, arguments, context, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLProgramEvaluator) EvaluateFunctionArgumentsContextUpdateContextError(function objectivec.IObject, arguments objectivec.IObject, context objectivec.IObject, context2 bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateFunction:arguments:context:updateContext:error:"), function, arguments, context, context2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLProgramEvaluator) NewContextAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("newContextAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLProgramEvaluator) PrepareArgumentsFromFeaturesContextForFunctionName(features objectivec.IObject, context objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("prepareArgumentsFromFeatures:context:forFunctionName:"), features, context, name)
	return objectivec.Object{ID: rv}
}
func (m MLProgramEvaluator) UpdateContextFunctionNameResult(context objectivec.IObject, name objectivec.IObject, result objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateContext:functionName:result:"), context, name, result)
}
func (m MLProgramEvaluator) InitWithProgramError(program objectivec.IObject) (MLProgramEvaluator, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithProgram:error:"), program, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramEvaluator{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramEvaluatorFromID(rv), nil

}

func (m MLProgramEvaluator) Model() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("model"))
	return rv
}
func (m MLProgramEvaluator) Program() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("program"))
	return rv
}
func (m MLProgramEvaluator) SetProgram(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setProgram:"), value)
}
