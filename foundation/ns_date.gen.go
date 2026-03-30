// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDate] class.
var (
	_NSDateClass     NSDateClass
	_NSDateClassOnce sync.Once
)

func getNSDateClass() NSDateClass {
	_NSDateClassOnce.Do(func() {
		_NSDateClass = NSDateClass{class: objc.GetClass("NSDate")}
	})
	return _NSDateClass
}

// GetNSDateClass returns the class object for NSDate.
func GetNSDateClass() NSDateClass {
	return getNSDateClass()
}

type NSDateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDateClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDateClass) Alloc() NSDate {
	rv := objc.Send[NSDate](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a specific point in time, independent of any calendar
// or time zone.
//
// # Overview
//
// In Swift, use this type when you need reference semantics or other
// Foundation-specific behavior.
//
// [NSDate] objects encapsulate a single point in time, independent of any
// particular calendrical system or time zone. Date objects are immutable,
// representing an invariant time interval relative to an absolute reference
// date (00:00:00 UTC on 1 January 2001).
//
// The [NSDate] class provides methods for comparing dates, calculating the
// time interval between two dates, and creating a new date from a time
// interval relative to another date. [NSDate] objects can be used in
// conjunction with [NSDateFormatter] objects to create localized
// representations of dates and times, as well as with [NSCalendar] objects to
// perform calendar arithmetic.
//
// [NSDate] is with its Core Foundation counterpart, [CFDate]. See [Toll-Free
// Bridging] for more information on toll-free bridging.
//
// # Subclassing Notes
//
// You might subclass [NSDate] in order to make it easier to work with a
// particular calendrical system, or to work with date and time values with a
// finer temporal granularity.
//
// # Methods to Override and Other Requirements
//
// If you want to subclass [NSDate] to obtain behavior different than that
// provided by the private or public subclasses, you must:
//
// - Declare a suitable instance variable to hold the date and time value
// (relative to an absolute reference date) - Override the
// [NSDate.TimeIntervalSinceReferenceDate] instance method to provide the correct
// date and time value based on your instance variable - Override
// [NSDate.InitWithTimeIntervalSinceReferenceDate], one of the designated initializer
// methods - If creating a subclass that represents a calendrical system,
// define methods that partition past and future periods into the units of
// this calendar - Implement the methods required by the [NSCopying] and
// [NSCoding] protocols, because [NSDate] adopts these protocols
//
// # Special Considerations
//
// Your subclass may use a different reference date than the absolute
// reference date used by [NSDate] (00:00:00 UTC on 1 January 2001). If it
// does, it must still use the absolute reference date in its implementations
// of the methods [NSDate.TimeIntervalSinceReferenceDate] and
// [NSDate.InitWithTimeIntervalSinceReferenceDate]. That is, the reference date
// referred to in the titles of these methods is the absolute reference date.
// If you do not use the absolute reference date in these methods, comparisons
// between [NSDate] objects of your subclass and [NSDate] objects of a private
// subclass will not work.
//
// # Initializing a Date
//
//   - [NSDate.InitWithTimeIntervalSinceNow]: Returns a date object initialized relative to the current date and time by a given number of seconds.
//   - [NSDate.InitWithTimeIntervalSinceReferenceDate]: Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds.
//   - [NSDate.InitWithTimeIntervalSince1970]: Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds.
//
// # Comparing Dates
//
//   - [NSDate.IsEqualToDate]: Returns a Boolean value that indicates whether a given object is a date that is exactly equal the receiver.
//   - [NSDate.EarlierDate]: Returns the earlier of the receiver and another given date.
//   - [NSDate.LaterDate]: Returns the later of the receiver and another given date.
//   - [NSDate.Compare]: Indicates the temporal ordering of the receiver and another given date.
//
// # Getting Time Intervals
//
//   - [NSDate.TimeIntervalSinceDate]: Returns the interval between the receiver and another given date.
//   - [NSDate.TimeIntervalSinceNow]: The interval between the date object and the current date and time.
//   - [NSDate.TimeIntervalSinceReferenceDate]: The interval between the date object and 00:00:00 UTC on 1 January 2001.
//   - [NSDate.TimeIntervalSince1970]: The interval between the date object and 00:00:00 UTC on 1 January 1970.
//   - [NSDate.NSTimeIntervalSince1970]: The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
//   - [NSDate.SetNSTimeIntervalSince1970]
//
// # Adding Time Intervals
//
//   - [NSDate.DateByAddingTimeInterval]: Returns a new date object that is set to a given number of seconds relative to the receiver.
//
// # Describing Dates
//
//   - [NSDate.Description]: A string representation of the date object.
//   - [NSDate.DescriptionWithLocale]: Returns a string representation of the date using the given locale.
//   - [NSDate.CustomPlaygroundQuickLook]: A custom playground Quick Look for this object.
//   - [NSDate.SetCustomPlaygroundQuickLook]
//
// # Recognizing Notifications
//
//   - [NSDate.NSSystemClockDidChange]: A notification posted whenever the system clock is changed.
//
// # Initializers
//
//   - [NSDate.InitWithTimeIntervalSinceDate]
//
// See: https://developer.apple.com/documentation/Foundation/NSDate
//
// [CFDate]: https://developer.apple.com/documentation/CoreFoundation/CFDate
// [NSCoding]: https://developer.apple.com/library/archive/documentation/Swift/Conceptual/BuildingCocoaApps/WritingSwiftClassesWithObjective-CBehavior.html#//apple_ref/doc/uid/TP40014216-CH5-ID152
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
type NSDate struct {
	objectivec.Object
}

// NSDateFromID constructs a [NSDate] from an objc.ID.
//
// A representation of a specific point in time, independent of any calendar
// or time zone.
func NSDateFromID(id objc.ID) NSDate {
	return NSDate{objectivec.Object{ID: id}}
}

// NOTE: NSDate adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDate] class.
//
// # Initializing a Date
//
//   - [INSDate.InitWithTimeIntervalSinceNow]: Returns a date object initialized relative to the current date and time by a given number of seconds.
//   - [INSDate.InitWithTimeIntervalSinceReferenceDate]: Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds.
//   - [INSDate.InitWithTimeIntervalSince1970]: Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds.
//
// # Comparing Dates
//
//   - [INSDate.IsEqualToDate]: Returns a Boolean value that indicates whether a given object is a date that is exactly equal the receiver.
//   - [INSDate.EarlierDate]: Returns the earlier of the receiver and another given date.
//   - [INSDate.LaterDate]: Returns the later of the receiver and another given date.
//   - [INSDate.Compare]: Indicates the temporal ordering of the receiver and another given date.
//
// # Getting Time Intervals
//
//   - [INSDate.TimeIntervalSinceDate]: Returns the interval between the receiver and another given date.
//   - [INSDate.TimeIntervalSinceNow]: The interval between the date object and the current date and time.
//   - [INSDate.TimeIntervalSinceReferenceDate]: The interval between the date object and 00:00:00 UTC on 1 January 2001.
//   - [INSDate.TimeIntervalSince1970]: The interval between the date object and 00:00:00 UTC on 1 January 1970.
//   - [INSDate.NSTimeIntervalSince1970]: The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
//   - [INSDate.SetNSTimeIntervalSince1970]
//
// # Adding Time Intervals
//
//   - [INSDate.DateByAddingTimeInterval]: Returns a new date object that is set to a given number of seconds relative to the receiver.
//
// # Describing Dates
//
//   - [INSDate.Description]: A string representation of the date object.
//   - [INSDate.DescriptionWithLocale]: Returns a string representation of the date using the given locale.
//   - [INSDate.CustomPlaygroundQuickLook]: A custom playground Quick Look for this object.
//   - [INSDate.SetCustomPlaygroundQuickLook]
//
// # Recognizing Notifications
//
//   - [INSDate.NSSystemClockDidChange]: A notification posted whenever the system clock is changed.
//
// # Initializers
//
//   - [INSDate.InitWithTimeIntervalSinceDate]
//
// See: https://developer.apple.com/documentation/Foundation/NSDate
type INSDate interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Initializing a Date

	// Returns a date object initialized relative to the current date and time by a given number of seconds.
	InitWithTimeIntervalSinceNow(secs float64) NSDate
	// Returns a date object initialized relative to 00:00:00 UTC on 1 January 2001 by a given number of seconds.
	InitWithTimeIntervalSinceReferenceDate(ti float64) NSDate
	// Returns a date object initialized relative to 00:00:00 UTC on 1 January 1970 by a given number of seconds.
	InitWithTimeIntervalSince1970(secs float64) NSDate

	// Topic: Comparing Dates

	// Returns a Boolean value that indicates whether a given object is a date that is exactly equal the receiver.
	IsEqualToDate(otherDate INSDate) bool
	// Returns the earlier of the receiver and another given date.
	EarlierDate(anotherDate INSDate) INSDate
	// Returns the later of the receiver and another given date.
	LaterDate(anotherDate INSDate) INSDate
	// Indicates the temporal ordering of the receiver and another given date.
	Compare(other INSDate) ComparisonResult

	// Topic: Getting Time Intervals

	// Returns the interval between the receiver and another given date.
	TimeIntervalSinceDate(anotherDate INSDate) float64
	// The interval between the date object and the current date and time.
	TimeIntervalSinceNow() float64
	// The interval between the date object and 00:00:00 UTC on 1 January 2001.
	TimeIntervalSinceReferenceDate() float64
	// The interval between the date object and 00:00:00 UTC on 1 January 1970.
	TimeIntervalSince1970() float64
	// The number of seconds from 1 January 1970 to the reference date, 1 January 2001.
	NSTimeIntervalSince1970() float64
	SetNSTimeIntervalSince1970(value float64)

	// Topic: Adding Time Intervals

	// Returns a new date object that is set to a given number of seconds relative to the receiver.
	DateByAddingTimeInterval(ti float64) INSDate

	// Topic: Describing Dates

	// A string representation of the date object.
	Description() string
	// Returns a string representation of the date using the given locale.
	DescriptionWithLocale(locale objectivec.IObject) string
	// A custom playground Quick Look for this object.
	CustomPlaygroundQuickLook() objectivec.IObject
	SetCustomPlaygroundQuickLook(value objectivec.IObject)

	// Topic: Recognizing Notifications

	// A notification posted whenever the system clock is changed.
	NSSystemClockDidChange() NSNotificationName

	// Topic: Initializers

	InitWithTimeIntervalSinceDate(secsToBeAdded float64, date INSDate) NSDate
}

// Init initializes the instance.
func (d NSDate) Init() NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDate) Autorelease() NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDate creates a new NSDate instance.
func NewNSDate() NSDate {
	class := getNSDateClass()
	rv := objc.Send[NSDate](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a date object initialized from data in the given unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(coder:)
func NewDateWithCoder(coder INSCoder) NSDate {
	instance := getNSDateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDateFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSDate/init(SRAbsoluteTime:)-886t8
func NewDateWithSRAbsoluteTime(time corefoundation.CFTimeInterval) NSDate {
	instance := getNSDateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSRAbsoluteTime:"), time)
	return NSDateFromID(rv)
}

// Returns a date object initialized relative to 00:00:00 UTC on 1 January
// 1970 by a given number of seconds.
//
// secs: The number of seconds from the reference date (00:00:00 UTC on 1 January
// 1970) for the new date. Use a negative argument to specify a date and time
// before the reference date.
//
// # Return Value
//
// An [NSDate] object set to `seconds` seconds from the reference date.
//
// # Discussion
//
// This method is useful for creating [NSDate] objects from time_t values
// returned by BSD system functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSince1970:)
func NewDateWithTimeIntervalSince1970(secs float64) NSDate {
	instance := getNSDateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeIntervalSince1970:"), secs)
	return NSDateFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeInterval:since:)
func NewDateWithTimeIntervalSinceDate(secsToBeAdded float64, date INSDate) NSDate {
	instance := getNSDateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	return NSDateFromID(rv)
}

// Returns a date object initialized relative to the current date and time by
// a given number of seconds.
//
// secs: The number of seconds from relative to the current date and time to which
// the receiver should be initialized. A negative value means the returned
// object will represent a date in the past.
//
// # Return Value
//
// An [NSDate] object initialized relative to the current date and time by
// `secs` seconds.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSinceNow:)
func NewDateWithTimeIntervalSinceNow(secs float64) NSDate {
	instance := getNSDateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeIntervalSinceNow:"), secs)
	return NSDateFromID(rv)
}

// Returns a date object initialized relative to 00:00:00 UTC on 1 January
// 2001 by a given number of seconds.
//
// ti: The number of seconds to add to the reference date (00:00:00 UTC on 1
// January 2001). A negative value means the receiver will be earlier than the
// reference date.
//
// # Return Value
//
// An [NSDate] object initialized relative to the absolute reference date by
// `ti` seconds.
//
// # Discussion
//
// This method is a designated initializer for the [NSDate] class and is
// declared primarily for the use of subclasses of [NSDate]. When you subclass
// [NSDate] to create a concrete date class, you must override this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSinceReferenceDate:)
func NewDateWithTimeIntervalSinceReferenceDate(ti float64) NSDate {
	instance := getNSDateClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeIntervalSinceReferenceDate:"), ti)
	return NSDateFromID(rv)
}

// Returns a date object initialized relative to the current date and time by
// a given number of seconds.
//
// secs: The number of seconds from relative to the current date and time to which
// the receiver should be initialized. A negative value means the returned
// object will represent a date in the past.
//
// # Return Value
//
// An [NSDate] object initialized relative to the current date and time by
// `secs` seconds.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSinceNow:)
func (d NSDate) InitWithTimeIntervalSinceNow(secs float64) NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("initWithTimeIntervalSinceNow:"), secs)
	return rv
}

// Returns a date object initialized relative to 00:00:00 UTC on 1 January
// 2001 by a given number of seconds.
//
// ti: The number of seconds to add to the reference date (00:00:00 UTC on 1
// January 2001). A negative value means the receiver will be earlier than the
// reference date.
//
// # Return Value
//
// An [NSDate] object initialized relative to the absolute reference date by
// `ti` seconds.
//
// # Discussion
//
// This method is a designated initializer for the [NSDate] class and is
// declared primarily for the use of subclasses of [NSDate]. When you subclass
// [NSDate] to create a concrete date class, you must override this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSinceReferenceDate:)
func (d NSDate) InitWithTimeIntervalSinceReferenceDate(ti float64) NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("initWithTimeIntervalSinceReferenceDate:"), ti)
	return rv
}

