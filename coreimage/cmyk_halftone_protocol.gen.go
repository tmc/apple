// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a CMYK halftone filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone
type CICMYKHalftone interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/angle
	Angle() float32

	// The x and y position to use as the center of the halftone pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/center
	Center() corefoundation.CGPoint

	// The gray component replacement value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/grayComponentReplacement
	GrayComponentReplacement() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/inputImage
	InputImage() ICIImage

	// The sharpness of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/sharpness
	Sharpness() float32

	// The under color removal value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/underColorRemoval
	UnderColorRemoval() float32

	// The distance between dots in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/width
	Width() float32

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the halftone pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/center
	SetCenter(value corefoundation.CGPoint)

	// The gray component replacement value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/grayComponentReplacement
	SetGrayComponentReplacement(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/inputImage
	SetInputImage(value ICIImage)

	// The sharpness of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/sharpness
	SetSharpness(value float32)

	// The under color removal value.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/underColorRemoval
	SetUnderColorRemoval(value float32)

	// The distance between dots in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/width
	SetWidth(value float32)
}

// CICMYKHalftoneObject wraps an existing Objective-C object that conforms to the CICMYKHalftone protocol.
type CICMYKHalftoneObject struct {
	objectivec.Object
}
func (o CICMYKHalftoneObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICMYKHalftoneObjectFromID constructs a [CICMYKHalftoneObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICMYKHalftoneObjectFromID(id objc.ID) CICMYKHalftoneObject {
	return CICMYKHalftoneObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/angle
func (o CICMYKHalftoneObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the halftone pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/center
func (o CICMYKHalftoneObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The gray component replacement value.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/grayComponentReplacement
func (o CICMYKHalftoneObject) GrayComponentReplacement() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("grayComponentReplacement"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/inputImage
func (o CICMYKHalftoneObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The sharpness of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/sharpness
func (o CICMYKHalftoneObject) Sharpness() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// The under color removal value.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/underColorRemoval
func (o CICMYKHalftoneObject) UnderColorRemoval() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("underColorRemoval"))
	return rv
	}
// The distance between dots in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICMYKHalftone/width
func (o CICMYKHalftoneObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICMYKHalftoneObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CICMYKHalftoneObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CICMYKHalftoneObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CICMYKHalftoneObject) SetGrayComponentReplacement(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setGrayComponentReplacement:"), value)
}

func (o CICMYKHalftoneObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CICMYKHalftoneObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

func (o CICMYKHalftoneObject) SetUnderColorRemoval(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setUnderColorRemoval:"), value)
}

func (o CICMYKHalftoneObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

