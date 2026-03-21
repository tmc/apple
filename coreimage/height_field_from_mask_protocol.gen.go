// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a height-field-from-mask filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask
type CIHeightFieldFromMask interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask/inputImage
	InputImage() ICIImage

	// The length of the height-field transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask/inputImage
	SetInputImage(value ICIImage)

	// The length of the height-field transition.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask/radius
	SetRadius(value float32)
}

// CIHeightFieldFromMaskObject wraps an existing Objective-C object that conforms to the CIHeightFieldFromMask protocol.
type CIHeightFieldFromMaskObject struct {
	objectivec.Object
}
func (o CIHeightFieldFromMaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHeightFieldFromMaskObjectFromID constructs a [CIHeightFieldFromMaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHeightFieldFromMaskObjectFromID(id objc.ID) CIHeightFieldFromMaskObject {
	return CIHeightFieldFromMaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask/inputImage
func (o CIHeightFieldFromMaskObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The length of the height-field transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHeightFieldFromMask/radius
func (o CIHeightFieldFromMaskObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHeightFieldFromMaskObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHeightFieldFromMaskObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIHeightFieldFromMaskObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

