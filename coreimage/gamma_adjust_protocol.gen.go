// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a gamma adjust filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust
type CIGammaAdjust interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust/inputImage
	InputImage() ICIImage

	// A gamma value to use to correct image brightness.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust/power
	Power() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust/inputImage
	SetInputImage(value ICIImage)

	// A gamma value to use to correct image brightness.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust/power
	SetPower(value float32)
}

// CIGammaAdjustObject wraps an existing Objective-C object that conforms to the CIGammaAdjust protocol.
type CIGammaAdjustObject struct {
	objectivec.Object
}
func (o CIGammaAdjustObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGammaAdjustObjectFromID constructs a [CIGammaAdjustObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGammaAdjustObjectFromID(id objc.ID) CIGammaAdjustObject {
	return CIGammaAdjustObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust/inputImage
func (o CIGammaAdjustObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A gamma value to use to correct image brightness.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGammaAdjust/power
func (o CIGammaAdjustObject) Power() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("power"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGammaAdjustObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIGammaAdjustObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIGammaAdjustObject) SetPower(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPower:"), value)
}

