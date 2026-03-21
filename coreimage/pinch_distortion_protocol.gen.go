// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIPinchDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion
type CIPinchDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/radius
	Radius() float32

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/scale
	Scale() float32

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/radius
	SetRadius(value float32)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/scale
	SetScale(value float32)
}

// CIPinchDistortionObject wraps an existing Objective-C object that conforms to the CIPinchDistortion protocol.
type CIPinchDistortionObject struct {
	objectivec.Object
}
func (o CIPinchDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPinchDistortionObjectFromID constructs a [CIPinchDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPinchDistortionObjectFromID(id objc.ID) CIPinchDistortionObject {
	return CIPinchDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/center
func (o CIPinchDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/inputImage
func (o CIPinchDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/radius
func (o CIPinchDistortionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIPinchDistortion/scale
func (o CIPinchDistortionObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPinchDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIPinchDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIPinchDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIPinchDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIPinchDistortionObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

