// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetWriterInput] class.
var (
	_AVAssetWriterInputClass     AVAssetWriterInputClass
	_AVAssetWriterInputClassOnce sync.Once
)

func getAVAssetWriterInputClass() AVAssetWriterInputClass {
	_AVAssetWriterInputClassOnce.Do(func() {
		_AVAssetWriterInputClass = AVAssetWriterInputClass{class: objc.GetClass("AVAssetWriterInput")}
	})
	return _AVAssetWriterInputClass
}

// GetAVAssetWriterInputClass returns the class object for AVAssetWriterInput.
func GetAVAssetWriterInputClass() AVAssetWriterInputClass {
	return getAVAssetWriterInputClass()
}

type AVAssetWriterInputClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetWriterInputClass) Alloc() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that appends media samples to a track in an asset writer’s
// output file.
//
// # Overview
// 
// Create an asset writer input to write a single track of media, and optional
// track-level metadata, to the output file. To write multiple concurrent
// tracks with ideal interleaving of media data, observe the value of the
// [isReadyForMoreMediaData] property of each input.
// 
// You can use an asset writer input to create tracks in a QuickTime movie
// file that aren’t self-contained, and instead reference sample data that
// exists in another file.
//
// [isReadyForMoreMediaData]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/isReadyForMoreMediaData
//
// # Creating an input
//
//   - [AVAssetWriterInput.InitWithMediaTypeOutputSettings]: Creates an input to append sample buffers of the specified type to the output file.
//   - [AVAssetWriterInput.InitWithMediaTypeOutputSettingsSourceFormatHint]: Creates an input that appends sample buffers of the specified type and format hint to the output file.
//
// # Configuring presentation
//
//   - [AVAssetWriterInput.NaturalSize]: The natural display dimensions of the output’s visual media.
//   - [AVAssetWriterInput.SetNaturalSize]
//   - [AVAssetWriterInput.Transform]: The transform to use for display of the output’s visual media.
//   - [AVAssetWriterInput.SetTransform]
//   - [AVAssetWriterInput.PreferredVolume]: The volume to prefer for playback of the output’s audio data.
//   - [AVAssetWriterInput.SetPreferredVolume]
//   - [AVAssetWriterInput.MediaTimeScale]: The time scale of the track in the output file.
//   - [AVAssetWriterInput.SetMediaTimeScale]
//   - [AVAssetWriterInput.MarksOutputTrackAsEnabled]: A Boolean value that indicates whether to enable a track in the output for playback and processing.
//   - [AVAssetWriterInput.SetMarksOutputTrackAsEnabled]
//
// # Configuring language support
//
//   - [AVAssetWriterInput.LanguageCode]: The language code of the input’s track.
//   - [AVAssetWriterInput.SetLanguageCode]
//   - [AVAssetWriterInput.ExtendedLanguageTag]: The extended language for the input’s track.
//   - [AVAssetWriterInput.SetExtendedLanguageTag]
//
// # Configuring metadata
//
//   - [AVAssetWriterInput.Metadata]: The track-level metadata to write to the output.
//   - [AVAssetWriterInput.SetMetadata]
//
// # Configuring media data layout
//
//   - [AVAssetWriterInput.PreferredMediaChunkAlignment]: The boundary, in bytes, for aligning media chunks.
//   - [AVAssetWriterInput.SetPreferredMediaChunkAlignment]
//   - [AVAssetWriterInput.PreferredMediaChunkDuration]: The duration to use for each chunk of sample data in the output file.
//   - [AVAssetWriterInput.SetPreferredMediaChunkDuration]
//   - [AVAssetWriterInput.SampleReferenceBaseURL]: The base URL sample references are relative to.
//   - [AVAssetWriterInput.SetSampleReferenceBaseURL]
//   - [AVAssetWriterInput.MediaDataLocation]: Specifies how the input lays out and interleaves media data.
//   - [AVAssetWriterInput.SetMediaDataLocation]
//
// # Configuring track associations
//
//   - [AVAssetWriterInput.CanAddTrackAssociationWithTrackOfInputType]: Determines whether it’s valid to associate another input’s track with this input’s track.
//   - [AVAssetWriterInput.AddTrackAssociationWithTrackOfInputType]: Adds an association between input tracks.
//
// # Appending media samples
//
//   - [AVAssetWriterInput.MarkAsFinished]: Marks the input as finished to indicate that you’re done appending samples to it.
//
// # Performing multiple-pass encoding
//
//   - [AVAssetWriterInput.CanPerformMultiplePasses]: A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//   - [AVAssetWriterInput.CurrentPassDescription]: An object that describes the requirements for the current pass.
//   - [AVAssetWriterInput.MarkCurrentPassAsFinished]: Tells the input to analyze the appended media to determine whether it can improve the results by reencoding certain segments.
//   - [AVAssetWriterInput.PerformsMultiPassEncodingIfSupported]: A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//   - [AVAssetWriterInput.SetPerformsMultiPassEncodingIfSupported]
//   - [AVAssetWriterInput.RespondToEachPassDescriptionOnQueueUsingBlock]: Tells the input to invoke a callback whenever it begins a new pass.
//
// # Inspecting an input
//
//   - [AVAssetWriterInput.MediaType]: The media type of the samples that the input accepts.
//   - [AVAssetWriterInput.OutputSettings]: The settings to use for encoding media data you append to the output.
//   - [AVAssetWriterInput.SourceFormatHint]: A hint about the format of the sample buffers to append to the input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput
type AVAssetWriterInput struct {
	objectivec.Object
}

