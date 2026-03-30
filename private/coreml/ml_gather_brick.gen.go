// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGatherBrick] class.
var (
	_MLGatherBrickClass     MLGatherBrickClass
	_MLGatherBrickClassOnce sync.Once
)

func getMLGatherBrickClass() MLGatherBrickClass {
	_MLGatherBrickClassOnce.Do(func() {
		_MLGatherBrickClass = MLGatherBrickClass{class: objc.GetClass("MLGatherBrick")}
	})
	return _MLGatherBrickClass
}

// GetMLGatherBrickClass returns the class object for MLGatherBrick.
func GetMLGatherBrickClass() MLGatherBrickClass {
	return getMLGatherBrickClass()
}

type MLGatherBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGatherBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGatherBrickClass) Alloc() MLGatherBrick {
	rv := objc.Send[MLGatherBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGatherBrick.Axis]
//   - [MLGatherBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLGatherBrick.HasGPUSupport]
//   - [MLGatherBrick.InputRanks]
//   - [MLGatherBrick.InputShapes]
//   - [MLGatherBrick.OutputRanks]
//   - [MLGatherBrick.OutputShapes]
//   - [MLGatherBrick.SetupForInputShapesWithParameters]
//   - [MLGatherBrick.ShapeInfoNeeded]
//   - [MLGatherBrick.InitWithParameters]
//   - [MLGatherBrick.DebugDescription]
//   - [MLGatherBrick.Description]
//   - [MLGatherBrick.Hash]
//   - [MLGatherBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick
type MLGatherBrick struct {
	objectivec.Object
}

// MLGatherBrickFromID constructs a [MLGatherBrick] from an objc.ID.
func MLGatherBrickFromID(id objc.ID) MLGatherBrick {
	return MLGatherBrick{objectivec.Object{ID: id}}
}

// Ensure MLGatherBrick implements IMLGatherBrick.
var _ IMLGatherBrick = MLGatherBrick{}

// An interface definition for the [MLGatherBrick] class.
//
// # Methods
//
//   - [IMLGatherBrick.Axis]
//   - [IMLGatherBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLGatherBrick.HasGPUSupport]
//   - [IMLGatherBrick.InputRanks]
//   - [IMLGatherBrick.InputShapes]
//   - [IMLGatherBrick.OutputRanks]
//   - [IMLGatherBrick.OutputShapes]
//   - [IMLGatherBrick.SetupForInputShapesWithParameters]
//   - [IMLGatherBrick.ShapeInfoNeeded]
//   - [IMLGatherBrick.InitWithParameters]
//   - [IMLGatherBrick.DebugDescription]
//   - [IMLGatherBrick.Description]
//   - [IMLGatherBrick.Hash]
//   - [IMLGatherBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick
type IMLGatherBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLGatherBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (g MLGatherBrick) Init() MLGatherBrick {
	rv := objc.Send[MLGatherBrick](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MLGatherBrick) Autorelease() MLGatherBrick {
	rv := objc.Send[MLGatherBrick](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGatherBrick creates a new MLGatherBrick instance.
func NewMLGatherBrick() MLGatherBrick {
	class := getMLGatherBrickClass()
	rv := objc.Send[MLGatherBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/initWithParameters:
func NewGatherBrickWithParameters(parameters objectivec.IObject) MLGatherBrick {
	instance := getMLGatherBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLGatherBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/computeOnCPUWithInputTensors:outputTensors:
func (g MLGatherBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/hasGPUSupport
func (g MLGatherBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/setupForInputShapes:withParameters:
func (g MLGatherBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/initWithParameters:
func (g MLGatherBrick) InitWithParameters(parameters objectivec.IObject) MLGatherBrick {
	rv := objc.Send[MLGatherBrick](g.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/axis
func (g MLGatherBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/debugDescription
func (g MLGatherBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/description
func (g MLGatherBrick) Description() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/hash
func (g MLGatherBrick) Hash() uint64 {
	rv := objc.Send[uint64](g.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/inputRanks
func (g MLGatherBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/inputShapes
func (g MLGatherBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/outputRanks
func (g MLGatherBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/outputShapes
func (g MLGatherBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/shapeInfoNeeded
func (g MLGatherBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLGatherBrick/superclass
func (g MLGatherBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](g.ID, objc.Sel("superclass"))
	return rv
}
