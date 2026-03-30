// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIMetaEvent] class.
var (
	_AVMIDIMetaEventClass     AVMIDIMetaEventClass
	_AVMIDIMetaEventClassOnce sync.Once
)

func getAVMIDIMetaEventClass() AVMIDIMetaEventClass {
	_AVMIDIMetaEventClassOnce.Do(func() {
		_AVMIDIMetaEventClass = AVMIDIMetaEventClass{class: objc.GetClass("AVMIDIMetaEvent")}
	})
	return _AVMIDIMetaEventClass
}

// GetAVMIDIMetaEventClass returns the class object for AVMIDIMetaEvent.
func GetAVMIDIMetaEventClass() AVMIDIMetaEventClass {
	return getAVMIDIMetaEventClass()
}

type AVMIDIMetaEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIMetaEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIMetaEventClass) Alloc() AVMIDIMetaEvent {
	rv := objc.Send[AVMIDIMetaEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents MIDI meta event messages.
//
// # Overview
//
// You can’t modify the size and contents of this event once you create it.
// This doesn’t verify that the content matches the MIDI specification.
//
// You can only add [AVMIDIMetaEventTypeTempo],
// [AVMIDIMetaEventTypeSmpteOffset], or [AVMIDIMetaEventTypeTimeSignature] to
// a sequence’s tempo track.
//
// # Creating a Meta Event
//
//   - [AVMIDIMetaEvent.InitWithTypeData]: Creates an event with a MIDI meta event type and data.
//
// # Getting the Meta Event Type
//
//   - [AVMIDIMetaEvent.Type]: The type of meta event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent
type AVMIDIMetaEvent struct {
	AVMusicEvent
}

// AVMIDIMetaEventFromID constructs a [AVMIDIMetaEvent] from an objc.ID.
//
// An object that represents MIDI meta event messages.
func AVMIDIMetaEventFromID(id objc.ID) AVMIDIMetaEvent {
	return AVMIDIMetaEvent{AVMusicEvent: AVMusicEventFromID(id)}
}

// NOTE: AVMIDIMetaEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIMetaEvent] class.
//
// # Creating a Meta Event
//
//   - [IAVMIDIMetaEvent.InitWithTypeData]: Creates an event with a MIDI meta event type and data.
//
// # Getting the Meta Event Type
//
//   - [IAVMIDIMetaEvent.Type]: The type of meta event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent
type IAVMIDIMetaEvent interface {
	IAVMusicEvent

	// Topic: Creating a Meta Event

	// Creates an event with a MIDI meta event type and data.
	InitWithTypeData(type_ AVMIDIMetaEventType, data foundation.INSData) AVMIDIMetaEvent

	// Topic: Getting the Meta Event Type

	// The type of meta event.
	Type() AVMIDIMetaEventType
}

// Init initializes the instance.
func (m AVMIDIMetaEvent) Init() AVMIDIMetaEvent {
	rv := objc.Send[AVMIDIMetaEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIMetaEvent) Autorelease() AVMIDIMetaEvent {
	rv := objc.Send[AVMIDIMetaEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIMetaEvent creates a new AVMIDIMetaEvent instance.
func NewAVMIDIMetaEvent() AVMIDIMetaEvent {
	class := getAVMIDIMetaEventClass()
	rv := objc.Send[AVMIDIMetaEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a MIDI meta event type and data.
//
// type: The meta event type.
//
// data: The data that contains the contents of the meta event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/init(type:data:)
func NewMIDIMetaEventWithTypeData(type_ AVMIDIMetaEventType, data foundation.INSData) AVMIDIMetaEvent {
	instance := getAVMIDIMetaEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:data:"), type_, data)
	return AVMIDIMetaEventFromID(rv)
}

// Creates an event with a MIDI meta event type and data.
//
// type: The meta event type.
//
// data: The data that contains the contents of the meta event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/init(type:data:)
func (m AVMIDIMetaEvent) InitWithTypeData(type_ AVMIDIMetaEventType, data foundation.INSData) AVMIDIMetaEvent {
	rv := objc.Send[AVMIDIMetaEvent](m.ID, objc.Sel("initWithType:data:"), type_, data)
	return rv
}

// The type of meta event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIMetaEvent/type
func (m AVMIDIMetaEvent) Type() AVMIDIMetaEventType {
	rv := objc.Send[AVMIDIMetaEventType](m.ID, objc.Sel("type"))
	return AVMIDIMetaEventType(rv)
}
