// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIGlassLozenge protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge
type CIGlassLozenge interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// point0 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/point0
	Point0() corefoundation.CGPoint
	SetPoint0(value corefoundation.CGPoint)

	// point1 protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/point1
	Point1() corefoundation.CGPoint
	SetPoint1(value corefoundation.CGPoint)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/radius
	Radius() float32
	SetRadius(value float32)

	// refraction protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/refraction
	Refraction() float32
	SetRefraction(value float32)
}

// CIGlassLozengeObject wraps an existing Objective-C object that conforms to the CIGlassLozenge protocol.
type CIGlassLozengeObject struct {
	objectivec.Object
}

func (o CIGlassLozengeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGlassLozengeObjectFromID constructs a [CIGlassLozengeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGlassLozengeObjectFromID(id objc.ID) CIGlassLozengeObject {
	return CIGlassLozengeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGlassLozengeObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/inputImage
func (o CIGlassLozengeObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIGlassLozengeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/point0
func (o CIGlassLozengeObject) Point0() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point0"))
	return corefoundation.CGPoint(rv)
}

func (o CIGlassLozengeObject) SetPoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint0:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/point1
func (o CIGlassLozengeObject) Point1() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point1"))
	return corefoundation.CGPoint(rv)
}

func (o CIGlassLozengeObject) SetPoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint1:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/radius
func (o CIGlassLozengeObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return float32(rv)
}

func (o CIGlassLozengeObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIGlassLozenge/refraction
func (o CIGlassLozengeObject) Refraction() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("refraction"))
	return float32(rv)
}

func (o CIGlassLozengeObject) SetRefraction(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRefraction:"), value)
}
