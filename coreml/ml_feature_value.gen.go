// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureValue] class.
var (
	_MLFeatureValueClass     MLFeatureValueClass
	_MLFeatureValueClassOnce sync.Once
)

func getMLFeatureValueClass() MLFeatureValueClass {
	_MLFeatureValueClassOnce.Do(func() {
		_MLFeatureValueClass = MLFeatureValueClass{class: objc.GetClass("MLFeatureValue")}
	})
	return _MLFeatureValueClass
}

// GetMLFeatureValueClass returns the class object for MLFeatureValue.
func GetMLFeatureValueClass() MLFeatureValueClass {
	return getMLFeatureValueClass()
}

type MLFeatureValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureValueClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureValueClass) Alloc() MLFeatureValue {
	rv := objc.Send[MLFeatureValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A generic wrapper around an underlying value and the value’s type.
//
// # Overview
// 
// A Core ML wraps an underlying value and bundles it with that value’s
// type, which is one of the types that [MLFeatureType] defines. Apps
// typically access feature values indirectly by using the methods in the
// wrapper class Xcode automatically generates for Core ML model files.
// 
// If your app accesses an [MLModel] directly, it must create and consume
// [MLFeatureProvider] instances. For each prediction, Core ML accepts a
// feature provider for its inputs, and generates a separate feature provider
// for its outputs. The input feature provider contains one [MLFeatureValue]
// instance per input, and the output feature provider contains one per
// output. See [MLFeatureDescription] for more information about the model
// input and output features.
//
// [MLFeatureType]: https://developer.apple.com/documentation/CoreML/MLFeatureType
//
// # Accessing the feature’s type
//
//   - [MLFeatureValue.Type]: The type of the feature value.
//
// # Accessing the feature’s value
//
//   - [MLFeatureValue.Undefined]: A Boolean value that indicates whether the feature value is undefined or missing.
//   - [MLFeatureValue.Int64Value]: The underlying integer of the feature value.
//   - [MLFeatureValue.DoubleValue]: The underlying double of the feature value.
//   - [MLFeatureValue.StringValue]: The underlying string of the feature value.
//   - [MLFeatureValue.ImageBufferValue]: The underlying image of the feature value as a pixel buffer.
//   - [MLFeatureValue.MultiArrayValue]: The underlying multiarray of the feature value.
//   - [MLFeatureValue.SequenceValue]: The underlying sequence of the feature value.
//   - [MLFeatureValue.DictionaryValue]: The underlying dictionary of the feature value.
//
// # Comparing feature values
//
//   - [MLFeatureValue.IsEqualToFeatureValue]: Returns a Boolean value that indicates whether a feature value is equal to another.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue
type MLFeatureValue struct {
	objectivec.Object
}

// MLFeatureValueFromID constructs a [MLFeatureValue] from an objc.ID.
//
// A generic wrapper around an underlying value and the value’s type.
func MLFeatureValueFromID(id objc.ID) MLFeatureValue {
	return MLFeatureValue{objectivec.Object{ID: id}}
}
// NOTE: MLFeatureValue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLFeatureValue] class.
//
// # Accessing the feature’s type
//
//   - [IMLFeatureValue.Type]: The type of the feature value.
//
// # Accessing the feature’s value
//
//   - [IMLFeatureValue.Undefined]: A Boolean value that indicates whether the feature value is undefined or missing.
//   - [IMLFeatureValue.Int64Value]: The underlying integer of the feature value.
//   - [IMLFeatureValue.DoubleValue]: The underlying double of the feature value.
//   - [IMLFeatureValue.StringValue]: The underlying string of the feature value.
//   - [IMLFeatureValue.ImageBufferValue]: The underlying image of the feature value as a pixel buffer.
//   - [IMLFeatureValue.MultiArrayValue]: The underlying multiarray of the feature value.
//   - [IMLFeatureValue.SequenceValue]: The underlying sequence of the feature value.
//   - [IMLFeatureValue.DictionaryValue]: The underlying dictionary of the feature value.
//
// # Comparing feature values
//
//   - [IMLFeatureValue.IsEqualToFeatureValue]: Returns a Boolean value that indicates whether a feature value is equal to another.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue
type IMLFeatureValue interface {
	objectivec.IObject

	// Topic: Accessing the feature’s type

	// The type of the feature value.
	Type() MLFeatureType

	// Topic: Accessing the feature’s value

	// A Boolean value that indicates whether the feature value is undefined or missing.
	Undefined() bool
	// The underlying integer of the feature value.
	Int64Value() int64
	// The underlying double of the feature value.
	DoubleValue() float64
	// The underlying string of the feature value.
	StringValue() string
	// The underlying image of the feature value as a pixel buffer.
	ImageBufferValue() corevideo.CVImageBufferRef
	// The underlying multiarray of the feature value.
	MultiArrayValue() IMLMultiArray
	// The underlying sequence of the feature value.
	SequenceValue() IMLSequence
	// The underlying dictionary of the feature value.
	DictionaryValue() foundation.INSDictionary

	// Topic: Comparing feature values

	// Returns a Boolean value that indicates whether a feature value is equal to another.
	IsEqualToFeatureValue(value IMLFeatureValue) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (f MLFeatureValue) Init() MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureValue) Autorelease() MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureValue creates a new MLFeatureValue instance.
func NewMLFeatureValue() MLFeatureValue {
	class := getMLFeatureValueClass()
	rv := objc.Send[MLFeatureValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a feature value with a type that represents an undefined or missing
// value.
//
// type: The type of the undefined or missing value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(undefined:)
func NewFeatureValueUndefinedFeatureValueWithType(type_ MLFeatureType) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("undefinedFeatureValueWithType:"), type_)
	return MLFeatureValueFromID(rv)
}

// Creates a feature value that contains an image defined by a core graphics
// image and a constraint.
//
// cgImage: A [CGImage] instance.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// constraint: An [MLImageConstraint] instance.
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:constraint:options:)
func NewFeatureValueWithCGImageConstraintOptionsError(cgImage coregraphics.CGImageRef, constraint IMLImageConstraint, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithCGImage:constraint:options:error:"), cgImage, constraint, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an image defined by a core graphics
// image, an orientation, and a constraint.
//
// cgImage: A [CGImage] instance.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// orientation: A [CGImagePropertyOrientation] instance.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// constraint: An [MLImageConstraint] instance.
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:constraint:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewFeatureValueWithCGImageOrientationConstraintOptionsError(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, constraint IMLImageConstraint, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithCGImage:orientation:constraint:options:error:"), cgImage, orientation, constraint, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an image defined by a core graphics
// image and its orientation, size, and pixel format.
//
// cgImage: A [CGImage] instance.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// orientation: A [CGImagePropertyOrientation] instance.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// pixelsWide: The image’s width in pixels.
//
// pixelsHigh: The image’s height in pixels.
//
// pixelFormatType: The image’s pixel format (see [Pixel Format Identifiers]).
// //
// [Pixel Format Identifiers]: https://developer.apple.com/documentation/CoreVideo/pixel-format-identifiers
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewFeatureValueWithCGImageOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage coregraphics.CGImageRef, orientation objectivec.IObject, pixelsWide int, pixelsHigh int, pixelFormatType uint32, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithCGImage:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an image defined by a core graphics
// image and its size and pixel format.
//
// cgImage: A [CGImage] instance.
// //
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// pixelsWide: The image’s width in pixels.
//
// pixelsHigh: The image’s height in pixels.
//
// pixelFormatType: The image’s pixel format (see [Pixel Format Identifiers]).
// //
// [Pixel Format Identifiers]: https://developer.apple.com/documentation/CoreVideo/pixel-format-identifiers
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(CGImage:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithCGImagePixelsWidePixelsHighPixelFormatTypeOptionsError(cgImage coregraphics.CGImageRef, pixelsWide int, pixelsHigh int, pixelFormatType uint32, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithCGImage:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), cgImage, pixelsWide, pixelsHigh, pixelFormatType, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains a dictionary of numbers.
//
// value: A dictionary of numbers.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(dictionary:)
func NewFeatureValueWithDictionaryError(value foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithDictionary:error:"), value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains a double.
//
// value: A double precision floating point value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(double:)
func NewFeatureValueWithDouble(value float64) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithDouble:"), value)
	return MLFeatureValueFromID(rv)
}

// Creates a feature value that contains an image defined by an image URL and
// a constraint.
//
// url: A [URL] (Swift) or [NSURL] (Objective-C) to an image.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [URL]: https://developer.apple.com/documentation/Foundation/URL
//
// constraint: An [MLImageConstraint] instance.
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:constraint:options:)
func NewFeatureValueWithImageAtURLConstraintOptionsError(url foundation.INSURL, constraint IMLImageConstraint, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), url, constraint, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an image defined by an image URL, an
// orientation, and a constraint.
//
// url: A [URL] (Swift) or [NSURL] (Objective-C) to an image.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [URL]: https://developer.apple.com/documentation/Foundation/URL
//
// orientation: A [CGImagePropertyOrientation] instance.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// constraint: An [MLImageConstraint] instance.
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:constraint:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewFeatureValueWithImageAtURLOrientationConstraintOptionsError(url foundation.INSURL, orientation objectivec.IObject, constraint IMLImageConstraint, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:orientation:constraint:options:error:"), url, orientation, constraint, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an image defined by an image URL and
// the image’s orientation, size, and pixel format.
//
// url: A [URL] (Swift) or [NSURL] (Objective-C) to an image.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [URL]: https://developer.apple.com/documentation/Foundation/URL
//
// orientation: A [CGImagePropertyOrientation] instance.
// //
// [CGImagePropertyOrientation]: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
//
// pixelsWide: The image’s width in pixels.
//
// pixelsHigh: The image’s height in pixels.
//
// pixelFormatType: The image’s pixel format (see [Pixel Format Identifiers]).
// //
// [Pixel Format Identifiers]: https://developer.apple.com/documentation/CoreVideo/pixel-format-identifiers
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:)
// orientation is a [imageio.CGImagePropertyOrientation].
func NewFeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.INSURL, orientation objectivec.IObject, pixelsWide int, pixelsHigh int, pixelFormatType uint32, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, orientation, pixelsWide, pixelsHigh, pixelFormatType, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an image defined by an image URL and
// the image’s size and pixel format.
//
// url: A [URL] (Swift) or [NSURL] (Objective-C) to an image.
// //
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [URL]: https://developer.apple.com/documentation/Foundation/URL
//
// pixelsWide: The image’s width in pixels.
//
// pixelsHigh: The image’s height in pixels.
//
// pixelFormatType: The image’s pixel format (see [Pixel Format Identifiers]).
// //
// [Pixel Format Identifiers]: https://developer.apple.com/documentation/CoreVideo/pixel-format-identifiers
//
// options: A dictionary of [VNImageCropAndScaleOption] values, each keyed by
// [MLFeatureValueImageOption].
// //
// [VNImageCropAndScaleOption]: https://developer.apple.com/documentation/Vision/VNImageCropAndScaleOption
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(imageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:)
func NewFeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.INSURL, pixelsWide int, pixelsHigh int, pixelFormatType uint32, options foundation.INSDictionary) (MLFeatureValue, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, pixelsWide, pixelsHigh, pixelFormatType, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLFeatureValue{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLFeatureValueFromID(rv), nil
}

// Creates a feature value that contains an integer.
//
// value: A 64-bit integer value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(int64:)
func NewFeatureValueWithInt64(value int64) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithInt64:"), value)
	return MLFeatureValueFromID(rv)
}

// Creates a feature value that contains a multidimensional array.
//
// value: An [MLMultiArray] instance.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(multiArray:)
func NewFeatureValueWithMultiArray(value IMLMultiArray) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithMultiArray:"), value)
	return MLFeatureValueFromID(rv)
}

// Creates a feature value that contains an image from a pixel buffer.
//
// value: A [CVPixelBuffer] (Swift) or [CVPixelBuffer] (Objective-C) instance.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// # Discussion
// 
// [Core ML] supports different pixel format types depending on the model’s
// feature description. For information about [ImageFeatureType], see [Core ML
// Format Reference]. When the image feature’s color space is [GRAYSCALE],
// use [kCVPixelFormatType_OneComponent8]; and when it’s
// `GRAYSCALE_FLOAT16`, use [kCVPixelFormatType_OneComponent16Half];
// otherwise, use [kCVPixelFormatType_32BGRA] when it’s set to [RGB] or
// [BGR].
//
// [Core ML Format Reference]: https://apple.github.io/coremltools/mlmodel/Format/FeatureTypes.html#imagefeaturetype
// [Core ML]: https://developer.apple.com/documentation/CoreML
// [kCVPixelFormatType_32BGRA]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_32BGRA
// [kCVPixelFormatType_OneComponent16Half]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent16Half
// [kCVPixelFormatType_OneComponent8]: https://developer.apple.com/documentation/CoreVideo/kCVPixelFormatType_OneComponent8
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(pixelBuffer:)
func NewFeatureValueWithPixelBuffer(value corevideo.CVImageBufferRef) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithPixelBuffer:"), value)
	return MLFeatureValueFromID(rv)
}

