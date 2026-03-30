// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIChannelPressureEvent] class.
var (
	_AVMIDIChannelPressureEventClass     AVMIDIChannelPressureEventClass
	_AVMIDIChannelPressureEventClassOnce sync.Once
)

func getAVMIDIChannelPressureEventClass() AVMIDIChannelPressureEventClass {
	_AVMIDIChannelPressureEventClassOnce.Do(func() {
		_AVMIDIChannelPressureEventClass = AVMIDIChannelPressureEventClass{class: objc.GetClass("AVMIDIChannelPressureEvent")}
	})
	return _AVMIDIChannelPressureEventClass
}

// GetAVMIDIChannelPressureEventClass returns the class object for AVMIDIChannelPressureEvent.
func GetAVMIDIChannelPressureEventClass() AVMIDIChannelPressureEventClass {
	return getAVMIDIChannelPressureEventClass()
}

type AVMIDIChannelPressureEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIChannelPressureEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIChannelPressureEventClass) Alloc() AVMIDIChannelPressureEvent {
	rv := objc.Send[AVMIDIChannelPressureEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI channel pressure message.
//
// # Overview
//
// The effect of this message depends on the [AVMusicTrack] destination audio
// unit, and the capabilities of the destination’s loaded instrument.
//
// # Creating a Pressure Event
//
//   - [AVMIDIChannelPressureEvent.InitWithChannelPressure]: Creates a pressure event with a channel and pressure value.
//
// # Configuring a Pressure Event
//
//   - [AVMIDIChannelPressureEvent.Pressure]: The MIDI channel pressure.
//   - [AVMIDIChannelPressureEvent.SetPressure]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent
type AVMIDIChannelPressureEvent struct {
	AVMIDIChannelEvent
}

// AVMIDIChannelPressureEventFromID constructs a [AVMIDIChannelPressureEvent] from an objc.ID.
//
// An object that represents a MIDI channel pressure message.
func AVMIDIChannelPressureEventFromID(id objc.ID) AVMIDIChannelPressureEvent {
	return AVMIDIChannelPressureEvent{AVMIDIChannelEvent: AVMIDIChannelEventFromID(id)}
}

// NOTE: AVMIDIChannelPressureEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIChannelPressureEvent] class.
//
// # Creating a Pressure Event
//
//   - [IAVMIDIChannelPressureEvent.InitWithChannelPressure]: Creates a pressure event with a channel and pressure value.
//
// # Configuring a Pressure Event
//
//   - [IAVMIDIChannelPressureEvent.Pressure]: The MIDI channel pressure.
//   - [IAVMIDIChannelPressureEvent.SetPressure]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent
type IAVMIDIChannelPressureEvent interface {
	IAVMIDIChannelEvent

	// Topic: Creating a Pressure Event

	// Creates a pressure event with a channel and pressure value.
	InitWithChannelPressure(channel uint32, pressure uint32) AVMIDIChannelPressureEvent

	// Topic: Configuring a Pressure Event

	// The MIDI channel pressure.
	Pressure() uint32
	SetPressure(value uint32)
}

// Init initializes the instance.
func (m AVMIDIChannelPressureEvent) Init() AVMIDIChannelPressureEvent {
	rv := objc.Send[AVMIDIChannelPressureEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIChannelPressureEvent) Autorelease() AVMIDIChannelPressureEvent {
	rv := objc.Send[AVMIDIChannelPressureEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIChannelPressureEvent creates a new AVMIDIChannelPressureEvent instance.
func NewAVMIDIChannelPressureEvent() AVMIDIChannelPressureEvent {
	class := getAVMIDIChannelPressureEventClass()
	rv := objc.Send[AVMIDIChannelPressureEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a pressure event with a channel and pressure value.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent/init(channel:pressure:)
func NewMIDIChannelPressureEventWithChannelPressure(channel uint32, pressure uint32) AVMIDIChannelPressureEvent {
	instance := getAVMIDIChannelPressureEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:pressure:"), channel, pressure)
	return AVMIDIChannelPressureEventFromID(rv)
}

// Creates a pressure event with a channel and pressure value.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent/init(channel:pressure:)
func (m AVMIDIChannelPressureEvent) InitWithChannelPressure(channel uint32, pressure uint32) AVMIDIChannelPressureEvent {
	rv := objc.Send[AVMIDIChannelPressureEvent](m.ID, objc.Sel("initWithChannel:pressure:"), channel, pressure)
	return rv
}

// The MIDI channel pressure.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelPressureEvent/pressure
func (m AVMIDIChannelPressureEvent) Pressure() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("pressure"))
	return rv
}
func (m AVMIDIChannelPressureEvent) SetPressure(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setPressure:"), value)
}
