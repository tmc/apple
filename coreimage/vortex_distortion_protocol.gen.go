// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIVortexDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion
type CIVortexDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/angle
	Angle() float32

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/radius
	Radius() float32

	// angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/angle
	SetAngle(value float32)

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/radius
	SetRadius(value float32)
}

// CIVortexDistortionObject wraps an existing Objective-C object that conforms to the CIVortexDistortion protocol.
type CIVortexDistortionObject struct {
	objectivec.Object
}

func (o CIVortexDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIVortexDistortionObjectFromID constructs a [CIVortexDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIVortexDistortionObjectFromID(id objc.ID) CIVortexDistortionObject {
	return CIVortexDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/angle
func (o CIVortexDistortionObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/center
func (o CIVortexDistortionObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/inputImage
func (o CIVortexDistortionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/radius
func (o CIVortexDistortionObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIVortexDistortionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/angle
func (o CIVortexDistortionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/center
func (o CIVortexDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/inputImage
func (o CIVortexDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIVortexDistortion/radius
func (o CIVortexDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
