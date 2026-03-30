// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a perspective tile filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile
type CIPerspectiveTile interface {
	objectivec.IObject
	CIFilterProtocol

	// The bottom-left coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomLeft
	BottomLeft() corefoundation.CGPoint

	// The bottom-right coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomRight
	BottomRight() corefoundation.CGPoint

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/inputImage
	InputImage() ICIImage

	// The top-left coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topLeft
	TopLeft() corefoundation.CGPoint

	// The top-right coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topRight
	TopRight() corefoundation.CGPoint

	// The bottom-left coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomLeft
	SetBottomLeft(value corefoundation.CGPoint)

	// The bottom-right coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomRight
	SetBottomRight(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/inputImage
	SetInputImage(value ICIImage)

	// The top-left coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topLeft
	SetTopLeft(value corefoundation.CGPoint)

	// The top-right coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topRight
	SetTopRight(value corefoundation.CGPoint)
}

// CIPerspectiveTileObject wraps an existing Objective-C object that conforms to the CIPerspectiveTile protocol.
type CIPerspectiveTileObject struct {
	objectivec.Object
}

func (o CIPerspectiveTileObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPerspectiveTileObjectFromID constructs a [CIPerspectiveTileObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPerspectiveTileObjectFromID(id objc.ID) CIPerspectiveTileObject {
	return CIPerspectiveTileObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The bottom-left coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomLeft
func (o CIPerspectiveTileObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return rv
}

// The bottom-right coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomRight
func (o CIPerspectiveTileObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/inputImage
func (o CIPerspectiveTileObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The top-left coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topLeft
func (o CIPerspectiveTileObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return rv
}

// The top-right coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topRight
func (o CIPerspectiveTileObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPerspectiveTileObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The bottom-left coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomLeft
func (o CIPerspectiveTileObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// The bottom-right coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomRight
func (o CIPerspectiveTileObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/inputImage
func (o CIPerspectiveTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The top-left coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topLeft
func (o CIPerspectiveTileObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// The top-right coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topRight
func (o CIPerspectiveTileObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
