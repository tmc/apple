// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlayerItemPlaybackSummaryEvent] class.
var (
	_AVMetricPlayerItemPlaybackSummaryEventClass     AVMetricPlayerItemPlaybackSummaryEventClass
	_AVMetricPlayerItemPlaybackSummaryEventClassOnce sync.Once
)

func getAVMetricPlayerItemPlaybackSummaryEventClass() AVMetricPlayerItemPlaybackSummaryEventClass {
	_AVMetricPlayerItemPlaybackSummaryEventClassOnce.Do(func() {
		_AVMetricPlayerItemPlaybackSummaryEventClass = AVMetricPlayerItemPlaybackSummaryEventClass{class: objc.GetClass("AVMetricPlayerItemPlaybackSummaryEvent")}
	})
	return _AVMetricPlayerItemPlaybackSummaryEventClass
}

// GetAVMetricPlayerItemPlaybackSummaryEventClass returns the class object for AVMetricPlayerItemPlaybackSummaryEvent.
func GetAVMetricPlayerItemPlaybackSummaryEventClass() AVMetricPlayerItemPlaybackSummaryEventClass {
	return getAVMetricPlayerItemPlaybackSummaryEventClass()
}

type AVMetricPlayerItemPlaybackSummaryEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemPlaybackSummaryEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemPlaybackSummaryEventClass) Alloc() AVMetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[AVMetricPlayerItemPlaybackSummaryEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents the combined metrics for the entire playback
// session.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemPlaybackSummaryEvent.ErrorEvent]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.MediaResourceRequestCount]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.PlaybackDuration]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.RecoverableErrorCount]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.StallCount]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.TimeSpentInInitialStartup]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.TimeSpentRecoveringFromStall]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.TimeWeightedAverageBitrate]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.TimeWeightedPeakBitrate]
//   - [AVMetricPlayerItemPlaybackSummaryEvent.VariantSwitchCount]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent
type AVMetricPlayerItemPlaybackSummaryEvent struct {
	AVMetricEvent
}

// AVMetricPlayerItemPlaybackSummaryEventFromID constructs a [AVMetricPlayerItemPlaybackSummaryEvent] from an objc.ID.
//
// An event that represents the combined metrics for the entire playback
// session.
func AVMetricPlayerItemPlaybackSummaryEventFromID(id objc.ID) AVMetricPlayerItemPlaybackSummaryEvent {
	return AVMetricPlayerItemPlaybackSummaryEvent{AVMetricEvent: AVMetricEventFromID(id)}
}

// NOTE: AVMetricPlayerItemPlaybackSummaryEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemPlaybackSummaryEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.ErrorEvent]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.MediaResourceRequestCount]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.PlaybackDuration]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.RecoverableErrorCount]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.StallCount]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.TimeSpentInInitialStartup]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.TimeSpentRecoveringFromStall]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.TimeWeightedAverageBitrate]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.TimeWeightedPeakBitrate]
//   - [IAVMetricPlayerItemPlaybackSummaryEvent.VariantSwitchCount]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent
type IAVMetricPlayerItemPlaybackSummaryEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	ErrorEvent() IAVMetricErrorEvent
	MediaResourceRequestCount() int
	PlaybackDuration() int
	RecoverableErrorCount() int
	StallCount() int
	TimeSpentInInitialStartup() float64
	TimeSpentRecoveringFromStall() float64
	TimeWeightedAverageBitrate() int
	TimeWeightedPeakBitrate() int
	VariantSwitchCount() int
}

// Init initializes the instance.
func (m AVMetricPlayerItemPlaybackSummaryEvent) Init() AVMetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[AVMetricPlayerItemPlaybackSummaryEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemPlaybackSummaryEvent) Autorelease() AVMetricPlayerItemPlaybackSummaryEvent {
	rv := objc.Send[AVMetricPlayerItemPlaybackSummaryEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemPlaybackSummaryEvent creates a new AVMetricPlayerItemPlaybackSummaryEvent instance.
func NewAVMetricPlayerItemPlaybackSummaryEvent() AVMetricPlayerItemPlaybackSummaryEvent {
	class := getAVMetricPlayerItemPlaybackSummaryEventClass()
	rv := objc.Send[AVMetricPlayerItemPlaybackSummaryEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/errorEvent
func (m AVMetricPlayerItemPlaybackSummaryEvent) ErrorEvent() IAVMetricErrorEvent {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("errorEvent"))
	return AVMetricErrorEventFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/mediaResourceRequestCount
func (m AVMetricPlayerItemPlaybackSummaryEvent) MediaResourceRequestCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("mediaResourceRequestCount"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/playbackDuration
func (m AVMetricPlayerItemPlaybackSummaryEvent) PlaybackDuration() int {
	rv := objc.Send[int](m.ID, objc.Sel("playbackDuration"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/recoverableErrorCount
func (m AVMetricPlayerItemPlaybackSummaryEvent) RecoverableErrorCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("recoverableErrorCount"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/stallCount
func (m AVMetricPlayerItemPlaybackSummaryEvent) StallCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("stallCount"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeSpentInInitialStartup
func (m AVMetricPlayerItemPlaybackSummaryEvent) TimeSpentInInitialStartup() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("timeSpentInInitialStartup"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeSpentRecoveringFromStall
func (m AVMetricPlayerItemPlaybackSummaryEvent) TimeSpentRecoveringFromStall() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("timeSpentRecoveringFromStall"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeWeightedAverageBitrate
func (m AVMetricPlayerItemPlaybackSummaryEvent) TimeWeightedAverageBitrate() int {
	rv := objc.Send[int](m.ID, objc.Sel("timeWeightedAverageBitrate"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/timeWeightedPeakBitrate
func (m AVMetricPlayerItemPlaybackSummaryEvent) TimeWeightedPeakBitrate() int {
	rv := objc.Send[int](m.ID, objc.Sel("timeWeightedPeakBitrate"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemPlaybackSummaryEvent/variantSwitchCount
func (m AVMetricPlayerItemPlaybackSummaryEvent) VariantSwitchCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("variantSwitchCount"))
	return rv
}
