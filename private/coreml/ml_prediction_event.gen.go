// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPredictionEvent] class.
var (
	_MLPredictionEventClass     MLPredictionEventClass
	_MLPredictionEventClassOnce sync.Once
)

func getMLPredictionEventClass() MLPredictionEventClass {
	_MLPredictionEventClassOnce.Do(func() {
		_MLPredictionEventClass = MLPredictionEventClass{class: objc.GetClass("MLPredictionEvent")}
	})
	return _MLPredictionEventClass
}

// GetMLPredictionEventClass returns the class object for MLPredictionEvent.
func GetMLPredictionEventClass() MLPredictionEventClass {
	return getMLPredictionEventClass()
}

type MLPredictionEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPredictionEventClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPredictionEventClass) Alloc() MLPredictionEvent {
	rv := objc.Send[MLPredictionEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLPredictionEvent.BundleIdentifier]
//   - [MLPredictionEvent.SetBundleIdentifier]
//   - [MLPredictionEvent.FirstPartyExecutable]
//   - [MLPredictionEvent.SetFirstPartyExecutable]
//   - [MLPredictionEvent.LastReportedMetric]
//   - [MLPredictionEvent.MaybeLogPredictionEvent]
//   - [MLPredictionEvent.ModelName]
//   - [MLPredictionEvent.SetModelName]
//   - [MLPredictionEvent.ModelType]
//   - [MLPredictionEvent.SetModelType]
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent
type MLPredictionEvent struct {
	objectivec.Object
}

// MLPredictionEventFromID constructs a [MLPredictionEvent] from an objc.ID.
func MLPredictionEventFromID(id objc.ID) MLPredictionEvent {
	return MLPredictionEvent{objectivec.Object{ID: id}}
}
// Ensure MLPredictionEvent implements IMLPredictionEvent.
var _ IMLPredictionEvent = MLPredictionEvent{}

// An interface definition for the [MLPredictionEvent] class.
//
// # Methods
//
//   - [IMLPredictionEvent.BundleIdentifier]
//   - [IMLPredictionEvent.SetBundleIdentifier]
//   - [IMLPredictionEvent.FirstPartyExecutable]
//   - [IMLPredictionEvent.SetFirstPartyExecutable]
//   - [IMLPredictionEvent.LastReportedMetric]
//   - [IMLPredictionEvent.MaybeLogPredictionEvent]
//   - [IMLPredictionEvent.ModelName]
//   - [IMLPredictionEvent.SetModelName]
//   - [IMLPredictionEvent.ModelType]
//   - [IMLPredictionEvent.SetModelType]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent
type IMLPredictionEvent interface {
	objectivec.IObject

	// Topic: Methods

	BundleIdentifier() string
	SetBundleIdentifier(value string)
	FirstPartyExecutable() foundation.NSNumber
	SetFirstPartyExecutable(value foundation.NSNumber)
	LastReportedMetric() objectivec.IObject
	MaybeLogPredictionEvent(event uint64)
	ModelName() string
	SetModelName(value string)
	ModelType() foundation.NSNumber
	SetModelType(value foundation.NSNumber)
}

// Init initializes the instance.
func (p MLPredictionEvent) Init() MLPredictionEvent {
	rv := objc.Send[MLPredictionEvent](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPredictionEvent) Autorelease() MLPredictionEvent {
	rv := objc.Send[MLPredictionEvent](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionEvent creates a new MLPredictionEvent instance.
func NewMLPredictionEvent() MLPredictionEvent {
	class := getMLPredictionEventClass()
	rv := objc.Send[MLPredictionEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent/lastReportedMetric
func (p MLPredictionEvent) LastReportedMetric() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("lastReportedMetric"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent/maybeLogPredictionEvent:
func (p MLPredictionEvent) MaybeLogPredictionEvent(event uint64) {
	objc.Send[objc.ID](p.ID, objc.Sel("maybeLogPredictionEvent:"), event)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent/bundleIdentifier
func (p MLPredictionEvent) BundleIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (p MLPredictionEvent) SetBundleIdentifier(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent/firstPartyExecutable
func (p MLPredictionEvent) FirstPartyExecutable() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("firstPartyExecutable"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (p MLPredictionEvent) SetFirstPartyExecutable(value foundation.NSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setFirstPartyExecutable:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent/modelName
func (p MLPredictionEvent) ModelName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (p MLPredictionEvent) SetModelName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setModelName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEvent/modelType
func (p MLPredictionEvent) ModelType() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (p MLPredictionEvent) SetModelType(value foundation.NSNumber) {
	objc.Send[struct{}](p.ID, objc.Sel("setModelType:"), value)
}

