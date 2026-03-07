// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitSpeed] class.
var (
	_UnitSpeedClass     UnitSpeedClass
	_UnitSpeedClassOnce sync.Once
)

func getUnitSpeedClass() UnitSpeedClass {
	_UnitSpeedClassOnce.Do(func() {
		_UnitSpeedClass = UnitSpeedClass{class: objc.GetClass("NSUnitSpeed")}
	})
	return _UnitSpeedClass
}

// GetUnitSpeedClass returns the class object for NSUnitSpeed.
func GetUnitSpeedClass() UnitSpeedClass {
	return getUnitSpeedClass()
}

type UnitSpeedClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitSpeedClass) Alloc() UnitSpeed {
	rv := objc.Send[UnitSpeed](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}







// A unit of measure for speed.
//
// # Overview
// 
// You typically use instances of [NSUnitSpeed] to represent specific
// quantities of speed using the [NSMeasurement] class.
// 
// # Speed
// 
// Speed is the magnitude of velocity, or the rate of change of position.
// Speed can be expressed by SI derived units in terms of meters per second
// (m/s), and is also commonly expressed in terms of kilometers per hour
// (km/h) and miles per hour (mph).
// 
// The [NSUnitSpeed] class defines its [BaseUnit] as [MetersPerSecond], and
// provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients:
// 
// [Table data omitted]
// 
// The base unit is [MetersPerSecond] and is accessed via [BaseUnit] on the
// [NSDimension] protocol.
//
// See: https://developer.apple.com/documentation/Foundation/UnitSpeed
type UnitSpeed struct {
	NSDimension
}

// UnitSpeedFromID constructs a [UnitSpeed] from an objc.ID.
//
// A unit of measure for speed.
func UnitSpeedFromID(id objc.ID) UnitSpeed {
	return NSUnitSpeed{NSDimension: NSDimensionFromID(id)}
}

// NSUnitSpeedFromID is an alias for [UnitSpeedFromID] for cross-framework compatibility.
func NSUnitSpeedFromID(id objc.ID) UnitSpeed { return UnitSpeedFromID(id) }
// NOTE: UnitSpeed adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [UnitSpeed] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitSpeed
type IUnitSpeed interface {
	INSDimension
}





// Init initializes the instance.
func (u UnitSpeed) Init() UnitSpeed {
	rv := objc.Send[UnitSpeed](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitSpeed) Autorelease() UnitSpeed {
	rv := objc.Send[UnitSpeed](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitSpeed creates a new UnitSpeed instance.
func NewUnitSpeed() UnitSpeed {
	class := getUnitSpeedClass()
	rv := objc.Send[UnitSpeed](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitSpeedWithCoder(coder INSCoder) UnitSpeed {
	instance := getUnitSpeedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitSpeedFromID(rv)
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
func NewUnitSpeedWithSymbol(symbol string) UnitSpeed {
	instance := getUnitSpeedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitSpeedFromID(rv)
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
func NewUnitSpeedWithSymbolConverter(symbol string, converter INSUnitConverter) UnitSpeed {
	instance := getUnitSpeedClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitSpeedFromID(rv)
}






















// The meter per second unit of speed.
//
// See: https://developer.apple.com/documentation/Foundation/UnitSpeed/metersPerSecond
func (_UnitSpeedClass UnitSpeedClass) MetersPerSecond() UnitSpeed {
	rv := objc.Send[objc.ID](objc.ID(_UnitSpeedClass.class), objc.Sel("metersPerSecond"))
	return NSUnitSpeedFromID(objc.ID(rv))
}



// The kilometers per hour unit of speed.
//
// See: https://developer.apple.com/documentation/Foundation/UnitSpeed/kilometersPerHour
func (_UnitSpeedClass UnitSpeedClass) KilometersPerHour() UnitSpeed {
	rv := objc.Send[objc.ID](objc.ID(_UnitSpeedClass.class), objc.Sel("kilometersPerHour"))
	return NSUnitSpeedFromID(objc.ID(rv))
}



// The miles per hour unit of speed.
//
// See: https://developer.apple.com/documentation/Foundation/UnitSpeed/milesPerHour
func (_UnitSpeedClass UnitSpeedClass) MilesPerHour() UnitSpeed {
	rv := objc.Send[objc.ID](objc.ID(_UnitSpeedClass.class), objc.Sel("milesPerHour"))
	return NSUnitSpeedFromID(objc.ID(rv))
}



// The knots unit of speed.
//
// See: https://developer.apple.com/documentation/Foundation/UnitSpeed/knots
func (_UnitSpeedClass UnitSpeedClass) Knots() UnitSpeed {
	rv := objc.Send[objc.ID](objc.ID(_UnitSpeedClass.class), objc.Sel("knots"))
	return NSUnitSpeedFromID(objc.ID(rv))
}
























