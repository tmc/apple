// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitConverterLinear] class.
var (
	_UnitConverterLinearClass     UnitConverterLinearClass
	_UnitConverterLinearClassOnce sync.Once
)

func getUnitConverterLinearClass() UnitConverterLinearClass {
	_UnitConverterLinearClassOnce.Do(func() {
		_UnitConverterLinearClass = UnitConverterLinearClass{class: objc.GetClass("NSUnitConverterLinear")}
	})
	return _UnitConverterLinearClass
}

// GetUnitConverterLinearClass returns the class object for NSUnitConverterLinear.
func GetUnitConverterLinearClass() UnitConverterLinearClass {
	return getUnitConverterLinearClass()
}

type UnitConverterLinearClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitConverterLinearClass) Alloc() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A description of how to convert between units using a linear equation.
//
// # Overview
// 
// A linear equation for unit conversion takes the form `y = mx + b`, such
// that the following is true:
// 
// - `y` is the value in terms of the base unit of the dimension. - `m` is the
// known coefficient to use for this unit’s conversion. - `x` is the value
// in terms of the unit on which you call this method. - `b` is the known
// constant to use for this unit’s conversion.
// 
// The `` method performs the conversion in the form of `y = mx + b`, where
// `x` represents the value passed in and `y` represents the value returned.
// The `` method performs the inverse conversion in the form of `x = (y - b) /
// m`, where `y` represents the value passed in and `x` represents the value
// returned.
// 
// For example, consider the [Fahrenheit] unit that [NSUnitTemperature]
// defines. The [BaseUnitValueFromValue] method calculates the value in the
// base unit, [Kelvin], using the formula `K = (0.55555555555556) × °F +
// 255.37222222222427`. The [ValueFromBaseUnitValue] method calculates the
// value in [Fahrenheit] using the formula `°F = (K — 255.37222222222427) /
// (0.55555555555556)`, where the [Coefficient] is `(0.55555555555556)` and
// the [Constant] is `255.37222222222427`.
// 
// Units that perform conversion using only a scale factor have a
// [Coefficient] equal to the scale factor and a [Constant] equal to `0`. For
// example, consider the [Kilometers] unit [NSUnitLength] defines. The
// [BaseUnitValueFromValue] method calculates the value in meters using the
// formula `valueInMeters = 1000 * valueInKilometers + 0`. The
// [ValueFromBaseUnitValue] calculates the value in kilometers using the
// formula `valueInKilometers = (valueInMeters - 0) / 1000`, where the
// coefficient is `1000` and the constant is `0`.
//
// # Accessing Linear Parameters
//
//   - [UnitConverterLinear.Coefficient]: The coefficient to use in the linear unit conversion calculation.
//   - [UnitConverterLinear.Constant]: The constant to use in the linear unit conversion calculation.
//
// # Creating Unit Converters
//
//   - [UnitConverterLinear.InitWithCoefficient]: Initializes the unit converter with the coefficient you specify.
//   - [UnitConverterLinear.InitWithCoefficientConstant]: Creates a unit converter with the coefficient and constant you specify.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear
type UnitConverterLinear struct {
	NSUnitConverter
}

// UnitConverterLinearFromID constructs a [UnitConverterLinear] from an objc.ID.
//
// A description of how to convert between units using a linear equation.
func UnitConverterLinearFromID(id objc.ID) UnitConverterLinear {
	return NSUnitConverterLinear{NSUnitConverter: NSUnitConverterFromID(id)}
}

// NSUnitConverterLinearFromID is an alias for [UnitConverterLinearFromID] for cross-framework compatibility.
func NSUnitConverterLinearFromID(id objc.ID) UnitConverterLinear { return UnitConverterLinearFromID(id) }
// NOTE: UnitConverterLinear adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitConverterLinear] class.
//
// # Accessing Linear Parameters
//
//   - [IUnitConverterLinear.Coefficient]: The coefficient to use in the linear unit conversion calculation.
//   - [IUnitConverterLinear.Constant]: The constant to use in the linear unit conversion calculation.
//
// # Creating Unit Converters
//
//   - [IUnitConverterLinear.InitWithCoefficient]: Initializes the unit converter with the coefficient you specify.
//   - [IUnitConverterLinear.InitWithCoefficientConstant]: Creates a unit converter with the coefficient and constant you specify.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear
type IUnitConverterLinear interface {
	INSUnitConverter
	NSCoding
	NSSecureCoding

	// Topic: Accessing Linear Parameters

	// The coefficient to use in the linear unit conversion calculation.
	Coefficient() float64
	// The constant to use in the linear unit conversion calculation.
	Constant() float64

	// Topic: Creating Unit Converters

	// Initializes the unit converter with the coefficient you specify.
	InitWithCoefficient(coefficient float64) UnitConverterLinear
	// Creates a unit converter with the coefficient and constant you specify.
	InitWithCoefficientConstant(coefficient float64, constant float64) UnitConverterLinear
}

