// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItem] class.
var (
	_AVPlayerItemClass     AVPlayerItemClass
	_AVPlayerItemClassOnce sync.Once
)

func getAVPlayerItemClass() AVPlayerItemClass {
	_AVPlayerItemClassOnce.Do(func() {
		_AVPlayerItemClass = AVPlayerItemClass{class: objc.GetClass("AVPlayerItem")}
	})
	return _AVPlayerItemClass
}

// GetAVPlayerItemClass returns the class object for AVPlayerItem.
func GetAVPlayerItemClass() AVPlayerItemClass {
	return getAVPlayerItemClass()
}

type AVPlayerItemClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemClass) Alloc() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that models the timing and presentation state of an asset during
// playback.
//
// # Overview
// 
// A player item stores a reference to an [AVAsset] object, which represents
// the media to play. If you require inspecting an asset before you enqueue it
// for playback, call its [load(_:isolation:)] method to retrieve the values
// of one or more properties. Alternatively, you can tell the player item to
// automatically load the required properties by passing them to its
// [init(asset:automaticallyLoadedAssetKeys:)] initializer. When the player
// item is ready to play, the asset properties you request are ready to use.
//
// [init(asset:automaticallyLoadedAssetKeys:)]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:automaticallyLoadedAssetKeys:)-5czjh
// [load(_:isolation:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading/load(_:isolation:)
//
// # Creating a player item
//
//   - [AVPlayerItem.InitWithURL]: Creates a player item with a specified URL.
//   - [AVPlayerItem.InitWithAsset]: Creates a player item for a specified asset.
//   - [AVPlayerItem.InitWithAssetAutomaticallyLoadedAssetKeys]: Creates a player item with the specified asset and the asset keys to automatically load.
//
// # Accessing tracks
//
//   - [AVPlayerItem.Tracks]: An array of player item track objects.
//
// # Determining readiness
//
//   - [AVPlayerItem.Status]: The status of the player item.
//   - [AVPlayerItem.Error]: The error that caused the player item to fail.
//
// # Determining playback capabilities
//
//   - [AVPlayerItem.CanPlayReverse]: A Boolean value that indicates whether the item can play in reverse.
//   - [AVPlayerItem.CanPlayFastForward]: A Boolean value that indicates whether the item can be fast forwarded.
//   - [AVPlayerItem.CanPlayFastReverse]: A Boolean value that indicates whether the item can be quickly reversed.
//   - [AVPlayerItem.CanPlaySlowForward]: A Boolean value that indicates whether the item can play slower than normal.
//   - [AVPlayerItem.CanPlaySlowReverse]: A Boolean value that indicates whether the item can play slowly backward.
//
// # Setting playback boundaries
//
//   - [AVPlayerItem.ForwardPlaybackEndTime]: The time at which forward playback ends.
//   - [AVPlayerItem.SetForwardPlaybackEndTime]
//   - [AVPlayerItem.ReversePlaybackEndTime]: The time at which reverse playback ends.
//   - [AVPlayerItem.SetReversePlaybackEndTime]
//
// # Stepping through media
//
//   - [AVPlayerItem.CanStepForward]: A Boolean value that indicates whether the item supports stepping forward.
//   - [AVPlayerItem.CanStepBackward]: A Boolean value that indicates whether the item supports stepping backward.
//   - [AVPlayerItem.StepByCount]: Moves the player item’s current time forward or backward by a specified number of steps.
//
// # Seeking through media
//
//   - [AVPlayerItem.SeekToTimeCompletionHandler]: Sets the current playback time to the specified time.
//   - [AVPlayerItem.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]: Sets the current playback time within a specified time bound and invokes the specified block when the seek operation completes or is interrupted.
//   - [AVPlayerItem.SeekToDateCompletionHandler]: Sets the current playback time to the time specified by the date object.
//   - [AVPlayerItem.CancelPendingSeeks]: Cancels any pending seek requests and invokes the corresponding completion handlers if present.
//
// # Selecting media options
//
//   - [AVPlayerItem.SelectMediaPresentationSettingForMediaSelectionGroup]: When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular presentation setting, replacing any previous preference for settings of the same media presentation selector.
//   - [AVPlayerItem.PreferredCustomMediaSelectionSchemes]: Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.
//   - [AVPlayerItem.SetPreferredCustomMediaSelectionSchemes]
//   - [AVPlayerItem.EffectiveMediaPresentationSettingsForMediaSelectionGroup]: Indicates the media presentation settings with media characteristics that are possessed by the currently selected AVMediaSelectionOption in the specified AVMediaSelectionGroup.
//   - [AVPlayerItem.SelectMediaPresentationLanguageForMediaSelectionGroup]: When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular language, replacing any previous preference for available languages of the specified group’s custom media selection scheme.
//   - [AVPlayerItem.SelectedMediaPresentationLanguageForMediaSelectionGroup]: Returns the selected media presentation language for the specified media selection group, if any language has previously been selected via use of -selectMediaPresentationLanguages:forMediaSelectionGroup:.
//   - [AVPlayerItem.SelectedMediaPresentationSettingsForMediaSelectionGroup]: Indicates the media presentation settings that have most recently been selected for each AVMediaPresentationSelector of the AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup.
//   - [AVPlayerItem.CurrentMediaSelection]: The current media selections for each of the receiver’s media selection groups.
//   - [AVPlayerItem.SelectMediaOptionInMediaSelectionGroup]: Selects a media option in a given media selection group and deselects all other options in that group.
//   - [AVPlayerItem.SelectMediaOptionAutomaticallyInMediaSelectionGroup]: Selects the media option in the specified media selection group that best matches the receiver’s automatic selection criteria.
//
// # Setting variant behavior
//
//   - [AVPlayerItem.VariantPreferences]: The preferences the player item uses when selecting variant playlists.
//   - [AVPlayerItem.SetVariantPreferences]
//   - [AVPlayerItem.StartsOnFirstEligibleVariant]: A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.
//   - [AVPlayerItem.SetStartsOnFirstEligibleVariant]
//
// # Configuring interstitial events
//
//   - [AVPlayerItem.IntegratedTimeline]: An integrated timeline that represents the player item timing including its scheduled interstitial events.
//   - [AVPlayerItem.AutomaticallyHandlesInterstitialEvents]: A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.
//   - [AVPlayerItem.SetAutomaticallyHandlesInterstitialEvents]
//   - [AVPlayerItem.TemplatePlayerItem]: The template player item that initializes this instance.
//
// # Accessing timing information
//
//   - [AVPlayerItem.CurrentTime]: Returns the current time of the item.
//   - [AVPlayerItem.CurrentDate]: Returns the current time of the item as a date.
//   - [AVPlayerItem.Duration]: The duration of the item.
//   - [AVPlayerItem.Timebase]: The timebase information for the item.
//
// # Determining available time ranges
//
//   - [AVPlayerItem.LoadedTimeRanges]: An array of time ranges indicating media data that is readily available.
//   - [AVPlayerItem.SeekableTimeRanges]: An array of time ranges within which it is possible to seek.
//
// # Determining buffering status
//
//   - [AVPlayerItem.PlaybackLikelyToKeepUp]: A Boolean value that indicates whether the item will likely play through without stalling.
//   - [AVPlayerItem.PlaybackBufferFull]: A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
//   - [AVPlayerItem.PlaybackBufferEmpty]: A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
//
// # Configuring expensive network behavior
//
//   - [AVPlayerItem.PreferredPeakBitRateForExpensiveNetworks]: A limit of network bandwidth consumption by the item when connecting over expensive networks.
//   - [AVPlayerItem.SetPreferredPeakBitRateForExpensiveNetworks]
//   - [AVPlayerItem.PreferredMaximumResolutionForExpensiveNetworks]: An upper limit on the resolution of video to download when connecting over expensive networks.
//   - [AVPlayerItem.SetPreferredMaximumResolutionForExpensiveNetworks]
//
// # Accessing text style rules
//
//   - [AVPlayerItem.TextStyleRules]: An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//   - [AVPlayerItem.SetTextStyleRules]
//
// # Accessing logging information
//
//   - [AVPlayerItem.AccessLog]: Returns an object that represents a snapshot of the network access log.
//   - [AVPlayerItem.ErrorLog]: Returns an object that represents a snapshot of the error log.
//
// # Managing time offsets
//
//   - [AVPlayerItem.AutomaticallyPreservesTimeOffsetFromLive]: A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.
//   - [AVPlayerItem.SetAutomaticallyPreservesTimeOffsetFromLive]
//   - [AVPlayerItem.RecommendedTimeOffsetFromLive]: A recommended time offset from the live time based on observed network conditions.
//   - [AVPlayerItem.ConfiguredTimeOffsetFromLive]: A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.
//   - [AVPlayerItem.SetConfiguredTimeOffsetFromLive]
//
// # Configuring presentation
//
//   - [AVPlayerItem.PresentationSize]: The size at which the visual portion of the item is presented by the player.
//   - [AVPlayerItem.PreferredMaximumResolution]: The desired maximum resolution of a video that is to be downloaded.
//   - [AVPlayerItem.SetPreferredMaximumResolution]
//   - [AVPlayerItem.VideoApertureMode]: The video aperture mode to apply during playback.
//   - [AVPlayerItem.SetVideoApertureMode]
//
// # Configuring HDR settings
//
//   - [AVPlayerItem.AppliesPerFrameHDRDisplayMetadata]: A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.
//   - [AVPlayerItem.SetAppliesPerFrameHDRDisplayMetadata]
//
// # Configuring video compositing
//
//   - [AVPlayerItem.VideoComposition]: The video composition settings to be applied during playback.
//   - [AVPlayerItem.SetVideoComposition]
//   - [AVPlayerItem.CustomVideoCompositor]: The custom video compositor.
//   - [AVPlayerItem.SeekingWaitsForVideoCompositionRendering]: A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.
//   - [AVPlayerItem.SetSeekingWaitsForVideoCompositionRendering]
//
// # Configuring audio
//
//   - [AVPlayerItem.AudioMix]: The audio mix parameters to be applied during playback.
//   - [AVPlayerItem.SetAudioMix]
//   - [AVPlayerItem.AudioTimePitchAlgorithm]: The processing algorithm used to manage audio pitch for scaled audio edits.
//   - [AVPlayerItem.SetAudioTimePitchAlgorithm]
//   - [AVPlayerItem.AllowedAudioSpatializationFormats]: The source audio channel layouts the player item supports for spatialization.
//   - [AVPlayerItem.SetAllowedAudioSpatializationFormats]
//   - [AVPlayerItem.AudioSpatializationAllowed]: A Boolean value that indicates whether the player item allows spatialized audio playback.
//   - [AVPlayerItem.SetAudioSpatializationAllowed]
//
// # Managing player item outputs
//
//   - [AVPlayerItem.Outputs]: An array of outputs associated with the player item.
//   - [AVPlayerItem.AddOutput]: Adds the specified player item output object to the receiver.
//   - [AVPlayerItem.RemoveOutput]: Removes the specified player item output object from the receiver.
//
// # Managing player item data collectors
//
//   - [AVPlayerItem.MediaDataCollectors]: The collection of associated media data collectors.
//   - [AVPlayerItem.AddMediaDataCollector]: Adds the specified media data collector to the player item’s collection of media collectors.
//   - [AVPlayerItem.RemoveMediaDataCollector]: Removes the specified media data collector from the player item’s collection of media collectors.
//
// # Configuring network behavior
//
//   - [AVPlayerItem.PreferredPeakBitRate]: The desired limit, in bits per second, of network bandwidth consumption for this item.
//   - [AVPlayerItem.SetPreferredPeakBitRate]
//   - [AVPlayerItem.PreferredForwardBufferDuration]: The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.
//   - [AVPlayerItem.SetPreferredForwardBufferDuration]
//   - [AVPlayerItem.CanUseNetworkResourcesForLiveStreamingWhilePaused]: A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.
//   - [AVPlayerItem.SetCanUseNetworkResourcesForLiveStreamingWhilePaused]
//
// # Managing playback authorization in macOS
//
//   - [AVPlayerItem.ContentAuthorizedForPlayback]: A Boolean value that indicates whether the content has been authorized by the user.
//   - [AVPlayerItem.AuthorizationRequiredForPlayback]: A Boolean value that indicates whether authorization is required to play the content.
//   - [AVPlayerItem.ApplicationAuthorizedForPlayback]: A Boolean value that indicates whether the application can be used to play the content.
//   - [AVPlayerItem.RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler]: Presents the user the opportunity to authorize the content for playback.
//   - [AVPlayerItem.ContentAuthorizationRequestStatus]: The status of the most recent content authorization request.
//   - [AVPlayerItem.CancelContentAuthorizationRequest]: Cancels the currently outstanding content authorization request.
//
// # Accessing initialization parameters
//
//   - [AVPlayerItem.Asset]: The asset provided during initialization.
//   - [AVPlayerItem.AutomaticallyLoadedAssetKeys]: The array of asset keys to be automatically loaded before the player item is ready to play.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem
type AVPlayerItem struct {
	objectivec.Object
}

