// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDateComponents] class.
var (
	_NSDateComponentsClass     NSDateComponentsClass
	_NSDateComponentsClassOnce sync.Once
)

func getNSDateComponentsClass() NSDateComponentsClass {
	_NSDateComponentsClassOnce.Do(func() {
		_NSDateComponentsClass = NSDateComponentsClass{class: objc.GetClass("NSDateComponents")}
	})
	return _NSDateComponentsClass
}

// GetNSDateComponentsClass returns the class object for NSDateComponents.
func GetNSDateComponentsClass() NSDateComponentsClass {
	return getNSDateComponentsClass()
}

type NSDateComponentsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDateComponentsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDateComponentsClass) Alloc() NSDateComponents {
	rv := objc.Send[NSDateComponents](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that specifies a date or time in terms of units (such as year,
// month, day, hour, and minute) to be evaluated in a calendar system and time
// zone.
//
// # Overview
//
// In Swift, this object bridges to [DateComponents]; use [NSDateComponents]
// when you need reference semantics or other Foundation-specific behavior.
//
// [NSDateComponents] encapsulates the components of a date in an extendable,
// object-oriented manner. It’s used to specify a date by providing the
// temporal components that make up a date and time: hour, minutes, seconds,
// day, month, year, and so on. You can also use it to specify a duration of
// time, for example, 5 hours and 16 minutes. An [NSDateComponents] object is
// not required to define all the component fields. When a new instance of
// [NSDateComponents] is created, the date components are set to
// [NSDateComponents.NSDateComponentUndefined].
//
// An instance of [NSDateComponents] is not responsible for answering
// questions about a date beyond the information with which it was
// initialized. For example, if you initialize one with May 4, 2017, its
// weekday is [NSDateComponents.NSDateComponentUndefined], not Thursday. To get the correct day
// of the week, you must create a suitable instance of [NSCalendar], create an
// [NSDate] object using [DateFromComponents] and then use
// [ComponentsFromDate] to retrieve the weekday—as illustrated in the
// following example.
//
// For more details, see [Calendars, Date Components, and Calendar Units] in
// [Date and Time Programming Guide].
//
// # Setting a Calendar and Time Zone
//
//   - [NSDateComponents.Calendar]: The calendar used to interpret the date components.
//   - [NSDateComponents.SetCalendar]
//   - [NSDateComponents.TimeZone]: The time zone used to interpret the date components.
//   - [NSDateComponents.SetTimeZone]
//
// # Validating a Date
//
//   - [NSDateComponents.ValidDate]: A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar.
//   - [NSDateComponents.IsValidDateInCalendar]: Returns a Boolean value that indicates whether the current combination of properties represents a date which exists in the specified calendar.
//   - [NSDateComponents.Date]: The date calculated from the current components using the stored calendar.
//
// # Accessing Years and Months
//
//   - [NSDateComponents.Era]: The number of eras.
//   - [NSDateComponents.SetEra]
//   - [NSDateComponents.Year]: The number of years.
//   - [NSDateComponents.SetYear]
//   - [NSDateComponents.YearForWeekOfYear]: The ISO 8601 week-numbering year.
//   - [NSDateComponents.SetYearForWeekOfYear]
//   - [NSDateComponents.Quarter]: The number of quarters.
//   - [NSDateComponents.SetQuarter]
//   - [NSDateComponents.Month]: The number of months.
//   - [NSDateComponents.SetMonth]
//   - [NSDateComponents.LeapMonth]: A Boolean value that indicates whether the month is a leap month.
//   - [NSDateComponents.SetLeapMonth]
//
// # Accessing Weeks and Days
//
//   - [NSDateComponents.Weekday]: The number of the weekdays.
//   - [NSDateComponents.SetWeekday]
//   - [NSDateComponents.WeekdayOrdinal]: The ordinal number of weekdays.
//   - [NSDateComponents.SetWeekdayOrdinal]
//   - [NSDateComponents.WeekOfMonth]: The week number of the months.
//   - [NSDateComponents.SetWeekOfMonth]
//   - [NSDateComponents.WeekOfYear]: The ISO 8601 week date of the year.
//   - [NSDateComponents.SetWeekOfYear]
//   - [NSDateComponents.Day]: The number of days.
//   - [NSDateComponents.SetDay]
//
// # Accessing Hours and Seconds
//
//   - [NSDateComponents.Hour]: The number of hour units for the receiver.
//   - [NSDateComponents.SetHour]
//   - [NSDateComponents.Minute]: The number of minute units for the receiver.
//   - [NSDateComponents.SetMinute]
//   - [NSDateComponents.Second]: The number of second units for the receiver.
//   - [NSDateComponents.SetSecond]
//   - [NSDateComponents.Nanosecond]: The number of nanosecond units for the receiver.
//   - [NSDateComponents.SetNanosecond]
//
// # Accessing Components as Calendrical Units
//
//   - [NSDateComponents.ValueForComponent]: Returns the value for a given calendar unit.
//   - [NSDateComponents.SetValueForComponent]: Sets a value for a given calendar unit.
//
// # Instance Properties
//
//   - [NSDateComponents.DayOfYear]
//   - [NSDateComponents.SetDayOfYear]
//   - [NSDateComponents.RepeatedDay]
//   - [NSDateComponents.SetRepeatedDay]
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
// [DateComponents]: https://developer.apple.com/documentation/Foundation/DateComponents
type NSDateComponents struct {
	objectivec.Object
}

// NSDateComponentsFromID constructs a [NSDateComponents] from an objc.ID.
//
// An object that specifies a date or time in terms of units (such as year,
// month, day, hour, and minute) to be evaluated in a calendar system and time
// zone.
func NSDateComponentsFromID(id objc.ID) NSDateComponents {
	return NSDateComponents{objectivec.Object{ID: id}}
}

// NOTE: NSDateComponents adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDateComponents] class.
//
// # Setting a Calendar and Time Zone
//
//   - [INSDateComponents.Calendar]: The calendar used to interpret the date components.
//   - [INSDateComponents.SetCalendar]
//   - [INSDateComponents.TimeZone]: The time zone used to interpret the date components.
//   - [INSDateComponents.SetTimeZone]
//
// # Validating a Date
//
//   - [INSDateComponents.ValidDate]: A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar.
//   - [INSDateComponents.IsValidDateInCalendar]: Returns a Boolean value that indicates whether the current combination of properties represents a date which exists in the specified calendar.
//   - [INSDateComponents.Date]: The date calculated from the current components using the stored calendar.
//
// # Accessing Years and Months
//
//   - [INSDateComponents.Era]: The number of eras.
//   - [INSDateComponents.SetEra]
//   - [INSDateComponents.Year]: The number of years.
//   - [INSDateComponents.SetYear]
//   - [INSDateComponents.YearForWeekOfYear]: The ISO 8601 week-numbering year.
//   - [INSDateComponents.SetYearForWeekOfYear]
//   - [INSDateComponents.Quarter]: The number of quarters.
//   - [INSDateComponents.SetQuarter]
//   - [INSDateComponents.Month]: The number of months.
//   - [INSDateComponents.SetMonth]
//   - [INSDateComponents.LeapMonth]: A Boolean value that indicates whether the month is a leap month.
//   - [INSDateComponents.SetLeapMonth]
//
// # Accessing Weeks and Days
//
//   - [INSDateComponents.Weekday]: The number of the weekdays.
//   - [INSDateComponents.SetWeekday]
//   - [INSDateComponents.WeekdayOrdinal]: The ordinal number of weekdays.
//   - [INSDateComponents.SetWeekdayOrdinal]
//   - [INSDateComponents.WeekOfMonth]: The week number of the months.
//   - [INSDateComponents.SetWeekOfMonth]
//   - [INSDateComponents.WeekOfYear]: The ISO 8601 week date of the year.
//   - [INSDateComponents.SetWeekOfYear]
//   - [INSDateComponents.Day]: The number of days.
//   - [INSDateComponents.SetDay]
//
// # Accessing Hours and Seconds
//
//   - [INSDateComponents.Hour]: The number of hour units for the receiver.
//   - [INSDateComponents.SetHour]
//   - [INSDateComponents.Minute]: The number of minute units for the receiver.
//   - [INSDateComponents.SetMinute]
//   - [INSDateComponents.Second]: The number of second units for the receiver.
//   - [INSDateComponents.SetSecond]
//   - [INSDateComponents.Nanosecond]: The number of nanosecond units for the receiver.
//   - [INSDateComponents.SetNanosecond]
//
// # Accessing Components as Calendrical Units
//
//   - [INSDateComponents.ValueForComponent]: Returns the value for a given calendar unit.
//   - [INSDateComponents.SetValueForComponent]: Sets a value for a given calendar unit.
//
// # Instance Properties
//
//   - [INSDateComponents.DayOfYear]
//   - [INSDateComponents.SetDayOfYear]
//   - [INSDateComponents.RepeatedDay]
//   - [INSDateComponents.SetRepeatedDay]
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents
type INSDateComponents interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Setting a Calendar and Time Zone

	// The calendar used to interpret the date components.
	Calendar() INSCalendar
	SetCalendar(value INSCalendar)
	// The time zone used to interpret the date components.
	TimeZone() INSTimeZone
	SetTimeZone(value INSTimeZone)

	// Topic: Validating a Date

	// A Boolean value that indicates whether the current combination of properties represents a date which exists in the current calendar.
	ValidDate() bool
	// Returns a Boolean value that indicates whether the current combination of properties represents a date which exists in the specified calendar.
	IsValidDateInCalendar(calendar INSCalendar) bool
	// The date calculated from the current components using the stored calendar.
	Date() INSDate

	// Topic: Accessing Years and Months

	// The number of eras.
	Era() int
	SetEra(value int)
	// The number of years.
	Year() int
	SetYear(value int)
	// The ISO 8601 week-numbering year.
	YearForWeekOfYear() int
	SetYearForWeekOfYear(value int)
	// The number of quarters.
	Quarter() int
	SetQuarter(value int)
	// The number of months.
	Month() int
	SetMonth(value int)
	// A Boolean value that indicates whether the month is a leap month.
	LeapMonth() bool
	SetLeapMonth(value bool)

	// Topic: Accessing Weeks and Days

	// The number of the weekdays.
	Weekday() int
	SetWeekday(value int)
	// The ordinal number of weekdays.
	WeekdayOrdinal() int
	SetWeekdayOrdinal(value int)
	// The week number of the months.
	WeekOfMonth() int
	SetWeekOfMonth(value int)
	// The ISO 8601 week date of the year.
	WeekOfYear() int
	SetWeekOfYear(value int)
	// The number of days.
	Day() int
	SetDay(value int)

	// Topic: Accessing Hours and Seconds

	// The number of hour units for the receiver.
	Hour() int
	SetHour(value int)
	// The number of minute units for the receiver.
	Minute() int
	SetMinute(value int)
	// The number of second units for the receiver.
	Second() int
	SetSecond(value int)
	// The number of nanosecond units for the receiver.
	Nanosecond() int
	SetNanosecond(value int)

	// Topic: Accessing Components as Calendrical Units

	// Returns the value for a given calendar unit.
	ValueForComponent(unit NSCalendarUnit) int
	// Sets a value for a given calendar unit.
	SetValueForComponent(value int, unit NSCalendarUnit)

	// Topic: Instance Properties

	DayOfYear() int
	SetDayOfYear(value int)
	RepeatedDay() bool
	SetRepeatedDay(value bool)

	// Specifies a date component without a value.
	NSDateComponentUndefined() int
	SetNSDateComponentUndefined(value int)
}

// Init initializes the instance.
func (d NSDateComponents) Init() NSDateComponents {
	rv := objc.Send[NSDateComponents](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDateComponents) Autorelease() NSDateComponents {
	rv := objc.Send[NSDateComponents](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDateComponents creates a new NSDateComponents instance.
func NewNSDateComponents() NSDateComponents {
	class := getNSDateComponentsClass()
	rv := objc.Send[NSDateComponents](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDateComponentsWithCoder(coder INSCoder) NSDateComponents {
	instance := getNSDateComponentsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDateComponentsFromID(rv)
}

// Returns a Boolean value that indicates whether the current combination of
// properties represents a date which exists in the specified calendar.
//
// calendar: The calendar for which to use in the calculation.
//
// # Return Value
//
// true if the date corresponding to the receiver’s values is valid and
// exists in the given calendar, otherwise false.
//
// # Discussion
//
// If the [TimeZone] property is set on the receiver, the time zone property
// value is used.
//
// This property should not be used for [NSDateComponents] objects that
// represent relative quantities of calendar components. To find the the next
// or previous date that matches a particular set of date components, use the
// [NSCalendar] instance method [NextDateAfterDateMatchingUnitValueOptions]
// instead.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/isValidDate(in:)
func (d NSDateComponents) IsValidDateInCalendar(calendar INSCalendar) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isValidDateInCalendar:"), calendar)
	return rv
}

// Returns the value for a given calendar unit.
//
// unit: The calendar unit for which to retrieve its value. Do not pass
// [NSCalendarUnitCalendar] or [NSCalendarUnitTimeZone].
//
// # Return Value
//
// The value for the given calendar unit.
//
// # Discussion
//
// This method allows for component values to be retrieved for an
// [NSCalendar.Unit] value.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/value(forComponent:)
//
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
func (d NSDateComponents) ValueForComponent(unit NSCalendarUnit) int {
	rv := objc.Send[int](d.ID, objc.Sel("valueForComponent:"), unit)
	return rv
}

// Sets a value for a given calendar unit.
//
// value: The value to set for the `unit` component.
//
// unit: The calendar unit for which to set `value`. Do not pass
// [NSCalendarUnitCalendar] or [NSCalendarUnitTimeZone].
//
// # Discussion
//
// This method allows for component values to be set for an [NSCalendar.Unit]
// value.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/setValue(_:forComponent:)
//
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
func (d NSDateComponents) SetValueForComponent(value int, unit NSCalendarUnit) {
	objc.Send[objc.ID](d.ID, objc.Sel("setValue:forComponent:"), value, unit)
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (d NSDateComponents) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (d NSDateComponents) InitWithCoder(coder INSCoder) NSDateComponents {
	rv := objc.Send[NSDateComponents](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// The calendar used to interpret the date components.
//
// # Discussion
//
// See [Calendars, Date Components, and Calendar Units] in [Date and Time
// Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/calendar
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Calendar() INSCalendar {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("calendar"))
	return NSCalendarFromID(objc.ID(rv))
}
func (d NSDateComponents) SetCalendar(value INSCalendar) {
	objc.Send[struct{}](d.ID, objc.Sel("setCalendar:"), value)
}

// The time zone used to interpret the date components.
//
// # Discussion
//
// See [Calendars, Date Components, and Calendar Units] in [Date and Time
// Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/timeZone
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) TimeZone() INSTimeZone {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("timeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
func (d NSDateComponents) SetTimeZone(value INSTimeZone) {
	objc.Send[struct{}](d.ID, objc.Sel("setTimeZone:"), value)
}

// A Boolean value that indicates whether the current combination of
// properties represents a date which exists in the current calendar.
//
// # Discussion
//
// If the [TimeZone] property is set on the receiver, the time zone property
// value is used. If the [Calendar] property is not set on the receiver, `nil`
// is returned.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/isValidDate
func (d NSDateComponents) ValidDate() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isValidDate"))
	return rv
}

// The date calculated from the current components using the stored calendar.
//
// # Discussion
//
// Returns `nil` if the [Calendar] property value of the receiver is `nil` or
// cannot convert the receiver into an [NSDate] object.
//
// See [Calendars, Date Components, and Calendar Units] in [Date and Time
// Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/date
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Date() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("date"))
	return NSDateFromID(objc.ID(rv))
}

// The number of eras.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/era
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Era() int {
	rv := objc.Send[int](d.ID, objc.Sel("era"))
	return rv
}
func (d NSDateComponents) SetEra(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setEra:"), value)
}

// The number of years.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/year
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Year() int {
	rv := objc.Send[int](d.ID, objc.Sel("year"))
	return rv
}
func (d NSDateComponents) SetYear(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setYear:"), value)
}

// The ISO 8601 week-numbering year.
//
// # Discussion
//
// The Gregorian calendar defines a week to have 7 days, and a year to have
// 365 days, or 366 in a leap year. However, neither 365 or 366 divide evenly
// into a 7-day week, so it is often the case that the last week of a year
// ends on a day in the next year, and the first week of a year begins in the
// preceding year. To reconcile this, ISO 8601 defines a week-numbering year,
// consisting of either 52 or 53 full weeks (364 or 371 days), such that the
// first week of a year is designated to be the week containing the first
// Thursday of the year. For a given date, the [WeekOfYear] property indicates
// which week the date falls in, and [YearForWeekOfYear] provides the
// corresponding week-numbering year.
//
// You can use the week-numbering year when specifying a date with
// [NSDateComponents], usually in combination with the [WeekOfYear]. The code
// listing below shows this approach. It creates an [NSDateComponents]
// instance specifying the first Friday (weekday 6) of the first week of 2016,
// which started on a Friday. Therefore, this date is January 1, 2016 in the
// Gregorian calendar. However, on the ISO 8601 calendar, the first week of
// 2016 begins on the following Monday. This means the first Friday in the
// first week of 2016 is January 8, 2016 on the ISO 8601 calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/yearForWeekOfYear
func (d NSDateComponents) YearForWeekOfYear() int {
	rv := objc.Send[int](d.ID, objc.Sel("yearForWeekOfYear"))
	return rv
}
func (d NSDateComponents) SetYearForWeekOfYear(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setYearForWeekOfYear:"), value)
}

// The number of quarters.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/quarter
func (d NSDateComponents) Quarter() int {
	rv := objc.Send[int](d.ID, objc.Sel("quarter"))
	return rv
}
func (d NSDateComponents) SetQuarter(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setQuarter:"), value)
}

// The number of months.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/month
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Month() int {
	rv := objc.Send[int](d.ID, objc.Sel("month"))
	return rv
}
func (d NSDateComponents) SetMonth(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setMonth:"), value)
}

// A Boolean value that indicates whether the month is a leap month.
//
// # Discussion
//
// true if the month is a leap month, false otherwise.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/isLeapMonth
func (d NSDateComponents) LeapMonth() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isLeapMonth"))
	return rv
}
func (d NSDateComponents) SetLeapMonth(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setLeapMonth:"), value)
}

// The number of the weekdays.
//
// # Discussion
//
// Weekday units are the numbers 1 through , where is the number of days in
// the week. For example, in the Gregorian calendar, is 7 and Sunday is
// represented by 1.
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekday
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Weekday() int {
	rv := objc.Send[int](d.ID, objc.Sel("weekday"))
	return rv
}
func (d NSDateComponents) SetWeekday(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setWeekday:"), value)
}

// The ordinal number of weekdays.
//
// # Discussion
//
// Weekday ordinal units represent the position of the weekday within the next
// larger calendar unit, such as the month. For example, is the weekday
// ordinal unit for the Friday of the month.
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekdayOrdinal
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) WeekdayOrdinal() int {
	rv := objc.Send[int](d.ID, objc.Sel("weekdayOrdinal"))
	return rv
}
func (d NSDateComponents) SetWeekdayOrdinal(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setWeekdayOrdinal:"), value)
}

// The week number of the months.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekOfMonth
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) WeekOfMonth() int {
	rv := objc.Send[int](d.ID, objc.Sel("weekOfMonth"))
	return rv
}
func (d NSDateComponents) SetWeekOfMonth(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setWeekOfMonth:"), value)
}

// The ISO 8601 week date of the year.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/weekOfYear
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) WeekOfYear() int {
	rv := objc.Send[int](d.ID, objc.Sel("weekOfYear"))
	return rv
}
func (d NSDateComponents) SetWeekOfYear(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setWeekOfYear:"), value)
}

// The number of days.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/day
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Day() int {
	rv := objc.Send[int](d.ID, objc.Sel("day"))
	return rv
}
func (d NSDateComponents) SetDay(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setDay:"), value)
}

// The number of hour units for the receiver.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/hour
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Hour() int {
	rv := objc.Send[int](d.ID, objc.Sel("hour"))
	return rv
}
func (d NSDateComponents) SetHour(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setHour:"), value)
}

// The number of minute units for the receiver.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/minute
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Minute() int {
	rv := objc.Send[int](d.ID, objc.Sel("minute"))
	return rv
}
func (d NSDateComponents) SetMinute(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setMinute:"), value)
}

// The number of second units for the receiver.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/second
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Second() int {
	rv := objc.Send[int](d.ID, objc.Sel("second"))
	return rv
}
func (d NSDateComponents) SetSecond(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setSecond:"), value)
}

// The number of nanosecond units for the receiver.
//
// # Discussion
//
// This value is interpreted in the context of the calendar with which it is
// used—see [Calendars, Date Components, and Calendar Units] in [Date and
// Time Programming Guide].
//
// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/nanosecond
//
// [Calendars, Date Components, and Calendar Units]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/Articles/dtCalendars.html#//apple_ref/doc/uid/TP40003470
// [Date and Time Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DatesAndTimes/DatesAndTimes.html#//apple_ref/doc/uid/10000039i
func (d NSDateComponents) Nanosecond() int {
	rv := objc.Send[int](d.ID, objc.Sel("nanosecond"))
	return rv
}
func (d NSDateComponents) SetNanosecond(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setNanosecond:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/dayOfYear
func (d NSDateComponents) DayOfYear() int {
	rv := objc.Send[int](d.ID, objc.Sel("dayOfYear"))
	return rv
}
func (d NSDateComponents) SetDayOfYear(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setDayOfYear:"), value)
}

// See: https://developer.apple.com/documentation/Foundation/NSDateComponents/isRepeatedDay
func (d NSDateComponents) RepeatedDay() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isRepeatedDay"))
	return rv
}
func (d NSDateComponents) SetRepeatedDay(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setRepeatedDay:"), value)
}

// Specifies a date component without a value.
//
// See: https://developer.apple.com/documentation/foundation/nsdatecomponentundefined
func (d NSDateComponents) NSDateComponentUndefined() int {
	rv := objc.Send[int](d.ID, objc.Sel("NSDateComponentUndefined"))
	return rv
}
func (d NSDateComponents) SetNSDateComponentUndefined(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setNSDateComponentUndefined:"), value)
}

// Protocol methods for NSCopying

// Protocol methods for NSSecureCoding
