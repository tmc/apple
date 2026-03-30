// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioBuffer] class.
var (
	_AVAudioBufferClass     AVAudioBufferClass
	_AVAudioBufferClassOnce sync.Once
)

func getAVAudioBufferClass() AVAudioBufferClass {
	_AVAudioBufferClassOnce.Do(func() {
		_AVAudioBufferClass = AVAudioBufferClass{class: objc.GetClass("AVAudioBuffer")}
	})
	return _AVAudioBufferClass
}

// GetAVAudioBufferClass returns the class object for AVAudioBuffer.
func GetAVAudioBufferClass() AVAudioBufferClass {
	return getAVAudioBufferClass()
}

type AVAudioBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioBufferClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioBufferClass) Alloc() AVAudioBuffer {
	rv := objc.Send[AVAudioBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVAudioBuffer.ByteCapacity]
//   - [AVAudioBuffer.ByteLength]
//   - [AVAudioBuffer.SetByteLength]
//   - [AVAudioBuffer.InitWithFormatByteCapacity]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer
type AVAudioBuffer struct {
	objectivec.Object
}

// AVAudioBufferFromID constructs a [AVAudioBuffer] from an objc.ID.
func AVAudioBufferFromID(id objc.ID) AVAudioBuffer {
	return AVAudioBuffer{objectivec.Object{ID: id}}
}

// Ensure AVAudioBuffer implements IAVAudioBuffer.
var _ IAVAudioBuffer = AVAudioBuffer{}

// An interface definition for the [AVAudioBuffer] class.
//
// # Methods
//
//   - [IAVAudioBuffer.ByteCapacity]
//   - [IAVAudioBuffer.ByteLength]
//   - [IAVAudioBuffer.SetByteLength]
//   - [IAVAudioBuffer.InitWithFormatByteCapacity]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer
type IAVAudioBuffer interface {
	objectivec.IObject

	// Topic: Methods

	ByteCapacity() uint32
	ByteLength() uint32
	SetByteLength(length uint32)
	InitWithFormatByteCapacity(format objectivec.IObject, capacity uint32) AVAudioBuffer
}

// Init initializes the instance.
func (a AVAudioBuffer) Init() AVAudioBuffer {
	rv := objc.Send[AVAudioBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioBuffer) Autorelease() AVAudioBuffer {
	rv := objc.Send[AVAudioBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioBuffer creates a new AVAudioBuffer instance.
func NewAVAudioBuffer() AVAudioBuffer {
	class := getAVAudioBufferClass()
	rv := objc.Send[AVAudioBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/initWithFormat:byteCapacity:
func NewAudioBufferWithFormatByteCapacity(format objectivec.IObject, capacity uint32) AVAudioBuffer {
	instance := getAVAudioBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:byteCapacity:"), format, capacity)
	return AVAudioBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/byteCapacity
func (a AVAudioBuffer) ByteCapacity() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("byteCapacity"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/byteLength
func (a AVAudioBuffer) ByteLength() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("byteLength"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/setByteLength:
func (a AVAudioBuffer) SetByteLength(length uint32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setByteLength:"), length)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/initWithFormat:byteCapacity:
func (a AVAudioBuffer) InitWithFormatByteCapacity(format objectivec.IObject, capacity uint32) AVAudioBuffer {
	rv := objc.Send[AVAudioBuffer](a.ID, objc.Sel("initWithFormat:byteCapacity:"), format, capacity)
	return rv
}
