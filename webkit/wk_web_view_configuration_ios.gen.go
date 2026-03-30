// Code generated from Apple documentation for WebKit. DO NOT EDIT.
//go:build ios
// +build ios

package webkit

import (
	"github.com/tmc/apple/objc"
)

// A Boolean value that determines whether a web view allows scaling of the
// webpage.
//
// # Discussion
//
// When set to true, this property overrides the `user-scalable` HTML property
// in a webpage, and lets the web view scale its webpage content regardless of
// the author’s intent. The default value of this property is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/ignoresViewportScaleLimits
func (w WKWebViewConfiguration) IgnoresViewportScaleLimits() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("ignoresViewportScaleLimits"))
	return rv
}
func (w WKWebViewConfiguration) SetIgnoresViewportScaleLimits(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setIgnoresViewportScaleLimits:"), value)
}

// A Boolean value that indicates whether HTML5 videos play inline or use the
// native full-screen controller.
//
// # Discussion
//
// Set this property to true to play videos inline, or false to use the native
// full-screen controller. When adding a video element to an HTML document on
// iPhone, you must also include the `playsinline` attribute.
//
// The default value of this property is false for iPhone and true for iPad.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsInlineMediaPlayback
func (w WKWebViewConfiguration) AllowsInlineMediaPlayback() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsInlineMediaPlayback"))
	return rv
}
func (w WKWebViewConfiguration) SetAllowsInlineMediaPlayback(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsInlineMediaPlayback:"), value)
}

// A Boolean value that indicates whether HTML5 videos can play Picture in
// Picture.
//
// # Discussion
//
// The default value of this property is true.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/allowsPictureInPictureMediaPlayback
func (w WKWebViewConfiguration) AllowsPictureInPictureMediaPlayback() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsPictureInPictureMediaPlayback"))
	return rv
}
func (w WKWebViewConfiguration) SetAllowsPictureInPictureMediaPlayback(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsPictureInPictureMediaPlayback:"), value)
}

// The types of data detectors to apply to the web view’s content.
//
// # Discussion
//
// Data detectors add interactivity to web content by creating links for
// specially formatted text. For example, the [WKDataDetectorTypeLink] type
// causes the `apple.Com()` portion of the text “Visit apple.com” to
// become a link to the Apple website.
//
// The default value of this property is [WKDataDetectorTypeNone]. For other
// possible values, see [WKDataDetectorTypes].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/dataDetectorTypes
//
// [WKDataDetectorTypes]: https://developer.apple.com/documentation/WebKit/WKDataDetectorTypes
func (w WKWebViewConfiguration) DataDetectorTypes() WKDataDetectorTypes {
	rv := objc.Send[WKDataDetectorTypes](w.ID, objc.Sel("dataDetectorTypes"))
	return WKDataDetectorTypes(rv)
}
func (w WKWebViewConfiguration) SetDataDetectorTypes(value WKDataDetectorTypes) {
	objc.Send[struct{}](w.ID, objc.Sel("setDataDetectorTypes:"), value)
}

// The level of granularity with which the user can interactively select web
// view content.
//
// # Discussion
//
// The value is one of the constants of the enumerated type
// [WKSelectionGranularity]. The default value is
// [WKSelectionGranularityDynamic].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/selectionGranularity
//
// [WKSelectionGranularity]: https://developer.apple.com/documentation/WebKit/WKSelectionGranularity
func (w WKWebViewConfiguration) SelectionGranularity() WKSelectionGranularity {
	rv := objc.Send[WKSelectionGranularity](w.ID, objc.Sel("selectionGranularity"))
	return WKSelectionGranularity(rv)
}
func (w WKWebViewConfiguration) SetSelectionGranularity(value WKSelectionGranularity) {
	objc.Send[struct{}](w.ID, objc.Sel("setSelectionGranularity:"), value)
}

// Deprecated property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaPlaybackAllowsAirPlay
func (w WKWebViewConfiguration) MediaPlaybackAllowsAirPlay() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("mediaPlaybackAllowsAirPlay"))
	return rv
}
func (w WKWebViewConfiguration) SetMediaPlaybackAllowsAirPlay(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setMediaPlaybackAllowsAirPlay:"), value)
}

// A Boolean value that indicates whether HTML5 videos require the user to
// start playing them (`true`) or whether the videos play automatically
// (`false`).
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/requiresUserActionForMediaPlayback
func (w WKWebViewConfiguration) RequiresUserActionForMediaPlayback() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("requiresUserActionForMediaPlayback"))
	return rv
}
func (w WKWebViewConfiguration) SetRequiresUserActionForMediaPlayback(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setRequiresUserActionForMediaPlayback:"), value)
}

// Deprecated property.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebViewConfiguration/mediaPlaybackRequiresUserAction
func (w WKWebViewConfiguration) MediaPlaybackRequiresUserAction() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("mediaPlaybackRequiresUserAction"))
	return rv
}
func (w WKWebViewConfiguration) SetMediaPlaybackRequiresUserAction(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setMediaPlaybackRequiresUserAction:"), value)
}
