// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCControllerElement] class.
var (
	_GCControllerElementClass     GCControllerElementClass
	_GCControllerElementClassOnce sync.Once
)

func getGCControllerElementClass() GCControllerElementClass {
	_GCControllerElementClassOnce.Do(func() {
		_GCControllerElementClass = GCControllerElementClass{class: objc.GetClass("GCControllerElement")}
	})
	return _GCControllerElementClass
}

// GetGCControllerElementClass returns the class object for GCControllerElement.
func GetGCControllerElementClass() GCControllerElementClass {
	return getGCControllerElementClass()
}

type GCControllerElementClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerElementClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerElementClass) Alloc() GCControllerElement {
	rv := objc.Send[GCControllerElement](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// An input for a physical control, such as a button or thumbstick.
//
// # Overview
//
// [GCControllerElement] is an abstract superclass for specific types of
// elements that represent controls on a game controller. Use the respective
// subclasses to either get the input of an element directly or set a handler
// that the element calls when the user changes a value. This class provides
// support for common features.
//
// For complex elements that have subelements, you can get the containing
// element using the [GCControllerElement.Collection] property. For example, a direction pad
// ([GCControllerDirectionPad]) has two axis control and four button
// subelements.
//
// If the user binds a controller element to a system gesture, the system
// sends the input to the system gesture recognizer first. If it doesn’t
// recognize a gesture, the system sends the input to your app but with a
// delay. If it does recognize a gesture, it doesn’t send any input to your
// app.
//
// To change this default behavior, you can set the
// [GCControllerElement.PreferredSystemGestureState] property to
// [GCSystemGestureStateAlwaysReceive] to receive the input simultaneously
// without delay. Alternatively, set it to [GCSystemGestureStateDisabled] to
// disable the system gesture and receive the input exclusively. Use the
// [GCControllerElement.BoundToSystemGesture] property to check whether the user included an
// element in a system gesture.
//
// Use the [GCControllerElement.Analog] property to determine whether an element’s input value
// is a range of values or a discrete digital value.
//
// # Accessing input values
//
//   - [GCControllerElement.IsAnalog]: A Boolean value that indicates whether the element provides analog data.
//
// # Getting a localized name
//
//   - [GCControllerElement.LocalizedName]: The localized name for the element or the remapped element.
//   - [GCControllerElement.SetLocalizedName]
//   - [GCControllerElement.UnmappedLocalizedName]: The element’s localized name, not the remapped name.
//   - [GCControllerElement.SetUnmappedLocalizedName]
//
// # Displaying a symbol
//
//   - [GCControllerElement.SfSymbolsName]: A system symbol for the element or the remapped element.
//   - [GCControllerElement.SetSfSymbolsName]
//   - [GCControllerElement.UnmappedSfSymbolsName]: The element’s system symbol, not the remapped symbol.
//   - [GCControllerElement.SetUnmappedSfSymbolsName]
//
// # Accessing elements by key
//
//   - [GCControllerElement.Aliases]: The element’s aliases you use when accessing it with the subscript notation.
//
// # Getting the containing element
//
//   - [GCControllerElement.Collection]: The enclosing element for this element.
//
// # Handling system gesture input
//
//   - [GCControllerElement.IsBoundToSystemGesture]: A Boolean value that indicates whether the user binds the element to a system gesture.
//   - [GCControllerElement.PreferredSystemGestureState]: The preferred state for handling input when the user binds the element to a system gesture.
//   - [GCControllerElement.SetPreferredSystemGestureState]
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement
type GCControllerElement struct {
	objectivec.Object
}

// GCControllerElementFromID constructs a [GCControllerElement] from an objc.ID.
//
// An input for a physical control, such as a button or thumbstick.
func GCControllerElementFromID(id objc.ID) GCControllerElement {
	return GCControllerElement{objectivec.Object{ID: id}}
}

// NOTE: GCControllerElement adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerElement] class.
//
// # Accessing input values
//
//   - [IGCControllerElement.IsAnalog]: A Boolean value that indicates whether the element provides analog data.
//
// # Getting a localized name
//
//   - [IGCControllerElement.LocalizedName]: The localized name for the element or the remapped element.
//   - [IGCControllerElement.SetLocalizedName]
//   - [IGCControllerElement.UnmappedLocalizedName]: The element’s localized name, not the remapped name.
//   - [IGCControllerElement.SetUnmappedLocalizedName]
//
// # Displaying a symbol
//
//   - [IGCControllerElement.SfSymbolsName]: A system symbol for the element or the remapped element.
//   - [IGCControllerElement.SetSfSymbolsName]
//   - [IGCControllerElement.UnmappedSfSymbolsName]: The element’s system symbol, not the remapped symbol.
//   - [IGCControllerElement.SetUnmappedSfSymbolsName]
//
// # Accessing elements by key
//
//   - [IGCControllerElement.Aliases]: The element’s aliases you use when accessing it with the subscript notation.
//
// # Getting the containing element
//
//   - [IGCControllerElement.Collection]: The enclosing element for this element.
//
// # Handling system gesture input
//
//   - [IGCControllerElement.IsBoundToSystemGesture]: A Boolean value that indicates whether the user binds the element to a system gesture.
//   - [IGCControllerElement.PreferredSystemGestureState]: The preferred state for handling input when the user binds the element to a system gesture.
//   - [IGCControllerElement.SetPreferredSystemGestureState]
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement
type IGCControllerElement interface {
	objectivec.IObject

	// Topic: Accessing input values

	// A Boolean value that indicates whether the element provides analog data.
	IsAnalog() bool

	// Topic: Getting a localized name

	// The localized name for the element or the remapped element.
	LocalizedName() string
	SetLocalizedName(value string)
	// The element’s localized name, not the remapped name.
	UnmappedLocalizedName() string
	SetUnmappedLocalizedName(value string)

	// Topic: Displaying a symbol

	// A system symbol for the element or the remapped element.
	SfSymbolsName() string
	SetSfSymbolsName(value string)
	// The element’s system symbol, not the remapped symbol.
	UnmappedSfSymbolsName() string
	SetUnmappedSfSymbolsName(value string)

	// Topic: Accessing elements by key

	// The element’s aliases you use when accessing it with the subscript notation.
	Aliases() foundation.INSSet

	// Topic: Getting the containing element

	// The enclosing element for this element.
	Collection() IGCControllerElement

	// Topic: Handling system gesture input

	// A Boolean value that indicates whether the user binds the element to a system gesture.
	IsBoundToSystemGesture() bool
	// The preferred state for handling input when the user binds the element to a system gesture.
	PreferredSystemGestureState() GCSystemGestureState
	SetPreferredSystemGestureState(value GCSystemGestureState)
}

// Init initializes the instance.
func (g GCControllerElement) Init() GCControllerElement {
	rv := objc.Send[GCControllerElement](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerElement) Autorelease() GCControllerElement {
	rv := objc.Send[GCControllerElement](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerElement creates a new GCControllerElement instance.
func NewGCControllerElement() GCControllerElement {
	class := getGCControllerElementClass()
	rv := objc.Send[GCControllerElement](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the element provides analog data.
//
// # Discussion
//
// If this property is true, the input value defined by the element can return
// a range (from a minimum to maximum) of possible values. For example, this
// element might be a pressure-sensitive button or an axis of a thumbstick
// that allows for a range of physical movement. If this property is false,
// the input value is a discrete value, such as `0` if the element is off, and
// `1` if the element is on.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/isAnalog
func (g GCControllerElement) IsAnalog() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isAnalog"))
	return rv
}

// The localized name for the element or the remapped element.
//
// # Discussion
//
// If the user remaps this element, this property is the remapped localized
// name.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/localizedName
func (g GCControllerElement) LocalizedName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCControllerElement) SetLocalizedName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setLocalizedName:"), objc.String(value))
}

// The element’s localized name, not the remapped name.
//
// # Discussion
//
// To present the element that a user wants to remap in your interface, use
// this property to get the original name. Otherwise, use the [LocalizedName]
// property to get the possibly remapped name.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/unmappedLocalizedName
func (g GCControllerElement) UnmappedLocalizedName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unmappedLocalizedName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCControllerElement) SetUnmappedLocalizedName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setUnmappedLocalizedName:"), objc.String(value))
}

// A system symbol for the element or the remapped element.
//
// # Discussion
//
// If the user remaps this element, this property is the remapped system
// symbol.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/sfSymbolsName
func (g GCControllerElement) SfSymbolsName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("sfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCControllerElement) SetSfSymbolsName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setSfSymbolsName:"), objc.String(value))
}