// Returns a date object initialized relative to 00:00:00 UTC on 1 January
// 1970 by a given number of seconds.
//
// secs: The number of seconds from the reference date (00:00:00 UTC on 1 January
// 1970) for the new date. Use a negative argument to specify a date and time
// before the reference date.
//
// # Return Value
//
// An [NSDate] object set to `seconds` seconds from the reference date.
//
// # Discussion
//
// This method is useful for creating [NSDate] objects from time_t values
// returned by BSD system functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeIntervalSince1970:)
func (d NSDate) InitWithTimeIntervalSince1970(secs float64) NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("initWithTimeIntervalSince1970:"), secs)
	return rv
}

// Returns a date object initialized from data in the given unarchiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/init(coder:)
func (d NSDate) InitWithCoder(coder INSCoder) NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns a Boolean value that indicates whether a given object is a date
// that is exactly equal the receiver.
//
// otherDate: The date to compare with the receiver.
//
// # Return Value
//
// true if the `otherDate` is an [NSDate] object and is exactly equal to the
// receiver, otherwise false.
//
// # Discussion
//
// This method detects sub-second differences between dates. If you want to
// compare dates with a less fine granularity, use [TimeIntervalSinceDate] to
// compare the two dates.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/isEqual(to:)
func (d NSDate) IsEqualToDate(otherDate INSDate) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEqualToDate:"), otherDate)
	return rv
}

