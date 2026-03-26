// Code generated from Apple documentation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// AVAssetTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAsset.LoadTrackWithTrackIDCompletionHandler]
//   - [AVURLAsset.FindCompatibleTrackForCompositionTrackCompletionHandler]
type AVAssetTrackErrorHandler = func(*AVAssetTrack, error)

// NewAVAssetTrackErrorBlock wraps a Go [AVAssetTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAsset.LoadTrackWithTrackIDCompletionHandler]
//   - [AVURLAsset.FindCompatibleTrackForCompositionTrackCompletionHandler]
func NewAVAssetTrackErrorBlock(handler AVAssetTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVAssetTrack
		if resultID != 0 {
			v := AVAssetTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAssetTrackSegmentErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - segment: The loaded track segment, or `nil` if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAssetTrack.LoadSegmentForTrackTimeCompletionHandler]
type AVAssetTrackSegmentErrorHandler = func(*AVAssetTrackSegment, error)

// NewAVAssetTrackSegmentErrorBlock wraps a Go [AVAssetTrackSegmentErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetTrack.LoadSegmentForTrackTimeCompletionHandler]
func NewAVAssetTrackSegmentErrorBlock(handler AVAssetTrackSegmentErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVAssetTrackSegment
		if resultID != 0 {
			v := AVAssetTrackSegmentFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVAsynchronousCIImageFilteringRequestHandler handles A block that AVFoundation calls when processing each video frame.
//   - request: An [AVAsynchronousCIImageFilteringRequest](<doc://com.apple.avfoundation/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest>) object representing the frame to be processed.
//
// Used by:
//   - [AVMutableVideoComposition.VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler]
//   - [AVMutableVideoComposition.VideoCompositionWithAssetApplyingCIFiltersWithHandler]
//   - [AVVideoComposition.VideoCompositionWithAssetApplyingCIFiltersWithHandler]
type AVAsynchronousCIImageFilteringRequestHandler = func(*objc.ID)

// AVCaptionConversionWarningHandler handles The callback the system invokes when it finishes validation.
//
// Used by:
//   - [AVCaptionConversionValidator.ValidateCaptionConversionWithWarningHandler]
type AVCaptionConversionWarningHandler = func(*AVCaptionConversionWarning)

// NewAVCaptionConversionWarningBlock wraps a Go [AVCaptionConversionWarningHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVCaptionConversionValidator.ValidateCaptionConversionWithWarningHandler]
func NewAVCaptionConversionWarningBlock(handler AVCaptionConversionWarningHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *AVCaptionConversionWarning
		if resultID != 0 {
			v := AVCaptionConversionWarningFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVCompositionTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVComposition.LoadTrackWithTrackIDCompletionHandler]
type AVCompositionTrackErrorHandler = func(*AVCompositionTrack, error)

// NewAVCompositionTrackErrorBlock wraps a Go [AVCompositionTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVComposition.LoadTrackWithTrackIDCompletionHandler]
func NewAVCompositionTrackErrorBlock(handler AVCompositionTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVCompositionTrack
		if resultID != 0 {
			v := AVCompositionTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVFragmentedAssetTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVFragmentedAsset.LoadTrackWithTrackIDCompletionHandler]
type AVFragmentedAssetTrackErrorHandler = func(*AVFragmentedAssetTrack, error)

// NewAVFragmentedAssetTrackErrorBlock wraps a Go [AVFragmentedAssetTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVFragmentedAsset.LoadTrackWithTrackIDCompletionHandler]
func NewAVFragmentedAssetTrackErrorBlock(handler AVFragmentedAssetTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVFragmentedAssetTrack
		if resultID != 0 {
			v := AVFragmentedAssetTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVFragmentedMovieTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVFragmentedMovie.LoadTrackWithTrackIDCompletionHandler]
type AVFragmentedMovieTrackErrorHandler = func(*AVFragmentedMovieTrack, error)

// NewAVFragmentedMovieTrackErrorBlock wraps a Go [AVFragmentedMovieTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVFragmentedMovie.LoadTrackWithTrackIDCompletionHandler]
func NewAVFragmentedMovieTrackErrorBlock(handler AVFragmentedMovieTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVFragmentedMovieTrack
		if resultID != 0 {
			v := AVFragmentedMovieTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMediaSelectionGroupErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - mediaSelectionGroup: The loaded media selection group, or `nil` if no group is available or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler]
type AVMediaSelectionGroupErrorHandler = func(*AVMediaSelectionGroup, error)

// NewAVMediaSelectionGroupErrorBlock wraps a Go [AVMediaSelectionGroupErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAsset.LoadMediaSelectionGroupForMediaCharacteristicCompletionHandler]
func NewAVMediaSelectionGroupErrorBlock(handler AVMediaSelectionGroupErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVMediaSelectionGroup
		if resultID != 0 {
			v := AVMediaSelectionGroupFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMetadataItemValueRequestHandler handles A callback the system invokes to load the value for the metadata item.
//
// Used by:
//   - [AVMetadataItem.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler]
//   - [AVMutableMetadataItem.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler]
type AVMetadataItemValueRequestHandler = func(*AVMetadataItemValueRequest)

// NewAVMetadataItemValueRequestBlock wraps a Go [AVMetadataItemValueRequestHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVMetadataItem.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler]
//   - [AVMutableMetadataItem.MetadataItemWithPropertiesOfMetadataItemValueLoadingHandler]
func NewAVMetadataItemValueRequestBlock(handler AVMetadataItemValueRequestHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *AVMetadataItemValueRequest
		if resultID != 0 {
			v := AVMetadataItemValueRequestFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMovieTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVMovie.LoadTrackWithTrackIDCompletionHandler]
type AVMovieTrackErrorHandler = func(*AVMovieTrack, error)

// NewAVMovieTrackErrorBlock wraps a Go [AVMovieTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVMovie.LoadTrackWithTrackIDCompletionHandler]
func NewAVMovieTrackErrorBlock(handler AVMovieTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVMovieTrack
		if resultID != 0 {
			v := AVMovieTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMutableCompositionTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVMutableComposition.LoadTrackWithTrackIDCompletionHandler]
type AVMutableCompositionTrackErrorHandler = func(*AVMutableCompositionTrack, error)

// NewAVMutableCompositionTrackErrorBlock wraps a Go [AVMutableCompositionTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVMutableComposition.LoadTrackWithTrackIDCompletionHandler]
func NewAVMutableCompositionTrackErrorBlock(handler AVMutableCompositionTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVMutableCompositionTrack
		if resultID != 0 {
			v := AVMutableCompositionTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMutableMovieTrackErrorHandler handles A callback that the system invokes after it finishes the loading request.
//   - track: The loaded track, or `nil` if no track with the specified identifier exists or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVMutableMovie.LoadTrackWithTrackIDCompletionHandler]
type AVMutableMovieTrackErrorHandler = func(*AVMutableMovieTrack, error)

// NewAVMutableMovieTrackErrorBlock wraps a Go [AVMutableMovieTrackErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVMutableMovie.LoadTrackWithTrackIDCompletionHandler]
func NewAVMutableMovieTrackErrorBlock(handler AVMutableMovieTrackErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVMutableMovieTrack
		if resultID != 0 {
			v := AVMutableMovieTrackFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVMutableVideoCompositionErrorHandler handles A callback the system invokes with the created video composition, or an error if a failure occurs.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVMutableVideoComposition.VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler]
//   - [AVMutableVideoComposition.VideoCompositionWithPropertiesOfAssetCompletionHandler]
//   - [AVMutableVideoComposition.VideoCompositionWithPropertiesOfAssetPrototypeInstructionCompletionHandler]
type AVMutableVideoCompositionErrorHandler = func(*AVMutableVideoComposition, error)

// NewAVMutableVideoCompositionErrorBlock wraps a Go [AVMutableVideoCompositionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVMutableVideoComposition.VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler]
//   - [AVMutableVideoComposition.VideoCompositionWithPropertiesOfAssetCompletionHandler]
//   - [AVMutableVideoComposition.VideoCompositionWithPropertiesOfAssetPrototypeInstructionCompletionHandler]
func NewAVMutableVideoCompositionErrorBlock(handler AVMutableVideoCompositionErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVMutableVideoComposition
		if resultID != 0 {
			v := AVMutableVideoCompositionFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVVideoCompositionErrorHandler handles A callback the system invokes with the created video composition, or an error if a failure occurs.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVVideoComposition.VideoCompositionWithPropertiesOfAssetCompletionHandler]
type AVVideoCompositionErrorHandler = func(*AVVideoComposition, error)

// NewAVVideoCompositionErrorBlock wraps a Go [AVVideoCompositionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVVideoComposition.VideoCompositionWithPropertiesOfAssetCompletionHandler]
func NewAVVideoCompositionErrorBlock(handler AVVideoCompositionErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *AVVideoComposition
		if resultID != 0 {
			v := AVVideoCompositionFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// AVVideoPerformanceMetricsHandler is the signature for a completion handler block.
//
// Used by:
//   - [AVSampleBufferVideoRenderer.LoadVideoPerformanceMetricsWithCompletionHandler]
type AVVideoPerformanceMetricsHandler = func(*AVVideoPerformanceMetrics)

// NewAVVideoPerformanceMetricsBlock wraps a Go [AVVideoPerformanceMetricsHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVSampleBufferVideoRenderer.LoadVideoPerformanceMetricsWithCompletionHandler]
func NewAVVideoPerformanceMetricsBlock(handler AVVideoPerformanceMetricsHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *AVVideoPerformanceMetrics
		if resultID != 0 {
			v := AVVideoPerformanceMetricsFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// ArrayErrorHandler handles A callback that the system invokes after it finishes the loading operation.
//   - tracks: An array of tracks, which may be empty if no tracks with the specified media type exist. The value is `nil` if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
//
// Used by:
//   - [AVAsset.LoadChapterMetadataGroupsBestMatchingPreferredLanguagesCompletionHandler]
//   - [AVAsset.LoadChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeysCompletionHandler]
//   - [AVAsset.LoadMetadataForFormatCompletionHandler]
//   - [AVAsset.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVAsset.LoadTracksWithMediaTypeCompletionHandler]
//   - [AVAssetTrack.LoadAssociatedTracksOfTypeCompletionHandler]
//   - [AVAssetTrack.LoadMetadataForFormatCompletionHandler]
//   - [AVComposition.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVComposition.LoadTracksWithMediaTypeCompletionHandler]
//   - [AVFragmentedAsset.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVFragmentedAsset.LoadTracksWithMediaTypeCompletionHandler]
//   - [AVFragmentedMovie.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVFragmentedMovie.LoadTracksWithMediaTypeCompletionHandler]
//   - [AVMovie.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVMovie.LoadTracksWithMediaTypeCompletionHandler]
//   - [AVMutableComposition.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVMutableComposition.LoadTracksWithMediaTypeCompletionHandler]
//   - [AVMutableMovie.LoadTracksWithMediaCharacteristicCompletionHandler]
//   - [AVMutableMovie.LoadTracksWithMediaTypeCompletionHandler]
type ArrayErrorHandler = func(*[]AVAssetTrack, error)

// BoolErrorHandler handles A completion block to be called on a serial dispatch queue after the photo output has finished preparing resources.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVCapturePhotoOutput.SetPreparedPhotoSettingsArrayCompletionHandler]
//   - [AVPlayerItem.RequestPlaybackRestrictionsAuthorization]
//   - [AVSampleBufferGenerator.NotifyOfDataReadyForSampleBufferCompletionHandler]
//   - [AVVideoComposition.DetermineValidityForAssetTimeRangeValidationDelegateCompletionHandler]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVCapturePhotoOutput.SetPreparedPhotoSettingsArrayCompletionHandler]
//   - [AVPlayerItem.RequestPlaybackRestrictionsAuthorization]
//   - [AVSampleBufferGenerator.NotifyOfDataReadyForSampleBufferCompletionHandler]
//   - [AVVideoComposition.DetermineValidityForAssetTimeRangeValidationDelegateCompletionHandler]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(primitiveVal, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// BoolHandler handles A callback the system passes a Boolean result when it determines the compatibility of the preset.
//   - finished: Indicates whether the seek operation completed.
//
// Used by:
//   - [AVAssetExportSession.DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler]
//   - [AVCaptureDevice.RequestAccessForMediaTypeCompletionHandler]
//   - [AVExternalStorageDevice.RequestAccessWithCompletionHandler]
//   - [AVPlayer.PrerollAtRateCompletionHandler]
//   - [AVPlayer.SeekToDateCompletionHandler]
//   - [AVPlayer.SeekToTimeCompletionHandler]
//   - [AVPlayer.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]
//   - [AVPlayerItem.SeekToDateCompletionHandler]
//   - [AVPlayerItem.SeekToTimeCompletionHandler]
//   - [AVPlayerItem.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]
//   - [AVPlayerItemIntegratedTimeline.AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock]
//   - [AVPlayerItemIntegratedTimeline.SeekToDateCompletionHandler]
//   - [AVPlayerItemIntegratedTimeline.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]
//   - [AVSampleBufferAudioRenderer.FlushFromSourceTimeCompletionHandler]
//   - [AVSampleBufferRenderSynchronizer.RemoveRendererAtTimeCompletionHandler]
type BoolHandler = func(bool)

// NewBoolBlock wraps a Go [BoolHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetExportSession.DetermineCompatibilityOfExportPresetWithAssetOutputFileTypeCompletionHandler]
//   - [AVCaptureDevice.RequestAccessForMediaTypeCompletionHandler]
//   - [AVExternalStorageDevice.RequestAccessWithCompletionHandler]
//   - [AVPlayer.PrerollAtRateCompletionHandler]
//   - [AVPlayer.SeekToDateCompletionHandler]
//   - [AVPlayer.SeekToTimeCompletionHandler]
//   - [AVPlayer.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]
//   - [AVPlayerItem.SeekToDateCompletionHandler]
//   - [AVPlayerItem.SeekToTimeCompletionHandler]
//   - [AVPlayerItem.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]
//   - [AVPlayerItemIntegratedTimeline.AddBoundaryTimeObserverForSegmentOffsetsIntoSegmentQueueUsingBlock]
//   - [AVPlayerItemIntegratedTimeline.SeekToDateCompletionHandler]
//   - [AVPlayerItemIntegratedTimeline.SeekToTimeToleranceBeforeToleranceAfterCompletionHandler]
//   - [AVSampleBufferAudioRenderer.FlushFromSourceTimeCompletionHandler]
//   - [AVSampleBufferRenderSynchronizer.RemoveRendererAtTimeCompletionHandler]
func NewBoolBlock(handler BoolHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// CGImageRefErrorHandler handles A callback that the image generator invokes with the result of the request.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAssetImageGenerator.GenerateCGImageAsynchronouslyForTimeCompletionHandler]
type CGImageRefErrorHandler = func(coregraphics.CGImageRef, error)

// NewCGImageRefErrorBlock wraps a Go [CGImageRefErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetImageGenerator.GenerateCGImageAsynchronouslyForTimeCompletionHandler]
func NewCGImageRefErrorBlock(handler CGImageRefErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal coregraphics.CGImageRef, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(primitiveVal, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// CMPersistentTrackIDErrorHandler handles A completion handler the system calls after it finishes the request.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAsset.FindUnusedTrackIDWithCompletionHandler]
type CMPersistentTrackIDErrorHandler = func(int32, error)

// NewCMPersistentTrackIDErrorBlock wraps a Go [CMPersistentTrackIDErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAsset.FindUnusedTrackIDWithCompletionHandler]
func NewCMPersistentTrackIDErrorBlock(handler CMPersistentTrackIDErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int32, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(primitiveVal, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// CMTimeErrorHandler handles A callback the system invokes when it finishes its estimation.
//   - time: A [CMTime](<doc://com.apple.documentation/documentation/CoreMedia/CMTime>) value, which is [invalid](<doc://com.apple.documentation/documentation/CoreMedia/CMTime/invalid>) if the track time is out of range or if an error occurs.
//   - error: An error object if the request fails; otherwise, `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAssetExportSession.EstimateMaximumDurationWithCompletionHandler]
//   - [AVAssetTrack.LoadSamplePresentationTimeForTrackTimeCompletionHandler]
//   - [AVCaptureDevice.SetDynamicAspectRatioCompletionHandler]
type CMTimeErrorHandler = func(objectivec.IObject, error)

// NewCMTimeErrorBlock wraps a Go [CMTimeErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetExportSession.EstimateMaximumDurationWithCompletionHandler]
//   - [AVAssetTrack.LoadSamplePresentationTimeForTrackTimeCompletionHandler]
//   - [AVCaptureDevice.SetDynamicAspectRatioCompletionHandler]
func NewCMTimeErrorBlock(handler CMTimeErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.IObject, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(primitiveVal, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// CMTimeHandler handles A callback the system invokes when the adjustment to the exposure duration and ISO is complete.
//   - time: The time at which the system invokes the block.
//
// Used by:
//   - [AVCaptureDevice.SetExposureModeCustomWithDurationISOCompletionHandler]
//   - [AVCaptureDevice.SetExposureTargetBiasCompletionHandler]
//   - [AVCaptureDevice.SetFocusModeLockedWithLensPositionCompletionHandler]
//   - [AVCaptureDevice.SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler]
//   - [AVCaptureDevice.SetWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValuesCompletionHandler]
//   - [AVPlayer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]
//   - [AVPlayerItemIntegratedTimeline.AddPeriodicTimeObserverForIntervalQueueUsingBlock]
//   - [AVSampleBufferRenderSynchronizer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]
type CMTimeHandler = func(objectivec.IObject)

// NewCMTimeBlock wraps a Go [CMTimeHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVCaptureDevice.SetExposureModeCustomWithDurationISOCompletionHandler]
//   - [AVCaptureDevice.SetExposureTargetBiasCompletionHandler]
//   - [AVCaptureDevice.SetFocusModeLockedWithLensPositionCompletionHandler]
//   - [AVCaptureDevice.SetWhiteBalanceModeLockedWithDeviceWhiteBalanceGainsCompletionHandler]
//   - [AVCaptureDevice.SetWhiteBalanceModeLockedWithDeviceWhiteBalanceTemperatureAndTintValuesCompletionHandler]
//   - [AVPlayer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]
//   - [AVPlayerItemIntegratedTimeline.AddPeriodicTimeObserverForIntervalQueueUsingBlock]
//   - [AVSampleBufferRenderSynchronizer.AddPeriodicTimeObserverForIntervalQueueUsingBlock]
func NewCMTimeBlock(handler CMTimeHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal objectivec.IObject) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// DataErrorHandler handles A block called after the streaming content key request has been prepared.
//   - contentKeyRequestData: The streaming content key request data.
//   - error: An object that describes the error, if one occurred; otherwise, the value is `nil`.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVContentKeyRequest.MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler]
//   - [AVContentKeySession.InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler]
//   - [AVContentKeySession.InvalidatePersistableContentKeyOptionsCompletionHandler]
//   - [AVContentKeySession.MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler]
type DataErrorHandler = func(*foundation.NSData, error)

// NewDataErrorBlock wraps a Go [DataErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVContentKeyRequest.MakeStreamingContentKeyRequestDataForAppContentIdentifierOptionsCompletionHandler]
//   - [AVContentKeySession.InvalidateAllPersistableContentKeysForAppOptionsCompletionHandler]
//   - [AVContentKeySession.InvalidatePersistableContentKeyOptionsCompletionHandler]
//   - [AVContentKeySession.MakeSecureTokenForExpirationDateOfPersistableContentKeyCompletionHandler]
func NewDataErrorBlock(handler DataErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSData
		if resultID != 0 {
			v := foundation.NSDataFromID(resultID)
			result = &v
		}
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(result, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The code to perform after the system displays Desk View.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAssetImageGenerator.GenerateCGImagesAsynchronouslyForTimesCompletionHandler]
//   - [AVCaptureDeskViewApplication.PresentWithCompletionHandler]
//   - [AVCaptureDeskViewApplication.PresentWithLaunchConfigurationCompletionHandler]
//   - [AVMutableComposition.InsertTimeRangeOfAssetAtTimeCompletionHandler]
//   - [AVSampleBufferGeneratorBatch.MakeDataReadyWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetImageGenerator.GenerateCGImagesAsynchronouslyForTimesCompletionHandler]
//   - [AVCaptureDeskViewApplication.PresentWithCompletionHandler]
//   - [AVCaptureDeskViewApplication.PresentWithLaunchConfigurationCompletionHandler]
//   - [AVMutableComposition.InsertTimeRangeOfAssetAtTimeCompletionHandler]
//   - [AVSampleBufferGeneratorBatch.MakeDataReadyWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

// Float32Handler handles The action to perform in response to changes to the slider’s value.
//
// Used by:
//   - [AVCaptureSlider.SetActionQueueAction]
//   - [AVCaptureSystemExposureBiasSlider.InitWithDeviceAction]
type Float32Handler = func(float32)

// NewFloat32Block wraps a Go [Float32Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVCaptureSlider.SetActionQueueAction]
//   - [AVCaptureSystemExposureBiasSlider.InitWithDeviceAction]
func NewFloat32Block(handler Float32Handler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal float32) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// Float64Handler handles An action the system calls on the main actor to respond to changes to the device’s videoZoomFactor property.
//
// Used by:
//   - [AVCaptureSystemZoomSlider.InitWithDeviceAction]
type Float64Handler = func(float64)

// NewFloat64Block wraps a Go [Float64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVCaptureSystemZoomSlider.InitWithDeviceAction]
func NewFloat64Block(handler Float64Handler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal float64) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// IntHandler handles A transformation from index to localized title.
//
// Used by:
//   - [AVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform]
//   - [AVCaptureIndexPicker.SetActionQueueAction]
type IntHandler = func(int)

// NewIntBlock wraps a Go [IntHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVCaptureIndexPicker.InitWithLocalizedTitleSymbolNameNumberOfIndexesLocalizedTitleTransform]
//   - [AVCaptureIndexPicker.SetActionQueueAction]
func NewIntBlock(handler IntHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles A callback the system passes an array of AVFileType structures when it determines the compatible file types.
//
// Used by:
//   - [AVAssetExportSession.DetermineCompatibleFileTypesWithCompletionHandler]
//   - [AVAssetExportSession.ExportAsynchronouslyWithCompletionHandler]
//   - [AVAssetPlaybackAssistant.LoadPlaybackConfigurationOptionsWithCompletionHandler]
//   - [AVAssetWriter.FinishWritingWithCompletionHandler]
//   - [AVMetadataItem.LoadValuesAsynchronouslyForKeysCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssuePauseCommandCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssuePlayCommandCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssueSeekCommandCompletionHandler]
//   - [AVPlayer.AddBoundaryTimeObserverForTimesQueueUsingBlock]
//   - [AVPlayerItem.RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler]
//   - [AVQueuedSampleBufferRendering.RequestMediaDataWhenReadyOnQueueUsingBlock]
//   - [AVSampleBufferAudioRenderer.RequestMediaDataWhenReadyOnQueueUsingBlock]
//   - [AVSampleBufferDisplayLayer.RequestMediaDataWhenReadyOnQueueUsingBlock]
//   - [AVSampleBufferRenderSynchronizer.AddBoundaryTimeObserverForTimesQueueUsingBlock]
//   - [AVSampleBufferVideoRenderer.FlushWithRemovalOfDisplayedImageCompletionHandler]
//   - [AVSampleBufferVideoRenderer.RequestMediaDataWhenReadyOnQueueUsingBlock]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetExportSession.DetermineCompatibleFileTypesWithCompletionHandler]
//   - [AVAssetExportSession.ExportAsynchronouslyWithCompletionHandler]
//   - [AVAssetPlaybackAssistant.LoadPlaybackConfigurationOptionsWithCompletionHandler]
//   - [AVAssetWriter.FinishWritingWithCompletionHandler]
//   - [AVMetadataItem.LoadValuesAsynchronouslyForKeysCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssueBufferingCommandCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssuePauseCommandCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssuePlayCommandCompletionHandler]
//   - [AVPlaybackCoordinatorPlaybackControlDelegate.PlaybackCoordinatorDidIssueSeekCommandCompletionHandler]
//   - [AVPlayer.AddBoundaryTimeObserverForTimesQueueUsingBlock]
//   - [AVPlayerItem.RequestContentAuthorizationAsynchronouslyWithTimeoutIntervalCompletionHandler]
//   - [AVQueuedSampleBufferRendering.RequestMediaDataWhenReadyOnQueueUsingBlock]
//   - [AVSampleBufferAudioRenderer.RequestMediaDataWhenReadyOnQueueUsingBlock]
//   - [AVSampleBufferDisplayLayer.RequestMediaDataWhenReadyOnQueueUsingBlock]
//   - [AVSampleBufferRenderSynchronizer.AddBoundaryTimeObserverForTimesQueueUsingBlock]
//   - [AVSampleBufferVideoRenderer.FlushWithRemovalOfDisplayedImageCompletionHandler]
//   - [AVSampleBufferVideoRenderer.RequestMediaDataWhenReadyOnQueueUsingBlock]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// int64_tErrorHandler handles A callback the system invokes when it finishes its estimation.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [AVAssetExportSession.EstimateOutputFileLengthWithCompletionHandler]
type int64_tErrorHandler = func(int64, error)

// Newint64_tErrorBlock wraps a Go [int64_tErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [AVAssetExportSession.EstimateOutputFileLengthWithCompletionHandler]
func Newint64_tErrorBlock(handler int64_tErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal int64, errID objc.ID) {
		var nserr *foundation.NSError
		if errID != 0 {
			e := foundation.NSErrorFromID(errID)
			nserr = &e
		}
		handler(primitiveVal, foundation.NSErrorToError(nserr))
	})
	return objc.ID(block), func() { block.Release() }
}

