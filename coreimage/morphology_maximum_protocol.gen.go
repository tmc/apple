// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a morphology maximum filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum
type CIMorphologyMaximum interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/inputImage
	InputImage() ICIImage

	// The radius of the circular morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/inputImage
	SetInputImage(value ICIImage)

	// The radius of the circular morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/radius
	SetRadius(value float32)
}

// CIMorphologyMaximumObject wraps an existing Objective-C object that conforms to the CIMorphologyMaximum protocol.
type CIMorphologyMaximumObject struct {
	objectivec.Object
}

func (o CIMorphologyMaximumObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMorphologyMaximumObjectFromID constructs a [CIMorphologyMaximumObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMorphologyMaximumObjectFromID(id objc.ID) CIMorphologyMaximumObject {
	return CIMorphologyMaximumObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/inputImage
func (o CIMorphologyMaximumObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The radius of the circular morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/radius
func (o CIMorphologyMaximumObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMorphologyMaximumObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/inputImage
func (o CIMorphologyMaximumObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The radius of the circular morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyMaximum/radius
func (o CIMorphologyMaximumObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
