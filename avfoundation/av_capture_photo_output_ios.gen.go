// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Tells the photo capture output to prepare resources for future capture
// requests with the specified settings.
//
// preparedPhotoSettingsArray: An array of photo capture settings objects indicating the types of capture
// for which the photo output should prepare resources.
//
// completionHandler: A completion block to be called on a serial dispatch queue after the photo
// output has finished preparing resources. Pass `nil` if you do not wish to
// be notified when preparation is complete.
//
// # Discussion
//
// Some types of photo capture, such as bracketed captures and RAW captures,
// require the photo output to allocate additional buffers or prepare other
// resources. To prevent photo capture requests from executing slowly due to
// lazy resource allocation, you may call this method with an array of
// settings objects representative of the types of capture you will be
// performing (such as settings for a bracketed capture, RAW capture, or
// capture with still image stabilization).
//
// By default, the photo output prepares sufficient resources to capture
// photos with default settings (as defined by the [AVCapturePhotoSettings]
// default initializer). To reclaim all possible resources, call this method
// with an empty array.
//
// Each time you call this method with an array of settings, the photo output
// evaluates what additional resources it needs to allocate, as well as
// existing resources that can be reclaimed, and calls your
// `completionHandler` block when it has finished preparing (and possibly
// reclaiming) needed resources. To provide an early hint as to which capture
// features you intend to use, you can call this method even before calling
// the [StartRunning] method on your capture session.
//
// Preparation for photo capture is always optional. You may call the
// [CapturePhotoWithSettingsDelegate] method without first calling this
// method—however, some of your photo captures may execute slowly as the
// capture system allocates additional resources on a just-in-time basis.
//
// If you call this method while your capture session is not running, your
// `completionHandler` block is not immediately called. The photo output calls
// your completion handler only after you’ve called the [StartRunning]
// method on your capture session and the needed resources have actually been
// prepared. If you call the [SetPreparedPhotoSettingsArrayCompletionHandler]
// method with an array of settings, and then call it a second time, the first
// call’s completion handler fires immediately with a `prepared` argument of
// false.
//
// Prepared settings persist across session starts/stops and committed
// configuration changes and participate in the deferred work behavior defined
// by the [AVCaptureSession] class. That is, if you call your capture
// session’s [BeginConfiguration] method, change your session’s
// input/output topology, and then call this method, the capture session
// defers any preparation work until you call the [CommitConfiguration]
// method. This pattern lets you atomically commit a new configuration as well
// as prepare to take photos in that new configuration.
//
// After you call this method and your `completionHandler` block has fired,
// the [PreparedPhotoSettingsArray] property lists the photo settings for
// which the capture system has prepared resources.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/setPreparedPhotoSettingsArray(_:completionHandler:)
func (c AVCapturePhotoOutput) SetPreparedPhotoSettingsArrayCompletionHandler(preparedPhotoSettingsArray []AVCapturePhotoSettings, completionHandler BoolErrorHandler) {
	_block1, _cleanup1 := NewBoolErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](c.ID, objc.Sel("setPreparedPhotoSettingsArray:completionHandler:"), preparedPhotoSettingsArray, _block1)
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedRawPhotoCodecTypes(forRawPhotoPixelFormatType:fileType:)
func (c AVCapturePhotoOutput) SupportedRawPhotoCodecTypesForRawPhotoPixelFormatTypeFileType(pixelFormatType uint32, fileType AVFileType) []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedRawPhotoCodecTypesForRawPhotoPixelFormatType:fileType:"), pixelFormatType, objc.String(string(fileType)))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the list of Bayer RAW pixel formats supported for photo data in the
// specified file type.
//
// fileType: The file type for which to obtain format information.
//
// # Return Value
//
// An array of pixel format types supported for encoding in the specified file
// type.
//
// # Discussion
//
// When you issue a photo capture request, you can separately specify the
// format for capturing or encoding image data and the container format for
// producing output files containing that data. However, each file type
// supports only a specific set of image data types.
//
// After choosing a file type from the [AvailableRawPhotoFileTypes] array, use
// this method to find a compatible image data format before creating a photo
// settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedRawPhotoPixelFormatTypesForFileType:
func (c AVCapturePhotoOutput) SupportedRawPhotoPixelFormatTypesForFileType(fileType AVFileType) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedRawPhotoPixelFormatTypesForFileType:"), objc.String(string(fileType)))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns a Boolean value that indicates whether the pixel format is an Apple
// ProRAW format.
//
// pixelFormat: The pixel format to query.
//
// # Return Value
//
// true if the pixel format is an Apple ProRAW format, otherwise false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAppleProRAWPixelFormat(_:)
func (_AVCapturePhotoOutputClass AVCapturePhotoOutputClass) IsAppleProRAWPixelFormat(pixelFormat uint32) bool {
	rv := objc.Send[bool](objc.ID(_AVCapturePhotoOutputClass.class), objc.Sel("isAppleProRAWPixelFormat:"), pixelFormat)
	return rv
}

