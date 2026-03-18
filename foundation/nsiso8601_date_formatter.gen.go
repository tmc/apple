// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [ISO8601DateFormatter] class.
var (
	_ISO8601DateFormatterClass     ISO8601DateFormatterClass
	_ISO8601DateFormatterClassOnce sync.Once
)

func getISO8601DateFormatterClass() ISO8601DateFormatterClass {
	_ISO8601DateFormatterClassOnce.Do(func() {
		_ISO8601DateFormatterClass = ISO8601DateFormatterClass{class: objc.GetClass("NSISO8601DateFormatter")}
	})
	return _ISO8601DateFormatterClass
}

// GetISO8601DateFormatterClass returns the class object for NSISO8601DateFormatter.
func GetISO8601DateFormatterClass() ISO8601DateFormatterClass {
	return getISO8601DateFormatterClass()
}

type ISO8601DateFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ic ISO8601DateFormatterClass) Alloc() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// A formatter that converts between dates and their ISO 8601 string
// representations.
//
// # Overview
// 
// The [NSISO8601DateFormatter] class generates and parses string
// representations of dates following the [ISO 8601] standard. Use this class
// to create ISO 8601 representations of dates and create dates from text
// strings in ISO 8601 format.
//
// [ISO 8601]: http://www.iso.org/iso/home/standards/iso8601
//
// # Configuring the Formatter
//
//   - [ISO8601DateFormatter.FormatOptions]: Options for generating and parsing ISO 8601 date representations. See [ISO8601DateFormatter.Options](<doc://com.apple.foundation/documentation/Foundation/ISO8601DateFormatter/Options>) for possible values.
//   - [ISO8601DateFormatter.SetFormatOptions]
//   - [ISO8601DateFormatter.TimeZone]: The time zone used to create and parse date representations. When unspecified, GMT is used.
//   - [ISO8601DateFormatter.SetTimeZone]
//
// # Converting ISO 8601 Dates
//
//   - [ISO8601DateFormatter.StringFromDate]: Creates and returns an ISO 8601 formatted string representation of the specified date.
//   - [ISO8601DateFormatter.DateFromString]: Creates and returns a date object from the specified ISO 8601 formatted string representation.
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter
type ISO8601DateFormatter struct {
	NSFormatter
}

