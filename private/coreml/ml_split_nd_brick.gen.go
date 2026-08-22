// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSplitNDBrick] class.
var (
	_MLSplitNDBrickClass     MLSplitNDBrickClass
	_MLSplitNDBrickClassOnce sync.Once
)

func getMLSplitNDBrickClass() MLSplitNDBrickClass {
	_MLSplitNDBrickClassOnce.Do(func() {
		_MLSplitNDBrickClass = MLSplitNDBrickClass{class: objc.GetClass("MLSplitNDBrick")}
	})
	return _MLSplitNDBrickClass
}

// GetMLSplitNDBrickClass returns the class object for MLSplitNDBrick.
func GetMLSplitNDBrickClass() MLSplitNDBrickClass {
	return getMLSplitNDBrickClass()
}

type MLSplitNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSplitNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSplitNDBrickClass) Alloc() MLSplitNDBrick {
	rv := objc.SendIfResponds[MLSplitNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSplitNDBrick.Axis]
//   - [MLSplitNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSplitNDBrick.HasGPUSupport]
//   - [MLSplitNDBrick.InputRanks]
//   - [MLSplitNDBrick.InputShapes]
//   - [MLSplitNDBrick.NumSplits]
//   - [MLSplitNDBrick.OutputRanks]
//   - [MLSplitNDBrick.OutputShapes]
//   - [MLSplitNDBrick.SetupForInputShapesWithParameters]
//   - [MLSplitNDBrick.ShapeInfoNeeded]
//   - [MLSplitNDBrick.SplitSizes]
//   - [MLSplitNDBrick.InitWithParameters]
//   - [MLSplitNDBrick.DebugDescription]
//   - [MLSplitNDBrick.Description]
//   - [MLSplitNDBrick.Hash]
//   - [MLSplitNDBrick.Superclass]
type MLSplitNDBrick struct {
	objectivec.Object
}

// MLSplitNDBrickFromID constructs a [MLSplitNDBrick] from an objc.ID.
func MLSplitNDBrickFromID(id objc.ID) MLSplitNDBrick {
	return MLSplitNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLSplitNDBrick implements IMLSplitNDBrick.
var _ IMLSplitNDBrick = MLSplitNDBrick{}

// An interface definition for the [MLSplitNDBrick] class.
//
// # Methods
//
//   - [IMLSplitNDBrick.Axis]
//   - [IMLSplitNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSplitNDBrick.HasGPUSupport]
//   - [IMLSplitNDBrick.InputRanks]
//   - [IMLSplitNDBrick.InputShapes]
//   - [IMLSplitNDBrick.NumSplits]
//   - [IMLSplitNDBrick.OutputRanks]
//   - [IMLSplitNDBrick.OutputShapes]
//   - [IMLSplitNDBrick.SetupForInputShapesWithParameters]
//   - [IMLSplitNDBrick.ShapeInfoNeeded]
//   - [IMLSplitNDBrick.SplitSizes]
//   - [IMLSplitNDBrick.InitWithParameters]
//   - [IMLSplitNDBrick.DebugDescription]
//   - [IMLSplitNDBrick.Description]
//   - [IMLSplitNDBrick.Hash]
//   - [IMLSplitNDBrick.Superclass]
type IMLSplitNDBrick interface {
	objectivec.IObject

	// Topic: Methods

	Axis() foundation.NSNumber
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	NumSplits() foundation.NSNumber
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	SplitSizes() foundation.INSArray
	InitWithParameters(parameters objectivec.IObject) MLSplitNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSplitNDBrick) Init() MLSplitNDBrick {
	rv := objc.SendIfResponds[MLSplitNDBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSplitNDBrick) Autorelease() MLSplitNDBrick {
	rv := objc.SendIfResponds[MLSplitNDBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSplitNDBrick creates a new MLSplitNDBrick instance.
func NewMLSplitNDBrick() MLSplitNDBrick {
	class := getMLSplitNDBrickClass()
	rv := objc.SendIfResponds[MLSplitNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSplitNDBrickWithParameters(parameters objectivec.IObject) MLSplitNDBrick {
	instance := getMLSplitNDBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSplitNDBrickFromID(rv)
}

func (m MLSplitNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLSplitNDBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLSplitNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLSplitNDBrick) InitWithParameters(parameters objectivec.IObject) MLSplitNDBrick {
	rv := objc.SendIfResponds[MLSplitNDBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLSplitNDBrick) Axis() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSplitNDBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSplitNDBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSplitNDBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) NumSplits() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("numSplits"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLSplitNDBrick) SplitSizes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("splitSizes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSplitNDBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
