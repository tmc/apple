// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Creates a photo settings object for RAW-format-only capture with the
// specified pixel format.
//
// rawPixelFormatType: The Bayer RAW pixel format type to use for capture. This value must be one
// of the format identifiers listed in the [AvailableRawPhotoPixelFormatTypes]
// array of your photo capture output.
//
// # Return Value
//
// A new photo settings object.
//
// # Discussion
//
// Use this initializer for RAW-only capture. To capture an image in both RAW
// format and a processed format (such as JPEG), use the
// [PhotoSettingsWithRawPixelFormatTypeProcessedFormat] initializer instead.
//
// Requesting RAW format capture adds requirements for other photo settings:
// for details, see the [RawPhotoPixelFormatType] property. The capture output
// validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate
// don’t meet these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:)
func (_AVCapturePhotoSettingsClass AVCapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatType(rawPixelFormatType uint32) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(_AVCapturePhotoSettingsClass.class), objc.Sel("photoSettingsWithRawPixelFormatType:"), rawPixelFormatType)
	return AVCapturePhotoSettingsFromID(rv)
}

// Creates a photo settings object for capture in both RAW format and a
// processed format.
//
// rawPixelFormatType: The Bayer RAW pixel format type to use for capture. This value must be one
// of the format identifiers listed in the [AvailableRawPhotoPixelFormatTypes]
// array of your photo capture output.
//
// processedFormat: A dictionary of Core Video pixel buffer attributes or AVFoundation video
// settings constants (see `Video Settings`).
//
// To capture a photo in an uncompressed format, such as 420f, 420v, or BGRA,
// set the key [kCVPixelBufferPixelFormatTypeKey] in the `format` dictionary.
// The corresponding value must be one of the pixel format identifiers listed
// in the [AvailablePhotoPixelFormatTypes] array of your photo capture output.
//
// To capture a photo in a compressed format, such as JPEG, set the key
// [AVVideoCodecKey] in the `format` dictionary. The corresponding value must
// be one of the codec identifiers listed in the [AvailablePhotoCodecTypes]
// array of your photo capture output. For a compressed format, you can also
// specify a compression level with the key [AVVideoQualityKey].
//
// # Return Value
//
// A new photo settings object.
//
// # Discussion
//
// Use this initializer to capture an image in both RAW format and a processed
// format (such as JPEG). For RAW-only capture, use the
// [PhotoSettingsWithRawPixelFormatType] initializer instead.
//
// Requesting both formats adds requirements for other photo settings: see the
// [Format] property for processed format requirements and the
// [RawPhotoPixelFormatType] property for RAW format requirements. The capture
// output validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate
// don’t meet these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:processedFormat:)
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoQualityKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoQualityKey
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
func (_AVCapturePhotoSettingsClass AVCapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatTypeProcessedFormat(rawPixelFormatType uint32, processedFormat foundation.INSDictionary) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(_AVCapturePhotoSettingsClass.class), objc.Sel("photoSettingsWithRawPixelFormatType:processedFormat:"), rawPixelFormatType, processedFormat)
	return AVCapturePhotoSettingsFromID(rv)
}

