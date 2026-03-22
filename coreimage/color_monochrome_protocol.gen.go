// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color monochrome filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome
type CIColorMonochrome interface {
	objectivec.IObject
	CIFilterProtocol

	// The monochrome color to apply to the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/color
	Color() ICIColor

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/inputImage
	InputImage() ICIImage

	// The intensity of the monochrome effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/intensity
	Intensity() float32

	// The monochrome color to apply to the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/color
	SetColor(value ICIColor)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the monochrome effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/intensity
	SetIntensity(value float32)
}

// CIColorMonochromeObject wraps an existing Objective-C object that conforms to the CIColorMonochrome protocol.
type CIColorMonochromeObject struct {
	objectivec.Object
}
func (o CIColorMonochromeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorMonochromeObjectFromID constructs a [CIColorMonochromeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorMonochromeObjectFromID(id objc.ID) CIColorMonochromeObject {
	return CIColorMonochromeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The monochrome color to apply to the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/color
func (o CIColorMonochromeObject) Color() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/inputImage
func (o CIColorMonochromeObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the monochrome effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMonochrome/intensity
func (o CIColorMonochromeObject) Intensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorMonochromeObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorMonochromeObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIColorMonochromeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorMonochromeObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

