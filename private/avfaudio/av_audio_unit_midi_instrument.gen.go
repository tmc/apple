// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitMIDIInstrument] class.
var (
	_AVAudioUnitMIDIInstrumentClass     AVAudioUnitMIDIInstrumentClass
	_AVAudioUnitMIDIInstrumentClassOnce sync.Once
)

func getAVAudioUnitMIDIInstrumentClass() AVAudioUnitMIDIInstrumentClass {
	_AVAudioUnitMIDIInstrumentClassOnce.Do(func() {
		_AVAudioUnitMIDIInstrumentClass = AVAudioUnitMIDIInstrumentClass{class: objc.GetClass("AVAudioUnitMIDIInstrument")}
	})
	return _AVAudioUnitMIDIInstrumentClass
}

// GetAVAudioUnitMIDIInstrumentClass returns the class object for AVAudioUnitMIDIInstrument.
func GetAVAudioUnitMIDIInstrumentClass() AVAudioUnitMIDIInstrumentClass {
	return getAVAudioUnitMIDIInstrumentClass()
}

type AVAudioUnitMIDIInstrumentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitMIDIInstrumentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitMIDIInstrumentClass) Alloc() AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitMIDIInstrument.DebugDescription]
//   - [AVAudioUnitMIDIInstrument.Description]
//   - [AVAudioUnitMIDIInstrument.Hash]
//   - [AVAudioUnitMIDIInstrument.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument
type AVAudioUnitMIDIInstrument struct {
	AVAudioUnit
}

// AVAudioUnitMIDIInstrumentFromID constructs a [AVAudioUnitMIDIInstrument] from an objc.ID.
func AVAudioUnitMIDIInstrumentFromID(id objc.ID) AVAudioUnitMIDIInstrument {
	return AVAudioUnitMIDIInstrument{AVAudioUnit: AVAudioUnitFromID(id)}
}

// Ensure AVAudioUnitMIDIInstrument implements IAVAudioUnitMIDIInstrument.
var _ IAVAudioUnitMIDIInstrument = AVAudioUnitMIDIInstrument{}

// An interface definition for the [AVAudioUnitMIDIInstrument] class.
//
// # Methods
//
//   - [IAVAudioUnitMIDIInstrument.DebugDescription]
//   - [IAVAudioUnitMIDIInstrument.Description]
//   - [IAVAudioUnitMIDIInstrument.Hash]
//   - [IAVAudioUnitMIDIInstrument.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument
type IAVAudioUnitMIDIInstrument interface {
	IAVAudioUnit

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a AVAudioUnitMIDIInstrument) Init() AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitMIDIInstrument) Autorelease() AVAudioUnitMIDIInstrument {
	rv := objc.Send[AVAudioUnitMIDIInstrument](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitMIDIInstrument creates a new AVAudioUnitMIDIInstrument instance.
func NewAVAudioUnitMIDIInstrument() AVAudioUnitMIDIInstrument {
	class := getAVAudioUnitMIDIInstrumentClass()
	rv := objc.Send[AVAudioUnitMIDIInstrument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioUnitMIDIInstrumentWithImpl(impl unsafe.Pointer) AVAudioUnitMIDIInstrument {
	instance := getAVAudioUnitMIDIInstrumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioUnitMIDIInstrumentFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/debugDescription
func (a AVAudioUnitMIDIInstrument) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/description
func (a AVAudioUnitMIDIInstrument) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/hash
func (a AVAudioUnitMIDIInstrument) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitMIDIInstrument/superclass
func (a AVAudioUnitMIDIInstrument) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}
