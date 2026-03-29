// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [DateComponentsFormatter] class.
var (
	_DateComponentsFormatterClass     DateComponentsFormatterClass
	_DateComponentsFormatterClassOnce sync.Once
)

func getDateComponentsFormatterClass() DateComponentsFormatterClass {
	_DateComponentsFormatterClassOnce.Do(func() {
		_DateComponentsFormatterClass = DateComponentsFormatterClass{class: objc.GetClass("NSDateComponentsFormatter")}
	})
	return _DateComponentsFormatterClass
}

// GetDateComponentsFormatterClass returns the class object for NSDateComponentsFormatter.
func GetDateComponentsFormatterClass() DateComponentsFormatterClass {
	return getDateComponentsFormatterClass()
}

type DateComponentsFormatterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DateComponentsFormatterClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DateComponentsFormatterClass) Alloc() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that creates string representations of quantities of time.
//
// # Overview
// 
// An [NSDateComponentsFormatter] object takes quantities of time and formats
// them as a user-readable string. Use a date components formatter to create
// strings for your app’s interface. The formatter object has many options
// for creating both abbreviated and expanded strings. The formatter takes the
// current user’s locale and language into account when generating strings.
// 
// To use this class, create an instance, configure its properties, and call
// one of its methods to generate an appropriate string. The properties of
// this class let you configure the calendar and specify the date and time
// units you want displayed in the resulting string. The listing below shows
// how to configure a formatter to create the string “About 5 minutes
// remaining”.
// 
// The methods of this class may be called safely from any thread of your app.
// It is also safe to share a single instance of this class from multiple
// threads, with the caveat that you should not change the configuration of
// the object while another thread is using it to generate a string.
//
// # Formatting Values
//
//   - [DateComponentsFormatter.StringFromDateComponents]: Returns a formatted string based on the specified date component information.
//   - [DateComponentsFormatter.StringFromDateToDate]: Returns a formatted string based on the time difference between two dates.
//   - [DateComponentsFormatter.StringFromTimeInterval]: Returns a formatted string based on the specified number of seconds.
//
// # Configuring the Formatter Options
//
//   - [DateComponentsFormatter.AllowedUnits]: The bitmask of calendrical units such as day and month to include in the output string.
//   - [DateComponentsFormatter.SetAllowedUnits]
//   - [DateComponentsFormatter.AllowsFractionalUnits]: A Boolean indicating whether non-integer units may be used for values.
//   - [DateComponentsFormatter.SetAllowsFractionalUnits]
//   - [DateComponentsFormatter.Calendar]: The default calendar to use when formatting date components.
//   - [DateComponentsFormatter.SetCalendar]
//   - [DateComponentsFormatter.CollapsesLargestUnit]: A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.
//   - [DateComponentsFormatter.SetCollapsesLargestUnit]
//   - [DateComponentsFormatter.IncludesApproximationPhrase]: A Boolean value indicating whether the resulting phrase reflects an inexact time value.
//   - [DateComponentsFormatter.SetIncludesApproximationPhrase]
//   - [DateComponentsFormatter.IncludesTimeRemainingPhrase]: A Boolean value indicating whether output strings reflect the amount of time remaining.
//   - [DateComponentsFormatter.SetIncludesTimeRemainingPhrase]
//   - [DateComponentsFormatter.MaximumUnitCount]: The maximum number of time units to include in the output string.
//   - [DateComponentsFormatter.SetMaximumUnitCount]
//   - [DateComponentsFormatter.UnitsStyle]: The formatting style for unit names.
//   - [DateComponentsFormatter.SetUnitsStyle]
//   - [DateComponentsFormatter.ZeroFormattingBehavior]: The formatting style for units whose value is 0.
//   - [DateComponentsFormatter.SetZeroFormattingBehavior]
//
// # Instance Properties
//
//   - [DateComponentsFormatter.FormattingContext]
//   - [DateComponentsFormatter.SetFormattingContext]
//   - [DateComponentsFormatter.ReferenceDate]
//   - [DateComponentsFormatter.SetReferenceDate]
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter
type DateComponentsFormatter struct {
	NSFormatter
}

