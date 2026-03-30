// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [DateIntervalFormatter] class.
var (
	_DateIntervalFormatterClass     DateIntervalFormatterClass
	_DateIntervalFormatterClassOnce sync.Once
)

func getDateIntervalFormatterClass() DateIntervalFormatterClass {
	_DateIntervalFormatterClassOnce.Do(func() {
		_DateIntervalFormatterClass = DateIntervalFormatterClass{class: objc.GetClass("NSDateIntervalFormatter")}
	})
	return _DateIntervalFormatterClass
}

// GetDateIntervalFormatterClass returns the class object for NSDateIntervalFormatter.
func GetDateIntervalFormatterClass() DateIntervalFormatterClass {
	return getDateIntervalFormatterClass()
}

type DateIntervalFormatterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DateIntervalFormatterClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DateIntervalFormatterClass) Alloc() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that creates string representations of time intervals.
//
// # Overview
//
// A [NSDateIntervalFormatter] object creates user-readable strings from pairs
// of dates. Use a date interval formatter to create user-readable strings of
// the form `-` for your app’s interface, where and are date values that you
// supply. The formatter uses locale and language information, along with
// custom formatting options, to define the content of the resulting string.
// You can specify different styles for the date and time information in each
// date value.
//
// To use this class, create an instance, configure its properties, and call
// the [StringFromDateToDate] method to generate a string. The properties of
// this class let you configure the calendar and specify the style to apply to
// date and time values. Given a current date of January 16, 2015, Configuring
// the Formatter Options shows how to configure a formatter object and
// generate the string “1/16/15 - 1/17/15”.
//
// # Configuring a formatter object
//
// The [StringFromDateToDate] method may be called safely from any thread of
// your app. It is also safe to share a single instance of this class from
// multiple threads, with the caveat that you should not change the
// configuration of the object while another thread is using it to generate a
// string.
//
// # Formatting a String
//
//   - [DateIntervalFormatter.StringFromDateToDate]: Returns a formatted string based on the specified start and end dates.
//
// # Configuring the Formatter Options
//
//   - [DateIntervalFormatter.DateStyle]: The style to use when formatting day, month, and year information.
//   - [DateIntervalFormatter.SetDateStyle]
//   - [DateIntervalFormatter.TimeStyle]: The style to use when formatting hour, minute, and second information.
//   - [DateIntervalFormatter.SetTimeStyle]
//   - [DateIntervalFormatter.DateTemplate]: The template for formatting one date and time value.
//   - [DateIntervalFormatter.SetDateTemplate]
//   - [DateIntervalFormatter.Calendar]: The calendar to use for date values.
//   - [DateIntervalFormatter.SetCalendar]
//   - [DateIntervalFormatter.Locale]: The locale to use when formatting date and time values.
//   - [DateIntervalFormatter.SetLocale]
//   - [DateIntervalFormatter.TimeZone]: The time zone with which to specify time values.
//   - [DateIntervalFormatter.SetTimeZone]
//
// # Instance Methods
//
//   - [DateIntervalFormatter.StringFromDateInterval]
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter
type DateIntervalFormatter struct {
	NSFormatter
}

