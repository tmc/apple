// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMetricPlayerItemVariantSwitchStartEvent] class.
var (
	_AVMetricPlayerItemVariantSwitchStartEventClass     AVMetricPlayerItemVariantSwitchStartEventClass
	_AVMetricPlayerItemVariantSwitchStartEventClassOnce sync.Once
)

func getAVMetricPlayerItemVariantSwitchStartEventClass() AVMetricPlayerItemVariantSwitchStartEventClass {
	_AVMetricPlayerItemVariantSwitchStartEventClassOnce.Do(func() {
		_AVMetricPlayerItemVariantSwitchStartEventClass = AVMetricPlayerItemVariantSwitchStartEventClass{class: objc.GetClass("AVMetricPlayerItemVariantSwitchStartEvent")}
	})
	return _AVMetricPlayerItemVariantSwitchStartEventClass
}

// GetAVMetricPlayerItemVariantSwitchStartEventClass returns the class object for AVMetricPlayerItemVariantSwitchStartEvent.
func GetAVMetricPlayerItemVariantSwitchStartEventClass() AVMetricPlayerItemVariantSwitchStartEventClass {
	return getAVMetricPlayerItemVariantSwitchStartEventClass()
}

type AVMetricPlayerItemVariantSwitchStartEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemVariantSwitchStartEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemVariantSwitchStartEventClass) Alloc() AVMetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[AVMetricPlayerItemVariantSwitchStartEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when the player attempts a variant switch.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemVariantSwitchStartEvent.FromVariant]
//   - [AVMetricPlayerItemVariantSwitchStartEvent.ToVariant]
//   - [AVMetricPlayerItemVariantSwitchStartEvent.AudioRendition]
//   - [AVMetricPlayerItemVariantSwitchStartEvent.VideoRendition]
//   - [AVMetricPlayerItemVariantSwitchStartEvent.SubtitleRendition]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent
type AVMetricPlayerItemVariantSwitchStartEvent struct {
	AVMetricEvent
}

// AVMetricPlayerItemVariantSwitchStartEventFromID constructs a [AVMetricPlayerItemVariantSwitchStartEvent] from an objc.ID.
//
// An event that represents when the player attempts a variant switch.
func AVMetricPlayerItemVariantSwitchStartEventFromID(id objc.ID) AVMetricPlayerItemVariantSwitchStartEvent {
	return AVMetricPlayerItemVariantSwitchStartEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricPlayerItemVariantSwitchStartEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemVariantSwitchStartEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemVariantSwitchStartEvent.FromVariant]
//   - [IAVMetricPlayerItemVariantSwitchStartEvent.ToVariant]
//   - [IAVMetricPlayerItemVariantSwitchStartEvent.AudioRendition]
//   - [IAVMetricPlayerItemVariantSwitchStartEvent.VideoRendition]
//   - [IAVMetricPlayerItemVariantSwitchStartEvent.SubtitleRendition]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent
type IAVMetricPlayerItemVariantSwitchStartEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	FromVariant() IAVAssetVariant
	ToVariant() IAVAssetVariant
	AudioRendition() IAVMetricMediaRendition
	VideoRendition() IAVMetricMediaRendition
	SubtitleRendition() IAVMetricMediaRendition

	LoadedTimeRanges() []foundation.NSValue
}

// Init initializes the instance.
func (m AVMetricPlayerItemVariantSwitchStartEvent) Init() AVMetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[AVMetricPlayerItemVariantSwitchStartEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemVariantSwitchStartEvent) Autorelease() AVMetricPlayerItemVariantSwitchStartEvent {
	rv := objc.Send[AVMetricPlayerItemVariantSwitchStartEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemVariantSwitchStartEvent creates a new AVMetricPlayerItemVariantSwitchStartEvent instance.
func NewAVMetricPlayerItemVariantSwitchStartEvent() AVMetricPlayerItemVariantSwitchStartEvent {
	class := getAVMetricPlayerItemVariantSwitchStartEventClass()
	rv := objc.Send[AVMetricPlayerItemVariantSwitchStartEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/fromVariant
func (m AVMetricPlayerItemVariantSwitchStartEvent) FromVariant() IAVAssetVariant {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fromVariant"))
	return AVAssetVariantFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/toVariant
func (m AVMetricPlayerItemVariantSwitchStartEvent) ToVariant() IAVAssetVariant {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("toVariant"))
	return AVAssetVariantFromID(objc.ID(rv))
}
//
// # Discussion
// 
// Contains information corresponding to the currently selected audio
// rendition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/audioRendition
func (m AVMetricPlayerItemVariantSwitchStartEvent) AudioRendition() IAVMetricMediaRendition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("audioRendition"))
	return AVMetricMediaRenditionFromID(objc.ID(rv))
}
//
// # Discussion
// 
// Contains information corresponding to the currently selected video
// rendition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/videoRendition
func (m AVMetricPlayerItemVariantSwitchStartEvent) VideoRendition() IAVMetricMediaRendition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("videoRendition"))
	return AVMetricMediaRenditionFromID(objc.ID(rv))
}
//
// # Discussion
// 
// Contains information corresponding to the currently selected subtitle
// rendition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/subtitleRendition
func (m AVMetricPlayerItemVariantSwitchStartEvent) SubtitleRendition() IAVMetricMediaRendition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("subtitleRendition"))
	return AVMetricMediaRenditionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchStartEvent/loadedTimeRanges-3svh3
func (m AVMetricPlayerItemVariantSwitchStartEvent) LoadedTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("loadedTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

