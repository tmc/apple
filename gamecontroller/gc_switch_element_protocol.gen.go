// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties for an element that represents a switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchElement
type GCSwitchElement interface {
	objectivec.IObject
	GCPhysicalInputElement

	// The input object that provides the position of the switch.
	//
	// See: https://developer.apple.com/documentation/GameController/GCSwitchElement/positionInput
	PositionInput() GCSwitchPositionInput
}

// GCSwitchElementObject wraps an existing Objective-C object that conforms to the GCSwitchElement protocol.
type GCSwitchElementObject struct {
	objectivec.Object
}

func (o GCSwitchElementObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCSwitchElementObjectFromID constructs a [GCSwitchElementObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCSwitchElementObjectFromID(id objc.ID) GCSwitchElementObject {
	return GCSwitchElementObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (o GCSwitchElementObject) LocalizedName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (o GCSwitchElementObject) SfSymbolsName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (o GCSwitchElementObject) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}

// The input object that provides the position of the switch.
//
// See: https://developer.apple.com/documentation/GameController/GCSwitchElement/positionInput
func (o GCSwitchElementObject) PositionInput() GCSwitchPositionInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("positionInput"))
	return GCSwitchPositionInputObjectFromID(rv)
}
