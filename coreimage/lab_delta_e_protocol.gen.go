// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a Lab Delta E filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE
type CILabDeltaE interface {
	objectivec.IObject
	CIFilterProtocol

	// The second input image for comparison.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE/image2
	Image2() ICIImage

	// The first input image for comparison.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE/inputImage
	InputImage() ICIImage

	// The second input image for comparison.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE/image2
	SetImage2(value ICIImage)

	// The first input image for comparison.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE/inputImage
	SetInputImage(value ICIImage)
}

// CILabDeltaEObject wraps an existing Objective-C object that conforms to the CILabDeltaE protocol.
type CILabDeltaEObject struct {
	objectivec.Object
}
func (o CILabDeltaEObject) BaseObject() objectivec.Object {
	return o.Object
}

// CILabDeltaEObjectFromID constructs a [CILabDeltaEObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CILabDeltaEObjectFromID(id objc.ID) CILabDeltaEObject {
	return CILabDeltaEObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The second input image for comparison.
//
// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE/image2
func (o CILabDeltaEObject) Image2() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("image2"))
	return CIImageFromID(rv)
	}
// The first input image for comparison.
//
// See: https://developer.apple.com/documentation/CoreImage/CILabDeltaE/inputImage
func (o CILabDeltaEObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CILabDeltaEObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CILabDeltaEObject) SetImage2(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setImage2:"), value)
}

func (o CILabDeltaEObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

