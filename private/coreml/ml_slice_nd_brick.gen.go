// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSliceNDBrick] class.
var (
	_MLSliceNDBrickClass     MLSliceNDBrickClass
	_MLSliceNDBrickClassOnce sync.Once
)

func getMLSliceNDBrickClass() MLSliceNDBrickClass {
	_MLSliceNDBrickClassOnce.Do(func() {
		_MLSliceNDBrickClass = MLSliceNDBrickClass{class: objc.GetClass("MLSliceNDBrick")}
	})
	return _MLSliceNDBrickClass
}

// GetMLSliceNDBrickClass returns the class object for MLSliceNDBrick.
func GetMLSliceNDBrickClass() MLSliceNDBrickClass {
	return getMLSliceNDBrickClass()
}

type MLSliceNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSliceNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSliceNDBrickClass) Alloc() MLSliceNDBrick {
	rv := objc.SendIfResponds[MLSliceNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSliceNDBrick.Begin_ids]
//   - [MLSliceNDBrick.Begin_masks]
//   - [MLSliceNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSliceNDBrick.End_ids]
//   - [MLSliceNDBrick.End_masks]
//   - [MLSliceNDBrick.HasGPUSupport]
//   - [MLSliceNDBrick.InputRanks]
//   - [MLSliceNDBrick.InputShapes]
//   - [MLSliceNDBrick.OutputRanks]
//   - [MLSliceNDBrick.OutputShapes]
//   - [MLSliceNDBrick.Rank]
//   - [MLSliceNDBrick.SetupForInputShapesWithParameters]
//   - [MLSliceNDBrick.ShapeInfoNeeded]
//   - [MLSliceNDBrick.Strides]
//   - [MLSliceNDBrick.InitWithParameters]
//   - [MLSliceNDBrick.DebugDescription]
//   - [MLSliceNDBrick.Description]
//   - [MLSliceNDBrick.Hash]
//   - [MLSliceNDBrick.Superclass]
type MLSliceNDBrick struct {
	objectivec.Object
}

// MLSliceNDBrickFromID constructs a [MLSliceNDBrick] from an objc.ID.
func MLSliceNDBrickFromID(id objc.ID) MLSliceNDBrick {
	return MLSliceNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLSliceNDBrick implements IMLSliceNDBrick.
var _ IMLSliceNDBrick = MLSliceNDBrick{}

// An interface definition for the [MLSliceNDBrick] class.
//
// # Methods
//
//   - [IMLSliceNDBrick.Begin_ids]
//   - [IMLSliceNDBrick.Begin_masks]
//   - [IMLSliceNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSliceNDBrick.End_ids]
//   - [IMLSliceNDBrick.End_masks]
//   - [IMLSliceNDBrick.HasGPUSupport]
//   - [IMLSliceNDBrick.InputRanks]
//   - [IMLSliceNDBrick.InputShapes]
//   - [IMLSliceNDBrick.OutputRanks]
//   - [IMLSliceNDBrick.OutputShapes]
//   - [IMLSliceNDBrick.Rank]
//   - [IMLSliceNDBrick.SetupForInputShapesWithParameters]
//   - [IMLSliceNDBrick.ShapeInfoNeeded]
//   - [IMLSliceNDBrick.Strides]
//   - [IMLSliceNDBrick.InitWithParameters]
//   - [IMLSliceNDBrick.DebugDescription]
//   - [IMLSliceNDBrick.Description]
//   - [IMLSliceNDBrick.Hash]
//   - [IMLSliceNDBrick.Superclass]
type IMLSliceNDBrick interface {
	objectivec.IObject

	// Topic: Methods

	Begin_ids() foundation.INSArray
	Begin_masks() foundation.INSArray
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	End_ids() foundation.INSArray
	End_masks() foundation.INSArray
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	Rank() int
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	Strides() foundation.INSArray
	InitWithParameters(parameters objectivec.IObject) MLSliceNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSliceNDBrick) Init() MLSliceNDBrick {
	rv := objc.SendIfResponds[MLSliceNDBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSliceNDBrick) Autorelease() MLSliceNDBrick {
	rv := objc.SendIfResponds[MLSliceNDBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSliceNDBrick creates a new MLSliceNDBrick instance.
func NewMLSliceNDBrick() MLSliceNDBrick {
	class := getMLSliceNDBrickClass()
	rv := objc.SendIfResponds[MLSliceNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSliceNDBrickWithParameters(parameters objectivec.IObject) MLSliceNDBrick {
	instance := getMLSliceNDBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSliceNDBrickFromID(rv)
}

func (m MLSliceNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLSliceNDBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLSliceNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLSliceNDBrick) InitWithParameters(parameters objectivec.IObject) MLSliceNDBrick {
	rv := objc.SendIfResponds[MLSliceNDBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLSliceNDBrick) Begin_ids() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("begin_ids"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) Begin_masks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("begin_masks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSliceNDBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSliceNDBrick) End_ids() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("end_ids"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) End_masks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("end_masks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSliceNDBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) Rank() int {
	rv := objc.SendIfResponds[int](m.ID, objc.Sel("rank"))
	return rv
}
func (m MLSliceNDBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLSliceNDBrick) Strides() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("strides"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSliceNDBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
