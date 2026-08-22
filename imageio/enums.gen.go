// Code generated from Apple documentation for ImageIO. DO NOT EDIT.

package imageio

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ImageIO/CGImageAnimationStatus
type CGImageAnimationStatus int32

const ()

// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataErrors
type CGImageMetadataErrors int32

const (
	// KCGImageMetadataErrorBadArgument: An error that indicates a parameter was malformed or contained invalid data.
	KCGImageMetadataErrorBadArgument CGImageMetadataErrors = 2
	// KCGImageMetadataErrorConflictingArguments: An error that indicates an attempt to save conflicting metadata values.
	KCGImageMetadataErrorConflictingArguments CGImageMetadataErrors = 3
	// KCGImageMetadataErrorPrefixConflict: An error that indicates an attempt to register a namespace with a prefix that is different than the namespace’s existing prefix.
	KCGImageMetadataErrorPrefixConflict CGImageMetadataErrors = 4
	// KCGImageMetadataErrorUnknown: An error that indicates an unknown condition occurred.
	KCGImageMetadataErrorUnknown CGImageMetadataErrors = 0
	// KCGImageMetadataErrorUnsupportedFormat: An error that indicates the metadata was in an unsupported format.
	KCGImageMetadataErrorUnsupportedFormat CGImageMetadataErrors = 1
)

