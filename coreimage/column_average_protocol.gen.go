// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIColumnAverage protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColumnAverage
type CIColumnAverage interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIColumnAverageObject wraps an existing Objective-C object that conforms to the CIColumnAverage protocol.
type CIColumnAverageObject struct {
	objectivec.Object
}

func (o CIColumnAverageObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColumnAverageObjectFromID constructs a [CIColumnAverageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColumnAverageObjectFromID(id objc.ID) CIColumnAverageObject {
	return CIColumnAverageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIColumnAverageObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIColumnAverageObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColumnAverageObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIColumnAverageObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIColumnAverageObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
