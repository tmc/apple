// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an exposure adjust filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust
type CIExposureAdjust interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount to adjust the exposure of the image by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust/ev
	EV() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust/inputImage
	InputImage() ICIImage

	// The amount to adjust the exposure of the image by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust/ev
	SetEV(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust/inputImage
	SetInputImage(value ICIImage)
}

// CIExposureAdjustObject wraps an existing Objective-C object that conforms to the CIExposureAdjust protocol.
type CIExposureAdjustObject struct {
	objectivec.Object
}
func (o CIExposureAdjustObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIExposureAdjustObjectFromID constructs a [CIExposureAdjustObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIExposureAdjustObjectFromID(id objc.ID) CIExposureAdjustObject {
	return CIExposureAdjustObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The amount to adjust the exposure of the image by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust/ev
func (o CIExposureAdjustObject) EV() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("EV"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIExposureAdjust/inputImage
func (o CIExposureAdjustObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIExposureAdjustObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIExposureAdjustObject) SetEV(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setEV:"), value)
}

func (o CIExposureAdjustObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

