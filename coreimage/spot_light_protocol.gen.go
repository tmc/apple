// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a spotlight filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight
type CISpotLight interface {
	objectivec.IObject
	CIFilterProtocol

	// The brightness of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/brightness
	Brightness() float32

	// The color of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/color
	Color() ICIColor

	// The size of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/concentration
	Concentration() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/inputImage
	InputImage() ICIImage

	// The x and y position that the spotlight points at.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/lightPointsAt
	LightPointsAt() ICIVector

	// The x and y position of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/lightPosition
	LightPosition() ICIVector

	// The brightness of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/brightness
	SetBrightness(value float32)

	// The color of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/color
	SetColor(value ICIColor)

	// The size of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/concentration
	SetConcentration(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/inputImage
	SetInputImage(value ICIImage)

	// The x and y position that the spotlight points at.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/lightPointsAt
	SetLightPointsAt(value ICIVector)

	// The x and y position of the spotlight.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/lightPosition
	SetLightPosition(value ICIVector)
}

// CISpotLightObject wraps an existing Objective-C object that conforms to the CISpotLight protocol.
type CISpotLightObject struct {
	objectivec.Object
}
func (o CISpotLightObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISpotLightObjectFromID constructs a [CISpotLightObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISpotLightObjectFromID(id objc.ID) CISpotLightObject {
	return CISpotLightObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The brightness of the spotlight.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/brightness
func (o CISpotLightObject) Brightness() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("brightness"))
	return rv
	}
// The color of the spotlight.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/color
func (o CISpotLightObject) Color() ICIColor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("color"))
	return CIColorFromID(rv)
	}
// The size of the spotlight.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/concentration
func (o CISpotLightObject) Concentration() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("concentration"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/inputImage
func (o CISpotLightObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The x and y position that the spotlight points at.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/lightPointsAt
func (o CISpotLightObject) LightPointsAt() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("lightPointsAt"))
	return CIVectorFromID(rv)
	}
// The x and y position of the spotlight.
//
// See: https://developer.apple.com/documentation/CoreImage/CISpotLight/lightPosition
func (o CISpotLightObject) LightPosition() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("lightPosition"))
	return CIVectorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISpotLightObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISpotLightObject) SetBrightness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setBrightness:"), value)
}

func (o CISpotLightObject) SetColor(value ICIColor) {
	objc.Send[struct{}](o.ID, objc.Sel("setColor:"), value)
}

func (o CISpotLightObject) SetConcentration(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setConcentration:"), value)
}

func (o CISpotLightObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISpotLightObject) SetLightPointsAt(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setLightPointsAt:"), value)
}

func (o CISpotLightObject) SetLightPosition(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setLightPosition:"), value)
}

