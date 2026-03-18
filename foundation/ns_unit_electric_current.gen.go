// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitElectricCurrent] class.
var (
	_UnitElectricCurrentClass     UnitElectricCurrentClass
	_UnitElectricCurrentClassOnce sync.Once
)

func getUnitElectricCurrentClass() UnitElectricCurrentClass {
	_UnitElectricCurrentClassOnce.Do(func() {
		_UnitElectricCurrentClass = UnitElectricCurrentClass{class: objc.GetClass("NSUnitElectricCurrent")}
	})
	return _UnitElectricCurrentClass
}

// GetUnitElectricCurrentClass returns the class object for NSUnitElectricCurrent.
func GetUnitElectricCurrentClass() UnitElectricCurrentClass {
	return getUnitElectricCurrentClass()
}

type UnitElectricCurrentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitElectricCurrentClass) Alloc() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for electric current.
//
// # Overview
// 
// You typically use instances of [NSUnitElectricCurrent] to represent
// specific quantities of electric current using the [NSMeasurement] class.
// 
// # Electric Current
// 
// Electric current is the flow of electric charge. The SI unit for electric
// current is the ampere (A), which is defined in terms the production of
// electromagnetic force between two parallel linear conductors. It can also
// be expressed as the flow of one coulomb per second (1A = 1C / s).
// 
// The [NSUnitElectricCurrent] class defines its [BaseUnit] as [Amperes], and
// provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent
type UnitElectricCurrent struct {
	NSDimension
}

// UnitElectricCurrentFromID constructs a [UnitElectricCurrent] from an objc.ID.
//
// A unit of measure for electric current.
func UnitElectricCurrentFromID(id objc.ID) UnitElectricCurrent {
	return NSUnitElectricCurrent{NSDimension: NSDimensionFromID(id)}
}

// NSUnitElectricCurrentFromID is an alias for [UnitElectricCurrentFromID] for cross-framework compatibility.
func NSUnitElectricCurrentFromID(id objc.ID) UnitElectricCurrent { return UnitElectricCurrentFromID(id) }
// NOTE: UnitElectricCurrent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitElectricCurrent] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent
type IUnitElectricCurrent interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitElectricCurrent) Init() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitElectricCurrent) Autorelease() UnitElectricCurrent {
	rv := objc.Send[UnitElectricCurrent](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricCurrent creates a new UnitElectricCurrent instance.
func NewUnitElectricCurrent() UnitElectricCurrent {
	class := getUnitElectricCurrentClass()
	rv := objc.Send[UnitElectricCurrent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitElectricCurrentWithCoder(coder INSCoder) UnitElectricCurrent {
	instance := getUnitElectricCurrentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitElectricCurrentFromID(rv)
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
func NewUnitElectricCurrentWithSymbol(symbol string) UnitElectricCurrent {
	instance := getUnitElectricCurrentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitElectricCurrentFromID(rv)
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
func NewUnitElectricCurrentWithSymbolConverter(symbol string, converter INSUnitConverter) UnitElectricCurrent {
	instance := getUnitElectricCurrentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitElectricCurrentFromID(rv)
}

// The megaamperes unit of electric current.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/megaamperes
func (_UnitElectricCurrentClass UnitElectricCurrentClass) Megaamperes() UnitElectricCurrent {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricCurrentClass.class), objc.Sel("megaamperes"))
	return NSUnitElectricCurrentFromID(objc.ID(rv))
}

// The kiloamperes unit of electric current.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/kiloamperes
func (_UnitElectricCurrentClass UnitElectricCurrentClass) Kiloamperes() UnitElectricCurrent {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricCurrentClass.class), objc.Sel("kiloamperes"))
	return NSUnitElectricCurrentFromID(objc.ID(rv))
}

// The amperes unit of electric current.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/amperes
func (_UnitElectricCurrentClass UnitElectricCurrentClass) Amperes() UnitElectricCurrent {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricCurrentClass.class), objc.Sel("amperes"))
	return NSUnitElectricCurrentFromID(objc.ID(rv))
}

// The milliamperes unit of electric current.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/milliamperes
func (_UnitElectricCurrentClass UnitElectricCurrentClass) Milliamperes() UnitElectricCurrent {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricCurrentClass.class), objc.Sel("milliamperes"))
	return NSUnitElectricCurrentFromID(objc.ID(rv))
}

// The microamperes unit of electric current.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricCurrent/microamperes
func (_UnitElectricCurrentClass UnitElectricCurrentClass) Microamperes() UnitElectricCurrent {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricCurrentClass.class), objc.Sel("microamperes"))
	return NSUnitElectricCurrentFromID(objc.ID(rv))
}

