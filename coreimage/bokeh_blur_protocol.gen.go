// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a bokeh blur filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur
type CIBokehBlur interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/inputImage
	InputImage() ICIImage

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/radius
	Radius() float32

	// The amount of extra emphasis at the ring of the bokeh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringAmount
	RingAmount() float32

	// The radius of the extra emphasis at the ring of the bokeh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringSize
	RingSize() float32

	// The softness of the bokeh effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/softness
	Softness() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/inputImage
	SetInputImage(value ICIImage)

	// The radius of the blur, in pixels.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/radius
	SetRadius(value float32)

	// The amount of extra emphasis at the ring of the bokeh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringAmount
	SetRingAmount(value float32)

	// The radius of the extra emphasis at the ring of the bokeh.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringSize
	SetRingSize(value float32)

	// The softness of the bokeh effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/softness
	SetSoftness(value float32)
}

// CIBokehBlurObject wraps an existing Objective-C object that conforms to the CIBokehBlur protocol.
type CIBokehBlurObject struct {
	objectivec.Object
}

func (o CIBokehBlurObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIBokehBlurObjectFromID constructs a [CIBokehBlurObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIBokehBlurObjectFromID(id objc.ID) CIBokehBlurObject {
	return CIBokehBlurObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/inputImage
func (o CIBokehBlurObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The radius of the blur, in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/radius
func (o CIBokehBlurObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// The amount of extra emphasis at the ring of the bokeh.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringAmount
func (o CIBokehBlurObject) RingAmount() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("ringAmount"))
	return rv
}

// The radius of the extra emphasis at the ring of the bokeh.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringSize
func (o CIBokehBlurObject) RingSize() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("ringSize"))
	return rv
}

// The softness of the bokeh effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/softness
func (o CIBokehBlurObject) Softness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("softness"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIBokehBlurObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/inputImage
func (o CIBokehBlurObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The radius of the blur, in pixels.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/radius
func (o CIBokehBlurObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

// The amount of extra emphasis at the ring of the bokeh.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringAmount
func (o CIBokehBlurObject) SetRingAmount(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRingAmount:"), value)
}

// The radius of the extra emphasis at the ring of the bokeh.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/ringSize
func (o CIBokehBlurObject) SetRingSize(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRingSize:"), value)
}

// The softness of the bokeh effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIBokehBlur/softness
func (o CIBokehBlurObject) SetSoftness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSoftness:"), value)
}
