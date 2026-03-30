// Code generated from Apple documentation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

var (
	// AVAssetChapterMetadataGroupsDidChangeNotification is posted when the collection of arrays of timed metadata groups representing chapters of an AVAsset change and when any of the contents of the timed metadata groups change, but only for changes that occur after the status of the value of @“availableChapterLocales” has reached AVKeyValueStatusLoaded.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetChapterMetadataGroupsDidChangeNotification
	AVAssetChapterMetadataGroupsDidChangeNotification string
	// AVAssetContainsFragmentsDidChangeNotification is a notification the system posts when an asset’s fragments change.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetContainsFragmentsDidChangeNotification
	AVAssetContainsFragmentsDidChangeNotification string
	// AVAssetDurationDidChangeNotification is a notification the system posts when a fragmented asset minder observes a change to a fragmented asset’s duration.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDurationDidChangeNotification
	AVAssetDurationDidChangeNotification string
	// AVAssetExportPreset1280x720 is a preset to export a 1280 by 720 movie that contains H.264 video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPreset1280x720
	AVAssetExportPreset1280x720 string
	// AVAssetExportPreset1920x1080 is a preset to export a 1920 by 1080 movie that contains H.264 video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPreset1920x1080
	AVAssetExportPreset1920x1080 string
	// AVAssetExportPreset3840x2160 is a preset to export a 3840 by 2160 movie that contains H.264 video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPreset3840x2160
	AVAssetExportPreset3840x2160 string
	// AVAssetExportPreset640x480 is a preset to export a 640 by 480 movie that contains H.264 video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPreset640x480
	AVAssetExportPreset640x480 string
	// AVAssetExportPreset960x540 is a preset to export a 960 by 540 movie that contains H.264 video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPreset960x540
	AVAssetExportPreset960x540 string
	// AVAssetExportPresetAppleM4A is a preset to export an audio-only MPEG 4 Audio file with appropriate iTunes gapless playback data.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4A
	AVAssetExportPresetAppleM4A string
	// AVAssetExportPresetAppleM4V1080pHD is a preset to export a 1080p High Definition format suitable for playing on Apple devices.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4V1080pHD
	AVAssetExportPresetAppleM4V1080pHD string
	// AVAssetExportPresetAppleM4V480pSD is a preset to export a 480p Standard Definition format suitable for playing on Apple devices.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4V480pSD
	AVAssetExportPresetAppleM4V480pSD string
	// AVAssetExportPresetAppleM4V720pHD is a preset to export a 720p High Definition format suitable for playing on Apple devices.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4V720pHD
	AVAssetExportPresetAppleM4V720pHD string
	// AVAssetExportPresetAppleM4VAppleTV is a preset to export a format suitable for playing on Apple TV.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4VAppleTV
	AVAssetExportPresetAppleM4VAppleTV string
	// AVAssetExportPresetAppleM4VCellular is a preset to export a format suitable for playing on Apple devices when it streams over a cellular network.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4VCellular
	AVAssetExportPresetAppleM4VCellular string
	// AVAssetExportPresetAppleM4VWiFi is a preset to export a format suitable for playing on Apple devices it streams over a WiFi network.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4VWiFi
	AVAssetExportPresetAppleM4VWiFi string
	// AVAssetExportPresetAppleM4ViPod is a preset to export a format suitable for playing on an iPod.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleM4ViPod
	AVAssetExportPresetAppleM4ViPod string
	// AVAssetExportPresetAppleProRes422LPCM is a preset to export a QuickTime movie with Apple ProRes 422 video and LPCM audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleProRes422LPCM
	AVAssetExportPresetAppleProRes422LPCM string
	// AVAssetExportPresetAppleProRes4444LPCM is a preset to export a QuickTime movie with Apple ProRes 4444 video and LPCM audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetAppleProRes4444LPCM
	AVAssetExportPresetAppleProRes4444LPCM string
	// AVAssetExportPresetHEVC1920x1080 is a preset to export a 1920 by 1080 movie that contains HEVC video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVC1920x1080
	AVAssetExportPresetHEVC1920x1080 string
	// AVAssetExportPresetHEVC1920x1080WithAlpha is a preset to export a 1920 by 1080 movie that contains HEVC video with alpha and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVC1920x1080WithAlpha
	AVAssetExportPresetHEVC1920x1080WithAlpha string
	// AVAssetExportPresetHEVC3840x2160 is a preset to export a 3840 by 2160 movie that contains HEVC video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVC3840x2160
	AVAssetExportPresetHEVC3840x2160 string
	// AVAssetExportPresetHEVC3840x2160WithAlpha is a preset to export a 3840 by 2160 movie that contains HEVC video with alpha and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVC3840x2160WithAlpha
	AVAssetExportPresetHEVC3840x2160WithAlpha string
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVC4320x2160
	AVAssetExportPresetHEVC4320x2160 string
	// AVAssetExportPresetHEVC7680x4320 is a preset to export a 7680 by 4320 movie that contains HEVC video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVC7680x4320
	AVAssetExportPresetHEVC7680x4320 string
	// AVAssetExportPresetHEVCHighestQuality is a preset to export the highest available video quality and HEVC video compression.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVCHighestQuality
	AVAssetExportPresetHEVCHighestQuality string
	// AVAssetExportPresetHEVCHighestQualityWithAlpha is a preset to export the highest available video quality and HEVC video compression with alpha.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHEVCHighestQualityWithAlpha
	AVAssetExportPresetHEVCHighestQualityWithAlpha string
	// AVAssetExportPresetHighestQuality is a preset to export a high-quality movie file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetHighestQuality
	AVAssetExportPresetHighestQuality string
	// AVAssetExportPresetLowQuality is a preset to export a low-quality movie file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetLowQuality
	AVAssetExportPresetLowQuality string
	// AVAssetExportPresetMVHEVC1440x1440 is a preset to export a 1440 by 1440 movie that contains MV-HEVC video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetMVHEVC1440x1440
	AVAssetExportPresetMVHEVC1440x1440 string
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetMVHEVC4320x4320
	AVAssetExportPresetMVHEVC4320x4320 string
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetMVHEVC7680x7680
	AVAssetExportPresetMVHEVC7680x7680 string
	// AVAssetExportPresetMVHEVC960x960 is a preset to export a 960 by 960 movie that contains MV-HEVC video and AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetMVHEVC960x960
	AVAssetExportPresetMVHEVC960x960 string
	// AVAssetExportPresetMediumQuality is a preset to export a medium-quality movie file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetMediumQuality
	AVAssetExportPresetMediumQuality string
	// AVAssetExportPresetPassthrough is a preset to export the asset in its current format, unless otherwise prohibited.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetPassthrough
	AVAssetExportPresetPassthrough string
	// AVAssetMediaSelectionGroupsDidChangeNotification is posted when the collection of media selection groups provided by an AVAsset changes and when any of the contents of its media selection groups change, but only for changes that occur after the status of the value of @“availableMediaCharacteristicsWithMediaSelectionOptions” has reached AVKeyValueStatusLoaded.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetMediaSelectionGroupsDidChangeNotification
	AVAssetMediaSelectionGroupsDidChangeNotification string
	// AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey is specifies whether the content key request requires a persistable key to be returned from the key vendor.
	//
	// Deprecated: Deprecated since macOS 15.0. Use -[AVPersistableContentKeyRequest persistableContentKeyFromKeyVendorResponse:options:error:] instead
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey
	AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey string
	// AVAssetTrackSegmentsDidChangeNotification is posted when the array of segments of an AVFragmentedAssetTrack changes while the associated instance of AVFragmentedAsset is being minded by an AVFragmentedAssetMinder, but only for changes that occur after the status of the value of @“segments” has reached AVKeyValueStatusLoaded.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegmentsDidChangeNotification
	AVAssetTrackSegmentsDidChangeNotification string
	// AVAssetTrackTimeRangeDidChangeNotification is posted when the timeRange of an AVFragmentedAssetTrack changes while the associated instance of AVFragmentedAsset is being minded by an AVFragmentedAssetMinder, but only for changes that occur after the status of the value of @“timeRange” has reached AVKeyValueStatusLoaded.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackTimeRangeDidChangeNotification
	AVAssetTrackTimeRangeDidChangeNotification string
	// AVAssetTrackTrackAssociationsDidChangeNotification is posted when the collection of track associations of an AVAssetTrack changes, but only for changes that occur after the status of the value of @“availableTrackAssociationTypes” has reached AVKeyValueStatusLoaded.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackTrackAssociationsDidChangeNotification
	AVAssetTrackTrackAssociationsDidChangeNotification string
	// AVAssetWasDefragmentedNotification is a notification the system posts when a fragmented asset minder observes that the system defragments the asset on disk.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWasDefragmentedNotification
	AVAssetWasDefragmentedNotification string
	// AVCaptureSessionErrorKey is key to retrieve the error object from a [runtimeErrorNotification] user info dictionary.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSessionErrorKey
	AVCaptureSessionErrorKey string
	// AVContentKeyRequestProtocolVersionsKey is a key that specifies the versions of the content protection protocol supported by the application.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequestProtocolVersionsKey
	AVContentKeyRequestProtocolVersionsKey string
	// AVContentKeyRequestRandomDeviceIdentifierSeedKey is value is an NSData containing a 16-byte seed to randomize the user’s deviceID contained in the SPC blob during FairPlay key exchange.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequestRandomDeviceIdentifierSeedKey
	AVContentKeyRequestRandomDeviceIdentifierSeedKey string
	// AVContentKeyRequestRequiresValidationDataInSecureTokenKey is a key that requires the secure token to have extended validation data.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequestRequiresValidationDataInSecureTokenKey
	AVContentKeyRequestRequiresValidationDataInSecureTokenKey string
	// AVContentKeyRequestShouldRandomizeDeviceIdentifierKey is value is an Boolean indicating whether the user’s deviceID contained in the SPC blob during FairPlay key exchange should be randomized using a system generated seed.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequestShouldRandomizeDeviceIdentifierKey
	AVContentKeyRequestShouldRandomizeDeviceIdentifierKey string
	// AVErrorDeviceKey is the user information key to retrieve the device name.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorDeviceKey
	AVErrorDeviceKey string
	// AVErrorDiscontinuityFlagsKey is the user information key to retrieve discontinuity flags.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorDiscontinuityFlagsKey
	AVErrorDiscontinuityFlagsKey string
	// AVErrorFileSizeKey is the user information key to retrieve the file size in bytes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorFileSizeKey
	AVErrorFileSizeKey string
	// AVErrorFileTypeKey is the user information key to retrieve the file type.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorFileTypeKey
	AVErrorFileTypeKey string
	// AVErrorMediaSubTypeKey is the user information key to retrieve the media subtype.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorMediaSubTypeKey
	AVErrorMediaSubTypeKey string
	// AVErrorMediaTypeKey is the user information key to retrieve the media type.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorMediaTypeKey
	AVErrorMediaTypeKey string
	// AVErrorPIDKey is the user information key to retrieve the process ID value.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorPIDKey
	AVErrorPIDKey string
	// AVErrorPersistentTrackIDKey is the user information key to retrieve the track’s persistent identifier.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorPersistentTrackIDKey
	AVErrorPersistentTrackIDKey string
	// AVErrorPresentationTimeStampKey is the user information key to retrieve the presentation time stamp.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorPresentationTimeStampKey
	AVErrorPresentationTimeStampKey string
	// AVErrorRecordingSuccessfullyFinishedKey is the user information key to retrieve a Boolean value that indicates whether recording finished successfully.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorRecordingSuccessfullyFinishedKey
	AVErrorRecordingSuccessfullyFinishedKey string
	// AVErrorTimeKey is the user information key to retrieve the error time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVErrorTimeKey
	AVErrorTimeKey string
	// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieContainsMovieFragmentsDidChangeNotification
	AVFragmentedMovieContainsMovieFragmentsDidChangeNotification string
	// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieDurationDidChangeNotification
	AVFragmentedMovieDurationDidChangeNotification string
	// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieTrackSegmentsDidChangeNotification
	AVFragmentedMovieTrackSegmentsDidChangeNotification string
	// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieTrackTimeRangeDidChangeNotification
	AVFragmentedMovieTrackTimeRangeDidChangeNotification string
	// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieWasDefragmentedNotification
	AVFragmentedMovieWasDefragmentedNotification string
	// AVMovieReferenceRestrictionsKey is a key that specifies restrictions for a movie to use when it resolves references to external media data.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMovieReferenceRestrictionsKey
	AVMovieReferenceRestrictionsKey string
	// AVMovieShouldSupportAliasDataReferencesKey is a key that specifies a Boolean value that indicates whether the system parses and resolves alias data references in the movie.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMovieShouldSupportAliasDataReferencesKey
	AVMovieShouldSupportAliasDataReferencesKey string
	// AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonKey is a key to retrieve the reason for an out-of-sync state notification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/snapshotsOutOfSyncReasonKey
	AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonKey string
	// AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeErrorKey is a key to retrieve the error related to a change in an interstitial event’s asset list response.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/assetListResponseStatusDidChangeErrorKey
	AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeErrorKey string
	// AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeEventKey is a key to retrieve the interstitial event that has an asset list response status change.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/assetListResponseStatusDidChangeEventKey
	AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeEventKey string
	// AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeStatusKey is a key to retrieve the asset list response status.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/assetListResponseStatusDidChangeStatusKey
	AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeStatusKey string
	// AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeEventKey is the dictionary key for the AVPlayerInterstitial event that had its skippable event state changed in the payload of the AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippableStateDidChangeEventKey
	AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeEventKey string
	// AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeSkipControlLabelKey is the dictionary key for the skip label of the event in the payload of the AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippableStateDidChangeSkipControlLabelKey
	AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeSkipControlLabelKey string
	// AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeStateKey is the dictionary key for the skippable event state in the payload of the AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippableStateDidChangeStateKey
	AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeStateKey string
	// AVPlayerInterstitialEventMonitorCurrentEventSkippedEventKey is the dictionary key for the AVPlayerInterstitialEvent that was skipped in the payload of the AVPlayerInterstitialEventMonitorCurrentEventSkippedNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippedEventKey
	AVPlayerInterstitialEventMonitorCurrentEventSkippedEventKey string
	// AVPlayerInterstitialEventMonitorInterstitialEventDidFinishDidPlayEntireEventKey is the dictionary key to indicate whether the event that finished playing was fully played out in the payload of the AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventDidFinishDidPlayEntireEventKey
	AVPlayerInterstitialEventMonitorInterstitialEventDidFinishDidPlayEntireEventKey string
	// AVPlayerInterstitialEventMonitorInterstitialEventDidFinishEventKey is the dictionary key for the AVPlayerInterstitialEvent that finished playing in the payload of the AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventDidFinishEventKey
	AVPlayerInterstitialEventMonitorInterstitialEventDidFinishEventKey string
	// AVPlayerInterstitialEventMonitorInterstitialEventDidFinishPlayoutTimeKey is the dictionary key for the playout time of the event that finished playing in the payload of the AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventDidFinishPlayoutTimeKey
	AVPlayerInterstitialEventMonitorInterstitialEventDidFinishPlayoutTimeKey string
	// AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledErrorKey is the dictionary key to indicate whether the event that was unscheduled was due to an error.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventWasUnscheduledErrorKey
	AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledErrorKey string
	// AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledEventKey is the dictionary key for the AVPlayerInterstitialEvent that was unscheduled in the payload of the AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledNotification.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventWasUnscheduledEventKey
	AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledEventKey string
	// AVPlayerInterstitialEventMonitorScheduleRequestErrorKey is userInfo dictionary key for the AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification. Value is NSError. Absent if the request succeeded.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitorScheduleRequestErrorKey
	AVPlayerInterstitialEventMonitorScheduleRequestErrorKey string
	// AVPlayerInterstitialEventMonitorScheduleRequestIdentifierKey is userInfo dictionary key for the AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification. Value is NSString.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitorScheduleRequestIdentifierKey
	AVPlayerInterstitialEventMonitorScheduleRequestIdentifierKey string
	// AVPlayerInterstitialEventMonitorScheduleRequestResponseKey is userInfo dictionary key for the AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification. Value is NSData. Absent if the request failed.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitorScheduleRequestResponseKey
	AVPlayerInterstitialEventMonitorScheduleRequestResponseKey string
	// AVPlayerItemFailedToPlayToEndTimeErrorKey is the key to retrieve the error object from the notification’s user information dictionary.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemFailedToPlayToEndTimeErrorKey
	AVPlayerItemFailedToPlayToEndTimeErrorKey string
	// AVPlayerItemTimeJumpedOriginatingParticipantKey is a key to retrieve a unique identifier of the participant that caused the time jump.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/timeJumpedOriginatingParticipantKey
	AVPlayerItemTimeJumpedOriginatingParticipantKey string
	// AVPlayerItemTrackVideoFieldModeDeinterlaceFields is a video field mode that requests deinterlacing of video fields.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemTrackVideoFieldModeDeinterlaceFields
	AVPlayerItemTrackVideoFieldModeDeinterlaceFields string
	// AVPlayerRateDidChangeOriginatingParticipantKey is a key to retrieve the identifier of the participant that originates the rate change.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/rateDidChangeOriginatingParticipantKey
	AVPlayerRateDidChangeOriginatingParticipantKey string
	// AVPlayerRateDidChangeReasonKey is a key to retrieve the reason for the rate change.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/rateDidChangeReasonKey
	AVPlayerRateDidChangeReasonKey string
	// AVSampleBufferAudioRendererFlushTimeKey is the key that indicates the presentation timestamp of the first queued sample that was flushed.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRendererFlushTimeKey
	AVSampleBufferAudioRendererFlushTimeKey string
	// AVSampleBufferDisplayLayerFailedToDecodeNotification is a notification the system posts when a sample buffer display layer fails to decode.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayerFailedToDecodeNotification
	AVSampleBufferDisplayLayerFailedToDecodeNotification string
	// AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey is the key for the corresponding error.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey
	AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey string
	// AVSampleBufferVideoRendererDidFailToDecodeNotificationErrorKey is a key to retrieve an error object that provides the details of the failure.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/didFailToDecodeNotificationErrorKey
	AVSampleBufferVideoRendererDidFailToDecodeNotificationErrorKey string
	// AVStreamingKeyDeliveryContentKeyType is a URL for a content key.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVStreamingKeyDeliveryContentKeyType
	AVStreamingKeyDeliveryContentKeyType string
	// AVStreamingKeyDeliveryPersistentContentKeyType is a URL for a persistent content key.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVStreamingKeyDeliveryPersistentContentKeyType
	AVStreamingKeyDeliveryPersistentContentKeyType string
	// AVURLAssetAllowsCellularAccessKey is a Boolean value that indicates whether the system can make network requests on behalf of the asset when connected to a cellular network.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetAllowsCellularAccessKey
	AVURLAssetAllowsCellularAccessKey string
	// AVURLAssetAllowsConstrainedNetworkAccessKey is a Boolean value that indicates whether the system allows network requests on behalf of this asset to use the constrained interface.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetAllowsConstrainedNetworkAccessKey
	AVURLAssetAllowsConstrainedNetworkAccessKey string
	// AVURLAssetAllowsExpensiveNetworkAccessKey is a Boolean value that indicates whether the system allows network requests on behalf of this asset to use the expensive interface.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetAllowsExpensiveNetworkAccessKey
	AVURLAssetAllowsExpensiveNetworkAccessKey string
	// AVURLAssetHTTPCookiesKey is the HTTP cookies that a URL asset may send with HTTP requests.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetHTTPCookiesKey
	AVURLAssetHTTPCookiesKey string
	// AVURLAssetHTTPUserAgentKey is a key that specifies the user agent of requests that an asset makes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetHTTPUserAgentKey
	AVURLAssetHTTPUserAgentKey string
	// AVURLAssetOverrideMIMETypeKey is a key that specifies the MIME type to use to identify the format of a media resource.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetOverrideMIMETypeKey
	AVURLAssetOverrideMIMETypeKey string
	// AVURLAssetPreferPreciseDurationAndTimingKey is a Boolean value that indicates whether the asset should provide accurate duration and precise random access by time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetPreferPreciseDurationAndTimingKey
	AVURLAssetPreferPreciseDurationAndTimingKey string
	// AVURLAssetPrimarySessionIdentifierKey is specifies a UUID to set as the session identifier for HTTP requests that the asset makes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetPrimarySessionIdentifierKey
	AVURLAssetPrimarySessionIdentifierKey string
	// AVURLAssetReferenceRestrictionsKey is a value that represents the restrictions used by the asset when resolving references to external media data.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetReferenceRestrictionsKey
	AVURLAssetReferenceRestrictionsKey string
	// AVURLAssetShouldParseExternalSphericalTagsKey is indicates whether additional projected media signaling in the asset should be parsed and resolved as format description extensions.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetShouldParseExternalSphericalTagsKey
	AVURLAssetShouldParseExternalSphericalTagsKey string
	// AVURLAssetShouldSupportAliasDataReferencesKey is a Boolean value that indicates whether the system parses and resolves alias data references in the asset.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetShouldSupportAliasDataReferencesKey
	AVURLAssetShouldSupportAliasDataReferencesKey string
	// AVURLAssetURLRequestAttributionKey is a value that specifies the attribution of the URLs that this asset requests.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVURLAssetURLRequestAttributionKey
	AVURLAssetURLRequestAttributionKey string
	// AVVideoAllowFrameReorderingKey is a key to access permission to reorder frames.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoAllowFrameReorderingKey
	AVVideoAllowFrameReorderingKey string
	// AVVideoAllowWideColorKey is the key for a dictionary that indicates whether the client can process wide color.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoAllowWideColorKey
	AVVideoAllowWideColorKey string
	// AVVideoAppleProRAWBitDepthKey is a key to access the Apple ProRAW bit depth.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoAppleProRAWBitDepthKey
	AVVideoAppleProRAWBitDepthKey string
	// AVVideoAverageBitRateKey is a key to access the average bit rate—as bits per second—used in compressing video.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoAverageBitRateKey
	AVVideoAverageBitRateKey string
	// AVVideoAverageNonDroppableFrameRateKey is the desired average number of non-droppable frames to be encoded for each second of video.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoAverageNonDroppableFrameRateKey
	AVVideoAverageNonDroppableFrameRateKey string
	// AVVideoCleanApertureHeightKey is a key to access the height of video that’s free from transition artifacts caused by signal encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureHeightKey
	AVVideoCleanApertureHeightKey string
	// AVVideoCleanApertureHorizontalOffsetKey is a key to access the horizontal offset of video that’s free from transition artifacts caused by signal encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureHorizontalOffsetKey
	AVVideoCleanApertureHorizontalOffsetKey string
	// AVVideoCleanApertureKey is a key that defines the region within the video dimension displayed during playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureKey
	AVVideoCleanApertureKey string
	// AVVideoCleanApertureVerticalOffsetKey is a key to access the vertical offset of video that’s free from transition artifacts caused by signal encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureVerticalOffsetKey
	AVVideoCleanApertureVerticalOffsetKey string
	// AVVideoCleanApertureWidthKey is a key to access the width of video that’s free from transition artifacts caused by signal encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCleanApertureWidthKey
	AVVideoCleanApertureWidthKey string
	// AVVideoCodecKey is a key to access the name of the codec for compressing video.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecKey
	AVVideoCodecKey string
	// AVVideoColorPrimariesKey is the key to identify color primaries in a color properties dictionary.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimariesKey
	AVVideoColorPrimariesKey string
	// AVVideoColorPrimaries_EBU_3213 is the color primary is in the EBU Tech. 3213 color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimaries_EBU_3213
	AVVideoColorPrimaries_EBU_3213 string
	// AVVideoColorPrimaries_ITU_R_2020 is the color primary is in the ITU_R BT.2020 color space for ultra high definition television.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimaries_ITU_R_2020
	AVVideoColorPrimaries_ITU_R_2020 string
	// AVVideoColorPrimaries_ITU_R_709_2 is the color primary is in the ITU_R BT.709 color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimaries_ITU_R_709_2
	AVVideoColorPrimaries_ITU_R_709_2 string
	// AVVideoColorPrimaries_P3_D65 is the color primary uses the DCI-P3 D65 color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimaries_P3_D65
	AVVideoColorPrimaries_P3_D65 string
	// AVVideoColorPrimaries_SMPTE_C is the color primary uses the SMPTE C color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPrimaries_SMPTE_C
	AVVideoColorPrimaries_SMPTE_C string
	// AVVideoColorPropertiesKey is the key for a dictionary that contains properties specifying video color.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoColorPropertiesKey
	AVVideoColorPropertiesKey string
	// AVVideoCompressionPropertiesKey is a key to access the dictionary of compression properties for a video asset.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompressionPropertiesKey
	AVVideoCompressionPropertiesKey string
	// AVVideoDecompressionPropertiesKey is the key that indicates the video decompression properties to pass to the video decoder.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoDecompressionPropertiesKey
	AVVideoDecompressionPropertiesKey string
	// AVVideoEncoderSpecificationKey is the video encoder specification includes options for choosing a specific video encoder.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoEncoderSpecificationKey
	AVVideoEncoderSpecificationKey string
	// AVVideoExpectedSourceFrameRateKey is the expected source frame rate.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoExpectedSourceFrameRateKey
	AVVideoExpectedSourceFrameRateKey string
	// AVVideoH264EntropyModeCABAC is the encoder uses Context-based Adaptive Binary Arithmetic Coding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoH264EntropyModeCABAC
	AVVideoH264EntropyModeCABAC string
	// AVVideoH264EntropyModeCAVLC is the encoder uses Context-based Adaptive Variable Length Coding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoH264EntropyModeCAVLC
	AVVideoH264EntropyModeCAVLC string
	// AVVideoH264EntropyModeKey is the entropy encoding mode for H.264 compression.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoH264EntropyModeKey
	AVVideoH264EntropyModeKey string
	// AVVideoHeightKey is a key to access the height of the video in pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoHeightKey
	AVVideoHeightKey string
	// AVVideoMaxKeyFrameIntervalDurationKey is a key to access the maximum interval duration between keyframes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoMaxKeyFrameIntervalDurationKey
	AVVideoMaxKeyFrameIntervalDurationKey string
	// AVVideoMaxKeyFrameIntervalKey is a key to access the maximum interval between keyframes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoMaxKeyFrameIntervalKey
	AVVideoMaxKeyFrameIntervalKey string
	// AVVideoPixelAspectRatioHorizontalSpacingKey is a key to access the pixel aspect ratio horizontal spacing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPixelAspectRatioHorizontalSpacingKey
	AVVideoPixelAspectRatioHorizontalSpacingKey string
	// AVVideoPixelAspectRatioKey is a key to access the video’s pixel aspect ratio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPixelAspectRatioKey
	AVVideoPixelAspectRatioKey string
	// AVVideoPixelAspectRatioVerticalSpacingKey is a key to access the pixel aspect ratio vertical spacing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoPixelAspectRatioVerticalSpacingKey
	AVVideoPixelAspectRatioVerticalSpacingKey string
	// AVVideoProfileLevelH264Baseline30 is a baseline-level 3.0 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Baseline30
	AVVideoProfileLevelH264Baseline30 string
	// AVVideoProfileLevelH264Baseline31 is a baseline-level 3.1 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Baseline31
	AVVideoProfileLevelH264Baseline31 string
	// AVVideoProfileLevelH264Baseline41 is a baseline-level 4.1 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Baseline41
	AVVideoProfileLevelH264Baseline41 string
	// AVVideoProfileLevelH264BaselineAutoLevel is a baseline auto level profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264BaselineAutoLevel
	AVVideoProfileLevelH264BaselineAutoLevel string
	// AVVideoProfileLevelH264High40 is a high-level 4.0 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264High40
	AVVideoProfileLevelH264High40 string
	// AVVideoProfileLevelH264High41 is a high-level 4.1 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264High41
	AVVideoProfileLevelH264High41 string
	// AVVideoProfileLevelH264HighAutoLevel is a high profile auto level profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264HighAutoLevel
	AVVideoProfileLevelH264HighAutoLevel string
	// AVVideoProfileLevelH264Main30 is a main-level 3.0 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Main30
	AVVideoProfileLevelH264Main30 string
	// AVVideoProfileLevelH264Main31 is a main-level 3.1 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Main31
	AVVideoProfileLevelH264Main31 string
	// AVVideoProfileLevelH264Main32 is a main-level 3.2 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Main32
	AVVideoProfileLevelH264Main32 string
	// AVVideoProfileLevelH264Main41 is a main-level 4.1 profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264Main41
	AVVideoProfileLevelH264Main41 string
	// AVVideoProfileLevelH264MainAutoLevel is a main profile auto level profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelH264MainAutoLevel
	AVVideoProfileLevelH264MainAutoLevel string
	// AVVideoProfileLevelKey is a key to access the video profile.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoProfileLevelKey
	AVVideoProfileLevelKey string
	// AVVideoQualityKey is a key to set the JPEG compression quality of the video.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoQualityKey
	AVVideoQualityKey string
	// AVVideoScalingModeFit is the string identifier for scaling a video to fit the surrounding view’s dimensions.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeFit
	AVVideoScalingModeFit string
	// AVVideoScalingModeKey is a key to retrieve the video scaling mode from a dictionary.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeKey
	AVVideoScalingModeKey string
	// AVVideoScalingModeResize is the string identifier for resizing a video to fit the surrounding view’s dimensions.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeResize
	AVVideoScalingModeResize string
	// AVVideoScalingModeResizeAspect is the string identifier for resizing a video to its surrounding view’s shorter dimension while preserving its aspect ratio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeResizeAspect
	AVVideoScalingModeResizeAspect string
	// AVVideoScalingModeResizeAspectFill is the string identifier for resizing a video to fit the surrounding view’s longer dimension while preserving aspect ratio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoScalingModeResizeAspectFill
	AVVideoScalingModeResizeAspectFill string
	// AVVideoTransferFunctionKey is the key to identify the transfer function in a color properties dictionary.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunctionKey
	AVVideoTransferFunctionKey string
	// AVVideoTransferFunction_IEC_sRGB is the transfer function for the IEC sRGB color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunction_IEC_sRGB
	AVVideoTransferFunction_IEC_sRGB string
	// AVVideoTransferFunction_ITU_R_2100_HLG is the transfer function for the ITU_R BT.2100 color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunction_ITU_R_2100_HLG
	AVVideoTransferFunction_ITU_R_2100_HLG string
	// AVVideoTransferFunction_ITU_R_709_2 is the transfer function for the ITU_R BT.709 color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunction_ITU_R_709_2
	AVVideoTransferFunction_ITU_R_709_2 string
	// AVVideoTransferFunction_Linear is the transfer function for the linear color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunction_Linear
	AVVideoTransferFunction_Linear string
	// AVVideoTransferFunction_SMPTE_240M_1995 is the transfer function for the SMPTE 240M color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunction_SMPTE_240M_1995
	AVVideoTransferFunction_SMPTE_240M_1995 string
	// AVVideoTransferFunction_SMPTE_ST_2084_PQ is the transfer function for the SMPTE 2084 color space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoTransferFunction_SMPTE_ST_2084_PQ
	AVVideoTransferFunction_SMPTE_ST_2084_PQ string
	// AVVideoWidthKey is a key to access the width of the video in pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoWidthKey
	AVVideoWidthKey string
	// AVVideoYCbCrMatrixKey is the key to identify the Y’CbCr matrix in a color properties dictionary.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoYCbCrMatrixKey
	AVVideoYCbCrMatrixKey string
	// AVVideoYCbCrMatrix_ITU_R_2020 is the Y’CbCr color matrix for ITU-R BT.2020 conversion.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoYCbCrMatrix_ITU_R_2020
	AVVideoYCbCrMatrix_ITU_R_2020 string
	// AVVideoYCbCrMatrix_ITU_R_601_4 is the Y’CbCr color matrix for ITU-R BT.601 conversion.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoYCbCrMatrix_ITU_R_601_4
	AVVideoYCbCrMatrix_ITU_R_601_4 string
	// AVVideoYCbCrMatrix_ITU_R_709_2 is the Y’CbCr color matrix for ITU-R BT.709 conversion.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoYCbCrMatrix_ITU_R_709_2
	AVVideoYCbCrMatrix_ITU_R_709_2 string
	// AVVideoYCbCrMatrix_SMPTE_240M_1995 is the Y’CbCr color matrix for SMPTE 240M conversion.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVVideoYCbCrMatrix_SMPTE_240M_1995
	AVVideoYCbCrMatrix_SMPTE_240M_1995 string
)

