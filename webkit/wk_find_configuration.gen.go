// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [WKFindConfiguration] class.
var (
	_WKFindConfigurationClass     WKFindConfigurationClass
	_WKFindConfigurationClassOnce sync.Once
)

func getWKFindConfigurationClass() WKFindConfigurationClass {
	_WKFindConfigurationClassOnce.Do(func() {
		_WKFindConfigurationClass = WKFindConfigurationClass{class: objc.GetClass("WKFindConfiguration")}
	})
	return _WKFindConfigurationClass
}

// GetWKFindConfigurationClass returns the class object for WKFindConfiguration.
func GetWKFindConfigurationClass() WKFindConfigurationClass {
	return getWKFindConfigurationClass()
}

type WKFindConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (wc WKFindConfigurationClass) Class() objc.Class {
	return wc.class
}

// Alloc allocates memory for a new instance of the class.
func (wc WKFindConfigurationClass) Alloc() WKFindConfiguration {
	rv := objc.Send[WKFindConfiguration](objc.ID(wc.class), objc.Sel("alloc"))
	return rv
}

// The configuration parameters to use when searching the contents of the web
// view.
//
// # Overview
//
// Create a [WKFindConfiguration] object and configure its attributes to
// specify how to perform searches within the web view’s contents. To
// initiate a search, call the appropriate method of [WKWebView] and pass this
// object along with the search string.
//
// # Configuring the Search Parameters
//
//   - [WKFindConfiguration.Backwards]: A Boolean value that indicates the search direction, relative to the current selection.
//   - [WKFindConfiguration.SetBackwards]
//   - [WKFindConfiguration.CaseSensitive]: A Boolean value that indicates whether to consider case when matching the search string.
//   - [WKFindConfiguration.SetCaseSensitive]
//   - [WKFindConfiguration.Wraps]: A Boolean value that indicates whether the search wraps around to the other side of the page.
//   - [WKFindConfiguration.SetWraps]
//
// See: https://developer.apple.com/documentation/WebKit/WKFindConfiguration
type WKFindConfiguration struct {
	objectivec.Object
}

// WKFindConfigurationFromID constructs a [WKFindConfiguration] from an objc.ID.
//
// The configuration parameters to use when searching the contents of the web
// view.
func WKFindConfigurationFromID(id objc.ID) WKFindConfiguration {
	return WKFindConfiguration{objectivec.Object{ID: id}}
}

// NOTE: WKFindConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [WKFindConfiguration] class.
//
// # Configuring the Search Parameters
//
//   - [IWKFindConfiguration.Backwards]: A Boolean value that indicates the search direction, relative to the current selection.
//   - [IWKFindConfiguration.SetBackwards]
//   - [IWKFindConfiguration.CaseSensitive]: A Boolean value that indicates whether to consider case when matching the search string.
//   - [IWKFindConfiguration.SetCaseSensitive]
//   - [IWKFindConfiguration.Wraps]: A Boolean value that indicates whether the search wraps around to the other side of the page.
//   - [IWKFindConfiguration.SetWraps]
//
// See: https://developer.apple.com/documentation/WebKit/WKFindConfiguration
type IWKFindConfiguration interface {
	objectivec.IObject

	// Topic: Configuring the Search Parameters

	// A Boolean value that indicates the search direction, relative to the current selection.
	Backwards() bool
	SetBackwards(value bool)
	// A Boolean value that indicates whether to consider case when matching the search string.
	CaseSensitive() bool
	SetCaseSensitive(value bool)
	// A Boolean value that indicates whether the search wraps around to the other side of the page.
	Wraps() bool
	SetWraps(value bool)
}

// Init initializes the instance.
func (f WKFindConfiguration) Init() WKFindConfiguration {
	rv := objc.Send[WKFindConfiguration](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f WKFindConfiguration) Autorelease() WKFindConfiguration {
	rv := objc.Send[WKFindConfiguration](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewWKFindConfiguration creates a new WKFindConfiguration instance.
func NewWKFindConfiguration() WKFindConfiguration {
	class := getWKFindConfigurationClass()
	rv := objc.Send[WKFindConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates the search direction, relative to the
// current selection.
//
// # Discussion
//
// When the value of this property is true, searches proceed backward from the
// current selection. When the value is false, searches proceed forward from
// the current selection. The default value of this property is false.
//
// The web view respects the writing direction of its content. For example, a
// forward search moves right-to-left and top-to-bottom for a right-to-left
// language.
//
// See: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/backwards
func (f WKFindConfiguration) Backwards() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("backwards"))
	return rv
}
func (f WKFindConfiguration) SetBackwards(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setBackwards:"), value)
}

// A Boolean value that indicates whether to consider case when matching the
// search string.
//
// # Discussion
//
// When the value of this property is true, the web view takes case into
// account when matching the search string. The default value of this property
// is false.
//
// See: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/caseSensitive
func (f WKFindConfiguration) CaseSensitive() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("caseSensitive"))
	return rv
}
func (f WKFindConfiguration) SetCaseSensitive(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setCaseSensitive:"), value)
}

// A Boolean value that indicates whether the search wraps around to the other
// side of the page.
//
// # Discussion
//
// When the value of this property is true, the search wraps and continues at
// the other end of the page. For example, when a forward search reaches the
// bottom of the page, the search wraps back to the top of the page and
// continues. When a backward search reaches the top of the page, the search
// wraps back to the bottom of the page. The default value of this property is
// true.
//
// See: https://developer.apple.com/documentation/WebKit/WKFindConfiguration/wraps
func (f WKFindConfiguration) Wraps() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("wraps"))
	return rv
}
func (f WKFindConfiguration) SetWraps(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setWraps:"), value)
}
