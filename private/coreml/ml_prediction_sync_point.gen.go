// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPredictionSyncPoint] class.
var (
	_MLPredictionSyncPointClass     MLPredictionSyncPointClass
	_MLPredictionSyncPointClassOnce sync.Once
)

func getMLPredictionSyncPointClass() MLPredictionSyncPointClass {
	_MLPredictionSyncPointClassOnce.Do(func() {
		_MLPredictionSyncPointClass = MLPredictionSyncPointClass{class: objc.GetClass("MLPredictionSyncPoint")}
	})
	return _MLPredictionSyncPointClass
}

// GetMLPredictionSyncPointClass returns the class object for MLPredictionSyncPoint.
func GetMLPredictionSyncPointClass() MLPredictionSyncPointClass {
	return getMLPredictionSyncPointClass()
}

type MLPredictionSyncPointClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPredictionSyncPointClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPredictionSyncPointClass) Alloc() MLPredictionSyncPoint {
	rv := objc.Send[MLPredictionSyncPoint](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLPredictionSyncPoint.Notify]
//   - [MLPredictionSyncPoint.SharedEvent]
//   - [MLPredictionSyncPoint.Value]
//   - [MLPredictionSyncPoint.InitWithSharedEventValue]
// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint
type MLPredictionSyncPoint struct {
	objectivec.Object
}

// MLPredictionSyncPointFromID constructs a [MLPredictionSyncPoint] from an objc.ID.
func MLPredictionSyncPointFromID(id objc.ID) MLPredictionSyncPoint {
	return MLPredictionSyncPoint{objectivec.Object{ID: id}}
}
// Ensure MLPredictionSyncPoint implements IMLPredictionSyncPoint.
var _ IMLPredictionSyncPoint = MLPredictionSyncPoint{}

// An interface definition for the [MLPredictionSyncPoint] class.
//
// # Methods
//
//   - [IMLPredictionSyncPoint.Notify]
//   - [IMLPredictionSyncPoint.SharedEvent]
//   - [IMLPredictionSyncPoint.Value]
//   - [IMLPredictionSyncPoint.InitWithSharedEventValue]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint
type IMLPredictionSyncPoint interface {
	objectivec.IObject

	// Topic: Methods

	Notify()
	SharedEvent() objectivec.IObject
	Value() uint64
	InitWithSharedEventValue(event objectivec.IObject, value uint64) MLPredictionSyncPoint
}

// Init initializes the instance.
func (p MLPredictionSyncPoint) Init() MLPredictionSyncPoint {
	rv := objc.Send[MLPredictionSyncPoint](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPredictionSyncPoint) Autorelease() MLPredictionSyncPoint {
	rv := objc.Send[MLPredictionSyncPoint](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionSyncPoint creates a new MLPredictionSyncPoint instance.
func NewMLPredictionSyncPoint() MLPredictionSyncPoint {
	class := getMLPredictionSyncPointClass()
	rv := objc.Send[MLPredictionSyncPoint](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint/initWithSharedEvent:value:
func NewPredictionSyncPointWithSharedEventValue(event objectivec.IObject, value uint64) MLPredictionSyncPoint {
	instance := getMLPredictionSyncPointClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSharedEvent:value:"), event, value)
	return MLPredictionSyncPointFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint/notify
func (p MLPredictionSyncPoint) Notify() {
	objc.Send[objc.ID](p.ID, objc.Sel("notify"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint/initWithSharedEvent:value:
func (p MLPredictionSyncPoint) InitWithSharedEventValue(event objectivec.IObject, value uint64) MLPredictionSyncPoint {
	rv := objc.Send[MLPredictionSyncPoint](p.ID, objc.Sel("initWithSharedEvent:value:"), event, value)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint/sharedEvent
func (p MLPredictionSyncPoint) SharedEvent() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("sharedEvent"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionSyncPoint/value
func (p MLPredictionSyncPoint) Value() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("value"))
	return rv
}

