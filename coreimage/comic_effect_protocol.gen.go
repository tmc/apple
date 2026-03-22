// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a comic effect filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIComicEffect
type CIComicEffect interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIComicEffect/inputImage
	InputImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIComicEffect/inputImage
	SetInputImage(value ICIImage)
}

// CIComicEffectObject wraps an existing Objective-C object that conforms to the CIComicEffect protocol.
type CIComicEffectObject struct {
	objectivec.Object
}
func (o CIComicEffectObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIComicEffectObjectFromID constructs a [CIComicEffectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIComicEffectObjectFromID(id objc.ID) CIComicEffectObject {
	return CIComicEffectObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIComicEffect/inputImage
func (o CIComicEffectObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIComicEffectObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIComicEffectObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

