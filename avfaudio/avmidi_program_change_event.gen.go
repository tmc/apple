// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIProgramChangeEvent] class.
var (
	_AVMIDIProgramChangeEventClass     AVMIDIProgramChangeEventClass
	_AVMIDIProgramChangeEventClassOnce sync.Once
)

func getAVMIDIProgramChangeEventClass() AVMIDIProgramChangeEventClass {
	_AVMIDIProgramChangeEventClassOnce.Do(func() {
		_AVMIDIProgramChangeEventClass = AVMIDIProgramChangeEventClass{class: objc.GetClass("AVMIDIProgramChangeEvent")}
	})
	return _AVMIDIProgramChangeEventClass
}

// GetAVMIDIProgramChangeEventClass returns the class object for AVMIDIProgramChangeEvent.
func GetAVMIDIProgramChangeEventClass() AVMIDIProgramChangeEventClass {
	return getAVMIDIProgramChangeEventClass()
}

type AVMIDIProgramChangeEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIProgramChangeEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIProgramChangeEventClass) Alloc() AVMIDIProgramChangeEvent {
	rv := objc.Send[AVMIDIProgramChangeEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI program or patch change message.
//
// # Overview
// 
// The effect of this message depends on the [AVMusicTrack] destination audio
// unit.
//
// # Creating a Program Change Event
//
//   - [AVMIDIProgramChangeEvent.InitWithChannelProgramNumber]: Creates a program change event with a channel and program number.
//
// # Configuring a Program Change Event
//
//   - [AVMIDIProgramChangeEvent.ProgramNumber]: The MIDI program number.
//   - [AVMIDIProgramChangeEvent.SetProgramNumber]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent
type AVMIDIProgramChangeEvent struct {
	AVMIDIChannelEvent
}

// AVMIDIProgramChangeEventFromID constructs a [AVMIDIProgramChangeEvent] from an objc.ID.
//
// An object that represents a MIDI program or patch change message.
func AVMIDIProgramChangeEventFromID(id objc.ID) AVMIDIProgramChangeEvent {
	return AVMIDIProgramChangeEvent{AVMIDIChannelEvent: AVMIDIChannelEventFromID(id)}
}
// NOTE: AVMIDIProgramChangeEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIProgramChangeEvent] class.
//
// # Creating a Program Change Event
//
//   - [IAVMIDIProgramChangeEvent.InitWithChannelProgramNumber]: Creates a program change event with a channel and program number.
//
// # Configuring a Program Change Event
//
//   - [IAVMIDIProgramChangeEvent.ProgramNumber]: The MIDI program number.
//   - [IAVMIDIProgramChangeEvent.SetProgramNumber]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent
type IAVMIDIProgramChangeEvent interface {
	IAVMIDIChannelEvent

	// Topic: Creating a Program Change Event

	// Creates a program change event with a channel and program number.
	InitWithChannelProgramNumber(channel uint32, programNumber uint32) AVMIDIProgramChangeEvent

	// Topic: Configuring a Program Change Event

	// The MIDI program number.
	ProgramNumber() uint32
	SetProgramNumber(value uint32)
}

// Init initializes the instance.
func (m AVMIDIProgramChangeEvent) Init() AVMIDIProgramChangeEvent {
	rv := objc.Send[AVMIDIProgramChangeEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIProgramChangeEvent) Autorelease() AVMIDIProgramChangeEvent {
	rv := objc.Send[AVMIDIProgramChangeEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIProgramChangeEvent creates a new AVMIDIProgramChangeEvent instance.
func NewAVMIDIProgramChangeEvent() AVMIDIProgramChangeEvent {
	class := getAVMIDIProgramChangeEventClass()
	rv := objc.Send[AVMIDIProgramChangeEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a program change event with a channel and program number.
//
// channel: The MIDI channel for the message, between `0` and `15`.
//
// programNumber: The program number to send, between `0` and `127`.
//
// # Discussion
// 
// The instrument this chooses depends on
// [MIDIControlChangeMessageTypeBankSelect] events sent prior to this event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/init(channel:programNumber:)
func NewMIDIProgramChangeEventWithChannelProgramNumber(channel uint32, programNumber uint32) AVMIDIProgramChangeEvent {
	instance := getAVMIDIProgramChangeEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:programNumber:"), channel, programNumber)
	return AVMIDIProgramChangeEventFromID(rv)
}

// Creates a program change event with a channel and program number.
//
// channel: The MIDI channel for the message, between `0` and `15`.
//
// programNumber: The program number to send, between `0` and `127`.
//
// # Discussion
// 
// The instrument this chooses depends on
// [MIDIControlChangeMessageTypeBankSelect] events sent prior to this event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/init(channel:programNumber:)
func (m AVMIDIProgramChangeEvent) InitWithChannelProgramNumber(channel uint32, programNumber uint32) AVMIDIProgramChangeEvent {
	rv := objc.Send[AVMIDIProgramChangeEvent](m.ID, objc.Sel("initWithChannel:programNumber:"), channel, programNumber)
	return rv
}

// The MIDI program number.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIProgramChangeEvent/programNumber
func (m AVMIDIProgramChangeEvent) ProgramNumber() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("programNumber"))
	return rv
}
func (m AVMIDIProgramChangeEvent) SetProgramNumber(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgramNumber:"), value)
}

