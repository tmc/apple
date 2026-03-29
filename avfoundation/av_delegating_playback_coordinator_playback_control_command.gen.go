// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVDelegatingPlaybackCoordinatorPlaybackControlCommand] class.
var (
	_AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass     AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass
	_AVDelegatingPlaybackCoordinatorPlaybackControlCommandClassOnce sync.Once
)

func getAVDelegatingPlaybackCoordinatorPlaybackControlCommandClass() AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass {
	_AVDelegatingPlaybackCoordinatorPlaybackControlCommandClassOnce.Do(func() {
		_AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass = AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass{class: objc.GetClass("AVDelegatingPlaybackCoordinatorPlaybackControlCommand")}
	})
	return _AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass
}

// GetAVDelegatingPlaybackCoordinatorPlaybackControlCommandClass returns the class object for AVDelegatingPlaybackCoordinatorPlaybackControlCommand.
func GetAVDelegatingPlaybackCoordinatorPlaybackControlCommandClass() AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass {
	return getAVDelegatingPlaybackCoordinatorPlaybackControlCommandClass()
}

type AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDelegatingPlaybackCoordinatorPlaybackControlCommandClass) Alloc() AVDelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlaybackControlCommand](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An abstract superclass for playback commands.
//
// # Overview
// 
// Playback commands inherit state that identifies their originator and
// applicable item.
//
// # Accessing command details
//
//   - [AVDelegatingPlaybackCoordinatorPlaybackControlCommand.ExpectedCurrentItemIdentifier]: An item identifier the coordinator issues the command for.
//   - [AVDelegatingPlaybackCoordinatorPlaybackControlCommand.Originator]: The participant that causes the coordinator to issue the command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand
type AVDelegatingPlaybackCoordinatorPlaybackControlCommand struct {
	objectivec.Object
}

// AVDelegatingPlaybackCoordinatorPlaybackControlCommandFromID constructs a [AVDelegatingPlaybackCoordinatorPlaybackControlCommand] from an objc.ID.
//
// An abstract superclass for playback commands.
func AVDelegatingPlaybackCoordinatorPlaybackControlCommandFromID(id objc.ID) AVDelegatingPlaybackCoordinatorPlaybackControlCommand {
	return AVDelegatingPlaybackCoordinatorPlaybackControlCommand{objectivec.Object{ID: id}}
}
// NOTE: AVDelegatingPlaybackCoordinatorPlaybackControlCommand adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDelegatingPlaybackCoordinatorPlaybackControlCommand] class.
//
// # Accessing command details
//
//   - [IAVDelegatingPlaybackCoordinatorPlaybackControlCommand.ExpectedCurrentItemIdentifier]: An item identifier the coordinator issues the command for.
//   - [IAVDelegatingPlaybackCoordinatorPlaybackControlCommand.Originator]: The participant that causes the coordinator to issue the command.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand
type IAVDelegatingPlaybackCoordinatorPlaybackControlCommand interface {
	objectivec.IObject

	// Topic: Accessing command details

	// An item identifier the coordinator issues the command for.
	ExpectedCurrentItemIdentifier() string
	// The participant that causes the coordinator to issue the command.
	Originator() IAVCoordinatedPlaybackParticipant
}

// Init initializes the instance.
func (d AVDelegatingPlaybackCoordinatorPlaybackControlCommand) Init() AVDelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlaybackControlCommand](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDelegatingPlaybackCoordinatorPlaybackControlCommand) Autorelease() AVDelegatingPlaybackCoordinatorPlaybackControlCommand {
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlaybackControlCommand](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDelegatingPlaybackCoordinatorPlaybackControlCommand creates a new AVDelegatingPlaybackCoordinatorPlaybackControlCommand instance.
func NewAVDelegatingPlaybackCoordinatorPlaybackControlCommand() AVDelegatingPlaybackCoordinatorPlaybackControlCommand {
	class := getAVDelegatingPlaybackCoordinatorPlaybackControlCommandClass()
	rv := objc.Send[AVDelegatingPlaybackCoordinatorPlaybackControlCommand](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An item identifier the coordinator issues the command for.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand/expectedCurrentItemIdentifier
func (d AVDelegatingPlaybackCoordinatorPlaybackControlCommand) ExpectedCurrentItemIdentifier() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("expectedCurrentItemIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The participant that causes the coordinator to issue the command.
//
// # Discussion
// 
// Only commands that the system issues on behalf of another participant
// contain an originator. Local commands to coordinate rate change, or those
// that originate from a call to
// [ReapplyCurrentItemStateToPlaybackControlDelegate], don’t.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorPlaybackControlCommand/originator
func (d AVDelegatingPlaybackCoordinatorPlaybackControlCommand) Originator() IAVCoordinatedPlaybackParticipant {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("originator"))
	return AVCoordinatedPlaybackParticipantFromID(objc.ID(rv))
}

