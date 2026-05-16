// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCPhysicalInputProfile] class.
var (
	_GCPhysicalInputProfileClass     GCPhysicalInputProfileClass
	_GCPhysicalInputProfileClassOnce sync.Once
)

func getGCPhysicalInputProfileClass() GCPhysicalInputProfileClass {
	_GCPhysicalInputProfileClassOnce.Do(func() {
		_GCPhysicalInputProfileClass = GCPhysicalInputProfileClass{class: objc.GetClass("GCPhysicalInputProfile")}
	})
	return _GCPhysicalInputProfileClass
}

// GetGCPhysicalInputProfileClass returns the class object for GCPhysicalInputProfile.
func GetGCPhysicalInputProfileClass() GCPhysicalInputProfileClass {
	return getGCPhysicalInputProfileClass()
}

type GCPhysicalInputProfileClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCPhysicalInputProfileClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCPhysicalInputProfileClass) Alloc() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The base class for controller profiles that support physical buttons,
// thumbsticks, and directional pads.
//
// # Overview
//
// This class provides properties and methods for accessing common elements of
// controllers, and for creating snapshots of profiles.
//
// # Getting the device
//
//   - [GCPhysicalInputProfile.Device]: The physical device that the profile represents.
//
// # Getting change information
//
//   - [GCPhysicalInputProfile.LastEventTimestamp]: The time of the most recent change to an element’s value.
//   - [GCPhysicalInputProfile.ValueDidChangeHandler]: The block that the profile calls when an element’s value changes.
//   - [GCPhysicalInputProfile.SetValueDidChangeHandler]
//
// # Accessing elements by name or key
//
//   - [GCPhysicalInputProfile.Elements]: The elements in the profile as key-value pairs for lookup by name.
//   - [GCPhysicalInputProfile.Buttons]: The buttons in the profile as key-value pairs for lookup by name.
//   - [GCPhysicalInputProfile.Axes]: The axes in the profile as key-value pairs for lookup by name.
//   - [GCPhysicalInputProfile.Dpads]: The directional pads in the profile as key-value pairs for lookup by name.
//   - [GCPhysicalInputProfile.Touchpads]: The touchpads in the profile as key-value pairs for lookup by name.
//   - [GCPhysicalInputProfile.ObjectForKeyedSubscript]: Returns the element that the key specifies.
//
// # Getting elements by type
//
//   - [GCPhysicalInputProfile.AllElements]: The elements in the profile.
//   - [GCPhysicalInputProfile.AllButtons]: The buttons in the profile.
//   - [GCPhysicalInputProfile.AllAxes]: The axes in the profile.
//   - [GCPhysicalInputProfile.AllDpads]: The directional pads in the profile.
//   - [GCPhysicalInputProfile.AllTouchpads]: The touchpads in the profile.
//
// # Setting snapshot values
//
//   - [GCPhysicalInputProfile.Capture]: Returns a snapshot of the profile with its current element values.
//   - [GCPhysicalInputProfile.SetStateFromPhysicalInput]: Copies the input values from a specified physical input profile to a snapshot of the profile.
//
// # Remapping input elements
//
//   - [GCPhysicalInputProfile.HasRemappedElements]: A Boolean value that indicates whether the user remaps elements in this profile.
//   - [GCPhysicalInputProfile.MappedElementAliasForPhysicalInputName]: Returns the name of the input element to which the user remaps the given physical element.
//   - [GCPhysicalInputProfile.MappedPhysicalInputNamesForElementAlias]: Returns the physical input elements to which the user remaps the given input element.
//   - [GCPhysicalInputProfile.GCControllerUserCustomizationsDidChange]: A notification that posts when the user customizes the button mappings or other settings of a controller.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile
type GCPhysicalInputProfile struct {
	objectivec.Object
}

// GCPhysicalInputProfileFromID constructs a [GCPhysicalInputProfile] from an objc.ID.
//
// The base class for controller profiles that support physical buttons,
// thumbsticks, and directional pads.
func GCPhysicalInputProfileFromID(id objc.ID) GCPhysicalInputProfile {
	return GCPhysicalInputProfile{objectivec.Object{ID: id}}
}

