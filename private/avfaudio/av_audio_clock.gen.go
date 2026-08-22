// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
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
	rv := objc.SendIfResponds[AVAudioClock](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioClock.AwaitIOCycle]
//   - [AVAudioClock.CurrentAudioTimeStamp]
//   - [AVAudioClock.CurrentIONumberFrames]
//   - [AVAudioClock.CurrentTime]
//   - [AVAudioClock.InitWithNode]
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
type IAVAudioClock interface {
	objectivec.IObject

	// Topic: Methods

	AwaitIOCycle(iOCycle *uint32) objectivec.IObject
	CurrentAudioTimeStamp() coreaudiotypes.AudioTimeStamp
	CurrentIONumberFrames() int64
	CurrentTime() IAVAudioTime
	InitWithNode(node unsafe.Pointer) AVAudioClock
}

// Init initializes the instance.
func (a AVAudioClock) Init() AVAudioClock {
	rv := objc.SendIfResponds[AVAudioClock](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioClock) Autorelease() AVAudioClock {
	rv := objc.SendIfResponds[AVAudioClock](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioClock creates a new AVAudioClock instance.
func NewAVAudioClock() AVAudioClock {
	class := getAVAudioClockClass()
	rv := objc.SendIfResponds[AVAudioClock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAudioClockWithNode(node unsafe.Pointer) AVAudioClock {
	instance := getAVAudioClockClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithNode:"), node)
	return AVAudioClockFromID(rv)
}

func (a AVAudioClock) AwaitIOCycle(iOCycle *uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("awaitIOCycle:"), unsafe.Pointer(iOCycle))
	return objectivec.Object{ID: rv}
}
func (a AVAudioClock) CurrentAudioTimeStamp() coreaudiotypes.AudioTimeStamp {
	rv := objc.SendIfResponds[coreaudiotypes.AudioTimeStamp](a.ID, objc.Sel("currentAudioTimeStamp"))
	return coreaudiotypes.AudioTimeStamp(rv)
}
func (a AVAudioClock) CurrentIONumberFrames() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("currentIONumberFrames"))
	return rv
}
func (a AVAudioClock) InitWithNode(node unsafe.Pointer) AVAudioClock {
	rv := objc.SendIfResponds[AVAudioClock](a.ID, objc.Sel("initWithNode:"), node)
	return rv
}

func (a AVAudioClock) CurrentTime() IAVAudioTime {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("currentTime"))
	return AVAudioTimeFromID(objc.ID(rv))
}
