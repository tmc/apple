// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a geometry adjustment filters that requires four coordinates.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter
type CIFourCoordinateGeometryFilter interface {
	objectivec.IObject
	CIFilterProtocol

	// bottomLeft protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
	BottomLeft() corefoundation.CGPoint
	SetBottomLeft(value corefoundation.CGPoint)

	// bottomRight protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
	BottomRight() corefoundation.CGPoint
	SetBottomRight(value corefoundation.CGPoint)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
	InputImage() ICIImage
	SetInputImage(value ICIImage)

	// topLeft protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
	TopLeft() corefoundation.CGPoint
	SetTopLeft(value corefoundation.CGPoint)

	// topRight protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
	TopRight() corefoundation.CGPoint
	SetTopRight(value corefoundation.CGPoint)
}

// CIFourCoordinateGeometryFilterObject wraps an existing Objective-C object that conforms to the CIFourCoordinateGeometryFilter protocol.
type CIFourCoordinateGeometryFilterObject struct {
	objectivec.Object
}

func (o CIFourCoordinateGeometryFilterObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIFourCoordinateGeometryFilterObjectFromID constructs a [CIFourCoordinateGeometryFilterObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIFourCoordinateGeometryFilterObjectFromID(id objc.ID) CIFourCoordinateGeometryFilterObject {
	return CIFourCoordinateGeometryFilterObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIFourCoordinateGeometryFilterObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIFourCoordinateGeometryFilterObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return corefoundation.CGPoint(rv)
}

func (o CIFourCoordinateGeometryFilterObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIFourCoordinateGeometryFilterObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return corefoundation.CGPoint(rv)
}

func (o CIFourCoordinateGeometryFilterObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIFourCoordinateGeometryFilterObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

func (o CIFourCoordinateGeometryFilterObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIFourCoordinateGeometryFilterObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return corefoundation.CGPoint(rv)
}

func (o CIFourCoordinateGeometryFilterObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIFourCoordinateGeometryFilterObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return corefoundation.CGPoint(rv)
}

func (o CIFourCoordinateGeometryFilterObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
