// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a mask-to-alpha filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaskToAlpha
type CIMaskToAlpha interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskToAlpha/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskToAlpha/inputImage
	SetInputImage(value ICIImage)
}

// CIMaskToAlphaObject wraps an existing Objective-C object that conforms to the CIMaskToAlpha protocol.
type CIMaskToAlphaObject struct {
	objectivec.Object
}
func (o CIMaskToAlphaObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMaskToAlphaObjectFromID constructs a [CIMaskToAlphaObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMaskToAlphaObjectFromID(id objc.ID) CIMaskToAlphaObject {
	return CIMaskToAlphaObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaskToAlpha/inputImage
func (o CIMaskToAlphaObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMaskToAlphaObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMaskToAlphaObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

