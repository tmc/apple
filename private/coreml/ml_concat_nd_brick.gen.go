// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLConcatNDBrick] class.
var (
	_MLConcatNDBrickClass     MLConcatNDBrickClass
	_MLConcatNDBrickClassOnce sync.Once
)

func getMLConcatNDBrickClass() MLConcatNDBrickClass {
	_MLConcatNDBrickClassOnce.Do(func() {
		_MLConcatNDBrickClass = MLConcatNDBrickClass{class: objc.GetClass("MLConcatNDBrick")}
	})
	return _MLConcatNDBrickClass
}

// GetMLConcatNDBrickClass returns the class object for MLConcatNDBrick.
func GetMLConcatNDBrickClass() MLConcatNDBrickClass {
	return getMLConcatNDBrickClass()
}

type MLConcatNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLConcatNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLConcatNDBrickClass) Alloc() MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLConcatNDBrick.Axis]
//   - [MLConcatNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLConcatNDBrick.HasGPUSupport]
//   - [MLConcatNDBrick.InputRanks]
//   - [MLConcatNDBrick.InputShapes]
//   - [MLConcatNDBrick.OutputRanks]
//   - [MLConcatNDBrick.OutputShapes]
//   - [MLConcatNDBrick.SetupForInputShapesWithParameters]
//   - [MLConcatNDBrick.ShapeInfoNeeded]
//   - [MLConcatNDBrick.InitWithParameters]
//   - [MLConcatNDBrick.DebugDescription]
//   - [MLConcatNDBrick.Description]
//   - [MLConcatNDBrick.Hash]
//   - [MLConcatNDBrick.Superclass]
type MLConcatNDBrick struct {
	objectivec.Object
}

// MLConcatNDBrickFromID constructs a [MLConcatNDBrick] from an objc.ID.
func MLConcatNDBrickFromID(id objc.ID) MLConcatNDBrick {
	return MLConcatNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLConcatNDBrick implements IMLConcatNDBrick.
var _ IMLConcatNDBrick = MLConcatNDBrick{}

// An interface definition for the [MLConcatNDBrick] class.
//
// # Methods
//
//   - [IMLConcatNDBrick.Axis]
//   - [IMLConcatNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLConcatNDBrick.HasGPUSupport]
//   - [IMLConcatNDBrick.InputRanks]
//   - [IMLConcatNDBrick.InputShapes]
//   - [IMLConcatNDBrick.OutputRanks]
//   - [IMLConcatNDBrick.OutputShapes]
//   - [IMLConcatNDBrick.SetupForInputShapesWithParameters]
//   - [IMLConcatNDBrick.ShapeInfoNeeded]
//   - [IMLConcatNDBrick.InitWithParameters]
//   - [IMLConcatNDBrick.DebugDescription]
//   - [IMLConcatNDBrick.Description]
//   - [IMLConcatNDBrick.Hash]
//   - [IMLConcatNDBrick.Superclass]
type IMLConcatNDBrick interface {
	objectivec.IObject

	// Topic: Methods

	Axis() foundation.NSNumber
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	InitWithParameters(parameters objectivec.IObject) MLConcatNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLConcatNDBrick) Init() MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLConcatNDBrick) Autorelease() MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLConcatNDBrick creates a new MLConcatNDBrick instance.
func NewMLConcatNDBrick() MLConcatNDBrick {
	class := getMLConcatNDBrickClass()
	rv := objc.Send[MLConcatNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewConcatNDBrickWithParameters(parameters objectivec.IObject) MLConcatNDBrick {
	instance := getMLConcatNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLConcatNDBrickFromID(rv)
}

func (m MLConcatNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLConcatNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLConcatNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLConcatNDBrick) InitWithParameters(parameters objectivec.IObject) MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLConcatNDBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLConcatNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLConcatNDBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLConcatNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLConcatNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLConcatNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLConcatNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLConcatNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLConcatNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLConcatNDBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
