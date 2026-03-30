// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoComposition] class.
var (
	_AVVideoCompositionClass     AVVideoCompositionClass
	_AVVideoCompositionClassOnce sync.Once
)

func getAVVideoCompositionClass() AVVideoCompositionClass {
	_AVVideoCompositionClassOnce.Do(func() {
		_AVVideoCompositionClass = AVVideoCompositionClass{class: objc.GetClass("AVVideoComposition")}
	})
	return _AVVideoCompositionClass
}

// GetAVVideoCompositionClass returns the class object for AVVideoComposition.
func GetAVVideoCompositionClass() AVVideoCompositionClass {
	return getAVVideoCompositionClass()
}

type AVVideoCompositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVideoCompositionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoCompositionClass) Alloc() AVVideoComposition {
	rv := objc.Send[AVVideoComposition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes how to compose video frames at particular points
// in time.
//
// # Overview
//
// If you use the built-in video compositor, the instructions a video
// composition contain can specify a spatial transformation, an opacity value,
// and a cropping rectangle for each video source. This values can vary over
// time by applying linear ramping functions.
//
// You can create a custom video compositor by implementing the
// [AVVideoCompositing] protocol. The system provides the custom video
// compositor with pixel buffers for each of its video sources during
// playback, and can perform arbitrary graphical operations on them to produce
// visual output.
//
// # Inspecting the video composition
//
//   - [AVVideoComposition.RenderSize]: The size at which the video composition should render.
//   - [AVVideoComposition.RenderScale]: The scale at which the video composition should render.
//   - [AVVideoComposition.FrameDuration]: A time interval for which the video composition should render composed video frames.
//   - [AVVideoComposition.AnimationTool]: A video composition tool to use with Core Animation in offline rendering.
//   - [AVVideoComposition.ColorPrimaries]: The color primaries used for video composition.
//   - [AVVideoComposition.ColorTransferFunction]: The transfer function used for video composition.
//   - [AVVideoComposition.ColorYCbCrMatrix]: The YCbCr matrix used for video composition.
//   - [AVVideoComposition.CustomVideoCompositorClass]: A custom compositor class to use.
//
// # Validating the time range
//
//   - [AVVideoComposition.IsValidForTracksAssetDurationTimeRangeValidationDelegate]: Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
//
// # Reading instructions
//
//   - [AVVideoComposition.Instructions]: The video composition instructions.
//
// # Identifying source tracks
//
//   - [AVVideoComposition.SourceTrackIDForFrameTiming]: An identifier of the source track from which the video composition derives frame timing.
//
// # Configuring HDR metadata
//
//   - [AVVideoComposition.PerFrameHDRDisplayMetadataPolicy]: The policy for display of HDR display metadata on the rendered frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition
type AVVideoComposition struct {
	objectivec.Object
}

// AVVideoCompositionFromID constructs a [AVVideoComposition] from an objc.ID.
//
// An object that describes how to compose video frames at particular points
// in time.
func AVVideoCompositionFromID(id objc.ID) AVVideoComposition {
	return AVVideoComposition{objectivec.Object{ID: id}}
}

// NOTE: AVVideoComposition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoComposition] class.
//
// # Inspecting the video composition
//
//   - [IAVVideoComposition.RenderSize]: The size at which the video composition should render.
//   - [IAVVideoComposition.RenderScale]: The scale at which the video composition should render.
//   - [IAVVideoComposition.FrameDuration]: A time interval for which the video composition should render composed video frames.
//   - [IAVVideoComposition.AnimationTool]: A video composition tool to use with Core Animation in offline rendering.
//   - [IAVVideoComposition.ColorPrimaries]: The color primaries used for video composition.
//   - [IAVVideoComposition.ColorTransferFunction]: The transfer function used for video composition.
//   - [IAVVideoComposition.ColorYCbCrMatrix]: The YCbCr matrix used for video composition.
//   - [IAVVideoComposition.CustomVideoCompositorClass]: A custom compositor class to use.
//
// # Validating the time range
//
//   - [IAVVideoComposition.IsValidForTracksAssetDurationTimeRangeValidationDelegate]: Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
//
// # Reading instructions
//
//   - [IAVVideoComposition.Instructions]: The video composition instructions.
//
// # Identifying source tracks
//
//   - [IAVVideoComposition.SourceTrackIDForFrameTiming]: An identifier of the source track from which the video composition derives frame timing.
//
// # Configuring HDR metadata
//
//   - [IAVVideoComposition.PerFrameHDRDisplayMetadataPolicy]: The policy for display of HDR display metadata on the rendered frame.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition
type IAVVideoComposition interface {
	objectivec.IObject

	// Topic: Inspecting the video composition

	// The size at which the video composition should render.
	RenderSize() corefoundation.CGSize
	// The scale at which the video composition should render.
	RenderScale() float32
	// A time interval for which the video composition should render composed video frames.
	FrameDuration() coremedia.CMTime
	// A video composition tool to use with Core Animation in offline rendering.
	AnimationTool() IAVVideoCompositionCoreAnimationTool
	// The color primaries used for video composition.
	ColorPrimaries() string
	// The transfer function used for video composition.
	ColorTransferFunction() string
	// The YCbCr matrix used for video composition.
	ColorYCbCrMatrix() string
	// A custom compositor class to use.
	CustomVideoCompositorClass() objc.Class

	// Topic: Validating the time range

	// Indicates whether the time ranges of the composition’s instructions conform to validation requirements.
	IsValidForTracksAssetDurationTimeRangeValidationDelegate(tracks []AVAssetTrack, duration coremedia.CMTime, timeRange coremedia.CMTimeRange, validationDelegate AVVideoCompositionValidationHandling) bool

	// Topic: Reading instructions

	// The video composition instructions.
	Instructions() []objectivec.IObject

	// Topic: Identifying source tracks

	// An identifier of the source track from which the video composition derives frame timing.
	SourceTrackIDForFrameTiming() int32

	// Topic: Configuring HDR metadata

	// The policy for display of HDR display metadata on the rendered frame.
	PerFrameHDRDisplayMetadataPolicy() AVVideoCompositionPerFrameHDRDisplayMetadataPolicy

	// The output buffers of the video composition can be specified with the outputBufferDescription. The value is an array of CMTagCollectionRef objects that describes the output buffers.
	OutputBufferDescription() foundation.INSArray
	// The identifiers of source sample data tracks in the composition that the compositor requires to compose frames.
	SourceSampleDataTrackIDs() []foundation.NSNumber
	// Indicates the spatial configurations that are available to associate with the output of the video composition.
	SpatialVideoConfigurations() []AVSpatialVideoConfiguration
}

// Init initializes the instance.
func (v AVVideoComposition) Init() AVVideoComposition {
	rv := objc.Send[AVVideoComposition](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoComposition) Autorelease() AVVideoComposition {
	rv := objc.Send[AVVideoComposition](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoComposition creates a new AVVideoComposition instance.
func NewAVVideoComposition() AVVideoComposition {
	class := getAVVideoCompositionClass()
	rv := objc.Send[AVVideoComposition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a video composition object configured to present the video tracks
// of the specified asset.
//
// asset: The asset whose configuration matches the intended use of the video
// composition.
//
// # Return Value
//
// A new video composition object.
//
// # Discussion
//
// Apple discourages using this method in iOS 16, tvOS 16, and macOS 13 or
// later. Create a video composition asynchronously using
// [VideoCompositionWithPropertiesOfAssetCompletionHandler] instead.
//
// This method creates the video composition object and configures it with the
// values and instructions suitable for presenting the video tracks of the
// specified asset. The returned object contains instructions that respect the
// spatial properties and time ranges of the specified asset’s video tracks.
// It also configures the object properties in the following way:
//
// - The value of the [FrameDuration] property is short enough to accommodate
// the greatest nominal frame rate value among the asset’s video tracks, as
// indicated by the [nominalFrameRate] property of each track. If all its
// tracks have a nominal frame rate of `0`, it uses a frame rate of 30 frames
// per second, with the frame duration set accordingly. - The value of the
// [RenderSize] property depends on whether the asset is an [AVComposition]
// object. For an [AVComposition], the render size is the composition’s
// [NaturalSize] value, and for other assets, its a size large enough to
// encompass all of its video tracks. - The value of the [RenderScale]
// property is `1.0`. - The value of the [AnimationTool] property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/init(propertiesOf:)
//
// [nominalFrameRate]: https://developer.apple.com/documentation/AVFoundation/AVPartialAsyncProperty/nominalFrameRate
func NewVideoCompositionWithPropertiesOfAsset(asset IAVAsset) AVVideoComposition {
	rv := objc.Send[objc.ID](objc.ID(getAVVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return AVVideoCompositionFromID(rv)
}

// Indicates whether the time ranges of the composition’s instructions
// conform to validation requirements.
//
// tracks: Pass a reference to an asset’s tracks if you wish to validate the track
// IDs of the layer instructions against the asset’s tracks. Pass `nil` to
// skip that validation. This method throws an exception if the tracks
// aren’t all from the same asset.
//
// duration: Pass the asset duration to validate the time ranges of the instructions.
// Pass [invalid] to skip that validation.
//
// timeRange: The composition only validates those instructions with time ranges that
// overlap with the specified time range. To validate all instructions that
// the composition may use for playback or other processing, regardless of
// time range, pass `CMTimeRange(XCUIElementTypeZero,
// XCUIElementTypePositiveInfinity)`.
//
// validationDelegate: A delegate that handles validation requests. May be `nil`.
//
// # Return Value
//
// true if the validation succeeds; otherwise; false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/isValid(for:assetDuration:timeRange:validationDelegate:)
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (v AVVideoComposition) IsValidForTracksAssetDurationTimeRangeValidationDelegate(tracks []AVAssetTrack, duration coremedia.CMTime, timeRange coremedia.CMTimeRange, validationDelegate AVVideoCompositionValidationHandling) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isValidForTracks:assetDuration:timeRange:validationDelegate:"), objectivec.IObjectSliceToNSArray(tracks), duration, timeRange, validationDelegate)
	return rv
}

// Returns a new video composition that’s configured to present the video
// tracks of the specified asset.
//
// asset: An asset to create a video composition for.
//
// completionHandler: A callback the system invokes with the created video composition, or an
// error if a failure occurs.
//
// # Discussion
//
// This method creates the video composition object and configures it with the
// values and instructions suitable for presenting the video tracks of the
// specified asset. The returned object contains instructions that respect the
// spatial properties and time ranges of the specified asset’s video tracks.
// It also configures the object properties in the following way:
//
// - The value of the [FrameDuration] property is short enough to accommodate
// the greatest nominal frame rate value among the asset’s video tracks, as
// indicated by the [nominalFrameRate] property of each track. If all its
// tracks have a nominal frame rate of `0`, it uses a frame rate of 30 frames
// per second, with the frame duration set accordingly. - The value of the
// [RenderSize] property depends on whether the asset is an [AVComposition]
// object. For an [AVComposition], the render size is the composition’s
// [NaturalSize] value, and for other assets, its a size large enough to
// encompass all of its video tracks. - The value of the [RenderScale]
// property is `1.0`. - The value of the [AnimationTool] property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoComposition(withPropertiesOf:completionHandler:)
//
// [nominalFrameRate]: https://developer.apple.com/documentation/AVFoundation/AVPartialAsyncProperty/nominalFrameRate
func (_AVVideoCompositionClass AVVideoCompositionClass) VideoCompositionWithPropertiesOfAssetCompletionHandler(asset IAVAsset, completionHandler AVVideoCompositionErrorHandler) {
	_block1, _ := NewAVVideoCompositionErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_AVVideoCompositionClass.class), objc.Sel("videoCompositionWithPropertiesOfAsset:completionHandler:"), asset, _block1)
}

// Pass-through initializer, for internal use in AVFoundation only
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/videoCompositionWithVideoComposition:
func (_AVVideoCompositionClass AVVideoCompositionClass) VideoCompositionWithVideoComposition(videoComposition IAVVideoComposition) AVVideoComposition {
	rv := objc.Send[objc.ID](objc.ID(_AVVideoCompositionClass.class), objc.Sel("videoCompositionWithVideoComposition:"), videoComposition)
	return AVVideoCompositionFromID(rv)
}

// The size at which the video composition should render.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderSize
func (v AVVideoComposition) RenderSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("renderSize"))
	return corefoundation.CGSize(rv)
}

// The scale at which the video composition should render.
//
// # Discussion
//
// This value must be `1.0` unless you set the composition on an
// [AVPlayerItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/renderScale
func (v AVVideoComposition) RenderScale() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("renderScale"))
	return rv
}

// A time interval for which the video composition should render composed
// video frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/frameDuration
func (v AVVideoComposition) FrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](v.ID, objc.Sel("frameDuration"))
	return coremedia.CMTime(rv)
}

