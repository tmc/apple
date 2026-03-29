// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCapturePhoto] class.
var (
	_AVCapturePhotoClass     AVCapturePhotoClass
	_AVCapturePhotoClassOnce sync.Once
)

func getAVCapturePhotoClass() AVCapturePhotoClass {
	_AVCapturePhotoClassOnce.Do(func() {
		_AVCapturePhotoClass = AVCapturePhotoClass{class: objc.GetClass("AVCapturePhoto")}
	})
	return _AVCapturePhotoClass
}

// GetAVCapturePhotoClass returns the class object for AVCapturePhoto.
func GetAVCapturePhotoClass() AVCapturePhotoClass {
	return getAVCapturePhotoClass()
}

type AVCapturePhotoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCapturePhotoClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCapturePhotoClass) Alloc() AVCapturePhoto {
	rv := objc.Send[AVCapturePhoto](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A container for image data from a photo capture output.
//
// # Overview
// 
// When you capture photos with the [AVCapturePhotoOutput] class, your
// delegate object receives each resulting image and related data in the form
// of an [AVCapturePhoto] object. This object is an immutable wrapper from
// which you can retrieve various results of the photo capture.
// 
// In addition to the photo image pixel buffer, an AVCapturePhoto object can
// also contain a preview-sized pixel buffer, capture metadata, and, on
// supported devices, depth data and camera calibration data. From an
// [AVCapturePhoto] object, you can generate data appropriate for writing to a
// file, such as HEVC encoded image data containerized in the HEIC file format
// and including a preview image, depth data and other attachments.
// 
// An [AVCapturePhoto] instance wraps a single image result. For example, if
// you request a bracketed capture of three images, your callback is called
// three times, each time delivering a single [AVCapturePhoto] object.
//
// # Resolving photo capture requests
//
//   - [AVCapturePhoto.ResolvedSettings]: The settings object that was used to request this photo capture.
//   - [AVCapturePhoto.PhotoCount]: The 1-based index of this photo capture relative to other results from the same capture request.
//   - [AVCapturePhoto.Timestamp]: The time at which the image was captured.
//
// # Accessing photo pixel data
//
//   - [AVCapturePhoto.PixelBuffer]: The uncompressed or RAW image sample buffer for the photo, if requested.
//
// # Packaging data for file output
//
//   - [AVCapturePhoto.FileDataRepresentation]: Generates and returns a flat data representation of the photo and its attachments.
//   - [AVCapturePhoto.CGImageRepresentation]: Extracts and returns the captured photo’s primary image as a Core Graphics image object.
//
// # Enabling constant color
//
//   - [AVCapturePhoto.ConstantColorCenterWeightedMeanConfidenceLevel]: A score that summarizes the overall confidence level of a constant color photo.
//   - [AVCapturePhoto.ConstantColorConfidenceMap]: A pixel buffer where each pixel value indicates how fully the system achieves the constant color effect in the corresponding region of the photo.
//   - [AVCapturePhoto.ConstantColorFallbackPhoto]: A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto
type AVCapturePhoto struct {
	objectivec.Object
}

// AVCapturePhotoFromID constructs a [AVCapturePhoto] from an objc.ID.
//
// A container for image data from a photo capture output.
func AVCapturePhotoFromID(id objc.ID) AVCapturePhoto {
	return AVCapturePhoto{objectivec.Object{ID: id}}
}
// NOTE: AVCapturePhoto adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCapturePhoto] class.
//
// # Resolving photo capture requests
//
//   - [IAVCapturePhoto.ResolvedSettings]: The settings object that was used to request this photo capture.
//   - [IAVCapturePhoto.PhotoCount]: The 1-based index of this photo capture relative to other results from the same capture request.
//   - [IAVCapturePhoto.Timestamp]: The time at which the image was captured.
//
// # Accessing photo pixel data
//
//   - [IAVCapturePhoto.PixelBuffer]: The uncompressed or RAW image sample buffer for the photo, if requested.
//
// # Packaging data for file output
//
//   - [IAVCapturePhoto.FileDataRepresentation]: Generates and returns a flat data representation of the photo and its attachments.
//   - [IAVCapturePhoto.CGImageRepresentation]: Extracts and returns the captured photo’s primary image as a Core Graphics image object.
//
// # Enabling constant color
//
//   - [IAVCapturePhoto.ConstantColorCenterWeightedMeanConfidenceLevel]: A score that summarizes the overall confidence level of a constant color photo.
//   - [IAVCapturePhoto.ConstantColorConfidenceMap]: A pixel buffer where each pixel value indicates how fully the system achieves the constant color effect in the corresponding region of the photo.
//   - [IAVCapturePhoto.ConstantColorFallbackPhoto]: A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto
type IAVCapturePhoto interface {
	objectivec.IObject

	// Topic: Resolving photo capture requests

	// The settings object that was used to request this photo capture.
	ResolvedSettings() IAVCaptureResolvedPhotoSettings
	// The 1-based index of this photo capture relative to other results from the same capture request.
	PhotoCount() int
	// The time at which the image was captured.
	Timestamp() coremedia.CMTime

	// Topic: Accessing photo pixel data

	// The uncompressed or RAW image sample buffer for the photo, if requested.
	PixelBuffer() corevideo.CVImageBufferRef

	// Topic: Packaging data for file output

	// Generates and returns a flat data representation of the photo and its attachments.
	FileDataRepresentation() foundation.INSData
	// Extracts and returns the captured photo’s primary image as a Core Graphics image object.
	CGImageRepresentation() coregraphics.CGImageRef

	// Topic: Enabling constant color

	// A score that summarizes the overall confidence level of a constant color photo.
	ConstantColorCenterWeightedMeanConfidenceLevel() float32
	// A pixel buffer where each pixel value indicates how fully the system achieves the constant color effect in the corresponding region of the photo.
	ConstantColorConfidenceMap() corevideo.CVImageBufferRef
	// A Boolean value that Indicates whether this photo is a fallback photo for a constant color capture.
	ConstantColorFallbackPhoto() bool
}

// Init initializes the instance.
func (c AVCapturePhoto) Init() AVCapturePhoto {
	rv := objc.Send[AVCapturePhoto](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCapturePhoto) Autorelease() AVCapturePhoto {
	rv := objc.Send[AVCapturePhoto](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCapturePhoto creates a new AVCapturePhoto instance.
func NewAVCapturePhoto() AVCapturePhoto {
	class := getAVCapturePhotoClass()
	rv := objc.Send[AVCapturePhoto](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Generates and returns a flat data representation of the photo and its
// attachments.
//
// # Return Value
// 
// Data appropriate for writing to a file of the type specified when
// requesting photo capture, or `nil` if the photo and attachment data cannot
// be flattened.
//
// # Discussion
// 
// When you request a photo capture with the [AVCapturePhotoOutput]
// [CapturePhotoWithSettingsDelegate] method, the [AVCapturePhotoSettings]
// object you provide specifies image data formats (such as JPEG and HEVC) and
// container file formats (such as JFIF and HEIF) for the resulting image
// file. Calling this method formats and packages the image pixel buffer,
// along with metadata and other attachments created during capture (such as
// preview photos and depth maps), into data appropriate for writing to a file
// of that type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/fileDataRepresentation()
func (c AVCapturePhoto) FileDataRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("fileDataRepresentation"))
	return foundation.NSDataFromID(rv)
}
// Extracts and returns the captured photo’s primary image as a Core
// Graphics image object.
//
// # Return Value
// 
// A Core Graphics image representation of the captured photo, or `nil` if the
// image cannot be converted.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/cgImageRepresentation()
func (c AVCapturePhoto) CGImageRepresentation() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](c.ID, objc.Sel("CGImageRepresentation"))
	return coregraphics.CGImageRef(rv)
}

// The settings object that was used to request this photo capture.
//
// # Discussion
// 
// To determine which [CapturePhotoWithSettingsDelegate] call produced this
// photo capture result, match this [AVCaptureResolvedPhotoSettings]
// object’s [UniqueID] value to the [UniqueID] property of the photo
// settings object you requested capture with. You can also use this object to
// find out which values the photo output has chosen for automatic settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/resolvedSettings
func (c AVCapturePhoto) ResolvedSettings() IAVCaptureResolvedPhotoSettings {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("resolvedSettings"))
	return AVCaptureResolvedPhotoSettingsFromID(objc.ID(rv))
}
// The 1-based index of this photo capture relative to other results from the
// same capture request.
//
// # Discussion
// 
// The [ExpectedPhotoCount] property of this capture result’s
// [ResolvedSettings] object indicates the total number of images that will be
// returned for a given capture request. When your delegate’s
// [CaptureOutputDidFinishProcessingPhotoError] method receives a photo whose
// [PhotoCount] value matches the [ExpectedPhotoCount] value, you know
// you’ve received the last one for the given capture request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/photoCount
func (c AVCapturePhoto) PhotoCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("photoCount"))
	return rv
}
// The time at which the image was captured.
//
// # Discussion
// 
// This timestamp is always synchronized to the [masterClock] time of the
// [AVCaptureSession] object to which the photo output is connected.
//
// [masterClock]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/masterClock
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/timestamp
func (c AVCapturePhoto) Timestamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("timestamp"))
	return coremedia.CMTime(rv)
}
// The uncompressed or RAW image sample buffer for the photo, if requested.
//
// # Discussion
// 
// If you requested photo capture in a RAW format, or in a processed format
// without compression such as TIFF, you can use this property to access the
// underlying sample buffer.
// 
// If you requested capture in a compressed format such as JPEG or HEVC/HEIF,
// this property’s value is `nil`. Use the [FileDataRepresentation] or
// [CGImageRepresentation] method to obtain compressed image data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/pixelBuffer
func (c AVCapturePhoto) PixelBuffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](c.ID, objc.Sel("pixelBuffer"))
	return corevideo.CVImageBufferRef(rv)
}
// A score that summarizes the overall confidence level of a constant color
// photo.
//
// # Discussion
// 
// A value of `1.0` means full confidence and `0.0` means zero confidence. The
// default is `0.0.`
// 
// In most use cases, such as document scanning, the system considers the
// central region of the photo more important than its edges. The system
// weights the confidence level of the central pixels more heavily than pixels
// on the edges of the photo.
// 
// Use [ConstantColorConfidenceMap] for more use case specific analyses of the
// confidence level.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/constantColorCenterWeightedMeanConfidenceLevel
func (c AVCapturePhoto) ConstantColorCenterWeightedMeanConfidenceLevel() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("constantColorCenterWeightedMeanConfidenceLevel"))
	return rv
}
// A pixel buffer where each pixel value indicates how fully the system
// achieves the constant color effect in the corresponding region of the
// photo.
//
// # Discussion
// 
// A value of `255` means full confidence and `0` means zero confidence.
// 
// This property provides a valid value only for constant color photos. The
// value is `nil` in all other cases.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/constantColorConfidenceMap
func (c AVCapturePhoto) ConstantColorConfidenceMap() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](c.ID, objc.Sel("constantColorConfidenceMap"))
	return corevideo.CVImageBufferRef(rv)
}
// A Boolean value that Indicates whether this photo is a fallback photo for a
// constant color capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhoto/isConstantColorFallbackPhoto
func (c AVCapturePhoto) ConstantColorFallbackPhoto() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isConstantColorFallbackPhoto"))
	return rv
}

