// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
)

// Retrieves the resolved dimensions of the semantic segmentation mattes that
// the photo output delivers.
//
// semanticSegmentationMatteType: The segmentation matte type for which to retrieve dimensions.
//
// # Return Value
//
// A [CMVideoDimensions] structure that provides the height and width of the
// image mattes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/dimensionsForSemanticSegmentationMatte(ofType:)
//
// [CMVideoDimensions]: https://developer.apple.com/documentation/CoreMedia/CMVideoDimensions
func (c AVCaptureResolvedPhotoSettings) DimensionsForSemanticSegmentationMatteOfType(semanticSegmentationMatteType AVSemanticSegmentationMatteType) coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("dimensionsForSemanticSegmentationMatteOfType:"), objc.String(string(semanticSegmentationMatteType)))
	return coremedia.CMVideoDimensions(rv)
}

// A Boolean value indicating whether the camera flash fires for this capture.
//
// # Discussion
//
// This property corresponds to the [AVCapturePhotoSettings] property
// [FlashMode].
//
// If you specify a flash mode of [AVCaptureFlashModeAuto] when requesting a
// capture, the device automatically chooses whether to use the flash based on
// the scene contents at the moment of capture. Therefore, you don’t know
// whether the flash will fire until right before the moment of capture. When
// the photo output calls your
// [CaptureOutputWillBeginCaptureForResolvedSettings] method (or other
// delegate methods that occur later in the capture process), you can use this
// property to determine whether a capture uses the flash.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isFlashEnabled
func (c AVCaptureResolvedPhotoSettings) FlashEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFlashEnabled"))
	return rv
}

// A Boolean value indicating whether the camera automatically reduces red-eye
// when capturing photos.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isRedEyeReductionEnabled
func (c AVCaptureResolvedPhotoSettings) RedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isRedEyeReductionEnabled"))
	return rv
}

// A Boolean value that specifies whether the system automatically uses
// virtual device image fusion.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isVirtualDeviceFusionEnabled
func (c AVCaptureResolvedPhotoSettings) VirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVirtualDeviceFusionEnabled"))
	return rv
}

// A Boolean value that indicates whether the system applies content-aware
// distortion correction when capturing the photo.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isContentAwareDistortionCorrectionEnabled
func (c AVCaptureResolvedPhotoSettings) ContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContentAwareDistortionCorrectionEnabled"))
	return rv
}

// A Boolean value indicating whether this capture uses image stabilization.
//
// # Discussion
//
// This property corresponds to the [AVCapturePhotoSettings] property
// [AutoStillImageStabilizationEnabled].
//
// When this value is true, the device automatically applies stabilization in
// low-light conditions to counteract hand shake. Automatic stabilization
// always includes digital image stabilization, and may also include optical
// lens stabilization, based on the current device.
//
// If you specify automatic stabilization when requesting a capture, the
// device automatically chooses whether to use image stabilization based on
// the scene contents at the moment of capture. Therefore, you don’t know
// whether the system uses stabilization until right before the moment of
// capture. When the photo output calls your
// [CaptureOutputWillBeginCaptureForResolvedSettings] method (or other
// delegate methods that occur later in the capture process), you can use this
// property to determine whether stabilization is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isStillImageStabilizationEnabled
func (c AVCaptureResolvedPhotoSettings) StillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStillImageStabilizationEnabled"))
	return rv
}

// A Boolean value indicating whether this capture combines image data from a
// dual camera.
//
// # Discussion
//
// This property corresponds to the [AVCapturePhotoSettings] property
// [AutoDualCameraFusionEnabled].
//
// When this value is true, a dual-camera device automatically combines
// samples from both cameras to produce a higher quality image. This property
// applies only when using the [builtInDualCamera] device type on supported
// devices.
//
// If you specify automatic image fusion when requesting a capture, the device
// automatically chooses whether to use image fusion based on the scene
// conditions at the moment of capture. Therefore, you don’t know whether
// the system uses image fusion until right before the moment of capture. When
// the photo output calls your
// [CaptureOutputWillBeginCaptureForResolvedSettings] method (or other
// delegate methods that occur later in the capture process), you can use this
// property to determine whether image fusion is active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isDualCameraFusionEnabled
//
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
func (c AVCaptureResolvedPhotoSettings) DualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDualCameraFusionEnabled"))
	return rv
}

