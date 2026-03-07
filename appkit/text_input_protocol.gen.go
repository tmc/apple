// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that text views need to implement to interact properly with the text input management system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInput
type NSTextInput interface {
	objectivec.IObject
}



// NSTextInputObject wraps an existing Objective-C object that conforms to the NSTextInput protocol.
type NSTextInputObject struct {
	objectivec.Object
}
func (o NSTextInputObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSTextInputObjectFromID constructs a [NSTextInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextInputObjectFromID(id objc.ID) NSTextInputObject {
	return NSTextInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}










