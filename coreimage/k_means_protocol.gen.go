// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIKMeans protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKMeans
type CIKMeans interface {
	objectivec.IObject
	CIAreaReductionFilter
	CIFilterProtocol

	// Count protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/count
	Count() int

	// InputMeans protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/inputMeans
	InputMeans() ICIImage

	// Passes protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/passes
	Passes() float32

	// Perceptual protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/perceptual
	Perceptual() bool

	// count protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/count
	SetCount(value int)

	// inputMeans protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/inputMeans
	SetInputMeans(value ICIImage)

	// passes protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/passes
	SetPasses(value float32)

	// perceptual protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/perceptual
	SetPerceptual(value bool)
}

// CIKMeansObject wraps an existing Objective-C object that conforms to the CIKMeans protocol.
type CIKMeansObject struct {
	objectivec.Object
}
func (o CIKMeansObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIKMeansObjectFromID constructs a [CIKMeansObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIKMeansObjectFromID(id objc.ID) CIKMeansObject {
	return CIKMeansObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/count
func (o CIKMeansObject) Count() int {
	rv := objc.Send[int](o.ID, objc.Sel("count"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/inputMeans
func (o CIKMeansObject) InputMeans() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputMeans"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/passes
func (o CIKMeansObject) Passes() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("passes"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIKMeans/perceptual
func (o CIKMeansObject) Perceptual() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("perceptual"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/extent
func (o CIKMeansObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIAreaReductionFilter/inputImage
func (o CIKMeansObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIKMeansObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIKMeansObject) SetCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setCount:"), value)
}

func (o CIKMeansObject) SetInputMeans(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputMeans:"), value)
}

func (o CIKMeansObject) SetPasses(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPasses:"), value)
}

func (o CIKMeansObject) SetPerceptual(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setPerceptual:"), value)
}

func (o CIKMeansObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIKMeansObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

