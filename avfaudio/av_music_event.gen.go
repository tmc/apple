// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMusicEvent] class.
var (
	_AVMusicEventClass     AVMusicEventClass
	_AVMusicEventClassOnce sync.Once
)

func getAVMusicEventClass() AVMusicEventClass {
	_AVMusicEventClassOnce.Do(func() {
		_AVMusicEventClass = AVMusicEventClass{class: objc.GetClass("AVMusicEvent")}
	})
	return _AVMusicEventClass
}

// GetAVMusicEventClass returns the class object for AVMusicEvent.
func GetAVMusicEventClass() AVMusicEventClass {
	return getAVMusicEventClass()
}

type AVMusicEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMusicEventClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMusicEventClass) Alloc() AVMusicEvent {
	rv := objc.Send[AVMusicEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A base class for the events you associate with a music track.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicEvent
type AVMusicEvent struct {
	objectivec.Object
}

// AVMusicEventFromID constructs a [AVMusicEvent] from an objc.ID.
//
// A base class for the events you associate with a music track.
func AVMusicEventFromID(id objc.ID) AVMusicEvent {
	return AVMusicEvent{objectivec.Object{ID: id}}
}
// NOTE: AVMusicEvent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMusicEvent] class.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMusicEvent
type IAVMusicEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m AVMusicEvent) Init() AVMusicEvent {
	rv := objc.Send[AVMusicEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMusicEvent) Autorelease() AVMusicEvent {
	rv := objc.Send[AVMusicEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMusicEvent creates a new AVMusicEvent instance.
func NewAVMusicEvent() AVMusicEvent {
	class := getAVMusicEventClass()
	rv := objc.Send[AVMusicEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

