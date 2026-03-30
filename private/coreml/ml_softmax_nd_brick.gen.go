// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSoftmaxNDBrick] class.
var (
	_MLSoftmaxNDBrickClass     MLSoftmaxNDBrickClass
	_MLSoftmaxNDBrickClassOnce sync.Once
)

func getMLSoftmaxNDBrickClass() MLSoftmaxNDBrickClass {
	_MLSoftmaxNDBrickClassOnce.Do(func() {
		_MLSoftmaxNDBrickClass = MLSoftmaxNDBrickClass{class: objc.GetClass("MLSoftmaxNDBrick")}
	})
	return _MLSoftmaxNDBrickClass
}

// GetMLSoftmaxNDBrickClass returns the class object for MLSoftmaxNDBrick.
func GetMLSoftmaxNDBrickClass() MLSoftmaxNDBrickClass {
	return getMLSoftmaxNDBrickClass()
}

type MLSoftmaxNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSoftmaxNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSoftmaxNDBrickClass) Alloc() MLSoftmaxNDBrick {
	rv := objc.Send[MLSoftmaxNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSoftmaxNDBrick.Axis]
//   - [MLSoftmaxNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSoftmaxNDBrick.HasGPUSupport]
//   - [MLSoftmaxNDBrick.InputRanks]
//   - [MLSoftmaxNDBrick.InputShapes]
//   - [MLSoftmaxNDBrick.OutputRanks]
//   - [MLSoftmaxNDBrick.OutputShapes]
//   - [MLSoftmaxNDBrick.SetupForInputShapesWithParameters]
//   - [MLSoftmaxNDBrick.ShapeInfoNeeded]
//   - [MLSoftmaxNDBrick.InitWithParameters]
//   - [MLSoftmaxNDBrick.DebugDescription]
//   - [MLSoftmaxNDBrick.Description]
//   - [MLSoftmaxNDBrick.Hash]
//   - [MLSoftmaxNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick
type MLSoftmaxNDBrick struct {
	objectivec.Object
}

// MLSoftmaxNDBrickFromID constructs a [MLSoftmaxNDBrick] from an objc.ID.
func MLSoftmaxNDBrickFromID(id objc.ID) MLSoftmaxNDBrick {
	return MLSoftmaxNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLSoftmaxNDBrick implements IMLSoftmaxNDBrick.
var _ IMLSoftmaxNDBrick = MLSoftmaxNDBrick{}

// An interface definition for the [MLSoftmaxNDBrick] class.
//
// # Methods
//
//   - [IMLSoftmaxNDBrick.Axis]
//   - [IMLSoftmaxNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSoftmaxNDBrick.HasGPUSupport]
//   - [IMLSoftmaxNDBrick.InputRanks]
//   - [IMLSoftmaxNDBrick.InputShapes]
//   - [IMLSoftmaxNDBrick.OutputRanks]
//   - [IMLSoftmaxNDBrick.OutputShapes]
//   - [IMLSoftmaxNDBrick.SetupForInputShapesWithParameters]
//   - [IMLSoftmaxNDBrick.ShapeInfoNeeded]
//   - [IMLSoftmaxNDBrick.InitWithParameters]
//   - [IMLSoftmaxNDBrick.DebugDescription]
//   - [IMLSoftmaxNDBrick.Description]
//   - [IMLSoftmaxNDBrick.Hash]
//   - [IMLSoftmaxNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick
type IMLSoftmaxNDBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLSoftmaxNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSoftmaxNDBrick) Init() MLSoftmaxNDBrick {
	rv := objc.Send[MLSoftmaxNDBrick](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSoftmaxNDBrick) Autorelease() MLSoftmaxNDBrick {
	rv := objc.Send[MLSoftmaxNDBrick](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSoftmaxNDBrick creates a new MLSoftmaxNDBrick instance.
func NewMLSoftmaxNDBrick() MLSoftmaxNDBrick {
	class := getMLSoftmaxNDBrickClass()
	rv := objc.Send[MLSoftmaxNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/initWithParameters:
func NewSoftmaxNDBrickWithParameters(parameters objectivec.IObject) MLSoftmaxNDBrick {
	instance := getMLSoftmaxNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSoftmaxNDBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/computeOnCPUWithInputTensors:outputTensors:
func (s MLSoftmaxNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/hasGPUSupport
func (s MLSoftmaxNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/setupForInputShapes:withParameters:
func (s MLSoftmaxNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/initWithParameters:
func (s MLSoftmaxNDBrick) InitWithParameters(parameters objectivec.IObject) MLSoftmaxNDBrick {
	rv := objc.Send[MLSoftmaxNDBrick](s.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/axis
func (s MLSoftmaxNDBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/debugDescription
func (s MLSoftmaxNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/description
func (s MLSoftmaxNDBrick) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/hash
func (s MLSoftmaxNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/inputRanks
func (s MLSoftmaxNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/inputShapes
func (s MLSoftmaxNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/outputRanks
func (s MLSoftmaxNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/outputShapes
func (s MLSoftmaxNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/shapeInfoNeeded
func (s MLSoftmaxNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSoftmaxNDBrick/superclass
func (s MLSoftmaxNDBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