var ()

var ()

var ()

var ()

var ()

var ()

var (
	// AVCaptionConversionAdjustmentTypeTimeRange is indicates a timing adjustment.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment/AdjustmentType-swift.struct/timeRange
	AVCaptionConversionAdjustmentTypeTimeRange AVCaptionConversionAdjustmentType
)

var (
	// AVCaptionConversionWarningTypeExcessMediaData is a type that indicates one or more captions exceed the media data capacity for media of the type and subtype that the conversion settings specify.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/WarningType-swift.struct/excessMediaData
	AVCaptionConversionWarningTypeExcessMediaData AVCaptionConversionWarningType
)

var (
	// AVCaptionMediaSubTypeKey is a key that identifies the output media subtype of a caption conversion operation.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSettingsKey/mediaSubType
	AVCaptionMediaSubTypeKey AVCaptionSettingsKey
	// AVCaptionMediaTypeKey is a key that identifies the output media type of a caption conversion operation.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSettingsKey/mediaType
	AVCaptionMediaTypeKey AVCaptionSettingsKey
	// AVCaptionTimeCodeFrameDurationKey is a key that identifies the frame duration that the system uses for the time code.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSettingsKey/timeCodeFrameDuration
	AVCaptionTimeCodeFrameDurationKey AVCaptionSettingsKey
	// AVCaptionUseDropFrameTimeCodeKey is a key that identifies whether the system uses drop frame time code.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSettingsKey/useDropFrameTimeCode
	AVCaptionUseDropFrameTimeCodeKey AVCaptionSettingsKey
)

var (
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/aspectratio/ratio16x9
	AVCaptureAspectRatio16x9 objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/aspectratio/ratio1x1
	AVCaptureAspectRatio1x1 objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/aspectratio/ratio3x4
	AVCaptureAspectRatio3x4 objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/aspectratio/ratio4x3
	AVCaptureAspectRatio4x3 objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/aspectratio/ratio9x16
	AVCaptureAspectRatio9x16 objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.struct/critical
	AVCaptureSystemPressureLevelCritical objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.struct/fair
	AVCaptureSystemPressureLevelFair objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.struct/nominal
	AVCaptureSystemPressureLevelNominal objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.struct/serious
	AVCaptureSystemPressureLevelSerious objc.ID
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/systempressurestate-swift.class/level-swift.struct/shutdown
	AVCaptureSystemPressureLevelShutdown objc.ID
)

var (
	// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/subjectareadidchangenotification
	AVCaptureDeviceSubjectAreaDidChangeNotification foundation.NSNotification
)

var ()

var (
	// AVCaptureDeviceWasConnectedNotification is a notification the system posts when a new capture device becomes available.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/wasConnectedNotification
	AVCaptureDeviceWasConnectedNotification foundation.NSNotificationName
	// AVCaptureDeviceWasDisconnectedNotification is a notification the system posts when an existing device becomes unavailable.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/wasDisconnectedNotification
	AVCaptureDeviceWasDisconnectedNotification foundation.NSNotificationName
	// AVCaptureInputPortFormatDescriptionDidChangeNotification is a notification the system posts when the capture input port’s format description changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput/Port/formatDescriptionDidChangeNotification
	AVCaptureInputPortFormatDescriptionDidChangeNotification foundation.NSNotificationName
	// AVCaptureSessionDidStartRunningNotification is a notification the system posts when a capture session starts.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/didStartRunningNotification
	AVCaptureSessionDidStartRunningNotification foundation.NSNotificationName
	// AVCaptureSessionDidStopRunningNotification is a notification the system posts when a capture session stops.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/didStopRunningNotification
	AVCaptureSessionDidStopRunningNotification foundation.NSNotificationName
	// AVCaptureSessionInterruptionEndedNotification is a notification the system posts when an interruption to a capture session finishes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/interruptionEndedNotification
	AVCaptureSessionInterruptionEndedNotification foundation.NSNotificationName
	// AVCaptureSessionRuntimeErrorNotification is a notification the system posts when an error occurs during a capture session.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runtimeErrorNotification
	AVCaptureSessionRuntimeErrorNotification foundation.NSNotificationName
	// AVCaptureSessionWasInterruptedNotification is a notification the system posts when it interrupts a capture session.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/wasInterruptedNotification
	AVCaptureSessionWasInterruptedNotification foundation.NSNotificationName
	// AVFoundationErrorDomain is the error domain of AVFoundation errors.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVFoundationErrorDomain
	AVFoundationErrorDomain foundation.NSErrorDomain
	// AVPlaybackCoordinatorOtherParticipantsDidChangeNotification is a notification that the coordinator posts when its other participants change.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/otherParticipantsDidChangeNotification
	AVPlaybackCoordinatorOtherParticipantsDidChangeNotification foundation.NSNotificationName
	// AVPlaybackCoordinatorSuspensionReasonsDidChangeNotification is a notification that the coordinator posts when its suspension reasons change.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlaybackCoordinator/suspensionReasonsDidChangeNotification
	AVPlaybackCoordinatorSuspensionReasonsDidChangeNotification foundation.NSNotificationName
	// AVPlayerEligibleForHDRPlaybackDidChangeNotification is a notification that’s posted whenever HDR playback eligibility changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/eligibleForHDRPlaybackDidChangeNotification
	AVPlayerEligibleForHDRPlaybackDidChangeNotification foundation.NSNotificationName
	// AVPlayerIntegratedTimelineSnapshotsOutOfSyncNotification is a notification the system posts when the snapshot objects provided by this timeline become out of sync with the current timeline state.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimeline/snapshotsOutOfSyncNotification
	AVPlayerIntegratedTimelineSnapshotsOutOfSyncNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeNotification is a notification the system posts when the status of an interstitial event’s asset list response changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/assetListResponseStatusDidChangeNotification
	AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorCurrentEventDidChangeNotification is a notification the system posts when the monitor’s current interstitial event changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventDidChangeNotification
	AVPlayerInterstitialEventMonitorCurrentEventDidChangeNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification is a notification that’s posted whenever the currentEventSkippableState of an AVPlayerInterstitialEventMonitor changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippableStateDidChangeNotification
	AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorCurrentEventSkippedNotification is a notification that’s posted whenever an event was skipped via skip control.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/currentEventSkippedNotification
	AVPlayerInterstitialEventMonitorCurrentEventSkippedNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorEventsDidChangeNotification is a notification the system posts when the monitor’s schedule of interstitial events changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/eventsDidChangeNotification
	AVPlayerInterstitialEventMonitorEventsDidChangeNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification is a notification that is posted whenever an AVPlayerInterstitialEvent finished playing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventDidFinishNotification
	AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledNotification is a notification that is posted whenever an AVPlayerInterstitialEvent with loaded assets was unscheduled prior to playing.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitor/interstitialEventWasUnscheduledNotification
	AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledNotification foundation.NSNotificationName
	// AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification is a notification that is posted whenever a daterange-schedule request completes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification
	AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification foundation.NSNotificationName
	// AVPlayerItemDidPlayToEndTimeNotification is a notification the system posts when a player item plays to its end time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/didPlayToEndTimeNotification
	AVPlayerItemDidPlayToEndTimeNotification foundation.NSNotificationName
	// AVPlayerItemFailedToPlayToEndTimeNotification is a notification that the system posts when a player item fails to play to its end time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/failedToPlayToEndTimeNotification
	AVPlayerItemFailedToPlayToEndTimeNotification foundation.NSNotificationName
	// AVPlayerItemMediaSelectionDidChangeNotification is a notification the player item posts when its media selection changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/mediaSelectionDidChangeNotification
	AVPlayerItemMediaSelectionDidChangeNotification foundation.NSNotificationName
	// AVPlayerItemNewAccessLogEntryNotification is a notification the system posts when a player item adds a new entry to its access log.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/newAccessLogEntryNotification
	AVPlayerItemNewAccessLogEntryNotification foundation.NSNotificationName
	// AVPlayerItemNewErrorLogEntryNotification is a notification the system posts when a player item adds a new entry to its error log.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/newErrorLogEntryNotification
	AVPlayerItemNewErrorLogEntryNotification foundation.NSNotificationName
	// AVPlayerItemPlaybackStalledNotification is a notification the system posts when a player item media doesn’t arrive in time to continue playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/playbackStalledNotification
	AVPlayerItemPlaybackStalledNotification foundation.NSNotificationName
	// AVPlayerItemRecommendedTimeOffsetFromLiveDidChangeNotification is a notification the player item posts when its offset from the live time changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/recommendedTimeOffsetFromLiveDidChangeNotification
	AVPlayerItemRecommendedTimeOffsetFromLiveDidChangeNotification foundation.NSNotificationName
	// AVPlayerItemTimeJumpedNotification is a notification the system posts when a player item’s time changes discontinuously.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/timeJumpedNotification
	AVPlayerItemTimeJumpedNotification foundation.NSNotificationName
	// AVPlayerRateDidChangeNotification is a notification that a player posts when its rate changes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/rateDidChangeNotification
	AVPlayerRateDidChangeNotification foundation.NSNotificationName
	// AVRouteDetectorMultipleRoutesDetectedDidChangeNotification is a notification the system posts when changes occur to its detected routes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVRouteDetectorMultipleRoutesDetectedDidChangeNotification
	AVRouteDetectorMultipleRoutesDetectedDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRendererOutputConfigurationDidChangeNotification
	AVSampleBufferAudioRendererOutputConfigurationDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRendererWasFlushedAutomaticallyNotification
	AVSampleBufferAudioRendererWasFlushedAutomaticallyNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayerOutputObscuredDueToInsufficientExternalProtectionDidChangeNotification
	AVSampleBufferDisplayLayerOutputObscuredDueToInsufficientExternalProtectionDidChangeNotification foundation.NSNotificationName
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayerReadyForDisplayDidChangeNotification
	AVSampleBufferDisplayLayerReadyForDisplayDidChangeNotification foundation.NSNotificationName
	// AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification is a notification the system posts when a sample buffer display layer changes its decoding requirements.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification
	AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification foundation.NSNotificationName
	// AVSampleBufferRenderSynchronizerRateDidChangeNotification is the synchronizer’s rendering rate changed.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRenderSynchronizer/rateDidChangeNotification
	AVSampleBufferRenderSynchronizerRateDidChangeNotification foundation.NSNotificationName
	// AVSampleBufferVideoRendererDidFailToDecodeNotification is a notification that indicates the video renderer fails to decode a sample buffer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/didFailToDecodeNotification
	AVSampleBufferVideoRendererDidFailToDecodeNotification foundation.NSNotificationName
	// AVSampleBufferVideoRendererRequiresFlushToResumeDecodingDidChangeNotification is a notification that indicates that the video renderer requires flushing to continue rendering sample buffers.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/requiresFlushToResumeDecodingDidChangeNotification
	AVSampleBufferVideoRendererRequiresFlushToResumeDecodingDidChangeNotification foundation.NSNotificationName
)

var (
	// AVCaptureMaxAvailableTorchLevel is a constant that indicates to set the torch to its maximum level.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxAvailableTorchLevel
	AVCaptureMaxAvailableTorchLevel float32
)

var ()

var (
	// AVCaptureSceneMonitoringStatusNotEnoughLight is the light level of the current scene is insufficient for the current set of features to function optimally.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSceneMonitoringStatus/notEnoughLight
	AVCaptureSceneMonitoringStatusNotEnoughLight AVCaptureSceneMonitoringStatus
)

var (
	// AVCaptureSessionPreset1280x720 is a preset suitable for capturing 720p quality (1280 x 720 pixel) video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/hd1280x720
	AVCaptureSessionPreset1280x720 AVCaptureSessionPreset
	// AVCaptureSessionPreset1920x1080 is a preset suitable for capturing 1080p-quality (1920 x 1080 pixels) video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/hd1920x1080
	AVCaptureSessionPreset1920x1080 AVCaptureSessionPreset
	// AVCaptureSessionPreset320x240 is a preset suitable for capturing 320 x 240 pixel video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/qvga320x240
	AVCaptureSessionPreset320x240 AVCaptureSessionPreset
	// AVCaptureSessionPreset352x288 is a preset suitable for capturing CIF quality (352 x 288 pixel) video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/cif352x288
	AVCaptureSessionPreset352x288 AVCaptureSessionPreset
	// AVCaptureSessionPreset3840x2160 is a preset suitable for capturing 2160p-quality (3840 x 2160 pixels) video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/hd4K3840x2160
	AVCaptureSessionPreset3840x2160 AVCaptureSessionPreset
	// AVCaptureSessionPreset640x480 is a preset suitable for capturing VGA quality (640 x 480 pixel) video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/vga640x480
	AVCaptureSessionPreset640x480 AVCaptureSessionPreset
	// AVCaptureSessionPreset960x540 is a preset suitable for capturing quarter HD quality (960 x 540 pixel) video output.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/qHD960x540
	AVCaptureSessionPreset960x540 AVCaptureSessionPreset
	// AVCaptureSessionPresetiFrame1280x720 is a preset suitable for capturing 1280 x 720 quality iFrame H.264 video at about 40 Mbits/sec with AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/iFrame1280x720
	AVCaptureSessionPresetiFrame1280x720 AVCaptureSessionPreset
	// AVCaptureSessionPresetiFrame960x540 is a preset suitable for capturing 960 x 540 quality iFrame H.264 video at about 30 Mbits/sec with AAC audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/iFrame960x540
	AVCaptureSessionPresetiFrame960x540 AVCaptureSessionPreset
)

var ()

var ()

var ()

var ()

var (
	// AVCoreAnimationBeginTimeAtZero is a value that sets an animation begin time to `0`.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCoreAnimationBeginTimeAtZero
	AVCoreAnimationBeginTimeAtZero float64
)

var (
	// AVFileType3GPP is the UTI for the 3GPP file format.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVFileType/mobile3GPP
	AVFileType3GPP AVFileType
	// AVFileType3GPP2 is the UTI for the 3GPP2 file format.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVFileType/mobile3GPP2
	AVFileType3GPP2 AVFileType
)

var ()

var ()

var ()

var ()

