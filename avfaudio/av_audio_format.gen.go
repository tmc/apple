// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioFormat] class.
var (
	_AVAudioFormatClass     AVAudioFormatClass
	_AVAudioFormatClassOnce sync.Once
)

func getAVAudioFormatClass() AVAudioFormatClass {
	_AVAudioFormatClassOnce.Do(func() {
		_AVAudioFormatClass = AVAudioFormatClass{class: objc.GetClass("AVAudioFormat")}
	})
	return _AVAudioFormatClass
}

// GetAVAudioFormatClass returns the class object for AVAudioFormat.
func GetAVAudioFormatClass() AVAudioFormatClass {
	return getAVAudioFormatClass()
}

type AVAudioFormatClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioFormatClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioFormatClass) Alloc() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the representation of an audio format.
//
// # Overview
//
// The [AVAudioFormat] class wraps Core Audio’s
// [AudioStreamBasicDescription], and includes convenience initializers and
// accessors for common formats, including Core Audio’s standard
// deinterleaved 32-bit floating point format.
//
// Instances of this class are immutable.
//
// # Creating a New Audio Format Representation
//
//   - [AVAudioFormat.InitStandardFormatWithSampleRateChannelLayout]: Creates an audio format instance as a deinterleaved float with the specified sample rate and channel layout.
//   - [AVAudioFormat.InitStandardFormatWithSampleRateChannels]: Creates an audio format instance with the specified sample rate and channel count.
//   - [AVAudioFormat.InitWithCommonFormatSampleRateChannelsInterleaved]: Creates an audio format instance.
//   - [AVAudioFormat.InitWithCommonFormatSampleRateInterleavedChannelLayout]: Creates an audio format instance with the specified audio format, sample rate, interleaved state, and channel layout.
//   - [AVAudioFormat.InitWithSettings]: Creates an audio format instance using the specified settings dictionary.
//   - [AVAudioFormat.InitWithStreamDescription]: Creates an audio format instance from a stream description.
//   - [AVAudioFormat.InitWithStreamDescriptionChannelLayout]: Creates an audio format instance from a stream description and channel layout.
//   - [AVAudioFormat.InitWithCMAudioFormatDescription]: Creates an audio format instance from a Core Media audio format description.
//
// # Getting the Audio Stream Description
//
//   - [AVAudioFormat.StreamDescription]: The audio format properties of a stream of audio data.
//
// # Getting Audio Format Values
//
//   - [AVAudioFormat.SampleRate]: The audio format sampling rate, in hertz.
//   - [AVAudioFormat.ChannelCount]: The number of channels of audio data.
//   - [AVAudioFormat.ChannelLayout]: The underlying audio channel layout.
//   - [AVAudioFormat.FormatDescription]: The audio format description to use with Core Media APIs.
//
// # Determining the Audio Format
//
//   - [AVAudioFormat.Interleaved]: A Boolean value that indicates whether the samples mix into one stream.
//   - [AVAudioFormat.Standard]: A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//   - [AVAudioFormat.CommonFormat]: The common format identifier instance.
//   - [AVAudioFormat.Settings]: A dictionary that represents the format as a dictionary using audio setting keys.
//   - [AVAudioFormat.MagicCookie]: An object that contains metadata that encoders and decoders require.
//   - [AVAudioFormat.SetMagicCookie]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
type AVAudioFormat struct {
	objectivec.Object
}

// AVAudioFormatFromID constructs a [AVAudioFormat] from an objc.ID.
//
// An object that describes the representation of an audio format.
func AVAudioFormatFromID(id objc.ID) AVAudioFormat {
	return AVAudioFormat{objectivec.Object{ID: id}}
}