func (e CGImageMetadataErrors) String() string {
	switch e {
	case KCGImageMetadataErrorBadArgument:
		return "KCGImageMetadataErrorBadArgument"
	case KCGImageMetadataErrorConflictingArguments:
		return "KCGImageMetadataErrorConflictingArguments"
	case KCGImageMetadataErrorPrefixConflict:
		return "KCGImageMetadataErrorPrefixConflict"
	case KCGImageMetadataErrorUnknown:
		return "KCGImageMetadataErrorUnknown"
	case KCGImageMetadataErrorUnsupportedFormat:
		return "KCGImageMetadataErrorUnsupportedFormat"
	default:
		return fmt.Sprintf("CGImageMetadataErrors(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageIO/CGImageMetadataType
type CGImageMetadataType int32

const (
	// KCGImageMetadataTypeAlternateArray: An ordered array, in which all elements are alternates for the same value.
	KCGImageMetadataTypeAlternateArray CGImageMetadataType = 4
	// KCGImageMetadataTypeAlternateText: An alternate array, in which all elements are localized strings for the same value.
	KCGImageMetadataTypeAlternateText CGImageMetadataType = 5
	// KCGImageMetadataTypeArrayOrdered: An array that preserves the order of items.
	KCGImageMetadataTypeArrayOrdered CGImageMetadataType = 3
	// KCGImageMetadataTypeArrayUnordered: An array that doesn’t preserve the order of items.
	KCGImageMetadataTypeArrayUnordered CGImageMetadataType = 2
	// KCGImageMetadataTypeDefault: The default type for new tags.
	KCGImageMetadataTypeDefault CGImageMetadataType = 0
	// KCGImageMetadataTypeInvalid: An invalid metadata type.
	KCGImageMetadataTypeInvalid CGImageMetadataType = -1
	// KCGImageMetadataTypeString: A string value.
	KCGImageMetadataTypeString CGImageMetadataType = 1
	// KCGImageMetadataTypeStructure: A collection of keys and values.
	KCGImageMetadataTypeStructure CGImageMetadataType = 6
)

func (e CGImageMetadataType) String() string {
	switch e {
	case KCGImageMetadataTypeAlternateArray:
		return "KCGImageMetadataTypeAlternateArray"
	case KCGImageMetadataTypeAlternateText:
		return "KCGImageMetadataTypeAlternateText"
	case KCGImageMetadataTypeArrayOrdered:
		return "KCGImageMetadataTypeArrayOrdered"
	case KCGImageMetadataTypeArrayUnordered:
		return "KCGImageMetadataTypeArrayUnordered"
	case KCGImageMetadataTypeDefault:
		return "KCGImageMetadataTypeDefault"
	case KCGImageMetadataTypeInvalid:
		return "KCGImageMetadataTypeInvalid"
	case KCGImageMetadataTypeString:
		return "KCGImageMetadataTypeString"
	case KCGImageMetadataTypeStructure:
		return "KCGImageMetadataTypeStructure"
	default:
		return fmt.Sprintf("CGImageMetadataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageIO/CGImagePropertyOrientation
type CGImagePropertyOrientation uint32

const (
	// KCGImagePropertyOrientationDown: The encoded image data is rotated 180° from the image’s intended display orientation.
	KCGImagePropertyOrientationDown CGImagePropertyOrientation = 3
	// KCGImagePropertyOrientationDownMirrored: The encoded image data is vertically flipped from the image’s intended display orientation.
	KCGImagePropertyOrientationDownMirrored CGImagePropertyOrientation = 4
	// KCGImagePropertyOrientationLeft: The encoded image data is rotated 90° clockwise from the image’s intended display orientation.
	KCGImagePropertyOrientationLeft CGImagePropertyOrientation = 8
	// KCGImagePropertyOrientationLeftMirrored: The encoded image data is horizontally flipped and rotated 90° counter-clockwise from the image’s intended display orientation.
	KCGImagePropertyOrientationLeftMirrored CGImagePropertyOrientation = 5
	// KCGImagePropertyOrientationRight: The encoded image data is rotated 90° counter-clockwise from the image’s intended display orientation.
	KCGImagePropertyOrientationRight CGImagePropertyOrientation = 6
	// KCGImagePropertyOrientationRightMirrored: The encoded image data is horizontally flipped and rotated 90° clockwise from the image’s intended display orientation.
	KCGImagePropertyOrientationRightMirrored CGImagePropertyOrientation = 7
	// KCGImagePropertyOrientationUp: The encoded image data matches the image’s intended display orientation.
	KCGImagePropertyOrientationUp CGImagePropertyOrientation = 1
	// KCGImagePropertyOrientationUpMirrored: The encoded image data is horizontally flipped from the image’s intended display orientation.
	KCGImagePropertyOrientationUpMirrored CGImagePropertyOrientation = 2
)

func (e CGImagePropertyOrientation) String() string {
	switch e {
	case KCGImagePropertyOrientationDown:
		return "KCGImagePropertyOrientationDown"
	case KCGImagePropertyOrientationDownMirrored:
		return "KCGImagePropertyOrientationDownMirrored"
	case KCGImagePropertyOrientationLeft:
		return "KCGImagePropertyOrientationLeft"
	case KCGImagePropertyOrientationLeftMirrored:
		return "KCGImagePropertyOrientationLeftMirrored"
	case KCGImagePropertyOrientationRight:
		return "KCGImagePropertyOrientationRight"
	case KCGImagePropertyOrientationRightMirrored:
		return "KCGImagePropertyOrientationRightMirrored"
	case KCGImagePropertyOrientationUp:
		return "KCGImagePropertyOrientationUp"
	case KCGImagePropertyOrientationUpMirrored:
		return "KCGImagePropertyOrientationUpMirrored"
	default:
		return fmt.Sprintf("CGImagePropertyOrientation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ImageIO/CGImagePropertyTGACompression
type CGImagePropertyTGACompression uint32

const ()

// See: https://developer.apple.com/documentation/ImageIO/CGImageSourceStatus
type CGImageSourceStatus int32

const (
	// KCGImageStatusComplete: The operation is complete.
	KCGImageStatusComplete CGImageSourceStatus = 0
	// KCGImageStatusIncomplete: The operation is not complete
	KCGImageStatusIncomplete CGImageSourceStatus = -1
	// KCGImageStatusInvalidData: The data is not valid.
	KCGImageStatusInvalidData CGImageSourceStatus = -4
	// KCGImageStatusReadingHeader: The image source is reading the header.
	KCGImageStatusReadingHeader CGImageSourceStatus = -2
	// KCGImageStatusUnexpectedEOF: The end of the file occurred unexpectedly.
	KCGImageStatusUnexpectedEOF CGImageSourceStatus = -5
	// KCGImageStatusUnknownType: The image is an unknown type.
	KCGImageStatusUnknownType CGImageSourceStatus = -3
)

func (e CGImageSourceStatus) String() string {
	switch e {
	case KCGImageStatusComplete:
		return "KCGImageStatusComplete"
	case KCGImageStatusIncomplete:
		return "KCGImageStatusIncomplete"
	case KCGImageStatusInvalidData:
		return "KCGImageStatusInvalidData"
	case KCGImageStatusReadingHeader:
		return "KCGImageStatusReadingHeader"
	case KCGImageStatusUnexpectedEOF:
		return "KCGImageStatusUnexpectedEOF"
	case KCGImageStatusUnknownType:
		return "KCGImageStatusUnknownType"
	default:
		return fmt.Sprintf("CGImageSourceStatus(%d)", e)
	}
}
