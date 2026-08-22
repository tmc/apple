// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSlidingWindowsBrick] class.
var (
	_MLSlidingWindowsBrickClass     MLSlidingWindowsBrickClass
	_MLSlidingWindowsBrickClassOnce sync.Once
)

func getMLSlidingWindowsBrickClass() MLSlidingWindowsBrickClass {
	_MLSlidingWindowsBrickClassOnce.Do(func() {
		_MLSlidingWindowsBrickClass = MLSlidingWindowsBrickClass{class: objc.GetClass("MLSlidingWindowsBrick")}
	})
	return _MLSlidingWindowsBrickClass
}

// GetMLSlidingWindowsBrickClass returns the class object for MLSlidingWindowsBrick.
func GetMLSlidingWindowsBrickClass() MLSlidingWindowsBrickClass {
	return getMLSlidingWindowsBrickClass()
}

type MLSlidingWindowsBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSlidingWindowsBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSlidingWindowsBrickClass) Alloc() MLSlidingWindowsBrick {
	rv := objc.SendIfResponds[MLSlidingWindowsBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSlidingWindowsBrick.Axis]
//   - [MLSlidingWindowsBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSlidingWindowsBrick.HasGPUSupport]
//   - [MLSlidingWindowsBrick.InputRanks]
//   - [MLSlidingWindowsBrick.InputShapes]
//   - [MLSlidingWindowsBrick.OutputRanks]
//   - [MLSlidingWindowsBrick.OutputShapes]
//   - [MLSlidingWindowsBrick.SetupForInputShapesWithParameters]
//   - [MLSlidingWindowsBrick.ShapeInfoNeeded]
//   - [MLSlidingWindowsBrick.Size]
//   - [MLSlidingWindowsBrick.Step]
//   - [MLSlidingWindowsBrick.InitWithParameters]
//   - [MLSlidingWindowsBrick.DebugDescription]
//   - [MLSlidingWindowsBrick.Description]
//   - [MLSlidingWindowsBrick.Hash]
//   - [MLSlidingWindowsBrick.Superclass]
type MLSlidingWindowsBrick struct {
	objectivec.Object
}

// MLSlidingWindowsBrickFromID constructs a [MLSlidingWindowsBrick] from an objc.ID.
func MLSlidingWindowsBrickFromID(id objc.ID) MLSlidingWindowsBrick {
	return MLSlidingWindowsBrick{objectivec.Object{ID: id}}
}

// Ensure MLSlidingWindowsBrick implements IMLSlidingWindowsBrick.
var _ IMLSlidingWindowsBrick = MLSlidingWindowsBrick{}

// An interface definition for the [MLSlidingWindowsBrick] class.
//
// # Methods
//
//   - [IMLSlidingWindowsBrick.Axis]
//   - [IMLSlidingWindowsBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSlidingWindowsBrick.HasGPUSupport]
//   - [IMLSlidingWindowsBrick.InputRanks]
//   - [IMLSlidingWindowsBrick.InputShapes]
//   - [IMLSlidingWindowsBrick.OutputRanks]
//   - [IMLSlidingWindowsBrick.OutputShapes]
//   - [IMLSlidingWindowsBrick.SetupForInputShapesWithParameters]
//   - [IMLSlidingWindowsBrick.ShapeInfoNeeded]
//   - [IMLSlidingWindowsBrick.Size]
//   - [IMLSlidingWindowsBrick.Step]
//   - [IMLSlidingWindowsBrick.InitWithParameters]
//   - [IMLSlidingWindowsBrick.DebugDescription]
//   - [IMLSlidingWindowsBrick.Description]
//   - [IMLSlidingWindowsBrick.Hash]
//   - [IMLSlidingWindowsBrick.Superclass]
type IMLSlidingWindowsBrick interface {
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
	Size() foundation.NSNumber
	Step() foundation.NSNumber
	InitWithParameters(parameters objectivec.IObject) MLSlidingWindowsBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSlidingWindowsBrick) Init() MLSlidingWindowsBrick {
	rv := objc.SendIfResponds[MLSlidingWindowsBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSlidingWindowsBrick) Autorelease() MLSlidingWindowsBrick {
	rv := objc.SendIfResponds[MLSlidingWindowsBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSlidingWindowsBrick creates a new MLSlidingWindowsBrick instance.
func NewMLSlidingWindowsBrick() MLSlidingWindowsBrick {
	class := getMLSlidingWindowsBrickClass()
	rv := objc.SendIfResponds[MLSlidingWindowsBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSlidingWindowsBrickWithParameters(parameters objectivec.IObject) MLSlidingWindowsBrick {
	instance := getMLSlidingWindowsBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSlidingWindowsBrickFromID(rv)
}

func (m MLSlidingWindowsBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLSlidingWindowsBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLSlidingWindowsBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLSlidingWindowsBrick) InitWithParameters(parameters objectivec.IObject) MLSlidingWindowsBrick {
	rv := objc.SendIfResponds[MLSlidingWindowsBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLSlidingWindowsBrick) Axis() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSlidingWindowsBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSlidingWindowsBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSlidingWindowsBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLSlidingWindowsBrick) Size() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("size"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) Step() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("step"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLSlidingWindowsBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