// NOTE: AVAudioFormat adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioFormat] class.
//
// # Creating a New Audio Format Representation
//
//   - [IAVAudioFormat.InitStandardFormatWithSampleRateChannelLayout]: Creates an audio format instance as a deinterleaved float with the specified sample rate and channel layout.
//   - [IAVAudioFormat.InitStandardFormatWithSampleRateChannels]: Creates an audio format instance with the specified sample rate and channel count.
//   - [IAVAudioFormat.InitWithCommonFormatSampleRateChannelsInterleaved]: Creates an audio format instance.
//   - [IAVAudioFormat.InitWithCommonFormatSampleRateInterleavedChannelLayout]: Creates an audio format instance with the specified audio format, sample rate, interleaved state, and channel layout.
//   - [IAVAudioFormat.InitWithSettings]: Creates an audio format instance using the specified settings dictionary.
//   - [IAVAudioFormat.InitWithStreamDescription]: Creates an audio format instance from a stream description.
//   - [IAVAudioFormat.InitWithStreamDescriptionChannelLayout]: Creates an audio format instance from a stream description and channel layout.
//   - [IAVAudioFormat.InitWithCMAudioFormatDescription]: Creates an audio format instance from a Core Media audio format description.
//
// # Getting the Audio Stream Description
//
//   - [IAVAudioFormat.StreamDescription]: The audio format properties of a stream of audio data.
//
// # Getting Audio Format Values
//
//   - [IAVAudioFormat.SampleRate]: The audio format sampling rate, in hertz.
//   - [IAVAudioFormat.ChannelCount]: The number of channels of audio data.
//   - [IAVAudioFormat.ChannelLayout]: The underlying audio channel layout.
//   - [IAVAudioFormat.FormatDescription]: The audio format description to use with Core Media APIs.
//
// # Determining the Audio Format
//
//   - [IAVAudioFormat.Interleaved]: A Boolean value that indicates whether the samples mix into one stream.
//   - [IAVAudioFormat.Standard]: A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
//   - [IAVAudioFormat.CommonFormat]: The common format identifier instance.
//   - [IAVAudioFormat.Settings]: A dictionary that represents the format as a dictionary using audio setting keys.
//   - [IAVAudioFormat.MagicCookie]: An object that contains metadata that encoders and decoders require.
//   - [IAVAudioFormat.SetMagicCookie]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat
type IAVAudioFormat interface {
	objectivec.IObject

	// Topic: Creating a New Audio Format Representation

	// Creates an audio format instance as a deinterleaved float with the specified sample rate and channel layout.
	InitStandardFormatWithSampleRateChannelLayout(sampleRate float64, layout IAVAudioChannelLayout) AVAudioFormat
	// Creates an audio format instance with the specified sample rate and channel count.
	InitStandardFormatWithSampleRateChannels(sampleRate float64, channels AVAudioChannelCount) AVAudioFormat
	// Creates an audio format instance.
	InitWithCommonFormatSampleRateChannelsInterleaved(format AVAudioCommonFormat, sampleRate float64, channels AVAudioChannelCount, interleaved bool) AVAudioFormat
	// Creates an audio format instance with the specified audio format, sample rate, interleaved state, and channel layout.
	InitWithCommonFormatSampleRateInterleavedChannelLayout(format AVAudioCommonFormat, sampleRate float64, interleaved bool, layout IAVAudioChannelLayout) AVAudioFormat
	// Creates an audio format instance using the specified settings dictionary.
	InitWithSettings(settings foundation.INSDictionary) AVAudioFormat
	// Creates an audio format instance from a stream description.
	InitWithStreamDescription(asbd unsafe.Pointer) AVAudioFormat
	// Creates an audio format instance from a stream description and channel layout.
	InitWithStreamDescriptionChannelLayout(asbd unsafe.Pointer, layout IAVAudioChannelLayout) AVAudioFormat
	// Creates an audio format instance from a Core Media audio format description.
	InitWithCMAudioFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) AVAudioFormat

	// Topic: Getting the Audio Stream Description

	// The audio format properties of a stream of audio data.
	StreamDescription() unsafe.Pointer

	// Topic: Getting Audio Format Values

	// The audio format sampling rate, in hertz.
	SampleRate() float64
	// The number of channels of audio data.
	ChannelCount() AVAudioChannelCount
	// The underlying audio channel layout.
	ChannelLayout() IAVAudioChannelLayout
	// The audio format description to use with Core Media APIs.
	FormatDescription() coremedia.CMFormatDescriptionRef

	// Topic: Determining the Audio Format

	// A Boolean value that indicates whether the samples mix into one stream.
	Interleaved() bool
	// A Boolean value that indicates whether the format is in a deinterleaved native-endian float state.
	Standard() bool
	// The common format identifier instance.
	CommonFormat() AVAudioCommonFormat
	// A dictionary that represents the format as a dictionary using audio setting keys.
	Settings() foundation.INSDictionary
	// An object that contains metadata that encoders and decoders require.
	MagicCookie() foundation.INSData
	SetMagicCookie(value foundation.INSData)

	AVChannelLayoutKey() string
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a AVAudioFormat) Init() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioFormat) Autorelease() AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioFormat creates a new AVAudioFormat instance.
func NewAVAudioFormat() AVAudioFormat {
	class := getAVAudioFormatClass()
	rv := objc.Send[AVAudioFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an audio format instance as a deinterleaved float with the
// specified sample rate and channel layout.
//
// sampleRate: The sampling rate, in hertz.
//
// layout: The channel layout, which must not be `nil`.
//
// # Return Value
//
// A new [AVAudioFormat] instance.
//
// # Discussion
//
// The returned [AVAudioFormat] instance uses the [AVAudioPCMFormatFloat32]
// format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(standardFormatWithSampleRate:channelLayout:)
func NewAudioFormatStandardFormatWithSampleRateChannelLayout(sampleRate float64, layout IAVAudioChannelLayout) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initStandardFormatWithSampleRate:channelLayout:"), sampleRate, layout)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance with the specified sample rate and channel
// count.
//
// sampleRate: The sampling rate, in hertz.
//
// channels: The channel count.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// The returned [AVAudioFormat] instance uses the [AVAudioPCMFormatFloat32]
// format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(standardFormatWithSampleRate:channels:)
func NewAudioFormatStandardFormatWithSampleRateChannels(sampleRate float64, channels AVAudioChannelCount) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initStandardFormatWithSampleRate:channels:"), sampleRate, channels)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance from a Core Media audio format
// description.
//
// formatDescription: The Core Media audio format description.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if `formatDescription` isn’t
// valid.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(cmAudioFormatDescription:)
func NewAudioFormatWithCMAudioFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCMAudioFormatDescription:"), formatDescription)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance.
//
// format: The audio format.
//
// sampleRate: The sampling rate, in hertz.
//
// channels: The channel count.
//
// interleaved: The Boolean value that indicates whether `format` is in an interleaved
// state.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// For information about possible `format` values, see [AVAudioCommonFormat].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(commonFormat:sampleRate:channels:interleaved:)
//
// [AVAudioCommonFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
func NewAudioFormatWithCommonFormatSampleRateChannelsInterleaved(format AVAudioCommonFormat, sampleRate float64, channels AVAudioChannelCount, interleaved bool) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommonFormat:sampleRate:channels:interleaved:"), format, sampleRate, channels, interleaved)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance with the specified audio format, sample
// rate, interleaved state, and channel layout.
//
// format: The audio format.
//
// sampleRate: The sampling rate, in hertz.
//
// interleaved: The Boolean value that indicates whether `format` is in an interleaved
// state.
//
// layout: The channel layout, which must not be `nil`.
//
// # Return Value
//
// A new [AVAudioFormat] instance.
//
// # Discussion
//
// For information about possible `format` values, see [AVAudioCommonFormat].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(commonFormat:sampleRate:interleaved:channelLayout:)
//
// [AVAudioCommonFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
func NewAudioFormatWithCommonFormatSampleRateInterleavedChannelLayout(format AVAudioCommonFormat, sampleRate float64, interleaved bool, layout IAVAudioChannelLayout) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommonFormat:sampleRate:interleaved:channelLayout:"), format, sampleRate, interleaved, layout)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance using the specified settings dictionary.
//
// settings: The settings dictionary.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// Note that many settings dictionary elements aren’t relevant for the
// format, so this method ignores them. For information about supported
// dictionary values, see [Audio settings].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(settings:)
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
func NewAudioFormatWithSettings(settings foundation.INSDictionary) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSettings:"), settings)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance from a stream description.
//
// asbd: The audio stream description.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// If the [AudioStreamBasicDescription] specifies more than two channels, this
// method fails and returns `nil`. Instead, use the
// [InitWithStreamDescriptionChannelLayout] method.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(streamDescription:)
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
func NewAudioFormatWithStreamDescription(asbd unsafe.Pointer) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamDescription:"), asbd)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance from a stream description and channel
// layout.
//
// asbd: The audio stream description.
//
// layout: The channel layout.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// When `layout` is `nil`, and `asbd` specifies one or two channels, this
// method assumes mono or stereo layout, respectively.
//
// If the [AudioStreamBasicDescription] specifies more than two channels and
// `layout` is `nil`, this method fails and returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(streamDescription:channelLayout:)
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
func NewAudioFormatWithStreamDescriptionChannelLayout(asbd unsafe.Pointer, layout IAVAudioChannelLayout) AVAudioFormat {
	instance := getAVAudioFormatClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamDescription:channelLayout:"), asbd, layout)
	return AVAudioFormatFromID(rv)
}

// Creates an audio format instance as a deinterleaved float with the
// specified sample rate and channel layout.
//
// sampleRate: The sampling rate, in hertz.
//
// layout: The channel layout, which must not be `nil`.
//
// # Return Value
//
// A new [AVAudioFormat] instance.
//
// # Discussion
//
// The returned [AVAudioFormat] instance uses the [AVAudioPCMFormatFloat32]
// format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(standardFormatWithSampleRate:channelLayout:)
func (a AVAudioFormat) InitStandardFormatWithSampleRateChannelLayout(sampleRate float64, layout IAVAudioChannelLayout) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initStandardFormatWithSampleRate:channelLayout:"), sampleRate, layout)
	return rv
}

// Creates an audio format instance with the specified sample rate and channel
// count.
//
// sampleRate: The sampling rate, in hertz.
//
// channels: The channel count.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// The returned [AVAudioFormat] instance uses the [AVAudioPCMFormatFloat32]
// format.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(standardFormatWithSampleRate:channels:)
func (a AVAudioFormat) InitStandardFormatWithSampleRateChannels(sampleRate float64, channels AVAudioChannelCount) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initStandardFormatWithSampleRate:channels:"), sampleRate, channels)
	return rv
}

