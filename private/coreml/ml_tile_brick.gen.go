// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLTileBrick] class.
var (
	_MLTileBrickClass     MLTileBrickClass
	_MLTileBrickClassOnce sync.Once
)

func getMLTileBrickClass() MLTileBrickClass {
	_MLTileBrickClassOnce.Do(func() {
		_MLTileBrickClass = MLTileBrickClass{class: objc.GetClass("MLTileBrick")}
	})
	return _MLTileBrickClass
}

// GetMLTileBrickClass returns the class object for MLTileBrick.
func GetMLTileBrickClass() MLTileBrickClass {
	return getMLTileBrickClass()
}

type MLTileBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTileBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTileBrickClass) Alloc() MLTileBrick {
	rv := objc.Send[MLTileBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLTileBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLTileBrick.HasGPUSupport]
//   - [MLTileBrick.InputRanks]
//   - [MLTileBrick.InputShapes]
//   - [MLTileBrick.OutputRanks]
//   - [MLTileBrick.OutputShapes]
//   - [MLTileBrick.Reps]
//   - [MLTileBrick.SetupForInputShapesWithParameters]
//   - [MLTileBrick.ShapeInfoNeeded]
//   - [MLTileBrick.InitWithParameters]
//   - [MLTileBrick.DebugDescription]
//   - [MLTileBrick.Description]
//   - [MLTileBrick.Hash]
//   - [MLTileBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLTileBrick
type MLTileBrick struct {
	objectivec.Object
}

// MLTileBrickFromID constructs a [MLTileBrick] from an objc.ID.
func MLTileBrickFromID(id objc.ID) MLTileBrick {
	return MLTileBrick{objectivec.Object{ID: id}}
}

// Ensure MLTileBrick implements IMLTileBrick.
var _ IMLTileBrick = MLTileBrick{}

// An interface definition for the [MLTileBrick] class.
//
// # Methods
//
//   - [IMLTileBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLTileBrick.HasGPUSupport]
//   - [IMLTileBrick.InputRanks]
//   - [IMLTileBrick.InputShapes]
//   - [IMLTileBrick.OutputRanks]
//   - [IMLTileBrick.OutputShapes]
//   - [IMLTileBrick.Reps]
//   - [IMLTileBrick.SetupForInputShapesWithParameters]
//   - [IMLTileBrick.ShapeInfoNeeded]
//   - [IMLTileBrick.InitWithParameters]
//   - [IMLTileBrick.DebugDescription]
//   - [IMLTileBrick.Description]
//   - [IMLTileBrick.Hash]
//   - [IMLTileBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLTileBrick
type IMLTileBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	Reps() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	InitWithParameters(parameters objectivec.IObject) MLTileBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (t MLTileBrick) Init() MLTileBrick {
	rv := objc.Send[MLTileBrick](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MLTileBrick) Autorelease() MLTileBrick {
	rv := objc.Send[MLTileBrick](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTileBrick creates a new MLTileBrick instance.
func NewMLTileBrick() MLTileBrick {
	class := getMLTileBrickClass()
	rv := objc.Send[MLTileBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/initWithParameters:
func NewTileBrickWithParameters(parameters objectivec.IObject) MLTileBrick {
	instance := getMLTileBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLTileBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/computeOnCPUWithInputTensors:outputTensors:
func (t MLTileBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/hasGPUSupport
func (t MLTileBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/setupForInputShapes:withParameters:
func (t MLTileBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/initWithParameters:
func (t MLTileBrick) InitWithParameters(parameters objectivec.IObject) MLTileBrick {
	rv := objc.Send[MLTileBrick](t.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/debugDescription
func (t MLTileBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/description
func (t MLTileBrick) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/hash
func (t MLTileBrick) Hash() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/inputRanks
func (t MLTileBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/inputShapes
func (t MLTileBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/outputRanks
func (t MLTileBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/outputShapes
func (t MLTileBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/reps
func (t MLTileBrick) Reps() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("reps"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/shapeInfoNeeded
func (t MLTileBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTileBrick/superclass
func (t MLTileBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](t.ID, objc.Sel("superclass"))
	return rv
}
