// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol for loading resources with URL schemes that WebKit doesn’t handle.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeHandler
type WKURLSchemeHandler interface {
	objectivec.IObject

	// Asks your handler to begin loading the data for the specified resource.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeHandler/webView(_:start:)
	WebViewStartURLSchemeTask(webView IWKWebView, urlSchemeTask WKURLSchemeTask)

	// Asks your handler to stop loading the data for the specified resource.
	//
	// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeHandler/webView(_:stop:)
	WebViewStopURLSchemeTask(webView IWKWebView, urlSchemeTask WKURLSchemeTask)
}

// WKURLSchemeHandlerObject wraps an existing Objective-C object that conforms to the WKURLSchemeHandler protocol.
type WKURLSchemeHandlerObject struct {
	objectivec.Object
}

func (o WKURLSchemeHandlerObject) BaseObject() objectivec.Object {
	return o.Object
}

// WKURLSchemeHandlerObjectFromID constructs a [WKURLSchemeHandlerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WKURLSchemeHandlerObjectFromID(id objc.ID) WKURLSchemeHandlerObject {
	return WKURLSchemeHandlerObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks your handler to begin loading the data for the specified resource.
//
// webView: The web view that requires the resource.
//
// urlSchemeTask: The task object that identifies the resource to load. You also use this
// object to report the progress of the load operation back to the web view.
//
// # Discussion
//
// When a web view encounters a resource with your custom URL scheme, it calls
// this method on the appropriate handler object. Use your implementation of
// this method to begin loading the resource. Call the methods of the provided
// [WKURLSchemeTask] object to report the progress of the loading operation
// back to the web view. You also use that object to deliver the resource data
// to the web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeHandler/webView(_:start:)
func (o WKURLSchemeHandlerObject) WebViewStartURLSchemeTask(webView IWKWebView, urlSchemeTask WKURLSchemeTask) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:startURLSchemeTask:"), webView, urlSchemeTask)
}

// Asks your handler to stop loading the data for the specified resource.
//
// webView: The web view that asked you to stop loading the resource.
//
// urlSchemeTask: The task object that identifies the resource the web view no longer needs.
//
// # Discussion
//
// If WebKit determines that it no longer needs a resource that your handler
// is loading, it calls this method to stop the load operation. Typically,
// WebKit calls this method when the user navigates to another page, but may
// call it for other reasons.
//
// When WebKit calls this method, stop the load operation immediately for the
// specified resource and free up any allocated memory for it. Don’t call
// any methods of the provided `urlSchemeTask` object to report your progress
// back to WebKit. If you do, the methods of that object raise an exception.
//
// See: https://developer.apple.com/documentation/WebKit/WKURLSchemeHandler/webView(_:stop:)
func (o WKURLSchemeHandlerObject) WebViewStopURLSchemeTask(webView IWKWebView, urlSchemeTask WKURLSchemeTask) {
	objc.Send[struct{}](o.ID, objc.Sel("webView:stopURLSchemeTask:"), webView, urlSchemeTask)
}
