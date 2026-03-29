// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayer] class.
var (
	_AVPlayerClass     AVPlayerClass
	_AVPlayerClassOnce sync.Once
)

func getAVPlayerClass() AVPlayerClass {
	_AVPlayerClassOnce.Do(func() {
		_AVPlayerClass = AVPlayerClass{class: objc.GetClass("AVPlayer")}
	})
	return _AVPlayerClass
}

// GetAVPlayerClass returns the class object for AVPlayer.
func GetAVPlayerClass() AVPlayerClass {
	return getAVPlayerClass()
}

type AVPlayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerClass) Alloc() AVPlayer {
	rv := objc.Send[AVPlayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides the interface to control the player’s transport
// behavior.
//
// # Overview
// 
// A player is a controller object that manages the playback and timing of a
// media asset. Use an instance of [AVPlayer] to play local and remote
// file-based media, such as QuickTime movies and MP3 audio files, as well as
// audiovisual media served using HTTP Live Streaming.
// 
// Use a player object to play a single media asset. You can reuse the player
// instance to play additional media assets using its
// [AVPlayer.ReplaceCurrentItemWithPlayerItem] method, but it manages the playback of
// only a single media asset at a time. The framework also provides a subclass
// called [AVQueuePlayer] that you can use to manage the playback of a queue
// of media assets.
// 
// You use an [AVPlayer] to play media assets, which AVFoundation represents
// using the [AVAsset] class. [AVAsset] only models the aspects of the media,
// such as its duration or creation date, and on its own, isn’t suitable for
// playback with an [AVPlayer]. To play an asset, you create an instance of
// its counterpart found in [AVPlayerItem]. This object models the timing and
// presentation state of an asset played by an instance of [AVPlayer]. See the
// [AVPlayerItem] reference for more details.
// 
// [AVPlayer] is a dynamic object whose state continuously changes. There are
// two approaches you can use to observe a player’s state:
// 
// - You can use key-value observing (KVO) to observe state changes to many of
// the player’s dynamic properties, such as its [AVPlayer.CurrentItem] or its
// playback [AVPlayer.Rate]. - KVO works well for general state observations, but
// isn’t intended for observing continuously changing state like the
// player’s time. [AVPlayer] provides two methods to observe time changes: -
// [AVPlayer.AddPeriodicTimeObserverForIntervalQueueUsingBlock] -
// [AVPlayer.AddBoundaryTimeObserverForTimesQueueUsingBlock]
// 
// These methods let you observe time changes either periodically or by
// boundary, respectively. As changes occur, invoke the callback block or
// closure you supply to these methods to give you the opportunity to take
// some action such as updating the state of your player’s user interface.
// 
// [AVPlayer] and [AVPlayerItem] are nonvisual objects, meaning that on their
// own they’re unable to present an asset’s video onscreen. There are two
// primary approaches you use to present your video content onscreen:
// 
// - The best way to present your video content is with the AVKit
// framework’s [AVPlayerViewController] class in iOS and tvOS, or the
// [AVPlayerView] class in macOS. These classes present the video content,
// along with playback controls and other media features giving you a
// full-featured playback experience. - When building a custom interface for
// your player, use [AVPlayerLayer]. You can set this layer a view’s backing
// layer or add it directly to the layer hierarchy. Unlike [AVPlayerView] and
// [AVPlayerViewController], a player layer doesn’t present any playback
// controls—it only presents the visual content onscreen. It’s up to you
// to build the playback transport controls to play, pause, and seek through
// the media.
// 
// Alongside the visual content presented with AVKit or [AVPlayerLayer], you
// can also present animated content synchronized with the player’s timing
// using [AVSynchronizedLayer]. Use a synchronized layer pass along player
// timing to its layer subtree. You can use [AVSynchronizedLayer] to build
// custom effects in Core Animation, such as animated lower thirds or video
// transitions, and have them play in sync with the timing of the player’s
// current [AVPlayerItem].
//
// [AVPlayerViewController]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController
// [AVPlayerView]: https://developer.apple.com/documentation/AVKit/AVPlayerView
//
// # Creating a player
//
//   - [AVPlayer.InitWithURL]: Creates a new player to play a single audiovisual resource referenced by a given URL.
//   - [AVPlayer.InitWithPlayerItem]: Creates a new player to play the specified player item.
//
// # Managing the player item
//
//   - [AVPlayer.CurrentItem]: The item for which the player is currently controlling playback.
//   - [AVPlayer.ReplaceCurrentItemWithPlayerItem]: Replaces the current item with a new item.
//
// # Determining player readiness
//
//   - [AVPlayer.Status]: A value that indicates the readiness of a player object for playback.
//   - [AVPlayer.Error]: An error that caused a failure.
//
// # Controlling playback
//
//   - [AVPlayer.DefaultRate]: A default rate at which to begin playback.
//   - [AVPlayer.SetDefaultRate]
//   - [AVPlayer.Play]: Begins playback of the current item.
//   - [AVPlayer.Pause]: Pauses playback of the current item.
//   - [AVPlayer.Rate]: The current playback rate.
//   - [AVPlayer.SetRate]
//
// # Observing playback time
//
//   - [AVPlayer.CurrentTime]: Returns the current time of the current player item.
//   - [AVPlayer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]: Requests the periodic invocation of a given block during playback to report changing time.
//   - [AVPlayer.AddBoundaryTimeObserverForTimesQueueUsingBlock]: Requests the invocation of a block when specified times are traversed during normal playback.
//   - [AVPlayer.RemoveTimeObserver]: Cancels a previously registered periodic or boundary time observer.
//
// # Seeking through media
//
//   - [AVPlayer.SeekToTime]: Requests that the player seek to a specified time.
//   - [AVPlayer.SeekToTimeCompletionHandler]: Requests that the player seek to a specified time, and to notify you when the seek is complete.
//   - [AVPlayer.SeekToTimeToleranceBeforeToleranceAfter]: Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values.
//   - [AVPlayer.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]: Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values, and to notify you when the seek is complete.
//   - [AVPlayer.SeekToDate]: Requests that the player seek to a specified date.
//   - [AVPlayer.SeekToDateCompletionHandler]: Requests that the player seek to a specified date, and to notify you when the seek is complete.
//
// # Configuring waiting behavior
//
//   - [AVPlayer.AutomaticallyWaitsToMinimizeStalling]: A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling.
//   - [AVPlayer.SetAutomaticallyWaitsToMinimizeStalling]
//   - [AVPlayer.ReasonForWaitingToPlay]: The reason the player is currently waiting for playback to begin or resume.
//   - [AVPlayer.TimeControlStatus]: A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//   - [AVPlayer.PlayImmediatelyAtRate]: Plays the available media data immediately, at the specified rate.
//
// # Responding when playback ends
//
//   - [AVPlayer.ActionAtItemEnd]: The action to perform when the current player item has finished playing.
//   - [AVPlayer.SetActionAtItemEnd]
//
// # Configuring media selection criteria
//
//   - [AVPlayer.AppliesMediaSelectionCriteriaAutomatically]: A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items.
//   - [AVPlayer.SetAppliesMediaSelectionCriteriaAutomatically]
//   - [AVPlayer.MediaSelectionCriteriaForMediaCharacteristic]: Returns the automatic selection criteria for media items with the specified media characteristic.
//   - [AVPlayer.SetMediaSelectionCriteriaForMediaCharacteristic]: Applies automatic selection criteria for media that has the specified media characteristic.
//
// # Accessing player output
//
//   - [AVPlayer.VideoOutput]: The video output for this player.
//   - [AVPlayer.SetVideoOutput]
//
// # Configuring audio behavior
//
//   - [AVPlayer.Volume]: The audio playback volume for the player.
//   - [AVPlayer.SetVolume]
//   - [AVPlayer.Muted]: A Boolean value that indicates whether the audio output of the player is muted.
//   - [AVPlayer.SetMuted]
//   - [AVPlayer.AllowedAudioSpatializationFormats]: The source audio channel layouts the player item supports for spatialization.
//   - [AVPlayer.SetAllowedAudioSpatializationFormats]
//   - [AVPlayer.IsAudioSpatializationAllowed]: A Boolean value that indicates whether the player item allows spatialized audio playback.
//   - [AVPlayer.SetIsAudioSpatializationAllowed]
//
// # Configuring background playback
//
//   - [AVPlayer.AudiovisualBackgroundPlaybackPolicy]: A policy that determines how playback of audiovisual media continues when the app transitions to the background.
//   - [AVPlayer.SetAudiovisualBackgroundPlaybackPolicy]
//
// # Managing external playback
//
//   - [AVPlayer.AllowsExternalPlayback]: A Boolean value that indicates whether the player allows switching to external playback mode.
//   - [AVPlayer.SetAllowsExternalPlayback]
//   - [AVPlayer.ExternalPlaybackActive]: A Boolean value that indicates whether the player is currently playing video in external playback mode.
//
// # Coordinating playback
//
//   - [AVPlayer.PlaybackCoordinator]: The playback coordinator for the player.
//
// # Synchronizing multiple players
//
//   - [AVPlayer.SetRateTimeAtHostTime]: Synchronizes the playback rate and time of the current item with an external source.
//   - [AVPlayer.PrerollAtRateCompletionHandler]: Begins loading media data to prime the media pipelines for playback.
//   - [AVPlayer.CancelPendingPrerolls]: Cancels any pending preroll requests and invokes the corresponding completion handlers, if present.
//   - [AVPlayer.SourceClock]: A clock the player uses for item time bases.
//   - [AVPlayer.SetSourceClock]
//   - [AVPlayer.MasterClock]: The host clock for item time bases.
//   - [AVPlayer.SetMasterClock]
//
// # Preventing sleep and backgrounding
//
//   - [AVPlayer.PreventsDisplaySleepDuringVideoPlayback]: A Boolean value that indicates whether video playback prevents display and device sleep.
//   - [AVPlayer.SetPreventsDisplaySleepDuringVideoPlayback]
//
// # Determining content protections
//
//   - [AVPlayer.OutputObscuredDueToInsufficientExternalProtection]: A Boolean value that indicates whether output is being obscured because of insufficient external protection.
//
// # Configuring audio and video devices
//
//   - [AVPlayer.AudioOutputDeviceUniqueID]: Specifies the unique ID of the Core Audio output device used to play audio.
//   - [AVPlayer.SetAudioOutputDeviceUniqueID]
//   - [AVPlayer.PreferredVideoDecoderGPURegistryID]: The registry identifier for the GPU used for video decoding.
//   - [AVPlayer.SetPreferredVideoDecoderGPURegistryID]
//
// # Configuring the network resource priority
//
//   - [AVPlayer.NetworkResourcePriority]: Indicates the priority of this player for network bandwidth resource distribution.
//   - [AVPlayer.SetNetworkResourcePriority]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer
type AVPlayer struct {
	objectivec.Object
}

// AVPlayerFromID constructs a [AVPlayer] from an objc.ID.
//
// An object that provides the interface to control the player’s transport
// behavior.
func AVPlayerFromID(id objc.ID) AVPlayer {
	return AVPlayer{objectivec.Object{ID: id}}
}
// NOTE: AVPlayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayer] class.
//
// # Creating a player
//
//   - [IAVPlayer.InitWithURL]: Creates a new player to play a single audiovisual resource referenced by a given URL.
//   - [IAVPlayer.InitWithPlayerItem]: Creates a new player to play the specified player item.
//
// # Managing the player item
//
//   - [IAVPlayer.CurrentItem]: The item for which the player is currently controlling playback.
//   - [IAVPlayer.ReplaceCurrentItemWithPlayerItem]: Replaces the current item with a new item.
//
// # Determining player readiness
//
//   - [IAVPlayer.Status]: A value that indicates the readiness of a player object for playback.
//   - [IAVPlayer.Error]: An error that caused a failure.
//
// # Controlling playback
//
//   - [IAVPlayer.DefaultRate]: A default rate at which to begin playback.
//   - [IAVPlayer.SetDefaultRate]
//   - [IAVPlayer.Play]: Begins playback of the current item.
//   - [IAVPlayer.Pause]: Pauses playback of the current item.
//   - [IAVPlayer.Rate]: The current playback rate.
//   - [IAVPlayer.SetRate]
//
// # Observing playback time
//
//   - [IAVPlayer.CurrentTime]: Returns the current time of the current player item.
//   - [IAVPlayer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]: Requests the periodic invocation of a given block during playback to report changing time.
//   - [IAVPlayer.AddBoundaryTimeObserverForTimesQueueUsingBlock]: Requests the invocation of a block when specified times are traversed during normal playback.
//   - [IAVPlayer.RemoveTimeObserver]: Cancels a previously registered periodic or boundary time observer.
//
// # Seeking through media
//
//   - [IAVPlayer.SeekToTime]: Requests that the player seek to a specified time.
//   - [IAVPlayer.SeekToTimeCompletionHandler]: Requests that the player seek to a specified time, and to notify you when the seek is complete.
//   - [IAVPlayer.SeekToTimeToleranceBeforeToleranceAfter]: Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values.
//   - [IAVPlayer.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]: Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values, and to notify you when the seek is complete.
//   - [IAVPlayer.SeekToDate]: Requests that the player seek to a specified date.
//   - [IAVPlayer.SeekToDateCompletionHandler]: Requests that the player seek to a specified date, and to notify you when the seek is complete.
//
// # Configuring waiting behavior
//
//   - [IAVPlayer.AutomaticallyWaitsToMinimizeStalling]: A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling.
//   - [IAVPlayer.SetAutomaticallyWaitsToMinimizeStalling]
//   - [IAVPlayer.ReasonForWaitingToPlay]: The reason the player is currently waiting for playback to begin or resume.
//   - [IAVPlayer.TimeControlStatus]: A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
//   - [IAVPlayer.PlayImmediatelyAtRate]: Plays the available media data immediately, at the specified rate.
//
// # Responding when playback ends
//
//   - [IAVPlayer.ActionAtItemEnd]: The action to perform when the current player item has finished playing.
//   - [IAVPlayer.SetActionAtItemEnd]
//
// # Configuring media selection criteria
//
//   - [IAVPlayer.AppliesMediaSelectionCriteriaAutomatically]: A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items.
//   - [IAVPlayer.SetAppliesMediaSelectionCriteriaAutomatically]
//   - [IAVPlayer.MediaSelectionCriteriaForMediaCharacteristic]: Returns the automatic selection criteria for media items with the specified media characteristic.
//   - [IAVPlayer.SetMediaSelectionCriteriaForMediaCharacteristic]: Applies automatic selection criteria for media that has the specified media characteristic.
//
// # Accessing player output
//
//   - [IAVPlayer.VideoOutput]: The video output for this player.
//   - [IAVPlayer.SetVideoOutput]
//
// # Configuring audio behavior
//
//   - [IAVPlayer.Volume]: The audio playback volume for the player.
//   - [IAVPlayer.SetVolume]
//   - [IAVPlayer.Muted]: A Boolean value that indicates whether the audio output of the player is muted.
//   - [IAVPlayer.SetMuted]
//   - [IAVPlayer.AllowedAudioSpatializationFormats]: The source audio channel layouts the player item supports for spatialization.
//   - [IAVPlayer.SetAllowedAudioSpatializationFormats]
//   - [IAVPlayer.IsAudioSpatializationAllowed]: A Boolean value that indicates whether the player item allows spatialized audio playback.
//   - [IAVPlayer.SetIsAudioSpatializationAllowed]
//
// # Configuring background playback
//
//   - [IAVPlayer.AudiovisualBackgroundPlaybackPolicy]: A policy that determines how playback of audiovisual media continues when the app transitions to the background.
//   - [IAVPlayer.SetAudiovisualBackgroundPlaybackPolicy]
//
// # Managing external playback
//
//   - [IAVPlayer.AllowsExternalPlayback]: A Boolean value that indicates whether the player allows switching to external playback mode.
//   - [IAVPlayer.SetAllowsExternalPlayback]
//   - [IAVPlayer.ExternalPlaybackActive]: A Boolean value that indicates whether the player is currently playing video in external playback mode.
//
// # Coordinating playback
//
//   - [IAVPlayer.PlaybackCoordinator]: The playback coordinator for the player.
//
// # Synchronizing multiple players
//
//   - [IAVPlayer.SetRateTimeAtHostTime]: Synchronizes the playback rate and time of the current item with an external source.
//   - [IAVPlayer.PrerollAtRateCompletionHandler]: Begins loading media data to prime the media pipelines for playback.
//   - [IAVPlayer.CancelPendingPrerolls]: Cancels any pending preroll requests and invokes the corresponding completion handlers, if present.
//   - [IAVPlayer.SourceClock]: A clock the player uses for item time bases.
//   - [IAVPlayer.SetSourceClock]
//   - [IAVPlayer.MasterClock]: The host clock for item time bases.
//   - [IAVPlayer.SetMasterClock]
//
// # Preventing sleep and backgrounding
//
//   - [IAVPlayer.PreventsDisplaySleepDuringVideoPlayback]: A Boolean value that indicates whether video playback prevents display and device sleep.
//   - [IAVPlayer.SetPreventsDisplaySleepDuringVideoPlayback]
//
// # Determining content protections
//
//   - [IAVPlayer.OutputObscuredDueToInsufficientExternalProtection]: A Boolean value that indicates whether output is being obscured because of insufficient external protection.
//
// # Configuring audio and video devices
//
//   - [IAVPlayer.AudioOutputDeviceUniqueID]: Specifies the unique ID of the Core Audio output device used to play audio.
//   - [IAVPlayer.SetAudioOutputDeviceUniqueID]
//   - [IAVPlayer.PreferredVideoDecoderGPURegistryID]: The registry identifier for the GPU used for video decoding.
//   - [IAVPlayer.SetPreferredVideoDecoderGPURegistryID]
//
// # Configuring the network resource priority
//
//   - [IAVPlayer.NetworkResourcePriority]: Indicates the priority of this player for network bandwidth resource distribution.
//   - [IAVPlayer.SetNetworkResourcePriority]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer
type IAVPlayer interface {
	objectivec.IObject

	// Topic: Creating a player

	// Creates a new player to play a single audiovisual resource referenced by a given URL.
	InitWithURL(URL foundation.INSURL) AVPlayer
	// Creates a new player to play the specified player item.
	InitWithPlayerItem(item IAVPlayerItem) AVPlayer

	// Topic: Managing the player item

	// The item for which the player is currently controlling playback.
	CurrentItem() IAVPlayerItem
	// Replaces the current item with a new item.
	ReplaceCurrentItemWithPlayerItem(item IAVPlayerItem)

	// Topic: Determining player readiness

	// A value that indicates the readiness of a player object for playback.
	Status() AVPlayerStatus
	// An error that caused a failure.
	Error() foundation.INSError

	// Topic: Controlling playback

	// A default rate at which to begin playback.
	DefaultRate() float32
	SetDefaultRate(value float32)
	// Begins playback of the current item.
	Play()
	// Pauses playback of the current item.
	Pause()
	// The current playback rate.
	Rate() float32
	SetRate(value float32)

	// Topic: Observing playback time

	// Returns the current time of the current player item.
	CurrentTime() coremedia.CMTime
	// Requests the periodic invocation of a given block during playback to report changing time.
	AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.CMTime, queue dispatch.Queue, block CMTimeHandler) objectivec.IObject
	// Requests the invocation of a block when specified times are traversed during normal playback.
	AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.NSValue, queue dispatch.Queue, block VoidHandler) objectivec.IObject
	// Cancels a previously registered periodic or boundary time observer.
	RemoveTimeObserver(observer objectivec.IObject)

	// Topic: Seeking through media

	// Requests that the player seek to a specified time.
	SeekToTime(time coremedia.CMTime)
	// Requests that the player seek to a specified time, and to notify you when the seek is complete.
	SeekToTimeCompletionHandler(time coremedia.CMTime, completionHandler BoolHandler)
	// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values.
	SeekToTimeToleranceBeforeToleranceAfter(time coremedia.CMTime, toleranceBefore coremedia.CMTime, toleranceAfter coremedia.CMTime)
	// Requests that the player seek to a specified time with the amount of accuracy specified by the time tolerance values, and to notify you when the seek is complete.
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time coremedia.CMTime, toleranceBefore coremedia.CMTime, toleranceAfter coremedia.CMTime, completionHandler BoolHandler)
	// Requests that the player seek to a specified date.
	SeekToDate(date foundation.INSDate)
	// Requests that the player seek to a specified date, and to notify you when the seek is complete.
	SeekToDateCompletionHandler(date foundation.INSDate, completionHandler BoolHandler)

	// Topic: Configuring waiting behavior

	// A Boolean value that indicates whether the player should automatically delay playback in order to minimize stalling.
	AutomaticallyWaitsToMinimizeStalling() bool
	SetAutomaticallyWaitsToMinimizeStalling(value bool)
	// The reason the player is currently waiting for playback to begin or resume.
	ReasonForWaitingToPlay() AVPlayerWaitingReason
	// A value that indicates whether playback is in progress, paused indefinitely, or waiting for network conditions to improve.
	TimeControlStatus() AVPlayerTimeControlStatus
	// Plays the available media data immediately, at the specified rate.
	PlayImmediatelyAtRate(rate float32)

	// Topic: Responding when playback ends

	// The action to perform when the current player item has finished playing.
	ActionAtItemEnd() AVPlayerActionAtItemEnd
	SetActionAtItemEnd(value AVPlayerActionAtItemEnd)

	// Topic: Configuring media selection criteria

	// A Boolean value that indicates whether the receiver should apply the current selection criteria automatically to player items.
	AppliesMediaSelectionCriteriaAutomatically() bool
	SetAppliesMediaSelectionCriteriaAutomatically(value bool)
	// Returns the automatic selection criteria for media items with the specified media characteristic.
	MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) IAVPlayerMediaSelectionCriteria
	// Applies automatic selection criteria for media that has the specified media characteristic.
	SetMediaSelectionCriteriaForMediaCharacteristic(criteria IAVPlayerMediaSelectionCriteria, mediaCharacteristic AVMediaCharacteristic)

	// Topic: Accessing player output

	// The video output for this player.
	VideoOutput() IAVPlayerVideoOutput
	SetVideoOutput(value IAVPlayerVideoOutput)

	// Topic: Configuring audio behavior

	// The audio playback volume for the player.
	Volume() float32
	SetVolume(value float32)
	// A Boolean value that indicates whether the audio output of the player is muted.
	Muted() bool
	SetMuted(value bool)
	// The source audio channel layouts the player item supports for spatialization.
	AllowedAudioSpatializationFormats() AVAudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AVAudioSpatializationFormats)
	// A Boolean value that indicates whether the player item allows spatialized audio playback.
	IsAudioSpatializationAllowed() bool
	SetIsAudioSpatializationAllowed(value bool)

	// Topic: Configuring background playback

	// A policy that determines how playback of audiovisual media continues when the app transitions to the background.
	AudiovisualBackgroundPlaybackPolicy() AVPlayerAudiovisualBackgroundPlaybackPolicy
	SetAudiovisualBackgroundPlaybackPolicy(value AVPlayerAudiovisualBackgroundPlaybackPolicy)

	// Topic: Managing external playback

	// A Boolean value that indicates whether the player allows switching to external playback mode.
	AllowsExternalPlayback() bool
	SetAllowsExternalPlayback(value bool)
	// A Boolean value that indicates whether the player is currently playing video in external playback mode.
	ExternalPlaybackActive() bool

	// Topic: Coordinating playback

	// The playback coordinator for the player.
	PlaybackCoordinator() IAVPlayerPlaybackCoordinator

	// Topic: Synchronizing multiple players

	// Synchronizes the playback rate and time of the current item with an external source.
	SetRateTimeAtHostTime(rate float32, itemTime coremedia.CMTime, hostClockTime coremedia.CMTime)
	// Begins loading media data to prime the media pipelines for playback.
	PrerollAtRateCompletionHandler(rate float32, completionHandler BoolHandler)
	// Cancels any pending preroll requests and invokes the corresponding completion handlers, if present.
	CancelPendingPrerolls()
	// A clock the player uses for item time bases.
	SourceClock() uintptr
	SetSourceClock(value uintptr)
	// The host clock for item time bases.
	MasterClock() uintptr
	SetMasterClock(value uintptr)

	// Topic: Preventing sleep and backgrounding

	// A Boolean value that indicates whether video playback prevents display and device sleep.
	PreventsDisplaySleepDuringVideoPlayback() bool
	SetPreventsDisplaySleepDuringVideoPlayback(value bool)

	// Topic: Determining content protections

	// A Boolean value that indicates whether output is being obscured because of insufficient external protection.
	OutputObscuredDueToInsufficientExternalProtection() bool

	// Topic: Configuring audio and video devices

	// Specifies the unique ID of the Core Audio output device used to play audio.
	AudioOutputDeviceUniqueID() string
	SetAudioOutputDeviceUniqueID(value string)
	// The registry identifier for the GPU used for video decoding.
	PreferredVideoDecoderGPURegistryID() uint64
	SetPreferredVideoDecoderGPURegistryID(value uint64)

	// Topic: Configuring the network resource priority

	// Indicates the priority of this player for network bandwidth resource distribution.
	NetworkResourcePriority() AVPlayerNetworkResourcePriority
	SetNetworkResourcePriority(value AVPlayerNetworkResourcePriority)
}

