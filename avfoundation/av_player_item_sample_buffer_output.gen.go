// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVPlayerItemSampleBufferOutput] class.
var (
	_AVPlayerItemSampleBufferOutputClass     AVPlayerItemSampleBufferOutputClass
	_AVPlayerItemSampleBufferOutputClassOnce sync.Once
)

func getAVPlayerItemSampleBufferOutputClass() AVPlayerItemSampleBufferOutputClass {
	_AVPlayerItemSampleBufferOutputClassOnce.Do(func() {
		_AVPlayerItemSampleBufferOutputClass = AVPlayerItemSampleBufferOutputClass{class: objc.GetClass("AVPlayerItemSampleBufferOutput")}
	})
	return _AVPlayerItemSampleBufferOutputClass
}

// GetAVPlayerItemSampleBufferOutputClass returns the class object for AVPlayerItemSampleBufferOutput.
func GetAVPlayerItemSampleBufferOutputClass() AVPlayerItemSampleBufferOutputClass {
	return getAVPlayerItemSampleBufferOutputClass()
}

type AVPlayerItemSampleBufferOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemSampleBufferOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemSampleBufferOutputClass) Alloc() AVPlayerItemSampleBufferOutput {
	rv := objc.Send[AVPlayerItemSampleBufferOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// [AVPlayerItemSampleBufferOutput] delivers [CMSampleBuffers] for
// [AVPlayerItem] playback.
//
// # Overview
//
// Playback only happens when the [AVPlayerItem] is the current item of its
// [AVPlayer].
//
// Create an [AVPlayerItemSampleBufferOutput] with a
// [AVPlayerItemSampleBufferOutputAudioConfiguration] to configure it to
// deliver [CMSampleBuffers] containing the decoded audio, and attach it to
// the [AVPlayerItem] using `-[AVPlayerItem ]`; the audio will be in the
// format specified by the configuration object’s `requestedAudioFormat`.
//
// Note that [AVPlayerItemSampleBufferOutput] may be used to pull
// [CMSampleBuffers] far ahead of the current play time. Practical use
// requires clients to monitor the item timebase time, and pause pulling when
// they have received CMSampleBuffers sufficient to prepare for
// near-term-future playback or processing.
//
// Marker-only [CMSampleBuffers] may be among those returned; you can detect
// and skip these by testing whether
// `CMSampleBufferGetNumSamples(sampleBuffer) == 0`.
//
// The output [CMSampleBuffers] will have appropriate
// OutputPresentationTimeStamps for playback, but beyond that, synchronizing
// presentation to the AVPlayerItem’s timebase is entirely up to the client.
//
// Currently supported for HLS [AVPlayerItems] only, and only for delivering
// decoded PCM audio.
//
// # Creating a sample buffer output
//
//   - [AVPlayerItemSampleBufferOutput.InitWithConfiguration]: Initializes an instance of [AVPlayerItemSampleBufferOutput](<https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput>).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput
type AVPlayerItemSampleBufferOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemSampleBufferOutputFromID constructs a [AVPlayerItemSampleBufferOutput] from an objc.ID.
//
// [AVPlayerItemSampleBufferOutput] delivers [CMSampleBuffers] for
// [AVPlayerItem] playback.
func AVPlayerItemSampleBufferOutputFromID(id objc.ID) AVPlayerItemSampleBufferOutput {
	return AVPlayerItemSampleBufferOutput{AVPlayerItemOutput: AVPlayerItemOutputFromID(id)}
}

// NOTE: AVPlayerItemSampleBufferOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemSampleBufferOutput] class.
//
// # Creating a sample buffer output
//
//   - [IAVPlayerItemSampleBufferOutput.InitWithConfiguration]: Initializes an instance of [AVPlayerItemSampleBufferOutput](<https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput>).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput
type IAVPlayerItemSampleBufferOutput interface {
	IAVPlayerItemOutput

	// Topic: Creating a sample buffer output

	// Initializes an instance of [AVPlayerItemSampleBufferOutput](<https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput>).
	InitWithConfiguration(configuration IAVPlayerItemSampleBufferOutputConfiguration) AVPlayerItemSampleBufferOutput
}

// Init initializes the instance.
func (p AVPlayerItemSampleBufferOutput) Init() AVPlayerItemSampleBufferOutput {
	rv := objc.Send[AVPlayerItemSampleBufferOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemSampleBufferOutput) Autorelease() AVPlayerItemSampleBufferOutput {
	rv := objc.Send[AVPlayerItemSampleBufferOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemSampleBufferOutput creates a new AVPlayerItemSampleBufferOutput instance.
func NewAVPlayerItemSampleBufferOutput() AVPlayerItemSampleBufferOutput {
	class := getAVPlayerItemSampleBufferOutputClass()
	rv := objc.Send[AVPlayerItemSampleBufferOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an instance of [AVPlayerItemSampleBufferOutput].
//
// configuration: Specifies the kind and format of media data to be delivered.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput/init(configuration:)
func NewPlayerItemSampleBufferOutputWithConfiguration(configuration IAVPlayerItemSampleBufferOutputConfiguration) AVPlayerItemSampleBufferOutput {
	instance := getAVPlayerItemSampleBufferOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return AVPlayerItemSampleBufferOutputFromID(rv)
}

// Initializes an instance of [AVPlayerItemSampleBufferOutput].
//
// configuration: Specifies the kind and format of media data to be delivered.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutput/init(configuration:)
func (p AVPlayerItemSampleBufferOutput) InitWithConfiguration(configuration IAVPlayerItemSampleBufferOutputConfiguration) AVPlayerItemSampleBufferOutput {
	rv := objc.Send[AVPlayerItemSampleBufferOutput](p.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}
