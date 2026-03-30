// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UnitAngle] class.
var (
	_UnitAngleClass     UnitAngleClass
	_UnitAngleClassOnce sync.Once
)

func getUnitAngleClass() UnitAngleClass {
	_UnitAngleClassOnce.Do(func() {
		_UnitAngleClass = UnitAngleClass{class: objc.GetClass("NSUnitAngle")}
	})
	return _UnitAngleClass
}

// GetUnitAngleClass returns the class object for NSUnitAngle.
func GetUnitAngleClass() UnitAngleClass {
	return getUnitAngleClass()
}

type UnitAngleClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UnitAngleClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UnitAngleClass) Alloc() UnitAngle {
	rv := objc.Send[UnitAngle](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A unit of measure for planar angle and rotation.
//
// # Overview
//
// You typically use instances of [NSUnitAngle] to represent specific
// quantities of planar angle using the [NSMeasurement] class.
//
// # Angle
//
// Angle is a quantity of rotation. The SI unit for angle is the radian (rad),
// which is dimensionless and defined to be the angle subtended by an arc that
// is equal in length to the radius of a circle. Angle is also commonly
// expressed in terms of degrees (°) and revolutions (rev).
//
// The [NSUnitAngle] class defines its [BaseUnit] as [Degrees], and provides
// the following units, which are initialized using [NSUnitConverterLinear]
// converters with the specified coefficients:
//
// [Table data omitted]
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle
type UnitAngle struct {
	NSDimension
}

// UnitAngleFromID constructs a [UnitAngle] from an objc.ID.
//
// A unit of measure for planar angle and rotation.
func UnitAngleFromID(id objc.ID) UnitAngle {
	return NSUnitAngle{NSDimension: NSDimensionFromID(id)}
}

// NSUnitAngleFromID is an alias for [UnitAngleFromID] for cross-framework compatibility.
func NSUnitAngleFromID(id objc.ID) UnitAngle { return UnitAngleFromID(id) }

// NOTE: UnitAngle adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UnitAngle] class.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle
type IUnitAngle interface {
	INSDimension
}

// Init initializes the instance.
func (u UnitAngle) Init() UnitAngle {
	rv := objc.Send[UnitAngle](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UnitAngle) Autorelease() UnitAngle {
	rv := objc.Send[UnitAngle](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUnitAngle creates a new UnitAngle instance.
func NewUnitAngle() UnitAngle {
	class := getUnitAngleClass()
	rv := objc.Send[UnitAngle](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewUnitAngleWithCoder(coder INSCoder) UnitAngle {
	instance := getUnitAngleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return UnitAngleFromID(rv)
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
func NewUnitAngleWithSymbol(symbol string) UnitAngle {
	instance := getUnitAngleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:"), objc.String(symbol))
	return UnitAngleFromID(rv)
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
func NewUnitAngleWithSymbolConverter(symbol string, converter INSUnitConverter) UnitAngle {
	instance := getUnitAngleClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSymbol:converter:"), objc.String(symbol), converter)
	return UnitAngleFromID(rv)
}

// The degrees unit of angle.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle/degrees
func (_UnitAngleClass UnitAngleClass) Degrees() UnitAngle {
	rv := objc.Send[objc.ID](objc.ID(_UnitAngleClass.class), objc.Sel("degrees"))
	return NSUnitAngleFromID(objc.ID(rv))
}

// The arc minutes unit of angle.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle/arcMinutes
func (_UnitAngleClass UnitAngleClass) ArcMinutes() UnitAngle {
	rv := objc.Send[objc.ID](objc.ID(_UnitAngleClass.class), objc.Sel("arcMinutes"))
	return NSUnitAngleFromID(objc.ID(rv))
}

// The arc seconds unit of angle.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle/arcSeconds
func (_UnitAngleClass UnitAngleClass) ArcSeconds() UnitAngle {
	rv := objc.Send[objc.ID](objc.ID(_UnitAngleClass.class), objc.Sel("arcSeconds"))
	return NSUnitAngleFromID(objc.ID(rv))
}

// The radians unit of angle.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle/radians
func (_UnitAngleClass UnitAngleClass) Radians() UnitAngle {
	rv := objc.Send[objc.ID](objc.ID(_UnitAngleClass.class), objc.Sel("radians"))
	return NSUnitAngleFromID(objc.ID(rv))
}

// The gradians unit of angle.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle/gradians
func (_UnitAngleClass UnitAngleClass) Gradians() UnitAngle {
	rv := objc.Send[objc.ID](objc.ID(_UnitAngleClass.class), objc.Sel("gradians"))
	return NSUnitAngleFromID(objc.ID(rv))
}

// The revolutions unit of angle.
//
// See: https://developer.apple.com/documentation/Foundation/UnitAngle/revolutions
func (_UnitAngleClass UnitAngleClass) Revolutions() UnitAngle {
	rv := objc.Send[objc.ID](objc.ID(_UnitAngleClass.class), objc.Sel("revolutions"))
	return NSUnitAngleFromID(objc.ID(rv))
}
