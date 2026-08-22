// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objc"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("ImageIO: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("ImageIO: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("ImageIO: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("ImageIO: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _cGAnimateImageAtURLWithBlock func(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef, block unsafe.Pointer) int32
var _cGAnimateImageAtURLWithBlockErr error

func tryCGAnimateImageAtURLWithBlock(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef, block CGImageSourceAnimationBlock) (int32, error) {
	if _cGAnimateImageAtURLWithBlock == nil {
		return 0, symbolCallError("CGAnimateImageAtURLWithBlock", "10.15", _cGAnimateImageAtURLWithBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 uint32, blockArg1 unsafe.Pointer, blockArg2 *bool) {
		block(blockArg0, (*coregraphics.CGImageRef)(blockArg1), blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cGAnimateImageAtURLWithBlock(url, options, _block0), nil
}

// CGAnimateImageAtURLWithBlock animate the sequence of images in the Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file at the specified URL.
//
// See: https://developer.apple.com/documentation/ImageIO/CGAnimateImageAtURLWithBlock(_:_:_:)
func CGAnimateImageAtURLWithBlock(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef, block CGImageSourceAnimationBlock) int32 {
	result, callErr := tryCGAnimateImageAtURLWithBlock(url, options, block)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGAnimateImageDataWithBlock func(data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, block unsafe.Pointer) int32
var _cGAnimateImageDataWithBlockErr error

func tryCGAnimateImageDataWithBlock(data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, block CGImageSourceAnimationBlock) (int32, error) {
	if _cGAnimateImageDataWithBlock == nil {
		return 0, symbolCallError("CGAnimateImageDataWithBlock", "10.15", _cGAnimateImageDataWithBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 uint32, blockArg1 unsafe.Pointer, blockArg2 *bool) {
		block(blockArg0, (*coregraphics.CGImageRef)(blockArg1), blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cGAnimateImageDataWithBlock(data, options, _block0), nil
}

// CGAnimateImageDataWithBlock animate the sequence of images using data from a Graphics Interchange Format (GIF) or Animated Portable Network Graphics (APNG) file file.
//
// See: https://developer.apple.com/documentation/ImageIO/CGAnimateImageDataWithBlock(_:_:_:)
func CGAnimateImageDataWithBlock(data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef, block CGImageSourceAnimationBlock) int32 {
	result, callErr := tryCGAnimateImageDataWithBlock(data, options, block)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationAddAuxiliaryDataInfo func(idst CGImageDestinationRef, auxiliaryImageDataType corefoundation.CFStringRef, auxiliaryDataInfoDictionary corefoundation.CFDictionaryRef)
var _cGImageDestinationAddAuxiliaryDataInfoErr error

func tryCGImageDestinationAddAuxiliaryDataInfo(idst CGImageDestinationRef, auxiliaryImageDataType corefoundation.CFStringRef, auxiliaryDataInfoDictionary corefoundation.CFDictionaryRef) error {
	if _cGImageDestinationAddAuxiliaryDataInfo == nil {
		return symbolCallError("CGImageDestinationAddAuxiliaryDataInfo", "10.13", _cGImageDestinationAddAuxiliaryDataInfoErr)
	}
	_cGImageDestinationAddAuxiliaryDataInfo(idst, auxiliaryImageDataType, auxiliaryDataInfoDictionary)
	return nil
}

// CGImageDestinationAddAuxiliaryDataInfo sets the auxiliary data, such as mattes and depth information, that accompany the image.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddAuxiliaryDataInfo(_:_:_:)
func CGImageDestinationAddAuxiliaryDataInfo(idst CGImageDestinationRef, auxiliaryImageDataType corefoundation.CFStringRef, auxiliaryDataInfoDictionary corefoundation.CFDictionaryRef) {
	if callErr := tryCGImageDestinationAddAuxiliaryDataInfo(idst, auxiliaryImageDataType, auxiliaryDataInfoDictionary); callErr != nil {
		panic(callErr)
	}
}

var _cGImageDestinationAddImage func(idst CGImageDestinationRef, image coregraphics.CGImageRef, properties corefoundation.CFDictionaryRef)
var _cGImageDestinationAddImageErr error

func tryCGImageDestinationAddImage(idst CGImageDestinationRef, image coregraphics.CGImageRef, properties corefoundation.CFDictionaryRef) error {
	if _cGImageDestinationAddImage == nil {
		return symbolCallError("CGImageDestinationAddImage", "10.4", _cGImageDestinationAddImageErr)
	}
	_cGImageDestinationAddImage(idst, image, properties)
	return nil
}

// CGImageDestinationAddImage adds an image to an image destination.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImage(_:_:_:)
func CGImageDestinationAddImage(idst CGImageDestinationRef, image coregraphics.CGImageRef, properties corefoundation.CFDictionaryRef) {
	if callErr := tryCGImageDestinationAddImage(idst, image, properties); callErr != nil {
		panic(callErr)
	}
}

var _cGImageDestinationAddImageAndMetadata func(idst CGImageDestinationRef, image coregraphics.CGImageRef, metadata CGImageMetadataRef, options corefoundation.CFDictionaryRef)
var _cGImageDestinationAddImageAndMetadataErr error

func tryCGImageDestinationAddImageAndMetadata(idst CGImageDestinationRef, image coregraphics.CGImageRef, metadata CGImageMetadataRef, options corefoundation.CFDictionaryRef) error {
	if _cGImageDestinationAddImageAndMetadata == nil {
		return symbolCallError("CGImageDestinationAddImageAndMetadata", "10.8", _cGImageDestinationAddImageAndMetadataErr)
	}
	_cGImageDestinationAddImageAndMetadata(idst, image, metadata, options)
	return nil
}

// CGImageDestinationAddImageAndMetadata.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageAndMetadata(_:_:_:_:)
func CGImageDestinationAddImageAndMetadata(idst CGImageDestinationRef, image coregraphics.CGImageRef, metadata CGImageMetadataRef, options corefoundation.CFDictionaryRef) {
	if callErr := tryCGImageDestinationAddImageAndMetadata(idst, image, metadata, options); callErr != nil {
		panic(callErr)
	}
}

var _cGImageDestinationAddImageFromSource func(idst CGImageDestinationRef, isrc CGImageSourceRef, index uintptr, properties corefoundation.CFDictionaryRef)
var _cGImageDestinationAddImageFromSourceErr error

func tryCGImageDestinationAddImageFromSource(idst CGImageDestinationRef, isrc CGImageSourceRef, index uintptr, properties corefoundation.CFDictionaryRef) error {
	if _cGImageDestinationAddImageFromSource == nil {
		return symbolCallError("CGImageDestinationAddImageFromSource", "10.4", _cGImageDestinationAddImageFromSourceErr)
	}
	_cGImageDestinationAddImageFromSource(idst, isrc, index, properties)
	return nil
}

// CGImageDestinationAddImageFromSource adds an image from an image source to an image destination.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationAddImageFromSource(_:_:_:_:)
func CGImageDestinationAddImageFromSource(idst CGImageDestinationRef, isrc CGImageSourceRef, index uintptr, properties corefoundation.CFDictionaryRef) {
	if callErr := tryCGImageDestinationAddImageFromSource(idst, isrc, index, properties); callErr != nil {
		panic(callErr)
	}
}

var _cGImageDestinationCopyImageSource func(idst CGImageDestinationRef, isrc CGImageSourceRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool
var _cGImageDestinationCopyImageSourceErr error

func tryCGImageDestinationCopyImageSource(idst CGImageDestinationRef, isrc CGImageSourceRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _cGImageDestinationCopyImageSource == nil {
		return false, symbolCallError("CGImageDestinationCopyImageSource", "10.8", _cGImageDestinationCopyImageSourceErr)
	}
	return _cGImageDestinationCopyImageSource(idst, isrc, options, err), nil
}

// CGImageDestinationCopyImageSource.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyImageSource(_:_:_:_:)
func CGImageDestinationCopyImageSource(idst CGImageDestinationRef, isrc CGImageSourceRef, options corefoundation.CFDictionaryRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryCGImageDestinationCopyImageSource(idst, isrc, options, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationCopyTypeIdentifiers func() corefoundation.CFArrayRef
var _cGImageDestinationCopyTypeIdentifiersErr error

func tryCGImageDestinationCopyTypeIdentifiers() (corefoundation.CFArrayRef, error) {
	if _cGImageDestinationCopyTypeIdentifiers == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CGImageDestinationCopyTypeIdentifiers", "10.4", _cGImageDestinationCopyTypeIdentifiersErr)
	}
	return _cGImageDestinationCopyTypeIdentifiers(), nil
}

// CGImageDestinationCopyTypeIdentifiers returns an array of the uniform type identifiers that are supported for image destinations.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCopyTypeIdentifiers()
func CGImageDestinationCopyTypeIdentifiers() corefoundation.CFArrayRef {
	result, callErr := tryCGImageDestinationCopyTypeIdentifiers()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationCreateWithData func(data corefoundation.CFMutableDataRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) CGImageDestinationRef
var _cGImageDestinationCreateWithDataErr error

func tryCGImageDestinationCreateWithData(data corefoundation.CFMutableDataRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) (CGImageDestinationRef, error) {
	if _cGImageDestinationCreateWithData == nil {
		return *new(CGImageDestinationRef), symbolCallError("CGImageDestinationCreateWithData", "10.4", _cGImageDestinationCreateWithDataErr)
	}
	return _cGImageDestinationCreateWithData(data, type_, count, options), nil
}

// CGImageDestinationCreateWithData creates an image destination that writes to a Core Foundation mutable data object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithData(_:_:_:_:)
func CGImageDestinationCreateWithData(data corefoundation.CFMutableDataRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) CGImageDestinationRef {
	result, callErr := tryCGImageDestinationCreateWithData(data, type_, count, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationCreateWithDataConsumer func(consumer coregraphics.CGDataConsumerRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) CGImageDestinationRef
var _cGImageDestinationCreateWithDataConsumerErr error

func tryCGImageDestinationCreateWithDataConsumer(consumer coregraphics.CGDataConsumerRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) (CGImageDestinationRef, error) {
	if _cGImageDestinationCreateWithDataConsumer == nil {
		return *new(CGImageDestinationRef), symbolCallError("CGImageDestinationCreateWithDataConsumer", "10.4", _cGImageDestinationCreateWithDataConsumerErr)
	}
	return _cGImageDestinationCreateWithDataConsumer(consumer, type_, count, options), nil
}

// CGImageDestinationCreateWithDataConsumer creates an image destination that writes to the specified data consumer.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithDataConsumer(_:_:_:_:)
func CGImageDestinationCreateWithDataConsumer(consumer coregraphics.CGDataConsumerRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) CGImageDestinationRef {
	result, callErr := tryCGImageDestinationCreateWithDataConsumer(consumer, type_, count, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationCreateWithURL func(url corefoundation.CFURLRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) CGImageDestinationRef
var _cGImageDestinationCreateWithURLErr error

func tryCGImageDestinationCreateWithURL(url corefoundation.CFURLRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) (CGImageDestinationRef, error) {
	if _cGImageDestinationCreateWithURL == nil {
		return *new(CGImageDestinationRef), symbolCallError("CGImageDestinationCreateWithURL", "10.4", _cGImageDestinationCreateWithURLErr)
	}
	return _cGImageDestinationCreateWithURL(url, type_, count, options), nil
}

// CGImageDestinationCreateWithURL creates an image destination that writes image data to the specified URL.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationCreateWithURL(_:_:_:_:)
func CGImageDestinationCreateWithURL(url corefoundation.CFURLRef, type_ corefoundation.CFStringRef, count uintptr, options corefoundation.CFDictionaryRef) CGImageDestinationRef {
	result, callErr := tryCGImageDestinationCreateWithURL(url, type_, count, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationFinalize func(idst CGImageDestinationRef) bool
var _cGImageDestinationFinalizeErr error

func tryCGImageDestinationFinalize(idst CGImageDestinationRef) (bool, error) {
	if _cGImageDestinationFinalize == nil {
		return false, symbolCallError("CGImageDestinationFinalize", "10.4", _cGImageDestinationFinalizeErr)
	}
	return _cGImageDestinationFinalize(idst), nil
}

// CGImageDestinationFinalize writes image data and properties to the data, URL, or data consumer associated with the image destination.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationFinalize(_:)
func CGImageDestinationFinalize(idst CGImageDestinationRef) bool {
	result, callErr := tryCGImageDestinationFinalize(idst)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationGetTypeID func() uint
var _cGImageDestinationGetTypeIDErr error

func tryCGImageDestinationGetTypeID() (uint, error) {
	if _cGImageDestinationGetTypeID == nil {
		return 0, symbolCallError("CGImageDestinationGetTypeID", "10.4", _cGImageDestinationGetTypeIDErr)
	}
	return _cGImageDestinationGetTypeID(), nil
}

// CGImageDestinationGetTypeID returns the unique type identifier of an image destination opaque type.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationGetTypeID()
func CGImageDestinationGetTypeID() uint {
	result, callErr := tryCGImageDestinationGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageDestinationSetProperties func(idst CGImageDestinationRef, properties corefoundation.CFDictionaryRef)
var _cGImageDestinationSetPropertiesErr error

func tryCGImageDestinationSetProperties(idst CGImageDestinationRef, properties corefoundation.CFDictionaryRef) error {
	if _cGImageDestinationSetProperties == nil {
		return symbolCallError("CGImageDestinationSetProperties", "10.4", _cGImageDestinationSetPropertiesErr)
	}
	_cGImageDestinationSetProperties(idst, properties)
	return nil
}

// CGImageDestinationSetProperties applies one or more properties to all images in an image destination.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageDestinationSetProperties(_:_:)
func CGImageDestinationSetProperties(idst CGImageDestinationRef, properties corefoundation.CFDictionaryRef) {
	if callErr := tryCGImageDestinationSetProperties(idst, properties); callErr != nil {
		panic(callErr)
	}
}

var _cGImageMetadataCopyStringValueWithPath func(metadata CGImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) corefoundation.CFStringRef
var _cGImageMetadataCopyStringValueWithPathErr error

func tryCGImageMetadataCopyStringValueWithPath(metadata CGImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _cGImageMetadataCopyStringValueWithPath == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CGImageMetadataCopyStringValueWithPath", "10.8", _cGImageMetadataCopyStringValueWithPathErr)
	}
	return _cGImageMetadataCopyStringValueWithPath(metadata, parent, path), nil
}

// CGImageMetadataCopyStringValueWithPath searches the metadata for the specified tag, and returns its string value if it exists.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyStringValueWithPath(_:_:_:)
func CGImageMetadataCopyStringValueWithPath(metadata CGImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryCGImageMetadataCopyStringValueWithPath(metadata, parent, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCopyTagMatchingImageProperty func(metadata CGImageMetadataRef, dictionaryName corefoundation.CFStringRef, propertyName corefoundation.CFStringRef) CGImageMetadataTagRef
var _cGImageMetadataCopyTagMatchingImagePropertyErr error

func tryCGImageMetadataCopyTagMatchingImageProperty(metadata CGImageMetadataRef, dictionaryName corefoundation.CFStringRef, propertyName corefoundation.CFStringRef) (CGImageMetadataTagRef, error) {
	if _cGImageMetadataCopyTagMatchingImageProperty == nil {
		return *new(CGImageMetadataTagRef), symbolCallError("CGImageMetadataCopyTagMatchingImageProperty", "10.8", _cGImageMetadataCopyTagMatchingImagePropertyErr)
	}
	return _cGImageMetadataCopyTagMatchingImageProperty(metadata, dictionaryName, propertyName), nil
}

// CGImageMetadataCopyTagMatchingImageProperty searches for the specified image property and, if found, returns the corresponding tag object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTagMatchingImageProperty(_:_:_:)
func CGImageMetadataCopyTagMatchingImageProperty(metadata CGImageMetadataRef, dictionaryName corefoundation.CFStringRef, propertyName corefoundation.CFStringRef) CGImageMetadataTagRef {
	result, callErr := tryCGImageMetadataCopyTagMatchingImageProperty(metadata, dictionaryName, propertyName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCopyTagWithPath func(metadata CGImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) CGImageMetadataTagRef
var _cGImageMetadataCopyTagWithPathErr error

func tryCGImageMetadataCopyTagWithPath(metadata CGImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) (CGImageMetadataTagRef, error) {
	if _cGImageMetadataCopyTagWithPath == nil {
		return *new(CGImageMetadataTagRef), symbolCallError("CGImageMetadataCopyTagWithPath", "10.8", _cGImageMetadataCopyTagWithPathErr)
	}
	return _cGImageMetadataCopyTagWithPath(metadata, parent, path), nil
}

// CGImageMetadataCopyTagWithPath searches for a specific metadata tag within a metadata collection.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTagWithPath(_:_:_:)
func CGImageMetadataCopyTagWithPath(metadata CGImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) CGImageMetadataTagRef {
	result, callErr := tryCGImageMetadataCopyTagWithPath(metadata, parent, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCopyTags func(metadata CGImageMetadataRef) corefoundation.CFArrayRef
var _cGImageMetadataCopyTagsErr error

func tryCGImageMetadataCopyTags(metadata CGImageMetadataRef) (corefoundation.CFArrayRef, error) {
	if _cGImageMetadataCopyTags == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CGImageMetadataCopyTags", "10.8", _cGImageMetadataCopyTagsErr)
	}
	return _cGImageMetadataCopyTags(metadata), nil
}

// CGImageMetadataCopyTags returns an array of root-level metadata tags from the specified metadata object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCopyTags(_:)
func CGImageMetadataCopyTags(metadata CGImageMetadataRef) corefoundation.CFArrayRef {
	result, callErr := tryCGImageMetadataCopyTags(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCreateFromXMPData func(data corefoundation.CFDataRef) CGImageMetadataRef
var _cGImageMetadataCreateFromXMPDataErr error

func tryCGImageMetadataCreateFromXMPData(data corefoundation.CFDataRef) (CGImageMetadataRef, error) {
	if _cGImageMetadataCreateFromXMPData == nil {
		return *new(CGImageMetadataRef), symbolCallError("CGImageMetadataCreateFromXMPData", "10.8", _cGImageMetadataCreateFromXMPDataErr)
	}
	return _cGImageMetadataCreateFromXMPData(data), nil
}

// CGImageMetadataCreateFromXMPData creates a collection of metadata tags from the specified XMP data.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateFromXMPData(_:)
func CGImageMetadataCreateFromXMPData(data corefoundation.CFDataRef) CGImageMetadataRef {
	result, callErr := tryCGImageMetadataCreateFromXMPData(data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCreateMutable func() CGMutableImageMetadataRef
var _cGImageMetadataCreateMutableErr error

func tryCGImageMetadataCreateMutable() (CGMutableImageMetadataRef, error) {
	if _cGImageMetadataCreateMutable == nil {
		return *new(CGMutableImageMetadataRef), symbolCallError("CGImageMetadataCreateMutable", "10.8", _cGImageMetadataCreateMutableErr)
	}
	return _cGImageMetadataCreateMutable(), nil
}

// CGImageMetadataCreateMutable creates an empty, mutable image metdata opaque type.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutable()
func CGImageMetadataCreateMutable() CGMutableImageMetadataRef {
	result, callErr := tryCGImageMetadataCreateMutable()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCreateMutableCopy func(metadata CGImageMetadataRef) CGMutableImageMetadataRef
var _cGImageMetadataCreateMutableCopyErr error

func tryCGImageMetadataCreateMutableCopy(metadata CGImageMetadataRef) (CGMutableImageMetadataRef, error) {
	if _cGImageMetadataCreateMutableCopy == nil {
		return *new(CGMutableImageMetadataRef), symbolCallError("CGImageMetadataCreateMutableCopy", "10.8", _cGImageMetadataCreateMutableCopyErr)
	}
	return _cGImageMetadataCreateMutableCopy(metadata), nil
}

// CGImageMetadataCreateMutableCopy creates a deep, mutable copy of the specified metadata information.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateMutableCopy(_:)
func CGImageMetadataCreateMutableCopy(metadata CGImageMetadataRef) CGMutableImageMetadataRef {
	result, callErr := tryCGImageMetadataCreateMutableCopy(metadata)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataCreateXMPData func(metadata CGImageMetadataRef, options corefoundation.CFDictionaryRef) corefoundation.CFDataRef
var _cGImageMetadataCreateXMPDataErr error

func tryCGImageMetadataCreateXMPData(metadata CGImageMetadataRef, options corefoundation.CFDictionaryRef) (corefoundation.CFDataRef, error) {
	if _cGImageMetadataCreateXMPData == nil {
		return *new(corefoundation.CFDataRef), symbolCallError("CGImageMetadataCreateXMPData", "10.8", _cGImageMetadataCreateXMPDataErr)
	}
	return _cGImageMetadataCreateXMPData(metadata, options), nil
}

// CGImageMetadataCreateXMPData returns a data object that contains the metadata object’s contents serialized into the XMP format.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataCreateXMPData(_:_:)
func CGImageMetadataCreateXMPData(metadata CGImageMetadataRef, options corefoundation.CFDictionaryRef) corefoundation.CFDataRef {
	result, callErr := tryCGImageMetadataCreateXMPData(metadata, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataEnumerateTagsUsingBlock func(metadata CGImageMetadataRef, rootPath corefoundation.CFStringRef, options corefoundation.CFDictionaryRef, block unsafe.Pointer)
var _cGImageMetadataEnumerateTagsUsingBlockErr error

func tryCGImageMetadataEnumerateTagsUsingBlock(metadata CGImageMetadataRef, rootPath corefoundation.CFStringRef, options corefoundation.CFDictionaryRef, block CGImageMetadataTagBlock) error {
	if _cGImageMetadataEnumerateTagsUsingBlock == nil {
		return symbolCallError("CGImageMetadataEnumerateTagsUsingBlock", "10.8", _cGImageMetadataEnumerateTagsUsingBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 corefoundation.CFStringRef, blockArg1 *CGImageMetadataTagRef) bool {
		return block(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	_cGImageMetadataEnumerateTagsUsingBlock(metadata, rootPath, options, _block0)
	return nil
}

// CGImageMetadataEnumerateTagsUsingBlock enumerates the tags of a metadata object and executes the specified block on each tag.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataEnumerateTagsUsingBlock(_:_:_:_:)
func CGImageMetadataEnumerateTagsUsingBlock(metadata CGImageMetadataRef, rootPath corefoundation.CFStringRef, options corefoundation.CFDictionaryRef, block CGImageMetadataTagBlock) {
	if callErr := tryCGImageMetadataEnumerateTagsUsingBlock(metadata, rootPath, options, block); callErr != nil {
		panic(callErr)
	}
}

var _cGImageMetadataGetTypeID func() uint
var _cGImageMetadataGetTypeIDErr error

func tryCGImageMetadataGetTypeID() (uint, error) {
	if _cGImageMetadataGetTypeID == nil {
		return 0, symbolCallError("CGImageMetadataGetTypeID", "10.8", _cGImageMetadataGetTypeIDErr)
	}
	return _cGImageMetadataGetTypeID(), nil
}

// CGImageMetadataGetTypeID returns the type identifier for metadata objects.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataGetTypeID()
func CGImageMetadataGetTypeID() uint {
	result, callErr := tryCGImageMetadataGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataRegisterNamespaceForPrefix func(metadata CGMutableImageMetadataRef, xmlns corefoundation.CFStringRef, prefix corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool
var _cGImageMetadataRegisterNamespaceForPrefixErr error

func tryCGImageMetadataRegisterNamespaceForPrefix(metadata CGMutableImageMetadataRef, xmlns corefoundation.CFStringRef, prefix corefoundation.CFStringRef, err *corefoundation.CFErrorRef) (bool, error) {
	if _cGImageMetadataRegisterNamespaceForPrefix == nil {
		return false, symbolCallError("CGImageMetadataRegisterNamespaceForPrefix", "10.8", _cGImageMetadataRegisterNamespaceForPrefixErr)
	}
	return _cGImageMetadataRegisterNamespaceForPrefix(metadata, xmlns, prefix, err), nil
}

// CGImageMetadataRegisterNamespaceForPrefix registers the specified namespace and prefix with the metadata object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataRegisterNamespaceForPrefix(_:_:_:_:)
func CGImageMetadataRegisterNamespaceForPrefix(metadata CGMutableImageMetadataRef, xmlns corefoundation.CFStringRef, prefix corefoundation.CFStringRef, err *corefoundation.CFErrorRef) bool {
	result, callErr := tryCGImageMetadataRegisterNamespaceForPrefix(metadata, xmlns, prefix, err)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataRemoveTagWithPath func(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) bool
var _cGImageMetadataRemoveTagWithPathErr error

func tryCGImageMetadataRemoveTagWithPath(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) (bool, error) {
	if _cGImageMetadataRemoveTagWithPath == nil {
		return false, symbolCallError("CGImageMetadataRemoveTagWithPath", "10.8", _cGImageMetadataRemoveTagWithPathErr)
	}
	return _cGImageMetadataRemoveTagWithPath(metadata, parent, path), nil
}

// CGImageMetadataRemoveTagWithPath removes the tag at the specified path from the metadata object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataRemoveTagWithPath(_:_:_:)
func CGImageMetadataRemoveTagWithPath(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef) bool {
	result, callErr := tryCGImageMetadataRemoveTagWithPath(metadata, parent, path)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataSetTagWithPath func(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef, tag CGImageMetadataTagRef) bool
var _cGImageMetadataSetTagWithPathErr error

func tryCGImageMetadataSetTagWithPath(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef, tag CGImageMetadataTagRef) (bool, error) {
	if _cGImageMetadataSetTagWithPath == nil {
		return false, symbolCallError("CGImageMetadataSetTagWithPath", "10.8", _cGImageMetadataSetTagWithPathErr)
	}
	return _cGImageMetadataSetTagWithPath(metadata, parent, path, tag), nil
}

// CGImageMetadataSetTagWithPath sets the tag at the specified path in the metadata object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetTagWithPath(_:_:_:_:)
func CGImageMetadataSetTagWithPath(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef, tag CGImageMetadataTagRef) bool {
	result, callErr := tryCGImageMetadataSetTagWithPath(metadata, parent, path, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataSetValueMatchingImageProperty func(metadata CGMutableImageMetadataRef, dictionaryName corefoundation.CFStringRef, propertyName corefoundation.CFStringRef, value corefoundation.CFTypeRef) bool
var _cGImageMetadataSetValueMatchingImagePropertyErr error

func tryCGImageMetadataSetValueMatchingImageProperty(metadata CGMutableImageMetadataRef, dictionaryName corefoundation.CFStringRef, propertyName corefoundation.CFStringRef, value corefoundation.CFTypeRef) (bool, error) {
	if _cGImageMetadataSetValueMatchingImageProperty == nil {
		return false, symbolCallError("CGImageMetadataSetValueMatchingImageProperty", "10.8", _cGImageMetadataSetValueMatchingImagePropertyErr)
	}
	return _cGImageMetadataSetValueMatchingImageProperty(metadata, dictionaryName, propertyName, value), nil
}

// CGImageMetadataSetValueMatchingImageProperty updates the value of the metadata tag assigned to the specified image property.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetValueMatchingImageProperty(_:_:_:_:)
func CGImageMetadataSetValueMatchingImageProperty(metadata CGMutableImageMetadataRef, dictionaryName corefoundation.CFStringRef, propertyName corefoundation.CFStringRef, value corefoundation.CFTypeRef) bool {
	result, callErr := tryCGImageMetadataSetValueMatchingImageProperty(metadata, dictionaryName, propertyName, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataSetValueWithPath func(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef, value corefoundation.CFTypeRef) bool
var _cGImageMetadataSetValueWithPathErr error

func tryCGImageMetadataSetValueWithPath(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef, value corefoundation.CFTypeRef) (bool, error) {
	if _cGImageMetadataSetValueWithPath == nil {
		return false, symbolCallError("CGImageMetadataSetValueWithPath", "10.8", _cGImageMetadataSetValueWithPathErr)
	}
	return _cGImageMetadataSetValueWithPath(metadata, parent, path, value), nil
}

// CGImageMetadataSetValueWithPath update the value of an existing metadata tag, or create a new tag using the specified information.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataSetValueWithPath(_:_:_:_:)
func CGImageMetadataSetValueWithPath(metadata CGMutableImageMetadataRef, parent CGImageMetadataTagRef, path corefoundation.CFStringRef, value corefoundation.CFTypeRef) bool {
	result, callErr := tryCGImageMetadataSetValueWithPath(metadata, parent, path, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagCopyName func(tag CGImageMetadataTagRef) corefoundation.CFStringRef
var _cGImageMetadataTagCopyNameErr error

func tryCGImageMetadataTagCopyName(tag CGImageMetadataTagRef) (corefoundation.CFStringRef, error) {
	if _cGImageMetadataTagCopyName == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CGImageMetadataTagCopyName", "10.8", _cGImageMetadataTagCopyNameErr)
	}
	return _cGImageMetadataTagCopyName(tag), nil
}

// CGImageMetadataTagCopyName returns an immutable copy of the tag’s name.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyName(_:)
func CGImageMetadataTagCopyName(tag CGImageMetadataTagRef) corefoundation.CFStringRef {
	result, callErr := tryCGImageMetadataTagCopyName(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagCopyNamespace func(tag CGImageMetadataTagRef) corefoundation.CFStringRef
var _cGImageMetadataTagCopyNamespaceErr error

func tryCGImageMetadataTagCopyNamespace(tag CGImageMetadataTagRef) (corefoundation.CFStringRef, error) {
	if _cGImageMetadataTagCopyNamespace == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CGImageMetadataTagCopyNamespace", "10.8", _cGImageMetadataTagCopyNamespaceErr)
	}
	return _cGImageMetadataTagCopyNamespace(tag), nil
}

// CGImageMetadataTagCopyNamespace returns an immutable copy of the tag’s XMP namespace.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyNamespace(_:)
func CGImageMetadataTagCopyNamespace(tag CGImageMetadataTagRef) corefoundation.CFStringRef {
	result, callErr := tryCGImageMetadataTagCopyNamespace(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagCopyPrefix func(tag CGImageMetadataTagRef) corefoundation.CFStringRef
var _cGImageMetadataTagCopyPrefixErr error

func tryCGImageMetadataTagCopyPrefix(tag CGImageMetadataTagRef) (corefoundation.CFStringRef, error) {
	if _cGImageMetadataTagCopyPrefix == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CGImageMetadataTagCopyPrefix", "10.8", _cGImageMetadataTagCopyPrefixErr)
	}
	return _cGImageMetadataTagCopyPrefix(tag), nil
}

// CGImageMetadataTagCopyPrefix returns an immutable copy of the tag’s prefix.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyPrefix(_:)
func CGImageMetadataTagCopyPrefix(tag CGImageMetadataTagRef) corefoundation.CFStringRef {
	result, callErr := tryCGImageMetadataTagCopyPrefix(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagCopyQualifiers func(tag CGImageMetadataTagRef) corefoundation.CFArrayRef
var _cGImageMetadataTagCopyQualifiersErr error

func tryCGImageMetadataTagCopyQualifiers(tag CGImageMetadataTagRef) (corefoundation.CFArrayRef, error) {
	if _cGImageMetadataTagCopyQualifiers == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CGImageMetadataTagCopyQualifiers", "10.8", _cGImageMetadataTagCopyQualifiersErr)
	}
	return _cGImageMetadataTagCopyQualifiers(tag), nil
}

// CGImageMetadataTagCopyQualifiers returns a shallow copy of the metadata tags that act as qualifiers for the current tag.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyQualifiers(_:)
func CGImageMetadataTagCopyQualifiers(tag CGImageMetadataTagRef) corefoundation.CFArrayRef {
	result, callErr := tryCGImageMetadataTagCopyQualifiers(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagCopyValue func(tag CGImageMetadataTagRef) corefoundation.CFTypeRef
var _cGImageMetadataTagCopyValueErr error

func tryCGImageMetadataTagCopyValue(tag CGImageMetadataTagRef) (corefoundation.CFTypeRef, error) {
	if _cGImageMetadataTagCopyValue == nil {
		return *new(corefoundation.CFTypeRef), symbolCallError("CGImageMetadataTagCopyValue", "10.8", _cGImageMetadataTagCopyValueErr)
	}
	return _cGImageMetadataTagCopyValue(tag), nil
}

// CGImageMetadataTagCopyValue returns a shallow copy of the tag’s value, which is suitable only for reading.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCopyValue(_:)
func CGImageMetadataTagCopyValue(tag CGImageMetadataTagRef) corefoundation.CFTypeRef {
	result, callErr := tryCGImageMetadataTagCopyValue(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagCreate func(xmlns corefoundation.CFStringRef, prefix corefoundation.CFStringRef, name corefoundation.CFStringRef, type_ CGImageMetadataType, value corefoundation.CFTypeRef) CGImageMetadataTagRef
var _cGImageMetadataTagCreateErr error

func tryCGImageMetadataTagCreate(xmlns corefoundation.CFStringRef, prefix corefoundation.CFStringRef, name corefoundation.CFStringRef, type_ CGImageMetadataType, value corefoundation.CFTypeRef) (CGImageMetadataTagRef, error) {
	if _cGImageMetadataTagCreate == nil {
		return *new(CGImageMetadataTagRef), symbolCallError("CGImageMetadataTagCreate", "10.8", _cGImageMetadataTagCreateErr)
	}
	return _cGImageMetadataTagCreate(xmlns, prefix, name, type_, value), nil
}

// CGImageMetadataTagCreate creates a new image metadata tag, and fills it with the specified information.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagCreate(_:_:_:_:_:)
func CGImageMetadataTagCreate(xmlns corefoundation.CFStringRef, prefix corefoundation.CFStringRef, name corefoundation.CFStringRef, type_ CGImageMetadataType, value corefoundation.CFTypeRef) CGImageMetadataTagRef {
	result, callErr := tryCGImageMetadataTagCreate(xmlns, prefix, name, type_, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagGetType func(tag CGImageMetadataTagRef) CGImageMetadataType
var _cGImageMetadataTagGetTypeErr error

func tryCGImageMetadataTagGetType(tag CGImageMetadataTagRef) (CGImageMetadataType, error) {
	if _cGImageMetadataTagGetType == nil {
		return *new(CGImageMetadataType), symbolCallError("CGImageMetadataTagGetType", "10.8", _cGImageMetadataTagGetTypeErr)
	}
	return _cGImageMetadataTagGetType(tag), nil
}

// CGImageMetadataTagGetType returns the type of the metadata tag’s value.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetType(_:)
func CGImageMetadataTagGetType(tag CGImageMetadataTagRef) CGImageMetadataType {
	result, callErr := tryCGImageMetadataTagGetType(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageMetadataTagGetTypeID func() uint
var _cGImageMetadataTagGetTypeIDErr error

func tryCGImageMetadataTagGetTypeID() (uint, error) {
	if _cGImageMetadataTagGetTypeID == nil {
		return 0, symbolCallError("CGImageMetadataTagGetTypeID", "10.8", _cGImageMetadataTagGetTypeIDErr)
	}
	return _cGImageMetadataTagGetTypeID(), nil
}

// CGImageMetadataTagGetTypeID returns the type identifier for the image metadata tag opaque type
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataTagGetTypeID()
func CGImageMetadataTagGetTypeID() uint {
	result, callErr := tryCGImageMetadataTagGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCopyAuxiliaryDataInfoAtIndex func(isrc CGImageSourceRef, index uintptr, auxiliaryImageDataType corefoundation.CFStringRef) corefoundation.CFDictionaryRef
var _cGImageSourceCopyAuxiliaryDataInfoAtIndexErr error

func tryCGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc CGImageSourceRef, index uintptr, auxiliaryImageDataType corefoundation.CFStringRef) (corefoundation.CFDictionaryRef, error) {
	if _cGImageSourceCopyAuxiliaryDataInfoAtIndex == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CGImageSourceCopyAuxiliaryDataInfoAtIndex", "10.13", _cGImageSourceCopyAuxiliaryDataInfoAtIndexErr)
	}
	return _cGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc, index, auxiliaryImageDataType), nil
}

// CGImageSourceCopyAuxiliaryDataInfoAtIndex returns auxiliary data, such as mattes and depth information, that accompany the image.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyAuxiliaryDataInfoAtIndex(_:_:_:)
func CGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc CGImageSourceRef, index uintptr, auxiliaryImageDataType corefoundation.CFStringRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCGImageSourceCopyAuxiliaryDataInfoAtIndex(isrc, index, auxiliaryImageDataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCopyMetadataAtIndex func(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) CGImageMetadataRef
var _cGImageSourceCopyMetadataAtIndexErr error

func tryCGImageSourceCopyMetadataAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) (CGImageMetadataRef, error) {
	if _cGImageSourceCopyMetadataAtIndex == nil {
		return *new(CGImageMetadataRef), symbolCallError("CGImageSourceCopyMetadataAtIndex", "10.8", _cGImageSourceCopyMetadataAtIndexErr)
	}
	return _cGImageSourceCopyMetadataAtIndex(isrc, index, options), nil
}

// CGImageSourceCopyMetadataAtIndex.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyMetadataAtIndex(_:_:_:)
func CGImageSourceCopyMetadataAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) CGImageMetadataRef {
	result, callErr := tryCGImageSourceCopyMetadataAtIndex(isrc, index, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCopyProperties func(isrc CGImageSourceRef, options corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef
var _cGImageSourceCopyPropertiesErr error

func tryCGImageSourceCopyProperties(isrc CGImageSourceRef, options corefoundation.CFDictionaryRef) (corefoundation.CFDictionaryRef, error) {
	if _cGImageSourceCopyProperties == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CGImageSourceCopyProperties", "10.4", _cGImageSourceCopyPropertiesErr)
	}
	return _cGImageSourceCopyProperties(isrc, options), nil
}

// CGImageSourceCopyProperties returns the properties of the image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyProperties(_:_:)
func CGImageSourceCopyProperties(isrc CGImageSourceRef, options corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCGImageSourceCopyProperties(isrc, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCopyPropertiesAtIndex func(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef
var _cGImageSourceCopyPropertiesAtIndexErr error

func tryCGImageSourceCopyPropertiesAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) (corefoundation.CFDictionaryRef, error) {
	if _cGImageSourceCopyPropertiesAtIndex == nil {
		return *new(corefoundation.CFDictionaryRef), symbolCallError("CGImageSourceCopyPropertiesAtIndex", "10.4", _cGImageSourceCopyPropertiesAtIndexErr)
	}
	return _cGImageSourceCopyPropertiesAtIndex(isrc, index, options), nil
}

// CGImageSourceCopyPropertiesAtIndex returns the properties of the image at a specified location in an image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyPropertiesAtIndex(_:_:_:)
func CGImageSourceCopyPropertiesAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCGImageSourceCopyPropertiesAtIndex(isrc, index, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCopyTypeIdentifiers func() corefoundation.CFArrayRef
var _cGImageSourceCopyTypeIdentifiersErr error

func tryCGImageSourceCopyTypeIdentifiers() (corefoundation.CFArrayRef, error) {
	if _cGImageSourceCopyTypeIdentifiers == nil {
		return *new(corefoundation.CFArrayRef), symbolCallError("CGImageSourceCopyTypeIdentifiers", "10.4", _cGImageSourceCopyTypeIdentifiersErr)
	}
	return _cGImageSourceCopyTypeIdentifiers(), nil
}

// CGImageSourceCopyTypeIdentifiers returns an array of uniform type identifiers that are supported for image sources.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCopyTypeIdentifiers()
func CGImageSourceCopyTypeIdentifiers() corefoundation.CFArrayRef {
	result, callErr := tryCGImageSourceCopyTypeIdentifiers()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCreateImageAtIndex func(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) coregraphics.CGImageRef
var _cGImageSourceCreateImageAtIndexErr error

func tryCGImageSourceCreateImageAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) (coregraphics.CGImageRef, error) {
	if _cGImageSourceCreateImageAtIndex == nil {
		return *new(coregraphics.CGImageRef), symbolCallError("CGImageSourceCreateImageAtIndex", "10.4", _cGImageSourceCreateImageAtIndexErr)
	}
	return _cGImageSourceCreateImageAtIndex(isrc, index, options), nil
}

// CGImageSourceCreateImageAtIndex creates an image object from the data at the specified index in an image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateImageAtIndex(_:_:_:)
func CGImageSourceCreateImageAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) coregraphics.CGImageRef {
	result, callErr := tryCGImageSourceCreateImageAtIndex(isrc, index, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCreateIncremental func(options corefoundation.CFDictionaryRef) CGImageSourceRef
var _cGImageSourceCreateIncrementalErr error

func tryCGImageSourceCreateIncremental(options corefoundation.CFDictionaryRef) (CGImageSourceRef, error) {
	if _cGImageSourceCreateIncremental == nil {
		return *new(CGImageSourceRef), symbolCallError("CGImageSourceCreateIncremental", "10.4", _cGImageSourceCreateIncrementalErr)
	}
	return _cGImageSourceCreateIncremental(options), nil
}

// CGImageSourceCreateIncremental creates an empty image source that you can use to accumulate incremental image data.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateIncremental(_:)
func CGImageSourceCreateIncremental(options corefoundation.CFDictionaryRef) CGImageSourceRef {
	result, callErr := tryCGImageSourceCreateIncremental(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCreateThumbnailAtIndex func(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) coregraphics.CGImageRef
var _cGImageSourceCreateThumbnailAtIndexErr error

func tryCGImageSourceCreateThumbnailAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) (coregraphics.CGImageRef, error) {
	if _cGImageSourceCreateThumbnailAtIndex == nil {
		return *new(coregraphics.CGImageRef), symbolCallError("CGImageSourceCreateThumbnailAtIndex", "10.4", _cGImageSourceCreateThumbnailAtIndexErr)
	}
	return _cGImageSourceCreateThumbnailAtIndex(isrc, index, options), nil
}

// CGImageSourceCreateThumbnailAtIndex creates a thumbnail version of the image at the specified index in an image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateThumbnailAtIndex(_:_:_:)
func CGImageSourceCreateThumbnailAtIndex(isrc CGImageSourceRef, index uintptr, options corefoundation.CFDictionaryRef) coregraphics.CGImageRef {
	result, callErr := tryCGImageSourceCreateThumbnailAtIndex(isrc, index, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCreateWithData func(data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef) CGImageSourceRef
var _cGImageSourceCreateWithDataErr error

func tryCGImageSourceCreateWithData(data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef) (CGImageSourceRef, error) {
	if _cGImageSourceCreateWithData == nil {
		return *new(CGImageSourceRef), symbolCallError("CGImageSourceCreateWithData", "10.4", _cGImageSourceCreateWithDataErr)
	}
	return _cGImageSourceCreateWithData(data, options), nil
}

// CGImageSourceCreateWithData creates an image source that reads from a Core Foundation data object.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithData(_:_:)
func CGImageSourceCreateWithData(data corefoundation.CFDataRef, options corefoundation.CFDictionaryRef) CGImageSourceRef {
	result, callErr := tryCGImageSourceCreateWithData(data, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCreateWithDataProvider func(provider coregraphics.CGDataProviderRef, options corefoundation.CFDictionaryRef) CGImageSourceRef
var _cGImageSourceCreateWithDataProviderErr error

func tryCGImageSourceCreateWithDataProvider(provider coregraphics.CGDataProviderRef, options corefoundation.CFDictionaryRef) (CGImageSourceRef, error) {
	if _cGImageSourceCreateWithDataProvider == nil {
		return *new(CGImageSourceRef), symbolCallError("CGImageSourceCreateWithDataProvider", "10.4", _cGImageSourceCreateWithDataProviderErr)
	}
	return _cGImageSourceCreateWithDataProvider(provider, options), nil
}

// CGImageSourceCreateWithDataProvider creates an image source that reads data from the specified data provider.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithDataProvider(_:_:)
func CGImageSourceCreateWithDataProvider(provider coregraphics.CGDataProviderRef, options corefoundation.CFDictionaryRef) CGImageSourceRef {
	result, callErr := tryCGImageSourceCreateWithDataProvider(provider, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceCreateWithURL func(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef) CGImageSourceRef
var _cGImageSourceCreateWithURLErr error

func tryCGImageSourceCreateWithURL(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef) (CGImageSourceRef, error) {
	if _cGImageSourceCreateWithURL == nil {
		return *new(CGImageSourceRef), symbolCallError("CGImageSourceCreateWithURL", "10.4", _cGImageSourceCreateWithURLErr)
	}
	return _cGImageSourceCreateWithURL(url, options), nil
}

// CGImageSourceCreateWithURL creates an image source that reads from a location specified by a URL.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceCreateWithURL(_:_:)
func CGImageSourceCreateWithURL(url corefoundation.CFURLRef, options corefoundation.CFDictionaryRef) CGImageSourceRef {
	result, callErr := tryCGImageSourceCreateWithURL(url, options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceGetCount func(isrc CGImageSourceRef) uintptr
var _cGImageSourceGetCountErr error

func tryCGImageSourceGetCount(isrc CGImageSourceRef) (uintptr, error) {
	if _cGImageSourceGetCount == nil {
		return 0, symbolCallError("CGImageSourceGetCount", "10.4", _cGImageSourceGetCountErr)
	}
	return _cGImageSourceGetCount(isrc), nil
}

// CGImageSourceGetCount returns the number of images (not including thumbnails) in the image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetCount(_:)
func CGImageSourceGetCount(isrc CGImageSourceRef) uintptr {
	result, callErr := tryCGImageSourceGetCount(isrc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceGetPrimaryImageIndex func(isrc CGImageSourceRef) uintptr
var _cGImageSourceGetPrimaryImageIndexErr error

func tryCGImageSourceGetPrimaryImageIndex(isrc CGImageSourceRef) (uintptr, error) {
	if _cGImageSourceGetPrimaryImageIndex == nil {
		return 0, symbolCallError("CGImageSourceGetPrimaryImageIndex", "10.14", _cGImageSourceGetPrimaryImageIndexErr)
	}
	return _cGImageSourceGetPrimaryImageIndex(isrc), nil
}

// CGImageSourceGetPrimaryImageIndex returns the index of the primary image for an High Efficiency Image File Format (HEIF) image.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetPrimaryImageIndex(_:)
func CGImageSourceGetPrimaryImageIndex(isrc CGImageSourceRef) uintptr {
	result, callErr := tryCGImageSourceGetPrimaryImageIndex(isrc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceGetStatus func(isrc CGImageSourceRef) CGImageSourceStatus
var _cGImageSourceGetStatusErr error

func tryCGImageSourceGetStatus(isrc CGImageSourceRef) (CGImageSourceStatus, error) {
	if _cGImageSourceGetStatus == nil {
		return *new(CGImageSourceStatus), symbolCallError("CGImageSourceGetStatus", "10.4", _cGImageSourceGetStatusErr)
	}
	return _cGImageSourceGetStatus(isrc), nil
}

// CGImageSourceGetStatus return the status of an image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatus(_:)
func CGImageSourceGetStatus(isrc CGImageSourceRef) CGImageSourceStatus {
	result, callErr := tryCGImageSourceGetStatus(isrc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceGetStatusAtIndex func(isrc CGImageSourceRef, index uintptr) CGImageSourceStatus
var _cGImageSourceGetStatusAtIndexErr error

func tryCGImageSourceGetStatusAtIndex(isrc CGImageSourceRef, index uintptr) (CGImageSourceStatus, error) {
	if _cGImageSourceGetStatusAtIndex == nil {
		return *new(CGImageSourceStatus), symbolCallError("CGImageSourceGetStatusAtIndex", "10.4", _cGImageSourceGetStatusAtIndexErr)
	}
	return _cGImageSourceGetStatusAtIndex(isrc, index), nil
}

// CGImageSourceGetStatusAtIndex returns the current status of an image at the specified location in the image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetStatusAtIndex(_:_:)
func CGImageSourceGetStatusAtIndex(isrc CGImageSourceRef, index uintptr) CGImageSourceStatus {
	result, callErr := tryCGImageSourceGetStatusAtIndex(isrc, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceGetType func(isrc CGImageSourceRef) corefoundation.CFStringRef
var _cGImageSourceGetTypeErr error

func tryCGImageSourceGetType(isrc CGImageSourceRef) (corefoundation.CFStringRef, error) {
	if _cGImageSourceGetType == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("CGImageSourceGetType", "10.4", _cGImageSourceGetTypeErr)
	}
	return _cGImageSourceGetType(isrc), nil
}

// CGImageSourceGetType returns the uniform type identifier of the source container.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetType(_:)
func CGImageSourceGetType(isrc CGImageSourceRef) corefoundation.CFStringRef {
	result, callErr := tryCGImageSourceGetType(isrc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceGetTypeID func() uint
var _cGImageSourceGetTypeIDErr error

func tryCGImageSourceGetTypeID() (uint, error) {
	if _cGImageSourceGetTypeID == nil {
		return 0, symbolCallError("CGImageSourceGetTypeID", "10.4", _cGImageSourceGetTypeIDErr)
	}
	return _cGImageSourceGetTypeID(), nil
}

// CGImageSourceGetTypeID returns the unique type identifier of an image source opaque type.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceGetTypeID()
func CGImageSourceGetTypeID() uint {
	result, callErr := tryCGImageSourceGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceRemoveCacheAtIndex func(isrc CGImageSourceRef, index uintptr)
var _cGImageSourceRemoveCacheAtIndexErr error

func tryCGImageSourceRemoveCacheAtIndex(isrc CGImageSourceRef, index uintptr) error {
	if _cGImageSourceRemoveCacheAtIndex == nil {
		return symbolCallError("CGImageSourceRemoveCacheAtIndex", "10.9", _cGImageSourceRemoveCacheAtIndexErr)
	}
	_cGImageSourceRemoveCacheAtIndex(isrc, index)
	return nil
}

// CGImageSourceRemoveCacheAtIndex.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceRemoveCacheAtIndex(_:_:)
func CGImageSourceRemoveCacheAtIndex(isrc CGImageSourceRef, index uintptr) {
	if callErr := tryCGImageSourceRemoveCacheAtIndex(isrc, index); callErr != nil {
		panic(callErr)
	}
}

var _cGImageSourceSetAllowableTypes func(allowableTypes corefoundation.CFArrayRef) int32
var _cGImageSourceSetAllowableTypesErr error

func tryCGImageSourceSetAllowableTypes(allowableTypes corefoundation.CFArrayRef) (int32, error) {
	if _cGImageSourceSetAllowableTypes == nil {
		return 0, symbolCallError("CGImageSourceSetAllowableTypes", "14.2", _cGImageSourceSetAllowableTypesErr)
	}
	return _cGImageSourceSetAllowableTypes(allowableTypes), nil
}

// CGImageSourceSetAllowableTypes restricts which image formats can be decoded in the current process.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceSetAllowableTypes(_:)
func CGImageSourceSetAllowableTypes(allowableTypes corefoundation.CFArrayRef) int32 {
	result, callErr := tryCGImageSourceSetAllowableTypes(allowableTypes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cGImageSourceUpdateData func(isrc CGImageSourceRef, data corefoundation.CFDataRef, final bool)
var _cGImageSourceUpdateDataErr error

func tryCGImageSourceUpdateData(isrc CGImageSourceRef, data corefoundation.CFDataRef, final bool) error {
	if _cGImageSourceUpdateData == nil {
		return symbolCallError("CGImageSourceUpdateData", "10.4", _cGImageSourceUpdateDataErr)
	}
	_cGImageSourceUpdateData(isrc, data, final)
	return nil
}

// CGImageSourceUpdateData updates the data in an incremental image source.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateData(_:_:_:)
func CGImageSourceUpdateData(isrc CGImageSourceRef, data corefoundation.CFDataRef, final bool) {
	if callErr := tryCGImageSourceUpdateData(isrc, data, final); callErr != nil {
		panic(callErr)
	}
}

var _cGImageSourceUpdateDataProvider func(isrc CGImageSourceRef, provider coregraphics.CGDataProviderRef, final bool)
var _cGImageSourceUpdateDataProviderErr error

func tryCGImageSourceUpdateDataProvider(isrc CGImageSourceRef, provider coregraphics.CGDataProviderRef, final bool) error {
	if _cGImageSourceUpdateDataProvider == nil {
		return symbolCallError("CGImageSourceUpdateDataProvider", "10.4", _cGImageSourceUpdateDataProviderErr)
	}
	_cGImageSourceUpdateDataProvider(isrc, provider, final)
	return nil
}

// CGImageSourceUpdateDataProvider updates an incremental image source with a new data provider.
//
// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceUpdateDataProvider(_:_:_:)
func CGImageSourceUpdateDataProvider(isrc CGImageSourceRef, provider coregraphics.CGDataProviderRef, final bool) {
	if callErr := tryCGImageSourceUpdateDataProvider(isrc, provider, final); callErr != nil {
		panic(callErr)
	}
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cGAnimateImageAtURLWithBlock, &_cGAnimateImageAtURLWithBlockErr, frameworkHandle, "CGAnimateImageAtURLWithBlock", "10.15")
	registerFunc(&_cGAnimateImageDataWithBlock, &_cGAnimateImageDataWithBlockErr, frameworkHandle, "CGAnimateImageDataWithBlock", "10.15")
	registerFunc(&_cGImageDestinationAddAuxiliaryDataInfo, &_cGImageDestinationAddAuxiliaryDataInfoErr, frameworkHandle, "CGImageDestinationAddAuxiliaryDataInfo", "10.13")
	registerFunc(&_cGImageDestinationAddImage, &_cGImageDestinationAddImageErr, frameworkHandle, "CGImageDestinationAddImage", "10.4")
	registerFunc(&_cGImageDestinationAddImageAndMetadata, &_cGImageDestinationAddImageAndMetadataErr, frameworkHandle, "CGImageDestinationAddImageAndMetadata", "10.8")
	registerFunc(&_cGImageDestinationAddImageFromSource, &_cGImageDestinationAddImageFromSourceErr, frameworkHandle, "CGImageDestinationAddImageFromSource", "10.4")
	registerFunc(&_cGImageDestinationCopyImageSource, &_cGImageDestinationCopyImageSourceErr, frameworkHandle, "CGImageDestinationCopyImageSource", "10.8")
	registerFunc(&_cGImageDestinationCopyTypeIdentifiers, &_cGImageDestinationCopyTypeIdentifiersErr, frameworkHandle, "CGImageDestinationCopyTypeIdentifiers", "10.4")
	registerFunc(&_cGImageDestinationCreateWithData, &_cGImageDestinationCreateWithDataErr, frameworkHandle, "CGImageDestinationCreateWithData", "10.4")
	registerFunc(&_cGImageDestinationCreateWithDataConsumer, &_cGImageDestinationCreateWithDataConsumerErr, frameworkHandle, "CGImageDestinationCreateWithDataConsumer", "10.4")
	registerFunc(&_cGImageDestinationCreateWithURL, &_cGImageDestinationCreateWithURLErr, frameworkHandle, "CGImageDestinationCreateWithURL", "10.4")
	registerFunc(&_cGImageDestinationFinalize, &_cGImageDestinationFinalizeErr, frameworkHandle, "CGImageDestinationFinalize", "10.4")
	registerFunc(&_cGImageDestinationGetTypeID, &_cGImageDestinationGetTypeIDErr, frameworkHandle, "CGImageDestinationGetTypeID", "10.4")
	registerFunc(&_cGImageDestinationSetProperties, &_cGImageDestinationSetPropertiesErr, frameworkHandle, "CGImageDestinationSetProperties", "10.4")
	registerFunc(&_cGImageMetadataCopyStringValueWithPath, &_cGImageMetadataCopyStringValueWithPathErr, frameworkHandle, "CGImageMetadataCopyStringValueWithPath", "10.8")
	registerFunc(&_cGImageMetadataCopyTagMatchingImageProperty, &_cGImageMetadataCopyTagMatchingImagePropertyErr, frameworkHandle, "CGImageMetadataCopyTagMatchingImageProperty", "10.8")
	registerFunc(&_cGImageMetadataCopyTagWithPath, &_cGImageMetadataCopyTagWithPathErr, frameworkHandle, "CGImageMetadataCopyTagWithPath", "10.8")
	registerFunc(&_cGImageMetadataCopyTags, &_cGImageMetadataCopyTagsErr, frameworkHandle, "CGImageMetadataCopyTags", "10.8")
	registerFunc(&_cGImageMetadataCreateFromXMPData, &_cGImageMetadataCreateFromXMPDataErr, frameworkHandle, "CGImageMetadataCreateFromXMPData", "10.8")
	registerFunc(&_cGImageMetadataCreateMutable, &_cGImageMetadataCreateMutableErr, frameworkHandle, "CGImageMetadataCreateMutable", "10.8")
	registerFunc(&_cGImageMetadataCreateMutableCopy, &_cGImageMetadataCreateMutableCopyErr, frameworkHandle, "CGImageMetadataCreateMutableCopy", "10.8")
	registerFunc(&_cGImageMetadataCreateXMPData, &_cGImageMetadataCreateXMPDataErr, frameworkHandle, "CGImageMetadataCreateXMPData", "10.8")
	registerFunc(&_cGImageMetadataEnumerateTagsUsingBlock, &_cGImageMetadataEnumerateTagsUsingBlockErr, frameworkHandle, "CGImageMetadataEnumerateTagsUsingBlock", "10.8")
	registerFunc(&_cGImageMetadataGetTypeID, &_cGImageMetadataGetTypeIDErr, frameworkHandle, "CGImageMetadataGetTypeID", "10.8")
	registerFunc(&_cGImageMetadataRegisterNamespaceForPrefix, &_cGImageMetadataRegisterNamespaceForPrefixErr, frameworkHandle, "CGImageMetadataRegisterNamespaceForPrefix", "10.8")
	registerFunc(&_cGImageMetadataRemoveTagWithPath, &_cGImageMetadataRemoveTagWithPathErr, frameworkHandle, "CGImageMetadataRemoveTagWithPath", "10.8")
	registerFunc(&_cGImageMetadataSetTagWithPath, &_cGImageMetadataSetTagWithPathErr, frameworkHandle, "CGImageMetadataSetTagWithPath", "10.8")
	registerFunc(&_cGImageMetadataSetValueMatchingImageProperty, &_cGImageMetadataSetValueMatchingImagePropertyErr, frameworkHandle, "CGImageMetadataSetValueMatchingImageProperty", "10.8")
	registerFunc(&_cGImageMetadataSetValueWithPath, &_cGImageMetadataSetValueWithPathErr, frameworkHandle, "CGImageMetadataSetValueWithPath", "10.8")
	registerFunc(&_cGImageMetadataTagCopyName, &_cGImageMetadataTagCopyNameErr, frameworkHandle, "CGImageMetadataTagCopyName", "10.8")
	registerFunc(&_cGImageMetadataTagCopyNamespace, &_cGImageMetadataTagCopyNamespaceErr, frameworkHandle, "CGImageMetadataTagCopyNamespace", "10.8")
	registerFunc(&_cGImageMetadataTagCopyPrefix, &_cGImageMetadataTagCopyPrefixErr, frameworkHandle, "CGImageMetadataTagCopyPrefix", "10.8")
	registerFunc(&_cGImageMetadataTagCopyQualifiers, &_cGImageMetadataTagCopyQualifiersErr, frameworkHandle, "CGImageMetadataTagCopyQualifiers", "10.8")
	registerFunc(&_cGImageMetadataTagCopyValue, &_cGImageMetadataTagCopyValueErr, frameworkHandle, "CGImageMetadataTagCopyValue", "10.8")
	registerFunc(&_cGImageMetadataTagCreate, &_cGImageMetadataTagCreateErr, frameworkHandle, "CGImageMetadataTagCreate", "10.8")
	registerFunc(&_cGImageMetadataTagGetType, &_cGImageMetadataTagGetTypeErr, frameworkHandle, "CGImageMetadataTagGetType", "10.8")
	registerFunc(&_cGImageMetadataTagGetTypeID, &_cGImageMetadataTagGetTypeIDErr, frameworkHandle, "CGImageMetadataTagGetTypeID", "10.8")
	registerFunc(&_cGImageSourceCopyAuxiliaryDataInfoAtIndex, &_cGImageSourceCopyAuxiliaryDataInfoAtIndexErr, frameworkHandle, "CGImageSourceCopyAuxiliaryDataInfoAtIndex", "10.13")
	registerFunc(&_cGImageSourceCopyMetadataAtIndex, &_cGImageSourceCopyMetadataAtIndexErr, frameworkHandle, "CGImageSourceCopyMetadataAtIndex", "10.8")
	registerFunc(&_cGImageSourceCopyProperties, &_cGImageSourceCopyPropertiesErr, frameworkHandle, "CGImageSourceCopyProperties", "10.4")
	registerFunc(&_cGImageSourceCopyPropertiesAtIndex, &_cGImageSourceCopyPropertiesAtIndexErr, frameworkHandle, "CGImageSourceCopyPropertiesAtIndex", "10.4")
	registerFunc(&_cGImageSourceCopyTypeIdentifiers, &_cGImageSourceCopyTypeIdentifiersErr, frameworkHandle, "CGImageSourceCopyTypeIdentifiers", "10.4")
	registerFunc(&_cGImageSourceCreateImageAtIndex, &_cGImageSourceCreateImageAtIndexErr, frameworkHandle, "CGImageSourceCreateImageAtIndex", "10.4")
	registerFunc(&_cGImageSourceCreateIncremental, &_cGImageSourceCreateIncrementalErr, frameworkHandle, "CGImageSourceCreateIncremental", "10.4")
	registerFunc(&_cGImageSourceCreateThumbnailAtIndex, &_cGImageSourceCreateThumbnailAtIndexErr, frameworkHandle, "CGImageSourceCreateThumbnailAtIndex", "10.4")
	registerFunc(&_cGImageSourceCreateWithData, &_cGImageSourceCreateWithDataErr, frameworkHandle, "CGImageSourceCreateWithData", "10.4")
	registerFunc(&_cGImageSourceCreateWithDataProvider, &_cGImageSourceCreateWithDataProviderErr, frameworkHandle, "CGImageSourceCreateWithDataProvider", "10.4")
	registerFunc(&_cGImageSourceCreateWithURL, &_cGImageSourceCreateWithURLErr, frameworkHandle, "CGImageSourceCreateWithURL", "10.4")
	registerFunc(&_cGImageSourceGetCount, &_cGImageSourceGetCountErr, frameworkHandle, "CGImageSourceGetCount", "10.4")
	registerFunc(&_cGImageSourceGetPrimaryImageIndex, &_cGImageSourceGetPrimaryImageIndexErr, frameworkHandle, "CGImageSourceGetPrimaryImageIndex", "10.14")
	registerFunc(&_cGImageSourceGetStatus, &_cGImageSourceGetStatusErr, frameworkHandle, "CGImageSourceGetStatus", "10.4")
	registerFunc(&_cGImageSourceGetStatusAtIndex, &_cGImageSourceGetStatusAtIndexErr, frameworkHandle, "CGImageSourceGetStatusAtIndex", "10.4")
	registerFunc(&_cGImageSourceGetType, &_cGImageSourceGetTypeErr, frameworkHandle, "CGImageSourceGetType", "10.4")
	registerFunc(&_cGImageSourceGetTypeID, &_cGImageSourceGetTypeIDErr, frameworkHandle, "CGImageSourceGetTypeID", "10.4")
	registerFunc(&_cGImageSourceRemoveCacheAtIndex, &_cGImageSourceRemoveCacheAtIndexErr, frameworkHandle, "CGImageSourceRemoveCacheAtIndex", "10.9")
	registerFunc(&_cGImageSourceSetAllowableTypes, &_cGImageSourceSetAllowableTypesErr, frameworkHandle, "CGImageSourceSetAllowableTypes", "14.2")
	registerFunc(&_cGImageSourceUpdateData, &_cGImageSourceUpdateDataErr, frameworkHandle, "CGImageSourceUpdateData", "10.4")
	registerFunc(&_cGImageSourceUpdateDataProvider, &_cGImageSourceUpdateDataProviderErr, frameworkHandle, "CGImageSourceUpdateDataProvider", "10.4")
}
