// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitVolume] class.
var (
	_UnitVolumeClass     UnitVolumeClass
	_UnitVolumeClassOnce sync.Once
)

func getUnitVolumeClass() UnitVolumeClass {
	_UnitVolumeClassOnce.Do(func() {
		_UnitVolumeClass = UnitVolumeClass{class: objc.GetClass("NSUnitVolume")}
	})
	return _UnitVolumeClass
}

// GetUnitVolumeClass returns the class object for NSUnitVolume.
func GetUnitVolumeClass() UnitVolumeClass {
	return getUnitVolumeClass()
}

type UnitVolumeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitVolumeClass) Alloc() UnitVolume {
	rv := objc.Send[UnitVolume](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for volume.
//
// # Overview
// 
// You typically use instances of [NSUnitVolume] to represent specific
// quantities of volume using the [NSMeasurement] class.
// 
// # Volume
// 
// Volume is a quantity of the extend of matter in three dimensions. The SI
// accepted unit of volume is the liter (L), which is derived as one cubic
// decimeter (1 dm3). Volume is also commonly expressed in terms of cubic
// meters (m3), gallons (gal), and cups (cup).
// 
// The [NSUnitVolume] class defines its [BaseUnit] as [Liters], and provides
// the following units, which are initialized using [NSUnitConverterLinear]
// converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume
type UnitVolume struct {
	NSDimension
}

// UnitVolumeFromID constructs a [UnitVolume] from an objc.ID.
//
// A unit of measure for volume.
func UnitVolumeFromID(id objc.ID) UnitVolume {
	return NSUnitVolume{NSDimension: NSDimensionFromID(id)}
}

// NSUnitVolumeFromID is an alias for [UnitVolumeFromID] for cross-framework compatibility.
func NSUnitVolumeFromID(id objc.ID) UnitVolume { return UnitVolumeFromID(id) }
// NOTE: UnitVolume adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitVolume] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume
type IUnitVolume interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitVolume) Init() UnitVolume {
	rv := objc.Send[UnitVolume](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitVolume) Autorelease() UnitVolume {
	rv := objc.Send[UnitVolume](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitVolume creates a new UnitVolume instance.
func NewUnitVolume() UnitVolume {
	class := getUnitVolumeClass()
	rv := objc.Send[UnitVolume](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitVolumeWithCoder(coder INSCoder) UnitVolume {
	instance := getUnitVolumeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitVolumeFromID(rv)
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
func NewUnitVolumeWithSymbol(symbol string) UnitVolume {
	instance := getUnitVolumeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitVolumeFromID(rv)
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
func NewUnitVolumeWithSymbolConverter(symbol string, converter INSUnitConverter) UnitVolume {
	instance := getUnitVolumeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitVolumeFromID(rv)
}

// The megaliters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/megaliters
func (_UnitVolumeClass UnitVolumeClass) Megaliters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("megaliters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The kiloliters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/kiloliters
func (_UnitVolumeClass UnitVolumeClass) Kiloliters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("kiloliters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The liters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/liters
func (_UnitVolumeClass UnitVolumeClass) Liters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("liters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The deciliters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/deciliters
func (_UnitVolumeClass UnitVolumeClass) Deciliters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("deciliters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The centiliters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/centiliters
func (_UnitVolumeClass UnitVolumeClass) Centiliters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("centiliters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The milliliters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/milliliters
func (_UnitVolumeClass UnitVolumeClass) Milliliters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("milliliters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic kilometers unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicKilometers
func (_UnitVolumeClass UnitVolumeClass) CubicKilometers() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicKilometers"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic meters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMeters
func (_UnitVolumeClass UnitVolumeClass) CubicMeters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicMeters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic decimeters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicDecimeters
func (_UnitVolumeClass UnitVolumeClass) CubicDecimeters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicDecimeters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic centimeters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicCentimeters
func (_UnitVolumeClass UnitVolumeClass) CubicCentimeters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicCentimeters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic millimeters unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMillimeters
func (_UnitVolumeClass UnitVolumeClass) CubicMillimeters() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicMillimeters"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic inches unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicInches
func (_UnitVolumeClass UnitVolumeClass) CubicInches() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicInches"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic feet unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicFeet
func (_UnitVolumeClass UnitVolumeClass) CubicFeet() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicFeet"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic yards unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicYards
func (_UnitVolumeClass UnitVolumeClass) CubicYards() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicYards"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cubic miles unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cubicMiles
func (_UnitVolumeClass UnitVolumeClass) CubicMiles() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cubicMiles"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The acre feet unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/acreFeet
func (_UnitVolumeClass UnitVolumeClass) AcreFeet() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("acreFeet"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The bushels unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/bushels
func (_UnitVolumeClass UnitVolumeClass) Bushels() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("bushels"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The teaspoons unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/teaspoons
func (_UnitVolumeClass UnitVolumeClass) Teaspoons() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("teaspoons"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The tablespoons unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/tablespoons
func (_UnitVolumeClass UnitVolumeClass) Tablespoons() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("tablespoons"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The fluid ounces unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/fluidOunces
func (_UnitVolumeClass UnitVolumeClass) FluidOunces() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("fluidOunces"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The cups unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/cups
func (_UnitVolumeClass UnitVolumeClass) Cups() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("cups"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The pints unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/pints
func (_UnitVolumeClass UnitVolumeClass) Pints() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("pints"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The quarts unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/quarts
func (_UnitVolumeClass UnitVolumeClass) Quarts() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("quarts"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The gallons unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/gallons
func (_UnitVolumeClass UnitVolumeClass) Gallons() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("gallons"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The imperial teaspoons unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialTeaspoons
func (_UnitVolumeClass UnitVolumeClass) ImperialTeaspoons() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("imperialTeaspoons"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The imperial tablespoons unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialTablespoons
func (_UnitVolumeClass UnitVolumeClass) ImperialTablespoons() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("imperialTablespoons"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The imperial fluid ounces unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialFluidOunces
func (_UnitVolumeClass UnitVolumeClass) ImperialFluidOunces() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("imperialFluidOunces"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The imperial pints unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialPints
func (_UnitVolumeClass UnitVolumeClass) ImperialPints() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("imperialPints"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The imperial quarts unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialQuarts
func (_UnitVolumeClass UnitVolumeClass) ImperialQuarts() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("imperialQuarts"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The imperial gallons unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/imperialGallons
func (_UnitVolumeClass UnitVolumeClass) ImperialGallons() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("imperialGallons"))
	return NSUnitVolumeFromID(objc.ID(rv))
}
// The metric cups unit of volume.
//
// See: https://developer.apple.com/documentation/Foundation/UnitVolume/metricCups
func (_UnitVolumeClass UnitVolumeClass) MetricCups() UnitVolume {
	rv := objc.Send[objc.ID](objc.ID(_UnitVolumeClass.class), objc.Sel("metricCups"))
	return NSUnitVolumeFromID(objc.ID(rv))
}

