// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties for an element that represents an absolute or relative input value along an axis.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisElement
type GCAxisElement interface {
	objectivec.IObject
	GCPhysicalInputElement

	// An input object that provides absolute axis values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisElement/absoluteInput
	AbsoluteInput() GCAxisInput

	// An input object that provides relative axis values.
	//
	// See: https://developer.apple.com/documentation/GameController/GCAxisElement/relativeInput
	RelativeInput() GCRelativeInput
}

// GCAxisElementObject wraps an existing Objective-C object that conforms to the GCAxisElement protocol.
type GCAxisElementObject struct {
	objectivec.Object
}

func (o GCAxisElementObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCAxisElementObjectFromID constructs a [GCAxisElementObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCAxisElementObjectFromID(id objc.ID) GCAxisElementObject {
	return GCAxisElementObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// An input object that provides absolute axis values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisElement/absoluteInput
func (o GCAxisElementObject) AbsoluteInput() GCAxisInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("absoluteInput"))
	return GCAxisInputObjectFromID(rv)
}

// An input object that provides relative axis values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisElement/relativeInput
func (o GCAxisElementObject) RelativeInput() GCRelativeInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("relativeInput"))
	return GCRelativeInputObjectFromID(rv)
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (o GCAxisElementObject) LocalizedName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (o GCAxisElementObject) SfSymbolsName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (o GCAxisElementObject) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}
