// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMIDIPlayer] class.
var (
	_AVMIDIPlayerClass     AVMIDIPlayerClass
	_AVMIDIPlayerClassOnce sync.Once
)

func getAVMIDIPlayerClass() AVMIDIPlayerClass {
	_AVMIDIPlayerClassOnce.Do(func() {
		_AVMIDIPlayerClass = AVMIDIPlayerClass{class: objc.GetClass("AVMIDIPlayer")}
	})
	return _AVMIDIPlayerClass
}

// GetAVMIDIPlayerClass returns the class object for AVMIDIPlayer.
func GetAVMIDIPlayerClass() AVMIDIPlayerClass {
	return getAVMIDIPlayerClass()
}

type AVMIDIPlayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMIDIPlayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMIDIPlayerClass) Alloc() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVMIDIPlayer.BeatsForHostTime]
//   - [AVMIDIPlayer.DestroyBase]
//   - [AVMIDIPlayer.HostTimeForBeats]
//   - [AVMIDIPlayer.InitBase]
//   - [AVMIDIPlayer.Playing]
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer
type AVMIDIPlayer struct {
	objectivec.Object
}

// AVMIDIPlayerFromID constructs a [AVMIDIPlayer] from an objc.ID.
func AVMIDIPlayerFromID(id objc.ID) AVMIDIPlayer {
	return AVMIDIPlayer{objectivec.Object{ID: id}}
}
// Ensure AVMIDIPlayer implements IAVMIDIPlayer.
var _ IAVMIDIPlayer = AVMIDIPlayer{}

// An interface definition for the [AVMIDIPlayer] class.
//
// # Methods
//
//   - [IAVMIDIPlayer.BeatsForHostTime]
//   - [IAVMIDIPlayer.DestroyBase]
//   - [IAVMIDIPlayer.HostTimeForBeats]
//   - [IAVMIDIPlayer.InitBase]
//   - [IAVMIDIPlayer.Playing]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer
type IAVMIDIPlayer interface {
	objectivec.IObject

	// Topic: Methods

	BeatsForHostTime(time uint64) float64
	DestroyBase()
	HostTimeForBeats(beats float64) uint64
	InitBase() AVMIDIPlayer
	Playing() bool
}

// Init initializes the instance.
func (m AVMIDIPlayer) Init() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMIDIPlayer) Autorelease() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIPlayer creates a new AVMIDIPlayer instance.
func NewAVMIDIPlayer() AVMIDIPlayer {
	class := getAVMIDIPlayerClass()
	rv := objc.Send[AVMIDIPlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/initBase
func NewMIDIPlayerBase() AVMIDIPlayer {
	instance := getAVMIDIPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initBase"))
	return AVMIDIPlayerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/beatsForHostTime:
func (m AVMIDIPlayer) BeatsForHostTime(time uint64) float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("beatsForHostTime:"), time)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/destroyBase
func (m AVMIDIPlayer) DestroyBase() {
	objc.Send[objc.ID](m.ID, objc.Sel("destroyBase"))
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/hostTimeForBeats:
func (m AVMIDIPlayer) HostTimeForBeats(beats float64) uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hostTimeForBeats:"), beats)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/initBase
func (m AVMIDIPlayer) InitBase() AVMIDIPlayer {
	rv := objc.Send[AVMIDIPlayer](m.ID, objc.Sel("initBase"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVMIDIPlayer/playing
func (m AVMIDIPlayer) Playing() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("playing"))
	return rv
}

