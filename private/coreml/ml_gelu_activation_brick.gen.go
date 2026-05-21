// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLGeluActivationBrick] class.
var (
	_MLGeluActivationBrickClass     MLGeluActivationBrickClass
	_MLGeluActivationBrickClassOnce sync.Once
)

func getMLGeluActivationBrickClass() MLGeluActivationBrickClass {
	_MLGeluActivationBrickClassOnce.Do(func() {
		_MLGeluActivationBrickClass = MLGeluActivationBrickClass{class: objc.GetClass("MLGeluActivationBrick")}
	})
	return _MLGeluActivationBrickClass
}

// GetMLGeluActivationBrickClass returns the class object for MLGeluActivationBrick.
func GetMLGeluActivationBrickClass() MLGeluActivationBrickClass {
	return getMLGeluActivationBrickClass()
}

type MLGeluActivationBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLGeluActivationBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLGeluActivationBrickClass) Alloc() MLGeluActivationBrick {
	rv := objc.Send[MLGeluActivationBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLGeluActivationBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLGeluActivationBrick.HasGPUSupport]
//   - [MLGeluActivationBrick.SetupForInputShapesWithParameters]
//   - [MLGeluActivationBrick.Size]
//   - [MLGeluActivationBrick.InitWithParameters]
//   - [MLGeluActivationBrick.DebugDescription]
//   - [MLGeluActivationBrick.Description]
//   - [MLGeluActivationBrick.Hash]
//   - [MLGeluActivationBrick.Superclass]
type MLGeluActivationBrick struct {
	objectivec.Object
}

// MLGeluActivationBrickFromID constructs a [MLGeluActivationBrick] from an objc.ID.
func MLGeluActivationBrickFromID(id objc.ID) MLGeluActivationBrick {
	return MLGeluActivationBrick{objectivec.Object{ID: id}}
}

// Ensure MLGeluActivationBrick implements IMLGeluActivationBrick.
var _ IMLGeluActivationBrick = MLGeluActivationBrick{}

// An interface definition for the [MLGeluActivationBrick] class.
//
// # Methods
//
//   - [IMLGeluActivationBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLGeluActivationBrick.HasGPUSupport]
//   - [IMLGeluActivationBrick.SetupForInputShapesWithParameters]
//   - [IMLGeluActivationBrick.Size]
//   - [IMLGeluActivationBrick.InitWithParameters]
//   - [IMLGeluActivationBrick.DebugDescription]
//   - [IMLGeluActivationBrick.Description]
//   - [IMLGeluActivationBrick.Hash]
//   - [IMLGeluActivationBrick.Superclass]
type IMLGeluActivationBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	Size() uint64
	InitWithParameters(parameters objectivec.IObject) MLGeluActivationBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLGeluActivationBrick) Init() MLGeluActivationBrick {
	rv := objc.Send[MLGeluActivationBrick](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLGeluActivationBrick) Autorelease() MLGeluActivationBrick {
	rv := objc.Send[MLGeluActivationBrick](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLGeluActivationBrick creates a new MLGeluActivationBrick instance.
func NewMLGeluActivationBrick() MLGeluActivationBrick {
	class := getMLGeluActivationBrickClass()
	rv := objc.Send[MLGeluActivationBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGeluActivationBrickWithParameters(parameters objectivec.IObject) MLGeluActivationBrick {
	instance := getMLGeluActivationBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLGeluActivationBrickFromID(rv)
}

func (m MLGeluActivationBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
func (m MLGeluActivationBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasGPUSupport"))
	return rv
}
func (m MLGeluActivationBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
func (m MLGeluActivationBrick) InitWithParameters(parameters objectivec.IObject) MLGeluActivationBrick {
	rv := objc.Send[MLGeluActivationBrick](m.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

func (m MLGeluActivationBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGeluActivationBrick) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLGeluActivationBrick) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLGeluActivationBrick) Size() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("size"))
	return rv
}
func (m MLGeluActivationBrick) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