// AVAssetWriterInputFromID constructs a [AVAssetWriterInput] from an objc.ID.
//
// An object that appends media samples to a track in an asset writer’s
// output file.
func AVAssetWriterInputFromID(id objc.ID) AVAssetWriterInput {
	return AVAssetWriterInput{objectivec.Object{ID: id}}
}
// NOTE: AVAssetWriterInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetWriterInput] class.
//
// # Creating an input
//
//   - [IAVAssetWriterInput.InitWithMediaTypeOutputSettings]: Creates an input to append sample buffers of the specified type to the output file.
//   - [IAVAssetWriterInput.InitWithMediaTypeOutputSettingsSourceFormatHint]: Creates an input that appends sample buffers of the specified type and format hint to the output file.
//
// # Configuring presentation
//
//   - [IAVAssetWriterInput.NaturalSize]: The natural display dimensions of the output’s visual media.
//   - [IAVAssetWriterInput.SetNaturalSize]
//   - [IAVAssetWriterInput.Transform]: The transform to use for display of the output’s visual media.
//   - [IAVAssetWriterInput.SetTransform]
//   - [IAVAssetWriterInput.PreferredVolume]: The volume to prefer for playback of the output’s audio data.
//   - [IAVAssetWriterInput.SetPreferredVolume]
//   - [IAVAssetWriterInput.MediaTimeScale]: The time scale of the track in the output file.
//   - [IAVAssetWriterInput.SetMediaTimeScale]
//   - [IAVAssetWriterInput.MarksOutputTrackAsEnabled]: A Boolean value that indicates whether to enable a track in the output for playback and processing.
//   - [IAVAssetWriterInput.SetMarksOutputTrackAsEnabled]
//
// # Configuring language support
//
//   - [IAVAssetWriterInput.LanguageCode]: The language code of the input’s track.
//   - [IAVAssetWriterInput.SetLanguageCode]
//   - [IAVAssetWriterInput.ExtendedLanguageTag]: The extended language for the input’s track.
//   - [IAVAssetWriterInput.SetExtendedLanguageTag]
//
// # Configuring metadata
//
//   - [IAVAssetWriterInput.Metadata]: The track-level metadata to write to the output.
//   - [IAVAssetWriterInput.SetMetadata]
//
// # Configuring media data layout
//
//   - [IAVAssetWriterInput.PreferredMediaChunkAlignment]: The boundary, in bytes, for aligning media chunks.
//   - [IAVAssetWriterInput.SetPreferredMediaChunkAlignment]
//   - [IAVAssetWriterInput.PreferredMediaChunkDuration]: The duration to use for each chunk of sample data in the output file.
//   - [IAVAssetWriterInput.SetPreferredMediaChunkDuration]
//   - [IAVAssetWriterInput.SampleReferenceBaseURL]: The base URL sample references are relative to.
//   - [IAVAssetWriterInput.SetSampleReferenceBaseURL]
//   - [IAVAssetWriterInput.MediaDataLocation]: Specifies how the input lays out and interleaves media data.
//   - [IAVAssetWriterInput.SetMediaDataLocation]
//
// # Configuring track associations
//
//   - [IAVAssetWriterInput.CanAddTrackAssociationWithTrackOfInputType]: Determines whether it’s valid to associate another input’s track with this input’s track.
//   - [IAVAssetWriterInput.AddTrackAssociationWithTrackOfInputType]: Adds an association between input tracks.
//
// # Appending media samples
//
//   - [IAVAssetWriterInput.MarkAsFinished]: Marks the input as finished to indicate that you’re done appending samples to it.
//
// # Performing multiple-pass encoding
//
//   - [IAVAssetWriterInput.CanPerformMultiplePasses]: A Boolean value that indicates whether the input may perform multiple passes over appended media data.
//   - [IAVAssetWriterInput.CurrentPassDescription]: An object that describes the requirements for the current pass.
//   - [IAVAssetWriterInput.MarkCurrentPassAsFinished]: Tells the input to analyze the appended media to determine whether it can improve the results by reencoding certain segments.
//   - [IAVAssetWriterInput.PerformsMultiPassEncodingIfSupported]: A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
//   - [IAVAssetWriterInput.SetPerformsMultiPassEncodingIfSupported]
//   - [IAVAssetWriterInput.RespondToEachPassDescriptionOnQueueUsingBlock]: Tells the input to invoke a callback whenever it begins a new pass.
//
// # Inspecting an input
//
//   - [IAVAssetWriterInput.MediaType]: The media type of the samples that the input accepts.
//   - [IAVAssetWriterInput.OutputSettings]: The settings to use for encoding media data you append to the output.
//   - [IAVAssetWriterInput.SourceFormatHint]: A hint about the format of the sample buffers to append to the input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput
type IAVAssetWriterInput interface {
	objectivec.IObject

	// Topic: Creating an input

	// Creates an input to append sample buffers of the specified type to the output file.
	InitWithMediaTypeOutputSettings(mediaType AVMediaType, outputSettings foundation.INSDictionary) AVAssetWriterInput
	// Creates an input that appends sample buffers of the specified type and format hint to the output file.
	InitWithMediaTypeOutputSettingsSourceFormatHint(mediaType AVMediaType, outputSettings foundation.INSDictionary, sourceFormatHint objectivec.IObject) AVAssetWriterInput

	// Topic: Configuring presentation

	// The natural display dimensions of the output’s visual media.
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	// The transform to use for display of the output’s visual media.
	Transform() corefoundation.CGAffineTransform
	SetTransform(value corefoundation.CGAffineTransform)
	// The volume to prefer for playback of the output’s audio data.
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	// The time scale of the track in the output file.
	MediaTimeScale() int32
	SetMediaTimeScale(value int32)
	// A Boolean value that indicates whether to enable a track in the output for playback and processing.
	MarksOutputTrackAsEnabled() bool
	SetMarksOutputTrackAsEnabled(value bool)

	// Topic: Configuring language support

	// The language code of the input’s track.
	LanguageCode() string
	SetLanguageCode(value string)
	// The extended language for the input’s track.
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)

	// Topic: Configuring metadata

	// The track-level metadata to write to the output.
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)

	// Topic: Configuring media data layout

	// The boundary, in bytes, for aligning media chunks.
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	// The duration to use for each chunk of sample data in the output file.
	PreferredMediaChunkDuration() objectivec.IObject
	SetPreferredMediaChunkDuration(value objectivec.IObject)
	// The base URL sample references are relative to.
	SampleReferenceBaseURL() foundation.INSURL
	SetSampleReferenceBaseURL(value foundation.INSURL)
	// Specifies how the input lays out and interleaves media data.
	MediaDataLocation() AVAssetWriterInputMediaDataLocation
	SetMediaDataLocation(value AVAssetWriterInputMediaDataLocation)

	// Topic: Configuring track associations

	// Determines whether it’s valid to associate another input’s track with this input’s track.
	CanAddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType string) bool
	// Adds an association between input tracks.
	AddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType string)

	// Topic: Appending media samples

	// Marks the input as finished to indicate that you’re done appending samples to it.
	MarkAsFinished()

	// Topic: Performing multiple-pass encoding

	// A Boolean value that indicates whether the input may perform multiple passes over appended media data.
	CanPerformMultiplePasses() bool
	// An object that describes the requirements for the current pass.
	CurrentPassDescription() IAVAssetWriterInputPassDescription
	// Tells the input to analyze the appended media to determine whether it can improve the results by reencoding certain segments.
	MarkCurrentPassAsFinished()
	// A Boolean value that indicates whether the input attempts to encode the source media data using multiple passes.
	PerformsMultiPassEncodingIfSupported() bool
	SetPerformsMultiPassEncodingIfSupported(value bool)
	// Tells the input to invoke a callback whenever it begins a new pass.
	RespondToEachPassDescriptionOnQueueUsingBlock(queue dispatch.Queue, block unsafe.Pointer)

	// Topic: Inspecting an input

	// The media type of the samples that the input accepts.
	MediaType() AVMediaType
	// The settings to use for encoding media data you append to the output.
	OutputSettings() foundation.INSDictionary
	// A hint about the format of the sample buffers to append to the input.
	SourceFormatHint() objectivec.IObject
}

