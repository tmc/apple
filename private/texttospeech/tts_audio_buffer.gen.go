// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TTSAudioBuffer] class.
var (
	_TTSAudioBufferClass     TTSAudioBufferClass
	_TTSAudioBufferClassOnce sync.Once
)

func getTTSAudioBufferClass() TTSAudioBufferClass {
	_TTSAudioBufferClassOnce.Do(func() {
		_TTSAudioBufferClass = TTSAudioBufferClass{class: objc.GetClass("TTSAudioBuffer")}
	})
	return _TTSAudioBufferClass
}

// GetTTSAudioBufferClass returns the class object for TTSAudioBuffer.
func GetTTSAudioBufferClass() TTSAudioBufferClass {
	return getTTSAudioBufferClass()
}

type TTSAudioBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSAudioBufferClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSAudioBufferClass) Alloc() TTSAudioBuffer {
	rv := objc.Send[TTSAudioBuffer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSAudioBuffer.AvBuffer]
//   - [TTSAudioBuffer.Format]
//   - [TTSAudioBuffer.FrameCapacity]
//   - [TTSAudioBuffer.FrameLength]
//   - [TTSAudioBuffer.SetFrameLength]
//   - [TTSAudioBuffer.MutableAudioBufferList]
//   - [TTSAudioBuffer.InitWithAVBuffer]
//   - [TTSAudioBuffer.InitWithFormatFrameCapacity]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer
type TTSAudioBuffer struct {
	objectivec.Object
}

// TTSAudioBufferFromID constructs a [TTSAudioBuffer] from an objc.ID.
func TTSAudioBufferFromID(id objc.ID) TTSAudioBuffer {
	return TTSAudioBuffer{objectivec.Object{ID: id}}
}

// Ensure TTSAudioBuffer implements ITTSAudioBuffer.
var _ ITTSAudioBuffer = TTSAudioBuffer{}

// An interface definition for the [TTSAudioBuffer] class.
//
// # Methods
//
//   - [ITTSAudioBuffer.AvBuffer]
//   - [ITTSAudioBuffer.Format]
//   - [ITTSAudioBuffer.FrameCapacity]
//   - [ITTSAudioBuffer.FrameLength]
//   - [ITTSAudioBuffer.SetFrameLength]
//   - [ITTSAudioBuffer.MutableAudioBufferList]
//   - [ITTSAudioBuffer.InitWithAVBuffer]
//   - [ITTSAudioBuffer.InitWithFormatFrameCapacity]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer
type ITTSAudioBuffer interface {
	objectivec.IObject

	// Topic: Methods

	AvBuffer() avfaudio.AVAudioPCMBuffer
	Format() ITTSAudioFormat
	FrameCapacity() uint32
	FrameLength() uint32
	SetFrameLength(value uint32)
	MutableAudioBufferList() unsafe.Pointer
	InitWithAVBuffer(aVBuffer objectivec.IObject) TTSAudioBuffer
	InitWithFormatFrameCapacity(format objectivec.IObject, capacity uint32) TTSAudioBuffer
}

// Init initializes the instance.
func (t TTSAudioBuffer) Init() TTSAudioBuffer {
	rv := objc.Send[TTSAudioBuffer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSAudioBuffer) Autorelease() TTSAudioBuffer {
	rv := objc.Send[TTSAudioBuffer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSAudioBuffer creates a new TTSAudioBuffer instance.
func NewTTSAudioBuffer() TTSAudioBuffer {
	class := getTTSAudioBufferClass()
	rv := objc.Send[TTSAudioBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/initWithAVBuffer:
func NewTTSAudioBufferWithAVBuffer(aVBuffer objectivec.IObject) TTSAudioBuffer {
	instance := getTTSAudioBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAVBuffer:"), aVBuffer)
	return TTSAudioBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/initWithFormat:frameCapacity:
func NewTTSAudioBufferWithFormatFrameCapacity(format objectivec.IObject, capacity uint32) TTSAudioBuffer {
	instance := getTTSAudioBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFormat:frameCapacity:"), format, capacity)
	return TTSAudioBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/initWithAVBuffer:
func (t TTSAudioBuffer) InitWithAVBuffer(aVBuffer objectivec.IObject) TTSAudioBuffer {
	rv := objc.Send[TTSAudioBuffer](t.ID, objc.Sel("initWithAVBuffer:"), aVBuffer)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/initWithFormat:frameCapacity:
func (t TTSAudioBuffer) InitWithFormatFrameCapacity(format objectivec.IObject, capacity uint32) TTSAudioBuffer {
	rv := objc.Send[TTSAudioBuffer](t.ID, objc.Sel("initWithFormat:frameCapacity:"), format, capacity)
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/avBuffer
func (t TTSAudioBuffer) AvBuffer() avfaudio.AVAudioPCMBuffer {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("avBuffer"))
	return avfaudio.AVAudioPCMBufferFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/format
func (t TTSAudioBuffer) Format() ITTSAudioFormat {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("format"))
	return TTSAudioFormatFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/frameCapacity
func (t TTSAudioBuffer) FrameCapacity() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("frameCapacity"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/frameLength
func (t TTSAudioBuffer) FrameLength() uint32 {
	rv := objc.Send[uint32](t.ID, objc.Sel("frameLength"))
	return rv
}
func (t TTSAudioBuffer) SetFrameLength(value uint32) {
	objc.Send[struct{}](t.ID, objc.Sel("setFrameLength:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSAudioBuffer/mutableAudioBufferList
func (t TTSAudioBuffer) MutableAudioBufferList() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("mutableAudioBufferList"))
	return rv
}