// Returns the earlier of the receiver and another given date.
//
// anotherDate: The date with which to compare the receiver.
//
// # Return Value
//
// The earlier of the receiver and `anotherDate`, determined using
// [TimeIntervalSinceDate]. If the receiver and `anotherDate` represent the
// same date, returns the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/earlierDate(_:)
func (d NSDate) EarlierDate(anotherDate INSDate) INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("earlierDate:"), anotherDate)
	return NSDateFromID(rv)
}

// Returns the later of the receiver and another given date.
//
// anotherDate: The date with which to compare the receiver.
//
// # Return Value
//
// The later of the receiver and `anotherDate`, determined using
// [TimeIntervalSinceDate]. If the receiver and `anotherDate` represent the
// same date, returns the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/laterDate(_:)
func (d NSDate) LaterDate(anotherDate INSDate) INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("laterDate:"), anotherDate)
	return NSDateFromID(rv)
}

// Indicates the temporal ordering of the receiver and another given date.
//
// other: The date with which to compare the receiver.
//
// This value must not be `nil`. If the value is `nil`, the behavior is
// undefined and may change in future versions of macOS.
//
// # Return Value
//
// If:
//
// - The receiver and `anotherDate` are exactly equal to each other,
// [NSOrderedSame] - The receiver is later in time than `anotherDate`,
// [NSOrderedDescending] - The receiver is earlier in time than `anotherDate`,
// [NSOrderedAscending].
//
// # Discussion
//
// This method detects sub-second differences between dates. If you want to
// compare dates with a less fine granularity, use [TimeIntervalSinceDate] to
// compare the two dates.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/compare(_:)
func (d NSDate) Compare(other INSDate) ComparisonResult {
	rv := objc.Send[ComparisonResult](d.ID, objc.Sel("compare:"), other)
	return ComparisonResult(rv)
}