// Init initializes the instance.
func (p AVPlayer) Init() AVPlayer {
	rv := objc.Send[AVPlayer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayer) Autorelease() AVPlayer {
	rv := objc.Send[AVPlayer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayer creates a new AVPlayer instance.
func NewAVPlayer() AVPlayer {
	class := getAVPlayerClass()
	rv := objc.Send[AVPlayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new player to play the specified player item.
//
// item: The player item to play.
//
// # Return Value
// 
// A new player initialized to play `item`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(playerItem:)
func NewPlayerWithPlayerItem(item IAVPlayerItem) AVPlayer {
	instance := getAVPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPlayerItem:"), item)
	return AVPlayerFromID(rv)
}

// Creates a new player to play a single audiovisual resource referenced by a
// given URL.
//
// URL: A URL that identifies an audiovisual resource.
//
// # Return Value
// 
// A new player instance initialized to play the audiovisual resource
// specified by [URL].
//
// # Discussion
// 
// This method implicitly creates an [AVPlayerItem] object. You can get the
// player item using [CurrentItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(url:)
func NewPlayerWithURL(URL foundation.INSURL) AVPlayer {
	instance := getAVPlayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return AVPlayerFromID(rv)
}

// Creates a new player to play a single audiovisual resource referenced by a
// given URL.
//
// URL: A URL that identifies an audiovisual resource.
//
// # Return Value
// 
// A new player instance initialized to play the audiovisual resource
// specified by [URL].
//
// # Discussion
// 
// This method implicitly creates an [AVPlayerItem] object. You can get the
// player item using [CurrentItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(url:)
func (p AVPlayer) InitWithURL(URL foundation.INSURL) AVPlayer {
	rv := objc.Send[AVPlayer](p.ID, objc.Sel("initWithURL:"), URL)
	return rv
}
// Creates a new player to play the specified player item.
//
// item: The player item to play.
//
// # Return Value
// 
// A new player initialized to play `item`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/init(playerItem:)
func (p AVPlayer) InitWithPlayerItem(item IAVPlayerItem) AVPlayer {
	rv := objc.Send[AVPlayer](p.ID, objc.Sel("initWithPlayerItem:"), item)
	return rv
}
// Replaces the current item with a new item.
//
// item: The new item for the player object to play.
//
// # Discussion
// 
// The player item replacement occurs immediately and the item becomes the
// player’s [CurrentItem]. Calling this method with the player’s current
// player item has no effect.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/replaceCurrentItem(with:)
func (p AVPlayer) ReplaceCurrentItemWithPlayerItem(item IAVPlayerItem) {
	objc.Send[objc.ID](p.ID, objc.Sel("replaceCurrentItemWithPlayerItem:"), item)
}
// Begins playback of the current item.
//
// # Discussion
// 
// Calling this method is the same as setting the [Rate] to `1.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/play()
func (p AVPlayer) Play() {
	objc.Send[objc.ID](p.ID, objc.Sel("play"))
}
// Pauses playback of the current item.
//
// # Discussion
// 
// Calling this method is the same as setting the [Rate] to `0.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/pause()
func (p AVPlayer) Pause() {
	objc.Send[objc.ID](p.ID, objc.Sel("pause"))
}
// Returns the current time of the current player item.
//
// # Return Value
// 
// The current time of the current player item.
//
// # Discussion
// 
// This property isn’t key-value observable. To observe the player’s time,
// use [AddPeriodicTimeObserverForIntervalQueueUsingBlock] or
// [AddBoundaryTimeObserverForTimesQueueUsingBlock].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/currentTime()
func (p AVPlayer) CurrentTime() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](p.ID, objc.Sel("currentTime"))
	return coremedia.CMTime(rv)
}
// Requests the periodic invocation of a given block during playback to report
// changing time.
//
// interval: The time interval at which the system invokes the block during normal
// playback, according to progress of the player’s current time.
//
// queue: The dispatch queue on which the system calls the block. Passing a
// concurrent queue isn’t supported and results in undefined behavior.
// 
// If you pass [NULL], the system uses the main queue.
//
// block: The block that the system periodically invokes.
// 
// The block takes a single parameter:
// 
// time: The time at which the system invokes the block.
//
// # Return Value
// 
// An opaque object that you pass as the argument to [RemoveTimeObserver] to
// cancel observation.
//
// # Discussion
// 
// You must maintain a strong reference to the returned value as long as you
// want the time observer to be invoked by the player. You must pair each
// invocation of this method with a corresponding call to
// [RemoveTimeObserver]. Releasing the observer object without invoking
// [RemoveTimeObserver] results in undefined behavior.
// 
// The system invokes the block periodically at the interval specified,
// interpreted according to the timeline of the current item. It also invokes
// the block whenever time jumps or playback starts or stops. If the interval
// corresponds to a very short interval in real time, the player may invoke
// the block less frequently than your app requested. Even so, the player will
// invoke the block sufficiently often for the client to update indications of
// the current time appropriately in its end-user interface.
// 
// The following example illustrates how you set up a callback the system
// invokes every half second during normal playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/addPeriodicTimeObserver(forInterval:queue:using:)
func (p AVPlayer) AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval coremedia.CMTime, queue dispatch.Queue, block CMTimeHandler) objectivec.IObject {
_block2, _ := NewCMTimeBlock(block)
	rv := objc.Send[objc.ID](p.ID, objc.Sel("addPeriodicTimeObserverForInterval:queue:usingBlock:"), interval, uintptr(queue.Handle()), _block2)
	return objectivec.Object{ID: rv}
}
// Requests the invocation of a block when specified times are traversed
// during normal playback.
//
// times: An array of [NSValue] objects containing [CMTime] values that represent the
// times at which to invoke the callback. The system raises an exception if
// you pass an empty array.
// //
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
//
// queue: A queue onto which `block` should be enqueued. Passing a concurrent queue
// is not supported and will result in undefined behavior.
// 
// If you pass `nil`, the main queue is used.
//
// block: The block to be invoked when any of the times in `times` is crossed during
// normal playback.
//
// # Return Value
// 
// An opaque object that you pass as the argument to [RemoveTimeObserver] to
// stop observation.
//
// # Discussion
// 
// Boundary times are arbitrary points of interest you define within the media
// timeline. As these times are traversed during normal playback, the block
// you provide to this method will be invoked. You must maintain a strong
// reference to the returned value as long as you want the time observer to be
// invoked by the player. Each invocation of this method should be paired with
// a corresponding call to [RemoveTimeObserver].
// 
// The player does not guarantee the callback block will always be invoked for
// each boundary time. If your times are very close together along the
// timeline (close enough that the execution of the block for one takes longer
// than the difference between them) or if a seek causes time to jump over one
// or more boundary times, time observation for any specific boundary time may
// not occur. The best practice is therefore to implement the callback block
// so it always performs its necessary calculations based solely on the
// player’s [CurrentTime].
// 
// The following example shows how you could define boundary times for each
// quarter of playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/addBoundaryTimeObserver(forTimes:queue:using:)
func (p AVPlayer) AddBoundaryTimeObserverForTimesQueueUsingBlock(times []foundation.NSValue, queue dispatch.Queue, block VoidHandler) objectivec.IObject {
_block2, _ := NewVoidBlock(block)
	rv := objc.Send[objc.ID](p.ID, objc.Sel("addBoundaryTimeObserverForTimes:queue:usingBlock:"), times, uintptr(queue.Handle()), _block2)
	return objectivec.Object{ID: rv}
}
// Cancels a previously registered periodic or boundary time observer.
//
// observer: An object returned by a previous call to
// [AddPeriodicTimeObserverForIntervalQueueUsingBlock] or
// [AddBoundaryTimeObserverForTimesQueueUsingBlock].
//
// # Discussion
// 
// Upon return, the caller is guaranteed that no new time observer blocks will
// begin executing. Depending on the calling thread and the queue used to add
// the time observer, an in-flight block may continue to execute after this
// method returns. You can guarantee synchronous time observer removal by
// enqueuing the call to `removeTimeObserver` on that queue. Alternatively,
// call `dispatch_sync(queue, ^{})` after `removeTimeObserver` to wait for any
// in-flight blocks to finish executing.
// 
// You should use this method to explicitly cancel each time observer added
// using [AddPeriodicTimeObserverForIntervalQueueUsingBlock] and
// [AddBoundaryTimeObserverForTimesQueueUsingBlock].
// 
// The following shows a common implementation to remove a registered time
// observer:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/removeTimeObserver(_:)
func (p AVPlayer) RemoveTimeObserver(observer objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeTimeObserver:"), observer)
}
// Requests that the player seek to a specified time.
//
// time: The time to which to seek.
//
// # Discussion
// 
// The time to which the player seeks may differ from the specified requested
// time for efficiency. For sample accurate seeking see
// [SeekToTimeToleranceBeforeToleranceAfter].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:)-87h2r
func (p AVPlayer) SeekToTime(time coremedia.CMTime) {
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:"), time)
}
// Requests that the player seek to a specified time, and to notify you when
// the seek is complete.
//
// time: The time to which to seek.
//
// completionHandler: The block to invoke when the seek operation has either been completed or
// been interrupted. The block takes one argument:
// 
// finished: Indicates whether the seek operation completed.
//
// # Discussion
// 
// Use this method to seek the current player item to the specified time and
// be notified when the operation completes. If the seek request completes
// without being interrupted (either by another seek request or by any other
// operation), the completion handler you provide is executed with the
// `finished` parameter set to [true].
// 
// If another seek request is already in progress when you call this method,
// the completion handler for the in-progress seek request is executed
// immediately with the `finished` parameter set to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:completionHandler:)-75bls
func (p AVPlayer) SeekToTimeCompletionHandler(time coremedia.CMTime, completionHandler BoolHandler) {
_block1, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:completionHandler:"), time, _block1)
}
// Requests that the player seek to a specified time with the amount of
// accuracy specified by the time tolerance values.
//
// time: A time to seek to.
//
// toleranceBefore: A tolerance before the target time to allow.
//
// toleranceAfter: A tolerance after the target time to allow.
//
// # Discussion
// 
// The player seeks within the range `[time-beforeTolerance,
// time+afterTolerance]`, and may differ from the specified time for
// efficiency. You can request sample accurate seeking by passing a time value
// of`kCMTimeZero` for both `toleranceBefore` and `toleranceAfter`. Sample
// accurate seeking may incur additional decoding delay which can impact
// seeking performance.
// 
// Passing `kCMTimePositiveInfinity` for both `toleranceBefore` and
// `toleranceAfter` is the same as messaging [SeekToTime] directly.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:toleranceBefore:toleranceAfter:)
func (p AVPlayer) SeekToTimeToleranceBeforeToleranceAfter(time coremedia.CMTime, toleranceBefore coremedia.CMTime, toleranceAfter coremedia.CMTime) {
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:"), time, toleranceBefore, toleranceAfter)
}
// Requests that the player seek to a specified time with the amount of
// accuracy specified by the time tolerance values, and to notify you when the
// seek is complete.
//
// time: The time to which to seek.
//
// toleranceBefore: The tolerance allowed before `time`.
//
// toleranceAfter: The tolerance allowed after `time`.
//
// completionHandler: The block to invoke when the seek operation has either been completed or
// been interrupted.
// 
// The block takes one argument:
// 
// finished: Indicated whether the seek operation completed.
//
// # Discussion
// 
// Use this method to seek to a specified time for the current player item and
// to be notified when the seek operation is complete.
// 
// The time seeked to will be within the range `[time-beforeTolerance,
// time+afterTolerance]`, and may differ from the specified time for
// efficiency. You can request sample accurate seeking by passing a time value
// of`kCMTimeZero` for both `toleranceBefore` and `toleranceAfter`. Sample
// accurate seeking may incur additional decoding delay which can impact
// seeking performance.
// 
// Invoking this method with `toleranceBefore` set to [positiveInfinity] and
// `toleranceAfter` set to [positiveInfinity] is the same as invoking
// [SeekToTime].
// 
// The completion handler for any prior seek request that is still in process
// will be invoked immediately with the `finished` parameter set to [false].
// If the new request completes without being interrupted by another seek
// request or by any other operation the specified completion handler will be
// invoked with the `finished` parameter set to [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (p AVPlayer) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time coremedia.CMTime, toleranceBefore coremedia.CMTime, toleranceAfter coremedia.CMTime, completionHandler BoolHandler) {
_block3, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, _block3)
}
// Requests that the player seek to a specified date.
//
// date: The time to which to seek.
//
// # Discussion
// 
// The time to which the player seeks may differ from the specified `date` for
// efficiency. For sample accurate seeking see
// [SeekToTimeToleranceBeforeToleranceAfter].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:)-9h9qr
func (p AVPlayer) SeekToDate(date foundation.INSDate) {
	objc.Send[objc.ID](p.ID, objc.Sel("seekToDate:"), date)
}
// Requests that the player seek to a specified date, and to notify you when
// the seek is complete.
//
// date: The time to which to seek.
//
// completionHandler: The block to invoke when the seek operation has either been completed or
// been interrupted. The block takes one argument:
// 
// finished: Indicates whether the seek operation completed.
//
// # Discussion
// 
// Use this method to seek the current player item to the specified time and
// be notified when the operation completes. If the seek request completes
// without being interrupted (either by another seek request or by any other
// operation), the completion handler you provide is executed with the
// `finished` parameter set to [true].
// 
// If another seek request is already in progress when you call this method,
// the completion handler for the in-progress seek request is executed
// immediately with the `finished` parameter set to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/seek(to:completionHandler:)-wr1l
func (p AVPlayer) SeekToDateCompletionHandler(date foundation.INSDate, completionHandler BoolHandler) {
_block1, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("seekToDate:completionHandler:"), date, _block1)
}
// Plays the available media data immediately, at the specified rate.
//
// rate: The specified playback rate.
//
// # Discussion
// 
// This method plays the available media data at the specified `rate`
// regardless of whether there is sufficient media buffered to ensure smooth
// playback. If media data exists in the playback buffer, calling this method
// changes the player’s playback rate to the specified `rate` and its
// [TimeControlStatus] to a value of [PlayerTimeControlStatusPlaying]. If the
// player has insufficient media data buffered to begin playback, the player
// will behave as if it has encountered a stall during playback, except that
// no [playbackStalledNotification] will be posted.
//
// [playbackStalledNotification]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playbackStalledNotification
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playImmediately(atRate:)
func (p AVPlayer) PlayImmediatelyAtRate(rate float32) {
	objc.Send[objc.ID](p.ID, objc.Sel("playImmediatelyAtRate:"), rate)
}
// Returns the automatic selection criteria for media items with the specified
// media characteristic.
//
// mediaCharacteristic: The media characteristic for which the selection criteria is to be
// returned. Supported values include [audible], [legible], and [visual].
// //
// [audible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/audible
// [legible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/legible
// [visual]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/visual
//
// # Return Value
// 
// The [AVPlayerMediaSelectionCriteria] for `mediaCharacteristic`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/mediaSelectionCriteria(forMediaCharacteristic:)
func (p AVPlayer) MediaSelectionCriteriaForMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) IAVPlayerMediaSelectionCriteria {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("mediaSelectionCriteriaForMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return AVPlayerMediaSelectionCriteriaFromID(rv)
}
// Applies automatic selection criteria for media that has the specified media
// characteristic.
//
// criteria: An instance of [AVPlayerMediaSelectionCriteria] that specifies the
// selection criteria.
//
// mediaCharacteristic: The media characteristic for which the selection criteria are to be
// applied. Supported values include [audible], [legible], and [visual]. See
// Media Characteristics in the `AVFoundation Constants`.
// //
// [audible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/audible
// [legible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/legible
// [visual]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/visual
//
// # Discussion
// 
// Criteria will be applied to an [AVPlayerItem] instance when:
// 
// - It is made ready to play. - Specific media selections are made by the
// [AVPlayerItem] instance using the method
// [SelectMediaOptionInMediaSelectionGroup] in a different group. The
// automatic choice in one group may be influenced by a specific selection in
// another group. - Underlying system preferences change, e.g. system
// language, accessibility captions.
// 
// Specific selections made by the [AVPlayerItem] instance using the method
// [SelectMediaOptionInMediaSelectionGroup] method within any group will
// override automatic selection in that group until the player item receives a
// [SelectMediaOptionAutomaticallyInMediaSelectionGroup] message.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/setMediaSelectionCriteria(_:forMediaCharacteristic:)
func (p AVPlayer) SetMediaSelectionCriteriaForMediaCharacteristic(criteria IAVPlayerMediaSelectionCriteria, mediaCharacteristic AVMediaCharacteristic) {
	objc.Send[objc.ID](p.ID, objc.Sel("setMediaSelectionCriteria:forMediaCharacteristic:"), criteria, objc.String(string(mediaCharacteristic)))
}
// Synchronizes the playback rate and time of the current item with an
// external source.
//
// rate: The playback rate for the item.
//
// itemTime: The precise time at which to match playback of the item. To use the current
// item’s current time, specify [invalid].
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// hostClockTime: The host time at which to synchronize playback. If you specify [invalid],
// the rate and time are set together without any external synchronization.
// //
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// # Discussion
// 
// This method adjusts the current item’s timebase so that the time in
// `itemTime` is in sync with the time in `hostClockTime`. Thus, if
// `hostClockTime` specifies a time in the past, the item’s timebase is
// adjusted to make it appear as if the item has been running at the specified
// rate since `itemTime`. And if `hostClockTime` specifies a time in the
// future, playback is adjusted backward (if possible) so that the value in
// `itemTime` occurs at the precise moment the host’s clock reaches the
// value in `hostClockTime`. If there is no content to play before the time
// specified by `itemTime`, playback holds until the two times come into sync.
// 
// This method does not ensure that media data is loaded before the timebase
// starts moving. However, if you specify a host time in the near future, that
// would give you some time to load the media data and prepare for playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/setRate(_:time:atHostTime:)
func (p AVPlayer) SetRateTimeAtHostTime(rate float32, itemTime coremedia.CMTime, hostClockTime coremedia.CMTime) {
	objc.Send[objc.ID](p.ID, objc.Sel("setRate:time:atHostTime:"), rate, itemTime, hostClockTime)
}
// Begins loading media data to prime the media pipelines for playback.
//
// rate: The playback rate to use when determining how much data to load.
//
// completionHandler: A block to execute when the player finishes the load attempt. This block
// takes a single Boolean parameter that contains [true] if the data was
// loaded or [false] if there was a problem. For example, the value might be
// [false] if the preroll was interrupted by a time change or incompatible
// rate change.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method loads data starting at the item’s current playback time. The
// current rate for the playback item should always be 0 prior to calling this
// method. After the method calls the completion handler, you can change the
// item’s playback rate to begin playback.
// 
// If the player object is not ready to play (its [Status] property is not
// [PlayerStatusReadyToPlay]), this method throws an exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preroll(atRate:completionHandler:)
func (p AVPlayer) PrerollAtRateCompletionHandler(rate float32, completionHandler BoolHandler) {
_block1, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](p.ID, objc.Sel("prerollAtRate:completionHandler:"), rate, _block1)
}
// Cancels any pending preroll requests and invokes the corresponding
// completion handlers, if present.
//
// # Discussion
// 
// This method cancels and releases the completion handlers for any pending
// prerolls. The finished parameter of the completion handlers passed to
// [PrerollAtRateCompletionHandler] will be set to `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/cancelPendingPrerolls()
func (p AVPlayer) CancelPendingPrerolls() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelPendingPrerolls"))
}

