// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitElectricResistance] class.
var (
	_UnitElectricResistanceClass     UnitElectricResistanceClass
	_UnitElectricResistanceClassOnce sync.Once
)

func getUnitElectricResistanceClass() UnitElectricResistanceClass {
	_UnitElectricResistanceClassOnce.Do(func() {
		_UnitElectricResistanceClass = UnitElectricResistanceClass{class: objc.GetClass("NSUnitElectricResistance")}
	})
	return _UnitElectricResistanceClass
}

// GetUnitElectricResistanceClass returns the class object for NSUnitElectricResistance.
func GetUnitElectricResistanceClass() UnitElectricResistanceClass {
	return getUnitElectricResistanceClass()
}

type UnitElectricResistanceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitElectricResistanceClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitElectricResistanceClass) Alloc() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for electric resistance.
//
// # Overview
//
// You typically use instances of [NSUnitElectricResistance] to represent
// specific quantities of electric resistance using the [NSMeasurement] class.
//
// # Electric Resistance
//
// Electric resistance is the difficulty of passing an electric current
// through a conductor. The SI unit for electric resistance is the ohm (Ω),
// which is derived as the electric resistance that produces one ampere of
// current between two points in conductor with one volt of electric potential
// difference (1Ω = 1V/1A).
//
// The [NSUnitElectricResistance] class defines its [BaseUnit] as [Ohms], and
// provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance
type UnitElectricResistance struct {
	NSDimension
}

// UnitElectricResistanceFromID constructs a [UnitElectricResistance] from an objc.ID.
//
// A unit of measure for electric resistance.
func UnitElectricResistanceFromID(id objc.ID) UnitElectricResistance {
	return NSUnitElectricResistance{NSDimension: NSDimensionFromID(id)}
}

// NSUnitElectricResistanceFromID is an alias for [UnitElectricResistanceFromID] for cross-framework compatibility.
func NSUnitElectricResistanceFromID(id objc.ID) UnitElectricResistance {
	return UnitElectricResistanceFromID(id)
}

// NOTE: UnitElectricResistance adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitElectricResistance] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance
type IUnitElectricResistance interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitElectricResistance) Init() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitElectricResistance) Autorelease() UnitElectricResistance {
	rv := objc.Send[UnitElectricResistance](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricResistance creates a new UnitElectricResistance instance.
func NewUnitElectricResistance() UnitElectricResistance {
	class := getUnitElectricResistanceClass()
	rv := objc.Send[UnitElectricResistance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitElectricResistanceWithCoder(coder INSCoder) UnitElectricResistance {
	instance := getUnitElectricResistanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitElectricResistanceFromID(rv)
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
func NewUnitElectricResistanceWithSymbol(symbol string) UnitElectricResistance {
	instance := getUnitElectricResistanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitElectricResistanceFromID(rv)
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
func NewUnitElectricResistanceWithSymbolConverter(symbol string, converter INSUnitConverter) UnitElectricResistance {
	instance := getUnitElectricResistanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitElectricResistanceFromID(rv)
}

// The megaohms unit of electric resistance.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/megaohms
func (_UnitElectricResistanceClass UnitElectricResistanceClass) Megaohms() UnitElectricResistance {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricResistanceClass.class), objc.Sel("megaohms"))
	return NSUnitElectricResistanceFromID(objc.ID(rv))
}

// The kiloohms unit of electric resistance.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/kiloohms
func (_UnitElectricResistanceClass UnitElectricResistanceClass) Kiloohms() UnitElectricResistance {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricResistanceClass.class), objc.Sel("kiloohms"))
	return NSUnitElectricResistanceFromID(objc.ID(rv))
}

// The ohms unit of electric resistance.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/ohms
func (_UnitElectricResistanceClass UnitElectricResistanceClass) Ohms() UnitElectricResistance {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricResistanceClass.class), objc.Sel("ohms"))
	return NSUnitElectricResistanceFromID(objc.ID(rv))
}

// The milliohms unit of electric resistance.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/milliohms
func (_UnitElectricResistanceClass UnitElectricResistanceClass) Milliohms() UnitElectricResistance {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricResistanceClass.class), objc.Sel("milliohms"))
	return NSUnitElectricResistanceFromID(objc.ID(rv))
}

// The microohms unit of electric resistance.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricResistance/microohms
func (_UnitElectricResistanceClass UnitElectricResistanceClass) Microohms() UnitElectricResistance {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricResistanceClass.class), objc.Sel("microohms"))
	return NSUnitElectricResistanceFromID(objc.ID(rv))
}
