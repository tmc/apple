// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCompilerNeuralNetworkOutput] class.
var (
	_MLCompilerNeuralNetworkOutputClass     MLCompilerNeuralNetworkOutputClass
	_MLCompilerNeuralNetworkOutputClassOnce sync.Once
)

func getMLCompilerNeuralNetworkOutputClass() MLCompilerNeuralNetworkOutputClass {
	_MLCompilerNeuralNetworkOutputClassOnce.Do(func() {
		_MLCompilerNeuralNetworkOutputClass = MLCompilerNeuralNetworkOutputClass{class: objc.GetClass("MLCompilerNeuralNetworkOutput")}
	})
	return _MLCompilerNeuralNetworkOutputClass
}

// GetMLCompilerNeuralNetworkOutputClass returns the class object for MLCompilerNeuralNetworkOutput.
func GetMLCompilerNeuralNetworkOutputClass() MLCompilerNeuralNetworkOutputClass {
	return getMLCompilerNeuralNetworkOutputClass()
}

type MLCompilerNeuralNetworkOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCompilerNeuralNetworkOutputClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCompilerNeuralNetworkOutputClass) Alloc() MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLCompilerNeuralNetworkOutput.Network]
//   - [MLCompilerNeuralNetworkOutput.Program]
//   - [MLCompilerNeuralNetworkOutput.InitWithEspressoNetwork]
//   - [MLCompilerNeuralNetworkOutput.InitWithMILProgram]
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput
type MLCompilerNeuralNetworkOutput struct {
	objectivec.Object
}

// MLCompilerNeuralNetworkOutputFromID constructs a [MLCompilerNeuralNetworkOutput] from an objc.ID.
func MLCompilerNeuralNetworkOutputFromID(id objc.ID) MLCompilerNeuralNetworkOutput {
	return MLCompilerNeuralNetworkOutput{objectivec.Object{ID: id}}
}
// Ensure MLCompilerNeuralNetworkOutput implements IMLCompilerNeuralNetworkOutput.
var _ IMLCompilerNeuralNetworkOutput = MLCompilerNeuralNetworkOutput{}

// An interface definition for the [MLCompilerNeuralNetworkOutput] class.
//
// # Methods
//
//   - [IMLCompilerNeuralNetworkOutput.Network]
//   - [IMLCompilerNeuralNetworkOutput.Program]
//   - [IMLCompilerNeuralNetworkOutput.InitWithEspressoNetwork]
//   - [IMLCompilerNeuralNetworkOutput.InitWithMILProgram]
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput
type IMLCompilerNeuralNetworkOutput interface {
	objectivec.IObject

	// Topic: Methods

	Network() objectivec.IObject
	Program() objectivec.IObject
	InitWithEspressoNetwork(network objectivec.IObject) MLCompilerNeuralNetworkOutput
	InitWithMILProgram(mILProgram objectivec.IObject) MLCompilerNeuralNetworkOutput
}

// Init initializes the instance.
func (c MLCompilerNeuralNetworkOutput) Init() MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCompilerNeuralNetworkOutput) Autorelease() MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerNeuralNetworkOutput creates a new MLCompilerNeuralNetworkOutput instance.
func NewMLCompilerNeuralNetworkOutput() MLCompilerNeuralNetworkOutput {
	class := getMLCompilerNeuralNetworkOutputClass()
	rv := objc.Send[MLCompilerNeuralNetworkOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/initWithEspressoNetwork:
func NewCompilerNeuralNetworkOutputWithEspressoNetwork(network objectivec.IObject) MLCompilerNeuralNetworkOutput {
	instance := getMLCompilerNeuralNetworkOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEspressoNetwork:"), network)
	return MLCompilerNeuralNetworkOutputFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/initWithMILProgram:
func NewCompilerNeuralNetworkOutputWithMILProgram(mILProgram objectivec.IObject) MLCompilerNeuralNetworkOutput {
	instance := getMLCompilerNeuralNetworkOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMILProgram:"), mILProgram)
	return MLCompilerNeuralNetworkOutputFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/initWithEspressoNetwork:
func (c MLCompilerNeuralNetworkOutput) InitWithEspressoNetwork(network objectivec.IObject) MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](c.ID, objc.Sel("initWithEspressoNetwork:"), network)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/initWithMILProgram:
func (c MLCompilerNeuralNetworkOutput) InitWithMILProgram(mILProgram objectivec.IObject) MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](c.ID, objc.Sel("initWithMILProgram:"), mILProgram)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/outputWithEspressoNetwork:
func (_MLCompilerNeuralNetworkOutputClass MLCompilerNeuralNetworkOutputClass) OutputWithEspressoNetwork(network objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerNeuralNetworkOutputClass.class), objc.Sel("outputWithEspressoNetwork:"), network)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/outputWithMILProgram:
func (_MLCompilerNeuralNetworkOutputClass MLCompilerNeuralNetworkOutputClass) OutputWithMILProgram(mILProgram objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLCompilerNeuralNetworkOutputClass.class), objc.Sel("outputWithMILProgram:"), mILProgram)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/network
func (c MLCompilerNeuralNetworkOutput) Network() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("network"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLCompilerNeuralNetworkOutput/program
func (c MLCompilerNeuralNetworkOutput) Program() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("program"))
	return objectivec.Object{ID: rv}
}

