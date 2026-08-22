// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiplyBroadcastableBrick] class.
var (
	_MLMultiplyBroadcastableBrickClass     MLMultiplyBroadcastableBrickClass
	_MLMultiplyBroadcastableBrickClassOnce sync.Once
)

func getMLMultiplyBroadcastableBrickClass() MLMultiplyBroadcastableBrickClass {
	_MLMultiplyBroadcastableBrickClassOnce.Do(func() {
		_MLMultiplyBroadcastableBrickClass = MLMultiplyBroadcastableBrickClass{class: objc.GetClass("MLMultiplyBroadcastableBrick")}
	})
	return _MLMultiplyBroadcastableBrickClass
}

// GetMLMultiplyBroadcastableBrickClass returns the class object for MLMultiplyBroadcastableBrick.
func GetMLMultiplyBroadcastableBrickClass() MLMultiplyBroadcastableBrickClass {
	return getMLMultiplyBroadcastableBrickClass()
}

type MLMultiplyBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiplyBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiplyBroadcastableBrickClass) Alloc() MLMultiplyBroadcastableBrick {
	rv := objc.SendIfResponds[MLMultiplyBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiplyBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLMultiplyBroadcastableBrick.HasGPUSupport]
//   - [MLMultiplyBroadcastableBrick.InputRanks]
//   - [MLMultiplyBroadcastableBrick.InputShapes]
//   - [MLMultiplyBroadcastableBrick.OutputRanks]
//   - [MLMultiplyBroadcastableBrick.OutputShapes]
//   - [MLMultiplyBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLMultiplyBroadcastableBrick.ShapeInfoNeeded]
//   - [MLMultiplyBroadcastableBrick.InitWithParameters]
//   - [MLMultiplyBroadcastableBrick.DebugDescription]
//   - [MLMultiplyBroadcastableBrick.Description]
//   - [MLMultiplyBroadcastableBrick.Hash]
//   - [MLMultiplyBroadcastableBrick.Superclass]
type MLMultiplyBroadcastableBrick struct {
	objectivec.Object
}

// MLMultiplyBroadcastableBrickFromID constructs a [MLMultiplyBroadcastableBrick] from an objc.ID.
func MLMultiplyBroadcastableBrickFromID(id objc.ID) MLMultiplyBroadcastableBrick {
	return MLMultiplyBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLMultiplyBroadcastableBrick implements IMLMultiplyBroadcastableBrick.
var _ IMLMultiplyBroadcastableBrick = MLMultiplyBroadcastableBrick{}

// An interface definition for the [MLMultiplyBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLMultiplyBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLMultiplyBroadcastableBrick.HasGPUSupport]
//   - [IMLMultiplyBroadcastableBrick.InputRanks]
//   - [IMLMultiplyBroadcastableBrick.InputShapes]
//   - [IMLMultiplyBroadcastableBrick.OutputRanks]
//   - [IMLMultiplyBroadcastableBrick.OutputShapes]
//   - [IMLMultiplyBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLMultiplyBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLMultiplyBroadcastableBrick.InitWithParameters]
//   - [IMLMultiplyBroadcastableBrick.DebugDescription]
//   - [IMLMultiplyBroadcastableBrick.Description]
//   - [IMLMultiplyBroadcastableBrick.Hash]
//   - [IMLMultiplyBroadcastableBrick.Superclass]
type IMLMultiplyBroadcastableBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	InitWithParameters(parameters objectivec.IObject) MLMultiplyBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLMultiplyBroadcastableBrick) Init() MLMultiplyBroadcastableBrick {
	rv := objc.SendIfResponds[MLMultiplyBroadcastableBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiplyBroadcastableBrick) Autorelease() MLMultiplyBroadcastableBrick {
	rv := objc.SendIfResponds[MLMultiplyBroadcastableBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiplyBroadcastableBrick creates a new MLMultiplyBroadcastableBrick instance.
func NewMLMultiplyBroadcastableBrick() MLMultiplyBroadcastableBrick {
	class := getMLMultiplyBroadcastableBrickClass()
	rv := objc.SendIfResponds[MLMultiplyBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMultiplyBroadcastableBrickWithParameters(parameters objectivec.IObject) MLMultiplyBroadcastableBrick {
	instance := getMLMultiplyBroadcastableBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLMultiplyBroadcastableBrickFromID(rv)
}

func (m MLMultiplyBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLMultiplyBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLMultiplyBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLMultiplyBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLMultiplyBroadcastableBrick {
	rv := objc.SendIfResponds[MLMultiplyBroadcastableBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLMultiplyBroadcastableBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLMultiplyBroadcastableBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLMultiplyBroadcastableBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLMultiplyBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLMultiplyBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLMultiplyBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLMultiplyBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLMultiplyBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLMultiplyBroadcastableBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
