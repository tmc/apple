// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// # Overview  A protocol to request information from NSView subclasses about the selected content in the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewContentSelectionInfo
type NSViewContentSelectionInfo interface {
	objectivec.IObject

	// SelectionAnchorRect protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewContentSelectionInfo/selectionAnchorRect
	SelectionAnchorRect() corefoundation.CGRect
}

// NSViewContentSelectionInfoObject wraps an existing Objective-C object that conforms to the NSViewContentSelectionInfo protocol.
type NSViewContentSelectionInfoObject struct {
	objectivec.Object
}

func (o NSViewContentSelectionInfoObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSViewContentSelectionInfoObjectFromID constructs a [NSViewContentSelectionInfoObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSViewContentSelectionInfoObjectFromID(id objc.ID) NSViewContentSelectionInfoObject {
	return NSViewContentSelectionInfoObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/AppKit/NSViewContentSelectionInfo/selectionAnchorRect
func (o NSViewContentSelectionInfoObject) SelectionAnchorRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("selectionAnchorRect"))
	return rv
}