// Creates an audio format instance.
//
// format: The audio format.
//
// sampleRate: The sampling rate, in hertz.
//
// channels: The channel count.
//
// interleaved: The Boolean value that indicates whether `format` is in an interleaved
// state.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// For information about possible `format` values, see [AVAudioCommonFormat].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(commonFormat:sampleRate:channels:interleaved:)
//
// [AVAudioCommonFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
func (a AVAudioFormat) InitWithCommonFormatSampleRateChannelsInterleaved(format AVAudioCommonFormat, sampleRate float64, channels AVAudioChannelCount, interleaved bool) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithCommonFormat:sampleRate:channels:interleaved:"), format, sampleRate, channels, interleaved)
	return rv
}

// Creates an audio format instance with the specified audio format, sample
// rate, interleaved state, and channel layout.
//
// format: The audio format.
//
// sampleRate: The sampling rate, in hertz.
//
// interleaved: The Boolean value that indicates whether `format` is in an interleaved
// state.
//
// layout: The channel layout, which must not be `nil`.
//
// # Return Value
//
// A new [AVAudioFormat] instance.
//
// # Discussion
//
// For information about possible `format` values, see [AVAudioCommonFormat].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(commonFormat:sampleRate:interleaved:channelLayout:)
//
// [AVAudioCommonFormat]: https://developer.apple.com/documentation/AVFAudio/AVAudioCommonFormat
func (a AVAudioFormat) InitWithCommonFormatSampleRateInterleavedChannelLayout(format AVAudioCommonFormat, sampleRate float64, interleaved bool, layout IAVAudioChannelLayout) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithCommonFormat:sampleRate:interleaved:channelLayout:"), format, sampleRate, interleaved, layout)
	return rv
}

