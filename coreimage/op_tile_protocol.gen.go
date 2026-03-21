// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an optical tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIOpTile
type CIOpTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/inputImage
	InputImage() ICIImage

	// A value that determines the number of tiles in the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/scale
	Scale() float32

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/width
	Width() float32

	// The angle of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/inputImage
	SetInputImage(value ICIImage)

	// A value that determines the number of tiles in the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/scale
	SetScale(value float32)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/width
	SetWidth(value float32)
}

// CIOpTileObject wraps an existing Objective-C object that conforms to the CIOpTile protocol.
type CIOpTileObject struct {
	objectivec.Object
}
func (o CIOpTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIOpTileObjectFromID constructs a [CIOpTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIOpTileObjectFromID(id objc.ID) CIOpTileObject {
	return CIOpTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/angle
func (o CIOpTileObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/center
func (o CIOpTileObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/inputImage
func (o CIOpTileObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A value that determines the number of tiles in the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/scale
func (o CIOpTileObject) Scale() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("scale"))
	return rv
	}
// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIOpTile/width
func (o CIOpTileObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIOpTileObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIOpTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIOpTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIOpTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIOpTileObject) SetScale(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setScale:"), value)
}

func (o CIOpTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

