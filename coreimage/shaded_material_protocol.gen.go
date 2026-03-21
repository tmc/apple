// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a shaded material filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial
type CIShadedMaterial interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/inputImage
	InputImage() ICIImage

	// The scale of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/scale
	Scale() float32

	// The image to use as the height field.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/shadingImage
	ShadingImage() ICIImage

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/inputImage
	SetInputImage(value ICIImage)

	// The scale of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/scale
	SetScale(value float32)

	// The image to use as the height field.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/shadingImage
	SetShadingImage(value ICIImage)
}

// CIShadedMaterialObject wraps an existing Objective-C object that conforms to the CIShadedMaterial protocol.
type CIShadedMaterialObject struct {
	objectivec.Object
}
func (o CIShadedMaterialObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIShadedMaterialObjectFromID constructs a [CIShadedMaterialObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIShadedMaterialObjectFromID(id objc.ID) CIShadedMaterialObject {
	return CIShadedMaterialObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/inputImage
func (o CIShadedMaterialObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The scale of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/scale
func (o CIShadedMaterialObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// The image to use as the height field.
//
// See: https://developer.apple.com/documentation/CoreImage/CIShadedMaterial/shadingImage
func (o CIShadedMaterialObject) ShadingImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("shadingImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIShadedMaterialObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIShadedMaterialObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIShadedMaterialObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

func (o CIShadedMaterialObject) SetShadingImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadingImage:"), value)
}

