// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIAreaLogarithmicHistogram protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram
type CIAreaLogarithmicHistogram interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol

	// Count protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/count
	Count() int

	// MaximumStop protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/maximumStop
	MaximumStop() float32

	// MinimumStop protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/minimumStop
	MinimumStop() float32

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/scale
	Scale() float32

	// count protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/count
	SetCount(value int)

	// maximumStop protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/maximumStop
	SetMaximumStop(value float32)

	// minimumStop protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/minimumStop
	SetMinimumStop(value float32)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/scale
	SetScale(value float32)
}

// CIAreaLogarithmicHistogramObject wraps an existing Objective-C object that conforms to the CIAreaLogarithmicHistogram protocol.
type CIAreaLogarithmicHistogramObject struct {
	objectivec.Object
}
func (o CIAreaLogarithmicHistogramObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIAreaLogarithmicHistogramObjectFromID constructs a [CIAreaLogarithmicHistogramObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIAreaLogarithmicHistogramObjectFromID(id objc.ID) CIAreaLogarithmicHistogramObject {
	return CIAreaLogarithmicHistogramObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/count
func (o CIAreaLogarithmicHistogramObject) Count() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("count"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/maximumStop
func (o CIAreaLogarithmicHistogramObject) MaximumStop() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("maximumStop"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/minimumStop
func (o CIAreaLogarithmicHistogramObject) MinimumStop() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("minimumStop"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaLogarithmicHistogram/scale
func (o CIAreaLogarithmicHistogramObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIAreaLogarithmicHistogramObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIAreaLogarithmicHistogramObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIAreaLogarithmicHistogramObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIAreaLogarithmicHistogramObject) SetCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setCount:"), value)
}

func (o CIAreaLogarithmicHistogramObject) SetMaximumStop(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaximumStop:"), value)
}

func (o CIAreaLogarithmicHistogramObject) SetMinimumStop(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMinimumStop:"), value)
}

func (o CIAreaLogarithmicHistogramObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

func (o CIAreaLogarithmicHistogramObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIAreaLogarithmicHistogramObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

