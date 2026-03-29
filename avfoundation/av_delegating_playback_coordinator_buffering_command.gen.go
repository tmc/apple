// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVDelegatingPlaybackCoordinatorBufferingCommand] class.
var (
	_AVDelegatingPlaybackCoordinatorBufferingCommandClass     AVDelegatingPlaybackCoordinatorBufferingCommandClass
	_AVDelegatingPlaybackCoordinatorBufferingCommandClassOnce sync.Once
)

func getAVDelegatingPlaybackCoordinatorBufferingCommandClass() AVDelegatingPlaybackCoordinatorBufferingCommandClass {
	_AVDelegatingPlaybackCoordinatorBufferingCommandClassOnce.Do(func() {
		_AVDelegatingPlaybackCoordinatorBufferingCommandClass = AVDelegatingPlaybackCoordinatorBufferingCommandClass{class: objc.GetClass("AVDelegatingPlaybackCoordinatorBufferingCommand")}
	})
	return _AVDelegatingPlaybackCoordinatorBufferingCommandClass
}

// GetAVDelegatingPlaybackCoordinatorBufferingCommandClass returns the class object for AVDelegatingPlaybackCoordinatorBufferingCommand.
func GetAVDelegatingPlaybackCoordinatorBufferingCommandClass() AVDelegatingPlaybackCoordinatorBufferingCommandClass {
	return getAVDelegatingPlaybackCoordinatorBufferingCommandClass()
}

type AVDelegatingPlaybackCoordinatorBufferingCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVDelegatingPlaybackCoordinatorBufferingCommandClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDelegatingPlaybackCoordinatorBufferingCommandClass) Alloc() AVDelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorBufferingCommand](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A command that indicates to start buffering data in preparation for
// playback.
//
// # Overview
// 
// When your app receives this command, update its user interface to indicate
// that playback is buffering.
//
// # Accessing command details
//
//   - [AVDelegatingPlaybackCoordinatorBufferingCommand.AnticipatedPlaybackRate]: The rate at which the coordinator expects the current item to play.
//   - [AVDelegatingPlaybackCoordinatorBufferingCommand.CompletionDueDate]: The deadline by which the coordinator expects the delegate to complete execution of a command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand
type AVDelegatingPlaybackCoordinatorBufferingCommand struct {
	AVDelegatingPlaybackCoordinatorPlaybackControlCommand
}

// AVDelegatingPlaybackCoordinatorBufferingCommandFromID constructs a [AVDelegatingPlaybackCoordinatorBufferingCommand] from an objc.ID.
//
// A command that indicates to start buffering data in preparation for
// playback.
func AVDelegatingPlaybackCoordinatorBufferingCommandFromID(id objc.ID) AVDelegatingPlaybackCoordinatorBufferingCommand {
	return AVDelegatingPlaybackCoordinatorBufferingCommand{AVDelegatingPlaybackCoordinatorPlaybackControlCommand: AVDelegatingPlaybackCoordinatorPlaybackControlCommandFromID(id)}
}
// NOTE: AVDelegatingPlaybackCoordinatorBufferingCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDelegatingPlaybackCoordinatorBufferingCommand] class.
//
// # Accessing command details
//
//   - [IAVDelegatingPlaybackCoordinatorBufferingCommand.AnticipatedPlaybackRate]: The rate at which the coordinator expects the current item to play.
//   - [IAVDelegatingPlaybackCoordinatorBufferingCommand.CompletionDueDate]: The deadline by which the coordinator expects the delegate to complete execution of a command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand
type IAVDelegatingPlaybackCoordinatorBufferingCommand interface {
	IAVDelegatingPlaybackCoordinatorPlaybackControlCommand

	// Topic: Accessing command details

	// The rate at which the coordinator expects the current item to play.
	AnticipatedPlaybackRate() float32
	// The deadline by which the coordinator expects the delegate to complete execution of a command.
	CompletionDueDate() foundation.INSDate
}

// Init initializes the instance.
func (d AVDelegatingPlaybackCoordinatorBufferingCommand) Init() AVDelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorBufferingCommand](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDelegatingPlaybackCoordinatorBufferingCommand) Autorelease() AVDelegatingPlaybackCoordinatorBufferingCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorBufferingCommand](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDelegatingPlaybackCoordinatorBufferingCommand creates a new AVDelegatingPlaybackCoordinatorBufferingCommand instance.
func NewAVDelegatingPlaybackCoordinatorBufferingCommand() AVDelegatingPlaybackCoordinatorBufferingCommand {
	class := getAVDelegatingPlaybackCoordinatorBufferingCommandClass()
	rv := objc.Send[AVDelegatingPlaybackCoordinatorBufferingCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The rate at which the coordinator expects the current item to play.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand/anticipatedPlaybackRate
func (d AVDelegatingPlaybackCoordinatorBufferingCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}
// The deadline by which the coordinator expects the delegate to complete
// execution of a command.
//
// # Discussion
// 
// A command that expects buffering in preparation for playback requires that
// the delegate call the command’s completion handler by the deadline. The
// delegate needs to complete the command by this date to keep up with the
// group. Alternatively, have the delegate begin a stall recovery suspension,
// and communicate that state to the other participants.
// 
// Completing the command after this date means that the coordinator likely
// sends a play command that isn’t for the current state.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorBufferingCommand/completionDueDate
func (d AVDelegatingPlaybackCoordinatorBufferingCommand) CompletionDueDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("completionDueDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

