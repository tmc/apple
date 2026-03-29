// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioSequencer] class.
var (
	_AVAudioSequencerClass     AVAudioSequencerClass
	_AVAudioSequencerClassOnce sync.Once
)

func getAVAudioSequencerClass() AVAudioSequencerClass {
	_AVAudioSequencerClassOnce.Do(func() {
		_AVAudioSequencerClass = AVAudioSequencerClass{class: objc.GetClass("AVAudioSequencer")}
	})
	return _AVAudioSequencerClass
}

// GetAVAudioSequencerClass returns the class object for AVAudioSequencer.
func GetAVAudioSequencerClass() AVAudioSequencerClass {
	return getAVAudioSequencerClass()
}

type AVAudioSequencerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSequencerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSequencerClass) Alloc() AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioSequencer.CleanTracks]
//   - [AVAudioSequencer.GetTempoTrack]
//   - [AVAudioSequencer.NumberOfTracks]
//   - [AVAudioSequencer.SetTempoTrack]
//   - [AVAudioSequencer.SetTrackArray]
//   - [AVAudioSequencer.SetupTrackArray]
//   - [AVAudioSequencer.SetupTracks]
//   - [AVAudioSequencer.TrackArray]
//   - [AVAudioSequencer.InitWithImpl]
//   - [AVAudioSequencer.Playing]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer
type AVAudioSequencer struct {
	objectivec.Object
}

// AVAudioSequencerFromID constructs a [AVAudioSequencer] from an objc.ID.
func AVAudioSequencerFromID(id objc.ID) AVAudioSequencer {
	return AVAudioSequencer{objectivec.Object{ID: id}}
}
// Ensure AVAudioSequencer implements IAVAudioSequencer.
var _ IAVAudioSequencer = AVAudioSequencer{}

// An interface definition for the [AVAudioSequencer] class.
//
// # Methods
//
//   - [IAVAudioSequencer.CleanTracks]
//   - [IAVAudioSequencer.GetTempoTrack]
//   - [IAVAudioSequencer.NumberOfTracks]
//   - [IAVAudioSequencer.SetTempoTrack]
//   - [IAVAudioSequencer.SetTrackArray]
//   - [IAVAudioSequencer.SetupTrackArray]
//   - [IAVAudioSequencer.SetupTracks]
//   - [IAVAudioSequencer.TrackArray]
//   - [IAVAudioSequencer.InitWithImpl]
//   - [IAVAudioSequencer.Playing]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer
type IAVAudioSequencer interface {
	objectivec.IObject

	// Topic: Methods

	CleanTracks()
	GetTempoTrack() objectivec.IObject
	NumberOfTracks() uint64
	SetTempoTrack(track objectivec.IObject)
	SetTrackArray(array objectivec.IObject)
	SetupTrackArray()
	SetupTracks()
	TrackArray() objectivec.IObject
	InitWithImpl(impl unsafe.Pointer) AVAudioSequencer
	Playing() bool
}

// Init initializes the instance.
func (a AVAudioSequencer) Init() AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSequencer) Autorelease() AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSequencer creates a new AVAudioSequencer instance.
func NewAVAudioSequencer() AVAudioSequencer {
	class := getAVAudioSequencerClass()
	rv := objc.Send[AVAudioSequencer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/initWithImpl:
func NewAudioSequencerWithImpl(impl unsafe.Pointer) AVAudioSequencer {
	instance := getAVAudioSequencerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioSequencerFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/cleanTracks
func (a AVAudioSequencer) CleanTracks() {
	objc.Send[objc.ID](a.ID, objc.Sel("cleanTracks"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/getTempoTrack
func (a AVAudioSequencer) GetTempoTrack() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getTempoTrack"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/numberOfTracks
func (a AVAudioSequencer) NumberOfTracks() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("numberOfTracks"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/setTempoTrack:
func (a AVAudioSequencer) SetTempoTrack(track objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setTempoTrack:"), track)
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/setTrackArray:
func (a AVAudioSequencer) SetTrackArray(array objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setTrackArray:"), array)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/setupTrackArray
func (a AVAudioSequencer) SetupTrackArray() {
	objc.Send[objc.ID](a.ID, objc.Sel("setupTrackArray"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/setupTracks
func (a AVAudioSequencer) SetupTracks() {
	objc.Send[objc.ID](a.ID, objc.Sel("setupTracks"))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/trackArray
func (a AVAudioSequencer) TrackArray() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("trackArray"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/initWithImpl:
func (a AVAudioSequencer) InitWithImpl(impl unsafe.Pointer) AVAudioSequencer {
	rv := objc.Send[AVAudioSequencer](a.ID, objc.Sel("initWithImpl:"), impl)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSequencer/playing
func (a AVAudioSequencer) Playing() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("playing"))
	return rv
}

