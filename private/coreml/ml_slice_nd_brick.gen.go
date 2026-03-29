// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSliceNDBrick] class.
var (
	_MLSliceNDBrickClass     MLSliceNDBrickClass
	_MLSliceNDBrickClassOnce sync.Once
)

func getMLSliceNDBrickClass() MLSliceNDBrickClass {
	_MLSliceNDBrickClassOnce.Do(func() {
		_MLSliceNDBrickClass = MLSliceNDBrickClass{class: objc.GetClass("MLSliceNDBrick")}
	})
	return _MLSliceNDBrickClass
}

// GetMLSliceNDBrickClass returns the class object for MLSliceNDBrick.
func GetMLSliceNDBrickClass() MLSliceNDBrickClass {
	return getMLSliceNDBrickClass()
}

type MLSliceNDBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSliceNDBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSliceNDBrickClass) Alloc() MLSliceNDBrick {
	rv := objc.Send[MLSliceNDBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLSliceNDBrick.Begin_ids]
//   - [MLSliceNDBrick.Begin_masks]
//   - [MLSliceNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLSliceNDBrick.End_ids]
//   - [MLSliceNDBrick.End_masks]
//   - [MLSliceNDBrick.HasGPUSupport]
//   - [MLSliceNDBrick.InputRanks]
//   - [MLSliceNDBrick.InputShapes]
//   - [MLSliceNDBrick.OutputRanks]
//   - [MLSliceNDBrick.OutputShapes]
//   - [MLSliceNDBrick.Rank]
//   - [MLSliceNDBrick.SetupForInputShapesWithParameters]
//   - [MLSliceNDBrick.ShapeInfoNeeded]
//   - [MLSliceNDBrick.Strides]
//   - [MLSliceNDBrick.InitWithParameters]
//   - [MLSliceNDBrick.DebugDescription]
//   - [MLSliceNDBrick.Description]
//   - [MLSliceNDBrick.Hash]
//   - [MLSliceNDBrick.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick
type MLSliceNDBrick struct {
	objectivec.Object
}

// MLSliceNDBrickFromID constructs a [MLSliceNDBrick] from an objc.ID.
func MLSliceNDBrickFromID(id objc.ID) MLSliceNDBrick {
	return MLSliceNDBrick{objectivec.Object{ID: id}}
}
// Ensure MLSliceNDBrick implements IMLSliceNDBrick.
var _ IMLSliceNDBrick = MLSliceNDBrick{}

// An interface definition for the [MLSliceNDBrick] class.
//
// # Methods
//
//   - [IMLSliceNDBrick.Begin_ids]
//   - [IMLSliceNDBrick.Begin_masks]
//   - [IMLSliceNDBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLSliceNDBrick.End_ids]
//   - [IMLSliceNDBrick.End_masks]
//   - [IMLSliceNDBrick.HasGPUSupport]
//   - [IMLSliceNDBrick.InputRanks]
//   - [IMLSliceNDBrick.InputShapes]
//   - [IMLSliceNDBrick.OutputRanks]
//   - [IMLSliceNDBrick.OutputShapes]
//   - [IMLSliceNDBrick.Rank]
//   - [IMLSliceNDBrick.SetupForInputShapesWithParameters]
//   - [IMLSliceNDBrick.ShapeInfoNeeded]
//   - [IMLSliceNDBrick.Strides]
//   - [IMLSliceNDBrick.InitWithParameters]
//   - [IMLSliceNDBrick.DebugDescription]
//   - [IMLSliceNDBrick.Description]
//   - [IMLSliceNDBrick.Hash]
//   - [IMLSliceNDBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick
type IMLSliceNDBrick interface {
	objectivec.IObject

	// Topic: Methods

	Begin_ids() foundation.INSArray
	Begin_masks() foundation.INSArray
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	End_ids() foundation.INSArray
	End_masks() foundation.INSArray
	HasGPUSupport() bool
	InputRanks() foundation.INSArray
	InputShapes() foundation.INSArray
	OutputRanks() foundation.INSArray
	OutputShapes() foundation.INSArray
	Rank() int
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	ShapeInfoNeeded() bool
	Strides() foundation.INSArray
	InitWithParameters(parameters objectivec.IObject) MLSliceNDBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (s MLSliceNDBrick) Init() MLSliceNDBrick {
	rv := objc.Send[MLSliceNDBrick](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSliceNDBrick) Autorelease() MLSliceNDBrick {
	rv := objc.Send[MLSliceNDBrick](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSliceNDBrick creates a new MLSliceNDBrick instance.
func NewMLSliceNDBrick() MLSliceNDBrick {
	class := getMLSliceNDBrickClass()
	rv := objc.Send[MLSliceNDBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/initWithParameters:
func NewSliceNDBrickWithParameters(parameters objectivec.IObject) MLSliceNDBrick {
	instance := getMLSliceNDBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSliceNDBrickFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/computeOnCPUWithInputTensors:outputTensors:
func (s MLSliceNDBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/hasGPUSupport
func (s MLSliceNDBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasGPUSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/setupForInputShapes:withParameters:
func (s MLSliceNDBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/initWithParameters:
func (s MLSliceNDBrick) InitWithParameters(parameters objectivec.IObject) MLSliceNDBrick {
	rv := objc.Send[MLSliceNDBrick](s.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/begin_ids
func (s MLSliceNDBrick) Begin_ids() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("begin_ids"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/begin_masks
func (s MLSliceNDBrick) Begin_masks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("begin_masks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/debugDescription
func (s MLSliceNDBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/description
func (s MLSliceNDBrick) Description() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/end_ids
func (s MLSliceNDBrick) End_ids() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("end_ids"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/end_masks
func (s MLSliceNDBrick) End_masks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("end_masks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/hash
func (s MLSliceNDBrick) Hash() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/inputRanks
func (s MLSliceNDBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/inputShapes
func (s MLSliceNDBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/outputRanks
func (s MLSliceNDBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/outputShapes
func (s MLSliceNDBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/rank
func (s MLSliceNDBrick) Rank() int {
	rv := objc.Send[int](s.ID, objc.Sel("rank"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/shapeInfoNeeded
func (s MLSliceNDBrick) ShapeInfoNeeded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shapeInfoNeeded"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/strides
func (s MLSliceNDBrick) Strides() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("strides"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLSliceNDBrick/superclass
func (s MLSliceNDBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](s.ID, objc.Sel("superclass"))
	return rv
}

