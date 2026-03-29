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

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramBlockClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramBlockClass) Alloc() MLModelStructureProgramBlock {
	rv := objc.Send[MLModelStructureProgramBlock](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelStructureProgramBlock.InitWithInputsOutputNamesOperations]
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock
type MLModelStructureProgramBlock struct {
	objectivec.Object
}

// MLModelStructureProgramBlockFromID constructs a [MLModelStructureProgramBlock] from an objc.ID.
func MLModelStructureProgramBlockFromID(id objc.ID) MLModelStructureProgramBlock {
	return MLModelStructureProgramBlock{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureProgramBlock implements IMLModelStructureProgramBlock.
var _ IMLModelStructureProgramBlock = MLModelStructureProgramBlock{}

// An interface definition for the [MLModelStructureProgramBlock] class.
//
// # Methods
//
//   - [IMLModelStructureProgramBlock.InitWithInputsOutputNamesOperations]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock
type IMLModelStructureProgramBlock interface {
	objectivec.IObject

	// Topic: Methods

	InitWithInputsOutputNamesOperations(inputs objectivec.IObject, names objectivec.IObject, operations objectivec.IObject) MLModelStructureProgramBlock
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

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/initWithInputs:outputNames:operations:
func NewModelStructureProgramBlockWithInputsOutputNamesOperations(inputs objectivec.IObject, names objectivec.IObject, operations objectivec.IObject) MLModelStructureProgramBlock {
	instance := getMLModelStructureProgramBlockClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputs:outputNames:operations:"), inputs, names, operations)
	return MLModelStructureProgramBlockFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBlock/initWithInputs:outputNames:operations:
func (m MLModelStructureProgramBlock) InitWithInputsOutputNamesOperations(inputs objectivec.IObject, names objectivec.IObject, operations objectivec.IObject) MLModelStructureProgramBlock {
	rv := objc.Send[MLModelStructureProgramBlock](m.ID, objc.Sel("initWithInputs:outputNames:operations:"), inputs, names, operations)
	return rv
}

