// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a morphology rectangle minimum filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum
type CIMorphologyRectangleMinimum interface {
	objectivec.IObject
	CIFilterProtocol

	// The height, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/height
	Height() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/inputImage
	InputImage() ICIImage

	// The width, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/width
	Width() float32

	// The height, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/height
	SetHeight(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/inputImage
	SetInputImage(value ICIImage)

	// The width, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/width
	SetWidth(value float32)
}

// CIMorphologyRectangleMinimumObject wraps an existing Objective-C object that conforms to the CIMorphologyRectangleMinimum protocol.
type CIMorphologyRectangleMinimumObject struct {
	objectivec.Object
}
func (o CIMorphologyRectangleMinimumObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMorphologyRectangleMinimumObjectFromID constructs a [CIMorphologyRectangleMinimumObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMorphologyRectangleMinimumObjectFromID(id objc.ID) CIMorphologyRectangleMinimumObject {
	return CIMorphologyRectangleMinimumObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The height, in pixels, of the morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/height
func (o CIMorphologyRectangleMinimumObject) Height() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("height"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/inputImage
func (o CIMorphologyRectangleMinimumObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width, in pixels, of the morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMinimum/width
func (o CIMorphologyRectangleMinimumObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMorphologyRectangleMinimumObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMorphologyRectangleMinimumObject) SetHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHeight:"), value)
}

func (o CIMorphologyRectangleMinimumObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIMorphologyRectangleMinimumObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

