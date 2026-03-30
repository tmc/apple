// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitPower] class.
var (
	_UnitPowerClass     UnitPowerClass
	_UnitPowerClassOnce sync.Once
)

func getUnitPowerClass() UnitPowerClass {
	_UnitPowerClassOnce.Do(func() {
		_UnitPowerClass = UnitPowerClass{class: objc.GetClass("NSUnitPower")}
	})
	return _UnitPowerClass
}

// GetUnitPowerClass returns the class object for NSUnitPower.
func GetUnitPowerClass() UnitPowerClass {
	return getUnitPowerClass()
}

type UnitPowerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitPowerClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitPowerClass) Alloc() UnitPower {
	rv := objc.Send[UnitPower](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for power.
//
// # Overview
//
// You typically use instances of [NSUnitPower] to represent specific
// quantities of power using the [NSMeasurement] class.
//
// # Power
//
// Power is the amount of energy used over time. The SI unit for power is the
// watt (W), which is derived as one joule per second (1W = 1J / 1s).
//
// The [NSUnitPower] class defines its [BaseUnit] as [Watts], and provides the
// following units, which are initialized using [NSUnitConverterLinear]
// converters with the specified coefficients:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower
type UnitPower struct {
	NSDimension
}

// UnitPowerFromID constructs a [UnitPower] from an objc.ID.
//
// A unit of measure for power.
func UnitPowerFromID(id objc.ID) UnitPower {
	return NSUnitPower{NSDimension: NSDimensionFromID(id)}
}

// NSUnitPowerFromID is an alias for [UnitPowerFromID] for cross-framework compatibility.
func NSUnitPowerFromID(id objc.ID) UnitPower { return UnitPowerFromID(id) }

// NOTE: UnitPower adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitPower] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower
type IUnitPower interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitPower) Init() UnitPower {
	rv := objc.Send[UnitPower](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitPower) Autorelease() UnitPower {
	rv := objc.Send[UnitPower](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitPower creates a new UnitPower instance.
func NewUnitPower() UnitPower {
	class := getUnitPowerClass()
	rv := objc.Send[UnitPower](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitPowerWithCoder(coder INSCoder) UnitPower {
	instance := getUnitPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitPowerFromID(rv)
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
func NewUnitPowerWithSymbol(symbol string) UnitPower {
	instance := getUnitPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitPowerFromID(rv)
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
func NewUnitPowerWithSymbolConverter(symbol string, converter INSUnitConverter) UnitPower {
	instance := getUnitPowerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitPowerFromID(rv)
}

// The terawatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/terawatts
func (_UnitPowerClass UnitPowerClass) Terawatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("terawatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The gigawatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/gigawatts
func (_UnitPowerClass UnitPowerClass) Gigawatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("gigawatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The megawatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/megawatts
func (_UnitPowerClass UnitPowerClass) Megawatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("megawatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The kilowatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/kilowatts
func (_UnitPowerClass UnitPowerClass) Kilowatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("kilowatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The watts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/watts
func (_UnitPowerClass UnitPowerClass) Watts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("watts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The milliwatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/milliwatts
func (_UnitPowerClass UnitPowerClass) Milliwatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("milliwatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The microwatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/microwatts
func (_UnitPowerClass UnitPowerClass) Microwatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("microwatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The nanowatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/nanowatts
func (_UnitPowerClass UnitPowerClass) Nanowatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("nanowatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The picowatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/picowatts
func (_UnitPowerClass UnitPowerClass) Picowatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("picowatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The femtowatts unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/femtowatts
func (_UnitPowerClass UnitPowerClass) Femtowatts() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("femtowatts"))
	return NSUnitPowerFromID(objc.ID(rv))
}

// The horsepower unit of power.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPower/horsepower
func (_UnitPowerClass UnitPowerClass) Horsepower() UnitPower {
	rv := objc.Send[objc.ID](objc.ID(_UnitPowerClass.class), objc.Sel("horsepower"))
	return NSUnitPowerFromID(objc.ID(rv))
}
