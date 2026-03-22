// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a rounded rectangle generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator
type CIRoundedRectangleGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The color of the rounded rectangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/color
	Color() ICIColor

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/extent
	Extent() corefoundation.CGRect

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/radius
	Radius() float32

	// A value to control the smoothness of the transition between the curved and linear edges of the shape.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/smoothness
	Smoothness() float32

	// The color of the rounded rectangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/color
	SetColor(value ICIColor)

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/extent
	SetExtent(value corefoundation.CGRect)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/radius
	SetRadius(value float32)

	// A value to control the smoothness of the transition between the curved and linear edges of the shape.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/smoothness
	SetSmoothness(value float32)
}

// CIRoundedRectangleGeneratorObject wraps an existing Objective-C object that conforms to the CIRoundedRectangleGenerator protocol.
type CIRoundedRectangleGeneratorObject struct {
	objectivec.Object
}
func (o CIRoundedRectangleGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRoundedRectangleGeneratorObjectFromID constructs a [CIRoundedRectangleGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRoundedRectangleGeneratorObjectFromID(id objc.ID) CIRoundedRectangleGeneratorObject {
	return CIRoundedRectangleGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The color of the rounded rectangle.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/color
func (o CIRoundedRectangleGeneratorObject) Color() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// A rectangle that defines the extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/extent
func (o CIRoundedRectangleGeneratorObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/radius
func (o CIRoundedRectangleGeneratorObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A value to control the smoothness of the transition between the curved and
// linear edges of the shape.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleGenerator/smoothness
func (o CIRoundedRectangleGeneratorObject) Smoothness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("smoothness"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRoundedRectangleGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIRoundedRectangleGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIRoundedRectangleGeneratorObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIRoundedRectangleGeneratorObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIRoundedRectangleGeneratorObject) SetSmoothness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmoothness:"), value)
}

