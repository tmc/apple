// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[AVMusicTrackEventIterator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

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
type IAVMusicTrackEventIterator interface {
	objectivec.IObject

	// Topic: Methods

	DeleteEvent()
	GetEventInfoOutEventTypeEventDataDataSize(info []float64, type_ *uint32, data unsafe.Pointer, size *uint32)
	HasCurrentEvent() bool
	HasNextEvent() bool
	HasPreviousEvent() bool
	NextEvent() int
	PreviousEvent() int
	Seek(seek float64)
	SetEventInfoData(info uint32, data unsafe.Pointer) bool
	SetEventTime(time float64) bool
	InitWithImpl(impl *MusicTrackEventIteratorImpl) AVMusicTrackEventIterator
}

// Init initializes the instance.
func (a AVMusicTrackEventIterator) Init() AVMusicTrackEventIterator {
	rv := objc.SendIfResponds[AVMusicTrackEventIterator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVMusicTrackEventIterator) Autorelease() AVMusicTrackEventIterator {
	rv := objc.SendIfResponds[AVMusicTrackEventIterator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMusicTrackEventIterator creates a new AVMusicTrackEventIterator instance.
func NewAVMusicTrackEventIterator() AVMusicTrackEventIterator {
	class := getAVMusicTrackEventIteratorClass()
	rv := objc.SendIfResponds[AVMusicTrackEventIterator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMusicTrackEventIteratorWithImpl(impl *MusicTrackEventIteratorImpl) AVMusicTrackEventIterator {
	instance := getAVMusicTrackEventIteratorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithImpl:"), unsafe.Pointer(impl))
	return AVMusicTrackEventIteratorFromID(rv)
}

func (a AVMusicTrackEventIterator) DeleteEvent() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("deleteEvent"))
}
func (a AVMusicTrackEventIterator) GetEventInfoOutEventTypeEventDataDataSize(info []float64, type_ *uint32, data unsafe.Pointer, size *uint32) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("getEventInfo:outEventType:eventData:dataSize:"), info, unsafe.Pointer(type_), data, unsafe.Pointer(size))
}
func (a AVMusicTrackEventIterator) HasCurrentEvent() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("hasCurrentEvent"))
	return rv
}
func (a AVMusicTrackEventIterator) HasNextEvent() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("hasNextEvent"))
	return rv
}
func (a AVMusicTrackEventIterator) HasPreviousEvent() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("hasPreviousEvent"))
	return rv
}
func (a AVMusicTrackEventIterator) NextEvent() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("nextEvent"))
	return rv
}
func (a AVMusicTrackEventIterator) PreviousEvent() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("previousEvent"))
	return rv
}
func (a AVMusicTrackEventIterator) Seek(seek float64) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("seek:"), seek)
}
func (a AVMusicTrackEventIterator) SetEventInfoData(info uint32, data unsafe.Pointer) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("setEventInfo:data:"), info, data)
	return rv
}
func (a AVMusicTrackEventIterator) SetEventTime(time float64) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("setEventTime:"), time)
	return rv
}
func (a AVMusicTrackEventIterator) InitWithImpl(impl *MusicTrackEventIteratorImpl) AVMusicTrackEventIterator {
	rv := objc.SendIfResponds[AVMusicTrackEventIterator](a.ID, objc.Sel("initWithImpl:"), unsafe.Pointer(impl))
	return rv
}
