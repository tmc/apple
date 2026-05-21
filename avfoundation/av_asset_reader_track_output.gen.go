// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVAssetReaderTrackOutput] class.
var (
	_AVAssetReaderTrackOutputClass     AVAssetReaderTrackOutputClass
	_AVAssetReaderTrackOutputClassOnce sync.Once
)

func getAVAssetReaderTrackOutputClass() AVAssetReaderTrackOutputClass {
	_AVAssetReaderTrackOutputClassOnce.Do(func() {
		_AVAssetReaderTrackOutputClass = AVAssetReaderTrackOutputClass{class: objc.GetClass("AVAssetReaderTrackOutput")}
	})
	return _AVAssetReaderTrackOutputClass
}

// GetAVAssetReaderTrackOutputClass returns the class object for AVAssetReaderTrackOutput.
func GetAVAssetReaderTrackOutputClass() AVAssetReaderTrackOutputClass {
	return getAVAssetReaderTrackOutputClass()
}

type AVAssetReaderTrackOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetReaderTrackOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetReaderTrackOutputClass) Alloc() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that reads media data from a single track of an asset.
//
// # Overview
//
// Read the media data of an asset track by adding a track output to an asset
// reader. You can read the media samples in their stored format, or you can
// convert them to an alternative format.
//
// A track output produces uncompressed output. For audio output settings,
// this means that [AVFormatIDKey] must be [kAudioFormatLinearPCM]. For video
// output settings, this means that the dictionary must contain values for
// uncompressed video output, as defined in `Video Settings`. A track output
// doesn’t support the [AVSampleRateConverterAudioQualityKey] audio setting
// key or the following video settings keys: [AVVideoCleanApertureKey],
// [AVVideoPixelAspectRatioKey], and [AVVideoScalingModeKey].
//
// When constructing video output settings, the choice of pixel format affects
// the performance and quality of the decompression. For optimal performance
// when decompressing video, the requested pixel format should be one that the
// decoder supports natively to avoid unnecessary conversions. Below are some
// recommendations:
//
// - For H.264, use [kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange] or
// [kCVPixelFormatType_420YpCbCr8BiPlanarFullRange] when you know the video is
// full range. - In iOS, use [kCVPixelFormatType_420YpCbCr8BiPlanarFullRange]
// for JPEG output. - In macOS, [kCVPixelFormatType_422YpCbCr8] is the
// preferred pixel format for video and generally provides the best
// performance when decoding. If you need to work in the RGB domain, use
// [kCVPixelFormatType_32BGRA] in iOS, and [kCVPixelFormatType_32ARGB] in
// macOS. - ProRes-encoded media can contain up to 12 bits per channel. For
// ProRes-encoded sources that you wish to preserve more than 8 bits per
// channel during decompression, use one of the following pixel formats:
// [kCVPixelFormatType_4444AYpCbCr16], [kCVPixelFormatType_422YpCbCr16],
// [kCVPixelFormatType_422YpCbCr10], or [kCVPixelFormatType_64ARGB].
// [AVAssetReader] doesn’t support scaling with any of these high-bit-depth
// pixel formats. If you use the above pixel formats, don’t specify
// [kCVPixelBufferWidthKey] or [kCVPixelBufferHeightKey] in the
// [AVAssetReaderTrackOutput.OutputSettings] dictionary. Only ProRes encoders
// support these pixel formats. - ProRes 4444-encoded media can contain a
// mathematically lossless alpha channel. To preserve the alpha channel during
// decompression, use a pixel format with an alpha component such as
// [kCVPixelFormatType_4444AYpCbCr16] or [kCVPixelFormatType_64ARGB]. To test
// whether your source contains an alpha channel, check that the track’s
// format description has a [kCMFormatDescriptionExtension_Depth] key with a
// value of `32`.
//
// # Creating a track output
//
//   - [AVAssetReaderTrackOutput.InitWithTrackOutputSettings]: Creates an object that reads media data from an asset track.
//
// # Configuring audio settings
//
//   - [AVAssetReaderTrackOutput.AudioTimePitchAlgorithm]: The processing algorithm to use for scaled audio edits.
//   - [AVAssetReaderTrackOutput.SetAudioTimePitchAlgorithm]
//
// # Inspecting an output
//
//   - [AVAssetReaderTrackOutput.OutputSettings]: The output settings for this track output.
//   - [AVAssetReaderTrackOutput.Track]: The track from which the output reads sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput
//
// [AVFormatIDKey]: https://developer.apple.com/documentation/AVFAudio/AVFormatIDKey
// [AVSampleRateConverterAudioQualityKey]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAudioQualityKey
// [AVVideoCleanApertureKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureKey
// [AVVideoPixelAspectRatioKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoPixelAspectRatioKey
// [AVVideoScalingModeKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeKey
// [kAudioFormatLinearPCM]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatLinearPCM
// [kCMFormatDescriptionExtension_Depth]: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Depth
// [kCVPixelBufferHeightKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferHeightKey
// [kCVPixelBufferWidthKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferWidthKey
// [kCVPixelFormatType_32ARGB]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
// [kCVPixelFormatType_32BGRA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
// [kCVPixelFormatType_420YpCbCr8BiPlanarFullRange]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarFullRange
// [kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange
// [kCVPixelFormatType_422YpCbCr10]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr10
// [kCVPixelFormatType_422YpCbCr16]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr16
// [kCVPixelFormatType_422YpCbCr8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
// [kCVPixelFormatType_4444AYpCbCr16]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_4444AYpCbCr16
// [kCVPixelFormatType_64ARGB]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64ARGB
type AVAssetReaderTrackOutput struct {
	AVAssetReaderOutput
}