// NOTE: GCPhysicalInputProfile adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCPhysicalInputProfile] class.
//
// # Getting the device
//
//   - [IGCPhysicalInputProfile.Device]: The physical device that the profile represents.
//
// # Getting change information
//
//   - [IGCPhysicalInputProfile.LastEventTimestamp]: The time of the most recent change to an element’s value.
//   - [IGCPhysicalInputProfile.ValueDidChangeHandler]: The block that the profile calls when an element’s value changes.
//   - [IGCPhysicalInputProfile.SetValueDidChangeHandler]
//
// # Accessing elements by name or key
//
//   - [IGCPhysicalInputProfile.Elements]: The elements in the profile as key-value pairs for lookup by name.
//   - [IGCPhysicalInputProfile.Buttons]: The buttons in the profile as key-value pairs for lookup by name.
//   - [IGCPhysicalInputProfile.Axes]: The axes in the profile as key-value pairs for lookup by name.
//   - [IGCPhysicalInputProfile.Dpads]: The directional pads in the profile as key-value pairs for lookup by name.
//   - [IGCPhysicalInputProfile.Touchpads]: The touchpads in the profile as key-value pairs for lookup by name.
//   - [IGCPhysicalInputProfile.ObjectForKeyedSubscript]: Returns the element that the key specifies.
//
// # Getting elements by type
//
//   - [IGCPhysicalInputProfile.AllElements]: The elements in the profile.
//   - [IGCPhysicalInputProfile.AllButtons]: The buttons in the profile.
//   - [IGCPhysicalInputProfile.AllAxes]: The axes in the profile.
//   - [IGCPhysicalInputProfile.AllDpads]: The directional pads in the profile.
//   - [IGCPhysicalInputProfile.AllTouchpads]: The touchpads in the profile.
//
// # Setting snapshot values
//
//   - [IGCPhysicalInputProfile.Capture]: Returns a snapshot of the profile with its current element values.
//   - [IGCPhysicalInputProfile.SetStateFromPhysicalInput]: Copies the input values from a specified physical input profile to a snapshot of the profile.
//
// # Remapping input elements
//
//   - [IGCPhysicalInputProfile.HasRemappedElements]: A Boolean value that indicates whether the user remaps elements in this profile.
//   - [IGCPhysicalInputProfile.MappedElementAliasForPhysicalInputName]: Returns the name of the input element to which the user remaps the given physical element.
//   - [IGCPhysicalInputProfile.MappedPhysicalInputNamesForElementAlias]: Returns the physical input elements to which the user remaps the given input element.
//   - [IGCPhysicalInputProfile.GCControllerUserCustomizationsDidChange]: A notification that posts when the user customizes the button mappings or other settings of a controller.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile
type IGCPhysicalInputProfile interface {
	objectivec.IObject

	// Topic: Getting the device

	// The physical device that the profile represents.
	Device() GCDevice

	// Topic: Getting change information

	// The time of the most recent change to an element’s value.
	LastEventTimestamp() float64
	// The block that the profile calls when an element’s value changes.
	ValueDidChangeHandler() GCPhysicalInputProfileGCControllerElementHandler
	SetValueDidChangeHandler(value GCPhysicalInputProfileGCControllerElementHandler)

	// Topic: Accessing elements by name or key

	// The elements in the profile as key-value pairs for lookup by name.
	Elements() foundation.INSDictionary
	// The buttons in the profile as key-value pairs for lookup by name.
	Buttons() foundation.INSDictionary
	// The axes in the profile as key-value pairs for lookup by name.
	Axes() foundation.INSDictionary
	// The directional pads in the profile as key-value pairs for lookup by name.
	Dpads() foundation.INSDictionary
	// The touchpads in the profile as key-value pairs for lookup by name.
	Touchpads() foundation.INSDictionary
	// Returns the element that the key specifies.
	ObjectForKeyedSubscript(key string) IGCControllerElement

	// Topic: Getting elements by type

	// The elements in the profile.
	AllElements() foundation.INSSet
	// The buttons in the profile.
	AllButtons() foundation.INSSet
	// The axes in the profile.
	AllAxes() foundation.INSSet
	// The directional pads in the profile.
	AllDpads() foundation.INSSet
	// The touchpads in the profile.
	AllTouchpads() foundation.INSSet

	// Topic: Setting snapshot values

	// Returns a snapshot of the profile with its current element values.
	Capture() IGCPhysicalInputProfile
	// Copies the input values from a specified physical input profile to a snapshot of the profile.
	SetStateFromPhysicalInput(physicalInput IGCPhysicalInputProfile)

	// Topic: Remapping input elements

	// A Boolean value that indicates whether the user remaps elements in this profile.
	HasRemappedElements() bool
	// Returns the name of the input element to which the user remaps the given physical element.
	MappedElementAliasForPhysicalInputName(inputName string) string
	// Returns the physical input elements to which the user remaps the given input element.
	MappedPhysicalInputNamesForElementAlias(elementAlias string) foundation.INSSet
	// A notification that posts when the user customizes the button mappings or other settings of a controller.
	GCControllerUserCustomizationsDidChange() foundation.NSString

	// The extended gamepad profile.
	ExtendedGamepad() IGCExtendedGamepad
	SetExtendedGamepad(value IGCExtendedGamepad)
	// The gamepad profile.
	Gamepad() unsafe.Pointer
	SetGamepad(value unsafe.Pointer)
	// The micro gamepad profile.
	MicroGamepad() IGCMicroGamepad
	SetMicroGamepad(value IGCMicroGamepad)
	// The motion input profile.
	Motion() IGCMotion
	SetMotion(value IGCMotion)
	// The physical input profile for the controller.
	PhysicalInputProfile() IGCPhysicalInputProfile
	SetPhysicalInputProfile(value IGCPhysicalInputProfile)
}

