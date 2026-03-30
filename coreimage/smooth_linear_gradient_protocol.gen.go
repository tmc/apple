// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a smooth linear gradient filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient
type CISmoothLinearGradient interface {
	objectivec.IObject
	CIFilterProtocol

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color0
	Color0() ICIColor

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color1
	Color1() ICIColor

	// The starting position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point0
	Point0() corefoundation.CGPoint

	// The ending position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point1
	Point1() corefoundation.CGPoint

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color0
	SetColor0(value ICIColor)

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color1
	SetColor1(value ICIColor)

	// The starting position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point0
	SetPoint0(value corefoundation.CGPoint)

	// The ending position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point1
	SetPoint1(value corefoundation.CGPoint)
}

// CISmoothLinearGradientObject wraps an existing Objective-C object that conforms to the CISmoothLinearGradient protocol.
type CISmoothLinearGradientObject struct {
	objectivec.Object
}

func (o CISmoothLinearGradientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISmoothLinearGradientObjectFromID constructs a [CISmoothLinearGradientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISmoothLinearGradientObjectFromID(id objc.ID) CISmoothLinearGradientObject {
	return CISmoothLinearGradientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The first color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color0
func (o CISmoothLinearGradientObject) Color0() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
}

// The second color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color1
func (o CISmoothLinearGradientObject) Color1() ICIColor {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
}

// The starting position of the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point0
func (o CISmoothLinearGradientObject) Point0() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point0"))
	return rv
}

// The ending position of the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point1
func (o CISmoothLinearGradientObject) Point1() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point1"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISmoothLinearGradientObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The first color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color0
func (o CISmoothLinearGradientObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

// The second color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/color1
func (o CISmoothLinearGradientObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

// The starting position of the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point0
func (o CISmoothLinearGradientObject) SetPoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint0:"), value)
}

// The ending position of the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CISmoothLinearGradient/point1
func (o CISmoothLinearGradientObject) SetPoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint1:"), value)
}
