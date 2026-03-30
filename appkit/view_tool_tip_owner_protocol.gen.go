// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods for dynamically associating a tool tip with a view.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewToolTipOwner
type NSViewToolTipOwner interface {
	objectivec.IObject

	// Returns the tool tip string to be displayed due to the cursor pausing at location `point` within the tool tip rectangle identified by `tag` in the view `view`.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewToolTipOwner/view(_:stringForToolTip:point:userData:)
	ViewStringForToolTipPointUserData(view INSView, tag NSToolTipTag, point corefoundation.CGPoint, data unsafe.Pointer) string
}

// NSViewToolTipOwnerObject wraps an existing Objective-C object that conforms to the NSViewToolTipOwner protocol.
type NSViewToolTipOwnerObject struct {
	objectivec.Object
}

func (o NSViewToolTipOwnerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSViewToolTipOwnerObjectFromID constructs a [NSViewToolTipOwnerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSViewToolTipOwnerObjectFromID(id objc.ID) NSViewToolTipOwnerObject {
	return NSViewToolTipOwnerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the tool tip string to be displayed due to the cursor pausing at
// location `point` within the tool tip rectangle identified by `tag` in the
// view `view`.
//
// # Discussion
//
// `userData` is additional information provided by the creator of the tool
// tip rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewToolTipOwner/view(_:stringForToolTip:point:userData:)
func (o NSViewToolTipOwnerObject) ViewStringForToolTipPointUserData(view INSView, tag NSToolTipTag, point corefoundation.CGPoint, data unsafe.Pointer) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("view:stringForToolTip:point:userData:"), view, tag, point, data)
	return foundation.NSStringFromID(rv).String()
}