// Creates a photo settings object for capture in both RAW format and a
// processed format with the specified output file types.
//
// rawPixelFormatType: The Bayer RAW pixel format type to use for capture. This value must be one
// of the format identifiers listed in the [AvailableRawPhotoPixelFormatTypes]
// array of your photo capture output.
//
// rawFileType: The container file format for eventual output of the RAW image.
//
// If you have no preferred file format, pass `nil` and the photo output will
// automatically choose a default file format appropriate to the
// `rawPixelFormatType` parameter.
//
// processedFormat: A dictionary of Core Video pixel buffer attributes or AVFoundation video
// settings constants (see `Video Settings`).
//
// To capture a photo in an uncompressed format, such as 420f, 420v, or BGRA,
// set the key [kCVPixelBufferPixelFormatTypeKey] in the `format` dictionary.
// The corresponding value must be one of the pixel format identifiers listed
// in the [AvailablePhotoPixelFormatTypes] array of your photo capture output.
//
// To capture a photo in a compressed format, such as JPEG, set the key
// [AVVideoCodecKey] in the `format` dictionary. The corresponding value must
// be one of the codec identifiers listed in the [AvailablePhotoCodecTypes]
// array of your photo capture output. For a compressed format, you can also
// specify a compression level with the key [AVVideoQualityKey].
//
// processedFileType: The container file format for eventual output of the processed image.
//
// If you have no preferred file format, pass `nil` and the photo output will
// automatically choose a default file format appropriate to the
// `processedFormat` parameter.
//
// # Return Value
//
// A new photo settings object.
//
// # Discussion
//
// Use this initializer to capture an image in both RAW format and a processed
// format (such as JPEG). For RAW-only capture, use the
// [PhotoSettingsWithRawPixelFormatType] initializer instead.
//
// Requesting both formats adds requirements for other photo settings: see the
// [Format] property for processed format requirements and the
// [RawPhotoPixelFormatType] property for RAW format requirements. The capture
// output validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate do
// not meet these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(rawPixelFormatType:rawFileType:processedFormat:processedFileType:)
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoQualityKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoQualityKey
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
func (_AVCapturePhotoSettingsClass AVCapturePhotoSettingsClass) PhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType(rawPixelFormatType uint32, rawFileType AVFileType, processedFormat foundation.INSDictionary, processedFileType AVFileType) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(_AVCapturePhotoSettingsClass.class), objc.Sel("photoSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:"), rawPixelFormatType, objc.String(string(rawFileType)), processedFormat, objc.String(string(processedFileType)))
	return AVCapturePhotoSettingsFromID(rv)
}

// The container file format for eventual output of the RAW image.
//
// # Discussion
//
// You specify a file format when creating capture settings with the
// [PhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType]
// initializer. If you didn’t specify a file format, this value is `nil`,
// and the photo output automatically choosea a default file format
// appropriate to the [RawPhotoPixelFormatType] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawFileType
func (c AVCapturePhotoSettings) RawFileType() AVFileType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rawFileType"))
	return AVFileType(foundation.NSStringFromID(rv).String())
}

// An identifier for the Bayer RAW pixel format to deliver captured RAW photos
// in.
//
// # Discussion
//
// This property is read-only—you specify a RAW pixel format when creating a
// settings object with the [PhotoSettingsWithRawPixelFormatType],
// [PhotoSettingsWithRawPixelFormatTypeProcessedFormat] initializer.
//
// When capturing RAW images, the following requirements apply:
//
// - The [AutoStillImageStabilizationEnabled] setting must be false. - Your
// delegate object must implement the
// [CaptureOutputDidFinishProcessingRawPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError]
// method. - The [isHighResolutionPhotoEnabled] setting may be true or false,
// but that setting applies only to the separate processed image.
//
// (You request separate processed images with the
// [PhotoSettingsWithRawPixelFormatTypeProcessedFormat] initializer. This
// restriction does not apply when you request RAW-only capture with the
// [PhotoSettingsWithRawPixelFormatType] initializer).
//
// The capture output validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate do
// not meet these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawPhotoPixelFormatType
//
// [isHighResolutionPhotoEnabled]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isHighResolutionPhotoEnabled
func (c AVCapturePhotoSettings) RawPhotoPixelFormatType() uint32 {
	rv := objc.Send[uint32](c.ID, objc.Sel("rawPhotoPixelFormatType"))
	return rv
}

// A Boolean value that indicates whether to use auto red-eye reduction on
// flash captures.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoRedEyeReductionEnabled
func (c AVCapturePhotoSettings) AutoRedEyeReductionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoRedEyeReductionEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetAutoRedEyeReductionEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoRedEyeReductionEnabled:"), value)
}

