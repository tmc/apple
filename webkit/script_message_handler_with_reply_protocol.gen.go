// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface for responding to messages from JavaScript code running in a webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessageHandlerWithReply
type WKScriptMessageHandlerWithReply interface {
	objectivec.IObject
}

// WKScriptMessageHandlerWithReplyObject wraps an existing Objective-C object that conforms to the WKScriptMessageHandlerWithReply protocol.
type WKScriptMessageHandlerWithReplyObject struct {
	objectivec.Object
}

func (o WKScriptMessageHandlerWithReplyObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKScriptMessageHandlerWithReplyObjectFromID constructs a [WKScriptMessageHandlerWithReplyObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKScriptMessageHandlerWithReplyObjectFromID(id objc.ID) WKScriptMessageHandlerWithReplyObject {
	return WKScriptMessageHandlerWithReplyObject{
		Object: objectivec.ObjectFromID(id),
	}
}
