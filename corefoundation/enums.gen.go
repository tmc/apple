// Code generated from Apple documentation for CoreFoundation. DO NOT EDIT.

package corefoundation

import (
	"fmt"
)

type CFByteOrder uint

const (
	// CFByteOrderBigEndian: Multi-byte values are stored with the most-significant bytes stored first.
	CFByteOrderBigEndian CFByteOrder = 2
	// CFByteOrderLittleEndian: Multi-byte values are stored with the least-significant bytes stored first.
	CFByteOrderLittleEndian CFByteOrder = 1
	// CFByteOrderUnknown: The byte order is unknown.
	CFByteOrderUnknown CFByteOrder = 0
)

func (e CFByteOrder) String() string {
	switch e {
	case CFByteOrderBigEndian:
		return "CFByteOrderBigEndian"
	case CFByteOrderLittleEndian:
		return "CFByteOrderLittleEndian"
	case CFByteOrderUnknown:
		return "CFByteOrderUnknown"
	default:
		return fmt.Sprintf("CFByteOrder(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFCalendarUnit
type CFCalendarUnit int

const (
	// KCFCalendarUnitDay: Specifies the day unit.
	KCFCalendarUnitDay CFCalendarUnit = 16
	KCFCalendarUnitDayOfYear CFCalendarUnit = 65536
	// KCFCalendarUnitEra: Specifies the era unit.
	KCFCalendarUnitEra CFCalendarUnit = 2
	// KCFCalendarUnitHour: Specifies the hour unit.
	KCFCalendarUnitHour CFCalendarUnit = 32
	// KCFCalendarUnitMinute: Specifies the minute unit.
	KCFCalendarUnitMinute CFCalendarUnit = 64
	// KCFCalendarUnitMonth: Specifies the month unit.
	KCFCalendarUnitMonth CFCalendarUnit = 8
	// KCFCalendarUnitQuarter: Specifies the quarter-year unit.
	KCFCalendarUnitQuarter CFCalendarUnit = 2048
	// KCFCalendarUnitSecond: Specifies the second unit.
	KCFCalendarUnitSecond CFCalendarUnit = 128
	// KCFCalendarUnitWeekOfMonth: Specifies the original week of a month calendar unit.
	KCFCalendarUnitWeekOfMonth CFCalendarUnit = 4096
	// KCFCalendarUnitWeekOfYear: Specifies the original week of the year calendar unit.
	KCFCalendarUnitWeekOfYear CFCalendarUnit = 8192
	// KCFCalendarUnitWeekday: Specifies the weekday unit.
	KCFCalendarUnitWeekday CFCalendarUnit = 512
	// KCFCalendarUnitWeekdayOrdinal: Specifies the ordinal weekday unit.
	KCFCalendarUnitWeekdayOrdinal CFCalendarUnit = 1024
	// KCFCalendarUnitYear: Specifies the year unit.
	KCFCalendarUnitYear CFCalendarUnit = 4
	// KCFCalendarUnitYearForWeekOfYear: Specifies the relative year for a week within a year calendar unit.
	KCFCalendarUnitYearForWeekOfYear CFCalendarUnit = 16384
	// Deprecated.
	KCFCalendarUnitWeek CFCalendarUnit = 256
)

func (e CFCalendarUnit) String() string {
	switch e {
	case KCFCalendarUnitDay:
		return "KCFCalendarUnitDay"
	case KCFCalendarUnitDayOfYear:
		return "KCFCalendarUnitDayOfYear"
	case KCFCalendarUnitEra:
		return "KCFCalendarUnitEra"
	case KCFCalendarUnitHour:
		return "KCFCalendarUnitHour"
	case KCFCalendarUnitMinute:
		return "KCFCalendarUnitMinute"
	case KCFCalendarUnitMonth:
		return "KCFCalendarUnitMonth"
	case KCFCalendarUnitQuarter:
		return "KCFCalendarUnitQuarter"
	case KCFCalendarUnitSecond:
		return "KCFCalendarUnitSecond"
	case KCFCalendarUnitWeekOfMonth:
		return "KCFCalendarUnitWeekOfMonth"
	case KCFCalendarUnitWeekOfYear:
		return "KCFCalendarUnitWeekOfYear"
	case KCFCalendarUnitWeekday:
		return "KCFCalendarUnitWeekday"
	case KCFCalendarUnitWeekdayOrdinal:
		return "KCFCalendarUnitWeekdayOrdinal"
	case KCFCalendarUnitYear:
		return "KCFCalendarUnitYear"
	case KCFCalendarUnitYearForWeekOfYear:
		return "KCFCalendarUnitYearForWeekOfYear"
	case KCFCalendarUnitWeek:
		return "KCFCalendarUnitWeek"
	default:
		return fmt.Sprintf("CFCalendarUnit(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFCharacterSetPredefinedSet
type CFCharacterSetPredefinedSet int

const (
	// KCFCharacterSetAlphaNumeric: Alpha Numeric character set (Unicode General Category L*, M*, & N*).
	KCFCharacterSetAlphaNumeric CFCharacterSetPredefinedSet = 10
	// KCFCharacterSetCapitalizedLetter: Titlecase character set (Unicode General Category Lt).
	KCFCharacterSetCapitalizedLetter CFCharacterSetPredefinedSet = 13
	// KCFCharacterSetControl: Control character set (Unicode General Category Cc and Cf).
	KCFCharacterSetControl CFCharacterSetPredefinedSet = 1
	// KCFCharacterSetDecimalDigit: Decimal digit character set.
	KCFCharacterSetDecimalDigit CFCharacterSetPredefinedSet = 4
	// KCFCharacterSetDecomposable: Canonically decomposable character set.
	KCFCharacterSetDecomposable CFCharacterSetPredefinedSet = 9
	// KCFCharacterSetIllegal: Illegal character set.
	KCFCharacterSetIllegal CFCharacterSetPredefinedSet = 12
	// KCFCharacterSetLetter: Letter character set (Unicode General Category L* & M*).
	KCFCharacterSetLetter CFCharacterSetPredefinedSet = 5
	// KCFCharacterSetLowercaseLetter: Lowercase character set (Unicode General Category Ll).
	KCFCharacterSetLowercaseLetter CFCharacterSetPredefinedSet = 6
	// KCFCharacterSetNewline: Newline character set (, , , and ).
	KCFCharacterSetNewline CFCharacterSetPredefinedSet = 15
	// KCFCharacterSetNonBase: Non-base character set (Unicode General Category M*).
	KCFCharacterSetNonBase CFCharacterSetPredefinedSet = 8
	// KCFCharacterSetPunctuation: Punctuation character set (Unicode General Category P*).
	KCFCharacterSetPunctuation CFCharacterSetPredefinedSet = 11
	// KCFCharacterSetSymbol: Symbol character set (Unicode General Category S*).
	KCFCharacterSetSymbol CFCharacterSetPredefinedSet = 14
	// KCFCharacterSetUppercaseLetter: Uppercase character set (Unicode General Category Lu and Lt).
	KCFCharacterSetUppercaseLetter CFCharacterSetPredefinedSet = 7
	// KCFCharacterSetWhitespace: Whitespace character set (Unicode General Category Zs and U0009 CHARACTER TABULATION).
	KCFCharacterSetWhitespace CFCharacterSetPredefinedSet = 2
	// KCFCharacterSetWhitespaceAndNewline: Whitespace and Newline character set (Unicode General Category Z*, , and ).
	KCFCharacterSetWhitespaceAndNewline CFCharacterSetPredefinedSet = 3
)

func (e CFCharacterSetPredefinedSet) String() string {
	switch e {
	case KCFCharacterSetAlphaNumeric:
		return "KCFCharacterSetAlphaNumeric"
	case KCFCharacterSetCapitalizedLetter:
		return "KCFCharacterSetCapitalizedLetter"
	case KCFCharacterSetControl:
		return "KCFCharacterSetControl"
	case KCFCharacterSetDecimalDigit:
		return "KCFCharacterSetDecimalDigit"
	case KCFCharacterSetDecomposable:
		return "KCFCharacterSetDecomposable"
	case KCFCharacterSetIllegal:
		return "KCFCharacterSetIllegal"
	case KCFCharacterSetLetter:
		return "KCFCharacterSetLetter"
	case KCFCharacterSetLowercaseLetter:
		return "KCFCharacterSetLowercaseLetter"
	case KCFCharacterSetNewline:
		return "KCFCharacterSetNewline"
	case KCFCharacterSetNonBase:
		return "KCFCharacterSetNonBase"
	case KCFCharacterSetPunctuation:
		return "KCFCharacterSetPunctuation"
	case KCFCharacterSetSymbol:
		return "KCFCharacterSetSymbol"
	case KCFCharacterSetUppercaseLetter:
		return "KCFCharacterSetUppercaseLetter"
	case KCFCharacterSetWhitespace:
		return "KCFCharacterSetWhitespace"
	case KCFCharacterSetWhitespaceAndNewline:
		return "KCFCharacterSetWhitespaceAndNewline"
	default:
		return fmt.Sprintf("CFCharacterSetPredefinedSet(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFComparisonResult
type CFComparisonResult int

const (
	// KCFCompareEqualTo: Returned by a comparison function if the first value is equal to the second value.
	KCFCompareEqualTo CFComparisonResult = 0
	// KCFCompareGreaterThan: Returned by a comparison function if the first value is greater than the second value.
	KCFCompareGreaterThan CFComparisonResult = 1
	// KCFCompareLessThan: Returned by a comparison function if the first value is less than the second value.
	KCFCompareLessThan CFComparisonResult = -1
)

func (e CFComparisonResult) String() string {
	switch e {
	case KCFCompareEqualTo:
		return "KCFCompareEqualTo"
	case KCFCompareGreaterThan:
		return "KCFCompareGreaterThan"
	case KCFCompareLessThan:
		return "KCFCompareLessThan"
	default:
		return fmt.Sprintf("CFComparisonResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFDataSearchFlags
type CFDataSearchFlags int

const (
	// KCFDataSearchAnchored: # Discussion
	KCFDataSearchAnchored CFDataSearchFlags = 2
	// KCFDataSearchBackwards: # Discussion
	KCFDataSearchBackwards CFDataSearchFlags = 1
)

func (e CFDataSearchFlags) String() string {
	switch e {
	case KCFDataSearchAnchored:
		return "KCFDataSearchAnchored"
	case KCFDataSearchBackwards:
		return "KCFDataSearchBackwards"
	default:
		return fmt.Sprintf("CFDataSearchFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFDateFormatterStyle
type CFDateFormatterStyle int

const (
	// KCFDateFormatterFullStyle: Specifies a full style with complete details, such as “Tuesday, April 12, 1952 AD” or “3:30:42pm PST”.
	KCFDateFormatterFullStyle CFDateFormatterStyle = 4
	// KCFDateFormatterLongStyle: Specifies a long style, typically with full text, such as “November 23, 1937” or “3:30:32pm”.
	KCFDateFormatterLongStyle CFDateFormatterStyle = 3
	// KCFDateFormatterMediumStyle: Specifies a medium style, typically with abbreviated text, such as “Nov 23, 1937”.
	KCFDateFormatterMediumStyle CFDateFormatterStyle = 2
	// KCFDateFormatterNoStyle: Specifies no output.
	KCFDateFormatterNoStyle CFDateFormatterStyle = 0
	// KCFDateFormatterShortStyle: Specifies a short style, typically numeric only, such as “11/23/37” or “3:30pm”.
	KCFDateFormatterShortStyle CFDateFormatterStyle = 1
)

func (e CFDateFormatterStyle) String() string {
	switch e {
	case KCFDateFormatterFullStyle:
		return "KCFDateFormatterFullStyle"
	case KCFDateFormatterLongStyle:
		return "KCFDateFormatterLongStyle"
	case KCFDateFormatterMediumStyle:
		return "KCFDateFormatterMediumStyle"
	case KCFDateFormatterNoStyle:
		return "KCFDateFormatterNoStyle"
	case KCFDateFormatterShortStyle:
		return "KCFDateFormatterShortStyle"
	default:
		return fmt.Sprintf("CFDateFormatterStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFFileSecurityClearOptions
type CFFileSecurityClearOptions int

const (
	// KCFFileSecurityClearAccessControlList: Clear the access control list.
	KCFFileSecurityClearAccessControlList CFFileSecurityClearOptions = 32
	// KCFFileSecurityClearGroup: Clear the (POSIX) group ID.
	KCFFileSecurityClearGroup CFFileSecurityClearOptions = 2
	// KCFFileSecurityClearGroupUUID: Clear the group UUID (for the access control list).
	KCFFileSecurityClearGroupUUID CFFileSecurityClearOptions = 16
	// KCFFileSecurityClearMode: Clear the file’s mode (POSIX permissions).
	KCFFileSecurityClearMode CFFileSecurityClearOptions = 4
	// KCFFileSecurityClearOwner: Clear the (POSIX) owner ID.
	KCFFileSecurityClearOwner CFFileSecurityClearOptions = 1
	// KCFFileSecurityClearOwnerUUID: Clear the owner UUID (for the access control list).
	KCFFileSecurityClearOwnerUUID CFFileSecurityClearOptions = 8
)

func (e CFFileSecurityClearOptions) String() string {
	switch e {
	case KCFFileSecurityClearAccessControlList:
		return "KCFFileSecurityClearAccessControlList"
	case KCFFileSecurityClearGroup:
		return "KCFFileSecurityClearGroup"
	case KCFFileSecurityClearGroupUUID:
		return "KCFFileSecurityClearGroupUUID"
	case KCFFileSecurityClearMode:
		return "KCFFileSecurityClearMode"
	case KCFFileSecurityClearOwner:
		return "KCFFileSecurityClearOwner"
	case KCFFileSecurityClearOwnerUUID:
		return "KCFFileSecurityClearOwnerUUID"
	default:
		return fmt.Sprintf("CFFileSecurityClearOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFGregorianUnitFlags
type CFGregorianUnitFlags int

const (
	// Deprecated.
	KCFGregorianAllUnits CFGregorianUnitFlags = 16777215
	// Deprecated.
	KCFGregorianUnitsDays CFGregorianUnitFlags = 4
	// Deprecated.
	KCFGregorianUnitsHours CFGregorianUnitFlags = 8
	// Deprecated.
	KCFGregorianUnitsMinutes CFGregorianUnitFlags = 16
	// Deprecated.
	KCFGregorianUnitsMonths CFGregorianUnitFlags = 2
	// Deprecated.
	KCFGregorianUnitsSeconds CFGregorianUnitFlags = 32
	// Deprecated.
	KCFGregorianUnitsYears CFGregorianUnitFlags = 1
)

func (e CFGregorianUnitFlags) String() string {
	switch e {
	case KCFGregorianAllUnits:
		return "KCFGregorianAllUnits"
	case KCFGregorianUnitsDays:
		return "KCFGregorianUnitsDays"
	case KCFGregorianUnitsHours:
		return "KCFGregorianUnitsHours"
	case KCFGregorianUnitsMinutes:
		return "KCFGregorianUnitsMinutes"
	case KCFGregorianUnitsMonths:
		return "KCFGregorianUnitsMonths"
	case KCFGregorianUnitsSeconds:
		return "KCFGregorianUnitsSeconds"
	case KCFGregorianUnitsYears:
		return "KCFGregorianUnitsYears"
	default:
		return fmt.Sprintf("CFGregorianUnitFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFISO8601DateFormatOptions
type CFISO8601DateFormatOptions int

const (
	KCFISO8601DateFormatWithColonSeparatorInTime CFISO8601DateFormatOptions = 512
	KCFISO8601DateFormatWithColonSeparatorInTimeZone CFISO8601DateFormatOptions = 1024
	KCFISO8601DateFormatWithDashSeparatorInDate CFISO8601DateFormatOptions = 256
	KCFISO8601DateFormatWithDay CFISO8601DateFormatOptions = 16
	KCFISO8601DateFormatWithFractionalSeconds CFISO8601DateFormatOptions = 2048
	KCFISO8601DateFormatWithFullDate CFISO8601DateFormatOptions = 1
	KCFISO8601DateFormatWithFullTime CFISO8601DateFormatOptions = 32
	KCFISO8601DateFormatWithInternetDateTime CFISO8601DateFormatOptions = 1
	KCFISO8601DateFormatWithMonth CFISO8601DateFormatOptions = 2
	KCFISO8601DateFormatWithSpaceBetweenDateAndTime CFISO8601DateFormatOptions = 128
	KCFISO8601DateFormatWithTime CFISO8601DateFormatOptions = 32
	KCFISO8601DateFormatWithTimeZone CFISO8601DateFormatOptions = 64
	KCFISO8601DateFormatWithWeekOfYear CFISO8601DateFormatOptions = 4
	KCFISO8601DateFormatWithYear CFISO8601DateFormatOptions = 1
)

func (e CFISO8601DateFormatOptions) String() string {
	switch e {
	case KCFISO8601DateFormatWithColonSeparatorInTime:
		return "KCFISO8601DateFormatWithColonSeparatorInTime"
	case KCFISO8601DateFormatWithColonSeparatorInTimeZone:
		return "KCFISO8601DateFormatWithColonSeparatorInTimeZone"
	case KCFISO8601DateFormatWithDashSeparatorInDate:
		return "KCFISO8601DateFormatWithDashSeparatorInDate"
	case KCFISO8601DateFormatWithDay:
		return "KCFISO8601DateFormatWithDay"
	case KCFISO8601DateFormatWithFractionalSeconds:
		return "KCFISO8601DateFormatWithFractionalSeconds"
	case KCFISO8601DateFormatWithFullDate:
		return "KCFISO8601DateFormatWithFullDate"
	case KCFISO8601DateFormatWithFullTime:
		return "KCFISO8601DateFormatWithFullTime"
	case KCFISO8601DateFormatWithMonth:
		return "KCFISO8601DateFormatWithMonth"
	case KCFISO8601DateFormatWithSpaceBetweenDateAndTime:
		return "KCFISO8601DateFormatWithSpaceBetweenDateAndTime"
	case KCFISO8601DateFormatWithTimeZone:
		return "KCFISO8601DateFormatWithTimeZone"
	case KCFISO8601DateFormatWithWeekOfYear:
		return "KCFISO8601DateFormatWithWeekOfYear"
	default:
		return fmt.Sprintf("CFISO8601DateFormatOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFLocaleLanguageDirection
type CFLocaleLanguageDirection int

const (
)

// See: https://developer.apple.com/documentation/CoreFoundation/CFNotificationSuspensionBehavior
type CFNotificationSuspensionBehavior int

const (
	// CFNotificationSuspensionBehaviorCoalesce: The server will only queue the last notification of the specified name and object; earlier notifications are dropped.
	CFNotificationSuspensionBehaviorCoalesce CFNotificationSuspensionBehavior = 2
	// CFNotificationSuspensionBehaviorDeliverImmediately: The server will deliver notifications of the specified name and object whether or not the application is in the background.
	CFNotificationSuspensionBehaviorDeliverImmediately CFNotificationSuspensionBehavior = 4
	// CFNotificationSuspensionBehaviorDrop: The server will not queue any notifications of the specified name and object while the receiving application is in the background.
	CFNotificationSuspensionBehaviorDrop CFNotificationSuspensionBehavior = 1
	// CFNotificationSuspensionBehaviorHold: The server will hold all matching notifications until the queue has been filled (queue size determined by the server) at which point the server may flush queued notifications.
	CFNotificationSuspensionBehaviorHold CFNotificationSuspensionBehavior = 3
)

func (e CFNotificationSuspensionBehavior) String() string {
	switch e {
	case CFNotificationSuspensionBehaviorCoalesce:
		return "CFNotificationSuspensionBehaviorCoalesce"
	case CFNotificationSuspensionBehaviorDeliverImmediately:
		return "CFNotificationSuspensionBehaviorDeliverImmediately"
	case CFNotificationSuspensionBehaviorDrop:
		return "CFNotificationSuspensionBehaviorDrop"
	case CFNotificationSuspensionBehaviorHold:
		return "CFNotificationSuspensionBehaviorHold"
	default:
		return fmt.Sprintf("CFNotificationSuspensionBehavior(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterOptionFlags
type CFNumberFormatterOptionFlags int

const (
	// KCFNumberFormatterParseIntegersOnly: Specifies that only integers should be parsed.
	KCFNumberFormatterParseIntegersOnly CFNumberFormatterOptionFlags = 1
)

func (e CFNumberFormatterOptionFlags) String() string {
	switch e {
	case KCFNumberFormatterParseIntegersOnly:
		return "KCFNumberFormatterParseIntegersOnly"
	default:
		return fmt.Sprintf("CFNumberFormatterOptionFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterPadPosition
type CFNumberFormatterPadPosition int

const (
	// KCFNumberFormatterPadAfterPrefix: Specifies the number of padding characters after the prefix.
	KCFNumberFormatterPadAfterPrefix CFNumberFormatterPadPosition = 1
	// KCFNumberFormatterPadAfterSuffix: Specifies the number of padding characters after the suffix.
	KCFNumberFormatterPadAfterSuffix CFNumberFormatterPadPosition = 3
	// KCFNumberFormatterPadBeforePrefix: Specifies the number of padding characters before the prefix.
	KCFNumberFormatterPadBeforePrefix CFNumberFormatterPadPosition = 0
	// KCFNumberFormatterPadBeforeSuffix: Specifies the number of padding characters before the suffix.
	KCFNumberFormatterPadBeforeSuffix CFNumberFormatterPadPosition = 2
)

func (e CFNumberFormatterPadPosition) String() string {
	switch e {
	case KCFNumberFormatterPadAfterPrefix:
		return "KCFNumberFormatterPadAfterPrefix"
	case KCFNumberFormatterPadAfterSuffix:
		return "KCFNumberFormatterPadAfterSuffix"
	case KCFNumberFormatterPadBeforePrefix:
		return "KCFNumberFormatterPadBeforePrefix"
	case KCFNumberFormatterPadBeforeSuffix:
		return "KCFNumberFormatterPadBeforeSuffix"
	default:
		return fmt.Sprintf("CFNumberFormatterPadPosition(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterRoundingMode
type CFNumberFormatterRoundingMode int

const (
	// KCFNumberFormatterRoundCeiling: Round towards positive infinity.
	KCFNumberFormatterRoundCeiling CFNumberFormatterRoundingMode = 0
	// KCFNumberFormatterRoundDown: Round towards zero.
	KCFNumberFormatterRoundDown CFNumberFormatterRoundingMode = 2
	// KCFNumberFormatterRoundFloor: Round towards negative infinity.
	KCFNumberFormatterRoundFloor CFNumberFormatterRoundingMode = 1
	// KCFNumberFormatterRoundHalfDown: Round towards the nearest integer, or towards zero if equidistant.
	KCFNumberFormatterRoundHalfDown CFNumberFormatterRoundingMode = 5
	// KCFNumberFormatterRoundHalfEven: Round towards the nearest integer, or towards an even number if equidistant.
	KCFNumberFormatterRoundHalfEven CFNumberFormatterRoundingMode = 4
	// KCFNumberFormatterRoundHalfUp: Round towards the nearest integer, or away from zero if equidistant.
	KCFNumberFormatterRoundHalfUp CFNumberFormatterRoundingMode = 6
	// KCFNumberFormatterRoundUp: Round away from zero.
	KCFNumberFormatterRoundUp CFNumberFormatterRoundingMode = 3
)

func (e CFNumberFormatterRoundingMode) String() string {
	switch e {
	case KCFNumberFormatterRoundCeiling:
		return "KCFNumberFormatterRoundCeiling"
	case KCFNumberFormatterRoundDown:
		return "KCFNumberFormatterRoundDown"
	case KCFNumberFormatterRoundFloor:
		return "KCFNumberFormatterRoundFloor"
	case KCFNumberFormatterRoundHalfDown:
		return "KCFNumberFormatterRoundHalfDown"
	case KCFNumberFormatterRoundHalfEven:
		return "KCFNumberFormatterRoundHalfEven"
	case KCFNumberFormatterRoundHalfUp:
		return "KCFNumberFormatterRoundHalfUp"
	case KCFNumberFormatterRoundUp:
		return "KCFNumberFormatterRoundUp"
	default:
		return fmt.Sprintf("CFNumberFormatterRoundingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberFormatterStyle
type CFNumberFormatterStyle int

const (
	// KCFNumberFormatterCurrencyStyle: Specifies a currency style format.
	KCFNumberFormatterCurrencyStyle CFNumberFormatterStyle = 2
	// KCFNumberFormatterDecimalStyle: Specifies a decimal style format.
	KCFNumberFormatterDecimalStyle CFNumberFormatterStyle = 1
	// KCFNumberFormatterNoStyle: Specifies no style.
	KCFNumberFormatterNoStyle CFNumberFormatterStyle = 0
	// KCFNumberFormatterPercentStyle: Specifies a percent style format.
	KCFNumberFormatterPercentStyle CFNumberFormatterStyle = 3
	// KCFNumberFormatterScientificStyle: Specifies a scientific style format.
	KCFNumberFormatterScientificStyle CFNumberFormatterStyle = 4
	// KCFNumberFormatterSpellOutStyle: Specifies a spelled out format.
	KCFNumberFormatterSpellOutStyle CFNumberFormatterStyle = 5
)

func (e CFNumberFormatterStyle) String() string {
	switch e {
	case KCFNumberFormatterCurrencyStyle:
		return "KCFNumberFormatterCurrencyStyle"
	case KCFNumberFormatterDecimalStyle:
		return "KCFNumberFormatterDecimalStyle"
	case KCFNumberFormatterNoStyle:
		return "KCFNumberFormatterNoStyle"
	case KCFNumberFormatterPercentStyle:
		return "KCFNumberFormatterPercentStyle"
	case KCFNumberFormatterScientificStyle:
		return "KCFNumberFormatterScientificStyle"
	case KCFNumberFormatterSpellOutStyle:
		return "KCFNumberFormatterSpellOutStyle"
	default:
		return fmt.Sprintf("CFNumberFormatterStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFNumberType
type CFNumberType int

const (
	// KCFNumberCFIndexType: CFIndex value.
	KCFNumberCFIndexType CFNumberType = 14
	// KCFNumberCGFloatType: value.
	KCFNumberCGFloatType CFNumberType = 16
	// KCFNumberCharType: Basic C  type.
	KCFNumberCharType CFNumberType = 7
	// KCFNumberDoubleType: Basic C  type.
	KCFNumberDoubleType CFNumberType = 13
	// KCFNumberFloat32Type: Thirty-two-bit real.
	KCFNumberFloat32Type CFNumberType = 5
	// KCFNumberFloat64Type: Sixty-four-bit real.
	KCFNumberFloat64Type CFNumberType = 6
	// KCFNumberFloatType: Basic C  type.
	KCFNumberFloatType CFNumberType = 12
	// KCFNumberIntType: Basic C  type.
	KCFNumberIntType CFNumberType = 9
	// KCFNumberLongLongType: Basic C  type.
	KCFNumberLongLongType CFNumberType = 11
	// KCFNumberLongType: Basic C  type.
	KCFNumberLongType CFNumberType = 10
	// KCFNumberNSIntegerType: value.
	KCFNumberNSIntegerType CFNumberType = 15
	// KCFNumberSInt16Type: Sixteen-bit, signed integer.
	KCFNumberSInt16Type CFNumberType = 2
	// KCFNumberSInt32Type: Thirty-two-bit, signed integer.
	KCFNumberSInt32Type CFNumberType = 3
	// KCFNumberSInt64Type: Sixty-four-bit, signed integer.
	KCFNumberSInt64Type CFNumberType = 4
	// KCFNumberSInt8Type: Eight-bit, signed integer.
	KCFNumberSInt8Type CFNumberType = 1
	// KCFNumberShortType: Basic C  type.
	KCFNumberShortType CFNumberType = 8
)

func (e CFNumberType) String() string {
	switch e {
	case KCFNumberCFIndexType:
		return "KCFNumberCFIndexType"
	case KCFNumberCGFloatType:
		return "KCFNumberCGFloatType"
	case KCFNumberCharType:
		return "KCFNumberCharType"
	case KCFNumberDoubleType:
		return "KCFNumberDoubleType"
	case KCFNumberFloat32Type:
		return "KCFNumberFloat32Type"
	case KCFNumberFloat64Type:
		return "KCFNumberFloat64Type"
	case KCFNumberFloatType:
		return "KCFNumberFloatType"
	case KCFNumberIntType:
		return "KCFNumberIntType"
	case KCFNumberLongLongType:
		return "KCFNumberLongLongType"
	case KCFNumberLongType:
		return "KCFNumberLongType"
	case KCFNumberNSIntegerType:
		return "KCFNumberNSIntegerType"
	case KCFNumberSInt16Type:
		return "KCFNumberSInt16Type"
	case KCFNumberSInt32Type:
		return "KCFNumberSInt32Type"
	case KCFNumberSInt64Type:
		return "KCFNumberSInt64Type"
	case KCFNumberSInt8Type:
		return "KCFNumberSInt8Type"
	case KCFNumberShortType:
		return "KCFNumberShortType"
	default:
		return fmt.Sprintf("CFNumberType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListFormat
type CFPropertyListFormat int

const (
	// KCFPropertyListBinaryFormat_v1_0: Binary format version 1.0.
	KCFPropertyListBinaryFormat_v1_0 CFPropertyListFormat = 200
	// KCFPropertyListOpenStepFormat: OpenStep format (use of this format is discouraged).
	KCFPropertyListOpenStepFormat CFPropertyListFormat = 1
	// KCFPropertyListXMLFormat_v1_0: XML format version 1.0.
	KCFPropertyListXMLFormat_v1_0 CFPropertyListFormat = 100
)

func (e CFPropertyListFormat) String() string {
	switch e {
	case KCFPropertyListBinaryFormat_v1_0:
		return "KCFPropertyListBinaryFormat_v1_0"
	case KCFPropertyListOpenStepFormat:
		return "KCFPropertyListOpenStepFormat"
	case KCFPropertyListXMLFormat_v1_0:
		return "KCFPropertyListXMLFormat_v1_0"
	default:
		return fmt.Sprintf("CFPropertyListFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFPropertyListMutabilityOptions
type CFPropertyListMutabilityOptions int

const (
	// KCFPropertyListMutableContainers: Specifies that the property list should have mutable containers but immutable leaves.
	KCFPropertyListMutableContainers CFPropertyListMutabilityOptions = 1
	// KCFPropertyListMutableContainersAndLeaves: Specifies that the property list should have mutable containers and mutable leaves.
	KCFPropertyListMutableContainersAndLeaves CFPropertyListMutabilityOptions = 2
)

func (e CFPropertyListMutabilityOptions) String() string {
	switch e {
	case KCFPropertyListMutableContainers:
		return "KCFPropertyListMutableContainers"
	case KCFPropertyListMutableContainersAndLeaves:
		return "KCFPropertyListMutableContainersAndLeaves"
	default:
		return fmt.Sprintf("CFPropertyListMutabilityOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopActivity
type CFRunLoopActivity int

const (
	// KCFRunLoopAfterWaiting: Inside the event processing loop after the run loop wakes up, but before processing the event that woke it up.
	KCFRunLoopAfterWaiting CFRunLoopActivity = 64
	// KCFRunLoopAllActivities: A combination of all the preceding stages.
	KCFRunLoopAllActivities CFRunLoopActivity = 268435455
	// KCFRunLoopBeforeSources: Inside the event processing loop before any sources are processed.
	KCFRunLoopBeforeSources CFRunLoopActivity = 4
	// KCFRunLoopBeforeTimers: Inside the event processing loop before any timers are processed.
	KCFRunLoopBeforeTimers CFRunLoopActivity = 2
	// KCFRunLoopBeforeWaiting: # Discussion
	KCFRunLoopBeforeWaiting CFRunLoopActivity = 32
	// KCFRunLoopEntry: The entrance of the run loop, before entering the event processing loop.
	KCFRunLoopEntry CFRunLoopActivity = 1
	// KCFRunLoopExit: The exit of the run loop, after exiting the event processing loop.
	KCFRunLoopExit CFRunLoopActivity = 128
)

func (e CFRunLoopActivity) String() string {
	switch e {
	case KCFRunLoopAfterWaiting:
		return "KCFRunLoopAfterWaiting"
	case KCFRunLoopAllActivities:
		return "KCFRunLoopAllActivities"
	case KCFRunLoopBeforeSources:
		return "KCFRunLoopBeforeSources"
	case KCFRunLoopBeforeTimers:
		return "KCFRunLoopBeforeTimers"
	case KCFRunLoopBeforeWaiting:
		return "KCFRunLoopBeforeWaiting"
	case KCFRunLoopEntry:
		return "KCFRunLoopEntry"
	case KCFRunLoopExit:
		return "KCFRunLoopExit"
	default:
		return fmt.Sprintf("CFRunLoopActivity(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopRunResult
type CFRunLoopRunResult int

const (
	// KCFRunLoopRunFinished: The running run loop mode has no sources or timers to process.
	KCFRunLoopRunFinished CFRunLoopRunResult = 1
	// KCFRunLoopRunHandledSource: A source has been processed.
	KCFRunLoopRunHandledSource CFRunLoopRunResult = 4
	// KCFRunLoopRunStopped: was called on the run loop.
	KCFRunLoopRunStopped CFRunLoopRunResult = 2
	// KCFRunLoopRunTimedOut: The specified time interval for running the run loop has passed.
	KCFRunLoopRunTimedOut CFRunLoopRunResult = 3
)

func (e CFRunLoopRunResult) String() string {
	switch e {
	case KCFRunLoopRunFinished:
		return "KCFRunLoopRunFinished"
	case KCFRunLoopRunHandledSource:
		return "KCFRunLoopRunHandledSource"
	case KCFRunLoopRunStopped:
		return "KCFRunLoopRunStopped"
	case KCFRunLoopRunTimedOut:
		return "KCFRunLoopRunTimedOut"
	default:
		return fmt.Sprintf("CFRunLoopRunResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketCallBackType
type CFSocketCallBackType int

const (
	// KCFSocketAcceptCallBack: New connections will be automatically accepted and the callback is called with the data argument being a pointer to a CFSocketNativeHandle of the child socket.
	KCFSocketAcceptCallBack CFSocketCallBackType = 2
	// KCFSocketConnectCallBack: # Discussion
	KCFSocketConnectCallBack CFSocketCallBackType = 4
	// KCFSocketDataCallBack: Incoming data will be read in chunks in the background and the callback is called with the data argument being a CFData object containing the read data.
	KCFSocketDataCallBack CFSocketCallBackType = 3
	// KCFSocketReadCallBack: The callback is called when data is available to be read or a new connection is waiting to be accepted.
	KCFSocketReadCallBack CFSocketCallBackType = 1
	// KCFSocketWriteCallBack: The callback is called when the socket is writable.
	KCFSocketWriteCallBack CFSocketCallBackType = 8
)

func (e CFSocketCallBackType) String() string {
	switch e {
	case KCFSocketAcceptCallBack:
		return "KCFSocketAcceptCallBack"
	case KCFSocketConnectCallBack:
		return "KCFSocketConnectCallBack"
	case KCFSocketDataCallBack:
		return "KCFSocketDataCallBack"
	case KCFSocketReadCallBack:
		return "KCFSocketReadCallBack"
	case KCFSocketWriteCallBack:
		return "KCFSocketWriteCallBack"
	default:
		return fmt.Sprintf("CFSocketCallBackType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFSocketError
type CFSocketError int

const (
	// KCFSocketError: The socket operation failed.
	KCFSocketError CFSocketError = -1
	// KCFSocketSuccess: The socket operation succeeded.
	KCFSocketSuccess CFSocketError = 0
	// KCFSocketTimeout: The socket operation timed out.
	KCFSocketTimeout CFSocketError = -2
)

func (e CFSocketError) String() string {
	switch e {
	case KCFSocketError:
		return "KCFSocketError"
	case KCFSocketSuccess:
		return "KCFSocketSuccess"
	case KCFSocketTimeout:
		return "KCFSocketTimeout"
	default:
		return fmt.Sprintf("CFSocketError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamErrorDomain
type CFStreamErrorDomain int

const (
	// KCFStreamErrorDomainCustom: The error code is a custom error code.
	KCFStreamErrorDomainCustom CFStreamErrorDomain = -1
	// KCFStreamErrorDomainMacOSStatus: The error is an OSStatus value defined in .
	KCFStreamErrorDomainMacOSStatus CFStreamErrorDomain = 2
	// KCFStreamErrorDomainPOSIX: The error code is an error code defined in .
	KCFStreamErrorDomainPOSIX CFStreamErrorDomain = 1
)

func (e CFStreamErrorDomain) String() string {
	switch e {
	case KCFStreamErrorDomainCustom:
		return "KCFStreamErrorDomainCustom"
	case KCFStreamErrorDomainMacOSStatus:
		return "KCFStreamErrorDomainMacOSStatus"
	case KCFStreamErrorDomainPOSIX:
		return "KCFStreamErrorDomainPOSIX"
	default:
		return fmt.Sprintf("CFStreamErrorDomain(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamEventType
type CFStreamEventType int

const (
	// KCFStreamEventCanAcceptBytes: The stream can accept bytes for writing.
	KCFStreamEventCanAcceptBytes CFStreamEventType = 4
	// KCFStreamEventEndEncountered: The end of the stream has been reached.
	KCFStreamEventEndEncountered CFStreamEventType = 16
	// KCFStreamEventErrorOccurred: An error has occurred on the stream.
	KCFStreamEventErrorOccurred CFStreamEventType = 8
	// KCFStreamEventHasBytesAvailable: The stream has bytes to be read.
	KCFStreamEventHasBytesAvailable CFStreamEventType = 2
	// KCFStreamEventOpenCompleted: The open has completed successfully.
	KCFStreamEventOpenCompleted CFStreamEventType = 1
)

func (e CFStreamEventType) String() string {
	switch e {
	case KCFStreamEventCanAcceptBytes:
		return "KCFStreamEventCanAcceptBytes"
	case KCFStreamEventEndEncountered:
		return "KCFStreamEventEndEncountered"
	case KCFStreamEventErrorOccurred:
		return "KCFStreamEventErrorOccurred"
	case KCFStreamEventHasBytesAvailable:
		return "KCFStreamEventHasBytesAvailable"
	case KCFStreamEventOpenCompleted:
		return "KCFStreamEventOpenCompleted"
	default:
		return fmt.Sprintf("CFStreamEventType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStreamStatus
type CFStreamStatus int

const (
	// KCFStreamStatusAtEnd: There is no more data to read, or no more data can be written.
	KCFStreamStatusAtEnd CFStreamStatus = 5
	// KCFStreamStatusClosed: The stream is closed.
	KCFStreamStatusClosed CFStreamStatus = 6
	// KCFStreamStatusError: An error occurred on the stream.
	KCFStreamStatusError CFStreamStatus = 7
	// KCFStreamStatusNotOpen: The stream is not open for reading or writing.
	KCFStreamStatusNotOpen CFStreamStatus = 0
	// KCFStreamStatusOpen: The stream is open.
	KCFStreamStatusOpen CFStreamStatus = 2
	// KCFStreamStatusOpening: The stream is being opened for reading or for writing.
	KCFStreamStatusOpening CFStreamStatus = 1
	// KCFStreamStatusReading: The stream is being read from.
	KCFStreamStatusReading CFStreamStatus = 3
	// KCFStreamStatusWriting: The stream is being written to.
	KCFStreamStatusWriting CFStreamStatus = 4
)

func (e CFStreamStatus) String() string {
	switch e {
	case KCFStreamStatusAtEnd:
		return "KCFStreamStatusAtEnd"
	case KCFStreamStatusClosed:
		return "KCFStreamStatusClosed"
	case KCFStreamStatusError:
		return "KCFStreamStatusError"
	case KCFStreamStatusNotOpen:
		return "KCFStreamStatusNotOpen"
	case KCFStreamStatusOpen:
		return "KCFStreamStatusOpen"
	case KCFStreamStatusOpening:
		return "KCFStreamStatusOpening"
	case KCFStreamStatusReading:
		return "KCFStreamStatusReading"
	case KCFStreamStatusWriting:
		return "KCFStreamStatusWriting"
	default:
		return fmt.Sprintf("CFStreamStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStringBuiltInEncodings
type CFStringBuiltInEncodings uint32

const (
	// KCFStringEncodingASCII: An encoding constant that identifies the ASCII encoding (decimal values 0 through 127).
	KCFStringEncodingASCII CFStringBuiltInEncodings = 1536
	// KCFStringEncodingISOLatin1: An encoding constant that identifies the ISO Latin 1 encoding (ISO 8859-1)
	KCFStringEncodingISOLatin1 CFStringBuiltInEncodings = 513
	// KCFStringEncodingMacRoman: An encoding constant that identifies the Mac Roman encoding.
	KCFStringEncodingMacRoman CFStringBuiltInEncodings = 0
	// KCFStringEncodingNextStepLatin: An encoding constant that identifies the NextStep/OpenStep encoding.
	KCFStringEncodingNextStepLatin CFStringBuiltInEncodings = 2817
	// KCFStringEncodingNonLossyASCII: An encoding constant that identifies non-lossy ASCII encoding.
	KCFStringEncodingNonLossyASCII CFStringBuiltInEncodings = 3071
	// KCFStringEncodingUTF16BE: An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16BEFormat encoding.
	KCFStringEncodingUTF16BE CFStringBuiltInEncodings = 268435712
	// KCFStringEncodingUTF16LE: An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF16LEFormat encoding.
	KCFStringEncodingUTF16LE CFStringBuiltInEncodings = 335544576
	// KCFStringEncodingUTF32: An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32Format encoding.
	KCFStringEncodingUTF32 CFStringBuiltInEncodings = 201326848
	// KCFStringEncodingUTF32BE: An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32BEFormat encoding.
	KCFStringEncodingUTF32BE CFStringBuiltInEncodings = 402653440
	// KCFStringEncodingUTF32LE: An encoding constant that identifies kTextEncodingUnicodeDefault + kUnicodeUTF32LEFormat encoding.
	KCFStringEncodingUTF32LE CFStringBuiltInEncodings = 469762304
	// KCFStringEncodingUTF8: An encoding constant that identifies the UTF 8 encoding.
	KCFStringEncodingUTF8 CFStringBuiltInEncodings = 134217984
	// KCFStringEncodingUnicode: An encoding constant that identifies the Unicode encoding.
	KCFStringEncodingUnicode CFStringBuiltInEncodings = 256
	// KCFStringEncodingWindowsLatin1: An encoding constant that identifies the Windows Latin 1 encoding (ANSI codepage 1252).
	KCFStringEncodingWindowsLatin1 CFStringBuiltInEncodings = 1280
)

func (e CFStringBuiltInEncodings) String() string {
	switch e {
	case KCFStringEncodingASCII:
		return "KCFStringEncodingASCII"
	case KCFStringEncodingISOLatin1:
		return "KCFStringEncodingISOLatin1"
	case KCFStringEncodingMacRoman:
		return "KCFStringEncodingMacRoman"
	case KCFStringEncodingNextStepLatin:
		return "KCFStringEncodingNextStepLatin"
	case KCFStringEncodingNonLossyASCII:
		return "KCFStringEncodingNonLossyASCII"
	case KCFStringEncodingUTF16BE:
		return "KCFStringEncodingUTF16BE"
	case KCFStringEncodingUTF16LE:
		return "KCFStringEncodingUTF16LE"
	case KCFStringEncodingUTF32:
		return "KCFStringEncodingUTF32"
	case KCFStringEncodingUTF32BE:
		return "KCFStringEncodingUTF32BE"
	case KCFStringEncodingUTF32LE:
		return "KCFStringEncodingUTF32LE"
	case KCFStringEncodingUTF8:
		return "KCFStringEncodingUTF8"
	case KCFStringEncodingUnicode:
		return "KCFStringEncodingUnicode"
	case KCFStringEncodingWindowsLatin1:
		return "KCFStringEncodingWindowsLatin1"
	default:
		return fmt.Sprintf("CFStringBuiltInEncodings(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStringCompareFlags
type CFStringCompareFlags int

const (
	// KCFCompareAnchored: Performs searching only on characters at the beginning or end of the range.
	KCFCompareAnchored CFStringCompareFlags = 8
	// KCFCompareBackwards: Specifies that the comparison should start at the last elements of the entities being compared (for example, strings or arrays).
	KCFCompareBackwards CFStringCompareFlags = 4
	// KCFCompareCaseInsensitive: Specifies that the comparison should ignore differences in case between alphabetical characters.
	KCFCompareCaseInsensitive CFStringCompareFlags = 1
	// KCFCompareDiacriticInsensitive: Specifies that the comparison should ignore diacritic markers.
	KCFCompareDiacriticInsensitive CFStringCompareFlags = 128
	// KCFCompareForcedOrdering: Specifies that the comparison is forced to return either `kCFCompareLessThan` or `kCFCompareGreaterThan` if the strings are equivalent but not strictly equal.
	KCFCompareForcedOrdering CFStringCompareFlags = 512
	// KCFCompareLocalized: Specifies that the comparison should take into account differences related to locale, such as the thousands separator character.
	KCFCompareLocalized CFStringCompareFlags = 32
	// KCFCompareNonliteral: Specifies that loose equivalence is acceptable, especially as pertains to diacritical marks.
	KCFCompareNonliteral CFStringCompareFlags = 16
	// KCFCompareNumerically: Specifies that represented numeric values should be used as the basis for comparison and not the actual character values.
	KCFCompareNumerically CFStringCompareFlags = 64
	// KCFCompareWidthInsensitive: Specifies that the comparison should ignore width differences.
	KCFCompareWidthInsensitive CFStringCompareFlags = 256
)

func (e CFStringCompareFlags) String() string {
	switch e {
	case KCFCompareAnchored:
		return "KCFCompareAnchored"
	case KCFCompareBackwards:
		return "KCFCompareBackwards"
	case KCFCompareCaseInsensitive:
		return "KCFCompareCaseInsensitive"
	case KCFCompareDiacriticInsensitive:
		return "KCFCompareDiacriticInsensitive"
	case KCFCompareForcedOrdering:
		return "KCFCompareForcedOrdering"
	case KCFCompareLocalized:
		return "KCFCompareLocalized"
	case KCFCompareNonliteral:
		return "KCFCompareNonliteral"
	case KCFCompareNumerically:
		return "KCFCompareNumerically"
	case KCFCompareWidthInsensitive:
		return "KCFCompareWidthInsensitive"
	default:
		return fmt.Sprintf("CFStringCompareFlags(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStringEncodings
type CFStringEncodings int

const (
	// KCFStringEncodingANSEL: ANSEL (ANSI Z39.47).
	KCFStringEncodingANSEL CFStringEncodings = 1537
	// KCFStringEncodingBig5: Big-5 (has variants)
	KCFStringEncodingBig5 CFStringEncodings = 2563
	// KCFStringEncodingBig5_E: Taiwan Big-5E standard.
	KCFStringEncodingBig5_E CFStringEncodings = 2569
	// KCFStringEncodingBig5_HKSCS_1999: Big-5 with Hong Kong special char set supplement.
	KCFStringEncodingBig5_HKSCS_1999 CFStringEncodings = 2566
	// KCFStringEncodingCNS_11643_92_P1: CNS 11643-1992 plane 1.
	KCFStringEncodingCNS_11643_92_P1 CFStringEncodings = 1617
	// KCFStringEncodingCNS_11643_92_P2: CNS 11643-1992 plane 2.
	KCFStringEncodingCNS_11643_92_P2 CFStringEncodings = 1618
	// KCFStringEncodingCNS_11643_92_P3: CNS 11643-1992 plane 3 (was plane 14 in 1986 version).
	KCFStringEncodingCNS_11643_92_P3 CFStringEncodings = 1619
	// KCFStringEncodingDOSArabic: Code page 864.
	KCFStringEncodingDOSArabic CFStringEncodings = 1049
	// KCFStringEncodingDOSBalticRim: Code page 775.
	KCFStringEncodingDOSBalticRim CFStringEncodings = 1030
	// KCFStringEncodingDOSCanadianFrench: Code page 863.
	KCFStringEncodingDOSCanadianFrench CFStringEncodings = 1048
	// KCFStringEncodingDOSChineseSimplif: Code page 936, also for Windows.
	KCFStringEncodingDOSChineseSimplif CFStringEncodings = 1057
	// KCFStringEncodingDOSChineseTrad: Code page 950, also for Windows.
	KCFStringEncodingDOSChineseTrad CFStringEncodings = 1059
	// KCFStringEncodingDOSCyrillic: Code page 855, IBM Cyrillic.
	KCFStringEncodingDOSCyrillic CFStringEncodings = 1043
	// KCFStringEncodingDOSGreek: Code page 737 (formerly code page 437G).
	KCFStringEncodingDOSGreek CFStringEncodings = 1029
	// KCFStringEncodingDOSGreek1: Code page 851.
	KCFStringEncodingDOSGreek1 CFStringEncodings = 1041
	// KCFStringEncodingDOSGreek2: Code page 869, IBM Modern Greek.
	KCFStringEncodingDOSGreek2 CFStringEncodings = 1052
	// KCFStringEncodingDOSHebrew: Code page 862.
	KCFStringEncodingDOSHebrew CFStringEncodings = 1047
	// KCFStringEncodingDOSIcelandic: Code page 861.
	KCFStringEncodingDOSIcelandic CFStringEncodings = 1046
	// KCFStringEncodingDOSJapanese: Code page 932, also for Windows.
	KCFStringEncodingDOSJapanese CFStringEncodings = 1056
	// KCFStringEncodingDOSKorean: Code page 949, also for Windows; Unified Hangul Code.
	KCFStringEncodingDOSKorean CFStringEncodings = 1058
	// KCFStringEncodingDOSLatin1: Code page 850, “Multilingual”.
	KCFStringEncodingDOSLatin1 CFStringEncodings = 1040
	// KCFStringEncodingDOSLatin2: Code page 852, Slavic.
	KCFStringEncodingDOSLatin2 CFStringEncodings = 1042
	// KCFStringEncodingDOSLatinUS: Code page 437.
	KCFStringEncodingDOSLatinUS CFStringEncodings = 1024
	// KCFStringEncodingDOSNordic: Code page 865.
	KCFStringEncodingDOSNordic CFStringEncodings = 1050
	// KCFStringEncodingDOSPortuguese: Code page 860.
	KCFStringEncodingDOSPortuguese CFStringEncodings = 1045
	// KCFStringEncodingDOSRussian: Code page 866.
	KCFStringEncodingDOSRussian CFStringEncodings = 1051
	// KCFStringEncodingDOSThai: Code page 874, also for Windows.
	KCFStringEncodingDOSThai CFStringEncodings = 1053
	// KCFStringEncodingDOSTurkish: Code page 857, IBM Turkish.
	KCFStringEncodingDOSTurkish CFStringEncodings = 1044
	// KCFStringEncodingEBCDIC_CP037: code page 037, extended EBCDIC (Latin-1 set) for US, Canada.
	KCFStringEncodingEBCDIC_CP037 CFStringEncodings = 3074
	// KCFStringEncodingEBCDIC_US: basic EBCDIC-US
	KCFStringEncodingEBCDIC_US CFStringEncodings = 3073
	// KCFStringEncodingEUC_CN: ISO 646, GB 2312-80.
	KCFStringEncodingEUC_CN CFStringEncodings = 2352
	// KCFStringEncodingEUC_JP: ISO 646, 1-byte katakana, JIS 208, JIS 212.
	KCFStringEncodingEUC_JP CFStringEncodings = 2336
	// KCFStringEncodingEUC_KR: ISO 646, KS C 5601-1987.
	KCFStringEncodingEUC_KR CFStringEncodings = 2368
	// KCFStringEncodingEUC_TW: ISO 646, CNS 11643-1992 Planes 1-16.
	KCFStringEncodingEUC_TW CFStringEncodings = 2353
	// KCFStringEncodingGBK_95: Annex to GB 13000-93; for Windows 95.
	KCFStringEncodingGBK_95 CFStringEncodings = 1585
	// KCFStringEncodingHZ_GB_2312: HZ (RFC 1842, for Chinese mail & news).
	KCFStringEncodingHZ_GB_2312 CFStringEncodings = 2565
	// KCFStringEncodingISOLatin10: ISO 8859-16.
	KCFStringEncodingISOLatin10 CFStringEncodings = 528
	// KCFStringEncodingISOLatin2: ISO 8859-2.
	KCFStringEncodingISOLatin2 CFStringEncodings = 514
	// KCFStringEncodingISOLatin3: ISO 8859-3.
	KCFStringEncodingISOLatin3 CFStringEncodings = 515
	// KCFStringEncodingISOLatin4: ISO 8859-4.
	KCFStringEncodingISOLatin4 CFStringEncodings = 516
	// KCFStringEncodingISOLatin5: ISO 8859-9.
	KCFStringEncodingISOLatin5 CFStringEncodings = 521
	// KCFStringEncodingISOLatin6: ISO 8859-10.
	KCFStringEncodingISOLatin6 CFStringEncodings = 522
	// KCFStringEncodingISOLatin7: ISO 8859-13.
	KCFStringEncodingISOLatin7 CFStringEncodings = 525
	// KCFStringEncodingISOLatin8: ISO 8859-14.
	KCFStringEncodingISOLatin8 CFStringEncodings = 526
	// KCFStringEncodingISOLatin9: ISO 8859-15.
	KCFStringEncodingISOLatin9 CFStringEncodings = 527
	// KCFStringEncodingISOLatinArabic: ISO 8859-6, =ASMO 708, =DOS CP 708.
	KCFStringEncodingISOLatinArabic CFStringEncodings = 518
	// KCFStringEncodingISOLatinCyrillic: ISO 8859-5.
	KCFStringEncodingISOLatinCyrillic CFStringEncodings = 517
	// KCFStringEncodingISOLatinGreek: ISO 8859-7.
	KCFStringEncodingISOLatinGreek CFStringEncodings = 519
	// KCFStringEncodingISOLatinHebrew: ISO 8859-8.
	KCFStringEncodingISOLatinHebrew CFStringEncodings = 520
	// KCFStringEncodingISOLatinThai: ISO 8859-11.
	KCFStringEncodingISOLatinThai CFStringEncodings = 523
	// KCFStringEncodingISO_2022_JP_1: RFC 2237.
	KCFStringEncodingISO_2022_JP_1 CFStringEncodings = 2082
	// KCFStringEncodingISO_2022_JP_3: JIS X0213.
	KCFStringEncodingISO_2022_JP_3 CFStringEncodings = 2083
	// KCFStringEncodingKOI8_R: Russian internet standard.
	KCFStringEncodingKOI8_R CFStringEncodings = 2562
	// KCFStringEncodingKOI8_U: RFC 2319, Ukrainian.
	KCFStringEncodingKOI8_U CFStringEncodings = 2568
	// KCFStringEncodingKSC_5601_87: Same as KSC 5601-92 without Johab annex.
	KCFStringEncodingKSC_5601_87 CFStringEncodings = 1600
	// KCFStringEncodingKSC_5601_92_Johab: KSC 5601-92 Johab annex.
	KCFStringEncodingKSC_5601_92_Johab CFStringEncodings = 1601
	// KCFStringEncodingMacFarsi: Like MacArabic but uses Farsi digits.
	KCFStringEncodingMacFarsi CFStringEncodings = 140
	// KCFStringEncodingMacHFS: Meta-value, should never appear in a table.
	KCFStringEncodingMacHFS CFStringEncodings = 255
	// KCFStringEncodingMacRomanLatin1: Mac OS Roman permuted to align with ISO Latin-1.
	KCFStringEncodingMacRomanLatin1 CFStringEncodings = 2564
	// KCFStringEncodingMacVT100: VT100102 font from Comm Toolbox: Latin-1 repertoire + box drawing etc.
	KCFStringEncodingMacVT100 CFStringEncodings = 252
	// KCFStringEncodingNextStepJapanese: NextStep Japanese encoding.
	KCFStringEncodingNextStepJapanese CFStringEncodings = 2818
	// KCFStringEncodingShiftJIS: Plain Shift-JIS.
	KCFStringEncodingShiftJIS CFStringEncodings = 2561
	// KCFStringEncodingShiftJIS_X0213: Shift-JIS format encoding of JIS X0213 planes 1 and 2.
	KCFStringEncodingShiftJIS_X0213 CFStringEncodings = 1576
	// KCFStringEncodingShiftJIS_X0213_MenKuTen: JIS X0213 in plane-row-column notation.
	KCFStringEncodingShiftJIS_X0213_MenKuTen CFStringEncodings = 1577
	// KCFStringEncodingUTF7: kTextEncodingUnicodeDefault + kUnicodeUTF7Format RFC2152.
	KCFStringEncodingUTF7 CFStringEncodings = 67109120
	// KCFStringEncodingUTF7_IMAP: UTF-7 (IMAP folder variant) RFC3501.
	KCFStringEncodingUTF7_IMAP CFStringEncodings = 2576
	// KCFStringEncodingVISCII: RFC 1456, Vietnamese.
	KCFStringEncodingVISCII CFStringEncodings = 2567
	// KCFStringEncodingWindowsArabic: Code page 1256.
	KCFStringEncodingWindowsArabic CFStringEncodings = 1286
	// KCFStringEncodingWindowsBalticRim: Code page 1257.
	KCFStringEncodingWindowsBalticRim CFStringEncodings = 1287
	// KCFStringEncodingWindowsCyrillic: Code page 1251, Slavic Cyrillic.
	KCFStringEncodingWindowsCyrillic CFStringEncodings = 1282
	// KCFStringEncodingWindowsGreek: Code page 1253.
	KCFStringEncodingWindowsGreek CFStringEncodings = 1283
	// KCFStringEncodingWindowsHebrew: Code page 1255.
	KCFStringEncodingWindowsHebrew CFStringEncodings = 1285
	// KCFStringEncodingWindowsKoreanJohab: Code page 1361, for Windows NT.
	KCFStringEncodingWindowsKoreanJohab CFStringEncodings = 1296
	// KCFStringEncodingWindowsLatin2: Code page 1250, Central Europe.
	KCFStringEncodingWindowsLatin2 CFStringEncodings = 1281
	// KCFStringEncodingWindowsLatin5: Code page 1254, Turkish.
	KCFStringEncodingWindowsLatin5 CFStringEncodings = 1284
	// KCFStringEncodingWindowsVietnamese: Code page 1258.
	KCFStringEncodingWindowsVietnamese CFStringEncodings = 1288
)

func (e CFStringEncodings) String() string {
	switch e {
	case KCFStringEncodingANSEL:
		return "KCFStringEncodingANSEL"
	case KCFStringEncodingBig5:
		return "KCFStringEncodingBig5"
	case KCFStringEncodingBig5_E:
		return "KCFStringEncodingBig5_E"
	case KCFStringEncodingBig5_HKSCS_1999:
		return "KCFStringEncodingBig5_HKSCS_1999"
	case KCFStringEncodingCNS_11643_92_P1:
		return "KCFStringEncodingCNS_11643_92_P1"
	case KCFStringEncodingCNS_11643_92_P2:
		return "KCFStringEncodingCNS_11643_92_P2"
	case KCFStringEncodingCNS_11643_92_P3:
		return "KCFStringEncodingCNS_11643_92_P3"
	case KCFStringEncodingDOSArabic:
		return "KCFStringEncodingDOSArabic"
	case KCFStringEncodingDOSBalticRim:
		return "KCFStringEncodingDOSBalticRim"
	case KCFStringEncodingDOSCanadianFrench:
		return "KCFStringEncodingDOSCanadianFrench"
	case KCFStringEncodingDOSChineseSimplif:
		return "KCFStringEncodingDOSChineseSimplif"
	case KCFStringEncodingDOSChineseTrad:
		return "KCFStringEncodingDOSChineseTrad"
	case KCFStringEncodingDOSCyrillic:
		return "KCFStringEncodingDOSCyrillic"
	case KCFStringEncodingDOSGreek:
		return "KCFStringEncodingDOSGreek"
	case KCFStringEncodingDOSGreek1:
		return "KCFStringEncodingDOSGreek1"
	case KCFStringEncodingDOSGreek2:
		return "KCFStringEncodingDOSGreek2"
	case KCFStringEncodingDOSHebrew:
		return "KCFStringEncodingDOSHebrew"
	case KCFStringEncodingDOSIcelandic:
		return "KCFStringEncodingDOSIcelandic"
	case KCFStringEncodingDOSJapanese:
		return "KCFStringEncodingDOSJapanese"
	case KCFStringEncodingDOSKorean:
		return "KCFStringEncodingDOSKorean"
	case KCFStringEncodingDOSLatin1:
		return "KCFStringEncodingDOSLatin1"
	case KCFStringEncodingDOSLatin2:
		return "KCFStringEncodingDOSLatin2"
	case KCFStringEncodingDOSLatinUS:
		return "KCFStringEncodingDOSLatinUS"
	case KCFStringEncodingDOSNordic:
		return "KCFStringEncodingDOSNordic"
	case KCFStringEncodingDOSPortuguese:
		return "KCFStringEncodingDOSPortuguese"
	case KCFStringEncodingDOSRussian:
		return "KCFStringEncodingDOSRussian"
	case KCFStringEncodingDOSThai:
		return "KCFStringEncodingDOSThai"
	case KCFStringEncodingDOSTurkish:
		return "KCFStringEncodingDOSTurkish"
	case KCFStringEncodingEBCDIC_CP037:
		return "KCFStringEncodingEBCDIC_CP037"
	case KCFStringEncodingEBCDIC_US:
		return "KCFStringEncodingEBCDIC_US"
	case KCFStringEncodingEUC_CN:
		return "KCFStringEncodingEUC_CN"
	case KCFStringEncodingEUC_JP:
		return "KCFStringEncodingEUC_JP"
	case KCFStringEncodingEUC_KR:
		return "KCFStringEncodingEUC_KR"
	case KCFStringEncodingEUC_TW:
		return "KCFStringEncodingEUC_TW"
	case KCFStringEncodingGBK_95:
		return "KCFStringEncodingGBK_95"
	case KCFStringEncodingHZ_GB_2312:
		return "KCFStringEncodingHZ_GB_2312"
	case KCFStringEncodingISOLatin10:
		return "KCFStringEncodingISOLatin10"
	case KCFStringEncodingISOLatin2:
		return "KCFStringEncodingISOLatin2"
	case KCFStringEncodingISOLatin3:
		return "KCFStringEncodingISOLatin3"
	case KCFStringEncodingISOLatin4:
		return "KCFStringEncodingISOLatin4"
	case KCFStringEncodingISOLatin5:
		return "KCFStringEncodingISOLatin5"
	case KCFStringEncodingISOLatin6:
		return "KCFStringEncodingISOLatin6"
	case KCFStringEncodingISOLatin7:
		return "KCFStringEncodingISOLatin7"
	case KCFStringEncodingISOLatin8:
		return "KCFStringEncodingISOLatin8"
	case KCFStringEncodingISOLatin9:
		return "KCFStringEncodingISOLatin9"
	case KCFStringEncodingISOLatinArabic:
		return "KCFStringEncodingISOLatinArabic"
	case KCFStringEncodingISOLatinCyrillic:
		return "KCFStringEncodingISOLatinCyrillic"
	case KCFStringEncodingISOLatinGreek:
		return "KCFStringEncodingISOLatinGreek"
	case KCFStringEncodingISOLatinHebrew:
		return "KCFStringEncodingISOLatinHebrew"
	case KCFStringEncodingISOLatinThai:
		return "KCFStringEncodingISOLatinThai"
	case KCFStringEncodingISO_2022_JP_1:
		return "KCFStringEncodingISO_2022_JP_1"
	case KCFStringEncodingISO_2022_JP_3:
		return "KCFStringEncodingISO_2022_JP_3"
	case KCFStringEncodingKOI8_R:
		return "KCFStringEncodingKOI8_R"
	case KCFStringEncodingKOI8_U:
		return "KCFStringEncodingKOI8_U"
	case KCFStringEncodingKSC_5601_87:
		return "KCFStringEncodingKSC_5601_87"
	case KCFStringEncodingKSC_5601_92_Johab:
		return "KCFStringEncodingKSC_5601_92_Johab"
	case KCFStringEncodingMacFarsi:
		return "KCFStringEncodingMacFarsi"
	case KCFStringEncodingMacHFS:
		return "KCFStringEncodingMacHFS"
	case KCFStringEncodingMacRomanLatin1:
		return "KCFStringEncodingMacRomanLatin1"
	case KCFStringEncodingMacVT100:
		return "KCFStringEncodingMacVT100"
	case KCFStringEncodingNextStepJapanese:
		return "KCFStringEncodingNextStepJapanese"
	case KCFStringEncodingShiftJIS:
		return "KCFStringEncodingShiftJIS"
	case KCFStringEncodingShiftJIS_X0213:
		return "KCFStringEncodingShiftJIS_X0213"
	case KCFStringEncodingShiftJIS_X0213_MenKuTen:
		return "KCFStringEncodingShiftJIS_X0213_MenKuTen"
	case KCFStringEncodingUTF7:
		return "KCFStringEncodingUTF7"
	case KCFStringEncodingUTF7_IMAP:
		return "KCFStringEncodingUTF7_IMAP"
	case KCFStringEncodingVISCII:
		return "KCFStringEncodingVISCII"
	case KCFStringEncodingWindowsArabic:
		return "KCFStringEncodingWindowsArabic"
	case KCFStringEncodingWindowsBalticRim:
		return "KCFStringEncodingWindowsBalticRim"
	case KCFStringEncodingWindowsCyrillic:
		return "KCFStringEncodingWindowsCyrillic"
	case KCFStringEncodingWindowsGreek:
		return "KCFStringEncodingWindowsGreek"
	case KCFStringEncodingWindowsHebrew:
		return "KCFStringEncodingWindowsHebrew"
	case KCFStringEncodingWindowsKoreanJohab:
		return "KCFStringEncodingWindowsKoreanJohab"
	case KCFStringEncodingWindowsLatin2:
		return "KCFStringEncodingWindowsLatin2"
	case KCFStringEncodingWindowsLatin5:
		return "KCFStringEncodingWindowsLatin5"
	case KCFStringEncodingWindowsVietnamese:
		return "KCFStringEncodingWindowsVietnamese"
	default:
		return fmt.Sprintf("CFStringEncodings(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStringNormalizationForm
type CFStringNormalizationForm int

const (
	// KCFStringNormalizationFormC: Canonical decomposition followed by canonical composition.
	KCFStringNormalizationFormC CFStringNormalizationForm = 2
	// KCFStringNormalizationFormD: Canonical decomposition.
	KCFStringNormalizationFormD CFStringNormalizationForm = 0
	// KCFStringNormalizationFormKC: Compatibility decomposition followed by canonical composition.
	KCFStringNormalizationFormKC CFStringNormalizationForm = 3
	// KCFStringNormalizationFormKD: Compatibility decomposition.
	KCFStringNormalizationFormKD CFStringNormalizationForm = 1
)

func (e CFStringNormalizationForm) String() string {
	switch e {
	case KCFStringNormalizationFormC:
		return "KCFStringNormalizationFormC"
	case KCFStringNormalizationFormD:
		return "KCFStringNormalizationFormD"
	case KCFStringNormalizationFormKC:
		return "KCFStringNormalizationFormKC"
	case KCFStringNormalizationFormKD:
		return "KCFStringNormalizationFormKD"
	default:
		return fmt.Sprintf("CFStringNormalizationForm(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFStringTokenizerTokenType
type CFStringTokenizerTokenType int

const (
	// KCFStringTokenizerTokenHasDerivedSubTokensMask: Compound token which may contain derived subtokens.
	KCFStringTokenizerTokenHasDerivedSubTokensMask CFStringTokenizerTokenType = 4
	// KCFStringTokenizerTokenHasHasNumbersMask: Appears to contain a number.
	KCFStringTokenizerTokenHasHasNumbersMask CFStringTokenizerTokenType = 8
	// KCFStringTokenizerTokenHasNonLettersMask: Contains punctuation, symbols, and so on.
	KCFStringTokenizerTokenHasNonLettersMask CFStringTokenizerTokenType = 16
	// KCFStringTokenizerTokenHasSubTokensMask: Compound token which may contain subtokens but with no derived subtokens.
	KCFStringTokenizerTokenHasSubTokensMask CFStringTokenizerTokenType = 2
	// KCFStringTokenizerTokenIsCJWordMask: Contains kana and/or ideographs.
	KCFStringTokenizerTokenIsCJWordMask CFStringTokenizerTokenType = 32
	// KCFStringTokenizerTokenNormal: Has a normal token.
	KCFStringTokenizerTokenNormal CFStringTokenizerTokenType = 1
)

func (e CFStringTokenizerTokenType) String() string {
	switch e {
	case KCFStringTokenizerTokenHasDerivedSubTokensMask:
		return "KCFStringTokenizerTokenHasDerivedSubTokensMask"
	case KCFStringTokenizerTokenHasHasNumbersMask:
		return "KCFStringTokenizerTokenHasHasNumbersMask"
	case KCFStringTokenizerTokenHasNonLettersMask:
		return "KCFStringTokenizerTokenHasNonLettersMask"
	case KCFStringTokenizerTokenHasSubTokensMask:
		return "KCFStringTokenizerTokenHasSubTokensMask"
	case KCFStringTokenizerTokenIsCJWordMask:
		return "KCFStringTokenizerTokenIsCJWordMask"
	case KCFStringTokenizerTokenNormal:
		return "KCFStringTokenizerTokenNormal"
	default:
		return fmt.Sprintf("CFStringTokenizerTokenType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFTimeZoneNameStyle
type CFTimeZoneNameStyle int

const (
	// KCFTimeZoneNameStyleDaylightSaving: Specifies the daylight saving name style; for example, “Central Daylight Time” for the Central time zone.
	KCFTimeZoneNameStyleDaylightSaving CFTimeZoneNameStyle = 2
	// KCFTimeZoneNameStyleGeneric: Specifies the generic name style, which does not distinguish between daylight saving and standard time; for example, “Central Time” for the Central time zone.
	KCFTimeZoneNameStyleGeneric CFTimeZoneNameStyle = 4
	// KCFTimeZoneNameStyleShortDaylightSaving: Specifies the short daylight saving name style; for example, “CDT” for the Central time zone.
	KCFTimeZoneNameStyleShortDaylightSaving CFTimeZoneNameStyle = 3
	// KCFTimeZoneNameStyleShortGeneric: Specifies the short generic name style, which does not distinguish between daylight saving and standard time; for example, “CT” for the Central time zone.
	KCFTimeZoneNameStyleShortGeneric CFTimeZoneNameStyle = 5
	// KCFTimeZoneNameStyleShortStandard: Specifies the short standard name style; for example, “CST” for the Central time zone.
	KCFTimeZoneNameStyleShortStandard CFTimeZoneNameStyle = 1
	// KCFTimeZoneNameStyleStandard: Specifies the standard name style; for example, “Central Standard Time” for the Central time zone.
	KCFTimeZoneNameStyleStandard CFTimeZoneNameStyle = 0
)

func (e CFTimeZoneNameStyle) String() string {
	switch e {
	case KCFTimeZoneNameStyleDaylightSaving:
		return "KCFTimeZoneNameStyleDaylightSaving"
	case KCFTimeZoneNameStyleGeneric:
		return "KCFTimeZoneNameStyleGeneric"
	case KCFTimeZoneNameStyleShortDaylightSaving:
		return "KCFTimeZoneNameStyleShortDaylightSaving"
	case KCFTimeZoneNameStyleShortGeneric:
		return "KCFTimeZoneNameStyleShortGeneric"
	case KCFTimeZoneNameStyleShortStandard:
		return "KCFTimeZoneNameStyleShortStandard"
	case KCFTimeZoneNameStyleStandard:
		return "KCFTimeZoneNameStyleStandard"
	default:
		return fmt.Sprintf("CFTimeZoneNameStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkCreationOptions
type CFURLBookmarkCreationOptions int

const (
	// KCFURLBookmarkCreationMinimalBookmarkMask: Specifies that an alias created with the bookmark data be created with minimal information, which may make it smaller but still able to resolve in certain ways.
	KCFURLBookmarkCreationMinimalBookmarkMask CFURLBookmarkCreationOptions = 512
	// KCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess: When combined with the KCFURLBookmarkCreationWithSecurityScope option, specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read-only access to a file-system resource; for use in an app that adopts App Sandbox.
	KCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess CFURLBookmarkCreationOptions = 4096
	// KCFURLBookmarkCreationSuitableForBookmarkFile: Specifies that the bookmark data include properties required to create Finder alias files.
	KCFURLBookmarkCreationSuitableForBookmarkFile CFURLBookmarkCreationOptions = 1024
	// KCFURLBookmarkCreationWithSecurityScope: Specifies that you want to create a security-scoped bookmark that, when resolved, provides a security-scoped URL allowing read/write access to a file-system resource; for use in an app that adopts App Sandbox.
	KCFURLBookmarkCreationWithSecurityScope CFURLBookmarkCreationOptions = 2048
	KCFURLBookmarkCreationWithoutImplicitSecurityScope CFURLBookmarkCreationOptions = 536870912
	// Deprecated.
	KCFURLBookmarkCreationPreferFileIDResolutionMask CFURLBookmarkCreationOptions = 536870913
)

func (e CFURLBookmarkCreationOptions) String() string {
	switch e {
	case KCFURLBookmarkCreationMinimalBookmarkMask:
		return "KCFURLBookmarkCreationMinimalBookmarkMask"
	case KCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess:
		return "KCFURLBookmarkCreationSecurityScopeAllowOnlyReadAccess"
	case KCFURLBookmarkCreationSuitableForBookmarkFile:
		return "KCFURLBookmarkCreationSuitableForBookmarkFile"
	case KCFURLBookmarkCreationWithSecurityScope:
		return "KCFURLBookmarkCreationWithSecurityScope"
	case KCFURLBookmarkCreationWithoutImplicitSecurityScope:
		return "KCFURLBookmarkCreationWithoutImplicitSecurityScope"
	case KCFURLBookmarkCreationPreferFileIDResolutionMask:
		return "KCFURLBookmarkCreationPreferFileIDResolutionMask"
	default:
		return fmt.Sprintf("CFURLBookmarkCreationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLBookmarkResolutionOptions
type CFURLBookmarkResolutionOptions int

const (
	// KCFBookmarkResolutionWithoutMountingMask: Specifies that no volume should be mounted during resolution of the bookmark data.
	KCFBookmarkResolutionWithoutMountingMask CFURLBookmarkResolutionOptions = 512
	// KCFBookmarkResolutionWithoutUIMask: Specifies that no UI feedback accompany resolution of the bookmark data.
	KCFBookmarkResolutionWithoutUIMask CFURLBookmarkResolutionOptions = 256
	// KCFURLBookmarkResolutionWithSecurityScope: Specifies that the security scope, applied to the bookmark when it was created, should be used during resolution of the bookmark data.
	KCFURLBookmarkResolutionWithSecurityScope CFURLBookmarkResolutionOptions = 1024
	KCFURLBookmarkResolutionWithoutImplicitStartAccessing CFURLBookmarkResolutionOptions = 32768
	KCFURLBookmarkResolutionWithoutMountingMask CFURLBookmarkResolutionOptions = 512
	KCFURLBookmarkResolutionWithoutUIMask CFURLBookmarkResolutionOptions = 256
)

func (e CFURLBookmarkResolutionOptions) String() string {
	switch e {
	case KCFBookmarkResolutionWithoutMountingMask:
		return "KCFBookmarkResolutionWithoutMountingMask"
	case KCFBookmarkResolutionWithoutUIMask:
		return "KCFBookmarkResolutionWithoutUIMask"
	case KCFURLBookmarkResolutionWithSecurityScope:
		return "KCFURLBookmarkResolutionWithSecurityScope"
	case KCFURLBookmarkResolutionWithoutImplicitStartAccessing:
		return "KCFURLBookmarkResolutionWithoutImplicitStartAccessing"
	default:
		return fmt.Sprintf("CFURLBookmarkResolutionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLComponentType
type CFURLComponentType int

const (
	// KCFURLComponentFragment: The URL’s fragment.
	KCFURLComponentFragment CFURLComponentType = 12
	// KCFURLComponentHost: The URL’s host.
	KCFURLComponentHost CFURLComponentType = 8
	// KCFURLComponentNetLocation: The URL’s network location.
	KCFURLComponentNetLocation CFURLComponentType = 2
	// KCFURLComponentParameterString: The URL’s parameter string.
	KCFURLComponentParameterString CFURLComponentType = 10
	// KCFURLComponentPassword: The user’s password.
	KCFURLComponentPassword CFURLComponentType = 6
	// KCFURLComponentPath: The URL’s path component.
	KCFURLComponentPath CFURLComponentType = 3
	// KCFURLComponentPort: The URL’s port.
	KCFURLComponentPort CFURLComponentType = 9
	// KCFURLComponentQuery: The URL’s query.
	KCFURLComponentQuery CFURLComponentType = 11
	// KCFURLComponentResourceSpecifier: The URL’s resource specifier.
	KCFURLComponentResourceSpecifier CFURLComponentType = 4
	// KCFURLComponentScheme: The URL’s scheme.
	KCFURLComponentScheme CFURLComponentType = 1
	// KCFURLComponentUser: The URL’s user.
	KCFURLComponentUser CFURLComponentType = 5
	// KCFURLComponentUserInfo: The user’s information.
	KCFURLComponentUserInfo CFURLComponentType = 7
)

func (e CFURLComponentType) String() string {
	switch e {
	case KCFURLComponentFragment:
		return "KCFURLComponentFragment"
	case KCFURLComponentHost:
		return "KCFURLComponentHost"
	case KCFURLComponentNetLocation:
		return "KCFURLComponentNetLocation"
	case KCFURLComponentParameterString:
		return "KCFURLComponentParameterString"
	case KCFURLComponentPassword:
		return "KCFURLComponentPassword"
	case KCFURLComponentPath:
		return "KCFURLComponentPath"
	case KCFURLComponentPort:
		return "KCFURLComponentPort"
	case KCFURLComponentQuery:
		return "KCFURLComponentQuery"
	case KCFURLComponentResourceSpecifier:
		return "KCFURLComponentResourceSpecifier"
	case KCFURLComponentScheme:
		return "KCFURLComponentScheme"
	case KCFURLComponentUser:
		return "KCFURLComponentUser"
	case KCFURLComponentUserInfo:
		return "KCFURLComponentUserInfo"
	default:
		return fmt.Sprintf("CFURLComponentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorOptions
type CFURLEnumeratorOptions int

const (
	// KCFURLEnumeratorDescendRecursively: The enumerator recurses into each subdirectory enumerated.
	KCFURLEnumeratorDescendRecursively CFURLEnumeratorOptions = 1
	// KCFURLEnumeratorGenerateFileReferenceURLs: The enumerator generates file reference URLs instead of file path URLs.
	KCFURLEnumeratorGenerateFileReferenceURLs CFURLEnumeratorOptions = 4
	KCFURLEnumeratorGenerateRelativePathURLs CFURLEnumeratorOptions = 64
	// KCFURLEnumeratorIncludeDirectoriesPostOrder: If provided along with the KCFURLEnumeratorDescendRecursively option, the recursive enumerator returns a directory’s URL after returning the URLs of the directory’s descendents.
	KCFURLEnumeratorIncludeDirectoriesPostOrder CFURLEnumeratorOptions = 32
	// KCFURLEnumeratorIncludeDirectoriesPreOrder: If provided along with the KCFURLEnumeratorDescendRecursively option, the recursive enumerator returns a directory’s URL before returning the URLs of the directory’s descendents.
	KCFURLEnumeratorIncludeDirectoriesPreOrder CFURLEnumeratorOptions = 16
	// KCFURLEnumeratorSkipInvisibles: The enumerator skips “hidden” or “invisible” objects.
	KCFURLEnumeratorSkipInvisibles CFURLEnumeratorOptions = 2
	// KCFURLEnumeratorSkipPackageContents: The enumerator skips package directory contents.
	KCFURLEnumeratorSkipPackageContents CFURLEnumeratorOptions = 8
)

func (e CFURLEnumeratorOptions) String() string {
	switch e {
	case KCFURLEnumeratorDescendRecursively:
		return "KCFURLEnumeratorDescendRecursively"
	case KCFURLEnumeratorGenerateFileReferenceURLs:
		return "KCFURLEnumeratorGenerateFileReferenceURLs"
	case KCFURLEnumeratorGenerateRelativePathURLs:
		return "KCFURLEnumeratorGenerateRelativePathURLs"
	case KCFURLEnumeratorIncludeDirectoriesPostOrder:
		return "KCFURLEnumeratorIncludeDirectoriesPostOrder"
	case KCFURLEnumeratorIncludeDirectoriesPreOrder:
		return "KCFURLEnumeratorIncludeDirectoriesPreOrder"
	case KCFURLEnumeratorSkipInvisibles:
		return "KCFURLEnumeratorSkipInvisibles"
	case KCFURLEnumeratorSkipPackageContents:
		return "KCFURLEnumeratorSkipPackageContents"
	default:
		return fmt.Sprintf("CFURLEnumeratorOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLEnumeratorResult
type CFURLEnumeratorResult int

const (
	// KCFURLEnumeratorDirectoryPostOrderSuccess: The recursive post-order enumerator returned the URL for a directory after having returned the URLs for all of the directory’s descendents.
	KCFURLEnumeratorDirectoryPostOrderSuccess CFURLEnumeratorResult = 4
	// KCFURLEnumeratorEnd: The enumeration is complete.
	KCFURLEnumeratorEnd CFURLEnumeratorResult = 2
	// KCFURLEnumeratorError: An error occurred during enumeration.
	KCFURLEnumeratorError CFURLEnumeratorResult = 3
	// KCFURLEnumeratorSuccess: The enumerator was advanced successfully and returned a valid URL.
	KCFURLEnumeratorSuccess CFURLEnumeratorResult = 1
)

func (e CFURLEnumeratorResult) String() string {
	switch e {
	case KCFURLEnumeratorDirectoryPostOrderSuccess:
		return "KCFURLEnumeratorDirectoryPostOrderSuccess"
	case KCFURLEnumeratorEnd:
		return "KCFURLEnumeratorEnd"
	case KCFURLEnumeratorError:
		return "KCFURLEnumeratorError"
	case KCFURLEnumeratorSuccess:
		return "KCFURLEnumeratorSuccess"
	default:
		return fmt.Sprintf("CFURLEnumeratorResult(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLError
type CFURLError int

const (
	// KCFURLImproperArgumentsError: Indicates one or more arguments are improper.
	KCFURLImproperArgumentsError CFURLError = -15
	// KCFURLPropertyKeyUnavailableError: Indicates a property key was unavailable.
	KCFURLPropertyKeyUnavailableError CFURLError = -17
	// KCFURLRemoteHostUnavailableError: Indicates a remote host is unavailable.
	KCFURLRemoteHostUnavailableError CFURLError = -14
	// KCFURLResourceAccessViolationError: Indicates an error in accessing a resource.
	KCFURLResourceAccessViolationError CFURLError = -13
	// KCFURLResourceNotFoundError: Indicates a resource was not found.
	KCFURLResourceNotFoundError CFURLError = -12
	// KCFURLTimeoutError: Indicates a timeout.
	KCFURLTimeoutError CFURLError = -18
	// KCFURLUnknownError: Indicates an unknown error.
	KCFURLUnknownError CFURLError = -10
	// KCFURLUnknownPropertyKeyError: Indicates a property key is unknown.
	KCFURLUnknownPropertyKeyError CFURLError = -16
	// KCFURLUnknownSchemeError: Indicates that the scheme is not recognized.
	KCFURLUnknownSchemeError CFURLError = -11
)

func (e CFURLError) String() string {
	switch e {
	case KCFURLImproperArgumentsError:
		return "KCFURLImproperArgumentsError"
	case KCFURLPropertyKeyUnavailableError:
		return "KCFURLPropertyKeyUnavailableError"
	case KCFURLRemoteHostUnavailableError:
		return "KCFURLRemoteHostUnavailableError"
	case KCFURLResourceAccessViolationError:
		return "KCFURLResourceAccessViolationError"
	case KCFURLResourceNotFoundError:
		return "KCFURLResourceNotFoundError"
	case KCFURLTimeoutError:
		return "KCFURLTimeoutError"
	case KCFURLUnknownError:
		return "KCFURLUnknownError"
	case KCFURLUnknownPropertyKeyError:
		return "KCFURLUnknownPropertyKeyError"
	case KCFURLUnknownSchemeError:
		return "KCFURLUnknownSchemeError"
	default:
		return fmt.Sprintf("CFURLError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFURLPathStyle
type CFURLPathStyle int

const (
	// KCFURLHFSPathStyle: Indicates a HFS style path name.
	KCFURLHFSPathStyle CFURLPathStyle = 1
	// KCFURLPOSIXPathStyle: Indicates a POSIX style path name.
	KCFURLPOSIXPathStyle CFURLPathStyle = 0
	// KCFURLWindowsPathStyle: Indicates a Windows style path name.
	KCFURLWindowsPathStyle CFURLPathStyle = 2
)

func (e CFURLPathStyle) String() string {
	switch e {
	case KCFURLHFSPathStyle:
		return "KCFURLHFSPathStyle"
	case KCFURLPOSIXPathStyle:
		return "KCFURLPOSIXPathStyle"
	case KCFURLWindowsPathStyle:
		return "KCFURLWindowsPathStyle"
	default:
		return fmt.Sprintf("CFURLPathStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLEntityTypeCode
type CFXMLEntityTypeCode int

const (
	// KCFXMLEntityTypeCharacter: Indicates a character entity type.
	KCFXMLEntityTypeCharacter CFXMLEntityTypeCode = 4
	// KCFXMLEntityTypeParameter: Implies a parsed, internal entity.
	KCFXMLEntityTypeParameter CFXMLEntityTypeCode = 0
	// KCFXMLEntityTypeParsedExternal: Indicates a parsed, external entity.
	KCFXMLEntityTypeParsedExternal CFXMLEntityTypeCode = 2
	// KCFXMLEntityTypeParsedInternal: Indicates a parsed, internal entity.
	KCFXMLEntityTypeParsedInternal CFXMLEntityTypeCode = 1
	// KCFXMLEntityTypeUnparsed: Indicates an unparsed entity.
	KCFXMLEntityTypeUnparsed CFXMLEntityTypeCode = 3
)

func (e CFXMLEntityTypeCode) String() string {
	switch e {
	case KCFXMLEntityTypeCharacter:
		return "KCFXMLEntityTypeCharacter"
	case KCFXMLEntityTypeParameter:
		return "KCFXMLEntityTypeParameter"
	case KCFXMLEntityTypeParsedExternal:
		return "KCFXMLEntityTypeParsedExternal"
	case KCFXMLEntityTypeParsedInternal:
		return "KCFXMLEntityTypeParsedInternal"
	case KCFXMLEntityTypeUnparsed:
		return "KCFXMLEntityTypeUnparsed"
	default:
		return fmt.Sprintf("CFXMLEntityTypeCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserOptions
type CFXMLParserOptions int

const (
	// KCFXMLParserAddImpliedAttributes: Where the DTD specifies implied attribute-value pairs for a particular element, add those pairs to any occurrences of the element in the element tree.
	KCFXMLParserAddImpliedAttributes CFXMLParserOptions = 32
	// KCFXMLParserAllOptions: Makes the parser do the most work, returning only the pure elementtree.
	KCFXMLParserAllOptions CFXMLParserOptions = 16777215
	// KCFXMLParserReplacePhysicalEntities: Replaces declared entities like `&lt`;.
	KCFXMLParserReplacePhysicalEntities CFXMLParserOptions = 4
	// KCFXMLParserResolveExternalEntities: Resolves all external entities.
	KCFXMLParserResolveExternalEntities CFXMLParserOptions = 16
	// KCFXMLParserSkipMetaData: Silently skip over metadata constructs (the DTD and comments).
	KCFXMLParserSkipMetaData CFXMLParserOptions = 2
	// KCFXMLParserSkipWhitespace: # Discussion
	KCFXMLParserSkipWhitespace CFXMLParserOptions = 8
	// KCFXMLParserValidateDocument: Validates the document against its grammar from the DTD, reporting any errors.
	KCFXMLParserValidateDocument CFXMLParserOptions = 1
)

func (e CFXMLParserOptions) String() string {
	switch e {
	case KCFXMLParserAddImpliedAttributes:
		return "KCFXMLParserAddImpliedAttributes"
	case KCFXMLParserAllOptions:
		return "KCFXMLParserAllOptions"
	case KCFXMLParserReplacePhysicalEntities:
		return "KCFXMLParserReplacePhysicalEntities"
	case KCFXMLParserResolveExternalEntities:
		return "KCFXMLParserResolveExternalEntities"
	case KCFXMLParserSkipMetaData:
		return "KCFXMLParserSkipMetaData"
	case KCFXMLParserSkipWhitespace:
		return "KCFXMLParserSkipWhitespace"
	case KCFXMLParserValidateDocument:
		return "KCFXMLParserValidateDocument"
	default:
		return fmt.Sprintf("CFXMLParserOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CFXMLParserStatusCode
type CFXMLParserStatusCode int

const (
	// KCFXMLErrorElementlessDocument: Indicates a document containing no elements.
	KCFXMLErrorElementlessDocument CFXMLParserStatusCode = 11
	// KCFXMLErrorEncodingConversionFailure: Indicates an encoding conversion error.
	KCFXMLErrorEncodingConversionFailure CFXMLParserStatusCode = 3
	// KCFXMLErrorMalformedCDSect: Indicates a malformed CDATA section.
	KCFXMLErrorMalformedCDSect CFXMLParserStatusCode = 7
	// KCFXMLErrorMalformedCharacterReference: Indicates a malformed character reference.
	KCFXMLErrorMalformedCharacterReference CFXMLParserStatusCode = 13
	// KCFXMLErrorMalformedCloseTag: Indicates a malformed close tag.
	KCFXMLErrorMalformedCloseTag CFXMLParserStatusCode = 8
	// KCFXMLErrorMalformedComment: Indicates a malformed comment.
	KCFXMLErrorMalformedComment CFXMLParserStatusCode = 12
	// KCFXMLErrorMalformedDTD: Indicates a malformed DTD.
	KCFXMLErrorMalformedDTD CFXMLParserStatusCode = 5
	// KCFXMLErrorMalformedDocument: Indicates a malformed document.
	KCFXMLErrorMalformedDocument CFXMLParserStatusCode = 10
	// KCFXMLErrorMalformedName: Indicates a malformed name.
	KCFXMLErrorMalformedName CFXMLParserStatusCode = 6
	// KCFXMLErrorMalformedParsedCharacterData: Indicates malformed character data.
	KCFXMLErrorMalformedParsedCharacterData CFXMLParserStatusCode = 14
	// KCFXMLErrorMalformedProcessingInstruction: Indicates a malformed processing instruction.
	KCFXMLErrorMalformedProcessingInstruction CFXMLParserStatusCode = 4
	// KCFXMLErrorMalformedStartTag: Indicates a malformed start tag.
	KCFXMLErrorMalformedStartTag CFXMLParserStatusCode = 9
	// KCFXMLErrorNoData: Indicates a no data error.
	KCFXMLErrorNoData CFXMLParserStatusCode = 15
	// KCFXMLErrorUnexpectedEOF: Indicates an unexpected EOF occurred.
	KCFXMLErrorUnexpectedEOF CFXMLParserStatusCode = 1
	// KCFXMLErrorUnknownEncoding: Indicates an unknown encoding error.
	KCFXMLErrorUnknownEncoding CFXMLParserStatusCode = 2
	// KCFXMLStatusParseInProgress: Indicates the parser is in progress.
	KCFXMLStatusParseInProgress CFXMLParserStatusCode = -1
	// KCFXMLStatusParseNotBegun: Indicates the parser has not begun.
	KCFXMLStatusParseNotBegun CFXMLParserStatusCode = -2
)

func (e CFXMLParserStatusCode) String() string {
	switch e {
	case KCFXMLErrorElementlessDocument:
		return "KCFXMLErrorElementlessDocument"
	case KCFXMLErrorEncodingConversionFailure:
		return "KCFXMLErrorEncodingConversionFailure"
	case KCFXMLErrorMalformedCDSect:
		return "KCFXMLErrorMalformedCDSect"
	case KCFXMLErrorMalformedCharacterReference:
		return "KCFXMLErrorMalformedCharacterReference"
	case KCFXMLErrorMalformedCloseTag:
		return "KCFXMLErrorMalformedCloseTag"
	case KCFXMLErrorMalformedComment:
		return "KCFXMLErrorMalformedComment"
	case KCFXMLErrorMalformedDTD:
		return "KCFXMLErrorMalformedDTD"
	case KCFXMLErrorMalformedDocument:
		return "KCFXMLErrorMalformedDocument"
	case KCFXMLErrorMalformedName:
		return "KCFXMLErrorMalformedName"
	case KCFXMLErrorMalformedParsedCharacterData:
		return "KCFXMLErrorMalformedParsedCharacterData"
	case KCFXMLErrorMalformedProcessingInstruction:
		return "KCFXMLErrorMalformedProcessingInstruction"
	case KCFXMLErrorMalformedStartTag:
		return "KCFXMLErrorMalformedStartTag"
	case KCFXMLErrorNoData:
		return "KCFXMLErrorNoData"
	case KCFXMLErrorUnexpectedEOF:
		return "KCFXMLErrorUnexpectedEOF"
	case KCFXMLErrorUnknownEncoding:
		return "KCFXMLErrorUnknownEncoding"
	case KCFXMLStatusParseInProgress:
		return "KCFXMLStatusParseInProgress"
	case KCFXMLStatusParseNotBegun:
		return "KCFXMLStatusParseNotBegun"
	default:
		return fmt.Sprintf("CFXMLParserStatusCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/CoreFoundation/CGRectEdge
type CGRectEdge uint32

const (
	CGRectMaxXEdge CGRectEdge = 2
	CGRectMinXEdge CGRectEdge = 0
	CGRectMinYEdge CGRectEdge = 1
)

func (e CGRectEdge) String() string {
	switch e {
	case CGRectMaxXEdge:
		return "CGRectMaxXEdge"
	case CGRectMinXEdge:
		return "CGRectMinXEdge"
	case CGRectMinYEdge:
		return "CGRectMinYEdge"
	default:
		return fmt.Sprintf("CGRectEdge(%d)", e)
	}
}

type KCFBundleExecutableArchitecture uint

const (
	KCFBundleExecutableArchitectureARM64 KCFBundleExecutableArchitecture = 16777228
	// KCFBundleExecutableArchitectureI386: Specifies the 32-bit Intel architecture.
	KCFBundleExecutableArchitectureI386 KCFBundleExecutableArchitecture = 7
	// KCFBundleExecutableArchitecturePPC: Specifies the 32-bit PowerPC architecture.
	KCFBundleExecutableArchitecturePPC KCFBundleExecutableArchitecture = 18
	// KCFBundleExecutableArchitecturePPC64: Specifies the 64-bit PowerPC architecture.
	KCFBundleExecutableArchitecturePPC64 KCFBundleExecutableArchitecture = 16777234
	// KCFBundleExecutableArchitectureX86_64: Specifies the 64-bit Intel architecture.
	KCFBundleExecutableArchitectureX86_64 KCFBundleExecutableArchitecture = 16777223
)

func (e KCFBundleExecutableArchitecture) String() string {
	switch e {
	case KCFBundleExecutableArchitectureARM64:
		return "KCFBundleExecutableArchitectureARM64"
	case KCFBundleExecutableArchitectureI386:
		return "KCFBundleExecutableArchitectureI386"
	case KCFBundleExecutableArchitecturePPC:
		return "KCFBundleExecutableArchitecturePPC"
	case KCFBundleExecutableArchitecturePPC64:
		return "KCFBundleExecutableArchitecturePPC64"
	case KCFBundleExecutableArchitectureX86_64:
		return "KCFBundleExecutableArchitectureX86_64"
	default:
		return fmt.Sprintf("KCFBundleExecutableArchitecture(%d)", e)
	}
}

const KCFCalendarComponentsWrap uint = 1

type KCFFileDescriptor uint

const (
	// KCFFileDescriptorReadCallBack: Identifies the read callback.
	KCFFileDescriptorReadCallBack KCFFileDescriptor = 1
	// KCFFileDescriptorWriteCallBack: Identifies the write callback.
	KCFFileDescriptorWriteCallBack KCFFileDescriptor = 2
)

func (e KCFFileDescriptor) String() string {
	switch e {
	case KCFFileDescriptorReadCallBack:
		return "KCFFileDescriptorReadCallBack"
	case KCFFileDescriptorWriteCallBack:
		return "KCFFileDescriptorWriteCallBack"
	default:
		return fmt.Sprintf("KCFFileDescriptor(%d)", e)
	}
}

type KCFMessagePort int

const (
	// KCFMessagePortBecameInvalidError: The message port was invalidated.
	KCFMessagePortBecameInvalidError KCFMessagePort = -5
	// KCFMessagePortIsInvalid: The message could not be sent because the message port is invalid.
	KCFMessagePortIsInvalid KCFMessagePort = -3
	// KCFMessagePortReceiveTimeout: No reply was received before the receive timeout.
	KCFMessagePortReceiveTimeout KCFMessagePort = -2
	// KCFMessagePortSendTimeout: The message could not be sent before the send timeout.
	KCFMessagePortSendTimeout KCFMessagePort = -1
	// KCFMessagePortSuccess: The message was successfully sent and, if a reply was expected, a reply was received.
	KCFMessagePortSuccess KCFMessagePort = 0
	// KCFMessagePortTransportError: An error occurred trying to send the message.
	KCFMessagePortTransportError KCFMessagePort = -4
)

func (e KCFMessagePort) String() string {
	switch e {
	case KCFMessagePortBecameInvalidError:
		return "KCFMessagePortBecameInvalidError"
	case KCFMessagePortIsInvalid:
		return "KCFMessagePortIsInvalid"
	case KCFMessagePortReceiveTimeout:
		return "KCFMessagePortReceiveTimeout"
	case KCFMessagePortSendTimeout:
		return "KCFMessagePortSendTimeout"
	case KCFMessagePortSuccess:
		return "KCFMessagePortSuccess"
	case KCFMessagePortTransportError:
		return "KCFMessagePortTransportError"
	default:
		return fmt.Sprintf("KCFMessagePort(%d)", e)
	}
}

type KCFNotification uint

const (
	// KCFNotificationDeliverImmediately: Delivers the notification immediately.
	KCFNotificationDeliverImmediately KCFNotification = 1
	// KCFNotificationPostToAllSessions: Delivers the notification to all sessions.
	KCFNotificationPostToAllSessions KCFNotification = 2
)

func (e KCFNotification) String() string {
	switch e {
	case KCFNotificationDeliverImmediately:
		return "KCFNotificationDeliverImmediately"
	case KCFNotificationPostToAllSessions:
		return "KCFNotificationPostToAllSessions"
	default:
		return fmt.Sprintf("KCFNotification(%d)", e)
	}
}

type KCFPropertyList uint

const (
	// KCFPropertyListReadCorruptError: Signifies an error parsing a property list.
	KCFPropertyListReadCorruptError KCFPropertyList = 3840
	// KCFPropertyListReadStreamError: Signifies a stream error reading a property list.
	KCFPropertyListReadStreamError KCFPropertyList = 3842
	// KCFPropertyListReadUnknownVersionError: Signifies the version number in the property list is unknown.
	KCFPropertyListReadUnknownVersionError KCFPropertyList = 3841
	// KCFPropertyListWriteStreamError: Signifies a stream error writing a property list.
	KCFPropertyListWriteStreamError KCFPropertyList = 3851
)

func (e KCFPropertyList) String() string {
	switch e {
	case KCFPropertyListReadCorruptError:
		return "KCFPropertyListReadCorruptError"
	case KCFPropertyListReadStreamError:
		return "KCFPropertyListReadStreamError"
	case KCFPropertyListReadUnknownVersionError:
		return "KCFPropertyListReadUnknownVersionError"
	case KCFPropertyListWriteStreamError:
		return "KCFPropertyListWriteStreamError"
	default:
		return fmt.Sprintf("KCFPropertyList(%d)", e)
	}
}

type KCFSocket uint

const (
	// KCFSocketAutomaticallyReenableAcceptCallBack: # Discussion
	KCFSocketAutomaticallyReenableAcceptCallBack KCFSocket = 2
	// KCFSocketAutomaticallyReenableDataCallBack: # Discussion
	KCFSocketAutomaticallyReenableDataCallBack KCFSocket = 3
	// KCFSocketAutomaticallyReenableReadCallBack: # Discussion
	KCFSocketAutomaticallyReenableReadCallBack KCFSocket = 1
	// KCFSocketAutomaticallyReenableWriteCallBack: # Discussion
	KCFSocketAutomaticallyReenableWriteCallBack KCFSocket = 8
	// KCFSocketCloseOnInvalidate: When enabled using CFSocketSetSocketFlags(_:_:), the native socket associated with a CFSocket object is closed when the CFSocket object is invalidated.
	KCFSocketCloseOnInvalidate KCFSocket = 128
	// KCFSocketLeaveErrors: # Discussion
	KCFSocketLeaveErrors KCFSocket = 64
)

func (e KCFSocket) String() string {
	switch e {
	case KCFSocketAutomaticallyReenableAcceptCallBack:
		return "KCFSocketAutomaticallyReenableAcceptCallBack"
	case KCFSocketAutomaticallyReenableDataCallBack:
		return "KCFSocketAutomaticallyReenableDataCallBack"
	case KCFSocketAutomaticallyReenableReadCallBack:
		return "KCFSocketAutomaticallyReenableReadCallBack"
	case KCFSocketAutomaticallyReenableWriteCallBack:
		return "KCFSocketAutomaticallyReenableWriteCallBack"
	case KCFSocketCloseOnInvalidate:
		return "KCFSocketCloseOnInvalidate"
	case KCFSocketLeaveErrors:
		return "KCFSocketLeaveErrors"
	default:
		return fmt.Sprintf("KCFSocket(%d)", e)
	}
}

type KCFStringTokenizer uint

const (
	// KCFStringTokenizerAttributeLanguage: Tells the tokenizer to prepare the language (specified as an RFC 3066bis string) when it tokenizes the string.
	KCFStringTokenizerAttributeLanguage KCFStringTokenizer = 131072
	// KCFStringTokenizerAttributeLatinTranscription: Used with `kCFStringTokenizerUnitWord`, tells the tokenizer to prepare the Latin transcription when it tokenizes the string.
	KCFStringTokenizerAttributeLatinTranscription KCFStringTokenizer = 65536
	// KCFStringTokenizerUnitLineBreak: Specifies that a string should be tokenized by line break.
	KCFStringTokenizerUnitLineBreak KCFStringTokenizer = 3
	// KCFStringTokenizerUnitParagraph: Specifies that a string should be tokenized by paragraph.
	KCFStringTokenizerUnitParagraph KCFStringTokenizer = 2
	// KCFStringTokenizerUnitSentence: Specifies that a string should be tokenized by sentence.
	KCFStringTokenizerUnitSentence KCFStringTokenizer = 1
	// KCFStringTokenizerUnitWord: Specifies that a string should be tokenized by word.
	KCFStringTokenizerUnitWord KCFStringTokenizer = 0
	// KCFStringTokenizerUnitWordBoundary: Specifies that a string should be tokenized by locale-sensitive word boundary.
	KCFStringTokenizerUnitWordBoundary KCFStringTokenizer = 4
)

func (e KCFStringTokenizer) String() string {
	switch e {
	case KCFStringTokenizerAttributeLanguage:
		return "KCFStringTokenizerAttributeLanguage"
	case KCFStringTokenizerAttributeLatinTranscription:
		return "KCFStringTokenizerAttributeLatinTranscription"
	case KCFStringTokenizerUnitLineBreak:
		return "KCFStringTokenizerUnitLineBreak"
	case KCFStringTokenizerUnitParagraph:
		return "KCFStringTokenizerUnitParagraph"
	case KCFStringTokenizerUnitSentence:
		return "KCFStringTokenizerUnitSentence"
	case KCFStringTokenizerUnitWord:
		return "KCFStringTokenizerUnitWord"
	case KCFStringTokenizerUnitWordBoundary:
		return "KCFStringTokenizerUnitWordBoundary"
	default:
		return fmt.Sprintf("KCFStringTokenizer(%d)", e)
	}
}

type KCFUserNotification uint

const (
	// KCFUserNotificationAlternateResponse: The alternate button was pressed.
	KCFUserNotificationAlternateResponse KCFUserNotification = 1
	// KCFUserNotificationCancelResponse: No button was pressed and the notification timed out.
	KCFUserNotificationCancelResponse KCFUserNotification = 3
	// KCFUserNotificationCautionAlertLevel: The notification is somewhat serious.
	KCFUserNotificationCautionAlertLevel KCFUserNotification = 2
	// KCFUserNotificationDefaultResponse: The default button was pressed.
	KCFUserNotificationDefaultResponse KCFUserNotification = 0
	// KCFUserNotificationNoDefaultButtonFlag: Displays the dialog without the default, alternate, or other buttons.
	KCFUserNotificationNoDefaultButtonFlag KCFUserNotification = 32
	// KCFUserNotificationNoteAlertLevel: The notification is not very serious.
	KCFUserNotificationNoteAlertLevel KCFUserNotification = 1
	// KCFUserNotificationOtherResponse: The third button was pressed.
	KCFUserNotificationOtherResponse KCFUserNotification = 2
	// KCFUserNotificationPlainAlertLevel: The notification is not serious.
	KCFUserNotificationPlainAlertLevel KCFUserNotification = 3
	// KCFUserNotificationStopAlertLevel: The notification is very serious.
	KCFUserNotificationStopAlertLevel KCFUserNotification = 0
	// KCFUserNotificationUseRadioButtonsFlag: Creates a group of radio buttons instead of checkboxes for the elements in the kCFUserNotificationCheckBoxTitlesKey array in the user notification’s description dictionary.
	KCFUserNotificationUseRadioButtonsFlag KCFUserNotification = 64
)

func (e KCFUserNotification) String() string {
	switch e {
	case KCFUserNotificationAlternateResponse:
		return "KCFUserNotificationAlternateResponse"
	case KCFUserNotificationCancelResponse:
		return "KCFUserNotificationCancelResponse"
	case KCFUserNotificationCautionAlertLevel:
		return "KCFUserNotificationCautionAlertLevel"
	case KCFUserNotificationDefaultResponse:
		return "KCFUserNotificationDefaultResponse"
	case KCFUserNotificationNoDefaultButtonFlag:
		return "KCFUserNotificationNoDefaultButtonFlag"
	case KCFUserNotificationUseRadioButtonsFlag:
		return "KCFUserNotificationUseRadioButtonsFlag"
	default:
		return fmt.Sprintf("KCFUserNotification(%d)", e)
	}
}

