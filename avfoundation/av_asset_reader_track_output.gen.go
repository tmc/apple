// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
// this means that [AVAssetReaderTrackOutput.AVFormatIDKey] must be [AVAssetReaderTrackOutput.KAudioFormatLinearPCM]. For video
// output settings, this means that the dictionary must contain values for
// uncompressed video output, as defined in `Video Settings`. A track output
// doesn’t support the [AVAssetReaderTrackOutput.AVSampleRateConverterAudioQualityKey] audio setting
// key or the following video settings keys: [AVAssetReaderTrackOutput.AVVideoCleanApertureKey],
// [AVAssetReaderTrackOutput.AVVideoPixelAspectRatioKey], and [AVAssetReaderTrackOutput.AVVideoScalingModeKey].
//
// When constructing video output settings, the choice of pixel format affects
// the performance and quality of the decompression. For optimal performance
// when decompressing video, the requested pixel format should be one that the
// decoder supports natively to avoid unnecessary conversions. Below are some
// recommendations:
//
// - For H.264, use [KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange] or
// [KCVPixelFormatType_420YpCbCr8BiPlanarFullRange] when you know the video is
// full range. - In iOS, use [KCVPixelFormatType_420YpCbCr8BiPlanarFullRange]
// for JPEG output. - In macOS, [KCVPixelFormatType_422YpCbCr8] is the
// preferred pixel format for video and generally provides the best
// performance when decoding. If you need to work in the RGB domain, use
// [KCVPixelFormatType_32BGRA] in iOS, and [KCVPixelFormatType_32ARGB] in
// macOS. - ProRes-encoded media can contain up to 12 bits per channel. For
// ProRes-encoded sources that you wish to preserve more than 8 bits per
// channel during decompression, use one of the following pixel formats:
// [KCVPixelFormatType_4444AYpCbCr16], [KCVPixelFormatType_422YpCbCr16],
// [KCVPixelFormatType_422YpCbCr10], or [KCVPixelFormatType_64ARGB].
// [AVAssetReader] doesn’t support scaling with any of these high-bit-depth
// pixel formats. If you use the above pixel formats, don’t specify
// [AVAssetReaderTrackOutput.KCVPixelBufferWidthKey] or [AVAssetReaderTrackOutput.KCVPixelBufferHeightKey] in the
// [AVAssetReaderTrackOutput.OutputSettings] dictionary. Only ProRes encoders support these pixel
// formats. - ProRes 4444-encoded media can contain a mathematically lossless
// alpha channel. To preserve the alpha channel during decompression, use a
// pixel format with an alpha component such as
// [KCVPixelFormatType_4444AYpCbCr16] or [KCVPixelFormatType_64ARGB]. To test
// whether your source contains an alpha channel, check that the track’s
// format description has a [KCMFormatDescriptionExtension_Depth] key with a
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

	// An integer value that represents the format of the audio data.
	AVFormatIDKey() string
	// An integer value that represents the audio quality for conversion.
	AVSampleRateConverterAudioQualityKey() string
	// A key that defines the region within the video dimension displayed during playback.
	AVVideoCleanApertureKey() string
	// A key to access the video’s pixel aspect ratio.
	AVVideoPixelAspectRatioKey() string
	// A key to retrieve the video scaling mode from a dictionary.
	AVVideoScalingModeKey() string
	// A key that specifies the linear PCM codec, and uses the standard flags.
	KAudioFormatLinearPCM() objectivec.IObject
	SetKAudioFormatLinearPCM(value objectivec.IObject)
	KCMFormatDescriptionExtension_Depth() corefoundation.CFStringRef
	// A key to the height of the pixel buffer.
	KCVPixelBufferHeightKey() corefoundation.CFStringRef
	// A key to the width of the pixel buffer.
	KCVPixelBufferWidthKey() corefoundation.CFStringRef
	// 32-bit ARGB.
	KCVPixelFormatType_32ARGB() uint32
	SetKCVPixelFormatType_32ARGB(value uint32)
	// 32-bit BGRA.
	KCVPixelFormatType_32BGRA() uint32
	SetKCVPixelFormatType_32BGRA(value uint32)
	// Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255] chroma=[1,255]).  `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
	KCVPixelFormatType_420YpCbCr8BiPlanarFullRange() uint32
	SetKCVPixelFormatType_420YpCbCr8BiPlanarFullRange(value uint32)
	// Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235] chroma=[16,240]).  `baseAddr` points to a big-endian `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
	KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange() uint32
	SetKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange(value uint32)
	// Component Y’CbCr 10-bit 4:2:2.
	KCVPixelFormatType_422YpCbCr10() uint32
	SetKCVPixelFormatType_422YpCbCr10(value uint32)
	// Component Y’CbCr 10,12,14,16-bit 4:2:2.
	KCVPixelFormatType_422YpCbCr16() uint32
	SetKCVPixelFormatType_422YpCbCr16(value uint32)
	// Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
	KCVPixelFormatType_422YpCbCr8() uint32
	SetKCVPixelFormatType_422YpCbCr8(value uint32)
	// Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha, video range Y’CbCr, 16-bit little-endian samples.
	KCVPixelFormatType_4444AYpCbCr16() uint32
	SetKCVPixelFormatType_4444AYpCbCr16(value uint32)
	// 64-bit ARGB, 16-bit big-endian samples.
	KCVPixelFormatType_64ARGB() uint32
	SetKCVPixelFormatType_64ARGB(value uint32)
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

