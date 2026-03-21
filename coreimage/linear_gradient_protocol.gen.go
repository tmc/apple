// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a linear gradient filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient
type CILinearGradient interface {
	objectivec.IObject
	CIFilterProtocol

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/color0
	Color0() ICIColor

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/color1
	Color1() ICIColor

	// The starting position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/point0
	Point0() corefoundation.CGPoint

	// The ending position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/point1
	Point1() corefoundation.CGPoint

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/color0
	SetColor0(value ICIColor)

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/color1
	SetColor1(value ICIColor)

	// The starting position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/point0
	SetPoint0(value corefoundation.CGPoint)

	// The ending position of the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/point1
	SetPoint1(value corefoundation.CGPoint)
}

// CILinearGradientObject wraps an existing Objective-C object that conforms to the CILinearGradient protocol.
type CILinearGradientObject struct {
	objectivec.Object
}
func (o CILinearGradientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILinearGradientObjectFromID constructs a [CILinearGradientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILinearGradientObjectFromID(id objc.ID) CILinearGradientObject {
	return CILinearGradientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The first color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/color0
func (o CILinearGradientObject) Color0() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
	}
// The second color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/color1
func (o CILinearGradientObject) Color1() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
	}
// The starting position of the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/point0
func (o CILinearGradientObject) Point0() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point0"))
	return rv
	}
// The ending position of the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CILinearGradient/point1
func (o CILinearGradientObject) Point1() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point1"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILinearGradientObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CILinearGradientObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

func (o CILinearGradientObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

func (o CILinearGradientObject) SetPoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint0:"), value)
}

func (o CILinearGradientObject) SetPoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint1:"), value)
}

