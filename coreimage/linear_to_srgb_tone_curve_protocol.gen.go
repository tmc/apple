// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a linear-to-sRGB filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearToSRGBToneCurve
type CILinearToSRGBToneCurve interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearToSRGBToneCurve/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearToSRGBToneCurve/inputImage
	SetInputImage(value ICIImage)
}

// CILinearToSRGBToneCurveObject wraps an existing Objective-C object that conforms to the CILinearToSRGBToneCurve protocol.
type CILinearToSRGBToneCurveObject struct {
	objectivec.Object
}

func (o CILinearToSRGBToneCurveObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILinearToSRGBToneCurveObjectFromID constructs a [CILinearToSRGBToneCurveObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILinearToSRGBToneCurveObjectFromID(id objc.ID) CILinearToSRGBToneCurveObject {
	return CILinearToSRGBToneCurveObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearToSRGBToneCurve/inputImage
func (o CILinearToSRGBToneCurveObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILinearToSRGBToneCurveObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearToSRGBToneCurve/inputImage
func (o CILinearToSRGBToneCurveObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
