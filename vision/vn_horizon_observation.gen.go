// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
)

// The class instance for the [VNHorizonObservation] class.
var (
	_VNHorizonObservationClass     VNHorizonObservationClass
	_VNHorizonObservationClassOnce sync.Once
)

func getVNHorizonObservationClass() VNHorizonObservationClass {
	_VNHorizonObservationClassOnce.Do(func() {
		_VNHorizonObservationClass = VNHorizonObservationClass{class: objc.GetClass("VNHorizonObservation")}
	})
	return _VNHorizonObservationClass
}

// GetVNHorizonObservationClass returns the class object for VNHorizonObservation.
func GetVNHorizonObservationClass() VNHorizonObservationClass {
	return getVNHorizonObservationClass()
}

type VNHorizonObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNHorizonObservationClass) Alloc() VNHorizonObservation {
	rv := objc.Send[VNHorizonObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The horizon angle information that an image-analysis request detects.
//
// # Overview
// 
// Instances of this class result from invoking a [VNDetectHorizonRequest],
// and report the [VNHorizonObservation.Angle] and [VNHorizonObservation.Transform] of the horizon in an image.
//
// # Evaluating the Horizon
//
//   - [VNHorizonObservation.Angle]: The angle of the observed horizon.
//   - [VNHorizonObservation.Transform]: The transform to apply to the detected horizon.
//   - [VNHorizonObservation.TransformForImageWidthHeight]: Creates an affine transform for the specified image width and height.
//
// See: https://developer.apple.com/documentation/Vision/VNHorizonObservation
type VNHorizonObservation struct {
	VNObservation
}

// VNHorizonObservationFromID constructs a [VNHorizonObservation] from an objc.ID.
//
// The horizon angle information that an image-analysis request detects.
func VNHorizonObservationFromID(id objc.ID) VNHorizonObservation {
	return VNHorizonObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNHorizonObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNHorizonObservation] class.
//
// # Evaluating the Horizon
//
//   - [IVNHorizonObservation.Angle]: The angle of the observed horizon.
//   - [IVNHorizonObservation.Transform]: The transform to apply to the detected horizon.
//   - [IVNHorizonObservation.TransformForImageWidthHeight]: Creates an affine transform for the specified image width and height.
//
// See: https://developer.apple.com/documentation/Vision/VNHorizonObservation
type IVNHorizonObservation interface {
	IVNObservation

	// Topic: Evaluating the Horizon

	// The angle of the observed horizon.
	Angle() float64
	// The transform to apply to the detected horizon.
	Transform() corefoundation.CGAffineTransform
	// Creates an affine transform for the specified image width and height.
	TransformForImageWidthHeight(width uintptr, height uintptr) corefoundation.CGAffineTransform
}

// Init initializes the instance.
func (h VNHorizonObservation) Init() VNHorizonObservation {
	rv := objc.Send[VNHorizonObservation](h.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (h VNHorizonObservation) Autorelease() VNHorizonObservation {
	rv := objc.Send[VNHorizonObservation](h.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNHorizonObservation creates a new VNHorizonObservation instance.
func NewVNHorizonObservation() VNHorizonObservation {
	class := getVNHorizonObservationClass()
	rv := objc.Send[VNHorizonObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an affine transform for the specified image width and height.
//
// width: The width of the image.
//
// height: The height of the image.
//
// # Return Value
// 
// An affine transform.
//
// See: https://developer.apple.com/documentation/Vision/VNHorizonObservation/transform(forImageWidth:height:)
func (h VNHorizonObservation) TransformForImageWidthHeight(width uintptr, height uintptr) corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](h.ID, objc.Sel("transformForImageWidth:height:"), width, height)
	return corefoundation.CGAffineTransform(rv)
}

// The angle of the observed horizon.
//
// # Discussion
// 
// Use the angle to orient the image in an upright position and make the
// detected horizon level.
//
// See: https://developer.apple.com/documentation/Vision/VNHorizonObservation/angle
func (h VNHorizonObservation) Angle() float64 {
	rv := objc.Send[float64](h.ID, objc.Sel("angle"))
	return rv
}
// The transform to apply to the detected horizon.
//
// # Discussion
// 
// Apply the transform’s inverse to orient the image in an upright position
// and make the detected horizon level.
//
// See: https://developer.apple.com/documentation/Vision/VNHorizonObservation/transform
func (h VNHorizonObservation) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](h.ID, objc.Sel("transform"))
	return corefoundation.CGAffineTransform(rv)
}

