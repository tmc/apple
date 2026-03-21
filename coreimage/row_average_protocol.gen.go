// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIRowAverage protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRowAverage
type CIRowAverage interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol
}

// CIRowAverageObject wraps an existing Objective-C object that conforms to the CIRowAverage protocol.
type CIRowAverageObject struct {
	objectivec.Object
}
func (o CIRowAverageObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRowAverageObjectFromID constructs a [CIRowAverageObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRowAverageObjectFromID(id objc.ID) CIRowAverageObject {
	return CIRowAverageObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIRowAverageObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIRowAverageObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRowAverageObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIRowAverageObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIRowAverageObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

