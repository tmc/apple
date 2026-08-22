// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAddBroadcastableBrick] class.
var (
	_MLAddBroadcastableBrickClass     MLAddBroadcastableBrickClass
	_MLAddBroadcastableBrickClassOnce sync.Once
)

func getMLAddBroadcastableBrickClass() MLAddBroadcastableBrickClass {
	_MLAddBroadcastableBrickClassOnce.Do(func() {
		_MLAddBroadcastableBrickClass = MLAddBroadcastableBrickClass{class: objc.GetClass("MLAddBroadcastableBrick")}
	})
	return _MLAddBroadcastableBrickClass
}

// GetMLAddBroadcastableBrickClass returns the class object for MLAddBroadcastableBrick.
func GetMLAddBroadcastableBrickClass() MLAddBroadcastableBrickClass {
	return getMLAddBroadcastableBrickClass()
}

type MLAddBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAddBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAddBroadcastableBrickClass) Alloc() MLAddBroadcastableBrick {
	rv := objc.SendIfResponds[MLAddBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAddBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLAddBroadcastableBrick.HasGPUSupport]
//   - [MLAddBroadcastableBrick.InputRanks]
//   - [MLAddBroadcastableBrick.InputShapes]
//   - [MLAddBroadcastableBrick.OutputRanks]
//   - [MLAddBroadcastableBrick.OutputShapes]
//   - [MLAddBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLAddBroadcastableBrick.ShapeInfoNeeded]
//   - [MLAddBroadcastableBrick.InitWithParameters]
//   - [MLAddBroadcastableBrick.DebugDescription]
//   - [MLAddBroadcastableBrick.Description]
//   - [MLAddBroadcastableBrick.Hash]
//   - [MLAddBroadcastableBrick.Superclass]
type MLAddBroadcastableBrick struct {
	objectivec.Object
}

// MLAddBroadcastableBrickFromID constructs a [MLAddBroadcastableBrick] from an objc.ID.
func MLAddBroadcastableBrickFromID(id objc.ID) MLAddBroadcastableBrick {
	return MLAddBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLAddBroadcastableBrick implements IMLAddBroadcastableBrick.
var _ IMLAddBroadcastableBrick = MLAddBroadcastableBrick{}

// An interface definition for the [MLAddBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLAddBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLAddBroadcastableBrick.HasGPUSupport]
//   - [IMLAddBroadcastableBrick.InputRanks]
//   - [IMLAddBroadcastableBrick.InputShapes]
//   - [IMLAddBroadcastableBrick.OutputRanks]
//   - [IMLAddBroadcastableBrick.OutputShapes]
//   - [IMLAddBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLAddBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLAddBroadcastableBrick.InitWithParameters]
//   - [IMLAddBroadcastableBrick.DebugDescription]
//   - [IMLAddBroadcastableBrick.Description]
//   - [IMLAddBroadcastableBrick.Hash]
//   - [IMLAddBroadcastableBrick.Superclass]
type IMLAddBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLAddBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLAddBroadcastableBrick) Init() MLAddBroadcastableBrick {
	rv := objc.SendIfResponds[MLAddBroadcastableBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLAddBroadcastableBrick) Autorelease() MLAddBroadcastableBrick {
	rv := objc.SendIfResponds[MLAddBroadcastableBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAddBroadcastableBrick creates a new MLAddBroadcastableBrick instance.
func NewMLAddBroadcastableBrick() MLAddBroadcastableBrick {
	class := getMLAddBroadcastableBrickClass()
	rv := objc.SendIfResponds[MLAddBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAddBroadcastableBrickWithParameters(parameters objectivec.IObject) MLAddBroadcastableBrick {
	instance := getMLAddBroadcastableBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLAddBroadcastableBrickFromID(rv)
}

func (m MLAddBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLAddBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLAddBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLAddBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLAddBroadcastableBrick {
	rv := objc.SendIfResponds[MLAddBroadcastableBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLAddBroadcastableBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAddBroadcastableBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLAddBroadcastableBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLAddBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLAddBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLAddBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLAddBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLAddBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLAddBroadcastableBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
