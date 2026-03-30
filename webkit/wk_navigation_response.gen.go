// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKNavigationResponse] class.
var (
	_WKNavigationResponseClass     WKNavigationResponseClass
	_WKNavigationResponseClassOnce sync.Once
)

func getWKNavigationResponseClass() WKNavigationResponseClass {
	_WKNavigationResponseClassOnce.Do(func() {
		_WKNavigationResponseClass = WKNavigationResponseClass{class: objc.GetClass("WKNavigationResponse")}
	})
	return _WKNavigationResponseClass
}

// GetWKNavigationResponseClass returns the class object for WKNavigationResponse.
func GetWKNavigationResponseClass() WKNavigationResponseClass {
	return getWKNavigationResponseClass()
}

type WKNavigationResponseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKNavigationResponseClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKNavigationResponseClass) Alloc() WKNavigationResponse {
	rv := objc.Send[WKNavigationResponse](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the response to a navigation request, and which you
// use to make navigation-related policy decisions.
//
// # Overview
//
// Use a [WKNavigationResponse] object to make policy decisions about whether
// to allow navigation within your app’s web view. You don’t create
// [WKNavigationResponse] objects directly. Instead, the web view creates them
// and delivers them to the appropriate delegate objects. Use the methods of
// your delegate to analyze the response and determine whether to allow the
// resulting navigation to occur.
//
// # Getting the Response Details
//
//   - [WKNavigationResponse.Response]: The frame’s response.
//
// # Getting Additional Response Information
//
//   - [WKNavigationResponse.CanShowMIMEType]: A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
//   - [WKNavigationResponse.ForMainFrame]: A Boolean value that indicates whether the response targets the web view’s main frame.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationResponse
type WKNavigationResponse struct {
	objectivec.Object
}

// WKNavigationResponseFromID constructs a [WKNavigationResponse] from an objc.ID.
//
// An object that contains the response to a navigation request, and which you
// use to make navigation-related policy decisions.
func WKNavigationResponseFromID(id objc.ID) WKNavigationResponse {
	return WKNavigationResponse{objectivec.Object{ID: id}}
}

// NOTE: WKNavigationResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKNavigationResponse] class.
//
// # Getting the Response Details
//
//   - [IWKNavigationResponse.Response]: The frame’s response.
//
// # Getting Additional Response Information
//
//   - [IWKNavigationResponse.CanShowMIMEType]: A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
//   - [IWKNavigationResponse.ForMainFrame]: A Boolean value that indicates whether the response targets the web view’s main frame.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationResponse
type IWKNavigationResponse interface {
	objectivec.IObject

	// Topic: Getting the Response Details

	// The frame’s response.
	Response() foundation.NSURLResponse

	// Topic: Getting Additional Response Information

	// A Boolean value that indicates whether WebKit is capable of displaying the response’s MIME type natively.
	CanShowMIMEType() bool
	// A Boolean value that indicates whether the response targets the web view’s main frame.
	ForMainFrame() bool
}

// Init initializes the instance.
func (n WKNavigationResponse) Init() WKNavigationResponse {
	rv := objc.Send[WKNavigationResponse](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n WKNavigationResponse) Autorelease() WKNavigationResponse {
	rv := objc.Send[WKNavigationResponse](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKNavigationResponse creates a new WKNavigationResponse instance.
func NewWKNavigationResponse() WKNavigationResponse {
	class := getWKNavigationResponseClass()
	rv := objc.Send[WKNavigationResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The frame’s response.
//
// # Discussion
//
// Allowing a navigation response with a MIME type that WebKit can’t display
// causes the navigation to fail.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/response
func (n WKNavigationResponse) Response() foundation.NSURLResponse {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("response"))
	return foundation.NSURLResponseFromID(objc.ID(rv))
}

// A Boolean value that indicates whether WebKit is capable of displaying the
// response’s MIME type natively.
//
// # Discussion
//
// The value of this property is true if WebKit is able to display the MIME
// type.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/canShowMIMEType
func (n WKNavigationResponse) CanShowMIMEType() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("canShowMIMEType"))
	return rv
}

// A Boolean value that indicates whether the response targets the web
// view’s main frame.
//
// # Discussion
//
// The value of this property is true if the response targets the main frame.
// The value is false if the response targets a different frame, such as the
// frame in a new window.
//
// See: https://developer.apple.com/documentation/WebKit/WKNavigationResponse/isForMainFrame
func (n WKNavigationResponse) ForMainFrame() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("isForMainFrame"))
	return rv
}
