// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCalendar] class.
var (
	_NSCalendarClass     NSCalendarClass
	_NSCalendarClassOnce sync.Once
)

func getNSCalendarClass() NSCalendarClass {
	_NSCalendarClassOnce.Do(func() {
		_NSCalendarClass = NSCalendarClass{class: objc.GetClass("NSCalendar")}
	})
	return _NSCalendarClass
}

// GetNSCalendarClass returns the class object for NSCalendar.
func GetNSCalendarClass() NSCalendarClass {
	return getNSCalendarClass()
}

type NSCalendarClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSCalendarClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCalendarClass) Alloc() NSCalendar {
	rv := objc.Send[NSCalendar](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A definition of the relationships between calendar units and absolute
// points in time, providing features for calculation and comparison of dates.
//
// # Overview
// 
// In Swift, this object bridges to [Calendar]; use [NSCalendar] when you need
// reference semantics or other Foundation-specific behavior.
// 
// [NSCalendar] objects encapsulate information about systems of reckoning
// time in which the beginning, length, and divisions of a year are defined.
// They provide information about the calendar and support for calendrical
// computations such as determining the range of a given calendrical unit and
// adding units to a given absolute time.
// 
// [NSCalendar] is with its Core Foundation counterpart, [CFCalendar]. See
// [Toll-Free Bridging] for more information on toll-free bridging.
// 
// # Locales and Calendars
// 
// Most locales use the most widely used civil calendar, called the
// ([gregorian]), but there remain exceptions to this trend. For example:
// 
// - In Saudi Arabia, some locales use primarily the Islamic Umm al-Qura
// calendar ([islamicUmmAlQura]). - In Ethiopia, some locales use primarily
// the Ethiopian calendar ([ethiopicAmeteMihret] or [ethiopicAmeteAlem]). - In
// Iran and Afghanistan, some locales use primarily the Persian calendar
// ([persian]). - In Thailand, some locales use primarily the Buddhist
// calendar ([buddhist]).
// 
// Other locales use another calendar alongside the Gregorian calendar. For
// example:
// 
// - India also uses the Indian national calendar ([indian]). - Israel also
// uses the Hebrew calendar ([hebrew]). - China mainland and other regions
// also use the Chinese calendar ([chinese]), primarily to calculate
// astronomical date and Chinese traditional holidays. - Japan also uses the
// Japanese calendar ([japanese]), primarily to add year names.
// 
// Independent of any particular locale, certain calendars are used primarily
// to calculate dates for religious observances. Among these are the Buddhist
// ([buddhist]), Coptic ([coptic]), Hebrew ([hebrew]), and Islamic ([islamic])
// calendars.
// 
// # How NSCalendar Models the Gregorian Calendar
// 
// The Gregorian calendar was first introduced in 1582, as a replacement for
// the Julian Calendar. According to the Julian calendar, a leap day is added
// to February for any year with a number divisible by 4, which results in an
// annual disparity of 11 minutes, or 1 day every 128 years. The Gregorian
// calendar revised the rules for leap day calculation, by skipping the leap
// day for any year with a number divisible by 100, unless that year number is
// also divisible by 400, resulting in an annual disparity of only 26 seconds,
// or 1 day every 3323 years.
// 
// To transition from the Julian calendar to the Gregorian calendar, 10 days
// were dropped from the Gregorian calendar (October 5–14).
// 
// After the Gregorian calendar was introduced, many regions continued to use
// the Julian calendar, with Turkey being the last country or region to adopt
// the Gregorian calendar, in 1926. As a result of the staggered adoption, the
// transition period for regions at the time of adoption have different start
// dates and a different number of skipped days to account for the additional
// disparity from leap day calculations.
// 
// [NSCalendar] models the behavior of a Gregorian calendar (), which extends
// the Gregorian calendar backward in time from the date of its introduction.
// This behavior should be taken into account when working with dates created
// before the transition period of the affected locales.
// 
// # Calendar Arithmetic
// 
// To do calendar arithmetic, you use [NSDate] objects in conjunction with a
// calendar. For example, to convert between a decomposed date in one calendar
// and another calendar, you must first convert the decomposed elements into a
// date using the first calendar, then decompose it using the second. [NSDate]
// provides the absolute scale and epoch (reference point) for dates and
// times, which can then be rendered into a particular calendar, for
// calendrical computations or user display.
// 
// Two [NSCalendar] methods that return a date object, [NSCalendar.DateFromComponents],
// [NSCalendar.DateByAddingComponentsToDateOptions], take as a parameter an
// [NSDateComponents] object that describes the calendrical components
// required for the computation. You can provide as many components as you
// need (or choose to). When there is incomplete information to compute an
// absolute time, default values similar to `0` and `1` are usually chosen by
// a calendar, but this is a calendar-specific choice. If you provide
// inconsistent information, calendar-specific disambiguation is performed
// (which may involve ignoring one or more of the parameters). Related methods
// ([NSCalendar.ComponentsFromDate] and [NSCalendar.ComponentsFromDateToDateOptions]) take a bit
// mask parameter that specifies which components to calculate when returning
// an [NSDateComponents] object. The bit mask is composed of [NSCalendar.Unit]
// constants (see [Constants]).
// 
// In a calendar, day, week, weekday, month, and year numbers are generally
// 1-based, but there may be calendar-specific exceptions. Ordinal numbers,
// where they occur, are 1-based. Some calendars represented by this API may
// have to map their basic unit concepts into year/month/week/day/…
// nomenclature. For example, a calendar composed of 4 quarters in a year
// instead of 12 months uses the month unit to represent quarters. The
// particular values of the unit are defined by each calendar, and are not
// necessarily consistent with values for that unit in another calendar.
//
// [CFCalendar]: https://developer.apple.com/documentation/CoreFoundation/CFCalendar
// [Calendar]: https://developer.apple.com/documentation/Foundation/Calendar
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [buddhist]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/buddhist
// [chinese]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/chinese
// [coptic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/coptic
// [ethiopicAmeteAlem]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/ethiopicAmeteAlem
// [ethiopicAmeteMihret]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/ethiopicAmeteMihret
// [gregorian]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/gregorian
// [hebrew]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/hebrew
// [indian]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/indian
// [islamicUmmAlQura]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/islamicUmmAlQura
// [islamic]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/islamic
// [japanese]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/japanese
// [persian]: https://developer.apple.com/documentation/Foundation/NSCalendar/Identifier/persian
//
// # Creating and Initializing Calendars
//
//   - [NSCalendar.InitWithCalendarIdentifier]: Initializes a calendar according to a given identifier.
//
// # Extracting Components
//
//   - [NSCalendar.DateMatchesComponents]: Returns whether a given date matches all of the given date components.
//   - [NSCalendar.ComponentFromDate]: Returns the specified date component from a given date.
//   - [NSCalendar.ComponentsFromDate]: Returns the date components representing a given date.
//   - [NSCalendar.ComponentsFromDateToDateOptions]: Returns the difference between two supplied dates as date components.
//   - [NSCalendar.ComponentsFromDateComponentsToDateComponentsOptions]: Returns the difference between start and end dates given as date components.
//   - [NSCalendar.ComponentsInTimeZoneFromDate]: Returns all the date components of a date, as if in a given time zone (instead of the receiving calendar’s time zone).
//   - [NSCalendar.GetEraYearMonthDayFromDate]: Returns by reference the era, year, week of year, and weekday component values for a given date.
//   - [NSCalendar.GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate]: Returns by reference the era, year, week of year, and weekday component values for a given date.
//   - [NSCalendar.GetHourMinuteSecondNanosecondFromDate]: Returns by reference the hour, minute, second, and nanosecond component values for a given date.
//
// # Getting Calendar Information
//
//   - [NSCalendar.CalendarIdentifier]: An identifier for the calendar.
//   - [NSCalendar.FirstWeekday]: The index of the first weekday of the receiver.
//   - [NSCalendar.SetFirstWeekday]
//   - [NSCalendar.Locale]: The locale of the receiver.
//   - [NSCalendar.SetLocale]
//   - [NSCalendar.TimeZone]: The time zone for the calendar.
//   - [NSCalendar.SetTimeZone]
//   - [NSCalendar.MaximumRangeOfUnit]: Returns the maximum range limits of the values that a given unit can take on.
//   - [NSCalendar.MinimumRangeOfUnit]: Returns the minimum range limits of the values that a given unit can take on.
//   - [NSCalendar.MinimumDaysInFirstWeek]: The minimum number of days in the first week of the receiver.
//   - [NSCalendar.SetMinimumDaysInFirstWeek]
//   - [NSCalendar.OrdinalityOfUnitInUnitForDate]: Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
//   - [NSCalendar.RangeOfUnitInUnitForDate]: Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
//   - [NSCalendar.RangeOfUnitStartDateIntervalForDate]: Returns by reference the starting time and duration of a given calendar unit that contains a given date.
//   - [NSCalendar.RangeOfWeekendStartDateIntervalContainingDate]: Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
//
// # Scanning Dates
//
//   - [NSCalendar.StartOfDayForDate]: Returns the first moment of a given date as a date instance.
//   - [NSCalendar.EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock]: Computes the dates that match (or most closely match) a given set of components, and calls the block once for each of them, until the enumeration is stopped.
//   - [NSCalendar.NextDateAfterDateMatchingComponentsOptions]: Returns the next date after a given date matching the given components.
//   - [NSCalendar.NextDateAfterDateMatchingHourMinuteSecondOptions]: Returns the next date after a given date that matches the given hour, minute, and second, component values.
//   - [NSCalendar.NextDateAfterDateMatchingUnitValueOptions]: Returns the next date after a given date matching the given calendar unit value.
//
// # Calculating Dates
//
//   - [NSCalendar.DateFromComponents]: Returns a date representing the absolute time calculated from given components.
//   - [NSCalendar.DateByAddingComponentsToDateOptions]: Returns a date representing the absolute time calculated by adding given components to a given date.
//   - [NSCalendar.DateByAddingUnitValueToDateOptions]: Returns a date representing the absolute time calculated by adding the value of a given component to a given date.
//   - [NSCalendar.DateBySettingHourMinuteSecondOfDateOptions]: Creates a new date calculated with the given time.
//   - [NSCalendar.DateBySettingUnitValueOfDateOptions]: Returns a new date representing the date calculated by setting a specific component of a given date to a given value, while trying to keep lower components the same.
//   - [NSCalendar.DateWithEraYearMonthDayHourMinuteSecondNanosecond]: Returns a date created with the given components.
//   - [NSCalendar.DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond]: Returns a new date created with the given components base on a week-of-year value.
//   - [NSCalendar.NextWeekendStartDateIntervalOptionsAfterDate]: Returns by reference the starting date and time interval range of the next weekend period after a given date.
//
// # Comparing Dates
//
//   - [NSCalendar.CompareDateToDateToUnitGranularity]: Indicates the ordering of two given dates based on their components down to a given unit granularity.
//   - [NSCalendar.IsDateEqualToDateToUnitGranularity]: Indicates whether two dates are equal to a given unit of granularity.
//   - [NSCalendar.IsDateInSameDayAsDate]: Indicates whether two dates are in the same day.
//   - [NSCalendar.IsDateInToday]: Indicates whether the given date is in “today.”
//   - [NSCalendar.IsDateInTomorrow]: Indicates whether the given date is in “tomorrow.”
//   - [NSCalendar.IsDateInWeekend]: Indicates whether a given date falls within a weekend period, as defined by the calendar and the calendar’s locale.
//   - [NSCalendar.IsDateInYesterday]: Indicates whether the given date is in “yesterday.”
//
// # Getting AM and PM Symbols
//
//   - [NSCalendar.AMSymbol]: The symbol used to represent “AM” for this calendar.
//   - [NSCalendar.PMSymbol]: The symbol used to represent “PM” for this calendar.
//
// # Getting Weekday Symbols
//
//   - [NSCalendar.WeekdaySymbols]: A list of weekdays in this calendar.
//   - [NSCalendar.ShortWeekdaySymbols]: A list of shorter-named weekdays in this calendar.
//   - [NSCalendar.VeryShortWeekdaySymbols]: A list of very-shortly-named weekdays in this calendar.
//   - [NSCalendar.StandaloneWeekdaySymbols]: A list of standalone weekday symbols for this calendar.
//   - [NSCalendar.ShortStandaloneWeekdaySymbols]: A list of short standalone weekday symbols for this calendar.
//   - [NSCalendar.VeryShortStandaloneWeekdaySymbols]: A list of very short standalone weekday symbols for this calendar.
//
// # Getting Month Symbols
//
//   - [NSCalendar.MonthSymbols]: A list of month symbols for this calendar.
//   - [NSCalendar.ShortMonthSymbols]: A list of short month symbols for this calendar.
//   - [NSCalendar.VeryShortMonthSymbols]: A list of very short month symbols for this calendar.
//   - [NSCalendar.StandaloneMonthSymbols]: A list of standalone month symbols for this calendar.
//   - [NSCalendar.ShortStandaloneMonthSymbols]: A list of short standalone month symbols for this calendar.
//   - [NSCalendar.VeryShortStandaloneMonthSymbols]: A list of very short month symbols for this calendar.
//
// # Getting Quarter Symbols
//
//   - [NSCalendar.QuarterSymbols]: A list of quarter symbols for this calendar.
//   - [NSCalendar.ShortQuarterSymbols]: A list of short quarter symbols for this calendar.
//   - [NSCalendar.StandaloneQuarterSymbols]: A list of standalone quarter symbols for this calendar.
//   - [NSCalendar.ShortStandaloneQuarterSymbols]: A list of short standalone quarter symbols for this calendar.
//
// # Getting Era Symbols
//
//   - [NSCalendar.EraSymbols]: A list of era symbols for this calendar.
//   - [NSCalendar.LongEraSymbols]: A list of long era symbols for this calendar.
//
// # Recognizing Notifications
//
//   - [NSCalendar.NSCalendarDayChanged]: A notification that is posted whenever the calendar day of the system changes, as determined by the system calendar, locale, and time zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar
type NSCalendar struct {
	objectivec.Object
}

// NSCalendarFromID constructs a [NSCalendar] from an objc.ID.
//
// A definition of the relationships between calendar units and absolute
// points in time, providing features for calculation and comparison of dates.
func NSCalendarFromID(id objc.ID) NSCalendar {
	return NSCalendar{objectivec.Object{ID: id}}
}
// NOTE: NSCalendar adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCalendar] class.
//
// # Creating and Initializing Calendars
//
//   - [INSCalendar.InitWithCalendarIdentifier]: Initializes a calendar according to a given identifier.
//
// # Extracting Components
//
//   - [INSCalendar.DateMatchesComponents]: Returns whether a given date matches all of the given date components.
//   - [INSCalendar.ComponentFromDate]: Returns the specified date component from a given date.
//   - [INSCalendar.ComponentsFromDate]: Returns the date components representing a given date.
//   - [INSCalendar.ComponentsFromDateToDateOptions]: Returns the difference between two supplied dates as date components.
//   - [INSCalendar.ComponentsFromDateComponentsToDateComponentsOptions]: Returns the difference between start and end dates given as date components.
//   - [INSCalendar.ComponentsInTimeZoneFromDate]: Returns all the date components of a date, as if in a given time zone (instead of the receiving calendar’s time zone).
//   - [INSCalendar.GetEraYearMonthDayFromDate]: Returns by reference the era, year, week of year, and weekday component values for a given date.
//   - [INSCalendar.GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate]: Returns by reference the era, year, week of year, and weekday component values for a given date.
//   - [INSCalendar.GetHourMinuteSecondNanosecondFromDate]: Returns by reference the hour, minute, second, and nanosecond component values for a given date.
//
// # Getting Calendar Information
//
//   - [INSCalendar.CalendarIdentifier]: An identifier for the calendar.
//   - [INSCalendar.FirstWeekday]: The index of the first weekday of the receiver.
//   - [INSCalendar.SetFirstWeekday]
//   - [INSCalendar.Locale]: The locale of the receiver.
//   - [INSCalendar.SetLocale]
//   - [INSCalendar.TimeZone]: The time zone for the calendar.
//   - [INSCalendar.SetTimeZone]
//   - [INSCalendar.MaximumRangeOfUnit]: Returns the maximum range limits of the values that a given unit can take on.
//   - [INSCalendar.MinimumRangeOfUnit]: Returns the minimum range limits of the values that a given unit can take on.
//   - [INSCalendar.MinimumDaysInFirstWeek]: The minimum number of days in the first week of the receiver.
//   - [INSCalendar.SetMinimumDaysInFirstWeek]
//   - [INSCalendar.OrdinalityOfUnitInUnitForDate]: Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
//   - [INSCalendar.RangeOfUnitInUnitForDate]: Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
//   - [INSCalendar.RangeOfUnitStartDateIntervalForDate]: Returns by reference the starting time and duration of a given calendar unit that contains a given date.
//   - [INSCalendar.RangeOfWeekendStartDateIntervalContainingDate]: Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
//
// # Scanning Dates
//
//   - [INSCalendar.StartOfDayForDate]: Returns the first moment of a given date as a date instance.
//   - [INSCalendar.EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock]: Computes the dates that match (or most closely match) a given set of components, and calls the block once for each of them, until the enumeration is stopped.
//   - [INSCalendar.NextDateAfterDateMatchingComponentsOptions]: Returns the next date after a given date matching the given components.
//   - [INSCalendar.NextDateAfterDateMatchingHourMinuteSecondOptions]: Returns the next date after a given date that matches the given hour, minute, and second, component values.
//   - [INSCalendar.NextDateAfterDateMatchingUnitValueOptions]: Returns the next date after a given date matching the given calendar unit value.
//
// # Calculating Dates
//
//   - [INSCalendar.DateFromComponents]: Returns a date representing the absolute time calculated from given components.
//   - [INSCalendar.DateByAddingComponentsToDateOptions]: Returns a date representing the absolute time calculated by adding given components to a given date.
//   - [INSCalendar.DateByAddingUnitValueToDateOptions]: Returns a date representing the absolute time calculated by adding the value of a given component to a given date.
//   - [INSCalendar.DateBySettingHourMinuteSecondOfDateOptions]: Creates a new date calculated with the given time.
//   - [INSCalendar.DateBySettingUnitValueOfDateOptions]: Returns a new date representing the date calculated by setting a specific component of a given date to a given value, while trying to keep lower components the same.
//   - [INSCalendar.DateWithEraYearMonthDayHourMinuteSecondNanosecond]: Returns a date created with the given components.
//   - [INSCalendar.DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond]: Returns a new date created with the given components base on a week-of-year value.
//   - [INSCalendar.NextWeekendStartDateIntervalOptionsAfterDate]: Returns by reference the starting date and time interval range of the next weekend period after a given date.
//
// # Comparing Dates
//
//   - [INSCalendar.CompareDateToDateToUnitGranularity]: Indicates the ordering of two given dates based on their components down to a given unit granularity.
//   - [INSCalendar.IsDateEqualToDateToUnitGranularity]: Indicates whether two dates are equal to a given unit of granularity.
//   - [INSCalendar.IsDateInSameDayAsDate]: Indicates whether two dates are in the same day.
//   - [INSCalendar.IsDateInToday]: Indicates whether the given date is in “today.”
//   - [INSCalendar.IsDateInTomorrow]: Indicates whether the given date is in “tomorrow.”
//   - [INSCalendar.IsDateInWeekend]: Indicates whether a given date falls within a weekend period, as defined by the calendar and the calendar’s locale.
//   - [INSCalendar.IsDateInYesterday]: Indicates whether the given date is in “yesterday.”
//
// # Getting AM and PM Symbols
//
//   - [INSCalendar.AMSymbol]: The symbol used to represent “AM” for this calendar.
//   - [INSCalendar.PMSymbol]: The symbol used to represent “PM” for this calendar.
//
// # Getting Weekday Symbols
//
//   - [INSCalendar.WeekdaySymbols]: A list of weekdays in this calendar.
//   - [INSCalendar.ShortWeekdaySymbols]: A list of shorter-named weekdays in this calendar.
//   - [INSCalendar.VeryShortWeekdaySymbols]: A list of very-shortly-named weekdays in this calendar.
//   - [INSCalendar.StandaloneWeekdaySymbols]: A list of standalone weekday symbols for this calendar.
//   - [INSCalendar.ShortStandaloneWeekdaySymbols]: A list of short standalone weekday symbols for this calendar.
//   - [INSCalendar.VeryShortStandaloneWeekdaySymbols]: A list of very short standalone weekday symbols for this calendar.
//
// # Getting Month Symbols
//
//   - [INSCalendar.MonthSymbols]: A list of month symbols for this calendar.
//   - [INSCalendar.ShortMonthSymbols]: A list of short month symbols for this calendar.
//   - [INSCalendar.VeryShortMonthSymbols]: A list of very short month symbols for this calendar.
//   - [INSCalendar.StandaloneMonthSymbols]: A list of standalone month symbols for this calendar.
//   - [INSCalendar.ShortStandaloneMonthSymbols]: A list of short standalone month symbols for this calendar.
//   - [INSCalendar.VeryShortStandaloneMonthSymbols]: A list of very short month symbols for this calendar.
//
// # Getting Quarter Symbols
//
//   - [INSCalendar.QuarterSymbols]: A list of quarter symbols for this calendar.
//   - [INSCalendar.ShortQuarterSymbols]: A list of short quarter symbols for this calendar.
//   - [INSCalendar.StandaloneQuarterSymbols]: A list of standalone quarter symbols for this calendar.
//   - [INSCalendar.ShortStandaloneQuarterSymbols]: A list of short standalone quarter symbols for this calendar.
//
// # Getting Era Symbols
//
//   - [INSCalendar.EraSymbols]: A list of era symbols for this calendar.
//   - [INSCalendar.LongEraSymbols]: A list of long era symbols for this calendar.
//
// # Recognizing Notifications
//
//   - [INSCalendar.NSCalendarDayChanged]: A notification that is posted whenever the calendar day of the system changes, as determined by the system calendar, locale, and time zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar
type INSCalendar interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating and Initializing Calendars

	// Initializes a calendar according to a given identifier.
	InitWithCalendarIdentifier(ident NSCalendarIdentifier) NSCalendar

	// Topic: Extracting Components

	// Returns whether a given date matches all of the given date components.
	DateMatchesComponents(date INSDate, components INSDateComponents) bool
	// Returns the specified date component from a given date.
	ComponentFromDate(unit NSCalendarUnit, date INSDate) int
	// Returns the date components representing a given date.
	ComponentsFromDate(unitFlags NSCalendarUnit, date INSDate) INSDateComponents
	// Returns the difference between two supplied dates as date components.
	ComponentsFromDateToDateOptions(unitFlags NSCalendarUnit, startingDate INSDate, resultDate INSDate, opts NSCalendarOptions) INSDateComponents
	// Returns the difference between start and end dates given as date components.
	ComponentsFromDateComponentsToDateComponentsOptions(unitFlags NSCalendarUnit, startingDateComp INSDateComponents, resultDateComp INSDateComponents, options NSCalendarOptions) INSDateComponents
	// Returns all the date components of a date, as if in a given time zone (instead of the receiving calendar’s time zone).
	ComponentsInTimeZoneFromDate(timezone INSTimeZone, date INSDate) INSDateComponents
	// Returns by reference the era, year, week of year, and weekday component values for a given date.
	GetEraYearMonthDayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, monthValuePointer unsafe.Pointer, dayValuePointer unsafe.Pointer, date INSDate)
	// Returns by reference the era, year, week of year, and weekday component values for a given date.
	GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, weekValuePointer unsafe.Pointer, weekdayValuePointer unsafe.Pointer, date INSDate)
	// Returns by reference the hour, minute, second, and nanosecond component values for a given date.
	GetHourMinuteSecondNanosecondFromDate(hourValuePointer unsafe.Pointer, minuteValuePointer unsafe.Pointer, secondValuePointer unsafe.Pointer, nanosecondValuePointer unsafe.Pointer, date INSDate)

	// Topic: Getting Calendar Information

	// An identifier for the calendar.
	CalendarIdentifier() NSCalendarIdentifier
	// The index of the first weekday of the receiver.
	FirstWeekday() uint
	SetFirstWeekday(value uint)
	// The locale of the receiver.
	Locale() INSLocale
	SetLocale(value INSLocale)
	// The time zone for the calendar.
	TimeZone() INSTimeZone
	SetTimeZone(value INSTimeZone)
	// Returns the maximum range limits of the values that a given unit can take on.
	MaximumRangeOfUnit(unit NSCalendarUnit) NSRange
	// Returns the minimum range limits of the values that a given unit can take on.
	MinimumRangeOfUnit(unit NSCalendarUnit) NSRange
	// The minimum number of days in the first week of the receiver.
	MinimumDaysInFirstWeek() uint
	SetMinimumDaysInFirstWeek(value uint)
	// Returns, for a given absolute time, the ordinal number of a smaller calendar unit (such as a day) within a specified larger calendar unit (such as a week).
	OrdinalityOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date INSDate) uint
	// Returns the range of absolute time values that a smaller calendar unit (such as a day) can take on in a larger calendar unit (such as a month) that includes a specified absolute time.
	RangeOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date INSDate) NSRange
	// Returns by reference the starting time and duration of a given calendar unit that contains a given date.
	RangeOfUnitStartDateIntervalForDate(unit NSCalendarUnit, datep INSDate, tip unsafe.Pointer, date INSDate) bool
	// Returns whether a given date falls within a weekend period, and if so, returns by reference the start date and time interval of the weekend range.
	RangeOfWeekendStartDateIntervalContainingDate(datep INSDate, tip unsafe.Pointer, date INSDate) bool

	// Topic: Scanning Dates

	// Returns the first moment of a given date as a date instance.
	StartOfDayForDate(date INSDate) INSDate
	// Computes the dates that match (or most closely match) a given set of components, and calls the block once for each of them, until the enumeration is stopped.
	EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start INSDate, comps INSDateComponents, opts NSCalendarOptions, block DateHandler)
	// Returns the next date after a given date matching the given components.
	NextDateAfterDateMatchingComponentsOptions(date INSDate, comps INSDateComponents, options NSCalendarOptions) INSDate
	// Returns the next date after a given date that matches the given hour, minute, and second, component values.
	NextDateAfterDateMatchingHourMinuteSecondOptions(date INSDate, hourValue int, minuteValue int, secondValue int, options NSCalendarOptions) INSDate
	// Returns the next date after a given date matching the given calendar unit value.
	NextDateAfterDateMatchingUnitValueOptions(date INSDate, unit NSCalendarUnit, value int, options NSCalendarOptions) INSDate

	// Topic: Calculating Dates

	// Returns a date representing the absolute time calculated from given components.
	DateFromComponents(comps INSDateComponents) INSDate
	// Returns a date representing the absolute time calculated by adding given components to a given date.
	DateByAddingComponentsToDateOptions(comps INSDateComponents, date INSDate, opts NSCalendarOptions) INSDate
	// Returns a date representing the absolute time calculated by adding the value of a given component to a given date.
	DateByAddingUnitValueToDateOptions(unit NSCalendarUnit, value int, date INSDate, options NSCalendarOptions) INSDate
	// Creates a new date calculated with the given time.
	DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date INSDate, opts NSCalendarOptions) INSDate
	// Returns a new date representing the date calculated by setting a specific component of a given date to a given value, while trying to keep lower components the same.
	DateBySettingUnitValueOfDateOptions(unit NSCalendarUnit, v int, date INSDate, opts NSCalendarOptions) INSDate
	// Returns a date created with the given components.
	DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) INSDate
	// Returns a new date created with the given components base on a week-of-year value.
	DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) INSDate
	// Returns by reference the starting date and time interval range of the next weekend period after a given date.
	NextWeekendStartDateIntervalOptionsAfterDate(datep INSDate, tip unsafe.Pointer, options NSCalendarOptions, date INSDate) bool

	// Topic: Comparing Dates

	// Indicates the ordering of two given dates based on their components down to a given unit granularity.
	CompareDateToDateToUnitGranularity(date1 INSDate, date2 INSDate, unit NSCalendarUnit) ComparisonResult
	// Indicates whether two dates are equal to a given unit of granularity.
	IsDateEqualToDateToUnitGranularity(date1 INSDate, date2 INSDate, unit NSCalendarUnit) bool
	// Indicates whether two dates are in the same day.
	IsDateInSameDayAsDate(date1 INSDate, date2 INSDate) bool
	// Indicates whether the given date is in “today.”
	IsDateInToday(date INSDate) bool
	// Indicates whether the given date is in “tomorrow.”
	IsDateInTomorrow(date INSDate) bool
	// Indicates whether a given date falls within a weekend period, as defined by the calendar and the calendar’s locale.
	IsDateInWeekend(date INSDate) bool
	// Indicates whether the given date is in “yesterday.”
	IsDateInYesterday(date INSDate) bool

	// Topic: Getting AM and PM Symbols

	// The symbol used to represent “AM” for this calendar.
	AMSymbol() string
	// The symbol used to represent “PM” for this calendar.
	PMSymbol() string

	// Topic: Getting Weekday Symbols

	// A list of weekdays in this calendar.
	WeekdaySymbols() []string
	// A list of shorter-named weekdays in this calendar.
	ShortWeekdaySymbols() []string
	// A list of very-shortly-named weekdays in this calendar.
	VeryShortWeekdaySymbols() []string
	// A list of standalone weekday symbols for this calendar.
	StandaloneWeekdaySymbols() []string
	// A list of short standalone weekday symbols for this calendar.
	ShortStandaloneWeekdaySymbols() []string
	// A list of very short standalone weekday symbols for this calendar.
	VeryShortStandaloneWeekdaySymbols() []string

	// Topic: Getting Month Symbols

	// A list of month symbols for this calendar.
	MonthSymbols() []string
	// A list of short month symbols for this calendar.
	ShortMonthSymbols() []string
	// A list of very short month symbols for this calendar.
	VeryShortMonthSymbols() []string
	// A list of standalone month symbols for this calendar.
	StandaloneMonthSymbols() []string
	// A list of short standalone month symbols for this calendar.
	ShortStandaloneMonthSymbols() []string
	// A list of very short month symbols for this calendar.
	VeryShortStandaloneMonthSymbols() []string

	// Topic: Getting Quarter Symbols

	// A list of quarter symbols for this calendar.
	QuarterSymbols() []string
	// A list of short quarter symbols for this calendar.
	ShortQuarterSymbols() []string
	// A list of standalone quarter symbols for this calendar.
	StandaloneQuarterSymbols() []string
	// A list of short standalone quarter symbols for this calendar.
	ShortStandaloneQuarterSymbols() []string

	// Topic: Getting Era Symbols

	// A list of era symbols for this calendar.
	EraSymbols() []string
	// A list of long era symbols for this calendar.
	LongEraSymbols() []string

	// Topic: Recognizing Notifications

	// A notification that is posted whenever the calendar day of the system changes, as determined by the system calendar, locale, and time zone.
	NSCalendarDayChanged() NSNotificationName
}

