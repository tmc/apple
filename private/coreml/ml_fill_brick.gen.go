// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFillBrick] class.
var (
	_MLFillBrickClass     MLFillBrickClass
	_MLFillBrickClassOnce sync.Once
)

func getMLFillBrickClass() MLFillBrickClass {
	_MLFillBrickClassOnce.Do(func() {
		_MLFillBrickClass = MLFillBrickClass{class: objc.GetClass("MLFillBrick")}
	})
	return _MLFillBrickClass
}

// GetMLFillBrickClass returns the class object for MLFillBrick.
func GetMLFillBrickClass() MLFillBrickClass {
	return getMLFillBrickClass()
}

type MLFillBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFillBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFillBrickClass) Alloc() MLFillBrick {
	rv := objc.Send[MLFillBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFillBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLFillBrick.HasGPUSupport]
//   - [MLFillBrick.InputRanks]
//   - [MLFillBrick.InputShapes]
//   - [MLFillBrick.OutputRanks]
//   - [MLFillBrick.OutputShapes]
//   - [MLFillBrick.SetupForInputShapesWithParameters]
//   - [MLFillBrick.InitWithParameters]
//   - [MLFillBrick.DebugDescription]
//   - [MLFillBrick.Description]
//   - [MLFillBrick.Hash]
//   - [MLFillBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLFillBrick
type MLFillBrick struct {
	objectivec.Object
}

// MLFillBrickFromID constructs a [MLFillBrick] from an objc.ID.
func MLFillBrickFromID(id objc.ID) MLFillBrick {
	return MLFillBrick{objectivec.Object{ID: id}}
}

// Ensure MLFillBrick implements IMLFillBrick.
var _ IMLFillBrick = MLFillBrick{}

// An interface definition for the [MLFillBrick] class.
//
// # Methods
//
//   - [IMLFillBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLFillBrick.HasGPUSupport]
//   - [IMLFillBrick.InputRanks]
//   - [IMLFillBrick.InputShapes]
//   - [IMLFillBrick.OutputRanks]
//   - [IMLFillBrick.OutputShapes]
//   - [IMLFillBrick.SetupForInputShapesWithParameters]
//   - [IMLFillBrick.InitWithParameters]
//   - [IMLFillBrick.DebugDescription]
//   - [IMLFillBrick.Description]
//   - [IMLFillBrick.Hash]
//   - [IMLFillBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLFillBrick
type IMLFillBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLFillBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (f MLFillBrick) Init() MLFillBrick {
	rv := objc.Send[MLFillBrick](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFillBrick) Autorelease() MLFillBrick {
	rv := objc.Send[MLFillBrick](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFillBrick creates a new MLFillBrick instance.
func NewMLFillBrick() MLFillBrick {
	class := getMLFillBrickClass()
	rv := objc.Send[MLFillBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/initWithParameters:
func NewFillBrickWithParameters(parameters objectivec.IObject) MLFillBrick {
	instance := getMLFillBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLFillBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/computeOnCPUWithInputTensors:outputTensors:
func (f MLFillBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](f.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/hasGPUSupport
func (f MLFillBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/setupForInputShapes:withParameters:
func (f MLFillBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/initWithParameters:
func (f MLFillBrick) InitWithParameters(parameters objectivec.IObject) MLFillBrick {
	rv := objc.Send[MLFillBrick](f.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/debugDescription
func (f MLFillBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/description
func (f MLFillBrick) Description() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/hash
func (f MLFillBrick) Hash() uint64 {
	rv := objc.Send[uint64](f.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/inputRanks
func (f MLFillBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/inputShapes
func (f MLFillBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/outputRanks
func (f MLFillBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/outputShapes
func (f MLFillBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFillBrick/superclass
func (f MLFillBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](f.ID, objc.Sel("superclass"))
	return rv
}
