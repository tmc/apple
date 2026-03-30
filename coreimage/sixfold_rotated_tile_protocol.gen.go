// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a sixfold rotated tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile
type CISixfoldRotatedTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/width
	Width() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/width
	SetWidth(value float32)
}

// CISixfoldRotatedTileObject wraps an existing Objective-C object that conforms to the CISixfoldRotatedTile protocol.
type CISixfoldRotatedTileObject struct {
	objectivec.Object
}

func (o CISixfoldRotatedTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISixfoldRotatedTileObjectFromID constructs a [CISixfoldRotatedTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISixfoldRotatedTileObjectFromID(id objc.ID) CISixfoldRotatedTileObject {
	return CISixfoldRotatedTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/angle
func (o CISixfoldRotatedTileObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/center
func (o CISixfoldRotatedTileObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/inputImage
func (o CISixfoldRotatedTileObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/width
func (o CISixfoldRotatedTileObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISixfoldRotatedTileObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/angle
func (o CISixfoldRotatedTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/center
func (o CISixfoldRotatedTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/inputImage
func (o CISixfoldRotatedTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldRotatedTile/width
func (o CISixfoldRotatedTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}
