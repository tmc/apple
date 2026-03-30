// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKDownload] class.
var (
	_WKDownloadClass     WKDownloadClass
	_WKDownloadClassOnce sync.Once
)

func getWKDownloadClass() WKDownloadClass {
	_WKDownloadClassOnce.Do(func() {
		_WKDownloadClass = WKDownloadClass{class: objc.GetClass("WKDownload")}
	})
	return _WKDownloadClass
}

// GetWKDownloadClass returns the class object for WKDownload.
func GetWKDownloadClass() WKDownloadClass {
	return getWKDownloadClass()
}

type WKDownloadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKDownloadClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKDownloadClass) Alloc() WKDownload {
	rv := objc.Send[WKDownload](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents the download of a web resource.
//
// # Managing the download
//
//   - [WKDownload.Delegate]: An object you use to track download progress and handle redirects, authentication challenges, and failures.
//   - [WKDownload.SetDelegate]
//   - [WKDownload.Cancel]: Cancels the download, and optionally captures data so that you can resume the download later.
//
// # Inspecting the download
//
//   - [WKDownload.OriginalRequest]: An object that represents the request that initiated the download.
//   - [WKDownload.WebView]: The web view where the download initiated.
//
// # Instance Properties
//
//   - [WKDownload.UserInitiated]
//   - [WKDownload.OriginatingFrame]
//
// See: https://developer.apple.com/documentation/WebKit/WKDownload
type WKDownload struct {
	objectivec.Object
}

// WKDownloadFromID constructs a [WKDownload] from an objc.ID.
//
// An object that represents the download of a web resource.
func WKDownloadFromID(id objc.ID) WKDownload {
	return WKDownload{objectivec.Object{ID: id}}
}

// NOTE: WKDownload adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKDownload] class.
//
// # Managing the download
//
//   - [IWKDownload.Delegate]: An object you use to track download progress and handle redirects, authentication challenges, and failures.
//   - [IWKDownload.SetDelegate]
//   - [IWKDownload.Cancel]: Cancels the download, and optionally captures data so that you can resume the download later.
//
// # Inspecting the download
//
//   - [IWKDownload.OriginalRequest]: An object that represents the request that initiated the download.
//   - [IWKDownload.WebView]: The web view where the download initiated.
//
// # Instance Properties
//
//   - [IWKDownload.UserInitiated]
//   - [IWKDownload.OriginatingFrame]
//
// See: https://developer.apple.com/documentation/WebKit/WKDownload
type IWKDownload interface {
	objectivec.IObject

	// Topic: Managing the download

	// An object you use to track download progress and handle redirects, authentication challenges, and failures.
	Delegate() WKDownloadDelegate
	SetDelegate(value WKDownloadDelegate)
	// Cancels the download, and optionally captures data so that you can resume the download later.
	Cancel(completionHandler DataHandler)

	// Topic: Inspecting the download

	// An object that represents the request that initiated the download.
	OriginalRequest() foundation.NSURLRequest
	// The web view where the download initiated.
	WebView() IWKWebView

	// Topic: Instance Properties

	UserInitiated() bool
	OriginatingFrame() IWKFrameInfo
}

// Init initializes the instance.
func (d WKDownload) Init() WKDownload {
	rv := objc.Send[WKDownload](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d WKDownload) Autorelease() WKDownload {
	rv := objc.Send[WKDownload](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKDownload creates a new WKDownload instance.
func NewWKDownload() WKDownload {
	class := getWKDownloadClass()
	rv := objc.Send[WKDownload](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Cancels the download, and optionally captures data so that you can resume
// the download later.
//
// completionHandler: A closure you provide to capture and store data so that you can resume the
// download later.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownload/cancel(_:)
func (d WKDownload) Cancel(completionHandler DataHandler) {
	_block0, _ := NewDataBlock(completionHandler)
	objc.Send[objc.ID](d.ID, objc.Sel("cancel:"), _block0)
}

// An object you use to track download progress and handle redirects,
// authentication challenges, and failures.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownload/delegate
func (d WKDownload) Delegate() WKDownloadDelegate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("delegate"))
	return WKDownloadDelegateObjectFromID(rv)
}
func (d WKDownload) SetDelegate(value WKDownloadDelegate) {
	objc.Send[struct{}](d.ID, objc.Sel("setDelegate:"), value)
}

// An object that represents the request that initiated the download.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownload/originalRequest
func (d WKDownload) OriginalRequest() foundation.NSURLRequest {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("originalRequest"))
	return foundation.NSURLRequestFromID(objc.ID(rv))
}

// The web view where the download initiated.
//
// See: https://developer.apple.com/documentation/WebKit/WKDownload/webView
func (d WKDownload) WebView() IWKWebView {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("webView"))
	return WKWebViewFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/WebKit/WKDownload/isUserInitiated
func (d WKDownload) UserInitiated() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isUserInitiated"))
	return rv
}

// See: https://developer.apple.com/documentation/WebKit/WKDownload/originatingFrame
func (d WKDownload) OriginatingFrame() IWKFrameInfo {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("originatingFrame"))
	return WKFrameInfoFromID(objc.ID(rv))
}

// CancelSync is a synchronous wrapper around [WKDownload.Cancel].
// It blocks until the completion handler fires or the context is cancelled.
func (d WKDownload) CancelSync(ctx context.Context) (*foundation.NSData, error) {
	done := make(chan *foundation.NSData, 1)
	d.Cancel(func(val *foundation.NSData) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
