// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a thermal filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIThermal
type CIThermal interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIThermal/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIThermal/inputImage
	SetInputImage(value ICIImage)
}

// CIThermalObject wraps an existing Objective-C object that conforms to the CIThermal protocol.
type CIThermalObject struct {
	objectivec.Object
}
func (o CIThermalObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIThermalObjectFromID constructs a [CIThermalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIThermalObjectFromID(id objc.ID) CIThermalObject {
	return CIThermalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIThermal/inputImage
func (o CIThermalObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIThermalObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIThermalObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

