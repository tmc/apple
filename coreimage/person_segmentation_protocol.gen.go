// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIPersonSegmentation protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation
type CIPersonSegmentation interface {
	objectivec.IObject
	CIFilterProtocol

	// InputImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/inputImage
	InputImage() ICIImage

	// QualityLevel protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/qualityLevel
	QualityLevel() uint

	// inputImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/inputImage
	SetInputImage(value ICIImage)

	// qualityLevel protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/qualityLevel
	SetQualityLevel(value uint)
}

// CIPersonSegmentationObject wraps an existing Objective-C object that conforms to the CIPersonSegmentation protocol.
type CIPersonSegmentationObject struct {
	objectivec.Object
}

func (o CIPersonSegmentationObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPersonSegmentationObjectFromID constructs a [CIPersonSegmentationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPersonSegmentationObjectFromID(id objc.ID) CIPersonSegmentationObject {
	return CIPersonSegmentationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/inputImage
func (o CIPersonSegmentationObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/qualityLevel
func (o CIPersonSegmentationObject) QualityLevel() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("qualityLevel"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPersonSegmentationObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/inputImage
func (o CIPersonSegmentationObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIPersonSegmentation/qualityLevel
func (o CIPersonSegmentationObject) SetQualityLevel(value uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setQualityLevel:"), value)
}
