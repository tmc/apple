// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLStackNDBrick] class.
var (
	_MLStackNDBrickClass     MLStackNDBrickClass
	_MLStackNDBrickClassOnce sync.Once
)

func getMLStackNDBrickClass() MLStackNDBrickClass {
	_MLStackNDBrickClassOnce.Do(func() {
		_MLStackNDBrickClass = MLStackNDBrickClass{class: objc.GetClass("MLStackNDBrick")}
	})
	return _MLStackNDBrickClass
}

// GetMLStackNDBrickClass returns the class object for MLStackNDBrick.
func GetMLStackNDBrickClass() MLStackNDBrickClass {
	return getMLStackNDBrickClass()
}

type MLStackNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLStackNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLStackNDBrickClass) Alloc() MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLStackNDBrick.Axis]
//   - [MLStackNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLStackNDBrick.HasGPUSupport]
//   - [MLStackNDBrick.InputRanks]
//   - [MLStackNDBrick.InputShapes]
//   - [MLStackNDBrick.OutputRanks]
//   - [MLStackNDBrick.OutputShapes]
//   - [MLStackNDBrick.SetupForInputShapesWithParameters]
//   - [MLStackNDBrick.ShapeInfoNeeded]
//   - [MLStackNDBrick.InitWithParameters]
//   - [MLStackNDBrick.DebugDescription]
//   - [MLStackNDBrick.Description]
//   - [MLStackNDBrick.Hash]
//   - [MLStackNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick
type MLStackNDBrick struct {
	objectivec.Object
}

// MLStackNDBrickFromID constructs a [MLStackNDBrick] from an objc.ID.
func MLStackNDBrickFromID(id objc.ID) MLStackNDBrick {
	return MLStackNDBrick{objectivec.Object{ID: id}}
}

// Ensure MLStackNDBrick implements IMLStackNDBrick.
var _ IMLStackNDBrick = MLStackNDBrick{}

// An interface definition for the [MLStackNDBrick] class.
//
// # Methods
//
//   - [IMLStackNDBrick.Axis]
//   - [IMLStackNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLStackNDBrick.HasGPUSupport]
//   - [IMLStackNDBrick.InputRanks]
//   - [IMLStackNDBrick.InputShapes]
//   - [IMLStackNDBrick.OutputRanks]
//   - [IMLStackNDBrick.OutputShapes]
//   - [IMLStackNDBrick.SetupForInputShapesWithParameters]
//   - [IMLStackNDBrick.ShapeInfoNeeded]
//   - [IMLStackNDBrick.InitWithParameters]
//   - [IMLStackNDBrick.DebugDescription]
//   - [IMLStackNDBrick.Description]
//   - [IMLStackNDBrick.Hash]
//   - [IMLStackNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick
type IMLStackNDBrick interface {
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
	InitWithParameters(parameters objectivec.IObject) MLStackNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLStackNDBrick) Init() MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLStackNDBrick) Autorelease() MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLStackNDBrick creates a new MLStackNDBrick instance.
func NewMLStackNDBrick() MLStackNDBrick {
	class := getMLStackNDBrickClass()
	rv := objc.Send[MLStackNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/initWithParameters:
func NewStackNDBrickWithParameters(parameters objectivec.IObject) MLStackNDBrick {
	instance := getMLStackNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLStackNDBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/computeOnCPUWithInputTensors:outputTensors:
func (s MLStackNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/hasGPUSupport
func (s MLStackNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/setupForInputShapes:withParameters:
func (s MLStackNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/initWithParameters:
func (s MLStackNDBrick) InitWithParameters(parameters objectivec.IObject) MLStackNDBrick {
	rv := objc.Send[MLStackNDBrick](s.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/axis
func (s MLStackNDBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/debugDescription
func (s MLStackNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/description
func (s MLStackNDBrick) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/hash
func (s MLStackNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/inputRanks
func (s MLStackNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/inputShapes
func (s MLStackNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/outputRanks
func (s MLStackNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/outputShapes
func (s MLStackNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/shapeInfoNeeded
func (s MLStackNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLStackNDBrick/superclass
func (s MLStackNDBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
