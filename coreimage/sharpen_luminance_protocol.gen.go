// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a sharpen luminance filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance
type CISharpenLuminance interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/inputImage
	InputImage() ICIImage

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/radius
	Radius() float32

	// The amount of sharpening to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/sharpness
	Sharpness() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/inputImage
	SetInputImage(value ICIImage)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/radius
	SetRadius(value float32)

	// The amount of sharpening to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/sharpness
	SetSharpness(value float32)
}

// CISharpenLuminanceObject wraps an existing Objective-C object that conforms to the CISharpenLuminance protocol.
type CISharpenLuminanceObject struct {
	objectivec.Object
}
func (o CISharpenLuminanceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISharpenLuminanceObjectFromID constructs a [CISharpenLuminanceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISharpenLuminanceObjectFromID(id objc.ID) CISharpenLuminanceObject {
	return CISharpenLuminanceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/inputImage
func (o CISharpenLuminanceObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/radius
func (o CISharpenLuminanceObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// The amount of sharpening to apply.
//
// See: https://developer.apple.com/documentation/CoreImage/CISharpenLuminance/sharpness
func (o CISharpenLuminanceObject) Sharpness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISharpenLuminanceObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISharpenLuminanceObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISharpenLuminanceObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CISharpenLuminanceObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

