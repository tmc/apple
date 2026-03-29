// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIPitchBendEvent] class.
var (
	_AVMIDIPitchBendEventClass     AVMIDIPitchBendEventClass
	_AVMIDIPitchBendEventClassOnce sync.Once
)

func getAVMIDIPitchBendEventClass() AVMIDIPitchBendEventClass {
	_AVMIDIPitchBendEventClassOnce.Do(func() {
		_AVMIDIPitchBendEventClass = AVMIDIPitchBendEventClass{class: objc.GetClass("AVMIDIPitchBendEvent")}
	})
	return _AVMIDIPitchBendEventClass
}

// GetAVMIDIPitchBendEventClass returns the class object for AVMIDIPitchBendEvent.
func GetAVMIDIPitchBendEventClass() AVMIDIPitchBendEventClass {
	return getAVMIDIPitchBendEventClass()
}

type AVMIDIPitchBendEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIPitchBendEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIPitchBendEventClass) Alloc() AVMIDIPitchBendEvent {
	rv := objc.Send[AVMIDIPitchBendEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI pitch bend message.
//
// # Creating a Pitch Bend Event
//
//   - [AVMIDIPitchBendEvent.InitWithChannelValue]: Creates an event with a channel and pitch bend value.
//
// # Configuring a Pitch Bend Event
//
//   - [AVMIDIPitchBendEvent.Value]: The value of the pitch bend event.
//   - [AVMIDIPitchBendEvent.SetValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent
type AVMIDIPitchBendEvent struct {
	AVMIDIChannelEvent
}

// AVMIDIPitchBendEventFromID constructs a [AVMIDIPitchBendEvent] from an objc.ID.
//
// An object that represents a MIDI pitch bend message.
func AVMIDIPitchBendEventFromID(id objc.ID) AVMIDIPitchBendEvent {
	return AVMIDIPitchBendEvent{AVMIDIChannelEvent: AVMIDIChannelEventFromID(id)}
}
// NOTE: AVMIDIPitchBendEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIPitchBendEvent] class.
//
// # Creating a Pitch Bend Event
//
//   - [IAVMIDIPitchBendEvent.InitWithChannelValue]: Creates an event with a channel and pitch bend value.
//
// # Configuring a Pitch Bend Event
//
//   - [IAVMIDIPitchBendEvent.Value]: The value of the pitch bend event.
//   - [IAVMIDIPitchBendEvent.SetValue]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent
type IAVMIDIPitchBendEvent interface {
	IAVMIDIChannelEvent

	// Topic: Creating a Pitch Bend Event

	// Creates an event with a channel and pitch bend value.
	InitWithChannelValue(channel uint32, value uint32) AVMIDIPitchBendEvent

	// Topic: Configuring a Pitch Bend Event

	// The value of the pitch bend event.
	Value() uint32
	SetValue(value uint32)
}

// Init initializes the instance.
func (m AVMIDIPitchBendEvent) Init() AVMIDIPitchBendEvent {
	rv := objc.Send[AVMIDIPitchBendEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIPitchBendEvent) Autorelease() AVMIDIPitchBendEvent {
	rv := objc.Send[AVMIDIPitchBendEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIPitchBendEvent creates a new AVMIDIPitchBendEvent instance.
func NewAVMIDIPitchBendEvent() AVMIDIPitchBendEvent {
	class := getAVMIDIPitchBendEventClass()
	rv := objc.Send[AVMIDIPitchBendEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a channel and pitch bend value.
//
// channel: The MIDI channel for the message, between `0` and `15`.
//
// value: The pitch bend value, between `0` and `16383`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent/init(channel:value:)
func NewMIDIPitchBendEventWithChannelValue(channel uint32, value uint32) AVMIDIPitchBendEvent {
	instance := getAVMIDIPitchBendEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:value:"), channel, value)
	return AVMIDIPitchBendEventFromID(rv)
}

// Creates an event with a channel and pitch bend value.
//
// channel: The MIDI channel for the message, between `0` and `15`.
//
// value: The pitch bend value, between `0` and `16383`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent/init(channel:value:)
func (m AVMIDIPitchBendEvent) InitWithChannelValue(channel uint32, value uint32) AVMIDIPitchBendEvent {
	rv := objc.Send[AVMIDIPitchBendEvent](m.ID, objc.Sel("initWithChannel:value:"), channel, value)
	return rv
}

// The value of the pitch bend event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPitchBendEvent/value
func (m AVMIDIPitchBendEvent) Value() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("value"))
	return rv
}
func (m AVMIDIPitchBendEvent) SetValue(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setValue:"), value)
}

