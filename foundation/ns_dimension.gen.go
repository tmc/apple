// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [Dimension] class.
var (
	_DimensionClass     DimensionClass
	_DimensionClassOnce sync.Once
)

func getDimensionClass() DimensionClass {
	_DimensionClassOnce.Do(func() {
		_DimensionClass = DimensionClass{class: objc.GetClass("NSDimension")}
	})
	return _DimensionClass
}

// GetDimensionClass returns the class object for NSDimension.
func GetDimensionClass() DimensionClass {
	return getDimensionClass()
}

type DimensionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DimensionClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DimensionClass) Alloc() Dimension {
	rv := objc.Send[Dimension](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class representing a dimensional unit of measure.
//
// # Overview
//
// The Foundation framework provides concrete subclasses for many of the most
// common types of physical units.
//
// Table 1: [NSDimension] subclasses.
//
// [Table data omitted]
//
// Each instance of a [NSDimension] subclass has a [Converter], which
// represents the unit in terms of the dimension’s [BaseUnit]. For example,
// the [NSLengthUnit] class uses [Meters] as its base unit. The system defines
// the predefined [Miles] unit by a [NSUnitConverterLinear] with a
// [Coefficient] of `1609.34`, which corresponds to the conversion ratio of
// miles to meters (1 mi = 1609.34 m); the system defines the predefined
// [Meters] unit by a [NSUnitConverterLinear] with a [Coefficient] of `1.0`
// because it’s the base unit.
//
// You typically use an [NSDimension] subclass in conjunction with the
// [NSMeasurement] class to represent specific quantities of a particular
// unit.
//
// # Working with Custom Units
//
// In addition to the Apple-provided units, you can define custom units. You
// can initialize custom units from a symbol and converter of an existing type
// or implemented as a class method of an existing type for additional
// convenience. You can also define your own [NSDimension] subclass to
// represent an entirely new unit dimension.
//
// # Initializing a Custom Unit with a Specified Symbol and Definition
//
// The simplest way to define a custom unit is to create a new instance of an
// existing [NSDimension] subclass using the [InitWithSymbolConverter] method.
//
// For example, the is a nonstandard unit of length (1 smoot = 1.70180 m). You
// can create a new instance of [NSUnitLength] as follows:
//
// # Extending Existing Dimension Subclasses
//
// Alternatively, if you use a custom unit extensively throughout an app,
// consider extending the corresponding [NSDimension] subclass and adding a
// static variable.
//
// For example, a measurement of speed can be furlongs per fortnight (1
// fur/ftn = 201.168 m / 1,209,600 s). If an app makes frequent use of this
// unit, you can extend [NSUnitSpeed] to add a `furlongsPerFortnight` static
// variable for convenient access as follows:
//
// # Creating a Custom Dimension Subclass
//
// You can create a new subclass of [NSDimension] to describe a new unit
// dimension.
//
// For example, the Foundation framework doesn’t define any units for
// radioactivity. Radioactivity is the process by which the nucleus of an atom
// emits radiation. The SI unit of measure for radioactivity is the becquerel
// (Bq), which is the quantity of radioactive material in which one nucleus
// decays per second (1 Bq = 1 s-1). Radioactivity is also commonly described
// in terms of curies (Ci), a unit defined relative to the decay of one gram
// of the radium-226 isotope (1 Ci = 3.7 × 1010 Bq). You can implement a
// [CustomUnitRadioactivity] class that defines both units of radioactivity as
// follows:
//
// # Subclassing Notes
//
// The system provides [NSDimension] for subclassing. Although the subclasses
// listed in Table 1 above are suitable for most purposes, you may want to
// define a custom unit type. For instance, you may need a custom unit type to
// represent a derived unit, such as magnetic flux (measured as the product of
// electric potential difference and time).
//
// To represent dimensionless units, subclass [NSUnit] directly.
//
// # Methods to Override
//
// All subclasses must fully implement the [BaseUnit] method designating the
// base unit, relative to which you define any additional units.
//
// You must also implement a class method named for the base unit itself, to
// use interchangeably. For example, the [NSUnitIlluminance] class defines its
// [BaseUnit] in terms of the lux (lx) and provides a corresponding [Lux]
// class method.
//
// # Alternatives to Subclassing
//
// As described in [NSDimension], you need to create a custom subclass of
// [NSDimension] only if you or the system haven’t defined a unit of the
// desired dimension. You can define a custom unit for an existing
// [NSDimension] subclass by either calling the [InitWithSymbolConverter]
// method or extending the subclass and adding a corresponding class method.
//
// # Creating Dimensions
//
//   - [Dimension.InitWithSymbolConverter]: Initializes a dimensional unit with the symbol and unit converter you specify.
//
// # Accessing the Unit Converter
//
//   - [Dimension.Converter]: The unit converter that represents the unit in terms of the dimension’s base unit.
//
// See: https://developer.apple.com/documentation/Foundation/Dimension
type Dimension struct {
	NSUnit
}

// DimensionFromID constructs a [Dimension] from an objc.ID.
//
// An abstract class representing a dimensional unit of measure.
func DimensionFromID(id objc.ID) Dimension {
	return NSDimension{NSUnit: NSUnitFromID(id)}
}

// NSDimensionFromID is an alias for [DimensionFromID] for cross-framework compatibility.
func NSDimensionFromID(id objc.ID) Dimension { return DimensionFromID(id) }

// NOTE: Dimension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [Dimension] class.
//
// # Creating Dimensions
//
//   - [IDimension.InitWithSymbolConverter]: Initializes a dimensional unit with the symbol and unit converter you specify.
//
// # Accessing the Unit Converter
//
//   - [IDimension.Converter]: The unit converter that represents the unit in terms of the dimension’s base unit.
//
// See: https://developer.apple.com/documentation/Foundation/Dimension
type IDimension interface {
	INSUnit

	// Topic: Creating Dimensions

	// Initializes a dimensional unit with the symbol and unit converter you specify.
	InitWithSymbolConverter(symbol string, converter INSUnitConverter) Dimension

	// Topic: Accessing the Unit Converter

	// The unit converter that represents the unit in terms of the dimension’s base unit.
	Converter() INSUnitConverter

	// The coefficient to use in the linear unit conversion calculation.
	Coefficient() float64
	SetCoefficient(value float64)
}

// Init initializes the instance.
func (d Dimension) Init() Dimension {
	rv := objc.Send[Dimension](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d Dimension) Autorelease() Dimension {
	rv := objc.Send[Dimension](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDimension creates a new Dimension instance.
func NewDimension() Dimension {
	class := getDimensionClass()
	rv := objc.Send[Dimension](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewDimensionWithCoder(coder INSCoder) Dimension {
	instance := getDimensionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return DimensionFromID(rv)
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
func NewDimensionWithSymbol(symbol string) Dimension {
	instance := getDimensionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return DimensionFromID(rv)
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
func NewDimensionWithSymbolConverter(symbol string, converter INSUnitConverter) Dimension {
	instance := getDimensionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return DimensionFromID(rv)
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
func (d Dimension) InitWithSymbolConverter(symbol string, converter INSUnitConverter) Dimension {
	rv := objc.Send[Dimension](d.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return rv
}

// Returns the base unit.
//
// # Return Value
//
// An [NSDimension] subclass object from which all other units provided by the
// subclass are defined.
//
// # Discussion
//
// The default implementation returns `nil` to indicate that the [NSDimension]
// class should not be used directly.
//
// When implementing a subclass, you should return a unit converter that
// returns the inputted value for both the “ and “ methods. You can create a
// unit converter for a base unit using the [NSUnitConverterLinear]
// [InitWithCoefficient] initializer, passing `1` as the coefficient.
//
// See: https://developer.apple.com/documentation/Foundation/Dimension/baseUnit()
func (_DimensionClass DimensionClass) BaseUnit() Dimension {
	rv := objc.Send[objc.ID](objc.ID(_DimensionClass.class), objc.Sel("baseUnit"))
	return NSDimensionFromID(rv)
}

// The unit converter that represents the unit in terms of the dimension’s
// base unit.
//
// See: https://developer.apple.com/documentation/Foundation/Dimension/converter
func (d Dimension) Converter() INSUnitConverter {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("converter"))
	return NSUnitConverterFromID(objc.ID(rv))
}

// The coefficient to use in the linear unit conversion calculation.
//
// See: https://developer.apple.com/documentation/foundation/unitconverterlinear/coefficient
func (d Dimension) Coefficient() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("coefficient"))
	return rv
}
func (d Dimension) SetCoefficient(value float64) {
	objc.Send[struct{}](d.ID, objc.Sel("setCoefficient:"), value)
}

// The lux unit of illuminance.
//
// See: https://developer.apple.com/documentation/foundation/unitilluminance/lux
func (_DimensionClass DimensionClass) Lux() UnitIlluminance {
	rv := objc.Send[objc.ID](objc.ID(_DimensionClass.class), objc.Sel("lux"))
	return NSUnitIlluminanceFromID(objc.ID(rv))
}
func (_DimensionClass DimensionClass) SetLux(value UnitIlluminance) {
	objc.Send[struct{}](objc.ID(_DimensionClass.class), objc.Sel("setLux:"), value)
}

// The meters unit of length.
//
// See: https://developer.apple.com/documentation/foundation/unitlength/meters
func (_DimensionClass DimensionClass) Meters() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_DimensionClass.class), objc.Sel("meters"))
	return NSUnitLengthFromID(objc.ID(rv))
}
func (_DimensionClass DimensionClass) SetMeters(value UnitLength) {
	objc.Send[struct{}](objc.ID(_DimensionClass.class), objc.Sel("setMeters:"), value)
}

// The miles unit of length.
//
// See: https://developer.apple.com/documentation/foundation/unitlength/miles
func (_DimensionClass DimensionClass) Miles() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_DimensionClass.class), objc.Sel("miles"))
	return NSUnitLengthFromID(objc.ID(rv))
}
func (_DimensionClass DimensionClass) SetMiles(value UnitLength) {
	objc.Send[struct{}](objc.ID(_DimensionClass.class), objc.Sel("setMiles:"), value)
}
