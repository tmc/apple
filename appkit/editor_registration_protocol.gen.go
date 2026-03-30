// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that controllers can implement to enable an editor view to inform the controller when it has uncommitted changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSEditorRegistration
type NSEditorRegistration interface {
	objectivec.IObject
}

// NSEditorRegistrationObject wraps an existing Objective-C object that conforms to the NSEditorRegistration protocol.
type NSEditorRegistrationObject struct {
	objectivec.Object
}

func (o NSEditorRegistrationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSEditorRegistrationObjectFromID constructs a [NSEditorRegistrationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSEditorRegistrationObjectFromID(id objc.ID) NSEditorRegistrationObject {
	return NSEditorRegistrationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppKit/NSEditorRegistration/objectDidBeginEditing(_:)
func (o NSEditorRegistrationObject) ObjectDidBeginEditing(editor NSEditor) {
	objc.Send[struct{}](o.ID, objc.Sel("objectDidBeginEditing:"), editor)
}

// See: https://developer.apple.com/documentation/AppKit/NSEditorRegistration/objectDidEndEditing(_:)
func (o NSEditorRegistrationObject) ObjectDidEndEditing(editor NSEditor) {
	objc.Send[struct{}](o.ID, objc.Sel("objectDidEndEditing:"), editor)
}
