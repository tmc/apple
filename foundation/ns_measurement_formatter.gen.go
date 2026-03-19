// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MeasurementFormatter] class.
var (
	_MeasurementFormatterClass     MeasurementFormatterClass
	_MeasurementFormatterClassOnce sync.Once
)

func getMeasurementFormatterClass() MeasurementFormatterClass {
	_MeasurementFormatterClassOnce.Do(func() {
		_MeasurementFormatterClass = MeasurementFormatterClass{class: objc.GetClass("NSMeasurementFormatter")}
	})
	return _MeasurementFormatterClass
}

// GetMeasurementFormatterClass returns the class object for NSMeasurementFormatter.
func GetMeasurementFormatterClass() MeasurementFormatterClass {
	return getMeasurementFormatterClass()
}

type MeasurementFormatterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MeasurementFormatterClass) Alloc() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A formatter that provides localized representations of units and
// measurements.
//
// # Overview
// 
// You use the [StringFromMeasurement] method to create a localized
// representation of an [NSMeasurement] object, and you use the
// [StringFromUnit] method to create a localized representation of an [NSUnit]
// object. The formatter takes into account the specified [Locale],
// [UnitStyle], and [UnitOptions] when producing string representations of
// units and measurements.
//
// # Specifying the Format
//
//   - [MeasurementFormatter.UnitOptions]: The options for how the unit is formatted.
//   - [MeasurementFormatter.SetUnitOptions]
//   - [MeasurementFormatter.UnitStyle]: The unit style.
//   - [MeasurementFormatter.SetUnitStyle]
//   - [MeasurementFormatter.Locale]: The locale of the formatter.
//   - [MeasurementFormatter.SetLocale]
//   - [MeasurementFormatter.NumberFormatter]: The number formatter used to format the quantity of a measurement.
//   - [MeasurementFormatter.SetNumberFormatter]
//
// # Converting Measurements
//
//   - [MeasurementFormatter.StringFromMeasurement]: Creates and returns a localized string representation of the provided measurement.
//   - [MeasurementFormatter.StringFromUnit]: Creates and returns a localized string representation of the provided unit of measure.
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter
type MeasurementFormatter struct {
	NSFormatter
}

// MeasurementFormatterFromID constructs a [MeasurementFormatter] from an objc.ID.
//
// A formatter that provides localized representations of units and
// measurements.
func MeasurementFormatterFromID(id objc.ID) MeasurementFormatter {
	return NSMeasurementFormatter{NSFormatter: NSFormatterFromID(id)}
}

// NSMeasurementFormatterFromID is an alias for [MeasurementFormatterFromID] for cross-framework compatibility.
func NSMeasurementFormatterFromID(id objc.ID) MeasurementFormatter { return MeasurementFormatterFromID(id) }
// NOTE: MeasurementFormatter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MeasurementFormatter] class.
//
// # Specifying the Format
//
//   - [IMeasurementFormatter.UnitOptions]: The options for how the unit is formatted.
//   - [IMeasurementFormatter.SetUnitOptions]
//   - [IMeasurementFormatter.UnitStyle]: The unit style.
//   - [IMeasurementFormatter.SetUnitStyle]
//   - [IMeasurementFormatter.Locale]: The locale of the formatter.
//   - [IMeasurementFormatter.SetLocale]
//   - [IMeasurementFormatter.NumberFormatter]: The number formatter used to format the quantity of a measurement.
//   - [IMeasurementFormatter.SetNumberFormatter]
//
// # Converting Measurements
//
//   - [IMeasurementFormatter.StringFromMeasurement]: Creates and returns a localized string representation of the provided measurement.
//   - [IMeasurementFormatter.StringFromUnit]: Creates and returns a localized string representation of the provided unit of measure.
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter
type IMeasurementFormatter interface {
	INSFormatter
	NSSecureCoding

	// Topic: Specifying the Format

	// The options for how the unit is formatted.
	UnitOptions() NSMeasurementFormatterUnitOptions
	SetUnitOptions(value NSMeasurementFormatterUnitOptions)
	// The unit style.
	UnitStyle() NSFormattingUnitStyle
	SetUnitStyle(value NSFormattingUnitStyle)
	// The locale of the formatter.
	Locale() INSLocale
	SetLocale(value INSLocale)
	// The number formatter used to format the quantity of a measurement.
	NumberFormatter() INSNumberFormatter
	SetNumberFormatter(value INSNumberFormatter)

	// Topic: Converting Measurements

	// Creates and returns a localized string representation of the provided measurement.
	StringFromMeasurement(measurement INSMeasurement) string
	// Creates and returns a localized string representation of the provided unit of measure.
	StringFromUnit(unit INSUnit) string
}

