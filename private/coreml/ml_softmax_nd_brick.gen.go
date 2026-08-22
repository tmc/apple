// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSoftmaxNDBrick] class.
var (
	_MLSoftmaxNDBrickClass     MLSoftmaxNDBrickClass
	_MLSoftmaxNDBrickClassOnce sync.Once
)

func getMLSoftmaxNDBrickClass() MLSoftmaxNDBrickClass {
	_MLSoftmaxNDBrickClassOnce.Do(func() {
		_MLSoftmaxNDBrickClass = MLSoftmaxNDBrickClass{class: objc.GetClass("MLSoftmaxNDBrick")}
	})
	return _MLSoftmaxNDBrickClass
}

// GetMLSoftmaxNDBrickClass returns the class object for MLSoftmaxNDBrick.
func GetMLSoftmaxNDBrickClass() MLSoftmaxNDBrickClass {
	return getMLSoftmaxNDBrickClass()
}

type MLSoftmaxNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSoftmaxNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSoftmaxNDBrickClass) Alloc() MLSoftmaxNDBrick {
	rv := objc.SendIfResponds[MLSoftmaxNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSoftmaxNDBrick.Axis]
//   - [MLSoftmaxNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSoftmaxNDBrick.HasGPUSupport]
//   - [MLSoftmaxNDBrick.InputRanks]
//   - [MLSoftmaxNDBrick.InputShapes]
//   - [MLSoftmaxNDBrick.OutputRanks]
//   - [MLSoftmaxNDBrick.OutputShapes]
//   - [MLSoftmaxNDBrick.SetupForInputShapesWithParameters]
//   - [MLSoftmaxNDBrick.ShapeInfoNeeded]
//   - [MLSoftmaxNDBrick.InitWithParameters]
//   - [MLSoftmaxNDBrick.DebugDescription]
//   - [MLSoftmaxNDBrick.Description]
//   - [MLSoftmaxNDBrick.Hash]
//   - [MLSoftmaxNDBrick.Superclass]
type MLSoftmaxNDBrick struct {
	objectivec.Object
}

// MLSoftmaxNDBrickFromID constructs a [MLSoftmaxNDBrick] from an objc.ID.
func MLSoftmaxNDBrickFromID(id objc.ID) MLSoftmaxNDBrick {
	return MLSoftmaxNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLSoftmaxNDBrick implements IMLSoftmaxNDBrick.
var _ IMLSoftmaxNDBrick = MLSoftmaxNDBrick{}

// An interface definition for the [MLSoftmaxNDBrick] class.
//
// # Methods
//
//   - [IMLSoftmaxNDBrick.Axis]
//   - [IMLSoftmaxNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSoftmaxNDBrick.HasGPUSupport]
//   - [IMLSoftmaxNDBrick.InputRanks]
//   - [IMLSoftmaxNDBrick.InputShapes]
//   - [IMLSoftmaxNDBrick.OutputRanks]
//   - [IMLSoftmaxNDBrick.OutputShapes]
//   - [IMLSoftmaxNDBrick.SetupForInputShapesWithParameters]
//   - [IMLSoftmaxNDBrick.ShapeInfoNeeded]
//   - [IMLSoftmaxNDBrick.InitWithParameters]
//   - [IMLSoftmaxNDBrick.DebugDescription]
//   - [IMLSoftmaxNDBrick.Description]
//   - [IMLSoftmaxNDBrick.Hash]
//   - [IMLSoftmaxNDBrick.Superclass]
type IMLSoftmaxNDBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLSoftmaxNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSoftmaxNDBrick) Init() MLSoftmaxNDBrick {
	rv := objc.SendIfResponds[MLSoftmaxNDBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSoftmaxNDBrick) Autorelease() MLSoftmaxNDBrick {
	rv := objc.SendIfResponds[MLSoftmaxNDBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSoftmaxNDBrick creates a new MLSoftmaxNDBrick instance.
func NewMLSoftmaxNDBrick() MLSoftmaxNDBrick {
	class := getMLSoftmaxNDBrickClass()
	rv := objc.SendIfResponds[MLSoftmaxNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSoftmaxNDBrickWithParameters(parameters objectivec.IObject) MLSoftmaxNDBrick {
	instance := getMLSoftmaxNDBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSoftmaxNDBrickFromID(rv)
}

func (m MLSoftmaxNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLSoftmaxNDBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLSoftmaxNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLSoftmaxNDBrick) InitWithParameters(parameters objectivec.IObject) MLSoftmaxNDBrick {
	rv := objc.SendIfResponds[MLSoftmaxNDBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLSoftmaxNDBrick) Axis() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLSoftmaxNDBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSoftmaxNDBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSoftmaxNDBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSoftmaxNDBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSoftmaxNDBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSoftmaxNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSoftmaxNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSoftmaxNDBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLSoftmaxNDBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
