// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPortraitEffectsMatte] class.
var (
	_AVPortraitEffectsMatteClass     AVPortraitEffectsMatteClass
	_AVPortraitEffectsMatteClassOnce sync.Once
)

func getAVPortraitEffectsMatteClass() AVPortraitEffectsMatteClass {
	_AVPortraitEffectsMatteClassOnce.Do(func() {
		_AVPortraitEffectsMatteClass = AVPortraitEffectsMatteClass{class: objc.GetClass("AVPortraitEffectsMatte")}
	})
	return _AVPortraitEffectsMatteClass
}

// GetAVPortraitEffectsMatteClass returns the class object for AVPortraitEffectsMatte.
func GetAVPortraitEffectsMatteClass() AVPortraitEffectsMatteClass {
	return getAVPortraitEffectsMatteClass()
}

type AVPortraitEffectsMatteClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPortraitEffectsMatteClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPortraitEffectsMatteClass) Alloc() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An auxiliary image used to separate foreground from background with high
// resolution.
//
// # Overview
//
// Before iOS 11, the iPhone camera software used depth maps to render a
// shallow depth of field (the effect) into still images taken in Portrait
// Mode before discarding the maps. Because the effect was part of the photo,
// you couldn’t access the maps separately, as metadata, for photos taken by
// devices running iOS 10 or earlier.
//
// Starting in iOS 11, apps accessing the photo library can use images
// containing embedded auxiliary depth maps to render creative depth effects,
// such as forced perspective, or image projection from 2D to 3D space. These
// depth maps are low-resolution compared to the full-resolution RGB image. As
// such, the depth effects you can render are limited by the resolution and
// accuracy of the maps. Fine detail, such as hair, is challenging to preserve
// faithfully at the resolution of these depth maps.
//
// Starting in iOS 12, the portrait effects matte helps achieve this
// fine-grain level of detail.
//
// [media-3030223]
//
// [Table data omitted]
//
// Using the auxiliary matte image, you can improve the quality of rendered
// portrait effects, such as Natural Light, Studio Light, Contour Light, Stage
// Light, and Stage Light Mono.
//
// Unlike the depth map, the portrait effects matte isn’t intended to
// faithfully preserve all gradations of depth in the scene. It’s a
// depth-guided, people-focused segmentation mask generated from a proprietary
// Apple neural network trained to detect people. It separates an individual
// in the foreground from whatever is in the background, with greater detail
// and clarity than with the depth map alone. It achieves this clarity in part
// because the matte image has higher resolution than the depth map.
//
// For information about capturing the portrait effects matte, see
// [Configuring camera capture to collect a Portrait Effects matte]. To learn
// how to extract a portrait effects matte from photos previously captured in
// portrait mode on a device running iOS 12, see [Extracting Portrait Effects
// matte image data from a photo].
//
// # Creating a Portrait Effects matte
//
//   - [AVPortraitEffectsMatte.PortraitEffectsMatteByApplyingExifOrientation]: Returns a derivative portrait effects matte after applying the specified EXIF orientation.
//   - [AVPortraitEffectsMatte.PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError]: Returns a portrait effects matte by wrapping the replacement pixel buffer.
//
// # Examining a Portrait Effects matte
//
//   - [AVPortraitEffectsMatte.MattingImage]: The portrait effects matte’s internal image, formatted as a pixel buffer.
//   - [AVPortraitEffectsMatte.PixelFormatType]: The pixel format type of this portrait effects matte’s internal image.
//   - [AVPortraitEffectsMatte.DictionaryRepresentationForAuxiliaryDataType]: A dictionary of primitive map information used for writing an image file with a portrait effects matte.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte
//
// [Configuring camera capture to collect a Portrait Effects matte]: https://developer.apple.com/documentation/AVFoundation/configuring-camera-capture-to-collect-a-portrait-effects-matte
// [Extracting Portrait Effects matte image data from a photo]: https://developer.apple.com/documentation/AVFoundation/extracting-portrait-effects-matte-image-data-from-a-photo
type AVPortraitEffectsMatte struct {
	objectivec.Object
}

// AVPortraitEffectsMatteFromID constructs a [AVPortraitEffectsMatte] from an objc.ID.
//
// An auxiliary image used to separate foreground from background with high
// resolution.
func AVPortraitEffectsMatteFromID(id objc.ID) AVPortraitEffectsMatte {
	return AVPortraitEffectsMatte{objectivec.Object{ID: id}}
}

// NOTE: AVPortraitEffectsMatte adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPortraitEffectsMatte] class.
//
// # Creating a Portrait Effects matte
//
//   - [IAVPortraitEffectsMatte.PortraitEffectsMatteByApplyingExifOrientation]: Returns a derivative portrait effects matte after applying the specified EXIF orientation.
//   - [IAVPortraitEffectsMatte.PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError]: Returns a portrait effects matte by wrapping the replacement pixel buffer.
//
// # Examining a Portrait Effects matte
//
//   - [IAVPortraitEffectsMatte.MattingImage]: The portrait effects matte’s internal image, formatted as a pixel buffer.
//   - [IAVPortraitEffectsMatte.PixelFormatType]: The pixel format type of this portrait effects matte’s internal image.
//   - [IAVPortraitEffectsMatte.DictionaryRepresentationForAuxiliaryDataType]: A dictionary of primitive map information used for writing an image file with a portrait effects matte.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte
type IAVPortraitEffectsMatte interface {
	objectivec.IObject

	// Topic: Creating a Portrait Effects matte

	// Returns a derivative portrait effects matte after applying the specified EXIF orientation.
	PortraitEffectsMatteByApplyingExifOrientation(exifOrientation objectivec.IObject) IAVPortraitEffectsMatte
	// Returns a portrait effects matte by wrapping the replacement pixel buffer.
	PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer corevideo.CVImageBufferRef) (IAVPortraitEffectsMatte, error)

	// Topic: Examining a Portrait Effects matte

	// The portrait effects matte’s internal image, formatted as a pixel buffer.
	MattingImage() corevideo.CVImageBufferRef
	// The pixel format type of this portrait effects matte’s internal image.
	PixelFormatType() uint32
	// A dictionary of primitive map information used for writing an image file with a portrait effects matte.
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.INSDictionary
}

