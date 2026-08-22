// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBatchedMatMulBrick] class.
var (
	_MLBatchedMatMulBrickClass     MLBatchedMatMulBrickClass
	_MLBatchedMatMulBrickClassOnce sync.Once
)

func getMLBatchedMatMulBrickClass() MLBatchedMatMulBrickClass {
	_MLBatchedMatMulBrickClassOnce.Do(func() {
		_MLBatchedMatMulBrickClass = MLBatchedMatMulBrickClass{class: objc.GetClass("MLBatchedMatMulBrick")}
	})
	return _MLBatchedMatMulBrickClass
}

// GetMLBatchedMatMulBrickClass returns the class object for MLBatchedMatMulBrick.
func GetMLBatchedMatMulBrickClass() MLBatchedMatMulBrickClass {
	return getMLBatchedMatMulBrickClass()
}

type MLBatchedMatMulBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBatchedMatMulBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBatchedMatMulBrickClass) Alloc() MLBatchedMatMulBrick {
	rv := objc.SendIfResponds[MLBatchedMatMulBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBatchedMatMulBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLBatchedMatMulBrick.HasGPUSupport]
//   - [MLBatchedMatMulBrick.InputRanks]
//   - [MLBatchedMatMulBrick.InputShapes]
//   - [MLBatchedMatMulBrick.OutputRanks]
//   - [MLBatchedMatMulBrick.OutputShapes]
//   - [MLBatchedMatMulBrick.SetupForInputShapesWithParameters]
//   - [MLBatchedMatMulBrick.ShapeInfoNeeded]
//   - [MLBatchedMatMulBrick.TransposeA]
//   - [MLBatchedMatMulBrick.TransposeB]
//   - [MLBatchedMatMulBrick.InitWithParameters]
//   - [MLBatchedMatMulBrick.DebugDescription]
//   - [MLBatchedMatMulBrick.Description]
//   - [MLBatchedMatMulBrick.Hash]
//   - [MLBatchedMatMulBrick.Superclass]
type MLBatchedMatMulBrick struct {
	objectivec.Object
}

// MLBatchedMatMulBrickFromID constructs a [MLBatchedMatMulBrick] from an objc.ID.
func MLBatchedMatMulBrickFromID(id objc.ID) MLBatchedMatMulBrick {
	return MLBatchedMatMulBrick{objectivec.Object{ID: id}}
}

// Ensure MLBatchedMatMulBrick implements IMLBatchedMatMulBrick.
var _ IMLBatchedMatMulBrick = MLBatchedMatMulBrick{}

// An interface definition for the [MLBatchedMatMulBrick] class.
//
// # Methods
//
//   - [IMLBatchedMatMulBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLBatchedMatMulBrick.HasGPUSupport]
//   - [IMLBatchedMatMulBrick.InputRanks]
//   - [IMLBatchedMatMulBrick.InputShapes]
//   - [IMLBatchedMatMulBrick.OutputRanks]
//   - [IMLBatchedMatMulBrick.OutputShapes]
//   - [IMLBatchedMatMulBrick.SetupForInputShapesWithParameters]
//   - [IMLBatchedMatMulBrick.ShapeInfoNeeded]
//   - [IMLBatchedMatMulBrick.TransposeA]
//   - [IMLBatchedMatMulBrick.TransposeB]
//   - [IMLBatchedMatMulBrick.InitWithParameters]
//   - [IMLBatchedMatMulBrick.DebugDescription]
//   - [IMLBatchedMatMulBrick.Description]
//   - [IMLBatchedMatMulBrick.Hash]
//   - [IMLBatchedMatMulBrick.Superclass]
type IMLBatchedMatMulBrick interface {
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
	TransposeA() foundation.NSNumber
	TransposeB() foundation.NSNumber
	InitWithParameters(parameters objectivec.IObject) MLBatchedMatMulBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLBatchedMatMulBrick) Init() MLBatchedMatMulBrick {
	rv := objc.SendIfResponds[MLBatchedMatMulBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBatchedMatMulBrick) Autorelease() MLBatchedMatMulBrick {
	rv := objc.SendIfResponds[MLBatchedMatMulBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBatchedMatMulBrick creates a new MLBatchedMatMulBrick instance.
func NewMLBatchedMatMulBrick() MLBatchedMatMulBrick {
	class := getMLBatchedMatMulBrickClass()
	rv := objc.SendIfResponds[MLBatchedMatMulBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBatchedMatMulBrickWithParameters(parameters objectivec.IObject) MLBatchedMatMulBrick {
	instance := getMLBatchedMatMulBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLBatchedMatMulBrickFromID(rv)
}

func (m MLBatchedMatMulBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLBatchedMatMulBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLBatchedMatMulBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLBatchedMatMulBrick) InitWithParameters(parameters objectivec.IObject) MLBatchedMatMulBrick {
	rv := objc.SendIfResponds[MLBatchedMatMulBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLBatchedMatMulBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBatchedMatMulBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBatchedMatMulBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLBatchedMatMulBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLBatchedMatMulBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLBatchedMatMulBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLBatchedMatMulBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLBatchedMatMulBrick) ShapeInfoNeeded() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
func (m MLBatchedMatMulBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLBatchedMatMulBrick) TransposeA() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("transposeA"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLBatchedMatMulBrick) TransposeB() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("transposeB"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
