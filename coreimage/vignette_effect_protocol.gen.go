// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a vignette-effect filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect
type CIVignetteEffect interface {
	objectivec.IObject
	CIFilterProtocol

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/center
	Center() corefoundation.CGPoint

	// The falloff of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/falloff
	Falloff() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/inputImage
	InputImage() ICIImage

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/intensity
	Intensity() float32

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/radius
	Radius() float32

	// The center of the effect as x and y coordinates.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/center
	SetCenter(value corefoundation.CGPoint)

	// The falloff of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/falloff
	SetFalloff(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/inputImage
	SetInputImage(value ICIImage)

	// The intensity of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/intensity
	SetIntensity(value float32)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/radius
	SetRadius(value float32)
}

// CIVignetteEffectObject wraps an existing Objective-C object that conforms to the CIVignetteEffect protocol.
type CIVignetteEffectObject struct {
	objectivec.Object
}
func (o CIVignetteEffectObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIVignetteEffectObjectFromID constructs a [CIVignetteEffectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIVignetteEffectObjectFromID(id objc.ID) CIVignetteEffectObject {
	return CIVignetteEffectObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The center of the effect as x and y coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/center
func (o CIVignetteEffectObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The falloff of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/falloff
func (o CIVignetteEffectObject) Falloff() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("falloff"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/inputImage
func (o CIVignetteEffectObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The intensity of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/intensity
func (o CIVignetteEffectObject) Intensity() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("intensity"))
	return rv
	}
// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIVignetteEffect/radius
func (o CIVignetteEffectObject) Radius() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIVignetteEffectObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIVignetteEffectObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIVignetteEffectObject) SetFalloff(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFalloff:"), value)
}

func (o CIVignetteEffectObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIVignetteEffectObject) SetIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setIntensity:"), value)
}

func (o CIVignetteEffectObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

