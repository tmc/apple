// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a fourfold reflected tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile
type CIFourfoldReflectedTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The primary angle for the repeating reflected tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/acuteAngle
	AcuteAngle() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/width
	Width() float32

	// The primary angle for the repeating reflected tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/acuteAngle
	SetAcuteAngle(value float32)

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/width
	SetWidth(value float32)
}

// CIFourfoldReflectedTileObject wraps an existing Objective-C object that conforms to the CIFourfoldReflectedTile protocol.
type CIFourfoldReflectedTileObject struct {
	objectivec.Object
}
func (o CIFourfoldReflectedTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFourfoldReflectedTileObjectFromID constructs a [CIFourfoldReflectedTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFourfoldReflectedTileObjectFromID(id objc.ID) CIFourfoldReflectedTileObject {
	return CIFourfoldReflectedTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The primary angle for the repeating reflected tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/acuteAngle
func (o CIFourfoldReflectedTileObject) AcuteAngle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("acuteAngle"))
	return rv
	}
// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/angle
func (o CIFourfoldReflectedTileObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/center
func (o CIFourfoldReflectedTileObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/inputImage
func (o CIFourfoldReflectedTileObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldReflectedTile/width
func (o CIFourfoldReflectedTileObject) Width() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFourfoldReflectedTileObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIFourfoldReflectedTileObject) SetAcuteAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAcuteAngle:"), value)
}

func (o CIFourfoldReflectedTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIFourfoldReflectedTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIFourfoldReflectedTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIFourfoldReflectedTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

