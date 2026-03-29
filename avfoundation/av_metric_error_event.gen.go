// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMetricErrorEvent] class.
var (
	_AVMetricErrorEventClass     AVMetricErrorEventClass
	_AVMetricErrorEventClassOnce sync.Once
)

func getAVMetricErrorEventClass() AVMetricErrorEventClass {
	_AVMetricErrorEventClassOnce.Do(func() {
		_AVMetricErrorEventClass = AVMetricErrorEventClass{class: objc.GetClass("AVMetricErrorEvent")}
	})
	return _AVMetricErrorEventClass
}

// GetAVMetricErrorEventClass returns the class object for AVMetricErrorEvent.
func GetAVMetricErrorEventClass() AVMetricErrorEventClass {
	return getAVMetricErrorEventClass()
}

type AVMetricErrorEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricErrorEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricErrorEventClass) Alloc() AVMetricErrorEvent {
	rv := objc.Send[AVMetricErrorEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a metric event when an error occurs.
//
// # Getting the error
//
//   - [AVMetricErrorEvent.Error]: Returns the error event.
//   - [AVMetricErrorEvent.DidRecover]: A Boolean value that indicates whether the error was recoverable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent
type AVMetricErrorEvent struct {
	AVMetricEvent
}

// AVMetricErrorEventFromID constructs a [AVMetricErrorEvent] from an objc.ID.
//
// An object that represents a metric event when an error occurs.
func AVMetricErrorEventFromID(id objc.ID) AVMetricErrorEvent {
	return AVMetricErrorEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricErrorEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricErrorEvent] class.
//
// # Getting the error
//
//   - [IAVMetricErrorEvent.Error]: Returns the error event.
//   - [IAVMetricErrorEvent.DidRecover]: A Boolean value that indicates whether the error was recoverable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent
type IAVMetricErrorEvent interface {
	IAVMetricEvent

	// Topic: Getting the error

	// Returns the error event.
	Error() foundation.INSError
	// A Boolean value that indicates whether the error was recoverable.
	DidRecover() bool
}

// Init initializes the instance.
func (m AVMetricErrorEvent) Init() AVMetricErrorEvent {
	rv := objc.Send[AVMetricErrorEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricErrorEvent) Autorelease() AVMetricErrorEvent {
	rv := objc.Send[AVMetricErrorEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricErrorEvent creates a new AVMetricErrorEvent instance.
func NewAVMetricErrorEvent() AVMetricErrorEvent {
	class := getAVMetricErrorEventClass()
	rv := objc.Send[AVMetricErrorEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the error event.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent/error
func (m AVMetricErrorEvent) Error() foundation.INSError {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the error was recoverable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricErrorEvent/didRecover
func (m AVMetricErrorEvent) DidRecover() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("didRecover"))
	return rv
}

