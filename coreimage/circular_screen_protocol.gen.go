// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a circular screen filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen
type CICircularScreen interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the circular screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/inputImage
	InputImage() ICIImage

	// The sharpness of the circles.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/sharpness
	Sharpness() float32

	// The distance between each circle in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/width
	Width() float32

	// The x and y position to use as the center of the circular screen pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/inputImage
	SetInputImage(value ICIImage)

	// The sharpness of the circles.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/sharpness
	SetSharpness(value float32)

	// The distance between each circle in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/width
	SetWidth(value float32)
}

// CICircularScreenObject wraps an existing Objective-C object that conforms to the CICircularScreen protocol.
type CICircularScreenObject struct {
	objectivec.Object
}

func (o CICircularScreenObject) BaseObject() objectivec.Object {
	return o.Object
}

// CICircularScreenObjectFromID constructs a [CICircularScreenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CICircularScreenObjectFromID(id objc.ID) CICircularScreenObject {
	return CICircularScreenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the circular screen pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/center
func (o CICircularScreenObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/inputImage
func (o CICircularScreenObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The sharpness of the circles.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/sharpness
func (o CICircularScreenObject) Sharpness() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sharpness"))
	return rv
}

// The distance between each circle in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/width
func (o CICircularScreenObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CICircularScreenObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The x and y position to use as the center of the circular screen pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/center
func (o CICircularScreenObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/inputImage
func (o CICircularScreenObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The sharpness of the circles.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/sharpness
func (o CICircularScreenObject) SetSharpness(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSharpness:"), value)
}

// The distance between each circle in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CICircularScreen/width
func (o CICircularScreenObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}