var (
	// AVMetadata3GPUserDataKeyAlbumAndTrack is a key that represents the text for the album and track titles.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyAlbumAndTrack
	AVMetadata3GPUserDataKeyAlbumAndTrack AVMetadataKey
	// AVMetadata3GPUserDataKeyAuthor is a key that represents the author of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyAuthor
	AVMetadata3GPUserDataKeyAuthor AVMetadataKey
	// AVMetadata3GPUserDataKeyCollection is a key that represents the collection name for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyCollection
	AVMetadata3GPUserDataKeyCollection AVMetadataKey
	// AVMetadata3GPUserDataKeyCopyright is a key that represents the copyright statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyCopyright
	AVMetadata3GPUserDataKeyCopyright AVMetadataKey
	// AVMetadata3GPUserDataKeyDescription is a key that represents the description for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyDescription
	AVMetadata3GPUserDataKeyDescription AVMetadataKey
	// AVMetadata3GPUserDataKeyGenre is a key that represents the genre of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyGenre
	AVMetadata3GPUserDataKeyGenre AVMetadataKey
	// AVMetadata3GPUserDataKeyKeywordList is a key that represents the list of keywords for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyKeywordList
	AVMetadata3GPUserDataKeyKeywordList AVMetadataKey
	// AVMetadata3GPUserDataKeyLocation is a key that represents the location information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyLocation
	AVMetadata3GPUserDataKeyLocation AVMetadataKey
	// AVMetadata3GPUserDataKeyMediaClassification is a key that represents the classification of the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyMediaClassification
	AVMetadata3GPUserDataKeyMediaClassification AVMetadataKey
	// AVMetadata3GPUserDataKeyMediaRating is a key that represents the rating of the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyMediaRating
	AVMetadata3GPUserDataKeyMediaRating AVMetadataKey
	// AVMetadata3GPUserDataKeyPerformer is a key that represents information about the performer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyPerformer
	AVMetadata3GPUserDataKeyPerformer AVMetadataKey
	// AVMetadata3GPUserDataKeyRecordingYear is a key that represents the recording year for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyRecordingYear
	AVMetadata3GPUserDataKeyRecordingYear AVMetadataKey
	// AVMetadata3GPUserDataKeyThumbnail is a key that represents the media thumbnail.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyThumbnail
	AVMetadata3GPUserDataKeyThumbnail AVMetadataKey
	// AVMetadata3GPUserDataKeyTitle is a key that represents the title for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyTitle
	AVMetadata3GPUserDataKeyTitle AVMetadataKey
	// AVMetadata3GPUserDataKeyUserRating is a key that represents the user rating.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/metadata3GPUserDataKeyUserRating
	AVMetadata3GPUserDataKeyUserRating AVMetadataKey
	// AVMetadataCommonKeyAccessibilityDescription is a key that represents the accessibility description for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyAccessibilityDescription
	AVMetadataCommonKeyAccessibilityDescription AVMetadataKey
	// AVMetadataCommonKeyAlbumName is a key that represents the name of the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyAlbumName
	AVMetadataCommonKeyAlbumName AVMetadataKey
	// AVMetadataCommonKeyArtist is a key that represents the name of the artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyArtist
	AVMetadataCommonKeyArtist AVMetadataKey
	// AVMetadataCommonKeyArtwork is a key that represents an image relating to the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyArtwork
	AVMetadataCommonKeyArtwork AVMetadataKey
	// AVMetadataCommonKeyAuthor is a key that represents the name of the author.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyAuthor
	AVMetadataCommonKeyAuthor AVMetadataKey
	// AVMetadataCommonKeyContributor is a key that represents the name of the contributor.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyContributor
	AVMetadataCommonKeyContributor AVMetadataKey
	// AVMetadataCommonKeyCopyrights is a key that represents the copyright statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyCopyrights
	AVMetadataCommonKeyCopyrights AVMetadataKey
	// AVMetadataCommonKeyCreationDate is a key that represents the date of the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyCreationDate
	AVMetadataCommonKeyCreationDate AVMetadataKey
	// AVMetadataCommonKeyCreator is a key that represents the name of the creator.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyCreator
	AVMetadataCommonKeyCreator AVMetadataKey
	// AVMetadataCommonKeyDescription is a key that represents the description of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyDescription
	AVMetadataCommonKeyDescription AVMetadataKey
	// AVMetadataCommonKeyFormat is a key that represents the file format of the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyFormat
	AVMetadataCommonKeyFormat AVMetadataKey
	// AVMetadataCommonKeyIdentifier is a key that represents the ID for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyIdentifier
	AVMetadataCommonKeyIdentifier AVMetadataKey
	// AVMetadataCommonKeyLanguage is a key that represents the language of the text or lyrics spoken or sung in the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyLanguage
	AVMetadataCommonKeyLanguage AVMetadataKey
	// AVMetadataCommonKeyLastModifiedDate is a key that represents the last modification date of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyLastModifiedDate
	AVMetadataCommonKeyLastModifiedDate AVMetadataKey
	// AVMetadataCommonKeyLocation is a key that represents the location information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyLocation
	AVMetadataCommonKeyLocation AVMetadataKey
	// AVMetadataCommonKeyMake is a key that represents the name of the camera maker.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyMake
	AVMetadataCommonKeyMake AVMetadataKey
	// AVMetadataCommonKeyModel is a key that represents the name of the camera model.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyModel
	AVMetadataCommonKeyModel AVMetadataKey
	// AVMetadataCommonKeyPublisher is a key that represents the name of the publisher.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyPublisher
	AVMetadataCommonKeyPublisher AVMetadataKey
	// AVMetadataCommonKeyRelation is a key that represents the relation information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyRelation
	AVMetadataCommonKeyRelation AVMetadataKey
	// AVMetadataCommonKeySoftware is a key that represents the name of the software used to create the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeySoftware
	AVMetadataCommonKeySoftware AVMetadataKey
	// AVMetadataCommonKeySource is a key that represents the source information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeySource
	AVMetadataCommonKeySource AVMetadataKey
	// AVMetadataCommonKeySubject is a key that represents the subject information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeySubject
	AVMetadataCommonKeySubject AVMetadataKey
	// AVMetadataCommonKeyTitle is a key that represents the title of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyTitle
	AVMetadataCommonKeyTitle AVMetadataKey
	// AVMetadataCommonKeyType is a key that represents the media type.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyType
	AVMetadataCommonKeyType AVMetadataKey
	// AVMetadataID3MetadataKeyAlbumSortOrder is a key that represents how to sort the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyAlbumSortOrder
	AVMetadataID3MetadataKeyAlbumSortOrder AVMetadataKey
	// AVMetadataID3MetadataKeyAlbumTitle is a key that represents the title of the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyAlbumTitle
	AVMetadataID3MetadataKeyAlbumTitle AVMetadataKey
	// AVMetadataID3MetadataKeyAttachedPicture is a key that represents an image relating to the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyAttachedPicture
	AVMetadataID3MetadataKeyAttachedPicture AVMetadataKey
	// AVMetadataID3MetadataKeyAudioEncryption is a key that represents the encryption details of the audio stream.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyAudioEncryption
	AVMetadataID3MetadataKeyAudioEncryption AVMetadataKey
	// AVMetadataID3MetadataKeyAudioSeekPointIndex is a key that represents the list of seek points within the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyAudioSeekPointIndex
	AVMetadataID3MetadataKeyAudioSeekPointIndex AVMetadataKey
	// AVMetadataID3MetadataKeyBand is a key that represents additional information about the performers in the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyBand
	AVMetadataID3MetadataKeyBand AVMetadataKey
	// AVMetadataID3MetadataKeyBeatsPerMinute is a key that represents the beats per minute of the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyBeatsPerMinute
	AVMetadataID3MetadataKeyBeatsPerMinute AVMetadataKey
	// AVMetadataID3MetadataKeyComments is a key that represents additional text information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyComments
	AVMetadataID3MetadataKeyComments AVMetadataKey
	// AVMetadataID3MetadataKeyCommercial is a key that represents the commercial details for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyCommercial
	AVMetadataID3MetadataKeyCommercial AVMetadataKey
	// AVMetadataID3MetadataKeyCommercialInformation is a key that represents the webpage containing purchasing information.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyCommercialInformation
	AVMetadataID3MetadataKeyCommercialInformation AVMetadataKey
	// AVMetadataID3MetadataKeyComposer is a key that represents the name of the composer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyComposer
	AVMetadataID3MetadataKeyComposer AVMetadataKey
	// AVMetadataID3MetadataKeyConductor is a key that represents the name of the conductor.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyConductor
	AVMetadataID3MetadataKeyConductor AVMetadataKey
	// AVMetadataID3MetadataKeyContentGroupDescription is a key that indicates the sound belongs to a larger category of sounds or music.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyContentGroupDescription
	AVMetadataID3MetadataKeyContentGroupDescription AVMetadataKey
	// AVMetadataID3MetadataKeyContentType is a key that represents the media content type.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyContentType
	AVMetadataID3MetadataKeyContentType AVMetadataKey
	// AVMetadataID3MetadataKeyCopyright is a key that represents the copyright statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyCopyright
	AVMetadataID3MetadataKeyCopyright AVMetadataKey
	// AVMetadataID3MetadataKeyCopyrightInformation is a key that represents the webpage describing the terms of use and ownership.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyCopyrightInformation
	AVMetadataID3MetadataKeyCopyrightInformation AVMetadataKey
	// AVMetadataID3MetadataKeyDate is a key that represents the date for the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyDate
	AVMetadataID3MetadataKeyDate AVMetadataKey
	// AVMetadataID3MetadataKeyEncodedBy is a key that represents the person or organization responsible for encoding the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEncodedBy
	AVMetadataID3MetadataKeyEncodedBy AVMetadataKey
	// AVMetadataID3MetadataKeyEncodedWith is a key that represents the software or hardware and settings used for encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEncodedWith
	AVMetadataID3MetadataKeyEncodedWith AVMetadataKey
	// AVMetadataID3MetadataKeyEncodingTime is a key that represents the encoding time of the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEncodingTime
	AVMetadataID3MetadataKeyEncodingTime AVMetadataKey
	// AVMetadataID3MetadataKeyEncryption is a key that represents the encryption method used.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEncryption
	AVMetadataID3MetadataKeyEncryption AVMetadataKey
	// AVMetadataID3MetadataKeyEqualization is a key that represents the equalization curve within the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEqualization
	AVMetadataID3MetadataKeyEqualization AVMetadataKey
	// AVMetadataID3MetadataKeyEqualization2 is a key that represents the equalization curve within the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEqualization2
	AVMetadataID3MetadataKeyEqualization2 AVMetadataKey
	// AVMetadataID3MetadataKeyEventTimingCodes is a key that represents the timing codes used for synchronization with key events in a song or sound.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyEventTimingCodes
	AVMetadataID3MetadataKeyEventTimingCodes AVMetadataKey
	// AVMetadataID3MetadataKeyFileOwner is a key that represents the name of the owner or licensee of the file and it’s contents.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyFileOwner
	AVMetadataID3MetadataKeyFileOwner AVMetadataKey
	// AVMetadataID3MetadataKeyFileType is a key that represents the file type of the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyFileType
	AVMetadataID3MetadataKeyFileType AVMetadataKey
	// AVMetadataID3MetadataKeyGeneralEncapsulatedObject is a key that represents the details of a file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyGeneralEncapsulatedObject
	AVMetadataID3MetadataKeyGeneralEncapsulatedObject AVMetadataKey
	// AVMetadataID3MetadataKeyGroupIdentifier is a key that represents the grouping of distinct frames.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyGroupIdentifier
	AVMetadataID3MetadataKeyGroupIdentifier AVMetadataKey
	// AVMetadataID3MetadataKeyInitialKey is a key that represents the musical key in which the sound starts.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyInitialKey
	AVMetadataID3MetadataKeyInitialKey AVMetadataKey
	// AVMetadataID3MetadataKeyInternationalStandardRecordingCode is a key that represents the international standard recording code.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyInternationalStandardRecordingCode
	AVMetadataID3MetadataKeyInternationalStandardRecordingCode AVMetadataKey
	// AVMetadataID3MetadataKeyInternetRadioStationName is a key that represents the name of the internet radio station streaming the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyInternetRadioStationName
	AVMetadataID3MetadataKeyInternetRadioStationName AVMetadataKey
	// AVMetadataID3MetadataKeyInternetRadioStationOwner is a key that represents the name of the owner of the internet radio station streaming the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyInternetRadioStationOwner
	AVMetadataID3MetadataKeyInternetRadioStationOwner AVMetadataKey
	// AVMetadataID3MetadataKeyInvolvedPeopleList_v23 is a key that represents the list of names of contributors to the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyInvolvedPeopleList_v23
	AVMetadataID3MetadataKeyInvolvedPeopleList_v23 AVMetadataKey
	// AVMetadataID3MetadataKeyInvolvedPeopleList_v24 is a key that represents the list of names of contributors to the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyInvolvedPeopleList_v24
	AVMetadataID3MetadataKeyInvolvedPeopleList_v24 AVMetadataKey
	// AVMetadataID3MetadataKeyLanguage is a key that represents the language of the text or lyrics spoken or sung in the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyLanguage
	AVMetadataID3MetadataKeyLanguage AVMetadataKey
	// AVMetadataID3MetadataKeyLeadPerformer is a key that represents the main artist of the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyLeadPerformer
	AVMetadataID3MetadataKeyLeadPerformer AVMetadataKey
	// AVMetadataID3MetadataKeyLength is a key that represents the length of the audio file in milliseconds.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyLength
	AVMetadataID3MetadataKeyLength AVMetadataKey
	// AVMetadataID3MetadataKeyLink is a key that represents the link information from an ID3 tag that might reside in another audio file or alone in a binary file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyLink
	AVMetadataID3MetadataKeyLink AVMetadataKey
	// AVMetadataID3MetadataKeyLyricist is a key that represents the writer(s) of the text or lyrics in the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyLyricist
	AVMetadataID3MetadataKeyLyricist AVMetadataKey
	// AVMetadataID3MetadataKeyMPEGLocationLookupTable is a key that represents the lookup table used to increase performance and accuracy of jumps within an MPEG audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyMPEGLocationLookupTable
	AVMetadataID3MetadataKeyMPEGLocationLookupTable AVMetadataKey
	// AVMetadataID3MetadataKeyMediaType is a key that represents which media the sound originated from.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyMediaType
	AVMetadataID3MetadataKeyMediaType AVMetadataKey
	// AVMetadataID3MetadataKeyModifiedBy is a key that represents the people behind a remix and similar interpretations of another existing piece.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyModifiedBy
	AVMetadataID3MetadataKeyModifiedBy AVMetadataKey
	// AVMetadataID3MetadataKeyMood is a key that represents the mood of the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyMood
	AVMetadataID3MetadataKeyMood AVMetadataKey
	// AVMetadataID3MetadataKeyMusicCDIdentifier is a key that represents the ID used to identify the CD in databases such as the Compact Disc Database.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyMusicCDIdentifier
	AVMetadataID3MetadataKeyMusicCDIdentifier AVMetadataKey
	// AVMetadataID3MetadataKeyMusicianCreditsList is a key that represents the mapping between an instrument and the musician that played it.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyMusicianCreditsList
	AVMetadataID3MetadataKeyMusicianCreditsList AVMetadataKey
	// AVMetadataID3MetadataKeyOfficialArtistWebpage is a key that represents the artist’s official webpage.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOfficialArtistWebpage
	AVMetadataID3MetadataKeyOfficialArtistWebpage AVMetadataKey
	// AVMetadataID3MetadataKeyOfficialAudioFileWebpage is a key that represents the official webpage for the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOfficialAudioFileWebpage
	AVMetadataID3MetadataKeyOfficialAudioFileWebpage AVMetadataKey
	// AVMetadataID3MetadataKeyOfficialAudioSourceWebpage is a key that represents the official webpage for the source of the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOfficialAudioSourceWebpage
	AVMetadataID3MetadataKeyOfficialAudioSourceWebpage AVMetadataKey
	// AVMetadataID3MetadataKeyOfficialInternetRadioStationHomepage is a key that represents the official homepage of the internet radio station.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOfficialInternetRadioStationHomepage
	AVMetadataID3MetadataKeyOfficialInternetRadioStationHomepage AVMetadataKey
	// AVMetadataID3MetadataKeyOfficialPublisherWebpage is a key that represents the official webpage for the publisher.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOfficialPublisherWebpage
	AVMetadataID3MetadataKeyOfficialPublisherWebpage AVMetadataKey
	// AVMetadataID3MetadataKeyOriginalAlbumTitle is a key that represents the title of the original recording or source of sound.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOriginalAlbumTitle
	AVMetadataID3MetadataKeyOriginalAlbumTitle AVMetadataKey
	// AVMetadataID3MetadataKeyOriginalArtist is a key that represents the performer(s) of the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOriginalArtist
	AVMetadataID3MetadataKeyOriginalArtist AVMetadataKey
	// AVMetadataID3MetadataKeyOriginalFilename is a key that represents the original filename for the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOriginalFilename
	AVMetadataID3MetadataKeyOriginalFilename AVMetadataKey
	// AVMetadataID3MetadataKeyOriginalLyricist is a key that represents the text writer(s) of the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOriginalLyricist
	AVMetadataID3MetadataKeyOriginalLyricist AVMetadataKey
	// AVMetadataID3MetadataKeyOriginalReleaseTime is a key that represents the release time for the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOriginalReleaseTime
	AVMetadataID3MetadataKeyOriginalReleaseTime AVMetadataKey
	// AVMetadataID3MetadataKeyOriginalReleaseYear is a key that represents the release year for the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOriginalReleaseYear
	AVMetadataID3MetadataKeyOriginalReleaseYear AVMetadataKey
	// AVMetadataID3MetadataKeyOwnership is a key that represents the transaction details indicating proof of ownership if signed.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyOwnership
	AVMetadataID3MetadataKeyOwnership AVMetadataKey
	// AVMetadataID3MetadataKeyPartOfASet is a key that represents the part of a set the audio came from.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPartOfASet
	AVMetadataID3MetadataKeyPartOfASet AVMetadataKey
	// AVMetadataID3MetadataKeyPayment is a key that represents the webpage that handles the process of paying for the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPayment
	AVMetadataID3MetadataKeyPayment AVMetadataKey
	// AVMetadataID3MetadataKeyPerformerSortOrder is a key that represents the performer sort order.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPerformerSortOrder
	AVMetadataID3MetadataKeyPerformerSortOrder AVMetadataKey
	// AVMetadataID3MetadataKeyPlayCounter is a key that represents the play count of the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPlayCounter
	AVMetadataID3MetadataKeyPlayCounter AVMetadataKey
	// AVMetadataID3MetadataKeyPlaylistDelay is a key that represents the number of milliseconds of silence between every song in a playlist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPlaylistDelay
	AVMetadataID3MetadataKeyPlaylistDelay AVMetadataKey
	// AVMetadataID3MetadataKeyPopularimeter is a key that represents the rating for the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPopularimeter
	AVMetadataID3MetadataKeyPopularimeter AVMetadataKey
	// AVMetadataID3MetadataKeyPositionSynchronization is a key that represents the time offset of the first frame in the stream.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPositionSynchronization
	AVMetadataID3MetadataKeyPositionSynchronization AVMetadataKey
	// AVMetadataID3MetadataKeyPrivate is a key that represents the information from a software producer that its program uses.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPrivate
	AVMetadataID3MetadataKeyPrivate AVMetadataKey
	// AVMetadataID3MetadataKeyProducedNotice is a key that represents the produced notice.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyProducedNotice
	AVMetadataID3MetadataKeyProducedNotice AVMetadataKey
	// AVMetadataID3MetadataKeyPublisher is a key that represents the name of the label or publisher.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyPublisher
	AVMetadataID3MetadataKeyPublisher AVMetadataKey
	// AVMetadataID3MetadataKeyRecommendedBufferSize is a key that represents the buffer size the server recommends.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyRecommendedBufferSize
	AVMetadataID3MetadataKeyRecommendedBufferSize AVMetadataKey
	// AVMetadataID3MetadataKeyRecordingDates is a key that represents additional recording dates that complement year, date, and time keys.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyRecordingDates
	AVMetadataID3MetadataKeyRecordingDates AVMetadataKey
	// AVMetadataID3MetadataKeyRecordingTime is a key that represents the recording time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyRecordingTime
	AVMetadataID3MetadataKeyRecordingTime AVMetadataKey
	// AVMetadataID3MetadataKeyRelativeVolumeAdjustment is a key that represents the increase or decrease of volume on each channel while the file plays.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyRelativeVolumeAdjustment
	AVMetadataID3MetadataKeyRelativeVolumeAdjustment AVMetadataKey
	// AVMetadataID3MetadataKeyRelativeVolumeAdjustment2 is a key that represents the increase or decrease of volume on each channel while the file plays.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyRelativeVolumeAdjustment2
	AVMetadataID3MetadataKeyRelativeVolumeAdjustment2 AVMetadataKey
	// AVMetadataID3MetadataKeyReleaseTime is a key that represents the time of the first release.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyReleaseTime
	AVMetadataID3MetadataKeyReleaseTime AVMetadataKey
	// AVMetadataID3MetadataKeyReverb is a key that represents the adjustments to echoes of different kinds.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyReverb
	AVMetadataID3MetadataKeyReverb AVMetadataKey
	// AVMetadataID3MetadataKeySeek is a key that represents the location of other tags in a file or stream.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySeek
	AVMetadataID3MetadataKeySeek AVMetadataKey
	// AVMetadataID3MetadataKeySetSubtitle is a key that represents the set subtitle the track belongs to.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySetSubtitle
	AVMetadataID3MetadataKeySetSubtitle AVMetadataKey
	// AVMetadataID3MetadataKeySignature is a key that represents the group of frames to sign.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySignature
	AVMetadataID3MetadataKeySignature AVMetadataKey
	// AVMetadataID3MetadataKeySize is a key that represents the size of the audio file in bytes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySize
	AVMetadataID3MetadataKeySize AVMetadataKey
	// AVMetadataID3MetadataKeySubTitle is a key that represents the information relating to the contents title.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySubTitle
	AVMetadataID3MetadataKeySubTitle AVMetadataKey
	// AVMetadataID3MetadataKeySynchronizedLyric is a key that represents the words in the audio file as text in sync with the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySynchronizedLyric
	AVMetadataID3MetadataKeySynchronizedLyric AVMetadataKey
	// AVMetadataID3MetadataKeySynchronizedTempoCodes is a key that represents the tempo codes used for a more accurate description of the tempo of a musical piece.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeySynchronizedTempoCodes
	AVMetadataID3MetadataKeySynchronizedTempoCodes AVMetadataKey
	// AVMetadataID3MetadataKeyTaggingTime is a key that represents the time of tagging.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyTaggingTime
	AVMetadataID3MetadataKeyTaggingTime AVMetadataKey
	// AVMetadataID3MetadataKeyTermsOfUse is a key that represents the brief description of the terms of use and ownership of the file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyTermsOfUse
	AVMetadataID3MetadataKeyTermsOfUse AVMetadataKey
	// AVMetadataID3MetadataKeyTime is a key that represents the time for the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyTime
	AVMetadataID3MetadataKeyTime AVMetadataKey
	// AVMetadataID3MetadataKeyTitleDescription is a key that represents the name of the piece.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyTitleDescription
	AVMetadataID3MetadataKeyTitleDescription AVMetadataKey
	// AVMetadataID3MetadataKeyTitleSortOrder is a key that represents the title sort order.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyTitleSortOrder
	AVMetadataID3MetadataKeyTitleSortOrder AVMetadataKey
	// AVMetadataID3MetadataKeyTrackNumber is a key that represents the order number of the audio file.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyTrackNumber
	AVMetadataID3MetadataKeyTrackNumber AVMetadataKey
	// AVMetadataID3MetadataKeyUniqueFileIdentifier is a key that represents the identifier used to indicate the audio file in a database.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyUniqueFileIdentifier
	AVMetadataID3MetadataKeyUniqueFileIdentifier AVMetadataKey
	// AVMetadataID3MetadataKeyUnsynchronizedLyric is a key that represents the lyrics of the song or a text transcription of other vocal activities.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyUnsynchronizedLyric
	AVMetadataID3MetadataKeyUnsynchronizedLyric AVMetadataKey
	// AVMetadataID3MetadataKeyUserText is a key that represents the user text information frame.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyUserText
	AVMetadataID3MetadataKeyUserText AVMetadataKey
	// AVMetadataID3MetadataKeyUserURL is a key that represents the user webpage frame.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyUserURL
	AVMetadataID3MetadataKeyUserURL AVMetadataKey
	// AVMetadataID3MetadataKeyYear is a key that represents the year of the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/id3MetadataKeyYear
	AVMetadataID3MetadataKeyYear AVMetadataKey
	// AVMetadataISOUserDataKeyAccessibilityDescription is a key that represents the accessibility description for the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/isoUserDataKeyAccessibilityDescription
	AVMetadataISOUserDataKeyAccessibilityDescription AVMetadataKey
	// AVMetadataISOUserDataKeyCopyright is a key that represents the copyright statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/isoUserDataKeyCopyright
	AVMetadataISOUserDataKeyCopyright AVMetadataKey
	// AVMetadataISOUserDataKeyDate is a key that represents the date for the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/isoUserDataKeyDate
	AVMetadataISOUserDataKeyDate AVMetadataKey
	// AVMetadataISOUserDataKeyTaggedCharacteristic is a key that represents the tagged media characteristic used for identifying accessibility features.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/isoUserDataKeyTaggedCharacteristic
	AVMetadataISOUserDataKeyTaggedCharacteristic AVMetadataKey
	// AVMetadataIcyMetadataKeyStreamTitle is a key that represents the title of a stream.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/icyMetadataKeyStreamTitle
	AVMetadataIcyMetadataKeyStreamTitle AVMetadataKey
	// AVMetadataIcyMetadataKeyStreamURL is a key that represents the web address of a stream.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/icyMetadataKeyStreamURL
	AVMetadataIcyMetadataKeyStreamURL AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyAccessibilityDescription is a key that represents the accessibility description for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyAccessibilityDescription
	AVMetadataQuickTimeMetadataKeyAccessibilityDescription AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyAlbum is a key that represents the name of the album or collection in QuickTime.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyAlbum
	AVMetadataQuickTimeMetadataKeyAlbum AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyArranger is a key that represents the name of the arranger of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyArranger
	AVMetadataQuickTimeMetadataKeyArranger AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyArtist is a key that represents the name of the artist of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyArtist
	AVMetadataQuickTimeMetadataKeyArtist AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyArtwork is a key that represents an image relating to the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyArtwork
	AVMetadataQuickTimeMetadataKeyArtwork AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyAuthor is a key that represents the name of the author of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyAuthor
	AVMetadataQuickTimeMetadataKeyAuthor AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraFocalLength35mmEquivalent is a value of type kCMMetadataBaseDataType_UTF8 indicating focal length normalized to the 35mm film equivalent value (e.g. “50.00mm”).
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraFocalLength35mmEquivalent
	AVMetadataQuickTimeMetadataKeyCameraFocalLength35mmEquivalent AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraFrameReadoutTime is a key that represents the camera frame readout time in QuickTime.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraFrameReadoutTime
	AVMetadataQuickTimeMetadataKeyCameraFrameReadoutTime AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraISOSensitivity is a value of type kCMMetadataBaseDataType_UTF8 indicating the sensitivity of the camera to light in terms of ISO exposure index (e.g. “800”). See SMPTE RDD 18.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraISOSensitivity
	AVMetadataQuickTimeMetadataKeyCameraISOSensitivity AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraIdentifier is a key that represents the camera identifier in QuickTime.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraIdentifier
	AVMetadataQuickTimeMetadataKeyCameraIdentifier AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraLensIrisFNumber is a value of type kCMMetadataBaseDataType_UTF8 indicating measure of the amount of light transmitted through the lens. It is the focal length divided by the effective lens aperture diameter (e.g. “F2.8” or “2.8”).
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraLensIrisFNumber
	AVMetadataQuickTimeMetadataKeyCameraLensIrisFNumber AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraLensModel is a value of type kCMMetadataBaseDataType_UTF8 indicating the lens model (e.g. “iPhone 16 Pro back camera 6.765mm f/1.78”).
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraLensModel
	AVMetadataQuickTimeMetadataKeyCameraLensModel AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraShutterSpeedAngle is a value of type kCMMetadataBaseDataType_UTF8 indicating the exposure period expressed as an angle in minutes (1/60 degree) (e.g. “21600” or “360.00deg””).
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraShutterSpeedAngle
	AVMetadataQuickTimeMetadataKeyCameraShutterSpeedAngle AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraShutterSpeedTime is a value of type kCMMetadataBaseDataType_UTF8 indicating the exposure period expressed as a time per one frame/field period in seconds.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraShutterSpeedTime
	AVMetadataQuickTimeMetadataKeyCameraShutterSpeedTime AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCameraWhiteBalance is a value of type kCMMetadataBaseDataType_UTF8 indicating the white balance value defined by the temperature in Kelvin units (e.g. “5500K” or “5500”). See SMPTE RDD 18.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCameraWhiteBalance
	AVMetadataQuickTimeMetadataKeyCameraWhiteBalance AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCinematicVideoIntent is a value of type `kCMMetadataBaseDataType_UInt8` indicating whether this movie is intended as a Cinematic Video (1) or not (0).
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCinematicVideoIntent
	AVMetadataQuickTimeMetadataKeyCinematicVideoIntent AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCollectionUser is a key that represents a name that indicates a user-defined collection.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCollectionUser
	AVMetadataQuickTimeMetadataKeyCollectionUser AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyComment is a key that represents a comment regarding the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyComment
	AVMetadataQuickTimeMetadataKeyComment AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyComposer is a key that represents the name of the composer of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyComposer
	AVMetadataQuickTimeMetadataKeyComposer AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyContentIdentifier is a key that represents the content identifier in QuickTime.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyContentIdentifier
	AVMetadataQuickTimeMetadataKeyContentIdentifier AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCopyright is a key that represents the copyright statement for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCopyright
	AVMetadataQuickTimeMetadataKeyCopyright AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCreationDate is a key that represents the creation date of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCreationDate
	AVMetadataQuickTimeMetadataKeyCreationDate AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyCredits is a key that represents the credits of the movie source content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyCredits
	AVMetadataQuickTimeMetadataKeyCredits AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyDescription is a key that represents the description of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyDescription
	AVMetadataQuickTimeMetadataKeyDescription AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyDirectionFacing is a key that represents the direction the camera is facing during the shot.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyDirectionFacing
	AVMetadataQuickTimeMetadataKeyDirectionFacing AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyDirectionMotion is a key that represents the direction the camera is moving during the shot.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyDirectionMotion
	AVMetadataQuickTimeMetadataKeyDirectionMotion AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyDirector is a key that represents the name of the director of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyDirector
	AVMetadataQuickTimeMetadataKeyDirector AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyDisplayName is a key that represents the display name of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyDisplayName
	AVMetadataQuickTimeMetadataKeyDisplayName AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyEncodedBy is a key that represents the name of the person or organization responsible for encoding the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyEncodedBy
	AVMetadataQuickTimeMetadataKeyEncodedBy AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyFullFrameRatePlaybackIntent is an key that represents whether this movie should play at full frame rate.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyFullFrameRatePlaybackIntent
	AVMetadataQuickTimeMetadataKeyFullFrameRatePlaybackIntent AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyGenre is a key that represents the genre or genres to which the movie content conforms.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyGenre
	AVMetadataQuickTimeMetadataKeyGenre AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyInformation is a key that represents general information about the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyInformation
	AVMetadataQuickTimeMetadataKeyInformation AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyIsMontage is a key that represents that the movie is a montage of other media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyIsMontage
	AVMetadataQuickTimeMetadataKeyIsMontage AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyKeywords is a key that represents the keywords for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyKeywords
	AVMetadataQuickTimeMetadataKeyKeywords AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyLocationBody is a key that represents the astronomical body for compatibility with the 3GPP format.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyLocationBody
	AVMetadataQuickTimeMetadataKeyLocationBody AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyLocationDate is a key that represents a date and time using the extended format defined in ISO 8601:2004.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyLocationDate
	AVMetadataQuickTimeMetadataKeyLocationDate AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyLocationISO6709 is a key that represents the geographic point location by coordinates as defined in ISO 6709:2008.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyLocationISO6709
	AVMetadataQuickTimeMetadataKeyLocationISO6709 AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyLocationName is a key that represents the name of the location.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyLocationName
	AVMetadataQuickTimeMetadataKeyLocationName AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyLocationNote is a key that represents a descriptive comment about the location.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyLocationNote
	AVMetadataQuickTimeMetadataKeyLocationNote AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyLocationRole is a key that represents the single byte describing the movie location.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyLocationRole
	AVMetadataQuickTimeMetadataKeyLocationRole AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyMake is a key that represents the name of the camera maker.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyMake
	AVMetadataQuickTimeMetadataKeyMake AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyModel is a key that represents the name of the camera model.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyModel
	AVMetadataQuickTimeMetadataKeyModel AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyOriginalArtist is a key that represents the name of the original artist of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyOriginalArtist
	AVMetadataQuickTimeMetadataKeyOriginalArtist AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyPerformer is a key that represents the name of the performer in the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyPerformer
	AVMetadataQuickTimeMetadataKeyPerformer AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyPhonogramRights is a key that represents the phonogram rights statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyPhonogramRights
	AVMetadataQuickTimeMetadataKeyPhonogramRights AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyProducer is a key that represents the name of the producer of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyProducer
	AVMetadataQuickTimeMetadataKeyProducer AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyPublisher is a key that represents the name of the publisher of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyPublisher
	AVMetadataQuickTimeMetadataKeyPublisher AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyRatingUser is a key that represents the rating or relative value of the movie.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyRatingUser
	AVMetadataQuickTimeMetadataKeyRatingUser AVMetadataKey
	// AVMetadataQuickTimeMetadataKeySoftware is a key that represents the name of software used to create the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeySoftware
	AVMetadataQuickTimeMetadataKeySoftware AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyTitle is a key that represents the title of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyTitle
	AVMetadataQuickTimeMetadataKeyTitle AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTColorMatrices is a value of type kCMMetadataBaseDataType_RawData indicating the reference color translation matrix data for ProRes RAW.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyWhiteBalanceByCCTColorMatrices
	AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTColorMatrices AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTWhiteBalanceFactors is a value of type kCMMetadataBaseDataType_RawData indicating the reference white balance multiplication factor data for ProRes RAW.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyWhiteBalanceByCCTWhiteBalanceFactors
	AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTWhiteBalanceFactors AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyYear is a key that represents the recording year of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyYear
	AVMetadataQuickTimeMetadataKeyYear AVMetadataKey
	// AVMetadataQuickTimeMetadataKeyiXML is a key that represents iXML information for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyiXML
	AVMetadataQuickTimeMetadataKeyiXML AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyAccessibilityDescription is a key that represents the accessibility description for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyAccessibilityDescription
	AVMetadataQuickTimeUserDataKeyAccessibilityDescription AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyAlbum is a key that represents the name of the album or collection in QuickTime.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyAlbum
	AVMetadataQuickTimeUserDataKeyAlbum AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyArranger is a key that represents the name of the arranger of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyArranger
	AVMetadataQuickTimeUserDataKeyArranger AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyArtist is a key that represents the name of the artist of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyArtist
	AVMetadataQuickTimeUserDataKeyArtist AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyAuthor is a key that represents the name of the author of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyAuthor
	AVMetadataQuickTimeUserDataKeyAuthor AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyChapter is a key that represents the name of the chapter.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyChapter
	AVMetadataQuickTimeUserDataKeyChapter AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyComment is a key that represents a comment regarding the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyComment
	AVMetadataQuickTimeUserDataKeyComment AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyComposer is a key that represents the name of the composer of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyComposer
	AVMetadataQuickTimeUserDataKeyComposer AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyCopyright is a key that represents the copyright statement in QuickTime.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyCopyright
	AVMetadataQuickTimeUserDataKeyCopyright AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyCreationDate is a key that represents the creation date of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyCreationDate
	AVMetadataQuickTimeUserDataKeyCreationDate AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyCredits is a key that represents the credits of movie source content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyCredits
	AVMetadataQuickTimeUserDataKeyCredits AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyDescription is a key that represents the description of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyDescription
	AVMetadataQuickTimeUserDataKeyDescription AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyDirector is a key that represents the name of the director of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyDirector
	AVMetadataQuickTimeUserDataKeyDirector AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyDisclaimer is a key that represents the disclaimer regarding the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyDisclaimer
	AVMetadataQuickTimeUserDataKeyDisclaimer AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyEncodedBy is a key that represents the name of the person or organization responsible for encoding the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyEncodedBy
	AVMetadataQuickTimeUserDataKeyEncodedBy AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyFullName is a key that represents the full name of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyFullName
	AVMetadataQuickTimeUserDataKeyFullName AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyGenre is a key that represents the genre or genres to which the movie content conforms.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyGenre
	AVMetadataQuickTimeUserDataKeyGenre AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyHostComputer is a key that represents the name of the host computer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyHostComputer
	AVMetadataQuickTimeUserDataKeyHostComputer AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyInformation is a key that represents general information about the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyInformation
	AVMetadataQuickTimeUserDataKeyInformation AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyKeywords is a key that represents the keywords for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyKeywords
	AVMetadataQuickTimeUserDataKeyKeywords AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyLocationISO6709 is a key that represents the geographic point location by coordinates as defined in ISO 6709:2008.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyLocationISO6709
	AVMetadataQuickTimeUserDataKeyLocationISO6709 AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyMake is a key that represents the name of the camera maker.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyMake
	AVMetadataQuickTimeUserDataKeyMake AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyModel is a key that represents the name of the camera model.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyModel
	AVMetadataQuickTimeUserDataKeyModel AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyOriginalArtist is a key that represents the name of the original artist of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyOriginalArtist
	AVMetadataQuickTimeUserDataKeyOriginalArtist AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyOriginalFormat is a key that represents the original format of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyOriginalFormat
	AVMetadataQuickTimeUserDataKeyOriginalFormat AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyOriginalSource is a key that represents the original source of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyOriginalSource
	AVMetadataQuickTimeUserDataKeyOriginalSource AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyPerformers is a key that represents the name of the performers in the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyPerformers
	AVMetadataQuickTimeUserDataKeyPerformers AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyPhonogramRights is a key that represents the phonogram rights statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyPhonogramRights
	AVMetadataQuickTimeUserDataKeyPhonogramRights AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyProducer is a key that represents the name of the producer of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyProducer
	AVMetadataQuickTimeUserDataKeyProducer AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyProduct is a key that represents the name of the product.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyProduct
	AVMetadataQuickTimeUserDataKeyProduct AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyPublisher is a key that represents the name of the publisher of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyPublisher
	AVMetadataQuickTimeUserDataKeyPublisher AVMetadataKey
	// AVMetadataQuickTimeUserDataKeySoftware is a key that represents the name of software used to create the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeySoftware
	AVMetadataQuickTimeUserDataKeySoftware AVMetadataKey
	// AVMetadataQuickTimeUserDataKeySpecialPlaybackRequirements is a key that represents the special hardware and software requirements for playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeySpecialPlaybackRequirements
	AVMetadataQuickTimeUserDataKeySpecialPlaybackRequirements AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyTaggedCharacteristic is a key that represents the tagged characteristic.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyTaggedCharacteristic
	AVMetadataQuickTimeUserDataKeyTaggedCharacteristic AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyTrack is a key that represents a track in the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyTrack
	AVMetadataQuickTimeUserDataKeyTrack AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyTrackName is a key that represents the name of a track in the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyTrackName
	AVMetadataQuickTimeUserDataKeyTrackName AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyURLLink is a key that represents the webpage for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyURLLink
	AVMetadataQuickTimeUserDataKeyURLLink AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyWarning is a key that represents the warning text for the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyWarning
	AVMetadataQuickTimeUserDataKeyWarning AVMetadataKey
	// AVMetadataQuickTimeUserDataKeyWriter is a key that represents the name of the writer of the movie file content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeUserDataKeyWriter
	AVMetadataQuickTimeUserDataKeyWriter AVMetadataKey
	// AVMetadataiTunesMetadataKeyAccountKind is a key that represents the kind of iTunes account.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyAccountKind
	AVMetadataiTunesMetadataKeyAccountKind AVMetadataKey
	// AVMetadataiTunesMetadataKeyAcknowledgement is a key that represents the acknowledgement information in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyAcknowledgement
	AVMetadataiTunesMetadataKeyAcknowledgement AVMetadataKey
	// AVMetadataiTunesMetadataKeyAlbum is a key that represents the name of the album in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyAlbum
	AVMetadataiTunesMetadataKeyAlbum AVMetadataKey
	// AVMetadataiTunesMetadataKeyAlbumArtist is a key that represents the artist for the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyAlbumArtist
	AVMetadataiTunesMetadataKeyAlbumArtist AVMetadataKey
	// AVMetadataiTunesMetadataKeyAppleID is a key that represents an Apple ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyAppleID
	AVMetadataiTunesMetadataKeyAppleID AVMetadataKey
	// AVMetadataiTunesMetadataKeyArranger is a key that represents the name of the arranger.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyArranger
	AVMetadataiTunesMetadataKeyArranger AVMetadataKey
	// AVMetadataiTunesMetadataKeyArtDirector is a key that represents the name of the art director.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyArtDirector
	AVMetadataiTunesMetadataKeyArtDirector AVMetadataKey
	// AVMetadataiTunesMetadataKeyArtist is a key that represents the name of the artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyArtist
	AVMetadataiTunesMetadataKeyArtist AVMetadataKey
	// AVMetadataiTunesMetadataKeyArtistID is a key that represents the ID for an artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyArtistID
	AVMetadataiTunesMetadataKeyArtistID AVMetadataKey
	// AVMetadataiTunesMetadataKeyAuthor is a key that represents the name of the author.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyAuthor
	AVMetadataiTunesMetadataKeyAuthor AVMetadataKey
	// AVMetadataiTunesMetadataKeyBeatsPerMin is a key that represents the beats per minute of a track in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyBeatsPerMin
	AVMetadataiTunesMetadataKeyBeatsPerMin AVMetadataKey
	// AVMetadataiTunesMetadataKeyComposer is a key that represents the name of the composer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyComposer
	AVMetadataiTunesMetadataKeyComposer AVMetadataKey
	// AVMetadataiTunesMetadataKeyConductor is a key that represents the name of the conductor.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyConductor
	AVMetadataiTunesMetadataKeyConductor AVMetadataKey
	// AVMetadataiTunesMetadataKeyContentRating is a key that represents the content rating in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyContentRating
	AVMetadataiTunesMetadataKeyContentRating AVMetadataKey
	// AVMetadataiTunesMetadataKeyCopyright is a key that represents the copyright statement in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyCopyright
	AVMetadataiTunesMetadataKeyCopyright AVMetadataKey
	// AVMetadataiTunesMetadataKeyCoverArt is a key that represents an album cover image.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyCoverArt
	AVMetadataiTunesMetadataKeyCoverArt AVMetadataKey
	// AVMetadataiTunesMetadataKeyCredits is a key that represents the credits for the source content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyCredits
	AVMetadataiTunesMetadataKeyCredits AVMetadataKey
	// AVMetadataiTunesMetadataKeyDescription is a key that represents the description information in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyDescription
	AVMetadataiTunesMetadataKeyDescription AVMetadataKey
	// AVMetadataiTunesMetadataKeyDirector is a key that represents the name of the director.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyDirector
	AVMetadataiTunesMetadataKeyDirector AVMetadataKey
	// AVMetadataiTunesMetadataKeyDiscCompilation is a key that represents whether an album is a compilation of tracks in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyDiscCompilation
	AVMetadataiTunesMetadataKeyDiscCompilation AVMetadataKey
	// AVMetadataiTunesMetadataKeyDiscNumber is a key that represents the disc number of a multi-CD album in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyDiscNumber
	AVMetadataiTunesMetadataKeyDiscNumber AVMetadataKey
	// AVMetadataiTunesMetadataKeyEQ is a key that represents the equalizer preset option in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyEQ
	AVMetadataiTunesMetadataKeyEQ AVMetadataKey
	// AVMetadataiTunesMetadataKeyEncodedBy is a key that represents the person or organization that encoded the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyEncodedBy
	AVMetadataiTunesMetadataKeyEncodedBy AVMetadataKey
	// AVMetadataiTunesMetadataKeyEncodingTool is a key that represents the software or hardware and settings used for encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyEncodingTool
	AVMetadataiTunesMetadataKeyEncodingTool AVMetadataKey
	// AVMetadataiTunesMetadataKeyExecProducer is a key that represents the name of the executive producer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyExecProducer
	AVMetadataiTunesMetadataKeyExecProducer AVMetadataKey
	// AVMetadataiTunesMetadataKeyGenreID is a key that represents the genre ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyGenreID
	AVMetadataiTunesMetadataKeyGenreID AVMetadataKey
	// AVMetadataiTunesMetadataKeyGrouping is a key that represents additional grouping information for an album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyGrouping
	AVMetadataiTunesMetadataKeyGrouping AVMetadataKey
	// AVMetadataiTunesMetadataKeyLinerNotes is a key that represents the digital booklet for an album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyLinerNotes
	AVMetadataiTunesMetadataKeyLinerNotes AVMetadataKey
	// AVMetadataiTunesMetadataKeyLyrics is a key that represents the lyrics in the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyLyrics
	AVMetadataiTunesMetadataKeyLyrics AVMetadataKey
	// AVMetadataiTunesMetadataKeyOnlineExtras is a key that represents the extra materials for an album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyOnlineExtras
	AVMetadataiTunesMetadataKeyOnlineExtras AVMetadataKey
	// AVMetadataiTunesMetadataKeyOriginalArtist is a key that represents the name of the original artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyOriginalArtist
	AVMetadataiTunesMetadataKeyOriginalArtist AVMetadataKey
	// AVMetadataiTunesMetadataKeyPerformer is a key that represents the name of the performer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyPerformer
	AVMetadataiTunesMetadataKeyPerformer AVMetadataKey
	// AVMetadataiTunesMetadataKeyPhonogramRights is a key that represents the phonogram rights statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyPhonogramRights
	AVMetadataiTunesMetadataKeyPhonogramRights AVMetadataKey
	// AVMetadataiTunesMetadataKeyPlaylistID is a key that represents the playlist ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyPlaylistID
	AVMetadataiTunesMetadataKeyPlaylistID AVMetadataKey
	// AVMetadataiTunesMetadataKeyPredefinedGenre is a key that represents the predefined genre.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyPredefinedGenre
	AVMetadataiTunesMetadataKeyPredefinedGenre AVMetadataKey
	// AVMetadataiTunesMetadataKeyProducer is a key that represents the name of the producer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyProducer
	AVMetadataiTunesMetadataKeyProducer AVMetadataKey
	// AVMetadataiTunesMetadataKeyPublisher is a key that represents the name of the publisher.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyPublisher
	AVMetadataiTunesMetadataKeyPublisher AVMetadataKey
	// AVMetadataiTunesMetadataKeyRecordCompany is a key that represents the name of the record company.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyRecordCompany
	AVMetadataiTunesMetadataKeyRecordCompany AVMetadataKey
	// AVMetadataiTunesMetadataKeyReleaseDate is a key that represents the release date for the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyReleaseDate
	AVMetadataiTunesMetadataKeyReleaseDate AVMetadataKey
	// AVMetadataiTunesMetadataKeySoloist is a key that represents the name of the soloist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeySoloist
	AVMetadataiTunesMetadataKeySoloist AVMetadataKey
	// AVMetadataiTunesMetadataKeySongID is a key that represents the song ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeySongID
	AVMetadataiTunesMetadataKeySongID AVMetadataKey
	// AVMetadataiTunesMetadataKeySongName is a key that represents the name of the song in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeySongName
	AVMetadataiTunesMetadataKeySongName AVMetadataKey
	// AVMetadataiTunesMetadataKeySoundEngineer is a key that represents the name of the sound engineer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeySoundEngineer
	AVMetadataiTunesMetadataKeySoundEngineer AVMetadataKey
	// AVMetadataiTunesMetadataKeyThanks is a key that represents the thanks statement in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyThanks
	AVMetadataiTunesMetadataKeyThanks AVMetadataKey
	// AVMetadataiTunesMetadataKeyTrackNumber is a key that represents the order number in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyTrackNumber
	AVMetadataiTunesMetadataKeyTrackNumber AVMetadataKey
	// AVMetadataiTunesMetadataKeyTrackSubTitle is a key that represents the information relating to the contents title.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyTrackSubTitle
	AVMetadataiTunesMetadataKeyTrackSubTitle AVMetadataKey
	// AVMetadataiTunesMetadataKeyUserComment is a key that represents a user comment regarding the content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyUserComment
	AVMetadataiTunesMetadataKeyUserComment AVMetadataKey
	// AVMetadataiTunesMetadataKeyUserGenre is a key that represents the genre set by a user in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyUserGenre
	AVMetadataiTunesMetadataKeyUserGenre AVMetadataKey
)

