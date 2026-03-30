// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitArea] class.
var (
	_UnitAreaClass     UnitAreaClass
	_UnitAreaClassOnce sync.Once
)

func getUnitAreaClass() UnitAreaClass {
	_UnitAreaClassOnce.Do(func() {
		_UnitAreaClass = UnitAreaClass{class: objc.GetClass("NSUnitArea")}
	})
	return _UnitAreaClass
}

// GetUnitAreaClass returns the class object for NSUnitArea.
func GetUnitAreaClass() UnitAreaClass {
	return getUnitAreaClass()
}

type UnitAreaClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitAreaClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitAreaClass) Alloc() UnitArea {
	rv := objc.Send[UnitArea](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for area.
//
// # Overview
//
// You typically use instances of [NSUnitArea] to represent specific
// quantities of area using the [NSMeasurement] class.
//
// # Area
//
// Area is a quantity of extent in two dimensions. Area can be expressed by SI
// derived units in terms of square meters (m2). Area is also commonly
// measured in square feet (ft2) and acres (ac).
//
// The [NSUnitArea] class defines its [BaseUnit] as [SquareMeters], and
// provides the following units, which are initialized using
// [NSUnitConverterLinear] converters with the specified coefficients:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea
type UnitArea struct {
	NSDimension
}

// UnitAreaFromID constructs a [UnitArea] from an objc.ID.
//
// A unit of measure for area.
func UnitAreaFromID(id objc.ID) UnitArea {
	return NSUnitArea{NSDimension: NSDimensionFromID(id)}
}

// NSUnitAreaFromID is an alias for [UnitAreaFromID] for cross-framework compatibility.
func NSUnitAreaFromID(id objc.ID) UnitArea { return UnitAreaFromID(id) }

// NOTE: UnitArea adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitArea] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea
type IUnitArea interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitArea) Init() UnitArea {
	rv := objc.Send[UnitArea](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitArea) Autorelease() UnitArea {
	rv := objc.Send[UnitArea](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitArea creates a new UnitArea instance.
func NewUnitArea() UnitArea {
	class := getUnitAreaClass()
	rv := objc.Send[UnitArea](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitAreaWithCoder(coder INSCoder) UnitArea {
	instance := getUnitAreaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitAreaFromID(rv)
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
func NewUnitAreaWithSymbol(symbol string) UnitArea {
	instance := getUnitAreaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitAreaFromID(rv)
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
func NewUnitAreaWithSymbolConverter(symbol string, converter INSUnitConverter) UnitArea {
	instance := getUnitAreaClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitAreaFromID(rv)
}

// The square megameters unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareMegameters
func (_UnitAreaClass UnitAreaClass) SquareMegameters() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareMegameters"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square kilometers unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareKilometers
func (_UnitAreaClass UnitAreaClass) SquareKilometers() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareKilometers"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square meters unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareMeters
func (_UnitAreaClass UnitAreaClass) SquareMeters() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareMeters"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square centimeters unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareCentimeters
func (_UnitAreaClass UnitAreaClass) SquareCentimeters() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareCentimeters"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square millimeters unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareMillimeters
func (_UnitAreaClass UnitAreaClass) SquareMillimeters() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareMillimeters"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square micrometers unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareMicrometers
func (_UnitAreaClass UnitAreaClass) SquareMicrometers() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareMicrometers"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square nanometers unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareNanometers
func (_UnitAreaClass UnitAreaClass) SquareNanometers() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareNanometers"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square inches unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareInches
func (_UnitAreaClass UnitAreaClass) SquareInches() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareInches"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square feet unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareFeet
func (_UnitAreaClass UnitAreaClass) SquareFeet() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareFeet"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square yards unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareYards
func (_UnitAreaClass UnitAreaClass) SquareYards() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareYards"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The square miles unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/squareMiles
func (_UnitAreaClass UnitAreaClass) SquareMiles() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("squareMiles"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The acres unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/acres
func (_UnitAreaClass UnitAreaClass) Acres() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("acres"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The ares unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/ares
func (_UnitAreaClass UnitAreaClass) Ares() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("ares"))
	return NSUnitAreaFromID(objc.ID(rv))
}

// The hectares unit of area.
//
// See: https://developer.apple.com/documentation/Foundation/UnitArea/hectares
func (_UnitAreaClass UnitAreaClass) Hectares() UnitArea {
	rv := objc.Send[objc.ID](objc.ID(_UnitAreaClass.class), objc.Sel("hectares"))
	return NSUnitAreaFromID(objc.ID(rv))
}
