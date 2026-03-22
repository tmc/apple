// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a box blur filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur
type CIBoxBlur interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur/inputImage
	InputImage() ICIImage

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur/inputImage
	SetInputImage(value ICIImage)

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur/radius
	SetRadius(value float32)
}

// CIBoxBlurObject wraps an existing Objective-C object that conforms to the CIBoxBlur protocol.
type CIBoxBlurObject struct {
	objectivec.Object
}
func (o CIBoxBlurObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBoxBlurObjectFromID constructs a [CIBoxBlurObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBoxBlurObjectFromID(id objc.ID) CIBoxBlurObject {
	return CIBoxBlurObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur/inputImage
func (o CIBoxBlurObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The radius of the blur, in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBoxBlur/radius
func (o CIBoxBlurObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBoxBlurObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBoxBlurObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIBoxBlurObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

