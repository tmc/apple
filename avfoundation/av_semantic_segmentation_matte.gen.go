// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSemanticSegmentationMatte] class.
var (
	_AVSemanticSegmentationMatteClass     AVSemanticSegmentationMatteClass
	_AVSemanticSegmentationMatteClassOnce sync.Once
)

func getAVSemanticSegmentationMatteClass() AVSemanticSegmentationMatteClass {
	_AVSemanticSegmentationMatteClassOnce.Do(func() {
		_AVSemanticSegmentationMatteClass = AVSemanticSegmentationMatteClass{class: objc.GetClass("AVSemanticSegmentationMatte")}
	})
	return _AVSemanticSegmentationMatteClass
}

// GetAVSemanticSegmentationMatteClass returns the class object for AVSemanticSegmentationMatte.
func GetAVSemanticSegmentationMatteClass() AVSemanticSegmentationMatteClass {
	return getAVSemanticSegmentationMatteClass()
}

type AVSemanticSegmentationMatteClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSemanticSegmentationMatteClass) Alloc() AVSemanticSegmentationMatte {
	rv := objc.Send[AVSemanticSegmentationMatte](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that wraps a matting image for a particular semantic
// segmentation.
//
// # Overview
// 
// The matting image stores its pixel data as [CVPixelBuffer] objects in
// [KCVPixelFormatType_OneComponent8] format. The image file contains the
// semantic segmentation matte as an auxiliary image, accessible using the
// ImageIO framework’s [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]
// function.
//
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
//
// # Creating a segmentation matte
//
//   - [AVSemanticSegmentationMatte.SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError]: Returns a semantic segmentation matte instance that wraps the replacement pixel buffer.
//   - [AVSemanticSegmentationMatte.SemanticSegmentationMatteByApplyingExifOrientation]: Returns a new semantic segmentation matte instance with the specified Exif orientation applied.
//   - [AVSemanticSegmentationMatte.DictionaryRepresentationForAuxiliaryDataType]: Returns a dictionary of primitive map information to use when writing an image file with a semantic segmentation matte.
//
// # Inspecting a segmentation matte
//
//   - [AVSemanticSegmentationMatte.MatteType]: The semantic segmentation matte image type.
//   - [AVSemanticSegmentationMatte.MattingImage]: The semantic segmentation matte’s internal image.
//   - [AVSemanticSegmentationMatte.PixelFormatType]: The pixel format type for this object’s internal matting image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte
type AVSemanticSegmentationMatte struct {
	objectivec.Object
}

// AVSemanticSegmentationMatteFromID constructs a [AVSemanticSegmentationMatte] from an objc.ID.
//
// An object that wraps a matting image for a particular semantic
// segmentation.
func AVSemanticSegmentationMatteFromID(id objc.ID) AVSemanticSegmentationMatte {
	return AVSemanticSegmentationMatte{objectivec.Object{ID: id}}
}
// NOTE: AVSemanticSegmentationMatte adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSemanticSegmentationMatte] class.
//
// # Creating a segmentation matte
//
//   - [IAVSemanticSegmentationMatte.SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError]: Returns a semantic segmentation matte instance that wraps the replacement pixel buffer.
//   - [IAVSemanticSegmentationMatte.SemanticSegmentationMatteByApplyingExifOrientation]: Returns a new semantic segmentation matte instance with the specified Exif orientation applied.
//   - [IAVSemanticSegmentationMatte.DictionaryRepresentationForAuxiliaryDataType]: Returns a dictionary of primitive map information to use when writing an image file with a semantic segmentation matte.
//
// # Inspecting a segmentation matte
//
//   - [IAVSemanticSegmentationMatte.MatteType]: The semantic segmentation matte image type.
//   - [IAVSemanticSegmentationMatte.MattingImage]: The semantic segmentation matte’s internal image.
//   - [IAVSemanticSegmentationMatte.PixelFormatType]: The pixel format type for this object’s internal matting image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte
type IAVSemanticSegmentationMatte interface {
	objectivec.IObject

	// Topic: Creating a segmentation matte

	// Returns a semantic segmentation matte instance that wraps the replacement pixel buffer.
	SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer corevideo.CVImageBufferRef) (IAVSemanticSegmentationMatte, error)
	// Returns a new semantic segmentation matte instance with the specified Exif orientation applied.
	SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation objectivec.IObject) IAVSemanticSegmentationMatte
	// Returns a dictionary of primitive map information to use when writing an image file with a semantic segmentation matte.
	DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.INSDictionary

	// Topic: Inspecting a segmentation matte

	// The semantic segmentation matte image type.
	MatteType() AVSemanticSegmentationMatteType
	// The semantic segmentation matte’s internal image.
	MattingImage() corevideo.CVImageBufferRef
	// The pixel format type for this object’s internal matting image.
	PixelFormatType() uint32

	// 8-bit one component, black is zero.
	KCVPixelFormatType_OneComponent8() uint32
	SetKCVPixelFormatType_OneComponent8(value uint32)
}