// AVPlayerItemFromID constructs a [AVPlayerItem] from an objc.ID.
//
// An object that models the timing and presentation state of an asset during
// playback.
func AVPlayerItemFromID(id objc.ID) AVPlayerItem {
	return AVPlayerItem{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItem] class.
//
// # Creating a player item
//
//   - [IAVPlayerItem.InitWithURL]: Creates a player item with a specified URL.
//   - [IAVPlayerItem.InitWithAsset]: Creates a player item for a specified asset.
//   - [IAVPlayerItem.InitWithAssetAutomaticallyLoadedAssetKeys]: Creates a player item with the specified asset and the asset keys to automatically load.
//
// # Accessing tracks
//
//   - [IAVPlayerItem.Tracks]: An array of player item track objects.
//
// # Determining readiness
//
//   - [IAVPlayerItem.Status]: The status of the player item.
//   - [IAVPlayerItem.Error]: The error that caused the player item to fail.
//
// # Determining playback capabilities
//
//   - [IAVPlayerItem.CanPlayReverse]: A Boolean value that indicates whether the item can play in reverse.
//   - [IAVPlayerItem.CanPlayFastForward]: A Boolean value that indicates whether the item can be fast forwarded.
//   - [IAVPlayerItem.CanPlayFastReverse]: A Boolean value that indicates whether the item can be quickly reversed.
//   - [IAVPlayerItem.CanPlaySlowForward]: A Boolean value that indicates whether the item can play slower than normal.
//   - [IAVPlayerItem.CanPlaySlowReverse]: A Boolean value that indicates whether the item can play slowly backward.
//
// # Setting playback boundaries
//
//   - [IAVPlayerItem.ForwardPlaybackEndTime]: The time at which forward playback ends.
//   - [IAVPlayerItem.SetForwardPlaybackEndTime]
//   - [IAVPlayerItem.ReversePlaybackEndTime]: The time at which reverse playback ends.
//   - [IAVPlayerItem.SetReversePlaybackEndTime]
//
// # Stepping through media
//
//   - [IAVPlayerItem.CanStepForward]: A Boolean value that indicates whether the item supports stepping forward.
//   - [IAVPlayerItem.CanStepBackward]: A Boolean value that indicates whether the item supports stepping backward.
//   - [IAVPlayerItem.StepByCount]: Moves the player item’s current time forward or backward by a specified number of steps.
//
// # Seeking through media
//
//   - [IAVPlayerItem.SeekToTimeCompletionHandler]: Sets the current playback time to the specified time.
//   - [IAVPlayerItem.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]: Sets the current playback time within a specified time bound and invokes the specified block when the seek operation completes or is interrupted.
//   - [IAVPlayerItem.SeekToDateCompletionHandler]: Sets the current playback time to the time specified by the date object.
//   - [IAVPlayerItem.CancelPendingSeeks]: Cancels any pending seek requests and invokes the corresponding completion handlers if present.
//
// # Selecting media options
//
//   - [IAVPlayerItem.SelectMediaPresentationSettingForMediaSelectionGroup]: When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular presentation setting, replacing any previous preference for settings of the same media presentation selector.
//   - [IAVPlayerItem.PreferredCustomMediaSelectionSchemes]: Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.
//   - [IAVPlayerItem.SetPreferredCustomMediaSelectionSchemes]
//   - [IAVPlayerItem.EffectiveMediaPresentationSettingsForMediaSelectionGroup]: Indicates the media presentation settings with media characteristics that are possessed by the currently selected AVMediaSelectionOption in the specified AVMediaSelectionGroup.
//   - [IAVPlayerItem.SelectMediaPresentationLanguageForMediaSelectionGroup]: When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular language, replacing any previous preference for available languages of the specified group’s custom media selection scheme.
//   - [IAVPlayerItem.SelectedMediaPresentationLanguageForMediaSelectionGroup]: Returns the selected media presentation language for the specified media selection group, if any language has previously been selected via use of -selectMediaPresentationLanguages:forMediaSelectionGroup:.
//   - [IAVPlayerItem.SelectedMediaPresentationSettingsForMediaSelectionGroup]: Indicates the media presentation settings that have most recently been selected for each AVMediaPresentationSelector of the AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup.
//   - [IAVPlayerItem.CurrentMediaSelection]: The current media selections for each of the receiver’s media selection groups.
//   - [IAVPlayerItem.SelectMediaOptionInMediaSelectionGroup]: Selects a media option in a given media selection group and deselects all other options in that group.
//   - [IAVPlayerItem.SelectMediaOptionAutomaticallyInMediaSelectionGroup]: Selects the media option in the specified media selection group that best matches the receiver’s automatic selection criteria.
//
// # Setting variant behavior
//
//   - [IAVPlayerItem.VariantPreferences]: The preferences the player item uses when selecting variant playlists.
//   - [IAVPlayerItem.SetVariantPreferences]
//   - [IAVPlayerItem.StartsOnFirstEligibleVariant]: A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.
//   - [IAVPlayerItem.SetStartsOnFirstEligibleVariant]
//
// # Configuring interstitial events
//
//   - [IAVPlayerItem.IntegratedTimeline]: An integrated timeline that represents the player item timing including its scheduled interstitial events.
//   - [IAVPlayerItem.AutomaticallyHandlesInterstitialEvents]: A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.
//   - [IAVPlayerItem.SetAutomaticallyHandlesInterstitialEvents]
//   - [IAVPlayerItem.TemplatePlayerItem]: The template player item that initializes this instance.
//
// # Accessing timing information
//
//   - [IAVPlayerItem.CurrentTime]: Returns the current time of the item.
//   - [IAVPlayerItem.CurrentDate]: Returns the current time of the item as a date.
//   - [IAVPlayerItem.Duration]: The duration of the item.
//   - [IAVPlayerItem.Timebase]: The timebase information for the item.
//
// # Determining available time ranges
//
//   - [IAVPlayerItem.LoadedTimeRanges]: An array of time ranges indicating media data that is readily available.
//   - [IAVPlayerItem.SeekableTimeRanges]: An array of time ranges within which it is possible to seek.
//
// # Determining buffering status
//
//   - [IAVPlayerItem.PlaybackLikelyToKeepUp]: A Boolean value that indicates whether the item will likely play through without stalling.
//   - [IAVPlayerItem.PlaybackBufferFull]: A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
//   - [IAVPlayerItem.PlaybackBufferEmpty]: A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
//
// # Configuring expensive network behavior
//
//   - [IAVPlayerItem.PreferredPeakBitRateForExpensiveNetworks]: A limit of network bandwidth consumption by the item when connecting over expensive networks.
//   - [IAVPlayerItem.SetPreferredPeakBitRateForExpensiveNetworks]
//   - [IAVPlayerItem.PreferredMaximumResolutionForExpensiveNetworks]: An upper limit on the resolution of video to download when connecting over expensive networks.
//   - [IAVPlayerItem.SetPreferredMaximumResolutionForExpensiveNetworks]
//
// # Accessing text style rules
//
//   - [IAVPlayerItem.TextStyleRules]: An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
//   - [IAVPlayerItem.SetTextStyleRules]
//
// # Accessing logging information
//
//   - [IAVPlayerItem.AccessLog]: Returns an object that represents a snapshot of the network access log.
//   - [IAVPlayerItem.ErrorLog]: Returns an object that represents a snapshot of the error log.
//
// # Managing time offsets
//
//   - [IAVPlayerItem.AutomaticallyPreservesTimeOffsetFromLive]: A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.
//   - [IAVPlayerItem.SetAutomaticallyPreservesTimeOffsetFromLive]
//   - [IAVPlayerItem.RecommendedTimeOffsetFromLive]: A recommended time offset from the live time based on observed network conditions.
//   - [IAVPlayerItem.ConfiguredTimeOffsetFromLive]: A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.
//   - [IAVPlayerItem.SetConfiguredTimeOffsetFromLive]
//
// # Configuring presentation
//
//   - [IAVPlayerItem.PresentationSize]: The size at which the visual portion of the item is presented by the player.
//   - [IAVPlayerItem.PreferredMaximumResolution]: The desired maximum resolution of a video that is to be downloaded.
//   - [IAVPlayerItem.SetPreferredMaximumResolution]
//   - [IAVPlayerItem.VideoApertureMode]: The video aperture mode to apply during playback.
//   - [IAVPlayerItem.SetVideoApertureMode]
//
// # Configuring HDR settings
//
//   - [IAVPlayerItem.AppliesPerFrameHDRDisplayMetadata]: A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.
//   - [IAVPlayerItem.SetAppliesPerFrameHDRDisplayMetadata]
//
// # Configuring video compositing
//
//   - [IAVPlayerItem.VideoComposition]: The video composition settings to be applied during playback.
//   - [IAVPlayerItem.SetVideoComposition]
//   - [IAVPlayerItem.CustomVideoCompositor]: The custom video compositor.
//   - [IAVPlayerItem.SeekingWaitsForVideoCompositionRendering]: A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.
//   - [IAVPlayerItem.SetSeekingWaitsForVideoCompositionRendering]
//
// # Configuring audio
//
//   - [IAVPlayerItem.AudioMix]: The audio mix parameters to be applied during playback.
//   - [IAVPlayerItem.SetAudioMix]
//   - [IAVPlayerItem.AudioTimePitchAlgorithm]: The processing algorithm used to manage audio pitch for scaled audio edits.
//   - [IAVPlayerItem.SetAudioTimePitchAlgorithm]
//   - [IAVPlayerItem.AllowedAudioSpatializationFormats]: The source audio channel layouts the player item supports for spatialization.
//   - [IAVPlayerItem.SetAllowedAudioSpatializationFormats]
//   - [IAVPlayerItem.AudioSpatializationAllowed]: A Boolean value that indicates whether the player item allows spatialized audio playback.
//   - [IAVPlayerItem.SetAudioSpatializationAllowed]
//
// # Managing player item outputs
//
//   - [IAVPlayerItem.Outputs]: An array of outputs associated with the player item.
//   - [IAVPlayerItem.AddOutput]: Adds the specified player item output object to the receiver.
//   - [IAVPlayerItem.RemoveOutput]: Removes the specified player item output object from the receiver.
//
// # Managing player item data collectors
//
//   - [IAVPlayerItem.MediaDataCollectors]: The collection of associated media data collectors.
//   - [IAVPlayerItem.AddMediaDataCollector]: Adds the specified media data collector to the player item’s collection of media collectors.
//   - [IAVPlayerItem.RemoveMediaDataCollector]: Removes the specified media data collector from the player item’s collection of media collectors.
//
// # Configuring network behavior
//
//   - [IAVPlayerItem.PreferredPeakBitRate]: The desired limit, in bits per second, of network bandwidth consumption for this item.
//   - [IAVPlayerItem.SetPreferredPeakBitRate]
//   - [IAVPlayerItem.PreferredForwardBufferDuration]: The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.
//   - [IAVPlayerItem.SetPreferredForwardBufferDuration]
//   - [IAVPlayerItem.CanUseNetworkResourcesForLiveStreamingWhilePaused]: A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.
//   - [IAVPlayerItem.SetCanUseNetworkResourcesForLiveStreamingWhilePaused]
//
// # Managing playback authorization in macOS
//
//   - [IAVPlayerItem.ContentAuthorizedForPlayback]: A Boolean value that indicates whether the content has been authorized by the user.
//   - [IAVPlayerItem.AuthorizationRequiredForPlayback]: A Boolean value that indicates whether authorization is required to play the content.
//   - [IAVPlayerItem.ApplicationAuthorizedForPlayback]: A Boolean value that indicates whether the application can be used to play the content.
//   - [IAVPlayerItem.RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler]: Presents the user the opportunity to authorize the content for playback.
//   - [IAVPlayerItem.ContentAuthorizationRequestStatus]: The status of the most recent content authorization request.
//   - [IAVPlayerItem.CancelContentAuthorizationRequest]: Cancels the currently outstanding content authorization request.
//
// # Accessing initialization parameters
//
//   - [IAVPlayerItem.Asset]: The asset provided during initialization.
//   - [IAVPlayerItem.AutomaticallyLoadedAssetKeys]: The array of asset keys to be automatically loaded before the player item is ready to play.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem
type IAVPlayerItem interface {
	objectivec.IObject
	AVMetricEventStreamPublisher

	// Topic: Creating a player item

	// Creates a player item with a specified URL.
	InitWithURL(URL foundation.INSURL) AVPlayerItem
	// Creates a player item for a specified asset.
	InitWithAsset(asset IAVAsset) AVPlayerItem
	// Creates a player item with the specified asset and the asset keys to automatically load.
	InitWithAssetAutomaticallyLoadedAssetKeys(asset IAVAsset, automaticallyLoadedAssetKeys []string) AVPlayerItem

	// Topic: Accessing tracks

	// An array of player item track objects.
	Tracks() []AVPlayerItemTrack

	// Topic: Determining readiness

	// The status of the player item.
	Status() AVPlayerItemStatus
	// The error that caused the player item to fail.
	Error() foundation.INSError

	// Topic: Determining playback capabilities

	// A Boolean value that indicates whether the item can play in reverse.
	CanPlayReverse() bool
	// A Boolean value that indicates whether the item can be fast forwarded.
	CanPlayFastForward() bool
	// A Boolean value that indicates whether the item can be quickly reversed.
	CanPlayFastReverse() bool
	// A Boolean value that indicates whether the item can play slower than normal.
	CanPlaySlowForward() bool
	// A Boolean value that indicates whether the item can play slowly backward.
	CanPlaySlowReverse() bool

	// Topic: Setting playback boundaries

	// The time at which forward playback ends.
	ForwardPlaybackEndTime() objectivec.IObject
	SetForwardPlaybackEndTime(value objectivec.IObject)
	// The time at which reverse playback ends.
	ReversePlaybackEndTime() objectivec.IObject
	SetReversePlaybackEndTime(value objectivec.IObject)

	// Topic: Stepping through media

	// A Boolean value that indicates whether the item supports stepping forward.
	CanStepForward() bool
	// A Boolean value that indicates whether the item supports stepping backward.
	CanStepBackward() bool
	// Moves the player item’s current time forward or backward by a specified number of steps.
	StepByCount(stepCount int)

	// Topic: Seeking through media

	// Sets the current playback time to the specified time.
	SeekToTimeCompletionHandler(time objectivec.IObject, completionHandler BoolHandler)
	// Sets the current playback time within a specified time bound and invokes the specified block when the seek operation completes or is interrupted.
	SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objectivec.IObject, toleranceBefore objectivec.IObject, toleranceAfter objectivec.IObject, completionHandler BoolHandler)
	// Sets the current playback time to the time specified by the date object.
	SeekToDateCompletionHandler(date foundation.INSDate, completionHandler BoolHandler) bool
	// Cancels any pending seek requests and invokes the corresponding completion handlers if present.
	CancelPendingSeeks()

	// Topic: Selecting media options

	// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular presentation setting, replacing any previous preference for settings of the same media presentation selector.
	SelectMediaPresentationSettingForMediaSelectionGroup(mediaPresentationSetting IAVMediaPresentationSetting, mediaSelectionGroup IAVMediaSelectionGroup)
	// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of the receiver’s asset with which an associated UI implementation should configure its interface for media selection.
	PreferredCustomMediaSelectionSchemes() []AVCustomMediaSelectionScheme
	SetPreferredCustomMediaSelectionSchemes(value []AVCustomMediaSelectionScheme)
	// Indicates the media presentation settings with media characteristics that are possessed by the currently selected AVMediaSelectionOption in the specified AVMediaSelectionGroup.
	EffectiveMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.INSDictionary
	// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically property is set to YES, configures the player item to prefer a particular language, replacing any previous preference for available languages of the specified group’s custom media selection scheme.
	SelectMediaPresentationLanguageForMediaSelectionGroup(language string, mediaSelectionGroup IAVMediaSelectionGroup)
	// Returns the selected media presentation language for the specified media selection group, if any language has previously been selected via use of -selectMediaPresentationLanguages:forMediaSelectionGroup:.
	SelectedMediaPresentationLanguageForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) string
	// Indicates the media presentation settings that have most recently been selected for each AVMediaPresentationSelector of the AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup.
	SelectedMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.INSDictionary
	// The current media selections for each of the receiver’s media selection groups.
	CurrentMediaSelection() IAVMediaSelection
	// Selects a media option in a given media selection group and deselects all other options in that group.
	SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup)
	// Selects the media option in the specified media selection group that best matches the receiver’s automatic selection criteria.
	SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup)

	// Topic: Setting variant behavior

	// The preferences the player item uses when selecting variant playlists.
	VariantPreferences() AVVariantPreferences
	SetVariantPreferences(value AVVariantPreferences)
	// A Boolean value that indicates whether playback starts with the first eligible variant that appears in the stream’s main playlist.
	StartsOnFirstEligibleVariant() bool
	SetStartsOnFirstEligibleVariant(value bool)

	// Topic: Configuring interstitial events

	// An integrated timeline that represents the player item timing including its scheduled interstitial events.
	IntegratedTimeline() IAVPlayerItemIntegratedTimeline
	// A Boolean value that indicates whether the player item automatically plays interstitial events according to server-side directives.
	AutomaticallyHandlesInterstitialEvents() bool
	SetAutomaticallyHandlesInterstitialEvents(value bool)
	// The template player item that initializes this instance.
	TemplatePlayerItem() IAVPlayerItem

	// Topic: Accessing timing information

	// Returns the current time of the item.
	CurrentTime() objectivec.IObject
	// Returns the current time of the item as a date.
	CurrentDate() foundation.INSDate
	// The duration of the item.
	Duration() objectivec.IObject
	// The timebase information for the item.
	Timebase() objectivec.IObject

	// Topic: Determining available time ranges

	// An array of time ranges indicating media data that is readily available.
	LoadedTimeRanges() []foundation.NSValue
	// An array of time ranges within which it is possible to seek.
	SeekableTimeRanges() []foundation.NSValue

	// Topic: Determining buffering status

	// A Boolean value that indicates whether the item will likely play through without stalling.
	PlaybackLikelyToKeepUp() bool
	// A Boolean value that indicates whether the internal media buffer is full and that further I/O is suspended.
	PlaybackBufferFull() bool
	// A Boolean value that indicates whether playback has consumed all buffered media and that playback will stall or end.
	PlaybackBufferEmpty() bool

	// Topic: Configuring expensive network behavior

	// A limit of network bandwidth consumption by the item when connecting over expensive networks.
	PreferredPeakBitRateForExpensiveNetworks() float64
	SetPreferredPeakBitRateForExpensiveNetworks(value float64)
	// An upper limit on the resolution of video to download when connecting over expensive networks.
	PreferredMaximumResolutionForExpensiveNetworks() corefoundation.CGSize
	SetPreferredMaximumResolutionForExpensiveNetworks(value corefoundation.CGSize)

	// Topic: Accessing text style rules

	// An array of text style rules that specify the formatting and presentation of Web Video Text Tracks (WebVTT) subtitles.
	TextStyleRules() []AVTextStyleRule
	SetTextStyleRules(value []AVTextStyleRule)

	// Topic: Accessing logging information

	// Returns an object that represents a snapshot of the network access log.
	AccessLog() IAVPlayerItemAccessLog
	// Returns an object that represents a snapshot of the error log.
	ErrorLog() IAVPlayerItemErrorLog

	// Topic: Managing time offsets

	// A Boolean value that indicates whether the player preserves its time offset from the live time after a buffering operation.
	AutomaticallyPreservesTimeOffsetFromLive() bool
	SetAutomaticallyPreservesTimeOffsetFromLive(value bool)
	// A recommended time offset from the live time based on observed network conditions.
	RecommendedTimeOffsetFromLive() objectivec.IObject
	// A time value that indicates the offset from the live time to start playback, or resume playback after a seek to positive infinity.
	ConfiguredTimeOffsetFromLive() objectivec.IObject
	SetConfiguredTimeOffsetFromLive(value objectivec.IObject)

	// Topic: Configuring presentation

	// The size at which the visual portion of the item is presented by the player.
	PresentationSize() corefoundation.CGSize
	// The desired maximum resolution of a video that is to be downloaded.
	PreferredMaximumResolution() corefoundation.CGSize
	SetPreferredMaximumResolution(value corefoundation.CGSize)
	// The video aperture mode to apply during playback.
	VideoApertureMode() AVVideoApertureMode
	SetVideoApertureMode(value AVVideoApertureMode)

	// Topic: Configuring HDR settings

	// A Boolean value that indicates whether the player item applies per-frame HDR display metadata during playback.
	AppliesPerFrameHDRDisplayMetadata() bool
	SetAppliesPerFrameHDRDisplayMetadata(value bool)

	// Topic: Configuring video compositing

	// The video composition settings to be applied during playback.
	VideoComposition() IAVVideoComposition
	SetVideoComposition(value IAVVideoComposition)
	// The custom video compositor.
	CustomVideoCompositor() AVVideoCompositing
	// A Boolean value that indicates whether the item’s timing follows the displayed video frame when seeking with a video composition.
	SeekingWaitsForVideoCompositionRendering() bool
	SetSeekingWaitsForVideoCompositionRendering(value bool)

	// Topic: Configuring audio

	// The audio mix parameters to be applied during playback.
	AudioMix() IAVAudioMix
	SetAudioMix(value IAVAudioMix)
	// The processing algorithm used to manage audio pitch for scaled audio edits.
	AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm)
	// The source audio channel layouts the player item supports for spatialization.
	AllowedAudioSpatializationFormats() AVAudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AVAudioSpatializationFormats)
	// A Boolean value that indicates whether the player item allows spatialized audio playback.
	AudioSpatializationAllowed() bool
	SetAudioSpatializationAllowed(value bool)

	// Topic: Managing player item outputs

	// An array of outputs associated with the player item.
	Outputs() []AVPlayerItemOutput
	// Adds the specified player item output object to the receiver.
	AddOutput(output IAVPlayerItemOutput)
	// Removes the specified player item output object from the receiver.
	RemoveOutput(output IAVPlayerItemOutput)

	// Topic: Managing player item data collectors

	// The collection of associated media data collectors.
	MediaDataCollectors() []AVPlayerItemMediaDataCollector
	// Adds the specified media data collector to the player item’s collection of media collectors.
	AddMediaDataCollector(collector IAVPlayerItemMediaDataCollector)
	// Removes the specified media data collector from the player item’s collection of media collectors.
	RemoveMediaDataCollector(collector IAVPlayerItemMediaDataCollector)

	// Topic: Configuring network behavior

	// The desired limit, in bits per second, of network bandwidth consumption for this item.
	PreferredPeakBitRate() float64
	SetPreferredPeakBitRate(value float64)
	// The duration the player should buffer media from the network ahead of the playhead to guard against playback disruption.
	PreferredForwardBufferDuration() float64
	SetPreferredForwardBufferDuration(value float64)
	// A Boolean value that indicates whether the player item can use network resources to keep the playback state up to date while paused.
	CanUseNetworkResourcesForLiveStreamingWhilePaused() bool
	SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool)

	// Topic: Managing playback authorization in macOS

	// A Boolean value that indicates whether the content has been authorized by the user.
	ContentAuthorizedForPlayback() bool
	// A Boolean value that indicates whether authorization is required to play the content.
	AuthorizationRequiredForPlayback() bool
	// A Boolean value that indicates whether the application can be used to play the content.
	ApplicationAuthorizedForPlayback() bool
	// Presents the user the opportunity to authorize the content for playback.
	RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval float64, handler VoidHandler)
	// The status of the most recent content authorization request.
	ContentAuthorizationRequestStatus() AVContentAuthorizationStatus
	// Cancels the currently outstanding content authorization request.
	CancelContentAuthorizationRequest()

	// Topic: Accessing initialization parameters

	// The asset provided during initialization.
	Asset() IAVAsset
	// The array of asset keys to be automatically loaded before the player item is ready to play.
	AutomaticallyLoadedAssetKeys() []string
}

