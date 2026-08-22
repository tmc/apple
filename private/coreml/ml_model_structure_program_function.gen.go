// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramFunction] class.
var (
	_MLModelStructureProgramFunctionClass     MLModelStructureProgramFunctionClass
	_MLModelStructureProgramFunctionClassOnce sync.Once
)

func getMLModelStructureProgramFunctionClass() MLModelStructureProgramFunctionClass {
	_MLModelStructureProgramFunctionClassOnce.Do(func() {
		_MLModelStructureProgramFunctionClass = MLModelStructureProgramFunctionClass{class: objc.GetClass("MLModelStructureProgramFunction")}
	})
	return _MLModelStructureProgramFunctionClass
}

// GetMLModelStructureProgramFunctionClass returns the class object for MLModelStructureProgramFunction.
func GetMLModelStructureProgramFunctionClass() MLModelStructureProgramFunctionClass {
	return getMLModelStructureProgramFunctionClass()
}

type MLModelStructureProgramFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramFunctionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramFunctionClass) Alloc() MLModelStructureProgramFunction {
	rv := objc.SendIfResponds[MLModelStructureProgramFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureProgramFunction.InitWithInputsBlock]
type MLModelStructureProgramFunction struct {
	objectivec.Object
}

// MLModelStructureProgramFunctionFromID constructs a [MLModelStructureProgramFunction] from an objc.ID.
func MLModelStructureProgramFunctionFromID(id objc.ID) MLModelStructureProgramFunction {
	return MLModelStructureProgramFunction{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgramFunction implements IMLModelStructureProgramFunction.
var _ IMLModelStructureProgramFunction = MLModelStructureProgramFunction{}

// An interface definition for the [MLModelStructureProgramFunction] class.
//
// # Methods
//
//   - [IMLModelStructureProgramFunction.InitWithInputsBlock]
type IMLModelStructureProgramFunction interface {
	objectivec.IObject

	// Topic: Methods

	InitWithInputsBlock(inputs objectivec.IObject, block objectivec.IObject) MLModelStructureProgramFunction
}

// Init initializes the instance.
func (m MLModelStructureProgramFunction) Init() MLModelStructureProgramFunction {
	rv := objc.SendIfResponds[MLModelStructureProgramFunction](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramFunction) Autorelease() MLModelStructureProgramFunction {
	rv := objc.SendIfResponds[MLModelStructureProgramFunction](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramFunction creates a new MLModelStructureProgramFunction instance.
func NewMLModelStructureProgramFunction() MLModelStructureProgramFunction {
	class := getMLModelStructureProgramFunctionClass()
	rv := objc.SendIfResponds[MLModelStructureProgramFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelStructureProgramFunctionWithInputsBlock(inputs objectivec.IObject, block objectivec.IObject) MLModelStructureProgramFunction {
	instance := getMLModelStructureProgramFunctionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputs:block:"), inputs, block)
	return MLModelStructureProgramFunctionFromID(rv)
}

func (m MLModelStructureProgramFunction) InitWithInputsBlock(inputs objectivec.IObject, block objectivec.IObject) MLModelStructureProgramFunction {
	rv := objc.SendIfResponds[MLModelStructureProgramFunction](m.ID, objc.Sel("initWithInputs:block:"), inputs, block)
	return rv
}
