// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVCaptionConversionTimeRangeAdjustment] class.
var (
	_AVCaptionConversionTimeRangeAdjustmentClass     AVCaptionConversionTimeRangeAdjustmentClass
	_AVCaptionConversionTimeRangeAdjustmentClassOnce sync.Once
)

func getAVCaptionConversionTimeRangeAdjustmentClass() AVCaptionConversionTimeRangeAdjustmentClass {
	_AVCaptionConversionTimeRangeAdjustmentClassOnce.Do(func() {
		_AVCaptionConversionTimeRangeAdjustmentClass = AVCaptionConversionTimeRangeAdjustmentClass{class: objc.GetClass("AVCaptionConversionTimeRangeAdjustment")}
	})
	return _AVCaptionConversionTimeRangeAdjustmentClass
}

// GetAVCaptionConversionTimeRangeAdjustmentClass returns the class object for AVCaptionConversionTimeRangeAdjustment.
func GetAVCaptionConversionTimeRangeAdjustmentClass() AVCaptionConversionTimeRangeAdjustmentClass {
	return getAVCaptionConversionTimeRangeAdjustmentClass()
}

type AVCaptionConversionTimeRangeAdjustmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionConversionTimeRangeAdjustmentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionConversionTimeRangeAdjustmentClass) Alloc() AVCaptionConversionTimeRangeAdjustment {
	rv := objc.Send[AVCaptionConversionTimeRangeAdjustment](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an adjustment to the time range of one or more
// captions.
//
// # Accessing time offsets
//
//   - [AVCaptionConversionTimeRangeAdjustment.StartTimeOffset]: The time value by which the system offsets the start times of captions to correct a problem.
//   - [AVCaptionConversionTimeRangeAdjustment.DurationOffset]: The time value by which the system offsets the durations of captions to correct a problem.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment
type AVCaptionConversionTimeRangeAdjustment struct {
	AVCaptionConversionAdjustment
}

// AVCaptionConversionTimeRangeAdjustmentFromID constructs a [AVCaptionConversionTimeRangeAdjustment] from an objc.ID.
//
// An object that describes an adjustment to the time range of one or more
// captions.
func AVCaptionConversionTimeRangeAdjustmentFromID(id objc.ID) AVCaptionConversionTimeRangeAdjustment {
	return AVCaptionConversionTimeRangeAdjustment{AVCaptionConversionAdjustment: AVCaptionConversionAdjustmentFromID(id)}
}

// NOTE: AVCaptionConversionTimeRangeAdjustment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionConversionTimeRangeAdjustment] class.
//
// # Accessing time offsets
//
//   - [IAVCaptionConversionTimeRangeAdjustment.StartTimeOffset]: The time value by which the system offsets the start times of captions to correct a problem.
//   - [IAVCaptionConversionTimeRangeAdjustment.DurationOffset]: The time value by which the system offsets the durations of captions to correct a problem.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment
type IAVCaptionConversionTimeRangeAdjustment interface {
	IAVCaptionConversionAdjustment

	// Topic: Accessing time offsets

	// The time value by which the system offsets the start times of captions to correct a problem.
	StartTimeOffset() coremedia.CMTime
	// The time value by which the system offsets the durations of captions to correct a problem.
	DurationOffset() coremedia.CMTime
}

// Init initializes the instance.
func (c AVCaptionConversionTimeRangeAdjustment) Init() AVCaptionConversionTimeRangeAdjustment {
	rv := objc.Send[AVCaptionConversionTimeRangeAdjustment](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionConversionTimeRangeAdjustment) Autorelease() AVCaptionConversionTimeRangeAdjustment {
	rv := objc.Send[AVCaptionConversionTimeRangeAdjustment](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionConversionTimeRangeAdjustment creates a new AVCaptionConversionTimeRangeAdjustment instance.
func NewAVCaptionConversionTimeRangeAdjustment() AVCaptionConversionTimeRangeAdjustment {
	class := getAVCaptionConversionTimeRangeAdjustmentClass()
	rv := objc.Send[AVCaptionConversionTimeRangeAdjustment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The time value by which the system offsets the start times of captions to
// correct a problem.
//
// # Discussion
//
// The value may any numeric value, positive, negative, or zero.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment/startTimeOffset
func (c AVCaptionConversionTimeRangeAdjustment) StartTimeOffset() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("startTimeOffset"))
	return coremedia.CMTime(rv)
}

// The time value by which the system offsets the durations of captions to
// correct a problem.
//
// # Discussion
//
// The value may any numeric value, positive, negative, or zero.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionTimeRangeAdjustment/durationOffset
func (c AVCaptionConversionTimeRangeAdjustment) DurationOffset() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("durationOffset"))
	return coremedia.CMTime(rv)
}
