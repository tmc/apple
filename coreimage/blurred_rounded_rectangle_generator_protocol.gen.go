// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The protocol for the Blurred Rounded Rectangle Generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator
type CIBlurredRoundedRectangleGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// A color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/color
	Color() ICIColor

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/extent
	Extent() corefoundation.CGRect

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/radius
	Radius() float32

	// The sigma for a gaussian blur.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/sigma
	Sigma() float32

	// A value to control the smoothness of the transition between the curved and linear edges of the shape.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/smoothness
	Smoothness() float32

	// A color.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/color
	SetColor(value ICIColor)

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/extent
	SetExtent(value corefoundation.CGRect)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/radius
	SetRadius(value float32)

	// The sigma for a gaussian blur.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/sigma
	SetSigma(value float32)

	// A value to control the smoothness of the transition between the curved and linear edges of the shape.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/smoothness
	SetSmoothness(value float32)
}

// CIBlurredRoundedRectangleGeneratorObject wraps an existing Objective-C object that conforms to the CIBlurredRoundedRectangleGenerator protocol.
type CIBlurredRoundedRectangleGeneratorObject struct {
	objectivec.Object
}
func (o CIBlurredRoundedRectangleGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBlurredRoundedRectangleGeneratorObjectFromID constructs a [CIBlurredRoundedRectangleGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBlurredRoundedRectangleGeneratorObjectFromID(id objc.ID) CIBlurredRoundedRectangleGeneratorObject {
	return CIBlurredRoundedRectangleGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A color.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/color
func (o CIBlurredRoundedRectangleGeneratorObject) Color() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// A rectangle that defines the extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/extent
func (o CIBlurredRoundedRectangleGeneratorObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/radius
func (o CIBlurredRoundedRectangleGeneratorObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// The sigma for a gaussian blur.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/sigma
func (o CIBlurredRoundedRectangleGeneratorObject) Sigma() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sigma"))
	return rv
	}
// A value to control the smoothness of the transition between the curved and
// linear edges of the shape.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlurredRoundedRectangleGenerator/smoothness
func (o CIBlurredRoundedRectangleGeneratorObject) Smoothness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("smoothness"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBlurredRoundedRectangleGeneratorObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIBlurredRoundedRectangleGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIBlurredRoundedRectangleGeneratorObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIBlurredRoundedRectangleGeneratorObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIBlurredRoundedRectangleGeneratorObject) SetSigma(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSigma:"), value)
}

func (o CIBlurredRoundedRectangleGeneratorObject) SetSmoothness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSmoothness:"), value)
}

