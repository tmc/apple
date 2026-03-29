// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [RelativeDateTimeFormatter] class.
var (
	_RelativeDateTimeFormatterClass     RelativeDateTimeFormatterClass
	_RelativeDateTimeFormatterClassOnce sync.Once
)

func getRelativeDateTimeFormatterClass() RelativeDateTimeFormatterClass {
	_RelativeDateTimeFormatterClassOnce.Do(func() {
		_RelativeDateTimeFormatterClass = RelativeDateTimeFormatterClass{class: objc.GetClass("NSRelativeDateTimeFormatter")}
	})
	return _RelativeDateTimeFormatterClass
}

// GetRelativeDateTimeFormatterClass returns the class object for NSRelativeDateTimeFormatter.
func GetRelativeDateTimeFormatterClass() RelativeDateTimeFormatterClass {
	return getRelativeDateTimeFormatterClass()
}

type RelativeDateTimeFormatterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (rc RelativeDateTimeFormatterClass) Class() objc.Class {
	return rc.class
}

// Alloc allocates memory for a new instance of the class.
func (rc RelativeDateTimeFormatterClass) Alloc() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that creates locale-aware string representations of a relative
// date or time.
//
// # Overview
// 
// Use the strings that the formatter produces, such as “1 hour ago”,
// “in 2 weeks”, “yesterday”, and “tomorrow” as standalone
// strings. Embedding them in other strings may not be grammatically correct.
//
// # Converting Dates to Formatted Strings
//
//   - [RelativeDateTimeFormatter.LocalizedStringForDateRelativeToDate]: Formats the date interval from the reference date to the specified date using the formatter’s calendar.
//   - [RelativeDateTimeFormatter.LocalizedStringFromDateComponents]: Formats a relative time represented by the specified date components.
//   - [RelativeDateTimeFormatter.LocalizedStringFromTimeInterval]: Formats the specified time interval using the formatter’s calendar.
//
// # Configuring Formatter Options
//
//   - [RelativeDateTimeFormatter.Calendar]: The calendar to use for formatting values that don’t have an inherent calendar of their own.
//   - [RelativeDateTimeFormatter.SetCalendar]
//   - [RelativeDateTimeFormatter.Locale]: The locale to use when formatting the date.
//   - [RelativeDateTimeFormatter.SetLocale]
//   - [RelativeDateTimeFormatter.DateTimeStyle]: The style to use when describing a relative date, for example “yesterday” or “1 day ago”.
//   - [RelativeDateTimeFormatter.SetDateTimeStyle]
//   - [RelativeDateTimeFormatter.UnitsStyle]: The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.
//   - [RelativeDateTimeFormatter.SetUnitsStyle]
//   - [RelativeDateTimeFormatter.FormattingContext]: A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.
//   - [RelativeDateTimeFormatter.SetFormattingContext]
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter
type RelativeDateTimeFormatter struct {
	NSFormatter
}

