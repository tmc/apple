// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitElectricCharge] class.
var (
	_UnitElectricChargeClass     UnitElectricChargeClass
	_UnitElectricChargeClassOnce sync.Once
)

func getUnitElectricChargeClass() UnitElectricChargeClass {
	_UnitElectricChargeClassOnce.Do(func() {
		_UnitElectricChargeClass = UnitElectricChargeClass{class: objc.GetClass("NSUnitElectricCharge")}
	})
	return _UnitElectricChargeClass
}

// GetUnitElectricChargeClass returns the class object for NSUnitElectricCharge.
func GetUnitElectricChargeClass() UnitElectricChargeClass {
	return getUnitElectricChargeClass()
}

type UnitElectricChargeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitElectricChargeClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitElectricChargeClass) Alloc() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for electric charge.
//
// # Overview
//
// You typically use instances of [NSUnitElectricCharge] to represent specific
// quantities of electric charge using the [NSMeasurement] class.
//
// # Electric Charge
//
// Electric charge is a fundamental physical property of matter that causes it
// to experience a force within an electromagnetic field. The SI unit for
// electric charge is the coulomb (C), which is defined as the amount of
// charge carried by a current of one ampere in one second (1C = 1A · 1s).
// Charge is also commonly expressed in terms of ampere hours (Ah).
//
// The [NSUnitElectricCharge] class defines its [BaseUnit] as [Coulombs], and
// provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge
type UnitElectricCharge struct {
	NSDimension
}

// UnitElectricChargeFromID constructs a [UnitElectricCharge] from an objc.ID.
//
// A unit of measure for electric charge.
func UnitElectricChargeFromID(id objc.ID) UnitElectricCharge {
	return NSUnitElectricCharge{NSDimension: NSDimensionFromID(id)}
}

// NSUnitElectricChargeFromID is an alias for [UnitElectricChargeFromID] for cross-framework compatibility.
func NSUnitElectricChargeFromID(id objc.ID) UnitElectricCharge { return UnitElectricChargeFromID(id) }

// NOTE: UnitElectricCharge adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitElectricCharge] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge
type IUnitElectricCharge interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitElectricCharge) Init() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitElectricCharge) Autorelease() UnitElectricCharge {
	rv := objc.Send[UnitElectricCharge](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricCharge creates a new UnitElectricCharge instance.
func NewUnitElectricCharge() UnitElectricCharge {
	class := getUnitElectricChargeClass()
	rv := objc.Send[UnitElectricCharge](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitElectricChargeWithCoder(coder INSCoder) UnitElectricCharge {
	instance := getUnitElectricChargeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitElectricChargeFromID(rv)
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
func NewUnitElectricChargeWithSymbol(symbol string) UnitElectricCharge {
	instance := getUnitElectricChargeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitElectricChargeFromID(rv)
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
func NewUnitElectricChargeWithSymbolConverter(symbol string, converter INSUnitConverter) UnitElectricCharge {
	instance := getUnitElectricChargeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitElectricChargeFromID(rv)
}

// The coulombs unit of electric charge.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/coulombs
func (_UnitElectricChargeClass UnitElectricChargeClass) Coulombs() UnitElectricCharge {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricChargeClass.class), objc.Sel("coulombs"))
	return NSUnitElectricChargeFromID(objc.ID(rv))
}

// The megaampere hours unit of electric charge.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/megaampereHours
func (_UnitElectricChargeClass UnitElectricChargeClass) MegaampereHours() UnitElectricCharge {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricChargeClass.class), objc.Sel("megaampereHours"))
	return NSUnitElectricChargeFromID(objc.ID(rv))
}

// The kiloampere hours unit of electric charge.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/kiloampereHours
func (_UnitElectricChargeClass UnitElectricChargeClass) KiloampereHours() UnitElectricCharge {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricChargeClass.class), objc.Sel("kiloampereHours"))
	return NSUnitElectricChargeFromID(objc.ID(rv))
}

// The ampere hours unit of electric charge.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/ampereHours
func (_UnitElectricChargeClass UnitElectricChargeClass) AmpereHours() UnitElectricCharge {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricChargeClass.class), objc.Sel("ampereHours"))
	return NSUnitElectricChargeFromID(objc.ID(rv))
}

// The milliampere hours unit of electric charge.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/milliampereHours
func (_UnitElectricChargeClass UnitElectricChargeClass) MilliampereHours() UnitElectricCharge {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricChargeClass.class), objc.Sel("milliampereHours"))
	return NSUnitElectricChargeFromID(objc.ID(rv))
}

// The microampere hours unit of electric charge.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCharge/microampereHours
func (_UnitElectricChargeClass UnitElectricChargeClass) MicroampereHours() UnitElectricCharge {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricChargeClass.class), objc.Sel("microampereHours"))
	return NSUnitElectricChargeFromID(objc.ID(rv))
}
