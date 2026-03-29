// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCapturePhotoSettings] class.
var (
	_AVCapturePhotoSettingsClass     AVCapturePhotoSettingsClass
	_AVCapturePhotoSettingsClassOnce sync.Once
)

func getAVCapturePhotoSettingsClass() AVCapturePhotoSettingsClass {
	_AVCapturePhotoSettingsClassOnce.Do(func() {
		_AVCapturePhotoSettingsClass = AVCapturePhotoSettingsClass{class: objc.GetClass("AVCapturePhotoSettings")}
	})
	return _AVCapturePhotoSettingsClass
}

// GetAVCapturePhotoSettingsClass returns the class object for AVCapturePhotoSettings.
func GetAVCapturePhotoSettingsClass() AVCapturePhotoSettingsClass {
	return getAVCapturePhotoSettingsClass()
}

type AVCapturePhotoSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCapturePhotoSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCapturePhotoSettingsClass) Alloc() AVCapturePhotoSettings {
	rv := objc.Send[AVCapturePhotoSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A specification of the features and settings to use for a single photo
// capture request.
//
// # Overview
// 
// To take a photo, you create and configure a [AVCapturePhotoSettings]
// object, then pass it to the [AVCapturePhotoOutput]
// [CapturePhotoWithSettingsDelegate] method.
// 
// A [AVCapturePhotoSettings] instance can include any combination of
// settings, regardless of whether that combination is valid for a given
// capture session. When you initiate a capture by passing a photo settings
// object to the [CapturePhotoWithSettingsDelegate] method, the photo capture
// output validates your settings to ensure deterministic behavior. For
// example, the [AVCapturePhotoSettings.FlashMode] setting must specify a value that’s present in
// the photo output’s [SupportedFlashModes] array. For detailed validation
// rules, see each property description below.
//
// # Inspecting settings
//
//   - [AVCapturePhotoSettings.UniqueID]: A unique identifier for this photo settings instance.
//   - [AVCapturePhotoSettings.Format]: A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
//   - [AVCapturePhotoSettings.ProcessedFileType]: The container file format for eventual output of the processed image.
//
// # Configuring photo settings
//
//   - [AVCapturePhotoSettings.FlashMode]: A setting for whether to fire the flash when capturing photos.
//   - [AVCapturePhotoSettings.SetFlashMode]
//   - [AVCapturePhotoSettings.MaxPhotoDimensions]: The maximum resolution of the photo to capture.
//   - [AVCapturePhotoSettings.SetMaxPhotoDimensions]
//   - [AVCapturePhotoSettings.PhotoQualityPrioritization]: A setting that indicates how to prioritize photo quality against speed of photo delivery.
//   - [AVCapturePhotoSettings.SetPhotoQualityPrioritization]
//
// # Suppressing the shutter sound
//
//   - [AVCapturePhotoSettings.ShutterSoundSuppressionEnabled]: A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//   - [AVCapturePhotoSettings.SetShutterSoundSuppressionEnabled]
//
// # Configuring constant color
//
//   - [AVCapturePhotoSettings.ConstantColorEnabled]: A Boolean value that indicates whether to capture the photo with constant color.
//   - [AVCapturePhotoSettings.SetConstantColorEnabled]
//   - [AVCapturePhotoSettings.ConstantColorFallbackPhotoDeliveryEnabled]: A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//   - [AVCapturePhotoSettings.SetConstantColorFallbackPhotoDeliveryEnabled]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings
type AVCapturePhotoSettings struct {
	objectivec.Object
}

// AVCapturePhotoSettingsFromID constructs a [AVCapturePhotoSettings] from an objc.ID.
//
// A specification of the features and settings to use for a single photo
// capture request.
func AVCapturePhotoSettingsFromID(id objc.ID) AVCapturePhotoSettings {
	return AVCapturePhotoSettings{objectivec.Object{ID: id}}
}
// NOTE: AVCapturePhotoSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCapturePhotoSettings] class.
//
// # Inspecting settings
//
//   - [IAVCapturePhotoSettings.UniqueID]: A unique identifier for this photo settings instance.
//   - [IAVCapturePhotoSettings.Format]: A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
//   - [IAVCapturePhotoSettings.ProcessedFileType]: The container file format for eventual output of the processed image.
//
// # Configuring photo settings
//
//   - [IAVCapturePhotoSettings.FlashMode]: A setting for whether to fire the flash when capturing photos.
//   - [IAVCapturePhotoSettings.SetFlashMode]
//   - [IAVCapturePhotoSettings.MaxPhotoDimensions]: The maximum resolution of the photo to capture.
//   - [IAVCapturePhotoSettings.SetMaxPhotoDimensions]
//   - [IAVCapturePhotoSettings.PhotoQualityPrioritization]: A setting that indicates how to prioritize photo quality against speed of photo delivery.
//   - [IAVCapturePhotoSettings.SetPhotoQualityPrioritization]
//
// # Suppressing the shutter sound
//
//   - [IAVCapturePhotoSettings.ShutterSoundSuppressionEnabled]: A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
//   - [IAVCapturePhotoSettings.SetShutterSoundSuppressionEnabled]
//
// # Configuring constant color
//
//   - [IAVCapturePhotoSettings.ConstantColorEnabled]: A Boolean value that indicates whether to capture the photo with constant color.
//   - [IAVCapturePhotoSettings.SetConstantColorEnabled]
//   - [IAVCapturePhotoSettings.ConstantColorFallbackPhotoDeliveryEnabled]: A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
//   - [IAVCapturePhotoSettings.SetConstantColorFallbackPhotoDeliveryEnabled]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings
type IAVCapturePhotoSettings interface {
	objectivec.IObject

	// Topic: Inspecting settings

	// A unique identifier for this photo settings instance.
	UniqueID() int64
	// A dictionary describing the processed format (for example, JPEG) to deliver captured photos in.
	Format() foundation.INSDictionary
	// The container file format for eventual output of the processed image.
	ProcessedFileType() AVFileType

	// Topic: Configuring photo settings

	// A setting for whether to fire the flash when capturing photos.
	FlashMode() AVCaptureFlashMode
	SetFlashMode(value AVCaptureFlashMode)
	// The maximum resolution of the photo to capture.
	MaxPhotoDimensions() objectivec.IObject
	SetMaxPhotoDimensions(value objectivec.IObject)
	// A setting that indicates how to prioritize photo quality against speed of photo delivery.
	PhotoQualityPrioritization() AVCapturePhotoQualityPrioritization
	SetPhotoQualityPrioritization(value AVCapturePhotoQualityPrioritization)

	// Topic: Suppressing the shutter sound

	// A Boolean value that indicates whether to suppress the built-in shutter sound when capturing a photo.
	ShutterSoundSuppressionEnabled() bool
	SetShutterSoundSuppressionEnabled(value bool)

	// Topic: Configuring constant color

	// A Boolean value that indicates whether to capture the photo with constant color.
	ConstantColorEnabled() bool
	SetConstantColorEnabled(value bool)
	// A Boolean value that indicates whether to deliver a fallback photo when taking a constant color capture.
	ConstantColorFallbackPhotoDeliveryEnabled() bool
	SetConstantColorFallbackPhotoDeliveryEnabled(value bool)

	// Name of an exception that occurs when you pass an invalid argument to a method, such as a `nil` pointer where a non-`nil` object is required.
	InvalidArgumentException() foundation.NSString
}

// Init initializes the instance.
func (c AVCapturePhotoSettings) Init() AVCapturePhotoSettings {
	rv := objc.Send[AVCapturePhotoSettings](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCapturePhotoSettings) Autorelease() AVCapturePhotoSettings {
	rv := objc.Send[AVCapturePhotoSettings](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCapturePhotoSettings creates a new AVCapturePhotoSettings instance.
func NewAVCapturePhotoSettings() AVCapturePhotoSettings {
	class := getAVCapturePhotoSettingsClass()
	rv := objc.Send[AVCapturePhotoSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a unique photo settings object, copying all settings values from
// the specified photo settings object.
//
// photoSettings: The photo settings object from which to copy settings.
//
// # Return Value
// 
// A new photo settings object.
//
// # Discussion
// 
// It is illegal to reuse a [AVCapturePhotoSettings] instance for multiple
// captures. Calling the [CapturePhotoWithSettingsDelegate] method throws an
// exception if the [UniqueID] value of the `settings` parameter matches that
// of any previously used settings object.
// 
// To reuse a specific combination of settings, use this initializer to create
// a new [AVCapturePhotoSettings] instance from an existing photo settings
// object. The newly created instance has a new, unique value for its
// [UniqueID] property, but copies the values for all other properties from
// the `photoSettings` parameter.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(from:)
func NewCapturePhotoSettingsFromPhotoSettings(photoSettings IAVCapturePhotoSettings) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(getAVCapturePhotoSettingsClass().class), objc.Sel("photoSettingsFromPhotoSettings:"), photoSettings)
	return AVCapturePhotoSettingsFromID(rv)
}

// Creates a photo settings object with the specified output format.
//
// format: A dictionary of Core Video pixel buffer attributes or AVFoundation video
// settings constants (see Video Settings).
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
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoQualityKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoQualityKey
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
//
// # Return Value
// 
// A new photo settings object.
//
// # Discussion
// 
// Requesting capture in a processed format adds requirements for other photo
// settings: for details, see the [Format] property. The capture output
// validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate
// don’t meet these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/init(format:)
func NewCapturePhotoSettingsWithFormat(format foundation.INSDictionary) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(getAVCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithFormat:"), format)
	return AVCapturePhotoSettingsFromID(rv)
}

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
func NewCapturePhotoSettingsWithRawPixelFormatType(rawPixelFormatType uint32) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(getAVCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:"), rawPixelFormatType)
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
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoQualityKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoQualityKey
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
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
func NewCapturePhotoSettingsWithRawPixelFormatTypeProcessedFormat(rawPixelFormatType uint32, processedFormat foundation.INSDictionary) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(getAVCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:processedFormat:"), rawPixelFormatType, processedFormat)
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
// //
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [AVVideoQualityKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoQualityKey
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
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
func NewCapturePhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType(rawPixelFormatType uint32, rawFileType AVFileType, processedFormat foundation.INSDictionary, processedFileType AVFileType) AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(getAVCapturePhotoSettingsClass().class), objc.Sel("photoSettingsWithRawPixelFormatType:rawFileType:processedFormat:processedFileType:"), rawPixelFormatType, objc.String(string(rawFileType)), processedFormat, objc.String(string(processedFileType)))
	return AVCapturePhotoSettingsFromID(rv)
}

// Creates a photo settings object with default settings.
//
// # Return Value
// 
// A new photo settings object.
//
// # Discussion
// 
// Capturing a photo with default settings delivers a single image in JPEG
// format.
// 
// Requesting capture in a processed format (such as JPEG) adds requirements
// for other photo settings: for details, see the [Format] property. The
// capture output validates these requirement when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate
// don’t meet these requirement, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoSettings
func (_AVCapturePhotoSettingsClass AVCapturePhotoSettingsClass) PhotoSettings() AVCapturePhotoSettings {
	rv := objc.Send[objc.ID](objc.ID(_AVCapturePhotoSettingsClass.class), objc.Sel("photoSettings"))
	return AVCapturePhotoSettingsFromID(rv)
}

// A unique identifier for this photo settings instance.
//
// # Discussion
// 
// Creating a [AVCapturePhotoSettings] instance automatically assigns a unique
// value to this property.
// 
// Use this property to track a photo capture request. After you call the
// [CapturePhotoWithSettingsDelegate] method, the photo capture output calls
// your delegate object to provide information about the progress and results
// of the capture. Each delegate method includes a
// [AVCaptureResolvedPhotoSettings] whose [UniqueID] property matches the
// [UniqueID] value of the [AVCapturePhotoSettings] object you used to request
// capture.
// 
// It is illegal to reuse a [AVCapturePhotoSettings] instance for multiple
// captures. Calling the [CapturePhotoWithSettingsDelegate] method throws an
// exception ([InvalidArgumentException]) if the `settings` object’s
// [UniqueID] value matches that of any previously used settings object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/uniqueID
func (c AVCapturePhotoSettings) UniqueID() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("uniqueID"))
	return rv
}
// A dictionary describing the processed format (for example, JPEG) to deliver
// captured photos in.
//
// # Discussion
// 
// This property is read-only—you specify a processed format when creating a
// settings object with the [PhotoSettings], [PhotoSettingsWithFormat], or
// [PhotoSettingsWithRawPixelFormatTypeProcessedFormat] initializer.
// 
// When capturing images in processed formats, the following requirements
// apply:
// 
// - This dictionary must contain a value for either the
// [kCVPixelBufferPixelFormatTypeKey] (to request an uncompressed format) or
// [AVVideoCodecKey] (to request a compressed format such as JPEG) key, but
// not both. - If this dictionary has the [kCVPixelBufferPixelFormatTypeKey]
// key, the value for that key must be listed in the photo output’s
// [AvailablePhotoPixelFormatTypes] array.
// 
// If this dictionary has the [AVVideoCodecKey] key, the value for that key
// must be listed in the photo output’s [AvailablePhotoCodecTypes] array.
// 
// - Your delegate method must implement the
// [CaptureOutputDidFinishProcessingPhotoSampleBufferPreviewPhotoSampleBufferResolvedSettingsBracketSettingsError]
// method.
// 
// The capture output validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings and delegate do
// not meet these requirements, that method raises an exception.
//
// [AVVideoCodecKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
// [kCVPixelBufferPixelFormatTypeKey]: https://developer.apple.com/documentation/CoreVideo/kCVPixelBufferPixelFormatTypeKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/format
func (c AVCapturePhotoSettings) Format() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("format"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The container file format for eventual output of the processed image.
//
// # Discussion
// 
// You specify a file format when creating capture settings with the
// [PhotoSettingsWithRawPixelFormatTypeRawFileTypeProcessedFormatProcessedFileType]
// initializer. If you didn’t specify a file format, this value is `nil`,
// and the photo output automatically chooses a default file format
// appropriate to the [Format] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/processedFileType
func (c AVCapturePhotoSettings) ProcessedFileType() AVFileType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("processedFileType"))
	return AVFileType(foundation.NSStringFromID(rv).String())
}
// A setting for whether to fire the flash when capturing photos.
//
// # Discussion
// 
// The default value for this setting is [CaptureFlashModeOff].
// 
// Assuming a static scene, using the [CaptureFlashModeAuto] setting is
// equivalent to testing the [AVCapturePhotoOutput] [IsFlashScene] property
// (which indicates whether flash is recommended for the scene currently
// visible to the camera), and then setting the [FlashMode] property of your
// photo settings output accordingly before requesting a capture. However, the
// visible scene can change between when you request a capture and when the
// camera hardware captures an image—the automatic setting ensures that the
// flash is enabled or disabled appropriately at the moment of capture. When
// the capture occurs, your [AVCapturePhotoCaptureDelegate] methods receive an
// [AVCaptureResolvedPhotoSettings] object whose [FlashEnabled] property
// indicates which flash mode was used for that capture.
// 
// When specifying a flash mode, the following requirements apply:
// 
// - The specified mode must be present in the photo output’s
// [SupportedFlashModes] array. - You may not enable image stabilization if
// the flash mode is [CaptureFlashModeOn]. (Enabling the flash takes priority
// over the [AutoStillImageStabilizationEnabled] setting).
// 
// The capture output validates these requirements when you call the
// [CapturePhotoWithSettingsDelegate] method. If your settings don’t meet
// these requirements, that method raises an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/flashMode
func (c AVCapturePhotoSettings) FlashMode() AVCaptureFlashMode {
	rv := objc.Send[AVCaptureFlashMode](c.ID, objc.Sel("flashMode"))
	return AVCaptureFlashMode(rv)
}
func (c AVCapturePhotoSettings) SetFlashMode(value AVCaptureFlashMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setFlashMode:"), value)
}
// The maximum resolution of the photo to capture.
//
// # Discussion
// 
// Setting a value specifies a maximum for the captured image. The dimensions
// must match one of the dimensions returned by capture device’s active
// format’s [supportedMaxPhotoDimensions], and be equal to or smaller than
// the value of the photo outputs [MaxPhotoDimensions].
// 
// This property defaults to the smallest dimensions returned by
// [supportedMaxPhotoDimensions].
//
// [supportedMaxPhotoDimensions]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/supportedMaxPhotoDimensions
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/maxPhotoDimensions
func (c AVCapturePhotoSettings) MaxPhotoDimensions() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("maxPhotoDimensions"))
	return objectivec.Object{ID: rv}
}
func (c AVCapturePhotoSettings) SetMaxPhotoDimensions(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxPhotoDimensions:"), value)
}
// A setting that indicates how to prioritize photo quality against speed of
// photo delivery.
//
// # Discussion
// 
// [AVCapturePhotoOutput] applies a variety of techniques to improve photo
// quality, depending on the source device’s [ActiveFormat]. Some of these
// techniques — which include reducing noise, preserving detail in low
// light, and freezing motion — can take significant processing time before
// the system returns a photo to your delegate callback. This property allows
// you to specify your preferred quality versus speed of delivery.
// 
// The default value of this property is
// [CapturePhotoQualityPrioritizationBalanced] and indicates that speed and
// quality are of equal importance to you.
// 
// When you need to prioritize speed at the expense of quality, use
// [CapturePhotoQualityPrioritizationSpeed]. Use
// [CapturePhotoQualityPrioritizationQuality] to prioritize the best quality
// at the expense of speed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/photoQualityPrioritization
func (c AVCapturePhotoSettings) PhotoQualityPrioritization() AVCapturePhotoQualityPrioritization {
	rv := objc.Send[AVCapturePhotoQualityPrioritization](c.ID, objc.Sel("photoQualityPrioritization"))
	return AVCapturePhotoQualityPrioritization(rv)
}
func (c AVCapturePhotoSettings) SetPhotoQualityPrioritization(value AVCapturePhotoQualityPrioritization) {
	objc.Send[struct{}](c.ID, objc.Sel("setPhotoQualityPrioritization:"), value)
}
// A Boolean value that indicates whether to suppress the built-in shutter
// sound when capturing a photo.
//
// # Discussion
// 
// The default value is [false]. Set the value to [true] to suppress the photo
// output’s built-in shutter sound for this request. The photo output throws
// an invalid argument exception when calling
// [CapturePhotoWithSettingsDelegate] if its
// [ShutterSoundSuppressionSupported] property returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isShutterSoundSuppressionEnabled
func (c AVCapturePhotoSettings) ShutterSoundSuppressionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isShutterSoundSuppressionEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetShutterSoundSuppressionEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShutterSoundSuppressionEnabled:"), value)
}
// A Boolean value that indicates whether to capture the photo with constant
// color.
//
// # Discussion
// 
// The default value is [false]. Set the value to [true] to capture a constant
// color photo.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorEnabled
func (c AVCapturePhotoSettings) ConstantColorEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isConstantColorEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetConstantColorEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setConstantColorEnabled:"), value)
}
// A Boolean value that indicates whether to deliver a fallback photo when
// taking a constant color capture.
//
// # Discussion
// 
// The default value is [false]. Set the value to [true] to receive a fallback
// photo that you can use if the main constant color photo’s confidence
// level doesn’t meet your requirement.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoSettings/isConstantColorFallbackPhotoDeliveryEnabled
func (c AVCapturePhotoSettings) ConstantColorFallbackPhotoDeliveryEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isConstantColorFallbackPhotoDeliveryEnabled"))
	return rv
}
func (c AVCapturePhotoSettings) SetConstantColorFallbackPhotoDeliveryEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setConstantColorFallbackPhotoDeliveryEnabled:"), value)
}
// Name of an exception that occurs when you pass an invalid argument to a
// method, such as a `nil` pointer where a non-`nil` object is required.
//
// See: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (c AVCapturePhotoSettings) InvalidArgumentException() foundation.NSString {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("invalidArgumentException"))
	return foundation.NSStringFromID(objc.ID(rv))
}

