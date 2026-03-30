// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a morphology gradient filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient
type CIMorphologyGradient interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/inputImage
	InputImage() ICIImage

	// The radius of the circular morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/inputImage
	SetInputImage(value ICIImage)

	// The radius of the circular morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/radius
	SetRadius(value float32)
}

// CIMorphologyGradientObject wraps an existing Objective-C object that conforms to the CIMorphologyGradient protocol.
type CIMorphologyGradientObject struct {
	objectivec.Object
}

func (o CIMorphologyGradientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMorphologyGradientObjectFromID constructs a [CIMorphologyGradientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMorphologyGradientObjectFromID(id objc.ID) CIMorphologyGradientObject {
	return CIMorphologyGradientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/inputImage
func (o CIMorphologyGradientObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The radius of the circular morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/radius
func (o CIMorphologyGradientObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMorphologyGradientObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/inputImage
func (o CIMorphologyGradientObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The radius of the circular morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyGradient/radius
func (o CIMorphologyGradientObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
