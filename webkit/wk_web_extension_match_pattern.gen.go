// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKWebExtensionMatchPattern] class.
var (
	_WKWebExtensionMatchPatternClass     WKWebExtensionMatchPatternClass
	_WKWebExtensionMatchPatternClassOnce sync.Once
)

func getWKWebExtensionMatchPatternClass() WKWebExtensionMatchPatternClass {
	_WKWebExtensionMatchPatternClassOnce.Do(func() {
		_WKWebExtensionMatchPatternClass = WKWebExtensionMatchPatternClass{class: objc.GetClass("WKWebExtensionMatchPattern")}
	})
	return _WKWebExtensionMatchPatternClass
}

// GetWKWebExtensionMatchPatternClass returns the class object for WKWebExtensionMatchPattern.
func GetWKWebExtensionMatchPatternClass() WKWebExtensionMatchPatternClass {
	return getWKWebExtensionMatchPatternClass()
}

type WKWebExtensionMatchPatternClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKWebExtensionMatchPatternClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKWebExtensionMatchPatternClass) Alloc() WKWebExtensionMatchPattern {
	rv := objc.Send[WKWebExtensionMatchPattern](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a way to specify groups of URLs.
//
// # Overview
//
// All match patterns are specified as strings. Apart from the special “
// pattern, match patterns consist of three parts: scheme, host, and path.
//
// # Initializers
//
//   - [WKWebExtensionMatchPattern.InitWithSchemeHostPathError]: Returns a pattern object for the specified scheme, host, and path strings.
//   - [WKWebExtensionMatchPattern.InitWithStringError]: Returns a pattern object for the specified pattern string.
//
// # Instance Properties
//
//   - [WKWebExtensionMatchPattern.Host]: The host part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
//   - [WKWebExtensionMatchPattern.MatchesAllHosts]: A Boolean value that indicates if the pattern is `<all_urls>` or has `*` as the host.
//   - [WKWebExtensionMatchPattern.MatchesAllURLs]: A Boolean value that indicates if the pattern is `<all_urls>`.
//   - [WKWebExtensionMatchPattern.Path]: The path part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
//   - [WKWebExtensionMatchPattern.Scheme]: The scheme part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
//   - [WKWebExtensionMatchPattern.String]: The original pattern string.
//
// # Instance Methods
//
//   - [WKWebExtensionMatchPattern.MatchesURL]: Matches the receiver pattern against the specified URL.
//   - [WKWebExtensionMatchPattern.MatchesPattern]: Matches the receiver pattern against the specified pattern.
//   - [WKWebExtensionMatchPattern.MatchesURLOptions]: Matches the receiver pattern against the specified URL with options.
//   - [WKWebExtensionMatchPattern.MatchesPatternOptions]: Matches the receiver pattern against the specified pattern with options.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern
type WKWebExtensionMatchPattern struct {
	objectivec.Object
}

// WKWebExtensionMatchPatternFromID constructs a [WKWebExtensionMatchPattern] from an objc.ID.
//
// An object that represents a way to specify groups of URLs.
func WKWebExtensionMatchPatternFromID(id objc.ID) WKWebExtensionMatchPattern {
	return WKWebExtensionMatchPattern{objectivec.Object{ID: id}}
}

// NOTE: WKWebExtensionMatchPattern adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKWebExtensionMatchPattern] class.
//
// # Initializers
//
//   - [IWKWebExtensionMatchPattern.InitWithSchemeHostPathError]: Returns a pattern object for the specified scheme, host, and path strings.
//   - [IWKWebExtensionMatchPattern.InitWithStringError]: Returns a pattern object for the specified pattern string.
//
// # Instance Properties
//
//   - [IWKWebExtensionMatchPattern.Host]: The host part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
//   - [IWKWebExtensionMatchPattern.MatchesAllHosts]: A Boolean value that indicates if the pattern is `<all_urls>` or has `*` as the host.
//   - [IWKWebExtensionMatchPattern.MatchesAllURLs]: A Boolean value that indicates if the pattern is `<all_urls>`.
//   - [IWKWebExtensionMatchPattern.Path]: The path part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
//   - [IWKWebExtensionMatchPattern.Scheme]: The scheme part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
//   - [IWKWebExtensionMatchPattern.String]: The original pattern string.
//
// # Instance Methods
//
//   - [IWKWebExtensionMatchPattern.MatchesURL]: Matches the receiver pattern against the specified URL.
//   - [IWKWebExtensionMatchPattern.MatchesPattern]: Matches the receiver pattern against the specified pattern.
//   - [IWKWebExtensionMatchPattern.MatchesURLOptions]: Matches the receiver pattern against the specified URL with options.
//   - [IWKWebExtensionMatchPattern.MatchesPatternOptions]: Matches the receiver pattern against the specified pattern with options.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern
type IWKWebExtensionMatchPattern interface {
	objectivec.IObject

	// Topic: Initializers

	// Returns a pattern object for the specified scheme, host, and path strings.
	InitWithSchemeHostPathError(scheme string, host string, path string) (WKWebExtensionMatchPattern, error)
	// Returns a pattern object for the specified pattern string.
	InitWithStringError(string_ string) (WKWebExtensionMatchPattern, error)

	// Topic: Instance Properties

	// The host part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
	Host() string
	// A Boolean value that indicates if the pattern is `<all_urls>` or has `*` as the host.
	MatchesAllHosts() bool
	// A Boolean value that indicates if the pattern is `<all_urls>`.
	MatchesAllURLs() bool
	// The path part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
	Path() string
	// The scheme part of the pattern string, unless [matchesAllURLs](<doc://com.apple.webkit/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs>) is [YES].
	Scheme() string
	// The original pattern string.
	String() string

	// Topic: Instance Methods

	// Matches the receiver pattern against the specified URL.
	MatchesURL(url foundation.INSURL) bool
	// Matches the receiver pattern against the specified pattern.
	MatchesPattern(pattern IWKWebExtensionMatchPattern) bool
	// Matches the receiver pattern against the specified URL with options.
	MatchesURLOptions(url foundation.INSURL, options WKWebExtensionMatchPatternOptions) bool
	// Matches the receiver pattern against the specified pattern with options.
	MatchesPatternOptions(pattern IWKWebExtensionMatchPattern, options WKWebExtensionMatchPatternOptions) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (w WKWebExtensionMatchPattern) Init() WKWebExtensionMatchPattern {
	rv := objc.Send[WKWebExtensionMatchPattern](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w WKWebExtensionMatchPattern) Autorelease() WKWebExtensionMatchPattern {
	rv := objc.Send[WKWebExtensionMatchPattern](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKWebExtensionMatchPattern creates a new WKWebExtensionMatchPattern instance.
func NewWKWebExtensionMatchPattern() WKWebExtensionMatchPattern {
	class := getWKWebExtensionMatchPatternClass()
	rv := objc.Send[WKWebExtensionMatchPattern](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a pattern object for the specified scheme, host, and path strings.
//
// scheme: On output, a pointer to an error object that describes why the method
// failed, or `nil` if no error occurred. If you are not interested in the
// error information, pass `nil` for this parameter.
//
// # Return Value
//
// A pattern object, or `nil` if any of the strings are invalid and an error
// will be set.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/init(scheme:host:path:)
func NewWebExtensionMatchPatternWithSchemeHostPathError(scheme string, host string, path string) (WKWebExtensionMatchPattern, error) {
	var errorPtr objc.ID
	instance := getWKWebExtensionMatchPatternClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScheme:host:path:error:"), objc.String(scheme), objc.String(host), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return WKWebExtensionMatchPattern{}, foundation.NSErrorFrom(errorPtr)
	}
	return WKWebExtensionMatchPatternFromID(rv), nil
}

// Returns a pattern object for the specified pattern string.
//
// string: On output, a pointer to an error object that describes why the method
// failed, or `nil` if no error occurred. If you are not interested in the
// error information, pass `nil` for this parameter.
//
// # Return Value
//
// A pattern object, or `nil` if the pattern string is invalid and an error
// will be set.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/init(string:)
func NewWebExtensionMatchPatternWithStringError(string_ string) (WKWebExtensionMatchPattern, error) {
	var errorPtr objc.ID
	instance := getWKWebExtensionMatchPatternClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:error:"), objc.String(string_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return WKWebExtensionMatchPattern{}, foundation.NSErrorFrom(errorPtr)
	}
	return WKWebExtensionMatchPatternFromID(rv), nil
}

// Returns a pattern object for the specified scheme, host, and path strings.
//
// scheme: On output, a pointer to an error object that describes why the method
// failed, or `nil` if no error occurred. If you are not interested in the
// error information, pass `nil` for this parameter.
//
// # Return Value
//
// A pattern object, or `nil` if any of the strings are invalid and an error
// will be set.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/init(scheme:host:path:)
func (w WKWebExtensionMatchPattern) InitWithSchemeHostPathError(scheme string, host string, path string) (WKWebExtensionMatchPattern, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("initWithScheme:host:path:error:"), objc.String(scheme), objc.String(host), objc.String(path), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return WKWebExtensionMatchPattern{}, foundation.NSErrorFrom(errorPtr)
	}
	return WKWebExtensionMatchPatternFromID(rv), nil

}

// Returns a pattern object for the specified pattern string.
//
// string: On output, a pointer to an error object that describes why the method
// failed, or `nil` if no error occurred. If you are not interested in the
// error information, pass `nil` for this parameter.
//
// # Return Value
//
// A pattern object, or `nil` if the pattern string is invalid and an error
// will be set.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/init(string:)
func (w WKWebExtensionMatchPattern) InitWithStringError(string_ string) (WKWebExtensionMatchPattern, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](w.ID, objc.Sel("initWithString:error:"), objc.String(string_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return WKWebExtensionMatchPattern{}, foundation.NSErrorFrom(errorPtr)
	}
	return WKWebExtensionMatchPatternFromID(rv), nil

}

// Matches the receiver pattern against the specified URL.
//
// url: The URL to match against the receiver pattern.
//
// # Return Value
//
// A Boolean value indicating if the pattern matches the specified URL.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:)-471rf
func (w WKWebExtensionMatchPattern) MatchesURL(url foundation.INSURL) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("matchesURL:"), url)
	return rv
}

// Matches the receiver pattern against the specified pattern.
//
// pattern: The pattern to match against the receiver pattern.
//
// # Return Value
//
// A Boolean value indicating if the receiver pattern matches the specified
// pattern.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:)-4d84f
func (w WKWebExtensionMatchPattern) MatchesPattern(pattern IWKWebExtensionMatchPattern) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("matchesPattern:"), pattern)
	return rv
}

// Matches the receiver pattern against the specified URL with options.
//
// url: The URL to match against the receiver pattern.
//
// options: The options to use while matching.
//
// # Return Value
//
// A Boolean value indicating if the pattern matches the specified URL.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:options:)-5wo3g
func (w WKWebExtensionMatchPattern) MatchesURLOptions(url foundation.INSURL, options WKWebExtensionMatchPatternOptions) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("matchesURL:options:"), url, options)
	return rv
}

