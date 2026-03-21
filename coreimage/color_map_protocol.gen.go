// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a color map filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMap
type CIColorMap interface {
	objectivec.IObject
	CIFilterProtocol

	// The image data that transforms the source image values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMap/gradientImage
	GradientImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMap/inputImage
	InputImage() ICIImage

	// The image data that transforms the source image values.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMap/gradientImage
	SetGradientImage(value ICIImage)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorMap/inputImage
	SetInputImage(value ICIImage)
}

// CIColorMapObject wraps an existing Objective-C object that conforms to the CIColorMap protocol.
type CIColorMapObject struct {
	objectivec.Object
}
func (o CIColorMapObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorMapObjectFromID constructs a [CIColorMapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorMapObjectFromID(id objc.ID) CIColorMapObject {
	return CIColorMapObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image data that transforms the source image values.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMap/gradientImage
func (o CIColorMapObject) GradientImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("gradientImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorMap/inputImage
func (o CIColorMapObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorMapObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorMapObject) SetGradientImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setGradientImage:"), value)
}

func (o CIColorMapObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

