// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIHistogramDisplay protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay
type CIHistogramDisplay interface {
	objectivec.IObject
	CIFilterProtocol

	// Height protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/height
	Height() float32

	// HighLimit protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/highLimit
	HighLimit() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/inputImage
	InputImage() ICIImage

	// LowLimit protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/lowLimit
	LowLimit() float32

	// height protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/height
	SetHeight(value float32)

	// highLimit protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/highLimit
	SetHighLimit(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/inputImage
	SetInputImage(value ICIImage)

	// lowLimit protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/lowLimit
	SetLowLimit(value float32)
}

// CIHistogramDisplayObject wraps an existing Objective-C object that conforms to the CIHistogramDisplay protocol.
type CIHistogramDisplayObject struct {
	objectivec.Object
}
func (o CIHistogramDisplayObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHistogramDisplayObjectFromID constructs a [CIHistogramDisplayObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHistogramDisplayObjectFromID(id objc.ID) CIHistogramDisplayObject {
	return CIHistogramDisplayObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/height
func (o CIHistogramDisplayObject) Height() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("height"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/highLimit
func (o CIHistogramDisplayObject) HighLimit() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("highLimit"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/inputImage
func (o CIHistogramDisplayObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIHistogramDisplay/lowLimit
func (o CIHistogramDisplayObject) LowLimit() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("lowLimit"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHistogramDisplayObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHistogramDisplayObject) SetHeight(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHeight:"), value)
}

func (o CIHistogramDisplayObject) SetHighLimit(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHighLimit:"), value)
}

func (o CIHistogramDisplayObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIHistogramDisplayObject) SetLowLimit(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setLowLimit:"), value)
}