// Returns a Boolean value that indicates whether the pixel format is a Bayer
// RAW format.
//
// pixelFormat: The pixel format to query.
//
// # Return Value
//
// true if the pixel format is a Bayer RAW format, otherwise false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isBayerRAWPixelFormat(_:)
func (_AVCapturePhotoOutputClass AVCapturePhotoOutputClass) IsBayerRAWPixelFormat(pixelFormat uint32) bool {
	rv := objc.Send[bool](objc.ID(_AVCapturePhotoOutputClass.class), objc.Sel("isBayerRAWPixelFormat:"), pixelFormat)
	return rv
}

// A Boolean value that indicates the enabled state of automatic deferred
// photo delivery.
//
// # Discussion
//
// Changing this value requires a lengthy reconfiguration of the capture
// pipeline, so you should set this property before calling [StartRunning] on
// the capture session.
//
// Setting this property to true throws an invalid argument exception the
// value of [AutoDeferredPhotoDeliverySupported] is false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAutoDeferredPhotoDeliveryEnabled
func (c AVCapturePhotoOutput) AutoDeferredPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoDeferredPhotoDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetAutoDeferredPhotoDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoDeferredPhotoDeliveryEnabled:"), value)
}

// A Boolean value that indicates whether the photo output supports deferred
// photo delivery.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAutoDeferredPhotoDeliverySupported
func (c AVCapturePhotoOutput) AutoDeferredPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoDeferredPhotoDeliverySupported"))
	return rv
}

// The list of file types currently supported for RAW format capture and
// output.
//
// # Discussion
//
// When you issue a photo capture request, you can separately specify the
// format for capturing or encoding image data and the container format for
// producing output files containing that data. However, each file type
// supports only a specific set of image data types.
//
// After choosing an output file type, use the
// [SupportedRawPhotoPixelFormatTypesForFileType] method to choose an
// appropriate data format before creating a photo settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoFileTypes
func (c AVCapturePhotoOutput) AvailableRawPhotoFileTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableRawPhotoFileTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// A Boolean value that indicates whether the current device and configuration
// supports Apple ProRAW pixel formats.
//
// # Discussion
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAppleProRAWSupported
func (c AVCapturePhotoOutput) AppleProRAWSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAppleProRAWSupported"))
	return rv
}

// A Boolean value that indicates whether you’ve configured the photo output
// to deliver Apple ProRAW formats.
//
// # Discussion
//
// If [AppleProRAWSupported] returns true, you can enable Apple ProRAW capture
// by setting this property to true. Compared to photos taken in Bayer RAW
// format, the system demosaics and partially processes Apple ProRAW photos.
// They’re still scene-referred, however, and allow capturing RAW photos in
// modes that don’t have a traditional Bayer RAW format available, such as
// modes that rely on fusing multiple captures.
//
// Apple ProRAW formats aren’t supported on all platforms and devices. You
// can determine the pixel formats the system supports by querying the
// [availableRawPhotoPixelFormatTypes] property. Use the
// [IsBayerRAWPixelFormat] or [IsAppleProRAWPixelFormat] method to determine
// whether the pixel format is Bayer RAW or Apple ProRAW, respectively.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAppleProRAWEnabled
//
// [availableRawPhotoPixelFormatTypes]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoPixelFormatTypes-9t9k5
func (c AVCapturePhotoOutput) AppleProRAWEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAppleProRAWEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetAppleProRAWEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAppleProRAWEnabled:"), value)
}

