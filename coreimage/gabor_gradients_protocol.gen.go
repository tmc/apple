// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Gabor gradients filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaborGradients
type CIGaborGradients interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaborGradients/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaborGradients/inputImage
	SetInputImage(value ICIImage)
}

// CIGaborGradientsObject wraps an existing Objective-C object that conforms to the CIGaborGradients protocol.
type CIGaborGradientsObject struct {
	objectivec.Object
}

func (o CIGaborGradientsObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGaborGradientsObjectFromID constructs a [CIGaborGradientsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGaborGradientsObjectFromID(id objc.ID) CIGaborGradientsObject {
	return CIGaborGradientsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaborGradients/inputImage
func (o CIGaborGradientsObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGaborGradientsObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaborGradients/inputImage
func (o CIGaborGradientsObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
