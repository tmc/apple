// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a lenticular halo generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator
type CILenticularHaloGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/center
	Center() corefoundation.CGPoint

	// The color of the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/color
	Color() ICIColor

	// The separation of colors in the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloOverlap
	HaloOverlap() float32

	// The radius of the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloRadius
	HaloRadius() float32

	// The width of the halo, from its inner radius to its outer radius.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloWidth
	HaloWidth() float32

	// The contrast of the halo colors.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/striationContrast
	StriationContrast() float32

	// The intensity of the halo colors.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/striationStrength
	StriationStrength() float32

	// The current time of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/time
	Time() float32

	// The x and y position to use as the center of the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/center
	SetCenter(value corefoundation.CGPoint)

	// The color of the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/color
	SetColor(value ICIColor)

	// The separation of colors in the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloOverlap
	SetHaloOverlap(value float32)

	// The radius of the halo.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloRadius
	SetHaloRadius(value float32)

	// The width of the halo, from its inner radius to its outer radius.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloWidth
	SetHaloWidth(value float32)

	// The contrast of the halo colors.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/striationContrast
	SetStriationContrast(value float32)

	// The intensity of the halo colors.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/striationStrength
	SetStriationStrength(value float32)

	// The current time of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/time
	SetTime(value float32)
}

// CILenticularHaloGeneratorObject wraps an existing Objective-C object that conforms to the CILenticularHaloGenerator protocol.
type CILenticularHaloGeneratorObject struct {
	objectivec.Object
}
func (o CILenticularHaloGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILenticularHaloGeneratorObjectFromID constructs a [CILenticularHaloGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILenticularHaloGeneratorObjectFromID(id objc.ID) CILenticularHaloGeneratorObject {
	return CILenticularHaloGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the halo.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/center
func (o CILenticularHaloGeneratorObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The color of the halo.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/color
func (o CILenticularHaloGeneratorObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The separation of colors in the halo.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloOverlap
func (o CILenticularHaloGeneratorObject) HaloOverlap() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("haloOverlap"))
	return rv
	}
// The radius of the halo.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloRadius
func (o CILenticularHaloGeneratorObject) HaloRadius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("haloRadius"))
	return rv
	}
// The width of the halo, from its inner radius to its outer radius.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/haloWidth
func (o CILenticularHaloGeneratorObject) HaloWidth() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("haloWidth"))
	return rv
	}
// The contrast of the halo colors.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/striationContrast
func (o CILenticularHaloGeneratorObject) StriationContrast() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("striationContrast"))
	return rv
	}
// The intensity of the halo colors.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/striationStrength
func (o CILenticularHaloGeneratorObject) StriationStrength() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("striationStrength"))
	return rv
	}
// The current time of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CILenticularHaloGenerator/time
func (o CILenticularHaloGeneratorObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILenticularHaloGeneratorObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CILenticularHaloGeneratorObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CILenticularHaloGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CILenticularHaloGeneratorObject) SetHaloOverlap(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHaloOverlap:"), value)
}

func (o CILenticularHaloGeneratorObject) SetHaloRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHaloRadius:"), value)
}

func (o CILenticularHaloGeneratorObject) SetHaloWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHaloWidth:"), value)
}

func (o CILenticularHaloGeneratorObject) SetStriationContrast(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStriationContrast:"), value)
}

func (o CILenticularHaloGeneratorObject) SetStriationStrength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStriationStrength:"), value)
}

func (o CILenticularHaloGeneratorObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}

