// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVCaptureAudioDataOutput] class.
var (
	_AVCaptureAudioDataOutputClass     AVCaptureAudioDataOutputClass
	_AVCaptureAudioDataOutputClassOnce sync.Once
)

func getAVCaptureAudioDataOutputClass() AVCaptureAudioDataOutputClass {
	_AVCaptureAudioDataOutputClassOnce.Do(func() {
		_AVCaptureAudioDataOutputClass = AVCaptureAudioDataOutputClass{class: objc.GetClass("AVCaptureAudioDataOutput")}
	})
	return _AVCaptureAudioDataOutputClass
}

// GetAVCaptureAudioDataOutputClass returns the class object for AVCaptureAudioDataOutput.
func GetAVCaptureAudioDataOutputClass() AVCaptureAudioDataOutputClass {
	return getAVCaptureAudioDataOutputClass()
}

type AVCaptureAudioDataOutputClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureAudioDataOutputClass) Alloc() AVCaptureAudioDataOutput {
	rv := objc.Send[AVCaptureAudioDataOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output that records audio and provides access to audio sample
// buffers as they are recorded.
//
// # Configuring audio capture
//
//   - [AVCaptureAudioDataOutput.AudioSettings]: The settings used to decode or re-encode audio before it’s output.
//   - [AVCaptureAudioDataOutput.SetAudioSettings]
//   - [AVCaptureAudioDataOutput.RecommendedAudioSettingsForAssetWriterWithOutputFileType]: Specifies the recommended settings for use with an [AVAssetWriterInput].
//
// # Receiving captured audio data
//
//   - [AVCaptureAudioDataOutput.SetSampleBufferDelegateQueue]: Sets the delegate that will accept captured buffers and the dispatch queue on which the delegate will be called.
//   - [AVCaptureAudioDataOutput.SampleBufferDelegate]: The capture object’s delegate.
//   - [AVCaptureAudioDataOutput.SampleBufferCallbackQueue]: The queue on which delegate callbacks are invoked
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput
type AVCaptureAudioDataOutput struct {
	AVCaptureOutput
}

// AVCaptureAudioDataOutputFromID constructs a [AVCaptureAudioDataOutput] from an objc.ID.
//
// A capture output that records audio and provides access to audio sample
// buffers as they are recorded.
func AVCaptureAudioDataOutputFromID(id objc.ID) AVCaptureAudioDataOutput {
	return AVCaptureAudioDataOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}
// NOTE: AVCaptureAudioDataOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureAudioDataOutput] class.
//
// # Configuring audio capture
//
//   - [IAVCaptureAudioDataOutput.AudioSettings]: The settings used to decode or re-encode audio before it’s output.
//   - [IAVCaptureAudioDataOutput.SetAudioSettings]
//   - [IAVCaptureAudioDataOutput.RecommendedAudioSettingsForAssetWriterWithOutputFileType]: Specifies the recommended settings for use with an [AVAssetWriterInput].
//
// # Receiving captured audio data
//
//   - [IAVCaptureAudioDataOutput.SetSampleBufferDelegateQueue]: Sets the delegate that will accept captured buffers and the dispatch queue on which the delegate will be called.
//   - [IAVCaptureAudioDataOutput.SampleBufferDelegate]: The capture object’s delegate.
//   - [IAVCaptureAudioDataOutput.SampleBufferCallbackQueue]: The queue on which delegate callbacks are invoked
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput
type IAVCaptureAudioDataOutput interface {
	IAVCaptureOutput

	// Topic: Configuring audio capture

	// The settings used to decode or re-encode audio before it’s output.
	AudioSettings() foundation.INSDictionary
	SetAudioSettings(value foundation.INSDictionary)
	// Specifies the recommended settings for use with an [AVAssetWriterInput].
	RecommendedAudioSettingsForAssetWriterWithOutputFileType(outputFileType AVFileType) foundation.INSDictionary

	// Topic: Receiving captured audio data

	// Sets the delegate that will accept captured buffers and the dispatch queue on which the delegate will be called.
	SetSampleBufferDelegateQueue(sampleBufferDelegate AVCaptureAudioDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue)
	// The capture object’s delegate.
	SampleBufferDelegate() AVCaptureAudioDataOutputSampleBufferDelegate
	// The queue on which delegate callbacks are invoked
	SampleBufferCallbackQueue() dispatch.Queue
}

// Init initializes the instance.
func (c AVCaptureAudioDataOutput) Init() AVCaptureAudioDataOutput {
	rv := objc.Send[AVCaptureAudioDataOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureAudioDataOutput) Autorelease() AVCaptureAudioDataOutput {
	rv := objc.Send[AVCaptureAudioDataOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureAudioDataOutput creates a new AVCaptureAudioDataOutput instance.
func NewAVCaptureAudioDataOutput() AVCaptureAudioDataOutput {
	class := getAVCaptureAudioDataOutputClass()
	rv := objc.Send[AVCaptureAudioDataOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Specifies the recommended settings for use with an [AVAssetWriterInput].
//
// outputFileType: Specifies the UTI of the file type to be written. See `File Format UTIs`
// for the defined UTIs.
//
// # Return Value
// 
// A fully populated dictionary of keys and values that are compatible with
// [AVAssetWriter].
//
// # Discussion
// 
// The value of this property is an [NSDictionary] containing values for
// compression settings keys defined in [Audio settings]. This dictionary is
// suitable for use as the [AssetWriterInputWithMediaTypeOutputSettings]
// method’s `outputSettings` parameter when creating an
// [AVAssetWriterInputPixelBufferAdaptor]; for example,
// 
// The dictionary returned contains all necessary keys and values needed to
// create an [AVAssetWriter] instance; see the
// [InitWithMediaTypeOutputSettings] method for a more in depth discussion.
// For QuickTime movie and ISO files, the recommended audio settings will
// always produce output comparable to that of [AVCaptureMovieFileOutput].
// 
// The dictionary of settings is dependent on the current configuration of the
// receiver’s [AVCaptureSession] and its inputs. The settings dictionary may
// change if the session’s configuration changes. As such, you should
// configure your session first, then query the recommended audio settings.
//
// [AVAssetWriterInputPixelBufferAdaptor]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInputPixelBufferAdaptor
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/recommendedAudioSettingsForAssetWriter(writingTo:)
func (c AVCaptureAudioDataOutput) RecommendedAudioSettingsForAssetWriterWithOutputFileType(outputFileType AVFileType) foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recommendedAudioSettingsForAssetWriterWithOutputFileType:"), objc.String(string(outputFileType)))
	return foundation.NSDictionaryFromID(rv)
}
// Sets the delegate that will accept captured buffers and the dispatch queue
// on which the delegate will be called.
//
// sampleBufferDelegate: An object conforming to the [AVCaptureAudioDataOutputSampleBufferDelegate]
// protocol that will receive sample buffers after they are captured.
//
// sampleBufferCallbackQueue: You must pass a serial dispatch to guarantee that audio samples will be
// delivered in order.
// 
// The value may not be [NULL], except when setting the `sampleBufferDelegate`
// to `nil`.
//
// # Discussion
// 
// When a new audio sample buffer is captured it is vended to the sample
// buffer delegate using the
// [CaptureOutputDidOutputSampleBufferFromConnection] delegate method. All
// delegate methods are called on the specified dispatch queue.
// 
// If the queue is blocked when new samples are captured, those samples will
// be automatically dropped when they become sufficiently late. This allows
// you to process existing samples on the same queue without having to manage
// the potential memory usage increases that would otherwise occur when that
// processing is unable to keep up with the rate of incoming samples.
// 
// If you need to minimize the chances of samples being dropped, you should
// specify a queue on which a sufficiently small amount of processing is being
// done outside of receiving sample buffers. When migrating extra processing
// to another queue, you are responsible for ensuring that memory usage does
// not grow without bound from samples that have not been processed.
// 
// # Special considerations
// 
// This method uses [dispatch_retain] and [dispatch_release] to manage the
// queue.
//
// [dispatch_release]: https://developer.apple.com/documentation/Dispatch/dispatch_release
// [dispatch_retain]: https://developer.apple.com/documentation/Dispatch/dispatch_retain
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/setSampleBufferDelegate(_:queue:)
func (c AVCaptureAudioDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate AVCaptureAudioDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, uintptr(sampleBufferCallbackQueue.Handle()))
}

// The settings used to decode or re-encode audio before it’s output.
//
// # Discussion
// 
// The value of this property is a dictionary containing values for audio
// settings keys defined in [Audio settings].
// 
// If the value of this property is `nil`, samples are output in their device
// native format.
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/audioSettings
func (c AVCaptureAudioDataOutput) AudioSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("audioSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c AVCaptureAudioDataOutput) SetAudioSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setAudioSettings:"), value)
}
// The capture object’s delegate.
//
// # Discussion
// 
// You use the delegate to manage incoming data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/sampleBufferDelegate
func (c AVCaptureAudioDataOutput) SampleBufferDelegate() AVCaptureAudioDataOutputSampleBufferDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sampleBufferDelegate"))
	return AVCaptureAudioDataOutputSampleBufferDelegateObjectFromID(rv)
}
// The queue on which delegate callbacks are invoked
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutput/sampleBufferCallbackQueue
func (c AVCaptureAudioDataOutput) SampleBufferCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("sampleBufferCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}

