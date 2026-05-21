// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFloorBrick] class.
var (
	_MLFloorBrickClass     MLFloorBrickClass
	_MLFloorBrickClassOnce sync.Once
)

func getMLFloorBrickClass() MLFloorBrickClass {
	_MLFloorBrickClassOnce.Do(func() {
		_MLFloorBrickClass = MLFloorBrickClass{class: objc.GetClass("MLFloorBrick")}
	})
	return _MLFloorBrickClass
}

// GetMLFloorBrickClass returns the class object for MLFloorBrick.
func GetMLFloorBrickClass() MLFloorBrickClass {
	return getMLFloorBrickClass()
}

type MLFloorBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFloorBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFloorBrickClass) Alloc() MLFloorBrick {
	rv := objc.Send[MLFloorBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFloorBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLFloorBrick.HasGPUSupport]
//   - [MLFloorBrick.InputRanks]
//   - [MLFloorBrick.InputShapes]
//   - [MLFloorBrick.OutputRanks]
//   - [MLFloorBrick.OutputShapes]
//   - [MLFloorBrick.SetupForInputShapesWithParameters]
//   - [MLFloorBrick.InitWithParameters]
//   - [MLFloorBrick.DebugDescription]
//   - [MLFloorBrick.Description]
//   - [MLFloorBrick.Hash]
//   - [MLFloorBrick.Superclass]
type MLFloorBrick struct {
	objectivec.Object
}

// MLFloorBrickFromID constructs a [MLFloorBrick] from an objc.ID.
func MLFloorBrickFromID(id objc.ID) MLFloorBrick {
	return MLFloorBrick{objectivec.Object{ID: id}}
}

// Ensure MLFloorBrick implements IMLFloorBrick.
var _ IMLFloorBrick = MLFloorBrick{}

// An interface definition for the [MLFloorBrick] class.
//
// # Methods
//
//   - [IMLFloorBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLFloorBrick.HasGPUSupport]
//   - [IMLFloorBrick.InputRanks]
//   - [IMLFloorBrick.InputShapes]
//   - [IMLFloorBrick.OutputRanks]
//   - [IMLFloorBrick.OutputShapes]
//   - [IMLFloorBrick.SetupForInputShapesWithParameters]
//   - [IMLFloorBrick.InitWithParameters]
//   - [IMLFloorBrick.DebugDescription]
//   - [IMLFloorBrick.Description]
//   - [IMLFloorBrick.Hash]
//   - [IMLFloorBrick.Superclass]
type IMLFloorBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLFloorBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLFloorBrick) Init() MLFloorBrick {
	rv := objc.Send[MLFloorBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLFloorBrick) Autorelease() MLFloorBrick {
	rv := objc.Send[MLFloorBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFloorBrick creates a new MLFloorBrick instance.
func NewMLFloorBrick() MLFloorBrick {
	class := getMLFloorBrickClass()
	rv := objc.Send[MLFloorBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewFloorBrickWithParameters(parameters objectivec.IObject) MLFloorBrick {
	instance := getMLFloorBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLFloorBrickFromID(rv)
}

func (m MLFloorBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLFloorBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLFloorBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLFloorBrick) InitWithParameters(parameters objectivec.IObject) MLFloorBrick {
	rv := objc.Send[MLFloorBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLFloorBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLFloorBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLFloorBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLFloorBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFloorBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFloorBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFloorBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFloorBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
