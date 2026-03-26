// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMetricMediaResourceRequestEvent] class.
var (
	_AVMetricMediaResourceRequestEventClass     AVMetricMediaResourceRequestEventClass
	_AVMetricMediaResourceRequestEventClassOnce sync.Once
)

func getAVMetricMediaResourceRequestEventClass() AVMetricMediaResourceRequestEventClass {
	_AVMetricMediaResourceRequestEventClassOnce.Do(func() {
		_AVMetricMediaResourceRequestEventClass = AVMetricMediaResourceRequestEventClass{class: objc.GetClass("AVMetricMediaResourceRequestEvent")}
	})
	return _AVMetricMediaResourceRequestEventClass
}

// GetAVMetricMediaResourceRequestEventClass returns the class object for AVMetricMediaResourceRequestEvent.
func GetAVMetricMediaResourceRequestEventClass() AVMetricMediaResourceRequestEventClass {
	return getAVMetricMediaResourceRequestEventClass()
}

type AVMetricMediaResourceRequestEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricMediaResourceRequestEventClass) Alloc() AVMetricMediaResourceRequestEvent {
	rv := objc.Send[AVMetricMediaResourceRequestEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents a media resource request.
//
// # Inspecting the event
//
//   - [AVMetricMediaResourceRequestEvent.ByteRange]
//   - [AVMetricMediaResourceRequestEvent.ErrorEvent]
//   - [AVMetricMediaResourceRequestEvent.NetworkTransactionMetrics]
//   - [AVMetricMediaResourceRequestEvent.RequestEndTime]
//   - [AVMetricMediaResourceRequestEvent.RequestStartTime]
//   - [AVMetricMediaResourceRequestEvent.ResponseEndTime]
//   - [AVMetricMediaResourceRequestEvent.ResponseStartTime]
//   - [AVMetricMediaResourceRequestEvent.ServerAddress]
//   - [AVMetricMediaResourceRequestEvent.Url]
//   - [AVMetricMediaResourceRequestEvent.ReadFromCache]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent
type AVMetricMediaResourceRequestEvent struct {
	AVMetricEvent
}

// AVMetricMediaResourceRequestEventFromID constructs a [AVMetricMediaResourceRequestEvent] from an objc.ID.
//
// An event that represents a media resource request.
func AVMetricMediaResourceRequestEventFromID(id objc.ID) AVMetricMediaResourceRequestEvent {
	return AVMetricMediaResourceRequestEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricMediaResourceRequestEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricMediaResourceRequestEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricMediaResourceRequestEvent.ByteRange]
//   - [IAVMetricMediaResourceRequestEvent.ErrorEvent]
//   - [IAVMetricMediaResourceRequestEvent.NetworkTransactionMetrics]
//   - [IAVMetricMediaResourceRequestEvent.RequestEndTime]
//   - [IAVMetricMediaResourceRequestEvent.RequestStartTime]
//   - [IAVMetricMediaResourceRequestEvent.ResponseEndTime]
//   - [IAVMetricMediaResourceRequestEvent.ResponseStartTime]
//   - [IAVMetricMediaResourceRequestEvent.ServerAddress]
//   - [IAVMetricMediaResourceRequestEvent.Url]
//   - [IAVMetricMediaResourceRequestEvent.ReadFromCache]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent
type IAVMetricMediaResourceRequestEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	ByteRange() foundation.NSRange
	ErrorEvent() IAVMetricErrorEvent
	NetworkTransactionMetrics() foundation.NSURLSessionTaskMetrics
	RequestEndTime() foundation.INSDate
	RequestStartTime() foundation.INSDate
	ResponseEndTime() foundation.INSDate
	ResponseStartTime() foundation.INSDate
	ServerAddress() string
	Url() foundation.INSURL
	ReadFromCache() bool
}

// Init initializes the instance.
func (m AVMetricMediaResourceRequestEvent) Init() AVMetricMediaResourceRequestEvent {
	rv := objc.Send[AVMetricMediaResourceRequestEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricMediaResourceRequestEvent) Autorelease() AVMetricMediaResourceRequestEvent {
	rv := objc.Send[AVMetricMediaResourceRequestEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricMediaResourceRequestEvent creates a new AVMetricMediaResourceRequestEvent instance.
func NewAVMetricMediaResourceRequestEvent() AVMetricMediaResourceRequestEvent {
	class := getAVMetricMediaResourceRequestEventClass()
	rv := objc.Send[AVMetricMediaResourceRequestEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/byteRange
func (m AVMetricMediaResourceRequestEvent) ByteRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("byteRange"))
	return foundation.NSRange(rv)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/errorEvent
func (m AVMetricMediaResourceRequestEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("errorEvent"))
	return AVMetricErrorEventFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/networkTransactionMetrics
func (m AVMetricMediaResourceRequestEvent) NetworkTransactionMetrics() foundation.NSURLSessionTaskMetrics {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("networkTransactionMetrics"))
	return foundation.NSURLSessionTaskMetricsFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/requestEndTime
func (m AVMetricMediaResourceRequestEvent) RequestEndTime() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("requestEndTime"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/requestStartTime
func (m AVMetricMediaResourceRequestEvent) RequestStartTime() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("requestStartTime"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/responseEndTime
func (m AVMetricMediaResourceRequestEvent) ResponseEndTime() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("responseEndTime"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/responseStartTime
func (m AVMetricMediaResourceRequestEvent) ResponseStartTime() foundation.INSDate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("responseStartTime"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/serverAddress
func (m AVMetricMediaResourceRequestEvent) ServerAddress() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("serverAddress"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/url
func (m AVMetricMediaResourceRequestEvent) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricMediaResourceRequestEvent/wasReadFromCache
func (m AVMetricMediaResourceRequestEvent) ReadFromCache() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("wasReadFromCache"))
	return rv
}

