// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color posterize filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize
type CIColorPosterize interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/inputImage
	InputImage() ICIImage

	// The number of brightness levels to use for each color component.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/levels
	Levels() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/inputImage
	SetInputImage(value ICIImage)

	// The number of brightness levels to use for each color component.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/levels
	SetLevels(value float32)
}

// CIColorPosterizeObject wraps an existing Objective-C object that conforms to the CIColorPosterize protocol.
type CIColorPosterizeObject struct {
	objectivec.Object
}

func (o CIColorPosterizeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorPosterizeObjectFromID constructs a [CIColorPosterizeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorPosterizeObjectFromID(id objc.ID) CIColorPosterizeObject {
	return CIColorPosterizeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/inputImage
func (o CIColorPosterizeObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The number of brightness levels to use for each color component.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/levels
func (o CIColorPosterizeObject) Levels() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("levels"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorPosterizeObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/inputImage
func (o CIColorPosterizeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The number of brightness levels to use for each color component.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorPosterize/levels
func (o CIColorPosterizeObject) SetLevels(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setLevels:"), value)
}