// Matches the receiver pattern against the specified pattern with options.
//
// pattern: The pattern to match against the receiver pattern.
//
// options: The options to use while matching.
//
// # Return Value
//
// A Boolean value indicating if the receiver pattern matches the specified
// pattern.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matches(_:options:)-fnde
func (w WKWebExtensionMatchPattern) MatchesPatternOptions(pattern IWKWebExtensionMatchPattern, options WKWebExtensionMatchPatternOptions) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("matchesPattern:options:"), pattern, options)
	return rv
}
func (w WKWebExtensionMatchPattern) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a pattern object that has `*` for scheme, host, and path.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/allHostsAndSchemes()
func (_WKWebExtensionMatchPatternClass WKWebExtensionMatchPatternClass) AllHostsAndSchemesMatchPattern() WKWebExtensionMatchPattern {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionMatchPatternClass.class), objc.Sel("allHostsAndSchemesMatchPattern"))
	return WKWebExtensionMatchPatternFromID(rv)
}

// Returns a pattern object for “.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/allURLs()
func (_WKWebExtensionMatchPatternClass WKWebExtensionMatchPatternClass) AllURLsMatchPattern() WKWebExtensionMatchPattern {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionMatchPatternClass.class), objc.Sel("allURLsMatchPattern"))
	return WKWebExtensionMatchPatternFromID(rv)
}

