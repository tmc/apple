// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIRoundedRectangleStrokeGenerator protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator
type CIRoundedRectangleStrokeGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// Color protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/color
	Color() ICIColor

	// Extent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/extent
	Extent() corefoundation.CGRect

	// Radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/radius
	Radius() float32

	// Width protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/width
	Width() float32

	// A value to control the smoothness of the transition between the curved and linear edges of the shape.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/smoothness
	Smoothness() float32

	// color protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/color
	SetColor(value ICIColor)

	// extent protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/extent
	SetExtent(value corefoundation.CGRect)

	// radius protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/radius
	SetRadius(value float32)

	// width protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/width
	SetWidth(value float32)

	// A value to control the smoothness of the transition between the curved and linear edges of the shape.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/smoothness
	SetSmoothness(value float32)
}

// CIRoundedRectangleStrokeGeneratorObject wraps an existing Objective-C object that conforms to the CIRoundedRectangleStrokeGenerator protocol.
type CIRoundedRectangleStrokeGeneratorObject struct {
	objectivec.Object
}
func (o CIRoundedRectangleStrokeGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRoundedRectangleStrokeGeneratorObjectFromID constructs a [CIRoundedRectangleStrokeGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRoundedRectangleStrokeGeneratorObjectFromID(id objc.ID) CIRoundedRectangleStrokeGeneratorObject {
	return CIRoundedRectangleStrokeGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/color
func (o CIRoundedRectangleStrokeGeneratorObject) Color() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/extent
func (o CIRoundedRectangleStrokeGeneratorObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/radius
func (o CIRoundedRectangleStrokeGeneratorObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/width
func (o CIRoundedRectangleStrokeGeneratorObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A value to control the smoothness of the transition between the curved and
// linear edges of the shape.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRoundedRectangleStrokeGenerator/smoothness
func (o CIRoundedRectangleStrokeGeneratorObject) Smoothness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("smoothness"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRoundedRectangleStrokeGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIRoundedRectangleStrokeGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIRoundedRectangleStrokeGeneratorObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIRoundedRectangleStrokeGeneratorObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIRoundedRectangleStrokeGeneratorObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

func (o CIRoundedRectangleStrokeGeneratorObject) SetSmoothness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmoothness:"), value)
}

