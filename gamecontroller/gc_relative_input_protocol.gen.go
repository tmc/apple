// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of inputs that provide positions along an axis that are relative to the previous position.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput
type GCRelativeInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/isAnalog
	IsAnalog() bool

	// A Boolean value that indicates whether the input provides analog values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/isAnalog
	Analog() bool

	// The most recent amount of change in values that the profile records.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/delta
	Delta() float32

	// A timestamp for when the profile reports the delta value.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/lastDeltaTimestamp
	LastDeltaTimestamp() float64

	// The time in seconds between the current and the previous delta values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/lastDeltaLatency
	LastDeltaLatency() float64

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/sources
	Sources() foundation.INSSet
}

// GCRelativeInputObject wraps an existing Objective-C object that conforms to the GCRelativeInput protocol.
type GCRelativeInputObject struct {
	objectivec.Object
}

func (o GCRelativeInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCRelativeInputObjectFromID constructs a [GCRelativeInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCRelativeInputObjectFromID(id objc.ID) GCRelativeInputObject {
	return GCRelativeInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the input provides analog values.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/isAnalog
func (o GCRelativeInputObject) IsAnalog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return rv
}

// A Boolean value that indicates whether the input provides analog values.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/isAnalog
func (o GCRelativeInputObject) Analog() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAnalog"))
	return bool(rv)
}

// The most recent amount of change in values that the profile records.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/delta
func (o GCRelativeInputObject) Delta() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("delta"))
	return float32(rv)
}

// A timestamp for when the profile reports the delta value.
//
// # Discussion
//
// This property isn’t a specific date and time. To determine the time
// between delta values, subtract a previous value from the current value.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/lastDeltaTimestamp
func (o GCRelativeInputObject) LastDeltaTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastDeltaTimestamp"))
	return float64(rv)
}

// The time in seconds between the current and the previous delta values.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/lastDeltaLatency
func (o GCRelativeInputObject) LastDeltaLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastDeltaLatency"))
	return float64(rv)
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCRelativeInput/sources
func (o GCRelativeInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
