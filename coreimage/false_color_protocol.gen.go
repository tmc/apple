// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a false color filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor
type CIFalseColor interface {
	objectivec.IObject
	CIFilterProtocol

	// The first color to use for the color ramp.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/color0
	Color0() ICIColor

	// The second color to use for the color ramp.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/color1
	Color1() ICIColor

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/inputImage
	InputImage() ICIImage

	// The first color to use for the color ramp.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/color0
	SetColor0(value ICIColor)

	// The second color to use for the color ramp.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/color1
	SetColor1(value ICIColor)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/inputImage
	SetInputImage(value ICIImage)
}

// CIFalseColorObject wraps an existing Objective-C object that conforms to the CIFalseColor protocol.
type CIFalseColorObject struct {
	objectivec.Object
}
func (o CIFalseColorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFalseColorObjectFromID constructs a [CIFalseColorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFalseColorObjectFromID(id objc.ID) CIFalseColorObject {
	return CIFalseColorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The first color to use for the color ramp.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/color0
func (o CIFalseColorObject) Color0() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
	}
// The second color to use for the color ramp.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/color1
func (o CIFalseColorObject) Color1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFalseColor/inputImage
func (o CIFalseColorObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFalseColorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIFalseColorObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

func (o CIFalseColorObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

func (o CIFalseColorObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

