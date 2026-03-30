// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCeilBrick] class.
var (
	_MLCeilBrickClass     MLCeilBrickClass
	_MLCeilBrickClassOnce sync.Once
)

func getMLCeilBrickClass() MLCeilBrickClass {
	_MLCeilBrickClassOnce.Do(func() {
		_MLCeilBrickClass = MLCeilBrickClass{class: objc.GetClass("MLCeilBrick")}
	})
	return _MLCeilBrickClass
}

// GetMLCeilBrickClass returns the class object for MLCeilBrick.
func GetMLCeilBrickClass() MLCeilBrickClass {
	return getMLCeilBrickClass()
}

type MLCeilBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCeilBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCeilBrickClass) Alloc() MLCeilBrick {
	rv := objc.Send[MLCeilBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLCeilBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLCeilBrick.HasGPUSupport]
//   - [MLCeilBrick.InputRanks]
//   - [MLCeilBrick.InputShapes]
//   - [MLCeilBrick.OutputRanks]
//   - [MLCeilBrick.OutputShapes]
//   - [MLCeilBrick.SetupForInputShapesWithParameters]
//   - [MLCeilBrick.InitWithParameters]
//   - [MLCeilBrick.DebugDescription]
//   - [MLCeilBrick.Description]
//   - [MLCeilBrick.Hash]
//   - [MLCeilBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick
type MLCeilBrick struct {
	objectivec.Object
}

// MLCeilBrickFromID constructs a [MLCeilBrick] from an objc.ID.
func MLCeilBrickFromID(id objc.ID) MLCeilBrick {
	return MLCeilBrick{objectivec.Object{ID: id}}
}

// Ensure MLCeilBrick implements IMLCeilBrick.
var _ IMLCeilBrick = MLCeilBrick{}

// An interface definition for the [MLCeilBrick] class.
//
// # Methods
//
//   - [IMLCeilBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLCeilBrick.HasGPUSupport]
//   - [IMLCeilBrick.InputRanks]
//   - [IMLCeilBrick.InputShapes]
//   - [IMLCeilBrick.OutputRanks]
//   - [IMLCeilBrick.OutputShapes]
//   - [IMLCeilBrick.SetupForInputShapesWithParameters]
//   - [IMLCeilBrick.InitWithParameters]
//   - [IMLCeilBrick.DebugDescription]
//   - [IMLCeilBrick.Description]
//   - [IMLCeilBrick.Hash]
//   - [IMLCeilBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick
type IMLCeilBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLCeilBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c MLCeilBrick) Init() MLCeilBrick {
	rv := objc.Send[MLCeilBrick](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCeilBrick) Autorelease() MLCeilBrick {
	rv := objc.Send[MLCeilBrick](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCeilBrick creates a new MLCeilBrick instance.
func NewMLCeilBrick() MLCeilBrick {
	class := getMLCeilBrickClass()
	rv := objc.Send[MLCeilBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/initWithParameters:
func NewCeilBrickWithParameters(parameters objectivec.IObject) MLCeilBrick {
	instance := getMLCeilBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLCeilBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/computeOnCPUWithInputTensors:outputTensors:
func (c MLCeilBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/hasGPUSupport
func (c MLCeilBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/setupForInputShapes:withParameters:
func (c MLCeilBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/initWithParameters:
func (c MLCeilBrick) InitWithParameters(parameters objectivec.IObject) MLCeilBrick {
	rv := objc.Send[MLCeilBrick](c.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/debugDescription
func (c MLCeilBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/description
func (c MLCeilBrick) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/hash
func (c MLCeilBrick) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/inputRanks
func (c MLCeilBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/inputShapes
func (c MLCeilBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/outputRanks
func (c MLCeilBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/outputShapes
func (c MLCeilBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCeilBrick/superclass
func (c MLCeilBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}
