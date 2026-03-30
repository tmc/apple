// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a fourfold rotated tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile
type CIFourfoldRotatedTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/width
	Width() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/width
	SetWidth(value float32)
}

// CIFourfoldRotatedTileObject wraps an existing Objective-C object that conforms to the CIFourfoldRotatedTile protocol.
type CIFourfoldRotatedTileObject struct {
	objectivec.Object
}

func (o CIFourfoldRotatedTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFourfoldRotatedTileObjectFromID constructs a [CIFourfoldRotatedTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFourfoldRotatedTileObjectFromID(id objc.ID) CIFourfoldRotatedTileObject {
	return CIFourfoldRotatedTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/angle
func (o CIFourfoldRotatedTileObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/center
func (o CIFourfoldRotatedTileObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/inputImage
func (o CIFourfoldRotatedTileObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/width
func (o CIFourfoldRotatedTileObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFourfoldRotatedTileObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/angle
func (o CIFourfoldRotatedTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/center
func (o CIFourfoldRotatedTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/inputImage
func (o CIFourfoldRotatedTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldRotatedTile/width
func (o CIFourfoldRotatedTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}
