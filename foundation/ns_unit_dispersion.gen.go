// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitDispersion] class.
var (
	_UnitDispersionClass     UnitDispersionClass
	_UnitDispersionClassOnce sync.Once
)

func getUnitDispersionClass() UnitDispersionClass {
	_UnitDispersionClassOnce.Do(func() {
		_UnitDispersionClass = UnitDispersionClass{class: objc.GetClass("NSUnitDispersion")}
	})
	return _UnitDispersionClass
}

// GetUnitDispersionClass returns the class object for NSUnitDispersion.
func GetUnitDispersionClass() UnitDispersionClass {
	return getUnitDispersionClass()
}

type UnitDispersionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitDispersionClass) Alloc() UnitDispersion {
	rv := objc.Send[UnitDispersion](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for specific quantities of dispersion.
//
// # Overview
// 
// You typically use instances of [NSUnitDispersion] to represent specific
// quantities of dispersion using the [NSMeasurement] class.
// 
// # Dispersion
// 
// Dispersion describes the amount of a constituent divided by the amount of
// all other constituents in a mixture. Dispersion is a dimensionless quantity
// that is commonly expressed in “parts-per” notation, such as “parts
// per million” (ppm), to describe small relative quantities.
// 
// The [NSUnitDispersion] class defines its [BaseUnit] as [PartsPerMillion].
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitDispersion
type UnitDispersion struct {
	NSDimension
}

// UnitDispersionFromID constructs a [UnitDispersion] from an objc.ID.
//
// A unit of measure for specific quantities of dispersion.
func UnitDispersionFromID(id objc.ID) UnitDispersion {
	return NSUnitDispersion{NSDimension: NSDimensionFromID(id)}
}

// NSUnitDispersionFromID is an alias for [UnitDispersionFromID] for cross-framework compatibility.
func NSUnitDispersionFromID(id objc.ID) UnitDispersion { return UnitDispersionFromID(id) }
// NOTE: UnitDispersion adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitDispersion] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDispersion
type IUnitDispersion interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitDispersion) Init() UnitDispersion {
	rv := objc.Send[UnitDispersion](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitDispersion) Autorelease() UnitDispersion {
	rv := objc.Send[UnitDispersion](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitDispersion creates a new UnitDispersion instance.
func NewUnitDispersion() UnitDispersion {
	class := getUnitDispersionClass()
	rv := objc.Send[UnitDispersion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitDispersionWithCoder(coder INSCoder) UnitDispersion {
	instance := getUnitDispersionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitDispersionFromID(rv)
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
func NewUnitDispersionWithSymbol(symbol string) UnitDispersion {
	instance := getUnitDispersionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitDispersionFromID(rv)
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
func NewUnitDispersionWithSymbolConverter(symbol string, converter INSUnitConverter) UnitDispersion {
	instance := getUnitDispersionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitDispersionFromID(rv)
}

// The parts per million unit.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDispersion/partsPerMillion
func (_UnitDispersionClass UnitDispersionClass) PartsPerMillion() UnitDispersion {
	rv := objc.Send[objc.ID](objc.ID(_UnitDispersionClass.class), objc.Sel("partsPerMillion"))
	return NSUnitDispersionFromID(objc.ID(rv))
}

