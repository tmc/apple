// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCoordinatedPlaybackSuspension] class.
var (
	_AVCoordinatedPlaybackSuspensionClass     AVCoordinatedPlaybackSuspensionClass
	_AVCoordinatedPlaybackSuspensionClassOnce sync.Once
)

func getAVCoordinatedPlaybackSuspensionClass() AVCoordinatedPlaybackSuspensionClass {
	_AVCoordinatedPlaybackSuspensionClassOnce.Do(func() {
		_AVCoordinatedPlaybackSuspensionClass = AVCoordinatedPlaybackSuspensionClass{class: objc.GetClass("AVCoordinatedPlaybackSuspension")}
	})
	return _AVCoordinatedPlaybackSuspensionClass
}

// GetAVCoordinatedPlaybackSuspensionClass returns the class object for AVCoordinatedPlaybackSuspension.
func GetAVCoordinatedPlaybackSuspensionClass() AVCoordinatedPlaybackSuspensionClass {
	return getAVCoordinatedPlaybackSuspensionClass()
}

type AVCoordinatedPlaybackSuspensionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCoordinatedPlaybackSuspensionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCoordinatedPlaybackSuspensionClass) Alloc() AVCoordinatedPlaybackSuspension {
	rv := objc.Send[AVCoordinatedPlaybackSuspension](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a temporary suspension of coordinated playback.
//
// # Overview
// 
// See the playback coordinator’s [BeginSuspensionForReason] method for
// details about suspending playback.
//
// # Inspecting a suspension
//
//   - [AVCoordinatedPlaybackSuspension.BeginDate]: The time the suspension begins.
//   - [AVCoordinatedPlaybackSuspension.Reason]: The reason for the suspension.
//
// # Ending a suspension
//
//   - [AVCoordinatedPlaybackSuspension.End]: Ends a suspension.
//   - [AVCoordinatedPlaybackSuspension.EndProposingNewTime]: Ends a suspension and proposes a new playback time to the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension
type AVCoordinatedPlaybackSuspension struct {
	objectivec.Object
}

// AVCoordinatedPlaybackSuspensionFromID constructs a [AVCoordinatedPlaybackSuspension] from an objc.ID.
//
// An object that represents a temporary suspension of coordinated playback.
func AVCoordinatedPlaybackSuspensionFromID(id objc.ID) AVCoordinatedPlaybackSuspension {
	return AVCoordinatedPlaybackSuspension{objectivec.Object{ID: id}}
}
// NOTE: AVCoordinatedPlaybackSuspension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCoordinatedPlaybackSuspension] class.
//
// # Inspecting a suspension
//
//   - [IAVCoordinatedPlaybackSuspension.BeginDate]: The time the suspension begins.
//   - [IAVCoordinatedPlaybackSuspension.Reason]: The reason for the suspension.
//
// # Ending a suspension
//
//   - [IAVCoordinatedPlaybackSuspension.End]: Ends a suspension.
//   - [IAVCoordinatedPlaybackSuspension.EndProposingNewTime]: Ends a suspension and proposes a new playback time to the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension
type IAVCoordinatedPlaybackSuspension interface {
	objectivec.IObject

	// Topic: Inspecting a suspension

	// The time the suspension begins.
	BeginDate() foundation.INSDate
	// The reason for the suspension.
	Reason() AVCoordinatedPlaybackSuspensionReason

	// Topic: Ending a suspension

	// Ends a suspension.
	End()
	// Ends a suspension and proposes a new playback time to the group.
	EndProposingNewTime(time coremedia.CMTime)
}

// Init initializes the instance.
func (c AVCoordinatedPlaybackSuspension) Init() AVCoordinatedPlaybackSuspension {
	rv := objc.Send[AVCoordinatedPlaybackSuspension](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCoordinatedPlaybackSuspension) Autorelease() AVCoordinatedPlaybackSuspension {
	rv := objc.Send[AVCoordinatedPlaybackSuspension](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCoordinatedPlaybackSuspension creates a new AVCoordinatedPlaybackSuspension instance.
func NewAVCoordinatedPlaybackSuspension() AVCoordinatedPlaybackSuspension {
	class := getAVCoordinatedPlaybackSuspensionClass()
	rv := objc.Send[AVCoordinatedPlaybackSuspension](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Ends a suspension.
//
// # Discussion
// 
// If this is the last suspension, the coordinator adjusts the timing of its
// playback object to match the group.
// 
// To end a suspension and simultaneously propose a new playback time to the
// group, call the [EndProposingNewTime] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/end()
func (c AVCoordinatedPlaybackSuspension) End() {
	objc.Send[objc.ID](c.ID, objc.Sel("end"))
}
// Ends a suspension and proposes a new playback time to the group.
//
// time: The proposed playback time. Passing a nonnumeric time results in the same
// behavior as calling the [End] method.
//
// # Discussion
// 
// If this is the last suspension, the coordinator proposes a new time to the
// group without changing the group’s playback rate. If it isn’t, the
// coordinator only proposes the new time after all other suspensions end.
// 
// A suspension that ends after this one ends can override the proposed time.
// Similarly, playback commands from the group that arrive after this
// suspension ends, override a pending proposal.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/end(proposingNewTime:)
func (c AVCoordinatedPlaybackSuspension) EndProposingNewTime(time coremedia.CMTime) {
	objc.Send[objc.ID](c.ID, objc.Sel("endProposingNewTime:"), time)
}

// The time the suspension begins.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/beginDate
func (c AVCoordinatedPlaybackSuspension) BeginDate() foundation.INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("beginDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// The reason for the suspension.
//
// # Discussion
// 
// The coordinator communicates the suspension reason to other participants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/reason-swift.property
func (c AVCoordinatedPlaybackSuspension) Reason() AVCoordinatedPlaybackSuspensionReason {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("reason"))
	return AVCoordinatedPlaybackSuspensionReason(foundation.NSStringFromID(rv).String())
}