// A video composition tool to use with Core Animation in offline rendering.
//
// # Discussion
//
// This attribute may be `nil`. Set an animation tool if you’re using the
// composition in conjunction with [AVAssetExportSession] for offline
// rendering, rather than with [AVPlayer].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/animationTool
func (v AVVideoComposition) AnimationTool() IAVVideoCompositionCoreAnimationTool {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("animationTool"))
	return AVVideoCompositionCoreAnimationToolFromID(objc.ID(rv))
}

// The color primaries used for video composition.
//
// # Discussion
//
// The default value is `nil`. When the value of this property is `nil`, the
// source’s color primaries are propagated and used. Valid values are those
// suitable for [AVVideoColorPrimariesKey].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorPrimaries
//
// [AVVideoColorPrimariesKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimariesKey
func (v AVVideoComposition) ColorPrimaries() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("colorPrimaries"))
	return foundation.NSStringFromID(rv).String()
}

// The transfer function used for video composition.
//
// # Discussion
//
// The default value is `nil`. When the value of this property is `nil`, the
// source’s transfer function is propagated and used. Valid values are those
// suitable for [AVVideoTransferFunctionKey].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorTransferFunction
//
// [AVVideoTransferFunctionKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunctionKey
func (v AVVideoComposition) ColorTransferFunction() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("colorTransferFunction"))
	return foundation.NSStringFromID(rv).String()
}

