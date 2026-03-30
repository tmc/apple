// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBroadcastToBrick] class.
var (
	_MLBroadcastToBrickClass     MLBroadcastToBrickClass
	_MLBroadcastToBrickClassOnce sync.Once
)

func getMLBroadcastToBrickClass() MLBroadcastToBrickClass {
	_MLBroadcastToBrickClassOnce.Do(func() {
		_MLBroadcastToBrickClass = MLBroadcastToBrickClass{class: objc.GetClass("MLBroadcastToBrick")}
	})
	return _MLBroadcastToBrickClass
}

// GetMLBroadcastToBrickClass returns the class object for MLBroadcastToBrick.
func GetMLBroadcastToBrickClass() MLBroadcastToBrickClass {
	return getMLBroadcastToBrickClass()
}

type MLBroadcastToBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBroadcastToBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBroadcastToBrickClass) Alloc() MLBroadcastToBrick {
	rv := objc.Send[MLBroadcastToBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBroadcastToBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLBroadcastToBrick.HasGPUSupport]
//   - [MLBroadcastToBrick.InputRanks]
//   - [MLBroadcastToBrick.InputShapes]
//   - [MLBroadcastToBrick.OutputRanks]
//   - [MLBroadcastToBrick.OutputShapes]
//   - [MLBroadcastToBrick.SetupForInputShapesWithParameters]
//   - [MLBroadcastToBrick.ShapeInfoNeeded]
//   - [MLBroadcastToBrick.TargetShape]
//   - [MLBroadcastToBrick.InitWithParameters]
//   - [MLBroadcastToBrick.DebugDescription]
//   - [MLBroadcastToBrick.Description]
//   - [MLBroadcastToBrick.Hash]
//   - [MLBroadcastToBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick
type MLBroadcastToBrick struct {
	objectivec.Object
}

// MLBroadcastToBrickFromID constructs a [MLBroadcastToBrick] from an objc.ID.
func MLBroadcastToBrickFromID(id objc.ID) MLBroadcastToBrick {
	return MLBroadcastToBrick{objectivec.Object{ID: id}}
}

// Ensure MLBroadcastToBrick implements IMLBroadcastToBrick.
var _ IMLBroadcastToBrick = MLBroadcastToBrick{}

// An interface definition for the [MLBroadcastToBrick] class.
//
// # Methods
//
//   - [IMLBroadcastToBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLBroadcastToBrick.HasGPUSupport]
//   - [IMLBroadcastToBrick.InputRanks]
//   - [IMLBroadcastToBrick.InputShapes]
//   - [IMLBroadcastToBrick.OutputRanks]
//   - [IMLBroadcastToBrick.OutputShapes]
//   - [IMLBroadcastToBrick.SetupForInputShapesWithParameters]
//   - [IMLBroadcastToBrick.ShapeInfoNeeded]
//   - [IMLBroadcastToBrick.TargetShape]
//   - [IMLBroadcastToBrick.InitWithParameters]
//   - [IMLBroadcastToBrick.DebugDescription]
//   - [IMLBroadcastToBrick.Description]
//   - [IMLBroadcastToBrick.Hash]
//   - [IMLBroadcastToBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick
type IMLBroadcastToBrick interface {
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
	TargetShape() foundation.INSArray
	InitWithParameters(parameters objectivec.IObject) MLBroadcastToBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (b MLBroadcastToBrick) Init() MLBroadcastToBrick {
	rv := objc.Send[MLBroadcastToBrick](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBroadcastToBrick) Autorelease() MLBroadcastToBrick {
	rv := objc.Send[MLBroadcastToBrick](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBroadcastToBrick creates a new MLBroadcastToBrick instance.
func NewMLBroadcastToBrick() MLBroadcastToBrick {
	class := getMLBroadcastToBrickClass()
	rv := objc.Send[MLBroadcastToBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/initWithParameters:
func NewBroadcastToBrickWithParameters(parameters objectivec.IObject) MLBroadcastToBrick {
	instance := getMLBroadcastToBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLBroadcastToBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/computeOnCPUWithInputTensors:outputTensors:
func (b MLBroadcastToBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/hasGPUSupport
func (b MLBroadcastToBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/setupForInputShapes:withParameters:
func (b MLBroadcastToBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/initWithParameters:
func (b MLBroadcastToBrick) InitWithParameters(parameters objectivec.IObject) MLBroadcastToBrick {
	rv := objc.Send[MLBroadcastToBrick](b.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/debugDescription
func (b MLBroadcastToBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/description
func (b MLBroadcastToBrick) Description() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/hash
func (b MLBroadcastToBrick) Hash() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/inputRanks
func (b MLBroadcastToBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/inputShapes
func (b MLBroadcastToBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/outputRanks
func (b MLBroadcastToBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/outputShapes
func (b MLBroadcastToBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/shapeInfoNeeded
func (b MLBroadcastToBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/superclass
func (b MLBroadcastToBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](b.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBroadcastToBrick/targetShape
func (b MLBroadcastToBrick) TargetShape() foundation.INSArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("targetShape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
