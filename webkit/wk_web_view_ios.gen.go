// Code generated from Apple documentation for WebKit. DO NOT EDIT.
//go:build ios
// +build ios

package webkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The scroll view associated with the web view.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebView/scrollView
func (w WKWebView) ScrollView() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("scrollView"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/isFindInteractionEnabled
func (w WKWebView) FindInteractionEnabled() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isFindInteractionEnabled"))
	return rv
}
func (w WKWebView) SetFindInteractionEnabled(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setFindInteractionEnabled:"), value)
}

// See: https://developer.apple.com/documentation/WebKit/WKWebView/findInteraction
func (w WKWebView) FindInteraction() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("findInteraction"))
	return objectivec.Object{ID: rv}
}
