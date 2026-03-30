// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricContentKeyRequestEvent] class.
var (
	_AVMetricContentKeyRequestEventClass     AVMetricContentKeyRequestEventClass
	_AVMetricContentKeyRequestEventClassOnce sync.Once
)

func getAVMetricContentKeyRequestEventClass() AVMetricContentKeyRequestEventClass {
	_AVMetricContentKeyRequestEventClassOnce.Do(func() {
		_AVMetricContentKeyRequestEventClass = AVMetricContentKeyRequestEventClass{class: objc.GetClass("AVMetricContentKeyRequestEvent")}
	})
	return _AVMetricContentKeyRequestEventClass
}

// GetAVMetricContentKeyRequestEventClass returns the class object for AVMetricContentKeyRequestEvent.
func GetAVMetricContentKeyRequestEventClass() AVMetricContentKeyRequestEventClass {
	return getAVMetricContentKeyRequestEventClass()
}

type AVMetricContentKeyRequestEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricContentKeyRequestEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricContentKeyRequestEventClass) Alloc() AVMetricContentKeyRequestEvent {
	rv := objc.Send[AVMetricContentKeyRequestEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents a live streaming content key resource request.
//
// # Inspecting the event
//
//   - [AVMetricContentKeyRequestEvent.ContentKeySpecifier]
//   - [AVMetricContentKeyRequestEvent.IsClientInitiated]
//   - [AVMetricContentKeyRequestEvent.MediaResourceRequestEvent]
//   - [AVMetricContentKeyRequestEvent.MediaType]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent
type AVMetricContentKeyRequestEvent struct {
	AVMetricEvent
}

// AVMetricContentKeyRequestEventFromID constructs a [AVMetricContentKeyRequestEvent] from an objc.ID.
//
// An event that represents a live streaming content key resource request.
func AVMetricContentKeyRequestEventFromID(id objc.ID) AVMetricContentKeyRequestEvent {
	return AVMetricContentKeyRequestEvent{AVMetricEvent: AVMetricEventFromID(id)}
}

// NOTE: AVMetricContentKeyRequestEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricContentKeyRequestEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricContentKeyRequestEvent.ContentKeySpecifier]
//   - [IAVMetricContentKeyRequestEvent.IsClientInitiated]
//   - [IAVMetricContentKeyRequestEvent.MediaResourceRequestEvent]
//   - [IAVMetricContentKeyRequestEvent.MediaType]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent
type IAVMetricContentKeyRequestEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	ContentKeySpecifier() IAVContentKeySpecifier
	IsClientInitiated() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() AVMediaType
}

// Init initializes the instance.
func (m AVMetricContentKeyRequestEvent) Init() AVMetricContentKeyRequestEvent {
	rv := objc.Send[AVMetricContentKeyRequestEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricContentKeyRequestEvent) Autorelease() AVMetricContentKeyRequestEvent {
	rv := objc.Send[AVMetricContentKeyRequestEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricContentKeyRequestEvent creates a new AVMetricContentKeyRequestEvent instance.
func NewAVMetricContentKeyRequestEvent() AVMetricContentKeyRequestEvent {
	class := getAVMetricContentKeyRequestEventClass()
	rv := objc.Send[AVMetricContentKeyRequestEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/contentKeySpecifier
func (m AVMetricContentKeyRequestEvent) ContentKeySpecifier() IAVContentKeySpecifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("contentKeySpecifier"))
	return AVContentKeySpecifierFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/isClientInitiated
func (m AVMetricContentKeyRequestEvent) IsClientInitiated() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isClientInitiated"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/mediaResourceRequestEvent
func (m AVMetricContentKeyRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaResourceRequestEvent"))
	return AVMetricMediaResourceRequestEventFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricContentKeyRequestEvent/mediaType
func (m AVMetricContentKeyRequestEvent) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
