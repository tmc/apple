// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVCaptureVideoDataOutput] class.
var (
	_AVCaptureVideoDataOutputClass     AVCaptureVideoDataOutputClass
	_AVCaptureVideoDataOutputClassOnce sync.Once
)

func getAVCaptureVideoDataOutputClass() AVCaptureVideoDataOutputClass {
	_AVCaptureVideoDataOutputClassOnce.Do(func() {
		_AVCaptureVideoDataOutputClass = AVCaptureVideoDataOutputClass{class: objc.GetClass("AVCaptureVideoDataOutput")}
	})
	return _AVCaptureVideoDataOutputClass
}

// GetAVCaptureVideoDataOutputClass returns the class object for AVCaptureVideoDataOutput.
func GetAVCaptureVideoDataOutputClass() AVCaptureVideoDataOutputClass {
	return getAVCaptureVideoDataOutputClass()
}

type AVCaptureVideoDataOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureVideoDataOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureVideoDataOutputClass) Alloc() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output that records video and provides access to video frames for
// processing.
//
// # Overview
// 
// Use this output to process compressed or uncompressed frames from the
// captured video. You can access the frames with the
// [CaptureOutputDidOutputSampleBufferFromConnection] delegate method.
// 
// This object supports compressed video data output for macOS only. It can
// output pixel buffers in several pixel formats. Consider the usability and
// performance characteristics of these formats and choose the best format for
// your app.
//
// # Configuring video capture
//
//   - [AVCaptureVideoDataOutput.VideoSettings]: A dictionary that contains the compression settings for the output.
//   - [AVCaptureVideoDataOutput.SetVideoSettings]
//   - [AVCaptureVideoDataOutput.AlwaysDiscardsLateVideoFrames]: Indicates whether to drop video frames if they arrive late.
//   - [AVCaptureVideoDataOutput.SetAlwaysDiscardsLateVideoFrames]
//   - [AVCaptureVideoDataOutput.PreservesDynamicHDRMetadata]: Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//   - [AVCaptureVideoDataOutput.SetPreservesDynamicHDRMetadata]
//   - [AVCaptureVideoDataOutput.RecommendedMediaTimeScaleForAssetWriter]: Indicates the recommended media timescale for the video track.
//   - [AVCaptureVideoDataOutput.RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType]: Recommends movie-level metadata for a particular video codec type and output file type, to be used with an asset writer input.
//   - [AVCaptureVideoDataOutput.RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType]: Returns a video settings dictionary appropriate for capturing video to a file with the specified codec and type.
//   - [AVCaptureVideoDataOutput.RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL]: Returns a dictionary of recommended output settings for writing the specified code, file type, and output URL.
//   - [AVCaptureVideoDataOutput.RecommendedVideoSettingsForAssetWriterWithOutputFileType]: Specifies the recommended settings for use with an AVAssetWriterInput.
//
// # Retrieving supported video types
//
//   - [AVCaptureVideoDataOutput.AvailableVideoPixelFormatTypes]: The video pixel formats the output supports.
//   - [AVCaptureVideoDataOutput.SetAvailableVideoPixelFormatTypes]
//   - [AVCaptureVideoDataOutput.AvailableVideoCodecTypes]: The video codecs that the output supports.
//   - [AVCaptureVideoDataOutput.AvailableVideoCodecTypesForAssetWriterWithOutputFileType]: The video codecs that the output supports for writing video to the output file.
//
// # Receiving captured video data
//
//   - [AVCaptureVideoDataOutput.SetSampleBufferDelegateQueue]: Sets the sample buffer delegate and the queue for invoking callbacks.
//   - [AVCaptureVideoDataOutput.SampleBufferDelegate]: The capture object’s delegate.
//   - [AVCaptureVideoDataOutput.SampleBufferCallbackQueue]: The queue on which the system invokes delegate callbacks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput
type AVCaptureVideoDataOutput struct {
	AVCaptureOutput
}

