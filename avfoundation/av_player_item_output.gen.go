// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemOutput] class.
var (
	_AVPlayerItemOutputClass     AVPlayerItemOutputClass
	_AVPlayerItemOutputClassOnce sync.Once
)

func getAVPlayerItemOutputClass() AVPlayerItemOutputClass {
	_AVPlayerItemOutputClassOnce.Do(func() {
		_AVPlayerItemOutputClass = AVPlayerItemOutputClass{class: objc.GetClass("AVPlayerItemOutput")}
	})
	return _AVPlayerItemOutputClass
}

// GetAVPlayerItemOutputClass returns the class object for AVPlayerItemOutput.
func GetAVPlayerItemOutputClass() AVPlayerItemOutputClass {
	return getAVPlayerItemOutputClass()
}

type AVPlayerItemOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemOutputClass) Alloc() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines the common interface to output media data
// from a player item.
//
// # Overview
//
// This class provides basic methods for converting time values to the
// timebase of the item. It also provides an option to suppress rendering of
// the output associated with the specific instance of this class.
//
// # Time conversion
//
//   - [AVPlayerItemOutput.ItemTimeForHostTime]: Converts a host time, specified in seconds, to the item’s timebase.
//   - [AVPlayerItemOutput.ItemTimeForMachAbsoluteTime]: Converts a Mach host time to the item’s timebase.
//   - [AVPlayerItemOutput.ItemTimeForCVTimeStamp]: Converts a Core Video timestamp to the item’s timebase.
//
// # Configuring the playback options
//
//   - [AVPlayerItemOutput.SuppressesPlayerRendering]: A Boolean value that indicates whether the player object renders the receiver’s output.
//   - [AVPlayerItemOutput.SetSuppressesPlayerRendering]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput
type AVPlayerItemOutput struct {
	objectivec.Object
}

// AVPlayerItemOutputFromID constructs a [AVPlayerItemOutput] from an objc.ID.
//
// An abstract class that defines the common interface to output media data
// from a player item.
func AVPlayerItemOutputFromID(id objc.ID) AVPlayerItemOutput {
	return AVPlayerItemOutput{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerItemOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemOutput] class.
//
// # Time conversion
//
//   - [IAVPlayerItemOutput.ItemTimeForHostTime]: Converts a host time, specified in seconds, to the item’s timebase.
//   - [IAVPlayerItemOutput.ItemTimeForMachAbsoluteTime]: Converts a Mach host time to the item’s timebase.
//   - [IAVPlayerItemOutput.ItemTimeForCVTimeStamp]: Converts a Core Video timestamp to the item’s timebase.
//
// # Configuring the playback options
//
//   - [IAVPlayerItemOutput.SuppressesPlayerRendering]: A Boolean value that indicates whether the player object renders the receiver’s output.
//   - [IAVPlayerItemOutput.SetSuppressesPlayerRendering]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput
type IAVPlayerItemOutput interface {
	objectivec.IObject

	// Topic: Time conversion

	// Converts a host time, specified in seconds, to the item’s timebase.
	ItemTimeForHostTime(hostTimeInSeconds float64) coremedia.CMTime
	// Converts a Mach host time to the item’s timebase.
	ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) coremedia.CMTime
	// Converts a Core Video timestamp to the item’s timebase.
	ItemTimeForCVTimeStamp(timestamp corevideo.CVTimeStamp) coremedia.CMTime

	// Topic: Configuring the playback options

	// A Boolean value that indicates whether the player object renders the receiver’s output.
	SuppressesPlayerRendering() bool
	SetSuppressesPlayerRendering(value bool)
}

// Init initializes the instance.
func (p AVPlayerItemOutput) Init() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemOutput) Autorelease() AVPlayerItemOutput {
	rv := objc.Send[AVPlayerItemOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemOutput creates a new AVPlayerItemOutput instance.
func NewAVPlayerItemOutput() AVPlayerItemOutput {
	class := getAVPlayerItemOutputClass()
	rv := objc.Send[AVPlayerItemOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Converts a host time, specified in seconds, to the item’s timebase.
//
// hostTimeInSeconds: A host time value, specified in seconds. For example, you might specify the
// time value returned by the [CACurrentMediaTime()] function or the timestamp
// from a [CADisplayLink] object for this parameter.
//
// # Return Value
//
// The equivalent time in the item’s timebase.
//
// # Discussion
//
// The timestamp associated with a [CADisplayLink] object represents the time
// of the most recent screen refresh, which is usually a time in the past. If
// you want to find the time associated with the next screen refresh, you need
// to increment the timestamp by the value in the display link’s `duration`
// property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(forHostTime:)
//
// [CACurrentMediaTime()]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
// [CADisplayLink]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
//
// [CADisplayLink]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
func (p AVPlayerItemOutput) ItemTimeForHostTime(hostTimeInSeconds float64) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("itemTimeForHostTime:"), hostTimeInSeconds)
	return coremedia.CMTime(rv)
}

// Converts a Mach host time to the item’s timebase.
//
// machAbsoluteTime: The Mach host time to convert. You typically retrieve this value using the
// `mach_absolute_time` function.
//
// # Return Value
//
// The equivalent time in the item’s timebase.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(forMachAbsoluteTime:)
func (p AVPlayerItemOutput) ItemTimeForMachAbsoluteTime(machAbsoluteTime int64) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("itemTimeForMachAbsoluteTime:"), machAbsoluteTime)
	return coremedia.CMTime(rv)
}

// Converts a Core Video timestamp to the item’s timebase.
//
// timestamp: A timestamp value provided by the Core Video framework.
//
// # Return Value
//
// The equivalent time in the item’s timebase.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/itemTime(for:)
func (p AVPlayerItemOutput) ItemTimeForCVTimeStamp(timestamp corevideo.CVTimeStamp) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("itemTimeForCVTimeStamp:"), timestamp)
	return coremedia.CMTime(rv)
}

// A Boolean value that indicates whether the player object renders the
// receiver’s output.
//
// # Discussion
//
// When the value of this property is false (the default), the player object
// handles the rendering of the receiver’s associated output. Change the
// value of this property to true to suppress the rendering of the media data
// associated with this object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutput/suppressesPlayerRendering
func (p AVPlayerItemOutput) SuppressesPlayerRendering() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("suppressesPlayerRendering"))
	return rv
}
func (p AVPlayerItemOutput) SetSuppressesPlayerRendering(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setSuppressesPlayerRendering:"), value)
}
