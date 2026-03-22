// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Gaussian blur filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur
type CIGaussianBlur interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur/inputImage
	InputImage() ICIImage

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur/inputImage
	SetInputImage(value ICIImage)

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur/radius
	SetRadius(value float32)
}

// CIGaussianBlurObject wraps an existing Objective-C object that conforms to the CIGaussianBlur protocol.
type CIGaussianBlurObject struct {
	objectivec.Object
}
func (o CIGaussianBlurObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGaussianBlurObjectFromID constructs a [CIGaussianBlurObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGaussianBlurObjectFromID(id objc.ID) CIGaussianBlurObject {
	return CIGaussianBlurObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur/inputImage
func (o CIGaussianBlurObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The radius of the blur, in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianBlur/radius
func (o CIGaussianBlurObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGaussianBlurObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIGaussianBlurObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIGaussianBlurObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

