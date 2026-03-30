// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlayerItemSeekDidCompleteEvent] class.
var (
	_AVMetricPlayerItemSeekDidCompleteEventClass     AVMetricPlayerItemSeekDidCompleteEventClass
	_AVMetricPlayerItemSeekDidCompleteEventClassOnce sync.Once
)

func getAVMetricPlayerItemSeekDidCompleteEventClass() AVMetricPlayerItemSeekDidCompleteEventClass {
	_AVMetricPlayerItemSeekDidCompleteEventClassOnce.Do(func() {
		_AVMetricPlayerItemSeekDidCompleteEventClass = AVMetricPlayerItemSeekDidCompleteEventClass{class: objc.GetClass("AVMetricPlayerItemSeekDidCompleteEvent")}
	})
	return _AVMetricPlayerItemSeekDidCompleteEventClass
}

// GetAVMetricPlayerItemSeekDidCompleteEventClass returns the class object for AVMetricPlayerItemSeekDidCompleteEvent.
func GetAVMetricPlayerItemSeekDidCompleteEventClass() AVMetricPlayerItemSeekDidCompleteEventClass {
	return getAVMetricPlayerItemSeekDidCompleteEventClass()
}

type AVMetricPlayerItemSeekDidCompleteEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlayerItemSeekDidCompleteEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlayerItemSeekDidCompleteEventClass) Alloc() AVMetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[AVMetricPlayerItemSeekDidCompleteEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An event that represents when the playback seek completes.
//
// # Inspecting the event
//
//   - [AVMetricPlayerItemSeekDidCompleteEvent.DidSeekInBuffer]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent
type AVMetricPlayerItemSeekDidCompleteEvent struct {
	AVMetricPlayerItemRateChangeEvent
}

// AVMetricPlayerItemSeekDidCompleteEventFromID constructs a [AVMetricPlayerItemSeekDidCompleteEvent] from an objc.ID.
//
// An event that represents when the playback seek completes.
func AVMetricPlayerItemSeekDidCompleteEventFromID(id objc.ID) AVMetricPlayerItemSeekDidCompleteEvent {
	return AVMetricPlayerItemSeekDidCompleteEvent{AVMetricPlayerItemRateChangeEvent: AVMetricPlayerItemRateChangeEventFromID(id)}
}

// NOTE: AVMetricPlayerItemSeekDidCompleteEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlayerItemSeekDidCompleteEvent] class.
//
// # Inspecting the event
//
//   - [IAVMetricPlayerItemSeekDidCompleteEvent.DidSeekInBuffer]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent
type IAVMetricPlayerItemSeekDidCompleteEvent interface {
	IAVMetricPlayerItemRateChangeEvent

	// Topic: Inspecting the event

	DidSeekInBuffer() bool
}

// Init initializes the instance.
func (m AVMetricPlayerItemSeekDidCompleteEvent) Init() AVMetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[AVMetricPlayerItemSeekDidCompleteEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlayerItemSeekDidCompleteEvent) Autorelease() AVMetricPlayerItemSeekDidCompleteEvent {
	rv := objc.Send[AVMetricPlayerItemSeekDidCompleteEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlayerItemSeekDidCompleteEvent creates a new AVMetricPlayerItemSeekDidCompleteEvent instance.
func NewAVMetricPlayerItemSeekDidCompleteEvent() AVMetricPlayerItemSeekDidCompleteEvent {
	class := getAVMetricPlayerItemSeekDidCompleteEventClass()
	rv := objc.Send[AVMetricPlayerItemSeekDidCompleteEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlayerItemSeekDidCompleteEvent/didSeekInBuffer
func (m AVMetricPlayerItemSeekDidCompleteEvent) DidSeekInBuffer() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("didSeekInBuffer"))
	return rv
}
