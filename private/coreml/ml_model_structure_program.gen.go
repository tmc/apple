// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgram] class.
var (
	_MLModelStructureProgramClass     MLModelStructureProgramClass
	_MLModelStructureProgramClassOnce sync.Once
)

func getMLModelStructureProgramClass() MLModelStructureProgramClass {
	_MLModelStructureProgramClassOnce.Do(func() {
		_MLModelStructureProgramClass = MLModelStructureProgramClass{class: objc.GetClass("MLModelStructureProgram")}
	})
	return _MLModelStructureProgramClass
}

// GetMLModelStructureProgramClass returns the class object for MLModelStructureProgram.
func GetMLModelStructureProgramClass() MLModelStructureProgramClass {
	return getMLModelStructureProgramClass()
}

type MLModelStructureProgramClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramClass) Alloc() MLModelStructureProgram {
	rv := objc.Send[MLModelStructureProgram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureProgram.MainFunction]
//   - [MLModelStructureProgram.InitWithFunctions]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram
type MLModelStructureProgram struct {
	objectivec.Object
}

// MLModelStructureProgramFromID constructs a [MLModelStructureProgram] from an objc.ID.
func MLModelStructureProgramFromID(id objc.ID) MLModelStructureProgram {
	return MLModelStructureProgram{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgram implements IMLModelStructureProgram.
var _ IMLModelStructureProgram = MLModelStructureProgram{}

// An interface definition for the [MLModelStructureProgram] class.
//
// # Methods
//
//   - [IMLModelStructureProgram.MainFunction]
//   - [IMLModelStructureProgram.InitWithFunctions]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram
type IMLModelStructureProgram interface {
	objectivec.IObject

	// Topic: Methods

	MainFunction() IMLModelStructureProgramFunction
	InitWithFunctions(functions objectivec.IObject) MLModelStructureProgram
}

// Init initializes the instance.
func (m MLModelStructureProgram) Init() MLModelStructureProgram {
	rv := objc.Send[MLModelStructureProgram](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgram) Autorelease() MLModelStructureProgram {
	rv := objc.Send[MLModelStructureProgram](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgram creates a new MLModelStructureProgram instance.
func NewMLModelStructureProgram() MLModelStructureProgram {
	class := getMLModelStructureProgramClass()
	rv := objc.Send[MLModelStructureProgram](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram/initWithFunctions:
func NewModelStructureProgramWithFunctions(functions objectivec.IObject) MLModelStructureProgram {
	instance := getMLModelStructureProgramClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFunctions:"), functions)
	return MLModelStructureProgramFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram/initWithFunctions:
func (m MLModelStructureProgram) InitWithFunctions(functions objectivec.IObject) MLModelStructureProgram {
	rv := objc.Send[MLModelStructureProgram](m.ID, objc.Sel("initWithFunctions:"), functions)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram/mainFunction
func (m MLModelStructureProgram) MainFunction() IMLModelStructureProgramFunction {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mainFunction"))
	return MLModelStructureProgramFunctionFromID(objc.ID(rv))
}
