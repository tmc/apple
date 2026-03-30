// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSineBrick] class.
var (
	_MLSineBrickClass     MLSineBrickClass
	_MLSineBrickClassOnce sync.Once
)

func getMLSineBrickClass() MLSineBrickClass {
	_MLSineBrickClassOnce.Do(func() {
		_MLSineBrickClass = MLSineBrickClass{class: objc.GetClass("MLSineBrick")}
	})
	return _MLSineBrickClass
}

// GetMLSineBrickClass returns the class object for MLSineBrick.
func GetMLSineBrickClass() MLSineBrickClass {
	return getMLSineBrickClass()
}

type MLSineBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSineBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSineBrickClass) Alloc() MLSineBrick {
	rv := objc.Send[MLSineBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSineBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSineBrick.HasGPUSupport]
//   - [MLSineBrick.InputRanks]
//   - [MLSineBrick.InputShapes]
//   - [MLSineBrick.OutputRanks]
//   - [MLSineBrick.OutputShapes]
//   - [MLSineBrick.SetupForInputShapesWithParameters]
//   - [MLSineBrick.InitWithParameters]
//   - [MLSineBrick.DebugDescription]
//   - [MLSineBrick.Description]
//   - [MLSineBrick.Hash]
//   - [MLSineBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSineBrick
type MLSineBrick struct {
	objectivec.Object
}

// MLSineBrickFromID constructs a [MLSineBrick] from an objc.ID.
func MLSineBrickFromID(id objc.ID) MLSineBrick {
	return MLSineBrick{objectivec.Object{ID: id}}
}

// Ensure MLSineBrick implements IMLSineBrick.
var _ IMLSineBrick = MLSineBrick{}

// An interface definition for the [MLSineBrick] class.
//
// # Methods
//
//   - [IMLSineBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSineBrick.HasGPUSupport]
//   - [IMLSineBrick.InputRanks]
//   - [IMLSineBrick.InputShapes]
//   - [IMLSineBrick.OutputRanks]
//   - [IMLSineBrick.OutputShapes]
//   - [IMLSineBrick.SetupForInputShapesWithParameters]
//   - [IMLSineBrick.InitWithParameters]
//   - [IMLSineBrick.DebugDescription]
//   - [IMLSineBrick.Description]
//   - [IMLSineBrick.Hash]
//   - [IMLSineBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSineBrick
type IMLSineBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	InitWithParameters(parameters objectivec.IObject) MLSineBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSineBrick) Init() MLSineBrick {
	rv := objc.Send[MLSineBrick](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSineBrick) Autorelease() MLSineBrick {
	rv := objc.Send[MLSineBrick](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSineBrick creates a new MLSineBrick instance.
func NewMLSineBrick() MLSineBrick {
	class := getMLSineBrickClass()
	rv := objc.Send[MLSineBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/initWithParameters:
func NewSineBrickWithParameters(parameters objectivec.IObject) MLSineBrick {
	instance := getMLSineBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSineBrickFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/computeOnCPUWithInputTensors:outputTensors:
func (s MLSineBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/hasGPUSupport
func (s MLSineBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasGPUSupport"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/setupForInputShapes:withParameters:
func (s MLSineBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/initWithParameters:
func (s MLSineBrick) InitWithParameters(parameters objectivec.IObject) MLSineBrick {
	rv := objc.Send[MLSineBrick](s.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/debugDescription
func (s MLSineBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/description
func (s MLSineBrick) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/hash
func (s MLSineBrick) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/inputRanks
func (s MLSineBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/inputShapes
func (s MLSineBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/outputRanks
func (s MLSineBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/outputShapes
func (s MLSineBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLSineBrick/superclass
func (s MLSineBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}
