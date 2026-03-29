// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVPlayerPlaybackCoordinator] class.
var (
	_AVPlayerPlaybackCoordinatorClass     AVPlayerPlaybackCoordinatorClass
	_AVPlayerPlaybackCoordinatorClassOnce sync.Once
)

func getAVPlayerPlaybackCoordinatorClass() AVPlayerPlaybackCoordinatorClass {
	_AVPlayerPlaybackCoordinatorClassOnce.Do(func() {
		_AVPlayerPlaybackCoordinatorClass = AVPlayerPlaybackCoordinatorClass{class: objc.GetClass("AVPlayerPlaybackCoordinator")}
	})
	return _AVPlayerPlaybackCoordinatorClass
}

// GetAVPlayerPlaybackCoordinatorClass returns the class object for AVPlayerPlaybackCoordinator.
func GetAVPlayerPlaybackCoordinatorClass() AVPlayerPlaybackCoordinatorClass {
	return getAVPlayerPlaybackCoordinatorClass()
}

type AVPlayerPlaybackCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerPlaybackCoordinatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerPlaybackCoordinatorClass) Alloc() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A playback coordinator subclass that coordinates the playback of player
// objects in a connected group.
//
// # Overview
// 
// This object coordinates the state of [AVPlayer] objects. You don’t create
// an instance of the coordinator, but instead access the player’s instance
// through its [PlaybackCoordinator] property.
// 
// Use the standard interfaces of [AVPlayer] to control playback in your app.
// The coordinator automatically intercepts calls that affect transport
// control state, like [SetRateTimeAtHostTime], [Pause], and
// [SeekToTimeCompletionHandler], and propagates them to other participants in
// the group when appropriate. Similarly, the coordinator observes rate and
// time changes from other participants and imposes them on the player. If
// this occurs, the player item posts notifications that identify the
// originating participant.
// 
// [media-3839391]
// 
// This object may automatically suspend coordinated playback when a system
// state change causes the player’s [AVPlayerPlaybackCoordinator.TimeControlStatus] value to change from
// a playing state to a waiting or paused state. A suspension that begins
// because the player enters a waiting state due to an event like a network
// stall or interstitial playback, ends automatically when the player finishes
// waiting. However, if the system pauses playback due to a system state
// change, such as an audio session interruption, the suspension ends only
// after the player’s rate changes back to nonzero.
//
// # Accessing the player
//
//   - [AVPlayerPlaybackCoordinator.Player]: A player that participates in coordinated playback.
//
// # Configuring the delegate
//
//   - [AVPlayerPlaybackCoordinator.Delegate]: A delegate object for the playback coordinator.
//   - [AVPlayerPlaybackCoordinator.SetDelegate]
//
// # Managing coordination
//
//   - [AVPlayerPlaybackCoordinator.CoordinateUsingCoordinationMediumError]: Connects the playback coordinator to the coordination medium
//   - [AVPlayerPlaybackCoordinator.PlaybackCoordinationMedium]: The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator
type AVPlayerPlaybackCoordinator struct {
	AVPlaybackCoordinator
}

// AVPlayerPlaybackCoordinatorFromID constructs a [AVPlayerPlaybackCoordinator] from an objc.ID.
//
// A playback coordinator subclass that coordinates the playback of player
// objects in a connected group.
func AVPlayerPlaybackCoordinatorFromID(id objc.ID) AVPlayerPlaybackCoordinator {
	return AVPlayerPlaybackCoordinator{AVPlaybackCoordinator: AVPlaybackCoordinatorFromID(id)}
}
// NOTE: AVPlayerPlaybackCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerPlaybackCoordinator] class.
//
// # Accessing the player
//
//   - [IAVPlayerPlaybackCoordinator.Player]: A player that participates in coordinated playback.
//
// # Configuring the delegate
//
//   - [IAVPlayerPlaybackCoordinator.Delegate]: A delegate object for the playback coordinator.
//   - [IAVPlayerPlaybackCoordinator.SetDelegate]
//
// # Managing coordination
//
//   - [IAVPlayerPlaybackCoordinator.CoordinateUsingCoordinationMediumError]: Connects the playback coordinator to the coordination medium
//   - [IAVPlayerPlaybackCoordinator.PlaybackCoordinationMedium]: The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator
type IAVPlayerPlaybackCoordinator interface {
	IAVPlaybackCoordinator

	// Topic: Accessing the player

	// A player that participates in coordinated playback.
	Player() IAVPlayer

	// Topic: Configuring the delegate

	// A delegate object for the playback coordinator.
	Delegate() AVPlayerPlaybackCoordinatorDelegate
	SetDelegate(value AVPlayerPlaybackCoordinatorDelegate)

	// Topic: Managing coordination

	// Connects the playback coordinator to the coordination medium
	CoordinateUsingCoordinationMediumError(coordinationMedium IAVPlaybackCoordinationMedium) (bool, error)
	// The AVPlaybackCoordinationMedium this playback coordinator is connected to.
	PlaybackCoordinationMedium() IAVPlaybackCoordinationMedium

	// The playback coordinator for the player.
	PlaybackCoordinator() IAVPlayerPlaybackCoordinator
	SetPlaybackCoordinator(value IAVPlayerPlaybackCoordinator)
	// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
	TimeControlStatus() AVPlayerTimeControlStatus
	SetTimeControlStatus(value AVPlayerTimeControlStatus)
}