// A Boolean value that determines whether a dual photo capture also delivers
// camera calibration data.
//
// # Discussion
//
// When this setting is false (the default), and the
// [DualCameraDualPhotoDeliveryEnabled] setting is true, dual photo capture
// doesn’t deliver additional data.
//
// If you change this setting to true, the [AVCapturePhoto] results from a
// dual photo capture include [AVCameraCalibrationData] objects that describe
// the imaging parameters for each camera. This data can be useful for
// performing computer vision tasks on the resulting images.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isCameraCalibrationDataDeliveryEnabled
func (c AVCapturePhotoSettings) CameraCalibrationDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCameraCalibrationDataDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetCameraCalibrationDataDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCameraCalibrationDataDeliveryEnabled:"), value)
}

// A Boolean value that specifies whether the photo output, at its discretion,
// uses content-aware distortion correction on this photo request.
//
// # Discussion
//
// The default value is false. Enabling this setting introduces a small amount
// of latency when processing the photo request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoContentAwareDistortionCorrectionEnabled
func (c AVCapturePhotoSettings) AutoContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoContentAwareDistortionCorrectionEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetAutoContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoContentAwareDistortionCorrectionEnabled:"), value)
}

// A Boolean value that specifies whether to use automatic virtual-device
// image fusion.
//
// # Discussion
//
// When [AutoVirtualDeviceFusionEnabled] and [VirtualDeviceFusionSupported]
// are true, the framework may fuse constituent camera images of a virtual
// device to improve still image quality, depending on the current zoom
// factor, light levels, and focus position. You can determine whether virtual
// device fusion is enabled for a particular capture request by inspecting the
// [VirtualDeviceFusionEnabled] property of [AVCaptureResolvedPhotoSettings].
//
// The default value for this property is true, unless you’re capturing a
// RAW photo or a bracket using [AVCapturePhotoBracketSettings].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoVirtualDeviceFusionEnabled
//
// [AVCapturePhotoBracketSettings]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings
func (c AVCapturePhotoSettings) AutoVirtualDeviceFusionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoVirtualDeviceFusionEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetAutoVirtualDeviceFusionEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoVirtualDeviceFusionEnabled:"), value)
}

// The constituent devices for which the virtual device should deliver photos.
//
// # Discussion
//
// You can opt in to constituent-device photo delivery by setting this
// property to any subset of the devices in the virtual device’s
// [ConstituentDevices] array. The framework calls your
// [CaptureOutputDidFinishProcessingPhotoError]callback once for each of the
// devices you include in the array.
//
// You may only set this property to a non-`nil` array if you’ve set your
// photo output’s [VirtualDeviceConstituentPhotoDeliveryEnabled] property to
// true, and your delegate implements the
// [CaptureOutputDidFinishProcessingPhotoError] method.
//
// The default value of this property is an empty array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/virtualDeviceConstituentPhotoDeliveryEnabledDevices
func (c AVCapturePhotoSettings) VirtualDeviceConstituentPhotoDeliveryEnabledDevices() []AVCaptureDevice {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("virtualDeviceConstituentPhotoDeliveryEnabledDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDevice {
		return AVCaptureDeviceFromID(id)
	})
}
func (c AVCapturePhotoSettings) SetVirtualDeviceConstituentPhotoDeliveryEnabledDevices(value []AVCaptureDevice) {
	objc.Send[struct{}](c.ID, objc.Sel("setVirtualDeviceConstituentPhotoDeliveryEnabledDevices:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that determines whether a dual camera device delivers
// images from both cameras.
//
// # Discussion
//
// When this property is false (the default), and the photo output is
// configured with a capture device of the [builtInDualCamera] type, the photo
// output delivers a single main photo image for each capture. (The device
// determines how to produce that image from one or both cameras).
//
// If you change this setting to true, your delegate’s
// [CaptureOutputDidFinishProcessingPhotoError] method fires at least twice
// for each main image—once for the telephoto image and again for the
// wide-angle image. Setting this property to true when not using a
// dual-camera device raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isDualCameraDualPhotoDeliveryEnabled
//
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
func (c AVCapturePhotoSettings) DualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDualCameraDualPhotoDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetDualCameraDualPhotoDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDualCameraDualPhotoDeliveryEnabled:"), value)
}

// A Boolean value that specifies whether captures automatically combine data
// from a dual camera device.
//
// # Discussion
//
// The default setting is true, unless you are capturing a RAW photo. (By
// definition, RAW photos are unprocessed, and image fusion involves
// processing the captured image).
//
// When you enable this setting, a dual-camera device automatically combines
// samples from both cameras to produce a higher quality image. This property
// applies only when using the [builtInDualCamera] device type on supported
// devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoDualCameraFusionEnabled
//
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
func (c AVCapturePhotoSettings) AutoDualCameraFusionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoDualCameraFusionEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetAutoDualCameraFusionEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoDualCameraFusionEnabled:"), value)
}

// A Boolean value that specifies whether captures use automatic image
// stabilization.
//
// # Discussion
//
// The default setting is true, unless you are capturing a RAW photo (By
// definition, RAW photos are unprocessed, and image stabilization involves
// processing the captured image).
//
// When you enable this setting, the device automatically applies
// stabilization in low-light conditions to counteract hand shake. Automatic
// stabilization always includes digital image stabilization, and may also
// include optical lens stabilization, based on the current device.
//
// Automatic image stabilization is not compatible with the
// [AVCaptureFlashModeOn] setting. If you explicitly enable the flash, the
// photo output ignores your image stabilization setting, and the
// [StillImageStabilizationEnabled] property of the
// [AVCaptureResolvedPhotoSettings] object provided to your photo capture
// delegate is always false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isAutoStillImageStabilizationEnabled
func (c AVCapturePhotoSettings) AutoStillImageStabilizationEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoStillImageStabilizationEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetAutoStillImageStabilizationEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoStillImageStabilizationEnabled:"), value)
}