// Returns a new player initialized to play the specified player item.
//
// item: The player item to play.
//
// # Return Value
// 
// A new player initialized to play `item`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playerWithPlayerItem:
func (_AVPlayerClass AVPlayerClass) PlayerWithPlayerItem(item IAVPlayerItem) AVPlayer {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerClass.class), objc.Sel("playerWithPlayerItem:"), item)
	return AVPlayerFromID(rv)
}
// Returns a new player to play a single audiovisual resource referenced by a
// given URL.
//
// URL: A URL that identifies an audiovisual resource.
//
// # Return Value
// 
// A new player initialized to play the audiovisual resource specified by
// [URL].
//
// # Discussion
// 
// This method implicitly creates an [AVPlayerItem] object. You can get the
// player item using [CurrentItem].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playerWithURL:
func (_AVPlayerClass AVPlayerClass) PlayerWithURL(URL foundation.INSURL) AVPlayer {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerClass.class), objc.Sel("playerWithURL:"), URL)
	return AVPlayerFromID(rv)
}

// The item for which the player is currently controlling playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/currentItem
func (p AVPlayer) CurrentItem() IAVPlayerItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentItem"))
	return AVPlayerItemFromID(objc.ID(rv))
}
// A value that indicates the readiness of a player object for playback.
//
// # Discussion
// 
// If the value of this property is [PlayerStatusFailed], check the value of
// the player’s [Error] property to determine the nature of the failure. If
// a player reaches a failed state, you can’t use it for playback, and
// instead need to create a new instance.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/status-swift.property
func (p AVPlayer) Status() AVPlayerStatus {
	rv := objc.Send[AVPlayerStatus](p.ID, objc.Sel("status"))
	return AVPlayerStatus(rv)
}
// An error that caused a failure.
//
// # Discussion
// 
// By default, this value is `nil`. If a player reaches a
// [PlayerStatusFailed], the system populates this value with an error that
// describes the failure.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/error
func (p AVPlayer) Error() foundation.INSError {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// A default rate at which to begin playback.
//
// # Discussion
// 
// This value represents the default playback rate the player uses when you
// call its [Play] method. After playback begins, the rate may differ from the
// default if you modify the player’s [Rate] value, such as by calling
// [Pause].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/defaultRate
func (p AVPlayer) DefaultRate() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("defaultRate"))
	return rv
}
func (p AVPlayer) SetDefaultRate(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setDefaultRate:"), value)
}
// The current playback rate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/rate
func (p AVPlayer) Rate() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("rate"))
	return rv
}
func (p AVPlayer) SetRate(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setRate:"), value)
}
// A Boolean value that indicates whether the player should automatically
// delay playback in order to minimize stalling.
//
// # Discussion
// 
// When playing media delivered over HTTP, this property is used to determine
// if the player should automatically delay playback in order to minimize
// stalling. When this property is [true] and the player changes from a paused
// state ([Rate] of `0.0`) to a played state ([Rate] > `0.0`), the player will
// try to determine if the current item can play to its end at the currently
// specified rate. If it determines that it’s likely to encounter a stall,
// the value of the player’s [TimeControlStatus] will change to
// [PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate] and playback will
// automatically start when the likelihood of stalling has been minimized. A
// similar condition will occur during playback if the current player item’s
// playback buffer is exhausted and playback stalls. Playback will
// automatically resume when the likelihood of stalling has been minimized.
// 
// You will need to set this property to [false] when you require precise
// control over playback start times, such as if you’re are synchronizing
// multiple player instances using the [SetRateTimeAtHostTime] method. If the
// value of this property is [false], playback will start immediately when
// requested as long as the playback buffer is not empty. If the playback
// buffer becomes empty and playback stalls, the player’s
// [TimeControlStatus] will switch to [PlayerTimeControlStatusPaused] and the
// playback rate will change to `0.0`.
// 
// Changing the value of this property to [false] while the player’s
// [TimeControlStatus] is
// [PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate] and its
// [ReasonForWaitingToPlay] is [toMinimizeStalls] will cause the player to
// immediately attempt playback at the specified rate.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [toMinimizeStalls]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/toMinimizeStalls
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/automaticallyWaitsToMinimizeStalling
func (p AVPlayer) AutomaticallyWaitsToMinimizeStalling() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("automaticallyWaitsToMinimizeStalling"))
	return rv
}
func (p AVPlayer) SetAutomaticallyWaitsToMinimizeStalling(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutomaticallyWaitsToMinimizeStalling:"), value)
}
// The reason the player is currently waiting for playback to begin or resume.
//
// # Discussion
// 
// When the value of the player’s [TimeControlStatus] is
// [PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate], you can use this
// property determine the reason the player is currently waiting for playback
// to begin or resume. Possible values for this property are:
// 
// - [toMinimizeStalls]
// - [noItemToPlay]
// - [evaluatingBufferingRate]
// 
// The value of this property will be `nil` if the player’s
// [TimeControlStatus] is a value other than
// [PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate].
// 
// You can use the value of this property to conditionally show UI indicating
// the player’s waiting state. This property is observable using key-value
// observing.
//
// [evaluatingBufferingRate]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/evaluatingBufferingRate
// [noItemToPlay]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/noItemToPlay
// [toMinimizeStalls]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/toMinimizeStalls
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/reasonForWaitingToPlay
func (p AVPlayer) ReasonForWaitingToPlay() AVPlayerWaitingReason {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("reasonForWaitingToPlay"))
	return AVPlayerWaitingReason(foundation.NSStringFromID(rv).String())
}
// A value that indicates whether playback is in progress, paused
// indefinitely, or waiting for network conditions to improve.
//
// # Discussion
// 
// When the value of [AutomaticallyWaitsToMinimizeStalling] is [true], the
// player waits until your app resumes playback.
// 
// During playback, the value of the property changes between
// [PlayerTimeControlStatusPlaying] and
// [PlayerTimeControlStatusWaitingToPlayAtSpecifiedRate] depending on whether
// the player has sufficient media data to continue playback.
// 
// This property is key-value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/timeControlStatus-swift.property
func (p AVPlayer) TimeControlStatus() AVPlayerTimeControlStatus {
	rv := objc.Send[AVPlayerTimeControlStatus](p.ID, objc.Sel("timeControlStatus"))
	return AVPlayerTimeControlStatus(rv)
}
// The action to perform when the current player item has finished playing.
//
// # Discussion
// 
// For possible values, see [AVPlayer.ActionAtItemEnd].
//
// [AVPlayer.ActionAtItemEnd]: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/actionAtItemEnd-swift.property
func (p AVPlayer) ActionAtItemEnd() AVPlayerActionAtItemEnd {
	rv := objc.Send[AVPlayerActionAtItemEnd](p.ID, objc.Sel("actionAtItemEnd"))
	return AVPlayerActionAtItemEnd(rv)
}
func (p AVPlayer) SetActionAtItemEnd(value AVPlayerActionAtItemEnd) {
	objc.Send[struct{}](p.ID, objc.Sel("setActionAtItemEnd:"), value)
}
// A Boolean value that indicates whether the receiver should apply the
// current selection criteria automatically to player items.
//
// # Discussion
// 
// By default, the [AVPlayer] instance applies selection criteria based on
// system accessibility preferences. To override the default criteria for any
// media selection group, use
// [SetMediaSelectionCriteriaForMediaCharacteristic].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/appliesMediaSelectionCriteriaAutomatically
func (p AVPlayer) AppliesMediaSelectionCriteriaAutomatically() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("appliesMediaSelectionCriteriaAutomatically"))
	return rv
}
func (p AVPlayer) SetAppliesMediaSelectionCriteriaAutomatically(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAppliesMediaSelectionCriteriaAutomatically:"), value)
}
// The video output for this player.
//
// # Discussion
// 
// The value of this property is `nil` by default.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/videoOutput
func (p AVPlayer) VideoOutput() IAVPlayerVideoOutput {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("videoOutput"))
	return AVPlayerVideoOutputFromID(objc.ID(rv))
}
func (p AVPlayer) SetVideoOutput(value IAVPlayerVideoOutput) {
	objc.Send[struct{}](p.ID, objc.Sel("setVideoOutput:"), value)
}
// The audio playback volume for the player.
//
// # Discussion
// 
// A value of `0.0` indicates silence; a value of `1.0` (the default)
// indicates full audio volume for the player instance.
// 
// This property is used to control the player audio volume relative to the
// system volume. There is no programmatic way to control the system volume in
// iOS, but you can use the MediaPlayer framework’s [MPVolumeView] class to
// present a standard user interface for controlling system volume.
//
// [MPVolumeView]: https://developer.apple.com/documentation/MediaPlayer/MPVolumeView
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/volume
func (p AVPlayer) Volume() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("volume"))
	return rv
}
func (p AVPlayer) SetVolume(value float32) {
	objc.Send[struct{}](p.ID, objc.Sel("setVolume:"), value)
}
// A Boolean value that indicates whether the audio output of the player is
// muted.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isMuted
func (p AVPlayer) Muted() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isMuted"))
	return rv
}
func (p AVPlayer) SetMuted(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setMuted:"), value)
}
// The source audio channel layouts the player item supports for
// spatialization.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritem/allowedaudiospatializationformats
func (p AVPlayer) AllowedAudioSpatializationFormats() AVAudioSpatializationFormats {
	rv := objc.Send[AVAudioSpatializationFormats](p.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return AVAudioSpatializationFormats(rv)
}
func (p AVPlayer) SetAllowedAudioSpatializationFormats(value AVAudioSpatializationFormats) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}
// A Boolean value that indicates whether the player item allows spatialized
// audio playback.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritem/isaudiospatializationallowed
func (p AVPlayer) IsAudioSpatializationAllowed() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("audioSpatializationAllowed"))
	return rv
}
func (p AVPlayer) SetIsAudioSpatializationAllowed(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAudioSpatializationAllowed:"), value)
}
// A policy that determines how playback of audiovisual media continues when
// the app transitions to the background.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audiovisualBackgroundPlaybackPolicy
func (p AVPlayer) AudiovisualBackgroundPlaybackPolicy() AVPlayerAudiovisualBackgroundPlaybackPolicy {
	rv := objc.Send[AVPlayerAudiovisualBackgroundPlaybackPolicy](p.ID, objc.Sel("audiovisualBackgroundPlaybackPolicy"))
	return AVPlayerAudiovisualBackgroundPlaybackPolicy(rv)
}
func (p AVPlayer) SetAudiovisualBackgroundPlaybackPolicy(value AVPlayerAudiovisualBackgroundPlaybackPolicy) {
	objc.Send[struct{}](p.ID, objc.Sel("setAudiovisualBackgroundPlaybackPolicy:"), value)
}
// A Boolean value that indicates whether the player allows switching to
// external playback mode.
//
// # Discussion
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/allowsExternalPlayback
func (p AVPlayer) AllowsExternalPlayback() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsExternalPlayback"))
	return rv
}
func (p AVPlayer) SetAllowsExternalPlayback(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowsExternalPlayback:"), value)
}
// A Boolean value that indicates whether the player is currently playing
// video in external playback mode.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isExternalPlaybackActive
func (p AVPlayer) ExternalPlaybackActive() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isExternalPlaybackActive"))
	return rv
}
// The playback coordinator for the player.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/playbackCoordinator
func (p AVPlayer) PlaybackCoordinator() IAVPlayerPlaybackCoordinator {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("playbackCoordinator"))
	return AVPlayerPlaybackCoordinatorFromID(objc.ID(rv))
}
// A clock the player uses for item time bases.
//
// # Discussion
// 
// The default value is `nil`. Setting an explicit source clock is useful to
// synchronize video-only movies with audio that plays through a different
// audio device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/sourceClock
func (p AVPlayer) SourceClock() uintptr {
	rv := objc.Send[uintptr](p.ID, objc.Sel("sourceClock"))
	return rv
}
func (p AVPlayer) SetSourceClock(value uintptr) {
	objc.Send[struct{}](p.ID, objc.Sel("setSourceClock:"), value)
}
// The host clock for item time bases.
//
// # Discussion
// 
// The default value of this property is [NULL], which means that the host
// clock is the automatic choice. When non-[NULL], this property overrides the
// automatic choice of host clock for item time bases. This is most useful
// when you’re synchronizing video-only movies with audio from another
// source.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/masterClock
func (p AVPlayer) MasterClock() uintptr {
	rv := objc.Send[uintptr](p.ID, objc.Sel("masterClock"))
	return rv
}
func (p AVPlayer) SetMasterClock(value uintptr) {
	objc.Send[struct{}](p.ID, objc.Sel("setMasterClock:"), value)
}
// A Boolean value that indicates whether video playback prevents display and
// device sleep.
//
// # Discussion
// 
// The default value is [true] in iOS, tvOS and Mac Catalyst apps, and [false]
// in macOS.
// 
// Setting this property to [false] doesn’t force the display to sleep, it
// only stops preventing display sleep. Other apps, or frameworks within your
// app may still prevent display sleep for various reasons.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preventsDisplaySleepDuringVideoPlayback
func (p AVPlayer) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}
func (p AVPlayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}
// A Boolean value that indicates whether output is being obscured because of
// insufficient external protection.
//
// # Discussion
// 
// Items that incorporate copy protection or other forms of security might
// have their visual content obscured by the player object if the current
// device configuration does not meet the requirements for protecting the
// item. This property reports whether the player is currently obscuring the
// item. If the current item does not require external protection or if the
// device configuration sufficiently protects the item, the value of this
// property is set to [false].
// 
// You can use this property to determine whether to change your app’s user
// interface to reflect the change in visibility. You can observe changes to
// the value of this property using key-value observing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isOutputObscuredDueToInsufficientExternalProtection
func (p AVPlayer) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}
// Specifies the unique ID of the Core Audio output device used to play audio.
//
// # Discussion
// 
// The default value of this property is `nil`, indicating that the default
// audio output device is used. Otherwise the value of this property is a
// string containing the unique ID of the Core Audio output device to be used
// for audio output.
// 
// Core Audio’s [kAudioDevicePropertyDeviceUID] is a suitable source of
// audio output device unique IDs.
//
// [kAudioDevicePropertyDeviceUID]: https://developer.apple.com/documentation/CoreAudio/kAudioDevicePropertyDeviceUID
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/audioOutputDeviceUniqueID
func (p AVPlayer) AudioOutputDeviceUniqueID() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return foundation.NSStringFromID(rv).String()
}
func (p AVPlayer) SetAudioOutputDeviceUniqueID(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), objc.String(value))
}
// The registry identifier for the GPU used for video decoding.
//
// # Discussion
// 
// By default, whenever possible, the GPU associated with the display
// presenting the [CALayer] performs the video decoding. Decode transitions to
// a new GPU, if appropriate, when the [CALayer] moves to a new display. This
// property overrides this default behavior, forcing decode to prefer an
// affinity to the GPU specified regardless of which GPU displays the
// associated [CALayer]. Obtain the GPU registry ID from the GPU [MTLDevice]
// using [registryID] or from OpenGL or OpenCL.
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [MTLDevice]: https://developer.apple.com/documentation/Metal/MTLDevice
// [registryID]: https://developer.apple.com/documentation/Metal/MTLDevice/registryID
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/preferredVideoDecoderGPURegistryID
func (p AVPlayer) PreferredVideoDecoderGPURegistryID() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("preferredVideoDecoderGPURegistryID"))
	return rv
}
func (p AVPlayer) SetPreferredVideoDecoderGPURegistryID(value uint64) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredVideoDecoderGPURegistryID:"), value)
}
// Indicates the priority of this player for network bandwidth resource
// distribution.
//
// # Discussion
// 
// This value determines the priority of the player during network resource
// allocation among all other players within the same application process. The
// default value for this is AVPlayerNetworkResourcePriorityDefault.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/networkResourcePriority-swift.property
func (p AVPlayer) NetworkResourcePriority() AVPlayerNetworkResourcePriority {
	rv := objc.Send[AVPlayerNetworkResourcePriority](p.ID, objc.Sel("networkResourcePriority"))
	return AVPlayerNetworkResourcePriority(rv)
}
func (p AVPlayer) SetNetworkResourcePriority(value AVPlayerNetworkResourcePriority) {
	objc.Send[struct{}](p.ID, objc.Sel("setNetworkResourcePriority:"), value)
}

