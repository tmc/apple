// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNSaliencyImageObservation] class.
var (
	_VNSaliencyImageObservationClass     VNSaliencyImageObservationClass
	_VNSaliencyImageObservationClassOnce sync.Once
)

func getVNSaliencyImageObservationClass() VNSaliencyImageObservationClass {
	_VNSaliencyImageObservationClassOnce.Do(func() {
		_VNSaliencyImageObservationClass = VNSaliencyImageObservationClass{class: objc.GetClass("VNSaliencyImageObservation")}
	})
	return _VNSaliencyImageObservationClass
}

// GetVNSaliencyImageObservationClass returns the class object for VNSaliencyImageObservation.
func GetVNSaliencyImageObservationClass() VNSaliencyImageObservationClass {
	return getVNSaliencyImageObservationClass()
}

type VNSaliencyImageObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNSaliencyImageObservationClass) Alloc() VNSaliencyImageObservation {
	rv := objc.Send[VNSaliencyImageObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An observation that contains a grayscale heat map of important areas across
// an image.
//
// # Overview
// 
// The heat map is a [CVPixelBuffer] in a one-component floating-point pixel
// format. Its dimensions are 64 x 64 when fetched in real time, or 68 x 68
// when requested in its deferred form.
//
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/CVPixelBuffer
//
// # Locating Salient Regions
//
//   - [VNSaliencyImageObservation.SalientObjects]: A collection of objects describing the distinct areas of the saliency heat map.
//
// See: https://developer.apple.com/documentation/Vision/VNSaliencyImageObservation
type VNSaliencyImageObservation struct {
	VNPixelBufferObservation
}

// VNSaliencyImageObservationFromID constructs a [VNSaliencyImageObservation] from an objc.ID.
//
// An observation that contains a grayscale heat map of important areas across
// an image.
func VNSaliencyImageObservationFromID(id objc.ID) VNSaliencyImageObservation {
	return VNSaliencyImageObservation{VNPixelBufferObservation: VNPixelBufferObservationFromID(id)}
}
// NOTE: VNSaliencyImageObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNSaliencyImageObservation] class.
//
// # Locating Salient Regions
//
//   - [IVNSaliencyImageObservation.SalientObjects]: A collection of objects describing the distinct areas of the saliency heat map.
//
// See: https://developer.apple.com/documentation/Vision/VNSaliencyImageObservation
type IVNSaliencyImageObservation interface {
	IVNPixelBufferObservation

	// Topic: Locating Salient Regions

	// A collection of objects describing the distinct areas of the saliency heat map.
	SalientObjects() []VNRectangleObservation
}

// Init initializes the instance.
func (s VNSaliencyImageObservation) Init() VNSaliencyImageObservation {
	rv := objc.Send[VNSaliencyImageObservation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VNSaliencyImageObservation) Autorelease() VNSaliencyImageObservation {
	rv := objc.Send[VNSaliencyImageObservation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNSaliencyImageObservation creates a new VNSaliencyImageObservation instance.
func NewVNSaliencyImageObservation() VNSaliencyImageObservation {
	class := getVNSaliencyImageObservationClass()
	rv := objc.Send[VNSaliencyImageObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A collection of objects describing the distinct areas of the saliency heat
// map.
//
// # Discussion
// 
// The objects in this array don’t follow any specific ordering. It’s up
// to your app to iterate across the observations and apply desired ordering.
// 
// Requesting this array lazily computes the bounds of salient objects within
// the image.
//
// See: https://developer.apple.com/documentation/Vision/VNSaliencyImageObservation/salientObjects
func (s VNSaliencyImageObservation) SalientObjects() []VNRectangleObservation {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("salientObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) VNRectangleObservation {
		return VNRectangleObservationFromID(id)
	})
}

