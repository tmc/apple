// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a masked variable blur filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur
type CIMaskedVariableBlur interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/inputImage
	InputImage() ICIImage

	// A grayscale mask that defines the blur amount.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/mask
	Mask() ICIImage

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/inputImage
	SetInputImage(value ICIImage)

	// A grayscale mask that defines the blur amount.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/mask
	SetMask(value ICIImage)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/radius
	SetRadius(value float32)
}

// CIMaskedVariableBlurObject wraps an existing Objective-C object that conforms to the CIMaskedVariableBlur protocol.
type CIMaskedVariableBlurObject struct {
	objectivec.Object
}
func (o CIMaskedVariableBlurObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMaskedVariableBlurObjectFromID constructs a [CIMaskedVariableBlurObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMaskedVariableBlurObjectFromID(id objc.ID) CIMaskedVariableBlurObject {
	return CIMaskedVariableBlurObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/inputImage
func (o CIMaskedVariableBlurObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A grayscale mask that defines the blur amount.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/mask
func (o CIMaskedVariableBlurObject) Mask() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("mask"))
	return CIImageFromID(rv)
	}
// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMaskedVariableBlur/radius
func (o CIMaskedVariableBlurObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMaskedVariableBlurObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMaskedVariableBlurObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIMaskedVariableBlurObject) SetMask(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setMask:"), value)
}

func (o CIMaskedVariableBlurObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

