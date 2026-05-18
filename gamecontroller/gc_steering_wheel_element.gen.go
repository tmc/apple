// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCSteeringWheelElement] class.
var (
	_GCSteeringWheelElementClass     GCSteeringWheelElementClass
	_GCSteeringWheelElementClassOnce sync.Once
)

func getGCSteeringWheelElementClass() GCSteeringWheelElementClass {
	_GCSteeringWheelElementClassOnce.Do(func() {
		_GCSteeringWheelElementClass = GCSteeringWheelElementClass{class: objc.GetClass("GCSteeringWheelElement")}
	})
	return _GCSteeringWheelElementClass
}

// GetGCSteeringWheelElementClass returns the class object for GCSteeringWheelElement.
func GetGCSteeringWheelElementClass() GCSteeringWheelElementClass {
	return getGCSteeringWheelElementClass()
}

type GCSteeringWheelElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCSteeringWheelElementClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCSteeringWheelElementClass) Alloc() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The element that represents the wheel of a racing wheel controller.
//
// # Getting the characteristics
//
//   - [GCSteeringWheelElement.MaximumDegreesOfRotation]: The maximum number of degrees that the user can rotate the wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCSteeringWheelElement
type GCSteeringWheelElement struct {
	objectivec.Object
}

// GCSteeringWheelElementFromID constructs a [GCSteeringWheelElement] from an objc.ID.
//
// The element that represents the wheel of a racing wheel controller.
func GCSteeringWheelElementFromID(id objc.ID) GCSteeringWheelElement {
	return GCSteeringWheelElement{objectivec.Object{ID: id}}
}

// NOTE: GCSteeringWheelElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCSteeringWheelElement] class.
//
// # Getting the characteristics
//
//   - [IGCSteeringWheelElement.MaximumDegreesOfRotation]: The maximum number of degrees that the user can rotate the wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCSteeringWheelElement
type IGCSteeringWheelElement interface {
	objectivec.IObject

	// Topic: Getting the characteristics

	// The maximum number of degrees that the user can rotate the wheel.
	MaximumDegreesOfRotation() float32
}

// Init initializes the instance.
func (g GCSteeringWheelElement) Init() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCSteeringWheelElement) Autorelease() GCSteeringWheelElement {
	rv := objc.Send[GCSteeringWheelElement](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCSteeringWheelElement creates a new GCSteeringWheelElement instance.
func NewGCSteeringWheelElement() GCSteeringWheelElement {
	class := getGCSteeringWheelElementClass()
	rv := objc.Send[GCSteeringWheelElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An input object that provides absolute axis values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisElement/absoluteInput
func (g GCSteeringWheelElement) AbsoluteInput() GCAxisInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("absoluteInput"))
	return GCAxisInputObjectFromID(rv)
}

// The element’s aliases to use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/aliases
func (g GCSteeringWheelElement) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(rv)
}

// The localized name for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/localizedName
func (g GCSteeringWheelElement) LocalizedName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// An input object that provides relative axis values.
//
// See: https://developer.apple.com/documentation/GameController/GCAxisElement/relativeInput
func (g GCSteeringWheelElement) RelativeInput() GCRelativeInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("relativeInput"))
	return GCRelativeInputObjectFromID(rv)
}

// A system symbol for the element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputElement/sfSymbolsName
func (g GCSteeringWheelElement) SfSymbolsName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}

// The maximum number of degrees that the user can rotate the wheel.
//
// See: https://developer.apple.com/documentation/GameController/GCSteeringWheelElement/maximumDegreesOfRotation
func (g GCSteeringWheelElement) MaximumDegreesOfRotation() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("maximumDegreesOfRotation"))
	return rv
}

// Protocol methods for GCAxisElement