// Init initializes the instance.
func (p AVPlayerItem) Init() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItem) Autorelease() AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItem creates a new AVPlayerItem instance.
func NewAVPlayerItem() AVPlayerItem {
	class := getAVPlayerItemClass()
	rv := objc.Send[AVPlayerItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a player item for a specified asset.
//
// asset: The [AVAsset] to be played.
//
// # Return Value
// 
// A new player item, initialized to play `asset`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:)-87rjl
func NewPlayerItemWithAsset(asset IAVAsset) AVPlayerItem {
	instance := getAVPlayerItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:"), asset)
	return AVPlayerItemFromID(rv)
}

// Creates a player item with the specified asset and the asset keys to
// automatically load.
//
// asset: An instance of [AVAsset].
//
// automaticallyLoadedAssetKeys: An array of strings, each representing a property defined by [AVAsset].
//
// # Return Value
// 
// An initialized instance of [AVPlayerItem].
//
// # Discussion
// 
// The value of each key in `automaticallyLoadedAssetKeys` will automatically
// be loaded by the underlying [AVAsset] before the player item achieves the
// status [PlayerItemStatusReadyToPlay]; i.e. when the item is ready to play,
// the value returned by invoking the [Asset] property’s
// [statusOfValue(forKey:error:)] method will be one of the terminal status
// values, either [KeyValueStatusLoaded], [KeyValueStatusFailed], or
// [KeyValueStatusCancelled].
//
// [statusOfValue(forKey:error:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading/statusOfValue(forKey:error:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:automaticallyLoadedAssetKeys:)-8x4
func NewPlayerItemWithAssetAutomaticallyLoadedAssetKeys(asset IAVAsset, automaticallyLoadedAssetKeys []string) AVPlayerItem {
	instance := getAVPlayerItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAsset:automaticallyLoadedAssetKeys:"), asset, objectivec.StringSliceToNSArray(automaticallyLoadedAssetKeys))
	return AVPlayerItemFromID(rv)
}

// Creates a player item with a specified URL.
//
// URL: A URL identifying the media resource to be played.
//
// # Return Value
// 
// A new player item, prepared to use [URL].
//
// # Discussion
// 
// This method immediately returns the item, but with the status
// [PlayerItemStatusUnknown].
// 
// Associating the player item with an [AVPlayer] immediately begins enqueuing
// its media and preparing it for playback. If the URL contains valid data
// that can be used by the player item, its status later changes to
// [PlayerItemStatusReadyToPlay]. If the URL contains no valid data or
// otherwise can’t be used by the player item, its status later changes to
// [PlayerItemStatusFailed]. You can determine the nature of the failure by
// querying the player item’s [Error] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(url:)
func NewPlayerItemWithURL(URL foundation.INSURL) AVPlayerItem {
	instance := getAVPlayerItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), URL)
	return AVPlayerItemFromID(rv)
}

