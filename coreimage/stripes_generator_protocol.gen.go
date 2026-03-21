// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a stripes generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator
type CIStripesGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the stripe pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/center
	Center() corefoundation.CGPoint

	// A color to use for the odd stripes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/color0
	Color0() ICIColor

	// A color to use for the even stripes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/color1
	Color1() ICIColor

	// The sharpness of the stripe pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/sharpness
	Sharpness() float32

	// The width of a stripe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/width
	Width() float32

	// The x and y position to use as the center of the stripe pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/center
	SetCenter(value corefoundation.CGPoint)

	// A color to use for the odd stripes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/color0
	SetColor0(value ICIColor)

	// A color to use for the even stripes.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/color1
	SetColor1(value ICIColor)

	// The sharpness of the stripe pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/sharpness
	SetSharpness(value float32)

	// The width of a stripe.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/width
	SetWidth(value float32)
}

// CIStripesGeneratorObject wraps an existing Objective-C object that conforms to the CIStripesGenerator protocol.
type CIStripesGeneratorObject struct {
	objectivec.Object
}
func (o CIStripesGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIStripesGeneratorObjectFromID constructs a [CIStripesGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIStripesGeneratorObjectFromID(id objc.ID) CIStripesGeneratorObject {
	return CIStripesGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the stripe pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/center
func (o CIStripesGeneratorObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// A color to use for the odd stripes.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/color0
func (o CIStripesGeneratorObject) Color0() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color0"))
	return CIColorFromID(rv)
	}
// A color to use for the even stripes.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/color1
func (o CIStripesGeneratorObject) Color1() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color1"))
	return CIColorFromID(rv)
	}
// The sharpness of the stripe pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/sharpness
func (o CIStripesGeneratorObject) Sharpness() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// The width of a stripe.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStripesGenerator/width
func (o CIStripesGeneratorObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIStripesGeneratorObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIStripesGeneratorObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIStripesGeneratorObject) SetColor0(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor0:"), value)
}

func (o CIStripesGeneratorObject) SetColor1(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor1:"), value)
}

func (o CIStripesGeneratorObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

func (o CIStripesGeneratorObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

