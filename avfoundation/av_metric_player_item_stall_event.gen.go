// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlayerItemStallEvent] class.
var (
	_AVMetricPlayerItemStallEventClass     AVMetricPlayerItemStallEventClass
	_AVMetricPlayerItemStallEventClassOnce sync.Once
)

func getAVMetricPlayerItemStallEventClass() AVMetricPlayerItemStallEventClass {
	_AVMetricPlayerItemStallEventClassOnce.Do(func() {
		_AVMetricPlayerItemStallEventClass = AVMetricPlayerItemStallEventClass{class: objc.GetClass("AVMetricPlayerItemStallEvent")}
	})
	return _AVMetricPlayerItemStallEventClass
}

// GetAVMetricPlayerItemStallEventClass returns the class object for AVMetricPlayerItemStallEvent.
func GetAVMetricPlayerItemStallEventClass() AVMetricPlayerItemStallEventClass {
	return getAVMetricPlayerItemStallEventClass()
}

type AVMetricPlayerItemStallEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemStallEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemStallEventClass) Alloc() AVMetricPlayerItemStallEvent {
	rv := objc.Send[AVMetricPlayerItemStallEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when playback stalls.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemStallEvent
type AVMetricPlayerItemStallEvent struct {
	AVMetricPlayerItemRateChangeEvent
}

// AVMetricPlayerItemStallEventFromID constructs a [AVMetricPlayerItemStallEvent] from an objc.ID.
//
// An event that represents when playback stalls.
func AVMetricPlayerItemStallEventFromID(id objc.ID) AVMetricPlayerItemStallEvent {
	return AVMetricPlayerItemStallEvent{AVMetricPlayerItemRateChangeEvent: AVMetricPlayerItemRateChangeEventFromID(id)}
}
// NOTE: AVMetricPlayerItemStallEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemStallEvent] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemStallEvent
type IAVMetricPlayerItemStallEvent interface {
	IAVMetricPlayerItemRateChangeEvent
}

// Init initializes the instance.
func (m AVMetricPlayerItemStallEvent) Init() AVMetricPlayerItemStallEvent {
	rv := objc.Send[AVMetricPlayerItemStallEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemStallEvent) Autorelease() AVMetricPlayerItemStallEvent {
	rv := objc.Send[AVMetricPlayerItemStallEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemStallEvent creates a new AVMetricPlayerItemStallEvent instance.
func NewAVMetricPlayerItemStallEvent() AVMetricPlayerItemStallEvent {
	class := getAVMetricPlayerItemStallEventClass()
	rv := objc.Send[AVMetricPlayerItemStallEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