// Init initializes the instance.
func (a AVAssetWriterInput) Init() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetWriterInput) Autorelease() AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetWriterInput creates a new AVAssetWriterInput instance.
func NewAVAssetWriterInput() AVAssetWriterInput {
	class := getAVAssetWriterInputClass()
	rv := objc.Send[AVAssetWriterInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an input to append sample buffers of the specified type to the
// output file.
//
// mediaType: The media type of the samples the input accepts.
//
// outputSettings: The settings to use for encoding the media you append to the output. Create
// an output settings dictionary manually, or use [AVOutputSettingsAssistant]
// to create preset-based settings.
//
// # Discussion
// 
// If you’re appending samples that are already in an acceptable compressed
// format, pass a value of `nil` for the output settings to pass the buffers
// to the output unaltered. However, if you’re not writing to a QuickTime
// movie file, an asset writer only supports passing through a restricted set
// of media types and subtypes. To pass through media data to files with a
// type other than [mov], pass a nonnull format hint using the
// [InitWithMediaTypeOutputSettingsSourceFormatHint] instead.
// 
// # Configuring audio settings
// 
// You must fully specify the audio settings dictionary when using this
// initializer, which means you must provide values for the following keys:
// 
// - [AVFormatIDKey]. The identifier of the audio format. For
// [KAudioFormatLinearPCM] format, you must include values for all relevant
// keys with a [AVLinearPCM] prefix, and for [kAudioFormatAppleLossless], you
// must specify a value for [AVEncoderBitDepthHintKey]. - [AVSampleRateKey].
// The sample rate of the audio. Common values are `44100` and `48000`. -
// [AVNumberOfChannelsKey]. If no other channel layout information is
// available, specifying a value of `1` results in mono output and a value of
// `2` results in stereo output. If [AVNumberOfChannelsKey] specifies a
// channel count greater than `2`, the dictionary must also specify a value
// for [AVChannelLayoutKey].
// 
// # Configuring video settings
// 
// A video output settings dictionary must request a compressed video format,
// which means that the value you specify must follow the rules for compressed
// video output.
// 
// You must fully specify the video settings dictionary when using this
// initializer, which means you must provide values for the following keys
// [AVVideoCodecKey], [AVVideoWidthKey], [AVVideoHeightKey].
//
// [AVChannelLayoutKey]: https://developer.apple.com/documentation/AVFAudio/AVChannelLayoutKey
// [AVEncoderBitDepthHintKey]: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitDepthHintKey
// [AVNumberOfChannelsKey]: https://developer.apple.com/documentation/AVFAudio/AVNumberOfChannelsKey
// [AVSampleRateKey]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateKey
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoHeightKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoHeightKey
// [AVVideoWidthKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoWidthKey
// [kAudioFormatAppleLossless]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatAppleLossless
// [mov]: https://developer.apple.com/documentation/AVFoundation/AVFileType/mov
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:)
func NewAssetWriterInputWithMediaTypeOutputSettings(mediaType AVMediaType, outputSettings foundation.INSDictionary) AVAssetWriterInput {
	instance := getAVAssetWriterInputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMediaType:outputSettings:"), objc.String(string(mediaType)), outputSettings)
	return AVAssetWriterInputFromID(rv)
}

// Creates an input that appends sample buffers of the specified type and
// format hint to the output file.
//
// mediaType: The type of media that an input accepts.
//
// outputSettings: The settings to use for encoding the media you append to the output. Create
// an output settings dictionary manually, or use [AVOutputSettingsAssistant]
// to create preset-based settings.
//
// sourceFormatHint: A hint about the format of the media data to append. The input uses the
// source format hint to fill in missing output settings. If you specify a
// hint, you only need to specify [AVFormatIDKey] for the audio output
// settings, and [AVVideoCodecKey] is the only required key for video output
// settings.
// 
// The system raises an error if the format description isn’t valid for the
// indicated media type.
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
//
// # Discussion
// 
// To guarantee successful file writing, ensure that sample buffers you append
// are of the specified format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:sourceFormatHint:)
// sourceFormatHint is a [coremedia.CMFormatDescriptionRef].
func NewAssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType AVMediaType, outputSettings foundation.INSDictionary, sourceFormatHint objectivec.IObject) AVAssetWriterInput {
	instance := getAVAssetWriterInputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMediaType:outputSettings:sourceFormatHint:"), objc.String(string(mediaType)), outputSettings, sourceFormatHint)
	return AVAssetWriterInputFromID(rv)
}

