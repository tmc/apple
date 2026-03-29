// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIChannelEvent] class.
var (
	_AVMIDIChannelEventClass     AVMIDIChannelEventClass
	_AVMIDIChannelEventClassOnce sync.Once
)

func getAVMIDIChannelEventClass() AVMIDIChannelEventClass {
	_AVMIDIChannelEventClassOnce.Do(func() {
		_AVMIDIChannelEventClass = AVMIDIChannelEventClass{class: objc.GetClass("AVMIDIChannelEvent")}
	})
	return _AVMIDIChannelEventClass
}

// GetAVMIDIChannelEventClass returns the class object for AVMIDIChannelEvent.
func GetAVMIDIChannelEventClass() AVMIDIChannelEventClass {
	return getAVMIDIChannelEventClass()
}

type AVMIDIChannelEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIChannelEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIChannelEventClass) Alloc() AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A base class for all MIDI messages that operate on a single MIDI channel.
//
// # Configuring a Channel Event
//
//   - [AVMIDIChannelEvent.Channel]: The MIDI channel.
//   - [AVMIDIChannelEvent.SetChannel]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent
type AVMIDIChannelEvent struct {
	AVMusicEvent
}

// AVMIDIChannelEventFromID constructs a [AVMIDIChannelEvent] from an objc.ID.
//
// A base class for all MIDI messages that operate on a single MIDI channel.
func AVMIDIChannelEventFromID(id objc.ID) AVMIDIChannelEvent {
	return AVMIDIChannelEvent{AVMusicEvent: AVMusicEventFromID(id)}
}
// NOTE: AVMIDIChannelEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIChannelEvent] class.
//
// # Configuring a Channel Event
//
//   - [IAVMIDIChannelEvent.Channel]: The MIDI channel.
//   - [IAVMIDIChannelEvent.SetChannel]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent
type IAVMIDIChannelEvent interface {
	IAVMusicEvent

	// Topic: Configuring a Channel Event

	// The MIDI channel.
	Channel() uint32
	SetChannel(value uint32)
}

// Init initializes the instance.
func (m AVMIDIChannelEvent) Init() AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIChannelEvent) Autorelease() AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIChannelEvent creates a new AVMIDIChannelEvent instance.
func NewAVMIDIChannelEvent() AVMIDIChannelEvent {
	class := getAVMIDIChannelEventClass()
	rv := objc.Send[AVMIDIChannelEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The MIDI channel.
//
// # Discussion
// 
// The valid range of values are between `0` and `15`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/channel
func (m AVMIDIChannelEvent) Channel() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("channel"))
	return rv
}
func (m AVMIDIChannelEvent) SetChannel(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setChannel:"), value)
}

