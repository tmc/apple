// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties for an element that has press state input, such as input from a button.
//
// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput
type GCPressedStateInput interface {
	objectivec.IObject

	// A Boolean value that indicates whether the user presses the button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/isPressed
	IsPressed() bool

	// A Boolean value that indicates whether the user presses the button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/isPressed
	Pressed() bool

	// The time of the most recent press state change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/lastPressedStateTimestamp
	LastPressedStateTimestamp() float64

	// The time in seconds between the last press state change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/lastPressedStateLatency
	LastPressedStateLatency() float64

	// One or more physical actions the user performs to manipulate the input.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/sources
	Sources() foundation.INSSet
}

// GCPressedStateInputObject wraps an existing Objective-C object that conforms to the GCPressedStateInput protocol.
type GCPressedStateInputObject struct {
	objectivec.Object
}

func (o GCPressedStateInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCPressedStateInputObjectFromID constructs a [GCPressedStateInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCPressedStateInputObjectFromID(id objc.ID) GCPressedStateInputObject {
	return GCPressedStateInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A Boolean value that indicates whether the user presses the button.
//
// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/isPressed
func (o GCPressedStateInputObject) IsPressed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPressed"))
	return rv
}

// A Boolean value that indicates whether the user presses the button.
//
// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/isPressed
func (o GCPressedStateInputObject) Pressed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPressed"))
	return bool(rv)
}

// The time of the most recent press state change.
//
// # Discussion
//
// This property isn’t a specific date and time. To determine the time
// between changes, subtract a previous value from the current value.
//
// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/lastPressedStateTimestamp
func (o GCPressedStateInputObject) LastPressedStateTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastPressedStateTimestamp"))
	return float64(rv)
}

// The time in seconds between the last press state change and the current
// time.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event.
//
// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/lastPressedStateLatency
func (o GCPressedStateInputObject) LastPressedStateLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastPressedStateLatency"))
	return float64(rv)
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCPressedStateInput/sources
func (o GCPressedStateInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
