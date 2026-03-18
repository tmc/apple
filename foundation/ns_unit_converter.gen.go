// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UnitConverter] class.
var (
	_UnitConverterClass     UnitConverterClass
	_UnitConverterClassOnce sync.Once
)

func getUnitConverterClass() UnitConverterClass {
	_UnitConverterClassOnce.Do(func() {
		_UnitConverterClass = UnitConverterClass{class: objc.GetClass("NSUnitConverter")}
	})
	return _UnitConverterClass
}

// GetUnitConverterClass returns the class object for NSUnitConverter.
func GetUnitConverterClass() UnitConverterClass {
	return getUnitConverterClass()
}

type UnitConverterClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitConverterClass) Alloc() UnitConverter {
	rv := objc.Send[UnitConverter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that provides a description of how to convert a unit to
// and from the base unit of its dimension.
//
// # Overview
// 
// For units that can be converted by a scale factor or linear equation, use
// the concrete subclass [NSUnitConverterLinear].
// 
// # Subclassing Notes
// 
// [NSUnitConverter] is an abstract class that is intended for subclassing.
// You can implement your own subclass of [NSUnitConverter] to convert between
// units according to any desired mapping function. For example, units may be
// converted using a logarithmic, exponential, or quantile scale.
// 
// # Methods to Override
// 
// All subclasses must fully implement the following methods:
// 
// - [BaseUnitValueFromValue]
// - [ValueFromBaseUnitValue]
// 
// # Alternatives to Subclassing
// 
// As stated above, most physical units can be converted using a linear
// equation with [NSUnitConverterLinear]. You should only create a custom
// subclass of [NSUnitConverter] for units that cannot be converted in this
// way.
//
// # Converting Between Units
//
//   - [UnitConverter.BaseUnitValueFromValue]: For a given unit, returns the specified value of that unit in terms of the base unit of its dimension.
//   - [UnitConverter.ValueFromBaseUnitValue]: For a given unit, returns the specified value of the base unit in terms of that unit.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverter
type UnitConverter struct {
	objectivec.Object
}

// UnitConverterFromID constructs a [UnitConverter] from an objc.ID.
//
// An abstract class that provides a description of how to convert a unit to
// and from the base unit of its dimension.
func UnitConverterFromID(id objc.ID) UnitConverter {
	return NSUnitConverter{objectivec.Object{ID: id}}
}

// NSUnitConverterFromID is an alias for [UnitConverterFromID] for cross-framework compatibility.
func NSUnitConverterFromID(id objc.ID) UnitConverter { return UnitConverterFromID(id) }
// NOTE: UnitConverter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitConverter] class.
//
// # Converting Between Units
//
//   - [IUnitConverter.BaseUnitValueFromValue]: For a given unit, returns the specified value of that unit in terms of the base unit of its dimension.
//   - [IUnitConverter.ValueFromBaseUnitValue]: For a given unit, returns the specified value of the base unit in terms of that unit.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverter
type IUnitConverter interface {
	objectivec.IObject

	// Topic: Converting Between Units

	// For a given unit, returns the specified value of that unit in terms of the base unit of its dimension.
	BaseUnitValueFromValue(value float64) float64
	// For a given unit, returns the specified value of the base unit in terms of that unit.
	ValueFromBaseUnitValue(baseUnitValue float64) float64
}

// Init initializes the instance.
func (u UnitConverter) Init() UnitConverter {
	rv := objc.Send[UnitConverter](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitConverter) Autorelease() UnitConverter {
	rv := objc.Send[UnitConverter](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConverter creates a new UnitConverter instance.
func NewUnitConverter() UnitConverter {
	class := getUnitConverterClass()
	rv := objc.Send[UnitConverter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// For a given unit, returns the specified value of that unit in terms of the
// base unit of its dimension.
//
// value: The value in terms of a given unit.
//
// # Return Value
// 
// The value in terms of the base unit.
//
// # Discussion
// 
// This method takes a value in a particular unit and returns the result of
// converting it into the base unit of that unit’s dimension. For example, a
// converter for the miles unit calling this method, passing `1.0` to the
// `value` parameter, results in `1609.34` ().
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverter/baseUnitValue(fromValue:)
func (u UnitConverter) BaseUnitValueFromValue(value float64) float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("baseUnitValueFromValue:"), value)
	return rv
}

// For a given unit, returns the specified value of the base unit in terms of
// that unit.
//
// baseUnitValue: The value in terms of the base unit.
//
// # Return Value
// 
// The value in terms of a given unit.
//
// # Discussion
// 
// This method takes a value in the base unit of a unit’s dimension and
// returns the result of converting it into that unit. For example, a
// converter for the pounds unit calling this method, passing `2.20462` to the
// `baseUnitValue` parameter, results in `1.0` ().
//
// See: https://developer.apple.com/documentation/Foundation/UnitConverter/value(fromBaseUnitValue:)
func (u UnitConverter) ValueFromBaseUnitValue(baseUnitValue float64) float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("valueFromBaseUnitValue:"), baseUnitValue)
	return rv
}

