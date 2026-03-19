// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitPressure] class.
var (
	_UnitPressureClass     UnitPressureClass
	_UnitPressureClassOnce sync.Once
)

func getUnitPressureClass() UnitPressureClass {
	_UnitPressureClassOnce.Do(func() {
		_UnitPressureClass = UnitPressureClass{class: objc.GetClass("NSUnitPressure")}
	})
	return _UnitPressureClass
}

// GetUnitPressureClass returns the class object for NSUnitPressure.
func GetUnitPressureClass() UnitPressureClass {
	return getUnitPressureClass()
}

type UnitPressureClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitPressureClass) Alloc() UnitPressure {
	rv := objc.Send[UnitPressure](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for pressure.
//
// # Overview
// 
// You typically use instances of [NSUnitPressure] to represent specific
// quantities of pressure using the [NSMeasurement] class.
// 
// # Pressure
// 
// Pressure is the normal force over a surface. The SI unit for pressure is
// the pascal (Pa), which is derived as one newton of force over one square
// meter (`1 Pa = 1 N / 1 m`2).
// 
// The [NSUnitPressure] class defines its [BaseUnit] as
// [NewtonsPerMetersSquared] and provides the following units, which
// [NSUnitConverterLinear] converters initialize with the given coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure
type UnitPressure struct {
	NSDimension
}

// UnitPressureFromID constructs a [UnitPressure] from an objc.ID.
//
// A unit of measure for pressure.
func UnitPressureFromID(id objc.ID) UnitPressure {
	return NSUnitPressure{NSDimension: NSDimensionFromID(id)}
}

// NSUnitPressureFromID is an alias for [UnitPressureFromID] for cross-framework compatibility.
func NSUnitPressureFromID(id objc.ID) UnitPressure { return UnitPressureFromID(id) }
// NOTE: UnitPressure adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitPressure] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure
type IUnitPressure interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitPressure) Init() UnitPressure {
	rv := objc.Send[UnitPressure](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitPressure) Autorelease() UnitPressure {
	rv := objc.Send[UnitPressure](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitPressure creates a new UnitPressure instance.
func NewUnitPressure() UnitPressure {
	class := getUnitPressureClass()
	rv := objc.Send[UnitPressure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitPressureWithCoder(coder INSCoder) UnitPressure {
	instance := getUnitPressureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitPressureFromID(rv)
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
func NewUnitPressureWithSymbol(symbol string) UnitPressure {
	instance := getUnitPressureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitPressureFromID(rv)
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
func NewUnitPressureWithSymbolConverter(symbol string, converter INSUnitConverter) UnitPressure {
	instance := getUnitPressureClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitPressureFromID(rv)
}

// The gigapascals unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/gigapascals
func (_UnitPressureClass UnitPressureClass) Gigapascals() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("gigapascals"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The megapascals unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/megapascals
func (_UnitPressureClass UnitPressureClass) Megapascals() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("megapascals"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The kilopascals unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/kilopascals
func (_UnitPressureClass UnitPressureClass) Kilopascals() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("kilopascals"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The hectopascals unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/hectopascals
func (_UnitPressureClass UnitPressureClass) Hectopascals() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("hectopascals"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The inches of mercury unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/inchesOfMercury
func (_UnitPressureClass UnitPressureClass) InchesOfMercury() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("inchesOfMercury"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The bars unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/bars
func (_UnitPressureClass UnitPressureClass) Bars() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("bars"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The millibars unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/millibars
func (_UnitPressureClass UnitPressureClass) Millibars() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("millibars"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The millimeters of mercury unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/millimetersOfMercury
func (_UnitPressureClass UnitPressureClass) MillimetersOfMercury() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("millimetersOfMercury"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The newtons per square meter unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/newtonsPerMetersSquared
func (_UnitPressureClass UnitPressureClass) NewtonsPerMetersSquared() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("newtonsPerMetersSquared"))
	return NSUnitPressureFromID(objc.ID(rv))
}
// The pounds per square inch unit of pressure.
//
// See: https://developer.apple.com/documentation/Foundation/UnitPressure/poundsForcePerSquareInch
func (_UnitPressureClass UnitPressureClass) PoundsForcePerSquareInch() UnitPressure {
	rv := objc.Send[objc.ID](objc.ID(_UnitPressureClass.class), objc.Sel("poundsForcePerSquareInch"))
	return NSUnitPressureFromID(objc.ID(rv))
}

