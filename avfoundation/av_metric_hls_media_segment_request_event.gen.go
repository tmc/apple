// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMetricHLSMediaSegmentRequestEvent] class.
var (
	_AVMetricHLSMediaSegmentRequestEventClass     AVMetricHLSMediaSegmentRequestEventClass
	_AVMetricHLSMediaSegmentRequestEventClassOnce sync.Once
)

func getAVMetricHLSMediaSegmentRequestEventClass() AVMetricHLSMediaSegmentRequestEventClass {
	_AVMetricHLSMediaSegmentRequestEventClassOnce.Do(func() {
		_AVMetricHLSMediaSegmentRequestEventClass = AVMetricHLSMediaSegmentRequestEventClass{class: objc.GetClass("AVMetricHLSMediaSegmentRequestEvent")}
	})
	return _AVMetricHLSMediaSegmentRequestEventClass
}

// GetAVMetricHLSMediaSegmentRequestEventClass returns the class object for AVMetricHLSMediaSegmentRequestEvent.
func GetAVMetricHLSMediaSegmentRequestEventClass() AVMetricHLSMediaSegmentRequestEventClass {
	return getAVMetricHLSMediaSegmentRequestEventClass()
}

type AVMetricHLSMediaSegmentRequestEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricHLSMediaSegmentRequestEventClass) Alloc() AVMetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[AVMetricHLSMediaSegmentRequestEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents a live streaming media segment resource request.
//
// # Inspecting the event
//
//   - [AVMetricHLSMediaSegmentRequestEvent.ByteRange]
//   - [AVMetricHLSMediaSegmentRequestEvent.IndexFileURL]
//   - [AVMetricHLSMediaSegmentRequestEvent.IsMapSegment]
//   - [AVMetricHLSMediaSegmentRequestEvent.MediaResourceRequestEvent]
//   - [AVMetricHLSMediaSegmentRequestEvent.MediaType]
//   - [AVMetricHLSMediaSegmentRequestEvent.SegmentDuration]: Returns the duration of segment in seconds.
//   - [AVMetricHLSMediaSegmentRequestEvent.Url]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent
type AVMetricHLSMediaSegmentRequestEvent struct {
	AVMetricEvent
}

// AVMetricHLSMediaSegmentRequestEventFromID constructs a [AVMetricHLSMediaSegmentRequestEvent] from an objc.ID.
//
// An event that represents a live streaming media segment resource request.
func AVMetricHLSMediaSegmentRequestEventFromID(id objc.ID) AVMetricHLSMediaSegmentRequestEvent {
	return AVMetricHLSMediaSegmentRequestEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricHLSMediaSegmentRequestEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricHLSMediaSegmentRequestEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricHLSMediaSegmentRequestEvent.ByteRange]
//   - [IAVMetricHLSMediaSegmentRequestEvent.IndexFileURL]
//   - [IAVMetricHLSMediaSegmentRequestEvent.IsMapSegment]
//   - [IAVMetricHLSMediaSegmentRequestEvent.MediaResourceRequestEvent]
//   - [IAVMetricHLSMediaSegmentRequestEvent.MediaType]
//   - [IAVMetricHLSMediaSegmentRequestEvent.SegmentDuration]: Returns the duration of segment in seconds.
//   - [IAVMetricHLSMediaSegmentRequestEvent.Url]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent
type IAVMetricHLSMediaSegmentRequestEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	ByteRange() foundation.NSRange
	IndexFileURL() foundation.INSURL
	IsMapSegment() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() AVMediaType
	// Returns the duration of segment in seconds.
	SegmentDuration() float64
	Url() foundation.INSURL
}

// Init initializes the instance.
func (m AVMetricHLSMediaSegmentRequestEvent) Init() AVMetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[AVMetricHLSMediaSegmentRequestEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricHLSMediaSegmentRequestEvent) Autorelease() AVMetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[AVMetricHLSMediaSegmentRequestEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricHLSMediaSegmentRequestEvent creates a new AVMetricHLSMediaSegmentRequestEvent instance.
func NewAVMetricHLSMediaSegmentRequestEvent() AVMetricHLSMediaSegmentRequestEvent {
	class := getAVMetricHLSMediaSegmentRequestEventClass()
	rv := objc.Send[AVMetricHLSMediaSegmentRequestEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/byteRange
func (m AVMetricHLSMediaSegmentRequestEvent) ByteRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("byteRange"))
	return foundation.NSRange(rv)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/indexFileURL
func (m AVMetricHLSMediaSegmentRequestEvent) IndexFileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("indexFileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/isMapSegment
func (m AVMetricHLSMediaSegmentRequestEvent) IsMapSegment() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isMapSegment"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/mediaResourceRequestEvent
func (m AVMetricHLSMediaSegmentRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaResourceRequestEvent"))
	return AVMetricMediaResourceRequestEventFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/mediaType
func (m AVMetricHLSMediaSegmentRequestEvent) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// Returns the duration of segment in seconds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/segmentDuration
func (m AVMetricHLSMediaSegmentRequestEvent) SegmentDuration() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("segmentDuration"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSMediaSegmentRequestEvent/url
func (m AVMetricHLSMediaSegmentRequestEvent) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

