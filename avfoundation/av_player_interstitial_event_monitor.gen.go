// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerInterstitialEventMonitor] class.
var (
	_AVPlayerInterstitialEventMonitorClass     AVPlayerInterstitialEventMonitorClass
	_AVPlayerInterstitialEventMonitorClassOnce sync.Once
)

func getAVPlayerInterstitialEventMonitorClass() AVPlayerInterstitialEventMonitorClass {
	_AVPlayerInterstitialEventMonitorClassOnce.Do(func() {
		_AVPlayerInterstitialEventMonitorClass = AVPlayerInterstitialEventMonitorClass{class: objc.GetClass("AVPlayerInterstitialEventMonitor")}
	})
	return _AVPlayerInterstitialEventMonitorClass
}

// GetAVPlayerInterstitialEventMonitorClass returns the class object for AVPlayerInterstitialEventMonitor.
func GetAVPlayerInterstitialEventMonitorClass() AVPlayerInterstitialEventMonitorClass {
	return getAVPlayerInterstitialEventMonitorClass()
}

type AVPlayerInterstitialEventMonitorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerInterstitialEventMonitorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerInterstitialEventMonitorClass) Alloc() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that monitors the scheduling and progress of interstitial events.
//
// # Overview
//
// This object monitors interstitial events that exist within the content of
// the primary items, such as events defined by an HLS media playlist, and
// also events managed by an [AVPlayerInterstitialEventController] object. You
// can access the schedule of interstitial events through the [AVPlayerInterstitialEventMonitor.Events]
// property.
//
// When it’s time to present an interstitial event, the system suspends
// playback of the primary item and changes its player’s [AVPlayerInterstitialEventMonitor.TimeControlStatus]
// to [AVPlayerTimeControlStatusWaitingToPlayAtSpecifiedRate] with a
// [AVPlayerInterstitialEventMonitor.ReasonForWaitingToPlay] value of [AVPlayerInterstitialEventMonitor.InterstitialEvent]. When the system
// suspends primary playback, it creates player items based on the event’s
// [AVPlayerInterstitialEventMonitor.TemplateItems] to play interstitial content. The interstitial player
// temporarily assumes the primary player’s output configuration, such as
// routing its visual output to player layers that reference the primary
// player. After the interstitial player finishes playback, or its current
// item otherwise becomes `nil`, playback of primary content resumes.
//
// # Creating a monitor
//
//   - [AVPlayerInterstitialEventMonitor.InitWithPrimaryPlayer]: Creates an observer with a player item.
//
// # Monitoring the current event
//
//   - [AVPlayerInterstitialEventMonitor.CurrentEvent]: The current interstitial event.
//
// # Monitoring the event schedule
//
//   - [AVPlayerInterstitialEventMonitor.Events]: The schedule of interstitial events.
//
// # Monitoring skipping
//
//   - [AVPlayerInterstitialEventMonitor.CurrentEventSkipControlLabel]: The skip control label for the currentEvent.
//   - [AVPlayerInterstitialEventMonitor.CurrentEventSkippableState]: The skippable event state for the currentEvent.
//
// # Accessing the players
//
//   - [AVPlayerInterstitialEventMonitor.PrimaryPlayer]: An object that plays primary content.
//   - [AVPlayerInterstitialEventMonitor.InterstitialPlayer]: An object that plays interstitial content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor
type AVPlayerInterstitialEventMonitor struct {
	objectivec.Object
}

// AVPlayerInterstitialEventMonitorFromID constructs a [AVPlayerInterstitialEventMonitor] from an objc.ID.
//
// An object that monitors the scheduling and progress of interstitial events.
func AVPlayerInterstitialEventMonitorFromID(id objc.ID) AVPlayerInterstitialEventMonitor {
	return AVPlayerInterstitialEventMonitor{objectivec.Object{ID: id}}
}

