// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a straighten filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStraighten
type CIStraighten interface {
	objectivec.IObject
	CIFilterProtocol

	// The rotation angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/angle
	Angle() float32

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/inputImage
	InputImage() ICIImage

	// The rotation angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/angle
	SetAngle(value float32)

	// The image to use as an input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/inputImage
	SetInputImage(value ICIImage)
}

// CIStraightenObject wraps an existing Objective-C object that conforms to the CIStraighten protocol.
type CIStraightenObject struct {
	objectivec.Object
}

func (o CIStraightenObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIStraightenObjectFromID constructs a [CIStraightenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIStraightenObjectFromID(id objc.ID) CIStraightenObject {
	return CIStraightenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The rotation angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/angle
func (o CIStraightenObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/inputImage
func (o CIStraightenObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIStraightenObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The rotation angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/angle
func (o CIStraightenObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIStraighten/inputImage
func (o CIStraightenObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}
