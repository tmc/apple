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

// An object that represents a buffer of audio data with a format.
//
// # Getting the Buffer Format
//
//   - [AVAudioBuffer.Format]: The format of the audio in the buffer.
//
// # Getting the Audio Buffers
//
//   - [AVAudioBuffer.AudioBufferList]: The buffer’s underlying audio buffer list.
//   - [AVAudioBuffer.MutableAudioBufferList]: A mutable version of the buffer’s underlying audio buffer list.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer
type AVAudioBuffer struct {
	objectivec.Object
}

// AVAudioBufferFromID constructs a [AVAudioBuffer] from an objc.ID.
//
// An object that represents a buffer of audio data with a format.
func AVAudioBufferFromID(id objc.ID) AVAudioBuffer {
	return AVAudioBuffer{objectivec.Object{ID: id}}
}

// NOTE: AVAudioBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioBuffer] class.
//
// # Getting the Buffer Format
//
//   - [IAVAudioBuffer.Format]: The format of the audio in the buffer.
//
// # Getting the Audio Buffers
//
//   - [IAVAudioBuffer.AudioBufferList]: The buffer’s underlying audio buffer list.
//   - [IAVAudioBuffer.MutableAudioBufferList]: A mutable version of the buffer’s underlying audio buffer list.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer
type IAVAudioBuffer interface {
	objectivec.IObject

	// Topic: Getting the Buffer Format

	// The format of the audio in the buffer.
	Format() IAVAudioFormat

	// Topic: Getting the Audio Buffers

	// The buffer’s underlying audio buffer list.
	AudioBufferList() objectivec.IObject
	// A mutable version of the buffer’s underlying audio buffer list.
	MutableAudioBufferList() objectivec.IObject
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

// The format of the audio in the buffer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/format
func (a AVAudioBuffer) Format() IAVAudioFormat {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("format"))
	return AVAudioFormatFromID(objc.ID(rv))
}

// The buffer’s underlying audio buffer list.
//
// # Discussion
//
// A buffer list is a variable length array that contains an array of audio
// buffer instances. You use it with lower-level Core Audio and Audio Toolbox
// API.
//
// You must not modify the buffer list structure, although you can modify
// buffer contents.
//
// The `mDataByteSize` fields of this audio buffer list express the buffer’s
// current [FrameLength].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/audioBufferList
func (a AVAudioBuffer) AudioBufferList() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioBufferList"))
	return objectivec.Object{ID: rv}
}

// A mutable version of the buffer’s underlying audio buffer list.
//
// # Discussion
//
// You use this with some lower-level Core Audio and Audio Toolbox APIs that
// require a mutable [AudioBufferList] (for example, the
// [AudioConverterConvertComplexBuffer(_:_:_:_:)] function).
//
// The `mDataByteSize` fields of this audio buffer list express the buffer’s
// current [FrameCapacity]. If you alter the capacity, modify the buffer’s
// `frameLength` to match.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioBuffer/mutableAudioBufferList
//
// [AudioBufferList]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBufferList
// [AudioConverterConvertComplexBuffer(_:_:_:_:)]: https://developer.apple.com/documentation/AudioToolbox/AudioConverterConvertComplexBuffer(_:_:_:_:)
func (a AVAudioBuffer) MutableAudioBufferList() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mutableAudioBufferList"))
	return objectivec.Object{ID: rv}
}
