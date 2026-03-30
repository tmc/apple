// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDivideBroadcastableBrick] class.
var (
	_MLDivideBroadcastableBrickClass     MLDivideBroadcastableBrickClass
	_MLDivideBroadcastableBrickClassOnce sync.Once
)

func getMLDivideBroadcastableBrickClass() MLDivideBroadcastableBrickClass {
	_MLDivideBroadcastableBrickClassOnce.Do(func() {
		_MLDivideBroadcastableBrickClass = MLDivideBroadcastableBrickClass{class: objc.GetClass("MLDivideBroadcastableBrick")}
	})
	return _MLDivideBroadcastableBrickClass
}

// GetMLDivideBroadcastableBrickClass returns the class object for MLDivideBroadcastableBrick.
func GetMLDivideBroadcastableBrickClass() MLDivideBroadcastableBrickClass {
	return getMLDivideBroadcastableBrickClass()
}

type MLDivideBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDivideBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDivideBroadcastableBrickClass) Alloc() MLDivideBroadcastableBrick {
	rv := objc.Send[MLDivideBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDivideBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLDivideBroadcastableBrick.HasGPUSupport]
//   - [MLDivideBroadcastableBrick.InputRanks]
//   - [MLDivideBroadcastableBrick.InputShapes]
//   - [MLDivideBroadcastableBrick.OutputRanks]
//   - [MLDivideBroadcastableBrick.OutputShapes]
//   - [MLDivideBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLDivideBroadcastableBrick.ShapeInfoNeeded]
//   - [MLDivideBroadcastableBrick.InitWithParameters]
//   - [MLDivideBroadcastableBrick.DebugDescription]
//   - [MLDivideBroadcastableBrick.Description]
//   - [MLDivideBroadcastableBrick.Hash]
//   - [MLDivideBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick
type MLDivideBroadcastableBrick struct {
	objectivec.Object
}

// MLDivideBroadcastableBrickFromID constructs a [MLDivideBroadcastableBrick] from an objc.ID.
func MLDivideBroadcastableBrickFromID(id objc.ID) MLDivideBroadcastableBrick {
	return MLDivideBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLDivideBroadcastableBrick implements IMLDivideBroadcastableBrick.
var _ IMLDivideBroadcastableBrick = MLDivideBroadcastableBrick{}

// An interface definition for the [MLDivideBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLDivideBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLDivideBroadcastableBrick.HasGPUSupport]
//   - [IMLDivideBroadcastableBrick.InputRanks]
//   - [IMLDivideBroadcastableBrick.InputShapes]
//   - [IMLDivideBroadcastableBrick.OutputRanks]
//   - [IMLDivideBroadcastableBrick.OutputShapes]
//   - [IMLDivideBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLDivideBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLDivideBroadcastableBrick.InitWithParameters]
//   - [IMLDivideBroadcastableBrick.DebugDescription]
//   - [IMLDivideBroadcastableBrick.Description]
//   - [IMLDivideBroadcastableBrick.Hash]
//   - [IMLDivideBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick
type IMLDivideBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLDivideBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (d MLDivideBroadcastableBrick) Init() MLDivideBroadcastableBrick {
	rv := objc.Send[MLDivideBroadcastableBrick](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDivideBroadcastableBrick) Autorelease() MLDivideBroadcastableBrick {
	rv := objc.Send[MLDivideBroadcastableBrick](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDivideBroadcastableBrick creates a new MLDivideBroadcastableBrick instance.
func NewMLDivideBroadcastableBrick() MLDivideBroadcastableBrick {
	class := getMLDivideBroadcastableBrickClass()
	rv := objc.Send[MLDivideBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/initWithParameters:
func NewDivideBroadcastableBrickWithParameters(parameters objectivec.IObject) MLDivideBroadcastableBrick {
	instance := getMLDivideBroadcastableBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLDivideBroadcastableBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/computeOnCPUWithInputTensors:outputTensors:
func (d MLDivideBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/hasGPUSupport
func (d MLDivideBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/setupForInputShapes:withParameters:
func (d MLDivideBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/initWithParameters:
func (d MLDivideBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLDivideBroadcastableBrick {
	rv := objc.Send[MLDivideBroadcastableBrick](d.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/debugDescription
func (d MLDivideBroadcastableBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/description
func (d MLDivideBroadcastableBrick) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/hash
func (d MLDivideBroadcastableBrick) Hash() uint64 {
	rv := objc.Send[uint64](d.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/inputRanks
func (d MLDivideBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/inputShapes
func (d MLDivideBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/outputRanks
func (d MLDivideBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/outputShapes
func (d MLDivideBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/shapeInfoNeeded
func (d MLDivideBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDivideBroadcastableBrick/superclass
func (d MLDivideBroadcastableBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](d.ID, objc.Sel("superclass"))
	return rv
}
