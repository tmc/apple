// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a sixfold reflected tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile
type CISixfoldReflectedTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/center
	Center() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/inputImage
	InputImage() ICIImage

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/width
	Width() float32

	// The angle, in radians, of the tiled pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/center
	SetCenter(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/inputImage
	SetInputImage(value ICIImage)

	// The width of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/width
	SetWidth(value float32)
}

// CISixfoldReflectedTileObject wraps an existing Objective-C object that conforms to the CISixfoldReflectedTile protocol.
type CISixfoldReflectedTileObject struct {
	objectivec.Object
}
func (o CISixfoldReflectedTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CISixfoldReflectedTileObjectFromID constructs a [CISixfoldReflectedTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CISixfoldReflectedTileObjectFromID(id objc.ID) CISixfoldReflectedTileObject {
	return CISixfoldReflectedTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle, in radians, of the tiled pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/angle
func (o CISixfoldReflectedTileObject) Angle() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
	}
// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/center
func (o CISixfoldReflectedTileObject) Center() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/inputImage
func (o CISixfoldReflectedTileObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The width of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CISixfoldReflectedTile/width
func (o CISixfoldReflectedTileObject) Width() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("width"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CISixfoldReflectedTileObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CISixfoldReflectedTileObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

func (o CISixfoldReflectedTileObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

func (o CISixfoldReflectedTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CISixfoldReflectedTileObject) SetWidth(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setWidth:"), value)
}