// Init initializes the instance.
func (p AVPlayerPlaybackCoordinator) Init() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerPlaybackCoordinator) Autorelease() AVPlayerPlaybackCoordinator {
	rv := objc.Send[AVPlayerPlaybackCoordinator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerPlaybackCoordinator creates a new AVPlayerPlaybackCoordinator instance.
func NewAVPlayerPlaybackCoordinator() AVPlayerPlaybackCoordinator {
	class := getAVPlayerPlaybackCoordinatorClass()
	rv := objc.Send[AVPlayerPlaybackCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Connects the playback coordinator to the coordination medium
//
// coordinationMedium: The coordination medium the playback coordinator connects to. If NULL, the
// playback coordinator disconnects from any existing coordination medium.
//
// # Discussion
// 
// This connects the playback coordinator to a coordination medium to enable
// sending and receiving messages from other connected playback coordinators.
// If the coordination medium is non-NULL, this will connect the playback
// coordinator to the specified coordination medium. If the coordination
// medium is set to NULL, this will disconnect the playback coordinator from
// the playback coordination medium. The player will no longer be coordinated
// with the other players connected to the coordination medium. The playback
// coordinator can either only coordinate with local players through an
// AVPlaybackCoordinationMedium or coordinate with a remote group session
// through the `coordinateWithSession` API. If the client attempts to connect
// to an AVPlaybackCoordinationMedium while already connected to a group
// session, this method will populate the outError parameter If the playback
// coordinator successfully connects to the coordination medium or disconnects
// from a coordination medium, the `outError` parameter will be nil. If the
// playback coordinator fails to connect to the specified coordination medium,
// the `outError` parameter will describe what went wrong.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/coordinate(using:)
func (p AVPlayerPlaybackCoordinator) CoordinateUsingCoordinationMediumError(coordinationMedium IAVPlaybackCoordinationMedium) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("coordinateUsingCoordinationMedium:error:"), coordinationMedium, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("coordinateUsingCoordinationMedium:error: returned NO with nil NSError")
	}
	return rv, nil

}

// A player that participates in coordinated playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/player
func (p AVPlayerPlaybackCoordinator) Player() IAVPlayer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("player"))
	return AVPlayerFromID(objc.ID(rv))
}
// A delegate object for the playback coordinator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/delegate
func (p AVPlayerPlaybackCoordinator) Delegate() AVPlayerPlaybackCoordinatorDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return AVPlayerPlaybackCoordinatorDelegateObjectFromID(rv)
}
func (p AVPlayerPlaybackCoordinator) SetDelegate(value AVPlayerPlaybackCoordinatorDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}
// The AVPlaybackCoordinationMedium this playback coordinator is connected to.
//
// # Discussion
// 
// This is the AVPlaybackCoordinationMedium the playback coordinator is
// connected to. If not NULL, the playback coordinator is connected to the
// specified coordination medium. The playback coordinator is not available to
// coordinate with a group session. If NULL, the playback coordinator is not
// connected to any playback coordination medium. The playback coordinator is
// available to coordinate with a group session through the
// `coordinateWithSession` API.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerPlaybackCoordinator/playbackCoordinationMedium
func (p AVPlayerPlaybackCoordinator) PlaybackCoordinationMedium() IAVPlaybackCoordinationMedium {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackCoordinationMedium"))
	return AVPlaybackCoordinationMediumFromID(objc.ID(rv))
}
// The playback coordinator for the player.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayer/playbackcoordinator
func (p AVPlayerPlaybackCoordinator) PlaybackCoordinator() IAVPlayerPlaybackCoordinator {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackCoordinator"))
	return AVPlayerPlaybackCoordinatorFromID(objc.ID(rv))
}
func (p AVPlayerPlaybackCoordinator) SetPlaybackCoordinator(value IAVPlayerPlaybackCoordinator) {
	objc.Send[struct{}](p.ID, objc.Sel("setPlaybackCoordinator:"), value)
}
// A value that indicates whether playback is in progress, paused
// indefinitely, or waiting for network conditions to improve.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p AVPlayerPlaybackCoordinator) TimeControlStatus() AVPlayerTimeControlStatus {
	rv := objc.Send[AVPlayerTimeControlStatus](p.ID, objc.Sel("timeControlStatus"))
	return AVPlayerTimeControlStatus(rv)
}
func (p AVPlayerPlaybackCoordinator) SetTimeControlStatus(value AVPlayerTimeControlStatus) {
	objc.Send[struct{}](p.ID, objc.Sel("setTimeControlStatus:"), value)
}

