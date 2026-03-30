// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVideoPerformanceMetrics] class.
var (
	_AVVideoPerformanceMetricsClass     AVVideoPerformanceMetricsClass
	_AVVideoPerformanceMetricsClassOnce sync.Once
)

func getAVVideoPerformanceMetricsClass() AVVideoPerformanceMetricsClass {
	_AVVideoPerformanceMetricsClassOnce.Do(func() {
		_AVVideoPerformanceMetricsClass = AVVideoPerformanceMetricsClass{class: objc.GetClass("AVVideoPerformanceMetrics")}
	})
	return _AVVideoPerformanceMetricsClass
}

// GetAVVideoPerformanceMetricsClass returns the class object for AVVideoPerformanceMetrics.
func GetAVVideoPerformanceMetricsClass() AVVideoPerformanceMetricsClass {
	return getAVVideoPerformanceMetricsClass()
}

type AVVideoPerformanceMetricsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVideoPerformanceMetricsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoPerformanceMetricsClass) Alloc() AVVideoPerformanceMetrics {
	rv := objc.Send[AVVideoPerformanceMetrics](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides metrics related to video playback quality.
//
// # Inspecting metrics
//
//   - [AVVideoPerformanceMetrics.NumberOfCorruptedFrames]: The total number of corrupted frames.
//   - [AVVideoPerformanceMetrics.NumberOfDroppedFrames]: The total number of frames the system drops prior to decoding or from missing the display deadline.
//   - [AVVideoPerformanceMetrics.NumberOfFramesDisplayedUsingOptimizedCompositing]: The total number of full screen frames rendered in a special power-efficient mode that didn’t require compositing with other UI elements.
//   - [AVVideoPerformanceMetrics.TotalAccumulatedFrameDelay]: The accumulated amount of time between the prescribed presentation times of displayed video frames and their actual time of display.
//   - [AVVideoPerformanceMetrics.TotalNumberOfFrames]: The total number of frames that display if no frames drop.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics
type AVVideoPerformanceMetrics struct {
	objectivec.Object
}

// AVVideoPerformanceMetricsFromID constructs a [AVVideoPerformanceMetrics] from an objc.ID.
//
// An object that provides metrics related to video playback quality.
func AVVideoPerformanceMetricsFromID(id objc.ID) AVVideoPerformanceMetrics {
	return AVVideoPerformanceMetrics{objectivec.Object{ID: id}}
}

// NOTE: AVVideoPerformanceMetrics adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoPerformanceMetrics] class.
//
// # Inspecting metrics
//
//   - [IAVVideoPerformanceMetrics.NumberOfCorruptedFrames]: The total number of corrupted frames.
//   - [IAVVideoPerformanceMetrics.NumberOfDroppedFrames]: The total number of frames the system drops prior to decoding or from missing the display deadline.
//   - [IAVVideoPerformanceMetrics.NumberOfFramesDisplayedUsingOptimizedCompositing]: The total number of full screen frames rendered in a special power-efficient mode that didn’t require compositing with other UI elements.
//   - [IAVVideoPerformanceMetrics.TotalAccumulatedFrameDelay]: The accumulated amount of time between the prescribed presentation times of displayed video frames and their actual time of display.
//   - [IAVVideoPerformanceMetrics.TotalNumberOfFrames]: The total number of frames that display if no frames drop.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics
type IAVVideoPerformanceMetrics interface {
	objectivec.IObject

	// Topic: Inspecting metrics

	// The total number of corrupted frames.
	NumberOfCorruptedFrames() int
	// The total number of frames the system drops prior to decoding or from missing the display deadline.
	NumberOfDroppedFrames() int
	// The total number of full screen frames rendered in a special power-efficient mode that didn’t require compositing with other UI elements.
	NumberOfFramesDisplayedUsingOptimizedCompositing() int
	// The accumulated amount of time between the prescribed presentation times of displayed video frames and their actual time of display.
	TotalAccumulatedFrameDelay() float64
	// The total number of frames that display if no frames drop.
	TotalNumberOfFrames() int

	NumberOfCorruptedVideoFrames() uint64
	NumberOfDisplayCompositedVideoFrames() uint64
	NumberOfDroppedVideoFrames() uint64
	NumberOfNonDisplayCompositedVideoFrames() uint64
	TotalFrameDelay() float64
	TotalNumberOfVideoFrames() uint64
}

// Init initializes the instance.
func (v AVVideoPerformanceMetrics) Init() AVVideoPerformanceMetrics {
	rv := objc.Send[AVVideoPerformanceMetrics](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoPerformanceMetrics) Autorelease() AVVideoPerformanceMetrics {
	rv := objc.Send[AVVideoPerformanceMetrics](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoPerformanceMetrics creates a new AVVideoPerformanceMetrics instance.
func NewAVVideoPerformanceMetrics() AVVideoPerformanceMetrics {
	class := getAVVideoPerformanceMetricsClass()
	rv := objc.Send[AVVideoPerformanceMetrics](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The total number of corrupted frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfCorruptedFrames
func (v AVVideoPerformanceMetrics) NumberOfCorruptedFrames() int {
	rv := objc.Send[int](v.ID, objc.Sel("numberOfCorruptedFrames"))
	return rv
}

// The total number of frames the system drops prior to decoding or from
// missing the display deadline.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfDroppedFrames
func (v AVVideoPerformanceMetrics) NumberOfDroppedFrames() int {
	rv := objc.Send[int](v.ID, objc.Sel("numberOfDroppedFrames"))
	return rv
}

// The total number of full screen frames rendered in a special
// power-efficient mode that didn’t require compositing with other UI
// elements.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfFramesDisplayedUsingOptimizedCompositing
func (v AVVideoPerformanceMetrics) NumberOfFramesDisplayedUsingOptimizedCompositing() int {
	rv := objc.Send[int](v.ID, objc.Sel("numberOfFramesDisplayedUsingOptimizedCompositing"))
	return rv
}

// The accumulated amount of time between the prescribed presentation times of
// displayed video frames and their actual time of display.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalAccumulatedFrameDelay
func (v AVVideoPerformanceMetrics) TotalAccumulatedFrameDelay() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("totalAccumulatedFrameDelay"))
	return rv
}

// The total number of frames that display if no frames drop.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalNumberOfFrames
func (v AVVideoPerformanceMetrics) TotalNumberOfFrames() int {
	rv := objc.Send[int](v.ID, objc.Sel("totalNumberOfFrames"))
	return rv
}

// # Discussion
//
// [SPI] The total number of corrupted frames that have been detected. Same as
// numberOfCorruptedFrames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfCorruptedVideoFrames
func (v AVVideoPerformanceMetrics) NumberOfCorruptedVideoFrames() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("numberOfCorruptedVideoFrames"))
	return rv
}

// # Discussion
//
// [SPI] The total number of frames that were composited in detached mode.
// Same as numberOfFramesDisplayedUsingOptimizedCompositing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfDisplayCompositedVideoFrames
func (v AVVideoPerformanceMetrics) NumberOfDisplayCompositedVideoFrames() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("numberOfDisplayCompositedVideoFrames"))
	return rv
}

// # Discussion
//
// [SPI] The total number of frames dropped prior to decoding or dropped
// because a frame missed its display deadline. Same as numberOfDroppedFrames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfDroppedVideoFrames
func (v AVVideoPerformanceMetrics) NumberOfDroppedVideoFrames() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("numberOfDroppedVideoFrames"))
	return rv
}

// # Discussion
//
// [SPI] The total number of frames that were composited in undetached mode.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/numberOfNonDisplayCompositedVideoFrames
func (v AVVideoPerformanceMetrics) NumberOfNonDisplayCompositedVideoFrames() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("numberOfNonDisplayCompositedVideoFrames"))
	return rv
}

// # Discussion
//
// [SPI] The accumulated amount of time, in microseconds, between the
// prescribed presentation times of displayed video frames and the actual time
// at which they were displayed.
//
// This delay is always greater than or equal to zero since frames must never
// be displayed before their presentation time. Non-zero delays are a sign of
// playback jitter and possible loss of A/V sync. Same as
// totalAccumulatedFrameDelay.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalFrameDelay
func (v AVVideoPerformanceMetrics) TotalFrameDelay() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("totalFrameDelay"))
	return rv
}

// # Discussion
//
// [SPI] The total number of frames that would have been displayed if no
// frames are dropped. Same as totalNumberOfFrames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPerformanceMetrics/totalNumberOfVideoFrames
func (v AVVideoPerformanceMetrics) TotalNumberOfVideoFrames() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("totalNumberOfVideoFrames"))
	return rv
}
