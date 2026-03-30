// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLConcatNDBrick] class.
var (
	_MLConcatNDBrickClass     MLConcatNDBrickClass
	_MLConcatNDBrickClassOnce sync.Once
)

func getMLConcatNDBrickClass() MLConcatNDBrickClass {
	_MLConcatNDBrickClassOnce.Do(func() {
		_MLConcatNDBrickClass = MLConcatNDBrickClass{class: objc.GetClass("MLConcatNDBrick")}
	})
	return _MLConcatNDBrickClass
}

// GetMLConcatNDBrickClass returns the class object for MLConcatNDBrick.
func GetMLConcatNDBrickClass() MLConcatNDBrickClass {
	return getMLConcatNDBrickClass()
}

type MLConcatNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLConcatNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLConcatNDBrickClass) Alloc() MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLConcatNDBrick.Axis]
//   - [MLConcatNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLConcatNDBrick.HasGPUSupport]
//   - [MLConcatNDBrick.InputRanks]
//   - [MLConcatNDBrick.InputShapes]
//   - [MLConcatNDBrick.OutputRanks]
//   - [MLConcatNDBrick.OutputShapes]
//   - [MLConcatNDBrick.SetupForInputShapesWithParameters]
//   - [MLConcatNDBrick.ShapeInfoNeeded]
//   - [MLConcatNDBrick.InitWithParameters]
//   - [MLConcatNDBrick.DebugDescription]
//   - [MLConcatNDBrick.Description]
//   - [MLConcatNDBrick.Hash]
//   - [MLConcatNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick
type MLConcatNDBrick struct {
	objectivec.Object
}

// MLConcatNDBrickFromID constructs a [MLConcatNDBrick] from an objc.ID.
func MLConcatNDBrickFromID(id objc.ID) MLConcatNDBrick {
	return MLConcatNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLConcatNDBrick implements IMLConcatNDBrick.
var _ IMLConcatNDBrick = MLConcatNDBrick{}

// An interface definition for the [MLConcatNDBrick] class.
//
// # Methods
//
//   - [IMLConcatNDBrick.Axis]
//   - [IMLConcatNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLConcatNDBrick.HasGPUSupport]
//   - [IMLConcatNDBrick.InputRanks]
//   - [IMLConcatNDBrick.InputShapes]
//   - [IMLConcatNDBrick.OutputRanks]
//   - [IMLConcatNDBrick.OutputShapes]
//   - [IMLConcatNDBrick.SetupForInputShapesWithParameters]
//   - [IMLConcatNDBrick.ShapeInfoNeeded]
//   - [IMLConcatNDBrick.InitWithParameters]
//   - [IMLConcatNDBrick.DebugDescription]
//   - [IMLConcatNDBrick.Description]
//   - [IMLConcatNDBrick.Hash]
//   - [IMLConcatNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick
type IMLConcatNDBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLConcatNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c MLConcatNDBrick) Init() MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLConcatNDBrick) Autorelease() MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLConcatNDBrick creates a new MLConcatNDBrick instance.
func NewMLConcatNDBrick() MLConcatNDBrick {
	class := getMLConcatNDBrickClass()
	rv := objc.Send[MLConcatNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/initWithParameters:
func NewConcatNDBrickWithParameters(parameters objectivec.IObject) MLConcatNDBrick {
	instance := getMLConcatNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLConcatNDBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/computeOnCPUWithInputTensors:outputTensors:
func (c MLConcatNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/hasGPUSupport
func (c MLConcatNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/setupForInputShapes:withParameters:
func (c MLConcatNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/initWithParameters:
func (c MLConcatNDBrick) InitWithParameters(parameters objectivec.IObject) MLConcatNDBrick {
	rv := objc.Send[MLConcatNDBrick](c.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/axis
func (c MLConcatNDBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/debugDescription
func (c MLConcatNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/description
func (c MLConcatNDBrick) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/hash
func (c MLConcatNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/inputRanks
func (c MLConcatNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/inputShapes
func (c MLConcatNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/outputRanks
func (c MLConcatNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/outputShapes
func (c MLConcatNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/shapeInfoNeeded
func (c MLConcatNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLConcatNDBrick/superclass
func (c MLConcatNDBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
