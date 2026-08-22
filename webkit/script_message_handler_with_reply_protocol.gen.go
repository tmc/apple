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

	// Tells the handler that a webpage sent a script message that included a reply.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKScriptMessageHandlerWithReply/userContentController(_:didReceive:replyHandler:)
	UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController IWKUserContentController, message IWKScriptMessage, replyHandler IObjectStringHandler)
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

// Tells the handler that a webpage sent a script message that included a
// reply.
//
// userContentController: The user content controller that delivered the message to your handler.
//
// message: An object that contains the message details.
//
// replyHandler: A reply handler block to execute with the response to send back to the
// webpage. This block has no return value and takes the following parameters:
//
// reply: An object that contains the data to return to the webpage. Allowed
// types for this parameter are [NSNumber], [NSString], [NSDate], [NSArray],
// [NSDictionary], and [NSNull]. Specify `nil` if an error occurred.
// errorMessage: `nil` on success, or a string that describes the error that
// occurred.
//
// # Discussion
//
// Use this method to handle a message from the webpage and provide an
// appropriate response.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessageHandlerWithReply/userContentController(_:didReceive:replyHandler:)
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (o WKScriptMessageHandlerWithReplyObject) UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController IWKUserContentController, message IWKScriptMessage, replyHandler IObjectStringHandler) {
	_block2, _ := NewIObjectStringBlock(replyHandler)
	objc.Send[struct{}](o.ID, objc.Sel("userContentController:didReceiveScriptMessage:replyHandler:"), userContentController, message, _block2)
}
