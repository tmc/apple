// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a minimum component filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMinimumComponent
type CIMinimumComponent interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMinimumComponent/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMinimumComponent/inputImage
	SetInputImage(value ICIImage)
}

// CIMinimumComponentObject wraps an existing Objective-C object that conforms to the CIMinimumComponent protocol.
type CIMinimumComponentObject struct {
	objectivec.Object
}

func (o CIMinimumComponentObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMinimumComponentObjectFromID constructs a [CIMinimumComponentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMinimumComponentObjectFromID(id objc.ID) CIMinimumComponentObject {
	return CIMinimumComponentObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMinimumComponent/inputImage
func (o CIMinimumComponentObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMinimumComponentObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMinimumComponent/inputImage
func (o CIMinimumComponentObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