// DateComponentsFormatterFromID constructs a [DateComponentsFormatter] from an objc.ID.
//
// A formatter that creates string representations of quantities of time.
func DateComponentsFormatterFromID(id objc.ID) DateComponentsFormatter {
	return NSDateComponentsFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSDateComponentsFormatterFromID is an alias for [DateComponentsFormatterFromID] for cross-framework compatibility.
func NSDateComponentsFormatterFromID(id objc.ID) DateComponentsFormatter { return DateComponentsFormatterFromID(id) }
// NOTE: DateComponentsFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [DateComponentsFormatter] class.
//
// # Formatting Values
//
//   - [IDateComponentsFormatter.StringFromDateComponents]: Returns a formatted string based on the specified date component information.
//   - [IDateComponentsFormatter.StringFromDateToDate]: Returns a formatted string based on the time difference between two dates.
//   - [IDateComponentsFormatter.StringFromTimeInterval]: Returns a formatted string based on the specified number of seconds.
//
// # Configuring the Formatter Options
//
//   - [IDateComponentsFormatter.AllowedUnits]: The bitmask of calendrical units such as day and month to include in the output string.
//   - [IDateComponentsFormatter.SetAllowedUnits]
//   - [IDateComponentsFormatter.AllowsFractionalUnits]: A Boolean indicating whether non-integer units may be used for values.
//   - [IDateComponentsFormatter.SetAllowsFractionalUnits]
//   - [IDateComponentsFormatter.Calendar]: The default calendar to use when formatting date components.
//   - [IDateComponentsFormatter.SetCalendar]
//   - [IDateComponentsFormatter.CollapsesLargestUnit]: A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.
//   - [IDateComponentsFormatter.SetCollapsesLargestUnit]
//   - [IDateComponentsFormatter.IncludesApproximationPhrase]: A Boolean value indicating whether the resulting phrase reflects an inexact time value.
//   - [IDateComponentsFormatter.SetIncludesApproximationPhrase]
//   - [IDateComponentsFormatter.IncludesTimeRemainingPhrase]: A Boolean value indicating whether output strings reflect the amount of time remaining.
//   - [IDateComponentsFormatter.SetIncludesTimeRemainingPhrase]
//   - [IDateComponentsFormatter.MaximumUnitCount]: The maximum number of time units to include in the output string.
//   - [IDateComponentsFormatter.SetMaximumUnitCount]
//   - [IDateComponentsFormatter.UnitsStyle]: The formatting style for unit names.
//   - [IDateComponentsFormatter.SetUnitsStyle]
//   - [IDateComponentsFormatter.ZeroFormattingBehavior]: The formatting style for units whose value is 0.
//   - [IDateComponentsFormatter.SetZeroFormattingBehavior]
//
// # Instance Properties
//
//   - [IDateComponentsFormatter.FormattingContext]
//   - [IDateComponentsFormatter.SetFormattingContext]
//   - [IDateComponentsFormatter.ReferenceDate]
//   - [IDateComponentsFormatter.SetReferenceDate]
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter
type IDateComponentsFormatter interface {
	INSFormatter

	// Topic: Formatting Values

	// Returns a formatted string based on the specified date component information.
	StringFromDateComponents(components INSDateComponents) string
	// Returns a formatted string based on the time difference between two dates.
	StringFromDateToDate(startDate INSDate, endDate INSDate) string
	// Returns a formatted string based on the specified number of seconds.
	StringFromTimeInterval(ti float64) string

	// Topic: Configuring the Formatter Options

	// The bitmask of calendrical units such as day and month to include in the output string.
	AllowedUnits() NSCalendarUnit
	SetAllowedUnits(value NSCalendarUnit)
	// A Boolean indicating whether non-integer units may be used for values.
	AllowsFractionalUnits() bool
	SetAllowsFractionalUnits(value bool)
	// The default calendar to use when formatting date components.
	Calendar() INSCalendar
	SetCalendar(value INSCalendar)
	// A Boolean value indicating whether to collapse the largest unit into smaller units when a certain threshold is met.
	CollapsesLargestUnit() bool
	SetCollapsesLargestUnit(value bool)
	// A Boolean value indicating whether the resulting phrase reflects an inexact time value.
	IncludesApproximationPhrase() bool
	SetIncludesApproximationPhrase(value bool)
	// A Boolean value indicating whether output strings reflect the amount of time remaining.
	IncludesTimeRemainingPhrase() bool
	SetIncludesTimeRemainingPhrase(value bool)
	// The maximum number of time units to include in the output string.
	MaximumUnitCount() int
	SetMaximumUnitCount(value int)
	// The formatting style for unit names.
	UnitsStyle() NSDateComponentsFormatterUnitsStyle
	SetUnitsStyle(value NSDateComponentsFormatterUnitsStyle)
	// The formatting style for units whose value is 0.
	ZeroFormattingBehavior() NSDateComponentsFormatterZeroFormattingBehavior
	SetZeroFormattingBehavior(value NSDateComponentsFormatterZeroFormattingBehavior)

	// Topic: Instance Properties

	FormattingContext() NSFormattingContext
	SetFormattingContext(value NSFormattingContext)
	ReferenceDate() INSDate
	SetReferenceDate(value INSDate)
}

// Init initializes the instance.
func (d DateComponentsFormatter) Init() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DateComponentsFormatter) Autorelease() DateComponentsFormatter {
	rv := objc.Send[DateComponentsFormatter](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDateComponentsFormatter creates a new DateComponentsFormatter instance.
func NewDateComponentsFormatter() DateComponentsFormatter {
	class := getDateComponentsFormatterClass()
	rv := objc.Send[DateComponentsFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDateComponentsFormatterWithCoder(coder INSCoder) DateComponentsFormatter {
	instance := getDateComponentsFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DateComponentsFormatterFromID(rv)
}

// Returns a formatted string based on the specified date component
// information.
//
// components: A date components object containing the date and time information to
// format. The [AllowedUnits] property determines which date components are
// actually used to generate the string. All other date components are
// ignored. This parameter must not be `nil`.
//
// # Return Value
// 
// A formatted string representing the specified date information.
//
// # Discussion
// 
// Use this method to format date information that is already broken down into
// the component day and time values.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(from:)-9exxn
func (d DateComponentsFormatter) StringFromDateComponents(components INSDateComponents) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stringFromDateComponents:"), components)
	return NSStringFromID(rv).String()
}
// Returns a formatted string based on the time difference between two dates.
//
// startDate: The start time. This parameter must not be `nil`.
//
// endDate: The end time. This parameter must not be `nil`.
//
// # Return Value
// 
// A formatted string representing the specified time information.
//
// # Discussion
// 
// This method calculates the elapsed time between the `startDate` and
// `endDate` values and uses that information to generate the string. For
// example, if there is exactly one hour and ten minutes difference between
// the start and end dates, generating an abbreviated string would result in a
// string of “1h 10m”.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(from:to:)
func (d DateComponentsFormatter) StringFromDateToDate(startDate INSDate, endDate INSDate) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stringFromDate:toDate:"), startDate, endDate)
	return NSStringFromID(rv).String()
}
// Returns a formatted string based on the specified number of seconds.
//
// ti: The time interval, measured in seconds. The value must be a finite number.
// Negative numbers are treated as positive numbers when creating the string.
//
// # Return Value
// 
// A formatted string representing the specified time interval.
//
// # Discussion
// 
// This method formats the specified number of seconds into the appropriate
// units. For example, if the formatter allows the display of minutes and
// seconds, creating an abbreviated string for the value 70 seconds results in
// the string “1m 10s”.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/string(from:)-7sj4j
func (d DateComponentsFormatter) StringFromTimeInterval(ti float64) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("stringFromTimeInterval:"), ti)
	return NSStringFromID(rv).String()
}

