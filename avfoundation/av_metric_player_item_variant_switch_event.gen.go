// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMetricPlayerItemVariantSwitchEvent] class.
var (
	_AVMetricPlayerItemVariantSwitchEventClass     AVMetricPlayerItemVariantSwitchEventClass
	_AVMetricPlayerItemVariantSwitchEventClassOnce sync.Once
)

func getAVMetricPlayerItemVariantSwitchEventClass() AVMetricPlayerItemVariantSwitchEventClass {
	_AVMetricPlayerItemVariantSwitchEventClassOnce.Do(func() {
		_AVMetricPlayerItemVariantSwitchEventClass = AVMetricPlayerItemVariantSwitchEventClass{class: objc.GetClass("AVMetricPlayerItemVariantSwitchEvent")}
	})
	return _AVMetricPlayerItemVariantSwitchEventClass
}

// GetAVMetricPlayerItemVariantSwitchEventClass returns the class object for AVMetricPlayerItemVariantSwitchEvent.
func GetAVMetricPlayerItemVariantSwitchEventClass() AVMetricPlayerItemVariantSwitchEventClass {
	return getAVMetricPlayerItemVariantSwitchEventClass()
}

type AVMetricPlayerItemVariantSwitchEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemVariantSwitchEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemVariantSwitchEventClass) Alloc() AVMetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[AVMetricPlayerItemVariantSwitchEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when the player completes a variant switch.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemVariantSwitchEvent.DidSucceed]
//   - [AVMetricPlayerItemVariantSwitchEvent.FromVariant]
//   - [AVMetricPlayerItemVariantSwitchEvent.ToVariant]
//   - [AVMetricPlayerItemVariantSwitchEvent.AudioRendition]: Represents the currently selected video rendition’s identifiers.
//   - [AVMetricPlayerItemVariantSwitchEvent.VideoRendition]: Represents the currently selected video rendition’s identifiers.
//   - [AVMetricPlayerItemVariantSwitchEvent.SubtitleRendition]: Represents the currently selected audio rendition’s identifiers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent
type AVMetricPlayerItemVariantSwitchEvent struct {
	AVMetricEvent
}

// AVMetricPlayerItemVariantSwitchEventFromID constructs a [AVMetricPlayerItemVariantSwitchEvent] from an objc.ID.
//
// An event that represents when the player completes a variant switch.
func AVMetricPlayerItemVariantSwitchEventFromID(id objc.ID) AVMetricPlayerItemVariantSwitchEvent {
	return AVMetricPlayerItemVariantSwitchEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricPlayerItemVariantSwitchEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemVariantSwitchEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemVariantSwitchEvent.DidSucceed]
//   - [IAVMetricPlayerItemVariantSwitchEvent.FromVariant]
//   - [IAVMetricPlayerItemVariantSwitchEvent.ToVariant]
//   - [IAVMetricPlayerItemVariantSwitchEvent.AudioRendition]: Represents the currently selected video rendition’s identifiers.
//   - [IAVMetricPlayerItemVariantSwitchEvent.VideoRendition]: Represents the currently selected video rendition’s identifiers.
//   - [IAVMetricPlayerItemVariantSwitchEvent.SubtitleRendition]: Represents the currently selected audio rendition’s identifiers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent
type IAVMetricPlayerItemVariantSwitchEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	DidSucceed() bool
	FromVariant() IAVAssetVariant
	ToVariant() IAVAssetVariant
	// Represents the currently selected video rendition’s identifiers.
	AudioRendition() IAVMetricMediaRendition
	// Represents the currently selected video rendition’s identifiers.
	VideoRendition() IAVMetricMediaRendition
	// Represents the currently selected audio rendition’s identifiers.
	SubtitleRendition() IAVMetricMediaRendition

	LoadedTimeRanges() []foundation.NSValue
}

// Init initializes the instance.
func (m AVMetricPlayerItemVariantSwitchEvent) Init() AVMetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[AVMetricPlayerItemVariantSwitchEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemVariantSwitchEvent) Autorelease() AVMetricPlayerItemVariantSwitchEvent {
	rv := objc.Send[AVMetricPlayerItemVariantSwitchEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemVariantSwitchEvent creates a new AVMetricPlayerItemVariantSwitchEvent instance.
func NewAVMetricPlayerItemVariantSwitchEvent() AVMetricPlayerItemVariantSwitchEvent {
	class := getAVMetricPlayerItemVariantSwitchEventClass()
	rv := objc.Send[AVMetricPlayerItemVariantSwitchEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/didSucceed
func (m AVMetricPlayerItemVariantSwitchEvent) DidSucceed() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("didSucceed"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/fromVariant
func (m AVMetricPlayerItemVariantSwitchEvent) FromVariant() IAVAssetVariant {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fromVariant"))
	return AVAssetVariantFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/toVariant
func (m AVMetricPlayerItemVariantSwitchEvent) ToVariant() IAVAssetVariant {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("toVariant"))
	return AVAssetVariantFromID(objc.ID(rv))
}
// Represents the currently selected video rendition’s identifiers.
//
// # Discussion
// 
// Subclasses of this type that are used from Swift must fulfill the
// requirements of a Sendable type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/audioRendition
func (m AVMetricPlayerItemVariantSwitchEvent) AudioRendition() IAVMetricMediaRendition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("audioRendition"))
	return AVMetricMediaRenditionFromID(objc.ID(rv))
}
// Represents the currently selected video rendition’s identifiers.
//
// # Discussion
// 
// Subclasses of this type that are used from Swift must fulfill the
// requirements of a Sendable type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/videoRendition
func (m AVMetricPlayerItemVariantSwitchEvent) VideoRendition() IAVMetricMediaRendition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("videoRendition"))
	return AVMetricMediaRenditionFromID(objc.ID(rv))
}
// Represents the currently selected audio rendition’s identifiers.
//
// # Discussion
// 
// Subclasses of this type that are used from Swift must fulfill the
// requirements of a Sendable type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/subtitleRendition
func (m AVMetricPlayerItemVariantSwitchEvent) SubtitleRendition() IAVMetricMediaRendition {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("subtitleRendition"))
	return AVMetricMediaRenditionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemVariantSwitchEvent/loadedTimeRanges-4rhjw
func (m AVMetricPlayerItemVariantSwitchEvent) LoadedTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("loadedTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

