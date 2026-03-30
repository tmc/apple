// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCoordinatedPlaybackParticipant] class.
var (
	_AVCoordinatedPlaybackParticipantClass     AVCoordinatedPlaybackParticipantClass
	_AVCoordinatedPlaybackParticipantClassOnce sync.Once
)

func getAVCoordinatedPlaybackParticipantClass() AVCoordinatedPlaybackParticipantClass {
	_AVCoordinatedPlaybackParticipantClassOnce.Do(func() {
		_AVCoordinatedPlaybackParticipantClass = AVCoordinatedPlaybackParticipantClass{class: objc.GetClass("AVCoordinatedPlaybackParticipant")}
	})
	return _AVCoordinatedPlaybackParticipantClass
}

// GetAVCoordinatedPlaybackParticipantClass returns the class object for AVCoordinatedPlaybackParticipant.
func GetAVCoordinatedPlaybackParticipantClass() AVCoordinatedPlaybackParticipantClass {
	return getAVCoordinatedPlaybackParticipantClass()
}

type AVCoordinatedPlaybackParticipantClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCoordinatedPlaybackParticipantClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCoordinatedPlaybackParticipantClass) Alloc() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a participant in a coordinated playback session.
//
// # Overview
//
// Access the other participants in a session through the playback
// coordinator’s [AVCoordinatedPlaybackParticipant.OtherParticipants] property to determine their playback
// readiness and suspension reasons.
//
// # Accessing participant status
//
//   - [AVCoordinatedPlaybackParticipant.Identifier]: A unique identifier for the participant.
//   - [AVCoordinatedPlaybackParticipant.ReadyToPlay]: A Boolean value that indicates whether the participant is ready to start coordinated playback.
//   - [AVCoordinatedPlaybackParticipant.SuspensionReasons]: The reasons a participant isn’t currently participating in coordinated playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant
type AVCoordinatedPlaybackParticipant struct {
	objectivec.Object
}

// AVCoordinatedPlaybackParticipantFromID constructs a [AVCoordinatedPlaybackParticipant] from an objc.ID.
//
// An object that represents a participant in a coordinated playback session.
func AVCoordinatedPlaybackParticipantFromID(id objc.ID) AVCoordinatedPlaybackParticipant {
	return AVCoordinatedPlaybackParticipant{objectivec.Object{ID: id}}
}

// NOTE: AVCoordinatedPlaybackParticipant adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCoordinatedPlaybackParticipant] class.
//
// # Accessing participant status
//
//   - [IAVCoordinatedPlaybackParticipant.Identifier]: A unique identifier for the participant.
//   - [IAVCoordinatedPlaybackParticipant.ReadyToPlay]: A Boolean value that indicates whether the participant is ready to start coordinated playback.
//   - [IAVCoordinatedPlaybackParticipant.SuspensionReasons]: The reasons a participant isn’t currently participating in coordinated playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant
type IAVCoordinatedPlaybackParticipant interface {
	objectivec.IObject

	// Topic: Accessing participant status

	// A unique identifier for the participant.
	Identifier() foundation.NSUUID
	// A Boolean value that indicates whether the participant is ready to start coordinated playback.
	ReadyToPlay() bool
	// The reasons a participant isn’t currently participating in coordinated playback.
	SuspensionReasons() []string

	// The identifiers of the other participants in a group.
	OtherParticipants() IAVCoordinatedPlaybackParticipant
	SetOtherParticipants(value IAVCoordinatedPlaybackParticipant)
}

// Init initializes the instance.
func (c AVCoordinatedPlaybackParticipant) Init() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCoordinatedPlaybackParticipant) Autorelease() AVCoordinatedPlaybackParticipant {
	rv := objc.Send[AVCoordinatedPlaybackParticipant](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCoordinatedPlaybackParticipant creates a new AVCoordinatedPlaybackParticipant instance.
func NewAVCoordinatedPlaybackParticipant() AVCoordinatedPlaybackParticipant {
	class := getAVCoordinatedPlaybackParticipantClass()
	rv := objc.Send[AVCoordinatedPlaybackParticipant](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A unique identifier for the participant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant/identifier
func (c AVCoordinatedPlaybackParticipant) Identifier() foundation.NSUUID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the participant is ready to start
// coordinated playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant/isReadyToPlay
func (c AVCoordinatedPlaybackParticipant) ReadyToPlay() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isReadyToPlay"))
	return rv
}

// The reasons a participant isn’t currently participating in coordinated
// playback.
//
// # Discussion
//
// This value is empty if the participant’s playback isn’t in a suspended
// state.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackParticipant/suspensionReasons
func (c AVCoordinatedPlaybackParticipant) SuspensionReasons() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("suspensionReasons"))
	return objc.ConvertSliceToStrings(rv)
}

// The identifiers of the other participants in a group.
//
// See: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/otherparticipants
func (c AVCoordinatedPlaybackParticipant) OtherParticipants() IAVCoordinatedPlaybackParticipant {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("otherParticipants"))
	return AVCoordinatedPlaybackParticipantFromID(objc.ID(rv))
}
func (c AVCoordinatedPlaybackParticipant) SetOtherParticipants(value IAVCoordinatedPlaybackParticipant) {
	objc.Send[struct{}](c.ID, objc.Sel("setOtherParticipants:"), value)
}

// A notification that the coordinator posts when its other participants
// change.
//
// See: https://developer.apple.com/documentation/avfoundation/avplaybackcoordinator/otherparticipantsdidchangenotification
func (_AVCoordinatedPlaybackParticipantClass AVCoordinatedPlaybackParticipantClass) OtherParticipantsDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_AVCoordinatedPlaybackParticipantClass.class), objc.Sel("AVPlaybackCoordinatorOtherParticipantsDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}