// Registers a custom URL scheme that can be used in match patterns.
//
// urlScheme: The custom URL scheme to register.
//
// # Discussion
//
// This method should be used to register any custom URL schemes used by the
// app for the extension base URLs, other than `webkit-extension`, or if
// extensions should have access to other supported URL schemes when using “.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/registerCustomURLScheme(_:)
func (_WKWebExtensionMatchPatternClass WKWebExtensionMatchPatternClass) RegisterCustomURLScheme(urlScheme string) {
	objc.Send[objc.ID](objc.ID(_WKWebExtensionMatchPatternClass.class), objc.Sel("registerCustomURLScheme:"), objc.String(urlScheme))
}

// Returns a pattern object for the specified scheme, host, and path strings.
//
// # Return Value
//
// A pattern object, or `nil` if any of the strings are invalid.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionMatchPattern/matchPatternWithScheme:host:path:
func (_WKWebExtensionMatchPatternClass WKWebExtensionMatchPatternClass) MatchPatternWithSchemeHostPath(scheme string, host string, path string) WKWebExtensionMatchPattern {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionMatchPatternClass.class), objc.Sel("matchPatternWithScheme:host:path:"), objc.String(scheme), objc.String(host), objc.String(path))
	return WKWebExtensionMatchPatternFromID(rv)
}

