// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVExtendedTempoEvent] class.
var (
	_AVExtendedTempoEventClass     AVExtendedTempoEventClass
	_AVExtendedTempoEventClassOnce sync.Once
)

func getAVExtendedTempoEventClass() AVExtendedTempoEventClass {
	_AVExtendedTempoEventClassOnce.Do(func() {
		_AVExtendedTempoEventClass = AVExtendedTempoEventClass{class: objc.GetClass("AVExtendedTempoEvent")}
	})
	return _AVExtendedTempoEventClass
}

// GetAVExtendedTempoEventClass returns the class object for AVExtendedTempoEvent.
func GetAVExtendedTempoEventClass() AVExtendedTempoEventClass {
	return getAVExtendedTempoEventClass()
}

type AVExtendedTempoEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVExtendedTempoEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVExtendedTempoEventClass) Alloc() AVExtendedTempoEvent {
	rv := objc.Send[AVExtendedTempoEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a tempo change to a specific beats-per-minute
// value.
//
// # Creating a Tempo Event
//
//   - [AVExtendedTempoEvent.InitWithTempo]: Creates an extended tempo event.
//
// # Configuring a Tempo Event
//
//   - [AVExtendedTempoEvent.Tempo]: The tempo in beats per minute as a positive value.
//   - [AVExtendedTempoEvent.SetTempo]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent
type AVExtendedTempoEvent struct {
	AVMusicEvent
}

// AVExtendedTempoEventFromID constructs a [AVExtendedTempoEvent] from an objc.ID.
//
// An object that represents a tempo change to a specific beats-per-minute
// value.
func AVExtendedTempoEventFromID(id objc.ID) AVExtendedTempoEvent {
	return AVExtendedTempoEvent{AVMusicEvent: AVMusicEventFromID(id)}
}

// NOTE: AVExtendedTempoEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVExtendedTempoEvent] class.
//
// # Creating a Tempo Event
//
//   - [IAVExtendedTempoEvent.InitWithTempo]: Creates an extended tempo event.
//
// # Configuring a Tempo Event
//
//   - [IAVExtendedTempoEvent.Tempo]: The tempo in beats per minute as a positive value.
//   - [IAVExtendedTempoEvent.SetTempo]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent
type IAVExtendedTempoEvent interface {
	IAVMusicEvent

	// Topic: Creating a Tempo Event

	// Creates an extended tempo event.
	InitWithTempo(tempo float64) AVExtendedTempoEvent

	// Topic: Configuring a Tempo Event

	// The tempo in beats per minute as a positive value.
	Tempo() float64
	SetTempo(value float64)
}

// Init initializes the instance.
func (e AVExtendedTempoEvent) Init() AVExtendedTempoEvent {
	rv := objc.Send[AVExtendedTempoEvent](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e AVExtendedTempoEvent) Autorelease() AVExtendedTempoEvent {
	rv := objc.Send[AVExtendedTempoEvent](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExtendedTempoEvent creates a new AVExtendedTempoEvent instance.
func NewAVExtendedTempoEvent() AVExtendedTempoEvent {
	class := getAVExtendedTempoEventClass()
	rv := objc.Send[AVExtendedTempoEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an extended tempo event.
//
// tempo: The tempo in beats per minute as a positive value.
//
// # Discussion
//
// The new tempo begins at the timestamp for this event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent/init(tempo:)
func NewExtendedTempoEventWithTempo(tempo float64) AVExtendedTempoEvent {
	instance := getAVExtendedTempoEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTempo:"), tempo)
	return AVExtendedTempoEventFromID(rv)
}

// Creates an extended tempo event.
//
// tempo: The tempo in beats per minute as a positive value.
//
// # Discussion
//
// The new tempo begins at the timestamp for this event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent/init(tempo:)
func (e AVExtendedTempoEvent) InitWithTempo(tempo float64) AVExtendedTempoEvent {
	rv := objc.Send[AVExtendedTempoEvent](e.ID, objc.Sel("initWithTempo:"), tempo)
	return rv
}

// The tempo in beats per minute as a positive value.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVExtendedTempoEvent/tempo
func (e AVExtendedTempoEvent) Tempo() float64 {
	rv := objc.Send[float64](e.ID, objc.Sel("tempo"))
	return rv
}
func (e AVExtendedTempoEvent) SetTempo(value float64) {
	objc.Send[struct{}](e.ID, objc.Sel("setTempo:"), value)
}
