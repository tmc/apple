// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a keystone correction vertical filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionVertical
type CIKeystoneCorrectionVertical interface {
	objectivec.IObject
	CIFilterProtocol
	CIFourCoordinateGeometryFilter

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionVertical/focalLength
	FocalLength() float32

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionVertical/focalLength
	SetFocalLength(value float32)
}

// CIKeystoneCorrectionVerticalObject wraps an existing Objective-C object that conforms to the CIKeystoneCorrectionVertical protocol.
type CIKeystoneCorrectionVerticalObject struct {
	objectivec.Object
}

func (o CIKeystoneCorrectionVerticalObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIKeystoneCorrectionVerticalObjectFromID constructs a [CIKeystoneCorrectionVerticalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIKeystoneCorrectionVerticalObjectFromID(id objc.ID) CIKeystoneCorrectionVerticalObject {
	return CIKeystoneCorrectionVerticalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionVertical/focalLength
func (o CIKeystoneCorrectionVerticalObject) FocalLength() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("focalLength"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIKeystoneCorrectionVerticalObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIKeystoneCorrectionVerticalObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIKeystoneCorrectionVerticalObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIKeystoneCorrectionVerticalObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIKeystoneCorrectionVerticalObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIKeystoneCorrectionVerticalObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return rv
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionVertical/focalLength
func (o CIKeystoneCorrectionVerticalObject) SetFocalLength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFocalLength:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIKeystoneCorrectionVerticalObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIKeystoneCorrectionVerticalObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIKeystoneCorrectionVerticalObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIKeystoneCorrectionVerticalObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIKeystoneCorrectionVerticalObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
