// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMIDIChannelEvent] class.
var (
	_AVMIDIChannelEventClass     AVMIDIChannelEventClass
	_AVMIDIChannelEventClassOnce sync.Once
)

func getAVMIDIChannelEventClass() AVMIDIChannelEventClass {
	_AVMIDIChannelEventClassOnce.Do(func() {
		_AVMIDIChannelEventClass = AVMIDIChannelEventClass{class: objc.GetClass("AVMIDIChannelEvent")}
	})
	return _AVMIDIChannelEventClass
}

// GetAVMIDIChannelEventClass returns the class object for AVMIDIChannelEvent.
func GetAVMIDIChannelEventClass() AVMIDIChannelEventClass {
	return getAVMIDIChannelEventClass()
}

type AVMIDIChannelEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIChannelEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIChannelEventClass) Alloc() AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVMIDIChannelEvent.Data1]
//   - [AVMIDIChannelEvent.Data2]
//   - [AVMIDIChannelEvent.SetData1]
//   - [AVMIDIChannelEvent.SetData2]
//   - [AVMIDIChannelEvent.InitWithChannelStatusData1Data2]
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent
type AVMIDIChannelEvent struct {
	objectivec.Object
}

// AVMIDIChannelEventFromID constructs a [AVMIDIChannelEvent] from an objc.ID.
func AVMIDIChannelEventFromID(id objc.ID) AVMIDIChannelEvent {
	return AVMIDIChannelEvent{objectivec.Object{ID: id}}
}
// NOTE: AVMIDIChannelEvent struct embeds objectivec.Object (parent type unavailable) but
// IAVMIDIChannelEvent embeds the parent interface; skip compile-time assertion.

// An interface definition for the [AVMIDIChannelEvent] class.
//
// # Methods
//
//   - [IAVMIDIChannelEvent.Data1]
//   - [IAVMIDIChannelEvent.Data2]
//   - [IAVMIDIChannelEvent.SetData1]
//   - [IAVMIDIChannelEvent.SetData2]
//   - [IAVMIDIChannelEvent.InitWithChannelStatusData1Data2]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent
type IAVMIDIChannelEvent interface {
	IAVMusicEvent

	// Topic: Methods

	Data1() byte
	Data2() byte
	SetData1(data1 byte)
	SetData2(data2 byte)
	InitWithChannelStatusData1Data2(channel byte, status byte, data1 byte, data2 byte) AVMIDIChannelEvent
}

// Init initializes the instance.
func (m AVMIDIChannelEvent) Init() AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIChannelEvent) Autorelease() AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIChannelEvent creates a new AVMIDIChannelEvent instance.
func NewAVMIDIChannelEvent() AVMIDIChannelEvent {
	class := getAVMIDIChannelEventClass()
	rv := objc.Send[AVMIDIChannelEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/initWithChannel:status:data1:data2:
func NewMIDIChannelEventWithChannelStatusData1Data2(channel byte, status byte, data1 byte, data2 byte) AVMIDIChannelEvent {
	instance := getAVMIDIChannelEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithChannel:status:data1:data2:"), channel, status, data1, data2)
	return AVMIDIChannelEventFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/data1
func (m AVMIDIChannelEvent) Data1() byte {
	rv := objc.Send[byte](m.ID, objc.Sel("data1"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/data2
func (m AVMIDIChannelEvent) Data2() byte {
	rv := objc.Send[byte](m.ID, objc.Sel("data2"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/setData1:
func (m AVMIDIChannelEvent) SetData1(data1 byte) {
	objc.Send[objc.ID](m.ID, objc.Sel("setData1:"), data1)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/setData2:
func (m AVMIDIChannelEvent) SetData2(data2 byte) {
	objc.Send[objc.ID](m.ID, objc.Sel("setData2:"), data2)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIChannelEvent/initWithChannel:status:data1:data2:
func (m AVMIDIChannelEvent) InitWithChannelStatusData1Data2(channel byte, status byte, data1 byte, data2 byte) AVMIDIChannelEvent {
	rv := objc.Send[AVMIDIChannelEvent](m.ID, objc.Sel("initWithChannel:status:data1:data2:"), channel, status, data1, data2)
	return rv
}

