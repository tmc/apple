// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color curves filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves
type CIColorCurves interface {
	objectivec.IObject
	CIFilterProtocol

	// The working color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/colorSpace
	ColorSpace() coregraphics.CGColorSpaceRef
	SetColorSpace(value coregraphics.CGColorSpaceRef)

	// Color values that determine the color curves transform.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesData
	CurvesData() foundation.NSData
	SetCurvesData(value foundation.NSData)

	// A two-element vector that defines the minimum and maximum values of the curve data.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesDomain
	CurvesDomain() ICIVector
	SetCurvesDomain(value ICIVector)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)
}

// CIColorCurvesObject wraps an existing Objective-C object that conforms to the CIColorCurves protocol.
type CIColorCurvesObject struct {
	objectivec.Object
}

func (o CIColorCurvesObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorCurvesObjectFromID constructs a [CIColorCurvesObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorCurvesObjectFromID(id objc.ID) CIColorCurvesObject {
	return CIColorCurvesObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorCurvesObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The working color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/colorSpace
func (o CIColorCurvesObject) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](o.ID, objc.Sel("colorSpace"))
	return coregraphics.CGColorSpaceRef(rv)
}

func (o CIColorCurvesObject) SetColorSpace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorSpace:"), value)
}

// Color values that determine the color curves transform.
//
// # Discussion
//
// Create the curves data as an [NSData] object containing a sequence of
// single-precision RGB values. These values represent a lookup table that’s
// applied to the input image.
//
// Core Image unpremultiplies the image before applying the effect, and
// premultiplies the result after applying the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesData
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
func (o CIColorCurvesObject) CurvesData() foundation.NSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("curvesData"))
	return foundation.NSDataFromID(rv)
}

func (o CIColorCurvesObject) SetCurvesData(value foundation.NSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setCurvesData:"), value)
}

// A two-element vector that defines the minimum and maximum values of the
// curve data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesDomain
func (o CIColorCurvesObject) CurvesDomain() ICIVector {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("curvesDomain"))
	return CIVectorFromID(rv)
}

func (o CIColorCurvesObject) SetCurvesDomain(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setCurvesDomain:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/inputImage
func (o CIColorCurvesObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIColorCurvesObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
