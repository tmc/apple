// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCosineBrick] class.
var (
	_MLCosineBrickClass     MLCosineBrickClass
	_MLCosineBrickClassOnce sync.Once
)

func getMLCosineBrickClass() MLCosineBrickClass {
	_MLCosineBrickClassOnce.Do(func() {
		_MLCosineBrickClass = MLCosineBrickClass{class: objc.GetClass("MLCosineBrick")}
	})
	return _MLCosineBrickClass
}

// GetMLCosineBrickClass returns the class object for MLCosineBrick.
func GetMLCosineBrickClass() MLCosineBrickClass {
	return getMLCosineBrickClass()
}

type MLCosineBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCosineBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCosineBrickClass) Alloc() MLCosineBrick {
	rv := objc.Send[MLCosineBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLCosineBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLCosineBrick.HasGPUSupport]
//   - [MLCosineBrick.InputRanks]
//   - [MLCosineBrick.InputShapes]
//   - [MLCosineBrick.OutputRanks]
//   - [MLCosineBrick.OutputShapes]
//   - [MLCosineBrick.SetupForInputShapesWithParameters]
//   - [MLCosineBrick.InitWithParameters]
//   - [MLCosineBrick.DebugDescription]
//   - [MLCosineBrick.Description]
//   - [MLCosineBrick.Hash]
//   - [MLCosineBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick
type MLCosineBrick struct {
	objectivec.Object
}

// MLCosineBrickFromID constructs a [MLCosineBrick] from an objc.ID.
func MLCosineBrickFromID(id objc.ID) MLCosineBrick {
	return MLCosineBrick{objectivec.Object{ID: id}}
}

// Ensure MLCosineBrick implements IMLCosineBrick.
var _ IMLCosineBrick = MLCosineBrick{}

// An interface definition for the [MLCosineBrick] class.
//
// # Methods
//
//   - [IMLCosineBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLCosineBrick.HasGPUSupport]
//   - [IMLCosineBrick.InputRanks]
//   - [IMLCosineBrick.InputShapes]
//   - [IMLCosineBrick.OutputRanks]
//   - [IMLCosineBrick.OutputShapes]
//   - [IMLCosineBrick.SetupForInputShapesWithParameters]
//   - [IMLCosineBrick.InitWithParameters]
//   - [IMLCosineBrick.DebugDescription]
//   - [IMLCosineBrick.Description]
//   - [IMLCosineBrick.Hash]
//   - [IMLCosineBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick
type IMLCosineBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLCosineBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLCosineBrick) Init() MLCosineBrick {
	rv := objc.Send[MLCosineBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCosineBrick) Autorelease() MLCosineBrick {
	rv := objc.Send[MLCosineBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCosineBrick creates a new MLCosineBrick instance.
func NewMLCosineBrick() MLCosineBrick {
	class := getMLCosineBrickClass()
	rv := objc.Send[MLCosineBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/initWithParameters:
func NewCosineBrickWithParameters(parameters objectivec.IObject) MLCosineBrick {
	instance := getMLCosineBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLCosineBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/computeOnCPUWithInputTensors:outputTensors:
func (m MLCosineBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/hasGPUSupport
func (m MLCosineBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/setupForInputShapes:withParameters:
func (m MLCosineBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/initWithParameters:
func (m MLCosineBrick) InitWithParameters(parameters objectivec.IObject) MLCosineBrick {
	rv := objc.Send[MLCosineBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/debugDescription
func (m MLCosineBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/description
func (m MLCosineBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/hash
func (m MLCosineBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/inputRanks
func (m MLCosineBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/inputShapes
func (m MLCosineBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/outputRanks
func (m MLCosineBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/outputShapes
func (m MLCosineBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLCosineBrick/superclass
func (m MLCosineBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
