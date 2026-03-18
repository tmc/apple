// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSInputServerMouseTracker protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSInputServerMouseTracker
type NSInputServerMouseTracker interface {
	objectivec.IObject
}

// NSInputServerMouseTrackerObject wraps an existing Objective-C object that conforms to the NSInputServerMouseTracker protocol.
type NSInputServerMouseTrackerObject struct {
	objectivec.Object
}
func (o NSInputServerMouseTrackerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSInputServerMouseTrackerObjectFromID constructs a [NSInputServerMouseTrackerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSInputServerMouseTrackerObjectFromID(id objc.ID) NSInputServerMouseTrackerObject {
	return NSInputServerMouseTrackerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

