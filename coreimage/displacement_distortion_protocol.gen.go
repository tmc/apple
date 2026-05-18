// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIDisplacementDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion
type CIDisplacementDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// displacementImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion/displacementImage
	DisplacementImage() ICIImage
	SetDisplacementImage(value ICIImage)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion/scale
	Scale() float32
	SetScale(value float32)
}

// CIDisplacementDistortionObject wraps an existing Objective-C object that conforms to the CIDisplacementDistortion protocol.
type CIDisplacementDistortionObject struct {
	objectivec.Object
}

func (o CIDisplacementDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDisplacementDistortionObjectFromID constructs a [CIDisplacementDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDisplacementDistortionObjectFromID(id objc.ID) CIDisplacementDistortionObject {
	return CIDisplacementDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDisplacementDistortionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion/displacementImage
func (o CIDisplacementDistortionObject) DisplacementImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("displacementImage"))
	return CIImageFromID(rv)
}

func (o CIDisplacementDistortionObject) SetDisplacementImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setDisplacementImage:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion/inputImage
func (o CIDisplacementDistortionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIDisplacementDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIDisplacementDistortion/scale
func (o CIDisplacementDistortionObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return float32(rv)
}

func (o CIDisplacementDistortionObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}
