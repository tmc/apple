// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitTemperature] class.
var (
	_UnitTemperatureClass     UnitTemperatureClass
	_UnitTemperatureClassOnce sync.Once
)

func getUnitTemperatureClass() UnitTemperatureClass {
	_UnitTemperatureClassOnce.Do(func() {
		_UnitTemperatureClass = UnitTemperatureClass{class: objc.GetClass("NSUnitTemperature")}
	})
	return _UnitTemperatureClass
}

// GetUnitTemperatureClass returns the class object for NSUnitTemperature.
func GetUnitTemperatureClass() UnitTemperatureClass {
	return getUnitTemperatureClass()
}

type UnitTemperatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitTemperatureClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitTemperatureClass) Alloc() UnitTemperature {
	rv := objc.Send[UnitTemperature](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for temperature.
//
// # Overview
// 
// You typically use instances of [NSUnitTemperature] to represent specific
// quantities of temperature using the [NSMeasurement] class.
// 
// # Temperature
// 
// Temperature is a comparative measure of thermal energy. The SI unit for
// temperature is the kelvin (K), which is defined in terms of the triple
// point of water. Temperature is also commonly measured by degrees of various
// scales, including Celsius (°C) and Fahrenheit (°F).
// 
// The [NSUnitTemperature] class defines its [BaseUnit] to be [Kelvin], and
// provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients and
// constants:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitTemperature
type UnitTemperature struct {
	NSDimension
}

// UnitTemperatureFromID constructs a [UnitTemperature] from an objc.ID.
//
// A unit of measure for temperature.
func UnitTemperatureFromID(id objc.ID) UnitTemperature {
	return NSUnitTemperature{NSDimension: NSDimensionFromID(id)}
}

// NSUnitTemperatureFromID is an alias for [UnitTemperatureFromID] for cross-framework compatibility.
func NSUnitTemperatureFromID(id objc.ID) UnitTemperature { return UnitTemperatureFromID(id) }
// NOTE: UnitTemperature adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitTemperature] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitTemperature
type IUnitTemperature interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitTemperature) Init() UnitTemperature {
	rv := objc.Send[UnitTemperature](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitTemperature) Autorelease() UnitTemperature {
	rv := objc.Send[UnitTemperature](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitTemperature creates a new UnitTemperature instance.
func NewUnitTemperature() UnitTemperature {
	class := getUnitTemperatureClass()
	rv := objc.Send[UnitTemperature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitTemperatureWithCoder(coder INSCoder) UnitTemperature {
	instance := getUnitTemperatureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitTemperatureFromID(rv)
}

// Initializes a new unit with the specified symbol.
//
// symbol: The symbol used to represent the unit.
//
// # Return Value
// 
// A new unit with the specified symbol.
//
// See: https://developer.apple.com/documentation/Foundation/Unit/init(symbol:)
func NewUnitTemperatureWithSymbol(symbol string) UnitTemperature {
	instance := getUnitTemperatureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitTemperatureFromID(rv)
}

// Initializes a dimensional unit with the symbol and unit converter you
// specify.
//
// symbol: The symbol used to represent the unit.
//
// converter: The unit converter used to represent the unit in terms of the dimension’s
// base unit.
//
// # Return Value
// 
// A new dimensional unit with the specified symbol and unit converter.
//
// # Discussion
// 
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/Dimension/init(symbol:converter:)
func NewUnitTemperatureWithSymbolConverter(symbol string, converter INSUnitConverter) UnitTemperature {
	instance := getUnitTemperatureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitTemperatureFromID(rv)
}

// The kelvin unit of temperature.
//
// See: https://developer.apple.com/documentation/Foundation/UnitTemperature/kelvin
func (_UnitTemperatureClass UnitTemperatureClass) Kelvin() UnitTemperature {
	rv := objc.Send[objc.ID](objc.ID(_UnitTemperatureClass.class), objc.Sel("kelvin"))
	return NSUnitTemperatureFromID(objc.ID(rv))
}
// The degree Celsius unit of temperature.
//
// See: https://developer.apple.com/documentation/Foundation/UnitTemperature/celsius
func (_UnitTemperatureClass UnitTemperatureClass) Celsius() UnitTemperature {
	rv := objc.Send[objc.ID](objc.ID(_UnitTemperatureClass.class), objc.Sel("celsius"))
	return NSUnitTemperatureFromID(objc.ID(rv))
}
// The degree Fahrenheit unit of temperature.
//
// See: https://developer.apple.com/documentation/Foundation/UnitTemperature/fahrenheit
func (_UnitTemperatureClass UnitTemperatureClass) Fahrenheit() UnitTemperature {
	rv := objc.Send[objc.ID](objc.ID(_UnitTemperatureClass.class), objc.Sel("fahrenheit"))
	return NSUnitTemperatureFromID(objc.ID(rv))
}

