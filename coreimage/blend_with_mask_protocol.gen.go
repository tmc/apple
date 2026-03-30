// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a blend with mask filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask
type CIBlendWithMask interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/backgroundImage
	BackgroundImage() ICIImage

	// The image to use as a foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/inputImage
	InputImage() ICIImage

	// A grayscale mask that defines the blend.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/maskImage
	MaskImage() ICIImage

	// The image to use as a background image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/backgroundImage
	SetBackgroundImage(value ICIImage)

	// The image to use as a foreground image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/inputImage
	SetInputImage(value ICIImage)

	// A grayscale mask that defines the blend.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/maskImage
	SetMaskImage(value ICIImage)
}

// CIBlendWithMaskObject wraps an existing Objective-C object that conforms to the CIBlendWithMask protocol.
type CIBlendWithMaskObject struct {
	objectivec.Object
}

func (o CIBlendWithMaskObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBlendWithMaskObjectFromID constructs a [CIBlendWithMaskObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBlendWithMaskObjectFromID(id objc.ID) CIBlendWithMaskObject {
	return CIBlendWithMaskObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as a background image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/backgroundImage
func (o CIBlendWithMaskObject) BackgroundImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("backgroundImage"))
	return CIImageFromID(rv)
}

// The image to use as a foreground image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/inputImage
func (o CIBlendWithMaskObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A grayscale mask that defines the blend.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/maskImage
func (o CIBlendWithMaskObject) MaskImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("maskImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBlendWithMaskObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as a background image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/backgroundImage
func (o CIBlendWithMaskObject) SetBackgroundImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setBackgroundImage:"), value)
}

// The image to use as a foreground image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/inputImage
func (o CIBlendWithMaskObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// A grayscale mask that defines the blend.
//
// # Discussion
//
// When a mask value is 0.0, the result is the background. When the mask value
// is 1.0, the result is the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBlendWithMask/maskImage
func (o CIBlendWithMaskObject) SetMaskImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setMaskImage:"), value)
}
