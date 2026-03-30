// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an affine clamp filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp
type CIAffineClamp interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/inputImage
	InputImage() ICIImage

	// The transform to apply to the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/transform
	Transform() corefoundation.CGAffineTransform

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/inputImage
	SetInputImage(value ICIImage)

	// The transform to apply to the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/transform
	SetTransform(value corefoundation.CGAffineTransform)
}

// CIAffineClampObject wraps an existing Objective-C object that conforms to the CIAffineClamp protocol.
type CIAffineClampObject struct {
	objectivec.Object
}

func (o CIAffineClampObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAffineClampObjectFromID constructs a [CIAffineClampObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAffineClampObjectFromID(id objc.ID) CIAffineClampObject {
	return CIAffineClampObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/inputImage
func (o CIAffineClampObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The transform to apply to the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/transform
func (o CIAffineClampObject) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](o.ID, objc.Sel("transform"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAffineClampObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/inputImage
func (o CIAffineClampObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The transform to apply to the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAffineClamp/transform
func (o CIAffineClampObject) SetTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](o.ID, objc.Sel("setTransform:"), value)
}