// Init initializes the instance.
func (c NSCalendar) Init() NSCalendar {
	rv := objc.Send[NSCalendar](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCalendar) Autorelease() NSCalendar {
	rv := objc.Send[NSCalendar](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCalendar creates a new NSCalendar instance.
func NewNSCalendar() NSCalendar {
	class := getNSCalendarClass()
	rv := objc.Send[NSCalendar](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a calendar according to a given identifier.
//
// ident: The identifier for the new calendar. For valid identifiers, see `Calendar
// Identifiers`.
//
// # Return Value
// 
// The initialized calendar, or `nil` if the identifier is unknown (if, for
// example, it is either an unrecognized string or the calendar is not
// supported by the current version of the operating system).
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/init(calendarIdentifier:)
func NewCalendarWithCalendarIdentifier(ident NSCalendarIdentifier) NSCalendar {
	instance := getNSCalendarClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCalendarIdentifier:"), objc.String(string(ident)))
	return NSCalendarFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewCalendarWithCoder(coder INSCoder) NSCalendar {
	instance := getNSCalendarClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCalendarFromID(rv)
}

// Creates a new calendar specified by a given identifier.
//
// calendarIdentifierConstant: The identifier for the new calendar. For valid identifiers, see `Calendar
// Identifiers`.
//
// # Return Value
// 
// The initialized calendar, or `nil` if the identifier is unknown (if, for
// example, it is either an unrecognized string or the calendar is not
// supported by the current version of the operating system).
//
// # Discussion
// 
// The returned calendar defaults to the current locale and default time zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/init(identifier:)
func NewCalendarWithIdentifier(calendarIdentifierConstant NSCalendarIdentifier) NSCalendar {
	rv := objc.Send[objc.ID](objc.ID(getNSCalendarClass().class), objc.Sel("calendarWithIdentifier:"), objc.String(string(calendarIdentifierConstant)))
	return NSCalendarFromID(rv)
}

// Initializes a calendar according to a given identifier.
//
// ident: The identifier for the new calendar. For valid identifiers, see `Calendar
// Identifiers`.
//
// # Return Value
// 
// The initialized calendar, or `nil` if the identifier is unknown (if, for
// example, it is either an unrecognized string or the calendar is not
// supported by the current version of the operating system).
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/init(calendarIdentifier:)
func (c NSCalendar) InitWithCalendarIdentifier(ident NSCalendarIdentifier) NSCalendar {
	rv := objc.Send[NSCalendar](c.ID, objc.Sel("initWithCalendarIdentifier:"), objc.String(string(ident)))
	return rv
}
// Returns whether a given date matches all of the given date components.
//
// date: The date for which to perform the calculation.
//
// components: The date components to match.
//
// # Return Value
// 
// [true] if the given date matches the given components, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is useful for determining whether dates calculated by methods
// like [NextDateAfterDateMatchingUnitValueOptions] or
// [EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock] are
// exact, or required an adjustment due to a nonexistent time.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(_:matchesComponents:)
func (c NSCalendar) DateMatchesComponents(date INSDate, components INSDateComponents) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("date:matchesComponents:"), date, components)
	return rv
}
// Returns the specified date component from a given date.
//
// unit: The component to return. For possible values, see [NSCalendar.Unit].
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// An [NSInteger] value for the requested component.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/component(_:from:)
func (c NSCalendar) ComponentFromDate(unit NSCalendarUnit, date INSDate) int {
	rv := objc.Send[int](c.ID, objc.Sel("component:fromDate:"), unit, date)
	return rv
}
// Returns the date components representing a given date.
//
// unitFlags: The components into which to decompose `date`.
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// An [NSDateComponents] object containing `date` decomposed into the
// components specified by `unitFlags`. Returns `nil` if `date` falls outside
// of the defined range of the receiver or if the computation cannot be
// performed.
//
// # Discussion
// 
// The Weekday ordinality, when requested, refers to the next larger (than
// Week) of the requested units. Some computations can take a relatively long
// time.
// 
// The following example shows how to use this method to determine the current
// year, month, and day, using an existing calendar (`gregorian`):
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:)
func (c NSCalendar) ComponentsFromDate(unitFlags NSCalendarUnit, date INSDate) INSDateComponents {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("components:fromDate:"), unitFlags, date)
	return NSDateComponentsFromID(rv)
}
// Returns the difference between two supplied dates as date components.
//
// unitFlags: Specifies the components for the returned [NSDateComponents] object.
//
// startingDate: The start date for the calculation.
//
// resultDate: The end date for the calculation.
//
// opts: Options for the calculation. For possible values, see [NSCalendar.Options].
// 
// If you specify a “wrap” option ([CalendarWrapComponents]), the
// specified components are incremented and wrap around to zero/one on
// overflow, but do not cause higher units to be incremented. When the wrap
// option is not specified, overflow in a unit carries into the higher units,
// as in typical addition.
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// An [NSDateComponents] object whose components are specified by `unitFlags`
// and calculated from the difference between the `resultDate` and `startDate`
// using the options specified by `options`. Returns `nil` if either date
// falls outside the defined range of the receiver or if the computation
// cannot be performed.
//
// # Discussion
// 
// The result is lossy if there is not a small enough unit requested to hold
// the full precision of the difference. Some operations can be ambiguous, and
// the behavior of the computation is calendar-specific, but generally larger
// components will be computed before smaller components; for example, in the
// Gregorian calendar a result might be 1 month and 5 days instead of, for
// example, 0 months and 35 days. The resulting component values may be
// negative if `resultDate` is before `startDate`.
// 
// The following example shows how to get the approximate number of months and
// days between two dates using an existing calendar (`gregorian`):
// 
// Note that some computations can take a relatively long time.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:to:options:)-84y5w
func (c NSCalendar) ComponentsFromDateToDateOptions(unitFlags NSCalendarUnit, startingDate INSDate, resultDate INSDate, opts NSCalendarOptions) INSDateComponents {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("components:fromDate:toDate:options:"), unitFlags, startingDate, resultDate, opts)
	return NSDateComponentsFromID(rv)
}
// Returns the difference between start and end dates given as date
// components.
//
// unitFlags: Specifies the components for the returned [NSDateComponents] object.
//
// startingDateComp: The start date for the calculation as an [NSDateComponents] object.
//
// resultDateComp: The end date for the calculation as an [NSDateComponents] object.
//
// options: The `options` parameter is currently unused.
//
// # Return Value
// 
// An [NSDateComponents] object whose components are specified by `unitFlags`
// and calculated from the difference between the `startingDateComp` and
// `resultDateComp` using the options specified by `options`. Returns `nil` if
// either date falls outside the defined range of the receiver or if the
// computation cannot be performed.
//
// # Discussion
// 
// If an [NSDateComponents] object does not specify a value for a calendar
// unit required to determine an absolute date, the base value of that unit is
// assumed. For example, given an [NSDateComponents] object with only a `year`
// and a `month` specified, the resulting [NSDate] object would be constructed
// using a `day` value of `1` and `hour`, `minute`, `second` and `nanosecond`
// values of `0`. Passing an [NSDateComponents] argument with an unspecified
// `era` or `year` value is not advised.
// 
// If an [NSDateComponents] object’s `timeZone` property is set, the time
// zone property value will be used in the calculation. If an
// [NSDateComponents] object’s `calendar` property is set, the calendar
// property value will be used instead of the receiving calendar. If both an
// [NSDateComponents] object’s `timeZone` and `calendar` properties are set,
// the time zone property value overrides the time zone of the calendar
// property value.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/components(_:from:to:options:)-49lo8
func (c NSCalendar) ComponentsFromDateComponentsToDateComponentsOptions(unitFlags NSCalendarUnit, startingDateComp INSDateComponents, resultDateComp INSDateComponents, options NSCalendarOptions) INSDateComponents {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("components:fromDateComponents:toDateComponents:options:"), unitFlags, startingDateComp, resultDateComp, options)
	return NSDateComponentsFromID(rv)
}
// Returns all the date components of a date, as if in a given time zone
// (instead of the receiving calendar’s time zone).
//
// timezone: The time zone to use when returning the components. This value overrides
// the time zone of the receiving [NSCalendar].
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// An [NSDateComponents] object containing all the components from the given
// date, calculated using the given time zone.
//
// # Discussion
// 
// If you want “date information in a given time zone” for the purpose to
// displaying it, you should use [NSDateFormatter] to format the date.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/components(in:from:)
func (c NSCalendar) ComponentsInTimeZoneFromDate(timezone INSTimeZone, date INSDate) INSDateComponents {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("componentsInTimeZone:fromDate:"), timezone, date)
	return NSDateComponentsFromID(rv)
}
// Returns by reference the era, year, week of year, and weekday component
// values for a given date.
//
// eraValuePointer: Upon return, contains the era of the given date.
//
// yearValuePointer: Upon return, contains the year of the given date.
//
// monthValuePointer: Upon return, contains the month of the given date.
//
// dayValuePointer: Upon return, contains the day of the given date.
//
// date: The date for which to perform the calculation.
//
// # Discussion
// 
// Pass [NULL] to ignore any individual component parameter.
// 
// This is a convenience method for getting the time components of a given
// date using [ComponentsFromDate]
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/getEra(_:year:month:day:from:)
func (c NSCalendar) GetEraYearMonthDayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, monthValuePointer unsafe.Pointer, dayValuePointer unsafe.Pointer, date INSDate) {
	objc.Send[objc.ID](c.ID, objc.Sel("getEra:year:month:day:fromDate:"), eraValuePointer, yearValuePointer, monthValuePointer, dayValuePointer, date)
}
// Returns by reference the era, year, week of year, and weekday component
// values for a given date.
//
// eraValuePointer: Upon return, contains the era of the given date.
//
// yearValuePointer: Upon return, contains the year of the given date.
//
// weekValuePointer: Upon return, contains the week of the given date.
//
// weekdayValuePointer: Upon return, contains the weekday of the given date.
//
// date: The date for which to perform the calculation.
//
// # Discussion
// 
// Pass [NULL] to ignore any individual component parameter.
// 
// This is a convenience method for getting the time components of a given
// date using [ComponentsFromDate]
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/getEra(_:yearForWeekOfYear:weekOfYear:weekday:from:)
func (c NSCalendar) GetEraYearForWeekOfYearWeekOfYearWeekdayFromDate(eraValuePointer unsafe.Pointer, yearValuePointer unsafe.Pointer, weekValuePointer unsafe.Pointer, weekdayValuePointer unsafe.Pointer, date INSDate) {
	objc.Send[objc.ID](c.ID, objc.Sel("getEra:yearForWeekOfYear:weekOfYear:weekday:fromDate:"), eraValuePointer, yearValuePointer, weekValuePointer, weekdayValuePointer, date)
}
// Returns by reference the hour, minute, second, and nanosecond component
// values for a given date.
//
// hourValuePointer: Upon return, contains the hour of the given date.
//
// minuteValuePointer: Upon return, contains the minute of the given date.
//
// secondValuePointer: Upon return, contains the second of the given date.
//
// nanosecondValuePointer: Upon return, contains the nanosecond of the given date.
//
// date: The date for which to perform the calculation.
//
// # Discussion
// 
// Pass [NULL] to ignore any individual component parameter.
// 
// This is a convenience method for getting the time components of a given
// date using [ComponentsFromDate]
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/getHour(_:minute:second:nanosecond:from:)
func (c NSCalendar) GetHourMinuteSecondNanosecondFromDate(hourValuePointer unsafe.Pointer, minuteValuePointer unsafe.Pointer, secondValuePointer unsafe.Pointer, nanosecondValuePointer unsafe.Pointer, date INSDate) {
	objc.Send[objc.ID](c.ID, objc.Sel("getHour:minute:second:nanosecond:fromDate:"), hourValuePointer, minuteValuePointer, secondValuePointer, nanosecondValuePointer, date)
}
// Returns the maximum range limits of the values that a given unit can take
// on.
//
// unit: The unit for which the maximum range is returned.
//
// # Return Value
// 
// The maximum range limits of the values that the unit specified by `unit`
// can take on in the receiver.
//
// # Discussion
// 
// As an example, in the Gregorian calendar the maximum range of values for
// the Day unit is 1-31.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/maximumRange(of:)
func (c NSCalendar) MaximumRangeOfUnit(unit NSCalendarUnit) NSRange {
	rv := objc.Send[NSRange](c.ID, objc.Sel("maximumRangeOfUnit:"), unit)
	return NSRange(rv)
}
// Returns the minimum range limits of the values that a given unit can take
// on.
//
// unit: The unit for which the maximum range is returned.
//
// # Return Value
// 
// The minimum range limits of the values that the unit specified by `unit`
// can take on in the receiver.
//
// # Discussion
// 
// As an example, in the Gregorian calendar the minimum range of values for
// the Day unit is 1-28.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumRange(of:)
func (c NSCalendar) MinimumRangeOfUnit(unit NSCalendarUnit) NSRange {
	rv := objc.Send[NSRange](c.ID, objc.Sel("minimumRangeOfUnit:"), unit)
	return NSRange(rv)
}
// Returns, for a given absolute time, the ordinal number of a smaller
// calendar unit (such as a day) within a specified larger calendar unit (such
// as a week).
//
// smaller: The smaller calendar unit
//
// larger: The larger calendar unit
//
// date: The absolute time for which the calculation is performed
//
// # Return Value
// 
// The ordinal number of `smaller` within `larger` at the time specified by
// `date`. Returns [NSNotFound] if `larger` is not logically bigger than
// `smaller` in the calendar, or the given combination of units does not make
// sense (or is a computation which is undefined).
//
// # Discussion
// 
// The ordinality is in most cases not the same as the decomposed value of the
// unit. Typically return values are `1` and greater. For example, the time
// `00:45` is in the first hour of the day, and for units Hour and Day
// respectively, the result would be `1`. An exception is the week-in-month
// calculation, which returns `0` for days before the first week in the month
// containing the date.
// 
// Note that some computations can take a relatively long time.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/ordinality(of:in:for:)
func (c NSCalendar) OrdinalityOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date INSDate) uint {
	rv := objc.Send[uint](c.ID, objc.Sel("ordinalityOfUnit:inUnit:forDate:"), smaller, larger, date)
	return rv
}
// Returns the range of absolute time values that a smaller calendar unit
// (such as a day) can take on in a larger calendar unit (such as a month)
// that includes a specified absolute time.
//
// smaller: The smaller calendar unit.
//
// larger: The larger calendar unit.
//
// date: The absolute time for which the calculation is performed.
//
// # Return Value
// 
// The range of absolute time values `smaller` can take on in `larger` at the
// time specified by `date`. Returns `{NSNotFound, NSNotFound`} if `larger` is
// not logically bigger than `smaller` in the calendar, or the given
// combination of units does not make sense (or is a computation which is
// undefined).
//
// # Discussion
// 
// You can use this method to calculate, for example, the range the Day unit
// can take on in the Month in which `date` lies.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:in:for:)
func (c NSCalendar) RangeOfUnitInUnitForDate(smaller NSCalendarUnit, larger NSCalendarUnit, date INSDate) NSRange {
	rv := objc.Send[NSRange](c.ID, objc.Sel("rangeOfUnit:inUnit:forDate:"), smaller, larger, date)
	return NSRange(rv)
}
// Returns by reference the starting time and duration of a given calendar
// unit that contains a given date.
//
// unit: A calendar unit (see [NSCalendar.Unit] for possible values).
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// datep: Upon return, contains the starting time of the calendar unit `unit` that
// contains the date `date`
//
// tip: Upon return, contains the duration of the calendar unit `unit` that
// contains the date `date`
//
// date: A date.
//
// # Return Value
// 
// [true] if the starting time and duration of a unit could be calculated,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/range(of:start:interval:for:)
func (c NSCalendar) RangeOfUnitStartDateIntervalForDate(unit NSCalendarUnit, datep INSDate, tip unsafe.Pointer, date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("rangeOfUnit:startDate:interval:forDate:"), unit, datep, tip, date)
	return rv
}
// Returns whether a given date falls within a weekend period, and if so,
// returns by reference the start date and time interval of the weekend range.
//
// datep: Upon return, contains the starting date of the next weekend period.
//
// tip: Upon return, contains the time interval of the next weekend period.
//
// date: The date to use to perform the calculation.
//
// # Return Value
// 
// [true] if the given date falls within a weekend period, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Note that a particular calendar day may not necessarily fall entirely
// within a weekend period, as weekends can start in the middle of a day in
// some calendars and locales.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/range(ofWeekendStart:interval:containing:)
func (c NSCalendar) RangeOfWeekendStartDateIntervalContainingDate(datep INSDate, tip unsafe.Pointer, date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("rangeOfWeekendStartDate:interval:containingDate:"), datep, tip, date)
	return rv
}
// Returns the first moment of a given date as a date instance.
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// An [NSDate] instance representing the first moment date of the given date.
//
// # Discussion
// 
// For example, passing `[NSDate date]` for the `date` parameter would give
// you the start of “today.”
// 
// # Special Considerations
// 
// If there were two midnights, this method returns the first. If there was
// none, it returns the first moment that did exist.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/startOfDay(for:)
func (c NSCalendar) StartOfDayForDate(date INSDate) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("startOfDayForDate:"), date)
	return NSDateFromID(rv)
}
// Computes the dates that match (or most closely match) a given set of
// components, and calls the block once for each of them, until the
// enumeration is stopped.
//
// start: The date for which to perform the calculation.
//
// comps: The date components to match. If no components are specified, the
// enumeration will not be executed. If the `nanoseconds` component is set to
// a nonzero value, the resulting dates will have floating point `seconds`
// values that most closely match the specified `nanoseconds` value.
// Otherwise, the resulting dates will have an integer `seconds` value.
//
// opts: Options for the enumeration. For possible values, see [NSCalendar.Options].
// For usage, see below.
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// block: The block to apply to each enumerated date. The block takes three
// arguments:
// 
// date: The enumerated date. idx: Whether `date` exactly matches the
// specified date components. stop: A reference to a Boolean value. The block
// can set the value to [true] to stop further processing of the array. The
// stop argument is an out-only argument. You should only ever set this
// Boolean to [true] within the Block.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Executes a given block with dates that most closely match a given set of
// components after a given date, until the enumeration is stopped.
// 
// Strict & Non-Strict Matching
// 
// If you specify a strict matching option ([CalendarMatchStrictly]), this
// method searches as far as necessary looking for a match, up to a an
// implementation-defined limit. If an exact match is not possible, `nil` is
// passed to the `date` argument of the block, and the enumeration is stopped.
// Otherwise, this method searches as far as the next instance of the next
// highest calendar unit in the given [NSDateComponents] object.
// 
// If you do not specify a strict matching option, you must specify one of the
// following options, or else an illegal argument exception will be thrown:
// 
// [CalendarMatchPreviousTimePreservingSmallerUnits]: If specified, and there
// is no matching time before the end of the next instance of the next highest
// unit specified in the given [NSDateComponents] object, this method uses the
// existing value of the missing unit and preserves the lower units’ values.
// [CalendarMatchNextTimePreservingSmallerUnits]: If specified, and there is
// no matching time before the end of the next instance of the next highest
// unit specified in the given [NSDateComponents] object, this method uses the
// existing value of the missing unit and preserves the lower units’ values.
// [CalendarMatchNextTime]: If specified, and there is no matching time before
// the end of the next instance of the next highest unit specified in the
// given [NSDateComponents] object, this method uses the existing value of the
// missing unit and preserve the lower units’ values.
// 
// For example, if the date “February 29th” does not exist for a
// particular year, a non-strict match would return “February 28th” of
// that year instead.
// 
// [media-2852008]
// 
// As another example, if the time “2:37AM” does not exist for a
// particular day, such as when advancing by an hour at the beginning of
// Daylight Savings Time, the following is true:
// 
// - If strict matching is specified, “2:37AM” on the following the next
// day is used. - If [CalendarMatchPreviousTimePreservingSmallerUnits] is
// specified, “1:37AM” would be used instead, if that time exists. - If
// [CalendarMatchNextTimePreservingSmallerUnits] is specified, the date at the
// time “3:37AM” would be used instead, if that time exists. - If
// [CalendarMatchNextTime] is specified, the date at the time “3:00AM”
// would be used instead, if that time exists.
// 
// [media-2852011]
// 
// Matching First or Last Occurrence
// 
// If you specify a “match first” option ([CalendarMatchFirst]) and there
// are two or more matching times (that is, all components are the same)
// before the end of the next instance of the next highest unit specified in
// the given [NSDateComponents] object, this method uses the occurrence.
// 
// If you specify a “match last” option ([CalendarMatchLast]) and there
// are two or more matching times (that is, all components are the same)
// before the end of the next instance of the next highest unit specified in
// the given [NSDateComponents] object, this method uses the occurrence.
// 
// If neither “match first” or “match last” options are specified or
// both options are specified, this method behaves as if [CalendarMatchFirst]
// was specified.
// 
// There is no option to return middle occurrences of more than two
// occurrences of a matching time, if such exist.
// 
// For example, when Daylight Savings Time ends, clocks are set back by one
// hour at 2:00AM, such that times between 1:00AM and 1:59AM occur twice that
// day. The [CalendarMatchFirst] and [CalendarMatchLast] search options return
// the first and last occurrence of these times.
// 
// [media-2852012]
// 
// Forward & Backward Search
// 
// If you specify a backward search option ([CalendarSearchBackwards]), this
// method will search for previous matches before the given date. This method
// will return the same results as if the search were made in the forward
// direction from the distant past, but with the results in reverse order
// starting with the match most recent to the given date. That is, when
// searching backwards for a particular hour with no specified minute or
// second value, the resulting time is not “59” minutes and “59”
// seconds for the matching hour. When enumerating dates that repeat a time,
// such as when the clock turns back to 1:00AM from 2:00AM when Daylight
// Savings Time ends, the “first” time is determined as if the search were
// performed in the forwards direction.
// 
// For example, given a date with a time of “5:00AM” when searching for a
// minute component equal to 30, a forward search returns the time
// “5:30AM” and a backward search returns “4:30AM”.
// 
// [media-2852013]
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/enumerateDates(startingAfter:matching:options:using:)
func (c NSCalendar) EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock(start INSDate, comps INSDateComponents, opts NSCalendarOptions, block DateHandler) {
_block3, _ := NewDateBlock(block)
	objc.Send[objc.ID](c.ID, objc.Sel("enumerateDatesStartingAfterDate:matchingComponents:options:usingBlock:"), start, comps, opts, _block3)
}
// Returns the next date after a given date matching the given components.
//
// date: The date for which to perform the calculation.
//
// comps: The date components to match.
//
// options: Options for the calculation. For possible values, see [NSCalendar.Options].
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] object.
//
// # Discussion
// 
// To compute a sequence of dates, use the
// [EnumerateDatesStartingAfterDateMatchingComponentsOptionsUsingBlock] method
// instead of calling this method in a loop with the previous loop
// iteration’s result.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:options:)
func (c NSCalendar) NextDateAfterDateMatchingComponentsOptions(date INSDate, comps INSDateComponents, options NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nextDateAfterDate:matchingComponents:options:"), date, comps, options)
	return NSDateFromID(rv)
}
// Returns the next date after a given date that matches the given hour,
// minute, and second, component values.
//
// date: The date for which to perform the calculation.
//
// hourValue: The value for the hour component.
//
// minuteValue: The value for the minute component.
//
// secondValue: The value for the second component.
//
// options: Options for the calculation. For possible values, see [NSCalendar.Options].
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matchingHour:minute:second:options:)
func (c NSCalendar) NextDateAfterDateMatchingHourMinuteSecondOptions(date INSDate, hourValue int, minuteValue int, secondValue int, options NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nextDateAfterDate:matchingHour:minute:second:options:"), date, hourValue, minuteValue, secondValue, options)
	return NSDateFromID(rv)
}
// Returns the next date after a given date matching the given calendar unit
// value.
//
// date: The date for which to perform the calculation.
//
// unit: The component to use. For possible values, see [NSCalendar.Unit].
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// value: The value for the given component.
//
// options: Options for the calculation. For possible values, see [NSCalendar.Options].
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/nextDate(after:matching:value:options:)
func (c NSCalendar) NextDateAfterDateMatchingUnitValueOptions(date INSDate, unit NSCalendarUnit, value int, options NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nextDateAfterDate:matchingUnit:value:options:"), date, unit, value, options)
	return NSDateFromID(rv)
}
// Returns a date representing the absolute time calculated from given
// components.
//
// comps: The components from which to calculate the returned date.
//
// # Return Value
// 
// A new [NSDate] object representing the absolute time calculated from
// `comps`. Returns `nil` if the receiver cannot convert the components given
// in `comps` into an [NSDate] object.
//
// # Discussion
// 
// When there are insufficient components provided to completely specify an
// absolute time, a calendar uses default values of its choice. When there is
// inconsistent information, a calendar may ignore some of the components
// parameters or the method may return `nil`. Unnecessary components are
// ignored (for example, Day takes precedence over Weekday and Weekday
// ordinals).
// 
// The following example shows how to use this method to create a date object
// to represent 14:10:00 on 6 January 1965, for a given calendar
// (`gregorian`).
// 
// Note that some computations can take a relatively long time to perform.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(from:)
func (c NSCalendar) DateFromComponents(comps INSDateComponents) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateFromComponents:"), comps)
	return NSDateFromID(rv)
}
// Returns a date representing the absolute time calculated by adding given
// components to a given date.
//
// comps: The components to add to `date`.
//
// date: The date to which `comps` are added.
//
// opts: Options for the calculation. See [NSCalendar.Options] for possible values.
// 
// If you specify no options, overflow in a unit carries into the higher units
// (as in typical addition).
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] object representing the absolute time calculated by adding
// to `date` the calendrical components specified by `comps` using the options
// specified by `opts`. Returns `nil` if `date` falls outside the defined
// range of the receiver or if the computation cannot be performed.
//
// # Discussion
// 
// Some operations can be ambiguous, and the behavior of the computation is
// calendar-specific, but generally components are added in the order
// specified.
// 
// The following example shows how to add 2 months and 3 days to the current
// date and time using an existing calendar (`gregorian`):
// 
// Note that some computations can take a relatively long time.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(byAdding:to:options:)
func (c NSCalendar) DateByAddingComponentsToDateOptions(comps INSDateComponents, date INSDate, opts NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateByAddingComponents:toDate:options:"), comps, date, opts)
	return NSDateFromID(rv)
}
// Returns a date representing the absolute time calculated by adding the
// value of a given component to a given date.
//
// unit: The unit to use for the calculation. For possible values, see
// [NSCalendar.Unit].
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// value: The value for the given unit.
//
// date: The date to use to perform the calculation.
//
// options: Options for the calculation. See [NSCalendar.Options] for possible values.
// 
// If you specify a “wrap” option ([CalendarWrapComponents]), the
// specified components are incremented and wrap around to zero/one on
// overflow, but do not cause higher units to be incremented. When the wrap
// option is false, overflow in a unit carries into the higher units, as in
// typical addition.
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] object representing the absolute time calculated by adding
// to `date` the `value` of the given calendrical `unit` using the options
// specified by `options`. Returns `nil` if `date` falls outside the defined
// range of the receiver or if the computation cannot be performed.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(byAdding:value:to:options:)
func (c NSCalendar) DateByAddingUnitValueToDateOptions(unit NSCalendarUnit, value int, date INSDate, options NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateByAddingUnit:value:toDate:options:"), unit, value, date, options)
	return NSDateFromID(rv)
}
// Creates a new date calculated with the given time.
//
// h: The hour value.
//
// m: The minute value.
//
// s: The second value.
//
// date: The date to use to perform the calculation.
//
// opts: Options for the calculation. For possible values, see [NSCalendar.Options].
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] instance representing the date calculated by setting the
// given hour, minute, and second, to a given time. If no such time exists for
// the specified components, the next available date is returned, which may be
// on a different calendar day.
//
// # Discussion
// 
// You can use this method to calculate a date at a different time of a
// particular calendar day.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(bySettingHour:minute:second:of:options:)
func (c NSCalendar) DateBySettingHourMinuteSecondOfDateOptions(h int, m int, s int, date INSDate, opts NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateBySettingHour:minute:second:ofDate:options:"), h, m, s, date, opts)
	return NSDateFromID(rv)
}
// Returns a new date representing the date calculated by setting a specific
// component of a given date to a given value, while trying to keep lower
// components the same.
//
// unit: The unit to set with the given value. For possible values, see
// [NSCalendar.Unit].
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// v: The value to set for the given calendar unit.
//
// date: The date to use to perform the calculation.
//
// opts: Options for the calculation. For possible values, see [NSCalendar.Options].
// //
// [NSCalendar.Options]: https://developer.apple.com/documentation/Foundation/NSCalendar/Options
//
// # Return Value
// 
// A new [NSDate] instance representing the date calculated by setting a
// specific component of a given date to a given value. If the unit already
// has that value, this may result in a date which is the same as the given
// date. If no such time exists for the specified components, the next
// available date is returned, which may be on a different calendar day.
//
// # Discussion
// 
// Changing a component’s value often requires higher or coupled components
// to change as well. For example, setting the `weekday` to “Thursday”
// will require the `day` component to change its value, and possibly the
// `month` and `year` as well. You can use the
// [NextDateAfterDateMatchingUnitValueOptions] method to specify more precise
// behavior for determining the next or previous date for a given date
// component.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(bySettingUnit:value:of:options:)
func (c NSCalendar) DateBySettingUnitValueOfDateOptions(unit NSCalendarUnit, v int, date INSDate, opts NSCalendarOptions) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateBySettingUnit:value:ofDate:options:"), unit, v, date, opts)
	return NSDateFromID(rv)
}
// Returns a date created with the given components.
//
// eraValue: The value to set for the era.
//
// yearValue: The value to set for the year.
//
// monthValue: The value to set for the month.
//
// dayValue: The value to set for the day.
//
// hourValue: The value to set for the hour.
//
// minuteValue: The value to set for the minute.
//
// secondValue: The value to set for the second.
//
// nanosecondValue: The value to set for the nanosecond.
//
// # Return Value
// 
// A new [NSDate] instance created with the given components, or `nil` if the
// components do not correspond to a valid date.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(era:year:month:day:hour:minute:second:nanosecond:)
func (c NSCalendar) DateWithEraYearMonthDayHourMinuteSecondNanosecond(eraValue int, yearValue int, monthValue int, dayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateWithEra:year:month:day:hour:minute:second:nanosecond:"), eraValue, yearValue, monthValue, dayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return NSDateFromID(rv)
}
// Returns a new date created with the given components base on a week-of-year
// value.
//
// eraValue: The value for the era component.
//
// yearValue: The value for the year component.
//
// weekValue: The value for the week-of-year component.
//
// weekdayValue: The value to use as the weekday.
//
// hourValue: The value for the hour component.
//
// minuteValue: The value for the minute component.
//
// secondValue: The value for the second component.
//
// nanosecondValue: The value for the nanosecond component.
//
// # Return Value
// 
// A new [NSDate] instance created with the given components based on a
// week-of-year calculation, or `nil` if the components do not correspond to a
// valid date.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/date(era:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:)
func (c NSCalendar) DateWithEraYearForWeekOfYearWeekOfYearWeekdayHourMinuteSecondNanosecond(eraValue int, yearValue int, weekValue int, weekdayValue int, hourValue int, minuteValue int, secondValue int, nanosecondValue int) INSDate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dateWithEra:yearForWeekOfYear:weekOfYear:weekday:hour:minute:second:nanosecond:"), eraValue, yearValue, weekValue, weekdayValue, hourValue, minuteValue, secondValue, nanosecondValue)
	return NSDateFromID(rv)
}
// Returns by reference the starting date and time interval range of the next
// weekend period after a given date.
//
// datep: Upon return, contains the starting date of the next weekend period.
//
// tip: Upon return, contains the time interval of the next weekend period.
//
// options: Options for the calculation. If you specify a backward search option
// ([CalendarSearchBackwards]), the starting date and time interval range of
// the preceding weekend period will be returned by reference instead.
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// [false] if the calendar and locale do not have the concept of a weekend,
// otherwise [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Note that a particular calendar day may not necessarily fall entirely
// within a weekend period, as weekends can start in the middle of a day in
// some calendars and locales.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/nextWeekendStart(_:interval:options:after:)
func (c NSCalendar) NextWeekendStartDateIntervalOptionsAfterDate(datep INSDate, tip unsafe.Pointer, options NSCalendarOptions, date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("nextWeekendStartDate:interval:options:afterDate:"), datep, tip, options, date)
	return rv
}
// Indicates the ordering of two given dates based on their components down to
// a given unit granularity.
//
// date1: The first date to compare.
//
// date2: The second date to compare.
//
// unit: The smallest unit that must, along with all larger units, be equal for the
// given dates to be considered the same. For possible values, see
// [NSCalendar.Unit].
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// # Return Value
// 
// [NSOrderedSame] if the dates are the same down to the given granularity,
// otherwise [NSOrderedAscending] or [NSOrderedDescending].
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/compare(_:to:toUnitGranularity:)
func (c NSCalendar) CompareDateToDateToUnitGranularity(date1 INSDate, date2 INSDate, unit NSCalendarUnit) ComparisonResult {
	rv := objc.Send[ComparisonResult](c.ID, objc.Sel("compareDate:toDate:toUnitGranularity:"), date1, date2, unit)
	return ComparisonResult(rv)
}
// Indicates whether two dates are equal to a given unit of granularity.
//
// date1: The first date to compare.
//
// date2: The second date to compare.
//
// unit: The smallest unit that must, along with all larger units, be equal in the
// given dates. For possible values, see [NSCalendar.Unit].
// //
// [NSCalendar.Unit]: https://developer.apple.com/documentation/Foundation/NSCalendar/Unit
//
// # Return Value
// 
// [true] if both dates have equal date component for all units greater than
// or equal to the given unit, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/isDate(_:equalTo:toUnitGranularity:)
func (c NSCalendar) IsDateEqualToDateToUnitGranularity(date1 INSDate, date2 INSDate, unit NSCalendarUnit) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDate:equalToDate:toUnitGranularity:"), date1, date2, unit)
	return rv
}
// Indicates whether two dates are in the same day.
//
// date1: The first date to compare.
//
// date2: The second date to compare.
//
// # Return Value
// 
// [true] if both dates are within the same day, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/isDate(_:inSameDayAs:)
func (c NSCalendar) IsDateInSameDayAsDate(date1 INSDate, date2 INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDate:inSameDayAsDate:"), date1, date2)
	return rv
}
// Indicates whether the given date is in “today.”
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// [true] if the given date is in “today,” otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInToday(_:)
func (c NSCalendar) IsDateInToday(date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDateInToday:"), date)
	return rv
}
// Indicates whether the given date is in “tomorrow.”
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// [true] if the given date is in “tomorrow,” otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInTomorrow(_:)
func (c NSCalendar) IsDateInTomorrow(date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDateInTomorrow:"), date)
	return rv
}
// Indicates whether a given date falls within a weekend period, as defined by
// the calendar and the calendar’s locale.
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// [true] if the given date is within a weekend period, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the date does fall within a weekend, you can use the
// [RangeOfWeekendStartDateIntervalContainingDate] method to determine the
// start date of that weekend period. Otherwise, you can use the
// [NextWeekendStartDateIntervalOptionsAfterDate] method to determine the
// start date of the next or previous weekend.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInWeekend(_:)
func (c NSCalendar) IsDateInWeekend(date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDateInWeekend:"), date)
	return rv
}
// Indicates whether the given date is in “yesterday.”
//
// date: The date for which to perform the calculation.
//
// # Return Value
// 
// [true] if the given date is in “yesterday,” otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/isDateInYesterday(_:)
func (c NSCalendar) IsDateInYesterday(date INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDateInYesterday:"), date)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (c NSCalendar) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (c NSCalendar) InitWithCoder(coder INSCoder) NSCalendar {
	rv := objc.Send[NSCalendar](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// An identifier for the calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/calendarIdentifier
func (c NSCalendar) CalendarIdentifier() NSCalendarIdentifier {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("calendarIdentifier"))
	return NSCalendarIdentifier(NSStringFromID(rv).String())
}
// The index of the first weekday of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/firstWeekday
func (c NSCalendar) FirstWeekday() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("firstWeekday"))
	return rv
}
func (c NSCalendar) SetFirstWeekday(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setFirstWeekday:"), value)
}
// The locale of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/locale
func (c NSCalendar) Locale() INSLocale {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (c NSCalendar) SetLocale(value INSLocale) {
	objc.Send[struct{}](c.ID, objc.Sel("setLocale:"), value)
}
// The time zone for the calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/timeZone
func (c NSCalendar) TimeZone() INSTimeZone {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("timeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
func (c NSCalendar) SetTimeZone(value INSTimeZone) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeZone:"), value)
}
// The minimum number of days in the first week of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/minimumDaysInFirstWeek
func (c NSCalendar) MinimumDaysInFirstWeek() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("minimumDaysInFirstWeek"))
	return rv
}
func (c NSCalendar) SetMinimumDaysInFirstWeek(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumDaysInFirstWeek:"), value)
}
// The symbol used to represent “AM” for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/amSymbol
func (c NSCalendar) AMSymbol() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AMSymbol"))
	return NSStringFromID(rv).String()
}
// The symbol used to represent “PM” for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/pmSymbol
func (c NSCalendar) PMSymbol() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("PMSymbol"))
	return NSStringFromID(rv).String()
}
// A list of weekdays in this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/weekdaySymbols
func (c NSCalendar) WeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("weekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of shorter-named weekdays in this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/shortWeekdaySymbols
func (c NSCalendar) ShortWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shortWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of very-shortly-named weekdays in this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortWeekdaySymbols
func (c NSCalendar) VeryShortWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("veryShortWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of standalone weekday symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneWeekdaySymbols
func (c NSCalendar) StandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("standaloneWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of short standalone weekday symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneWeekdaySymbols
func (c NSCalendar) ShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shortStandaloneWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of very short standalone weekday symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortStandaloneWeekdaySymbols
func (c NSCalendar) VeryShortStandaloneWeekdaySymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("veryShortStandaloneWeekdaySymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of month symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/monthSymbols
func (c NSCalendar) MonthSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("monthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of short month symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/shortMonthSymbols
func (c NSCalendar) ShortMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shortMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of very short month symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortMonthSymbols
func (c NSCalendar) VeryShortMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("veryShortMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of standalone month symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneMonthSymbols
func (c NSCalendar) StandaloneMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("standaloneMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of short standalone month symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneMonthSymbols
func (c NSCalendar) ShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shortStandaloneMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of very short month symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/veryShortStandaloneMonthSymbols
func (c NSCalendar) VeryShortStandaloneMonthSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("veryShortStandaloneMonthSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of quarter symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/quarterSymbols
func (c NSCalendar) QuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("quarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of short quarter symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/shortQuarterSymbols
func (c NSCalendar) ShortQuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shortQuarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of standalone quarter symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/standaloneQuarterSymbols
func (c NSCalendar) StandaloneQuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("standaloneQuarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of short standalone quarter symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/shortStandaloneQuarterSymbols
func (c NSCalendar) ShortStandaloneQuarterSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("shortStandaloneQuarterSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of era symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/eraSymbols
func (c NSCalendar) EraSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("eraSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A list of long era symbols for this calendar.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/longEraSymbols
func (c NSCalendar) LongEraSymbols() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("longEraSymbols"))
	return objc.ConvertSliceToStrings(rv)
}
// A notification that is posted whenever the calendar day of the system
// changes, as determined by the system calendar, locale, and time zone.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nscalendardaychanged
func (c NSCalendar) NSCalendarDayChanged() NSNotificationName {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("NSCalendarDayChangedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// The user’s current calendar.
//
// # Return Value
// 
// The logical calendar for the current user.
// 
// # Discussion
// 
// The returned calendar is formed from the settings for the current user’s
// chosen system locale overlaid with any custom settings the user has
// specified in System Preferences. Settings you get from this calendar do not
// change as System Preferences are changed, so that your operations are
// consistent (contrast with [AutoupdatingCurrentCalendar]).
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/current
func (_NSCalendarClass NSCalendarClass) CurrentCalendar() NSCalendar {
	rv := objc.Send[objc.ID](objc.ID(_NSCalendarClass.class), objc.Sel("currentCalendar"))
	return NSCalendarFromID(objc.ID(rv))
}
// A calendar that tracks changes to user’s preferred calendar.
//
// # Return Value
// 
// The current logical calendar for the current user.
// 
// # Discussion
// 
// Settings you get from this calendar do change as the user’s settings
// change (contrast with [CurrentCalendar]).
// 
// Note that if you cache values based on the calendar or related information
// those caches will of course not be automatically updated by the updating of
// the calendar object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCalendar/autoupdatingCurrent
func (_NSCalendarClass NSCalendarClass) AutoupdatingCurrentCalendar() NSCalendar {
	rv := objc.Send[objc.ID](objc.ID(_NSCalendarClass.class), objc.Sel("autoupdatingCurrentCalendar"))
	return NSCalendarFromID(objc.ID(rv))
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

