// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCapturePhotoOutput] class.
var (
	_AVCapturePhotoOutputClass     AVCapturePhotoOutputClass
	_AVCapturePhotoOutputClassOnce sync.Once
)

func getAVCapturePhotoOutputClass() AVCapturePhotoOutputClass {
	_AVCapturePhotoOutputClassOnce.Do(func() {
		_AVCapturePhotoOutputClass = AVCapturePhotoOutputClass{class: objc.GetClass("AVCapturePhotoOutput")}
	})
	return _AVCapturePhotoOutputClass
}

// GetAVCapturePhotoOutputClass returns the class object for AVCapturePhotoOutput.
func GetAVCapturePhotoOutputClass() AVCapturePhotoOutputClass {
	return getAVCapturePhotoOutputClass()
}

type AVCapturePhotoOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCapturePhotoOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCapturePhotoOutputClass) Alloc() AVCapturePhotoOutput {
	rv := objc.Send[AVCapturePhotoOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A capture output for still image, Live Photos, and other photography
// workflows.
//
// # Overview
// 
// [AVCapturePhotoOutput] provides an interface for capture workflows related
// to still photography. In addition to basic capture of still images, a photo
// output supports RAW-format capture, bracketed capture of multiple images,
// Live Photos, and wide-gamut color. You can output captured photos in a
// variety of formats and codecs, including RAW format DNG files, HEVC format
// HEIF files, and JPEG files.
// 
// To capture photos with the [AVCapturePhotoOutput] class, follow these
// steps:
// 
// - Create an [AVCapturePhotoOutput] object. Use its properties to determine
// supported capture settings and to enable certain features (for example,
// whether to capture Live Photos). - Create and configure an
// [AVCapturePhotoSettings] object to choose features and settings for a
// specific capture (for example, whether to enable image stabilization or
// flash). - Capture an image by passing your photo settings object to the
// [AVCapturePhotoOutput.CapturePhotoWithSettingsDelegate] method along with a delegate object
// implementing the [AVCapturePhotoCaptureDelegate] protocol. The photo
// capture output then calls your delegate to notify you of significant events
// during the capture process.
// 
// Some photo capture settings, such as the [AVCapturePhotoOutput.FlashMode] property, include
// options for automatic behavior. For such settings, the photo output
// determines whether to use that feature at the moment of capture—you
// don’t know when requesting a capture whether the feature will be enabled
// when the capture completes. When the photo capture output calls your
// [AVCapturePhotoCaptureDelegate] methods with information about the
// completed or in-progress capture, it also provides an
// [AVCaptureResolvedPhotoSettings] object that details which automatic
// features are set for that capture. The resolved settings object’s
// [AVCapturePhotoOutput.UniqueID] property matches the [AVCapturePhotoOutput.UniqueID] value of the
// [AVCapturePhotoSettings] object you used to request capture.
// 
// Enabling certain photo features (Live Photo capture and high resolution
// capture) requires a reconfiguration of the capture render pipeline. To opt
// into these features, set the [isHighResolutionCaptureEnabled],
// [AVCapturePhotoOutput.LivePhotoCaptureEnabled], and [AVCapturePhotoOutput.LivePhotoAutoTrimmingEnabled] properties
// before calling your [AVCaptureSession] object’s [StartRunning] method.
// Changing any of these properties while the session is running disrupts the
// capture render pipeline: Live Photo captures in progress end immediately,
// unfulfilled photo requests abort, and video preview temporarily freezes.
// 
// Using a photo capture output adds other requirements to your
// [AVCaptureSession] object:
// 
// - A capture session can’t support both Live Photo capture and movie file
// output. If your capture session includes an [AVCaptureMovieFileOutput]
// object, the [AVCapturePhotoOutput.LivePhotoCaptureSupported] property becomes [false]. (As an
// alternative, you can use the [AVCaptureVideoDataOutput] class to output
// video buffers at the same resolution as a simultaneous Live Photo capture).
// - A capture session can’t contain both an [AVCapturePhotoOutput] object
// and an [AVCaptureStillImageOutput] object. The [AVCapturePhotoOutput] class
// includes all functionality of (and deprecates) the
// [AVCaptureStillImageOutput] class.
// 
// The [AVCapturePhotoOutput] class implicitly supports wide-gamut color
// photography. If the source [AVCaptureDevice] object’s [AVCapturePhotoOutput.ActiveColorSpace]
// value is [CaptureColorSpace_P3_D65], the capture output produces photos
// with wide color information (unless your [AVCapturePhotoSettings] object
// specifies an output format that doesn’t support wide color).
//
// [AVCaptureStillImageOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureStillImageOutput
// [false]: https://developer.apple.com/documentation/Swift/false
// [isHighResolutionCaptureEnabled]: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isHighResolutionCaptureEnabled
//
// # Capturing a photo
//
//   - [AVCapturePhotoOutput.CapturePhotoWithSettingsDelegate]: Initiates a photo capture using the specified settings.
//
// # Managing responsive capture
//
//   - [AVCapturePhotoOutput.CaptureReadiness]: A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.
//   - [AVCapturePhotoOutput.FastCapturePrioritizationSupported]: A Boolean value that indicates whether the photo output supports fast capture prioritization.
//   - [AVCapturePhotoOutput.SetFastCapturePrioritizationSupported]
//   - [AVCapturePhotoOutput.FastCapturePrioritizationEnabled]: A Boolean value that indicates whether the output enables fast capture prioritization.
//   - [AVCapturePhotoOutput.SetFastCapturePrioritizationEnabled]
//   - [AVCapturePhotoOutput.ResponsiveCaptureSupported]: A Boolean value that indicates whether the photo output supports responsive capture.
//   - [AVCapturePhotoOutput.ResponsiveCaptureEnabled]: A Boolean value that indicates whether the photo output configuration enables responsive capture.
//   - [AVCapturePhotoOutput.SetResponsiveCaptureEnabled]
//   - [AVCapturePhotoOutput.ZeroShutterLagSupported]: A Boolean value that indicates whether the photo output supports zero shutter lag.
//   - [AVCapturePhotoOutput.ZeroShutterLagEnabled]: A Boolean value that indicates whether the photo output configuration enables zero shutter lag.
//   - [AVCapturePhotoOutput.SetZeroShutterLagEnabled]
//
// # Determining supported codec types
//
//   - [AVCapturePhotoOutput.AvailablePhotoCodecTypes]: The compression codecs this capture output currently supports for photo capture.
//   - [AVCapturePhotoOutput.SupportedPhotoCodecTypesForFileType]: Returns the list of photo codecs (such as JPEG or HEVC) supported for photo data in the specified file type.
//
// # Determining supported file types
//
//   - [AVCapturePhotoOutput.AvailablePhotoFileTypes]: The list of file types currently supported for photo capture and output.
//
// # Suppressing the shutter sound
//
//   - [AVCapturePhotoOutput.ShutterSoundSuppressionSupported]: A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.
//
// # Configuring high-resolution still capture
//
//   - [AVCapturePhotoOutput.MaxPhotoDimensions]: The maximum resolution of the requested photo.
//   - [AVCapturePhotoOutput.SetMaxPhotoDimensions]
//
// # Configuring Live Photo capture
//
//   - [AVCapturePhotoOutput.PreservesLivePhotoCaptureSuspendedOnSessionStop]: A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.
//   - [AVCapturePhotoOutput.SetPreservesLivePhotoCaptureSuspendedOnSessionStop]
//
// # Configuring Portrait Effects matte capture
//
//   - [AVCapturePhotoOutput.PortraitEffectsMatte]: The portrait effects matte captured with the photo.
//   - [AVCapturePhotoOutput.SetPortraitEffectsMatte]
//
// # Configuring constant color
//
//   - [AVCapturePhotoOutput.ConstantColorSupported]: A Boolean value that indicates whether a photo output supports constant color capture.
//   - [AVCapturePhotoOutput.ConstantColorEnabled]: A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//   - [AVCapturePhotoOutput.SetConstantColorEnabled]
//
// # Setting the capture prioritization
//
//   - [AVCapturePhotoOutput.MaxPhotoQualityPrioritization]: The highest quality the photo output should prepare to deliver on a capture-by-capture basis.
//   - [AVCapturePhotoOutput.SetMaxPhotoQualityPrioritization]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput
type AVCapturePhotoOutput struct {
	AVCaptureOutput
}

// AVCapturePhotoOutputFromID constructs a [AVCapturePhotoOutput] from an objc.ID.
//
// A capture output for still image, Live Photos, and other photography
// workflows.
func AVCapturePhotoOutputFromID(id objc.ID) AVCapturePhotoOutput {
	return AVCapturePhotoOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}
// NOTE: AVCapturePhotoOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCapturePhotoOutput] class.
//
// # Capturing a photo
//
//   - [IAVCapturePhotoOutput.CapturePhotoWithSettingsDelegate]: Initiates a photo capture using the specified settings.
//
// # Managing responsive capture
//
//   - [IAVCapturePhotoOutput.CaptureReadiness]: A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.
//   - [IAVCapturePhotoOutput.FastCapturePrioritizationSupported]: A Boolean value that indicates whether the photo output supports fast capture prioritization.
//   - [IAVCapturePhotoOutput.SetFastCapturePrioritizationSupported]
//   - [IAVCapturePhotoOutput.FastCapturePrioritizationEnabled]: A Boolean value that indicates whether the output enables fast capture prioritization.
//   - [IAVCapturePhotoOutput.SetFastCapturePrioritizationEnabled]
//   - [IAVCapturePhotoOutput.ResponsiveCaptureSupported]: A Boolean value that indicates whether the photo output supports responsive capture.
//   - [IAVCapturePhotoOutput.ResponsiveCaptureEnabled]: A Boolean value that indicates whether the photo output configuration enables responsive capture.
//   - [IAVCapturePhotoOutput.SetResponsiveCaptureEnabled]
//   - [IAVCapturePhotoOutput.ZeroShutterLagSupported]: A Boolean value that indicates whether the photo output supports zero shutter lag.
//   - [IAVCapturePhotoOutput.ZeroShutterLagEnabled]: A Boolean value that indicates whether the photo output configuration enables zero shutter lag.
//   - [IAVCapturePhotoOutput.SetZeroShutterLagEnabled]
//
// # Determining supported codec types
//
//   - [IAVCapturePhotoOutput.AvailablePhotoCodecTypes]: The compression codecs this capture output currently supports for photo capture.
//   - [IAVCapturePhotoOutput.SupportedPhotoCodecTypesForFileType]: Returns the list of photo codecs (such as JPEG or HEVC) supported for photo data in the specified file type.
//
// # Determining supported file types
//
//   - [IAVCapturePhotoOutput.AvailablePhotoFileTypes]: The list of file types currently supported for photo capture and output.
//
// # Suppressing the shutter sound
//
//   - [IAVCapturePhotoOutput.ShutterSoundSuppressionSupported]: A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.
//
// # Configuring high-resolution still capture
//
//   - [IAVCapturePhotoOutput.MaxPhotoDimensions]: The maximum resolution of the requested photo.
//   - [IAVCapturePhotoOutput.SetMaxPhotoDimensions]
//
// # Configuring Live Photo capture
//
//   - [IAVCapturePhotoOutput.PreservesLivePhotoCaptureSuspendedOnSessionStop]: A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.
//   - [IAVCapturePhotoOutput.SetPreservesLivePhotoCaptureSuspendedOnSessionStop]
//
// # Configuring Portrait Effects matte capture
//
//   - [IAVCapturePhotoOutput.PortraitEffectsMatte]: The portrait effects matte captured with the photo.
//   - [IAVCapturePhotoOutput.SetPortraitEffectsMatte]
//
// # Configuring constant color
//
//   - [IAVCapturePhotoOutput.ConstantColorSupported]: A Boolean value that indicates whether a photo output supports constant color capture.
//   - [IAVCapturePhotoOutput.ConstantColorEnabled]: A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
//   - [IAVCapturePhotoOutput.SetConstantColorEnabled]
//
// # Setting the capture prioritization
//
//   - [IAVCapturePhotoOutput.MaxPhotoQualityPrioritization]: The highest quality the photo output should prepare to deliver on a capture-by-capture basis.
//   - [IAVCapturePhotoOutput.SetMaxPhotoQualityPrioritization]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput
type IAVCapturePhotoOutput interface {
	IAVCaptureOutput

	// Topic: Capturing a photo

	// Initiates a photo capture using the specified settings.
	CapturePhotoWithSettingsDelegate(settings IAVCapturePhotoSettings, delegate AVCapturePhotoCaptureDelegate)

	// Topic: Managing responsive capture

	// A value that specifies whether the photo output is ready to respond to new capture requests in a timely manner.
	CaptureReadiness() AVCapturePhotoOutputCaptureReadiness
	// A Boolean value that indicates whether the photo output supports fast capture prioritization.
	FastCapturePrioritizationSupported() bool
	SetFastCapturePrioritizationSupported(value bool)
	// A Boolean value that indicates whether the output enables fast capture prioritization.
	FastCapturePrioritizationEnabled() bool
	SetFastCapturePrioritizationEnabled(value bool)
	// A Boolean value that indicates whether the photo output supports responsive capture.
	ResponsiveCaptureSupported() bool
	// A Boolean value that indicates whether the photo output configuration enables responsive capture.
	ResponsiveCaptureEnabled() bool
	SetResponsiveCaptureEnabled(value bool)
	// A Boolean value that indicates whether the photo output supports zero shutter lag.
	ZeroShutterLagSupported() bool
	// A Boolean value that indicates whether the photo output configuration enables zero shutter lag.
	ZeroShutterLagEnabled() bool
	SetZeroShutterLagEnabled(value bool)

	// Topic: Determining supported codec types

	// The compression codecs this capture output currently supports for photo capture.
	AvailablePhotoCodecTypes() []string
	// Returns the list of photo codecs (such as JPEG or HEVC) supported for photo data in the specified file type.
	SupportedPhotoCodecTypesForFileType(fileType AVFileType) []string

	// Topic: Determining supported file types

	// The list of file types currently supported for photo capture and output.
	AvailablePhotoFileTypes() []string

	// Topic: Suppressing the shutter sound

	// A Boolean value that indicates whether the photo output supports suppressing the system shutter sound.
	ShutterSoundSuppressionSupported() bool

	// Topic: Configuring high-resolution still capture

	// The maximum resolution of the requested photo.
	MaxPhotoDimensions() coremedia.CMVideoDimensions
	SetMaxPhotoDimensions(value coremedia.CMVideoDimensions)

	// Topic: Configuring Live Photo capture

	// A Boolean value that indicates whether to preserve the suspended state of Live Photo capture when the session stops.
	PreservesLivePhotoCaptureSuspendedOnSessionStop() bool
	SetPreservesLivePhotoCaptureSuspendedOnSessionStop(value bool)

	// Topic: Configuring Portrait Effects matte capture

	// The portrait effects matte captured with the photo.
	PortraitEffectsMatte() IAVPortraitEffectsMatte
	SetPortraitEffectsMatte(value IAVPortraitEffectsMatte)

	// Topic: Configuring constant color

	// A Boolean value that indicates whether a photo output supports constant color capture.
	ConstantColorSupported() bool
	// A Boolean value that indicates whether the photo output configures the render pipeline to perform constant color capture.
	ConstantColorEnabled() bool
	SetConstantColorEnabled(value bool)

	// Topic: Setting the capture prioritization

	// The highest quality the photo output should prepare to deliver on a capture-by-capture basis.
	MaxPhotoQualityPrioritization() AVCapturePhotoQualityPrioritization
	SetMaxPhotoQualityPrioritization(value AVCapturePhotoQualityPrioritization)

	// The currently active color space for capture.
	ActiveColorSpace() AVCaptureColorSpace
	SetActiveColorSpace(value AVCaptureColorSpace)
	// The pixel formats the capture output supports for photo capture.
	AvailablePhotoPixelFormatTypes() []foundation.NSNumber
	// A setting for whether to fire the flash when capturing photos.
	FlashMode() AVCaptureFlashMode
	SetFlashMode(value AVCaptureFlashMode)
	// The flash settings this capture output currently supports.
	SupportedFlashModes() []foundation.NSNumber
	// A unique identifier for this photo settings instance.
	UniqueID() objectivec.IObject
	SetUniqueID(value objectivec.IObject)
	// Returns the list of uncompressed pixel formats supported for photo data in the specified file type.
	SupportedPhotoPixelFormatTypesForFileType(fileType AVFileType) []foundation.NSNumber
}

// Init initializes the instance.
func (c AVCapturePhotoOutput) Init() AVCapturePhotoOutput {
	rv := objc.Send[AVCapturePhotoOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCapturePhotoOutput) Autorelease() AVCapturePhotoOutput {
	rv := objc.Send[AVCapturePhotoOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCapturePhotoOutput creates a new AVCapturePhotoOutput instance.
func NewAVCapturePhotoOutput() AVCapturePhotoOutput {
	class := getAVCapturePhotoOutputClass()
	rv := objc.Send[AVCapturePhotoOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initiates a photo capture using the specified settings.
//
// settings: The settings for the photo capture, such as the output pixel format and
// flash mode. This method copies the provided [AVCapturePhotoSettings]
// object, so future changes to that object do not affect the capture in
// progress.
//
// delegate: A delegate object to receive messages about capture progress and results.
// The photo output calls your delegate methods as the photo advances from
// capture to processing to delivery of finished images.
//
// # Discussion
// 
// Use this method for all variations of still photography, including single
// photo capture, RAW format capture (with or without a secondary format such
// as JPEG), bracketed capture of multiple images, and Live Photo capture.
// 
// When you call this method, the photo output validates the properties of
// your `settings` object to ensure deterministic behavior. For example, the
// [FlashMode] setting must specify a value that is present in the photo
// output’s [SupportedFlashModes] array. See each property’s description
// in the [AVCapturePhotoSettings] class reference for detailed validation
// rules.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/capturePhoto(with:delegate:)
func (c AVCapturePhotoOutput) CapturePhotoWithSettingsDelegate(settings IAVCapturePhotoSettings, delegate AVCapturePhotoCaptureDelegate) {
	objc.Send[objc.ID](c.ID, objc.Sel("capturePhotoWithSettings:delegate:"), settings, delegate)
}
// Returns the list of photo codecs (such as JPEG or HEVC) supported for photo
// data in the specified file type.
//
// fileType: The file type (such as JFIF or HEIF) for which to obtain codec information.
//
// # Return Value
// 
// An array of video codec types supported for encoding in the specified file
// type.
//
// # Discussion
// 
// When you issue a photo capture request, you can separately specify the
// format for capturing or encoding image data and the container format for
// producing output files containing that data. However, each file type
// supports only a specific set of image data types.
// 
// After choosing a file type from the [AvailablePhotoFileTypes] array, use
// this method to find a compatible image data codec before creating a photo
// settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedPhotoCodecTypes(for:)
func (c AVCapturePhotoOutput) SupportedPhotoCodecTypesForFileType(fileType AVFileType) []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedPhotoCodecTypesForFileType:"), objc.String(string(fileType)))
	return objc.ConvertSliceToStrings(rv)
}
// Returns the list of uncompressed pixel formats supported for photo data in
// the specified file type.
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
// After choosing a file type from the [AvailablePhotoFileTypes] array, use
// this method to find a compatible image data format before creating a photo
// settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedPhotoPixelFormatTypesForFileType:
func (c AVCapturePhotoOutput) SupportedPhotoPixelFormatTypesForFileType(fileType AVFileType) []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedPhotoPixelFormatTypesForFileType:"), objc.String(string(fileType)))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// A value that specifies whether the photo output is ready to respond to new
// capture requests in a timely manner.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/captureReadiness-swift.property
func (c AVCapturePhotoOutput) CaptureReadiness() AVCapturePhotoOutputCaptureReadiness {
	rv := objc.Send[AVCapturePhotoOutputCaptureReadiness](c.ID, objc.Sel("captureReadiness"))
	return AVCapturePhotoOutputCaptureReadiness(rv)
}
// A Boolean value that indicates whether the photo output supports fast
// capture prioritization.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isFastCapturePrioritizationSupported
func (c AVCapturePhotoOutput) FastCapturePrioritizationSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFastCapturePrioritizationSupported"))
	return rv
}
func (c AVCapturePhotoOutput) SetFastCapturePrioritizationSupported(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setFastCapturePrioritizationSupported:"), value)
}
// A Boolean value that indicates whether the output enables fast capture
// prioritization.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isFastCapturePrioritizationEnabled
func (c AVCapturePhotoOutput) FastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFastCapturePrioritizationEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetFastCapturePrioritizationEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setFastCapturePrioritizationEnabled:"), value)
}
// A Boolean value that indicates whether the photo output supports responsive
// capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isResponsiveCaptureSupported
func (c AVCapturePhotoOutput) ResponsiveCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isResponsiveCaptureSupported"))
	return rv
}
// A Boolean value that indicates whether the photo output configuration
// enables responsive capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isResponsiveCaptureEnabled
func (c AVCapturePhotoOutput) ResponsiveCaptureEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isResponsiveCaptureEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetResponsiveCaptureEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setResponsiveCaptureEnabled:"), value)
}
// A Boolean value that indicates whether the photo output supports zero
// shutter lag.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isZeroShutterLagSupported
func (c AVCapturePhotoOutput) ZeroShutterLagSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isZeroShutterLagSupported"))
	return rv
}
// A Boolean value that indicates whether the photo output configuration
// enables zero shutter lag.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isZeroShutterLagEnabled
func (c AVCapturePhotoOutput) ZeroShutterLagEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isZeroShutterLagEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetZeroShutterLagEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setZeroShutterLagEnabled:"), value)
}
// The compression codecs this capture output currently supports for photo
// capture.
//
// # Discussion
// 
// To capture a photo in a compressed format, such as JPEG, use the
// [PhotoSettingsWithFormat] initializer to create your photo settings object.
// In that initializer’s `format` dictionary, pass the key
// [AVVideoCodecKey], whose value must be one of the codec identifiers listed
// in this array.
// 
// This property supports key-value observing.
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availablePhotoCodecTypes
func (c AVCapturePhotoOutput) AvailablePhotoCodecTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availablePhotoCodecTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// The list of file types currently supported for photo capture and output.
//
// # Discussion
// 
// When you issue a photo capture request, you can separately specify the
// format for capturing or encoding image data and the container format for
// producing output files containing that data. However, each file type
// supports only a specific set of image data types.
// 
// After choosing an output file type, use the
// [SupportedPhotoCodecTypesForFileType] (for capture in compressed formats
// such as HEVC and JPEG) or [SupportedPhotoPixelFormatTypesForFileType] (for
// capture in uncompressed formats such as TIFF) method to choose an
// appropriate data format before creating a photo settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availablePhotoFileTypes
func (c AVCapturePhotoOutput) AvailablePhotoFileTypes() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availablePhotoFileTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// A Boolean value that indicates whether the photo output supports
// suppressing the system shutter sound.
//
// # Discussion
// 
// In iOS, the value is [true], except in jurisdictions where you can’t
// disable the shutter sound. On all other platforms, the value is always
// [false].
// 
// If the output supports this feature, you can supress the shutter sound when
// capturing a photo using the [ShutterSoundSuppressionEnabled] property of
// [AVCapturePhotoSettings].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isShutterSoundSuppressionSupported
func (c AVCapturePhotoOutput) ShutterSoundSuppressionSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isShutterSoundSuppressionSupported"))
	return rv
}
// The maximum resolution of the requested photo.
//
// # Discussion
// 
// Set a value for this property to request images up to the specified
// dimensions. Images that a photo output returns may be smaller than the
// dimensions, but are never be larger. Once set, you can request images with
// any valid maximum photo dimensions by setting [MaxPhotoDimensions] on
// [AVCapturePhotoSettings] on a per photo basis.
// 
// The dimensions you set must match one returned by
// [supportedMaxPhotoDimensions] for the current active format.
//
// [supportedMaxPhotoDimensions]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedMaxPhotoDimensions
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/maxPhotoDimensions
func (c AVCapturePhotoOutput) MaxPhotoDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("maxPhotoDimensions"))
	return coremedia.CMVideoDimensions(rv)
}
func (c AVCapturePhotoOutput) SetMaxPhotoDimensions(value coremedia.CMVideoDimensions) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}
// A Boolean value that indicates whether to preserve the suspended state of
// Live Photo capture when the session stops.
//
// # Discussion
// 
// This value defaults to [false], which means that Live Photo capture resumes
// when the session stops. Set the value to [true] to save the state of the
// [LivePhotoCaptureSuspended] property across session restarts.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/preservesLivePhotoCaptureSuspendedOnSessionStop
func (c AVCapturePhotoOutput) PreservesLivePhotoCaptureSuspendedOnSessionStop() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("preservesLivePhotoCaptureSuspendedOnSessionStop"))
	return rv
}
func (c AVCapturePhotoOutput) SetPreservesLivePhotoCaptureSuspendedOnSessionStop(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreservesLivePhotoCaptureSuspendedOnSessionStop:"), value)
}
// The portrait effects matte captured with the photo.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturephoto/portraiteffectsmatte
func (c AVCapturePhotoOutput) PortraitEffectsMatte() IAVPortraitEffectsMatte {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("portraitEffectsMatte"))
	return AVPortraitEffectsMatteFromID(objc.ID(rv))
}
func (c AVCapturePhotoOutput) SetPortraitEffectsMatte(value IAVPortraitEffectsMatte) {
	objc.Send[struct{}](c.ID, objc.Sel("setPortraitEffectsMatte:"), value)
}
// A Boolean value that indicates whether a photo output supports constant
// color capture.
//
// # Discussion
// 
// The light sources illuminating a scene affect an object’s color in a
// photograph, so the color of the same object captured in warm light might
// look significantly different than in colder light. In some cases, ambient
// light-induced color variation is undesirable, and you may prefer an
// estimate of what these materials would look like under a standard light
// such as daylight (D65), regardless of the lighting conditions at the time
// of capture. Some devices are capable of producing such constant color
// photos.
// 
// Constant color captures require the flash to fire and may require a
// pre-flash sequence to determine the correct focus and exposure, therefore
// it might take several seconds to acquire a constant color photo. Due to
// this flash requirement, you can only take a constant color capture with
// [CaptureFlashModeAuto] or [CaptureFlashModeOn] as the flash mode, otherwise
// the system throws an exception.
// 
// You can only achieve constant color when the flash has a discernible effect
// on the scene, so it may not perform well in bright conditions such as
// direct sunlight. Use the photo’s [ConstantColorConfidenceMap] property to
// examine the confidence level, and therefore the usefulness, of each region
// of a constant color photo.
// 
// The property value is [true] if the session’s current configuration
// allows the output to capture photos with constant color.
// 
// When switching cameras or formats this property may change. When this
// property changes from [true] to [false], the value of
// [ConstantColorEnabled] also reverts to [false]. If you’ve previously
// opted in to constant color and then change configurations, you may need to
// set the value of [ConstantColorEnabled] to [true] again.
// 
// This property is key-value observable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isConstantColorSupported
func (c AVCapturePhotoOutput) ConstantColorSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isConstantColorSupported"))
	return rv
}
// A Boolean value that indicates whether the photo output configures the
// render pipeline to perform constant color capture.
//
// # Discussion
// 
// The default value is [false]. Set the value to [true] to enable support for
// taking constant color photos. You can only enable constant color capture if
// the value of [ConstantColorSupported] is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/isConstantColorEnabled
func (c AVCapturePhotoOutput) ConstantColorEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}
func (c AVCapturePhotoOutput) SetConstantColorEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setConstantColorEnabled:"), value)
}
// The highest quality the photo output should prepare to deliver on a
// capture-by-capture basis.
//
// # Discussion
// 
// [AVCapturePhotoOutput] can apply a variety of techniques to improve photo
// quality, such as reducing noise, preserving detail in low light, freezing
// motion, and so on. Some techniques improve image quality at the expense of
// the shot-to-shot time. Before starting your session, you may set this
// property to indicate the highest quality prioritization you intend to
// request when calling the [CapturePhotoWithSettingsDelegate] method.
// 
// When configuring an [AVCapturePhotoSettings] object, you can’t exceed
// this quality prioritization level, but you may select a lower
// prioritization level that favors speed over quality.
// 
// When you attach the photo output to an [AVCaptureSession], the default
// value of this property is [CapturePhotoQualityPrioritizationBalanced]. If
// you instead attach it to an [AVCaptureMultiCamSession], the default value
// is [CapturePhotoQualityPrioritizationSpeed].
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/maxPhotoQualityPrioritization
func (c AVCapturePhotoOutput) MaxPhotoQualityPrioritization() AVCapturePhotoQualityPrioritization {
	rv := objc.Send[AVCapturePhotoQualityPrioritization](c.ID, objc.Sel("maxPhotoQualityPrioritization"))
	return AVCapturePhotoQualityPrioritization(rv)
}
func (c AVCapturePhotoOutput) SetMaxPhotoQualityPrioritization(value AVCapturePhotoQualityPrioritization) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxPhotoQualityPrioritization:"), value)
}
// The currently active color space for capture.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/activecolorspace
func (c AVCapturePhotoOutput) ActiveColorSpace() AVCaptureColorSpace {
	rv := objc.Send[AVCaptureColorSpace](c.ID, objc.Sel("activeColorSpace"))
	return AVCaptureColorSpace(rv)
}
func (c AVCapturePhotoOutput) SetActiveColorSpace(value AVCaptureColorSpace) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveColorSpace:"), value)
}
// The pixel formats the capture output supports for photo capture.
//
// # Discussion
// 
// To capture a photo in an uncompressed format, such as 420f, 420v, or BGRA,
// use the [PhotoSettingsWithFormat] initializer to create your photo settings
// object. In that initializer’s `format` dictionary, pass the key
// [kCVPixelBufferPixelFormatTypeKey], whose value must be one of the pixel
// format identifiers listed in this array.
// 
// This property supports key-value observing.
//
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/availablePhotoPixelFormatTypes-6eyb
func (c AVCapturePhotoOutput) AvailablePhotoPixelFormatTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availablePhotoPixelFormatTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// A setting for whether to fire the flash when capturing photos.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c AVCapturePhotoOutput) FlashMode() AVCaptureFlashMode {
	rv := objc.Send[AVCaptureFlashMode](c.ID, objc.Sel("flashMode"))
	return AVCaptureFlashMode(rv)
}
func (c AVCapturePhotoOutput) SetFlashMode(value AVCaptureFlashMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setFlashMode:"), value)
}
// The flash settings this capture output currently supports.
//
// # Discussion
// 
// To set the flash mode for a capture, set the [FlashMode] property of your
// photo settings object to one of the [AVCaptureDevice.FlashMode] values
// listed in this array.
// 
// This property supports key-value observing.
//
// [AVCaptureDevice.FlashMode]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FlashMode-swift.enum
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/supportedFlashModes-4u69s
func (c AVCapturePhotoOutput) SupportedFlashModes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedFlashModes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// A unique identifier for this photo settings instance.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/uniqueid
func (c AVCapturePhotoOutput) UniqueID() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("uniqueID"))
	return objectivec.Object{ID: rv}
}
func (c AVCapturePhotoOutput) SetUniqueID(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setUniqueID:"), value)
}

