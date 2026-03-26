// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/coregraphics"
"github.com/tmc/apple/corevideo"
"github.com/tmc/apple/foundation"
"github.com/tmc/apple/objectivec"
)

// Gets a customized representation of the photo data.
//
// customizer: An object that customizes the returned metadata, image thumbnail, or depth
// data.
//
// # Return Value
// 
// A data representation of the photo.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/fileDataRepresentation(with:)
func (c AVCapturePhoto) FileDataRepresentationWithCustomizer(customizer objectivec.IObject) foundation.INSData {
rv := objc.Send[objc.ID](c.ID, objc.Sel("fileDataRepresentationWithCustomizer:"), customizer)
return foundation.NSDataFromID(rv)
}
// Extracts and returns the captured photo’s preview image as a Core
// Graphics image object.
//
// # Return Value
// 
// A Core Graphics image representation of the captured photo, or `nil` if
// either the image cannot be converted or a preview image was not requested
// as part of the photo capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/previewCGImageRepresentation()
func (c AVCapturePhoto) PreviewCGImageRepresentation() coregraphics.CGImageRef {
rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("previewCGImageRepresentation"))
return coregraphics.CGImageRef(rv)
}
// Retrieves the semantic segmentation matte associated with this photo.
//
// semanticSegmentationMatteType: The type of semantic segmentation matte to retrieve from the photo.
//
// # Return Value
// 
// An instance of [AVSemanticSegmentationMatte], or `nil` of you didn’t
// request semantic segmentation matte delivery or if no mattes of the
// specified type were found.
//
// # Discussion
// 
// If you requested one or more semantic segmentation mattes by calling
// [EnabledSemanticSegmentationMatteTypes] with a nonempty array of types,
// this property offers access to the resulting [AVSemanticSegmentationMatte]
// objects.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/semanticSegmentationMatte(for:)
func (c AVCapturePhoto) SemanticSegmentationMatteForType(semanticSegmentationMatteType AVSemanticSegmentationMatteType) IAVSemanticSegmentationMatte {
rv := objc.Send[objc.ID](c.ID, objc.Sel("semanticSegmentationMatteForType:"), objc.String(string(semanticSegmentationMatteType)))
return AVSemanticSegmentationMatteFromID(rv)
}

