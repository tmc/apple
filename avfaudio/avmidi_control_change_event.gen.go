// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIControlChangeEvent] class.
var (
	_AVMIDIControlChangeEventClass     AVMIDIControlChangeEventClass
	_AVMIDIControlChangeEventClassOnce sync.Once
)

func getAVMIDIControlChangeEventClass() AVMIDIControlChangeEventClass {
	_AVMIDIControlChangeEventClassOnce.Do(func() {
		_AVMIDIControlChangeEventClass = AVMIDIControlChangeEventClass{class: objc.GetClass("AVMIDIControlChangeEvent")}
	})
	return _AVMIDIControlChangeEventClass
}

// GetAVMIDIControlChangeEventClass returns the class object for AVMIDIControlChangeEvent.
func GetAVMIDIControlChangeEventClass() AVMIDIControlChangeEventClass {
	return getAVMIDIControlChangeEventClass()
}

type AVMIDIControlChangeEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIControlChangeEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIControlChangeEventClass) Alloc() AVMIDIControlChangeEvent {
	rv := objc.Send[AVMIDIControlChangeEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI control change message.
//
// # Creating a Control Change Event
//
//   - [AVMIDIControlChangeEvent.InitWithChannelMessageTypeValue]: Creates an event with a channel, control change type, and a value.
//
// # Inspecting a Control Change Event
//
//   - [AVMIDIControlChangeEvent.Value]: The value of the control change event.
//   - [AVMIDIControlChangeEvent.MessageType]: The type of control change message.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent
type AVMIDIControlChangeEvent struct {
	AVMIDIChannelEvent
}

// AVMIDIControlChangeEventFromID constructs a [AVMIDIControlChangeEvent] from an objc.ID.
//
// An object that represents a MIDI control change message.
func AVMIDIControlChangeEventFromID(id objc.ID) AVMIDIControlChangeEvent {
	return AVMIDIControlChangeEvent{AVMIDIChannelEvent: AVMIDIChannelEventFromID(id)}
}

// NOTE: AVMIDIControlChangeEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIControlChangeEvent] class.
//
// # Creating a Control Change Event
//
//   - [IAVMIDIControlChangeEvent.InitWithChannelMessageTypeValue]: Creates an event with a channel, control change type, and a value.
//
// # Inspecting a Control Change Event
//
//   - [IAVMIDIControlChangeEvent.Value]: The value of the control change event.
//   - [IAVMIDIControlChangeEvent.MessageType]: The type of control change message.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent
type IAVMIDIControlChangeEvent interface {
	IAVMIDIChannelEvent

	// Topic: Creating a Control Change Event

	// Creates an event with a channel, control change type, and a value.
	InitWithChannelMessageTypeValue(channel uint32, messageType AVMIDIControlChangeMessageType, value uint32) AVMIDIControlChangeEvent

	// Topic: Inspecting a Control Change Event

	// The value of the control change event.
	Value() uint32
	// The type of control change message.
	MessageType() AVMIDIControlChangeMessageType
}

// Init initializes the instance.
func (m AVMIDIControlChangeEvent) Init() AVMIDIControlChangeEvent {
	rv := objc.Send[AVMIDIControlChangeEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIControlChangeEvent) Autorelease() AVMIDIControlChangeEvent {
	rv := objc.Send[AVMIDIControlChangeEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIControlChangeEvent creates a new AVMIDIControlChangeEvent instance.
func NewAVMIDIControlChangeEvent() AVMIDIControlChangeEvent {
	class := getAVMIDIControlChangeEventClass()
	rv := objc.Send[AVMIDIControlChangeEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a channel, control change type, and a value.
//
// channel: The MIDI channel for the control change, between `0` and `15`.
//
// messageType: The type that indicates which MIDI control change message to send.
//
// value: The value for the control change.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/init(channel:messageType:value:)
func NewMIDIControlChangeEventWithChannelMessageTypeValue(channel uint32, messageType AVMIDIControlChangeMessageType, value uint32) AVMIDIControlChangeEvent {
	instance := getAVMIDIControlChangeEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:messageType:value:"), channel, messageType, value)
	return AVMIDIControlChangeEventFromID(rv)
}

// Creates an event with a channel, control change type, and a value.
//
// channel: The MIDI channel for the control change, between `0` and `15`.
//
// messageType: The type that indicates which MIDI control change message to send.
//
// value: The value for the control change.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/init(channel:messageType:value:)
func (m AVMIDIControlChangeEvent) InitWithChannelMessageTypeValue(channel uint32, messageType AVMIDIControlChangeMessageType, value uint32) AVMIDIControlChangeEvent {
	rv := objc.Send[AVMIDIControlChangeEvent](m.ID, objc.Sel("initWithChannel:messageType:value:"), channel, messageType, value)
	return rv
}

// The value of the control change event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/value
func (m AVMIDIControlChangeEvent) Value() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("value"))
	return rv
}

// The type of control change message.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIControlChangeEvent/messageType-swift.property
func (m AVMIDIControlChangeEvent) MessageType() AVMIDIControlChangeMessageType {
	rv := objc.Send[AVMIDIControlChangeMessageType](m.ID, objc.Sel("messageType"))
	return AVMIDIControlChangeMessageType(rv)
}