// A Boolean value that indicates whether the session’s current
// configuration supports content-aware distortion correction.
//
// # Discussion
//
// Optical design and geometric distortion correction use a rectilinear model
// that preserves lines but not area, angles, or distance. The wider the field
// of view of a lens, the greater the areal distortion along the edges of
// images. Content-aware distortion correction intelligently adjusts its
// behavior to correct distortions based on the photo’s content. For
// example, the algorithm may not apply correction to faces in the center of a
// photo, but may apply it to faces near the photo’s edges.
//
// Switching cameras or formats, or enabling depth data delivery, may result
// in a change to this property value. When the property changes from true to
// false, [ContentAwareDistortionCorrectionEnabled] also reverts to false.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isContentAwareDistortionCorrectionSupported
func (c AVCapturePhotoOutput) ContentAwareDistortionCorrectionSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContentAwareDistortionCorrectionSupported"))
	return rv
}

// A Boolean value that indicates whether the photo render pipeline can
// perform content-aware distortion correction.
//
// # Discussion
//
// You can set this value to true only if
// [ContentAwareDistortionCorrectionSupported] returns true.
//
// Applying distortion correction to preserve natural-looking content may
// result in a small change in the field of view compared to what you see in
// [AVCaptureVideoPreviewLayer]. The amount lost or gained is content specific
// and varies from photo to photo.
//
// Enabling this property requires a lengthy reconfiguration of the capture
// render pipeline, so set this property to true before calling [StartRunning]
// on the capture session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isContentAwareDistortionCorrectionEnabled
func (c AVCapturePhotoOutput) ContentAwareDistortionCorrectionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContentAwareDistortionCorrectionEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetContentAwareDistortionCorrectionEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setContentAwareDistortionCorrectionEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports
// lens stabilization during bracketed image capture.
//
// # Discussion
//
// To make use of optical image stabilization across the entire duration of a
// bracketed capture, set the [isLensStabilizationEnabled] property of your
// bracketed photo settings object.
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLensStabilizationDuringBracketedCaptureSupported
//
// [isLensStabilizationEnabled]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings/isLensStabilizationEnabled
func (c AVCapturePhotoOutput) LensStabilizationDuringBracketedCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLensStabilizationDuringBracketedCaptureSupported"))
	return rv
}

// The maximum number of images that the photo capture output can support in a
// single bracketed capture.
//
// # Discussion
//
// To perform a bracketed capture of multiple images with varied capture
// settings, create a [AVCapturePhotoBracketSettings] instance containing the
// combination of settings and bracketed variations you want. The maximum
// number of photos per capture depends on the size and format of images to be
// captured.
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/maxBracketedCapturePhotoCount
//
// [AVCapturePhotoBracketSettings]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoBracketSettings
func (c AVCapturePhotoOutput) MaxBracketedCapturePhotoCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maxBracketedCapturePhotoCount"))
	return rv
}

// A Boolean value indicating whether the capture output supports automatic
// red-eye reduction.
//
// # Discussion
//
// In iOS 12 and later, AVFoundation applies a red-eye reduction effect only
// when needed. It won’t apply red-eye reduction in the following
// situations:
//
// - When you force a flash capture in good light - When only one eye is
// visible - When doing a bracketed capture - When taking a picture with depth
// on the dual camera
//
// When taking RAW + processed (JPEG or HEIC) still-photo capture with auto
// red-eye reduction enabled, AVFoundation applies correction to only the
// processed photo, the RAW photo.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isAutoRedEyeReductionSupported
func (c AVCapturePhotoOutput) AutoRedEyeReductionSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoRedEyeReductionSupported"))
	return rv
}

