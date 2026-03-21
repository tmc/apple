// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an X-ray filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIXRay
type CIXRay interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIXRay/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIXRay/inputImage
	SetInputImage(value ICIImage)
}

// CIXRayObject wraps an existing Objective-C object that conforms to the CIXRay protocol.
type CIXRayObject struct {
	objectivec.Object
}
func (o CIXRayObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIXRayObjectFromID constructs a [CIXRayObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIXRayObjectFromID(id objc.ID) CIXRayObject {
	return CIXRayObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIXRay/inputImage
func (o CIXRayObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIXRayObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIXRayObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

