// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSlidingWindowsBrick] class.
var (
	_MLSlidingWindowsBrickClass     MLSlidingWindowsBrickClass
	_MLSlidingWindowsBrickClassOnce sync.Once
)

func getMLSlidingWindowsBrickClass() MLSlidingWindowsBrickClass {
	_MLSlidingWindowsBrickClassOnce.Do(func() {
		_MLSlidingWindowsBrickClass = MLSlidingWindowsBrickClass{class: objc.GetClass("MLSlidingWindowsBrick")}
	})
	return _MLSlidingWindowsBrickClass
}

// GetMLSlidingWindowsBrickClass returns the class object for MLSlidingWindowsBrick.
func GetMLSlidingWindowsBrickClass() MLSlidingWindowsBrickClass {
	return getMLSlidingWindowsBrickClass()
}

type MLSlidingWindowsBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSlidingWindowsBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSlidingWindowsBrickClass) Alloc() MLSlidingWindowsBrick {
	rv := objc.Send[MLSlidingWindowsBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSlidingWindowsBrick.Axis]
//   - [MLSlidingWindowsBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSlidingWindowsBrick.HasGPUSupport]
//   - [MLSlidingWindowsBrick.InputRanks]
//   - [MLSlidingWindowsBrick.InputShapes]
//   - [MLSlidingWindowsBrick.OutputRanks]
//   - [MLSlidingWindowsBrick.OutputShapes]
//   - [MLSlidingWindowsBrick.SetupForInputShapesWithParameters]
//   - [MLSlidingWindowsBrick.ShapeInfoNeeded]
//   - [MLSlidingWindowsBrick.Size]
//   - [MLSlidingWindowsBrick.Step]
//   - [MLSlidingWindowsBrick.InitWithParameters]
//   - [MLSlidingWindowsBrick.DebugDescription]
//   - [MLSlidingWindowsBrick.Description]
//   - [MLSlidingWindowsBrick.Hash]
//   - [MLSlidingWindowsBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick
type MLSlidingWindowsBrick struct {
	objectivec.Object
}

// MLSlidingWindowsBrickFromID constructs a [MLSlidingWindowsBrick] from an objc.ID.
func MLSlidingWindowsBrickFromID(id objc.ID) MLSlidingWindowsBrick {
	return MLSlidingWindowsBrick{objectivec.Object{ID: id}}
}

// Ensure MLSlidingWindowsBrick implements IMLSlidingWindowsBrick.
var _ IMLSlidingWindowsBrick = MLSlidingWindowsBrick{}

// An interface definition for the [MLSlidingWindowsBrick] class.
//
// # Methods
//
//   - [IMLSlidingWindowsBrick.Axis]
//   - [IMLSlidingWindowsBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSlidingWindowsBrick.HasGPUSupport]
//   - [IMLSlidingWindowsBrick.InputRanks]
//   - [IMLSlidingWindowsBrick.InputShapes]
//   - [IMLSlidingWindowsBrick.OutputRanks]
//   - [IMLSlidingWindowsBrick.OutputShapes]
//   - [IMLSlidingWindowsBrick.SetupForInputShapesWithParameters]
//   - [IMLSlidingWindowsBrick.ShapeInfoNeeded]
//   - [IMLSlidingWindowsBrick.Size]
//   - [IMLSlidingWindowsBrick.Step]
//   - [IMLSlidingWindowsBrick.InitWithParameters]
//   - [IMLSlidingWindowsBrick.DebugDescription]
//   - [IMLSlidingWindowsBrick.Description]
//   - [IMLSlidingWindowsBrick.Hash]
//   - [IMLSlidingWindowsBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick
type IMLSlidingWindowsBrick interface {
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
	Size() foundation.NSNumber
	Step() foundation.NSNumber
	InitWithParameters(parameters objectivec.IObject) MLSlidingWindowsBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLSlidingWindowsBrick) Init() MLSlidingWindowsBrick {
	rv := objc.Send[MLSlidingWindowsBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSlidingWindowsBrick) Autorelease() MLSlidingWindowsBrick {
	rv := objc.Send[MLSlidingWindowsBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSlidingWindowsBrick creates a new MLSlidingWindowsBrick instance.
func NewMLSlidingWindowsBrick() MLSlidingWindowsBrick {
	class := getMLSlidingWindowsBrickClass()
	rv := objc.Send[MLSlidingWindowsBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/initWithParameters:
func NewSlidingWindowsBrickWithParameters(parameters objectivec.IObject) MLSlidingWindowsBrick {
	instance := getMLSlidingWindowsBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSlidingWindowsBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/computeOnCPUWithInputTensors:outputTensors:
func (m MLSlidingWindowsBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/hasGPUSupport
func (m MLSlidingWindowsBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/setupForInputShapes:withParameters:
func (m MLSlidingWindowsBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/initWithParameters:
func (m MLSlidingWindowsBrick) InitWithParameters(parameters objectivec.IObject) MLSlidingWindowsBrick {
	rv := objc.Send[MLSlidingWindowsBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/axis
func (m MLSlidingWindowsBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/debugDescription
func (m MLSlidingWindowsBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/description
func (m MLSlidingWindowsBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/hash
func (m MLSlidingWindowsBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/inputRanks
func (m MLSlidingWindowsBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/inputShapes
func (m MLSlidingWindowsBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/outputRanks
func (m MLSlidingWindowsBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/outputShapes
func (m MLSlidingWindowsBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/shapeInfoNeeded
func (m MLSlidingWindowsBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/size
func (m MLSlidingWindowsBrick) Size() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("size"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/step
func (m MLSlidingWindowsBrick) Step() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("step"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSlidingWindowsBrick/superclass
func (m MLSlidingWindowsBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
