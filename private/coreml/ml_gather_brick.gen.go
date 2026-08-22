// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGatherBrick] class.
var (
	_MLGatherBrickClass     MLGatherBrickClass
	_MLGatherBrickClassOnce sync.Once
)

func getMLGatherBrickClass() MLGatherBrickClass {
	_MLGatherBrickClassOnce.Do(func() {
		_MLGatherBrickClass = MLGatherBrickClass{class: objc.GetClass("MLGatherBrick")}
	})
	return _MLGatherBrickClass
}

// GetMLGatherBrickClass returns the class object for MLGatherBrick.
func GetMLGatherBrickClass() MLGatherBrickClass {
	return getMLGatherBrickClass()
}

type MLGatherBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGatherBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGatherBrickClass) Alloc() MLGatherBrick {
	rv := objc.SendIfResponds[MLGatherBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGatherBrick.Axis]
//   - [MLGatherBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLGatherBrick.HasGPUSupport]
//   - [MLGatherBrick.InputRanks]
//   - [MLGatherBrick.InputShapes]
//   - [MLGatherBrick.OutputRanks]
//   - [MLGatherBrick.OutputShapes]
//   - [MLGatherBrick.SetupForInputShapesWithParameters]
//   - [MLGatherBrick.ShapeInfoNeeded]
//   - [MLGatherBrick.InitWithParameters]
//   - [MLGatherBrick.DebugDescription]
//   - [MLGatherBrick.Description]
//   - [MLGatherBrick.Hash]
//   - [MLGatherBrick.Superclass]
type MLGatherBrick struct {
	objectivec.Object
}

// MLGatherBrickFromID constructs a [MLGatherBrick] from an objc.ID.
func MLGatherBrickFromID(id objc.ID) MLGatherBrick {
	return MLGatherBrick{objectivec.Object{ID: id}}
}

// Ensure MLGatherBrick implements IMLGatherBrick.
var _ IMLGatherBrick = MLGatherBrick{}

// An interface definition for the [MLGatherBrick] class.
//
// # Methods
//
//   - [IMLGatherBrick.Axis]
//   - [IMLGatherBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLGatherBrick.HasGPUSupport]
//   - [IMLGatherBrick.InputRanks]
//   - [IMLGatherBrick.InputShapes]
//   - [IMLGatherBrick.OutputRanks]
//   - [IMLGatherBrick.OutputShapes]
//   - [IMLGatherBrick.SetupForInputShapesWithParameters]
//   - [IMLGatherBrick.ShapeInfoNeeded]
//   - [IMLGatherBrick.InitWithParameters]
//   - [IMLGatherBrick.DebugDescription]
//   - [IMLGatherBrick.Description]
//   - [IMLGatherBrick.Hash]
//   - [IMLGatherBrick.Superclass]
type IMLGatherBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLGatherBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLGatherBrick) Init() MLGatherBrick {
	rv := objc.SendIfResponds[MLGatherBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLGatherBrick) Autorelease() MLGatherBrick {
	rv := objc.SendIfResponds[MLGatherBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGatherBrick creates a new MLGatherBrick instance.
func NewMLGatherBrick() MLGatherBrick {
	class := getMLGatherBrickClass()
	rv := objc.SendIfResponds[MLGatherBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGatherBrickWithParameters(parameters objectivec.IObject) MLGatherBrick {
	instance := getMLGatherBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLGatherBrickFromID(rv)
}

func (m MLGatherBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLGatherBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLGatherBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLGatherBrick) InitWithParameters(parameters objectivec.IObject) MLGatherBrick {
	rv := objc.SendIfResponds[MLGatherBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLGatherBrick) Axis() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLGatherBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGatherBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGatherBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLGatherBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLGatherBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLGatherBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLGatherBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLGatherBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLGatherBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
