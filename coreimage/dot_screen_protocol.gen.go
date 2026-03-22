// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a dot screen filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen
type CIDotScreen interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/angle
	Angle() float32

	// The x and y position to use as the center of the dot screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/inputImage
	InputImage() ICIImage

	// The sharpness of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/sharpness
	Sharpness() float32

	// The distance between dots in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/width
	Width() float32

	// The angle of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the dot screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/inputImage
	SetInputImage(value ICIImage)

	// The sharpness of the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/sharpness
	SetSharpness(value float32)

	// The distance between dots in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/width
	SetWidth(value float32)
}

// CIDotScreenObject wraps an existing Objective-C object that conforms to the CIDotScreen protocol.
type CIDotScreenObject struct {
	objectivec.Object
}
func (o CIDotScreenObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIDotScreenObjectFromID constructs a [CIDotScreenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIDotScreenObjectFromID(id objc.ID) CIDotScreenObject {
	return CIDotScreenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/angle
func (o CIDotScreenObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the dot screen pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/center
func (o CIDotScreenObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/inputImage
func (o CIDotScreenObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The sharpness of the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/sharpness
func (o CIDotScreenObject) Sharpness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
	}
// The distance between dots in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIDotScreen/width
func (o CIDotScreenObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIDotScreenObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIDotScreenObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIDotScreenObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIDotScreenObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIDotScreenObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

func (o CIDotScreenObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

