// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLStackNDBrick] class.
var (
	_MLStackNDBrickClass     MLStackNDBrickClass
	_MLStackNDBrickClassOnce sync.Once
)

func getMLStackNDBrickClass() MLStackNDBrickClass {
	_MLStackNDBrickClassOnce.Do(func() {
		_MLStackNDBrickClass = MLStackNDBrickClass{class: objc.GetClass("MLStackNDBrick")}
	})
	return _MLStackNDBrickClass
}

// GetMLStackNDBrickClass returns the class object for MLStackNDBrick.
func GetMLStackNDBrickClass() MLStackNDBrickClass {
	return getMLStackNDBrickClass()
}

type MLStackNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLStackNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLStackNDBrickClass) Alloc() MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLStackNDBrick.Axis]
//   - [MLStackNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLStackNDBrick.HasGPUSupport]
//   - [MLStackNDBrick.InputRanks]
//   - [MLStackNDBrick.InputShapes]
//   - [MLStackNDBrick.OutputRanks]
//   - [MLStackNDBrick.OutputShapes]
//   - [MLStackNDBrick.SetupForInputShapesWithParameters]
//   - [MLStackNDBrick.ShapeInfoNeeded]
//   - [MLStackNDBrick.InitWithParameters]
//   - [MLStackNDBrick.DebugDescription]
//   - [MLStackNDBrick.Description]
//   - [MLStackNDBrick.Hash]
//   - [MLStackNDBrick.Superclass]
type MLStackNDBrick struct {
	objectivec.Object
}

// MLStackNDBrickFromID constructs a [MLStackNDBrick] from an objc.ID.
func MLStackNDBrickFromID(id objc.ID) MLStackNDBrick {
	return MLStackNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLStackNDBrick implements IMLStackNDBrick.
var _ IMLStackNDBrick = MLStackNDBrick{}

// An interface definition for the [MLStackNDBrick] class.
//
// # Methods
//
//   - [IMLStackNDBrick.Axis]
//   - [IMLStackNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLStackNDBrick.HasGPUSupport]
//   - [IMLStackNDBrick.InputRanks]
//   - [IMLStackNDBrick.InputShapes]
//   - [IMLStackNDBrick.OutputRanks]
//   - [IMLStackNDBrick.OutputShapes]
//   - [IMLStackNDBrick.SetupForInputShapesWithParameters]
//   - [IMLStackNDBrick.ShapeInfoNeeded]
//   - [IMLStackNDBrick.InitWithParameters]
//   - [IMLStackNDBrick.DebugDescription]
//   - [IMLStackNDBrick.Description]
//   - [IMLStackNDBrick.Hash]
//   - [IMLStackNDBrick.Superclass]
type IMLStackNDBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLStackNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLStackNDBrick) Init() MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLStackNDBrick) Autorelease() MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLStackNDBrick creates a new MLStackNDBrick instance.
func NewMLStackNDBrick() MLStackNDBrick {
	class := getMLStackNDBrickClass()
	rv := objc.Send[MLStackNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewStackNDBrickWithParameters(parameters objectivec.IObject) MLStackNDBrick {
	instance := getMLStackNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLStackNDBrickFromID(rv)
}

func (m MLStackNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLStackNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLStackNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLStackNDBrick) InitWithParameters(parameters objectivec.IObject) MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLStackNDBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLStackNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLStackNDBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLStackNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLStackNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLStackNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLStackNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLStackNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLStackNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLStackNDBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
