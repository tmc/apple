// Code generated from Apple documentation for WebKit. DO NOT EDIT.
//go:build ios
// +build ios

package webkit

import (
	"github.com/tmc/apple/objc"
)

// A view controller that presents a web view loaded with the pop-up page for
// this action, or `nil` if no popup is specified.
//
// # Discussion
//
// The view controller adaptively adjusts its presentation style based on
// where it is presented from, preferring popover.
//
// It contains a web view preloaded with the pop-up page and automatically
// adjusts its `preferredContentSize` to fit the web view’s content size.
// The [PresentsPopup] property should be checked to determine the
// availability of a pop-up before using this property.
//
// Dismissing the view controller will close the pop-up and unload the web
// view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/Action/popupViewController
func (w WKWebExtensionAction) PopupViewController() objc.ID {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("popupViewController"))
	return rv
}
