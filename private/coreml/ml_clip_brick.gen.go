// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLClipBrick] class.
var (
	_MLClipBrickClass     MLClipBrickClass
	_MLClipBrickClassOnce sync.Once
)

func getMLClipBrickClass() MLClipBrickClass {
	_MLClipBrickClassOnce.Do(func() {
		_MLClipBrickClass = MLClipBrickClass{class: objc.GetClass("MLClipBrick")}
	})
	return _MLClipBrickClass
}

// GetMLClipBrickClass returns the class object for MLClipBrick.
func GetMLClipBrickClass() MLClipBrickClass {
	return getMLClipBrickClass()
}

type MLClipBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLClipBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLClipBrickClass) Alloc() MLClipBrick {
	rv := objc.Send[MLClipBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLClipBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLClipBrick.HasGPUSupport]
//   - [MLClipBrick.InputRanks]
//   - [MLClipBrick.InputShapes]
//   - [MLClipBrick.OutputRanks]
//   - [MLClipBrick.OutputShapes]
//   - [MLClipBrick.SetupForInputShapesWithParameters]
//   - [MLClipBrick.InitWithParameters]
//   - [MLClipBrick.DebugDescription]
//   - [MLClipBrick.Description]
//   - [MLClipBrick.Hash]
//   - [MLClipBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLClipBrick
type MLClipBrick struct {
	objectivec.Object
}

// MLClipBrickFromID constructs a [MLClipBrick] from an objc.ID.
func MLClipBrickFromID(id objc.ID) MLClipBrick {
	return MLClipBrick{objectivec.Object{ID: id}}
}

// Ensure MLClipBrick implements IMLClipBrick.
var _ IMLClipBrick = MLClipBrick{}

// An interface definition for the [MLClipBrick] class.
//
// # Methods
//
//   - [IMLClipBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLClipBrick.HasGPUSupport]
//   - [IMLClipBrick.InputRanks]
//   - [IMLClipBrick.InputShapes]
//   - [IMLClipBrick.OutputRanks]
//   - [IMLClipBrick.OutputShapes]
//   - [IMLClipBrick.SetupForInputShapesWithParameters]
//   - [IMLClipBrick.InitWithParameters]
//   - [IMLClipBrick.DebugDescription]
//   - [IMLClipBrick.Description]
//   - [IMLClipBrick.Hash]
//   - [IMLClipBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLClipBrick
type IMLClipBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLClipBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c MLClipBrick) Init() MLClipBrick {
	rv := objc.Send[MLClipBrick](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLClipBrick) Autorelease() MLClipBrick {
	rv := objc.Send[MLClipBrick](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLClipBrick creates a new MLClipBrick instance.
func NewMLClipBrick() MLClipBrick {
	class := getMLClipBrickClass()
	rv := objc.Send[MLClipBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/initWithParameters:
func NewClipBrickWithParameters(parameters objectivec.IObject) MLClipBrick {
	instance := getMLClipBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLClipBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/computeOnCPUWithInputTensors:outputTensors:
func (c MLClipBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/hasGPUSupport
func (c MLClipBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/setupForInputShapes:withParameters:
func (c MLClipBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/initWithParameters:
func (c MLClipBrick) InitWithParameters(parameters objectivec.IObject) MLClipBrick {
	rv := objc.Send[MLClipBrick](c.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/debugDescription
func (c MLClipBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/description
func (c MLClipBrick) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/hash
func (c MLClipBrick) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/inputRanks
func (c MLClipBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/inputShapes
func (c MLClipBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/outputRanks
func (c MLClipBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/outputShapes
func (c MLClipBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLClipBrick/superclass
func (c MLClipBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