// Init initializes the instance.
func (p AVPortraitEffectsMatte) Init() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPortraitEffectsMatte) Autorelease() AVPortraitEffectsMatte {
	rv := objc.Send[AVPortraitEffectsMatte](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPortraitEffectsMatte creates a new AVPortraitEffectsMatte instance.
func NewAVPortraitEffectsMatte() AVPortraitEffectsMatte {
	class := getAVPortraitEffectsMatteClass()
	rv := objc.Send[AVPortraitEffectsMatte](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a portrait effects matte instance from auxiliary image
// information in an image file.
//
// imageSourceAuxDataInfoDictionary: A dictionary of information related to primitive portrait effects matte;
// obtained from [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)].
//
// # Discussion
//
// When using the [Image I/O] API to read from a HEIF or JPEG file containing
// a portrait effects matte, you can create an [AVPortraitEffectsMatte] object
// from the result of [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)].
// This function returns a [CFDictionary] of primitive map information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/init(fromDictionaryRepresentation:)
//
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
// [CFDictionary]: https://developer.apple.com/documentation/CoreFoundation/CFDictionary
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
//
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
func NewPortraitEffectsMatteFromDictionaryRepresentationError(imageSourceAuxDataInfoDictionary foundation.INSDictionary) (AVPortraitEffectsMatte, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getAVPortraitEffectsMatteClass().class), objc.Sel("portraitEffectsMatteFromDictionaryRepresentation:error:"), imageSourceAuxDataInfoDictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVPortraitEffectsMatte{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVPortraitEffectsMatteFromID(rv), nil
}

// Returns a derivative portrait effects matte after applying the specified
// EXIF orientation.
//
// exifOrientation: One of the standard EXIF orientation tags expressing how the portrait
// effects matte should be rotated or mirrored.
//
// exifOrientation is a [imageio.CGImagePropertyOrientation].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/applyingExifOrientation(_:)
// exifOrientation is a [imageio.CGImagePropertyOrientation].
func (p AVPortraitEffectsMatte) PortraitEffectsMatteByApplyingExifOrientation(exifOrientation objectivec.IObject) IAVPortraitEffectsMatte {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("portraitEffectsMatteByApplyingExifOrientation:"), exifOrientation)
	return AVPortraitEffectsMatteFromID(rv)
}

// Returns a portrait effects matte by wrapping the replacement pixel buffer.
//
// pixelBuffer: A pixel buffer containing a portrait effects matte image, represented as
// [KCVPixelFormatType_OneComponent8] with
// [kCVImageBufferColorPrimaries_ITU_R_709_2] color primaries and a
// [kCVImageBufferTransferFunction_Linear] transfer function.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/replacingPortraitEffectsMatte(with:)
//
// [kCVImageBufferColorPrimaries_ITU_R_709_2]: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferColorPrimaries_ITU_R_709_2
// [kCVImageBufferTransferFunction_Linear]: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_Linear
func (p AVPortraitEffectsMatte) PortraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBufferError(pixelBuffer corevideo.CVImageBufferRef) (IAVPortraitEffectsMatte, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("portraitEffectsMatteByReplacingPortraitEffectsMatteWithPixelBuffer:error:"), pixelBuffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVPortraitEffectsMatte{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVPortraitEffectsMatteFromID(rv), nil

}

// A dictionary of primitive map information used for writing an image file
// with a portrait effects matte.
//
// outAuxDataType: Must be [kCGImageAuxiliaryDataTypePortraitEffectsMatte].
//
// # Return Value
//
// A dictionary of primitive map information for
// [CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/dictionaryRepresentation(forAuxiliaryDataType:)
//
// [kCGImageAuxiliaryDataTypePortraitEffectsMatte]: https://developer.apple.com/documentation/ImageIO/kCGImageAuxiliaryDataTypePortraitEffectsMatte
// [CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
func (p AVPortraitEffectsMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), objc.String(outAuxDataType))
	return foundation.NSDictionaryFromID(rv)
}

// The portrait effects matte’s internal image, formatted as a pixel buffer.
//
// # Discussion
//
// Query the pixel format using the [PixelFormatType] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/mattingImage
func (p AVPortraitEffectsMatte) MattingImage() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](p.ID, objc.Sel("mattingImage"))
	return corevideo.CVImageBufferRef(rv)
}

// The pixel format type of this portrait effects matte’s internal image.
//
// # Discussion
//
// The only supported pixel format type for the matting image is
// [KCVPixelFormatType_OneComponent8].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPortraitEffectsMatte/pixelFormatType
func (p AVPortraitEffectsMatte) PixelFormatType() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("pixelFormatType"))
	return rv
}