// Creates a player item with a specified URL.
//
// URL: A URL identifying the media resource to be played.
//
// # Return Value
// 
// A new player item, prepared to use [URL].
//
// # Discussion
// 
// This method immediately returns the item, but with the status
// [PlayerItemStatusUnknown].
// 
// Associating the player item with an [AVPlayer] immediately begins enqueuing
// its media and preparing it for playback. If the URL contains valid data
// that can be used by the player item, its status later changes to
// [PlayerItemStatusReadyToPlay]. If the URL contains no valid data or
// otherwise can’t be used by the player item, its status later changes to
// [PlayerItemStatusFailed]. You can determine the nature of the failure by
// querying the player item’s [Error] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(url:)
func (p AVPlayerItem) InitWithURL(URL foundation.INSURL) AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p.ID, objc.Sel("initWithURL:"), URL)
	return rv
}
// Creates a player item for a specified asset.
//
// asset: The [AVAsset] to be played.
//
// # Return Value
// 
// A new player item, initialized to play `asset`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:)-87rjl
func (p AVPlayerItem) InitWithAsset(asset IAVAsset) AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p.ID, objc.Sel("initWithAsset:"), asset)
	return rv
}
// Creates a player item with the specified asset and the asset keys to
// automatically load.
//
// asset: An instance of [AVAsset].
//
// automaticallyLoadedAssetKeys: An array of strings, each representing a property defined by [AVAsset].
//
// # Return Value
// 
// An initialized instance of [AVPlayerItem].
//
// # Discussion
// 
// The value of each key in `automaticallyLoadedAssetKeys` will automatically
// be loaded by the underlying [AVAsset] before the player item achieves the
// status [PlayerItemStatusReadyToPlay]; i.e. when the item is ready to play,
// the value returned by invoking the [Asset] property’s
// [statusOfValue(forKey:error:)] method will be one of the terminal status
// values, either [KeyValueStatusLoaded], [KeyValueStatusFailed], or
// [KeyValueStatusCancelled].
//
// [statusOfValue(forKey:error:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading/statusOfValue(forKey:error:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/init(asset:automaticallyLoadedAssetKeys:)-8x4
func (p AVPlayerItem) InitWithAssetAutomaticallyLoadedAssetKeys(asset IAVAsset, automaticallyLoadedAssetKeys []string) AVPlayerItem {
	rv := objc.Send[AVPlayerItem](p.ID, objc.Sel("initWithAsset:automaticallyLoadedAssetKeys:"), asset, objectivec.StringSliceToNSArray(automaticallyLoadedAssetKeys))
	return rv
}
// Moves the player item’s current time forward or backward by a specified
// number of steps.
//
// stepCount: The number of steps by which to move.
// 
// A positive number steps forward, a negative number steps backward.
//
// # Discussion
// 
// The size of each step depends on the receiver’s enabled
// [AVPlayerItemTrack] objects (see [Tracks]).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/step(byCount:)
func (p AVPlayerItem) StepByCount(stepCount int) {
	objc.Send[objc.ID](p.ID, objc.Sel("stepByCount:"), stepCount)
}
// Sets the current playback time to the specified time.
//
// time: The time to which to seek.
//
// completionHandler: The block to invoke when the seek operation has either been completed or
// been interrupted. The block takes one argument:
// 
// finished: Indicates whether the seek operation completed.
//
// time is a [coremedia.CMTime].
//
// # Discussion
// 
// Use this method to seek to a specified time in the player item and be
// notified when the operation completes. If the seek request completes
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
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seek(to:completionHandler:)-91gnw
func (p AVPlayerItem) SeekToTimeCompletionHandler(time objectivec.IObject, completionHandler BoolHandler) {
_block1, _cleanup1 := NewBoolBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:completionHandler:"), time, _block1)
}
// Sets the current playback time within a specified time bound and invokes
// the specified block when the seek operation completes or is interrupted.
//
// time: The time to which to seek.
//
// toleranceBefore: The temporal tolerance before `time`.
// 
// Pass [zero] to request sample accurate seeking (this may incur additional
// decoding delay).
// //
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// toleranceAfter: The temporal tolerance after `time`.
// 
// Pass [zero] to request sample accurate seeking (this may incur additional
// decoding delay).
// //
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// completionHandler: The block to invoke when the seek operation has finished.
//
// time is a [coremedia.CMTime].
//
// toleranceBefore is a [coremedia.CMTime].
//
// toleranceAfter is a [coremedia.CMTime].
//
// # Discussion
// 
// Use this method to seek to a specified time for the item.
// 
// The time seeked to will be within the range `[time-toleranceBefore,
// time+toleranceAfter]` and may differ from `time` for efficiency.
// 
// Invoking this method with [positiveInfinity] for `toleranceBefore` and
// `toleranceAfter` is the same as invoking [SeekToTimeCompletionHandler]
// directly.
// 
// Seeking is constrained by the collection of seekable time ranges. If you
// seek to a time outside all of the seekable ranges, the seek will result in
// a current time within the seekable ranges.
//
// [positiveInfinity]: https://developer.apple.com/documentation/CoreMedia/CMTime/positiveInfinity
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seek(to:toleranceBefore:toleranceAfter:completionHandler:)
func (p AVPlayerItem) SeekToTimeToleranceBeforeToleranceAfterCompletionHandler(time objectivec.IObject, toleranceBefore objectivec.IObject, toleranceAfter objectivec.IObject, completionHandler BoolHandler) {
_block3, _cleanup3 := NewBoolBlock(completionHandler)
	defer _cleanup3()
	objc.Send[objc.ID](p.ID, objc.Sel("seekToTime:toleranceBefore:toleranceAfter:completionHandler:"), time, toleranceBefore, toleranceAfter, _block3)
}
// Sets the current playback time to the time specified by the date object.
//
// date: The time to which to seek.
//
// completionHandler: The block to invoke when the seek operation has either been completed or
// been interrupted. The block takes one argument:
// 
// finished: Indicates whether the seek operation completed.
//
// # Return Value
// 
// [true] if the playhead moved to the specified date or [false] if it did
// not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to seek to a specified time in the player item and be
// notified when the operation completes. If the seek request completes
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
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seek(to:completionHandler:)-1dibq
func (p AVPlayerItem) SeekToDateCompletionHandler(date foundation.INSDate, completionHandler BoolHandler) bool {
_block1, _cleanup1 := NewBoolBlock(completionHandler)
	defer _cleanup1()
	rv := objc.Send[bool](p.ID, objc.Sel("seekToDate:completionHandler:"), date, _block1)
	return rv
}
// Cancels any pending seek requests and invokes the corresponding completion
// handlers if present.
//
// # Discussion
// 
// Use this method to cancel and release the completion handlers of pending
// seeks.
// 
// The `finished` parameter of the completion handlers will be set to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelPendingSeeks()
func (p AVPlayerItem) CancelPendingSeeks() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelPendingSeeks"))
}
// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically
// property is set to YES, configures the player item to prefer a particular
// presentation setting, replacing any previous preference for settings of the
// same media presentation selector.
//
// mediaPresentationSetting: The setting to select.
//
// mediaSelectionGroup: The media selection group, obtained from the receiver’s asset, to which
// the specified setting is to be applied.
//
// # Discussion
// 
// Note that preferences for media characteristics indicated by selected
// AVMediaPresentationSettings are treated as supplemental to the associated
// AVPlayer’s media selection criteria for the AVMediaSelectionGroup. An
// AVPlayer’s default media selection criteria can also indicate preferences
// for media characteristics, such as those indicating the availability of
// accessibility affordances such as audio descriptions, and these media
// characteristics can be left up to the AVPlayer to manage even when an
// AVCustomMediaSelectionScheme is in use. But if you wish to do so, you can
// use AVMediaPresentationSettings offered by a AVCustomMediaSelectionScheme
// in combination with custom AVPlayerMediaSelectionCriteria. If the specified
// setting isn’t offered by an AVMediaPresentationSelector of the
// AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup, no
// change in the presentation of the media will result. This method has no
// effect when the associated AVPlayer’s
// appliesMediaSelectionCriteriaAutomatically property has a value of NO, in
// which case you must use -selectMediaOption:inMediaSelectionGroup: instead
// in order to alter the presentation state of the media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/select(_:for:)
func (p AVPlayerItem) SelectMediaPresentationSettingForMediaSelectionGroup(mediaPresentationSetting IAVMediaPresentationSetting, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectMediaPresentationSetting:forMediaSelectionGroup:"), mediaPresentationSetting, mediaSelectionGroup)
}
// Indicates the media presentation settings with media characteristics that
// are possessed by the currently selected AVMediaSelectionOption in the
// specified AVMediaSelectionGroup.
//
// mediaSelectionGroup: An AVMediaSelectionGroup obtained from the receiver’s asset for which the
// currently effective media presentation settings are desired.
//
// # Return Value
// 
// A dictionary with AVMediaPresentationSelectors as keys and
// AVMediaPresentationSettings as values, unless the AVMediaSelectionOption
// currently selected in the group possesses none of the characteristics
// associated with the selector’s settings. In that case the dictionary
// value will be NSNull.
//
// # Discussion
// 
// Effective media presentation settings can differ from the currently
// effective media presentation settings if no AVMediaSelectionOption of the
// specified AVMediaSelectionGroup with the currently selected media
// presentation language possesses all of the characteristics associated with
// the currently selected settings. A value of NSNull for an
// AVMediaPresentationSelector can occur if either the content is
// inappropriately authored for the use of the AVCustomMediaSelectionScheme or
// if the currently selected AVMediaSelectionOption has been selected by means
// other than through the use of AVMediaPresentationSettings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/effectiveMediaPresentationSettings(for:)
func (p AVPlayerItem) EffectiveMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("effectiveMediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return foundation.NSDictionaryFromID(rv)
}
// When the associated AVPlayer’s appliesMediaSelectionCriteriaAutomatically
// property is set to YES, configures the player item to prefer a particular
// language, replacing any previous preference for available languages of the
// specified group’s custom media selection scheme.
//
// mediaSelectionGroup: The media selection group, obtained from the receiver’s asset, to which
// the specified setting is to be applied.
//
// # Discussion
// 
// Overrides preferences for languages specified by the AVPlayer’s current
// media selection criteria. This method has no effect when the associated
// AVPlayer’s appliesMediaSelectionCriteriaAutomatically property has a
// value of NO, in which case you must use
// -selectMediaOption:inMediaSelectionGroup: instead in order to alter the
// presentation state of the media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectMediaPresentationLanguage(_:for:)
func (p AVPlayerItem) SelectMediaPresentationLanguageForMediaSelectionGroup(language string, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectMediaPresentationLanguage:forMediaSelectionGroup:"), objc.String(language), mediaSelectionGroup)
}
// Returns the selected media presentation language for the specified media
// selection group, if any language has previously been selected via use of
// -selectMediaPresentationLanguages:forMediaSelectionGroup:.
//
// mediaSelectionGroup: The media selection group, obtained from the receiver’s asset, for which
// the selected media presentation language is requested.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedMediaPresentationLanguage(for:)
func (p AVPlayerItem) SelectedMediaPresentationLanguageForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectedMediaPresentationLanguageForMediaSelectionGroup:"), mediaSelectionGroup)
	return foundation.NSStringFromID(rv).String()
}
// Indicates the media presentation settings that have most recently been
// selected for each AVMediaPresentationSelector of the
// AVCustomMediaSelectionScheme of the specified AVMediaSelectionGroup.
//
// mediaSelectionGroup: An AVMediaSelectionGroup obtained from the receiver’s asset for which the
// currently selected media presentation settings are desired.
//
// # Return Value
// 
// A dictionary with AVMediaPresentationSelectors as keys and
// AVMediaPresentationSettings as values, providing the most recently selected
// setting for each selector or, if no setting has previously been selected,
// NSNull.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedMediaPresentationSettings(for:)
func (p AVPlayerItem) SelectedMediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectedMediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return foundation.NSDictionaryFromID(rv)
}
// Selects a media option in a given media selection group and deselects all
// other options in that group.
//
// mediaSelectionOption: The option to select.
// 
// If the value of the [AllowsEmptySelection] property of
// `mediaSelectionGroup` is [true], you can pass `nil` to deselect all media
// selection options in the group.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// mediaSelectionGroup: The media selection group, obtained from the receiver’s asset, that
// contains `mediaSelectionOption`.
//
// # Discussion
// 
// If `mediaSelectionOption` isn’t a member of the `mediaSelectionGroup`, no
// change in presentation state will result.
// 
// If multiple options within a group meet your criteria for selection
// according to locale or other considerations, and if these options are
// otherwise indistinguishable to you according to media characteristics that
// are meaningful for your application, content is typically authored so that
// the first available option that meets your criteria is appropriate for
// selection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/select(_:in:)
func (p AVPlayerItem) SelectMediaOptionInMediaSelectionGroup(mediaSelectionOption IAVMediaSelectionOption, mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectMediaOption:inMediaSelectionGroup:"), mediaSelectionOption, mediaSelectionGroup)
}
// Selects the media option in the specified media selection group that best
// matches the receiver’s automatic selection criteria.
//
// mediaSelectionGroup: The media selection group, obtained from the receiver’s [Asset], that
// contains the specified option.
//
// # Discussion
// 
// This method has no effect unless the
// [AppliesMediaSelectionCriteriaAutomatically] property of the associated
// [AVPlayer] is [true] and unless automatic media selection has previously
// been overridden by invoking [SelectMediaOptionInMediaSelectionGroup].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectMediaOptionAutomatically(in:)
func (p AVPlayerItem) SelectMediaOptionAutomaticallyInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectMediaOptionAutomaticallyInMediaSelectionGroup:"), mediaSelectionGroup)
}
// Returns the current time of the item.
//
// # Return Value
// 
// The current time of the item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentTime()
func (p AVPlayerItem) CurrentTime() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentTime"))
	return objectivec.Object{ID: rv}
}
// Returns the current time of the item as a date.
//
// # Return Value
// 
// The current time of the item as a date, or `nil` if there isn’t a mapped
// date for the item.
//
// # Discussion
// 
// The system calculates this value from the `EXT-X-PROGRAM-DATE-TIME` tag.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentDate()
func (p AVPlayerItem) CurrentDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentDate"))
	return foundation.NSDateFromID(rv)
}
// Returns an object that represents a snapshot of the network access log.
//
// # Return Value
// 
// An object that represents a snapshot of the network access log. The
// returned value can be `nil`.
//
// # Discussion
// 
// If the method returns `nil`, there is no logging information currently
// available for the player item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/accessLog()
func (p AVPlayerItem) AccessLog() IAVPlayerItemAccessLog {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessLog"))
	return AVPlayerItemAccessLogFromID(rv)
}
// Returns an object that represents a snapshot of the error log.
//
// # Return Value
// 
// An object that represents a snapshot of the error log. The returned value
// can be `nil`.
//
// # Discussion
// 
// If the method returns `nil`, there is no logging information currently
// available for the player item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/errorLog()
func (p AVPlayerItem) ErrorLog() IAVPlayerItemErrorLog {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("errorLog"))
	return AVPlayerItemErrorLogFromID(rv)
}
// Adds the specified player item output object to the receiver.
//
// output: The player item output object to associate with the item.
//
// # Discussion
// 
// When you add an [AVPlayerItemOutput] object to an item, the samples
// associated with that output object are processed according to the rules for
// mixing, composing, or excluding content that the [AVPlayer] object honors
// for the specific media type. For example, video media is composed according
// to the instructions provided by the player item’s video composition
// object and audio media is mixed according to the parameters of its audio
// mix object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/add(_:)-16ctk
func (p AVPlayerItem) AddOutput(output IAVPlayerItemOutput) {
	objc.Send[objc.ID](p.ID, objc.Sel("addOutput:"), output)
}
// Removes the specified player item output object from the receiver.
//
// output: The player item output object to remove.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/remove(_:)-46b1r
func (p AVPlayerItem) RemoveOutput(output IAVPlayerItemOutput) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeOutput:"), output)
}
// Adds the specified media data collector to the player item’s collection
// of media collectors.
//
// collector: An instance of [AVPlayerItemMediaDataCollector].
//
// # Discussion
// 
// This method may incur additional I/O to collect the requested media data
// asynchronously.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/add(_:)-9l3to
func (p AVPlayerItem) AddMediaDataCollector(collector IAVPlayerItemMediaDataCollector) {
	objc.Send[objc.ID](p.ID, objc.Sel("addMediaDataCollector:"), collector)
}
// Removes the specified media data collector from the player item’s
// collection of media collectors.
//
// collector: The instance of [AVPlayerItemMediaDataCollector] to remove.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/remove(_:)-29iuz
func (p AVPlayerItem) RemoveMediaDataCollector(collector IAVPlayerItemMediaDataCollector) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeMediaDataCollector:"), collector)
}
// Presents the user the opportunity to authorize the content for playback.
//
// timeoutInterval: The maximum amount of time in seconds to wait for the user to authorize the
// content before calling the handler block with a timeout result.
//
// handler: The block to be called upon completion.
//
// # Discussion
// 
// Calling this method will present the user with the opportunity to authorize
// the content (for example, by launching iTunes and prompting the user to
// enter their Apple ID and password).
// 
// When the user has taken action (or the timeout has elapsed), the completion
// handler is invoked. You determine the status of the authorization attempt
// by checking the value of the [ContentAuthorizationRequestStatus] property.
// 
// Even if the status indicates a completed authorization, the content may
// still not be authorized (for example, if the user authorizes an Apple ID
// other than that associated with the content). You should re-check the value
// of [ContentAuthorizationRequestStatus] to verify whether the content has
// actually been authorized before continuing. It is not necessary to call
// this method if the value of [ContentAuthorizationRequestStatus] is already
// true.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/requestContentAuthorizationAsynchronously(withTimeoutInterval:completionHandler:)
func (p AVPlayerItem) RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval float64, handler VoidHandler) {
_block1, _cleanup1 := NewVoidBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](p.ID, objc.Sel("requestContentAuthorizationAsynchronouslyWithTimeoutInterval:completionHandler:"), timeoutInterval, _block1)
}
// Cancels the currently outstanding content authorization request.
//
// # Discussion
// 
// Calling this method while a content authorization request is pending will
// cause that request to be cancelled and its completion handler to be invoked
// with a status of [ContentAuthorizationCancelled].
// 
// This method does not block.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelContentAuthorizationRequest()
func (p AVPlayerItem) CancelContentAuthorizationRequest() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelContentAuthorizationRequest"))
}

