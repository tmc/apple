// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSubtractBroadcastableBrick] class.
var (
	_MLSubtractBroadcastableBrickClass     MLSubtractBroadcastableBrickClass
	_MLSubtractBroadcastableBrickClassOnce sync.Once
)

func getMLSubtractBroadcastableBrickClass() MLSubtractBroadcastableBrickClass {
	_MLSubtractBroadcastableBrickClassOnce.Do(func() {
		_MLSubtractBroadcastableBrickClass = MLSubtractBroadcastableBrickClass{class: objc.GetClass("MLSubtractBroadcastableBrick")}
	})
	return _MLSubtractBroadcastableBrickClass
}

// GetMLSubtractBroadcastableBrickClass returns the class object for MLSubtractBroadcastableBrick.
func GetMLSubtractBroadcastableBrickClass() MLSubtractBroadcastableBrickClass {
	return getMLSubtractBroadcastableBrickClass()
}

type MLSubtractBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSubtractBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSubtractBroadcastableBrickClass) Alloc() MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSubtractBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSubtractBroadcastableBrick.HasGPUSupport]
//   - [MLSubtractBroadcastableBrick.InputRanks]
//   - [MLSubtractBroadcastableBrick.InputShapes]
//   - [MLSubtractBroadcastableBrick.OutputRanks]
//   - [MLSubtractBroadcastableBrick.OutputShapes]
//   - [MLSubtractBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLSubtractBroadcastableBrick.ShapeInfoNeeded]
//   - [MLSubtractBroadcastableBrick.InitWithParameters]
//   - [MLSubtractBroadcastableBrick.DebugDescription]
//   - [MLSubtractBroadcastableBrick.Description]
//   - [MLSubtractBroadcastableBrick.Hash]
//   - [MLSubtractBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick
type MLSubtractBroadcastableBrick struct {
	objectivec.Object
}

// MLSubtractBroadcastableBrickFromID constructs a [MLSubtractBroadcastableBrick] from an objc.ID.
func MLSubtractBroadcastableBrickFromID(id objc.ID) MLSubtractBroadcastableBrick {
	return MLSubtractBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLSubtractBroadcastableBrick implements IMLSubtractBroadcastableBrick.
var _ IMLSubtractBroadcastableBrick = MLSubtractBroadcastableBrick{}

// An interface definition for the [MLSubtractBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLSubtractBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSubtractBroadcastableBrick.HasGPUSupport]
//   - [IMLSubtractBroadcastableBrick.InputRanks]
//   - [IMLSubtractBroadcastableBrick.InputShapes]
//   - [IMLSubtractBroadcastableBrick.OutputRanks]
//   - [IMLSubtractBroadcastableBrick.OutputShapes]
//   - [IMLSubtractBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLSubtractBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLSubtractBroadcastableBrick.InitWithParameters]
//   - [IMLSubtractBroadcastableBrick.DebugDescription]
//   - [IMLSubtractBroadcastableBrick.Description]
//   - [IMLSubtractBroadcastableBrick.Hash]
//   - [IMLSubtractBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick
type IMLSubtractBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLSubtractBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSubtractBroadcastableBrick) Init() MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSubtractBroadcastableBrick) Autorelease() MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSubtractBroadcastableBrick creates a new MLSubtractBroadcastableBrick instance.
func NewMLSubtractBroadcastableBrick() MLSubtractBroadcastableBrick {
	class := getMLSubtractBroadcastableBrickClass()
	rv := objc.Send[MLSubtractBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/initWithParameters:
func NewSubtractBroadcastableBrickWithParameters(parameters objectivec.IObject) MLSubtractBroadcastableBrick {
	instance := getMLSubtractBroadcastableBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSubtractBroadcastableBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/computeOnCPUWithInputTensors:outputTensors:
func (s MLSubtractBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/hasGPUSupport
func (s MLSubtractBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/setupForInputShapes:withParameters:
func (s MLSubtractBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/initWithParameters:
func (s MLSubtractBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLSubtractBroadcastableBrick {
	rv := objc.Send[MLSubtractBroadcastableBrick](s.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/debugDescription
func (s MLSubtractBroadcastableBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/description
func (s MLSubtractBroadcastableBrick) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/hash
func (s MLSubtractBroadcastableBrick) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/inputRanks
func (s MLSubtractBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/inputShapes
func (s MLSubtractBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/outputRanks
func (s MLSubtractBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/outputShapes
func (s MLSubtractBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/shapeInfoNeeded
func (s MLSubtractBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSubtractBroadcastableBrick/superclass
func (s MLSubtractBroadcastableBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
