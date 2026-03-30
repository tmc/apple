// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetReaderAudioMixOutput] class.
var (
	_AVAssetReaderAudioMixOutputClass     AVAssetReaderAudioMixOutputClass
	_AVAssetReaderAudioMixOutputClassOnce sync.Once
)

func getAVAssetReaderAudioMixOutputClass() AVAssetReaderAudioMixOutputClass {
	_AVAssetReaderAudioMixOutputClassOnce.Do(func() {
		_AVAssetReaderAudioMixOutputClass = AVAssetReaderAudioMixOutputClass{class: objc.GetClass("AVAssetReaderAudioMixOutput")}
	})
	return _AVAssetReaderAudioMixOutputClass
}

// GetAVAssetReaderAudioMixOutputClass returns the class object for AVAssetReaderAudioMixOutput.
func GetAVAssetReaderAudioMixOutputClass() AVAssetReaderAudioMixOutputClass {
	return getAVAssetReaderAudioMixOutputClass()
}

type AVAssetReaderAudioMixOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetReaderAudioMixOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetReaderAudioMixOutputClass) Alloc() AVAssetReaderAudioMixOutput {
	rv := objc.Send[AVAssetReaderAudioMixOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that reads audio samples that result from mixing audio from one
// or more tracks.
//
// # Overview
//
// Read audio data that you mix from one or more asset tracks by adding an
// audio mix output to an asset reader. You can read the samples in their
// stored format or you can convert them to an alternative format.
//
// # Creating an audio mix output
//
//   - [AVAssetReaderAudioMixOutput.InitWithAudioTracksAudioSettings]: Creates an object that reads mixed audio from the specified audio tracks.
//
// # Configuring audio settings
//
//   - [AVAssetReaderAudioMixOutput.AudioMix]: The audio mix to use with this output.
//   - [AVAssetReaderAudioMixOutput.SetAudioMix]
//   - [AVAssetReaderAudioMixOutput.AudioTimePitchAlgorithm]: The processing algorithm to use for scaled audio edits.
//   - [AVAssetReaderAudioMixOutput.SetAudioTimePitchAlgorithm]
//
// # Inspecting an output
//
//   - [AVAssetReaderAudioMixOutput.AudioTracks]: The tracks from which the output reads audio.
//   - [AVAssetReaderAudioMixOutput.AudioSettings]: The audio settings that the output uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput
type AVAssetReaderAudioMixOutput struct {
	AVAssetReaderOutput
}

// AVAssetReaderAudioMixOutputFromID constructs a [AVAssetReaderAudioMixOutput] from an objc.ID.
//
// An object that reads audio samples that result from mixing audio from one
// or more tracks.
func AVAssetReaderAudioMixOutputFromID(id objc.ID) AVAssetReaderAudioMixOutput {
	return AVAssetReaderAudioMixOutput{AVAssetReaderOutput: AVAssetReaderOutputFromID(id)}
}

// NOTE: AVAssetReaderAudioMixOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetReaderAudioMixOutput] class.
//
// # Creating an audio mix output
//
//   - [IAVAssetReaderAudioMixOutput.InitWithAudioTracksAudioSettings]: Creates an object that reads mixed audio from the specified audio tracks.
//
// # Configuring audio settings
//
//   - [IAVAssetReaderAudioMixOutput.AudioMix]: The audio mix to use with this output.
//   - [IAVAssetReaderAudioMixOutput.SetAudioMix]
//   - [IAVAssetReaderAudioMixOutput.AudioTimePitchAlgorithm]: The processing algorithm to use for scaled audio edits.
//   - [IAVAssetReaderAudioMixOutput.SetAudioTimePitchAlgorithm]
//
// # Inspecting an output
//
//   - [IAVAssetReaderAudioMixOutput.AudioTracks]: The tracks from which the output reads audio.
//   - [IAVAssetReaderAudioMixOutput.AudioSettings]: The audio settings that the output uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput
type IAVAssetReaderAudioMixOutput interface {
	IAVAssetReaderOutput

	// Topic: Creating an audio mix output

	// Creates an object that reads mixed audio from the specified audio tracks.
	InitWithAudioTracksAudioSettings(audioTracks []AVAssetTrack, audioSettings foundation.INSDictionary) AVAssetReaderAudioMixOutput

	// Topic: Configuring audio settings

	// The audio mix to use with this output.
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	// The processing algorithm to use for scaled audio edits.
	AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm)

	// Topic: Inspecting an output

	// The tracks from which the output reads audio.
	AudioTracks() []AVAssetTrack
	// The audio settings that the output uses.
	AudioSettings() foundation.INSDictionary
}

// Init initializes the instance.
func (a AVAssetReaderAudioMixOutput) Init() AVAssetReaderAudioMixOutput {
	rv := objc.Send[AVAssetReaderAudioMixOutput](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetReaderAudioMixOutput) Autorelease() AVAssetReaderAudioMixOutput {
	rv := objc.Send[AVAssetReaderAudioMixOutput](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderAudioMixOutput creates a new AVAssetReaderAudioMixOutput instance.
func NewAVAssetReaderAudioMixOutput() AVAssetReaderAudioMixOutput {
	class := getAVAssetReaderAudioMixOutputClass()
	rv := objc.Send[AVAssetReaderAudioMixOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that reads mixed audio from the specified audio tracks.
//
// audioTracks: An array of track objects of type [audio] from which to source the sample
// buffers to mix.
//
// audioSettings: Optional audio settings to use for audio output. Pass `nil` to receive the
// decoded samples in an uncompressed format. To determine the specific
// format, examine the value of the sample buffer’s [formatDescription]
// property.
//
// For non-`nil` audio settings, the dictionary must contain values for the
// [Linear PCM format settings] keys. The output doesn’t support the
// [AVSampleRateConverterAudioQualityKey] constant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/init(audioTracks:audioSettings:)
//
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [Linear PCM format settings]: https://developer.apple.com/documentation/AVFoundation/linear-pcm-format-settings
// [formatDescription]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/formatDescription
func NewAssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []AVAssetTrack, audioSettings foundation.INSDictionary) AVAssetReaderAudioMixOutput {
	instance := getAVAssetReaderAudioMixOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAudioTracks:audioSettings:"), objectivec.IObjectSliceToNSArray(audioTracks), audioSettings)
	return AVAssetReaderAudioMixOutputFromID(rv)
}

// Creates an object that reads mixed audio from the specified audio tracks.
//
// audioTracks: An array of track objects of type [audio] from which to source the sample
// buffers to mix.
//
// audioSettings: Optional audio settings to use for audio output. Pass `nil` to receive the
// decoded samples in an uncompressed format. To determine the specific
// format, examine the value of the sample buffer’s [formatDescription]
// property.
//
// For non-`nil` audio settings, the dictionary must contain values for the
// [Linear PCM format settings] keys. The output doesn’t support the
// [AVSampleRateConverterAudioQualityKey] constant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/init(audioTracks:audioSettings:)
//
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [Linear PCM format settings]: https://developer.apple.com/documentation/AVFoundation/linear-pcm-format-settings
// [formatDescription]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/formatDescription
func (a AVAssetReaderAudioMixOutput) InitWithAudioTracksAudioSettings(audioTracks []AVAssetTrack, audioSettings foundation.INSDictionary) AVAssetReaderAudioMixOutput {
	rv := objc.Send[AVAssetReaderAudioMixOutput](a.ID, objc.Sel("initWithAudioTracks:audioSettings:"), objectivec.IObjectSliceToNSArray(audioTracks), audioSettings)
	return rv
}

// Creates an object that reads mixed audio from the specified audio tracks.
//
// audioTracks: An array of track objects of type [audio] from which to source the sample
// buffers to mix.
//
// audioSettings: Optional audio settings to use for audio output. Pass `nil` to receive the
// decoded samples in an uncompressed format. To determine the specific
// format, examine the value of the sample buffer’s [formatDescription]
// property.
//
// For non-`nil` audio settings, the dictionary must contain values for the
// [Linear PCM format settings] keys. The output doesn’t support the
// [AVSampleRateConverterAudioQualityKey] constant.
//
// # Return Value
//
// A new audio mix output, or `nil` if initialization fails.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/assetReaderAudioMixOutputWithAudioTracks:audioSettings:
//
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [Linear PCM format settings]: https://developer.apple.com/documentation/AVFoundation/linear-pcm-format-settings
// [formatDescription]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer/formatDescription
func (_AVAssetReaderAudioMixOutputClass AVAssetReaderAudioMixOutputClass) AssetReaderAudioMixOutputWithAudioTracksAudioSettings(audioTracks []AVAssetTrack, audioSettings foundation.INSDictionary) AVAssetReaderAudioMixOutput {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetReaderAudioMixOutputClass.class), objc.Sel("assetReaderAudioMixOutputWithAudioTracks:audioSettings:"), objectivec.IObjectSliceToNSArray(audioTracks), audioSettings)
	return AVAssetReaderAudioMixOutputFromID(rv)
}

// The audio mix to use with this output.
//
// # Discussion
//
// Use an audio mix to specify how an audio track’s volume changes over the
// media’s timeline.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioMix
func (a AVAssetReaderAudioMixOutput) AudioMix() IAVAudioMix {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioMix"))
	return AVAudioMixFromID(objc.ID(rv))
}
func (a AVAssetReaderAudioMixOutput) SetAudioMix(value IAVAudioMix) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioMix:"), value)
}

// The processing algorithm to use for scaled audio edits.
//
// # Discussion
//
// See [Time pitch algorithm settings] for possible values. The system throws
// an exception if you set this property to a value other than one of the
// defined constants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioTimePitchAlgorithm
//
// [Time pitch algorithm settings]: https://developer.apple.com/documentation/AVFoundation/time-pitch-algorithm-settings
func (a AVAssetReaderAudioMixOutput) AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioTimePitchAlgorithm"))
	return AVAudioTimePitchAlgorithm(foundation.NSStringFromID(rv).String())
}
func (a AVAssetReaderAudioMixOutput) SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm) {
	objc.Send[struct{}](a.ID, objc.Sel("setAudioTimePitchAlgorithm:"), objc.String(string(value)))
}

// The tracks from which the output reads audio.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioTracks
func (a AVAssetReaderAudioMixOutput) AudioTracks() []AVAssetTrack {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("audioTracks"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetTrack {
		return AVAssetTrackFromID(id)
	})
}

// The audio settings that the output uses.
//
// # Discussion
//
// The dictionary must contain values for the keys in [Linear PCM format
// settings].
//
// Setting the property value to `nil` indicates that the output returns audio
// samples in an uncompressed format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderAudioMixOutput/audioSettings
//
// [Linear PCM format settings]: https://developer.apple.com/documentation/AVFoundation/linear-pcm-format-settings
func (a AVAssetReaderAudioMixOutput) AudioSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("audioSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
