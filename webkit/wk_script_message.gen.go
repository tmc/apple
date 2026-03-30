// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKScriptMessage] class.
var (
	_WKScriptMessageClass     WKScriptMessageClass
	_WKScriptMessageClassOnce sync.Once
)

func getWKScriptMessageClass() WKScriptMessageClass {
	_WKScriptMessageClassOnce.Do(func() {
		_WKScriptMessageClass = WKScriptMessageClass{class: objc.GetClass("WKScriptMessage")}
	})
	return _WKScriptMessageClass
}

// GetWKScriptMessageClass returns the class object for WKScriptMessage.
func GetWKScriptMessageClass() WKScriptMessageClass {
	return getWKScriptMessageClass()
}

type WKScriptMessageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKScriptMessageClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKScriptMessageClass) Alloc() WKScriptMessage {
	rv := objc.Send[WKScriptMessage](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that encapsulates a message sent by JavaScript code from a
// webpage.
//
// # Overview
//
// Use a [WKScriptMessage] object to get details about a JavaScript message
// sent to a custom message handler in your app. You don’t create
// [WKScriptMessage] objects directly. When JavaScript code targets one of
// your app’s message handlers, the [WKUserContentController] object of the
// web view creates a [WKScriptMessage] object and delivers it to the message
// handler’s delegate method. Use the object you’re provided to process
// the message and provide an appropriate response.
//
// For more information about handling script messages, see the
// [WKScriptMessageHandler] and [WKScriptMessageHandlerWithReply] protocols.
// For information about how to register message handlers, see the methods of
// [WKUserContentController].
//
// # Getting the Message Contents
//
//   - [WKScriptMessage.Body]: The body of the message.
//
// # Getting Message-Related Information
//
//   - [WKScriptMessage.FrameInfo]: The frame that sent the message.
//   - [WKScriptMessage.WebView]: The web view that sent the message.
//   - [WKScriptMessage.World]: The namespace in which the JavaScript code executes.
//
// # Getting the Message Handler’s Name
//
//   - [WKScriptMessage.Name]: The name of the message handler to which the message is sent.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage
type WKScriptMessage struct {
	objectivec.Object
}

// WKScriptMessageFromID constructs a [WKScriptMessage] from an objc.ID.
//
// An object that encapsulates a message sent by JavaScript code from a
// webpage.
func WKScriptMessageFromID(id objc.ID) WKScriptMessage {
	return WKScriptMessage{objectivec.Object{ID: id}}
}

// NOTE: WKScriptMessage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKScriptMessage] class.
//
// # Getting the Message Contents
//
//   - [IWKScriptMessage.Body]: The body of the message.
//
// # Getting Message-Related Information
//
//   - [IWKScriptMessage.FrameInfo]: The frame that sent the message.
//   - [IWKScriptMessage.WebView]: The web view that sent the message.
//   - [IWKScriptMessage.World]: The namespace in which the JavaScript code executes.
//
// # Getting the Message Handler’s Name
//
//   - [IWKScriptMessage.Name]: The name of the message handler to which the message is sent.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage
type IWKScriptMessage interface {
	objectivec.IObject

	// Topic: Getting the Message Contents

	// The body of the message.
	Body() objectivec.IObject

	// Topic: Getting Message-Related Information

	// The frame that sent the message.
	FrameInfo() IWKFrameInfo
	// The web view that sent the message.
	WebView() IWKWebView
	// The namespace in which the JavaScript code executes.
	World() IWKContentWorld

	// Topic: Getting the Message Handler’s Name

	// The name of the message handler to which the message is sent.
	Name() string
}

// Init initializes the instance.
func (s WKScriptMessage) Init() WKScriptMessage {
	rv := objc.Send[WKScriptMessage](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s WKScriptMessage) Autorelease() WKScriptMessage {
	rv := objc.Send[WKScriptMessage](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKScriptMessage creates a new WKScriptMessage instance.
func NewWKScriptMessage() WKScriptMessage {
	class := getWKScriptMessageClass()
	rv := objc.Send[WKScriptMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The body of the message.
//
// # Discussion
//
// Allowed types are [NSNumber], [NSString], [NSDate], [NSArray],
// [NSDictionary], and [NSNull].
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage/body
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (s WKScriptMessage) Body() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("body"))
	return objectivec.Object{ID: rv}
}

// The frame that sent the message.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage/frameInfo
func (s WKScriptMessage) FrameInfo() IWKFrameInfo {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("frameInfo"))
	return WKFrameInfoFromID(objc.ID(rv))
}

// The web view that sent the message.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage/webView
func (s WKScriptMessage) WebView() IWKWebView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("webView"))
	return WKWebViewFromID(objc.ID(rv))
}

// The namespace in which the JavaScript code executes.
//
// # Discussion
//
// For more information about content worlds, see [WKContentWorld].
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage/world
func (s WKScriptMessage) World() IWKContentWorld {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("world"))
	return WKContentWorldFromID(objc.ID(rv))
}

// The name of the message handler to which the message is sent.
//
// See: https://developer.apple.com/documentation/WebKit/WKScriptMessage/name
func (s WKScriptMessage) Name() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
