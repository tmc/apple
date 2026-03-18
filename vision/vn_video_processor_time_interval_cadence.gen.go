// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNVideoProcessorTimeIntervalCadence] class.
var (
	_VNVideoProcessorTimeIntervalCadenceClass     VNVideoProcessorTimeIntervalCadenceClass
	_VNVideoProcessorTimeIntervalCadenceClassOnce sync.Once
)

func getVNVideoProcessorTimeIntervalCadenceClass() VNVideoProcessorTimeIntervalCadenceClass {
	_VNVideoProcessorTimeIntervalCadenceClassOnce.Do(func() {
		_VNVideoProcessorTimeIntervalCadenceClass = VNVideoProcessorTimeIntervalCadenceClass{class: objc.GetClass("VNVideoProcessorTimeIntervalCadence")}
	})
	return _VNVideoProcessorTimeIntervalCadenceClass
}

// GetVNVideoProcessorTimeIntervalCadenceClass returns the class object for VNVideoProcessorTimeIntervalCadence.
func GetVNVideoProcessorTimeIntervalCadenceClass() VNVideoProcessorTimeIntervalCadenceClass {
	return getVNVideoProcessorTimeIntervalCadenceClass()
}

type VNVideoProcessorTimeIntervalCadenceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNVideoProcessorTimeIntervalCadenceClass) Alloc() VNVideoProcessorTimeIntervalCadence {
	rv := objc.Send[VNVideoProcessorTimeIntervalCadence](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that defines a time-based cadence for processing a video stream.
//
// # Creating a Cadence
//
//   - [VNVideoProcessorTimeIntervalCadence.InitWithTimeInterval]: Creates a new time-based cadence with a time interval.
//
// # Inspecting the Time Interval
//
//   - [VNVideoProcessorTimeIntervalCadence.TimeInterval]: The time interval of the cadence.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence
type VNVideoProcessorTimeIntervalCadence struct {
	VNVideoProcessorCadence
}

// VNVideoProcessorTimeIntervalCadenceFromID constructs a [VNVideoProcessorTimeIntervalCadence] from an objc.ID.
//
// An object that defines a time-based cadence for processing a video stream.
func VNVideoProcessorTimeIntervalCadenceFromID(id objc.ID) VNVideoProcessorTimeIntervalCadence {
	return VNVideoProcessorTimeIntervalCadence{VNVideoProcessorCadence: VNVideoProcessorCadenceFromID(id)}
}
// NOTE: VNVideoProcessorTimeIntervalCadence adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNVideoProcessorTimeIntervalCadence] class.
//
// # Creating a Cadence
//
//   - [IVNVideoProcessorTimeIntervalCadence.InitWithTimeInterval]: Creates a new time-based cadence with a time interval.
//
// # Inspecting the Time Interval
//
//   - [IVNVideoProcessorTimeIntervalCadence.TimeInterval]: The time interval of the cadence.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence
type IVNVideoProcessorTimeIntervalCadence interface {
	IVNVideoProcessorCadence

	// Topic: Creating a Cadence

	// Creates a new time-based cadence with a time interval.
	InitWithTimeInterval(timeInterval float64) VNVideoProcessorTimeIntervalCadence

	// Topic: Inspecting the Time Interval

	// The time interval of the cadence.
	TimeInterval() float64
}

// Init initializes the instance.
func (v VNVideoProcessorTimeIntervalCadence) Init() VNVideoProcessorTimeIntervalCadence {
	rv := objc.Send[VNVideoProcessorTimeIntervalCadence](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VNVideoProcessorTimeIntervalCadence) Autorelease() VNVideoProcessorTimeIntervalCadence {
	rv := objc.Send[VNVideoProcessorTimeIntervalCadence](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNVideoProcessorTimeIntervalCadence creates a new VNVideoProcessorTimeIntervalCadence instance.
func NewVNVideoProcessorTimeIntervalCadence() VNVideoProcessorTimeIntervalCadence {
	class := getVNVideoProcessorTimeIntervalCadenceClass()
	rv := objc.Send[VNVideoProcessorTimeIntervalCadence](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new time-based cadence with a time interval.
//
// timeInterval: The time interval at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence/init(_:)
func NewVideoProcessorTimeIntervalCadenceWithTimeInterval(timeInterval float64) VNVideoProcessorTimeIntervalCadence {
	instance := getVNVideoProcessorTimeIntervalCadenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeInterval:"), timeInterval)
	return VNVideoProcessorTimeIntervalCadenceFromID(rv)
}

// Creates a new time-based cadence with a time interval.
//
// timeInterval: The time interval at which to process video.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence/init(_:)
func (v VNVideoProcessorTimeIntervalCadence) InitWithTimeInterval(timeInterval float64) VNVideoProcessorTimeIntervalCadence {
	rv := objc.Send[VNVideoProcessorTimeIntervalCadence](v.ID, objc.Sel("initWithTimeInterval:"), timeInterval)
	return rv
}

// The time interval of the cadence.
//
// See: https://developer.apple.com/documentation/Vision/VNVideoProcessor/TimeIntervalCadence/timeInterval
func (v VNVideoProcessorTimeIntervalCadence) TimeInterval() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("timeInterval"))
	return rv
}

