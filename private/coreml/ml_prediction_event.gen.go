// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
func (m MLPredictionEvent) Init() MLPredictionEvent {
	rv := objc.Send[MLPredictionEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPredictionEvent) Autorelease() MLPredictionEvent {
	rv := objc.Send[MLPredictionEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionEvent creates a new MLPredictionEvent instance.
func NewMLPredictionEvent() MLPredictionEvent {
	class := getMLPredictionEventClass()
	rv := objc.Send[MLPredictionEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLPredictionEvent) LastReportedMetric() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lastReportedMetric"))
	return objectivec.Object{ID: rv}
}
func (m MLPredictionEvent) MaybeLogPredictionEvent(event uint64) {
	objc.Send[objc.ID](m.ID, objc.Sel("maybeLogPredictionEvent:"), event)
}

func (m MLPredictionEvent) BundleIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEvent) SetBundleIdentifier(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setBundleIdentifier:"), objc.String(value))
}
func (m MLPredictionEvent) FirstPartyExecutable() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("firstPartyExecutable"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLPredictionEvent) SetFirstPartyExecutable(value foundation.NSNumber) {
	objc.Send[struct{}](m.ID, objc.Sel("setFirstPartyExecutable:"), value)
}
func (m MLPredictionEvent) ModelName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEvent) SetModelName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelName:"), objc.String(value))
}
func (m MLPredictionEvent) ModelType() foundation.NSNumber {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLPredictionEvent) SetModelType(value foundation.NSNumber) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelType:"), value)
}
