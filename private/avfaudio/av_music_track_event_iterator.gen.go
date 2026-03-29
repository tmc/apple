// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMusicTrackEventIterator] class.
var (
	_AVMusicTrackEventIteratorClass     AVMusicTrackEventIteratorClass
	_AVMusicTrackEventIteratorClassOnce sync.Once
)

func getAVMusicTrackEventIteratorClass() AVMusicTrackEventIteratorClass {
	_AVMusicTrackEventIteratorClassOnce.Do(func() {
		_AVMusicTrackEventIteratorClass = AVMusicTrackEventIteratorClass{class: objc.GetClass("AVMusicTrackEventIterator")}
	})
	return _AVMusicTrackEventIteratorClass
}

// GetAVMusicTrackEventIteratorClass returns the class object for AVMusicTrackEventIterator.
func GetAVMusicTrackEventIteratorClass() AVMusicTrackEventIteratorClass {
	return getAVMusicTrackEventIteratorClass()
}

type AVMusicTrackEventIteratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMusicTrackEventIteratorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMusicTrackEventIteratorClass) Alloc() AVMusicTrackEventIterator {
	rv := objc.Send[AVMusicTrackEventIterator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVMusicTrackEventIterator.DeleteEvent]
//   - [AVMusicTrackEventIterator.GetEventInfoOutEventTypeEventDataDataSize]
//   - [AVMusicTrackEventIterator.HasCurrentEvent]
//   - [AVMusicTrackEventIterator.HasNextEvent]
//   - [AVMusicTrackEventIterator.HasPreviousEvent]
//   - [AVMusicTrackEventIterator.NextEvent]
//   - [AVMusicTrackEventIterator.PreviousEvent]
//   - [AVMusicTrackEventIterator.Seek]
//   - [AVMusicTrackEventIterator.SetEventInfoData]
//   - [AVMusicTrackEventIterator.SetEventTime]
//   - [AVMusicTrackEventIterator.InitWithImpl]
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator
type AVMusicTrackEventIterator struct {
	objectivec.Object
}

// AVMusicTrackEventIteratorFromID constructs a [AVMusicTrackEventIterator] from an objc.ID.
func AVMusicTrackEventIteratorFromID(id objc.ID) AVMusicTrackEventIterator {
	return AVMusicTrackEventIterator{objectivec.Object{ID: id}}
}
// Ensure AVMusicTrackEventIterator implements IAVMusicTrackEventIterator.
var _ IAVMusicTrackEventIterator = AVMusicTrackEventIterator{}

// An interface definition for the [AVMusicTrackEventIterator] class.
//
// # Methods
//
//   - [IAVMusicTrackEventIterator.DeleteEvent]
//   - [IAVMusicTrackEventIterator.GetEventInfoOutEventTypeEventDataDataSize]
//   - [IAVMusicTrackEventIterator.HasCurrentEvent]
//   - [IAVMusicTrackEventIterator.HasNextEvent]
//   - [IAVMusicTrackEventIterator.HasPreviousEvent]
//   - [IAVMusicTrackEventIterator.NextEvent]
//   - [IAVMusicTrackEventIterator.PreviousEvent]
//   - [IAVMusicTrackEventIterator.Seek]
//   - [IAVMusicTrackEventIterator.SetEventInfoData]
//   - [IAVMusicTrackEventIterator.SetEventTime]
//   - [IAVMusicTrackEventIterator.InitWithImpl]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator
type IAVMusicTrackEventIterator interface {
	objectivec.IObject

	// Topic: Methods

	DeleteEvent()
	GetEventInfoOutEventTypeEventDataDataSize(info []float64, type_ unsafe.Pointer, data unsafe.Pointer, size unsafe.Pointer)
	HasCurrentEvent() bool
	HasNextEvent() bool
	HasPreviousEvent() bool
	NextEvent() int
	PreviousEvent() int
	Seek(seek float64)
	SetEventInfoData(info uint32, data unsafe.Pointer) bool
	SetEventTime(time float64) bool
	InitWithImpl(impl unsafe.Pointer) AVMusicTrackEventIterator
}

// Init initializes the instance.
func (m AVMusicTrackEventIterator) Init() AVMusicTrackEventIterator {
	rv := objc.Send[AVMusicTrackEventIterator](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMusicTrackEventIterator) Autorelease() AVMusicTrackEventIterator {
	rv := objc.Send[AVMusicTrackEventIterator](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMusicTrackEventIterator creates a new AVMusicTrackEventIterator instance.
func NewAVMusicTrackEventIterator() AVMusicTrackEventIterator {
	class := getAVMusicTrackEventIteratorClass()
	rv := objc.Send[AVMusicTrackEventIterator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/initWithImpl:
func NewMusicTrackEventIteratorWithImpl(impl unsafe.Pointer) AVMusicTrackEventIterator {
	instance := getAVMusicTrackEventIteratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVMusicTrackEventIteratorFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/deleteEvent
func (m AVMusicTrackEventIterator) DeleteEvent() {
	objc.Send[objc.ID](m.ID, objc.Sel("deleteEvent"))
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/getEventInfo:outEventType:eventData:dataSize:
func (m AVMusicTrackEventIterator) GetEventInfoOutEventTypeEventDataDataSize(info []float64, type_ unsafe.Pointer, data unsafe.Pointer, size unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("getEventInfo:outEventType:eventData:dataSize:"), info, type_, data, size)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/hasCurrentEvent
func (m AVMusicTrackEventIterator) HasCurrentEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasCurrentEvent"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/hasNextEvent
func (m AVMusicTrackEventIterator) HasNextEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasNextEvent"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/hasPreviousEvent
func (m AVMusicTrackEventIterator) HasPreviousEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasPreviousEvent"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/nextEvent
func (m AVMusicTrackEventIterator) NextEvent() int {
	rv := objc.Send[int](m.ID, objc.Sel("nextEvent"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/previousEvent
func (m AVMusicTrackEventIterator) PreviousEvent() int {
	rv := objc.Send[int](m.ID, objc.Sel("previousEvent"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/seek:
func (m AVMusicTrackEventIterator) Seek(seek float64) {
	objc.Send[objc.ID](m.ID, objc.Sel("seek:"), seek)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/setEventInfo:data:
func (m AVMusicTrackEventIterator) SetEventInfoData(info uint32, data unsafe.Pointer) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("setEventInfo:data:"), info, data)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/setEventTime:
func (m AVMusicTrackEventIterator) SetEventTime(time float64) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("setEventTime:"), time)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicTrackEventIterator/initWithImpl:
func (m AVMusicTrackEventIterator) InitWithImpl(impl unsafe.Pointer) AVMusicTrackEventIterator {
	rv := objc.Send[AVMusicTrackEventIterator](m.ID, objc.Sel("initWithImpl:"), impl)
	return rv
}

