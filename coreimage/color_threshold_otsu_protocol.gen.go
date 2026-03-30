// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIColorThresholdOtsu protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorThresholdOtsu
type CIColorThresholdOtsu interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorThresholdOtsu/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorThresholdOtsu/inputImage
	SetInputImage(value ICIImage)
}

// CIColorThresholdOtsuObject wraps an existing Objective-C object that conforms to the CIColorThresholdOtsu protocol.
type CIColorThresholdOtsuObject struct {
	objectivec.Object
}

func (o CIColorThresholdOtsuObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorThresholdOtsuObjectFromID constructs a [CIColorThresholdOtsuObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorThresholdOtsuObjectFromID(id objc.ID) CIColorThresholdOtsuObject {
	return CIColorThresholdOtsuObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorThresholdOtsu/inputImage
func (o CIColorThresholdOtsuObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorThresholdOtsuObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorThresholdOtsu/inputImage
func (o CIColorThresholdOtsuObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
