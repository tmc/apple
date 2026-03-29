// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVDelegatingPlaybackCoordinatorSeekCommand] class.
var (
	_AVDelegatingPlaybackCoordinatorSeekCommandClass     AVDelegatingPlaybackCoordinatorSeekCommandClass
	_AVDelegatingPlaybackCoordinatorSeekCommandClassOnce sync.Once
)

func getAVDelegatingPlaybackCoordinatorSeekCommandClass() AVDelegatingPlaybackCoordinatorSeekCommandClass {
	_AVDelegatingPlaybackCoordinatorSeekCommandClassOnce.Do(func() {
		_AVDelegatingPlaybackCoordinatorSeekCommandClass = AVDelegatingPlaybackCoordinatorSeekCommandClass{class: objc.GetClass("AVDelegatingPlaybackCoordinatorSeekCommand")}
	})
	return _AVDelegatingPlaybackCoordinatorSeekCommandClass
}

// GetAVDelegatingPlaybackCoordinatorSeekCommandClass returns the class object for AVDelegatingPlaybackCoordinatorSeekCommand.
func GetAVDelegatingPlaybackCoordinatorSeekCommandClass() AVDelegatingPlaybackCoordinatorSeekCommandClass {
	return getAVDelegatingPlaybackCoordinatorSeekCommandClass()
}

type AVDelegatingPlaybackCoordinatorSeekCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVDelegatingPlaybackCoordinatorSeekCommandClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDelegatingPlaybackCoordinatorSeekCommandClass) Alloc() AVDelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorSeekCommand](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A command that indicates to seek to a new time in the item timeline.
//
// # Accessing command details
//
//   - [AVDelegatingPlaybackCoordinatorSeekCommand.ShouldBufferInAnticipationOfPlayback]: A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.
//   - [AVDelegatingPlaybackCoordinatorSeekCommand.AnticipatedPlaybackRate]: The rate at which the coordinator expects playback to resume.
//   - [AVDelegatingPlaybackCoordinatorSeekCommand.ItemTime]: The time to seek to in the item timeline.
//   - [AVDelegatingPlaybackCoordinatorSeekCommand.CompletionDueDate]: The deadline by which the coordinator expects the delegate to handle the command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand
type AVDelegatingPlaybackCoordinatorSeekCommand struct {
	AVDelegatingPlaybackCoordinatorPlaybackControlCommand
}

// AVDelegatingPlaybackCoordinatorSeekCommandFromID constructs a [AVDelegatingPlaybackCoordinatorSeekCommand] from an objc.ID.
//
// A command that indicates to seek to a new time in the item timeline.
func AVDelegatingPlaybackCoordinatorSeekCommandFromID(id objc.ID) AVDelegatingPlaybackCoordinatorSeekCommand {
	return AVDelegatingPlaybackCoordinatorSeekCommand{AVDelegatingPlaybackCoordinatorPlaybackControlCommand: AVDelegatingPlaybackCoordinatorPlaybackControlCommandFromID(id)}
}
// NOTE: AVDelegatingPlaybackCoordinatorSeekCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDelegatingPlaybackCoordinatorSeekCommand] class.
//
// # Accessing command details
//
//   - [IAVDelegatingPlaybackCoordinatorSeekCommand.ShouldBufferInAnticipationOfPlayback]: A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.
//   - [IAVDelegatingPlaybackCoordinatorSeekCommand.AnticipatedPlaybackRate]: The rate at which the coordinator expects playback to resume.
//   - [IAVDelegatingPlaybackCoordinatorSeekCommand.ItemTime]: The time to seek to in the item timeline.
//   - [IAVDelegatingPlaybackCoordinatorSeekCommand.CompletionDueDate]: The deadline by which the coordinator expects the delegate to handle the command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand
type IAVDelegatingPlaybackCoordinatorSeekCommand interface {
	IAVDelegatingPlaybackCoordinatorPlaybackControlCommand

	// Topic: Accessing command details

	// A Boolean value that indicates whether the player starts buffering in anticipation of a request to begin playback.
	ShouldBufferInAnticipationOfPlayback() bool
	// The rate at which the coordinator expects playback to resume.
	AnticipatedPlaybackRate() float32
	// The time to seek to in the item timeline.
	ItemTime() coremedia.CMTime
	// The deadline by which the coordinator expects the delegate to handle the command.
	CompletionDueDate() foundation.INSDate
}

// Init initializes the instance.
func (d AVDelegatingPlaybackCoordinatorSeekCommand) Init() AVDelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorSeekCommand](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDelegatingPlaybackCoordinatorSeekCommand) Autorelease() AVDelegatingPlaybackCoordinatorSeekCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorSeekCommand](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDelegatingPlaybackCoordinatorSeekCommand creates a new AVDelegatingPlaybackCoordinatorSeekCommand instance.
func NewAVDelegatingPlaybackCoordinatorSeekCommand() AVDelegatingPlaybackCoordinatorSeekCommand {
	class := getAVDelegatingPlaybackCoordinatorSeekCommandClass()
	rv := objc.Send[AVDelegatingPlaybackCoordinatorSeekCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the player starts buffering in
// anticipation of a request to begin playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/shouldBufferInAnticipationOfPlayback
func (d AVDelegatingPlaybackCoordinatorSeekCommand) ShouldBufferInAnticipationOfPlayback() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shouldBufferInAnticipationOfPlayback"))
	return rv
}
// The rate at which the coordinator expects playback to resume.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/anticipatedPlaybackRate
func (d AVDelegatingPlaybackCoordinatorSeekCommand) AnticipatedPlaybackRate() float32 {
	rv := objc.Send[float32](d.ID, objc.Sel("anticipatedPlaybackRate"))
	return rv
}
// The time to seek to in the item timeline.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/itemTime
func (d AVDelegatingPlaybackCoordinatorSeekCommand) ItemTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](d.ID, objc.Sel("itemTime"))
	return coremedia.CMTime(rv)
}
// The deadline by which the coordinator expects the delegate to handle the
// command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekCommand/completionDueDate
func (d AVDelegatingPlaybackCoordinatorSeekCommand) CompletionDueDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("completionDueDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

