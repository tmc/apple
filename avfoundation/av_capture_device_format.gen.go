// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureDeviceFormat] class.
var (
	_AVCaptureDeviceFormatClass     AVCaptureDeviceFormatClass
	_AVCaptureDeviceFormatClassOnce sync.Once
)

func getAVCaptureDeviceFormatClass() AVCaptureDeviceFormatClass {
	_AVCaptureDeviceFormatClassOnce.Do(func() {
		_AVCaptureDeviceFormatClass = AVCaptureDeviceFormatClass{class: objc.GetClass("AVCaptureDeviceFormat")}
	})
	return _AVCaptureDeviceFormatClass
}

// GetAVCaptureDeviceFormatClass returns the class object for AVCaptureDeviceFormat.
func GetAVCaptureDeviceFormatClass() AVCaptureDeviceFormatClass {
	return getAVCaptureDeviceFormatClass()
}

type AVCaptureDeviceFormatClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeviceFormatClass) Alloc() AVCaptureDeviceFormat {
	rv := objc.Send[AVCaptureDeviceFormat](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A class that defines media formats and capture settings that capture
// devices support.
//
// # Overview
// 
// A format object provides information about a media capture format to use
// with a capture device, such as video frame rates and zoom factors.
// 
// You can find more information about a capture format using its associated
// Core Media format description (see [CMFormatDescription]), available using
// the [AVCaptureDeviceFormat.FormatDescription] property.
//
// [CMFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
//
// # Determining spatial capture support
//
//   - [AVCaptureDeviceFormat.SpatialVideoCaptureSupported]: A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// # Determining background replacement support
//
//   - [AVCaptureDeviceFormat.BackgroundReplacementSupported]: A Boolean value that indicates whether the format supports background replacement.
//   - [AVCaptureDeviceFormat.VideoFrameRateRangeForBackgroundReplacement]: The minimum and maximum frame rates available when Background Replacement is active.
//
// # Determining video capture support
//
//   - [AVCaptureDeviceFormat.AutoVideoFrameRateSupported]: A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//   - [AVCaptureDeviceFormat.VideoSupportedFrameRateRanges]: A list of frame rate ranges that a format supports.
//
// # Determining reaction effects support
//
//   - [AVCaptureDeviceFormat.ReactionEffectsSupported]: A Boolean value that indicates whether the device supports reaction effects.
//   - [AVCaptureDeviceFormat.VideoFrameRateRangeForReactionEffectsInProgress]: Indicates the minimum and maximum frame rates available when a reaction effect runs.
//
// # Determining supported media formats
//
//   - [AVCaptureDeviceFormat.MediaType]: A constant describing the media type of an [AVCaptureDevice] active or supported format.
//   - [AVCaptureDeviceFormat.FormatDescription]: An object describing the capture format.
//
// # Determining photo quality
//
//   - [AVCaptureDeviceFormat.HighPhotoQualitySupported]: A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// # Determining zoom capabilities
//
//   - [AVCaptureDeviceFormat.ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported]: A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// # Determining the auto focus system
//
//   - [AVCaptureDeviceFormat.AutoFocusSystem]: The auto focus system the format uses.
//
// # Determining Cinematic video support
//
//   - [AVCaptureDeviceFormat.CinematicVideoCaptureSupported]: Indicates whether the format supports Cinematic Video capture.
//   - [AVCaptureDeviceFormat.DefaultSimulatedAperture]: Default shallow depth of field simulated aperture.
//   - [AVCaptureDeviceFormat.MinSimulatedAperture]: Minimum supported shallow depth of field simulated aperture.
//   - [AVCaptureDeviceFormat.MaxSimulatedAperture]: Maximum supported shallow depth of field simulated aperture.
//   - [AVCaptureDeviceFormat.VideoMaxZoomFactorForCinematicVideo]: Indicates the maximum zoom factor available for the [videoZoomFactor](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor>) property when Cinematic Video capture is enabled on the device input.
//   - [AVCaptureDeviceFormat.VideoMinZoomFactorForCinematicVideo]: Indicates the minimum zoom factor available for the [videoZoomFactor](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor>) property when Cinematic Video capture is enabled on the device input.
//   - [AVCaptureDeviceFormat.VideoFrameRateRangeForCinematicVideo]: Indicates the minimum / maximum frame rates available when Cinematic Video capture is enabled on the device input.
//
// # Determining lens smudge detection support
//
//   - [AVCaptureDeviceFormat.CameraLensSmudgeDetectionSupported]: Whether camera lens smudge detection is supported.
//
// # Determining Center Stage support
//
//   - [AVCaptureDeviceFormat.CenterStageSupported]: A Boolean value that indicates whether the format supports Center Stage.
//   - [AVCaptureDeviceFormat.VideoFrameRateRangeForCenterStage]: The range of frame rates available when Center Stage is active.
//   - [AVCaptureDeviceFormat.VideoMinZoomFactorForCenterStage]: The minimum zoom factor available when Center Stage is active.
//   - [AVCaptureDeviceFormat.VideoMaxZoomFactorForCenterStage]: The maximum zoom factor available when Center Stage is active.
//
// # Determining Portrait Effects support
//
//   - [AVCaptureDeviceFormat.PortraitEffectSupported]: A Boolean value that indicates whether the format supports the Portrait Effect feature.
//   - [AVCaptureDeviceFormat.VideoFrameRateRangeForPortraitEffect]: The range of frame rates available when Portrait Effect is active.
//
// # Determining Studio Light support
//
//   - [AVCaptureDeviceFormat.StudioLightSupported]: A Boolean value that indicates whether the format supports Studio Light.
//   - [AVCaptureDeviceFormat.VideoFrameRateRangeForStudioLight]: A value that indicates the minimum and maximum frame rates available when a user enables Studio Light.
//
// # Instance Properties
//
//   - [AVCaptureDeviceFormat.EdgeLightSupported]: Indicates whether the format supports the Edge Light feature.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format
type AVCaptureDeviceFormat struct {
	objectivec.Object
}

// AVCaptureDeviceFormatFromID constructs a [AVCaptureDeviceFormat] from an objc.ID.
//
// A class that defines media formats and capture settings that capture
// devices support.
func AVCaptureDeviceFormatFromID(id objc.ID) AVCaptureDeviceFormat {
	return AVCaptureDeviceFormat{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDeviceFormat adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeviceFormat] class.
//
// # Determining spatial capture support
//
//   - [IAVCaptureDeviceFormat.SpatialVideoCaptureSupported]: A Boolean value that indicates whether the format supports capturing spatial video to a file.
//
// # Determining background replacement support
//
//   - [IAVCaptureDeviceFormat.BackgroundReplacementSupported]: A Boolean value that indicates whether the format supports background replacement.
//   - [IAVCaptureDeviceFormat.VideoFrameRateRangeForBackgroundReplacement]: The minimum and maximum frame rates available when Background Replacement is active.
//
// # Determining video capture support
//
//   - [IAVCaptureDeviceFormat.AutoVideoFrameRateSupported]: A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
//   - [IAVCaptureDeviceFormat.VideoSupportedFrameRateRanges]: A list of frame rate ranges that a format supports.
//
// # Determining reaction effects support
//
//   - [IAVCaptureDeviceFormat.ReactionEffectsSupported]: A Boolean value that indicates whether the device supports reaction effects.
//   - [IAVCaptureDeviceFormat.VideoFrameRateRangeForReactionEffectsInProgress]: Indicates the minimum and maximum frame rates available when a reaction effect runs.
//
// # Determining supported media formats
//
//   - [IAVCaptureDeviceFormat.MediaType]: A constant describing the media type of an [AVCaptureDevice] active or supported format.
//   - [IAVCaptureDeviceFormat.FormatDescription]: An object describing the capture format.
//
// # Determining photo quality
//
//   - [IAVCaptureDeviceFormat.HighPhotoQualitySupported]: A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
//
// # Determining zoom capabilities
//
//   - [IAVCaptureDeviceFormat.ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported]: A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
//
// # Determining the auto focus system
//
//   - [IAVCaptureDeviceFormat.AutoFocusSystem]: The auto focus system the format uses.
//
// # Determining Cinematic video support
//
//   - [IAVCaptureDeviceFormat.CinematicVideoCaptureSupported]: Indicates whether the format supports Cinematic Video capture.
//   - [IAVCaptureDeviceFormat.DefaultSimulatedAperture]: Default shallow depth of field simulated aperture.
//   - [IAVCaptureDeviceFormat.MinSimulatedAperture]: Minimum supported shallow depth of field simulated aperture.
//   - [IAVCaptureDeviceFormat.MaxSimulatedAperture]: Maximum supported shallow depth of field simulated aperture.
//   - [IAVCaptureDeviceFormat.VideoMaxZoomFactorForCinematicVideo]: Indicates the maximum zoom factor available for the [videoZoomFactor](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor>) property when Cinematic Video capture is enabled on the device input.
//   - [IAVCaptureDeviceFormat.VideoMinZoomFactorForCinematicVideo]: Indicates the minimum zoom factor available for the [videoZoomFactor](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor>) property when Cinematic Video capture is enabled on the device input.
//   - [IAVCaptureDeviceFormat.VideoFrameRateRangeForCinematicVideo]: Indicates the minimum / maximum frame rates available when Cinematic Video capture is enabled on the device input.
//
// # Determining lens smudge detection support
//
//   - [IAVCaptureDeviceFormat.CameraLensSmudgeDetectionSupported]: Whether camera lens smudge detection is supported.
//
// # Determining Center Stage support
//
//   - [IAVCaptureDeviceFormat.CenterStageSupported]: A Boolean value that indicates whether the format supports Center Stage.
//   - [IAVCaptureDeviceFormat.VideoFrameRateRangeForCenterStage]: The range of frame rates available when Center Stage is active.
//   - [IAVCaptureDeviceFormat.VideoMinZoomFactorForCenterStage]: The minimum zoom factor available when Center Stage is active.
//   - [IAVCaptureDeviceFormat.VideoMaxZoomFactorForCenterStage]: The maximum zoom factor available when Center Stage is active.
//
// # Determining Portrait Effects support
//
//   - [IAVCaptureDeviceFormat.PortraitEffectSupported]: A Boolean value that indicates whether the format supports the Portrait Effect feature.
//   - [IAVCaptureDeviceFormat.VideoFrameRateRangeForPortraitEffect]: The range of frame rates available when Portrait Effect is active.
//
// # Determining Studio Light support
//
//   - [IAVCaptureDeviceFormat.StudioLightSupported]: A Boolean value that indicates whether the format supports Studio Light.
//   - [IAVCaptureDeviceFormat.VideoFrameRateRangeForStudioLight]: A value that indicates the minimum and maximum frame rates available when a user enables Studio Light.
//
// # Instance Properties
//
//   - [IAVCaptureDeviceFormat.EdgeLightSupported]: Indicates whether the format supports the Edge Light feature.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format
type IAVCaptureDeviceFormat interface {
	objectivec.IObject

	// Topic: Determining spatial capture support

	// A Boolean value that indicates whether the format supports capturing spatial video to a file.
	SpatialVideoCaptureSupported() bool

	// Topic: Determining background replacement support

	// A Boolean value that indicates whether the format supports background replacement.
	BackgroundReplacementSupported() bool
	// The minimum and maximum frame rates available when Background Replacement is active.
	VideoFrameRateRangeForBackgroundReplacement() IAVFrameRateRange

	// Topic: Determining video capture support

	// A Boolean value that Indicates whether the format supports performing automatic video frame rate adjustments.
	AutoVideoFrameRateSupported() bool
	// A list of frame rate ranges that a format supports.
	VideoSupportedFrameRateRanges() []AVFrameRateRange

	// Topic: Determining reaction effects support

	// A Boolean value that indicates whether the device supports reaction effects.
	ReactionEffectsSupported() bool
	// Indicates the minimum and maximum frame rates available when a reaction effect runs.
	VideoFrameRateRangeForReactionEffectsInProgress() IAVFrameRateRange

	// Topic: Determining supported media formats

	// A constant describing the media type of an [AVCaptureDevice] active or supported format.
	MediaType() AVMediaType
	// An object describing the capture format.
	FormatDescription() objectivec.IObject

	// Topic: Determining photo quality

	// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting.
	HighPhotoQualitySupported() bool

	// Topic: Determining zoom capabilities

	// A Boolean value that indicates whether the format supports zoom factors outside the range supported for depth delivery.
	ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool

	// Topic: Determining the auto focus system

	// The auto focus system the format uses.
	AutoFocusSystem() AVCaptureAutoFocusSystem

	// Topic: Determining Cinematic video support

	// Indicates whether the format supports Cinematic Video capture.
	CinematicVideoCaptureSupported() bool
	// Default shallow depth of field simulated aperture.
	DefaultSimulatedAperture() float32
	// Minimum supported shallow depth of field simulated aperture.
	MinSimulatedAperture() float32
	// Maximum supported shallow depth of field simulated aperture.
	MaxSimulatedAperture() float32
	// Indicates the maximum zoom factor available for the [videoZoomFactor](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor>) property when Cinematic Video capture is enabled on the device input.
	VideoMaxZoomFactorForCinematicVideo() float64
	// Indicates the minimum zoom factor available for the [videoZoomFactor](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureDevice/videoZoomFactor>) property when Cinematic Video capture is enabled on the device input.
	VideoMinZoomFactorForCinematicVideo() float64
	// Indicates the minimum / maximum frame rates available when Cinematic Video capture is enabled on the device input.
	VideoFrameRateRangeForCinematicVideo() IAVFrameRateRange

	// Topic: Determining lens smudge detection support

	// Whether camera lens smudge detection is supported.
	CameraLensSmudgeDetectionSupported() bool

	// Topic: Determining Center Stage support

	// A Boolean value that indicates whether the format supports Center Stage.
	CenterStageSupported() bool
	// The range of frame rates available when Center Stage is active.
	VideoFrameRateRangeForCenterStage() IAVFrameRateRange
	// The minimum zoom factor available when Center Stage is active.
	VideoMinZoomFactorForCenterStage() float64
	// The maximum zoom factor available when Center Stage is active.
	VideoMaxZoomFactorForCenterStage() float64

	// Topic: Determining Portrait Effects support

	// A Boolean value that indicates whether the format supports the Portrait Effect feature.
	PortraitEffectSupported() bool
	// The range of frame rates available when Portrait Effect is active.
	VideoFrameRateRangeForPortraitEffect() IAVFrameRateRange

	// Topic: Determining Studio Light support

	// A Boolean value that indicates whether the format supports Studio Light.
	StudioLightSupported() bool
	// A value that indicates the minimum and maximum frame rates available when a user enables Studio Light.
	VideoFrameRateRangeForStudioLight() IAVFrameRateRange

	// Topic: Instance Properties

	// Indicates whether the format supports the Edge Light feature.
	EdgeLightSupported() bool

	// The currently active depth data format of the capture device.
	ActiveDepthDataFormat() IAVCaptureDeviceFormat
	SetActiveDepthDataFormat(value IAVCaptureDeviceFormat)
	// The capture format in use by the device.
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	// The capture formats a device supports.
	Formats() IAVCaptureDeviceFormat
	SetFormats(value IAVCaptureDeviceFormat)
	// The zoom factors at which this device transitions to secondary native resolution modes.
	SecondaryNativeResolutionZoomFactors() []foundation.NSNumber
	// The list of color spaces the format supports for image and video capture.
	SupportedColorSpaces() []foundation.NSNumber
	// The maximum photo dimension this format supports.
	SupportedMaxPhotoDimensions() []foundation.NSValue
	// The zoom ranges that support the delivery of depth data.
	SupportedVideoZoomRangesForDepthDataDelivery() []AVZoomRange
	// The system’s recommended exposure bias range for this device format.
	SystemRecommendedExposureBiasRange() IAVExposureBiasRange
	// The system’s recommended zoom range for this device format.
	SystemRecommendedVideoZoomRange() IAVZoomRange
	// A value that controls the cropping and enlargement of images captured by the device.
	VideoZoomFactor() float64
	SetVideoZoomFactor(value float64)
}

// Init initializes the instance.
func (c AVCaptureDeviceFormat) Init() AVCaptureDeviceFormat {
	rv := objc.Send[AVCaptureDeviceFormat](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeviceFormat) Autorelease() AVCaptureDeviceFormat {
	rv := objc.Send[AVCaptureDeviceFormat](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceFormat creates a new AVCaptureDeviceFormat instance.
func NewAVCaptureDeviceFormat() AVCaptureDeviceFormat {
	class := getAVCaptureDeviceFormatClass()
	rv := objc.Send[AVCaptureDeviceFormat](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the format supports capturing
// spatial video to a file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isSpatialVideoCaptureSupported
func (c AVCaptureDeviceFormat) SpatialVideoCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSpatialVideoCaptureSupported"))
	return rv
}
// A Boolean value that indicates whether the format supports background
// replacement.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isBackgroundReplacementSupported
func (c AVCaptureDeviceFormat) BackgroundReplacementSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBackgroundReplacementSupported"))
	return rv
}
// The minimum and maximum frame rates available when Background Replacement
// is active.
//
// # Discussion
// 
// Devices may support a limited frame rate range when Background Replacement
// is active. If this device format doesn’t support this feature, the value
// of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForBackgroundReplacement
func (c AVCaptureDeviceFormat) VideoFrameRateRangeForBackgroundReplacement() IAVFrameRateRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoFrameRateRangeForBackgroundReplacement"))
	return AVFrameRateRangeFromID(objc.ID(rv))
}
// A Boolean value that Indicates whether the format supports performing
// automatic video frame rate adjustments.
//
// # Discussion
// 
// This property determines whether you can enable a capture device’s
// [AutoVideoFrameRateEnabled] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isAutoVideoFrameRateSupported
func (c AVCaptureDeviceFormat) AutoVideoFrameRateSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoVideoFrameRateSupported"))
	return rv
}
// A list of frame rate ranges that a format supports.
//
// # Discussion
// 
// The value is an array of [AVFrameRateRange] objects, one for each of the
// format’s supported video frame rate ranges.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoSupportedFrameRateRanges
func (c AVCaptureDeviceFormat) VideoSupportedFrameRateRanges() []AVFrameRateRange {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("videoSupportedFrameRateRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVFrameRateRange {
		return AVFrameRateRangeFromID(id)
	})
}
// A Boolean value that indicates whether the device supports reaction
// effects.
//
// # Discussion
// 
// See [ReactionEffectsEnabled] for more information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/reactionEffectsSupported
func (c AVCaptureDeviceFormat) ReactionEffectsSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("reactionEffectsSupported"))
	return rv
}
// Indicates the minimum and maximum frame rates available when a reaction
// effect runs.
//
// # Discussion
// 
// Unlike other video effects, enabling reaction effects doesn’t limit the
// stream’s frame rate because most of the time the system isn’t rendering
// the effect. The frame rate only ramps down when the system renders a
// reaction on the stream.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForReactionEffectsInProgress
func (c AVCaptureDeviceFormat) VideoFrameRateRangeForReactionEffectsInProgress() IAVFrameRateRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoFrameRateRangeForReactionEffectsInProgress"))
	return AVFrameRateRangeFromID(objc.ID(rv))
}
// A constant describing the media type of an [AVCaptureDevice] active or
// supported format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/mediaType
func (c AVCaptureDeviceFormat) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// An object describing the capture format.
//
// # Discussion
// 
// Calling this method doesn’t assume ownership of the returned
// [CMFormatDescription].
//
// [CMFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/formatDescription
func (c AVCaptureDeviceFormat) FormatDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("formatDescription"))
	return objectivec.Object{ID: rv}
}
// A Boolean value that indicates whether this format supports high-quality
// capture with the current quality prioritization setting.
//
// # Discussion
// 
// When this value is [true], the format produces higher image quality when
// selecting a quality prioritization of
// [CapturePhotoQualityPrioritizationBalanced] or
// [CapturePhotoQualityPrioritizationQuality] in comparison to
// [CapturePhotoQualityPrioritizationSpeed].
// 
// High-quality formats adhere to the following rules:
// 
// - Photo requests that prioritize speed produce the fastest image result,
// which makes it a good choice for burst captures. - Photo requests that
// prioritize speed and quality equally produce higher image quality without
// dropping frames if a video recording is underway. - Photo requests that
// prioritize quality produce high-quality images and may cause frame drops if
// a video recording is underway. For maximum backward compatibility, photo
// requests on high photo quality formats only cause video frame drops if your
// app links against iOS 15 or later.
// 
// Formats that don’t support high photo quality produce the same image
// quality regardless of the current [PhotoQualityPrioritization] setting.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isHighPhotoQualitySupported
func (c AVCaptureDeviceFormat) HighPhotoQualitySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isHighPhotoQualitySupported"))
	return rv
}
// A Boolean value that indicates whether the format supports zoom factors
// outside the range supported for depth delivery.
//
// # Discussion
// 
// Setting a zoom factor outside the range defined by the
// [supportedVideoZoomFactorsForDepthDataDelivery] property results in the
// system suspending depth data delivery. It resumes delivery when you set the
// zoom factor back to a supported value.
//
// [supportedVideoZoomFactorsForDepthDataDelivery]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedVideoZoomFactorsForDepthDataDelivery
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported
func (c AVCaptureDeviceFormat) ZoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("zoomFactorsOutsideOfVideoZoomRangesForDepthDeliverySupported"))
	return rv
}
// The auto focus system the format uses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/autoFocusSystem-swift.property
func (c AVCaptureDeviceFormat) AutoFocusSystem() AVCaptureAutoFocusSystem {
	rv := objc.Send[AVCaptureAutoFocusSystem](c.ID, objc.Sel("autoFocusSystem"))
	return AVCaptureAutoFocusSystem(rv)
}
// Indicates whether the format supports Cinematic Video capture.
//
// # Discussion
// 
// This property returns `true` if the format supports Cinematic Video that
// produces a controllable, simulated depth of field and adds beautiful focus
// transitions for a cinema-grade look.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isCinematicVideoCaptureSupported
func (c AVCaptureDeviceFormat) CinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}
// Default shallow depth of field simulated aperture.
//
// # Discussion
// 
// This property return a non-zero value on devices that support the shallow
// depth of field effect.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/defaultSimulatedAperture
func (c AVCaptureDeviceFormat) DefaultSimulatedAperture() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("defaultSimulatedAperture"))
	return rv
}
// Minimum supported shallow depth of field simulated aperture.
//
// # Discussion
// 
// On devices that do not support changing the simulated aperture value, this
// returns a value of `0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/minSimulatedAperture
func (c AVCaptureDeviceFormat) MinSimulatedAperture() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("minSimulatedAperture"))
	return rv
}
// Maximum supported shallow depth of field simulated aperture.
//
// # Discussion
// 
// On devices that do not support changing the simulated aperture value, this
// returns a value of `0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/maxSimulatedAperture
func (c AVCaptureDeviceFormat) MaxSimulatedAperture() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("maxSimulatedAperture"))
	return rv
}
// Indicates the maximum zoom factor available for the [VideoZoomFactor]
// property when Cinematic Video capture is enabled on the device input.
//
// # Discussion
// 
// Devices support a limited zoom range when Cinematic Video capture is
// active. If this device format does not support Cinematic Video capture,
// this property returns `1.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactorForCinematicVideo
func (c AVCaptureDeviceFormat) VideoMaxZoomFactorForCinematicVideo() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoMaxZoomFactorForCinematicVideo"))
	return rv
}
// Indicates the minimum zoom factor available for the [VideoZoomFactor]
// property when Cinematic Video capture is enabled on the device input.
//
// # Discussion
// 
// Devices support a limited zoom range when Cinematic Video capture is
// active. If this device format does not support Cinematic Video capture,
// this property returns `1.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMinZoomFactorForCinematicVideo
func (c AVCaptureDeviceFormat) VideoMinZoomFactorForCinematicVideo() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoMinZoomFactorForCinematicVideo"))
	return rv
}
// Indicates the minimum / maximum frame rates available when Cinematic Video
// capture is enabled on the device input.
//
// # Discussion
// 
// Devices may support a limited frame rate range when Cinematic Video capture
// is active. If this device format does not support Cinematic Video capture,
// this property returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForCinematicVideo
func (c AVCaptureDeviceFormat) VideoFrameRateRangeForCinematicVideo() IAVFrameRateRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoFrameRateRangeForCinematicVideo"))
	return AVFrameRateRangeFromID(objc.ID(rv))
}
// Whether camera lens smudge detection is supported.
//
// # Discussion
// 
// This property returns `true` if the session’s current configuration
// supports lens smudge detection. When switching cameras or formats, this
// property may change. When this property changes from `true` to `false`,
// [CameraLensSmudgeDetectionEnabled] also reverts to `false`. If you opt in
// for lens smudge detection and then change configurations, you should set
// [CameraLensSmudgeDetectionEnabled] to `true` again.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isCameraLensSmudgeDetectionSupported
func (c AVCaptureDeviceFormat) CameraLensSmudgeDetectionSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCameraLensSmudgeDetectionSupported"))
	return rv
}
// A Boolean value that indicates whether the format supports Center Stage.
//
// # Discussion
// 
// See [AVCaptureDevice] for more information on using Center Stage.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isCenterStageSupported
func (c AVCaptureDeviceFormat) CenterStageSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCenterStageSupported"))
	return rv
}
// The range of frame rates available when Center Stage is active.
//
// # Discussion
// 
// Devices may support a limited frame rate range when Center Stage is active.
// The value is `nil` if the device doesn’t support Center Stage.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForCenterStage
func (c AVCaptureDeviceFormat) VideoFrameRateRangeForCenterStage() IAVFrameRateRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoFrameRateRangeForCenterStage"))
	return AVFrameRateRangeFromID(objc.ID(rv))
}
// The minimum zoom factor available when Center Stage is active.
//
// # Discussion
// 
// Devices support a limited zoom range when Center Stage is active. If the
// device doesn’t support Center Stage, the value is 1.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMinZoomFactorForCenterStage
func (c AVCaptureDeviceFormat) VideoMinZoomFactorForCenterStage() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoMinZoomFactorForCenterStage"))
	return rv
}
// The maximum zoom factor available when Center Stage is active.
//
// # Discussion
// 
// Devices support a limited zoom range when Center Stage is active. If the
// device doesn’t support Center Stage, the value is [VideoMaxZoomFactor].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactorForCenterStage
func (c AVCaptureDeviceFormat) VideoMaxZoomFactorForCenterStage() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoMaxZoomFactorForCenterStage"))
	return rv
}
// A Boolean value that indicates whether the format supports the Portrait
// Effect feature.
//
// # Discussion
// 
// Enabling a Portrait Effect applies a shallow depth-of-field effect to
// objects in the background. See the [PortraitEffectEnabled] property of
// [AVCaptureDevice] for more information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isPortraitEffectSupported
func (c AVCaptureDeviceFormat) PortraitEffectSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPortraitEffectSupported"))
	return rv
}
// The range of frame rates available when Portrait Effect is active.
//
// # Discussion
// 
// Devices may support a limited range of frame rates when Portrait Effect is
// active. If a device format doesn’t support Portrait Effect, the value of
// this property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForPortraitEffect
func (c AVCaptureDeviceFormat) VideoFrameRateRangeForPortraitEffect() IAVFrameRateRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoFrameRateRangeForPortraitEffect"))
	return AVFrameRateRangeFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the format supports Studio Light.
