// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CITwirlDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion
type CITwirlDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/angle
	Angle() float32

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/radius
	Radius() float32

	// angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/angle
	SetAngle(value float32)

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/radius
	SetRadius(value float32)
}

// CITwirlDistortionObject wraps an existing Objective-C object that conforms to the CITwirlDistortion protocol.
type CITwirlDistortionObject struct {
	objectivec.Object
}
func (o CITwirlDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITwirlDistortionObjectFromID constructs a [CITwirlDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITwirlDistortionObjectFromID(id objc.ID) CITwirlDistortionObject {
	return CITwirlDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/angle
func (o CITwirlDistortionObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/center
func (o CITwirlDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/inputImage
func (o CITwirlDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CITwirlDistortion/radius
func (o CITwirlDistortionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITwirlDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITwirlDistortionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CITwirlDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CITwirlDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CITwirlDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

