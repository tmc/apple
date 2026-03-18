// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"fmt"
)

type NS int

const (
	// NSASCIIStringEncoding: Strict 7-bit ASCII encoding within 8-bit chars; ASCII values 0…127 only.
	NSASCIIStringEncoding NS = 1
	// NSArgumentEvaluationScriptError: The object specified by an argument could not be found.
	NSArgumentEvaluationScriptError NS = 3
	// NSArgumentsWrongScriptError: An argument (or more than one argument) is of the wrong type or is otherwise invalid.
	NSArgumentsWrongScriptError NS = 6
	// NSBundleErrorMaximum: The end of the range of error codes reserved for bundle errors.
	NSBundleErrorMaximum NS = 5119
	// NSBundleErrorMinimum: The start of the range of error codes reserved for bundle errors.
	NSBundleErrorMinimum NS = 4992
	// NSBundleOnDemandResourceExceededMaximumSizeError: The application exceeded the amount of on-demand resources content in use at one time.
	NSBundleOnDemandResourceExceededMaximumSizeError NS = 4993
	// NSBundleOnDemandResourceInvalidTagError: The application specified a tag that the system couldn’t find in the application tag manifest.
	NSBundleOnDemandResourceInvalidTagError NS = 4994
	NSBundleOnDemandResourceOutOfSpaceError NS = 4992
	// NSCannotCreateScriptCommandError: Could not create the script command; an invalid or unrecognized Apple event was received.
	NSCannotCreateScriptCommandError NS = 10
	// NSCloudSharingConflictError: A conflict occurred during an attempt to save changes.
	NSCloudSharingConflictError NS = 5123
	// NSCloudSharingErrorMaximum: The end of the range of error codes reserved for cloud-sharing errors.
	NSCloudSharingErrorMaximum NS = 5375
	// NSCloudSharingErrorMinimum: The start of the range of error codes reserved for cloud-sharing errors.
	NSCloudSharingErrorMinimum NS = 5120
	// NSCloudSharingNetworkFailureError: Sharing failed due to a network failure.
	NSCloudSharingNetworkFailureError NS = 5120
	// NSCloudSharingNoPermissionError: The current user doesn’t have permission to perform the requested actions.
	NSCloudSharingNoPermissionError NS = 5124
	// NSCloudSharingOtherError: An otherwise unspecified cloud-sharing error occurred.
	NSCloudSharingOtherError NS = 5375
	// NSCloudSharingQuotaExceededError: The user doesn’t have enough storage space available to share the requested items.
	NSCloudSharingQuotaExceededError NS = 5121
	// NSCloudSharingTooManyParticipantsError: Additional participants couldn’t be added to the share, because the limit was reached.
	NSCloudSharingTooManyParticipantsError NS = 5122
	// NSCoderErrorMaximum: The end of the range of error codes reserved for coder errors.
	NSCoderErrorMaximum NS = 4991
	// NSCoderErrorMinimum: The start of the range of error codes reserved for coder errors.
	NSCoderErrorMinimum NS = 4864
	// NSCoderInvalidValueError: Data wasn’t valid to encode.
	NSCoderInvalidValueError NS = 4866
	// NSCoderReadCorruptError: Decoding failed due to corrupt data.
	NSCoderReadCorruptError NS = 4864
	// NSCoderValueNotFoundError: The requested data wasn’t found.
	NSCoderValueNotFoundError NS = 4865
	// NSCollectorDisabledOption: Specifies that the block is retained, and therefore ineligible for collection.
	NSCollectorDisabledOption NS = 2
	NSCompressionErrorMaximum NS = 5503
	// NSCompressionErrorMinimum: The start of the range of error codes reserved for compression errors.
	NSCompressionErrorMinimum NS = 5376
	// NSCompressionFailedError: An error code value that indicates a failure to compress data using the provided algorithm.
	NSCompressionFailedError NS = 5376
	// NSContainerSpecifierError: Error evaluating container specifier.
	NSContainerSpecifierError NS = 2
	// NSDateComponentUndefined: Specifies a date component without a value.
	NSDateComponentUndefined NS = 9223372036854775807
	// NSDecompressionFailedError: An error code value that indicates a failure to decompress data using the provided algorithm.
	NSDecompressionFailedError NS = 5377
	// NSExecutableArchitectureMismatchError: The executable doesn’t provide an architecture compatible with the current process.
	NSExecutableArchitectureMismatchError NS = 3585
	// NSExecutableErrorMaximum: The end of the range of error codes reserved for errors related to executable files.
	NSExecutableErrorMaximum NS = 3839
	// NSExecutableErrorMinimum: The beginning of the range of error codes reserved for errors related to executable files.
	NSExecutableErrorMinimum NS = 3584
	// NSExecutableLinkError: The executable failed due to linking issues.
	NSExecutableLinkError NS = 3588
	// NSExecutableLoadError: Executable cannot be loaded for an otherwise-unspecified reason.
	NSExecutableLoadError NS = 3587
	// NSExecutableNotLoadableError: The executable type isn’t loadable in the current process.
	NSExecutableNotLoadableError NS = 3584
	// NSExecutableRuntimeMismatchError: The executable has Objective-C runtime information that’s incompatible with the current process.
	NSExecutableRuntimeMismatchError NS = 3586
	// NSFeatureUnsupportedError: The feature isn’t supported, because the file system lacks the feature, or required libraries are missing, or other similar reasons.
	NSFeatureUnsupportedError NS = 3328
	// NSFileErrorMaximum: The end of the range of error codes reserved for file errors.
	NSFileErrorMaximum NS = 1023
	// NSFileErrorMinimum: The start of the range of error codes reserved for file errors.
	NSFileErrorMinimum NS = 0
	// NSFileLockingError: The file could not be locked.
	NSFileLockingError NS = 255
	// NSFileManagerUnmountBusyError: The volume couldn’t be unmounted because it’s in use.
	NSFileManagerUnmountBusyError NS = 769
	// NSFileManagerUnmountUnknownError: The volume couldn’t be unmounted, for unknown reasons.
	NSFileManagerUnmountUnknownError NS = 768
	// NSFileNoSuchFileError: A filesystem operation was attempted on a non-existent file.
	NSFileNoSuchFileError NS = 4
	// NSFileReadCorruptFileError: Could not read because of a corrupted file, bad format, or similar reason.
	NSFileReadCorruptFileError NS = 259
	// NSFileReadInapplicableStringEncodingError: Could not read because the string encoding wasn’t applicable.
	NSFileReadInapplicableStringEncodingError NS = 261
	// NSFileReadInvalidFileNameError: Could not read because of an invalid file name.
	NSFileReadInvalidFileNameError NS = 258
	// NSFileReadNoPermissionError: Could not read because of a permission problem.
	NSFileReadNoPermissionError NS = 257
	// NSFileReadNoSuchFileError: Could not read because no such file was found.
	NSFileReadNoSuchFileError NS = 260
	// NSFileReadTooLargeError: Could not read because the specified file was too large.
	NSFileReadTooLargeError NS = 263
	// NSFileReadUnknownError: Could not read, for unknown reasons.
	NSFileReadUnknownError NS = 256
	// NSFileReadUnknownStringEncodingError: Could not read because the string coding of the file couldn’t be determined.
	NSFileReadUnknownStringEncodingError NS = 264
	// NSFileReadUnsupportedSchemeError: Could not read because the specified URL scheme is unsupported.
	NSFileReadUnsupportedSchemeError NS = 262
	// NSFileWriteFileExistsError: Could not perform an operation because the destination file already exists.
	NSFileWriteFileExistsError NS = 516
	// NSFileWriteInapplicableStringEncodingError: Could not write because the string encoding was not applicable.
	NSFileWriteInapplicableStringEncodingError NS = 517
	// NSFileWriteInvalidFileNameError: Could not write because of an invalid file name.
	NSFileWriteInvalidFileNameError NS = 514
	// NSFileWriteNoPermissionError: Could not write because of a permission problem.
	NSFileWriteNoPermissionError NS = 513
	// NSFileWriteOutOfSpaceError: Could not write because of a lack of disk space.
	NSFileWriteOutOfSpaceError NS = 640
	// NSFileWriteUnknownError: Could not write, for unknown reasons.
	NSFileWriteUnknownError NS = 512
	// NSFileWriteUnsupportedSchemeError: Could not write because the specified URL scheme is unsupported.
	NSFileWriteUnsupportedSchemeError NS = 518
	// NSFileWriteVolumeReadOnlyError: Could not write because the volume is read-only.
	NSFileWriteVolumeReadOnlyError NS = 642
	// NSFormattingError: A formatter couldn’t generate a string for an object, or parse a string into an object.
	NSFormattingError NS = 2048
	// NSFormattingErrorMaximum: The end of the range of error codes reserved for formatting errors.
	NSFormattingErrorMaximum NS = 2559
	// NSFormattingErrorMinimum: The start of the range of error codes reserved for formatting errors.
	NSFormattingErrorMinimum NS = 2048
	// NSHPUXOperatingSystem: Indicates the HP UX operating system.
	NSHPUXOperatingSystem NS = 4
	// NSISO2022JPStringEncoding: ISO 2022 Japanese encoding for email.
	NSISO2022JPStringEncoding NS = 21
	// NSISOLatin1StringEncoding: 8-bit ISO Latin 1 encoding.
	NSISOLatin1StringEncoding NS = 5
	// NSISOLatin2StringEncoding: 8-bit ISO Latin 2 encoding.
	NSISOLatin2StringEncoding NS = 9
	// NSInternalScriptError: An unidentified internal error occurred; indicates an error in the scripting support of your application.
	NSInternalScriptError NS = 8
	// NSInternalSpecifierError: Other internal error.
	NSInternalSpecifierError NS = 5
	// NSInvalidIndexSpecifierError: Index out of bounds.
	NSInvalidIndexSpecifierError NS = 4
	// NSJapaneseEUCStringEncoding: 8-bit EUC encoding for Japanese text.
	NSJapaneseEUCStringEncoding NS = 3
	// NSKeySpecifierEvaluationScriptError: The object or objects specified by a key (for commands that support key specifiers) could not be found.
	NSKeySpecifierEvaluationScriptError NS = 2
	// NSKeyValueValidationError: A key-value coding validation error.
	NSKeyValueValidationError NS = 1024
	// NSMACHOperatingSystem: Indicates the macOS operating system.
	NSMACHOperatingSystem NS = 5
	// NSMacOSRomanStringEncoding: Classic Macintosh Roman encoding.
	NSMacOSRomanStringEncoding NS = 30
	// NSNEXTSTEPStringEncoding: 8-bit ASCII encoding with NEXTSTEP extensions.
	NSNEXTSTEPStringEncoding NS = 2
	// NSNoScriptError: No error.
	NSNoScriptError NS = 0
	// NSNoSpecifierError: No error encountered.
	NSNoSpecifierError NS = 0
	// NSNoTopLevelContainersSpecifierError: Someone called `evaluate` with `nil`.
	NSNoTopLevelContainersSpecifierError NS = 1
	// NSNonLossyASCIIStringEncoding: 7-bit verbose ASCII to represent all Unicode characters.
	NSNonLossyASCIIStringEncoding NS = 7
	// NSOSF1OperatingSystem: Indicates the OSF/1 operating system.
	NSOSF1OperatingSystem NS = 7
	// NSOperationNotSupportedForKeyScriptError: The implementation of a scripting command signaled an error.
	NSOperationNotSupportedForKeyScriptError NS = 9
	// NSOperationNotSupportedForKeySpecifierError: Attempt made to perform an unsupported operation on some key.
	NSOperationNotSupportedForKeySpecifierError NS = 6
	// NSPropertyListErrorMaximum: The end of the range of error codes reserved for property list errors.
	NSPropertyListErrorMaximum NS = 4095
	// NSPropertyListErrorMinimum: The start of the range of error codes reserved for property list errors.
	NSPropertyListErrorMinimum NS = 3840
	// NSPropertyListReadCorruptError: Parsing of the property list failed.
	NSPropertyListReadCorruptError NS = 3840
	// NSPropertyListReadStreamError: Reading of the property list failed.
	NSPropertyListReadStreamError NS = 3842
	// NSPropertyListReadUnknownVersionError: The version number of the property list cannot be determined.
	NSPropertyListReadUnknownVersionError NS = 3841
	// NSPropertyListWriteInvalidError: Writing failed because of an invalid property list object, or an invalid property list type was specified.
	NSPropertyListWriteInvalidError NS = 3852
	// NSPropertyListWriteStreamError: Writing to the property list failed.
	NSPropertyListWriteStreamError NS = 3851
	// NSReceiverEvaluationScriptError: The object or objects specified by the direct parameter to a command could not be found.
	NSReceiverEvaluationScriptError NS = 1
	// NSReceiversCantHandleCommandScriptError: The receivers don’t support the command sent to them.
	NSReceiversCantHandleCommandScriptError NS = 4
	// NSRequiredArgumentsMissingScriptError: An argument (or more than one argument) is missing.
	NSRequiredArgumentsMissingScriptError NS = 5
	// NSScannedOption: Specifies allocation of scanned memory.
	NSScannedOption NS = 1
	// NSShiftJISStringEncoding: 8-bit Shift-JIS encoding for Japanese text.
	NSShiftJISStringEncoding NS = 8
	// NSSolarisOperatingSystem: Indicates the Solaris operating system.
	NSSolarisOperatingSystem NS = 3
	// NSSunOSOperatingSystem: Indicates the Sun OS operating system.
	NSSunOSOperatingSystem NS = 6
	// NSSymbolStringEncoding: 8-bit Adobe Symbol encoding vector.
	NSSymbolStringEncoding NS = 6
	// NSUTF16BigEndianStringEncoding: [NSUTF16StringEncoding] encoding with explicit endianness specified.
	NSUTF16BigEndianStringEncoding NS = 2415919360
	// NSUTF16LittleEndianStringEncoding: [NSUTF16StringEncoding] encoding with explicit endianness specified.
	NSUTF16LittleEndianStringEncoding NS = 2483028224
	// NSUTF16StringEncoding: # Discussion
	NSUTF16StringEncoding NS = 10
	// NSUTF32BigEndianStringEncoding: [NSUTF32StringEncoding] encoding with explicit endianness specified.
	NSUTF32BigEndianStringEncoding NS = 2550137088
	// NSUTF32LittleEndianStringEncoding: [NSUTF32StringEncoding] encoding with explicit endianness specified.
	NSUTF32LittleEndianStringEncoding NS = 2617245952
	// NSUTF32StringEncoding: 32-bit UTF encoding.
	NSUTF32StringEncoding NS = 2348810496
	// NSUTF8StringEncoding: An 8-bit representation of Unicode characters, suitable for transmission or storage by ASCII-based systems.
	NSUTF8StringEncoding NS = 4
	// NSUbiquitousFileErrorMaximum: The maximum error code value that represents an iCloud error.
	NSUbiquitousFileErrorMaximum NS = 4607
	// NSUbiquitousFileErrorMinimum: The minimum error code value that represents an iCloud error.
	NSUbiquitousFileErrorMinimum NS = 4352
	// NSUbiquitousFileNotUploadedDueToQuotaError: The item could not be uploaded to iCloud because it would make the account go over its quota.
	NSUbiquitousFileNotUploadedDueToQuotaError NS = 4354
	// NSUbiquitousFileUbiquityServerNotAvailable: A failure to connect to the iCloud servers.
	NSUbiquitousFileUbiquityServerNotAvailable NS = 4355
	// NSUbiquitousFileUnavailableError: The item has not been uploaded to iCloud by another device yet.
	NSUbiquitousFileUnavailableError NS = 4353
	// NSUnicodeStringEncoding: The canonical Unicode encoding for string objects.
	NSUnicodeStringEncoding NS = 10
	// NSUnknownKeyScriptError: An unidentified error occurred; indicates an error in the scripting support of your application.
	NSUnknownKeyScriptError NS = 7
	// NSUnknownKeySpecifierError: Receivers do not understand the key.
	NSUnknownKeySpecifierError NS = 3
	// NSUserActivityConnectionUnavailableError: The user activity couldn’t be continued because a required connection wasn’t available.
	NSUserActivityConnectionUnavailableError NS = 4609
	// NSUserActivityErrorMaximum: The end of the range of error codes reserved for user activity errors.
	NSUserActivityErrorMaximum NS = 4863
	// NSUserActivityErrorMinimum: The start of the range of error codes reserved for user activity errors.
	NSUserActivityErrorMinimum NS = 4608
	// NSUserActivityHandoffFailedError: The data for the user activity wasn’t available.
	NSUserActivityHandoffFailedError NS = 4608
	// NSUserActivityHandoffUserInfoTooLargeError: The user info dictionary was too large to receive.
	NSUserActivityHandoffUserInfoTooLargeError NS = 4611
	// NSUserActivityRemoteApplicationTimedOutError: The remote application failed to send data within the specified time.
	NSUserActivityRemoteApplicationTimedOutError NS = 4610
	// NSUserCancelledError: The user canceled the operation (for example, by pressing Command-period).
	NSUserCancelledError NS = 3072
	// NSValidationErrorMaximum: The end of the range of error codes reserved for validation errors.
	NSValidationErrorMaximum NS = 2047
	// NSValidationErrorMinimum: The start of the range of error codes reserved for validation errors.
	NSValidationErrorMinimum NS = 1024
	// NSWindows95OperatingSystem: Indicates the Windows 95 operating system.
	NSWindows95OperatingSystem NS = 2
	// NSWindowsCP1250StringEncoding: Microsoft Windows codepage 1250; equivalent to WinLatin2.
	NSWindowsCP1250StringEncoding NS = 15
	// NSWindowsCP1251StringEncoding: Microsoft Windows codepage 1251, encoding Cyrillic characters; equivalent to AdobeStandardCyrillic font encoding.
	NSWindowsCP1251StringEncoding NS = 11
	// NSWindowsCP1252StringEncoding: Microsoft Windows codepage 1252; equivalent to WinLatin1.
	NSWindowsCP1252StringEncoding NS = 12
	// NSWindowsCP1253StringEncoding: Microsoft Windows codepage 1253, encoding Greek characters.
	NSWindowsCP1253StringEncoding NS = 13
	// NSWindowsCP1254StringEncoding: Microsoft Windows codepage 1254, encoding Turkish characters.
	NSWindowsCP1254StringEncoding NS = 14
	// NSWindowsNTOperatingSystem: Indicates the Windows NT operating system.
	NSWindowsNTOperatingSystem NS = 1
	// NSXPCConnectionCodeSigningRequirementFailure: A code-signing requirement check failed.
	NSXPCConnectionCodeSigningRequirementFailure NS = 4102
	// NSXPCConnectionErrorMaximum: The upper bounds of XPC connection error code values.
	NSXPCConnectionErrorMaximum NS = 4224
	// NSXPCConnectionErrorMinimum: The lower bounds of XPC connection error code values.
	NSXPCConnectionErrorMinimum NS = 4096
	// NSXPCConnectionInterrupted: The XPC connection was interrupted.
	NSXPCConnectionInterrupted NS = 4097
	// NSXPCConnectionInvalid: The XPC connection was invalid.
	NSXPCConnectionInvalid NS = 4099
	// NSXPCConnectionReplyInvalid: The XPC connection reply was invalid.
	NSXPCConnectionReplyInvalid NS = 4101
	// NS_BigEndian: The byte order is big endian.
	NS_BigEndian NS = 0
	// NS_LittleEndian: The byte order is little endian.
	NS_LittleEndian NS = 1
	// NS_UnknownByteOrder: The byte order is unknown.
	NS_UnknownByteOrder NS = 0
	// Deprecated.
	NSUndefinedDateComponent NS = -9223372036854775808
)