// Returns the interval between the receiver and another given date.
//
// anotherDate: The date with which to compare the receiver. You must pass a non-`nil` date
// object.
//
// # Return Value
//
// The interval between the receiver and the `anotherDate` parameter. If the
// receiver is earlier than `anotherDate`, the return value is negative. If
// `anotherDate` is `nil`, the results are undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSince(_:)
func (d NSDate) TimeIntervalSinceDate(anotherDate INSDate) float64 {
	rv := objc.Send[NSTimeInterval](d.ID, objc.Sel("timeIntervalSinceDate:"), anotherDate)
	return float64(rv)
}

// Returns a new date object that is set to a given number of seconds relative
// to the receiver.
//
// ti: The number of seconds to add to the receiver. Use a negative value for
// seconds to have the returned object specify a date before the receiver.
//
// # Return Value
//
// A new [NSDate] object that is set to `seconds` seconds relative to the
// receiver. The date returned might have a representation different from the
// receiver’s.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/addingTimeInterval(_:)
func (d NSDate) DateByAddingTimeInterval(ti float64) INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dateByAddingTimeInterval:"), ti)
	return NSDateFromID(rv)
}

// Returns a string representation of the date using the given locale.
//
// locale: An [NSLocale] object.
//
// If you pass `nil`, [NSDate] formats the date in the same way as the
// [Description] property.
//
// In OS X v10.4 and earlier, this parameter was an [NSDictionary] object. If
// you pass in an [NSDictionary] object in OS X v10.5, [NSDate] uses the
// default user locale—the same as if you passed in `[NSLocale
// currentLocale].`
//
// # Return Value
//
// A string representation of the receiver, using the given locale, or if the
// locale argument is `nil`, in the international format `YYYY-MM-DD HH:MM:SS
// ±HHMM`, where `±HHMM` represents the time zone offset in hours and
// minutes from UTC (for example, “`2001-03-24 10:45:32 +0600`”)
//
// # Discussion
//
// In OS X v10.4 and earlier, `localeDictionary` is an [NSDictionary] object
// containing locale data. To use the user’s preferences, you can use
// `[[NSUserDefaults standardUserDefaults] dictionaryRepresentation].`
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/description(with:)
func (d NSDate) DescriptionWithLocale(locale objectivec.IObject) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("descriptionWithLocale:"), locale)
	return NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Foundation/NSDate/init(timeInterval:since:)
