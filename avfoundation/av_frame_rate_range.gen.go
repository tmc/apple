// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVFrameRateRange] class.
var (
	_AVFrameRateRangeClass     AVFrameRateRangeClass
	_AVFrameRateRangeClassOnce sync.Once
)

func getAVFrameRateRangeClass() AVFrameRateRangeClass {
	_AVFrameRateRangeClassOnce.Do(func() {
		_AVFrameRateRangeClass = AVFrameRateRangeClass{class: objc.GetClass("AVFrameRateRange")}
	})
	return _AVFrameRateRangeClass
}

// GetAVFrameRateRangeClass returns the class object for AVFrameRateRange.
func GetAVFrameRateRangeClass() AVFrameRateRangeClass {
	return getAVFrameRateRangeClass()
}

type AVFrameRateRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVFrameRateRangeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFrameRateRangeClass) Alloc() AVFrameRateRange {
	rv := objc.Send[AVFrameRateRange](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An immutable type that represents a range of valid frame rates.
//
// # Overview
//
// An AVFrameRateRange object is immutable.
//
// An [AVCaptureDeviceFormat] object wraps a CMFormatDescription and expresses
// a range of valid video frame rates as an array of [AVFrameRateRange]
// objects.
//
// An [AVCaptureDevice] object uses [AVCaptureDeviceFormat] to describe the
// formats it supports and the currently-active format.
//
// # Accessing properties
//
//   - [AVFrameRateRange.MaxFrameDuration]: The maximum frame duration supported by the range.
//   - [AVFrameRateRange.MaxFrameRate]: The maximum frame rate supported by the range.
//   - [AVFrameRateRange.MinFrameDuration]: The minimum frame duration supported by the range.
//   - [AVFrameRateRange.MinFrameRate]: The minimum frame rate supported by the range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange
type AVFrameRateRange struct {
	objectivec.Object
}

// AVFrameRateRangeFromID constructs a [AVFrameRateRange] from an objc.ID.
//
// An immutable type that represents a range of valid frame rates.
func AVFrameRateRangeFromID(id objc.ID) AVFrameRateRange {
	return AVFrameRateRange{objectivec.Object{ID: id}}
}

// NOTE: AVFrameRateRange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFrameRateRange] class.
//
// # Accessing properties
//
//   - [IAVFrameRateRange.MaxFrameDuration]: The maximum frame duration supported by the range.
//   - [IAVFrameRateRange.MaxFrameRate]: The maximum frame rate supported by the range.
//   - [IAVFrameRateRange.MinFrameDuration]: The minimum frame duration supported by the range.
//   - [IAVFrameRateRange.MinFrameRate]: The minimum frame rate supported by the range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange
type IAVFrameRateRange interface {
	objectivec.IObject

	// Topic: Accessing properties

	// The maximum frame duration supported by the range.
	MaxFrameDuration() coremedia.CMTime
	// The maximum frame rate supported by the range.
	MaxFrameRate() float64
	// The minimum frame duration supported by the range.
	MinFrameDuration() coremedia.CMTime
	// The minimum frame rate supported by the range.
	MinFrameRate() float64
}

// Init initializes the instance.
func (f AVFrameRateRange) Init() AVFrameRateRange {
	rv := objc.Send[AVFrameRateRange](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFrameRateRange) Autorelease() AVFrameRateRange {
	rv := objc.Send[AVFrameRateRange](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFrameRateRange creates a new AVFrameRateRange instance.
func NewAVFrameRateRange() AVFrameRateRange {
	class := getAVFrameRateRangeClass()
	rv := objc.Send[AVFrameRateRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The maximum frame duration supported by the range.
//
// # Discussion
//
// This value is the reciprocal of [AVFrameRateRange.MinFrameRate], and
// expresses the minimum frame rate as a duration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/maxFrameDuration
func (f AVFrameRateRange) MaxFrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](f.ID, objc.Sel("maxFrameDuration"))
	return coremedia.CMTime(rv)
}

// The maximum frame rate supported by the range.
//
// # Discussion
//
// The frame is given in frames per second.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/maxFrameRate
func (f AVFrameRateRange) MaxFrameRate() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("maxFrameRate"))
	return rv
}

// The minimum frame duration supported by the range.
//
// # Discussion
//
// This value is the reciprocal of [AVFrameRateRange.MaxFrameRate], and
// expresses the maximum frame rate as a duration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/minFrameDuration
func (f AVFrameRateRange) MinFrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](f.ID, objc.Sel("minFrameDuration"))
	return coremedia.CMTime(rv)
}

// The minimum frame rate supported by the range.
//
// # Discussion
//
// The frame is given in frames per second.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFrameRateRange/minFrameRate
func (f AVFrameRateRange) MinFrameRate() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("minFrameRate"))
	return rv
}
