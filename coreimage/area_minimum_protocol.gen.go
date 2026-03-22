// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaMinimum protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaMinimum
type CIAreaMinimum interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIAreaMinimumObject wraps an existing Objective-C object that conforms to the CIAreaMinimum protocol.
type CIAreaMinimumObject struct {
	objectivec.Object
}
func (o CIAreaMinimumObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaMinimumObjectFromID constructs a [CIAreaMinimumObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaMinimumObjectFromID(id objc.ID) CIAreaMinimumObject {
	return CIAreaMinimumObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaMinimumObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaMinimumObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaMinimumObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaMinimumObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaMinimumObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