// The YCbCr matrix used for video composition.
//
// # Discussion
//
// The default value is `nil`. When the value of this property is `nil`, the
// source’s matrix is propagated and used. Valid values are those suitable
// for [AVVideoYCbCrMatrixKey].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/colorYCbCrMatrix
//
// [AVVideoYCbCrMatrixKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoYCbCrMatrixKey
func (v AVVideoComposition) ColorYCbCrMatrix() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("colorYCbCrMatrix"))
	return foundation.NSStringFromID(rv).String()
}

// A custom compositor class to use.
//
// # Discussion
//
// The default is `nil`, which indicates to use the internal video compositor.
// The `customVideoCompositorClass` must implement the [AVVideoCompositing]
// protocol.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/customVideoCompositorClass
func (v AVVideoComposition) CustomVideoCompositorClass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("customVideoCompositorClass"))
	return rv
}

// The video composition instructions.
//
// # Discussion
//
// The array contains instances of [AVVideoCompositionInstruction]. For the
// first instruction in the array, `timeRange.Start()` must be less than or
// equal to the earliest time for which playback or other processing will be
// attempted (typically `kCMTimeZero`). For subsequent instructions,
// `timeRange.Start()` must be equal to the prior instruction’s end time.
// The end time of the last instruction must be greater than or equal to the
// latest time for which playback or other processing will be attempted
// (typically be the duration of the asset with which the instance of
// [AVVideoComposition] is associated).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/instructions
func (v AVVideoComposition) Instructions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("instructions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// An identifier of the source track from which the video composition derives
// frame timing.
//
// # Discussion
//
// If an empty edit is encountered in the source asset’s track, the
// compositor composes frames as needed up to the frequency specified in
// [FrameDuration] property. Otherwise the frame timing for the video
// composition is derived from the source asset’s track with the
// corresponding ID.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/sourceTrackIDForFrameTiming
func (v AVVideoComposition) SourceTrackIDForFrameTiming() int32 {
	rv := objc.Send[int32](v.ID, objc.Sel("sourceTrackIDForFrameTiming"))
	return rv
}

// The policy for display of HDR display metadata on the rendered frame.
//
// # Discussion
//
// Allows the system to identify situations where it can generate HDR metadata
// and attach it to the rendered video frame.
//
// The default value is [propagate], which indicates the system propagates any
// HDR metadata attached to the composed frame to the rendered video frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/perFrameHDRDisplayMetadataPolicy-swift.property
//
// [propagate]: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/PerFrameHDRDisplayMetadataPolicy-swift.struct/propagate
func (v AVVideoComposition) PerFrameHDRDisplayMetadataPolicy() AVVideoCompositionPerFrameHDRDisplayMetadataPolicy {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("perFrameHDRDisplayMetadataPolicy"))
	return AVVideoCompositionPerFrameHDRDisplayMetadataPolicy(foundation.NSStringFromID(rv).String())
}

