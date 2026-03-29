// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureReactionEffectState] class.
var (
	_AVCaptureReactionEffectStateClass     AVCaptureReactionEffectStateClass
	_AVCaptureReactionEffectStateClassOnce sync.Once
)

func getAVCaptureReactionEffectStateClass() AVCaptureReactionEffectStateClass {
	_AVCaptureReactionEffectStateClassOnce.Do(func() {
		_AVCaptureReactionEffectStateClass = AVCaptureReactionEffectStateClass{class: objc.GetClass("AVCaptureReactionEffectState")}
	})
	return _AVCaptureReactionEffectStateClass
}

// GetAVCaptureReactionEffectStateClass returns the class object for AVCaptureReactionEffectState.
func GetAVCaptureReactionEffectStateClass() AVCaptureReactionEffectStateClass {
	return getAVCaptureReactionEffectStateClass()
}

type AVCaptureReactionEffectStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureReactionEffectStateClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureReactionEffectStateClass) Alloc() AVCaptureReactionEffectState {
	rv := objc.Send[AVCaptureReactionEffectState](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that reports the state of a reaction effect performed on a
// capture device.
//
// # Overview
// 
// Obtain an instance of this class by querying a capture device’s
// [AVCaptureReactionEffectState.ReactionEffectsInProgress] property. The system adds new entries to this
// array when you call [PerformEffectForReaction] or by gesture detection in
// the capture stream when the value of [AVCaptureReactionEffectState.ReactionEffectGesturesEnabled] is
// [true].
// 
// The system renders the effect before providing frames to your app, and
// these status objects let you know when it performs the effect.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring the effect state
//
//   - [AVCaptureReactionEffectState.ReactionType]: The type of reaction.
//   - [AVCaptureReactionEffectState.StartTime]: The presentation time of the first frame where the system renders the effect.
//   - [AVCaptureReactionEffectState.EndTime]: The presentation time of the first frame following the end of a reaction effect.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState
type AVCaptureReactionEffectState struct {
	objectivec.Object
}

// AVCaptureReactionEffectStateFromID constructs a [AVCaptureReactionEffectState] from an objc.ID.
//
// An object that reports the state of a reaction effect performed on a
// capture device.
func AVCaptureReactionEffectStateFromID(id objc.ID) AVCaptureReactionEffectState {
	return AVCaptureReactionEffectState{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureReactionEffectState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureReactionEffectState] class.
//
// # Configuring the effect state
//
//   - [IAVCaptureReactionEffectState.ReactionType]: The type of reaction.
//   - [IAVCaptureReactionEffectState.StartTime]: The presentation time of the first frame where the system renders the effect.
//   - [IAVCaptureReactionEffectState.EndTime]: The presentation time of the first frame following the end of a reaction effect.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState
type IAVCaptureReactionEffectState interface {
	objectivec.IObject

	// Topic: Configuring the effect state

	// The type of reaction.
	ReactionType() AVCaptureReactionType
	// The presentation time of the first frame where the system renders the effect.
	StartTime() uintptr
	// The presentation time of the first frame following the end of a reaction effect.
	EndTime() uintptr

	// A set of reactions types that a device supports performing.
	AvailableReactionTypes() AVCaptureReactionType
	SetAvailableReactionTypes(value AVCaptureReactionType)
	// A Boolean value that indicates whether you can perform reaction effects on a capture device.
	CanPerformReactionEffects() bool
	SetCanPerformReactionEffects(value bool)
	// An array of reaction effects that the device is currently performing, sorted by timestamp.
	ReactionEffectsInProgress() IAVCaptureReactionEffectState
	SetReactionEffectsInProgress(value IAVCaptureReactionEffectState)
}

// Init initializes the instance.
func (c AVCaptureReactionEffectState) Init() AVCaptureReactionEffectState {
	rv := objc.Send[AVCaptureReactionEffectState](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureReactionEffectState) Autorelease() AVCaptureReactionEffectState {
	rv := objc.Send[AVCaptureReactionEffectState](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureReactionEffectState creates a new AVCaptureReactionEffectState instance.
func NewAVCaptureReactionEffectState() AVCaptureReactionEffectState {
	class := getAVCaptureReactionEffectStateClass()
	rv := objc.Send[AVCaptureReactionEffectState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type of reaction.
//
// # Discussion
// 
// There may be multiple reactions of the same type at a given time. Some may
// come from calls to [PerformEffectForReaction] and others from gesture
// detection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState/reactionType
func (c AVCaptureReactionEffectState) ReactionType() AVCaptureReactionType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("reactionType"))
	return AVCaptureReactionType(foundation.NSStringFromID(rv).String())
}
// The presentation time of the first frame where the system renders the
// effect.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState/startTime
func (c AVCaptureReactionEffectState) StartTime() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("startTime"))
	return rv
}
// The presentation time of the first frame following the end of a reaction
// effect.
//
// # Discussion
// 
// The value is [invalid] while the effect is in progress, but changes to a
// valid time when the reaction effect completes and the system removes it
// from the list of [ReactionEffectsInProgress].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionEffectState/endTime
func (c AVCaptureReactionEffectState) EndTime() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("endTime"))
	return rv
}
// A set of reactions types that a device supports performing.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/availablereactiontypes
func (c AVCaptureReactionEffectState) AvailableReactionTypes() AVCaptureReactionType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableReactionTypes"))
	return AVCaptureReactionType(foundation.NSStringFromID(rv).String())
}
func (c AVCaptureReactionEffectState) SetAvailableReactionTypes(value AVCaptureReactionType) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableReactionTypes:"), objc.String(string(value)))
}
// A Boolean value that indicates whether you can perform reaction effects on
// a capture device.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/canperformreactioneffects
func (c AVCaptureReactionEffectState) CanPerformReactionEffects() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canPerformReactionEffects"))
	return rv
}
func (c AVCaptureReactionEffectState) SetCanPerformReactionEffects(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCanPerformReactionEffects:"), value)
}
// An array of reaction effects that the device is currently performing,
// sorted by timestamp.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/reactioneffectsinprogress
func (c AVCaptureReactionEffectState) ReactionEffectsInProgress() IAVCaptureReactionEffectState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("reactionEffectsInProgress"))
	return AVCaptureReactionEffectStateFromID(objc.ID(rv))
}
func (c AVCaptureReactionEffectState) SetReactionEffectsInProgress(value IAVCaptureReactionEffectState) {
	objc.Send[struct{}](c.ID, objc.Sel("setReactionEffectsInProgress:"), value)
}