// A Boolean value indicating whether the scene currently being previewed by
// the camera warrants use of the flash.
//
// # Discussion
//
// This property’s value changes depending on the scene currently visible to
// the camera. For example, you might use this property to highlight the flash
// control in your app’s camera UI, indicating to the user that the scene is
// dark enough that enabling the flash might be desirable.
//
// If the photo capture output’s [SupportedFlashModes] value is
// [AVCaptureFlashModeOff], this property’s value is always false.
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isFlashScene
func (c AVCapturePhotoOutput) IsFlashScene() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFlashScene"))
	return rv
}

// A photo settings object that controls how the photo output detects and
// handles automatic flash and stabilization modes.
//
// # Discussion
//
// Set the [FlashMode] and [AutoStillImageStabilizationEnabled] properties of
// this photo settings object to influence the values of the photo output’s
// scene monitoring properties ([IsFlashScene] and
// [IsStillImageStabilizationScene]). For example, if you set the [FlashMode]
// property of this photo settings object to [AVCaptureFlashModeOff], the
// photo output’s [IsFlashScene] property reports false regardless of
// lighting conditions in the visible scene. If you set this photo settings
// object’s [FlashMode] property to [AVCaptureFlashModeAuto] or
// [AVCaptureFlashModeOn], the photo output’s [IsFlashScene] property
// reverts to its default behavior of returning true or false based on the
// visible light level.
//
// The default value is an [AVCapturePhotoSettings] object with the following
// settings:
//
// - [FlashMode]: [AVCaptureFlashModeAuto] -
// [AutoStillImageStabilizationEnabled]: true
//
// The photo output ignores all other properties of this photo settings
// object. To control other photo settings when requesting capture, create a
// photo settings object to pass to the [CapturePhotoWithSettingsDelegate]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/photoSettingsForSceneMonitoring
func (c AVCapturePhotoOutput) PhotoSettingsForSceneMonitoring() IAVCapturePhotoSettings {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("photoSettingsForSceneMonitoring"))
	return AVCapturePhotoSettingsFromID(rv)
}
func (c AVCapturePhotoOutput) SetPhotoSettingsForSceneMonitoring(value IAVCapturePhotoSettings) {
	objc.Send[struct{}](c.ID, objc.Sel("setPhotoSettingsForSceneMonitoring:"), value)
}

// A Boolean value that indicates whether the capture output currently
// supports Live Photo capture.
//
// # Discussion
//
// Live Photo captures both a still image and a short movie centered on the
// moment of capture, and the system presents them together in user interfaces
// such as the Photos app.
//
// Not all devices and capture formats support Live Photo capture. This
// property’s value can change if the [SessionPreset] property of the
// current capture session or the [ActiveFormat] property of the underlying
// capture device changes.
//
// When this value changes to false, the [LivePhotoCaptureEnabled]
// property’s value also changes to false. If you previously opted in for
// Live Photo capture and then change configurations, you may need to set
// [LivePhotoCaptureEnabled] to true again.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoCaptureSupported
func (c AVCapturePhotoOutput) LivePhotoCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLivePhotoCaptureSupported"))
	return rv
}

// A Boolean value that indicates whether to configure the capture pipeline
// for Live Photo capture.
//
// # Discussion
//
// This value defaults to false. Changing this value while your session is
// running requires a lengthy reconfiguration of the capture render pipeline.
// If you intend to take any Live Photo captures, set this value to true
// before calling [AVCaptureSession] [StartRunning]. If you change this
// property while the session is running, in-progress Live Photo captures end
// immediately, unfulfilled photo requests cancel, and the video preview
// temporarily freezes.
//
// You must enable this option before initiating a photo capture with the
// [LivePhotoMovieFileURL] property of your photo settings object set to
// non-`nil`. However, after you’ve enabled this option, you can issue photo
// capture requests for both Live Photo captures and still photos.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoCaptureEnabled
func (c AVCapturePhotoOutput) LivePhotoCaptureEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLivePhotoCaptureEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetLivePhotoCaptureEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setLivePhotoCaptureEnabled:"), value)
}

