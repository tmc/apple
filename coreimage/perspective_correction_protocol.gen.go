// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a perspective correction filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveCorrection
type CIPerspectiveCorrection interface {
	objectivec.IObject
	CIFilterProtocol
	CIFourCoordinateGeometryFilter

	// A rectangle that specifies the extent of the corrected image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveCorrection/crop
	Crop() bool

	// A rectangle that specifies the extent of the corrected image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveCorrection/crop
	SetCrop(value bool)
}

// CIPerspectiveCorrectionObject wraps an existing Objective-C object that conforms to the CIPerspectiveCorrection protocol.
type CIPerspectiveCorrectionObject struct {
	objectivec.Object
}

func (o CIPerspectiveCorrectionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPerspectiveCorrectionObjectFromID constructs a [CIPerspectiveCorrectionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPerspectiveCorrectionObjectFromID(id objc.ID) CIPerspectiveCorrectionObject {
	return CIPerspectiveCorrectionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A rectangle that specifies the extent of the corrected image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveCorrection/crop
func (o CIPerspectiveCorrectionObject) Crop() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("crop"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPerspectiveCorrectionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIPerspectiveCorrectionObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIPerspectiveCorrectionObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIPerspectiveCorrectionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIPerspectiveCorrectionObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIPerspectiveCorrectionObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return rv
}

// A rectangle that specifies the extent of the corrected image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveCorrection/crop
func (o CIPerspectiveCorrectionObject) SetCrop(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setCrop:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIPerspectiveCorrectionObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIPerspectiveCorrectionObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIPerspectiveCorrectionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIPerspectiveCorrectionObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIPerspectiveCorrectionObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