// A dictionary describing the format for delivery of preview-sized images
// alongside the main photo.
//
// # Discussion
//
// By default, this property is `nil`, specifying that the photo output should
// return only the main image requested. To also receive a preview-sized
// image, set this property to a dictionary describing the format for that
// image, containing the following keys and values:
//
// The dictionary must contain the [kCVPixelBufferPixelFormatTypeKey] key,
// whose corresponding value must be one of the pixel format types listed in
// the [AvailablePreviewPhotoPixelFormatTypes] array.
//
// Optionally, you can also include the [KCVPixelBufferWidthKey] and
// [KCVPixelBufferHeightKey] keys to specify the size of the preview image.
// (If you specify either width or height, you must specify both.) If the size
// you specify does not match the aspect ratio of the primary photo, the photo
// output provides a preview image whose size matches the longer of the two
// specified dimensions, preserving the original aspect ratio.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/previewPhotoFormat
//
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
func (c AVCapturePhotoSettings) PreviewPhotoFormat() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previewPhotoFormat"))
	return foundation.NSDictionaryFromID(rv)
}
func (c AVCapturePhotoSettings) SetPreviewPhotoFormat(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreviewPhotoFormat:"), value)
}

// A dictionary describing the format for delivery of thumbnail images
// embedded in photo file output.
//
// # Discussion
//
// By default, this property is `nil`, specifying that the photo output should
// not embed thumbnail images in photo file output. To enable embedding, set
// this property to a dictionary describing the format for thumbnail images,
// containing the following keys and values:
//
// - The dictionary must contain the key [AVVideoCodecKey], whose
// corresponding value must be one of the pixel format types listed in the
// [AvailableEmbeddedThumbnailPhotoCodecTypes] array. - Optionally, you can
// also include the [AVVideoWidthKey] and [AVVideoHeightKey] keys to specify
// the size of the thumbnail image. (If you specify either width or height,
// you must specify both.) If the size you specify does not match the aspect
// ratio of the primary photo, the photo output provides a thumbnail image
// whose size matches the longer of the two specified dimensions, preserving
// the original aspect ratio.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embeddedThumbnailPhotoFormat
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoHeightKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoHeightKey
// [AVVideoWidthKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoWidthKey
func (c AVCapturePhotoSettings) EmbeddedThumbnailPhotoFormat() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("embeddedThumbnailPhotoFormat"))
	return foundation.NSDictionaryFromID(rv)
}
func (c AVCapturePhotoSettings) SetEmbeddedThumbnailPhotoFormat(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setEmbeddedThumbnailPhotoFormat:"), value)
}

