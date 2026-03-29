// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioSharedPCMBuffer] class.
var (
	_AVAudioSharedPCMBufferClass     AVAudioSharedPCMBufferClass
	_AVAudioSharedPCMBufferClassOnce sync.Once
)

func getAVAudioSharedPCMBufferClass() AVAudioSharedPCMBufferClass {
	_AVAudioSharedPCMBufferClassOnce.Do(func() {
		_AVAudioSharedPCMBufferClass = AVAudioSharedPCMBufferClass{class: objc.GetClass("AVAudioSharedPCMBuffer")}
	})
	return _AVAudioSharedPCMBufferClass
}

// GetAVAudioSharedPCMBufferClass returns the class object for AVAudioSharedPCMBuffer.
func GetAVAudioSharedPCMBufferClass() AVAudioSharedPCMBufferClass {
	return getAVAudioSharedPCMBufferClass()
}

type AVAudioSharedPCMBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioSharedPCMBufferClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioSharedPCMBufferClass) Alloc() AVAudioSharedPCMBuffer {
	rv := objc.Send[AVAudioSharedPCMBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVAudioSharedPCMBuffer.BackedBySharedMemory]
//   - [AVAudioSharedPCMBuffer.SharedBufferToken]
//   - [AVAudioSharedPCMBuffer.InitWithPCMFormatFrameCapacity]
//   - [AVAudioSharedPCMBuffer.InitWithPCMFormatSharedBufferToken]
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer
type AVAudioSharedPCMBuffer struct {
	AVAudioPCMBuffer
}

// AVAudioSharedPCMBufferFromID constructs a [AVAudioSharedPCMBuffer] from an objc.ID.
func AVAudioSharedPCMBufferFromID(id objc.ID) AVAudioSharedPCMBuffer {
	return AVAudioSharedPCMBuffer{AVAudioPCMBuffer: AVAudioPCMBufferFromID(id)}
}
// Ensure AVAudioSharedPCMBuffer implements IAVAudioSharedPCMBuffer.
var _ IAVAudioSharedPCMBuffer = AVAudioSharedPCMBuffer{}

// An interface definition for the [AVAudioSharedPCMBuffer] class.
//
// # Methods
//
//   - [IAVAudioSharedPCMBuffer.BackedBySharedMemory]
//   - [IAVAudioSharedPCMBuffer.SharedBufferToken]
//   - [IAVAudioSharedPCMBuffer.InitWithPCMFormatFrameCapacity]
//   - [IAVAudioSharedPCMBuffer.InitWithPCMFormatSharedBufferToken]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer
type IAVAudioSharedPCMBuffer interface {
	IAVAudioPCMBuffer

	// Topic: Methods

	BackedBySharedMemory() bool
	SharedBufferToken() IAVAudioSharedBufferToken
	InitWithPCMFormatFrameCapacity(pCMFormat objectivec.IObject, capacity uint32) AVAudioSharedPCMBuffer
	InitWithPCMFormatSharedBufferToken(pCMFormat objectivec.IObject, token objectivec.IObject) AVAudioSharedPCMBuffer
}

// Init initializes the instance.
func (a AVAudioSharedPCMBuffer) Init() AVAudioSharedPCMBuffer {
	rv := objc.Send[AVAudioSharedPCMBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioSharedPCMBuffer) Autorelease() AVAudioSharedPCMBuffer {
	rv := objc.Send[AVAudioSharedPCMBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioSharedPCMBuffer creates a new AVAudioSharedPCMBuffer instance.
func NewAVAudioSharedPCMBuffer() AVAudioSharedPCMBuffer {
	class := getAVAudioSharedPCMBufferClass()
	rv := objc.Send[AVAudioSharedPCMBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/initWithFormat:byteCapacity:
func NewAudioSharedPCMBufferWithFormatByteCapacity(format objectivec.IObject, capacity uint32) AVAudioSharedPCMBuffer {
	instance := getAVAudioSharedPCMBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:byteCapacity:"), format, capacity)
	return AVAudioSharedPCMBufferFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer/initWithPCMFormat:frameCapacity:
func NewAudioSharedPCMBufferWithPCMFormatFrameCapacity(pCMFormat objectivec.IObject, capacity uint32) AVAudioSharedPCMBuffer {
	instance := getAVAudioSharedPCMBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPCMFormat:frameCapacity:"), pCMFormat, capacity)
	return AVAudioSharedPCMBufferFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer/initWithPCMFormat:sharedBufferToken:
func NewAudioSharedPCMBufferWithPCMFormatSharedBufferToken(pCMFormat objectivec.IObject, token objectivec.IObject) AVAudioSharedPCMBuffer {
	instance := getAVAudioSharedPCMBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPCMFormat:sharedBufferToken:"), pCMFormat, token)
	return AVAudioSharedPCMBufferFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer/initWithPCMFormat:frameCapacity:
func (a AVAudioSharedPCMBuffer) InitWithPCMFormatFrameCapacity(pCMFormat objectivec.IObject, capacity uint32) AVAudioSharedPCMBuffer {
	rv := objc.Send[AVAudioSharedPCMBuffer](a.ID, objc.Sel("initWithPCMFormat:frameCapacity:"), pCMFormat, capacity)
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer/initWithPCMFormat:sharedBufferToken:
func (a AVAudioSharedPCMBuffer) InitWithPCMFormatSharedBufferToken(pCMFormat objectivec.IObject, token objectivec.IObject) AVAudioSharedPCMBuffer {
	rv := objc.Send[AVAudioSharedPCMBuffer](a.ID, objc.Sel("initWithPCMFormat:sharedBufferToken:"), pCMFormat, token)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer/backedBySharedMemory
func (a AVAudioSharedPCMBuffer) BackedBySharedMemory() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("backedBySharedMemory"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioSharedPCMBuffer/sharedBufferToken
func (a AVAudioSharedPCMBuffer) SharedBufferToken() IAVAudioSharedBufferToken {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sharedBufferToken"))
	return AVAudioSharedBufferTokenFromID(objc.ID(rv))
}

