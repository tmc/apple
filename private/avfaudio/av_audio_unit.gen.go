// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAudioUnit] class.
var (
	_AVAudioUnitClass     AVAudioUnitClass
	_AVAudioUnitClassOnce sync.Once
)

func getAVAudioUnitClass() AVAudioUnitClass {
	_AVAudioUnitClassOnce.Do(func() {
		_AVAudioUnitClass = AVAudioUnitClass{class: objc.GetClass("AVAudioUnit")}
	})
	return _AVAudioUnitClass
}

// GetAVAudioUnitClass returns the class object for AVAudioUnit.
func GetAVAudioUnitClass() AVAudioUnitClass {
	return getAVAudioUnitClass()
}

type AVAudioUnitClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioUnitClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioUnitClass) Alloc() AVAudioUnit {
	rv := objc.Send[AVAudioUnit](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioUnit.SetValueForParam]
//   - [AVAudioUnit.ValueForParam]
//   - [AVAudioUnit.AUAudioUnit]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
type AVAudioUnit struct {
	AVAudioNode
}

// AVAudioUnitFromID constructs a [AVAudioUnit] from an objc.ID.
func AVAudioUnitFromID(id objc.ID) AVAudioUnit {
	return AVAudioUnit{AVAudioNode: AVAudioNodeFromID(id)}
}

// Ensure AVAudioUnit implements IAVAudioUnit.
var _ IAVAudioUnit = AVAudioUnit{}

// An interface definition for the [AVAudioUnit] class.
//
// # Methods
//
//   - [IAVAudioUnit.SetValueForParam]
//   - [IAVAudioUnit.ValueForParam]
//   - [IAVAudioUnit.AUAudioUnit]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit
type IAVAudioUnit interface {
	IAVAudioNode

	// Topic: Methods

	SetValueForParam(value float32, param uint32) bool
	ValueForParam(param uint32) float32
	AUAudioUnit() unsafe.Pointer
}

// Init initializes the instance.
func (a AVAudioUnit) Init() AVAudioUnit {
	rv := objc.Send[AVAudioUnit](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioUnit) Autorelease() AVAudioUnit {
	rv := objc.Send[AVAudioUnit](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioUnit creates a new AVAudioUnit instance.
func NewAVAudioUnit() AVAudioUnit {
	class := getAVAudioUnitClass()
	rv := objc.Send[AVAudioUnit](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioNode/initWithImpl:
func NewAudioUnitWithImpl(impl unsafe.Pointer) AVAudioUnit {
	instance := getAVAudioUnitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithImpl:"), impl)
	return AVAudioUnitFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/setValue:forParam:
func (a AVAudioUnit) SetValueForParam(value float32, param uint32) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setValue:forParam:"), value, param)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/valueForParam:
func (a AVAudioUnit) ValueForParam(param uint32) float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("valueForParam:"), param)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioUnit/AUAudioUnit
func (a AVAudioUnit) AUAudioUnit() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("AUAudioUnit"))
	return rv
}