// Init initializes the instance.
func (g GCPhysicalInputProfile) Init() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCPhysicalInputProfile) Autorelease() GCPhysicalInputProfile {
	rv := objc.Send[GCPhysicalInputProfile](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCPhysicalInputProfile creates a new GCPhysicalInputProfile instance.
func NewGCPhysicalInputProfile() GCPhysicalInputProfile {
	class := getGCPhysicalInputProfileClass()
	rv := objc.Send[GCPhysicalInputProfile](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the element that the key specifies.
//
// key: A key that identifies an element.
//
// # Return Value
//
// The element that matches the key.
//
// # Discussion
//
// You can access elements of a profile using a subscript notation. For
// example, get the button with the X label from an instance of
// [GCMicroGamepad] using `microGamepad[”Button X”]`.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/subscript(_:)
func (g GCPhysicalInputProfile) ObjectForKeyedSubscript(key string) IGCControllerElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(key))
	return GCControllerElementFromID(rv)
}

// Returns a snapshot of the profile with its current element values.
//
// # Return Value
//
// A snapshot of the profile.
//
// # Discussion
//
// A snapshot is a copy of profile at a moment in time with its current
// element values. Unlike other profiles, you can set the values of a
// snapshot’s elements.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/capture()
func (g GCPhysicalInputProfile) Capture() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCPhysicalInputProfileFromID(rv)
}

// Copies the input values from a specified physical input profile to a
// snapshot of the profile.
//
// physicalInput: The physical input profile to copy the input values from.
//
// # Discussion
//
// If the associated controller isn’t a snapshot, this method does nothing.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/setStateFromPhysicalInput(_:)
func (g GCPhysicalInputProfile) SetStateFromPhysicalInput(physicalInput IGCPhysicalInputProfile) {
	objc.Send[objc.ID](g.ID, objc.Sel("setStateFromPhysicalInput:"), physicalInput)
}

// Returns the name of the input element to which the user remaps the given
// physical element.
//
// inputName: The name of the physical element. For possible values, see [Extended
// gamepad input names].
//
// # Return Value
//
// The name of the input element to which the user remaps the physical
// element, or `nil` if the user doesn’t remap the physical element.
//
// # Discussion
//
// Use this method to get the alias for an input element. For example, if the
// user remaps a physical press of the controller’s A button to button B,
// then passing [GCInputButtonA] to this method returns [GCInputButtonB].
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedElementAlias(forPhysicalInputName:)
//
// [Extended gamepad input names]: https://developer.apple.com/documentation/GameController/extended-gamepad-input-names
// [GCInputButtonA]: https://developer.apple.com/documentation/GameController/GCInputButtonA-8z15w
// [GCInputButtonB]: https://developer.apple.com/documentation/GameController/GCInputButtonB-6z361
func (g GCPhysicalInputProfile) MappedElementAliasForPhysicalInputName(inputName string) string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mappedElementAliasForPhysicalInputName:"), objc.String(inputName))
	return foundation.NSStringFromID(rv).String()
}

