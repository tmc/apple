// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIHoleDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion
type CIHoleDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/radius
	Radius() float32

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/radius
	SetRadius(value float32)
}

// CIHoleDistortionObject wraps an existing Objective-C object that conforms to the CIHoleDistortion protocol.
type CIHoleDistortionObject struct {
	objectivec.Object
}
func (o CIHoleDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHoleDistortionObjectFromID constructs a [CIHoleDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHoleDistortionObjectFromID(id objc.ID) CIHoleDistortionObject {
	return CIHoleDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/center
func (o CIHoleDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/inputImage
func (o CIHoleDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIHoleDistortion/radius
func (o CIHoleDistortionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHoleDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHoleDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIHoleDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIHoleDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

