// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a depth-to-disparity filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthToDisparity
type CIDepthToDisparity interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthToDisparity/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthToDisparity/inputImage
	SetInputImage(value ICIImage)
}

// CIDepthToDisparityObject wraps an existing Objective-C object that conforms to the CIDepthToDisparity protocol.
type CIDepthToDisparityObject struct {
	objectivec.Object
}
func (o CIDepthToDisparityObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDepthToDisparityObjectFromID constructs a [CIDepthToDisparityObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDepthToDisparityObjectFromID(id objc.ID) CIDepthToDisparityObject {
	return CIDepthToDisparityObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthToDisparity/inputImage
func (o CIDepthToDisparityObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDepthToDisparityObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIDepthToDisparityObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

