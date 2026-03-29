// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVDelegatingPlaybackCoordinator] class.
var (
	_AVDelegatingPlaybackCoordinatorClass     AVDelegatingPlaybackCoordinatorClass
	_AVDelegatingPlaybackCoordinatorClassOnce sync.Once
)

func getAVDelegatingPlaybackCoordinatorClass() AVDelegatingPlaybackCoordinatorClass {
	_AVDelegatingPlaybackCoordinatorClassOnce.Do(func() {
		_AVDelegatingPlaybackCoordinatorClass = AVDelegatingPlaybackCoordinatorClass{class: objc.GetClass("AVDelegatingPlaybackCoordinator")}
	})
	return _AVDelegatingPlaybackCoordinatorClass
}

// GetAVDelegatingPlaybackCoordinatorClass returns the class object for AVDelegatingPlaybackCoordinator.
func GetAVDelegatingPlaybackCoordinatorClass() AVDelegatingPlaybackCoordinatorClass {
	return getAVDelegatingPlaybackCoordinatorClass()
}

type AVDelegatingPlaybackCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVDelegatingPlaybackCoordinatorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVDelegatingPlaybackCoordinatorClass) Alloc() AVDelegatingPlaybackCoordinator {
	rv := objc.Send[AVDelegatingPlaybackCoordinator](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A playback coordinator subclass that coordinates the playback of custom
// player objects in a connected group.
//
// # Overview
// 
// This object coordinates the state of custom player objects, such as those
// that render media using [AVSampleBufferDisplayLayer] and
// [AVSampleBufferAudioRenderer], or that play audio using [AVAudioEngine].
// 
// Adopt the [AVPlaybackCoordinatorPlaybackControlDelegate] protocol so that
// your app responds to playback commands from the coordinator. The commands
// provide the details of a requested state change so you can control your
// player object accordingly.
// 
// [media-3783472]
//
// [AVAudioEngine]: https://developer.apple.com/documentation/AVFAudio/AVAudioEngine
//
// # Creating a coordinator
//
//   - [AVDelegatingPlaybackCoordinator.InitWithPlaybackControlDelegate]: Creates a playback coordinator for a custom playback object.
//
// # Identifying items
//
//   - [AVDelegatingPlaybackCoordinator.CurrentItemIdentifier]: An identifier of the current item.
//
// # Accessing the delegate
//
//   - [AVDelegatingPlaybackCoordinator.PlaybackControlDelegate]: The delegate object for the playback coordinator.
//
// # Coordinating state changes
//
//   - [AVDelegatingPlaybackCoordinator.CoordinateRateChangeToRateOptions]: Coordinates a rate change across all participants, waiting for others to become ready, if necessary.
//   - [AVDelegatingPlaybackCoordinator.CoordinateSeekToTimeOptions]: Coordinates a seek to the specified time for all connected participants.
//   - [AVDelegatingPlaybackCoordinator.TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase]: Tells the coordinator to transition to a new item.
//   - [AVDelegatingPlaybackCoordinator.ReapplyCurrentItemStateToPlaybackControlDelegate]: Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator
type AVDelegatingPlaybackCoordinator struct {
	AVPlaybackCoordinator
}

// AVDelegatingPlaybackCoordinatorFromID constructs a [AVDelegatingPlaybackCoordinator] from an objc.ID.
//
// A playback coordinator subclass that coordinates the playback of custom
// player objects in a connected group.
func AVDelegatingPlaybackCoordinatorFromID(id objc.ID) AVDelegatingPlaybackCoordinator {
	return AVDelegatingPlaybackCoordinator{AVPlaybackCoordinator: AVPlaybackCoordinatorFromID(id)}
}
// NOTE: AVDelegatingPlaybackCoordinator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVDelegatingPlaybackCoordinator] class.
//
// # Creating a coordinator
//
//   - [IAVDelegatingPlaybackCoordinator.InitWithPlaybackControlDelegate]: Creates a playback coordinator for a custom playback object.
//
// # Identifying items
//
//   - [IAVDelegatingPlaybackCoordinator.CurrentItemIdentifier]: An identifier of the current item.
//
// # Accessing the delegate
//
//   - [IAVDelegatingPlaybackCoordinator.PlaybackControlDelegate]: The delegate object for the playback coordinator.
//
// # Coordinating state changes
//
//   - [IAVDelegatingPlaybackCoordinator.CoordinateRateChangeToRateOptions]: Coordinates a rate change across all participants, waiting for others to become ready, if necessary.
//   - [IAVDelegatingPlaybackCoordinator.CoordinateSeekToTimeOptions]: Coordinates a seek to the specified time for all connected participants.
//   - [IAVDelegatingPlaybackCoordinator.TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase]: Tells the coordinator to transition to a new item.
//   - [IAVDelegatingPlaybackCoordinator.ReapplyCurrentItemStateToPlaybackControlDelegate]: Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator
type IAVDelegatingPlaybackCoordinator interface {
	IAVPlaybackCoordinator

	// Topic: Creating a coordinator

	// Creates a playback coordinator for a custom playback object.
	InitWithPlaybackControlDelegate(playbackControlDelegate AVPlaybackCoordinatorPlaybackControlDelegate) AVDelegatingPlaybackCoordinator

	// Topic: Identifying items

	// An identifier of the current item.
	CurrentItemIdentifier() string

	// Topic: Accessing the delegate

	// The delegate object for the playback coordinator.
	PlaybackControlDelegate() AVPlaybackCoordinatorPlaybackControlDelegate

	// Topic: Coordinating state changes

	// Coordinates a rate change across all participants, waiting for others to become ready, if necessary.
	CoordinateRateChangeToRateOptions(rate float32, options AVDelegatingPlaybackCoordinatorRateChangeOptions)
	// Coordinates a seek to the specified time for all connected participants.
	CoordinateSeekToTimeOptions(time coremedia.CMTime, options AVDelegatingPlaybackCoordinatorSeekOptions)
	// Tells the coordinator to transition to a new item.
	TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier string, snapshotTimebase uintptr)
	// Tells the coordinator to reissue current play state commands to synchronize the current item to the state of other participants.
	ReapplyCurrentItemStateToPlaybackControlDelegate()
}

// Init initializes the instance.
func (d AVDelegatingPlaybackCoordinator) Init() AVDelegatingPlaybackCoordinator {
	rv := objc.Send[AVDelegatingPlaybackCoordinator](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d AVDelegatingPlaybackCoordinator) Autorelease() AVDelegatingPlaybackCoordinator {
	rv := objc.Send[AVDelegatingPlaybackCoordinator](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVDelegatingPlaybackCoordinator creates a new AVDelegatingPlaybackCoordinator instance.
func NewAVDelegatingPlaybackCoordinator() AVDelegatingPlaybackCoordinator {
	class := getAVDelegatingPlaybackCoordinatorClass()
	rv := objc.Send[AVDelegatingPlaybackCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a playback coordinator for a custom playback object.
//
// playbackControlDelegate: The playback control delegate for the playback coordinator.
//
// # Discussion
// 
// If your app doesn’t use [AVPlayer] for playback, create an instance of
// this class to coordinate playback of your customer player.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/init(playbackControlDelegate:)
func NewDelegatingPlaybackCoordinatorWithPlaybackControlDelegate(playbackControlDelegate AVPlaybackCoordinatorPlaybackControlDelegate) AVDelegatingPlaybackCoordinator {
	instance := getAVDelegatingPlaybackCoordinatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPlaybackControlDelegate:"), playbackControlDelegate)
	return AVDelegatingPlaybackCoordinatorFromID(rv)
}

// Creates a playback coordinator for a custom playback object.
//
// playbackControlDelegate: The playback control delegate for the playback coordinator.
//
// # Discussion
// 
// If your app doesn’t use [AVPlayer] for playback, create an instance of
// this class to coordinate playback of your customer player.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/init(playbackControlDelegate:)
func (d AVDelegatingPlaybackCoordinator) InitWithPlaybackControlDelegate(playbackControlDelegate AVPlaybackCoordinatorPlaybackControlDelegate) AVDelegatingPlaybackCoordinator {
	rv := objc.Send[AVDelegatingPlaybackCoordinator](d.ID, objc.Sel("initWithPlaybackControlDelegate:"), playbackControlDelegate)
	return rv
}
// Coordinates a rate change across all participants, waiting for others to
// become ready, if necessary.
//
// rate: The playback rate for the group to use.
//
// options: Additional configuration of the rate change.
//
// # Discussion
// 
// When the rate changes from zero to nonzero, the coordinator may also wait
// for participant suspensions from the [SuspensionReasonsThatTriggerWaiting]
// property.
// 
// Don’t call this method if the rate change doesn’t affect the group, or
// if the group doesn’t have control over local playback temporarily. For
// example, don’t call the method for a pause that occurs due to an audio
// session interruption. In those cases, inform the coordinator by beginning a
// suspension with an appropriate reason.
// 
// The suspension stops the coordinator from issuing further commands to its
// delegate. After beginning a suspension, you can reconfigure your app’s
// playback object as necessary.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/coordinateRateChange(to:options:)
func (d AVDelegatingPlaybackCoordinator) CoordinateRateChangeToRateOptions(rate float32, options AVDelegatingPlaybackCoordinatorRateChangeOptions) {
	objc.Send[objc.ID](d.ID, objc.Sel("coordinateRateChangeToRate:options:"), rate, options)
}
// Coordinates a seek to the specified time for all connected participants.
//
// time: A time the group seeks to when the command ends.
//
// options: Additional configuration of the seek.
//
// # Discussion
// 
// To end a suspension and also affect the group timing, see
// [EndProposingNewTime].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/coordinateSeek(to:options:)
func (d AVDelegatingPlaybackCoordinator) CoordinateSeekToTimeOptions(time coremedia.CMTime, options AVDelegatingPlaybackCoordinatorSeekOptions) {
	objc.Send[objc.ID](d.ID, objc.Sel("coordinateSeekToTime:options:"), time, options)
}
// Tells the coordinator to transition to a new item.
//
// itemIdentifier: The identifier for the new current item, which is `nil` if there isn’t
// anything to play.
//
// snapshotTimebase: A time base that communicates the initial playback state of the new item.
// If you specify `nil`, the coordinator assumes that the player pauses at
// [zero].
// 
// You can retrieve an appropriate time base to pass for this value from
// AVFoundation playback objects like [AVSampleBufferRenderSynchronizer]. You
// can also create one manually using the
// [CMTimebaseCreateWithSourceClock(allocator:sourceClock:timebaseOut:)]
// function.
// //
// [CMTimebaseCreateWithSourceClock(allocator:sourceClock:timebaseOut:)]: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceClock(allocator:sourceClock:timebaseOut:)
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/transitionToItem(withIdentifier:proposingInitialTimingBasedOn:)
func (d AVDelegatingPlaybackCoordinator) TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase(itemIdentifier string, snapshotTimebase uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("transitionToItemWithIdentifier:proposingInitialTimingBasedOnTimebase:"), objc.String(itemIdentifier), snapshotTimebase)
}
// Tells the coordinator to reissue current play state commands to synchronize
// the current item to the state of other participants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/reapplyCurrentItemStateToPlaybackControlDelegate()
func (d AVDelegatingPlaybackCoordinator) ReapplyCurrentItemStateToPlaybackControlDelegate() {
	objc.Send[objc.ID](d.ID, objc.Sel("reapplyCurrentItemStateToPlaybackControlDelegate"))
}

// An identifier of the current item.
//
// # Discussion
// 
// The coordinator sets this value in a previous call to
// [TransitionToItemWithIdentifierProposingInitialTimingBasedOnTimebase].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/currentItemIdentifier
func (d AVDelegatingPlaybackCoordinator) CurrentItemIdentifier() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("currentItemIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The delegate object for the playback coordinator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinator/playbackControlDelegate
func (d AVDelegatingPlaybackCoordinator) PlaybackControlDelegate() AVPlaybackCoordinatorPlaybackControlDelegate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("playbackControlDelegate"))
	return AVPlaybackCoordinatorPlaybackControlDelegateObjectFromID(rv)
}

