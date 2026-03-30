// Code generated from Apple documentation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
)

// AVAssetDownloadedAssetEvictionPriority is constants that define eviction priorities for a storage management policy.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetDownloadedAssetEvictionPriority
type AVAssetDownloadedAssetEvictionPriority = string

// AVAssetImageGeneratorApertureMode is constants that define aperture modes to use when generating images.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/ApertureMode-swift.struct
type AVAssetImageGeneratorApertureMode = string

// AVAssetImageGeneratorCompletionHandler is a type alias for a closure that provides the result of an image generation request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGeneratorCompletionHandler
type AVAssetImageGeneratorCompletionHandler = func(coremedia.CMTime, *coregraphics.CGImageRef, coremedia.CMTime, AVAssetImageGeneratorResult, foundation.NSError)

// AVAssetImageGeneratorDynamicRangePolicy is a type that specifies the dynamic range policy to apply when generating images.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/DynamicRangePolicy-swift.struct
type AVAssetImageGeneratorDynamicRangePolicy = string

// AVAssetPlaybackConfigurationOption is a structure that defines playback configuration options for an asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetPlaybackConfigurationOption
type AVAssetPlaybackConfigurationOption = string

// AVAssetWriterInputMediaDataLocation is a structure that indicates how to lay out and interleave media data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriterInput/MediaDataLocation-swift.struct
type AVAssetWriterInputMediaDataLocation = string

// AVAudioTimePitchAlgorithm is an algorithm used to set the audio pitch as the rate changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAudioTimePitchAlgorithm
type AVAudioTimePitchAlgorithm = string

// AVCaptionConversionAdjustmentType is constants that indicate an adjustment type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionAdjustment/AdjustmentType-swift.struct
type AVCaptionConversionAdjustmentType = string

// AVCaptionConversionWarningType is the type of a caption conversion warning.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionWarning/WarningType-swift.struct
type AVCaptionConversionWarningType = string

// AVCaptionSettingsKey is a structure that defines dictionary keys to configure the caption converter and validator.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSettingsKey
type AVCaptionSettingsKey = string

// AVCaptureDeviceTransportControlsSpeed is a constant that specifies speed of transport controls.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TransportControlsSpeed-swift.typealias
type AVCaptureDeviceTransportControlsSpeed = float32

// AVCaptureDeviceType is a structure that defines the device types the framework supports.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct
type AVCaptureDeviceType = string

// AVCaptureReactionType is constants that indicate the type of reaction that an effect can perform.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionType
type AVCaptureReactionType = string

// AVCaptureSceneMonitoringStatus is an informative status about the scene observed by the device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSceneMonitoringStatus
type AVCaptureSceneMonitoringStatus = string

// AVCaptureSessionPreset is presets that define standard configurations for a capture session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset
type AVCaptureSessionPreset = string

// AVContentKeyRequestRetryReason is the reason for asking the client to retry a content key request.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/RetryReason
type AVContentKeyRequestRetryReason = string

// AVContentKeySessionServerPlaybackContextOption is options for specifying additional information for generating server playback context (SPC).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySessionServerPlaybackContextOption
type AVContentKeySessionServerPlaybackContextOption = string

// AVContentKeySystem is a key-delivery method for a content key session.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeySystem
type AVContentKeySystem = string

// AVCoordinatedPlaybackSuspensionReason is constants that identify playback suspension reasons.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCoordinatedPlaybackSuspension/Reason-swift.struct
type AVCoordinatedPlaybackSuspensionReason = string

// AVFileType is the uniform type identifiers for various file formats.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFileType
type AVFileType = string

// AVFileTypeProfile is file type profiles for streaming formats.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFileTypeProfile
type AVFileTypeProfile = string

// AVLayerVideoGravity is a structure that defines how a layer displays a player’s visual content within the layer’s bounds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity
type AVLayerVideoGravity = string

// AVMediaCharacteristic is a structure that defines media data characteristics.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic
type AVMediaCharacteristic = string

// AVMediaType is an identifier for various media types.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaType
type AVMediaType = string

// AVMetadataExtraAttributeKey is a structure that defines keys for extra metadata attributes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataExtraAttributeKey
type AVMetadataExtraAttributeKey = string

// AVMetadataFormat is a structure that defines metadata formats.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataFormat
type AVMetadataFormat = string

// AVMetadataIdentifier is a structure that defines identifiers for metadata formats.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier
type AVMetadataIdentifier = string

// AVMetadataKey is a structure that defines a metadata key.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey
type AVMetadataKey = string

// AVMetadataKeySpace is a structure that defines a metadata key space.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataKeySpace
type AVMetadataKeySpace = string

// AVMetadataObjectType is constants that identify metadata object types.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMetadataObject/ObjectType
type AVMetadataObjectType = string

// AVOutputSettingsPreset is a structure that defines preset configurations for an output settings assistant.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVOutputSettingsPreset
type AVOutputSettingsPreset = string

// AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason is constants that represent the reason for an out-of-sync state.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason
type AVPlayerIntegratedTimelineSnapshotsOutOfSyncReason = string

// AVPlayerInterstitialEventCue is a structure that defines standard cues to play interstitial content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Cue-swift.struct
type AVPlayerInterstitialEventCue = string

// AVPlayerItemLegibleOutputTextStylingResolution is a text styling resolution.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemLegibleOutput/TextStylingResolution-swift.struct
type AVPlayerItemLegibleOutputTextStylingResolution = string

// AVPlayerRateDidChangeReason is a structure that represents a rate change reason.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/RateDidChangeReason
type AVPlayerRateDidChangeReason = string

// AVPlayerWaitingReason is the reasons a player is waiting to begin or resume playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/WaitingReason
type AVPlayerWaitingReason = string

// AVSemanticSegmentationMatteType is a structure that defines the types of segmentation matte images that you can capture along with the primary image.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSemanticSegmentationMatte/MatteType-swift.struct
type AVSemanticSegmentationMatteType = string

// AVSpatialCaptureDiscomfortReason is constants that indicate the suitability of the current scene to create a comfortable viewing experience.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSpatialCaptureDiscomfortReason
type AVSpatialCaptureDiscomfortReason = string

// AVTrackAssociationType is constants that define track association types.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/AssociationType
type AVTrackAssociationType = string

// AVVideoApertureMode is a value that describes how a video is scaled or cropped.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoApertureMode
type AVVideoApertureMode = string

// AVVideoCodecType is a set of constants that describe the codecs the system supports for video capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCodecType
type AVVideoCodecType = string

// AVVideoCompositionPerFrameHDRDisplayMetadataPolicy is a type that defines the policy for handling of per frame HDR metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoComposition/PerFrameHDRDisplayMetadataPolicy-swift.struct
type AVVideoCompositionPerFrameHDRDisplayMetadataPolicy = string

// AVVideoRange is constants that describe a video variant’s dynamic range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoRange
type AVVideoRange = string
