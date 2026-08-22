// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLClipBrick] class.
var (
	_MLClipBrickClass     MLClipBrickClass
	_MLClipBrickClassOnce sync.Once
)

func getMLClipBrickClass() MLClipBrickClass {
	_MLClipBrickClassOnce.Do(func() {
		_MLClipBrickClass = MLClipBrickClass{class: objc.GetClass("MLClipBrick")}
	})
	return _MLClipBrickClass
}

// GetMLClipBrickClass returns the class object for MLClipBrick.
func GetMLClipBrickClass() MLClipBrickClass {
	return getMLClipBrickClass()
}

type MLClipBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLClipBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLClipBrickClass) Alloc() MLClipBrick {
	rv := objc.SendIfResponds[MLClipBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLClipBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLClipBrick.HasGPUSupport]
//   - [MLClipBrick.InputRanks]
//   - [MLClipBrick.InputShapes]
//   - [MLClipBrick.OutputRanks]
//   - [MLClipBrick.OutputShapes]
//   - [MLClipBrick.SetupForInputShapesWithParameters]
//   - [MLClipBrick.InitWithParameters]
//   - [MLClipBrick.DebugDescription]
//   - [MLClipBrick.Description]
//   - [MLClipBrick.Hash]
//   - [MLClipBrick.Superclass]
type MLClipBrick struct {
	objectivec.Object
}

// MLClipBrickFromID constructs a [MLClipBrick] from an objc.ID.
func MLClipBrickFromID(id objc.ID) MLClipBrick {
	return MLClipBrick{objectivec.Object{ID: id}}
}

// Ensure MLClipBrick implements IMLClipBrick.
var _ IMLClipBrick = MLClipBrick{}

// An interface definition for the [MLClipBrick] class.
//
// # Methods
//
//   - [IMLClipBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLClipBrick.HasGPUSupport]
//   - [IMLClipBrick.InputRanks]
//   - [IMLClipBrick.InputShapes]
//   - [IMLClipBrick.OutputRanks]
//   - [IMLClipBrick.OutputShapes]
//   - [IMLClipBrick.SetupForInputShapesWithParameters]
//   - [IMLClipBrick.InitWithParameters]
//   - [IMLClipBrick.DebugDescription]
//   - [IMLClipBrick.Description]
//   - [IMLClipBrick.Hash]
//   - [IMLClipBrick.Superclass]
type IMLClipBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLClipBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLClipBrick) Init() MLClipBrick {
	rv := objc.SendIfResponds[MLClipBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLClipBrick) Autorelease() MLClipBrick {
	rv := objc.SendIfResponds[MLClipBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLClipBrick creates a new MLClipBrick instance.
func NewMLClipBrick() MLClipBrick {
	class := getMLClipBrickClass()
	rv := objc.SendIfResponds[MLClipBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewClipBrickWithParameters(parameters objectivec.IObject) MLClipBrick {
	instance := getMLClipBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLClipBrickFromID(rv)
}

func (m MLClipBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLClipBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLClipBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLClipBrick) InitWithParameters(parameters objectivec.IObject) MLClipBrick {
	rv := objc.SendIfResponds[MLClipBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLClipBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLClipBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLClipBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLClipBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLClipBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLClipBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLClipBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLClipBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
