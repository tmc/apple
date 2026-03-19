// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A protocol that describes specific kinds of input content types.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContent
type NSTextContent interface {
	objectivec.IObject

	// The semantic meaning for a text input area.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextContent/contentType
	ContentType() NSTextContentType

	// The semantic meaning for a text input area.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextContent/contentType
	SetContentType(value NSTextContentType)
}

// NSTextContentObject wraps an existing Objective-C object that conforms to the NSTextContent protocol.
type NSTextContentObject struct {
	objectivec.Object
}
func (o NSTextContentObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextContentObjectFromID constructs a [NSTextContentObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextContentObjectFromID(id objc.ID) NSTextContentObject {
	return NSTextContentObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The semantic meaning for a text input area.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContent/contentType
func (o NSTextContentObject) ContentType() NSTextContentType {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentType"))
	return NSTextContentType(foundation.NSStringFromID(rv).String())
	}

func (o NSTextContentObject) SetContentType(value NSTextContentType) {
	objc.Send[struct{}](o.ID, objc.Sel("setContentType:"), objc.String(string(value)))
}

