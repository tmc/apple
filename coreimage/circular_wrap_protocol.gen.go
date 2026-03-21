// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CICircularWrap protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap
type CICircularWrap interface {
	objectivec.IObject
	CIFilterProtocol

	// Angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/angle
	Angle() float32

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/inputImage
	InputImage() ICIImage

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/radius
	Radius() float32

	// angle protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/angle
	SetAngle(value float32)

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/inputImage
	SetInputImage(value ICIImage)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/radius
	SetRadius(value float32)
}

// CICircularWrapObject wraps an existing Objective-C object that conforms to the CICircularWrap protocol.
type CICircularWrapObject struct {
	objectivec.Object
}
func (o CICircularWrapObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICircularWrapObjectFromID constructs a [CICircularWrapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICircularWrapObjectFromID(id objc.ID) CICircularWrapObject {
	return CICircularWrapObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/angle
func (o CICircularWrapObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/center
func (o CICircularWrapObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/inputImage
func (o CICircularWrapObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CICircularWrap/radius
func (o CICircularWrapObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICircularWrapObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CICircularWrapObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CICircularWrapObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CICircularWrapObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CICircularWrapObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

