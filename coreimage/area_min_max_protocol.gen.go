// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaMinMax protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaMinMax
type CIAreaMinMax interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIAreaMinMaxObject wraps an existing Objective-C object that conforms to the CIAreaMinMax protocol.
type CIAreaMinMaxObject struct {
	objectivec.Object
}
func (o CIAreaMinMaxObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaMinMaxObjectFromID constructs a [CIAreaMinMaxObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaMinMaxObjectFromID(id objc.ID) CIAreaMinMaxObject {
	return CIAreaMinMaxObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaMinMaxObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaMinMaxObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaMinMaxObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaMinMaxObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaMinMaxObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