// Returns a localized string based on the specified date components and style
// option.
//
// components: The value to format.
//
// unitsStyle: The style for the resulting units. Use this parameter to specify whether
// you want to the resulting string to use an abbreviated or more spelled out
// format.
//
// # Return Value
// 
// A string containing the localized date and time information.
//
// # Discussion
// 
// Use this convenience method to format a string using the default formatter
// values, with the exception of the `unitsStyle` value.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/localizedString(from:unitsStyle:)
func (_DateComponentsFormatterClass DateComponentsFormatterClass) LocalizedStringFromDateComponentsUnitsStyle(components INSDateComponents, unitsStyle NSDateComponentsFormatterUnitsStyle) string {
	rv := objc.Send[objc.ID](objc.ID(_DateComponentsFormatterClass.class), objc.Sel("localizedStringFromDateComponents:unitsStyle:"), components, unitsStyle)
	return NSStringFromID(rv).String()
}

// The bitmask of calendrical units such as day and month to include in the
// output string.
//
// # Discussion
// 
// The allowed calendar units are:
// 
// - [CalendarUnitYear] - [CalendarUnitMonth] - [CalendarUnitWeekOfMonth] -
// [CalendarUnitDay] - [CalendarUnitHour] - [CalendarUnitMinute] -
// [CalendarUnitSecond]
// 
// Assigning any other calendar units to this property results in an
// exception.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/allowedUnits
func (d DateComponentsFormatter) AllowedUnits() NSCalendarUnit {
	rv := objc.Send[NSCalendarUnit](d.ID, objc.Sel("allowedUnits"))
	return NSCalendarUnit(rv)
}
func (d DateComponentsFormatter) SetAllowedUnits(value NSCalendarUnit) {
	objc.Send[struct{}](d.ID, objc.Sel("setAllowedUnits:"), value)
}
// A Boolean indicating whether non-integer units may be used for values.
//
// # Discussion
// 
// Fractional units may be used when a value cannot be exactly represented
// using the available units. For example, if minutes are not allowed, the
// value “1h 30m” could be formatted as “1.5h”.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/allowsFractionalUnits
func (d DateComponentsFormatter) AllowsFractionalUnits() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowsFractionalUnits"))
	return rv
}
func (d DateComponentsFormatter) SetAllowsFractionalUnits(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setAllowsFractionalUnits:"), value)
}
// The default calendar to use when formatting date components.
//
// # Discussion
// 
// The formatter uses the calendar in this property to format values that do
// not have an inherent calendar of their own. For example, the formatter uses
// this calendar when formatting an [NSTimeInterval] value.
// 
// The default value of this property is the calendar returned by the
// [AutoupdatingCurrentCalendar] method of [NSCalendar]. Setting this property
// to `nil` causes the formatter to use the Gregorian calendar with the
// `en_US_POSIX` locale.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/calendar
func (d DateComponentsFormatter) Calendar() INSCalendar {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("calendar"))
	return NSCalendarFromID(objc.ID(rv))
}
func (d DateComponentsFormatter) SetCalendar(value INSCalendar) {
	objc.Send[struct{}](d.ID, objc.Sel("setCalendar:"), value)
}
// A Boolean value indicating whether to collapse the largest unit into
// smaller units when a certain threshold is met.
//
// # Discussion
// 
// An example of when this property might apply is when expressing 63 seconds
// worth of time. When this property is set to [true], the formatted value
// would be “63s”. When the value of this property is [false], the
// formatted value would be “1m 3s”.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/collapsesLargestUnit
func (d DateComponentsFormatter) CollapsesLargestUnit() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("collapsesLargestUnit"))
	return rv
}
func (d DateComponentsFormatter) SetCollapsesLargestUnit(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setCollapsesLargestUnit:"), value)
}
// A Boolean value indicating whether the resulting phrase reflects an inexact
// time value.
//
// # Discussion
// 
// Setting the value of this property to [true] adds phrasing to output
// strings to reflect that the given time value is approximate and not exact.
// Using this property yields more correct phrasing than simply prepending the
// string “About” to an output string.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/includesApproximationPhrase
func (d DateComponentsFormatter) IncludesApproximationPhrase() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("includesApproximationPhrase"))
	return rv
}
func (d DateComponentsFormatter) SetIncludesApproximationPhrase(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setIncludesApproximationPhrase:"), value)
}
// A Boolean value indicating whether output strings reflect the amount of
// time remaining.
//
// # Discussion
// 
// Setting this property to [true] results in output strings like “30
// minutes remaining”.
// 
// The default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/includesTimeRemainingPhrase
func (d DateComponentsFormatter) IncludesTimeRemainingPhrase() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("includesTimeRemainingPhrase"))
	return rv
}
func (d DateComponentsFormatter) SetIncludesTimeRemainingPhrase(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setIncludesTimeRemainingPhrase:"), value)
}
// The maximum number of time units to include in the output string.
//
// # Discussion
// 
// Use this property to limit the number of units displayed in the resulting
// string. For example, with this property set to 2, instead of “1h 10m,
// 30s”, the resulting string would be “1h 10m”. Use this property when
// you are constrained for space or want to round up values to the nearest
// large unit.
// 
// The default value of this property is `0`, which does not cause the
// elimination of any units.
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/maximumUnitCount
func (d DateComponentsFormatter) MaximumUnitCount() int {
	rv := objc.Send[int](d.ID, objc.Sel("maximumUnitCount"))
	return rv
}
func (d DateComponentsFormatter) SetMaximumUnitCount(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaximumUnitCount:"), value)
}
// The formatting style for unit names.
//
// # Discussion
// 
// Configures the strings to use (if any) for unit names such as days, hours,
// minutes, and seconds. Use this property to specify whether you want
// abbreviated or shortened versions of unit names—for example, `hrs`
// instead of `hours`.
// 
// The default value of this property is
// [DateComponentsFormatterUnitsStylePositional].
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/unitsStyle-swift.property
func (d DateComponentsFormatter) UnitsStyle() NSDateComponentsFormatterUnitsStyle {
	rv := objc.Send[NSDateComponentsFormatterUnitsStyle](d.ID, objc.Sel("unitsStyle"))
	return NSDateComponentsFormatterUnitsStyle(rv)
}
func (d DateComponentsFormatter) SetUnitsStyle(value NSDateComponentsFormatterUnitsStyle) {
	objc.Send[struct{}](d.ID, objc.Sel("setUnitsStyle:"), value)
}
// The formatting style for units whose value is 0.
//
// # Discussion
// 
// When the value for a particular unit is 0, the zero formatting behavior
// determines whether that value is retained or omitted from any resulting
// strings. For example, when the formatting behavior is
// [DateComponentsFormatterZeroFormattingBehaviorDropTrailing], the value of
// one hour, ten minutes, and zero seconds would omit the mention of seconds.
// 
// The default value of this property is
// [DateComponentsFormatterZeroFormattingBehaviorDefault].
//
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/zeroFormattingBehavior-swift.property
func (d DateComponentsFormatter) ZeroFormattingBehavior() NSDateComponentsFormatterZeroFormattingBehavior {
	rv := objc.Send[NSDateComponentsFormatterZeroFormattingBehavior](d.ID, objc.Sel("zeroFormattingBehavior"))
	return NSDateComponentsFormatterZeroFormattingBehavior(rv)
}
func (d DateComponentsFormatter) SetZeroFormattingBehavior(value NSDateComponentsFormatterZeroFormattingBehavior) {
	objc.Send[struct{}](d.ID, objc.Sel("setZeroFormattingBehavior:"), value)
}
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/formattingContext
func (d DateComponentsFormatter) FormattingContext() NSFormattingContext {
	rv := objc.Send[NSFormattingContext](d.ID, objc.Sel("formattingContext"))
	return NSFormattingContext(rv)
}
func (d DateComponentsFormatter) SetFormattingContext(value NSFormattingContext) {
	objc.Send[struct{}](d.ID, objc.Sel("setFormattingContext:"), value)
}
// See: https://developer.apple.com/documentation/Foundation/DateComponentsFormatter/referenceDate
func (d DateComponentsFormatter) ReferenceDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("referenceDate"))
	return NSDateFromID(objc.ID(rv))
}
func (d DateComponentsFormatter) SetReferenceDate(value INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setReferenceDate:"), value)
}