// Init initializes the instance.
func (m MeasurementFormatter) Init() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MeasurementFormatter) Autorelease() MeasurementFormatter {
	rv := objc.Send[MeasurementFormatter](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMeasurementFormatter creates a new MeasurementFormatter instance.
func NewMeasurementFormatter() MeasurementFormatter {
	class := getMeasurementFormatterClass()
	rv := objc.Send[MeasurementFormatter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewMeasurementFormatterWithCoder(coder INSCoder) MeasurementFormatter {
	instance := getMeasurementFormatterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MeasurementFormatterFromID(rv)
}

// Creates and returns a localized string representation of the provided
// measurement.
//
// measurement: The measurement to be represented.
//
// # Return Value
// 
// A user-readable string that represents the measurement.
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/string(from:)-wt9y
func (m MeasurementFormatter) StringFromMeasurement(measurement INSMeasurement) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringFromMeasurement:"), measurement)
	return NSStringFromID(rv).String()
}
// Creates and returns a localized string representation of the provided unit
// of measure.
//
// unit: The unit of measure to be represented.
//
// # Return Value
// 
// A user-readable string that represents the unit of measure. If the unit
// cannot be localized, the unit’s [Symbol] value is used.
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/string(from:)-4hwjz
func (m MeasurementFormatter) StringFromUnit(unit INSUnit) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stringFromUnit:"), unit)
	return NSStringFromID(rv).String()
}

// The options for how the unit is formatted.
//
// # Discussion
// 
// You can set this property to ensure that the formatter chooses the
// preferred unit to format for the measurement based on the formatter’s
// locale. For possible values, see [MeasurementFormatter.UnitOptions].
// 
// If no options are specified, the formatter localizes according to the
// preferences of the formatter’s [Locale]. For example, a measurement in
// kilocalories may be formatted as [C] instead of `kcal`, or a measurement in
// kilometers per hour may be formatted as `miles per hour` for US and UK
// locales, but `kilometers per hour` for other locales. However, if the
// `providedUnit` option is specified, a measurement with [Kilocalories] units
// would be formatted as `kcal`, even if the locale prefers [C], and a
// measurement with [KilometersPerHour] units would be formatted as
// `kilometers per hour` for US and UK locales, even though they prefer `miles
// per hour`.
//
// [MeasurementFormatter.UnitOptions]: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/UnitOptions-swift.struct
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/unitOptions-swift.property
func (m MeasurementFormatter) UnitOptions() NSMeasurementFormatterUnitOptions {
	rv := objc.Send[NSMeasurementFormatterUnitOptions](m.ID, objc.Sel("unitOptions"))
	return NSMeasurementFormatterUnitOptions(rv)
}
func (m MeasurementFormatter) SetUnitOptions(value NSMeasurementFormatterUnitOptions) {
	objc.Send[struct{}](m.ID, objc.Sel("setUnitOptions:"), value)
}
// The unit style.
//
// # Discussion
// 
// The possible values are [FormattingUnitStyleShort],
// [FormattingUnitStyleMedium], and [FormattingUnitStyleLong]. The default
// value is [FormattingUnitStyleMedium].
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/unitStyle
func (m MeasurementFormatter) UnitStyle() NSFormattingUnitStyle {
	rv := objc.Send[NSFormattingUnitStyle](m.ID, objc.Sel("unitStyle"))
	return NSFormattingUnitStyle(rv)
}
func (m MeasurementFormatter) SetUnitStyle(value NSFormattingUnitStyle) {
	objc.Send[struct{}](m.ID, objc.Sel("setUnitStyle:"), value)
}
// The locale of the formatter.
//
// # Discussion
// 
// If unspecified, an [NSLocale] object representing the current system locale
// is used.
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/locale
func (m MeasurementFormatter) Locale() INSLocale {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("locale"))
	return NSLocaleFromID(objc.ID(rv))
}
func (m MeasurementFormatter) SetLocale(value INSLocale) {
	objc.Send[struct{}](m.ID, objc.Sel("setLocale:"), value)
}
// The number formatter used to format the quantity of a measurement.
//
// # Discussion
// 
// If unspecified, an [NSNumberFormatter] object with
// [NumberFormatterDecimalStyle] style is used.
//
// See: https://developer.apple.com/documentation/Foundation/MeasurementFormatter/numberFormatter
func (m MeasurementFormatter) NumberFormatter() INSNumberFormatter {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("numberFormatter"))
	return NSNumberFormatterFromID(objc.ID(rv))
}
func (m MeasurementFormatter) SetNumberFormatter(value INSNumberFormatter) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumberFormatter:"), value)
}

			// Protocol methods for NSSecureCoding
			

