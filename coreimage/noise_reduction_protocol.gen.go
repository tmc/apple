// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a noise reduction filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction
type CINoiseReduction interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/inputImage
	InputImage() ICIImage

	// The amount of noise reduction.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/noiseLevel
	NoiseLevel() float32

	// The sharpness of the final image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/sharpness
	Sharpness() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/inputImage
	SetInputImage(value ICIImage)

	// The amount of noise reduction.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/noiseLevel
	SetNoiseLevel(value float32)

	// The sharpness of the final image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/sharpness
	SetSharpness(value float32)
}

// CINoiseReductionObject wraps an existing Objective-C object that conforms to the CINoiseReduction protocol.
type CINoiseReductionObject struct {
	objectivec.Object
}
func (o CINoiseReductionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CINoiseReductionObjectFromID constructs a [CINoiseReductionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CINoiseReductionObjectFromID(id objc.ID) CINoiseReductionObject {
	return CINoiseReductionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/inputImage
func (o CINoiseReductionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The amount of noise reduction.
//
// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/noiseLevel
func (o CINoiseReductionObject) NoiseLevel() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("noiseLevel"))
	return rv
	}
// The sharpness of the final image.
//
// See: https://developer.apple.com/documentation/CoreImage/CINoiseReduction/sharpness
func (o CINoiseReductionObject) Sharpness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CINoiseReductionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CINoiseReductionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CINoiseReductionObject) SetNoiseLevel(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setNoiseLevel:"), value)
}

func (o CINoiseReductionObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

