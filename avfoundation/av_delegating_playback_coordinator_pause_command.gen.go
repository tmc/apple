// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVDelegatingPlaybackCoordinatorPauseCommand] class.
var (
	_AVDelegatingPlaybackCoordinatorPauseCommandClass     AVDelegatingPlaybackCoordinatorPauseCommandClass
	_AVDelegatingPlaybackCoordinatorPauseCommandClassOnce sync.Once
)

func getAVDelegatingPlaybackCoordinatorPauseCommandClass() AVDelegatingPlaybackCoordinatorPauseCommandClass {
	_AVDelegatingPlaybackCoordinatorPauseCommandClassOnce.Do(func() {
		_AVDelegatingPlaybackCoordinatorPauseCommandClass = AVDelegatingPlaybackCoordinatorPauseCommandClass{class: objc.GetClass("AVDelegatingPlaybackCoordinatorPauseCommand")}
	})
	return _AVDelegatingPlaybackCoordinatorPauseCommandClass
}

// GetAVDelegatingPlaybackCoordinatorPauseCommandClass returns the class object for AVDelegatingPlaybackCoordinatorPauseCommand.
func GetAVDelegatingPlaybackCoordinatorPauseCommandClass() AVDelegatingPlaybackCoordinatorPauseCommandClass {
	return getAVDelegatingPlaybackCoordinatorPauseCommandClass()
}

type AVDelegatingPlaybackCoordinatorPauseCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVDelegatingPlaybackCoordinatorPauseCommandClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDelegatingPlaybackCoordinatorPauseCommandClass) Alloc() AVDelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPauseCommand](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A command that indicates to pause playback.
//
// # Accessing command details
//
//   - [AVDelegatingPlaybackCoordinatorPauseCommand.ShouldBufferInAnticipationOfPlayback]: A Boolean value that indicates whether the player starts buffering in preparation for a request to begin playback.
//   - [AVDelegatingPlaybackCoordinatorPauseCommand.AnticipatedPlaybackRate]: The rate at which the coordinator expects the current item to play.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand
type AVDelegatingPlaybackCoordinatorPauseCommand struct {
	AVDelegatingPlaybackCoordinatorPlaybackControlCommand
}

// AVDelegatingPlaybackCoordinatorPauseCommandFromID constructs a [AVDelegatingPlaybackCoordinatorPauseCommand] from an objc.ID.
//
// A command that indicates to pause playback.
func AVDelegatingPlaybackCoordinatorPauseCommandFromID(id objc.ID) AVDelegatingPlaybackCoordinatorPauseCommand {
	return AVDelegatingPlaybackCoordinatorPauseCommand{AVDelegatingPlaybackCoordinatorPlaybackControlCommand: AVDelegatingPlaybackCoordinatorPlaybackControlCommandFromID(id)}
}

// NOTE: AVDelegatingPlaybackCoordinatorPauseCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDelegatingPlaybackCoordinatorPauseCommand] class.
//
// # Accessing command details
//
//   - [IAVDelegatingPlaybackCoordinatorPauseCommand.ShouldBufferInAnticipationOfPlayback]: A Boolean value that indicates whether the player starts buffering in preparation for a request to begin playback.
//   - [IAVDelegatingPlaybackCoordinatorPauseCommand.AnticipatedPlaybackRate]: The rate at which the coordinator expects the current item to play.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand
type IAVDelegatingPlaybackCoordinatorPauseCommand interface {
	IAVDelegatingPlaybackCoordinatorPlaybackControlCommand

	// Topic: Accessing command details

	// A Boolean value that indicates whether the player starts buffering in preparation for a request to begin playback.
	ShouldBufferInAnticipationOfPlayback() bool
	// The rate at which the coordinator expects the current item to play.
	AnticipatedPlaybackRate() float32
}

// Init initializes the instance.
func (d AVDelegatingPlaybackCoordinatorPauseCommand) Init() AVDelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPauseCommand](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDelegatingPlaybackCoordinatorPauseCommand) Autorelease() AVDelegatingPlaybackCoordinatorPauseCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPauseCommand](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDelegatingPlaybackCoordinatorPauseCommand creates a new AVDelegatingPlaybackCoordinatorPauseCommand instance.
func NewAVDelegatingPlaybackCoordinatorPauseCommand() AVDelegatingPlaybackCoordinatorPauseCommand {
	class := getAVDelegatingPlaybackCoordinatorPauseCommandClass()
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPauseCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the player starts buffering in
// preparation for a request to begin playback.
//
// # Discussion
//
// A true value indicates that a participant player requests starting playback
// at the [AnticipatedPlaybackRate] value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand/shouldBufferInAnticipationOfPlayback
func (d AVDelegatingPlaybackCoordinatorPauseCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}

// The rate at which the coordinator expects the current item to play.
//
// # Discussion
//
// Consider this command complete after the player is ready to start playback
// at the indicated rate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPauseCommand/anticipatedPlaybackRate
func (d AVDelegatingPlaybackCoordinatorPauseCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}
