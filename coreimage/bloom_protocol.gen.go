// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a bloom filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom
type CIBloom interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBloom/inputImage
	InputImage() ICIImage

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBloom/intensity
	Intensity() float32

	// The radius, in pixels, of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBloom/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBloom/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBloom/intensity
	SetIntensity(value float32)

	// The radius, in pixels, of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBloom/radius
	SetRadius(value float32)
}

// CIBloomObject wraps an existing Objective-C object that conforms to the CIBloom protocol.
type CIBloomObject struct {
	objectivec.Object
}

func (o CIBloomObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBloomObjectFromID constructs a [CIBloomObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBloomObjectFromID(id objc.ID) CIBloomObject {
	return CIBloomObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom/inputImage
func (o CIBloomObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom/intensity
func (o CIBloomObject) Intensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
}

// The radius, in pixels, of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom/radius
func (o CIBloomObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBloomObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom/inputImage
func (o CIBloomObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The intensity of the effect.
//
// # Discussion
//
// A value of 0.0 is no effect. A value of 1.0 is the maximum effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom/intensity
func (o CIBloomObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

// The radius, in pixels, of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBloom/radius
func (o CIBloomObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
