// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that restoration classes must implement to handle the recreation of windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowRestoration
type NSWindowRestoration interface {
	objectivec.IObject
}

// NSWindowRestorationObject wraps an existing Objective-C object that conforms to the NSWindowRestoration protocol.
type NSWindowRestorationObject struct {
	objectivec.Object
}
func (o NSWindowRestorationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSWindowRestorationObjectFromID constructs a [NSWindowRestorationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSWindowRestorationObjectFromID(id objc.ID) NSWindowRestorationObject {
	return NSWindowRestorationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

