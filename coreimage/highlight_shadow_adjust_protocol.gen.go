// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a highlight-shadow adjust filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust
type CIHighlightShadowAdjust interface {
	objectivec.IObject
	CIFilterProtocol

	// The amount of adjustment to the highlights in the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/highlightAmount
	HighlightAmount() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/inputImage
	InputImage() ICIImage

	// The shadow highlight radius.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/radius
	Radius() float32

	// The amount of adjustment to the shadows in the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/shadowAmount
	ShadowAmount() float32

	// The amount of adjustment to the highlights in the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/highlightAmount
	SetHighlightAmount(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/inputImage
	SetInputImage(value ICIImage)

	// The shadow highlight radius.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/radius
	SetRadius(value float32)

	// The amount of adjustment to the shadows in the image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/shadowAmount
	SetShadowAmount(value float32)
}

// CIHighlightShadowAdjustObject wraps an existing Objective-C object that conforms to the CIHighlightShadowAdjust protocol.
type CIHighlightShadowAdjustObject struct {
	objectivec.Object
}

func (o CIHighlightShadowAdjustObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHighlightShadowAdjustObjectFromID constructs a [CIHighlightShadowAdjustObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHighlightShadowAdjustObjectFromID(id objc.ID) CIHighlightShadowAdjustObject {
	return CIHighlightShadowAdjustObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The amount of adjustment to the highlights in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/highlightAmount
func (o CIHighlightShadowAdjustObject) HighlightAmount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("highlightAmount"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/inputImage
func (o CIHighlightShadowAdjustObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The shadow highlight radius.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/radius
func (o CIHighlightShadowAdjustObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// The amount of adjustment to the shadows in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/shadowAmount
func (o CIHighlightShadowAdjustObject) ShadowAmount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("shadowAmount"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHighlightShadowAdjustObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The amount of adjustment to the highlights in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/highlightAmount
func (o CIHighlightShadowAdjustObject) SetHighlightAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setHighlightAmount:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/inputImage
func (o CIHighlightShadowAdjustObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The shadow highlight radius.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/radius
func (o CIHighlightShadowAdjustObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

// The amount of adjustment to the shadows in the image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHighlightShadowAdjust/shadowAmount
func (o CIHighlightShadowAdjustObject) SetShadowAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setShadowAmount:"), value)
}
