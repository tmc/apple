// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitMass] class.
var (
	_UnitMassClass     UnitMassClass
	_UnitMassClassOnce sync.Once
)

func getUnitMassClass() UnitMassClass {
	_UnitMassClassOnce.Do(func() {
		_UnitMassClass = UnitMassClass{class: objc.GetClass("NSUnitMass")}
	})
	return _UnitMassClass
}

// GetUnitMassClass returns the class object for NSUnitMass.
func GetUnitMassClass() UnitMassClass {
	return getUnitMassClass()
}

type UnitMassClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitMassClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitMassClass) Alloc() UnitMass {
	rv := objc.Send[UnitMass](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for mass.
//
// # Overview
// 
// You typically use instances of [NSUnitMass] to represent specific
// quantities of mass using the [NSMeasurement] class.
// 
// # Mass
// 
// Mass is a fundamental property of matter that causes it to resist a force
// accelerating it. The SI unit for mass is the kilogram (kg), which defined
// in terms of the mass of the international prototype kilogram.
// 
// The [NSUnitMass] class defines its [BaseUnit] as [Kilograms], and provides
// the following units, which [NSUnitConverterLinear] converters initialize
// with the given coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass
type UnitMass struct {
	NSDimension
}

// UnitMassFromID constructs a [UnitMass] from an objc.ID.
//
// A unit of measure for mass.
func UnitMassFromID(id objc.ID) UnitMass {
	return NSUnitMass{NSDimension: NSDimensionFromID(id)}
}

// NSUnitMassFromID is an alias for [UnitMassFromID] for cross-framework compatibility.
func NSUnitMassFromID(id objc.ID) UnitMass { return UnitMassFromID(id) }
// NOTE: UnitMass adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitMass] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass
type IUnitMass interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitMass) Init() UnitMass {
	rv := objc.Send[UnitMass](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitMass) Autorelease() UnitMass {
	rv := objc.Send[UnitMass](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitMass creates a new UnitMass instance.
func NewUnitMass() UnitMass {
	class := getUnitMassClass()
	rv := objc.Send[UnitMass](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitMassWithCoder(coder INSCoder) UnitMass {
	instance := getUnitMassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitMassFromID(rv)
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
func NewUnitMassWithSymbol(symbol string) UnitMass {
	instance := getUnitMassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitMassFromID(rv)
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
func NewUnitMassWithSymbolConverter(symbol string, converter INSUnitConverter) UnitMass {
	instance := getUnitMassClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitMassFromID(rv)
}

// The kilograms unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/kilograms
func (_UnitMassClass UnitMassClass) Kilograms() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("kilograms"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The grams unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/grams
func (_UnitMassClass UnitMassClass) Grams() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("grams"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The decigrams unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/decigrams
func (_UnitMassClass UnitMassClass) Decigrams() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("decigrams"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The centigrams unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/centigrams
func (_UnitMassClass UnitMassClass) Centigrams() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("centigrams"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The milligrams unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/milligrams
func (_UnitMassClass UnitMassClass) Milligrams() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("milligrams"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The micrograms unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/micrograms
func (_UnitMassClass UnitMassClass) Micrograms() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("micrograms"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The nanograms unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/nanograms
func (_UnitMassClass UnitMassClass) Nanograms() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("nanograms"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The picograms unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/picograms
func (_UnitMassClass UnitMassClass) Picograms() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("picograms"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The ounces unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/ounces
func (_UnitMassClass UnitMassClass) Ounces() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("ounces"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The pounds unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/pounds
func (_UnitMassClass UnitMassClass) PoundsMass() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("poundsMass"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The stone unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/stones
func (_UnitMassClass UnitMassClass) Stones() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("stones"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The metric tons unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/metricTons
func (_UnitMassClass UnitMassClass) MetricTons() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("metricTons"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The short tons unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/shortTons
func (_UnitMassClass UnitMassClass) ShortTons() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("shortTons"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The carats unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/carats
func (_UnitMassClass UnitMassClass) Carats() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("carats"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The ounces troy unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/ouncesTroy
func (_UnitMassClass UnitMassClass) OuncesTroy() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("ouncesTroy"))
	return NSUnitMassFromID(objc.ID(rv))
}
// The slugs unit of mass.
//
// See: https://developer.apple.com/documentation/Foundation/UnitMass/slugs
func (_UnitMassClass UnitMassClass) Slugs() UnitMass {
	rv := objc.Send[objc.ID](objc.ID(_UnitMassClass.class), objc.Sel("slugs"))
	return NSUnitMassFromID(objc.ID(rv))
}

