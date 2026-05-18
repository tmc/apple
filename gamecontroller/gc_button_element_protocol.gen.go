// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of an element that represents a momentary switch, such as a push button.
//
// See: https://developer.apple.com/documentation/GameController/GCButtonElement
type GCButtonElement interface {
	objectivec.IObject
	GCPhysicalInputElement

	// The input object that provides the touch state of the element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCButtonElement/touchedInput
	TouchedInput() GCTouchedStateInput

	// The input object that provides the linear and press state of the element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCButtonElement/pressedInput
	PressedInput() objectivec.IObject

	// Get the input containing the measured force applied to the button.
	//
	// See: https://developer.apple.com/documentation/GameController/GCButtonElement/forceInput
	ForceInput() GCLinearInput
}

// GCButtonElementObject wraps an existing Objective-C object that conforms to the GCButtonElement protocol.
type GCButtonElementObject struct {
	objectivec.Object
}

func (o GCButtonElementObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCButtonElementObjectFromID constructs a [GCButtonElementObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCButtonElementObjectFromID(id objc.ID) GCButtonElementObject {
	return GCButtonElementObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (o GCButtonElementObject) LocalizedName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (o GCButtonElementObject) SfSymbolsName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (o GCButtonElementObject) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}

// The input object that provides the touch state of the element.
//
// See: https://developer.apple.com/documentation/GameController/GCButtonElement/touchedInput
func (o GCButtonElementObject) TouchedInput() GCTouchedStateInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("touchedInput"))
	return GCTouchedStateInputObjectFromID(rv)
}

// The input object that provides the linear and press state of the element.
//
// See: https://developer.apple.com/documentation/GameController/GCButtonElement/pressedInput
func (o GCButtonElementObject) PressedInput() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pressedInput"))
	return objectivec.Object{ID: rv}
}

// Get the input containing the measured force applied to the button.
//
// # Discussion
//
// Some buttons feature load cells (also known as button force transducers)
// capable of measuring applied mechanical force.
//
// See: https://developer.apple.com/documentation/GameController/GCButtonElement/forceInput
func (o GCButtonElementObject) ForceInput() GCLinearInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("forceInput"))
	return GCLinearInputObjectFromID(rv)
}
