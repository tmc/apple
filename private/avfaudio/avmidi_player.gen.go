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
	rv := objc.SendIfResponds[AVMIDIPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVMIDIPlayer.BeatsForHostTime]
//   - [AVMIDIPlayer.DestroyBase]
//   - [AVMIDIPlayer.HostTimeForBeats]
//   - [AVMIDIPlayer.InitBase]
//   - [AVMIDIPlayer.Playing]
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
func (a AVMIDIPlayer) Init() AVMIDIPlayer {
	rv := objc.SendIfResponds[AVMIDIPlayer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVMIDIPlayer) Autorelease() AVMIDIPlayer {
	rv := objc.SendIfResponds[AVMIDIPlayer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMIDIPlayer creates a new AVMIDIPlayer instance.
func NewAVMIDIPlayer() AVMIDIPlayer {
	class := getAVMIDIPlayerClass()
	rv := objc.SendIfResponds[AVMIDIPlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMIDIPlayerBase() AVMIDIPlayer {
	instance := getAVMIDIPlayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initBase"))
	return AVMIDIPlayerFromID(rv)
}

func (a AVMIDIPlayer) BeatsForHostTime(time uint64) float64 {
	rv := objc.SendIfResponds[float64](a.ID, objc.Sel("beatsForHostTime:"), time)
	return rv
}
func (a AVMIDIPlayer) DestroyBase() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("destroyBase"))
}
func (a AVMIDIPlayer) HostTimeForBeats(beats float64) uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("hostTimeForBeats:"), beats)
	return rv
}
func (a AVMIDIPlayer) InitBase() AVMIDIPlayer {
	rv := objc.SendIfResponds[AVMIDIPlayer](a.ID, objc.Sel("initBase"))
	return rv
}

func (a AVMIDIPlayer) Playing() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("playing"))
	return rv
}
