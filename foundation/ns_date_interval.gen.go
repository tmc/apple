// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDateInterval] class.
var (
	_NSDateIntervalClass     NSDateIntervalClass
	_NSDateIntervalClassOnce sync.Once
)

func getNSDateIntervalClass() NSDateIntervalClass {
	_NSDateIntervalClassOnce.Do(func() {
		_NSDateIntervalClass = NSDateIntervalClass{class: objc.GetClass("NSDateInterval")}
	})
	return _NSDateIntervalClass
}

// GetNSDateIntervalClass returns the class object for NSDateInterval.
func GetNSDateIntervalClass() NSDateIntervalClass {
	return getNSDateIntervalClass()
}

type NSDateIntervalClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDateIntervalClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDateIntervalClass) Alloc() NSDateInterval {
	rv := objc.Send[NSDateInterval](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object representing the span of time between a specific start date and
// end date.
//
// # Overview
// 
// In Swift, this object bridges to [DateInterval]; use [NSDateInterval] when
// you need reference semantics or other Foundation-specific behavior.
// 
// An [NSDateInterval] object represents a closed interval between two dates.
// The [NSDateInterval] class provides a programmatic interface for
// calculating the duration of a time interval and determining whether a date
// falls within it, as well as comparing date intervals and checking to see
// whether they intersect.
// 
// An [NSDateInterval] object consists of a [NSDateInterval.StartDate] and an [NSDateInterval.EndDate]. The
// [NSDateInterval.StartDate] and [NSDateInterval.EndDate] of a date interval can be equal, in which case
// its [NSDateInterval.Duration] is `0`. However, [NSDateInterval.EndDate] cannot occur earlier than
// [NSDateInterval.StartDate].
// 
// You can use the [NSDateIntervalFormatter] class to create string
// representations of [NSDateInterval] objects that are suitable for display
// in the current locale.
//
// [DateInterval]: https://developer.apple.com/documentation/Foundation/DateInterval
//
// # Creating Date Intervals
//
//   - [NSDateInterval.InitWithStartDateDuration]: Initializes a date interval with a given start date and duration.
//   - [NSDateInterval.InitWithStartDateEndDate]: Initializes a date interval from a given start date and end date.
//
// # Accessing Start Date, End Date, and Duration
//
//   - [NSDateInterval.StartDate]: The start date of the date interval.
//   - [NSDateInterval.EndDate]: The end date of the date interval.
//   - [NSDateInterval.Duration]: The duration of the date interval.
//
// # Comparing Date Intervals
//
//   - [NSDateInterval.Compare]: Compares the receiver with the specified date interval.
//   - [NSDateInterval.IsEqualToDateInterval]: Indicates whether the receiver is equal to the specified date interval.
//
// # Determining Intersections
//
//   - [NSDateInterval.IntersectsDateInterval]: Indicates whether the receiver intersects with the specified date interval.
//   - [NSDateInterval.IntersectionWithDateInterval]: Returns the intersection between the receiver and the specified date interval.
//
// # Determining Whether a Date Occurs Within a Date Interval
//
//   - [NSDateInterval.ContainsDate]: Indicates whether the receiver contains the specified date.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval
type NSDateInterval struct {
	objectivec.Object
}

// NSDateIntervalFromID constructs a [NSDateInterval] from an objc.ID.
//
// An object representing the span of time between a specific start date and
// end date.
func NSDateIntervalFromID(id objc.ID) NSDateInterval {
	return NSDateInterval{objectivec.Object{ID: id}}
}
// NOTE: NSDateInterval adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDateInterval] class.
//
// # Creating Date Intervals
//
//   - [INSDateInterval.InitWithStartDateDuration]: Initializes a date interval with a given start date and duration.
//   - [INSDateInterval.InitWithStartDateEndDate]: Initializes a date interval from a given start date and end date.
//
// # Accessing Start Date, End Date, and Duration
//
//   - [INSDateInterval.StartDate]: The start date of the date interval.
//   - [INSDateInterval.EndDate]: The end date of the date interval.
//   - [INSDateInterval.Duration]: The duration of the date interval.
//
// # Comparing Date Intervals
//
//   - [INSDateInterval.Compare]: Compares the receiver with the specified date interval.
//   - [INSDateInterval.IsEqualToDateInterval]: Indicates whether the receiver is equal to the specified date interval.
//
// # Determining Intersections
//
//   - [INSDateInterval.IntersectsDateInterval]: Indicates whether the receiver intersects with the specified date interval.
//   - [INSDateInterval.IntersectionWithDateInterval]: Returns the intersection between the receiver and the specified date interval.
//
// # Determining Whether a Date Occurs Within a Date Interval
//
//   - [INSDateInterval.ContainsDate]: Indicates whether the receiver contains the specified date.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval
type INSDateInterval interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating Date Intervals

	// Initializes a date interval with a given start date and duration.
	InitWithStartDateDuration(startDate INSDate, duration float64) NSDateInterval
	// Initializes a date interval from a given start date and end date.
	InitWithStartDateEndDate(startDate INSDate, endDate INSDate) NSDateInterval

	// Topic: Accessing Start Date, End Date, and Duration

	// The start date of the date interval.
	StartDate() INSDate
	// The end date of the date interval.
	EndDate() INSDate
	// The duration of the date interval.
	Duration() float64

	// Topic: Comparing Date Intervals

	// Compares the receiver with the specified date interval.
	Compare(dateInterval INSDateInterval) ComparisonResult
	// Indicates whether the receiver is equal to the specified date interval.
	IsEqualToDateInterval(dateInterval INSDateInterval) bool

	// Topic: Determining Intersections

	// Indicates whether the receiver intersects with the specified date interval.
	IntersectsDateInterval(dateInterval INSDateInterval) bool
	// Returns the intersection between the receiver and the specified date interval.
	IntersectionWithDateInterval(dateInterval INSDateInterval) INSDateInterval

	// Topic: Determining Whether a Date Occurs Within a Date Interval

	// Indicates whether the receiver contains the specified date.
	ContainsDate(date INSDate) bool
}

// Init initializes the instance.
func (d NSDateInterval) Init() NSDateInterval {
	rv := objc.Send[NSDateInterval](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDateInterval) Autorelease() NSDateInterval {
	rv := objc.Send[NSDateInterval](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDateInterval creates a new NSDateInterval instance.
func NewNSDateInterval() NSDateInterval {
	class := getNSDateIntervalClass()
	rv := objc.Send[NSDateInterval](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a date interval initialized from data in the given unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(coder:)
func NewDateIntervalWithCoder(coder INSCoder) NSDateInterval {
	instance := getNSDateIntervalClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDateIntervalFromID(rv)
}

// Initializes a date interval with a given start date and duration.
//
// startDate: The start date of the date interval.
//
// duration: The duration from the start date for the date interval.
//
// # Discussion
// 
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:duration:)
func NewDateIntervalWithStartDateDuration(startDate INSDate, duration float64) NSDateInterval {
	instance := getNSDateIntervalClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStartDate:duration:"), startDate, duration)
	return NSDateIntervalFromID(rv)
}

// Initializes a date interval from a given start date and end date.
//
// startDate: The start date of the date interval.
//
// endDate: The end date of the date interval.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:end:)
func NewDateIntervalWithStartDateEndDate(startDate INSDate, endDate INSDate) NSDateInterval {
	instance := getNSDateIntervalClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStartDate:endDate:"), startDate, endDate)
	return NSDateIntervalFromID(rv)
}

// Initializes a date interval with a given start date and duration.
//
// startDate: The start date of the date interval.
//
// duration: The duration from the start date for the date interval.
//
// # Discussion
// 
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:duration:)
func (d NSDateInterval) InitWithStartDateDuration(startDate INSDate, duration float64) NSDateInterval {
	rv := objc.Send[NSDateInterval](d.ID, objc.Sel("initWithStartDate:duration:"), startDate, duration)
	return rv
}
// Initializes a date interval from a given start date and end date.
//
// startDate: The start date of the date interval.
//
// endDate: The end date of the date interval.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(start:end:)
func (d NSDateInterval) InitWithStartDateEndDate(startDate INSDate, endDate INSDate) NSDateInterval {
	rv := objc.Send[NSDateInterval](d.ID, objc.Sel("initWithStartDate:endDate:"), startDate, endDate)
	return rv
}
// Returns a date interval initialized from data in the given unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/init(coder:)
func (d NSDateInterval) InitWithCoder(coder INSCoder) NSDateInterval {
	rv := objc.Send[NSDateInterval](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Compares the receiver with the specified date interval.
//
// dateInterval: The date interval with which to compare the receiver.
//
// # Return Value
// 
// Returns an [ComparisonResult] value that indicates the temporal ordering of
// the receiver and a given date interval:
// 
// - [OrderedAscending] if the receiver’s [StartDate] occurs earlier than
// that of `dateInterval`, or both [StartDate] values are equal and the
// [Duration] of the receiver is less than that of `dateInterval`. -
// [OrderedDescending] if the receiver’s [StartDate] occurs later than that
// of `dateInterval`, or both [StartDate] values are equal and the [Duration]
// of the receiver is greater than that of `dateInterval`. - [OrderedSame] if
// the receiver’s [StartDate] and [Duration] values are equal to those of
// `dateInterval`.
//
// [ComparisonResult]: https://developer.apple.com/documentation/Foundation/ComparisonResult
//
// # Discussion
// 
// The following figure illustrates four [NSDateInterval] objects plotted on
// an arbitrary time axis. Each date interval spans its [Duration] from left
// to right, from its [StartDate] to its [EndDate].
// 
// [media-2556955]
// 
// The result of comparing the date interval labeled with the date interval
// labeled is [OrderedAscending], because has a [StartDate] that occurs
// earlier than that of .
// 
// The result of comparing the date interval labeled with the date interval
// labeled is [OrderedDescending], because because and have the same
// [StartDate], and has a [Duration] greater than that of .
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/compare(_:)
func (d NSDateInterval) Compare(dateInterval INSDateInterval) ComparisonResult {
	rv := objc.Send[ComparisonResult](d.ID, objc.Sel("compare:"), dateInterval)
	return ComparisonResult(rv)
}
// Indicates whether the receiver is equal to the specified date interval.
//
// dateInterval: The date interval with which to check the receiver for equality.
//
// # Return Value
// 
// [true] if the [StartDate] and [Duration] of `dateInterval` and the receiver
// are equal. Otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/isEqual(to:)
func (d NSDateInterval) IsEqualToDateInterval(dateInterval INSDateInterval) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEqualToDateInterval:"), dateInterval)
	return rv
}
// Indicates whether the receiver intersects with the specified date interval.
//
// dateInterval: The date interval with which to check the receiver for intersection.
//
// # Discussion
// 
// See [IntersectionWithDateInterval] for more information about determining
// whether two date intervals intersect.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/intersects(_:)
func (d NSDateInterval) IntersectsDateInterval(dateInterval INSDateInterval) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("intersectsDateInterval:"), dateInterval)
	return rv
}
// Returns the intersection between the receiver and the specified date
// interval.
//
// dateInterval: The date interval with which to calculate the intersection of the receiver.
//
// # Return Value
// 
// A date interval for the intersection of the receiver and `dateInterval`, or
// `nil` if no intersection occurs.
//
// # Discussion
// 
// Calculating the intersection of date intervals is a commutative and
// associative operation. The intersection of a date interval with itself is
// equal to itself.
// 
// The following figure illustrates five [NSDateInterval] objects plotted on
// an arbitrary time axis. Each date interval spans its [Duration] from left
// to right, from its [StartDate] to its [EndDate].
// 
// [media-2556958]
// 
// The date intervals labeled and do not intersect, because the [StartDate] of
// occurs later than the [EndDate] of .
// 
// The date intervals labeled and do intersect. The date interval labeled
// represents the result of calculating the intersection between and .
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/intersection(with:)
func (d NSDateInterval) IntersectionWithDateInterval(dateInterval INSDateInterval) INSDateInterval {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("intersectionWithDateInterval:"), dateInterval)
	return NSDateIntervalFromID(rv)
}
// Indicates whether the receiver contains the specified date.
//
// date: The date for which to test membership of the date interval.
//
// # Return Value
// 
// [true] if the receiver contains `date`. Otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/contains(_:)
func (d NSDateInterval) ContainsDate(date INSDate) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("containsDate:"), date)
	return rv
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (d NSDateInterval) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The start date of the date interval.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/startDate
func (d NSDateInterval) StartDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("startDate"))
	return NSDateFromID(objc.ID(rv))
}
// The end date of the date interval.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/endDate
func (d NSDateInterval) EndDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("endDate"))
	return NSDateFromID(objc.ID(rv))
}
// The duration of the date interval.
//
// See: https://developer.apple.com/documentation/Foundation/NSDateInterval/duration
func (d NSDateInterval) Duration() float64 {
	rv := objc.Send[NSTimeInterval](d.ID, objc.Sel("duration"))
	return float64(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

