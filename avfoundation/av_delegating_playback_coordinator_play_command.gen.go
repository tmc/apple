// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVDelegatingPlaybackCoordinatorPlayCommand] class.
var (
	_AVDelegatingPlaybackCoordinatorPlayCommandClass     AVDelegatingPlaybackCoordinatorPlayCommandClass
	_AVDelegatingPlaybackCoordinatorPlayCommandClassOnce sync.Once
)

func getAVDelegatingPlaybackCoordinatorPlayCommandClass() AVDelegatingPlaybackCoordinatorPlayCommandClass {
	_AVDelegatingPlaybackCoordinatorPlayCommandClassOnce.Do(func() {
		_AVDelegatingPlaybackCoordinatorPlayCommandClass = AVDelegatingPlaybackCoordinatorPlayCommandClass{class: objc.GetClass("AVDelegatingPlaybackCoordinatorPlayCommand")}
	})
	return _AVDelegatingPlaybackCoordinatorPlayCommandClass
}

// GetAVDelegatingPlaybackCoordinatorPlayCommandClass returns the class object for AVDelegatingPlaybackCoordinatorPlayCommand.
func GetAVDelegatingPlaybackCoordinatorPlayCommandClass() AVDelegatingPlaybackCoordinatorPlayCommandClass {
	return getAVDelegatingPlaybackCoordinatorPlayCommandClass()
}

type AVDelegatingPlaybackCoordinatorPlayCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVDelegatingPlaybackCoordinatorPlayCommandClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDelegatingPlaybackCoordinatorPlayCommandClass) Alloc() AVDelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlayCommand](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A command that indicates to play at a specific rate and time.
//
// # Accessing command details
//
//   - [AVDelegatingPlaybackCoordinatorPlayCommand.Rate]: A rate to use when starting playback.
//   - [AVDelegatingPlaybackCoordinatorPlayCommand.ItemTime]: A time in the item timeline to use to begin playback.
//   - [AVDelegatingPlaybackCoordinatorPlayCommand.HostClockTime]: A host clock time to use to begin playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand
type AVDelegatingPlaybackCoordinatorPlayCommand struct {
	AVDelegatingPlaybackCoordinatorPlaybackControlCommand
}

// AVDelegatingPlaybackCoordinatorPlayCommandFromID constructs a [AVDelegatingPlaybackCoordinatorPlayCommand] from an objc.ID.
//
// A command that indicates to play at a specific rate and time.
func AVDelegatingPlaybackCoordinatorPlayCommandFromID(id objc.ID) AVDelegatingPlaybackCoordinatorPlayCommand {
	return AVDelegatingPlaybackCoordinatorPlayCommand{AVDelegatingPlaybackCoordinatorPlaybackControlCommand: AVDelegatingPlaybackCoordinatorPlaybackControlCommandFromID(id)}
}
// NOTE: AVDelegatingPlaybackCoordinatorPlayCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDelegatingPlaybackCoordinatorPlayCommand] class.
//
// # Accessing command details
//
//   - [IAVDelegatingPlaybackCoordinatorPlayCommand.Rate]: A rate to use when starting playback.
//   - [IAVDelegatingPlaybackCoordinatorPlayCommand.ItemTime]: A time in the item timeline to use to begin playback.
//   - [IAVDelegatingPlaybackCoordinatorPlayCommand.HostClockTime]: A host clock time to use to begin playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand
type IAVDelegatingPlaybackCoordinatorPlayCommand interface {
	IAVDelegatingPlaybackCoordinatorPlaybackControlCommand

	// Topic: Accessing command details

	// A rate to use when starting playback.
	Rate() float32
	// A time in the item timeline to use to begin playback.
	ItemTime() uintptr
	// A host clock time to use to begin playback.
	HostClockTime() uintptr
}

// Init initializes the instance.
func (d AVDelegatingPlaybackCoordinatorPlayCommand) Init() AVDelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlayCommand](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDelegatingPlaybackCoordinatorPlayCommand) Autorelease() AVDelegatingPlaybackCoordinatorPlayCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlayCommand](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDelegatingPlaybackCoordinatorPlayCommand creates a new AVDelegatingPlaybackCoordinatorPlayCommand instance.
func NewAVDelegatingPlaybackCoordinatorPlayCommand() AVDelegatingPlaybackCoordinatorPlayCommand {
	class := getAVDelegatingPlaybackCoordinatorPlayCommandClass()
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlayCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A rate to use when starting playback.
//
// # Discussion
// 
// This value is always nonzero.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand/rate
func (d AVDelegatingPlaybackCoordinatorPlayCommand) Rate() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("rate"))
	return rv
}
// A time in the item timeline to use to begin playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand/itemTime
func (d AVDelegatingPlaybackCoordinatorPlayCommand) ItemTime() uintptr {
	rv := objc.Send[uintptr](d.ID, objc.Sel("itemTime"))
	return rv
}
// A host clock time to use to begin playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlayCommand/hostClockTime
func (d AVDelegatingPlaybackCoordinatorPlayCommand) HostClockTime() uintptr {
	rv := objc.Send[uintptr](d.ID, objc.Sel("hostClockTime"))
	return rv
}