// The element’s system symbol, not the remapped symbol.
//
// # Discussion
//
// Use this property to get the original unmapped name. Otherwise, use the
// [SfSymbolsName] property to get the possibly remapped symbol.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/unmappedSfSymbolsName
func (g GCControllerElement) UnmappedSfSymbolsName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unmappedSfSymbolsName"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCControllerElement) SetUnmappedSfSymbolsName(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setUnmappedSfSymbolsName:"), objc.String(value))
}

// The element’s aliases you use when accessing it with the subscript
// notation.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/aliases
func (g GCControllerElement) Aliases() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("aliases"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The enclosing element for this element.
//
// # Discussion
//
// If this element is part of another element, this property is the containing
// element; otherwise, it’s `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/collection
func (g GCControllerElement) Collection() IGCControllerElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("collection"))
	return GCControllerElementFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the user binds the element to a
// system gesture.
//
// # Discussion
//
// This property is true if the user binds this element to a gesture;
// otherwise, it’s false.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/isBoundToSystemGesture
func (g GCControllerElement) IsBoundToSystemGesture() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isBoundToSystemGesture"))
	return rv
}

// The preferred state for handling input when the user binds the element to a
// system gesture.
//
// # Discussion
//
// In rare situations, you may use this property to disable system gestures.
// However, the system isn’t guaranteed to respect this property. The
// default value for this property is [GCSystemGestureStateEnabled].
//
// See: https://developer.apple.com/documentation/GameController/GCControllerElement/preferredSystemGestureState
func (g GCControllerElement) PreferredSystemGestureState() GCSystemGestureState {
	rv := objc.Send[GCSystemGestureState](g.ID, objc.Sel("preferredSystemGestureState"))
	return GCSystemGestureState(rv)
}
func (g GCControllerElement) SetPreferredSystemGestureState(value GCSystemGestureState) {
	objc.Send[struct{}](g.ID, objc.Sel("setPreferredSystemGestureState:"), value)
}
