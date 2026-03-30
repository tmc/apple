// Code generated from Apple documentation for WebKit. DO NOT EDIT.
//go:build ios
// +build ios

package webkit

import (
	"github.com/tmc/apple/objc"
)

// # Discussion
//
// A Boolean value indicating whether LookToScroll is enabled.
//
// The default value is [NO].
//
// See: https://developer.apple.com/documentation/WebKit/WKPreferences/isLookToScrollEnabled
func (p WKPreferences) IsLookToScrollEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isLookToScrollEnabled"))
	return rv
}
func (p WKPreferences) SetIsLookToScrollEnabled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setIsLookToScrollEnabled:"), value)
}