// A Boolean value that indicates whether Live Photo capture is currently in a
// suspended state.
//
// # Discussion
//
// By default, this property’s value is false. Set this value to true to
// stop any current Live Photo movie captures in progress. Doing this prevents
// recording additional actions in the Live Photo movie. For example, if you
// want to capture a still photo that makes a shutter sound, you can prevent
// recording that action.
//
// When you change this value to true, the system trims any Live Photo movie
// captures in progress to the current time. Likewise, when you change this
// value from true to false, subsequent Live Photo movie captures won’t
// contain any earlier recordings.
//
// By default, this property resets to false when the [AVCaptureSession]
// stops. You can prevent this behavior by setting
// [PreservesLivePhotoCaptureSuspendedOnSessionStop] to true before stopping
// the session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoCaptureSuspended
func (c AVCapturePhotoOutput) LivePhotoCaptureSuspended() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLivePhotoCaptureSuspended"))
	return rv
}
func (c AVCapturePhotoOutput) SetLivePhotoCaptureSuspended(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setLivePhotoCaptureSuspended:"), value)
}

// A Boolean value that indicates whether to automatically trim Live Photo
// movie captures to avoid excessive movement.
//
// # Discussion
//
// This value defaults to true when [LivePhotoCaptureSupported] is true.
//
// Use this option to enable the same automatic trimming behavior found in the
// Camera app. By default, a Live Photo capture is about three seconds in
// duration, centered on the time of the capture request. If the user moves
// the camera during capture, iOS analyzes the capture and automatically trims
// the duration of the Live Photo to avoid capturing excess movement.
//
// Changing this value while your session is running requires a lengthy
// reconfiguration of the session. If you intend to take any Live Photo
// captures, set this value to true before calling [AVCaptureSession]
// [StartRunning]. If you change this property while the session is running,
// in-progress Live Photo captures end immediately, unfulfilled photo requests
// cancel, and the video preview temporarily freezes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isLivePhotoAutoTrimmingEnabled
func (c AVCapturePhotoOutput) LivePhotoAutoTrimmingEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLivePhotoAutoTrimmingEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetLivePhotoAutoTrimmingEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setLivePhotoAutoTrimmingEnabled:"), value)
}

// An array of video codecs currently available for Live Photo movie captures.
//
// # Discussion
//
// By default, Live Photo capture encodes the movie portion of a Live Photo
// using the H.264 codec. To use a different codec, set the
// [LivePhotoVideoCodecType] property of your photo settings object to one of
// the values in this array.
//
// The system always presents its default video codec first. If you haven’t
// added the photo output to an [AVCaptureSession] with a video source, no
// codecs are available.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableLivePhotoVideoCodecTypes
func (c AVCapturePhotoOutput) AvailableLivePhotoVideoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableLivePhotoVideoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// A Boolean value indicating whether the capture output currently supports
// depth data capture.
//
// # Discussion
//
// Depth data captures a per-pixel map of scene depth information delivered
// alongside the photo image and optionally embedded in image file output.
// Depth data can be used for purposes such as applying depth-sensitive photo
// filter effects (like that seen in the iOS Camera app’s Portrait mode) and
// performing computer vision tasks.
//
// Not all devices and capture formats support depth capture. This
// property’s value can change if the [SessionPreset] property of the
// current capture session or the [ActiveFormat] property of the underlying
// capture device changes. If a camera or format change causes this
// property’s value to become false, the [DepthDataDeliveryEnabled]
// property’s value also becomes false.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDepthDataDeliverySupported
func (c AVCapturePhotoOutput) DepthDataDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDepthDataDeliverySupported"))
	return rv
}

