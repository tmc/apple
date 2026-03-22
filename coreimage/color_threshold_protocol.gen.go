// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIColorThreshold protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold
type CIColorThreshold interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold/inputImage
	InputImage() ICIImage

	// The threshold to use.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold/threshold
	Threshold() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold/inputImage
	SetInputImage(value ICIImage)

	// The threshold to use.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold/threshold
	SetThreshold(value float32)
}

// CIColorThresholdObject wraps an existing Objective-C object that conforms to the CIColorThreshold protocol.
type CIColorThresholdObject struct {
	objectivec.Object
}
func (o CIColorThresholdObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorThresholdObjectFromID constructs a [CIColorThresholdObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorThresholdObjectFromID(id objc.ID) CIColorThresholdObject {
	return CIColorThresholdObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold/inputImage
func (o CIColorThresholdObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The threshold to use.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorThreshold/threshold
func (o CIColorThresholdObject) Threshold() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("threshold"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorThresholdObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorThresholdObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorThresholdObject) SetThreshold(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setThreshold:"), value)
}

