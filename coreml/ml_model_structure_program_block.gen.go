// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramBlock] class.
var (
	_MLModelStructureProgramBlockClass     MLModelStructureProgramBlockClass
	_MLModelStructureProgramBlockClassOnce sync.Once
)

func getMLModelStructureProgramBlockClass() MLModelStructureProgramBlockClass {
	_MLModelStructureProgramBlockClassOnce.Do(func() {
		_MLModelStructureProgramBlockClass = MLModelStructureProgramBlockClass{class: objc.GetClass("MLModelStructureProgramBlock")}
	})
	return _MLModelStructureProgramBlockClass
}

// GetMLModelStructureProgramBlockClass returns the class object for MLModelStructureProgramBlock.
func GetMLModelStructureProgramBlockClass() MLModelStructureProgramBlockClass {
	return getMLModelStructureProgramBlockClass()
}

type MLModelStructureProgramBlockClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramBlockClass) Alloc() MLModelStructureProgramBlock {
	rv := objc.Send[MLModelStructureProgramBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing a block in the Program.
//
// # Accessing the program block properties
//
//   - [MLModelStructureProgramBlock.Inputs]: The named inputs to the block.
//   - [MLModelStructureProgramBlock.Operations]: The list of topologically sorted operations in the block.
//   - [MLModelStructureProgramBlock.OutputNames]: The output names.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock
type MLModelStructureProgramBlock struct {
	objectivec.Object
}

// MLModelStructureProgramBlockFromID constructs a [MLModelStructureProgramBlock] from an objc.ID.
//
// A class representing a block in the Program.
func MLModelStructureProgramBlockFromID(id objc.ID) MLModelStructureProgramBlock {
	return MLModelStructureProgramBlock{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureProgramBlock implements IMLModelStructureProgramBlock.
var _ IMLModelStructureProgramBlock = MLModelStructureProgramBlock{}

// An interface definition for the [MLModelStructureProgramBlock] class.
//
// # Accessing the program block properties
//
//   - [IMLModelStructureProgramBlock.Inputs]: The named inputs to the block.
//   - [IMLModelStructureProgramBlock.Operations]: The list of topologically sorted operations in the block.
//   - [IMLModelStructureProgramBlock.OutputNames]: The output names.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock
type IMLModelStructureProgramBlock interface {
	objectivec.IObject

	// Topic: Accessing the program block properties

	// The named inputs to the block.
	Inputs() []MLModelStructureProgramNamedValueType
	// The list of topologically sorted operations in the block.
	Operations() []MLModelStructureProgramOperation
	// The output names.
	OutputNames() []string
}

// Init initializes the instance.
func (m MLModelStructureProgramBlock) Init() MLModelStructureProgramBlock {
	rv := objc.Send[MLModelStructureProgramBlock](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramBlock) Autorelease() MLModelStructureProgramBlock {
	rv := objc.Send[MLModelStructureProgramBlock](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramBlock creates a new MLModelStructureProgramBlock instance.
func NewMLModelStructureProgramBlock() MLModelStructureProgramBlock {
	class := getMLModelStructureProgramBlockClass()
	rv := objc.Send[MLModelStructureProgramBlock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The named inputs to the block.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/inputs
func (m MLModelStructureProgramBlock) Inputs() []MLModelStructureProgramNamedValueType {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("inputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureProgramNamedValueType {
		return MLModelStructureProgramNamedValueTypeFromID(id)
	})
}
// The list of topologically sorted operations in the block.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/operations
func (m MLModelStructureProgramBlock) Operations() []MLModelStructureProgramOperation {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("operations"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureProgramOperation {
		return MLModelStructureProgramOperationFromID(id)
	})
}
// The output names.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/outputNames
func (m MLModelStructureProgramBlock) OutputNames() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("outputNames"))
	return objc.ConvertSliceToStrings(rv)
}

