// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure an eightfold reflected tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile
type CIEightfoldReflectedTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/width
	Width() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/width
	SetWidth(value float32)
}

// CIEightfoldReflectedTileObject wraps an existing Objective-C object that conforms to the CIEightfoldReflectedTile protocol.
type CIEightfoldReflectedTileObject struct {
	objectivec.Object
}
func (o CIEightfoldReflectedTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIEightfoldReflectedTileObjectFromID constructs a [CIEightfoldReflectedTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIEightfoldReflectedTileObjectFromID(id objc.ID) CIEightfoldReflectedTileObject {
	return CIEightfoldReflectedTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/angle
func (o CIEightfoldReflectedTileObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/center
func (o CIEightfoldReflectedTileObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/inputImage
func (o CIEightfoldReflectedTileObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIEightfoldReflectedTile/width
func (o CIEightfoldReflectedTileObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIEightfoldReflectedTileObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIEightfoldReflectedTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CIEightfoldReflectedTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CIEightfoldReflectedTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIEightfoldReflectedTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

