// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties of elements that represent directional pads.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement
type GCDirectionPadElement interface {
	objectivec.IObject
	GCPhysicalInputElement

	// The input object that represents the left button on the directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/left
	Left() objectivec.IObject

	// The input object that represents the right button on the directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/right
	Right() objectivec.IObject

	// The input object that represents the up button on the directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/up
	Up() objectivec.IObject

	// The input object that represents the down button on the directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/down
	Down() objectivec.IObject

	// The input object that represents the x-axis on the directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/xAxis
	XAxis() GCAxisInput

	// The input object that represents the y-axis on the directional pad.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/yAxis
	YAxis() GCAxisInput

	// The location of the directional pad represented as a point.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/xyAxes
	XyAxes() GCAxis2DInput
}

// GCDirectionPadElementObject wraps an existing Objective-C object that conforms to the GCDirectionPadElement protocol.
type GCDirectionPadElementObject struct {
	objectivec.Object
}

func (o GCDirectionPadElementObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCDirectionPadElementObjectFromID constructs a [GCDirectionPadElementObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCDirectionPadElementObjectFromID(id objc.ID) GCDirectionPadElementObject {
	return GCDirectionPadElementObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The input object that represents the left button on the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/left
func (o GCDirectionPadElementObject) Left() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("left"))
	return objectivec.Object{ID: rv}
}

// The input object that represents the right button on the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/right
func (o GCDirectionPadElementObject) Right() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("right"))
	return objectivec.Object{ID: rv}
}

// The input object that represents the up button on the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/up
func (o GCDirectionPadElementObject) Up() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("up"))
	return objectivec.Object{ID: rv}
}

// The input object that represents the down button on the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/down
func (o GCDirectionPadElementObject) Down() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("down"))
	return objectivec.Object{ID: rv}
}

// The input object that represents the x-axis on the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/xAxis
func (o GCDirectionPadElementObject) XAxis() GCAxisInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("xAxis"))
	return GCAxisInputObjectFromID(rv)
}

// The input object that represents the y-axis on the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/yAxis
func (o GCDirectionPadElementObject) YAxis() GCAxisInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("yAxis"))
	return GCAxisInputObjectFromID(rv)
}

// The location of the directional pad represented as a point.
//
// See: https://developer.apple.com/documentation/GameController/GCDirectionPadElement/xyAxes
func (o GCDirectionPadElementObject) XyAxes() GCAxis2DInput {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("xyAxes"))
	return GCAxis2DInputObjectFromID(rv)
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (o GCDirectionPadElementObject) LocalizedName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (o GCDirectionPadElementObject) SfSymbolsName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (o GCDirectionPadElementObject) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}