// An array of video codec types compatible with the photo settings for
// embedding raw thumbnail images in photo file output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/availableRawEmbeddedThumbnailPhotoCodecTypes
func (c AVCapturePhotoSettings) AvailableRawEmbeddedThumbnailPhotoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableRawEmbeddedThumbnailPhotoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// A dictionary describing the format for delivery of raw thumbnail images
// embedded in photo file output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawEmbeddedThumbnailPhotoFormat
func (c AVCapturePhotoSettings) RawEmbeddedThumbnailPhotoFormat() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rawEmbeddedThumbnailPhotoFormat"))
	return foundation.NSDictionaryFromID(rv)
}
func (c AVCapturePhotoSettings) SetRawEmbeddedThumbnailPhotoFormat(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setRawEmbeddedThumbnailPhotoFormat:"), value)
}

// An array of video codec types compatible with the photo settings for
// embedding thumbnail images in photo file output.
//
// # Discussion
//
// To enable embedding thumbnail images in photo file output, set the
// [EmbeddedThumbnailPhotoFormat] property using one of the codec types listed
// in this array.
//
// The order of this array is such that the most backward-compatible codec is
// listed first.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/availableEmbeddedThumbnailPhotoCodecTypes
func (c AVCapturePhotoSettings) AvailableEmbeddedThumbnailPhotoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableEmbeddedThumbnailPhotoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// A URL at which to write Live Photo movie output.
//
// # Discussion
//
// Live Photos capture both a still image and a short movie, which the system
// presents together in user interfaces such as the Photos app. By default,
// this property’s value is `nil`, disabling Live Photo capture. Set this
// property to a file URL to capture Live Photos.
//
// When you enable Live Photo capture, the following requirements apply:
//
// - The photo output’s [LivePhotoCaptureEnabled] property must be true, and
// its and [LivePhotoCaptureSuspended] property must be false. - The URL you
// specify must be a file URL to an accessible location in your app’s
// sandbox. - Your delegate object must implement the
// [CaptureOutputDidFinishProcessingLivePhotoToMovieFileAtURLDurationPhotoDisplayTimeResolvedSettingsError]
// method.
//
// The capture output validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate
// don’t meet these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/livePhotoMovieFileURL
func (c AVCapturePhotoSettings) LivePhotoMovieFileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("livePhotoMovieFileURL"))
	return foundation.NSURLFromID(rv)
}
func (c AVCapturePhotoSettings) SetLivePhotoMovieFileURL(value foundation.INSURL) {
	objc.Send[struct{}](c.ID, objc.Sel("setLivePhotoMovieFileURL:"), value)
}

// A dictionary of metadata to include in the Live Photo movie file.
//
// # Discussion
//
// Live Photos capture both a still image and a short movie, which the system
// presents together in user interfaces such as the Photos app. A Live Photo
// movie always contains a [AVMetadataQuickTimeMetadataKeyContentIdentifier]
// key, associating the movie with a similar identifier in the
// [kCGImagePropertyExifMakerNote] property of the corresponding still image.
// The photo capture output automatically generates a unique content
// identifier for you if you don’t specify one of your own. You can also use
// this property to specify additional movie metadata.
//
// This property applies only if the value of the [LivePhotoMovieFileURL]
// property is to non-`nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/livePhotoMovieMetadata
//
// [kCGImagePropertyExifMakerNote]: https://developer.apple.com/documentation/ImageIO/kCGImagePropertyExifMakerNote
func (c AVCapturePhotoSettings) LivePhotoMovieMetadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("livePhotoMovieMetadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (c AVCapturePhotoSettings) SetLivePhotoMovieMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setLivePhotoMovieMetadata:"), objectivec.IObjectSliceToNSArray(value))
}

