// Code generated from Apple documentation for texttospeech. DO NOT EDIT.

package texttospeech

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/private/avfaudio"
)

// The class instance for the [TTSWrappedAudioQueueBuffer] class.
var (
	_TTSWrappedAudioQueueBufferClass     TTSWrappedAudioQueueBufferClass
	_TTSWrappedAudioQueueBufferClassOnce sync.Once
)

func getTTSWrappedAudioQueueBufferClass() TTSWrappedAudioQueueBufferClass {
	_TTSWrappedAudioQueueBufferClassOnce.Do(func() {
		_TTSWrappedAudioQueueBufferClass = TTSWrappedAudioQueueBufferClass{class: objc.GetClass("TTSWrappedAudioQueueBuffer")}
	})
	return _TTSWrappedAudioQueueBufferClass
}

// GetTTSWrappedAudioQueueBufferClass returns the class object for TTSWrappedAudioQueueBuffer.
func GetTTSWrappedAudioQueueBufferClass() TTSWrappedAudioQueueBufferClass {
	return getTTSWrappedAudioQueueBufferClass()
}

type TTSWrappedAudioQueueBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TTSWrappedAudioQueueBufferClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TTSWrappedAudioQueueBufferClass) Alloc() TTSWrappedAudioQueueBuffer {
	rv := objc.Send[TTSWrappedAudioQueueBuffer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [TTSWrappedAudioQueueBuffer.AqBuffer]
//   - [TTSWrappedAudioQueueBuffer.SetAqBuffer]
//   - [TTSWrappedAudioQueueBuffer.ByteSize]
//   - [TTSWrappedAudioQueueBuffer.QueuedTimeStamp]
//   - [TTSWrappedAudioQueueBuffer.SetQueuedTimeStamp]
//   - [TTSWrappedAudioQueueBuffer.SetCompletionHandler]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueueBuffer
type TTSWrappedAudioQueueBuffer struct {
	objectivec.Object
}

// TTSWrappedAudioQueueBufferFromID constructs a [TTSWrappedAudioQueueBuffer] from an objc.ID.
func TTSWrappedAudioQueueBufferFromID(id objc.ID) TTSWrappedAudioQueueBuffer {
	return TTSWrappedAudioQueueBuffer{objectivec.Object{ID: id}}
}

// Ensure TTSWrappedAudioQueueBuffer implements ITTSWrappedAudioQueueBuffer.
var _ ITTSWrappedAudioQueueBuffer = TTSWrappedAudioQueueBuffer{}

// An interface definition for the [TTSWrappedAudioQueueBuffer] class.
//
// # Methods
//
//   - [ITTSWrappedAudioQueueBuffer.AqBuffer]
//   - [ITTSWrappedAudioQueueBuffer.SetAqBuffer]
//   - [ITTSWrappedAudioQueueBuffer.ByteSize]
//   - [ITTSWrappedAudioQueueBuffer.QueuedTimeStamp]
//   - [ITTSWrappedAudioQueueBuffer.SetQueuedTimeStamp]
//   - [ITTSWrappedAudioQueueBuffer.SetCompletionHandler]
//
// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueueBuffer
type ITTSWrappedAudioQueueBuffer interface {
	objectivec.IObject

	// Topic: Methods

	AqBuffer() *avfaudio.AudioQueueBufferRef
	SetAqBuffer(value *avfaudio.AudioQueueBufferRef)
	ByteSize() uint64
	QueuedTimeStamp() objectivec.IObject
	SetQueuedTimeStamp(value objectivec.IObject)
	SetCompletionHandler(handler ErrorHandler)
}

// Init initializes the instance.
func (t TTSWrappedAudioQueueBuffer) Init() TTSWrappedAudioQueueBuffer {
	rv := objc.Send[TTSWrappedAudioQueueBuffer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TTSWrappedAudioQueueBuffer) Autorelease() TTSWrappedAudioQueueBuffer {
	rv := objc.Send[TTSWrappedAudioQueueBuffer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTTSWrappedAudioQueueBuffer creates a new TTSWrappedAudioQueueBuffer instance.
func NewTTSWrappedAudioQueueBuffer() TTSWrappedAudioQueueBuffer {
	class := getTTSWrappedAudioQueueBufferClass()
	rv := objc.Send[TTSWrappedAudioQueueBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueueBuffer/setCompletionHandler:
func (t TTSWrappedAudioQueueBuffer) SetCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](t.ID, objc.Sel("setCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueueBuffer/aqBuffer
func (t TTSWrappedAudioQueueBuffer) AqBuffer() *avfaudio.AudioQueueBufferRef {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("aqBuffer"))
	return (*avfaudio.AudioQueueBufferRef)(rv)
}
func (t TTSWrappedAudioQueueBuffer) SetAqBuffer(value *avfaudio.AudioQueueBufferRef) {
	objc.Send[struct{}](t.ID, objc.Sel("setAqBuffer:"), value)
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueueBuffer/byteSize
func (t TTSWrappedAudioQueueBuffer) ByteSize() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("byteSize"))
	return rv
}

// See: https://developer.apple.com/documentation/TextToSpeech/TTSWrappedAudioQueueBuffer/queuedTimeStamp
func (t TTSWrappedAudioQueueBuffer) QueuedTimeStamp() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("queuedTimeStamp"))
	return objectivec.Object{ID: rv}
}
func (t TTSWrappedAudioQueueBuffer) SetQueuedTimeStamp(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setQueuedTimeStamp:"), value)
}

// Set is a synchronous wrapper around [TTSWrappedAudioQueueBuffer.SetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t TTSWrappedAudioQueueBuffer) Set(ctx context.Context) error {
	done := make(chan error, 1)
	t.SetCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
