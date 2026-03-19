// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitFrequency] class.
var (
	_UnitFrequencyClass     UnitFrequencyClass
	_UnitFrequencyClassOnce sync.Once
)

func getUnitFrequencyClass() UnitFrequencyClass {
	_UnitFrequencyClassOnce.Do(func() {
		_UnitFrequencyClass = UnitFrequencyClass{class: objc.GetClass("NSUnitFrequency")}
	})
	return _UnitFrequencyClass
}

// GetUnitFrequencyClass returns the class object for NSUnitFrequency.
func GetUnitFrequencyClass() UnitFrequencyClass {
	return getUnitFrequencyClass()
}

type UnitFrequencyClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitFrequencyClass) Alloc() UnitFrequency {
	rv := objc.Send[UnitFrequency](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for frequency.
//
// # Overview
// 
// You typically use instances of [NSUnitFrequency] to represent specific
// quantities of frequency using the [NSMeasurement] class.
// 
// # Frequency
// 
// Frequency is a quantity of occurrences for a repeating event over time. The
// SI unit for frequency is the hertz (Hz), which is a derived as one
// occurrence per second (`1 Hz = 1 / 1s`).
// 
// The [NSUnitFrequency] class defines its [BaseUnit] as [Hertz], and provides
// the following units, which are initialized using [NSUnitConverterLinear]
// converters with the specified coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency
type UnitFrequency struct {
	NSDimension
}

// UnitFrequencyFromID constructs a [UnitFrequency] from an objc.ID.
//
// A unit of measure for frequency.
func UnitFrequencyFromID(id objc.ID) UnitFrequency {
	return NSUnitFrequency{NSDimension: NSDimensionFromID(id)}
}

// NSUnitFrequencyFromID is an alias for [UnitFrequencyFromID] for cross-framework compatibility.
func NSUnitFrequencyFromID(id objc.ID) UnitFrequency { return UnitFrequencyFromID(id) }
// NOTE: UnitFrequency adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitFrequency] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency
type IUnitFrequency interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitFrequency) Init() UnitFrequency {
	rv := objc.Send[UnitFrequency](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitFrequency) Autorelease() UnitFrequency {
	rv := objc.Send[UnitFrequency](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitFrequency creates a new UnitFrequency instance.
func NewUnitFrequency() UnitFrequency {
	class := getUnitFrequencyClass()
	rv := objc.Send[UnitFrequency](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitFrequencyWithCoder(coder INSCoder) UnitFrequency {
	instance := getUnitFrequencyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitFrequencyFromID(rv)
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
func NewUnitFrequencyWithSymbol(symbol string) UnitFrequency {
	instance := getUnitFrequencyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitFrequencyFromID(rv)
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
func NewUnitFrequencyWithSymbolConverter(symbol string, converter INSUnitConverter) UnitFrequency {
	instance := getUnitFrequencyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitFrequencyFromID(rv)
}

// The terahertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/terahertz
func (_UnitFrequencyClass UnitFrequencyClass) Terahertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("terahertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The gigahertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/gigahertz
func (_UnitFrequencyClass UnitFrequencyClass) Gigahertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("gigahertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The megahertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/megahertz
func (_UnitFrequencyClass UnitFrequencyClass) Megahertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("megahertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The kilohertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/kilohertz
func (_UnitFrequencyClass UnitFrequencyClass) Kilohertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("kilohertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The hertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/hertz
func (_UnitFrequencyClass UnitFrequencyClass) Hertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("hertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The millihertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/millihertz
func (_UnitFrequencyClass UnitFrequencyClass) Millihertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("millihertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The microhertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/microhertz
func (_UnitFrequencyClass UnitFrequencyClass) Microhertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("microhertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The nanohertz unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/nanohertz
func (_UnitFrequencyClass UnitFrequencyClass) Nanohertz() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("nanohertz"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}
// The frames per second unit of frequency.
//
// See: https://developer.apple.com/documentation/Foundation/UnitFrequency/framesPerSecond
func (_UnitFrequencyClass UnitFrequencyClass) FramesPerSecond() UnitFrequency {
	rv := objc.Send[objc.ID](objc.ID(_UnitFrequencyClass.class), objc.Sel("framesPerSecond"))
	return NSUnitFrequencyFromID(objc.ID(rv))
}

