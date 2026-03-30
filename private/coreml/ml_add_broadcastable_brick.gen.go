// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAddBroadcastableBrick] class.
var (
	_MLAddBroadcastableBrickClass     MLAddBroadcastableBrickClass
	_MLAddBroadcastableBrickClassOnce sync.Once
)

func getMLAddBroadcastableBrickClass() MLAddBroadcastableBrickClass {
	_MLAddBroadcastableBrickClassOnce.Do(func() {
		_MLAddBroadcastableBrickClass = MLAddBroadcastableBrickClass{class: objc.GetClass("MLAddBroadcastableBrick")}
	})
	return _MLAddBroadcastableBrickClass
}

// GetMLAddBroadcastableBrickClass returns the class object for MLAddBroadcastableBrick.
func GetMLAddBroadcastableBrickClass() MLAddBroadcastableBrickClass {
	return getMLAddBroadcastableBrickClass()
}

type MLAddBroadcastableBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAddBroadcastableBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAddBroadcastableBrickClass) Alloc() MLAddBroadcastableBrick {
	rv := objc.Send[MLAddBroadcastableBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAddBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLAddBroadcastableBrick.HasGPUSupport]
//   - [MLAddBroadcastableBrick.InputRanks]
//   - [MLAddBroadcastableBrick.InputShapes]
//   - [MLAddBroadcastableBrick.OutputRanks]
//   - [MLAddBroadcastableBrick.OutputShapes]
//   - [MLAddBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [MLAddBroadcastableBrick.ShapeInfoNeeded]
//   - [MLAddBroadcastableBrick.InitWithParameters]
//   - [MLAddBroadcastableBrick.DebugDescription]
//   - [MLAddBroadcastableBrick.Description]
//   - [MLAddBroadcastableBrick.Hash]
//   - [MLAddBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick
type MLAddBroadcastableBrick struct {
	objectivec.Object
}

// MLAddBroadcastableBrickFromID constructs a [MLAddBroadcastableBrick] from an objc.ID.
func MLAddBroadcastableBrickFromID(id objc.ID) MLAddBroadcastableBrick {
	return MLAddBroadcastableBrick{objectivec.Object{ID: id}}
}

// Ensure MLAddBroadcastableBrick implements IMLAddBroadcastableBrick.
var _ IMLAddBroadcastableBrick = MLAddBroadcastableBrick{}

// An interface definition for the [MLAddBroadcastableBrick] class.
//
// # Methods
//
//   - [IMLAddBroadcastableBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLAddBroadcastableBrick.HasGPUSupport]
//   - [IMLAddBroadcastableBrick.InputRanks]
//   - [IMLAddBroadcastableBrick.InputShapes]
//   - [IMLAddBroadcastableBrick.OutputRanks]
//   - [IMLAddBroadcastableBrick.OutputShapes]
//   - [IMLAddBroadcastableBrick.SetupForInputShapesWithParameters]
//   - [IMLAddBroadcastableBrick.ShapeInfoNeeded]
//   - [IMLAddBroadcastableBrick.InitWithParameters]
//   - [IMLAddBroadcastableBrick.DebugDescription]
//   - [IMLAddBroadcastableBrick.Description]
//   - [IMLAddBroadcastableBrick.Hash]
//   - [IMLAddBroadcastableBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick
type IMLAddBroadcastableBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLAddBroadcastableBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a MLAddBroadcastableBrick) Init() MLAddBroadcastableBrick {
	rv := objc.Send[MLAddBroadcastableBrick](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAddBroadcastableBrick) Autorelease() MLAddBroadcastableBrick {
	rv := objc.Send[MLAddBroadcastableBrick](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAddBroadcastableBrick creates a new MLAddBroadcastableBrick instance.
func NewMLAddBroadcastableBrick() MLAddBroadcastableBrick {
	class := getMLAddBroadcastableBrickClass()
	rv := objc.Send[MLAddBroadcastableBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/initWithParameters:
func NewAddBroadcastableBrickWithParameters(parameters objectivec.IObject) MLAddBroadcastableBrick {
	instance := getMLAddBroadcastableBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLAddBroadcastableBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/computeOnCPUWithInputTensors:outputTensors:
func (a MLAddBroadcastableBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/hasGPUSupport
func (a MLAddBroadcastableBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/setupForInputShapes:withParameters:
func (a MLAddBroadcastableBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/initWithParameters:
func (a MLAddBroadcastableBrick) InitWithParameters(parameters objectivec.IObject) MLAddBroadcastableBrick {
	rv := objc.Send[MLAddBroadcastableBrick](a.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/debugDescription
func (a MLAddBroadcastableBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/description
func (a MLAddBroadcastableBrick) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/hash
func (a MLAddBroadcastableBrick) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/inputRanks
func (a MLAddBroadcastableBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/inputShapes
func (a MLAddBroadcastableBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/outputRanks
func (a MLAddBroadcastableBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/outputShapes
func (a MLAddBroadcastableBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/shapeInfoNeeded
func (a MLAddBroadcastableBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLAddBroadcastableBrick/superclass
func (a MLAddBroadcastableBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}