// RelativeDateTimeFormatterFromID constructs a [RelativeDateTimeFormatter] from an objc.ID.
//
// A formatter that creates locale-aware string representations of a relative
// date or time.
func RelativeDateTimeFormatterFromID(id objc.ID) RelativeDateTimeFormatter {
	return NSRelativeDateTimeFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSRelativeDateTimeFormatterFromID is an alias for [RelativeDateTimeFormatterFromID] for cross-framework compatibility.
func NSRelativeDateTimeFormatterFromID(id objc.ID) RelativeDateTimeFormatter { return RelativeDateTimeFormatterFromID(id) }
// NOTE: RelativeDateTimeFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [RelativeDateTimeFormatter] class.
//
// # Converting Dates to Formatted Strings
//
//   - [IRelativeDateTimeFormatter.LocalizedStringForDateRelativeToDate]: Formats the date interval from the reference date to the specified date using the formatter’s calendar.
//   - [IRelativeDateTimeFormatter.LocalizedStringFromDateComponents]: Formats a relative time represented by the specified date components.
//   - [IRelativeDateTimeFormatter.LocalizedStringFromTimeInterval]: Formats the specified time interval using the formatter’s calendar.
//
// # Configuring Formatter Options
//
//   - [IRelativeDateTimeFormatter.Calendar]: The calendar to use for formatting values that don’t have an inherent calendar of their own.
//   - [IRelativeDateTimeFormatter.SetCalendar]
//   - [IRelativeDateTimeFormatter.Locale]: The locale to use when formatting the date.
//   - [IRelativeDateTimeFormatter.SetLocale]
//   - [IRelativeDateTimeFormatter.DateTimeStyle]: The style to use when describing a relative date, for example “yesterday” or “1 day ago”.
//   - [IRelativeDateTimeFormatter.SetDateTimeStyle]
//   - [IRelativeDateTimeFormatter.UnitsStyle]: The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.
//   - [IRelativeDateTimeFormatter.SetUnitsStyle]
//   - [IRelativeDateTimeFormatter.FormattingContext]: A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.
//   - [IRelativeDateTimeFormatter.SetFormattingContext]
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter
type IRelativeDateTimeFormatter interface {
	INSFormatter

	// Topic: Converting Dates to Formatted Strings

	// Formats the date interval from the reference date to the specified date using the formatter’s calendar.
	LocalizedStringForDateRelativeToDate(date INSDate, referenceDate INSDate) string
	// Formats a relative time represented by the specified date components.
	LocalizedStringFromDateComponents(dateComponents INSDateComponents) string
	// Formats the specified time interval using the formatter’s calendar.
	LocalizedStringFromTimeInterval(timeInterval float64) string

	// Topic: Configuring Formatter Options

	// The calendar to use for formatting values that don’t have an inherent calendar of their own.
	Calendar() INSCalendar
	SetCalendar(value INSCalendar)
	// The locale to use when formatting the date.
	Locale() INSLocale
	SetLocale(value INSLocale)
	// The style to use when describing a relative date, for example “yesterday” or “1 day ago”.
	DateTimeStyle() NSRelativeDateTimeFormatterStyle
	SetDateTimeStyle(value NSRelativeDateTimeFormatterStyle)
	// The style to use when formatting the quantity or the name of the unit, such as “1 day ago” or “one day ago”.
	UnitsStyle() NSRelativeDateTimeFormatterUnitsStyle
	SetUnitsStyle(value NSRelativeDateTimeFormatterUnitsStyle)
	// A description of where the formatted string will appear, allowing the formatter to capitalize the output appropriately.
	FormattingContext() NSFormattingContext
	SetFormattingContext(value NSFormattingContext)
}

// Init initializes the instance.
func (r RelativeDateTimeFormatter) Init() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RelativeDateTimeFormatter) Autorelease() RelativeDateTimeFormatter {
	rv := objc.Send[RelativeDateTimeFormatter](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRelativeDateTimeFormatter creates a new RelativeDateTimeFormatter instance.
func NewRelativeDateTimeFormatter() RelativeDateTimeFormatter {
	class := getRelativeDateTimeFormatterClass()
	rv := objc.Send[RelativeDateTimeFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewRelativeDateTimeFormatterWithCoder(coder INSCoder) RelativeDateTimeFormatter {
	instance := getRelativeDateTimeFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return RelativeDateTimeFormatterFromID(rv)
}

// Formats the date interval from the reference date to the specified date
// using the formatter’s calendar.
//
// date: The end date of the interval to format.
//
// referenceDate: The start date of the interval to format.
//
// # Return Value
// 
// A string that represents the date interval between two dates.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/localizedString(for:relativeTo:)
func (r RelativeDateTimeFormatter) LocalizedStringForDateRelativeToDate(date INSDate, referenceDate INSDate) string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("localizedStringForDate:relativeToDate:"), date, referenceDate)
	return NSStringFromID(rv).String()
}
// Formats a relative time represented by the specified date components.
//
// dateComponents: The date components to format.
//
// # Return Value
// 
// A string that represents the formatted relative time from date components.
//
// # Discussion
// 
// The formatter interprets a negative component value as a date in the past.
// 
// This method formats the value of the least granular unit in the
// [NSDateComponents] object, and doesn’t provide a compound format of the
// date component.
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/localizedString(from:)
func (r RelativeDateTimeFormatter) LocalizedStringFromDateComponents(dateComponents INSDateComponents) string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("localizedStringFromDateComponents:"), dateComponents)
	return NSStringFromID(rv).String()
}
// Formats the specified time interval using the formatter’s calendar.
//
// timeInterval: The time interval to format.
//
// # Return Value
// 
// A string that represents the formatted time interval.
//
// # Discussion
// 
// The formatter interprets a negative time interval as a date in the past.
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/localizedString(fromTimeInterval:)
func (r RelativeDateTimeFormatter) LocalizedStringFromTimeInterval(timeInterval float64) string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("localizedStringFromTimeInterval:"), timeInterval)
	return NSStringFromID(rv).String()
}