// AVAssetReaderTrackOutputFromID constructs a [AVAssetReaderTrackOutput] from an objc.ID.
//
// An object that reads media data from a single track of an asset.
func AVAssetReaderTrackOutputFromID(id objc.ID) AVAssetReaderTrackOutput {
	return AVAssetReaderTrackOutput{AVAssetReaderOutput: AVAssetReaderOutputFromID(id)}
}

// NOTE: AVAssetReaderTrackOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetReaderTrackOutput] class.
//
// # Creating a track output
//
//   - [IAVAssetReaderTrackOutput.InitWithTrackOutputSettings]: Creates an object that reads media data from an asset track.
//
// # Configuring audio settings
//
//   - [IAVAssetReaderTrackOutput.AudioTimePitchAlgorithm]: The processing algorithm to use for scaled audio edits.
//   - [IAVAssetReaderTrackOutput.SetAudioTimePitchAlgorithm]
//
// # Inspecting an output
//
//   - [IAVAssetReaderTrackOutput.OutputSettings]: The output settings for this track output.
//   - [IAVAssetReaderTrackOutput.Track]: The track from which the output reads sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput
type IAVAssetReaderTrackOutput interface {
	IAVAssetReaderOutput

	// Topic: Creating a track output

	// Creates an object that reads media data from an asset track.
	InitWithTrackOutputSettings(track IAVAssetTrack, outputSettings foundation.INSDictionary) AVAssetReaderTrackOutput

	// Topic: Configuring audio settings

	// The processing algorithm to use for scaled audio edits.
	AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm)

	// Topic: Inspecting an output

	// The output settings for this track output.
	OutputSettings() foundation.INSDictionary
	// The track from which the output reads sample buffers.
	Track() IAVAssetTrack
}

// Init initializes the instance.
func (a AVAssetReaderTrackOutput) Init() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetReaderTrackOutput) Autorelease() AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderTrackOutput creates a new AVAssetReaderTrackOutput instance.
func NewAVAssetReaderTrackOutput() AVAssetReaderTrackOutput {
	class := getAVAssetReaderTrackOutputClass()
	rv := objc.Send[AVAssetReaderTrackOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that reads media data from an asset track.
//
// track: The track from which to read media samples.
//
// outputSettings: A dictionary of settings to use for sample output. Specify `nil` to receive
// samples in their storage format.
//
// You use keys and values from [Audio settings], [Video settings], or
// [CVPixelBuffer], depending on the media type and the output format you
// require.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/init(track:outputSettings:)
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
func NewAssetReaderTrackOutputWithTrackOutputSettings(track IAVAssetTrack, outputSettings foundation.INSDictionary) AVAssetReaderTrackOutput {
	instance := getAVAssetReaderTrackOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTrack:outputSettings:"), track, outputSettings)
	return AVAssetReaderTrackOutputFromID(rv)
}

// Creates an object that reads media data from an asset track.
//
// track: The track from which to read media samples.
//
// outputSettings: A dictionary of settings to use for sample output. Specify `nil` to receive
// samples in their storage format.
//
// You use keys and values from [Audio settings], [Video settings], or
// [CVPixelBuffer], depending on the media type and the output format you
// require.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/init(track:outputSettings:)
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
func (a AVAssetReaderTrackOutput) InitWithTrackOutputSettings(track IAVAssetTrack, outputSettings foundation.INSDictionary) AVAssetReaderTrackOutput {
	rv := objc.Send[AVAssetReaderTrackOutput](a.ID, objc.Sel("initWithTrack:outputSettings:"), track, outputSettings)
	return rv
}

// Returns a new object that reads media data from an asset track.
//
// track: The track from which to read media samples.
//
// outputSettings: A dictionary of settings to use for sample output. Specify `nil` to receive
// samples in their storage format.
//
// You use keys and values from [Audio settings], [Video settings], or
// [CVPixelBuffer], depending on the media type and the output format you
// require.
//
// # Return Value
//
// A new asset reader, or `nil` if initialization fails.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/assetReaderTrackOutputWithTrack:outputSettings:
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
func (_AVAssetReaderTrackOutputClass AVAssetReaderTrackOutputClass) AssetReaderTrackOutputWithTrackOutputSettings(track IAVAssetTrack, outputSettings foundation.INSDictionary) AVAssetReaderTrackOutput {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetReaderTrackOutputClass.class), objc.Sel("assetReaderTrackOutputWithTrack:outputSettings:"), track, outputSettings)
	return AVAssetReaderTrackOutputFromID(rv)
}

// The processing algorithm to use for scaled audio edits.
//
// # Discussion
//
// See [Time pitch algorithm settings] for possible values. The system throws
// an exception if you set this property to a value other than one of the
// defined constants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/audioTimePitchAlgorithm
//
// [Time pitch algorithm settings]: https://developer.apple.com/documentation/AVFoundation/time-pitch-algorithm-settings
func (a AVAssetReaderTrackOutput) AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioTimePitchAlgorithm"))
	return AVAudioTimePitchAlgorithm(foundation.NSStringFromID(rv).String())
}
func (a AVAssetReaderTrackOutput) SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioTimePitchAlgorithm:"), objc.String(string(value)))
}

// The output settings for this track output.
//
// # Discussion
//
// The value is a dictionary that contains values for audio and video settings
// keys. A value of `nil` indicates that the track output vends samples in
// their original format as stored in the target track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/outputSettings
func (a AVAssetReaderTrackOutput) OutputSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The track from which the output reads sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderTrackOutput/track
func (a AVAssetReaderTrackOutput) Track() IAVAssetTrack {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("track"))
	return AVAssetTrackFromID(objc.ID(rv))
}
