// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDIPolyPressureEvent] class.
var (
	_AVMIDIPolyPressureEventClass     AVMIDIPolyPressureEventClass
	_AVMIDIPolyPressureEventClassOnce sync.Once
)

func getAVMIDIPolyPressureEventClass() AVMIDIPolyPressureEventClass {
	_AVMIDIPolyPressureEventClassOnce.Do(func() {
		_AVMIDIPolyPressureEventClass = AVMIDIPolyPressureEventClass{class: objc.GetClass("AVMIDIPolyPressureEvent")}
	})
	return _AVMIDIPolyPressureEventClass
}

// GetAVMIDIPolyPressureEventClass returns the class object for AVMIDIPolyPressureEvent.
func GetAVMIDIPolyPressureEventClass() AVMIDIPolyPressureEventClass {
	return getAVMIDIPolyPressureEventClass()
}

type AVMIDIPolyPressureEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIPolyPressureEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIPolyPressureEventClass) Alloc() AVMIDIPolyPressureEvent {
	rv := objc.Send[AVMIDIPolyPressureEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI poly or key pressure event.
//
// # Creating a Poly Pressure Event
//
//   - [AVMIDIPolyPressureEvent.InitWithChannelKeyPressure]: Creates an event with a channel, MIDI key number, and a key pressure value.
//
// # Configuring a Poly Pressure Event
//
//   - [AVMIDIPolyPressureEvent.Key]: The MIDI key number.
//   - [AVMIDIPolyPressureEvent.SetKey]
//   - [AVMIDIPolyPressureEvent.Pressure]: The poly pressure value for the requested key.
//   - [AVMIDIPolyPressureEvent.SetPressure]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent
type AVMIDIPolyPressureEvent struct {
	AVMIDIChannelEvent
}

// AVMIDIPolyPressureEventFromID constructs a [AVMIDIPolyPressureEvent] from an objc.ID.
//
// An object that represents a MIDI poly or key pressure event.
func AVMIDIPolyPressureEventFromID(id objc.ID) AVMIDIPolyPressureEvent {
	return AVMIDIPolyPressureEvent{AVMIDIChannelEvent: AVMIDIChannelEventFromID(id)}
}
// NOTE: AVMIDIPolyPressureEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDIPolyPressureEvent] class.
//
// # Creating a Poly Pressure Event
//
//   - [IAVMIDIPolyPressureEvent.InitWithChannelKeyPressure]: Creates an event with a channel, MIDI key number, and a key pressure value.
//
// # Configuring a Poly Pressure Event
//
//   - [IAVMIDIPolyPressureEvent.Key]: The MIDI key number.
//   - [IAVMIDIPolyPressureEvent.SetKey]
//   - [IAVMIDIPolyPressureEvent.Pressure]: The poly pressure value for the requested key.
//   - [IAVMIDIPolyPressureEvent.SetPressure]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent
type IAVMIDIPolyPressureEvent interface {
	IAVMIDIChannelEvent

	// Topic: Creating a Poly Pressure Event

	// Creates an event with a channel, MIDI key number, and a key pressure value.
	InitWithChannelKeyPressure(channel uint32, key uint32, pressure uint32) AVMIDIPolyPressureEvent

	// Topic: Configuring a Poly Pressure Event

	// The MIDI key number.
	Key() uint32
	SetKey(value uint32)
	// The poly pressure value for the requested key.
	Pressure() uint32
	SetPressure(value uint32)
}

// Init initializes the instance.
func (m AVMIDIPolyPressureEvent) Init() AVMIDIPolyPressureEvent {
	rv := objc.Send[AVMIDIPolyPressureEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIPolyPressureEvent) Autorelease() AVMIDIPolyPressureEvent {
	rv := objc.Send[AVMIDIPolyPressureEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIPolyPressureEvent creates a new AVMIDIPolyPressureEvent instance.
func NewAVMIDIPolyPressureEvent() AVMIDIPolyPressureEvent {
	class := getAVMIDIPolyPressureEventClass()
	rv := objc.Send[AVMIDIPolyPressureEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a channel, MIDI key number, and a key pressure value.
//
// channel: The MIDI channel for the message, between `0` and `15`.
//
// key: The MIDI key number to apply the pressure to.
//
// pressure: The poly pressure value.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/init(channel:key:pressure:)
func NewMIDIPolyPressureEventWithChannelKeyPressure(channel uint32, key uint32, pressure uint32) AVMIDIPolyPressureEvent {
	instance := getAVMIDIPolyPressureEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:key:pressure:"), channel, key, pressure)
	return AVMIDIPolyPressureEventFromID(rv)
}

// Creates an event with a channel, MIDI key number, and a key pressure value.
//
// channel: The MIDI channel for the message, between `0` and `15`.
//
// key: The MIDI key number to apply the pressure to.
//
// pressure: The poly pressure value.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/init(channel:key:pressure:)
func (m AVMIDIPolyPressureEvent) InitWithChannelKeyPressure(channel uint32, key uint32, pressure uint32) AVMIDIPolyPressureEvent {
	rv := objc.Send[AVMIDIPolyPressureEvent](m.ID, objc.Sel("initWithChannel:key:pressure:"), channel, key, pressure)
	return rv
}

// The MIDI key number.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/key
func (m AVMIDIPolyPressureEvent) Key() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("key"))
	return rv
}
func (m AVMIDIPolyPressureEvent) SetKey(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setKey:"), value)
}
// The poly pressure value for the requested key.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPolyPressureEvent/pressure
func (m AVMIDIPolyPressureEvent) Pressure() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("pressure"))
	return rv
}
func (m AVMIDIPolyPressureEvent) SetPressure(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setPressure:"), value)
}

