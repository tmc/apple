// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMIDISysexEvent] class.
var (
	_AVMIDISysexEventClass     AVMIDISysexEventClass
	_AVMIDISysexEventClassOnce sync.Once
)

func getAVMIDISysexEventClass() AVMIDISysexEventClass {
	_AVMIDISysexEventClassOnce.Do(func() {
		_AVMIDISysexEventClass = AVMIDISysexEventClass{class: objc.GetClass("AVMIDISysexEvent")}
	})
	return _AVMIDISysexEventClass
}

// GetAVMIDISysexEventClass returns the class object for AVMIDISysexEvent.
func GetAVMIDISysexEventClass() AVMIDISysexEventClass {
	return getAVMIDISysexEventClass()
}

type AVMIDISysexEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDISysexEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDISysexEventClass) Alloc() AVMIDISysexEvent {
	rv := objc.Send[AVMIDISysexEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a MIDI system exclusive message.
//
// # Overview
// 
// You can’t modify the size and contents of this event once you create it.
//
// # Creates a System Event
//
//   - [AVMIDISysexEvent.InitWithData]: Creates a system event with the data you specify.
//
// # Getting the Size of the Event
//
//   - [AVMIDISysexEvent.SizeInBytes]: The size of the data that this event contains.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent
type AVMIDISysexEvent struct {
	AVMusicEvent
}

// AVMIDISysexEventFromID constructs a [AVMIDISysexEvent] from an objc.ID.
//
// An object that represents a MIDI system exclusive message.
func AVMIDISysexEventFromID(id objc.ID) AVMIDISysexEvent {
	return AVMIDISysexEvent{AVMusicEvent: AVMusicEventFromID(id)}
}
// NOTE: AVMIDISysexEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMIDISysexEvent] class.
//
// # Creates a System Event
//
//   - [IAVMIDISysexEvent.InitWithData]: Creates a system event with the data you specify.
//
// # Getting the Size of the Event
//
//   - [IAVMIDISysexEvent.SizeInBytes]: The size of the data that this event contains.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent
type IAVMIDISysexEvent interface {
	IAVMusicEvent

	// Topic: Creates a System Event

	// Creates a system event with the data you specify.
	InitWithData(data foundation.INSData) AVMIDISysexEvent

	// Topic: Getting the Size of the Event

	// The size of the data that this event contains.
	SizeInBytes() uint32
}

// Init initializes the instance.
func (m AVMIDISysexEvent) Init() AVMIDISysexEvent {
	rv := objc.Send[AVMIDISysexEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDISysexEvent) Autorelease() AVMIDISysexEvent {
	rv := objc.Send[AVMIDISysexEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDISysexEvent creates a new AVMIDISysexEvent instance.
func NewAVMIDISysexEvent() AVMIDISysexEvent {
	class := getAVMIDISysexEventClass()
	rv := objc.Send[AVMIDISysexEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a system event with the data you specify.
//
// data: The data that contains the contents of the system event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/init(data:)
func NewMIDISysexEventWithData(data foundation.INSData) AVMIDISysexEvent {
	instance := getAVMIDISysexEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return AVMIDISysexEventFromID(rv)
}

// Creates a system event with the data you specify.
//
// data: The data that contains the contents of the system event.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/init(data:)
func (m AVMIDISysexEvent) InitWithData(data foundation.INSData) AVMIDISysexEvent {
	rv := objc.Send[AVMIDISysexEvent](m.ID, objc.Sel("initWithData:"), data)
	return rv
}

// The size of the data that this event contains.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDISysexEvent/sizeInBytes
func (m AVMIDISysexEvent) SizeInBytes() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("sizeInBytes"))
	return rv
}