// AVCaptureVideoDataOutputFromID constructs a [AVCaptureVideoDataOutput] from an objc.ID.
//
// A capture output that records video and provides access to video frames for
// processing.
func AVCaptureVideoDataOutputFromID(id objc.ID) AVCaptureVideoDataOutput {
	return AVCaptureVideoDataOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}
// NOTE: AVCaptureVideoDataOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureVideoDataOutput] class.
//
// # Configuring video capture
//
//   - [IAVCaptureVideoDataOutput.VideoSettings]: A dictionary that contains the compression settings for the output.
//   - [IAVCaptureVideoDataOutput.SetVideoSettings]
//   - [IAVCaptureVideoDataOutput.AlwaysDiscardsLateVideoFrames]: Indicates whether to drop video frames if they arrive late.
//   - [IAVCaptureVideoDataOutput.SetAlwaysDiscardsLateVideoFrames]
//   - [IAVCaptureVideoDataOutput.PreservesDynamicHDRMetadata]: Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
//   - [IAVCaptureVideoDataOutput.SetPreservesDynamicHDRMetadata]
//   - [IAVCaptureVideoDataOutput.RecommendedMediaTimeScaleForAssetWriter]: Indicates the recommended media timescale for the video track.
//   - [IAVCaptureVideoDataOutput.RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType]: Recommends movie-level metadata for a particular video codec type and output file type, to be used with an asset writer input.
//   - [IAVCaptureVideoDataOutput.RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType]: Returns a video settings dictionary appropriate for capturing video to a file with the specified codec and type.
//   - [IAVCaptureVideoDataOutput.RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL]: Returns a dictionary of recommended output settings for writing the specified code, file type, and output URL.
//   - [IAVCaptureVideoDataOutput.RecommendedVideoSettingsForAssetWriterWithOutputFileType]: Specifies the recommended settings for use with an AVAssetWriterInput.
//
// # Retrieving supported video types
//
//   - [IAVCaptureVideoDataOutput.AvailableVideoPixelFormatTypes]: The video pixel formats the output supports.
//   - [IAVCaptureVideoDataOutput.SetAvailableVideoPixelFormatTypes]
//   - [IAVCaptureVideoDataOutput.AvailableVideoCodecTypes]: The video codecs that the output supports.
//   - [IAVCaptureVideoDataOutput.AvailableVideoCodecTypesForAssetWriterWithOutputFileType]: The video codecs that the output supports for writing video to the output file.
//
// # Receiving captured video data
//
//   - [IAVCaptureVideoDataOutput.SetSampleBufferDelegateQueue]: Sets the sample buffer delegate and the queue for invoking callbacks.
//   - [IAVCaptureVideoDataOutput.SampleBufferDelegate]: The capture object’s delegate.
//   - [IAVCaptureVideoDataOutput.SampleBufferCallbackQueue]: The queue on which the system invokes delegate callbacks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput
type IAVCaptureVideoDataOutput interface {
	IAVCaptureOutput

	// Topic: Configuring video capture

	// A dictionary that contains the compression settings for the output.
	VideoSettings() foundation.INSDictionary
	SetVideoSettings(value foundation.INSDictionary)
	// Indicates whether to drop video frames if they arrive late.
	AlwaysDiscardsLateVideoFrames() bool
	SetAlwaysDiscardsLateVideoFrames(value bool)
	// Indicates whether the receiver should preserve dynamic HDR metadata as an attachment on the output sample buffer’s underlying pixel buffer.
	PreservesDynamicHDRMetadata() bool
	SetPreservesDynamicHDRMetadata(value bool)
	// Indicates the recommended media timescale for the video track.
	RecommendedMediaTimeScaleForAssetWriter() int32
	// Recommends movie-level metadata for a particular video codec type and output file type, to be used with an asset writer input.
	RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType AVVideoCodecType, outputFileType AVFileType) []AVMetadataItem
	// Returns a video settings dictionary appropriate for capturing video to a file with the specified codec and type.
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType AVVideoCodecType, outputFileType AVFileType) foundation.INSDictionary
	// Returns a dictionary of recommended output settings for writing the specified code, file type, and output URL.
	RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL(videoCodecType AVVideoCodecType, outputFileType AVFileType, outputFileURL foundation.INSURL) foundation.INSDictionary
	// Specifies the recommended settings for use with an AVAssetWriterInput.
	RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType AVFileType) foundation.INSDictionary

	// Topic: Retrieving supported video types

	// The video pixel formats the output supports.
	AvailableVideoPixelFormatTypes() uint32
	SetAvailableVideoPixelFormatTypes(value uint32)
	// The video codecs that the output supports.
	AvailableVideoCodecTypes() []string
	// The video codecs that the output supports for writing video to the output file.
	AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType AVFileType) []string

	// Topic: Receiving captured video data

	// Sets the sample buffer delegate and the queue for invoking callbacks.
	SetSampleBufferDelegateQueue(sampleBufferDelegate AVCaptureVideoDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue)
	// The capture object’s delegate.
	SampleBufferDelegate() AVCaptureVideoDataOutputSampleBufferDelegate
	// The queue on which the system invokes delegate callbacks.
	SampleBufferCallbackQueue() dispatch.Queue

	// The video pixel formats the output supports.
	AvailableVideoCVPixelFormatTypes() []foundation.NSNumber
}

