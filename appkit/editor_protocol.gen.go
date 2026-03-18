// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSEditor protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSEditor
type NSEditor interface {
	objectivec.IObject

	// CommitEditing protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditing()
	CommitEditing() bool

	// CommitEditingAndReturnError protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditingWithoutPresentingError()
	CommitEditingAndReturnError() (bool, error)

	// DiscardEditing protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSEditor/discardEditing()
	DiscardEditing()
}

// NSEditorObject wraps an existing Objective-C object that conforms to the NSEditor protocol.
type NSEditorObject struct {
	objectivec.Object
}
func (o NSEditorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSEditorObjectFromID constructs a [NSEditorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSEditorObjectFromID(id objc.ID) NSEditorObject {
	return NSEditorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditing()

func (o NSEditorObject) CommitEditing() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("commitEditing"))
	return rv
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditing(withDelegate:didCommit:contextInfo:)

func (o NSEditorObject) CommitEditingWithDelegateDidCommitSelectorContextInfo(delegate objectivec.IObject, didCommitSelector objc.SEL, contextInfo uintptr) {
	
	objc.Send[struct{}](o.ID, objc.Sel("commitEditingWithDelegate:didCommitSelector:contextInfo:"), delegate, didCommitSelector, contextInfo)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSEditor/commitEditingWithoutPresentingError()

func (o NSEditorObject) CommitEditingAndReturnError() (bool, error) {
	
	rv, err := objc.SendWithError[bool](o.ID, objc.Sel("commitEditingAndReturnError:"))
	if err != nil {
		return false, err
	}
	return rv, nil
	}

// See: https://developer.apple.com/documentation/AppKit/NSEditor/discardEditing()

func (o NSEditorObject) DiscardEditing() {
	
	objc.Send[struct{}](o.ID, objc.Sel("discardEditing"))
	}

