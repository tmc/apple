// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties for an element that has touch state input.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput
type GCTouchedStateInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the user touches the button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/isTouched
	IsTouched() bool

	// A Boolean value that indicates whether the user touches the button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/isTouched
	Touched() bool

	// The time of the most recent touch state change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateTimestamp
	LastTouchedStateTimestamp() foundation.NSTimeInterval

	// The time in seconds between the last touch state change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateLatency
	LastTouchedStateLatency() foundation.NSTimeInterval

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/sources
	Sources() foundation.INSSet
}

// GCTouchedStateInputObject wraps an existing Objective-C object that conforms to the GCTouchedStateInput protocol.
type GCTouchedStateInputObject struct {
	objectivec.Object
}

func (o GCTouchedStateInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCTouchedStateInputObjectFromID constructs a [GCTouchedStateInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCTouchedStateInputObjectFromID(id objc.ID) GCTouchedStateInputObject {
	return GCTouchedStateInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the user touches the button.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/isTouched
func (o GCTouchedStateInputObject) IsTouched() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isTouched"))
	return rv
}

// A Boolean value that indicates whether the user touches the button.
//
// # Discussion
//
// For controllers that support capacitive touch, the user can start touching
// the button without pressure when the value property is `0`. For controllers
// that don’t support capacitive touch, the user starts touching the button
// when the value property is greater than `0`.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/isTouched
func (o GCTouchedStateInputObject) Touched() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isTouched"))
	return bool(rv)
}

// The time of the most recent touch state change.
//
// # Discussion
//
// This property isn’t a specific date and time. To determine the time
// between changes, subtract a previous value from the current value.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateTimestamp
func (o GCTouchedStateInputObject) LastTouchedStateTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastTouchedStateTimestamp"))
	return foundation.NSTimeInterval(rv)
}

// The time in seconds between the last touch state change and the current
// time.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateLatency
func (o GCTouchedStateInputObject) LastTouchedStateLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastTouchedStateLatency"))
	return foundation.NSTimeInterval(rv)
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/sources
func (o GCTouchedStateInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
