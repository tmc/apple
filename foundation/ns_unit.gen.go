// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Unit] class.
var (
	_UnitClass     UnitClass
	_UnitClassOnce sync.Once
)

func getUnitClass() UnitClass {
	_UnitClassOnce.Do(func() {
		_UnitClass = UnitClass{class: objc.GetClass("NSUnit")}
	})
	return _UnitClass
}

// GetUnitClass returns the class object for NSUnit.
func GetUnitClass() UnitClass {
	return getUnitClass()
}

type UnitClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitClass) Alloc() Unit {
	rv := objc.Send[Unit](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}







// An abstract class representing a unit of measure.
//
// # Overview
// 
// Each instance of an [NSUnit] subclass consists of a [Symbol], which can be
// used to create string representations of [NSMeasurement] objects with the
// [NSMeasurementFormatter] class.
// 
// The [NSDimension] subclass is an abstract class that represents a
// dimensional unit, which can be converted into different units of the same
// type. The Foundation framework provides several concrete [NSDimension]
// subclasses to represent the most common physical quantities, including
// mass, length, duration, and speed.
// 
// # Subclassing Notes
// 
// [NSUnit] is intended for subclassing. For dimensional units, you should use
// one of the Apple provided [NSDimension] subclasses listed in Table 1 of
// [NSDimension], or create a custom subclass of [NSDimension]. You can create
// a direct subclass of [NSUnit] to represent a custom dimensionless unit,
// such as a count, score, or ratio.
//
// # Accessing Properties
//
//   - [Unit.Symbol]: The symbolic representation of the unit.
//
// # Creating Units
//
//   - [Unit.InitWithSymbol]: Initializes a new unit with the specified symbol.
//
// See: https://developer.apple.com/documentation/Foundation/Unit
type Unit struct {
	objectivec.Object
}

// UnitFromID constructs a [Unit] from an objc.ID.
//
// An abstract class representing a unit of measure.
func UnitFromID(id objc.ID) Unit {
	return NSUnit{objectivec.Object{id}}
}

// NSUnitFromID is an alias for [UnitFromID] for cross-framework compatibility.
func NSUnitFromID(id objc.ID) Unit { return UnitFromID(id) }
// NOTE: Unit adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Unit] class.
//
// # Accessing Properties
//
//   - [IUnit.Symbol]: The symbolic representation of the unit.
//
// # Creating Units
//
//   - [IUnit.InitWithSymbol]: Initializes a new unit with the specified symbol.
//
// See: https://developer.apple.com/documentation/Foundation/Unit
type IUnit interface {
	objectivec.IObject
	NSCopying

	// Topic: Accessing Properties

	// The symbolic representation of the unit.
	Symbol() string

	// Topic: Creating Units

	// Initializes a new unit with the specified symbol.
	InitWithSymbol(symbol string) Unit

	InitWithCoder(coder INSCoder) Unit
	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
}





// Init initializes the instance.
func (u Unit) Init() Unit {
	rv := objc.Send[Unit](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u Unit) Autorelease() Unit {
	rv := objc.Send[Unit](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnit creates a new Unit instance.
func NewUnit() Unit {
	class := getUnitClass()
	rv := objc.Send[Unit](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitWithCoder(coder INSCoder) Unit {
	instance := getUnitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitFromID(rv)
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
func NewUnitWithSymbol(symbol string) Unit {
	instance := getUnitClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitFromID(rv)
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
func (u Unit) InitWithSymbol(symbol string) Unit {
	rv := objc.Send[Unit](u.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (u Unit) InitWithCoder(coder INSCoder) Unit {
	rv := objc.Send[Unit](u.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (u Unit) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The symbolic representation of the unit.
//
// # Discussion
// 
// The symbol of a unit is a string that can be used to designate a number as
// a quantity of a particular unit in user-readable representations. Units
// typically have symbols that are abbreviated and standardized, so as to be
// easily and unambiguously conveyed. For example, the `milePerHour` unit has
// the symbol `mph`. If a unit does not have a standardized or well-understood
// symbol, the lowercase name of the unit can be used. For example, the
// `metricCup` unit has the symbol `metric cup`.
// 
// Unit symbols may incorporate a metric prefix to indicate a multiple or
// fraction of existing unit symbols. For example, the `kilogram` unit has the
// symbol `kg`, which uses the SI prefix k for kilo- to indicate a magnitude
// of 103 for the `gram` unit, and the `microgram` unit has the symbol `µg`,
// which uses the SI prefix µ for micro- to indicate a magnitude of 10-6 for
// the `gram` unit.
//
// See: https://developer.apple.com/documentation/Foundation/Unit/symbol
func (u Unit) Symbol() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("symbol"))
	return NSStringFromID(rv).String()
}














			// Protocol methods for NSCopying
			

















