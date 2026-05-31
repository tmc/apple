// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech

import (
	"sync"

	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SFSpeechAudioBufferRecognitionRequest] class.
var (
	_SFSpeechAudioBufferRecognitionRequestClass     SFSpeechAudioBufferRecognitionRequestClass
	_SFSpeechAudioBufferRecognitionRequestClassOnce sync.Once
)

func getSFSpeechAudioBufferRecognitionRequestClass() SFSpeechAudioBufferRecognitionRequestClass {
	_SFSpeechAudioBufferRecognitionRequestClassOnce.Do(func() {
		_SFSpeechAudioBufferRecognitionRequestClass = SFSpeechAudioBufferRecognitionRequestClass{class: objc.GetClass("SFSpeechAudioBufferRecognitionRequest")}
	})
	return _SFSpeechAudioBufferRecognitionRequestClass
}

// GetSFSpeechAudioBufferRecognitionRequestClass returns the class object for SFSpeechAudioBufferRecognitionRequest.
func GetSFSpeechAudioBufferRecognitionRequestClass() SFSpeechAudioBufferRecognitionRequestClass {
	return getSFSpeechAudioBufferRecognitionRequestClass()
}

type SFSpeechAudioBufferRecognitionRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SFSpeechAudioBufferRecognitionRequestClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SFSpeechAudioBufferRecognitionRequestClass) Alloc() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A request to recognize speech from captured audio content, such as audio
// from the device’s microphone.
//
// # Overview
//
// Use an [SFSpeechAudioBufferRecognitionRequest] object to perform speech
// recognition on live audio, or on a set of existing audio buffers. For
// example, use this request object to route audio from a device’s
// microphone to the speech recognizer.
//
// The request object contains no audio initially. As you capture audio, call
// [SFSpeechAudioBufferRecognitionRequest.AppendAudioPCMBuffer] or
// [SFSpeechAudioBufferRecognitionRequest.AppendAudioSampleBuffer] to add
// audio samples to the request object. The speech recognizer continuously
// analyzes the audio you appended, stopping only when you call the
// [SFSpeechAudioBufferRecognitionRequest.EndAudio] method. You must call
// [SFSpeechAudioBufferRecognitionRequest.EndAudio] explicitly to stop the
// speech recognition process.
//
// For a complete example of how to use audio buffers with speech recognition,
// see [SpeakToMe: Using Speech Recognition with AVAudioEngine].
//
// # Appending Audio Buffers
//
//   - [SFSpeechAudioBufferRecognitionRequest.AppendAudioPCMBuffer]: Appends audio in the PCM format to the end of the recognition request.
//   - [SFSpeechAudioBufferRecognitionRequest.AppendAudioSampleBuffer]: Appends audio to the end of the recognition request.
//   - [SFSpeechAudioBufferRecognitionRequest.EndAudio]: Marks the end of audio input for the recognition request.
//
// # Getting the Audio Format
//
//   - [SFSpeechAudioBufferRecognitionRequest.NativeAudioFormat]: The preferred audio format for optimal speech recognition.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest
//
// [SpeakToMe: Using Speech Recognition with AVAudioEngine]: https://developer.apple.com/library/archive/samplecode/SpeakToMe/Introduction/Intro.html#//apple_ref/doc/uid/TP40017110
type SFSpeechAudioBufferRecognitionRequest struct {
	SFSpeechRecognitionRequest
}

// SFSpeechAudioBufferRecognitionRequestFromID constructs a [SFSpeechAudioBufferRecognitionRequest] from an objc.ID.
//
// A request to recognize speech from captured audio content, such as audio
// from the device’s microphone.
func SFSpeechAudioBufferRecognitionRequestFromID(id objc.ID) SFSpeechAudioBufferRecognitionRequest {
	return SFSpeechAudioBufferRecognitionRequest{SFSpeechRecognitionRequest: SFSpeechRecognitionRequestFromID(id)}
}

