// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The protocol for the Area Average and Maximum Red filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaAverageMaximumRed
type CIAreaAverageMaximumRed interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIAreaAverageMaximumRedObject wraps an existing Objective-C object that conforms to the CIAreaAverageMaximumRed protocol.
type CIAreaAverageMaximumRedObject struct {
	objectivec.Object
}

func (o CIAreaAverageMaximumRedObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaAverageMaximumRedObjectFromID constructs a [CIAreaAverageMaximumRedObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaAverageMaximumRedObjectFromID(id objc.ID) CIAreaAverageMaximumRedObject {
	return CIAreaAverageMaximumRedObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaAverageMaximumRedObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaAverageMaximumRedObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaAverageMaximumRedObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaAverageMaximumRedObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaAverageMaximumRedObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