// Creates an input to append sample buffers of the specified type to the
// output file.
//
// mediaType: The media type of the samples the input accepts.
//
// outputSettings: The settings to use for encoding the media you append to the output. Create
// an output settings dictionary manually, or use [AVOutputSettingsAssistant]
// to create preset-based settings.
//
// # Discussion
// 
// If you’re appending samples that are already in an acceptable compressed
// format, pass a value of `nil` for the output settings to pass the buffers
// to the output unaltered. However, if you’re not writing to a QuickTime
// movie file, an asset writer only supports passing through a restricted set
// of media types and subtypes. To pass through media data to files with a
// type other than [mov], pass a nonnull format hint using the
// [InitWithMediaTypeOutputSettingsSourceFormatHint] instead.
// 
// # Configuring audio settings
// 
// You must fully specify the audio settings dictionary when using this
// initializer, which means you must provide values for the following keys:
// 
// - [AVFormatIDKey]. The identifier of the audio format. For
// [KAudioFormatLinearPCM] format, you must include values for all relevant
// keys with a [AVLinearPCM] prefix, and for [kAudioFormatAppleLossless], you
// must specify a value for [AVEncoderBitDepthHintKey]. - [AVSampleRateKey].
// The sample rate of the audio. Common values are `44100` and `48000`. -
// [AVNumberOfChannelsKey]. If no other channel layout information is
// available, specifying a value of `1` results in mono output and a value of
// `2` results in stereo output. If [AVNumberOfChannelsKey] specifies a
// channel count greater than `2`, the dictionary must also specify a value
// for [AVChannelLayoutKey].
// 
// # Configuring video settings
// 
// A video output settings dictionary must request a compressed video format,
// which means that the value you specify must follow the rules for compressed
// video output.
// 
// You must fully specify the video settings dictionary when using this
// initializer, which means you must provide values for the following keys
// [AVVideoCodecKey], [AVVideoWidthKey], [AVVideoHeightKey].
//
// [AVChannelLayoutKey]: https://developer.apple.com/documentation/AVFAudio/AVChannelLayoutKey
// [AVEncoderBitDepthHintKey]: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitDepthHintKey
// [AVNumberOfChannelsKey]: https://developer.apple.com/documentation/AVFAudio/AVNumberOfChannelsKey
// [AVSampleRateKey]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateKey
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoHeightKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoHeightKey
// [AVVideoWidthKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoWidthKey
// [kAudioFormatAppleLossless]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatAppleLossless
// [mov]: https://developer.apple.com/documentation/AVFoundation/AVFileType/mov
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:)
func (a AVAssetWriterInput) InitWithMediaTypeOutputSettings(mediaType AVMediaType, outputSettings foundation.INSDictionary) AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a.ID, objc.Sel("initWithMediaType:outputSettings:"), objc.String(string(mediaType)), outputSettings)
	return rv
}
// Creates an input that appends sample buffers of the specified type and
// format hint to the output file.
//
// mediaType: The type of media that an input accepts.
//
// outputSettings: The settings to use for encoding the media you append to the output. Create
// an output settings dictionary manually, or use [AVOutputSettingsAssistant]
// to create preset-based settings.
//
// sourceFormatHint: A hint about the format of the media data to append. The input uses the
// source format hint to fill in missing output settings. If you specify a
// hint, you only need to specify [AVFormatIDKey] for the audio output
// settings, and [AVVideoCodecKey] is the only required key for video output
// settings.
// 
// The system raises an error if the format description isn’t valid for the
// indicated media type.
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
//
// sourceFormatHint is a [coremedia.CMFormatDescriptionRef].
//
// # Discussion
// 
// To guarantee successful file writing, ensure that sample buffers you append
// are of the specified format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/init(mediaType:outputSettings:sourceFormatHint:)
func (a AVAssetWriterInput) InitWithMediaTypeOutputSettingsSourceFormatHint(mediaType AVMediaType, outputSettings foundation.INSDictionary, sourceFormatHint objectivec.IObject) AVAssetWriterInput {
	rv := objc.Send[AVAssetWriterInput](a.ID, objc.Sel("initWithMediaType:outputSettings:sourceFormatHint:"), objc.String(string(mediaType)), outputSettings, sourceFormatHint)
	return rv
}
// Determines whether it’s valid to associate another input’s track with
// this input’s track.
//
// input: An asset writer input that contains the track to associate.
//
// trackAssociationType: The type of track association to test.
//
// # Return Value
// 
// [true] if the system can make the association between tracks; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method returns [false] if the association type requires tracks of a
// media type that doesn’t match the input’s type, or if the output file
// type doesn’t support track associations.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/canAddTrackAssociation(withTrackOf:type:)
func (a AVAssetWriterInput) CanAddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType string) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canAddTrackAssociationWithTrackOfInput:type:"), input, objc.String(trackAssociationType))
	return rv
}
// Adds an association between input tracks.
//
// input: The input that contains the track to associate with this input’s track.
//
// trackAssociationType: The type of track association to add.
//
// # Discussion
// 
// The system raises an error if the association type requires tracks of a
// media type that doesn’t match the input’s type, or if the output file
// type doesn’t support track associations.
// 
// You can’t add track associations after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/addTrackAssociation(withTrackOf:type:)
func (a AVAssetWriterInput) AddTrackAssociationWithTrackOfInputType(input IAVAssetWriterInput, trackAssociationType string) {
	objc.Send[objc.ID](a.ID, objc.Sel("addTrackAssociationWithTrackOfInput:type:"), input, objc.String(trackAssociationType))
}
// Marks the input as finished to indicate that you’re done appending
// samples to it.
//
// # Discussion
// 
// Apps that monitor an input’s [isReadyForMoreMediaData] value must call
// this method when they finish appending to it. This is necessary to prevent
// other inputs from stalling because they’re waiting on the input’s media
// data to complete the ideal interleaving pattern.
// 
// After calling this method from the serial queue passed to
// [requestMediaDataWhenReady(on:using:)], the input issues no more
// invocations of the callback it passes to that method.
//
// [isReadyForMoreMediaData]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/isReadyForMoreMediaData
// [requestMediaDataWhenReady(on:using:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/requestMediaDataWhenReady(on:using:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/markAsFinished()
func (a AVAssetWriterInput) MarkAsFinished() {
	objc.Send[objc.ID](a.ID, objc.Sel("markAsFinished"))
}
// Tells the input to analyze the appended media to determine whether it can
// improve the results by reencoding certain segments.
//
// # Discussion
// 
// When the value of the [CanPerformMultiplePasses] property is [true], call
// this method after you append all of your media data. After the input
// determines if it warrants performing an additional pass, the value of
// [CurrentPassDescription] changes (typically asynchronously) to describe how
// to set up for the next pass. Although it’s possible to use key-value
// observing to determine when the value of [CurrentPassDescription] changes,
// it’s typically more convenient to call the
// [RespondToEachPassDescriptionOnQueueUsingBlock] method to start the work
// for each pass.
// 
// After reappending the media data for all of the time ranges of the new
// pass, call this method again to determine whether to reappend segments in
// another pass.
// 
// Calling this method effectively cancels any previous invocation of
// [requestMediaDataWhenReady(on:using:)], which means that you may call
// [requestMediaDataWhenReady(on:using:)] again for each new pass. This method
// provides a convenient way to consolidate these invocations in your code.
// 
// After each pass, you have the option of keeping the most recent results by
// calling [MarkAsFinished], instead of this method. If the value of
// [CurrentPassDescription] is `nil` at the beginning of a pass, call
// [MarkAsFinished] to tell the input to not expect any further media data.
// 
// If the value of [CanPerformMultiplePasses] is [false], the value of
// [CurrentPassDescription] immediately becomes `nil` after calling this
// method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [requestMediaDataWhenReady(on:using:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/requestMediaDataWhenReady(on:using:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/markCurrentPassAsFinished()
func (a AVAssetWriterInput) MarkCurrentPassAsFinished() {
	objc.Send[objc.ID](a.ID, objc.Sel("markCurrentPassAsFinished"))
}
// Tells the input to invoke a callback whenever it begins a new pass.
//
// queue: The queue on which to invoke the callback.
//
// block: A callback the input invokes at the beginning of each pass.
//
// # Discussion
// 
// A typical implemementation of the callback block performs the following
// steps:
// 
// - Gets the value of the [CurrentPassDescription] property and configures
// media data source accordingly. - Calls the
// [requestMediaDataWhenReady(on:using:)] method to begin appending data for
// the current pass.
// 
// After you’ve appended all media data for the current pass, call the
// [MarkCurrentPassAsFinished] method have the system determine whether to
// perform another pass. If it performs an additional pass the system invokes
// the callback to begin the next pass. When it determines that it requires no
// additional passes, the system invokes the callback one final time so the
// client can invoke [MarkAsFinished] in response to the value of
// [CurrentPassDescription] becoming `nil`.
//
// [requestMediaDataWhenReady(on:using:)]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/requestMediaDataWhenReady(on:using:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/respondToEachPassDescription(on:using:)
func (a AVAssetWriterInput) RespondToEachPassDescriptionOnQueueUsingBlock(queue dispatch.Queue, block unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("respondToEachPassDescriptionOnQueue:usingBlock:"), uintptr(queue.Handle()), block)
}

