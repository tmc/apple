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
	rv := objc.SendIfResponds[MLSineBrick](objc.ID(mc.class), objc.Sel("alloc"))
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLSineBrick) Init() MLSineBrick {
	rv := objc.SendIfResponds[MLSineBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSineBrick) Autorelease() MLSineBrick {
	rv := objc.SendIfResponds[MLSineBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSineBrick creates a new MLSineBrick instance.
func NewMLSineBrick() MLSineBrick {
	class := getMLSineBrickClass()
	rv := objc.SendIfResponds[MLSineBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSineBrickWithParameters(parameters objectivec.IObject) MLSineBrick {
	instance := getMLSineBrickClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLSineBrickFromID(rv)
}

func (m MLSineBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLSineBrick) HasGPUSupport() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLSineBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLSineBrick) InitWithParameters(parameters objectivec.IObject) MLSineBrick {
	rv := objc.SendIfResponds[MLSineBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLSineBrick) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSineBrick) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLSineBrick) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLSineBrick) InputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSineBrick) InputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSineBrick) OutputRanks() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputRanks"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSineBrick) OutputShapes() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputShapes"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSineBrick) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
