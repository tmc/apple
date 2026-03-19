// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTimeZone] class.
var (
	_NSTimeZoneClass     NSTimeZoneClass
	_NSTimeZoneClassOnce sync.Once
)

func getNSTimeZoneClass() NSTimeZoneClass {
	_NSTimeZoneClassOnce.Do(func() {
		_NSTimeZoneClass = NSTimeZoneClass{class: objc.GetClass("NSTimeZone")}
	})
	return _NSTimeZoneClass
}

// GetNSTimeZoneClass returns the class object for NSTimeZone.
func GetNSTimeZoneClass() NSTimeZoneClass {
	return getNSTimeZoneClass()
}

type NSTimeZoneClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTimeZoneClass) Alloc() NSTimeZone {
	rv := objc.Send[NSTimeZone](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Information about standard time conventions associated with a specific
// geopolitical region.
//
// # Overview
// 
// In Swift, this type bridges to [TimeZone]; use [NSTimeZone] when you need
// reference semantics or other Foundation-specific behavior.
// 
// Time zones represent the standard time policies for a geopolitical region.
// Time zones have identifiers like “America/Los_Angeles” and can also be
// identified by abbreviations, such as PST for Pacific Standard Time. You can
// create time zone objects by ID with [NSTimeZone.InitWithName] and by abbreviation with
// [NSTimeZone.TimeZoneWithAbbreviation].
// 
// Time zones can also represent a temporal offset—either plus or
// minus—from Greenwich Mean Time (GMT). For example, the temporal offset of
// Pacific Standard Time is 8 hours behind Greenwich Mean Time (GMT-8). You
// can create time zone objects with a temporal offset by using
// [NSTimeZone.TimeZoneForSecondsFromGMT].
// 
// You typically work with system time zones rather than creating time zones
// by identifier or by offset. The [NSTimeZone.SystemTimeZone] class property returns the
// time zone currently used by the system, if known. This value is cached once
// the property is accessed and doesn’t reflect any system time zone changes
// until you call the [NSTimeZone.ResetSystemTimeZone] method. The [NSTimeZone.LocalTimeZone] class
// property returns an autoupdating proxy object that always returns the
// current time zone used by the system. You can also set the
// [NSTimeZone.DefaultTimeZone] class property to make your app run as if it were in a
// different time zone than the system.
// 
// [NSTimeZone] is with its Core Foundation counterpart, [CFTimeZone]. See
// [Toll-Free Bridging] for more information on toll-free bridging.
//
// [CFTimeZone]: https://developer.apple.com/documentation/CoreFoundation/CFTimeZone
// [TimeZone]: https://developer.apple.com/documentation/Foundation/TimeZone
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating Time Zones
//
//   - [NSTimeZone.InitWithName]: Returns a time zone initialized with a given identifier.
//   - [NSTimeZone.InitWithNameData]: Initializes a time zone with a given identifier and time zone data.
//
// # Getting Time Zone Information
//
//   - [NSTimeZone.Name]: The geopolitical region ID that identifies the receiver.
//   - [NSTimeZone.Abbreviation]: The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).
//   - [NSTimeZone.AbbreviationForDate]: Returns the abbreviation for the receiver at a given date.
//   - [NSTimeZone.SecondsFromGMT]: The current difference in seconds between the receiver and Greenwich Mean Time.
//   - [NSTimeZone.SecondsFromGMTForDate]: Returns the difference in seconds between the receiver and Greenwich Mean Time at a given date.
//   - [NSTimeZone.Data]: The data that stores the information used by the receiver.
//
// # Working with Daylight Savings
//
//   - [NSTimeZone.DaylightSavingTime]: A Boolean value that indicates whether the receiver is currently using daylight saving time.
//   - [NSTimeZone.IsDaylightSavingTimeForDate]: Indicates whether the receiver uses daylight saving time on a given date.
//   - [NSTimeZone.DaylightSavingTimeOffset]: The current daylight saving time offset of the receiver.
//   - [NSTimeZone.DaylightSavingTimeOffsetForDate]: Returns the daylight saving time offset for a given date.
//   - [NSTimeZone.NextDaylightSavingTimeTransition]: The date of the next daylight saving time transition for the receiver.
//   - [NSTimeZone.NextDaylightSavingTimeTransitionAfterDate]: Returns the next daylight saving time transition after a given date.
//
// # Comparing Time Zones
//
//   - [NSTimeZone.IsEqualToTimeZone]: Indicates whether the receiver has the same name and data as the specified time zone.
//
// # Describing Time Zones
//
//   - [NSTimeZone.LocalizedNameLocale]: Returns the localized name of the time zone.
//   - [NSTimeZone.Description]: A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
//
// # Recognizing Notifications
//
//   - [NSTimeZone.NSSystemTimeZoneDidChange]: A notification posted when the time zone changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone
type NSTimeZone struct {
	objectivec.Object
}

// NSTimeZoneFromID constructs a [NSTimeZone] from an objc.ID.
//
// Information about standard time conventions associated with a specific
// geopolitical region.
func NSTimeZoneFromID(id objc.ID) NSTimeZone {
	return NSTimeZone{objectivec.Object{ID: id}}
}
// NOTE: NSTimeZone adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTimeZone] class.
//
// # Creating Time Zones
//
//   - [INSTimeZone.InitWithName]: Returns a time zone initialized with a given identifier.
//   - [INSTimeZone.InitWithNameData]: Initializes a time zone with a given identifier and time zone data.
//
// # Getting Time Zone Information
//
//   - [INSTimeZone.Name]: The geopolitical region ID that identifies the receiver.
//   - [INSTimeZone.Abbreviation]: The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).
//   - [INSTimeZone.AbbreviationForDate]: Returns the abbreviation for the receiver at a given date.
//   - [INSTimeZone.SecondsFromGMT]: The current difference in seconds between the receiver and Greenwich Mean Time.
//   - [INSTimeZone.SecondsFromGMTForDate]: Returns the difference in seconds between the receiver and Greenwich Mean Time at a given date.
//   - [INSTimeZone.Data]: The data that stores the information used by the receiver.
//
// # Working with Daylight Savings
//
//   - [INSTimeZone.DaylightSavingTime]: A Boolean value that indicates whether the receiver is currently using daylight saving time.
//   - [INSTimeZone.IsDaylightSavingTimeForDate]: Indicates whether the receiver uses daylight saving time on a given date.
//   - [INSTimeZone.DaylightSavingTimeOffset]: The current daylight saving time offset of the receiver.
//   - [INSTimeZone.DaylightSavingTimeOffsetForDate]: Returns the daylight saving time offset for a given date.
//   - [INSTimeZone.NextDaylightSavingTimeTransition]: The date of the next daylight saving time transition for the receiver.
//   - [INSTimeZone.NextDaylightSavingTimeTransitionAfterDate]: Returns the next daylight saving time transition after a given date.
//
// # Comparing Time Zones
//
//   - [INSTimeZone.IsEqualToTimeZone]: Indicates whether the receiver has the same name and data as the specified time zone.
//
// # Describing Time Zones
//
//   - [INSTimeZone.LocalizedNameLocale]: Returns the localized name of the time zone.
//   - [INSTimeZone.Description]: A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
//
// # Recognizing Notifications
//
//   - [INSTimeZone.NSSystemTimeZoneDidChange]: A notification posted when the time zone changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone
type INSTimeZone interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating Time Zones

	// Returns a time zone initialized with a given identifier.
	InitWithName(tzName string) NSTimeZone
	// Initializes a time zone with a given identifier and time zone data.
	InitWithNameData(tzName string, aData INSData) NSTimeZone

	// Topic: Getting Time Zone Information

	// The geopolitical region ID that identifies the receiver.
	Name() string
	// The abbreviation for the receiver, such as “EDT” (Eastern Daylight Time).
	Abbreviation() string
	// Returns the abbreviation for the receiver at a given date.
	AbbreviationForDate(aDate INSDate) string
	// The current difference in seconds between the receiver and Greenwich Mean Time.
	SecondsFromGMT() int
	// Returns the difference in seconds between the receiver and Greenwich Mean Time at a given date.
	SecondsFromGMTForDate(aDate INSDate) int
	// The data that stores the information used by the receiver.
	Data() INSData

	// Topic: Working with Daylight Savings

	// A Boolean value that indicates whether the receiver is currently using daylight saving time.
	DaylightSavingTime() bool
	// Indicates whether the receiver uses daylight saving time on a given date.
	IsDaylightSavingTimeForDate(aDate INSDate) bool
	// The current daylight saving time offset of the receiver.
	DaylightSavingTimeOffset() float64
	// Returns the daylight saving time offset for a given date.
	DaylightSavingTimeOffsetForDate(aDate INSDate) float64
	// The date of the next daylight saving time transition for the receiver.
	NextDaylightSavingTimeTransition() INSDate
	// Returns the next daylight saving time transition after a given date.
	NextDaylightSavingTimeTransitionAfterDate(aDate INSDate) INSDate

	// Topic: Comparing Time Zones

	// Indicates whether the receiver has the same name and data as the specified time zone.
	IsEqualToTimeZone(aTimeZone INSTimeZone) bool

	// Topic: Describing Time Zones

	// Returns the localized name of the time zone.
	LocalizedNameLocale(style NSTimeZoneNameStyle, locale INSLocale) string
	// A textual description of the time zone including the name, abbreviation, offset from GMT, and whether or not daylight saving time is currently in effect.
	Description() string

	// Topic: Recognizing Notifications

	// A notification posted when the time zone changes.
	NSSystemTimeZoneDidChange() NSNotificationName
}

// Init initializes the instance.
func (t NSTimeZone) Init() NSTimeZone {
	rv := objc.Send[NSTimeZone](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTimeZone) Autorelease() NSTimeZone {
	rv := objc.Send[NSTimeZone](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTimeZone creates a new NSTimeZone instance.
func NewNSTimeZone() NSTimeZone {
	class := getNSTimeZoneClass()
	rv := objc.Send[NSTimeZone](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a time zone object offset from Greenwich Mean Time by a given
// number of seconds.
//
// seconds: The number of seconds by which the new time zone is offset from GMT.
//
// # Return Value
// 
// A time zone object offset from Greenwich Mean Time by `seconds`.
//
// # Discussion
// 
// The name of the new time zone is GMT +/– the offset, in hours and
// minutes. Time zones created with this method never have daylight savings,
// and the offset is constant no matter the date.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(forSecondsFromGMT:)
func NewTimeZoneForSecondsFromGMT(seconds int) NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(getNSTimeZoneClass().class), objc.Sel("timeZoneForSecondsFromGMT:"), seconds)
	return NSTimeZoneFromID(rv)
}

// Returns the time zone object identified by a given abbreviation.
//
// abbreviation: An abbreviation for a time zone.
//
// # Return Value
// 
// The time zone object identified by `abbreviation` determined by resolving
// the abbreviation to a name using the abbreviation dictionary and then
// returning the time zone for that name. Returns `nil` if there is no match
// for `abbreviation`.
//
// # Discussion
// 
// In general, you are discouraged from using abbreviations except for unique
// instances such as “GMT”. Time Zone abbreviations are not standardized
// and so a given abbreviation may have multiple meanings—for example,
// “EST” refers to Eastern Time in both the United States and Australia
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(abbreviation:)
func NewTimeZoneWithAbbreviation(abbreviation string) NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(getNSTimeZoneClass().class), objc.Sel("timeZoneWithAbbreviation:"), objc.String(abbreviation))
	return NSTimeZoneFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewTimeZoneWithCoder(coder INSCoder) NSTimeZone {
	instance := getNSTimeZoneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTimeZoneFromID(rv)
}

// Returns a time zone initialized with a given identifier.
//
// tzName: The identifier for the time zone. Providing `nil` for this parameter raises
// an invalid argument exception.
//
// # Return Value
// 
// A time zone object initialized with the identifier `tzName`.
//
// # Discussion
// 
// If `tzName` is a known identifier, this method calls [InitWithNameData]
// with the appropriate data object.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:)
func NewTimeZoneWithName(tzName string) NSTimeZone {
	instance := getNSTimeZoneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), objc.String(tzName))
	return NSTimeZoneFromID(rv)
}

// Initializes a time zone with a given identifier and time zone data.
//
// tzName: The identifier for the time zone. Providing `nil` for this parameter raises
// an invalid argument exception.
//
// aData: This parameter is ignored.
//
// # Discussion
// 
// As of macOS 10.6, the underlying implementation of this method has been
// changed to ignore the specified `data` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:data:)
func NewTimeZoneWithNameData(tzName string, aData INSData) NSTimeZone {
	instance := getNSTimeZoneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:data:"), objc.String(tzName), aData)
	return NSTimeZoneFromID(rv)
}

// Returns a time zone initialized with a given identifier.
//
// tzName: The identifier for the time zone. Providing `nil` for this parameter raises
// an invalid argument exception.
//
// # Return Value
// 
// A time zone object initialized with the identifier `tzName`.
//
// # Discussion
// 
// If `tzName` is a known identifier, this method calls [InitWithNameData]
// with the appropriate data object.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:)
func (t NSTimeZone) InitWithName(tzName string) NSTimeZone {
	rv := objc.Send[NSTimeZone](t.ID, objc.Sel("initWithName:"), objc.String(tzName))
	return rv
}
// Initializes a time zone with a given identifier and time zone data.
//
// tzName: The identifier for the time zone. Providing `nil` for this parameter raises
// an invalid argument exception.
//
// aData: This parameter is ignored.
//
// # Discussion
// 
// As of macOS 10.6, the underlying implementation of this method has been
// changed to ignore the specified `data` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/init(name:data:)
func (t NSTimeZone) InitWithNameData(tzName string, aData INSData) NSTimeZone {
	rv := objc.Send[NSTimeZone](t.ID, objc.Sel("initWithName:data:"), objc.String(tzName), aData)
	return rv
}
// Returns the abbreviation for the receiver at a given date.
//
// aDate: The date for which to get the abbreviation for the receiver.
//
// # Return Value
// 
// The abbreviation for the receiver at `aDate`.
//
// # Discussion
// 
// Note that the abbreviation may be different at different dates. For
// example, during daylight saving time the US/Eastern time zone has an
// abbreviation of “EDT.” At other times, its abbreviation is “EST.”
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviation(for:)
func (t NSTimeZone) AbbreviationForDate(aDate INSDate) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("abbreviationForDate:"), aDate)
	return NSStringFromID(rv).String()
}
// Returns the difference in seconds between the receiver and Greenwich Mean
// Time at a given date.
//
// aDate: The date against which to test the receiver.
//
// # Return Value
// 
// The difference in seconds between the receiver and Greenwich Mean Time at
// `aDate`.
//
// # Discussion
// 
// The difference may be different from the current difference if the time
// zone changes its offset from GMT at different points in the year—for
// example, the U.S. time zones change with daylight saving time.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/secondsFromGMT(for:)
func (t NSTimeZone) SecondsFromGMTForDate(aDate INSDate) int {
	rv := objc.Send[int](t.ID, objc.Sel("secondsFromGMTForDate:"), aDate)
	return rv
}
// Indicates whether the receiver uses daylight saving time on a given date.
//
// aDate: The date against which to test the receiver.
//
// # Return Value
// 
// [true] if the receiver uses daylight saving time at `aDate`, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/isDaylightSavingTime(for:)
func (t NSTimeZone) IsDaylightSavingTimeForDate(aDate INSDate) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isDaylightSavingTimeForDate:"), aDate)
	return rv
}
// Returns the daylight saving time offset for a given date.
//
// aDate: A date.
//
// # Return Value
// 
// The daylight saving time offset for `aDate`.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/daylightSavingTimeOffset(for:)
func (t NSTimeZone) DaylightSavingTimeOffsetForDate(aDate INSDate) float64 {
	rv := objc.Send[NSTimeInterval](t.ID, objc.Sel("daylightSavingTimeOffsetForDate:"), aDate)
	return float64(rv)
}
// Returns the next daylight saving time transition after a given date.
//
// aDate: A date.
//
// # Return Value
// 
// The next daylight saving time transition after `aDate`. Depending on the
// time zone of the receiver, this method may return a change of the time
// zone’s offset from GMT. Returns `nil` if the time zone of the receiver
// does not observe daylight savings time as of `aDate`.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/nextDaylightSavingTimeTransition(after:)
func (t NSTimeZone) NextDaylightSavingTimeTransitionAfterDate(aDate INSDate) INSDate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("nextDaylightSavingTimeTransitionAfterDate:"), aDate)
	return NSDateFromID(rv)
}
// Indicates whether the receiver has the same name and data as the specified
// time zone.
//
// aTimeZone: The time zone to compare with the receiver.
//
// # Return Value
// 
// [true] if `aTimeZone` and the receiver have the same name and data,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/isEqual(to:)
func (t NSTimeZone) IsEqualToTimeZone(aTimeZone INSTimeZone) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEqualToTimeZone:"), aTimeZone)
	return rv
}
// Returns the localized name of the time zone.
//
// style: The format style for the returned string.
//
// locale: The locale for which to format the name.
//
// # Return Value
// 
// The name of the receiver localized for `locale` using `style`.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/localizedName(_:locale:)
func (t NSTimeZone) LocalizedNameLocale(style NSTimeZoneNameStyle, locale INSLocale) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("localizedName:locale:"), style, locale)
	return NSStringFromID(rv).String()
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (t NSTimeZone) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (t NSTimeZone) InitWithCoder(coder INSCoder) NSTimeZone {
	rv := objc.Send[NSTimeZone](t.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Clears any time zone value cached for the [SystemTimeZone] property.
//
// # Discussion
// 
// If the app has cached the system time zone by accessing the
// [SystemTimeZone] class property, this method clears that cached value. If
// you subsequently access the [SystemTimeZone] class property, a new time
// zone object is created and cached.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/resetSystemTimeZone()
func (_NSTimeZoneClass NSTimeZoneClass) ResetSystemTimeZone() {
	objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("resetSystemTimeZone"))
}
// Returns the time zone object identified by a given identifier.
//
// tzName: The ID for the time zone.
//
// # Return Value
// 
// The time zone in the information directory with a name matching `tzName`.
// Returns `nil` if there is no match for the name.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneWithName:
func (_NSTimeZoneClass NSTimeZoneClass) TimeZoneWithName(tzName string) NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("timeZoneWithName:"), objc.String(tzName))
	return NSTimeZoneFromID(rv)
}
// Returns the time zone with a given identifier whose data has been
// initialized using given data.
//
// tzName: The ID for the time zone.
//
// aData: This parameter is ignored.
//
// # Return Value
// 
// The time zone with the ID `tzName`.
//
// # Discussion
// 
// As of macOS 10.6, the underlying implementation of this method has been
// changed to ignore the specified `data` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneWithName:data:
func (_NSTimeZoneClass NSTimeZoneClass) TimeZoneWithNameData(tzName string, aData INSData) NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("timeZoneWithName:data:"), objc.String(tzName), aData)
	return NSTimeZoneFromID(rv)
}