// The video codec to use for encoding the movie portion of Live Photo output.
//
// # Discussion
//
// This value must be one of the video codec types listed in the photo
// output’s [AvailableLivePhotoVideoCodecTypes] array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/livePhotoVideoCodecType
func (c AVCapturePhotoSettings) LivePhotoVideoCodecType() AVVideoCodecType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("livePhotoVideoCodecType"))
	return AVVideoCodecType(foundation.NSStringFromID(rv).String())
}
func (c AVCapturePhotoSettings) SetLivePhotoVideoCodecType(value AVVideoCodecType) {
	objc.Send[struct{}](c.ID, objc.Sel("setLivePhotoVideoCodecType:"), value)
}

// A Boolean value that determines whether the photo output captures depth
// data along with the photo.
//
// # Discussion
//
// When this property is false (the default), the capture output produces only
// photo data and metadata.
//
// If you change this property to true, the capture output records per-pixel
// scene depth information and delivers an [AVDepthData] object in the photo
// capture results. Enabling depth capture for a photo capture request
// requires that the photo output first be configured for depth capture using
// its own [DepthDataDeliveryEnabled] property (and raises an exception
// otherwise).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isDepthDataDeliveryEnabled
func (c AVCapturePhotoSettings) DepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetDepthDataDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDepthDataDeliveryEnabled:"), value)
}

// A Boolean value that determines whether any depth data captured with the
// photo is included when generating output file data.
//
// # Discussion
//
// When this property is true (the default), and depth data capture is enabled
// with the [DepthDataDeliveryEnabled] property, the [AVCapturePhoto] class
// includes the depth map as an embedded attachment when you flatten the photo
// data for output in compatible file formats.
//
// Set this property to false if you wish to capture depth data with a photo
// but not include depth data in output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embedsDepthDataInPhoto
func (c AVCapturePhotoSettings) EmbedsDepthDataInPhoto() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("embedsDepthDataInPhoto"))
	return rv
}
func (c AVCapturePhotoSettings) SetEmbedsDepthDataInPhoto(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEmbedsDepthDataInPhoto:"), value)
}

// A Boolean value that determines whether to smooth noise and fill in missing
// values in depth data output.
//
// # Discussion
//
// When this property is true (the default), and depth data capture is enabled
// with the [DepthDataDeliveryEnabled] property, the capture system smooths
// noise and fills in missing values (caused by low light or lens occlusion)
// in depth data maps by temporally interpolating between previous and
// subsequent frames of captured depth data.
//
// Filtering depth data makes it more useful for applying visual effects to a
// companion image, but alters the data such that it may no longer be suitable
// for computer vision tasks. (In an unfiltered depth map, missing values are
// represented as [NaN].) Set this property to false to disable filtering and
// receive unfiltered depth data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isDepthDataFiltered
func (c AVCapturePhotoSettings) DepthDataFiltered() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDepthDataFiltered"))
	return rv
}
func (c AVCapturePhotoSettings) SetDepthDataFiltered(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDepthDataFiltered:"), value)
}

// Specifies whether a portrait effects matte should be captured along with
// the photo.
//
// # Discussion
//
// The default is [NO]. Set to [YES] if you wish to receive a portrait effects
// matte with your photo. AVFoundation throws an exception if
// [PortraitEffectsMatteDeliveryEnabled] is not set to [YES], or if your
// delegate doesn’t respond to the
// [CaptureOutputDidFinishProcessingPhotoError] selector.
//
// Setting this property to [YES] doen’t guarantee that a portrait effects
// matte will be present in the resulting [AVCapturePhoto]. The matte is
// primarily used to improve the rendering quality of portrait effects on the
// image. If the photo’s content lacks a clear foreground subject, no
// portrait effects matte is generated, and the property returns `nil`.
// Setting this property to [YES] may add significant processing time to the
// delivery of your [CaptureOutputDidFinishProcessingPhotoError] callback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isPortraitEffectsMatteDeliveryEnabled
func (c AVCapturePhotoSettings) PortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPortraitEffectsMatteDeliveryEnabled:"), value)
}

