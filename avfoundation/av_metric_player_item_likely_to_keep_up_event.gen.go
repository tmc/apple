// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMetricPlayerItemLikelyToKeepUpEvent] class.
var (
	_AVMetricPlayerItemLikelyToKeepUpEventClass     AVMetricPlayerItemLikelyToKeepUpEventClass
	_AVMetricPlayerItemLikelyToKeepUpEventClassOnce sync.Once
)

func getAVMetricPlayerItemLikelyToKeepUpEventClass() AVMetricPlayerItemLikelyToKeepUpEventClass {
	_AVMetricPlayerItemLikelyToKeepUpEventClassOnce.Do(func() {
		_AVMetricPlayerItemLikelyToKeepUpEventClass = AVMetricPlayerItemLikelyToKeepUpEventClass{class: objc.GetClass("AVMetricPlayerItemLikelyToKeepUpEvent")}
	})
	return _AVMetricPlayerItemLikelyToKeepUpEventClass
}

// GetAVMetricPlayerItemLikelyToKeepUpEventClass returns the class object for AVMetricPlayerItemLikelyToKeepUpEvent.
func GetAVMetricPlayerItemLikelyToKeepUpEventClass() AVMetricPlayerItemLikelyToKeepUpEventClass {
	return getAVMetricPlayerItemLikelyToKeepUpEventClass()
}

type AVMetricPlayerItemLikelyToKeepUpEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemLikelyToKeepUpEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemLikelyToKeepUpEventClass) Alloc() AVMetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[AVMetricPlayerItemLikelyToKeepUpEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when playback is likely to continue without
// stalling.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemLikelyToKeepUpEvent.TimeTaken]
//   - [AVMetricPlayerItemLikelyToKeepUpEvent.Variant]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent
type AVMetricPlayerItemLikelyToKeepUpEvent struct {
	AVMetricEvent
}

// AVMetricPlayerItemLikelyToKeepUpEventFromID constructs a [AVMetricPlayerItemLikelyToKeepUpEvent] from an objc.ID.
//
// An event that represents when playback is likely to continue without
// stalling.
func AVMetricPlayerItemLikelyToKeepUpEventFromID(id objc.ID) AVMetricPlayerItemLikelyToKeepUpEvent {
	return AVMetricPlayerItemLikelyToKeepUpEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricPlayerItemLikelyToKeepUpEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemLikelyToKeepUpEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemLikelyToKeepUpEvent.TimeTaken]
//   - [IAVMetricPlayerItemLikelyToKeepUpEvent.Variant]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent
type IAVMetricPlayerItemLikelyToKeepUpEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	TimeTaken() float64
	Variant() IAVAssetVariant

	LoadedTimeRanges() []foundation.NSValue
}

// Init initializes the instance.
func (m AVMetricPlayerItemLikelyToKeepUpEvent) Init() AVMetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[AVMetricPlayerItemLikelyToKeepUpEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemLikelyToKeepUpEvent) Autorelease() AVMetricPlayerItemLikelyToKeepUpEvent {
	rv := objc.Send[AVMetricPlayerItemLikelyToKeepUpEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemLikelyToKeepUpEvent creates a new AVMetricPlayerItemLikelyToKeepUpEvent instance.
func NewAVMetricPlayerItemLikelyToKeepUpEvent() AVMetricPlayerItemLikelyToKeepUpEvent {
	class := getAVMetricPlayerItemLikelyToKeepUpEventClass()
	rv := objc.Send[AVMetricPlayerItemLikelyToKeepUpEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/timeTaken
func (m AVMetricPlayerItemLikelyToKeepUpEvent) TimeTaken() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("timeTaken"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/variant
func (m AVMetricPlayerItemLikelyToKeepUpEvent) Variant() IAVAssetVariant {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("variant"))
	return AVAssetVariantFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemLikelyToKeepUpEvent/loadedTimeRanges-960vi
func (m AVMetricPlayerItemLikelyToKeepUpEvent) LoadedTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("loadedTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

