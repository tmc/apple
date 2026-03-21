// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaHistogram protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram
type CIAreaHistogram interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol

	// Count protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram/count
	Count() int

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram/scale
	Scale() float32

	// count protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram/count
	SetCount(value int)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram/scale
	SetScale(value float32)
}

// CIAreaHistogramObject wraps an existing Objective-C object that conforms to the CIAreaHistogram protocol.
type CIAreaHistogramObject struct {
	objectivec.Object
}
func (o CIAreaHistogramObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaHistogramObjectFromID constructs a [CIAreaHistogramObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaHistogramObjectFromID(id objc.ID) CIAreaHistogramObject {
	return CIAreaHistogramObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram/count
func (o CIAreaHistogramObject) Count() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("count"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaHistogram/scale
func (o CIAreaHistogramObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaHistogramObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaHistogramObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaHistogramObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaHistogramObject) SetCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setCount:"), value)
}

func (o CIAreaHistogramObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

func (o CIAreaHistogramObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaHistogramObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

