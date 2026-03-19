// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// # Overview  When the user selects a color in an [NSColorPanel](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel>) object, the panel tries to call this method on the first responder. You can override this method in any responder that needs to respond to a color change.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorChanging
type NSColorChanging interface {
	objectivec.IObject

	// Sent to the first responder when the user selects a color in an NSColorPanel object.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSColorChanging/changeColor(_:)
	ChangeColor(sender INSColorPanel)
}

// NSColorChangingObject wraps an existing Objective-C object that conforms to the NSColorChanging protocol.
type NSColorChangingObject struct {
	objectivec.Object
}
func (o NSColorChangingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSColorChangingObjectFromID constructs a [NSColorChangingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSColorChangingObjectFromID(id objc.ID) NSColorChangingObject {
	return NSColorChangingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sent to the first responder when the user selects a color in an
// NSColorPanel object.
//
// sender: The [NSColorPanel] sending the message.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorChanging/changeColor(_:)
func (o NSColorChangingObject) ChangeColor(sender INSColorPanel) {
	
	objc.Send[struct{}](o.ID, objc.Sel("changeColor:"), sender)
	}

