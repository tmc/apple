// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a perspective rotate filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate
type CIPerspectiveRotate interface {
	objectivec.IObject
	CIFilterProtocol

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/focalLength
	FocalLength() float32

	// The image to process.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/inputImage
	InputImage() ICIImage

	// The pitch angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/pitch
	Pitch() float32

	// The roll angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/roll
	Roll() float32

	// The yaw angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/yaw
	Yaw() float32

	// The 35mm equivalent focal length of the input image.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/focalLength
	SetFocalLength(value float32)

	// The image to process.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/inputImage
	SetInputImage(value ICIImage)

	// The pitch angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/pitch
	SetPitch(value float32)

	// The roll angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/roll
	SetRoll(value float32)

	// The yaw angle, in radians.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/yaw
	SetYaw(value float32)
}

// CIPerspectiveRotateObject wraps an existing Objective-C object that conforms to the CIPerspectiveRotate protocol.
type CIPerspectiveRotateObject struct {
	objectivec.Object
}

func (o CIPerspectiveRotateObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIPerspectiveRotateObjectFromID constructs a [CIPerspectiveRotateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIPerspectiveRotateObjectFromID(id objc.ID) CIPerspectiveRotateObject {
	return CIPerspectiveRotateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/focalLength
func (o CIPerspectiveRotateObject) FocalLength() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("focalLength"))
	return rv
}

// The image to process.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/inputImage
func (o CIPerspectiveRotateObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The pitch angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/pitch
func (o CIPerspectiveRotateObject) Pitch() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("pitch"))
	return rv
}

// The roll angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/roll
func (o CIPerspectiveRotateObject) Roll() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("roll"))
	return rv
}

// The yaw angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/yaw
func (o CIPerspectiveRotateObject) Yaw() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("yaw"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIPerspectiveRotateObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The 35mm equivalent focal length of the input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/focalLength
func (o CIPerspectiveRotateObject) SetFocalLength(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setFocalLength:"), value)
}

// The image to process.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/inputImage
func (o CIPerspectiveRotateObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The pitch angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/pitch
func (o CIPerspectiveRotateObject) SetPitch(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setPitch:"), value)
}

// The roll angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/roll
func (o CIPerspectiveRotateObject) SetRoll(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRoll:"), value)
}

// The yaw angle, in radians.
//
// See: https://developer.apple.com/documentation/CoreImage/CIPerspectiveRotate/yaw
func (o CIPerspectiveRotateObject) SetYaw(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setYaw:"), value)
}
