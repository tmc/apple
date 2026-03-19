// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DateFormatter] class.
var (
	_DateFormatterClass     DateFormatterClass
	_DateFormatterClassOnce sync.Once
)

func getDateFormatterClass() DateFormatterClass {
	_DateFormatterClassOnce.Do(func() {
		_DateFormatterClass = DateFormatterClass{class: objc.GetClass("NSDateFormatter")}
	})
	return _DateFormatterClass
}

// GetDateFormatterClass returns the class object for NSDateFormatter.
func GetDateFormatterClass() DateFormatterClass {
	return getDateFormatterClass()
}

type DateFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (dc DateFormatterClass) Alloc() DateFormatter {
	rv := objc.Send[DateFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that converts between dates and their textual representations.
//
// # Overview
// 
// Instances of [NSDateFormatter] create string representations of [NSDate]
// objects, and convert textual representations of dates and times into
// [NSDate] objects. For user-visible representations of dates and times,
// [NSDateFormatter] provides a variety of localized presets and configuration
// options. For fixed format representations of dates and times, you can
// specify a custom format string.
// 
// When working with date representations in ISO 8601 format, use
// [NSISO8601DateFormatter] instead.
// 
// To represent an interval between two [NSDate] objects, use
// [NSDateIntervalFormatter] instead.
// 
// To represent a quantity of time specified by an [NSDateComponents] object,
// use [NSDateComponentsFormatter] instead.
// 
// # Working With User-Visible Representations of Dates and Times
// 
// When displaying a date to a user, you set the [DateStyle] and [TimeStyle]
// properties of the date formatter according to your particular needs. For
// example, if you want to show the month, day, and year without showing the
// time, you would set the [DateStyle] property to [DateFormatterLongStyle]
// and the [TimeStyle] property to [DateFormatterNoStyle]. Conversely, if you
// want to show only the time, you would set the `dateStyle` property to
// [DateFormatterNoStyle] and the [TimeStyle] property to
// [DateFormatterShortStyle]. Based on the values of the [DateStyle] and
// [TimeStyle] properties, [NSDateFormatter] provides a representation of a
// specified date that is appropriate for a given locale.
// 
// If you need to define a format that cannot be achieved using the predefined
// styles, you can use the [SetLocalizedDateFormatFromTemplate] to specify a
// localized date format from a template.
// 
// # Working With Fixed Format Date Representations
// 
// When working with fixed format dates, such as RFC 3339, you set the
// [DateFormat] property to specify a format string. For most fixed formats,
// you should also set the [Locale] property to a POSIX locale
// (`"en_US_POSIX"`), and set the [TimeZone] property to UTC.
// 
// For more information, see [Technical Q&A QA1480 “NSDateFormatter and
// Internet Dates”].
// 
// # Thread Safety
// 
// On iOS 7 and later [NSDateFormatter] is thread safe.
// 
// In macOS 10.9 and later [NSDateFormatter] is thread safe so long as you are
// using the modern behavior in a 64-bit app.
// 
// On earlier versions of the operating system, or when using the legacy
// formatter behavior or running in 32-bit in macOS, [NSDateFormatter] is not
// thread safe, and you therefore must not mutate a date formatter
// simultaneously from multiple threads.
//
// [Technical Q&A QA1480 “NSDateFormatter and Internet Dates”]: https://developer.apple.com/library/mac/qa/qa1480/
//
// # Converting Objects
//
//   - [DateFormatter.DateFromString]: Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
//   - [DateFormatter.StringFromDate]: Returns a string representation of a specified date that the system formats using the receiver’s current settings.
//   - [DateFormatter.GetObjectValueForStringRangeError]: Returns by reference a date representation of a specified string and its date range, as well as a Boolean value that indicates whether the system can parse the string.
//
// # Managing Formats and Styles
//
//   - [DateFormatter.DateStyle]: The date style of the receiver.
//   - [DateFormatter.SetDateStyle]
//   - [DateFormatter.TimeStyle]: The time style of the receiver.
//   - [DateFormatter.SetTimeStyle]
//   - [DateFormatter.DateFormat]: The date format string used by the receiver.
//   - [DateFormatter.SetDateFormat]
//   - [DateFormatter.SetLocalizedDateFormatFromTemplate]: Sets the date format from a template using the specified locale for the receiver.
//   - [DateFormatter.FormattingContext]: The capitalization formatting context used when formatting a date.
//   - [DateFormatter.SetFormattingContext]
//
// # Managing Attributes
//
//   - [DateFormatter.Calendar]: The calendar for the receiver.
//   - [DateFormatter.SetCalendar]
//   - [DateFormatter.DefaultDate]: The default date for the receiver.
//   - [DateFormatter.SetDefaultDate]
//   - [DateFormatter.Locale]: The locale for the receiver.
//   - [DateFormatter.SetLocale]
//   - [DateFormatter.TimeZone]: The time zone for the receiver.
//   - [DateFormatter.SetTimeZone]
//   - [DateFormatter.TwoDigitStartDate]: The earliest date that can be denoted by a two-digit year specifier.
//   - [DateFormatter.SetTwoDigitStartDate]
//   - [DateFormatter.GregorianStartDate]: The start date of the Gregorian calendar for the receiver.
//   - [DateFormatter.SetGregorianStartDate]
//
// # Managing Behavior Version
//
//   - [DateFormatter.FormatterBehavior]: The formatter behavior for the receiver.
//   - [DateFormatter.SetFormatterBehavior]
//
// # Managing Natural Language Support
//
//   - [DateFormatter.Lenient]: A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//   - [DateFormatter.SetLenient]
//   - [DateFormatter.DoesRelativeDateFormatting]: A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//   - [DateFormatter.SetDoesRelativeDateFormatting]
//
// # Managing AM and PM Symbols
//
//   - [DateFormatter.AMSymbol]: The AM symbol for the receiver.
//   - [DateFormatter.SetAMSymbol]
//   - [DateFormatter.PMSymbol]: The PM symbol for the receiver.
//   - [DateFormatter.SetPMSymbol]
//
// # Managing Weekday Symbols
//
//   - [DateFormatter.WeekdaySymbols]: The array of weekday symbols for the receiver.
//   - [DateFormatter.SetWeekdaySymbols]
//   - [DateFormatter.ShortWeekdaySymbols]: The array of short weekday symbols for the receiver.
//   - [DateFormatter.SetShortWeekdaySymbols]
//   - [DateFormatter.VeryShortWeekdaySymbols]: The array of very short weekday symbols for the receiver.
//   - [DateFormatter.SetVeryShortWeekdaySymbols]
//   - [DateFormatter.StandaloneWeekdaySymbols]: The array of standalone weekday symbols for the receiver.
//   - [DateFormatter.SetStandaloneWeekdaySymbols]
//   - [DateFormatter.ShortStandaloneWeekdaySymbols]: The array of short standalone weekday symbols for the receiver.
//   - [DateFormatter.SetShortStandaloneWeekdaySymbols]
//   - [DateFormatter.VeryShortStandaloneWeekdaySymbols]: The array of very short standalone weekday symbols for the receiver.
//   - [DateFormatter.SetVeryShortStandaloneWeekdaySymbols]
//
// # Managing Month Symbols
//
//   - [DateFormatter.MonthSymbols]: The month symbols for the receiver.
//   - [DateFormatter.SetMonthSymbols]
//   - [DateFormatter.ShortMonthSymbols]: The array of short month symbols for the receiver.
//   - [DateFormatter.SetShortMonthSymbols]
//   - [DateFormatter.VeryShortMonthSymbols]: The very short month symbols for the receiver.
//   - [DateFormatter.SetVeryShortMonthSymbols]
//   - [DateFormatter.StandaloneMonthSymbols]: The standalone month symbols for the receiver.
//   - [DateFormatter.SetStandaloneMonthSymbols]
//   - [DateFormatter.ShortStandaloneMonthSymbols]: The short standalone month symbols for the receiver.
//   - [DateFormatter.SetShortStandaloneMonthSymbols]
//   - [DateFormatter.VeryShortStandaloneMonthSymbols]: The very short month symbols for the receiver.
//   - [DateFormatter.SetVeryShortStandaloneMonthSymbols]
//
// # Managing Quarter Symbols
//
//   - [DateFormatter.QuarterSymbols]: The quarter symbols for the receiver.
//   - [DateFormatter.SetQuarterSymbols]
//   - [DateFormatter.ShortQuarterSymbols]: The short quarter symbols for the receiver.
//   - [DateFormatter.SetShortQuarterSymbols]
//   - [DateFormatter.StandaloneQuarterSymbols]: The standalone quarter symbols for the receiver.
//   - [DateFormatter.SetStandaloneQuarterSymbols]
//   - [DateFormatter.ShortStandaloneQuarterSymbols]: The short standalone quarter symbols for the receiver.
//   - [DateFormatter.SetShortStandaloneQuarterSymbols]
//
// # Managing Era Symbols
//
//   - [DateFormatter.EraSymbols]: The era symbols for the receiver.
//   - [DateFormatter.SetEraSymbols]
//   - [DateFormatter.LongEraSymbols]: The long era symbols for the receiver
//   - [DateFormatter.SetLongEraSymbols]
//
// # Deprecated
//
//   - [DateFormatter.GeneratesCalendarDates]: Indicates whether the formatter generates the deprecated calendar date type.
//   - [DateFormatter.SetGeneratesCalendarDates]
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter
type DateFormatter struct {
	NSFormatter
}

// DateFormatterFromID constructs a [DateFormatter] from an objc.ID.
//
// A formatter that converts between dates and their textual representations.
func DateFormatterFromID(id objc.ID) DateFormatter {
	return NSDateFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSDateFormatterFromID is an alias for [DateFormatterFromID] for cross-framework compatibility.
func NSDateFormatterFromID(id objc.ID) DateFormatter { return DateFormatterFromID(id) }
// NOTE: DateFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [DateFormatter] class.
//
// # Converting Objects
//
//   - [IDateFormatter.DateFromString]: Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
//   - [IDateFormatter.StringFromDate]: Returns a string representation of a specified date that the system formats using the receiver’s current settings.
//   - [IDateFormatter.GetObjectValueForStringRangeError]: Returns by reference a date representation of a specified string and its date range, as well as a Boolean value that indicates whether the system can parse the string.
//
// # Managing Formats and Styles
//
//   - [IDateFormatter.DateStyle]: The date style of the receiver.
//   - [IDateFormatter.SetDateStyle]
//   - [IDateFormatter.TimeStyle]: The time style of the receiver.
//   - [IDateFormatter.SetTimeStyle]
//   - [IDateFormatter.DateFormat]: The date format string used by the receiver.
//   - [IDateFormatter.SetDateFormat]
//   - [IDateFormatter.SetLocalizedDateFormatFromTemplate]: Sets the date format from a template using the specified locale for the receiver.
//   - [IDateFormatter.FormattingContext]: The capitalization formatting context used when formatting a date.
//   - [IDateFormatter.SetFormattingContext]
//
// # Managing Attributes
//
//   - [IDateFormatter.Calendar]: The calendar for the receiver.
//   - [IDateFormatter.SetCalendar]
//   - [IDateFormatter.DefaultDate]: The default date for the receiver.
//   - [IDateFormatter.SetDefaultDate]
//   - [IDateFormatter.Locale]: The locale for the receiver.
//   - [IDateFormatter.SetLocale]
//   - [IDateFormatter.TimeZone]: The time zone for the receiver.
//   - [IDateFormatter.SetTimeZone]
//   - [IDateFormatter.TwoDigitStartDate]: The earliest date that can be denoted by a two-digit year specifier.
//   - [IDateFormatter.SetTwoDigitStartDate]
//   - [IDateFormatter.GregorianStartDate]: The start date of the Gregorian calendar for the receiver.
//   - [IDateFormatter.SetGregorianStartDate]
//
// # Managing Behavior Version
//
//   - [IDateFormatter.FormatterBehavior]: The formatter behavior for the receiver.
//   - [IDateFormatter.SetFormatterBehavior]
//
// # Managing Natural Language Support
//
//   - [IDateFormatter.Lenient]: A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
//   - [IDateFormatter.SetLenient]
//   - [IDateFormatter.DoesRelativeDateFormatting]: A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
//   - [IDateFormatter.SetDoesRelativeDateFormatting]
//
// # Managing AM and PM Symbols
//
//   - [IDateFormatter.AMSymbol]: The AM symbol for the receiver.
//   - [IDateFormatter.SetAMSymbol]
//   - [IDateFormatter.PMSymbol]: The PM symbol for the receiver.
//   - [IDateFormatter.SetPMSymbol]
//
// # Managing Weekday Symbols
//
//   - [IDateFormatter.WeekdaySymbols]: The array of weekday symbols for the receiver.
//   - [IDateFormatter.SetWeekdaySymbols]
//   - [IDateFormatter.ShortWeekdaySymbols]: The array of short weekday symbols for the receiver.
//   - [IDateFormatter.SetShortWeekdaySymbols]
//   - [IDateFormatter.VeryShortWeekdaySymbols]: The array of very short weekday symbols for the receiver.
//   - [IDateFormatter.SetVeryShortWeekdaySymbols]
//   - [IDateFormatter.StandaloneWeekdaySymbols]: The array of standalone weekday symbols for the receiver.
//   - [IDateFormatter.SetStandaloneWeekdaySymbols]
//   - [IDateFormatter.ShortStandaloneWeekdaySymbols]: The array of short standalone weekday symbols for the receiver.
//   - [IDateFormatter.SetShortStandaloneWeekdaySymbols]
//   - [IDateFormatter.VeryShortStandaloneWeekdaySymbols]: The array of very short standalone weekday symbols for the receiver.
//   - [IDateFormatter.SetVeryShortStandaloneWeekdaySymbols]
//
// # Managing Month Symbols
//
//   - [IDateFormatter.MonthSymbols]: The month symbols for the receiver.
//   - [IDateFormatter.SetMonthSymbols]
//   - [IDateFormatter.ShortMonthSymbols]: The array of short month symbols for the receiver.
//   - [IDateFormatter.SetShortMonthSymbols]
//   - [IDateFormatter.VeryShortMonthSymbols]: The very short month symbols for the receiver.
//   - [IDateFormatter.SetVeryShortMonthSymbols]
//   - [IDateFormatter.StandaloneMonthSymbols]: The standalone month symbols for the receiver.
//   - [IDateFormatter.SetStandaloneMonthSymbols]
//   - [IDateFormatter.ShortStandaloneMonthSymbols]: The short standalone month symbols for the receiver.
//   - [IDateFormatter.SetShortStandaloneMonthSymbols]
//   - [IDateFormatter.VeryShortStandaloneMonthSymbols]: The very short month symbols for the receiver.
//   - [IDateFormatter.SetVeryShortStandaloneMonthSymbols]
//
// # Managing Quarter Symbols
//
//   - [IDateFormatter.QuarterSymbols]: The quarter symbols for the receiver.
//   - [IDateFormatter.SetQuarterSymbols]
//   - [IDateFormatter.ShortQuarterSymbols]: The short quarter symbols for the receiver.
//   - [IDateFormatter.SetShortQuarterSymbols]
//   - [IDateFormatter.StandaloneQuarterSymbols]: The standalone quarter symbols for the receiver.
//   - [IDateFormatter.SetStandaloneQuarterSymbols]
//   - [IDateFormatter.ShortStandaloneQuarterSymbols]: The short standalone quarter symbols for the receiver.
//   - [IDateFormatter.SetShortStandaloneQuarterSymbols]
//
// # Managing Era Symbols
//
//   - [IDateFormatter.EraSymbols]: The era symbols for the receiver.
//   - [IDateFormatter.SetEraSymbols]
//   - [IDateFormatter.LongEraSymbols]: The long era symbols for the receiver
//   - [IDateFormatter.SetLongEraSymbols]
//
// # Deprecated
//
//   - [IDateFormatter.GeneratesCalendarDates]: Indicates whether the formatter generates the deprecated calendar date type.
//   - [IDateFormatter.SetGeneratesCalendarDates]
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter
type IDateFormatter interface {
	INSFormatter

	// Topic: Converting Objects

	// Returns a date representation of a specified string that the system interprets using the receiver’s current settings.
	DateFromString(string_ string) INSDate
	// Returns a string representation of a specified date that the system formats using the receiver’s current settings.
	StringFromDate(date INSDate) string
	// Returns by reference a date representation of a specified string and its date range, as well as a Boolean value that indicates whether the system can parse the string.
	GetObjectValueForStringRangeError(obj []objectivec.IObject, string_ string, rangep NSRange) (bool, error)

	// Topic: Managing Formats and Styles

	// The date style of the receiver.
	DateStyle() NSDateFormatterStyle
	SetDateStyle(value NSDateFormatterStyle)
	// The time style of the receiver.
	TimeStyle() NSDateFormatterStyle
	SetTimeStyle(value NSDateFormatterStyle)
	// The date format string used by the receiver.
	DateFormat() string
	SetDateFormat(value string)
	// Sets the date format from a template using the specified locale for the receiver.
	SetLocalizedDateFormatFromTemplate(dateFormatTemplate string)
	// The capitalization formatting context used when formatting a date.
	FormattingContext() NSFormattingContext
	SetFormattingContext(value NSFormattingContext)

	// Topic: Managing Attributes

	// The calendar for the receiver.
	Calendar() INSCalendar
	SetCalendar(value INSCalendar)
	// The default date for the receiver.
	DefaultDate() INSDate
	SetDefaultDate(value INSDate)
	// The locale for the receiver.
	Locale() INSLocale
	SetLocale(value INSLocale)
	// The time zone for the receiver.
	TimeZone() INSTimeZone
	SetTimeZone(value INSTimeZone)
	// The earliest date that can be denoted by a two-digit year specifier.
	TwoDigitStartDate() INSDate
	SetTwoDigitStartDate(value INSDate)
	// The start date of the Gregorian calendar for the receiver.
	GregorianStartDate() INSDate
	SetGregorianStartDate(value INSDate)

	// Topic: Managing Behavior Version

	// The formatter behavior for the receiver.
	FormatterBehavior() NSDateFormatterBehavior
	SetFormatterBehavior(value NSDateFormatterBehavior)

	// Topic: Managing Natural Language Support

	// A Boolean value that indicates whether the receiver uses heuristics when parsing a string.
	Lenient() bool
	SetLenient(value bool)
	// A Boolean value that indicates whether the receiver uses phrases such as “today” and “tomorrow” for the date component.
	DoesRelativeDateFormatting() bool
	SetDoesRelativeDateFormatting(value bool)

	// Topic: Managing AM and PM Symbols

	// The AM symbol for the receiver.
	AMSymbol() string
	SetAMSymbol(value string)
	// The PM symbol for the receiver.
	PMSymbol() string
	SetPMSymbol(value string)

	// Topic: Managing Weekday Symbols

	// The array of weekday symbols for the receiver.
	WeekdaySymbols() []string
	SetWeekdaySymbols(value []string)
	// The array of short weekday symbols for the receiver.
	ShortWeekdaySymbols() []string
	SetShortWeekdaySymbols(value []string)
	// The array of very short weekday symbols for the receiver.
	VeryShortWeekdaySymbols() []string
	SetVeryShortWeekdaySymbols(value []string)
	// The array of standalone weekday symbols for the receiver.
	StandaloneWeekdaySymbols() []string
	SetStandaloneWeekdaySymbols(value []string)
	// The array of short standalone weekday symbols for the receiver.
	ShortStandaloneWeekdaySymbols() []string
	SetShortStandaloneWeekdaySymbols(value []string)
	// The array of very short standalone weekday symbols for the receiver.
	VeryShortStandaloneWeekdaySymbols() []string
	SetVeryShortStandaloneWeekdaySymbols(value []string)

	// Topic: Managing Month Symbols

	// The month symbols for the receiver.
	MonthSymbols() []string
	SetMonthSymbols(value []string)
	// The array of short month symbols for the receiver.
	ShortMonthSymbols() []string
	SetShortMonthSymbols(value []string)
	// The very short month symbols for the receiver.
	VeryShortMonthSymbols() []string
	SetVeryShortMonthSymbols(value []string)
	// The standalone month symbols for the receiver.
	StandaloneMonthSymbols() []string
	SetStandaloneMonthSymbols(value []string)
	// The short standalone month symbols for the receiver.
	ShortStandaloneMonthSymbols() []string
	SetShortStandaloneMonthSymbols(value []string)
	// The very short month symbols for the receiver.
	VeryShortStandaloneMonthSymbols() []string
	SetVeryShortStandaloneMonthSymbols(value []string)

	// Topic: Managing Quarter Symbols

	// The quarter symbols for the receiver.
	QuarterSymbols() []string
	SetQuarterSymbols(value []string)
	// The short quarter symbols for the receiver.
	ShortQuarterSymbols() []string
	SetShortQuarterSymbols(value []string)
	// The standalone quarter symbols for the receiver.
	StandaloneQuarterSymbols() []string
	SetStandaloneQuarterSymbols(value []string)
	// The short standalone quarter symbols for the receiver.
	ShortStandaloneQuarterSymbols() []string
	SetShortStandaloneQuarterSymbols(value []string)

	// Topic: Managing Era Symbols

	// The era symbols for the receiver.
	EraSymbols() []string
	SetEraSymbols(value []string)
	// The long era symbols for the receiver
	LongEraSymbols() []string
	SetLongEraSymbols(value []string)

	// Topic: Deprecated

	// Indicates whether the formatter generates the deprecated calendar date type.
	GeneratesCalendarDates() bool
	SetGeneratesCalendarDates(value bool)
}

// Init initializes the instance.
func (d DateFormatter) Init() DateFormatter {
	rv := objc.Send[DateFormatter](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DateFormatter) Autorelease() DateFormatter {
	rv := objc.Send[DateFormatter](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateFormatter creates a new DateFormatter instance.
func NewDateFormatter() DateFormatter {
	class := getDateFormatterClass()
	rv := objc.Send[DateFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDateFormatterWithCoder(coder INSCoder) DateFormatter {
	instance := getDateFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DateFormatterFromID(rv)
}

// Returns a date representation of a specified string that the system
// interprets using the receiver’s current settings.
//
// string: The string to parse.
//
// # Return Value
// 
// A date representation of `string`. If [DateFromString] can’t parse the
// string, returns `nil`.
//
// # Discussion
// 
// For more information about using [NSDateFormatter] to convert a string to a
// date, see [NSDateFormatter]. For a sample code playground, see [Displaying
// Human-Friendly Content].
//
// [Displaying Human-Friendly Content]: https://developer.apple.com/documentation/Foundation/displaying-human-friendly-content
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/date(from:)
func (d DateFormatter) DateFromString(string_ string) INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dateFromString:"), objc.String(string_))
	return NSDateFromID(rv)
}
// Returns a string representation of a specified date that the system formats
// using the receiver’s current settings.
//
// date: The date to format.
//
// # Return Value
// 
// A string representation of `date`.
//
// # Discussion
// 
// For more information about using [NSDateFormatter] to produce a string
// representation of a date, see [NSDateFormatter]. For a sample code
// playground, see [Displaying Human-Friendly Content].
//
// [Displaying Human-Friendly Content]: https://developer.apple.com/documentation/Foundation/displaying-human-friendly-content
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/string(from:)
func (d DateFormatter) StringFromDate(date INSDate) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stringFromDate:"), date)
	return NSStringFromID(rv).String()
}
// Returns by reference a date representation of a specified string and its
// date range, as well as a Boolean value that indicates whether the system
// can parse the string.
//
// obj: If the receiver is able to parse `string`, upon return contains a date
// representation of `string`.
//
// string: The string to parse.
//
// rangep: If the receiver is able to parse `string`, upon return contains the range
// of `string` used to create the date.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/getObjectValue(_:for:range:)
func (d DateFormatter) GetObjectValueForStringRangeError(obj []objectivec.IObject, string_ string, rangep NSRange) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("getObjectValue:forString:range:error:"), objectivec.IObjectSliceToNSArray(obj), objc.String(string_), rangep, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("getObjectValue:forString:range:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the date format from a template using the specified locale for the
// receiver.
//
// dateFormatTemplate: A string containing date format patterns (such as “MM” or “h”).
// 
// For full details, see [Date and Time Programming Guide].
// //
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
//
// # Discussion
// 
// See [Data Formatting Guide] for a list of the conversion specifiers
// permitted in date format strings.
// 
// Calling this method is equivalent to, but not necessarily implemented as,
// setting the [DateFormat] property to the result of calling the
// [DateFormatFromTemplateOptionsLocale] method, passing no options and the
// [Locale] property value.
//
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/setLocalizedDateFormatFromTemplate(_:)
func (d DateFormatter) SetLocalizedDateFormatFromTemplate(dateFormatTemplate string) {
	objc.Send[objc.ID](d.ID, objc.Sel("setLocalizedDateFormatFromTemplate:"), objc.String(dateFormatTemplate))
}

// Returns a string representation of a specified date, that the system
// formats for the current locale using the specified date and time styles.
//
// date: A date.
//
// dstyle: A format style for the date. For possible values, see
// [DateFormatter.Style].
// //
// [DateFormatter.Style]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style
//
// tstyle: A format style for the time. For possible values, see
// [DateFormatter.Style].
// //
// [DateFormatter.Style]: https://developer.apple.com/documentation/Foundation/DateFormatter/Style
//
// # Return Value
// 
// A localized string representation of `date` using the specified date and
// time styles.
//
// # Discussion
// 
// This method uses a date formatter configured with the current default
// settings. The returned string is the same as if you configured and used a
// date formatter as shown in the following example:
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/localizedString(from:dateStyle:timeStyle:)
func (_DateFormatterClass DateFormatterClass) LocalizedStringFromDateDateStyleTimeStyle(date INSDate, dstyle NSDateFormatterStyle, tstyle NSDateFormatterStyle) string {
	rv := objc.Send[objc.ID](objc.ID(_DateFormatterClass.class), objc.Sel("localizedStringFromDate:dateStyle:timeStyle:"), date, dstyle, tstyle)
	return NSStringFromID(rv).String()
}
// Returns a localized date format string representing the given date format
// components arranged appropriately for the specified locale.
//
// tmplate: A string containing date format patterns (such as “MM” or “h”).
// 
// For full details, see [Date and Time Programming Guide].
// //
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
//
// opts: No options are currently defined.
//
// locale: The locale for which the template is required.
//
// # Return Value
// 
// A localized date format string representing the date format components
// given in `template`, arranged appropriately for the locale specified by
// `locale`. The returned string may not contain exactly those components
// given in `template`, but may—for example—have locale-specific
// adjustments applied.
//
// # Discussion
// 
// Different locales have different conventions for the ordering of date
// components. You use this method to get an appropriate format string for a
// given set of components for a specified locale (typically you use the
// current locale—see [CurrentLocale]).
// 
// The following example shows the difference between the date formats for
// British and American English:
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat(fromTemplate:options:locale:)
func (_DateFormatterClass DateFormatterClass) DateFormatFromTemplateOptionsLocale(tmplate string, opts uint, locale INSLocale) string {
	rv := objc.Send[objc.ID](objc.ID(_DateFormatterClass.class), objc.Sel("dateFormatFromTemplate:options:locale:"), objc.String(tmplate), opts, locale)
	return NSStringFromID(rv).String()
}

// The date style of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/dateStyle
func (d DateFormatter) DateStyle() NSDateFormatterStyle {
	rv := objc.Send[NSDateFormatterStyle](d.ID, objc.Sel("dateStyle"))
	return NSDateFormatterStyle(rv)
}
func (d DateFormatter) SetDateStyle(value NSDateFormatterStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setDateStyle:"), value)
}
// The time style of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/timeStyle
func (d DateFormatter) TimeStyle() NSDateFormatterStyle {
	rv := objc.Send[NSDateFormatterStyle](d.ID, objc.Sel("timeStyle"))
	return NSDateFormatterStyle(rv)
}
func (d DateFormatter) SetTimeStyle(value NSDateFormatterStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeStyle:"), value)
}
// The date format string used by the receiver.
//
// # Discussion
// 
// See [Data Formatting Guide] for a list of the conversion specifiers
// permitted in date format strings.
// 
// You should only set this property when working with fixed format
// representations, as discussed in [NSDateFormatter]. For user-visible
// representations, you should use the [DateStyle] and [TimeStyle] properties,
// or the [SetLocalizedDateFormatFromTemplate] method if your desired format
// cannot be achieved using the predefined styles; both of these properties
// and this method provide a localized date representation appropriate for
// display to the user.
//
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/dateFormat
func (d DateFormatter) DateFormat() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dateFormat"))
	return NSStringFromID(rv).String()
}
func (d DateFormatter) SetDateFormat(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setDateFormat:"), objc.String(value))
}
// The capitalization formatting context used when formatting a date.
//
// # Discussion
// 
// The formatting context allows the formatter to apply appropriate
// capitalization depending on how the how the string will be used, and
// whether the locale makes capitalization distinctions.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/formattingContext
func (d DateFormatter) FormattingContext() NSFormattingContext {
	rv := objc.Send[NSFormattingContext](d.ID, objc.Sel("formattingContext"))
	return NSFormattingContext(rv)
}
func (d DateFormatter) SetFormattingContext(value NSFormattingContext) {
	objc.Send[struct{}](d.ID, objc.Sel("setFormattingContext:"), value)
}
// The calendar for the receiver.
//
// # Discussion
// 
// If unspecified, the logical calendar for the current user is used.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/calendar
func (d DateFormatter) Calendar() INSCalendar {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("calendar"))
	return NSCalendarFromID(objc.ID(rv))
}
func (d DateFormatter) SetCalendar(value INSCalendar) {
	objc.Send[struct{}](d.ID, objc.Sel("setCalendar:"), value)
}
// The default date for the receiver.
//
// # Discussion
// 
// By default, this property is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultDate
func (d DateFormatter) DefaultDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("defaultDate"))
	return NSDateFromID(objc.ID(rv))
}
func (d DateFormatter) SetDefaultDate(value INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDefaultDate:"), value)
}
// The locale for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/locale
func (d DateFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (d DateFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocale:"), value)
}
// The time zone for the receiver.
//
// # Discussion
// 
// If unspecified, the system time zone is used.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/timeZone
func (d DateFormatter) TimeZone() INSTimeZone {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
func (d DateFormatter) SetTimeZone(value INSTimeZone) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeZone:"), value)
}
// The earliest date that can be denoted by a two-digit year specifier.
//
// # Discussion
// 
// If the two-digit start date is set to January 6, 1976, then “January 1,
// 76” is interpreted as New Year’s Day in 2076, whereas “February 14,
// 76” is interpreted as Valentine’s Day in 1976.
// 
// By default, this property is equal to December 31, 1949.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/twoDigitStartDate
func (d DateFormatter) TwoDigitStartDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("twoDigitStartDate"))
	return NSDateFromID(objc.ID(rv))
}
func (d DateFormatter) SetTwoDigitStartDate(value INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setTwoDigitStartDate:"), value)
}
// The start date of the Gregorian calendar for the receiver.
//
// # Discussion
// 
// This is used to specify the start date for the Gregorian calendar switch
// from the Julian calendar. Different locales switched at different times.
// Normally you should just accept the locale’s default date for the switch.
// 
// See [NSCalendar] for more information.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/gregorianStartDate
func (d DateFormatter) GregorianStartDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("gregorianStartDate"))
	return NSDateFromID(objc.ID(rv))
}
func (d DateFormatter) SetGregorianStartDate(value INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setGregorianStartDate:"), value)
}
// The formatter behavior for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/formatterBehavior
func (d DateFormatter) FormatterBehavior() NSDateFormatterBehavior {
	rv := objc.Send[NSDateFormatterBehavior](d.ID, objc.Sel("formatterBehavior"))
	return NSDateFormatterBehavior(rv)
}
func (d DateFormatter) SetFormatterBehavior(value NSDateFormatterBehavior) {
	objc.Send[struct{}](d.ID, objc.Sel("setFormatterBehavior:"), value)
}
// A Boolean value that indicates whether the receiver uses heuristics when
// parsing a string.
//
// # Discussion
// 
// [true] if the receiver has been set to use heuristics when parsing a string
// to guess at the date which is intended, otherwise [false].
// 
// If a formatter is set to be lenient, when parsing a string it uses
// heuristics to guess at the date which is intended. As with any guessing, it
// may get the result date wrong (that is, a date other than that which was
// intended).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/isLenient
func (d DateFormatter) Lenient() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isLenient"))
	return rv
}
func (d DateFormatter) SetLenient(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setLenient:"), value)
}
// A Boolean value that indicates whether the receiver uses phrases such as
// “today” and “tomorrow” for the date component.
//
// # Discussion
// 
// [true] if the receiver uses relative date formatting, otherwise [false].
// 
// If a date formatter uses relative date formatting, where possible it
// replaces the date component of its output with a phrase—such as
// “today” or “tomorrow”—that indicates a relative date. The
// available phrases depend on the locale for the date formatter; whereas, for
// dates in the future, English may only allow “tomorrow,” French may
// allow “the day after the day after tomorrow,” as illustrated in the
// following example.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/doesRelativeDateFormatting
func (d DateFormatter) DoesRelativeDateFormatting() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("doesRelativeDateFormatting"))
	return rv
}
func (d DateFormatter) SetDoesRelativeDateFormatting(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDoesRelativeDateFormatting:"), value)
}
// The AM symbol for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/amSymbol
func (d DateFormatter) AMSymbol() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("AMSymbol"))
	return NSStringFromID(rv).String()
}
func (d DateFormatter) SetAMSymbol(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setAMSymbol:"), objc.String(value))
}
// The PM symbol for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/pmSymbol
func (d DateFormatter) PMSymbol() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("PMSymbol"))
	return NSStringFromID(rv).String()
}
func (d DateFormatter) SetPMSymbol(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setPMSymbol:"), objc.String(value))
}
// The array of weekday symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/weekdaySymbols
func (d DateFormatter) WeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("weekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetWeekdaySymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setWeekdaySymbols:"), objectivec.StringSliceToNSArray(value))
}
// The array of short weekday symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/shortWeekdaySymbols
func (d DateFormatter) ShortWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("shortWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetShortWeekdaySymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setShortWeekdaySymbols:"), objectivec.StringSliceToNSArray(value))
}
// The array of very short weekday symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortWeekdaySymbols
func (d DateFormatter) VeryShortWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("veryShortWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetVeryShortWeekdaySymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setVeryShortWeekdaySymbols:"), objectivec.StringSliceToNSArray(value))
}
// The array of standalone weekday symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneWeekdaySymbols
func (d DateFormatter) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("standaloneWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetStandaloneWeekdaySymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setStandaloneWeekdaySymbols:"), objectivec.StringSliceToNSArray(value))
}
// The array of short standalone weekday symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneWeekdaySymbols
func (d DateFormatter) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetShortStandaloneWeekdaySymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setShortStandaloneWeekdaySymbols:"), objectivec.StringSliceToNSArray(value))
}
// The array of very short standalone weekday symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneWeekdaySymbols
func (d DateFormatter) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetVeryShortStandaloneWeekdaySymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setVeryShortStandaloneWeekdaySymbols:"), objectivec.StringSliceToNSArray(value))
}
// The month symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/monthSymbols
func (d DateFormatter) MonthSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("monthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetMonthSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setMonthSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The array of short month symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/shortMonthSymbols
func (d DateFormatter) ShortMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("shortMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetShortMonthSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setShortMonthSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The very short month symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortMonthSymbols
func (d DateFormatter) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("veryShortMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetVeryShortMonthSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setVeryShortMonthSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The standalone month symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneMonthSymbols
func (d DateFormatter) StandaloneMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("standaloneMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetStandaloneMonthSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setStandaloneMonthSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The short standalone month symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneMonthSymbols
func (d DateFormatter) ShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetShortStandaloneMonthSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setShortStandaloneMonthSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The very short month symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/veryShortStandaloneMonthSymbols
func (d DateFormatter) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetVeryShortStandaloneMonthSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setVeryShortStandaloneMonthSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The quarter symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/quarterSymbols
func (d DateFormatter) QuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("quarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetQuarterSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setQuarterSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The short quarter symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/shortQuarterSymbols
func (d DateFormatter) ShortQuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("shortQuarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetShortQuarterSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setShortQuarterSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The standalone quarter symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/standaloneQuarterSymbols
func (d DateFormatter) StandaloneQuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("standaloneQuarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetStandaloneQuarterSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setStandaloneQuarterSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The short standalone quarter symbols for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/shortStandaloneQuarterSymbols
func (d DateFormatter) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetShortStandaloneQuarterSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setShortStandaloneQuarterSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The era symbols for the receiver.
//
// # Discussion
// 
// An array containing [NSString] objects representing the era symbols for the
// receiver (for example, {“B.C.E.”, “C.E.”}).
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/eraSymbols
func (d DateFormatter) EraSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("eraSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetEraSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setEraSymbols:"), objectivec.StringSliceToNSArray(value))
}
// The long era symbols for the receiver
//
// # Discussion
// 
// An array containing [NSString] objects representing the era symbols for the
// receiver (for example, {“Before Common Era”, “Common Era”}).
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/longEraSymbols
func (d DateFormatter) LongEraSymbols() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("longEraSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
func (d DateFormatter) SetLongEraSymbols(value []string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLongEraSymbols:"), objectivec.StringSliceToNSArray(value))
}
// Indicates whether the formatter generates the deprecated calendar date
// type.
//
// # Discussion
// 
// This property is [true] if the formatter generates the deprecated
// [NSCalendarDate] type, and is [false] otherwise. You should use [Date] and
// [Calendar] rather than [NSCalendarDate].
//
// [Calendar]: https://developer.apple.com/documentation/Foundation/Calendar
// [Date]: https://developer.apple.com/documentation/Foundation/Date
// [NSCalendarDate]: https://developer.apple.com/documentation/Foundation/NSCalendarDate
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/generatesCalendarDates
func (d DateFormatter) GeneratesCalendarDates() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("generatesCalendarDates"))
	return rv
}
func (d DateFormatter) SetGeneratesCalendarDates(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setGeneratesCalendarDates:"), value)
}

// Returns the default formatting behavior for instances of the class.
//
// # Return Value
// 
// The default formatting behavior for instances of the class. For possible
// values, see [DateFormatter.Behavior].
// 
// # Discussion
// 
// For iOS and for macOS applications linked against macOS 10.5 and later, the
// default is `NSDateFormatterBehavior10_4`.
//
// [DateFormatter.Behavior]: https://developer.apple.com/documentation/Foundation/DateFormatter/Behavior
//
// See: https://developer.apple.com/documentation/Foundation/DateFormatter/defaultFormatterBehavior
func (_DateFormatterClass DateFormatterClass) DefaultFormatterBehavior() NSDateFormatterBehavior {
	rv := objc.Send[NSDateFormatterBehavior](objc.ID(_DateFormatterClass.class), objc.Sel("defaultFormatterBehavior"))
	return NSDateFormatterBehavior(rv)
}
func (_DateFormatterClass DateFormatterClass) SetDefaultFormatterBehavior(value NSDateFormatterBehavior) {
	objc.Send[struct{}](objc.ID(_DateFormatterClass.class), objc.Sel("setDefaultFormatterBehavior:"), value)
}