// DateIntervalFormatterFromID constructs a [DateIntervalFormatter] from an objc.ID.
//
// A formatter that creates string representations of time intervals.
func DateIntervalFormatterFromID(id objc.ID) DateIntervalFormatter {
	return NSDateIntervalFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSDateIntervalFormatterFromID is an alias for [DateIntervalFormatterFromID] for cross-framework compatibility.
func NSDateIntervalFormatterFromID(id objc.ID) DateIntervalFormatter {
	return DateIntervalFormatterFromID(id)
}

// NOTE: DateIntervalFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [DateIntervalFormatter] class.
//
// # Formatting a String
//
//   - [IDateIntervalFormatter.StringFromDateToDate]: Returns a formatted string based on the specified start and end dates.
//
// # Configuring the Formatter Options
//
//   - [IDateIntervalFormatter.DateStyle]: The style to use when formatting day, month, and year information.
//   - [IDateIntervalFormatter.SetDateStyle]
//   - [IDateIntervalFormatter.TimeStyle]: The style to use when formatting hour, minute, and second information.
//   - [IDateIntervalFormatter.SetTimeStyle]
//   - [IDateIntervalFormatter.DateTemplate]: The template for formatting one date and time value.
//   - [IDateIntervalFormatter.SetDateTemplate]
//   - [IDateIntervalFormatter.Calendar]: The calendar to use for date values.
//   - [IDateIntervalFormatter.SetCalendar]
//   - [IDateIntervalFormatter.Locale]: The locale to use when formatting date and time values.
//   - [IDateIntervalFormatter.SetLocale]
//   - [IDateIntervalFormatter.TimeZone]: The time zone with which to specify time values.
//   - [IDateIntervalFormatter.SetTimeZone]
//
// # Instance Methods
//
//   - [IDateIntervalFormatter.StringFromDateInterval]
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter
type IDateIntervalFormatter interface {
	INSFormatter

	// Topic: Formatting a String

	// Returns a formatted string based on the specified start and end dates.
	StringFromDateToDate(fromDate INSDate, toDate INSDate) string

	// Topic: Configuring the Formatter Options

	// The style to use when formatting day, month, and year information.
	DateStyle() NSDateIntervalFormatterStyle
	SetDateStyle(value NSDateIntervalFormatterStyle)
	// The style to use when formatting hour, minute, and second information.
	TimeStyle() NSDateIntervalFormatterStyle
	SetTimeStyle(value NSDateIntervalFormatterStyle)
	// The template for formatting one date and time value.
	DateTemplate() string
	SetDateTemplate(value string)
	// The calendar to use for date values.
	Calendar() INSCalendar
	SetCalendar(value INSCalendar)
	// The locale to use when formatting date and time values.
	Locale() INSLocale
	SetLocale(value INSLocale)
	// The time zone with which to specify time values.
	TimeZone() INSTimeZone
	SetTimeZone(value INSTimeZone)

	// Topic: Instance Methods

	StringFromDateInterval(dateInterval INSDateInterval) string
}

// Init initializes the instance.
func (d DateIntervalFormatter) Init() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DateIntervalFormatter) Autorelease() DateIntervalFormatter {
	rv := objc.Send[DateIntervalFormatter](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateIntervalFormatter creates a new DateIntervalFormatter instance.
func NewDateIntervalFormatter() DateIntervalFormatter {
	class := getDateIntervalFormatterClass()
	rv := objc.Send[DateIntervalFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDateIntervalFormatterWithCoder(coder INSCoder) DateIntervalFormatter {
	instance := getDateIntervalFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DateIntervalFormatterFromID(rv)
}

// Returns a formatted string based on the specified start and end dates.
//
// fromDate: The start date. This date appears first in the resulting string.
//
// toDate: The end date. This date appears last after the hyphen in the resulting
// string.
//
// # Return Value
//
// A formatted string representing the specified date interval.
//
// # Discussion
//
// The formatter includes both `fromDate` and `toDate` in the resulting string
// only when there is enough of a difference in their values to warrant the
// inclusion of both. If the date and time difference cannot be adequately
// displayed, the formatter displays one date value. For example, if the
// [TimeStyle] property was set to [NSDateIntervalFormatterNoStyle], the two
// dates would need to be at least one day apart in order for both to be
// displayed.
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/string(from:to:)
func (d DateIntervalFormatter) StringFromDateToDate(fromDate INSDate, toDate INSDate) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stringFromDate:toDate:"), fromDate, toDate)
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/string(from:)
func (d DateIntervalFormatter) StringFromDateInterval(dateInterval INSDateInterval) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stringFromDateInterval:"), dateInterval)
	return NSStringFromID(rv).String()
}

// The style to use when formatting day, month, and year information.
//
// # Discussion
//
// Set this property to an appropriate value before generating string values.
// The default value of this property is [NSDateIntervalFormatterNoStyle].
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/dateStyle
func (d DateIntervalFormatter) DateStyle() NSDateIntervalFormatterStyle {
	rv := objc.Send[NSDateIntervalFormatterStyle](d.ID, objc.Sel("dateStyle"))
	return NSDateIntervalFormatterStyle(rv)
}
func (d DateIntervalFormatter) SetDateStyle(value NSDateIntervalFormatterStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setDateStyle:"), value)
}

// The style to use when formatting hour, minute, and second information.
//
// # Discussion
//
// Set this property to an appropriate value before generating string values.
// The default value of this property is [NSDateIntervalFormatterNoStyle].
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/timeStyle
func (d DateIntervalFormatter) TimeStyle() NSDateIntervalFormatterStyle {
	rv := objc.Send[NSDateIntervalFormatterStyle](d.ID, objc.Sel("timeStyle"))
	return NSDateIntervalFormatterStyle(rv)
}
func (d DateIntervalFormatter) SetTimeStyle(value NSDateIntervalFormatterStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeStyle:"), value)
}

// The template for formatting one date and time value.
//
// # Discussion
//
// Use this string to specify a custom fixed format for each of the date and
// time values. The string you specify is based on the Unicode Technical
// Standard #35, which uses characters to represent the day, time, year, hour,
// minute, and other pieces of date or time information.
//
// If you do not assign a value to this string, the formatter object
// automatically updates the string based on the values in the [DateStyle] and
// [TimeStyle] properties.
//
// For information about how to define a custom formatting string, see [Date
// Formatters] in [Data Formatting Guide].
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/dateTemplate
//
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
// [Date Formatters]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/Articles/dfDateFormatting10_4.html#//apple_ref/doc/uid/TP40002369
func (d DateIntervalFormatter) DateTemplate() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dateTemplate"))
	return NSStringFromID(rv).String()
}
func (d DateIntervalFormatter) SetDateTemplate(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setDateTemplate:"), objc.String(value))
}

// The calendar to use for date values.
//
// # Discussion
//
// The default value of this property is the calendar associated with the
// current [Locale] object. You can change this value to use a different
// calendar for interpreting dates.
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/calendar
func (d DateIntervalFormatter) Calendar() INSCalendar {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("calendar"))
	return NSCalendarFromID(objc.ID(rv))
}
func (d DateIntervalFormatter) SetCalendar(value INSCalendar) {
	objc.Send[struct{}](d.ID, objc.Sel("setCalendar:"), value)
}

// The locale to use when formatting date and time values.
//
// # Discussion
//
// The default value of this property is the current user’s locale, which is
// accessible from the [CurrentLocale] method of [NSLocale]. You can change
// this value to a different locale to generate strings based on that locale.
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/locale
func (d DateIntervalFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (d DateIntervalFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](d.ID, objc.Sel("setLocale:"), value)
}

// The time zone with which to specify time values.
//
// # Discussion
//
// The default value of this property is the default time zone for the current
// user, which is accessible from the [DefaultTimeZone] method of
// [NSTimeZone]. You can change this value to a different time zone to
// generate strings based on that time zone.
//
// See: https://developer.apple.com/documentation/Foundation/DateIntervalFormatter/timeZone
func (d DateIntervalFormatter) TimeZone() INSTimeZone {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
func (d DateIntervalFormatter) SetTimeZone(value INSTimeZone) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeZone:"), value)
}
