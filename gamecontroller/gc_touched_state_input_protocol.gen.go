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

	// The time of the most recent touch state change.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateTimestamp
	LastTouchedStateTimestamp() float64

	// The time in seconds between the last touch state change and the current time.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateLatency
	LastTouchedStateLatency() float64

	// A block that the element calls when its touch value changes.
	//
	// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/touchedDidChangeHandler
	TouchedDidChangeHandler() func(objc.ID, bool)

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

// The time of the most recent touch state change.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateTimestamp
func (o GCTouchedStateInputObject) LastTouchedStateTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastTouchedStateTimestamp"))
	return rv
}

// The time in seconds between the last touch state change and the current
// time.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/lastTouchedStateLatency
func (o GCTouchedStateInputObject) LastTouchedStateLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastTouchedStateLatency"))
	return rv
}

// A block that the element calls when its touch value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/touchedDidChangeHandler
func (o GCTouchedStateInputObject) TouchedDidChangeHandler() func(objc.ID, bool) {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("touchedDidChangeHandler"))
	// Block/function return - cannot convert from objc.ID to Go func
	_ = rv
	return nil
}

// One or more physical actions the user performs to manipulate the input.
//
// See: https://developer.apple.com/documentation/GameController/GCTouchedStateInput/sources
func (o GCTouchedStateInputObject) Sources() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sources"))
	return foundation.NSSetFromID(rv)
}
