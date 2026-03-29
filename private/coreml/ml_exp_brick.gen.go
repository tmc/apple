// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLExpBrick] class.
var (
	_MLExpBrickClass     MLExpBrickClass
	_MLExpBrickClassOnce sync.Once
)

func getMLExpBrickClass() MLExpBrickClass {
	_MLExpBrickClassOnce.Do(func() {
		_MLExpBrickClass = MLExpBrickClass{class: objc.GetClass("MLExpBrick")}
	})
	return _MLExpBrickClass
}

// GetMLExpBrickClass returns the class object for MLExpBrick.
func GetMLExpBrickClass() MLExpBrickClass {
	return getMLExpBrickClass()
}

type MLExpBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLExpBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLExpBrickClass) Alloc() MLExpBrick {
	rv := objc.Send[MLExpBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLExpBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLExpBrick.HasGPUSupport]
//   - [MLExpBrick.InputRanks]
//   - [MLExpBrick.InputShapes]
//   - [MLExpBrick.OutputRanks]
//   - [MLExpBrick.OutputShapes]
//   - [MLExpBrick.SetupForInputShapesWithParameters]
//   - [MLExpBrick.WithBase2]
//   - [MLExpBrick.InitWithParameters]
//   - [MLExpBrick.DebugDescription]
//   - [MLExpBrick.Description]
//   - [MLExpBrick.Hash]
//   - [MLExpBrick.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick
type MLExpBrick struct {
	objectivec.Object
}

// MLExpBrickFromID constructs a [MLExpBrick] from an objc.ID.
func MLExpBrickFromID(id objc.ID) MLExpBrick {
	return MLExpBrick{objectivec.Object{ID: id}}
}
// Ensure MLExpBrick implements IMLExpBrick.
var _ IMLExpBrick = MLExpBrick{}

// An interface definition for the [MLExpBrick] class.
//
// # Methods
//
//   - [IMLExpBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLExpBrick.HasGPUSupport]
//   - [IMLExpBrick.InputRanks]
//   - [IMLExpBrick.InputShapes]
//   - [IMLExpBrick.OutputRanks]
//   - [IMLExpBrick.OutputShapes]
//   - [IMLExpBrick.SetupForInputShapesWithParameters]
//   - [IMLExpBrick.WithBase2]
//   - [IMLExpBrick.InitWithParameters]
//   - [IMLExpBrick.DebugDescription]
//   - [IMLExpBrick.Description]
//   - [IMLExpBrick.Hash]
//   - [IMLExpBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick
type IMLExpBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	WithBase2() bool
	InitWithParameters(parameters objectivec.IObject) MLExpBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLExpBrick) Init() MLExpBrick {
	rv := objc.Send[MLExpBrick](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLExpBrick) Autorelease() MLExpBrick {
	rv := objc.Send[MLExpBrick](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLExpBrick creates a new MLExpBrick instance.
func NewMLExpBrick() MLExpBrick {
	class := getMLExpBrickClass()
	rv := objc.Send[MLExpBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/initWithParameters:
func NewExpBrickWithParameters(parameters objectivec.IObject) MLExpBrick {
	instance := getMLExpBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLExpBrickFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/computeOnCPUWithInputTensors:outputTensors:
func (e MLExpBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/hasGPUSupport
func (e MLExpBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("hasGPUSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/setupForInputShapes:withParameters:
func (e MLExpBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/initWithParameters:
func (e MLExpBrick) InitWithParameters(parameters objectivec.IObject) MLExpBrick {
	rv := objc.Send[MLExpBrick](e.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/debugDescription
func (e MLExpBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/description
func (e MLExpBrick) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/hash
func (e MLExpBrick) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/inputRanks
func (e MLExpBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/inputShapes
func (e MLExpBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/outputRanks
func (e MLExpBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/outputShapes
func (e MLExpBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/superclass
func (e MLExpBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLExpBrick/withBase2
func (e MLExpBrick) WithBase2() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("withBase2"))
	return rv
}