// Returns a new player item for a specified asset.
//
// asset: The [AVAsset] to be played.
//
// # Return Value
// 
// A new player item, initialized to play `asset`.
//
// # Discussion
// 
// This method is equivalent to invoking
// [PlayerItemWithAssetAutomaticallyLoadedAssetKeys], passing `["duration"]`
// as the value of `automaticallyLoadedAssetKeys`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playerItemWithAsset:
func (_AVPlayerItemClass AVPlayerItemClass) PlayerItemWithAsset(asset IAVAsset) AVPlayerItem {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerItemClass.class), objc.Sel("playerItemWithAsset:"), asset)
	return AVPlayerItemFromID(rv)
}
// Creates a player item with the specified asset and the asset keys to
// automatically load.
//
// asset: An instance of [AVAsset].
//
// automaticallyLoadedAssetKeys: An array of strings, each representing a property key defined by [AVAsset].
//
// # Return Value
// 
// An initialized instance of [AVPlayerItem].
//
// # Discussion
// 
// The value of each key in `automaticallyLoadedAssetKeys` will automatically
// be loaded by the underlying [AVAsset] before the player item achieves the
// status [PlayerItemStatusReadyToPlay]; i.e. when the item is ready to play,
// the value returned by invoking the [Asset] property’s
// [statusOfValue(forKey:error:)] method will be one of the terminal status
// values, either [KeyValueStatusLoaded], [KeyValueStatusFailed], or
// [KeyValueStatusCancelled].
//
// [statusOfValue(forKey:error:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading/statusOfValue(forKey:error:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playerItemWithAsset:automaticallyLoadedAssetKeys:
func (_AVPlayerItemClass AVPlayerItemClass) PlayerItemWithAssetAutomaticallyLoadedAssetKeys(asset IAVAsset, automaticallyLoadedAssetKeys []string) AVPlayerItem {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerItemClass.class), objc.Sel("playerItemWithAsset:automaticallyLoadedAssetKeys:"), asset, objectivec.StringSliceToNSArray(automaticallyLoadedAssetKeys))
	return AVPlayerItemFromID(rv)
}
// Returns a new player item with a specified URL.
//
// URL: A URL identifying the media resource to be played.
//
// # Return Value
// 
// A new player item, prepared to use [URL].
//
// # Discussion
// 
// This method immediately returns the item, but with the status
// [PlayerItemStatusUnknown].
// 
// Associating the player item with an [AVPlayer] immediately begins enqueuing
// its media and preparing it for playback. If the URL contains valid data
// that can be used by the player item, its status later changes to
// [PlayerItemStatusReadyToPlay]. If the URL contains no valid data or
// otherwise can’t be used by the player item, its status later changes to
// [PlayerItemStatusFailed]. You can determine the nature of the failure by
// querying the player item’s [Error] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playerItemWithURL:
func (_AVPlayerItemClass AVPlayerItemClass) PlayerItemWithURL(URL foundation.INSURL) AVPlayerItem {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerItemClass.class), objc.Sel("playerItemWithURL:"), URL)
	return AVPlayerItemFromID(rv)
}

