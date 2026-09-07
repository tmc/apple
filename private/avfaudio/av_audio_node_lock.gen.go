// Code generated from Apple documentation for avfaudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioNodeLock] class.
var (
	_AVAudioNodeLockClass     AVAudioNodeLockClass
	_AVAudioNodeLockClassOnce sync.Once
)

func getAVAudioNodeLockClass() AVAudioNodeLockClass {
	_AVAudioNodeLockClassOnce.Do(func() {
		_AVAudioNodeLockClass = AVAudioNodeLockClass{class: objc.GetClass("AVAudioNodeLock")}
	})
	return _AVAudioNodeLockClass
}

// GetAVAudioNodeLockClass returns the class object for AVAudioNodeLock.
func GetAVAudioNodeLockClass() AVAudioNodeLockClass {
	return getAVAudioNodeLockClass()
}

type AVAudioNodeLockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioNodeLockClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioNodeLockClass) Alloc() AVAudioNodeLock {
	rv := objc.SendIfResponds[AVAudioNodeLock](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioNodeLock.Unlock]
//   - [AVAudioNodeLock.InitWithLock]
type AVAudioNodeLock struct {
	objectivec.Object
}

// AVAudioNodeLockFromID constructs a [AVAudioNodeLock] from an objc.ID.
func AVAudioNodeLockFromID(id objc.ID) AVAudioNodeLock {
	return AVAudioNodeLock{objectivec.Object{ID: id}}
}

// Ensure AVAudioNodeLock implements IAVAudioNodeLock.
var _ IAVAudioNodeLock = AVAudioNodeLock{}

// An interface definition for the [AVAudioNodeLock] class.
//
// # Methods
//
//   - [IAVAudioNodeLock.Unlock]
//   - [IAVAudioNodeLock.InitWithLock]
type IAVAudioNodeLock interface {
	objectivec.IObject

	// Topic: Methods

	Unlock()
	InitWithLock(lock unsafe.Pointer) AVAudioNodeLock
}

// Init initializes the instance.
func (a AVAudioNodeLock) Init() AVAudioNodeLock {
	rv := objc.SendIfResponds[AVAudioNodeLock](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioNodeLock) Autorelease() AVAudioNodeLock {
	rv := objc.SendIfResponds[AVAudioNodeLock](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioNodeLock creates a new AVAudioNodeLock instance.
func NewAVAudioNodeLock() AVAudioNodeLock {
	class := getAVAudioNodeLockClass()
	rv := objc.SendIfResponds[AVAudioNodeLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAVAudioNodeLockWithLock(lock unsafe.Pointer) AVAudioNodeLock {
	instance := getAVAudioNodeLockClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLock:"), lock)
	return AVAudioNodeLockFromID(rv)
}

func (a AVAudioNodeLock) Unlock() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("unlock"))
}
func (a AVAudioNodeLock) InitWithLock(lock unsafe.Pointer) AVAudioNodeLock {
	rv := objc.SendIfResponds[AVAudioNodeLock](a.ID, objc.Sel("initWithLock:"), lock)
	return rv
}
