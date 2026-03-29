// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelExecutionSchedule] class.
var (
	_MLModelExecutionScheduleClass     MLModelExecutionScheduleClass
	_MLModelExecutionScheduleClassOnce sync.Once
)

func getMLModelExecutionScheduleClass() MLModelExecutionScheduleClass {
	_MLModelExecutionScheduleClassOnce.Do(func() {
		_MLModelExecutionScheduleClass = MLModelExecutionScheduleClass{class: objc.GetClass("MLModelExecutionSchedule")}
	})
	return _MLModelExecutionScheduleClass
}

// GetMLModelExecutionScheduleClass returns the class object for MLModelExecutionSchedule.
func GetMLModelExecutionScheduleClass() MLModelExecutionScheduleClass {
	return getMLModelExecutionScheduleClass()
}

type MLModelExecutionScheduleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelExecutionScheduleClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelExecutionScheduleClass) Alloc() MLModelExecutionSchedule {
	rv := objc.Send[MLModelExecutionSchedule](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLModelExecutionSchedule.ModelExecutionSchedule]
//   - [MLModelExecutionSchedule.SetModelExecutionSchedule]
//   - [MLModelExecutionSchedule.ModelExecutionScheduleByModelStructurePath]
//   - [MLModelExecutionSchedule.SetModelExecutionScheduleByModelStructurePath]
// See: https://developer.apple.com/documentation/CoreML/MLModelExecutionSchedule
type MLModelExecutionSchedule struct {
	objectivec.Object
}

// MLModelExecutionScheduleFromID constructs a [MLModelExecutionSchedule] from an objc.ID.
func MLModelExecutionScheduleFromID(id objc.ID) MLModelExecutionSchedule {
	return MLModelExecutionSchedule{objectivec.Object{ID: id}}
}
// Ensure MLModelExecutionSchedule implements IMLModelExecutionSchedule.
var _ IMLModelExecutionSchedule = MLModelExecutionSchedule{}

// An interface definition for the [MLModelExecutionSchedule] class.
//
// # Methods
//
//   - [IMLModelExecutionSchedule.ModelExecutionSchedule]
//   - [IMLModelExecutionSchedule.SetModelExecutionSchedule]
//   - [IMLModelExecutionSchedule.ModelExecutionScheduleByModelStructurePath]
//   - [IMLModelExecutionSchedule.SetModelExecutionScheduleByModelStructurePath]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelExecutionSchedule
type IMLModelExecutionSchedule interface {
	objectivec.IObject

	// Topic: Methods

	ModelExecutionSchedule() foundation.INSDictionary
	SetModelExecutionSchedule(value foundation.INSDictionary)
	ModelExecutionScheduleByModelStructurePath() foundation.INSDictionary
	SetModelExecutionScheduleByModelStructurePath(value foundation.INSDictionary)
}

// Init initializes the instance.
func (m MLModelExecutionSchedule) Init() MLModelExecutionSchedule {
	rv := objc.Send[MLModelExecutionSchedule](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelExecutionSchedule) Autorelease() MLModelExecutionSchedule {
	rv := objc.Send[MLModelExecutionSchedule](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelExecutionSchedule creates a new MLModelExecutionSchedule instance.
func NewMLModelExecutionSchedule() MLModelExecutionSchedule {
	class := getMLModelExecutionScheduleClass()
	rv := objc.Send[MLModelExecutionSchedule](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelExecutionSchedule/modelExecutionSchedule
func (m MLModelExecutionSchedule) ModelExecutionSchedule() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelExecutionSchedule"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelExecutionSchedule) SetModelExecutionSchedule(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelExecutionSchedule:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLModelExecutionSchedule/modelExecutionScheduleByModelStructurePath
func (m MLModelExecutionSchedule) ModelExecutionScheduleByModelStructurePath() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelExecutionScheduleByModelStructurePath"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelExecutionSchedule) SetModelExecutionScheduleByModelStructurePath(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelExecutionScheduleByModelStructurePath:"), value)
}

