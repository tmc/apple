// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMIDINoteEvent] class.
var (
	_AVMIDINoteEventClass     AVMIDINoteEventClass
	_AVMIDINoteEventClassOnce sync.Once
)

func getAVMIDINoteEventClass() AVMIDINoteEventClass {
	_AVMIDINoteEventClassOnce.Do(func() {
		_AVMIDINoteEventClass = AVMIDINoteEventClass{class: objc.GetClass("AVMIDINoteEvent")}
	})
	return _AVMIDINoteEventClass
}

// GetAVMIDINoteEventClass returns the class object for AVMIDINoteEvent.
func GetAVMIDINoteEventClass() AVMIDINoteEventClass {
	return getAVMIDINoteEventClass()
}

type AVMIDINoteEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDINoteEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDINoteEventClass) Alloc() AVMIDINoteEvent {
	rv := objc.Send[AVMIDINoteEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents MIDI note on or off messages.
//
// # Creating a MIDI Note Event
//
//   - [AVMIDINoteEvent.InitWithChannelKeyVelocityDuration]: Creates an event with a MIDI channel, key number, velocity, and duration.
//
// # Configuring a MIDI Note Event
//
//   - [AVMIDINoteEvent.Channel]: The MIDI channel.
//   - [AVMIDINoteEvent.SetChannel]
//   - [AVMIDINoteEvent.Key]: The MIDI key number.
//   - [AVMIDINoteEvent.SetKey]
//   - [AVMIDINoteEvent.Velocity]: The MIDI velocity.
//   - [AVMIDINoteEvent.SetVelocity]
//   - [AVMIDINoteEvent.Duration]: The duration for the note, in beats.
//   - [AVMIDINoteEvent.SetDuration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent
type AVMIDINoteEvent struct {
	AVMusicEvent
}

// AVMIDINoteEventFromID constructs a [AVMIDINoteEvent] from an objc.ID.
//
// An object that represents MIDI note on or off messages.
func AVMIDINoteEventFromID(id objc.ID) AVMIDINoteEvent {
	return AVMIDINoteEvent{AVMusicEvent: AVMusicEventFromID(id)}
}

// NOTE: AVMIDINoteEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDINoteEvent] class.
//
// # Creating a MIDI Note Event
//
//   - [IAVMIDINoteEvent.InitWithChannelKeyVelocityDuration]: Creates an event with a MIDI channel, key number, velocity, and duration.
//
// # Configuring a MIDI Note Event
//
//   - [IAVMIDINoteEvent.Channel]: The MIDI channel.
//   - [IAVMIDINoteEvent.SetChannel]
//   - [IAVMIDINoteEvent.Key]: The MIDI key number.
//   - [IAVMIDINoteEvent.SetKey]
//   - [IAVMIDINoteEvent.Velocity]: The MIDI velocity.
//   - [IAVMIDINoteEvent.SetVelocity]
//   - [IAVMIDINoteEvent.Duration]: The duration for the note, in beats.
//   - [IAVMIDINoteEvent.SetDuration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent
type IAVMIDINoteEvent interface {
	IAVMusicEvent

	// Topic: Creating a MIDI Note Event

	// Creates an event with a MIDI channel, key number, velocity, and duration.
	InitWithChannelKeyVelocityDuration(channel uint32, keyNum uint32, velocity uint32, duration AVMusicTimeStamp) AVMIDINoteEvent

	// Topic: Configuring a MIDI Note Event

	// The MIDI channel.
	Channel() uint32
	SetChannel(value uint32)
	// The MIDI key number.
	Key() uint32
	SetKey(value uint32)
	// The MIDI velocity.
	Velocity() uint32
	SetVelocity(value uint32)
	// The duration for the note, in beats.
	Duration() AVMusicTimeStamp
	SetDuration(value AVMusicTimeStamp)
}

// Init initializes the instance.
func (m AVMIDINoteEvent) Init() AVMIDINoteEvent {
	rv := objc.Send[AVMIDINoteEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDINoteEvent) Autorelease() AVMIDINoteEvent {
	rv := objc.Send[AVMIDINoteEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDINoteEvent creates a new AVMIDINoteEvent instance.
func NewAVMIDINoteEvent() AVMIDINoteEvent {
	class := getAVMIDINoteEventClass()
	rv := objc.Send[AVMIDINoteEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a MIDI channel, key number, velocity, and duration.
//
// channel: The MIDI channel, between `0` and `15`.
//
// keyNum: The MIDI key number, between `0` and `127`.
//
// velocity: The MIDI velocity, between `0` and `127`.
//
// duration: The duration for this note, in beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/init(channel:key:velocity:duration:)
func NewMIDINoteEventWithChannelKeyVelocityDuration(channel uint32, keyNum uint32, velocity uint32, duration AVMusicTimeStamp) AVMIDINoteEvent {
	instance := getAVMIDINoteEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:key:velocity:duration:"), channel, keyNum, velocity, duration)
	return AVMIDINoteEventFromID(rv)
}

// Creates an event with a MIDI channel, key number, velocity, and duration.
//
// channel: The MIDI channel, between `0` and `15`.
//
// keyNum: The MIDI key number, between `0` and `127`.
//
// velocity: The MIDI velocity, between `0` and `127`.
//
// duration: The duration for this note, in beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/init(channel:key:velocity:duration:)
func (m AVMIDINoteEvent) InitWithChannelKeyVelocityDuration(channel uint32, keyNum uint32, velocity uint32, duration AVMusicTimeStamp) AVMIDINoteEvent {
	rv := objc.Send[AVMIDINoteEvent](m.ID, objc.Sel("initWithChannel:key:velocity:duration:"), channel, keyNum, velocity, duration)
	return rv
}

// The MIDI channel.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/channel
func (m AVMIDINoteEvent) Channel() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("channel"))
	return rv
}
func (m AVMIDINoteEvent) SetChannel(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setChannel:"), value)
}

// The MIDI key number.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/key
func (m AVMIDINoteEvent) Key() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("key"))
	return rv
}
func (m AVMIDINoteEvent) SetKey(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setKey:"), value)
}

// The MIDI velocity.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/velocity
func (m AVMIDINoteEvent) Velocity() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("velocity"))
	return rv
}
func (m AVMIDINoteEvent) SetVelocity(value uint32) {
	objc.Send[struct{}](m.ID, objc.Sel("setVelocity:"), value)
}

// The duration for the note, in beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDINoteEvent/duration
func (m AVMIDINoteEvent) Duration() AVMusicTimeStamp {
	rv := objc.Send[AVMusicTimeStamp](m.ID, objc.Sel("duration"))
	return AVMusicTimeStamp(rv)
}
func (m AVMIDINoteEvent) SetDuration(value AVMusicTimeStamp) {
	objc.Send[struct{}](m.ID, objc.Sel("setDuration:"), value)
}
