// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNVideoProcessorFrameRateCadence] class.
var (
	_VNVideoProcessorFrameRateCadenceClass     VNVideoProcessorFrameRateCadenceClass
	_VNVideoProcessorFrameRateCadenceClassOnce sync.Once
)

func getVNVideoProcessorFrameRateCadenceClass() VNVideoProcessorFrameRateCadenceClass {
	_VNVideoProcessorFrameRateCadenceClassOnce.Do(func() {
		_VNVideoProcessorFrameRateCadenceClass = VNVideoProcessorFrameRateCadenceClass{class: objc.GetClass("VNVideoProcessorFrameRateCadence")}
	})
	return _VNVideoProcessorFrameRateCadenceClass
}

// GetVNVideoProcessorFrameRateCadenceClass returns the class object for VNVideoProcessorFrameRateCadence.
func GetVNVideoProcessorFrameRateCadenceClass() VNVideoProcessorFrameRateCadenceClass {
	return getVNVideoProcessorFrameRateCadenceClass()
}

type VNVideoProcessorFrameRateCadenceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNVideoProcessorFrameRateCadenceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNVideoProcessorFrameRateCadenceClass) Alloc() VNVideoProcessorFrameRateCadence {
	rv := objc.Send[VNVideoProcessorFrameRateCadence](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines a frame-based cadence for processing a video stream.
//
// # Creating a Cadence
//
//   - [VNVideoProcessorFrameRateCadence.InitWithFrameRate]: Creates a new frame-based cadence with a frame rate.
//
// # Inspecting the Frame Rate
//
//   - [VNVideoProcessorFrameRateCadence.FrameRate]: The frame rate at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence
type VNVideoProcessorFrameRateCadence struct {
	VNVideoProcessorCadence
}

// VNVideoProcessorFrameRateCadenceFromID constructs a [VNVideoProcessorFrameRateCadence] from an objc.ID.
//
// An object that defines a frame-based cadence for processing a video stream.
func VNVideoProcessorFrameRateCadenceFromID(id objc.ID) VNVideoProcessorFrameRateCadence {
	return VNVideoProcessorFrameRateCadence{VNVideoProcessorCadence: VNVideoProcessorCadenceFromID(id)}
}
// NOTE: VNVideoProcessorFrameRateCadence adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNVideoProcessorFrameRateCadence] class.
//
// # Creating a Cadence
//
//   - [IVNVideoProcessorFrameRateCadence.InitWithFrameRate]: Creates a new frame-based cadence with a frame rate.
//
// # Inspecting the Frame Rate
//
//   - [IVNVideoProcessorFrameRateCadence.FrameRate]: The frame rate at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence
type IVNVideoProcessorFrameRateCadence interface {
	IVNVideoProcessorCadence

	// Topic: Creating a Cadence

	// Creates a new frame-based cadence with a frame rate.
	InitWithFrameRate(frameRate int) VNVideoProcessorFrameRateCadence

	// Topic: Inspecting the Frame Rate

	// The frame rate at which to process video.
	FrameRate() int
}

// Init initializes the instance.
func (v VNVideoProcessorFrameRateCadence) Init() VNVideoProcessorFrameRateCadence {
	rv := objc.Send[VNVideoProcessorFrameRateCadence](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VNVideoProcessorFrameRateCadence) Autorelease() VNVideoProcessorFrameRateCadence {
	rv := objc.Send[VNVideoProcessorFrameRateCadence](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNVideoProcessorFrameRateCadence creates a new VNVideoProcessorFrameRateCadence instance.
func NewVNVideoProcessorFrameRateCadence() VNVideoProcessorFrameRateCadence {
	class := getVNVideoProcessorFrameRateCadenceClass()
	rv := objc.Send[VNVideoProcessorFrameRateCadence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new frame-based cadence with a frame rate.
//
// frameRate: The frame rate at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence/init(_:)
func NewVideoProcessorFrameRateCadenceWithFrameRate(frameRate int) VNVideoProcessorFrameRateCadence {
	instance := getVNVideoProcessorFrameRateCadenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrameRate:"), frameRate)
	return VNVideoProcessorFrameRateCadenceFromID(rv)
}

// Creates a new frame-based cadence with a frame rate.
//
// frameRate: The frame rate at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence/init(_:)
func (v VNVideoProcessorFrameRateCadence) InitWithFrameRate(frameRate int) VNVideoProcessorFrameRateCadence {
	rv := objc.Send[VNVideoProcessorFrameRateCadence](v.ID, objc.Sel("initWithFrameRate:"), frameRate)
	return rv
}

// The frame rate at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/FrameRateCadence/frameRate
func (v VNVideoProcessorFrameRateCadence) FrameRate() int {
	rv := objc.Send[int](v.ID, objc.Sel("frameRate"))
	return rv
}

