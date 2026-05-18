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
	SetBottomLeft(value corefoundation.CGPoint)

	// The bottom-right coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomRight
	BottomRight() corefoundation.CGPoint
	SetBottomRight(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// The top-left coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topLeft
	TopLeft() corefoundation.CGPoint
	SetTopLeft(value corefoundation.CGPoint)

	// The top-right coordinate of a tile.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topRight
	TopRight() corefoundation.CGPoint
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
func (o CIPerspectiveTileObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return corefoundation.CGPoint(rv)
}

func (o CIPerspectiveTileObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// The bottom-right coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/bottomRight
func (o CIPerspectiveTileObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return corefoundation.CGPoint(rv)
}

func (o CIPerspectiveTileObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/inputImage
func (o CIPerspectiveTileObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIPerspectiveTileObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The top-left coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topLeft
func (o CIPerspectiveTileObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return corefoundation.CGPoint(rv)
}

func (o CIPerspectiveTileObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// The top-right coordinate of a tile.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTile/topRight
func (o CIPerspectiveTileObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return corefoundation.CGPoint(rv)
}

func (o CIPerspectiveTileObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
