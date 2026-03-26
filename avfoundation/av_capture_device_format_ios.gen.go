// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/foundation"
"github.com/tmc/apple/objectivec"
)

// A Boolean value that indicates whether the format supports a given video
// stabilization mode.
//
// videoStabilizationMode: The stabilization mode to test.
//
// # Return Value
// 
// [true] if video stabilization is supported; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoStabilizationModeSupported(_:)
func (c AVCaptureDeviceFormat) IsVideoStabilizationModeSupported(videoStabilizationMode AVCaptureVideoStabilizationMode) bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoStabilizationModeSupported:"), videoStabilizationMode)
return rv
}
// Indicates the horizontal field of view for an aspect ratio, either
// uncorrected or corrected for geometric distortion.
//
// # Discussion
// 
// A float indicating the field of view for the corresponding
// [AVCaptureDevice.AspectRatio]. Set
// `AVCaptureDevice/geometricDistortionCorrected` to `true` to receive the
// field of view corrected for geometric distortion. If this device format
// does not support dynamic aspect ratio, this function returns `0`.
//
// [AVCaptureDevice.AspectRatio]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AspectRatio
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFieldOfView(for:geometricDistortionCorrected:)
func (c AVCaptureDeviceFormat) VideoFieldOfViewForAspectRatioGeometricDistortionCorrected(aspectRatio objectivec.IObject, geometricDistortionCorrected bool) float32 {
rv := objc.Send[float32](c.ID, objc.Sel("videoFieldOfViewForAspectRatio:geometricDistortionCorrected:"), aspectRatio, geometricDistortionCorrected)
return rv
}

