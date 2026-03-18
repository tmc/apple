// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MassFormatter] class.
var (
	_MassFormatterClass     MassFormatterClass
	_MassFormatterClassOnce sync.Once
)

func getMassFormatterClass() MassFormatterClass {
	_MassFormatterClassOnce.Do(func() {
		_MassFormatterClass = MassFormatterClass{class: objc.GetClass("NSMassFormatter")}
	})
	return _MassFormatterClass
}

// GetMassFormatterClass returns the class object for NSMassFormatter.
func GetMassFormatterClass() MassFormatterClass {
	return getMassFormatterClass()
}

type MassFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MassFormatterClass) Alloc() MassFormatter {
	rv := objc.Send[MassFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that provides localized descriptions of mass and weight values.
//
// # Overview
//
// # Formatting Mass Strings
//
//   - [MassFormatter.ForPersonMassUse]: A Boolean value that indicates whether the resulting string represents a person’s mass.
//   - [MassFormatter.SetForPersonMassUse]
//   - [MassFormatter.NumberFormatter]: The number formatter used to format the numbers in a mass strings.
//   - [MassFormatter.SetNumberFormatter]
//   - [MassFormatter.StringFromKilograms]: Returns a mass string for the provided value.
//   - [MassFormatter.StringFromValueUnit]: Returns a properly formatted mass string for the given value and unit.
//   - [MassFormatter.UnitStringFromKilogramsUsedUnit]: Returns the unit string for the provided value.
//   - [MassFormatter.UnitStringFromValueUnit]: Returns the unit string based on the provided value and unit.
//   - [MassFormatter.UnitStyle]: The unit style used by this formatter.
//   - [MassFormatter.SetUnitStyle]
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter
type MassFormatter struct {
	NSFormatter
}

// MassFormatterFromID constructs a [MassFormatter] from an objc.ID.
//
// A formatter that provides localized descriptions of mass and weight values.
func MassFormatterFromID(id objc.ID) MassFormatter {
	return NSMassFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSMassFormatterFromID is an alias for [MassFormatterFromID] for cross-framework compatibility.
func NSMassFormatterFromID(id objc.ID) MassFormatter { return MassFormatterFromID(id) }
// NOTE: MassFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MassFormatter] class.
//
// # Formatting Mass Strings
//
//   - [IMassFormatter.ForPersonMassUse]: A Boolean value that indicates whether the resulting string represents a person’s mass.
//   - [IMassFormatter.SetForPersonMassUse]
//   - [IMassFormatter.NumberFormatter]: The number formatter used to format the numbers in a mass strings.
//   - [IMassFormatter.SetNumberFormatter]
//   - [IMassFormatter.StringFromKilograms]: Returns a mass string for the provided value.
//   - [IMassFormatter.StringFromValueUnit]: Returns a properly formatted mass string for the given value and unit.
//   - [IMassFormatter.UnitStringFromKilogramsUsedUnit]: Returns the unit string for the provided value.
//   - [IMassFormatter.UnitStringFromValueUnit]: Returns the unit string based on the provided value and unit.
//   - [IMassFormatter.UnitStyle]: The unit style used by this formatter.
//   - [IMassFormatter.SetUnitStyle]
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter
type IMassFormatter interface {
	INSFormatter

	// Topic: Formatting Mass Strings

	// A Boolean value that indicates whether the resulting string represents a person’s mass.
	ForPersonMassUse() bool
	SetForPersonMassUse(value bool)
	// The number formatter used to format the numbers in a mass strings.
	NumberFormatter() INSNumberFormatter
	SetNumberFormatter(value INSNumberFormatter)
	// Returns a mass string for the provided value.
	StringFromKilograms(numberInKilograms float64) string
	// Returns a properly formatted mass string for the given value and unit.
	StringFromValueUnit(value float64, unit NSMassFormatterUnit) string
	// Returns the unit string for the provided value.
	UnitStringFromKilogramsUsedUnit(numberInKilograms float64, unitp NSMassFormatterUnit) string
	// Returns the unit string based on the provided value and unit.
	UnitStringFromValueUnit(value float64, unit NSMassFormatterUnit) string
	// The unit style used by this formatter.
	UnitStyle() NSFormattingUnitStyle
	SetUnitStyle(value NSFormattingUnitStyle)
}

// Init initializes the instance.
func (m MassFormatter) Init() MassFormatter {
	rv := objc.Send[MassFormatter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MassFormatter) Autorelease() MassFormatter {
	rv := objc.Send[MassFormatter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMassFormatter creates a new MassFormatter instance.
func NewMassFormatter() MassFormatter {
	class := getMassFormatterClass()
	rv := objc.Send[MassFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMassFormatterWithCoder(coder INSCoder) MassFormatter {
	instance := getMassFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MassFormatterFromID(rv)
}

// Returns a mass string for the provided value.
//
// numberInKilograms: The mass’s value in kilograms.
//
// # Return Value
// 
// A string that combines a value and a unit string appropriate for the
// formatter’s locale.
//
// # Discussion
// 
// This method converts the provided mass in kilograms into units appropriate
// for the formatter’s locale.
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/string(fromKilograms:)
func (m MassFormatter) StringFromKilograms(numberInKilograms float64) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringFromKilograms:"), numberInKilograms)
	return NSStringFromID(rv).String()
}

// Returns a properly formatted mass string for the given value and unit.
//
// value: The mass’s value in the given unit.
//
// unit: The unit used in the resulting mass string.
//
// # Return Value
// 
// A localized string that combines the provided value and unit.
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/string(fromValue:unit:)
func (m MassFormatter) StringFromValueUnit(value float64, unit NSMassFormatterUnit) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return NSStringFromID(rv).String()
}

// Returns the unit string for the provided value.
//
// numberInKilograms: The mass’s value in kilograms.
//
// unitp: An output parameter. This will hold the [MassFormatter.Unit] value that
// corresponds to the returned units.
// //
// [MassFormatter.Unit]: https://developer.apple.com/documentation/Foundation/MassFormatter/Unit
//
// # Return Value
// 
// A localized string representing the unit.
//
// # Discussion
// 
// This method selects the correct unit based on the formatter’s locale, the
// magnitude of the value, and the [ForPersonMassUse] property. The value,
// once converted into the appropriate unit, determines whether the unit
// string is plural or singular.
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/unitString(fromKilograms:usedUnit:)
func (m MassFormatter) UnitStringFromKilogramsUsedUnit(numberInKilograms float64, unitp NSMassFormatterUnit) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("unitStringFromKilograms:usedUnit:"), numberInKilograms, unitp)
	return NSStringFromID(rv).String()
}

// Returns the unit string based on the provided value and unit.
//
// value: The mass’s value for the provided unit.
//
// unit: The unit to use in the resulting mass string.
//
// # Return Value
// 
// A localized string representing the given unit. The provided value
// determines whether the unit is plural or singular.
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/unitString(fromValue:unit:)
func (m MassFormatter) UnitStringFromValueUnit(value float64, unit NSMassFormatterUnit) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the resulting string represents a
// person’s mass.
//
// # Discussion
// 
// Returns [true] if the value passed to [StringFromKilograms] or
// [UnitStringFromKilogramsUsedUnit] is a person’s mass; otherwise, [false].
// By default, this property returns [false].
// 
// The mass formatter uses this property when determining the best unit for a
// given locale (for example, in the [StringFromKilograms] method).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/isForPersonMassUse
func (m MassFormatter) ForPersonMassUse() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isForPersonMassUse"))
	return rv
}
func (m MassFormatter) SetForPersonMassUse(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setForPersonMassUse:"), value)
}

// The number formatter used to format the numbers in a mass strings.
//
// # Discussion
// 
// This property defaults to a number formatter using the
// [NumberFormatterDecimalStyle] style. You can provide a different number
// formatter to customize the mass string’s appearance.
//
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/numberFormatter
func (m MassFormatter) NumberFormatter() INSNumberFormatter {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("numberFormatter"))
	return NSNumberFormatterFromID(objc.ID(rv))
}
func (m MassFormatter) SetNumberFormatter(value INSNumberFormatter) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumberFormatter:"), value)
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
// See: https://developer.apple.com/documentation/Foundation/MassFormatter/unitStyle
func (m MassFormatter) UnitStyle() NSFormattingUnitStyle {
	rv := objc.Send[NSFormattingUnitStyle](m.ID, objc.Sel("unitStyle"))
	return NSFormattingUnitStyle(rv)
}
func (m MassFormatter) SetUnitStyle(value NSFormattingUnitStyle) {
	objc.Send[struct{}](m.ID, objc.Sel("setUnitStyle:"), value)
}

