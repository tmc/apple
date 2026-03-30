// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a mix filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix
type CIMix interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMix/amount
	Amount() float32

	// The image to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMix/backgroundImage
	BackgroundImage() ICIImage

	// The image to use as a foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMix/inputImage
	InputImage() ICIImage

	// The amount of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMix/amount
	SetAmount(value float32)

	// The image to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMix/backgroundImage
	SetBackgroundImage(value ICIImage)

	// The image to use as a foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMix/inputImage
	SetInputImage(value ICIImage)
}

// CIMixObject wraps an existing Objective-C object that conforms to the CIMix protocol.
type CIMixObject struct {
	objectivec.Object
}

func (o CIMixObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMixObjectFromID constructs a [CIMixObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMixObjectFromID(id objc.ID) CIMixObject {
	return CIMixObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The amount of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix/amount
func (o CIMixObject) Amount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("amount"))
	return rv
}

// The image to use as a background image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix/backgroundImage
func (o CIMixObject) BackgroundImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("backgroundImage"))
	return CIImageFromID(rv)
}

// The image to use as a foreground image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix/inputImage
func (o CIMixObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMixObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The amount of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix/amount
func (o CIMixObject) SetAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAmount:"), value)
}

// The image to use as a background image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix/backgroundImage
func (o CIMixObject) SetBackgroundImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setBackgroundImage:"), value)
}

// The image to use as a foreground image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMix/inputImage
func (o CIMixObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
