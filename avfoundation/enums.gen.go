// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetExportSession/Status-swift.enum
type AVAssetExportSessionStatus int

const (
	// AVAssetExportSessionStatusCancelled: You canceled the export.
	AVAssetExportSessionStatusCancelled AVAssetExportSessionStatus = 5
	// AVAssetExportSessionStatusCompleted: The export completes successfully.
	AVAssetExportSessionStatusCompleted AVAssetExportSessionStatus = 3
	// AVAssetExportSessionStatusExporting: The export is in progress.
	AVAssetExportSessionStatusExporting AVAssetExportSessionStatus = 2
	// AVAssetExportSessionStatusFailed: The export fails.
	AVAssetExportSessionStatusFailed AVAssetExportSessionStatus = 4
	// AVAssetExportSessionStatusUnknown: The session status is unknown.
	AVAssetExportSessionStatusUnknown AVAssetExportSessionStatus = 0
	// AVAssetExportSessionStatusWaiting: The session is waiting to export more data.
	AVAssetExportSessionStatusWaiting AVAssetExportSessionStatus = 1
)

func (e AVAssetExportSessionStatus) String() string {
	switch e {
	case AVAssetExportSessionStatusCancelled:
		return "AVAssetExportSessionStatusCancelled"
	case AVAssetExportSessionStatusCompleted:
		return "AVAssetExportSessionStatusCompleted"
	case AVAssetExportSessionStatusExporting:
		return "AVAssetExportSessionStatusExporting"
	case AVAssetExportSessionStatusFailed:
		return "AVAssetExportSessionStatusFailed"
	case AVAssetExportSessionStatusUnknown:
		return "AVAssetExportSessionStatusUnknown"
	case AVAssetExportSessionStatusWaiting:
		return "AVAssetExportSessionStatusWaiting"
	default:
		return fmt.Sprintf("AVAssetExportSessionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetImageGenerator/Result
type AVAssetImageGeneratorResult int

const (
	// AVAssetImageGeneratorCancelled: A result that indicates you canceled image generation.
	AVAssetImageGeneratorCancelled AVAssetImageGeneratorResult = 2
	// AVAssetImageGeneratorFailed: A result that indicates that image generation failed.
	AVAssetImageGeneratorFailed AVAssetImageGeneratorResult = 1
	// AVAssetImageGeneratorSucceeded: A result that indicates that image generation succeeded.
	AVAssetImageGeneratorSucceeded AVAssetImageGeneratorResult = 0
)

func (e AVAssetImageGeneratorResult) String() string {
	switch e {
	case AVAssetImageGeneratorCancelled:
		return "AVAssetImageGeneratorCancelled"
	case AVAssetImageGeneratorFailed:
		return "AVAssetImageGeneratorFailed"
	case AVAssetImageGeneratorSucceeded:
		return "AVAssetImageGeneratorSucceeded"
	default:
		return fmt.Sprintf("AVAssetImageGeneratorResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReader/Status-swift.enum
type AVAssetReaderStatus int

const (
	// AVAssetReaderStatusCancelled: The asset reader can no longer read samples because you canceled reading.
	AVAssetReaderStatusCancelled AVAssetReaderStatus = 4
	// AVAssetReaderStatusCompleted: The asset reader completes reading all samples within its specified time range.
	AVAssetReaderStatusCompleted AVAssetReaderStatus = 2
	// AVAssetReaderStatusFailed: The asset reader can no longer read samples from its asset because of an error.
	AVAssetReaderStatusFailed AVAssetReaderStatus = 3
	// AVAssetReaderStatusReading: The asset reader is successfully reading samples from its asset.
	AVAssetReaderStatusReading AVAssetReaderStatus = 1
	// AVAssetReaderStatusUnknown: The asset reader is in an unknown state.
	AVAssetReaderStatusUnknown AVAssetReaderStatus = 0
)

func (e AVAssetReaderStatus) String() string {
	switch e {
	case AVAssetReaderStatusCancelled:
		return "AVAssetReaderStatusCancelled"
	case AVAssetReaderStatusCompleted:
		return "AVAssetReaderStatusCompleted"
	case AVAssetReaderStatusFailed:
		return "AVAssetReaderStatusFailed"
	case AVAssetReaderStatusReading:
		return "AVAssetReaderStatusReading"
	case AVAssetReaderStatusUnknown:
		return "AVAssetReaderStatusUnknown"
	default:
		return fmt.Sprintf("AVAssetReaderStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetReferenceRestrictions
type AVAssetReferenceRestrictions uint

const (
	// AVAssetReferenceRestrictionDefaultPolicy: The asset should use the default reference restrictions policy.
	AVAssetReferenceRestrictionDefaultPolicy AVAssetReferenceRestrictions = 0
	// AVAssetReferenceRestrictionForbidAll: The asset can only reference media stored within its container file.
	AVAssetReferenceRestrictionForbidAll AVAssetReferenceRestrictions = 65535
	// AVAssetReferenceRestrictionForbidCrossSiteReference: A remote asset shouldn’t follow references to remote media data stored at a different host.
	AVAssetReferenceRestrictionForbidCrossSiteReference AVAssetReferenceRestrictions = 4
	// AVAssetReferenceRestrictionForbidLocalReferenceToLocal: A local asset shouldn’t follow references to local media data stored outside its container file.
	AVAssetReferenceRestrictionForbidLocalReferenceToLocal AVAssetReferenceRestrictions = 8
	// AVAssetReferenceRestrictionForbidLocalReferenceToRemote: A local asset shouldn’t follow references to remote media.
	AVAssetReferenceRestrictionForbidLocalReferenceToRemote AVAssetReferenceRestrictions = 2
	// AVAssetReferenceRestrictionForbidNone: The asset should follow all media references.
	AVAssetReferenceRestrictionForbidNone AVAssetReferenceRestrictions = 0
	// AVAssetReferenceRestrictionForbidRemoteReferenceToLocal: A remote asset shouldn’t follow references to local media.
	AVAssetReferenceRestrictionForbidRemoteReferenceToLocal AVAssetReferenceRestrictions = 1
)

func (e AVAssetReferenceRestrictions) String() string {
	switch e {
	case AVAssetReferenceRestrictionDefaultPolicy:
		return "AVAssetReferenceRestrictionDefaultPolicy"
	case AVAssetReferenceRestrictionForbidAll:
		return "AVAssetReferenceRestrictionForbidAll"
	case AVAssetReferenceRestrictionForbidCrossSiteReference:
		return "AVAssetReferenceRestrictionForbidCrossSiteReference"
	case AVAssetReferenceRestrictionForbidLocalReferenceToLocal:
		return "AVAssetReferenceRestrictionForbidLocalReferenceToLocal"
	case AVAssetReferenceRestrictionForbidLocalReferenceToRemote:
		return "AVAssetReferenceRestrictionForbidLocalReferenceToRemote"
	case AVAssetReferenceRestrictionForbidRemoteReferenceToLocal:
		return "AVAssetReferenceRestrictionForbidRemoteReferenceToLocal"
	default:
		return fmt.Sprintf("AVAssetReferenceRestrictions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetSegmentType
type AVAssetSegmentType int

const (
	// AVAssetSegmentTypeInitialization: An initialization segment type.
	AVAssetSegmentTypeInitialization AVAssetSegmentType = 1
	// AVAssetSegmentTypeSeparable: A separable segment type.
	AVAssetSegmentTypeSeparable AVAssetSegmentType = 2
)

func (e AVAssetSegmentType) String() string {
	switch e {
	case AVAssetSegmentTypeInitialization:
		return "AVAssetSegmentTypeInitialization"
	case AVAssetSegmentTypeSeparable:
		return "AVAssetSegmentTypeSeparable"
	default:
		return fmt.Sprintf("AVAssetSegmentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroupOutputHandling
type AVAssetTrackGroupOutputHandling uint

const (
	// AVAssetTrackGroupOutputHandlingDefaultPolicy: The default track group output handling policy.
	AVAssetTrackGroupOutputHandlingDefaultPolicy AVAssetTrackGroupOutputHandling = 0
	// AVAssetTrackGroupOutputHandlingNone: A policy that doesn’t pass through alternate audio tracks from the source asset during export.
	AVAssetTrackGroupOutputHandlingNone AVAssetTrackGroupOutputHandling = 0
	// AVAssetTrackGroupOutputHandlingPreserveAlternateTracks: A policy that passes through alternate audio tracks from the source asset during export.
	AVAssetTrackGroupOutputHandlingPreserveAlternateTracks AVAssetTrackGroupOutputHandling = 1
)

func (e AVAssetTrackGroupOutputHandling) String() string {
	switch e {
	case AVAssetTrackGroupOutputHandlingDefaultPolicy:
		return "AVAssetTrackGroupOutputHandlingDefaultPolicy"
	case AVAssetTrackGroupOutputHandlingPreserveAlternateTracks:
		return "AVAssetTrackGroupOutputHandlingPreserveAlternateTracks"
	default:
		return fmt.Sprintf("AVAssetTrackGroupOutputHandling(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAssetWriter/Status-swift.enum
type AVAssetWriterStatus int

const (
	// AVAssetWriterStatusCancelled: The asset writer canceled the writing operation.
	AVAssetWriterStatusCancelled AVAssetWriterStatus = 4
	// AVAssetWriterStatusCompleted: The asset writer finishes writing successfully.
	AVAssetWriterStatusCompleted AVAssetWriterStatus = 2
	// AVAssetWriterStatusFailed: The asset writer fails to write the output file.
	AVAssetWriterStatusFailed AVAssetWriterStatus = 3
	// AVAssetWriterStatusUnknown: The asset writer’s status isn’t known.
	AVAssetWriterStatusUnknown AVAssetWriterStatus = 0
	// AVAssetWriterStatusWriting: The asset writer is writing.
	AVAssetWriterStatusWriting AVAssetWriterStatus = 1
)

func (e AVAssetWriterStatus) String() string {
	switch e {
	case AVAssetWriterStatusCancelled:
		return "AVAssetWriterStatusCancelled"
	case AVAssetWriterStatusCompleted:
		return "AVAssetWriterStatusCompleted"
	case AVAssetWriterStatusFailed:
		return "AVAssetWriterStatusFailed"
	case AVAssetWriterStatusUnknown:
		return "AVAssetWriterStatusUnknown"
	case AVAssetWriterStatusWriting:
		return "AVAssetWriterStatusWriting"
	default:
		return fmt.Sprintf("AVAssetWriterStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAudioSpatializationFormats
type AVAudioSpatializationFormats uint

const (
	// AVAudioSpatializationFormatMonoAndStereo: A value that indicates the player item only supports mono and stereo layouts for audio spatialization.
	AVAudioSpatializationFormatMonoAndStereo AVAudioSpatializationFormats = 3
	// AVAudioSpatializationFormatMonoStereoAndMultichannel: A value that indicates the player item supports mono, stereo, and multichannel layouts for audio spatialization.
	AVAudioSpatializationFormatMonoStereoAndMultichannel AVAudioSpatializationFormats = 7
	// AVAudioSpatializationFormatMultichannel: A value that indicates the player item only supports multichannel layouts for audio spatialization.
	AVAudioSpatializationFormatMultichannel AVAudioSpatializationFormats = 4
	// AVAudioSpatializationFormatNone: A value that indicates the player item doesn’t support audio spatialization.
	AVAudioSpatializationFormatNone AVAudioSpatializationFormats = 0
)

func (e AVAudioSpatializationFormats) String() string {
	switch e {
	case AVAudioSpatializationFormatMonoAndStereo:
		return "AVAudioSpatializationFormatMonoAndStereo"
	case AVAudioSpatializationFormatMonoStereoAndMultichannel:
		return "AVAudioSpatializationFormatMonoStereoAndMultichannel"
	case AVAudioSpatializationFormatMultichannel:
		return "AVAudioSpatializationFormatMultichannel"
	case AVAudioSpatializationFormatNone:
		return "AVAudioSpatializationFormatNone"
	default:
		return fmt.Sprintf("AVAudioSpatializationFormats(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVAuthorizationStatus
type AVAuthorizationStatus int

const (
	// AVAuthorizationStatusAuthorized: A status that indicates the user has explicitly granted an app permission to capture media.
	AVAuthorizationStatusAuthorized AVAuthorizationStatus = 3
	// AVAuthorizationStatusDenied: A status that indicates the user has explicitly denied an app permission to capture media.
	AVAuthorizationStatusDenied AVAuthorizationStatus = 2
	// AVAuthorizationStatusNotDetermined: A status that indicates the user hasn’t yet granted or denied authorization.
	AVAuthorizationStatusNotDetermined AVAuthorizationStatus = 0
	// AVAuthorizationStatusRestricted: A status that indicates the app isn’t permitted to use media capture devices.
	AVAuthorizationStatusRestricted AVAuthorizationStatus = 1
)

func (e AVAuthorizationStatus) String() string {
	switch e {
	case AVAuthorizationStatusAuthorized:
		return "AVAuthorizationStatusAuthorized"
	case AVAuthorizationStatusDenied:
		return "AVAuthorizationStatusDenied"
	case AVAuthorizationStatusNotDetermined:
		return "AVAuthorizationStatusNotDetermined"
	case AVAuthorizationStatusRestricted:
		return "AVAuthorizationStatusRestricted"
	default:
		return fmt.Sprintf("AVAuthorizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Animation-swift.enum
type AVCaptionAnimation int

const (
	// AVCaptionAnimationCharacterReveal: A character reveal animation.
	AVCaptionAnimationCharacterReveal AVCaptionAnimation = 1
	// AVCaptionAnimationNone: No animation.
	AVCaptionAnimationNone AVCaptionAnimation = 0
)

func (e AVCaptionAnimation) String() string {
	switch e {
	case AVCaptionAnimationCharacterReveal:
		return "AVCaptionAnimationCharacterReveal"
	case AVCaptionAnimationNone:
		return "AVCaptionAnimationNone"
	default:
		return fmt.Sprintf("AVCaptionAnimation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionConversionValidator/Status-swift.enum
type AVCaptionConversionValidatorStatus int

const (
	// AVCaptionConversionValidatorStatusCompleted: A status that indicates the system validation is complete.
	AVCaptionConversionValidatorStatusCompleted AVCaptionConversionValidatorStatus = 2
	// AVCaptionConversionValidatorStatusStopped: A status that indicates the system validation stopped prior to completion.
	AVCaptionConversionValidatorStatusStopped AVCaptionConversionValidatorStatus = 3
	// AVCaptionConversionValidatorStatusUnknown: A status that indicates the system didn’t initialize the validation operation.
	AVCaptionConversionValidatorStatusUnknown AVCaptionConversionValidatorStatus = 0
	// AVCaptionConversionValidatorStatusValidating: A status that indicates the system validation is in progress.
	AVCaptionConversionValidatorStatusValidating AVCaptionConversionValidatorStatus = 1
)

func (e AVCaptionConversionValidatorStatus) String() string {
	switch e {
	case AVCaptionConversionValidatorStatusCompleted:
		return "AVCaptionConversionValidatorStatusCompleted"
	case AVCaptionConversionValidatorStatusStopped:
		return "AVCaptionConversionValidatorStatusStopped"
	case AVCaptionConversionValidatorStatusUnknown:
		return "AVCaptionConversionValidatorStatusUnknown"
	case AVCaptionConversionValidatorStatusValidating:
		return "AVCaptionConversionValidatorStatusValidating"
	default:
		return fmt.Sprintf("AVCaptionConversionValidatorStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/Decoration
type AVCaptionDecoration uint

const (
	// AVCaptionDecorationLineThrough: A decoration representing a line through the text.
	AVCaptionDecorationLineThrough AVCaptionDecoration = 2
	// AVCaptionDecorationNone: No text decoration.
	AVCaptionDecorationNone AVCaptionDecoration = 0
	// AVCaptionDecorationOverline: A decoration representing a line over the text.
	AVCaptionDecorationOverline AVCaptionDecoration = 4
	// AVCaptionDecorationUnderline: A decoration representing a line under the text.
	AVCaptionDecorationUnderline AVCaptionDecoration = 1
)

func (e AVCaptionDecoration) String() string {
	switch e {
	case AVCaptionDecorationLineThrough:
		return "AVCaptionDecorationLineThrough"
	case AVCaptionDecorationNone:
		return "AVCaptionDecorationNone"
	case AVCaptionDecorationOverline:
		return "AVCaptionDecorationOverline"
	case AVCaptionDecorationUnderline:
		return "AVCaptionDecorationUnderline"
	default:
		return fmt.Sprintf("AVCaptionDecoration(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontStyle
type AVCaptionFontStyle int

const (
	// AVCaptionFontStyleItalic: An italic font style.
	AVCaptionFontStyleItalic AVCaptionFontStyle = 2
	// AVCaptionFontStyleNormal: A normal font style.
	AVCaptionFontStyleNormal AVCaptionFontStyle = 1
	// AVCaptionFontStyleUnknown: An unknown font style.
	AVCaptionFontStyleUnknown AVCaptionFontStyle = 0
)

func (e AVCaptionFontStyle) String() string {
	switch e {
	case AVCaptionFontStyleItalic:
		return "AVCaptionFontStyleItalic"
	case AVCaptionFontStyleNormal:
		return "AVCaptionFontStyleNormal"
	case AVCaptionFontStyleUnknown:
		return "AVCaptionFontStyleUnknown"
	default:
		return fmt.Sprintf("AVCaptionFontStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/FontWeight
type AVCaptionFontWeight int

const (
	// AVCaptionFontWeightBold: A bold font weight.
	AVCaptionFontWeightBold AVCaptionFontWeight = 2
	// AVCaptionFontWeightNormal: A normal font weight.
	AVCaptionFontWeightNormal AVCaptionFontWeight = 1
	// AVCaptionFontWeightUnknown: An unknown font weight.
	AVCaptionFontWeightUnknown AVCaptionFontWeight = 0
)

func (e AVCaptionFontWeight) String() string {
	switch e {
	case AVCaptionFontWeightBold:
		return "AVCaptionFontWeightBold"
	case AVCaptionFontWeightNormal:
		return "AVCaptionFontWeightNormal"
	case AVCaptionFontWeightUnknown:
		return "AVCaptionFontWeightUnknown"
	default:
		return fmt.Sprintf("AVCaptionFontWeight(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/DisplayAlignment-swift.enum
type AVCaptionRegionDisplayAlignment int

const (
	// AVCaptionRegionDisplayAlignmentAfter: An alignment that positions lines at the bottom of the block progression direction.
	AVCaptionRegionDisplayAlignmentAfter AVCaptionRegionDisplayAlignment = 2
	// AVCaptionRegionDisplayAlignmentBefore: An alignment that positions lines at the top of the block progression direction.
	AVCaptionRegionDisplayAlignmentBefore AVCaptionRegionDisplayAlignment = 0
	// AVCaptionRegionDisplayAlignmentCenter: An alignment that positions lines in the middle of the block progression direction.
	AVCaptionRegionDisplayAlignmentCenter AVCaptionRegionDisplayAlignment = 1
)

func (e AVCaptionRegionDisplayAlignment) String() string {
	switch e {
	case AVCaptionRegionDisplayAlignmentAfter:
		return "AVCaptionRegionDisplayAlignmentAfter"
	case AVCaptionRegionDisplayAlignmentBefore:
		return "AVCaptionRegionDisplayAlignmentBefore"
	case AVCaptionRegionDisplayAlignmentCenter:
		return "AVCaptionRegionDisplayAlignmentCenter"
	default:
		return fmt.Sprintf("AVCaptionRegionDisplayAlignment(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/Scroll-swift.enum
type AVCaptionRegionScroll int

const (
	// AVCaptionRegionScrollNone: A type that indicates no scrolling.
	AVCaptionRegionScrollNone AVCaptionRegionScroll = 0
	// AVCaptionRegionScrollRollUp: A type that indicates a roll-up scroll effect.
	AVCaptionRegionScrollRollUp AVCaptionRegionScroll = 1
)

func (e AVCaptionRegionScroll) String() string {
	switch e {
	case AVCaptionRegionScrollNone:
		return "AVCaptionRegionScrollNone"
	case AVCaptionRegionScrollRollUp:
		return "AVCaptionRegionScrollRollUp"
	default:
		return fmt.Sprintf("AVCaptionRegionScroll(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRegion/WritingMode-swift.enum
type AVCaptionRegionWritingMode int

const (
	// AVCaptionRegionWritingModeLeftToRightAndTopToBottom: A left-to-right and top-to-bottom writing mode.
	AVCaptionRegionWritingModeLeftToRightAndTopToBottom AVCaptionRegionWritingMode = 0
	// AVCaptionRegionWritingModeTopToBottomAndRightToLeft: A top-to-bottom and right-to-left writing mode.
	AVCaptionRegionWritingModeTopToBottomAndRightToLeft AVCaptionRegionWritingMode = 2
)

func (e AVCaptionRegionWritingMode) String() string {
	switch e {
	case AVCaptionRegionWritingModeLeftToRightAndTopToBottom:
		return "AVCaptionRegionWritingModeLeftToRightAndTopToBottom"
	case AVCaptionRegionWritingModeTopToBottomAndRightToLeft:
		return "AVCaptionRegionWritingModeTopToBottomAndRightToLeft"
	default:
		return fmt.Sprintf("AVCaptionRegionWritingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyAlignment
type AVCaptionRubyAlignment int

const (
	// AVCaptionRubyAlignmentCenter: An alignment with the ruby text at the center of ruby base.
	AVCaptionRubyAlignmentCenter AVCaptionRubyAlignment = 1
	// AVCaptionRubyAlignmentDistributeSpaceAround: An alignment with the ruby text so the spaces around each ruby text character are equal.
	AVCaptionRubyAlignmentDistributeSpaceAround AVCaptionRubyAlignment = 3
	// AVCaptionRubyAlignmentDistributeSpaceBetween: An alignment with the ruby text so the spaces between the ruby text characters are equal.
	AVCaptionRubyAlignmentDistributeSpaceBetween AVCaptionRubyAlignment = 2
	// AVCaptionRubyAlignmentStart: An alignment with the ruby base and text at the left edge of horizontal text in a left-to-right inline progression, or at top of the vertical text in a top-to-bottom inline progression.
	AVCaptionRubyAlignmentStart AVCaptionRubyAlignment = 0
)

func (e AVCaptionRubyAlignment) String() string {
	switch e {
	case AVCaptionRubyAlignmentCenter:
		return "AVCaptionRubyAlignmentCenter"
	case AVCaptionRubyAlignmentDistributeSpaceAround:
		return "AVCaptionRubyAlignmentDistributeSpaceAround"
	case AVCaptionRubyAlignmentDistributeSpaceBetween:
		return "AVCaptionRubyAlignmentDistributeSpaceBetween"
	case AVCaptionRubyAlignmentStart:
		return "AVCaptionRubyAlignmentStart"
	default:
		return fmt.Sprintf("AVCaptionRubyAlignment(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRubyPosition
type AVCaptionRubyPosition int

const (
	// AVCaptionRubyPositionAfter: Display ruby text below horizontal text, or to the left of vertical text in a right-to-left block progression.
	AVCaptionRubyPositionAfter AVCaptionRubyPosition = 1
	// AVCaptionRubyPositionBefore: Display ruby text above horizontal text, or to the right of vertical text in a right-to-left block progression.
	AVCaptionRubyPositionBefore AVCaptionRubyPosition = 0
)

func (e AVCaptionRubyPosition) String() string {
	switch e {
	case AVCaptionRubyPositionAfter:
		return "AVCaptionRubyPositionAfter"
	case AVCaptionRubyPositionBefore:
		return "AVCaptionRubyPositionBefore"
	default:
		return fmt.Sprintf("AVCaptionRubyPosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextAlignment-swift.enum
type AVCaptionTextAlignment int

const (
	// AVCaptionTextAlignmentCenter: An alignment of the text to the center of the display.
	AVCaptionTextAlignmentCenter AVCaptionTextAlignment = 2
	// AVCaptionTextAlignmentEnd: An alignment of the text to the end of the inline progression direction.
	AVCaptionTextAlignmentEnd AVCaptionTextAlignment = 1
	// AVCaptionTextAlignmentLeft: An alignment of the text to the left in horizontal writing mode, and top in vertical writing mode.
	AVCaptionTextAlignmentLeft AVCaptionTextAlignment = 3
	// AVCaptionTextAlignmentRight: An alignment of the text to the right in horizontal writing mode, and bottom in vertical writing mode.
	AVCaptionTextAlignmentRight AVCaptionTextAlignment = 4
	// AVCaptionTextAlignmentStart: An alignment of the text to the start of the inline progression direction.
	AVCaptionTextAlignmentStart AVCaptionTextAlignment = 0
)

func (e AVCaptionTextAlignment) String() string {
	switch e {
	case AVCaptionTextAlignmentCenter:
		return "AVCaptionTextAlignmentCenter"
	case AVCaptionTextAlignmentEnd:
		return "AVCaptionTextAlignmentEnd"
	case AVCaptionTextAlignmentLeft:
		return "AVCaptionTextAlignmentLeft"
	case AVCaptionTextAlignmentRight:
		return "AVCaptionTextAlignmentRight"
	case AVCaptionTextAlignmentStart:
		return "AVCaptionTextAlignmentStart"
	default:
		return fmt.Sprintf("AVCaptionTextAlignment(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaption/TextCombine
type AVCaptionTextCombine int

const (
	// AVCaptionTextCombineAll: An option that combines all of the characters.
	AVCaptionTextCombineAll AVCaptionTextCombine = -1
	// AVCaptionTextCombineFourDigits: An option that combines four consecutive digits.
	AVCaptionTextCombineFourDigits AVCaptionTextCombine = 4
	// AVCaptionTextCombineNone: An option that doesn’t combine text upright.
	AVCaptionTextCombineNone AVCaptionTextCombine = 0
	// AVCaptionTextCombineOneDigit: An option that makes one digit upright.
	AVCaptionTextCombineOneDigit AVCaptionTextCombine = 1
	// AVCaptionTextCombineThreeDigits: An option that combines three consecutive digits.
	AVCaptionTextCombineThreeDigits AVCaptionTextCombine = 3
	// AVCaptionTextCombineTwoDigits: An option that combines two consecutive digits.
	AVCaptionTextCombineTwoDigits AVCaptionTextCombine = 2
)

func (e AVCaptionTextCombine) String() string {
	switch e {
	case AVCaptionTextCombineAll:
		return "AVCaptionTextCombineAll"
	case AVCaptionTextCombineFourDigits:
		return "AVCaptionTextCombineFourDigits"
	case AVCaptionTextCombineNone:
		return "AVCaptionTextCombineNone"
	case AVCaptionTextCombineOneDigit:
		return "AVCaptionTextCombineOneDigit"
	case AVCaptionTextCombineThreeDigits:
		return "AVCaptionTextCombineThreeDigits"
	case AVCaptionTextCombineTwoDigits:
		return "AVCaptionTextCombineTwoDigits"
	default:
		return fmt.Sprintf("AVCaptionTextCombine(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionUnitsType
type AVCaptionUnitsType int

const (
	// AVCaptionUnitsTypeCells: A cell-based unit type.
	AVCaptionUnitsTypeCells AVCaptionUnitsType = 1
	// AVCaptionUnitsTypePercent: A percentage-based unit type.
	AVCaptionUnitsTypePercent AVCaptionUnitsType = 2
	// AVCaptionUnitsTypeUnspecified: An unspecified unit type.
	AVCaptionUnitsTypeUnspecified AVCaptionUnitsType = 0
)

func (e AVCaptionUnitsType) String() string {
	switch e {
	case AVCaptionUnitsTypeCells:
		return "AVCaptionUnitsTypeCells"
	case AVCaptionUnitsTypePercent:
		return "AVCaptionUnitsTypePercent"
	case AVCaptionUnitsTypeUnspecified:
		return "AVCaptionUnitsTypeUnspecified"
	default:
		return fmt.Sprintf("AVCaptionUnitsType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/AutoFocusRangeRestriction-swift.enum
type AVCaptureAutoFocusRangeRestriction int

const (
	// AVCaptureAutoFocusRangeRestrictionFar: The device primarily attempts to focus on subjects far away from the camera.
	AVCaptureAutoFocusRangeRestrictionFar AVCaptureAutoFocusRangeRestriction = 2
	// AVCaptureAutoFocusRangeRestrictionNear: The device primarily attempts to focus on subjects near the camera.
	AVCaptureAutoFocusRangeRestrictionNear AVCaptureAutoFocusRangeRestriction = 1
	// AVCaptureAutoFocusRangeRestrictionNone: The device attempts to focus on objects at any range.
	AVCaptureAutoFocusRangeRestrictionNone AVCaptureAutoFocusRangeRestriction = 0
)

func (e AVCaptureAutoFocusRangeRestriction) String() string {
	switch e {
	case AVCaptureAutoFocusRangeRestrictionFar:
		return "AVCaptureAutoFocusRangeRestrictionFar"
	case AVCaptureAutoFocusRangeRestrictionNear:
		return "AVCaptureAutoFocusRangeRestrictionNear"
	case AVCaptureAutoFocusRangeRestrictionNone:
		return "AVCaptureAutoFocusRangeRestrictionNone"
	default:
		return fmt.Sprintf("AVCaptureAutoFocusRangeRestriction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Format/AutoFocusSystem-swift.enum
type AVCaptureAutoFocusSystem int

const (
	// AVCaptureAutoFocusSystemContrastDetection: A slower autofocus system based on differences in contrast.
	AVCaptureAutoFocusSystemContrastDetection AVCaptureAutoFocusSystem = 1
	// AVCaptureAutoFocusSystemNone: Autofocus isn’t available.
	AVCaptureAutoFocusSystemNone AVCaptureAutoFocusSystem = 0
	// AVCaptureAutoFocusSystemPhaseDetection: A faster autofoscus system based on differences in light phase.
	AVCaptureAutoFocusSystemPhaseDetection AVCaptureAutoFocusSystem = 2
)

func (e AVCaptureAutoFocusSystem) String() string {
	switch e {
	case AVCaptureAutoFocusSystemContrastDetection:
		return "AVCaptureAutoFocusSystemContrastDetection"
	case AVCaptureAutoFocusSystemNone:
		return "AVCaptureAutoFocusSystemNone"
	case AVCaptureAutoFocusSystemPhaseDetection:
		return "AVCaptureAutoFocusSystemPhaseDetection"
	default:
		return fmt.Sprintf("AVCaptureAutoFocusSystem(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus
type AVCaptureCameraLensSmudgeDetectionStatus int

const (
	// AVCaptureCameraLensSmudgeDetectionStatusDisabled: Indicates that the detection is not enabled.
	AVCaptureCameraLensSmudgeDetectionStatusDisabled AVCaptureCameraLensSmudgeDetectionStatus = 0
	// AVCaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected: Indicates that the most recent detection found no smudge on the camera lens.
	AVCaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected AVCaptureCameraLensSmudgeDetectionStatus = 1
	// AVCaptureCameraLensSmudgeDetectionStatusSmudged: Indicates that the most recent detection found the camera lens to be smudged.
	AVCaptureCameraLensSmudgeDetectionStatusSmudged AVCaptureCameraLensSmudgeDetectionStatus = 2
	// AVCaptureCameraLensSmudgeDetectionStatusUnknown: Indicates that the detection result has not settled, commonly caused by excessive camera movement or the content of the scene.
	AVCaptureCameraLensSmudgeDetectionStatusUnknown AVCaptureCameraLensSmudgeDetectionStatus = 3
)

func (e AVCaptureCameraLensSmudgeDetectionStatus) String() string {
	switch e {
	case AVCaptureCameraLensSmudgeDetectionStatusDisabled:
		return "AVCaptureCameraLensSmudgeDetectionStatusDisabled"
	case AVCaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected:
		return "AVCaptureCameraLensSmudgeDetectionStatusSmudgeNotDetected"
	case AVCaptureCameraLensSmudgeDetectionStatusSmudged:
		return "AVCaptureCameraLensSmudgeDetectionStatusSmudged"
	case AVCaptureCameraLensSmudgeDetectionStatusUnknown:
		return "AVCaptureCameraLensSmudgeDetectionStatusUnknown"
	default:
		return fmt.Sprintf("AVCaptureCameraLensSmudgeDetectionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CenterStageControlMode-swift.enum
type AVCaptureCenterStageControlMode int

const (
	// AVCaptureCenterStageControlModeApp: The app controls Center Stage.
	AVCaptureCenterStageControlModeApp AVCaptureCenterStageControlMode = 1
	// AVCaptureCenterStageControlModeCooperative: A user and app cooperatively share control of Center Stage.
	AVCaptureCenterStageControlModeCooperative AVCaptureCenterStageControlMode = 2
	// AVCaptureCenterStageControlModeUser: The user controls Center Stage.
	AVCaptureCenterStageControlModeUser AVCaptureCenterStageControlMode = 0
)

func (e AVCaptureCenterStageControlMode) String() string {
	switch e {
	case AVCaptureCenterStageControlModeApp:
		return "AVCaptureCenterStageControlModeApp"
	case AVCaptureCenterStageControlModeCooperative:
		return "AVCaptureCenterStageControlModeCooperative"
	case AVCaptureCenterStageControlModeUser:
		return "AVCaptureCenterStageControlModeUser"
	default:
		return fmt.Sprintf("AVCaptureCenterStageControlMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode
type AVCaptureCinematicVideoFocusMode int

const (
	// AVCaptureCinematicVideoFocusModeNone: Indicates that no focus mode is specified, in which case weak focus is used as default.
	AVCaptureCinematicVideoFocusModeNone AVCaptureCinematicVideoFocusMode = 0
	// AVCaptureCinematicVideoFocusModeStrong: Indicates that the subject should remain in focus until it exits the scene.
	AVCaptureCinematicVideoFocusModeStrong AVCaptureCinematicVideoFocusMode = 1
	// AVCaptureCinematicVideoFocusModeWeak: Indicates that the Cinematic Video algorithm should automatically adjust focus according to the prominence of the subjects in the scene.
	AVCaptureCinematicVideoFocusModeWeak AVCaptureCinematicVideoFocusMode = 2
)

func (e AVCaptureCinematicVideoFocusMode) String() string {
	switch e {
	case AVCaptureCinematicVideoFocusModeNone:
		return "AVCaptureCinematicVideoFocusModeNone"
	case AVCaptureCinematicVideoFocusModeStrong:
		return "AVCaptureCinematicVideoFocusModeStrong"
	case AVCaptureCinematicVideoFocusModeWeak:
		return "AVCaptureCinematicVideoFocusModeWeak"
	default:
		return fmt.Sprintf("AVCaptureCinematicVideoFocusMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureColorSpace
type AVCaptureColorSpace int

const (
	// AVCaptureColorSpace_AppleLog: The Apple Log Color space, which uses BT2020 as the color primaries, and an Apple-defined Log curve as a transfer function.
	AVCaptureColorSpace_AppleLog AVCaptureColorSpace = 3
	// AVCaptureColorSpace_AppleLog2: The Apple Log 2 Color space, which uses Apple Gamut as the color primaries, and an Apple defined Log curve as a transfer function.
	AVCaptureColorSpace_AppleLog2 AVCaptureColorSpace = 4
	// AVCaptureColorSpace_HLG_BT2020: The BT.2020 wide color space that uses Illuminant D65 as the white point and Hybrid Log-Gamma (HLG) as the transfer function.
	AVCaptureColorSpace_HLG_BT2020 AVCaptureColorSpace = 2
	// AVCaptureColorSpace_P3_D65: The P3 D65 wide color space that uses Illuminant D65 as the white point.
	AVCaptureColorSpace_P3_D65 AVCaptureColorSpace = 1
	// AVCaptureColorSpace_sRGB: The standard RGB color space.
	AVCaptureColorSpace_sRGB AVCaptureColorSpace = 0
)

func (e AVCaptureColorSpace) String() string {
	switch e {
	case AVCaptureColorSpace_AppleLog:
		return "AVCaptureColorSpace_AppleLog"
	case AVCaptureColorSpace_AppleLog2:
		return "AVCaptureColorSpace_AppleLog2"
	case AVCaptureColorSpace_HLG_BT2020:
		return "AVCaptureColorSpace_HLG_BT2020"
	case AVCaptureColorSpace_P3_D65:
		return "AVCaptureColorSpace_P3_D65"
	case AVCaptureColorSpace_sRGB:
		return "AVCaptureColorSpace_sRGB"
	default:
		return fmt.Sprintf("AVCaptureColorSpace(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/Position-swift.enum
type AVCaptureDevicePosition int

const (
	// AVCaptureDevicePositionBack: A position on the subject-facing side of an iOS device.
	AVCaptureDevicePositionBack AVCaptureDevicePosition = 1
	// AVCaptureDevicePositionFront: A position on the user-facing side of an iOS device.
	AVCaptureDevicePositionFront AVCaptureDevicePosition = 2
	// AVCaptureDevicePositionUnspecified: A position that’s unspecified.
	AVCaptureDevicePositionUnspecified AVCaptureDevicePosition = 0
)

func (e AVCaptureDevicePosition) String() string {
	switch e {
	case AVCaptureDevicePositionBack:
		return "AVCaptureDevicePositionBack"
	case AVCaptureDevicePositionFront:
		return "AVCaptureDevicePositionFront"
	case AVCaptureDevicePositionUnspecified:
		return "AVCaptureDevicePositionUnspecified"
	default:
		return fmt.Sprintf("AVCaptureDevicePosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TransportControlsPlaybackMode-swift.enum
type AVCaptureDeviceTransportControlsPlaybackMode int

const (
	// AVCaptureDeviceTransportControlsNotPlayingMode: A value that indicates that the tape transport isn’t threaded through the play head.
	AVCaptureDeviceTransportControlsNotPlayingMode AVCaptureDeviceTransportControlsPlaybackMode = 0
	// AVCaptureDeviceTransportControlsPlayingMode: A value that indicates that the tape transport is threaded through the play head.
	AVCaptureDeviceTransportControlsPlayingMode AVCaptureDeviceTransportControlsPlaybackMode = 1
)

func (e AVCaptureDeviceTransportControlsPlaybackMode) String() string {
	switch e {
	case AVCaptureDeviceTransportControlsNotPlayingMode:
		return "AVCaptureDeviceTransportControlsNotPlayingMode"
	case AVCaptureDeviceTransportControlsPlayingMode:
		return "AVCaptureDeviceTransportControlsPlayingMode"
	default:
		return fmt.Sprintf("AVCaptureDeviceTransportControlsPlaybackMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/ExposureMode-swift.enum
type AVCaptureExposureMode int

const (
	// AVCaptureExposureModeAutoExpose: A mode that automatically adjusts the exposure one time, and then locks exposure for the device.
	AVCaptureExposureModeAutoExpose AVCaptureExposureMode = 1
	// AVCaptureExposureModeContinuousAutoExposure: A mode that continuously monitors exposure levels and automatically adjusts exposure when necessary.
	AVCaptureExposureModeContinuousAutoExposure AVCaptureExposureMode = 2
	// AVCaptureExposureModeCustom: A mode where an app manually sets the exposure duration and ISO values.
	AVCaptureExposureModeCustom AVCaptureExposureMode = 3
	// AVCaptureExposureModeLocked: A mode that locks exposure for the device.
	AVCaptureExposureModeLocked AVCaptureExposureMode = 0
)

func (e AVCaptureExposureMode) String() string {
	switch e {
	case AVCaptureExposureModeAutoExpose:
		return "AVCaptureExposureModeAutoExpose"
	case AVCaptureExposureModeContinuousAutoExposure:
		return "AVCaptureExposureModeContinuousAutoExposure"
	case AVCaptureExposureModeCustom:
		return "AVCaptureExposureModeCustom"
	case AVCaptureExposureModeLocked:
		return "AVCaptureExposureModeLocked"
	default:
		return fmt.Sprintf("AVCaptureExposureMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FlashMode-swift.enum
type AVCaptureFlashMode int

const (
	// AVCaptureFlashModeAuto: A mode that indicates the device continuously monitors light levels and uses the flash when necessary.
	AVCaptureFlashModeAuto AVCaptureFlashMode = 2
	// AVCaptureFlashModeOff: A mode that indicates the flash is off.
	AVCaptureFlashModeOff AVCaptureFlashMode = 0
	// AVCaptureFlashModeOn: A mode that indicates the flash is on.
	AVCaptureFlashModeOn AVCaptureFlashMode = 1
)

func (e AVCaptureFlashMode) String() string {
	switch e {
	case AVCaptureFlashModeAuto:
		return "AVCaptureFlashModeAuto"
	case AVCaptureFlashModeOff:
		return "AVCaptureFlashModeOff"
	case AVCaptureFlashModeOn:
		return "AVCaptureFlashModeOn"
	default:
		return fmt.Sprintf("AVCaptureFlashMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/FocusMode-swift.enum
type AVCaptureFocusMode int

const (
	// AVCaptureFocusModeAutoFocus: A mode that automatically adjusts the focus one time, and then locks focus.
	AVCaptureFocusModeAutoFocus AVCaptureFocusMode = 1
	// AVCaptureFocusModeContinuousAutoFocus: A mode that continuously monitors focus and autofocuses when necessary.
	AVCaptureFocusModeContinuousAutoFocus AVCaptureFocusMode = 2
	// AVCaptureFocusModeLocked: A mode that locks device focus.
	AVCaptureFocusModeLocked AVCaptureFocusMode = 0
)

func (e AVCaptureFocusMode) String() string {
	switch e {
	case AVCaptureFocusModeAutoFocus:
		return "AVCaptureFocusModeAutoFocus"
	case AVCaptureFocusModeContinuousAutoFocus:
		return "AVCaptureFocusModeContinuousAutoFocus"
	case AVCaptureFocusModeLocked:
		return "AVCaptureFocusModeLocked"
	default:
		return fmt.Sprintf("AVCaptureFocusMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/LensStabilizationStatus
type AVCaptureLensStabilizationStatus int

const (
	// AVCaptureLensStabilizationStatusActive: Lens stabilization was active for the full duration of the photo capture.
	AVCaptureLensStabilizationStatusActive AVCaptureLensStabilizationStatus = 2
	// AVCaptureLensStabilizationStatusOff: Lens stabilization isn’t specified for this photo capture.
	AVCaptureLensStabilizationStatusOff AVCaptureLensStabilizationStatus = 1
	// AVCaptureLensStabilizationStatusOutOfRange: Lens stabilization was enabled for the photo capture, but device motion or capture duration exceeded the stabilization module’s correction limits.
	AVCaptureLensStabilizationStatusOutOfRange AVCaptureLensStabilizationStatus = 3
	// AVCaptureLensStabilizationStatusUnavailable: Lens stabilization was temporarily unavailable during the photo capture.
	AVCaptureLensStabilizationStatusUnavailable AVCaptureLensStabilizationStatus = 4
	// AVCaptureLensStabilizationStatusUnsupported: Lens stabilization isn’t available on the device or device configuration that captured this photo.
	AVCaptureLensStabilizationStatusUnsupported AVCaptureLensStabilizationStatus = 0
)

func (e AVCaptureLensStabilizationStatus) String() string {
	switch e {
	case AVCaptureLensStabilizationStatusActive:
		return "AVCaptureLensStabilizationStatusActive"
	case AVCaptureLensStabilizationStatusOff:
		return "AVCaptureLensStabilizationStatusOff"
	case AVCaptureLensStabilizationStatusOutOfRange:
		return "AVCaptureLensStabilizationStatusOutOfRange"
	case AVCaptureLensStabilizationStatusUnavailable:
		return "AVCaptureLensStabilizationStatusUnavailable"
	case AVCaptureLensStabilizationStatusUnsupported:
		return "AVCaptureLensStabilizationStatusUnsupported"
	default:
		return fmt.Sprintf("AVCaptureLensStabilizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/MicrophoneMode
type AVCaptureMicrophoneMode int

const (
	// AVCaptureMicrophoneModeStandard: A mode that processes microphone audio with standard voice DSP.
	AVCaptureMicrophoneModeStandard AVCaptureMicrophoneMode = 0
	// AVCaptureMicrophoneModeVoiceIsolation: A mode that processes microphone audio to isolate the voice and attenuate other signals.
	AVCaptureMicrophoneModeVoiceIsolation AVCaptureMicrophoneMode = 2
	// AVCaptureMicrophoneModeWideSpectrum: A mode that minimizes microphone audio processing to capture all sounds in the room.
	AVCaptureMicrophoneModeWideSpectrum AVCaptureMicrophoneMode = 1
)

func (e AVCaptureMicrophoneMode) String() string {
	switch e {
	case AVCaptureMicrophoneModeStandard:
		return "AVCaptureMicrophoneModeStandard"
	case AVCaptureMicrophoneModeVoiceIsolation:
		return "AVCaptureMicrophoneModeVoiceIsolation"
	case AVCaptureMicrophoneModeWideSpectrum:
		return "AVCaptureMicrophoneModeWideSpectrum"
	default:
		return fmt.Sprintf("AVCaptureMicrophoneMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultichannelAudioMode
type AVCaptureMultichannelAudioMode int

const (
	// AVCaptureMultichannelAudioModeFirstOrderAmbisonics: An audio mode that indicates the recording uses first-order ambisonics.
	AVCaptureMultichannelAudioModeFirstOrderAmbisonics AVCaptureMultichannelAudioMode = 2
	// AVCaptureMultichannelAudioModeNone: A mode that indicates there’s no multichannel audio.
	AVCaptureMultichannelAudioModeNone AVCaptureMultichannelAudioMode = 0
	// AVCaptureMultichannelAudioModeStereo: A mode that indicates the recording uses stereo audio.
	AVCaptureMultichannelAudioModeStereo AVCaptureMultichannelAudioMode = 1
)

func (e AVCaptureMultichannelAudioMode) String() string {
	switch e {
	case AVCaptureMultichannelAudioModeFirstOrderAmbisonics:
		return "AVCaptureMultichannelAudioModeFirstOrderAmbisonics"
	case AVCaptureMultichannelAudioModeNone:
		return "AVCaptureMultichannelAudioModeNone"
	case AVCaptureMultichannelAudioModeStereo:
		return "AVCaptureMultichannelAudioModeStereo"
	default:
		return fmt.Sprintf("AVCaptureMultichannelAudioMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput/DataDroppedReason
type AVCaptureOutputDataDroppedReason int

const (
	// AVCaptureOutputDataDroppedReasonDiscontinuity: The system dropped data because the device providing data experienced a discontinuity, and the output lost an unknown number of data objects.
	AVCaptureOutputDataDroppedReasonDiscontinuity AVCaptureOutputDataDroppedReason = 3
	// AVCaptureOutputDataDroppedReasonLateData: The system dropped data because you’ve configured capture output to drop data when delegate queue is in a blocked state, and there’s data to deliver.
	AVCaptureOutputDataDroppedReasonLateData AVCaptureOutputDataDroppedReason = 1
	// AVCaptureOutputDataDroppedReasonNone: The system didn’t drop data.
	AVCaptureOutputDataDroppedReasonNone AVCaptureOutputDataDroppedReason = 0
	// AVCaptureOutputDataDroppedReasonOutOfBuffers: The system dropped data because the capture output exhausted its internal pool of memory buffers.
	AVCaptureOutputDataDroppedReasonOutOfBuffers AVCaptureOutputDataDroppedReason = 2
)

func (e AVCaptureOutputDataDroppedReason) String() string {
	switch e {
	case AVCaptureOutputDataDroppedReasonDiscontinuity:
		return "AVCaptureOutputDataDroppedReasonDiscontinuity"
	case AVCaptureOutputDataDroppedReasonLateData:
		return "AVCaptureOutputDataDroppedReasonLateData"
	case AVCaptureOutputDataDroppedReasonNone:
		return "AVCaptureOutputDataDroppedReasonNone"
	case AVCaptureOutputDataDroppedReasonOutOfBuffers:
		return "AVCaptureOutputDataDroppedReasonOutOfBuffers"
	default:
		return fmt.Sprintf("AVCaptureOutputDataDroppedReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/CaptureReadiness-swift.enum
type AVCapturePhotoOutputCaptureReadiness int

const (
	// AVCapturePhotoOutputCaptureReadinessNotReadyMomentarily: Indicates that the output isn’t ready to receive requests, but may be ready shortly.
	AVCapturePhotoOutputCaptureReadinessNotReadyMomentarily AVCapturePhotoOutputCaptureReadiness = 2
	// AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture: Indicates that the output isn’t ready to receive requests for a longer duration because it’s busy capturing.
	AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture AVCapturePhotoOutputCaptureReadiness = 3
	// AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing: Indicates that the output isn’t ready to receive requests for a longer duration because it’s busy processing.
	AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing AVCapturePhotoOutputCaptureReadiness = 4
	// AVCapturePhotoOutputCaptureReadinessReady: Indicates that the output is ready to receive new requests.
	AVCapturePhotoOutputCaptureReadinessReady AVCapturePhotoOutputCaptureReadiness = 1
	// AVCapturePhotoOutputCaptureReadinessSessionNotRunning: Indicates that the session isn’t running and the output isn’t ready to receive requests.
	AVCapturePhotoOutputCaptureReadinessSessionNotRunning AVCapturePhotoOutputCaptureReadiness = 0
)

func (e AVCapturePhotoOutputCaptureReadiness) String() string {
	switch e {
	case AVCapturePhotoOutputCaptureReadinessNotReadyMomentarily:
		return "AVCapturePhotoOutputCaptureReadinessNotReadyMomentarily"
	case AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture:
		return "AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForCapture"
	case AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing:
		return "AVCapturePhotoOutputCaptureReadinessNotReadyWaitingForProcessing"
	case AVCapturePhotoOutputCaptureReadinessReady:
		return "AVCapturePhotoOutputCaptureReadinessReady"
	case AVCapturePhotoOutputCaptureReadinessSessionNotRunning:
		return "AVCapturePhotoOutputCaptureReadinessSessionNotRunning"
	default:
		return fmt.Sprintf("AVCapturePhotoOutputCaptureReadiness(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCapturePhotoOutput/QualityPrioritization
type AVCapturePhotoQualityPrioritization int

const (
	// AVCapturePhotoQualityPrioritizationBalanced: Priority is balanced between photo quality and speed of delivery.
	AVCapturePhotoQualityPrioritizationBalanced AVCapturePhotoQualityPrioritization = 2
	// AVCapturePhotoQualityPrioritizationQuality: Photo quality is most important, even at the expense of shot-to-shot time.
	AVCapturePhotoQualityPrioritizationQuality AVCapturePhotoQualityPrioritization = 3
	// AVCapturePhotoQualityPrioritizationSpeed: Speed of photo delivery is most important, even at the expense of quality.
	AVCapturePhotoQualityPrioritizationSpeed AVCapturePhotoQualityPrioritization = 1
)

func (e AVCapturePhotoQualityPrioritization) String() string {
	switch e {
	case AVCapturePhotoQualityPrioritizationBalanced:
		return "AVCapturePhotoQualityPrioritizationBalanced"
	case AVCapturePhotoQualityPrioritizationQuality:
		return "AVCapturePhotoQualityPrioritizationQuality"
	case AVCapturePhotoQualityPrioritizationSpeed:
		return "AVCapturePhotoQualityPrioritizationSpeed"
	default:
		return fmt.Sprintf("AVCapturePhotoQualityPrioritization(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.struct
type AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions uint

const (
	// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged: Restrict switching to a fallback camera only when the device’s exposure mode changes.
	AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 4
	// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged: Restrict switching to a fallback camera only when the device’s focus mode changes.
	AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 2
	// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone: Disallow switching to a fallback camera.
	AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 0
	// AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged: Restrict switching to a fallback camera only when the device’s video zoom changes.
	AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions = 1
)

func (e AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) String() string {
	switch e {
	case AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged:
		return "AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionExposureModeChanged"
	case AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged:
		return "AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionFocusModeChanged"
	case AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone:
		return "AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone"
	case AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged:
		return "AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionVideoZoomChanged"
	default:
		return fmt.Sprintf("AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/PrimaryConstituentDeviceSwitchingBehavior-swift.enum
type AVCapturePrimaryConstituentDeviceSwitchingBehavior int

const (
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorAuto: The device automatically selects the best camera for the current scene.
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorAuto AVCapturePrimaryConstituentDeviceSwitchingBehavior = 1
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorLocked: The device locks camera switching to the active primary constituent device.
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorLocked AVCapturePrimaryConstituentDeviceSwitchingBehavior = 3
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorRestricted: The device restricts fallback camera selection to certain conditions.
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorRestricted AVCapturePrimaryConstituentDeviceSwitchingBehavior = 2
	// AVCapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported: The device doesn’t support constituent device switching.
	AVCapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported AVCapturePrimaryConstituentDeviceSwitchingBehavior = 0
)

func (e AVCapturePrimaryConstituentDeviceSwitchingBehavior) String() string {
	switch e {
	case AVCapturePrimaryConstituentDeviceSwitchingBehaviorAuto:
		return "AVCapturePrimaryConstituentDeviceSwitchingBehaviorAuto"
	case AVCapturePrimaryConstituentDeviceSwitchingBehaviorLocked:
		return "AVCapturePrimaryConstituentDeviceSwitchingBehaviorLocked"
	case AVCapturePrimaryConstituentDeviceSwitchingBehaviorRestricted:
		return "AVCapturePrimaryConstituentDeviceSwitchingBehaviorRestricted"
	case AVCapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported:
		return "AVCapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported"
	default:
		return fmt.Sprintf("AVCapturePrimaryConstituentDeviceSwitchingBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/InterruptionReason
type AVCaptureSessionInterruptionReason int

const (
	// AVCaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient: An interruption caused by the audio hardware temporarily being made unavailable (for example, for a phone call or alarm).
	AVCaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient AVCaptureSessionInterruptionReason = 2
	// AVCaptureSessionInterruptionReasonSensitiveContentMitigationActivated: An interruption caused by a [SCVideoStreamAnalyzer] when it detects sensitive content on an associated AVCaptureDeviceInput.
	AVCaptureSessionInterruptionReasonSensitiveContentMitigationActivated AVCaptureSessionInterruptionReason = 6
	// AVCaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient: An interruption caused by the video device temporarily being made unavailable (for example, when used by another capture session).
	AVCaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient AVCaptureSessionInterruptionReason = 3
	// AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure: An interruption due to system pressure, such as thermal duress.
	AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure AVCaptureSessionInterruptionReason = 5
	// AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground: An interruption caused by the app being sent to the background while using a camera.
	AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground AVCaptureSessionInterruptionReason = 1
	// AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps: An interruption caused when your app is running in Slide Over, Split View, or Picture in Picture mode on iPad.
	AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps AVCaptureSessionInterruptionReason = 4
)

func (e AVCaptureSessionInterruptionReason) String() string {
	switch e {
	case AVCaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient:
		return "AVCaptureSessionInterruptionReasonAudioDeviceInUseByAnotherClient"
	case AVCaptureSessionInterruptionReasonSensitiveContentMitigationActivated:
		return "AVCaptureSessionInterruptionReasonSensitiveContentMitigationActivated"
	case AVCaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient:
		return "AVCaptureSessionInterruptionReasonVideoDeviceInUseByAnotherClient"
	case AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure:
		return "AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableDueToSystemPressure"
	case AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground:
		return "AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableInBackground"
	case AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps:
		return "AVCaptureSessionInterruptionReasonVideoDeviceNotAvailableWithMultipleForegroundApps"
	default:
		return fmt.Sprintf("AVCaptureSessionInterruptionReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemPressureState-swift.class/Factors-swift.struct
type AVCaptureSystemPressureFactors uint

const (
	// AVCaptureSystemPressureFactorCameraTemperature: The camera module is operating at an elevated temperature.
	AVCaptureSystemPressureFactorCameraTemperature AVCaptureSystemPressureFactors = 8
	// AVCaptureSystemPressureFactorDepthModuleTemperature: The module capturing depth information is operating at an elevated temperature.
	AVCaptureSystemPressureFactorDepthModuleTemperature AVCaptureSystemPressureFactors = 4
	// AVCaptureSystemPressureFactorNone: System pressure is currently nominal.
	AVCaptureSystemPressureFactorNone AVCaptureSystemPressureFactors = 0
	// AVCaptureSystemPressureFactorPeakPower: The system’s peak power requirements exceed the battery’s current capacity.
	AVCaptureSystemPressureFactorPeakPower AVCaptureSystemPressureFactors = 2
	// AVCaptureSystemPressureFactorSystemTemperature: The entire system is under elevated thermal load.
	AVCaptureSystemPressureFactorSystemTemperature AVCaptureSystemPressureFactors = 1
)

func (e AVCaptureSystemPressureFactors) String() string {
	switch e {
	case AVCaptureSystemPressureFactorCameraTemperature:
		return "AVCaptureSystemPressureFactorCameraTemperature"
	case AVCaptureSystemPressureFactorDepthModuleTemperature:
		return "AVCaptureSystemPressureFactorDepthModuleTemperature"
	case AVCaptureSystemPressureFactorNone:
		return "AVCaptureSystemPressureFactorNone"
	case AVCaptureSystemPressureFactorPeakPower:
		return "AVCaptureSystemPressureFactorPeakPower"
	case AVCaptureSystemPressureFactorSystemTemperature:
		return "AVCaptureSystemPressureFactorSystemTemperature"
	default:
		return fmt.Sprintf("AVCaptureSystemPressureFactors(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/SystemUserInterface
type AVCaptureSystemUserInterface int

const (
	// AVCaptureSystemUserInterfaceMicrophoneModes: The system user interface for selecting microphone modes.
	AVCaptureSystemUserInterfaceMicrophoneModes AVCaptureSystemUserInterface = 2
	// AVCaptureSystemUserInterfaceVideoEffects: The system user interface for changing the state of video effects.
	AVCaptureSystemUserInterfaceVideoEffects AVCaptureSystemUserInterface = 1
)

func (e AVCaptureSystemUserInterface) String() string {
	switch e {
	case AVCaptureSystemUserInterfaceMicrophoneModes:
		return "AVCaptureSystemUserInterfaceMicrophoneModes"
	case AVCaptureSystemUserInterfaceVideoEffects:
		return "AVCaptureSystemUserInterfaceVideoEffects"
	default:
		return fmt.Sprintf("AVCaptureSystemUserInterface(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecodeGenerator/SynchronizationStatus
type AVCaptureTimecodeGeneratorSynchronizationStatus int

const (
	// AVCaptureTimecodeGeneratorSynchronizationStatusNotRequired: The timecode generator does not require active synchronization for a given source.
	AVCaptureTimecodeGeneratorSynchronizationStatusNotRequired AVCaptureTimecodeGeneratorSynchronizationStatus = 7
	// AVCaptureTimecodeGeneratorSynchronizationStatusSourceSelected: A timecode source has been selected, but synchronization has not yet started.
	AVCaptureTimecodeGeneratorSynchronizationStatusSourceSelected AVCaptureTimecodeGeneratorSynchronizationStatus = 1
	// AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnavailable: The timecode generator has failed to establish a connection with a given source.
	AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnavailable AVCaptureTimecodeGeneratorSynchronizationStatus = 5
	// AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnsupported: The timecode generator is receiving data from the source in an unrecognized format.
	AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnsupported AVCaptureTimecodeGeneratorSynchronizationStatus = 6
	// AVCaptureTimecodeGeneratorSynchronizationStatusSynchronized: The timecode generator is successfully synchronized to the selected source, maintaining active timing alignment.
	AVCaptureTimecodeGeneratorSynchronizationStatusSynchronized AVCaptureTimecodeGeneratorSynchronizationStatus = 3
	// AVCaptureTimecodeGeneratorSynchronizationStatusSynchronizing: The timecode generator is actively synchronizing to the selected source.
	AVCaptureTimecodeGeneratorSynchronizationStatusSynchronizing AVCaptureTimecodeGeneratorSynchronizationStatus = 2
	// AVCaptureTimecodeGeneratorSynchronizationStatusTimedOut: The synchronization has timed out.
	AVCaptureTimecodeGeneratorSynchronizationStatusTimedOut AVCaptureTimecodeGeneratorSynchronizationStatus = 4
	// AVCaptureTimecodeGeneratorSynchronizationStatusUnknown: The initial state before a source is selected or during error conditions.
	AVCaptureTimecodeGeneratorSynchronizationStatusUnknown AVCaptureTimecodeGeneratorSynchronizationStatus = 0
)

func (e AVCaptureTimecodeGeneratorSynchronizationStatus) String() string {
	switch e {
	case AVCaptureTimecodeGeneratorSynchronizationStatusNotRequired:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusNotRequired"
	case AVCaptureTimecodeGeneratorSynchronizationStatusSourceSelected:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusSourceSelected"
	case AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnavailable:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnavailable"
	case AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnsupported:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusSourceUnsupported"
	case AVCaptureTimecodeGeneratorSynchronizationStatusSynchronized:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusSynchronized"
	case AVCaptureTimecodeGeneratorSynchronizationStatusSynchronizing:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusSynchronizing"
	case AVCaptureTimecodeGeneratorSynchronizationStatusTimedOut:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusTimedOut"
	case AVCaptureTimecodeGeneratorSynchronizationStatusUnknown:
		return "AVCaptureTimecodeGeneratorSynchronizationStatusUnknown"
	default:
		return fmt.Sprintf("AVCaptureTimecodeGeneratorSynchronizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/SourceType-swift.enum
type AVCaptureTimecodeSourceType int

const (
	// AVCaptureTimecodeSourceTypeExternal: Synchronizes timecode to an external timecode data stream.
	AVCaptureTimecodeSourceTypeExternal AVCaptureTimecodeSourceType = 2
	// AVCaptureTimecodeSourceTypeFrameCount: No internal or external source is adopted.
	AVCaptureTimecodeSourceTypeFrameCount AVCaptureTimecodeSourceType = 0
	// AVCaptureTimecodeSourceTypeRealTimeClock: Synchronizes timecode to the system clock for real-time applications.
	AVCaptureTimecodeSourceTypeRealTimeClock AVCaptureTimecodeSourceType = 1
)

func (e AVCaptureTimecodeSourceType) String() string {
	switch e {
	case AVCaptureTimecodeSourceTypeExternal:
		return "AVCaptureTimecodeSourceTypeExternal"
	case AVCaptureTimecodeSourceTypeFrameCount:
		return "AVCaptureTimecodeSourceTypeFrameCount"
	case AVCaptureTimecodeSourceTypeRealTimeClock:
		return "AVCaptureTimecodeSourceTypeRealTimeClock"
	default:
		return fmt.Sprintf("AVCaptureTimecodeSourceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TorchMode-swift.enum
type AVCaptureTorchMode int

const (
	// AVCaptureTorchModeAuto: The capture device continuously monitors light levels and uses the torch when necessary.
	AVCaptureTorchModeAuto AVCaptureTorchMode = 2
	// AVCaptureTorchModeOff: The capture device torch is always off.
	AVCaptureTorchModeOff AVCaptureTorchMode = 0
	// AVCaptureTorchModeOn: The capture device torch is always on.
	AVCaptureTorchModeOn AVCaptureTorchMode = 1
)

func (e AVCaptureTorchMode) String() string {
	switch e {
	case AVCaptureTorchModeAuto:
		return "AVCaptureTorchModeAuto"
	case AVCaptureTorchModeOff:
		return "AVCaptureTorchModeOff"
	case AVCaptureTorchModeOn:
		return "AVCaptureTorchModeOn"
	default:
		return fmt.Sprintf("AVCaptureTorchMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoOrientation
type AVCaptureVideoOrientation int

const (
	// Deprecated.
	AVCaptureVideoOrientationLandscapeLeft AVCaptureVideoOrientation = 4
	// Deprecated.
	AVCaptureVideoOrientationLandscapeRight AVCaptureVideoOrientation = 3
	// Deprecated.
	AVCaptureVideoOrientationPortrait AVCaptureVideoOrientation = 1
	// Deprecated.
	AVCaptureVideoOrientationPortraitUpsideDown AVCaptureVideoOrientation = 2
)

func (e AVCaptureVideoOrientation) String() string {
	switch e {
	case AVCaptureVideoOrientationLandscapeLeft:
		return "AVCaptureVideoOrientationLandscapeLeft"
	case AVCaptureVideoOrientationLandscapeRight:
		return "AVCaptureVideoOrientationLandscapeRight"
	case AVCaptureVideoOrientationPortrait:
		return "AVCaptureVideoOrientationPortrait"
	case AVCaptureVideoOrientationPortraitUpsideDown:
		return "AVCaptureVideoOrientationPortraitUpsideDown"
	default:
		return fmt.Sprintf("AVCaptureVideoOrientation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoStabilizationMode
type AVCaptureVideoStabilizationMode int

const (
	// AVCaptureVideoStabilizationModeAuto: A mode that indicates the system chooses the most appropriate video stabilization mode for the device and format.
	AVCaptureVideoStabilizationModeAuto AVCaptureVideoStabilizationMode = -1
	// AVCaptureVideoStabilizationModeCinematic: A mode that uses the cinematic stabilization algorithm.
	AVCaptureVideoStabilizationModeCinematic AVCaptureVideoStabilizationMode = 2
	// AVCaptureVideoStabilizationModeCinematicExtended: A mode that uses the extended cinematic stabilization algorithm.
	AVCaptureVideoStabilizationModeCinematicExtended AVCaptureVideoStabilizationMode = 3
	// AVCaptureVideoStabilizationModeCinematicExtendedEnhanced: A mode that stabilizes video using the enhanced extended cinematic stabilization algorithm.
	AVCaptureVideoStabilizationModeCinematicExtendedEnhanced AVCaptureVideoStabilizationMode = 5
	// AVCaptureVideoStabilizationModeLowLatency: Indicates that video should be stabilized using the low latency stabilization algorithm.
	AVCaptureVideoStabilizationModeLowLatency AVCaptureVideoStabilizationMode = 6
	// AVCaptureVideoStabilizationModeOff: A mode that doesn’t stabilize video capture.
	AVCaptureVideoStabilizationModeOff AVCaptureVideoStabilizationMode = 0
	// AVCaptureVideoStabilizationModePreviewOptimized: A mode that uses the preview optimized stabilization algorithm.
	AVCaptureVideoStabilizationModePreviewOptimized AVCaptureVideoStabilizationMode = 4
	// AVCaptureVideoStabilizationModeStandard: A mode that uses the standard algorithm.
	AVCaptureVideoStabilizationModeStandard AVCaptureVideoStabilizationMode = 1
)

func (e AVCaptureVideoStabilizationMode) String() string {
	switch e {
	case AVCaptureVideoStabilizationModeAuto:
		return "AVCaptureVideoStabilizationModeAuto"
	case AVCaptureVideoStabilizationModeCinematic:
		return "AVCaptureVideoStabilizationModeCinematic"
	case AVCaptureVideoStabilizationModeCinematicExtended:
		return "AVCaptureVideoStabilizationModeCinematicExtended"
	case AVCaptureVideoStabilizationModeCinematicExtendedEnhanced:
		return "AVCaptureVideoStabilizationModeCinematicExtendedEnhanced"
	case AVCaptureVideoStabilizationModeLowLatency:
		return "AVCaptureVideoStabilizationModeLowLatency"
	case AVCaptureVideoStabilizationModeOff:
		return "AVCaptureVideoStabilizationModeOff"
	case AVCaptureVideoStabilizationModePreviewOptimized:
		return "AVCaptureVideoStabilizationModePreviewOptimized"
	case AVCaptureVideoStabilizationModeStandard:
		return "AVCaptureVideoStabilizationModeStandard"
	default:
		return fmt.Sprintf("AVCaptureVideoStabilizationMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/WhiteBalanceMode-swift.enum
type AVCaptureWhiteBalanceMode int

const (
	// AVCaptureWhiteBalanceModeAutoWhiteBalance: A mode that automatically manages white balance.
	AVCaptureWhiteBalanceModeAutoWhiteBalance AVCaptureWhiteBalanceMode = 1
	// AVCaptureWhiteBalanceModeContinuousAutoWhiteBalance: A mode that continuously monitors white balance and adjusts when necessary.
	AVCaptureWhiteBalanceModeContinuousAutoWhiteBalance AVCaptureWhiteBalanceMode = 2
	// AVCaptureWhiteBalanceModeLocked: A mode that locks the white balance state.
	AVCaptureWhiteBalanceModeLocked AVCaptureWhiteBalanceMode = 0
)

func (e AVCaptureWhiteBalanceMode) String() string {
	switch e {
	case AVCaptureWhiteBalanceModeAutoWhiteBalance:
		return "AVCaptureWhiteBalanceModeAutoWhiteBalance"
	case AVCaptureWhiteBalanceModeContinuousAutoWhiteBalance:
		return "AVCaptureWhiteBalanceModeContinuousAutoWhiteBalance"
	case AVCaptureWhiteBalanceModeLocked:
		return "AVCaptureWhiteBalanceModeLocked"
	default:
		return fmt.Sprintf("AVCaptureWhiteBalanceMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVContentAuthorizationStatus
type AVContentAuthorizationStatus int

const (
	// AVContentAuthorizationBusy: The last call to request content authorization couldn’t be completed because another asset is currently attempting authorization.
	AVContentAuthorizationBusy AVContentAuthorizationStatus = 4
	// AVContentAuthorizationCancelled: The last call to request content authorization was cancelled by the user.
	AVContentAuthorizationCancelled AVContentAuthorizationStatus = 2
	// AVContentAuthorizationCompleted: The last completed call to request content authorization completed.
	AVContentAuthorizationCompleted AVContentAuthorizationStatus = 1
	// AVContentAuthorizationNotAvailable: The last call to request content authorization couldn’t be completed because there was no known mechanism by which to attempt authorization.
	AVContentAuthorizationNotAvailable AVContentAuthorizationStatus = 5
	// AVContentAuthorizationNotPossible: The last call to request content authorization couldn’t be completed in a non-recoverable way.
	AVContentAuthorizationNotPossible AVContentAuthorizationStatus = 6
	// AVContentAuthorizationTimedOut: The last call to request content authorization was cancelled because the timeout interval was reached.
	AVContentAuthorizationTimedOut AVContentAuthorizationStatus = 3
	// AVContentAuthorizationUnknown: The content authorization content request hasn’t completed.
	AVContentAuthorizationUnknown AVContentAuthorizationStatus = 0
)

func (e AVContentAuthorizationStatus) String() string {
	switch e {
	case AVContentAuthorizationBusy:
		return "AVContentAuthorizationBusy"
	case AVContentAuthorizationCancelled:
		return "AVContentAuthorizationCancelled"
	case AVContentAuthorizationCompleted:
		return "AVContentAuthorizationCompleted"
	case AVContentAuthorizationNotAvailable:
		return "AVContentAuthorizationNotAvailable"
	case AVContentAuthorizationNotPossible:
		return "AVContentAuthorizationNotPossible"
	case AVContentAuthorizationTimedOut:
		return "AVContentAuthorizationTimedOut"
	case AVContentAuthorizationUnknown:
		return "AVContentAuthorizationUnknown"
	default:
		return fmt.Sprintf("AVContentAuthorizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVContentKeyRequest/Status-swift.enum
type AVContentKeyRequestStatus int

const (
	// AVContentKeyRequestStatusCancelled: The key request was canceled.
	AVContentKeyRequestStatusCancelled AVContentKeyRequestStatus = 4
	// AVContentKeyRequestStatusFailed: The key request failed.
	AVContentKeyRequestStatusFailed AVContentKeyRequestStatus = 5
	// AVContentKeyRequestStatusReceivedResponse: The key request was received, and the key is in use.
	AVContentKeyRequestStatusReceivedResponse AVContentKeyRequestStatus = 1
	// AVContentKeyRequestStatusRenewed: The key request was renewed.
	AVContentKeyRequestStatusRenewed AVContentKeyRequestStatus = 2
	// AVContentKeyRequestStatusRequestingResponse: The key request was just created.
	AVContentKeyRequestStatusRequestingResponse AVContentKeyRequestStatus = 0
	// AVContentKeyRequestStatusRetried: The key request was retried.
	AVContentKeyRequestStatusRetried AVContentKeyRequestStatus = 3
)

func (e AVContentKeyRequestStatus) String() string {
	switch e {
	case AVContentKeyRequestStatusCancelled:
		return "AVContentKeyRequestStatusCancelled"
	case AVContentKeyRequestStatusFailed:
		return "AVContentKeyRequestStatusFailed"
	case AVContentKeyRequestStatusReceivedResponse:
		return "AVContentKeyRequestStatusReceivedResponse"
	case AVContentKeyRequestStatusRenewed:
		return "AVContentKeyRequestStatusRenewed"
	case AVContentKeyRequestStatusRequestingResponse:
		return "AVContentKeyRequestStatusRequestingResponse"
	case AVContentKeyRequestStatusRetried:
		return "AVContentKeyRequestStatusRetried"
	default:
		return fmt.Sprintf("AVContentKeyRequestStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorRateChangeOptions
type AVDelegatingPlaybackCoordinatorRateChangeOptions uint

const (
	// AVDelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately: Indicates that the coordinator should begin playback as soon as possible, regardless of other participant’s readiness or suspensions.
	AVDelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately AVDelegatingPlaybackCoordinatorRateChangeOptions = 1
)

func (e AVDelegatingPlaybackCoordinatorRateChangeOptions) String() string {
	switch e {
	case AVDelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately:
		return "AVDelegatingPlaybackCoordinatorRateChangeOptionPlayImmediately"
	default:
		return fmt.Sprintf("AVDelegatingPlaybackCoordinatorRateChangeOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVDelegatingPlaybackCoordinatorSeekOptions
type AVDelegatingPlaybackCoordinatorSeekOptions uint

const (
	// AVDelegatingPlaybackCoordinatorSeekOptionResumeImmediately: An option that Indicates that the coordinator needs to resume playback as soon as possible, regardless of other participant’s readiness or suspensions.
	AVDelegatingPlaybackCoordinatorSeekOptionResumeImmediately AVDelegatingPlaybackCoordinatorSeekOptions = 1
)

func (e AVDelegatingPlaybackCoordinatorSeekOptions) String() string {
	switch e {
	case AVDelegatingPlaybackCoordinatorSeekOptionResumeImmediately:
		return "AVDelegatingPlaybackCoordinatorSeekOptionResumeImmediately"
	default:
		return fmt.Sprintf("AVDelegatingPlaybackCoordinatorSeekOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Accuracy
type AVDepthDataAccuracy int

const (
	// AVDepthDataAccuracyAbsolute: Values within the depth map are absolutely accurate within the physical world.
	AVDepthDataAccuracyAbsolute AVDepthDataAccuracy = 1
	// AVDepthDataAccuracyRelative: Values within the depth data map are usable for foreground/background separation, but are not absolutely accurate in the physical world.
	AVDepthDataAccuracyRelative AVDepthDataAccuracy = 0
)

func (e AVDepthDataAccuracy) String() string {
	switch e {
	case AVDepthDataAccuracyAbsolute:
		return "AVDepthDataAccuracyAbsolute"
	case AVDepthDataAccuracyRelative:
		return "AVDepthDataAccuracyRelative"
	default:
		return fmt.Sprintf("AVDepthDataAccuracy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVDepthData/Quality
type AVDepthDataQuality int

const (
	// AVDepthDataQualityHigh: The depth map is a good candidate for rendering high-quality depth effects or reconstructing a 3D scene.
	AVDepthDataQualityHigh AVDepthDataQuality = 1
	// AVDepthDataQualityLow: The depth map is a poor candidate for rendering high-quality depth effects or reconstructing a 3D scene.
	AVDepthDataQualityLow AVDepthDataQuality = 0
)

func (e AVDepthDataQuality) String() string {
	switch e {
	case AVDepthDataQualityHigh:
		return "AVDepthDataQualityHigh"
	case AVDepthDataQualityLow:
		return "AVDepthDataQualityLow"
	default:
		return fmt.Sprintf("AVDepthDataQuality(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVError-swift.struct/Code
type AVError int

const (
	// AVErrorAirPlayControllerRequiresInternet: The AirPlay controller requires an internet connection to function.
	AVErrorAirPlayControllerRequiresInternet AVError = -11856
	// AVErrorAirPlayReceiverRequiresInternet: The AirPlay receiver requires an internet connection to function.
	AVErrorAirPlayReceiverRequiresInternet AVError = -11857
	// AVErrorAirPlayReceiverTemporarilyUnavailable: An AirPlay receiver is temporarily unavailable.
	AVErrorAirPlayReceiverTemporarilyUnavailable AVError = -11882
	// AVErrorApplicationIsNotAuthorized: The app isn’t authorized to play media.
	AVErrorApplicationIsNotAuthorized AVError = -11836
	// AVErrorApplicationIsNotAuthorizedToUseDevice: The user denied this app permission to capture media.
	AVErrorApplicationIsNotAuthorizedToUseDevice AVError = -11852
	AVErrorAutoWhiteBalanceNotLocked AVError = -11891
	// AVErrorCompositionTrackSegmentsNotContiguous: The composition can’t add the source media because it contains gaps.
	AVErrorCompositionTrackSegmentsNotContiguous AVError = -11824
	// AVErrorContentIsNotAuthorized: The user isn’t authorized to play the media.
	AVErrorContentIsNotAuthorized AVError = -11835
	// AVErrorContentIsProtected: The app isn’t authorized to open the media.
	AVErrorContentIsProtected AVError = -11831
	// AVErrorContentIsUnavailable: The captured content is unavailable.
	AVErrorContentIsUnavailable AVError = -11863
	AVErrorContentKeyInvalid AVError = -11889
	// AVErrorContentKeyRequestCancelled: The app canceled a request to retrieve a content key.
	AVErrorContentKeyRequestCancelled AVError = -11879
	AVErrorContentKeyRequestPlaybackDestinationDoesNotSupportDeviceIdentifierRandomization AVError = -11888
	// AVErrorContentNotUpdated: The system couldn’t update the captured content.
	AVErrorContentNotUpdated AVError = -11866
	// AVErrorCreateContentKeyRequestFailed: The app couldn’t create a content key request.
	AVErrorCreateContentKeyRequestFailed AVError = -11860
	// AVErrorDecodeFailed: The system failed to decode the media.
	AVErrorDecodeFailed AVError = -11821
	// AVErrorDecoderNotFound: The system can’t find a suitable decoder for the media.
	AVErrorDecoderNotFound AVError = -11833
	// AVErrorDecoderTemporarilyUnavailable: A suitable decoder for the media is temporarily available.
	AVErrorDecoderTemporarilyUnavailable AVError = -11839
	// AVErrorDeviceAlreadyUsedByAnotherSession: Your app can’t access the device because another session is currently using it.
	AVErrorDeviceAlreadyUsedByAnotherSession AVError = -11804
	// AVErrorDeviceInUseByAnotherApplication: Your app can’t access the device because another app is currently using it.
	AVErrorDeviceInUseByAnotherApplication AVError = -11815
	// AVErrorDeviceLockedForConfigurationByAnotherProcess: Your app can’t change device settings because another process currently controls the device.
	AVErrorDeviceLockedForConfigurationByAnotherProcess AVError = -11817
	// AVErrorDeviceNotConnected: You app can’t access the device because it isn’t connected.
	AVErrorDeviceNotConnected AVError = -11814
	// AVErrorDeviceWasDisconnected: A previously connected device is no longer accessible.
	AVErrorDeviceWasDisconnected AVError = -11808
	// AVErrorDiskFull: Recording stopped because the disk is full.
	AVErrorDiskFull AVError = -11807
	// AVErrorDisplayWasDisabled: Screen capture failed because the display was inactive.
	AVErrorDisplayWasDisabled AVError = -11845
	// AVErrorEncodeFailed: The system couldn’t encode the media data.
	AVErrorEncodeFailed AVError = -11883
	// AVErrorEncoderNotFound: The requested encoder isn’t found.
	AVErrorEncoderNotFound AVError = -11834
	// AVErrorEncoderTemporarilyUnavailable: An appropriate encoder isn’t currently available.
	AVErrorEncoderTemporarilyUnavailable AVError = -11840
	// AVErrorExportFailed: The requested export operation failed.
	AVErrorExportFailed AVError = -11820
	// AVErrorExternalPlaybackNotSupportedForAsset: The current asset doesn’t support playback.
	AVErrorExternalPlaybackNotSupportedForAsset AVError = -11870
	// AVErrorFailedToLoadMediaData: The system can’t load the requested media data.
	AVErrorFailedToLoadMediaData AVError = -11849
	// AVErrorFailedToLoadSampleData: The system can’t load the requested sample data.
	AVErrorFailedToLoadSampleData AVError = -11881
	// AVErrorFailedToParse: The system can’t parse the media.
	AVErrorFailedToParse AVError = -11853
	// AVErrorFileAlreadyExists: A file with the same name exists at the location and you can’t overwrite it.
	AVErrorFileAlreadyExists AVError = -11823
	// AVErrorFileFailedToParse: The file is corrupt or in an unrecognized format.
	AVErrorFileFailedToParse AVError = -11829
	// AVErrorFileFormatNotRecognized: The system can’t open the file because it’s in an unrecognized format.
	AVErrorFileFormatNotRecognized AVError = -11828
	// AVErrorFileTypeDoesNotSupportSampleReferences: The file type doesn’t support sample references.
	AVErrorFileTypeDoesNotSupportSampleReferences AVError = -11854
	AVErrorFollowExternalSyncDeviceTimedOut AVError = -11892
	// AVErrorFormatUnsupported: The current asset format isn’t supported.
	AVErrorFormatUnsupported AVError = -11864
	// AVErrorIncompatibleAsset: You can’t display the media because the device isn’t capable of playing the content.
	AVErrorIncompatibleAsset AVError = -11848
	// AVErrorIncorrectlyConfigured: The system is incorrectly configured for the requested operation.
	AVErrorIncorrectlyConfigured AVError = -11875
	// AVErrorInvalidCompositionTrackSegmentDuration: You can’t add the source media because its duration in the destination is invalid.
	AVErrorInvalidCompositionTrackSegmentDuration AVError = -11825
	// AVErrorInvalidCompositionTrackSegmentSourceDuration: You can’t add the source media because it has no duration.
	AVErrorInvalidCompositionTrackSegmentSourceDuration AVError = -11827
	// AVErrorInvalidCompositionTrackSegmentSourceStartTime: You can’t add the source media because its start time in the destination is invalid.
	AVErrorInvalidCompositionTrackSegmentSourceStartTime AVError = -11826
	// AVErrorInvalidOutputURLPathExtension: The path extension of the output URL is invalid.
	AVErrorInvalidOutputURLPathExtension AVError = -11843
	// AVErrorInvalidSampleCursor: An invalid sample cursor produced an error.
	AVErrorInvalidSampleCursor AVError = -11880
	// AVErrorInvalidSourceMedia: The system couldn’t read the source media.
	AVErrorInvalidSourceMedia AVError = -11822
	// AVErrorInvalidVideoComposition: You attempted to present an unsupported video composition.
	AVErrorInvalidVideoComposition AVError = -11841
	// AVErrorMalformedDepth: The depth data isn’t properly structured.
	AVErrorMalformedDepth AVError = -11865
	// AVErrorMaximumDurationReached: The recording stopped because it reached the file’s maximum duration.
	AVErrorMaximumDurationReached AVError = -11810
	// AVErrorMaximumFileSizeReached: The recording stopped because it reached the file’s maximum size.
	AVErrorMaximumFileSizeReached AVError = -11811
	// AVErrorMaximumNumberOfSamplesForFileFormatReached: The recording stopped because it reached the file’s maximum number of samples.
	AVErrorMaximumNumberOfSamplesForFileFormatReached AVError = -11813
	// AVErrorMaximumStillImageCaptureRequestsExceeded: Your app can’t take a photo because there are too many unfinished photo capture requests.
	AVErrorMaximumStillImageCaptureRequestsExceeded AVError = -11830
	// AVErrorMediaChanged: Recording stopped because the format of the source media changed.
	AVErrorMediaChanged AVError = -11809
	// AVErrorMediaDiscontinuity: Recording stopped because there was an interruption in the input media.
	AVErrorMediaDiscontinuity AVError = -11812
	AVErrorMediaExtensionConflict AVError = -11887
	AVErrorMediaExtensionDisabled AVError = -11886
	// AVErrorMediaServicesWereReset: The system couldn’t perform the operation because media services were unavailable.
	AVErrorMediaServicesWereReset AVError = 0
	// AVErrorNoCompatibleAlternatesForExternalDisplay: The system found no compatible external displays.
	AVErrorNoCompatibleAlternatesForExternalDisplay AVError = -11868
	// AVErrorNoDataCaptured: The recording failed because the system received no data.
	AVErrorNoDataCaptured AVError = -11805
	// AVErrorNoImageAtTime: No image is available in the media at the indicated time.
	AVErrorNoImageAtTime AVError = -11832
	// AVErrorNoLongerPlayable: The asset is no longer playable.
	AVErrorNoLongerPlayable AVError = -11867
	AVErrorNoSmartFramingsEnabled AVError = -11890
	// AVErrorNoSourceTrack: The asset doesn’t contain a source track.
	AVErrorNoSourceTrack AVError = -11869
	// AVErrorOperationCancelled: The asset handled a request to cancel loading a property value asynchronously.
	AVErrorOperationCancelled AVError = -11878
	// AVErrorOperationInterrupted: An interruption occurred while performing a reading or writing operation.
	AVErrorOperationInterrupted AVError = 0
	// AVErrorOperationNotAllowed: The requested operation isn’t allowed.
	AVErrorOperationNotAllowed AVError = -11862
	// AVErrorOperationNotSupportedForAsset: Your app attempted to perform an unsupported operation with the asset.
	AVErrorOperationNotSupportedForAsset AVError = -11838
	// AVErrorOperationNotSupportedForPreset: Your app attempted to perform an unsupported operation for the current preset.
	AVErrorOperationNotSupportedForPreset AVError = -11871
	// AVErrorOutOfMemory: The operation couldn’t finish because there isn’t enough memory available to process the media.
	AVErrorOutOfMemory AVError = -11801
	// AVErrorRecordingAlreadyInProgress: Your app attempted to start recording a movie file while an existing recording is underway.
	AVErrorRecordingAlreadyInProgress AVError = -11859
	// AVErrorReferenceForbiddenByReferencePolicy: The current reference restrictions prevent the system from loading referenced media.
	AVErrorReferenceForbiddenByReferencePolicy AVError = -11842
	// AVErrorRosettaNotInstalled: The system doesn’t have Rosetta installed and can’t perform the requested operation.
	AVErrorRosettaNotInstalled AVError = -11877
	// AVErrorSandboxExtensionDenied: The system denied issuing the sandbox extension.
	AVErrorSandboxExtensionDenied AVError = -11884
	// AVErrorScreenCaptureFailed: An unexpected problem occurred that prevented screen capture.
	AVErrorScreenCaptureFailed AVError = -11844
	// AVErrorSegmentStartedWithNonSyncSample: The operation attempted to write a new MPEG-4 segment that didn’t start with a sync sample.
	AVErrorSegmentStartedWithNonSyncSample AVError = -11876
	// AVErrorServerIncorrectlyConfigured: The configuration of the HTTP server that streams the media resource isn’t correct.
	AVErrorServerIncorrectlyConfigured AVError = -11850
	// AVErrorSessionConfigurationChanged: Recording stopped because the configuration of media sources and destinations changed.
	AVErrorSessionConfigurationChanged AVError = -11806
	// AVErrorSessionHardwareCostOverage: Your app requested too many camera hardware resources.
	AVErrorSessionHardwareCostOverage AVError = -11872
	// AVErrorSessionNotRunning: The recording couldn’t start because the session isn’t running.
	AVErrorSessionNotRunning AVError = -11803
	// AVErrorSessionWasInterrupted: The recording stopped because the system interrupted the audio session.
	AVErrorSessionWasInterrupted AVError = 0
	// AVErrorToneMappingFailed: The requested tone mapping failed.
	AVErrorToneMappingFailed AVError = -11885
	// AVErrorTorchLevelUnavailable: The specified torch level is valid but currently unavailable, possibly due to overheating.
	AVErrorTorchLevelUnavailable AVError = -11846
	// AVErrorUndecodableMediaData: The system couldn’t decode the media data.
	AVErrorUndecodableMediaData AVError = -11855
	// AVErrorUnknown: An unknown error occurred.
	AVErrorUnknown AVError = -11800
	// AVErrorUnsupportedDeviceActiveFormat: The capture session doesn’t support the camera device’s active format.
	AVErrorUnsupportedDeviceActiveFormat AVError = -11873
	// AVErrorUnsupportedOutputSettings: Your app requested unsupported output settings.
	AVErrorUnsupportedOutputSettings AVError = -11861
	// AVErrorVideoCompositorFailed: The compositor couldn’t composite video frames.
	AVErrorVideoCompositorFailed AVError = -11858
	// Deprecated.
	AVErrorDeviceIsNotAvailableInBackground AVError = 0
)

func (e AVError) String() string {
	switch e {
	case AVErrorAirPlayControllerRequiresInternet:
		return "AVErrorAirPlayControllerRequiresInternet"
	case AVErrorAirPlayReceiverRequiresInternet:
		return "AVErrorAirPlayReceiverRequiresInternet"
	case AVErrorAirPlayReceiverTemporarilyUnavailable:
		return "AVErrorAirPlayReceiverTemporarilyUnavailable"
	case AVErrorApplicationIsNotAuthorized:
		return "AVErrorApplicationIsNotAuthorized"
	case AVErrorApplicationIsNotAuthorizedToUseDevice:
		return "AVErrorApplicationIsNotAuthorizedToUseDevice"
	case AVErrorAutoWhiteBalanceNotLocked:
		return "AVErrorAutoWhiteBalanceNotLocked"
	case AVErrorCompositionTrackSegmentsNotContiguous:
		return "AVErrorCompositionTrackSegmentsNotContiguous"
	case AVErrorContentIsNotAuthorized:
		return "AVErrorContentIsNotAuthorized"
	case AVErrorContentIsProtected:
		return "AVErrorContentIsProtected"
	case AVErrorContentIsUnavailable:
		return "AVErrorContentIsUnavailable"
	case AVErrorContentKeyInvalid:
		return "AVErrorContentKeyInvalid"
	case AVErrorContentKeyRequestCancelled:
		return "AVErrorContentKeyRequestCancelled"
	case AVErrorContentKeyRequestPlaybackDestinationDoesNotSupportDeviceIdentifierRandomization:
		return "AVErrorContentKeyRequestPlaybackDestinationDoesNotSupportDeviceIdentifierRandomization"
	case AVErrorContentNotUpdated:
		return "AVErrorContentNotUpdated"
	case AVErrorCreateContentKeyRequestFailed:
		return "AVErrorCreateContentKeyRequestFailed"
	case AVErrorDecodeFailed:
		return "AVErrorDecodeFailed"
	case AVErrorDecoderNotFound:
		return "AVErrorDecoderNotFound"
	case AVErrorDecoderTemporarilyUnavailable:
		return "AVErrorDecoderTemporarilyUnavailable"
	case AVErrorDeviceAlreadyUsedByAnotherSession:
		return "AVErrorDeviceAlreadyUsedByAnotherSession"
	case AVErrorDeviceInUseByAnotherApplication:
		return "AVErrorDeviceInUseByAnotherApplication"
	case AVErrorDeviceLockedForConfigurationByAnotherProcess:
		return "AVErrorDeviceLockedForConfigurationByAnotherProcess"
	case AVErrorDeviceNotConnected:
		return "AVErrorDeviceNotConnected"
	case AVErrorDeviceWasDisconnected:
		return "AVErrorDeviceWasDisconnected"
	case AVErrorDiskFull:
		return "AVErrorDiskFull"
	case AVErrorDisplayWasDisabled:
		return "AVErrorDisplayWasDisabled"
	case AVErrorEncodeFailed:
		return "AVErrorEncodeFailed"
	case AVErrorEncoderNotFound:
		return "AVErrorEncoderNotFound"
	case AVErrorEncoderTemporarilyUnavailable:
		return "AVErrorEncoderTemporarilyUnavailable"
	case AVErrorExportFailed:
		return "AVErrorExportFailed"
	case AVErrorExternalPlaybackNotSupportedForAsset:
		return "AVErrorExternalPlaybackNotSupportedForAsset"
	case AVErrorFailedToLoadMediaData:
		return "AVErrorFailedToLoadMediaData"
	case AVErrorFailedToLoadSampleData:
		return "AVErrorFailedToLoadSampleData"
	case AVErrorFailedToParse:
		return "AVErrorFailedToParse"
	case AVErrorFileAlreadyExists:
		return "AVErrorFileAlreadyExists"
	case AVErrorFileFailedToParse:
		return "AVErrorFileFailedToParse"
	case AVErrorFileFormatNotRecognized:
		return "AVErrorFileFormatNotRecognized"
	case AVErrorFileTypeDoesNotSupportSampleReferences:
		return "AVErrorFileTypeDoesNotSupportSampleReferences"
	case AVErrorFollowExternalSyncDeviceTimedOut:
		return "AVErrorFollowExternalSyncDeviceTimedOut"
	case AVErrorFormatUnsupported:
		return "AVErrorFormatUnsupported"
	case AVErrorIncompatibleAsset:
		return "AVErrorIncompatibleAsset"
	case AVErrorIncorrectlyConfigured:
		return "AVErrorIncorrectlyConfigured"
	case AVErrorInvalidCompositionTrackSegmentDuration:
		return "AVErrorInvalidCompositionTrackSegmentDuration"
	case AVErrorInvalidCompositionTrackSegmentSourceDuration:
		return "AVErrorInvalidCompositionTrackSegmentSourceDuration"
	case AVErrorInvalidCompositionTrackSegmentSourceStartTime:
		return "AVErrorInvalidCompositionTrackSegmentSourceStartTime"
	case AVErrorInvalidOutputURLPathExtension:
		return "AVErrorInvalidOutputURLPathExtension"
	case AVErrorInvalidSampleCursor:
		return "AVErrorInvalidSampleCursor"
	case AVErrorInvalidSourceMedia:
		return "AVErrorInvalidSourceMedia"
	case AVErrorInvalidVideoComposition:
		return "AVErrorInvalidVideoComposition"
	case AVErrorMalformedDepth:
		return "AVErrorMalformedDepth"
	case AVErrorMaximumDurationReached:
		return "AVErrorMaximumDurationReached"
	case AVErrorMaximumFileSizeReached:
		return "AVErrorMaximumFileSizeReached"
	case AVErrorMaximumNumberOfSamplesForFileFormatReached:
		return "AVErrorMaximumNumberOfSamplesForFileFormatReached"
	case AVErrorMaximumStillImageCaptureRequestsExceeded:
		return "AVErrorMaximumStillImageCaptureRequestsExceeded"
	case AVErrorMediaChanged:
		return "AVErrorMediaChanged"
	case AVErrorMediaDiscontinuity:
		return "AVErrorMediaDiscontinuity"
	case AVErrorMediaExtensionConflict:
		return "AVErrorMediaExtensionConflict"
	case AVErrorMediaExtensionDisabled:
		return "AVErrorMediaExtensionDisabled"
	case AVErrorMediaServicesWereReset:
		return "AVErrorMediaServicesWereReset"
	case AVErrorNoCompatibleAlternatesForExternalDisplay:
		return "AVErrorNoCompatibleAlternatesForExternalDisplay"
	case AVErrorNoDataCaptured:
		return "AVErrorNoDataCaptured"
	case AVErrorNoImageAtTime:
		return "AVErrorNoImageAtTime"
	case AVErrorNoLongerPlayable:
		return "AVErrorNoLongerPlayable"
	case AVErrorNoSmartFramingsEnabled:
		return "AVErrorNoSmartFramingsEnabled"
	case AVErrorNoSourceTrack:
		return "AVErrorNoSourceTrack"
	case AVErrorOperationCancelled:
		return "AVErrorOperationCancelled"
	case AVErrorOperationNotAllowed:
		return "AVErrorOperationNotAllowed"
	case AVErrorOperationNotSupportedForAsset:
		return "AVErrorOperationNotSupportedForAsset"
	case AVErrorOperationNotSupportedForPreset:
		return "AVErrorOperationNotSupportedForPreset"
	case AVErrorOutOfMemory:
		return "AVErrorOutOfMemory"
	case AVErrorRecordingAlreadyInProgress:
		return "AVErrorRecordingAlreadyInProgress"
	case AVErrorReferenceForbiddenByReferencePolicy:
		return "AVErrorReferenceForbiddenByReferencePolicy"
	case AVErrorRosettaNotInstalled:
		return "AVErrorRosettaNotInstalled"
	case AVErrorSandboxExtensionDenied:
		return "AVErrorSandboxExtensionDenied"
	case AVErrorScreenCaptureFailed:
		return "AVErrorScreenCaptureFailed"
	case AVErrorSegmentStartedWithNonSyncSample:
		return "AVErrorSegmentStartedWithNonSyncSample"
	case AVErrorServerIncorrectlyConfigured:
		return "AVErrorServerIncorrectlyConfigured"
	case AVErrorSessionConfigurationChanged:
		return "AVErrorSessionConfigurationChanged"
	case AVErrorSessionHardwareCostOverage:
		return "AVErrorSessionHardwareCostOverage"
	case AVErrorSessionNotRunning:
		return "AVErrorSessionNotRunning"
	case AVErrorToneMappingFailed:
		return "AVErrorToneMappingFailed"
	case AVErrorTorchLevelUnavailable:
		return "AVErrorTorchLevelUnavailable"
	case AVErrorUndecodableMediaData:
		return "AVErrorUndecodableMediaData"
	case AVErrorUnknown:
		return "AVErrorUnknown"
	case AVErrorUnsupportedDeviceActiveFormat:
		return "AVErrorUnsupportedDeviceActiveFormat"
	case AVErrorUnsupportedOutputSettings:
		return "AVErrorUnsupportedOutputSettings"
	case AVErrorVideoCompositorFailed:
		return "AVErrorVideoCompositorFailed"
	default:
		return fmt.Sprintf("AVError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVExternalContentProtectionStatus
type AVExternalContentProtectionStatus int

const (
	// AVExternalContentProtectionStatusInsufficient: A status that indicates insufficient protections exists for display.
	AVExternalContentProtectionStatusInsufficient AVExternalContentProtectionStatus = 2
	// AVExternalContentProtectionStatusPending: A status that indicates content protections are pending.
	AVExternalContentProtectionStatusPending AVExternalContentProtectionStatus = 0
	// AVExternalContentProtectionStatusSufficient: A status that indicates sufficient protections exists for display.
	AVExternalContentProtectionStatusSufficient AVExternalContentProtectionStatus = 1
)

func (e AVExternalContentProtectionStatus) String() string {
	switch e {
	case AVExternalContentProtectionStatusInsufficient:
		return "AVExternalContentProtectionStatusInsufficient"
	case AVExternalContentProtectionStatusPending:
		return "AVExternalContentProtectionStatusPending"
	case AVExternalContentProtectionStatusSufficient:
		return "AVExternalContentProtectionStatusSufficient"
	default:
		return fmt.Sprintf("AVExternalContentProtectionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVExternalSyncDeviceStatus
type AVExternalSyncDeviceStatus int

const (
	// AVExternalSyncDeviceStatusActiveSync: Indicates that the AVExternalSyncDevice object is running and that the clock property on AVExternalSyncDevice is calibrated to the external sync signal.
	AVExternalSyncDeviceStatusActiveSync AVExternalSyncDeviceStatus = 3
	// AVExternalSyncDeviceStatusCalibrating: Indicates that the external sync signal is connected and that the AVExternalSyncDevice object is calibrating to follow.
	AVExternalSyncDeviceStatusCalibrating AVExternalSyncDeviceStatus = 2
	// AVExternalSyncDeviceStatusFreeRunSync: Indicates that the AVExternalSyncDevice was calibrated to follow the external sync, but the sync signal has been lost.
	AVExternalSyncDeviceStatusFreeRunSync AVExternalSyncDeviceStatus = 4
	// AVExternalSyncDeviceStatusReady: Indicates that a device supporting external sync is connected, but calibration has not started.
	AVExternalSyncDeviceStatusReady AVExternalSyncDeviceStatus = 1
	// AVExternalSyncDeviceStatusUnavailable: Indicates that external sync signal is not connected, or has transitioned to a state that is not recoverable.
	AVExternalSyncDeviceStatusUnavailable AVExternalSyncDeviceStatus = 0
)

func (e AVExternalSyncDeviceStatus) String() string {
	switch e {
	case AVExternalSyncDeviceStatusActiveSync:
		return "AVExternalSyncDeviceStatusActiveSync"
	case AVExternalSyncDeviceStatusCalibrating:
		return "AVExternalSyncDeviceStatusCalibrating"
	case AVExternalSyncDeviceStatusFreeRunSync:
		return "AVExternalSyncDeviceStatusFreeRunSync"
	case AVExternalSyncDeviceStatusReady:
		return "AVExternalSyncDeviceStatusReady"
	case AVExternalSyncDeviceStatusUnavailable:
		return "AVExternalSyncDeviceStatusUnavailable"
	default:
		return fmt.Sprintf("AVExternalSyncDeviceStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVKeyValueStatus
type AVKeyValueStatus int

const (
	// Deprecated.
	AVKeyValueStatusCancelled AVKeyValueStatus = 4
	// Deprecated.
	AVKeyValueStatusFailed AVKeyValueStatus = 3
	// Deprecated.
	AVKeyValueStatusLoaded AVKeyValueStatus = 2
	// Deprecated.
	AVKeyValueStatusLoading AVKeyValueStatus = 1
	// Deprecated.
	AVKeyValueStatusUnknown AVKeyValueStatus = 0
)

func (e AVKeyValueStatus) String() string {
	switch e {
	case AVKeyValueStatusCancelled:
		return "AVKeyValueStatusCancelled"
	case AVKeyValueStatusFailed:
		return "AVKeyValueStatusFailed"
	case AVKeyValueStatusLoaded:
		return "AVKeyValueStatusLoaded"
	case AVKeyValueStatusLoading:
		return "AVKeyValueStatusLoading"
	case AVKeyValueStatusUnknown:
		return "AVKeyValueStatusUnknown"
	default:
		return fmt.Sprintf("AVKeyValueStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVMovieWritingOptions
type AVMovieWritingOptions uint

const (
	// AVMovieWritingAddMovieHeaderToDestination: The new movie header overwrites any existing movie header.
	AVMovieWritingAddMovieHeaderToDestination AVMovieWritingOptions = 0
	// AVMovieWritingTruncateDestinationToMovieHeaderOnly: The movie header overwrites all existing data and creates a pure reference movie file.
	AVMovieWritingTruncateDestinationToMovieHeaderOnly AVMovieWritingOptions = 1
)

func (e AVMovieWritingOptions) String() string {
	switch e {
	case AVMovieWritingAddMovieHeaderToDestination:
		return "AVMovieWritingAddMovieHeaderToDestination"
	case AVMovieWritingTruncateDestinationToMovieHeaderOnly:
		return "AVMovieWritingTruncateDestinationToMovieHeaderOnly"
	default:
		return fmt.Sprintf("AVMovieWritingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/ActionAtItemEnd-swift.enum
type AVPlayerActionAtItemEnd int

const (
	// AVPlayerActionAtItemEndAdvance: The player should advance to the next item, if there is one.
	AVPlayerActionAtItemEndAdvance AVPlayerActionAtItemEnd = 0
	// AVPlayerActionAtItemEndNone: The player should do nothing.
	AVPlayerActionAtItemEndNone AVPlayerActionAtItemEnd = 2
	// AVPlayerActionAtItemEndPause: The player should pause playing.
	AVPlayerActionAtItemEndPause AVPlayerActionAtItemEnd = 1
)

func (e AVPlayerActionAtItemEnd) String() string {
	switch e {
	case AVPlayerActionAtItemEndAdvance:
		return "AVPlayerActionAtItemEndAdvance"
	case AVPlayerActionAtItemEndNone:
		return "AVPlayerActionAtItemEndNone"
	case AVPlayerActionAtItemEndPause:
		return "AVPlayerActionAtItemEndPause"
	default:
		return fmt.Sprintf("AVPlayerActionAtItemEnd(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerAudiovisualBackgroundPlaybackPolicy
type AVPlayerAudiovisualBackgroundPlaybackPolicy int

const (
	// AVPlayerAudiovisualBackgroundPlaybackPolicyAutomatic: The system decides whether playback continues.
	AVPlayerAudiovisualBackgroundPlaybackPolicyAutomatic AVPlayerAudiovisualBackgroundPlaybackPolicy = 1
	// AVPlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible: The app continues playback, if possible.
	AVPlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible AVPlayerAudiovisualBackgroundPlaybackPolicy = 3
	// AVPlayerAudiovisualBackgroundPlaybackPolicyPauses: The app pauses playback.
	AVPlayerAudiovisualBackgroundPlaybackPolicyPauses AVPlayerAudiovisualBackgroundPlaybackPolicy = 2
)

func (e AVPlayerAudiovisualBackgroundPlaybackPolicy) String() string {
	switch e {
	case AVPlayerAudiovisualBackgroundPlaybackPolicyAutomatic:
		return "AVPlayerAudiovisualBackgroundPlaybackPolicyAutomatic"
	case AVPlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible:
		return "AVPlayerAudiovisualBackgroundPlaybackPolicyContinuesIfPossible"
	case AVPlayerAudiovisualBackgroundPlaybackPolicyPauses:
		return "AVPlayerAudiovisualBackgroundPlaybackPolicyPauses"
	default:
		return fmt.Sprintf("AVPlayerAudiovisualBackgroundPlaybackPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/HDRMode
type AVPlayerHDRMode int

const (
	// Deprecated.
	AVPlayerHDRModeDolbyVision AVPlayerHDRMode = 4
	// Deprecated.
	AVPlayerHDRModeHDR10 AVPlayerHDRMode = 2
	// Deprecated.
	AVPlayerHDRModeHLG AVPlayerHDRMode = 1
)

func (e AVPlayerHDRMode) String() string {
	switch e {
	case AVPlayerHDRModeDolbyVision:
		return "AVPlayerHDRModeDolbyVision"
	case AVPlayerHDRModeHDR10:
		return "AVPlayerHDRModeHDR10"
	case AVPlayerHDRModeHLG:
		return "AVPlayerHDRModeHLG"
	default:
		return fmt.Sprintf("AVPlayerHDRMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEventAssetListResponseStatus
type AVPlayerInterstitialEventAssetListResponseStatus int

const (
	// AVPlayerInterstitialEventAssetListResponseStatusAvailable: Indicates that a valid asset list response is available.
	AVPlayerInterstitialEventAssetListResponseStatusAvailable AVPlayerInterstitialEventAssetListResponseStatus = 0
	// AVPlayerInterstitialEventAssetListResponseStatusCleared: Indicates that the system cleared the asset list response.
	AVPlayerInterstitialEventAssetListResponseStatusCleared AVPlayerInterstitialEventAssetListResponseStatus = 1
	// AVPlayerInterstitialEventAssetListResponseStatusUnavailable: Indicates that the asset list response is unavailable.
	AVPlayerInterstitialEventAssetListResponseStatusUnavailable AVPlayerInterstitialEventAssetListResponseStatus = 2
)

func (e AVPlayerInterstitialEventAssetListResponseStatus) String() string {
	switch e {
	case AVPlayerInterstitialEventAssetListResponseStatusAvailable:
		return "AVPlayerInterstitialEventAssetListResponseStatusAvailable"
	case AVPlayerInterstitialEventAssetListResponseStatusCleared:
		return "AVPlayerInterstitialEventAssetListResponseStatusCleared"
	case AVPlayerInterstitialEventAssetListResponseStatusUnavailable:
		return "AVPlayerInterstitialEventAssetListResponseStatusUnavailable"
	default:
		return fmt.Sprintf("AVPlayerInterstitialEventAssetListResponseStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/Restrictions-swift.struct
type AVPlayerInterstitialEventRestrictions uint

const (
	// AVPlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent: A restriction that indicates the event doesn’t allow seeking forward within an interstitial item.
	AVPlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent AVPlayerInterstitialEventRestrictions = 1
	// AVPlayerInterstitialEventRestrictionDefaultPolicy: The default restriction policy.
	AVPlayerInterstitialEventRestrictionDefaultPolicy AVPlayerInterstitialEventRestrictions = 0
	// AVPlayerInterstitialEventRestrictionNone: A value that indicates no restrictions on playback of primary or interstitial content.
	AVPlayerInterstitialEventRestrictionNone AVPlayerInterstitialEventRestrictions = 0
	// AVPlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement: A restriction that indicates the event doesn’t allow advancing the current time within an interstitial item.
	AVPlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement AVPlayerInterstitialEventRestrictions = 4
)

func (e AVPlayerInterstitialEventRestrictions) String() string {
	switch e {
	case AVPlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent:
		return "AVPlayerInterstitialEventRestrictionConstrainsSeekingForwardInPrimaryContent"
	case AVPlayerInterstitialEventRestrictionDefaultPolicy:
		return "AVPlayerInterstitialEventRestrictionDefaultPolicy"
	case AVPlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement:
		return "AVPlayerInterstitialEventRestrictionRequiresPlaybackAtPreferredRateForAdvancement"
	default:
		return fmt.Sprintf("AVPlayerInterstitialEventRestrictions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/SkippableEventState
type AVPlayerInterstitialEventSkippableEventState int

const (
	// AVPlayerInterstitialEventSkippableEventStateEligible: Indicates that the interstitial event is currently skippable.
	AVPlayerInterstitialEventSkippableEventStateEligible AVPlayerInterstitialEventSkippableEventState = 2
	// AVPlayerInterstitialEventSkippableEventStateNoLongerEligible: Indicates that the interstitial event is no longer eligible to be skipped.
	AVPlayerInterstitialEventSkippableEventStateNoLongerEligible AVPlayerInterstitialEventSkippableEventState = 3
	// AVPlayerInterstitialEventSkippableEventStateNotSkippable: Indicates that the interstitial event is not skippable.
	AVPlayerInterstitialEventSkippableEventStateNotSkippable AVPlayerInterstitialEventSkippableEventState = 0
	// AVPlayerInterstitialEventSkippableEventStateNotYetEligible: Indicates that the interstitial event will eventually become eligible to be skipped.
	AVPlayerInterstitialEventSkippableEventStateNotYetEligible AVPlayerInterstitialEventSkippableEventState = 1
)

func (e AVPlayerInterstitialEventSkippableEventState) String() string {
	switch e {
	case AVPlayerInterstitialEventSkippableEventStateEligible:
		return "AVPlayerInterstitialEventSkippableEventStateEligible"
	case AVPlayerInterstitialEventSkippableEventStateNoLongerEligible:
		return "AVPlayerInterstitialEventSkippableEventStateNoLongerEligible"
	case AVPlayerInterstitialEventSkippableEventStateNotSkippable:
		return "AVPlayerInterstitialEventSkippableEventStateNotSkippable"
	case AVPlayerInterstitialEventSkippableEventStateNotYetEligible:
		return "AVPlayerInterstitialEventSkippableEventStateNotYetEligible"
	default:
		return fmt.Sprintf("AVPlayerInterstitialEventSkippableEventState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerInterstitialEvent/TimelineOccupancy-swift.enum
type AVPlayerInterstitialEventTimelineOccupancy int

const (
	// AVPlayerInterstitialEventTimelineOccupancyFill: The event fills the integrated timeline with the duration of this event.
	AVPlayerInterstitialEventTimelineOccupancyFill AVPlayerInterstitialEventTimelineOccupancy = 1
	// AVPlayerInterstitialEventTimelineOccupancySinglePoint: The event occupies a single point on the integrated timeline.
	AVPlayerInterstitialEventTimelineOccupancySinglePoint AVPlayerInterstitialEventTimelineOccupancy = 0
)

func (e AVPlayerInterstitialEventTimelineOccupancy) String() string {
	switch e {
	case AVPlayerInterstitialEventTimelineOccupancyFill:
		return "AVPlayerInterstitialEventTimelineOccupancyFill"
	case AVPlayerInterstitialEventTimelineOccupancySinglePoint:
		return "AVPlayerInterstitialEventTimelineOccupancySinglePoint"
	default:
		return fmt.Sprintf("AVPlayerInterstitialEventTimelineOccupancy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/SegmentType-swift.enum
type AVPlayerItemSegmentType int

const (
	// AVPlayerItemSegmentTypeInterstitial: A segment that represents playback of an interstitial event.
	AVPlayerItemSegmentTypeInterstitial AVPlayerItemSegmentType = 1
	// AVPlayerItemSegmentTypePrimary: A segment that represents playback of a primary item.
	AVPlayerItemSegmentTypePrimary AVPlayerItemSegmentType = 0
)

func (e AVPlayerItemSegmentType) String() string {
	switch e {
	case AVPlayerItemSegmentTypeInterstitial:
		return "AVPlayerItemSegmentTypeInterstitial"
	case AVPlayerItemSegmentTypePrimary:
		return "AVPlayerItemSegmentTypePrimary"
	default:
		return fmt.Sprintf("AVPlayerItemSegmentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/Status-swift.enum
type AVPlayerItemStatus int

const (
	// AVPlayerItemStatusFailed: The item no longer plays due to an error.
	AVPlayerItemStatusFailed AVPlayerItemStatus = 2
	// AVPlayerItemStatusReadyToPlay: The item is ready to play.
	AVPlayerItemStatusReadyToPlay AVPlayerItemStatus = 1
	// AVPlayerItemStatusUnknown: The item’s status is unknown.
	AVPlayerItemStatusUnknown AVPlayerItemStatus = 0
)

func (e AVPlayerItemStatus) String() string {
	switch e {
	case AVPlayerItemStatusFailed:
		return "AVPlayerItemStatusFailed"
	case AVPlayerItemStatusReadyToPlay:
		return "AVPlayerItemStatusReadyToPlay"
	case AVPlayerItemStatusUnknown:
		return "AVPlayerItemStatusUnknown"
	default:
		return fmt.Sprintf("AVPlayerItemStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/ItemOrdering
type AVPlayerLooperItemOrdering int

const (
	// AVPlayerLooperItemOrderingLoopingItemsFollowExistingItems: Indicates to insert replica items after any existing items in the specified player’s queue.
	AVPlayerLooperItemOrderingLoopingItemsFollowExistingItems AVPlayerLooperItemOrdering = 1
	// AVPlayerLooperItemOrderingLoopingItemsPrecedeExistingItems: Indicates to insert replica items before any existing items in the specified player’s queue.
	AVPlayerLooperItemOrderingLoopingItemsPrecedeExistingItems AVPlayerLooperItemOrdering = 0
)

func (e AVPlayerLooperItemOrdering) String() string {
	switch e {
	case AVPlayerLooperItemOrderingLoopingItemsFollowExistingItems:
		return "AVPlayerLooperItemOrderingLoopingItemsFollowExistingItems"
	case AVPlayerLooperItemOrderingLoopingItemsPrecedeExistingItems:
		return "AVPlayerLooperItemOrderingLoopingItemsPrecedeExistingItems"
	default:
		return fmt.Sprintf("AVPlayerLooperItemOrdering(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerLooper/Status-swift.enum
type AVPlayerLooperStatus int

const (
	// AVPlayerLooperStatusCancelled: The app canceled looping on the player.
	AVPlayerLooperStatusCancelled AVPlayerLooperStatus = 3
	// AVPlayerLooperStatusFailed: The looper isn’t able to perform looping playback due to an error.
	AVPlayerLooperStatusFailed AVPlayerLooperStatus = 2
	// AVPlayerLooperStatusReady: The looper is ready to perform looping playback.
	AVPlayerLooperStatusReady AVPlayerLooperStatus = 1
	// AVPlayerLooperStatusUnknown: The status isn’t known.
	AVPlayerLooperStatusUnknown AVPlayerLooperStatus = 0
)

func (e AVPlayerLooperStatus) String() string {
	switch e {
	case AVPlayerLooperStatusCancelled:
		return "AVPlayerLooperStatusCancelled"
	case AVPlayerLooperStatusFailed:
		return "AVPlayerLooperStatusFailed"
	case AVPlayerLooperStatusReady:
		return "AVPlayerLooperStatusReady"
	case AVPlayerLooperStatusUnknown:
		return "AVPlayerLooperStatusUnknown"
	default:
		return fmt.Sprintf("AVPlayerLooperStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/NetworkResourcePriority-swift.enum
type AVPlayerNetworkResourcePriority int

const (
	// AVPlayerNetworkResourcePriorityDefault: The default priority level given to a player for loading network resources.
	AVPlayerNetworkResourcePriorityDefault AVPlayerNetworkResourcePriority = 0
	// AVPlayerNetworkResourcePriorityHigh: Indicates a high priority level for loading network resources.
	AVPlayerNetworkResourcePriorityHigh AVPlayerNetworkResourcePriority = 2
	// AVPlayerNetworkResourcePriorityLow: Indicates a low priority level for loading network resources.
	AVPlayerNetworkResourcePriorityLow AVPlayerNetworkResourcePriority = 1
)

func (e AVPlayerNetworkResourcePriority) String() string {
	switch e {
	case AVPlayerNetworkResourcePriorityDefault:
		return "AVPlayerNetworkResourcePriorityDefault"
	case AVPlayerNetworkResourcePriorityHigh:
		return "AVPlayerNetworkResourcePriorityHigh"
	case AVPlayerNetworkResourcePriorityLow:
		return "AVPlayerNetworkResourcePriorityLow"
	default:
		return fmt.Sprintf("AVPlayerNetworkResourcePriority(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/Status-swift.enum
type AVPlayerStatus int

const (
	// AVPlayerStatusFailed: A value that indicates the player can no longer play media due to an error.
	AVPlayerStatusFailed AVPlayerStatus = 2
	// AVPlayerStatusReadyToPlay: A value that indicates the player is ready to media.
	AVPlayerStatusReadyToPlay AVPlayerStatus = 1
	// AVPlayerStatusUnknown: A value that indicates a player hasn’t attempted to load media for playback.
	AVPlayerStatusUnknown AVPlayerStatus = 0
)

func (e AVPlayerStatus) String() string {
	switch e {
	case AVPlayerStatusFailed:
		return "AVPlayerStatusFailed"
	case AVPlayerStatusReadyToPlay:
		return "AVPlayerStatusReadyToPlay"
	case AVPlayerStatusUnknown:
		return "AVPlayerStatusUnknown"
	default:
		return fmt.Sprintf("AVPlayerStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVPlayer/TimeControlStatus-swift.enum
type AVPlayerTimeControlStatus int

const (
	// AVPlayerTimeControlStatusPaused: A state that indicates the player paused playback indefinitely.
	AVPlayerTimeControlStatusPaused AVPlayerTimeControlStatus = 0
	// AVPlayerTimeControlStatusPlaying: A state that indicates that the player is currently playing media.
	AVPlayerTimeControlStatusPlaying AVPlayerTimeControlStatus = 2
	// AVPlayerTimeControlStatusWaitingToPlayAtSpecifiedRate: A state that indicates that the player is waiting for network conditions to improve before it can start or resume playback.
	AVPlayerTimeControlStatusWaitingToPlayAtSpecifiedRate AVPlayerTimeControlStatus = 1
)

func (e AVPlayerTimeControlStatus) String() string {
	switch e {
	case AVPlayerTimeControlStatusPaused:
		return "AVPlayerTimeControlStatusPaused"
	case AVPlayerTimeControlStatusPlaying:
		return "AVPlayerTimeControlStatusPlaying"
	case AVPlayerTimeControlStatusWaitingToPlayAtSpecifiedRate:
		return "AVPlayerTimeControlStatusWaitingToPlayAtSpecifiedRate"
	default:
		return fmt.Sprintf("AVPlayerTimeControlStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRenderingStatus
type AVQueuedSampleBufferRenderingStatus int

const (
	// AVQueuedSampleBufferRenderingStatusFailed: The object can no longer render sample buffers because of an error.
	AVQueuedSampleBufferRenderingStatusFailed AVQueuedSampleBufferRenderingStatus = 2
	// AVQueuedSampleBufferRenderingStatusRendering: The object is rendering the sample buffer.
	AVQueuedSampleBufferRenderingStatusRendering AVQueuedSampleBufferRenderingStatus = 1
	// AVQueuedSampleBufferRenderingStatusUnknown: The object doesn’t have any sample buffers enqueued.
	AVQueuedSampleBufferRenderingStatusUnknown AVQueuedSampleBufferRenderingStatus = 0
)

func (e AVQueuedSampleBufferRenderingStatus) String() string {
	switch e {
	case AVQueuedSampleBufferRenderingStatusFailed:
		return "AVQueuedSampleBufferRenderingStatusFailed"
	case AVQueuedSampleBufferRenderingStatusRendering:
		return "AVQueuedSampleBufferRenderingStatusRendering"
	case AVQueuedSampleBufferRenderingStatusUnknown:
		return "AVQueuedSampleBufferRenderingStatusUnknown"
	default:
		return fmt.Sprintf("AVQueuedSampleBufferRenderingStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Direction-swift.enum
type AVSampleBufferRequestDirection int

const (
	// AVSampleBufferRequestDirectionForward: The number of following samples may be zero or greater.
	AVSampleBufferRequestDirectionForward AVSampleBufferRequestDirection = 0
	// AVSampleBufferRequestDirectionNone: A single sample will be loaded.
	AVSampleBufferRequestDirectionNone AVSampleBufferRequestDirection = 0
	// AVSampleBufferRequestDirectionReverse: The number of previous samples may be zero or greater.
	AVSampleBufferRequestDirectionReverse AVSampleBufferRequestDirection = -1
)

func (e AVSampleBufferRequestDirection) String() string {
	switch e {
	case AVSampleBufferRequestDirectionForward:
		return "AVSampleBufferRequestDirectionForward"
	case AVSampleBufferRequestDirectionReverse:
		return "AVSampleBufferRequestDirectionReverse"
	default:
		return fmt.Sprintf("AVSampleBufferRequestDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferRequest/Mode-swift.enum
type AVSampleBufferRequestMode int

const (
	// AVSampleBufferRequestModeImmediate: A mode that indicates that sample buffer creation requests load data as soon as possible.
	AVSampleBufferRequestModeImmediate AVSampleBufferRequestMode = 0
	// AVSampleBufferRequestModeOpportunistic: A mode that indicates that opportunistic sample buffer creation requests load data as soon as possible.
	AVSampleBufferRequestModeOpportunistic AVSampleBufferRequestMode = 2
	// AVSampleBufferRequestModeScheduled: A mode that indicates that sample buffer creation requests load data according to a scheduled deadline.
	AVSampleBufferRequestModeScheduled AVSampleBufferRequestMode = 1
)

func (e AVSampleBufferRequestMode) String() string {
	switch e {
	case AVSampleBufferRequestModeImmediate:
		return "AVSampleBufferRequestModeImmediate"
	case AVSampleBufferRequestModeOpportunistic:
		return "AVSampleBufferRequestModeOpportunistic"
	case AVSampleBufferRequestModeScheduled:
		return "AVSampleBufferRequestModeScheduled"
	default:
		return fmt.Sprintf("AVSampleBufferRequestMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVVariantPreferences
type AVVariantPreferences uint

const (
	// AVVariantPreferenceNone: Indicates that the player item uses the default behavior for determining variant playlist selection.
	AVVariantPreferenceNone AVVariantPreferences = 0
	// AVVariantPreferenceScalabilityToLosslessAudio: A preference that indicates the player item supports variant playlists that contain losslessly encoded audio when sufficient bandwidth is available.
	AVVariantPreferenceScalabilityToLosslessAudio AVVariantPreferences = 1
)

func (e AVVariantPreferences) String() string {
	switch e {
	case AVVariantPreferenceNone:
		return "AVVariantPreferenceNone"
	case AVVariantPreferenceScalabilityToLosslessAudio:
		return "AVVariantPreferenceScalabilityToLosslessAudio"
	default:
		return fmt.Sprintf("AVVariantPreferences(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/AVVideoFieldMode
type AVVideoFieldMode int

const (
	// AVVideoFieldModeBoth: A value that indicates that a video connection passes both the top and bottom video fields.
	AVVideoFieldModeBoth AVVideoFieldMode = 0
	// AVVideoFieldModeBottomOnly: A value that indicates that a video connection only passes the bottom video field.
	AVVideoFieldModeBottomOnly AVVideoFieldMode = 2
	// AVVideoFieldModeDeinterlace: A value that indicates that a video connection deinterlaces the top and bottom video fields.
	AVVideoFieldModeDeinterlace AVVideoFieldMode = 3
	// AVVideoFieldModeTopOnly: A value that indicates that a video connection only passes the top video field.
	AVVideoFieldModeTopOnly AVVideoFieldMode = 1
)

func (e AVVideoFieldMode) String() string {
	switch e {
	case AVVideoFieldModeBoth:
		return "AVVideoFieldModeBoth"
	case AVVideoFieldModeBottomOnly:
		return "AVVideoFieldModeBottomOnly"
	case AVVideoFieldModeDeinterlace:
		return "AVVideoFieldModeDeinterlace"
	case AVVideoFieldModeTopOnly:
		return "AVVideoFieldModeTopOnly"
	default:
		return fmt.Sprintf("AVVideoFieldMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionVideoOutputPreset
type CMTagCollectionVideoOutputPreset uint32

const (
	// KCMTagCollectionVideoOutputPreset_Monoscopic: An output preset for monoscopic video.
	KCMTagCollectionVideoOutputPreset_Monoscopic CMTagCollectionVideoOutputPreset = 0
	// KCMTagCollectionVideoOutputPreset_Stereoscopic: An output preset for stereoscopic video.
	KCMTagCollectionVideoOutputPreset_Stereoscopic CMTagCollectionVideoOutputPreset = 1
)

func (e CMTagCollectionVideoOutputPreset) String() string {
	switch e {
	case KCMTagCollectionVideoOutputPreset_Monoscopic:
		return "KCMTagCollectionVideoOutputPreset_Monoscopic"
	case KCMTagCollectionVideoOutputPreset_Stereoscopic:
		return "KCMTagCollectionVideoOutputPreset_Stereoscopic"
	default:
		return fmt.Sprintf("CMTagCollectionVideoOutputPreset(%d)", e)
	}
}

