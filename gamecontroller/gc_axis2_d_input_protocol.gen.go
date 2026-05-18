// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of inputs that provide a normalized point in a two-dimensional coordinate system with a fixed origin.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput
type GCAxis2DInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/isAnalog
	IsAnalog() bool

	// A Boolean value that indicates whether the value wraps when it reaches the range’s minimum or maximum value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/canWrap
	CanWrap() bool

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/isAnalog
	Analog() bool

	// The axis input represented as a normalized point in a two-dimensional coordinate system.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/value
	Value() GCPoint2

	// The time of the most recent value change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/lastValueTimestamp
	LastValueTimestamp() float64

	// The time in seconds between the last value change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/lastValueLatency
	LastValueLatency() float64

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/sources
	Sources() foundation.INSSet
}

// GCAxis2DInputObject wraps an existing Objective-C object that conforms to the GCAxis2DInput protocol.
type GCAxis2DInputObject struct {
	objectivec.Object
}

func (o GCAxis2DInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCAxis2DInputObjectFromID constructs a [GCAxis2DInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCAxis2DInputObjectFromID(id objc.ID) GCAxis2DInputObject {
	return GCAxis2DInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the input provides analog values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/isAnalog
func (o GCAxis2DInputObject) IsAnalog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return rv
}

// A Boolean value that indicates whether the value wraps when it reaches the
// range’s minimum or maximum value.
//
// # Discussion
//
// This property is false for most axis inputs.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/canWrap
func (o GCAxis2DInputObject) CanWrap() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canWrap"))
	return bool(rv)
}

// A Boolean value that indicates whether the input provides analog values.
//
// # Discussion
//
// This property is true for most axis inputs.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/isAnalog
func (o GCAxis2DInputObject) Analog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return bool(rv)
}

// The axis input represented as a normalized point in a two-dimensional
// coordinate system.
//
// # Discussion
//
// The values of the coordinates range between `-1` and `1` where `(0,0)` is
// the fixed origin. Game Controller deadzones and saturates the values so
// there’s no value outside this range. A zero coordinate is inside the
// deadzone and any coordinate greater than or less than zero is outside the
// deadzone.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/value
func (o GCAxis2DInputObject) Value() GCPoint2 {
	rv := objc.Send[GCPoint2](o.ID, objc.Sel("value"))
	return GCPoint2(rv)
}

// The time of the most recent value change.
//
// # Discussion
//
// This property isn’t a specific date and time. To determine the time
// between value changes in seconds, subtract a previous time from the current
// time.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/lastValueTimestamp
func (o GCAxis2DInputObject) LastValueTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastValueTimestamp"))
	return float64(rv)
}

// The time in seconds between the last value change and the current time.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/lastValueLatency
func (o GCAxis2DInputObject) LastValueLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastValueLatency"))
	return float64(rv)
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCAxis2DInput/sources
func (o GCAxis2DInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