// Creates an audio format instance using the specified settings dictionary.
//
// settings: The settings dictionary.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// Note that many settings dictionary elements aren’t relevant for the
// format, so this method ignores them. For information about supported
// dictionary values, see [Audio settings].
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(settings:)
//
// [Audio settings]: https://developer.apple.com/documentation/AVFoundation/audio-settings
func (a AVAudioFormat) InitWithSettings(settings foundation.INSDictionary) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithSettings:"), settings)
	return rv
}

// Creates an audio format instance from a stream description.
//
// asbd: The audio stream description.
//
// asbd is a [*coreaudiotypes.AudioStreamBasicDescription].
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// If the [AudioStreamBasicDescription] specifies more than two channels, this
// method fails and returns `nil`. Instead, use the
// [InitWithStreamDescriptionChannelLayout] method.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(streamDescription:)
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
func (a AVAudioFormat) InitWithStreamDescription(asbd unsafe.Pointer) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithStreamDescription:"), asbd)
	return rv
}

// Creates an audio format instance from a stream description and channel
// layout.
//
// asbd: The audio stream description.
//
// layout: The channel layout.
//
// asbd is a [*coreaudiotypes.AudioStreamBasicDescription].
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if the initialization fails.
//
// # Discussion
//
// When `layout` is `nil`, and `asbd` specifies one or two channels, this
// method assumes mono or stereo layout, respectively.
//
// If the [AudioStreamBasicDescription] specifies more than two channels and
// `layout` is `nil`, this method fails and returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(streamDescription:channelLayout:)
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
func (a AVAudioFormat) InitWithStreamDescriptionChannelLayout(asbd unsafe.Pointer, layout IAVAudioChannelLayout) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithStreamDescription:channelLayout:"), asbd, layout)
	return rv
}