func (d NSDate) InitWithTimeIntervalSinceDate(secsToBeAdded float64, date INSDate) NSDate {
	rv := objc.Send[NSDate](d.ID, objc.Sel("initWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (d NSDate) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns a new date object set to the current date and time.
//
// # Return Value
//
// A new date object set to the current date and time.
//
// # Discussion
//
// This method uses the default initializer method for the class, [Init].
//
// The following code sample shows how to use [Date] to get the current date:
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/date
func (_NSDateClass NSDateClass) Date() NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("date"))
	return NSDateFromID(rv)
}

// Creates and returns a date object set to the given number of seconds from
// 00:00:00 UTC on 1 January 1970.
//
// secs: The number of seconds from the reference date (00:00:00 UTC on 1 January
// 1970) for the new date. Use a negative argument to specify a date and time
// before the reference date.
//
// # Return Value
//
// An [NSDate] object set to `secs` seconds from the reference date.
//
// # Discussion
//
// This method is useful for creating [NSDate] objects from time_t values
// returned by BSD system functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeIntervalSince1970:
func (_NSDateClass NSDateClass) DateWithTimeIntervalSince1970(secs float64) NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("dateWithTimeIntervalSince1970:"), secs)
	return NSDateFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeInterval:sinceDate:
func (_NSDateClass NSDateClass) DateWithTimeIntervalSinceDate(secsToBeAdded float64, date INSDate) NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("dateWithTimeInterval:sinceDate:"), secsToBeAdded, date)
	return NSDateFromID(rv)
}

