// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a motion blur filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur
type CIMotionBlur interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of the motion, in radians, that determines which direction the blur smears.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/angle
	Angle() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/inputImage
	InputImage() ICIImage

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/radius
	Radius() float32

	// The angle of the motion, in radians, that determines which direction the blur smears.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/angle
	SetAngle(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/inputImage
	SetInputImage(value ICIImage)

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/radius
	SetRadius(value float32)
}

// CIMotionBlurObject wraps an existing Objective-C object that conforms to the CIMotionBlur protocol.
type CIMotionBlurObject struct {
	objectivec.Object
}

func (o CIMotionBlurObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIMotionBlurObjectFromID constructs a [CIMotionBlurObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIMotionBlurObjectFromID(id objc.ID) CIMotionBlurObject {
	return CIMotionBlurObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the motion, in radians, that determines which direction the
// blur smears.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/angle
func (o CIMotionBlurObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/inputImage
func (o CIMotionBlurObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The radius of the blur, in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/radius
func (o CIMotionBlurObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIMotionBlurObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The angle of the motion, in radians, that determines which direction the
// blur smears.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/angle
func (o CIMotionBlurObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/inputImage
func (o CIMotionBlurObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The radius of the blur, in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIMotionBlur/radius
func (o CIMotionBlurObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
