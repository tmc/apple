// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a photo-effect filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect
type CIPhotoEffect interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect/inputImage
	InputImage() ICIImage

	// Extrapolate for RGB values outside of the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect/extrapolate
	Extrapolate() bool

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect/inputImage
	SetInputImage(value ICIImage)

	// Extrapolate for RGB values outside of the range 0.0 to 1.0.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect/extrapolate
	SetExtrapolate(value bool)
}

// CIPhotoEffectObject wraps an existing Objective-C object that conforms to the CIPhotoEffect protocol.
type CIPhotoEffectObject struct {
	objectivec.Object
}
func (o CIPhotoEffectObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPhotoEffectObjectFromID constructs a [CIPhotoEffectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPhotoEffectObjectFromID(id objc.ID) CIPhotoEffectObject {
	return CIPhotoEffectObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect/inputImage
func (o CIPhotoEffectObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// Extrapolate for RGB values outside of the range 0.0 to 1.0.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPhotoEffect/extrapolate
func (o CIPhotoEffectObject) Extrapolate() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("extrapolate"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPhotoEffectObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIPhotoEffectObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIPhotoEffectObject) SetExtrapolate(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtrapolate:"), value)
}