// An array of player item track objects.
//
// # Discussion
// 
// The value is an empty array before the player loads the underlying tracks.
// Key-value observe this property value to access valid tracks as soon as
// they’re available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/tracks
func (p AVPlayerItem) Tracks() []AVPlayerItemTrack {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("tracks"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItemTrack {
		return AVPlayerItemTrackFromID(id)
	})
}
// The status of the player item.
//
// # Discussion
// 
// When a player item is created, its [Status] is [PlayerItemStatusUnknown],
// meaning its media hasn’t been loaded and has not yet been enqueued for
// playback. Associating a player item with an [AVPlayer] immediately begins
// enqueuing the item’s media and preparing it for playback. When the player
// item’s media has been loaded and is ready for use, its status will change
// to [PlayerItemStatusReadyToPlay]. You can observe this change using
// key-value observing.
// 
// For possible values, see [AVPlayerItem.Status].
//
// [AVPlayerItem.Status]: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/status-swift.property
func (p AVPlayerItem) Status() AVPlayerItemStatus {
	rv := objc.Send[AVPlayerItemStatus](p.ID, objc.Sel("status"))
	return AVPlayerItemStatus(rv)
}
// The error that caused the player item to fail.
//
// # Discussion
// 
// The value of this property is an error that describes what caused the
// player item to no longer be able to be played.
// 
// If the receiver’s status is not [PlayerItemStatusFailed], the value of
// this property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/error
func (p AVPlayerItem) Error() foundation.INSError {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the item can play in reverse.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlayReverse
func (p AVPlayerItem) CanPlayReverse() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canPlayReverse"))
	return rv
}
// A Boolean value that indicates whether the item can be fast forwarded.
//
// # Discussion
// 
// An item can be fast forwarded if its rate can be greater than `1.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlayFastForward
func (p AVPlayerItem) CanPlayFastForward() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canPlayFastForward"))
	return rv
}
// A Boolean value that indicates whether the item can be quickly reversed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlayFastReverse
func (p AVPlayerItem) CanPlayFastReverse() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canPlayFastReverse"))
	return rv
}
// A Boolean value that indicates whether the item can play slower than
// normal.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlaySlowForward
func (p AVPlayerItem) CanPlaySlowForward() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canPlaySlowForward"))
	return rv
}
// A Boolean value that indicates whether the item can play slowly backward.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canPlaySlowReverse
func (p AVPlayerItem) CanPlaySlowReverse() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canPlaySlowReverse"))
	return rv
}
// The time at which forward playback ends.
//
// # Discussion
// 
// The value indicates the time at which playback should end when the playback
// rate is positive (see [AVPlayer]’s [Rate] property).
// 
// The default value is [invalid], which indicates that no end time for
// forward playback is specified. In this case, the effective end time for
// forward playback is the item’s duration.
// 
// The value of this property has no effect on playback when the rate is
// negative.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/forwardPlaybackEndTime
func (p AVPlayerItem) ForwardPlaybackEndTime() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("forwardPlaybackEndTime"))
	return objectivec.Object{ID: rv}
}
func (p AVPlayerItem) SetForwardPlaybackEndTime(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setForwardPlaybackEndTime:"), value)
}
// The time at which reverse playback ends.
//
// # Discussion
// 
// The value indicated the time at which playback should end when the playback
// rate is negative (see [AVPlayer]’s [Rate] property).
// 
// The default value is [invalid], which indicates that no end time for
// reverse playback is specified. In this case, the effective end time for
// reverse playback is [zero].
// 
// The value of this property has no effect on playback when the rate is
// positive.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/reversePlaybackEndTime
func (p AVPlayerItem) ReversePlaybackEndTime() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("reversePlaybackEndTime"))
	return objectivec.Object{ID: rv}
}
func (p AVPlayerItem) SetReversePlaybackEndTime(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setReversePlaybackEndTime:"), value)
}
// A Boolean value that indicates whether the item supports stepping forward.
//
// # Discussion
// 
// Once the item becomes ready to play, the value of this property does not
// change. This behavior applies even when boundary conditions, such as when
// the item’s current time is equal to its end time, have been reached.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canStepForward
func (p AVPlayerItem) CanStepForward() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canStepForward"))
	return rv
}
// A Boolean value that indicates whether the item supports stepping backward.
//
// # Discussion
// 
// Once the item becomes ready to play, the value of this property does not
// change. This behavior applies even when boundary conditions, such as when
// the item’s current time is [zero], have been reached.
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canStepBackward
func (p AVPlayerItem) CanStepBackward() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canStepBackward"))
	return rv
}
// Indicates the AVCustomMediaSelectionSchemes of AVMediaSelectionGroups of
// the receiver’s asset with which an associated UI implementation should
// configure its interface for media selection.
//
// # Discussion
// 
// Recommended usage: if use of a custom media selection scheme is desired,
// set this property before either replacing an AVPlayer’s current item with
// the receiver or adding the receiver to an AVQueuePlayer’s play queue.
// This will satisfy requirements of UI implementations that commit to a
// configuration of UI elements as the receiver becomes ready to play.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredCustomMediaSelectionSchemes
func (p AVPlayerItem) PreferredCustomMediaSelectionSchemes() []AVCustomMediaSelectionScheme {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("preferredCustomMediaSelectionSchemes"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCustomMediaSelectionScheme {
		return AVCustomMediaSelectionSchemeFromID(id)
	})
}
func (p AVPlayerItem) SetPreferredCustomMediaSelectionSchemes(value []AVCustomMediaSelectionScheme) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredCustomMediaSelectionSchemes:"), objectivec.IObjectSliceToNSArray(value))
}
// The current media selections for each of the receiver’s media selection
// groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/currentMediaSelection
func (p AVPlayerItem) CurrentMediaSelection() IAVMediaSelection {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentMediaSelection"))
	return AVMediaSelectionFromID(objc.ID(rv))
}
// The preferences the player item uses when selecting variant playlists.
//
// # Discussion
// 
// The default value is [VariantPreferenceNone].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/variantPreferences
func (p AVPlayerItem) VariantPreferences() AVVariantPreferences {
	rv := objc.Send[AVVariantPreferences](p.ID, objc.Sel("variantPreferences"))
	return AVVariantPreferences(rv)
}
func (p AVPlayerItem) SetVariantPreferences(value AVVariantPreferences) {
	objc.Send[struct{}](p.ID, objc.Sel("setVariantPreferences:"), value)
}
// A Boolean value that indicates whether playback starts with the first
// eligible variant that appears in the stream’s main playlist.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/startsOnFirstEligibleVariant
func (p AVPlayerItem) StartsOnFirstEligibleVariant() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("startsOnFirstEligibleVariant"))
	return rv
}
func (p AVPlayerItem) SetStartsOnFirstEligibleVariant(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setStartsOnFirstEligibleVariant:"), value)
}
// An integrated timeline that represents the player item timing including its
// scheduled interstitial events.
//
// # Discussion
// 
// The value is `nil` for player items in an interstitial player.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/integratedTimeline
func (p AVPlayerItem) IntegratedTimeline() IAVPlayerItemIntegratedTimeline {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("integratedTimeline"))
	return AVPlayerItemIntegratedTimelineFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the player item automatically plays
