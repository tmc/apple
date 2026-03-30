// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlaybackCoordinator] class.
var (
	_AVPlaybackCoordinatorClass     AVPlaybackCoordinatorClass
	_AVPlaybackCoordinatorClassOnce sync.Once
)

func getAVPlaybackCoordinatorClass() AVPlaybackCoordinatorClass {
	_AVPlaybackCoordinatorClassOnce.Do(func() {
		_AVPlaybackCoordinatorClass = AVPlaybackCoordinatorClass{class: objc.GetClass("AVPlaybackCoordinator")}
	})
	return _AVPlaybackCoordinatorClass
}

// GetAVPlaybackCoordinatorClass returns the class object for AVPlaybackCoordinator.
func GetAVPlaybackCoordinatorClass() AVPlaybackCoordinatorClass {
	return getAVPlaybackCoordinatorClass()
}

type AVPlaybackCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlaybackCoordinatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlaybackCoordinatorClass) Alloc() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that coordinates the playback of players in a connected group.
//
// # Overview
//
// The framework provides two playback coordinator subclasses that manage
// different types of player objects:
//
// - [AVPlayerPlaybackCoordinator] coordinates the state of [AVPlayer]
// objects. If your app uses [AVPlayer], continue to use its standard
// interfaces to control playback. The coordinator intercepts changes to the
// player’s rate and time, and propagates them to other players in the
// group. - [AVDelegatingPlaybackCoordinator] coordinates the state of custom
// player objects. If your app uses a custom player, such as one that renders
// media using [AVSampleBufferDisplayLayer] and [AVSampleBufferAudioRenderer],
// use this object to coordinate group playback. Adopt the coordinator’s
// delegate protocol so that your player responds to the commands that the
// coordinator issues.
//
// # Configuring playback policies
//
//   - [AVPlaybackCoordinator.ParticipantLimitForWaitingOutSuspensionsWithReason]: Returns the limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
//   - [AVPlaybackCoordinator.SetParticipantLimitForWaitingOutSuspensionsWithReason]: Sets a limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
//   - [AVPlaybackCoordinator.SuspensionReasonsThatTriggerWaiting]: The reasons that cause a coordinator to suspend playback.
//   - [AVPlaybackCoordinator.SetSuspensionReasonsThatTriggerWaiting]
//   - [AVPlaybackCoordinator.PauseSnapsToMediaTimeOfOriginator]: A Boolean value that indicates whether participants mirror the originator’s stop time when they pause.
//   - [AVPlaybackCoordinator.SetPauseSnapsToMediaTimeOfOriginator]
//
// # Suspending state coordination
//
//   - [AVPlaybackCoordinator.BeginSuspensionForReason]: Tells the coordinator to stop sending playback commands temporarily when the playback object disconnects from the group activity.
//   - [AVPlaybackCoordinator.ExpectedItemTimeAtHostTime]: Returns a time in the current item’s timeline that the coordinator expects to play at the specified host time.
//
// # Observing suspension reasons
//
//   - [AVPlaybackCoordinator.SuspensionReasons]: The reasons a coordinator is currently unable to participate in a group playback activity.
//
// # Observing other participants
//
//   - [AVPlaybackCoordinator.OtherParticipants]: The identifiers of the other participants in a group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator
type AVPlaybackCoordinator struct {
	objectivec.Object
}

// AVPlaybackCoordinatorFromID constructs a [AVPlaybackCoordinator] from an objc.ID.
//
// An object that coordinates the playback of players in a connected group.
func AVPlaybackCoordinatorFromID(id objc.ID) AVPlaybackCoordinator {
	return AVPlaybackCoordinator{objectivec.Object{ID: id}}
}