// A Boolean value that indicates whether gesture detection triggers reaction
// effects on the video stream.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/reactioneffectgesturesenabled
func (_AVCaptureReactionEffectStateClass AVCaptureReactionEffectStateClass) ReactionEffectGesturesEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureReactionEffectStateClass.class), objc.Sel("reactionEffectGesturesEnabled"))
	return rv
}
func (_AVCaptureReactionEffectStateClass AVCaptureReactionEffectStateClass) SetReactionEffectGesturesEnabled(value bool) {
	objc.Send[struct{}](objc.ID(_AVCaptureReactionEffectStateClass.class), objc.Sel("setReactionEffectGesturesEnabled:"), value)
}
// A Boolean value that indicates whether the app supports performing reaction
// effects.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/reactioneffectsenabled
func (_AVCaptureReactionEffectStateClass AVCaptureReactionEffectStateClass) ReactionEffectsEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureReactionEffectStateClass.class), objc.Sel("reactionEffectsEnabled"))
	return rv
}
func (_AVCaptureReactionEffectStateClass AVCaptureReactionEffectStateClass) SetReactionEffectsEnabled(value bool) {
	objc.Send[struct{}](objc.ID(_AVCaptureReactionEffectStateClass.class), objc.Sel("setReactionEffectsEnabled:"), value)
}

