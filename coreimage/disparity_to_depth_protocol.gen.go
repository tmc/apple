// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a disparity-to-depth filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisparityToDepth
type CIDisparityToDepth interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisparityToDepth/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisparityToDepth/inputImage
	SetInputImage(value ICIImage)
}

// CIDisparityToDepthObject wraps an existing Objective-C object that conforms to the CIDisparityToDepth protocol.
type CIDisparityToDepthObject struct {
	objectivec.Object
}
func (o CIDisparityToDepthObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDisparityToDepthObjectFromID constructs a [CIDisparityToDepthObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDisparityToDepthObjectFromID(id objc.ID) CIDisparityToDepthObject {
	return CIDisparityToDepthObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisparityToDepth/inputImage
func (o CIDisparityToDepthObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDisparityToDepthObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIDisparityToDepthObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

