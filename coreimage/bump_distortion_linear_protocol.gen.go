// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIBumpDistortionLinear protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear
type CIBumpDistortionLinear interface {
	objectivec.IObject
	CIFilterProtocol

	// Angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/angle
	Angle() float32

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/radius
	Radius() float32

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/scale
	Scale() float32

	// angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/angle
	SetAngle(value float32)

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/radius
	SetRadius(value float32)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/scale
	SetScale(value float32)
}

// CIBumpDistortionLinearObject wraps an existing Objective-C object that conforms to the CIBumpDistortionLinear protocol.
type CIBumpDistortionLinearObject struct {
	objectivec.Object
}
func (o CIBumpDistortionLinearObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBumpDistortionLinearObjectFromID constructs a [CIBumpDistortionLinearObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBumpDistortionLinearObjectFromID(id objc.ID) CIBumpDistortionLinearObject {
	return CIBumpDistortionLinearObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/angle
func (o CIBumpDistortionLinearObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/center
func (o CIBumpDistortionLinearObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/inputImage
func (o CIBumpDistortionLinearObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/radius
func (o CIBumpDistortionLinearObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortionLinear/scale
func (o CIBumpDistortionLinearObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBumpDistortionLinearObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBumpDistortionLinearObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIBumpDistortionLinearObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIBumpDistortionLinearObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIBumpDistortionLinearObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIBumpDistortionLinearObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

