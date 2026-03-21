// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an affine tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile
type CIAffineTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile/inputImage
	InputImage() ICIImage

	// The transform to apply to the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile/transform
	Transform() corefoundation.CGAffineTransform

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile/inputImage
	SetInputImage(value ICIImage)

	// The transform to apply to the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile/transform
	SetTransform(value corefoundation.CGAffineTransform)
}

// CIAffineTileObject wraps an existing Objective-C object that conforms to the CIAffineTile protocol.
type CIAffineTileObject struct {
	objectivec.Object
}
func (o CIAffineTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAffineTileObjectFromID constructs a [CIAffineTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAffineTileObjectFromID(id objc.ID) CIAffineTileObject {
	return CIAffineTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile/inputImage
func (o CIAffineTileObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The transform to apply to the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineTile/transform
func (o CIAffineTileObject) Transform() corefoundation.CGAffineTransform {
	
	rv := objc.Send[corefoundation.CGAffineTransform](o.ID, objc.Sel("transform"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAffineTileObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAffineTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIAffineTileObject) SetTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](o.ID, objc.Sel("setTransform:"), value)
}

