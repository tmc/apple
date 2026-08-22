// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of inputs that provide absolute values along an axis with a fixed origin.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput
type GCAxisInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the value wraps when it reaches the range’s minimum or maximum value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/canWrap
	CanWrap() bool

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/isAnalog
	IsAnalog() bool

	// The value along the axis, in unit coordinates.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/value
	Value() float32

	// The time of the most recent value change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueTimestamp
	LastValueTimestamp() foundation.NSTimeInterval

	// The time in seconds between the last value change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueLatency
	LastValueLatency() foundation.NSTimeInterval

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisInput/sources
	Sources() foundation.INSSet
}

// GCAxisInputObject wraps an existing Objective-C object that conforms to the GCAxisInput protocol.
type GCAxisInputObject struct {
	objectivec.Object
}

func (o GCAxisInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCAxisInputObjectFromID constructs a [GCAxisInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCAxisInputObjectFromID(id objc.ID) GCAxisInputObject {
	return GCAxisInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the value wraps when it reaches the
// range’s minimum or maximum value.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/canWrap
func (o GCAxisInputObject) CanWrap() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canWrap"))
	return bool(rv)
}

// A Boolean value that indicates whether the input provides analog values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/isAnalog
func (o GCAxisInputObject) IsAnalog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return bool(rv)
}

// The value along the axis, in unit coordinates.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/value
func (o GCAxisInputObject) Value() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("value"))
	return float32(rv)
}

// The time of the most recent value change.
//
// # Discussion
//
// This property isn’t a specific date and time. To determine the time
// between value changes, subtract a previous time from the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueTimestamp
func (o GCAxisInputObject) LastValueTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastValueTimestamp"))
	return foundation.NSTimeInterval(rv)
}

// The time in seconds between the last value change and the current time.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/lastValueLatency
func (o GCAxisInputObject) LastValueLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastValueLatency"))
	return foundation.NSTimeInterval(rv)
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisInput/sources
func (o GCAxisInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