// interstitial events according to server-side directives.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyHandlesInterstitialEvents
func (p AVPlayerItem) AutomaticallyHandlesInterstitialEvents() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("automaticallyHandlesInterstitialEvents"))
	return rv
}
func (p AVPlayerItem) SetAutomaticallyHandlesInterstitialEvents(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutomaticallyHandlesInterstitialEvents:"), value)
}
// The template player item that initializes this instance.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/template
func (p AVPlayerItem) TemplatePlayerItem() IAVPlayerItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("templatePlayerItem"))
	return AVPlayerItemFromID(objc.ID(rv))
}
// The duration of the item.
//
// # Discussion
// 
// This property indicates the duration of the item, not considering either
// its [ForwardPlaybackEndTime] or [ReversePlaybackEndTime].
// 
// The system reports the value of this property as [indefinite] until it
// loads the duration of the underlying asset. There are two ways to make sure
// you don’t access the value of duration until the system makes it
// available:
// 
// - Wait until the [Status] of the player item is
// [PlayerItemStatusReadyToPlay]. - Register for key-value observation of the
// property and request the initial value. If the system reports the initial
// value as [indefinite], wait for the player item to notify you when
// [Duration] becomes available.
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/duration
func (p AVPlayerItem) Duration() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("duration"))
	return objectivec.Object{ID: rv}
}
// The timebase information for the item.
//
// # Discussion
// 
// The system uses timebase information to synchronize playback of the current
// item with the host clock. You can use this property to access the timebase
// information, but you can’t use it to set the time or the rate of
// playback.
// 
// If you need to respond to changes in the effective playback rate, listen
// for [kCMTimebaseNotification_EffectiveRateChanged] notifications that the
// player item’s [Timebase] posts. These notifications announce when the
// effective playback rate changes, which includes any compensation necessary
// for drifting behaviors of audio output hardware.
//
// [kCMTimebaseNotification_EffectiveRateChanged]: https://developer.apple.com/documentation/CoreMedia/kCMTimebaseNotification_EffectiveRateChanged
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/timebase
func (p AVPlayerItem) Timebase() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("timebase"))
	return objectivec.Object{ID: rv}
}
// An array of time ranges indicating media data that is readily available.
//
// # Discussion
// 
// The array contains [NSValue] objects containing a [CMTimeRange] value
// indicating the times ranges for which the player item has media data
// readily available. The time ranges returned may be discontinuous.
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/loadedTimeRanges
func (p AVPlayerItem) LoadedTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("loadedTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// An array of time ranges within which it is possible to seek.
//
// # Discussion
// 
// The array contains [NSValue] objects containing a [CMTimeRange] value
// indicating the times ranges to which the player item can seek. The time
// ranges returned may be discontinuous.
//
// [CMTimeRange]: https://developer.apple.com/documentation/CoreMedia/CMTimeRange
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seekableTimeRanges
func (p AVPlayerItem) SeekableTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("seekableTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// A Boolean value that indicates whether the item will likely play through
// without stalling.
//
// # Discussion
// 
// This property communicates a prediction of playability. Factors considered
// in this prediction include I/O throughput and media decode performance. It
// is possible for `playbackLikelyToKeepUp` to indicate [false] while the
// property [PlaybackBufferFull] indicates [true]. In this event the playback
// buffer has reached capacity but there isn’t the statistical data to
// support a prediction that playback is likely to keep up in the future. It
// is up to you to decide whether to continue media playback.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isPlaybackLikelyToKeepUp
func (p AVPlayerItem) PlaybackLikelyToKeepUp() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPlaybackLikelyToKeepUp"))
	return rv
}
// A Boolean value that indicates whether the internal media buffer is full
// and that further I/O is suspended.
//
// # Discussion
// 
// Despite the playback buffer reaching capacity there might not exist
// sufficient statistical data to support a [PlaybackLikelyToKeepUp]
// prediction of [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isPlaybackBufferFull
func (p AVPlayerItem) PlaybackBufferFull() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPlaybackBufferFull"))
	return rv
}
// A Boolean value that indicates whether playback has consumed all buffered
// media and that playback will stall or end.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isPlaybackBufferEmpty
func (p AVPlayerItem) PlaybackBufferEmpty() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isPlaybackBufferEmpty"))
	return rv
}
// A limit of network bandwidth consumption by the item when connecting over
// expensive networks.
//
// # Discussion
// 
// When this value is nonzero, the player attempts to limit item playback by
// the specified bit rate when streaming over an expensive network, such as a
// cellular data plan. If the system can’t reduce the bit rate to meet this
// value, it reduces it as much as possible while it continues to play the
// item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRateForExpensiveNetworks
func (p AVPlayerItem) PreferredPeakBitRateForExpensiveNetworks() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("preferredPeakBitRateForExpensiveNetworks"))
	return rv
}
func (p AVPlayerItem) SetPreferredPeakBitRateForExpensiveNetworks(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredPeakBitRateForExpensiveNetworks:"), value)
}
// An upper limit on the resolution of video to download when connecting over
// expensive networks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredMaximumResolutionForExpensiveNetworks
func (p AVPlayerItem) PreferredMaximumResolutionForExpensiveNetworks() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("preferredMaximumResolutionForExpensiveNetworks"))
	return corefoundation.CGSize(rv)
}
func (p AVPlayerItem) SetPreferredMaximumResolutionForExpensiveNetworks(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredMaximumResolutionForExpensiveNetworks:"), value)
}
// An array of text style rules that specify the formatting and presentation
// of Web Video Text Tracks (WebVTT) subtitles.
//
// # Discussion
// 
// Text style rules apply only to WebVTT subtitles. They don’t apply to
// other subtitle formats and legible text.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/textStyleRules
func (p AVPlayerItem) TextStyleRules() []AVTextStyleRule {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("textStyleRules"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVTextStyleRule {
		return AVTextStyleRuleFromID(id)
	})
}
func (p AVPlayerItem) SetTextStyleRules(value []AVTextStyleRule) {
	objc.Send[struct{}](p.ID, objc.Sel("setTextStyleRules:"), objectivec.IObjectSliceToNSArray(value))
}
// A Boolean value that indicates whether the player preserves its time offset
// from the live time after a buffering operation.
//
// # Discussion
// 
// The default value of this property is [false]. If the value is [true], the
// player seeks forward after it finishes buffering to restore the position
// that the playhead had when buffering began, relative to the end of the
// player item’s [SeekableTimeRanges] property value.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyPreservesTimeOffsetFromLive
func (p AVPlayerItem) AutomaticallyPreservesTimeOffsetFromLive() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("automaticallyPreservesTimeOffsetFromLive"))
	return rv
}
func (p AVPlayerItem) SetAutomaticallyPreservesTimeOffsetFromLive(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutomaticallyPreservesTimeOffsetFromLive:"), value)
}
// A recommended time offset from the live time based on observed network
// conditions.
//
// # Discussion
// 
// For nonlive stream content, the value is [invalid].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/recommendedTimeOffsetFromLive
func (p AVPlayerItem) RecommendedTimeOffsetFromLive() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("recommendedTimeOffsetFromLive"))
	return objectivec.Object{ID: rv}
}
// A time value that indicates the offset from the live time to start
// playback, or resume playback after a seek to positive infinity.
//
// # Discussion
// 
// For nonlive stream content, the value is [invalid].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/configuredTimeOffsetFromLive
func (p AVPlayerItem) ConfiguredTimeOffsetFromLive() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("configuredTimeOffsetFromLive"))
	return objectivec.Object{ID: rv}
}
func (p AVPlayerItem) SetConfiguredTimeOffsetFromLive(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setConfiguredTimeOffsetFromLive:"), value)
}
// The size at which the visual portion of the item is presented by the
// player.
//
// # Discussion
// 
// This property can be accessed at any time, but may return a value of
// [CGSizeZero] prior to the player item becoming ready to play. You can use
// key-value observing to obtain the player item’s valid presentation size
// as early as possible.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/presentationSize
func (p AVPlayerItem) PresentationSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("presentationSize"))
	return corefoundation.CGSize(rv)
}
// The desired maximum resolution of a video that is to be downloaded.
//
// # Discussion
// 
// Defaults to [CGSizeZero], which indicates there is no limit on the video
// resolution. Any other value indicates a preferred maximum video resolution.
// This property only applies to HTTP Live Streaming assets.
//
// [CGSizeZero]: https://developer.apple.com/documentation/CoreGraphics/CGSizeZero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredMaximumResolution
func (p AVPlayerItem) PreferredMaximumResolution() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("preferredMaximumResolution"))
	return corefoundation.CGSize(rv)
}
func (p AVPlayerItem) SetPreferredMaximumResolution(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredMaximumResolution:"), value)
}
// The video aperture mode to apply during playback.
//
// # Discussion
// 
// The default value for this property is [cleanAperture].
//
// [cleanAperture]: https://developer.apple.com/documentation/AVFoundation/AVVideoApertureMode/cleanAperture
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/videoApertureMode
func (p AVPlayerItem) VideoApertureMode() AVVideoApertureMode {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("videoApertureMode"))
	return AVVideoApertureMode(foundation.NSStringFromID(rv).String())
}
func (p AVPlayerItem) SetVideoApertureMode(value AVVideoApertureMode) {
	objc.Send[struct{}](p.ID, objc.Sel("setVideoApertureMode:"), objc.String(string(value)))
}
// A Boolean value that indicates whether the player item applies per-frame
// HDR display metadata during playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/appliesPerFrameHDRDisplayMetadata
func (p AVPlayerItem) AppliesPerFrameHDRDisplayMetadata() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("appliesPerFrameHDRDisplayMetadata"))
	return rv
}
func (p AVPlayerItem) SetAppliesPerFrameHDRDisplayMetadata(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAppliesPerFrameHDRDisplayMetadata:"), value)
}
// The video composition settings to be applied during playback.
//
// # Discussion
// 
// A video composition can only be used with file-based media and is not
// supported for use with media served using HTTP Live Streaming.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/videoComposition
func (p AVPlayerItem) VideoComposition() IAVVideoComposition {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("videoComposition"))
	return AVVideoCompositionFromID(objc.ID(rv))
}
func (p AVPlayerItem) SetVideoComposition(value IAVVideoComposition) {
	objc.Send[struct{}](p.ID, objc.Sel("setVideoComposition:"), value)
}
// The custom video compositor.
//
// # Discussion
// 
// The custom video compositor instance that is used during image generation
// is accessible via this property after the value of [VideoComposition] is
// set to an [AVVideoComposition] instance that specifies a custom video
// compositor class. Any additional communication between the application and
// that instance of the custom video compositor, if any is required for
// configuration or other purposes, can only occur once that has happened.
// 
// If the value of [VideoComposition] is changed from an [AVVideoComposition]
// that specifies a custom video compositor class to another instance of
// [AVVideoComposition] that specifies the same custom video compositor class,
// the instance of the custom video compositor that was previously created
// will receive the [RenderContextChanged] message and remain in use for
// subsequent image generation.
// 
// This property is `nil` if there is no video compositor, or if the internal
// video compositor is in use.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/customVideoCompositor
func (p AVPlayerItem) CustomVideoCompositor() AVVideoCompositing {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("customVideoCompositor"))
	return AVVideoCompositingObjectFromID(rv)
}
// A Boolean value that indicates whether the item’s timing follows the
// displayed video frame when seeking with a video composition.
//
// # Discussion
// 
// By default, item timing is updated as quickly as possible during seeking.
// Specifically, the item does not wait for new frames to be rendered when
// seeking during normal playback. In most situations, the latency between the
// completion of a seek operation and the display of a video frame at the new
// time is negligible. However, when video compositions are in use, the
// processing of video may introduce noticeable latency. Setting the value of
// this property to [true] causes the item’s timing to be updated only after
// the corresponding video frame has been displayed. For example, this allows
// an AVSynchronizedLayer object associated with the item to remain in sync
// with the displayed video.
// 
// This property has no effect on items whose [VideoComposition] property is
// `nil`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/seekingWaitsForVideoCompositionRendering
func (p AVPlayerItem) SeekingWaitsForVideoCompositionRendering() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("seekingWaitsForVideoCompositionRendering"))
	return rv
}
func (p AVPlayerItem) SetSeekingWaitsForVideoCompositionRendering(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setSeekingWaitsForVideoCompositionRendering:"), value)
}
// The audio mix parameters to be applied during playback.
//
// # Discussion
// 
// An audio mix can only be used with file-based media and is not supported
// for use with media served using HTTP Live Streaming.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/audioMix
func (p AVPlayerItem) AudioMix() IAVAudioMix {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("audioMix"))
	return AVAudioMixFromID(objc.ID(rv))
}
func (p AVPlayerItem) SetAudioMix(value IAVAudioMix) {
	objc.Send[struct{}](p.ID, objc.Sel("setAudioMix:"), value)
}
// The processing algorithm used to manage audio pitch for scaled audio edits.
//
// # Discussion
// 
// The supported constants are defined in Time Pitch Algorithm Settings.
// 
// An [InvalidArgumentException] will be raised if this property is set to a
// value other than the defined constants.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/audioTimePitchAlgorithm
func (p AVPlayerItem) AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("audioTimePitchAlgorithm"))
	return AVAudioTimePitchAlgorithm(foundation.NSStringFromID(rv).String())
}
func (p AVPlayerItem) SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm) {
	objc.Send[struct{}](p.ID, objc.Sel("setAudioTimePitchAlgorithm:"), objc.String(string(value)))
}
// The source audio channel layouts the player item supports for
// spatialization.
//
// # Discussion
// 
// Spatialization uses psychoacoustic methods to create a more immersive audio
// experience when playing content on specialized headphones and speaker
// arrangements.
// 
// The default value for video content is
// [AudioSpatializationFormatMonoStereoAndMultichannel], and
// [AudioSpatializationFormatMultichannel] for audio-only content. Your app
// can set a preferred spatialization format, but a user can change the audio
// spatialization behavior in Control Center.
// 
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/allowedAudioSpatializationFormats
func (p AVPlayerItem) AllowedAudioSpatializationFormats() AVAudioSpatializationFormats {
	rv := objc.Send[AVAudioSpatializationFormats](p.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return AVAudioSpatializationFormats(rv)
}
func (p AVPlayerItem) SetAllowedAudioSpatializationFormats(value AVAudioSpatializationFormats) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}
// A Boolean value that indicates whether the player item allows spatialized
// audio playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isAudioSpatializationAllowed
func (p AVPlayerItem) AudioSpatializationAllowed() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAudioSpatializationAllowed"))
	return rv
}
func (p AVPlayerItem) SetAudioSpatializationAllowed(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAudioSpatializationAllowed:"), value)
}
// An array of outputs associated with the player item.
//
// # Discussion
// 
// This property contains the collection of [AVPlayerItemOutput] objects used
// to transfer media data to the player object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/outputs
func (p AVPlayerItem) Outputs() []AVPlayerItemOutput {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("outputs"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItemOutput {
		return AVPlayerItemOutputFromID(id)
	})
}
// The collection of associated media data collectors.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/mediaDataCollectors
func (p AVPlayerItem) MediaDataCollectors() []AVPlayerItemMediaDataCollector {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("mediaDataCollectors"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItemMediaDataCollector {
		return AVPlayerItemMediaDataCollectorFromID(id)
	})
}
// The desired limit, in bits per second, of network bandwidth consumption for
// this item.
//
// # Discussion
// 
// Set `preferredPeakBitRate` to nonzero to indicate that the player should
// attempt to limit item playback to that bit rate, expressed in bits per
// second.
// 
// If the system can’t lower network bandwidth consumption to meet the this
// value, it reduces it as much as possible while it continues to play the
// item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredPeakBitRate
func (p AVPlayerItem) PreferredPeakBitRate() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("preferredPeakBitRate"))
	return rv
}
func (p AVPlayerItem) SetPreferredPeakBitRate(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredPeakBitRate:"), value)
}
// The duration the player should buffer media from the network ahead of the
// playhead to guard against playback disruption.
//
// # Discussion
// 
// This property defines the preferred forward buffer duration in seconds. If
// set to 0, the player will choose an appropriate level of buffering for most
// use cases. Setting this property to a low value will increase the chance
// that playback will stall and re-buffer, while setting it to a high value
// will increase demand on system resources.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/preferredForwardBufferDuration
func (p AVPlayerItem) PreferredForwardBufferDuration() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("preferredForwardBufferDuration"))
	return rv
}
func (p AVPlayerItem) SetPreferredForwardBufferDuration(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredForwardBufferDuration:"), value)
}
// A Boolean value that indicates whether the player item can use network
// resources to keep the playback state up to date while paused.
//
// # Discussion
// 
// For live streaming content, the player item may need to use extra
// networking and power resources to keep playback state up to date when
// paused. For example, when this property is set to true, the
// [SeekableTimeRanges] property will be periodically updated to reflect the
// current state of the live stream.
// 
// To minimize power usage, avoid setting this property to `true` when you do
// not need playback state to stay up to date while paused.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/canUseNetworkResourcesForLiveStreamingWhilePaused
func (p AVPlayerItem) CanUseNetworkResourcesForLiveStreamingWhilePaused() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("canUseNetworkResourcesForLiveStreamingWhilePaused"))
	return rv
}
func (p AVPlayerItem) SetCanUseNetworkResourcesForLiveStreamingWhilePaused(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setCanUseNetworkResourcesForLiveStreamingWhilePaused:"), value)
}
// A Boolean value that indicates whether the content has been authorized by
// the user.
//
// # Discussion
// 
// This property reports whether the user has provided the necessary
// credentials to the system in order for the content to be decrypted for
// playback.
// 
// Content authorization is independent of application authorization (see
// [ApplicationAuthorizedForPlayback]) and that both must be granted in order
// for an application to be allowed to play protected content.
// 
// This property is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isContentAuthorizedForPlayback
func (p AVPlayerItem) ContentAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isContentAuthorizedForPlayback"))
	return rv
}
// A Boolean value that indicates whether authorization is required to play
// the content.
//
// # Discussion
// 
// This property reports whether authorization is required for the item’s
// content to be played. If it does not require authorization, then none of
// the other authorization-related methods or properties apply (though they
// will return sensible values where possible).
// 
// This property is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isAuthorizationRequiredForPlayback
func (p AVPlayerItem) AuthorizationRequiredForPlayback() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAuthorizationRequiredForPlayback"))
	return rv
}
// A Boolean value that indicates whether the application can be used to play
// the content.
//
// # Discussion
// 
// This property reports whether or not the calling application is authorized
// to play the content associated with the item.
// 
// Application authorization is independent of content authorization (see
// [ContentAuthorizedForPlayback]) and that both must be granted in order for
// an application to be allowed to play protected content. Also, unlike
// content authorization, application authorization is not dependent on user
// credentials (that is, if `applicationAuthorizedForPlayback` is [false],
// there are no means to obtain authorization).
// 
// This property is not key-value observable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/isApplicationAuthorizedForPlayback
func (p AVPlayerItem) ApplicationAuthorizedForPlayback() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isApplicationAuthorizedForPlayback"))
	return rv
}
// The status of the most recent content authorization request.
//
// # Discussion
// 
// This property reports the authorization status as determined by the most
// recent call to
// [RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler].
// 
// The value will be [ContentAuthorizationUnknown] before the first call and
// between the time a request call is made and just prior to the completion
// handler being executed (thus it is safe to query this property from the
// completion handler).
// 
// This value is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/contentAuthorizationRequestStatus
func (p AVPlayerItem) ContentAuthorizationRequestStatus() AVContentAuthorizationStatus {
	rv := objc.Send[AVContentAuthorizationStatus](p.ID, objc.Sel("contentAuthorizationRequestStatus"))
	return AVContentAuthorizationStatus(rv)
}
// The asset provided during initialization.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/asset
func (p AVPlayerItem) Asset() IAVAsset {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("asset"))
	return AVAssetFromID(objc.ID(rv))
}
// The array of asset keys to be automatically loaded before the player item
// is ready to play.
//
// # Discussion
// 
// The value of each key in `automaticallyLoadedAssetKeys` will automatically
// be loaded by the [Asset] prior to the player item reaching a status of
// [PlayerItemStatusReadyToPlay]. When this status is reached, the asset’s
// [statusOfValue(forKey:error:)] method returns [KeyValueStatusLoaded] for
// the status of all keys in the array. If loading of any of the asset’s key
// values fails, the player item’s [Status] will change to
// [PlayerItemStatusFailed].
//
// [statusOfValue(forKey:error:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading/statusOfValue(forKey:error:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/automaticallyLoadedAssetKeys
func (p AVPlayerItem) AutomaticallyLoadedAssetKeys() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("automaticallyLoadedAssetKeys"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for AVMetricEventStreamPublisher
			

// SeekToTime is a synchronous wrapper around [AVPlayerItem.SeekToTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItem) SeekToTime(ctx context.Context, time objectivec.IObject) (bool, error) {
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

// SeekToTimeToleranceBeforeToleranceAfter is a synchronous wrapper around [AVPlayerItem.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItem) SeekToTimeToleranceBeforeToleranceAfter(ctx context.Context, time objectivec.IObject, toleranceBefore objectivec.IObject, toleranceAfter objectivec.IObject) (bool, error) {
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

// SeekToDate is a synchronous wrapper around [AVPlayerItem.SeekToDateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItem) SeekToDate(ctx context.Context, date foundation.INSDate) (bool, error) {
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

// RequestContentAuthorizationAsynchronouslyWithTimeoutInterval is a synchronous wrapper around [AVPlayerItem.RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (p AVPlayerItem) RequestContentAuthorizationAsynchronouslyWithTimeoutInterval(ctx context.Context, timeoutInterval float64) error {
	done := make(chan struct{}, 1)
	p.RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler(timeoutInterval, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

