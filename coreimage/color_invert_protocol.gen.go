// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color invert filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorInvert
type CIColorInvert interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorInvert/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorInvert/inputImage
	SetInputImage(value ICIImage)
}

// CIColorInvertObject wraps an existing Objective-C object that conforms to the CIColorInvert protocol.
type CIColorInvertObject struct {
	objectivec.Object
}

func (o CIColorInvertObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorInvertObjectFromID constructs a [CIColorInvertObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorInvertObjectFromID(id objc.ID) CIColorInvertObject {
	return CIColorInvertObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorInvert/inputImage
func (o CIColorInvertObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorInvertObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorInvert/inputImage
func (o CIColorInvertObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
