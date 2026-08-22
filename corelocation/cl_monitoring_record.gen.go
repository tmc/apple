// Code generated from Apple documentation for CoreLocation. DO NOT EDIT.

package corelocation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CLMonitoringRecord] class.
var (
	_CLMonitoringRecordClass     CLMonitoringRecordClass
	_CLMonitoringRecordClassOnce sync.Once
)

func getCLMonitoringRecordClass() CLMonitoringRecordClass {
	_CLMonitoringRecordClassOnce.Do(func() {
		_CLMonitoringRecordClass = CLMonitoringRecordClass{class: objc.GetClass("CLMonitoringRecord")}
	})
	return _CLMonitoringRecordClass
}

// GetCLMonitoringRecordClass returns the class object for CLMonitoringRecord.
func GetCLMonitoringRecordClass() CLMonitoringRecordClass {
	return getCLMonitoringRecordClass()
}

type CLMonitoringRecordClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CLMonitoringRecordClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CLMonitoringRecordClass) Alloc() CLMonitoringRecord {
	rv := objc.Send[CLMonitoringRecord](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a condition and its associated information that a
// location monitor is monitoring.
//
// # Overview
//
// When handling a new [CLMonitoringEvent], the [CLMonitoringRecord] available
// for the indicated identifier from the [CLMonitor] contains the prior event.
// The [CLMonitoringRecord] updates with the new event when the handling is
// complete.
//
// # Event properties
//
//   - [CLMonitoringRecord.Condition]: The condition that the framework is monitoring events for.
//   - [CLMonitoringRecord.LastEvent]: An object that contains the specifics of the most recent event.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringRecord
//
// [CLMonitor]: https://developer.apple.com/documentation/CoreLocation/CLMonitor-2r51v
type CLMonitoringRecord struct {
	objectivec.Object
}

// CLMonitoringRecordFromID constructs a [CLMonitoringRecord] from an objc.ID.
//
// An object that represents a condition and its associated information that a
// location monitor is monitoring.
func CLMonitoringRecordFromID(id objc.ID) CLMonitoringRecord {
	return CLMonitoringRecord{objectivec.Object{ID: id}}
}

// NOTE: CLMonitoringRecord adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CLMonitoringRecord] class.
//
// # Event properties
//
//   - [ICLMonitoringRecord.Condition]: The condition that the framework is monitoring events for.
//   - [ICLMonitoringRecord.LastEvent]: An object that contains the specifics of the most recent event.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringRecord
type ICLMonitoringRecord interface {
	objectivec.IObject

	// Topic: Event properties

	// The condition that the framework is monitoring events for.
	Condition() ICLCondition
	// An object that contains the specifics of the most recent event.
	LastEvent() ICLMonitoringEvent

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m CLMonitoringRecord) Init() CLMonitoringRecord {
	rv := objc.Send[CLMonitoringRecord](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CLMonitoringRecord) Autorelease() CLMonitoringRecord {
	rv := objc.Send[CLMonitoringRecord](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCLMonitoringRecord creates a new CLMonitoringRecord instance.
func NewCLMonitoringRecord() CLMonitoringRecord {
	class := getCLMonitoringRecordClass()
	rv := objc.Send[CLMonitoringRecord](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m CLMonitoringRecord) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The condition that the framework is monitoring events for.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringRecord/condition
func (m CLMonitoringRecord) Condition() ICLCondition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("condition"))
	return CLConditionFromID(objc.ID(rv))
}

// An object that contains the specifics of the most recent event.
//
// # Discussion
//
// This includes the state, the date, and the specifics of the condition, if
// applicable.
//
// See: https://developer.apple.com/documentation/CoreLocation/CLMonitoringRecord/lastEvent
func (m CLMonitoringRecord) LastEvent() ICLMonitoringEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lastEvent"))
	return CLMonitoringEventFromID(objc.ID(rv))
}
