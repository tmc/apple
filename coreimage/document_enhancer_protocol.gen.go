// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a document enhancer filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer
type CIDocumentEnhancer interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount of enhancement.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer/amount
	Amount() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer/inputImage
	InputImage() ICIImage

	// The amount of enhancement.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer/amount
	SetAmount(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer/inputImage
	SetInputImage(value ICIImage)
}

// CIDocumentEnhancerObject wraps an existing Objective-C object that conforms to the CIDocumentEnhancer protocol.
type CIDocumentEnhancerObject struct {
	objectivec.Object
}
func (o CIDocumentEnhancerObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDocumentEnhancerObjectFromID constructs a [CIDocumentEnhancerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDocumentEnhancerObjectFromID(id objc.ID) CIDocumentEnhancerObject {
	return CIDocumentEnhancerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The amount of enhancement.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer/amount
func (o CIDocumentEnhancerObject) Amount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("amount"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDocumentEnhancer/inputImage
func (o CIDocumentEnhancerObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDocumentEnhancerObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIDocumentEnhancerObject) SetAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAmount:"), value)
}

func (o CIDocumentEnhancerObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

