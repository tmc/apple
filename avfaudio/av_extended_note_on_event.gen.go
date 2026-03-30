// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVExtendedNoteOnEvent] class.
var (
	_AVExtendedNoteOnEventClass     AVExtendedNoteOnEventClass
	_AVExtendedNoteOnEventClassOnce sync.Once
)

func getAVExtendedNoteOnEventClass() AVExtendedNoteOnEventClass {
	_AVExtendedNoteOnEventClassOnce.Do(func() {
		_AVExtendedNoteOnEventClass = AVExtendedNoteOnEventClass{class: objc.GetClass("AVExtendedNoteOnEvent")}
	})
	return _AVExtendedNoteOnEventClass
}

// GetAVExtendedNoteOnEventClass returns the class object for AVExtendedNoteOnEvent.
func GetAVExtendedNoteOnEventClass() AVExtendedNoteOnEventClass {
	return getAVExtendedNoteOnEventClass()
}

type AVExtendedNoteOnEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVExtendedNoteOnEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExtendedNoteOnEventClass) Alloc() AVExtendedNoteOnEvent {
	rv := objc.Send[AVExtendedNoteOnEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a custom extension of a MIDI note on event.
//
// # Overview
//
// Use this to allow an app to trigger a custom note on event on one of
// several Apple audio units that support it. The floating point note and
// velocity numbers allow for optional fractional control of the note’s
// runtime properties that the system modulates by those inputs. This event
// supports the possibility of an audio unit with more than the standard 16
// MIDI channels.
//
// # Creating a Note On Event
//
//   - [AVExtendedNoteOnEvent.InitWithMIDINoteVelocityGroupIDDuration]: Creates an event with a MIDI note, velocity, group identifier, and duration.
//   - [AVExtendedNoteOnEvent.InitWithMIDINoteVelocityInstrumentIDGroupIDDuration]: Creates a note on event with the default instrument.
//
// # Configuring a Note On Event
//
//   - [AVExtendedNoteOnEvent.MidiNote]: The MIDI note number.
//   - [AVExtendedNoteOnEvent.SetMidiNote]
//   - [AVExtendedNoteOnEvent.Velocity]: The MDI velocity.
//   - [AVExtendedNoteOnEvent.SetVelocity]
//   - [AVExtendedNoteOnEvent.InstrumentID]: The instrument identifier.
//   - [AVExtendedNoteOnEvent.SetInstrumentID]
//   - [AVExtendedNoteOnEvent.GroupID]: The audio unit channel that handles the event.
//   - [AVExtendedNoteOnEvent.SetGroupID]
//   - [AVExtendedNoteOnEvent.Duration]: The duration of the event, in beats.
//   - [AVExtendedNoteOnEvent.SetDuration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent
type AVExtendedNoteOnEvent struct {
	AVMusicEvent
}

// AVExtendedNoteOnEventFromID constructs a [AVExtendedNoteOnEvent] from an objc.ID.
//
// An object that represents a custom extension of a MIDI note on event.
func AVExtendedNoteOnEventFromID(id objc.ID) AVExtendedNoteOnEvent {
	return AVExtendedNoteOnEvent{AVMusicEvent: AVMusicEventFromID(id)}
}

// NOTE: AVExtendedNoteOnEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVExtendedNoteOnEvent] class.
//
// # Creating a Note On Event
//
//   - [IAVExtendedNoteOnEvent.InitWithMIDINoteVelocityGroupIDDuration]: Creates an event with a MIDI note, velocity, group identifier, and duration.
//   - [IAVExtendedNoteOnEvent.InitWithMIDINoteVelocityInstrumentIDGroupIDDuration]: Creates a note on event with the default instrument.
//
// # Configuring a Note On Event
//
//   - [IAVExtendedNoteOnEvent.MidiNote]: The MIDI note number.
//   - [IAVExtendedNoteOnEvent.SetMidiNote]
//   - [IAVExtendedNoteOnEvent.Velocity]: The MDI velocity.
//   - [IAVExtendedNoteOnEvent.SetVelocity]
//   - [IAVExtendedNoteOnEvent.InstrumentID]: The instrument identifier.
//   - [IAVExtendedNoteOnEvent.SetInstrumentID]
//   - [IAVExtendedNoteOnEvent.GroupID]: The audio unit channel that handles the event.
//   - [IAVExtendedNoteOnEvent.SetGroupID]
//   - [IAVExtendedNoteOnEvent.Duration]: The duration of the event, in beats.
//   - [IAVExtendedNoteOnEvent.SetDuration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent
type IAVExtendedNoteOnEvent interface {
	IAVMusicEvent

	// Topic: Creating a Note On Event

	// Creates an event with a MIDI note, velocity, group identifier, and duration.
	InitWithMIDINoteVelocityGroupIDDuration(midiNote float32, velocity float32, groupID uint32, duration AVMusicTimeStamp) AVExtendedNoteOnEvent
	// Creates a note on event with the default instrument.
	InitWithMIDINoteVelocityInstrumentIDGroupIDDuration(midiNote float32, velocity float32, instrumentID uint32, groupID uint32, duration AVMusicTimeStamp) AVExtendedNoteOnEvent

	// Topic: Configuring a Note On Event

	// The MIDI note number.
	MidiNote() float32
	SetMidiNote(value float32)
	// The MDI velocity.
	Velocity() float32
	SetVelocity(value float32)
	// The instrument identifier.
	InstrumentID() uint32
	SetInstrumentID(value uint32)
	// The audio unit channel that handles the event.
	GroupID() uint32
	SetGroupID(value uint32)
	// The duration of the event, in beats.
	Duration() AVMusicTimeStamp
	SetDuration(value AVMusicTimeStamp)
}

// Init initializes the instance.
func (e AVExtendedNoteOnEvent) Init() AVExtendedNoteOnEvent {
	rv := objc.Send[AVExtendedNoteOnEvent](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExtendedNoteOnEvent) Autorelease() AVExtendedNoteOnEvent {
	rv := objc.Send[AVExtendedNoteOnEvent](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExtendedNoteOnEvent creates a new AVExtendedNoteOnEvent instance.
func NewAVExtendedNoteOnEvent() AVExtendedNoteOnEvent {
	class := getAVExtendedNoteOnEventClass()
	rv := objc.Send[AVExtendedNoteOnEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event with a MIDI note, velocity, group identifier, and
// duration.
//
// midiNote: The MIDI note number.
//
// velocity: The MIDI velocity.
//
// groupID: The identifier that represents the audio unit channel that handles the
// event.
//
// duration: The duration of the event, in beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/init(midiNote:velocity:groupID:duration:)
func NewExtendedNoteOnEventWithMIDINoteVelocityGroupIDDuration(midiNote float32, velocity float32, groupID uint32, duration AVMusicTimeStamp) AVExtendedNoteOnEvent {
	instance := getAVExtendedNoteOnEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMIDINote:velocity:groupID:duration:"), midiNote, velocity, groupID, duration)
	return AVExtendedNoteOnEventFromID(rv)
}

// Creates a note on event with the default instrument.
//
// midiNote: The MIDI note number.
//
// velocity: The MIDI velocity.
//
// instrumentID: The default instrument.
//
// groupID: The identifier that represents the audio unit channel that handles the
// event.
//
// duration: The duration of the event, in beats.
//
// # Discussion
//
// Use [defaultInstrument] when you set `instrumentID`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/init(midiNote:velocity:instrumentID:groupID:duration:)
//
// [defaultInstrument]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/defaultInstrument
func NewExtendedNoteOnEventWithMIDINoteVelocityInstrumentIDGroupIDDuration(midiNote float32, velocity float32, instrumentID uint32, groupID uint32, duration AVMusicTimeStamp) AVExtendedNoteOnEvent {
	instance := getAVExtendedNoteOnEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMIDINote:velocity:instrumentID:groupID:duration:"), midiNote, velocity, instrumentID, groupID, duration)
	return AVExtendedNoteOnEventFromID(rv)
}

// Creates an event with a MIDI note, velocity, group identifier, and
// duration.
//
// midiNote: The MIDI note number.
//
// velocity: The MIDI velocity.
//
// groupID: The identifier that represents the audio unit channel that handles the
// event.
//
// duration: The duration of the event, in beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/init(midiNote:velocity:groupID:duration:)
func (e AVExtendedNoteOnEvent) InitWithMIDINoteVelocityGroupIDDuration(midiNote float32, velocity float32, groupID uint32, duration AVMusicTimeStamp) AVExtendedNoteOnEvent {
	rv := objc.Send[AVExtendedNoteOnEvent](e.ID, objc.Sel("initWithMIDINote:velocity:groupID:duration:"), midiNote, velocity, groupID, duration)
	return rv
}

// Creates a note on event with the default instrument.
//
// midiNote: The MIDI note number.
//
// velocity: The MIDI velocity.
//
// instrumentID: The default instrument.
//
// groupID: The identifier that represents the audio unit channel that handles the
// event.
//
// duration: The duration of the event, in beats.
//
// # Discussion
//
// Use [defaultInstrument] when you set `instrumentID`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/init(midiNote:velocity:instrumentID:groupID:duration:)
//
// [defaultInstrument]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/defaultInstrument
func (e AVExtendedNoteOnEvent) InitWithMIDINoteVelocityInstrumentIDGroupIDDuration(midiNote float32, velocity float32, instrumentID uint32, groupID uint32, duration AVMusicTimeStamp) AVExtendedNoteOnEvent {
	rv := objc.Send[AVExtendedNoteOnEvent](e.ID, objc.Sel("initWithMIDINote:velocity:instrumentID:groupID:duration:"), midiNote, velocity, instrumentID, groupID, duration)
	return rv
}

// The MIDI note number.
//
// # Discussion
//
// If the instrument within the [AVMusicTrack] destination audio unit supports
// fractional values, you use this to generate arbitrary tunings. The valid
// range of values depends on the destination audio unit, and is usually
// between `0.0` and `127.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/midiNote
func (e AVExtendedNoteOnEvent) MidiNote() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("midiNote"))
	return rv
}
func (e AVExtendedNoteOnEvent) SetMidiNote(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setMidiNote:"), value)
}

// The MDI velocity.
//
// # Discussion
//
// If the instrument in the [AVMusicTrack] destination audio unit supports
// fractional values, use this to generate precise changes in gain and other
// values. The valid range of values depend on the destination audio unit, and
// is usually between `0.0` and `127.0`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/velocity
func (e AVExtendedNoteOnEvent) Velocity() float32 {
	rv := objc.Send[float32](e.ID, objc.Sel("velocity"))
	return rv
}
func (e AVExtendedNoteOnEvent) SetVelocity(value float32) {
	objc.Send[struct{}](e.ID, objc.Sel("setVelocity:"), value)
}

// The instrument identifier.
//
// # Discussion
//
// Set this value to [defaultInstrument].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/instrumentID
//
// [defaultInstrument]: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/defaultInstrument
func (e AVExtendedNoteOnEvent) InstrumentID() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("instrumentID"))
	return rv
}
func (e AVExtendedNoteOnEvent) SetInstrumentID(value uint32) {
	objc.Send[struct{}](e.ID, objc.Sel("setInstrumentID:"), value)
}

// The audio unit channel that handles the event.
//
// # Discussion
//
// The valid range of values are between `0` and `15`, but can be higher if
// the [AVMusicTrack] destination audio unit supports more channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/groupID
func (e AVExtendedNoteOnEvent) GroupID() uint32 {
	rv := objc.Send[uint32](e.ID, objc.Sel("groupID"))
	return rv
}
func (e AVExtendedNoteOnEvent) SetGroupID(value uint32) {
	objc.Send[struct{}](e.ID, objc.Sel("setGroupID:"), value)
}

// The duration of the event, in beats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedNoteOnEvent/duration
func (e AVExtendedNoteOnEvent) Duration() AVMusicTimeStamp {
	rv := objc.Send[AVMusicTimeStamp](e.ID, objc.Sel("duration"))
	return AVMusicTimeStamp(rv)
}
func (e AVExtendedNoteOnEvent) SetDuration(value AVMusicTimeStamp) {
	objc.Send[struct{}](e.ID, objc.Sel("setDuration:"), value)
}
