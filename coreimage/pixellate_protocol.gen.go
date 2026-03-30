// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a pixellate filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate
type CIPixellate interface {
	objectivec.IObject
	CIFilterProtocol

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/inputImage
	InputImage() ICIImage

	// A value that determines the size of the squares.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/scale
	Scale() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/inputImage
	SetInputImage(value ICIImage)

	// A value that determines the size of the squares.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/scale
	SetScale(value float32)
}

// CIPixellateObject wraps an existing Objective-C object that conforms to the CIPixellate protocol.
type CIPixellateObject struct {
	objectivec.Object
}

func (o CIPixellateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPixellateObjectFromID constructs a [CIPixellateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPixellateObjectFromID(id objc.ID) CIPixellateObject {
	return CIPixellateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/center
func (o CIPixellateObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/inputImage
func (o CIPixellateObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A value that determines the size of the squares.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/scale
func (o CIPixellateObject) Scale() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPixellateObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/center
func (o CIPixellateObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/inputImage
func (o CIPixellateObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// A value that determines the size of the squares.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPixellate/scale
func (o CIPixellateObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}
