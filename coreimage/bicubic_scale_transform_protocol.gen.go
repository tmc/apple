// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a bicubic scale transform filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform
type CIBicubicScaleTransform interface {
	objectivec.IObject
	CIFilterProtocol

	// The additional horizontal scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/aspectRatio
	AspectRatio() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/inputImage
	InputImage() ICIImage

	// The value of B to use for the cubic resampling function.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/parameterB
	ParameterB() float32

	// The value of C to use for the cubic resampling function.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/parameterC
	ParameterC() float32

	// The scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/scale
	Scale() float32

	// The additional horizontal scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/aspectRatio
	SetAspectRatio(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/inputImage
	SetInputImage(value ICIImage)

	// The value of B to use for the cubic resampling function.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/parameterB
	SetParameterB(value float32)

	// The value of C to use for the cubic resampling function.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/parameterC
	SetParameterC(value float32)

	// The scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/scale
	SetScale(value float32)
}

// CIBicubicScaleTransformObject wraps an existing Objective-C object that conforms to the CIBicubicScaleTransform protocol.
type CIBicubicScaleTransformObject struct {
	objectivec.Object
}
func (o CIBicubicScaleTransformObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBicubicScaleTransformObjectFromID constructs a [CIBicubicScaleTransformObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBicubicScaleTransformObjectFromID(id objc.ID) CIBicubicScaleTransformObject {
	return CIBicubicScaleTransformObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The additional horizontal scaling factor to use on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/aspectRatio
func (o CIBicubicScaleTransformObject) AspectRatio() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("aspectRatio"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/inputImage
func (o CIBicubicScaleTransformObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The value of B to use for the cubic resampling function.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/parameterB
func (o CIBicubicScaleTransformObject) ParameterB() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("parameterB"))
	return rv
	}
// The value of C to use for the cubic resampling function.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/parameterC
func (o CIBicubicScaleTransformObject) ParameterC() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("parameterC"))
	return rv
	}
// The scaling factor to use on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBicubicScaleTransform/scale
func (o CIBicubicScaleTransformObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBicubicScaleTransformObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBicubicScaleTransformObject) SetAspectRatio(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAspectRatio:"), value)
}

func (o CIBicubicScaleTransformObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIBicubicScaleTransformObject) SetParameterB(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setParameterB:"), value)
}

func (o CIBicubicScaleTransformObject) SetParameterC(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setParameterC:"), value)
}

func (o CIBicubicScaleTransformObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

