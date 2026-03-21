// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a radial gradient filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient
type CIRadialGradient interface {
	objectivec.IObject
	CIFilterProtocol

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/center
	Center() corefoundation.CGPoint

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/color0
	Color0() ICIColor

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/color1
	Color1() ICIColor

	// The radius of the starting circle to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/radius0
	Radius0() float32

	// The radius of the ending circle to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/radius1
	Radius1() float32

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/center
	SetCenter(value corefoundation.CGPoint)

	// The first color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/color0
	SetColor0(value ICIColor)

	// The second color to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/color1
	SetColor1(value ICIColor)

	// The radius of the starting circle to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/radius0
	SetRadius0(value float32)

	// The radius of the ending circle to use in the gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/radius1
	SetRadius1(value float32)
}

// CIRadialGradientObject wraps an existing Objective-C object that conforms to the CIRadialGradient protocol.
type CIRadialGradientObject struct {
	objectivec.Object
}
func (o CIRadialGradientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIRadialGradientObjectFromID constructs a [CIRadialGradientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIRadialGradientObjectFromID(id objc.ID) CIRadialGradientObject {
	return CIRadialGradientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/center
func (o CIRadialGradientObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The first color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/color0
func (o CIRadialGradientObject) Color0() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
	}
// The second color to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/color1
func (o CIRadialGradientObject) Color1() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
	}
// The radius of the starting circle to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/radius0
func (o CIRadialGradientObject) Radius0() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius0"))
	return rv
	}
// The radius of the ending circle to use in the gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIRadialGradient/radius1
func (o CIRadialGradientObject) Radius1() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius1"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIRadialGradientObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIRadialGradientObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIRadialGradientObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

func (o CIRadialGradientObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

func (o CIRadialGradientObject) SetRadius0(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius0:"), value)
}

func (o CIRadialGradientObject) SetRadius1(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius1:"), value)
}

