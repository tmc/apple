// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [EnergyFormatter] class.
var (
	_EnergyFormatterClass     EnergyFormatterClass
	_EnergyFormatterClassOnce sync.Once
)

func getEnergyFormatterClass() EnergyFormatterClass {
	_EnergyFormatterClassOnce.Do(func() {
		_EnergyFormatterClass = EnergyFormatterClass{class: objc.GetClass("NSEnergyFormatter")}
	})
	return _EnergyFormatterClass
}

// GetEnergyFormatterClass returns the class object for NSEnergyFormatter.
func GetEnergyFormatterClass() EnergyFormatterClass {
	return getEnergyFormatterClass()
}

type EnergyFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ec EnergyFormatterClass) Alloc() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](objc.ID(ec.class), objc.Sel("alloc"))
	return rv
}

// A formatter that provides localized descriptions of energy values.
//
// # Overview
//
// # Formatting Energy Strings
//
//   - [EnergyFormatter.ForFoodEnergyUse]: A Boolean value that indicates whether the energy value is used to measure food energy.
//   - [EnergyFormatter.SetForFoodEnergyUse]
//   - [EnergyFormatter.NumberFormatter]: The number formatter used to format the numbers in energy strings.
//   - [EnergyFormatter.SetNumberFormatter]
//   - [EnergyFormatter.StringFromJoules]: Returns an energy string for the provided value.
//   - [EnergyFormatter.StringFromValueUnit]: Returns a properly formatted energy string for the given value and unit.
//   - [EnergyFormatter.UnitStringFromJoulesUsedUnit]: Returns the unit string for the provided value.
//   - [EnergyFormatter.UnitStringFromValueUnit]: Returns the unit string based on the provided value and unit.
//   - [EnergyFormatter.UnitStyle]: The unit style used by this formatter.
//   - [EnergyFormatter.SetUnitStyle]
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter
type EnergyFormatter struct {
	NSFormatter
}

// EnergyFormatterFromID constructs a [EnergyFormatter] from an objc.ID.
//
// A formatter that provides localized descriptions of energy values.
func EnergyFormatterFromID(id objc.ID) EnergyFormatter {
	return NSEnergyFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSEnergyFormatterFromID is an alias for [EnergyFormatterFromID] for cross-framework compatibility.
func NSEnergyFormatterFromID(id objc.ID) EnergyFormatter { return EnergyFormatterFromID(id) }
// NOTE: EnergyFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [EnergyFormatter] class.
//
// # Formatting Energy Strings
//
//   - [IEnergyFormatter.ForFoodEnergyUse]: A Boolean value that indicates whether the energy value is used to measure food energy.
//   - [IEnergyFormatter.SetForFoodEnergyUse]
//   - [IEnergyFormatter.NumberFormatter]: The number formatter used to format the numbers in energy strings.
//   - [IEnergyFormatter.SetNumberFormatter]
//   - [IEnergyFormatter.StringFromJoules]: Returns an energy string for the provided value.
//   - [IEnergyFormatter.StringFromValueUnit]: Returns a properly formatted energy string for the given value and unit.
//   - [IEnergyFormatter.UnitStringFromJoulesUsedUnit]: Returns the unit string for the provided value.
//   - [IEnergyFormatter.UnitStringFromValueUnit]: Returns the unit string based on the provided value and unit.
//   - [IEnergyFormatter.UnitStyle]: The unit style used by this formatter.
//   - [IEnergyFormatter.SetUnitStyle]
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter
type IEnergyFormatter interface {
	INSFormatter

	// Topic: Formatting Energy Strings

	// A Boolean value that indicates whether the energy value is used to measure food energy.
	ForFoodEnergyUse() bool
	SetForFoodEnergyUse(value bool)
	// The number formatter used to format the numbers in energy strings.
	NumberFormatter() INSNumberFormatter
	SetNumberFormatter(value INSNumberFormatter)
	// Returns an energy string for the provided value.
	StringFromJoules(numberInJoules float64) string
	// Returns a properly formatted energy string for the given value and unit.
	StringFromValueUnit(value float64, unit NSEnergyFormatterUnit) string
	// Returns the unit string for the provided value.
	UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp NSEnergyFormatterUnit) string
	// Returns the unit string based on the provided value and unit.
	UnitStringFromValueUnit(value float64, unit NSEnergyFormatterUnit) string
	// The unit style used by this formatter.
	UnitStyle() NSFormattingUnitStyle
	SetUnitStyle(value NSFormattingUnitStyle)
}

