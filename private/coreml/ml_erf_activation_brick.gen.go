// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLErfActivationBrick] class.
var (
	_MLErfActivationBrickClass     MLErfActivationBrickClass
	_MLErfActivationBrickClassOnce sync.Once
)

func getMLErfActivationBrickClass() MLErfActivationBrickClass {
	_MLErfActivationBrickClassOnce.Do(func() {
		_MLErfActivationBrickClass = MLErfActivationBrickClass{class: objc.GetClass("MLErfActivationBrick")}
	})
	return _MLErfActivationBrickClass
}

// GetMLErfActivationBrickClass returns the class object for MLErfActivationBrick.
func GetMLErfActivationBrickClass() MLErfActivationBrickClass {
	return getMLErfActivationBrickClass()
}

type MLErfActivationBrickClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLErfActivationBrickClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLErfActivationBrickClass) Alloc() MLErfActivationBrick {
	rv := objc.Send[MLErfActivationBrick](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLErfActivationBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [MLErfActivationBrick.HasGPUSupport]
//   - [MLErfActivationBrick.SetupForInputShapesWithParameters]
//   - [MLErfActivationBrick.Size]
//   - [MLErfActivationBrick.InitWithParameters]
//   - [MLErfActivationBrick.DebugDescription]
//   - [MLErfActivationBrick.Description]
//   - [MLErfActivationBrick.Hash]
//   - [MLErfActivationBrick.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick
type MLErfActivationBrick struct {
	objectivec.Object
}

// MLErfActivationBrickFromID constructs a [MLErfActivationBrick] from an objc.ID.
func MLErfActivationBrickFromID(id objc.ID) MLErfActivationBrick {
	return MLErfActivationBrick{objectivec.Object{ID: id}}
}
// Ensure MLErfActivationBrick implements IMLErfActivationBrick.
var _ IMLErfActivationBrick = MLErfActivationBrick{}

// An interface definition for the [MLErfActivationBrick] class.
//
// # Methods
//
//   - [IMLErfActivationBrick.ComputeOnCPUWithInputTensorsOutputTensors]
//   - [IMLErfActivationBrick.HasGPUSupport]
//   - [IMLErfActivationBrick.SetupForInputShapesWithParameters]
//   - [IMLErfActivationBrick.Size]
//   - [IMLErfActivationBrick.InitWithParameters]
//   - [IMLErfActivationBrick.DebugDescription]
//   - [IMLErfActivationBrick.Description]
//   - [IMLErfActivationBrick.Hash]
//   - [IMLErfActivationBrick.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick
type IMLErfActivationBrick interface {
	objectivec.IObject

	// Topic: Methods

	ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject)
	HasGPUSupport() bool
	SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject
	Size() uint64
	InitWithParameters(parameters objectivec.IObject) MLErfActivationBrick
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLErfActivationBrick) Init() MLErfActivationBrick {
	rv := objc.Send[MLErfActivationBrick](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLErfActivationBrick) Autorelease() MLErfActivationBrick {
	rv := objc.Send[MLErfActivationBrick](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLErfActivationBrick creates a new MLErfActivationBrick instance.
func NewMLErfActivationBrick() MLErfActivationBrick {
	class := getMLErfActivationBrickClass()
	rv := objc.Send[MLErfActivationBrick](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/initWithParameters:
func NewErfActivationBrickWithParameters(parameters objectivec.IObject) MLErfActivationBrick {
	instance := getMLErfActivationBrickClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:"), parameters)
	return MLErfActivationBrickFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/computeOnCPUWithInputTensors:outputTensors:
func (e MLErfActivationBrick) ComputeOnCPUWithInputTensorsOutputTensors(tensors objectivec.IObject, tensors2 objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("computeOnCPUWithInputTensors:outputTensors:"), tensors, tensors2)
}
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/hasGPUSupport
func (e MLErfActivationBrick) HasGPUSupport() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("hasGPUSupport"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/setupForInputShapes:withParameters:
func (e MLErfActivationBrick) SetupForInputShapesWithParameters(shapes objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("setupForInputShapes:withParameters:"), shapes, parameters)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/initWithParameters:
func (e MLErfActivationBrick) InitWithParameters(parameters objectivec.IObject) MLErfActivationBrick {
	rv := objc.Send[MLErfActivationBrick](e.ID, objc.Sel("initWithParameters:"), parameters)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/debugDescription
func (e MLErfActivationBrick) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/description
func (e MLErfActivationBrick) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/hash
func (e MLErfActivationBrick) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/size
func (e MLErfActivationBrick) Size() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("size"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLErfActivationBrick/superclass
func (e MLErfActivationBrick) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}