// Specifies whether the portrait effects matte captured with ths photo should
// be written to the photo’s file structure.
//
// # Discussion
//
// The default is true, which tells AV Foundation to embed the portrait
// effects matte images as HEIF and JPEG in the photo.
//
// This property is ignored if [PortraitEffectsMatteDeliveryEnabled] is set to
// false. AV Foundation includes the portrait effects matte only if both this
// property and [PortraitEffectsMatteDeliveryEnabled] are set to true.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embedsPortraitEffectsMatteInPhoto
func (c AVCapturePhotoSettings) EmbedsPortraitEffectsMatteInPhoto() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("embedsPortraitEffectsMatteInPhoto"))
	return rv
}
func (c AVCapturePhotoSettings) SetEmbedsPortraitEffectsMatteInPhoto(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEmbedsPortraitEffectsMatteInPhoto:"), value)
}

// A Boolean value that specifies whether to write the enabled semantic
// segmentation matte types captured with this photo to the photo’s file
// structure.
//
// # Discussion
//
// Semantic segmentation mattes are only supported in HEIF and JPEG. The photo
// output ignores this property if you set
// [EnabledSemanticSegmentationMatteTypes] to an empty array.
//
// The property’s default value is true.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/embedsSemanticSegmentationMattesInPhoto
func (c AVCapturePhotoSettings) EmbedsSemanticSegmentationMattesInPhoto() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("embedsSemanticSegmentationMattesInPhoto"))
	return rv
}
func (c AVCapturePhotoSettings) SetEmbedsSemanticSegmentationMattesInPhoto(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEmbedsSemanticSegmentationMattesInPhoto:"), value)
}

// An array of semantic segmentation matte types that the photo render
// pipeline can deliver.
//
// # Discussion
//
// You may set this property to the array of matte types you’d like
// delivered with [AVCapturePhoto]. The array may only contain values present
// in [AvailableSemanticSegmentationMatteTypes].
//
// The default value of this property is an empty array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/enabledSemanticSegmentationMatteTypes
func (c AVCapturePhotoSettings) EnabledSemanticSegmentationMatteTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return objc.ConvertSliceToStrings(rv)
}
func (c AVCapturePhotoSettings) SetEnabledSemanticSegmentationMatteTypes(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabledSemanticSegmentationMatteTypes:"), objectivec.StringSliceToNSArray(value))
}

// A dictionary of metadata keys and values to embed in photo file output.
//
// # Discussion
//
// The capture output automatically writes metadata including image
// orientation, EXIF camera properties, and Live Photo metadata, but you can
// override those values or specify additional metadata using the keys and
// values listed in [CGImageProperties]. (Setting this property with any other
// keys raises an exception.)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/metadata
func (c AVCapturePhotoSettings) Metadata() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("metadata"))
	return foundation.NSDictionaryFromID(rv)
}
func (c AVCapturePhotoSettings) SetMetadata(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadata:"), value)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/rawFileFormat
func (c AVCapturePhotoSettings) RawFileFormat() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rawFileFormat"))
	return foundation.NSDictionaryFromID(rv)
}
func (c AVCapturePhotoSettings) SetRawFileFormat(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setRawFileFormat:"), value)
}

// An array of pixel format types compatible with the photo settings for
// delivery of preview-sized images.
//
// # Discussion
//
// To enable delivery of preview-sized images along with the main image
// captured by a photo output, set the [PreviewPhotoFormat] property using one
// of the [kCVPixelBufferPixelFormatTypeKey] values listed in this array.
//
// The order of this array is such that the format requiring the least
// conversion is listed first.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/availablePreviewPhotoPixelFormatTypes-2vfwu
//
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
func (c AVCapturePhotoSettings) AvailablePreviewPhotoPixelFormatTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availablePreviewPhotoPixelFormatTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
