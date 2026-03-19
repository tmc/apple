// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitLength] class.
var (
	_UnitLengthClass     UnitLengthClass
	_UnitLengthClassOnce sync.Once
)

func getUnitLengthClass() UnitLengthClass {
	_UnitLengthClassOnce.Do(func() {
		_UnitLengthClass = UnitLengthClass{class: objc.GetClass("NSUnitLength")}
	})
	return _UnitLengthClass
}

// GetUnitLengthClass returns the class object for NSUnitLength.
func GetUnitLengthClass() UnitLengthClass {
	return getUnitLengthClass()
}

type UnitLengthClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitLengthClass) Alloc() UnitLength {
	rv := objc.Send[UnitLength](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for length.
//
// # Overview
// 
// You typically use instances of [NSUnitLength] to represent specific
// quantities of length using the [NSMeasurement] class.
// 
// # Length
// 
// Length is the dimensional extent of matter. The SI unit for length is the
// meter (m), which is defined in terms of the distance traveled by light in a
// vacuum.
// 
// The [NSUnitLength] class defines its [BaseUnit] as [Meters], and provides
// the following units, which are initialized using [NSUnitConverterLinear]
// converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength
type UnitLength struct {
	NSDimension
}

// UnitLengthFromID constructs a [UnitLength] from an objc.ID.
//
// A unit of measure for length.
func UnitLengthFromID(id objc.ID) UnitLength {
	return NSUnitLength{NSDimension: NSDimensionFromID(id)}
}

// NSUnitLengthFromID is an alias for [UnitLengthFromID] for cross-framework compatibility.
func NSUnitLengthFromID(id objc.ID) UnitLength { return UnitLengthFromID(id) }
// NOTE: UnitLength adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitLength] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength
type IUnitLength interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitLength) Init() UnitLength {
	rv := objc.Send[UnitLength](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitLength) Autorelease() UnitLength {
	rv := objc.Send[UnitLength](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitLength creates a new UnitLength instance.
func NewUnitLength() UnitLength {
	class := getUnitLengthClass()
	rv := objc.Send[UnitLength](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitLengthWithCoder(coder INSCoder) UnitLength {
	instance := getUnitLengthClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitLengthFromID(rv)
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
func NewUnitLengthWithSymbol(symbol string) UnitLength {
	instance := getUnitLengthClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitLengthFromID(rv)
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
func NewUnitLengthWithSymbolConverter(symbol string, converter INSUnitConverter) UnitLength {
	instance := getUnitLengthClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitLengthFromID(rv)
}

// The megameters unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/megameters
func (_UnitLengthClass UnitLengthClass) Megameters() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("megameters"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The kilometers unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/kilometers
func (_UnitLengthClass UnitLengthClass) Kilometers() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("kilometers"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The hectometers unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/hectometers
func (_UnitLengthClass UnitLengthClass) Hectometers() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("hectometers"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The decameters unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/decameters
func (_UnitLengthClass UnitLengthClass) Decameters() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("decameters"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The decimeters unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/decimeters
func (_UnitLengthClass UnitLengthClass) Decimeters() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("decimeters"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The centimeters unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/centimeters
func (_UnitLengthClass UnitLengthClass) Centimeters() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("centimeters"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The millimeters unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/millimeters
func (_UnitLengthClass UnitLengthClass) Millimeters() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("millimeters"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The micrometers unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/micrometers
func (_UnitLengthClass UnitLengthClass) Micrometers() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("micrometers"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The nanometers unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/nanometers
func (_UnitLengthClass UnitLengthClass) Nanometers() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("nanometers"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The picometers unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/picometers
func (_UnitLengthClass UnitLengthClass) Picometers() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("picometers"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The inches unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/inches
func (_UnitLengthClass UnitLengthClass) Inches() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("inches"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The feet unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/feet
func (_UnitLengthClass UnitLengthClass) Feet() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("feet"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The yards unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/yards
func (_UnitLengthClass UnitLengthClass) Yards() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("yards"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The Scandinavian miles unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/scandinavianMiles
func (_UnitLengthClass UnitLengthClass) ScandinavianMiles() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("scandinavianMiles"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The light years unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/lightyears
func (_UnitLengthClass UnitLengthClass) Lightyears() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("lightyears"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The nautical miles unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/nauticalMiles
func (_UnitLengthClass UnitLengthClass) NauticalMiles() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("nauticalMiles"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The fathoms unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/fathoms
func (_UnitLengthClass UnitLengthClass) Fathoms() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("fathoms"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The furlongs unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/furlongs
func (_UnitLengthClass UnitLengthClass) Furlongs() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("furlongs"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The astronomical units unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/astronomicalUnits
func (_UnitLengthClass UnitLengthClass) AstronomicalUnits() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("astronomicalUnits"))
	return NSUnitLengthFromID(objc.ID(rv))
}
// The parsecs unit of length.
//
// See: https://developer.apple.com/documentation/Foundation/UnitLength/parsecs
func (_UnitLengthClass UnitLengthClass) Parsecs() UnitLength {
	rv := objc.Send[objc.ID](objc.ID(_UnitLengthClass.class), objc.Sel("parsecs"))
	return NSUnitLengthFromID(objc.ID(rv))
}

