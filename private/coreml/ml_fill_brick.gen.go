// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFillBrick] class.
var (
	_MLFillBrickClass     MLFillBrickClass
	_MLFillBrickClassOnce sync.Once
)

func getMLFillBrickClass() MLFillBrickClass {
	_MLFillBrickClassOnce.Do(func() {
		_MLFillBrickClass = MLFillBrickClass{class: objc.GetClass("MLFillBrick")}
	})
	return _MLFillBrickClass
}

// GetMLFillBrickClass returns the class object for MLFillBrick.
func GetMLFillBrickClass() MLFillBrickClass {
	return getMLFillBrickClass()
}

type MLFillBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFillBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFillBrickClass) Alloc() MLFillBrick {
	rv := objc.SendIfResponds[MLFillBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFillBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLFillBrick.HasGPUSupport]
//   - [MLFillBrick.InputRanks]
//   - [MLFillBrick.InputShapes]
//   - [MLFillBrick.OutputRanks]
//   - [MLFillBrick.OutputShapes]
//   - [MLFillBrick.SetupForInputShapesWithParameters]
//   - [MLFillBrick.InitWithParameters]
//   - [MLFillBrick.DebugDescription]
//   - [MLFillBrick.Description]
//   - [MLFillBrick.Hash]
//   - [MLFillBrick.Superclass]
type MLFillBrick struct {
	objectivec.Object
}

// MLFillBrickFromID constructs a [MLFillBrick] from an objc.ID.
func MLFillBrickFromID(id objc.ID) MLFillBrick {
	return MLFillBrick{objectivec.Object{ID: id}}
}

// Ensure MLFillBrick implements IMLFillBrick.
var _ IMLFillBrick = MLFillBrick{}

// An interface definition for the [MLFillBrick] class.
//
// # Methods
//
//   - [IMLFillBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLFillBrick.HasGPUSupport]
//   - [IMLFillBrick.InputRanks]
//   - [IMLFillBrick.InputShapes]
//   - [IMLFillBrick.OutputRanks]
//   - [IMLFillBrick.OutputShapes]
//   - [IMLFillBrick.SetupForInputShapesWithParameters]
//   - [IMLFillBrick.InitWithParameters]
//   - [IMLFillBrick.DebugDescription]
//   - [IMLFillBrick.Description]
//   - [IMLFillBrick.Hash]
//   - [IMLFillBrick.Superclass]
type IMLFillBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLFillBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLFillBrick) Init() MLFillBrick {
	rv := objc.SendIfResponds[MLFillBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLFillBrick) Autorelease() MLFillBrick {
	rv := objc.SendIfResponds[MLFillBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFillBrick creates a new MLFillBrick instance.
func NewMLFillBrick() MLFillBrick {
	class := getMLFillBrickClass()
	rv := objc.SendIfResponds[MLFillBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewFillBrickWithParameters(parameters objectivec.IObject) MLFillBrick {
	instance := getMLFillBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLFillBrickFromID(rv)
}

func (m MLFillBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLFillBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLFillBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLFillBrick) InitWithParameters(parameters objectivec.IObject) MLFillBrick {
	rv := objc.SendIfResponds[MLFillBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLFillBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLFillBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLFillBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLFillBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFillBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFillBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFillBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLFillBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
