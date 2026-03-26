// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMetricEvent] class.
var (
	_AVMetricEventClass     AVMetricEventClass
	_AVMetricEventClassOnce sync.Once
)

func getAVMetricEventClass() AVMetricEventClass {
	_AVMetricEventClassOnce.Do(func() {
		_AVMetricEventClass = AVMetricEventClass{class: objc.GetClass("AVMetricEvent")}
	})
	return _AVMetricEventClass
}

// GetAVMetricEventClass returns the class object for AVMetricEvent.
func GetAVMetricEventClass() AVMetricEventClass {
	return getAVMetricEventClass()
}

type AVMetricEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricEventClass) Alloc() AVMetricEvent {
	rv := objc.Send[AVMetricEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A base class that represents a metric event.
//
// # Inspecting an event
//
//   - [AVMetricEvent.Date]
//   - [AVMetricEvent.MediaTime]
//   - [AVMetricEvent.SessionID]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent
type AVMetricEvent struct {
	objectivec.Object
}

// AVMetricEventFromID constructs a [AVMetricEvent] from an objc.ID.
//
// A base class that represents a metric event.
func AVMetricEventFromID(id objc.ID) AVMetricEvent {
	return AVMetricEvent{objectivec.Object{ID: id}}
}
// NOTE: AVMetricEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricEvent] class.
//
// # Inspecting an event
//
//   - [IAVMetricEvent.Date]
//   - [IAVMetricEvent.MediaTime]
//   - [IAVMetricEvent.SessionID]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent
type IAVMetricEvent interface {
	objectivec.IObject

	// Topic: Inspecting an event

	Date() foundation.INSDate
	MediaTime() objectivec.IObject
	SessionID() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m AVMetricEvent) Init() AVMetricEvent {
	rv := objc.Send[AVMetricEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricEvent) Autorelease() AVMetricEvent {
	rv := objc.Send[AVMetricEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricEvent creates a new AVMetricEvent instance.
func NewAVMetricEvent() AVMetricEvent {
	class := getAVMetricEventClass()
	rv := objc.Send[AVMetricEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m AVMetricEvent) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/date
func (m AVMetricEvent) Date() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("date"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/mediaTime
func (m AVMetricEvent) MediaTime() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaTime"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/sessionID
func (m AVMetricEvent) SessionID() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sessionID"))
	return foundation.NSStringFromID(rv).String()
}

