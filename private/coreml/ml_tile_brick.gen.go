// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLTileBrick] class.
var (
	_MLTileBrickClass     MLTileBrickClass
	_MLTileBrickClassOnce sync.Once
)

func getMLTileBrickClass() MLTileBrickClass {
	_MLTileBrickClassOnce.Do(func() {
		_MLTileBrickClass = MLTileBrickClass{class: objc.GetClass("MLTileBrick")}
	})
	return _MLTileBrickClass
}

// GetMLTileBrickClass returns the class object for MLTileBrick.
func GetMLTileBrickClass() MLTileBrickClass {
	return getMLTileBrickClass()
}

type MLTileBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTileBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTileBrickClass) Alloc() MLTileBrick {
	rv := objc.SendIfResponds[MLTileBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLTileBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLTileBrick.HasGPUSupport]
//   - [MLTileBrick.InputRanks]
//   - [MLTileBrick.InputShapes]
//   - [MLTileBrick.OutputRanks]
//   - [MLTileBrick.OutputShapes]
//   - [MLTileBrick.Reps]
//   - [MLTileBrick.SetupForInputShapesWithParameters]
//   - [MLTileBrick.ShapeInfoNeeded]
//   - [MLTileBrick.InitWithParameters]
//   - [MLTileBrick.DebugDescription]
//   - [MLTileBrick.Description]
//   - [MLTileBrick.Hash]
//   - [MLTileBrick.Superclass]
type MLTileBrick struct {
	objectivec.Object
}

// MLTileBrickFromID constructs a [MLTileBrick] from an objc.ID.
func MLTileBrickFromID(id objc.ID) MLTileBrick {
	return MLTileBrick{objectivec.Object{ID: id}}
}

// Ensure MLTileBrick implements IMLTileBrick.
var _ IMLTileBrick = MLTileBrick{}

// An interface definition for the [MLTileBrick] class.
//
// # Methods
//
//   - [IMLTileBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLTileBrick.HasGPUSupport]
//   - [IMLTileBrick.InputRanks]
//   - [IMLTileBrick.InputShapes]
//   - [IMLTileBrick.OutputRanks]
//   - [IMLTileBrick.OutputShapes]
//   - [IMLTileBrick.Reps]
//   - [IMLTileBrick.SetupForInputShapesWithParameters]
//   - [IMLTileBrick.ShapeInfoNeeded]
//   - [IMLTileBrick.InitWithParameters]
//   - [IMLTileBrick.DebugDescription]
//   - [IMLTileBrick.Description]
//   - [IMLTileBrick.Hash]
//   - [IMLTileBrick.Superclass]
type IMLTileBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	Reps() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	InitWithParameters(parameters objectivec.IObject) MLTileBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLTileBrick) Init() MLTileBrick {
	rv := objc.SendIfResponds[MLTileBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLTileBrick) Autorelease() MLTileBrick {
	rv := objc.SendIfResponds[MLTileBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTileBrick creates a new MLTileBrick instance.
func NewMLTileBrick() MLTileBrick {
	class := getMLTileBrickClass()
	rv := objc.SendIfResponds[MLTileBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTileBrickWithParameters(parameters objectivec.IObject) MLTileBrick {
	instance := getMLTileBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLTileBrickFromID(rv)
}

func (m MLTileBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLTileBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLTileBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLTileBrick) InitWithParameters(parameters objectivec.IObject) MLTileBrick {
	rv := objc.SendIfResponds[MLTileBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLTileBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTileBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTileBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLTileBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTileBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTileBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTileBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTileBrick) Reps() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("reps"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTileBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLTileBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
