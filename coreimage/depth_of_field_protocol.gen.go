// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a depth-of-field filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField
type CIDepthOfField interface {
	objectivec.IObject
	CIFilterProtocol

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/inputImage
	InputImage() ICIImage

	// The first point in the focused region of the output image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point0
	Point0() corefoundation.CGPoint

	// The second point in the focused region of the output image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point1
	Point1() corefoundation.CGPoint

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/radius
	Radius() float32

	// The amount to adjust the saturation by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/saturation
	Saturation() float32

	// The intensity of the unsharp mask effect applied to the in-focus area.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskIntensity
	UnsharpMaskIntensity() float32

	// The radius of the unsharp mask effect applied to the in-focus area.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskRadius
	UnsharpMaskRadius() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/inputImage
	SetInputImage(value ICIImage)

	// The first point in the focused region of the output image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point0
	SetPoint0(value corefoundation.CGPoint)

	// The second point in the focused region of the output image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point1
	SetPoint1(value corefoundation.CGPoint)

	// The distance from the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/radius
	SetRadius(value float32)

	// The amount to adjust the saturation by.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/saturation
	SetSaturation(value float32)

	// The intensity of the unsharp mask effect applied to the in-focus area.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskIntensity
	SetUnsharpMaskIntensity(value float32)

	// The radius of the unsharp mask effect applied to the in-focus area.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskRadius
	SetUnsharpMaskRadius(value float32)
}

// CIDepthOfFieldObject wraps an existing Objective-C object that conforms to the CIDepthOfField protocol.
type CIDepthOfFieldObject struct {
	objectivec.Object
}

func (o CIDepthOfFieldObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDepthOfFieldObjectFromID constructs a [CIDepthOfFieldObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDepthOfFieldObjectFromID(id objc.ID) CIDepthOfFieldObject {
	return CIDepthOfFieldObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/inputImage
func (o CIDepthOfFieldObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The first point in the focused region of the output image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point0
func (o CIDepthOfFieldObject) Point0() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point0"))
	return rv
}

// The second point in the focused region of the output image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point1
func (o CIDepthOfFieldObject) Point1() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point1"))
	return rv
}

// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/radius
func (o CIDepthOfFieldObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// The amount to adjust the saturation by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/saturation
func (o CIDepthOfFieldObject) Saturation() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("saturation"))
	return rv
}

// The intensity of the unsharp mask effect applied to the in-focus area.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskIntensity
func (o CIDepthOfFieldObject) UnsharpMaskIntensity() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("unsharpMaskIntensity"))
	return rv
}

// The radius of the unsharp mask effect applied to the in-focus area.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskRadius
func (o CIDepthOfFieldObject) UnsharpMaskRadius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("unsharpMaskRadius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDepthOfFieldObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/inputImage
func (o CIDepthOfFieldObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The first point in the focused region of the output image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point0
func (o CIDepthOfFieldObject) SetPoint0(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint0:"), value)
}

// The second point in the focused region of the output image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/point1
func (o CIDepthOfFieldObject) SetPoint1(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint1:"), value)
}

// The distance from the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/radius
func (o CIDepthOfFieldObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

// The amount to adjust the saturation by.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/saturation
func (o CIDepthOfFieldObject) SetSaturation(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSaturation:"), value)
}

// The intensity of the unsharp mask effect applied to the in-focus area.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskIntensity
func (o CIDepthOfFieldObject) SetUnsharpMaskIntensity(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setUnsharpMaskIntensity:"), value)
}

// The radius of the unsharp mask effect applied to the in-focus area.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDepthOfField/unsharpMaskRadius
func (o CIDepthOfFieldObject) SetUnsharpMaskRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setUnsharpMaskRadius:"), value)
}