// NOTE: AVPlayerInterstitialEventMonitor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerInterstitialEventMonitor] class.
//
// # Creating a monitor
//
//   - [IAVPlayerInterstitialEventMonitor.InitWithPrimaryPlayer]: Creates an observer with a player item.
//
// # Monitoring the current event
//
//   - [IAVPlayerInterstitialEventMonitor.CurrentEvent]: The current interstitial event.
//
// # Monitoring the event schedule
//
//   - [IAVPlayerInterstitialEventMonitor.Events]: The schedule of interstitial events.
//
// # Monitoring skipping
//
//   - [IAVPlayerInterstitialEventMonitor.CurrentEventSkipControlLabel]: The skip control label for the currentEvent.
//   - [IAVPlayerInterstitialEventMonitor.CurrentEventSkippableState]: The skippable event state for the currentEvent.
//
// # Accessing the players
//
//   - [IAVPlayerInterstitialEventMonitor.PrimaryPlayer]: An object that plays primary content.
//   - [IAVPlayerInterstitialEventMonitor.InterstitialPlayer]: An object that plays interstitial content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor
type IAVPlayerInterstitialEventMonitor interface {
	objectivec.IObject

	// Topic: Creating a monitor

	// Creates an observer with a player item.
	InitWithPrimaryPlayer(primaryPlayer IAVPlayer) AVPlayerInterstitialEventMonitor

	// Topic: Monitoring the current event

	// The current interstitial event.
	CurrentEvent() IAVPlayerInterstitialEvent

	// Topic: Monitoring the event schedule

	// The schedule of interstitial events.
	Events() []AVPlayerInterstitialEvent

	// Topic: Monitoring skipping

	// The skip control label for the currentEvent.
	CurrentEventSkipControlLabel() string
	// The skippable event state for the currentEvent.
	CurrentEventSkippableState() AVPlayerInterstitialEventSkippableEventState

	// Topic: Accessing the players

	// An object that plays primary content.
	PrimaryPlayer() IAVPlayer
	// An object that plays interstitial content.
	InterstitialPlayer() IAVQueuePlayer

	// The player is waiting for an interstitial event to complete.
	InterstitialEvent() AVPlayerWaitingReason
	// The reason the player is currently waiting for playback to begin or resume.
	ReasonForWaitingToPlay() AVPlayerWaitingReason
	SetReasonForWaitingToPlay(value AVPlayerWaitingReason)
	// An array of player item configurations to use as templates for player items that play interstitial content.
	TemplateItems() IAVPlayerItem
	SetTemplateItems(value IAVPlayerItem)
	// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
	TimeControlStatus() AVPlayerTimeControlStatus
	SetTimeControlStatus(value AVPlayerTimeControlStatus)
}

// Init initializes the instance.
func (p AVPlayerInterstitialEventMonitor) Init() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerInterstitialEventMonitor) Autorelease() AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerInterstitialEventMonitor creates a new AVPlayerInterstitialEventMonitor instance.
func NewAVPlayerInterstitialEventMonitor() AVPlayerInterstitialEventMonitor {
	class := getAVPlayerInterstitialEventMonitorClass()
	rv := objc.Send[AVPlayerInterstitialEventMonitor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an observer with a player item.
//
// primaryPlayer: An object that plays the primary content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/init(primaryPlayer:)
func NewPlayerInterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IAVPlayer) AVPlayerInterstitialEventMonitor {
	instance := getAVPlayerInterstitialEventMonitorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPrimaryPlayer:"), primaryPlayer)
	return AVPlayerInterstitialEventMonitorFromID(rv)
}

// Creates an observer with a player item.
//
// primaryPlayer: An object that plays the primary content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/init(primaryPlayer:)
func (p AVPlayerInterstitialEventMonitor) InitWithPrimaryPlayer(primaryPlayer IAVPlayer) AVPlayerInterstitialEventMonitor {
	rv := objc.Send[AVPlayerInterstitialEventMonitor](p.ID, objc.Sel("initWithPrimaryPlayer:"), primaryPlayer)
	return rv
}

// A convenience initializer that creates an observer with a player item.
//
// primaryPlayer: An object that plays the primary content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventMonitorWithPrimaryPlayer:
func (_AVPlayerInterstitialEventMonitorClass AVPlayerInterstitialEventMonitorClass) InterstitialEventMonitorWithPrimaryPlayer(primaryPlayer IAVPlayer) AVPlayerInterstitialEventMonitor {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerInterstitialEventMonitorClass.class), objc.Sel("interstitialEventMonitorWithPrimaryPlayer:"), primaryPlayer)
	return AVPlayerInterstitialEventMonitorFromID(rv)
}

// The current interstitial event.
//
// # Discussion
//
// The value is `nil` when primary content is playing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEvent
func (p AVPlayerInterstitialEventMonitor) CurrentEvent() IAVPlayerInterstitialEvent {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentEvent"))
	return AVPlayerInterstitialEventFromID(objc.ID(rv))
}

