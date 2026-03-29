// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitConcentrationMass] class.
var (
	_UnitConcentrationMassClass     UnitConcentrationMassClass
	_UnitConcentrationMassClassOnce sync.Once
)

func getUnitConcentrationMassClass() UnitConcentrationMassClass {
	_UnitConcentrationMassClassOnce.Do(func() {
		_UnitConcentrationMassClass = UnitConcentrationMassClass{class: objc.GetClass("NSUnitConcentrationMass")}
	})
	return _UnitConcentrationMassClass
}

// GetUnitConcentrationMassClass returns the class object for NSUnitConcentrationMass.
func GetUnitConcentrationMassClass() UnitConcentrationMassClass {
	return getUnitConcentrationMassClass()
}

type UnitConcentrationMassClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitConcentrationMassClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitConcentrationMassClass) Alloc() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for concentration of mass.
//
// # Overview
// 
// You typically use instances of [NSUnitConcentrationMass] to represent
// specific quantities of concentration using the [NSMeasurement] class.
// 
// # Concentration of Mass
// 
// Concentration is the abundance of a constituent within a volume.
// Concentration can be expressed by SI derived units in terms of kilograms
// per cubic meter (kg/m3).
// 
// The [NSUnitConcentrationMass] class defines its [BaseUnit] as
// [GramsPerLiter], and provides the following units, which are initialized
// using [NSUnitConverterLinear] converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass
type UnitConcentrationMass struct {
	NSDimension
}

// UnitConcentrationMassFromID constructs a [UnitConcentrationMass] from an objc.ID.
//
// A unit of measure for concentration of mass.
func UnitConcentrationMassFromID(id objc.ID) UnitConcentrationMass {
	return NSUnitConcentrationMass{NSDimension: NSDimensionFromID(id)}
}

// NSUnitConcentrationMassFromID is an alias for [UnitConcentrationMassFromID] for cross-framework compatibility.
func NSUnitConcentrationMassFromID(id objc.ID) UnitConcentrationMass { return UnitConcentrationMassFromID(id) }
// NOTE: UnitConcentrationMass adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitConcentrationMass] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass
type IUnitConcentrationMass interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitConcentrationMass) Init() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitConcentrationMass) Autorelease() UnitConcentrationMass {
	rv := objc.Send[UnitConcentrationMass](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitConcentrationMass creates a new UnitConcentrationMass instance.
func NewUnitConcentrationMass() UnitConcentrationMass {
	class := getUnitConcentrationMassClass()
	rv := objc.Send[UnitConcentrationMass](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitConcentrationMassWithCoder(coder INSCoder) UnitConcentrationMass {
	instance := getUnitConcentrationMassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitConcentrationMassFromID(rv)
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
func NewUnitConcentrationMassWithSymbol(symbol string) UnitConcentrationMass {
	instance := getUnitConcentrationMassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitConcentrationMassFromID(rv)
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
func NewUnitConcentrationMassWithSymbolConverter(symbol string, converter INSUnitConverter) UnitConcentrationMass {
	instance := getUnitConcentrationMassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitConcentrationMassFromID(rv)
}

// Returns the millimoles per liter unit with the specified number of grams
// per mole.
//
// gramsPerMole: The mass, in grams, for a mole of a given constituent.
//
// # Return Value
// 
// A unit expressing millimoles per liter with the specified molar mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass/millimolesPerLiter(withGramsPerMole:)
func (_UnitConcentrationMassClass UnitConcentrationMassClass) MillimolesPerLiterWithGramsPerMole(gramsPerMole float64) UnitConcentrationMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitConcentrationMassClass.class), objc.Sel("millimolesPerLiterWithGramsPerMole:"), gramsPerMole)
	return NSUnitConcentrationMassFromID(rv)
}

// The grams per liter unit of concentration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass/gramsPerLiter
func (_UnitConcentrationMassClass UnitConcentrationMassClass) GramsPerLiter() UnitConcentrationMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitConcentrationMassClass.class), objc.Sel("gramsPerLiter"))
	return NSUnitConcentrationMassFromID(objc.ID(rv))
}
// The milligrams per deciliter unit of concentration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitConcentrationMass/milligramsPerDeciliter
func (_UnitConcentrationMassClass UnitConcentrationMassClass) MilligramsPerDeciliter() UnitConcentrationMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitConcentrationMassClass.class), objc.Sel("milligramsPerDeciliter"))
	return NSUnitConcentrationMassFromID(objc.ID(rv))
}

