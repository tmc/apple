// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructure] class.
var (
	_MLModelStructureClass     MLModelStructureClass
	_MLModelStructureClassOnce sync.Once
)

func getMLModelStructureClass() MLModelStructureClass {
	_MLModelStructureClassOnce.Do(func() {
		_MLModelStructureClass = MLModelStructureClass{class: objc.GetClass("MLModelStructure")}
	})
	return _MLModelStructureClass
}

// GetMLModelStructureClass returns the class object for MLModelStructure.
func GetMLModelStructureClass() MLModelStructureClass {
	return getMLModelStructureClass()
}

type MLModelStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureClass) Alloc() MLModelStructure {
	rv := objc.Send[MLModelStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelStructure.InitWithNeuralNetwork]
//   - [MLModelStructure.InitWithNeuralNetworkProgramPipeline]
//   - [MLModelStructure.InitWithPipeline]
//   - [MLModelStructure.InitWithProgram]
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure
type MLModelStructure struct {
	objectivec.Object
}

// MLModelStructureFromID constructs a [MLModelStructure] from an objc.ID.
func MLModelStructureFromID(id objc.ID) MLModelStructure {
	return MLModelStructure{objectivec.Object{ID: id}}
}
// Ensure MLModelStructure implements IMLModelStructure.
var _ IMLModelStructure = MLModelStructure{}

// An interface definition for the [MLModelStructure] class.
//
// # Methods
//
//   - [IMLModelStructure.InitWithNeuralNetwork]
//   - [IMLModelStructure.InitWithNeuralNetworkProgramPipeline]
//   - [IMLModelStructure.InitWithPipeline]
//   - [IMLModelStructure.InitWithProgram]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure
type IMLModelStructure interface {
	objectivec.IObject

	// Topic: Methods

	InitWithNeuralNetwork(network objectivec.IObject) MLModelStructure
	InitWithNeuralNetworkProgramPipeline(network objectivec.IObject, program objectivec.IObject, pipeline objectivec.IObject) MLModelStructure
	InitWithPipeline(pipeline objectivec.IObject) MLModelStructure
	InitWithProgram(program objectivec.IObject) MLModelStructure
}

// Init initializes the instance.
func (m MLModelStructure) Init() MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructure) Autorelease() MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructure creates a new MLModelStructure instance.
func NewMLModelStructure() MLModelStructure {
	class := getMLModelStructureClass()
	rv := objc.Send[MLModelStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithNeuralNetwork:
func NewModelStructureWithNeuralNetwork(network objectivec.IObject) MLModelStructure {
	instance := getMLModelStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNeuralNetwork:"), network)
	return MLModelStructureFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithNeuralNetwork:program:pipeline:
func NewModelStructureWithNeuralNetworkProgramPipeline(network objectivec.IObject, program objectivec.IObject, pipeline objectivec.IObject) MLModelStructure {
	instance := getMLModelStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNeuralNetwork:program:pipeline:"), network, program, pipeline)
	return MLModelStructureFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithPipeline:
func NewModelStructureWithPipeline(pipeline objectivec.IObject) MLModelStructure {
	instance := getMLModelStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPipeline:"), pipeline)
	return MLModelStructureFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithProgram:
func NewModelStructureWithProgram(program objectivec.IObject) MLModelStructure {
	instance := getMLModelStructureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgram:"), program)
	return MLModelStructureFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithNeuralNetwork:
func (m MLModelStructure) InitWithNeuralNetwork(network objectivec.IObject) MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("initWithNeuralNetwork:"), network)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithNeuralNetwork:program:pipeline:
func (m MLModelStructure) InitWithNeuralNetworkProgramPipeline(network objectivec.IObject, program objectivec.IObject, pipeline objectivec.IObject) MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("initWithNeuralNetwork:program:pipeline:"), network, program, pipeline)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithPipeline:
func (m MLModelStructure) InitWithPipeline(pipeline objectivec.IObject) MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("initWithPipeline:"), pipeline)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure/initWithProgram:
func (m MLModelStructure) InitWithProgram(program objectivec.IObject) MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("initWithProgram:"), program)
	return rv
}