// The calendar to use for formatting values that don’t have an inherent
// calendar of their own.
//
// # Discussion
// 
// Defaults to [AutoupdatingCurrentCalendar]. If you set this property to
// `nil`, the formatter resets to using [AutoupdatingCurrentCalendar].
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/calendar
func (r RelativeDateTimeFormatter) Calendar() INSCalendar {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("calendar"))
	return NSCalendarFromID(objc.ID(rv))
}
func (r RelativeDateTimeFormatter) SetCalendar(value INSCalendar) {
	objc.Send[struct{}](r.ID, objc.Sel("setCalendar:"), value)
}
// The locale to use when formatting the date.
//
// # Discussion
// 
// The default value is [AutoupdatingCurrentLocale]. If you set this property
// to `nil`, the formatter resets to using [AutoupdatingCurrentLocale].
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/locale
func (r RelativeDateTimeFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (r RelativeDateTimeFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](r.ID, objc.Sel("setLocale:"), value)
}
// The style to use when describing a relative date, for example
// “yesterday” or “1 day ago”.
//
// # Discussion
// 
// Default is `numeric`.
// 
// To display relative dates using named styles, set this property to `named`.
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/dateTimeStyle-swift.property
func (r RelativeDateTimeFormatter) DateTimeStyle() NSRelativeDateTimeFormatterStyle {
	rv := objc.Send[NSRelativeDateTimeFormatterStyle](r.ID, objc.Sel("dateTimeStyle"))
	return NSRelativeDateTimeFormatterStyle(rv)
}
func (r RelativeDateTimeFormatter) SetDateTimeStyle(value NSRelativeDateTimeFormatterStyle) {
	objc.Send[struct{}](r.ID, objc.Sel("setDateTimeStyle:"), value)
}
// The style to use when formatting the quantity or the name of the unit, such
// as “1 day ago” or “one day ago”.
//
// # Discussion
// 
// The default value is `full`.
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/unitsStyle-swift.property
func (r RelativeDateTimeFormatter) UnitsStyle() NSRelativeDateTimeFormatterUnitsStyle {
	rv := objc.Send[NSRelativeDateTimeFormatterUnitsStyle](r.ID, objc.Sel("unitsStyle"))
	return NSRelativeDateTimeFormatterUnitsStyle(rv)
}
func (r RelativeDateTimeFormatter) SetUnitsStyle(value NSRelativeDateTimeFormatterUnitsStyle) {
	objc.Send[struct{}](r.ID, objc.Sel("setUnitsStyle:"), value)
}
// A description of where the formatted string will appear, allowing the
// formatter to capitalize the output appropriately.
//
// # Discussion
// 
// The default value is [FormattingContextUnknown]. For additional details
// about specifying contexts, see [Formatter.Context].
//
// [Formatter.Context]: https://developer.apple.com/documentation/Foundation/Formatter/Context
//
// See: https://developer.apple.com/documentation/Foundation/RelativeDateTimeFormatter/formattingContext
func (r RelativeDateTimeFormatter) FormattingContext() NSFormattingContext {
	rv := objc.Send[NSFormattingContext](r.ID, objc.Sel("formattingContext"))
	return NSFormattingContext(rv)
}
func (r RelativeDateTimeFormatter) SetFormattingContext(value NSFormattingContext) {
	objc.Send[struct{}](r.ID, objc.Sel("setFormattingContext:"), value)
}

