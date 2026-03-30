// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnitSampler] class.
var (
	_AVAudioUnitSamplerClass     AVAudioUnitSamplerClass
	_AVAudioUnitSamplerClassOnce sync.Once
)

func getAVAudioUnitSamplerClass() AVAudioUnitSamplerClass {
	_AVAudioUnitSamplerClassOnce.Do(func() {
		_AVAudioUnitSamplerClass = AVAudioUnitSamplerClass{class: objc.GetClass("AVAudioUnitSampler")}
	})
	return _AVAudioUnitSamplerClass
}

// GetAVAudioUnitSamplerClass returns the class object for AVAudioUnitSampler.
func GetAVAudioUnitSamplerClass() AVAudioUnitSamplerClass {
	return getAVAudioUnitSamplerClass()
}

type AVAudioUnitSamplerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitSamplerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitSamplerClass) Alloc() AVAudioUnitSampler {
	rv := objc.Send[AVAudioUnitSampler](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnitSampler.MasterGain]
//   - [AVAudioUnitSampler.SetMasterGain]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler
type AVAudioUnitSampler struct {
	AVAudioUnitMIDIInstrument
}

// AVAudioUnitSamplerFromID constructs a [AVAudioUnitSampler] from an objc.ID.
func AVAudioUnitSamplerFromID(id objc.ID) AVAudioUnitSampler {
	return AVAudioUnitSampler{AVAudioUnitMIDIInstrument: AVAudioUnitMIDIInstrumentFromID(id)}
}

// Ensure AVAudioUnitSampler implements IAVAudioUnitSampler.
var _ IAVAudioUnitSampler = AVAudioUnitSampler{}

// An interface definition for the [AVAudioUnitSampler] class.
//
// # Methods
//
//   - [IAVAudioUnitSampler.MasterGain]
//   - [IAVAudioUnitSampler.SetMasterGain]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler
type IAVAudioUnitSampler interface {
	IAVAudioUnitMIDIInstrument

	// Topic: Methods

	MasterGain() float32
	SetMasterGain(value float32)
}

// Init initializes the instance.
func (a AVAudioUnitSampler) Init() AVAudioUnitSampler {
	rv := objc.Send[AVAudioUnitSampler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnitSampler) Autorelease() AVAudioUnitSampler {
	rv := objc.Send[AVAudioUnitSampler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnitSampler creates a new AVAudioUnitSampler instance.
func NewAVAudioUnitSampler() AVAudioUnitSampler {
	class := getAVAudioUnitSamplerClass()
	rv := objc.Send[AVAudioUnitSampler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioUnitSamplerWithImpl(impl unsafe.Pointer) AVAudioUnitSampler {
	instance := getAVAudioUnitSamplerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioUnitSamplerFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnitSampler/masterGain
func (a AVAudioUnitSampler) MasterGain() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("masterGain"))
	return rv
}
func (a AVAudioUnitSampler) SetMasterGain(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMasterGain:"), value)
}
