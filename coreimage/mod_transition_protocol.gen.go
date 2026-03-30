// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The properties you use to configure a mod transition filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition
type CIModTransition interface {
	objectivec.IObject
	CIFilterProtocol
	CITransitionFilter

	// The angle of the mod hole pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/angle
	Angle() float32

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/center
	Center() corefoundation.CGPoint

	// The amount of stretching applied to the mod hole pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/compression
	Compression() float32

	// The radius of the undistorted mod holes in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/radius
	Radius() float32

	// The angle of the mod hole pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/angle
	SetAngle(value float32)

	// The x and y position to use as the center of the effect.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/center
	SetCenter(value corefoundation.CGPoint)

	// The amount of stretching applied to the mod hole pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/compression
	SetCompression(value float32)

	// The radius of the undistorted mod holes in the pattern.
	//
	// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/radius
	SetRadius(value float32)
}

// CIModTransitionObject wraps an existing Objective-C object that conforms to the CIModTransition protocol.
type CIModTransitionObject struct {
	objectivec.Object
}

func (o CIModTransitionObject) BaseObject() objectivec.Object {
	return o.Object
}

// CIModTransitionObjectFromID constructs a [CIModTransitionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CIModTransitionObjectFromID(id objc.ID) CIModTransitionObject {
	return CIModTransitionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The angle of the mod hole pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/angle
func (o CIModTransitionObject) Angle() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("angle"))
	return rv
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/center
func (o CIModTransitionObject) Center() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("center"))
	return rv
}

// The amount of stretching applied to the mod hole pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/compression
func (o CIModTransitionObject) Compression() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("compression"))
	return rv
}

// The radius of the undistorted mod holes in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/radius
func (o CIModTransitionObject) Radius() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("radius"))
	return rv
}

// A [CIImage] object that encapsulates the operations configured in the
// filter.
//
// See: https://developer.apple.com/documentation/CoreImage/CIFilterProtocol/outputImage
func (o CIModTransitionObject) OutputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("outputImage"))
	return CIImageFromID(rv)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIModTransitionObject) InputImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputImage"))
	return CIImageFromID(rv)
}

// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIModTransitionObject) TargetImage() ICIImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("targetImage"))
	return CIImageFromID(rv)
}

// The parametric time of the transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIModTransitionObject) Time() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("time"))
	return rv
}

// The angle of the mod hole pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/angle
func (o CIModTransitionObject) SetAngle(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAngle:"), value)
}

// The x and y position to use as the center of the effect.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/center
func (o CIModTransitionObject) SetCenter(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setCenter:"), value)
}

// The amount of stretching applied to the mod hole pattern.
//
// # Discussion
//
// Holes in the center aren’t distorted as much as those at the edge of the
// image.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/compression
func (o CIModTransitionObject) SetCompression(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setCompression:"), value)
}

// The radius of the undistorted mod holes in the pattern.
//
// See: https://developer.apple.com/documentation/CoreImage/CIModTransition/radius
func (o CIModTransitionObject) SetRadius(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setRadius:"), value)
}

// The image to use as an input image.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/inputImage
func (o CIModTransitionObject) SetInputImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputImage:"), value)
}

// The target image for a transition.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/targetImage
func (o CIModTransitionObject) SetTargetImage(value ICIImage) {
	objc.Send[struct{}](o.ID, objc.Sel("setTargetImage:"), value)
}

// The parametric time of the transition.
//
// # Discussion
//
// This value drives the transition from start, at time 0, to end, at time 1.
//
// See: https://developer.apple.com/documentation/CoreImage/CITransitionFilter/time
func (o CIModTransitionObject) SetTime(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setTime:"), value)
}
