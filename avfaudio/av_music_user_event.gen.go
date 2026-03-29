// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVMusicUserEvent] class.
var (
	_AVMusicUserEventClass     AVMusicUserEventClass
	_AVMusicUserEventClassOnce sync.Once
)

func getAVMusicUserEventClass() AVMusicUserEventClass {
	_AVMusicUserEventClassOnce.Do(func() {
		_AVMusicUserEventClass = AVMusicUserEventClass{class: objc.GetClass("AVMusicUserEvent")}
	})
	return _AVMusicUserEventClass
}

// GetAVMusicUserEventClass returns the class object for AVMusicUserEvent.
func GetAVMusicUserEventClass() AVMusicUserEventClass {
	return getAVMusicUserEventClass()
}

type AVMusicUserEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMusicUserEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMusicUserEventClass) Alloc() AVMusicUserEvent {
	rv := objc.Send[AVMusicUserEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a custom user message.
//
// # Overview
// 
// When playback of an [AVMusicTrack] reaches this event, the system calls the
// track’s callback. You can’t modify the size and contents of an
// [AVMusicUserEvent] once you create it.
//
// # Creating a User Event
//
//   - [AVMusicUserEvent.InitWithData]: Creates a user event with the data you specify.
//
// # Inspecting a User Event
//
//   - [AVMusicUserEvent.SizeInBytes]: The size of the data that the user event represents.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent
type AVMusicUserEvent struct {
	AVMusicEvent
}

// AVMusicUserEventFromID constructs a [AVMusicUserEvent] from an objc.ID.
//
// An object that represents a custom user message.
func AVMusicUserEventFromID(id objc.ID) AVMusicUserEvent {
	return AVMusicUserEvent{AVMusicEvent: AVMusicEventFromID(id)}
}
// NOTE: AVMusicUserEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMusicUserEvent] class.
//
// # Creating a User Event
//
//   - [IAVMusicUserEvent.InitWithData]: Creates a user event with the data you specify.
//
// # Inspecting a User Event
//
//   - [IAVMusicUserEvent.SizeInBytes]: The size of the data that the user event represents.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent
type IAVMusicUserEvent interface {
	IAVMusicEvent

	// Topic: Creating a User Event

	// Creates a user event with the data you specify.
	InitWithData(data foundation.INSData) AVMusicUserEvent

	// Topic: Inspecting a User Event

	// The size of the data that the user event represents.
	SizeInBytes() uint32
}

// Init initializes the instance.
func (m AVMusicUserEvent) Init() AVMusicUserEvent {
	rv := objc.Send[AVMusicUserEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMusicUserEvent) Autorelease() AVMusicUserEvent {
	rv := objc.Send[AVMusicUserEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMusicUserEvent creates a new AVMusicUserEvent instance.
func NewAVMusicUserEvent() AVMusicUserEvent {
	class := getAVMusicUserEventClass()
	rv := objc.Send[AVMusicUserEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a user event with the data you specify.
//
// data: The contents a music track returns on callback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/init(data:)
func NewMusicUserEventWithData(data foundation.INSData) AVMusicUserEvent {
	instance := getAVMusicUserEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return AVMusicUserEventFromID(rv)
}

// Creates a user event with the data you specify.
//
// data: The contents a music track returns on callback.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/init(data:)
func (m AVMusicUserEvent) InitWithData(data foundation.INSData) AVMusicUserEvent {
	rv := objc.Send[AVMusicUserEvent](m.ID, objc.Sel("initWithData:"), data)
	return rv
}

// The size of the data that the user event represents.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicUserEvent/sizeInBytes
func (m AVMusicUserEvent) SizeInBytes() uint32 {
	rv := objc.Send[uint32](m.ID, objc.Sel("sizeInBytes"))
	return rv
}