// A Boolean value that indicates whether the format produces video data in a
// binned format.
//
// # Discussion
// 
// Binning is a pixel-combining process which can result in greater low light
// sensitivity at the cost of reduced resolution.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoBinned
func (c AVCaptureDeviceFormat) VideoBinned() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoBinned"))
		return rv
}
// A Boolean value that indicates whether the format supports high dynamic
// range streaming.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoHDRSupported
func (c AVCaptureDeviceFormat) VideoHDRSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoHDRSupported"))
		return rv
}
// A Boolean value that indicates whether a multi-camera capture session
// supports this format.
//
// # Discussion
// 
// When performing single-camera capture using [AVCaptureSession], you may set
// any of the device’s formats as its [ActiveFormat]. However, when using
// [AVCaptureMultiCamSession], you may only set the device’s format to one
// in which [MultiCamSupported] is [true]. Only this limited subset of capture
// formats can run sustainably in a multi-camera capture scenario.
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isMultiCamSupported
func (c AVCaptureDeviceFormat) MultiCamSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isMultiCamSupported"))
		return rv
}
// The list of capture output subclasses not allowed for capture with this
// format, if any.
//
// # Discussion
// 
// As a rule, capture formats with a given [MediaType] are available for use
// with all [AVCaptureOutput] subclasses that accept that media type. However,
// this isn’t always the case. For example, formats for high-resolution
// photo capture may not support the [AVCaptureMovieFileOutput] class due to
// bandwidth limitations.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/unsupportedCaptureOutputClasses
func (c AVCaptureDeviceFormat) UnsupportedCaptureOutputClasses() []objc.Class {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("unsupportedCaptureOutputClasses"))
return objc.ConvertSlice(rv, func(id objc.ID) objc.Class {
	return objc.Class(id)
})
}
// Indicates the format’s horizontal field of view in degrees.
//
// # Discussion
// 
// Returns zero if the format’s field of view is unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoFieldOfView
func (c AVCaptureDeviceFormat) VideoFieldOfView() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("videoFieldOfView"))
		return rv
}
// A horizontal field of view for the format after correction for geometric
// distortion.
//
// # Discussion
// 
// If the capture device doesn’t support geometric distortion correction
// (GDC), the value of this property is equal to the value of
// [VideoFieldOfView].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/geometricDistortionCorrectedVideoFieldOfView
func (c AVCaptureDeviceFormat) GeometricDistortionCorrectedVideoFieldOfView() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("geometricDistortionCorrectedVideoFieldOfView"))
		return rv
}
// A Boolean value that indicates whether this format supports the highest
// photo quality that the platform delivers.
//
// # Discussion
// 
// The simplest way to capture the highest quality photos is to set [photo] as
// your session’s preset. If you’re instead manually setting the capture
// device’s [ActiveFormat] value, select the format whose
// [HighestPhotoQualitySupported] property is [true].
//
// [photo]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/photo
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isHighestPhotoQualitySupported
func (c AVCaptureDeviceFormat) HighestPhotoQualitySupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isHighestPhotoQualitySupported"))
		return rv
}
// A Boolean value that indicates whether the format supports global tone
// mapping.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isGlobalToneMappingSupported
func (c AVCaptureDeviceFormat) GlobalToneMappingSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isGlobalToneMappingSupported"))
		return rv
}
// A floating-point number that indicates the minimum supported exposure ISO
// value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/minISO
func (c AVCaptureDeviceFormat) MinISO() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("minISO"))
		return rv
}
// A floating-point number that indicates the maximum supported exposure ISO
// value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/maxISO
func (c AVCaptureDeviceFormat) MaxISO() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("maxISO"))
		return rv
}
// A time value that indicates the minimum supported exposure duration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/minExposureDuration
func (c AVCaptureDeviceFormat) MinExposureDuration() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("minExposureDuration"))
return objectivec.Object{ID: rv}
}
// A time value that indicates the maximum supported exposure duration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/maxExposureDuration
func (c AVCaptureDeviceFormat) MaxExposureDuration() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("maxExposureDuration"))
return objectivec.Object{ID: rv}
}
// A maximum zoom factor the format allows.
//
// # Discussion
// 
// A maximum factor of `1.0` indicates that the format isn’t capable of
// zooming.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactor
func (c AVCaptureDeviceFormat) VideoMaxZoomFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoMaxZoomFactor"))
		return rv
}
// A threshold at which the system upscales pixel data.
//
// # Discussion
// 
// The device achieves a zoom effect by cropping around the center of the
// image captured by the sensor. At low zoom factors, the cropped images is
// equal to or larger than the output size. At higher zoom factors, the device
// must scale the cropped image up to the output size, resulting in a loss of
// image quality. This property indicates the factors at which upscaling
// occurs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoZoomFactorUpscaleThreshold
func (c AVCaptureDeviceFormat) VideoZoomFactorUpscaleThreshold() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoZoomFactorUpscaleThreshold"))
		return rv
}
// Returns `true` if smart framing is supported by the current format.
//
// # Discussion
// 
// An ultra wide camera device that supports dynamic aspect ratio
// configuration may also support “smart framing monitoring” on particular
// formats.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isSmartFramingSupported
func (c AVCaptureDeviceFormat) SmartFramingSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isSmartFramingSupported"))
		return rv
}
// Indicates the supported aspect ratios for the device format.
//
// # Discussion
// 
// An array that describes the aspect ratios that are supported for this
// format. If this device format does not support dynamic aspect ratio, this
// property returns an empty array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedDynamicAspectRatios
func (c AVCaptureDeviceFormat) SupportedDynamicAspectRatios() []string {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedDynamicAspectRatios"))
return objc.ConvertSliceToStrings(rv)
}
// A Boolean indicating whether the device supports portrait matte effects in
// still-image capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isPortraitEffectsMatteStillImageDeliverySupported
func (c AVCaptureDeviceFormat) PortraitEffectsMatteStillImageDeliverySupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isPortraitEffectsMatteStillImageDeliverySupported"))
		return rv
}
// The list of data formats compatible with this video format.
//
// # Discussion
// 
// Depth data capture requires a compatible pairing of video format and depth
// data format. After you set a capture device’s [ActiveFormat] property to
// this format, you can set the device’s [ActiveDepthDataFormat] property to
// one of the formats in this array.
// 
// Supported depth data formats always match the aspect ratio of their
// corresponding video format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedDepthDataFormats
func (c AVCaptureDeviceFormat) SupportedDepthDataFormats() []AVCaptureDeviceFormat {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedDepthDataFormats"))
return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDeviceFormat {
	return AVCaptureDeviceFormatFromID(id)
})
}
// The highest resolution still image the system can produce for this format.
//
// # Discussion
// 
// Normally, the [AVCaptureStillImageOutput] class emits images with the same
// dimensions as the source [AVCaptureDevice] instance’s [ActiveFormat].
// However, if you set `highResolutionStillImageOutputEnabled` to [true],
// [AVCaptureStillImageOutput] emits still images with its source
// [AVCaptureDevice] instance’s
// `activeFormat.HighResolutionStillImageDimensions()` dimensions.
//
// [AVCaptureStillImageOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/highResolutionStillImageDimensions
func (c AVCaptureDeviceFormat) HighResolutionStillImageDimensions() objectivec.IObject {
rv := objc.Send[objc.ID](c.ID, objc.Sel("highResolutionStillImageDimensions"))
return objectivec.Object{ID: rv}
}
// The zoom factors that a format supports for depth data delivery.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceFormat/supportedVideoZoomFactorsForDepthDataDelivery
func (c AVCaptureDeviceFormat) SupportedVideoZoomFactorsForDepthDataDelivery() []foundation.NSNumber {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedVideoZoomFactorsForDepthDataDelivery"))
return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
	return foundation.NSNumberFromID(id)
})
}
// A maximum zoom factor the device supports when configured for depth data
// delivery.
//
// # Discussion
// 
// Depth data capture requires coordinating the zoom factors of the two
// cameras on a dual-camera device. Therefore, when you enable depth data
// delivery for a capture format using the [AVCaptureDepthDataOutput] class,
// the range of available values for the device’s [VideoZoomFactor] property
// is reduced.
// 
// If this format doesn’t support depth capture, this property’s value is
// the same as that of the [VideoMaxZoomFactor] property.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMaxZoomFactorForDepthDataDelivery
func (c AVCaptureDeviceFormat) VideoMaxZoomFactorForDepthDataDelivery() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoMaxZoomFactorForDepthDataDelivery"))
		return rv
}
// A minimum zoom factor the device supports when configured for depth data
// delivery.
//
// # Discussion
// 
// Depth data capture requires coordinating the zoom factors of the two
// cameras on a dual-camera device. Therefore, when you enable depth data
// delivery for a capture format using the [AVCaptureDepthDataOutput] class,
// the range of available values for the device’s [VideoZoomFactor] property
// is reduced.
// 
// If this format doesn’t support depth capture, this property’s value is
// `1.0`.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/videoMinZoomFactorForDepthDataDelivery
func (c AVCaptureDeviceFormat) VideoMinZoomFactorForDepthDataDelivery() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoMinZoomFactorForDepthDataDelivery"))
		return rv
}
// A Boolean value that indicates whether the format supports video
// stabilization.
//
// # Discussion
// 
// If the format supports video stabilization, you can enable it on an
// [AVCaptureConnection] instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/isVideoStabilizationSupported
func (c AVCaptureDeviceFormat) VideoStabilizationSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoStabilizationSupported"))
		return rv
}

