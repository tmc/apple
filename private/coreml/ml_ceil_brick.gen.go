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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLCeilBrick) Init() MLCeilBrick {
	rv := objc.Send[MLCeilBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLCeilBrick) Autorelease() MLCeilBrick {
	rv := objc.Send[MLCeilBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCeilBrick creates a new MLCeilBrick instance.
func NewMLCeilBrick() MLCeilBrick {
	class := getMLCeilBrickClass()
	rv := objc.Send[MLCeilBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewCeilBrickWithParameters(parameters objectivec.IObject) MLCeilBrick {
	instance := getMLCeilBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLCeilBrickFromID(rv)
}

func (m MLCeilBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLCeilBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLCeilBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLCeilBrick) InitWithParameters(parameters objectivec.IObject) MLCeilBrick {
	rv := objc.Send[MLCeilBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLCeilBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCeilBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLCeilBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLCeilBrick) InputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLCeilBrick) InputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLCeilBrick) OutputRanks() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLCeilBrick) OutputShapes() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLCeilBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
