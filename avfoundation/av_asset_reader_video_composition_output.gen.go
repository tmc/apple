// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetReaderVideoCompositionOutput] class.
var (
	_AVAssetReaderVideoCompositionOutputClass     AVAssetReaderVideoCompositionOutputClass
	_AVAssetReaderVideoCompositionOutputClassOnce sync.Once
)

func getAVAssetReaderVideoCompositionOutputClass() AVAssetReaderVideoCompositionOutputClass {
	_AVAssetReaderVideoCompositionOutputClassOnce.Do(func() {
		_AVAssetReaderVideoCompositionOutputClass = AVAssetReaderVideoCompositionOutputClass{class: objc.GetClass("AVAssetReaderVideoCompositionOutput")}
	})
	return _AVAssetReaderVideoCompositionOutputClass
}

// GetAVAssetReaderVideoCompositionOutputClass returns the class object for AVAssetReaderVideoCompositionOutput.
func GetAVAssetReaderVideoCompositionOutputClass() AVAssetReaderVideoCompositionOutputClass {
	return getAVAssetReaderVideoCompositionOutputClass()
}

type AVAssetReaderVideoCompositionOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetReaderVideoCompositionOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetReaderVideoCompositionOutputClass) Alloc() AVAssetReaderVideoCompositionOutput {
	rv := objc.Send[AVAssetReaderVideoCompositionOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that reads composited video frames from one or more tracks of an
// asset.
//
// # Creating a video composition output
//
//   - [AVAssetReaderVideoCompositionOutput.InitWithVideoTracksVideoSettings]: Creates an object that reads composited video frames from the specified video tracks.
//
// # Configuring video settings
//
//   - [AVAssetReaderVideoCompositionOutput.VideoComposition]: The video composition to use for the output.
//   - [AVAssetReaderVideoCompositionOutput.SetVideoComposition]
//   - [AVAssetReaderVideoCompositionOutput.CustomVideoCompositor]: A custom video compositor for the output.
//
// # Inspecting an output
//
//   - [AVAssetReaderVideoCompositionOutput.VideoTracks]: The tracks from which the output reads the composited video.
//   - [AVAssetReaderVideoCompositionOutput.VideoSettings]: The video settings that the output uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput
type AVAssetReaderVideoCompositionOutput struct {
	AVAssetReaderOutput
}

// AVAssetReaderVideoCompositionOutputFromID constructs a [AVAssetReaderVideoCompositionOutput] from an objc.ID.
//
// An object that reads composited video frames from one or more tracks of an
// asset.
func AVAssetReaderVideoCompositionOutputFromID(id objc.ID) AVAssetReaderVideoCompositionOutput {
	return AVAssetReaderVideoCompositionOutput{AVAssetReaderOutput: AVAssetReaderOutputFromID(id)}
}

// NOTE: AVAssetReaderVideoCompositionOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetReaderVideoCompositionOutput] class.
//
// # Creating a video composition output
//
//   - [IAVAssetReaderVideoCompositionOutput.InitWithVideoTracksVideoSettings]: Creates an object that reads composited video frames from the specified video tracks.
//
// # Configuring video settings
//
//   - [IAVAssetReaderVideoCompositionOutput.VideoComposition]: The video composition to use for the output.
//   - [IAVAssetReaderVideoCompositionOutput.SetVideoComposition]
//   - [IAVAssetReaderVideoCompositionOutput.CustomVideoCompositor]: A custom video compositor for the output.
//
// # Inspecting an output
//
//   - [IAVAssetReaderVideoCompositionOutput.VideoTracks]: The tracks from which the output reads the composited video.
//   - [IAVAssetReaderVideoCompositionOutput.VideoSettings]: The video settings that the output uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput
type IAVAssetReaderVideoCompositionOutput interface {
	IAVAssetReaderOutput

	// Topic: Creating a video composition output

	// Creates an object that reads composited video frames from the specified video tracks.
	InitWithVideoTracksVideoSettings(videoTracks []AVAssetTrack, videoSettings foundation.INSDictionary) AVAssetReaderVideoCompositionOutput

	// Topic: Configuring video settings

	// The video composition to use for the output.
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
	// A custom video compositor for the output.
	CustomVideoCompositor() AVVideoCompositing

	// Topic: Inspecting an output

	// The tracks from which the output reads the composited video.
	VideoTracks() []AVAssetTrack
	// The video settings that the output uses.
	VideoSettings() foundation.INSDictionary
}

// Init initializes the instance.
func (a AVAssetReaderVideoCompositionOutput) Init() AVAssetReaderVideoCompositionOutput {
	rv := objc.Send[AVAssetReaderVideoCompositionOutput](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetReaderVideoCompositionOutput) Autorelease() AVAssetReaderVideoCompositionOutput {
	rv := objc.Send[AVAssetReaderVideoCompositionOutput](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetReaderVideoCompositionOutput creates a new AVAssetReaderVideoCompositionOutput instance.
func NewAVAssetReaderVideoCompositionOutput() AVAssetReaderVideoCompositionOutput {
	class := getAVAssetReaderVideoCompositionOutputClass()
	rv := objc.Send[AVAssetReaderVideoCompositionOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that reads composited video frames from the specified
// video tracks.
//
// videoTracks: An array of asset tracks from which to read video frames for compositing.
// The media type of each track must be [video].
//
// videoSettings: Specifying a `nil` value configures the output to return samples in an
// uncompressed format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/init(videoTracks:videoSettings:)
//
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
func NewAssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []AVAssetTrack, videoSettings foundation.INSDictionary) AVAssetReaderVideoCompositionOutput {
	instance := getAVAssetReaderVideoCompositionOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVideoTracks:videoSettings:"), objectivec.IObjectSliceToNSArray(videoTracks), videoSettings)
	return AVAssetReaderVideoCompositionOutputFromID(rv)
}

// Creates an object that reads composited video frames from the specified
// video tracks.
//
// videoTracks: An array of asset tracks from which to read video frames for compositing.
// The media type of each track must be [video].
//
// videoSettings: Specifying a `nil` value configures the output to return samples in an
// uncompressed format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/init(videoTracks:videoSettings:)
//
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
func (a AVAssetReaderVideoCompositionOutput) InitWithVideoTracksVideoSettings(videoTracks []AVAssetTrack, videoSettings foundation.INSDictionary) AVAssetReaderVideoCompositionOutput {
	rv := objc.Send[AVAssetReaderVideoCompositionOutput](a.ID, objc.Sel("initWithVideoTracks:videoSettings:"), objectivec.IObjectSliceToNSArray(videoTracks), videoSettings)
	return rv
}

// Returns a new object that reads composited video from the specified video
// tracks.
//
// videoTracks: An array of asset tracks from which the created object should read video
// frames for compositing. The media type of each track must be [video].
//
// videoSettings: A dictionary of video settings to use for sample output, or `nil` if you
// want to receive decoded samples in a convenient uncompressed format, with
// properties determined according to the properties of the specified video
// tracks.
//
// You use keys from [CVPixelBuffer], depending on the output format you want.
//
// # Return Value
//
// A new video composition output, or `nil` if initialization fails.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/assetReaderVideoCompositionOutputWithVideoTracks:videoSettings:
//
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
func (_AVAssetReaderVideoCompositionOutputClass AVAssetReaderVideoCompositionOutputClass) AssetReaderVideoCompositionOutputWithVideoTracksVideoSettings(videoTracks []AVAssetTrack, videoSettings foundation.INSDictionary) AVAssetReaderVideoCompositionOutput {
	rv := objc.Send[objc.ID](objc.ID(_AVAssetReaderVideoCompositionOutputClass.class), objc.Sel("assetReaderVideoCompositionOutputWithVideoTracks:videoSettings:"), objectivec.IObjectSliceToNSArray(videoTracks), videoSettings)
	return AVAssetReaderVideoCompositionOutputFromID(rv)
}

// The video composition to use for the output.
//
// # Discussion
//
// The value is an [AVVideoComposition] object that specifies the visual
// arrangement of video frames read from each source track over the timeline
// of the source asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoComposition
func (a AVAssetReaderVideoCompositionOutput) VideoComposition() IAVVideoComposition {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoComposition"))
	return AVVideoCompositionFromID(objc.ID(rv))
}
func (a AVAssetReaderVideoCompositionOutput) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[struct{}](a.ID, objc.Sel("setVideoComposition:"), value)
}

// A custom video compositor for the output.
//
// # Discussion
//
// This property is `nil` if there isn’t a custom video compositor, or if
// the internal video compositor is in use.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/customVideoCompositor
func (a AVAssetReaderVideoCompositionOutput) CustomVideoCompositor() AVVideoCompositing {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("customVideoCompositor"))
	return AVVideoCompositingObjectFromID(rv)
}

// The tracks from which the output reads the composited video.
//
// # Discussion
//
// The array contains [AVAssetTrack] objects owned by the target asset
// reader’s asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoTracks
func (a AVAssetReaderVideoCompositionOutput) VideoTracks() []AVAssetTrack {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("videoTracks"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetTrack {
		return AVAssetTrackFromID(id)
	})
}

// The video settings that the output uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReaderVideoCompositionOutput/videoSettings
func (a AVAssetReaderVideoCompositionOutput) VideoSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("videoSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