// Returns the physical input elements to which the user remaps the given
// input element.
//
// elementAlias: The name of the input element too which physical input elements remap. For
// possible values, see [Extended gamepad input names].
//
// # Return Value
//
// The names of the physical input element to which the user remaps the given
// element.
//
// # Discussion
//
// For example, if the user maps a physical press of A button , B button, and
// X button to button B, then passing [GCInputButtonB] returns a set that
// contains [GCInputButtonA], [GCInputButtonB], and [GCInputButtonX].
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/mappedPhysicalInputNames(forElementAlias:)
//
// [Extended gamepad input names]: https://developer.apple.com/documentation/GameController/extended-gamepad-input-names
// [GCInputButtonA]: https://developer.apple.com/documentation/GameController/GCInputButtonA-8z15w
// [GCInputButtonB]: https://developer.apple.com/documentation/GameController/GCInputButtonB-6z361
// [GCInputButtonX]: https://developer.apple.com/documentation/GameController/GCInputButtonX-32i2z
func (g GCPhysicalInputProfile) MappedPhysicalInputNamesForElementAlias(elementAlias string) foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("mappedPhysicalInputNamesForElementAlias:"), objc.String(elementAlias))
	return foundation.NSSetFromID(rv)
}

// The physical device that the profile represents.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/device
func (g GCPhysicalInputProfile) Device() GCDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// The time of the most recent change to an element’s value.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/lastEventTimestamp
func (g GCPhysicalInputProfile) LastEventTimestamp() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("lastEventTimestamp"))
	return rv
}

// The block that the profile calls when an element’s value changes.
//
// # Discussion
//
// The block’s parameters are:
//
// `profile`: The controller profile that contains the element. `element`: The
// element with the value that changes.
//
// If multiple elements change values at the same time, the profile calls this
// block once for each element that changes. If the value of a subelement
// changes, the profile only calls the block for the containing element.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/valueDidChangeHandler
func (g GCPhysicalInputProfile) ValueDidChangeHandler() GCPhysicalInputProfileGCControllerElementHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("valueDidChangeHandler"))
	_ = rv
	return nil
}
func (g GCPhysicalInputProfile) SetValueDidChangeHandler(value GCPhysicalInputProfileGCControllerElementHandler) {
	block, cleanup := NewGCPhysicalInputProfileGCControllerElementBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setValueDidChangeHandler:"), block)
}

