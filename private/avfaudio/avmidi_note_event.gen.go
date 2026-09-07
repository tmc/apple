// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
	rv := objc.SendIfResponds[AVMIDINoteEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

type AVMIDINoteEvent struct {
	objectivec.Object
}

// AVMIDINoteEventFromID constructs a [AVMIDINoteEvent] from an objc.ID.
func AVMIDINoteEventFromID(id objc.ID) AVMIDINoteEvent {
	return AVMIDINoteEvent{objectivec.Object{ID: id}}
}

// NOTE: AVMIDINoteEvent embeds objectivec.Object because the parent type is
// unavailable, but IAVMIDINoteEvent embeds IAVMusicEvent, which that fallback
// cannot satisfy; skip compile-time assertion.

// An interface definition for the [AVMIDINoteEvent] class.
type IAVMIDINoteEvent interface {
	IAVMusicEvent
}

// Init initializes the instance.
func (a AVMIDINoteEvent) Init() AVMIDINoteEvent {
	rv := objc.SendIfResponds[AVMIDINoteEvent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVMIDINoteEvent) Autorelease() AVMIDINoteEvent {
	rv := objc.SendIfResponds[AVMIDINoteEvent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDINoteEvent creates a new AVMIDINoteEvent instance.
func NewAVMIDINoteEvent() AVMIDINoteEvent {
	class := getAVMIDINoteEventClass()
	rv := objc.SendIfResponds[AVMIDINoteEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