// The schedule of interstitial events.
//
// # Discussion
//
// When the primary player’s content specifies the schedule of interstitial
// events intrinsically, this property value typically changes whenever
// primary player’s [CurrentItem] changes. For HLS content that specifies
// interstitials using of [DATERANGE] tags, the value of this property may
// also change whenever the set of [DATERANGE] tags in the current item’s
// media playlist changes.
//
// When you specify the schedule of interstitial events using an
// [AVPlayerInterstitialEventController], this property value changes only
// when you update the interstitial event controller’s schedule.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/events
func (p AVPlayerInterstitialEventMonitor) Events() []AVPlayerInterstitialEvent {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("events"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerInterstitialEvent {
		return AVPlayerInterstitialEventFromID(id)
	})
}

// The skip control label for the currentEvent.
//
// # Discussion
//
// If a localizedStringsBundle has been set on the
// AVPlayerInterstitialEventController, and a
// skipControlLocalizedLabelBundleKey is set on the currentEvent, then this
// value will be the localized string that was matched to the event’s
// skipControlLocalizedLabelBundleKey for the corresponding system language in
// the supplied Bundle, if any. If currentEvent is nil, then the value will be
// nil.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkipControlLabel
func (p AVPlayerInterstitialEventMonitor) CurrentEventSkipControlLabel() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentEventSkipControlLabel"))
	return foundation.NSStringFromID(rv).String()
}

// The skippable event state for the currentEvent.
//
// # Discussion
//
// If currentEvent is nil, then the value will be
// AVPlayerInterstitialEventSkippableEventStateNotSkippable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippableState
func (p AVPlayerInterstitialEventMonitor) CurrentEventSkippableState() AVPlayerInterstitialEventSkippableEventState {
	rv := objc.Send[AVPlayerInterstitialEventSkippableEventState](p.ID, objc.Sel("currentEventSkippableState"))
	return AVPlayerInterstitialEventSkippableEventState(rv)
}

// An object that plays primary content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/primaryPlayer
func (p AVPlayerInterstitialEventMonitor) PrimaryPlayer() IAVPlayer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("primaryPlayer"))
	return AVPlayerFromID(objc.ID(rv))
}

// An object that plays interstitial content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialPlayer
func (p AVPlayerInterstitialEventMonitor) InterstitialPlayer() IAVQueuePlayer {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("interstitialPlayer"))
	return AVQueuePlayerFromID(objc.ID(rv))
}

// The player is waiting for an interstitial event to complete.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayer/waitingreason/interstitialevent
func (p AVPlayerInterstitialEventMonitor) InterstitialEvent() AVPlayerWaitingReason {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("AVPlayerWaitingDuringInterstitialEventReason"))
	return AVPlayerWaitingReason(foundation.NSStringFromID(rv).String())
}

// The reason the player is currently waiting for playback to begin or resume.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayer/reasonforwaitingtoplay
func (p AVPlayerInterstitialEventMonitor) ReasonForWaitingToPlay() AVPlayerWaitingReason {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("reasonForWaitingToPlay"))
	return AVPlayerWaitingReason(foundation.NSStringFromID(rv).String())
}
func (p AVPlayerInterstitialEventMonitor) SetReasonForWaitingToPlay(value AVPlayerWaitingReason) {
	objc.Send[struct{}](p.ID, objc.Sel("setReasonForWaitingToPlay:"), objc.String(string(value)))
}

// An array of player item configurations to use as templates for player items
// that play interstitial content.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayerinterstitialevent/templateitems
func (p AVPlayerInterstitialEventMonitor) TemplateItems() IAVPlayerItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("templateItems"))
	return AVPlayerItemFromID(objc.ID(rv))
}
func (p AVPlayerInterstitialEventMonitor) SetTemplateItems(value IAVPlayerItem) {
	objc.Send[struct{}](p.ID, objc.Sel("setTemplateItems:"), value)
}

// A value that indicates whether playback is in progress, paused
// indefinitely, or waiting for network conditions to improve.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayer/timecontrolstatus-swift.property
func (p AVPlayerInterstitialEventMonitor) TimeControlStatus() AVPlayerTimeControlStatus {
	rv := objc.Send[AVPlayerTimeControlStatus](p.ID, objc.Sel("timeControlStatus"))
	return AVPlayerTimeControlStatus(rv)
}
func (p AVPlayerInterstitialEventMonitor) SetTimeControlStatus(value AVPlayerTimeControlStatus) {
	objc.Send[struct{}](p.ID, objc.Sel("setTimeControlStatus:"), value)
}
