// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a fourfold translated tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile
type CIFourfoldTranslatedTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The primary angle for the repeating translated tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/acuteAngle
	AcuteAngle() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/width
	Width() float32

	// The primary angle for the repeating translated tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/acuteAngle
	SetAcuteAngle(value float32)

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/width
	SetWidth(value float32)
}

// CIFourfoldTranslatedTileObject wraps an existing Objective-C object that conforms to the CIFourfoldTranslatedTile protocol.
type CIFourfoldTranslatedTileObject struct {
	objectivec.Object
}
func (o CIFourfoldTranslatedTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFourfoldTranslatedTileObjectFromID constructs a [CIFourfoldTranslatedTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFourfoldTranslatedTileObjectFromID(id objc.ID) CIFourfoldTranslatedTileObject {
	return CIFourfoldTranslatedTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The primary angle for the repeating translated tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/acuteAngle
func (o CIFourfoldTranslatedTileObject) AcuteAngle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("acuteAngle"))
	return rv
	}
// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/angle
func (o CIFourfoldTranslatedTileObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/center
func (o CIFourfoldTranslatedTileObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/inputImage
func (o CIFourfoldTranslatedTileObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourfoldTranslatedTile/width
func (o CIFourfoldTranslatedTileObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFourfoldTranslatedTileObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIFourfoldTranslatedTileObject) SetAcuteAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAcuteAngle:"), value)
}

func (o CIFourfoldTranslatedTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIFourfoldTranslatedTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIFourfoldTranslatedTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIFourfoldTranslatedTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

