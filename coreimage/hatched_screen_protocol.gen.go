// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a hatched screen filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen
type CIHatchedScreen interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/angle
	Angle() float32

	// The x and y position to use as the center of the hatched screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/inputImage
	InputImage() ICIImage

	// The amount of sharpening to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/sharpness
	Sharpness() float32

	// The distance between lines in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/width
	Width() float32

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the hatched screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/inputImage
	SetInputImage(value ICIImage)

	// The amount of sharpening to apply.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/sharpness
	SetSharpness(value float32)

	// The distance between lines in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/width
	SetWidth(value float32)
}

// CIHatchedScreenObject wraps an existing Objective-C object that conforms to the CIHatchedScreen protocol.
type CIHatchedScreenObject struct {
	objectivec.Object
}
func (o CIHatchedScreenObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIHatchedScreenObjectFromID constructs a [CIHatchedScreenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIHatchedScreenObjectFromID(id objc.ID) CIHatchedScreenObject {
	return CIHatchedScreenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/angle
func (o CIHatchedScreenObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the hatched screen pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/center
func (o CIHatchedScreenObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/inputImage
func (o CIHatchedScreenObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The amount of sharpening to apply.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/sharpness
func (o CIHatchedScreenObject) Sharpness() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// The distance between lines in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIHatchedScreen/width
func (o CIHatchedScreenObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIHatchedScreenObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIHatchedScreenObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIHatchedScreenObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIHatchedScreenObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIHatchedScreenObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

func (o CIHatchedScreenObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

