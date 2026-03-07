// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitElectricPotentialDifference] class.
var (
	_UnitElectricPotentialDifferenceClass     UnitElectricPotentialDifferenceClass
	_UnitElectricPotentialDifferenceClassOnce sync.Once
)

func getUnitElectricPotentialDifferenceClass() UnitElectricPotentialDifferenceClass {
	_UnitElectricPotentialDifferenceClassOnce.Do(func() {
		_UnitElectricPotentialDifferenceClass = UnitElectricPotentialDifferenceClass{class: objc.GetClass("NSUnitElectricPotentialDifference")}
	})
	return _UnitElectricPotentialDifferenceClass
}

// GetUnitElectricPotentialDifferenceClass returns the class object for NSUnitElectricPotentialDifference.
func GetUnitElectricPotentialDifferenceClass() UnitElectricPotentialDifferenceClass {
	return getUnitElectricPotentialDifferenceClass()
}

type UnitElectricPotentialDifferenceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitElectricPotentialDifferenceClass) Alloc() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}







// A unit of measure for electric potential difference.
//
// # Overview
// 
// You typically use instances of [NSUnitElectricPotentialDifference] to
// represent specific quantities of electric potential difference using the
// [NSMeasurement] class.
// 
// # Electric Potential Difference
// 
// Electric potential difference is the amount of electric potential energy of
// a point charge at a point in space. The SI unit for electric potential
// difference is the volt (V), which is derived as the difference in electric
// potential energy between two points of a linear conductor when an electric
// current of one ampere dissipates one watt of power between those points (1V
// = 1W/1A).
// 
// The [NSUnitElectricPotentialDifference] class defines its [BaseUnit] as
// [Volts], and provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference
type UnitElectricPotentialDifference struct {
	NSDimension
}

// UnitElectricPotentialDifferenceFromID constructs a [UnitElectricPotentialDifference] from an objc.ID.
//
// A unit of measure for electric potential difference.
func UnitElectricPotentialDifferenceFromID(id objc.ID) UnitElectricPotentialDifference {
	return NSUnitElectricPotentialDifference{NSDimension: NSDimensionFromID(id)}
}

// NSUnitElectricPotentialDifferenceFromID is an alias for [UnitElectricPotentialDifferenceFromID] for cross-framework compatibility.
func NSUnitElectricPotentialDifferenceFromID(id objc.ID) UnitElectricPotentialDifference { return UnitElectricPotentialDifferenceFromID(id) }
// NOTE: UnitElectricPotentialDifference adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [UnitElectricPotentialDifference] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference
type IUnitElectricPotentialDifference interface {
	INSDimension
}





// Init initializes the instance.
func (u UnitElectricPotentialDifference) Init() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitElectricPotentialDifference) Autorelease() UnitElectricPotentialDifference {
	rv := objc.Send[UnitElectricPotentialDifference](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitElectricPotentialDifference creates a new UnitElectricPotentialDifference instance.
func NewUnitElectricPotentialDifference() UnitElectricPotentialDifference {
	class := getUnitElectricPotentialDifferenceClass()
	rv := objc.Send[UnitElectricPotentialDifference](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitElectricPotentialDifferenceWithCoder(coder INSCoder) UnitElectricPotentialDifference {
	instance := getUnitElectricPotentialDifferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitElectricPotentialDifferenceFromID(rv)
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
func NewUnitElectricPotentialDifferenceWithSymbol(symbol string) UnitElectricPotentialDifference {
	instance := getUnitElectricPotentialDifferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitElectricPotentialDifferenceFromID(rv)
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
func NewUnitElectricPotentialDifferenceWithSymbolConverter(symbol string, converter INSUnitConverter) UnitElectricPotentialDifference {
	instance := getUnitElectricPotentialDifferenceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitElectricPotentialDifferenceFromID(rv)
}






















// The megavolts unit of electric potential difference.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/megavolts
func (_UnitElectricPotentialDifferenceClass UnitElectricPotentialDifferenceClass) Megavolts() UnitElectricPotentialDifference {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricPotentialDifferenceClass.class), objc.Sel("megavolts"))
	return NSUnitElectricPotentialDifferenceFromID(objc.ID(rv))
}



// The kilovolts unit of electric potential difference.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/kilovolts
func (_UnitElectricPotentialDifferenceClass UnitElectricPotentialDifferenceClass) Kilovolts() UnitElectricPotentialDifference {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricPotentialDifferenceClass.class), objc.Sel("kilovolts"))
	return NSUnitElectricPotentialDifferenceFromID(objc.ID(rv))
}



// The volts unit of electric potential difference.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/volts
func (_UnitElectricPotentialDifferenceClass UnitElectricPotentialDifferenceClass) Volts() UnitElectricPotentialDifference {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricPotentialDifferenceClass.class), objc.Sel("volts"))
	return NSUnitElectricPotentialDifferenceFromID(objc.ID(rv))
}



// The millivolts unit of electric potential difference.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/millivolts
func (_UnitElectricPotentialDifferenceClass UnitElectricPotentialDifferenceClass) Millivolts() UnitElectricPotentialDifference {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricPotentialDifferenceClass.class), objc.Sel("millivolts"))
	return NSUnitElectricPotentialDifferenceFromID(objc.ID(rv))
}



// The microvolts unit of electric potential difference.
//
// See: https://developer.apple.com/documentation/Foundation/UnitElectricPotentialDifference/microvolts
func (_UnitElectricPotentialDifferenceClass UnitElectricPotentialDifferenceClass) Microvolts() UnitElectricPotentialDifference {
	rv := objc.Send[objc.ID](objc.ID(_UnitElectricPotentialDifferenceClass.class), objc.Sel("microvolts"))
	return NSUnitElectricPotentialDifferenceFromID(objc.ID(rv))
}
