// Creates an audio format instance from a Core Media audio format
// description.
//
// formatDescription: The Core Media audio format description.
//
// # Return Value
//
// A new [AVAudioFormat] instance, or `nil` if `formatDescription` isn’t
// valid.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/init(cmAudioFormatDescription:)
func (a AVAudioFormat) InitWithCMAudioFormatDescription(formatDescription coremedia.CMFormatDescriptionRef) AVAudioFormat {
	rv := objc.Send[AVAudioFormat](a.ID, objc.Sel("initWithCMAudioFormatDescription:"), formatDescription)
	return rv
}
func (a AVAudioFormat) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The audio format properties of a stream of audio data.
//
// # Discussion
//
// Returns an [AudioStreamBasicDescription] that you use with lower-level
// audio APIs.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/streamDescription
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
func (a AVAudioFormat) StreamDescription() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a.ID, objc.Sel("streamDescription"))
	return rv
}

// The audio format sampling rate, in hertz.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/sampleRate
func (a AVAudioFormat) SampleRate() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("sampleRate"))
	return rv
}

// The number of channels of audio data.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/channelCount
func (a AVAudioFormat) ChannelCount() AVAudioChannelCount {
	rv := objc.Send[AVAudioChannelCount](a.ID, objc.Sel("channelCount"))
	return AVAudioChannelCount(rv)
}

// The underlying audio channel layout.
//
// # Discussion
//
// Only formats with more than two channels require channel layouts.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/channelLayout
func (a AVAudioFormat) ChannelLayout() IAVAudioChannelLayout {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("channelLayout"))
	return AVAudioChannelLayoutFromID(objc.ID(rv))
}

// The audio format description to use with Core Media APIs.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/formatDescription
func (a AVAudioFormat) FormatDescription() coremedia.CMFormatDescriptionRef {
	rv := objc.Send[coremedia.CMFormatDescriptionRef](a.ID, objc.Sel("formatDescription"))
	return coremedia.CMFormatDescriptionRef(rv)
}

// A Boolean value that indicates whether the samples mix into one stream.
//
// # Discussion
//
// This value is only valid for PCM formats.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/isInterleaved
func (a AVAudioFormat) Interleaved() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInterleaved"))
	return rv
}

// A Boolean value that indicates whether the format is in a deinterleaved
// native-endian float state.
//
// # Discussion
//
// This value returns true if the format is [AVAudioPCMFormatFloat32];
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/isStandard
func (a AVAudioFormat) Standard() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isStandard"))
	return rv
}

// The common format identifier instance.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/commonFormat
func (a AVAudioFormat) CommonFormat() AVAudioCommonFormat {
	rv := objc.Send[AVAudioCommonFormat](a.ID, objc.Sel("commonFormat"))
	return AVAudioCommonFormat(rv)
}

// A dictionary that represents the format as a dictionary using audio setting
// keys.
//
// # Discussion
//
// The settings dictionary doesn’t support all formats that
// [AudioStreamBasicDescription] represents (the underlying implementation),
// in which case, this property returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/settings
//
// [AudioStreamBasicDescription]: https://developer.apple.com/documentation/CoreAudioTypes/AudioStreamBasicDescription
func (a AVAudioFormat) Settings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("settings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// An object that contains metadata that encoders and decoders require.
//
// # Discussion
//
// Encoders produce a `magicCookie` object, and some decoders require it to
// decode properly.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioFormat/magicCookie
func (a AVAudioFormat) MagicCookie() foundation.INSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("magicCookie"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (a AVAudioFormat) SetMagicCookie(value foundation.INSData) {
	objc.Send[struct{}](a.ID, objc.Sel("setMagicCookie:"), value)
}

// See: https://developer.apple.com/documentation/avfaudio/avchannellayoutkey
func (a AVAudioFormat) AVChannelLayoutKey() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("AVChannelLayoutKey"))
	return foundation.NSStringFromID(rv).String()
}