// A Boolean value that specifies whether to configure the capture pipeline
// for depth data capture.
//
// # Discussion
//
// Depth data captures a per-pixel map of scene depth information delivered
// alongside the photo image and optionally embedded in image file output.
// Depth data can be used for purposes such as applying depth-sensitive photo
// filter effects (like that seen in the iOS Camera app’s Portrait mode) and
// performing computer vision tasks.
//
// Capturing depth data requires that a capture session set up its internal
// rendering pipeline differently. If you intend to capture depth data at all,
// set this property to true before calling the [AVCaptureSession]
// [StartRunning] method. Changing this property while the session is running
// requires a lengthy reconfiguration of the capture render pipeline: Live
// Photo captures in progress will end immediately, unfulfilled photo requests
// will abort, and video preview will temporarily freeze.
//
// You must enable this option before initiating a photo capture with the
// [DepthDataDeliveryEnabled] property of your photo settings object set to
// true. However, after you’ve enabled this option, you are free to issue
// photo capture requests both with and without depth data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDepthDataDeliveryEnabled
func (c AVCapturePhotoOutput) DepthDataDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDepthDataDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetDepthDataDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDepthDataDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output generates a portrait
// effects matte.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isPortraitEffectsMatteDeliveryEnabled
func (c AVCapturePhotoOutput) PortraitEffectsMatteDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPortraitEffectsMatteDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetPortraitEffectsMatteDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPortraitEffectsMatteDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports
// delivery of a portrait effects matte.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isPortraitEffectsMatteDeliverySupported
func (c AVCapturePhotoOutput) PortraitEffectsMatteDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPortraitEffectsMatteDeliverySupported"))
	return rv
}

// # Discussion
//
// A read-only BOOL value indicating whether still image buffers may be
// rotated to match the sensor orientation of earlier generation hardware.
//
// Value is YES for camera configurations which support compensation for the
// sensor orientation, which is applied to HEIC, JPEG, and uncompressed
// processed photos only; compensation is never applied to Bayer RAW or Apple
// ProRaw captures.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isCameraSensorOrientationCompensationSupported
func (c AVCapturePhotoOutput) CameraSensorOrientationCompensationSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCameraSensorOrientationCompensationSupported"))
	return rv
}

// # Discussion
//
// A BOOL value indicating that still image buffers will be rotated to match
// the sensor orientation of earlier generation hardware.
//
// Default is YES when cameraSensorOrientationCompensationSupported is YES.
// Set to NO if your app does not require sensor orientation compensation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isCameraSensorOrientationCompensationEnabled
func (c AVCapturePhotoOutput) CameraSensorOrientationCompensationEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCameraSensorOrientationCompensationEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetCameraSensorOrientationCompensationEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCameraSensorOrientationCompensationEnabled:"), value)
}

// A Boolean value that indicates whether the device supports virtual device
// image fusion.
//
// # Discussion
//
// When using a virtual capture device, the system can fuse the images from
// its constituent cameras to improve image quality.
//
// If the current configuration doesn’t support virtual device fusion, your
// capture requests always resolve [VirtualDeviceFusionEnabled] to false.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isVirtualDeviceFusionSupported
func (c AVCapturePhotoOutput) VirtualDeviceFusionSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVirtualDeviceFusionSupported"))
	return rv
}

// A Boolean value that indicates whether the photo output configuration
// supports delivery of photos from constituent cameras of a virtual device.
//
// # Discussion
//
// The system only supports virtual device constituent photo delivery for
// certain capture session presets and capture device formats.
//
// When switching cameras or formats, this property may change. When this
// property changes from true to false,
// [VirtualDeviceConstituentPhotoDeliveryEnabled] also reverts to false.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isVirtualDeviceConstituentPhotoDeliverySupported
func (c AVCapturePhotoOutput) VirtualDeviceConstituentPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVirtualDeviceConstituentPhotoDeliverySupported"))
	return rv
}

