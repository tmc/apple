// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProgramTrainingDelta] class.
var (
	_MLProgramTrainingDeltaClass     MLProgramTrainingDeltaClass
	_MLProgramTrainingDeltaClassOnce sync.Once
)

func getMLProgramTrainingDeltaClass() MLProgramTrainingDeltaClass {
	_MLProgramTrainingDeltaClassOnce.Do(func() {
		_MLProgramTrainingDeltaClass = MLProgramTrainingDeltaClass{class: objc.GetClass("MLProgramTrainingDelta")}
	})
	return _MLProgramTrainingDeltaClass
}

// GetMLProgramTrainingDeltaClass returns the class object for MLProgramTrainingDelta.
func GetMLProgramTrainingDeltaClass() MLProgramTrainingDeltaClass {
	return getMLProgramTrainingDeltaClass()
}

type MLProgramTrainingDeltaClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramTrainingDeltaClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramTrainingDeltaClass) Alloc() MLProgramTrainingDelta {
	rv := objc.SendIfResponds[MLProgramTrainingDelta](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProgramTrainingDelta.FlattenedModelUpdate]
//   - [MLProgramTrainingDelta.InitWithFlattenedModelUpdate]
type MLProgramTrainingDelta struct {
	objectivec.Object
}

// MLProgramTrainingDeltaFromID constructs a [MLProgramTrainingDelta] from an objc.ID.
func MLProgramTrainingDeltaFromID(id objc.ID) MLProgramTrainingDelta {
	return MLProgramTrainingDelta{objectivec.Object{ID: id}}
}

// Ensure MLProgramTrainingDelta implements IMLProgramTrainingDelta.
var _ IMLProgramTrainingDelta = MLProgramTrainingDelta{}

// An interface definition for the [MLProgramTrainingDelta] class.
//
// # Methods
//
//   - [IMLProgramTrainingDelta.FlattenedModelUpdate]
//   - [IMLProgramTrainingDelta.InitWithFlattenedModelUpdate]
type IMLProgramTrainingDelta interface {
	objectivec.IObject

	// Topic: Methods

	FlattenedModelUpdate() foundation.NSData
	InitWithFlattenedModelUpdate(update objectivec.IObject) MLProgramTrainingDelta
}

// Init initializes the instance.
func (m MLProgramTrainingDelta) Init() MLProgramTrainingDelta {
	rv := objc.SendIfResponds[MLProgramTrainingDelta](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProgramTrainingDelta) Autorelease() MLProgramTrainingDelta {
	rv := objc.SendIfResponds[MLProgramTrainingDelta](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramTrainingDelta creates a new MLProgramTrainingDelta instance.
func NewMLProgramTrainingDelta() MLProgramTrainingDelta {
	class := getMLProgramTrainingDeltaClass()
	rv := objc.SendIfResponds[MLProgramTrainingDelta](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewProgramTrainingDeltaWithFlattenedModelUpdate(update objectivec.IObject) MLProgramTrainingDelta {
	instance := getMLProgramTrainingDeltaClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFlattenedModelUpdate:"), update)
	return MLProgramTrainingDeltaFromID(rv)
}

func (m MLProgramTrainingDelta) InitWithFlattenedModelUpdate(update objectivec.IObject) MLProgramTrainingDelta {
	rv := objc.SendIfResponds[MLProgramTrainingDelta](m.ID, objc.Sel("initWithFlattenedModelUpdate:"), update)
	return rv
}

func (m MLProgramTrainingDelta) FlattenedModelUpdate() foundation.NSData {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("flattenedModelUpdate"))
	return foundation.NSDataFromID(objc.ID(rv))
}
