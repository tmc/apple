// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSInputServiceProvider protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSInputServiceProvider
type NSInputServiceProvider interface {
	objectivec.IObject
}



// NSInputServiceProviderObject wraps an existing Objective-C object that conforms to the NSInputServiceProvider protocol.
type NSInputServiceProviderObject struct {
	objectivec.Object
}
func (o NSInputServiceProviderObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSInputServiceProviderObjectFromID constructs a [NSInputServiceProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSInputServiceProviderObjectFromID(id objc.ID) NSInputServiceProviderObject {
	return NSInputServiceProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}