// The geopolitical region ID that identifies the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/name
func (t NSTimeZone) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
// The abbreviation for the receiver, such as “EDT” (Eastern Daylight
// Time).
//
// # Discussion
// 
// Invokes [AbbreviationForDate] with the current date as the argument.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviation
func (t NSTimeZone) Abbreviation() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("abbreviation"))
	return NSStringFromID(rv).String()
}
// The current difference in seconds between the receiver and Greenwich Mean
// Time.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/secondsFromGMT
func (t NSTimeZone) SecondsFromGMT() int {
	rv := objc.Send[int](t.ID, objc.Sel("secondsFromGMT"))
	return rv
}
// The data that stores the information used by the receiver.
//
// # Discussion
// 
// Treat this data as an opaque object.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/data
func (t NSTimeZone) Data() INSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("data"))
	return NSDataFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the receiver is currently using
// daylight saving time.
//
// # Discussion
// 
// If [true], the receiver is currently using daylight saving time, otherwise
// [false]. This property invokes [IsDaylightSavingTimeForDate] with the
// current date as the argument.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/isDaylightSavingTime
func (t NSTimeZone) DaylightSavingTime() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isDaylightSavingTime"))
	return rv
}
// The current daylight saving time offset of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/daylightSavingTimeOffset
func (t NSTimeZone) DaylightSavingTimeOffset() float64 {
	rv := objc.Send[NSTimeInterval](t.ID, objc.Sel("daylightSavingTimeOffset"))
	return float64(rv)
}
// The date of the next daylight saving time transition for the receiver.
//
// # Discussion
// 
// This property contains the date of the next (after the current instant)
// daylight saving time transition for the receiver. Depending on the time
// zone of the receiver, the value of this property may represent a change of
// the time zone’s offset from GMT. Returns `nil` if the time zone of the
// receiver does not currently observe daylight saving time.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/nextDaylightSavingTimeTransition
func (t NSTimeZone) NextDaylightSavingTimeTransition() INSDate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("nextDaylightSavingTimeTransition"))
	return NSDateFromID(objc.ID(rv))
}
// A textual description of the time zone including the name, abbreviation,
// offset from GMT, and whether or not daylight saving time is currently in
// effect.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/description
func (t NSTimeZone) Description() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}
// A notification posted when the time zone changes.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nssystemtimezonedidchange
func (t NSTimeZone) NSSystemTimeZoneDidChange() NSNotificationName {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSSystemTimeZoneDidChangeNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// An object that tracks the current system time zone.
//
// # Discussion
// 
// Use this property when you want an object that always reflects the current
// system time zone. Contrast this behavior with that of the [SystemTimeZone]
// class property, which has its value cached until you manually clear it by
// calling the [ResetSystemTimeZone] method.
// 
// Although the time zone obtained here automatically updates with the system,
// it provides no indication when system settings change. To receive
// notification of time zone changes, add an observer to the
// [NSSystemTimeZoneDidChange] notification by using the
// [AddObserverSelectorNameObject].
//
// [NSSystemTimeZoneDidChange]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSSystemTimeZoneDidChange
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/local
func (_NSTimeZoneClass NSTimeZoneClass) LocalTimeZone() NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("localTimeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
// The time zone currently used by the system.
//
// # Discussion
// 
// If the current system time zone cannot be determined, the GMT time zone is
// used instead.
// 
// If you access the [SystemTimeZone] class property, its value is cached by
// the app and doesn’t update if the user subsequently changes the system
// time zone. In order for the [SystemTimeZone] property to reflect the new
// time zone, you must first call the [ResetSystemTimeZone] method to clear
// the cached value. Then, the next time you access the [SystemTimeZone]
// property, it returns the current system time zone, and caches that value.
// 
// If you access the [SystemTimeZone] class property, assign its value to a
// variable, and clear the cached value for the property by calling the
// [ResetSystemTimeZone] method, the object stored in the variable doesn’t
// update to reflect the new system time zone. Contrast this behavior with
// that of the [LocalTimeZone] class property, which returns a proxy object
// that always reflects the current system time zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/system
func (_NSTimeZoneClass NSTimeZoneClass) SystemTimeZone() NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("systemTimeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
// The default time zone for the current app.
//
// # Discussion
// 
// If no [DefaultTimeZone] time zone has been set, the current system time
// zone is used. If the current system time zone cannot be determined, the GMT
// time zone is used instead.
// 
// The [DefaultTimeZone] time zone is used by the app for date and time
// operations. You can set it to cause the app to run as if it were in a
// different time zone. Setting the [DefaultTimeZone] property clears any
// value that was previously set.
// 
// If you access the [DefaultTimeZone] class property, assign its value to a
// variable, and set a new [DefaultTimeZone] time zone, the object stored in
// the variable doesn’t update to reflect the new [DefaultTimeZone] time
// zone. Contrast this behavior with that of the [LocalTimeZone] class
// property, which returns a proxy object that always reflects the current
// system time zone.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/default
func (_NSTimeZoneClass NSTimeZoneClass) DefaultTimeZone() NSTimeZone {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("defaultTimeZone"))
	return NSTimeZoneFromID(objc.ID(rv))
}
func (_NSTimeZoneClass NSTimeZoneClass) SetDefaultTimeZone(value NSTimeZone) {
	objc.Send[struct{}](objc.ID(_NSTimeZoneClass.class), objc.Sel("setDefaultTimeZone:"), value)
}
// Returns an array of strings listing the IDs of all the time zones known to
// the system.
//
// # Return Value
// 
// An array of strings listing the IDs of all the time zones known to the
// system.
// 
// # Discussion
// 
// An array of strings listing the IDs of all the time zones known to the
// system.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/knownTimeZoneNames
func (_NSTimeZoneClass NSTimeZoneClass) KnownTimeZoneNames() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("knownTimeZoneNames"))
	return objc.ConvertSliceToStrings(rv)
}
// Returns a dictionary holding the mappings of time zone abbreviations to
// time zone names.
//
// # Return Value
// 
// A dictionary holding the mappings of time zone abbreviations to time zone
// names.
// 
// # Discussion
// 
// Note that more than one time zone may have the same abbreviation—for
// example, US/Pacific and Canada/Pacific both use the abbreviation “PST.”
// In these cases, [AbbreviationDictionary] chooses a single name to map the
// abbreviation to.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/abbreviationDictionary
func (_NSTimeZoneClass NSTimeZoneClass) AbbreviationDictionary() INSDictionary {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("abbreviationDictionary"))
	return NSDictionaryFromID(objc.ID(rv))
}
func (_NSTimeZoneClass NSTimeZoneClass) SetAbbreviationDictionary(value INSDictionary) {
	objc.Send[struct{}](objc.ID(_NSTimeZoneClass.class), objc.Sel("setAbbreviationDictionary:"), value)
}
// Returns the time zone data version.
//
// # Return Value
// 
// A string containing the time zone data version.
//
// See: https://developer.apple.com/documentation/Foundation/NSTimeZone/timeZoneDataVersion
func (_NSTimeZoneClass NSTimeZoneClass) TimeZoneDataVersion() string {
	rv := objc.Send[objc.ID](objc.ID(_NSTimeZoneClass.class), objc.Sel("timeZoneDataVersion"))
	return NSStringFromID(rv).String()
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