// A Boolean value indicating whether this photo object contains RAW format
// data.
//
// # Discussion
// 
// When you request capture in RAW format, the capture output calls your
// delegate’s [CaptureOutputDidFinishProcessingPhotoError] method one or
// more times, delivering both the RAW photo data and (if requested)
// equivalent processed photos. Use this property to distinguish between the
// RAW and processed results from the same capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/isRawPhoto
func (c AVCapturePhoto) RawPhoto() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isRawPhoto"))
		return rv
}
// A dictionary describing the data format for a preview-sized image
// accompanying the captured photo.
//
// # Discussion
// 
// See [Video settings] for possible keys and values.
// 
// If you requested an embedded thumbnail image by specifying the
// [EmbeddedThumbnailPhotoFormat] property of your photo settings when
// requesting capture, this property’s value is the resolved video settings
// dictionary for the embedded thumbnail image. If you did not request an
// embedded thumbnail image, this property’s value is `nil`.
//
// [Video settings]: https://developer.apple.com/documentation/AVFoundation/video-settings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/embeddedThumbnailPhotoFormat
func (c AVCapturePhoto) EmbeddedThumbnailPhotoFormat() foundation.INSDictionary {
rv := objc.Send[objc.ID](c.ID, objc.Sel("embeddedThumbnailPhotoFormat"))
return foundation.NSDictionaryFromID(rv)
}
// The pixel data for a preview-sized version of the photo, if requested.
//
// # Discussion
// 
// If you requested a preview image by specifying the [PreviewPhotoFormat]
// property of your photo settings when requesting capture, this property
// offers access to the resulting preview image pixel data. The pixel buffer
// contains only the minimal attachments required for correct display. If you
// did not request a preview image, this property’s value is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/previewPixelBuffer
func (c AVCapturePhoto) PreviewPixelBuffer() corevideo.CVImageBufferRef {
rv := objc.Send[corevideo.CVImageBufferRef](c.ID, objc.Sel("previewPixelBuffer"))
		return corevideo.CVImageBufferRef(rv)
}
// Depth or disparity map data captured with the photo.
//
// # Discussion
// 
// To request capture of depth data alongside a photo (on supported devices),
// set the [DepthDataDeliveryEnabled] property of your photo settings object
// to [true] when requesting photo capture. If you did not request depth data
// delivery, this property’s value is `nil`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/depthData
func (c AVCapturePhoto) DepthData() IAVDepthData {
rv := objc.Send[objc.ID](c.ID, objc.Sel("depthData"))
return AVDepthDataFromID(rv)
}
// Calibration information for the camera device that captured the photo.
//
// # Discussion
// 
// Camera calibration data is present only if you specified the
// [CameraCalibrationDataDeliveryEnabled] and
// [DualCameraDualPhotoDeliveryEnabled] settings when requesting capture. For
// camera calibration data in a capture that includes depth data, see the
// [AVDepthData] [CameraCalibrationData] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/cameraCalibrationData
func (c AVCapturePhoto) CameraCalibrationData() IAVCameraCalibrationData {
rv := objc.Send[objc.ID](c.ID, objc.Sel("cameraCalibrationData"))
return AVCameraCalibrationDataFromID(rv)
}
// The type of device that captured the photo.
//
// # Discussion
// 
// When you capture dual photos with a [builtInDualCamera] device and the
// [DualCameraDualPhotoDeliveryEnabled] setting, use this property to
// determine which of the two resulting photo objects is from the
// [builtInWideAngleCamera] or [builtInTelephotoCamera] device.
// 
// For all other captures, this property’s value is equal to the
// [DeviceType] property of the capture device to which the photo output is
// connected.
// 
// This property’s value can be `nil` if the [AVCapturePhoto] object did not
// come from an [AVCaptureDevice] capture.
//
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
// [builtInTelephotoCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInTelephotoCamera
// [builtInWideAngleCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInWideAngleCamera
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/sourceDeviceType
func (c AVCapturePhoto) SourceDeviceType() AVCaptureDeviceType {
rv := objc.Send[objc.ID](c.ID, objc.Sel("sourceDeviceType"))
return AVCaptureDeviceType(foundation.NSStringFromID(rv).String())
}
// A dictionary of metadata describing the captured image.
//
// # Discussion
// 
// See [CGImageProperties] for possible keys and values. Metadata captured
// with a photo may include image orientation, EXIF camera properties, and
// Live Photo metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/metadata
func (c AVCapturePhoto) Metadata() foundation.INSDictionary {
rv := objc.Send[objc.ID](c.ID, objc.Sel("metadata"))
return foundation.NSDictionaryFromID(rv)
}
// The portrait effects matte captured with the photo.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/portraitEffectsMatte
func (c AVCapturePhoto) PortraitEffectsMatte() IAVPortraitEffectsMatte {
rv := objc.Send[objc.ID](c.ID, objc.Sel("portraitEffectsMatte"))
return AVPortraitEffectsMatteFromID(rv)
}
// The variations available for bracketed capture settings for this photo.
//
// # Discussion
// 
// When you request a bracketed capture using the
// [AVCapturePhotoBracketSettings] class, you specify an array of
// [AVCaptureBracketedStillImageSettings] objects indicating the capture
// setting variations (such as exposure compensation) to apply to each image
// in the bracket. This property indicates the settings associated with this
// particular photo, or `nil` if this photo is not part of a bracketed
// capture.
//
// [AVCaptureBracketedStillImageSettings]: https://developer.apple.com/documentation/AVFoundation/AVCaptureBracketedStillImageSettings
// [AVCapturePhotoBracketSettings]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/bracketSettings
func (c AVCapturePhoto) BracketSettings() objc.ID {
rv := objc.Send[objc.ID](c.ID, objc.Sel("bracketSettings"))
		return rv
}
// The 1-based index of this photo in a bracketed capture sequence.
//
// # Discussion
// 
// If this photo is part of a bracketed capture (requested with the
// [AVCapturePhotoBracketSettings] class), this property indicates the current
// result’s count in the sequence, starting with `1` for the first result.
// 
// If this photo is not part of a bracketed capture, this property’s value
// is `0`.
//
// [AVCapturePhotoBracketSettings]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/sequenceCount
func (c AVCapturePhoto) SequenceCount() int {
rv := objc.Send[int](c.ID, objc.Sel("sequenceCount"))
		return rv
}
// Information about the use of lens stabilization during bracketed photo
// capture.
//
// # Discussion
// 
// This property applies only to capture results for which you requested
// optical image stabilization (OIS) across all frames of a bracketed photo
// capture (using the [AVCapturePhotoBracketSettings]
// [isLensStabilizationEnabled] property).
// 
// If the device configuration does not support OIS, this property’s value
// is [CaptureLensStabilizationStatusUnsupported]. If OIS is supported, but
// this captured photo is not from a bracketed capture where OIS was
// requested, this property’s value is [CaptureLensStabilizationStatusOff].
// Otherwise, this property indicates how the device applied OIS across the
// duration of the bracketed capture.
//
// [AVCapturePhotoBracketSettings]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings
// [isLensStabilizationEnabled]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/isLensStabilizationEnabled
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/lensStabilizationStatus
func (c AVCapturePhoto) LensStabilizationStatus() AVCaptureLensStabilizationStatus {
rv := objc.Send[AVCaptureLensStabilizationStatus](c.ID, objc.Sel("lensStabilizationStatus"))
		return AVCaptureLensStabilizationStatus(rv)
}