func (e NS) String() string {
	switch e {
	case NSASCIIStringEncoding:
		return "NSASCIIStringEncoding"
	case NSArgumentEvaluationScriptError:
		return "NSArgumentEvaluationScriptError"
	case NSArgumentsWrongScriptError:
		return "NSArgumentsWrongScriptError"
	case NSBundleErrorMaximum:
		return "NSBundleErrorMaximum"
	case NSBundleErrorMinimum:
		return "NSBundleErrorMinimum"
	case NSBundleOnDemandResourceExceededMaximumSizeError:
		return "NSBundleOnDemandResourceExceededMaximumSizeError"
	case NSBundleOnDemandResourceInvalidTagError:
		return "NSBundleOnDemandResourceInvalidTagError"
	case NSCannotCreateScriptCommandError:
		return "NSCannotCreateScriptCommandError"
	case NSCloudSharingConflictError:
		return "NSCloudSharingConflictError"
	case NSCloudSharingErrorMaximum:
		return "NSCloudSharingErrorMaximum"
	case NSCloudSharingErrorMinimum:
		return "NSCloudSharingErrorMinimum"
	case NSCloudSharingNoPermissionError:
		return "NSCloudSharingNoPermissionError"
	case NSCloudSharingQuotaExceededError:
		return "NSCloudSharingQuotaExceededError"
	case NSCloudSharingTooManyParticipantsError:
		return "NSCloudSharingTooManyParticipantsError"
	case NSCoderErrorMaximum:
		return "NSCoderErrorMaximum"
	case NSCoderErrorMinimum:
		return "NSCoderErrorMinimum"
	case NSCoderInvalidValueError:
		return "NSCoderInvalidValueError"
	case NSCoderValueNotFoundError:
		return "NSCoderValueNotFoundError"
	case NSCollectorDisabledOption:
		return "NSCollectorDisabledOption"
	case NSCompressionErrorMaximum:
		return "NSCompressionErrorMaximum"
	case NSCompressionErrorMinimum:
		return "NSCompressionErrorMinimum"
	case NSDateComponentUndefined:
		return "NSDateComponentUndefined"
	case NSDecompressionFailedError:
		return "NSDecompressionFailedError"
	case NSExecutableArchitectureMismatchError:
		return "NSExecutableArchitectureMismatchError"
	case NSExecutableErrorMaximum:
		return "NSExecutableErrorMaximum"
	case NSExecutableErrorMinimum:
		return "NSExecutableErrorMinimum"
	case NSExecutableLinkError:
		return "NSExecutableLinkError"
	case NSExecutableLoadError:
		return "NSExecutableLoadError"
	case NSExecutableRuntimeMismatchError:
		return "NSExecutableRuntimeMismatchError"
	case NSFeatureUnsupportedError:
		return "NSFeatureUnsupportedError"
	case NSFileErrorMaximum:
		return "NSFileErrorMaximum"
	case NSFileErrorMinimum:
		return "NSFileErrorMinimum"
	case NSFileLockingError:
		return "NSFileLockingError"
	case NSFileManagerUnmountBusyError:
		return "NSFileManagerUnmountBusyError"
	case NSFileManagerUnmountUnknownError:
		return "NSFileManagerUnmountUnknownError"
	case NSFileNoSuchFileError:
		return "NSFileNoSuchFileError"
	case NSFileReadCorruptFileError:
		return "NSFileReadCorruptFileError"
	case NSFileReadInapplicableStringEncodingError:
		return "NSFileReadInapplicableStringEncodingError"
	case NSFileReadInvalidFileNameError:
		return "NSFileReadInvalidFileNameError"
	case NSFileReadNoPermissionError:
		return "NSFileReadNoPermissionError"
	case NSFileReadNoSuchFileError:
		return "NSFileReadNoSuchFileError"
	case NSFileReadTooLargeError:
		return "NSFileReadTooLargeError"
	case NSFileReadUnknownError:
		return "NSFileReadUnknownError"
	case NSFileReadUnknownStringEncodingError:
		return "NSFileReadUnknownStringEncodingError"
	case NSFileReadUnsupportedSchemeError:
		return "NSFileReadUnsupportedSchemeError"
	case NSFileWriteFileExistsError:
		return "NSFileWriteFileExistsError"
	case NSFileWriteInapplicableStringEncodingError:
		return "NSFileWriteInapplicableStringEncodingError"
	case NSFileWriteInvalidFileNameError:
		return "NSFileWriteInvalidFileNameError"
	case NSFileWriteNoPermissionError:
		return "NSFileWriteNoPermissionError"
	case NSFileWriteOutOfSpaceError:
		return "NSFileWriteOutOfSpaceError"
	case NSFileWriteUnknownError:
		return "NSFileWriteUnknownError"
	case NSFileWriteUnsupportedSchemeError:
		return "NSFileWriteUnsupportedSchemeError"
	case NSFileWriteVolumeReadOnlyError:
		return "NSFileWriteVolumeReadOnlyError"
	case NSFormattingError:
		return "NSFormattingError"
	case NSFormattingErrorMaximum:
		return "NSFormattingErrorMaximum"
	case NSISO2022JPStringEncoding:
		return "NSISO2022JPStringEncoding"
	case NSISOLatin1StringEncoding:
		return "NSISOLatin1StringEncoding"
	case NSISOLatin2StringEncoding:
		return "NSISOLatin2StringEncoding"
	case NSInternalScriptError:
		return "NSInternalScriptError"
	case NSKeyValueValidationError:
		return "NSKeyValueValidationError"
	case NSMacOSRomanStringEncoding:
		return "NSMacOSRomanStringEncoding"
	case NSNonLossyASCIIStringEncoding:
		return "NSNonLossyASCIIStringEncoding"
	case NSPropertyListErrorMaximum:
		return "NSPropertyListErrorMaximum"
	case NSPropertyListErrorMinimum:
		return "NSPropertyListErrorMinimum"
	case NSPropertyListReadStreamError:
		return "NSPropertyListReadStreamError"
	case NSPropertyListReadUnknownVersionError:
		return "NSPropertyListReadUnknownVersionError"
	case NSPropertyListWriteInvalidError:
		return "NSPropertyListWriteInvalidError"
	case NSPropertyListWriteStreamError:
		return "NSPropertyListWriteStreamError"
	case NSUTF16BigEndianStringEncoding:
		return "NSUTF16BigEndianStringEncoding"
	case NSUTF16LittleEndianStringEncoding:
		return "NSUTF16LittleEndianStringEncoding"
	case NSUTF32BigEndianStringEncoding:
		return "NSUTF32BigEndianStringEncoding"
	case NSUTF32LittleEndianStringEncoding:
		return "NSUTF32LittleEndianStringEncoding"
	case NSUTF32StringEncoding:
		return "NSUTF32StringEncoding"
	case NSUbiquitousFileErrorMaximum:
		return "NSUbiquitousFileErrorMaximum"
	case NSUbiquitousFileErrorMinimum:
		return "NSUbiquitousFileErrorMinimum"
	case NSUbiquitousFileNotUploadedDueToQuotaError:
		return "NSUbiquitousFileNotUploadedDueToQuotaError"
	case NSUbiquitousFileUbiquityServerNotAvailable:
		return "NSUbiquitousFileUbiquityServerNotAvailable"
	case NSUbiquitousFileUnavailableError:
		return "NSUbiquitousFileUnavailableError"
	case NSUserActivityConnectionUnavailableError:
		return "NSUserActivityConnectionUnavailableError"
	case NSUserActivityErrorMaximum:
		return "NSUserActivityErrorMaximum"
	case NSUserActivityErrorMinimum:
		return "NSUserActivityErrorMinimum"
	case NSUserActivityHandoffUserInfoTooLargeError:
		return "NSUserActivityHandoffUserInfoTooLargeError"
	case NSUserActivityRemoteApplicationTimedOutError:
		return "NSUserActivityRemoteApplicationTimedOutError"
	case NSUserCancelledError:
		return "NSUserCancelledError"
	case NSValidationErrorMaximum:
		return "NSValidationErrorMaximum"
	case NSWindowsCP1250StringEncoding:
		return "NSWindowsCP1250StringEncoding"
	case NSWindowsCP1251StringEncoding:
		return "NSWindowsCP1251StringEncoding"
	case NSWindowsCP1252StringEncoding:
		return "NSWindowsCP1252StringEncoding"
	case NSWindowsCP1253StringEncoding:
		return "NSWindowsCP1253StringEncoding"
	case NSWindowsCP1254StringEncoding:
		return "NSWindowsCP1254StringEncoding"
	case NSXPCConnectionCodeSigningRequirementFailure:
		return "NSXPCConnectionCodeSigningRequirementFailure"
	case NSXPCConnectionErrorMaximum:
		return "NSXPCConnectionErrorMaximum"
	case NSXPCConnectionErrorMinimum:
		return "NSXPCConnectionErrorMinimum"
	case NSXPCConnectionInterrupted:
		return "NSXPCConnectionInterrupted"
	case NSXPCConnectionInvalid:
		return "NSXPCConnectionInvalid"
	case NSXPCConnectionReplyInvalid:
		return "NSXPCConnectionReplyInvalid"
	case NSUndefinedDateComponent:
		return "NSUndefinedDateComponent"
	default:
		return fmt.Sprintf("NS(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/ActivityOptions
type NSActivityOptions int

const (
	// NSActivityAnimationTrackingEnabled: A flag to track the activity with an animation signpost interval.
	NSActivityAnimationTrackingEnabled NSActivityOptions = 32769
	// NSActivityAutomaticTerminationDisabled: A flag to prevent automatic termination.
	NSActivityAutomaticTerminationDisabled NSActivityOptions = 32768
	// NSActivityBackground: A flag to indicate the app has initiated some kind of work, but not as the direct result of user request.
	NSActivityBackground NSActivityOptions = 255
	// NSActivityIdleDisplaySleepDisabled: A flag to require the screen to stay powered on.
	NSActivityIdleDisplaySleepDisabled NSActivityOptions = 1099511627776
	// NSActivityIdleSystemSleepDisabled: A flag to prevent idle sleep.
	NSActivityIdleSystemSleepDisabled NSActivityOptions = 1048576
	// NSActivityLatencyCritical: A flag to indicate the activity requires the highest amount of timer and I/O precision available.
	NSActivityLatencyCritical NSActivityOptions = 1095216660480
	// NSActivitySuddenTerminationDisabled: A flag to prevent sudden termination.
	NSActivitySuddenTerminationDisabled NSActivityOptions = 16384
	// NSActivityTrackingEnabled: A flag to track the activity with a signpost interval.
	NSActivityTrackingEnabled NSActivityOptions = 70368744177664
	// NSActivityUserInitiated: A flag to indicate the app is performing a user-requested action.
	NSActivityUserInitiated NSActivityOptions = 0
	// NSActivityUserInitiatedAllowingIdleSystemSleep: A flag to indicate the app is performing a user-requested action, but that the system can sleep on idle.
	NSActivityUserInitiatedAllowingIdleSystemSleep NSActivityOptions = 0
	// NSActivityUserInteractive: A flag to indicate the app is responding to user interaction.
	NSActivityUserInteractive NSActivityOptions = 1095216660481
)

func (e NSActivityOptions) String() string {
	switch e {
	case NSActivityAnimationTrackingEnabled:
		return "NSActivityAnimationTrackingEnabled"
	case NSActivityAutomaticTerminationDisabled:
		return "NSActivityAutomaticTerminationDisabled"
	case NSActivityBackground:
		return "NSActivityBackground"
	case NSActivityIdleDisplaySleepDisabled:
		return "NSActivityIdleDisplaySleepDisabled"
	case NSActivityIdleSystemSleepDisabled:
		return "NSActivityIdleSystemSleepDisabled"
	case NSActivityLatencyCritical:
		return "NSActivityLatencyCritical"
	case NSActivitySuddenTerminationDisabled:
		return "NSActivitySuddenTerminationDisabled"
	case NSActivityTrackingEnabled:
		return "NSActivityTrackingEnabled"
	case NSActivityUserInitiated:
		return "NSActivityUserInitiated"
	case NSActivityUserInteractive:
		return "NSActivityUserInteractive"
	default:
		return fmt.Sprintf("NSActivityOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/AlignmentOptions
type AlignmentOptions int

const (
	// NSAlignAllEdgesInward: Aligns all edges inward.
	NSAlignAllEdgesInward AlignmentOptions = 0
	// NSAlignAllEdgesNearest: Aligns all edges to the nearest value.
	NSAlignAllEdgesNearest AlignmentOptions = 0
	// NSAlignAllEdgesOutward: Aligns all edges outwards.
	NSAlignAllEdgesOutward AlignmentOptions = 0
	// NSAlignHeightInward: Specifies that alignment of the height should be to the nearest inward integral value.
	NSAlignHeightInward AlignmentOptions = 32
	// NSAlignHeightNearest: Specifies that alignment of the height should be to the nearest integral value.
	NSAlignHeightNearest AlignmentOptions = 2097152
	// NSAlignHeightOutward: Specifies that alignment of the height should be to the nearest outward integral value.
	NSAlignHeightOutward AlignmentOptions = 8192
	// NSAlignMaxXInward: Specifies that alignment of the maximum X coordinate should be to the nearest inward integral value.
	NSAlignMaxXInward AlignmentOptions = 4
	// NSAlignMaxXNearest: Specifies that alignment of the maximum X coordinate should be to the nearest integral value.
	NSAlignMaxXNearest AlignmentOptions = 262144
	// NSAlignMaxXOutward: Specifies that alignment of the maximum X coordinate should be to the nearest outward integral value.
	NSAlignMaxXOutward AlignmentOptions = 1024
	// NSAlignMaxYInward: Specifies that alignment of the maximum X coordinate should be to the nearest inward integral value.
	NSAlignMaxYInward AlignmentOptions = 8
	// NSAlignMaxYNearest: Specifies that alignment of the maximum Y coordinate should be to the nearest integral value.
	NSAlignMaxYNearest AlignmentOptions = 524288
	// NSAlignMaxYOutward: Specifies that alignment of the maximum Y coordinate should be to the nearest outward integral value.
	NSAlignMaxYOutward AlignmentOptions = 2048
	// NSAlignMinXInward: Specifies that alignment of the minimum X coordinate should be to the nearest inward integral value.
	NSAlignMinXInward AlignmentOptions = 1
	// NSAlignMinXNearest: Specifies that alignment of the minimum X coordinate should be to the nearest integral value.
	NSAlignMinXNearest AlignmentOptions = 65536
	// NSAlignMinXOutward: Specifies that alignment of the minimum X coordinate should be to the nearest outward integral value.
	NSAlignMinXOutward AlignmentOptions = 256
	// NSAlignMinYInward: Specifies that alignment of the minimum Y coordinate should be to the nearest inward integral value.
	NSAlignMinYInward AlignmentOptions = 2
	// NSAlignMinYNearest: Specifies that alignment of the minimum Y coordinate should be to the nearest integral value.
	NSAlignMinYNearest AlignmentOptions = 131072
	// NSAlignMinYOutward: Specifies that alignment of the minimum Y coordinate should be to the nearest outward integral value.
	NSAlignMinYOutward AlignmentOptions = 512
	// NSAlignRectFlipped: This option should be included  if the rectangle is in a flipped coordinate system.
	NSAlignRectFlipped AlignmentOptions = -9223372036854775808
	// NSAlignWidthInward: Specifies that alignment of the width should be to the nearest inward integral value.
	NSAlignWidthInward AlignmentOptions = 16
	// NSAlignWidthNearest: Specifies that alignment of the width should be to the nearest integral value.
	NSAlignWidthNearest AlignmentOptions = 1048576
	// NSAlignWidthOutward: Specifies that alignment of the width should be to the nearest outward integral value.
	NSAlignWidthOutward AlignmentOptions = 4096
)

func (e AlignmentOptions) String() string {
	switch e {
	case NSAlignAllEdgesInward:
		return "NSAlignAllEdgesInward"
	case NSAlignHeightInward:
		return "NSAlignHeightInward"
	case NSAlignHeightNearest:
		return "NSAlignHeightNearest"
	case NSAlignHeightOutward:
		return "NSAlignHeightOutward"
	case NSAlignMaxXInward:
		return "NSAlignMaxXInward"
	case NSAlignMaxXNearest:
		return "NSAlignMaxXNearest"
	case NSAlignMaxXOutward:
		return "NSAlignMaxXOutward"
	case NSAlignMaxYInward:
		return "NSAlignMaxYInward"
	case NSAlignMaxYNearest:
		return "NSAlignMaxYNearest"
	case NSAlignMaxYOutward:
		return "NSAlignMaxYOutward"
	case NSAlignMinXInward:
		return "NSAlignMinXInward"
	case NSAlignMinXNearest:
		return "NSAlignMinXNearest"
	case NSAlignMinXOutward:
		return "NSAlignMinXOutward"
	case NSAlignMinYInward:
		return "NSAlignMinYInward"
	case NSAlignMinYNearest:
		return "NSAlignMinYNearest"
	case NSAlignMinYOutward:
		return "NSAlignMinYOutward"
	case NSAlignRectFlipped:
		return "NSAlignRectFlipped"
	case NSAlignWidthInward:
		return "NSAlignWidthInward"
	case NSAlignWidthNearest:
		return "NSAlignWidthNearest"
	case NSAlignWidthOutward:
		return "NSAlignWidthOutward"
	default:
		return fmt.Sprintf("AlignmentOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSAppleEventDescriptor/SendOptions
type NSAppleEventSendOptions int

const (
	NSAppleEventSendAlwaysInteract NSAppleEventSendOptions = 48
	NSAppleEventSendCanInteract NSAppleEventSendOptions = 32
	NSAppleEventSendCanSwitchLayer NSAppleEventSendOptions = 64
	NSAppleEventSendDefaultOptions NSAppleEventSendOptions = 3
	NSAppleEventSendDontAnnotate NSAppleEventSendOptions = 65536
	NSAppleEventSendDontExecute NSAppleEventSendOptions = 8192
	NSAppleEventSendDontRecord NSAppleEventSendOptions = 4096
	NSAppleEventSendNeverInteract NSAppleEventSendOptions = 16
	NSAppleEventSendNoReply NSAppleEventSendOptions = 1
	NSAppleEventSendQueueReply NSAppleEventSendOptions = 2
	NSAppleEventSendWaitForReply NSAppleEventSendOptions = 3
)

func (e NSAppleEventSendOptions) String() string {
	switch e {
	case NSAppleEventSendAlwaysInteract:
		return "NSAppleEventSendAlwaysInteract"
	case NSAppleEventSendCanInteract:
		return "NSAppleEventSendCanInteract"
	case NSAppleEventSendCanSwitchLayer:
		return "NSAppleEventSendCanSwitchLayer"
	case NSAppleEventSendDefaultOptions:
		return "NSAppleEventSendDefaultOptions"
	case NSAppleEventSendDontAnnotate:
		return "NSAppleEventSendDontAnnotate"
	case NSAppleEventSendDontExecute:
		return "NSAppleEventSendDontExecute"
	case NSAppleEventSendDontRecord:
		return "NSAppleEventSendDontRecord"
	case NSAppleEventSendNeverInteract:
		return "NSAppleEventSendNeverInteract"
	case NSAppleEventSendNoReply:
		return "NSAppleEventSendNoReply"
	case NSAppleEventSendQueueReply:
		return "NSAppleEventSendQueueReply"
	default:
		return fmt.Sprintf("NSAppleEventSendOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/EnumerationOptions
type NSAttributedStringEnumerationOptions int

const (
	// NSAttributedStringEnumerationLongestEffectiveRangeNotRequired: If [NSAttributedStringEnumerationLongestEffectiveRangeNotRequired] option is supplied, then the longest effective range computation is not performed; the blocks may be invoked with consecutive attribute runs that have the same value.
	NSAttributedStringEnumerationLongestEffectiveRangeNotRequired NSAttributedStringEnumerationOptions = 1048576
	// NSAttributedStringEnumerationReverse: Causes the enumeration to occur in reverse.
	NSAttributedStringEnumerationReverse NSAttributedStringEnumerationOptions = 2
)

func (e NSAttributedStringEnumerationOptions) String() string {
	switch e {
	case NSAttributedStringEnumerationLongestEffectiveRangeNotRequired:
		return "NSAttributedStringEnumerationLongestEffectiveRangeNotRequired"
	case NSAttributedStringEnumerationReverse:
		return "NSAttributedStringEnumerationReverse"
	default:
		return fmt.Sprintf("NSAttributedStringEnumerationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringFormattingOptions
type NSAttributedStringFormattingOptions uint

const (
	// NSAttributedStringFormattingApplyReplacementIndexAttribute: An option to apply to the replaced portions of text in a format string.
	NSAttributedStringFormattingApplyReplacementIndexAttribute NSAttributedStringFormattingOptions = 1
	// NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging: An option to replace the attributes in a substituted string with those of the provided attributed string.
	NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging NSAttributedStringFormattingOptions = 0
)

func (e NSAttributedStringFormattingOptions) String() string {
	switch e {
	case NSAttributedStringFormattingApplyReplacementIndexAttribute:
		return "NSAttributedStringFormattingApplyReplacementIndexAttribute"
	case NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging:
		return "NSAttributedStringFormattingInsertArgumentAttributesWithoutMerging"
	default:
		return fmt.Sprintf("NSAttributedStringFormattingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownInterpretedSyntax
type NSAttributedStringMarkdownInterpretedSyntax int

const (
	// NSAttributedStringMarkdownInterpretedSyntaxFull: A syntax value that interprets the full Markdown syntax and produces all relevant attributes.
	NSAttributedStringMarkdownInterpretedSyntaxFull NSAttributedStringMarkdownInterpretedSyntax = 0
	// NSAttributedStringMarkdownInterpretedSyntaxInlineOnly: A syntax value that parses all Markdown text, but interprets only attributes that apply to inline spans.
	NSAttributedStringMarkdownInterpretedSyntaxInlineOnly NSAttributedStringMarkdownInterpretedSyntax = 1
	// NSAttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace: A syntax value that parses all Markdown text, but interprets only attributes that apply to inline spans, perserving white space.
	NSAttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace NSAttributedStringMarkdownInterpretedSyntax = 2
)

func (e NSAttributedStringMarkdownInterpretedSyntax) String() string {
	switch e {
	case NSAttributedStringMarkdownInterpretedSyntaxFull:
		return "NSAttributedStringMarkdownInterpretedSyntaxFull"
	case NSAttributedStringMarkdownInterpretedSyntaxInlineOnly:
		return "NSAttributedStringMarkdownInterpretedSyntaxInlineOnly"
	case NSAttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace:
		return "NSAttributedStringMarkdownInterpretedSyntaxInlineOnlyPreservingWhitespace"
	default:
		return fmt.Sprintf("NSAttributedStringMarkdownInterpretedSyntax(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSAttributedStringMarkdownParsingFailurePolicy
type NSAttributedStringMarkdownParsingFailurePolicy int

const (
	// NSAttributedStringMarkdownParsingFailureReturnError: A policy to return an error from the initializer if parsing fails.
	NSAttributedStringMarkdownParsingFailureReturnError NSAttributedStringMarkdownParsingFailurePolicy = 0
	// NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible: A policy to return a partially parsed string, if possible.
	NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible NSAttributedStringMarkdownParsingFailurePolicy = 1
)

func (e NSAttributedStringMarkdownParsingFailurePolicy) String() string {
	switch e {
	case NSAttributedStringMarkdownParsingFailureReturnError:
		return "NSAttributedStringMarkdownParsingFailureReturnError"
	case NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible:
		return "NSAttributedStringMarkdownParsingFailureReturnPartiallyParsedIfPossible"
	default:
		return fmt.Sprintf("NSAttributedStringMarkdownParsingFailurePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result
type NSBackgroundActivityResult int

const (
	// NSBackgroundActivityResultDeferred: System conditions have changed since the time the activity began executing, and deferral of additional work is recommended.
	NSBackgroundActivityResultDeferred NSBackgroundActivityResult = 2
	// NSBackgroundActivityResultFinished: The activity has finished executing.
	NSBackgroundActivityResultFinished NSBackgroundActivityResult = 1
)

func (e NSBackgroundActivityResult) String() string {
	switch e {
	case NSBackgroundActivityResultDeferred:
		return "NSBackgroundActivityResultDeferred"
	case NSBackgroundActivityResultFinished:
		return "NSBackgroundActivityResultFinished"
	default:
		return fmt.Sprintf("NSBackgroundActivityResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSBinarySearchingOptions
type NSBinarySearchingOptions int

const (
	// NSBinarySearchingFirstEqual: Specifies that the search should return the first object in the range that is equal to the given object.
	NSBinarySearchingFirstEqual NSBinarySearchingOptions = 256
	// NSBinarySearchingInsertionIndex: Returns the index at which you should insert the object in order to maintain a sorted array.
	NSBinarySearchingInsertionIndex NSBinarySearchingOptions = 1024
	// NSBinarySearchingLastEqual: Specifies that the search should return the last object in the range that is equal to the given object.
	NSBinarySearchingLastEqual NSBinarySearchingOptions = 512
)

func (e NSBinarySearchingOptions) String() string {
	switch e {
	case NSBinarySearchingFirstEqual:
		return "NSBinarySearchingFirstEqual"
	case NSBinarySearchingInsertionIndex:
		return "NSBinarySearchingInsertionIndex"
	case NSBinarySearchingLastEqual:
		return "NSBinarySearchingLastEqual"
	default:
		return fmt.Sprintf("NSBinarySearchingOptions(%d)", e)
	}
}

type NSBundleExecutableArchitecture uint

const (
	// NSBundleExecutableArchitectureARM64: The 64-bit ARM architecture.
	NSBundleExecutableArchitectureARM64 NSBundleExecutableArchitecture = 16777228
	// NSBundleExecutableArchitectureI386: The 32-bit Intel architecture.
	NSBundleExecutableArchitectureI386 NSBundleExecutableArchitecture = 7
	// NSBundleExecutableArchitecturePPC: The 32-bit PowerPC architecture.
	NSBundleExecutableArchitecturePPC NSBundleExecutableArchitecture = 18
	// NSBundleExecutableArchitecturePPC64: The 64-bit PowerPC architecture.
	NSBundleExecutableArchitecturePPC64 NSBundleExecutableArchitecture = 16777234
	// NSBundleExecutableArchitectureX86_64: The 64-bit Intel architecture.
	NSBundleExecutableArchitectureX86_64 NSBundleExecutableArchitecture = 16777223
)

func (e NSBundleExecutableArchitecture) String() string {
	switch e {
	case NSBundleExecutableArchitectureARM64:
		return "NSBundleExecutableArchitectureARM64"
	case NSBundleExecutableArchitectureI386:
		return "NSBundleExecutableArchitectureI386"
	case NSBundleExecutableArchitecturePPC:
		return "NSBundleExecutableArchitecturePPC"
	case NSBundleExecutableArchitecturePPC64:
		return "NSBundleExecutableArchitecturePPC64"
	case NSBundleExecutableArchitectureX86_64:
		return "NSBundleExecutableArchitectureX86_64"
	default:
		return fmt.Sprintf("NSBundleExecutableArchitecture(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/CountStyle-swift.enum
type NSByteCountFormatterCountStyle int

const (
	// NSByteCountFormatterCountStyleBinary: Causes 1024 bytes to be shown as 1 KB.
	NSByteCountFormatterCountStyleBinary NSByteCountFormatterCountStyle = 3
	// NSByteCountFormatterCountStyleDecimal: Causes 1000 bytes to be shown as 1 KB.
	NSByteCountFormatterCountStyleDecimal NSByteCountFormatterCountStyle = 2
	// NSByteCountFormatterCountStyleFile: Specifies display of file byte counts.
	NSByteCountFormatterCountStyleFile NSByteCountFormatterCountStyle = 0
	// NSByteCountFormatterCountStyleMemory: Specifies display of memory byte counts.
	NSByteCountFormatterCountStyleMemory NSByteCountFormatterCountStyle = 1
)

func (e NSByteCountFormatterCountStyle) String() string {
	switch e {
	case NSByteCountFormatterCountStyleBinary:
		return "NSByteCountFormatterCountStyleBinary"
	case NSByteCountFormatterCountStyleDecimal:
		return "NSByteCountFormatterCountStyleDecimal"
	case NSByteCountFormatterCountStyleFile:
		return "NSByteCountFormatterCountStyleFile"
	case NSByteCountFormatterCountStyleMemory:
		return "NSByteCountFormatterCountStyleMemory"
	default:
		return fmt.Sprintf("NSByteCountFormatterCountStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/ByteCountFormatter/Units
type NSByteCountFormatterUnits int

const (
	// NSByteCountFormatterUseAll: Can use any unit in the formatter content.
	NSByteCountFormatterUseAll NSByteCountFormatterUnits = 65535
	// NSByteCountFormatterUseBytes: Displays bytes in the formatter content.
	NSByteCountFormatterUseBytes NSByteCountFormatterUnits = 1
	// NSByteCountFormatterUseDefault: This causes default units appropriate for the platform to be used.
	NSByteCountFormatterUseDefault NSByteCountFormatterUnits = 0
	// NSByteCountFormatterUseEB: Displays exabytes in the formatter content.
	NSByteCountFormatterUseEB NSByteCountFormatterUnits = 64
	// NSByteCountFormatterUseGB: Displays gigabytes in the formatter content.
	NSByteCountFormatterUseGB NSByteCountFormatterUnits = 8
	// NSByteCountFormatterUseKB: Displays kilobytes in the formatter content.
	NSByteCountFormatterUseKB NSByteCountFormatterUnits = 2
	// NSByteCountFormatterUseMB: Displays megabytes in the formatter content.
	NSByteCountFormatterUseMB NSByteCountFormatterUnits = 4
	// NSByteCountFormatterUsePB: Displays petabyte in the formatter content.
	NSByteCountFormatterUsePB NSByteCountFormatterUnits = 32
	// NSByteCountFormatterUseTB: Displays terabytes in the formatter content.
	NSByteCountFormatterUseTB NSByteCountFormatterUnits = 16
	// NSByteCountFormatterUseYBOrHigher: Displays yottabytes in the formatter content.
	NSByteCountFormatterUseYBOrHigher NSByteCountFormatterUnits = 255
	// NSByteCountFormatterUseZB: Displays zettabytes in the formatter content.
	NSByteCountFormatterUseZB NSByteCountFormatterUnits = 128
)

func (e NSByteCountFormatterUnits) String() string {
	switch e {
	case NSByteCountFormatterUseAll:
		return "NSByteCountFormatterUseAll"
	case NSByteCountFormatterUseBytes:
		return "NSByteCountFormatterUseBytes"
	case NSByteCountFormatterUseDefault:
		return "NSByteCountFormatterUseDefault"
	case NSByteCountFormatterUseEB:
		return "NSByteCountFormatterUseEB"
	case NSByteCountFormatterUseGB:
		return "NSByteCountFormatterUseGB"
	case NSByteCountFormatterUseKB:
		return "NSByteCountFormatterUseKB"
	case NSByteCountFormatterUseMB:
		return "NSByteCountFormatterUseMB"
	case NSByteCountFormatterUsePB:
		return "NSByteCountFormatterUsePB"
	case NSByteCountFormatterUseTB:
		return "NSByteCountFormatterUseTB"
	case NSByteCountFormatterUseYBOrHigher:
		return "NSByteCountFormatterUseYBOrHigher"
	case NSByteCountFormatterUseZB:
		return "NSByteCountFormatterUseZB"
	default:
		return fmt.Sprintf("NSByteCountFormatterUnits(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/CalculationError
type NSCalculationError int

const (
	// NSCalculationDivideByZero: The caller tried to divide by `0`.
	NSCalculationDivideByZero NSCalculationError = 4
	// NSCalculationLossOfPrecision: The number can’t be represented in 38 significant digits.
	NSCalculationLossOfPrecision NSCalculationError = 1
	// NSCalculationNoError: No error occurred.
	NSCalculationNoError NSCalculationError = 0
	// NSCalculationOverflow: The number is too large to represent.
	NSCalculationOverflow NSCalculationError = 3
	// NSCalculationUnderflow: The number is too small to represent.
	NSCalculationUnderflow NSCalculationError = 2
)

func (e NSCalculationError) String() string {
	switch e {
	case NSCalculationDivideByZero:
		return "NSCalculationDivideByZero"
	case NSCalculationLossOfPrecision:
		return "NSCalculationLossOfPrecision"
	case NSCalculationNoError:
		return "NSCalculationNoError"
	case NSCalculationOverflow:
		return "NSCalculationOverflow"
	case NSCalculationUnderflow:
		return "NSCalculationUnderflow"
	default:
		return fmt.Sprintf("NSCalculationError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
type NSCalendarOptions int

const (
	// NSCalendarMatchFirst: Specifies that, if there are two or more matching times, the operation should return the first occurrence.
	NSCalendarMatchFirst NSCalendarOptions = 4096
	// NSCalendarMatchLast: Specifies that, if there are two or more matching times, the operation should return the last occurrence.
	NSCalendarMatchLast NSCalendarOptions = 8192
	// NSCalendarMatchNextTime: Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given [NSDateComponents] object, this method uses the  existing value of the missing unit and  preserve the lower units’ values.
	NSCalendarMatchNextTime NSCalendarOptions = 1024
	// NSCalendarMatchNextTimePreservingSmallerUnits: Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given [NSDateComponents] object, this method uses the  existing value of the missing unit and preserves the lower units’ values.
	NSCalendarMatchNextTimePreservingSmallerUnits NSCalendarOptions = 512
	// NSCalendarMatchPreviousTimePreservingSmallerUnits: Specifies that, when there is no matching time before the end of the next instance of the next highest unit specified in the given [NSDateComponents] object, this method uses the  existing value of the missing unit and preserves the lower units’ values.
	NSCalendarMatchPreviousTimePreservingSmallerUnits NSCalendarOptions = 256
	// NSCalendarMatchStrictly: Specifies that the operation should travel as far forward or backward as necessary looking for a match.
	NSCalendarMatchStrictly NSCalendarOptions = 2
	// NSCalendarSearchBackwards: Specifies that the operation should travel backwards to find the previous match before the given date.
	NSCalendarSearchBackwards NSCalendarOptions = 4
	// NSCalendarWrapComponents: Specifies that the components specified for an [NSDateComponents] object should be incremented and wrap around to zero/one on overflow, but should not cause higher units to be incremented.
	NSCalendarWrapComponents NSCalendarOptions = 1
)

func (e NSCalendarOptions) String() string {
	switch e {
	case NSCalendarMatchFirst:
		return "NSCalendarMatchFirst"
	case NSCalendarMatchLast:
		return "NSCalendarMatchLast"
	case NSCalendarMatchNextTime:
		return "NSCalendarMatchNextTime"
	case NSCalendarMatchNextTimePreservingSmallerUnits:
		return "NSCalendarMatchNextTimePreservingSmallerUnits"
	case NSCalendarMatchPreviousTimePreservingSmallerUnits:
		return "NSCalendarMatchPreviousTimePreservingSmallerUnits"
	case NSCalendarMatchStrictly:
		return "NSCalendarMatchStrictly"
	case NSCalendarSearchBackwards:
		return "NSCalendarSearchBackwards"
	case NSCalendarWrapComponents:
		return "NSCalendarWrapComponents"
	default:
		return fmt.Sprintf("NSCalendarOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
type NSCalendarUnit int

const (
	// NSCalendarUnitCalendar: Identifier for the calendar of a date components object.
	NSCalendarUnitCalendar NSCalendarUnit = 1048576
	// NSCalendarUnitDay: Identifier for the day unit.
	NSCalendarUnitDay NSCalendarUnit = 16
	NSCalendarUnitDayOfYear NSCalendarUnit = 65536
	// NSCalendarUnitEra: Identifier for the era unit.
	NSCalendarUnitEra NSCalendarUnit = 2
	// NSCalendarUnitHour: Identifier for the hour unit.
	NSCalendarUnitHour NSCalendarUnit = 32
	NSCalendarUnitIsLeapMonth NSCalendarUnit = 1073741824
	NSCalendarUnitIsRepeatedDay NSCalendarUnit = 2147483648
	// NSCalendarUnitMinute: Identifier for the minute unit.
	NSCalendarUnitMinute NSCalendarUnit = 64
	// NSCalendarUnitMonth: Identifier for the month unit.
	NSCalendarUnitMonth NSCalendarUnit = 8
	// NSCalendarUnitNanosecond: Identifier for the nanosecond unit.
	NSCalendarUnitNanosecond NSCalendarUnit = 32768
	// NSCalendarUnitQuarter: Identifier for the quarter of the calendar.
	NSCalendarUnitQuarter NSCalendarUnit = 2048
	// NSCalendarUnitSecond: Identifier for the second unit.
	NSCalendarUnitSecond NSCalendarUnit = 128
	// NSCalendarUnitTimeZone: Identifier for the time zone of a date components object.
	NSCalendarUnitTimeZone NSCalendarUnit = 2097152
	// NSCalendarUnitWeekOfMonth: Identifier for the week of the month calendar unit.
	NSCalendarUnitWeekOfMonth NSCalendarUnit = 4096
	// NSCalendarUnitWeekOfYear: Identifier for the week of the year calendar unit.
	NSCalendarUnitWeekOfYear NSCalendarUnit = 8192
	// NSCalendarUnitWeekday: Identifier for the weekday unit.
	NSCalendarUnitWeekday NSCalendarUnit = 512
	// NSCalendarUnitWeekdayOrdinal: Identifier for the ordinal weekday unit.
	NSCalendarUnitWeekdayOrdinal NSCalendarUnit = 1024
	// NSCalendarUnitYear: Identifier for the year unit.
	NSCalendarUnitYear NSCalendarUnit = 4
	// NSCalendarUnitYearForWeekOfYear: Identifier for the week-counting year unit.
	NSCalendarUnitYearForWeekOfYear NSCalendarUnit = 16384
	// Deprecated.
	NSCalendarCalendarUnit NSCalendarUnit = 1048576
	// Deprecated.
	NSDayCalendarUnit NSCalendarUnit = 16
	// Deprecated.
	NSEraCalendarUnit NSCalendarUnit = 2
	// Deprecated.
	NSHourCalendarUnit NSCalendarUnit = 32
	// Deprecated.
	NSMinuteCalendarUnit NSCalendarUnit = 64
	// Deprecated.
	NSMonthCalendarUnit NSCalendarUnit = 8
	// Deprecated.
	NSQuarterCalendarUnit NSCalendarUnit = 2048
	// Deprecated.
	NSSecondCalendarUnit NSCalendarUnit = 128
	// Deprecated.
	NSTimeZoneCalendarUnit NSCalendarUnit = 2097152
	// Deprecated.
	NSWeekCalendarUnit NSCalendarUnit = 256
	// Deprecated.
	NSWeekOfMonthCalendarUnit NSCalendarUnit = 4096
	// Deprecated.
	NSWeekOfYearCalendarUnit NSCalendarUnit = 8192
	// Deprecated.
	NSWeekdayCalendarUnit NSCalendarUnit = 512
	// Deprecated.
	NSWeekdayOrdinalCalendarUnit NSCalendarUnit = 1024
	// Deprecated.
	NSYearCalendarUnit NSCalendarUnit = 4
	// Deprecated.
	NSYearForWeekOfYearCalendarUnit NSCalendarUnit = 16384
)

func (e NSCalendarUnit) String() string {
	switch e {
	case NSCalendarUnitCalendar:
		return "NSCalendarUnitCalendar"
	case NSCalendarUnitDay:
		return "NSCalendarUnitDay"
	case NSCalendarUnitDayOfYear:
		return "NSCalendarUnitDayOfYear"
	case NSCalendarUnitEra:
		return "NSCalendarUnitEra"
	case NSCalendarUnitHour:
		return "NSCalendarUnitHour"
	case NSCalendarUnitIsLeapMonth:
		return "NSCalendarUnitIsLeapMonth"
	case NSCalendarUnitIsRepeatedDay:
		return "NSCalendarUnitIsRepeatedDay"
	case NSCalendarUnitMinute:
		return "NSCalendarUnitMinute"
	case NSCalendarUnitMonth:
		return "NSCalendarUnitMonth"
	case NSCalendarUnitNanosecond:
		return "NSCalendarUnitNanosecond"
	case NSCalendarUnitQuarter:
		return "NSCalendarUnitQuarter"
	case NSCalendarUnitSecond:
		return "NSCalendarUnitSecond"
	case NSCalendarUnitTimeZone:
		return "NSCalendarUnitTimeZone"
	case NSCalendarUnitWeekOfMonth:
		return "NSCalendarUnitWeekOfMonth"
	case NSCalendarUnitWeekOfYear:
		return "NSCalendarUnitWeekOfYear"
	case NSCalendarUnitWeekday:
		return "NSCalendarUnitWeekday"
	case NSCalendarUnitWeekdayOrdinal:
		return "NSCalendarUnitWeekdayOrdinal"
	case NSCalendarUnitYear:
		return "NSCalendarUnitYear"
	case NSCalendarUnitYearForWeekOfYear:
		return "NSCalendarUnitYearForWeekOfYear"
	case NSWeekCalendarUnit:
		return "NSWeekCalendarUnit"
	default:
		return fmt.Sprintf("NSCalendarUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSCollectionChangeType
type NSCollectionChangeType int

const (
	// NSCollectionChangeInsert: A change type that represents the insertion of an object into an ordered collection.
	NSCollectionChangeInsert NSCollectionChangeType = 0
	// NSCollectionChangeRemove: A change type that represents the removal of an object from an ordered collection.
	NSCollectionChangeRemove NSCollectionChangeType = 1
)

func (e NSCollectionChangeType) String() string {
	switch e {
	case NSCollectionChangeInsert:
		return "NSCollectionChangeInsert"
	case NSCollectionChangeRemove:
		return "NSCollectionChangeRemove"
	default:
		return fmt.Sprintf("NSCollectionChangeType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Modifier
type NSComparisonPredicateModifier int

const (
	// NSAllPredicateModifier: A predicate to compare all entries in the destination of a to-many relationship.
	NSAllPredicateModifier NSComparisonPredicateModifier = 1
	// NSAnyPredicateModifier: A predicate to match with any entry in the destination of a to-many relationship.
	NSAnyPredicateModifier NSComparisonPredicateModifier = 2
	// NSDirectPredicateModifier: A predicate to compare directly the left and right hand sides.
	NSDirectPredicateModifier NSComparisonPredicateModifier = 0
)

func (e NSComparisonPredicateModifier) String() string {
	switch e {
	case NSAllPredicateModifier:
		return "NSAllPredicateModifier"
	case NSAnyPredicateModifier:
		return "NSAnyPredicateModifier"
	case NSDirectPredicateModifier:
		return "NSDirectPredicateModifier"
	default:
		return fmt.Sprintf("NSComparisonPredicateModifier(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Options-swift.struct
type NSComparisonPredicateOptions int

const (
	// NSCaseInsensitivePredicateOption: A case-insensitive predicate.
	NSCaseInsensitivePredicateOption NSComparisonPredicateOptions = 1
	// NSDiacriticInsensitivePredicateOption: A diacritic-insensitive predicate.
	NSDiacriticInsensitivePredicateOption NSComparisonPredicateOptions = 2
	// NSNormalizedPredicateOption: A predicate that indicates you’ve preprocessed the strings to compare.
	NSNormalizedPredicateOption NSComparisonPredicateOptions = 4
)

func (e NSComparisonPredicateOptions) String() string {
	switch e {
	case NSCaseInsensitivePredicateOption:
		return "NSCaseInsensitivePredicateOption"
	case NSDiacriticInsensitivePredicateOption:
		return "NSDiacriticInsensitivePredicateOption"
	case NSNormalizedPredicateOption:
		return "NSNormalizedPredicateOption"
	default:
		return fmt.Sprintf("NSComparisonPredicateOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/ComparisonResult
type ComparisonResult int

const (
	// NSOrderedAscending: The left operand is smaller than the right operand.
	NSOrderedAscending ComparisonResult = -1
	// NSOrderedDescending: The left operand is greater than the right operand.
	NSOrderedDescending ComparisonResult = 1
	// NSOrderedSame: The two operands are equal.
	NSOrderedSame ComparisonResult = 0
)

func (e ComparisonResult) String() string {
	switch e {
	case NSOrderedAscending:
		return "NSOrderedAscending"
	case NSOrderedDescending:
		return "NSOrderedDescending"
	case NSOrderedSame:
		return "NSOrderedSame"
	default:
		return fmt.Sprintf("ComparisonResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSCompoundPredicate/LogicalType
type NSCompoundPredicateType int

const (
	// NSAndPredicateType: A logical AND predicate.
	NSAndPredicateType NSCompoundPredicateType = 1
	// NSNotPredicateType: A logical NOT predicate.
	NSNotPredicateType NSCompoundPredicateType = 0
	// NSOrPredicateType: A logical OR predicate.
	NSOrPredicateType NSCompoundPredicateType = 2
)

func (e NSCompoundPredicateType) String() string {
	switch e {
	case NSAndPredicateType:
		return "NSAndPredicateType"
	case NSNotPredicateType:
		return "NSNotPredicateType"
	case NSOrPredicateType:
		return "NSOrPredicateType"
	default:
		return fmt.Sprintf("NSCompoundPredicateType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSData/Base64DecodingOptions
type NSDataBase64DecodingOptions int

const (
	// NSDataBase64DecodingIgnoreUnknownCharacters: Modify the decoding algorithm so that it ignores unknown non-Base-64 bytes, including line ending characters.
	NSDataBase64DecodingIgnoreUnknownCharacters NSDataBase64DecodingOptions = 1
)

func (e NSDataBase64DecodingOptions) String() string {
	switch e {
	case NSDataBase64DecodingIgnoreUnknownCharacters:
		return "NSDataBase64DecodingIgnoreUnknownCharacters"
	default:
		return fmt.Sprintf("NSDataBase64DecodingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSData/Base64EncodingOptions
type NSDataBase64EncodingOptions int

const (
	// NSDataBase64Encoding64CharacterLineLength: Set the maximum line length to 64 characters, after which a line ending is inserted.
	NSDataBase64Encoding64CharacterLineLength NSDataBase64EncodingOptions = 1
	// NSDataBase64Encoding76CharacterLineLength: Set the maximum line length to 76 characters, after which a line ending is inserted.
	NSDataBase64Encoding76CharacterLineLength NSDataBase64EncodingOptions = 2
	// NSDataBase64EncodingEndLineWithCarriageReturn: When a maximum line length is set, specify that the line ending to insert should include a carriage return.
	NSDataBase64EncodingEndLineWithCarriageReturn NSDataBase64EncodingOptions = 16
	// NSDataBase64EncodingEndLineWithLineFeed: When a maximum line length is set, specify that the line ending to insert should include a line feed.
	NSDataBase64EncodingEndLineWithLineFeed NSDataBase64EncodingOptions = 32
)

func (e NSDataBase64EncodingOptions) String() string {
	switch e {
	case NSDataBase64Encoding64CharacterLineLength:
		return "NSDataBase64Encoding64CharacterLineLength"
	case NSDataBase64Encoding76CharacterLineLength:
		return "NSDataBase64Encoding76CharacterLineLength"
	case NSDataBase64EncodingEndLineWithCarriageReturn:
		return "NSDataBase64EncodingEndLineWithCarriageReturn"
	case NSDataBase64EncodingEndLineWithLineFeed:
		return "NSDataBase64EncodingEndLineWithLineFeed"
	default:
		return fmt.Sprintf("NSDataBase64EncodingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSData/CompressionAlgorithm
type NSDataCompressionAlgorithm int

const (
	// NSDataCompressionAlgorithmLZ4: The LZ4 compression algorithm, recommended for fast compression.
	NSDataCompressionAlgorithmLZ4 NSDataCompressionAlgorithm = 1
	// NSDataCompressionAlgorithmLZFSE: The LZFSE compression algorithm, recommended for use on Apple platforms.
	NSDataCompressionAlgorithmLZFSE NSDataCompressionAlgorithm = 0
	// NSDataCompressionAlgorithmLZMA: The LZMA compression algorithm, recommended for high-compression ratio.
	NSDataCompressionAlgorithmLZMA NSDataCompressionAlgorithm = 2
	// NSDataCompressionAlgorithmZlib: The zlib compression algorithm, recommended for cross-platform compression.
	NSDataCompressionAlgorithmZlib NSDataCompressionAlgorithm = 3
)

func (e NSDataCompressionAlgorithm) String() string {
	switch e {
	case NSDataCompressionAlgorithmLZ4:
		return "NSDataCompressionAlgorithmLZ4"
	case NSDataCompressionAlgorithmLZFSE:
		return "NSDataCompressionAlgorithmLZFSE"
	case NSDataCompressionAlgorithmLZMA:
		return "NSDataCompressionAlgorithmLZMA"
	case NSDataCompressionAlgorithmZlib:
		return "NSDataCompressionAlgorithmZlib"
	default:
		return fmt.Sprintf("NSDataCompressionAlgorithm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSData/ReadingOptions
type NSDataReadingOptions int

const (
	// NSDataReadingMappedAlways: Hint to map the file in if possible.
	NSDataReadingMappedAlways NSDataReadingOptions = 8
	// NSDataReadingMappedIfSafe: A hint indicating the file should be mapped into virtual memory, if possible and safe.
	NSDataReadingMappedIfSafe NSDataReadingOptions = 1
	// NSDataReadingUncached: A hint indicating the file should not be stored in the file-system caches.
	NSDataReadingUncached NSDataReadingOptions = 2
	// Deprecated: use NSDataReadingMappedIfSafe.
	NSDataReadingMapped NSDataReadingOptions = 1
	// Deprecated: use NSDataReadingMapped.
	NSMappedRead NSDataReadingOptions = 1
	// Deprecated: use NSDataReadingUncached.
	NSUncachedRead NSDataReadingOptions = 2
)

func (e NSDataReadingOptions) String() string {
	switch e {
	case NSDataReadingMappedAlways:
		return "NSDataReadingMappedAlways"
	case NSDataReadingMappedIfSafe:
		return "NSDataReadingMappedIfSafe"
	case NSDataReadingUncached:
		return "NSDataReadingUncached"
	default:
		return fmt.Sprintf("NSDataReadingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSData/SearchOptions
type NSDataSearchOptions int

const (
	// NSDataSearchAnchored: Search is limited to start (or end, if searching backwards) of the data object.
	NSDataSearchAnchored NSDataSearchOptions = 2
	// NSDataSearchBackwards: Search from the end of the data object.
	NSDataSearchBackwards NSDataSearchOptions = 1
)

func (e NSDataSearchOptions) String() string {
	switch e {
	case NSDataSearchAnchored:
		return "NSDataSearchAnchored"
	case NSDataSearchBackwards:
		return "NSDataSearchBackwards"
	default:
		return fmt.Sprintf("NSDataSearchOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSData/WritingOptions
type NSDataWritingOptions int

const (
	// NSDataWritingAtomic: An option to write data to an auxiliary file first and then replace the original file with the auxiliary file when the write completes.
	NSDataWritingAtomic NSDataWritingOptions = 1
	// NSDataWritingFileProtectionComplete: An option to make the file accessible only while the device is unlocked.
	NSDataWritingFileProtectionComplete NSDataWritingOptions = 536870912
	// NSDataWritingFileProtectionCompleteUnlessOpen: An option to allow the file to be accessible while the device is unlocked or the file is already open.
	NSDataWritingFileProtectionCompleteUnlessOpen NSDataWritingOptions = 805306368
	// NSDataWritingFileProtectionCompleteUntilFirstUserAuthentication: An option to allow the file to be accessible after a user first unlocks the device.
	NSDataWritingFileProtectionCompleteUntilFirstUserAuthentication NSDataWritingOptions = 1073741824
	NSDataWritingFileProtectionCompleteWhenUserInactive NSDataWritingOptions = 1342177280
	// NSDataWritingFileProtectionMask: An option the system uses when determining the file protection options that the system assigns to the data.
	NSDataWritingFileProtectionMask NSDataWritingOptions = 4026531840
	// NSDataWritingFileProtectionNone: An option to not encrypt the file when writing it out.
	NSDataWritingFileProtectionNone NSDataWritingOptions = 268435456
	// NSDataWritingWithoutOverwriting: An option that attempts to write data to a file and fails with an error if the destination file already exists.
	NSDataWritingWithoutOverwriting NSDataWritingOptions = 2
	// Deprecated.
	NSAtomicWrite NSDataWritingOptions = 4026531841
)

func (e NSDataWritingOptions) String() string {
	switch e {
	case NSDataWritingAtomic:
		return "NSDataWritingAtomic"
	case NSDataWritingFileProtectionComplete:
		return "NSDataWritingFileProtectionComplete"
	case NSDataWritingFileProtectionCompleteUnlessOpen:
		return "NSDataWritingFileProtectionCompleteUnlessOpen"
	case NSDataWritingFileProtectionCompleteUntilFirstUserAuthentication:
		return "NSDataWritingFileProtectionCompleteUntilFirstUserAuthentication"
	case NSDataWritingFileProtectionCompleteWhenUserInactive:
		return "NSDataWritingFileProtectionCompleteWhenUserInactive"
	case NSDataWritingFileProtectionMask:
		return "NSDataWritingFileProtectionMask"
	case NSDataWritingFileProtectionNone:
		return "NSDataWritingFileProtectionNone"
	case NSDataWritingWithoutOverwriting:
		return "NSDataWritingWithoutOverwriting"
	case NSAtomicWrite:
		return "NSAtomicWrite"
	default:
		return fmt.Sprintf("NSDataWritingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/UnitsStyle-swift.enum
type NSDateComponentsFormatterUnitsStyle int

const (
	// NSDateComponentsFormatterUnitsStyleAbbreviated: A style that uses the most abbreviated spelling for units of time.
	NSDateComponentsFormatterUnitsStyleAbbreviated NSDateComponentsFormatterUnitsStyle = 1
	// NSDateComponentsFormatterUnitsStyleBrief: A style that uses a shortened spelling for units of time that is shorter than DateComponentsFormatter.UnitsStyle.short.
	NSDateComponentsFormatterUnitsStyleBrief NSDateComponentsFormatterUnitsStyle = 5
	// NSDateComponentsFormatterUnitsStyleFull: A style that spells out the units of time, but not the quantities.
	NSDateComponentsFormatterUnitsStyleFull NSDateComponentsFormatterUnitsStyle = 3
	// NSDateComponentsFormatterUnitsStylePositional: A style that uses the position of a unit of time to identify its value.
	NSDateComponentsFormatterUnitsStylePositional NSDateComponentsFormatterUnitsStyle = 0
	// NSDateComponentsFormatterUnitsStyleShort: A style that uses a shortened spelling for units.
	NSDateComponentsFormatterUnitsStyleShort NSDateComponentsFormatterUnitsStyle = 2
	// NSDateComponentsFormatterUnitsStyleSpellOut: A style that spells out the units and quantities of time.
	NSDateComponentsFormatterUnitsStyleSpellOut NSDateComponentsFormatterUnitsStyle = 4
)

func (e NSDateComponentsFormatterUnitsStyle) String() string {
	switch e {
	case NSDateComponentsFormatterUnitsStyleAbbreviated:
		return "NSDateComponentsFormatterUnitsStyleAbbreviated"
	case NSDateComponentsFormatterUnitsStyleBrief:
		return "NSDateComponentsFormatterUnitsStyleBrief"
	case NSDateComponentsFormatterUnitsStyleFull:
		return "NSDateComponentsFormatterUnitsStyleFull"
	case NSDateComponentsFormatterUnitsStylePositional:
		return "NSDateComponentsFormatterUnitsStylePositional"
	case NSDateComponentsFormatterUnitsStyleShort:
		return "NSDateComponentsFormatterUnitsStyleShort"
	case NSDateComponentsFormatterUnitsStyleSpellOut:
		return "NSDateComponentsFormatterUnitsStyleSpellOut"
	default:
		return fmt.Sprintf("NSDateComponentsFormatterUnitsStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/ZeroFormattingBehavior-swift.struct
type NSDateComponentsFormatterZeroFormattingBehavior int

const (
	// NSDateComponentsFormatterZeroFormattingBehaviorDefault: The default formatting behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorDefault NSDateComponentsFormatterZeroFormattingBehavior = 1
	// NSDateComponentsFormatterZeroFormattingBehaviorDropAll: The drop all zero units behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorDropAll NSDateComponentsFormatterZeroFormattingBehavior = 0
	// NSDateComponentsFormatterZeroFormattingBehaviorDropLeading: The drop leading zeroes formatting behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorDropLeading NSDateComponentsFormatterZeroFormattingBehavior = 2
	// NSDateComponentsFormatterZeroFormattingBehaviorDropMiddle: The drop middle zero units behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorDropMiddle NSDateComponentsFormatterZeroFormattingBehavior = 4
	// NSDateComponentsFormatterZeroFormattingBehaviorDropTrailing: The drop trailing zero units behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorDropTrailing NSDateComponentsFormatterZeroFormattingBehavior = 8
	// NSDateComponentsFormatterZeroFormattingBehaviorNone: No formatting behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorNone NSDateComponentsFormatterZeroFormattingBehavior = 0
	// NSDateComponentsFormatterZeroFormattingBehaviorPad: The add padding zeroes behavior.
	NSDateComponentsFormatterZeroFormattingBehaviorPad NSDateComponentsFormatterZeroFormattingBehavior = 65536
)

func (e NSDateComponentsFormatterZeroFormattingBehavior) String() string {
	switch e {
	case NSDateComponentsFormatterZeroFormattingBehaviorDefault:
		return "NSDateComponentsFormatterZeroFormattingBehaviorDefault"
	case NSDateComponentsFormatterZeroFormattingBehaviorDropAll:
		return "NSDateComponentsFormatterZeroFormattingBehaviorDropAll"
	case NSDateComponentsFormatterZeroFormattingBehaviorDropLeading:
		return "NSDateComponentsFormatterZeroFormattingBehaviorDropLeading"
	case NSDateComponentsFormatterZeroFormattingBehaviorDropMiddle:
		return "NSDateComponentsFormatterZeroFormattingBehaviorDropMiddle"
	case NSDateComponentsFormatterZeroFormattingBehaviorDropTrailing:
		return "NSDateComponentsFormatterZeroFormattingBehaviorDropTrailing"
	case NSDateComponentsFormatterZeroFormattingBehaviorPad:
		return "NSDateComponentsFormatterZeroFormattingBehaviorPad"
	default:
		return fmt.Sprintf("NSDateComponentsFormatterZeroFormattingBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior
type NSDateFormatterBehavior int

const (
	// NSDateFormatterBehavior10_0: # Discussion
	NSDateFormatterBehavior10_0 NSDateFormatterBehavior = 1000
	// NSDateFormatterBehavior10_4: # Discussion
	NSDateFormatterBehavior10_4 NSDateFormatterBehavior = 1040
	// NSDateFormatterBehaviorDefault: # Discussion
	NSDateFormatterBehaviorDefault NSDateFormatterBehavior = 0
)

func (e NSDateFormatterBehavior) String() string {
	switch e {
	case NSDateFormatterBehavior10_0:
		return "NSDateFormatterBehavior10_0"
	case NSDateFormatterBehavior10_4:
		return "NSDateFormatterBehavior10_4"
	case NSDateFormatterBehaviorDefault:
		return "NSDateFormatterBehaviorDefault"
	default:
		return fmt.Sprintf("NSDateFormatterBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DateFormatter/Style
type NSDateFormatterStyle int

const (
	// NSDateFormatterFullStyle: # Discussion
	NSDateFormatterFullStyle NSDateFormatterStyle = 0
	// NSDateFormatterLongStyle: # Discussion
	NSDateFormatterLongStyle NSDateFormatterStyle = 3
	// NSDateFormatterMediumStyle: # Discussion
	NSDateFormatterMediumStyle NSDateFormatterStyle = 2
	// NSDateFormatterNoStyle: # Discussion
	NSDateFormatterNoStyle NSDateFormatterStyle = 0
	// NSDateFormatterShortStyle: # Discussion
	NSDateFormatterShortStyle NSDateFormatterStyle = 1
)

func (e NSDateFormatterStyle) String() string {
	switch e {
	case NSDateFormatterFullStyle:
		return "NSDateFormatterFullStyle"
	case NSDateFormatterLongStyle:
		return "NSDateFormatterLongStyle"
	case NSDateFormatterMediumStyle:
		return "NSDateFormatterMediumStyle"
	case NSDateFormatterShortStyle:
		return "NSDateFormatterShortStyle"
	default:
		return fmt.Sprintf("NSDateFormatterStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/Style
type NSDateIntervalFormatterStyle int

const (
	// NSDateIntervalFormatterFullStyle: A fully spelled out date or time format.
	NSDateIntervalFormatterFullStyle NSDateIntervalFormatterStyle = 4
	// NSDateIntervalFormatterLongStyle: A long length date or time format.
	NSDateIntervalFormatterLongStyle NSDateIntervalFormatterStyle = 3
	// NSDateIntervalFormatterMediumStyle: A medium length date or time format.
	NSDateIntervalFormatterMediumStyle NSDateIntervalFormatterStyle = 2
	// NSDateIntervalFormatterNoStyle: No information for the date or time.
	NSDateIntervalFormatterNoStyle NSDateIntervalFormatterStyle = 0
	// NSDateIntervalFormatterShortStyle: An abbreviated date or time format.
	NSDateIntervalFormatterShortStyle NSDateIntervalFormatterStyle = 1
)

func (e NSDateIntervalFormatterStyle) String() string {
	switch e {
	case NSDateIntervalFormatterFullStyle:
		return "NSDateIntervalFormatterFullStyle"
	case NSDateIntervalFormatterLongStyle:
		return "NSDateIntervalFormatterLongStyle"
	case NSDateIntervalFormatterMediumStyle:
		return "NSDateIntervalFormatterMediumStyle"
	case NSDateIntervalFormatterNoStyle:
		return "NSDateIntervalFormatterNoStyle"
	case NSDateIntervalFormatterShortStyle:
		return "NSDateIntervalFormatterShortStyle"
	default:
		return fmt.Sprintf("NSDateIntervalFormatterStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSCoder/DecodingFailurePolicy-swift.enum
type NSDecodingFailurePolicy int

const (
	// NSDecodingFailurePolicyRaiseException: A failure policy that directs the coder to raise an exception.
	NSDecodingFailurePolicyRaiseException NSDecodingFailurePolicy = 0
	// NSDecodingFailurePolicySetErrorAndReturn: A failure policy that directs the coder to capture the failure as an error object.
	NSDecodingFailurePolicySetErrorAndReturn NSDecodingFailurePolicy = 1
)

func (e NSDecodingFailurePolicy) String() string {
	switch e {
	case NSDecodingFailurePolicyRaiseException:
		return "NSDecodingFailurePolicyRaiseException"
	case NSDecodingFailurePolicySetErrorAndReturn:
		return "NSDecodingFailurePolicySetErrorAndReturn"
	default:
		return fmt.Sprintf("NSDecodingFailurePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/DirectoryEnumerationOptions
type NSDirectoryEnumerationOptions int

const (
	NSDirectoryEnumerationIncludesDirectoriesPostOrder NSDirectoryEnumerationOptions = 8
	NSDirectoryEnumerationProducesRelativePathURLs NSDirectoryEnumerationOptions = 16
	// NSDirectoryEnumerationSkipsHiddenFiles: An option to skip hidden files.
	NSDirectoryEnumerationSkipsHiddenFiles NSDirectoryEnumerationOptions = 4
	// NSDirectoryEnumerationSkipsPackageDescendants: An option to treat packages like files and not descend into their contents.
	NSDirectoryEnumerationSkipsPackageDescendants NSDirectoryEnumerationOptions = 2
	// NSDirectoryEnumerationSkipsSubdirectoryDescendants: An option to perform a shallow enumeration that doesn’t descend into directories.
	NSDirectoryEnumerationSkipsSubdirectoryDescendants NSDirectoryEnumerationOptions = 1
)

func (e NSDirectoryEnumerationOptions) String() string {
	switch e {
	case NSDirectoryEnumerationIncludesDirectoriesPostOrder:
		return "NSDirectoryEnumerationIncludesDirectoriesPostOrder"
	case NSDirectoryEnumerationProducesRelativePathURLs:
		return "NSDirectoryEnumerationProducesRelativePathURLs"
	case NSDirectoryEnumerationSkipsHiddenFiles:
		return "NSDirectoryEnumerationSkipsHiddenFiles"
	case NSDirectoryEnumerationSkipsPackageDescendants:
		return "NSDirectoryEnumerationSkipsPackageDescendants"
	case NSDirectoryEnumerationSkipsSubdirectoryDescendants:
		return "NSDirectoryEnumerationSkipsSubdirectoryDescendants"
	default:
		return fmt.Sprintf("NSDirectoryEnumerationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/Options
type NSDistributedNotificationOptions int

const (
	NSDistributedNotificationDeliverImmediately NSDistributedNotificationOptions = 1
	NSDistributedNotificationPostToAllSessions NSDistributedNotificationOptions = 2
)

func (e NSDistributedNotificationOptions) String() string {
	switch e {
	case NSDistributedNotificationDeliverImmediately:
		return "NSDistributedNotificationDeliverImmediately"
	case NSDistributedNotificationPostToAllSessions:
		return "NSDistributedNotificationPostToAllSessions"
	default:
		return fmt.Sprintf("NSDistributedNotificationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit
type NSEnergyFormatterUnit int

const (
	// NSEnergyFormatterUnitCalorie: The calorie unit.
	NSEnergyFormatterUnitCalorie NSEnergyFormatterUnit = 1792
	// NSEnergyFormatterUnitJoule: The joule unit.
	NSEnergyFormatterUnitJoule NSEnergyFormatterUnit = 11
	// NSEnergyFormatterUnitKilocalorie: The kilocalorie unit.
	NSEnergyFormatterUnitKilocalorie NSEnergyFormatterUnit = 1792
	// NSEnergyFormatterUnitKilojoule: The kilojoule unit.
	NSEnergyFormatterUnitKilojoule NSEnergyFormatterUnit = 14
)

func (e NSEnergyFormatterUnit) String() string {
	switch e {
	case NSEnergyFormatterUnitCalorie:
		return "NSEnergyFormatterUnitCalorie"
	case NSEnergyFormatterUnitJoule:
		return "NSEnergyFormatterUnitJoule"
	case NSEnergyFormatterUnitKilojoule:
		return "NSEnergyFormatterUnitKilojoule"
	default:
		return fmt.Sprintf("NSEnergyFormatterUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSEnumerationOptions
type NSEnumerationOptions int

const (
	// NSEnumerationConcurrent: Specifies that the Block enumeration should be concurrent.
	NSEnumerationConcurrent NSEnumerationOptions = 1
	// NSEnumerationReverse: Specifies that the enumeration should be performed in reverse.
	NSEnumerationReverse NSEnumerationOptions = 2
)

func (e NSEnumerationOptions) String() string {
	switch e {
	case NSEnumerationConcurrent:
		return "NSEnumerationConcurrent"
	case NSEnumerationReverse:
		return "NSEnumerationReverse"
	default:
		return fmt.Sprintf("NSEnumerationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSExpression/ExpressionType-swift.enum
type NSExpressionType int

const (
	// NSAggregateExpressionType: An expression that defines an aggregate of [NSExpression] objects.
	NSAggregateExpressionType NSExpressionType = 14
	// NSAnyKeyExpressionType: An expression that represents any key.
	NSAnyKeyExpressionType NSExpressionType = 15
	// NSBlockExpressionType: An expression that uses a Block.
	NSBlockExpressionType NSExpressionType = 19
	NSConditionalExpressionType NSExpressionType = 20
	// NSConstantValueExpressionType: An expression that always returns the same value.
	NSConstantValueExpressionType NSExpressionType = 0
	// NSEvaluatedObjectExpressionType: An expression that always returns the parameter object itself.
	NSEvaluatedObjectExpressionType NSExpressionType = 1
	// NSFunctionExpressionType: An expression that returns the result of evaluating a function.
	NSFunctionExpressionType NSExpressionType = 4
	// NSIntersectSetExpressionType: An expression that creates an intersection of the results of two nested expressions.
	NSIntersectSetExpressionType NSExpressionType = 6
	// NSKeyPathExpressionType: An expression that returns something that can be used as a key path.
	NSKeyPathExpressionType NSExpressionType = 3
	// NSMinusSetExpressionType: An expression that combines two nested expression results by set subtraction.
	NSMinusSetExpressionType NSExpressionType = 7
	// NSSubqueryExpressionType: An expression that filters a collection using a subpredicate.
	NSSubqueryExpressionType NSExpressionType = 13
	// NSUnionSetExpressionType: An expression that creates a union of the results of two nested expressions.
	NSUnionSetExpressionType NSExpressionType = 5
	// NSVariableExpressionType: An expression that always returns whatever value is associated with the key specified by ‘variable’ in the bindings dictionary.
	NSVariableExpressionType NSExpressionType = 2
)

func (e NSExpressionType) String() string {
	switch e {
	case NSAggregateExpressionType:
		return "NSAggregateExpressionType"
	case NSAnyKeyExpressionType:
		return "NSAnyKeyExpressionType"
	case NSBlockExpressionType:
		return "NSBlockExpressionType"
	case NSConditionalExpressionType:
		return "NSConditionalExpressionType"
	case NSConstantValueExpressionType:
		return "NSConstantValueExpressionType"
	case NSEvaluatedObjectExpressionType:
		return "NSEvaluatedObjectExpressionType"
	case NSFunctionExpressionType:
		return "NSFunctionExpressionType"
	case NSIntersectSetExpressionType:
		return "NSIntersectSetExpressionType"
	case NSKeyPathExpressionType:
		return "NSKeyPathExpressionType"
	case NSMinusSetExpressionType:
		return "NSMinusSetExpressionType"
	case NSSubqueryExpressionType:
		return "NSSubqueryExpressionType"
	case NSUnionSetExpressionType:
		return "NSUnionSetExpressionType"
	case NSVariableExpressionType:
		return "NSVariableExpressionType"
	default:
		return fmt.Sprintf("NSExpressionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/ReadingOptions
type NSFileCoordinatorReadingOptions int

const (
	// NSFileCoordinatorReadingForUploading: Specify this content when reading an item for the purpose of uploading its contents.
	NSFileCoordinatorReadingForUploading NSFileCoordinatorReadingOptions = 8
	// NSFileCoordinatorReadingImmediatelyAvailableMetadataOnly: Specify this constant if you want to read an item’s metadata without triggering a download.
	NSFileCoordinatorReadingImmediatelyAvailableMetadataOnly NSFileCoordinatorReadingOptions = 4
	// NSFileCoordinatorReadingResolvesSymbolicLink: # Discussion
	NSFileCoordinatorReadingResolvesSymbolicLink NSFileCoordinatorReadingOptions = 2
	// NSFileCoordinatorReadingWithoutChanges: # Discussion
	NSFileCoordinatorReadingWithoutChanges NSFileCoordinatorReadingOptions = 1
)

func (e NSFileCoordinatorReadingOptions) String() string {
	switch e {
	case NSFileCoordinatorReadingForUploading:
		return "NSFileCoordinatorReadingForUploading"
	case NSFileCoordinatorReadingImmediatelyAvailableMetadataOnly:
		return "NSFileCoordinatorReadingImmediatelyAvailableMetadataOnly"
	case NSFileCoordinatorReadingResolvesSymbolicLink:
		return "NSFileCoordinatorReadingResolvesSymbolicLink"
	case NSFileCoordinatorReadingWithoutChanges:
		return "NSFileCoordinatorReadingWithoutChanges"
	default:
		return fmt.Sprintf("NSFileCoordinatorReadingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/WritingOptions
type NSFileCoordinatorWritingOptions int

const (
	// NSFileCoordinatorWritingContentIndependentMetadataOnly: Select this option when writing to change the file’s metadata only and not its contents.
	NSFileCoordinatorWritingContentIndependentMetadataOnly NSFileCoordinatorWritingOptions = 16
	// NSFileCoordinatorWritingForDeleting: # Discussion
	NSFileCoordinatorWritingForDeleting NSFileCoordinatorWritingOptions = 1
	// NSFileCoordinatorWritingForMerging: # Discussion
	NSFileCoordinatorWritingForMerging NSFileCoordinatorWritingOptions = 4
	// NSFileCoordinatorWritingForMoving: # Discussion
	NSFileCoordinatorWritingForMoving NSFileCoordinatorWritingOptions = 2
	// NSFileCoordinatorWritingForReplacing: # Discussion
	NSFileCoordinatorWritingForReplacing NSFileCoordinatorWritingOptions = 8
)

func (e NSFileCoordinatorWritingOptions) String() string {
	switch e {
	case NSFileCoordinatorWritingContentIndependentMetadataOnly:
		return "NSFileCoordinatorWritingContentIndependentMetadataOnly"
	case NSFileCoordinatorWritingForDeleting:
		return "NSFileCoordinatorWritingForDeleting"
	case NSFileCoordinatorWritingForMerging:
		return "NSFileCoordinatorWritingForMerging"
	case NSFileCoordinatorWritingForMoving:
		return "NSFileCoordinatorWritingForMoving"
	case NSFileCoordinatorWritingForReplacing:
		return "NSFileCoordinatorWritingForReplacing"
	default:
		return fmt.Sprintf("NSFileCoordinatorWritingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions
type NSFileManagerItemReplacementOptions int

const (
	// NSFileManagerItemReplacementUsingNewMetadataOnly: Only metadata from the new item is used, and metadata from the original item isn’t preserved (default).
	NSFileManagerItemReplacementUsingNewMetadataOnly NSFileManagerItemReplacementOptions = 1
	// NSFileManagerItemReplacementWithoutDeletingBackupItem: The backup item remains in place after a successful replacement.
	NSFileManagerItemReplacementWithoutDeletingBackupItem NSFileManagerItemReplacementOptions = 2
)

func (e NSFileManagerItemReplacementOptions) String() string {
	switch e {
	case NSFileManagerItemReplacementUsingNewMetadataOnly:
		return "NSFileManagerItemReplacementUsingNewMetadataOnly"
	case NSFileManagerItemReplacementWithoutDeletingBackupItem:
		return "NSFileManagerItemReplacementWithoutDeletingBackupItem"
	default:
		return fmt.Sprintf("NSFileManagerItemReplacementOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileManagerResumeSyncBehavior
type NSFileManagerResumeSyncBehavior int

const (
	// NSFileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict: Resumes sync by first uploading the local version of the file, failing if the provider detects a conflict.
	NSFileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict NSFileManagerResumeSyncBehavior = 1
	// NSFileManagerResumeSyncBehaviorDropLocalChanges: Resumes synchronizing by overwriting any local changes with the remote version of the file.
	NSFileManagerResumeSyncBehaviorDropLocalChanges NSFileManagerResumeSyncBehavior = 2
	// NSFileManagerResumeSyncBehaviorPreserveLocalChanges: Resumes synchronizing by uploading the local version of the file.
	NSFileManagerResumeSyncBehaviorPreserveLocalChanges NSFileManagerResumeSyncBehavior = 0
)

func (e NSFileManagerResumeSyncBehavior) String() string {
	switch e {
	case NSFileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict:
		return "NSFileManagerResumeSyncBehaviorAfterUploadWithFailOnConflict"
	case NSFileManagerResumeSyncBehaviorDropLocalChanges:
		return "NSFileManagerResumeSyncBehaviorDropLocalChanges"
	case NSFileManagerResumeSyncBehaviorPreserveLocalChanges:
		return "NSFileManagerResumeSyncBehaviorPreserveLocalChanges"
	default:
		return fmt.Sprintf("NSFileManagerResumeSyncBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileManagerSupportedSyncControls
type NSFileManagerSupportedSyncControls int

const (
	// NSFileManagerSupportedSyncControlsFailUploadOnConflict: The file provider supports failing an upload if the local and server versions conflict.
	NSFileManagerSupportedSyncControlsFailUploadOnConflict NSFileManagerSupportedSyncControls = 2
	// NSFileManagerSupportedSyncControlsPauseSync: The file provider supports pausing the sync on the item.
	NSFileManagerSupportedSyncControlsPauseSync NSFileManagerSupportedSyncControls = 1
)

func (e NSFileManagerSupportedSyncControls) String() string {
	switch e {
	case NSFileManagerSupportedSyncControlsFailUploadOnConflict:
		return "NSFileManagerSupportedSyncControlsFailUploadOnConflict"
	case NSFileManagerSupportedSyncControlsPauseSync:
		return "NSFileManagerSupportedSyncControlsPauseSync"
	default:
		return fmt.Sprintf("NSFileManagerSupportedSyncControls(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/UnmountOptions
type NSFileManagerUnmountOptions int

const (
	// NSFileManagerUnmountAllPartitionsAndEjectDisk: Specifies that all partitions on an unmountable disk should be unmounted.
	NSFileManagerUnmountAllPartitionsAndEjectDisk NSFileManagerUnmountOptions = 1
	// NSFileManagerUnmountWithoutUI: Specifies that no UI should accompany the unmount operation.
	NSFileManagerUnmountWithoutUI NSFileManagerUnmountOptions = 2
)

func (e NSFileManagerUnmountOptions) String() string {
	switch e {
	case NSFileManagerUnmountAllPartitionsAndEjectDisk:
		return "NSFileManagerUnmountAllPartitionsAndEjectDisk"
	case NSFileManagerUnmountWithoutUI:
		return "NSFileManagerUnmountWithoutUI"
	default:
		return fmt.Sprintf("NSFileManagerUnmountOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileManagerUploadLocalVersionConflictPolicy
type NSFileManagerUploadLocalVersionConflictPolicy int

const (
	// NSFileManagerUploadConflictPolicyDefault: Resolves the conflict using the policy defined by the file provider.
	NSFileManagerUploadConflictPolicyDefault NSFileManagerUploadLocalVersionConflictPolicy = 0
	// NSFileManagerUploadConflictPolicyFailOnConflict: Resolves the conflict by causing the upload to fail.
	NSFileManagerUploadConflictPolicyFailOnConflict NSFileManagerUploadLocalVersionConflictPolicy = 1
)

func (e NSFileManagerUploadLocalVersionConflictPolicy) String() string {
	switch e {
	case NSFileManagerUploadConflictPolicyDefault:
		return "NSFileManagerUploadConflictPolicyDefault"
	case NSFileManagerUploadConflictPolicyFailOnConflict:
		return "NSFileManagerUploadConflictPolicyFailOnConflict"
	default:
		return fmt.Sprintf("NSFileManagerUploadLocalVersionConflictPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/AddingOptions
type NSFileVersionAddingOptions int

const (
	// NSFileVersionAddingByMoving: # Discussion
	NSFileVersionAddingByMoving NSFileVersionAddingOptions = 1
)

func (e NSFileVersionAddingOptions) String() string {
	switch e {
	case NSFileVersionAddingByMoving:
		return "NSFileVersionAddingByMoving"
	default:
		return fmt.Sprintf("NSFileVersionAddingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions
type NSFileVersionReplacingOptions int

const (
	// NSFileVersionReplacingByMoving: An option to perform replacing by moving a file.
	NSFileVersionReplacingByMoving NSFileVersionReplacingOptions = 1
)

func (e NSFileVersionReplacingOptions) String() string {
	switch e {
	case NSFileVersionReplacingByMoving:
		return "NSFileVersionReplacingByMoving"
	default:
		return fmt.Sprintf("NSFileVersionReplacingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileWrapper/ReadingOptions
type NSFileWrapperReadingOptions int

const (
	// NSFileWrapperReadingImmediate: The option to read files immediately after creating a file wrapper.
	NSFileWrapperReadingImmediate NSFileWrapperReadingOptions = 1
	// NSFileWrapperReadingWithoutMapping: Whether file mapping for regular file wrappers is disallowed.
	NSFileWrapperReadingWithoutMapping NSFileWrapperReadingOptions = 2
)

func (e NSFileWrapperReadingOptions) String() string {
	switch e {
	case NSFileWrapperReadingImmediate:
		return "NSFileWrapperReadingImmediate"
	case NSFileWrapperReadingWithoutMapping:
		return "NSFileWrapperReadingWithoutMapping"
	default:
		return fmt.Sprintf("NSFileWrapperReadingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileWrapper/WritingOptions
type NSFileWrapperWritingOptions int

const (
	// NSFileWrapperWritingAtomic: Whether writing is done atomically.
	NSFileWrapperWritingAtomic NSFileWrapperWritingOptions = 1
	// NSFileWrapperWritingWithNameUpdating: Whether descendant file wrappers’filename properties are set if the writing succeeds.
	NSFileWrapperWritingWithNameUpdating NSFileWrapperWritingOptions = 2
)

func (e NSFileWrapperWritingOptions) String() string {
	switch e {
	case NSFileWrapperWritingAtomic:
		return "NSFileWrapperWritingAtomic"
	case NSFileWrapperWritingWithNameUpdating:
		return "NSFileWrapperWritingWithNameUpdating"
	default:
		return fmt.Sprintf("NSFileWrapperWritingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/Formatter/Context
type NSFormattingContext int

const (
	// NSFormattingContextBeginningOfSentence: The formatting context for the beginning of a sentence.
	NSFormattingContextBeginningOfSentence NSFormattingContext = 4
	// NSFormattingContextDynamic: A formatting context determined automatically at runtime.
	NSFormattingContextDynamic NSFormattingContext = 1
	// NSFormattingContextListItem: The formatting context for a list or menu item.
	NSFormattingContextListItem NSFormattingContext = 3
	// NSFormattingContextMiddleOfSentence: The formatting context for the middle of a sentence.
	NSFormattingContextMiddleOfSentence NSFormattingContext = 5
	// NSFormattingContextStandalone: The formatting context for stand-alone usage.
	NSFormattingContextStandalone NSFormattingContext = 2
	// NSFormattingContextUnknown: An unknown formatting context.
	NSFormattingContextUnknown NSFormattingContext = 0
)

func (e NSFormattingContext) String() string {
	switch e {
	case NSFormattingContextBeginningOfSentence:
		return "NSFormattingContextBeginningOfSentence"
	case NSFormattingContextDynamic:
		return "NSFormattingContextDynamic"
	case NSFormattingContextListItem:
		return "NSFormattingContextListItem"
	case NSFormattingContextMiddleOfSentence:
		return "NSFormattingContextMiddleOfSentence"
	case NSFormattingContextStandalone:
		return "NSFormattingContextStandalone"
	case NSFormattingContextUnknown:
		return "NSFormattingContextUnknown"
	default:
		return fmt.Sprintf("NSFormattingContext(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle
type NSFormattingUnitStyle int

const (
	// NSFormattingUnitStyleLong: Specifies a long unit style.
	NSFormattingUnitStyleLong NSFormattingUnitStyle = 3
	// NSFormattingUnitStyleMedium: Specifies a medium unit style.
	NSFormattingUnitStyleMedium NSFormattingUnitStyle = 2
	// NSFormattingUnitStyleShort: Specifies a short unit style.
	NSFormattingUnitStyleShort NSFormattingUnitStyle = 1
)

func (e NSFormattingUnitStyle) String() string {
	switch e {
	case NSFormattingUnitStyleLong:
		return "NSFormattingUnitStyleLong"
	case NSFormattingUnitStyleMedium:
		return "NSFormattingUnitStyleMedium"
	case NSFormattingUnitStyleShort:
		return "NSFormattingUnitStyleShort"
	default:
		return fmt.Sprintf("NSFormattingUnitStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalCase
type NSGrammaticalCase int

const (
	NSGrammaticalCaseAblative NSGrammaticalCase = 6
	NSGrammaticalCaseAccusative NSGrammaticalCase = 2
	NSGrammaticalCaseAdessive NSGrammaticalCase = 7
	NSGrammaticalCaseAllative NSGrammaticalCase = 8
	NSGrammaticalCaseDative NSGrammaticalCase = 3
	NSGrammaticalCaseElative NSGrammaticalCase = 9
	NSGrammaticalCaseEssive NSGrammaticalCase = 11
	NSGrammaticalCaseGenitive NSGrammaticalCase = 4
	NSGrammaticalCaseInessive NSGrammaticalCase = 12
	NSGrammaticalCaseLocative NSGrammaticalCase = 13
	NSGrammaticalCaseNotSet NSGrammaticalCase = 0
	NSGrammaticalCasePrepositional NSGrammaticalCase = 5
	NSGrammaticalCaseTranslative NSGrammaticalCase = 14
)

func (e NSGrammaticalCase) String() string {
	switch e {
	case NSGrammaticalCaseAblative:
		return "NSGrammaticalCaseAblative"
	case NSGrammaticalCaseAccusative:
		return "NSGrammaticalCaseAccusative"
	case NSGrammaticalCaseAdessive:
		return "NSGrammaticalCaseAdessive"
	case NSGrammaticalCaseAllative:
		return "NSGrammaticalCaseAllative"
	case NSGrammaticalCaseDative:
		return "NSGrammaticalCaseDative"
	case NSGrammaticalCaseElative:
		return "NSGrammaticalCaseElative"
	case NSGrammaticalCaseEssive:
		return "NSGrammaticalCaseEssive"
	case NSGrammaticalCaseGenitive:
		return "NSGrammaticalCaseGenitive"
	case NSGrammaticalCaseInessive:
		return "NSGrammaticalCaseInessive"
	case NSGrammaticalCaseLocative:
		return "NSGrammaticalCaseLocative"
	case NSGrammaticalCaseNotSet:
		return "NSGrammaticalCaseNotSet"
	case NSGrammaticalCasePrepositional:
		return "NSGrammaticalCasePrepositional"
	case NSGrammaticalCaseTranslative:
		return "NSGrammaticalCaseTranslative"
	default:
		return fmt.Sprintf("NSGrammaticalCase(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalDefiniteness
type NSGrammaticalDefiniteness int

const (
	NSGrammaticalDefinitenessDefinite NSGrammaticalDefiniteness = 2
	NSGrammaticalDefinitenessIndefinite NSGrammaticalDefiniteness = 1
	NSGrammaticalDefinitenessNotSet NSGrammaticalDefiniteness = 0
)

func (e NSGrammaticalDefiniteness) String() string {
	switch e {
	case NSGrammaticalDefinitenessDefinite:
		return "NSGrammaticalDefinitenessDefinite"
	case NSGrammaticalDefinitenessIndefinite:
		return "NSGrammaticalDefinitenessIndefinite"
	case NSGrammaticalDefinitenessNotSet:
		return "NSGrammaticalDefinitenessNotSet"
	default:
		return fmt.Sprintf("NSGrammaticalDefiniteness(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalDetermination
type NSGrammaticalDetermination int

const (
	NSGrammaticalDeterminationDependent NSGrammaticalDetermination = 2
	NSGrammaticalDeterminationIndependent NSGrammaticalDetermination = 1
	NSGrammaticalDeterminationNotSet NSGrammaticalDetermination = 0
)

func (e NSGrammaticalDetermination) String() string {
	switch e {
	case NSGrammaticalDeterminationDependent:
		return "NSGrammaticalDeterminationDependent"
	case NSGrammaticalDeterminationIndependent:
		return "NSGrammaticalDeterminationIndependent"
	case NSGrammaticalDeterminationNotSet:
		return "NSGrammaticalDeterminationNotSet"
	default:
		return fmt.Sprintf("NSGrammaticalDetermination(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalGender
type NSGrammaticalGender int

const (
	// NSGrammaticalGenderFeminine: The feminine grammatical gender.
	NSGrammaticalGenderFeminine NSGrammaticalGender = 1
	// NSGrammaticalGenderMasculine: The masculine grammatical gender.
	NSGrammaticalGenderMasculine NSGrammaticalGender = 2
	// NSGrammaticalGenderNeuter: A value to not specify gender when inflecting a string.
	NSGrammaticalGenderNeuter NSGrammaticalGender = 3
	// NSGrammaticalGenderNotSet: A value that indicates the gender is unset.
	NSGrammaticalGenderNotSet NSGrammaticalGender = 0
)

func (e NSGrammaticalGender) String() string {
	switch e {
	case NSGrammaticalGenderFeminine:
		return "NSGrammaticalGenderFeminine"
	case NSGrammaticalGenderMasculine:
		return "NSGrammaticalGenderMasculine"
	case NSGrammaticalGenderNeuter:
		return "NSGrammaticalGenderNeuter"
	case NSGrammaticalGenderNotSet:
		return "NSGrammaticalGenderNotSet"
	default:
		return fmt.Sprintf("NSGrammaticalGender(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalNumber
type NSGrammaticalNumber int

const (
	// NSGrammaticalNumberNotSet: A value that indicates the number is unset.
	NSGrammaticalNumberNotSet NSGrammaticalNumber = 0
	// NSGrammaticalNumberPlural: Multiple persons or things, as used for a grammatical number.
	NSGrammaticalNumberPlural NSGrammaticalNumber = 3
	// NSGrammaticalNumberPluralFew: A small number of persons or things, as used for a grammatical number.
	NSGrammaticalNumberPluralFew NSGrammaticalNumber = 5
	// NSGrammaticalNumberPluralMany: A large number of persons or things, as used for a grammatical number.
	NSGrammaticalNumberPluralMany NSGrammaticalNumber = 6
	// NSGrammaticalNumberPluralTwo: Two persons or things, as used for a grammatical number.
	NSGrammaticalNumberPluralTwo NSGrammaticalNumber = 4
	// NSGrammaticalNumberSingular: A single person or thing, as used for a grammatical number.
	NSGrammaticalNumberSingular NSGrammaticalNumber = 1
	// NSGrammaticalNumberZero: Zero persons or things, as used for a grammatical number.
	NSGrammaticalNumberZero NSGrammaticalNumber = 2
)

func (e NSGrammaticalNumber) String() string {
	switch e {
	case NSGrammaticalNumberNotSet:
		return "NSGrammaticalNumberNotSet"
	case NSGrammaticalNumberPlural:
		return "NSGrammaticalNumberPlural"
	case NSGrammaticalNumberPluralFew:
		return "NSGrammaticalNumberPluralFew"
	case NSGrammaticalNumberPluralMany:
		return "NSGrammaticalNumberPluralMany"
	case NSGrammaticalNumberPluralTwo:
		return "NSGrammaticalNumberPluralTwo"
	case NSGrammaticalNumberSingular:
		return "NSGrammaticalNumberSingular"
	case NSGrammaticalNumberZero:
		return "NSGrammaticalNumberZero"
	default:
		return fmt.Sprintf("NSGrammaticalNumber(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalPartOfSpeech
type NSGrammaticalPartOfSpeech int

const (
	// NSGrammaticalPartOfSpeechAbbreviation: An abbreviation, as used as a part of speech.
	NSGrammaticalPartOfSpeechAbbreviation NSGrammaticalPartOfSpeech = 14
	// NSGrammaticalPartOfSpeechAdjective: An adjective, as used as a part of speech.
	NSGrammaticalPartOfSpeechAdjective NSGrammaticalPartOfSpeech = 6
	// NSGrammaticalPartOfSpeechAdposition: An adposition, as used as a part of speech.
	NSGrammaticalPartOfSpeechAdposition NSGrammaticalPartOfSpeech = 7
	// NSGrammaticalPartOfSpeechAdverb: An adverb, as used as a part of speech.
	NSGrammaticalPartOfSpeechAdverb NSGrammaticalPartOfSpeech = 4
	// NSGrammaticalPartOfSpeechConjunction: A conjunction, as used as a part of speech.
	NSGrammaticalPartOfSpeechConjunction NSGrammaticalPartOfSpeech = 10
	// NSGrammaticalPartOfSpeechDeterminer: A determiner, as used as a part of speech.
	NSGrammaticalPartOfSpeechDeterminer NSGrammaticalPartOfSpeech = 1
	// NSGrammaticalPartOfSpeechInterjection: An interjection, as used as a part of speech.
	NSGrammaticalPartOfSpeechInterjection NSGrammaticalPartOfSpeech = 12
	// NSGrammaticalPartOfSpeechLetter: A letter, as used as a part of speech.
	NSGrammaticalPartOfSpeechLetter NSGrammaticalPartOfSpeech = 3
	// NSGrammaticalPartOfSpeechNotSet: A value that indicates the part of speech is unset.
	NSGrammaticalPartOfSpeechNotSet NSGrammaticalPartOfSpeech = 0
	// NSGrammaticalPartOfSpeechNoun: A noun, as used as a part of speech.
	NSGrammaticalPartOfSpeechNoun NSGrammaticalPartOfSpeech = 9
	// NSGrammaticalPartOfSpeechNumeral: A numeral, as used as a part of speech.
	NSGrammaticalPartOfSpeechNumeral NSGrammaticalPartOfSpeech = 11
	// NSGrammaticalPartOfSpeechParticle: A particle, as used as a part of speech.
	NSGrammaticalPartOfSpeechParticle NSGrammaticalPartOfSpeech = 5
	// NSGrammaticalPartOfSpeechPreposition: A preposition, as used as a part of speech.
	NSGrammaticalPartOfSpeechPreposition NSGrammaticalPartOfSpeech = 13
	// NSGrammaticalPartOfSpeechPronoun: A pronoun, as used as a part of speech.
	NSGrammaticalPartOfSpeechPronoun NSGrammaticalPartOfSpeech = 2
	// NSGrammaticalPartOfSpeechVerb: A verb, as used as a part of speech.
	NSGrammaticalPartOfSpeechVerb NSGrammaticalPartOfSpeech = 8
)

func (e NSGrammaticalPartOfSpeech) String() string {
	switch e {
	case NSGrammaticalPartOfSpeechAbbreviation:
		return "NSGrammaticalPartOfSpeechAbbreviation"
	case NSGrammaticalPartOfSpeechAdjective:
		return "NSGrammaticalPartOfSpeechAdjective"
	case NSGrammaticalPartOfSpeechAdposition:
		return "NSGrammaticalPartOfSpeechAdposition"
	case NSGrammaticalPartOfSpeechAdverb:
		return "NSGrammaticalPartOfSpeechAdverb"
	case NSGrammaticalPartOfSpeechConjunction:
		return "NSGrammaticalPartOfSpeechConjunction"
	case NSGrammaticalPartOfSpeechDeterminer:
		return "NSGrammaticalPartOfSpeechDeterminer"
	case NSGrammaticalPartOfSpeechInterjection:
		return "NSGrammaticalPartOfSpeechInterjection"
	case NSGrammaticalPartOfSpeechLetter:
		return "NSGrammaticalPartOfSpeechLetter"
	case NSGrammaticalPartOfSpeechNotSet:
		return "NSGrammaticalPartOfSpeechNotSet"
	case NSGrammaticalPartOfSpeechNoun:
		return "NSGrammaticalPartOfSpeechNoun"
	case NSGrammaticalPartOfSpeechNumeral:
		return "NSGrammaticalPartOfSpeechNumeral"
	case NSGrammaticalPartOfSpeechParticle:
		return "NSGrammaticalPartOfSpeechParticle"
	case NSGrammaticalPartOfSpeechPreposition:
		return "NSGrammaticalPartOfSpeechPreposition"
	case NSGrammaticalPartOfSpeechPronoun:
		return "NSGrammaticalPartOfSpeechPronoun"
	case NSGrammaticalPartOfSpeechVerb:
		return "NSGrammaticalPartOfSpeechVerb"
	default:
		return fmt.Sprintf("NSGrammaticalPartOfSpeech(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalPerson
type NSGrammaticalPerson int

const (
	NSGrammaticalPersonFirst NSGrammaticalPerson = 1
	NSGrammaticalPersonNotSet NSGrammaticalPerson = 0
	NSGrammaticalPersonSecond NSGrammaticalPerson = 2
	NSGrammaticalPersonThird NSGrammaticalPerson = 3
)

func (e NSGrammaticalPerson) String() string {
	switch e {
	case NSGrammaticalPersonFirst:
		return "NSGrammaticalPersonFirst"
	case NSGrammaticalPersonNotSet:
		return "NSGrammaticalPersonNotSet"
	case NSGrammaticalPersonSecond:
		return "NSGrammaticalPersonSecond"
	case NSGrammaticalPersonThird:
		return "NSGrammaticalPersonThird"
	default:
		return fmt.Sprintf("NSGrammaticalPerson(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSGrammaticalPronounType
type NSGrammaticalPronounType int

const (
	NSGrammaticalPronounTypeNotSet NSGrammaticalPronounType = 0
	NSGrammaticalPronounTypePersonal NSGrammaticalPronounType = 1
	NSGrammaticalPronounTypePossessive NSGrammaticalPronounType = 3
	NSGrammaticalPronounTypeReflexive NSGrammaticalPronounType = 2
)

func (e NSGrammaticalPronounType) String() string {
	switch e {
	case NSGrammaticalPronounTypeNotSet:
		return "NSGrammaticalPronounTypeNotSet"
	case NSGrammaticalPronounTypePersonal:
		return "NSGrammaticalPronounTypePersonal"
	case NSGrammaticalPronounTypePossessive:
		return "NSGrammaticalPronounTypePossessive"
	case NSGrammaticalPronounTypeReflexive:
		return "NSGrammaticalPronounTypeReflexive"
	default:
		return fmt.Sprintf("NSGrammaticalPronounType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/HTTPCookie/AcceptPolicy
type NSHTTPCookieAcceptPolicy int

const (
	// NSHTTPCookieAcceptPolicyAlways: Accept all cookies.
	NSHTTPCookieAcceptPolicyAlways NSHTTPCookieAcceptPolicy = 0
	// NSHTTPCookieAcceptPolicyNever: Reject all cookies.
	NSHTTPCookieAcceptPolicyNever NSHTTPCookieAcceptPolicy = 1
	// NSHTTPCookieAcceptPolicyOnlyFromMainDocumentDomain: Accept cookies only from the main document domain.
	NSHTTPCookieAcceptPolicyOnlyFromMainDocumentDomain NSHTTPCookieAcceptPolicy = 2
)

func (e NSHTTPCookieAcceptPolicy) String() string {
	switch e {
	case NSHTTPCookieAcceptPolicyAlways:
		return "NSHTTPCookieAcceptPolicyAlways"
	case NSHTTPCookieAcceptPolicyNever:
		return "NSHTTPCookieAcceptPolicyNever"
	case NSHTTPCookieAcceptPolicyOnlyFromMainDocumentDomain:
		return "NSHTTPCookieAcceptPolicyOnlyFromMainDocumentDomain"
	default:
		return fmt.Sprintf("NSHTTPCookieAcceptPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options
type NSISO8601DateFormatOptions int

const (
	// NSISO8601DateFormatWithColonSeparatorInTime: # Discussion
	NSISO8601DateFormatWithColonSeparatorInTime NSISO8601DateFormatOptions = 512
	// NSISO8601DateFormatWithColonSeparatorInTimeZone: # Discussion
	NSISO8601DateFormatWithColonSeparatorInTimeZone NSISO8601DateFormatOptions = 1024
	// NSISO8601DateFormatWithDashSeparatorInDate: # Discussion
	NSISO8601DateFormatWithDashSeparatorInDate NSISO8601DateFormatOptions = 256
	// NSISO8601DateFormatWithDay: # Discussion
	NSISO8601DateFormatWithDay NSISO8601DateFormatOptions = 16
	NSISO8601DateFormatWithFractionalSeconds NSISO8601DateFormatOptions = 2048
	// NSISO8601DateFormatWithFullDate: # Discussion
	NSISO8601DateFormatWithFullDate NSISO8601DateFormatOptions = 1
	// NSISO8601DateFormatWithFullTime: # Discussion
	NSISO8601DateFormatWithFullTime NSISO8601DateFormatOptions = 32
	// NSISO8601DateFormatWithInternetDateTime: # Discussion
	NSISO8601DateFormatWithInternetDateTime NSISO8601DateFormatOptions = 1
	// NSISO8601DateFormatWithMonth: # Discussion
	NSISO8601DateFormatWithMonth NSISO8601DateFormatOptions = 2
	// NSISO8601DateFormatWithSpaceBetweenDateAndTime: # Discussion
	NSISO8601DateFormatWithSpaceBetweenDateAndTime NSISO8601DateFormatOptions = 128
	// NSISO8601DateFormatWithTime: # Discussion
	NSISO8601DateFormatWithTime NSISO8601DateFormatOptions = 32
	// NSISO8601DateFormatWithTimeZone: # Discussion
	NSISO8601DateFormatWithTimeZone NSISO8601DateFormatOptions = 64
	// NSISO8601DateFormatWithWeekOfYear: # Discussion
	NSISO8601DateFormatWithWeekOfYear NSISO8601DateFormatOptions = 4
	// NSISO8601DateFormatWithYear: # Discussion
	NSISO8601DateFormatWithYear NSISO8601DateFormatOptions = 1
)

func (e NSISO8601DateFormatOptions) String() string {
	switch e {
	case NSISO8601DateFormatWithColonSeparatorInTime:
		return "NSISO8601DateFormatWithColonSeparatorInTime"
	case NSISO8601DateFormatWithColonSeparatorInTimeZone:
		return "NSISO8601DateFormatWithColonSeparatorInTimeZone"
	case NSISO8601DateFormatWithDashSeparatorInDate:
		return "NSISO8601DateFormatWithDashSeparatorInDate"
	case NSISO8601DateFormatWithDay:
		return "NSISO8601DateFormatWithDay"
	case NSISO8601DateFormatWithFractionalSeconds:
		return "NSISO8601DateFormatWithFractionalSeconds"
	case NSISO8601DateFormatWithFullDate:
		return "NSISO8601DateFormatWithFullDate"
	case NSISO8601DateFormatWithFullTime:
		return "NSISO8601DateFormatWithFullTime"
	case NSISO8601DateFormatWithMonth:
		return "NSISO8601DateFormatWithMonth"
	case NSISO8601DateFormatWithSpaceBetweenDateAndTime:
		return "NSISO8601DateFormatWithSpaceBetweenDateAndTime"
	case NSISO8601DateFormatWithTimeZone:
		return "NSISO8601DateFormatWithTimeZone"
	case NSISO8601DateFormatWithWeekOfYear:
		return "NSISO8601DateFormatWithWeekOfYear"
	default:
		return fmt.Sprintf("NSISO8601DateFormatOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/InlinePresentationIntent
type InlinePresentationIntent int

const (
	// NSInlinePresentationIntentBlockHTML: An intent that represents a block HTML presentation.
	NSInlinePresentationIntentBlockHTML InlinePresentationIntent = 512
	// NSInlinePresentationIntentCode: An intent that represents a code voice presentation.
	NSInlinePresentationIntentCode InlinePresentationIntent = 4
	// NSInlinePresentationIntentEmphasized: An intent that represents an emphasized presentation.
	NSInlinePresentationIntentEmphasized InlinePresentationIntent = 1
	// NSInlinePresentationIntentInlineHTML: An intent that represents an inline HTML presentation.
	NSInlinePresentationIntentInlineHTML InlinePresentationIntent = 256
	// NSInlinePresentationIntentLineBreak: An intent that represents a line break.
	NSInlinePresentationIntentLineBreak InlinePresentationIntent = 128
	// NSInlinePresentationIntentSoftBreak: An intent that represents a soft line break.
	NSInlinePresentationIntentSoftBreak InlinePresentationIntent = 64
	// NSInlinePresentationIntentStrikethrough: An intent that represents a strikethrough presentation.
	NSInlinePresentationIntentStrikethrough InlinePresentationIntent = 32
	// NSInlinePresentationIntentStronglyEmphasized: An intent that represents a strongly emphasized presentation.
	NSInlinePresentationIntentStronglyEmphasized InlinePresentationIntent = 2
)

func (e InlinePresentationIntent) String() string {
	switch e {
	case NSInlinePresentationIntentBlockHTML:
		return "NSInlinePresentationIntentBlockHTML"
	case NSInlinePresentationIntentCode:
		return "NSInlinePresentationIntentCode"
	case NSInlinePresentationIntentEmphasized:
		return "NSInlinePresentationIntentEmphasized"
	case NSInlinePresentationIntentInlineHTML:
		return "NSInlinePresentationIntentInlineHTML"
	case NSInlinePresentationIntentLineBreak:
		return "NSInlinePresentationIntentLineBreak"
	case NSInlinePresentationIntentSoftBreak:
		return "NSInlinePresentationIntentSoftBreak"
	case NSInlinePresentationIntentStrikethrough:
		return "NSInlinePresentationIntentStrikethrough"
	case NSInlinePresentationIntentStronglyEmphasized:
		return "NSInlinePresentationIntentStronglyEmphasized"
	default:
		return fmt.Sprintf("InlinePresentationIntent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSPositionalSpecifier/InsertionPosition
type NSInsertionPosition int

const (
	// NSPositionAfter: Specifies a position after another object.
	NSPositionAfter NSInsertionPosition = 0
	// NSPositionBefore: Specifies a position before another object.
	NSPositionBefore NSInsertionPosition = 1
	// NSPositionBeginning: Specifies a position at the beginning of a collection.
	NSPositionBeginning NSInsertionPosition = 2
	// NSPositionEnd: Specifies a position at the end of a collection.
	NSPositionEnd NSInsertionPosition = 3
	// NSPositionReplace: Specifies a position in the place of another object.
	NSPositionReplace NSInsertionPosition = 4
)

func (e NSInsertionPosition) String() string {
	switch e {
	case NSPositionAfter:
		return "NSPositionAfter"
	case NSPositionBefore:
		return "NSPositionBefore"
	case NSPositionBeginning:
		return "NSPositionBeginning"
	case NSPositionEnd:
		return "NSPositionEnd"
	case NSPositionReplace:
		return "NSPositionReplace"
	default:
		return fmt.Sprintf("NSInsertionPosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/ErrorCode
type NSItemProviderErrorCode int

const (
	// NSItemProviderItemUnavailableError: An error code indicating that the requested data was unavailable from an item provider.
	NSItemProviderItemUnavailableError NSItemProviderErrorCode = -1000
	// NSItemProviderUnavailableCoercionError: An error code indicating that the requested data type coercion is unavailable from an item provider.
	NSItemProviderUnavailableCoercionError NSItemProviderErrorCode = -1200
	// NSItemProviderUnexpectedValueClassError: An error code indicating that type coercion to the requested class failed.
	NSItemProviderUnexpectedValueClassError NSItemProviderErrorCode = -1100
	// NSItemProviderUnknownError: An error code indicating an unknown error with consuming data from an item provider.
	NSItemProviderUnknownError NSItemProviderErrorCode = -1
)

func (e NSItemProviderErrorCode) String() string {
	switch e {
	case NSItemProviderItemUnavailableError:
		return "NSItemProviderItemUnavailableError"
	case NSItemProviderUnavailableCoercionError:
		return "NSItemProviderUnavailableCoercionError"
	case NSItemProviderUnexpectedValueClassError:
		return "NSItemProviderUnexpectedValueClassError"
	case NSItemProviderUnknownError:
		return "NSItemProviderUnknownError"
	default:
		return fmt.Sprintf("NSItemProviderErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSItemProviderFileOptions
type NSItemProviderFileOptions int

const (
	// NSItemProviderFileOptionOpenInPlace: A data-access specification declaring that items should open in place, rather than being copied.
	NSItemProviderFileOptionOpenInPlace NSItemProviderFileOptions = 1
)

func (e NSItemProviderFileOptions) String() string {
	switch e {
	case NSItemProviderFileOptionOpenInPlace:
		return "NSItemProviderFileOptionOpenInPlace"
	default:
		return fmt.Sprintf("NSItemProviderFileOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSItemProviderRepresentationVisibility
type NSItemProviderRepresentationVisibility int

const (
	// NSItemProviderRepresentationVisibilityAll: A representation visibility specification conferring item visibility to all processes.
	NSItemProviderRepresentationVisibilityAll NSItemProviderRepresentationVisibility = 0
	// NSItemProviderRepresentationVisibilityGroup: A representation visibility specification confining item visibility to the app’s app group.
	NSItemProviderRepresentationVisibilityGroup NSItemProviderRepresentationVisibility = 2
	// NSItemProviderRepresentationVisibilityOwnProcess: A representation visibility specification confining item visibility to the app that is the source of the item.
	NSItemProviderRepresentationVisibilityOwnProcess NSItemProviderRepresentationVisibility = 3
	// NSItemProviderRepresentationVisibilityTeam: A representation visibility specification confining item visibility to processes created by the app’s development team.
	NSItemProviderRepresentationVisibilityTeam NSItemProviderRepresentationVisibility = 1
)

func (e NSItemProviderRepresentationVisibility) String() string {
	switch e {
	case NSItemProviderRepresentationVisibilityAll:
		return "NSItemProviderRepresentationVisibilityAll"
	case NSItemProviderRepresentationVisibilityGroup:
		return "NSItemProviderRepresentationVisibilityGroup"
	case NSItemProviderRepresentationVisibilityOwnProcess:
		return "NSItemProviderRepresentationVisibilityOwnProcess"
	case NSItemProviderRepresentationVisibilityTeam:
		return "NSItemProviderRepresentationVisibilityTeam"
	default:
		return fmt.Sprintf("NSItemProviderRepresentationVisibility(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/ReadingOptions
type NSJSONReadingOptions int

const (
	// NSJSONReadingFragmentsAllowed: Specifies that the parser allows top-level objects that aren’t arrays or dictionaries.
	NSJSONReadingFragmentsAllowed NSJSONReadingOptions = 4
	// NSJSONReadingJSON5Allowed: Specifies that reading serialized JSON data supports the JSON5 syntax.
	NSJSONReadingJSON5Allowed NSJSONReadingOptions = 8
	// NSJSONReadingMutableContainers: Specifies that arrays and dictionaries in the returned object are mutable.
	NSJSONReadingMutableContainers NSJSONReadingOptions = 1
	// NSJSONReadingMutableLeaves: Specifies that leaf strings in the JSON object graph are mutable.
	NSJSONReadingMutableLeaves NSJSONReadingOptions = 2
	// NSJSONReadingTopLevelDictionaryAssumed: Specifies that the parser assumes the top level of the JSON data is a dictionary, even if it doesn’t begin and end with curly braces.
	NSJSONReadingTopLevelDictionaryAssumed NSJSONReadingOptions = 16
	// Deprecated.
	NSJSONReadingAllowFragments NSJSONReadingOptions = 4
)

func (e NSJSONReadingOptions) String() string {
	switch e {
	case NSJSONReadingFragmentsAllowed:
		return "NSJSONReadingFragmentsAllowed"
	case NSJSONReadingJSON5Allowed:
		return "NSJSONReadingJSON5Allowed"
	case NSJSONReadingMutableContainers:
		return "NSJSONReadingMutableContainers"
	case NSJSONReadingMutableLeaves:
		return "NSJSONReadingMutableLeaves"
	case NSJSONReadingTopLevelDictionaryAssumed:
		return "NSJSONReadingTopLevelDictionaryAssumed"
	default:
		return fmt.Sprintf("NSJSONReadingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/JSONSerialization/WritingOptions
type NSJSONWritingOptions int

const (
	// NSJSONWritingFragmentsAllowed: Specifies that the parser should allow top-level objects that aren’t arrays or dictionaries.
	NSJSONWritingFragmentsAllowed NSJSONWritingOptions = 4
	// NSJSONWritingPrettyPrinted: Specifies that the output uses white space and indentation to make the resulting data more readable.
	NSJSONWritingPrettyPrinted NSJSONWritingOptions = 1
	// NSJSONWritingSortedKeys: Specifies that the output sorts keys in lexicographic order.
	NSJSONWritingSortedKeys NSJSONWritingOptions = 2
	// NSJSONWritingWithoutEscapingSlashes: Specifies that the output doesn’t prefix slash characters with escape characters.
	NSJSONWritingWithoutEscapingSlashes NSJSONWritingOptions = 8
)

func (e NSJSONWritingOptions) String() string {
	switch e {
	case NSJSONWritingFragmentsAllowed:
		return "NSJSONWritingFragmentsAllowed"
	case NSJSONWritingPrettyPrinted:
		return "NSJSONWritingPrettyPrinted"
	case NSJSONWritingSortedKeys:
		return "NSJSONWritingSortedKeys"
	case NSJSONWritingWithoutEscapingSlashes:
		return "NSJSONWritingWithoutEscapingSlashes"
	default:
		return fmt.Sprintf("NSJSONWritingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSKeyValueChange
type NSKeyValueChange uint

const (
	// NSKeyValueChangeReplacement: Indicates that an object has been replaced in the to-many relationship that is being observed.
	NSKeyValueChangeReplacement NSKeyValueChange = 4
	// NSKeyValueChangeSetting: Indicates that the value of the observed key path was set to a new value.
	NSKeyValueChangeSetting NSKeyValueChange = 1
)

func (e NSKeyValueChange) String() string {
	switch e {
	case NSKeyValueChangeReplacement:
		return "NSKeyValueChangeReplacement"
	case NSKeyValueChangeSetting:
		return "NSKeyValueChangeSetting"
	default:
		return fmt.Sprintf("NSKeyValueChange(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSKeyValueObservingOptions
type NSKeyValueObservingOptions uint

const (
	// NSKeyValueObservingOptionInitial: If specified, a notification should be sent to the observer immediately, before the observer registration method even returns.
	NSKeyValueObservingOptionInitial NSKeyValueObservingOptions = 4
	// NSKeyValueObservingOptionNew: Indicates that the change dictionary should provide the new attribute value, if applicable.
	NSKeyValueObservingOptionNew NSKeyValueObservingOptions = 1
	// NSKeyValueObservingOptionOld: Indicates that the change dictionary should contain the old attribute value, if applicable.
	NSKeyValueObservingOptionOld NSKeyValueObservingOptions = 2
	// NSKeyValueObservingOptionPrior: Whether separate notifications should be sent to the observer before and after each change, instead of a single notification after the change.
	NSKeyValueObservingOptionPrior NSKeyValueObservingOptions = 8
)

func (e NSKeyValueObservingOptions) String() string {
	switch e {
	case NSKeyValueObservingOptionInitial:
		return "NSKeyValueObservingOptionInitial"
	case NSKeyValueObservingOptionNew:
		return "NSKeyValueObservingOptionNew"
	case NSKeyValueObservingOptionOld:
		return "NSKeyValueObservingOptionOld"
	case NSKeyValueObservingOptionPrior:
		return "NSKeyValueObservingOptionPrior"
	default:
		return fmt.Sprintf("NSKeyValueObservingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSKeyValueSetMutationKind
type NSKeyValueSetMutationKind uint

const (
	NSKeyValueIntersectSetMutation NSKeyValueSetMutationKind = 3
	NSKeyValueMinusSetMutation NSKeyValueSetMutationKind = 2
	NSKeyValueSetSetMutation NSKeyValueSetMutationKind = 4
	NSKeyValueUnionSetMutation NSKeyValueSetMutationKind = 1
)

func (e NSKeyValueSetMutationKind) String() string {
	switch e {
	case NSKeyValueIntersectSetMutation:
		return "NSKeyValueIntersectSetMutation"
	case NSKeyValueMinusSetMutation:
		return "NSKeyValueMinusSetMutation"
	case NSKeyValueSetSetMutation:
		return "NSKeyValueSetSetMutation"
	case NSKeyValueUnionSetMutation:
		return "NSKeyValueUnionSetMutation"
	default:
		return fmt.Sprintf("NSKeyValueSetMutationKind(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit
type NSLengthFormatterUnit int

const (
	// NSLengthFormatterUnitCentimeter: The centimeter unit.
	NSLengthFormatterUnitCentimeter NSLengthFormatterUnit = 9
	// NSLengthFormatterUnitFoot: The foot unit.
	NSLengthFormatterUnitFoot NSLengthFormatterUnit = 1280
	// NSLengthFormatterUnitInch: The inch unit.
	NSLengthFormatterUnitInch NSLengthFormatterUnit = 1280
	// NSLengthFormatterUnitKilometer: The kilometer unit.
	NSLengthFormatterUnitKilometer NSLengthFormatterUnit = 14
	// NSLengthFormatterUnitMeter: The meter unit.
	NSLengthFormatterUnitMeter NSLengthFormatterUnit = 11
	// NSLengthFormatterUnitMile: The mile unit.
	NSLengthFormatterUnitMile NSLengthFormatterUnit = 1280
	// NSLengthFormatterUnitMillimeter: The millimeter unit.
	NSLengthFormatterUnitMillimeter NSLengthFormatterUnit = 8
	// NSLengthFormatterUnitYard: The yard unit.
	NSLengthFormatterUnitYard NSLengthFormatterUnit = 1280
)

func (e NSLengthFormatterUnit) String() string {
	switch e {
	case NSLengthFormatterUnitCentimeter:
		return "NSLengthFormatterUnitCentimeter"
	case NSLengthFormatterUnitFoot:
		return "NSLengthFormatterUnitFoot"
	case NSLengthFormatterUnitKilometer:
		return "NSLengthFormatterUnitKilometer"
	case NSLengthFormatterUnitMeter:
		return "NSLengthFormatterUnitMeter"
	case NSLengthFormatterUnitMillimeter:
		return "NSLengthFormatterUnitMillimeter"
	default:
		return fmt.Sprintf("NSLengthFormatterUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTagger/Options
type NSLinguisticTaggerOptions int

const (
	// NSLinguisticTaggerJoinNames: Typically, multiple-word names will be returned as multiple tokens, following the standard tokenization practice of the tagger.
	NSLinguisticTaggerJoinNames NSLinguisticTaggerOptions = 16
	// NSLinguisticTaggerOmitOther: Omit tokens of type other (non-linguistic items, such as symbols).
	NSLinguisticTaggerOmitOther NSLinguisticTaggerOptions = 8
	// NSLinguisticTaggerOmitPunctuation: Omit tokens of type punctuation (all punctuation).
	NSLinguisticTaggerOmitPunctuation NSLinguisticTaggerOptions = 2
	// NSLinguisticTaggerOmitWhitespace: Omit tokens of type whitespace (whitespace of all sorts).
	NSLinguisticTaggerOmitWhitespace NSLinguisticTaggerOptions = 4
	// NSLinguisticTaggerOmitWords: Omit tokens of type word (items considered to be words).
	NSLinguisticTaggerOmitWords NSLinguisticTaggerOptions = 1
)

func (e NSLinguisticTaggerOptions) String() string {
	switch e {
	case NSLinguisticTaggerJoinNames:
		return "NSLinguisticTaggerJoinNames"
	case NSLinguisticTaggerOmitOther:
		return "NSLinguisticTaggerOmitOther"
	case NSLinguisticTaggerOmitPunctuation:
		return "NSLinguisticTaggerOmitPunctuation"
	case NSLinguisticTaggerOmitWhitespace:
		return "NSLinguisticTaggerOmitWhitespace"
	case NSLinguisticTaggerOmitWords:
		return "NSLinguisticTaggerOmitWords"
	default:
		return fmt.Sprintf("NSLinguisticTaggerOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSLinguisticTaggerUnit
type NSLinguisticTaggerUnit int

const (
	// NSLinguisticTaggerUnitDocument: The document in its entirety.
	NSLinguisticTaggerUnitDocument NSLinguisticTaggerUnit = 3
	// NSLinguisticTaggerUnitParagraph: An individual paragraph.
	NSLinguisticTaggerUnitParagraph NSLinguisticTaggerUnit = 2
	// NSLinguisticTaggerUnitSentence: An individual sentence.
	NSLinguisticTaggerUnitSentence NSLinguisticTaggerUnit = 1
	// NSLinguisticTaggerUnitWord: An individual word.
	NSLinguisticTaggerUnitWord NSLinguisticTaggerUnit = 0
)

func (e NSLinguisticTaggerUnit) String() string {
	switch e {
	case NSLinguisticTaggerUnitDocument:
		return "NSLinguisticTaggerUnitDocument"
	case NSLinguisticTaggerUnitParagraph:
		return "NSLinguisticTaggerUnitParagraph"
	case NSLinguisticTaggerUnitSentence:
		return "NSLinguisticTaggerUnitSentence"
	case NSLinguisticTaggerUnitWord:
		return "NSLinguisticTaggerUnitWord"
	default:
		return fmt.Sprintf("NSLinguisticTaggerUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSLocale/LanguageDirection
type NSLocaleLanguageDirection int

const (
	// NSLocaleLanguageDirectionBottomToTop: The language direction is from bottom to top.
	NSLocaleLanguageDirectionBottomToTop NSLocaleLanguageDirection = 0
	// NSLocaleLanguageDirectionLeftToRight: The language direction is from left to right.
	NSLocaleLanguageDirectionLeftToRight NSLocaleLanguageDirection = 1
	// NSLocaleLanguageDirectionRightToLeft: The language direction is from right to left.
	NSLocaleLanguageDirectionRightToLeft NSLocaleLanguageDirection = 2
	// NSLocaleLanguageDirectionTopToBottom: The language direction is from top to bottom.
	NSLocaleLanguageDirectionTopToBottom NSLocaleLanguageDirection = 3
	// NSLocaleLanguageDirectionUnknown: The direction of the language is unknown.
	NSLocaleLanguageDirectionUnknown NSLocaleLanguageDirection = 0
)

func (e NSLocaleLanguageDirection) String() string {
	switch e {
	case NSLocaleLanguageDirectionBottomToTop:
		return "NSLocaleLanguageDirectionBottomToTop"
	case NSLocaleLanguageDirectionLeftToRight:
		return "NSLocaleLanguageDirectionLeftToRight"
	case NSLocaleLanguageDirectionRightToLeft:
		return "NSLocaleLanguageDirectionRightToLeft"
	case NSLocaleLanguageDirectionTopToBottom:
		return "NSLocaleLanguageDirectionTopToBottom"
	default:
		return fmt.Sprintf("NSLocaleLanguageDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSMachPort/Options
type NSMachPortOptions int

const (
	// NSMachPortDeallocateReceiveRight: Remove a receive right when the [NSMachPort] object is invalidated or destroyed.
	NSMachPortDeallocateReceiveRight NSMachPortOptions = 2
	// NSMachPortDeallocateSendRight: Deallocate a send right when the [NSMachPort] object is invalidated or destroyed.
	NSMachPortDeallocateSendRight NSMachPortOptions = 1
)

func (e NSMachPortOptions) String() string {
	switch e {
	case NSMachPortDeallocateReceiveRight:
		return "NSMachPortDeallocateReceiveRight"
	case NSMachPortDeallocateSendRight:
		return "NSMachPortDeallocateSendRight"
	default:
		return fmt.Sprintf("NSMachPortOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit
type NSMassFormatterUnit int

const (
	// NSMassFormatterUnitGram: The gram unit.
	NSMassFormatterUnitGram NSMassFormatterUnit = 11
	// NSMassFormatterUnitKilogram: The kilogram unit.
	NSMassFormatterUnitKilogram NSMassFormatterUnit = 14
	// NSMassFormatterUnitOunce: The ounce unit.
	NSMassFormatterUnitOunce NSMassFormatterUnit = 1536
	// NSMassFormatterUnitPound: The pound unit.
	NSMassFormatterUnitPound NSMassFormatterUnit = 1536
	// NSMassFormatterUnitStone: The stone unit.
	NSMassFormatterUnitStone NSMassFormatterUnit = 1536
)

func (e NSMassFormatterUnit) String() string {
	switch e {
	case NSMassFormatterUnitGram:
		return "NSMassFormatterUnitGram"
	case NSMassFormatterUnitKilogram:
		return "NSMassFormatterUnitKilogram"
	case NSMassFormatterUnitOunce:
		return "NSMassFormatterUnitOunce"
	default:
		return fmt.Sprintf("NSMassFormatterUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingFlags
type NSMatchingFlags int

const (
	// NSMatchingCompleted: Set when the Block is called after matching has completed.
	NSMatchingCompleted NSMatchingFlags = 2
	// NSMatchingHitEnd: Set when the current match operation reached the end of the search range.
	NSMatchingHitEnd NSMatchingFlags = 4
	// NSMatchingInternalError: Set when matching failed due to an internal error.
	NSMatchingInternalError NSMatchingFlags = 16
	// NSMatchingProgress: Set when the Block is called to report progress during a long-running match operation.
	NSMatchingProgress NSMatchingFlags = 1
	// NSMatchingRequiredEnd: Set when the current match depended on the location of the end of the search range.
	NSMatchingRequiredEnd NSMatchingFlags = 8
)

func (e NSMatchingFlags) String() string {
	switch e {
	case NSMatchingCompleted:
		return "NSMatchingCompleted"
	case NSMatchingHitEnd:
		return "NSMatchingHitEnd"
	case NSMatchingInternalError:
		return "NSMatchingInternalError"
	case NSMatchingProgress:
		return "NSMatchingProgress"
	case NSMatchingRequiredEnd:
		return "NSMatchingRequiredEnd"
	default:
		return fmt.Sprintf("NSMatchingFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/MatchingOptions
type NSMatchingOptions int

const (
	// NSMatchingAnchored: Specifies that matches are limited to those at the start of the search range.
	NSMatchingAnchored NSMatchingOptions = 4
	// NSMatchingReportCompletion: Call the Block once after the completion of any matching.
	NSMatchingReportCompletion NSMatchingOptions = 2
	// NSMatchingReportProgress: Call the Block periodically during long-running match operations.
	NSMatchingReportProgress NSMatchingOptions = 1
	// NSMatchingWithTransparentBounds: Specifies that matching may examine parts of the string beyond the bounds of the search range, for purposes such as word boundary detection, lookahead, etc.
	NSMatchingWithTransparentBounds NSMatchingOptions = 8
	// NSMatchingWithoutAnchoringBounds: Specifies that `^` and `$` will not automatically match the beginning and end of the search range, but will still match the beginning and end of the entire string.
	NSMatchingWithoutAnchoringBounds NSMatchingOptions = 16
)

func (e NSMatchingOptions) String() string {
	switch e {
	case NSMatchingAnchored:
		return "NSMatchingAnchored"
	case NSMatchingReportCompletion:
		return "NSMatchingReportCompletion"
	case NSMatchingReportProgress:
		return "NSMatchingReportProgress"
	case NSMatchingWithTransparentBounds:
		return "NSMatchingWithTransparentBounds"
	case NSMatchingWithoutAnchoringBounds:
		return "NSMatchingWithoutAnchoringBounds"
	default:
		return fmt.Sprintf("NSMatchingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct
type NSMeasurementFormatterUnitOptions int

const (
	// NSMeasurementFormatterUnitOptionsNaturalScale: # Discussion
	NSMeasurementFormatterUnitOptionsNaturalScale NSMeasurementFormatterUnitOptions = 2
	// NSMeasurementFormatterUnitOptionsProvidedUnit: # Discussion
	NSMeasurementFormatterUnitOptionsProvidedUnit NSMeasurementFormatterUnitOptions = 1
	// NSMeasurementFormatterUnitOptionsTemperatureWithoutUnit: # Discussion
	NSMeasurementFormatterUnitOptionsTemperatureWithoutUnit NSMeasurementFormatterUnitOptions = 4
)

func (e NSMeasurementFormatterUnitOptions) String() string {
	switch e {
	case NSMeasurementFormatterUnitOptionsNaturalScale:
		return "NSMeasurementFormatterUnitOptionsNaturalScale"
	case NSMeasurementFormatterUnitOptionsProvidedUnit:
		return "NSMeasurementFormatterUnitOptionsProvidedUnit"
	case NSMeasurementFormatterUnitOptionsTemperatureWithoutUnit:
		return "NSMeasurementFormatterUnitOptionsTemperatureWithoutUnit"
	default:
		return fmt.Sprintf("NSMeasurementFormatterUnitOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NetService/Options
type NSNetServiceOptions int

const (
	// NSNetServiceListenForConnections: # Discussion
	NSNetServiceListenForConnections NSNetServiceOptions = 2
	// NSNetServiceNoAutoRename: Specifies that the network service should not rename itself in the event of a name collision.
	NSNetServiceNoAutoRename NSNetServiceOptions = 1
)

func (e NSNetServiceOptions) String() string {
	switch e {
	case NSNetServiceListenForConnections:
		return "NSNetServiceListenForConnections"
	case NSNetServiceNoAutoRename:
		return "NSNetServiceNoAutoRename"
	default:
		return fmt.Sprintf("NSNetServiceOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NetService/ErrorCode-swift.enum
type NSNetServicesError int

const (
	// NSNetServicesActivityInProgress: The net service cannot process the request at this time.
	NSNetServicesActivityInProgress NSNetServicesError = -72003
	// NSNetServicesBadArgumentError: An invalid argument was used when creating the [NSNetService] object.
	NSNetServicesBadArgumentError NSNetServicesError = -72004
	// NSNetServicesCancelledError: The client canceled the action.
	NSNetServicesCancelledError NSNetServicesError = -72005
	// NSNetServicesCollisionError: The service could not be published because the name is already in use.
	NSNetServicesCollisionError NSNetServicesError = -72001
	// NSNetServicesInvalidError: The net service was improperly configured.
	NSNetServicesInvalidError NSNetServicesError = -72006
	// NSNetServicesNotFoundError: The service could not be found on the network.
	NSNetServicesNotFoundError NSNetServicesError = -72002
	// NSNetServicesTimeoutError: The net service has timed out.
	NSNetServicesTimeoutError NSNetServicesError = -72007
	// NSNetServicesUnknownError: An unknown error occurred.
	NSNetServicesUnknownError NSNetServicesError = -72000
)

func (e NSNetServicesError) String() string {
	switch e {
	case NSNetServicesActivityInProgress:
		return "NSNetServicesActivityInProgress"
	case NSNetServicesBadArgumentError:
		return "NSNetServicesBadArgumentError"
	case NSNetServicesCancelledError:
		return "NSNetServicesCancelledError"
	case NSNetServicesCollisionError:
		return "NSNetServicesCollisionError"
	case NSNetServicesInvalidError:
		return "NSNetServicesInvalidError"
	case NSNetServicesNotFoundError:
		return "NSNetServicesNotFoundError"
	case NSNetServicesTimeoutError:
		return "NSNetServicesTimeoutError"
	case NSNetServicesUnknownError:
		return "NSNetServicesUnknownError"
	default:
		return fmt.Sprintf("NSNetServicesError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/NotificationCoalescing
type NSNotificationCoalescing int

const (
	// NSNotificationCoalescingOnName: Coalesce notifications with the same name.
	NSNotificationCoalescingOnName NSNotificationCoalescing = 1
	// NSNotificationCoalescingOnSender: Coalesce notifications with the same object.
	NSNotificationCoalescingOnSender NSNotificationCoalescing = 2
	// NSNotificationNoCoalescing: Do not coalesce notifications in the queue.
	NSNotificationNoCoalescing NSNotificationCoalescing = 0
)

func (e NSNotificationCoalescing) String() string {
	switch e {
	case NSNotificationCoalescingOnName:
		return "NSNotificationCoalescingOnName"
	case NSNotificationCoalescingOnSender:
		return "NSNotificationCoalescingOnSender"
	case NSNotificationNoCoalescing:
		return "NSNotificationNoCoalescing"
	default:
		return fmt.Sprintf("NSNotificationCoalescing(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/DistributedNotificationCenter/SuspensionBehavior
type NSNotificationSuspensionBehavior int

const (
	// NSNotificationSuspensionBehaviorCoalesce: The server only queues the last notification of the specified name and object; earlier notifications are dropped.
	NSNotificationSuspensionBehaviorCoalesce NSNotificationSuspensionBehavior = 2
	// NSNotificationSuspensionBehaviorDeliverImmediately: # Discussion
	NSNotificationSuspensionBehaviorDeliverImmediately NSNotificationSuspensionBehavior = 4
	// NSNotificationSuspensionBehaviorDrop: The server doesn’t queue any notifications with this name and object until the notification center resumes notification delivery.
	NSNotificationSuspensionBehaviorDrop NSNotificationSuspensionBehavior = 1
	// NSNotificationSuspensionBehaviorHold: The server holds all matching notifications until the queue has been filled (queue size determined by the server), at which point the server may flush queued notifications.
	NSNotificationSuspensionBehaviorHold NSNotificationSuspensionBehavior = 3
)

func (e NSNotificationSuspensionBehavior) String() string {
	switch e {
	case NSNotificationSuspensionBehaviorCoalesce:
		return "NSNotificationSuspensionBehaviorCoalesce"
	case NSNotificationSuspensionBehaviorDeliverImmediately:
		return "NSNotificationSuspensionBehaviorDeliverImmediately"
	case NSNotificationSuspensionBehaviorDrop:
		return "NSNotificationSuspensionBehaviorDrop"
	case NSNotificationSuspensionBehaviorHold:
		return "NSNotificationSuspensionBehaviorHold"
	default:
		return fmt.Sprintf("NSNotificationSuspensionBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/Behavior
type NSNumberFormatterBehavior int

const (
	// NSNumberFormatterBehavior10_0: The number-formatter behavior as it existed prior to macOS 10.4.
	NSNumberFormatterBehavior10_0 NSNumberFormatterBehavior = 1000
	// NSNumberFormatterBehavior10_4: The number-formatter behavior since macOS 10.4.
	NSNumberFormatterBehavior10_4 NSNumberFormatterBehavior = 1040
	// NSNumberFormatterBehaviorDefault: The number-formatter behavior set as the default for new instances.
	NSNumberFormatterBehaviorDefault NSNumberFormatterBehavior = 0
)

func (e NSNumberFormatterBehavior) String() string {
	switch e {
	case NSNumberFormatterBehavior10_0:
		return "NSNumberFormatterBehavior10_0"
	case NSNumberFormatterBehavior10_4:
		return "NSNumberFormatterBehavior10_4"
	case NSNumberFormatterBehaviorDefault:
		return "NSNumberFormatterBehaviorDefault"
	default:
		return fmt.Sprintf("NSNumberFormatterBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/PadPosition
type NSNumberFormatterPadPosition int

const (
	// NSNumberFormatterPadAfterPrefix: Specifies that the padding should occur after the prefix.
	NSNumberFormatterPadAfterPrefix NSNumberFormatterPadPosition = 1
	// NSNumberFormatterPadAfterSuffix: Specifies that the padding should occur after the suffix.
	NSNumberFormatterPadAfterSuffix NSNumberFormatterPadPosition = 0
	// NSNumberFormatterPadBeforePrefix: Specifies that the padding should occur before the prefix.
	NSNumberFormatterPadBeforePrefix NSNumberFormatterPadPosition = 0
	// NSNumberFormatterPadBeforeSuffix: Specifies that the padding should occur before the suffix.
	NSNumberFormatterPadBeforeSuffix NSNumberFormatterPadPosition = 2
)

func (e NSNumberFormatterPadPosition) String() string {
	switch e {
	case NSNumberFormatterPadAfterPrefix:
		return "NSNumberFormatterPadAfterPrefix"
	case NSNumberFormatterPadAfterSuffix:
		return "NSNumberFormatterPadAfterSuffix"
	case NSNumberFormatterPadBeforeSuffix:
		return "NSNumberFormatterPadBeforeSuffix"
	default:
		return fmt.Sprintf("NSNumberFormatterPadPosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/RoundingMode-swift.enum
type NSNumberFormatterRoundingMode int

const (
	// NSNumberFormatterRoundCeiling: Round towards positive infinity.
	NSNumberFormatterRoundCeiling NSNumberFormatterRoundingMode = 0
	// NSNumberFormatterRoundDown: Round towards zero.
	NSNumberFormatterRoundDown NSNumberFormatterRoundingMode = 2
	// NSNumberFormatterRoundFloor: Round towards negative infinity.
	NSNumberFormatterRoundFloor NSNumberFormatterRoundingMode = 1
	// NSNumberFormatterRoundHalfDown: Round towards the nearest integer, or towards zero if equidistant.
	NSNumberFormatterRoundHalfDown NSNumberFormatterRoundingMode = 5
	// NSNumberFormatterRoundHalfEven: Round towards the nearest integer, or towards an even number if equidistant.
	NSNumberFormatterRoundHalfEven NSNumberFormatterRoundingMode = 4
	// NSNumberFormatterRoundHalfUp: Round towards the nearest integer, or away from zero if equidistant.
	NSNumberFormatterRoundHalfUp NSNumberFormatterRoundingMode = 0
	// NSNumberFormatterRoundUp: Round away from zero.
	NSNumberFormatterRoundUp NSNumberFormatterRoundingMode = 3
)

func (e NSNumberFormatterRoundingMode) String() string {
	switch e {
	case NSNumberFormatterRoundCeiling:
		return "NSNumberFormatterRoundCeiling"
	case NSNumberFormatterRoundDown:
		return "NSNumberFormatterRoundDown"
	case NSNumberFormatterRoundFloor:
		return "NSNumberFormatterRoundFloor"
	case NSNumberFormatterRoundHalfDown:
		return "NSNumberFormatterRoundHalfDown"
	case NSNumberFormatterRoundHalfEven:
		return "NSNumberFormatterRoundHalfEven"
	case NSNumberFormatterRoundUp:
		return "NSNumberFormatterRoundUp"
	default:
		return fmt.Sprintf("NSNumberFormatterRoundingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NumberFormatter/Style
type NSNumberFormatterStyle int

const (
	// NSNumberFormatterCurrencyAccountingStyle: An accounting currency style format that uses the currency symbol defined by the number formatter locale.
	NSNumberFormatterCurrencyAccountingStyle NSNumberFormatterStyle = 10
	// NSNumberFormatterCurrencyISOCodeStyle: A currency style format that uses the ISO 4217 currency code defined by the number formatter locale.
	NSNumberFormatterCurrencyISOCodeStyle NSNumberFormatterStyle = 8
	// NSNumberFormatterCurrencyPluralStyle: A currency style format that uses the pluralized denomination defined by the number formatter locale.
	NSNumberFormatterCurrencyPluralStyle NSNumberFormatterStyle = 9
	// NSNumberFormatterCurrencyStyle: A currency style format that uses the currency symbol defined by the number formatter locale.
	NSNumberFormatterCurrencyStyle NSNumberFormatterStyle = 2
	// NSNumberFormatterDecimalStyle: A decimal style format.
	NSNumberFormatterDecimalStyle NSNumberFormatterStyle = 1
	// NSNumberFormatterNoStyle: An integer representation.
	NSNumberFormatterNoStyle NSNumberFormatterStyle = 0
	// NSNumberFormatterOrdinalStyle: An ordinal style format.
	NSNumberFormatterOrdinalStyle NSNumberFormatterStyle = 6
	// NSNumberFormatterPercentStyle: A percent style format.
	NSNumberFormatterPercentStyle NSNumberFormatterStyle = 3
	// NSNumberFormatterScientificStyle: A scientific style format.
	NSNumberFormatterScientificStyle NSNumberFormatterStyle = 4
	// NSNumberFormatterSpellOutStyle: A style format in which numbers are spelled out in the language defined by the number formatter locale.
	NSNumberFormatterSpellOutStyle NSNumberFormatterStyle = 5
)

func (e NSNumberFormatterStyle) String() string {
	switch e {
	case NSNumberFormatterCurrencyAccountingStyle:
		return "NSNumberFormatterCurrencyAccountingStyle"
	case NSNumberFormatterCurrencyISOCodeStyle:
		return "NSNumberFormatterCurrencyISOCodeStyle"
	case NSNumberFormatterCurrencyPluralStyle:
		return "NSNumberFormatterCurrencyPluralStyle"
	case NSNumberFormatterCurrencyStyle:
		return "NSNumberFormatterCurrencyStyle"
	case NSNumberFormatterDecimalStyle:
		return "NSNumberFormatterDecimalStyle"
	case NSNumberFormatterNoStyle:
		return "NSNumberFormatterNoStyle"
	case NSNumberFormatterOrdinalStyle:
		return "NSNumberFormatterOrdinalStyle"
	case NSNumberFormatterPercentStyle:
		return "NSNumberFormatterPercentStyle"
	case NSNumberFormatterScientificStyle:
		return "NSNumberFormatterScientificStyle"
	case NSNumberFormatterSpellOutStyle:
		return "NSNumberFormatterSpellOutStyle"
	default:
		return fmt.Sprintf("NSNumberFormatterStyle(%d)", e)
	}
}

type NSOpenStepUnicodeReserved uint

const (
	// NSOpenStepUnicodeReservedBase: Specifies lower bound for a Unicode character range reserved for Apple’s corporate use (the range is `0xF400–0xF8FF`).
	NSOpenStepUnicodeReservedBase NSOpenStepUnicodeReserved = 62464
)

func (e NSOpenStepUnicodeReserved) String() string {
	switch e {
	case NSOpenStepUnicodeReservedBase:
		return "NSOpenStepUnicodeReservedBase"
	default:
		return fmt.Sprintf("NSOpenStepUnicodeReserved(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/Operation/QueuePriority-swift.enum
type NSOperationQueuePriority int

const (
	// NSOperationQueuePriorityHigh: Operations receive high priority for execution.
	NSOperationQueuePriorityHigh NSOperationQueuePriority = 4
	// NSOperationQueuePriorityLow: Operations receive low priority for execution.
	NSOperationQueuePriorityLow NSOperationQueuePriority = -4
	// NSOperationQueuePriorityNormal: Operations receive the normal priority for execution.
	NSOperationQueuePriorityNormal NSOperationQueuePriority = 0
	// NSOperationQueuePriorityVeryHigh: Operations receive very high priority for execution.
	NSOperationQueuePriorityVeryHigh NSOperationQueuePriority = 8
	// NSOperationQueuePriorityVeryLow: Operations receive very low priority for execution.
	NSOperationQueuePriorityVeryLow NSOperationQueuePriority = -8
)

func (e NSOperationQueuePriority) String() string {
	switch e {
	case NSOperationQueuePriorityHigh:
		return "NSOperationQueuePriorityHigh"
	case NSOperationQueuePriorityLow:
		return "NSOperationQueuePriorityLow"
	case NSOperationQueuePriorityNormal:
		return "NSOperationQueuePriorityNormal"
	case NSOperationQueuePriorityVeryHigh:
		return "NSOperationQueuePriorityVeryHigh"
	case NSOperationQueuePriorityVeryLow:
		return "NSOperationQueuePriorityVeryLow"
	default:
		return fmt.Sprintf("NSOperationQueuePriority(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSOrderedCollectionDifferenceCalculationOptions
type NSOrderedCollectionDifferenceCalculationOptions int

const (
	// NSOrderedCollectionDifferenceCalculationInferMoves: An option that identifies insertions or removals as moves.
	NSOrderedCollectionDifferenceCalculationInferMoves NSOrderedCollectionDifferenceCalculationOptions = 4
	// NSOrderedCollectionDifferenceCalculationOmitInsertedObjects: An option that indicates that the difference should omit references to the insertions.
	NSOrderedCollectionDifferenceCalculationOmitInsertedObjects NSOrderedCollectionDifferenceCalculationOptions = 1
	// NSOrderedCollectionDifferenceCalculationOmitRemovedObjects: An option that indicates that the difference should omit references to the removals.
	NSOrderedCollectionDifferenceCalculationOmitRemovedObjects NSOrderedCollectionDifferenceCalculationOptions = 2
)

func (e NSOrderedCollectionDifferenceCalculationOptions) String() string {
	switch e {
	case NSOrderedCollectionDifferenceCalculationInferMoves:
		return "NSOrderedCollectionDifferenceCalculationInferMoves"
	case NSOrderedCollectionDifferenceCalculationOmitInsertedObjects:
		return "NSOrderedCollectionDifferenceCalculationOmitInsertedObjects"
	case NSOrderedCollectionDifferenceCalculationOmitRemovedObjects:
		return "NSOrderedCollectionDifferenceCalculationOmitRemovedObjects"
	default:
		return fmt.Sprintf("NSOrderedCollectionDifferenceCalculationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Options
type NSPersonNameComponentsFormatterOptions int

const (
	// NSPersonNameComponentsFormatterPhonetic: The formatter should format the component object’s `phoneticRepresentation` components instead of its own components.
	NSPersonNameComponentsFormatterPhonetic NSPersonNameComponentsFormatterOptions = 2
)

func (e NSPersonNameComponentsFormatterOptions) String() string {
	switch e {
	case NSPersonNameComponentsFormatterPhonetic:
		return "NSPersonNameComponentsFormatterPhonetic"
	default:
		return fmt.Sprintf("NSPersonNameComponentsFormatterOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/PersonNameComponentsFormatter/Style-swift.enum
type NSPersonNameComponentsFormatterStyle int

const (
	// NSPersonNameComponentsFormatterStyleAbbreviated: # Discussion
	NSPersonNameComponentsFormatterStyleAbbreviated NSPersonNameComponentsFormatterStyle = 4
	// NSPersonNameComponentsFormatterStyleDefault: # Discussion
	NSPersonNameComponentsFormatterStyleDefault NSPersonNameComponentsFormatterStyle = 0
	// NSPersonNameComponentsFormatterStyleLong: # Discussion
	NSPersonNameComponentsFormatterStyleLong NSPersonNameComponentsFormatterStyle = 3
	// NSPersonNameComponentsFormatterStyleMedium: # Discussion
	NSPersonNameComponentsFormatterStyleMedium NSPersonNameComponentsFormatterStyle = 2
	// NSPersonNameComponentsFormatterStyleShort: # Discussion
	NSPersonNameComponentsFormatterStyleShort NSPersonNameComponentsFormatterStyle = 1
)

func (e NSPersonNameComponentsFormatterStyle) String() string {
	switch e {
	case NSPersonNameComponentsFormatterStyleAbbreviated:
		return "NSPersonNameComponentsFormatterStyleAbbreviated"
	case NSPersonNameComponentsFormatterStyleDefault:
		return "NSPersonNameComponentsFormatterStyleDefault"
	case NSPersonNameComponentsFormatterStyleLong:
		return "NSPersonNameComponentsFormatterStyleLong"
	case NSPersonNameComponentsFormatterStyleMedium:
		return "NSPersonNameComponentsFormatterStyleMedium"
	case NSPersonNameComponentsFormatterStyleShort:
		return "NSPersonNameComponentsFormatterStyleShort"
	default:
		return fmt.Sprintf("NSPersonNameComponentsFormatterStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
type NSPointerFunctionsOptions int

const (
	// NSPointerFunctionsCStringPersonality: Use a string hash and `strcmp`; C-string ‘`%s`’ style description.
	NSPointerFunctionsCStringPersonality NSPointerFunctionsOptions = 768
	// NSPointerFunctionsCopyIn: Use the memory acquire function to allocate and copy items on input (see acquireFunction).
	NSPointerFunctionsCopyIn NSPointerFunctionsOptions = 65536
	// NSPointerFunctionsIntegerPersonality: Use unshifted value as hash and equality.
	NSPointerFunctionsIntegerPersonality NSPointerFunctionsOptions = 1280
	// NSPointerFunctionsMachVirtualMemory: Use Mach memory.
	NSPointerFunctionsMachVirtualMemory NSPointerFunctionsOptions = 4
	// NSPointerFunctionsMallocMemory: Use `free()` on removal, `calloc()` on copy in.
	NSPointerFunctionsMallocMemory NSPointerFunctionsOptions = 3
	// NSPointerFunctionsObjectPersonality: Use `hash` and `isEqual` methods for hashing and equality comparisons, use the `description` method for a description.
	NSPointerFunctionsObjectPersonality NSPointerFunctionsOptions = 0
	// NSPointerFunctionsObjectPointerPersonality: Use shifted pointer for the hash value and direct comparison to determine equality; use the `description` method for a description.
	NSPointerFunctionsObjectPointerPersonality NSPointerFunctionsOptions = 512
	// NSPointerFunctionsOpaqueMemory: Take no action when pointers are deleted.
	NSPointerFunctionsOpaqueMemory NSPointerFunctionsOptions = 2
	// NSPointerFunctionsOpaquePersonality: Use shifted pointer for the hash value and direct comparison to determine equality.
	NSPointerFunctionsOpaquePersonality NSPointerFunctionsOptions = 256
	// NSPointerFunctionsStrongMemory: Use strong write-barriers to backing store; use garbage-collected memory on copy-in.
	NSPointerFunctionsStrongMemory NSPointerFunctionsOptions = 0
	// NSPointerFunctionsStructPersonality: Use a memory hash and `memcmp` (using a size function that you must set—see sizeFunction).
	NSPointerFunctionsStructPersonality NSPointerFunctionsOptions = 1024
	// NSPointerFunctionsWeakMemory: Uses weak read and write barriers appropriate for ARC or GC.
	NSPointerFunctionsWeakMemory NSPointerFunctionsOptions = 5
	// Deprecated.
	NSPointerFunctionsZeroingWeakMemory NSPointerFunctionsOptions = 1
)

func (e NSPointerFunctionsOptions) String() string {
	switch e {
	case NSPointerFunctionsCStringPersonality:
		return "NSPointerFunctionsCStringPersonality"
	case NSPointerFunctionsCopyIn:
		return "NSPointerFunctionsCopyIn"
	case NSPointerFunctionsIntegerPersonality:
		return "NSPointerFunctionsIntegerPersonality"
	case NSPointerFunctionsMachVirtualMemory:
		return "NSPointerFunctionsMachVirtualMemory"
	case NSPointerFunctionsMallocMemory:
		return "NSPointerFunctionsMallocMemory"
	case NSPointerFunctionsObjectPersonality:
		return "NSPointerFunctionsObjectPersonality"
	case NSPointerFunctionsObjectPointerPersonality:
		return "NSPointerFunctionsObjectPointerPersonality"
	case NSPointerFunctionsOpaqueMemory:
		return "NSPointerFunctionsOpaqueMemory"
	case NSPointerFunctionsOpaquePersonality:
		return "NSPointerFunctionsOpaquePersonality"
	case NSPointerFunctionsStructPersonality:
		return "NSPointerFunctionsStructPersonality"
	case NSPointerFunctionsWeakMemory:
		return "NSPointerFunctionsWeakMemory"
	case NSPointerFunctionsZeroingWeakMemory:
		return "NSPointerFunctionsZeroingWeakMemory"
	default:
		return fmt.Sprintf("NSPointerFunctionsOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NotificationQueue/PostingStyle
type NSPostingStyle int

const (
	// NSPostASAP: The notification is posted at the end of the current notification callout or timer.
	NSPostASAP NSPostingStyle = 2
	// NSPostNow: The notification is posted immediately after coalescing.
	NSPostNow NSPostingStyle = 3
	// NSPostWhenIdle: The notification is posted when the run loop is idle.
	NSPostWhenIdle NSPostingStyle = 1
)

func (e NSPostingStyle) String() string {
	switch e {
	case NSPostASAP:
		return "NSPostASAP"
	case NSPostNow:
		return "NSPostNow"
	case NSPostWhenIdle:
		return "NSPostWhenIdle"
	default:
		return fmt.Sprintf("NSPostingStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSComparisonPredicate/Operator
type NSPredicateOperatorType int

const (
	// NSBeginsWithPredicateOperatorType: A begins-with predicate.
	NSBeginsWithPredicateOperatorType NSPredicateOperatorType = 8
	// NSBetweenPredicateOperatorType: A predicate to determine if the left hand side lies at or between bounds specified by the right hand side.
	NSBetweenPredicateOperatorType NSPredicateOperatorType = 100
	// NSContainsPredicateOperatorType: A predicate to determine if the left hand side contains the right hand side.
	NSContainsPredicateOperatorType NSPredicateOperatorType = 99
	// NSCustomSelectorPredicateOperatorType: A predicate that uses a custom selector that takes a single argument and returns a [BOOL] value.
	NSCustomSelectorPredicateOperatorType NSPredicateOperatorType = 11
	// NSEndsWithPredicateOperatorType: An ends-with predicate.
	NSEndsWithPredicateOperatorType NSPredicateOperatorType = 9
	// NSEqualToPredicateOperatorType: An equal-to predicate.
	NSEqualToPredicateOperatorType NSPredicateOperatorType = 4
	// NSGreaterThanOrEqualToPredicateOperatorType: A greater-than-or-equal-to predicate.
	NSGreaterThanOrEqualToPredicateOperatorType NSPredicateOperatorType = 3
	// NSGreaterThanPredicateOperatorType: A greater-than predicate.
	NSGreaterThanPredicateOperatorType NSPredicateOperatorType = 2
	// NSInPredicateOperatorType: A predicate to determine if the left hand side is in the right hand side.
	NSInPredicateOperatorType NSPredicateOperatorType = 10
	// NSLessThanOrEqualToPredicateOperatorType: A less-than-or-equal-to predicate.
	NSLessThanOrEqualToPredicateOperatorType NSPredicateOperatorType = 1
	// NSLessThanPredicateOperatorType: A less-than predicate.
	NSLessThanPredicateOperatorType NSPredicateOperatorType = 0
	// NSLikePredicateOperatorType: A simple subset of the MATCHES predicate, similar in behavior to SQL [LIKE].
	NSLikePredicateOperatorType NSPredicateOperatorType = 7
	// NSMatchesPredicateOperatorType: A full regular expression matching predicate.
	NSMatchesPredicateOperatorType NSPredicateOperatorType = 6
	// NSNotEqualToPredicateOperatorType: A not-equal-to predicate.
	NSNotEqualToPredicateOperatorType NSPredicateOperatorType = 5
)

func (e NSPredicateOperatorType) String() string {
	switch e {
	case NSBeginsWithPredicateOperatorType:
		return "NSBeginsWithPredicateOperatorType"
	case NSBetweenPredicateOperatorType:
		return "NSBetweenPredicateOperatorType"
	case NSContainsPredicateOperatorType:
		return "NSContainsPredicateOperatorType"
	case NSCustomSelectorPredicateOperatorType:
		return "NSCustomSelectorPredicateOperatorType"
	case NSEndsWithPredicateOperatorType:
		return "NSEndsWithPredicateOperatorType"
	case NSEqualToPredicateOperatorType:
		return "NSEqualToPredicateOperatorType"
	case NSGreaterThanOrEqualToPredicateOperatorType:
		return "NSGreaterThanOrEqualToPredicateOperatorType"
	case NSGreaterThanPredicateOperatorType:
		return "NSGreaterThanPredicateOperatorType"
	case NSInPredicateOperatorType:
		return "NSInPredicateOperatorType"
	case NSLessThanOrEqualToPredicateOperatorType:
		return "NSLessThanOrEqualToPredicateOperatorType"
	case NSLessThanPredicateOperatorType:
		return "NSLessThanPredicateOperatorType"
	case NSLikePredicateOperatorType:
		return "NSLikePredicateOperatorType"
	case NSMatchesPredicateOperatorType:
		return "NSMatchesPredicateOperatorType"
	case NSNotEqualToPredicateOperatorType:
		return "NSNotEqualToPredicateOperatorType"
	default:
		return fmt.Sprintf("NSPredicateOperatorType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntentKind
type NSPresentationIntentKind int

const (
	// NSPresentationIntentKindBlockQuote: A presentation style for a block quote.
	NSPresentationIntentKindBlockQuote NSPresentationIntentKind = 6
	// NSPresentationIntentKindCodeBlock: A presentation style for a block of code.
	NSPresentationIntentKindCodeBlock NSPresentationIntentKind = 5
	// NSPresentationIntentKindHeader: A presentation style for a section header.
	NSPresentationIntentKindHeader NSPresentationIntentKind = 1
	// NSPresentationIntentKindListItem: A presentation style for a list of items.
	NSPresentationIntentKindListItem NSPresentationIntentKind = 4
	// NSPresentationIntentKindOrderedList: A presentation style for an ordered list of items.
	NSPresentationIntentKindOrderedList NSPresentationIntentKind = 2
	// NSPresentationIntentKindParagraph: A presentation style for a paragraph of text.
	NSPresentationIntentKindParagraph NSPresentationIntentKind = 0
	// NSPresentationIntentKindTable: A presentation style for a table.
	NSPresentationIntentKindTable NSPresentationIntentKind = 8
	// NSPresentationIntentKindTableHeaderRow: A presentation style for the header row of a table.
	NSPresentationIntentKindTableHeaderRow NSPresentationIntentKind = 9
	// NSPresentationIntentKindTableRow: A presentation style for a row of a table.
	NSPresentationIntentKindTableRow NSPresentationIntentKind = 10
	// NSPresentationIntentKindThematicBreak: A presentation style for a horizontal rule.
	NSPresentationIntentKindThematicBreak NSPresentationIntentKind = 7
	// NSPresentationIntentKindUnorderedList: A presentation style for an unordered list of items.
	NSPresentationIntentKindUnorderedList NSPresentationIntentKind = 3
)

func (e NSPresentationIntentKind) String() string {
	switch e {
	case NSPresentationIntentKindBlockQuote:
		return "NSPresentationIntentKindBlockQuote"
	case NSPresentationIntentKindCodeBlock:
		return "NSPresentationIntentKindCodeBlock"
	case NSPresentationIntentKindHeader:
		return "NSPresentationIntentKindHeader"
	case NSPresentationIntentKindListItem:
		return "NSPresentationIntentKindListItem"
	case NSPresentationIntentKindOrderedList:
		return "NSPresentationIntentKindOrderedList"
	case NSPresentationIntentKindParagraph:
		return "NSPresentationIntentKindParagraph"
	case NSPresentationIntentKindTable:
		return "NSPresentationIntentKindTable"
	case NSPresentationIntentKindTableHeaderRow:
		return "NSPresentationIntentKindTableHeaderRow"
	case NSPresentationIntentKindTableRow:
		return "NSPresentationIntentKindTableRow"
	case NSPresentationIntentKindThematicBreak:
		return "NSPresentationIntentKindThematicBreak"
	case NSPresentationIntentKindUnorderedList:
		return "NSPresentationIntentKindUnorderedList"
	default:
		return fmt.Sprintf("NSPresentationIntentKind(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSPresentationIntentTableColumnAlignment
type NSPresentationIntentTableColumnAlignment int

const (
	// NSPresentationIntentTableColumnAlignmentCenter: A presentation style for columns with center-aligned text.
	NSPresentationIntentTableColumnAlignmentCenter NSPresentationIntentTableColumnAlignment = 1
	// NSPresentationIntentTableColumnAlignmentLeft: A presentation style for columns with left-aligned text.
	NSPresentationIntentTableColumnAlignmentLeft NSPresentationIntentTableColumnAlignment = 0
	// NSPresentationIntentTableColumnAlignmentRight: A presentation style for columns with right-aligned text.
	NSPresentationIntentTableColumnAlignmentRight NSPresentationIntentTableColumnAlignment = 2
)

func (e NSPresentationIntentTableColumnAlignment) String() string {
	switch e {
	case NSPresentationIntentTableColumnAlignmentCenter:
		return "NSPresentationIntentTableColumnAlignmentCenter"
	case NSPresentationIntentTableColumnAlignmentLeft:
		return "NSPresentationIntentTableColumnAlignmentLeft"
	case NSPresentationIntentTableColumnAlignmentRight:
		return "NSPresentationIntentTableColumnAlignmentRight"
	default:
		return fmt.Sprintf("NSPresentationIntentTableColumnAlignment(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/ProcessInfo/ThermalState-swift.enum
type NSProcessInfoThermalState int

const (
	// NSProcessInfoThermalStateCritical: The thermal state is significantly impacting the performance of the system and the device needs to cool down.
	NSProcessInfoThermalStateCritical NSProcessInfoThermalState = 3
	// NSProcessInfoThermalStateFair: The thermal state is slightly elevated.
	NSProcessInfoThermalStateFair NSProcessInfoThermalState = 1
	// NSProcessInfoThermalStateNominal: The thermal state is within normal limits.
	NSProcessInfoThermalStateNominal NSProcessInfoThermalState = 0
	// NSProcessInfoThermalStateSerious: The thermal state is high.
	NSProcessInfoThermalStateSerious NSProcessInfoThermalState = 2
)

func (e NSProcessInfoThermalState) String() string {
	switch e {
	case NSProcessInfoThermalStateCritical:
		return "NSProcessInfoThermalStateCritical"
	case NSProcessInfoThermalStateFair:
		return "NSProcessInfoThermalStateFair"
	case NSProcessInfoThermalStateNominal:
		return "NSProcessInfoThermalStateNominal"
	case NSProcessInfoThermalStateSerious:
		return "NSProcessInfoThermalStateSerious"
	default:
		return fmt.Sprintf("NSProcessInfoThermalState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/PropertyListFormat
type NSPropertyListFormat int

const (
	// NSPropertyListBinaryFormat_v1_0: Specifies the binary property list format.
	NSPropertyListBinaryFormat_v1_0 NSPropertyListFormat = 0
	// NSPropertyListOpenStepFormat: Specifies the ASCII property list format inherited from the OpenStep APIs.
	NSPropertyListOpenStepFormat NSPropertyListFormat = 1
	// NSPropertyListXMLFormat_v1_0: Specifies the XML property list format.
	NSPropertyListXMLFormat_v1_0 NSPropertyListFormat = 100
)

func (e NSPropertyListFormat) String() string {
	switch e {
	case NSPropertyListBinaryFormat_v1_0:
		return "NSPropertyListBinaryFormat_v1_0"
	case NSPropertyListOpenStepFormat:
		return "NSPropertyListOpenStepFormat"
	case NSPropertyListXMLFormat_v1_0:
		return "NSPropertyListXMLFormat_v1_0"
	default:
		return fmt.Sprintf("NSPropertyListFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/PropertyListSerialization/MutabilityOptions
type NSPropertyListMutabilityOptions int

const (
	// NSPropertyListMutableContainers: Causes the returned property list to have mutable containers but immutable leaves.
	NSPropertyListMutableContainers NSPropertyListMutabilityOptions = 1
	// NSPropertyListMutableContainersAndLeaves: Causes the returned property list to have mutable containers and leaves.
	NSPropertyListMutableContainersAndLeaves NSPropertyListMutabilityOptions = 0
)

func (e NSPropertyListMutabilityOptions) String() string {
	switch e {
	case NSPropertyListMutableContainers:
		return "NSPropertyListMutableContainers"
	case NSPropertyListMutableContainersAndLeaves:
		return "NSPropertyListMutableContainersAndLeaves"
	default:
		return fmt.Sprintf("NSPropertyListMutabilityOptions(%d)", e)
	}
}

type NSProprietaryString uint

const (
	// NSProprietaryStringEncoding: Installation-specific encoding.
	NSProprietaryStringEncoding NSProprietaryString = 65536
)

func (e NSProprietaryString) String() string {
	switch e {
	case NSProprietaryStringEncoding:
		return "NSProprietaryStringEncoding"
	default:
		return fmt.Sprintf("NSProprietaryString(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/QualityOfService
type QualityOfService int

const (
	// NSQualityOfServiceBackground: # Discussion
	NSQualityOfServiceBackground QualityOfService = 9
	// NSQualityOfServiceDefault: # Discussion
	NSQualityOfServiceDefault QualityOfService = -1
	// NSQualityOfServiceUserInitiated: # Discussion
	NSQualityOfServiceUserInitiated QualityOfService = 25
	// NSQualityOfServiceUserInteractive: # Discussion
	NSQualityOfServiceUserInteractive QualityOfService = 33
	// NSQualityOfServiceUtility: # Discussion
	NSQualityOfServiceUtility QualityOfService = 17
)

func (e QualityOfService) String() string {
	switch e {
	case NSQualityOfServiceBackground:
		return "NSQualityOfServiceBackground"
	case NSQualityOfServiceDefault:
		return "NSQualityOfServiceDefault"
	case NSQualityOfServiceUserInitiated:
		return "NSQualityOfServiceUserInitiated"
	case NSQualityOfServiceUserInteractive:
		return "NSQualityOfServiceUserInteractive"
	case NSQualityOfServiceUtility:
		return "NSQualityOfServiceUtility"
	default:
		return fmt.Sprintf("QualityOfService(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSRectEdge
type NSRectEdge int

const (
	// NSMaxXEdge: The maximum X edge.
	NSMaxXEdge NSRectEdge = 2
	// NSMaxYEdge: The maximum Y edge.
	NSMaxYEdge NSRectEdge = 3
	// NSMinXEdge: The minimum X edge.
	NSMinXEdge NSRectEdge = 0
	// NSMinYEdge: The minimum Y edge.
	NSMinYEdge NSRectEdge = 1
	// NSRectEdgeMaxX: The maximum X edge.
	NSRectEdgeMaxX NSRectEdge = 2
	// NSRectEdgeMaxY: The maximum Y edge.
	NSRectEdgeMaxY NSRectEdge = 3
	// NSRectEdgeMinX: The minimum X edge.
	NSRectEdgeMinX NSRectEdge = 0
	// NSRectEdgeMinY: The minimum Y edge.
	NSRectEdgeMinY NSRectEdge = 1
)

func (e NSRectEdge) String() string {
	switch e {
	case NSMaxXEdge:
		return "NSMaxXEdge"
	case NSMaxYEdge:
		return "NSMaxYEdge"
	case NSMinXEdge:
		return "NSMinXEdge"
	case NSMinYEdge:
		return "NSMinYEdge"
	default:
		return fmt.Sprintf("NSRectEdge(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSRegularExpression/Options-swift.struct
type NSRegularExpressionOptions int

const (
	// NSRegularExpressionAllowCommentsAndWhitespace: Ignore whitespace and #-prefixed comments in the pattern.
	NSRegularExpressionAllowCommentsAndWhitespace NSRegularExpressionOptions = 2
	// NSRegularExpressionAnchorsMatchLines: Allow `^` and `$` to match the start and end of lines.
	NSRegularExpressionAnchorsMatchLines NSRegularExpressionOptions = 16
	// NSRegularExpressionCaseInsensitive: Match letters in the pattern independent of case.
	NSRegularExpressionCaseInsensitive NSRegularExpressionOptions = 1
	// NSRegularExpressionDotMatchesLineSeparators: Allow `.` to match any character, including line separators.
	NSRegularExpressionDotMatchesLineSeparators NSRegularExpressionOptions = 8
	// NSRegularExpressionIgnoreMetacharacters: Treat the entire pattern as a literal string.
	NSRegularExpressionIgnoreMetacharacters NSRegularExpressionOptions = 4
	// NSRegularExpressionUseUnicodeWordBoundaries: Use Unicode `TR#29` to specify word boundaries (otherwise, traditional regular expression word boundaries are used).
	NSRegularExpressionUseUnicodeWordBoundaries NSRegularExpressionOptions = 64
	// NSRegularExpressionUseUnixLineSeparators: Treat only `\n` as a line separator (otherwise, all standard line separators are used).
	NSRegularExpressionUseUnixLineSeparators NSRegularExpressionOptions = 32
)

func (e NSRegularExpressionOptions) String() string {
	switch e {
	case NSRegularExpressionAllowCommentsAndWhitespace:
		return "NSRegularExpressionAllowCommentsAndWhitespace"
	case NSRegularExpressionAnchorsMatchLines:
		return "NSRegularExpressionAnchorsMatchLines"
	case NSRegularExpressionCaseInsensitive:
		return "NSRegularExpressionCaseInsensitive"
	case NSRegularExpressionDotMatchesLineSeparators:
		return "NSRegularExpressionDotMatchesLineSeparators"
	case NSRegularExpressionIgnoreMetacharacters:
		return "NSRegularExpressionIgnoreMetacharacters"
	case NSRegularExpressionUseUnicodeWordBoundaries:
		return "NSRegularExpressionUseUnicodeWordBoundaries"
	case NSRegularExpressionUseUnixLineSeparators:
		return "NSRegularExpressionUseUnixLineSeparators"
	default:
		return fmt.Sprintf("NSRegularExpressionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/DateTimeStyle-swift.enum
type NSRelativeDateTimeFormatterStyle int

const (
	// NSRelativeDateTimeFormatterStyleNamed: A style that uses named styles to describe relative dates, such as “yesterday”, “last week”, or “next week”.
	NSRelativeDateTimeFormatterStyleNamed NSRelativeDateTimeFormatterStyle = 1
	// NSRelativeDateTimeFormatterStyleNumeric: A style that uses a numeric style to describe relative dates, such as “1 day ago” or “in 3 weeks”.
	NSRelativeDateTimeFormatterStyleNumeric NSRelativeDateTimeFormatterStyle = 0
)

func (e NSRelativeDateTimeFormatterStyle) String() string {
	switch e {
	case NSRelativeDateTimeFormatterStyleNamed:
		return "NSRelativeDateTimeFormatterStyleNamed"
	case NSRelativeDateTimeFormatterStyleNumeric:
		return "NSRelativeDateTimeFormatterStyleNumeric"
	default:
		return fmt.Sprintf("NSRelativeDateTimeFormatterStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/UnitsStyle-swift.enum
type NSRelativeDateTimeFormatterUnitsStyle int

const (
	// NSRelativeDateTimeFormatterUnitsStyleAbbreviated: A style that uses abbreviated units, such as “2 mo.
	NSRelativeDateTimeFormatterUnitsStyleAbbreviated NSRelativeDateTimeFormatterUnitsStyle = 3
	// NSRelativeDateTimeFormatterUnitsStyleFull: A style that uses full units, such as “2 months ago”.
	NSRelativeDateTimeFormatterUnitsStyleFull NSRelativeDateTimeFormatterUnitsStyle = 0
	// NSRelativeDateTimeFormatterUnitsStyleShort: A style that uses shortened units, such as “2 mo.
	NSRelativeDateTimeFormatterUnitsStyleShort NSRelativeDateTimeFormatterUnitsStyle = 2
	// NSRelativeDateTimeFormatterUnitsStyleSpellOut: A style that spells out units such as “two months ago”.
	NSRelativeDateTimeFormatterUnitsStyleSpellOut NSRelativeDateTimeFormatterUnitsStyle = 1
)

func (e NSRelativeDateTimeFormatterUnitsStyle) String() string {
	switch e {
	case NSRelativeDateTimeFormatterUnitsStyleAbbreviated:
		return "NSRelativeDateTimeFormatterUnitsStyleAbbreviated"
	case NSRelativeDateTimeFormatterUnitsStyleFull:
		return "NSRelativeDateTimeFormatterUnitsStyleFull"
	case NSRelativeDateTimeFormatterUnitsStyleShort:
		return "NSRelativeDateTimeFormatterUnitsStyleShort"
	case NSRelativeDateTimeFormatterUnitsStyleSpellOut:
		return "NSRelativeDateTimeFormatterUnitsStyleSpellOut"
	default:
		return fmt.Sprintf("NSRelativeDateTimeFormatterUnitsStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSRelativeSpecifier/RelativePosition-swift.enum
type NSRelativePosition int

const (
	// NSRelativeAfter: Specifies a position after another object.
	NSRelativeAfter NSRelativePosition = 0
	// NSRelativeBefore: Specifies a position before another object.
	NSRelativeBefore NSRelativePosition = 1
)

func (e NSRelativePosition) String() string {
	switch e {
	case NSRelativeAfter:
		return "NSRelativeAfter"
	case NSRelativeBefore:
		return "NSRelativeBefore"
	default:
		return fmt.Sprintf("NSRelativePosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSDecimalNumber/RoundingMode
type NSRoundingMode int

const (
	// NSRoundBankers: Round to the closest possible return value; when halfway between two possibilities, return the possibility whose last digit is even.
	NSRoundBankers NSRoundingMode = 3
	// NSRoundDown: Round return values down.
	NSRoundDown NSRoundingMode = 1
	// NSRoundPlain: Round to the closest possible return value; when caught halfway between two positive numbers, round up; when caught between two negative numbers, round down.
	NSRoundPlain NSRoundingMode = 0
	// NSRoundUp: Round return values up.
	NSRoundUp NSRoundingMode = 2
)

func (e NSRoundingMode) String() string {
	switch e {
	case NSRoundBankers:
		return "NSRoundBankers"
	case NSRoundDown:
		return "NSRoundDown"
	case NSRoundPlain:
		return "NSRoundPlain"
	case NSRoundUp:
		return "NSRoundUp"
	default:
		return fmt.Sprintf("NSRoundingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSSaveOptions
type NSSaveOptions int

const (
	// NSSaveOptionsAsk: Indicates the user should be asked before saving any modified documents on closing.
	NSSaveOptionsAsk NSSaveOptions = 2
	// NSSaveOptionsNo: Indicates a modified document should not be saved on closing.
	NSSaveOptionsNo NSSaveOptions = 1
	// NSSaveOptionsYes: Indicates a modified document should be saved on closing without asking the user.
	NSSaveOptionsYes NSSaveOptions = 0
)

func (e NSSaveOptions) String() string {
	switch e {
	case NSSaveOptionsAsk:
		return "NSSaveOptionsAsk"
	case NSSaveOptionsNo:
		return "NSSaveOptionsNo"
	case NSSaveOptionsYes:
		return "NSSaveOptionsYes"
	default:
		return fmt.Sprintf("NSSaveOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDirectory
type NSSearchPathDirectory int

const (
	// NSAdminApplicationDirectory: System and network administration applications.
	NSAdminApplicationDirectory NSSearchPathDirectory = 4
	// NSAllApplicationsDirectory: All directories where applications can be stored.
	NSAllApplicationsDirectory NSSearchPathDirectory = 100
	// NSAllLibrariesDirectory: All directories where resources can be stored.
	NSAllLibrariesDirectory NSSearchPathDirectory = 101
	// NSApplicationDirectory: Supported applications (`/Applications`).
	NSApplicationDirectory NSSearchPathDirectory = 1
	// NSApplicationScriptsDirectory: The user scripts folder for the calling application (`~/Library/Application Scripts/<code-signing-id>`.
	NSApplicationScriptsDirectory NSSearchPathDirectory = 23
	// NSApplicationSupportDirectory: Application support files (`Library/Application Support`).
	NSApplicationSupportDirectory NSSearchPathDirectory = 14
	// NSAutosavedInformationDirectory: The user’s autosaved documents (`Library/Autosave Information`).
	NSAutosavedInformationDirectory NSSearchPathDirectory = 11
	// NSCachesDirectory: Discardable cache files (`Library/Caches`).
	NSCachesDirectory NSSearchPathDirectory = 13
	// NSCoreServiceDirectory: Core services (`System/Library/CoreServices`).
	NSCoreServiceDirectory NSSearchPathDirectory = 10
	// NSDemoApplicationDirectory: Unsupported applications and demonstration versions.
	NSDemoApplicationDirectory NSSearchPathDirectory = 2
	// NSDesktopDirectory: The user’s desktop directory.
	NSDesktopDirectory NSSearchPathDirectory = 12
	// NSDeveloperApplicationDirectory: Developer applications (`/Developer/Applications`).
	NSDeveloperApplicationDirectory NSSearchPathDirectory = 3
	// NSDeveloperDirectory: Developer resources (`/Developer`).
	NSDeveloperDirectory NSSearchPathDirectory = 6
	// NSDocumentDirectory: Document directory.
	NSDocumentDirectory NSSearchPathDirectory = 9
	// NSDocumentationDirectory: Documentation.
	NSDocumentationDirectory NSSearchPathDirectory = 8
	// NSDownloadsDirectory: The user’s downloads directory.
	NSDownloadsDirectory NSSearchPathDirectory = 15
	// NSInputMethodsDirectory: Input Methods `(Library/Input Methods)`.
	NSInputMethodsDirectory NSSearchPathDirectory = 16
	// NSItemReplacementDirectory: The constant used to create a temporary directory.
	NSItemReplacementDirectory NSSearchPathDirectory = 99
	// NSLibraryDirectory: Various user-visible documentation, support, and configuration files (`/Library`).
	NSLibraryDirectory NSSearchPathDirectory = 5
	// NSMoviesDirectory: The user’s Movies directory `(~/Movies`).
	NSMoviesDirectory NSSearchPathDirectory = 17
	// NSMusicDirectory: The user’s Music directory (`~/Music`).
	NSMusicDirectory NSSearchPathDirectory = 18
	// NSPicturesDirectory: The user’s Pictures directory (`~/Pictures`).
	NSPicturesDirectory NSSearchPathDirectory = 19
	// NSPreferencePanesDirectory: The PreferencePanes directory for use with System Preferences (`Library/PreferencePanes`).
	NSPreferencePanesDirectory NSSearchPathDirectory = 22
	// NSPrinterDescriptionDirectory: The system’s PPDs directory (`Library/Printers/PPDs`).
	NSPrinterDescriptionDirectory NSSearchPathDirectory = 20
	// NSSharedPublicDirectory: The user’s Public sharing directory (`~/Public`).
	NSSharedPublicDirectory NSSearchPathDirectory = 21
	// NSTrashDirectory: The trash directory.
	NSTrashDirectory NSSearchPathDirectory = 102
	// NSUserDirectory: User home directories (`/Users`).
	NSUserDirectory NSSearchPathDirectory = 7
)

func (e NSSearchPathDirectory) String() string {
	switch e {
	case NSAdminApplicationDirectory:
		return "NSAdminApplicationDirectory"
	case NSAllApplicationsDirectory:
		return "NSAllApplicationsDirectory"
	case NSAllLibrariesDirectory:
		return "NSAllLibrariesDirectory"
	case NSApplicationDirectory:
		return "NSApplicationDirectory"
	case NSApplicationScriptsDirectory:
		return "NSApplicationScriptsDirectory"
	case NSApplicationSupportDirectory:
		return "NSApplicationSupportDirectory"
	case NSAutosavedInformationDirectory:
		return "NSAutosavedInformationDirectory"
	case NSCachesDirectory:
		return "NSCachesDirectory"
	case NSCoreServiceDirectory:
		return "NSCoreServiceDirectory"
	case NSDemoApplicationDirectory:
		return "NSDemoApplicationDirectory"
	case NSDesktopDirectory:
		return "NSDesktopDirectory"
	case NSDeveloperApplicationDirectory:
		return "NSDeveloperApplicationDirectory"
	case NSDeveloperDirectory:
		return "NSDeveloperDirectory"
	case NSDocumentDirectory:
		return "NSDocumentDirectory"
	case NSDocumentationDirectory:
		return "NSDocumentationDirectory"
	case NSDownloadsDirectory:
		return "NSDownloadsDirectory"
	case NSInputMethodsDirectory:
		return "NSInputMethodsDirectory"
	case NSItemReplacementDirectory:
		return "NSItemReplacementDirectory"
	case NSLibraryDirectory:
		return "NSLibraryDirectory"
	case NSMoviesDirectory:
		return "NSMoviesDirectory"
	case NSMusicDirectory:
		return "NSMusicDirectory"
	case NSPicturesDirectory:
		return "NSPicturesDirectory"
	case NSPreferencePanesDirectory:
		return "NSPreferencePanesDirectory"
	case NSPrinterDescriptionDirectory:
		return "NSPrinterDescriptionDirectory"
	case NSSharedPublicDirectory:
		return "NSSharedPublicDirectory"
	case NSTrashDirectory:
		return "NSTrashDirectory"
	case NSUserDirectory:
		return "NSUserDirectory"
	default:
		return fmt.Sprintf("NSSearchPathDirectory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/SearchPathDomainMask
type NSSearchPathDomainMask int

const (
	// NSAllDomainsMask: All domains.
	NSAllDomainsMask NSSearchPathDomainMask = 65535
	// NSLocalDomainMask: The place to install items available to everyone on this machine.
	NSLocalDomainMask NSSearchPathDomainMask = 2
	// NSNetworkDomainMask: The place to install items available on the network (`/Network`).
	NSNetworkDomainMask NSSearchPathDomainMask = 4
	// NSSystemDomainMask: A directory for system files provided by Apple (`/System`) .
	NSSystemDomainMask NSSearchPathDomainMask = 8
	// NSUserDomainMask: The user’s home directory—the place to install user’s personal items (`~`).
	NSUserDomainMask NSSearchPathDomainMask = 1
)

func (e NSSearchPathDomainMask) String() string {
	switch e {
	case NSAllDomainsMask:
		return "NSAllDomainsMask"
	case NSLocalDomainMask:
		return "NSLocalDomainMask"
	case NSNetworkDomainMask:
		return "NSNetworkDomainMask"
	case NSSystemDomainMask:
		return "NSSystemDomainMask"
	case NSUserDomainMask:
		return "NSUserDomainMask"
	default:
		return fmt.Sprintf("NSSearchPathDomainMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSSortOptions
type NSSortOptions int

const (
	// NSSortConcurrent: Specifies that the Block sort operation should be concurrent.
	NSSortConcurrent NSSortOptions = 1
	// NSSortStable: Specifies that the sorted results should return compared items having equal value in the order they occurred originally.
	NSSortStable NSSortOptions = 16
)

func (e NSSortOptions) String() string {
	switch e {
	case NSSortConcurrent:
		return "NSSortConcurrent"
	case NSSortStable:
		return "NSSortStable"
	default:
		return fmt.Sprintf("NSSortOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSAttributedString/SpellingState
type NSSpellingState int

const (
)

// See: https://developer.apple.com/documentation/Foundation/Stream/Event
type NSStreamEvent int

const (
	// NSStreamEventEndEncountered: The end of the stream has been reached.
	NSStreamEventEndEncountered NSStreamEvent = 16
	// NSStreamEventErrorOccurred: An error has occurred on the stream.
	NSStreamEventErrorOccurred NSStreamEvent = 8
	// NSStreamEventHasBytesAvailable: The stream has bytes to be read.
	NSStreamEventHasBytesAvailable NSStreamEvent = 2
	// NSStreamEventHasSpaceAvailable: The stream can accept bytes for writing.
	NSStreamEventHasSpaceAvailable NSStreamEvent = 4
	// NSStreamEventOpenCompleted: The open has completed successfully.
	NSStreamEventOpenCompleted NSStreamEvent = 1
)

func (e NSStreamEvent) String() string {
	switch e {
	case NSStreamEventEndEncountered:
		return "NSStreamEventEndEncountered"
	case NSStreamEventErrorOccurred:
		return "NSStreamEventErrorOccurred"
	case NSStreamEventHasBytesAvailable:
		return "NSStreamEventHasBytesAvailable"
	case NSStreamEventHasSpaceAvailable:
		return "NSStreamEventHasSpaceAvailable"
	case NSStreamEventOpenCompleted:
		return "NSStreamEventOpenCompleted"
	default:
		return fmt.Sprintf("NSStreamEvent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/Stream/Status
type NSStreamStatus int

const (
	// NSStreamStatusAtEnd: There is no more data to read, or no more data can be written to the stream.
	NSStreamStatusAtEnd NSStreamStatus = 5
	// NSStreamStatusClosed: The stream is closed (close() has been called on it).
	NSStreamStatusClosed NSStreamStatus = 6
	// NSStreamStatusError: The remote end of the connection can’t be contacted, or the connection has been severed for some other reason.
	NSStreamStatusError NSStreamStatus = 7
	// NSStreamStatusNotOpen: The stream is not open for reading or writing.
	NSStreamStatusNotOpen NSStreamStatus = 0
	// NSStreamStatusOpen: The stream is open, but no reading or writing is occurring.
	NSStreamStatusOpen NSStreamStatus = 2
	// NSStreamStatusOpening: The stream is in the process of being opened for reading or for writing.
	NSStreamStatusOpening NSStreamStatus = 1
	// NSStreamStatusReading: Data is being read from the stream.
	NSStreamStatusReading NSStreamStatus = 3
	// NSStreamStatusWriting: Data is being written to the stream.
	NSStreamStatusWriting NSStreamStatus = 4
)

func (e NSStreamStatus) String() string {
	switch e {
	case NSStreamStatusAtEnd:
		return "NSStreamStatusAtEnd"
	case NSStreamStatusClosed:
		return "NSStreamStatusClosed"
	case NSStreamStatusError:
		return "NSStreamStatusError"
	case NSStreamStatusNotOpen:
		return "NSStreamStatusNotOpen"
	case NSStreamStatusOpen:
		return "NSStreamStatusOpen"
	case NSStreamStatusOpening:
		return "NSStreamStatusOpening"
	case NSStreamStatusReading:
		return "NSStreamStatusReading"
	case NSStreamStatusWriting:
		return "NSStreamStatusWriting"
	default:
		return fmt.Sprintf("NSStreamStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSString/CompareOptions
type NSStringCompareOptions int

const (
	// NSAnchoredSearch: Search is limited to start (or end, if [NSBackwardsSearch]) of source string.
	NSAnchoredSearch NSStringCompareOptions = 8
	// NSBackwardsSearch: Search from end of source string.
	NSBackwardsSearch NSStringCompareOptions = 4
	// NSCaseInsensitiveSearch: A case-insensitive search.
	NSCaseInsensitiveSearch NSStringCompareOptions = 1
	// NSDiacriticInsensitiveSearch: Search ignores diacritic marks.
	NSDiacriticInsensitiveSearch NSStringCompareOptions = 128
	// NSForcedOrderingSearch: Comparisons are forced to return either [NSOrderedAscending] or [NSOrderedDescending] if the strings are equivalent but not strictly equal.
	NSForcedOrderingSearch NSStringCompareOptions = 512
	// NSLiteralSearch: Exact character-by-character equivalence.
	NSLiteralSearch NSStringCompareOptions = 2
	// NSNumericSearch: Numbers within strings are compared using numeric value, that is, `Name2.Txt()` < `Name7.Txt()` < `Name25.Txt()`.
	NSNumericSearch NSStringCompareOptions = 64
	// NSRegularExpressionSearch: The search string is treated as an ICU-compatible regular expression.
	NSRegularExpressionSearch NSStringCompareOptions = 1024
	// NSWidthInsensitiveSearch: Search ignores width differences in characters that have full-width and half-width forms, as occurs in East Asian character sets.
	NSWidthInsensitiveSearch NSStringCompareOptions = 256
)

func (e NSStringCompareOptions) String() string {
	switch e {
	case NSAnchoredSearch:
		return "NSAnchoredSearch"
	case NSBackwardsSearch:
		return "NSBackwardsSearch"
	case NSCaseInsensitiveSearch:
		return "NSCaseInsensitiveSearch"
	case NSDiacriticInsensitiveSearch:
		return "NSDiacriticInsensitiveSearch"
	case NSForcedOrderingSearch:
		return "NSForcedOrderingSearch"
	case NSLiteralSearch:
		return "NSLiteralSearch"
	case NSNumericSearch:
		return "NSNumericSearch"
	case NSRegularExpressionSearch:
		return "NSRegularExpressionSearch"
	case NSWidthInsensitiveSearch:
		return "NSWidthInsensitiveSearch"
	default:
		return fmt.Sprintf("NSStringCompareOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSString/DrawingOptions
type NSStringDrawingOptions int

const (
	OptionsResolvesNaturalAlignmentWithBaseWritingDirection NSStringDrawingOptions = 0
	TruncatesLastVisibleLine NSStringDrawingOptions = 0
	UsesDeviceMetrics NSStringDrawingOptions = 0
	UsesFontLeading NSStringDrawingOptions = 0
	UsesLineFragmentOrigin NSStringDrawingOptions = 0
	// Deprecated.
	DisableScreenFontSubstitution NSStringDrawingOptions = 0
	// Deprecated.
	OneShot NSStringDrawingOptions = 0
)

func (e NSStringDrawingOptions) String() string {
	switch e {
	case OptionsResolvesNaturalAlignmentWithBaseWritingDirection:
		return "OptionsResolvesNaturalAlignmentWithBaseWritingDirection"
	default:
		return fmt.Sprintf("NSStringDrawingOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSString/EncodingConversionOptions
type NSStringEncodingConversionOptions int

const (
	// NSStringEncodingConversionAllowLossy: Allows lossy conversion.
	NSStringEncodingConversionAllowLossy NSStringEncodingConversionOptions = 1
	// NSStringEncodingConversionExternalRepresentation: Specifies an external representation (with a byte-order mark, if necessary, to indicate endianness).
	NSStringEncodingConversionExternalRepresentation NSStringEncodingConversionOptions = 2
)

func (e NSStringEncodingConversionOptions) String() string {
	switch e {
	case NSStringEncodingConversionAllowLossy:
		return "NSStringEncodingConversionAllowLossy"
	case NSStringEncodingConversionExternalRepresentation:
		return "NSStringEncodingConversionExternalRepresentation"
	default:
		return fmt.Sprintf("NSStringEncodingConversionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSString/EnumerationOptions
type NSStringEnumerationOptions int

const (
	NSStringEnumerationByCaretPositions NSStringEnumerationOptions = 5
	// NSStringEnumerationByComposedCharacterSequences: # Discussion
	NSStringEnumerationByComposedCharacterSequences NSStringEnumerationOptions = 2
	NSStringEnumerationByDeletionClusters NSStringEnumerationOptions = 6
	// NSStringEnumerationByLines: # Discussion
	NSStringEnumerationByLines NSStringEnumerationOptions = 0
	// NSStringEnumerationByParagraphs: # Discussion
	NSStringEnumerationByParagraphs NSStringEnumerationOptions = 1
	// NSStringEnumerationBySentences: # Discussion
	NSStringEnumerationBySentences NSStringEnumerationOptions = 4
	// NSStringEnumerationByWords: # Discussion
	NSStringEnumerationByWords NSStringEnumerationOptions = 3
	// NSStringEnumerationLocalized: # Discussion
	NSStringEnumerationLocalized NSStringEnumerationOptions = 1024
	// NSStringEnumerationReverse: # Discussion
	NSStringEnumerationReverse NSStringEnumerationOptions = 256
	// NSStringEnumerationSubstringNotRequired: # Discussion
	NSStringEnumerationSubstringNotRequired NSStringEnumerationOptions = 512
)

func (e NSStringEnumerationOptions) String() string {
	switch e {
	case NSStringEnumerationByCaretPositions:
		return "NSStringEnumerationByCaretPositions"
	case NSStringEnumerationByComposedCharacterSequences:
		return "NSStringEnumerationByComposedCharacterSequences"
	case NSStringEnumerationByDeletionClusters:
		return "NSStringEnumerationByDeletionClusters"
	case NSStringEnumerationByLines:
		return "NSStringEnumerationByLines"
	case NSStringEnumerationByParagraphs:
		return "NSStringEnumerationByParagraphs"
	case NSStringEnumerationBySentences:
		return "NSStringEnumerationBySentences"
	case NSStringEnumerationByWords:
		return "NSStringEnumerationByWords"
	case NSStringEnumerationLocalized:
		return "NSStringEnumerationLocalized"
	case NSStringEnumerationReverse:
		return "NSStringEnumerationReverse"
	case NSStringEnumerationSubstringNotRequired:
		return "NSStringEnumerationSubstringNotRequired"
	default:
		return fmt.Sprintf("NSStringEnumerationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/Process/TerminationReason-swift.enum
type NSTaskTerminationReason int

const (
	// NSTaskTerminationReasonExit: The task exited normally.
	NSTaskTerminationReasonExit NSTaskTerminationReason = 1
	// NSTaskTerminationReasonUncaughtSignal: The task exited due to an uncaught signal.
	NSTaskTerminationReasonUncaughtSignal NSTaskTerminationReason = 2
)

func (e NSTaskTerminationReason) String() string {
	switch e {
	case NSTaskTerminationReasonExit:
		return "NSTaskTerminationReasonExit"
	case NSTaskTerminationReasonUncaughtSignal:
		return "NSTaskTerminationReasonUncaughtSignal"
	default:
		return fmt.Sprintf("NSTaskTerminationReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSSpecifierTest/TestComparisonOperation
type NSTestComparisonOperation int

const (
	// NSBeginsWithComparison: Binary containment operator that results in true if the test object is a list or string that matches the beginning of the other object (which is also a list or string).
	NSBeginsWithComparison NSTestComparisonOperation = 5
	// NSContainsComparison: Binary containment operator that results in true if the test object is a list or string that matches the other object (which is also a list or string) at any location.
	NSContainsComparison NSTestComparisonOperation = 7
	// NSEndsWithComparison: Binary containment operator that results in true if the test object is a list or string that matches the end of the other object (which is also a list or string).
	NSEndsWithComparison NSTestComparisonOperation = 6
	// NSEqualToComparison: Binary comparison operator that results in true if the two objects are equal.
	NSEqualToComparison NSTestComparisonOperation = 0
	// NSGreaterThanComparison: Binary comparison operator that results in true if the value of the test object is greater than the value of the other object.
	NSGreaterThanComparison NSTestComparisonOperation = 4
	// NSGreaterThanOrEqualToComparison: Binary comparison operator that results in true if the value of the test object is greater than or equal to the value of the other object.
	NSGreaterThanOrEqualToComparison NSTestComparisonOperation = 3
	// NSLessThanComparison: Binary comparison operator that results in true if the value of the test object is less than the value of the other object.
	NSLessThanComparison NSTestComparisonOperation = 2
	// NSLessThanOrEqualToComparison: Binary comparison operator that results in true if the value of the test object is equal to or less than the value of the other object.
	NSLessThanOrEqualToComparison NSTestComparisonOperation = 1
)

func (e NSTestComparisonOperation) String() string {
	switch e {
	case NSBeginsWithComparison:
		return "NSBeginsWithComparison"
	case NSContainsComparison:
		return "NSContainsComparison"
	case NSEndsWithComparison:
		return "NSEndsWithComparison"
	case NSEqualToComparison:
		return "NSEqualToComparison"
	case NSGreaterThanComparison:
		return "NSGreaterThanComparison"
	case NSGreaterThanOrEqualToComparison:
		return "NSGreaterThanOrEqualToComparison"
	case NSLessThanComparison:
		return "NSLessThanComparison"
	case NSLessThanOrEqualToComparison:
		return "NSLessThanOrEqualToComparison"
	default:
		return fmt.Sprintf("NSTestComparisonOperation(%d)", e)
	}
}

type NSTextCheckingAll uint

const (
	// NSTextCheckingAllCustomTypes: Checking types that can be used by clients.
	NSTextCheckingAllCustomTypes NSTextCheckingAll = 4294967295
	// NSTextCheckingAllSystemTypes: Checking types supported by the system.
	NSTextCheckingAllSystemTypes NSTextCheckingAll = 4294967295
	// NSTextCheckingAllTypes: All possible checking types, both system- and user-supported.
	NSTextCheckingAllTypes NSTextCheckingAll = 0
)

func (e NSTextCheckingAll) String() string {
	switch e {
	case NSTextCheckingAllCustomTypes:
		return "NSTextCheckingAllCustomTypes"
	case NSTextCheckingAllTypes:
		return "NSTextCheckingAllTypes"
	default:
		return fmt.Sprintf("NSTextCheckingAll(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult/CheckingType
type NSTextCheckingType int

const (
	// NSTextCheckingTypeAddress: Attempts to locate addresses.
	NSTextCheckingTypeAddress NSTextCheckingType = 16
	// NSTextCheckingTypeCorrection: Performs autocorrection on misspelled words.
	NSTextCheckingTypeCorrection NSTextCheckingType = 512
	// NSTextCheckingTypeDash: Replaces dashes with em-dashes.
	NSTextCheckingTypeDash NSTextCheckingType = 128
	// NSTextCheckingTypeDate: Attempts to locate dates.
	NSTextCheckingTypeDate NSTextCheckingType = 8
	// NSTextCheckingTypeGrammar: Checks grammar.
	NSTextCheckingTypeGrammar NSTextCheckingType = 4
	// NSTextCheckingTypeLink: Attempts to locate URL links.
	NSTextCheckingTypeLink NSTextCheckingType = 32
	// NSTextCheckingTypeOrthography: Attempts to identify the language
	NSTextCheckingTypeOrthography NSTextCheckingType = 1
	// NSTextCheckingTypePhoneNumber: Matches a phone number.
	NSTextCheckingTypePhoneNumber NSTextCheckingType = 2048
	// NSTextCheckingTypeQuote: Replaces quotes with smart quotes.
	NSTextCheckingTypeQuote NSTextCheckingType = 64
	// NSTextCheckingTypeRegularExpression: Matches a regular expression.
	NSTextCheckingTypeRegularExpression NSTextCheckingType = 1024
	// NSTextCheckingTypeReplacement: Replaces characters such as (c) with the appropriate symbol (in this case ©).
	NSTextCheckingTypeReplacement NSTextCheckingType = 256
	// NSTextCheckingTypeSpelling: Checks spelling.
	NSTextCheckingTypeSpelling NSTextCheckingType = 2
	// NSTextCheckingTypeTransitInformation: Matches a transit information, for example, flight information.
	NSTextCheckingTypeTransitInformation NSTextCheckingType = 4096
)

func (e NSTextCheckingType) String() string {
	switch e {
	case NSTextCheckingTypeAddress:
		return "NSTextCheckingTypeAddress"
	case NSTextCheckingTypeCorrection:
		return "NSTextCheckingTypeCorrection"
	case NSTextCheckingTypeDash:
		return "NSTextCheckingTypeDash"
	case NSTextCheckingTypeDate:
		return "NSTextCheckingTypeDate"
	case NSTextCheckingTypeGrammar:
		return "NSTextCheckingTypeGrammar"
	case NSTextCheckingTypeLink:
		return "NSTextCheckingTypeLink"
	case NSTextCheckingTypeOrthography:
		return "NSTextCheckingTypeOrthography"
	case NSTextCheckingTypePhoneNumber:
		return "NSTextCheckingTypePhoneNumber"
	case NSTextCheckingTypeQuote:
		return "NSTextCheckingTypeQuote"
	case NSTextCheckingTypeRegularExpression:
		return "NSTextCheckingTypeRegularExpression"
	case NSTextCheckingTypeReplacement:
		return "NSTextCheckingTypeReplacement"
	case NSTextCheckingTypeSpelling:
		return "NSTextCheckingTypeSpelling"
	case NSTextCheckingTypeTransitInformation:
		return "NSTextCheckingTypeTransitInformation"
	default:
		return fmt.Sprintf("NSTextCheckingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/NameStyle
type NSTimeZoneNameStyle int

const (
	// NSTimeZoneNameStyleDaylightSaving: Specifies a daylight saving name style.
	NSTimeZoneNameStyleDaylightSaving NSTimeZoneNameStyle = 2
	// NSTimeZoneNameStyleGeneric: Specifies a generic name style.
	NSTimeZoneNameStyleGeneric NSTimeZoneNameStyle = 4
	// NSTimeZoneNameStyleShortDaylightSaving: Specifies a short daylight saving name style.
	NSTimeZoneNameStyleShortDaylightSaving NSTimeZoneNameStyle = 3
	// NSTimeZoneNameStyleShortGeneric: Specifies a generic time zone name.
	NSTimeZoneNameStyleShortGeneric NSTimeZoneNameStyle = 5
	// NSTimeZoneNameStyleShortStandard: Specifies a short name style.
	NSTimeZoneNameStyleShortStandard NSTimeZoneNameStyle = 1
	// NSTimeZoneNameStyleStandard: Specifies a standard name style.
	NSTimeZoneNameStyleStandard NSTimeZoneNameStyle = 0
)

func (e NSTimeZoneNameStyle) String() string {
	switch e {
	case NSTimeZoneNameStyleDaylightSaving:
		return "NSTimeZoneNameStyleDaylightSaving"
	case NSTimeZoneNameStyleGeneric:
		return "NSTimeZoneNameStyleGeneric"
	case NSTimeZoneNameStyleShortDaylightSaving:
		return "NSTimeZoneNameStyleShortDaylightSaving"
	case NSTimeZoneNameStyleShortGeneric:
		return "NSTimeZoneNameStyleShortGeneric"
	case NSTimeZoneNameStyleShortStandard:
		return "NSTimeZoneNameStyleShortStandard"
	case NSTimeZoneNameStyleStandard:
		return "NSTimeZoneNameStyleStandard"
	default:
		return fmt.Sprintf("NSTimeZoneNameStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkCreationOptions
type NSURLBookmarkCreationOptions int

const (
	// NSURLBookmarkCreationMinimalBookmark: Specifies that when creating a bookmark, it includes minimal information.
	NSURLBookmarkCreationMinimalBookmark NSURLBookmarkCreationOptions = 512
	// NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess: Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read-only access to a file-system resource.
	NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess NSURLBookmarkCreationOptions = 4096
	// NSURLBookmarkCreationSuitableForBookmarkFile: Specifies that the bookmark data includes the required properties for creating Finder alias files.
	NSURLBookmarkCreationSuitableForBookmarkFile NSURLBookmarkCreationOptions = 1024
	// NSURLBookmarkCreationWithSecurityScope: Specifies that when creating a security-scoped bookmark, upon resolution, it provides a security-scoped URL allowing read/write access to a file-system resource.
	NSURLBookmarkCreationWithSecurityScope NSURLBookmarkCreationOptions = 2048
	// NSURLBookmarkCreationWithoutImplicitSecurityScope: Prevents inclusion of a bookmark’s implicit ephemeral security scope, when creating one without security scope.
	NSURLBookmarkCreationWithoutImplicitSecurityScope NSURLBookmarkCreationOptions = 536870912
	// Deprecated.
	NSURLBookmarkCreationPreferFileIDResolution NSURLBookmarkCreationOptions = 256
)

func (e NSURLBookmarkCreationOptions) String() string {
	switch e {
	case NSURLBookmarkCreationMinimalBookmark:
		return "NSURLBookmarkCreationMinimalBookmark"
	case NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess:
		return "NSURLBookmarkCreationSecurityScopeAllowOnlyReadAccess"
	case NSURLBookmarkCreationSuitableForBookmarkFile:
		return "NSURLBookmarkCreationSuitableForBookmarkFile"
	case NSURLBookmarkCreationWithSecurityScope:
		return "NSURLBookmarkCreationWithSecurityScope"
	case NSURLBookmarkCreationWithoutImplicitSecurityScope:
		return "NSURLBookmarkCreationWithoutImplicitSecurityScope"
	case NSURLBookmarkCreationPreferFileIDResolution:
		return "NSURLBookmarkCreationPreferFileIDResolution"
	default:
		return fmt.Sprintf("NSURLBookmarkCreationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURL/BookmarkResolutionOptions
type NSURLBookmarkResolutionOptions int

const (
	// NSURLBookmarkResolutionWithSecurityScope: Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	NSURLBookmarkResolutionWithSecurityScope NSURLBookmarkResolutionOptions = 1024
	// NSURLBookmarkResolutionWithoutImplicitStartAccessing: A property that specifies that resolution doesn’t implicitly start accessing the ephemeral security-scoped resource.
	NSURLBookmarkResolutionWithoutImplicitStartAccessing NSURLBookmarkResolutionOptions = 32768
	// NSURLBookmarkResolutionWithoutMounting: Specifies that no volume should be mounted during resolution of the bookmark data.
	NSURLBookmarkResolutionWithoutMounting NSURLBookmarkResolutionOptions = 512
	// NSURLBookmarkResolutionWithoutUI: Specifies that no UI feedback should accompany resolution of the bookmark data.
	NSURLBookmarkResolutionWithoutUI NSURLBookmarkResolutionOptions = 256
)

func (e NSURLBookmarkResolutionOptions) String() string {
	switch e {
	case NSURLBookmarkResolutionWithSecurityScope:
		return "NSURLBookmarkResolutionWithSecurityScope"
	case NSURLBookmarkResolutionWithoutImplicitStartAccessing:
		return "NSURLBookmarkResolutionWithoutImplicitStartAccessing"
	case NSURLBookmarkResolutionWithoutMounting:
		return "NSURLBookmarkResolutionWithoutMounting"
	case NSURLBookmarkResolutionWithoutUI:
		return "NSURLBookmarkResolutionWithoutUI"
	default:
		return fmt.Sprintf("NSURLBookmarkResolutionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLCache/StoragePolicy
type NSURLCacheStoragePolicy int

const (
	// NSURLCacheStorageAllowed: Storage in URLCache is allowed without restriction.
	NSURLCacheStorageAllowed NSURLCacheStoragePolicy = 0
	// NSURLCacheStorageAllowedInMemoryOnly: Storage in URLCache is allowed; however storage should be restricted to memory only.
	NSURLCacheStorageAllowedInMemoryOnly NSURLCacheStoragePolicy = 1
	// NSURLCacheStorageNotAllowed: Storage in URLCache is not allowed in any fashion, either in memory or on disk.
	NSURLCacheStorageNotAllowed NSURLCacheStoragePolicy = 2
)

func (e NSURLCacheStoragePolicy) String() string {
	switch e {
	case NSURLCacheStorageAllowed:
		return "NSURLCacheStorageAllowed"
	case NSURLCacheStorageAllowedInMemoryOnly:
		return "NSURLCacheStorageAllowedInMemoryOnly"
	case NSURLCacheStorageNotAllowed:
		return "NSURLCacheStorageNotAllowed"
	default:
		return fmt.Sprintf("NSURLCacheStoragePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLCredential/Persistence-swift.enum
type NSURLCredentialPersistence int

const (
	// NSURLCredentialPersistenceSynchronizable: The credential should be stored permanently in the keychain, and in addition should be distributed to other devices based on the owning Apple ID.
	NSURLCredentialPersistenceSynchronizable NSURLCredentialPersistence = 3
)

func (e NSURLCredentialPersistence) String() string {
	switch e {
	case NSURLCredentialPersistenceSynchronizable:
		return "NSURLCredentialPersistenceSynchronizable"
	default:
		return fmt.Sprintf("NSURLCredentialPersistence(%d)", e)
	}
}

type NSURLError int

const (
	// NSURLErrorAppTransportSecurityRequiresSecureConnection: App Transport Security disallowed a connection because there is no secure network connection.
	NSURLErrorAppTransportSecurityRequiresSecureConnection NSURLError = -1016
	// NSURLErrorBackgroundSessionInUseByAnotherProcess: An app or app extension attempted to connect to a background session that is already connected to a process.
	NSURLErrorBackgroundSessionInUseByAnotherProcess NSURLError = -3001
	// NSURLErrorBackgroundSessionRequiresSharedContainer: The shared container identifier of the URL session configuration is needed but hasn’t been set.
	NSURLErrorBackgroundSessionRequiresSharedContainer NSURLError = -3002
	// NSURLErrorBackgroundSessionWasDisconnected: The app is suspended or exits while a background data task is processing.
	NSURLErrorBackgroundSessionWasDisconnected NSURLError = -3000
	// NSURLErrorBadServerResponse: The URL Loading System received bad data from the server.
	NSURLErrorBadServerResponse NSURLError = -1011
	// NSURLErrorBadURL: A malformed URL prevented a URL request from being initiated.
	NSURLErrorBadURL NSURLError = -1000
	// NSURLErrorCallIsActive: A connection was attempted while a phone call was active on a network that doesn’t support simultaneous phone and data communication, such as EDGE or GPRS.
	NSURLErrorCallIsActive NSURLError = -3005
	// NSURLErrorCancelled: An asynchronous load has been canceled.
	NSURLErrorCancelled NSURLError = -999
	// NSURLErrorCannotCloseFile: A download task couldn’t close the downloaded file on disk.
	NSURLErrorCannotCloseFile NSURLError = -3002
	// NSURLErrorCannotConnectToHost: An attempt to connect to a host failed.
	NSURLErrorCannotConnectToHost NSURLError = -1004
	// NSURLErrorCannotCreateFile: A download task couldn’t create the downloaded file on disk because of an I/O failure.
	NSURLErrorCannotCreateFile NSURLError = -3000
	// NSURLErrorCannotDecodeContentData: Content data received during a connection request had an unknown content encoding.
	NSURLErrorCannotDecodeContentData NSURLError = -1016
	// NSURLErrorCannotDecodeRawData: Content data received during a connection request couldn’t be decoded for a known content encoding.
	NSURLErrorCannotDecodeRawData NSURLError = -1015
	// NSURLErrorCannotFindHost: The host name for a URL couldn’t be resolved.
	NSURLErrorCannotFindHost NSURLError = -1003
	// NSURLErrorCannotLoadFromNetwork: A specific request to load an item only from the cache couldn’t be satisfied.
	NSURLErrorCannotLoadFromNetwork NSURLError = -2000
	// NSURLErrorCannotMoveFile: A downloaded file on disk couldn’t be moved.
	NSURLErrorCannotMoveFile NSURLError = -3005
	// NSURLErrorCannotOpenFile: A downloaded file on disk couldn’t be opened.
	NSURLErrorCannotOpenFile NSURLError = -3001
	// NSURLErrorCannotParseResponse: A response to a connection request couldn’t be parsed.
	NSURLErrorCannotParseResponse NSURLError = -1017
	// NSURLErrorCannotRemoveFile: A downloaded file couldn’t be removed from disk.
	NSURLErrorCannotRemoveFile NSURLError = -3004
	// NSURLErrorCannotWriteToFile: A download task couldn’t write the file to disk.
	NSURLErrorCannotWriteToFile NSURLError = -3003
	// NSURLErrorClientCertificateRejected: A server certificate was rejected.
	NSURLErrorClientCertificateRejected NSURLError = -1205
	// NSURLErrorClientCertificateRequired: A client certificate was required to authenticate an SSL connection during a connection request.
	NSURLErrorClientCertificateRequired NSURLError = -1206
	// NSURLErrorDNSLookupFailed: The host address couldn’t be found via DNS lookup.
	NSURLErrorDNSLookupFailed NSURLError = -1006
	// NSURLErrorDataLengthExceedsMaximum: The length of the resource data exceeded the maximum allowed.
	NSURLErrorDataLengthExceedsMaximum NSURLError = -1101
	// NSURLErrorDataNotAllowed: The cellular network disallowed a connection.
	NSURLErrorDataNotAllowed NSURLError = -3004
	// NSURLErrorDownloadDecodingFailedMidStream: A download task failed to decode an encoded file during the download.
	NSURLErrorDownloadDecodingFailedMidStream NSURLError = -3006
	// NSURLErrorDownloadDecodingFailedToComplete: A download task failed to decode an encoded file after downloading.
	NSURLErrorDownloadDecodingFailedToComplete NSURLError = -3007
	// NSURLErrorFileDoesNotExist: The specified file doesn’t exist.
	NSURLErrorFileDoesNotExist NSURLError = -1100
	// NSURLErrorFileIsDirectory: A request for an FTP file resulted in the server responding that the file is not a plain file, but a directory.
	NSURLErrorFileIsDirectory NSURLError = -1101
	// NSURLErrorFileOutsideSafeArea: An internal file operation failed.
	NSURLErrorFileOutsideSafeArea NSURLError = -1100
	// NSURLErrorHTTPTooManyRedirects: A redirect loop was detected or the threshold for number of allowable redirects was exceeded (currently 16).
	NSURLErrorHTTPTooManyRedirects NSURLError = -1007
	// NSURLErrorInternationalRoamingOff: The attempted connection required activating a data context while roaming, but international roaming is disabled.
	NSURLErrorInternationalRoamingOff NSURLError = -3006
	// NSURLErrorNetworkConnectionLost: A client or server connection was severed in the middle of an in-progress load.
	NSURLErrorNetworkConnectionLost NSURLError = -1005
	// NSURLErrorNoPermissionsToReadFile: A resource couldn’t be read because of insufficient permissions.
	NSURLErrorNoPermissionsToReadFile NSURLError = -1102
	// NSURLErrorNotConnectedToInternet: A network resource was requested, but an internet connection has not been established and can’t be established automatically.
	NSURLErrorNotConnectedToInternet NSURLError = -1009
	// NSURLErrorRedirectToNonExistentLocation: A redirect was specified by way of server response code, but the server didn’t accompany this code with a redirect URL.
	NSURLErrorRedirectToNonExistentLocation NSURLError = -1010
	// NSURLErrorRequestBodyStreamExhausted: A body stream was needed but the client did not provide one.
	NSURLErrorRequestBodyStreamExhausted NSURLError = -3003
	// NSURLErrorResourceUnavailable: A requested resource couldn’t be retrieved.
	NSURLErrorResourceUnavailable NSURLError = -1008
	// NSURLErrorSecureConnectionFailed: An attempt to establish a secure connection failed for reasons that can’t be expressed more specifically.
	NSURLErrorSecureConnectionFailed NSURLError = -1200
	// NSURLErrorServerCertificateHasBadDate: A server certificate is expired, or is not yet valid.
	NSURLErrorServerCertificateHasBadDate NSURLError = -1201
	// NSURLErrorServerCertificateHasUnknownRoot: A server certificate wasn’t signed by any root server.
	NSURLErrorServerCertificateHasUnknownRoot NSURLError = -1203
	// NSURLErrorServerCertificateNotYetValid: A server certificate isn’t valid yet.
	NSURLErrorServerCertificateNotYetValid NSURLError = -1204
	// NSURLErrorServerCertificateUntrusted: A server certificate was signed by a root server that isn’t trusted.
	NSURLErrorServerCertificateUntrusted NSURLError = -1202
	// NSURLErrorTimedOut: An asynchronous operation timed out.
	NSURLErrorTimedOut NSURLError = -1001
	// NSURLErrorUnknown: The URL Loading System encountered an error that it can’t interpret.
	NSURLErrorUnknown NSURLError = -1
	// NSURLErrorUnsupportedURL: A properly formed URL couldn’t be handled by the framework.
	NSURLErrorUnsupportedURL NSURLError = -1002
	// NSURLErrorUserAuthenticationRequired: Authentication was required to access a resource.
	NSURLErrorUserAuthenticationRequired NSURLError = -1013
	// NSURLErrorUserCancelledAuthentication: An asynchronous request for authentication has been canceled by the user.
	NSURLErrorUserCancelledAuthentication NSURLError = -1012
	// NSURLErrorZeroByteResource: A server reported that a URL has a non-zero content length, but terminated the network connection gracefully without sending any data.
	NSURLErrorZeroByteResource NSURLError = -1014
)

func (e NSURLError) String() string {
	switch e {
	case NSURLErrorAppTransportSecurityRequiresSecureConnection:
		return "NSURLErrorAppTransportSecurityRequiresSecureConnection"
	case NSURLErrorBackgroundSessionInUseByAnotherProcess:
		return "NSURLErrorBackgroundSessionInUseByAnotherProcess"
	case NSURLErrorBackgroundSessionRequiresSharedContainer:
		return "NSURLErrorBackgroundSessionRequiresSharedContainer"
	case NSURLErrorBackgroundSessionWasDisconnected:
		return "NSURLErrorBackgroundSessionWasDisconnected"
	case NSURLErrorBadServerResponse:
		return "NSURLErrorBadServerResponse"
	case NSURLErrorBadURL:
		return "NSURLErrorBadURL"
	case NSURLErrorCallIsActive:
		return "NSURLErrorCallIsActive"
	case NSURLErrorCancelled:
		return "NSURLErrorCancelled"
	case NSURLErrorCannotConnectToHost:
		return "NSURLErrorCannotConnectToHost"
	case NSURLErrorCannotDecodeRawData:
		return "NSURLErrorCannotDecodeRawData"
	case NSURLErrorCannotFindHost:
		return "NSURLErrorCannotFindHost"
	case NSURLErrorCannotLoadFromNetwork:
		return "NSURLErrorCannotLoadFromNetwork"
	case NSURLErrorCannotParseResponse:
		return "NSURLErrorCannotParseResponse"
	case NSURLErrorCannotRemoveFile:
		return "NSURLErrorCannotRemoveFile"
	case NSURLErrorCannotWriteToFile:
		return "NSURLErrorCannotWriteToFile"
	case NSURLErrorClientCertificateRejected:
		return "NSURLErrorClientCertificateRejected"
	case NSURLErrorClientCertificateRequired:
		return "NSURLErrorClientCertificateRequired"
	case NSURLErrorDNSLookupFailed:
		return "NSURLErrorDNSLookupFailed"
	case NSURLErrorDataLengthExceedsMaximum:
		return "NSURLErrorDataLengthExceedsMaximum"
	case NSURLErrorDownloadDecodingFailedMidStream:
		return "NSURLErrorDownloadDecodingFailedMidStream"
	case NSURLErrorDownloadDecodingFailedToComplete:
		return "NSURLErrorDownloadDecodingFailedToComplete"
	case NSURLErrorFileDoesNotExist:
		return "NSURLErrorFileDoesNotExist"
	case NSURLErrorHTTPTooManyRedirects:
		return "NSURLErrorHTTPTooManyRedirects"
	case NSURLErrorNetworkConnectionLost:
		return "NSURLErrorNetworkConnectionLost"
	case NSURLErrorNoPermissionsToReadFile:
		return "NSURLErrorNoPermissionsToReadFile"
	case NSURLErrorNotConnectedToInternet:
		return "NSURLErrorNotConnectedToInternet"
	case NSURLErrorRedirectToNonExistentLocation:
		return "NSURLErrorRedirectToNonExistentLocation"
	case NSURLErrorResourceUnavailable:
		return "NSURLErrorResourceUnavailable"
	case NSURLErrorSecureConnectionFailed:
		return "NSURLErrorSecureConnectionFailed"
	case NSURLErrorServerCertificateHasBadDate:
		return "NSURLErrorServerCertificateHasBadDate"
	case NSURLErrorServerCertificateHasUnknownRoot:
		return "NSURLErrorServerCertificateHasUnknownRoot"
	case NSURLErrorServerCertificateNotYetValid:
		return "NSURLErrorServerCertificateNotYetValid"
	case NSURLErrorServerCertificateUntrusted:
		return "NSURLErrorServerCertificateUntrusted"
	case NSURLErrorTimedOut:
		return "NSURLErrorTimedOut"
	case NSURLErrorUnknown:
		return "NSURLErrorUnknown"
	case NSURLErrorUnsupportedURL:
		return "NSURLErrorUnsupportedURL"
	case NSURLErrorUserAuthenticationRequired:
		return "NSURLErrorUserAuthenticationRequired"
	case NSURLErrorUserCancelledAuthentication:
		return "NSURLErrorUserCancelledAuthentication"
	case NSURLErrorZeroByteResource:
		return "NSURLErrorZeroByteResource"
	default:
		return fmt.Sprintf("NSURLError(%d)", e)
	}
}

type NSURLErrorCancelledReason int

const (
	// NSURLErrorCancelledReasonBackgroundUpdatesDisabled: A reason that indicates the system canceled the background task because background tasks are disabled.
	NSURLErrorCancelledReasonBackgroundUpdatesDisabled NSURLErrorCancelledReason = 1
	// NSURLErrorCancelledReasonInsufficientSystemResources: A reason that indicates the system canceled the background task because it lacks sufficient resources to perform the task.
	NSURLErrorCancelledReasonInsufficientSystemResources NSURLErrorCancelledReason = 2
	// NSURLErrorCancelledReasonUserForceQuitApplication: A reason that indicates the system canceled the background task because the user force-quit the application.
	NSURLErrorCancelledReasonUserForceQuitApplication NSURLErrorCancelledReason = 0
)

func (e NSURLErrorCancelledReason) String() string {
	switch e {
	case NSURLErrorCancelledReasonBackgroundUpdatesDisabled:
		return "NSURLErrorCancelledReasonBackgroundUpdatesDisabled"
	case NSURLErrorCancelledReasonInsufficientSystemResources:
		return "NSURLErrorCancelledReasonInsufficientSystemResources"
	case NSURLErrorCancelledReasonUserForceQuitApplication:
		return "NSURLErrorCancelledReasonUserForceQuitApplication"
	default:
		return fmt.Sprintf("NSURLErrorCancelledReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURLErrorNetworkUnavailableReason
type NSURLErrorNetworkUnavailableReason int

const (
	// NSURLErrorNetworkUnavailableReasonCellular: A reason that indicates network is unavailable because the interface is cellular and cellular network is disabled.
	NSURLErrorNetworkUnavailableReasonCellular NSURLErrorNetworkUnavailableReason = 0
	// NSURLErrorNetworkUnavailableReasonConstrained: A reason that indicates network is unavailable because the user enabled “Low Data Mode” in the Settings app.
	NSURLErrorNetworkUnavailableReasonConstrained NSURLErrorNetworkUnavailableReason = 2
	// NSURLErrorNetworkUnavailableReasonExpensive: A reason that indicates network is unavailable because the system marked the interface as expensive.
	NSURLErrorNetworkUnavailableReasonExpensive NSURLErrorNetworkUnavailableReason = 1
	NSURLErrorNetworkUnavailableReasonUltraConstrained NSURLErrorNetworkUnavailableReason = 3
)

func (e NSURLErrorNetworkUnavailableReason) String() string {
	switch e {
	case NSURLErrorNetworkUnavailableReasonCellular:
		return "NSURLErrorNetworkUnavailableReasonCellular"
	case NSURLErrorNetworkUnavailableReasonConstrained:
		return "NSURLErrorNetworkUnavailableReasonConstrained"
	case NSURLErrorNetworkUnavailableReasonExpensive:
		return "NSURLErrorNetworkUnavailableReasonExpensive"
	case NSURLErrorNetworkUnavailableReasonUltraConstrained:
		return "NSURLErrorNetworkUnavailableReasonUltraConstrained"
	default:
		return fmt.Sprintf("NSURLErrorNetworkUnavailableReason(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURLHandle/Status-swift.enum
type NSURLHandleStatus int

const (
	// NSURLHandleLoadFailed: The resource data failed to load.
	NSURLHandleLoadFailed NSURLHandleStatus = 3
	// NSURLHandleLoadInProgress: The resource data is in the process of loading.
	NSURLHandleLoadInProgress NSURLHandleStatus = 2
	// NSURLHandleLoadSucceeded: The resource data was successfully loaded.
	NSURLHandleLoadSucceeded NSURLHandleStatus = 1
	// NSURLHandleNotLoaded: The resource data has not been loaded.
	NSURLHandleNotLoaded NSURLHandleStatus = 0
)

func (e NSURLHandleStatus) String() string {
	switch e {
	case NSURLHandleLoadFailed:
		return "NSURLHandleLoadFailed"
	case NSURLHandleLoadInProgress:
		return "NSURLHandleLoadInProgress"
	case NSURLHandleLoadSucceeded:
		return "NSURLHandleLoadSucceeded"
	case NSURLHandleNotLoaded:
		return "NSURLHandleNotLoaded"
	default:
		return fmt.Sprintf("NSURLHandleStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/URLRelationship
type NSURLRelationship int

const (
	// NSURLRelationshipContains: The directory contains the specified item.
	NSURLRelationshipContains NSURLRelationship = 0
	// NSURLRelationshipOther: The directory does not contain the item and is not the same as the item.
	NSURLRelationshipOther NSURLRelationship = 2
	// NSURLRelationshipSame: The directory and the item are the same.
	NSURLRelationshipSame NSURLRelationship = 1
)

func (e NSURLRelationship) String() string {
	switch e {
	case NSURLRelationshipContains:
		return "NSURLRelationshipContains"
	case NSURLRelationshipOther:
		return "NSURLRelationshipOther"
	case NSURLRelationshipSame:
		return "NSURLRelationshipSame"
	default:
		return fmt.Sprintf("NSURLRelationship(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/Attribution-swift.enum
type NSURLRequestAttribution int

const (
	// NSURLRequestAttributionDeveloper: A developer-initiated network request.
	NSURLRequestAttributionDeveloper NSURLRequestAttribution = 0
	// NSURLRequestAttributionUser: The user explicitly directs the app to make a network request.
	NSURLRequestAttributionUser NSURLRequestAttribution = 1
)

func (e NSURLRequestAttribution) String() string {
	switch e {
	case NSURLRequestAttributionDeveloper:
		return "NSURLRequestAttributionDeveloper"
	case NSURLRequestAttributionUser:
		return "NSURLRequestAttributionUser"
	default:
		return fmt.Sprintf("NSURLRequestAttribution(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/CachePolicy-swift.enum
type NSURLRequestCachePolicy int

const (
	// NSURLRequestReloadIgnoringLocalCacheData: The URL load should be loaded only from the originating source.
	NSURLRequestReloadIgnoringLocalCacheData NSURLRequestCachePolicy = 1
	// NSURLRequestReturnCacheDataDontLoad: Use existing cache data, regardless or age or expiration date, and fail if no cached data is available.
	NSURLRequestReturnCacheDataDontLoad NSURLRequestCachePolicy = 3
	// NSURLRequestReturnCacheDataElseLoad: Use existing cache data, regardless or age or expiration date, loading from originating source only if there is no cached data.
	NSURLRequestReturnCacheDataElseLoad NSURLRequestCachePolicy = 2
	// NSURLRequestUseProtocolCachePolicy: Use the caching logic defined in the protocol implementation, if any, for a particular URL load request.
	NSURLRequestUseProtocolCachePolicy NSURLRequestCachePolicy = 0
)

func (e NSURLRequestCachePolicy) String() string {
	switch e {
	case NSURLRequestReloadIgnoringLocalCacheData:
		return "NSURLRequestReloadIgnoringLocalCacheData"
	case NSURLRequestReturnCacheDataDontLoad:
		return "NSURLRequestReturnCacheDataDontLoad"
	case NSURLRequestReturnCacheDataElseLoad:
		return "NSURLRequestReturnCacheDataElseLoad"
	case NSURLRequestUseProtocolCachePolicy:
		return "NSURLRequestUseProtocolCachePolicy"
	default:
		return fmt.Sprintf("NSURLRequestCachePolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURLRequest/NetworkServiceType-swift.enum
type NSURLRequestNetworkServiceType int

const (
	// NSURLNetworkServiceTypeAVStreaming: A service type for medium-delay tolerant, low-medium-loss tolerant, elastic flow, constant packet interval, and variable rate and size connections.
	NSURLNetworkServiceTypeAVStreaming NSURLRequestNetworkServiceType = 8
	// NSURLNetworkServiceTypeBackground: A service type for high-delay tolerant, high-loss tolerant, elastic flow, and variable size connections.
	NSURLNetworkServiceTypeBackground NSURLRequestNetworkServiceType = 3
	// NSURLNetworkServiceTypeCallSignaling: A service for low-loss tolerant, inelastic flow, jitter tolerant, short but bursty rate, and variable size connections.
	NSURLNetworkServiceTypeCallSignaling NSURLRequestNetworkServiceType = 11
	// NSURLNetworkServiceTypeDefault: A service type for standard network traffic.
	NSURLNetworkServiceTypeDefault NSURLRequestNetworkServiceType = 0
	// NSURLNetworkServiceTypeResponsiveAV: A service type for low-delay tolerant, low-to-medium-loss tolerant, elastic flow, variable packet interval, rate, size responsive and time-sensitive connections.
	NSURLNetworkServiceTypeResponsiveAV NSURLRequestNetworkServiceType = 9
	// NSURLNetworkServiceTypeResponsiveData: A service type for medium-delay tolerant, elastic and inelastic flow, bursty, and long-lived connections.
	NSURLNetworkServiceTypeResponsiveData NSURLRequestNetworkServiceType = 6
	// NSURLNetworkServiceTypeVideo: A service type for low-delay tolerant, very low-loss tolerant, inelastic flow, and constant packet rate connections.
	NSURLNetworkServiceTypeVideo NSURLRequestNetworkServiceType = 2
	// NSURLNetworkServiceTypeVoIP: A service type for VoIP traffic.
	NSURLNetworkServiceTypeVoIP NSURLRequestNetworkServiceType = 1
	// NSURLNetworkServiceTypeVoice: A service type for low-delay tolerant, very low-loss tolerant, inelastic flow, and constant packet rate connections.
	NSURLNetworkServiceTypeVoice NSURLRequestNetworkServiceType = 4
)

func (e NSURLRequestNetworkServiceType) String() string {
	switch e {
	case NSURLNetworkServiceTypeAVStreaming:
		return "NSURLNetworkServiceTypeAVStreaming"
	case NSURLNetworkServiceTypeBackground:
		return "NSURLNetworkServiceTypeBackground"
	case NSURLNetworkServiceTypeCallSignaling:
		return "NSURLNetworkServiceTypeCallSignaling"
	case NSURLNetworkServiceTypeDefault:
		return "NSURLNetworkServiceTypeDefault"
	case NSURLNetworkServiceTypeResponsiveAV:
		return "NSURLNetworkServiceTypeResponsiveAV"
	case NSURLNetworkServiceTypeResponsiveData:
		return "NSURLNetworkServiceTypeResponsiveData"
	case NSURLNetworkServiceTypeVideo:
		return "NSURLNetworkServiceTypeVideo"
	case NSURLNetworkServiceTypeVoIP:
		return "NSURLNetworkServiceTypeVoIP"
	case NSURLNetworkServiceTypeVoice:
		return "NSURLNetworkServiceTypeVoice"
	default:
		return fmt.Sprintf("NSURLRequestNetworkServiceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSession/AuthChallengeDisposition
type NSURLSessionAuthChallengeDisposition int

const (
	// NSURLSessionAuthChallengeCancelAuthenticationChallenge: Cancel the entire request.
	NSURLSessionAuthChallengeCancelAuthenticationChallenge NSURLSessionAuthChallengeDisposition = 2
	// NSURLSessionAuthChallengePerformDefaultHandling: Use the default handling for the challenge as though this delegate method were not implemented.
	NSURLSessionAuthChallengePerformDefaultHandling NSURLSessionAuthChallengeDisposition = 1
	// NSURLSessionAuthChallengeUseCredential: Use the specified credential, which may be `nil`.
	NSURLSessionAuthChallengeUseCredential NSURLSessionAuthChallengeDisposition = 0
)

func (e NSURLSessionAuthChallengeDisposition) String() string {
	switch e {
	case NSURLSessionAuthChallengeCancelAuthenticationChallenge:
		return "NSURLSessionAuthChallengeCancelAuthenticationChallenge"
	case NSURLSessionAuthChallengePerformDefaultHandling:
		return "NSURLSessionAuthChallengePerformDefaultHandling"
	case NSURLSessionAuthChallengeUseCredential:
		return "NSURLSessionAuthChallengeUseCredential"
	default:
		return fmt.Sprintf("NSURLSessionAuthChallengeDisposition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSession/DelayedRequestDisposition
type NSURLSessionDelayedRequestDisposition int

const (
	// NSURLSessionDelayedRequestCancel: A disposition indicating that the task should be canceled.
	NSURLSessionDelayedRequestCancel NSURLSessionDelayedRequestDisposition = 2
	// NSURLSessionDelayedRequestContinueLoading: A disposition indicating that the task should proceed with the original request.
	NSURLSessionDelayedRequestContinueLoading NSURLSessionDelayedRequestDisposition = 0
	// NSURLSessionDelayedRequestUseNewRequest: A disposition indicating that the task should use a new request to perform the network load.
	NSURLSessionDelayedRequestUseNewRequest NSURLSessionDelayedRequestDisposition = 1
)

func (e NSURLSessionDelayedRequestDisposition) String() string {
	switch e {
	case NSURLSessionDelayedRequestCancel:
		return "NSURLSessionDelayedRequestCancel"
	case NSURLSessionDelayedRequestContinueLoading:
		return "NSURLSessionDelayedRequestContinueLoading"
	case NSURLSessionDelayedRequestUseNewRequest:
		return "NSURLSessionDelayedRequestUseNewRequest"
	default:
		return fmt.Sprintf("NSURLSessionDelayedRequestDisposition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSessionConfiguration/MultipathServiceType-swift.enum
type NSURLSessionMultipathServiceType int

const (
	// NSURLSessionMultipathServiceTypeAggregate: A service that aggregates the capacities of other Multipath options in an attempt to increase throughput and minimize latency.
	NSURLSessionMultipathServiceTypeAggregate NSURLSessionMultipathServiceType = 3
	// NSURLSessionMultipathServiceTypeHandover: A Multipath TCP service that provides seamless handover between Wi-Fi and cellular in order to preserve the connection.
	NSURLSessionMultipathServiceTypeHandover NSURLSessionMultipathServiceType = 1
	// NSURLSessionMultipathServiceTypeInteractive: A service whereby Multipath TCP attempts to use the lowest-latency interface.
	NSURLSessionMultipathServiceTypeInteractive NSURLSessionMultipathServiceType = 2
	// NSURLSessionMultipathServiceTypeNone: The default service type indicating that Multipath TCP should not be used.
	NSURLSessionMultipathServiceTypeNone NSURLSessionMultipathServiceType = 0
)

func (e NSURLSessionMultipathServiceType) String() string {
	switch e {
	case NSURLSessionMultipathServiceTypeAggregate:
		return "NSURLSessionMultipathServiceTypeAggregate"
	case NSURLSessionMultipathServiceTypeHandover:
		return "NSURLSessionMultipathServiceTypeHandover"
	case NSURLSessionMultipathServiceTypeInteractive:
		return "NSURLSessionMultipathServiceTypeInteractive"
	case NSURLSessionMultipathServiceTypeNone:
		return "NSURLSessionMultipathServiceTypeNone"
	default:
		return fmt.Sprintf("NSURLSessionMultipathServiceType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSession/ResponseDisposition
type NSURLSessionResponseDisposition int

const (
	// NSURLSessionResponseAllow: Allow the load operation to continue.
	NSURLSessionResponseAllow NSURLSessionResponseDisposition = 1
	// NSURLSessionResponseBecomeDownload: Convert the response for this request to use a .
	NSURLSessionResponseBecomeDownload NSURLSessionResponseDisposition = 2
	// NSURLSessionResponseBecomeStream: Convert the response for this request to use a .
	NSURLSessionResponseBecomeStream NSURLSessionResponseDisposition = 3
	// NSURLSessionResponseCancel: Cancel the load.
	NSURLSessionResponseCancel NSURLSessionResponseDisposition = 0
)

func (e NSURLSessionResponseDisposition) String() string {
	switch e {
	case NSURLSessionResponseAllow:
		return "NSURLSessionResponseAllow"
	case NSURLSessionResponseBecomeDownload:
		return "NSURLSessionResponseBecomeDownload"
	case NSURLSessionResponseBecomeStream:
		return "NSURLSessionResponseBecomeStream"
	case NSURLSessionResponseCancel:
		return "NSURLSessionResponseCancel"
	default:
		return fmt.Sprintf("NSURLSessionResponseDisposition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/DomainResolutionProtocol
type NSURLSessionTaskMetricsDomainResolutionProtocol int

const (
	NSURLSessionTaskMetricsDomainResolutionProtocolHTTPS NSURLSessionTaskMetricsDomainResolutionProtocol = 4
	NSURLSessionTaskMetricsDomainResolutionProtocolTCP NSURLSessionTaskMetricsDomainResolutionProtocol = 2
	NSURLSessionTaskMetricsDomainResolutionProtocolTLS NSURLSessionTaskMetricsDomainResolutionProtocol = 3
	NSURLSessionTaskMetricsDomainResolutionProtocolUDP NSURLSessionTaskMetricsDomainResolutionProtocol = 1
	NSURLSessionTaskMetricsDomainResolutionProtocolUnknown NSURLSessionTaskMetricsDomainResolutionProtocol = 0
)

func (e NSURLSessionTaskMetricsDomainResolutionProtocol) String() string {
	switch e {
	case NSURLSessionTaskMetricsDomainResolutionProtocolHTTPS:
		return "NSURLSessionTaskMetricsDomainResolutionProtocolHTTPS"
	case NSURLSessionTaskMetricsDomainResolutionProtocolTCP:
		return "NSURLSessionTaskMetricsDomainResolutionProtocolTCP"
	case NSURLSessionTaskMetricsDomainResolutionProtocolTLS:
		return "NSURLSessionTaskMetricsDomainResolutionProtocolTLS"
	case NSURLSessionTaskMetricsDomainResolutionProtocolUDP:
		return "NSURLSessionTaskMetricsDomainResolutionProtocolUDP"
	case NSURLSessionTaskMetricsDomainResolutionProtocolUnknown:
		return "NSURLSessionTaskMetricsDomainResolutionProtocolUnknown"
	default:
		return fmt.Sprintf("NSURLSessionTaskMetricsDomainResolutionProtocol(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSessionTaskMetrics/ResourceFetchType
type NSURLSessionTaskMetricsResourceFetchType int

const (
	// NSURLSessionTaskMetricsResourceFetchTypeLocalCache: The resource was retrieved from the local storage.
	NSURLSessionTaskMetricsResourceFetchTypeLocalCache NSURLSessionTaskMetricsResourceFetchType = 3
	// NSURLSessionTaskMetricsResourceFetchTypeNetworkLoad: The resource was loaded over the network.
	NSURLSessionTaskMetricsResourceFetchTypeNetworkLoad NSURLSessionTaskMetricsResourceFetchType = 1
	// NSURLSessionTaskMetricsResourceFetchTypeServerPush: The resource was pushed by the server to the client.
	NSURLSessionTaskMetricsResourceFetchTypeServerPush NSURLSessionTaskMetricsResourceFetchType = 2
	// NSURLSessionTaskMetricsResourceFetchTypeUnknown: The manner in which the resource was fetched could not be determined.
	NSURLSessionTaskMetricsResourceFetchTypeUnknown NSURLSessionTaskMetricsResourceFetchType = 0
)

func (e NSURLSessionTaskMetricsResourceFetchType) String() string {
	switch e {
	case NSURLSessionTaskMetricsResourceFetchTypeLocalCache:
		return "NSURLSessionTaskMetricsResourceFetchTypeLocalCache"
	case NSURLSessionTaskMetricsResourceFetchTypeNetworkLoad:
		return "NSURLSessionTaskMetricsResourceFetchTypeNetworkLoad"
	case NSURLSessionTaskMetricsResourceFetchTypeServerPush:
		return "NSURLSessionTaskMetricsResourceFetchTypeServerPush"
	case NSURLSessionTaskMetricsResourceFetchTypeUnknown:
		return "NSURLSessionTaskMetricsResourceFetchTypeUnknown"
	default:
		return fmt.Sprintf("NSURLSessionTaskMetricsResourceFetchType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSessionTask/State-swift.enum
type NSURLSessionTaskState int

const (
	// NSURLSessionTaskStateCanceling: The task has received a  message.
	NSURLSessionTaskStateCanceling NSURLSessionTaskState = 2
	// NSURLSessionTaskStateCompleted: The task has completed (without being canceled), and the task’s delegate receives no further callbacks.
	NSURLSessionTaskStateCompleted NSURLSessionTaskState = 3
	// NSURLSessionTaskStateRunning: The task is currently being serviced by the session.
	NSURLSessionTaskStateRunning NSURLSessionTaskState = 0
	// NSURLSessionTaskStateSuspended: The task was suspended by the app.
	NSURLSessionTaskStateSuspended NSURLSessionTaskState = 1
)

func (e NSURLSessionTaskState) String() string {
	switch e {
	case NSURLSessionTaskStateCanceling:
		return "NSURLSessionTaskStateCanceling"
	case NSURLSessionTaskStateCompleted:
		return "NSURLSessionTaskStateCompleted"
	case NSURLSessionTaskStateRunning:
		return "NSURLSessionTaskStateRunning"
	case NSURLSessionTaskStateSuspended:
		return "NSURLSessionTaskStateSuspended"
	default:
		return fmt.Sprintf("NSURLSessionTaskState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/URLSessionWebSocketTask/CloseCode-swift.enum
type NSURLSessionWebSocketCloseCode int

const (
	// NSURLSessionWebSocketCloseCodeAbnormalClosure: A reserved code that indicates the connection closed without a close control frame.
	NSURLSessionWebSocketCloseCodeAbnormalClosure NSURLSessionWebSocketCloseCode = 1006
	// NSURLSessionWebSocketCloseCodeGoingAway: A code that indicates an endpoint is going away.
	NSURLSessionWebSocketCloseCodeGoingAway NSURLSessionWebSocketCloseCode = 1001
	// NSURLSessionWebSocketCloseCodeInternalServerError: A code that indicates the server terminated the connection because it encountered an unexpected condition.
	NSURLSessionWebSocketCloseCodeInternalServerError NSURLSessionWebSocketCloseCode = 1011
	// NSURLSessionWebSocketCloseCodeInvalid: A code that indicates the connection is still open.
	NSURLSessionWebSocketCloseCodeInvalid NSURLSessionWebSocketCloseCode = 0
	// NSURLSessionWebSocketCloseCodeInvalidFramePayloadData: A code that indicates the server terminated the connection because it received data inconsistent with the message’s type.
	NSURLSessionWebSocketCloseCodeInvalidFramePayloadData NSURLSessionWebSocketCloseCode = 1007
	// NSURLSessionWebSocketCloseCodeMandatoryExtensionMissing: A code that indicates the client terminated the connection because the server didn’t negotiate a required extension.
	NSURLSessionWebSocketCloseCodeMandatoryExtensionMissing NSURLSessionWebSocketCloseCode = 1010
	// NSURLSessionWebSocketCloseCodeMessageTooBig: A code that indicates an endpoint is terminating the connection because it received a message too big for it to process.
	NSURLSessionWebSocketCloseCodeMessageTooBig NSURLSessionWebSocketCloseCode = 1009
	// NSURLSessionWebSocketCloseCodeNoStatusReceived: A reserved code that indicates an endpoint expected a status code and didn’t receive one.
	NSURLSessionWebSocketCloseCodeNoStatusReceived NSURLSessionWebSocketCloseCode = 1005
	// NSURLSessionWebSocketCloseCodeNormalClosure: A code that indicates normal connection closure.
	NSURLSessionWebSocketCloseCodeNormalClosure NSURLSessionWebSocketCloseCode = 1000
	// NSURLSessionWebSocketCloseCodePolicyViolation: A code that indicates an endpoint terminated the connection because it received a message that violates its policy.
	NSURLSessionWebSocketCloseCodePolicyViolation NSURLSessionWebSocketCloseCode = 1008
	// NSURLSessionWebSocketCloseCodeProtocolError: A code that indicates an endpoint terminated the connection due to a protocol error.
	NSURLSessionWebSocketCloseCodeProtocolError NSURLSessionWebSocketCloseCode = 1002
	// NSURLSessionWebSocketCloseCodeTLSHandshakeFailure: A reserved code that indicates the connection closed due to the failure to perform a TLS handshake.
	NSURLSessionWebSocketCloseCodeTLSHandshakeFailure NSURLSessionWebSocketCloseCode = 1015
	// NSURLSessionWebSocketCloseCodeUnsupportedData: A code that indicates an endpoint terminated the connection after receiving a type of data it can’t accept.
	NSURLSessionWebSocketCloseCodeUnsupportedData NSURLSessionWebSocketCloseCode = 1003
)

func (e NSURLSessionWebSocketCloseCode) String() string {
	switch e {
	case NSURLSessionWebSocketCloseCodeAbnormalClosure:
		return "NSURLSessionWebSocketCloseCodeAbnormalClosure"
	case NSURLSessionWebSocketCloseCodeGoingAway:
		return "NSURLSessionWebSocketCloseCodeGoingAway"
	case NSURLSessionWebSocketCloseCodeInternalServerError:
		return "NSURLSessionWebSocketCloseCodeInternalServerError"
	case NSURLSessionWebSocketCloseCodeInvalid:
		return "NSURLSessionWebSocketCloseCodeInvalid"
	case NSURLSessionWebSocketCloseCodeInvalidFramePayloadData:
		return "NSURLSessionWebSocketCloseCodeInvalidFramePayloadData"
	case NSURLSessionWebSocketCloseCodeMandatoryExtensionMissing:
		return "NSURLSessionWebSocketCloseCodeMandatoryExtensionMissing"
	case NSURLSessionWebSocketCloseCodeMessageTooBig:
		return "NSURLSessionWebSocketCloseCodeMessageTooBig"
	case NSURLSessionWebSocketCloseCodeNoStatusReceived:
		return "NSURLSessionWebSocketCloseCodeNoStatusReceived"
	case NSURLSessionWebSocketCloseCodeNormalClosure:
		return "NSURLSessionWebSocketCloseCodeNormalClosure"
	case NSURLSessionWebSocketCloseCodePolicyViolation:
		return "NSURLSessionWebSocketCloseCodePolicyViolation"
	case NSURLSessionWebSocketCloseCodeProtocolError:
		return "NSURLSessionWebSocketCloseCodeProtocolError"
	case NSURLSessionWebSocketCloseCodeTLSHandshakeFailure:
		return "NSURLSessionWebSocketCloseCodeTLSHandshakeFailure"
	case NSURLSessionWebSocketCloseCodeUnsupportedData:
		return "NSURLSessionWebSocketCloseCodeUnsupportedData"
	default:
		return fmt.Sprintf("NSURLSessionWebSocketCloseCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSURLSessionWebSocketMessageType
type NSURLSessionWebSocketMessageType int

const (
	NSURLSessionWebSocketMessageTypeData NSURLSessionWebSocketMessageType = 0
)

func (e NSURLSessionWebSocketMessageType) String() string {
	switch e {
	case NSURLSessionWebSocketMessageTypeData:
		return "NSURLSessionWebSocketMessageTypeData"
	default:
		return fmt.Sprintf("NSURLSessionWebSocketMessageType(%d)", e)
	}
}

const (
	// NSUbiquitousKeyValueStoreAccountChange: A constant that indicates the current Apple account changed.
	NSUbiquitousKeyValueStoreAccountChange uint = 3
	// NSUbiquitousKeyValueStoreInitialSyncChange: A constant that indicates the initial attempt to load keys and values from iCloud is in progress.
	NSUbiquitousKeyValueStoreInitialSyncChange uint = 1
	// NSUbiquitousKeyValueStoreQuotaViolationChange: A constant that indicates an attempt to write data exceeded the quota limits.
	NSUbiquitousKeyValueStoreQuotaViolationChange uint = 2
	// NSUbiquitousKeyValueStoreServerChange: A constant that indicates a value changed in iCloud.
	NSUbiquitousKeyValueStoreServerChange uint = 0
)

// See: https://developer.apple.com/documentation/Foundation/NSUserNotification/ActivationType-swift.enum
type NSUserNotificationActivationType int

const (
	// Deprecated.
	NSUserNotificationActivationTypeActionButtonClicked NSUserNotificationActivationType = 2
	// Deprecated.
	NSUserNotificationActivationTypeAdditionalActionClicked NSUserNotificationActivationType = 4
	// Deprecated.
	NSUserNotificationActivationTypeContentsClicked NSUserNotificationActivationType = 1
	// Deprecated.
	NSUserNotificationActivationTypeNone NSUserNotificationActivationType = 0
	// Deprecated.
	NSUserNotificationActivationTypeReplied NSUserNotificationActivationType = 3
)

func (e NSUserNotificationActivationType) String() string {
	switch e {
	case NSUserNotificationActivationTypeActionButtonClicked:
		return "NSUserNotificationActivationTypeActionButtonClicked"
	case NSUserNotificationActivationTypeAdditionalActionClicked:
		return "NSUserNotificationActivationTypeAdditionalActionClicked"
	case NSUserNotificationActivationTypeContentsClicked:
		return "NSUserNotificationActivationTypeContentsClicked"
	case NSUserNotificationActivationTypeNone:
		return "NSUserNotificationActivationTypeNone"
	case NSUserNotificationActivationTypeReplied:
		return "NSUserNotificationActivationTypeReplied"
	default:
		return fmt.Sprintf("NSUserNotificationActivationType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/FileManager/VolumeEnumerationOptions
type NSVolumeEnumerationOptions int

const (
	// NSVolumeEnumerationProduceFileReferenceURLs: The enumeration produces file reference URLs rather than path-based URLs.
	NSVolumeEnumerationProduceFileReferenceURLs NSVolumeEnumerationOptions = 4
	// NSVolumeEnumerationSkipHiddenVolumes: The enumeration skips hidden volumes.
	NSVolumeEnumerationSkipHiddenVolumes NSVolumeEnumerationOptions = 2
)

func (e NSVolumeEnumerationOptions) String() string {
	switch e {
	case NSVolumeEnumerationProduceFileReferenceURLs:
		return "NSVolumeEnumerationProduceFileReferenceURLs"
	case NSVolumeEnumerationSkipHiddenVolumes:
		return "NSVolumeEnumerationSkipHiddenVolumes"
	default:
		return fmt.Sprintf("NSVolumeEnumerationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSWhoseSpecifier/SubelementIdentifier
type NSWhoseSubelementIdentifier int

const (
	// NSEverySubelement: Every element that meets the specifier test.
	NSEverySubelement NSWhoseSubelementIdentifier = 1
	// NSIndexSubelement: An element at a given index that meets the specifier test.
	NSIndexSubelement NSWhoseSubelementIdentifier = 0
	// NSMiddleSubelement: The middle element that meets the specifier test.
	NSMiddleSubelement NSWhoseSubelementIdentifier = 2
	// NSNoSubelement: No sub-element met the specifier test.
	NSNoSubelement NSWhoseSubelementIdentifier = 4
	// NSRandomSubelement: Any element that meets the specifier test.
	NSRandomSubelement NSWhoseSubelementIdentifier = 3
)

func (e NSWhoseSubelementIdentifier) String() string {
	switch e {
	case NSEverySubelement:
		return "NSEverySubelement"
	case NSIndexSubelement:
		return "NSIndexSubelement"
	case NSMiddleSubelement:
		return "NSMiddleSubelement"
	case NSNoSubelement:
		return "NSNoSubelement"
	case NSRandomSubelement:
		return "NSRandomSubelement"
	default:
		return fmt.Sprintf("NSWhoseSubelementIdentifier(%d)", e)
	}
}

type NSWrapCalendar uint

const (
	// Deprecated.
	NSWrapCalendarComponents NSWrapCalendar = 1
)

func (e NSWrapCalendar) String() string {
	switch e {
	case NSWrapCalendarComponents:
		return "NSWrapCalendarComponents"
	default:
		return fmt.Sprintf("NSWrapCalendar(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/XMLDTDNode/DTDKind-swift.enum
type NSXMLDTDNodeKind int

const (
	// NSXMLAttributeCDATAKind: Identifies an attribute-list declaration with a [CDATA] (character data) value type.
	NSXMLAttributeCDATAKind NSXMLDTDNodeKind = 6
	// NSXMLAttributeEntitiesKind: Identifies an attribute-list declaration with an [ENTITIES] value type (refers to multiple unparsed entities declared elsewhere in document).
	NSXMLAttributeEntitiesKind NSXMLDTDNodeKind = 11
	// NSXMLAttributeEntityKind: Identifies an attribute-list declaration with an [ENTITY] value type (refers to unparsed entity declared in document).
	NSXMLAttributeEntityKind NSXMLDTDNodeKind = 10
	// NSXMLAttributeEnumerationKind: Identifies an attribute-list declaration with an enumeration value type (list of all possible values).
	NSXMLAttributeEnumerationKind NSXMLDTDNodeKind = 14
	// NSXMLAttributeIDKind: Identifies an attribute-list declaration with an [ID] value type (per-document unique element name).
	NSXMLAttributeIDKind NSXMLDTDNodeKind = 7
	// NSXMLAttributeIDRefKind: Identifies an attribute-list declaration with an [IDREF] value type (refers to element [ID] type).
	NSXMLAttributeIDRefKind NSXMLDTDNodeKind = 8
	// NSXMLAttributeIDRefsKind: Identifies an attribute-list declaration with an [IDREFS] value type (refers to multiple elements of [ID] type).
	NSXMLAttributeIDRefsKind NSXMLDTDNodeKind = 9
	// NSXMLAttributeNMTokenKind: Identifies an attribute-list declaration with a [NMTOKEN] value type (name token).
	NSXMLAttributeNMTokenKind NSXMLDTDNodeKind = 12
	// NSXMLAttributeNMTokensKind: Identifies an attribute-list declaration with a [NMTOKENS] value type (multiple name tokens)
	NSXMLAttributeNMTokensKind NSXMLDTDNodeKind = 13
	// NSXMLAttributeNotationKind: Identifies an attribute-list declaration with a [NOTATION] value type (name of declared notation).
	NSXMLAttributeNotationKind NSXMLDTDNodeKind = 15
	// NSXMLElementDeclarationAnyKind: Identifies an [ANY] element declaration.
	NSXMLElementDeclarationAnyKind NSXMLDTDNodeKind = 18
	// NSXMLElementDeclarationElementKind: Identifies a declaration of an element with child elements.
	NSXMLElementDeclarationElementKind NSXMLDTDNodeKind = 20
	// NSXMLElementDeclarationEmptyKind: Identifies a declaration ([EMPTY]) of an empty element.
	NSXMLElementDeclarationEmptyKind NSXMLDTDNodeKind = 17
	// NSXMLElementDeclarationMixedKind: Identifies a declaration of an element with mixed content (`(#PCDATA | child)`).
	NSXMLElementDeclarationMixedKind NSXMLDTDNodeKind = 19
	// NSXMLElementDeclarationUndefinedKind: Identifies an undefined element declaration.
	NSXMLElementDeclarationUndefinedKind NSXMLDTDNodeKind = 16
	// NSXMLEntityGeneralKind: Identifies a general entity declaration.
	NSXMLEntityGeneralKind NSXMLDTDNodeKind = 1
	// NSXMLEntityParameterKind: Identifies a parameter entity declaration.
	NSXMLEntityParameterKind NSXMLDTDNodeKind = 4
	// NSXMLEntityParsedKind: Identifies a parsed entity declaration.
	NSXMLEntityParsedKind NSXMLDTDNodeKind = 2
	// NSXMLEntityPredefined: Identifies a predefined entity declaration.
	NSXMLEntityPredefined NSXMLDTDNodeKind = 5
	// NSXMLEntityUnparsedKind: Identifies an unparsed entity declaration.
	NSXMLEntityUnparsedKind NSXMLDTDNodeKind = 3
)

func (e NSXMLDTDNodeKind) String() string {
	switch e {
	case NSXMLAttributeCDATAKind:
		return "NSXMLAttributeCDATAKind"
	case NSXMLAttributeEntitiesKind:
		return "NSXMLAttributeEntitiesKind"
	case NSXMLAttributeEntityKind:
		return "NSXMLAttributeEntityKind"
	case NSXMLAttributeEnumerationKind:
		return "NSXMLAttributeEnumerationKind"
	case NSXMLAttributeIDKind:
		return "NSXMLAttributeIDKind"
	case NSXMLAttributeIDRefKind:
		return "NSXMLAttributeIDRefKind"
	case NSXMLAttributeIDRefsKind:
		return "NSXMLAttributeIDRefsKind"
	case NSXMLAttributeNMTokenKind:
		return "NSXMLAttributeNMTokenKind"
	case NSXMLAttributeNMTokensKind:
		return "NSXMLAttributeNMTokensKind"
	case NSXMLAttributeNotationKind:
		return "NSXMLAttributeNotationKind"
	case NSXMLElementDeclarationAnyKind:
		return "NSXMLElementDeclarationAnyKind"
	case NSXMLElementDeclarationElementKind:
		return "NSXMLElementDeclarationElementKind"
	case NSXMLElementDeclarationEmptyKind:
		return "NSXMLElementDeclarationEmptyKind"
	case NSXMLElementDeclarationMixedKind:
		return "NSXMLElementDeclarationMixedKind"
	case NSXMLElementDeclarationUndefinedKind:
		return "NSXMLElementDeclarationUndefinedKind"
	case NSXMLEntityGeneralKind:
		return "NSXMLEntityGeneralKind"
	case NSXMLEntityParameterKind:
		return "NSXMLEntityParameterKind"
	case NSXMLEntityParsedKind:
		return "NSXMLEntityParsedKind"
	case NSXMLEntityPredefined:
		return "NSXMLEntityPredefined"
	case NSXMLEntityUnparsedKind:
		return "NSXMLEntityUnparsedKind"
	default:
		return fmt.Sprintf("NSXMLDTDNodeKind(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/XMLDocument/ContentKind
type NSXMLDocumentContentKind int

const (
	// NSXMLDocumentHTMLKind: Outputs empty tags in HTML without a close tag, such as `<br>`.
	NSXMLDocumentHTMLKind NSXMLDocumentContentKind = 2
	// NSXMLDocumentTextKind: Outputs the string value of the document by extracting the string values from all text nodes.
	NSXMLDocumentTextKind NSXMLDocumentContentKind = 3
	// NSXMLDocumentXHTMLKind: The document output is XHTML.
	NSXMLDocumentXHTMLKind NSXMLDocumentContentKind = 1
	// NSXMLDocumentXMLKind: The default type of document content type, which is XML.
	NSXMLDocumentXMLKind NSXMLDocumentContentKind = 0
)

func (e NSXMLDocumentContentKind) String() string {
	switch e {
	case NSXMLDocumentHTMLKind:
		return "NSXMLDocumentHTMLKind"
	case NSXMLDocumentTextKind:
		return "NSXMLDocumentTextKind"
	case NSXMLDocumentXHTMLKind:
		return "NSXMLDocumentXHTMLKind"
	case NSXMLDocumentXMLKind:
		return "NSXMLDocumentXMLKind"
	default:
		return fmt.Sprintf("NSXMLDocumentContentKind(%d)", e)
	}
}

type NSXMLNodeKind int

const (
	NSXMLAttributeDeclarationKind NSXMLNodeKind = 10
	NSXMLAttributeKind NSXMLNodeKind = 3
	NSXMLCommentKind NSXMLNodeKind = 6
	NSXMLDTDKind NSXMLNodeKind = 8
	NSXMLDocumentKind NSXMLNodeKind = 1
	NSXMLElementDeclarationKind NSXMLNodeKind = 11
	NSXMLElementKind NSXMLNodeKind = 2
	NSXMLEntityDeclarationKind NSXMLNodeKind = 9
	NSXMLInvalidKind NSXMLNodeKind = 0
	NSXMLNamespaceKind NSXMLNodeKind = 4
	NSXMLNotationDeclarationKind NSXMLNodeKind = 12
	NSXMLProcessingInstructionKind NSXMLNodeKind = 5
	NSXMLTextKind NSXMLNodeKind = 7
)

func (e NSXMLNodeKind) String() string {
	switch e {
	case NSXMLAttributeDeclarationKind:
		return "NSXMLAttributeDeclarationKind"
	case NSXMLAttributeKind:
		return "NSXMLAttributeKind"
	case NSXMLCommentKind:
		return "NSXMLCommentKind"
	case NSXMLDTDKind:
		return "NSXMLDTDKind"
	case NSXMLDocumentKind:
		return "NSXMLDocumentKind"
	case NSXMLElementDeclarationKind:
		return "NSXMLElementDeclarationKind"
	case NSXMLElementKind:
		return "NSXMLElementKind"
	case NSXMLEntityDeclarationKind:
		return "NSXMLEntityDeclarationKind"
	case NSXMLInvalidKind:
		return "NSXMLInvalidKind"
	case NSXMLNamespaceKind:
		return "NSXMLNamespaceKind"
	case NSXMLNotationDeclarationKind:
		return "NSXMLNotationDeclarationKind"
	case NSXMLProcessingInstructionKind:
		return "NSXMLProcessingInstructionKind"
	case NSXMLTextKind:
		return "NSXMLTextKind"
	default:
		return fmt.Sprintf("NSXMLNodeKind(%d)", e)
	}
}

type NSXMLNodeOptions int

const (
	NSXMLDocumentIncludeContentTypeDeclaration NSXMLNodeOptions = 262144
	NSXMLDocumentTidyHTML NSXMLNodeOptions = 512
	NSXMLDocumentTidyXML NSXMLNodeOptions = 1024
	NSXMLDocumentValidate NSXMLNodeOptions = 8192
	NSXMLDocumentXInclude NSXMLNodeOptions = 65536
	NSXMLNodeCompactEmptyElement NSXMLNodeOptions = 4
	NSXMLNodeExpandEmptyElement NSXMLNodeOptions = 2
	NSXMLNodeIsCDATA NSXMLNodeOptions = 1
	NSXMLNodeLoadExternalEntitiesAlways NSXMLNodeOptions = 16384
	NSXMLNodeLoadExternalEntitiesNever NSXMLNodeOptions = 524288
	NSXMLNodeLoadExternalEntitiesSameOriginOnly NSXMLNodeOptions = 32768
	NSXMLNodeNeverEscapeContents NSXMLNodeOptions = 32
	NSXMLNodeOptionsNone NSXMLNodeOptions = 0
	NSXMLNodePreserveAttributeOrder NSXMLNodeOptions = 2097152
	NSXMLNodePreserveCDATA NSXMLNodeOptions = 16777216
	NSXMLNodePreserveCharacterReferences NSXMLNodeOptions = 134217728
	NSXMLNodePreserveDTD NSXMLNodeOptions = 67108864
	NSXMLNodePreserveEntities NSXMLNodeOptions = 4194304
	NSXMLNodePreserveNamespaceOrder NSXMLNodeOptions = 1048576
	NSXMLNodePreservePrefixes NSXMLNodeOptions = 8388608
	NSXMLNodePreserveWhitespace NSXMLNodeOptions = 33554432
	NSXMLNodePrettyPrint NSXMLNodeOptions = 131072
	NSXMLNodePromoteSignificantWhitespace NSXMLNodeOptions = 268435456
	NSXMLNodeUseDoubleQuotes NSXMLNodeOptions = 16
	NSXMLNodeUseSingleQuotes NSXMLNodeOptions = 8
)

func (e NSXMLNodeOptions) String() string {
	switch e {
	case NSXMLDocumentIncludeContentTypeDeclaration:
		return "NSXMLDocumentIncludeContentTypeDeclaration"
	case NSXMLDocumentTidyHTML:
		return "NSXMLDocumentTidyHTML"
	case NSXMLDocumentTidyXML:
		return "NSXMLDocumentTidyXML"
	case NSXMLDocumentValidate:
		return "NSXMLDocumentValidate"
	case NSXMLDocumentXInclude:
		return "NSXMLDocumentXInclude"
	case NSXMLNodeCompactEmptyElement:
		return "NSXMLNodeCompactEmptyElement"
	case NSXMLNodeExpandEmptyElement:
		return "NSXMLNodeExpandEmptyElement"
	case NSXMLNodeIsCDATA:
		return "NSXMLNodeIsCDATA"
	case NSXMLNodeLoadExternalEntitiesAlways:
		return "NSXMLNodeLoadExternalEntitiesAlways"
	case NSXMLNodeLoadExternalEntitiesNever:
		return "NSXMLNodeLoadExternalEntitiesNever"
	case NSXMLNodeLoadExternalEntitiesSameOriginOnly:
		return "NSXMLNodeLoadExternalEntitiesSameOriginOnly"
	case NSXMLNodeNeverEscapeContents:
		return "NSXMLNodeNeverEscapeContents"
	case NSXMLNodeOptionsNone:
		return "NSXMLNodeOptionsNone"
	case NSXMLNodePreserveAttributeOrder:
		return "NSXMLNodePreserveAttributeOrder"
	case NSXMLNodePreserveCDATA:
		return "NSXMLNodePreserveCDATA"
	case NSXMLNodePreserveCharacterReferences:
		return "NSXMLNodePreserveCharacterReferences"
	case NSXMLNodePreserveDTD:
		return "NSXMLNodePreserveDTD"
	case NSXMLNodePreserveEntities:
		return "NSXMLNodePreserveEntities"
	case NSXMLNodePreserveNamespaceOrder:
		return "NSXMLNodePreserveNamespaceOrder"
	case NSXMLNodePreservePrefixes:
		return "NSXMLNodePreservePrefixes"
	case NSXMLNodePreserveWhitespace:
		return "NSXMLNodePreserveWhitespace"
	case NSXMLNodePrettyPrint:
		return "NSXMLNodePrettyPrint"
	case NSXMLNodePromoteSignificantWhitespace:
		return "NSXMLNodePromoteSignificantWhitespace"
	case NSXMLNodeUseDoubleQuotes:
		return "NSXMLNodeUseDoubleQuotes"
	case NSXMLNodeUseSingleQuotes:
		return "NSXMLNodeUseSingleQuotes"
	default:
		return fmt.Sprintf("NSXMLNodeOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/XMLParser/ErrorCode
type NSXMLParserError int

const (
	// NSXMLParserAttributeHasNoValueError: Attribute doesn’t contain a value.
	NSXMLParserAttributeHasNoValueError NSXMLParserError = 41
	// NSXMLParserAttributeListNotFinishedError: Attribute list is not finished.
	NSXMLParserAttributeListNotFinishedError NSXMLParserError = 51
	// NSXMLParserAttributeListNotStartedError: Attribute list is not started.
	NSXMLParserAttributeListNotStartedError NSXMLParserError = 50
	// NSXMLParserAttributeNotFinishedError: Attribute is not finished.
	NSXMLParserAttributeNotFinishedError NSXMLParserError = 40
	// NSXMLParserAttributeNotStartedError: Attribute is not started.
	NSXMLParserAttributeNotStartedError NSXMLParserError = 39
	// NSXMLParserAttributeRedefinedError: Attribute is redefined.
	NSXMLParserAttributeRedefinedError NSXMLParserError = 42
	// NSXMLParserCDATANotFinishedError: CDATA block is not finished.
	NSXMLParserCDATANotFinishedError NSXMLParserError = 63
	// NSXMLParserCharacterRefAtEOFError: Target of character reference cannot be found.
	NSXMLParserCharacterRefAtEOFError NSXMLParserError = 10
	// NSXMLParserCharacterRefInDTDError: Invalid character encountered in the DTD.
	NSXMLParserCharacterRefInDTDError NSXMLParserError = 13
	// NSXMLParserCharacterRefInEpilogError: Invalid character found in the epilog.
	NSXMLParserCharacterRefInEpilogError NSXMLParserError = 12
	// NSXMLParserCharacterRefInPrologError: Invalid character found in the prolog.
	NSXMLParserCharacterRefInPrologError NSXMLParserError = 11
	// NSXMLParserCommentContainsDoubleHyphenError: Comment contains double hyphen.
	NSXMLParserCommentContainsDoubleHyphenError NSXMLParserError = 80
	// NSXMLParserCommentNotFinishedError: Comment is not finished.
	NSXMLParserCommentNotFinishedError NSXMLParserError = 45
	// NSXMLParserConditionalSectionNotFinishedError: Conditional section is not finished.
	NSXMLParserConditionalSectionNotFinishedError NSXMLParserError = 59
	// NSXMLParserConditionalSectionNotStartedError: Conditional section is not started.
	NSXMLParserConditionalSectionNotStartedError NSXMLParserError = 58
	// NSXMLParserDOCTYPEDeclNotFinishedError: Document type declaration is not finished.
	NSXMLParserDOCTYPEDeclNotFinishedError NSXMLParserError = 61
	// NSXMLParserDelegateAbortedParseError: Delegate aborted parse.
	NSXMLParserDelegateAbortedParseError NSXMLParserError = 512
	// NSXMLParserDocumentStartError: The parser object is unable to start parsing.
	NSXMLParserDocumentStartError NSXMLParserError = 3
	// NSXMLParserElementContentDeclNotFinishedError: Element content declaration is not finished.
	NSXMLParserElementContentDeclNotFinishedError NSXMLParserError = 55
	// NSXMLParserElementContentDeclNotStartedError: Element content declaration is not started.
	NSXMLParserElementContentDeclNotStartedError NSXMLParserError = 54
	// NSXMLParserEmptyDocumentError: The document is empty.
	NSXMLParserEmptyDocumentError NSXMLParserError = 4
	// NSXMLParserEncodingNotSupportedError: Document encoding is not supported.
	NSXMLParserEncodingNotSupportedError NSXMLParserError = 32
	// NSXMLParserEntityBoundaryError: Entity boundary error.
	NSXMLParserEntityBoundaryError NSXMLParserError = 90
	// NSXMLParserEntityIsExternalError: Cannot parse external entity.
	NSXMLParserEntityIsExternalError NSXMLParserError = 29
	// NSXMLParserEntityIsParameterError: Entity is a parameter.
	NSXMLParserEntityIsParameterError NSXMLParserError = 30
	// NSXMLParserEntityNotFinishedError: Entity is not finished.
	NSXMLParserEntityNotFinishedError NSXMLParserError = 37
	// NSXMLParserEntityNotStartedError: Entity is not started.
	NSXMLParserEntityNotStartedError NSXMLParserError = 36
	// NSXMLParserEntityRefAtEOFError: Target of entity reference is not found.
	NSXMLParserEntityRefAtEOFError NSXMLParserError = 14
	// NSXMLParserEntityRefInDTDError: Invalid entity reference found in the DTD.
	NSXMLParserEntityRefInDTDError NSXMLParserError = 17
	// NSXMLParserEntityRefInEpilogError: Invalid entity reference found in the epilog.
	NSXMLParserEntityRefInEpilogError NSXMLParserError = 16
	// NSXMLParserEntityRefInPrologError: Invalid entity reference found in the prolog.
	NSXMLParserEntityRefInPrologError NSXMLParserError = 15
	// NSXMLParserEntityRefLoopError: Entity reference loop encountered.
	NSXMLParserEntityRefLoopError NSXMLParserError = 89
	// NSXMLParserEntityReferenceMissingSemiError: Entity reference is missing semicolon.
	NSXMLParserEntityReferenceMissingSemiError NSXMLParserError = 23
	// NSXMLParserEntityReferenceWithoutNameError: Entity reference is without name.
	NSXMLParserEntityReferenceWithoutNameError NSXMLParserError = 22
	// NSXMLParserEntityValueRequiredError: Entity value is required.
	NSXMLParserEntityValueRequiredError NSXMLParserError = 84
	// NSXMLParserEqualExpectedError: Equal sign expected.
	NSXMLParserEqualExpectedError NSXMLParserError = 75
	// NSXMLParserExternalStandaloneEntityError: External standalone entity.
	NSXMLParserExternalStandaloneEntityError NSXMLParserError = 82
	// NSXMLParserExternalSubsetNotFinishedError: External subset is not finished.
	NSXMLParserExternalSubsetNotFinishedError NSXMLParserError = 60
	// NSXMLParserExtraContentError: Error in content found.
	NSXMLParserExtraContentError NSXMLParserError = 86
	// NSXMLParserGTRequiredError: Right angle bracket is required.
	NSXMLParserGTRequiredError NSXMLParserError = 73
	// NSXMLParserInternalError: The parser object encountered an internal error.
	NSXMLParserInternalError NSXMLParserError = 1
	// NSXMLParserInvalidCharacterError: Invalid character encountered.
	NSXMLParserInvalidCharacterError NSXMLParserError = 9
	// NSXMLParserInvalidCharacterInEntityError: Invalid character in entity found.
	NSXMLParserInvalidCharacterInEntityError NSXMLParserError = 87
	// NSXMLParserInvalidCharacterRefError: Invalid character reference encountered.
	NSXMLParserInvalidCharacterRefError NSXMLParserError = 8
	// NSXMLParserInvalidConditionalSectionError: Invalid conditional section.
	NSXMLParserInvalidConditionalSectionError NSXMLParserError = 83
	// NSXMLParserInvalidDecimalCharacterRefError: Invalid decimal character reference encountered.
	NSXMLParserInvalidDecimalCharacterRefError NSXMLParserError = 7
	// NSXMLParserInvalidEncodingError: Invalid encoding.
	NSXMLParserInvalidEncodingError NSXMLParserError = 81
	// NSXMLParserInvalidEncodingNameError: Invalid encoding name found.
	NSXMLParserInvalidEncodingNameError NSXMLParserError = 79
	// NSXMLParserInvalidHexCharacterRefError: Invalid hexadecimal character reference encountered.
	NSXMLParserInvalidHexCharacterRefError NSXMLParserError = 6
	// NSXMLParserInvalidURIError: Invalid URI specified.
	NSXMLParserInvalidURIError NSXMLParserError = 91
	// NSXMLParserLTRequiredError: Left angle bracket is required.
	NSXMLParserLTRequiredError NSXMLParserError = 72
	// NSXMLParserLTSlashRequiredError: Left angle bracket slash is required.
	NSXMLParserLTSlashRequiredError NSXMLParserError = 74
	// NSXMLParserLessThanSymbolInAttributeError: Angle bracket is used in attribute.
	NSXMLParserLessThanSymbolInAttributeError NSXMLParserError = 38
	// NSXMLParserLiteralNotFinishedError: Literal is not finished.
	NSXMLParserLiteralNotFinishedError NSXMLParserError = 44
	// NSXMLParserLiteralNotStartedError: Literal is not started.
	NSXMLParserLiteralNotStartedError NSXMLParserError = 43
	// NSXMLParserMisplacedCDATAEndStringError: Misplaced CDATA end string.
	NSXMLParserMisplacedCDATAEndStringError NSXMLParserError = 62
	// NSXMLParserMisplacedXMLDeclarationError: Misplaced XML declaration.
	NSXMLParserMisplacedXMLDeclarationError NSXMLParserError = 64
	// NSXMLParserMixedContentDeclNotFinishedError: Mixed content declaration is not finished.
	NSXMLParserMixedContentDeclNotFinishedError NSXMLParserError = 53
	// NSXMLParserMixedContentDeclNotStartedError: Mixed content declaration is not started.
	NSXMLParserMixedContentDeclNotStartedError NSXMLParserError = 52
	// NSXMLParserNAMERequiredError: Name is required.
	NSXMLParserNAMERequiredError NSXMLParserError = 68
	// NSXMLParserNMTOKENRequiredError: Name token is required.
	NSXMLParserNMTOKENRequiredError NSXMLParserError = 67
	// NSXMLParserNamespaceDeclarationError: Invalid namespace declaration encountered.
	NSXMLParserNamespaceDeclarationError NSXMLParserError = 35
	// NSXMLParserNoDTDError: Missing DTD.
	NSXMLParserNoDTDError NSXMLParserError = 94
	// NSXMLParserNotWellBalancedError: Document is not well balanced.
	NSXMLParserNotWellBalancedError NSXMLParserError = 85
	// NSXMLParserNotationNotFinishedError: Notation is not finished.
	NSXMLParserNotationNotFinishedError NSXMLParserError = 49
	// NSXMLParserNotationNotStartedError: Notation is not started.
	NSXMLParserNotationNotStartedError NSXMLParserError = 48
	// NSXMLParserOutOfMemoryError: The parser object ran out of memory.
	NSXMLParserOutOfMemoryError NSXMLParserError = 2
	// NSXMLParserPCDATARequiredError: CDATA is required.
	NSXMLParserPCDATARequiredError NSXMLParserError = 69
	// NSXMLParserParsedEntityRefAtEOFError: Target of parsed entity reference is not found.
	NSXMLParserParsedEntityRefAtEOFError NSXMLParserError = 18
	// NSXMLParserParsedEntityRefInEpilogError: Target of parsed entity reference is not found in epilog.
	NSXMLParserParsedEntityRefInEpilogError NSXMLParserError = 20
	// NSXMLParserParsedEntityRefInInternalError: Internal error in parsed entity reference found.
	NSXMLParserParsedEntityRefInInternalError NSXMLParserError = 88
	// NSXMLParserParsedEntityRefInInternalSubsetError: Target of parsed entity reference is not found in internal subset.
	NSXMLParserParsedEntityRefInInternalSubsetError NSXMLParserError = 21
	// NSXMLParserParsedEntityRefInPrologError: Target of parsed entity reference is not found in prolog.
	NSXMLParserParsedEntityRefInPrologError NSXMLParserError = 19
	// NSXMLParserParsedEntityRefMissingSemiError: Parsed entity reference is missing semicolon.
	NSXMLParserParsedEntityRefMissingSemiError NSXMLParserError = 25
	// NSXMLParserParsedEntityRefNoNameError: Parsed entity reference is without an entity name.
	NSXMLParserParsedEntityRefNoNameError NSXMLParserError = 24
	// NSXMLParserPrematureDocumentEndError: The document ended unexpectedly.
	NSXMLParserPrematureDocumentEndError NSXMLParserError = 5
	// NSXMLParserProcessingInstructionNotFinishedError: Processing instruction is not finished.
	NSXMLParserProcessingInstructionNotFinishedError NSXMLParserError = 47
	// NSXMLParserProcessingInstructionNotStartedError: Processing instruction is not started.
	NSXMLParserProcessingInstructionNotStartedError NSXMLParserError = 46
	// NSXMLParserPublicIdentifierRequiredError: Public identifier is required.
	NSXMLParserPublicIdentifierRequiredError NSXMLParserError = 71
	// NSXMLParserSeparatorRequiredError: Separator is required.
	NSXMLParserSeparatorRequiredError NSXMLParserError = 66
	// NSXMLParserSpaceRequiredError: Space is required.
	NSXMLParserSpaceRequiredError NSXMLParserError = 65
	// NSXMLParserStandaloneValueError: Standalone value found.
	NSXMLParserStandaloneValueError NSXMLParserError = 78
	// NSXMLParserStringNotClosedError: String is not closed.
	NSXMLParserStringNotClosedError NSXMLParserError = 34
	// NSXMLParserStringNotStartedError: String is not started.
	NSXMLParserStringNotStartedError NSXMLParserError = 33
	// NSXMLParserTagNameMismatchError: Tag name mismatch.
	NSXMLParserTagNameMismatchError NSXMLParserError = 76
	// NSXMLParserURIFragmentError: URI fragment.
	NSXMLParserURIFragmentError NSXMLParserError = 92
	// NSXMLParserURIRequiredError: URI is required.
	NSXMLParserURIRequiredError NSXMLParserError = 70
	// NSXMLParserUndeclaredEntityError: Entity is not declared.
	NSXMLParserUndeclaredEntityError NSXMLParserError = 26
	// NSXMLParserUnfinishedTagError: Unfinished tag found.
	NSXMLParserUnfinishedTagError NSXMLParserError = 77
	// NSXMLParserUnknownEncodingError: Document encoding is unknown.
	NSXMLParserUnknownEncodingError NSXMLParserError = 31
	// NSXMLParserUnparsedEntityError: Cannot parse entity.
	NSXMLParserUnparsedEntityError NSXMLParserError = 28
	// NSXMLParserXMLDeclNotFinishedError: XML declaration is not finished.
	NSXMLParserXMLDeclNotFinishedError NSXMLParserError = 57
	// NSXMLParserXMLDeclNotStartedError: XML declaration is not started.
	NSXMLParserXMLDeclNotStartedError NSXMLParserError = 56
)

func (e NSXMLParserError) String() string {
	switch e {
	case NSXMLParserAttributeHasNoValueError:
		return "NSXMLParserAttributeHasNoValueError"
	case NSXMLParserAttributeListNotFinishedError:
		return "NSXMLParserAttributeListNotFinishedError"
	case NSXMLParserAttributeListNotStartedError:
		return "NSXMLParserAttributeListNotStartedError"
	case NSXMLParserAttributeNotFinishedError:
		return "NSXMLParserAttributeNotFinishedError"
	case NSXMLParserAttributeNotStartedError:
		return "NSXMLParserAttributeNotStartedError"
	case NSXMLParserAttributeRedefinedError:
		return "NSXMLParserAttributeRedefinedError"
	case NSXMLParserCDATANotFinishedError:
		return "NSXMLParserCDATANotFinishedError"
	case NSXMLParserCharacterRefAtEOFError:
		return "NSXMLParserCharacterRefAtEOFError"
	case NSXMLParserCharacterRefInDTDError:
		return "NSXMLParserCharacterRefInDTDError"
	case NSXMLParserCharacterRefInEpilogError:
		return "NSXMLParserCharacterRefInEpilogError"
	case NSXMLParserCharacterRefInPrologError:
		return "NSXMLParserCharacterRefInPrologError"
	case NSXMLParserCommentContainsDoubleHyphenError:
		return "NSXMLParserCommentContainsDoubleHyphenError"
	case NSXMLParserCommentNotFinishedError:
		return "NSXMLParserCommentNotFinishedError"
	case NSXMLParserConditionalSectionNotFinishedError:
		return "NSXMLParserConditionalSectionNotFinishedError"
	case NSXMLParserConditionalSectionNotStartedError:
		return "NSXMLParserConditionalSectionNotStartedError"
	case NSXMLParserDOCTYPEDeclNotFinishedError:
		return "NSXMLParserDOCTYPEDeclNotFinishedError"
	case NSXMLParserDelegateAbortedParseError:
		return "NSXMLParserDelegateAbortedParseError"
	case NSXMLParserDocumentStartError:
		return "NSXMLParserDocumentStartError"
	case NSXMLParserElementContentDeclNotFinishedError:
		return "NSXMLParserElementContentDeclNotFinishedError"
	case NSXMLParserElementContentDeclNotStartedError:
		return "NSXMLParserElementContentDeclNotStartedError"
	case NSXMLParserEmptyDocumentError:
		return "NSXMLParserEmptyDocumentError"
	case NSXMLParserEncodingNotSupportedError:
		return "NSXMLParserEncodingNotSupportedError"
	case NSXMLParserEntityBoundaryError:
		return "NSXMLParserEntityBoundaryError"
	case NSXMLParserEntityIsExternalError:
		return "NSXMLParserEntityIsExternalError"
	case NSXMLParserEntityIsParameterError:
		return "NSXMLParserEntityIsParameterError"
	case NSXMLParserEntityNotFinishedError:
		return "NSXMLParserEntityNotFinishedError"
	case NSXMLParserEntityNotStartedError:
		return "NSXMLParserEntityNotStartedError"
	case NSXMLParserEntityRefAtEOFError:
		return "NSXMLParserEntityRefAtEOFError"
	case NSXMLParserEntityRefInDTDError:
		return "NSXMLParserEntityRefInDTDError"
	case NSXMLParserEntityRefInEpilogError:
		return "NSXMLParserEntityRefInEpilogError"
	case NSXMLParserEntityRefInPrologError:
		return "NSXMLParserEntityRefInPrologError"
	case NSXMLParserEntityRefLoopError:
		return "NSXMLParserEntityRefLoopError"
	case NSXMLParserEntityReferenceMissingSemiError:
		return "NSXMLParserEntityReferenceMissingSemiError"
	case NSXMLParserEntityReferenceWithoutNameError:
		return "NSXMLParserEntityReferenceWithoutNameError"
	case NSXMLParserEntityValueRequiredError:
		return "NSXMLParserEntityValueRequiredError"
	case NSXMLParserEqualExpectedError:
		return "NSXMLParserEqualExpectedError"
	case NSXMLParserExternalStandaloneEntityError:
		return "NSXMLParserExternalStandaloneEntityError"
	case NSXMLParserExternalSubsetNotFinishedError:
		return "NSXMLParserExternalSubsetNotFinishedError"
	case NSXMLParserExtraContentError:
		return "NSXMLParserExtraContentError"
	case NSXMLParserGTRequiredError:
		return "NSXMLParserGTRequiredError"
	case NSXMLParserInternalError:
		return "NSXMLParserInternalError"
	case NSXMLParserInvalidCharacterError:
		return "NSXMLParserInvalidCharacterError"
	case NSXMLParserInvalidCharacterInEntityError:
		return "NSXMLParserInvalidCharacterInEntityError"
	case NSXMLParserInvalidCharacterRefError:
		return "NSXMLParserInvalidCharacterRefError"
	case NSXMLParserInvalidConditionalSectionError:
		return "NSXMLParserInvalidConditionalSectionError"
	case NSXMLParserInvalidDecimalCharacterRefError:
		return "NSXMLParserInvalidDecimalCharacterRefError"
	case NSXMLParserInvalidEncodingError:
		return "NSXMLParserInvalidEncodingError"
	case NSXMLParserInvalidEncodingNameError:
		return "NSXMLParserInvalidEncodingNameError"
	case NSXMLParserInvalidHexCharacterRefError:
		return "NSXMLParserInvalidHexCharacterRefError"
	case NSXMLParserInvalidURIError:
		return "NSXMLParserInvalidURIError"
	case NSXMLParserLTRequiredError:
		return "NSXMLParserLTRequiredError"
	case NSXMLParserLTSlashRequiredError:
		return "NSXMLParserLTSlashRequiredError"
	case NSXMLParserLessThanSymbolInAttributeError:
		return "NSXMLParserLessThanSymbolInAttributeError"
	case NSXMLParserLiteralNotFinishedError:
		return "NSXMLParserLiteralNotFinishedError"
	case NSXMLParserLiteralNotStartedError:
		return "NSXMLParserLiteralNotStartedError"
	case NSXMLParserMisplacedCDATAEndStringError:
		return "NSXMLParserMisplacedCDATAEndStringError"
	case NSXMLParserMisplacedXMLDeclarationError:
		return "NSXMLParserMisplacedXMLDeclarationError"
	case NSXMLParserMixedContentDeclNotFinishedError:
		return "NSXMLParserMixedContentDeclNotFinishedError"
	case NSXMLParserMixedContentDeclNotStartedError:
		return "NSXMLParserMixedContentDeclNotStartedError"
	case NSXMLParserNAMERequiredError:
		return "NSXMLParserNAMERequiredError"
	case NSXMLParserNMTOKENRequiredError:
		return "NSXMLParserNMTOKENRequiredError"
	case NSXMLParserNamespaceDeclarationError:
		return "NSXMLParserNamespaceDeclarationError"
	case NSXMLParserNoDTDError:
		return "NSXMLParserNoDTDError"
	case NSXMLParserNotWellBalancedError:
		return "NSXMLParserNotWellBalancedError"
	case NSXMLParserNotationNotFinishedError:
		return "NSXMLParserNotationNotFinishedError"
	case NSXMLParserNotationNotStartedError:
		return "NSXMLParserNotationNotStartedError"
	case NSXMLParserOutOfMemoryError:
		return "NSXMLParserOutOfMemoryError"
	case NSXMLParserPCDATARequiredError:
		return "NSXMLParserPCDATARequiredError"
	case NSXMLParserParsedEntityRefAtEOFError:
		return "NSXMLParserParsedEntityRefAtEOFError"
	case NSXMLParserParsedEntityRefInEpilogError:
		return "NSXMLParserParsedEntityRefInEpilogError"
	case NSXMLParserParsedEntityRefInInternalError:
		return "NSXMLParserParsedEntityRefInInternalError"
	case NSXMLParserParsedEntityRefInInternalSubsetError:
		return "NSXMLParserParsedEntityRefInInternalSubsetError"
	case NSXMLParserParsedEntityRefInPrologError:
		return "NSXMLParserParsedEntityRefInPrologError"
	case NSXMLParserParsedEntityRefMissingSemiError:
		return "NSXMLParserParsedEntityRefMissingSemiError"
	case NSXMLParserParsedEntityRefNoNameError:
		return "NSXMLParserParsedEntityRefNoNameError"
	case NSXMLParserPrematureDocumentEndError:
		return "NSXMLParserPrematureDocumentEndError"
	case NSXMLParserProcessingInstructionNotFinishedError:
		return "NSXMLParserProcessingInstructionNotFinishedError"
	case NSXMLParserProcessingInstructionNotStartedError:
		return "NSXMLParserProcessingInstructionNotStartedError"
	case NSXMLParserPublicIdentifierRequiredError:
		return "NSXMLParserPublicIdentifierRequiredError"
	case NSXMLParserSeparatorRequiredError:
		return "NSXMLParserSeparatorRequiredError"
	case NSXMLParserSpaceRequiredError:
		return "NSXMLParserSpaceRequiredError"
	case NSXMLParserStandaloneValueError:
		return "NSXMLParserStandaloneValueError"
	case NSXMLParserStringNotClosedError:
		return "NSXMLParserStringNotClosedError"
	case NSXMLParserStringNotStartedError:
		return "NSXMLParserStringNotStartedError"
	case NSXMLParserTagNameMismatchError:
		return "NSXMLParserTagNameMismatchError"
	case NSXMLParserURIFragmentError:
		return "NSXMLParserURIFragmentError"
	case NSXMLParserURIRequiredError:
		return "NSXMLParserURIRequiredError"
	case NSXMLParserUndeclaredEntityError:
		return "NSXMLParserUndeclaredEntityError"
	case NSXMLParserUnfinishedTagError:
		return "NSXMLParserUnfinishedTagError"
	case NSXMLParserUnknownEncodingError:
		return "NSXMLParserUnknownEncodingError"
	case NSXMLParserUnparsedEntityError:
		return "NSXMLParserUnparsedEntityError"
	case NSXMLParserXMLDeclNotFinishedError:
		return "NSXMLParserXMLDeclNotFinishedError"
	case NSXMLParserXMLDeclNotStartedError:
		return "NSXMLParserXMLDeclNotStartedError"
	default:
		return fmt.Sprintf("NSXMLParserError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/XMLParser/ExternalEntityResolvingPolicy-swift.enum
type NSXMLParserExternalEntityResolvingPolicy int

const (
	NSXMLParserResolveExternalEntitiesAlways NSXMLParserExternalEntityResolvingPolicy = 3
	NSXMLParserResolveExternalEntitiesNever NSXMLParserExternalEntityResolvingPolicy = 0
	NSXMLParserResolveExternalEntitiesNoNetwork NSXMLParserExternalEntityResolvingPolicy = 1
	NSXMLParserResolveExternalEntitiesSameOriginOnly NSXMLParserExternalEntityResolvingPolicy = 2
)

func (e NSXMLParserExternalEntityResolvingPolicy) String() string {
	switch e {
	case NSXMLParserResolveExternalEntitiesAlways:
		return "NSXMLParserResolveExternalEntitiesAlways"
	case NSXMLParserResolveExternalEntitiesNever:
		return "NSXMLParserResolveExternalEntitiesNever"
	case NSXMLParserResolveExternalEntitiesNoNetwork:
		return "NSXMLParserResolveExternalEntitiesNoNetwork"
	case NSXMLParserResolveExternalEntitiesSameOriginOnly:
		return "NSXMLParserResolveExternalEntitiesSameOriginOnly"
	default:
		return fmt.Sprintf("NSXMLParserExternalEntityResolvingPolicy(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/Options
type NSXPCConnectionOptions int

const (
	// NSXPCConnectionPrivileged: # Discussion
	NSXPCConnectionPrivileged NSXPCConnectionOptions = 4096
)

func (e NSXPCConnectionOptions) String() string {
	switch e {
	case NSXPCConnectionPrivileged:
		return "NSXPCConnectionPrivileged"
	default:
		return fmt.Sprintf("NSXPCConnectionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Foundation/NSItemProvider/PreferredPresentationStyle-swift.enum
type UIPreferredPresentationStyle int

const (
)

