// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitIlluminance] class.
var (
	_UnitIlluminanceClass     UnitIlluminanceClass
	_UnitIlluminanceClassOnce sync.Once
)

func getUnitIlluminanceClass() UnitIlluminanceClass {
	_UnitIlluminanceClassOnce.Do(func() {
		_UnitIlluminanceClass = UnitIlluminanceClass{class: objc.GetClass("NSUnitIlluminance")}
	})
	return _UnitIlluminanceClass
}

// GetUnitIlluminanceClass returns the class object for NSUnitIlluminance.
func GetUnitIlluminanceClass() UnitIlluminanceClass {
	return getUnitIlluminanceClass()
}

type UnitIlluminanceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitIlluminanceClass) Alloc() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for illuminance.
//
// # Overview
// 
// You typically use instances of [NSUnitIlluminance] to represent specific
// quantities of illuminance using the [NSMeasurement] class.
// 
// # Illuminance
// 
// Illuminance is the luminous flux incident on a surface. The SI unit for
// illuminance is the lux (lx), which is derived as one lumen per square meter
// (1lm / 1m2).
// 
// The [NSUnitIlluminance] class defines its [BaseUnit] as [Lux].
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitIlluminance
type UnitIlluminance struct {
	NSDimension
}

// UnitIlluminanceFromID constructs a [UnitIlluminance] from an objc.ID.
//
// A unit of measure for illuminance.
func UnitIlluminanceFromID(id objc.ID) UnitIlluminance {
	return NSUnitIlluminance{NSDimension: NSDimensionFromID(id)}
}

// NSUnitIlluminanceFromID is an alias for [UnitIlluminanceFromID] for cross-framework compatibility.
func NSUnitIlluminanceFromID(id objc.ID) UnitIlluminance { return UnitIlluminanceFromID(id) }
// NOTE: UnitIlluminance adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitIlluminance] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitIlluminance
type IUnitIlluminance interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitIlluminance) Init() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitIlluminance) Autorelease() UnitIlluminance {
	rv := objc.Send[UnitIlluminance](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitIlluminance creates a new UnitIlluminance instance.
func NewUnitIlluminance() UnitIlluminance {
	class := getUnitIlluminanceClass()
	rv := objc.Send[UnitIlluminance](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitIlluminanceWithCoder(coder INSCoder) UnitIlluminance {
	instance := getUnitIlluminanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitIlluminanceFromID(rv)
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
func NewUnitIlluminanceWithSymbol(symbol string) UnitIlluminance {
	instance := getUnitIlluminanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitIlluminanceFromID(rv)
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
func NewUnitIlluminanceWithSymbolConverter(symbol string, converter INSUnitConverter) UnitIlluminance {
	instance := getUnitIlluminanceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitIlluminanceFromID(rv)
}

