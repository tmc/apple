// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a maximum component filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaximumComponent
type CIMaximumComponent interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumComponent/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaximumComponent/inputImage
	SetInputImage(value ICIImage)
}

// CIMaximumComponentObject wraps an existing Objective-C object that conforms to the CIMaximumComponent protocol.
type CIMaximumComponentObject struct {
	objectivec.Object
}
func (o CIMaximumComponentObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMaximumComponentObjectFromID constructs a [CIMaximumComponentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMaximumComponentObjectFromID(id objc.ID) CIMaximumComponentObject {
	return CIMaximumComponentObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaximumComponent/inputImage
func (o CIMaximumComponentObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMaximumComponentObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMaximumComponentObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

