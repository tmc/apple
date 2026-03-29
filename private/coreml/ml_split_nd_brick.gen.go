// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSplitNDBrick] class.
var (
	_MLSplitNDBrickClass     MLSplitNDBrickClass
	_MLSplitNDBrickClassOnce sync.Once
)

func getMLSplitNDBrickClass() MLSplitNDBrickClass {
	_MLSplitNDBrickClassOnce.Do(func() {
		_MLSplitNDBrickClass = MLSplitNDBrickClass{class: objc.GetClass("MLSplitNDBrick")}
	})
	return _MLSplitNDBrickClass
}

// GetMLSplitNDBrickClass returns the class object for MLSplitNDBrick.
func GetMLSplitNDBrickClass() MLSplitNDBrickClass {
	return getMLSplitNDBrickClass()
}

type MLSplitNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSplitNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSplitNDBrickClass) Alloc() MLSplitNDBrick {
	rv := objc.Send[MLSplitNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSplitNDBrick.Axis]
//   - [MLSplitNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSplitNDBrick.HasGPUSupport]
//   - [MLSplitNDBrick.InputRanks]
//   - [MLSplitNDBrick.InputShapes]
//   - [MLSplitNDBrick.NumSplits]
//   - [MLSplitNDBrick.OutputRanks]
//   - [MLSplitNDBrick.OutputShapes]
//   - [MLSplitNDBrick.SetupForInputShapesWithParameters]
//   - [MLSplitNDBrick.ShapeInfoNeeded]
//   - [MLSplitNDBrick.SplitSizes]
//   - [MLSplitNDBrick.InitWithParameters]
//   - [MLSplitNDBrick.DebugDescription]
//   - [MLSplitNDBrick.Description]
//   - [MLSplitNDBrick.Hash]
//   - [MLSplitNDBrick.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick
type MLSplitNDBrick struct {
	objectivec.Object
}

// MLSplitNDBrickFromID constructs a [MLSplitNDBrick] from an objc.ID.
func MLSplitNDBrickFromID(id objc.ID) MLSplitNDBrick {
	return MLSplitNDBrick{objectivec.Object{ID: id}}
}
// Ensure MLSplitNDBrick implements IMLSplitNDBrick.
var _ IMLSplitNDBrick = MLSplitNDBrick{}

// An interface definition for the [MLSplitNDBrick] class.
//
// # Methods
//
//   - [IMLSplitNDBrick.Axis]
//   - [IMLSplitNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSplitNDBrick.HasGPUSupport]
//   - [IMLSplitNDBrick.InputRanks]
//   - [IMLSplitNDBrick.InputShapes]
//   - [IMLSplitNDBrick.NumSplits]
//   - [IMLSplitNDBrick.OutputRanks]
//   - [IMLSplitNDBrick.OutputShapes]
//   - [IMLSplitNDBrick.SetupForInputShapesWithParameters]
//   - [IMLSplitNDBrick.ShapeInfoNeeded]
//   - [IMLSplitNDBrick.SplitSizes]
//   - [IMLSplitNDBrick.InitWithParameters]
//   - [IMLSplitNDBrick.DebugDescription]
//   - [IMLSplitNDBrick.Description]
//   - [IMLSplitNDBrick.Hash]
//   - [IMLSplitNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick
type IMLSplitNDBrick interface {
	objectivec.IObject

	// Topic: Methods

	Axis() foundation.NSNumber
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	NumSplits() foundation.NSNumber
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	SplitSizes() foundation.INSArray
	InitWithParameters(parameters objectivec.IObject) MLSplitNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSplitNDBrick) Init() MLSplitNDBrick {
	rv := objc.Send[MLSplitNDBrick](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSplitNDBrick) Autorelease() MLSplitNDBrick {
	rv := objc.Send[MLSplitNDBrick](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSplitNDBrick creates a new MLSplitNDBrick instance.
func NewMLSplitNDBrick() MLSplitNDBrick {
	class := getMLSplitNDBrickClass()
	rv := objc.Send[MLSplitNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/initWithParameters:
func NewSplitNDBrickWithParameters(parameters objectivec.IObject) MLSplitNDBrick {
	instance := getMLSplitNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSplitNDBrickFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/computeOnCPUWithInputTensors:outputTensors:
func (s MLSplitNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/hasGPUSupport
func (s MLSplitNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasGPUSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/setupForInputShapes:withParameters:
func (s MLSplitNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/initWithParameters:
func (s MLSplitNDBrick) InitWithParameters(parameters objectivec.IObject) MLSplitNDBrick {
	rv := objc.Send[MLSplitNDBrick](s.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/axis
func (s MLSplitNDBrick) Axis() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("axis"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/debugDescription
func (s MLSplitNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/description
func (s MLSplitNDBrick) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/hash
func (s MLSplitNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/inputRanks
func (s MLSplitNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/inputShapes
func (s MLSplitNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/numSplits
func (s MLSplitNDBrick) NumSplits() foundation.NSNumber {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("numSplits"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/outputRanks
func (s MLSplitNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/outputShapes
func (s MLSplitNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/shapeInfoNeeded
func (s MLSplitNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/splitSizes
func (s MLSplitNDBrick) SplitSizes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("splitSizes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSplitNDBrick/superclass
func (s MLSplitNDBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