// A Boolean value that indicates whether the current device can present
// content to an HDR display.
//
// # Discussion
// 
// This property is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/eligibleForHDRPlayback
func (_AVPlayerClass AVPlayerClass) EligibleForHDRPlayback() bool {
	rv := objc.Send[bool](objc.ID(_AVPlayerClass.class), objc.Sel("eligibleForHDRPlayback"))
	return rv
}
// AVPlayer and other AVFoundation types can optionally be observed using
// Swift Observation.
//
// # Discussion
// 
// When set to YES, new instances of AVPlayer, AVQueuePlayer, AVPlayerItem,
// and AVPlayerItemTrack are observable with Swift Observation. The default
// value is NO (not observable). An exception is thrown if this property is
// set YES after initializing any objects of these types, or if it is set to
// NO after any observable objects are initialized. In other words, all
// objects of these types must either be observable or not observable in an
// application instance.
// 
// For more information regarding management of class objects in SwiftUI,
// please refer to https://developer.apple.com/documentation/swiftui/state.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/isObservationEnabled
func (_AVPlayerClass AVPlayerClass) ObservationEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVPlayerClass.class), objc.Sel("isObservationEnabled"))
	return rv
}
func (_AVPlayerClass AVPlayerClass) SetObservationEnabled(value bool) {
	objc.Send[struct{}](objc.ID(_AVPlayerClass.class), objc.Sel("setObservationEnabled:"), value)
}

