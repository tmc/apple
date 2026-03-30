// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKFrameInfo] class.
var (
	_WKFrameInfoClass     WKFrameInfoClass
	_WKFrameInfoClassOnce sync.Once
)

func getWKFrameInfoClass() WKFrameInfoClass {
	_WKFrameInfoClassOnce.Do(func() {
		_WKFrameInfoClass = WKFrameInfoClass{class: objc.GetClass("WKFrameInfo")}
	})
	return _WKFrameInfoClass
}

// GetWKFrameInfoClass returns the class object for WKFrameInfo.
func GetWKFrameInfoClass() WKFrameInfoClass {
	return getWKFrameInfoClass()
}

type WKFrameInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKFrameInfoClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKFrameInfoClass) Alloc() WKFrameInfo {
	rv := objc.Send[WKFrameInfo](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains information about a frame on a webpage.
//
// # Overview
//
// An instance of this class is a transient, data-only object; it does not
// uniquely identify a frame across multiple delegate method calls.
//
// # Inspecting frame information
//
//   - [WKFrameInfo.MainFrame]: A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//   - [WKFrameInfo.Request]: The frame’s current request.
//   - [WKFrameInfo.SecurityOrigin]: The frame’s security origin.
//   - [WKFrameInfo.WebView]: The web view that contains this frame and the containing webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKFrameInfo
type WKFrameInfo struct {
	objectivec.Object
}

// WKFrameInfoFromID constructs a [WKFrameInfo] from an objc.ID.
//
// An object that contains information about a frame on a webpage.
func WKFrameInfoFromID(id objc.ID) WKFrameInfo {
	return WKFrameInfo{objectivec.Object{ID: id}}
}

// NOTE: WKFrameInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKFrameInfo] class.
//
// # Inspecting frame information
//
//   - [IWKFrameInfo.MainFrame]: A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
//   - [IWKFrameInfo.Request]: The frame’s current request.
//   - [IWKFrameInfo.SecurityOrigin]: The frame’s security origin.
//   - [IWKFrameInfo.WebView]: The web view that contains this frame and the containing webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKFrameInfo
type IWKFrameInfo interface {
	objectivec.IObject

	// Topic: Inspecting frame information

	// A Boolean value indicating whether the frame is the web site’s main frame or a subframe.
	MainFrame() bool
	// The frame’s current request.
	Request() foundation.NSURLRequest
	// The frame’s security origin.
	SecurityOrigin() IWKSecurityOrigin
	// The web view that contains this frame and the containing webpage.
	WebView() IWKWebView
}

// Init initializes the instance.
func (f WKFrameInfo) Init() WKFrameInfo {
	rv := objc.Send[WKFrameInfo](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f WKFrameInfo) Autorelease() WKFrameInfo {
	rv := objc.Send[WKFrameInfo](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKFrameInfo creates a new WKFrameInfo instance.
func NewWKFrameInfo() WKFrameInfo {
	class := getWKFrameInfoClass()
	rv := objc.Send[WKFrameInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value indicating whether the frame is the web site’s main frame
// or a subframe.
//
// See: https://developer.apple.com/documentation/WebKit/WKFrameInfo/isMainFrame
func (f WKFrameInfo) MainFrame() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isMainFrame"))
	return rv
}

// The frame’s current request.
//
// See: https://developer.apple.com/documentation/WebKit/WKFrameInfo/request
func (f WKFrameInfo) Request() foundation.NSURLRequest {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("request"))
	return foundation.NSURLRequestFromID(objc.ID(rv))
}

// The frame’s security origin.
//
// # Discussion
//
// The [WKSecurityOrigin] object consists of a host name, a protocol, and a
// port number.
//
// See: https://developer.apple.com/documentation/WebKit/WKFrameInfo/securityOrigin
func (f WKFrameInfo) SecurityOrigin() IWKSecurityOrigin {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("securityOrigin"))
	return WKSecurityOriginFromID(objc.ID(rv))
}

// The web view that contains this frame and the containing webpage.
//
// See: https://developer.apple.com/documentation/WebKit/WKFrameInfo/webView
func (f WKFrameInfo) WebView() IWKWebView {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("webView"))
	return WKWebViewFromID(objc.ID(rv))
}
