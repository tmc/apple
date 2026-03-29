// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitDuration] class.
var (
	_UnitDurationClass     UnitDurationClass
	_UnitDurationClassOnce sync.Once
)

func getUnitDurationClass() UnitDurationClass {
	_UnitDurationClassOnce.Do(func() {
		_UnitDurationClass = UnitDurationClass{class: objc.GetClass("NSUnitDuration")}
	})
	return _UnitDurationClass
}

// GetUnitDurationClass returns the class object for NSUnitDuration.
func GetUnitDurationClass() UnitDurationClass {
	return getUnitDurationClass()
}

type UnitDurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitDurationClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitDurationClass) Alloc() UnitDuration {
	rv := objc.Send[UnitDuration](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for a duration of time.
//
// # Overview
// 
// You typically use instances of [NSUnitDuration] to represent specific
// quantities of planar angle using the [NSMeasurement] class.
// 
// # Duration
// 
// Duration is a quantity of time. The SI unit for time is the second (sec),
// which is defined in terms of the radioactivity of a cesium-133 atom.
// Duration is also commonly expressed in terms of minutes (min) and hours
// (hr).
// 
// The [NSUnitDuration] class defines its [BaseUnit] as [Seconds], and
// provides the following units, which [NSUnitConverterLinear] converters
// initialize with the given coefficients:
// 
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration
type UnitDuration struct {
	NSDimension
}

// UnitDurationFromID constructs a [UnitDuration] from an objc.ID.
//
// A unit of measure for a duration of time.
func UnitDurationFromID(id objc.ID) UnitDuration {
	return NSUnitDuration{NSDimension: NSDimensionFromID(id)}
}

// NSUnitDurationFromID is an alias for [UnitDurationFromID] for cross-framework compatibility.
func NSUnitDurationFromID(id objc.ID) UnitDuration { return UnitDurationFromID(id) }
// NOTE: UnitDuration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitDuration] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration
type IUnitDuration interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitDuration) Init() UnitDuration {
	rv := objc.Send[UnitDuration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitDuration) Autorelease() UnitDuration {
	rv := objc.Send[UnitDuration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitDuration creates a new UnitDuration instance.
func NewUnitDuration() UnitDuration {
	class := getUnitDurationClass()
	rv := objc.Send[UnitDuration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitDurationWithCoder(coder INSCoder) UnitDuration {
	instance := getUnitDurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitDurationFromID(rv)
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
func NewUnitDurationWithSymbol(symbol string) UnitDuration {
	instance := getUnitDurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitDurationFromID(rv)
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
func NewUnitDurationWithSymbolConverter(symbol string, converter INSUnitConverter) UnitDuration {
	instance := getUnitDurationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitDurationFromID(rv)
}

// The hour unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/hours
func (_UnitDurationClass UnitDurationClass) Hours() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("hours"))
	return NSUnitDurationFromID(objc.ID(rv))
}
// The minute unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/minutes
func (_UnitDurationClass UnitDurationClass) Minutes() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("minutes"))
	return NSUnitDurationFromID(objc.ID(rv))
}
// The second unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/seconds
func (_UnitDurationClass UnitDurationClass) Seconds() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("seconds"))
	return NSUnitDurationFromID(objc.ID(rv))
}
// The millisecond unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/milliseconds
func (_UnitDurationClass UnitDurationClass) Milliseconds() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("milliseconds"))
	return NSUnitDurationFromID(objc.ID(rv))
}
// The microsecond unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/microseconds
func (_UnitDurationClass UnitDurationClass) Microseconds() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("microseconds"))
	return NSUnitDurationFromID(objc.ID(rv))
}
// The nanosecond unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/nanoseconds
func (_UnitDurationClass UnitDurationClass) Nanoseconds() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("nanoseconds"))
	return NSUnitDurationFromID(objc.ID(rv))
}
// The picosecond unit of duration.
//
// See: https://developer.apple.com/documentation/Foundation/UnitDuration/picoseconds
func (_UnitDurationClass UnitDurationClass) Picoseconds() UnitDuration {
	rv := objc.Send[objc.ID](objc.ID(_UnitDurationClass.class), objc.Sel("picoseconds"))
	return NSUnitDurationFromID(objc.ID(rv))
}

