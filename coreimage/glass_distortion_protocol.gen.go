// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// CIGlassDistortion protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion
type CIGlassDistortion interface {
	objectivec.IObject
	CIFilterProtocol

	// Center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/inputImage
	InputImage() ICIImage

	// Scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/scale
	Scale() float32

	// TextureImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/textureImage
	TextureImage() ICIImage

	// center protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/inputImage
	SetInputImage(value ICIImage)

	// scale protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/scale
	SetScale(value float32)

	// textureImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/textureImage
	SetTextureImage(value ICIImage)
}

// CIGlassDistortionObject wraps an existing Objective-C object that conforms to the CIGlassDistortion protocol.
type CIGlassDistortionObject struct {
	objectivec.Object
}
func (o CIGlassDistortionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIGlassDistortionObjectFromID constructs a [CIGlassDistortionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIGlassDistortionObjectFromID(id objc.ID) CIGlassDistortionObject {
	return CIGlassDistortionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/center
func (o CIGlassDistortionObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/inputImage
func (o CIGlassDistortionObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/scale
func (o CIGlassDistortionObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIGlassDistortion/textureImage
func (o CIGlassDistortionObject) TextureImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textureImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIGlassDistortionObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIGlassDistortionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIGlassDistortionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIGlassDistortionObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

func (o CIGlassDistortionObject) SetTextureImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTextureImage:"), value)
}

