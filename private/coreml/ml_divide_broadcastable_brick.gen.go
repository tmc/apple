// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDivideBroadcastableBrick] class.
var (
	_MLDivideBroadcastableBrickClass     MLDivideBroadcastableBrickClass
	_MLDivideBroadcastableBrickClassOnce sync.Once
)

func getMLDivideBroadcastableBrickClass() MLDivideBroadcastableBrickClass {
	_MLDivideBroadcastableBrickClassOnce.Do(func() {
		_MLDivideBroadcastableBrickClass = MLDivideBroadcastableBrickClass{class: objc.GetClass("MLDivideBroadcastableBrick")}
	})
	return _MLDivideBroadcastableBrickClass
}

// GetMLDivideBroadcastableBrickClass returns the class object for MLDivideBroadcastableBrick.
func GetMLDivideBroadcastableBrickClass() MLDivideBroadcastableBrickClass {
	return getMLDivideBroadcastableBrickClass()
}

type MLDivideBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDivideBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDivideBroadcastableBrickClass) Alloc() MLDivideBroadcastableBrick {
	rv := objc.SendIfResponds[MLDivideBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDivideBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLDivideBroadcastableBrick.HasGPUSupport]
//   - [MLDivideBroadcastableBrick.InputRanks]
//   - [MLDivideBroadcastableBrick.InputShapes]
//   - [MLDivideBroadcastableBrick.OutputRanks]
//   - [MLDivideBroadcastableBrick.OutputShapes]
//   - [MLDivideBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLDivideBroadcastableBrick.ShapeInfoNeeded]
//   - [MLDivideBroadcastableBrick.InitWithParameters]
//   - [MLDivideBroadcastableBrick.DebugDescription]
//   - [MLDivideBroadcastableBrick.Description]
//   - [MLDivideBroadcastableBrick.Hash]
//   - [MLDivideBroadcastableBrick.Superclass]
type MLDivideBroadcastableBrick struct {
	objectivec.Object
}

// MLDivideBroadcastableBrickFromID constructs a [MLDivideBroadcastableBrick] from an objc.ID.
func MLDivideBroadcastableBrickFromID(id objc.ID) MLDivideBroadcastableBrick {
	return MLDivideBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLDivideBroadcastableBrick implements IMLDivideBroadcastableBrick.
var _ IMLDivideBroadcastableBrick = MLDivideBroadcastableBrick{}

// An interface definition for the [MLDivideBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLDivideBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLDivideBroadcastableBrick.HasGPUSupport]
//   - [IMLDivideBroadcastableBrick.InputRanks]
//   - [IMLDivideBroadcastableBrick.InputShapes]
//   - [IMLDivideBroadcastableBrick.OutputRanks]
//   - [IMLDivideBroadcastableBrick.OutputShapes]
//   - [IMLDivideBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLDivideBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLDivideBroadcastableBrick.InitWithParameters]
//   - [IMLDivideBroadcastableBrick.DebugDescription]
//   - [IMLDivideBroadcastableBrick.Description]
//   - [IMLDivideBroadcastableBrick.Hash]
//   - [IMLDivideBroadcastableBrick.Superclass]
type IMLDivideBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLDivideBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLDivideBroadcastableBrick) Init() MLDivideBroadcastableBrick {
	rv := objc.SendIfResponds[MLDivideBroadcastableBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDivideBroadcastableBrick) Autorelease() MLDivideBroadcastableBrick {
	rv := objc.SendIfResponds[MLDivideBroadcastableBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDivideBroadcastableBrick creates a new MLDivideBroadcastableBrick instance.
func NewMLDivideBroadcastableBrick() MLDivideBroadcastableBrick {
	class := getMLDivideBroadcastableBrickClass()
	rv := objc.SendIfResponds[MLDivideBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDivideBroadcastableBrickWithParameters(parameters objectivec.IObject) MLDivideBroadcastableBrick {
	instance := getMLDivideBroadcastableBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLDivideBroadcastableBrickFromID(rv)
}

func (m MLDivideBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLDivideBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLDivideBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLDivideBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLDivideBroadcastableBrick {
	rv := objc.SendIfResponds[MLDivideBroadcastableBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLDivideBroadcastableBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLDivideBroadcastableBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLDivideBroadcastableBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLDivideBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLDivideBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLDivideBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLDivideBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLDivideBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLDivideBroadcastableBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
