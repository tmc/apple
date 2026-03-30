// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlayerItemInitialLikelyToKeepUpEvent] class.
var (
	_AVMetricPlayerItemInitialLikelyToKeepUpEventClass     AVMetricPlayerItemInitialLikelyToKeepUpEventClass
	_AVMetricPlayerItemInitialLikelyToKeepUpEventClassOnce sync.Once
)

func getAVMetricPlayerItemInitialLikelyToKeepUpEventClass() AVMetricPlayerItemInitialLikelyToKeepUpEventClass {
	_AVMetricPlayerItemInitialLikelyToKeepUpEventClassOnce.Do(func() {
		_AVMetricPlayerItemInitialLikelyToKeepUpEventClass = AVMetricPlayerItemInitialLikelyToKeepUpEventClass{class: objc.GetClass("AVMetricPlayerItemInitialLikelyToKeepUpEvent")}
	})
	return _AVMetricPlayerItemInitialLikelyToKeepUpEventClass
}

// GetAVMetricPlayerItemInitialLikelyToKeepUpEventClass returns the class object for AVMetricPlayerItemInitialLikelyToKeepUpEvent.
func GetAVMetricPlayerItemInitialLikelyToKeepUpEventClass() AVMetricPlayerItemInitialLikelyToKeepUpEventClass {
	return getAVMetricPlayerItemInitialLikelyToKeepUpEventClass()
}

type AVMetricPlayerItemInitialLikelyToKeepUpEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemInitialLikelyToKeepUpEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemInitialLikelyToKeepUpEventClass) Alloc() AVMetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[AVMetricPlayerItemInitialLikelyToKeepUpEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents the initial state for whether playback is likely
// to continue without stalling.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemInitialLikelyToKeepUpEvent.ContentKeyRequestEvents]
//   - [AVMetricPlayerItemInitialLikelyToKeepUpEvent.MediaSegmentRequestEvents]
//   - [AVMetricPlayerItemInitialLikelyToKeepUpEvent.PlaylistRequestEvents]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent
type AVMetricPlayerItemInitialLikelyToKeepUpEvent struct {
	AVMetricPlayerItemLikelyToKeepUpEvent
}

// AVMetricPlayerItemInitialLikelyToKeepUpEventFromID constructs a [AVMetricPlayerItemInitialLikelyToKeepUpEvent] from an objc.ID.
//
// An event that represents the initial state for whether playback is likely
// to continue without stalling.
func AVMetricPlayerItemInitialLikelyToKeepUpEventFromID(id objc.ID) AVMetricPlayerItemInitialLikelyToKeepUpEvent {
	return AVMetricPlayerItemInitialLikelyToKeepUpEvent{AVMetricPlayerItemLikelyToKeepUpEvent: AVMetricPlayerItemLikelyToKeepUpEventFromID(id)}
}

// NOTE: AVMetricPlayerItemInitialLikelyToKeepUpEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemInitialLikelyToKeepUpEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemInitialLikelyToKeepUpEvent.ContentKeyRequestEvents]
//   - [IAVMetricPlayerItemInitialLikelyToKeepUpEvent.MediaSegmentRequestEvents]
//   - [IAVMetricPlayerItemInitialLikelyToKeepUpEvent.PlaylistRequestEvents]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent
type IAVMetricPlayerItemInitialLikelyToKeepUpEvent interface {
	IAVMetricPlayerItemLikelyToKeepUpEvent

	// Topic: Inspecting the event

	ContentKeyRequestEvents() []AVMetricContentKeyRequestEvent
	MediaSegmentRequestEvents() []AVMetricHLSMediaSegmentRequestEvent
	PlaylistRequestEvents() []AVMetricHLSPlaylistRequestEvent
}

// Init initializes the instance.
func (m AVMetricPlayerItemInitialLikelyToKeepUpEvent) Init() AVMetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[AVMetricPlayerItemInitialLikelyToKeepUpEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemInitialLikelyToKeepUpEvent) Autorelease() AVMetricPlayerItemInitialLikelyToKeepUpEvent {
	rv := objc.Send[AVMetricPlayerItemInitialLikelyToKeepUpEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemInitialLikelyToKeepUpEvent creates a new AVMetricPlayerItemInitialLikelyToKeepUpEvent instance.
func NewAVMetricPlayerItemInitialLikelyToKeepUpEvent() AVMetricPlayerItemInitialLikelyToKeepUpEvent {
	class := getAVMetricPlayerItemInitialLikelyToKeepUpEventClass()
	rv := objc.Send[AVMetricPlayerItemInitialLikelyToKeepUpEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/contentKeyRequestEvents
func (m AVMetricPlayerItemInitialLikelyToKeepUpEvent) ContentKeyRequestEvents() []AVMetricContentKeyRequestEvent {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("contentKeyRequestEvents"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetricContentKeyRequestEvent {
		return AVMetricContentKeyRequestEventFromID(id)
	})
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/mediaSegmentRequestEvents
func (m AVMetricPlayerItemInitialLikelyToKeepUpEvent) MediaSegmentRequestEvents() []AVMetricHLSMediaSegmentRequestEvent {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("mediaSegmentRequestEvents"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetricHLSMediaSegmentRequestEvent {
		return AVMetricHLSMediaSegmentRequestEventFromID(id)
	})
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemInitialLikelyToKeepUpEvent/playlistRequestEvents
func (m AVMetricPlayerItemInitialLikelyToKeepUpEvent) PlaylistRequestEvents() []AVMetricHLSPlaylistRequestEvent {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("playlistRequestEvents"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetricHLSPlaylistRequestEvent {
		return AVMetricHLSPlaylistRequestEventFromID(id)
	})
}
