// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CITorusLensDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion
type CITorusLensDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/radius
	Radius() float32

	// Refraction protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/refraction
	Refraction() float32

	// Width protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/width
	Width() float32

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/radius
	SetRadius(value float32)

	// refraction protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/refraction
	SetRefraction(value float32)

	// width protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/width
	SetWidth(value float32)
}

// CITorusLensDistortionObject wraps an existing Objective-C object that conforms to the CITorusLensDistortion protocol.
type CITorusLensDistortionObject struct {
	objectivec.Object
}
func (o CITorusLensDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITorusLensDistortionObjectFromID constructs a [CITorusLensDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITorusLensDistortionObjectFromID(id objc.ID) CITorusLensDistortionObject {
	return CITorusLensDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/center
func (o CITorusLensDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/inputImage
func (o CITorusLensDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/radius
func (o CITorusLensDistortionObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/refraction
func (o CITorusLensDistortionObject) Refraction() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("refraction"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CITorusLensDistortion/width
func (o CITorusLensDistortionObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITorusLensDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITorusLensDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CITorusLensDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CITorusLensDistortionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CITorusLensDistortionObject) SetRefraction(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRefraction:"), value)
}

func (o CITorusLensDistortionObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

