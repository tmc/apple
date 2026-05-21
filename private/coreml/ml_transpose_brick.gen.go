// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLTransposeBrick] class.
var (
	_MLTransposeBrickClass     MLTransposeBrickClass
	_MLTransposeBrickClassOnce sync.Once
)

func getMLTransposeBrickClass() MLTransposeBrickClass {
	_MLTransposeBrickClassOnce.Do(func() {
		_MLTransposeBrickClass = MLTransposeBrickClass{class: objc.GetClass("MLTransposeBrick")}
	})
	return _MLTransposeBrickClass
}

// GetMLTransposeBrickClass returns the class object for MLTransposeBrick.
func GetMLTransposeBrickClass() MLTransposeBrickClass {
	return getMLTransposeBrickClass()
}

type MLTransposeBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTransposeBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTransposeBrickClass) Alloc() MLTransposeBrick {
	rv := objc.Send[MLTransposeBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLTransposeBrick.Axes]
//   - [MLTransposeBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLTransposeBrick.HasGPUSupport]
//   - [MLTransposeBrick.InputRanks]
//   - [MLTransposeBrick.InputShapes]
//   - [MLTransposeBrick.OutputRanks]
//   - [MLTransposeBrick.OutputShapes]
//   - [MLTransposeBrick.SetupForInputShapesWithParameters]
//   - [MLTransposeBrick.InitWithParameters]
//   - [MLTransposeBrick.DebugDescription]
//   - [MLTransposeBrick.Description]
//   - [MLTransposeBrick.Hash]
//   - [MLTransposeBrick.Superclass]
type MLTransposeBrick struct {
	objectivec.Object
}

// MLTransposeBrickFromID constructs a [MLTransposeBrick] from an objc.ID.
func MLTransposeBrickFromID(id objc.ID) MLTransposeBrick {
	return MLTransposeBrick{objectivec.Object{ID: id}}
}

// Ensure MLTransposeBrick implements IMLTransposeBrick.
var _ IMLTransposeBrick = MLTransposeBrick{}

// An interface definition for the [MLTransposeBrick] class.
//
// # Methods
//
//   - [IMLTransposeBrick.Axes]
//   - [IMLTransposeBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLTransposeBrick.HasGPUSupport]
//   - [IMLTransposeBrick.InputRanks]
//   - [IMLTransposeBrick.InputShapes]
//   - [IMLTransposeBrick.OutputRanks]
//   - [IMLTransposeBrick.OutputShapes]
//   - [IMLTransposeBrick.SetupForInputShapesWithParameters]
//   - [IMLTransposeBrick.InitWithParameters]
//   - [IMLTransposeBrick.DebugDescription]
//   - [IMLTransposeBrick.Description]
//   - [IMLTransposeBrick.Hash]
//   - [IMLTransposeBrick.Superclass]
type IMLTransposeBrick interface {
	objectivec.IObject

	// Topic: Methods

	Axes() foundation.INSArray
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLTransposeBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLTransposeBrick) Init() MLTransposeBrick {
	rv := objc.Send[MLTransposeBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLTransposeBrick) Autorelease() MLTransposeBrick {
	rv := objc.Send[MLTransposeBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTransposeBrick creates a new MLTransposeBrick instance.
func NewMLTransposeBrick() MLTransposeBrick {
	class := getMLTransposeBrickClass()
	rv := objc.Send[MLTransposeBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTransposeBrickWithParameters(parameters objectivec.IObject) MLTransposeBrick {
	instance := getMLTransposeBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLTransposeBrickFromID(rv)
}

func (m MLTransposeBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLTransposeBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLTransposeBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLTransposeBrick) InitWithParameters(parameters objectivec.IObject) MLTransposeBrick {
	rv := objc.Send[MLTransposeBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLTransposeBrick) Axes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("axes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTransposeBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTransposeBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTransposeBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLTransposeBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTransposeBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTransposeBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTransposeBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTransposeBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
