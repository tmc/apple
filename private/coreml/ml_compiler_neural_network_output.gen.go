// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

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

// # Methods
//
//   - [MLCompilerNeuralNetworkOutput.Network]
//   - [MLCompilerNeuralNetworkOutput.Program]
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
type IMLCompilerNeuralNetworkOutput interface {
	objectivec.IObject

	// Topic: Methods

	Network() unsafe.Pointer
	Program() unsafe.Pointer
}

// Init initializes the instance.
func (m MLCompilerNeuralNetworkOutput) Init() MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCompilerNeuralNetworkOutput) Autorelease() MLCompilerNeuralNetworkOutput {
	rv := objc.Send[MLCompilerNeuralNetworkOutput](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCompilerNeuralNetworkOutput creates a new MLCompilerNeuralNetworkOutput instance.
func NewMLCompilerNeuralNetworkOutput() MLCompilerNeuralNetworkOutput {
	class := getMLCompilerNeuralNetworkOutputClass()
	rv := objc.Send[MLCompilerNeuralNetworkOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCompilerNeuralNetworkOutputWithEspressoNetwork(network unsafe.Pointer) MLCompilerNeuralNetworkOutput {
	instance := getMLCompilerNeuralNetworkOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEspressoNetwork:"), network)
	return MLCompilerNeuralNetworkOutputFromID(rv)
}

func NewCompilerNeuralNetworkOutputWithMILProgram(mILProgram unsafe.Pointer) MLCompilerNeuralNetworkOutput {
	instance := getMLCompilerNeuralNetworkOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMILProgram:"), mILProgram)
	return MLCompilerNeuralNetworkOutputFromID(rv)
}

func (m MLCompilerNeuralNetworkOutput) Network() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("network"))
	return rv
}
func (m MLCompilerNeuralNetworkOutput) Program() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("program"))
	return rv
}
