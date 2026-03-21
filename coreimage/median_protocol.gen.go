// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a median filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMedian
type CIMedian interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMedian/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMedian/inputImage
	SetInputImage(value ICIImage)
}

// CIMedianObject wraps an existing Objective-C object that conforms to the CIMedian protocol.
type CIMedianObject struct {
	objectivec.Object
}
func (o CIMedianObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMedianObjectFromID constructs a [CIMedianObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMedianObjectFromID(id objc.ID) CIMedianObject {
	return CIMedianObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMedian/inputImage
func (o CIMedianObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMedianObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMedianObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

