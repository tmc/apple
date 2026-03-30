// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Lanczos scale transform filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform
type CILanczosScaleTransform interface {
	objectivec.IObject
	CIFilterProtocol

	// The additional horizontal scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/aspectRatio
	AspectRatio() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/inputImage
	InputImage() ICIImage

	// The scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/scale
	Scale() float32

	// The additional horizontal scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/aspectRatio
	SetAspectRatio(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/inputImage
	SetInputImage(value ICIImage)

	// The scaling factor to use on the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/scale
	SetScale(value float32)
}

// CILanczosScaleTransformObject wraps an existing Objective-C object that conforms to the CILanczosScaleTransform protocol.
type CILanczosScaleTransformObject struct {
	objectivec.Object
}

func (o CILanczosScaleTransformObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILanczosScaleTransformObjectFromID constructs a [CILanczosScaleTransformObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILanczosScaleTransformObjectFromID(id objc.ID) CILanczosScaleTransformObject {
	return CILanczosScaleTransformObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The additional horizontal scaling factor to use on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/aspectRatio
func (o CILanczosScaleTransformObject) AspectRatio() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("aspectRatio"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/inputImage
func (o CILanczosScaleTransformObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The scaling factor to use on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/scale
func (o CILanczosScaleTransformObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILanczosScaleTransformObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The additional horizontal scaling factor to use on the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/aspectRatio
func (o CILanczosScaleTransformObject) SetAspectRatio(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAspectRatio:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/inputImage
func (o CILanczosScaleTransformObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The scaling factor to use on the image.
//
// # Discussion
//
// Values less than 1.0 scale down the images. Values greater than 1.0 scale
// up the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILanczosScaleTransform/scale
func (o CILanczosScaleTransformObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}