// Creates and returns a date object set to a given number of seconds from the
// current date and time.
//
// secs: The number of seconds from the current date and time for the new date. Use
// a negative value to specify a date before the current date.
//
// # Return Value
//
// An [NSDate] object set to `secs` seconds from the current date and time.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeIntervalSinceNow:
func (_NSDateClass NSDateClass) DateWithTimeIntervalSinceNow(secs float64) NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("dateWithTimeIntervalSinceNow:"), secs)
	return NSDateFromID(rv)
}

// Creates and returns a date object set to a given number of seconds from
// 00:00:00 UTC on 1 January 2001.
//
// ti: The number of seconds from the absolute reference date (00:00:00 UTC on 1
// January 2001) for the new date. Use a negative argument to specify a date
// and time before the reference date.
//
// # Return Value
//
// An [NSDate] object set to `ti` seconds from the absolute reference date.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/dateWithTimeIntervalSinceReferenceDate:
func (_NSDateClass NSDateClass) DateWithTimeIntervalSinceReferenceDate(ti float64) NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("dateWithTimeIntervalSinceReferenceDate:"), ti)
	return NSDateFromID(rv)
}

// The interval between the date object and the current date and time.
//
// # Discussion
//
// If the date object is earlier than the current date and time, this
// property’s value is negative.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSinceNow
func (d NSDate) TimeIntervalSinceNow() float64 {
	rv := objc.Send[NSTimeInterval](d.ID, objc.Sel("timeIntervalSinceNow"))
	return float64(rv)
}

// The interval between the date object and 00:00:00 UTC on 1 January 2001.
//
// # Discussion
//
// This property’s value is negative if the date object is earlier than the
// system’s absolute reference date (00:00:00 UTC on 1 January 2001).
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSinceReferenceDate-swift.property
func (d NSDate) TimeIntervalSinceReferenceDate() float64 {
	rv := objc.Send[NSTimeInterval](d.ID, objc.Sel("timeIntervalSinceReferenceDate"))
	return float64(rv)
}

// The interval between the date object and 00:00:00 UTC on 1 January 1970.
//
// # Discussion
//
// This property’s value is negative if the date object is earlier than
// 00:00:00 UTC on 1 January 1970.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/timeIntervalSince1970
func (d NSDate) TimeIntervalSince1970() float64 {
	rv := objc.Send[NSTimeInterval](d.ID, objc.Sel("timeIntervalSince1970"))
	return float64(rv)
}

