// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSubtractBroadcastableBrick] class.
var (
	_MLSubtractBroadcastableBrickClass     MLSubtractBroadcastableBrickClass
	_MLSubtractBroadcastableBrickClassOnce sync.Once
)

func getMLSubtractBroadcastableBrickClass() MLSubtractBroadcastableBrickClass {
	_MLSubtractBroadcastableBrickClassOnce.Do(func() {
		_MLSubtractBroadcastableBrickClass = MLSubtractBroadcastableBrickClass{class: objc.GetClass("MLSubtractBroadcastableBrick")}
	})
	return _MLSubtractBroadcastableBrickClass
}

// GetMLSubtractBroadcastableBrickClass returns the class object for MLSubtractBroadcastableBrick.
func GetMLSubtractBroadcastableBrickClass() MLSubtractBroadcastableBrickClass {
	return getMLSubtractBroadcastableBrickClass()
}

type MLSubtractBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSubtractBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSubtractBroadcastableBrickClass) Alloc() MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSubtractBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSubtractBroadcastableBrick.HasGPUSupport]
//   - [MLSubtractBroadcastableBrick.InputRanks]
//   - [MLSubtractBroadcastableBrick.InputShapes]
//   - [MLSubtractBroadcastableBrick.OutputRanks]
//   - [MLSubtractBroadcastableBrick.OutputShapes]
//   - [MLSubtractBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLSubtractBroadcastableBrick.ShapeInfoNeeded]
//   - [MLSubtractBroadcastableBrick.InitWithParameters]
//   - [MLSubtractBroadcastableBrick.DebugDescription]
//   - [MLSubtractBroadcastableBrick.Description]
//   - [MLSubtractBroadcastableBrick.Hash]
//   - [MLSubtractBroadcastableBrick.Superclass]
type MLSubtractBroadcastableBrick struct {
	objectivec.Object
}

// MLSubtractBroadcastableBrickFromID constructs a [MLSubtractBroadcastableBrick] from an objc.ID.
func MLSubtractBroadcastableBrickFromID(id objc.ID) MLSubtractBroadcastableBrick {
	return MLSubtractBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLSubtractBroadcastableBrick implements IMLSubtractBroadcastableBrick.
var _ IMLSubtractBroadcastableBrick = MLSubtractBroadcastableBrick{}

// An interface definition for the [MLSubtractBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLSubtractBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSubtractBroadcastableBrick.HasGPUSupport]
//   - [IMLSubtractBroadcastableBrick.InputRanks]
//   - [IMLSubtractBroadcastableBrick.InputShapes]
//   - [IMLSubtractBroadcastableBrick.OutputRanks]
//   - [IMLSubtractBroadcastableBrick.OutputShapes]
//   - [IMLSubtractBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLSubtractBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLSubtractBroadcastableBrick.InitWithParameters]
//   - [IMLSubtractBroadcastableBrick.DebugDescription]
//   - [IMLSubtractBroadcastableBrick.Description]
//   - [IMLSubtractBroadcastableBrick.Hash]
//   - [IMLSubtractBroadcastableBrick.Superclass]
type IMLSubtractBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLSubtractBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSubtractBroadcastableBrick) Init() MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSubtractBroadcastableBrick) Autorelease() MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSubtractBroadcastableBrick creates a new MLSubtractBroadcastableBrick instance.
func NewMLSubtractBroadcastableBrick() MLSubtractBroadcastableBrick {
	class := getMLSubtractBroadcastableBrickClass()
	rv := objc.Send[MLSubtractBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSubtractBroadcastableBrickWithParameters(parameters objectivec.IObject) MLSubtractBroadcastableBrick {
	instance := getMLSubtractBroadcastableBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSubtractBroadcastableBrickFromID(rv)
}

func (m MLSubtractBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLSubtractBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLSubtractBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLSubtractBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLSubtractBroadcastableBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSubtractBroadcastableBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSubtractBroadcastableBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSubtractBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSubtractBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSubtractBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSubtractBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSubtractBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLSubtractBroadcastableBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
