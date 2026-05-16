// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of physical input elements.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement
type GCPhysicalInputElement interface {
	objectivec.IObject

	// The localized name for the element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
	LocalizedName() string

	// A system symbol for the element.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
	SfSymbolsName() string

	// The element’s aliases to use when accessing it with the subscript notation.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
	Aliases() foundation.INSSet
}

// GCPhysicalInputElementObject wraps an existing Objective-C object that conforms to the GCPhysicalInputElement protocol.
type GCPhysicalInputElementObject struct {
	objectivec.Object
}

func (o GCPhysicalInputElementObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCPhysicalInputElementObjectFromID constructs a [GCPhysicalInputElementObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCPhysicalInputElementObjectFromID(id objc.ID) GCPhysicalInputElementObject {
	return GCPhysicalInputElementObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (o GCPhysicalInputElementObject) LocalizedName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (o GCPhysicalInputElementObject) SfSymbolsName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (o GCPhysicalInputElementObject) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}
