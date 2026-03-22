// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a line screen filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineScreen
type CILineScreen interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/angle
	Angle() float32

	// The x and y position to use as the center of the line screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/inputImage
	InputImage() ICIImage

	// The sharpness of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/sharpness
	Sharpness() float32

	// The distance between lines in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/width
	Width() float32

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the line screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/inputImage
	SetInputImage(value ICIImage)

	// The sharpness of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/sharpness
	SetSharpness(value float32)

	// The distance between lines in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/width
	SetWidth(value float32)
}

// CILineScreenObject wraps an existing Objective-C object that conforms to the CILineScreen protocol.
type CILineScreenObject struct {
	objectivec.Object
}
func (o CILineScreenObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILineScreenObjectFromID constructs a [CILineScreenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILineScreenObjectFromID(id objc.ID) CILineScreenObject {
	return CILineScreenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/angle
func (o CILineScreenObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the line screen pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/center
func (o CILineScreenObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/inputImage
func (o CILineScreenObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The sharpness of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/sharpness
func (o CILineScreenObject) Sharpness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// The distance between lines in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CILineScreen/width
func (o CILineScreenObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILineScreenObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CILineScreenObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CILineScreenObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CILineScreenObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CILineScreenObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

func (o CILineScreenObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

