// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramOperation] class.
var (
	_MLModelStructureProgramOperationClass     MLModelStructureProgramOperationClass
	_MLModelStructureProgramOperationClassOnce sync.Once
)

func getMLModelStructureProgramOperationClass() MLModelStructureProgramOperationClass {
	_MLModelStructureProgramOperationClassOnce.Do(func() {
		_MLModelStructureProgramOperationClass = MLModelStructureProgramOperationClass{class: objc.GetClass("MLModelStructureProgramOperation")}
	})
	return _MLModelStructureProgramOperationClass
}

// GetMLModelStructureProgramOperationClass returns the class object for MLModelStructureProgramOperation.
func GetMLModelStructureProgramOperationClass() MLModelStructureProgramOperationClass {
	return getMLModelStructureProgramOperationClass()
}

type MLModelStructureProgramOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramOperationClass) Alloc() MLModelStructureProgramOperation {
	rv := objc.Send[MLModelStructureProgramOperation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A class representing an Operation in a Program.
//
// # Accessing the program operation properties
//
//   - [MLModelStructureProgramOperation.Blocks]: Nested blocks for loops and conditionals, e.g., a conditional block will have two entries here.
//   - [MLModelStructureProgramOperation.Inputs]: The arguments to the Operation.
//   - [MLModelStructureProgramOperation.OperatorName]: The name of the operator, e.g., “conv”, “pool”, “softmax”, etc.
//   - [MLModelStructureProgramOperation.Outputs]: The outputs of the Operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation
type MLModelStructureProgramOperation struct {
	objectivec.Object
}

// MLModelStructureProgramOperationFromID constructs a [MLModelStructureProgramOperation] from an objc.ID.
//
// A class representing an Operation in a Program.
func MLModelStructureProgramOperationFromID(id objc.ID) MLModelStructureProgramOperation {
	return MLModelStructureProgramOperation{objectivec.Object{id}}
}
// Ensure MLModelStructureProgramOperation implements IMLModelStructureProgramOperation.
var _ IMLModelStructureProgramOperation = MLModelStructureProgramOperation{}





// An interface definition for the [MLModelStructureProgramOperation] class.
//
// # Accessing the program operation properties
//
//   - [IMLModelStructureProgramOperation.Blocks]: Nested blocks for loops and conditionals, e.g., a conditional block will have two entries here.
//   - [IMLModelStructureProgramOperation.Inputs]: The arguments to the Operation.
//   - [IMLModelStructureProgramOperation.OperatorName]: The name of the operator, e.g., “conv”, “pool”, “softmax”, etc.
//   - [IMLModelStructureProgramOperation.Outputs]: The outputs of the Operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation
type IMLModelStructureProgramOperation interface {
	objectivec.IObject

	// Topic: Accessing the program operation properties

	// Nested blocks for loops and conditionals, e.g., a conditional block will have two entries here.
	Blocks() []MLModelStructureProgramBlock
	// The arguments to the Operation.
	Inputs() foundation.INSDictionary
	// The name of the operator, e.g., “conv”, “pool”, “softmax”, etc.
	OperatorName() string
	// The outputs of the Operation.
	Outputs() []MLModelStructureProgramNamedValueType
}





// Init initializes the instance.
func (m MLModelStructureProgramOperation) Init() MLModelStructureProgramOperation {
	rv := objc.Send[MLModelStructureProgramOperation](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramOperation) Autorelease() MLModelStructureProgramOperation {
	rv := objc.Send[MLModelStructureProgramOperation](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramOperation creates a new MLModelStructureProgramOperation instance.
func NewMLModelStructureProgramOperation() MLModelStructureProgramOperation {
	class := getMLModelStructureProgramOperationClass()
	rv := objc.Send[MLModelStructureProgramOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// Nested blocks for loops and conditionals, e.g., a conditional block will
// have two entries here.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/blocks
func (m MLModelStructureProgramOperation) Blocks() []MLModelStructureProgramBlock {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("blocks"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureProgramBlock {
		return MLModelStructureProgramBlockFromID(id)
	})
}



// The arguments to the Operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/inputs
func (m MLModelStructureProgramOperation) Inputs() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputs"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}



// The name of the operator, e.g., “conv”, “pool”, “softmax”, etc.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/operatorName
func (m MLModelStructureProgramOperation) OperatorName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("operatorName"))
	return foundation.NSStringFromID(rv).String()
}



// The outputs of the Operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramOperation/outputs
func (m MLModelStructureProgramOperation) Outputs() []MLModelStructureProgramNamedValueType {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("outputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureProgramNamedValueType {
		return MLModelStructureProgramNamedValueTypeFromID(id)
	})
}

