// An integer value that represents the format of the audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVFormatIDKey
func (a AVAssetReaderTrackOutput) AVFormatIDKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVFormatIDKey"))
	return foundation.NSStringFromID(rv).String()
}

// An integer value that represents the audio quality for conversion.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVSampleRateConverterAudioQualityKey
func (a AVAssetReaderTrackOutput) AVSampleRateConverterAudioQualityKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVSampleRateConverterAudioQualityKey"))
	return foundation.NSStringFromID(rv).String()
}

// A key that defines the region within the video dimension displayed during
// playback.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideocleanaperturekey
func (a AVAssetReaderTrackOutput) AVVideoCleanApertureKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVVideoCleanApertureKey"))
	return foundation.NSStringFromID(rv).String()
}

// A key to access the video’s pixel aspect ratio.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideopixelaspectratiokey
func (a AVAssetReaderTrackOutput) AVVideoPixelAspectRatioKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVVideoPixelAspectRatioKey"))
	return foundation.NSStringFromID(rv).String()
}

// A key to retrieve the video scaling mode from a dictionary.
//
// See: https://developer.apple.com/documentation/avfoundation/avvideoscalingmodekey
func (a AVAssetReaderTrackOutput) AVVideoScalingModeKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVVideoScalingModeKey"))
	return foundation.NSStringFromID(rv).String()
}

// A key that specifies the linear PCM codec, and uses the standard flags.
//
// See: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatLinearPCM
func (a AVAssetReaderTrackOutput) KAudioFormatLinearPCM() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("kAudioFormatLinearPCM"))
	return objectivec.Object{ID: rv}
}
func (a AVAssetReaderTrackOutput) SetKAudioFormatLinearPCM(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setKAudioFormatLinearPCM:"), value)
}

// See: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_Depth
func (a AVAssetReaderTrackOutput) KCMFormatDescriptionExtension_Depth() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](a.ID, objc.Sel("kCMFormatDescriptionExtension_Depth"))
	return corefoundation.CFStringRef(rv)
}

// A key to the height of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferHeightKey
func (a AVAssetReaderTrackOutput) KCVPixelBufferHeightKey() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](a.ID, objc.Sel("kCVPixelBufferHeightKey"))
	return corefoundation.CFStringRef(rv)
}

// A key to the width of the pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferWidthKey
func (a AVAssetReaderTrackOutput) KCVPixelBufferWidthKey() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](a.ID, objc.Sel("kCVPixelBufferWidthKey"))
	return corefoundation.CFStringRef(rv)
}

// 32-bit ARGB.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32ARGB
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_32ARGB() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_32ARGB"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_32ARGB(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_32ARGB:"), value)
}

// 32-bit BGRA.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_32BGRA() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_32BGRA"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_32BGRA(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_32BGRA:"), value)
}

// Bi-Planar Component Y’CbCr 8-bit 4:2:0, full-range (luma=[0,255]
// chroma=[1,255]). `baseAddr` points to a big-endian
// `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarFullRange
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_420YpCbCr8BiPlanarFullRange() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_420YpCbCr8BiPlanarFullRange"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_420YpCbCr8BiPlanarFullRange(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_420YpCbCr8BiPlanarFullRange:"), value)
}

// Bi-Planar Component Y’CbCr 8-bit 4:2:0, video-range (luma=[16,235]
// chroma=[16,240]). `baseAddr` points to a big-endian
// `CVPlanarPixelBufferInfo_YCbCrBiPlanar` struct.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_420YpCbCr8BiPlanarVideoRange() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_420YpCbCr8BiPlanarVideoRange"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_420YpCbCr8BiPlanarVideoRange:"), value)
}

// Component Y’CbCr 10-bit 4:2:2.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr10
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr10() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_422YpCbCr10"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr10(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr10:"), value)
}

// Component Y’CbCr 10,12,14,16-bit 4:2:2.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr16
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr16() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_422YpCbCr16"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr16(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr16:"), value)
}

// Component Y’CbCr 8-bit 4:2:2, ordered Cb Y’0 Cr Y’1.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_422YpCbCr8
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_422YpCbCr8() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_422YpCbCr8"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_422YpCbCr8(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_422YpCbCr8:"), value)
}

// Component Y’CbCrA 16-bit 4:4:4:4, ordered A Y’ Cb Cr, full range alpha,
// video range Y’CbCr, 16-bit little-endian samples.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_4444AYpCbCr16
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_4444AYpCbCr16() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_4444AYpCbCr16"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_4444AYpCbCr16(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_4444AYpCbCr16:"), value)
}

// 64-bit ARGB, 16-bit big-endian samples.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_64ARGB
func (a AVAssetReaderTrackOutput) KCVPixelFormatType_64ARGB() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("kCVPixelFormatType_64ARGB"))
	return rv
}
func (a AVAssetReaderTrackOutput) SetKCVPixelFormatType_64ARGB(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setKCVPixelFormatType_64ARGB:"), value)
}
