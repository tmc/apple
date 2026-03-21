// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a keystone correction horizontal filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionHorizontal
type CIKeystoneCorrectionHorizontal interface {
	objectivec.IObject
	CIFilterProtocol
	CIFourCoordinateGeometryFilter

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionHorizontal/focalLength
	FocalLength() float32

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionHorizontal/focalLength
	SetFocalLength(value float32)
}

// CIKeystoneCorrectionHorizontalObject wraps an existing Objective-C object that conforms to the CIKeystoneCorrectionHorizontal protocol.
type CIKeystoneCorrectionHorizontalObject struct {
	objectivec.Object
}
func (o CIKeystoneCorrectionHorizontalObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIKeystoneCorrectionHorizontalObjectFromID constructs a [CIKeystoneCorrectionHorizontalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIKeystoneCorrectionHorizontalObjectFromID(id objc.ID) CIKeystoneCorrectionHorizontalObject {
	return CIKeystoneCorrectionHorizontalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIKeystoneCorrectionHorizontal/focalLength
func (o CIKeystoneCorrectionHorizontalObject) FocalLength() float32 {
	
	rv := objc.Send[float32](o.ID, objc.Sel("focalLength"))
	return rv
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIKeystoneCorrectionHorizontalObject) OutputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomLeft
func (o CIKeystoneCorrectionHorizontalObject) BottomLeft() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomLeft"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/bottomRight
func (o CIKeystoneCorrectionHorizontalObject) BottomRight() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("bottomRight"))
	return rv
	}
// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/inputImage
func (o CIKeystoneCorrectionHorizontalObject) InputImage() ICIImage {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topLeft
func (o CIKeystoneCorrectionHorizontalObject) TopLeft() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topLeft"))
	return rv
	}
// See: https://developer.apple.com/documentation/CoreImage/CIFourCoordinateGeometryFilter/topRight
func (o CIKeystoneCorrectionHorizontalObject) TopRight() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("topRight"))
	return rv
	}

func (o CIKeystoneCorrectionHorizontalObject) SetFocalLength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFocalLength:"), value)
}

func (o CIKeystoneCorrectionHorizontalObject) SetBottomLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomLeft:"), value)
}

func (o CIKeystoneCorrectionHorizontalObject) SetBottomRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBottomRight:"), value)
}

func (o CIKeystoneCorrectionHorizontalObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIKeystoneCorrectionHorizontalObject) SetTopLeft(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopLeft:"), value)
}

func (o CIKeystoneCorrectionHorizontalObject) SetTopRight(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTopRight:"), value)
}

