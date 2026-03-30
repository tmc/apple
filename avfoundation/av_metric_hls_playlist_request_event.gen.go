// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricHLSPlaylistRequestEvent] class.
var (
	_AVMetricHLSPlaylistRequestEventClass     AVMetricHLSPlaylistRequestEventClass
	_AVMetricHLSPlaylistRequestEventClassOnce sync.Once
)

func getAVMetricHLSPlaylistRequestEventClass() AVMetricHLSPlaylistRequestEventClass {
	_AVMetricHLSPlaylistRequestEventClassOnce.Do(func() {
		_AVMetricHLSPlaylistRequestEventClass = AVMetricHLSPlaylistRequestEventClass{class: objc.GetClass("AVMetricHLSPlaylistRequestEvent")}
	})
	return _AVMetricHLSPlaylistRequestEventClass
}

// GetAVMetricHLSPlaylistRequestEventClass returns the class object for AVMetricHLSPlaylistRequestEvent.
func GetAVMetricHLSPlaylistRequestEventClass() AVMetricHLSPlaylistRequestEventClass {
	return getAVMetricHLSPlaylistRequestEventClass()
}

type AVMetricHLSPlaylistRequestEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricHLSPlaylistRequestEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricHLSPlaylistRequestEventClass) Alloc() AVMetricHLSPlaylistRequestEvent {
	rv := objc.Send[AVMetricHLSPlaylistRequestEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents a live streaming playlist resource request.
//
// # Inspecting the event
//
//   - [AVMetricHLSPlaylistRequestEvent.IsMultivariantPlaylist]
//   - [AVMetricHLSPlaylistRequestEvent.MediaResourceRequestEvent]
//   - [AVMetricHLSPlaylistRequestEvent.MediaType]
//   - [AVMetricHLSPlaylistRequestEvent.Url]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent
type AVMetricHLSPlaylistRequestEvent struct {
	AVMetricEvent
}

// AVMetricHLSPlaylistRequestEventFromID constructs a [AVMetricHLSPlaylistRequestEvent] from an objc.ID.
//
// An event that represents a live streaming playlist resource request.
func AVMetricHLSPlaylistRequestEventFromID(id objc.ID) AVMetricHLSPlaylistRequestEvent {
	return AVMetricHLSPlaylistRequestEvent{AVMetricEvent: AVMetricEventFromID(id)}
}

// NOTE: AVMetricHLSPlaylistRequestEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricHLSPlaylistRequestEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricHLSPlaylistRequestEvent.IsMultivariantPlaylist]
//   - [IAVMetricHLSPlaylistRequestEvent.MediaResourceRequestEvent]
//   - [IAVMetricHLSPlaylistRequestEvent.MediaType]
//   - [IAVMetricHLSPlaylistRequestEvent.Url]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent
type IAVMetricHLSPlaylistRequestEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	IsMultivariantPlaylist() bool
	MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent
	MediaType() AVMediaType
	Url() foundation.INSURL
}

// Init initializes the instance.
func (m AVMetricHLSPlaylistRequestEvent) Init() AVMetricHLSPlaylistRequestEvent {
	rv := objc.Send[AVMetricHLSPlaylistRequestEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricHLSPlaylistRequestEvent) Autorelease() AVMetricHLSPlaylistRequestEvent {
	rv := objc.Send[AVMetricHLSPlaylistRequestEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricHLSPlaylistRequestEvent creates a new AVMetricHLSPlaylistRequestEvent instance.
func NewAVMetricHLSPlaylistRequestEvent() AVMetricHLSPlaylistRequestEvent {
	class := getAVMetricHLSPlaylistRequestEventClass()
	rv := objc.Send[AVMetricHLSPlaylistRequestEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/isMultivariantPlaylist
func (m AVMetricHLSPlaylistRequestEvent) IsMultivariantPlaylist() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isMultivariantPlaylist"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/mediaResourceRequestEvent
func (m AVMetricHLSPlaylistRequestEvent) MediaResourceRequestEvent() IAVMetricMediaResourceRequestEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaResourceRequestEvent"))
	return AVMetricMediaResourceRequestEventFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/mediaType
func (m AVMetricHLSPlaylistRequestEvent) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricHLSPlaylistRequestEvent/url
func (m AVMetricHLSPlaylistRequestEvent) Url() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