// Init initializes the instance.
func (u UnitConverterLinear) Init() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitConverterLinear) Autorelease() UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConverterLinear creates a new UnitConverterLinear instance.
func NewUnitConverterLinear() UnitConverterLinear {
	class := getUnitConverterLinearClass()
	rv := objc.Send[UnitConverterLinear](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitConverterLinearWithCoder(coder INSCoder) UnitConverterLinear {
	instance := getUnitConverterLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitConverterLinearFromID(rv)
}

// Initializes the unit converter with the coefficient you specify.
//
// coefficient: The coefficient used in the linear unit conversion calculation.
//
// # Return Value
// 
// A unit converter initialized with the specified coefficient.
//
// # Discussion
// 
// Calling this initializer is equivalent to calling
// [InitWithCoefficientConstant], passing `0` for the `constant` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/init(coefficient:)
func NewUnitConverterLinearWithCoefficient(coefficient float64) UnitConverterLinear {
	instance := getUnitConverterLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoefficient:"), coefficient)
	return UnitConverterLinearFromID(rv)
}

// Creates a unit converter with the coefficient and constant you specify.
//
// coefficient: The coefficient used in the linear unit conversion calculation.
//
// constant: The constant used in the linear unit conversion calculation.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/init(coefficient:constant:)
func NewUnitConverterLinearWithCoefficientConstant(coefficient float64, constant float64) UnitConverterLinear {
	instance := getUnitConverterLinearClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoefficient:constant:"), coefficient, constant)
	return UnitConverterLinearFromID(rv)
}

// Initializes the unit converter with the coefficient you specify.
//
// coefficient: The coefficient used in the linear unit conversion calculation.
//
// # Return Value
// 
// A unit converter initialized with the specified coefficient.
//
// # Discussion
// 
// Calling this initializer is equivalent to calling
// [InitWithCoefficientConstant], passing `0` for the `constant` parameter.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/init(coefficient:)
func (u UnitConverterLinear) InitWithCoefficient(coefficient float64) UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](u.ID, objc.Sel("initWithCoefficient:"), coefficient)
	return rv
}

// Creates a unit converter with the coefficient and constant you specify.
//
// coefficient: The coefficient used in the linear unit conversion calculation.
//
// constant: The constant used in the linear unit conversion calculation.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/init(coefficient:constant:)
func (u UnitConverterLinear) InitWithCoefficientConstant(coefficient float64, constant float64) UnitConverterLinear {
	rv := objc.Send[UnitConverterLinear](u.ID, objc.Sel("initWithCoefficient:constant:"), coefficient, constant)
	return rv
}

// The coefficient to use in the linear unit conversion calculation.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/coefficient
func (u UnitConverterLinear) Coefficient() float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("coefficient"))
	return rv
}

// The constant to use in the linear unit conversion calculation.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverterLinear/constant
func (u UnitConverterLinear) Constant() float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("constant"))
	return rv
}

// The degree Fahrenheit unit of temperature.
//
// See: https://developer.apple.com/documentation/foundation/unittemperature/fahrenheit
func (_UnitConverterLinearClass UnitConverterLinearClass) Fahrenheit() UnitTemperature {
	rv := objc.Send[objc.ID](objc.ID(_UnitConverterLinearClass.class), objc.Sel("fahrenheit"))
	return NSUnitTemperatureFromID(objc.ID(rv))
}
func (_UnitConverterLinearClass UnitConverterLinearClass) SetFahrenheit(value UnitTemperature) {
	objc.Send[struct{}](objc.ID(_UnitConverterLinearClass.class), objc.Sel("setFahrenheit:"), value)
}

// The kelvin unit of temperature.
//
// See: https://developer.apple.com/documentation/foundation/unittemperature/kelvin
func (_UnitConverterLinearClass UnitConverterLinearClass) Kelvin() UnitTemperature {
	rv := objc.Send[objc.ID](objc.ID(_UnitConverterLinearClass.class), objc.Sel("kelvin"))
	return NSUnitTemperatureFromID(objc.ID(rv))
}
func (_UnitConverterLinearClass UnitConverterLinearClass) SetKelvin(value UnitTemperature) {
	objc.Send[struct{}](objc.ID(_UnitConverterLinearClass.class), objc.Sel("setKelvin:"), value)
}

// The kilometers unit of length.
//
// See: https://developer.apple.com/documentation/foundation/unitlength/kilometers
func (_UnitConverterLinearClass UnitConverterLinearClass) Kilometers() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitConverterLinearClass.class), objc.Sel("kilometers"))
	return NSUnitLengthFromID(objc.ID(rv))
}
func (_UnitConverterLinearClass UnitConverterLinearClass) SetKilometers(value UnitLength) {
	objc.Send[struct{}](objc.ID(_UnitConverterLinearClass.class), objc.Sel("setKilometers:"), value)
}

			// Protocol methods for NSSecureCoding
			

