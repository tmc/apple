// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
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

	// Color values that determine the color curves transform.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesData
	CurvesData() foundation.INSData

	// A two-element vector that defines the minimum and maximum values of the curve data.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesDomain
	CurvesDomain() ICIVector

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/inputImage
	InputImage() ICIImage

	// The working color space.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/colorSpace
	SetColorSpace(value coregraphics.CGColorSpaceRef)

	// Color values that determine the color curves transform.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesData
	SetCurvesData(value foundation.INSData)

	// A two-element vector that defines the minimum and maximum values of the curve data.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesDomain
	SetCurvesDomain(value ICIVector)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/inputImage
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

// The working color space.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/colorSpace
func (o CIColorCurvesObject) ColorSpace() coregraphics.CGColorSpaceRef {
	
	rv := objc.Send[coregraphics.CGColorSpaceRef](o.ID, objc.Sel("colorSpace"))
	return rv
	}
// Color values that determine the color curves transform.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesData
func (o CIColorCurvesObject) CurvesData() foundation.INSData {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("curvesData"))
	return foundation.NSDataFromID(rv)
	}
// A two-element vector that defines the minimum and maximum values of the
// curve data.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/curvesDomain
func (o CIColorCurvesObject) CurvesDomain() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("curvesDomain"))
	return CIVectorFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorCurves/inputImage
func (o CIColorCurvesObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorCurvesObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorCurvesObject) SetColorSpace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorSpace:"), value)
}

func (o CIColorCurvesObject) SetCurvesData(value foundation.INSData) {
	objc.Send[struct{}](o.ID, objc.Sel("setCurvesData:"), value)
}

func (o CIColorCurvesObject) SetCurvesDomain(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setCurvesDomain:"), value)
}

func (o CIColorCurvesObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