// The output buffers of the video composition can be specified with the
// outputBufferDescription. The value is an array of CMTagCollectionRef
// objects that describes the output buffers.
//
// # Discussion
//
// If the video composition will output tagged buffers, the details of those
// buffers should be specified with CMTags. Specifically, the StereoView
// (eyes) and ProjectionKind must be specified. The behavior is undefined if
// the output tagged buffers do not match the outputBufferDescription. The
// default is nil, which means monoscopic output. Note that an empty array is
// not valid. An exception will be thrown if the objects in the array are not
// of type CMTagCollectionRef. Note that tagged buffers are only supported for
// custom compositors.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/outputBufferDescription-3wsar
func (v AVVideoComposition) OutputBufferDescription() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("outputBufferDescription"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// The identifiers of source sample data tracks in the composition that the
// compositor requires to compose frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/sourceSampleDataTrackIDs-3nrgi
func (v AVVideoComposition) SourceSampleDataTrackIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("sourceSampleDataTrackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Indicates the spatial configurations that are available to associate with
// the output of the video composition.
//
// # Discussion
//
// A custom compositor can output spatial video by specifying one of these
// spatial configurations. A spatial configuration with all nil values
// indicates the video is not spatial. A nil spatial configuration also
// indicates the video is not spatial. The value can be nil, which indicates
// the output will not be spatial. NOTE: If this property is not empty, then
// the client must attach one of the spatial configurations in this array to
// all of the pixel buffers, otherwise an exception will be thrown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/spatialVideoConfigurations-2ipps
func (v AVVideoComposition) SpatialVideoConfigurations() []AVSpatialVideoConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("spatialVideoConfigurations"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVSpatialVideoConfiguration {
		return AVSpatialVideoConfigurationFromID(id)
	})
}

// VideoCompositionWithPropertiesOfAssetSync is a synchronous wrapper around [AVVideoComposition.VideoCompositionWithPropertiesOfAssetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (vc AVVideoCompositionClass) VideoCompositionWithPropertiesOfAssetSync(ctx context.Context, asset IAVAsset) (*AVVideoComposition, error) {
	type result struct {
		val *AVVideoComposition
		err error
	}
	done := make(chan result, 1)
	vc.VideoCompositionWithPropertiesOfAssetCompletionHandler(asset, func(val *AVVideoComposition, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