// The number of seconds from 1 January 1970 to the reference date, 1 January
// 2001.
//
// See: https://developer.apple.com/documentation/foundation/nstimeintervalsince1970
func (d NSDate) NSTimeIntervalSince1970() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("NSTimeIntervalSince1970"))
	return rv
}
func (d NSDate) SetNSTimeIntervalSince1970(value float64) {
	objc.Send[struct{}](d.ID, objc.Sel("setNSTimeIntervalSince1970:"), value)
}

// A string representation of the date object.
//
// # Discussion
//
// The representation is useful for debugging only.
//
// There are a number of options to acquire a formatted string for a date
// including: date formatters (see [NSDateFormatter] and [Data Formatting
// Guide]), and the [NSDate] methods [DescriptionWithLocale],
// [date(withCalendarFormat:timeZone:)], and
// [description(withCalendarFormat:timeZone:locale:)]
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/description
//
// [Data Formatting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/DataFormatting/DataFormatting.html#//apple_ref/doc/uid/10000029i
// [date(withCalendarFormat:timeZone:)]: https://developer.apple.com/documentation/Foundation/NSDate/date(withCalendarFormat:timeZone:)
// [description(withCalendarFormat:timeZone:locale:)]: https://developer.apple.com/documentation/Foundation/NSDate/description(withCalendarFormat:timeZone:locale:)
func (d NSDate) Description() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}

// A custom playground Quick Look for this object.
//
// See: https://developer.apple.com/documentation/foundation/nsdate/customplaygroundquicklook
func (d NSDate) CustomPlaygroundQuickLook() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("customPlaygroundQuickLook"))
	return objectivec.Object{ID: rv}
}
func (d NSDate) SetCustomPlaygroundQuickLook(value objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setCustomPlaygroundQuickLook:"), value)
}

// A notification posted whenever the system clock is changed.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nssystemclockdidchange
func (d NSDate) NSSystemClockDidChange() NSNotificationName {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("NSSystemClockDidChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// A date object representing a date in the distant future.
//
// # Return Value
//
// An [NSDate] object representing a date in the distant future (in terms of
// centuries).
//
// # Discussion
//
// You can pass this value when an [NSDate] object is required to have the
// date argument essentially ignored. For example, the [NSWindow] method
// [nextEvent(matching:until:inMode:dequeue:)] returns `nil` if an event
// specified in the event mask does not happen before the specified date. You
// can use the object returned by [DistantFuture] as the date argument to wait
// indefinitely for the event to occur.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/distantFuture
//
// [NSWindow]: https://developer.apple.com/documentation/AppKit/NSWindow
// [nextEvent(matching:until:inMode:dequeue:)]: https://developer.apple.com/documentation/AppKit/NSWindow/nextEvent(matching:until:inMode:dequeue:)
func (_NSDateClass NSDateClass) DistantFuture() NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("distantFuture"))
	return NSDateFromID(objc.ID(rv))
}

// A date object representing a date in the distant past.
//
// # Return Value
//
// An [NSDate] object representing a date in the distant past (in terms of
// centuries).
//
// # Discussion
//
// You can use this object as a control date, a guaranteed temporal boundary.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/distantPast
func (_NSDateClass NSDateClass) DistantPast() NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("distantPast"))
	return NSDateFromID(objc.ID(rv))
}

// The current date and time, as of the time of access.
//
// # Discussion
//
// This is equivalent to initializing a new instance with `NSDate()` (or
// `[[NSDate alloc] init]` in Objective-C). The [NSDate] instance doesn’t
// automatically update its time after you retrieve it.
//
// See: https://developer.apple.com/documentation/Foundation/NSDate/now
func (_NSDateClass NSDateClass) Now() NSDate {
	rv := objc.Send[objc.ID](objc.ID(_NSDateClass.class), objc.Sel("now"))
	return NSDateFromID(objc.ID(rv))
}

// Protocol methods for NSCopying

// Protocol methods for NSSecureCoding