// Returns a new input to append sample buffers of the specified type to the
// output file.
//
// mediaType: The media type of the samples the input accepts.
//
// outputSettings: The settings to use for encoding the media you append to the output. Create
// an output settings dictionary manually, or use [AVOutputSettingsAssistant]
// to create preset-based settings.
//
// # Return Value
// 
// A new asset writer input.
//
// # Discussion
// 
// If you’re appending samples that are already in an acceptable compressed
// format, pass a value of `nil` for the output settings to pass the buffers
// to the output unaltered. However, if you’re not writing to a QuickTime
// movie file, an asset writer only supports passing through a restricted set
// of media types and subtypes. To pass through media data to files with a
// type other than [mov], pass a nonnull format hint using the
// [InitWithMediaTypeOutputSettingsSourceFormatHint] instead.
// 
// # Configuring audio settings
// 
// You must fully specify the audio settings dictionary when using this
// initializer, which means you must provide values for the following keys:
// 
// - [AVFormatIDKey]. The identifier of the audio format. For
// [KAudioFormatLinearPCM] format, you must include values for all relevant
// keys with a [AVLinearPCM] prefix, and for [kAudioFormatAppleLossless], you
// must specify a value for [AVEncoderBitDepthHintKey]. - [AVSampleRateKey].
// The sample rate of the audio. Common values are `44100` and `48000`. -
// [AVNumberOfChannelsKey]. If no other channel layout information is
// available, specifying a value of `1` results in mono output and a value of
// `2` results in stereo output. If [AVNumberOfChannelsKey] specifies a
// channel count greater than `2`, the dictionary must also specify a value
// for [AVChannelLayoutKey].
// 
// # Configuring video settings
// 
// A video output settings dictionary must request a compressed video format,
// which means that the value you specify must follow the rules for compressed
// video output.
// 
// You must fully specify the video settings dictionary when using this
// initializer, which means you must provide values for the following keys
// [AVVideoCodecKey], [AVVideoWidthKey], [AVVideoHeightKey].
//
// [AVChannelLayoutKey]: https://developer.apple.com/documentation/AVFAudio/AVChannelLayoutKey
// [AVEncoderBitDepthHintKey]: https://developer.apple.com/documentation/AVFAudio/AVEncoderBitDepthHintKey
// [AVNumberOfChannelsKey]: https://developer.apple.com/documentation/AVFAudio/AVNumberOfChannelsKey
// [AVSampleRateKey]: https://developer.apple.com/documentation/AVFAudio/AVSampleRateKey
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoHeightKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoHeightKey
// [AVVideoWidthKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoWidthKey
// [kAudioFormatAppleLossless]: https://developer.apple.com/documentation/CoreAudioTypes/kAudioFormatAppleLossless
// [mov]: https://developer.apple.com/documentation/AVFoundation/AVFileType/mov
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/assetWriterInputWithMediaType:outputSettings:
func (_AVAssetWriterInputClass AVAssetWriterInputClass) AssetWriterInputWithMediaTypeOutputSettings(mediaType AVMediaType, outputSettings foundation.INSDictionary) AVAssetWriterInput {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetWriterInputClass.class), objc.Sel("assetWriterInputWithMediaType:outputSettings:"), objc.String(string(mediaType)), outputSettings)
	return AVAssetWriterInputFromID(rv)
}
// Returns a new input that appends sample buffers of the specified type and
// format hint to the output file.
//
// mediaType: The type of media that an input accepts.
//
// outputSettings: The settings to use for encoding the media you append to the output. Create
// an output settings dictionary manually, or use [AVOutputSettingsAssistant]
// to create preset-based settings.
//
// sourceFormatHint: A hint about the format of the media data to append. The input uses the
// source format hint to fill in missing output settings. If you specify a
// hint, you only need to specify [AVFormatIDKey] for the audio output
// settings, and [AVVideoCodecKey] is the only required key for video output
// settings.
// 
// The system raises an error if the format description isn’t valid for the
// indicated media type.
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
//
// sourceFormatHint is a [coremedia.CMFormatDescriptionRef].
//
// # Return Value
// 
// A new asset writer input.
//
// # Discussion
// 
// To guarantee successful file writing, ensure that sample buffers you append
// are of the specified format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/assetWriterInputWithMediaType:outputSettings:sourceFormatHint:
func (_AVAssetWriterInputClass AVAssetWriterInputClass) AssetWriterInputWithMediaTypeOutputSettingsSourceFormatHint(mediaType AVMediaType, outputSettings foundation.INSDictionary, sourceFormatHint objectivec.IObject) AVAssetWriterInput {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetWriterInputClass.class), objc.Sel("assetWriterInputWithMediaType:outputSettings:sourceFormatHint:"), objc.String(string(mediaType)), outputSettings, sourceFormatHint)
	return AVAssetWriterInputFromID(rv)
}

