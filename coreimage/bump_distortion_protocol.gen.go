// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIBumpDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion
type CIBumpDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/radius
	Radius() float32

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/scale
	Scale() float32

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/radius
	SetRadius(value float32)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/scale
	SetScale(value float32)
}

// CIBumpDistortionObject wraps an existing Objective-C object that conforms to the CIBumpDistortion protocol.
type CIBumpDistortionObject struct {
	objectivec.Object
}
func (o CIBumpDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBumpDistortionObjectFromID constructs a [CIBumpDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBumpDistortionObjectFromID(id objc.ID) CIBumpDistortionObject {
	return CIBumpDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/center
func (o CIBumpDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/inputImage
func (o CIBumpDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/radius
func (o CIBumpDistortionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIBumpDistortion/scale
func (o CIBumpDistortionObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBumpDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBumpDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIBumpDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIBumpDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIBumpDistortionObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

