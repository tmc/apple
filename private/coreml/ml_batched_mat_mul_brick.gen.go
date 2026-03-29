// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[MLBatchedMatMulBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick
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
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick
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
	Superclass() objc.Class
}

// Init initializes the instance.
func (b MLBatchedMatMulBrick) Init() MLBatchedMatMulBrick {
	rv := objc.Send[MLBatchedMatMulBrick](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBatchedMatMulBrick) Autorelease() MLBatchedMatMulBrick {
	rv := objc.Send[MLBatchedMatMulBrick](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBatchedMatMulBrick creates a new MLBatchedMatMulBrick instance.
func NewMLBatchedMatMulBrick() MLBatchedMatMulBrick {
	class := getMLBatchedMatMulBrickClass()
	rv := objc.Send[MLBatchedMatMulBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/initWithParameters:
func NewBatchedMatMulBrickWithParameters(parameters objectivec.IObject) MLBatchedMatMulBrick {
	instance := getMLBatchedMatMulBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLBatchedMatMulBrickFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/computeOnCPUWithInputTensors:outputTensors:
func (b MLBatchedMatMulBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/hasGPUSupport
func (b MLBatchedMatMulBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("hasGPUSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/setupForInputShapes:withParameters:
func (b MLBatchedMatMulBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/initWithParameters:
func (b MLBatchedMatMulBrick) InitWithParameters(parameters objectivec.IObject) MLBatchedMatMulBrick {
	rv := objc.Send[MLBatchedMatMulBrick](b.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/debugDescription
func (b MLBatchedMatMulBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/description
func (b MLBatchedMatMulBrick) Description() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/hash
func (b MLBatchedMatMulBrick) Hash() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/inputRanks
func (b MLBatchedMatMulBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/inputShapes
func (b MLBatchedMatMulBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/outputRanks
func (b MLBatchedMatMulBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/outputShapes
func (b MLBatchedMatMulBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/shapeInfoNeeded
func (b MLBatchedMatMulBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/superclass
func (b MLBatchedMatMulBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](b.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/transposeA
func (b MLBatchedMatMulBrick) TransposeA() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("transposeA"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLBatchedMatMulBrick/transposeB
func (b MLBatchedMatMulBrick) TransposeB() foundation.NSNumber {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("transposeB"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

