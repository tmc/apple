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

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramFunctionClass) Alloc() MLModelStructureProgramFunction {
	rv := objc.Send[MLModelStructureProgramFunction](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing a function in the Program.
//
// # Accessing the program function properties
//
//   - [MLModelStructureProgramFunction.Block]: The active block in the function.
//   - [MLModelStructureProgramFunction.Inputs]: The named inputs to the function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction
type MLModelStructureProgramFunction struct {
	objectivec.Object
}

// MLModelStructureProgramFunctionFromID constructs a [MLModelStructureProgramFunction] from an objc.ID.
//
// A class representing a function in the Program.
func MLModelStructureProgramFunctionFromID(id objc.ID) MLModelStructureProgramFunction {
	return MLModelStructureProgramFunction{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureProgramFunction implements IMLModelStructureProgramFunction.
var _ IMLModelStructureProgramFunction = MLModelStructureProgramFunction{}

// An interface definition for the [MLModelStructureProgramFunction] class.
//
// # Accessing the program function properties
//
//   - [IMLModelStructureProgramFunction.Block]: The active block in the function.
//   - [IMLModelStructureProgramFunction.Inputs]: The named inputs to the function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction
type IMLModelStructureProgramFunction interface {
	objectivec.IObject

	// Topic: Accessing the program function properties

	// The active block in the function.
	Block() IMLModelStructureProgramBlock
	// The named inputs to the function.
	Inputs() []MLModelStructureProgramNamedValueType
}

// Init initializes the instance.
func (m MLModelStructureProgramFunction) Init() MLModelStructureProgramFunction {
	rv := objc.Send[MLModelStructureProgramFunction](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramFunction) Autorelease() MLModelStructureProgramFunction {
	rv := objc.Send[MLModelStructureProgramFunction](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramFunction creates a new MLModelStructureProgramFunction instance.
func NewMLModelStructureProgramFunction() MLModelStructureProgramFunction {
	class := getMLModelStructureProgramFunctionClass()
	rv := objc.Send[MLModelStructureProgramFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The active block in the function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction/block
func (m MLModelStructureProgramFunction) Block() IMLModelStructureProgramBlock {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("block"))
	return MLModelStructureProgramBlockFromID(objc.ID(rv))
}

// The named inputs to the function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramFunction/inputs
func (m MLModelStructureProgramFunction) Inputs() []MLModelStructureProgramNamedValueType {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("inputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureProgramNamedValueType {
		return MLModelStructureProgramNamedValueTypeFromID(id)
	})
}