// Init initializes the instance.
func (c AVCaptureVideoDataOutput) Init() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureVideoDataOutput) Autorelease() AVCaptureVideoDataOutput {
	rv := objc.Send[AVCaptureVideoDataOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureVideoDataOutput creates a new AVCaptureVideoDataOutput instance.
func NewAVCaptureVideoDataOutput() AVCaptureVideoDataOutput {
	class := getAVCaptureVideoDataOutputClass()
	rv := objc.Send[AVCaptureVideoDataOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Recommends movie-level metadata for a particular video codec type and
// output file type, to be used with an asset writer input.
//
// videoCodecType: The desired [AVVideoCodecKey] to be used for compression (see [Video
// settings]).
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// outputFileType: Specifies the UTI of the file type to be written (see [AVFileType]).
//
// # Return Value
// 
// A fully populated array of [AVMetadataItem] objects compatible with
// [AVAssetWriter].
//
// # Discussion
// 
// The value of this property is an array of [AVMetadataItem] objects
// representing the collection of top-level metadata to be written in each
// output file. This array is suitable to use as the [Metadata] property
// before you have called [startWriting()]. For more details see
// [startWriting()].
// 
// The `videoCodecType` string you provide must be present in
// [AvailableVideoCodecTypesForAssetWriterWithOutputFileType] array, or an
// [NSInvalidArgumentException] is thrown.
// 
// For clients writing files using a ProRes Raw codec type, white balance must
// be locked (call
// [SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler])
// before querying this property, or an [NSIvalidArgumentException] is thrown.
//
// [startWriting()]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startWriting()
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMovieMetadata(forVideoCodecType:assetWriterOutputFileType:)
func (c AVCaptureVideoDataOutput) RecommendedMovieMetadataForVideoCodecTypeAssetWriterOutputFileType(videoCodecType AVVideoCodecType, outputFileType AVFileType) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recommendedMovieMetadataForVideoCodecType:assetWriterOutputFileType:"), objc.String(string(videoCodecType)), objc.String(string(outputFileType)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
// Returns a video settings dictionary appropriate for capturing video to a
// file with the specified codec and type.
//
// videoCodecType: The video codec type to write.
//
// outputFileType: The Uniform Type Identifier of the file type to write. See `File Format
// UTIs` for supported types.
//
// # Return Value
// 
// A fully populated dictionary of keys and values that are compatible with
// [AVAssetWriter].
//
// # Discussion
// 
// This dictionary contains keys and values described in [Video settings] and
// is suitable for use when creating an [AVAssetWriterInput] with the
// [InitWithMediaTypeOutputSettings] initializer.
// 
// For QuickTime movie and ISO file types, the recommended video settings
// produce output comparable to that of [AVCaptureMovieFileOutput].
// 
// Note that the dictionary of settings is dependent on the current
// configuration of the output’s [AVCaptureSession] and its inputs. The
// settings dictionary may change if the session’s configuration changes. As
// such, configure your session first, then query the recommended video
// settings.
//
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettings(forVideoCodecType:assetWriterOutputFileType:)
func (c AVCaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileType(videoCodecType AVVideoCodecType, outputFileType AVFileType) foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:"), objc.String(string(videoCodecType)), objc.String(string(outputFileType)))
	return foundation.NSDictionaryFromID(rv)
}
// Returns a dictionary of recommended output settings for writing the
// specified code, file type, and output URL.
//
// videoCodecType: The type of video codec to use.
//
// outputFileType: The type of output file to write.
//
// outputFileURL: The URL of the output file to write.
//
// # Return Value
// 
// A fully populated output settings dictionary suitable for configuring an
// asset writer input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettings(forVideoCodecType:assetWriterOutputFileType:outputFileURL:)
func (c AVCaptureVideoDataOutput) RecommendedVideoSettingsForVideoCodecTypeAssetWriterOutputFileTypeOutputFileURL(videoCodecType AVVideoCodecType, outputFileType AVFileType, outputFileURL foundation.INSURL) foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recommendedVideoSettingsForVideoCodecType:assetWriterOutputFileType:outputFileURL:"), objc.String(string(videoCodecType)), objc.String(string(outputFileType)), outputFileURL)
	return foundation.NSDictionaryFromID(rv)
}
// Specifies the recommended settings for use with an AVAssetWriterInput.
//
// outputFileType: The Uniform Type Identifier of the file type to write. See `File Format
// UTIs` for supported types.
//
// # Return Value
// 
// A fully populated dictionary of keys and values that are compatible with
// [AVAssetWriter].
//
// # Discussion
// 
// This dictionary contains keys and values described in [Video settings] and
// is suitable for use when creating an [AVAssetWriterInput] with the
// [InitWithMediaTypeOutputSettings] initializer. For QuickTime movie and ISO
// file types, the recommended video settings produce output comparable to
// that of [AVCaptureMovieFileOutput].
// 
// Note that the dictionary of settings is dependent on the current
// configuration of the output’s [AVCaptureSession] and its inputs. The
// settings dictionary may change if the session’s configuration changes. As
// such, configure your session first, then query the recommended video
// settings.
//
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedVideoSettingsForAssetWriter(writingTo:)
func (c AVCaptureVideoDataOutput) RecommendedVideoSettingsForAssetWriterWithOutputFileType(outputFileType AVFileType) foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recommendedVideoSettingsForAssetWriterWithOutputFileType:"), objc.String(string(outputFileType)))
	return foundation.NSDictionaryFromID(rv)
}
// The video codecs that the output supports for writing video to the output
// file.
//
// outputFileType: The UTI of the output file type.
//
// # Return Value
// 
// An array of video codecs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCodecTypesForAssetWriter(writingTo:)
func (c AVCaptureVideoDataOutput) AvailableVideoCodecTypesForAssetWriterWithOutputFileType(outputFileType AVFileType) []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableVideoCodecTypesForAssetWriterWithOutputFileType:"), objc.String(string(outputFileType)))
	return objc.ConvertSliceToStrings(rv)
}
// Sets the sample buffer delegate and the queue for invoking callbacks.
//
// sampleBufferDelegate: An object conforming to the [AVCaptureVideoDataOutputSampleBufferDelegate]
// protocol that will receive sample buffers after they are captured.
//
// sampleBufferCallbackQueue: The queue on which callbacks should be invoked. You must use a serial
// dispatch queue, to guarantee that video frames will be delivered in order.
// 
// The sampleBufferCallbackQueue parameter may not be [NULL], except when
// setting the `sampleBufferDelegate` to `nil`.
//
// # Discussion
// 
// When a new video sample buffer is captured, it is sent to the sample buffer
// delegate using [CaptureOutputDidOutputSampleBufferFromConnection]. All
// delegate methods are invoked on the specified dispatch queue.
// 
// If the queue is blocked when new frames are captured, those frames will be
// automatically dropped at a time determined by the value of the
// [AlwaysDiscardsLateVideoFrames] property. This allows you to process
// existing frames on the same queue without having to manage the potential
// memory usage increases that would otherwise occur when that processing is
// unable to keep up with the rate of incoming frames.
// 
// If your frame processing is consistently unable to keep up with the rate of
// incoming frames, you should consider using the [MinFrameDuration] property,
// which will generally yield better performance characteristics and more
// consistent frame rates than frame dropping alone.
// 
// If you need to minimize the chances of frames being dropped, you should
// specify a queue on which a sufficiently small amount of processing is being
// done outside of receiving sample buffers. However, if you migrate extra
// processing to another queue, you are responsible for ensuring that memory
// usage does not grow without bound from frames that have not been processed.
// 
// # Special considerations
// 
// This method uses [dispatch_retain] and [dispatch_release] to manage the
// queue.
//
// [dispatch_release]: https://developer.apple.com/documentation/Dispatch/dispatch_release
// [dispatch_retain]: https://developer.apple.com/documentation/Dispatch/dispatch_retain
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/setSampleBufferDelegate(_:queue:)
func (c AVCaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate AVCaptureVideoDataOutputSampleBufferDelegate, sampleBufferCallbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, uintptr(sampleBufferCallbackQueue.Handle()))
}