// The resolved dimensions of the photo proxy when using deferred photo
// delivery.
//
// # Discussion
//
// When the system returns an [AVCaptureDeferredPhotoProxy], the
// [PhotoDimensions] property of this object represents the dimensions of the
// final photo. If you don’t opt in to deferred photo delivery, this value
// has a width and height of 0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/deferredPhotoProxyDimensions
//
// [AVCaptureDeferredPhotoProxy]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeferredPhotoProxy
func (c AVCaptureResolvedPhotoSettings) DeferredPhotoProxyDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("deferredPhotoProxyDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The size, in pixels, of the RAW-format photo image that the capture
// delivers.
//
// # Discussion
//
// The output dimensions of a captured image are set at the moment of capture,
// depending on device orientation and capture session configuration. (For
// example, when the capture session includes a video output and video
// stabilization is in use, captured photos are smaller.)
//
// This property provides the dimensions of the image to be delivered in the
// [CaptureOutputDidFinishProcessingPhotoError] method. Use this property in
// earlier delegate methods to find the size of the image before delivery.
//
// If you do not request capture in RAW format, this property’s value has
// zero width and zero height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/rawPhotoDimensions
func (c AVCaptureResolvedPhotoSettings) RawPhotoDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("rawPhotoDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The size, in pixels, of the preview image that the system delivers with the
// capture.
//
// # Discussion
//
// Use the [PreviewPhotoFormat] property in your photo settings object to
// request delivery of a preview image alongside the main photo output from
// the capture. When you request a preview, the photo output chooses
// dimensions that best match your requested size while preserving the aspect
// ratio of the captured photo. Aspect ratio is determined by capture format
// and by device orientation at the moment of capture.
//
// This property provides the dimensions of the requested preview image, which
// is delivered in the [CaptureOutputDidFinishProcessingPhotoError] method.
// Use this property in earlier delegate methods to find the size of the image
// before delivery.
//
// If you do not request a preview image, this property’s value has zero
// width and zero height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/previewDimensions
func (c AVCaptureResolvedPhotoSettings) PreviewDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("previewDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The size, in pixels, of the thumbnail image that the capture delivers.
//
// # Discussion
//
// Use the [EmbeddedThumbnailPhotoFormat] property in your photo settings
// object to request delivery of a thumbnail image alongside the main photo
// output from the capture. When you request a thumbnail, the photo output
// chooses dimensions that best match your requested size while preserving the
// aspect ratio of the captured photo. Aspect ratio is determined by capture
// format and by device orientation at the moment of capture.
//
// This property provides the dimensions of the requested thumbnail image,
// which is delivered in the [CaptureOutputDidFinishProcessingPhotoError]
// method. Use this property in earlier delegate methods to find the size of
// the image before delivery.
//
// If you do not request a thumbnail image, this property’s value has zero
// width and zero height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/embeddedThumbnailDimensions
func (c AVCaptureResolvedPhotoSettings) EmbeddedThumbnailDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("embeddedThumbnailDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The size, in pixels, of the RAW-format embedded thumbnail image that the
// capture delivers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/rawEmbeddedThumbnailDimensions
func (c AVCaptureResolvedPhotoSettings) RawEmbeddedThumbnailDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("rawEmbeddedThumbnailDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The size, in pixels, of the Live Photo movie content that the capture
// delivers.
//
// # Discussion
//
// Use the [LivePhotoMovieFileURL] property in your photo settings object to
// request Live Photo capture. Live Photo movie dimensions can change
// depending on which device camera is used for capture.
//
// This property provides dimensions for the movie content of the Live Photo,
// which is delivered in the
// [CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError]
// method. Use this property in earlier delegate methods to find the
// dimensions of the movie before delivery.
//
// For the dimensions of the static photo content of a Live Photo, see the
// [PhotoDimensions] property.
//
// If you do not request Live Photo capture, this property’s value has zero
// width and zero height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/livePhotoMovieDimensions
func (c AVCaptureResolvedPhotoSettings) LivePhotoMovieDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("livePhotoMovieDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The size, in pixels, of the portrait effects matte that the capture
// delivers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/portraitEffectsMatteDimensions
func (c AVCaptureResolvedPhotoSettings) PortraitEffectsMatteDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("portraitEffectsMatteDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// The time range in which to expect the system to deliver the photo to the
// delegate.
//
// # Discussion
//
// Indicates the processing time range you can expect for this photo to be
// delivered to your delegate. the .start field of the CMTimeRange is
// zero-based. In other words, if photoProcessingTimeRange.start is equal to
// .5 seconds, then the minimum processing time for this photo is .5 seconds.
// The .start field plus the .duration field of the CMTimeRange indicate the
// max expected processing time for this photo. Consider implementing a UI
// affordance if the max processing time is uncomfortably long.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/photoProcessingTimeRange
func (c AVCaptureResolvedPhotoSettings) PhotoProcessingTimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](c.ID, objc.Sel("photoProcessingTimeRange"))
	return coremedia.CMTimeRange(rv)
}