// A Boolean value that indicates whether the photo output delivers photos
// from constituent cameras of a virtual device.
//
// # Discussion
//
// You can only set this value to true when
// [VirtualDeviceConstituentPhotoDeliverySupported] is true.
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isVirtualDeviceConstituentPhotoDeliveryEnabled
func (c AVCapturePhotoOutput) VirtualDeviceConstituentPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVirtualDeviceConstituentPhotoDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetVirtualDeviceConstituentPhotoDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setVirtualDeviceConstituentPhotoDeliveryEnabled:"), value)
}

// An array of photo settings for which the photo output has prepared capture
// resources.
//
// # Discussion
//
// Some types of photo capture, such as bracketed captures and RAW captures,
// require the photo output to allocate additional buffers or prepare other
// resources. To prevent photo capture requests from executing slowly due to
// lazy resource allocation, you may call the
// [SetPreparedPhotoSettingsArrayCompletionHandler] method with an array of
// settings objects representative of the types of capture you will be
// performing (such as settings for a bracketed capture, RAW capture, or
// capture with still image stabilization).
//
// By default, the photo output prepares sufficient resources to capture
// photos with default settings (as defined by the [AVCapturePhotoSettings]
// default initializer).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/preparedPhotoSettingsArray
func (c AVCapturePhotoOutput) PreparedPhotoSettingsArray() []AVCapturePhotoSettings {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("preparedPhotoSettingsArray"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCapturePhotoSettings {
		return AVCapturePhotoSettingsFromID(id)
	})
}

// An array of semantic segmentation matte types that may be captured and
// delivered along with the primary photo.
//
// # Discussion
//
// This property returns the array of semantic segmentation types that’s
// available given the current session configuration.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableSemanticSegmentationMatteTypes
func (c AVCapturePhotoOutput) AvailableSemanticSegmentationMatteTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableSemanticSegmentationMatteTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// The semantic segmentation matte types that the photo render pipeline
// delivers.
//
// # Discussion
//
// Set this property value to the array of matte types you’d like delivered
// with your primary photos. The array may only contain values present in
// [AvailableSemanticSegmentationMatteTypes].
//
// The default value of this property is an empty array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/enabledSemanticSegmentationMatteTypes
func (c AVCapturePhotoOutput) EnabledSemanticSegmentationMatteTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("enabledSemanticSegmentationMatteTypes"))
	return objc.ConvertSliceToStrings(rv)
}
func (c AVCapturePhotoOutput) SetEnabledSemanticSegmentationMatteTypes(value []string) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabledSemanticSegmentationMatteTypes:"), objectivec.StringSliceToNSArray(value))
}

