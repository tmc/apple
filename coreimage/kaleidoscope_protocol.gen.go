// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a kaleidoscope filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope
type CIKaleidoscope interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of the reflection.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/center
	Center() corefoundation.CGPoint

	// The number of reflections in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/count
	Count() int

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/inputImage
	InputImage() ICIImage

	// The angle of the reflection.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/center
	SetCenter(value corefoundation.CGPoint)

	// The number of reflections in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/count
	SetCount(value int)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/inputImage
	SetInputImage(value ICIImage)
}

// CIKaleidoscopeObject wraps an existing Objective-C object that conforms to the CIKaleidoscope protocol.
type CIKaleidoscopeObject struct {
	objectivec.Object
}
func (o CIKaleidoscopeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIKaleidoscopeObjectFromID constructs a [CIKaleidoscopeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIKaleidoscopeObjectFromID(id objc.ID) CIKaleidoscopeObject {
	return CIKaleidoscopeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the reflection.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/angle
func (o CIKaleidoscopeObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/center
func (o CIKaleidoscopeObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The number of reflections in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/count
func (o CIKaleidoscopeObject) Count() int {
	rv := objc.Send[int](o.ID, objc.Sel("count"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKaleidoscope/inputImage
func (o CIKaleidoscopeObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIKaleidoscopeObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIKaleidoscopeObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIKaleidoscopeObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIKaleidoscopeObject) SetCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setCount:"), value)
}

func (o CIKaleidoscopeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