//
// # Discussion
// 
// See [StudioLightEnabled] for more information on the Studio Light feature.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isStudioLightSupported
func (c AVCaptureDeviceFormat) StudioLightSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStudioLightSupported"))
	return rv
}
// A value that indicates the minimum and maximum frame rates available when a
// user enables Studio Light.
//
// # Discussion
// 
// Devices may support a limited frame rate range when Studio Light is active.
// If the format doesn’t support Studio Light, this property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFrameRateRangeForStudioLight
func (c AVCaptureDeviceFormat) VideoFrameRateRangeForStudioLight() IAVFrameRateRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoFrameRateRangeForStudioLight"))
	return AVFrameRateRangeFromID(objc.ID(rv))
}
// Indicates whether the format supports the Edge Light feature.
//
// # Discussion
// 
// This property returns YES if the device supports the Edge Light feature.
// See +AVCaptureDevice.edgeLightEnabled.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isEdgeLightSupported
func (c AVCaptureDeviceFormat) EdgeLightSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEdgeLightSupported"))
	return rv
}
// The currently active depth data format of the capture device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activedepthdataformat
func (c AVCaptureDeviceFormat) ActiveDepthDataFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeDepthDataFormat"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureDeviceFormat) SetActiveDepthDataFormat(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveDepthDataFormat:"), value)
}
// The capture format in use by the device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activeformat
func (c AVCaptureDeviceFormat) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeFormat"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureDeviceFormat) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveFormat:"), value)
}
// The capture formats a device supports.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/formats
func (c AVCaptureDeviceFormat) Formats() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("formats"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureDeviceFormat) SetFormats(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setFormats:"), value)
}
// The zoom factors at which this device transitions to secondary native
// resolution modes.
//
// # Discussion
// 
// Devices that provide secondary native resolution zoom factors can switch
// their pixel sampling mode dynamically to produce high-fidelity images
// without upscaling at a fixed zoom factor beyond 1.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/secondaryNativeResolutionZoomFactors
func (c AVCaptureDeviceFormat) SecondaryNativeResolutionZoomFactors() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("secondaryNativeResolutionZoomFactors"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The list of color spaces the format supports for image and video capture.
//
// # Discussion
// 
// The [NSNumber] objects in this array contain [AVCaptureColorSpace] values.
// The ordering of the array is such that spaces with a narrower color gamut
// appear before those with wider color gamuts.
// 
// All devices and formats support capture in the sRGB color space. Some
// devices and formats can also capture in the P3 color space, which includes
// a much wider gamut of colors than the sRGB color space. (Content captured
// in the P3 color space is viewable on all devices. Devices without
// wide-color displays render P3 content as accurately as possible in the sRGB
// color gamut). By default, a capture session automatically enables
// wide-gamut capture for supported devices and capture workflows (for
// details, see the [AVCaptureSession] property
// [AutomaticallyConfiguresCaptureDeviceForWideColor]).
//
// [AVCaptureColorSpace]: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedColorSpaces
func (c AVCaptureDeviceFormat) SupportedColorSpaces() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedColorSpaces"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The maximum photo dimension this format supports.
//
// # Discussion
// 
// The array contains [NSValue] objects that hold a [CMVideoDimensions]
// structure.
//
// [CMVideoDimensions]: https://developer.apple.com/documentation/CoreMedia/CMVideoDimensions
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedMaxPhotoDimensions
func (c AVCaptureDeviceFormat) SupportedMaxPhotoDimensions() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedMaxPhotoDimensions"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// The zoom ranges that support the delivery of depth data.
//
// # Discussion
// 
// Virtual devices support limited zoom ranges when delivering depth data to
// any output. If a device format has no [SupportedDepthDataFormats] values,
// this property value is an empty array.
// 
// The presence of one or more ranges where the minimum and maximum zoom
// factors aren’t equal means the system supports continuous zoom with
// depth. For example, if the value of this property contains zoom ranges with
// equal minimum and maximum values, the system only allows you to set zoom
// factors equal to these values when you enable depth data delivery. Setting
// a zoom factor other than these values results in an exception.
// Alternatively, when a range’s minimum and maximum values aren’t the
// same, the system supports depth data delivery with across a range of zoom
// factors. You can set a zoom factor outside this range, but results in a
// loss of depth data. Setting the zoom factor back to the supported range
// resumes depth data delivery.
// 
// When you enable depth data delivery, the effective
// [VideoZoomFactorUpscaleThreshold] is `1.0`, which means that all zoom
// factors that aren’t native zoom factors (see
// [VirtualDeviceSwitchOverVideoZoomFactors] and
// [secondaryNativeResolutionZoomFactors]) result in digital upscaling.
//
// [secondaryNativeResolutionZoomFactors]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/secondaryNativeResolutionZoomFactors
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedVideoZoomRangesForDepthDataDelivery
func (c AVCaptureDeviceFormat) SupportedVideoZoomRangesForDepthDataDelivery() []AVZoomRange {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedVideoZoomRangesForDepthDataDelivery"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVZoomRange {
		return AVZoomRangeFromID(id)
	})
}
// The system’s recommended exposure bias range for this device format.
//
// # Discussion
// 
// Use this value to create a slider in your app’s user interface that
// controls a device’s exposure bias within a system-recommended range. When
// a recommendation isn’t available, this property returns `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/systemRecommendedExposureBiasRange
func (c AVCaptureDeviceFormat) SystemRecommendedExposureBiasRange() IAVExposureBiasRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("systemRecommendedExposureBiasRange"))
	return AVExposureBiasRangeFromID(objc.ID(rv))
}
// The system’s recommended zoom range for this device format.
//
// # Discussion
// 
// Use this value to create a slider in your app’s user interface that
// controls a device’s zoom within a system-recommended range. When a
// recommendation isn’t available, this property returns `nil`.
// 
// Apps can key-value observe a capture device’s
// [MinAvailableVideoZoomFactor] and [MaxAvailableVideoZoomFactor] property
// values to know when a device limits its supported zoom to the recommended
// range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/systemRecommendedVideoZoomRange
func (c AVCaptureDeviceFormat) SystemRecommendedVideoZoomRange() IAVZoomRange {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("systemRecommendedVideoZoomRange"))
	return AVZoomRangeFromID(objc.ID(rv))
}
// A value that controls the cropping and enlargement of images captured by
// the device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/videozoomfactor
func (c AVCaptureDeviceFormat) VideoZoomFactor() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("videoZoomFactor"))
	return rv
}
func (c AVCaptureDeviceFormat) SetVideoZoomFactor(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setVideoZoomFactor:"), value)
}