// NOTE: AVPlaybackCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlaybackCoordinator] class.
//
// # Configuring playback policies
//
//   - [IAVPlaybackCoordinator.ParticipantLimitForWaitingOutSuspensionsWithReason]: Returns the limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
//   - [IAVPlaybackCoordinator.SetParticipantLimitForWaitingOutSuspensionsWithReason]: Sets a limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
//   - [IAVPlaybackCoordinator.SuspensionReasonsThatTriggerWaiting]: The reasons that cause a coordinator to suspend playback.
//   - [IAVPlaybackCoordinator.SetSuspensionReasonsThatTriggerWaiting]
//   - [IAVPlaybackCoordinator.PauseSnapsToMediaTimeOfOriginator]: A Boolean value that indicates whether participants mirror the originator’s stop time when they pause.
//   - [IAVPlaybackCoordinator.SetPauseSnapsToMediaTimeOfOriginator]
//
// # Suspending state coordination
//
//   - [IAVPlaybackCoordinator.BeginSuspensionForReason]: Tells the coordinator to stop sending playback commands temporarily when the playback object disconnects from the group activity.
//   - [IAVPlaybackCoordinator.ExpectedItemTimeAtHostTime]: Returns a time in the current item’s timeline that the coordinator expects to play at the specified host time.
//
// # Observing suspension reasons
//
//   - [IAVPlaybackCoordinator.SuspensionReasons]: The reasons a coordinator is currently unable to participate in a group playback activity.
//
// # Observing other participants
//
//   - [IAVPlaybackCoordinator.OtherParticipants]: The identifiers of the other participants in a group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator
type IAVPlaybackCoordinator interface {
	objectivec.IObject

	// Topic: Configuring playback policies

	// Returns the limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
	ParticipantLimitForWaitingOutSuspensionsWithReason(reason AVCoordinatedPlaybackSuspensionReason) int
	// Sets a limit on the number of partipants that a group may contain before the coordinator stops waiting on suspensions that occur for a particular reason.
	SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason AVCoordinatedPlaybackSuspensionReason)
	// The reasons that cause a coordinator to suspend playback.
	SuspensionReasonsThatTriggerWaiting() []string
	SetSuspensionReasonsThatTriggerWaiting(value []string)
	// A Boolean value that indicates whether participants mirror the originator’s stop time when they pause.
	PauseSnapsToMediaTimeOfOriginator() bool
	SetPauseSnapsToMediaTimeOfOriginator(value bool)

	// Topic: Suspending state coordination

	// Tells the coordinator to stop sending playback commands temporarily when the playback object disconnects from the group activity.
	BeginSuspensionForReason(suspensionReason AVCoordinatedPlaybackSuspensionReason) IAVCoordinatedPlaybackSuspension
	// Returns a time in the current item’s timeline that the coordinator expects to play at the specified host time.
	ExpectedItemTimeAtHostTime(hostClockTime coremedia.CMTime) coremedia.CMTime

	// Topic: Observing suspension reasons

	// The reasons a coordinator is currently unable to participate in a group playback activity.
	SuspensionReasons() []string

	// Topic: Observing other participants

	// The identifiers of the other participants in a group.
	OtherParticipants() []AVCoordinatedPlaybackParticipant
}