var (
	// AVMetadataCommonIdentifierAccessibilityDescription is an identifier that represents the accessibility description for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierAccessibilityDescription
	AVMetadataCommonIdentifierAccessibilityDescription AVMetadataIdentifier
	// AVMetadataCommonIdentifierAlbumName is an identifier that represents the name of the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierAlbumName
	AVMetadataCommonIdentifierAlbumName AVMetadataIdentifier
	// AVMetadataCommonIdentifierArtist is an identifier that represents the name of the artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierArtist
	AVMetadataCommonIdentifierArtist AVMetadataIdentifier
	// AVMetadataCommonIdentifierArtwork is an identifier that represents an image relating to the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierArtwork
	AVMetadataCommonIdentifierArtwork AVMetadataIdentifier
	// AVMetadataCommonIdentifierAssetIdentifier is an identifier that represents the asset ID for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierAssetIdentifier
	AVMetadataCommonIdentifierAssetIdentifier AVMetadataIdentifier
	// AVMetadataCommonIdentifierAuthor is an identifier that represents the name of the author.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierAuthor
	AVMetadataCommonIdentifierAuthor AVMetadataIdentifier
	// AVMetadataCommonIdentifierContributor is an identifier that represents the name of the contributor.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierContributor
	AVMetadataCommonIdentifierContributor AVMetadataIdentifier
	// AVMetadataCommonIdentifierCopyrights is an identifier that represents the copyright statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierCopyrights
	AVMetadataCommonIdentifierCopyrights AVMetadataIdentifier
	// AVMetadataCommonIdentifierCreationDate is an identifier that represents the date of the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierCreationDate
	AVMetadataCommonIdentifierCreationDate AVMetadataIdentifier
	// AVMetadataCommonIdentifierCreator is an identifier that represents the name of the creator.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierCreator
	AVMetadataCommonIdentifierCreator AVMetadataIdentifier
	// AVMetadataCommonIdentifierDescription is an identifier that represents the description of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierDescription
	AVMetadataCommonIdentifierDescription AVMetadataIdentifier
	// AVMetadataCommonIdentifierFormat is an identifier that represents the file format of the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierFormat
	AVMetadataCommonIdentifierFormat AVMetadataIdentifier
	// AVMetadataCommonIdentifierLanguage is an identifier that represents the language of the text or lyrics spoken or sung in the audio.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierLanguage
	AVMetadataCommonIdentifierLanguage AVMetadataIdentifier
	// AVMetadataCommonIdentifierLastModifiedDate is an identifier that represents the last modification date of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierLastModifiedDate
	AVMetadataCommonIdentifierLastModifiedDate AVMetadataIdentifier
	// AVMetadataCommonIdentifierLocation is an identifier that represents the location information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierLocation
	AVMetadataCommonIdentifierLocation AVMetadataIdentifier
	// AVMetadataCommonIdentifierMake is an identifier that represents the name of the camera maker.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierMake
	AVMetadataCommonIdentifierMake AVMetadataIdentifier
	// AVMetadataCommonIdentifierModel is an identifier that represents the name of the camera model.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierModel
	AVMetadataCommonIdentifierModel AVMetadataIdentifier
	// AVMetadataCommonIdentifierPublisher is an identifier that represents the name of the publisher.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierPublisher
	AVMetadataCommonIdentifierPublisher AVMetadataIdentifier
	// AVMetadataCommonIdentifierRelation is an identifier that represents the relation information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierRelation
	AVMetadataCommonIdentifierRelation AVMetadataIdentifier
	// AVMetadataCommonIdentifierSoftware is an identifier that represents the name of the software used to create the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierSoftware
	AVMetadataCommonIdentifierSoftware AVMetadataIdentifier
	// AVMetadataCommonIdentifierSource is an identifier that represents the source information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierSource
	AVMetadataCommonIdentifierSource AVMetadataIdentifier
	// AVMetadataCommonIdentifierSubject is an identifier that represents the subject information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierSubject
	AVMetadataCommonIdentifierSubject AVMetadataIdentifier
	// AVMetadataCommonIdentifierTitle is an identifier that represents the title of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierTitle
	AVMetadataCommonIdentifierTitle AVMetadataIdentifier
	// AVMetadataCommonIdentifierType is an identifier that represents the media type.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierType
	AVMetadataCommonIdentifierType AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataAlbumAndTrack is an identifier that represents the text for the album and track titles.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataAlbumAndTrack
	AVMetadataIdentifier3GPUserDataAlbumAndTrack AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataAuthor is an identifier that represents the author of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataAuthor
	AVMetadataIdentifier3GPUserDataAuthor AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataCollection is an identifier that represents the collection name for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataCollection
	AVMetadataIdentifier3GPUserDataCollection AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataCopyright is an identifier that represents the copyright statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataCopyright
	AVMetadataIdentifier3GPUserDataCopyright AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataDescription is an identifier that represents the description for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataDescription
	AVMetadataIdentifier3GPUserDataDescription AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataGenre is an identifier that represents the genre of the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataGenre
	AVMetadataIdentifier3GPUserDataGenre AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataKeywordList is an identifier that represents the list of keywords for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataKeywordList
	AVMetadataIdentifier3GPUserDataKeywordList AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataLocation is an identifier that represents the location information for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataLocation
	AVMetadataIdentifier3GPUserDataLocation AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataMediaClassification is an identifier that represents the classification of the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataMediaClassification
	AVMetadataIdentifier3GPUserDataMediaClassification AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataMediaRating is an identifier that represents the rating of the media content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataMediaRating
	AVMetadataIdentifier3GPUserDataMediaRating AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataPerformer is an identifier that represents information about the performer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataPerformer
	AVMetadataIdentifier3GPUserDataPerformer AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataRecordingYear is an identifier that represents the recording year for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataRecordingYear
	AVMetadataIdentifier3GPUserDataRecordingYear AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataThumbnail is an identifier that represents the media thumbnail.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataThumbnail
	AVMetadataIdentifier3GPUserDataThumbnail AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataTitle is an identifier that represents the title for the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataTitle
	AVMetadataIdentifier3GPUserDataTitle AVMetadataIdentifier
	// AVMetadataIdentifier3GPUserDataUserRating is an identifier that represents the user rating.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/identifier3GPUserDataUserRating
	AVMetadataIdentifier3GPUserDataUserRating AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataAccountKind is an identifier that represents the kind of iTunes account.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataAccountKind
	AVMetadataIdentifieriTunesMetadataAccountKind AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataAcknowledgement is an identifier that represents the acknowledgement information in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataAcknowledgement
	AVMetadataIdentifieriTunesMetadataAcknowledgement AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataAlbum is an identifier that represents the name of the album in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataAlbum
	AVMetadataIdentifieriTunesMetadataAlbum AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataAlbumArtist is an identifier that represents the artist for the album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataAlbumArtist
	AVMetadataIdentifieriTunesMetadataAlbumArtist AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataAppleID is an identifier that represents an Apple ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataAppleID
	AVMetadataIdentifieriTunesMetadataAppleID AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataArranger is an identifier that represents the name of the arranger.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataArranger
	AVMetadataIdentifieriTunesMetadataArranger AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataArtDirector is an identifier that represents the name of the art director.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataArtDirector
	AVMetadataIdentifieriTunesMetadataArtDirector AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataArtist is an identifier that represents the name of the artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataArtist
	AVMetadataIdentifieriTunesMetadataArtist AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataArtistID is an identifier that represents the ID for an artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataArtistID
	AVMetadataIdentifieriTunesMetadataArtistID AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataAuthor is an identifier that represents the name of the author.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataAuthor
	AVMetadataIdentifieriTunesMetadataAuthor AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataBeatsPerMin is an identifier that represents the beats per minute of a track in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataBeatsPerMin
	AVMetadataIdentifieriTunesMetadataBeatsPerMin AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataComposer is an identifier that represents the name of the composer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataComposer
	AVMetadataIdentifieriTunesMetadataComposer AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataConductor is an identifier that represents the name of the conductor.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataConductor
	AVMetadataIdentifieriTunesMetadataConductor AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataContentRating is an identifier that represents the content rating in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataContentRating
	AVMetadataIdentifieriTunesMetadataContentRating AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataCopyright is an identifier that represents the copyright statement in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataCopyright
	AVMetadataIdentifieriTunesMetadataCopyright AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataCoverArt is an identifier that represents an album cover image.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataCoverArt
	AVMetadataIdentifieriTunesMetadataCoverArt AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataCredits is an identifier that represents the credits for the source content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataCredits
	AVMetadataIdentifieriTunesMetadataCredits AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataDescription is an identifier that represents the description information in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataDescription
	AVMetadataIdentifieriTunesMetadataDescription AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataDirector is an identifier that represents the name of the director.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataDirector
	AVMetadataIdentifieriTunesMetadataDirector AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataDiscCompilation is an identifier that represents whether an album is a compilation of tracks in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataDiscCompilation
	AVMetadataIdentifieriTunesMetadataDiscCompilation AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataDiscNumber is an identifier that represents the disc number of a multi-CD album in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataDiscNumber
	AVMetadataIdentifieriTunesMetadataDiscNumber AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataEQ is an identifier that represents the equalizer preset option in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataEQ
	AVMetadataIdentifieriTunesMetadataEQ AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataEncodedBy is an identifier that represents the person or organization responsible for encoding the media.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataEncodedBy
	AVMetadataIdentifieriTunesMetadataEncodedBy AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataEncodingTool is an identifier that represents the software or hardware and settings used for encoding.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataEncodingTool
	AVMetadataIdentifieriTunesMetadataEncodingTool AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataExecProducer is an identifier that represents the name of the executive producer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataExecProducer
	AVMetadataIdentifieriTunesMetadataExecProducer AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataGenreID is an identifier that represents the genre ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataGenreID
	AVMetadataIdentifieriTunesMetadataGenreID AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataGrouping is an identifier that represents additional grouping information for an album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataGrouping
	AVMetadataIdentifieriTunesMetadataGrouping AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataLinerNotes is an identifier that represents the digital booklet for an album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataLinerNotes
	AVMetadataIdentifieriTunesMetadataLinerNotes AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataLyrics is an identifier that represents the lyrics in the recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataLyrics
	AVMetadataIdentifieriTunesMetadataLyrics AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataOnlineExtras is an identifier that represents the extra materials for an album.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataOnlineExtras
	AVMetadataIdentifieriTunesMetadataOnlineExtras AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataOriginalArtist is an identifier that represents the name of the original artist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataOriginalArtist
	AVMetadataIdentifieriTunesMetadataOriginalArtist AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataPerformer is an identifier that represents the name of the performer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataPerformer
	AVMetadataIdentifieriTunesMetadataPerformer AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataPhonogramRights is an identifier that represents the phonogram rights statement.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataPhonogramRights
	AVMetadataIdentifieriTunesMetadataPhonogramRights AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataPlaylistID is an identifier that represents the playlist ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataPlaylistID
	AVMetadataIdentifieriTunesMetadataPlaylistID AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataPredefinedGenre is an identifier that represents the predefined genre.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataPredefinedGenre
	AVMetadataIdentifieriTunesMetadataPredefinedGenre AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataProducer is an identifier that represents the name of the producer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataProducer
	AVMetadataIdentifieriTunesMetadataProducer AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataPublisher is an identifier that represents the name of the publisher.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataPublisher
	AVMetadataIdentifieriTunesMetadataPublisher AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataRecordCompany is an identifier that represents the name of the record company.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataRecordCompany
	AVMetadataIdentifieriTunesMetadataRecordCompany AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataReleaseDate is an identifier that represents the release date for the original recording.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataReleaseDate
	AVMetadataIdentifieriTunesMetadataReleaseDate AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataSoloist is an identifier that represents the name of the soloist.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataSoloist
	AVMetadataIdentifieriTunesMetadataSoloist AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataSongID is an identifier that represents the song ID.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataSongID
	AVMetadataIdentifieriTunesMetadataSongID AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataSongName is an identifier that represents the name of the song in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataSongName
	AVMetadataIdentifieriTunesMetadataSongName AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataSoundEngineer is an identifier that represents the name of the sound engineer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataSoundEngineer
	AVMetadataIdentifieriTunesMetadataSoundEngineer AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataThanks is an identifier that represents the thanks statement in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataThanks
	AVMetadataIdentifieriTunesMetadataThanks AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataTrackNumber is an identifier that represents the order number in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataTrackNumber
	AVMetadataIdentifieriTunesMetadataTrackNumber AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataTrackSubTitle is an identifier that represents the information relating to the contents title.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataTrackSubTitle
	AVMetadataIdentifieriTunesMetadataTrackSubTitle AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataUserComment is an identifier that represents a user comment regarding the content.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataUserComment
	AVMetadataIdentifieriTunesMetadataUserComment AVMetadataIdentifier
	// AVMetadataIdentifieriTunesMetadataUserGenre is an identifier that represents the genre set by a user in iTunes.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataUserGenre
	AVMetadataIdentifieriTunesMetadataUserGenre AVMetadataIdentifier
)

var (
	// AVMetadataExtraAttributeBaseURIKey is a key that identifies the base URI the item uses to resolve its related URIs.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataExtraAttributeKey/baseURI
	AVMetadataExtraAttributeBaseURIKey AVMetadataExtraAttributeKey
	// AVMetadataExtraAttributeInfoKey is a key that identifies more information about the item.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataExtraAttributeKey/info
	AVMetadataExtraAttributeInfoKey AVMetadataExtraAttributeKey
	// AVMetadataExtraAttributeValueURIKey is a key that identifies a resource to use as the item’s value.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataExtraAttributeKey/valueURI
	AVMetadataExtraAttributeValueURIKey AVMetadataExtraAttributeKey
)

var (
	// AVMetadataFormatiTunesMetadata is the iTunes metadata format.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFormat/iTunesMetadata
	AVMetadataFormatiTunesMetadata AVMetadataFormat
)

var (
	// AVMetadataKeySpaceiTunes is the iTunes key space.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKeySpace/iTunes
	AVMetadataKeySpaceiTunes AVMetadataKeySpace
)

var ()

var (
	// AVOutputSettingsPreset1280x720 is a preset for H.264 video at 1280 by 720 pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/preset1280x720
	AVOutputSettingsPreset1280x720 AVOutputSettingsPreset
	// AVOutputSettingsPreset1920x1080 is a preset for H.264 video at 1920 by 1080 pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/preset1920x1080
	AVOutputSettingsPreset1920x1080 AVOutputSettingsPreset
	// AVOutputSettingsPreset3840x2160 is a preset for H.264 video at 3840 by 2160 pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/preset3840x2160
	AVOutputSettingsPreset3840x2160 AVOutputSettingsPreset
	// AVOutputSettingsPreset640x480 is a preset for H.264 video at 640 by 480 pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/preset640x480
	AVOutputSettingsPreset640x480 AVOutputSettingsPreset
	// AVOutputSettingsPreset960x540 is a preset for H.264 video at 960 by 540 pixels.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset/preset960x540
	AVOutputSettingsPreset960x540 AVOutputSettingsPreset
)

var ()

var (
	// AVPlayerInterstitialEventJoinCue is a cue that indicates that playback occurs before starting primary playback, regardless of initial primary playback position.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Cue-swift.struct/joinCue
	AVPlayerInterstitialEventJoinCue AVPlayerInterstitialEventCue
	// AVPlayerInterstitialEventLeaveCue is a cue that indicates event playback occurs after primary playback ends without error, either at the end of the primary asset or at the client-specified forward playback end time.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Cue-swift.struct/leaveCue
	AVPlayerInterstitialEventLeaveCue AVPlayerInterstitialEventCue
	// AVPlayerInterstitialEventNoCue is a cue that indicates that playback starts at the interstitial event time or date.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Cue-swift.struct/noCue
	AVPlayerInterstitialEventNoCue AVPlayerInterstitialEventCue
)

var ()

var ()

var (
	// AVPlayerWaitingDuringInterstitialEventReason is the player is waiting for an interstitial event to complete.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/interstitialEvent
	AVPlayerWaitingDuringInterstitialEventReason AVPlayerWaitingReason
	// AVPlayerWaitingForCoordinatedPlaybackReason is the player is waiting for another participant in a coordinated playback session.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/waitingForCoordinatedPlayback
	AVPlayerWaitingForCoordinatedPlaybackReason AVPlayerWaitingReason
	// AVPlayerWaitingToMinimizeStallsReason is the player is waiting for appropriate playback conditions before starting playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/toMinimizeStalls
	AVPlayerWaitingToMinimizeStallsReason AVPlayerWaitingReason
	// AVPlayerWaitingWhileEvaluatingBufferingRateReason is the player is waiting because it’s monitoring the buffer’s fill rate to determine whether playback is likely to complete without interruptions.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/evaluatingBufferingRate
	AVPlayerWaitingWhileEvaluatingBufferingRateReason AVPlayerWaitingReason
	// AVPlayerWaitingWithNoItemToPlayReason is the player is waiting because there’s no item to play.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason/noItemToPlay
	AVPlayerWaitingWithNoItemToPlayReason AVPlayerWaitingReason
)

var ()

var ()

var ()

var ()

var ()

var ()

