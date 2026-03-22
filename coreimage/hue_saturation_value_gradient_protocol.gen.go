// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a hue-saturation-value gradient filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient
type CIHueSaturationValueGradient interface {
	objectivec.IObject
	CIFilterProtocol

	// The color space for the generated color wheel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/colorSpace
	ColorSpace() coregraphics.CGColorSpaceRef

	// A Boolean value specifying whether the dither the generated output.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/dither
	Dither() float32

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/radius
	Radius() float32

	// The softness of the generated color wheel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/softness
	Softness() float32

	// The lightness of the hue-saturation gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/value
	Value() float32

	// The color space for the generated color wheel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/colorSpace
	SetColorSpace(value coregraphics.CGColorSpaceRef)

	// A Boolean value specifying whether the dither the generated output.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/dither
	SetDither(value float32)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/radius
	SetRadius(value float32)

	// The softness of the generated color wheel.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/softness
	SetSoftness(value float32)

	// The lightness of the hue-saturation gradient.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/value
	SetValue(value float32)
}

// CIHueSaturationValueGradientObject wraps an existing Objective-C object that conforms to the CIHueSaturationValueGradient protocol.
type CIHueSaturationValueGradientObject struct {
	objectivec.Object
}
func (o CIHueSaturationValueGradientObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHueSaturationValueGradientObjectFromID constructs a [CIHueSaturationValueGradientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHueSaturationValueGradientObjectFromID(id objc.ID) CIHueSaturationValueGradientObject {
	return CIHueSaturationValueGradientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The color space for the generated color wheel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/colorSpace
func (o CIHueSaturationValueGradientObject) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](o.ID, objc.Sel("colorSpace"))
	return rv
	}
// A Boolean value specifying whether the dither the generated output.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/dither
func (o CIHueSaturationValueGradientObject) Dither() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("dither"))
	return rv
	}
// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/radius
func (o CIHueSaturationValueGradientObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// The softness of the generated color wheel.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/softness
func (o CIHueSaturationValueGradientObject) Softness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("softness"))
	return rv
	}
// The lightness of the hue-saturation gradient.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueSaturationValueGradient/value
func (o CIHueSaturationValueGradientObject) Value() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("value"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHueSaturationValueGradientObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHueSaturationValueGradientObject) SetColorSpace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorSpace:"), value)
}

func (o CIHueSaturationValueGradientObject) SetDither(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setDither:"), value)
}

func (o CIHueSaturationValueGradientObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

func (o CIHueSaturationValueGradientObject) SetSoftness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSoftness:"), value)
}

func (o CIHueSaturationValueGradientObject) SetValue(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setValue:"), value)
}

