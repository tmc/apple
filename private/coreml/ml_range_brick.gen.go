// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRangeBrick] class.
var (
	_MLRangeBrickClass     MLRangeBrickClass
	_MLRangeBrickClassOnce sync.Once
)

func getMLRangeBrickClass() MLRangeBrickClass {
	_MLRangeBrickClassOnce.Do(func() {
		_MLRangeBrickClass = MLRangeBrickClass{class: objc.GetClass("MLRangeBrick")}
	})
	return _MLRangeBrickClass
}

// GetMLRangeBrickClass returns the class object for MLRangeBrick.
func GetMLRangeBrickClass() MLRangeBrickClass {
	return getMLRangeBrickClass()
}

type MLRangeBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLRangeBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRangeBrickClass) Alloc() MLRangeBrick {
	rv := objc.Send[MLRangeBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLRangeBrick.ComputeDynamicOutputShape]
//   - [MLRangeBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLRangeBrick.EndValueParameter]
//   - [MLRangeBrick.HasDynamicOutputShape]
//   - [MLRangeBrick.HasGPUSupport]
//   - [MLRangeBrick.SetupForInputShapesWithParameters]
//   - [MLRangeBrick.Size]
//   - [MLRangeBrick.Start]
//   - [MLRangeBrick.StartValueParameter]
//   - [MLRangeBrick.StepSize]
//   - [MLRangeBrick.StepSizeValueParameter]
//   - [MLRangeBrick.InitWithParameters]
//   - [MLRangeBrick.DebugDescription]
//   - [MLRangeBrick.Description]
//   - [MLRangeBrick.Hash]
//   - [MLRangeBrick.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick
type MLRangeBrick struct {
	objectivec.Object
}

// MLRangeBrickFromID constructs a [MLRangeBrick] from an objc.ID.
func MLRangeBrickFromID(id objc.ID) MLRangeBrick {
	return MLRangeBrick{objectivec.Object{ID: id}}
}
// Ensure MLRangeBrick implements IMLRangeBrick.
var _ IMLRangeBrick = MLRangeBrick{}

// An interface definition for the [MLRangeBrick] class.
//
// # Methods
//
//   - [IMLRangeBrick.ComputeDynamicOutputShape]
//   - [IMLRangeBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLRangeBrick.EndValueParameter]
//   - [IMLRangeBrick.HasDynamicOutputShape]
//   - [IMLRangeBrick.HasGPUSupport]
//   - [IMLRangeBrick.SetupForInputShapesWithParameters]
//   - [IMLRangeBrick.Size]
//   - [IMLRangeBrick.Start]
//   - [IMLRangeBrick.StartValueParameter]
//   - [IMLRangeBrick.StepSize]
//   - [IMLRangeBrick.StepSizeValueParameter]
//   - [IMLRangeBrick.InitWithParameters]
//   - [IMLRangeBrick.DebugDescription]
//   - [IMLRangeBrick.Description]
//   - [IMLRangeBrick.Hash]
//   - [IMLRangeBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick
type IMLRangeBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeDynamicOutputShape(shape objectivec.IObject) objectivec.IObject
	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	EndValueParameter() float32
	HasDynamicOutputShape(shape uint64) bool
	HasGPUSupport() bool
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	Size() int
	Start() float32
	StartValueParameter() float32
	StepSize() float32
	StepSizeValueParameter() float32
	InitWithParameters(parameters objectivec.IObject) MLRangeBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (r MLRangeBrick) Init() MLRangeBrick {
	rv := objc.Send[MLRangeBrick](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRangeBrick) Autorelease() MLRangeBrick {
	rv := objc.Send[MLRangeBrick](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRangeBrick creates a new MLRangeBrick instance.
func NewMLRangeBrick() MLRangeBrick {
	class := getMLRangeBrickClass()
	rv := objc.Send[MLRangeBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/initWithParameters:
func NewRangeBrickWithParameters(parameters objectivec.IObject) MLRangeBrick {
	instance := getMLRangeBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLRangeBrickFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/computeDynamicOutputShape:
func (r MLRangeBrick) ComputeDynamicOutputShape(shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("computeDynamicOutputShape:"), shape)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/computeOnCPUWithInputTensors:outputTensors:
func (r MLRangeBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/hasDynamicOutputShape:
func (r MLRangeBrick) HasDynamicOutputShape(shape uint64) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("hasDynamicOutputShape:"), shape)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/hasGPUSupport
func (r MLRangeBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("hasGPUSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/setupForInputShapes:withParameters:
func (r MLRangeBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/initWithParameters:
func (r MLRangeBrick) InitWithParameters(parameters objectivec.IObject) MLRangeBrick {
	rv := objc.Send[MLRangeBrick](r.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/debugDescription
func (r MLRangeBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/description
func (r MLRangeBrick) Description() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/endValueParameter
func (r MLRangeBrick) EndValueParameter() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("endValueParameter"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/hash
func (r MLRangeBrick) Hash() uint64 {
	rv := objc.Send[uint64](r.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/size
func (r MLRangeBrick) Size() int {
	rv := objc.Send[int](r.ID, objc.Sel("size"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/start
func (r MLRangeBrick) Start() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("start"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/startValueParameter
func (r MLRangeBrick) StartValueParameter() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("startValueParameter"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/stepSize
func (r MLRangeBrick) StepSize() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("stepSize"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/stepSizeValueParameter
func (r MLRangeBrick) StepSizeValueParameter() float32 {
	rv := objc.Send[float32](r.ID, objc.Sel("stepSizeValueParameter"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLRangeBrick/superclass
func (r MLRangeBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](r.ID, objc.Sel("superclass"))
	return rv
}