// ISO8601DateFormatterFromID constructs a [ISO8601DateFormatter] from an objc.ID.
//
// A formatter that converts between dates and their ISO 8601 string
// representations.
func ISO8601DateFormatterFromID(id objc.ID) ISO8601DateFormatter {
	return NSISO8601DateFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSISO8601DateFormatterFromID is an alias for [ISO8601DateFormatterFromID] for cross-framework compatibility.
func NSISO8601DateFormatterFromID(id objc.ID) ISO8601DateFormatter { return ISO8601DateFormatterFromID(id) }
// NOTE: ISO8601DateFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ISO8601DateFormatter] class.
//
// # Configuring the Formatter
//
//   - [IISO8601DateFormatter.FormatOptions]: Options for generating and parsing ISO 8601 date representations. See [ISO8601DateFormatter.Options](<doc://com.apple.foundation/documentation/Foundation/ISO8601DateFormatter/Options>) for possible values.
//   - [IISO8601DateFormatter.SetFormatOptions]
//   - [IISO8601DateFormatter.TimeZone]: The time zone used to create and parse date representations. When unspecified, GMT is used.
//   - [IISO8601DateFormatter.SetTimeZone]
//
// # Converting ISO 8601 Dates
//
//   - [IISO8601DateFormatter.StringFromDate]: Creates and returns an ISO 8601 formatted string representation of the specified date.
//   - [IISO8601DateFormatter.DateFromString]: Creates and returns a date object from the specified ISO 8601 formatted string representation.
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter
type IISO8601DateFormatter interface {
	INSFormatter
	NSSecureCoding

	// Topic: Configuring the Formatter

	// Options for generating and parsing ISO 8601 date representations. See [ISO8601DateFormatter.Options](<doc://com.apple.foundation/documentation/Foundation/ISO8601DateFormatter/Options>) for possible values.
	FormatOptions() NSISO8601DateFormatOptions
	SetFormatOptions(value NSISO8601DateFormatOptions)
	// The time zone used to create and parse date representations. When unspecified, GMT is used.
	TimeZone() INSTimeZone
	SetTimeZone(value INSTimeZone)

	// Topic: Converting ISO 8601 Dates

	// Creates and returns an ISO 8601 formatted string representation of the specified date.
	StringFromDate(date INSDate) string
	// Creates and returns a date object from the specified ISO 8601 formatted string representation.
	DateFromString(string_ string) INSDate
}

// Init initializes the instance.
func (i ISO8601DateFormatter) Init() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i ISO8601DateFormatter) Autorelease() ISO8601DateFormatter {
	rv := objc.Send[ISO8601DateFormatter](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewISO8601DateFormatter creates a new ISO8601DateFormatter instance.
func NewISO8601DateFormatter() ISO8601DateFormatter {
	class := getISO8601DateFormatterClass()
	rv := objc.Send[ISO8601DateFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewISO8601DateFormatterWithCoder(coder INSCoder) ISO8601DateFormatter {
	instance := getISO8601DateFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ISO8601DateFormatterFromID(rv)
}

// Creates and returns an ISO 8601 formatted string representation of the
// specified date.
//
// date: The date to be represented.
//
// # Return Value
// 
// A user-readable string representing the date.
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/string(from:)
func (i ISO8601DateFormatter) StringFromDate(date INSDate) string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("stringFromDate:"), date)
	return NSStringFromID(rv).String()
}

// Creates and returns a date object from the specified ISO 8601 formatted
// string representation.
//
// string: The ISO 8601 formatted string representation of a date.
//
// # Return Value
// 
// A date object, or `nil` if no valid date was found.
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/date(from:)
func (i ISO8601DateFormatter) DateFromString(string_ string) INSDate {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("dateFromString:"), objc.String(string_))
	return NSDateFromID(rv)
}

// Creates a representation of the specified date with a given time zone and
// format options.
//
// date: The date to be represented.
//
// timeZone: The time zone used.
//
// formatOptions: The options used. For possible values, see [ISO8601DateFormatter.Options].
// //
// [ISO8601DateFormatter.Options]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options
//
// # Return Value
// 
// A user-readable string representing the date.
//
// # Discussion
// 
// This method uses a date formatter configured with the specified time zone
// and format options. The following code examples produce the same string
// value:
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/string(from:timeZone:formatOptions:)
func (_ISO8601DateFormatterClass ISO8601DateFormatterClass) StringFromDateTimeZoneFormatOptions(date INSDate, timeZone INSTimeZone, formatOptions NSISO8601DateFormatOptions) string {
	rv := objc.Send[objc.ID](objc.ID(_ISO8601DateFormatterClass.class), objc.Sel("stringFromDate:timeZone:formatOptions:"), date, timeZone, formatOptions)
	return NSStringFromID(rv).String()
}

// Options for generating and parsing ISO 8601 date representations. See
// [ISO8601DateFormatter.Options] for possible values.
//
// [ISO8601DateFormatter.Options]: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/Options
//
// # Discussion
// 
// The ISO 8601 specification allows for dates to be expressed in a variety of
// ways. You can configure the format used to parse and generate
// representations by specifying various combinations of format options.
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/formatOptions
func (i ISO8601DateFormatter) FormatOptions() NSISO8601DateFormatOptions {
	rv := objc.Send[NSISO8601DateFormatOptions](i.ID, objc.Sel("formatOptions"))
	return NSISO8601DateFormatOptions(rv)
}
func (i ISO8601DateFormatter) SetFormatOptions(value NSISO8601DateFormatOptions) {
	objc.Send[struct{}](i.ID, objc.Sel("setFormatOptions:"), value)
}

// The time zone used to create and parse date representations. When
// unspecified, GMT is used.
//
// # Discussion
// 
// Resetting this property can incur a significant performance cost, as it may
// cause internal state to be regenerated.
//
// See: https://developer.apple.com/documentation/Foundation/ISO8601DateFormatter/timeZone
func (i ISO8601DateFormatter) TimeZone() INSTimeZone {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("timeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
func (i ISO8601DateFormatter) SetTimeZone(value INSTimeZone) {
	objc.Send[struct{}](i.ID, objc.Sel("setTimeZone:"), value)
}

			// Protocol methods for NSSecureCoding
			