// Init initializes the instance.
func (e EnergyFormatter) Init() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e EnergyFormatter) Autorelease() EnergyFormatter {
	rv := objc.Send[EnergyFormatter](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewEnergyFormatter creates a new EnergyFormatter instance.
func NewEnergyFormatter() EnergyFormatter {
	class := getEnergyFormatterClass()
	rv := objc.Send[EnergyFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewEnergyFormatterWithCoder(coder INSCoder) EnergyFormatter {
	instance := getEnergyFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return EnergyFormatterFromID(rv)
}

// Returns an energy string for the provided value.
//
// numberInJoules: The energy value in joules.
//
// # Return Value
// 
// A string that combines a value and a unit string appropriate for the
// formatter’s locale.
//
// # Discussion
// 
// This method converts the provided value in joules into units appropriate to
// the formatter’s locale.
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/string(fromJoules:)
func (e EnergyFormatter) StringFromJoules(numberInJoules float64) string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("stringFromJoules:"), numberInJoules)
	return NSStringFromID(rv).String()
}

// Returns a properly formatted energy string for the given value and unit.
//
// value: The energy value in the given unit.
//
// unit: The unit used in the resulting energy string.
//
// # Return Value
// 
// A localized string that combines the provided value and unit.
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/string(fromValue:unit:)
func (e EnergyFormatter) StringFromValueUnit(value float64, unit NSEnergyFormatterUnit) string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("stringFromValue:unit:"), value, unit)
	return NSStringFromID(rv).String()
}

// Returns the unit string for the provided value.
//
// numberInJoules: The energy value in joules.
//
// unitp: An output parameter. This will hold the [EnergyFormatter.Unit] value that
// corresponds to the returned units.
// //
// [EnergyFormatter.Unit]: https://developer.apple.com/documentation/Foundation/EnergyFormatter/Unit
//
// # Return Value
// 
// A localized string representing the unit.
//
// # Discussion
// 
// This method selects the correct unit based on the formatter’s locale, the
// magnitude of the value, and the [ForFoodEnergyUse] property.
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitString(fromJoules:usedUnit:)
func (e EnergyFormatter) UnitStringFromJoulesUsedUnit(numberInJoules float64, unitp NSEnergyFormatterUnit) string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("unitStringFromJoules:usedUnit:"), numberInJoules, unitp)
	return NSStringFromID(rv).String()
}

// Returns the unit string based on the provided value and unit.
//
// value: The energy value in the provided unit.
//
// unit: The unit to use in the resulting energy string.
//
// # Return Value
// 
// A localized string representing the given unit. The provided value
// determines whether the unit is plural or singular.
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitString(fromValue:unit:)
func (e EnergyFormatter) UnitStringFromValueUnit(value float64, unit NSEnergyFormatterUnit) string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("unitStringFromValue:unit:"), value, unit)
	return NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the energy value is used to measure
// food energy.
//
// # Discussion
// 
// Returns [true] if the energy is used to measure food energy; otherwise,
// [false]. If set to [true], [EnergyFormatterUnitKilocalorie] may be
// represented using “C” instead of “kcal”. By default, this property
// returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/isForFoodEnergyUse
func (e EnergyFormatter) ForFoodEnergyUse() bool {
	rv := objc.Send[bool](e.ID, objc.Sel("isForFoodEnergyUse"))
	return rv
}
func (e EnergyFormatter) SetForFoodEnergyUse(value bool) {
	objc.Send[struct{}](e.ID, objc.Sel("setForFoodEnergyUse:"), value)
}

// The number formatter used to format the numbers in energy strings.
//
// # Discussion
// 
// This property defaults to a number formatter using the
// [NumberFormatterDecimalStyle] style. You can provide a different number
// formatter to customize the energy string’s appearance.
//
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/numberFormatter
func (e EnergyFormatter) NumberFormatter() INSNumberFormatter {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("numberFormatter"))
	return NSNumberFormatterFromID(objc.ID(rv))
}
func (e EnergyFormatter) SetNumberFormatter(value INSNumberFormatter) {
	objc.Send[struct{}](e.ID, objc.Sel("setNumberFormatter:"), value)
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
// See: https://developer.apple.com/documentation/Foundation/EnergyFormatter/unitStyle
func (e EnergyFormatter) UnitStyle() NSFormattingUnitStyle {
	rv := objc.Send[NSFormattingUnitStyle](e.ID, objc.Sel("unitStyle"))
	return NSFormattingUnitStyle(rv)
}
func (e EnergyFormatter) SetUnitStyle(value NSFormattingUnitStyle) {
	objc.Send[struct{}](e.ID, objc.Sel("setUnitStyle:"), value)
}

