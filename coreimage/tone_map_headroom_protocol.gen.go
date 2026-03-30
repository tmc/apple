// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CIToneMapHeadroom protocol.
//
// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom
type CIToneMapHeadroom interface {
	objectivec.IObject
	CIFilterProtocol

	// InputImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/inputImage
	InputImage() ICIImage

	// SourceHeadroom protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/sourceHeadroom
	SourceHeadroom() float32

	// TargetHeadroom protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/targetHeadroom
	TargetHeadroom() float32

	// inputImage protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/inputImage
	SetInputImage(value ICIImage)

	// sourceHeadroom protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/sourceHeadroom
	SetSourceHeadroom(value float32)

	// targetHeadroom protocol.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/targetHeadroom
	SetTargetHeadroom(value float32)
}

// CIToneMapHeadroomObject wraps an existing Objective-C object that conforms to the CIToneMapHeadroom protocol.
type CIToneMapHeadroomObject struct {
	objectivec.Object
}

func (o CIToneMapHeadroomObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIToneMapHeadroomObjectFromID constructs a [CIToneMapHeadroomObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIToneMapHeadroomObjectFromID(id objc.ID) CIToneMapHeadroomObject {
	return CIToneMapHeadroomObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/inputImage
func (o CIToneMapHeadroomObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/sourceHeadroom
func (o CIToneMapHeadroomObject) SourceHeadroom() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("sourceHeadroom"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/targetHeadroom
func (o CIToneMapHeadroomObject) TargetHeadroom() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("targetHeadroom"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIToneMapHeadroomObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/inputImage
func (o CIToneMapHeadroomObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/sourceHeadroom
func (o CIToneMapHeadroomObject) SetSourceHeadroom(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setSourceHeadroom:"), value)
}

// See: https://developer.apple.com/documentation/CoreImage/CIToneMapHeadroom/targetHeadroom
func (o CIToneMapHeadroomObject) SetTargetHeadroom(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetHeadroom:"), value)
}
