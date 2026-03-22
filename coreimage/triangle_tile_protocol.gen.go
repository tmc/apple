// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a triangle tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile
type CITriangleTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/width
	Width() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/width
	SetWidth(value float32)
}

// CITriangleTileObject wraps an existing Objective-C object that conforms to the CITriangleTile protocol.
type CITriangleTileObject struct {
	objectivec.Object
}
func (o CITriangleTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CITriangleTileObjectFromID constructs a [CITriangleTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CITriangleTileObjectFromID(id objc.ID) CITriangleTileObject {
	return CITriangleTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/angle
func (o CITriangleTileObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/center
func (o CITriangleTileObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/inputImage
func (o CITriangleTileObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CITriangleTile/width
func (o CITriangleTileObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CITriangleTileObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CITriangleTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CITriangleTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CITriangleTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CITriangleTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

