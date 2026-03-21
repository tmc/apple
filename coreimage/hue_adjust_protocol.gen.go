// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a hue adjust filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust
type CIHueAdjust interface {
	objectivec.IObject
	CIFilterProtocol

	// An angle, in radians, to use to correct the hue of an image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust/angle
	Angle() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust/inputImage
	InputImage() ICIImage

	// An angle, in radians, to use to correct the hue of an image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust/angle
	SetAngle(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust/inputImage
	SetInputImage(value ICIImage)
}

// CIHueAdjustObject wraps an existing Objective-C object that conforms to the CIHueAdjust protocol.
type CIHueAdjustObject struct {
	objectivec.Object
}
func (o CIHueAdjustObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHueAdjustObjectFromID constructs a [CIHueAdjustObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHueAdjustObjectFromID(id objc.ID) CIHueAdjustObject {
	return CIHueAdjustObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An angle, in radians, to use to correct the hue of an image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust/angle
func (o CIHueAdjustObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHueAdjust/inputImage
func (o CIHueAdjustObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHueAdjustObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHueAdjustObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIHueAdjustObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

