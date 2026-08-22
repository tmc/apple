// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCGearShifterElement] class.
var (
	_GCGearShifterElementClass     GCGearShifterElementClass
	_GCGearShifterElementClassOnce sync.Once
)

func getGCGearShifterElementClass() GCGearShifterElementClass {
	_GCGearShifterElementClassOnce.Do(func() {
		_GCGearShifterElementClass = GCGearShifterElementClass{class: objc.GetClass("GCGearShifterElement")}
	})
	return _GCGearShifterElementClass
}

// GetGCGearShifterElementClass returns the class object for GCGearShifterElement.
func GetGCGearShifterElementClass() GCGearShifterElementClass {
	return getGCGearShifterElementClass()
}

type GCGearShifterElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCGearShifterElementClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCGearShifterElementClass) Alloc() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// An element that represents either a pattern or a sequential gear shift.
//
// # Accessing input values
//
//   - [GCGearShifterElement.PatternInput]: The input object for a pattern gear shift.
//   - [GCGearShifterElement.SequentialInput]: The input object for a sequential gear shift.
//
// See: https://developer.apple.com/documentation/GameController/GCGearShifterElement
type GCGearShifterElement struct {
	objectivec.Object
}

// GCGearShifterElementFromID constructs a [GCGearShifterElement] from an objc.ID.
//
// An element that represents either a pattern or a sequential gear shift.
func GCGearShifterElementFromID(id objc.ID) GCGearShifterElement {
	return GCGearShifterElement{objectivec.Object{ID: id}}
}

// NOTE: GCGearShifterElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCGearShifterElement] class.
//
// # Accessing input values
//
//   - [IGCGearShifterElement.PatternInput]: The input object for a pattern gear shift.
//   - [IGCGearShifterElement.SequentialInput]: The input object for a sequential gear shift.
//
// See: https://developer.apple.com/documentation/GameController/GCGearShifterElement
type IGCGearShifterElement interface {
	objectivec.IObject

	// Topic: Accessing input values

	// The input object for a pattern gear shift.
	PatternInput() GCSwitchPositionInput
	// The input object for a sequential gear shift.
	SequentialInput() GCRelativeInput

	// The element’s aliases to use when accessing it with the subscript notation.
	Aliases() foundation.INSSet
	// The localized name for the element.
	LocalizedName() string
	// A system symbol for the element.
	SfSymbolsName() string
}

// Init initializes the instance.
func (g GCGearShifterElement) Init() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCGearShifterElement) Autorelease() GCGearShifterElement {
	rv := objc.Send[GCGearShifterElement](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCGearShifterElement creates a new GCGearShifterElement instance.
func NewGCGearShifterElement() GCGearShifterElement {
	class := getGCGearShifterElementClass()
	rv := objc.Send[GCGearShifterElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (g GCGearShifterElement) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (g GCGearShifterElement) LocalizedName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (g GCGearShifterElement) SfSymbolsName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The input object for a pattern gear shift.
//
// # Discussion
//
// If this property is `nil`, the gear shift isn’t a pattern gear shift. A
// pattern gear shift lays out the gears in a pattern that lets the user move
// to any gear. If the [Position] property of this property is `0`, the gear
// shift is in neutral. If it’s `-1`, the gear shift is in reverse.
//
// See: https://developer.apple.com/documentation/GameController/GCGearShifterElement/patternInput
func (g GCGearShifterElement) PatternInput() GCSwitchPositionInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("patternInput"))
	return GCSwitchPositionInputObjectFromID(rv)
}

// The input object for a sequential gear shift.
//
// # Discussion
//
// If this property is `nil`, this gear shift isn’t a sequential gear shift.
// A sequential gear shift requires the user to move through the gears in
// sequence.
//
// See: https://developer.apple.com/documentation/GameController/GCGearShifterElement/sequentialInput
func (g GCGearShifterElement) SequentialInput() GCRelativeInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sequentialInput"))
	return GCRelativeInputObjectFromID(rv)
}

// Protocol methods for GCPhysicalInputElement
