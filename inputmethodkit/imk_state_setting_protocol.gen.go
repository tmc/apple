// Code generated from Apple documentation for InputMethodKit. DO NOT EDIT.

package inputmethodkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The [IMKStateSetting] protocol defines methods for setting or accessing values that indicate the state of an input method.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting
type IMKStateSetting interface {
	objectivec.IObject

	// Activates the input method server.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/activateServer(_:)
	ActivateServer(sender objectivec.IObject)

	// Deactivates the input method server.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/deactivateServer(_:)
	DeactivateServer(sender objectivec.IObject)

	// Displays a preferences window.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/showPreferences(_:)
	ShowPreferences(sender objectivec.IObject)

	// Returns an unsigned integer that contains a union of event masks
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/recognizedEvents(_:)
	RecognizedEvents(sender objectivec.IObject) uint

	// Returns the modes dictionary associated with the input method.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/modes(_:)
	Modes(sender objectivec.IObject) foundation.INSDictionary

	// Returns a value object whose key is the provided tag.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/value(forTag:client:)
	ValueForTagClient(tag int, sender objectivec.IObject) objectivec.IObject

	// Set the value for the provided key.
	//
	// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/setValue(_:forTag:client:)
	SetValueForTagClient(value objectivec.IObject, tag int, sender objectivec.IObject)
}

// IMKStateSettingObject wraps an existing Objective-C object that conforms to the IMKStateSetting protocol.
type IMKStateSettingObject struct {
	objectivec.Object
}

func (o IMKStateSettingObject) BaseObject() objectivec.Object {
	return o.Object
}

// IMKStateSettingObjectFromID constructs a [IMKStateSettingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func IMKStateSettingObjectFromID(id objc.ID) IMKStateSettingObject {
	return IMKStateSettingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Activates the input method server.
//
// sender: The object sending the activation message.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/activateServer(_:)
func (o IMKStateSettingObject) ActivateServer(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("activateServer:"), sender)
}

// Deactivates the input method server.
//
// sender: The object sending the deactivation message.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/deactivateServer(_:)
func (o IMKStateSettingObject) DeactivateServer(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deactivateServer:"), sender)
}

// Displays a preferences window.
//
// sender: The object sending the message to show the preference window.
//
// # Discussion
//
// This method looks for a nib file that contains a window controller class
// and a preferences utility. If found, it displays the window. To use this
// method you must create a menu item in your input method menu whose action
// is “. When a user selects that item, the Input Method Kit invokes your “
// method. The default implementation looks for a nib file named
// `preferences.Nib()`. If found, it allocates a window controller class loads
// the nib file. You can provide a custom window controller class by naming
// the class in your input method `info.Plist()` file, providing a key-value
// pair. The key must be [InputMethodServerPreferencesWindowControllerClass]
// and the associated value must be the name of your custom class.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/showPreferences(_:)
func (o IMKStateSettingObject) ShowPreferences(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("showPreferences:"), sender)
}

// Returns an unsigned integer that contains a union of event masks
//
// sender: The client object requesting the supported events.
//
// # Return Value
//
// An unsigned integer that contains a union of event masks (See the
// `NSEvent.H()` header file.
//
// # Discussion
//
// A client calls this method to check whether an input method supports an
// event. The default implementation returns [NSKeyDownMask]. If your input
// method handles only key down events, the Input Method Kit provides the
// default mouse handling. The default mouse-down handling behavior is as
// follows: If there is an active composition area and the user clicks in the
// text but outside of the composition area, the Input Method Kit sends your
// input method a “ message. This happens only for input methods that return
// only the default value—[NSKeyDownMask].
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/recognizedEvents(_:)
func (o IMKStateSettingObject) RecognizedEvents(sender objectivec.IObject) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("recognizedEvents:"), sender)
	return rv
}

// Returns the modes dictionary associated with the input method.
//
// sender: The client object requesting the modes dictionary.
//
// # Return Value
//
// The modes dictionary associated with the input method.
//
// # Discussion
//
// Typically a client object calls this method to to build the text input
// menu. By calling the input method rather than reading the modes from the
// `Info.Plist()` file, the input method can dynamically modify the modes
// supported.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/modes(_:)
func (o IMKStateSettingObject) Modes(sender objectivec.IObject) foundation.INSDictionary {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modes:"), sender)
	return foundation.NSDictionaryFromID(rv)
}

// Returns a value object whose key is the provided tag.
//
// tag: The key whose value you want to retrieve.
//
// sender: The client requesting the value.
//
// # Return Value
//
// The value object.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/value(forTag:client:)
func (o IMKStateSettingObject) ValueForTagClient(tag int, sender objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("valueForTag:client:"), tag, sender)
	return objectivec.Object{ID: rv}
}

// Set the value for the provided key.
//
// value: The value, specified as the appropriate object (such as [NSNumber]), to
// set.
//
// tag: The key whose value you want to set.
//
// sender: The client setting the value.
//
// See: https://developer.apple.com/documentation/InputMethodKit/IMKStateSetting/setValue(_:forTag:client:)
func (o IMKStateSettingObject) SetValueForTagClient(value objectivec.IObject, tag int, sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setValue:forTag:client:"), value, tag, sender)
}
