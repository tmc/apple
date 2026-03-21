// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a flash transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition
type CIFlashTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/center
	Center() corefoundation.CGPoint

	// The color of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/color
	Color() ICIColor

	// The extent of the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/extent
	Extent() corefoundation.CGRect

	// The amount of fade between the flash and the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/fadeThreshold
	FadeThreshold() float32

	// The radius of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/maxStriationRadius
	MaxStriationRadius() float32

	// The contrast of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/striationContrast
	StriationContrast() float32

	// The strength of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/striationStrength
	StriationStrength() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/center
	SetCenter(value corefoundation.CGPoint)

	// The color of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/color
	SetColor(value ICIColor)

	// The extent of the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/extent
	SetExtent(value corefoundation.CGRect)

	// The amount of fade between the flash and the target image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/fadeThreshold
	SetFadeThreshold(value float32)

	// The radius of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/maxStriationRadius
	SetMaxStriationRadius(value float32)

	// The contrast of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/striationContrast
	SetStriationContrast(value float32)

	// The strength of the light rays emanating from the flash.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/striationStrength
	SetStriationStrength(value float32)
}

// CIFlashTransitionObject wraps an existing Objective-C object that conforms to the CIFlashTransition protocol.
type CIFlashTransitionObject struct {
	objectivec.Object
}
func (o CIFlashTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFlashTransitionObjectFromID constructs a [CIFlashTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFlashTransitionObjectFromID(id objc.ID) CIFlashTransitionObject {
	return CIFlashTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/center
func (o CIFlashTransitionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The color of the light rays emanating from the flash.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/color
func (o CIFlashTransitionObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The extent of the flash.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/extent
func (o CIFlashTransitionObject) Extent() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
	}
// The amount of fade between the flash and the target image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/fadeThreshold
func (o CIFlashTransitionObject) FadeThreshold() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("fadeThreshold"))
	return rv
	}
// The radius of the light rays emanating from the flash.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/maxStriationRadius
func (o CIFlashTransitionObject) MaxStriationRadius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("maxStriationRadius"))
	return rv
	}
// The contrast of the light rays emanating from the flash.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/striationContrast
func (o CIFlashTransitionObject) StriationContrast() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("striationContrast"))
	return rv
	}
// The strength of the light rays emanating from the flash.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFlashTransition/striationStrength
func (o CIFlashTransitionObject) StriationStrength() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("striationStrength"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFlashTransitionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIFlashTransitionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIFlashTransitionObject) TargetImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
	}
// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIFlashTransitionObject) Time() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
	}

func (o CIFlashTransitionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIFlashTransitionObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CIFlashTransitionObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

func (o CIFlashTransitionObject) SetFadeThreshold(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFadeThreshold:"), value)
}

func (o CIFlashTransitionObject) SetMaxStriationRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaxStriationRadius:"), value)
}

func (o CIFlashTransitionObject) SetStriationContrast(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStriationContrast:"), value)
}

func (o CIFlashTransitionObject) SetStriationStrength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setStriationStrength:"), value)
}

func (o CIFlashTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIFlashTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

func (o CIFlashTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}

