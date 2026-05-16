// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for a description of an element without any system-level remapping of the controls.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource
type GCPhysicalInputSource interface {
	objectivec.IObject

	// The localized name for the element without any system-level remapping of the controls.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/elementLocalizedName
	ElementLocalizedName() string

	// A system symbol for the element without any system-level remapping of the controls.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/sfSymbolsName
	SfSymbolsName() string

	// The element’s true aliases without any system-level remapping of the controls.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/elementAliases
	ElementAliases() foundation.INSSet

	// The directional input, if any, that a physical input source involves.
	//
	// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/direction
	Direction() GCPhysicalInputSourceDirection
}

// GCPhysicalInputSourceObject wraps an existing Objective-C object that conforms to the GCPhysicalInputSource protocol.
type GCPhysicalInputSourceObject struct {
	objectivec.Object
}

func (o GCPhysicalInputSourceObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCPhysicalInputSourceObjectFromID constructs a [GCPhysicalInputSourceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCPhysicalInputSourceObjectFromID(id objc.ID) GCPhysicalInputSourceObject {
	return GCPhysicalInputSourceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The localized name for the element without any system-level remapping of
// the controls.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/elementLocalizedName
func (o GCPhysicalInputSourceObject) ElementLocalizedName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elementLocalizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element without any system-level remapping of the
// controls.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/sfSymbolsName
func (o GCPhysicalInputSourceObject) SfSymbolsName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The element’s true aliases without any system-level remapping of the
// controls.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/elementAliases
func (o GCPhysicalInputSourceObject) ElementAliases() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elementAliases"))
	return foundation.NSSetFromID(rv)
}

// The directional input, if any, that a physical input source involves.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSource/direction
func (o GCPhysicalInputSourceObject) Direction() GCPhysicalInputSourceDirection {
	rv := objc.Send[GCPhysicalInputSourceDirection](o.ID, objc.Sel("direction"))
	return rv
}
