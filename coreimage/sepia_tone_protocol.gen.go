// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a sepia-tone filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone
type CISepiaTone interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone/inputImage
	InputImage() ICIImage

	// The intensity of the sepia effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone/intensity
	Intensity() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the sepia effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone/intensity
	SetIntensity(value float32)
}

// CISepiaToneObject wraps an existing Objective-C object that conforms to the CISepiaTone protocol.
type CISepiaToneObject struct {
	objectivec.Object
}
func (o CISepiaToneObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISepiaToneObjectFromID constructs a [CISepiaToneObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISepiaToneObjectFromID(id objc.ID) CISepiaToneObject {
	return CISepiaToneObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone/inputImage
func (o CISepiaToneObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the sepia effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISepiaTone/intensity
func (o CISepiaToneObject) Intensity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISepiaToneObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISepiaToneObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISepiaToneObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

