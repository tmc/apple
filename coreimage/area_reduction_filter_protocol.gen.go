// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaReductionFilter protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter
type CIAreaReductionFilter interface {
	objectivec.IObject
	CIFilterProtocol

	// Extent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
	Extent() corefoundation.CGRect

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
	InputImage() ICIImage

	// extent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
	SetExtent(value corefoundation.CGRect)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
	SetInputImage(value ICIImage)
}

// CIAreaReductionFilterObject wraps an existing Objective-C object that conforms to the CIAreaReductionFilter protocol.
type CIAreaReductionFilterObject struct {
	objectivec.Object
}
func (o CIAreaReductionFilterObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaReductionFilterObjectFromID constructs a [CIAreaReductionFilterObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaReductionFilterObjectFromID(id objc.ID) CIAreaReductionFilterObject {
	return CIAreaReductionFilterObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaReductionFilterObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaReductionFilterObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaReductionFilterObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaReductionFilterObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaReductionFilterObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

