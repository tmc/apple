// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface for receiving messages from JavaScript code running in a webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessageHandler
type WKScriptMessageHandler interface {
	objectivec.IObject

	// Tells the handler that a webpage sent a script message.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKScriptMessageHandler/userContentController(_:didReceive:)
	UserContentControllerDidReceiveScriptMessage(userContentController IWKUserContentController, message IWKScriptMessage)
}

// WKScriptMessageHandlerObject wraps an existing Objective-C object that conforms to the WKScriptMessageHandler protocol.
type WKScriptMessageHandlerObject struct {
	objectivec.Object
}

func (o WKScriptMessageHandlerObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKScriptMessageHandlerObjectFromID constructs a [WKScriptMessageHandlerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKScriptMessageHandlerObjectFromID(id objc.ID) WKScriptMessageHandlerObject {
	return WKScriptMessageHandlerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the handler that a webpage sent a script message.
//
// userContentController: The user content controller that delivered the message to your handler.
//
// message: An object that contains the message details.
//
// # Discussion
//
// Use this method to respond to a message sent from the webpage’s
// JavaScript code. Use the message parameter to get the message contents and
// to determine the originating web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessageHandler/userContentController(_:didReceive:)
func (o WKScriptMessageHandlerObject) UserContentControllerDidReceiveScriptMessage(userContentController IWKUserContentController, message IWKScriptMessage) {
	objc.Send[struct{}](o.ID, objc.Sel("userContentController:didReceiveScriptMessage:"), userContentController, message)
}
