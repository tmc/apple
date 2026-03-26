// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerInterstitialEventController] class.
var (
	_AVPlayerInterstitialEventControllerClass     AVPlayerInterstitialEventControllerClass
	_AVPlayerInterstitialEventControllerClassOnce sync.Once
)

func getAVPlayerInterstitialEventControllerClass() AVPlayerInterstitialEventControllerClass {
	_AVPlayerInterstitialEventControllerClassOnce.Do(func() {
		_AVPlayerInterstitialEventControllerClass = AVPlayerInterstitialEventControllerClass{class: objc.GetClass("AVPlayerInterstitialEventController")}
	})
	return _AVPlayerInterstitialEventControllerClass
}

// GetAVPlayerInterstitialEventControllerClass returns the class object for AVPlayerInterstitialEventController.
func GetAVPlayerInterstitialEventControllerClass() AVPlayerInterstitialEventControllerClass {
	return getAVPlayerInterstitialEventControllerClass()
}

type AVPlayerInterstitialEventControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerInterstitialEventControllerClass) Alloc() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that schedules interstitial events for items played by the
// primary player.
//
// # Overview
// 
// This class is a subclass of [AVPlayerInterstitialEventMonitor] that you use
// to manage the schedule of interstitial events to present during playback of
// primary content.
//
// # Configuring the event schedule
//
//   - [AVPlayerInterstitialEventController.CancelCurrentEventWithResumptionOffset]: Cancels the playback of all currently playing and scheduled interstitial events, and resumes playback of primary content.
//   - [AVPlayerInterstitialEventController.SkipCurrentEvent]: Causes the playback of the currently playing interstital event to be abandoned.
//
// # Accessing strings
//
//   - [AVPlayerInterstitialEventController.LocalizedStringsBundle]: The bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//   - [AVPlayerInterstitialEventController.SetLocalizedStringsBundle]
//   - [AVPlayerInterstitialEventController.LocalizedStringsTableName]: The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//   - [AVPlayerInterstitialEventController.SetLocalizedStringsTableName]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController
type AVPlayerInterstitialEventController struct {
	AVPlayerInterstitialEventMonitor
}

// AVPlayerInterstitialEventControllerFromID constructs a [AVPlayerInterstitialEventController] from an objc.ID.
//
// An object that schedules interstitial events for items played by the
// primary player.
func AVPlayerInterstitialEventControllerFromID(id objc.ID) AVPlayerInterstitialEventController {
	return AVPlayerInterstitialEventController{AVPlayerInterstitialEventMonitor: AVPlayerInterstitialEventMonitorFromID(id)}
}
// NOTE: AVPlayerInterstitialEventController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerInterstitialEventController] class.
//
// # Configuring the event schedule
//
//   - [IAVPlayerInterstitialEventController.CancelCurrentEventWithResumptionOffset]: Cancels the playback of all currently playing and scheduled interstitial events, and resumes playback of primary content.
//   - [IAVPlayerInterstitialEventController.SkipCurrentEvent]: Causes the playback of the currently playing interstital event to be abandoned.
//
// # Accessing strings
//
//   - [IAVPlayerInterstitialEventController.LocalizedStringsBundle]: The bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//   - [IAVPlayerInterstitialEventController.SetLocalizedStringsBundle]
//   - [IAVPlayerInterstitialEventController.LocalizedStringsTableName]: The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
//   - [IAVPlayerInterstitialEventController.SetLocalizedStringsTableName]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController
type IAVPlayerInterstitialEventController interface {
	IAVPlayerInterstitialEventMonitor

	// Topic: Configuring the event schedule

	// Cancels the playback of all currently playing and scheduled interstitial events, and resumes playback of primary content.
	CancelCurrentEventWithResumptionOffset(resumptionOffset objectivec.IObject)
	// Causes the playback of the currently playing interstital event to be abandoned.
	SkipCurrentEvent()

	// Topic: Accessing strings

	// The bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
	LocalizedStringsBundle() foundation.NSBundle
	SetLocalizedStringsBundle(value foundation.NSBundle)
	// The name of the table in the bundle that contains the localized strings to be used by the AVPlayerInterstitialEventController.
	LocalizedStringsTableName() string
	SetLocalizedStringsTableName(value string)
}

