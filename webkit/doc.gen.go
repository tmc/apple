// Code generated from Apple documentation for WebKit. DO NOT EDIT.

// Package webkit provides Go bindings for the WebKit framework.
//
// Integrate web content seamlessly into your app, and customize content interactions to meet your app’s needs.
//
// Use the WebKit framework to integrate richly styled web content into your app’s native content. WebKit offers a full browsing experience for your content, offering a platform-native view and supporting classes to:
//
// # WebKit APIs
//
//   - WebKit for AppKit and UIKit: Display web content in AppKit or UIKit apps, or apps built with Objective-C. ([WKWebView], [WKUIDelegate], [WKWebViewConfiguration], [WKWindowFeatures], [WKPreferences])
//   - WebKit for SwiftUI: Integrate web content into your SwiftUI apps with new standard views you connect to webpages. ([WebView], [URLSchemeHandler])
//
// # Safari Support
//
//   - Optimizing Your Website for Safari: Improve your website by optimizing it for Safari.
//   - Delivering Video Content for Safari: Improve the performance and appearance of video in your website in Safari.
//   - Promoting Apps with Smart App Banners: Create a banner to promote your app on the App Store from a website.
//
// # WebDriver
//
//   - macOS WebDriver Commands for Safari 11.1 and earlier: Test your web content using the WebDriver commands supported by Safari 11.1 and earlier.
//   - macOS WebDriver Commands for Safari 12 and later: Test your web content using the WebDriver commands supported by Safari 12 and later.
//   - About WebDriver for Safari: Enhance testing of your web content using Safari’s enhancements to WebDriver.
//   - Testing with WebDriver in Safari: Enable WebDriver and run a test.
//
// # Enumerations
//
//   - WKSecurityRestrictionMode
//
// # Key Types
//
//   - [WKWebView] - An object that displays interactive web content, such as for an in-app browser.
//   - [WKWebExtensionContext] - An object that represents the runtime environment for a web extension.
//   - [WKWebExtension] - An object that encapsulates a web extension’s resources that the manifest file defines.
//   - [WKWebExtensionController] - An object that manages a set of loaded extension contexts.
//   - [WKWebViewConfiguration] - A collection of properties that you use to initialize a web view.
//   - [WKWebExtensionMatchPattern] - An object that represents a way to specify groups of URLs.
//   - [WKWebsiteDataStore] - An object that manages cookies, disk and memory caches, and other types of data for a web view.
//   - [WKUserContentController] - An object for managing interactions between JavaScript code and your web view, and for filtering content in your web view.
//   - [WKWebExtensionAction] - An object that encapsulates the properties for an individual web extension action.
//   - [WKPreferences] - An object that encapsulates the standard behaviors to apply to websites.
//
// [WebKit Documentation]: https://developer.apple.com/documentation/WebKit
package webkit

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the WebKit library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/WebKit.framework/WebKit",
	"/usr/lib/libWebKit.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: WebKit: failed to load framework from any known path\n")
}
