// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramClass) Alloc() MLModelStructureProgram {
	rv := objc.Send[MLModelStructureProgram](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing the structure of an ML Program model.
//
// # Accessing the program functions
//
//   - [MLModelStructureProgram.Functions]: The functions in the program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram
type MLModelStructureProgram struct {
	objectivec.Object
}

// MLModelStructureProgramFromID constructs a [MLModelStructureProgram] from an objc.ID.
//
// A class representing the structure of an ML Program model.
func MLModelStructureProgramFromID(id objc.ID) MLModelStructureProgram {
	return MLModelStructureProgram{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureProgram implements IMLModelStructureProgram.
var _ IMLModelStructureProgram = MLModelStructureProgram{}

// An interface definition for the [MLModelStructureProgram] class.
//
// # Accessing the program functions
//
//   - [IMLModelStructureProgram.Functions]: The functions in the program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram
type IMLModelStructureProgram interface {
	objectivec.IObject

	// Topic: Accessing the program functions

	// The functions in the program.
	Functions() foundation.INSDictionary
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

// The functions in the program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgram/functions
func (m MLModelStructureProgram) Functions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

