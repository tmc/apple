// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIColorAbsoluteDifference protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference
type CIColorAbsoluteDifference interface {
	objectivec.IObject
	CIFilterProtocol

	// The first image to use for differencing.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference/inputImage
	InputImage() ICIImage

	// The second image to use for differencing.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference/inputImage2
	InputImage2() ICIImage

	// The first image to use for differencing.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference/inputImage
	SetInputImage(value ICIImage)

	// The second image to use for differencing.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference/inputImage2
	SetInputImage2(value ICIImage)
}

// CIColorAbsoluteDifferenceObject wraps an existing Objective-C object that conforms to the CIColorAbsoluteDifference protocol.
type CIColorAbsoluteDifferenceObject struct {
	objectivec.Object
}
func (o CIColorAbsoluteDifferenceObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIColorAbsoluteDifferenceObjectFromID constructs a [CIColorAbsoluteDifferenceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIColorAbsoluteDifferenceObjectFromID(id objc.ID) CIColorAbsoluteDifferenceObject {
	return CIColorAbsoluteDifferenceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The first image to use for differencing.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference/inputImage
func (o CIColorAbsoluteDifferenceObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
	}
// The second image to use for differencing.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColorAbsoluteDifference/inputImage2
func (o CIColorAbsoluteDifferenceObject) InputImage2() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage2"))
	return CIImageFromID(rv)
	}
// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIColorAbsoluteDifferenceObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
	}

func (o CIColorAbsoluteDifferenceObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

func (o CIColorAbsoluteDifferenceObject) SetInputImage2(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage2:"), value)
}

