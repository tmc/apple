// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
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
		return fmt.Sprintf("Vision: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("Vision: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("Vision: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("Vision: register symbol %s: %v", name, r)
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

var _vNElementTypeSize func(elementType VNElementType) uint
var _vNElementTypeSizeErr error

func tryVNElementTypeSize(elementType VNElementType) (uint, error) {
	if _vNElementTypeSize == nil {
		return 0, symbolCallError("VNElementTypeSize", "10.15", _vNElementTypeSizeErr)
	}
	return _vNElementTypeSize(elementType), nil
}

// VNElementTypeSize returns the size of a feature print element.
//
// See: https://developer.apple.com/documentation/Vision/VNElementTypeSize(_:)
func VNElementTypeSize(elementType VNElementType) uint {
	result, callErr := tryVNElementTypeSize(elementType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNImagePointForFaceLandmarkPoint func(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint
var _vNImagePointForFaceLandmarkPointErr error

func tryVNImagePointForFaceLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) (corefoundation.CGPoint, error) {
	if _vNImagePointForFaceLandmarkPoint == nil {
		return corefoundation.CGPoint{}, symbolCallError("VNImagePointForFaceLandmarkPoint", "10.13", _vNImagePointForFaceLandmarkPointErr)
	}
	return _vNImagePointForFaceLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight), nil
}

// VNImagePointForFaceLandmarkPoint returns the image coordinates of a specified face landmark point.
//
// See: https://developer.apple.com/documentation/Vision/VNImagePointForFaceLandmarkPoint(_:_:_:_:)
func VNImagePointForFaceLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	result, callErr := tryVNImagePointForFaceLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNImagePointForNormalizedPoint func(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint
var _vNImagePointForNormalizedPointErr error

func tryVNImagePointForNormalizedPoint(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) (corefoundation.CGPoint, error) {
	if _vNImagePointForNormalizedPoint == nil {
		return corefoundation.CGPoint{}, symbolCallError("VNImagePointForNormalizedPoint", "10.13", _vNImagePointForNormalizedPointErr)
	}
	return _vNImagePointForNormalizedPoint(normalizedPoint, imageWidth, imageHeight), nil
}

// VNImagePointForNormalizedPoint projects a point in normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPoint(_:_:_:)
func VNImagePointForNormalizedPoint(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	result, callErr := tryVNImagePointForNormalizedPoint(normalizedPoint, imageWidth, imageHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNImagePointForNormalizedPointUsingRegionOfInterest func(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint
var _vNImagePointForNormalizedPointUsingRegionOfInterestErr error

func tryVNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) (corefoundation.CGPoint, error) {
	if _vNImagePointForNormalizedPointUsingRegionOfInterest == nil {
		return corefoundation.CGPoint{}, symbolCallError("VNImagePointForNormalizedPointUsingRegionOfInterest", "12.0", _vNImagePointForNormalizedPointUsingRegionOfInterestErr)
	}
	return _vNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint, imageWidth, imageHeight, roi), nil
}

// VNImagePointForNormalizedPointUsingRegionOfInterest projects a point from a region of interest within the normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPointUsingRegionOfInterest(_:_:_:_:)
func VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint {
	result, callErr := tryVNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint, imageWidth, imageHeight, roi)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNImageRectForNormalizedRect func(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect
var _vNImageRectForNormalizedRectErr error

func tryVNImageRectForNormalizedRect(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) (corefoundation.CGRect, error) {
	if _vNImageRectForNormalizedRect == nil {
		return corefoundation.CGRect{}, symbolCallError("VNImageRectForNormalizedRect", "10.13", _vNImageRectForNormalizedRectErr)
	}
	return _vNImageRectForNormalizedRect(normalizedRect, imageWidth, imageHeight), nil
}

// VNImageRectForNormalizedRect projects a rectangle from normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRect(_:_:_:)
func VNImageRectForNormalizedRect(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect {
	result, callErr := tryVNImageRectForNormalizedRect(normalizedRect, imageWidth, imageHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNImageRectForNormalizedRectUsingRegionOfInterest func(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect
var _vNImageRectForNormalizedRectUsingRegionOfInterestErr error

func tryVNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _vNImageRectForNormalizedRectUsingRegionOfInterest == nil {
		return corefoundation.CGRect{}, symbolCallError("VNImageRectForNormalizedRectUsingRegionOfInterest", "12.0", _vNImageRectForNormalizedRectUsingRegionOfInterestErr)
	}
	return _vNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect, imageWidth, imageHeight, roi), nil
}

// VNImageRectForNormalizedRectUsingRegionOfInterest projects a rectangle from a region of interest within the normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRectUsingRegionOfInterest(_:_:_:_:)
func VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryVNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect, imageWidth, imageHeight, roi)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNNormalizedFaceBoundingBoxPointForLandmarkPoint func(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint
var _vNNormalizedFaceBoundingBoxPointForLandmarkPointErr error

func tryVNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) (corefoundation.CGPoint, error) {
	if _vNNormalizedFaceBoundingBoxPointForLandmarkPoint == nil {
		return corefoundation.CGPoint{}, symbolCallError("VNNormalizedFaceBoundingBoxPointForLandmarkPoint", "10.13", _vNNormalizedFaceBoundingBoxPointForLandmarkPointErr)
	}
	return _vNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight), nil
}

// VNNormalizedFaceBoundingBoxPointForLandmarkPoint returns the coordinates of a specified face landmark point, in bounding box coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedFaceBoundingBoxPointForLandmarkPoint(_:_:_:_:)
func VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint unsafe.Pointer, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	result, callErr := tryVNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNNormalizedPointForImagePoint func(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint
var _vNNormalizedPointForImagePointErr error

func tryVNNormalizedPointForImagePoint(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) (corefoundation.CGPoint, error) {
	if _vNNormalizedPointForImagePoint == nil {
		return corefoundation.CGPoint{}, symbolCallError("VNNormalizedPointForImagePoint", "11.0", _vNNormalizedPointForImagePointErr)
	}
	return _vNNormalizedPointForImagePoint(imagePoint, imageWidth, imageHeight), nil
}

// VNNormalizedPointForImagePoint projects a point from image coordinates into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePoint(_:_:_:)
func VNNormalizedPointForImagePoint(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	result, callErr := tryVNNormalizedPointForImagePoint(imagePoint, imageWidth, imageHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNNormalizedPointForImagePointUsingRegionOfInterest func(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint
var _vNNormalizedPointForImagePointUsingRegionOfInterestErr error

func tryVNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) (corefoundation.CGPoint, error) {
	if _vNNormalizedPointForImagePointUsingRegionOfInterest == nil {
		return corefoundation.CGPoint{}, symbolCallError("VNNormalizedPointForImagePointUsingRegionOfInterest", "12.0", _vNNormalizedPointForImagePointUsingRegionOfInterestErr)
	}
	return _vNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint, imageWidth, imageHeight, roi), nil
}

// VNNormalizedPointForImagePointUsingRegionOfInterest projects a point from a region of interest within the image coordinates into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePointUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint {
	result, callErr := tryVNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint, imageWidth, imageHeight, roi)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNNormalizedRectForImageRect func(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect
var _vNNormalizedRectForImageRectErr error

func tryVNNormalizedRectForImageRect(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) (corefoundation.CGRect, error) {
	if _vNNormalizedRectForImageRect == nil {
		return corefoundation.CGRect{}, symbolCallError("VNNormalizedRectForImageRect", "10.13", _vNNormalizedRectForImageRectErr)
	}
	return _vNNormalizedRectForImageRect(imageRect, imageWidth, imageHeight), nil
}

// VNNormalizedRectForImageRect projects a rectangle from image coordinates into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRect(_:_:_:)
func VNNormalizedRectForImageRect(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect {
	result, callErr := tryVNNormalizedRectForImageRect(imageRect, imageWidth, imageHeight)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNNormalizedRectForImageRectUsingRegionOfInterest func(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect
var _vNNormalizedRectForImageRectUsingRegionOfInterestErr error

func tryVNNormalizedRectForImageRectUsingRegionOfInterest(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _vNNormalizedRectForImageRectUsingRegionOfInterest == nil {
		return corefoundation.CGRect{}, symbolCallError("VNNormalizedRectForImageRectUsingRegionOfInterest", "12.0", _vNNormalizedRectForImageRectUsingRegionOfInterestErr)
	}
	return _vNNormalizedRectForImageRectUsingRegionOfInterest(imageRect, imageWidth, imageHeight, roi), nil
}

// VNNormalizedRectForImageRectUsingRegionOfInterest projects a rectangle from a region of interest within the image coordinates space into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRectUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryVNNormalizedRectForImageRectUsingRegionOfInterest(imageRect, imageWidth, imageHeight, roi)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vNNormalizedRectIsIdentityRect func(normalizedRect corefoundation.CGRect) bool
var _vNNormalizedRectIsIdentityRectErr error

func tryVNNormalizedRectIsIdentityRect(normalizedRect corefoundation.CGRect) (bool, error) {
	if _vNNormalizedRectIsIdentityRect == nil {
		return false, symbolCallError("VNNormalizedRectIsIdentityRect", "10.13", _vNNormalizedRectIsIdentityRectErr)
	}
	return _vNNormalizedRectIsIdentityRect(normalizedRect), nil
}

// VNNormalizedRectIsIdentityRect returns a Boolean value that indicates whether the rectangle has an origin of zero and unit length and width.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedRectIsIdentityRect(_:)
func VNNormalizedRectIsIdentityRect(normalizedRect corefoundation.CGRect) bool {
	result, callErr := tryVNNormalizedRectIsIdentityRect(normalizedRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_vNElementTypeSize, &_vNElementTypeSizeErr, frameworkHandle, "VNElementTypeSize", "10.15")
	registerFunc(&_vNImagePointForFaceLandmarkPoint, &_vNImagePointForFaceLandmarkPointErr, frameworkHandle, "VNImagePointForFaceLandmarkPoint", "10.13")
	registerFunc(&_vNImagePointForNormalizedPoint, &_vNImagePointForNormalizedPointErr, frameworkHandle, "VNImagePointForNormalizedPoint", "10.13")
	registerFunc(&_vNImagePointForNormalizedPointUsingRegionOfInterest, &_vNImagePointForNormalizedPointUsingRegionOfInterestErr, frameworkHandle, "VNImagePointForNormalizedPointUsingRegionOfInterest", "12.0")
	registerFunc(&_vNImageRectForNormalizedRect, &_vNImageRectForNormalizedRectErr, frameworkHandle, "VNImageRectForNormalizedRect", "10.13")
	registerFunc(&_vNImageRectForNormalizedRectUsingRegionOfInterest, &_vNImageRectForNormalizedRectUsingRegionOfInterestErr, frameworkHandle, "VNImageRectForNormalizedRectUsingRegionOfInterest", "12.0")
	registerFunc(&_vNNormalizedFaceBoundingBoxPointForLandmarkPoint, &_vNNormalizedFaceBoundingBoxPointForLandmarkPointErr, frameworkHandle, "VNNormalizedFaceBoundingBoxPointForLandmarkPoint", "10.13")
	registerFunc(&_vNNormalizedPointForImagePoint, &_vNNormalizedPointForImagePointErr, frameworkHandle, "VNNormalizedPointForImagePoint", "11.0")
	registerFunc(&_vNNormalizedPointForImagePointUsingRegionOfInterest, &_vNNormalizedPointForImagePointUsingRegionOfInterestErr, frameworkHandle, "VNNormalizedPointForImagePointUsingRegionOfInterest", "12.0")
	registerFunc(&_vNNormalizedRectForImageRect, &_vNNormalizedRectForImageRectErr, frameworkHandle, "VNNormalizedRectForImageRect", "10.13")
	registerFunc(&_vNNormalizedRectForImageRectUsingRegionOfInterest, &_vNNormalizedRectForImageRectUsingRegionOfInterestErr, frameworkHandle, "VNNormalizedRectForImageRectUsingRegionOfInterest", "12.0")
	registerFunc(&_vNNormalizedRectIsIdentityRect, &_vNNormalizedRectIsIdentityRectErr, frameworkHandle, "VNNormalizedRectIsIdentityRect", "10.13")
}