// The elements in the profile as key-value pairs for lookup by name.
//
// # Discussion
//
// Use this property to access elements by name. For example, use the name
// `“Button A”` to get the face button of an extended gamepad profile.
//
// For more button names, see [Extended gamepad input names].
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/elements
//
// [Extended gamepad input names]: https://developer.apple.com/documentation/GameController/extended-gamepad-input-names
func (g GCPhysicalInputProfile) Elements() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elements"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The buttons in the profile as key-value pairs for lookup by name.
//
// # Discussion
//
// Use the [GCInputXboxPaddleOne] constant to get the P1 paddle button for an
// Xbox controller.
//
// For more button names, see [Extended gamepad input names] and [Xbox
// controller input names].
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/buttons
//
// [Extended gamepad input names]: https://developer.apple.com/documentation/GameController/extended-gamepad-input-names
// [GCInputXboxPaddleOne]: https://developer.apple.com/documentation/GameController/GCInputXboxPaddleOne-offv
// [Xbox controller input names]: https://developer.apple.com/documentation/GameController/xbox-controller-input-names
func (g GCPhysicalInputProfile) Buttons() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttons"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The axes in the profile as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/axes
func (g GCPhysicalInputProfile) Axes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("axes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The directional pads in the profile as key-value pairs for lookup by name.
//
// # Discussion
//
// Use the [GCInputDualShockTouchpadOne] name to access an element on a
// DualShock controller.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/dpads
//
// [GCInputDualShockTouchpadOne]: https://developer.apple.com/documentation/GameController/GCInputDualShockTouchpadOne-55wgz
func (g GCPhysicalInputProfile) Dpads() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dpads"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The touchpads in the profile as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/touchpads
func (g GCPhysicalInputProfile) Touchpads() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpads"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// The elements in the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allElements
func (g GCPhysicalInputProfile) AllElements() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allElements"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The buttons in the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allButtons
func (g GCPhysicalInputProfile) AllButtons() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allButtons"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The axes in the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allAxes
func (g GCPhysicalInputProfile) AllAxes() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allAxes"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The directional pads in the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allDpads
func (g GCPhysicalInputProfile) AllDpads() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allDpads"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// The touchpads in the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/allTouchpads
func (g GCPhysicalInputProfile) AllTouchpads() foundation.INSSet {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("allTouchpads"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the user remaps elements in this
// profile.
//
// # Discussion
//
// If true, the user remaps one or more elements; otherwise, this property is
// false.
//
// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputProfile/hasRemappedElements
func (g GCPhysicalInputProfile) HasRemappedElements() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("hasRemappedElements"))
	return rv
}

// A notification that posts when the user customizes the button mappings or
// other settings of a controller.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCControllerUserCustomizationsDidChange
func (g GCPhysicalInputProfile) GCControllerUserCustomizationsDidChange() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCControllerUserCustomizationsDidChange"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The extended gamepad profile.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/extendedgamepad
func (g GCPhysicalInputProfile) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("extendedGamepad"))
	return GCExtendedGamepadFromID(objc.ID(rv))
}
func (g GCPhysicalInputProfile) SetExtendedGamepad(value IGCExtendedGamepad) {
	objc.Send[struct{}](g.ID, objc.Sel("setExtendedGamepad:"), value)
}

// The gamepad profile.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/gamepad
func (g GCPhysicalInputProfile) Gamepad() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("gamepad"))
	return rv
}
func (g GCPhysicalInputProfile) SetGamepad(value unsafe.Pointer) {
	objc.Send[struct{}](g.ID, objc.Sel("setGamepad:"), value)
}

// The micro gamepad profile.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/microgamepad
func (g GCPhysicalInputProfile) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("microGamepad"))
	return GCMicroGamepadFromID(objc.ID(rv))
}
func (g GCPhysicalInputProfile) SetMicroGamepad(value IGCMicroGamepad) {
	objc.Send[struct{}](g.ID, objc.Sel("setMicroGamepad:"), value)
}

// The motion input profile.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/motion
func (g GCPhysicalInputProfile) Motion() IGCMotion {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("motion"))
	return GCMotionFromID(objc.ID(rv))
}
func (g GCPhysicalInputProfile) SetMotion(value IGCMotion) {
	objc.Send[struct{}](g.ID, objc.Sel("setMotion:"), value)
}

// The physical input profile for the controller.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/physicalinputprofile
func (g GCPhysicalInputProfile) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("physicalInputProfile"))
	return GCPhysicalInputProfileFromID(objc.ID(rv))
}
func (g GCPhysicalInputProfile) SetPhysicalInputProfile(value IGCPhysicalInputProfile) {
	objc.Send[struct{}](g.ID, objc.Sel("setPhysicalInputProfile:"), value)
}
