// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a sunbeams generator filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator
type CISunbeamsGenerator interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the sunbeam pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/center
	Center() corefoundation.CGPoint

	// The color of the sun.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/color
	Color() ICIColor

	// The radius of the sunbeams.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/maxStriationRadius
	MaxStriationRadius() float32

	// The contrast of the sunbeams.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/striationContrast
	StriationContrast() float32

	// The intensity of the sunbeams.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/striationStrength
	StriationStrength() float32

	// The radius of the sun.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/sunRadius
	SunRadius() float32

	// The duration of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/time
	Time() float32

	// The x and y position to use as the center of the sunbeam pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/center
	SetCenter(value corefoundation.CGPoint)

	// The color of the sun.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/color
	SetColor(value ICIColor)

	// The radius of the sunbeams.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/maxStriationRadius
	SetMaxStriationRadius(value float32)

	// The contrast of the sunbeams.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/striationContrast
	SetStriationContrast(value float32)

	// The intensity of the sunbeams.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/striationStrength
	SetStriationStrength(value float32)

	// The radius of the sun.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/sunRadius
	SetSunRadius(value float32)

	// The duration of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/time
	SetTime(value float32)
}

// CISunbeamsGeneratorObject wraps an existing Objective-C object that conforms to the CISunbeamsGenerator protocol.
type CISunbeamsGeneratorObject struct {
	objectivec.Object
}
func (o CISunbeamsGeneratorObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISunbeamsGeneratorObjectFromID constructs a [CISunbeamsGeneratorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISunbeamsGeneratorObjectFromID(id objc.ID) CISunbeamsGeneratorObject {
	return CISunbeamsGeneratorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the sunbeam pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/center
func (o CISunbeamsGeneratorObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The color of the sun.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/color
func (o CISunbeamsGeneratorObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The radius of the sunbeams.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/maxStriationRadius
func (o CISunbeamsGeneratorObject) MaxStriationRadius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("maxStriationRadius"))
	return rv
	}
// The contrast of the sunbeams.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/striationContrast
func (o CISunbeamsGeneratorObject) StriationContrast() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("striationContrast"))
	return rv
	}
// The intensity of the sunbeams.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/striationStrength
func (o CISunbeamsGeneratorObject) StriationStrength() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("striationStrength"))
	return rv
	}
// The radius of the sun.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/sunRadius
func (o CISunbeamsGeneratorObject) SunRadius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("sunRadius"))
	return rv
	}
// The duration of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISunbeamsGenerator/time
func (o CISunbeamsGeneratorObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISunbeamsGeneratorObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISunbeamsGeneratorObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CISunbeamsGeneratorObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CISunbeamsGeneratorObject) SetMaxStriationRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaxStriationRadius:"), value)
}

func (o CISunbeamsGeneratorObject) SetStriationContrast(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStriationContrast:"), value)
}

func (o CISunbeamsGeneratorObject) SetStriationStrength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStriationStrength:"), value)
}

func (o CISunbeamsGeneratorObject) SetSunRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSunRadius:"), value)
}

func (o CISunbeamsGeneratorObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}