// AddPeriodicTimeObserverForIntervalQueueUsingBlockSync is a synchronous wrapper around [AVPlayer.AddPeriodicTimeObserverForIntervalQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayer) AddPeriodicTimeObserverForIntervalQueueUsingBlockSync(ctx context.Context, interval coremedia.CMTime, queue dispatch.Queue) (coremedia.CMTime, error) {
	done := make(chan coremedia.CMTime, 1)
	p.AddPeriodicTimeObserverForIntervalQueueUsingBlock(interval, queue, func(val coremedia.CMTime) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return coremedia.CMTime{}, ctx.Err()
	}
}

// SeekToTimeSync is a synchronous wrapper around [AVPlayer.SeekToTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayer) SeekToTimeSync(ctx context.Context, time coremedia.CMTime) (bool, error) {
	done := make(chan bool, 1)
	p.SeekToTimeCompletionHandler(time, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SeekToTimeToleranceBeforeToleranceAfterSync is a synchronous wrapper around [AVPlayer.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayer) SeekToTimeToleranceBeforeToleranceAfterSync(ctx context.Context, time coremedia.CMTime, toleranceBefore coremedia.CMTime, toleranceAfter coremedia.CMTime) (bool, error) {
	done := make(chan bool, 1)
	p.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time, toleranceBefore, toleranceAfter, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// SeekToDateSync is a synchronous wrapper around [AVPlayer.SeekToDateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayer) SeekToDateSync(ctx context.Context, date foundation.INSDate) (bool, error) {
	done := make(chan bool, 1)
	p.SeekToDateCompletionHandler(date, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// PrerollAtRate is a synchronous wrapper around [AVPlayer.PrerollAtRateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayer) PrerollAtRate(ctx context.Context, rate float32) (bool, error) {
	done := make(chan bool, 1)
	p.PrerollAtRateCompletionHandler(rate, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