// Init initializes the instance.
func (p AVPlayerInterstitialEventController) Init() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerInterstitialEventController) Autorelease() AVPlayerInterstitialEventController {
	rv := objc.Send[AVPlayerInterstitialEventController](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerInterstitialEventController creates a new AVPlayerInterstitialEventController instance.
func NewAVPlayerInterstitialEventController() AVPlayerInterstitialEventController {
	class := getAVPlayerInterstitialEventControllerClass()
	rv := objc.Send[AVPlayerInterstitialEventController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an event controller with a player item.
//
// primaryPlayer: A player that plays primary content. The system raises an exception you
// specify an interstitial player for this argument.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/init(primaryPlayer:)
func NewPlayerInterstitialEventControllerWithPrimaryPlayer(primaryPlayer IAVPlayer) AVPlayerInterstitialEventController {
	instance := getAVPlayerInterstitialEventControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPrimaryPlayer:"), primaryPlayer)
	return AVPlayerInterstitialEventControllerFromID(rv)
}

// Cancels the playback of all currently playing and scheduled interstitial
// events, and resumes playback of primary content.
//
// resumptionOffset: The time offset at which playback of the primary content resumes after
// interstitial playback finishes.
//
// resumptionOffset is a [coremedia.CMTime].
//
// # Discussion
// 
// When you cancel interstitial events using this method, the resumption
// offset value that you specify overrides the events’s [ResumptionOffset]
// value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/cancelCurrentEvent(withResumptionOffset:)
func (p AVPlayerInterstitialEventController) CancelCurrentEventWithResumptionOffset(resumptionOffset objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelCurrentEventWithResumptionOffset:"), resumptionOffset)
}
// Causes the playback of the currently playing interstital event to be
// abandoned.
//
// # Discussion
// 
// Note that coinciding events will NOT be skipped. This results in
// AVPlayerInterstitialEventMonitorCurrentEventSkippedNotification being
// posted. Has no effect while the currentEvent is nil.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/skipCurrentEvent()
func (p AVPlayerInterstitialEventController) SkipCurrentEvent() {
	objc.Send[objc.ID](p.ID, objc.Sel("skipCurrentEvent"))
}

// A convenience initializer that creates an event controller with a player
// item.
//
// primaryPlayer: A player that plays primary content. The system raises an exception you
// specify an interstitial player for this argument.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/interstitialEventControllerWithPrimaryPlayer:
func (_AVPlayerInterstitialEventControllerClass AVPlayerInterstitialEventControllerClass) InterstitialEventControllerWithPrimaryPlayer(primaryPlayer IAVPlayer) AVPlayerInterstitialEventController {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerInterstitialEventControllerClass.class), objc.Sel("interstitialEventControllerWithPrimaryPlayer:"), primaryPlayer)
	return AVPlayerInterstitialEventControllerFromID(rv)
}

// The bundle that contains the localized strings to be used by the
// AVPlayerInterstitialEventController.
//
// # Discussion
// 
// If the value of the property is nil, any UI elements triggered by the
// AVPlayerInterstitialEventController, such as the skip button, may contain a
// generic label based on the implementation of the UI that’s in use. To
// ensure the best available user experience in various playback
// configurations, including external playback, set a value for this property
// that provides localized translations of skip control labels.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsBundle
func (p AVPlayerInterstitialEventController) LocalizedStringsBundle() foundation.NSBundle {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("localizedStringsBundle"))
	return foundation.NSBundleFromID(objc.ID(rv))
}
func (p AVPlayerInterstitialEventController) SetLocalizedStringsBundle(value foundation.NSBundle) {
	objc.Send[struct{}](p.ID, objc.Sel("setLocalizedStringsBundle:"), value)
}
// The name of the table in the bundle that contains the localized strings to
// be used by the AVPlayerInterstitialEventController.
//
// # Discussion
// 
// If the value of the property is nil, it will default to “Localizable”
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventController/localizedStringsTableName
func (p AVPlayerInterstitialEventController) LocalizedStringsTableName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("localizedStringsTableName"))
	return foundation.NSStringFromID(rv).String()
}
func (p AVPlayerInterstitialEventController) SetLocalizedStringsTableName(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setLocalizedStringsTableName:"), objc.String(value))
}