// The natural display dimensions of the output’s visual media.
//
// # Discussion
// 
// The default value is [CGSizeZero], which indicates the system sets the
// natural size according to the dimensions of the output track’s format
// descriptions.
// 
// You can’t set this value after writing starts.
//
// [CGSizeZero]: https://developer.apple.com/documentation/CoreGraphics/CGSizeZero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/naturalSize
func (a AVAssetWriterInput) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("naturalSize"))
	return corefoundation.CGSize(rv)
}
func (a AVAssetWriterInput) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[struct{}](a.ID, objc.Sel("setNaturalSize:"), value)
}
// The transform to use for display of the output’s visual media.
//
// # Discussion
// 
// By default, the input uses the [CGAffineTransformIdentity] transform.
// 
// You can’t set this value after writing starts.
//
// [CGAffineTransformIdentity]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIdentity
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/transform
func (a AVAssetWriterInput) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](a.ID, objc.Sel("transform"))
	return corefoundation.CGAffineTransform(rv)
}
func (a AVAssetWriterInput) SetTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransform:"), value)
}
// The volume to prefer for playback of the output’s audio data.
//
// # Discussion
// 
// The default value for audio data is `1.0`, which indicates typical playback
// level. Set the value for this property in the range of `0.0` to `1.0`. For
// nonaudio media, the default value is `0.0`.
// 
// You can’t set this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredVolume
func (a AVAssetWriterInput) PreferredVolume() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("preferredVolume"))
	return rv
}
func (a AVAssetWriterInput) SetPreferredVolume(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreferredVolume:"), value)
}
// The time scale of the track in the output file.
//
// # Discussion
// 
// The default value is `0`, which indicates that the input chooses an
// appropriate value, if applicable. It’s an error to set this value if the
// input’s media type is [audio].
// 
// You can’t set this value after writing starts.
//
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaTimeScale
func (a AVAssetWriterInput) MediaTimeScale() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("mediaTimeScale"))
	return rv
}
func (a AVAssetWriterInput) SetMediaTimeScale(value int32) {
	objc.Send[struct{}](a.ID, objc.Sel("setMediaTimeScale:"), value)
}
// A Boolean value that indicates whether to enable a track in the output for
// playback and processing.
//
// # Discussion
// 
// The default value is [true]. If the format you’re writing supports
// disabling tracks, you can disable a track by setting this value to [false].
// 
// You can’t set this value after writing starts.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/marksOutputTrackAsEnabled
func (a AVAssetWriterInput) MarksOutputTrackAsEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("marksOutputTrackAsEnabled"))
	return rv
}
func (a AVAssetWriterInput) SetMarksOutputTrackAsEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setMarksOutputTrackAsEnabled:"), value)
}
// The language code of the input’s track.
//
// # Discussion
// 
// Specify an ISO 639-2/T language code value, or `nil` to prevent the writer
// from writing a language code.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/languageCode
func (a AVAssetWriterInput) LanguageCode() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("languageCode"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAssetWriterInput) SetLanguageCode(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}
// The extended language for the input’s track.
//
// # Discussion
// 
// Extended language tags are normally set only when an ISO 639-2/T language
// code alone is ambiguous. For example, you may use an extended language tag
// to distinguish media by the regional dialect in use or the writing system
// employed.
// 
// Specify the value as an RFC 4646 language tag, or `nil` to prevent the
// writer from writing an extended language tag.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/extendedLanguageTag
func (a AVAssetWriterInput) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVAssetWriterInput) SetExtendedLanguageTag(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}
// The track-level metadata to write to the output.
//
// # Discussion
// 
// You can’t set this property after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/metadata
func (a AVAssetWriterInput) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (a AVAssetWriterInput) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](a.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
}
// The boundary, in bytes, for aligning media chunks.
//
// # Discussion
// 
// This property supports file types that support media chunk alignment, such
// as QuickTime Movie files. The default value is `0`, which means that the
// input chooses an appropriate default value. A value of `1` indicates not to
// use padding to achieve a particular chunk alignment. It’s an error to set
// a negative value for chunk alignment.
// 
// You can’t set this value after writing starts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredMediaChunkAlignment
func (a AVAssetWriterInput) PreferredMediaChunkAlignment() int {
	rv := objc.Send[int](a.ID, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}
func (a AVAssetWriterInput) SetPreferredMediaChunkAlignment(value int) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}
// The duration to use for each chunk of sample data in the output file.
//
// # Discussion
// 
// This property supports file types that support media chunk alignment, such
// as QuickTime Movie files. Chunk duration can influence the granularity of
// the I/O the system performs when reading a media file; for example, during
// playback. A larger chunk duration can result in fewer reads from disk, at
// the potential expense of a higher memory footprint.
// 
// A chunk contains one or more samples. The total duration of the samples in
// a chunk is no greater than the preferred chunk duration, or the duration of
// a single sample if the sample’s duration is greater than this preferred
// chunk duration.
// 
// The default value is [invalid], which means that the input chooses an
// appropriate default value. It’s an error to set a chunk duration that’s
// negative or nonnumeric.
// 
// You can’t set this value after writing starts.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/preferredMediaChunkDuration
func (a AVAssetWriterInput) PreferredMediaChunkDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("preferredMediaChunkDuration"))
	return objectivec.Object{ID: rv}
}
func (a AVAssetWriterInput) SetPreferredMediaChunkDuration(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreferredMediaChunkDuration:"), value)
}
// The base URL sample references are relative to.
//
// # Discussion
// 
// This property is only valid for file types that support writing sample
// references, such as QuickTime files. If the system resolves the value of
// this property to an absolute URL, the sample references it appends are
// relative to this URL. The URL must point to a location that’s in a
// directory that’s a parent of the sample reference location.
// 
// For example, setting the value of this property to
// `///User/johnappleseed/Movies/` and appending sample buffers with the
// [KCMSampleBufferAttachmentKey_SampleReferenceURL] attachment set to
// `///User/johnappleseed/Movies/data/movie1.Mov()` writes a sample reference
// of `data/movie1.Mov()` to the movie.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/sampleReferenceBaseURL
func (a AVAssetWriterInput) SampleReferenceBaseURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sampleReferenceBaseURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a AVAssetWriterInput) SetSampleReferenceBaseURL(value foundation.INSURL) {
	objc.Send[struct{}](a.ID, objc.Sel("setSampleReferenceBaseURL:"), value)
}
// Specifies how the input lays out and interleaves media data.
//
// # Discussion
// 
// Use this property to optimize tracks that contain a small amount of data
// that you need all at once, such as chapter tracks. Setting the value to
// [beforeMainMediaDataNotInterleaved] causes the asset writer to try to write
// the media data for this track before interleaved inputs. For file types
// that support preloading media data, such as QuickTime files, setting this
// value also writes an indicator to the file to preload its media data.
// 
// Set the value to [interleavedWithMainMediaData] for tracks whose media data
// you need only as it approaches its presentation time, or when multiple
// inputs exist that supply media data that plays concurrently.
// 
// You can’t set this value after writing starts.
//
// [beforeMainMediaDataNotInterleaved]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/MediaDataLocation-swift.struct/beforeMainMediaDataNotInterleaved
// [interleavedWithMainMediaData]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/MediaDataLocation-swift.struct/interleavedWithMainMediaData
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaDataLocation-swift.property
func (a AVAssetWriterInput) MediaDataLocation() AVAssetWriterInputMediaDataLocation {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaDataLocation"))
	return AVAssetWriterInputMediaDataLocation(foundation.NSStringFromID(rv).String())
}
func (a AVAssetWriterInput) SetMediaDataLocation(value AVAssetWriterInputMediaDataLocation) {
	objc.Send[struct{}](a.ID, objc.Sel("setMediaDataLocation:"), objc.String(string(value)))
}
// A Boolean value that indicates whether the input may perform multiple
// passes over appended media data.
//
// # Discussion
// 
// When the value for this property is [true], configure your source media
// data for random access. After appending the media data for the current
// pass, as specified by the [CurrentPassDescription] property, call
// [MarkCurrentPassAsFinished] so the system can determine whether it needs to
// perform additional passes. The system may perform only the initial pass if
// it determines there’s no benefit to performing multiple passes.
// 
// When the value for this property is [false], your source for media data
// only needs to support sequential access. In this case, append all of the
// source media one time and call [MarkAsFinished].
// 
// The default value is [false]. Currently the only way for this property to
// become [true] is when the value of [PerformsMultiPassEncodingIfSupported]
// is [true]. The final value is available after you call [startWriting()].
// 
// This property is key-value observable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [startWriting()]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startWriting()
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/canPerformMultiplePasses
func (a AVAssetWriterInput) CanPerformMultiplePasses() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("canPerformMultiplePasses"))
	return rv
}
// An object that describes the requirements for the current pass.
//
// # Discussion
// 
// If the value of this property is `nil`, call the asset writer input’s
// [MarkAsFinished] method because there are no more requests to fulfill.
// 
// During the first pass, the request contains a single time range value, from
// zero to positive infinity, that indicates to append all media from the
// source. This condition is also true when [CanPerformMultiplePasses] is
// [false], in which case the asset writer only performs a single pass.
// 
// The value of this property is `nil` before you call [startWriting()] on the
// containing asset writer. It transitions to an initial non-`nil` value
// during the call to [startWriting()], and changes only after a call to
// [MarkCurrentPassAsFinished]. You can use the
// [RespondToEachPassDescriptionOnQueueUsingBlock] to have the system call you
// at the beginning of each pass.
// 
// This property is key-value observable. The system doesn’t notify an
// observer on a specific thread.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [startWriting()]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startWriting()
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/currentPassDescription
func (a AVAssetWriterInput) CurrentPassDescription() IAVAssetWriterInputPassDescription {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentPassDescription"))
	return AVAssetWriterInputPassDescriptionFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the input attempts to encode the
// source media data using multiple passes.
//
// # Discussion
// 
// An asset writer input may be able to achieve higher quality and lower data
// rate by performing multiple passes over the source media. It does this by
// analyzing the media data it appends and reencodes certain segments with
// different parameters. To perform the reencoding, you must append the media
// data for the segments again. See the [CurrentPassDescription] property and
// [MarkCurrentPassAsFinished] method for the means by which the input
// nominates segments to reappend.
// 
// When the value of this property is [true], the value of
// [isReadyForMoreMediaData] for other inputs attached to the same asset
// writer may be [false] more often and for longer periods of time. In
// particular, the value of [isReadyForMoreMediaData] for inputs that don’t
// perform multiple passes may start as [false] after calling the method asset
// writer’s [startWriting()] method, and may not change to [true] until
// after all multipass inputs complete their final pass.
// 
// When the value of this property is [true], the input may store data in one
// or more temporary files before writing compressed samples to the output
// file. Use the asset writer’s [DirectoryForTemporaryFiles] property if you
// need to specify the location of temporary file writing.
// 
// The default value is [false], which means that no additional analysis
// occurs and it doesn’t reencode segments. Not all asset writer input
// configurations benefit from performing multiple passes over the source
// media. To determine whether the selected encoder can perform multiple
// passes, query the value of [CanPerformMultiplePasses] after calling
// [startWriting()].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [isReadyForMoreMediaData]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/isReadyForMoreMediaData
// [startWriting()]: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/startWriting()
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/performsMultiPassEncodingIfSupported
func (a AVAssetWriterInput) PerformsMultiPassEncodingIfSupported() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("performsMultiPassEncodingIfSupported"))
	return rv
}
func (a AVAssetWriterInput) SetPerformsMultiPassEncodingIfSupported(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setPerformsMultiPassEncodingIfSupported:"), value)
}
// The media type of the samples that the input accepts.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/mediaType
func (a AVAssetWriterInput) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// The settings to use for encoding media data you append to the output.
//
// # Discussion
// 
// A value of `nil` indicates that the input passes the samples through to the
// output without reencoding them.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/outputSettings
func (a AVAssetWriterInput) OutputSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A hint about the format of the sample buffers to append to the input.
//
// # Discussion
// 
// An input may use this hint to fill in missing output settings or perform
// additional upfront validation of samples.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/sourceFormatHint
func (a AVAssetWriterInput) SourceFormatHint() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sourceFormatHint"))
	return objectivec.Object{ID: rv}
}

