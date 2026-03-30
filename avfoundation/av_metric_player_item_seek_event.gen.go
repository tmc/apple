// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlayerItemSeekEvent] class.
var (
	_AVMetricPlayerItemSeekEventClass     AVMetricPlayerItemSeekEventClass
	_AVMetricPlayerItemSeekEventClassOnce sync.Once
)

func getAVMetricPlayerItemSeekEventClass() AVMetricPlayerItemSeekEventClass {
	_AVMetricPlayerItemSeekEventClassOnce.Do(func() {
		_AVMetricPlayerItemSeekEventClass = AVMetricPlayerItemSeekEventClass{class: objc.GetClass("AVMetricPlayerItemSeekEvent")}
	})
	return _AVMetricPlayerItemSeekEventClass
}

// GetAVMetricPlayerItemSeekEventClass returns the class object for AVMetricPlayerItemSeekEvent.
func GetAVMetricPlayerItemSeekEventClass() AVMetricPlayerItemSeekEventClass {
	return getAVMetricPlayerItemSeekEventClass()
}

type AVMetricPlayerItemSeekEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemSeekEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemSeekEventClass) Alloc() AVMetricPlayerItemSeekEvent {
	rv := objc.Send[AVMetricPlayerItemSeekEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when a playback seek occurs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekEvent
type AVMetricPlayerItemSeekEvent struct {
	AVMetricPlayerItemRateChangeEvent
}

// AVMetricPlayerItemSeekEventFromID constructs a [AVMetricPlayerItemSeekEvent] from an objc.ID.
//
// An event that represents when a playback seek occurs.
func AVMetricPlayerItemSeekEventFromID(id objc.ID) AVMetricPlayerItemSeekEvent {
	return AVMetricPlayerItemSeekEvent{AVMetricPlayerItemRateChangeEvent: AVMetricPlayerItemRateChangeEventFromID(id)}
}

// NOTE: AVMetricPlayerItemSeekEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemSeekEvent] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekEvent
type IAVMetricPlayerItemSeekEvent interface {
	IAVMetricPlayerItemRateChangeEvent
}

// Init initializes the instance.
func (m AVMetricPlayerItemSeekEvent) Init() AVMetricPlayerItemSeekEvent {
	rv := objc.Send[AVMetricPlayerItemSeekEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemSeekEvent) Autorelease() AVMetricPlayerItemSeekEvent {
	rv := objc.Send[AVMetricPlayerItemSeekEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemSeekEvent creates a new AVMetricPlayerItemSeekEvent instance.
func NewAVMetricPlayerItemSeekEvent() AVMetricPlayerItemSeekEvent {
	class := getAVMetricPlayerItemSeekEventClass()
	rv := objc.Send[AVMetricPlayerItemSeekEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
