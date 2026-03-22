// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a morphology rectangle maximum filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum
type CIMorphologyRectangleMaximum interface {
	objectivec.IObject
	CIFilterProtocol

	// The height, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/height
	Height() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/inputImage
	InputImage() ICIImage

	// The width, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/width
	Width() float32

	// The height, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/height
	SetHeight(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/inputImage
	SetInputImage(value ICIImage)

	// The width, in pixels, of the morphological structuring element.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/width
	SetWidth(value float32)
}

// CIMorphologyRectangleMaximumObject wraps an existing Objective-C object that conforms to the CIMorphologyRectangleMaximum protocol.
type CIMorphologyRectangleMaximumObject struct {
	objectivec.Object
}
func (o CIMorphologyRectangleMaximumObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMorphologyRectangleMaximumObjectFromID constructs a [CIMorphologyRectangleMaximumObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMorphologyRectangleMaximumObjectFromID(id objc.ID) CIMorphologyRectangleMaximumObject {
	return CIMorphologyRectangleMaximumObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The height, in pixels, of the morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/height
func (o CIMorphologyRectangleMaximumObject) Height() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("height"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/inputImage
func (o CIMorphologyRectangleMaximumObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width, in pixels, of the morphological structuring element.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMorphologyRectangleMaximum/width
func (o CIMorphologyRectangleMaximumObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMorphologyRectangleMaximumObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIMorphologyRectangleMaximumObject) SetHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHeight:"), value)
}

func (o CIMorphologyRectangleMaximumObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIMorphologyRectangleMaximumObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

