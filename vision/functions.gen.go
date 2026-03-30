// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Vision: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: Vision: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Vision: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _vNElementTypeSize func(elementType VNElementType) uint

// VNElementTypeSize returns the size of a feature print element.
//
// See: https://developer.apple.com/documentation/Vision/VNElementTypeSize(_:)
func VNElementTypeSize(elementType VNElementType) uint {
	if _vNElementTypeSize == nil {
		panic("Vision: symbol VNElementTypeSize not loaded")
	}
	return _vNElementTypeSize(elementType)
}

var _vNImagePointForFaceLandmarkPoint func(faceLandmarkPoint uintptr, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint

// VNImagePointForFaceLandmarkPoint returns the image coordinates of a specified face landmark point.
//
// See: https://developer.apple.com/documentation/Vision/VNImagePointForFaceLandmarkPoint(_:_:_:_:)
func VNImagePointForFaceLandmarkPoint(faceLandmarkPoint uintptr, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	if _vNImagePointForFaceLandmarkPoint == nil {
		panic("Vision: symbol VNImagePointForFaceLandmarkPoint not loaded")
	}
	return _vNImagePointForFaceLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
}

var _vNImagePointForNormalizedPoint func(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint

// VNImagePointForNormalizedPoint projects a point in normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPoint(_:_:_:)
func VNImagePointForNormalizedPoint(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	if _vNImagePointForNormalizedPoint == nil {
		panic("Vision: symbol VNImagePointForNormalizedPoint not loaded")
	}
	return _vNImagePointForNormalizedPoint(normalizedPoint, imageWidth, imageHeight)
}

var _vNImagePointForNormalizedPointUsingRegionOfInterest func(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint

// VNImagePointForNormalizedPointUsingRegionOfInterest projects a point from a region of interest within the normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImagePointForNormalizedPointUsingRegionOfInterest(_:_:_:_:)
func VNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint {
	if _vNImagePointForNormalizedPointUsingRegionOfInterest == nil {
		panic("Vision: symbol VNImagePointForNormalizedPointUsingRegionOfInterest not loaded")
	}
	return _vNImagePointForNormalizedPointUsingRegionOfInterest(normalizedPoint, imageWidth, imageHeight, roi)
}

var _vNImageRectForNormalizedRect func(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect

// VNImageRectForNormalizedRect projects a rectangle from normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRect(_:_:_:)
func VNImageRectForNormalizedRect(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect {
	if _vNImageRectForNormalizedRect == nil {
		panic("Vision: symbol VNImageRectForNormalizedRect not loaded")
	}
	return _vNImageRectForNormalizedRect(normalizedRect, imageWidth, imageHeight)
}

var _vNImageRectForNormalizedRectUsingRegionOfInterest func(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect

// VNImageRectForNormalizedRectUsingRegionOfInterest projects a rectangle from a region of interest within the normalized coordinates into image coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNImageRectForNormalizedRectUsingRegionOfInterest(_:_:_:_:)
func VNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect {
	if _vNImageRectForNormalizedRectUsingRegionOfInterest == nil {
		panic("Vision: symbol VNImageRectForNormalizedRectUsingRegionOfInterest not loaded")
	}
	return _vNImageRectForNormalizedRectUsingRegionOfInterest(normalizedRect, imageWidth, imageHeight, roi)
}

var _vNNormalizedFaceBoundingBoxPointForLandmarkPoint func(faceLandmarkPoint uintptr, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint

// VNNormalizedFaceBoundingBoxPointForLandmarkPoint returns the coordinates of a specified face landmark point, in bounding box coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedFaceBoundingBoxPointForLandmarkPoint(_:_:_:_:)
func VNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint uintptr, faceBoundingBox corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	if _vNNormalizedFaceBoundingBoxPointForLandmarkPoint == nil {
		panic("Vision: symbol VNNormalizedFaceBoundingBoxPointForLandmarkPoint not loaded")
	}
	return _vNNormalizedFaceBoundingBoxPointForLandmarkPoint(faceLandmarkPoint, faceBoundingBox, imageWidth, imageHeight)
}

var _vNNormalizedPointForImagePoint func(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint

// VNNormalizedPointForImagePoint projects a point from image coordinates into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePoint(_:_:_:)
func VNNormalizedPointForImagePoint(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr) corefoundation.CGPoint {
	if _vNNormalizedPointForImagePoint == nil {
		panic("Vision: symbol VNNormalizedPointForImagePoint not loaded")
	}
	return _vNNormalizedPointForImagePoint(imagePoint, imageWidth, imageHeight)
}

var _vNNormalizedPointForImagePointUsingRegionOfInterest func(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint

// VNNormalizedPointForImagePointUsingRegionOfInterest projects a point from a region of interest within the image coordinates into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedPointForImagePointUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint corefoundation.CGPoint, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGPoint {
	if _vNNormalizedPointForImagePointUsingRegionOfInterest == nil {
		panic("Vision: symbol VNNormalizedPointForImagePointUsingRegionOfInterest not loaded")
	}
	return _vNNormalizedPointForImagePointUsingRegionOfInterest(imagePoint, imageWidth, imageHeight, roi)
}

var _vNNormalizedRectForImageRect func(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect

// VNNormalizedRectForImageRect projects a rectangle from image coordinates into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRect(_:_:_:)
func VNNormalizedRectForImageRect(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr) corefoundation.CGRect {
	if _vNNormalizedRectForImageRect == nil {
		panic("Vision: symbol VNNormalizedRectForImageRect not loaded")
	}
	return _vNNormalizedRectForImageRect(imageRect, imageWidth, imageHeight)
}

var _vNNormalizedRectForImageRectUsingRegionOfInterest func(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect

// VNNormalizedRectForImageRectUsingRegionOfInterest projects a rectangle from a region of interest within the image coordinates space into normalized coordinates.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedRectForImageRectUsingRegionOfInterest(_:_:_:_:)
func VNNormalizedRectForImageRectUsingRegionOfInterest(imageRect corefoundation.CGRect, imageWidth uintptr, imageHeight uintptr, roi corefoundation.CGRect) corefoundation.CGRect {
	if _vNNormalizedRectForImageRectUsingRegionOfInterest == nil {
		panic("Vision: symbol VNNormalizedRectForImageRectUsingRegionOfInterest not loaded")
	}
	return _vNNormalizedRectForImageRectUsingRegionOfInterest(imageRect, imageWidth, imageHeight, roi)
}

var _vNNormalizedRectIsIdentityRect func(normalizedRect corefoundation.CGRect) bool

// VNNormalizedRectIsIdentityRect returns a Boolean value that indicates whether the rectangle has an origin of zero and unit length and width.
//
// See: https://developer.apple.com/documentation/Vision/VNNormalizedRectIsIdentityRect(_:)
func VNNormalizedRectIsIdentityRect(normalizedRect corefoundation.CGRect) bool {
	if _vNNormalizedRectIsIdentityRect == nil {
		panic("Vision: symbol VNNormalizedRectIsIdentityRect not loaded")
	}
	return _vNNormalizedRectIsIdentityRect(normalizedRect)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_vNElementTypeSize, frameworkHandle, "VNElementTypeSize")
	registerFunc(&_vNImagePointForFaceLandmarkPoint, frameworkHandle, "VNImagePointForFaceLandmarkPoint")
	registerFunc(&_vNImagePointForNormalizedPoint, frameworkHandle, "VNImagePointForNormalizedPoint")
	registerFunc(&_vNImagePointForNormalizedPointUsingRegionOfInterest, frameworkHandle, "VNImagePointForNormalizedPointUsingRegionOfInterest")
	registerFunc(&_vNImageRectForNormalizedRect, frameworkHandle, "VNImageRectForNormalizedRect")
	registerFunc(&_vNImageRectForNormalizedRectUsingRegionOfInterest, frameworkHandle, "VNImageRectForNormalizedRectUsingRegionOfInterest")
	registerFunc(&_vNNormalizedFaceBoundingBoxPointForLandmarkPoint, frameworkHandle, "VNNormalizedFaceBoundingBoxPointForLandmarkPoint")
	registerFunc(&_vNNormalizedPointForImagePoint, frameworkHandle, "VNNormalizedPointForImagePoint")
	registerFunc(&_vNNormalizedPointForImagePointUsingRegionOfInterest, frameworkHandle, "VNNormalizedPointForImagePointUsingRegionOfInterest")
	registerFunc(&_vNNormalizedRectForImageRect, frameworkHandle, "VNNormalizedRectForImageRect")
	registerFunc(&_vNNormalizedRectForImageRectUsingRegionOfInterest, frameworkHandle, "VNNormalizedRectForImageRectUsingRegionOfInterest")
	registerFunc(&_vNNormalizedRectIsIdentityRect, frameworkHandle, "VNNormalizedRectIsIdentityRect")
}
