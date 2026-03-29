// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitEnergy] class.
var (
	_UnitEnergyClass     UnitEnergyClass
	_UnitEnergyClassOnce sync.Once
)

func getUnitEnergyClass() UnitEnergyClass {
	_UnitEnergyClassOnce.Do(func() {
		_UnitEnergyClass = UnitEnergyClass{class: objc.GetClass("NSUnitEnergy")}
	})
	return _UnitEnergyClass
}

// GetUnitEnergyClass returns the class object for NSUnitEnergy.
func GetUnitEnergyClass() UnitEnergyClass {
	return getUnitEnergyClass()
}

type UnitEnergyClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitEnergyClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitEnergyClass) Alloc() UnitEnergy {
	rv := objc.Send[UnitEnergy](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for energy.
//
// # Overview
// 
// You typically use instances of [NSUnitEnergy] to represent specific
// quantities of energy using the [NSMeasurement] class.
// 
// # Energy
// 
// Energy is a fundamental property of matter than can be transferred and
// converted into different forms, such as kinetic, electric, and thermal. The
// SI unit for energy is the joule (J), which is derived as the work of one
// meter of displacement in the direction of a force of one newton (1J = 1N
// ∙ 1m). It can also be derived as the work required to displace an
// electric charge of one coulomb through an electrical potential difference
// of one volt (1J = 1C ∙ 1V), or the work required to produce one watt of
// power for one second (1J = 1W ∙ 1s). Energy is also commonly expressed in
// terms of the calorie (cal), or the energy needed to raise the temperature
// of one gram of water by one degree Celsius at a pressure of one atmosphere
// (1cal ≡ 4.184J).
// 
// The [NSUnitEnergy] class defines its [BaseUnit] as [Joules], and provides
// the following units, which are initialized using [NSUnitConverterLinear]
// converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy
type UnitEnergy struct {
	NSDimension
}

// UnitEnergyFromID constructs a [UnitEnergy] from an objc.ID.
//
// A unit of measure for energy.
func UnitEnergyFromID(id objc.ID) UnitEnergy {
	return NSUnitEnergy{NSDimension: NSDimensionFromID(id)}
}

// NSUnitEnergyFromID is an alias for [UnitEnergyFromID] for cross-framework compatibility.
func NSUnitEnergyFromID(id objc.ID) UnitEnergy { return UnitEnergyFromID(id) }
// NOTE: UnitEnergy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitEnergy] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy
type IUnitEnergy interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitEnergy) Init() UnitEnergy {
	rv := objc.Send[UnitEnergy](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitEnergy) Autorelease() UnitEnergy {
	rv := objc.Send[UnitEnergy](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitEnergy creates a new UnitEnergy instance.
func NewUnitEnergy() UnitEnergy {
	class := getUnitEnergyClass()
	rv := objc.Send[UnitEnergy](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitEnergyWithCoder(coder INSCoder) UnitEnergy {
	instance := getUnitEnergyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitEnergyFromID(rv)
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
func NewUnitEnergyWithSymbol(symbol string) UnitEnergy {
	instance := getUnitEnergyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitEnergyFromID(rv)
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
func NewUnitEnergyWithSymbolConverter(symbol string, converter INSUnitConverter) UnitEnergy {
	instance := getUnitEnergyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitEnergyFromID(rv)
}

// The kilojoules unit of energy.
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy/kilojoules
func (_UnitEnergyClass UnitEnergyClass) Kilojoules() UnitEnergy {
	rv := objc.Send[objc.ID](objc.ID(_UnitEnergyClass.class), objc.Sel("kilojoules"))
	return NSUnitEnergyFromID(objc.ID(rv))
}
// The joules unit of energy.
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy/joules
func (_UnitEnergyClass UnitEnergyClass) Joules() UnitEnergy {
	rv := objc.Send[objc.ID](objc.ID(_UnitEnergyClass.class), objc.Sel("joules"))
	return NSUnitEnergyFromID(objc.ID(rv))
}
// The kilocalories unit of energy.
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy/kilocalories
func (_UnitEnergyClass UnitEnergyClass) Kilocalories() UnitEnergy {
	rv := objc.Send[objc.ID](objc.ID(_UnitEnergyClass.class), objc.Sel("kilocalories"))
	return NSUnitEnergyFromID(objc.ID(rv))
}
// The calories unit of energy.
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy/calories
func (_UnitEnergyClass UnitEnergyClass) Calories() UnitEnergy {
	rv := objc.Send[objc.ID](objc.ID(_UnitEnergyClass.class), objc.Sel("calories"))
	return NSUnitEnergyFromID(objc.ID(rv))
}
// The kilowatt hours unit of energy.
//
// See: https://developer.apple.com/documentation/Foundation/UnitEnergy/kilowattHours
func (_UnitEnergyClass UnitEnergyClass) KilowattHours() UnitEnergy {
	rv := objc.Send[objc.ID](objc.ID(_UnitEnergyClass.class), objc.Sel("kilowattHours"))
	return NSUnitEnergyFromID(objc.ID(rv))
}

