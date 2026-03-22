// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CISobelGradients protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CISobelGradients
type CISobelGradients interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISobelGradients/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISobelGradients/inputImage
	SetInputImage(value ICIImage)
}

// CISobelGradientsObject wraps an existing Objective-C object that conforms to the CISobelGradients protocol.
type CISobelGradientsObject struct {
	objectivec.Object
}
func (o CISobelGradientsObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISobelGradientsObjectFromID constructs a [CISobelGradientsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISobelGradientsObjectFromID(id objc.ID) CISobelGradientsObject {
	return CISobelGradientsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISobelGradients/inputImage
func (o CISobelGradientsObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISobelGradientsObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISobelGradientsObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

