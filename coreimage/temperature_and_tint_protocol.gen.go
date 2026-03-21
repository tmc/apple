// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a temperature and tint filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint
type CITemperatureAndTint interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/inputImage
	InputImage() ICIImage

	// A vector containing the source white point defined by color temperature and tint.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/neutral
	Neutral() ICIVector

	// A vector containing the desired white point defined by color temperature and tint.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/targetNeutral
	TargetNeutral() ICIVector

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/inputImage
	SetInputImage(value ICIImage)

	// A vector containing the source white point defined by color temperature and tint.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/neutral
	SetNeutral(value ICIVector)

	// A vector containing the desired white point defined by color temperature and tint.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/targetNeutral
	SetTargetNeutral(value ICIVector)
}

// CITemperatureAndTintObject wraps an existing Objective-C object that conforms to the CITemperatureAndTint protocol.
type CITemperatureAndTintObject struct {
	objectivec.Object
}
func (o CITemperatureAndTintObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITemperatureAndTintObjectFromID constructs a [CITemperatureAndTintObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITemperatureAndTintObjectFromID(id objc.ID) CITemperatureAndTintObject {
	return CITemperatureAndTintObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/inputImage
func (o CITemperatureAndTintObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A vector containing the source white point defined by color temperature and
// tint.
//
// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/neutral
func (o CITemperatureAndTintObject) Neutral() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("neutral"))
	return CIVectorFromID(rv)
	}
// A vector containing the desired white point defined by color temperature
// and tint.
//
// See: https://developer.apple.com/documentation/CoreImage/CITemperatureAndTint/targetNeutral
func (o CITemperatureAndTintObject) TargetNeutral() ICIVector {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetNeutral"))
	return CIVectorFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITemperatureAndTintObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITemperatureAndTintObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CITemperatureAndTintObject) SetNeutral(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setNeutral:"), value)
}

func (o CITemperatureAndTintObject) SetTargetNeutral(value ICIVector) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetNeutral:"), value)
}