// A Boolean value that indicates whether the photo output currently supports
// the delivery of camera calibration data.
//
// # Discussion
//
// A photo output can deliver camera calibration data only when it’s
// [VirtualDeviceConstituentPhotoDeliveryEnabled] property is `true` and its
// [ContentAwareDistortionCorrectionEnabled] property is `false`.
// Additionally, the source capture device’s
// [GeometricDistortionCorrectionEnabled] property must be `false`.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isCameraCalibrationDataDeliverySupported
func (c AVCapturePhotoOutput) CameraCalibrationDataDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCameraCalibrationDataDeliverySupported"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoCodecTypes
func (c AVCapturePhotoOutput) AvailableRawPhotoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableRawPhotoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// The pixel formats the capture output supports for RAW photo capture.
//
// # Discussion
//
// To capture a photo in RAW format, use the
// [PhotoSettingsWithRawPixelFormatType] or
// [PhotoSettingsWithRawPixelFormatTypeProcessedFormat] initializer to create
// your photo settings object. The value for that initializer’s
// `rawPixelFormatType` parameter must be one of the Bayer RAW format
// identifiers listed in this array.
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availableRawPhotoPixelFormatTypes-5fatm
func (c AVCapturePhotoOutput) AvailableRawPhotoPixelFormatTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableRawPhotoPixelFormatTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// A Boolean value that specifies whether to configure the capture pipeline
// for simultaneous photo capture with both cameras on a dual-camera device.
//
// # Discussion
//
// Enabling this property (on a supported device) allows the capture output to
// deliver separate images from both the wide-angle and telephoto cameras in a
// single capture.
//
// Dual photo delivery requires that a capture session set up its internal
// rendering pipeline differently. If you intend to capture with dual photo
// delivery at all, set this property to true before calling the
// [AVCaptureSession] [StartRunning] method. Changing this property while the
// session is running requires a lengthy reconfiguration of the capture render
// pipeline: Live Photo captures in progress will end immediately, unfulfilled
// photo requests will abort, and video preview will temporarily freeze.
//
// You must enable this option before initiating a photo capture with the
// [DualCameraDualPhotoDeliveryEnabled] property of your photo settings object
// set to true. However, after you’ve enabled this option, you are free to
// issue photo capture requests both with and without dual photo delivery.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDualCameraDualPhotoDeliveryEnabled
func (c AVCapturePhotoOutput) DualCameraDualPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDualCameraDualPhotoDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetDualCameraDualPhotoDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDualCameraDualPhotoDeliveryEnabled:"), value)
}

// A Boolean value indicating whether the capture output currently supports
// simultaneous photo capture with both cameras on a dual-camera device.
//
// # Discussion
//
// Not all devices and capture formats support dual camera capture. This
// property’s value can change if the [SessionPreset] property of the
// current capture session or the [ActiveFormat] property of the underlying
// capture device changes. If a camera or format change causes this
// property’s value to become false, the
// [DualCameraDualPhotoDeliveryEnabled] property’s value also becomes false.
//
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDualCameraDualPhotoDeliverySupported
func (c AVCapturePhotoOutput) DualCameraDualPhotoDeliverySupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDualCameraDualPhotoDeliverySupported"))
	return rv
}

// A Boolean value indicating whether the capture output currently supports
// automatically combining image data on a dual camera device.
//
// # Discussion
//
// On devices equipped with a dual camera, image fusion combines samples from
// both cameras to produce a higher quality image.
//
// To capture a photo with image fusion, set the [AutoDualCameraFusionEnabled]
// property of your photo settings object. If a device does not support image
// fusion, setting the [AutoDualCameraFusionEnabled] property has no effect
// (that is, the resolved [DualCameraFusionEnabled] setting will always be
// false).
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isDualCameraFusionSupported
func (c AVCapturePhotoOutput) DualCameraFusionSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDualCameraFusionSupported"))
	return rv
}

// A Boolean value indicating whether the scene currently being previewed by
// the camera warrants image stabilization.
//
// # Discussion
//
// This property’s value changes depending on the scene currently visible to
// the camera. For example, you might use this property to highlight controls
// in your app’s camera UI related to image stabilization, indicating to the
// user that the scene is dark enough that enabling image stabilization might
// be desirable.
//
// If the photo capture output’s [StillImageStabilizationSupported] value is
// false, this property’s value is always false.
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isStillImageStabilizationScene
func (c AVCapturePhotoOutput) IsStillImageStabilizationScene() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStillImageStabilizationScene"))
	return rv
}

// A Boolean value indicating whether the capture output currently supports
// automatic stabilization for still image capture.
//
// # Discussion
//
// To capture a photo with image stabilization, set the
// [AutoStillImageStabilizationEnabled] property of your photo settings
// object. Automatic stabilization always includes digital image
// stabilization, and may also include optical lens stabilization, based on
// the current device. If a device does not support still image stabilization,
// set the [AutoStillImageStabilizationEnabled] property has no effect (that
// is, the resolved [StillImageStabilizationEnabled] setting will always be
// false).
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isStillImageStabilizationSupported
func (c AVCapturePhotoOutput) StillImageStabilizationSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStillImageStabilizationSupported"))
	return rv
}
