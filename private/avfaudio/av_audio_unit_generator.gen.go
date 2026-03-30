// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitGenerator] class.
var (
	_AVAudioUnitGeneratorClass     AVAudioUnitGeneratorClass
	_AVAudioUnitGeneratorClassOnce sync.Once
)

func getAVAudioUnitGeneratorClass() AVAudioUnitGeneratorClass {
	_AVAudioUnitGeneratorClassOnce.Do(func() {
		_AVAudioUnitGeneratorClass = AVAudioUnitGeneratorClass{class: objc.GetClass("AVAudioUnitGenerator")}
	})
	return _AVAudioUnitGeneratorClass
}

// GetAVAudioUnitGeneratorClass returns the class object for AVAudioUnitGenerator.
func GetAVAudioUnitGeneratorClass() AVAudioUnitGeneratorClass {
	return getAVAudioUnitGeneratorClass()
}

type AVAudioUnitGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitGeneratorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitGeneratorClass) Alloc() AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitGenerator.DebugDescription]
//   - [AVAudioUnitGenerator.Description]
//   - [AVAudioUnitGenerator.Hash]
//   - [AVAudioUnitGenerator.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator
type AVAudioUnitGenerator struct {
	AVAudioUnit
}

// AVAudioUnitGeneratorFromID constructs a [AVAudioUnitGenerator] from an objc.ID.
func AVAudioUnitGeneratorFromID(id objc.ID) AVAudioUnitGenerator {
	return AVAudioUnitGenerator{AVAudioUnit: AVAudioUnitFromID(id)}
}

// Ensure AVAudioUnitGenerator implements IAVAudioUnitGenerator.
var _ IAVAudioUnitGenerator = AVAudioUnitGenerator{}

// An interface definition for the [AVAudioUnitGenerator] class.
//
// # Methods
//
//   - [IAVAudioUnitGenerator.DebugDescription]
//   - [IAVAudioUnitGenerator.Description]
//   - [IAVAudioUnitGenerator.Hash]
//   - [IAVAudioUnitGenerator.Superclass]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator
type IAVAudioUnitGenerator interface {
	IAVAudioUnit

	// Topic: Methods

	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (a AVAudioUnitGenerator) Init() AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitGenerator) Autorelease() AVAudioUnitGenerator {
	rv := objc.Send[AVAudioUnitGenerator](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitGenerator creates a new AVAudioUnitGenerator instance.
func NewAVAudioUnitGenerator() AVAudioUnitGenerator {
	class := getAVAudioUnitGeneratorClass()
	rv := objc.Send[AVAudioUnitGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioUnitGeneratorWithImpl(impl unsafe.Pointer) AVAudioUnitGenerator {
	instance := getAVAudioUnitGeneratorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioUnitGeneratorFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/debugDescription
func (a AVAudioUnitGenerator) DebugDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/description
func (a AVAudioUnitGenerator) Description() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/hash
func (a AVAudioUnitGenerator) Hash() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitGenerator/superclass
func (a AVAudioUnitGenerator) Superclass() objc.Class {
	rv := objc.Send[objc.Class](a.ID, objc.Sel("superclass"))
	return rv
}
