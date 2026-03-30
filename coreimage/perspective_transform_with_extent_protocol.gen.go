// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a perspective transform with extent filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTransformWithExtent
type CIPerspectiveTransformWithExtent interface {
	objectivec.IObject
	CIFilterProtocol
	CIFourCoordinateGeometryFilter

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTransformWithExtent/extent
	Extent() corefoundation.CGRect

	// A rectangle that defines the extent of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTransformWithExtent/extent
	SetExtent(value corefoundation.CGRect)
}

// CIPerspectiveTransformWithExtentObject wraps an existing Objective-C object that conforms to the CIPerspectiveTransformWithExtent protocol.
type CIPerspectiveTransformWithExtentObject struct {
	objectivec.Object
}

func (o CIPerspectiveTransformWithExtentObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPerspectiveTransformWithExtentObjectFromID constructs a [CIPerspectiveTransformWithExtentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPerspectiveTransformWithExtentObjectFromID(id objc.ID) CIPerspectiveTransformWithExtentObject {
	return CIPerspectiveTransformWithExtentObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A rectangle that defines the extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTransformWithExtent/extent
func (o CIPerspectiveTransformWithExtentObject) Extent() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("extent"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPerspectiveTransformWithExtentObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIPerspectiveTransformWithExtentObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIPerspectiveTransformWithExtentObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIPerspectiveTransformWithExtentObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIPerspectiveTransformWithExtentObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIPerspectiveTransformWithExtentObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return rv
}

// A rectangle that defines the extent of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveTransformWithExtent/extent
func (o CIPerspectiveTransformWithExtentObject) SetExtent(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setExtent:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIPerspectiveTransformWithExtentObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIPerspectiveTransformWithExtentObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIPerspectiveTransformWithExtentObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIPerspectiveTransformWithExtentObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIPerspectiveTransformWithExtentObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