// A dictionary that contains the compression settings for the output.
//
// # Discussion
// 
// To receive samples in their device-native format, set this value to an
// empty dictionary:
// 
// To receive samples in a default uncompressed format, set this value to
// `nil`. Then you can query this value to receive a dictionary of the
// settings the session uses.
// 
// In iOS versions prior to iOS 16, the only key supported is
// [kCVPixelBufferPixelFormatTypeKey]. In iOS 16 and later, the supported keys
// include the following:
// 
// - For compressed video output, only use [AVVideoPixelAspectRatioKey],
// [AVVideoCleanApertureKey], [AVVideoScalingModeKey],
// [AVVideoColorPropertiesKey], and [AVVideoAllowWideColorKey]. - For
// uncompressed video output, you can also use
// [kCVPixelBufferPixelFormatTypeKey], [KCVPixelBufferWidthKey], and
// [KCVPixelBufferHeightKey], in addition to the compressed video output keys.
// 
// You can use [availableVideoPixelFormatTypes] and [AvailableVideoCodecTypes]
// to get a list of the supported pixel formats and video codecs,
// respectively. The width and height need to match the [VideoOrientation]
// specified in the output’s [AVCaptureConnection], otherwise the system
// throws an [InvalidArgumentException]. The aspect ratio of the width and
// height also need to match the aspect ratio of the source’s
// [ActiveFormat], corrected for the connection’s [VideoOrientation],
// otherwise the system throws an [InvalidArgumentException]. If the width or
// height exceeds the source’s `activeFormat`‘s width or height, the
// system throws an [InvalidArgumentException]. Don’t change the width and
// height if [DeliversPreviewSizedOutputBuffers] is [true], otherwise the
// system throws an [InvalidArgumentException].
//
// [AVVideoAllowWideColorKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoAllowWideColorKey
// [AVVideoCleanApertureKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureKey
// [AVVideoColorPropertiesKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPropertiesKey
// [AVVideoPixelAspectRatioKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoPixelAspectRatioKey
// [AVVideoScalingModeKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeKey
// [availableVideoPixelFormatTypes]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoPixelFormatTypes
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/videoSettings
func (c AVCaptureVideoDataOutput) VideoSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c AVCaptureVideoDataOutput) SetVideoSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoSettings:"), value)
}
// Indicates whether to drop video frames if they arrive late.
//
// # Discussion
// 
// If this property is [true], the output immediately discard frames that
// captured while the dispatch queue handling existing frames blocks in the
// [CaptureOutputDidOutputSampleBufferFromConnection] delegate method. When
// set to [false], the output gives delegates more time to process old frames
// before it discards new frames, but application memory usage may increase
// significantly as a result.
// 
// The default is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/alwaysDiscardsLateVideoFrames
func (c AVCaptureVideoDataOutput) AlwaysDiscardsLateVideoFrames() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("alwaysDiscardsLateVideoFrames"))
	return rv
}
func (c AVCaptureVideoDataOutput) SetAlwaysDiscardsLateVideoFrames(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlwaysDiscardsLateVideoFrames:"), value)
}
// Indicates whether the receiver should preserve dynamic HDR metadata as an
// attachment on the output sample buffer’s underlying pixel buffer.
//
// # Discussion
// 
// Set this property to `true` if you wish to use [AVCaptureVideoDataOutput]
// with [AVAssetWriter] to record HDR movies. You must also set
// `kVTCompressionPropertyKey_PreserveDynamicHDRMetadata` to `true` in the
// compression settings you pass to your [AVAssetWriterInput]. These
// compression settings are represented under the
// [AVVideoCompressionPropertiesKey] sub-dictionary of your top-level
// AVVideoSettings (see [Video settings]). When you set this key to `true`,
// performance improves, as the encoder is able to skip HDR metadata
// calculation for every frame. The default value is `false`.
//
// [AVVideoCompressionPropertiesKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCompressionPropertiesKey
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preservesDynamicHDRMetadata
func (c AVCaptureVideoDataOutput) PreservesDynamicHDRMetadata() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("preservesDynamicHDRMetadata"))
	return rv
}
func (c AVCaptureVideoDataOutput) SetPreservesDynamicHDRMetadata(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreservesDynamicHDRMetadata:"), value)
}
// Indicates the recommended media timescale for the video track.
//
// # Return Value
// 
// The recommended media timescale based on the active capture session’s
// inputs. It is never less than 600. It may or may not be a multiple of 600.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/recommendedMediaTimeScaleForAssetWriter
func (c AVCaptureVideoDataOutput) RecommendedMediaTimeScaleForAssetWriter() int32 {
	rv := objc.Send[int32](c.ID, objc.Sel("recommendedMediaTimeScaleForAssetWriter"))
	return rv
}
// The video pixel formats the output supports.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutput/availablevideopixelformattypes
func (c AVCaptureVideoDataOutput) AvailableVideoPixelFormatTypes() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("availableVideoPixelFormatTypes"))
	return rv
}
func (c AVCaptureVideoDataOutput) SetAvailableVideoPixelFormatTypes(value uint32) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableVideoPixelFormatTypes:"), value)
}
// The video codecs that the output supports.
//
// # Discussion
// 
// The value contains an array of video codecs that the output supports.
// Specify the codec it uses by setting a supported value for the
// [AVVideoCodecKey] entry in its [VideoSettings] dictionary. The first format
// in the returned list is the most efficient output format.
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCodecTypes
func (c AVCaptureVideoDataOutput) AvailableVideoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableVideoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// The capture object’s delegate.
//
// # Discussion
// 
// The delegate receives sample buffers after they are captured.
// 
// You set the delegate using [SetSampleBufferDelegateQueue].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferDelegate
func (c AVCaptureVideoDataOutput) SampleBufferDelegate() AVCaptureVideoDataOutputSampleBufferDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sampleBufferDelegate"))
	return AVCaptureVideoDataOutputSampleBufferDelegateObjectFromID(rv)
}
// The queue on which the system invokes delegate callbacks.
//
// # Discussion
// 
// You set the queue using [SetSampleBufferDelegateQueue].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/sampleBufferCallbackQueue
func (c AVCaptureVideoDataOutput) SampleBufferCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("sampleBufferCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}
// The video pixel formats the output supports.
//
// # Discussion
// 
// This value contains an array of video formats, in unspecified order, that
// the output supports. You can set the format by specifying it as the
// [kCVPixelBufferPixelFormatTypeKey] entry in the output’s [VideoSettings]
// dictionary.
//
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/availableVideoCVPixelFormatTypes
func (c AVCaptureVideoDataOutput) AvailableVideoCVPixelFormatTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableVideoCVPixelFormatTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