// Returns a pattern object for the specified pattern string.
//
// # Return Value
//
// Returns `nil` if the pattern string is invalid.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtensionMatchPattern/matchPatternWithString:
func (_WKWebExtensionMatchPatternClass WKWebExtensionMatchPatternClass) MatchPatternWithString(string_ string) WKWebExtensionMatchPattern {
	rv := objc.Send[objc.ID](objc.ID(_WKWebExtensionMatchPatternClass.class), objc.Sel("matchPatternWithString:"), objc.String(string_))
	return WKWebExtensionMatchPatternFromID(rv)
}

// The host part of the pattern string, unless [MatchesAllURLs] is [YES].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/host
func (w WKWebExtensionMatchPattern) Host() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("host"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates if the pattern is “ or has `*` as the host.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllHosts
func (w WKWebExtensionMatchPattern) MatchesAllHosts() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("matchesAllHosts"))
	return rv
}

// A Boolean value that indicates if the pattern is “.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/matchesAllURLs
func (w WKWebExtensionMatchPattern) MatchesAllURLs() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("matchesAllURLs"))
	return rv
}

// The path part of the pattern string, unless [MatchesAllURLs] is [YES].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/path
func (w WKWebExtensionMatchPattern) Path() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("path"))
	return foundation.NSStringFromID(rv).String()
}

// The scheme part of the pattern string, unless [MatchesAllURLs] is [YES].
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/scheme
func (w WKWebExtensionMatchPattern) Scheme() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("scheme"))
	return foundation.NSStringFromID(rv).String()
}

// The original pattern string.
//
// See: https://developer.apple.com/documentation/WebKit/WKWebExtension/MatchPattern/string
func (w WKWebExtensionMatchPattern) String() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("string"))
	return foundation.NSStringFromID(rv).String()
}
