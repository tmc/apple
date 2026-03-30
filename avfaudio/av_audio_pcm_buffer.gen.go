// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioPCMBuffer] class.
var (
	_AVAudioPCMBufferClass     AVAudioPCMBufferClass
	_AVAudioPCMBufferClassOnce sync.Once
)

func getAVAudioPCMBufferClass() AVAudioPCMBufferClass {
	_AVAudioPCMBufferClassOnce.Do(func() {
		_AVAudioPCMBufferClass = AVAudioPCMBufferClass{class: objc.GetClass("AVAudioPCMBuffer")}
	})
	return _AVAudioPCMBufferClass
}

// GetAVAudioPCMBufferClass returns the class object for AVAudioPCMBuffer.
func GetAVAudioPCMBufferClass() AVAudioPCMBufferClass {
	return getAVAudioPCMBufferClass()
}

type AVAudioPCMBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioPCMBufferClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioPCMBufferClass) Alloc() AVAudioPCMBuffer {
	rv := objc.Send[AVAudioPCMBuffer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an audio buffer you use with PCM audio formats.
//
// # Overview
//
// The PCM buffer class provides methods that are useful for manipulating
// buffers of audio in PCM format.
//
// # Creating a PCM Audio Buffer
//
//   - [AVAudioPCMBuffer.InitWithPCMFormatFrameCapacity]: Creates a PCM audio buffer instance for PCM audio data.
//   - [AVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator]: Creates a PCM audio buffer instance without copying samples, for PCM audio data, with a specified buffer list and a deallocator closure.
//
// # Getting and Setting the Frame Length
//
//   - [AVAudioPCMBuffer.FrameLength]: The current number of valid sample frames in the buffer.
//   - [AVAudioPCMBuffer.SetFrameLength]
//
// # Accessing PCM Buffer Data
//
//   - [AVAudioPCMBuffer.FloatChannelData]: The buffer’s audio samples as floating point values.
//   - [AVAudioPCMBuffer.FrameCapacity]: The buffer’s capacity, in audio sample frames.
//   - [AVAudioPCMBuffer.Int16ChannelData]: The buffer’s 16-bit integer audio samples.
//   - [AVAudioPCMBuffer.Int32ChannelData]: The buffer’s 32-bit integer audio samples.
//   - [AVAudioPCMBuffer.Stride]: The buffer’s number of interleaved channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer
type AVAudioPCMBuffer struct {
	AVAudioBuffer
}

// AVAudioPCMBufferFromID constructs a [AVAudioPCMBuffer] from an objc.ID.
//
// An object that represents an audio buffer you use with PCM audio formats.
func AVAudioPCMBufferFromID(id objc.ID) AVAudioPCMBuffer {
	return AVAudioPCMBuffer{AVAudioBuffer: AVAudioBufferFromID(id)}
}

// NOTE: AVAudioPCMBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioPCMBuffer] class.
//
// # Creating a PCM Audio Buffer
//
//   - [IAVAudioPCMBuffer.InitWithPCMFormatFrameCapacity]: Creates a PCM audio buffer instance for PCM audio data.
//   - [IAVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator]: Creates a PCM audio buffer instance without copying samples, for PCM audio data, with a specified buffer list and a deallocator closure.
//
// # Getting and Setting the Frame Length
//
//   - [IAVAudioPCMBuffer.FrameLength]: The current number of valid sample frames in the buffer.
//   - [IAVAudioPCMBuffer.SetFrameLength]
//
// # Accessing PCM Buffer Data
//
//   - [IAVAudioPCMBuffer.FloatChannelData]: The buffer’s audio samples as floating point values.
//   - [IAVAudioPCMBuffer.FrameCapacity]: The buffer’s capacity, in audio sample frames.
//   - [IAVAudioPCMBuffer.Int16ChannelData]: The buffer’s 16-bit integer audio samples.
//   - [IAVAudioPCMBuffer.Int32ChannelData]: The buffer’s 32-bit integer audio samples.
//   - [IAVAudioPCMBuffer.Stride]: The buffer’s number of interleaved channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer
type IAVAudioPCMBuffer interface {
	IAVAudioBuffer

	// Topic: Creating a PCM Audio Buffer

	// Creates a PCM audio buffer instance for PCM audio data.
	InitWithPCMFormatFrameCapacity(format IAVAudioFormat, frameCapacity AVAudioFrameCount) AVAudioPCMBuffer
	// Creates a PCM audio buffer instance without copying samples, for PCM audio data, with a specified buffer list and a deallocator closure.
	InitWithPCMFormatBufferListNoCopyDeallocator(format IAVAudioFormat, bufferList objectivec.IObject, deallocator constAudioBufferListHandler) AVAudioPCMBuffer

	// Topic: Getting and Setting the Frame Length

	// The current number of valid sample frames in the buffer.
	FrameLength() AVAudioFrameCount
	SetFrameLength(value AVAudioFrameCount)

	// Topic: Accessing PCM Buffer Data

	// The buffer’s audio samples as floating point values.
	FloatChannelData() unsafe.Pointer
	// The buffer’s capacity, in audio sample frames.
	FrameCapacity() AVAudioFrameCount
	// The buffer’s 16-bit integer audio samples.
	Int16ChannelData() unsafe.Pointer
	// The buffer’s 32-bit integer audio samples.
	Int32ChannelData() unsafe.Pointer
	// The buffer’s number of interleaved channels.
	Stride() uint
}

// Init initializes the instance.
func (a AVAudioPCMBuffer) Init() AVAudioPCMBuffer {
	rv := objc.Send[AVAudioPCMBuffer](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioPCMBuffer) Autorelease() AVAudioPCMBuffer {
	rv := objc.Send[AVAudioPCMBuffer](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioPCMBuffer creates a new AVAudioPCMBuffer instance.
func NewAVAudioPCMBuffer() AVAudioPCMBuffer {
	class := getAVAudioPCMBufferClass()
	rv := objc.Send[AVAudioPCMBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a PCM audio buffer instance for PCM audio data.
//
// format: The format of the PCM audio the buffer contains.
//
// frameCapacity: The capacity of the buffer in PCM sample frames.
//
// # Return Value
//
// A new [AVAudioPCMBuffer] instance, or `nil` if it’s not possible.
//
// # Discussion
//
// The method returns `nil` due to the following reasons:
//
// - The format has zero bytes per frame. - The system can’t represent the
// buffer byte capacity as an unsigned bit-32 integer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/init(pcmFormat:frameCapacity:)
func NewAudioPCMBufferWithPCMFormatFrameCapacity(format IAVAudioFormat, frameCapacity AVAudioFrameCount) AVAudioPCMBuffer {
	instance := getAVAudioPCMBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPCMFormat:frameCapacity:"), format, frameCapacity)
	return AVAudioPCMBufferFromID(rv)
}

// Creates a PCM audio buffer instance for PCM audio data.
//
// format: The format of the PCM audio the buffer contains.
//
// frameCapacity: The capacity of the buffer in PCM sample frames.
//
// # Return Value
//
// A new [AVAudioPCMBuffer] instance, or `nil` if it’s not possible.
//
// # Discussion
//
// The method returns `nil` due to the following reasons:
//
// - The format has zero bytes per frame. - The system can’t represent the
// buffer byte capacity as an unsigned bit-32 integer.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/init(pcmFormat:frameCapacity:)
func (a AVAudioPCMBuffer) InitWithPCMFormatFrameCapacity(format IAVAudioFormat, frameCapacity AVAudioFrameCount) AVAudioPCMBuffer {
	rv := objc.Send[AVAudioPCMBuffer](a.ID, objc.Sel("initWithPCMFormat:frameCapacity:"), format, frameCapacity)
	return rv
}

// Creates a PCM audio buffer instance without copying samples, for PCM audio
// data, with a specified buffer list and a deallocator closure.
//
// format: The format of the PCM audio the buffer contains.
//
// bufferList: The buffer list with the memory to contain the PCM audio data.
//
// deallocator: The closure the method invokes when the resulting PCM buffer object
// deallocates.
//
// # Return Value
//
// A new [AVAudioPCMBuffer] instance, or `nil` if it’s not possible.
//
// # Discussion
//
// Use the deallocator parameter to define your own deallocation behavior for
// the audio buffer list’s underlying memory. The buffer list sent to the
// deallocator is identical to the one you specify, in term of buffer count
// and each buffer’s [mData] and [mDataByteSize] members.
//
// The method returns `nil` due to the following reasons:
//
// - The format has zero bytes per frame. - The buffer you specify has zero
// number of buffers. - The buffer list’s pointer to the buffer of audio
// data is in a `nil` state. - Each of the buffer’s data byte size aren’t
// equal, or if any of the buffers’ data byte size is zero. - There’s a
// mismatch between the format’s number of buffers and the buffer list’s
// size (1 if interleaved, [mChannelsPerFrame] if deinterleaved).
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/init(pcmFormat:bufferListNoCopy:deallocator:)
//
// [mChannelsPerFrame]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription/mChannelsPerFrame
// [mDataByteSize]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBuffer/mDataByteSize
// [mData]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBuffer/mData
func (a AVAudioPCMBuffer) InitWithPCMFormatBufferListNoCopyDeallocator(format IAVAudioFormat, bufferList objectivec.IObject, deallocator constAudioBufferListHandler) AVAudioPCMBuffer {
	_block2, _ := NewconstAudioBufferListBlock(deallocator)
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithPCMFormat:bufferListNoCopy:deallocator:"), format, bufferList, _block2)
	return AVAudioPCMBufferFromID(rv)
}

// The current number of valid sample frames in the buffer.
//
// # Discussion
//
// By default, the `frameLength` property doesn’t have a useful value upon
// creation, so you must set this property before using the buffer. The length
// must be less than or equal to the [FrameCapacity] of the buffer. For
// deinterleaved formats, [FrameCapacity] refers to the size of one
// channel’s worth of audio samples.
//
// You may modify the length of the buffer as part of an operation that
// modifies its contents. Modifying `frameLength` updates the `mDataByteSize`
// field in each of the underlying [AudioBufferList] structure’s
// [AudioBuffer] properties correspondingly, and vice versa.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/frameLength
//
// [AudioBufferList]: https://developer.apple.com/documentation/CoreAudioTypes/AudioBufferList
func (a AVAudioPCMBuffer) FrameLength() AVAudioFrameCount {
	rv := objc.Send[AVAudioFrameCount](a.ID, objc.Sel("frameLength"))
	return AVAudioFrameCount(rv)
}
func (a AVAudioPCMBuffer) SetFrameLength(value AVAudioFrameCount) {
	objc.Send[struct{}](a.ID, objc.Sel("setFrameLength:"), value)
}

// The buffer’s audio samples as floating point values.
//
// # Discussion
//
// The `floatChannelData` property returns pointers to the buffer’s audio
// samples if the buffer’s format is 32-bit float. It returns `nil` if
// it’s another format.
//
// The returned pointer is to `format.ChannelCount()` pointers to float. Each
// of these pointers is to [FrameLength] valid samples, which the class spaces
// by [Stride] samples.
//
// If the format isn’t interleaved, as with the standard deinterleaved float
// format, the pointers point to separate chunks of memory, and the [Stride]
// property value is `1`.
//
// When the format is in an interleaved state, the pointers refer to the same
// buffer of interleaved samples, each offset by `1` frame, and the [Stride]
// property value is the number of interleaved channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/floatChannelData
func (a AVAudioPCMBuffer) FloatChannelData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("floatChannelData"))
	return rv
}

// The buffer’s capacity, in audio sample frames.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/frameCapacity
func (a AVAudioPCMBuffer) FrameCapacity() AVAudioFrameCount {
	rv := objc.Send[AVAudioFrameCount](a.ID, objc.Sel("frameCapacity"))
	return AVAudioFrameCount(rv)
}

// The buffer’s 16-bit integer audio samples.
//
// # Discussion
//
// The `int16ChannelData` property returns the buffer’s audio samples if the
// buffer’s format has 2-byte integer samples, or `nil` if it’s another
// format. For more information, see [FloatChannelData].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/int16ChannelData
func (a AVAudioPCMBuffer) Int16ChannelData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("int16ChannelData"))
	return rv
}

// The buffer’s 32-bit integer audio samples.
//
// # Discussion
//
// The `int32ChannelData` property returns the buffer’s audio samples if the
// buffer’s format has 4-byte integer samples, or `nil` if it’s another
// format. For more information, see [FloatChannelData].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/int32ChannelData
func (a AVAudioPCMBuffer) Int32ChannelData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("int32ChannelData"))
	return rv
}

// The buffer’s number of interleaved channels.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioPCMBuffer/stride
func (a AVAudioPCMBuffer) Stride() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("stride"))
	return rv
}

// InitWithPCMFormatBufferListNoCopyDeallocatorSync is a synchronous wrapper around [AVAudioPCMBuffer.InitWithPCMFormatBufferListNoCopyDeallocator].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioPCMBuffer) InitWithPCMFormatBufferListNoCopyDeallocatorSync(ctx context.Context, format IAVAudioFormat, bufferList objectivec.IObject) (*objectivec.Object, error) {
	done := make(chan *objectivec.Object, 1)
	a.InitWithPCMFormatBufferListNoCopyDeallocator(format, bufferList, func(val *objectivec.Object) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
