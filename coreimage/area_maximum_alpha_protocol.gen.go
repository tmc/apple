// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaMaximumAlpha protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaMaximumAlpha
type CIAreaMaximumAlpha interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIAreaMaximumAlphaObject wraps an existing Objective-C object that conforms to the CIAreaMaximumAlpha protocol.
type CIAreaMaximumAlphaObject struct {
	objectivec.Object
}
func (o CIAreaMaximumAlphaObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaMaximumAlphaObjectFromID constructs a [CIAreaMaximumAlphaObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaMaximumAlphaObjectFromID(id objc.ID) CIAreaMaximumAlphaObject {
	return CIAreaMaximumAlphaObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaMaximumAlphaObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaMaximumAlphaObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaMaximumAlphaObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaMaximumAlphaObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaMaximumAlphaObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