var ()

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetChapterMetadataGroupsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetChapterMetadataGroupsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetContainsFragmentsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetContainsFragmentsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetDownloadedAssetEvictionPriorityDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetDownloadedAssetEvictionPrioritys.Default = AVAssetDownloadedAssetEvictionPriority(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetDownloadedAssetEvictionPriorityImportant"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetDownloadedAssetEvictionPrioritys.Important = AVAssetDownloadedAssetEvictionPriority(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetDurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetDurationDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPreset1280x720"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPreset1280x720 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPreset1920x1080"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPreset1920x1080 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPreset3840x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPreset3840x2160 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPreset640x480"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPreset640x480 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPreset960x540"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPreset960x540 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4A"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4A = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4V1080pHD"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4V1080pHD = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4V480pSD"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4V480pSD = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4V720pHD"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4V720pHD = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4VAppleTV"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4VAppleTV = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4VCellular"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4VCellular = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4VWiFi"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4VWiFi = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleM4ViPod"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleM4ViPod = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleProRes422LPCM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleProRes422LPCM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetAppleProRes4444LPCM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetAppleProRes4444LPCM = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVC1920x1080"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVC1920x1080 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVC1920x1080WithAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVC1920x1080WithAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVC3840x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVC3840x2160 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVC3840x2160WithAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVC3840x2160WithAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVC4320x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVC4320x2160 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVC7680x4320"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVC7680x4320 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVCHighestQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVCHighestQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHEVCHighestQualityWithAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHEVCHighestQualityWithAlpha = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetHighestQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetHighestQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetLowQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetLowQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetMVHEVC1440x1440"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetMVHEVC1440x1440 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetMVHEVC4320x4320"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetMVHEVC4320x4320 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetMVHEVC7680x7680"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetMVHEVC7680x7680 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetMVHEVC960x960"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetMVHEVC960x960 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetMediumQuality"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetMediumQuality = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetExportPresetPassthrough"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetExportPresetPassthrough = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetImageGeneratorApertureModeCleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetImageGeneratorApertureModes.CleanAperture = AVAssetImageGeneratorApertureMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetImageGeneratorApertureModeEncodedPixels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetImageGeneratorApertureModes.EncodedPixels = AVAssetImageGeneratorApertureMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetImageGeneratorApertureModeProductionAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetImageGeneratorApertureModes.ProductionAperture = AVAssetImageGeneratorApertureMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetImageGeneratorDynamicRangePolicyForceSDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetImageGeneratorDynamicRangePolicys.ForceSDR = AVAssetImageGeneratorDynamicRangePolicy(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetImageGeneratorDynamicRangePolicyMatchSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetImageGeneratorDynamicRangePolicys.MatchSource = AVAssetImageGeneratorDynamicRangePolicy(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetMediaSelectionGroupsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetMediaSelectionGroupsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetPlaybackConfigurationOptionAppleImmersiveVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetPlaybackConfigurationOptions.AppleImmersiveVideo = AVAssetPlaybackConfigurationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetPlaybackConfigurationOptionNonRectilinearProjection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetPlaybackConfigurationOptions.NonRectilinearProjection = AVAssetPlaybackConfigurationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetPlaybackConfigurationOptionSpatialVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetPlaybackConfigurationOptions.SpatialVideo = AVAssetPlaybackConfigurationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetPlaybackConfigurationOptionStereoMultiviewVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetPlaybackConfigurationOptions.StereoMultiviewVideo = AVAssetPlaybackConfigurationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetPlaybackConfigurationOptionStereoVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetPlaybackConfigurationOptions.StereoVideo = AVAssetPlaybackConfigurationOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetResourceLoadingRequestStreamingContentKeyRequestRequiresPersistentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetTrackSegmentsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetTrackSegmentsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetTrackTimeRangeDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetTrackTimeRangeDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetTrackTrackAssociationsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetTrackTrackAssociationsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetWasDefragmentedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetWasDefragmentedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetWriterInputMediaDataLocationBeforeMainMediaDataNotInterleaved"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetWriterInputMediaDataLocations.BeforeMainMediaDataNotInterleaved = AVAssetWriterInputMediaDataLocation(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetWriterInputMediaDataLocationInterleavedWithMainMediaData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetWriterInputMediaDataLocations.InterleavedWithMainMediaData = AVAssetWriterInputMediaDataLocation(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAssetWriterInputMediaDataLocationSparselyInterleavedWithMainMediaData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAssetWriterInputMediaDataLocations.SparselyInterleavedWithMainMediaData = AVAssetWriterInputMediaDataLocation(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioTimePitchAlgorithmLowQualityZeroLatency"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioTimePitchAlgorithms.LowQualityZeroLatency = AVAudioTimePitchAlgorithm(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioTimePitchAlgorithmSpectral"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioTimePitchAlgorithms.Spectral = AVAudioTimePitchAlgorithm(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioTimePitchAlgorithmTimeDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioTimePitchAlgorithms.TimeDomain = AVAudioTimePitchAlgorithm(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVAudioTimePitchAlgorithmVarispeed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVAudioTimePitchAlgorithms.Varispeed = AVAudioTimePitchAlgorithm(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptionConversionAdjustmentTypeTimeRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptionConversionAdjustmentTypeTimeRange = AVCaptionConversionAdjustmentType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptionConversionWarningTypeExcessMediaData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptionConversionWarningTypeExcessMediaData = AVCaptionConversionWarningType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptionMediaSubTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptionMediaSubTypeKey = AVCaptionSettingsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptionMediaTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptionMediaTypeKey = AVCaptionSettingsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptionTimeCodeFrameDurationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptionTimeCodeFrameDurationKey = AVCaptionSettingsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptionUseDropFrameTimeCodeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptionUseDropFrameTimeCodeKey = AVCaptionSettingsKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureAspectRatio16x9"); err == nil && ptr != 0 {
		AVCaptureAspectRatio16x9 = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureAspectRatio1x1"); err == nil && ptr != 0 {
		AVCaptureAspectRatio1x1 = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureAspectRatio3x4"); err == nil && ptr != 0 {
		AVCaptureAspectRatio3x4 = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureAspectRatio4x3"); err == nil && ptr != 0 {
		AVCaptureAspectRatio4x3 = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureAspectRatio9x16"); err == nil && ptr != 0 {
		AVCaptureAspectRatio9x16 = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceSubjectAreaDidChangeNotification"); err == nil && ptr != 0 {
		AVCaptureDeviceSubjectAreaDidChangeNotification = *(*foundation.NSNotification)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInDualCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInDualCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInDualWideCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInDualWideCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInDuoCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInDuoCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInLiDARDepthCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInLiDARDepthCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInMicrophone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInMicrophone = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInTelephotoCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInTelephotoCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInTripleCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInTripleCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInTrueDepthCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInTrueDepthCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInUltraWideCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInUltraWideCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeBuiltInWideAngleCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.BuiltInWideAngleCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeContinuityCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.ContinuityCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeDeskViewCamera"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.DeskViewCamera = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeExternal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.External = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeExternalUnknown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.ExternalUnknown = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceTypeMicrophone"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceTypes.Microphone = AVCaptureDeviceType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceWasConnectedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceWasConnectedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureDeviceWasDisconnectedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureDeviceWasDisconnectedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureInputPortFormatDescriptionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureInputPortFormatDescriptionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureMaxAvailableTorchLevel"); err == nil && ptr != 0 {
		AVCaptureMaxAvailableTorchLevel = *(*float32)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeBalloons"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.Balloons = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeConfetti"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.Confetti = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeFireworks"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.Fireworks = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeHeart"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.Heart = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeLasers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.Lasers = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeRain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.Rain = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeThumbsDown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.ThumbsDown = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureReactionTypeThumbsUp"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureReactionTypes.ThumbsUp = AVCaptureReactionType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSceneMonitoringStatusNotEnoughLight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSceneMonitoringStatusNotEnoughLight = AVCaptureSceneMonitoringStatus(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionDidStartRunningNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionDidStartRunningNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionDidStopRunningNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionDidStopRunningNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionInterruptionEndedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionInterruptionEndedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset1280x720"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset1280x720 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset1920x1080"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset1920x1080 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset320x240"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset320x240 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset352x288"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset352x288 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset3840x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset3840x2160 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset640x480"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset640x480 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPreset960x540"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPreset960x540 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetHigh"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresets.High = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetInputPriority"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresets.InputPriority = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetLow"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresets.Low = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetMedium"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresets.Medium = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetPhoto"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresets.Photo = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetiFrame1280x720"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresetiFrame1280x720 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionPresetiFrame960x540"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionPresetiFrame960x540 = AVCaptureSessionPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionRuntimeErrorNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionRuntimeErrorNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSessionWasInterruptedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCaptureSessionWasInterruptedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSystemPressureLevelCritical"); err == nil && ptr != 0 {
		AVCaptureSystemPressureLevelCritical = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSystemPressureLevelFair"); err == nil && ptr != 0 {
		AVCaptureSystemPressureLevelFair = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSystemPressureLevelNominal"); err == nil && ptr != 0 {
		AVCaptureSystemPressureLevelNominal = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSystemPressureLevelSerious"); err == nil && ptr != 0 {
		AVCaptureSystemPressureLevelSerious = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCaptureSystemPressureLevelShutdown"); err == nil && ptr != 0 {
		AVCaptureSystemPressureLevelShutdown = *(*objc.ID)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestProtocolVersionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestProtocolVersionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestRandomDeviceIdentifierSeedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestRandomDeviceIdentifierSeedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestRequiresValidationDataInSecureTokenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestRequiresValidationDataInSecureTokenKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestRetryReasonReceivedObsoleteContentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestRetryReasons.ReceivedObsoleteContentKey = AVContentKeyRequestRetryReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestRetryReasonReceivedResponseWithExpiredLease"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestRetryReasons.ReceivedResponseWithExpiredLease = AVContentKeyRequestRetryReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestRetryReasonTimedOut"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestRetryReasons.TimedOut = AVContentKeyRequestRetryReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeyRequestShouldRandomizeDeviceIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeyRequestShouldRandomizeDeviceIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeySessionServerPlaybackContextOptionProtocolVersions"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeySessionServerPlaybackContextOptions.ProtocolVersions = AVContentKeySessionServerPlaybackContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeySessionServerPlaybackContextOptionServerChallenge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeySessionServerPlaybackContextOptions.ServerChallenge = AVContentKeySessionServerPlaybackContextOption(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeySystemAuthorizationToken"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeySystems.AuthorizationToken = AVContentKeySystem(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeySystemClearKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeySystems.ClearKey = AVContentKeySystem(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVContentKeySystemFairPlayStreaming"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVContentKeySystems.FairPlayStreaming = AVContentKeySystem(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoordinatedPlaybackSuspensionReasonAudioSessionInterrupted"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCoordinatedPlaybackSuspensionReasons.AudioSessionInterrupted = AVCoordinatedPlaybackSuspensionReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoordinatedPlaybackSuspensionReasonCoordinatedPlaybackNotPossible"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCoordinatedPlaybackSuspensionReasons.CoordinatedPlaybackNotPossible = AVCoordinatedPlaybackSuspensionReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoordinatedPlaybackSuspensionReasonPlayingInterstitial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCoordinatedPlaybackSuspensionReasons.PlayingInterstitial = AVCoordinatedPlaybackSuspensionReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoordinatedPlaybackSuspensionReasonStallRecovery"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCoordinatedPlaybackSuspensionReasons.StallRecovery = AVCoordinatedPlaybackSuspensionReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoordinatedPlaybackSuspensionReasonUserActionRequired"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCoordinatedPlaybackSuspensionReasons.UserActionRequired = AVCoordinatedPlaybackSuspensionReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoordinatedPlaybackSuspensionReasonUserIsChangingCurrentTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVCoordinatedPlaybackSuspensionReasons.UserIsChangingCurrentTime = AVCoordinatedPlaybackSuspensionReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVCoreAnimationBeginTimeAtZero"); err == nil && ptr != 0 {
		AVCoreAnimationBeginTimeAtZero = *(*float64)(unsafe.Pointer(ptr))
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorDeviceKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorDeviceKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorDiscontinuityFlagsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorDiscontinuityFlagsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorFileSizeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorFileSizeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorFileTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorFileTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorMediaSubTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorMediaSubTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorMediaTypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorMediaTypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorPIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorPIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorPersistentTrackIDKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorPersistentTrackIDKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorPresentationTimeStampKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorPresentationTimeStampKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorRecordingSuccessfullyFinishedKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorRecordingSuccessfullyFinishedKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVErrorTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVErrorTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileType3GPP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileType3GPP = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileType3GPP2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileType3GPP2 = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAC3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AC3 = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAHAP"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AHAP = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAIFC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AIFC = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAIFF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AIFF = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAMR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AMR = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAVCI"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AVCI = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAppleM4A"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AppleM4A = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAppleM4V"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AppleM4V = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeAppleiTT"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.AppleiTT = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeCoreAudioFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.CoreAudioFormat = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeDICOM"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.DICOM = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeDNG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.DNG = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeEnhancedAC3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.EnhancedAC3 = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeHEIC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.HEIC = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeHEIF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.HEIF = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeJPEG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.JPEG = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeMPEG4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.MPEG4 = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeMPEGLayer3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.MPEGLayer3 = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeProfileMPEG4AppleHLS"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypeProfiles.MPEG4AppleHLS = AVFileTypeProfile(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeProfileMPEG4CMAFCompliant"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypeProfiles.MPEG4CMAFCompliant = AVFileTypeProfile(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeQuickTimeAudio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.QuickTimeAudio = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeQuickTimeMovie"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.QuickTimeMovie = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeSCC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.SCC = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeSunAU"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.SunAU = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeTIFF"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.TIFF = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFileTypeWAVE"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFileTypes.WAVE = AVFileType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFoundationErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFoundationErrorDomain = foundation.NSErrorDomain(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFragmentedMovieContainsMovieFragmentsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFragmentedMovieContainsMovieFragmentsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFragmentedMovieDurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFragmentedMovieDurationDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFragmentedMovieTrackSegmentsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFragmentedMovieTrackSegmentsDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFragmentedMovieTrackTimeRangeDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFragmentedMovieTrackTimeRangeDidChangeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVFragmentedMovieWasDefragmentedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVFragmentedMovieWasDefragmentedNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLayerVideoGravityResize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLayerVideoGravitys.Resize = AVLayerVideoGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLayerVideoGravityResizeAspect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLayerVideoGravitys.ResizeAspect = AVLayerVideoGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVLayerVideoGravityResizeAspectFill"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVLayerVideoGravitys.ResizeAspectFill = AVLayerVideoGravity(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicAudible"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.Audible = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicCarriesVideoStereoMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.CarriesVideoStereoMetadata = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicContainsAlphaChannel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.ContainsAlphaChannel = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicContainsHDRVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.ContainsHDRVideo = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicContainsOnlyForcedSubtitles"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.ContainsOnlyForcedSubtitles = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicContainsStereoMultiviewVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.ContainsStereoMultiviewVideo = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicDescribesMusicAndSoundForAccessibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.DescribesMusicAndSoundForAccessibility = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicDescribesVideoForAccessibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.DescribesVideoForAccessibility = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicDubbedTranslation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.DubbedTranslation = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicEasyToRead"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.EasyToRead = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicEnhancesSpeechIntelligibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.EnhancesSpeechIntelligibility = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicFrameBased"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.FrameBased = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicIndicatesHorizontalFieldOfView"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.IndicatesHorizontalFieldOfView = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicIndicatesNonRectilinearProjection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.IndicatesNonRectilinearProjection = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicIsAuxiliaryContent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.IsAuxiliaryContent = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicIsMainProgramContent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.IsMainProgramContent = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicIsOriginalContent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.IsOriginalContent = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicLanguageTranslation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.LanguageTranslation = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicLegible"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.Legible = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicMachineGenerated"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.MachineGenerated = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicTactileMinimal"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.TactileMinimal = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicTranscribesSpokenDialogForAccessibility"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.TranscribesSpokenDialogForAccessibility = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicUsesWideGamutColorSpace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.UsesWideGamutColorSpace = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicVisual"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.Visual = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaCharacteristicVoiceOverTranslation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaCharacteristics.VoiceOverTranslation = AVMediaCharacteristic(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeAudio"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Audio = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeAuxiliaryPicture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.AuxiliaryPicture = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeClosedCaption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.ClosedCaption = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeDepthData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.DepthData = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeHaptic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Haptic = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Metadata = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeMetadataObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.MetadataObject = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeMuxed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Muxed = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeSubtitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Subtitle = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Text = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeTimecode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Timecode = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMediaTypeVideo"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMediaTypes.Video = AVMediaType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyAlbumAndTrack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyAlbumAndTrack = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyAuthor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyCollection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyCollection = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyCopyright = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyGenre = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyKeywordList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyKeywordList = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyLocation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyMediaClassification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyMediaClassification = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyMediaRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyMediaRating = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyPerformer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyRecordingYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyRecordingYear = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyThumbnail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyThumbnail = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadata3GPUserDataKeyUserRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadata3GPUserDataKeyUserRating = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierAccessibilityDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierAlbumName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierAlbumName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierArtwork"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierArtwork = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierAssetIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierAssetIdentifier = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierAuthor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierContributor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierContributor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierCopyrights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierCopyrights = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierCreationDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierCreator = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierFormat = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierLanguage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierLanguage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierLastModifiedDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierLastModifiedDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierLocation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierMake = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierModel = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierPublisher = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierRelation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierRelation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierSoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierSoftware = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierSource = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierSubject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierSubject = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonIdentifierType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonIdentifierType = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyAccessibilityDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyAlbumName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyAlbumName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyArtwork"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyArtwork = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyAuthor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyContributor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyContributor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyCopyrights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyCopyrights = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyCreationDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyCreator"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyCreator = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyFormat = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyIdentifier = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyLanguage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyLanguage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyLastModifiedDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyLastModifiedDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyLocation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyMake = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyModel = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyPublisher = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyRelation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyRelation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeySoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeySoftware = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeySource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeySource = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeySubject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeySubject = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataCommonKeyType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataCommonKeyType = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataExtraAttributeBaseURIKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataExtraAttributeBaseURIKey = AVMetadataExtraAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataExtraAttributeInfoKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataExtraAttributeInfoKey = AVMetadataExtraAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataExtraAttributeValueURIKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataExtraAttributeValueURIKey = AVMetadataExtraAttributeKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatHLSMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormats.HLSMetadata = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatID3Metadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormats.ID3Metadata = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatISOUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormats.ISOUserData = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatQuickTimeMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormats.QuickTimeMetadata = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatQuickTimeUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormats.QuickTimeUserData = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatUnknown"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormats.Unknown = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataFormatiTunesMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataFormatiTunesMetadata = AVMetadataFormat(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyAlbumSortOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyAlbumSortOrder = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyAlbumTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyAlbumTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyAttachedPicture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyAttachedPicture = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyAudioEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyAudioEncryption = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyAudioSeekPointIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyAudioSeekPointIndex = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyBand"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyBand = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyBeatsPerMinute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyBeatsPerMinute = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyComments"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyComments = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyCommercial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyCommercial = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyCommercialInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyCommercialInformation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyComposer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyConductor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyConductor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyContentGroupDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyContentGroupDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyContentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyContentType = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyCopyright = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyCopyrightInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyCopyrightInformation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEncodedBy = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEncodedWith"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEncodedWith = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEncodingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEncodingTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEncryption = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEqualization"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEqualization = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEqualization2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEqualization2 = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyEventTimingCodes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyEventTimingCodes = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyFileOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyFileOwner = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyFileType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyFileType = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyGeneralEncapsulatedObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyGeneralEncapsulatedObject = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyGroupIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyGroupIdentifier = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyInitialKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyInitialKey = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyInternationalStandardRecordingCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyInternationalStandardRecordingCode = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyInternetRadioStationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyInternetRadioStationName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyInternetRadioStationOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyInternetRadioStationOwner = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyInvolvedPeopleList_v23"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyInvolvedPeopleList_v23 = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyInvolvedPeopleList_v24"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyInvolvedPeopleList_v24 = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyLanguage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyLanguage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyLeadPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyLeadPerformer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyLength = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyLink = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyLyricist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyLyricist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyMPEGLocationLookupTable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyMPEGLocationLookupTable = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyMediaType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyMediaType = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyModifiedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyModifiedBy = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyMood"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyMood = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyMusicCDIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyMusicCDIdentifier = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyMusicianCreditsList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyMusicianCreditsList = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOfficialArtistWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOfficialArtistWebpage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOfficialAudioFileWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOfficialAudioFileWebpage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOfficialAudioSourceWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOfficialAudioSourceWebpage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOfficialInternetRadioStationHomepage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOfficialInternetRadioStationHomepage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOfficialPublisherWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOfficialPublisherWebpage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOriginalAlbumTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOriginalAlbumTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOriginalArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOriginalFilename"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOriginalFilename = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOriginalLyricist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOriginalLyricist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOriginalReleaseTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOriginalReleaseTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOriginalReleaseYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOriginalReleaseYear = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyOwnership"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyOwnership = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPartOfASet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPartOfASet = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPayment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPayment = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPerformerSortOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPerformerSortOrder = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPlayCounter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPlayCounter = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPlaylistDelay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPlaylistDelay = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPopularimeter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPopularimeter = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPositionSynchronization"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPositionSynchronization = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPrivate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPrivate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyProducedNotice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyProducedNotice = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyPublisher = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyRecommendedBufferSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyRecommendedBufferSize = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyRecordingDates"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyRecordingDates = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyRecordingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyRecordingTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyRelativeVolumeAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyRelativeVolumeAdjustment = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyRelativeVolumeAdjustment2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyRelativeVolumeAdjustment2 = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyReleaseTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyReleaseTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyReverb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyReverb = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySeek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySeek = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySetSubtitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySetSubtitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySignature = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySize = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySubTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySubTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySynchronizedLyric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySynchronizedLyric = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeySynchronizedTempoCodes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeySynchronizedTempoCodes = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyTaggingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyTaggingTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyTermsOfUse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyTermsOfUse = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyTitleDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyTitleDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyTitleSortOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyTitleSortOrder = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyTrackNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyTrackNumber = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyUniqueFileIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyUniqueFileIdentifier = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyUnsynchronizedLyric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyUnsynchronizedLyric = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyUserText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyUserText = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyUserURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyUserURL = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataID3MetadataKeyYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataID3MetadataKeyYear = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataISOUserDataKeyAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataISOUserDataKeyAccessibilityDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataISOUserDataKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataISOUserDataKeyCopyright = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataISOUserDataKeyDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataISOUserDataKeyDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataISOUserDataKeyTaggedCharacteristic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataISOUserDataKeyTaggedCharacteristic = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIcyMetadataKeyStreamTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIcyMetadataKeyStreamTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIcyMetadataKeyStreamURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIcyMetadataKeyStreamURL = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataAlbumAndTrack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataAlbumAndTrack = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataAuthor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataCollection"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataCollection = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataCopyright = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataGenre = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataKeywordList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataKeywordList = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataLocation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataLocation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataMediaClassification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataMediaClassification = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataMediaRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataMediaRating = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataPerformer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataRecordingYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataRecordingYear = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataThumbnail"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataThumbnail = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifier3GPUserDataUserRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifier3GPUserDataUserRating = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataAlbumSortOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataAlbumSortOrder = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataAlbumTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataAlbumTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataAttachedPicture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataAttachedPicture = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataAudioEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataAudioEncryption = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataAudioSeekPointIndex"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataAudioSeekPointIndex = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataBand"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataBand = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataBeatsPerMinute"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataBeatsPerMinute = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataComments"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataComments = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataCommercial"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataCommercial = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataCommercialInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataCommercialInformation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataComposer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataConductor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataConductor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataContentGroupDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataContentGroupDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataContentType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataContentType = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataCopyright = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataCopyrightInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataCopyrightInformation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEncodedBy = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEncodedWith"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEncodedWith = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEncodingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEncodingTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEncryption"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEncryption = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEqualization"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEqualization = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEqualization2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEqualization2 = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataEventTimingCodes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataEventTimingCodes = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataFileOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataFileOwner = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataFileType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataFileType = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataGeneralEncapsulatedObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataGeneralEncapsulatedObject = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataGroupIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataGroupIdentifier = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataInitialKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataInitialKey = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataInternationalStandardRecordingCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataInternationalStandardRecordingCode = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataInternetRadioStationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataInternetRadioStationName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataInternetRadioStationOwner"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataInternetRadioStationOwner = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataInvolvedPeopleList_v23"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataInvolvedPeopleList_v23 = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataInvolvedPeopleList_v24"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataInvolvedPeopleList_v24 = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataLanguage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataLanguage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataLeadPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataLeadPerformer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataLength"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataLength = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataLink = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataLyricist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataLyricist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataMPEGLocationLookupTable"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataMPEGLocationLookupTable = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataMediaType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataMediaType = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataModifiedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataModifiedBy = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataMood"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataMood = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataMusicCDIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataMusicCDIdentifier = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataMusicianCreditsList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataMusicianCreditsList = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOfficialArtistWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOfficialArtistWebpage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOfficialAudioFileWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOfficialAudioFileWebpage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOfficialAudioSourceWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOfficialAudioSourceWebpage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOfficialInternetRadioStationHomepage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOfficialInternetRadioStationHomepage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOfficialPublisherWebpage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOfficialPublisherWebpage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOriginalAlbumTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOriginalAlbumTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOriginalArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOriginalFilename"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOriginalFilename = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOriginalLyricist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOriginalLyricist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOriginalReleaseTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOriginalReleaseTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOriginalReleaseYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOriginalReleaseYear = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataOwnership"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataOwnership = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPartOfASet"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPartOfASet = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPayment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPayment = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPerformerSortOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPerformerSortOrder = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPlayCounter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPlayCounter = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPlaylistDelay"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPlaylistDelay = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPopularimeter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPopularimeter = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPositionSynchronization"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPositionSynchronization = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPrivate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPrivate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataProducedNotice"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataProducedNotice = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataPublisher = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataRecommendedBufferSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataRecommendedBufferSize = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataRecordingDates"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataRecordingDates = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataRecordingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataRecordingTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataRelativeVolumeAdjustment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataRelativeVolumeAdjustment = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataRelativeVolumeAdjustment2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataRelativeVolumeAdjustment2 = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataReleaseTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataReleaseTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataReverb"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataReverb = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSeek"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSeek = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSetSubtitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSetSubtitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSignature"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSignature = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSize = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSubTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSubTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSynchronizedLyric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSynchronizedLyric = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataSynchronizedTempoCodes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataSynchronizedTempoCodes = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataTaggingTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataTaggingTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataTermsOfUse"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataTermsOfUse = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataTitleDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataTitleDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataTitleSortOrder"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataTitleSortOrder = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataTrackNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataTrackNumber = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataUniqueFileIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataUniqueFileIdentifier = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataUnsynchronizedLyric"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataUnsynchronizedLyric = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataUserText"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataUserText = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataUserURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataUserURL = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierID3MetadataYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ID3MetadataYear = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierISOUserDataAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ISOUserDataAccessibilityDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierISOUserDataCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ISOUserDataCopyright = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierISOUserDataDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ISOUserDataDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierISOUserDataTaggedCharacteristic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.ISOUserDataTaggedCharacteristic = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierIcyMetadataStreamTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.IcyMetadataStreamTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierIcyMetadataStreamURL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.IcyMetadataStreamURL = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataAIMEData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataAIMEData = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataAccessibilityDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataAlbum = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataArranger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataArranger = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataArtwork"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataArtwork = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataAuthor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataAutoLivePhoto"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataAutoLivePhoto = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraFocalLength35mmEquivalent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraFocalLength35mmEquivalent = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraFrameReadoutTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraFrameReadoutTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraISOSensitivity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraISOSensitivity = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraIdentifier = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraLensIrisFNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraLensIrisFNumber = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraLensModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraLensModel = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraShutterSpeedAngle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraShutterSpeedAngle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraShutterSpeedTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraShutterSpeedTime = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCameraWhiteBalance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCameraWhiteBalance = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCinematicVideoIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCinematicVideoIntent = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCollectionUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCollectionUser = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataComment = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataComposer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataContentIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataContentIdentifier = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCopyright = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCreationDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataCredits = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDetectedCatBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDetectedCatBody = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDetectedDogBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDetectedDogBody = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDetectedFace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDetectedFace = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDetectedHumanBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDetectedHumanBody = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDetectedSalientObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDetectedSalientObject = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDirectionFacing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDirectionFacing = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDirectionMotion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDirectionMotion = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDirector = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataDisplayName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataDisplayName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataEncodedBy = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataFullFrameRatePlaybackIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataFullFrameRatePlaybackIntent = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataGenre = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataInformation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataIsMontage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataIsMontage = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataKeywords"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataKeywords = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLivePhotoVitalityScore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLivePhotoVitalityScore = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLivePhotoVitalityScoringVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLivePhotoVitalityScoringVersion = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationBody = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationHorizontalAccuracyInMeters"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationHorizontalAccuracyInMeters = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationISO6709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationISO6709 = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationNote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationNote = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataLocationRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataLocationRole = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataMake = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataModel = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataOriginalArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataPerformer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataPhonogramRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataPhonogramRights = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataPreferredAffineTransform"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataPreferredAffineTransform = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataPresentationImmersiveMedia"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataPresentationImmersiveMedia = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataProducer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataPublisher = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataRatingUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataRatingUser = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataSoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataSoftware = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataSpatialOverCaptureQualityScore"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataSpatialOverCaptureQualityScore = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataSpatialOverCaptureQualityScoringVersion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataSpatialOverCaptureQualityScoringVersion = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataVideoOrientation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataVideoOrientation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataWhiteBalanceByCCTColorMatrices"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataWhiteBalanceByCCTColorMatrices = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataWhiteBalanceByCCTWhiteBalanceFactors"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataWhiteBalanceByCCTWhiteBalanceFactors = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataYear = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeMetadataiXML"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeMetadataiXML = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataAccessibilityDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataAlbum = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataArranger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataArranger = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataAuthor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataChapter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataChapter = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataComment = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataComposer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataCopyright = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataCreationDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataCredits = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataDirector = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataDisclaimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataDisclaimer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataEncodedBy = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataFullName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataFullName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataGenre = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataHostComputer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataHostComputer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataInformation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataKeywords"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataKeywords = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataLocationISO6709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataLocationISO6709 = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataMake = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataModel = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataOriginalArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataOriginalFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataOriginalFormat = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataOriginalSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataOriginalSource = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataPerformers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataPerformers = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataPhonogramRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataPhonogramRights = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataProducer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataProduct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataProduct = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataPublisher = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataSoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataSoftware = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataSpecialPlaybackRequirements"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataSpecialPlaybackRequirements = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataTaggedCharacteristic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataTaggedCharacteristic = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataTrack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataTrack = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataTrackName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataTrackName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataURLLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataURLLink = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataWarning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataWarning = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifierQuickTimeUserDataWriter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifiers.QuickTimeUserDataWriter = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataAccountKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataAccountKind = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataAcknowledgement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataAcknowledgement = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataAlbum = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataAlbumArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataAlbumArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataAppleID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataAppleID = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataArranger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataArranger = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataArtDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataArtDirector = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataArtistID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataArtistID = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataAuthor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataBeatsPerMin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataBeatsPerMin = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataComposer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataConductor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataConductor = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataContentRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataContentRating = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataCopyright = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataCoverArt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataCoverArt = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataCredits = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataDescription = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataDirector = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataDiscCompilation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataDiscCompilation = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataDiscNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataDiscNumber = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataEQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataEQ = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataEncodedBy = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataEncodingTool"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataEncodingTool = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataExecProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataExecProducer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataGenreID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataGenreID = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataGrouping"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataGrouping = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataLinerNotes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataLinerNotes = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataLyrics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataLyrics = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataOnlineExtras"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataOnlineExtras = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataOriginalArtist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataPerformer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataPhonogramRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataPhonogramRights = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataPlaylistID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataPlaylistID = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataPredefinedGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataPredefinedGenre = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataProducer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataPublisher = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataRecordCompany"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataRecordCompany = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataReleaseDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataReleaseDate = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataSoloist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataSoloist = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataSongID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataSongID = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataSongName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataSongName = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataSoundEngineer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataSoundEngineer = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataThanks"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataThanks = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataTrackNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataTrackNumber = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataTrackSubTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataTrackSubTitle = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataUserComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataUserComment = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataIdentifieriTunesMetadataUserGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataIdentifieriTunesMetadataUserGenre = AVMetadataIdentifier(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceAudioFile"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.AudioFile = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceCommon"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.Common = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceHLSDateRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.HLSDateRange = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceID3"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.ID3 = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceISOUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.ISOUserData = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceIcy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.Icy = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceQuickTimeMetadata"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.QuickTimeMetadata = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceQuickTimeUserData"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaces.QuickTimeUserData = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataKeySpaceiTunes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataKeySpaceiTunes = AVMetadataKeySpace(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeAztecCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.AztecCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCatBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.CatBody = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCatHead"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.CatHead = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCodabarCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.CodabarCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCode128Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.Code128Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCode39Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.Code39Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCode39Mod43Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.Code39Mod43Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeCode93Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.Code93Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeDataMatrixCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.DataMatrixCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeDogBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.DogBody = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeDogHead"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.DogHead = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeEAN13Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.EAN13Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeEAN8Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.EAN8Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeFace"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.Face = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeGS1DataBarCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.GS1DataBarCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeGS1DataBarExpandedCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.GS1DataBarExpandedCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeGS1DataBarLimitedCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.GS1DataBarLimitedCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeHumanBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.HumanBody = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeHumanFullBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.HumanFullBody = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeITF14Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.ITF14Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeInterleaved2of5Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.Interleaved2of5Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeMicroPDF417Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.MicroPDF417Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeMicroQRCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.MicroQRCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypePDF417Code"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.PDF417Code = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeQRCode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.QRCode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeSalientObject"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.SalientObject = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataObjectTypeUPCECode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataObjectTypes.UPCECode = AVMetadataObjectType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyAccessibilityDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyAlbum = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyArranger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyArranger = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyArtwork"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyArtwork = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyAuthor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraFocalLength35mmEquivalent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraFocalLength35mmEquivalent = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraFrameReadoutTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraFrameReadoutTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraISOSensitivity"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraISOSensitivity = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraIdentifier = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraLensIrisFNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraLensIrisFNumber = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraLensModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraLensModel = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraShutterSpeedAngle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraShutterSpeedAngle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraShutterSpeedTime"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraShutterSpeedTime = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCameraWhiteBalance"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCameraWhiteBalance = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCinematicVideoIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCinematicVideoIntent = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCollectionUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCollectionUser = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyComment = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyComposer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyContentIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyContentIdentifier = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCopyright = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCreationDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyCredits = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyDirectionFacing"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyDirectionFacing = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyDirectionMotion"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyDirectionMotion = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyDirector = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyDisplayName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyDisplayName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyEncodedBy = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyFullFrameRatePlaybackIntent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyFullFrameRatePlaybackIntent = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyGenre = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyInformation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyIsMontage"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyIsMontage = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyKeywords"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyKeywords = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyLocationBody"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyLocationBody = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyLocationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyLocationDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyLocationISO6709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyLocationISO6709 = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyLocationName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyLocationName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyLocationNote"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyLocationNote = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyLocationRole"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyLocationRole = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyMake = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyModel = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyOriginalArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyPerformer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyPhonogramRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyPhonogramRights = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyProducer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyPublisher = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyRatingUser"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyRatingUser = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeySoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeySoftware = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTColorMatrices"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTColorMatrices = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTWhiteBalanceFactors"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyWhiteBalanceByCCTWhiteBalanceFactors = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyYear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyYear = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeMetadataKeyiXML"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeMetadataKeyiXML = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyAccessibilityDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyAccessibilityDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyAlbum = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyArranger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyArranger = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyAuthor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyChapter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyChapter = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyComment = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyComposer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyCopyright = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyCreationDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyCreationDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyCredits = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyDirector = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyDisclaimer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyDisclaimer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyEncodedBy = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyFullName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyFullName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyGenre = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyHostComputer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyHostComputer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyInformation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyInformation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyKeywords"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyKeywords = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyLocationISO6709"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyLocationISO6709 = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyMake"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyMake = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyModel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyModel = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyOriginalArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyOriginalFormat"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyOriginalFormat = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyOriginalSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyOriginalSource = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyPerformers"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyPerformers = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyPhonogramRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyPhonogramRights = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyProducer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyProduct"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyProduct = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyPublisher = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeySoftware"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeySoftware = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeySpecialPlaybackRequirements"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeySpecialPlaybackRequirements = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyTaggedCharacteristic"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyTaggedCharacteristic = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyTrack"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyTrack = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyTrackName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyTrackName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyURLLink"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyURLLink = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyWarning"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyWarning = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataQuickTimeUserDataKeyWriter"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataQuickTimeUserDataKeyWriter = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyAccountKind"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyAccountKind = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyAcknowledgement"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyAcknowledgement = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyAlbum"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyAlbum = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyAlbumArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyAlbumArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyAppleID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyAppleID = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyArranger"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyArranger = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyArtDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyArtDirector = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyArtistID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyArtistID = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyAuthor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyAuthor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyBeatsPerMin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyBeatsPerMin = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyComposer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyComposer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyConductor"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyConductor = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyContentRating"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyContentRating = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyCopyright"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyCopyright = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyCoverArt"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyCoverArt = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyCredits"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyCredits = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyDescription"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyDescription = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyDirector"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyDirector = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyDiscCompilation"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyDiscCompilation = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyDiscNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyDiscNumber = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyEQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyEQ = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyEncodedBy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyEncodedBy = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyEncodingTool"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyEncodingTool = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyExecProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyExecProducer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyGenreID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyGenreID = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyGrouping"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyGrouping = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyLinerNotes"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyLinerNotes = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyLyrics"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyLyrics = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyOnlineExtras"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyOnlineExtras = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyOriginalArtist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyOriginalArtist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyPerformer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyPerformer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyPhonogramRights"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyPhonogramRights = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyPlaylistID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyPlaylistID = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyPredefinedGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyPredefinedGenre = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyProducer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyProducer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyPublisher"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyPublisher = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyRecordCompany"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyRecordCompany = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyReleaseDate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyReleaseDate = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeySoloist"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeySoloist = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeySongID"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeySongID = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeySongName"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeySongName = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeySoundEngineer"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeySoundEngineer = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyThanks"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyThanks = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyTrackNumber"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyTrackNumber = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyTrackSubTitle"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyTrackSubTitle = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyUserComment"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyUserComment = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMetadataiTunesMetadataKeyUserGenre"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMetadataiTunesMetadataKeyUserGenre = AVMetadataKey(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMovieReferenceRestrictionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMovieReferenceRestrictionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVMovieShouldSupportAliasDataReferencesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVMovieShouldSupportAliasDataReferencesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPreset1280x720"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPreset1280x720 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPreset1920x1080"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPreset1920x1080 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPreset3840x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPreset3840x2160 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPreset640x480"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPreset640x480 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPreset960x540"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPreset960x540 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetHEVC1920x1080"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.HEVC1920x1080 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetHEVC1920x1080WithAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.HEVC1920x1080WithAlpha = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetHEVC3840x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.HEVC3840x2160 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetHEVC3840x2160WithAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.HEVC3840x2160WithAlpha = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetHEVC4320x2160"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.HEVC4320x2160 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetHEVC7680x4320"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.HEVC7680x4320 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetMVHEVC1440x1440"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.MVHEVC1440x1440 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetMVHEVC4320x4320"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.MVHEVC4320x4320 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetMVHEVC7680x7680"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.MVHEVC7680x7680 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVOutputSettingsPresetMVHEVC960x960"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVOutputSettingsPresets.MVHEVC960x960 = AVOutputSettingsPreset(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlaybackCoordinatorOtherParticipantsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlaybackCoordinatorOtherParticipantsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlaybackCoordinatorSuspensionReasonsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlaybackCoordinatorSuspensionReasonsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerEligibleForHDRPlaybackDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerEligibleForHDRPlaybackDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerIntegratedTimelineSnapshotsOutOfSyncNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerIntegratedTimelineSnapshotsOutOfSyncNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonCurrentSegmentChanged"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasons.CurrentSegmentChanged = AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonLoadedTimeRangesChanged"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasons.LoadedTimeRangesChanged = AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasonSegmentsChanged"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasons.SegmentsChanged = AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventJoinCue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventJoinCue = AVPlayerInterstitialEventCue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventLeaveCue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventLeaveCue = AVPlayerInterstitialEventCue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeEventKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeEventKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeStatusKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorAssetListResponseStatusDidChangeStatusKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeEventKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeEventKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeSkipControlLabelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeSkipControlLabelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeStateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventSkippableStateDidChangeStateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventSkippedEventKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventSkippedEventKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorCurrentEventSkippedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorCurrentEventSkippedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorEventsDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorEventsDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventDidFinishDidPlayEntireEventKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventDidFinishDidPlayEntireEventKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventDidFinishEventKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventDidFinishEventKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventDidFinishNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventDidFinishPlayoutTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventDidFinishPlayoutTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledEventKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledEventKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorInterstitialEventWasUnscheduledNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorScheduleRequestCompletedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorScheduleRequestErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorScheduleRequestErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorScheduleRequestIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorScheduleRequestIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventMonitorScheduleRequestResponseKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventMonitorScheduleRequestResponseKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerInterstitialEventNoCue"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerInterstitialEventNoCue = AVPlayerInterstitialEventCue(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemDidPlayToEndTimeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemDidPlayToEndTimeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemFailedToPlayToEndTimeErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemFailedToPlayToEndTimeErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemFailedToPlayToEndTimeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemFailedToPlayToEndTimeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemLegibleOutputTextStylingResolutionDefault"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemLegibleOutputTextStylingResolutions.Default = AVPlayerItemLegibleOutputTextStylingResolution(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemLegibleOutputTextStylingResolutionSourceAndRulesOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemLegibleOutputTextStylingResolutions.SourceAndRulesOnly = AVPlayerItemLegibleOutputTextStylingResolution(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemMediaSelectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemMediaSelectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemNewAccessLogEntryNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemNewAccessLogEntryNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemNewErrorLogEntryNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemNewErrorLogEntryNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemPlaybackStalledNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemPlaybackStalledNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemRecommendedTimeOffsetFromLiveDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemRecommendedTimeOffsetFromLiveDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemTimeJumpedNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemTimeJumpedNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemTimeJumpedOriginatingParticipantKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemTimeJumpedOriginatingParticipantKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerItemTrackVideoFieldModeDeinterlaceFields"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerItemTrackVideoFieldModeDeinterlaceFields = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeOriginatingParticipantKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeOriginatingParticipantKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonAppBackgrounded"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasons.AppBackgrounded = AVPlayerRateDidChangeReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonAudioSessionInterrupted"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasons.AudioSessionInterrupted = AVPlayerRateDidChangeReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasonKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonPlayheadReachedLiveEdge"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasons.PlayheadReachedLiveEdge = AVPlayerRateDidChangeReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonReversePlaybackReachedStartOfSeekableRange"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasons.ReversePlaybackReachedStartOfSeekableRange = AVPlayerRateDidChangeReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonSetRateCalled"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasons.SetRateCalled = AVPlayerRateDidChangeReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerRateDidChangeReasonSetRateFailed"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerRateDidChangeReasons.SetRateFailed = AVPlayerRateDidChangeReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerWaitingDuringInterstitialEventReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerWaitingDuringInterstitialEventReason = AVPlayerWaitingReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerWaitingForCoordinatedPlaybackReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerWaitingForCoordinatedPlaybackReason = AVPlayerWaitingReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerWaitingToMinimizeStallsReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerWaitingToMinimizeStallsReason = AVPlayerWaitingReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerWaitingWhileEvaluatingBufferingRateReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerWaitingWhileEvaluatingBufferingRateReason = AVPlayerWaitingReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVPlayerWaitingWithNoItemToPlayReason"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVPlayerWaitingWithNoItemToPlayReason = AVPlayerWaitingReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVRouteDetectorMultipleRoutesDetectedDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVRouteDetectorMultipleRoutesDetectedDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferAudioRendererFlushTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferAudioRendererFlushTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferAudioRendererOutputConfigurationDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferAudioRendererOutputConfigurationDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferAudioRendererWasFlushedAutomaticallyNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferAudioRendererWasFlushedAutomaticallyNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferDisplayLayerFailedToDecodeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferDisplayLayerFailedToDecodeNotification = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferDisplayLayerOutputObscuredDueToInsufficientExternalProtectionDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferDisplayLayerOutputObscuredDueToInsufficientExternalProtectionDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferDisplayLayerReadyForDisplayDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferDisplayLayerReadyForDisplayDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferRenderSynchronizerRateDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferRenderSynchronizerRateDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferVideoRendererDidFailToDecodeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferVideoRendererDidFailToDecodeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferVideoRendererDidFailToDecodeNotificationErrorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferVideoRendererDidFailToDecodeNotificationErrorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSampleBufferVideoRendererRequiresFlushToResumeDecodingDidChangeNotification"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSampleBufferVideoRendererRequiresFlushToResumeDecodingDidChangeNotification = foundation.NSNotificationName(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSemanticSegmentationMatteTypeGlasses"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSemanticSegmentationMatteTypes.Glasses = AVSemanticSegmentationMatteType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSemanticSegmentationMatteTypeHair"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSemanticSegmentationMatteTypes.Hair = AVSemanticSegmentationMatteType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSemanticSegmentationMatteTypeSkin"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSemanticSegmentationMatteTypes.Skin = AVSemanticSegmentationMatteType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSemanticSegmentationMatteTypeTeeth"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSemanticSegmentationMatteTypes.Teeth = AVSemanticSegmentationMatteType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpatialCaptureDiscomfortReasonNotEnoughLight"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSpatialCaptureDiscomfortReasons.NotEnoughLight = AVSpatialCaptureDiscomfortReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVSpatialCaptureDiscomfortReasonSubjectTooClose"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVSpatialCaptureDiscomfortReasons.SubjectTooClose = AVSpatialCaptureDiscomfortReason(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVStreamingKeyDeliveryContentKeyType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVStreamingKeyDeliveryContentKeyType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVStreamingKeyDeliveryPersistentContentKeyType"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVStreamingKeyDeliveryPersistentContentKeyType = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeAudioFallback"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.AudioFallback = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeChapterList"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.ChapterList = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeForcedSubtitlesOnly"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.ForcedSubtitlesOnly = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeMetadataReferent"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.MetadataReferent = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeRenderMetadataSource"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.RenderMetadataSource = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeSelectionFollower"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.SelectionFollower = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVTrackAssociationTypeTimecode"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVTrackAssociationTypes.Timecode = AVTrackAssociationType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetAllowsCellularAccessKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetAllowsCellularAccessKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetAllowsConstrainedNetworkAccessKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetAllowsConstrainedNetworkAccessKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetAllowsExpensiveNetworkAccessKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetAllowsExpensiveNetworkAccessKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetHTTPCookiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetHTTPCookiesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetHTTPUserAgentKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetHTTPUserAgentKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetOverrideMIMETypeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetOverrideMIMETypeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetPreferPreciseDurationAndTimingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetPreferPreciseDurationAndTimingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetPrimarySessionIdentifierKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetPrimarySessionIdentifierKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetReferenceRestrictionsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetReferenceRestrictionsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetShouldParseExternalSphericalTagsKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetShouldParseExternalSphericalTagsKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetShouldSupportAliasDataReferencesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetShouldSupportAliasDataReferencesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVURLAssetURLRequestAttributionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVURLAssetURLRequestAttributionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoAllowFrameReorderingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoAllowFrameReorderingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoAllowWideColorKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoAllowWideColorKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoApertureModeCleanAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoApertureModes.CleanAperture = AVVideoApertureMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoApertureModeEncodedPixels"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoApertureModes.EncodedPixels = AVVideoApertureMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoApertureModeProductionAperture"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoApertureModes.ProductionAperture = AVVideoApertureMode(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoAppleProRAWBitDepthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoAppleProRAWBitDepthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoAverageBitRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoAverageBitRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoAverageNonDroppableFrameRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoAverageNonDroppableFrameRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCleanApertureHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCleanApertureHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCleanApertureHorizontalOffsetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCleanApertureHorizontalOffsetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCleanApertureKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCleanApertureKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCleanApertureVerticalOffsetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCleanApertureVerticalOffsetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCleanApertureWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCleanApertureWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProRes422"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProRes422 = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProRes422HQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProRes422HQ = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProRes422LT"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProRes422LT = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProRes422Proxy"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProRes422Proxy = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProRes4444"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProRes4444 = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProRes4444XQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProRes4444XQ = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProResRAW"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProResRAW = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeAppleProResRAWHQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.AppleProResRAWHQ = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeH264"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.H264 = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeHEVC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.HEVC = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeHEVCWithAlpha"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.HEVCWithAlpha = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeJPEG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.JPEG = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCodecTypeJPEGXL"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCodecTypes.JPEGXL = AVVideoCodecType(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPrimariesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPrimariesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPrimaries_EBU_3213"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPrimaries_EBU_3213 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPrimaries_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPrimaries_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPrimaries_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPrimaries_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPrimaries_P3_D65"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPrimaries_P3_D65 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPrimaries_SMPTE_C"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPrimaries_SMPTE_C = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoColorPropertiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoColorPropertiesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCompositionPerFrameHDRDisplayMetadataPolicyGenerate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCompositionPerFrameHDRDisplayMetadataPolicys.Generate = AVVideoCompositionPerFrameHDRDisplayMetadataPolicy(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCompositionPerFrameHDRDisplayMetadataPolicyPropagate"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCompositionPerFrameHDRDisplayMetadataPolicys.Propagate = AVVideoCompositionPerFrameHDRDisplayMetadataPolicy(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoCompressionPropertiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoCompressionPropertiesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoDecompressionPropertiesKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoDecompressionPropertiesKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoEncoderSpecificationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoEncoderSpecificationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoExpectedSourceFrameRateKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoExpectedSourceFrameRateKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoH264EntropyModeCABAC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoH264EntropyModeCABAC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoH264EntropyModeCAVLC"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoH264EntropyModeCAVLC = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoH264EntropyModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoH264EntropyModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoHeightKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoHeightKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoMaxKeyFrameIntervalDurationKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoMaxKeyFrameIntervalDurationKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoMaxKeyFrameIntervalKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoMaxKeyFrameIntervalKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoPixelAspectRatioHorizontalSpacingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoPixelAspectRatioHorizontalSpacingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoPixelAspectRatioKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoPixelAspectRatioKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoPixelAspectRatioVerticalSpacingKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoPixelAspectRatioVerticalSpacingKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Baseline30"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Baseline30 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Baseline31"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Baseline31 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Baseline41"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Baseline41 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264BaselineAutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264BaselineAutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264High40"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264High40 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264High41"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264High41 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264HighAutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264HighAutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Main30"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Main30 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Main31"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Main31 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Main32"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Main32 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264Main41"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264Main41 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelH264MainAutoLevel"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelH264MainAutoLevel = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoProfileLevelKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoProfileLevelKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoQualityKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoQualityKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoRangeHLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoRanges.HLG = AVVideoRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoRangePQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoRanges.PQ = AVVideoRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoRangeSDR"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoRanges.SDR = AVVideoRange(objc.GoString(cstr))
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoScalingModeFit"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoScalingModeFit = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoScalingModeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoScalingModeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoScalingModeResize"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoScalingModeResize = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoScalingModeResizeAspect"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoScalingModeResizeAspect = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoScalingModeResizeAspectFill"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoScalingModeResizeAspectFill = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunctionKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunctionKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunction_IEC_sRGB"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunction_IEC_sRGB = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunction_ITU_R_2100_HLG"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunction_ITU_R_2100_HLG = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunction_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunction_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunction_Linear"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunction_Linear = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunction_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunction_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoTransferFunction_SMPTE_ST_2084_PQ"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoTransferFunction_SMPTE_ST_2084_PQ = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoWidthKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoWidthKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoYCbCrMatrixKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoYCbCrMatrixKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoYCbCrMatrix_ITU_R_2020"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoYCbCrMatrix_ITU_R_2020 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoYCbCrMatrix_ITU_R_601_4"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoYCbCrMatrix_ITU_R_601_4 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoYCbCrMatrix_ITU_R_709_2"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoYCbCrMatrix_ITU_R_709_2 = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "AVVideoYCbCrMatrix_SMPTE_240M_1995"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				AVVideoYCbCrMatrix_SMPTE_240M_1995 = objc.GoString(cstr)
			}
		}
	}

}

// AVAssetDownloadedAssetEvictionPrioritys provides typed accessors for [AVAssetDownloadedAssetEvictionPriority] constants.
var AVAssetDownloadedAssetEvictionPrioritys struct {
	// Default: The default eviction priority.
	Default AVAssetDownloadedAssetEvictionPriority
	// Important: An eviction priority that indicates that this asset is important and the system should evict lower-priority assets first.
	Important AVAssetDownloadedAssetEvictionPriority
}

// AVAssetImageGeneratorApertureModes provides typed accessors for [AVAssetImageGeneratorApertureMode] constants.
var AVAssetImageGeneratorApertureModes struct {
	// CleanAperture: A mode that applies both pixel aspect ratio and clean aperture.
	CleanAperture AVAssetImageGeneratorApertureMode
	// EncodedPixels: A mode that applies neither pixel aspect ratio nor clean aperture.
	EncodedPixels AVAssetImageGeneratorApertureMode
	// ProductionAperture: A mode that applies only pixel aspect ratio.
	ProductionAperture AVAssetImageGeneratorApertureMode
}

// AVAssetImageGeneratorDynamicRangePolicys provides typed accessors for [AVAssetImageGeneratorDynamicRangePolicy] constants.
var AVAssetImageGeneratorDynamicRangePolicys struct {
	// ForceSDR: A policy that forces conversion to standard dynamic range.
	ForceSDR AVAssetImageGeneratorDynamicRangePolicy
	// MatchSource: A policy that preserves the color parameters of the source media.
	MatchSource AVAssetImageGeneratorDynamicRangePolicy
}

// AVAssetPlaybackConfigurationOptions provides typed accessors for [AVAssetPlaybackConfigurationOption] constants.
var AVAssetPlaybackConfigurationOptions struct {
	// AppleImmersiveVideo: Indicates whether the asset is Apple Immersive Video.
	AppleImmersiveVideo AVAssetPlaybackConfigurationOption
	// NonRectilinearProjection: Indicates whether the asset calls for the use of a non-rectilinear projection for rendering video.
	NonRectilinearProjection AVAssetPlaybackConfigurationOption
	// SpatialVideo: An option that indicates whether the asset can render as spatial video.
	SpatialVideo AVAssetPlaybackConfigurationOption
	// StereoMultiviewVideo: An option that indicates whether the asset is in a multiview compression format and can render as stereo video.
	StereoMultiviewVideo AVAssetPlaybackConfigurationOption
	// StereoVideo: An option that indicates whether the asset can render as stereo video.
	StereoVideo AVAssetPlaybackConfigurationOption
}

// AVAssetWriterInputMediaDataLocations provides typed accessors for [AVAssetWriterInputMediaDataLocation] constants.
var AVAssetWriterInputMediaDataLocations struct {
	// BeforeMainMediaDataNotInterleaved: A value that indicates to use noninterleaved data, and write it before interleaved data.
	BeforeMainMediaDataNotInterleaved AVAssetWriterInputMediaDataLocation
	// InterleavedWithMainMediaData: A value that indicates to interleave the input’s media data with other media data.
	InterleavedWithMainMediaData AVAssetWriterInputMediaDataLocation
	// SparselyInterleavedWithMainMediaData: Indicates that there may be large segments of time without any media data from this track. When mediaDataLocation is set to this value, AVAssetWriter will interleave the media data, but will not wait for media data from this track to achieve tight interleaving with other tracks.
	SparselyInterleavedWithMainMediaData AVAssetWriterInputMediaDataLocation
}

// AVAudioTimePitchAlgorithms provides typed accessors for [AVAudioTimePitchAlgorithm] constants.
var AVAudioTimePitchAlgorithms struct {
	// LowQualityZeroLatency: A low-quality and very low computationally intensive pitch algorithm.
	LowQualityZeroLatency AVAudioTimePitchAlgorithm
	// Spectral: A highest-quality time pitch algorithm that’s suitable for music.
	Spectral AVAudioTimePitchAlgorithm
	// TimeDomain: A modest quality time pitch algorithm that’s suitable for voice.
	TimeDomain AVAudioTimePitchAlgorithm
	// Varispeed: A high-quality time pitch algorithm that doesn’t perform pitch correction.
	Varispeed AVAudioTimePitchAlgorithm
}

// AVCaptureDeviceTypes provides typed accessors for [AVCaptureDeviceType] constants.
var AVCaptureDeviceTypes struct {
	// BuiltInDualCamera: A built-in camera device type that consists of a wide-angle and telephoto camera.
	BuiltInDualCamera AVCaptureDeviceType
	// BuiltInDualWideCamera: A built-in camera device type that consists of two cameras of fixed focal length, one ultrawide angle and one wide angle.
	BuiltInDualWideCamera AVCaptureDeviceType
	// BuiltInDuoCamera: A built-in dual camera device type.
	BuiltInDuoCamera AVCaptureDeviceType
	// BuiltInLiDARDepthCamera: A device that consists of two cameras, one LiDAR and one YUV.
	BuiltInLiDARDepthCamera AVCaptureDeviceType
	// BuiltInMicrophone: A built-in microphone.
	BuiltInMicrophone AVCaptureDeviceType
	// BuiltInTelephotoCamera: A built-in camera device type with a longer focal length than a wide-angle camera.
	BuiltInTelephotoCamera AVCaptureDeviceType
	// BuiltInTripleCamera: A built-in camera device type that consists of three cameras of fixed focal length, one ultrawide angle, one wide angle, and one telephoto.
	BuiltInTripleCamera AVCaptureDeviceType
	// BuiltInTrueDepthCamera: A device that consists of two cameras, one Infrared and one YUV.
	BuiltInTrueDepthCamera AVCaptureDeviceType
	// BuiltInUltraWideCamera: A built-in camera device type with a shorter focal length than a wide-angle camera.
	BuiltInUltraWideCamera AVCaptureDeviceType
	// BuiltInWideAngleCamera: A built-in wide-angle camera device type.
	BuiltInWideAngleCamera AVCaptureDeviceType
	// ContinuityCamera: A Continuity Camera device type.
	ContinuityCamera AVCaptureDeviceType
	// DeskViewCamera: A virtual overhead camera that captures a user’s desk.
	DeskViewCamera AVCaptureDeviceType
	// External: An external device type.
	External AVCaptureDeviceType
	// ExternalUnknown: An unknown external device type.
	ExternalUnknown AVCaptureDeviceType
	// Microphone: A microphone device type.
	Microphone AVCaptureDeviceType
}

// AVCaptureReactionTypes provides typed accessors for [AVCaptureReactionType] constants.
var AVCaptureReactionTypes struct {
	// Balloons: A reaction that displays balloons rising through the scene.
	Balloons AVCaptureReactionType
	// Confetti: A reaction that displays festive spots of color falling through the scene.
	Confetti AVCaptureReactionType
	// Fireworks: A reaction that displays fireworks bursting in the background.
	Fireworks AVCaptureReactionType
	// Heart: A reaction that displays one or more heart symbols.
	Heart AVCaptureReactionType
	// Lasers: A reaction that displays a bright laser show projecting into the scene.
	Lasers AVCaptureReactionType
	// Rain: A reaction that displays a dark and stormy night.
	Rain AVCaptureReactionType
	// ThumbsDown: A reaction that displays a thumbs-down symbol.
	ThumbsDown AVCaptureReactionType
	// ThumbsUp: A reaction that displays a thumbs-up symbol.
	ThumbsUp AVCaptureReactionType
}

// AVCaptureSessionPresets provides typed accessors for [AVCaptureSessionPreset] constants.
var AVCaptureSessionPresets struct {
	// High: A preset suitable for capturing high-quality output.
	High AVCaptureSessionPreset
	// InputPriority: A preset that doesn’t specify audio and video output settings for a capture session.
	InputPriority AVCaptureSessionPreset
	// Low: A preset suitable for capturing low-quality output.
	Low AVCaptureSessionPreset
	// Medium: A preset suitable for capturing medium-quality output.
	Medium AVCaptureSessionPreset
	// Photo: A preset suitable for capturing high-resolution photo quality output.
	Photo AVCaptureSessionPreset
}

// AVContentKeyRequestRetryReasons provides typed accessors for [AVContentKeyRequestRetryReason] constants.
var AVContentKeyRequestRetryReasons struct {
	// ReceivedObsoleteContentKey: An obsolete key response that was set on the previous content key request.
	ReceivedObsoleteContentKey AVContentKeyRequestRetryReason
	// ReceivedResponseWithExpiredLease: A key response with an expired lease that was set on the previous content key request.
	ReceivedResponseWithExpiredLease AVContentKeyRequestRetryReason
	// TimedOut: A key response that wasn’t set soon enough.
	TimedOut AVContentKeyRequestRetryReason
}

// AVContentKeySessionServerPlaybackContextOptions provides typed accessors for [AVContentKeySessionServerPlaybackContextOption] constants.
var AVContentKeySessionServerPlaybackContextOptions struct {
	// ProtocolVersions: Specifies the versions of the content protection protocols supported by the application.
	ProtocolVersions AVContentKeySessionServerPlaybackContextOption
	// ServerChallenge: Specifies a nonce to include in the secure server playback context (SPC) to prevent replay attacks.
	ServerChallenge AVContentKeySessionServerPlaybackContextOption
}

// AVContentKeySystems provides typed accessors for [AVContentKeySystem] constants.
var AVContentKeySystems struct {
	// AuthorizationToken: A method of key delivery that uses a token to authorize playback.
	AuthorizationToken AVContentKeySystem
	// ClearKey: A method of key delivery that uses a clear key system.
	ClearKey AVContentKeySystem
	// FairPlayStreaming: A method of key delivery that uses FairPlay Streaming.
	FairPlayStreaming AVContentKeySystem
}

// AVCoordinatedPlaybackSuspensionReasons provides typed accessors for [AVCoordinatedPlaybackSuspensionReason] constants.
var AVCoordinatedPlaybackSuspensionReasons struct {
	// AudioSessionInterrupted: The system interrupts a participant’s audio session.
	AudioSessionInterrupted AVCoordinatedPlaybackSuspensionReason
	// CoordinatedPlaybackNotPossible: It’s not possible for a participant to start or resume coordinated playback.
	CoordinatedPlaybackNotPossible AVCoordinatedPlaybackSuspensionReason
	// PlayingInterstitial: A participant is playing content other than the primary content.
	PlayingInterstitial AVCoordinatedPlaybackSuspensionReason
	// StallRecovery: The player object is buffering media data after a stall.
	StallRecovery AVCoordinatedPlaybackSuspensionReason
	// UserActionRequired: A playback object requires user intervention to resume playback.
	UserActionRequired AVCoordinatedPlaybackSuspensionReason
	// UserIsChangingCurrentTime: A participant is actively changing the current time.
	UserIsChangingCurrentTime AVCoordinatedPlaybackSuspensionReason
}

// AVFileTypeProfiles provides typed accessors for [AVFileTypeProfile] constants.
var AVFileTypeProfiles struct {
	// MPEG4AppleHLS: A file type profile for Apple HTTP Live Streaming.
	MPEG4AppleHLS AVFileTypeProfile
	// MPEG4CMAFCompliant: A file type profile that complies with the Common Media Application Format (CMAF).
	MPEG4CMAFCompliant AVFileTypeProfile
}

// AVFileTypes provides typed accessors for [AVFileType] constants.
var AVFileTypes struct {
	// AC3: The UTI for the AC3 audio file format.
	AC3 AVFileType
	// AHAP: The UTI for the Apple Haptics Audio Pattern file format.
	AHAP AVFileType
	// AIFC: The UTI for the AIFC audio file format.
	AIFC AVFileType
	// AIFF: The UTI for the AIFF audio file format.
	AIFF AVFileType
	// AMR: The UTI for the adaptive multirate audio file format.
	AMR AVFileType
	// AVCI: The UTI for the high-efficiency image file format that contains H.264 compressed images.
	AVCI AVFileType
	// AppleM4A: The UTI for the Apple m4a audio file format.
	AppleM4A AVFileType
	// AppleM4V: The UTI for the iTunes video file format.
	AppleM4V AVFileType
	// AppleiTT: The UTI for the Apple iTT caption file format.
	AppleiTT AVFileType
	// CoreAudioFormat: The UTI for the Core Audio Format.
	CoreAudioFormat AVFileType
	// DICOM: A UTI for the Digital Imaging and Communications in Medicine (DICOM) file format.
	DICOM AVFileType
	// DNG: The UTI for the Adobe Digital Negative file format.
	DNG AVFileType
	// EnhancedAC3: The UTI for the enhanced AC3 audio file format.
	EnhancedAC3 AVFileType
	// HEIC: The UTI for the high-efficiency image file format that contains HEVC compressed images.
	HEIC AVFileType
	// HEIF: The UTI for the high-efficiency image file format that contains compressed images from any codec.
	HEIF AVFileType
	// JPEG: The UTI for the JPEG (JFIF) format.
	JPEG AVFileType
	// MPEG4: The UTI for the MPEG-4 file format.
	MPEG4 AVFileType
	// MPEGLayer3: The UTI for the MPEG Audio Layer III file format.
	MPEGLayer3 AVFileType
	// QuickTimeAudio: A UTI for the QuickTime audio file format
	QuickTimeAudio AVFileType
	// QuickTimeMovie: The UTI for the QuickTime movie file format.
	QuickTimeMovie AVFileType
	// SCC: The UTI for the Scenarist closed-caption file format.
	SCC AVFileType
	// SunAU: The UTI for the Sun/NeXT audio file format.
	SunAU AVFileType
	// TIFF: The UTI for the tagged image file format.
	TIFF AVFileType
	// WAVE: The UTI for the WAVE audio file format.
	WAVE AVFileType
}

// AVLayerVideoGravitys provides typed accessors for [AVLayerVideoGravity] constants.
var AVLayerVideoGravitys struct {
	// Resize: The video stretches to fill the layer’s bounds.
	Resize AVLayerVideoGravity
	// ResizeAspect: The video preserves its aspect ratio and fits it within the layer’s bounds.
	ResizeAspect AVLayerVideoGravity
	// ResizeAspectFill: The video preserves its aspect ratio and fills the layer’s bounds.
	ResizeAspectFill AVLayerVideoGravity
}

// AVMediaCharacteristics provides typed accessors for [AVMediaCharacteristic] constants.
var AVMediaCharacteristics struct {
	// Audible: A media characteristic that indicates that a track or media selection option includes audible content.
	Audible AVMediaCharacteristic
	// CarriesVideoStereoMetadata: A media characteristic that indicates that the stereoscopic video track carries additional information related to the stereoscopic video.
	CarriesVideoStereoMetadata AVMediaCharacteristic
	// ContainsAlphaChannel: A media characteristic that indicates that a track contains an alpha channel.
	ContainsAlphaChannel AVMediaCharacteristic
	// ContainsHDRVideo: A media characteristic that indicates that a track contains HDR video.
	ContainsHDRVideo AVMediaCharacteristic
	// ContainsOnlyForcedSubtitles: A media characteristic that indicates that a track or media selection option presents only forced subtitles.
	ContainsOnlyForcedSubtitles AVMediaCharacteristic
	// ContainsStereoMultiviewVideo: A media characteristic that indicates that a track contains stereoscopic video captured in a multiview compression format.
	ContainsStereoMultiviewVideo AVMediaCharacteristic
	// DescribesMusicAndSoundForAccessibility: A media characteristic that indicates that a track or media selection option includes legible content in the language of its specified locale.
	DescribesMusicAndSoundForAccessibility AVMediaCharacteristic
	// DescribesVideoForAccessibility: A media characteristic that indicates the media includes audible content that describes the visual portion of the presentation.
	DescribesVideoForAccessibility AVMediaCharacteristic
	// DubbedTranslation: A media characteristic that indicates that a track or media selection option contains audio language or dialect translation of the original content.
	DubbedTranslation AVMediaCharacteristic
	// EasyToRead: A media characteristic that indicates a track or media selection option provides legible content that’s edited for easy reading.
	EasyToRead AVMediaCharacteristic
	// EnhancesSpeechIntelligibility: A media characteristic that indicates a track or media selection option includes audio processed to enhance the intelligibility of speech.
	EnhancesSpeechIntelligibility AVMediaCharacteristic
	// FrameBased: A media characteristic that indicates that a track or media selection option includes frame-based content.
	FrameBased AVMediaCharacteristic
	// IndicatesHorizontalFieldOfView: A media characteristic that indicates the video track carries information related to the horizontal field of view.
	IndicatesHorizontalFieldOfView AVMediaCharacteristic
	// IndicatesNonRectilinearProjection: A media characteristic that indicates the video track carries information related to how it should be projected for display.
	IndicatesNonRectilinearProjection AVMediaCharacteristic
	// IsAuxiliaryContent: A media characteristic that indicates a track or media selection option includes content its author indicates is auxiliary to the asset’s presentation.
	IsAuxiliaryContent AVMediaCharacteristic
	// IsMainProgramContent: A media characteristic that indicates a track or media selection option includes content its author indicates is essential to the asset’s presentation.
	IsMainProgramContent AVMediaCharacteristic
	// IsOriginalContent: A media characteristic that indicates that a track or media selection option contains original content.
	IsOriginalContent AVMediaCharacteristic
	// LanguageTranslation: A media characteristic that indicates that a track or media selection option contains a language or dialect translation of the original content.
	LanguageTranslation AVMediaCharacteristic
	// Legible: A media characteristic that indicates that a track or media selection option includes legible content.
	Legible AVMediaCharacteristic
	// MachineGenerated: A media characteristic that indicates that a track was generated in an automated fashion by a machine.
	MachineGenerated AVMediaCharacteristic
	// TactileMinimal: A media characteristic that indicates that a track or media selection option includes haptic content.
	TactileMinimal AVMediaCharacteristic
	// TranscribesSpokenDialogForAccessibility: A media characteristic that indicates that a media selection option includes legible content that transcribes spoken dialog.
	TranscribesSpokenDialogForAccessibility AVMediaCharacteristic
	// UsesWideGamutColorSpace: A media characteristic that indicates that a track uses a wide-gamut color space.
	UsesWideGamutColorSpace AVMediaCharacteristic
	// Visual: A media characteristic that indicates that a track or media selection option includes visual content.
	Visual AVMediaCharacteristic
	// VoiceOverTranslation: A media characteristic that indicates that a track or media selection option contains a language translation and verbal interpretation of spoken dialog.
	VoiceOverTranslation AVMediaCharacteristic
}

// AVMediaTypes provides typed accessors for [AVMediaType] constants.
var AVMediaTypes struct {
	// Audio: The media contains audio media.
	Audio            AVMediaType
	AuxiliaryPicture AVMediaType
	// ClosedCaption: The media contains closed-caption content.
	ClosedCaption AVMediaType
	// DepthData: The media contains depth data.
	DepthData AVMediaType
	// Haptic: The media contains haptic content.
	Haptic AVMediaType
	// Metadata: The media contains metadata.
	Metadata AVMediaType
	// MetadataObject: The media contains metadata objects.
	MetadataObject AVMediaType
	// Muxed: The media contains muxed media.
	Muxed AVMediaType
	// Subtitle: The media contains subtitles.
	Subtitle AVMediaType
	// Text: The media contains text.
	Text AVMediaType
	// Timecode: The media contains a time code.
	Timecode AVMediaType
	// Video: The media contains video.
	Video AVMediaType
}

// AVMetadataFormats provides typed accessors for [AVMetadataFormat] constants.
var AVMetadataFormats struct {
	// HLSMetadata: The HLS metadata format.
	HLSMetadata AVMetadataFormat
	// ID3Metadata: The ID3 metadata format.
	ID3Metadata AVMetadataFormat
	// ISOUserData: The ISO user data metadata format.
	ISOUserData AVMetadataFormat
	// QuickTimeMetadata: The QuickTime metadata format.
	QuickTimeMetadata AVMetadataFormat
	// QuickTimeUserData: The QuickTime user data metadata format.
	QuickTimeUserData AVMetadataFormat
	// Unknown: An unknown metadata format.
	Unknown AVMetadataFormat
}

// AVMetadataIdentifiers provides typed accessors for [AVMetadataIdentifier] constants.
var AVMetadataIdentifiers struct {
	// ID3MetadataAlbumSortOrder: An identifier that represents how to sort the album.
	ID3MetadataAlbumSortOrder AVMetadataIdentifier
	// ID3MetadataAlbumTitle: An identifier that represents the title of the recording.
	ID3MetadataAlbumTitle AVMetadataIdentifier
	// ID3MetadataAttachedPicture: An identifier that represents an image relating to the audio file.
	ID3MetadataAttachedPicture AVMetadataIdentifier
	// ID3MetadataAudioEncryption: An identifier that represents the encryption details of the audio stream.
	ID3MetadataAudioEncryption AVMetadataIdentifier
	// ID3MetadataAudioSeekPointIndex: An identifier that represents the list of seek points within the audio file.
	ID3MetadataAudioSeekPointIndex AVMetadataIdentifier
	// ID3MetadataBand: An identifier that represents additional information about the performers in the recording.
	ID3MetadataBand AVMetadataIdentifier
	// ID3MetadataBeatsPerMinute: An identifier that represents the beats per minute of the audio.
	ID3MetadataBeatsPerMinute AVMetadataIdentifier
	// ID3MetadataComments: An identifier that represents additional text information for the media.
	ID3MetadataComments AVMetadataIdentifier
	// ID3MetadataCommercial: An identifier that represents the commercial details for the media.
	ID3MetadataCommercial AVMetadataIdentifier
	// ID3MetadataCommercialInformation: An identifier that represents the webpage containing purchasing information.
	ID3MetadataCommercialInformation AVMetadataIdentifier
	// ID3MetadataComposer: An identifier that represents the name of the composer.
	ID3MetadataComposer AVMetadataIdentifier
	// ID3MetadataConductor: An identifier that represents the name of the conductor.
	ID3MetadataConductor AVMetadataIdentifier
	// ID3MetadataContentGroupDescription: An identifier that indicates the sound belongs to a larger category of sounds or music.
	ID3MetadataContentGroupDescription AVMetadataIdentifier
	// ID3MetadataContentType: An identifier that represents the media content type.
	ID3MetadataContentType AVMetadataIdentifier
	// ID3MetadataCopyright: An identifier that represents the copyright statement.
	ID3MetadataCopyright AVMetadataIdentifier
	// ID3MetadataCopyrightInformation: An identifier that represents the webpage describing the terms of use and ownership.
	ID3MetadataCopyrightInformation AVMetadataIdentifier
	// ID3MetadataDate: An identifier that represents the date for the recording.
	ID3MetadataDate AVMetadataIdentifier
	// ID3MetadataEncodedBy: An identifier that represents the person or organization responsible for encoding the media.
	ID3MetadataEncodedBy AVMetadataIdentifier
	// ID3MetadataEncodedWith: An identifier that represents the software or hardware and settings used for encoding.
	ID3MetadataEncodedWith AVMetadataIdentifier
	// ID3MetadataEncodingTime: An identifier that represents the encoding time of the audio.
	ID3MetadataEncodingTime AVMetadataIdentifier
	// ID3MetadataEncryption: An identifier that represents the encryption method used.
	ID3MetadataEncryption AVMetadataIdentifier
	// ID3MetadataEqualization: An identifier that represents the equalization curve within the audio file.
	ID3MetadataEqualization AVMetadataIdentifier
	// ID3MetadataEqualization2: An identifier that represents the equalization curve within the audio file.
	ID3MetadataEqualization2 AVMetadataIdentifier
	// ID3MetadataEventTimingCodes: An identifier that represents the timing codes used for synchronization with key events in a song or sound.
	ID3MetadataEventTimingCodes AVMetadataIdentifier
	// ID3MetadataFileOwner: An identifier that represents the name of the owner or licensee of the file and it’s contents.
	ID3MetadataFileOwner AVMetadataIdentifier
	// ID3MetadataFileType: An identifier that represents the file type of the audio.
	ID3MetadataFileType AVMetadataIdentifier
	// ID3MetadataGeneralEncapsulatedObject: An identifier that represents the details of a file.
	ID3MetadataGeneralEncapsulatedObject AVMetadataIdentifier
	// ID3MetadataGroupIdentifier: An identifier that represents the grouping of distinct frames.
	ID3MetadataGroupIdentifier AVMetadataIdentifier
	// ID3MetadataInitialKey: An identifier that represents the musical key in which the sound starts.
	ID3MetadataInitialKey AVMetadataIdentifier
	// ID3MetadataInternationalStandardRecordingCode: An identifier that represents the international standard recording code.
	ID3MetadataInternationalStandardRecordingCode AVMetadataIdentifier
	// ID3MetadataInternetRadioStationName: An identifier that represents the name of the internet radio station streaming the audio.
	ID3MetadataInternetRadioStationName AVMetadataIdentifier
	// ID3MetadataInternetRadioStationOwner: An identifier that represents the name of the owner of the internet radio station streaming the audio.
	ID3MetadataInternetRadioStationOwner AVMetadataIdentifier
	// ID3MetadataInvolvedPeopleList_v23: An identifier that represents the list of names of contributors to the media.
	ID3MetadataInvolvedPeopleList_v23 AVMetadataIdentifier
	// ID3MetadataInvolvedPeopleList_v24: An identifier that represents the list of names of contributors to the media.
	ID3MetadataInvolvedPeopleList_v24 AVMetadataIdentifier
	// ID3MetadataLanguage: An identifier that represents the language of the text or lyrics spoken or sung in the audio.
	ID3MetadataLanguage AVMetadataIdentifier
	// ID3MetadataLeadPerformer: An identifier that represents the main artist of the recording.
	ID3MetadataLeadPerformer AVMetadataIdentifier
	// ID3MetadataLength: An identifier that represents the length of the audio file in milliseconds.
	ID3MetadataLength AVMetadataIdentifier
	// ID3MetadataLink: An identifier that represents the link information from an ID3 tag that might reside in another audio file or alone in a binary file.
	ID3MetadataLink AVMetadataIdentifier
	// ID3MetadataLyricist: An identifier that represents the writer(s) of the text or lyrics in the recording.
	ID3MetadataLyricist AVMetadataIdentifier
	// ID3MetadataMPEGLocationLookupTable: An identifier that represents the lookup table used to increase performance and accuracy of jumps within an MPEG audio file.
	ID3MetadataMPEGLocationLookupTable AVMetadataIdentifier
	// ID3MetadataMediaType: An identifier that represents which media the sound originated from.
	ID3MetadataMediaType AVMetadataIdentifier
	// ID3MetadataModifiedBy: An identifier that represents the people behind a remix and similar interpretations of another existing piece.
	ID3MetadataModifiedBy AVMetadataIdentifier
	// ID3MetadataMood: An identifier that represents the mood of the audio.
	ID3MetadataMood AVMetadataIdentifier
	// ID3MetadataMusicCDIdentifier: An identifier that represents the ID used to identify the CD in databases such as the Compact Disc Database.
	ID3MetadataMusicCDIdentifier AVMetadataIdentifier
	// ID3MetadataMusicianCreditsList: An identifier that represents the mapping between an instrument and the musician that played it.
	ID3MetadataMusicianCreditsList AVMetadataIdentifier
	// ID3MetadataOfficialArtistWebpage: An identifier that represents the artist’s official webpage.
	ID3MetadataOfficialArtistWebpage AVMetadataIdentifier
	// ID3MetadataOfficialAudioFileWebpage: An identifier that represents the official webpage for the audio file.
	ID3MetadataOfficialAudioFileWebpage AVMetadataIdentifier
	// ID3MetadataOfficialAudioSourceWebpage: An identifier that represents the official webpage for the source of the audio file.
	ID3MetadataOfficialAudioSourceWebpage AVMetadataIdentifier
	// ID3MetadataOfficialInternetRadioStationHomepage: An identifier that represents the official homepage of the internet radio station.
	ID3MetadataOfficialInternetRadioStationHomepage AVMetadataIdentifier
	// ID3MetadataOfficialPublisherWebpage: An identifier that represents the official webpage for the publisher.
	ID3MetadataOfficialPublisherWebpage AVMetadataIdentifier
	// ID3MetadataOriginalAlbumTitle: An identifier that represents the title of the original recording or source of sound.
	ID3MetadataOriginalAlbumTitle AVMetadataIdentifier
	// ID3MetadataOriginalArtist: An identifier that represents the performer(s) of the original recording.
	ID3MetadataOriginalArtist AVMetadataIdentifier
	// ID3MetadataOriginalFilename: An identifier that represents the original filename for the recording.
	ID3MetadataOriginalFilename AVMetadataIdentifier
	// ID3MetadataOriginalLyricist: An identifier that represents the text writer(s) of the original recording.
	ID3MetadataOriginalLyricist AVMetadataIdentifier
	// ID3MetadataOriginalReleaseTime: An identifier that represents the release time for the original recording.
	ID3MetadataOriginalReleaseTime AVMetadataIdentifier
	// ID3MetadataOriginalReleaseYear: An identifier that represents the release year for the original recording.
	ID3MetadataOriginalReleaseYear AVMetadataIdentifier
	// ID3MetadataOwnership: An identifier that represents the transaction details indicating proof of ownership if signed.
	ID3MetadataOwnership AVMetadataIdentifier
	// ID3MetadataPartOfASet: An identifier that represents the part of a set the audio came from.
	ID3MetadataPartOfASet AVMetadataIdentifier
	// ID3MetadataPayment: An identifier that represents the webpage that handles the process of paying for the audio file.
	ID3MetadataPayment AVMetadataIdentifier
	// ID3MetadataPerformerSortOrder: An identifier that represents the performer sort order.
	ID3MetadataPerformerSortOrder AVMetadataIdentifier
	// ID3MetadataPlayCounter: An identifier that represents the play count of the audio file.
	ID3MetadataPlayCounter AVMetadataIdentifier
	// ID3MetadataPlaylistDelay: An identifier that represents the number of milliseconds of silence between every song in a playlist.
	ID3MetadataPlaylistDelay AVMetadataIdentifier
	// ID3MetadataPopularimeter: An identifier that represents the rating for the audio file.
	ID3MetadataPopularimeter AVMetadataIdentifier
	// ID3MetadataPositionSynchronization: An identifier that represents the time offset of the first frame in the stream.
	ID3MetadataPositionSynchronization AVMetadataIdentifier
	// ID3MetadataPrivate: An identifier that represents the information from a software producer that its program uses.
	ID3MetadataPrivate AVMetadataIdentifier
	// ID3MetadataProducedNotice: An identifier that represents the produced notice.
	ID3MetadataProducedNotice AVMetadataIdentifier
	// ID3MetadataPublisher: An identifier that represents the name of the label or publisher.
	ID3MetadataPublisher AVMetadataIdentifier
	// ID3MetadataRecommendedBufferSize: An identifier that represents the buffer size the server recommends.
	ID3MetadataRecommendedBufferSize AVMetadataIdentifier
	// ID3MetadataRecordingDates: An identifier that represents additional recording dates that complement year, date, and time identifiers.
	ID3MetadataRecordingDates AVMetadataIdentifier
	// ID3MetadataRecordingTime: An identifier that represents the recording time.
	ID3MetadataRecordingTime AVMetadataIdentifier
	// ID3MetadataRelativeVolumeAdjustment: An identifier that represents the increase or decrease of volume on each channel while the file plays.
	ID3MetadataRelativeVolumeAdjustment AVMetadataIdentifier
	// ID3MetadataRelativeVolumeAdjustment2: An identifier that represents the increase or decrease of volume on each channel while the file plays.
	ID3MetadataRelativeVolumeAdjustment2 AVMetadataIdentifier
	// ID3MetadataReleaseTime: An identifier that represents the time of the first release.
	ID3MetadataReleaseTime AVMetadataIdentifier
	// ID3MetadataReverb: An identifier that represents the adjustments to echoes of different kinds.
	ID3MetadataReverb AVMetadataIdentifier
	// ID3MetadataSeek: An identifier that represents the location of other tags in a file or stream.
	ID3MetadataSeek AVMetadataIdentifier
	// ID3MetadataSetSubtitle: An identifier that represents the set subtitle the track belongs to.
	ID3MetadataSetSubtitle AVMetadataIdentifier
	// ID3MetadataSignature: An identifier that represents the group of frames to sign.
	ID3MetadataSignature AVMetadataIdentifier
	// ID3MetadataSize: An identifier that represents the size of the audio file in bytes.
	ID3MetadataSize AVMetadataIdentifier
	// ID3MetadataSubTitle: An identifier that represents the information relating to the contents title.
	ID3MetadataSubTitle AVMetadataIdentifier
	// ID3MetadataSynchronizedLyric: An identifier that represents the words in the audio file as text in sync with the audio.
	ID3MetadataSynchronizedLyric AVMetadataIdentifier
	// ID3MetadataSynchronizedTempoCodes: An identifier that represents the tempo codes used for a more accurate description of the tempo of a musical piece.
	ID3MetadataSynchronizedTempoCodes AVMetadataIdentifier
	// ID3MetadataTaggingTime: An identifier that represents the time of tagging.
	ID3MetadataTaggingTime AVMetadataIdentifier
	// ID3MetadataTermsOfUse: An identifier that represents the brief description of the terms of use and ownership of the file.
	ID3MetadataTermsOfUse AVMetadataIdentifier
	// ID3MetadataTime: An identifier that represents the time for the recording.
	ID3MetadataTime AVMetadataIdentifier
	// ID3MetadataTitleDescription: An identifier that represents the name of the piece.
	ID3MetadataTitleDescription AVMetadataIdentifier
	// ID3MetadataTitleSortOrder: An identifier that represents the title sort order.
	ID3MetadataTitleSortOrder AVMetadataIdentifier
	// ID3MetadataTrackNumber: An identifier that represents the order number of the audio file.
	ID3MetadataTrackNumber AVMetadataIdentifier
	// ID3MetadataUniqueFileIdentifier: An identifier that represents the identifier used to indicate the audio file in a database.
	ID3MetadataUniqueFileIdentifier AVMetadataIdentifier
	// ID3MetadataUnsynchronizedLyric: An identifier that represents the lyrics of the song or a text transcription of other vocal activities.
	ID3MetadataUnsynchronizedLyric AVMetadataIdentifier
	// ID3MetadataUserText: An identifier that represents the user text information frame.
	ID3MetadataUserText AVMetadataIdentifier
	// ID3MetadataUserURL: An identifier that represents the user webpage frame.
	ID3MetadataUserURL AVMetadataIdentifier
	// ID3MetadataYear: An identifier that represents the year of the recording.
	ID3MetadataYear AVMetadataIdentifier
	// ISOUserDataAccessibilityDescription: An identifier that represents the accessibility description for the media content.
	ISOUserDataAccessibilityDescription AVMetadataIdentifier
	// ISOUserDataCopyright: An identifier that represents the copyright statement.
	ISOUserDataCopyright AVMetadataIdentifier
	// ISOUserDataDate: An identifier that represents the date for the media content.
	ISOUserDataDate AVMetadataIdentifier
	// ISOUserDataTaggedCharacteristic: An identifier that represents the tagged media characteristic used for identifying accessibility features.
	ISOUserDataTaggedCharacteristic AVMetadataIdentifier
	// IcyMetadataStreamTitle: An identifier that represents the title of a stream.
	IcyMetadataStreamTitle AVMetadataIdentifier
	// IcyMetadataStreamURL: An identifier that represents the web address of a stream.
	IcyMetadataStreamURL AVMetadataIdentifier
	// QuickTimeMetadataAIMEData: A value of type kCMMetadataBaseDataType_RawData
	QuickTimeMetadataAIMEData AVMetadataIdentifier
	// QuickTimeMetadataAccessibilityDescription: An identifier that represents the accessibility description for the movie file content.
	QuickTimeMetadataAccessibilityDescription AVMetadataIdentifier
	// QuickTimeMetadataAlbum: An identifier that represents the name of the album or collection in QuickTime.
	QuickTimeMetadataAlbum AVMetadataIdentifier
	// QuickTimeMetadataArranger: An identifier that represents the name of the arranger of the movie file content.
	QuickTimeMetadataArranger AVMetadataIdentifier
	// QuickTimeMetadataArtist: An identifier that represents the name of the artist of the movie file content.
	QuickTimeMetadataArtist AVMetadataIdentifier
	// QuickTimeMetadataArtwork: An identifier that represents an image relating to the movie file content.
	QuickTimeMetadataArtwork AVMetadataIdentifier
	// QuickTimeMetadataAuthor: An identifier that represents the name of the author of the movie file content.
	QuickTimeMetadataAuthor AVMetadataIdentifier
	// QuickTimeMetadataAutoLivePhoto: An identifier that represents whether the live photo movie used auto mode.
	QuickTimeMetadataAutoLivePhoto AVMetadataIdentifier
	// QuickTimeMetadataCameraFocalLength35mmEquivalent: A value of type kCMMetadataBaseDataType_UTF8 indicating focal length normalized to the 35mm film equivalent value (e.g. “50.00mm”).
	QuickTimeMetadataCameraFocalLength35mmEquivalent AVMetadataIdentifier
	// QuickTimeMetadataCameraFrameReadoutTime: An identifier that represents the camera frame readout time in QuickTime.
	QuickTimeMetadataCameraFrameReadoutTime AVMetadataIdentifier
	// QuickTimeMetadataCameraISOSensitivity: A value of type kCMMetadataBaseDataType_UTF8 indicating the sensitivity of the camera to light in terms of ISO exposure index (e.g. “800”). See SMPTE RDD 18.
	QuickTimeMetadataCameraISOSensitivity AVMetadataIdentifier
	// QuickTimeMetadataCameraIdentifier: An identifier that represents the camera identifier in QuickTime.
	QuickTimeMetadataCameraIdentifier AVMetadataIdentifier
	// QuickTimeMetadataCameraLensIrisFNumber: A value of type kCMMetadataBaseDataType_UTF8 indicating measure of the amount of light transmitted through the lens. It is the focal length divided by the effective lens aperture diameter (e.g. “F2.8” or “2.8”).
	QuickTimeMetadataCameraLensIrisFNumber AVMetadataIdentifier
	// QuickTimeMetadataCameraLensModel: A value of type kCMMetadataBaseDataType_UTF8 indicating the lens model (e.g. “iPhone 16 Pro back camera 6.765mm f/1.78”).
	QuickTimeMetadataCameraLensModel AVMetadataIdentifier
	// QuickTimeMetadataCameraShutterSpeedAngle: A value of type kCMMetadataBaseDataType_UTF8 indicating the exposure period expressed as an angle in minutes (1/60 degree) (e.g. “21600” or “360.00deg””).
	QuickTimeMetadataCameraShutterSpeedAngle AVMetadataIdentifier
	// QuickTimeMetadataCameraShutterSpeedTime: A value of type kCMMetadataBaseDataType_UTF8 indicating the exposure period expressed as a time per one frame/field period in seconds.
	QuickTimeMetadataCameraShutterSpeedTime AVMetadataIdentifier
	// QuickTimeMetadataCameraWhiteBalance: A value of type kCMMetadataBaseDataType_UTF8 indicating the white balance value defined by the temperature in Kelvin units (e.g. “5500K” or “5500”). See SMPTE RDD 18.
	QuickTimeMetadataCameraWhiteBalance AVMetadataIdentifier
	// QuickTimeMetadataCinematicVideoIntent: A value of type `kCMMetadataBaseDataType_UInt8` indicating whether this movie is intended as a Cinematic Video (1) or not (0).
	QuickTimeMetadataCinematicVideoIntent AVMetadataIdentifier
	// QuickTimeMetadataCollectionUser: An identifier that represents a name that indicates a user-defined collection.
	QuickTimeMetadataCollectionUser AVMetadataIdentifier
	// QuickTimeMetadataComment: An identifier that represents a comment regarding the movie file content.
	QuickTimeMetadataComment AVMetadataIdentifier
	// QuickTimeMetadataComposer: An identifier that represents the name of the composer of the movie file content.
	QuickTimeMetadataComposer AVMetadataIdentifier
	// QuickTimeMetadataContentIdentifier: An identifier that represents the content identifier in QuickTime.
	QuickTimeMetadataContentIdentifier AVMetadataIdentifier
	// QuickTimeMetadataCopyright: An identifier that represents the copyright statement for the movie file content.
	QuickTimeMetadataCopyright AVMetadataIdentifier
	// QuickTimeMetadataCreationDate: An identifier that represents the creation date of the movie file content.
	QuickTimeMetadataCreationDate AVMetadataIdentifier
	// QuickTimeMetadataCredits: An identifier that represents the credits of the movie source content.
	QuickTimeMetadataCredits AVMetadataIdentifier
	// QuickTimeMetadataDescription: An identifier that represents the description of the movie file content.
	QuickTimeMetadataDescription AVMetadataIdentifier
	// QuickTimeMetadataDetectedCatBody: An identifier that represents a detected cat body in the movie file content.
	QuickTimeMetadataDetectedCatBody AVMetadataIdentifier
	// QuickTimeMetadataDetectedDogBody: An identifier that represents a detected dog body in the movie file content.
	QuickTimeMetadataDetectedDogBody AVMetadataIdentifier
	// QuickTimeMetadataDetectedFace: An identifier that represents a detected face in the movie file content.
	QuickTimeMetadataDetectedFace AVMetadataIdentifier
	// QuickTimeMetadataDetectedHumanBody: An identifier that represents a detected human body in the movie file content.
	QuickTimeMetadataDetectedHumanBody AVMetadataIdentifier
	// QuickTimeMetadataDetectedSalientObject: An identifier that represents a detected salient object in the movie file content.
	QuickTimeMetadataDetectedSalientObject AVMetadataIdentifier
	// QuickTimeMetadataDirectionFacing: An identifier that represents the direction the camera is facing during the shot.
	QuickTimeMetadataDirectionFacing AVMetadataIdentifier
	// QuickTimeMetadataDirectionMotion: An identifier that represents the direction the camera is moving during the shot.
	QuickTimeMetadataDirectionMotion AVMetadataIdentifier
	// QuickTimeMetadataDirector: An identifier that represents the name of the director of the movie file content.
	QuickTimeMetadataDirector AVMetadataIdentifier
	// QuickTimeMetadataDisplayName: An identifier that represents the display name of the movie file content.
	QuickTimeMetadataDisplayName AVMetadataIdentifier
	// QuickTimeMetadataEncodedBy: An identifier that represents the name of the person or organization responsible for encoding the movie file content.
	QuickTimeMetadataEncodedBy AVMetadataIdentifier
	// QuickTimeMetadataFullFrameRatePlaybackIntent: An identifier that represents whether this movie should play at full frame rate.
	QuickTimeMetadataFullFrameRatePlaybackIntent AVMetadataIdentifier
	// QuickTimeMetadataGenre: An identifier that represents the genre or genres to which the movie content conforms.
	QuickTimeMetadataGenre AVMetadataIdentifier
	// QuickTimeMetadataInformation: An identifier that represents general information about the movie file content.
	QuickTimeMetadataInformation AVMetadataIdentifier
	// QuickTimeMetadataIsMontage: An identifier that represents that a movie is a montage of other media.
	QuickTimeMetadataIsMontage AVMetadataIdentifier
	// QuickTimeMetadataKeywords: An identifier that represents the keywords for the movie file content.
	QuickTimeMetadataKeywords AVMetadataIdentifier
	// QuickTimeMetadataLivePhotoVitalityScore: An identifier that represents the vitality score of the Live Photo movie.
	QuickTimeMetadataLivePhotoVitalityScore AVMetadataIdentifier
	// QuickTimeMetadataLivePhotoVitalityScoringVersion: An identifier that represents the version of the algorithm responsible for scoring the Live Photo movie for vitality.
	QuickTimeMetadataLivePhotoVitalityScoringVersion AVMetadataIdentifier
	// QuickTimeMetadataLocationBody: An identifier that represents the astronomical body for compatibility with the 3GPP format.
	QuickTimeMetadataLocationBody AVMetadataIdentifier
	// QuickTimeMetadataLocationDate: An identifier that represents a date and time using the extended format defined in ISO 8601:2004.
	QuickTimeMetadataLocationDate AVMetadataIdentifier
	// QuickTimeMetadataLocationHorizontalAccuracyInMeters: An identifier that represents the horizontal accuracy of the location data.
	QuickTimeMetadataLocationHorizontalAccuracyInMeters AVMetadataIdentifier
	// QuickTimeMetadataLocationISO6709: An identifier that represents the geographic point location by coordinates as defined in ISO 6709:2008.
	QuickTimeMetadataLocationISO6709 AVMetadataIdentifier
	// QuickTimeMetadataLocationName: An identifier that represents the name of the location.
	QuickTimeMetadataLocationName AVMetadataIdentifier
	// QuickTimeMetadataLocationNote: An identifier that represents a descriptive comment about the location.
	QuickTimeMetadataLocationNote AVMetadataIdentifier
	// QuickTimeMetadataLocationRole: An identifier that represents the single byte describing the movie location.
	QuickTimeMetadataLocationRole AVMetadataIdentifier
	// QuickTimeMetadataMake: An identifier that represents the name of the camera maker.
	QuickTimeMetadataMake AVMetadataIdentifier
	// QuickTimeMetadataModel: An identifier that represents the name of the camera model.
	QuickTimeMetadataModel AVMetadataIdentifier
	// QuickTimeMetadataOriginalArtist: An identifier that represents the name of the original artist of the movie file content.
	QuickTimeMetadataOriginalArtist AVMetadataIdentifier
	// QuickTimeMetadataPerformer: An identifier that represents the name of the performer in the movie file content.
	QuickTimeMetadataPerformer AVMetadataIdentifier
	// QuickTimeMetadataPhonogramRights: An identifier that represents the phonogram rights statement.
	QuickTimeMetadataPhonogramRights AVMetadataIdentifier
	// QuickTimeMetadataPreferredAffineTransform: An identifier that represents the affine transform preference for the movie file content.
	QuickTimeMetadataPreferredAffineTransform AVMetadataIdentifier
	// QuickTimeMetadataPresentationImmersiveMedia: A value of type kCMMetadataBaseDataType_RawData
	QuickTimeMetadataPresentationImmersiveMedia AVMetadataIdentifier
	// QuickTimeMetadataProducer: An identifier that represents the name of the producer of the movie file content.
	QuickTimeMetadataProducer AVMetadataIdentifier
	// QuickTimeMetadataPublisher: An identifier that represents the name of the publisher of the movie file content.
	QuickTimeMetadataPublisher AVMetadataIdentifier
	// QuickTimeMetadataRatingUser: An identifier that represents the rating or relative value of the movie.
	QuickTimeMetadataRatingUser AVMetadataIdentifier
	// QuickTimeMetadataSoftware: An identifier that represents the name of software used to create the movie file content.
	QuickTimeMetadataSoftware AVMetadataIdentifier
	// QuickTimeMetadataSpatialOverCaptureQualityScore: An identifier that represents a score that indicates the quality of an asset.
	QuickTimeMetadataSpatialOverCaptureQualityScore AVMetadataIdentifier
	// QuickTimeMetadataSpatialOverCaptureQualityScoringVersion: An identifier that represents the version of the algorithm responsible for generating a score.
	QuickTimeMetadataSpatialOverCaptureQualityScoringVersion AVMetadataIdentifier
	// QuickTimeMetadataTitle: An identifier that represents the title of the movie file content.
	QuickTimeMetadataTitle AVMetadataIdentifier
	// QuickTimeMetadataVideoOrientation: An identifier that represents the orientation of the movie file content.
	QuickTimeMetadataVideoOrientation AVMetadataIdentifier
	// QuickTimeMetadataWhiteBalanceByCCTColorMatrices: A value of type kCMMetadataBaseDataType_RawData indicating the reference color translation matrix data for ProRes RAW.
	QuickTimeMetadataWhiteBalanceByCCTColorMatrices AVMetadataIdentifier
	// QuickTimeMetadataWhiteBalanceByCCTWhiteBalanceFactors: A value of type kCMMetadataBaseDataType_RawData indicating the reference white balance multiplication factor data for ProRes RAW.
	QuickTimeMetadataWhiteBalanceByCCTWhiteBalanceFactors AVMetadataIdentifier
	// QuickTimeMetadataYear: An identifier that represents the recording year of the movie file content.
	QuickTimeMetadataYear AVMetadataIdentifier
	// QuickTimeMetadataiXML: An identifier that represents iXML information for the movie file content.
	QuickTimeMetadataiXML AVMetadataIdentifier
	// QuickTimeUserDataAccessibilityDescription: An identifier that represents the accessibility description for the movie file content.
	QuickTimeUserDataAccessibilityDescription AVMetadataIdentifier
	// QuickTimeUserDataAlbum: An identifier that represents the name of the album or collection in QuickTime.
	QuickTimeUserDataAlbum AVMetadataIdentifier
	// QuickTimeUserDataArranger: An identifier that represents the name of the arranger of the movie file content.
	QuickTimeUserDataArranger AVMetadataIdentifier
	// QuickTimeUserDataArtist: An identifier that represents the name of the artist of the movie file content.
	QuickTimeUserDataArtist AVMetadataIdentifier
	// QuickTimeUserDataAuthor: An identifier that represents the name of the author of the movie file content.
	QuickTimeUserDataAuthor AVMetadataIdentifier
	// QuickTimeUserDataChapter: An identifier that represents the name of the chapter.
	QuickTimeUserDataChapter AVMetadataIdentifier
	// QuickTimeUserDataComment: An identifier that represents a comment regarding the movie file content.
	QuickTimeUserDataComment AVMetadataIdentifier
	// QuickTimeUserDataComposer: An identifier that represents the name of the composer of the movie file content.
	QuickTimeUserDataComposer AVMetadataIdentifier
	// QuickTimeUserDataCopyright: An identifier that represents the copyright statement in QuickTime.
	QuickTimeUserDataCopyright AVMetadataIdentifier
	// QuickTimeUserDataCreationDate: An identifier that represents the creation date of the movie file content.
	QuickTimeUserDataCreationDate AVMetadataIdentifier
	// QuickTimeUserDataCredits: An identifier that represents the credits of movie source content.
	QuickTimeUserDataCredits AVMetadataIdentifier
	// QuickTimeUserDataDescription: An identifier that represents the description of the movie file content.
	QuickTimeUserDataDescription AVMetadataIdentifier
	// QuickTimeUserDataDirector: An identifier that represents the name of the director of the movie file content.
	QuickTimeUserDataDirector AVMetadataIdentifier
	// QuickTimeUserDataDisclaimer: An identifier that represents the disclaimer regarding the movie file content.
	QuickTimeUserDataDisclaimer AVMetadataIdentifier
	// QuickTimeUserDataEncodedBy: An identifier that represents the name of the person or organization responsible for encoding the movie file content.
	QuickTimeUserDataEncodedBy AVMetadataIdentifier
	// QuickTimeUserDataFullName: An identifier that represents the full name of the movie file content.
	QuickTimeUserDataFullName AVMetadataIdentifier
	// QuickTimeUserDataGenre: An identifier that represents the genre or genres to which the movie content conforms.
	QuickTimeUserDataGenre AVMetadataIdentifier
	// QuickTimeUserDataHostComputer: An identifier that represents the name of the host computer.
	QuickTimeUserDataHostComputer AVMetadataIdentifier
	// QuickTimeUserDataInformation: An identifier that represents general information about the movie file content.
	QuickTimeUserDataInformation AVMetadataIdentifier
	// QuickTimeUserDataKeywords: An identifier that represents the keywords for the movie file content.
	QuickTimeUserDataKeywords AVMetadataIdentifier
	// QuickTimeUserDataLocationISO6709: An identifier that represents the geographic point location by coordinates as defined in ISO 6709:2008.
	QuickTimeUserDataLocationISO6709 AVMetadataIdentifier
	// QuickTimeUserDataMake: An identifier that represents the name of the camera maker.
	QuickTimeUserDataMake AVMetadataIdentifier
	// QuickTimeUserDataModel: An identifier that represents the name of the camera model.
	QuickTimeUserDataModel AVMetadataIdentifier
	// QuickTimeUserDataOriginalArtist: An identifier that represents the name of the original artist of the movie file content.
	QuickTimeUserDataOriginalArtist AVMetadataIdentifier
	// QuickTimeUserDataOriginalFormat: An identifier that represents the original format of the movie file content.
	QuickTimeUserDataOriginalFormat AVMetadataIdentifier
	// QuickTimeUserDataOriginalSource: An identifier that represents the original source of the movie file content.
	QuickTimeUserDataOriginalSource AVMetadataIdentifier
	// QuickTimeUserDataPerformers: An identifier that represents the name of the performers in the movie file content.
	QuickTimeUserDataPerformers AVMetadataIdentifier
	// QuickTimeUserDataPhonogramRights: An identifier that represents the phonogram rights statement.
	QuickTimeUserDataPhonogramRights AVMetadataIdentifier
	// QuickTimeUserDataProducer: An identifier that represents the name of the producer of the movie file content.
	QuickTimeUserDataProducer AVMetadataIdentifier
	// QuickTimeUserDataProduct: An identifier that represents the name of the product.
	QuickTimeUserDataProduct AVMetadataIdentifier
	// QuickTimeUserDataPublisher: An identifier that represents the name of the publisher of the movie file content.
	QuickTimeUserDataPublisher AVMetadataIdentifier
	// QuickTimeUserDataSoftware: An identifier that represents the name of software used to create the movie file content.
	QuickTimeUserDataSoftware AVMetadataIdentifier
	// QuickTimeUserDataSpecialPlaybackRequirements: An identifier that represents the special hardware and software requirements for playback.
	QuickTimeUserDataSpecialPlaybackRequirements AVMetadataIdentifier
	// QuickTimeUserDataTaggedCharacteristic: An identifier that represents the tagged characteristic.
	QuickTimeUserDataTaggedCharacteristic AVMetadataIdentifier
	// QuickTimeUserDataTrack: An identifier that represents a track in the movie file content.
	QuickTimeUserDataTrack AVMetadataIdentifier
	// QuickTimeUserDataTrackName: An identifier that represents the name of a track in the movie file content.
	QuickTimeUserDataTrackName AVMetadataIdentifier
	// QuickTimeUserDataURLLink: An identifier that represents the webpage for the movie file content.
	QuickTimeUserDataURLLink AVMetadataIdentifier
	// QuickTimeUserDataWarning: An identifier that represents the warning text for the movie file content.
	QuickTimeUserDataWarning AVMetadataIdentifier
	// QuickTimeUserDataWriter: An identifier that represents the name of the writer of the movie file content.
	QuickTimeUserDataWriter AVMetadataIdentifier
}

// AVMetadataKeySpaces provides typed accessors for [AVMetadataKeySpace] constants.
var AVMetadataKeySpaces struct {
	// AudioFile: The AudioToolbox audio file key space.
	AudioFile AVMetadataKeySpace
	// Common: The common key space.
	Common AVMetadataKeySpace
	// HLSDateRange: The HTTP Live Streaming key space.
	HLSDateRange AVMetadataKeySpace
	// ID3: The ID3 key space.
	ID3 AVMetadataKeySpace
	// ISOUserData: The ISO key space.
	ISOUserData AVMetadataKeySpace
	// Icy: The Icecast/ShoutCAST streaming key space.
	Icy AVMetadataKeySpace
	// QuickTimeMetadata: The QuickTime metadata key space.
	QuickTimeMetadata AVMetadataKeySpace
	// QuickTimeUserData: The QuickTime user data key space.
	QuickTimeUserData AVMetadataKeySpace
}

// AVMetadataObjectTypes provides typed accessors for [AVMetadataObjectType] constants.
var AVMetadataObjectTypes struct {
	// AztecCode: A constant that identifies the Aztec symbology.
	AztecCode AVMetadataObjectType
	// CatBody: A constant that identifies cat body metadata.
	CatBody AVMetadataObjectType
	// CatHead: An identifier for an instance of a cat head object.
	CatHead AVMetadataObjectType
	// CodabarCode: A constant that identifies the Codabar symbology.
	CodabarCode AVMetadataObjectType
	// Code128Code: A constant that identifies the Code 128 symbology.
	Code128Code AVMetadataObjectType
	// Code39Code: A constant that identifies the Code 39 symbology.
	Code39Code AVMetadataObjectType
	// Code39Mod43Code: A constant that identifies the Code 39 mod 43 symbology.
	Code39Mod43Code AVMetadataObjectType
	// Code93Code: A constant that identifies the Code 93 symbology.
	Code93Code AVMetadataObjectType
	// DataMatrixCode: A constant that identifies the DataMatrix symbology.
	DataMatrixCode AVMetadataObjectType
	// DogBody: A constant that identifies dog body metadata.
	DogBody AVMetadataObjectType
	// DogHead: An identifier for an instance of a dog head object.
	DogHead AVMetadataObjectType
	// EAN13Code: A constant that identifies the EAN-13 symbology.
	EAN13Code AVMetadataObjectType
	// EAN8Code: A constant that identifies the EAN-8 symbology.
	EAN8Code AVMetadataObjectType
	// Face: A constant that identifies face metadata.
	Face AVMetadataObjectType
	// GS1DataBarCode: A constant that identifies the GS1 DataBar symbology.
	GS1DataBarCode AVMetadataObjectType
	// GS1DataBarExpandedCode: A constant that identifies the GS1 DataBar Expanded symbology.
	GS1DataBarExpandedCode AVMetadataObjectType
	// GS1DataBarLimitedCode: A constant that identifies the GS1 DataBar Limited symbology.
	GS1DataBarLimitedCode AVMetadataObjectType
	// HumanBody: A constant that identifies human body metadata.
	HumanBody AVMetadataObjectType
	// HumanFullBody: A constant that identifies human full body metadata.
	HumanFullBody AVMetadataObjectType
	// ITF14Code: A constant that identifies the ITF14 symbology.
	ITF14Code AVMetadataObjectType
	// Interleaved2of5Code: A constant that identifies the Interleaved 2 of 5 symbology.
	Interleaved2of5Code AVMetadataObjectType
	// MicroPDF417Code: A constant that identifies the Micro PDF417 symbology.
	MicroPDF417Code AVMetadataObjectType
	// MicroQRCode: A constant that identifies the Micro QR symbology.
	MicroQRCode AVMetadataObjectType
	// PDF417Code: A constant that identifies the PDF417 symbology.
	PDF417Code AVMetadataObjectType
	// QRCode: A constant that identifies the QR symbology.
	QRCode AVMetadataObjectType
	// SalientObject: A constant that identifies saliency metadata.
	SalientObject AVMetadataObjectType
	// UPCECode: A constant that identifies the UPC-E symbology.
	UPCECode AVMetadataObjectType
}

// AVOutputSettingsPresets provides typed accessors for [AVOutputSettingsPreset] constants.
var AVOutputSettingsPresets struct {
	// HEVC1920x1080: A preset for HEVC video at 1920 by 1080 pixels.
	HEVC1920x1080 AVOutputSettingsPreset
	// HEVC1920x1080WithAlpha: A preset for HEVC with Alpha video at 1920 by 1080 pixels.
	HEVC1920x1080WithAlpha AVOutputSettingsPreset
	// HEVC3840x2160: A preset for HEVC video at 3840 by 2160 pixels.
	HEVC3840x2160 AVOutputSettingsPreset
	// HEVC3840x2160WithAlpha: A preset for HEVC with Alpha video at 3840 by 2160 pixels.
	HEVC3840x2160WithAlpha AVOutputSettingsPreset
	HEVC4320x2160          AVOutputSettingsPreset
	// HEVC7680x4320: A preset for HEVC video at 7680 by 4320 pixels.
	HEVC7680x4320 AVOutputSettingsPreset
	// MVHEVC1440x1440: A preset for MV-HEVC video at 1440 by 1440 pixels.
	MVHEVC1440x1440 AVOutputSettingsPreset
	MVHEVC4320x4320 AVOutputSettingsPreset
	MVHEVC7680x7680 AVOutputSettingsPreset
	// MVHEVC960x960: A preset for MV-HEVC video at 960 by 960 pixels.
	MVHEVC960x960 AVOutputSettingsPreset
}

// AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasons provides typed accessors for [AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason] constants.
var AVPlayerIntegratedTimelineSnapshotsOutOfSyncReasons struct {
	// CurrentSegmentChanged: The snapshot is out of sync due to a change of the current segment.
	CurrentSegmentChanged AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason
	// LoadedTimeRangesChanged: The snapshot is out of sync due to a change of the loaded time ranges.
	LoadedTimeRangesChanged AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason
	// SegmentsChanged: The snapshot is out of sync due to a change of segments.
	SegmentsChanged AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason
}

// AVPlayerItemLegibleOutputTextStylingResolutions provides typed accessors for [AVPlayerItemLegibleOutputTextStylingResolution] constants.
var AVPlayerItemLegibleOutputTextStylingResolutions struct {
	// Default: The text styling information is the same level of information that AVFoundation uses within a player layer.
	Default AVPlayerItemLegibleOutputTextStylingResolution
	// SourceAndRulesOnly: The level of resolution excludes styling provided by the user-level Media Accessibility settings.
	SourceAndRulesOnly AVPlayerItemLegibleOutputTextStylingResolution
}

// AVPlayerRateDidChangeReasons provides typed accessors for [AVPlayerRateDidChangeReason] constants.
var AVPlayerRateDidChangeReasons struct {
	// AppBackgrounded: An app transitions to the background.
	AppBackgrounded AVPlayerRateDidChangeReason
	// AudioSessionInterrupted: The system interrupts the app’s audio session.
	AudioSessionInterrupted                    AVPlayerRateDidChangeReason
	PlayheadReachedLiveEdge                    AVPlayerRateDidChangeReason
	ReversePlaybackReachedStartOfSeekableRange AVPlayerRateDidChangeReason
	// SetRateCalled: An app makes a call to set the player’s rate.
	SetRateCalled AVPlayerRateDidChangeReason
	// SetRateFailed: An attempt to change the player’s rate fails.
	SetRateFailed AVPlayerRateDidChangeReason
}

// AVSemanticSegmentationMatteTypes provides typed accessors for [AVSemanticSegmentationMatteType] constants.
var AVSemanticSegmentationMatteTypes struct {
	// Glasses: A matting image that segments eyeglasses and sunglasses from all people in the visible field of view of an image.
	Glasses AVSemanticSegmentationMatteType
	// Hair: A matting image that segments the hair from all people in the visible field of view of an image.
	Hair AVSemanticSegmentationMatteType
	// Skin: A matting image that segments the skin from all people in the visible field of view of an image.
	Skin AVSemanticSegmentationMatteType
	// Teeth: A matting image that segments the teeth from all people in the visible field of view of an image.
	Teeth AVSemanticSegmentationMatteType
}

// AVSpatialCaptureDiscomfortReasons provides typed accessors for [AVSpatialCaptureDiscomfortReason] constants.
var AVSpatialCaptureDiscomfortReasons struct {
	// NotEnoughLight: A value that indicates the lighting of the current scene isn’t bright enough.
	NotEnoughLight AVSpatialCaptureDiscomfortReason
	// SubjectTooClose: A value that indicates the focus point of the current scene is too close.
	SubjectTooClose AVSpatialCaptureDiscomfortReason
}

// AVTrackAssociationTypes provides typed accessors for [AVTrackAssociationType] constants.
var AVTrackAssociationTypes struct {
	// AudioFallback: The track contains the same content as another track, but in a more widely supported format.
	AudioFallback AVTrackAssociationType
	// ChapterList: The associated track contains chapter information for the base track.
	ChapterList AVTrackAssociationType
	// ForcedSubtitlesOnly: An association between a subtitle track containing forced and nonforced subtitles and one with a subtitle track containing only forced subtitles.
	ForcedSubtitlesOnly AVTrackAssociationType
	// MetadataReferent: An association between a metadata track and the track that it describes or annotates.
	MetadataReferent AVTrackAssociationType
	// RenderMetadataSource: Indicates an association between a metadata track and another track where the metadata provides additional information for rendering of that track.
	RenderMetadataSource AVTrackAssociationType
	// SelectionFollower: An association between two tracks that specifies when a user selects the main track, the system should follow that selection by automatically selecting the associated track.
	SelectionFollower AVTrackAssociationType
	// Timecode: An association between a timecode track providing timing information for the main track.
	Timecode AVTrackAssociationType
}

// AVVideoApertureModes provides typed accessors for [AVVideoApertureMode] constants.
var AVVideoApertureModes struct {
	// CleanAperture: The pixel aspect ratio and clean aperture will be applied.
	CleanAperture AVVideoApertureMode
	// EncodedPixels: The encoded dimensions of the image description are displayed.
	EncodedPixels AVVideoApertureMode
	// ProductionAperture: The pixel aspect ratio will be applied.
	ProductionAperture AVVideoApertureMode
}

// AVVideoCodecTypes provides typed accessors for [AVVideoCodecType] constants.
var AVVideoCodecTypes struct {
	// AppleProRes422: The Apple ProRes 422 video codec.
	AppleProRes422 AVVideoCodecType
	// AppleProRes422HQ: The Apple ProRes 422 HQ video codec.
	AppleProRes422HQ AVVideoCodecType
	// AppleProRes422LT: The Apple ProRes 422 LT video codec.
	AppleProRes422LT AVVideoCodecType
	// AppleProRes422Proxy: The Apple ProRes 422 Proxy video codec.
	AppleProRes422Proxy AVVideoCodecType
	// AppleProRes4444: The Apple ProRes 4444 video codec.
	AppleProRes4444 AVVideoCodecType
	// AppleProRes4444XQ: The Apple ProRes 4444 XQ video codec.
	AppleProRes4444XQ AVVideoCodecType
	AppleProResRAW    AVVideoCodecType
	AppleProResRAWHQ  AVVideoCodecType
	// H264: The H.264 video codec.
	H264 AVVideoCodecType
	// HEVC: The HEVC video codec.
	HEVC AVVideoCodecType
	// HEVCWithAlpha: The HEVC video codec that supports an alpha channel.
	HEVCWithAlpha AVVideoCodecType
	// JPEG: The JPEG video codec.
	JPEG AVVideoCodecType
	// JPEGXL: The JPEG XL video codec.
	JPEGXL AVVideoCodecType
}

// AVVideoCompositionPerFrameHDRDisplayMetadataPolicys provides typed accessors for [AVVideoCompositionPerFrameHDRDisplayMetadataPolicy] constants.
var AVVideoCompositionPerFrameHDRDisplayMetadataPolicys struct {
	// Generate: A video composition may generate HDR metadata and attach it to the rendered frame.
	Generate AVVideoCompositionPerFrameHDRDisplayMetadataPolicy
	// Propagate: A policy that passes HDR metadata through, if present on the composed frame.
	Propagate AVVideoCompositionPerFrameHDRDisplayMetadataPolicy
}

// AVVideoRanges provides typed accessors for [AVVideoRange] constants.
var AVVideoRanges struct {
	// HLG: Indicates Hybrid-Log Gamma (HLG) high-dynamic-range video.
	HLG AVVideoRange
	// PQ: Indicates Perceptual Quantizer (PQ) high-dynamic-range video.
	PQ AVVideoRange
	// SDR: Indicates standard-dynamic-range (SDR) video.
	SDR AVVideoRange
}
