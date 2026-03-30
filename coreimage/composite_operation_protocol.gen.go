// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a composite operation filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation
type CICompositeOperation interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/backgroundImage
	BackgroundImage() ICIImage

	// The image to use as a foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/inputImage
	InputImage() ICIImage

	// The image to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/backgroundImage
	SetBackgroundImage(value ICIImage)

	// The image to use as a foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/inputImage
	SetInputImage(value ICIImage)
}

// CICompositeOperationObject wraps an existing Objective-C object that conforms to the CICompositeOperation protocol.
type CICompositeOperationObject struct {
	objectivec.Object
}

func (o CICompositeOperationObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICompositeOperationObjectFromID constructs a [CICompositeOperationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICompositeOperationObjectFromID(id objc.ID) CICompositeOperationObject {
	return CICompositeOperationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as a background image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/backgroundImage
func (o CICompositeOperationObject) BackgroundImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("backgroundImage"))
	return CIImageFromID(rv)
}

// The image to use as a foreground image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/inputImage
func (o CICompositeOperationObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICompositeOperationObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as a background image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/backgroundImage
func (o CICompositeOperationObject) SetBackgroundImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setBackgroundImage:"), value)
}

// The image to use as a foreground image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICompositeOperation/inputImage
func (o CICompositeOperationObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
