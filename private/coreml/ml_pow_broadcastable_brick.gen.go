// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPowBroadcastableBrick] class.
var (
	_MLPowBroadcastableBrickClass     MLPowBroadcastableBrickClass
	_MLPowBroadcastableBrickClassOnce sync.Once
)

func getMLPowBroadcastableBrickClass() MLPowBroadcastableBrickClass {
	_MLPowBroadcastableBrickClassOnce.Do(func() {
		_MLPowBroadcastableBrickClass = MLPowBroadcastableBrickClass{class: objc.GetClass("MLPowBroadcastableBrick")}
	})
	return _MLPowBroadcastableBrickClass
}

// GetMLPowBroadcastableBrickClass returns the class object for MLPowBroadcastableBrick.
func GetMLPowBroadcastableBrickClass() MLPowBroadcastableBrickClass {
	return getMLPowBroadcastableBrickClass()
}

type MLPowBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPowBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPowBroadcastableBrickClass) Alloc() MLPowBroadcastableBrick {
	rv := objc.Send[MLPowBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPowBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLPowBroadcastableBrick.HasGPUSupport]
//   - [MLPowBroadcastableBrick.InputRanks]
//   - [MLPowBroadcastableBrick.InputShapes]
//   - [MLPowBroadcastableBrick.OutputRanks]
//   - [MLPowBroadcastableBrick.OutputShapes]
//   - [MLPowBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLPowBroadcastableBrick.ShapeInfoNeeded]
//   - [MLPowBroadcastableBrick.InitWithParameters]
//   - [MLPowBroadcastableBrick.DebugDescription]
//   - [MLPowBroadcastableBrick.Description]
//   - [MLPowBroadcastableBrick.Hash]
//   - [MLPowBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick
type MLPowBroadcastableBrick struct {
	objectivec.Object
}

// MLPowBroadcastableBrickFromID constructs a [MLPowBroadcastableBrick] from an objc.ID.
func MLPowBroadcastableBrickFromID(id objc.ID) MLPowBroadcastableBrick {
	return MLPowBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLPowBroadcastableBrick implements IMLPowBroadcastableBrick.
var _ IMLPowBroadcastableBrick = MLPowBroadcastableBrick{}

// An interface definition for the [MLPowBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLPowBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLPowBroadcastableBrick.HasGPUSupport]
//   - [IMLPowBroadcastableBrick.InputRanks]
//   - [IMLPowBroadcastableBrick.InputShapes]
//   - [IMLPowBroadcastableBrick.OutputRanks]
//   - [IMLPowBroadcastableBrick.OutputShapes]
//   - [IMLPowBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLPowBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLPowBroadcastableBrick.InitWithParameters]
//   - [IMLPowBroadcastableBrick.DebugDescription]
//   - [IMLPowBroadcastableBrick.Description]
//   - [IMLPowBroadcastableBrick.Hash]
//   - [IMLPowBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick
type IMLPowBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLPowBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (p MLPowBroadcastableBrick) Init() MLPowBroadcastableBrick {
	rv := objc.Send[MLPowBroadcastableBrick](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPowBroadcastableBrick) Autorelease() MLPowBroadcastableBrick {
	rv := objc.Send[MLPowBroadcastableBrick](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPowBroadcastableBrick creates a new MLPowBroadcastableBrick instance.
func NewMLPowBroadcastableBrick() MLPowBroadcastableBrick {
	class := getMLPowBroadcastableBrickClass()
	rv := objc.Send[MLPowBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/initWithParameters:
func NewPowBroadcastableBrickWithParameters(parameters objectivec.IObject) MLPowBroadcastableBrick {
	instance := getMLPowBroadcastableBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLPowBroadcastableBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/computeOnCPUWithInputTensors:outputTensors:
func (p MLPowBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/hasGPUSupport
func (p MLPowBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/setupForInputShapes:withParameters:
func (p MLPowBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/initWithParameters:
func (p MLPowBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLPowBroadcastableBrick {
	rv := objc.Send[MLPowBroadcastableBrick](p.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/debugDescription
func (p MLPowBroadcastableBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/description
func (p MLPowBroadcastableBrick) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/hash
func (p MLPowBroadcastableBrick) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/inputRanks
func (p MLPowBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/inputShapes
func (p MLPowBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/outputRanks
func (p MLPowBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/outputShapes
func (p MLPowBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/shapeInfoNeeded
func (p MLPowBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPowBroadcastableBrick/superclass
func (p MLPowBroadcastableBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("superclass"))
	return rv
}
