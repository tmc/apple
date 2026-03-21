// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an unsharp mask filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask
type CIUnsharpMask interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/inputImage
	InputImage() ICIImage

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/intensity
	Intensity() float32

	// The radius of the unsharp mask effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/radius
	Radius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/intensity
	SetIntensity(value float32)

	// The radius of the unsharp mask effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/radius
	SetRadius(value float32)
}

// CIUnsharpMaskObject wraps an existing Objective-C object that conforms to the CIUnsharpMask protocol.
type CIUnsharpMaskObject struct {
	objectivec.Object
}
func (o CIUnsharpMaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIUnsharpMaskObjectFromID constructs a [CIUnsharpMaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIUnsharpMaskObjectFromID(id objc.ID) CIUnsharpMaskObject {
	return CIUnsharpMaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/inputImage
func (o CIUnsharpMaskObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/intensity
func (o CIUnsharpMaskObject) Intensity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// The radius of the unsharp mask effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIUnsharpMask/radius
func (o CIUnsharpMaskObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIUnsharpMaskObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIUnsharpMaskObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIUnsharpMaskObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

func (o CIUnsharpMaskObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

