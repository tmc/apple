// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKRange] class.
var (
	_SKRangeClass     SKRangeClass
	_SKRangeClassOnce sync.Once
)

func getSKRangeClass() SKRangeClass {
	_SKRangeClassOnce.Do(func() {
		_SKRangeClass = SKRangeClass{class: objc.GetClass("SKRange")}
	})
	return _SKRangeClass
}

// GetSKRangeClass returns the class object for SKRange.
func GetSKRangeClass() SKRangeClass {
	return getSKRangeClass()
}

type SKRangeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKRangeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKRangeClass) Alloc() SKRange {
	rv := objc.Send[SKRange](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A definition of a range of floating-point values.
//
// # Overview
//
// You typically use a [SKRange] to clamp a value so that it is within the
// specified range.
//
// # Creating a Range Object
//
//   - [SKRange.InitWithLowerLimitUpperLimit]: Initializes a new range object.
//
// # Inspecting a Range Object’s Limits
//
//   - [SKRange.LowerLimit]: The minimum possible value.
//   - [SKRange.SetLowerLimit]
//   - [SKRange.UpperLimit]: The maximum possible value.
//   - [SKRange.SetUpperLimit]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange
type SKRange struct {
	objectivec.Object
}

// SKRangeFromID constructs a [SKRange] from an objc.ID.
//
// A definition of a range of floating-point values.
func SKRangeFromID(id objc.ID) SKRange {
	return SKRange{objectivec.Object{ID: id}}
}

// NOTE: SKRange adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKRange] class.
//
// # Creating a Range Object
//
//   - [ISKRange.InitWithLowerLimitUpperLimit]: Initializes a new range object.
//
// # Inspecting a Range Object’s Limits
//
//   - [ISKRange.LowerLimit]: The minimum possible value.
//   - [ISKRange.SetLowerLimit]
//   - [ISKRange.UpperLimit]: The maximum possible value.
//   - [ISKRange.SetUpperLimit]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange
type ISKRange interface {
	objectivec.IObject

	// Topic: Creating a Range Object

	// Initializes a new range object.
	InitWithLowerLimitUpperLimit(lower float64, upper float64) SKRange

	// Topic: Inspecting a Range Object’s Limits

	// The minimum possible value.
	LowerLimit() float64
	SetLowerLimit(value float64)
	// The maximum possible value.
	UpperLimit() float64
	SetUpperLimit(value float64)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (r SKRange) Init() SKRange {
	rv := objc.Send[SKRange](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SKRange) Autorelease() SKRange {
	rv := objc.Send[SKRange](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKRange creates a new SKRange instance.
func NewSKRange() SKRange {
	class := getSKRangeClass()
	rv := objc.Send[SKRange](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and initializes a new range object that specifies a constant value.
//
// value: A constant.
//
// # Return Value
//
// A newly initialized range object whose minimum and maximum value are both
// equal to `value`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/init(constantValue:)
func NewRangeWithConstantValue(value float64) SKRange {
	rv := objc.Send[objc.ID](objc.ID(getSKRangeClass().class), objc.Sel("rangeWithConstantValue:"), value)
	return SKRangeFromID(rv)
}

// Creates and initializes a new range object that specifies only a minimum
// value.
//
// lower: The minimum value for the range.
//
// # Return Value
//
// A newly initialized range object whose minimum value is `lower` and whose
// maximum value is `+Inf`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/init(lowerLimit:)
func NewRangeWithLowerLimit(lower float64) SKRange {
	rv := objc.Send[objc.ID](objc.ID(getSKRangeClass().class), objc.Sel("rangeWithLowerLimit:"), lower)
	return SKRangeFromID(rv)
}

// Initializes a new range object.
//
// lower: The minimum value for the range.
//
// upper: The maximum value for the range.
//
// # Return Value
//
// A newly initialized range object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/init(lowerLimit:upperLimit:)
func NewRangeWithLowerLimitUpperLimit(lower float64, upper float64) SKRange {
	instance := getSKRangeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLowerLimit:upperLimit:"), lower, upper)
	return SKRangeFromID(rv)
}

// Creates and initializes a new range object that specifies only a maximum
// value.
//
// upper: The maximum value for the range.
//
// # Return Value
//
// A newly initialized range object whose minimum value is `-Inf` and whose
// maximum value is `upper`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/init(upperLimit:)
func NewRangeWithUpperLimit(upper float64) SKRange {
	rv := objc.Send[objc.ID](objc.ID(getSKRangeClass().class), objc.Sel("rangeWithUpperLimit:"), upper)
	return SKRangeFromID(rv)
}

// Creates and initializes a new range object using a value and a maximum
// distance from that value.
//
// value: The midpoint for the range.
//
// variance: The maximum amount that a value may differ from the midpoint.
//
// # Return Value
//
// A newly initialized range object whose minimum value is `value-variance`
// and whose maximum value is `value+variance`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/init(value:variance:)
func NewRangeWithValueVariance(value float64, variance float64) SKRange {
	rv := objc.Send[objc.ID](objc.ID(getSKRangeClass().class), objc.Sel("rangeWithValue:variance:"), value, variance)
	return SKRangeFromID(rv)
}

// Initializes a new range object.
//
// lower: The minimum value for the range.
//
// upper: The maximum value for the range.
//
// # Return Value
//
// A newly initialized range object.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/init(lowerLimit:upperLimit:)
func (r SKRange) InitWithLowerLimitUpperLimit(lower float64, upper float64) SKRange {
	rv := objc.Send[SKRange](r.ID, objc.Sel("initWithLowerLimit:upperLimit:"), lower, upper)
	return rv
}
func (r SKRange) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](r.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and initializes a new range object that encompasses all possible
// values.
//
// # Return Value
//
// A newly initialized range object whose minimum value is `—Inf` and whose
// maximum value is `+Inf`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/withNoLimits()
func (_SKRangeClass SKRangeClass) RangeWithNoLimits() SKRange {
	rv := objc.Send[objc.ID](objc.ID(_SKRangeClass.class), objc.Sel("rangeWithNoLimits"))
	return SKRangeFromID(rv)
}

// The minimum possible value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/lowerLimit
func (r SKRange) LowerLimit() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("lowerLimit"))
	return rv
}
func (r SKRange) SetLowerLimit(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setLowerLimit:"), value)
}

// The maximum possible value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRange/upperLimit
func (r SKRange) UpperLimit() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("upperLimit"))
	return rv
}
func (r SKRange) SetUpperLimit(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setUpperLimit:"), value)
}
