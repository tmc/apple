// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitAcceleration] class.
var (
	_UnitAccelerationClass     UnitAccelerationClass
	_UnitAccelerationClassOnce sync.Once
)

func getUnitAccelerationClass() UnitAccelerationClass {
	_UnitAccelerationClassOnce.Do(func() {
		_UnitAccelerationClass = UnitAccelerationClass{class: objc.GetClass("NSUnitAcceleration")}
	})
	return _UnitAccelerationClass
}

// GetUnitAccelerationClass returns the class object for NSUnitAcceleration.
func GetUnitAccelerationClass() UnitAccelerationClass {
	return getUnitAccelerationClass()
}

type UnitAccelerationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitAccelerationClass) Alloc() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for acceleration.
//
// # Overview
// 
// You typically use instances of [NSUnitAcceleration] to represent specific
// quantities of acceleration using the [NSMeasurement] class.
// 
// # Acceleration
// 
// Acceleration is the rate of change of velocity. Acceleration can be
// expressed by SI derived units in terms of meters per second squared (m/s2).
// 
// The [NSUnitAcceleration] class defines its [BaseUnit] as
// [MetersPerSecondSquared], and provides the following units, which are
// initialized using [NSUnitConverterLinear] converters with the specified
// coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitAcceleration
type UnitAcceleration struct {
	NSDimension
}

// UnitAccelerationFromID constructs a [UnitAcceleration] from an objc.ID.
//
// A unit of measure for acceleration.
func UnitAccelerationFromID(id objc.ID) UnitAcceleration {
	return NSUnitAcceleration{NSDimension: NSDimensionFromID(id)}
}

// NSUnitAccelerationFromID is an alias for [UnitAccelerationFromID] for cross-framework compatibility.
func NSUnitAccelerationFromID(id objc.ID) UnitAcceleration { return UnitAccelerationFromID(id) }
// NOTE: UnitAcceleration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitAcceleration] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAcceleration
type IUnitAcceleration interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitAcceleration) Init() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitAcceleration) Autorelease() UnitAcceleration {
	rv := objc.Send[UnitAcceleration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitAcceleration creates a new UnitAcceleration instance.
func NewUnitAcceleration() UnitAcceleration {
	class := getUnitAccelerationClass()
	rv := objc.Send[UnitAcceleration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitAccelerationWithCoder(coder INSCoder) UnitAcceleration {
	instance := getUnitAccelerationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitAccelerationFromID(rv)
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
func NewUnitAccelerationWithSymbol(symbol string) UnitAcceleration {
	instance := getUnitAccelerationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitAccelerationFromID(rv)
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
func NewUnitAccelerationWithSymbolConverter(symbol string, converter INSUnitConverter) UnitAcceleration {
	instance := getUnitAccelerationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitAccelerationFromID(rv)
}

// Returns the meter per second squared unit of acceleration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAcceleration/metersPerSecondSquared
func (_UnitAccelerationClass UnitAccelerationClass) MetersPerSecondSquared() UnitAcceleration {
	rv := objc.Send[objc.ID](objc.ID(_UnitAccelerationClass.class), objc.Sel("metersPerSecondSquared"))
	return NSUnitAccelerationFromID(objc.ID(rv))
}

// Returns the gravity unit of acceleration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAcceleration/gravity
func (_UnitAccelerationClass UnitAccelerationClass) Gravity() UnitAcceleration {
	rv := objc.Send[objc.ID](objc.ID(_UnitAccelerationClass.class), objc.Sel("gravity"))
	return NSUnitAccelerationFromID(objc.ID(rv))
}

