// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a convolution filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIConvolution
type CIConvolution interface {
	objectivec.IObject
	CIFilterProtocol

	// A value that’s added to each output pixel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/bias
	Bias() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/inputImage
	InputImage() ICIImage

	// The convolution kernel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/weights
	Weights() ICIVector

	// A value that’s added to each output pixel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/bias
	SetBias(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/inputImage
	SetInputImage(value ICIImage)

	// The convolution kernel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/weights
	SetWeights(value ICIVector)
}

// CIConvolutionObject wraps an existing Objective-C object that conforms to the CIConvolution protocol.
type CIConvolutionObject struct {
	objectivec.Object
}
func (o CIConvolutionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIConvolutionObjectFromID constructs a [CIConvolutionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIConvolutionObjectFromID(id objc.ID) CIConvolutionObject {
	return CIConvolutionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A value that’s added to each output pixel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/bias
func (o CIConvolutionObject) Bias() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("bias"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/inputImage
func (o CIConvolutionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The convolution kernel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIConvolution/weights
func (o CIConvolutionObject) Weights() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("weights"))
	return CIVectorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIConvolutionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIConvolutionObject) SetBias(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setBias:"), value)
}

func (o CIConvolutionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIConvolutionObject) SetWeights(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setWeights:"), value)
}

