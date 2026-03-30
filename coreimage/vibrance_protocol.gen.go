// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a vibrance filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVibrance
type CIVibrance interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount to adjust the saturation by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/amount
	Amount() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/inputImage
	InputImage() ICIImage

	// The amount to adjust the saturation by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/amount
	SetAmount(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/inputImage
	SetInputImage(value ICIImage)
}

// CIVibranceObject wraps an existing Objective-C object that conforms to the CIVibrance protocol.
type CIVibranceObject struct {
	objectivec.Object
}

func (o CIVibranceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIVibranceObjectFromID constructs a [CIVibranceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIVibranceObjectFromID(id objc.ID) CIVibranceObject {
	return CIVibranceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The amount to adjust the saturation by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/amount
func (o CIVibranceObject) Amount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("amount"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/inputImage
func (o CIVibranceObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIVibranceObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The amount to adjust the saturation by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/amount
func (o CIVibranceObject) SetAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAmount:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVibrance/inputImage
func (o CIVibranceObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