// NOTE: SFSpeechAudioBufferRecognitionRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SFSpeechAudioBufferRecognitionRequest] class.
//
// # Appending Audio Buffers
//
//   - [ISFSpeechAudioBufferRecognitionRequest.AppendAudioPCMBuffer]: Appends audio in the PCM format to the end of the recognition request.
//   - [ISFSpeechAudioBufferRecognitionRequest.AppendAudioSampleBuffer]: Appends audio to the end of the recognition request.
//   - [ISFSpeechAudioBufferRecognitionRequest.EndAudio]: Marks the end of audio input for the recognition request.
//
// # Getting the Audio Format
//
//   - [ISFSpeechAudioBufferRecognitionRequest.NativeAudioFormat]: The preferred audio format for optimal speech recognition.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest
type ISFSpeechAudioBufferRecognitionRequest interface {
	ISFSpeechRecognitionRequest

	// Topic: Appending Audio Buffers

	// Appends audio in the PCM format to the end of the recognition request.
	AppendAudioPCMBuffer(audioPCMBuffer avfaudio.AVAudioPCMBuffer)
	// Appends audio to the end of the recognition request.
	AppendAudioSampleBuffer(sampleBuffer coremedia.CMSampleBufferRef)
	// Marks the end of audio input for the recognition request.
	EndAudio()

	// Topic: Getting the Audio Format

	// The preferred audio format for optimal speech recognition.
	NativeAudioFormat() avfaudio.AVAudioFormat
}

// Init initializes the instance.
func (s SFSpeechAudioBufferRecognitionRequest) Init() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SFSpeechAudioBufferRecognitionRequest) Autorelease() SFSpeechAudioBufferRecognitionRequest {
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSFSpeechAudioBufferRecognitionRequest creates a new SFSpeechAudioBufferRecognitionRequest instance.
func NewSFSpeechAudioBufferRecognitionRequest() SFSpeechAudioBufferRecognitionRequest {
	class := getSFSpeechAudioBufferRecognitionRequestClass()
	rv := objc.Send[SFSpeechAudioBufferRecognitionRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Appends audio in the PCM format to the end of the recognition request.
//
// audioPCMBuffer: An audio buffer that contains audio in the PCM format.
//
// # Discussion
//
// The audio must be in a native format and uncompressed.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/append(_:)
func (s SFSpeechAudioBufferRecognitionRequest) AppendAudioPCMBuffer(audioPCMBuffer avfaudio.AVAudioPCMBuffer) {
	objc.Send[objc.ID](s.ID, objc.Sel("appendAudioPCMBuffer:"), audioPCMBuffer)
}

// Appends audio to the end of the recognition request.
//
// sampleBuffer: A buffer of audio.
//
// # Discussion
//
// The audio must be in a native format.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/appendAudioSampleBuffer(_:)
func (s SFSpeechAudioBufferRecognitionRequest) AppendAudioSampleBuffer(sampleBuffer coremedia.CMSampleBufferRef) {
	objc.Send[objc.ID](s.ID, objc.Sel("appendAudioSampleBuffer:"), sampleBuffer)
}

// Marks the end of audio input for the recognition request.
//
// # Discussion
//
// Call this method explicitly to let the speech recognizer know that no more
// audio input is coming.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/endAudio()
func (s SFSpeechAudioBufferRecognitionRequest) EndAudio() {
	objc.Send[objc.ID](s.ID, objc.Sel("endAudio"))
}

// The preferred audio format for optimal speech recognition.
//
// # Discussion
//
// Use the audio format in this property as a hint for optimal recording, but
// don’t depend on the value remaining unchanged.
//
// See: https://developer.apple.com/documentation/Speech/SFSpeechAudioBufferRecognitionRequest/nativeAudioFormat
func (s SFSpeechAudioBufferRecognitionRequest) NativeAudioFormat() avfaudio.AVAudioFormat {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nativeAudioFormat"))
	return avfaudio.AVAudioFormatFromID(objc.ID(rv))
}
