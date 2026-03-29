// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioClock] class.
var (
	_AVAudioClockClass     AVAudioClockClass
	_AVAudioClockClassOnce sync.Once
)

func getAVAudioClockClass() AVAudioClockClass {
	_AVAudioClockClassOnce.Do(func() {
		_AVAudioClockClass = AVAudioClockClass{class: objc.GetClass("AVAudioClock")}
	})
	return _AVAudioClockClass
}

// GetAVAudioClockClass returns the class object for AVAudioClock.
func GetAVAudioClockClass() AVAudioClockClass {
	return getAVAudioClockClass()
}

type AVAudioClockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioClockClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioClockClass) Alloc() AVAudioClock {
	rv := objc.Send[AVAudioClock](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioClock.AwaitIOCycle]
//   - [AVAudioClock.CurrentAudioTimeStamp]
//   - [AVAudioClock.CurrentIONumberFrames]
//   - [AVAudioClock.CurrentTime]
//   - [AVAudioClock.InitWithNode]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock
type AVAudioClock struct {
	objectivec.Object
}

// AVAudioClockFromID constructs a [AVAudioClock] from an objc.ID.
func AVAudioClockFromID(id objc.ID) AVAudioClock {
	return AVAudioClock{objectivec.Object{ID: id}}
}
// Ensure AVAudioClock implements IAVAudioClock.
var _ IAVAudioClock = AVAudioClock{}

// An interface definition for the [AVAudioClock] class.
//
// # Methods
//
//   - [IAVAudioClock.AwaitIOCycle]
//   - [IAVAudioClock.CurrentAudioTimeStamp]
//   - [IAVAudioClock.CurrentIONumberFrames]
//   - [IAVAudioClock.CurrentTime]
//   - [IAVAudioClock.InitWithNode]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock
type IAVAudioClock interface {
	objectivec.IObject

	// Topic: Methods

	AwaitIOCycle(iOCycle unsafe.Pointer) objectivec.IObject
	CurrentAudioTimeStamp() objectivec.IObject
	CurrentIONumberFrames() int64
	CurrentTime() IAVAudioTime
	InitWithNode(node unsafe.Pointer) AVAudioClock
}

// Init initializes the instance.
func (a AVAudioClock) Init() AVAudioClock {
	rv := objc.Send[AVAudioClock](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioClock) Autorelease() AVAudioClock {
	rv := objc.Send[AVAudioClock](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioClock creates a new AVAudioClock instance.
func NewAVAudioClock() AVAudioClock {
	class := getAVAudioClockClass()
	rv := objc.Send[AVAudioClock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock/initWithNode:
func NewAudioClockWithNode(node unsafe.Pointer) AVAudioClock {
	instance := getAVAudioClockClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNode:"), node)
	return AVAudioClockFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock/awaitIOCycle:
func (a AVAudioClock) AwaitIOCycle(iOCycle unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("awaitIOCycle:"), iOCycle)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock/currentAudioTimeStamp
func (a AVAudioClock) CurrentAudioTimeStamp() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentAudioTimeStamp"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock/currentIONumberFrames
func (a AVAudioClock) CurrentIONumberFrames() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("currentIONumberFrames"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock/initWithNode:
func (a AVAudioClock) InitWithNode(node unsafe.Pointer) AVAudioClock {
	rv := objc.Send[AVAudioClock](a.ID, objc.Sel("initWithNode:"), node)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioClock/currentTime
func (a AVAudioClock) CurrentTime() IAVAudioTime {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentTime"))
	return AVAudioTimeFromID(objc.ID(rv))
}

