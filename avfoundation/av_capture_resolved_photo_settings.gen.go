// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureResolvedPhotoSettings] class.
var (
	_AVCaptureResolvedPhotoSettingsClass     AVCaptureResolvedPhotoSettingsClass
	_AVCaptureResolvedPhotoSettingsClassOnce sync.Once
)

func getAVCaptureResolvedPhotoSettingsClass() AVCaptureResolvedPhotoSettingsClass {
	_AVCaptureResolvedPhotoSettingsClassOnce.Do(func() {
		_AVCaptureResolvedPhotoSettingsClass = AVCaptureResolvedPhotoSettingsClass{class: objc.GetClass("AVCaptureResolvedPhotoSettings")}
	})
	return _AVCaptureResolvedPhotoSettingsClass
}

// GetAVCaptureResolvedPhotoSettingsClass returns the class object for AVCaptureResolvedPhotoSettings.
func GetAVCaptureResolvedPhotoSettingsClass() AVCaptureResolvedPhotoSettingsClass {
	return getAVCaptureResolvedPhotoSettingsClass()
}

type AVCaptureResolvedPhotoSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureResolvedPhotoSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureResolvedPhotoSettingsClass) Alloc() AVCaptureResolvedPhotoSettings {
	rv := objc.Send[AVCaptureResolvedPhotoSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A description of the features and settings in use for an in-progress or
// complete photo capture request.
//
// # Overview
//
// When you request a photo capture using the [AVCapturePhotoOutput]
// [CapturePhotoWithSettingsDelegate] method, you describe the settings for
// that capture request in an [AVCapturePhotoSettings] object. When the
// capture begins, the photo output calls your delegate methods and provides
// an [AVCaptureResolvedPhotoSettings] object detailing the settings that are
// in effect for that capture. Resolved photo settings objects are immutable;
// they describe a request that has already been made.
//
// The [AVCaptureResolvedPhotoSettings.UniqueID] property of a resolved photo settings object passed to one
// of your [AVCapturePhotoCaptureDelegate] methods matches the [AVCaptureResolvedPhotoSettings.UniqueID]
// value of the [AVCapturePhotoSettings] object you passed when requesting
// capture. Use this value to determine which delegate method calls correspond
// to which capture requests.
//
// Some photo capture settings are automatic, such as the [AVCaptureResolvedPhotoSettings.FlashMode]
// property. For such settings, the photo output determines whether to use
// that feature at the moment of capture—you don’t know when requesting a
// capture whether the feature is active when the capture completes. When the
// photo output calls your delegate methods, the provided
// [AVCaptureResolvedPhotoSettings] object details which automatic features
// have been set for that capture.
//
// Likewise, the dimensions of an output image or movie may not be set until
// the moment of capture. For example, when you specify a thumbnail size with
// the [AVCaptureResolvedPhotoSettings.PreviewPhotoFormat] setting, the photo output chooses dimensions that
// best match your requested size while preserving the aspect ratio of the
// captured photo. When the photo output calls your delegate methods, use the
// [AVCaptureResolvedPhotoSettings.PreviewDimensions] property of the resolved settings to find the actual
// preview image dimensions. See the methods listed in Examining Output
// Dimensions for other cases where output dimensions can change at capture
// time.
//
// # Resolving photo capture requests
//
//   - [AVCaptureResolvedPhotoSettings.UniqueID]: The unique identifier for the photo capture this settings object corresponds to.
//   - [AVCaptureResolvedPhotoSettings.ExpectedPhotoCount]: The number of photo capture results in the capture request.
//
// # Examining photo capture settings
//
//   - [AVCaptureResolvedPhotoSettings.FastCapturePrioritizationEnabled]: A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// # Examining output dimensions
//
//   - [AVCaptureResolvedPhotoSettings.PhotoDimensions]: The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings
type AVCaptureResolvedPhotoSettings struct {
	objectivec.Object
}

// AVCaptureResolvedPhotoSettingsFromID constructs a [AVCaptureResolvedPhotoSettings] from an objc.ID.
//
// A description of the features and settings in use for an in-progress or
// complete photo capture request.
func AVCaptureResolvedPhotoSettingsFromID(id objc.ID) AVCaptureResolvedPhotoSettings {
	return AVCaptureResolvedPhotoSettings{objectivec.Object{ID: id}}
}

// NOTE: AVCaptureResolvedPhotoSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureResolvedPhotoSettings] class.
//
// # Resolving photo capture requests
//
//   - [IAVCaptureResolvedPhotoSettings.UniqueID]: The unique identifier for the photo capture this settings object corresponds to.
//   - [IAVCaptureResolvedPhotoSettings.ExpectedPhotoCount]: The number of photo capture results in the capture request.
//
// # Examining photo capture settings
//
//   - [IAVCaptureResolvedPhotoSettings.FastCapturePrioritizationEnabled]: A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
//
// # Examining output dimensions
//
//   - [IAVCaptureResolvedPhotoSettings.PhotoDimensions]: The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings
type IAVCaptureResolvedPhotoSettings interface {
	objectivec.IObject

	// Topic: Resolving photo capture requests

	// The unique identifier for the photo capture this settings object corresponds to.
	UniqueID() int64
	// The number of photo capture results in the capture request.
	ExpectedPhotoCount() uint

	// Topic: Examining photo capture settings

	// A Boolean value that indicates whether the system uses fast capture prioritization when capturing the photo.
	FastCapturePrioritizationEnabled() bool

	// Topic: Examining output dimensions

	// The size, in pixels, of the photo image (in a processed format, such as JPEG) that the capture delivers.
	PhotoDimensions() coremedia.CMVideoDimensions

	// A setting for whether to fire the flash when capturing photos.
	FlashMode() AVCaptureFlashMode
	SetFlashMode(value AVCaptureFlashMode)
	// A dictionary describing the format for delivery of preview-sized images alongside the main photo.
	PreviewPhotoFormat() string
	SetPreviewPhotoFormat(value string)
}

// Init initializes the instance.
func (c AVCaptureResolvedPhotoSettings) Init() AVCaptureResolvedPhotoSettings {
	rv := objc.Send[AVCaptureResolvedPhotoSettings](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureResolvedPhotoSettings) Autorelease() AVCaptureResolvedPhotoSettings {
	rv := objc.Send[AVCaptureResolvedPhotoSettings](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureResolvedPhotoSettings creates a new AVCaptureResolvedPhotoSettings instance.
func NewAVCaptureResolvedPhotoSettings() AVCaptureResolvedPhotoSettings {
	class := getAVCaptureResolvedPhotoSettingsClass()
	rv := objc.Send[AVCaptureResolvedPhotoSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The unique identifier for the photo capture this settings object
// corresponds to.
//
// # Discussion
//
// The value of this property matches the matches the [UniqueID] value of the
// [AVCapturePhotoSettings] object you passed when initiating a photo capture
// with the [CapturePhotoWithSettingsDelegate] method. Use this value to
// determine which delegate method calls correspond to which capture requests.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/uniqueID
func (c AVCaptureResolvedPhotoSettings) UniqueID() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("uniqueID"))
	return rv
}

// The number of photo capture results in the capture request.
//
// # Discussion
//
// When you request a photo capture, the photo output calls your delegate’s
// [CaptureOutputDidFinishProcessingPhotoError] method many times based on the
// settings you choose. For example, if you request a bracket of three
// exposures with image delivery in both JPEG and RAW formats, the expected
// photo count is `6`.
//
// The [PhotoCount] property of each [AVCapturePhoto] object delivered to your
// delegate indicates where that capture result relates to this sequence. When
// your delegate receives a photo whose [PhotoCount] value matches the
// [ExpectedPhotoCount], you know you’ve received the last one for the given
// capture request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/expectedPhotoCount
func (c AVCaptureResolvedPhotoSettings) ExpectedPhotoCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("expectedPhotoCount"))
	return rv
}

// A Boolean value that indicates whether the system uses fast capture
// prioritization when capturing the photo.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/isFastCapturePrioritizationEnabled
func (c AVCaptureResolvedPhotoSettings) FastCapturePrioritizationEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFastCapturePrioritizationEnabled"))
	return rv
}

// The size, in pixels, of the photo image (in a processed format, such as
// JPEG) that the capture delivers.
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
// If you do not request capture in a processed format (that is, if you
// request capture in RAW format only), this property’s value has zero width
// and zero height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureResolvedPhotoSettings/photoDimensions
func (c AVCaptureResolvedPhotoSettings) PhotoDimensions() coremedia.CMVideoDimensions {
	rv := objc.Send[coremedia.CMVideoDimensions](c.ID, objc.Sel("photoDimensions"))
	return coremedia.CMVideoDimensions(rv)
}

// A setting for whether to fire the flash when capturing photos.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/flashmode
func (c AVCaptureResolvedPhotoSettings) FlashMode() AVCaptureFlashMode {
	rv := objc.Send[AVCaptureFlashMode](c.ID, objc.Sel("flashMode"))
	return AVCaptureFlashMode(rv)
}
func (c AVCaptureResolvedPhotoSettings) SetFlashMode(value AVCaptureFlashMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setFlashMode:"), value)
}

// A dictionary describing the format for delivery of preview-sized images
// alongside the main photo.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturephotosettings/previewphotoformat
func (c AVCaptureResolvedPhotoSettings) PreviewPhotoFormat() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previewPhotoFormat"))
	return foundation.NSStringFromID(rv).String()
}
func (c AVCaptureResolvedPhotoSettings) SetPreviewPhotoFormat(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreviewPhotoFormat:"), objc.String(value))
}