// Init initializes the instance.
func (s AVSemanticSegmentationMatte) Init() AVSemanticSegmentationMatte {
	rv := objc.Send[AVSemanticSegmentationMatte](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSemanticSegmentationMatte) Autorelease() AVSemanticSegmentationMatte {
	rv := objc.Send[AVSemanticSegmentationMatte](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSemanticSegmentationMatte creates a new AVSemanticSegmentationMatte instance.
func NewAVSemanticSegmentationMatte() AVSemanticSegmentationMatte {
	class := getAVSemanticSegmentationMatteClass()
	rv := objc.Send[AVSemanticSegmentationMatte](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a new semantic segmentation matte instance from auxiliary image
// information in an image file.
//
// imageSourceAuxiliaryDataType: The `kCGImageAuxiliaryDataType` constants corresponding to the semantic
// segmentation matte being created (see [CGImageProperties]).
//
// imageSourceAuxiliaryDataInfoDictionary: A dictionary of primitive semantic segmentation matte information obtained
// from [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)].
// //
// [CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
//
// # Return Value
// 
// A new semantic segmentation matte instance, or `nil` if the auxiliary data
// info dictionary is malformed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/init(fromImageSourceAuxiliaryDataType:dictionaryRepresentation:)
func NewSemanticSegmentationMatteFromImageSourceAuxiliaryDataTypeDictionaryRepresentationError(imageSourceAuxiliaryDataType corefoundation.CFStringRef, imageSourceAuxiliaryDataInfoDictionary foundation.INSDictionary) (AVSemanticSegmentationMatte, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getAVSemanticSegmentationMatteClass().class), objc.Sel("semanticSegmentationMatteFromImageSourceAuxiliaryDataType:dictionaryRepresentation:error:"), imageSourceAuxiliaryDataType, imageSourceAuxiliaryDataInfoDictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVSemanticSegmentationMatte{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVSemanticSegmentationMatteFromID(rv), nil
}

// Returns a semantic segmentation matte instance that wraps the replacement
// pixel buffer.
//
// pixelBuffer: A pixel buffer containing a semantic segmentation matting image,
// represented as [KCVPixelFormatType_OneComponent8] with a
// [kCVImageBufferTransferFunction_Linear] transfer function.
// //
// [kCVImageBufferTransferFunction_Linear]: https://developer.apple.com/documentation/CoreVideo/kCVImageBufferTransferFunction_Linear
//
// # Return Value
// 
// A new semantic segmentation matte instance, or `nil` if the pixel buffer is
// malformed.
//
// # Discussion
// 
// When applying complex edits to media containing a semantic segmentation
// matte, you may create a derivative matte with arbitrary transforms applied
// to it. You can then use this method to create a new semantic segmentation
// matte instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/replacingSemanticSegmentationMatte(with:)
func (s AVSemanticSegmentationMatte) SemanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBufferError(pixelBuffer corevideo.CVImageBufferRef) (IAVSemanticSegmentationMatte, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("semanticSegmentationMatteByReplacingSemanticSegmentationMatteWithPixelBuffer:error:"), pixelBuffer, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVSemanticSegmentationMatte{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVSemanticSegmentationMatteFromID(rv), nil

}
// Returns a new semantic segmentation matte instance with the specified Exif
// orientation applied.
//
// exifOrientation: A [CGImagePropertyOrientation] value expressing how the matte should be
// rotated or mirrored.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// exifOrientation is a [imageio.CGImagePropertyOrientation].
//
// # Return Value
// 
// A new semantic segmentation matte instance.
//
// # Discussion
// 
// This method throws an [InvalidArgumentException] if you pass an
// unrecognized `exifOrientation`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/applyingExifOrientation(_:)
func (s AVSemanticSegmentationMatte) SemanticSegmentationMatteByApplyingExifOrientation(exifOrientation objectivec.IObject) IAVSemanticSegmentationMatte {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("semanticSegmentationMatteByApplyingExifOrientation:"), exifOrientation)
	return AVSemanticSegmentationMatteFromID(rv)
}
// Returns a dictionary of primitive map information to use when writing an
// image file with a semantic segmentation matte.
//
// outAuxDataType: On output, the auxiliary data type to be used when calling the ImageIO
// framework’s [CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)] function.
// Currently supported auxiliary data types are enumerated in
// [CGImageProperties].
// //
// [CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)]: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
//
// # Return Value
// 
// A dictionary of [CGImageDestination]-compatible semantic segmentation matte
// information, or `nil` if the auxiliary data type is unsupported.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/dictionaryRepresentation(forAuxiliaryDataType:)
func (s AVSemanticSegmentationMatte) DictionaryRepresentationForAuxiliaryDataType(outAuxDataType string) foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("dictionaryRepresentationForAuxiliaryDataType:"), objc.String(outAuxDataType))
	return foundation.NSDictionaryFromID(rv)
}

// The semantic segmentation matte image type.
//
// # Discussion
// 
// A semantic segmentation matte’s [MatteType] is immutable for the life of
// the object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/matteType-swift.property
func (s AVSemanticSegmentationMatte) MatteType() AVSemanticSegmentationMatteType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("matteType"))
	return AVSemanticSegmentationMatteType(foundation.NSStringFromID(rv).String())
}
// The semantic segmentation matte’s internal image.
//
// # Discussion
// 
// You can determine the pixel buffer’s format type using the
// [PixelFormatType] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/mattingImage
func (s AVSemanticSegmentationMatte) MattingImage() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](s.ID, objc.Sel("mattingImage"))
	return corevideo.CVImageBufferRef(rv)
}
// The pixel format type for this object’s internal matting image.
//
// # Discussion
// 
// Currently, the only supported pixel format type for the matting image is
// [KCVPixelFormatType_OneComponent8].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/pixelFormatType
func (s AVSemanticSegmentationMatte) PixelFormatType() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("pixelFormatType"))
	return rv
}
// 8-bit one component, black is zero.
//
// See: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
func (s AVSemanticSegmentationMatte) KCVPixelFormatType_OneComponent8() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("kCVPixelFormatType_OneComponent8"))
	return rv
}
func (s AVSemanticSegmentationMatte) SetKCVPixelFormatType_OneComponent8(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setKCVPixelFormatType_OneComponent8:"), value)
}