// Creates a feature value that contains a sequence.
//
// sequence: An [MLSequence] instance.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(sequence:)
func NewFeatureValueWithSequence(sequence IMLSequence) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithSequence:"), sequence)
	return MLFeatureValueFromID(rv)
}

// Creates a feature value that contains a string.
//
// value: A string.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/init(string:)
func NewFeatureValueWithString(value string) MLFeatureValue {
	rv := objc.Send[objc.ID](objc.ID(getMLFeatureValueClass().class), objc.Sel("featureValueWithString:"), objc.String(value))
	return MLFeatureValueFromID(rv)
}

// Returns a Boolean value that indicates whether a feature value is equal to
// another.
//
// value: Another feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/isEqual(to:)
func (f MLFeatureValue) IsEqualToFeatureValue(value IMLFeatureValue) bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isEqualToFeatureValue:"), value)
	return rv
}
func (f MLFeatureValue) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](f.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The type of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/type
func (f MLFeatureValue) Type() MLFeatureType {
	rv := objc.Send[MLFeatureType](f.ID, objc.Sel("type"))
	return MLFeatureType(rv)
}
// A Boolean value that indicates whether the feature value is undefined or
// missing.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/isUndefined
func (f MLFeatureValue) Undefined() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isUndefined"))
	return rv
}
// The underlying integer of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/int64Value
func (f MLFeatureValue) Int64Value() int64 {
	rv := objc.Send[int64](f.ID, objc.Sel("int64Value"))
	return rv
}
// The underlying double of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/doubleValue
func (f MLFeatureValue) DoubleValue() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("doubleValue"))
	return rv
}
// The underlying string of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/stringValue
func (f MLFeatureValue) StringValue() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stringValue"))
	return foundation.NSStringFromID(rv).String()
}
// The underlying image of the feature value as a pixel buffer.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/imageBufferValue
func (f MLFeatureValue) ImageBufferValue() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](f.ID, objc.Sel("imageBufferValue"))
	return corevideo.CVImageBufferRef(rv)
}
// The underlying multiarray of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/multiArrayValue
func (f MLFeatureValue) MultiArrayValue() IMLMultiArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("multiArrayValue"))
	return MLMultiArrayFromID(objc.ID(rv))
}
// The underlying sequence of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/sequenceValue
func (f MLFeatureValue) SequenceValue() IMLSequence {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("sequenceValue"))
	return MLSequenceFromID(objc.ID(rv))
}
// The underlying dictionary of the feature value.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/dictionaryValue
func (f MLFeatureValue) DictionaryValue() foundation.INSDictionary {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("dictionaryValue"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