// Init initializes the instance.
func (p AVPlaybackCoordinator) Init() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlaybackCoordinator) Autorelease() AVPlaybackCoordinator {
	rv := objc.Send[AVPlaybackCoordinator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlaybackCoordinator creates a new AVPlaybackCoordinator instance.
func NewAVPlaybackCoordinator() AVPlaybackCoordinator {
	class := getAVPlaybackCoordinatorClass()
	rv := objc.Send[AVPlaybackCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the limit on the number of partipants that a group may contain
// before the coordinator stops waiting on suspensions that occur for a
// particular reason.
//
// reason: The suspension reason to find a participant limit for.
//
// # Return Value
//
// The participant limit.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/participantLimitForWaitingOutSuspensions(withReason:)
func (p AVPlaybackCoordinator) ParticipantLimitForWaitingOutSuspensionsWithReason(reason AVCoordinatedPlaybackSuspensionReason) int {
	rv := objc.Send[int](p.ID, objc.Sel("participantLimitForWaitingOutSuspensionsWithReason:"), objc.String(string(reason)))
	return rv
}

// Sets a limit on the number of partipants that a group may contain before
// the coordinator stops waiting on suspensions that occur for a particular
// reason.
//
// participantLimit: The number of participants.
//
// reason: The suspension reason to set a limit for.
//
// # Discussion
//
// This method provides additional configuration of the values your app sets
// for the [SuspensionReasonsThatTriggerWaiting] property. When a coordinator
// decides whether one participant’s suspensions cause others to wait, it
// also considers any participant limits that you set on the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/setParticipantLimit(_:forWaitingOutSuspensionsWithReason:)
func (p AVPlaybackCoordinator) SetParticipantLimitForWaitingOutSuspensionsWithReason(participantLimit int, reason AVCoordinatedPlaybackSuspensionReason) {
	objc.Send[objc.ID](p.ID, objc.Sel("setParticipantLimit:forWaitingOutSuspensionsWithReason:"), participantLimit, objc.String(string(reason)))
}

// Tells the coordinator to stop sending playback commands temporarily when
// the playback object disconnects from the group activity.
//
// suspensionReason: The reason for the suspension. Indicate a system-defined value or a custom
// suspension reason.
//
// # Return Value
//
// A suspension object.
//
// # Discussion
//
// End a suspension by calling its [End] or [EndProposingNewTime] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/beginSuspension(for:)
func (p AVPlaybackCoordinator) BeginSuspensionForReason(suspensionReason AVCoordinatedPlaybackSuspensionReason) IAVCoordinatedPlaybackSuspension {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("beginSuspensionForReason:"), objc.String(string(suspensionReason)))
	return AVCoordinatedPlaybackSuspensionFromID(rv)
}

// Returns a time in the current item’s timeline that the coordinator
// expects to play at the specified host time.
//
// hostClockTime: The host time to return a player item time for.
//
// # Return Value
//
// A time in the current item’s timeline.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/expectedItemTime(atHostTime:)
func (p AVPlaybackCoordinator) ExpectedItemTimeAtHostTime(hostClockTime coremedia.CMTime) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("expectedItemTimeAtHostTime:"), hostClockTime)
	return coremedia.CMTime(rv)
}

// The reasons that cause a coordinator to suspend playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/suspensionReasonsThatTriggerWaiting
func (p AVPlaybackCoordinator) SuspensionReasonsThatTriggerWaiting() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("suspensionReasonsThatTriggerWaiting"))
	return objc.ConvertSliceToStrings(rv)
}
func (p AVPlaybackCoordinator) SetSuspensionReasonsThatTriggerWaiting(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setSuspensionReasonsThatTriggerWaiting:"), objectivec.StringSliceToNSArray(value))
}

// A Boolean value that indicates whether participants mirror the
// originator’s stop time when they pause.
//
// # Discussion
//
// If this value is true, all participants seek to the originator’s stop
// time after they pause. Use this setting if it counteracts network delays
// that result from communicating the originator’s pause state to the other
// participants.
//
// If this value is false, it’s acceptable for participants to stop at
// slightly different times, and a pause doesn’t cause the time of other
// participants to jump back.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/pauseSnapsToMediaTimeOfOriginator
func (p AVPlaybackCoordinator) PauseSnapsToMediaTimeOfOriginator() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("pauseSnapsToMediaTimeOfOriginator"))
	return rv
}
func (p AVPlaybackCoordinator) SetPauseSnapsToMediaTimeOfOriginator(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPauseSnapsToMediaTimeOfOriginator:"), value)
}

// The reasons a coordinator is currently unable to participate in a group
// playback activity.
//
// # Discussion
//
// The coordinator doesn’t respond to changes in group playback state when
// this property value contains suspension reasons.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/suspensionReasons
func (p AVPlaybackCoordinator) SuspensionReasons() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("suspensionReasons"))
	return objc.ConvertSliceToStrings(rv)
}

// The identifiers of the other participants in a group.
//
// # Discussion
//
// Use this property value to create a user interface that informs the user
// about the state of other participants in the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/otherParticipants
func (p AVPlaybackCoordinator) OtherParticipants() []AVCoordinatedPlaybackParticipant {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("otherParticipants"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCoordinatedPlaybackParticipant {
		return AVCoordinatedPlaybackParticipantFromID(id)
	})
}
