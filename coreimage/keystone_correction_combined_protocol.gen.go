// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a keystone correction combined filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionCombined
type CIKeystoneCorrectionCombined interface {
	objectivec.IObject
	CIFilterProtocol
	CIFourCoordinateGeometryFilter

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionCombined/focalLength
	FocalLength() float32

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionCombined/focalLength
	SetFocalLength(value float32)
}

// CIKeystoneCorrectionCombinedObject wraps an existing Objective-C object that conforms to the CIKeystoneCorrectionCombined protocol.
type CIKeystoneCorrectionCombinedObject struct {
	objectivec.Object
}

func (o CIKeystoneCorrectionCombinedObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIKeystoneCorrectionCombinedObjectFromID constructs a [CIKeystoneCorrectionCombinedObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIKeystoneCorrectionCombinedObjectFromID(id objc.ID) CIKeystoneCorrectionCombinedObject {
	return CIKeystoneCorrectionCombinedObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionCombined/focalLength
func (o CIKeystoneCorrectionCombinedObject) FocalLength() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("focalLength"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIKeystoneCorrectionCombinedObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIKeystoneCorrectionCombinedObject) BottomLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIKeystoneCorrectionCombinedObject) BottomRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIKeystoneCorrectionCombinedObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIKeystoneCorrectionCombinedObject) TopLeft() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIKeystoneCorrectionCombinedObject) TopRight() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return rv
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionCombined/focalLength
func (o CIKeystoneCorrectionCombinedObject) SetFocalLength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFocalLength:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIKeystoneCorrectionCombinedObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIKeystoneCorrectionCombinedObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIKeystoneCorrectionCombinedObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIKeystoneCorrectionCombinedObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIKeystoneCorrectionCombinedObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}
