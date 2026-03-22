// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIStretchCrop protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop
type CIStretchCrop interface {
	objectivec.IObject
	CIFilterProtocol

	// CenterStretchAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/centerStretchAmount
	CenterStretchAmount() float32

	// CropAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/cropAmount
	CropAmount() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/inputImage
	InputImage() ICIImage

	// Size protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/size
	Size() corefoundation.CGPoint

	// centerStretchAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/centerStretchAmount
	SetCenterStretchAmount(value float32)

	// cropAmount protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/cropAmount
	SetCropAmount(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/inputImage
	SetInputImage(value ICIImage)

	// size protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/size
	SetSize(value corefoundation.CGPoint)
}

// CIStretchCropObject wraps an existing Objective-C object that conforms to the CIStretchCrop protocol.
type CIStretchCropObject struct {
	objectivec.Object
}
func (o CIStretchCropObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIStretchCropObjectFromID constructs a [CIStretchCropObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIStretchCropObjectFromID(id objc.ID) CIStretchCropObject {
	return CIStretchCropObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/centerStretchAmount
func (o CIStretchCropObject) CenterStretchAmount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("centerStretchAmount"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/cropAmount
func (o CIStretchCropObject) CropAmount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("cropAmount"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/inputImage
func (o CIStretchCropObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIStretchCrop/size
func (o CIStretchCropObject) Size() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("size"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIStretchCropObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIStretchCropObject) SetCenterStretchAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenterStretchAmount:"), value)
}

func (o CIStretchCropObject) SetCropAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCropAmount:"), value)
}

func (o CIStretchCropObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIStretchCropObject) SetSize(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setSize:"), value)
}

