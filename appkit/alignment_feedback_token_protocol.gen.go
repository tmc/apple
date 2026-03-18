// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSAlignmentFeedbackToken protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlignmentFeedbackToken
type NSAlignmentFeedbackToken interface {
	objectivec.IObject
}

// NSAlignmentFeedbackTokenObject wraps an existing Objective-C object that conforms to the NSAlignmentFeedbackToken protocol.
type NSAlignmentFeedbackTokenObject struct {
	objectivec.Object
}
func (o NSAlignmentFeedbackTokenObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAlignmentFeedbackTokenObjectFromID constructs a [NSAlignmentFeedbackTokenObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAlignmentFeedbackTokenObjectFromID(id objc.ID) NSAlignmentFeedbackTokenObject {
	return NSAlignmentFeedbackTokenObject{
		Object: objectivec.ObjectFromID(id),
	}
}

