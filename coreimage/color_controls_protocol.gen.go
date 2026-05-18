// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color controls filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorControls
type CIColorControls interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount of brightness to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/brightness
	Brightness() float32
	SetBrightness(value float32)

	// The amount of contrast to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/contrast
	Contrast() float32
	SetContrast(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// The amount of saturation to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/saturation
	Saturation() float32
	SetSaturation(value float32)
}

// CIColorControlsObject wraps an existing Objective-C object that conforms to the CIColorControls protocol.
type CIColorControlsObject struct {
	objectivec.Object
}

func (o CIColorControlsObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorControlsObjectFromID constructs a [CIColorControlsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorControlsObjectFromID(id objc.ID) CIColorControlsObject {
	return CIColorControlsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorControlsObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The amount of brightness to apply.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/brightness
func (o CIColorControlsObject) Brightness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("brightness"))
	return float32(rv)
}

func (o CIColorControlsObject) SetBrightness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setBrightness:"), value)
}

// The amount of contrast to apply.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/contrast
func (o CIColorControlsObject) Contrast() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("contrast"))
	return float32(rv)
}

func (o CIColorControlsObject) SetContrast(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setContrast:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/inputImage
func (o CIColorControlsObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIColorControlsObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The amount of saturation to apply.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorControls/saturation
func (o CIColorControlsObject) Saturation() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("saturation"))
	return float32(rv)
}

func (o CIColorControlsObject) SetSaturation(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSaturation:"), value)
}
