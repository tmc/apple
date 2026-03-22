// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a gloom filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGloom
type CIGloom interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGloom/inputImage
	InputImage() ICIImage

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGloom/intensity
	Intensity() float32

	// The radius, in pixels, of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGloom/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGloom/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGloom/intensity
	SetIntensity(value float32)

	// The radius, in pixels, of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGloom/radius
	SetRadius(value float32)
}

// CIGloomObject wraps an existing Objective-C object that conforms to the CIGloom protocol.
type CIGloomObject struct {
	objectivec.Object
}
func (o CIGloomObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGloomObjectFromID constructs a [CIGloomObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGloomObjectFromID(id objc.ID) CIGloomObject {
	return CIGloomObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGloom/inputImage
func (o CIGloomObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGloom/intensity
func (o CIGloomObject) Intensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// The radius, in pixels, of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGloom/radius
func (o CIGloomObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGloomObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIGloomObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIGloomObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

func (o CIGloomObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

