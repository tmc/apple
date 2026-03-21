// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CICircleSplashDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion
type CICircleSplashDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/radius
	Radius() float32

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/radius
	SetRadius(value float32)
}

// CICircleSplashDistortionObject wraps an existing Objective-C object that conforms to the CICircleSplashDistortion protocol.
type CICircleSplashDistortionObject struct {
	objectivec.Object
}
func (o CICircleSplashDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICircleSplashDistortionObjectFromID constructs a [CICircleSplashDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICircleSplashDistortionObjectFromID(id objc.ID) CICircleSplashDistortionObject {
	return CICircleSplashDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/center
func (o CICircleSplashDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/inputImage
func (o CICircleSplashDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CICircleSplashDistortion/radius
func (o CICircleSplashDistortionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICircleSplashDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CICircleSplashDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CICircleSplashDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CICircleSplashDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

