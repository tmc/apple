// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIAreaMaximum protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaMaximum
type CIAreaMaximum interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIAreaMaximumObject wraps an existing Objective-C object that conforms to the CIAreaMaximum protocol.
type CIAreaMaximumObject struct {
	objectivec.Object
}

func (o CIAreaMaximumObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaMaximumObjectFromID constructs a [CIAreaMaximumObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaMaximumObjectFromID(id objc.ID) CIAreaMaximumObject {
	return CIAreaMaximumObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaMaximumObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaMaximumObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaMaximumObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaMaximumObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaMaximumObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
