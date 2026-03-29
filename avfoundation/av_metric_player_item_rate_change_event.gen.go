// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlayerItemRateChangeEvent] class.
var (
	_AVMetricPlayerItemRateChangeEventClass     AVMetricPlayerItemRateChangeEventClass
	_AVMetricPlayerItemRateChangeEventClassOnce sync.Once
)

func getAVMetricPlayerItemRateChangeEventClass() AVMetricPlayerItemRateChangeEventClass {
	_AVMetricPlayerItemRateChangeEventClassOnce.Do(func() {
		_AVMetricPlayerItemRateChangeEventClass = AVMetricPlayerItemRateChangeEventClass{class: objc.GetClass("AVMetricPlayerItemRateChangeEvent")}
	})
	return _AVMetricPlayerItemRateChangeEventClass
}

// GetAVMetricPlayerItemRateChangeEventClass returns the class object for AVMetricPlayerItemRateChangeEvent.
func GetAVMetricPlayerItemRateChangeEventClass() AVMetricPlayerItemRateChangeEventClass {
	return getAVMetricPlayerItemRateChangeEventClass()
}

type AVMetricPlayerItemRateChangeEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemRateChangeEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemRateChangeEventClass) Alloc() AVMetricPlayerItemRateChangeEvent {
	rv := objc.Send[AVMetricPlayerItemRateChangeEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when the playback rate changes.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemRateChangeEvent.PreviousRate]
//   - [AVMetricPlayerItemRateChangeEvent.Rate]
//   - [AVMetricPlayerItemRateChangeEvent.Variant]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent
type AVMetricPlayerItemRateChangeEvent struct {
	AVMetricEvent
}

// AVMetricPlayerItemRateChangeEventFromID constructs a [AVMetricPlayerItemRateChangeEvent] from an objc.ID.
//
// An event that represents when the playback rate changes.
func AVMetricPlayerItemRateChangeEventFromID(id objc.ID) AVMetricPlayerItemRateChangeEvent {
	return AVMetricPlayerItemRateChangeEvent{AVMetricEvent: AVMetricEventFromID(id)}
}
// NOTE: AVMetricPlayerItemRateChangeEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemRateChangeEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemRateChangeEvent.PreviousRate]
//   - [IAVMetricPlayerItemRateChangeEvent.Rate]
//   - [IAVMetricPlayerItemRateChangeEvent.Variant]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent
type IAVMetricPlayerItemRateChangeEvent interface {
	IAVMetricEvent

	// Topic: Inspecting the event

	PreviousRate() float64
	Rate() float64
	Variant() IAVAssetVariant
}

// Init initializes the instance.
func (m AVMetricPlayerItemRateChangeEvent) Init() AVMetricPlayerItemRateChangeEvent {
	rv := objc.Send[AVMetricPlayerItemRateChangeEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemRateChangeEvent) Autorelease() AVMetricPlayerItemRateChangeEvent {
	rv := objc.Send[AVMetricPlayerItemRateChangeEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemRateChangeEvent creates a new AVMetricPlayerItemRateChangeEvent instance.
func NewAVMetricPlayerItemRateChangeEvent() AVMetricPlayerItemRateChangeEvent {
	class := getAVMetricPlayerItemRateChangeEventClass()
	rv := objc.Send[AVMetricPlayerItemRateChangeEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent/previousRate
func (m AVMetricPlayerItemRateChangeEvent) PreviousRate() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("previousRate"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent/rate
func (m AVMetricPlayerItemRateChangeEvent) Rate() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("rate"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemRateChangeEvent/variant
func (m AVMetricPlayerItemRateChangeEvent) Variant() IAVAssetVariant {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("variant"))
	return AVAssetVariantFromID(objc.ID(rv))
}

