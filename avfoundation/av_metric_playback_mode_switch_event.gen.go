// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMetricPlaybackModeSwitchEvent] class.
var (
	_AVMetricPlaybackModeSwitchEventClass     AVMetricPlaybackModeSwitchEventClass
	_AVMetricPlaybackModeSwitchEventClassOnce sync.Once
)

func getAVMetricPlaybackModeSwitchEventClass() AVMetricPlaybackModeSwitchEventClass {
	_AVMetricPlaybackModeSwitchEventClassOnce.Do(func() {
		_AVMetricPlaybackModeSwitchEventClass = AVMetricPlaybackModeSwitchEventClass{class: objc.GetClass("AVMetricPlaybackModeSwitchEvent")}
	})
	return _AVMetricPlaybackModeSwitchEventClass
}

// GetAVMetricPlaybackModeSwitchEventClass returns the class object for AVMetricPlaybackModeSwitchEvent.
func GetAVMetricPlaybackModeSwitchEventClass() AVMetricPlaybackModeSwitchEventClass {
	return getAVMetricPlaybackModeSwitchEventClass()
}

type AVMetricPlaybackModeSwitchEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMetricPlaybackModeSwitchEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMetricPlaybackModeSwitchEventClass) Alloc() AVMetricPlaybackModeSwitchEvent {
	rv := objc.Send[AVMetricPlaybackModeSwitchEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// Represents a change in playback state, entering one of AVMetricPlaybackMode
//
// # Overview
//
// Subclasses of this type that are used from Swift must fulfill the
// requirements of a Sendable type.
//
// # Identifying the playback mode
//
//   - [AVMetricPlaybackModeSwitchEvent.Mode]: Returns the mode into which playback entered.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlaybackModeSwitchEvent
type AVMetricPlaybackModeSwitchEvent struct {
	AVMetricEvent
}

// AVMetricPlaybackModeSwitchEventFromID constructs a [AVMetricPlaybackModeSwitchEvent] from an objc.ID.
//
// Represents a change in playback state, entering one of AVMetricPlaybackMode
func AVMetricPlaybackModeSwitchEventFromID(id objc.ID) AVMetricPlaybackModeSwitchEvent {
	return AVMetricPlaybackModeSwitchEvent{AVMetricEvent: AVMetricEventFromID(id)}
}

// NOTE: AVMetricPlaybackModeSwitchEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMetricPlaybackModeSwitchEvent] class.
//
// # Identifying the playback mode
//
//   - [IAVMetricPlaybackModeSwitchEvent.Mode]: Returns the mode into which playback entered.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlaybackModeSwitchEvent
type IAVMetricPlaybackModeSwitchEvent interface {
	IAVMetricEvent

	// Topic: Identifying the playback mode

	// Returns the mode into which playback entered.
	Mode() AVMetricPlaybackMode
}

// Init initializes the instance.
func (m AVMetricPlaybackModeSwitchEvent) Init() AVMetricPlaybackModeSwitchEvent {
	rv := objc.Send[AVMetricPlaybackModeSwitchEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMetricPlaybackModeSwitchEvent) Autorelease() AVMetricPlaybackModeSwitchEvent {
	rv := objc.Send[AVMetricPlaybackModeSwitchEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMetricPlaybackModeSwitchEvent creates a new AVMetricPlaybackModeSwitchEvent instance.
func NewAVMetricPlaybackModeSwitchEvent() AVMetricPlaybackModeSwitchEvent {
	class := getAVMetricPlaybackModeSwitchEventClass()
	rv := objc.Send[AVMetricPlaybackModeSwitchEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMetricEvent/init(coder:)
func NewMetricPlaybackModeSwitchEventWithCoder(coder foundation.INSCoder) AVMetricPlaybackModeSwitchEvent {
	instance := getAVMetricPlaybackModeSwitchEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return AVMetricPlaybackModeSwitchEventFromID(rv)
}

// Returns the mode into which playback entered.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetricPlaybackModeSwitchEvent/mode
func (m AVMetricPlaybackModeSwitchEvent) Mode() AVMetricPlaybackMode {
	rv := objc.Send[AVMetricPlaybackMode](m.ID, objc.Sel("mode"))
	return AVMetricPlaybackMode(rv)
}
