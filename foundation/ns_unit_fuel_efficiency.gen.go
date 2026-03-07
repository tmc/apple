// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitFuelEfficiency] class.
var (
	_UnitFuelEfficiencyClass     UnitFuelEfficiencyClass
	_UnitFuelEfficiencyClassOnce sync.Once
)

func getUnitFuelEfficiencyClass() UnitFuelEfficiencyClass {
	_UnitFuelEfficiencyClassOnce.Do(func() {
		_UnitFuelEfficiencyClass = UnitFuelEfficiencyClass{class: objc.GetClass("NSUnitFuelEfficiency")}
	})
	return _UnitFuelEfficiencyClass
}

// GetUnitFuelEfficiencyClass returns the class object for NSUnitFuelEfficiency.
func GetUnitFuelEfficiencyClass() UnitFuelEfficiencyClass {
	return getUnitFuelEfficiencyClass()
}

type UnitFuelEfficiencyClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitFuelEfficiencyClass) Alloc() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}







// A unit of measure for fuel efficiency.
//
// # Overview
// 
// You typically use instances of [NSUnitFuelEfficiency] to represent specific
// quantities of fuel efficiency using the [NSMeasurement] class.
// 
// # Fuel Efficiency
// 
// Fuel efficiency corresponds to the thermal efficiency of a process that
// converts the chemical potential energy of a fuel into kinetic energy. Fuel
// efficiency can be expressed by SI derived units in terms of cubic meters
// per meter (m3/m), but is more commonly expressed in terms of liters per
// kilometer (L/km) and miles per gallon (mpg).
// 
// The [NSUnitFuelEfficiency] class defines its [BaseUnit] as
// [LitersPer100Kilometers], and provides the following units:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency
type UnitFuelEfficiency struct {
	NSDimension
}

// UnitFuelEfficiencyFromID constructs a [UnitFuelEfficiency] from an objc.ID.
//
// A unit of measure for fuel efficiency.
func UnitFuelEfficiencyFromID(id objc.ID) UnitFuelEfficiency {
	return NSUnitFuelEfficiency{NSDimension: NSDimensionFromID(id)}
}

// NSUnitFuelEfficiencyFromID is an alias for [UnitFuelEfficiencyFromID] for cross-framework compatibility.
func NSUnitFuelEfficiencyFromID(id objc.ID) UnitFuelEfficiency { return UnitFuelEfficiencyFromID(id) }
// NOTE: UnitFuelEfficiency adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [UnitFuelEfficiency] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency
type IUnitFuelEfficiency interface {
	INSDimension
}





// Init initializes the instance.
func (u UnitFuelEfficiency) Init() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitFuelEfficiency) Autorelease() UnitFuelEfficiency {
	rv := objc.Send[UnitFuelEfficiency](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitFuelEfficiency creates a new UnitFuelEfficiency instance.
func NewUnitFuelEfficiency() UnitFuelEfficiency {
	class := getUnitFuelEfficiencyClass()
	rv := objc.Send[UnitFuelEfficiency](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitFuelEfficiencyWithCoder(coder INSCoder) UnitFuelEfficiency {
	instance := getUnitFuelEfficiencyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitFuelEfficiencyFromID(rv)
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
func NewUnitFuelEfficiencyWithSymbol(symbol string) UnitFuelEfficiency {
	instance := getUnitFuelEfficiencyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitFuelEfficiencyFromID(rv)
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
func NewUnitFuelEfficiencyWithSymbolConverter(symbol string, converter INSUnitConverter) UnitFuelEfficiency {
	instance := getUnitFuelEfficiencyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitFuelEfficiencyFromID(rv)
}






















// The miles per imperial gallon unit of fuel efficiency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerImperialGallon
func (_UnitFuelEfficiencyClass UnitFuelEfficiencyClass) MilesPerImperialGallon() UnitFuelEfficiency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFuelEfficiencyClass.class), objc.Sel("milesPerImperialGallon"))
	return NSUnitFuelEfficiencyFromID(objc.ID(rv))
}



// The liters per 100 kilometers unit of fuel efficiency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/litersPer100Kilometers
func (_UnitFuelEfficiencyClass UnitFuelEfficiencyClass) LitersPer100Kilometers() UnitFuelEfficiency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFuelEfficiencyClass.class), objc.Sel("litersPer100Kilometers"))
	return NSUnitFuelEfficiencyFromID(objc.ID(rv))
}



// The miles per gallon unit of fuel efficiency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFuelEfficiency/milesPerGallon
func (_UnitFuelEfficiencyClass UnitFuelEfficiencyClass) MilesPerGallon() UnitFuelEfficiency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFuelEfficiencyClass.class), objc.Sel("milesPerGallon"))
	return NSUnitFuelEfficiencyFromID(objc.ID(rv))
}
























