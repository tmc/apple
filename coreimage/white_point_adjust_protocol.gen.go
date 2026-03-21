// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a white-point adjust filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust
type CIWhitePointAdjust interface {
	objectivec.IObject
	CIFilterProtocol

	// A color to use as the white point.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust/color
	Color() ICIColor

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust/inputImage
	InputImage() ICIImage

	// A color to use as the white point.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust/color
	SetColor(value ICIColor)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust/inputImage
	SetInputImage(value ICIImage)
}

// CIWhitePointAdjustObject wraps an existing Objective-C object that conforms to the CIWhitePointAdjust protocol.
type CIWhitePointAdjustObject struct {
	objectivec.Object
}
func (o CIWhitePointAdjustObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIWhitePointAdjustObjectFromID constructs a [CIWhitePointAdjustObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIWhitePointAdjustObjectFromID(id objc.ID) CIWhitePointAdjustObject {
	return CIWhitePointAdjustObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A color to use as the white point.
//
// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust/color
func (o CIWhitePointAdjustObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIWhitePointAdjust/inputImage
func (o CIWhitePointAdjustObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIWhitePointAdjustObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIWhitePointAdjustObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIWhitePointAdjustObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

