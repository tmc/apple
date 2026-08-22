// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The [IMKMouseHandling] protocol defines methods that your input method can implement to handle mouse events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling
type IMKMouseHandling interface {
	objectivec.IObject

	// Handles mouse-down event send to an input method.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseDown(onCharacterIndex:coordinate:withModifier:continueTracking:client:)
	MouseDownOnCharacterIndexCoordinateWithModifierContinueTrackingClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) (bool, bool)

	// Handles a mouse-up event sent to an input method.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseUp(onCharacterIndex:coordinate:withModifier:client:)
	MouseUpOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool

	// Handles a mouse-moved event sent to an input method.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseMoved(onCharacterIndex:coordinate:withModifier:client:)
	MouseMovedOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool
}

// IMKMouseHandlingObject wraps an existing Objective-C object that conforms to the IMKMouseHandling protocol.
type IMKMouseHandlingObject struct {
	objectivec.Object
}

func (o IMKMouseHandlingObject) BaseObject() objectivec.Object {
	return o.Object
}

// IMKMouseHandlingObjectFromID constructs a [IMKMouseHandlingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IMKMouseHandlingObjectFromID(id objc.ID) IMKMouseHandlingObject {
	return IMKMouseHandlingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Handles mouse-down event send to an input method.
//
// index: The index within the sender’s text storage where the mouse-down event
// occurred.
//
// point: The point at which the mouse-down event occurred.
//
// flags: The modifier keys.
//
// keepTracking: Set this parameter to true if you want to receive subsequent mouse-moved
// and mouse -up events.
//
// sender: The client object.
//
// # Return Value
//
// true if handled; otherwise false.
//
// # Discussion
//
// Implement this method if your input method handles mouse-down events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseDown(onCharacterIndex:coordinate:withModifier:continueTracking:client:)
func (o IMKMouseHandlingObject) MouseDownOnCharacterIndexCoordinateWithModifierContinueTrackingClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) (bool, bool) {
	var keepTracking bool
	rv := objc.Send[bool](o.ID, objc.Sel("mouseDownOnCharacterIndex:coordinate:withModifier:continueTracking:client:"), index, point, flags, unsafe.Pointer(&keepTracking), sender)
	return keepTracking, rv
}

// Handles a mouse-up event sent to an input method.
//
// index: The index within the sender’s text storage where the mouse-up event
// occurred.
//
// point: The point at which the mouse-up event occurred.
//
// flags: The modifier keys.
//
// sender: The client object.
//
// # Return Value
//
// true if handled; otherwise false.
//
// # Discussion
//
// Implement this method if your input method handles mouse-up events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseUp(onCharacterIndex:coordinate:withModifier:client:)
func (o IMKMouseHandlingObject) MouseUpOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("mouseUpOnCharacterIndex:coordinate:withModifier:client:"), index, point, flags, sender)
	return rv
}

// Handles a mouse-moved event sent to an input method.
//
// index: The index within the sender’s text storage where the mouse-moved event
// occurred.
//
// point: The point at which the mouse-moved event occurred.
//
// flags: The modifier keys.
//
// sender: The client object.
//
// # Return Value
//
// true if handled; otherwise false.
//
// # Discussion
//
// Implement this method if your input method handles mouse-moved events.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKMouseHandling/mouseMoved(onCharacterIndex:coordinate:withModifier:client:)
func (o IMKMouseHandlingObject) MouseMovedOnCharacterIndexCoordinateWithModifierClient(index uint, point corefoundation.CGPoint, flags uint, sender objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("mouseMovedOnCharacterIndex:coordinate:withModifier:client:"), index, point, flags, sender)
	return rv
}
