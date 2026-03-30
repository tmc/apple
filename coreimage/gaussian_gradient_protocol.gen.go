// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Gaussian gradient filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient
type CIGaussianGradient interface {
	objectivec.IObject
	CIFilterProtocol

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/center
	Center() corefoundation.CGPoint

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color0
	Color0() ICIColor

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color1
	Color1() ICIColor

	// The radius of the Gaussian distribution.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/radius
	Radius() float32

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/center
	SetCenter(value corefoundation.CGPoint)

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color0
	SetColor0(value ICIColor)

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color1
	SetColor1(value ICIColor)

	// The radius of the Gaussian distribution.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/radius
	SetRadius(value float32)
}

// CIGaussianGradientObject wraps an existing Objective-C object that conforms to the CIGaussianGradient protocol.
type CIGaussianGradientObject struct {
	objectivec.Object
}

func (o CIGaussianGradientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGaussianGradientObjectFromID constructs a [CIGaussianGradientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGaussianGradientObjectFromID(id objc.ID) CIGaussianGradientObject {
	return CIGaussianGradientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/center
func (o CIGaussianGradientObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The first color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color0
func (o CIGaussianGradientObject) Color0() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
}

// The second color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color1
func (o CIGaussianGradientObject) Color1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
}

// The radius of the Gaussian distribution.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/radius
func (o CIGaussianGradientObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGaussianGradientObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/center
func (o CIGaussianGradientObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The first color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color0
func (o CIGaussianGradientObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

// The second color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/color1
func (o CIGaussianGradientObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

// The radius of the Gaussian distribution.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGaussianGradient/radius
func (o CIGaussianGradientObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}
