// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a dither filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDither
type CIDither interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDither/inputImage
	InputImage() ICIImage

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDither/intensity
	Intensity() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDither/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDither/intensity
	SetIntensity(value float32)
}

// CIDitherObject wraps an existing Objective-C object that conforms to the CIDither protocol.
type CIDitherObject struct {
	objectivec.Object
}
func (o CIDitherObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDitherObjectFromID constructs a [CIDitherObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDitherObjectFromID(id objc.ID) CIDitherObject {
	return CIDitherObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDither/inputImage
func (o CIDitherObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDither/intensity
func (o CIDitherObject) Intensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDitherObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIDitherObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIDitherObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

