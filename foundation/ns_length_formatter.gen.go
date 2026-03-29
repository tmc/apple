// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [LengthFormatter] class.
var (
	_LengthFormatterClass     LengthFormatterClass
	_LengthFormatterClassOnce sync.Once
)

func getLengthFormatterClass() LengthFormatterClass {
	_LengthFormatterClassOnce.Do(func() {
		_LengthFormatterClass = LengthFormatterClass{class: objc.GetClass("NSLengthFormatter")}
	})
	return _LengthFormatterClass
}

// GetLengthFormatterClass returns the class object for NSLengthFormatter.
func GetLengthFormatterClass() LengthFormatterClass {
	return getLengthFormatterClass()
}

type LengthFormatterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LengthFormatterClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LengthFormatterClass) Alloc() LengthFormatter {
	rv := objc.Send[LengthFormatter](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that provides localized descriptions of linear distances, such
// as length and height measurements.
//
// # Overview
//
// # Formatting Length Strings
//
//   - [LengthFormatter.ForPersonHeightUse]: A Boolean value that indicates whether the resulting string represents a person’s height.
//   - [LengthFormatter.SetForPersonHeightUse]
//   - [LengthFormatter.NumberFormatter]: The number formatter used to format the numbers in length strings.
//   - [LengthFormatter.SetNumberFormatter]
//   - [LengthFormatter.StringFromMeters]: Returns a length string for the provided value.
//   - [LengthFormatter.StringFromValueUnit]: Returns a properly formatted length string for the given value and unit.
//   - [LengthFormatter.UnitStringFromMetersUsedUnit]: Returns the unit string for the provided value.
//   - [LengthFormatter.UnitStringFromValueUnit]: Returns the unit string based on the provided value and unit.
//   - [LengthFormatter.UnitStyle]: The unit style used by this formatter.
//   - [LengthFormatter.SetUnitStyle]
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter
type LengthFormatter struct {
	NSFormatter
}

// LengthFormatterFromID constructs a [LengthFormatter] from an objc.ID.
//
// A formatter that provides localized descriptions of linear distances, such
// as length and height measurements.
func LengthFormatterFromID(id objc.ID) LengthFormatter {
	return NSLengthFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSLengthFormatterFromID is an alias for [LengthFormatterFromID] for cross-framework compatibility.
func NSLengthFormatterFromID(id objc.ID) LengthFormatter { return LengthFormatterFromID(id) }
// NOTE: LengthFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [LengthFormatter] class.
//
// # Formatting Length Strings
//
//   - [ILengthFormatter.ForPersonHeightUse]: A Boolean value that indicates whether the resulting string represents a person’s height.
//   - [ILengthFormatter.SetForPersonHeightUse]
//   - [ILengthFormatter.NumberFormatter]: The number formatter used to format the numbers in length strings.
//   - [ILengthFormatter.SetNumberFormatter]
//   - [ILengthFormatter.StringFromMeters]: Returns a length string for the provided value.
//   - [ILengthFormatter.StringFromValueUnit]: Returns a properly formatted length string for the given value and unit.
//   - [ILengthFormatter.UnitStringFromMetersUsedUnit]: Returns the unit string for the provided value.
//   - [ILengthFormatter.UnitStringFromValueUnit]: Returns the unit string based on the provided value and unit.
//   - [ILengthFormatter.UnitStyle]: The unit style used by this formatter.
//   - [ILengthFormatter.SetUnitStyle]
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter
type ILengthFormatter interface {
	INSFormatter

	// Topic: Formatting Length Strings

	// A Boolean value that indicates whether the resulting string represents a person’s height.
	ForPersonHeightUse() bool
	SetForPersonHeightUse(value bool)
	// The number formatter used to format the numbers in length strings.
	NumberFormatter() INSNumberFormatter
	SetNumberFormatter(value INSNumberFormatter)
	// Returns a length string for the provided value.
	StringFromMeters(numberInMeters float64) string
	// Returns a properly formatted length string for the given value and unit.
	StringFromValueUnit(value float64, unit NSLengthFormatterUnit) string
	// Returns the unit string for the provided value.
	UnitStringFromMetersUsedUnit(numberInMeters float64, unitp NSLengthFormatterUnit) string
	// Returns the unit string based on the provided value and unit.
	UnitStringFromValueUnit(value float64, unit NSLengthFormatterUnit) string
	// The unit style used by this formatter.
	UnitStyle() NSFormattingUnitStyle
	SetUnitStyle(value NSFormattingUnitStyle)
}

// Init initializes the instance.
func (l LengthFormatter) Init() LengthFormatter {
	rv := objc.Send[LengthFormatter](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LengthFormatter) Autorelease() LengthFormatter {
	rv := objc.Send[LengthFormatter](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLengthFormatter creates a new LengthFormatter instance.
func NewLengthFormatter() LengthFormatter {
	class := getLengthFormatterClass()
	rv := objc.Send[LengthFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewLengthFormatterWithCoder(coder INSCoder) LengthFormatter {
	instance := getLengthFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return LengthFormatterFromID(rv)
}

// Returns a length string for the provided value.
//
// numberInMeters: The length’s value in meters.
//
// # Return Value
// 
// A string that combines a value and a unit string appropriate for the
// formatter’s locale.
//
// # Discussion
// 
// This method converts the provided length into units appropriate for the
// formatter’s locale.
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/string(fromMeters:)
func (l LengthFormatter) StringFromMeters(numberInMeters float64) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("stringFromMeters:"), numberInMeters)
	return NSStringFromID(rv).String()
}
// Returns a properly formatted length string for the given value and unit.
//
// value: The length’s value in the given unit.
//
// unit: The unit used in the resulting length string.
//
// # Return Value
// 
// A localized string that combines the provided value and unit.
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/string(fromValue:unit:)
func (l LengthFormatter) StringFromValueUnit(value float64, unit NSLengthFormatterUnit) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return NSStringFromID(rv).String()
}
// Returns the unit string for the provided value.
//
// numberInMeters: The length’s value in meters.
//
// unitp: An output parameter. This will hold the [LengthFormatter.Unit] value that
// corresponds to the returned units.
// //
// [LengthFormatter.Unit]: https://developer.apple.com/documentation/Foundation/LengthFormatter/Unit
//
// # Return Value
// 
// A localized string representing the unit.
//
// # Discussion
// 
// This method selects the correct unit based on the formatter’s locale, the
// magnitude of the value, and the [ForPersonHeightUse] property.
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromMeters:usedUnit:)
func (l LengthFormatter) UnitStringFromMetersUsedUnit(numberInMeters float64, unitp NSLengthFormatterUnit) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("unitStringFromMeters:usedUnit:"), numberInMeters, unitp)
	return NSStringFromID(rv).String()
}
// Returns the unit string based on the provided value and unit.
//
// value: The length’s value for the provided unit.
//
// unit: The unit to use in the resulting length string.
//
// # Return Value
// 
// A localized string representing the given unit. The provided value
// determines whether the unit is plural or singular.
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitString(fromValue:unit:)
func (l LengthFormatter) UnitStringFromValueUnit(value float64, unit NSLengthFormatterUnit) string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the resulting string represents a
// person’s height.
//
// # Discussion
// 
// Returns [true] if the value passed to [StringFromMeters] or
// [UnitStringFromMetersUsedUnit] is a person’s height; otherwise, [false].
// By default, this property returns [false].
// 
// The length formatter uses this property when determining the best unit for
// a given locale (for example, in the [StringFromMeters] method).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/isForPersonHeightUse
func (l LengthFormatter) ForPersonHeightUse() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("isForPersonHeightUse"))
	return rv
}
func (l LengthFormatter) SetForPersonHeightUse(value bool) {
	objc.Send[struct{}](l.ID, objc.Sel("setForPersonHeightUse:"), value)
}
// The number formatter used to format the numbers in length strings.
//
// # Discussion
// 
// This property defaults to a number formatter using the
// [NumberFormatterDecimalStyle]. You can provide a different number formatter
// to customize the length string’s appearance.
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/numberFormatter
func (l LengthFormatter) NumberFormatter() INSNumberFormatter {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("numberFormatter"))
	return NSNumberFormatterFromID(objc.ID(rv))
}
func (l LengthFormatter) SetNumberFormatter(value INSNumberFormatter) {
	objc.Send[struct{}](l.ID, objc.Sel("setNumberFormatter:"), value)
}
// The unit style used by this formatter.
//
// # Discussion
// 
// This property defaults to [FormattingUnitStyleMedium]. For a complete list
// of unit styles, see [Formatter.UnitStyle].
//
// [Formatter.UnitStyle]: https://developer.apple.com/documentation/Foundation/Formatter/UnitStyle
//
// See: https://developer.apple.com/documentation/Foundation/LengthFormatter/unitStyle
func (l LengthFormatter) UnitStyle() NSFormattingUnitStyle {
	rv := objc.Send[NSFormattingUnitStyle](l.ID, objc.Sel("unitStyle"))
	return NSFormattingUnitStyle(rv)
}
func (l LengthFormatter) SetUnitStyle(value NSFormattingUnitStyle) {
	objc.Send[struct{}](l.ID, objc.Sel("setUnitStyle:"), value)
}

