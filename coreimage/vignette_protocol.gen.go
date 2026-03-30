// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a vignette filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette
type CIVignette interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignette/inputImage
	InputImage() ICIImage

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignette/intensity
	Intensity() float32

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignette/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignette/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignette/intensity
	SetIntensity(value float32)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignette/radius
	SetRadius(value float32)
}

// CIVignetteObject wraps an existing Objective-C object that conforms to the CIVignette protocol.
type CIVignetteObject struct {
	objectivec.Object
}

func (o CIVignetteObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIVignetteObjectFromID constructs a [CIVignetteObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIVignetteObjectFromID(id objc.ID) CIVignetteObject {
	return CIVignetteObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette/inputImage
func (o CIVignetteObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette/intensity
func (o CIVignetteObject) Intensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
}

// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette/radius
func (o CIVignetteObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIVignetteObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette/inputImage
func (o CIVignetteObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette/intensity
func (o CIVignetteObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignette/radius
func (o CIVignetteObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
