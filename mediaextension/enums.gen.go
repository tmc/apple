// Code generated from Apple documentation for MediaExtension. DO NOT EDIT.

package mediaextension

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/MediaExtension/MEDecodeFrameStatus
type MEDecodeFrameStatus uint

const (
	// MEDecodeFrameFrameDropped: A frame decode operation status that indicates the system dropped the output of the frame for a reason other than an error.
	MEDecodeFrameFrameDropped MEDecodeFrameStatus = 1
	MEDecodeFrameNoStatus     MEDecodeFrameStatus = 0
)

func (e MEDecodeFrameStatus) String() string {
	switch e {
	case MEDecodeFrameFrameDropped:
		return "MEDecodeFrameFrameDropped"
	case MEDecodeFrameNoStatus:
		return "MEDecodeFrameNoStatus"
	default:
		return fmt.Sprintf("MEDecodeFrameStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MediaExtension/MEError-swift.struct/Code
type MEError int

const (
	// MEErrorAllocationFailure: An error code that indicates the extension can’t allocate memory.
	MEErrorAllocationFailure MEError = -19321
	// MEErrorEndOfStream: An error code that indicates the extension reached the end of the source file.
	MEErrorEndOfStream MEError = -19329
	// MEErrorInternalFailure: An error code that indicates the extension encountered an internal operation failure, such as code loading.
	MEErrorInternalFailure MEError = -19324
	// MEErrorInvalidParameter: An error code that indicates the extension received an invalid parameter.
	MEErrorInvalidParameter MEError = -19322
	// MEErrorLocationNotAvailable: An error code that indicates specific sample isn’t contiguous, spans more than one file, or is for some other reason unsuitable for reading directly from a file.
	MEErrorLocationNotAvailable MEError = -19328
	// MEErrorNoSamples: An error code that indicates there are no samples in the track or a request to load a sample buffer fails.
	MEErrorNoSamples MEError = -19327
	// MEErrorNoSuchEdit: An error code that indicates the plug-in track reader received a request to return an edit that’s out of range.
	MEErrorNoSuchEdit MEError = -19326
	// MEErrorParsingFailure: An error code that indicates the extension encountered an error while parsing the media.
	MEErrorParsingFailure MEError = -19323
	// MEErrorPermissionDenied: An error code that indicates the extension received a request to perform an invalid operation on a byte source.
	MEErrorPermissionDenied MEError = -19330
	// MEErrorPropertyNotSupported: An error code that indicates the extension encountered a property it doesn’t support reading and writing to.
	MEErrorPropertyNotSupported MEError = -19325
	// MEErrorReferenceMissing: An error code that indicates the decoder received a request to decode a sample without decoding the required reference frame dependencies first.
	MEErrorReferenceMissing MEError = -19331
	// MEErrorUnsupportedFeature: An error code that indicates the extension doesn’t support an aspect of the media.
	MEErrorUnsupportedFeature MEError = -19320
)

func (e MEError) String() string {
	switch e {
	case MEErrorAllocationFailure:
		return "MEErrorAllocationFailure"
	case MEErrorEndOfStream:
		return "MEErrorEndOfStream"
	case MEErrorInternalFailure:
		return "MEErrorInternalFailure"
	case MEErrorInvalidParameter:
		return "MEErrorInvalidParameter"
	case MEErrorLocationNotAvailable:
		return "MEErrorLocationNotAvailable"
	case MEErrorNoSamples:
		return "MEErrorNoSamples"
	case MEErrorNoSuchEdit:
		return "MEErrorNoSuchEdit"
	case MEErrorParsingFailure:
		return "MEErrorParsingFailure"
	case MEErrorPermissionDenied:
		return "MEErrorPermissionDenied"
	case MEErrorPropertyNotSupported:
		return "MEErrorPropertyNotSupported"
	case MEErrorReferenceMissing:
		return "MEErrorReferenceMissing"
	case MEErrorUnsupportedFeature:
		return "MEErrorUnsupportedFeature"
	default:
		return fmt.Sprintf("MEError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MediaExtension/MEFileInfo/FragmentsStatus-swift.enum
type MEFileInfoFragmentsStatus int

const (
	// MEFileInfoContainsFragments: The file is extendable by fragments and contains at least one fragment.
	MEFileInfoContainsFragments MEFileInfoFragmentsStatus = 1
	// MEFileInfoCouldContainButDoesNotContainFragments: The file is extendable by fragments, but doesn’t contain any fragments.
	MEFileInfoCouldContainButDoesNotContainFragments MEFileInfoFragmentsStatus = 2
	// MEFileInfoCouldNotContainFragments: The file isn’t extendable by fragments.
	MEFileInfoCouldNotContainFragments MEFileInfoFragmentsStatus = 0
)

func (e MEFileInfoFragmentsStatus) String() string {
	switch e {
	case MEFileInfoContainsFragments:
		return "MEFileInfoContainsFragments"
	case MEFileInfoCouldContainButDoesNotContainFragments:
		return "MEFileInfoCouldContainButDoesNotContainFragments"
	case MEFileInfoCouldNotContainFragments:
		return "MEFileInfoCouldNotContainFragments"
	default:
		return fmt.Sprintf("MEFileInfoFragmentsStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/MediaExtension/MEFormatReaderParseAdditionalFragmentsStatus
type MEFormatReaderParseAdditionalFragmentsStatus uint

const (
	// MEFormatReaderParseAdditionalFragmentsStatusFragmentAdded: Indicates that the format reader received one or more fragments.
	MEFormatReaderParseAdditionalFragmentsStatusFragmentAdded MEFormatReaderParseAdditionalFragmentsStatus = 2
	// MEFormatReaderParseAdditionalFragmentsStatusFragmentsComplete: Indicates that the format reader can’t receive any more fragments.
	MEFormatReaderParseAdditionalFragmentsStatusFragmentsComplete MEFormatReaderParseAdditionalFragmentsStatus = 4
	// MEFormatReaderParseAdditionalFragmentsStatusSizeIncreased: Indicates that the format reader file size increased.
	MEFormatReaderParseAdditionalFragmentsStatusSizeIncreased MEFormatReaderParseAdditionalFragmentsStatus = 1
)

func (e MEFormatReaderParseAdditionalFragmentsStatus) String() string {
	switch e {
	case MEFormatReaderParseAdditionalFragmentsStatusFragmentAdded:
		return "MEFormatReaderParseAdditionalFragmentsStatusFragmentAdded"
	case MEFormatReaderParseAdditionalFragmentsStatusFragmentsComplete:
		return "MEFormatReaderParseAdditionalFragmentsStatusFragmentsComplete"
	case MEFormatReaderParseAdditionalFragmentsStatusSizeIncreased:
		return "MEFormatReaderParseAdditionalFragmentsStatusSizeIncreased"
	default:
		return fmt.Sprintf("MEFormatReaderParseAdditionalFragmentsStatus(%d)", e)
	}
}
