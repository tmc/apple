// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a morphology minimum filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum
type CIMorphologyMinimum interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/inputImage
	InputImage() ICIImage

	// The radius of the circular morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/inputImage
	SetInputImage(value ICIImage)

	// The radius of the circular morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/radius
	SetRadius(value float32)
}

// CIMorphologyMinimumObject wraps an existing Objective-C object that conforms to the CIMorphologyMinimum protocol.
type CIMorphologyMinimumObject struct {
	objectivec.Object
}

func (o CIMorphologyMinimumObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMorphologyMinimumObjectFromID constructs a [CIMorphologyMinimumObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMorphologyMinimumObjectFromID(id objc.ID) CIMorphologyMinimumObject {
	return CIMorphologyMinimumObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/inputImage
func (o CIMorphologyMinimumObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The radius of the circular morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/radius
func (o CIMorphologyMinimumObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMorphologyMinimumObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/inputImage
func (o CIMorphologyMinimumObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The radius of the circular morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMinimum/radius
func (o CIMorphologyMinimumObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
