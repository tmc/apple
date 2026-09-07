// Code generated from Apple documentation for avfaudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioEngineLock] class.
var (
	_AVAudioEngineLockClass     AVAudioEngineLockClass
	_AVAudioEngineLockClassOnce sync.Once
)

func getAVAudioEngineLockClass() AVAudioEngineLockClass {
	_AVAudioEngineLockClassOnce.Do(func() {
		_AVAudioEngineLockClass = AVAudioEngineLockClass{class: objc.GetClass("AVAudioEngineLock")}
	})
	return _AVAudioEngineLockClass
}

// GetAVAudioEngineLockClass returns the class object for AVAudioEngineLock.
func GetAVAudioEngineLockClass() AVAudioEngineLockClass {
	return getAVAudioEngineLockClass()
}

type AVAudioEngineLockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioEngineLockClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioEngineLockClass) Alloc() AVAudioEngineLock {
	rv := objc.SendIfResponds[AVAudioEngineLock](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioEngineLock.Unlock]
//   - [AVAudioEngineLock.InitWithLock]
type AVAudioEngineLock struct {
	objectivec.Object
}

// AVAudioEngineLockFromID constructs a [AVAudioEngineLock] from an objc.ID.
func AVAudioEngineLockFromID(id objc.ID) AVAudioEngineLock {
	return AVAudioEngineLock{objectivec.Object{ID: id}}
}

// Ensure AVAudioEngineLock implements IAVAudioEngineLock.
var _ IAVAudioEngineLock = AVAudioEngineLock{}

// An interface definition for the [AVAudioEngineLock] class.
//
// # Methods
//
//   - [IAVAudioEngineLock.Unlock]
//   - [IAVAudioEngineLock.InitWithLock]
type IAVAudioEngineLock interface {
	objectivec.IObject

	// Topic: Methods

	Unlock()
	InitWithLock(lock unsafe.Pointer) AVAudioEngineLock
}

// Init initializes the instance.
func (a AVAudioEngineLock) Init() AVAudioEngineLock {
	rv := objc.SendIfResponds[AVAudioEngineLock](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioEngineLock) Autorelease() AVAudioEngineLock {
	rv := objc.SendIfResponds[AVAudioEngineLock](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioEngineLock creates a new AVAudioEngineLock instance.
func NewAVAudioEngineLock() AVAudioEngineLock {
	class := getAVAudioEngineLockClass()
	rv := objc.SendIfResponds[AVAudioEngineLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAVAudioEngineLockWithLock(lock unsafe.Pointer) AVAudioEngineLock {
	instance := getAVAudioEngineLockClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLock:"), lock)
	return AVAudioEngineLockFromID(rv)
}

func (a AVAudioEngineLock) Unlock() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("unlock"))
}
func (a AVAudioEngineLock) InitWithLock(lock unsafe.Pointer) AVAudioEngineLock {
	rv := objc.SendIfResponds[AVAudioEngineLock](a.ID, objc.Sel("initWithLock:"), lock)
	return rv
}
