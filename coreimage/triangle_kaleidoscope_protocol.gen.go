// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a triangle kaleidoscope filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope
type CITriangleKaleidoscope interface {
	objectivec.IObject
	CIFilterProtocol

	// A value that determines how fast the color fades from the center triangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/decay
	Decay() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/inputImage
	InputImage() ICIImage

	// The x and y position to use as the center of the triangular area in the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/point
	Point() corefoundation.CGPoint

	// The rotation angle of the triangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/rotation
	Rotation() float32

	// The size, in pixels, of the triangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/size
	Size() float32

	// A value that determines how fast the color fades from the center triangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/decay
	SetDecay(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/inputImage
	SetInputImage(value ICIImage)

	// The x and y position to use as the center of the triangular area in the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/point
	SetPoint(value corefoundation.CGPoint)

	// The rotation angle of the triangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/rotation
	SetRotation(value float32)

	// The size, in pixels, of the triangle.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/size
	SetSize(value float32)
}

// CITriangleKaleidoscopeObject wraps an existing Objective-C object that conforms to the CITriangleKaleidoscope protocol.
type CITriangleKaleidoscopeObject struct {
	objectivec.Object
}
func (o CITriangleKaleidoscopeObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITriangleKaleidoscopeObjectFromID constructs a [CITriangleKaleidoscopeObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITriangleKaleidoscopeObjectFromID(id objc.ID) CITriangleKaleidoscopeObject {
	return CITriangleKaleidoscopeObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A value that determines how fast the color fades from the center triangle.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/decay
func (o CITriangleKaleidoscopeObject) Decay() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("decay"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/inputImage
func (o CITriangleKaleidoscopeObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The x and y position to use as the center of the triangular area in the
// input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/point
func (o CITriangleKaleidoscopeObject) Point() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("point"))
	return rv
	}
// The rotation angle of the triangle.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/rotation
func (o CITriangleKaleidoscopeObject) Rotation() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("rotation"))
	return rv
	}
// The size, in pixels, of the triangle.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleKaleidoscope/size
func (o CITriangleKaleidoscopeObject) Size() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("size"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITriangleKaleidoscopeObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITriangleKaleidoscopeObject) SetDecay(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setDecay:"), value)
}

func (o CITriangleKaleidoscopeObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CITriangleKaleidoscopeObject) SetPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setPoint:"), value)
}

func (o CITriangleKaleidoscopeObject) SetRotation(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRotation:"), value)
}

func (o CITriangleKaleidoscopeObject) SetSize(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSize:"), value)
}

