// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioMixingDestination] class.
var (
	_AVAudioMixingDestinationClass     AVAudioMixingDestinationClass
	_AVAudioMixingDestinationClassOnce sync.Once
)

func getAVAudioMixingDestinationClass() AVAudioMixingDestinationClass {
	_AVAudioMixingDestinationClassOnce.Do(func() {
		_AVAudioMixingDestinationClass = AVAudioMixingDestinationClass{class: objc.GetClass("AVAudioMixingDestination")}
	})
	return _AVAudioMixingDestinationClass
}

// GetAVAudioMixingDestinationClass returns the class object for AVAudioMixingDestination.
func GetAVAudioMixingDestinationClass() AVAudioMixingDestinationClass {
	return getAVAudioMixingDestinationClass()
}

type AVAudioMixingDestinationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioMixingDestinationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioMixingDestinationClass) Alloc() AVAudioMixingDestination {
	rv := objc.Send[AVAudioMixingDestination](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioMixingDestination.DebugDescription]
//   - [AVAudioMixingDestination.Description]
//   - [AVAudioMixingDestination.Hash]
//   - [AVAudioMixingDestination.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination
type AVAudioMixingDestination struct {
	objectivec.Object
}

// AVAudioMixingDestinationFromID constructs a [AVAudioMixingDestination] from an objc.ID.
func AVAudioMixingDestinationFromID(id objc.ID) AVAudioMixingDestination {
	return AVAudioMixingDestination{objectivec.Object{ID: id}}
}

// Ensure AVAudioMixingDestination implements IAVAudioMixingDestination.
var _ IAVAudioMixingDestination = AVAudioMixingDestination{}

// An interface definition for the [AVAudioMixingDestination] class.
//
// # Methods
//
//   - [IAVAudioMixingDestination.DebugDescription]
//   - [IAVAudioMixingDestination.Description]
//   - [IAVAudioMixingDestination.Hash]
//   - [IAVAudioMixingDestination.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination
type IAVAudioMixingDestination interface {
	objectivec.IObject

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a AVAudioMixingDestination) Init() AVAudioMixingDestination {
	rv := objc.Send[AVAudioMixingDestination](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioMixingDestination) Autorelease() AVAudioMixingDestination {
	rv := objc.Send[AVAudioMixingDestination](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioMixingDestination creates a new AVAudioMixingDestination instance.
func NewAVAudioMixingDestination() AVAudioMixingDestination {
	class := getAVAudioMixingDestinationClass()
	rv := objc.Send[AVAudioMixingDestination](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination/debugDescription
func (a AVAudioMixingDestination) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination/description
func (a AVAudioMixingDestination) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination/hash
func (a AVAudioMixingDestination) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioMixingDestination/superclass
func (a AVAudioMixingDestination) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}
