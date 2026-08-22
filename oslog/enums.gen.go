// Code generated from Apple documentation for OSLog. DO NOT EDIT.

package oslog

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/OSLog/OSLogEntryLog/Level-swift.enum
type OSLogEntryLogLevel int

const (
	// OSLogEntryLogLevelDebug: A log level that captures diagnostic information.
	OSLogEntryLogLevelDebug OSLogEntryLogLevel = 1
	// OSLogEntryLogLevelError: The log level that captures errors.
	OSLogEntryLogLevelError OSLogEntryLogLevel = 4
	// OSLogEntryLogLevelFault: The log level that captures fault information.
	OSLogEntryLogLevelFault OSLogEntryLogLevel = 5
	// OSLogEntryLogLevelInfo: The log level that captures additional information.
	OSLogEntryLogLevelInfo OSLogEntryLogLevel = 2
	// OSLogEntryLogLevelNotice: The log level that captures notifications.
	OSLogEntryLogLevelNotice OSLogEntryLogLevel = 3
	// OSLogEntryLogLevelUndefined: The log level was never specified.
	OSLogEntryLogLevelUndefined OSLogEntryLogLevel = 0
)

func (e OSLogEntryLogLevel) String() string {
	switch e {
	case OSLogEntryLogLevelDebug:
		return "OSLogEntryLogLevelDebug"
	case OSLogEntryLogLevelError:
		return "OSLogEntryLogLevelError"
	case OSLogEntryLogLevelFault:
		return "OSLogEntryLogLevelFault"
	case OSLogEntryLogLevelInfo:
		return "OSLogEntryLogLevelInfo"
	case OSLogEntryLogLevelNotice:
		return "OSLogEntryLogLevelNotice"
	case OSLogEntryLogLevelUndefined:
		return "OSLogEntryLogLevelUndefined"
	default:
		return fmt.Sprintf("OSLogEntryLogLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/OSLog/OSLogEntrySignpost/SignpostType-swift.enum
type OSLogEntrySignpostType int

const (
	// OSLogEntrySignpostTypeEvent: The signpost marks an event.
	OSLogEntrySignpostTypeEvent OSLogEntrySignpostType = 3
	// OSLogEntrySignpostTypeIntervalBegin: The signpost marks the start of a time interval.
	OSLogEntrySignpostTypeIntervalBegin OSLogEntrySignpostType = 1
	// OSLogEntrySignpostTypeIntervalEnd: The signpost marks the end of a time interval.
	OSLogEntrySignpostTypeIntervalEnd OSLogEntrySignpostType = 2
	// OSLogEntrySignpostTypeUndefined: The signpost does not have a type.
	OSLogEntrySignpostTypeUndefined OSLogEntrySignpostType = 0
)

func (e OSLogEntrySignpostType) String() string {
	switch e {
	case OSLogEntrySignpostTypeEvent:
		return "OSLogEntrySignpostTypeEvent"
	case OSLogEntrySignpostTypeIntervalBegin:
		return "OSLogEntrySignpostTypeIntervalBegin"
	case OSLogEntrySignpostTypeIntervalEnd:
		return "OSLogEntrySignpostTypeIntervalEnd"
	case OSLogEntrySignpostTypeUndefined:
		return "OSLogEntrySignpostTypeUndefined"
	default:
		return fmt.Sprintf("OSLogEntrySignpostType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/OSLog/OSLogEntry/StoreCategory-swift.enum
type OSLogEntryStoreCategory int

const (
	// OSLogEntryStoreCategoryLongTerm1: The entry was tagged with a hint indicating the system should try to preserve it for approximately 1 day.
	OSLogEntryStoreCategoryLongTerm1 OSLogEntryStoreCategory = 4
	// OSLogEntryStoreCategoryLongTerm14: The entry was tagged with a hint indicating the system should try to preserve it for approximately 14 days.
	OSLogEntryStoreCategoryLongTerm14 OSLogEntryStoreCategory = 7
	// OSLogEntryStoreCategoryLongTerm3: The entry was tagged with a hint indicating the system should try to preserve it for approximately 3 days.
	OSLogEntryStoreCategoryLongTerm3 OSLogEntryStoreCategory = 5
	// OSLogEntryStoreCategoryLongTerm30: The entry was tagged with a hint indicating the system should try to preserve it for approximately 30 days.
	OSLogEntryStoreCategoryLongTerm30 OSLogEntryStoreCategory = 8
	// OSLogEntryStoreCategoryLongTerm7: The entry was tagged with a hint indicating the system should try to preserve it for approximately 7 days.
	OSLogEntryStoreCategoryLongTerm7 OSLogEntryStoreCategory = 6
	// OSLogEntryStoreCategoryLongTermAuto: The entry was tagged with a hint indicating the system should try to preserve it based on the amount of space available.
	OSLogEntryStoreCategoryLongTermAuto OSLogEntryStoreCategory = 3
	// OSLogEntryStoreCategoryMetadata: This entry was generated as information about the other entries or about the sequence of entries as a whole.
	OSLogEntryStoreCategoryMetadata OSLogEntryStoreCategory = 1
	// OSLogEntryStoreCategoryShortTerm: This entry was not intended to be long-lived, and was captured in the ring buffer.
	OSLogEntryStoreCategoryShortTerm OSLogEntryStoreCategory = 2
	// OSLogEntryStoreCategoryUndefined: This entry’s purpose is unknown.
	OSLogEntryStoreCategoryUndefined OSLogEntryStoreCategory = 0
)

func (e OSLogEntryStoreCategory) String() string {
	switch e {
	case OSLogEntryStoreCategoryLongTerm1:
		return "OSLogEntryStoreCategoryLongTerm1"
	case OSLogEntryStoreCategoryLongTerm14:
		return "OSLogEntryStoreCategoryLongTerm14"
	case OSLogEntryStoreCategoryLongTerm3:
		return "OSLogEntryStoreCategoryLongTerm3"
	case OSLogEntryStoreCategoryLongTerm30:
		return "OSLogEntryStoreCategoryLongTerm30"
	case OSLogEntryStoreCategoryLongTerm7:
		return "OSLogEntryStoreCategoryLongTerm7"
	case OSLogEntryStoreCategoryLongTermAuto:
		return "OSLogEntryStoreCategoryLongTermAuto"
	case OSLogEntryStoreCategoryMetadata:
		return "OSLogEntryStoreCategoryMetadata"
	case OSLogEntryStoreCategoryShortTerm:
		return "OSLogEntryStoreCategoryShortTerm"
	case OSLogEntryStoreCategoryUndefined:
		return "OSLogEntryStoreCategoryUndefined"
	default:
		return fmt.Sprintf("OSLogEntryStoreCategory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/OSLog/OSLogEnumerator/Options
type OSLogEnumeratorOptions uint

const (
	// OSLogEnumeratorReverse: Tells the framework to iterate backwards.
	OSLogEnumeratorReverse OSLogEnumeratorOptions = 0x1
)

func (e OSLogEnumeratorOptions) String() string {
	switch e {
	case OSLogEnumeratorReverse:
		return "OSLogEnumeratorReverse"
	default:
		return fmt.Sprintf("OSLogEnumeratorOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/OSLog/OSLogMessageComponent/ArgumentCategory-swift.enum
type OSLogMessageComponentArgumentCategory int

const (
	// OSLogMessageComponentArgumentCategoryData: The argument is an NSData object.
	OSLogMessageComponentArgumentCategoryData OSLogMessageComponentArgumentCategory = 1
	// OSLogMessageComponentArgumentCategoryDouble: The argument is a double.
	OSLogMessageComponentArgumentCategoryDouble OSLogMessageComponentArgumentCategory = 2
	// OSLogMessageComponentArgumentCategoryInt64: The argument is a 64-bit signed integer.
	OSLogMessageComponentArgumentCategoryInt64 OSLogMessageComponentArgumentCategory = 3
	// OSLogMessageComponentArgumentCategoryString: The argument is a string.
	OSLogMessageComponentArgumentCategoryString OSLogMessageComponentArgumentCategory = 4
	// OSLogMessageComponentArgumentCategoryUInt64: The argument is a 64-bit unsigned integer.
	OSLogMessageComponentArgumentCategoryUInt64 OSLogMessageComponentArgumentCategory = 5
	// OSLogMessageComponentArgumentCategoryUndefined: The argument’s type is not defined.
	OSLogMessageComponentArgumentCategoryUndefined OSLogMessageComponentArgumentCategory = 0
)

func (e OSLogMessageComponentArgumentCategory) String() string {
	switch e {
	case OSLogMessageComponentArgumentCategoryData:
		return "OSLogMessageComponentArgumentCategoryData"
	case OSLogMessageComponentArgumentCategoryDouble:
		return "OSLogMessageComponentArgumentCategoryDouble"
	case OSLogMessageComponentArgumentCategoryInt64:
		return "OSLogMessageComponentArgumentCategoryInt64"
	case OSLogMessageComponentArgumentCategoryString:
		return "OSLogMessageComponentArgumentCategoryString"
	case OSLogMessageComponentArgumentCategoryUInt64:
		return "OSLogMessageComponentArgumentCategoryUInt64"
	case OSLogMessageComponentArgumentCategoryUndefined:
		return "OSLogMessageComponentArgumentCategoryUndefined"
	default:
		return fmt.Sprintf("OSLogMessageComponentArgumentCategory(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/OSLog/OSLogStore/Scope
type OSLogStoreScope int

const (
	OSLogStoreCurrentProcessIdentifier OSLogStoreScope = 1
	OSLogStoreSystem                   OSLogStoreScope = 0
)

func (e OSLogStoreScope) String() string {
	switch e {
	case OSLogStoreCurrentProcessIdentifier:
		return "OSLogStoreCurrentProcessIdentifier"
	case OSLogStoreSystem:
		return "OSLogStoreSystem"
	default:
		return fmt.Sprintf("OSLogStoreScope(%d)", e)
	}
}
