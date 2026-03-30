// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an sRGB-to-linear filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISRGBToneCurveToLinear
type CISRGBToneCurveToLinear interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISRGBToneCurveToLinear/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISRGBToneCurveToLinear/inputImage
	SetInputImage(value ICIImage)
}

// CISRGBToneCurveToLinearObject wraps an existing Objective-C object that conforms to the CISRGBToneCurveToLinear protocol.
type CISRGBToneCurveToLinearObject struct {
	objectivec.Object
}

func (o CISRGBToneCurveToLinearObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISRGBToneCurveToLinearObjectFromID constructs a [CISRGBToneCurveToLinearObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISRGBToneCurveToLinearObjectFromID(id objc.ID) CISRGBToneCurveToLinearObject {
	return CISRGBToneCurveToLinearObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISRGBToneCurveToLinear/inputImage
func (o CISRGBToneCurveToLinearObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISRGBToneCurveToLinearObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISRGBToneCurveToLinear/inputImage
func (o CISRGBToneCurveToLinearObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
